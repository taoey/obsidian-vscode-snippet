package util

import (
	"fmt"
	"regexp"
	"strings"
)

// 特殊字符转换
func EscapeSpecialChars(code string, noNeedConvertSpecialChar []string) string {
	noNeedConvertSpecialCharSet := make(map[string]struct{})
	for _, item := range noNeedConvertSpecialChar {
		noNeedConvertSpecialCharSet[item] = struct{}{}
		noNeedConvertSpecialCharSet["{"+item+"}"] = struct{}{}
	}

	// 匹配 $VAR 或 ${VAR} 格式
	// 正则提取所有 $的变量，以空格为界限
	pattern := `\$(\w+|\{\w+\})`
	re := regexp.MustCompile(pattern)
	// 打印所有结果
	matches := re.FindAllStringSubmatch(code, -1)
	for _, match := range matches {
		varVal := match[0]
		varName := match[1]
		fmt.Println("匹配到：", varVal, varName)
		if _, ok := noNeedConvertSpecialCharSet[varName]; ok {
			continue
		}
		code = strings.ReplaceAll(code, varVal, "\\"+varVal)
	}

	return code
}
