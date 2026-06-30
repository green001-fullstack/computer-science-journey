# The difference between a Node struct and a Node object.

### A Node struct is a blueprint that defines the fields a node should have. A Node object is an actual instance created in memory from that blueprint

# Why creating nodes does not automatically create a tree.

### The tree only happens when the left and right component of the struct starts referencing the address of another node

# How nodes become connected.

### Nodes become connected when a node's Left or Right field stores a reference to another node.

# Draw the tree you built and explain how it exists in memory

  root
          │
          ▼

      +---------+
      |   15    |
      +---------+
      | Left ---|------------------+
      | Right --|-----------+      |
      +---------+           |      |
                            |      |
                            ▼      ▼

                      +---------+   +---------+
                      |   10    |   |   25    |
                      +---------+   +---------+
                      | Left ---|---------+
                      | Right --|--> nil
                      +---------+
                           |
                           ▼

                      +---------+
                      |    8    |
                      +---------+
                      | Left --> nil
                      | Right-> nil
                      +---------+