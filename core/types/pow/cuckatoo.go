package pow

import (
	"errors"
	"github.com/HalalChain/qitmeer-lib/common/hash"
	"github.com/HalalChain/qitmeer-lib/crypto/cuckoo"
	"github.com/HalalChain/qitmeer-lib/log"
	"math/big"
)

type Cuckatoo struct {
	Cuckoo
}

func (this *Cuckatoo) Verify(headerWithoutProofData []byte,targetDiff uint64) error{
	h := hash.HashH(headerWithoutProofData)
	nonces := this.GetCircleNonces()
	err := cuckoo.VerifyCuckatoo(h[:],nonces[:])
	if err != nil{
		log.Debug("Verify Error!",err)
		return err
	}
	if this.CalcCuckooDiff(this.GetScale(),this.GetBlockHash([]byte{})) < targetDiff{
		return errors.New("difficulty is too easy!")
	}
	return nil
}

func (this *Cuckatoo) GetMinDiff(env int) uint64{
	//env 0 private 1 test 2 main
	return 3
}

func (this *Cuckatoo) GetNextDiffBig(weightedSumDiv *big.Int,oldDiffBig *big.Int) *big.Int{
	nextDiffBig := oldDiffBig.Div(oldDiffBig, weightedSumDiv)
	return nextDiffBig
}