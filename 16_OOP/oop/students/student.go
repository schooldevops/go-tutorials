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
