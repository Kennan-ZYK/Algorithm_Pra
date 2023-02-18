package chapter1

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	us := Init(isConnected)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if isConnected[i][j] == 1 {
				us.Union(i, j)
			}
		}
	}
	return us.sum
}

type unionSet struct {
	count  []int // rank[i]表示以i为根的树的高度
	parent []int
	sum    int
}

func Init(isConnected [][]int) *unionSet {
	us := &unionSet{}
	us.count = make([]int, len(isConnected))
	us.parent = make([]int, len(isConnected))
	us.sum = len(isConnected)
	for i := 0; i < len(isConnected); i++ {
		us.parent[i] = i
		us.count[i] = 1
	}
	return us
}

func (us *unionSet) FindHead(x int) int {
	if us.parent[x] != x {
		us.parent[x] = us.FindHead(us.parent[x])
	}
	return us.parent[x]
}

func (us *unionSet) Union(n1, n2 int) {
	p := us.FindHead(n1)
	q := us.FindHead(n2)
	if p != q {
		if us.count[p] < us.count[q] {
			us.parent[p] = q
			us.count[p] = us.count[p] + us.count[q]
			us.sum--
		} else { //到这步说明高度相等,谁指向谁都可以,高度要+1
			us.parent[p] = q
			us.count[p] = us.count[p] + us.count[q]
			us.sum--
		}
	}
}
