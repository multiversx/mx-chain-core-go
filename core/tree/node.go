package tree

type node struct {
	interval     *interval
	left         *node
	right        *node
	max          uint64
	offset       int
	parentOffset int
}

func newNode(interval *interval) *node {
	return &node{
		interval: interval,
	}
}

func (node *node) contains(value uint64) bool {
	return node.interval.contains(value)
}

func (node *node) low() uint64 {
	return node.interval.low
}

func (node *node) high() uint64 {
	return node.interval.high
}

func (node *node) toString() string {
	return node.interval.toString()
}

func (node *node) isLeaf() bool {
	return node.left == nil && node.right == nil
}
