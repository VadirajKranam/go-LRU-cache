package main

import "fmt"
const SIZE=5
type Node struct{
Val string
left *Node
right *Node
}
type Queue struct{
Head *Node
Tail *Node
Length int64
}
type Cache struct{
Queue Queue
Hash Hash
}
type Hash map[string]*Node
func NewCache() Cache{
return Cache{Queue: NewQueue(),Hash:Hash{}}
}

func NewQueue() Queue{
head:=&Node{}
tail:=&Node{}
head.right=tail
tail.left=head
return Queue{Head: head,Tail: tail}
}

func (c *Cache) Check(str string){
node:=&Node{}
if val,ok:=c.Hash[str];ok{
node=c.Remove(val)
}else{
node=&Node{Val: str}
}
c.Add(node)
c.Hash[str]=node
}
func (c *Cache) Remove(n *Node) *Node{
fmt.Printf("remove:%s\n",n.Val)
left:=n.left
right:=n.right
left.right=right
right.left=left

c.Queue.Length-=1
delete(c.Hash,n.Val)
return n
}
func (c *Cache) Add(n *Node) {
fmt.Printf("add :%s\n",n.Val)
temp:=c.Queue.Head.right
c.Queue.Head.right=n
n.left=c.Queue.Head
n.right=temp
temp.left=n
c.Queue.Length+=1
if c.Queue.Length>SIZE{
c.Remove(c.Queue.Tail.left)
} 
}
func (c *Cache) Display(){
c.Queue.Display()
}
func (q *Queue) Display(){
node:=q.Head.right
fmt.Printf("%d-[",q.Length)
for i:=0;i<int(q.Length);i++{
fmt.Printf("{%s}",node.Val)
if i<int(q.Length)-1{
fmt.Printf("<-->")
}
node=node.right
}
fmt.Println("]")
}
func main() {
	fmt.Println("Start cache")
    cache:=NewCache()
    for _,word:=range []string{"parrot", "avocado", "dragonfruit", "tree", "potato", "tomato", "tree", "dog"}{
   cache.Check(word)
   cache.Display()
}
}