package bitree

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strings"
	"sync"
	"tsp/models"
)

// Unicode символы для визуализации
var (
	chHor     = "─"
	chVer     = "│"
	chDDia    = "┌"
	chRDia    = "┐"
	chUDia    = "└"
	chVerHor  = "├─"
	chUDiaHor = "└─"
	chDDiaHor = "┌─"
	chVerSpa  = "│ "
)

type TreeNode struct {
	Val   string
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	ID       int
	ParentID int
	LeftID   int
	RightID  int
	W        int
	In       int
	Out      int
	Sign     string
	Node     *TreeNode
	Mxs      [][]int
}

type Results struct {
	Tour []Node
	Back []Node
}

type BiTree struct {
	CurWeight   int
	State       map[int]*TreeNode
	Count       int
	Result      Results
	CurrentNode *TreeNode
	RootNode    *TreeNode
	AllNodes    map[int]Node
	CurrentID   int
	mutex       sync.Mutex
}

var BT *BiTree

func NewBiTree(mx [][]int, weight int) *BiTree {
	bt := &BiTree{
		CurWeight:   math.MaxInt,
		State:       make(map[int]*TreeNode),
		Count:       0,
		Result:      Results{},
		CurrentNode: &TreeNode{},
		RootNode:    &TreeNode{},
		AllNodes:    make(map[int]Node),
		CurrentID:   0,
		mutex:       sync.Mutex{},
	}

	// создаем Root и сохраняем в мапе
	nd := Node{ID: bt.Count, W: weight}
	bt.RootNode = &TreeNode{Val: fmt.Sprintf("w%d:Root", nd.W)}
	bt.State[bt.Count] = bt.RootNode
	bt.CurrentNode = bt.RootNode
	bt.Result.Back = append(bt.Result.Back, Node{
		ID:   bt.Count,
		W:    weight,
		In:   0,
		Out:  0,
		Sign: "",
		Node: bt.RootNode,
		//Mxs:  CloneMx(mx),
	})
	bt.AllNodes[0] = Node{
		ID:   0,
		W:    weight,
		In:   0,
		Out:  0,
		Sign: "-",
		Node: bt.RootNode,
		Mxs:  mx,
	}
	bt.Count++
	return bt
}

func CloneMx(mx [][]int) [][]int {
	lenRows := len(mx)
	lenCols := len(mx[0])
	mxClone := make([][]int, lenRows)
	for i := range mxClone {
		mxClone[i] = make([]int, lenCols)
	}
	for i := 0; i < lenRows; i++ {
		for j := 0; j < lenCols; j++ {
			mxClone[i][j] = mx[i][j]
		}
	}
	return mxClone
}

func (bt *BiTree) CreateLeftNode(parID int, mx [][]int, w, o, i int, setCurrent bool) {
	// func (bt *BiTree) CreateLeftNode(w, o, i int, setCurrent bool) {
	bt.mutex.Lock()
	defer bt.mutex.Unlock()
	if models.Debug {
		fmt.Printf("Left: id:%d, w:%d, out:%d,in:%d\n", bt.Count, w, o, i)
	}

	// nd := Node{ID: bt.Count, W: w, Out: o, Sign: "-", In: i, Mxs: CloneMx(mx)}
	nd := Node{ID: bt.Count, W: w, Out: o, Sign: "-", In: i, Mxs: CloneMx(mx), ParentID: parID}
	curNode := bt.AllNodes[bt.CurrentID]
	curNode.LeftID = nd.ID
	bt.AllNodes[bt.CurrentID] = curNode

	err := bt.CurrentNode.InsertLeft(fmt.Sprintf("id%d:w%d:-%d.%d|%d", nd.ID, nd.W, nd.Out, nd.In, nd.ParentID))
	if err != nil {
		log.Fatal("Insert Left node is failure: ", err)
	}
	nd.Node = bt.CurrentNode.Left
	bt.State[bt.Count] = bt.CurrentNode.Left
	bt.AllNodes[bt.Count] = nd

	if setCurrent {
		bt.CurrentID = nd.ID
		bt.Result.Tour = nil
		//bt.Result.Tour = append(bt.Result.Tour, nd)
		//bt.Result.Tour = nil
		// bt.CurrentNode = bt.CurrentNode.Left
	} else {
		bt.Result.Back = append(bt.Result.Back, nd)
	}

	bt.Count++
	if models.Debug {
		fmt.Println("________________________ Узел создан _______________________")
	}
}

