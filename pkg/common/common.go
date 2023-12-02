package common

// Origin 原始数据
type Origin struct {
	SourceText     string // 待翻译源文本
	SourceLanguage string // 源文本语言
	TargetLanguage string // 目标语言
	SourceFormat   string // 源文本格式
}
