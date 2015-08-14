package gokatas

import (
	"testing"
)

/**
 * An array and its derived slice have same elements
 */
func TestArraySlice0(t *testing.T) {
	a := [5]string{"a", "b", "c", "d", "e"}
	b := a[2:]
	if !compareSliceAndSubArray5(b, 2, a) {
		t.Fail()
	}
}

/**
 * An array literal and its derived slice have same elements
 */
func TestArraySlice1(t *testing.T) {
	a := [...]string{"a", "b", "c", "d", "e"}
	b := a[2:]
	if !compareSliceAndSubArray5(b, 2, a) {
		t.Fail()
	}
}

/**
 * Setting an array element affects the derived slice
 */
func TestArraySlice2(t *testing.T) {
	a := [5]string{"a", "b", "c", "d", "e"}
	b := a[2:]
	a[4] = "z"
	if !compareSliceAndSubArray5(b, 2, a) {
		t.Fail()
	}
}

/**
 * Setting a slice element affects the underlying array but not a derived array
 */
func TestArraySlice3(t *testing.T) {
	a := [5]string{"a", "b", "c", "d", "e"}
	c := a
	b := c[2:]
	a[4] = "z"
	if compareSliceAndSubArray5(b, 2, a) {
		t.Fail()
	}
}

/**
 * Setting a slice element affects the underlying array
 */
func TestArraySlice4(t *testing.T) {
	a := [5]string{"a", "b", "c", "d", "e"}
	b := a[2:]
	b[2] = "z"
	if !compareSliceAndSubArray5(b, 2, a) {
		t.Fail()
	}
}

/**
 * Setting a slice element affects the underlying array and all derived slices
 */
func TestArraySlice5(t *testing.T) {
	a := [5]string{"a", "b", "c", "d", "e"}
	b := a[2:]
	c := b
	b[2] = "z"
	if !compareSliceAndSubArray5(c, 2, a) {
		t.Fail()
	}
}

/**
 * When a slice is expanded it "looses" contact with its original underlying array
 */
func TestArraySlice6(t *testing.T) {
	a := [5]string{"a", "b", "c", "d", "e"}
	b := a[2:]
	b = append(b, "x", "x", "x", "x", "x", "x")
	c := b[0:3]
	b[2] = "z"
	if compareSliceAndSubArray5(c, 2, a) {
		t.Fail()
	}
}

/**
 * Arrays are copied when passed as arguments...
 */
func TestArraySlice7(t *testing.T) {
	a := [5]string{"a", "b", "c", "d", "e"}
	b := modifyArray0(a)
	if a == b {
		t.Fail()
	}
}

func modifyArray0(a [5]string) [5]string {
	a[4] = "z"
	return a
}

/**
 * ...Until they are passed by reference
 */
func TestArraySlice8(t *testing.T) {
	a := [5]string{"a", "b", "c", "d", "e"}
	b := modifyArray1(&a)
	if a != b {
		t.Fail()
	}
}

func modifyArray1(a *[5]string) [5]string {
	a[4] = "z"
	return *a
}

func TestArraySlice9(t *testing.T) {
	a := [5]string{"a", "b", "c", "d", "e"}
	b := modifyArray2(&a)
	if &a != b {
		t.Fail()
	}
}

func modifyArray2(a *[5]string) *[5]string {
	a[4] = "z"
	return a
}

/**
 * When passed as a slice, the underlying array is modified...
 */
func TestArraySlice10(t *testing.T) {
	a := [5]string{"a", "b", "c", "d", "e"}
	b := modifyArray3(a[:])
	if !compareSliceAndSubArray5(b, 0, a) {
		t.Fail()
	}
}

func modifyArray3(a []string) []string {
	a[4] = "z"
	return a
}

/**
 * ... Until the slice is expanded during an append.
 */
func TestArraySlice11(t *testing.T) {
	a := [5]string{"a", "b", "c", "d", "e"}
	b := modifyArray4(a[:])
	if compareSliceAndSubArray5(b[:5], 0, a) {
		t.Fail()
	}
}

func modifyArray4(a []string) []string {
	a = append(a, a...)
	a[4] = "z"
	return a
}

func compareSliceAndSubArray5(ss []string, start int, a [5]string) bool {
	if len(ss) == len(a)-start {
		for i, s := range ss {
			if s != a[i+start] {
				return false
			}
		}
		return true
	}
	return false
}
