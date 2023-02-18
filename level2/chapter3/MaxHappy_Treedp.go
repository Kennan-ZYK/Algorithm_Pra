package chapter3

/*
派对的最大快乐值
员工信息的定义如下class Employee {
publicint happy;// 这名员工可以带来的快乐值
List<Employee> subordinates; // 这名员工有哪些直接下级
}
公司的每个员工都符合Employee 类的描述。整个公司的人员结构可以看作是一棵标准的、没有环的多叉树。
树的头节点是公司唯一的老板。除老板之外的每个员工都有唯一的直接上级。 叶节点是没有任何下属的基层员工(subordinates列表为空)，
除基层员工外，每个员工都有一个或多个直接下级。这个公司现在要办party，你可以决定哪些员工来，哪些员工不来。
但是要遵循如下规则
1.如果某个员工来了，那么这个员工的所有直接下级都不能来
2.派对的整体快乐值是所有到场员工快乐值的累加
3.你的目标是让派对的整体快乐值尽量大
给定一棵多叉树的头节点boss，请返回派对的最大快乐值。
*/

/*
			8
		20		 30
	1		2  31   20
选择是 8 不来， 20 来，30不来，31，20来
*/

/*
思路：这类题都可以用树型dp，树型dp只要确定好条件，就可以了
			x
	a		b	  c
q	w	e   s   y	z
情况：
x来，其下级不来abc, qwe,s,yz可以来
x不来，其下级可以来a,b,c可以来

所以结合多种情况可以得出
x参加， x.happy + a整个a不来happy的最大值 + b整个b不来happy的最大值 + c整个c不来的最大值
x不参加， 0 + Max(a来整个happy最大值，a不来整个happy最大值) + Max(b来整个happy最大值，b不来整个happy最大值) + Max(c来整个happy最大值，c不来整个happy最大值)
*/

type Employee struct {
	happy int
	next  *Employee
}

type Info struct {
	comeMax     int
	dontcomeMax int
}

func MaxHappy(boss Employee) int {
	headInfo := process(boss)
	return Max(headInfo.comeMax, headInfo.dontcomeMax)
}

func process(x Employee) Info {
	// 基层员工
	if x.next == nil {
		return Info{x.happy, 0}
	}
	// x参与，x不参与
	come := x.happy
	dontcome := 0
	for x.next != nil {
		nextInfo := process(*x.next)
		// x来，x的下级员工就不来
		come += nextInfo.dontcomeMax
		// x不来，x的下级员工就来
		dontcome += nextInfo.comeMax
	}
	return Info{come, dontcome}

}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
