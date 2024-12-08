package text_test

import (
	"fmt"

	"github.com/poly-gun/text"
)

func ExampleLowercase() {
	v := "Field-Name"

	fmt.Println(text.Lowercase(v))
	// Output: field-name
}

func ExampleTitle() {
	// v represents a name that should be otherwise capitalized (titled)
	v := "jacob b. sanders"

	fmt.Println(text.Title(v))
	// Output: Jacob B. Sanders
}

func ExampleVariadic() {
	v := text.Dereference(nil, func(o *text.Options) {
		o.Log = true
	})

	pointer := text.Pointer("example")
	v = text.Dereference(pointer, func(o *text.Options) {
		o.Log = true
	})

	fmt.Println(v)
	// Output: example
}

func ExampleDereference() {
	// initialize a string pointer of underlying value "example"
	pointer := text.Pointer("example")

	// establish variable "v" of type string
	v := text.Dereference(pointer, func(o *text.Options) {
		o.Log = true // log if the pointer is nil
	})

	fmt.Println(v)

	// establish variable "v" of type string, without using functional options.
	v = text.Dereference(pointer)

	fmt.Println(v)
	// Output: example
	// Output: example
}

func ExamplePointer() {
	// create a pointer of type string with reference value: "example"
	pointer := text.Pointer("example", func(o *text.Options) {
		o.Log = true // log if the string value is an empty string
	})

	fmt.Println(*(pointer))

	// create a pointer of type string with reference value: "example", without using functional options.
	pointer = text.Pointer("example")

	fmt.Println(*(pointer))
	// Output: example
	// Output: example
}
