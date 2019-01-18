package randomfts

import (
	"github.com/ethereum/go-ethereum/common"
	"fmt"
	"strconv"
)

/**
 * A stakeholder in the Merkle tree. Each stakeholder has a name,
 * and controls an amount of coins.
 */
type StakeHolder struct {
	name string
	addr common.Address
	coins int
}

func NewStakeHolder(name string,addr common.Address,coins int) StakeHolder {
	return  StakeHolder{name:name,addr:addr,coins:coins}
}

func (s StakeHolder)getName() string {
	return s.name
}

func (s StakeHolder) getAddr() common.Address {
	return s.addr
}

func (s StakeHolder) getCoins() int {
	return s.coins
}

func (s StakeHolder) toBytes() []byte {
	coin :=strconv.Itoa(s.coins)
	if len(coin)%2!=0 {
		coin ="0"+coin
	}
	temp :=fmt.Sprintf("%x%x%s",s.name,s.addr,coin)
	var dst []byte
	fmt.Sscanf(temp, "%x", &dst)
	return dst
}
