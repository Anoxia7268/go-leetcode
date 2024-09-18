package sf

import (
	"fmt"
	"testing"
)

func TestTwoSum1(t *testing.T) {
	nums := []int{2, 3, 0}
	target := 6
	k := twoSum1(nums, target)
	fmt.Println(k)
}
