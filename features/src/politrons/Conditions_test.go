package main

import (
	"fmt"
	"testing"
)
// Unit test for the for loop functionality
// Pretty much the same like in other languajes, quite extend and simple to use like in Scala
func ForLoop() {
	foo := 0
	for i := 0; i < 10; i++ {
		foo += i
	}
	fmt.Println(foo)

	for foo = 0; foo < 100; {
		foo = +1
	}
	fmt.Println(foo)

	for foo < 100 {
		foo = +1
	}
	fmt.Println(foo)
}

//Using [range] operator we can extract and interate elements from an array and get index and value.
//You can also do it to iterte over a map and extract pair of key/value per entry.
func TestForRange(t *testing.T) {
	var array = []int{1,4,30,10,5,68}
    for index,value := range array {
		println("Index" ,index) 
		println("Value" ,value)
	}
	myMap := map[int]string{1:"value1",2:"value2",3:"value3"}
	for key, value := range myMap{
		println("Key", key, "Value", value)
	}
}

/* 
[Defer] is an operator to make wait the execute of a function until the end of the current function where
remains ends. Once the function ends start executing the defer functions just like recursively from the bottom to 
the top of the function.
*/
func TestForDefer(t *testing.T) {
	fmt.Println("Start")

	defer println("Iteration done")
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
	defer println("Iteration start")

	fmt.Println("End")
}