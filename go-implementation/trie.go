package main

type node struct {
	children map[rune]*node
	eow      bool
}

// Create a new node
func newNode() *node {
	return &node{
		children: make(map[rune]*node),
		eow:      false,
	}
}

// Insert a new node to a trie
func (n *node) insert(key string) {
	for _, r := range key {
		if _, ok := n.children[r]; !ok {
			n.children[r] = newNode()
		}

		n = n.children[r]
	}

	n.eow = true
}

// Search for all nodes that matches the key prefix
func (n *node) search(key string) []string {
	for _, r := range key {
		if _, ok := n.children[r]; !ok {
			return nil
		}

		n = n.children[r]
	}

	if n.eow {
		return append(n.traverse(key), key)
	}

	return n.traverse(key)
}

// Traverse through the trie and return all its nodes with a prefix
func (n *node) traverse(prefix string) (results []string) {
	for key, node := range n.children {
		if node.eow {
			results = append(results, prefix+string(key))
		}

		results = append(results, node.traverse(prefix+string(key))...)
	}

	return results
}
