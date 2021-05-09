package main

import "fmt"

const MAXCAP = 26 // a-z

type Trie struct {
	next map[rune]*Trie  //存放下一个节点全部的消息
	isWord bool
}
//Initialize your data structure here.
func Constructor() Trie {
	//初始化根节点
	root := new(Trie)
	root.next = make(map[rune]*Trie,MAXCAP)
	root.isWord = false
	return *root
}
//插入一个单词到trie
func (this *Trie) Insert(word string) {
	for _, v := range word {
		//如果再节点的下一层位置没有，则新建下一个节点
		if this.next[v] == nil{
			//新建一个子节点
			node := new(Trie)
			node.next = make(map[rune]*Trie,MAXCAP)
			node.isWord = false
			//该节点指向子节点
			this.next[v] = node
		}
		//如果有的话，则进入下一个节点
		this = this.next[v]
	}
	this.isWord = true
}
func (this *Trie) Search(word string) bool {
	for _, v := range word {
		if this.next[v] == nil{
			return false
		}
		this = this.next[v]
	}
	return this.isWord
}
func (this *Trie) StartsWith(prefix string) bool {
	for _, v := range prefix {
		if this.next[v] == nil{
			return false
		}
		this = this.next[v]
	}
	return false
}
func main() {
	t := Constructor()
	t.Insert("Hello")
	fmt.Println(t.Search("Hello"))
	fmt.Println(t.Search("Hallo"))
	//client, err := rpc.Dial("tcp", "localhost:1234")
	//if err != nil{
	//	log.Fatal("dialing:",err)
	//}
	//var reply string
	//err = client.Call("HelloService.Hello","hello",&reply)
	//if err != nil{}
}
