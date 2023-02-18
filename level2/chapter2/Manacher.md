

# 最长回文字符串

## 经典解

![image-20221112100728114](E:\linux学习PDF\linux\img\image-20221112100728114.png)

 ```go
 package main
 
 import (
 	"fmt"
 	"math"
 )
 
 // 最长回文字符串
 
 /*
 暴力思路：在每个字符前后都加入一个字符(可以避免奇偶数情况)，然后从左往右开始遍历字符串，
 		 每个字符以自己为中心，向两边扩，记录最长的相同字符串长度,最后除于2就是答案
 举例:
 # a # b # b # a # g # c # b # b # a #
 1 3 1 3 9 3 1 3 1 3 1 3 1 3 5 3 1 3 1
 
 9 / 2 = 4 (abba)
 */
 // O(N)
 func longpalindrome(str string) []int {
 
 	n := len(str)
 	tmp := make([]byte, n+n+1)
 	// 插入字符
 	for i := 0; i < n+n+1; i = i + 2 {
 		tmp[i] = '#'
 	}
 	for i, j := 1, 0; i < n+n+1; i, j = i+2, j+1 {
 		tmp[i] = str[j]
 	}
 	// fmt.Println(string(tmp))
 	p, q := 0, 0
 
 	ans := make([]int, n+n+1)
 	for i := 0; i < n+n+1; i++ {
 		t := i
 		p = t - 1
 		q = t + 1
 		cnt := 1
 		// 边界情况
 		if p < 0 || q > n+n {
 			ans[i] = cnt
 			continue
 		}
 		// fmt.Println("ASdsad")
 		for p >= 0 && q <= n+n {
 			if tmp[p] != tmp[q] {
 				ans[i] = cnt
 				break
 			} else {
 				cnt = cnt + 2
 				ans[i] = cnt
 			}
 			p--
 			q++
 		}
 		// 走else，然后p--,q++越界，cnt要重新初始
 	}
 	// fmt.Println(ans)
 	return ans
 }
 func Max(a, b int) int {
 	if a > b {
 		return a
 	}
 	return b
 }
 
 func main() {
 	st := "abaaxjsowlaabaa"
 	ans := longpalindrome(st)
 	max := math.MinInt64
 	for i := 0; i < len(ans); i++ {
 		max = Max(max, ans[i])
 	}
 	fmt.Println(max / 2)
 }
 
 ```





## Manache算法





### 示例

**初始化：** 到达回文最右边界`R`，回文中心`C`初始都令其为 -1 

<img src="E:\linux学习PDF\linux\img\image-20221113213535630.png" alt="image-20221113213535630" style="zoom:67%;" />



**开始遍历**，来到第0位，左右两边都不是回文，范围是[0,0], 更新R，C

<img src="E:\linux学习PDF\linux\img\image-20221113213751978.png" alt="image-20221113213751978" style="zoom:67%;" />

第1位，左右两边有回文，虽然是自填加的字符，但是也是满足，所以范围为[0,2]，R，C更新

<img src="E:\linux学习PDF\linux\img\image-20221113214017893.png" alt="image-20221113214017893" style="zoom:67%;" />

第2位，左右不回文，没有超过最大右边界，所以R，C不更新。

<img src="E:\linux学习PDF\linux\img\image-20221113214313617.png" alt="image-20221113214313617" style="zoom:67%;" />

第三位，左右回文，且范围是[0,6]，R是到达的最右有边界所以R更新为6。

<img src="E:\linux学习PDF\linux\img\image-20221113215218835.png" alt="image-20221113215218835" style="zoom:67%;" />

第四位，左右不回文，范围[4,4]，在R内不是，未到达的最右边界，所以不更新R，C

<img src="E:\linux学习PDF\linux\img\image-20221113215306179.png" alt="image-20221113215306179" style="zoom:67%;" />

第5第6位，和第四位一样都在范围内，所以都不更新



### 可能情况

**情况1 ：**不在R内，直接扩，更新R



**情况2：**分为3中小情况

当到`i`位置时，发现在R内部，此时C一定在`i`左侧，因为C是最右回文边界的中心点，包含`i`位置，所以在C左侧一定存在`i`的对称点`i'`,R也是，[L,R]这个区域内是回文字符串

<img src="E:\linux学习PDF\linux\img\image-20221113224403684.png" alt="image-20221113224403684" style="zoom:67%;" />



1） （这里特殊字符没写，当作还是有）`i'`的回文区域在[L,R]内部，那么此时 `i`的回文字符串长度和 `i'`长度一样

<img src="E:\linux学习PDF\linux\img\image-20221113225810087.png" alt="image-20221113225810087" style="zoom:67%;" />

==`i`的回文字符串至少是乙那么大，为啥不可能更大呢？==

甲当时为啥范围没有扩更大，因为`X != Y`，说明甲的回文子串就大，`i`关于C对称，可得`Y==Z，X==Z`，所以`Z != L` 

<img src="E:\linux学习PDF\linux\img\image-20221113232124832.png" alt="image-20221113232124832" style="zoom:57%;" />



