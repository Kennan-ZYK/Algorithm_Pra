package chapter3

type Node struct {
	val   int
	left  *Node
	right *Node
}

/*
Morris遍历细节
假设来到当前节点cur，开始时cur来到头节点位置
1)如果cur没有左孩子，cur向右移动(cur = cur.right)
2)如果cur有左孩子，找到左子树上最右的节点mostRight:
	a.如果mostRight的右指针指向空，让其指向cur然后cur向左移动(cur = cur.left)
	b.如果mostRight的右指针指向cur，让其指向null然后cur向右移动(cur = cur.right)
3)cur为空时遍历停止
*/

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
