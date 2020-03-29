/*
** Merge Sort (Chapter 2)
 */

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type vector []int

func (v vector) String() string {

	var s []string
	for _, token := range v {
		s = append(s, fmt.Sprintf("%d", token))
	}
	return strings.Join(s, ", ")
}

func (v vector) Divide() (vector, vector) {
	var half = len(v) / 2
	return v[:half], v[half:]
}

func (v vector) MergeSort() vector {
	/*
	** Recursively split the vector into left and right pairs until we can split no more
	** then merge the pairs
	**
	 */

	if len(v) > 1 {
		var l, r = v.Divide()
		l = l.MergeSort()
		r = r.MergeSort()
		v = l.Merge(r)
	}
	return v
}

func (v vector) Merge(w vector) vector {
	/*
	** Merges two (sorted) vectors together and returns a single
	** ordered vector so [4, 6, 7, 9].Merge([1, 5, 8]) would return
	** [1, 4, 5, 6, 7, 8, 9]
	 */
	var r vector
	for len(v) != 0 || len(w) != 0 {
		if len(v) == 0 {
			r = append(r, w[0])
			w = w[1:]
		} else if len(w) == 0 {
			r = append(r, v[0])
			v = v[1:]
		} else if v[0] < w[0] {
			r = append(r, v[0])
			v = v[1:]
		} else {
			r = append(r, w[0])
			w = w[1:]
		}
	}
	return r
}

func (v vector) Check() error {
	/*
	** Checks that the vector is ordered. Returns nil if that is the case.
	 */

	for i := 0; i < len(v)-1; i++ {
		if v[i+1] < v[i] {
			return fmt.Errorf("Elements %d (%d) and %d (%d) are out of sequence", i, v[i], i+1, v[i+1])
		}
	}
	return nil
}

func (v *vector) Initialize(n int, m int) {
	/*
	** Initialises a vector with n random integers in the range [0 .. m)
	 */
	var i int

	for i = 0; i < n; i++ {
		*v = append(*v, rand.Intn(m))
	}
}
func main() {

	var e int // Exponent
	var c = 1
	for e = 1; e <= 6; e++ {
		var start = time.Now()
		var v = new(vector)
		c *= 10
		v.Initialize(c, 100)
		var w = v.MergeSort()
		var err error

		if err = w.Check(); err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("Sorted 10**%d random integers in %v\n", e, time.Since(start).Round(time.Millisecond))
		}
	}

}
