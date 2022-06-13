# Go에서 테스팅 수행하기. 

- Go 에서 테스트를 수행하고 안전한 코드를 작성하는 것에 대해서 단계적으로 알아보자. 

## Testing Package

- 테스트를 위해서는 테스팅 패키지를 가져와야한다. 
- 여기에는 T 타입이 존재하며 이는 테스트에서 에러 혹은 실패등을 체크할 수 있다. 
  
```go
import "testing"
```

- 위와 같이 testing 패키지를 임포트 한다. 

## testing.T 

- testing.T 는 테스를 수행하고 error, failure 를 체크할 수 있도록 해준다. 
- t.Fail() 과 t.Error() 등의 메소드를 제공하고 있다. 

- t.Fail() 
  - 테스트가 실패 했음을 나타낸다. 다만 다른 테스트가 계속해서 수행될 수 있도록 한다. 
- t.FailNow()
  - 현재 테스트가 실패했는지 표시하고, 실행을 종료한다. 
- t.Fatal() 
  - 에러 메시지를 로그로 출력하고, t.FailNow() 메소드를 호출한다. 
- t.Errorf("expected %s, but got %s", expected, got) 
  - 테스트 에러를 로깅하고, 테스트가 실패 했음을 마킹한다. 
- t.Log() 
  - 에러 메시지를 노출하고, 실파된 위치를 노출한다. 



## 테스트 하기

- email_check.go 파일을 다음과 같이 작성한다. 
  
```go
package email

import "strings"

var freemails = []string{"gmail.com", "yahoo.com", "outlook.com"}

func IsFreemail(email string) bool {
	for _, provider := range freemails {
		if strings.Contains(email, provider) {
			return true
		}
	}
	return false
}

```

- email_check_test.go 파일을 다음과 같이 작성한다. 

```go
package email

import "testing"

func TestGmail(t *testing.T) {
	gmail := "user@gmail.com"
	if !IsFreemail(gmail) {
		t.Fail()
	}


```

- 테스트 실행하기. 

```go
go test

PASS
ok      github.com/schooldevops/go-test 0.340s
```

- 에러가 있는경우 다음과 같다. 

```go
go test 

    email_check_test.go:8: 
FAIL
exit status 1
FAIL    github.com/schooldevops/go-test 0.343s
```

- '-run=Pattern' 옵션을 이용하여 레귤러 익스프레션을 통해 테스트를 선정할 수 있다. 

```go
go test -run=G

PASS
ok      github.com/schooldevops/go-test 0.339s
```

```go
go test -run=Gmail

PASS
ok      github.com/schooldevops/go-test 0.170s
```

## t.Parallel() 

- t.Parallel() 메소드를 호출하여 다른 테스트와 병행으로 안전하게 테스트를 수행할 수 있다. 

```go
func TestGmail(t *testing.T) {
  t.Parallel()
  ...
}
```

- 테스트가 병렬로 수행되고 패키지 레벨 상태가 없다 그리고 다른 테스트가 수행될때 사이드 이펙트가 없다.