package main

import (
	"fmt"
	"github.com/emicklei/dot"
)

type Node struct {
	Parent *Node
	Right  *Node
	Left   *Node
	Value int
}

type Queue []*Node

// Dequeue takes a queue and return
// the queue without it's head element
// and return it's head element.
func dequeue(q Queue)(Queue, *Node) {
	if len(q) > 0 {
		head := q[0]
		q = q[1:]
		return q, head
	}
	return nil, nil
}

// appendChild takes a int value of the child, it make
// the child node and add with a referent to it's parent
// It append nodes first to left and then to right
func (n *Node) appendChild(childValue int) (*Node, bool) {
	child := &Node{n, nil, nil, childValue}
	if n.Left == nil {
		n.Left = child
	} else if n.Right == nil {
		n.Right = child
	}else {
		return child, false
	}
	return child, true
}

// Create the DOT string
func printGraph(root *Node){
	g := dot.NewGraph(dot.Directed)
	printRecursive(root, g)
	fmt.Println(g.String())
}

// Recursive function to traverse the tree
func printRecursive(root *Node, g *dot.Graph){
	if root==nil{
		return
	}
	n1 := g.Node(fmt.Sprintf("%d", root.Value))
	if (root.Left != nil){
		n2 := g.Node(fmt.Sprintf("%d", root.Left.Value))
		g.Edge(n1,n2)
	}
	if (root.Right != nil){
		n3 := g.Node(fmt.Sprintf("%d", root.Right.Value))
		g.Edge(n1,n3)
	}
	printRecursive(root.Left, g)
	printRecursive(root.Right, g)
}


func main() {
	// Visited is the closed set, visited[nodeValue] == true means it
	// is in the closed set
	visited := make(map[int]bool)
	// Open list is a Queue for implementing breadth-first search
	var open Queue

	first := &Node{nil, nil, nil, 1}
	open = append(open, first)
	var node *Node


	for open!=nil {
		open, node = dequeue(open)

		if node.Value == 100 {
			break
		}

		// Search the successors of the current node
		// calculating the sum with it's parents value.
		itNode := node.Parent
		for itNode != nil {
			childValue := itNode.Value + node.Value
			if childValue <=100 && !visited[childValue]{
				child, success := node.appendChild(childValue)
				if success {
					open = append(open, child)
					visited[childValue] = true
				}
			}
			itNode = itNode.Parent
		}
		// If it has only one successor or none of it, append the sum
		// with itself (if it's not in closed set)
		if (node.Left == nil || node.Right == nil) && node.Value*2 <= 100 && !visited[node.Value*2]{
			child, _ := node.appendChild(node.Value*2)
			open = append(open, child)
			visited[node.Value*2] = true
		}
	}
	printGraph(first)
}
