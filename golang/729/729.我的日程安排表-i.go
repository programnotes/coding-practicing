/*
 * @lc app=leetcode.cn id=729 lang=golang
 *
 * [729] 我的日程安排表 I
 *
 * https://leetcode.cn/problems/my-calendar-i/description/
 *
 * algorithms
 * Medium (53.95%)
 * Likes:    123
 * Dislikes: 0
 * Total Accepted:    14.5K
 * Total Submissions: 26.8K
 * Testcase Example:  '["MyCalendar","book","book","book"]\n[[],[10,20],[15,25],[20,30]]'
 *
 * 实现一个 MyCalendar 类来存放你的日程安排。如果要添加的日程安排不会造成 重复预订 ，则可以存储这个新的日程安排。
 *
 * 当两个日程安排有一些时间上的交叉时（例如两个日程安排都在同一时间内），就会产生 重复预订 。
 *
 * 日程可以用一对整数 start 和 end 表示，这里的时间是半开区间，即 [start, end), 实数 x 的范围为，  start <= x <
 * end 。
 *
 * 实现 MyCalendar 类：
 *
 *
 * MyCalendar() 初始化日历对象。
 * boolean book(int start, int end) 如果可以将日程安排成功添加到日历中而不会导致重复预订，返回 true 。否则，返回
 * false 并且不要将该日程安排添加到日历中。
 *
 *
 *
 *
 * 示例：
 *
 *
 * 输入：
 * ["MyCalendar", "book", "book", "book"]
 * [[], [10, 20], [15, 25], [20, 30]]
 * 输出：
 * [null, true, false, true]
 *
 * 解释：
 * MyCalendar myCalendar = new MyCalendar();
 * myCalendar.book(10, 20); // return True
 * myCalendar.book(15, 25); // return False ，这个日程安排不能添加到日历中，因为时间 15
 * 已经被另一个日程安排预订了。
 * myCalendar.book(20, 30); // return True ，这个日程安排可以添加到日历中，因为第一个日程安排预订的每个时间都小于
 * 20 ，且不包含时间 20 。
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= start < end <= 10^9
 * 每个测试用例，调用 book 方法的次数最多不超过 1000 次。
 *
 *
 */
//
package jzoffer

import (
	"log"
	"math"
)

type Node struct {
	prev, next *Node
	start, end int
}

// @lc code=start
type MyCalendar struct {
	head, tail *Node
	nsMap      map[int]*Node
}

func Constructor() MyCalendar {

	return MyCalendar{nsMap: make(map[int]*Node)}
}

func (this *MyCalendar) Book(start int, end int) bool {
	if len(this.nsMap) == 0 {
		node := &Node{start: start, end: end}
		this.nsMap[start/100] = node
		this.head = node
		this.tail = node
		return true
	}
	if this.head.start > end {
		node := &Node{start: start, end: end}
		this.nsMap[start/100] = node
		tmp := this.head
		this.head = node
		this.head.next = tmp
		tmp.prev = this.head
		return true
	}
	if this.tail.end < start {
		node := &Node{start: start, end: end}
		this.nsMap[start/100] = node
		tmp := this.tail
		this.tail = node
		this.head.prev = tmp
		tmp.next = this.head
		return true
	}
	modStart := start / 100
	modeEnd := end / 100

	var sNode, eNode *Node
	sK, eK := math.MinInt32, math.MaxInt32
	for k, v := range this.nsMap {
		if k <= modStart && k > sK {
			sNode = v
			sK = k
		}
		if k >= modeEnd && k < eK {
			eNode = v
			eK = k
		}
	}
	if sNode == nil && eNode == nil {
		log.Println("sNode == nil && eNode == nil")
		return false
	}

	for sNode != nil {
		if sNode.start < start && sNode.end <= start && (sNode.next == nil || (sNode.next != nil && sNode.start > start && sNode.next.start > end)) {
			node := &Node{start: start, end: end, prev: sNode, next: sNode.next}
			this.nsMap[start/100] = node
			tmp := sNode.next
			sNode.next = node
			if tmp != nil {
				tmp.prev = node
			}
			return true
		}
		sNode = sNode.next
	}
	for eNode != nil {
		if eNode.start < start && eNode.end <= start && (eNode.next == nil || (eNode.next != nil && eNode.start > start && eNode.next.start > end)) {
			node := &Node{start: start, end: end, prev: sNode, next: sNode.next}
			this.nsMap[start/100] = node
			tmp := sNode.next
			sNode.next = node
			if tmp != nil {
				tmp.prev = node
			}
			return true
		}
		eNode = eNode.next
	}
	return false
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
// @lc code=end
