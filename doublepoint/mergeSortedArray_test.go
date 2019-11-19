package main

import (
	"reflect"
	"testing"
)

func Testmerge(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	if !reflect.DeepEqual(nums1, []int{1, 2, 2, 3, 5, 6}) {
		t.Errorf(`merge result %v`, nums1)
	}
}
