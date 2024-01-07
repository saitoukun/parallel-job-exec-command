package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/cheggaaa/pb"
	"log"
	"os"
	"sync"
	"time"
)

type jobInfo struct {
	url    string
	device string
	secret string
}

// lineごとに必要な処理を行う
func processLine(ji jobInfo, wg *sync.WaitGroup, semaphore chan struct{}, bar *pb.ProgressBar) {
	defer func() {
		<-semaphore
		wg.Done()
		bar.Increment()
	}()
	doSomething(ji)
}

// 並列で実行する処理。例として数秒待機する。
func doSomething(ji jobInfo) {
	time.Sleep(3 * time.Second)
	fmt.Fprintf(os.Stderr, "num: %s\n", ji.url)
}

func main() {
	var (
		filePath = flag.String("f", "test.txt", "File path to list of URLs separated by line")
		parallel = flag.Int("P", 10, "Maximum number of goroutines to run  parallel")
	)
	flag.Parse()

	file, err := os.Open(*filePath)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer file.Close()

	// 進捗率のためにファイルの行数を取得
	scanner := bufio.NewScanner(file)
	var totalLines int
	for scanner.Scan() {
		totalLines++
	}
	bar := pb.StartNew(totalLines)
	file.Seek(0, 0)

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, *parallel)

	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		// 新しいgoroutineを起動するたびにWaitGroupに+1
		semaphore <- struct{}{}
		wg.Add(1)

		ji := jobInfo{ url: scanner.Text()}
		go processLine(ji, &wg, semaphore, bar)
	}

	// 終了
	wg.Wait()
	close(semaphore)
	bar.Finish()
}
