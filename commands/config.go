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
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "key",
						Aliases: []string{"k"},
						Value:   "openid",
					},
				},
				Action: func(c *cli.Context) error {
					return AddCommand(c)
				},
			},
			{
				Name:    "reset",
				Aliases: []string{"reset"},
				Usage:   "reset a memos openId",
				Action: func(c *cli.Context) error {
					return ResetCommand(c)
				},
			},
			{
				Name:    "show",
				Aliases: []string{"s"},
				Usage:   "show memos config info",
				Action: func(c *cli.Context) error {
					return ShowCommand(c)
				},
			},
		},
		Action: func(c *cli.Context) error {
			return fmt.Errorf("please use subcommand")
		},
	}
}

// AddCommand 配置 add 命令 逻辑处理
func AddCommand(c *cli.Context) error {
	// 配置 memos openId
	value := c.Args().First()
	if value == "" {
		return fmt.Errorf("memos openId can not empty")
	}
	key := c.String("key")
	value, err := config.Set(key, value)
	if err == nil {
		fmt.Println(fmt.Sprintf("set %s success: %s", key, value))
	}
	return err
}

// ResetCommand 配置 reset 命令 逻辑处理
func ResetCommand(c *cli.Context) error {
	// reset memos openId
	msg, err := config.Reset()
	if err == nil {
		fmt.Println(msg)
	}
	return err
}

// ShowCommand 配置 show 命令 逻辑处理
func ShowCommand(c *cli.Context) error {
	// show memos config info
	jsonByte, err := config.Show()
	if err != nil {
		return err
	}
	fmt.Println(string(jsonByte))
	return nil
}
