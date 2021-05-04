package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(rs []Record) (*Node, error) {
	if len(rs) == 0 {
		return nil, nil
	}

	idToNodeMap := make(map[int]*Node, len(rs))
	idToRecMap := make(map[int]Record, len(rs))
	sortedKeys := make([]int, len(rs))

	for k, r := range rs {
		if r.ID != 0 && r.ID == r.Parent {
			return nil, errors.New("parent cannot be equal to its own id")
		}

		if r.ID < r.Parent {
			return nil, errors.New("ParentID can't be greater than child id")
		}

		if _, ok := idToNodeMap[r.ID]; ok {
			return nil, errors.New("there cannot be duplicate nodes")
		}

		idToNodeMap[r.ID] = &Node{ID: r.ID}
		idToRecMap[r.ID] = r
		sortedKeys[k] = r.ID
	}

	if _, ok := idToNodeMap[0]; !ok {
		return nil, errors.New("no root node")
	}

	sort.Ints(sortedKeys)

	if len(rs) != sortedKeys[len(sortedKeys)-1]+1 {
		return nil, errors.New("there is a gap in node ids")
	}

	for _, id := range sortedKeys {
		if idToRecMap[id].ID == 0 {
			continue
		}

		idToNodeMap[idToRecMap[id].Parent].Children = append(idToNodeMap[idToRecMap[id].Parent].Children, idToNodeMap[id])
	}

	return idToNodeMap[0], nil
}
