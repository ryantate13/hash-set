package hash_set

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	tests := []struct{ it string }{
		{"instantiates a new empty set"},
	}
	for _, tt := range tests {
		t.Run(tt.it, func(t *testing.T) {
			s := New[int]()
			require.NotNil(t, s)
			require.True(t, s.Empty())
		})
	}
}

func TestOf(t *testing.T) {
	tests := []struct{ it string }{
		{"instantiates a set with provided elements"},
	}
	for _, tt := range tests {
		t.Run(tt.it, func(t *testing.T) {
			elements := make([]int, 0)
			for i := 0; i < 10; i++ {
				for j := 0; j <= i; j++ {
					elements = append(elements, i)
				}
			}
			s := Of(elements...)
			for i := 0; i < 10; i++ {
				assert.True(t, s.Has(i))
			}
			assert.Equal(t, 10, s.Len())
		})
	}
}

func TestSet_Empty(t *testing.T) {
	tests := []struct {
		it     string
		s      *Set[int]
		expect bool
	}{
		{
			it:     "returns true if a set is empty",
			s:      New[int](),
			expect: true,
		},
		{
			it:     "returns false a set is not empty",
			s:      Of(1),
			expect: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.it, func(t *testing.T) {
			require.Equal(t, tt.expect, tt.s.Empty())
		})
	}
}

func TestSet_Has(t *testing.T) {
	tests := []struct {
		it     string
		s      *Set[int]
		el     int
		expect bool
	}{
		{
			it:     "returns true if an element is in a set",
			s:      Of(1),
			el:     1,
			expect: true,
		},
		{
			it:     "returns false if an element is not in a set",
			s:      New[int](),
			el:     0,
			expect: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.it, func(t *testing.T) {
			require.Equal(t, tt.expect, tt.s.Has(tt.el))
		})
	}
}

func TestSet_Subset(t *testing.T) {
	tests := []struct {
		it     string
		s1     *Set[int]
		s2     *Set[int]
		expect bool
	}{
		{
			it:     "returns true if the set is a subset of another set",
			s1:     Of(1),
			s2:     Of(1, 2, 3),
			expect: true,
		},
		{
			it:     "the empty set is a subset of every other set",
			s1:     New[int](),
			s2:     Of(1, 2, 3),
			expect: true,
		},
		{
			it:     "the empty set is a subset of itself",
			s1:     New[int](),
			s2:     New[int](),
			expect: true,
		},
		{
			it:     "returns false if the set is not a subset of another set",
			s1:     Of(1, 2, 3),
			s2:     Of(2, 3, 4),
			expect: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.it, func(t *testing.T) {
			assert.Equal(t, tt.expect, tt.s1.Subset(tt.s2))
		})
	}
}

func TestSet_Add(t *testing.T) {
	tests := []struct{ it string }{
		{"adds an element to the set, optionally supports chaining"},
	}
	for _, tt := range tests {
		t.Run(tt.it, func(t *testing.T) {
			s := New[int]()
			require.False(t, s.Has(1))
			s.Add(1)
			require.True(t, s.Has(1))
			s.Add(2).Add(3).Add(4)
			for i := 2; i < 5; i++ {
				require.True(t, s.Has(i))
			}
			require.Equal(t, 4, s.Len())
		})
	}
}

func TestSet_Remove(t *testing.T) {
	tests := []struct{ it string }{
		{"removes an element from the set, optionally supports chaining"},
	}
	for _, tt := range tests {
		t.Run(tt.it, func(t *testing.T) {
			s := Of(1)
			require.True(t, s.Has(1))
			s.Remove(1)
			require.False(t, s.Has(1))
			s.Add(1, 2, 3)
			for i := 1; i <= 3; i++ {
				require.True(t, s.Has(i))
			}
			s.Remove(1).Remove(2, 3)
			for i := 1; i <= 3; i++ {
				require.False(t, s.Has(i))
			}
		})
	}
}

