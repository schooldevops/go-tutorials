# Json Encoding

- Json 은 일반적인 어플리케이션에서 가장 자주 사용하는 데이터 표기법일 것이다. 
- Javascript Object Notation 으로 객체를 표현하는 방법으로 사용법이 간단하고, 객체를 표현하는데 매우 편리하다. 

## Decoder

- Json 으로 표현된 문자열 데이터를 읽어서 (문자열 변수 or 파일변수) 타입에 할당하는 예를 살펴보자. 

```go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	const jsonStream = `
		{"Name": "Ed", "Text": "Knock knock."}
		{"Name": "Sam", "Text": "Who's there?"}
		{"Name": "Ed", "Text": "Go fmt."}
		{"Name": "Sam", "Text": "Go fmt who?"}
		{"Name": "Ed", "Text": "Go fmt yourself!"}
	`

	type Message struct {
		Name, Text string
	}

	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Text)
	}
}

```

### 복수 라인 JSON 지정하기 

- 문자열을 복수 라인으로 표현하기 위해서 다음과 같이 지정할 수 있다. 

```go
	const jsonStream = `
		{"Name": "Ed", "Text": "Knock knock."}
		{"Name": "Sam", "Text": "Who's there?"}
		{"Name": "Ed", "Text": "Go fmt."}
		{"Name": "Sam", "Text": "Go fmt who?"}
		{"Name": "Ed", "Text": "Go fmt yourself!"}
	`
```

- '`' 을 이용하여 문자열을 작성하는 경우 복수 라인으로 문자열을 정의할 수 있다. 
- 표기법의 경우 json 표현식 한 라인씩 지정하였다. 이런 형식을 jsonl (json line) 이라고 표현한다. 

### 타입 설정하기 

```go
	type Message struct {
		Name, Text string
	}
```

- Message 라는 구조체 타입을 지정하였다. 
- 내용은 Name, Text 로 문자열을 담는 그릇을 만들었다. 

### 문자열 스트림을 디코딩하기

```go
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Text)
	}
}
```

- 위와 깉이 json.NewDecoder 객체를 이용하여 문자열을 스트림으로 읽어들였다. 
- 이후 dec.Decode 를 통해서 각 라인단위로 JSON 데이터를 읽어서 Message 구조체에 담았다. 
- err == io.EOF 일때까지 루프를 반복하고, 파일/스트림의 끝까지 읽었다면 break로 종료한다. 
- 실제 err이 발생된경우 log.Fatal을 통해서 출력한다. 
- 이후 결과를 출력한다. 

### 결과 보기 

```go
$go run 01_decoder_json.go 

Ed: Knock knock.
Sam: Who's there?
Ed: Go fmt.
Sam: Go fmt who?
Ed: Go fmt yourself!
```

- 결과는 이름: 텍스트의 형태로 출력이 되었음을 알 수 있다. 

## HTML Escape 수행하기. 

- 웹 프런트로 부터 데이터를 읽어 들이면 html 태그 혹은 Javascript 코드 등과 같은 위험한 코드가 들어와서 Inject 공격을 수행할 수 있게 된다. 
- 이 경우 태깅 문자열들을 Escape 하여 화면에 출력될때에만 해당 문자 타입이 변환되도록 코드 변환이 필요하다. 

```go
package main

import (
	"bytes"
	"encoding/json"
	"os"
)

func main() {
	var out bytes.Buffer
	json.HTMLEscape(&out, []byte(`{"Name":"<b>HTML content</b>"}`))
	out.WriteTo(os.Stdout)
}

```

- json.HTMLEscape 라는 메소드를 이용하여 입력된 문자열을 출력 버퍼로 복사하게 된다. 
- 복사하는 데이터는 escape된 문자로 변환되어 저장된다. 


### 결과보기 

```go
$ go run 01_html_escapes.go 

{"Name":"\u003cb\u003eHTML content\u003c/b\u003e"}
```

- 위와 같이 `{"Name":"<b>HTML content</b>"}` 가 인코딩 되어 {"Name":"\u003cb\u003eHTML content\u003c/b\u003e"} 값으로 변환되어 출력됨을 확인할 수 있다. 

## Marshal 수행하기

- marshal 은 Object 혹은 데이터 객체를 json 으로 변환하는 역할을 수행한다. 

```go
package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}

	group := ColorGroup{
		ID:     1,
		Name:   "Redis",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	b, err := json.Marshal(group)
	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.Write(b)
}

```

### 구조체 타입 지정하기. 

```go
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
```

- ColorGroup 구조체를 지정하였다. ID, Name, Color 슬라이스를 지정하였다. 

### 구조체 객체 생성하기. 

