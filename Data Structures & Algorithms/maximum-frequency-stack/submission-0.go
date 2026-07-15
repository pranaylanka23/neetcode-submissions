type FreqStack struct {
	cnt map[int]int
	stacks [][]int
}

func Constructor() FreqStack {
	return FreqStack{
		cnt: make(map[int]int),
		stacks: [][]int{{}},
	}
}

func (this *FreqStack) Push(val int) {
	this.cnt[val]++
	valCnt := this.cnt[val]
	if valCnt == len(this.stacks){
		this.stacks = append(this.stacks, []int{})
	}
	this.stacks[valCnt] = append(this.stacks[valCnt], val)
}

func (this *FreqStack) Pop() int {
	lastStack := this.stacks[len(this.stacks)-1]
	ans := lastStack[len(lastStack)-1]
	this.stacks[len(this.stacks)-1] = lastStack[:len(lastStack)-1]
	this.cnt[ans]--
	if len(this.stacks[len(this.stacks)-1]) == 0 {
		this.stacks = this.stacks[:len(this.stacks)-1]
	}
	return ans
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor()
 * obj.Push(val)
 * param2 := obj.Pop()
 */
 