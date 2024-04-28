package main // import "github.com/airylinus/ask4me"

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/airylinus/ask4me/entities"
	"github.com/atotto/clipboard"
	"gopkg.in/yaml.v3"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ask4me [prompt_id]")
		os.Exit(1)
	}
	promptID := os.Args[1]
	prompts := readPrompts()
	for _, prompt := range prompts {
		if prompt.ID == promptID {
			question := generatePromptsAndCode(prompt)
			done := copyToClipboard(question)
			if done {
				fmt.Println("ask4me: Let's GO")
				os.Exit(0)
			}
		}
	}
	fmt.Println("prompt not found")
	os.Exit(1)
	// fmt.Println("prompts", prompts)
}

func generatePromptsAndCode(prompt entities.Prompt) (question string) {
	question += prompt.Description + "\n\n"
	// 获取当前用户的 home 目录
	for _, fp := range prompt.Files {
		question += fp + " 文件内容如下: \n\n"
		sourceCode := readFromFile(filepath.Join(prompt.BasePath, fp))
		question += "```\n" + sourceCode + "\n```\n\n"
	}
	return question
}

func readFromFile(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return ""
	}
	return string(data)
}

func copyToClipboard(str string) (Done bool) {
	err := clipboard.WriteAll(str)
	if err != nil {
		return false
	} else {
		return true
	}
}

func readPrompts() (prompts []entities.Prompt) {
	// 获取当前用户的 home 目录
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}

	// 构建 prompts.yaml 文件的路径
	promptsFile := filepath.Join(homeDir, ".ask4me", "prompts.yaml")

	// 读取 prompts.yaml 文件内容
	data, err := os.ReadFile(promptsFile)
	if err != nil {
		fmt.Println("Error reading prompts.yaml file:", err)
		return
	}

	// 解析 YAML 文件内容
	err = yaml.Unmarshal(data, &prompts)
	if err != nil {
		fmt.Println("Error parsing YAML file:", err)
		return
	}

	// 输出 prompts 内容
	// fmt.Println("Prompts:", prompts)
	return prompts
}
