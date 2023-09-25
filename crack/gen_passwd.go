package crack

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
	"wpacrack/logging"
)

var logger = logging.GetSugar()

func GenPasswdDict(dictPath string) <-chan string {
	var queue = make(chan string, 1000000)
	go func() {
		defer close(queue)
		err := filepath.Walk(dictPath, func(path string, info os.FileInfo, err error) error {
			stat, err := os.Stat(path)
			if err != nil {
				return err
			}
			if !stat.IsDir() {
				f, err := os.Open(path)
				if err != nil {
					logger.Error(err)
					return err
				}
				reader := bufio.NewReader(f)
				for {
					readString, err := reader.ReadString('\n')
					if err != nil && err != io.EOF {
						logger.Error(err)
						break
					}
					if err == io.EOF {
						break
					}
					readString = strings.TrimSpace(readString)
					if len(readString) >= 8 && len(readString) <= 12 {
						queue <- readString
					}
				}
			}
			return nil
		})
		if err != nil {
			logger.Error(err)
		}
	}()
	time.Sleep(10 * time.Second)
	return queue
}

func WritePasswordFile(passwordFile string, queue <-chan string) (isFinish bool) {
	queueLen := len(queue)
	logger.Info(queueLen)
	if queueLen == 0 {
		isFinish = true
		return
	}
	if _, err := os.Stat(passwordFile); err == nil {
		_ = os.RemoveAll(passwordFile)
	}

	file, err := os.OpenFile(passwordFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	write := bufio.NewWriter(file)

	for i := 0; i < queueLen; i++ {
		_, _ = write.WriteString(<-queue + "\n")
	}
	_ = write.Flush()
	return
}