func (bt *BiTree) CreateRightNode(parID int, mx [][]int, w, o, i int, setCurrent bool) {
	// func (bt *BiTree) CreateRightNode(w, o, i int, setCurrent bool) {
	bt.mutex.Lock()
	defer bt.mutex.Unlock()
	if models.Debug {
		fmt.Printf("Right: id:%d, w:%d, out:%d,in:%d\n", bt.Count, w, o, i)
	}

	// nd := Node{ID: bt.Count, W: w, Out: o, Sign: "+", In: i, Mxs: CloneMx(mx)}
	nd := Node{ID: bt.Count, W: w, Out: o, Sign: "+", In: i, Mxs: CloneMx(mx), ParentID: parID}
	curNode := bt.AllNodes[bt.CurrentID]
	curNode.RightID = nd.ID
	bt.AllNodes[bt.CurrentID] = curNode

	err := bt.CurrentNode.InsertRight(fmt.Sprintf("id%d:w%d:%d.%d|%d", nd.ID, nd.W, nd.Out, nd.In, nd.ParentID))
	if err != nil {
		log.Fatal("Insert Right node is failure: ", err)
	}
	nd.Node = bt.CurrentNode.Right
	bt.State[bt.Count] = bt.CurrentNode.Right
	bt.AllNodes[bt.Count] = nd

	if setCurrent {
		bt.CurrentID = nd.ID
		bt.Result.Tour = append(bt.Result.Tour, nd)
		// bt.CurrentNode = bt.CurrentNode.Right
	} else {
		bt.Result.Back = append(bt.Result.Back, nd)
	}

	bt.Count++
	if models.Debug {
		fmt.Println("________________________ Узел создан _______________________")
	}
}

func (bt *BiTree) CreateLastNode(parID int, mx [][]int, w, o, i int) {
	// func (bt *BiTree) CreateLastNode(w, o, i int) {
	bt.mutex.Lock()
	defer bt.mutex.Unlock()
	if models.Debug {
		fmt.Printf("Last: id:%d, w:%d, out:%d,in:%d\n", bt.Count, w, o, i)
	}
	// nd := Node{ID: bt.Count, W: w, Out: o, Sign: "+", In: i, Mxs: CloneMx(mx)}
	nd := Node{ID: bt.Count, W: w, Out: o, Sign: "+", In: i, Mxs: CloneMx(mx), ParentID: parID}
	curNode := bt.AllNodes[bt.CurrentID]
	curNode.RightID = nd.ID
	bt.AllNodes[bt.CurrentID] = curNode

	err := bt.CurrentNode.InsertRight(fmt.Sprintf("id%d:w%d:%d.%d|%d", nd.ID, nd.W, nd.Out, nd.In, nd.ParentID))
	if err != nil {
		log.Fatal("Insert Last node is failure: ", err)
	}
	//nd.Node = bt.CurrentNode.Right
	bt.State[bt.Count] = bt.CurrentNode.Right
	bt.AllNodes[bt.Count] = nd
	bt.CurrentID = nd.ID
	bt.Result.Tour = append(bt.Result.Tour, nd)
	bt.Count++
	if models.Debug {
		fmt.Println("________________________ Узел создан _______________________")
	}
}

// PrintInorder prints the elements in order
func (t *TreeNode) PrintInorder() {
	if t == nil {
		return
	}
	t.Left.PrintInorder()
	// fmt.Printf("%s,", t.Val)
	t.Right.PrintInorder()
}

// Insert inserts a new node into the binary tree while adhering to the rules of a perfect BST.
func (t *TreeNode) InsertLeft(value string) error {
	if t == nil {
		return errors.New("tree is nil")
	}

	// if t.val == value {
	// 	return errors.New("this node value already exists")
	// }

	if t.Left == nil {
		t.Left = &TreeNode{Val: value}
		return nil
	}
	return fmt.Errorf("левый лист узла: \"%s\" занят", t.Val)
}

// Insert inserts a new node into the binary tree while adhering to the rules of a perfect BST.
func (t *TreeNode) InsertRight(value string) error {
	if t == nil {
		return errors.New("tree is nil")
	}

	// if t.val == value {
	// 	return errors.New("this node value already exists")
	// }

	if t.Right == nil {
		t.Right = &TreeNode{Val: value}
		return nil
	}
	return fmt.Errorf("правый лист узла: \"%s\" занят", t.Val)
}

// Insert inserts a new node into the binary tree while adhering to the rules of a perfect BST.
func (t *TreeNode) Insert(value string) error {
	if t == nil {
		return errors.New("tree is nil")
	}

	if t.Val == value {
		return errors.New("this node value already exists")
	}

	if value < t.Val {
		if t.Left == nil {
			t.Left = &TreeNode{Val: value}
			return nil
		}
		return t.Left.Insert(value)
	}

	if value > t.Val {
		if t.Right == nil {
			t.Right = &TreeNode{Val: value}
			return nil
		}
		return t.Right.Insert(value)
	}
	return nil
}

