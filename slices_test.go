package gonerics_test

import (
	"testing"

	. "github.com/abenz1267/gonerics"
	"github.com/stretchr/testify/assert"
)

func TestCombine(t *testing.T) {
	t.Parallel()

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := Combine(arr[:4], Slice(arr[4]), arr[5:])

	assert.Equal(t, arr, res)
}

var res []int //nolint

func BenchmarkCombine(b *testing.B) {
	b.ReportAllocs()

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	var r []int

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		r = Combine(arr[:4], Slice(arr[4]), arr[5:])
	}

	res = r
}

func TestSlice(t *testing.T) {
	t.Parallel()

	arr := []int{1, 2, 3, 4}

	r := Slice(1, 2, 3, 4)

	assert.Equal(t, arr, r)
}

func BenchmarkSlice(b *testing.B) {
	b.ReportAllocs()

	var r []int

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		r = Slice(1, 2, 3, 4)
	}

	res = r
}
