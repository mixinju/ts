package main

import (
	basic "ts/config"
	"ts/pkg"
	"ts/pkg/aliyun"
)

func main() {
	command := pkg.Parse()
	basic.Init()
	aliyun.Init()
	aliyun.Translate(command)
}