```go
	group := ColorGroup{
		ID:     1,
		Name:   "Redis",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
```

- 구조체를 이용하여 ColorGroup 객체를 생성하였다. 

### 마샬링 수행하기 

```go
	b, err := json.Marshal(group)
	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.Write(b)
```

- json.Marshal 을 이용하여 group 객체를 json 객체로 변환한다. 

### 출력결과 확인하기

```go
$ go run 01_json_marshal.go 

{"ID":1,"Name":"Redis","Colors":["Crimson","Red","Ruby","Maroon"]}
```

- 우리가 지정한 객체가 json 형식으로 마샬링 되었다. 

## UnMarshal

- 마샬링은 객체를 json 으로 변경하는 작업이다. 
- 이제는 json을 객체로 매핑할 차례이다. 이것을 UnMarshaling 이라고 한다. 

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var jsonBlob = []byte(
		`[
			{"Name": "Platypus", "Order": "Monotremata"},
			{"Name": "Quoll", "Order": "Dasyuromorphia"}
		]`)

	type Animal struct {
		Name  string
		Order string
	}

	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", animals)
}
```

### 바이트 배열 준비

```go
	var jsonBlob = []byte(
		`[
			{"Name": "Platypus", "Order": "Monotremata"},
			{"Name": "Quoll", "Order": "Dasyuromorphia"}
		]`)
```

- 위 내용과 같이 byte 배열을 이용하여 json 문자열을 초기화 했다. 
- 만약 파일에서 json 문자열 데이터가 조회된다면 위와 같은 형식이 될 것이다. 

### json 객체를 언마샬링해서 담을 구조체 생성하기  

```go
	type Animal struct {
		Name  string
		Order string
	}
```

- 위와 같이 Animal 이라는 이름으로 구조체를 생성했다. 
- 이는 json 객체를 담을 그릇으로 생각하면 된다. 

### 

```go
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", animals)
```

- 이제 animals 를 담을 슬라이스를 정의했다. 
- 그리고 json.Unamrshal 메소드를 이용하여 jsonBlob를 언마샬링해서 animals 에 추 대입하는 것을 확인하자. 
- 참고로 배열 슬라이스는 배열의 주소가 필요하다. 그러므로 &animals 라고 값을 전달했다. 

### 결과 확인하기. 

- 이제 결과를 확인하자. 

```go
$ go run 01_json_unmarshal.go 

[{Name:Platypus Order:Monotremata} {Name:Quoll Order:Dasyuromorphia}]
```

- 위와 같이 언마샬링된 값을 출력하면 animals 라는 슬라이스값이 화면에 출력됨을 확인할 수 있다. 

## json indent 적용하여 출력하기. 

- json역시 문자열이기 때문에 출력하면 일렬로 나열이 된다. 
- 사용자가 이 json 문자열을 확인하기 위해서는 indent를 이용하여 적절히 포매팅 되어야 한다. 

```go
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

```

### 타입 지정하기.

- 이제 우리가 데이터를 담을 타입 구조체를 정의하자. 

```go
	type Road struct {
		Name   string
		Number int
	}
```

- 이름, 숫자를 담는 Road라는 타입을 구성하였다. 

### 데이터 변수에 담기 

```go
	roads := []Road{
		{"Diamond Fork", 29},
		{"Sheep Creek", 51},
	}
```

- Road 타입 슬라이스를 새로 생성하였다. 

### 인덴트로 출력하기 

```go
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
```

- 타입 데이터를 json 형식으로 변환하였다. 이때 Marshal 이라는 함수를 이용한다. 
- json 타입으로 변형을 하게 되면 byte 버퍼로 담긴다. b 변수는 이렇게 변경된 데이터를 담고 있는 버퍼이다. 
- json.Indent 를 이용하여 소스(b)에서 목적지 버포(out) 으로 인덴트를 적용하여 변환한후 변수에 할당한다. 
- prefix는 '=' 으로 지정하였고, 인덴트는 '\t' 로 탭으로 인덴트를 부여했다. 

### 결과 확인하기. 

```go
$ go run 01_json_indent.go 

[
=       {
=               "Name": "Diamond Fork",
=               "Number": 29
=       },
=       {
=               "Name": "Sheep Creek",
=               "Number": 51
=       }
=]
```

- 위와 같이 json 으로 변형된 객체는 인덴트(탭)이 적용되어 출력이 되었다. 
- prefix는 = 으로 출력된 것을 알 수 있다. 

## 종합 하기 및 json 매핑 이용하기 

- 이번에는 지금까지 내용을 모두 종합해서 코드를 작성하자. 
- 그리고 json 데이터를 언마샬링 할때, 구조체에 타입 매핑을 지정하는 방법도 함게 알아보자. 

```go
package main

