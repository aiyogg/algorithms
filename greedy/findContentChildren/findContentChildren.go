package findContentChildren

import "sort"

/**
455. 分发饼干
题目描述：每个孩子都有一个满足度 grid，每个饼干都有一个大小 size，只有饼干的大小大于等于一个孩子的满足度，该孩子才会获得满足。求解最多可以获得满足的孩子数量。

贪心思想：
- 给一个孩子的饼干应当尽量小并且又能满足该孩子，这样大饼干才能拿来给满足度比较大的孩子。
- 因为满足度最小的孩子最容易得到满足，所以先满足满足度最小的孩子。
*/
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	gi, si := 0, 0
	for gi < len(g) && si < len(s) {
		if s[si] >= g[gi] {
			gi++
		}
		si++
	}
	return gi
}
