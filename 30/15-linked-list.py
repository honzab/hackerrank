class Node:
    def __init__(self, data):
        self.data = data
        self.next = None


class Solution:
    def display(self, head):
        current = head
        while current:
            print current.data,
            current = current.next

    def insert(self, head, data):
        n = Node(data)

        if not head:
            return n

        # Iterate to the tail
        current = head
        while current.next is not None:
            current = current.next

        # Set the new node
        current.next = n
        return head


mylist = Solution()
T = int(input())
head = None
for i in range(T):
    data = int(input())
    head = mylist.insert(head, data)
mylist.display(head)
