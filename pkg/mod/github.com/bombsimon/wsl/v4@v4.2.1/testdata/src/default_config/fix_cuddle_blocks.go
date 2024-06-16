package testpkg

import "fmt"

func FailBlocks() {
	if false {
		return
	}
	_, err := maybeErr() // want "assignments should only be cuddled with other assignments"
	if err != nil {      // want "only one cuddle assignment allowed before if statement"
		return
	}

	if true {
		return
	}
	_, x := maybeErr() // want "assignments should only be cuddled with other assignments"
	if err != nil {    // want "only one cuddle assignment allowed before if statement"
		return
	}

	x = nil
	_, err = maybeErr()
	if err != nil { // want "only one cuddle assignment allowed before if statement"
		panic(err)
	}

	if x != nil {
	}
	switch 1 { // want "anonymous switch statements should never be cuddled"
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	foo := 1
	bar := 2
	baz := 3
	if 1 == 1 { // want "only one cuddle assignment allowed before if statement"
		baz = 4
	}

	_ = foo
	_ = bar
	_ = baz

	a := 1
	b := 2
	for b < bar { // want "only one cuddle assignment allowed before for statement"
		b--
	}

	_ = a
	_ = b
}

func maybeErr() (int, error) {
	return 0, nil
}
