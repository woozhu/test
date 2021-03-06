package wxXunHuan

import (
	"errors"
	"fmt"
)

type ElemType interface{}

//结点
type Node struct {
	Data ElemType
	Pre  *Node
	Next *Node
}

//链表
type List struct {
	First *Node
	Last  *Node
	Size  int
}

//工厂函数
func CreateList() *List {
	s := new(Node)
	s.Next, s.Pre = s, s
	return &List{s, s, 0}
}

//尾插法
//尾插法，改变的是最后一个节点的定义，新定义节点的前后节点，链表的首末节点
func (list *List) PushBack(x ElemType) {
	s := new(Node)
	s.Data = x                  //s节点的数据定义
	list.Last.Next = s          //改变最后一个节点的next节点
	s.Pre = list.Last           //s节点的前节点定义
	list.Last = s               //重定义Last节点 尾插所以First节点不变
	list.Last.Next = list.First //最后一个节点的下一个节点改变
	list.First.Pre = list.Last  //第一个节点的前节点改变
	list.Size++
}

//头插法
//头插法，改变第一个节点，
func (list *List) PushFront(x ElemType) {
	s := new(Node)
	s.Data = x

	s.Next = list.First.Next
	list.First.Next.Pre = s

	list.First.Next = s
	s.Pre = list.First
	if list.Size == 0 {
		list.Last = s
	}
	list.Size++
}

//尾删法
func (list *List) PopBack() bool {
	if list.IsEmpty() {
		return false
	}
	s := list.Last.Pre //找到最后一个节点的前驱
	s.Data = list.Last.Pre.Data
	s.Next = list.First
	s.Pre = list.Last.Pre.Pre
	list.Last = s
	if list.Size == 1 {
		list.First = list.Last
	}
	list.Size--
	return true
}

//头删法
func (list *List) PopFront() bool {
	if list.IsEmpty() {
		return false
	}
	s := list.First.Next //找到第一个节点
	list.First.Next = s.Next
	s.Next.Pre = list.First
	if list.Size == 1 {
		list.Last = list.First
	}
	list.Size--
	return true
}

//向前转动列表
//将列表向前转动n个节点之后的列表
func (list *List) ForWard(n int) {
	s := new(Node)
	if list.Size > 1 {
		for i := 0; i < n%list.Size; i++ {
			s = list.First.Next
			s.Data = list.First.Next.Data
			list.PopFront()
			list.PushBack(s.Data)
		}
	}
}

//向后转动列表
//将列表向后转动n个节点之后的列表
func (list *List) BackWard(n int) {
	s := new(Node)
	if list.Size > 1 {
		for i := 0; i < n%list.Size; i++ {
			s = list.Last
			s.Data = list.Last.Data
			list.PopBack()
			list.PushFront(s.Data)
		}
	}
}

//查找指定元素
func (list *List) Find(x ElemType) *Node {
	s := list.First.Next
	for s != list.First {
		if x == s.Data {
			return s
		} else {
			s = s.Next
		}
	}
	return nil
}

//获取指定元素所在的角度(0,360)。
func (list *List) GetAngle(x ElemType) int {
	s := list.First.Next
	i := 0
	for s != list.First {
		if x == s.Data {
			return 360 * i / list.Size
		} else {
			s = s.Next
			i++
		}
	}
	return -1
}

//把所有元素按照相等的角度排列一圈，共计360度,第一个元素在正北方。查找指定索引所在的元素，索引值为整数。正负均可。
func (list *List) FindByIndex(a int) *Node {
	b := a % list.Size //取度数值除以Size的余数，多于一圈的按照一圈内计算。
	if b < 0 {
		b = list.Size + b
	}
	s := list.First.Next
	i := 0
	for s != list.First {
		if i == b {
			return s
		} else {
			s = s.Next
			i++
		}
	}
	return nil
}

//根据角度来得到指定的元素(0<a<360),以正北为0度，正南为180度。
func (list *List) FindByAngle(a int) *Node {
	b := a % 360
	if b < 0 {
		b = 360 + b
	}
	m := 360 / list.Size
	s := list.First.Next
	if b >= (list.Size-1)*m+m/2 || b < m/2 {
		return s
	} else {
		for i := 1; i < list.Size-1; {
			s = s.Next
			if b >= i*m-m/2 && b < i*m+m/2 {
				return s
			}
			i++
		}
	}
	return nil
}

//按值删除结点
func (list *List) DeleteVal(x ElemType) bool {
	s := list.Find(x)
	if s != nil {
		s.Pre.Next = s.Next
		s.Next.Pre = s.Pre
		list.Size--
		//如果删除的是最后一个结点
		if s == list.Last {
			list.Last = s.Pre
		}
		return true
	}
	return false
}

//把值为x的元素的值修改为y
func (list *List) Modify(x, y ElemType) bool {
	s := list.Find(x)
	if s != nil {
		s.Data = y
		return true
	}
	return false
}

//判断链表是否为空
func (list *List) IsEmpty() bool {
	return list.Size == 0
}

//反转链表
//保留第一个结点，将剩余的结点游离出来，然后依次头插到保留的结点中
func (list *List) Reverse() {
	if list.Size > 1 {
		s := list.First.Next
		p := s.Next
		s.Next = list.First //第一个结点逆置后成为最后一个结点
		list.Last = s

		for p != list.First {
			s = p
			p = p.Next

			s.Next = list.First.Next
			list.First.Next.Pre = s

			s.Pre = list.First
			list.First.Next = s
		}
	}
}

//打印链表
func (list *List) Print() error {
	if list.IsEmpty() {
		return errors.New("this is an empty list")
	}
	s := list.First.Next
	for s != list.First {
		fmt.Printf("%v  ", s.Data)
		s = s.Next
	}
	fmt.Println()
	return nil
}
