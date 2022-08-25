package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://naver.com")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// bodyByte 로 서버로 부터 컨텐츠 내용을 읽어들인다.
	bodyByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bodyByte))
}
