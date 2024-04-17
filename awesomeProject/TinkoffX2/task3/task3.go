package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type node struct {
	Name     string
	Children []*node
}

func main() {
	var n int
	fmt.Scan(&n)

	scanner := bufio.NewScanner(os.Stdin)
	directories := make([]string, n)
	for i := 0; i < n && scanner.Scan(); i++ {
		directories[i] = scanner.Text()
	}

	root := &node{Name: ""}
	for _, directory := range directories {
		addChild(root, directory)
	}
	printTree(root, "")
}

func addChild(root *node, directory string) {
	parts := strings.Split(directory, "/")
	for i := 0; i < len(parts); i++ {
		dir := parts[i]
		found := false
		for _, child := range root.Children {
			if child.Name == dir {
				root = child
				found = true
				break
			}
		}
		if !found {
			newDir := &node{Name: dir}
			root.Children = append(root.Children, newDir)
			root = newDir
		}
	}
}

func printTree(node *node, indent string) {
	sort.Slice(node.Children, func(i, j int) bool {
		return node.Children[i].Name < node.Children[j].Name
	})

	for _, child := range node.Children {
		fmt.Println(indent + child.Name)
		printTree(child, indent+"  ")
	}
}
