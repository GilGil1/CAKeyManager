package node

import (
	"fmt"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
)

type Node struct {
	host host.Host
}

func (n *Node) Start() {
	// start a libp2p node with default settings
	node, err := libp2p.New()
	if err != nil {
		panic(err)
	}
	n.host = node

	// print the node's listening addresses
	fmt.Println("Listen addresses:", node.Addrs())

}

func (n *Node) Stop() {
	// shut the node down
	if err := n.host.Close(); err != nil {
		panic(err)
	}

}
