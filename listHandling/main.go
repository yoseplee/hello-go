package main

import (
	"fmt"
)

type Man struct {
	Name string
	Age  int
}

func PrintIntArray(array *[]int) {
	for _, v := range *array {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func CartesianProduct(array *[]int, toProduct int) {
	for idx, val := range *array {
		fmt.Printf("idx: %d | val: %d", idx, val)
		(*array)[idx] = val * toProduct
		fmt.Println()
	}
}

func main() {
	fmt.Println("hello. practicing array, slice and map")

	// array
	fmt.Println("====array====")
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	arr2 := [5]int{10, 20, 30, 40, 50}

	// slice
	fmt.Println("====slice====")
	var slicedArray2 []int = arr2[:]
	PrintIntArray(&slicedArray2)
	CartesianProduct(&slicedArray2, 2)
	PrintIntArray(&slicedArray2)

	var slicedAll []int = arr2[3:5]
	slicedAll = append(slicedAll, 99, 999, 9999)
	PrintIntArray(&slicedAll)

	fmt.Println("====original arr====")
	tmpArr := arr2[:]
	PrintIntArray(&tmpArr)

	// 구조체 슬라이스 리터럴
	fmt.Println("====Struct Slice Literal====")
	structSliceLiteral := []struct {
		i int
		b bool
	}{
		{2, true},
		{5, true},
		{4, true},
		{3, true},
		{2, true},
		{2, true},
	}
	for _, v := range structSliceLiteral {
		fmt.Println(v.i)
	}

	//동적
	fmt.Println("dynamically allocated slice")
	makeSlice := make([]int, 5)
	var makeSlice2 []int = make([]int, 1, 5)

	// slice에 append할 때 반드시 len과 cap이 일치하지는 않는다
	// 일정한 범위로 cap이 증가함. 5->10->12->14
	makeSlice = append(makeSlice, 3, 3, 3, 3, 3, 3, 3, 3)

	fmt.Println(len(makeSlice), " || ", cap(makeSlice))
	for _, val := range makeSlice {
		fmt.Println(val)
	}

	fmt.Println(len(makeSlice2), " || ", cap(makeSlice2))
	for _, val := range makeSlice2 {
		fmt.Println(val)
	}

	// is ref? or value?
	var original []int = make([]int, 5)
	var assigned []int = original[:]
	assigned[3] = 33

	for idx, val := range original {
		fmt.Printf("original %d: %d || assigned %d: %d\n", idx, val, idx, assigned[idx])
	}
	fmt.Println("COPY")

	var copied []int = make([]int, 5)
	copy(copied, original)
	copied[4] = 44
	for idx, val := range original {
		fmt.Printf("original %d: %d || copied %d: %d\n", idx, val, idx, copied[idx])
	}

	// map
	fmt.Println("====Map====")
	var myMap map[string]int
	myMap = make(map[string]int)
	myMap["first"] = 111
	myMap["second"] = 222
	fmt.Println(myMap["first"])

	var myMap2 = map[string]Man{
		"head": Man{
			"lee", 36,
		},
		"col": Man{
			"kim", 22,
		},
	}
	fmt.Println(myMap2)

	fmt.Println("===map handling===")
	//add
	myMap["third"] = 333
	fmt.Println(myMap)

	//get
	var tmpInt int = myMap["first"]
	fmt.Println(tmpInt)

	//delete
	delete(myMap, "third")
	fmt.Println(myMap)

	//check
	tmpInt, ok := myMap["fifth"]
	if !ok {
		fmt.Println(ok)
	}
}
