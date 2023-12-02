package pkg

import (
	"flag"
	"fmt"
	"os"
	"unicode"
)

type Command struct {
	Source         *string
	SourceLanguage *string
	TargetLanguage *string
	Scene          *string
	FormatType     *string
}

func Parse() *Command {

	command := &Command{}
	set := flag.NewFlagSet("ts", flag.ExitOnError)

	command.Source = &os.Args[1]
	command.Scene = set.String("scene", "general", "场景")
	command.SourceLanguage = set.String("sl", "zh", "原文本语言,默认中文")
	command.FormatType = set.String("format", "text", "格式化文本")
	err := set.Parse(os.Args[2:])

	if err != nil {
		fmt.Println("解析参数出错")
	}

	return command
}

// Parse 解析命令行参数
func (c *Command) Parse() {

}

func (c *Command) isChinese(s string) bool {

	for _, r := range s {
		if unicode.Is(unicode.Han, r) {
			return true
		}
	}

	return false
}
