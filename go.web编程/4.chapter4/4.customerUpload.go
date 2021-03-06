package main

import (
	"fmt"
	"mime/multipart"
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func postFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	// 关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile",filename)
	if err != nil {
		fmt.Println("error writing to Buffer")
	}

	// 打开文件句柄操作

	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}

	//iocopy

	_,err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()


	resp, err := http.Post(targetUrl,contentType,bodyBuf)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	resp_body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}

func main(){
	targetUrl := "http://127.0.0.1:9090/upload"
	filename := "./bbb.csv"
	postFile(filename,targetUrl)
}