/*
 add by jinchaozhang. 2021/5/26
*/
package alg_test_test

import (
	"fmt"
	"math"
	"testing"
)

type TreeNode struct {
	val int ;
	left *TreeNode;
	right *TreeNode;
}

func TestSymmetricTree(t *testing.T) {
	tr := &TreeNode{val: 1}
	tr.left = &TreeNode{val: 2}
	tr.left.right = &TreeNode{val: 3}
	tr.right = &TreeNode{val: 2}
	tr.right.left = &TreeNode{val: 3}
	testOk := isMirror(tr,tr)
	fmt.Println(testOk)
}

func isMirror(t1 *TreeNode, t2 *TreeNode) bool{

	if (t1 == nil) != (t2 == nil) {
		return false;
	}
	if (t1 == nil && t2 == nil) {
		return true;
	}
	return t1.val == t2.val && isMirror(t1.left, t2.right) && isMirror(t1.right, t2.left);
}

func TestTreeDept(t *testing.T) {
	tr := &TreeNode{val: 1}
	tr.left = &TreeNode{val: 2}
	tr.left.right = &TreeNode{val: 3}
	tr.right = &TreeNode{val: 2}
	tr.right.left = &TreeNode{val: 3}
	dept := maxDepth(tr)
	fmt.Println(dept)
}

func maxDepth(t *TreeNode) int{
	if t == nil{
		return 0
	}
	return int(math.Max(float64(maxDepth(t.left)),float64(maxDepth(t.right)))) + 1
}

func TestBuildTreeWithPreOrderAndInOrder(t *testing.T) {
	preOrder := []int{1,2,4,3,5}
	inOrder := []int{4,2,1,5,3}
	th := buildTree(preOrder,inOrder)
	fmt.Println("----",th.val)
	fmt.Println("--",th.left.val,"--",th.right.val)
	fmt.Println(th.left.left.val,"------",th.right.left.val)
}

func buildTree(pre,in []int) *TreeNode{
	if len(pre) != len(in){
		return nil
	}
	valueIndexMap := make(map[int]int,0)
	for index,val := range in{
		valueIndexMap[val] = index
	}
	return doBuildTree(pre,in,0,len(pre)- 1,0,len(in)-1,valueIndexMap)
}

func doBuildTree(pre,in []int,pl,pr,il,ir int,inValueIndexMap map[int]int)*TreeNode{
	if pl > pr{
		return nil
	}
	head := &TreeNode{val: pre[pl]}
	if pl == pr{
		return head
	}
	find := inValueIndexMap[pre[pl]]
	head.left = doBuildTree(pre,in,pl + 1,pl + find - il,il,find - 1,inValueIndexMap)
	head.right = doBuildTree(pre,in,pl + find - il + 1,pr,find + 1,ir,inValueIndexMap)
	return head
}


