package main

/*

func main() {

	// Array - 固定長度
	var arrayDemo [2]int
	arrayDemo[0] = 1
	arrayDemo[1] = 2
	fmt.Println(arrayDemo)

	otherArrayDemo := [3]string{"cxcxc","aws-saa","aws-da"}
	fmt.Println(otherArrayDemo)

	// Slice - 動態長度
	//	Slice: The size specifies the length. The capacity of the slice is
	//	equal to its length. A second integer argument may be provided to
	//	specify a different capacity; it must be no smaller than the
	//	length. For example, make([]int, 0, 10) allocates an underlying array
	//	of size 10 and returns a slice of length 0 and capacity 10 that is
	//	backed by this underlying array.
	var stringSlice []string
	stringSlice = make([]string, 5, 5)
	stringSlice = append(stringSlice,"123")
	fmt.Println(stringSlice)


	// map
	// create a map without assign any value
	mapDemo := make(map[string]int)

	// assign value
	mapDemo["k1"] = 7
	mapDemo["k2"] = 13
	fmt.Println("len:", len(mapDemo))

	// delete value
	delete(mapDemo, "k2")

	// get value 與確認是否存在
	mapValue, doesExist := mapDemo["k1"]
	fmt.Println("prs:", mapValue, doesExist)

	mapValue2, doesExist2 := mapDemo["k2"]
	fmt.Println("prs:", mapValue2, doesExist2)

	// 直接宣告完整的map
	directDeclareMap := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(directDeclareMap)

}

 */
