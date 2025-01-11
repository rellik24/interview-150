package interview150

import "fmt"

func FMinStack() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin()) // return -3
	minStack.Pop()
	minStack.Top()
}

type MinStack struct {
	data []int // elements stack
}

func Constructor() MinStack {
	return MinStack{}
}

func (mstack *MinStack) Push(val int) {
	mstack.data = append(mstack.data, val)
}

func (mstack *MinStack) Pop() {
	mstack.data = mstack.data[:len(mstack.data)-1]
}

func (mstack *MinStack) Top() int {
	return mstack.data[len(mstack.data)-1]
}

func (mstack *MinStack) GetMin() int {
	min := mstack.data[0]
	for _, v := range mstack.data {
		if v < min {
			min = v
		}
	}
	return min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
