package studyhall

type Node struct {
	key int
	value string
	leftNode *Node
	rightNode *Node
}

type BST struct {
	root *Node
}

func (bst *BST) Insert(key int, value string) {
	newNode := &Node{
		key: key,
		value: value,
	}

	if bst.root == nil {
		bst.root = newNode
	} else {
		insertNode(bst.root, newNode)
	}
}

func insertNode(node, newNode *Node) {
	if newNode.key < node.key {
		if node.leftNode == nil {
			node.leftNode = newNode
		} else {
			insertNode(node.leftNode, newNode)
		}
	} else {
		if node.rightNode == nil {
			node.rightNode = newNode
		} else {
			insertNode(node.rightNode, newNode)
		}
	}
}

func (bst *BST) Search(key int) string {
	if found, val := search(bst.root, key); found {
		return val.value
	}

	return ""
}

func search(n *Node, key int) (bool, *Node) {
	if n == nil {
		return false, nil
	}
	if key < n.key {
		return search(n.leftNode, key)
	}
	if key > n.key {
		return search(n.rightNode, key)
	}

	return true, n
}

func (bst *BST) Delete(key int) {
	found, n := search(bst.root, key)
	if !found {
		return
	}

	if n.leftNode == nil && n.rightNode == nil {
		n = nil
		return
	}

	if n.leftNode == nil {
		n = n.rightNode
		return
	}

	if n.rightNode == nil {
		n = n.leftNode
		return
	}

	nleftNode := n.leftNode
	n = n.rightNode

	insertNode(bst.root, nleftNode)
}

