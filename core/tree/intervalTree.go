package tree

// BlocksExceptionInterval represents an exception config for an interval of rounds
type BlocksExceptionInterval struct {
	Low  uint64
	High uint64
}

type intervalTree struct {
	root *node
}

// NewIntervalTree returns a new instance of interval tree
func NewIntervalTree(intervals []BlocksExceptionInterval) *intervalTree {
	return createTree(intervals)
}

// Contains returns true if the provided value is part of any node
func (tree *intervalTree) Contains(value uint64) bool {
	return contains(tree.root, value)
}

// String returns a printable string form for the tree
func (tree *intervalTree) String() string {
	return treeToString(tree.root)
}

func createTree(intervals []BlocksExceptionInterval) *intervalTree {
	tree := &intervalTree{}
	if len(intervals) == 0 {
		return tree
	}

	intervals = sortIntervals(intervals)

	for _, blockExceptionInterval := range intervals {
		i := newInterval(blockExceptionInterval.Low, blockExceptionInterval.High)
		managedNode := newNode(i)
		tree.addNode(tree.root, managedNode)
	}

	computeMaxFields(tree.root)
	return tree
}

func (tree *intervalTree) addNode(currentNode *node, nextNode *node) {
	if currentNode == nil {
		tree.root = nextNode
		return
	}

	if nextNode.low() <= currentNode.low() {
		if currentNode.left == nil {
			currentNode.left = nextNode
			return
		}

		tree.addNode(currentNode.left, nextNode)
		return
	}

	if currentNode.right == nil {
		currentNode.right = nextNode
		return
	}

	tree.addNode(currentNode.right, nextNode)
}

// IsInterfaceNil returns true if there is no value under the interface
func (tree *intervalTree) IsInterfaceNil() bool {
	return tree == nil
}
