package main

import "fmt"

func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func Must2[T any, U any](value T, value2 U, err error) (T, U) {
	if err != nil {
		panic(err)
	}
	return value, value2
}

func However[T any](value T, err error) T {
	if err != nil {
		fmt.Println(err)
	}
	return value
}
