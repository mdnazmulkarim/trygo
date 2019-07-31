package main

import (
		"fmt"
		"errors"
		"math"
		)
		
		
type person struct {
	name string
	age int
}

func main() {
	fmt.Printf("hello, world\n")
	
	var x int
	var z int = 5
	y := 99
	
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	
	
	//array
	
	var a [5]int
	a[2] = 8
	//a[6] = 9
	fmt.Println(a)
	
	b := [5]int{2,5,7,4,2}
	//b = append(b,6)
	fmt.Println(b)
	
	
	//slice
	c := []int{1,4,2,9,3}
	c = append(c,10)
	
	
	fmt.Println(c)
	
	
	//map
	
	vertices := make(map[string]int)
	
	vertices["circle"] = 1
	vertices["triangle"] = 2
	vertices["rectangle"] = 3
	
	fmt.Println(vertices)
	
	delete(vertices, "circle")
	
	fmt.Println("Triangle:",vertices["triangle"])
	
	fmt.Println("Map after delete:", vertices)
	
	
	//for loop:
	
	for i:= 0; i<5; i++ {
		fmt.Println(i)
	}
	
	//while loop:
	
	fmt.Println("While loop:")
	
	j := 0
	
	for j<5 {
		fmt.Println(j)
		j++
	}
	
	arr := []string{"a","b","c"}
	
	for index, value := range arr {
		fmt.Println("Index:",index, "Value:", value)
	}
	
	
	for key, value := range vertices {
		fmt.Println("Key:",key, "Value:", value)
	}
	
	fmt.Println(sum(3,5))
	
	result , err := sqrt(-9)
	
	if  err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	
	// struct
	
	p :=  person{name: "nazmul", age: 23}
	fmt.Println(p)
	fmt.Println(p.name)
	
	p.name = "karim"
	fmt.Println(p)
	
	
	//pointer
	
	k := 7
	inc(&k)
	fmt.Println(k)
	
}

func inc(x *int){
	*x++
}

func sum(x int, y int) int {
	return x +y
}

func sqrt(x float64) (float64, error) {
    if x < 0 {
		return 0, errors.New("Undefined")
	} else {
		return math.Sqrt(x), nil
	}
}