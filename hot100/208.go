package hot100

import "fmt"

type Test208 struct{}

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor2() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *Trie) SearchPrefix(prefix string) *Trie {
	node := t
	for _, ch := range prefix {
		ch -= 'a'
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

func (t *Trie) Search(word string) bool {
	node := t.SearchPrefix(word)
	return node != nil && node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.SearchPrefix(prefix) != nil
}

func (T Test208) Run() {
	obj := Constructor2()
	obj.Insert("w")
	param_2 := obj.Search("word")
	param_3 := obj.StartsWith("prefix")
	fmt.Println(param_2)
	fmt.Println(param_3)
	fmt.Println(obj.children['w'-'a'])
}
