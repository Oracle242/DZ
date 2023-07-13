package BinaryTree

import "fmt"

// Определение типа узла дерева
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Определение типа дерева
type BinaryTree struct {
	Root *Node
}

// Функция для вставки элемента в дерево
func (tree *BinaryTree) Insert(value int) {
	if tree.Root == nil {
		tree.Root = &Node{Value: value, Left: nil, Right: nil}
	} else {
		tree.Root.insert(value)
	}
}

// Вспомогательная функция для вставки элемента в узел
func (node *Node) insert(value int) {
	if value < node.Value {
		if node.Left == nil {
			node.Left = &Node{Value: value, Left: nil, Right: nil}
		} else {
			node.Left.insert(value)
		}
	} else {
		if node.Right == nil {
			node.Right = &Node{Value: value, Left: nil, Right: nil}
		} else {
			node.Right.insert(value)
		}
	}
}

// Функция для удаления элемента из дерева
func (tree *BinaryTree) Delete(value int) {
	if tree.Root == nil {
		return
	}
	tree.Root.delete(value)
}

// Вспомогательная функция для удаления элемента из узла
func (node *Node) delete(value int) *Node {
	if node == nil {
		return nil
	}
	if value < node.Value {
		node.Left = node.Left.delete(value)
	} else if value > node.Value {
		node.Right = node.Right.delete(value)
	} else {
		if node.Left == nil && node.Right == nil {
			return nil
		} else if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}
		minValue := node.Right.findMinValue()
		node.Value = minValue
		node.Right = node.Right.delete(minValue)
	}
	return node
}

// Вспомогательная функция для поиска минимального значения в дереве
func (node *Node) findMinValue() int {
	if node.Left == nil {
		return node.Value
	}
	return node.Left.findMinValue()
}

// Функция для поиска элемента в дереве
func (tree *BinaryTree) Search(value int) bool {
	if tree.Root == nil {
		return false
	}
	return tree.Root.search(value)
}

// Вспомогательная функция для поиска элемента в узле
func (node *Node) search(value int) bool {
	if node == nil {
		return false
	}
	if value < node.Value {
		return node.Left.search(value)
	} else if value > node.Value {
		return node.Right.search(value)
	}
	return true
}

// Функция для печати элементов дерева в порядке возрастания
func (tree *BinaryTree) PrintInOrder() {
	if tree.Root == nil {
		fmt.Println("Дерево пустое")
	} else {
		tree.Root.printInOrder()
	}
}

// Вспомогательная функция для печати элементов узла в порядке возрастания
func (node *Node) printInOrder() {
	if node != nil {
		node.Left.printInOrder()
		fmt.Println(node.Value)
		node.Right.printInOrder()
	}
}
