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

type nodeDisplayInfo struct {
	*node
	nodeOffset   int
	parentOffset int
}

func treeToString(root *node) string {
	if root == nil {
		return ""
	}

	elementSize := len(root.toString())
	nodesOnLevels, minOffset := extractNodesOnLevels(root)
	arrowsLines := make([]string, len(nodesOnLevels))
	nodesLines := make([]string, len(nodesOnLevels))
	for level := 0; level < len(nodesOnLevels); level++ {
		lastEmptyOffset := 0
		for _, currentNodeInfo := range nodesOnLevels[level] {
			newOffset := currentNodeInfo.nodeOffset - minOffset
			distanceFromPrevious := newOffset - lastEmptyOffset
			emptySpace := strings.Repeat(space, distanceFromPrevious*elementSize)
			nodesLines[level] += emptySpace
			nodesLines[level] += currentNodeInfo.toString()
			lastEmptyOffset = newOffset + 1

			// skip arrow for root
			if level == 0 {
				continue
			}

			arrowsLines[level] += computeArrowWithSpacing(currentNodeInfo, len(arrowsLines[level]), minOffset, elementSize)
		}
	}

	printableTreeStr := ""
	for i := 0; i < len(nodesLines); i++ {
		printableTreeStr += arrowsLines[i] + newLine
		printableTreeStr += nodesLines[i] + newLine
	}

	return printableTreeStr
}

func extractNodesOnLevels(root *node) ([][]*nodeDisplayInfo, int) {
	maxOffset := 0
	minOffset := 0
	queue := make([]*nodeDisplayInfo, 0)
	rootInfo := &nodeDisplayInfo{
		node:         root,
		nodeOffset:   0,
		parentOffset: 0,
	}
	queue = append(queue, rootInfo)
	nodesOnLevels := make([][]*nodeDisplayInfo, 0)
	nodesOnLevels = append(nodesOnLevels, []*nodeDisplayInfo{rootInfo})
	offsetsMap := make(map[int]struct{}) // used to handle offsets collisions

	// extract nodes through level order traversal
	for len(queue) > 0 {
		queueSize := len(queue)
		currentLevel := make([]*nodeDisplayInfo, 0)

		for i := 0; i < queueSize; i++ {
			currentNodeInfo := queue[0]
			queue = queue[1:]

			if maxOffset < currentNodeInfo.nodeOffset {
				maxOffset = currentNodeInfo.nodeOffset
			}
			if minOffset > currentNodeInfo.nodeOffset {
				minOffset = currentNodeInfo.nodeOffset
			}

			if currentNodeInfo.left != nil {
				nextLeftOffset := currentNodeInfo.nodeOffset - 1
				// if next left offset already exists, this node must be shifted to right
				// this is done only on left due to level order traversal iterating from left to right
				_, exists := offsetsMap[nextLeftOffset]
				if exists {
					currentNodeInfo.nodeOffset++
				}
				nextNodeInfo := &nodeDisplayInfo{
					node:         currentNodeInfo.left,
					nodeOffset:   currentNodeInfo.nodeOffset - 1,
					parentOffset: currentNodeInfo.nodeOffset,
				}
				offsetsMap[nextNodeInfo.nodeOffset] = struct{}{}
				queue = append(queue, nextNodeInfo)
				currentLevel = append(currentLevel, nextNodeInfo)
			}
			if currentNodeInfo.right != nil {
				nextNodeInfo := &nodeDisplayInfo{
					node:         currentNodeInfo.right,
					nodeOffset:   currentNodeInfo.nodeOffset + 1,
					parentOffset: currentNodeInfo.nodeOffset,
				}
				offsetsMap[nextNodeInfo.nodeOffset] = struct{}{}
				queue = append(queue, nextNodeInfo)
				currentLevel = append(currentLevel, nextNodeInfo)
			}
		}

		if len(currentLevel) > 0 {
			nodesOnLevels = append(nodesOnLevels, currentLevel)
		}
	}

	return nodesOnLevels, minOffset
}

func computeArrowWithSpacing(currentNodeInfo *nodeDisplayInfo, currentLineLen int, minOffset int, elementSize int) string {
	arrowForNodeWithSpacing := ""
	arrowsOffset := currentNodeInfo.parentOffset - minOffset
	if currentNodeInfo.nodeOffset > currentNodeInfo.parentOffset {
		numUnderscores := (currentNodeInfo.nodeOffset-currentNodeInfo.parentOffset)*elementSize - elementSize/2
		underscores := strings.Repeat(underscore, numUnderscores)
		emptyArrowSpace := strings.Repeat(space, arrowsOffset*elementSize-currentLineLen+elementSize/2)
		arrowForNodeWithSpacing = emptyArrowSpace + rightArrow + underscores
	} else {
		numUnderscores := (currentNodeInfo.parentOffset-currentNodeInfo.nodeOffset)*elementSize - elementSize/2 + 1
		underscores := strings.Repeat(underscore, numUnderscores)
		emptyArrowSpace := strings.Repeat(space, arrowsOffset*elementSize-currentLineLen-numUnderscores-1+elementSize/2)
		arrowForNodeWithSpacing = emptyArrowSpace + underscores + leftArrow
	}
	return arrowForNodeWithSpacing
}

func sortIntervals(intervals []BlocksExceptionInterval) []BlocksExceptionInterval {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Low <= intervals[j].Low
	})

	midIdx := len(intervals) / 2
	tmp := intervals[0]
	intervals[0] = intervals[midIdx]
	intervals[midIdx] = tmp

	return intervals
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