func TestSet_Len(t *testing.T) {
	tests := []struct{ it string }{
		{"returns the number of elements in the set"},
	}
	for _, tt := range tests {
		t.Run(tt.it, func(t *testing.T) {
			require.Equal(t, 0, New[int]().Len())
			require.Equal(t, 1, New[int]().Add(1).Len())
			require.Equal(t, 1, Of(1).Len())
			require.Equal(t, 10, Of(0).Add(1).Add(2).Add(3, 4, 5, 6, 7, 8).Add(9).Len())
		})
	}
}

func TestSet_Slice(t *testing.T) {
	tests := []struct{ it string }{
		{"returns the elements of the set as a slice"},
	}
	for _, tt := range tests {
		t.Run(tt.it, func(t *testing.T) {
			s := []int{1, 2, 3, 4, 5}
			got := Of(s...).Add(s...).Slice()
			sort.Ints(got)
			require.Equal(t, s, got)
		})
	}
}

func TestSet_Foreach(t *testing.T) {
	tests := []struct{ it string }{
		{"executes a callback function with every member of the set"},
	}
	for _, tt := range tests {
		t.Run(tt.it, func(t *testing.T) {
			sum := 0
			Of(1, 2, 3, 4, 4, 3, 2, 1).Foreach(func(i int) { sum += i })
			require.Equal(t, 10, sum)
		})
	}
}

func TestSet_Filter(t *testing.T) {
	tests := []struct{ it string }{
		{"returns the subset of a set that pass a given a filter func"},
	}
	for _, tt := range tests {
		t.Run(tt.it, func(t *testing.T) {
			evens := Of(1, 2, 3, 4, 4, 3, 2, 1).Filter(func(i int) bool { return i%2 == 0 })
			require.Equal(t, Of(2, 4), evens)
		})
	}
}

func TestSet_Intersection(t *testing.T) {
	tests := []struct{ it string }{
		{"returns the set of elements common to both sets"},
	}
	for _, tt := range tests {
		t.Run(tt.it, func(t *testing.T) {
			s1 := Of(1, 2, 3, 4, 5, 6, 7, 8)
			s2 := Of(5, 6, 7, 8, 9, 10, 11, 12)
			i := Of(5, 6, 7, 8)
			require.Equal(t, i, s1.Intersection(s2))
			require.Equal(t, i, s2.Intersection(s1))
			require.Equal(t, s1, s1.Intersection(s1))
			require.Equal(t, s2, s2.Intersection(s2))
		})
	}
}

func TestSet_Union(t *testing.T) {
	tests := []struct{ it string }{
		{"returns the combined elements of both sets"},
	}
	for _, tt := range tests {
		t.Run(tt.it, func(t *testing.T) {
			s1 := Of(1, 2, 3, 4, 5, 6, 7, 8)
			s2 := Of(5, 6, 7, 8, 9, 10, 11, 12)
			u := Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
			require.Equal(t, u, s1.Union(s2))
			require.Equal(t, u, s2.Union(s1))
			require.Equal(t, s1, s1.Union(s1))
			require.Equal(t, s2, s2.Union(s2))
		})
	}
}

func TestSet_Difference(t *testing.T) {
	tests := []struct{ it string }{
		{"returns the set of elements not in the other set"},
	}
	for _, tt := range tests {
		t.Run(tt.it, func(t *testing.T) {
			s1 := Of(1, 2, 3, 4, 5, 6, 7, 8)
			s2 := Of(5, 6, 7, 8, 9, 10, 11, 12)
			d1 := Of(1, 2, 3, 4)
			d2 := Of(9, 10, 11, 12)
			require.Equal(t, d1, s1.Difference(s2))
			require.Equal(t, d2, s2.Difference(s1))
			require.Equal(t, Of[int](), s1.Difference(s1))
			require.Equal(t, Of[int](), s2.Difference(s2))
		})
	}
}
