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

	once.Do(func() {

	})

	var err error
	Config, err = ini.Load("/Users/mixinju/code/go/ts/config/config.ini")
	if err != nil {
		fmt.Println(err.Error())
		panic("读取配置文件出错")
	}
}
