type ListNode struct {
	Key, Val int
	Next *ListNode
}

type MyHashMap struct {
	data []*ListNode
}

func Constructor() MyHashMap {
	data := make([]*ListNode, 1000)
	return MyHashMap{data: data}
}

func (this *MyHashMap) hash(key int) int {
	return key % len(this.data)
}

func (this *MyHashMap) Put(key int, value int) {
	index := this.hash(key)
	cur := this.data[index]
	if cur == nil {
		this.data[index] = &ListNode{Key: key, Val: value}
		return
	}
	if cur.Key == key {
		cur.Val = value
		return
	}

	for cur.Next != nil {
		if cur.Next.Key == key {
			cur.Next.Val = value
			return
		}
		cur = cur.Next
	}
	cur.Next = &ListNode{Key: key, Val: value}
}

func (this *MyHashMap) Get(key int) int {
	cur := this.data[this.hash(key)]
	if cur == nil {
		return -1
	}
	for cur != nil {
		if cur.Key == key {
			return cur.Val
		}
		cur = cur.Next
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	index := this.hash(key)
	cur := this.data[index]
	if cur == nil {
		return
	}
	if cur.Key == key {
		this.data[index] = cur.Next
		return
	}
	for cur.Next != nil {
		if cur.Next.Key == key {
			cur.Next = cur.Next.Next
			return
		}
		cur = cur.Next
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */