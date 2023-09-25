package crack

import (
	"bytes"
	"time"
)

func GenerateAllPasswords() <-chan string {
	var queue = make(chan string, 1000000)
	go func() {
		defer close(queue)
		charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

		// 生成长度为i的所有密码
		for i := 8; i <= 12; i++ {
			generateLengthAllPasswords(charset, i, queue)
		}
	}()
	time.Sleep(10 * time.Second)
	return queue
}
func generateLengthAllPasswords(charset string, length int, queue chan<- string) {
	// 初始化第一个密码
	password := make([]byte, length)
	for i := 0; i < length; i++ {
		password[i] = charset[0]
	}
	charsetByte := []byte(charset)

	// 生成所有可能的密码组合
	for {
		queue <- string(password)

		// 递增密码
		for i := length - 1; i >= 0; i-- {
			if password[i] == charset[len(charset)-1] {
				password[i] = charset[0]
			} else {
				password[i] = charset[bytes.IndexByte(charsetByte, password[i])+1]
				break
			}
		}

		// 检查是否已经生成了所有可能的密码
		if password[0] == charset[len(charset)-1] {
			break
		}
	}
}
