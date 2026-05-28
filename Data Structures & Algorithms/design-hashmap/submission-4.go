type ListNode struct {
	Key, Val int
	Next *ListNode
}

type MyHashMap struct {
	data []*ListNode
}

func Constructor() MyHashMap {
	data := make([]*ListNode, 1000)
	for i := range data {
		data[i] = &ListNode{Key: -1, Val: -1}
	}
	return MyHashMap{data: data}
}

func (this *MyHashMap) hash(key int) int {
	return key % len(this.data)
}

func (this *MyHashMap) Put(key int, value int) {
	bucket := this.hash(key)
	cur := this.data[bucket]
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
	bucket := this.hash(key)
	cur := this.data[bucket].Next
	for cur != nil {
		if cur.Key == key {
			return cur.Val
		}
		cur = cur.Next
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	bucket := this.hash(key)
	cur := this.data[bucket]
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