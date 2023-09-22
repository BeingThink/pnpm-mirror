package github

type Asset struct {
	Url string `json:"url"` // 需要指定结构体描述，给json解析提供接口体和json key的映射
	BrowserDownloadUrl string `json:"browser_download_url"`
}

type Release struct {
	TagName string `json:"tag_name"`
	Name string `json:"name"`
	Id int `json:"id"`
	Assets []Asset `json:"assets"`
}