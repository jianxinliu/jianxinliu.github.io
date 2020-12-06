package list

type list interface{
	init() *linkedList
	NewNode(value int) *Node
	Append(value int) *Node
	PreAppend(value int) *Node
	Size() (int)
	Remove(value int) error
	RemoveTail()
	Print()
}