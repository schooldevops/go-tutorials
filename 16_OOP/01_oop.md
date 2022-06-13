# OOP (Object Oriented Protramming) 

- go는 순수한 객체지향 프로그래밍 언어가 아니다. 
- go는 타입과 메소드를 가지고 있으며, 이는 객체 지향 언어 스타일의 요소이다. 
- 타입에 대한 계층은 없다. 
- Go 인터페이스 개념은 사용하기가 쉽고 어떤 면에서는 더 일반적인 방식에서 다른 접근 방식을 제공한다. 
- 서브 클래싱과 유사하지만, 다른 방식을 제공하기 위해 포함된 타입을 제공하는 방법도 있다. 
- 더군다나 메소드들은 C++이나 Java보다 더 일반적이다. 
- 이들은 모든 종류의 데이터에 대해서 정의할 수 있으며, 심지어 일반 "Unbounded" 정수와 같은 내장 유형도 정의할 수 있다. 
- 이는 구조체 에 국한하지 않는다. 

## 클래스 대신 구조체 이용

- Go는 클래스를 제공하지 않는다. 그러나 structs를 제공한다. 
- 메소드들은 struct에 추가될 수 있다. 
- 이는 데이터와 메소드의 번들링을 제공한다. 이는 데이터와 함께 클래스에서 운영할 수 있도록 하는 방법이다. 

- 클래스를 생성하기 위해서 다음 작업에 따라 코딩하자. 

### 모듈 생성하기 

- oop 디렉토리를 다음과 같이 생성하고, 모듈을 초기화 하자. 

```go
mkdir oop
cd oop

go mod init oop
```

- 위 처리를 수행하고 다음과 같은 구조로 파일과 디렉토리를 생성하다. 

```go
-- Documents
  -- oop
    -- students
      -- student.go
    -- go.mod
    -- main.go
```

- 위 형태대로 디렉토리와 파일을 생성했다. 
- 그리고 student.go 파일을 다음과 같이 작성한다. 

```go
package students

import "fmt"

type Student struct {
	Name    string
	Korean  int
	English int
	Math    int
}

func (s Student) TotalScore() {
	fmt.Printf("%s's total score is %d\n", s.Name, ((s.Korean + s.English + s.Math) / 3))
}
```

- 위와 같이 struct와 TotalScore메소드를 생성하였다. 
- 해당 메소드를 생성하면서 리시버를 (s Student)로 사용하였으므로 해당 메소드는 Student 구조체의 메소드로 처리된다. 

```go
package main

import "oop/students"

func main() {
	kido := students.Student{
		Name:    "Kido",
		Korean:  95,
		English: 90,
		Math:    80,
	}

	kido.TotalScore()
}

```

- 위 처리를 위해서 oop/students 를 임포트 했다. 
- 그리고 students.Student 를 이용하여 객체 인스턴스를 생성했다. 
- 이후 TotalScore를 호출할 수 있게 되었다. 

- 처리 결과는 다음과 같다. 

```go
Kido's total score is 88
```

## 생성자 대신에 New 함수를 이용한다. 

- Go에서는 New 메소드를 이용하여 생성자를 생성할 수 있다. 

```go
package students

import "fmt"

type student struct {
	name    string
	korean  int
	english int
	math    int
}

func New(name string, korean int, english int, math int) student {
  s := student {name, korean, english, math}
  return s
}

func (s student) TotalScore() {
	fmt.Printf("%s's total score is %d\n", s.Name, ((s.Korean + s.English + s.Math) / 3))
}
```

- 다음 main.go 파일을 작성하고 아래와 같이 코드를 작성하자. 

```go
package main

import (
	"oop_new/students"
)

func main() {
	s := students.New("Kido", 95, 90, 80)
	s.TotalScore()
}
```

- 위와 같이 New를 이용하여 접근할 수 있다. 
- Go 에서는 메소드, 필드는 대문자로 시작하는 이름을 작성하면, public 접근이 가능하다. 
- 소문자로 시작하는 메소드 필드는 private 접근으로 제한된다.
- 그러므로 New 메소드를 이용하여 객체 인스턴스를 생성하도록 작성한 것이다. 

## WrapUp 

- 위와 같이 Go에서는 공식적으로 클래스를 지원하지 않는다. 
- 그러므로 Go에서는 struct와 함수와 리시버를 이용하여 구현한다. 
- 생성자를 사용할때에는 private를 이용하여, struct를 이용하여 직접 접근할 수 없도록 설정한다. 
- 그리고 New함수를 이용하여 인스턴스를 생성할 수 있도록 작업할 수 있음을 알 수 있었다. 

