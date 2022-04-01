package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容

func Get(url string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
//var client *http.Client

func Post(url string, data interface{}, contentType string) (string,error) {
	url = "http://127.0.0.1:5000/data" + url
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)

	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))

	if err != nil {
		if err.Error() != "" {
			log.Printf(err.Error())
			return "",err
		}
		return "", errors.New(resp.Status)
	}

	if resp.StatusCode != 200 {
		return resp.Status,errors.New(resp.Status)
	}

	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result),err
}


