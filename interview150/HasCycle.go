package interview150

func HasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	m := make(map[*ListNode]bool)
	m[head] = true
	for head.Next != nil {
		if v, exist := m[head.Next]; !exist {
			m[head.Next] = true
		} else {
			return v
		}
		head = head.Next
	}
	return false
}
