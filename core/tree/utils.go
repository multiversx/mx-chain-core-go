package tree

import (
	"sort"
	"strings"
)

const (
	newLine    = "\n"
	underscore = "_"
	space      = " "
	leftArrow  = "/"
	rightArrow = "\\"
)

type currentNodeDetails struct {
	*node
	parent *node
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

			// skip arrow for root
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
	queue := make([]*currentNodeDetails, 0)
	queue = append(queue, &currentNodeDetails{
		node:   root,
		parent: nil,
	})
	nodesOnLevels := make([][]*node, 0)
	nodesOnLevels = append(nodesOnLevels, []*node{root})

	// extract nodes through level order traversal
	for len(queue) > 0 {
		queueSize := len(queue)
		currentLevel := make([]*node, 0)

		for i := 0; i < queueSize; i++ {
			currentNode := queue[0]
			queue = queue[1:]

			if maxOffset < currentNode.offset {
				maxOffset = currentNode.offset
			}
			if minOffset > currentNode.offset {
				minOffset = currentNode.offset
			}

			if currentNode.left != nil {
				// if next left offset already exists, the offsets must be updated
				// if the current node is in the left subtree of the tree, shift the subtree to left
				// otherwise, shift it to right
				if isOffsetDuplicateInLevel(currentLevel, currentNode.left.offset) {
					if currentNode.low() < root.low() {
						shiftSubTreeOffsets(currentNode.parent.left, -1)
					} else {
						shiftSubTreeOffsets(currentNode.node, 1)
					}
				}

				queue = append(queue, &currentNodeDetails{
					node:   currentNode.left,
					parent: currentNode.node,
				})
				currentLevel = append(currentLevel, currentNode.left)
			}
			if currentNode.right != nil {
				queue = append(queue, &currentNodeDetails{
					node:   currentNode.right,
					parent: currentNode.node,
				})
				currentLevel = append(currentLevel, currentNode.right)
			}
		}

		if len(currentLevel) > 0 {
			nodesOnLevels = append(nodesOnLevels, currentLevel)
		}
	}

	return nodesOnLevels, minOffset
}

func isOffsetDuplicateInLevel(nodesOnLevel []*node, currentOffset int) bool {
	for _, levelNode := range nodesOnLevel {
		if levelNode.offset == currentOffset {
			return true
		}
	}
	return false
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
	arrowForNodeWithSpacing := ""
	arrowsOffset := currentNode.parentOffset - minOffset
	if currentNode.offset > currentNode.parentOffset {
		numUnderscores := (currentNode.offset-currentNode.parentOffset)*elementSize - elementSize/2 + 1
		underscores := strings.Repeat(underscore, numUnderscores)
		emptyArrowSpace := strings.Repeat(space, arrowsOffset*elementSize-currentLineLen+elementSize/2)
		arrowForNodeWithSpacing = emptyArrowSpace + rightArrow + underscores
	} else {
		numUnderscores := (currentNode.parentOffset-currentNode.offset)*elementSize - elementSize/2 + 1
		underscores := strings.Repeat(underscore, numUnderscores)
		emptyArrowSpace := strings.Repeat(space, arrowsOffset*elementSize-currentLineLen-numUnderscores-1+elementSize/2)
		arrowForNodeWithSpacing = emptyArrowSpace + underscores + leftArrow
	}
	return arrowForNodeWithSpacing
}

func sortAndOptimizeIntervals(intervals []BlocksExceptionInterval) []BlocksExceptionInterval {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Low <= intervals[j].Low
	})

	optimizedIntervals := make([]BlocksExceptionInterval, 0, len(intervals))
	optimizeIntervalsSlice(intervals, &optimizedIntervals)

	return optimizedIntervals
}

func optimizeIntervalsSlice(intervals []BlocksExceptionInterval, finalIntervals *[]BlocksExceptionInterval) {
	midIdx := len(intervals) / 2
	*finalIntervals = append(*finalIntervals, intervals[midIdx])
	nextLeftInterval := intervals[:midIdx]
	if len(nextLeftInterval) > 0 {
		optimizeIntervalsSlice(nextLeftInterval, finalIntervals)
	}
	nextRightInterval := intervals[midIdx+1:]
	if len(nextRightInterval) > 0 {
		optimizeIntervalsSlice(nextRightInterval, finalIntervals)
	}
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
