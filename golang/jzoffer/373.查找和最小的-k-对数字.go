/*
 * @Author: yiGmMk marvelousme@163.com
 * @Date: 2022-06-02 17:25:21
 * @LastEditors: yiGmMk marvelousme@163.com
 * @LastEditTime: 2022-06-02 17:25:22
 * @FilePath: /go-tool/home/admin/code/coding-practicing/golang/jzoffer/373.查找和最小的-k-对数字.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
/*
 * @lc app=leetcode.cn id=373 lang=golang
 *
 * [373] 查找和最小的 K 对数字
 *
 * https://leetcode.cn/problems/find-k-pairs-with-smallest-sums/description/
 *
 * algorithms
 * Medium (41.58%)
 * Likes:    401
 * Dislikes: 0
 * Total Accepted:    49K
 * Total Submissions: 117.8K
 * Testcase Example:  '[1,7,11]\n[2,4,6]\n3'
 *
 * 给定两个以 升序排列 的整数数组 nums1 和 nums2 , 以及一个整数 k 。
 *
 * 定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。
 *
 * 请找到和最小的 k 个数对 (u1,v1),  (u2,v2)  ...  (uk,vk) 。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
 * 输出: [1,2],[1,4],[1,6]
 * 解释: 返回序列中的前 3 对数：
 * ⁠    [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
 *
 *
 * 示例 2:
 *
 *
 * 输入: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
 * 输出: [1,1],[1,1]
 * 解释: 返回序列中的前 2 对数：
 * [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
 *
 *
 * 示例 3:
 *
 *
 * 输入: nums1 = [1,2], nums2 = [3], k = 3
 * 输出: [1,3],[2,3]
 * 解释: 也可能序列中所有的数对都被返回:[1,3],[2,3]
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= nums1.length, nums2.length <= 10^5
 * -10^9 <= nums1[i], nums2[i] <= 10^9
 * nums1 和 nums2 均为升序排列
 * 1 <= k <= 1000
 *
 *
 */

// top k的问题首先应该想到用堆解决
package jzoffer


// import "container/heap"

// func kSmallestPairs(nums1, nums2 []int, k int) (ans [][]int) {
// 	m, n := len(nums1), len(nums2)
// 	h := HP{nil, nums1, nums2}
// 	for i := 0; i < k && i < m; i++ {
// 		h.data = append(h.data, pair{i, 0})
// 	}
// 	for h.Len() > 0 && len(ans) < k {
// 		p := heap.Pop(&h).(pair)
// 		i, j := p.i, p.j
// 		ans = append(ans, []int{nums1[i], nums2[j]})
// 		if j+1 < n {
// 			heap.Push(&h, pair{i, j + 1})
// 		}
// 	}
// 	return
// }

// @lc code=end
