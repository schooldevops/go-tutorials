# Context 알아보기

- context 는 어플리케이션에서 다양한 레이어 사이에 정보를 전달하는 역할을 수행한다. 
- 어플리케이션의 가장 외곽 부분에서 정보를 생성하고, 이 정보는 서비스 레이어를 통과하고, 저장 영역으로 전달이 되는 구조이며 이때 정보 전달을 Context를 통해서 수행이 가능하다. 
- Context 전달 내용
  - 추가적인 정보 : 어플리케이션의 레이어 체인에 추가정보를 전달
  - 취소에 대한 제어: 추가정보 이외에 전달된 정보가 언제까지 처리 되어야하는지, 해당시간에 종료하지 못한경우 취소 처리
- 중요: Context가 각 레이어에 값을 안전하게 전달하지만, 모든 곳에 Context로 전달하는것은 좋지 않다. 필요한 경우에만 사용하자. 
- context는 thread safe 하다. 변경이 발생하면 새로운 복사본을 생성하므로 어려 루틴에서도 안전하게 이용이 가능하다. 

## Context Methods

- context.TODO()
  - nil이 아닌 빈 컨텍스트 반환
  - 사용할지 아닐지 모르는 경우 TODO 함수를 이용하여 컨텍스트를 생성할 수 있다. 
- context.Background()
  - nil이 아닌 빈 컨텍스트 반환
  - canceled가 안되고, 값이 없고, 데드라인도 없는 컨텍스트 생성
  - 일반적으로 메인 함수에서 이용하며, 테스트, 요청의 최상위 컨텍스트 생성시 주로 이용
- context.WithValue()
  - 컨텍스트에 키/값 쌍으로 구성된 컨텍스트를 생성한다.
  - 부모 컨텍스트의 복제본을 생성하고 반환한다. 
  - 프로스세의 요청 스콥 내에서 꼭 전달해야하는 데이터를 전달할때 이용하며, 일반적인 파라미터용도가 아님을 명심하자. 
- context.WithCancel()
  - WithCancel은 새 Done 채널이 있는 부모 복사본을 반환한다.
  - 반환 컨텍스트의 Done 채널은 반환된 취소 함수가 호출될 때 닫힌다. 또는 상위 컨텍스트의 Done 채널이 닫힐 때 먼저 발생하는 것이다.
  - 이 컨텍스트를 취소하면 연관된 리소스가 해제되므로 코드는 이 컨텍스트에서 실행 중인 작업이 완료되는 즉시 취소를 호출해야한다.
- context.WithDeadline()
  - WithDeadline은 기한이 조정된 상위 컨텍스트의 복사본을 반환한다.
- context.WithTimeout()
  - WithDeadline과 동일한 처리를 수행한다. 
  - context.WithDeadline을 사용하면 컨텍스트가 종료되는 특정 time.Time을 제공
  - context.WithTimeout 함수를 사용하면 컨텍스트를 지속할 시간에 대한 time.Duration 값만 제공하면 된다.

## Context 에 Value 전달 (추가적인 정보 전달)

- 값을 전달하는 컨텍스트를 생성하고, 메소드에 전달하는 예제이다. 

```go
package main

import (
	"context"
	"fmt"
)

func createContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "my-key", "my-secret-key")
}

func doSomething(ctx context.Context) {
	secretKey := ctx.Value("my-key")
	fmt.Printf("Get my key from parent... key is [%v]\n", secretKey)
	secretKey2 := ctx.Value("my-key2")
	fmt.Printf("Get my second key from parent... key is [%v]\n", secretKey2)
}

func main() {
	fmt.Println("Context with Value.")
	defer fmt.Println("Done example.")
	// Background는 nil 이 아닌 기본 컨텍스트를 생성한다.
	// 취소되지 않고, 값이 없으며, 데드라인도 없는 컨텍스트를 생성한다.
	// 보통 메인함수, 테스트용도, 요청에 대해 최상위 컨텍스트를 생성할때 주로 이용한다.
	ctx := context.Background()

	// 컨텍스트를 withValue로 생성한다.
	ctx = createContext(ctx)

	// 컨텍스트를 전달한 함수 실행
	doSomething(ctx)
}

```

