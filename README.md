# これは何
HTTP1.1のパイプライニングを使って、テキストファイルの中身を1行ずつHTTPで送りつけると、並列処理で分かち書きして返してくれるサーバーを試しに実装してみました。serverディレクトリ内のコードを実行するとlocalhost:8888で立ち上がります。

これを利用するためのクライアントコードもclientディレクトリにあります。

「[Goならわかるシステムプログラミング](https://www.lambdanote.com/products/go)」の第6章「TCPソケットとHTTPの実装」で紹介されている実装を参考にしました。

分かち書きには、形態素解析器として[kagome](https://github.com/ikawaha/kagome)を利用しています。

# 使い方

## セットアップ
```bash
git clone https://github.com/kecbigmt/go-kagome-pipeline-demo.git
cd go-kagome-pipeline-demo
```

## サーバー側
```bash
cd ./server
go run main.go
```

## クライアント側
```bash
cd ./client
go run main.go sample.txt #分かち書きしたいテキストファイルのパスを渡す
```

# サンプル
sample.txtには青空文庫の「としよりのお祖父さんと孫」が入っています。

https://www.aozora.gr.jp/cards/001091/card59849.html

これを分かち書きすると、以下のようになります。

```bash
$ cd ./client
$ go run main.go sample.txt
Access
HTTP/1.1 200 OK
Content-Length: 375

としよりのお祖父さんと孫

としより        動詞,自立,*,*,五段・ラ行,連用形,としよる,トシヨリ,トシヨリ
の      助詞,連体化,*,*,*,*,の,ノ,ノ
お祖父さん      名詞,一般,*,*,*,*,お祖父さん,オジイサン,オジーサン
と      助詞,並立助詞,*,*,*,*,と,ト,ト
孫      名詞,一般,*,*,*,*,孫,マゴ,マゴ

        記号,空白,*,*,*,*,*

HTTP/1.1 200 OK
Content-Length: 209

ヤーコップ、ウィルヘルム・グリム

ヤーコップ      名詞,一般,*,*,*,*,*
、      記号,読点,*,*,*,*,、,、,、
ウィルヘルム・グリム    名詞,一般,*,*,*,*,*

        記号,空白,*,*,*,*,*

HTTP/1.1 200 OK
Content-Length: 230

金田鬼一訳

金田    名詞,固有名詞,人名,姓,*,*,金田,カネダ,カネダ
鬼一    名詞,固有名詞,人名,名,*,*,鬼一,キイチ,キイチ
訳      名詞,接尾,一般,*,*,*,訳,ヤク,ヤク

        記号,空白,*,*,*,*,*

HTTP/1.1 200 OK
Content-Length: 3005

むかし昔、あるところに石みたようにとしをとったおじいさんがありました。おじいさんは、目はかすんでしまい、耳はつんぼになって、膝は、ぶるぶるふるえていました。

むかし  名詞,副詞可能,*,*,*,*,むかし,ムカシ,ムカシ
昔      名詞,副詞可能,*,*,*,*,昔,ムカシ,ムカシ
、      記号,読点,*,*,*,*,、,、,、
ある    連体詞,*,*,*,*,*,ある,アル,アル
ところ  名詞,非自立,副詞可能,*,*,*,ところ,トコロ,トコロ
に      助詞,格助詞,一般,*,*,*,に,ニ,ニ
石      名詞,一般,*,*,*,*,石,イシ,イシ
み      動詞,自立,*,*,一段,連用形,みる,ミ,ミ
た      助動詞,*,*,*,特殊・タ,基本形,た,タ,タ
よう    名詞,非自立,助動詞語幹,*,*,*,よう,ヨウ,ヨー
に      助詞,副詞化,*,*,*,*,に,ニ,ニ
と      助詞,格助詞,引用,*,*,*,と,ト,ト
し      動詞,自立,*,*,サ変・スル,連用形,する,シ,シ
を      助詞,格助詞,一般,*,*,*,を,ヲ,ヲ
とっ    動詞,自立,*,*,五段・ラ行,連用タ接続,とる,トッ,トッ
た      助動詞,*,*,*,特殊・タ,基本形,た,タ,タ
おじいさん      名詞,一般,*,*,*,*,おじいさん,オジイサン,オジーサン
が      助詞,格助詞,一般,*,*,*,が,ガ,ガ
あり    動詞,自立,*,*,五段・ラ行,連用形,ある,アリ,アリ
まし    助動詞,*,*,*,特殊・マス,連用形,ます,マシ,マシ
た      助動詞,*,*,*,特殊・タ,基本形,た,タ,タ
。      記号,句点,*,*,*,*,。,。,。
おじいさん      名詞,一般,*,*,*,*,おじいさん,オジイサン,オジーサン
は      助詞,係助詞,*,*,*,*,は,ハ,ワ
、      記号,読点,*,*,*,*,、,、,、
目      名詞,一般,*,*,*,*,目,メ,メ
は      助詞,係助詞,*,*,*,*,は,ハ,ワ
かすん  動詞,自立,*,*,五段・マ行,連用タ接続,かすむ,カスン,カスン
で      助詞,接続助詞,*,*,*,*,で,デ,デ
しまい  動詞,非自立,*,*,五段・ワ行促音便,連用形,しまう,シマイ,シマイ
、      記号,読点,*,*,*,*,、,、,、
耳      名詞,一般,*,*,*,*,耳,ミミ,ミミ
は      助詞,係助詞,*,*,*,*,は,ハ,ワ
つん    動詞,自立,*,*,五段・マ行,連用タ接続,つむ,ツン,ツン
ぼ      動詞,自立,*,*,五段・ラ行,体言接続特殊２,ぼる,ボ,ボ
に      助詞,格助詞,一般,*,*,*,に,ニ,ニ
なっ    動詞,自立,*,*,五段・ラ行,連用タ接続,なる,ナッ,ナッ
て      助詞,接続助詞,*,*,*,*,て,テ,テ
、      記号,読点,*,*,*,*,、,、,、
膝      名詞,一般,*,*,*,*,膝,ヒザ,ヒザ
は      助詞,係助詞,*,*,*,*,は,ハ,ワ
、      記号,読点,*,*,*,*,、,、,、
ぶるぶる        副詞,一般,*,*,*,*,ぶるぶる,ブルブル,ブルブル
ふるえ  動詞,自立,*,*,一段,連用形,ふるえる,フルエ,フルエ
て      助詞,接続助詞,*,*,*,*,て,テ,テ
い      動詞,非自立,*,*,一段,連用形,いる,イ,イ
まし    助動詞,*,*,*,特殊・マス,連用形,ます,マシ,マシ
た      助動詞,*,*,*,特殊・タ,基本形,た,タ,タ
。      記号,句点,*,*,*,*,。,。,。

        記号,空白,*,*,*,*,*

... 以下略
```