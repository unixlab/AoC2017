package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	Name     string
	Weight   int
	Parent   *node
	Children []*node
}

func sumNode(node *node) int {
	sum := 0
	if node.Children != nil {
		for _, c := range node.Children {
			sum += sumNode(c)
		}
	}
	sum += node.Weight
	return sum
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

		nodes = append(nodes, node{nodeName, nodeWeight, nil, nil})
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
				nodeAddr.Children = append(nodeAddr.Children, nodeIndex[childNode])
			}
		}
	}

	var root *node

	for nodePos, node := range nodes {
		if node.Parent == nil {
			root = &nodes[nodePos]
		}
	}

	fmt.Println(root.Name)

	foundFault := false
	for {
		dist := make(map[int]int, 2)
		weights := make(map[*node]int, len(root.Children))
		for _, c := range root.Children {
			sum := sumNode(c)
			weights[c] = sum
			dist[sum]++
		}

		outlier := 0
		for sum, counter := range dist {
			if counter == 1 {
				outlier = sum
			}
		}

		if foundFault {
			var numbers [2]int
			for k, v := range dist {
				numbers[v-1] = k
			}
			adjust := numbers[1] - numbers[0]
			for node, weight := range weights {
				if weight == outlier {
					fmt.Println(node.Weight + adjust)
				}
			}
			break
		}

		if outlier == 0 {
			root = root.Parent
			foundFault = true
		} else {
			for node, weight := range weights {
				if weight == outlier {
					root = node
				}
			}
		}
	}
}
