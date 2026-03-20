package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	// 获取当前时间，格式化为 YYYYMMDDHHMM
	// 注意：Go 的格式化时间基准点必须是 "2006-01-02 15:04:05"
	timestamp := time.Now().Format("200601021504")

	// 遍历所有拖入的文件路径 (从 os.Args[1] 开始)
	for _, oldPath := range os.Args[1:] {
		// 获取文件所在的目录
		dir := filepath.Dir(oldPath)
		// 获取带有扩展名的完整文件名 (例如: a.txt)
		base := filepath.Base(oldPath)
		// 获取扩展名 (例如: .txt)
		ext := filepath.Ext(base)
		// 获取去掉扩展名的纯文件名 (例如: a)
		nameWithoutExt := strings.TrimSuffix(base, ext)

		// 拼接新文件名：原名(时间戳).扩展名
		newName := fmt.Sprintf("%s(%s)%s", nameWithoutExt, timestamp, ext)
		// 拼接完整的新文件路径
		newPath := filepath.Join(dir, newName)

		// 执行重命名操作
		err := os.Rename(oldPath, newPath)
		if err != nil {
			fmt.Printf("❌ 重命名失败: %s\n   原因: %v\n", base, err)
		} else {
			fmt.Printf("✅ 成功: %s -> %s\n", base, newName)
		}
	}
}
