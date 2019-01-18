package randomperm

import (
	"testing"
	"github.com/ethereum/go-ethereum/common"
	"fmt"
)

func TestRandomPerm(t *testing.T) {
	m :=int(21)
	n :=int(6)
	seed := SeedFromBytes([]byte("address"))
	members :=make([]common.Address,0,m)
	for i := 0; i < m; i++ {
		a :=seed.Deri(i)
		b :=a[:]
		members =append(members,common.BytesToAddress(b))
	}
	for k,v := range members {
		fmt.Printf("addres%d:%#x\n",k,v)
	}
	// SortAddresses from small to big
	SortAddresses(members)
	fmt.Println("SortAddresses from small to big:")
	for k,v := range members {
		fmt.Printf("addres%d:%#x\n",k,v)
	}
	indices := M_RandomSelect_N([]byte("tm01"),m,n)
	validators := make([]common.Address, n)
	fmt.Println("randomPerm:")
	for j, idx := range indices {
		validators[j] = members[idx]
		fmt.Printf("addres%d:%#x\n",idx,members[idx])
	}
}
