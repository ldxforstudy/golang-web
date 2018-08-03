package example

import (
	"testing"
	"fmt"
)

type driver1 struct {
}

func (p driver1) Info() string {
	return "Driver1"
}

func TestFinishRate_Query(t *testing.T) {
	f := FinishRate{driver: driver1{}}
	fmt.Println(f.Query())
}
