package main

import (
	"container/heap"
	"fmt"
)

const MAX_LENGTH = 256

type HuffmanNode struct {
	char     string
	priority int
	lnode    *HuffmanNode
	rnode    *HuffmanNode
}

func newNode(char string, priority int, lnode, rnode *HuffmanNode) *HuffmanNode {
	node := new(HuffmanNode)
	node.char = char
	node.priority = priority
	node.lnode = lnode
	node.rnode = rnode

	return node
}

// Priority queue implementation
type PriorityQueue []*HuffmanNode

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(v any) {
	letter := v.(*HuffmanNode)
	*pq = append(*pq, letter)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	length := len(old)
	popped_node := old[length-1]
	old[length-1] = nil

	*pq = old[0 : length-1]
	return popped_node
}

// pre-order traversal
func display(root, node *HuffmanNode, display_root bool) {
	if display_root {
		fmt.Printf("ROOT\nNode char: %q\nNode priority: %v\n\n", string(root.char), root.priority)
	}

	if lnode := node.lnode; lnode != nil {
		fmt.Printf("LNODE\nNode char: %q\nNode priority: %v\n\n", string(lnode.char), lnode.priority)
		display(root, node.lnode, false)
	}

	if rnode := node.rnode; rnode != nil {
		fmt.Printf("RNODE\nNode char: %q\nNode priority: %v\n\n", string(rnode.char), rnode.priority)
		display(root, node.rnode, false)
	}
}

func tableify(root, node *HuffmanNode, tbl_val *uint8, huffman_table map[string]uint8) {
	if lnode := node.lnode; lnode != nil {
		if len(lnode.char) == 1 {
			huffman_table[lnode.char] += *tbl_val
		}

		tableify(root, node.lnode, tbl_val, huffman_table)
	}

	if rnode := node.rnode; rnode != nil {
		*tbl_val++
		if len(rnode.char) == 1 {
			huffman_table[rnode.char] += *tbl_val
		}

		tableify(root, node.rnode, tbl_val, huffman_table)
	}
}

func compress(input string) (map[string]uint8, error) {
	letter_map := make(map[string]int)
	for _, char := range input {
		letter_map[string(char)] += 1
	}

	pq := make(PriorityQueue, len(letter_map))
	var idx int = 0

	for char, priority := range letter_map {
		v := newNode(char, priority, nil, nil)
		pq[idx] = v
		idx++
	}

	heap.Init(&pq)

	var node1 *HuffmanNode
	var node2 *HuffmanNode
	huff_table := make(map[string]uint8, MAX_LENGTH)
	huff_len := 0

	for pq.Len() > 1 {
		node1 = pq.Pop().(*HuffmanNode)
		node2 = pq.Pop().(*HuffmanNode)

		parent := newNode(node1.char+node2.char, node1.priority+node2.priority, node1, node2)

		pq.Push(parent)
		huff_len += 2
	}

	if huff_len > MAX_LENGTH {
		return nil, fmt.Errorf("Huffman tree exceeded max length: %v\n", huff_len)
	}

	root := pq.Pop().(*HuffmanNode)
	code := uint8(0)
	tableify(root, root, &code, huff_table)

	return huff_table, nil
}

func main() {
	str := "The quick brown fox jumps over the lazy dog"
	huff_tbl, err := compress(str)
	if err != nil {
		panic("AAAHAHHAAH")
	}

	for str, huff_code := range huff_tbl {
		fmt.Printf("%v; code -> %b\n", str, huff_code)
	}
}
