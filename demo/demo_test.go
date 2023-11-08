package demo

import (
	"fmt"
	"testing"
)

func TestDemo(t *testing.T) {
	if Demo() != "Your demo package!" {
		t.Errorf("Error demo :(")
	}
	fmt.Println("successful demo, good job!")
}
