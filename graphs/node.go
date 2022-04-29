package graphs

type Node struct {
	Visited bool
	Key     string
}

func (n Node) compare(other Node) int {
	if n.Key == other.Key {
		return 0
	}

	return 1
}
