/*
Remove Nth node from end of the list
linked list
input: head, n

n: 2 # start 1 right?
head>1>2>3>4>5>6>6

# first thought
identify numb of nodes
num-n

# second thought


*/

type ll struct {
	var data int
	var next *ll
}
// 1>2>3

// 1;
// head = nil
// 
func countLL(head *ll) int {
	var c int = 0
	// if head == nil {
	// 	return 0
	// }
	for head != nil {
		c += 1
		head = head.next
	}
	return c
}

func removeNthNode(head *ll, n int)  {
	var c int = 0

	if head == nil {
		return
	}

	n = countLL(head) - n + 1

	for head.next != nil {
		c += 1
		if c == n {
			head.next = head.next.next
			break
		}
		head = head.next
	}
}
