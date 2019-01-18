package randomfts

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Node struct {
	left *Node
	right *Node
	stakeHolder StakeHolder
	hash common.Hash
	isLeaf bool
}

func NewLeafNode(stakeHolder StakeHolder) Node {
	h := crypto.Keccak256Hash(stakeHolder.toBytes())
	return  Node{stakeHolder: stakeHolder, isLeaf:true, hash:h}
}

func NewNotleafNode(left *Node,right *Node,hash common.Hash) Node {
	return  Node{left: left, right: right, hash: hash}
}

func (d Node)isLeafNode() bool {
	return d.isLeaf
}

func (d Node)getStakeHolder() StakeHolder {
	return d.stakeHolder
}

func (d Node)getLeftNode() *Node {
	return d.left
}

func (d Node)getRightNode() *Node {
	return d.right
}

func (d Node)getMerkleHash() common.Hash {
	return d.hash
}

func (d Node)getCoins() int {
	if(d.isLeaf){
		return  d.stakeHolder.getCoins()
	}
	return d.left.getCoins() + d.right.getCoins()
}



