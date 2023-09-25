package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// 定义消息结构体
type WeChatWorkMessage struct {
	MsgType  string   `json:"msgtype"`
	Text     Text     `json:"text"`
	Markdown Markdown `json:"markdown"`
}

type Text struct {
	Content string `json:"content"`
}

type Markdown struct {
	Content string `json:"content"`
}

// 发送消息到企业微信机器人
func SendWeChatWorkMessage(accessKey, message string) error {
	webhookURL := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + accessKey

	// 构建消息结构体
	wechatWorkMessage := WeChatWorkMessage{
		MsgType: "text",
		Text: Text{
			Content: message,
		},
	}

	// 将消息转换为JSON
	jsonMessage, err := json.Marshal(wechatWorkMessage)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(jsonMessage))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // 忽略证书验证
			},
		}}
	// 发送HTTP POST请求
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request status not ok")
	}

	return nil
}
