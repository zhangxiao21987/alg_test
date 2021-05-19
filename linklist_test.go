/*
 add by jinchaozhang. 2021/4/22
*/
package alg_test_test

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(node *ListNode) *ListNode {
	if node == nil {
		return node
	}
	var preNode *ListNode
	var next *ListNode
	head := node
	for head != nil {
		next = head.Next
		head.Next = preNode
		preNode = head
		head = next
	}

	return preNode
}

func printNode(node *ListNode) {
	if node == nil {
		fmt.Println("")
		return
	}
	currNode := node
	for currNode != nil {
		fmt.Print(currNode.Val)
		fmt.Print(" ")
		currNode = currNode.Next
	}
	fmt.Println("")
}

func generateListNode(maxLen, maxVal int) *ListNode {
	var head *ListNode
	len := rand.Intn(maxLen)
	if len > 0 {
		head = &ListNode{Val: rand.Intn(maxVal)}
		curNode := head
		for i := 1; i < len; i++ {
			curNode.Next = &ListNode{Val: rand.Intn(maxVal)}
			curNode = curNode.Next
		}
	}
	return head
}

func getListNodeArr(node *ListNode) []int {
	nodeArr := make([]int, 0)
	curNode := node
	for curNode != nil {
		nodeArr = append(nodeArr, curNode.Val)
		curNode = curNode.Next
	}
	return nodeArr
}

func checkReversedListNode(arr []int, node *ListNode) bool {
	if len(arr) == 0 && node == nil {
		return true
	}
	curNode := node
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] != curNode.Val {
			return false
		}
		curNode = curNode.Next
	}
	return true
}

func TestReverseNode(t *testing.T) {
	fmt.Println("start")
	maxLen := 30
	maxVal := 1000
	maxTimes := 1000000
	for i := 0; i < maxTimes; i++ {
		node := generateListNode(maxLen, maxVal)
		arr := getListNodeArr(node)
		rNode := reverse(node)
		ok := checkReversedListNode(arr, rNode)
		if !ok {
			fmt.Println(arr)
			printNode(rNode)
			break
		}
	}
	fmt.Println("done")
}

type MyQueue struct {
	head *ListNode
	tail *ListNode
	size int
}

func (q *MyQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *MyQueue) Size() int {
	return q.size
}

func (q *MyQueue) Offer(val int) {
	node := &ListNode{Val: val}
	if q.head == nil {
		q.head = node
	} else {
		q.tail.Next = node
	}
	q.tail = node
	q.size = q.size + 1
}

func (q *MyQueue) Poll() (int, error) {
	if q.head == nil {
		return -1, errors.New("queue node is nil")
	}
	curNode := q.head
	val := curNode.Val
	q.head = curNode.Next
	if q.head == nil {
		q.tail = nil
	}
	return val, nil
}

func (q *MyQueue) Peek() (int, error) {
	if q.head == nil {
		return -1, errors.New("queue node is nil")
	}
	return q.head.Val, nil
}

func TestQueue(t *testing.T) {
	fmt.Println("start")
	q := new(MyQueue)
	q.Offer(3)
	q.Offer(5)
	q.Offer(4)
	printNode(q.head)
	val, _ := q.Poll()
	fmt.Println(val)
	val, _ = q.Peek()
	fmt.Println(val)
	printNode(q.head)
	fmt.Println("done")
}

type MyStack struct {
	head *ListNode
	size int
}

func (s *MyStack) IsEmpty() bool {
	return s.size == 0
}

func (s *MyStack) Size() int {
	return s.size
}

func (s *MyStack) Push(val int) {
	node := &ListNode{Val: val}
	node.Next = s.head
	s.head = node
	s.size = s.size + 1
}

func (s *MyStack) Pop() (int, error) {
	if s.head == nil {
		return -1, errors.New("stack node is nil")
	}
	curNode := s.head
	val := curNode.Val
	s.head = curNode.Next
	return val, nil
}

func (s *MyStack) Peek() (int, error) {
	if s.head == nil {
		return -1, errors.New("queue node is nil")
	}
	return s.head.Val, nil
}

func TestStack(t *testing.T) {
	fmt.Println("start")
	s := new(MyStack)
	s.Push(3)
	s.Push(5)
	s.Push(6)
	printNode(s.head)
	val, _ := s.Pop()
	fmt.Println(val)
	val, _ = s.Peek()
	fmt.Println(val)
	printNode(s.head)
	fmt.Println("done")
}

type DoubleListNode struct {
	Val  int
	Next *DoubleListNode
	Last *DoubleListNode
}

type MyDeque struct {
	head *DoubleListNode
	tail *DoubleListNode
	size int
}

func (d *MyDeque) IsEmpty() bool {
	return d.size == 0
}

func (d *MyDeque) Size() int {
	return d.size
}

func (d *MyDeque) PushHead(val int) {
	node := &DoubleListNode{Val: val}
	if d.head == nil {
		d.tail = node
	} else {
		d.head.Last = node
		node.Next = d.head
	}
	d.head = node
	d.size = d.size + 1
}

func (d *MyDeque) PushTail(val int) {
	node := &DoubleListNode{Val: val}
	if d.head == nil {
		d.head = node
	} else {
		d.tail.Next = node
		node.Last = d.tail
	}
	d.tail = node
	d.size = d.size + 1
}

func (d *MyDeque) PopHead() (int, error) {
	if d.head == nil {
		return -1, errors.New("deque node is nil")
	}
	curNode := d.head
	d.head = curNode.Next
	if d.head == nil {
		d.tail = nil
	} else {
		d.head.Last = nil
	}
	return curNode.Val, nil
}

