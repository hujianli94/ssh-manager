package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"ssh-manager/cmd/add"
	"ssh-manager/cmd/list"
	"ssh-manager/cmd/remove"
)

var (
	// Version 编译时注入
	Version = "v2.0.0"
)

func main() {
	app := &cli.App{
		Name:    "ssh-manager",
		Usage:   "A simple SSH alias manager",
		Version: Version, // 显示版本号
		Commands: []*cli.Command{
			add.Command,    // 注册 add 子命令
			list.Command,   // 注册 list 子命令
			remove.Command, // 注册 remove 子命令
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
