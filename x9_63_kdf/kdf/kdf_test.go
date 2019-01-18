package kdf

import (
	"testing"
	"fmt"
)

func TestKdf(t *testing.T) {
	pri :=[]byte("1122334455667788")
	var klen int = 32
	out,_ :=X9_63_kdf(pri,klen)
	fmt.Printf("(pri)%s (klen)%d (k)%X \n",pri,klen, out)
	klen =64
	out,_ =X9_63_kdf(pri,klen)
	fmt.Printf("(pri)%s (klen)%d (k)%X \n",pri,klen, out)
	klen =48
	out,_ =X9_63_kdf(pri,klen)
	fmt.Printf("(pri)%s (klen)%d (k)%X \n",pri,klen, out)
	klen =25
	out,_ =X9_63_kdf(pri,klen)
	fmt.Printf("(pri)%s (klen)%d (k)%X \n",pri,klen, out)
}