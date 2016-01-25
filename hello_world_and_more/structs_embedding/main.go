package main

import "fmt"

// START OMIT
type node struct {
	name string
}

type superNode struct {
	node
}

type cluster struct {
	name string
	node node
}

func main() {
	n := node{name: "Primary Node"}
	sn := superNode{n}
	c := cluster{name: "Servers", node: n}

	fmt.Printf("SuperNode: %+v, %s\n", sn, sn.name) // SuperNode IS a node
	fmt.Printf("Cluster:   %+v, %s\n", c, c.name)   // Cluster HAS a node
}

// END OMIT
