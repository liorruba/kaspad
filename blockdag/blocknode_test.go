package blockdag

import (
	"testing"
)

func TestChainHeight(t *testing.T) {
	phantomK := uint32(2)
	buildNode := buildNodeGenerator(phantomK)
	buildWithChildren := func(parents blockSet) *blockNode {
		node := buildNode(parents)
		addNodeAsChildToParents(node)
		return node
	}

	node0 := buildWithChildren(setFromSlice())
	node1 := buildWithChildren(setFromSlice(node0))
	node2 := buildWithChildren(setFromSlice(node0))
	node3 := buildWithChildren(setFromSlice(node0))
	node4 := buildWithChildren(setFromSlice(node1, node2, node3))
	node5 := buildWithChildren(setFromSlice(node1, node2, node3))
	node6 := buildWithChildren(setFromSlice(node1, node2, node3))
	node7 := buildWithChildren(setFromSlice(node0))
	node8 := buildWithChildren(setFromSlice(node7))
	node9 := buildWithChildren(setFromSlice(node8))
	node10 := buildWithChildren(setFromSlice(node9, node6))

	// Because nodes 7 & 8 were mined secretly, node10's selected
	// parent will be node6, although node9 is higher. So in this
	// case, node10.height and node10.chainHeight will be different

	tests := []struct {
		node                *blockNode
		expectedChainHeight uint32
	}{
		{
			node:                node0,
			expectedChainHeight: 0,
		},
		{
			node:                node1,
			expectedChainHeight: 1,
		},
		{
			node:                node2,
			expectedChainHeight: 1,
		},
		{
			node:                node3,
			expectedChainHeight: 1,
		},
		{
			node:                node4,
			expectedChainHeight: 2,
		},
		{
			node:                node5,
			expectedChainHeight: 2,
		},
		{
			node:                node6,
			expectedChainHeight: 2,
		},
		{
			node:                node7,
			expectedChainHeight: 1,
		},
		{
			node:                node8,
			expectedChainHeight: 2,
		},
		{
			node:                node9,
			expectedChainHeight: 3,
		},
		{
			node:                node10,
			expectedChainHeight: 3,
		},
	}

	for _, test := range tests {
		if test.node.chainHeight != test.expectedChainHeight {
			t.Errorf("block %v expected chain height %v but got %v", test.node, test.expectedChainHeight, test.node.chainHeight)
		}
	}

}