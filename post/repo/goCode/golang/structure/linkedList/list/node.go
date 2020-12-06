package list

type Node struct{
	Value int
	Next *Node
}

type bidiNode struct{
	Pre *bidiNode
	Value int
	Next *bidiNode
}