package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"memos-tool/config"
)

// GetConfig 获取配置命令
func GetConfig() *cli.Command {
	return &cli.Command{
		Name:    "config",
		Aliases: []string{"c"},
		Usage:   "config memos openId",
		Subcommands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a memos openId",
				Action: func(c *cli.Context) error {
					// 配置 memos openId
					openId := c.Args().First()
					if openId == "" {
						return fmt.Errorf("memos openId can not empty")
					}
					openId, err := config.SetOpenId(openId)
					if err == nil {
						fmt.Println("set openId success:", openId)
					}
					return err
				},
			},
			{
				Name:    "reset",
				Aliases: []string{"reset"},
				Usage:   "reset a memos openId",
				Action: func(c *cli.Context) error {
					// reset memos openId
					msg, err := config.Reset()
					if err == nil {
						fmt.Println(msg)
					}
					return err
				},
			},
		},
		Action: func(c *cli.Context) error {
			return fmt.Errorf("please use subcommand")
		},
	}
}