package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const XinsanbanListUrl = "https://xinsanban.eastmoney.com/api/gg/list"

const XinsanbanDtlUrl = "https://xinsanban.eastmoney.com/Article/NoticeContent?id="

/*
目前的最后一页是2500
*/
var maxPageSize = 2500

/*
开始执行页数
*/
var currentPageSize = 1

/*
执行间隔
*/
var intervalSecond = 8

const NeedMatchPart = "公开发行股票并在北交所上市辅导"
const segmentFlag = "|"

type Xinsanban_Result struct {
	Re     bool                   `json:"re"`
	Result []Xinsanban_Result_Dtl `json:"result"`
}

type Xinsanban_Result_Dtl struct {
	ArtCode    string `json:"art_code"`
	Title      string `json:"title"`
	NoticeDate string `json:"notice_date"`
}

// BatchGetXinsanbanListByPage 批跑递归任务/**
func BatchGetXinsanbanListByPage(isProxy bool) {
	time.Sleep(time.Duration(intervalSecond) * time.Second)
	getXinsanbanListByPage(currentPageSize, isProxy)
	currentPageSize++
	if currentPageSize <= maxPageSize {
		BatchGetXinsanbanListByPage(isProxy)
	}
}

func getXinsanbanListByPage(pageSize int, isProxy bool) {
	var client *http.Client
	if isProxy {
		client = getProxyHttpClient()
	} else {
		client = getHttpClient()
	}
	url := XinsanbanListUrl + "?page_index=" + strconv.Itoa(pageSize) + "&type=0&begin=&end=&securitycodes=&content=&sortRule=1"
	//url := "https://www.baidu.com"
	println("请求地址:" + url)
	req, err := http.NewRequest("GET", url, nil)
	check(err)
	getReqHeader(req)
	printAllHeader(req)
	//q := req.URL.Query()
	//q.Add("page_index", strconv.Itoa(int(pageSize)))
	//req.URL.RawQuery = q.Encode()
	//response, err := http.DefaultClient.Do(req)
	response, err := client.Do(req)
	check(err)
	if response.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", response.StatusCode, response.Status)
		return
	}
	body, _ := io.ReadAll(response.Body)
	defer response.Body.Close()
	println(string(body))
	if string(body) == "" {
		currentPageSize--
		return
	}
	responseJSON := &Xinsanban_Result{}
	err = json.Unmarshal(body, responseJSON)
	check(err)
	resList := responseJSON.Result
	needWriteData := ""
	for _, res := range resList {
		title := res.Title
		if title == "" {
			continue
		}
		if strings.Contains(title, NeedMatchPart) {
			needWriteData += title + segmentFlag + res.NoticeDate + segmentFlag + res.ArtCode +
				segmentFlag + XinsanbanDtlUrl + res.ArtCode + "\n"
		}
	}
	if needWriteData != "" {
		writeDataToTxtFile(needWriteData)
	}
}

func getProxyHttpClient() *http.Client {
	//使用代理服务器时设置超时
	var client = &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			// 设置代理
			Proxy: http.ProxyURL(getRandomProxyIp()),
		},
	}
	return client
}

func getHttpClient() *http.Client {
	return http.DefaultClient
}

func getReqHeader(req *http.Request) {
	//req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) "+
	//	"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36")
	//req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	//req.Header.Add("Pragma", "no-cache")
	//req.Header.Set("Connection", "keep-alive")
}

func printAllHeader(req *http.Request) {
	for name, values := range req.Header {
		// Loop over all values for the name.
		for _, value := range values {
			fmt.Println(name, value)
		}
	}
}
