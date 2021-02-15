package tree

type Node struct {
	key   int
	left  *Node
	right *Node
}

func NewNode(value int) *Node {
	return &Node{key: value}
}

type BinarySearch struct {
	root *Node
}

func NewBinary() *BinarySearch {
	return &BinarySearch{}
}

func (b *BinarySearch) Insert(key int) {
	var (
		newNode = NewNode(key)
	)
	if b.root == nil {
		b.root = newNode
		return
	} else {
		b.insertNode(b.root, newNode)
	}
}

func (b *BinarySearch) Search(key int) bool {
	return b.searchNode(b.root, key)
}

func (b *BinarySearch) Remove(key int) {
	b.removeNode(b.root, key)
}

func (b *BinarySearch) insertNode(node, newNode *Node) {
	if newNode == nil {
		return
	}

	if node == nil {
		return
	}

	if newNode.key < node.key {
		if node.left == nil {
			node.left = newNode
		} else {
			b.insertNode(node.left, newNode)
		}
	}

	if newNode.key > node.key {
		if node.right == nil {
			node.right = newNode
		} else {
			b.insertNode(node.right, newNode)
		}
	}
}

func (b *BinarySearch) searchNode(node *Node, key int) bool {
	if node == nil {
		return false
	}

	if key < node.key {
		return b.searchNode(node.left, key)
	}

	if key > node.key {
		return b.searchNode(node.right, key)
	}
	return true
}

func (b *BinarySearch) removeNode(node *Node, key int) *Node {
	if node == nil {
		return nil
	}

	if key < node.key {
		node.left = b.removeNode(node.left, key)
		return node
	}
	if key > node.key {
		node.right = b.removeNode(node.right, key)
		return node
	}

	if node.left == nil && node.right == nil {
		node = nil
		return node
	}

	if node.left == nil {
		node = node.right
		return node
	}

	if node.right == nil {
		node = node.left
		return node
	}

	// node has two children, take inorderNextGreater value => choose the smallest key that is bigger than removed key
	inOrderSuccessor := node.right
	for inOrderSuccessor.left != nil {
		inOrderSuccessor = inOrderSuccessor.left
	}

	node.key = inOrderSuccessor.key
	node.right = b.removeNode(node.right, inOrderSuccessor.key)
	return node
}
