package gee

import (
	"fmt"
	"strings"
)

type node struct {
	pattern string  // 待匹配路由,例如/p/:lang
 	part string     // 路由中的一部分，例如:lang
	children []*node  //子节点[doc,tutorial,intro]
	isWild bool  //是否精确匹配，part 含有 : 或 * 时为true
}

//第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild{
			return child
		}
	}
	return nil
}
//在该节点中查找是否存在，如果存放，则收集起来
//part    记录路径中的一部分
//pattern 记录完整的路径
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {

		if child.part == part || child.isWild{
			nodes = append(nodes,child)
		}
	}
	return nodes
}
/**
1 pattern：路由完全路径
2 parts：
3 height：树高
 */
func (n *node) insert(pattern string,parts []string,height int)  {
	if len(parts) == height{
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil{
		//将该节点作为孩子加入
		child = &node{part: part,isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children,child)
	}
	child.insert(pattern,parts,height+1)
}
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part,"*"){
		if n.pattern == ""{
			return nil
		}
		return n
	}
	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
func main() {
	pattern := parsePattern("/p/:lang/doc")
	for _, v := range pattern {
		fmt.Println(v)
	}
}