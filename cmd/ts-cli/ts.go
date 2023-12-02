package main

import (
	"ts/pkg"
	"ts/pkg/aliyun"
)

func main() {

	command := pkg.Parse()
	aliyun.Init()
	aliyun.Translate(command)
}
