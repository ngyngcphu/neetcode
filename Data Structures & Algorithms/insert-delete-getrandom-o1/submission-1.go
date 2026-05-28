type RandomizedSet struct {
	hmap map[int]int
	arr []int
}

func Constructor() RandomizedSet {
	return RandomizedSet {
		hmap: make(map[int]int),
		arr: []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.hmap[val]; ok {
		return false
	}
	this.hmap[val] = len(this.arr)
	this.arr = append(this.arr, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.hmap[val]; !ok {
		return false
	}
	index := this.hmap[val]
	last := this.arr[len(this.arr) - 1]
	this.arr[index] = last
	this.hmap[last] = index
	delete(this.hmap, val) 
	this.arr = this.arr[:len(this.arr) - 1]
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.arr[rand.Intn(len(this.arr))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
 