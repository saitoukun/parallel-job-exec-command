# parallel-job-exec-client

Channel, WaitGroupを使うことで、常に一定の数のgoroutinesを実行し、効率的にジョブを実行させる。
https://github.com/cheggaaa/pb のプログレスバーを使って進捗率を表示する。

## コマンド引数

| 引数 | 説明                                             |
|----|------------------------------------------------|
| -f | File path to job list separated by line    |
| -P | Maximum number of goroutines to run  parallel  |

## How to dev

Dockerコンテナを用意しています。
go言語が使用できる環境ならローカルでも使用可能。

https://hub.docker.com/_/golang

### Dockerの環境で実行したい場合
```
# Docker imageをビルド
docker build -t parallel-job-exec-client .

# コンテナへ接続
docker run -it parallel-job-exec-client bash

# 実行
go run main.go -f test.txt -P 3 2> aaa
go run main.go -f test.txt -P 3 2>/dev/null
```

### goからバイナリを生成して実行したい場合
``` 
go build main.go
./main -f test.txt -P 10 2>/dev/null
```
