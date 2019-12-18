package reconstructQueue

/**
406. 根据身高重建队列
假设有打乱顺序的一群人站成一个队列。 每个人由一个整数对(h, k)表示，其中h是这个人的身高，
k是排在这个人前面且身高大于或等于h的人数。 编写一个算法来重建这个队列。

示例
输入:
[[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]
输出:
[[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]

为了使插入操作不影响后续的操作，身高较高的学生应该先做插入操作，否则身高较小的学生原先正确插入的第 k 个位置可能会变成第 k+1 个位置。
*/

import "sort"

func reconstructQueue(people [][]int) [][]int {
	res := make([][]int, 0, len(people))

	sort.Sort(persons(people))

	insert := func(idx int, person []int) {
		res = append(res, person)
		// 插入到末尾
		if len(res)-1 == idx {
			return
		}
		// 插入到中间位置
		copy(res[idx+1:], res[idx:])
		res[idx] = person
	}

	for _, p := range people {
		insert(p[1], p)
	}

	return res
}

// 优先 h 降序，其次 k 升序
type persons [][]int

func (p persons) Len() int {
	return len(p)
}

func (p persons) Less(i, j int) bool {
	if p[i][0] == p[j][0] {
		return p[i][1] < p[j][1]
	}
	return p[i][0] > p[j][0]
}

func (p persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
