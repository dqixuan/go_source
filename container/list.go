package container

// 实现带头结点的环形、双向链表

// 节点结构体
type Element struct {
	prev, next *Element // 前驱、后驱节点

	list *List // 节点所属的链表

	Value any // 节点存储内容
}

type List struct {
	root Element // 哨兵节点
	len  int     // 链表长度
}

// Next 返回当前节点
func (e *Element) Next() *Element {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

func (e *Element) Prev() *Element {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// Init 创建一个空链表或是将一个链表置为空
func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

// Len 返回链表长度
func (l *List) Len() int {
	return l.len
}

func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

func (l *List) Back() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

// insert 将节点e放在节点at之后
func (l *List) insert(e, at *Element) *Element {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++
	return e
}

func (l *List) insertValue(value any, at *Element) *Element {
	return l.insert(&Element{Value: value}, at)
}

func (l *List) remove(e *Element) {
	// 将节点e从链表中删除
	e.prev.next = e.next
	e.next.prev = e.prev

	// 将节点e的指向置空，防止内存泄露
	e.next = nil
	e.prev = nil
	e.list = nil

	// 链表长度-1
	l.len--
	return
}

func (l *List) Remove(e *Element) any {
	// 外部可调用的方法判断节点是否属于当前链表, remove方法没有判断
	if e.list == l {
		l.remove(e)
	}
	return e.Value
}

func (l *List) move(e, at *Element) {
	if e == at {
		return
	}
	e.prev.next = e.next
	e.next.prev = e.prev

	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
}
