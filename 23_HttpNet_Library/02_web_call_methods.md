# net/http 패키지로 REST API 호출하기 

- net/http 를 이용하면 REST API서버로 API요청을 수행할 수 있다. 

## Get 메소드로 데이터 조회하기 

- http.Get 메소드를 이용하면 서버에 Get 메소드 호출을 하여 데이터를 조회할 수 있다. 

```go
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

```

- resp, err := http.Get(서버주소)
  - 위 내용과 같이 서버 주소에 접근하여 결과값을 resp에 담는다. 
  - 오류가 발생한경우 err에 오류값을 확인할 수 있다. 
- defer resp.Body.Close() 
  - defer 를 이용하여 함수가 종료되면 요청 핸들러를 닫는다. 
- bodyByte, err := ioutil.ReadAll(resp.Body)
  - 응답값은 resp.Body 를 이용하여 접근할 수 있다. 
  - 정상적으로 읽어들인 값을 bodyByte라는 변수에 할당한다. 
- fmt.Println(string(bodyByte))
  - 읽어들인 바이트 배열을 string으로 변환할 수 있다. 

## Echo 서버 생성하기

- 우선 POST 로 요청을 받아서 응답에 날짜를 추가한 echo 서버를 구성하자. 

```go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/echo", mainPage)
	fmt.Println("Server Start with 8080 port.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type UserInfo struct {
	Name string
	Age  int
	Job  string
	Tm   time.Time
}

func mainPage(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":

		b, err := io.ReadAll(r.Body)

		if err != nil {
			log.Fatal(err)
		}

    var userInfo UserInfo
		json.Unmarshal(b, &userInfo)
		userInfo.Tm = time.Now()
		writeByte, _ := json.Marshal(userInfo)

    w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(writeByte)
	}
}

```

- 요청으로 부터 r.Method 를 이용하면 어떠한 형태의 Method가 요청되었는지 파악할 수 있다. 
- switch 를 이용하여 요청값을 비교하고 POST 인 경우만 동작하게 코드를 작성하였다 
- b, err := io.ReadAll(r.Body)
  - 위 코드를 이용하면 요청 body로 부터 모든 바이트를 읽어 들인다. 
- json.Unmarshal(b, &userInfo)
  - 타입으로 요청 정보를 언 마샬링한다. 
  - 언 마샬링은 바이트를 UserInfo타입 인스턴스로 변환한다. 
- time.Now() 
  - 이 메소드를 이용하여 현재 날짜를 조회한다. 
- writeByte, _ := json.Marshal(userInfo)
  - json.Marshal 을 이용하면 인스턴스 값을 바이트로 변환한다. 
- w.Header().Set("Content-Type", "application/json")
  - 응답 헤더의 Content-Type 값을 application/json으로 변환한다. 
- w.WriteHeader(http.StatusOK)
  - 처리 결과를 OK로 반환한다. 
- w.Write(writeByte)
  - 변환된 데이터를 쓰기한다. 

### 서버 수행하기 

```go
$ go run 01_simple_server_post.go

Server Start with 8080 port.
```

## POST 요청 수행하기 

- Post 는 특정 URL에 post방식으로 요청을 보낼 수 있다. 
- 여기서는 json 타입으로 post요청을 수행하고, 응닶 값을 json 객체로 받아들이는 작업을 수행한다. 

```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	requestData := map[string]any{"Name": "Kido", "Age": 20, "Job": "Developer"}
	json_data, err := json.Marshal(requestData)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://localhost:8080/echo", "application/json", bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	var resMap map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&resMap)

	fmt.Println(resMap)

}

```

- requestData := map[string]any{"Name": "Kido", "Age": 20, "Job": "Developer"}
  - 요청할 데이터를 위와 같이 작성하였다. 
  - 이름, 나이, 직업을 맵으로 받았다. 
- json_data, err := json.Marshal(requestData)
  - 요청할 데이터를 json_data로 변환하기 위해서 json.Marshal 메소드를 이용했다. 
- resp, err := http.Post("http://localhost:8080/echo", "application/json", bytes.NewBuffer(json_data))
  - http.Post 로 서버로 요청을 보낸다. 
  - application/json 형색의 컨텐츠 타입으로 요청한다. 
  - bytes.NewBuffer 를 이용하여 요청 내용으로 요청한다. 
  - 응답값은 resp로 응답한다. 
- var resMap map[string]interface{}
  - 요청한 값을 받을 맵을 구성한다. 
  - interface{} 라고 값타입을 지정하면 어떠한 json도 변경이 가능하다. 
- json.NewDecoder(resp.Body).Decode(&resMap)
  - 새로운 디코더를 이용하여 응답받은 값을 맵으로 변환한다. 
- fmt.Println(resMap)
  - 결과 내용을 출력했다. 

### 요청 프로그램 수행하기 

```go
$ go run 02_sample_post_json_methods.go

map[Age:20 Job:Developer Name:Kido Tm:2022-08-25T19:28:08.175353+09:00]
```

## Wrap Up

- 지금까지 Get, Post로 요청을 해 보았다. 
- Get, Post 등으로 요청을 하고, 응답값 파싱해 보았다. 
- 또한 Post를 처리하는 서버를 구성하고, 서버의 응답을 클라이언트에서 파싱하는 방법도 알아 보았다. 
- json.Marshal, json.Unmarshal 을 이용하여 바이트로 변경하거나, 오브젝트로 변경하는 작업도 할 수 있다는 것을 확인했다. 