package pkg

import (
	"flag"
	"fmt"
	"os"
	"regexp"
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

	var sourceLanguage string
	var targetLanguage string
	if command.isChinese() {
		sourceLanguage = Han
		targetLanguage = En
	}

	if command.isEnglish() {
		sourceLanguage = En
		targetLanguage = Han
	}

	command.SourceLanguage = set.String("sl", sourceLanguage, "源文本语言")
	command.Scene = set.String("scene", "general", "场景")
	command.FormatType = set.String("format", "text", "格式化文本")
	command.TargetLanguage = set.String("tl", targetLanguage, "目标语言")
	err := set.Parse(os.Args[2:])
	if err != nil {
		panic("解析参数出错")
	}

	return command
}

// Parse 解析命令行参数
func (c *Command) Parse() {

}

func (c *Command) isChinese() bool {

	for _, r := range *c.Source {
		if unicode.Is(unicode.Han, r) {
			return true
		}
	}

	return false
}

func (c *Command) isEnglish() bool {
	ok, err := regexp.MatchString("^([A-z]+)$", *c.Source)
	if err != nil {
		fmt.Println(err)
	}
	return ok
}
