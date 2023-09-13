package src

import (
	"sort"
)

func FirstPositiveInteger(A []int) int {
	// assume output is 1
	output := 1

	if A == nil {
		return output
	}

	// sort the slice
	sort.Ints(A)

	// iterate through the slice, A
	for i := 0; i < len(A); i++ {
		if A[i] < 0 {
			// negative integer, disregard.
			continue
		}

		if A[i] > output {
			// already found our lowest possible positive integer
			// because of gap in the sorted slice.
			break
		}

		if A[i] == output {
			output += 1
		}
	}

	return output
}
