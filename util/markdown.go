package util

import (
	"fmt"
	"regexp"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
	"gopkg.in/yaml.v3"
)

// Snippet markdown格式化
type Snippet struct {
	// 文档表头属性，自动加载
	Prefix      string `yaml:"prefix"`
	Description string `yaml:"description"`
	Scope       string `yaml:"Scope"`
	// 代码块内容
	Code         string
	CodeLanguage string
}

type VscodeSnippet struct {
	Prefix      string   `json:"prefix"`
	Scope       string   `json:"scope"`
	Description string   `json:"description"`
	Body        []string `json:"body"`
}

// ExtractFirstCodeBlock 提取第一个代码块
func ExtractFirstCodeBlock(markdownContent string) (string, string) {
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

func ParseSnippetFromMD(mdText string) (*Snippet, error) {
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
	language, code := ExtractFirstCodeBlock(remaining)
	// fmt.Println("匹配到的文档", string(metadata), snippet.Prefix)
	// fmt.Println("匹配结果", snippet.Description)
	// fmt.Println("剩余文档", remaining)
	snippet.CodeLanguage = language
	snippet.Code = code

	return snippet, nil
}
