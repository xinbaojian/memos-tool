package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"memos-tool/commands"
	"memos-tool/enums"
	"memos-tool/utils"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "memos",
		Usage: "memos helper",
		Commands: []*cli.Command{
			commands.GetList(),
			commands.GetConfig(),
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "tag",
				Aliases:     []string{"t"},
				Value:       "memos",
				Usage:       "发送 memos 内容的标签",
				DefaultText: "memos",
			},
			&cli.IntFlag{
				Name:        "visibility",
				Aliases:     []string{"v"},
				Value:       0,
				Usage:       "发送 memos 内容可见性，0: PRIVATE, 1: PUBLIC, 2: PROTECTED",
				DefaultText: "0: PRIVATE",
			},
		},
		Action: func(c *cli.Context) error {
			tag := fmt.Sprintf("#%s", c.String("tag"))
			visibility := enums.AccessLevel(c.Int("visibility")).String()
			content := c.Args().First()
			if content == "" {
				return fmt.Errorf("memos content is empty")
			}
			memos := utils.Memos{Content: fmt.Sprintf("%s %s", tag, content), Visibility: visibility}
			utils.SendMemos(&memos)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
