package main

/*小和问题
在一个数组中，每一个数左边比当前数小的数累加起来，叫做这个数组的小和。
求一个数组的小和。
例子:[1,3,4,2,5] 1左边比1小的数，没有;  3左边比3小的数，1; 4左边比4小的数，1、3;
2左边比2小的数，1; 5左边比5小的数，1、3、4、2; 所以小和为1+1+3+1+1+3+4+2=16
*/

func SmallSum(nums []int) int {
	return process(nums, 0, len(nums)-1)
}

func process(nums []int, left int, right int) int {
	if left == right {
		return 0
	}
	mid := left + (right-left)>>1
	return process(nums, left, mid) + process(nums, mid+1, right) + Merge(nums, left, mid, right)
}

func Merge(nums []int, left int, mid int, right int) int {
	tmp := make([]int, right-left+1)
	p1, p2 := left, mid+1
	i := 0
	res := 0
	for p1 <= mid && p2 <= right {
		if nums[p1] < nums[p2] {
			res += (right - p2 + 1) * nums[p1]
			p1++
			i++
		} else {
			res += 0
			p2++
			i++
		}
	}
	//右边数组分治结束，左边还未
	for p1 <= mid {
		tmp[i] = arr[p1]
		p1++
		i++
	}
	//左边数组分治结束，右边还未
	for p2 <= right {
		tmp[i] = arr[p2]
		p2++
		i++
	}
	//分治结束，替换原值
	for i := 0; i < len(tmp); i++ {
		arr[left+i] = tmp[i]
	}
	return res

}
