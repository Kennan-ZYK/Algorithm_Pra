# 并查集

并查集Union find，用来表示网络节点之间是否连接的集合。
 这里的网络是一抽抽象的概念，不仅包括互联网，人与人形成的网络，道路之间形成的网络，迷宫网络等等。
 数据存储格式如下

![image-20221110162244033](E:\linux学习PDF\linux\img\image-20221110162244033.png)



1）i 为节点索引， set[i] 为所属集合，若两个索引所属集合相同，表示这两节点相连
 2）跟传统树结构不同的是，并查集是一种子节点指向父节点的数据结构

```
对并查集来说，主要支持两个动作：
1）union（p，q） // 将p，q两个节点相连
2）isConnected（p，q）  // 查询p，q是否相连，即是否属于同一个集合
```

<img src="E:\linux学习PDF\linux\img\image-20221110162254378.png" alt="image-20221110162254378" style="zoom:70%;" />

```go
下面从简单到深入，一步步优化并查集，实现5版不同的并查集，并对其性能进行对比  
//
1）第一版union find
特点：union（p，q）操作时，将set[p] = set[q]，之后遍历一遍set，
     找出所有set[i] == set[p]的索引，执行 set[i] = set[q]

// 实现
type unionSet struct {
    set  []int
}

func NewUnionSet(size int) *unionSet {
    buf := make([]int, size)
    for i := 0; i < size; i++ {
        buf[i] = i  // 初始时，所有节点均指向自己
    }
    return &unionSet{set:  buf}
}

func (set *unionSet) GetSize() int {
    return len(set.set)
}

func (set *unionSet) GetID(p int) (int, error) {
    if p < 0 || p >  len(set.set) {
        return 0, errors.New(
            "failed to get ID,index is illegal.")
    }
    return set.set[p], nil
}

func (set *unionSet) IsConnected(p, q int) (bool, error) {
    if p < 0 || p >  len(set.set) || q < 0 || q >  len(set.set) {
        return false, errors.New(
            "failed to get ID,index is illegal.")
    }
    return set.set[p] == set.set[q], nil
}

func (set *unionSet) Union(p, q int) error {
    b, err := set.IsConnected(p, q)
    if err != nil {
        return err
    }

    if !b {
        pID := set.set[p]
        qID := set.set[q]
        for k, v := range set.set {
            if v == pID {
                set.set[k] = qID
            }
        }
    }
    return nil
}
时间复杂度分析：
union()          O(n)
isConnected()    O(1)
（2）第二版union find
特点：union（p，q）操作时，找出p的根节点n，让set[q] = n 
// 实现
type unionSet struct {
    set  []int
}

func NewUnionSet(size int) *unionSet {
    buf := make([]int, size)
    for i := 0; i < size; i++ {
        buf[i] = i
    }
    return &unionSet{ set: buf}
}

func (set *unionSet) GetSize() int {
    return len(set.set)
}

func (set *unionSet) GetID(p int) (int, error) {
    if p < 0 || p > len(set.set) {
        return 0, errors.New(
            "failed to get ID,index is illegal.")
    }

    return set.getRoot(p), nil
}

// getRoot: 找出p的根节点,时间复杂度为O(h),h为树的高度
func (set *unionSet) getRoot(p int) int {
    if set.set[p] == p {
        return p
    }
    return set.getRoot(set.set[p])
}

func (set *unionSet) IsConnected(p, q int) (bool, error) {
    if p < 0 || p > len(set.set) || q < 0 || q > len(set.set) {
        return false, errors.New(
            "error: index is illegal.")
    }
    return set.getRoot(set.set[p]) == set.getRoot(set.set[q]), nil
}

func (set *unionSet) Union(p, q int) error {
    if p < 0 || p > len(set.set) || q < 0 || q > len(set.set) {
        return errors.New(
            "error: index is illegal.")
    }

    pRoot := set.getRoot(p)
    qRoot := set.getRoot(q)

    if pRoot != qRoot {
        set.set[pRoot] = qRoot
    }
    return nil
}
时间复杂度分析： // h为树的高度
union()          O(h)
isConnected()    O(h)
相比第一版，这个时间复杂度更好接受，不过在极端的情况下，
比如执行union(0,1),union(0,2)...union(0,n),这样树变成单链，
时间复杂度都为O(n)，要继续优化
（3）第三版union find
特点：加入numNode数组，numNode[ i ] 表示集合 i 中节点数量，每次执行union时，
     取出numNode的节点数量，让节点数量小的指向节点数量大的。如下图：
```

<img src="E:\linux学习PDF\linux\img\image-20221110162318788.png" alt="image-20221110162318788" style="zoom:67%;" />

```go
// 实现
type unionSet struct {
    numNode  []int // numNode[i]表示集合i中是节点数量
    set []int
}

func NewUnionSet(size int) *unionSet {
    buf1 := make([]int, size)
    for i := 0; i < size; i++ {
        buf1[i] = i
    }
    buf2 := make([]int, size)
    for i := 0; i < size; i++ {
        buf2[i] = 1  // 初始节点数量均为1
    }

    return &unionSet{
        numNode:  buf2,
        set: buf1,
    }
}

func (set *unionSet) GetSize() int {
    return len(set.set)
}

func (set *unionSet) GetID(p int) (int, error) {
    if p < 0 || p > len(set.set) {
        return 0, errors.New(
            "failed to get ID,index is illegal.")
    }
    return set.getRoot(p), nil
}

func (set *unionSet) getRoot(p int) int {
    for p != set.set[p] {
        p = set.set[p]
    }
    return p
}

func (set *unionSet) IsConnected(p, q int) (bool, error) {
    if p < 0 || p > len(set.set) || q < 0 || q > len(set.set) {
        return false, errors.New(
            "error: index is illegal.")
    }
    return set.getRoot(set.set[p]) == set.getRoot(set.set[q]), nil
}

func (set *unionSet) Union(p, q int) error {
    if p < 0 || p > len(set.set) || q < 0 || q > len(set.set) {
        return errors.New(
            "error: index is illegal.")
    }

    pRoot := set.getRoot(p)
    qRoot := set.getRoot(q)

    if pRoot != qRoot {
        if set.numNode[pRoot] < set.numNode[qRoot] {
            set.set[pRoot] = qRoot 
            set.numNode[qRoot] += set.numNode[pRoot]
        } else {
            set.set[qRoot] = pRoot
            set.numNode[pRoot] += set.numNode[qRoot]
        }
    }
    return nil
}
时间复杂度分析： // h为树的高度
union()          O(h)
isConnected()    O(h)
相比第二版，这版能有效避免树变为单链的情况，不过还能继续优化，看下图一个反面例子：
```

