package tree

import (
	"strings"
)

const (
	newLine    = "\n"
	underscore = "_"
	space      = " "
	leftArrow  = "/"
	rightArrow = "\\"
)

type nodeDetails struct {
	*node
	lastEdgeParent *node
}

func treeToString(root *node) string {
	if root == nil {
		return ""
	}

	elementSize := len(root.toString())
	nodesOnLevels, minOffset := extractNodesOnLevelsWithOffsetsUpdated(root)
	arrowsLines := make([]string, len(nodesOnLevels))
	nodesLines := make([]string, len(nodesOnLevels))
	for level := 0; level < len(nodesOnLevels); level++ {
		lastEmptyOffset := 0
		for _, currentNode := range nodesOnLevels[level] {
			newOffset := currentNode.offset - minOffset
			distanceFromPrevious := newOffset - lastEmptyOffset
			emptySpace := strings.Repeat(space, distanceFromPrevious*elementSize)
			nodesLines[level] += emptySpace
			nodesLines[level] += currentNode.toString()
			lastEmptyOffset = newOffset + 1

			// skip arrow line for root
			if level == 0 {
				continue
			}

			arrowsLines[level] += computeArrowWithSpacing(currentNode, len(arrowsLines[level]), minOffset, elementSize)
		}
	}

	printableTreeStr := ""
	for i := 0; i < len(nodesLines); i++ {
		printableTreeStr += arrowsLines[i] + newLine
		printableTreeStr += nodesLines[i] + newLine
	}

	return printableTreeStr
}

func extractNodesOnLevelsWithOffsetsUpdated(root *node) ([][]*node, int) {
	maxOffset := 0
	minOffset := 0
	queue := make([]*nodeDetails, 0)
	rootDetails := &nodeDetails{
		node:           root,
		lastEdgeParent: root,
	}
	queue = append(queue, rootDetails)
	nodesOnLevels := make([][]*nodeDetails, 0)
	nodesOnLevels = append(nodesOnLevels, []*nodeDetails{rootDetails})

	// extract nodes through level order traversal
	for len(queue) > 0 {
		queueSize := len(queue)
		currentLevel := make([]*nodeDetails, 0)

		for i := 0; i < queueSize; i++ {
			currentNode := queue[0]
			queue = queue[1:]

			minOffset, maxOffset = updateOffsetsIfNeeded(currentNode.offset, minOffset, maxOffset)

			if currentNode.left != nil {
				nextNodeDetails := getNextLeftNodeUpdatingOffsets(currentNode, root, currentLevel)
				queue = append(queue, nextNodeDetails)
				currentLevel = append(currentLevel, nextNodeDetails)
			}
			if currentNode.right != nil {
				nextNodeDetails := getNextRightNode(currentNode, root)
				queue = append(queue, nextNodeDetails)
				currentLevel = append(currentLevel, nextNodeDetails)
			}
		}

		if len(currentLevel) > 0 {
			nodesOnLevels = append(nodesOnLevels, currentLevel)
		}
	}

	nodesSlices := make([][]*node, len(nodesOnLevels))
	for level, nodesDetailsOnLevel := range nodesOnLevels {
		nodesSlices[level] = make([]*node, 0, len(nodesDetailsOnLevel))
		for _, nodeDetailsOnLevel := range nodesDetailsOnLevel {
			nodesSlices[level] = append(nodesSlices[level], nodeDetailsOnLevel.node)
		}
	}

	return nodesSlices, minOffset
}

func getNextLeftNodeUpdatingOffsets(currentNode *nodeDetails, root *node, currentLevel []*nodeDetails) *nodeDetails {
	// if next left offset already exists, the offsets must be updated, thus all possible collisions must be checked
	isLeftSubtree := currentNode.left.low() <= root.low()
	collidedNode := getFirstCollidedNode(currentLevel, currentNode.left)
	if collidedNode != nil {
		shiftingDistance := collidedNode.offset - currentNode.left.offset + 1
		if isLeftSubtree {
			shiftLeftSubtree(currentNode, collidedNode, shiftingDistance)
		} else {
			shiftRightSubtree(currentNode, collidedNode, shiftingDistance, root)
		}
	}

	isMovingLeft := currentNode.left.low() <= currentNode.lastEdgeParent.low()
	lastEdgeParent := currentNode.lastEdgeParent
	// if new edge detected, update it
	if isLeftSubtree && isMovingLeft {
		lastEdgeParent = currentNode.left
	}

	return &nodeDetails{
		node:           currentNode.left,
		lastEdgeParent: lastEdgeParent,
	}
}

func shiftLeftSubtree(currentNode *nodeDetails, collidedNode *nodeDetails, shiftingDistance int) {
	// if the collided nodes are in the same subtree from the last known edge, find the last common node,
	// first shift the left subtree of the last known edge to left,
	// then shift the current node's subtree from last common to right
	if isSameSubtree(currentNode.left, collidedNode.node, currentNode.lastEdgeParent.low()) {
		lastCommonNode := getLastCommonNode(currentNode.left, collidedNode.node, currentNode.lastEdgeParent)
		shiftSubTreeOffsets(currentNode.lastEdgeParent, -shiftingDistance)
		shiftSubTreeOffsets(lastCommonNode.right, shiftingDistance)
		return
	}

	// if the nodes are in the same subtree, simply shift the left subtree of the last known edge to left
	shiftSubTreeOffsets(currentNode.lastEdgeParent.left, -shiftingDistance)
}

