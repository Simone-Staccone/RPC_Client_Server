package main

type OverlayNetwork struct {
	Network [256]Resource
}

type Resource struct {
	Id    int
	Value string
}

type Node struct {
	Id        int
	Value     string
	Successor int
}

type Reply struct {
	RET int
}

type Ret int
