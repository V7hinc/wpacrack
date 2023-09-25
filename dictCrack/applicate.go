package dictCrack

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

type dictBrute struct {
	capFileFolder        string
	passwdDictFileFolder string
	wechatBotKey         string
	testWechatBot        bool
}

func newDictBrute(capFileFolder, passwdDictFileFolder, wechatBotKey string, testWechatBot bool) *dictBrute {
	return &dictBrute{
		capFileFolder:        capFileFolder,
		passwdDictFileFolder: passwdDictFileFolder,
		wechatBotKey:         wechatBotKey,
		testWechatBot:        testWechatBot,
	}
}

func (b *dictBrute) BruteCrack() error {
	startTime := time.Now()
	var (
		impMsg       string
		foundKey     bool
		err          error
		passwordFile = "waitTestPasswd.txt"
	)
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
		queue := crack.GenPasswdDict(b.passwdDictFileFolder)
		logger.Info(capFile)
		for !foundKey && err == nil {
			isFinish := crack.WritePasswordFile(passwordFile, queue)
			impMsg, foundKey, err = crack.Crack(passwordFile, capFile)
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
