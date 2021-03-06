package _1_bi_tree

import (
	"errors"
)

/*
八
给定二叉树和其中一个节点，找到该节点在中序遍历序列的后继节点。
树中节点有指向父节点的指针。
*/

type BiTreeWithP struct {
	Data   interface{}
	lChild *BiTreeWithP
	rChild *BiTreeWithP
	parent *BiTreeWithP
}

/*
获取中序遍历的后继节点
主要有两种情况：该节点的右子节点不为nil，在右子树上找最左节点；该节点的右子节点为nil，寻找节点与父节点的一种相对位置。
*/
func getSuccessor(node *BiTreeWithP) *BiTreeWithP {
	getLeft := func(node *BiTreeWithP) *BiTreeWithP {
		if node == nil {
			return nil
		}
		for {
			if node.lChild != nil {
				node = node.lChild
			} else {
				return node
			}
		}
	}

	//参照寻找直接前驱的函数对比着看
	if node != nil {
		if node.rChild != nil { //第一种情况：节点的右子节点不是nil，那么就在右子树中寻找
			return getLeft(node.rChild)
		} else { //第二种情况：节点的右子节点是nil，那么一直判断节点与父节点的相对位置，逐级向上寻找
			for {
				if node == nil || node.parent == nil {
					break
				}
				if node == node.parent.lChild {
					return node.parent
				}
				node = node.parent
			}
		}
	}

	return nil
}

func constructBiTreeByPreIn(preOrder, inOrder []int) (*BiTreeWithP, error) {
	preLen := len(preOrder)
	inLen := len(inOrder)
	if preLen == 0 || inLen == 0 || preLen != inLen {
		return nil, errors.New("input invalid")
	}
	return constructBiTreeByPreInCore(preOrder, 0, preLen-1, inOrder, 0, inLen-1)
}

func constructBiTreeByPreInCore(preOrder []int, preStart, preEnd int,
	inOrder []int, inStart, inEnd int) (*BiTreeWithP, error) {
	node := &BiTreeWithP{
		Data: preOrder[preStart],
	}
	if preEnd == preStart && inEnd == inStart {
		if preOrder[preStart] == inOrder[inStart] {
			return node, nil
		} else {
			return nil, errors.New("input invalid")
		}
	}

	inOrderIndex := inStart
	for inOrderIndex <= inEnd {
		if inOrder[inOrderIndex] == preOrder[preStart] {
			break
		} else {
			inOrderIndex++
		}
	}
	if inOrderIndex > inEnd {
		return nil, errors.New("input invalid")
	}
	leftLen := inOrderIndex - inStart
	rightLen := inEnd - inOrderIndex
	if leftLen > 0 {
		left, err := constructBiTreeByPreInCore(preOrder, preStart+1, preStart+leftLen,
			inOrder, inStart, inOrderIndex-1)
		if err != nil {
			return nil, err
		} else {
			node.lChild = left
			left.parent = node
		}
	}
	if rightLen > 0 {
		right, err := constructBiTreeByPreInCore(preOrder, preStart+leftLen+1, preEnd,
			inOrder, inOrderIndex+1, inEnd)
		if err != nil {
			return nil, err
		} else {
			node.rChild = right
			right.parent = node
		}
	}
	return node, nil
}

// PreOrderTraverse 二叉树的前序遍历
func (bi *BiTreeWithP) PreOrderTraverse() (res []interface{}) {
	if bi == nil {
		return
	}
	res = append(res, bi.Data)
	res = append(res, bi.lChild.PreOrderTraverse()...)
	res = append(res, bi.rChild.PreOrderTraverse()...)
	return res
}

// InOrderTraverse 二叉树的中序遍历
func (bi *BiTreeWithP) InOrderTraverse() (res []interface{}) {
	if bi == nil {
		return
	}
	res = append(res, bi.lChild.InOrderTraverse()...)
	res = append(res, bi.Data)
	res = append(res, bi.rChild.InOrderTraverse()...)
	return res
}

// PostOrderTraverse 二叉树的后序遍历
func (bi *BiTreeWithP) PostOrderTraverse() (res []interface{}) {
	if bi == nil {
		return
	}
	res = append(res, bi.lChild.PostOrderTraverse()...)
	res = append(res, bi.rChild.PostOrderTraverse()...)
	res = append(res, bi.Data)
	return res
}
