package algo9

import (
	"fmt"
	"testing"
)

func TestLowerCase(t *testing.T) {
	s := "Hello 中文"
	res := toLowerCase(s)
	fmt.Println(res)
}
