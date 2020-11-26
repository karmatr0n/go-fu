// https://en.wikipedia.org/wiki/Heap_(data_structure)
package heap

type Heap struct {
	items []int
}

func initHeap() *Heap {
	return &Heap{}
}

func (h *Heap) GetLeftIndex(parentIndex int) int {
	return 2*parentIndex + 1
}

func (h *Heap) GetRightIndex(parentIndex int) int {
	return 2*parentIndex + 2
}

func (h *Heap) GetParentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

func (h *Heap) HasLeft(index int) bool {
	return h.GetLeftIndex(index) < len(h.items)
}

func (h *Heap) HasRight(index int) bool {
	return h.GetRightIndex(index) < len(h.items)
}

func (h *Heap) HasParent(index int) bool {
	return h.GetParentIndex(index) >= 0
}

func (h *Heap) Right(index int) int {
	return h.items[h.GetRightIndex(index)]
}

func (h *Heap) Left(index int) int {
	return h.items[h.GetLeftIndex(index)]
}

func (h *Heap) Parent(index int) int {
	return h.items[h.GetParentIndex(index)]
}

func (h *Heap) Swap(indexOne, indexTwo int) {
	h.items[indexOne], h.items[indexTwo] = h.items[indexTwo], h.items[indexOne]
}
