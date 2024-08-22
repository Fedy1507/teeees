package main

import (
	"fmt"
	"testing"
)

func TestSayHello(t *testing.T) {
	msg := sayHello("Ali")
	fmt.Println(msg)

	if msg != "Hello Ali" {
		t.Errorf("Expected 'Hello Ali', got '%s'", msg)
	}
}
