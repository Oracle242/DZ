package main

import (
	"fmt"
	"operation/BinaryTree"
)

func main() {
	tree := BinaryTree.BinaryTree{}

	// Вставка элементов в дерево
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(8)
	tree.Insert(2)
	tree.Insert(4)

	// Печать элементов дерева
	tree.PrintInOrder() // Вывод: 2 3 4 5 8

	// Поиск элемента в дереве
	fmt.Println(tree.Search(4)) // Вывод: true

	// Удаление элемента из дерева
	tree.Delete(3)
	tree.PrintInOrder() // Вывод: 2 4 5 8
}