func shiftRightSubtree(currentNode *nodeDetails, collidedNode *nodeDetails, shiftingDistance int, root *node) {
	// if the collided nodes are in different subtrees from the root, simply shift the right subtree to right
	if !isSameSubtree(currentNode.left, collidedNode.node, root.low()) {
		shiftSubTreeOffsets(currentNode.lastEdgeParent, shiftingDistance)
		return
	}

	// if they are on the same subtree from root but different subtrees from last edge, find the last common node,
	// shift the right subtree of the last known edge to right,
	// then shift the current node's subtree from last common to right
	if isSameSubtree(currentNode.left, collidedNode.node, collidedNode.lastEdgeParent.low()) {
		lastCommonNode := getLastCommonNode(currentNode.left, collidedNode.node, collidedNode.lastEdgeParent)
		shiftSubTreeOffsets(currentNode.lastEdgeParent.right, shiftingDistance)
		shiftSubTreeOffsets(lastCommonNode.right, shiftingDistance)
		return
	}

	// if they are in different subtrees from last known edge, simply shift the last known edge to right
	shiftSubTreeOffsets(currentNode.lastEdgeParent, shiftingDistance)
}

func getLastCommonNode(currentNode *node, collidedNode *node, lastCommonNode *node) *node {
	if bothOnRight(currentNode, collidedNode, lastCommonNode.low()) {
		return getLastCommonNode(currentNode, collidedNode, lastCommonNode.right)
	}
	if bothOnLeft(currentNode, collidedNode, lastCommonNode.low()) {
		return getLastCommonNode(currentNode, collidedNode, lastCommonNode.left)
	}
	return lastCommonNode
}

func getNextRightNode(currentNode *nodeDetails, root *node) *nodeDetails {
	isRightSubtree := currentNode.right.low() > root.low()
	isMovingRight := currentNode.right.high() >= currentNode.lastEdgeParent.high()
	lastEdgeParent := currentNode.lastEdgeParent
	// if new edge detected, update it
	if isRightSubtree && isMovingRight {
		lastEdgeParent = currentNode.node.right
	}

	return &nodeDetails{
		node:           currentNode.right,
		lastEdgeParent: lastEdgeParent,
	}
}

func updateOffsetsIfNeeded(currentOffset int, minOffset int, maxOffset int) (int, int) {
	if maxOffset < currentOffset {
		maxOffset = currentOffset
	}
	if minOffset > currentOffset {
		minOffset = currentOffset
	}
	return minOffset, maxOffset
}

func getFirstCollidedNode(nodesOnLevel []*nodeDetails, nextNode *node) *nodeDetails {
	for i := len(nodesOnLevel) - 1; i >= 0; i-- {
		levelNode := nodesOnLevel[i]
		if levelNode.offset >= nextNode.offset {
			return levelNode
		}
	}
	return nil
}

func bothOnLeft(l1 *node, l2 *node, rootLow uint64) bool {
	return l1.low() <= rootLow && l2.low() <= rootLow
}

func bothOnRight(l1 *node, l2 *node, rootLow uint64) bool {
	return l1.low() >= rootLow && l2.low() >= rootLow
}

func isSameSubtree(l1 *node, l2 *node, rootLow uint64) bool {
	return bothOnLeft(l1, l2, rootLow) || bothOnRight(l1, l2, rootLow)
}

func shiftSubTreeOffsets(root *node, direction int) {
	if root == nil {
		return
	}

	root.offset += direction
	if root.left != nil {
		root.left.parentOffset = root.offset
		shiftSubTreeOffsets(root.left, direction)
	}
	if root.right != nil {
		root.right.parentOffset = root.offset
		shiftSubTreeOffsets(root.right, direction)
	}
}

func computeArrowWithSpacing(currentNode *node, currentLineLen int, minOffset int, elementSize int) string {
	arrowsOffset := currentNode.parentOffset - minOffset
	if currentNode.offset > currentNode.parentOffset {
		numUnderscores := (currentNode.offset-currentNode.parentOffset)*elementSize - elementSize/2 + 1
		underscores := strings.Repeat(underscore, numUnderscores)
		emptyArrowSpace := strings.Repeat(space, arrowsOffset*elementSize-currentLineLen+elementSize/2)
		return emptyArrowSpace + rightArrow + underscores
	}

	numUnderscores := (currentNode.parentOffset-currentNode.offset)*elementSize - elementSize/2 + 1
	underscores := strings.Repeat(underscore, numUnderscores)
	emptyArrowSpace := strings.Repeat(space, arrowsOffset*elementSize-currentLineLen-numUnderscores-1+elementSize/2)
	return emptyArrowSpace + underscores + leftArrow
}
