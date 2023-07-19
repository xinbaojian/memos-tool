package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"memos-tool/utils"
)

// DeleteCommand 根据ID删除Memos
func DeleteCommand() *cli.Command {
	return &cli.Command{
		Name:    "delete",
		Aliases: []string{"del"},
		Usage:   "根据ID删除Memos",
		Action: func(c *cli.Context) error {
			var id = c.Args().First()
			if id == "" {
				fmt.Println("请输入要删除的Memos ID")
				return nil
			}
			return utils.DelById(id)
		},
	}
}
