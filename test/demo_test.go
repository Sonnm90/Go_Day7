package test

import (
	"fmt"
	"testing"
)

func hello(name string) string {
	if name == "" {
		return fmt.Sprintf("What is your name ?")
	} else {
		return fmt.Sprintf("Hello %s", name)
	}
}
func TestHello(t *testing.T) {
	emptyNameResult := hello("")

	if emptyNameResult != "What is your name ?" {
		t.Errorf("Output expect What is your name ? instead of %v", emptyNameResult)
	}

	result := hello("SOn")

	if result != "Hello Gopher" {
		t.Errorf("Output expect Hello Gopher instead of %v", result)
	}
}
