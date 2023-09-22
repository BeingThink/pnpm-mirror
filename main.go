package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"os"
	"github.com/beingthink/pnpm-mirror/src/github"
	"github.com/beingthink/pnpm-mirror/src/server"
)

// 配置项
var GITHUB_TOKEN string
var OWNER string
var REPO string

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Print(err)
	}
	GITHUB_TOKEN = os.Getenv("GITHUB_TOKEN")
	OWNER = os.Getenv("OWNER")
	REPO = os.Getenv("REPO")
}

func main() {
	// 加载配置文件
	loadConfig()
	server.StartServer()

	// 获取github release
	// pnpmRelease := getPublicRepoReleases(OWNER, REPO)

	// for _, release := range pnpmRelease {
	// 	for _, asset := range release.Assets {
	// 		println(release.TagName, asset.BrowserDownloadUrl)
	// 		go github.DownloadFile(asset.BrowserDownloadUrl, release.TagName)
	// 	}
	// }
}

func getPublicRepoReleases(owner string, repo string) []github.Release {
	url := "https://api.github.com/repos/" + owner + "/" + repo + "/releases"
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err)
	}

	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", "Bearer "+GITHUB_TOKEN)
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	res, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
	}

	resultBuf, _ := io.ReadAll(res.Body)

	// 关闭连接
	defer res.Body.Close()

	var resultArr []github.Release

	err = json.Unmarshal(resultBuf, &resultArr)
	if err != nil {
		fmt.Print(err)
	}

	return resultArr
	// fmt.Print(resultStr)
}
