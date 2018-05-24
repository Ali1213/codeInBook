package main
import (
	
)

// tree 这个类型
type tree {
	value int
	left,right *tree
}

// 加入一个数组，返回tree之后的数组
func Sort(values []int){
	var root *tree
	for _,v := range values {
		add(root, v)
	}
	appendValues(values[:0], root)
}

// 传入一个values
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values,t.right)
	}
	return values
} 

// 往tree里面增加
func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}

	if value < t.value {
		add(t.left, value)
	}else {
		add(t.right,value)
	}
	return t
}

func main(){
	
}