import (
	"encoding/json"
	"fmt"
)

type userInfo struct {
	Id   int
	Name string
	Age  int
	Tall int
}

type userInfo2 struct {
	Id   int    `json:"id"`
	Name string `json:"name,omitempty"`
	Age  int    `json:"age"`
	Tall int    `json:"tall"`
}

func main() {

	user01 := &userInfo{
		Id:   1,
		Name: "Kido",
		Age:  30,
		Tall: 180,
	}
	user01_m, _ := json.Marshal(user01)
	fmt.Println(string(user01_m))

	user02 := &userInfo2{
		Id:   2,
		Name: "Kido2",
		Age:  33,
		Tall: 177,
	}
	user02_m, _ := json.Marshal(user02)
	fmt.Println(string(user02_m))

	user03 := []byte(`{"id":3,"name":"kido3", "age": 40, "tall": 181}`)

	var user03_m map[string]interface{}

	if err := json.Unmarshal(user03, &user03_m); err != nil {
		panic(err)
	}
	fmt.Println(user03_m)

	num := user03_m["age"].(float64)
	fmt.Println(num)

	str := `{"id":4,"name":"kido4", "age": 46, "tall": 176}`
	res := userInfo2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Name)
}

```

### 구조체 지정하기 

```go
type userInfo struct {
	Id   int
	Name string
	Age  int
	Tall int
}
```

- 지금까지 보아왔던 구조체 설정 케이스이다. 

### 구조체 json 매핑 정의하기 

```go
type userInfo2 struct {
	Id   int    `json:"id"`
	Name string `json:"name,omitempty"`
	Age  int    `json:"age"`
	Tall int    `json:"tall"`
}
```

- 위 내용은 json 의 키와 struct 구조체의 필드를 매핑하였다. 
- 'json:"키이름"' 형식이 기본 형태이다. 
- 값이 비어 있다면 빈값을 넣을 경우 omitempty 를 지정하면 json 값이 비어있는경우 빈값으로 세팅하게 된다. (기본값 세팅)

### 기본 구조체 마샬링 

```go
	user01 := &userInfo{
		Id:   1,
		Name: "Kido",
		Age:  30,
		Tall: 180,
	}
	user01_m, _ := json.Marshal(user01)
	fmt.Println(string(user01_m))
```

- userInfo 를 이용하여 json으로 마샬링 한다. 
- 결과는 다음과 같다. 
  - {"Id":1,"Name":"Kido","Age":30,"Tall":180}

### 구조체 json 매핑 마샬링 

```go
	user02 := &userInfo2{
		Id:   2,
		Name: "Kido2",
		Age:  33,
		Tall: 177,
	}
	user02_m, _ := json.Marshal(user02)
	fmt.Println(string(user02_m))
```

- 위와 같이 매핑된 객체를 마샬링한다. 
- 결과는 다음과 같다. 
  - {"id":2,"name":"Kido2","age":33,"tall":177}

- 이 경우 매핑된 키 이름으로 출력됨을 확인할 수 있다. 

### 바이트 배열 언마샬링

```go
	user03 := []byte(`{"id":3,"name":"kido3", "age": 40, "tall": 181}`)

	var user03_m map[string]interface{}

	if err := json.Unmarshal(user03, &user03_m); err != nil {
		panic(err)
	}
	fmt.Println(user03_m)

	num := user03_m["age"].(float64)
	fmt.Println(num)
```

- 위와 같이 바이트 배열을 보면 id, name 등이 첫글자가 소문자이다. 
- 그리고 map 타입으로 언마샬링 한다. 
- 결과는 다음과 같다. 
	- map[age:40 id:3 name:kido3 tall:181]
	- 40 <-- 맵에서 값을 가져온 케이스 


### userInfo2 로 구조체 json 객체 언마샬링 

```go
	str := `{"id":4,"name":"kido4", "age": 46, "tall": 176}`
	res := userInfo2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Name)
```

- 위와 같이 소문자로된 키를 가진 문자열을 정의했다. 
- 이후 userInfo2 객체로 언마샬링을 수행한다. 
- 결과는 다음과 같다. 
  - {4 kido4 46 176}
  - kido4

## WrapUp

- 지금까지 json 객체를 golang에서 어떻게 변환하는지 알아보았다. 
- json을 golang 객체로 으로 전환하는 언마샬링, 객체를 json 으로 변환하는 마샬링에 대해서 알수 있게 되었다.
- 또한 인덴트 및 웹 개발시 필요한 이스케이프 처리도 같이 확인해 보았다. 

