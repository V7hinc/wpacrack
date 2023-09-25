package allCrack

import (
	"github.com/urfave/cli/v2"
	"wpacrack/utils"
)

var Command = &cli.Command{
	Name:  "allCrack",
	Usage: "start the all password brute scan",
	Action: func(c *cli.Context) error {
		capFileFolder := c.String("capFileFolder")
		wechatBotKey := c.String("wechatBotKey")
		testWechatBot := c.Bool("testWechatBot")
		brute := newAllBrute(capFileFolder, wechatBotKey, testWechatBot)
		if brute.testWechatBot {
			err := utils.SendWeChatWorkMessage(brute.wechatBotKey, "test")
			if err != nil {
				return err
			}
		}
		return brute.BruteCrack()
	},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "capFileFolder",
			Aliases:  []string{"c"},
			Required: true,
			Usage:    "cap file Folder path",
		},
		&cli.StringFlag{
			Name:     "wechatBotKey",
			Aliases:  []string{"k"},
			Required: true,
			Value:    "",
			Usage:    "wechat bot key",
		},
		&cli.BoolFlag{
			Name:     "testWechatBot",
			Aliases:  []string{"b"},
			Required: false,
			Usage:    "test wechat bot key whether valid， will send a test message",
		},
	},
}
