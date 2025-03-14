package cmd

import (
	"wxcloudrun-golang/cmd/err"
	"wxcloudrun-golang/cmd/help"
	"wxcloudrun-golang/cmd/wolf"
)

func Init() {
	help.Init()
	wolf.Init()
	err.Init()
}
