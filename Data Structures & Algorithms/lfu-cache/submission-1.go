type Node struct{
    key, val, freq int
    prev, next *Node
}

type LFUCache struct {
    cap int
    left, right *Node
    cache map[int]*Node
}


func Constructor(capacity int) LFUCache {
    lfu := LFUCache{
        cap: capacity,
        cache: make(map[int]*Node),
        left: &Node{freq: math.MaxInt},
        right: &Node{freq: math.MinInt},
    }
    lfu.left.next = lfu.right
    lfu.right.prev = lfu.left
    return lfu
}

func (this *LFUCache) remove(node *Node){
    prev, next := node.prev, node.next
    prev.next = next
    next.prev = prev
}

func (this *LFUCache) insert(node *Node){
    curr := this.left.next
    for curr.freq > node.freq{
        curr = curr.next
    }
    prev, next := curr.prev, curr
    node.prev = prev
    node.next = next
    prev.next = node
    next.prev = node
}

func (this *LFUCache) Get(key int) int {
    if node,exists := this.cache[key]; exists{
        this.remove(node)
        node.freq++
        this.insert(node)
        return node.val
    }
    return -1
}


func (this *LFUCache) Put(key int, value int)  {
    freq := 0
    if node, exists := this.cache[key]; exists{
        delete(this.cache, node.key)
        this.remove(node)
        freq = node.freq
        this.cap++
    }
    freq++
    node := &Node{key: key, val: value, freq: freq}

    if this.cap==0{
        lfu := this.right.prev
        this.remove(lfu)
        delete(this.cache, lfu.key)
        this.cap++
    }

    this.insert(node)
    this.cache[key] = node
    this.cap--
}


/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param1 := obj.Get(key);
 * obj.Put(key,value);
 */
