package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

// 順番に従ってconnに書き出しをする（goroutineで実行される）
func writeToConn(sessionResponses chan chan *http.Response, conn net.Conn) {
	defer conn.Close()
	for sessionResponse := range sessionResponses {
		// 選択された仕事が終わるまで待つ
		resp := <-sessionResponse
		resp.Write(conn)
		close(sessionResponse)
	}
}

// セッション内のリクエストを処理する
func handleRequest(t *tokenizer.Tokenizer, req *http.Request, resultReceiver chan *http.Response) {
	text := req.URL.Query().Get("text")
	tokens := t.Tokenize(text)

	content := text + "\n"
	for _, token := range tokens {
		features := strings.Join(token.Features(), ",")
		content += fmt.Sprintf("%s\t%v\n", token.Surface, features)
	}

	resp := http.Response{
		StatusCode:    200,
		ProtoMajor:    1,
		ProtoMinor:    1,
		ContentLength: int64(len(content)),
		Body:          io.NopCloser(strings.NewReader(content)),
	}
	resultReceiver <- &resp
}

// セッション一つを処理
func processSession(t *tokenizer.Tokenizer, conn net.Conn) {
	fmt.Printf("Accept %v\n", conn.RemoteAddr())
	// セッション内のリクエストを順次処理するためのチャネル
	sessionResponses := make(chan chan *http.Response, 50)
	defer close(sessionResponses)
	go writeToConn(sessionResponses, conn)
	reader := bufio.NewReader(conn)
	// レスポンスを受け取ってセッションのキューに入れる
	for {
		conn.SetReadDeadline(time.Now().Add(time.Second * 5)) // 5秒でタイムアウト
		req, err := http.ReadRequest(reader)
		if err != nil {
			neterr, ok := err.(net.Error)
			// タイムアウトしたらそのことをクライアントに通知してからループを終了
			if ok && neterr.Timeout() {
				content := "Request timeout"
				resp := http.Response{
					StatusCode:    408,
					ProtoMajor:    1,
					ProtoMinor:    1,
					Body:          io.NopCloser(strings.NewReader(content)),
					ContentLength: int64(len(content)),
				}
				resp.Write(conn)
				break
			}
			// リクエストの最後まで来ていたらループを終了
			if err == io.EOF {
				break
			}
			panic(err)
		}

		// レスポンスを待ち受けるためのバッファなしチャネルを作成
		sessionResponse := make(chan *http.Response)
		// レスポンスが届いたらwriteToConnに届くようにする
		sessionResponses <- sessionResponse
		// 非同期で今回のリクエストを処理。次のリクエストへ
		go handleRequest(t, req, sessionResponse)
	}
}

func main() {
	t, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running at localhost:8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go processSession(t, conn)
	}
}
