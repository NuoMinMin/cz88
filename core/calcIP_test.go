package core

import (
	"fmt"
	"testing"
)

func TestCalcIP(t *testing.T) {
	res := calcIP("255.255.255.255")
	fmt.Println("res", res)
}
