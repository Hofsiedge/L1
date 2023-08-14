package bisect

type ComparatorOrder int

const (
	OrderedRight ComparatorOrder = iota
	OrderedWrong
	OrderedEqual
)

type Comparator[T any] func(T, T) ComparatorOrder

// bisect implements binary search over slice of T
//
// slice must be sorted, matching comparator logic
//
// comparator(a, b) tells what is the order of (a, b)
//
// bisect looks for value in the slice, returns (position, true) if found,
// (position, false) if not found (position is the index where value can be inserted)
func Bisect[T comparable](slice []T, comparator Comparator[T], value T) (int, bool) {
	L, R := 0, len(slice)-1
	for L <= R && L < len(slice) {
		M := (L + R) / 2
		switch comparator(value, slice[M]) {
		case OrderedRight:
			R = M - 1
		case OrderedWrong:
			L = M + 1
		case OrderedEqual:
			return M, true
		}
	}
	return L, L < len(slice) && slice[L] == value
}
