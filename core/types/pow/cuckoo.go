package pow

import (
	"encoding/binary"
	"fmt"
	"github.com/HalalChain/qitmeer-lib/common/hash"
	"github.com/HalalChain/qitmeer-lib/common/util"
	"github.com/HalalChain/qitmeer-lib/crypto/cuckoo"
	"math/big"
	"sort"
)

type Cuckoo struct {
	Pow
}

const (
	PROOF_DATA_EDGE_BITS_START = 0
	PROOF_DATA_EDGE_BITS_END = 4
	PROOF_DATA_EDGE_SCALE_START = 4
	PROOF_DATA_EDGE_SCALE_END = 8
	PROOF_DATA_CIRCLE_NONCE_END = 176
)

// set edge bits
func (this *Cuckoo) SetEdgeBits (edge_bits uint32) {
	binary.LittleEndian.PutUint32(this.ProofData[PROOF_DATA_EDGE_BITS_START:PROOF_DATA_EDGE_BITS_END],uint32(edge_bits))
}

// get edge bits
func (this *Cuckoo) GetEdgeBits () uint32 {
	return binary.LittleEndian.Uint32(this.ProofData[PROOF_DATA_EDGE_BITS_START:PROOF_DATA_EDGE_BITS_END])
}

// set edge circles
func (this *Cuckoo) SetCircleEdges (edges []uint32) {
	for i:=0 ;i<len(edges);i++{
		b := make([]byte,4)
		binary.LittleEndian.PutUint32(b,edges[i])
		copy(this.ProofData[(i*4)+PROOF_DATA_EDGE_SCALE_END:(i*4)+PROOF_DATA_EDGE_SCALE_END+4],b)
	}
}

func (this *Cuckoo) GetCircleNonces () (nonces [cuckoo.ProofSize]uint32) {
	nonces = [cuckoo.ProofSize]uint32{}
	j := 0
	for i :=PROOF_DATA_EDGE_SCALE_END;i<PROOF_DATA_CIRCLE_NONCE_END;i+=4{
		nonceBytes := this.ProofData[i:i+4]
		nonces[j] = binary.LittleEndian.Uint32(nonceBytes)
		j++
	}
	return
}

// set scale ,the diff scale of circle
func (this *Cuckoo) SetScale (scale uint32) {
	binary.LittleEndian.PutUint32(this.ProofData[PROOF_DATA_EDGE_SCALE_START:PROOF_DATA_EDGE_SCALE_END],uint32(scale))
}

//get scale ,the diff scale of circle
func (this *Cuckoo) GetScale () int64 {
	return int64(binary.LittleEndian.Uint32(this.ProofData[PROOF_DATA_EDGE_SCALE_START:PROOF_DATA_EDGE_SCALE_END]))
}

func (this *Cuckoo)GetBlockHash (data []byte) hash.Hash {
	circlNonces := [cuckoo.ProofSize]uint64{}
	nonces := this.GetCircleNonces()
	for i:=0;i<len(nonces);i++{
		circlNonces[i] = uint64(nonces[i])
	}
	return this.CuckooHash(circlNonces[:],int(this.GetEdgeBits()))
}

//calc cuckoo hash
func (this *Cuckoo)CuckooHash(nonces []uint64,nonce_bits int) hash.Hash {
	sort.Slice(nonces, func(i, j int) bool {
		return nonces[i] < nonces[j]
	})
	bitvec,_ := util.New(nonce_bits*cuckoo.ProofSize)
	for i:=41;i>=0;i--{
		n := i
		nonce := nonces[i]
		for bit:= 0;bit < nonce_bits;bit++{
			if nonce & (1 << uint(bit)) != 0 {
				bitvec.SetBitAt(n * nonce_bits + bit)
			}
		}
	}
	h := hash.HashH(bitvec.Bytes())
	util.ReverseBytes(h[:])
	return h
}

//calc cuckoo diff
func CalcCuckooDiff(scale int64,blockHash hash.Hash) uint64 {
	c := &big.Int{}
	util.ReverseBytes(blockHash[:])
	c.SetUint64(binary.BigEndian.Uint64(blockHash[:8]))
	a := big.NewInt(scale)
	d := big.NewInt(1)
	d.Lsh(d,64)
	a.Mul(a,d)
	e := a.Div(a,c)
	fmt.Println(fmt.Sprintf("===================solution difficulty:%d",e.Uint64()))
	return e.Uint64()
}