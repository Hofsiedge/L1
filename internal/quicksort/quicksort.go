package quicksort

// Comparator(a, b) returns 1 if (a, b) is the desired order,
// 0 if a == b, -1 if (b, a) is the desired order.
type Comparator[T any] func(T, T) int

// partition splits slice into two parts: one is left to pivot, another is right to pivot
//
// partition returns the position of pivot point. pivot point is in the right half.
func partition[T any](slice []T, cmp Comparator[T], L int, R int) int {
	if R-L <= 1 {
		return R
	}
	pivot := slice[(L+R)/2]
	l, r := L, R
	for {
		// searching for an out-of-order element "left" to pivot
		for l <= R && cmp(slice[l], pivot) == 1 {
			l++
		}
		// searching for an out-of-order element "right" to pivot
		for r >= L && cmp(pivot, slice[r]) == 1 {
			r--
		}
		// quit if not found
		if l >= r {
			return r
		}
		slice[l], slice[r] = slice[r], slice[l]
	}
}

func quicksort[T any](slice []T, cmp Comparator[T], L int, R int) {
	if L >= 0 && R < len(slice) && L < R {
		p := partition(slice, cmp, L, R)
		quicksort(slice, cmp, L, p-1)
		quicksort(slice, cmp, p, R)
	}
}

func Quicksort[T any](slice []T, cmp Comparator[T]) {
	quicksort(slice, cmp, 0, len(slice)-1)
}
