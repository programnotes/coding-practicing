/*
* @lc app=leetcode.cn id=240 lang=golang
*
* [240] 搜索二维矩阵 II
*
* https://leetcode.cn/problems/search-a-2d-matrix-ii/description/
*
  - algorithms
  - Medium (52.36%)
  - Likes:    1184
  - Dislikes: 0
  - Total Accepted:    333.2K
  - Total Submissions: 636.3K
  - Testcase Example:  '[[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]]\n' +
    '5'

*
* 编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：
*
*
* 每行的元素从左到右升序排列。
* 每列的元素从上到下升序排列。
*
*
*
*
* 示例 1：
*
*
* 输入：matrix =
* [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]],
* target = 5
* 输出：true
*
*
* 示例 2：
*
*
* 输入：matrix =
* [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]],
* target = 20
* 输出：false
*
*
*
*
* 提示：
*
*
* m == matrix.length
* n == matrix[i].length
* 1 <= n, m <= 300
* -10^9 <= matrix[i][j] <= 10^9
* 每行的所有元素从左到右升序排列
* 每列的所有元素从上到下升序排列
* -10^9 <= target <= 10^9
*
*
*/
package jzoffer

import "sort"

// @lc code=start
// 逐行二分
func searchMatrix(matrix [][]int, target int) bool {

	for _, v := range matrix {
		i := sort.SearchInts(v, target)
		if i < len(v) && v[i] == target {
			return true
		}
	}
	return false
}

// TODO
// Z 字形搜索
// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/search-a-2d-matrix-ii/solution/sou-suo-er-wei-ju-zhen-ii-by-leetcode-so-9hcx/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func searchMatrix1(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	x, y := 0, n-1
	for x < m && y >= 0 {
		if matrix[x][y] == target {
			return true
		}
		if matrix[x][y] > target {
			y--
		} else {
			x++
		}
	}
	return false
}

// @lc code=end
