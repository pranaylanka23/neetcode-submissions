type Node struct{
    val int
    next *Node
    prev *Node
}

type MyCircularQueue struct {
    cap int
    left *Node
    right *Node

}


func Constructor(k int) MyCircularQueue {
    left := &Node{val: 0}
    right := &Node{val: 0, prev: left}
    left.next = right
    return MyCircularQueue{
        cap: k,
        left: left,
        right: right,
    }
}


func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.IsFull(){ return false}
    node := &Node{val: value, prev: this.right.prev, next: this.right}
    this.right.prev.next = node
    this.right.prev = node
    this.cap--
    return true
}


func (this *MyCircularQueue) DeQueue() bool {
    if this.IsEmpty(){ return false }
    this.left.next = this.left.next.next
    this.left.next.prev = this.left
    this.cap++
    return true
}


func (this *MyCircularQueue) Front() int {
    if this.IsEmpty() { return -1}
    return this.left.next.val
}


func (this *MyCircularQueue) Rear() int {
    if this.IsEmpty() { return -1}
    return this.right.prev.val
}


func (this *MyCircularQueue) IsEmpty() bool {
    return this.left.next==this.right
}


func (this *MyCircularQueue) IsFull() bool {
    return this.cap==0
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param1 := obj.EnQueue(value);
 * param2 := obj.DeQueue();
 * param3 := obj.Front();
 * param4 := obj.Rear();
 * param5 := obj.IsEmpty();
 * param6 := obj.IsFull();
 */
 