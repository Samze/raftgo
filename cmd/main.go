package main

import "github.com/samze/raftgo"

func main() {
	//Get all nodes ip addresses

	n := raftgo.NewNode()

	n.FSM.Event("BecomeLeader")
}
