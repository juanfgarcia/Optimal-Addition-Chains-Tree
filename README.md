# Optimal addition chains tree generator
A small program for creating a tree of Optimal Addition Chains created for exercise 2.1 of "From Mathematics to Generic Programming"

## Tree of optimal addition chains under 100
![optimal_addition_chains_tree_under_100](https://github.com/juanfgarcia/Optimal-Addition-Chains-Tree/blob/master/tree.png)

## How to execute
The program outputs a string with the graph in DOT notation, so DOT is a needed to execute
To install dot

### In Arch Linux

```
pacman -S graphviz
```

### In Debian (or Debian based distros)
```
apt-get install graphviz
```

### Creating the tree
```
go run main.go | dot -Tpng > tree.png && xdg-open tree.png 
```

Note that xdg-open is only for opening and can change in your OS.


