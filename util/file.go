package util

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadMDFile(path string) (string, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", fmt.Errorf("文件不存在: %s", path)
	}
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// GetDirSubFilepath 返回指定目录下所有文件的完整路径（递归）
func GetDirSubFilepath(dirFilepath string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(dirFilepath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

// 筛选出文件中的所有markdown文档
func GetDirSubMDFilepath(dirFilepath string) ([]string, error) {
	result := []string{}
	filepath, err := GetDirSubFilepath(dirFilepath)
	if err != nil {
		return result, nil
	}
	for _, filepath := range filepath {
		if strings.HasSuffix(filepath, ".md") {
			result = append(result, filepath)
		}
	}
	return result, nil
}

const configFilepath = "config/config.json"

type Config struct {
	ObsidianDir              string   `json:"obsidian_dir"`
	OutputFilepath           []string `json:"output_filepath"`
	NoNeedConvertSpecialChar []string `json:"no_need_convert_special_char"`
}

func GetConfig() *Config {
	// 读取本地josn配置文件
	// 读取文件内容
	data, err := os.ReadFile(configFilepath)
	if err != nil {
		return nil
	}
	// 解析 JSON 到结构体
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil
	}

	return &cfg
}
