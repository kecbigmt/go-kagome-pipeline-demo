package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

func main() {
	filepath := os.Args[1]
	if filepath == "" {
		panic("specify path of text file")
	}

	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("Access")
	defer conn.Close()

	var requests []*http.Request
	fileReader := bufio.NewReader(file)

	// 1ループ前に読み込んだ行を使ってリクエスト（最後の行でConnection closeするため）
	var line string
	var isFinal bool
	for {
		if line != "" {
			textParam := url.QueryEscape(strings.TrimLeft(line, "\n")) // 末尾の改行を削除してURLエンコーディング
			req, err := http.NewRequest("GET", "http://localhost:8888?text="+textParam, nil)
			if err != nil {
				panic(err)
			}
			if isFinal {
				req.Header.Set("Connection", "close")
			} else {
				req.Header.Set("Connection", "keep-alive")
			}
			err = req.Write(conn)
			requests = append(requests, req)
		}

		if isFinal {
			break
		}

		// 次のループでリクエストする行を読み込み
		line, err = fileReader.ReadString('\n')
		if err != nil {
			// EOFであれば次のループを最後にする
			if err == io.EOF {
				isFinal = true
			} else {
				panic(err)
			}
		}
	}

	// リクエストを受信
	connReader := bufio.NewReader(conn)
	for _, request := range requests {
		resp, err := http.ReadResponse(connReader, request)
		if err != nil {
			panic(err)
		}
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))
	}
}
