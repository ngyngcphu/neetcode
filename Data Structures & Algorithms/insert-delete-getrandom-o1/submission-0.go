type RandomizedSet struct {
	data map[int]bool
	size int
}

func Constructor() RandomizedSet {
	return RandomizedSet{data: make(map[int]bool), size: 0}
}

func (this *RandomizedSet) Insert(val int) bool {
	if this.data[val] {
		return false
	}
	this.data[val] = true
	this.size++
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if !this.data[val] {
		return false
	}
	this.size--
	delete(this.data, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(this.size)
	for key := range this.data {
		if index == 0 {
			return key
		}
		index--
	}
	return 0
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
 