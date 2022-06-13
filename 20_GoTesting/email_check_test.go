package email

import "testing"

func TestGmail(t *testing.T) {
	gmail := "user@gmail.com"
	if !IsFreemail(gmail) {
		t.Log()
		t.Fail()
	}
}
