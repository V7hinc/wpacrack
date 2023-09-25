package allCrack

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
	"wpacrack/crack"
	"wpacrack/logging"
	"wpacrack/utils"
)

var logger = logging.GetSugar()

type allBrute struct {
	capFileFolder string
	wechatBotKey  string
	testWechatBot bool
}

func newAllBrute(capFileFolder, wechatBotKey string, testWechatBot bool) *allBrute {
	return &allBrute{
		capFileFolder: capFileFolder,
		wechatBotKey:  wechatBotKey,
		testWechatBot: testWechatBot,
	}
}

func (b *allBrute) BruteCrack() error {
	startTime := time.Now()
	var (
		impMsg   string
		foundKey bool
		err      error
	)
	passwordFile, err := os.CreateTemp("/tmp", "waitTestPasswd")
	passwordFile.Close()

	defer os.RemoveAll(passwordFile.Name())

	var files = make([]string, 0)
	err = filepath.Walk(b.capFileFolder, func(path string, info os.FileInfo, err error) error {
		stat, err := os.Stat(path)
		if err != nil {
			return err
		}
		if !stat.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return err
	}
	for _, capFile := range files {
		queue := crack.GenerateAllPasswords()
		logger.Info(capFile)
		for !foundKey && err == nil {
			isFinish := crack.WritePasswordFile(passwordFile.Name(), queue)
			impMsg, foundKey, err = crack.Crack(passwordFile.Name(), capFile)
			if isFinish {
				break
			}
		}

		var msg string
		if err != nil {
			msg = impMsg + "\n" + err.Error()
		} else {
			msg = fmt.Sprintf("%s\nfoundKey:%t", impMsg, foundKey)
		}
		msg += "\n" + time.Now().Sub(startTime).String()
		err = utils.SendWeChatWorkMessage(b.wechatBotKey, msg)
		if err != nil {
			logger.Error(err)
		}
	}
	return err
}
