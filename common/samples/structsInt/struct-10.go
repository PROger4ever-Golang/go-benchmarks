package structsInt

type Struct10 struct {
	Field01, Field02, Field03, Field04, Field05, Field06, Field07, Field08, Field09 int
	Field10                                                                int
}

func NewSampleStruct10() *Struct10 {
	return &Struct10{
		Field01: 1,
		Field02: 2,
		Field03: 3,
		Field04: 4,
		Field05: 5,
		Field06: 6,
		Field07: 7,
		Field08: 8,
		Field09: 9,
		Field10: 10,
	}
}