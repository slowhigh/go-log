package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

// go test 명령어는 "_test.go"라고 끝나는 테스트 파일의 "Test"라고 시작하는 function들을 실행한다.
// -v 옵션을 추가하면 각각의 테스트의 결과를 나열한다.

// $ go test
// PASS
// ok      example.com/greetings   0.364s

// $ go test -v
// === RUN   TestHelloName
// --- PASS: TestHelloName (0.00s)
// === RUN   TestHelloEmpty
// --- PASS: TestHelloEmpty (0.00s)
// PASS
// ok      example.com/greetings   0.372s