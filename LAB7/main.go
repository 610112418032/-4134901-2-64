package main

import (
	"fmt"

	"pengssk/LAB7/structs"
)

func main ()  {
	const rname = "Suksan"
	var age = 20

	fmt.Printf("Hello...%s : %v \n", rname,age)

	if age > 20 {
		fmt.Println(" > 20")
	} else{
		fmt.Println("< 20")
	}

	switch age {
	case 20:
		fmt.Println("age is 20")
	case 25:
		fmt.Println("age is 20")
	}

	var count = 0
	for count < 10{
		fmt.Printf("%v",count)
		count++
	}

	result := add(10,20)
	fmt.Printf("\n%v",result)


	var ptr *int
	fmt.Printf("\n%v",&ptr)

	city := "Buriram"
	var cityPointer *string
	cityPointer = &city
	fmt.Println(*cityPointer)

	structs.Learn()
}

func add(a int,b int) int  {
	c := a+b
	return c
}