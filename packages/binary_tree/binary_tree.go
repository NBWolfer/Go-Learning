package packages

import "fmt"

// Node represents a node in the binary tree
type Node struct {
	data  int
	left  *Node
	right *Node
}

// BinaryTree represents a binary tree
type BinaryTree struct {
	root *Node
}

// Insert inserts a new node in the binary tree recursively
func (b *BinaryTree) Insert(data int) {
	if b.root == nil {
		b.root = &Node{data: data}
		return
	}
	b.insert(b.root, data)
}

func (b *BinaryTree) insert(node *Node, data int) {
	if data < node.data {
		if node.left == nil {
			node.left = &Node{data: data}
		} else {
			b.insert(node.left, data)
		}
	} else {
		if node.right == nil {
			node.right = &Node{data: data}
		} else {
			b.insert(node.right, data)
		}
	}
}

// InOrderTraversal traverses the binary tree in InOrder fashion
func (b *BinaryTree) InOrderTraversal(node *Node) {
	if node == nil {
		return
	}
	b.InOrderTraversal(node.left)
	fmt.Printf("%d ", node.data)
	b.InOrderTraversal(node.right)
}

// PreOrderTraversal traverses the binary tree in PreOrder fashion
func (b *BinaryTree) PreOrderTraversal(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.data)
	b.PreOrderTraversal(node.left)
	b.PreOrderTraversal(node.right)
}

// PostOrderTraversal traverses the binary tree in PostOrder fashion
func (b *BinaryTree) PostOrderTraversal(node *Node) {
	if node == nil {
		return
	}
	b.PostOrderTraversal(node.left)
	b.PostOrderTraversal(node.right)
	fmt.Printf("%d ", node.data)
}

// BFS traverses the binary tree in Breadth First Searching and prints the nodes
func (b *BinaryTree) BFS() {
	if b.root == nil {
		return
	}
	queue := []*Node{b.root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", node.data)
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
}

// DFS traverses the binary tree in Depth First Searching and prints the nodes
func (b *BinaryTree) DFS() {
	if b.root == nil {
		return
	}
	stack := []*Node{b.root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Printf("%d ", node.data)
		if node.right != nil {
			stack = append(stack, node.right)
		}
		if node.left != nil {
			stack = append(stack, node.left)
		}
	}
}

// Delete deletes a node from the binary tree
func (b *BinaryTree) Delete(data int) {
	b.root = b.delete(b.root, data)
}

func (b *BinaryTree) delete(node *Node, data int) *Node {
	if node == nil {
		return nil
	}
	if data < node.data {
		node.left = b.delete(node.left, data)
	} else if data > node.data {
		node.right = b.delete(node.right, data)
	} else {
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}
		node.data = b.minValue(node.right)
		node.right = b.delete(node.right, node.data)
	}
	return node
}

func (b *BinaryTree) minValue(node *Node) int {
	minValue := node.data
	for node.left != nil {
		minValue = node.left.data
		node = node.left
	}
	return minValue
}

// Depth returns the depth of the binary tree
func (b *BinaryTree) Depth() int {
	return b.depth(b.root)
}

func (b *BinaryTree) depth(node *Node) int {
	if node == nil {
		return 0
	}
	leftDepth := b.depth(node.left)
	rightDepth := b.depth(node.right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

// BinaryTreeDemo demonstrates the binary tree operations
func BinaryTreeDemo() {
	b := BinaryTree{}
	b.Insert(10)
	b.Insert(6)
	b.Insert(15)
	b.Insert(3)
	b.Insert(8)
	b.Insert(20)
	b.Insert(6) // Duplicate value
	b.Insert(12)
	b.Insert(18)
	b.Insert(22)

	fmt.Println("Inserted nodes in the binary tree: 10, 6, 15, 3, 8, 20 and 6, 12, 18, 22")

	fmt.Println("InOrder Traversal:")
	b.InOrderTraversal(b.root)
	fmt.Println()

	fmt.Println("PreOrder Traversal:")
	b.PreOrderTraversal(b.root)
	fmt.Println()

	fmt.Println("PostOrder Traversal:")
	b.PostOrderTraversal(b.root)
	fmt.Println()

	fmt.Println("BFS:")
	b.BFS()
	fmt.Println()

	fmt.Println("DFS:")
	b.DFS()
	fmt.Println()

	fmt.Println("Depth of the binary tree:", b.Depth())

	fmt.Println("Binary Tree:")
	fmt.Println("    10")
	fmt.Println("   /  \\")
	fmt.Println("  6    15")
	fmt.Println(" / \\    / \\")
	fmt.Println("3   8  12  20")
	fmt.Println("       / \\    \\")
	fmt.Println("      6  18   22")

	fmt.Println()

	fmt.Println("Deleting 6 from the binary tree")
	b.Delete(6)

	fmt.Println("InOrder Traversal:")
	b.InOrderTraversal(b.root)
	fmt.Println()

	fmt.Println("Depth of the binary tree:", b.Depth())

	fmt.Println("Binary Tree:")
	fmt.Println("    10")
	fmt.Println("   /  \\")
	fmt.Println("  8    15")
	fmt.Println(" /    /  \\")
	fmt.Println("3    12   20")
	fmt.Println("        \\   \\")
	fmt.Println("        18   22")
}
