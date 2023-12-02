package basic

import (
	"fmt"
	"gopkg.in/ini.v1"
	"sync"
)

var (
	once   sync.Once
	Config *ini.File
)

func Init() {

	var err error
	Config, err = ini.Load("/Users/mixinju/.config/ts/config.ini")
	if err != nil {
		fmt.Println(err.Error())
		panic("读取配置文件出错")
	}
}
