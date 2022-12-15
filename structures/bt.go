package structures

type BT[T any] struct {
	Data  T
	Left  *BT[T]
	Right *BT[T]
}
