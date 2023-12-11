package main

import (
	"math/rand"
	"net/url"
)

var arr = [1]string{
	"http://58.221.40.175:80",
}

func getRandomProxyIp() *url.URL {
	length := len(arr)
	r := rand.Intn(length)
	println("使用的代理ip是:" + arr[r])
	randomUrl, err := url.Parse(arr[r])
	check(err)
	return randomUrl
}

func getProxyIp(index int) *url.URL {
	randomUrl, err := url.Parse(arr[index])
	check(err)
	return randomUrl
}
