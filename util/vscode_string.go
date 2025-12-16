package util

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// 获取单个文件的代码片段
func GetOneFileVscodeSnippet(filePath string, noNeedConvertSpecialChar []string) (*VscodeSnippet, error) {
	mdContent, err := ReadMDFile(filePath)
	if err != nil {
		fmt.Println("读取文件错误: ", os.Stderr, err)
		return nil, err
	}
	snippet, err := ParseSnippetFromMD(mdContent)
	if err != nil {
		return nil, err
	}
	// 特殊字符转换
	snippet.Code = EscapeSpecialChars(snippet.Code, noNeedConvertSpecialChar)

	codelines := strings.Split(snippet.Code, "\n")
	vscodeSnippet := VscodeSnippet{
		Prefix:      snippet.Prefix,
		Scope:       snippet.Scope,
		Description: snippet.Description,
		Body:        codelines,
	}
	return &vscodeSnippet, nil
}

// 特殊字符转换
func EscapeSpecialChars(code string, noNeedConvertSpecialChar []string) string {
	noNeedConvertSpecialCharSet := make(map[string]struct{})
	for _, item := range noNeedConvertSpecialChar {
		noNeedConvertSpecialCharSet[item] = struct{}{}
		noNeedConvertSpecialCharSet["{"+item+"}"] = struct{}{}
	}
	// 匹配 $VAR 或 ${VAR} 格式
	// 正则提取所有 $的变量，以空格为界限
	// 添加反斜杠
	pattern := `\$(\w+|\{\w+\})`
	re := regexp.MustCompile(pattern)
	// 打印所有结果
	matches := re.FindAllStringSubmatch(code, -1)
	varMap := make(map[string]string)
	for _, match := range matches {
		varVal := match[0]
		varName := match[1]
		fmt.Println("匹配到：", varVal, varName)
		if _, ok := noNeedConvertSpecialCharSet[varName]; ok {
			continue
		}
		varMap[varVal] = varName
	}
	// 匹配到的变量进行去重，保障只替换一次
	for varVal := range varMap {
		code = strings.ReplaceAll(code, varVal, `\`+varVal)
	}
	return code
}
