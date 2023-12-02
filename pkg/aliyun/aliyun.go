package aliyun

import (
	"fmt"
	alimt "github.com/alibabacloud-go/alimt-20181012/v2/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	basic "ts/config"
	"ts/pkg"
)

// https://api.aliyun.com/api/alimt/2018-10-12/TranslateGeneral?tab=DEMO&lang=GO
const packageName = "aliyun"

var (
	accessKeyId     string
	accessKeySecret string
)

// LTAI5tQ2kHepgv2QpxAWo5RG
// Plqof8Sm3i16KQu8uin6vDmRGoNUgN

func Init() {

	section, err := basic.Config.GetSection(packageName)
	if err != nil {
		panic("阿里云读取配置失败")
	}

	accessKeyId = section.Key("accessKey").String()
	accessKeySecret = section.Key("Secret").String()

}

func createClient() *alimt.Client {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: &accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: &accessKeySecret,
	}
	// Endpoint 请参考 https://api.aliyun.com/product/alimt
	config.Endpoint = tea.String("mt.cn-hangzhou.aliyuncs.com")

	client, err := alimt.NewClient(config)
	if err != nil {
		panic("创建客户端失败")
	}
	return client
}

func Translate(command *pkg.Command) {
	req := alimt.TranslateRequest{
		SourceText:     command.Source,
		SourceLanguage: command.SourceLanguage,
		FormatType:     command.FormatType,
		Scene:          command.Scene,
		TargetLanguage: command.TargetLanguage,
	}

	client := createClient()
	resp, err := client.Translate(&req)

	if err != nil {
		fmt.Println(err.Error())
		panic("请求阿里云翻译失败")
	}

	if *resp.Body.Code != 200 {
		fmt.Printf("请求出错, Code: %v\n", *resp.Body.Code)
		return
	}

	fmt.Printf("翻译结果: %v \n", *resp.Body.Data.Translated)

}
