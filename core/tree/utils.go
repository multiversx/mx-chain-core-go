package tree

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
