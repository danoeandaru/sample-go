package helper

import (
	"testing"
)

// nama file diakhiri _test
// nama func diawali Test
// go test -v (verbose)
// go test -v -run=TestHelloWorld, testing satu func
// go hanya testing di package yang ada di package tsb
// jika ingin run semua unit test, go test -v ./...
// don't use panic for throwing error, jadi stopper untuk semua func test
// use Fail, FailNow, Error, Fatal
func TestHelloWorld(t *testing.T) {

	result := HelloWorld("Danu")
	if result != "Hello World Danu" {
		// if fail, lanjut testing untuk code lainnya
		t.Fail()
	}
}

func TestHelloWorld2(t *testing.T) {

	result := HelloWorld("Danu")
	if result != "Hello World Danu" {
		// if fail, then stop
		t.FailNow()
	}
}

func TestHelloWorld3(t *testing.T) {

	result := HelloWorld("Danu")
	if result != "Hello World Danu" {
		// throw log error, then call t.Fail
		t.Error("Result must be Hello World Danu")
	}
}

func TestHelloWorld4(t *testing.T) {

	result := HelloWorld("Danu")
	if result != "Hello World Danu" {
		// throw log error, then call t.FailNow
		t.Fatal("Result must be Hello World Danu")
	}
}
