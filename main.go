package main // import "github.com/airylinus/ask4me"

import (
	"fmt"

	"github.com/atotto/clipboard"
)

func main() {
	// 要复制的字符串
	str := "Hello, world!"

	// 复制字符串到剪贴板中
	err := clipboard.WriteAll(str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("字符串已成功复制到剪贴板!")
	}
}
