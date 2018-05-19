package Matrix

type Matrix interface {
	GetM() int
	GetN() int
	GetElement(int, int) (interface{}, error)
	SetElement(int, int, interface{}) error
}
