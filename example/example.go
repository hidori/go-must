package main

import (
	"errors"
	"fmt"

	"github.com/hidori/go-must"
)

func main() {
	do()
	do2()
	do3()
}

func do() {
	v := must.Do(func() (int, error) {
		return 1, nil
	})
	fmt.Println(v)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	must.Do(func() (int, error) {
		panic(errors.New("error1"))
	})
}

func do2() {
	v, v1 := must.Do2(func() (int, int, error) {
		return 1, 2, nil
	})
	fmt.Println(v, v1)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	must.Do2(func() (int, int, error) {
		panic(errors.New("error2"))
	})
}

func do3() {
	v, v1, v2 := must.Do3(func() (int, int, int, error) {
		return 1, 2, 3, nil
	})

	fmt.Println(v, v1, v2)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	must.Do3(func() (int, int, int, error) {
		panic(errors.New("error3"))
	})
}
