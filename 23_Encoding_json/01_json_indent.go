package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
)

func main() {
	type Road struct {
		Name   string
		Number int
	}

	roads := []Road{
		{"Diamond Fork", 29},
		{"Sheep Creek", 51},
	}

	b, err := json.Marshal(roads)
	if err != nil {
		log.Fatal(err)
	}

	var out bytes.Buffer

	//	func Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error
	//	dst: 출력버퍼 (즉, 타겟 버퍼)
	//	src: 인덴트를 수행할 소스 바이트 배열
	//	prefis: 인덴트시 prefix
	//	indent: 인덴트 대체 문자열
	json.Indent(&out, b, "=", "\t")
	out.WriteTo(os.Stdout)
}