- 위 예제에서 context.Background() 함수를 이용하여 컨텍스트를 최상위 컨텍스트를 생성한다. 
- 이후 createContext 를 통해서 컨텍스트에 값을 전달하도록 WithValue메소드를 사용하고 있다. 
  - 참고. WithValue 함수는 기존 컨텍스트의 복사본을 반환하며, 원래 컨텍스트를 수정하지 않는다. (중요)
- ctx.Value('키') 메소드를 이용하면 키에 대한 값을 조회할 수 있다. 
- 참고로 키값이 context에 없다면 nil을 반환한다. 

### 결과보기 

```go
$ go run 01_context_with_value.go

Context with Value.
Get my key from parent... key is [my-secret-key]
Get my second key from parent... key is [<nil>]
Done example.
```

- 위와 같이 전달한 키에 대한 값을 조회하였다. 
- 두번째 키라고 해서 컨텍스트에 값이 존재하지 않는다면 nil 이 반환된다. 

## Deadline 이 있는 Context

- 처리를 완료하는데 특정 시간내에 처리가 되어야 하는 경우가 있다. 
- 이때 사용할 수 있는 것이 WithTimeout 메소드이다. 해당 시간에 처리를 완료하지 못한경우 취소된다. 

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func doSomethingDuringTime(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Time out !!!! You Failed.")
			return
		default:
			fmt.Printf("Get key from context... [%v]\n", ctx.Value("my-key"))
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Start Context with Timeout...")
	defer fmt.Println("End of example.")

	ctx := context.WithValue(context.Background(), "my-key", "Hello This is key...")
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	go doSomethingDuringTime(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("Process have exceeded dead line.")
	}

	time.Sleep(2 * time.Second)
}

```

### go routine 함수 생성하기

- 메인 함수에서 새로운 go routine을 실행하기 위해서 doSomethingDuringTime 메소드를 생성한다. 

```go
func doSomethingDuringTime(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Time out !!!! You Failed.")
			return
		default:
			fmt.Printf("Get key from context... [%v]\n", ctx.Value("my-key"))
		}
		time.Sleep(500 * time.Millisecond)
	}
}
```

- 위 함수는 컨텍스트를 전달 받았다. 이 컨텍스트는 withTimeout 을 통해서 제한시간이 있는 컨텍스트이다. 
- ctx.Done() 라는 채널 이벤트를 받았다면 "Time out !!!! You Failed." 라는 내용을 출력하고 루틴을 종료한다. 
- 그렇지 않은경우 Get Key ... 으로 전달받은 값을 출력한다. 
- 또한 매번 500 밀리초동안 쉬었다가 다시 loop를 반복한다 

### 컨텍스트 생성하고 go routine 실행하기

```go
func main() {
	fmt.Println("Start Context with Timeout...")
	defer fmt.Println("End of example.")

	ctx := context.WithValue(context.Background(), "my-key", "Hello This is key...")
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	go doSomethingDuringTime(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("Process have exceeded dead line.")
	}

	time.Sleep(2 * time.Second)
}
```

- context.Background() 는 최상위 기본 컨텍스트이다. 
- context.WithValue 를 통해서 기본 컨텍스트에 키/값 쌍으로 값을 할당한다. (역시 기본 컨텍스트의 복사본을 반환한다는 것이 중요하다.)
- context.WithTimeout 을 수행하여 컨텍스트에 타임아웃을 설정한다. 여기서는 2초를 설정했다. 
- 최종적으로 cancel() 메소드가 수행될 수 있도록 defer를 사용했다. 이를 이용하면 컨텍스트의 리소스를 타임아웃 이후에 릴리즈 하게 된다. 
- go doSomethingDuringTime 메소드를 go routine으로 수행한다. 
- 메인 태널에서 컨텍스트가 종료되기를 기다렸다가, 종료되면 "Progress have ... " 를 출력하고 종료된다. 
- time.Sleep 을 가장 마지막에 사용한 이유는 고루틴이 실행되고 2초동안 메인 루틴을 그대로 유지하기 위함이다. 

### 결과 확인하기 

```go
$ go run 02_context_with_deadline.go

