package github

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var mirrorPath string

func init() {
	workDir, _ := os.Getwd()
	mirrorPath = filepath.Join(workDir, "mirror")
}

// https://github.com/pnpm/pnpm/releases/download/v8.7.6/pnpm-linux-arm64
// DownloadFile 下载文件并保存到指定路径
func DownloadFile(url string, version string) error {
	// 发起 HTTP GET 请求
	client := &http.Client{	
		Timeout: time.Second * 5,
	}

	// 发起 GET 请求
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 创建目标文件
	pnpmArcName := filepath.Base(url)
	pnpmVerPath := filepath.Join(mirrorPath, version)
	pnpmItemPath := filepath.Join(pnpmVerPath, pnpmArcName)

	if !isExsit(pnpmVerPath) {
		os.Mkdir(pnpmVerPath, os.ModePerm)
		println("成功创建文件路径", pnpmVerPath)
	}

	if isExsit(pnpmItemPath) {
		return nil
	}

	out, err := os.Create(pnpmItemPath)
	if err != nil {
		println("文件创建失败")
		return err
	}
	defer out.Close()

	// 将 HTTP 响应的内容拷贝到目标文件
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	resp.Close = true
	return nil
}

func isExsit(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
