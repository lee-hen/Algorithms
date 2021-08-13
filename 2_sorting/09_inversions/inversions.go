package inversions

// 2.2.19 Inversions. Develop and implement a linearithmic algorithm for computing the number of inversions in a given array
// (the number of exchanges that would be performed by insertion sort for that array—see SECTION 2.1).
// This quantity is related to the Kendall tau distance; see SECTION 2.5.


type intSlice []int

func (s intSlice) less (i, j int) bool {
	return s[i] < s[j]
}

func merge(a, aux intSlice, lo, mid, hi int) int {
	// copy to aux[]
	for k := lo; k <= hi; k++ {
		aux[k] = a[k]
	}

	// merge back to a[]
	i, j := lo, mid+1
	var inversions int
	for k := lo; k <= hi; k++ {
		if i > mid {
			a[k] = aux[j]
			j++
		} else if j > hi {
			a[k] = aux[i]
			i++
		} else if aux.less(j, i) {
			a[k] = aux[j]
			j++

			inversions += mid-i+1
		} else {
			a[k] = aux[i]
			i++
		}
	}
	return inversions
}

func count(a, aux intSlice, lo, hi int) int {
	if hi <= lo {
		return 0
	}
	var inversions int

	mid := lo + (hi-lo)/2
	inversions += count(a, aux, lo, mid)
	inversions += count(a, aux, mid+1, hi)
	inversions += merge(a, aux, lo, mid, hi)

	return inversions
}

func Count(a []int) int {
	aux := make(intSlice, len(a), len(a))
	return count(a, aux, 0, len(a)-1)
}