Start Context with Timeout...
Get key from context... [Hello This is key...]
Get key from context... [Hello This is key...]
Get key from context... [Hello This is key...]
Get key from context... [Hello This is key...]
Process have exceeded dead line.
Time out !!!! You Failed.
End of example.
```

## Error 컨텍스트 이용하기 

- 컨텍스트 객체는 기능을 중지해야할 만한 오류를 반환해야할 때 Err() 를 이용할 수 있다. 
- Err() 함수를 호출하면 Done 이 아직 닫히지 않은경우 nil 을 반환한다. 
- Done이 닫혔다면 Err은 닫힌 이유를 설명하는 오류를 반환한다. 

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Time out occur...")
			err := ctx.Err()
			fmt.Println("Err: ", err)
			return
		default:
			fmt.Println("do something ..... ")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Start Context with Err...")
	defer fmt.Println("End of example.")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go doSomething(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("Process have exceeded dead line.")
	}

	time.Sleep(2 * time.Second)
}

```

### Go routine 생성하기 

```go
func doSomething(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Time out occur...")
			err := ctx.Err()
			fmt.Println("Err: ", err)
			return
		default:
			fmt.Println("do something ..... ")
		}
		time.Sleep(500 * time.Millisecond)
	}
}
```

- 컨텍스트 타임아웃이 발생한경우 Time out occur... 을 출력한다. 
  - 그리고 ctx.Err() 함수에서 에러 원인을 추출하고, 화면에 출력한다. 이후 루틴을 종료한다. 
- 그렇지 않으면 do something ..... 반복적으로 500ms 주기로 출력한다. 

### 메인함수 구현하기

```go
func main() {
	fmt.Println("Start Context with Err...")
	defer fmt.Println("End of example.")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go doSomething(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("Process have exceeded dead line.")
	}

	time.Sleep(2 * time.Second)
}
```

- 컨텍스트를 타임아웃과 함께 생성한다. 
  - context.Backgroune는 기본 최상위 컨텍스트이다. 
  - context.WithTimeout 함수를 이용하여 2초 동안 타임아웃을 설정했다. 
- defer cancel() 을 이용하여 타임아웃된 경우 리소스를 릴리즈 하도록 한다. 
- go doSomething() 을 호출하여 고 루틴을 실행한다. 
- 메인 함수에서 go routine 이 타임아웃 신호를 받을때 까지 대기한다. 
- 이후 2초동안 쉬고, 프로그램을 종료한다. 

### 결과 확인하기 

```go
$ go run 03_context_err.go 

Start Context with Err...
do something ..... 
do something ..... 
do something ..... 
do something ..... 
Process have exceeded dead line.
Time out occur...
Err:  context deadline exceeded
End of example.
```

## WrapUp

- 여기서는 컨텍스트가 무엇인지, 컨텍스트를 이용하여 값을 전달하는 context.WithValue 메소드의 사용 방법을 알아 보았다. 
- 그리고 context.WithTimeout 메소드를 통해서 타임아웃이 있는 컨텍스트를 전달하여 컨텍스트를 특정 시간동안 처리하도록 제한을 걸어보았다. 
- context.Err 함수를 이용하여 컨텍스트 정지 및 오류에 대해서 사유를 출력하고 확인도 해 보았다. 
- 컨텍스트의 경우 꼭 필요한 정보를 하위 루틴에 안전하게 전달하는 방법을 제시하고 있다. 

- 참고:
  - https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go
  - https://tutorialedge.net/golang/go-context-tutorial/