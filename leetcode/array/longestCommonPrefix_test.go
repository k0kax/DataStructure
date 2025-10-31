package array

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{"flower", "flow", "floight"}
	prefix := longestCommonPrefix2(strs)
	fmt.Println(prefix)
}
