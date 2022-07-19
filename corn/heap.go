package corn

type data []Event

func (h data) Len() int {
	return len(h)
}

func (h data) Less(i, j int) bool {
	return h[i].Time.Unix() < h[j].Time.Unix()
}

func (h data) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *data) Push(v interface{}) {
	*h = append(*h, v.(Event))
}

func (h *data) Pop() interface{} {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}
