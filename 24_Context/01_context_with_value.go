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
