# Composition

- Go에서는 상속을 지원하지 않는다. 
- 다만 Composition을 지원하며, 여러 struct를 생성해두고, 필요한경우 생성하고자 하는 struct에서 포함관계로 객체를 생성할 수 있도록 지원한다. 

## 샘플코드

- 다음은 학생 구조체가 있고, 학생은 다시 신체사항 정보를 가지는 예를 수행해 보자. 

```go
package students

import (
  "fmt"
)

type BodyInfo struct {
  Height float32
  Weight float32
}

func (b BodyInfo) Bmi() float32 {
	return (b.Weight / (b.Height * b.Height))
}
```

- 이제 학생 정보를 저장하는 struct를 생성해 보자. 

```go
package students

import ("fmt")

type Student struct {
  Name string
  Korean int
  English int
  Math int
  BodyInfo
  
}

func (s Student) TotalScore() float32 {
  return (float32)(s.Korean + s.English + s.Math) / 3.0
}
```

- 위 코드와 같이 Student struct는 BodyInfo를 포함하고 있는 Composition관계를 가진다. 

- 이제 위 코드를 사용하는 main 코드를 작성하자. 

```go
package main

import (
  "fmt"
  "school/students"
)

func main() {
  bodyInfo := students.BodyInfo{
    Weight: 82, 
    Height: 177,
  }

  kido := students.Student{
    Name: "Kido",
    Korean: 95,
    English: 90,
    Math: 80,
    bodyInfo,
  }

  fmt.Println("Student Name: ", kido.Name)
  fmt.Println("Total Score: ", kido.TotalScore())
  fmt.Println("Bmi:", kido.Bmi())
}
```

- 위 처리결과는 다음과 같다. 

```go
Student Name:  Kido
Total Score:  88.333336
Bmi: 26.173832
```

- 위 코드는 우선 BodyInfo라는 구조체를 생성하고, Student라는 구조체를 생성하면서 BodyInfo를 포함하도록 생성했다. 
- 이런 포함관계를 Composition 구조이며, BodyInfo 인스턴스를 생성하고, Student에 파라미터로 전달하였다.