2）`i'`的回文串在[L,R]外，那么此时`i`的回文子串就是 [R',R] ，因为在以C为中心点的区间[L,R]内，以`i'`为中心对称，L'  ，[L, L' ]必是回文，反之可以证明 [R',R]这个区间也是回文。

<img src="E:\linux学习PDF\linux\img\image-20221114102626383.png" alt="image-20221114102626383" style="zoom:77%;" />

示例：

（这里特殊字符没写，当作还是有）

<img src="E:\linux学习PDF\linux\img\image-20221114103407428.png" alt="image-20221114103407428" style="zoom:57%;" />

==那么`i`的回文区间有没有可能在R区间外呢？==

因 X和 Y在 `i'`的回文区间内，所以X == Y，又因 [L,R]是回文区间，Y == Z，

为啥[L,R]不能扩的更远呢？因为X != P，所以等量关系可得 Z != P

答案是： 不能

<img src="E:\linux学习PDF\linux\img\image-20221114103728585.png" alt="image-20221114103728585" style="zoom:67%;" />



3）`i'`的回文区间刚好压线L， 那么 `i`的回文区间**至少**是 [R',R]，因为C的回文区间是 [L,R] ，`i'`刚好左区间与C的重合，又因为对称关系，[L,I'] 等于 [i,R]，做`i`的对称R'， 所以 [R',R]也是回文区间。 

<img src="E:\linux学习PDF\linux\img\image-20221114105333512.png" alt="image-20221114105333512" style="zoom:67%;" />

**示例**

（这里特殊字符没写，当作还是有）

<img src="E:\linux学习PDF\linux\img\image-20221114104816613.png" alt="image-20221114104816613" style="zoom:67%;" />



==`i`的回文区间为什么说是**至少**[R',R]呢？有没有可能变得更远呢？==

不知道，需要验证X 是否等于 Y，以上面示例来说，当 `?`是 k 时候就可以更远，不是就不行

<img src="E:\linux学习PDF\linux\img\image-20221114105216034.png" alt="image-20221114105216034" style="zoom:60%;" />



### 时间复杂度计算

```go
// 伪代码

for i :=  0 ; i < len(str) ; i++{
    if (i 在 R的外部){
        从 i 开始往两边扩，R变大
    } else if (i’的回文区间在[L,R]内){
        // 直接等于 i’ 的
        pArr[i] = 某个O(1)表达式
    } else if (i’的回文区间一部分在[L,R]外部){
        // 与 [L'..i..L] 等大的区间，[R'..i..R]
        pArr[i] = 某个O(1)表达式
    } else (i’的回文区间左边界与[L,R]压线){ 
        需要判断从R之外的字符开始，往外扩，确定pArr[i]的答案
        若第一步就外扩失败，R不变，否则R会变大
    }
}
```

每个位置最多扩失败一次 O失败(N) ，情况1：1次， 情况2：1) 0 次  2）0次  3）1次

i 的变化幅度为 O(N) ，R也是只增不退变化幅度也是 O(N)

|             |  i   |  R   |
| ----------- | :--: | :--: |
| 情况1       | 变大 | 变大 |
| 情况2   1） | 变大 | 不变 |
| 2）         | 变大 | 不变 |
| 3）         | 变大 | 变大 |

时间复杂度为 O(N)



### 代码

```go
package chapter2
import "math"
func manacheString(str string) []byte {

	ch := []byte(str)
	n := len(str)
	ans := make([]byte, 2*n+1)
	j := 0
	for i := 0; i < 2*n+1; i++ {
		if i%2 == 0 {
			ans[i] = '#'
		} else {
			ans[i] = ch[j]
			j++
		}
	}
	return ans
}

func MaxLcpsLenght(s string) int {
	if len(s) == 0 {
		return -1
	}
	// 加入特殊字符
	str := manacheString(s)
	// 记录每个位置的回文半径数组
	pArr := make([]int, len(str))
	// 回文中心，回文右边界
	C, R := -1, -1
	max := math.MinInt64
	for i := 0; i < len(str); i++ {
		// pArr 至少的回文区间
		if R > i {
			// 2*C-i 是 i'位置，R-i是 [i,R]
			//最短的回文半径可能是i对应的i'的回文半径或者i到R的距离
			pArr[i] = min(pArr[2*C-i], R-i)
		} else {
			// 情况1 长度至少1
			pArr[i] = 1
		}

		// 往左右两边扩是否越界，不越界就进行判断
		// 这么写是为了缩短代码
		for i+pArr[i] < len(str) && i-pArr[i] > -1 {
			// 情况1 和 情况2 的 3)
			if str[i+pArr[i]] == str[i-pArr[i]] {
				pArr[i]++
			} else {
				// 情况2 的 1) 2)
				break
			}
		}
		// 查看是否有更往右
		if i+pArr[i] > R {
			R = i + pArr[i]
			C = i
		}
		max = Max(max, pArr[i])
	}

	return max - 1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

```





### 代码可能不懂的问题

pArr[i] = min(pArr[2*C-i], R-i) 

代表3种情况:

1. i'  在 [L,R]内，i 的 回文半径就为 i' 的， 就直接取 pArr[2*C-i]

   <img src="E:\linux学习PDF\linux\img\image-20221114161934668.png" alt="image-20221114161934668" style="zoom:50%;" />

2. i' 左边界超过 L，此时 i 的 回文半径就为 R-i

3. i' 的左边界压线L，i 回文半径至少 R - i 

总结起来就是取两者最小值。