func (d *MyDeque) PopTail() (int, error) {
	if d.head == nil {
		return -1, errors.New("deque node is nil")
	}
	curNode := d.tail
	d.tail = curNode.Last
	if d.tail == nil {
		d.head = nil
	} else {
		d.tail.Next = nil
	}
	return curNode.Val, nil
}

func (d *MyDeque) PeekHead() (int, error) {
	if d.head == nil {
		return -1, errors.New("deque node is nil")
	}

	return d.head.Val, nil
}

func (d *MyDeque) PeekTail() (int, error) {
	if d.head == nil {
		return -1, errors.New("deque node is nil")
	}

	return d.tail.Val, nil
}

func printDeNode(node *DoubleListNode, revertFlag int) {
	if node == nil {
		fmt.Println("")
		return
	}
	currNode := node
	for currNode != nil {
		fmt.Print(currNode.Val)
		fmt.Print(" ")
		if revertFlag > 0 {
			currNode = currNode.Next
		} else {
			currNode = currNode.Last
		}

	}
	fmt.Println("")
}

func TestDeque(t *testing.T) {
	fmt.Println("start")
	q := new(MyDeque)
	q.PushHead(3)
	q.PushHead(5)
	q.PushTail(6)
	q.PushTail(8)
	printDeNode(q.head, 1)
	printDeNode(q.tail, -1)
	val, _ := q.PopTail()
	fmt.Println(val)
	val, _ = q.PeekTail()
	fmt.Println(val)
	printDeNode(q.head, 1)
	val, _ = q.PopHead()
	fmt.Println(val)
	val, _ = q.PeekHead()
	fmt.Println(val)
	printDeNode(q.head, 1)
	fmt.Println("done")
}

func getKGroupEnd(node *ListNode, k int) *ListNode {
	if node == nil {
		return nil
	}
	curNode := node
	count := k
	for curNode != nil {
		count--
		if count == 0 {
			return curNode
		}
		curNode = curNode.Next
	}

	return curNode
}

func reverseRangeGroup(start, end *ListNode) *ListNode {
	if end == nil {
		return start
	}
	var pre *ListNode
	cur := start
	for {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
		if pre == end {
			return pre
		}
	}

}

func reverseKGroup(node *ListNode, k int) *ListNode {
	var groupStartNode *ListNode
	rangeEndNode := getKGroupEnd(node, k)
	if rangeEndNode == nil {
		return node
	}
	groupStartNode = rangeEndNode.Next

	head := reverseRangeGroup(node, rangeEndNode)
	groupEndNode := node
	for groupStartNode != nil {

		rangeEndNode = getKGroupEnd(groupStartNode, k)
		var nextNode *ListNode
		if rangeEndNode != nil {
			nextNode = rangeEndNode.Next
		}
		rangeHeadNode := reverseRangeGroup(groupStartNode, rangeEndNode)
		groupEndNode.Next = rangeHeadNode
		groupEndNode = groupStartNode
		groupStartNode = nextNode

	}

	return head
}

func TestKReverse(t *testing.T) {
	head := &ListNode{Val: 1}
	curNode := head
	for i := 2; i <= 14; i++ {
		curNode.Next = &ListNode{Val: i}
		curNode = curNode.Next
	}
	printNode(head)
	node := reverseKGroup(head, 3)
	printNode(node)
}

func addNode(node1, node2 *ListNode) *ListNode {
	node1Size := size(node1)
	node2Size := size(node2)
	var shortNode, longNode *ListNode
	if node1Size >= node2Size {
		shortNode = node2
		longNode = node1
	} else {
		shortNode = node1
		longNode = node2
	}
	curSNode := shortNode
	curLNode := longNode
	var carry int = 0
	for curSNode != nil {
		curNum := curSNode.Val + curLNode.Val + carry
		carry = curNum / 10
		curLNode.Val = curNum % 10
		curSNode = curSNode.Next
		curLNode = curLNode.Next
	}
	preNode := curLNode
	for curLNode != nil {
		curNum := curLNode.Val + carry
		carry = curNum / 10
		curLNode.Val = curNum % 10
		preNode = curLNode
		curLNode = curLNode.Next
	}
	if carry > 0 {
		preNode.Next = &ListNode{Val: 1}
	}
	return longNode
}

func size(node *ListNode) int {
	size := 0
	curNode := node
	for curNode != nil {
		curNode = curNode.Next
		size++
	}
	return size
}

func TestAddNode(t *testing.T) {
	nodeS := &ListNode{Val: 2}
	nodeS.Next = &ListNode{Val: 6}

	nodeL := &ListNode{Val: 7}
	nodeL.Next = &ListNode{Val: 8}
	nodeL.Next.Next = &ListNode{Val: 9}

	node := addNode(nodeS, nodeL)
	printNode(node)
}

func mergeNode(node1, node2 *ListNode) *ListNode {
	var head, cur2 *ListNode
	if node1.Val <= node2.Val {
		head = node1
		cur2 = node2
	} else {
		head = node2
		cur2 = node1
	}
	cur1 := head.Next
	pre := head

	for cur1 != nil && cur2 != nil {
		if cur1.Val <= cur2.Val {
			pre.Next = cur1
			cur1 = cur1.Next
		} else {
			pre.Next = cur2
			cur2 = cur2.Next
		}
		pre = pre.Next
	}
	if cur1 != nil {
		pre.Next = cur1
	}
	if cur2 != nil {
		pre.Next = cur2
	}
	return head
}

func TestMergeNode(t *testing.T) {
	node1 := &ListNode{Val: 2}
	node1.Next = &ListNode{Val: 6}

	node2 := &ListNode{Val: 3}
	node2.Next = &ListNode{Val: 5}
	node2.Next.Next = &ListNode{Val: 9}

	node := mergeNode(node1, node2)
	printNode(node)
}
