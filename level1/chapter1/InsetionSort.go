package chapter1

func swap(arr []int, i int, j int) {
	//使用位运算的前提是数组的值不能是相同的，否则最后会为0
	// arr[i] = arr[i] ^ arr[j]
	// arr[j] = arr[i] ^ arr[j]
	// arr[i] = arr[i] ^ arr[j]

	var tmp int
	tmp = arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

//形参不需要带指针，因为切片是引用类型
func insertionSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	for i := 1; i < len(arr); i++ {
		for j := i; j >= 0; j-- {
			if arr[j-1] > arr[j] {
				swap(arr, j-1, j)
			}
		}
	}

}
