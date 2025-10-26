package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
	"gopkg.in/yaml.v3"
)

type VscodeSnippet struct {
	Prefix      string   `json:"prefix"`
	Scope       string   `json:"scope"`
	Description string   `json:"description"`
	Body        []string `json:"body"`
}

// Snippet 表示解析后的代码片段结构
type Snippet struct {
	Prefix       string `yaml:"prefix"`
	Description  string `yaml:"description"`
	Code         string
	CodeLanguage string
}

// 生成代码片段的json
func genCodeJson() string {
	return ""
}

// extractFirstCodeBlock 提取第一个代码块
func extractFirstCodeBlock(markdownContent string) (string, string) {
	// 创建解析器
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	p := parser.NewWithExtensions(extensions)

	// 解析Markdown
	doc := p.Parse([]byte(markdownContent))

	var language string
	var codeContent string

	// 遍历AST查找第一个代码块
	ast.WalkFunc(doc, func(node ast.Node, entering bool) ast.WalkStatus {
		if !entering {
			return ast.GoToNext
		}

		// 检查是否是代码块
		if codeBlock, ok := node.(*ast.CodeBlock); ok {
			if codeContent == "" { // 只取第一个代码块
				language = string(codeBlock.Info)
				codeContent = string(codeBlock.Literal)
				return ast.Terminate // 找到第一个后停止遍历
			}
		}
		return ast.GoToNext
	})

	return language, codeContent
}

func parseSnippetFromMD(mdText string) (*Snippet, error) {
	// 正则匹配 YAML front matter: 以 --- 开始和结束的部分
	re := regexp.MustCompile(`(?s)^---\s*\n(.*?)\n---\s*`)
	matches := re.FindStringSubmatch(mdText)

	var metadata []byte
	remaining := mdText

	if len(matches) >= 2 {
		metadata = []byte(matches[1])
		remaining = mdText[len(matches[0]):]
	}
	snippet := &Snippet{}
	// 解析 YAML
	if len(metadata) > 0 {
		err := yaml.Unmarshal(metadata, snippet)
		if err != nil {
			return nil, fmt.Errorf("YAML 解析失败: %w", err)
		}
	}
	language, code := extractFirstCodeBlock(remaining)
	// fmt.Println("匹配到的文档", string(metadata), snippet.Prefix)
	// fmt.Println("匹配结果", snippet.Description)
	// fmt.Println("剩余文档", remaining)
	snippet.CodeLanguage = language
	snippet.Code = code

	return snippet, nil
}

func readMDFile(path string) (string, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", fmt.Errorf("文件不存在: %s", path)
	}
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "用法: %s <Markdown文件路径>\n", os.Args[0])
		os.Exit(1)
	}

	filePath := os.Args[1]
	mdContent, err := readMDFile(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "读取文件错误: %v\n", err)
		os.Exit(1)
	}

	snippet, err := parseSnippetFromMD(mdContent)
	if err != nil {
		fmt.Fprintf(os.Stderr, "解析错误: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ 解析成功:")
	fmt.Printf("prefix:      %q\n", snippet.Prefix)
	fmt.Printf("description: %q\n", snippet.Description)
	fmt.Printf("code:        %q\n", snippet.Code)

	codelines := strings.Split(snippet.Code, "\n")

	vscodeSnippet := VscodeSnippet{
		Prefix:      snippet.Prefix,
		Scope:       snippet.CodeLanguage,
		Description: snippet.Description,
		Body:        codelines,
	}

	fmt.Println(vscodeSnippet)

}
