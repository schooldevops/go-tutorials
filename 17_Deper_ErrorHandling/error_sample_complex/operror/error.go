// error를 생성하기 위한 패키지이다.
package operror

import "fmt"

// operationError 구조체로 스트링을 가진다.
type OperationError struct {
	err     string
	width   float64
	height  float64
	dividen float64
}

// Error() 함수를 지정하며 이는 OperationError의 포인터를 수신 받는다. 결과 값으로 string을 반환한다.
func (e *OperationError) Error() string {
	return fmt.Sprintf("%s, errorValue is (width: %f, height: %f, dividen: %f)", e.err, e.width, e.height, e.dividen)
}

func (e *OperationError) IsZeroDividen() bool {
	return e.dividen <= 0
}

func (e *OperationError) IsNotValidValue() bool {
	return e.height <= 0 || e.width <= 0
}

// New 함수는 에러 객체를 생성한다.
func New(text string, width float64, height float64, dividen float64) error {
	return &OperationError{text, width, height, dividen}
}
