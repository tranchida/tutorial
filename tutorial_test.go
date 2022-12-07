package tutorial

import (
	"fmt"
	"testing"
)

func TestComposite(t *testing.T) {

	message := Composite()

	fmt.Println(message)

	if len(message) == 0 {
		t.Fatal("can not be empty")
	}
}
