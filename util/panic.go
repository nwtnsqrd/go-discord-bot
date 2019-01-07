package util

// PanicOnErr wraps the panic function
func PanicOnErr(err error) {

	if err != nil {
		panic(err)
	}
}
