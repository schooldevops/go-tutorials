# Simple Web Application 개발하기

- golang의 net/http 패키지를 이용하면 기본적인 웹 어플리케이션을 개발할 수 있다. 
- 다양한 프레임워크가 있으나 여기서는 net/http만을 이용하는 방법을 알아볼 것이다. 
- 우리는 여기서 REST API만을 구성할 것이다. 

## 기본 페이지 구성하기

- endpoint: /
- response: 
  - 응답값으로 {"Hello": "World"} 라는 값을 반환한다. 

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"Hello\": \"World\"}")
}

func main() {
	http.HandleFunc("/", mainPage)
	fmt.Println("Server Start with 8080 port.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

```

- func mainPage(w http.ResponseWriter, r *http.Request)
  - 핸들러 함수를 작성하는 기본 폼이다. 
  - http.ResponseWriter 는 HTTP 응답을 내려보내주는 Writer 이다. 
  - *http.Request 는 HTTP 요청을 받아들인다. 이때 포인터 타입으로 전달되는 것을 확인하자. 
- fmt.Fprintf(w, "{\"Hello\": \"World\"}")
  - Fprintf 는 파일 스트름으로 결과를 작성하는 함수이다. 
  - w 아규먼트로 어떤 Writer에 쓰기를 할지 지정하고, 내용을 작성하고 있다. 
- http.HandleFunc("/", mainPage)
  - http.HandleFunc 함수는 http uri에 대해서 어떠한 핸들러가 이 작업을 처리할지 등록해주는 작업을 한다. 
- log.Fatal
  - 함수가 정상으로 수행되면 그냥 지나가지만, 에러가 있으면 Fatal 로 결과를 로깅한다. 
- http.ListenAndServe
  - http 서버를 기동하고 요청을 대기한다. 


### 결과 확인하기 

- 아래 명령어로 서버를 기동한다. 
  
```go
$ go run 01_simple_server_basic.go

Server Start with 8080 port.
```

- 다음으로 내용을 확인하자. 

```go
$ curl -i http://localhost:8080

HTTP/1.1 200 OK
Date: Wed, 24 Aug 2022 21:36:43 GMT
Content-Length: 18
Content-Type: text/plain; charset=utf-8

{"Hello": "World"}
```

- 기본 페이지 요청에 대한 정상적인 응답을 확인할 수 있다. 

## html 파일을 읽어서 브라우저에 보여주기 

- 이번에는 /view/파일명 을 브라우저에 기술하면 서버는 파일명을 읽어서 브라우저로 내용을 보여주는 서비스를 만들어 보자.

```go
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/view/", viewPage)
	fmt.Println("Server Start with 8080 port.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"Hello\": \"World\"}")
}

func viewPage(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/view/"):]
	fmt.Println("--------- ", path)
	fileName := fmt.Sprintf("%s.%s", path, "html")
	writeFileToHttp(w, fileName)
}

func writeFileToHttp(w http.ResponseWriter, fileName string) {

	buff := make([]byte, 1024)

	sourceFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Fail to oepn source file/n", err)
	}

	defer sourceFile.Close()

	for {
		readSize, err := sourceFile.Read(buff)
		if err != nil && err != io.EOF {
			fmt.Errorf("Error occur when read file data.", err)
		}

		if readSize == 0 {
			break
		}

		writeSize, err := w.Write(buff[:readSize])
		if err != nil {
			fmt.Errorf("Error occur write data to dest file", err)
		}
		fmt.Printf("WriteFile Size is %d \n", writeSize)
	}
}

```

### view Handler 살펴보기 

- /view/ uri 에 대한 요청이 들어오면 처리하는 핸들러 함수를 다음과 같이 정의했다. 

```go
func viewPage(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/view/"):]
	fmt.Println("--------- ", path)
	fileName := fmt.Sprintf("%s.%s", path, "html")
	writeFileToHttp(w, fileName)
}
```

- 위 내용을 보면 Request 에서 URL.Path 를 통해서 현재 들어온 url 경로를 추출할 수 있다. 
- 경로에서 /view/ 이후 파일이름을 추출하여 writeFileToHttp 를 통해서 파일을 읽어 쓰기를 수행하는 메소드를 호출한다. 

### 파일을 읽어서 HTTP 로 내보내기

- 이제 파일 이름을 알았다면, 해당 파일을 읽어서 HTTP로 응답하는 코드를 작성하자. 

```go
func writeFileToHttp(w http.ResponseWriter, fileName string) {

	buff := make([]byte, 1024)

	sourceFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Fail to oepn source file/n", err)
	}

	defer sourceFile.Close()

	for {
		readSize, err := sourceFile.Read(buff)
		if err != nil && err != io.EOF {
			fmt.Errorf("Error occur when read file data.", err)
		}

		if readSize == 0 {
			break
		}

		writeSize, err := w.Write(buff[:readSize])
		if err != nil {
			fmt.Errorf("Error occur write data to dest file", err)
		}
		fmt.Printf("WriteFile Size is %d \n", writeSize)
	}
}
```

- 위 입력된 파일 이름을 대상으로 서버내 디렉토리에서 파일을 찾는다. 
- 찾은 파일을 버퍼로 읽어 들이고, 해당 내용을 바로 http.ResponseWriter 로 쓰기를 수행한다. 
- 파일의 끝에 도달할때까지 반복해서 내용을 Writer로 쓰게 된다. 

### html 파일 보기 

- test.html 파일을 다음과 같이 작성하였다. 

```html
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Test Hello</title>
</head>
<body>
  <h1>Hello World</h1>
