package main

import "fmt"

type Node struct {
	Value int64
	Next  *Node
}

func NewNode(v int64) *Node {
	return &Node{Value: v}
}

func OpsOnNode() {
	node := NewNode(5)
	node.Next = NewNode(10)
}

func AddElement(rootNode *Node, value int64) {
	//Traverse the node and add the value at the end
	for rootNode.Next != nil {
		rootNode = rootNode.Next
	}
	rootNode.Next = NewNode(value)
}
func TraverseAndPrintNodeValues(rootNode *Node) {
	//Traverse and print each value
	//temp pointing to root. Then you traverse till Next pointer is not nil
	temp := rootNode
	for temp != nil {
		fmt.Printf("%v -> ", temp.Value)
		temp = temp.Next
	}

}

func LinkedListOps() {
	// OpsOnNode()
	rootNode := NewNode(5)
	AddElement(rootNode, 10)
	AddElement(rootNode, 20)
	AddElement(rootNode, 25)

	//Should print 5 -> 10 -> 20 -> 25
	TraverseAndPrintNodeValues(rootNode)
}
