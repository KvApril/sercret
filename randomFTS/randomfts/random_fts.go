package randomfts

import "fmt"
import (
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"math/rand"
    "bytes"
)
// Create the Merkle tree
func createMerkleTree(stakeholders []StakeHolder)[]Node {
	tree :=make([]Node,len(stakeholders)*2)
	fmt.Printf("Creating Merkle tree with %d nodes\n",len(tree)-1)
	for i:=int(0);i<len(stakeholders);i++{
		tree[len(stakeholders)+i] = NewLeafNode(stakeholders[i])
	}
	for j :=len(stakeholders)-1;j>0;j--{
		left :=tree[j*2]
		right :=tree[j*2 + 1]
		x1 := new(big.Int)
		x1.SetInt64(int64(left.getCoins()))
		x2 := new(big.Int)
		x2.SetInt64(int64(right.getCoins()))
		h := crypto.Keccak256Hash(left.getMerkleHash().Bytes(),right.getMerkleHash().Bytes(),x1.Bytes(),x2.Bytes())
		tree[j] = NewNotleafNode(&left, &right,h)
	}
	for v :=int(1);v<len(tree);v++{
		fmt.Printf("Hash%d: %x\n",v,tree[v].getMerkleHash())
	}
	return tree
}

func ftsTree(tree []Node,seed int64) FtsResult {
	i :=int(1)
	rand.Seed(seed)
	merkleProof := make([]ProofEntry,0)
	for {
		if tree[i].isLeafNode(){
			return NewFtsResult(merkleProof,tree[i].getStakeHolder())
		}
		x1 := tree[i].getLeftNode().getCoins()
		x2 := tree[i].getRightNode().getCoins()
		fmt.Printf("Left subtree %d coins / right subtree %d coins\n",x1,x2)
		r := rand.Intn(x1+x2)+1
		if r <= x1 {
			fmt.Println("Choosing left subtree...")
			i *= 2
			merkleProof =append(merkleProof,NewProofEntry(tree[i + 1].getMerkleHash(),x1,x2))
		}else{
			fmt.Println("Choosing right subtree...")
			i = i*2 + 1
			merkleProof =append(merkleProof,NewProofEntry(tree[i - 1].getMerkleHash(),x1,x2))
		}
	}
}

func ftsVerify(seed int64, merkleRootHash common.Hash, ftsResult FtsResult) bool {
	rand.Seed(seed)
	audit := make([]bool,0)
	for _,proofEntry := range ftsResult.getMerkleProof() {
		x1 := proofEntry.getLeftBound()
		x2 := proofEntry.getRightBound()
		r := rand.Intn(x1 + x2)+1
		if r <= x1 {
			audit =append(audit,false)
		}else{
			audit =append(audit,true)
		}
	}
	hx := crypto.Keccak256Hash(ftsResult.getStakeHolder().toBytes())
	for i := len(ftsResult.getMerkleProof())-1; i>=0 ; i-- {
		pEntry := ftsResult.getMerkleProof()[i]
		x1 := new(big.Int)
		x1.SetInt64(int64(pEntry.getLeftBound()))
		x2 := new(big.Int)
		x2.SetInt64(int64(pEntry.getRightBound()))
		hy := pEntry.getMerkleHash()
		if audit[i]==false {
			hx = crypto.Keccak256Hash(hx.Bytes(),hy.Bytes(),x1.Bytes(),x2.Bytes())
		}else{
			hx = crypto.Keccak256Hash(hy.Bytes(),hx.Bytes(),x1.Bytes(),x2.Bytes())
		}
	}
	result := bytes.Equal(hx.Bytes(),merkleRootHash.Bytes())
	return  result
}


