class Node:
    def __init__(self, val):
        self.val = val
        self.next = None
        self.prev = None

    def __str__(self):
        return self.val


if __name__ == "__main__":
    node_a = Node("A")
    node_b = Node("B")
    node_c = Node("C")
    node_a.next = node_b
    node_b.prev = node_a
    node_b.next = node_c
    node_c.prev = node_b

    print(node_a, node_b.next)
