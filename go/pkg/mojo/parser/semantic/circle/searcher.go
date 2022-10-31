package circle

import (
	"fmt"
	"strings"
)

type fileNode struct {
	name         string
	dependencies []string

	depNodes      []*fileNode
	dependedNodes []*fileNode

	circleNodes []*fileNode
}

func (f *fileNode) IsCircle() bool {
	return len(f.circleNodes) > 0
}

type Searcher struct {
}

func NewSearcher() *Searcher {
	return &Searcher{}
}

func (s Searcher) normalizeNodes(files map[string]*fileNode) map[string]*fileNode {
	for _, node := range files {
		for _, dep := range node.dependencies {
			depNode := files[dep]
			if depNode == nil {
				// error happened
				return map[string]*fileNode{}
			}
			node.depNodes = append(node.depNodes, depNode)
			depNode.dependedNodes = append(depNode.dependedNodes, node)
		}
	}
	return files
}

func (s Searcher) cleanNodes(files map[string]*fileNode) map[string]*fileNode {
	if len(files) == 0 {
		return files
	}

	fileNodes := make(map[string]*fileNode)
	independentNodes := make(map[string]*fileNode)

	for key, node := range files {
		if len(node.depNodes) > 0 && len(node.dependedNodes) > 0 {
			fileNodes[key] = node
		} else {
			independentNodes[key] = node
		}
	}

	dirty := false
	for _, node := range fileNodes {
		var depNodes []*fileNode
		for _, n := range node.depNodes {
			if independentNodes[n.name] == nil {
				depNodes = append(depNodes, n)
			}
		}
		node.depNodes = depNodes

		var dependedNodes []*fileNode
		for _, n := range node.dependedNodes {
			if independentNodes[n.name] == nil {
				dependedNodes = append(dependedNodes, n)
			}
		}
		node.dependedNodes = dependedNodes

		if len(node.depNodes) == 0 || len(node.dependedNodes) == 0 {
			dirty = true
		}
	}

	if dirty {
		fileNodes = s.cleanNodes(fileNodes)
	}

	return fileNodes
}

func inStringArray(key string, values []string) bool {
	for _, v := range values {
		if v == key {
			return true
		}
	}
	return false
}

func checkFiles(files map[string]*fileNode) bool {
	for _, file := range files {
		for _, dep := range file.depNodes {
			if !strings.HasPrefix(file.name, "circle-") &&
				!strings.HasPrefix(dep.name, "circle-") &&
				!inStringArray(dep.name, file.dependencies) {
				return false
			}
		}

		for _, dep := range file.depNodes {
			found := false
			for _, depended := range dep.dependedNodes {
				if depended.name == file.name {
					found = true
				}
			}
			if !found {
				return false
			}
		}

		for _, depended := range file.dependedNodes {
			found := false
			for _, dep := range depended.depNodes {
				if dep.name == file.name {
					found = true
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}

func (s Searcher) Search(files map[string]*fileNode) [][]string {
	files = s.normalizeNodes(files)
	if len(files) <= 1 {
		return [][]string{}
	}

	fileNodes := s.cleanNodes(files)
	if len(fileNodes) <= 1 {
		return [][]string{}
	}

	index := 0
	var circleNodes []*fileNode
	for len(fileNodes) > 0 {
		for _, node := range fileNodes {
			visited := make(map[string]bool)
			track := make([]*fileNode, 0)
			found := searchCircle(node, visited, &track)
			if found {
				circle := makeCircle(track, index)

				for _, n := range track {
					delete(fileNodes, n.name)
				}
				fileNodes[circle.name] = circle
				index++
			} else {
				for _, n := range track {
					if n.IsCircle() {
						circleNodes = append(circleNodes, n)
					}
					delete(fileNodes, n.name)
				}
			}

			if !checkFiles(fileNodes) {
				fmt.Print("error")
			}

			break
		}
	}

	var circles [][]string
	for _, node := range circleNodes {
		circles = append(circles, flatCircle(node))
	}

	return circles
}

func flatCircle(circle *fileNode) []string {
	files := make([]string, 0)
	for _, node := range circle.circleNodes {
		if node.IsCircle() {
			fs := flatCircle(node)
			files = append(files, fs...)
		} else {
			files = append(files, node.name)
		}
	}
	return files
}

func makeCircle(track []*fileNode, index int) *fileNode {
	node := &fileNode{
		name:        fmt.Sprintf("circle-%d", index),
		depNodes:    nil,
		circleNodes: track,
	}

	// update the depended nodes
	circle := make(map[string]bool)
	for _, c := range node.circleNodes {
		circle[c.name] = true
	}

	// remove the duplicated dep nodes
	deps := make(map[string]*fileNode)
	dependedNodes := make(map[string]*fileNode)
	for _, n := range track {
		for _, d := range n.depNodes {
			if !circle[d.name] {
				deps[d.name] = d
			}
		}

		for _, d := range n.dependedNodes {
			if !circle[d.name] {
				dependedNodes[d.name] = d
			}
		}
	}

	for _, d := range deps {
		node.depNodes = append(node.depNodes, d)
	}
	for _, d := range dependedNodes {
		node.dependedNodes = append(node.dependedNodes, d)
	}

	for _, n := range node.dependedNodes {
		nodes := make([]*fileNode, 0, len(n.depNodes))
		included := false
		for _, d := range n.depNodes {
			if circle[d.name] {
				if !included {
					included = true
					nodes = append(nodes, node)
				}
			} else {
				nodes = append(nodes, d)
			}
		}
		n.depNodes = nodes
	}

	for _, n := range node.depNodes {
		nodes := make([]*fileNode, 0, len(n.dependedNodes))
		included := false
		for _, d := range n.dependedNodes {
			if circle[d.name] {
				if !included {
					included = true
					nodes = append(nodes, node)
				}
			} else {
				nodes = append(nodes, d)
			}
		}
		n.dependedNodes = nodes
	}

	return node
}

func searchCircle(node *fileNode, visited map[string]bool, track *[]*fileNode) bool {
	visited[node.name] = true
	*track = append(*track, node)
	for _, dependency := range node.depNodes {
		if visited[dependency.name] { // found circle
			// update the track to only circle related
			for i, n := range *track {
				if n.name == dependency.name {
					*track = (*track)[i:]
					break
				}
			}
			return true
		} else {
			found := searchCircle(dependency, visited, track)
			if found {
				return found
			} else {
				*track = (*track)[:len(*track)-1]
			}
		}
	}
	return false
}