<img src="E:\linux学习PDF\linux\img\image-20221110162331324.png" alt="image-20221110162331324" style="zoom:67%;" />

```go
（4）第四版union find
特点：加入rank数组，rank[ i ] 表示集合高度，每次执行union(p,q)，取出p，q所在集合的高度，
      让高度低的指向高度高的。
// 实现
type unionSet struct {
    rank []int // rank[i]表示以i为根的树的高度
    set  []int
}

func NewUnionSet(size int) *unionSet {
    buf1 := make([]int, size)
    for i := 0; i < size; i++ {
        buf1[i] = i
    }
    buf2 := make([]int, size)
    for i := 0; i < size; i++ {
        buf2[i] = 1
    }

    return &unionSet{
        rank: buf2,
        set:  buf1,
    }
}

func (set *unionSet) GetSize() int {
    return len(set.set)
}

func (set *unionSet) GetID(p int) (int, error) {
    if p < 0 || p > len(set.set) {
        return 0, errors.New(
            "failed to get ID,index is illegal.")
    }

    return set.getRoot(p), nil
}

func (set *unionSet) getRoot(p int) int {
    for p != set.set[p] {
        p = set.set[p]
    }
    return p
}

func (set *unionSet) IsConnected(p, q int) (bool, error) {
    if p < 0 || p > len(set.set) || q < 0 || q > len(set.set) {
        return false, errors.New(
            "error: index is illegal.")
    }
    return set.getRoot(set.set[p]) == set.getRoot(set.set[q]), nil
}

func (set *unionSet) Union(p, q int) error {
    if p < 0 || p > len(set.set) || q < 0 || q > len(set.set) {
        return errors.New(
            "error: index is illegal.")
    }

    pRoot := set.getRoot(p)
    qRoot := set.getRoot(q)

    if pRoot != qRoot {
        if set.rank[pRoot] < set.rank[qRoot] {
            set.set[pRoot] = qRoot
        } else if set.rank[qRoot] < set.rank[pRoot] {
            set.set[qRoot] = pRoot
        } else { //到这步说明高度相等,谁指向谁都可以,高度要+1
            set.set[pRoot] = qRoot
            set.rank[qRoot] += 1
        }
    }
    return nil
}
时间复杂度分析： // h为树的高度
union()          O(h)
isConnected()    O(h)
从理论上讲，是比第三版更好的优化方式
先对前4版做一个性能测试对比
const count = 1000

func main() {
    nums := []int{}
    for i := 0; i < 50*count; i++ {
        nums = append(nums, rand.Intn(count))
    }
    test1(nums)
    test2(nums)
    test3(nums)
    test4(nums)
}

func test1(nums []int) {
    s := unionfind1.NewUnionSet(count)
    t := time.Now()
    for i := 1; i < 50*count; i++ {
        s.GetID(nums[i])
        s.Union(nums[i-1], nums[i])
    }
    fmt.Println("第一版时间: ", t.Sub(time.Now()))
}
func test2(nums []int) {
    s := unionfind2.NewUnionSet(count)
    t := time.Now()
    for i := 1; i < 50*count; i++ {
        s.GetID(nums[i])
        s.Union(nums[i-1], nums[i])
    }
    fmt.Println("第二版时间: ", t.Sub(time.Now()))
}
func test3(nums []int) {
    s := unionfind4.NewUnionSet(count)
    t := time.Now()
    for i := 1; i < 50*count; i++ {
        s.GetID(nums[i])
        s.Union(nums[i-1], nums[i])
    }
    fmt.Println("第三版时间: ", t.Sub(time.Now()))
}

func test4(nums []int) {
    s := unionfind5.NewUnionSet(count)
    t := time.Now()
    for i := 1; i < 50*count; i++ {
        s.GetID(nums[i])
        s.Union(nums[i-1], nums[i])
    }
    fmt.Println("第四版时间: ", t.Sub(time.Now()))
}
// 测试结果
第一版时间:  -5.101884ms
第二版时间:  -423.963439ms
第三版时间:  -1.912094ms
第四版时间:  -1.788539ms
分析：
1）加入数量优化或者高度优化的并查集性能明显优于前两版，后续要加大数量级，不再对前两版测试；
2）第二版时间更慢是因为在go的底层中，访问连续的内存地址速度要快于通过索引访问，而第二版
寻找根节点是通过索引访问的
3）第三和第四版的优化明显抵消了通过索引访问慢的劣势；
当数量级增加到1000万后，执行时间是：
第三版时间:  -1.100688529s        第三版时间:  -972.220009ms
第四版时间:  -1.066303094s        第四版时间:  -1.090987724s
两者性能差异非常小，不过第四版理论上能避免第三版最坏情况的出现，实际应用中使用第四版居多，
还能优化，加入路径压缩功能
```



