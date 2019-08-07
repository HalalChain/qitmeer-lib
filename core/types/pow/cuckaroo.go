package pow

import (
	"errors"
	"fmt"
	"github.com/HalalChain/qitmeer-lib/common/hash"
	"github.com/HalalChain/qitmeer-lib/crypto/cuckoo"
	"github.com/HalalChain/qitmeer-lib/log"
	"math/big"
)

const (
	EDGE_BITS_START = 4
	EDGE_BITS_END = 8
	CIRCLE_NONCE_START = 12
	CIRCLE_NONCE_END = 180
)

type Cuckaroo struct {
	Cuckoo
}

func (this *Cuckaroo) Verify(headerWithoutProofData []byte,targetDiff uint64) error{
	h := hash.HashH(headerWithoutProofData)
	nonces := this.GetCircleNonces()
	err := cuckoo.VerifyCuckaroo(h[:],nonces[:])
	if err != nil{
		log.Debug("Verify Error!",err)
		return err
	}
	fmt.Println(fmt.Sprintf("===================target difficulty:%d",targetDiff))
	if this.CalcCuckooDiff(this.GetScale(),this.GetBlockHash([]byte{})) < targetDiff{
		return errors.New("difficulty is too easy!")
	}
	return nil
}

func (this *Cuckaroo) GetNextDiffBig(weightedSumDiv *big.Int,oldDiffBig *big.Int,currentPowPercent *big.Int,param *PowConfig) *big.Int{
	nextDiffBig := oldDiffBig.Div(oldDiffBig, weightedSumDiv)
	targetPercent := this.PowPercent(param)
	if currentPowPercent.Cmp(targetPercent) > 0{
		currentPowPercent.Div(currentPowPercent,targetPercent)
		nextDiffBig.Mul(nextDiffBig,currentPowPercent)
	}
	return nextDiffBig
}

func (this *Cuckaroo) PowPercent(param *PowConfig) *big.Int{
	targetPercent := big.NewInt(int64(param.CuckarooPercent))
	targetPercent.Lsh(targetPercent,32)
	return targetPercent
}

func (this *Cuckaroo) GetMinDiff(param *PowConfig) uint64{
	return uint64(param.CuckarooPowLimitBits)
}