package students

import "fmt"

type student struct {
	name    string
	korean  int
	english int
	math    int
}

func New(name string, korean int, english int, math int) student {
	s := student{name, korean, english, math}
	return s
}

func (s student) TotalScore() {
	fmt.Printf("%s's total score is %d\n", s.name, ((s.korean + s.english + s.math) / 3))
}
