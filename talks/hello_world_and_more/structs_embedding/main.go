package main

import "fmt"

// START OMIT
type node struct {
	name string
}

type altNode struct {
	node
}

type cluster struct {
	name string
	node node
}

func main() {
	n := node{name: "Primary"}
	a := altNode{n}
	c := cluster{name: "Nodes", node: n}

	fmt.Printf("node   : %+v, %s\n", n, n.name)
	fmt.Printf("altNode: %+v, %s\n", a, a.name) // altNode embeds a node
	fmt.Printf("cluster: %+v, %s\n", c, c.name) // Cluster contains a node field
}

// END OMIT
