package raftgo

import (
	"fmt"
	"time"

	"github.com/looplab/fsm"
)

type Node struct {
	FSM     *fsm.FSM
	timeout time.Duration

	leader <-chan string
	nodes  []chan string
}

func NewNode() *Node {
	n := &Node{
		timeout: time.Second * 5,
	}

	n.FSM = fsm.NewFSM(
		"follower",
		fsm.Events{
			{Name: "StartElection", Src: []string{"follower"}, Dst: "candidate"},
			{Name: "WinElection", Src: []string{"candidate"}, Dst: "leader"},
			{Name: "DiscoverLeader", Src: []string{"candidate"}, Dst: "follower"},
			{Name: "LeaderWithHigherTerm", Src: []string{"leader"}, Dst: "follower"},
		},
		fsm.Callbacks{
			"enter_state": func(e *fsm.Event) { n.changeState(e) },
		},
	)

	return n
}

func (n *Node) changeState(e *fsm.Event) {
	fmt.Printf("State changed to %s", e.Dst)
}

func (n *Node) followerMode() {
	for {
		select {
		case <-time.After(n.timeout):
			n.FSM.Event("StartElection")
		case <-n.leader:
			//Ping received
		}
	}
}
