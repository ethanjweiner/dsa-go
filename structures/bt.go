package structures

import (
	"golang.org/x/exp/constraints"
)

type BT[T constraints.Ordered] struct {
	Data  T
	Left  *BT[T]
	Right *BT[T]
}