// Find finds the treenode for the given node val
func (t *TreeNode) Find(value string) (TreeNode, bool) {
	if t == nil {
		return TreeNode{}, false
	}

	switch {
	case value == t.Val:
		return *t, true
	case value < t.Val:
		return t.Left.Find(value)
	default:
		return t.Right.Find(value)
	}
}

// Delete removes the Item with value from the tree
func (t *TreeNode) Delete(value string) {
	t.remove(value)
}

func (t *TreeNode) remove(value string) *TreeNode {
	if t == nil {
		return nil
	}
	if value < t.Val {
		t.Left = t.Left.remove(value)
		return t
	}
	if value > t.Val {
		t.Right = t.Right.remove(value)
		return t
	}
	if t.Left == nil && t.Right == nil {
		t = nil
		return nil
	}
	if t.Left == nil {
		t = t.Right
		return t
	}
	if t.Right == nil {
		t = t.Left
		return t
	}

	smallestValOnRight := t.Right
	for {
		//find smallest value on the right side
		if smallestValOnRight != nil && smallestValOnRight.Left != nil {
			smallestValOnRight = smallestValOnRight.Left
		} else {
			break
		}
	}

	t.Val = smallestValOnRight.Val
	t.Right = t.Right.remove(t.Val)
	return t
}

// FindMax finds the max element in the given BST
func (t *TreeNode) FindMax() string {
	if t.Right == nil {
		return t.Val
	}
	return t.Right.FindMax()
}

// FindMin finds the min element in the given BST
func (t *TreeNode) FindMin() string {
	if t.Left == nil {
		return t.Val
	}
	return t.Left.FindMin()
}

// Вспомогательная функция для повторения строки несколько раз
func RepStr(s string, cnt int) string {
	if cnt < 0 {
		panic(fmt.Sprintf("RepStr: Некорректное значение %d!", cnt))
	}
	return strings.Repeat(s, cnt)
}

// Функция PrintTree - реализация визуализации дерева
func PrintTree(node *TreeNode) {
	// Вспомогательная рекурсивная функция
	var Rec func(*TreeNode, bool) ([]string, int, int)
	Rec = func(node *TreeNode, left bool) ([]string, int, int) {
		if node == nil {
			return []string{}, 0, 0
		}
		// sval := strconv.Itoa(node.val)
		sval := node.Val
		resl, cl, lss := Rec(node.Left, true)
		resr, cr, rss := Rec(node.Right, false)

		vl := resl
		vr := resr

		lv := len(sval)
		ls := 0
		if len(vl) > 0 {
			ls = lss
		}
		rs := 0
		if len(vr) > 0 {
			rs = rss
		}

		lis := 0
		if ls == 0 {
			lis = lv / 2
		} else {
			lis = max(lv/2+1-(ls-cl), 0)
		}
		ris := 0
		if rs == 0 {
			ris = (lv + 1) / 2
		} else {
			ris = max((lv+1)/2-cr, 0)
			if lis == 0 {
				ris = max(ris, 1)
			}
		}

		dashls := 0
		if ls != 0 {
			dashls = ls - cl - 1 + lis - lv/2
		}
		dashrs := 0
		if rs != 0 {
			dashrs = cr + ris - (lv+1)/2
		}

		// Формируем первую строку
		line := ""
		if ls != 0 {
			line += RepStr(" ", cl) + chDDia + RepStr(chHor, dashls)
		}
		line += sval
		if rs != 0 {
			line += RepStr(chHor, dashrs) + chRDia + RepStr(" ", rs-cr-1)
		}

		lines := []string{line}

		// Формируем остальные строки
		for i := 0; i < max(len(vl), len(vr)); i++ {
			sl := RepStr(" ", ls)
			sr := RepStr(" ", rs)
			if i < len(vl) {
				sl = vl[i]
			}
			if i < len(vr) {
				sr = vr[i]
			}
			sl += RepStr(" ", lis)
			sr = RepStr(" ", ris) + sr
			lines = append(lines, sl+sr)
		}
		var adjustedWidth int
		if left || ls+lis == 0 || lv%2 == 1 {
			adjustedWidth = ls + lis
		} else {
			adjustedWidth = ls + lis - 1
		}

		return lines, adjustedWidth, ls + lis + ris + rs
		// return lines, (left || ls+lis == 0 || lv%2 == 1) ? ls+lis : ls+lis-1, ls+lis+ris+rs
	}

	// Запуск визуализации дерева
	lines, _, _ := Rec(node, true)
	for _, line := range lines {
		fmt.Println(line)
	}
}
