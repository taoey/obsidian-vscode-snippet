package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"taoey/obsidian-vscode-snippet/util"
)

// 获取单个文件的代码片段
func GetOneFileVscodeSnippet(filePath string, noNeedConvertSpecialChar []string) (*util.VscodeSnippet, error) {
	mdContent, err := util.ReadMDFile(filePath)
	if err != nil {
		fmt.Println("读取文件错误: ", os.Stderr, err)
		return nil, err
	}
	snippet, err := util.ParseSnippetFromMD(mdContent)
	if err != nil {
		return nil, err
	}
	// 特殊字符转换
	snippet.Code = util.EscapeSpecialChars(snippet.Code, noNeedConvertSpecialChar)

	codelines := strings.Split(snippet.Code, "\n")
	vscodeSnippet := util.VscodeSnippet{
		Prefix:      snippet.Prefix,
		Scope:       snippet.Scope,
		Description: snippet.Description,
		Body:        codelines,
	}
	return &vscodeSnippet, nil
}

func main() {
	config := util.GetConfig()
	fmt.Printf("load config: %+v\n", util.MustJsonString(config))
	// 1、加载文件列表
	dirpath := config.ObsidianDir

	mdFilepathList, err := util.GetDirSubMDFilepath(dirpath)
	if err != nil {
		fmt.Println("获取markdown文件失败:", err)
		return
	}
	// 2、处理每个文件，生成对应的vscode snippet
	snippetMap := make(map[string]util.VscodeSnippet)
	for _, filePath := range mdFilepathList {
		vscodeSnippet, err := GetOneFileVscodeSnippet(filePath, config.NoNeedConvertSpecialChar)
		if err != nil {
			fmt.Println("处理文件失败:", filePath, err)
			continue
		}
		// 只有有前缀的才有效
		if vscodeSnippet.Prefix == "" {
			continue
		}
		snippetMap[vscodeSnippet.Prefix] = *vscodeSnippet
	}
	resultJsonByte, _ := json.MarshalIndent(snippetMap, "", "  ")

	// 3、存储json到指定文件中，需要格式化输出
	outFilePath := config.OutputFilepath
	for _, outFilepath := range outFilePath {
		err = os.WriteFile(outFilepath, resultJsonByte, 0644)
		if err != nil {
			fmt.Println("写入文件失败:", outFilepath, err)
		}
		fmt.Println("写入文件成功:", outFilepath)
	}
}
