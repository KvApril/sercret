package kdf
import "errors"
import "github.com/slifeio/go-slifeio/crypto"

func X9_63_kdf(z []byte, klen int) ( []byte, error) {
	rlen :=klen
	if rlen < 1{
		errors.New("klen lenth less than 1")
	}
	var dgstlen int =32
	ct := []byte{0,0,0,1}
	out :=make([]byte,klen)
	var len int =0
	for rlen >0{
		d :=crypto.Keccak256(z, ct)
		if rlen >=dgstlen{
			copy(out[len:],d[:dgstlen])
			len +=dgstlen
		}else{
			copy(out[len:],d[:rlen])
			len += rlen
		}
		rlen -= dgstlen
		ct[3]+=1
	}
	return out,nil
}
