package students

type BodyInfo struct {
	Height float32
	Weight float32
}

func (b BodyInfo) Bmi() float32 {
	return (b.Weight / (b.Height / 100 * b.Height / 100))
}
