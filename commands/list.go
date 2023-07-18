package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"memos-tool/enums"
)

// GetList 获取添加命令
func GetList() *cli.Command {
	return &cli.Command{
		Name:                   "list",
		Aliases:                []string{"a"},
		Usage:                  "拉取Memos列表",
		UseShortOptionHandling: true,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "limit",
				Aliases:     []string{"l"},
				Value:       20,
				Usage:       "查询Memos条数限制",
				DefaultText: "20条",
			},
			&cli.IntFlag{
				Name:        "status",
				Aliases:     []string{"s"},
				Value:       0,
				Usage:       "memos 内容状态，0: NORMAL, 1: ARCHIVED",
				DefaultText: "查询Memos内容状态，默认为NORMAL",
			},
		},
		Action: func(c *cli.Context) error {
			// TODO 查询Memos列表
			limit := c.Int("limit")
			rowStatus := enums.RowStatus(c.Int("status")).String()
			fmt.Println(limit, rowStatus)
			return nil
		},
	}
}
