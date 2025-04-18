package smErrors

/*
VectorLike
Description:

	An interface for all objects that can be treated as vectors.
*/
type VectorLike interface {
	Len() int
	Dims() []int
}
