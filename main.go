package main

import (
	"encoding/json"
	"fmt"
	"os"
	"taoey/obsidian-vscode-snippet/util"
)

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
		vscodeSnippet, err := util.GetOneFileVscodeSnippet(filePath, config.NoNeedConvertSpecialChar)
		if err != nil {
			fmt.Println("处理文件失败:", filePath, err)
			continue
		}
		// 只有有前缀的才有效
		if vscodeSnippet.Prefix == "" {
			continue
		}
		desc := vscodeSnippet.Description
		// 如果描述为空，则使用前缀作为描述
		if desc == "" {
			desc = vscodeSnippet.Prefix
		}
		vscodeSnippet.Description = ""
		snippetMap[desc] = *vscodeSnippet
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
