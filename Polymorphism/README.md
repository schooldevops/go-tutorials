# Go Polymorphism 구현하기 

- Go 에서 Polymorphism 을 구현하기 위해서는 interface를 이용하여 타입을 구성하고, 이를 실제 구현하는 방법이 있다. 
- 이번 아티클에서는 Polymorphism을 이용하여 도형의 넓이를 구하고 도형의 이름을 출력하는 예제를 만들어 볼 것이다. 

## Go 모듈 생성하기 

- Go 에서는 기본적으로 GOPATH에서 소스를 작성하는 것으로 가정하고 개발이 진행된다. 
- 그러나 이렇게 하나의 PATH로 지정하면 여러 프로그램을 개발할때 매우 불편할 것이다. 
- 이를 해결하기 위해서 Go에서는 모듈을 제공하고 각기 다른 모듈을 생성할 수 있도록 해주고 있다. 
- 우리는 polymorphism 이라는 모듈을 생성할 것이다. 

```py
go mod init polymorphism
```

- 위 결과 go.mod 파일이 생성된다. 

```go
module polymorphism

go 1.18
```

## Polymorphism 구현하기

- 이제 다형성 (Polymorphism) 을 구성하기 위해서 'geometries' 라는 디렉토리를 하나 생성할 것이다. 

### 디렉토리 구조 

```py
- Polymorphism
  - geometries
    - circle.go
    - geometry.go
    - rectangle.go
    - square.go
    - triangle.go
  - go.mod
  - main.go
```

- 위와 같이 디렉토리 구조와 계층을 유심히 확인하자.
- go module 의 경우 모듈 위치와 참조를 위해서 디렉토리 계층은 매우 중요하다. 

### 인터페이스 정의하기 

- 그리고 해당 디렉토리에 geometry.go 파일을 생성하여 인터페이스를 정의할 것이다. 

```go
package geometries

type Geometry interface {
	GetName() string
	Area() float64
}
```

- 패키지는 geometries 라고 지정했다. 
- 그리고 type으로 Geometry 라는 인터페이스를 지정했다. 
- 인터페이스에는 다음 메소드를 정의한다. 
  - GetName() string: 도형의 이름을 반환한다. 
  - Area() float64: 도형의 넓이를 계산하여 반환한다. 

### Circle 구현체 생성하기 

- 원의 넓이를 구하는 구현체를 생성할 것이다. 
- circle.go 파일을 생성하고 다음과 같이 작성하자. 

```go
package geometries 

import "math"

type Circle struct {
	Name string
	R float64
}

func (circle Circle) GetName() string {
	return circle.Name
}

func (circle Circle) Area() float64 {
	return math.Pi * circle.R * circle.R
}
```

- 위와 같이 패키지는 인터페이스와 동일하게 생성한다. 
- Circle 구조체 타입을 정의한다. 
- GetName, Area에 대한 각각의 메소드를 구현한다. 
- (circle Circle) 은 Circle 구조체에 대해 구현 메소드라는 것을 알려준다. 

### Rectangle 구현체 생성하기.

- 직 사각형의 넓이를 구하는 구현체를 생성할 것이다. 
- rectangle.go 파일을 생성하고 다음과 같이 작성하자. 

```go
package geometries

type Rectangle struct {
	Name string
	X    float64
	Y    float64
}

func (rectangle Rectangle) GetName() string {
	return rectangle.Name
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.X * rectangle.Y
}
```

- 역시 동일한 패키지인 geometries 내부에 있도록 지정한다. 
- 그리고 Rectangle 구조체를 생성한다. 이름, 밑변, 높이를 지정할 수 있도록 속성을 정의했다. 
- GetName, Area 메소드를 구현했다. 

### Square 구현체 생성하기 

- 정 사각형의 넓이를 구하는 구현체를 생성할 것이다. 
- square.go 파일을 생성하고 다음과 같이 작성하자. 

```go
package geometries

type Square struct {
	Name string
	X    float64
}

func (square Square) GetName() string {
	return square.Name
}

func (square Square) Area() float64 {
	return square.X * square.X
}
```

- 역시 동일한 패키지인 geometries 내부에 있도록 지정한다. 
- 그리고 Square 구조체를 생성한다. 이름, 한 변을 지정할 수 있도록 속성을 정의했다. 
- GetName, Area 메소드를 구현했다. 

### Triangle 구현체 생성하기 

- 삼각형의 넓이를 구하는 구현체를 생성할 것이다. 
- triangle.go 파일을 생성하고 다음과 같이 작성하자. 

```go
package geometries

type Triangle struct {
	Name string
	X    float64
	Y    float64
}

func (triangle Triangle) GetName() string {
	return triangle.Name
}

func (triangle Triangle) Area() float64 {
	return triangle.X * triangle.Y / 2
}
```

- 역시 동일한 패키지인 geometries 내부에 있도록 지정한다. 
- 그리고 Triangle 구조체를 생성한다. 이름, 밑변, 높이를 지정할 수 있도록 속성을 정의했다. 
- GetName, Area 메소드를 구현했다. 

## 다형성 이용하기. 

- 이제는 작성한 다형성 코드를 이용해 보자. 
- main.go 파일을 생성하고, 다음과 같이 작성한다. 

```go
package main

import (
	geometries "polymorphism/geometries"
	"fmt"
)

func main() {

	circle := geometries.Circle{Name: "원", R: 10}
	rectangle := geometries.Rectangle{Name: "직사각형", X: 5, Y: 10}
	square := geometries.Square{Name: "정사각형", X: 5}
	triangle := geometries.Triangle{Name: "삼각형", X: 10, Y: 5}

	geometries := []geometries.Geometry{circle, rectangle, square, triangle}
	for _, geometry := range geometries {
		fmt.Println(geometry.GetName())
		fmt.Printf("Area: %0.2f\n", geometry.Area())
		fmt.Println("----------------------")
	}
}
```

- import 구문에서 이전에 작성한 모듈을 임포트 하자. 이때 임포트를 위해서 "polymorphism" 모듈내에 "geometries" 패미지의 의미로 "polymorphism/geometries" 라고 지정한다.
- geometries.Circle 객체를 파라미터와 함께 생성한다. 나머지 도형들도 유사한 방식으로 생성할 것이다. 
- []geometries.Geometry 를 이용하여 슬라이스를 생성하였다. 
- 그리고 for 구분을 이용하여 슬라이스를 반복하면서 이름, 넓이를 출력한다. 

## WrapUp

- 지금까지 다형성을 Golang로 구현해 보았다. 
- 다형성을 위해서는 type, interface로 정의하고, 이후 동일한 패키지 내에서 구현체를 생성하고 있다. 
- 모듈을 생성하고, 패키지를 생성하는 방법도 알아보았다. 