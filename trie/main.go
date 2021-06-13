package main

const Aplhabet int = 26

//node struct holds the array and flag to check is the node is leaf node
type Node struct {
	child [26]*Node
	isEnd bool
}

type Trie struct {
	root *Node
}

//Insert method to add the characters of word into its respective nodes
func (t *Trie) Insert(w string) {
	wordLenght := len(w)
	currentNode := t.root
	for i := 0; i < wordLenght; i++ {
		charIndex := w[i] - 'a'
		if currentNode.child[charIndex] == nil { //there is no node present
			currentNode.child[charIndex] = &Node{}
		}
		currentNode = currentNode.child[charIndex] //move to the nex t node
	}
	currentNode.isEnd = true
}

//Search method return true if it finds the words we are searching for

func (t *Trie) Search(w string) bool {
	wordLenght := len(w)
	currentNode := t.root
	for i := 0; i < wordLenght; i++ {
		charIndex := w[i] - 'a'
		if currentNode.child[charIndex] == nil {
			return false
		}
		currentNode = currentNode.child[charIndex]
	}
	return currentNode.isEnd
}

func InitTrie() *Trie {
	res := &Trie{&Node{}}
	return res
}

func main() {
	trie := InitTrie()
	trie.Insert("danish")
	println(trie.Search("danish"))
}
