package randomperm

import (
	"math/big"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"
	"strconv"
)
/// Seed

// SeedLength --
const SeedLength = 32

// Seed --
type Seed [SeedLength]byte

// Constructors

// SeedFromBytes --
func SeedFromBytes(b []byte) (r Seed) {
	h := crypto.Keccak256Hash(b)
	copy(r[:SeedLength], h[:])
	return
}

// DerivedRand -- Derived Randomness hierarchically
func (r Seed) DerivedRand(idx []byte) Seed {
	// Keccak is not susceptible to length-extension-attacks, so we can use it as-is to implement an HMAC
	return SeedFromBytes(crypto.Keccak256(r[:], idx))
}

// Shortcuts to the derivation function

// Ders --
// ... by string
func (r Seed) Ders(s ...string) Seed {
	ri := r
	for _, si := range s {
		ri = ri.DerivedRand([]byte(si))
	}
	return ri
}

// Deri --
// ... by int
func (r Seed) Deri(i int) Seed {
	return r.Ders(strconv.Itoa(i))
}

func (r Seed) Modulo(n int) int {
	// modulo len(groups) with big.Ints
	//var b big.Int
	b := big.NewInt(0)
	b.SetBytes(r[:])
	b.Mod(b, big.NewInt(int64(n)))
	return int(b.Int64())
}
// RandomPerm --
// Convert to a random permutation
func M_RandomSelect_N(seed []byte,m int, n int) []int {
	h := SeedFromBytes(seed)
	l := make([]int, m)
	for i := range l {
		l[i] = i
	}
	for i := 0; i < n; i++ {
		j := h.Deri(i).Modulo(m-i) + i
		l[i], l[j] = l[j], l[i]
	}
	return l[:n]
}

func sortByHex(addresses []common.Address, l int, r int) {
	if l < r {
		pivot := addresses[(l + r) / 2].Hex()
		i := l
		j := r
		var tmp common.Address
		for i <= j {
			for addresses[i].Hex() < pivot { i++ }
			for addresses[j].Hex() > pivot { j-- }
			if i <= j {
				tmp = addresses[i]
				addresses[i] = addresses[j]
				addresses[j] = tmp
				i++
				j--
			}
		}
		if l < j {
			sortByHex(addresses, l, j)
		}
		if i < r {
			sortByHex(addresses, i, r)
		}
	}
}
// SortAddresses - Sort a list of address.
func SortAddresses(addresses []common.Address) {
	n := len(addresses)
	sortByHex(addresses, 0, n - 1)
}


