# Morris

## 简介

Morris遍历：一种遍历二叉树的方式，并且[时间复杂度](https://so.csdn.net/so/search?q=时间复杂度&spm=1001.2101.3001.7020)O(N)，额外空间复杂度O(1)。**通过利用原树中大量空闲右指针的方式，达到节省空间的目的**。

#### Morris遍历的关键

**利用一棵树上大量的右指针空闲空间**

## Morris 遍历细节

![image-20221121193236508](E:\linux学习PDF\linux\img\image-20221121193236508.png)

## 演示过程

第1步,`cur` 指向头结点

<img src="E:\linux学习PDF\linux\img\image-20221121200916467.png" alt="image-20221121200916467" style="zoom:80%;" />

第2步，判断`cur`是否有左子树，有左子树则找到其左子树最右结点`mostRight` 

<img src="E:\linux学习PDF\linux\img\image-20221121200932004.png" alt="image-20221121200932004" style="zoom:80%;" />

第3步，，如果 `mostRight`的右指针为空，让其指向cur，然后cur 向左移动

<img src="E:\linux学习PDF\linux\img\image-20221121201010636.png" alt="image-20221121201010636" style="zoom:80%;" />

第4步，判断`cur`是否有左子树，如果有，找到其左子树最右结点`mostRight`，并让`mostRight`指向`cur`

<img src="E:\linux学习PDF\linux\img\image-20221121201334318.png" alt="image-20221121201334318" style="zoom:80%;" />

第5步，判断`cur`是否有左子树，**没有**则让 `cur`向右移动，来到右子树，右子树的`null->2`，所以此时cur来到2位置

<img src="E:\linux学习PDF\linux\img\image-20221121202035415.png" alt="image-20221121202035415" style="zoom:80%;" />

第6步，判断`cur`是否有左子树，有但是`mostRight == cur `，让`mostRight==nil`,又因为此时 `mostRight==nil`所以`cur`向右移动，来到5

<img src="E:\linux学习PDF\linux\img\image-20221121203209072.png" alt="image-20221121203209072" style="zoom:80%;" />

第7步，判断`cur`是否有左子树，没有，`cur` 向右移动，来到 1 位置

<img src="E:\linux学习PDF\linux\img\image-20221121203355227.png" alt="image-20221121203355227" style="zoom:80%;" />

第8步，判断 `cur`是否有左子树，有则找到左子树的最右结点`mostRight`，但此时 `mostRight == cur`，所以`mostRight == nil`,

cur向右移动到3

<img src="E:\linux学习PDF\linux\img\image-20221121203746501.png" alt="image-20221121203746501" style="zoom:80%;" />

<img src="E:\linux学习PDF\linux\img\image-20221121203913376.png" alt="image-20221121203913376" style="zoom:80%;" />



第9步，判断3的左子树，有左子树找到左子树最右结点`mostRight`，若`mostRight==nil`，则指向`cur`

<img src="E:\linux学习PDF\linux\img\image-20221121203959486.png" alt="image-20221121203959486" style="zoom:80%;" />

第10步，`cur`向左移动，然后判断是否有左子树，没有`cur`向右移动，移动到 3 

<img src="E:\linux学习PDF\linux\img\image-20221121204151233.png" alt="image-20221121204151233" style="zoom:80%;" />

<img src="E:\linux学习PDF\linux\img\image-20221121204259539.png" alt="image-20221121204259539" style="zoom:80%;" />

第11步，判断`cur`是否有左子树，有，找到左子树最右结点`mostRight`，若`mostRight == cur`则 让`mostRight == nil`，`cur`向右移动

<img src="E:\linux学习PDF\linux\img\image-20221121204441814.png" alt="image-20221121204441814" style="zoom:80%;" />

第12步，判断`cur`是否有左子树，没有则`cur`向右移动，此时 `cur == nil` 遍历结束

<img src="E:\linux学习PDF\linux\img\image-20221121204551740.png" alt="image-20221121204551740" style="zoom:80%;" />

## 代码

```go
package chapter3

type Node struct {
	val   int
	left  *Node
	right *Node
}

func Morris(head *Node) {
	if head == nil {
		return
	}
	cur := head
	var mostRight *Node = nil
	for cur != nil {
		// mostRight 为cur的左结点
		mostRight = cur.left
		if mostRight != nil { // 左子树不为空
			// 找到其左子树最右结点
			for mostRight.right != nil && mostRight.right != cur {
				mostRight = mostRight.right
			}
			// 此时 mostRight是左子树最右结点
			if mostRight.right == nil { // 第一次来到cur
				mostRight.right = cur
				cur = cur.left
				continue // 返回到cur，继续
			} else { // mostRight.right == cur
				mostRight.right = nil
			}
		}
		cur = cur.right
	}
}

```



## 应用场景

### Morris先序

已知Morris遍历信息，求二叉树先序遍历

Morris遍历 :1，2，4，8，2，1，5，3，6，3，7

先序遍历：1，2，4，8，5，3，6，7

思路：只有一次直接打印，有两次的，遇见第一次打印

```go
func MorrisPre(head *Node) {
	if head == nil {
		return
	}
	cur := head
	var mostRight *Node = nil
	for cur != nil {
		mostRight = cur.left
		// 有左子树
		if mostRight != nil {
			for mostRight.right != nil && mostRight.right != cur {
				mostRight = mostRight.right
			}
			// mostRight 变成左子树上最右结点
			if mostRight.right == nil { // 第一次来到cur
				fmt.Printf("%v", cur.val)
				mostRight.right = cur
				cur = cur.left
				continue
			} else { // mostRight.right == cur
				mostRight.right = nil
			}

		} else { // 没有左子树
			fmt.Printf("%v", cur.val)
		}
		cur = cur.right
	}
}
```



### Morris中序

Morris遍历 :1，2，4，8，2，1，5，3，6，3，7

中序遍历：4，8，2，5，1，6，3，7

思路：只有一次直接打印，有两次的，遇见第二次打印

```go
func MorriOrd(head *Node) {
	if head == nil {
		return
	}
	cur := head
	var mostRight *Node = nil
	for cur != nil {
		mostRight = cur.left
		// 有左子树
		if mostRight != nil {
			for mostRight.right != nil && mostRight.right != cur {
				mostRight = mostRight.right
			}
			// mostRight 变成左子树上最右结点
			if mostRight.right == nil { // 第一次来到cur
				fmt.Printf("%v", cur.val)
				mostRight.right = cur
				cur = cur.left
				continue
			} else { // mostRight.right == cur
				mostRight.right = nil
			}

		} else { // 没有左子树
			fmt.Printf("%v", cur.val)
		}
		cur = cur.right
	}
}
```



### Morris后序

Morris遍历 :1，2，4，8，2，1，5，3，6，3，7

后序遍历：8，4，5，2，6，7，3，1

思路：只出现一次或者第一次遇见的结点都跳过，第二次遇见，逆序打印此结点**左树右边界**，遍历结束后，逆序打印整棵树的右边界

示例：

1，2，4，8都跳过，第二次遇见 2，逆序打印其右边界，左树右边界为8，4

<img src="E:\linux学习PDF\linux\img\image-20221122211005774.png" alt="image-20221122211005774" style="zoom:67%;" />

继续判断，遇到第二次 1 ，逆序打印左树的右边界 2

<img src="E:\linux学习PDF\linux\img\image-20221122211459872.png" alt="image-20221122211459872" style="zoom:67%;" />

第二次遇到 3 ，逆序打印左树的右边界 6 ，结束

<img src="E:\linux学习PDF\linux\img\image-20221122211628595.png" alt="image-20221122211628595" style="zoom:67%;" />

遍历结束后，逆序打印整棵树的右边界  7

<img src="E:\linux学习PDF\linux\img\image-20221122211745765.png" alt="image-20221122211745765" style="zoom:67%;" />



```go
func MorrisPos(head *Node) {
	if head == nil {
		return
	}
	cur := head
	var mostRight *Node = nil
	for cur != nil {
		mostRight = cur.left
		// 有左子树
		if mostRight != nil {
			for mostRight.right != nil && mostRight.right != cur {
				mostRight = mostRight.right
			}
			// mostRight 变成左子树上最右结点
			if mostRight.right == nil { // 第一次来到cur
				mostRight.right = cur
				cur = cur.left
				continue
			} else { // mostRight.right == cur
				mostRight.right = nil
				// 逆序打印此结点左树右边界
				printEdge(cur.left)
			}
		}
		cur = cur.right
	}
	// 打印整棵树的右边界
	printEdge(head)
}

func printEdge(x *Node) {
	// 逆序
	tail := reverseEdge(x)
	cur := tail
	for cur != nil {
		fmt.Printf("%v", cur.val)
		cur = cur.right
	}
	// 复原
	reverseEdge(tail)
}

func reverseEdge(from *Node) *Node {
	var pre *Node = nil
	var next *Node = nil
	for from != nil {
		next = from.right
		from.right = pre
		pre = from
		from = next
	}
	return pre
}
```



### 检验是否是二叉搜索树(BST)

正常思路：中序遍历，然后记录，打印看是不是递增序列



```go
func Morris_isBTS(head *Node) bool {
	if head == nil {
		return true
	}
	cur := head
	var mostRight *Node = nil
	preVal := math.MinInt64
	for cur != nil {
		mostRight = cur.left
		// 如果有左子树，则找到其左子树最右结点
		if mostRight != nil {
			for mostRight.right != nil && mostRight.right != cur {
				mostRight = mostRight.right
			}
			if mostRight.right == nil {
				mostRight.right = cur
				cur = cur.left
				continue
			} else { //mostRight.right == cur
				mostRight.right = nil
			}
		}
		if preVal <= cur.val {
			return false
		}
		preVal = cur.val
		cur = cur.right
	}
	return true
}

```

