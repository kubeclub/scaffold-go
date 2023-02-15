package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	// 爬取的目标网站页面
	url := "https://www.baidu.com"
	// 向该网站发送请求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("http get error", err)
		return
	}
	// 读取网页内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read error", err)
		return
	}
	// 将网页内容转化为字符串
	html := string(body)
	// 正则表达式，用于提取图片链接
	imgRe := regexp.MustCompile(`<img[^>]+src="(.*?)"`)
	// 提取图片链接
	imgs := imgRe.FindAllStringSubmatch(html, -1)
	for _, img := range imgs {
		// 去除不完整的链接
		if strings.HasPrefix(img[1], "http") {
			fmt.Println(img[1])
		}
	}
}
