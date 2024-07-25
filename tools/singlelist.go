package tools

// 单链表节点的结构
type ListNode struct {
	val  int
	next *ListNode
}

func NewListNode(x int) *ListNode {
	return &ListNode{
		val: x,
	}
}
func ReverseList(head *ListNode) *ListNode {
	if head.next == nil {
		return head
	}
	last := ReverseList(head.next)
	head.next.next = head
	head.next = nil
	return last
}

var successor *ListNode // 后驱节点
// 將链表的前 n 个节点反轉（n <= 链表长度）
func ReverseListN(head *ListNode, n int) *ListNode {
	if n == 1 {
		// 紀錄第 n + 1 个节点
		successor = head.next
		return head
	}
	// 以 head.next 為起点，需要反轉前 n - 1 个节点
	last := ReverseListN(head.next, n-1)
	head.next.next = head
	// 让反轉之后的 head 节点和后面的节点連起来
	head.next = successor
	return last
}

func ReverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == 1 {
		return ReverseListN(head, m)
	}
	// 前進到反轉的起点触發 base case
	head.next = ReverseBetween(head.next, m-1, n-1)
	return head
}

/** 反轉区间 [a, b) 的元素，注意是左闭右开 */
func ReverseSingleList(a *ListNode, b *ListNode) *ListNode {
	var (
		pre, cur, nxt *ListNode
	)
	pre = nil
	cur = a
	nxt = a
	//终止的条件改一下就行了
	for cur != b {
		nxt = cur.next
		// 逐个结点反轉
		cur.next = pre
		// 更新指针位置
		pre = cur
		cur = nxt
	}
	// 返回反轉后的头结点
	return pre
}
func ReverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	// 区间 [a, b) 包含 k 个待反轉元素
	var (
		a, b *ListNode
	)
	b = head
	a = head
	for i := 0; i < k; i++ {
		// 不足 k 个，不需要反轉，base case
		if b == nil {
			return head
		}
		b = b.next
	}
	// 反轉前 k 个元素
	newHead := ReverseSingleList(a, b)
	// 递归反轉后续链表並連接起来
	a.next = ReverseKGroup(b, k)
	return newHead
}
