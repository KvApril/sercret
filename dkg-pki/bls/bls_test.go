package bls

import (
	"testing"
	"github.com/dfinity/dkg-pki/blscgo"
)

func TestComparison(t *testing.T) {
	t.Log("testComparison")
	blscgo.Init(blscgo.CurveFp254BNb)
	b := Decimal2Big("16798108731015832284940804142231733909759579603404752749028378864165570215948")
	sec := SeckeyFromBigInt(&b)
	t.Log("sec.Hex: ", sec.Hex())
	t.Log("sec.String: ", sec.String())
//	fmt.Println("sec.Hex: ", sec.Hex())
//	fmt.Println("sec.String: ", sec.String())

	// Add Seckeys
	sum := AggregateSeckeys([]Seckey{sec, sec})
	t.Log("sum: ", sum.Hex())
//	fmt.Println("sum: ", sum.Hex())
//	fmt.Println("sum: ", sum.Hex()[2:])

	sk := sec.SecretKey()
	t.Log("sk = sec.SecretKey(): ", sk.GetHexString())
//	fmt.Println("sk = sec.SecretKey(): ", sk.GetHexString())

	// Pubkey
	pk := sk.GetPublicKey()
	t.Log("pk: ", pk.GetHexString())
//	fmt.Println("pk: ", pk.GetHexString())
	pub := PubkeyFromSeckey(sec)
	t.Log("pub: ", pub.String())
//	fmt.Println("pub: ", pub.String())
	//pub2 := PublicKeyFromSeckey(sec)
	//t.Log("pub2: ", pub2.String())

	// Add SecretKeys
	sk.Add(sk)
	t.Log("sksum: ", sk.GetHexString())
//	fmt.Println("sksum: ", sk.GetHexString())

	if sk.GetHexString() != sum.Hex()[2:] {
		t.Error("Mismatch in secret key addition")
	}

	// Sig
	sig := Sign(sec, []byte("hi"))
	asig := AggregateSigs([]Signature{sig, sig})
	if !VerifyAggregateSig([]Pubkey{pub, pub}, []byte("hi"), asig) {
		t.Error("Aggregated signature does not verify")
	}
}
