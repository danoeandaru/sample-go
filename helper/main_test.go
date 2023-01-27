package helper

import (
	"fmt"
	"testing"
)

// hanya run di satu package
// before and after testing
// m.Run, menjalankan semua unit testing dalam package

func TestMain(m *testing.M) {
	fmt.Println("before unit testing")
	m.Run()
	fmt.Println("after unit testing")
}
