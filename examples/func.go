package main

// add
// @desc
// @param a int
// @param b int
// @return d int
// @return e int
// @return c error
func add(a int, b int) (d int, e int, c error) {
	return a + b, a, nil
}

type myInt int

// add
// @desc
// @param a int
// @param b int
// @return d int
// @return e int
// @return c error
func (m myInt) add(a int, b int) (d int, e int, c error) {
	return a + b, a, nil
}

type MyIntface interface {
	// add
	// @desc
	// @param a int
	// @param b int
	// @return d int
	// @return e int
	// @return c error
	add(a int, b int) (d int, e int, c error)
}
