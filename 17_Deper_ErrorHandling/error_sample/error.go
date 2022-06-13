// error를 생성하기 위한 패키지이다.
package errors

// New 함수는 에러 객체를 생성한다.
func New(text string) error {
	return &errorString{text}
}

// errorString 구조체로 스트링을 가진다.
type errorString struct {
	s string
}

// Error() 함수를 지정하며 이는 errorString의 포인터를 수신 받는다. 결과 값으로 string을 반환한다.
func (e *errorString) Error() string {
	return e.s
}
