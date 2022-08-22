package here_test

import (
	"fmt"

	"github.com/0x416e746f6e/here"
)

func ExampleDoc() {
	lorem := here.Doc(`
		    Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
		eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad
		minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip
		ex ea commodo consequat.

		    Duis aute irure dolor in reprehenderit in voluptate velit esse
		cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat
		cupidatat non proident, sunt in culpa qui officia deserunt mollit anim
		id est laborum.
	`)

	fmt.Print(lorem)
	// Output:
	//     Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
	// eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad
	// minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip
	// ex ea commodo consequat.
	//
	//     Duis aute irure dolor in reprehenderit in voluptate velit esse
	// cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat
	// cupidatat non proident, sunt in culpa qui officia deserunt mollit anim
	// id est laborum.
	//
}