</body>
</html>
```

## html/template 패키지 이용하기 

- html/template 패키지를 이용하면 HTML을 자동으로 파싱하여 브라우저로 전달하고, 동적으로 변하는 변수 영역을 플레이스홀더화 하여 교체하기 쉬운 장점이 있다. 

### 전체코드 살펴보기 

```go
package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/view/", viewPage)
	http.HandleFunc("/template/", viewTemplate)
	fmt.Println("Server Start with 8080 port.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"Hello\": \"World\"}")
}

func viewPage(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/view/"):]
	fmt.Println("--------- ", path)
	fileName := fmt.Sprintf("%s.%s", path, "html")
	writeFileToHttp(w, fileName)
}

type templateData struct {
	Title    string
	Greeting string
}

func viewTemplate(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/template/"):]
	fmt.Println("--------- ", path)
	fileName := fmt.Sprintf("%s.%s", path, "html")

	p := templateData{
		Title:    "Go Study",
		Greeting: "Hello World Study Mates",
	}

	t, _ := template.ParseFiles(fileName)
	t.Execute(w, p)
}

func writeFileToHttp(w http.ResponseWriter, fileName string) {

	buff := make([]byte, 1024)

	sourceFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Fail to oepn source file/n", err)
	}

	defer sourceFile.Close()

	for {
		readSize, err := sourceFile.Read(buff)
		if err != nil && err != io.EOF {
			fmt.Errorf("Error occur when read file data.", err)
		}

		if readSize == 0 {
			break
		}

		writeSize, err := w.Write(buff[:readSize])
		if err != nil {
			fmt.Errorf("Error occur write data to dest file", err)
		}
		fmt.Printf("WriteFile Size is %d \n", writeSize)
	}
}

```

- 위 코드는 이전 코드에서 연결하여 /template/ uri 에 대해서 요청을 처리하는 코드를 추가하였다. 
- 이전 코드와 차이점은 html/template 패키지를 이용하여 html 파일을 파싱하고, 교체해야할 내용을 execute 를 통해서 교체한다. 

### 구조체 작성하기 

- 우선 다음과 같은 templateData 구조체를 구성하였다. 
- 구조체는 Title, Greeting 을 추가하였다. 
  
```go
type templateData struct {
	Title    string
	Greeting string
}
```

### Template 파일 읽고, 내용 교체하기 

- 다음 코드는 Template를 이용하여 html을 읽고, 교체해야할 문자를 교체하는 코드이다. 
- html을 직접 읽어서 replace 하는 것보다 간단하게 작업이 처리가 된다. 

```go
func viewTemplate(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/template/"):]
	fmt.Println("--------- ", path)
	fileName := fmt.Sprintf("%s.%s", path, "html")

	p := templateData{
		Title:    "Go Study",
		Greeting: "Hello World Study Mates",
	}

	t, _ := template.ParseFiles(fileName)
	t.Execute(w, p)
}
```

- template.ParseFiles 를 이용하여 html 파일을 파싱하고 플레이스홀더 영역을 확인한다. 
- t.Execute(Write, TemplateData) 의 형식으로 템플릿의 플레이스 홀더를 교체하게 된다. 

### 핸들러 등록하기. 

```go
http.HandleFunc("/template/", viewTemplate)
```

- 위 코드와 같이 핸들러를 등록하였다. 

### template 파일 확인하기. 

```html
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>{{.Title}}</title>
</head>
<body>
  <h1>{{.Greeting}}</h1>
</body>
</html>
```

- 위 코드에서와 같이 {{.Title}}, {{.Greeting}} 등과 같은 플레이스 홀더를 이용하여 값이 들어오면 템플릿이 교체하게 된다. 

## WrapUp

- 지금까지 간단한 서버를 구성해 보았다. 
- golang 에 있는 net/http 패키지와 html/template 패키지를 이용하면 간단하게 브라우저로 부터 들어오는 요청을 처리하여 결과를 브라우저로 전송할 수 있게 된다. 