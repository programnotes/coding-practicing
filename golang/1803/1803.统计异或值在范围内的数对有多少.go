/*
 * @lc app=leetcode.cn id=1803 lang=golang
 *
 * [1803] 统计异或值在范围内的数对有多少
 *
 * https://leetcode.cn/problems/count-pairs-with-xor-in-a-range/description/
 *
 * algorithms
 * Hard (42.96%)
 * Likes:    87
 * Dislikes: 0
 * Total Accepted:    5.6K
 * Total Submissions: 11.5K
 * Testcase Example:  '[1,4,2,7]\n2\n6'
 *
 * 给你一个整数数组 nums （下标 从 0 开始 计数）以及两个整数：low 和 high ，请返回 漂亮数对 的数目。
 *
 * 漂亮数对 是一个形如 (i, j) 的数对，其中 0 <= i < j < nums.length 且 low <= (nums[i] XOR
 * nums[j]) <= high 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [1,4,2,7], low = 2, high = 6
 * 输出：6
 * 解释：所有漂亮数对 (i, j) 列出如下：
 * ⁠   - (0, 1): nums[0] XOR nums[1] = 5
 * ⁠   - (0, 2): nums[0] XOR nums[2] = 3
 * ⁠   - (0, 3): nums[0] XOR nums[3] = 6
 * ⁠   - (1, 2): nums[1] XOR nums[2] = 6
 * ⁠   - (1, 3): nums[1] XOR nums[3] = 3
 * ⁠   - (2, 3): nums[2] XOR nums[3] = 5
 *
 *
 * 示例 2：
 *
 * 输入：nums = [9,8,4,2,1], low = 5, high = 14
 * 输出：8
 * 解释：所有漂亮数对 (i, j) 列出如下：
 * ​​​​​    - (0, 2): nums[0] XOR nums[2] = 13
 * - (0, 3): nums[0] XOR nums[3] = 11
 * - (0, 4): nums[0] XOR nums[4] = 8
 * - (1, 2): nums[1] XOR nums[2] = 12
 * - (1, 3): nums[1] XOR nums[3] = 10
 * - (1, 4): nums[1] XOR nums[4] = 9
 * - (2, 3): nums[2] XOR nums[3] = 6
 * - (2, 4): nums[2] XOR nums[4] = 5
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 2 * 10^4
 * 1 <= nums[i] <= 2 * 10^4
 * 1 <= low <= high <= 2 * 10^4
 *
 *
 */

package jzoffer

// @lc code=start

// TODO:
//  1. Let's note that we can count all pairs with XOR ≤ K, so the answer would be to
//     subtract the number of pairs withs XOR < low from the number of pairs with XOR ≤ high.
//  2. For each value, find out the number of values when you XOR it with the result is ≤ K using a trie.

const _trieBitLen = 14

type trieNode struct {
	son [2]*trieNode
	cnt int
}

type Trie struct {
	root *trieNode
}

func (t *Trie) Put(v int) *trieNode {
	o := t.root
	for i := _trieBitLen; i >= 0; i-- {
		b := v >> i & 1
		if o.son[b] == nil {
			o.son[b] = &trieNode{}
		}
		o = o.son[b]
		o.cnt++
	}
	return o
}

func (t *Trie) CountLimitXOR(v, limit int) (cnt int) {
	o := t.root
	for i := _trieBitLen; i >= 0; i-- {
		b := v >> i & 1
		if limit>>i&1 > 0 {
			if o.son[b] != nil {
				cnt += o.son[b].cnt
			}
			b ^= 1
		}
		if o.son[b] == nil {
			return
		}
		o = o.son[b]
	}
	return
}

func countPairs(nums []int, low int, high int) (res int) {
	t := &Trie{root: &trieNode{}}
	t.Put(nums[0])
	for _, v := range nums[1:] {
		res += t.CountLimitXOR(v, high+1) - t.CountLimitXOR(v, low)
		t.Put(v)
	}
	return
}

// O(n*n)
func countPairsN2(nums []int, low int, high int) (res int) {
	// x^y=t 等价 y=x^t
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			v := nums[i] ^ nums[j]
			if v >= low && v <= high {
				res++
			}
		}
	}
	return
}

// @lc code=end
