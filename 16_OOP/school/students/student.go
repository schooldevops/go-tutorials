package students

type Student struct {
	Name    string
	Korean  int
	English int
	Math    int
	BodyInfo
}

func (s Student) TotalScore() float32 {
	return (float32)(s.Korean+s.English+s.Math) / 3.0
}
