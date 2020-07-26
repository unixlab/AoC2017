package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	Name   string
	Weight int
	Parent *node
}

func main() {
	var lines []string
	var nodes []node

	nodeIndex := make(map[string]*node, 100)

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for _, line := range lines {
		whitespacePos := strings.Index(line, " ")
		openBracketPos := strings.Index(line, "(")
		closeBracketPos := strings.Index(line, ")")

		nodeName := line[:whitespacePos]
		nodeWeight, _ := strconv.Atoi(line[openBracketPos+1 : closeBracketPos])

		nodes = append(nodes, node{nodeName, nodeWeight, nil})
	}

	for pos, node := range nodes {
		nodeIndex[node.Name] = &nodes[pos]
	}

	for _, line := range lines {
		pos := strings.Index(line, ">")
		if pos > 0 {
			pos += 2

			whitespacePos := strings.Index(line, " ")
			nodeName := line[:whitespacePos]
			nodeAddr := nodeIndex[nodeName]

			for _, childNode := range strings.Split(line[pos:], ", ") {
				nodeIndex[childNode].Parent = nodeAddr
			}
		}
	}

	for _, node := range nodes {
		if node.Parent == nil {
			fmt.Println(node.Name)
		}
	}
}
