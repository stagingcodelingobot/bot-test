package main

import (
	"fmt"
)


func Export1(msg string) error {
	fmt.Println("Exported")
	fmt.Println(msg)
	return nil
}

func GetSubSlice1(start int, slice []int) []int {
	return slice[start:len(slice)]
}

func AnotherGetSubSlice1(start int, slice []int) []int {
	// a comment
	return nil
}

func Do1() error {
	err := Export("Hello")
	if err != nil {
		return err
	}
	return nil
}

func Another1() {
	fmt.Println("Another")
}
