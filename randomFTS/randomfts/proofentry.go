package randomfts

import (
	"github.com/ethereum/go-ethereum/common"
)

/**
 * A class containing an entry (x1, x2, hash) in a Merkle proof, where
 * "x1" is the number of coins in the left subtree, "x2" is the number
 * of coins the right subtree and "hash" is the Merkle hash
 * H(left hash | right hash | x1 | x2).
 */
type ProofEntry struct {
	hash common.Hash
	x1 int
	x2 int
}

func NewProofEntry(hash common.Hash,x1 int,x2 int) ProofEntry {
	return  ProofEntry{hash: hash, x1: x1, x2: x2}
}

func (f ProofEntry)getLeftBound() int {
	return f.x1
}

func (f ProofEntry)getRightBound() int {
	return f.x2
}

func (f ProofEntry)getMerkleHash() common.Hash {
	return f.hash
}