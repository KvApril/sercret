package randomfts

import (
	"testing"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"
)

func TestFts(test *testing.T) {
	// Create some stakeholders
	 stakeHolders := make([]StakeHolder,0)
	 c := []int{25,7,11,9,18,5,31,2}
	 for i := int(0) ; i < 8; i++ {
         name :=fmt.Sprintf("Stakeholder %d",i)
		 h := crypto.Keccak256Hash([]byte(name))
		 b :=h[:]
         stakeHolders = append(stakeHolders,NewStakeHolder(name,common.BytesToAddress(b),c[i]))

	}
	// Create the Merkle tree
	tree := createMerkleTree(stakeHolders)
	fmt.Println("Doing follow-the-satoshi in the stake tree")
	ftsResult := ftsTree(tree, 100)
	fmt.Println("Verifying the result")
	ftsVerify(100, tree[1].getMerkleHash(), ftsResult)
}