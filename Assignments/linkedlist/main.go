package main

func main() {
	// OpsOnNode()
	rootNode := NewNode(5)
	AddElement(rootNode, 10)
	AddElement(rootNode, 20)
	AddElement(rootNode, 25)

	//Should print 5 -> 10 -> 20 -> 25
	TraverseAndPrintNodeValues(rootNode)

}
