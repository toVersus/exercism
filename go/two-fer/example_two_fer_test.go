package twofer

import "fmt"

func ExampleShareWith() {
	fmt.Println(ShareWith(""))
	fmt.Println(ShareWith("Alice"))
	// Output:
	// One for you, one for me.
	// One for Alice, one for me.
}
