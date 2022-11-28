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

// IsLeftMargin returns true if the provided value is left margin of any node
func (tree *intervalTree) IsLeftMargin(value uint64) bool {
	return isLeftMargin(tree.root, value)
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

	intervals = sortAndOptimizeIntervals(intervals)

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

	nextNode.parentOffset = currentNode.offset
	if nextNode.low() <= currentNode.low() {
		if currentNode.left == nil {
			nextNode.offset = currentNode.offset - 1
			currentNode.left = nextNode
			return
		}

		tree.addNode(currentNode.left, nextNode)
		return
	}

	if currentNode.right == nil {
		nextNode.offset = currentNode.offset + 1
		currentNode.right = nextNode
		return
	}

	tree.addNode(currentNode.right, nextNode)
}

func contains(currentNode *node, value uint64) bool {
	if currentNode == nil {
		return false
	}

	if currentNode.contains(value) {
		return true
	}

	if currentNode.left != nil {
		if currentNode.left.max >= value {
			return contains(currentNode.left, value)
		}
	}

	return contains(currentNode.right, value)
}

func isLeftMargin(currentNode *node, value uint64) bool {
	if currentNode == nil {
		return false
	}

	if currentNode.low() == value {
		return true
	}

	if currentNode.low() >= value {
		return isLeftMargin(currentNode.left, value)
	}

	return isLeftMargin(currentNode.right, value)
}

func computeMaxFields(currentNode *node) {
	if currentNode == nil {
		return
	}

	if currentNode.isLeaf() {
		currentNode.max = currentNode.high()
		return
	}

	maxLeft := uint64(0)
	maxRight := uint64(0)
	if currentNode.left != nil {
		computeMaxFields(currentNode.left)
		maxLeft = currentNode.left.max
	}
	if currentNode.right != nil {
		computeMaxFields(currentNode.right)
		maxRight = currentNode.right.max
	}

	currentNode.max = max(currentNode.high(), max(maxLeft, maxRight))
}

func max(a, b uint64) uint64 {
	if a >= b {
		return a
	}
	return b
}

// IsInterfaceNil returns true if there is no value under the interface
func (tree *intervalTree) IsInterfaceNil() bool {
	return tree == nil
}
