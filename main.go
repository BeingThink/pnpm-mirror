package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

// 配置项
var GITHUB_TOKEN string
var OWNER string
var REPO string

func loadConfig() {
	err := godotenv.Load()
	if (err != nil) {
		fmt.Print(err)
	}
	GITHUB_TOKEN = os.Getenv("GITHUB_TOKEN")
	OWNER = os.Getenv("OWNER")
	REPO = os.Getenv("REPO")
}

func main() {
	// 加载配置文件
	loadConfig()

	// 获取github release
	getPublicRepoReleases(OWNER, REPO)
}

func getPublicRepoReleases(owner string, repo string) {
	url := "https://api.github.com/repos/" + owner + "/" + repo +"/releases"
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if (err != nil) {
		fmt.Print(err)
	}

	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", "Bearer " + GITHUB_TOKEN)
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	res, err := client.Do(req)
	if (err != nil) {
		fmt.Print(err)
	}

	resultBuf, _ := io.ReadAll(res.Body)

	type responseBody struct {
		Url string `json:"url"` // 需要指定结构体描述，给json解析提供接口体和json key的映射
		HtmlUrl string `json:"html_url"`
	}

	var resultArr []responseBody

	err = json.Unmarshal(resultBuf, &resultArr)
	if (err != nil) {
		fmt.Print(err)
	}

	fmt.Print(resultArr)
	// fmt.Print(resultStr)
}
