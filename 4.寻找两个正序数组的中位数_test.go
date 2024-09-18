package sf

import (
	"fmt"
	"testing"
)

func TestfindMedianSortedArrays4(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	k := findMedianSortedArrays4(nums1, nums2)
	fmt.Println(k)
}
