package main

//import (
//	"bytes"
//	"fmt"
//	jsoniter "github.com/json-iterator/go"
//	"io/ioutil"
//	"net/http"
//)

type Comment struct{
	PostId int `json:postId`
	Id int `json:id`
	Name string `json:name`
	Email string `json:email`
	Body string `json:body`
}

/*

func main() {

	// 簡易型訪問

	// Get Website
	resp, err := http.Get("http://jsonplaceholder.typicode.com/comments/1")
	if err != nil {
		fmt.Println(err)
	}

	// 取得的結果轉成Byte
	tempByteArray,_ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(tempByteArray[:]))

	// 將json轉換成object
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var comment Comment
	json.Unmarshal(tempByteArray,&comment)
	fmt.Println(comment)

	// 將object轉換成json
	byteJson, err := json.Marshal(comment)
	// 將結果以字串做打印
	fmt.Println(string(byteJson))

	// Post Website
	// 將Request body 轉成Reader
	jsonReader := bytes.NewReader(byteJson)
	// 將Request 以 Post方式，送至遠端網站
	postResp, _ := http.Post("http://jsonplaceholder.typicode.com/comments", "application/json", jsonReader)
	postRespByte,_ := ioutil.ReadAll(postResp.Body)
	// 將結果以字串做打印
	fmt.Println(string(postRespByte[:]))

	// 進階型設定訪問

	// Request設定
	request, err := http.NewRequest("POST", "http://jsonplaceholder.typicode.com/comments", jsonReader)

	// 追加Header （Optional）
	request.Header.Set("Content-Type","application/json;charset=UTF-8")
	// 新建Client
	client := http.Client{}
	// 發送，並取得結果
	advanceHttpResp, _ := client.Do(request)
	advanceHttpRespBytes, err := ioutil.ReadAll(advanceHttpResp.Body)
	fmt.Println(string(advanceHttpRespBytes[:]))

}
*/