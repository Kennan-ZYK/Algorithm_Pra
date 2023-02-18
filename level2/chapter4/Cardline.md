

**basecase**

```go
	// 根据 i == j return arr[i]，可得
	for i := range fdp {
		fdp[i][i] = arr[i]
	}
```

<img src="E:\linux学习PDF\linux\img\image-20221210092722458.png" alt="image-20221210092722458" style="zoom:57%;" />



分析**普遍位置**如何依赖的

对于 **f** 来说调用了 **s** 的子过程，**s** 调用了 **f** 的子过程，

根据暴力递归可以知道，**f** 表 和 **s** 表中的 `(L+1,R)，(L,R-1)`有关，**s** 表 和 **f** 表中的 `(L+1,R)，(L,R-1)`，所以两张表可以交替完成

```go
func f(arr []int, L, R int) int {
    // 先手拿
    if L == R {
        return arr[L]
    }
    return Max(arr[L]+s(arr, L+1, R), arr[R]+s(arr, L, R-1))
}

func s(arr []int, L, R int) int {
    if L == R {
        return 0
    }
    return Min(f(arr, L+1, R), f(arr, L, R-1))
}
```



<img src="E:\linux学习PDF\linux\img\image-20221210094932350.png" alt="image-20221210094932350" style="zoom:67%;" />



**代码实现**

```go
for stratCol := 1; stratCol < N; stratCol++ {
    L := 0
    R := stratCol
    for R < N {
        fdp[L][R] = Max(arr[L]+sdp[L+1][R], arr[R]+sdp[L][R-1])
        fdp[L][R] = Min(fdp[L+1][R], fdp[L][R-1])
        L++
        R++
    }
}
```

**举例：[1,4,5,4]**

画图自己试试

![image-20221210103928463](E:\linux学习PDF\linux\img\image-20221210103928463.png)