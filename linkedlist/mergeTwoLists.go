package linkedlist

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	curr1 := list1
	curr2 := list2

	for curr1 != nil && curr2 != nil {
		if curr1.Val <= curr2.Val {
			tail.Next = curr1
			curr1 = curr1.Next
		} else {
			tail.Next = curr2
			curr2 = curr2.Next
		}

		tail = tail.Next
	}

	if curr1 != nil {
		tail.Next = curr1
	}

	if curr2 != nil {
		tail.Next = curr2
	}

	return dummy.Next
}

func mergeTwoListsRC(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		list1.Next = mergeTwoListsRC(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoListsRC(list1, list2.Next)
	return list2
}
