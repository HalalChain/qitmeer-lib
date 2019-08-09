// Copyright (c) 2017-2018 The qitmeer developers
// Copyright (c) 2014-2016 The btcsuite developers
// Copyright (c) 2015-2017 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package params

import (
	"github.com/HalalChain/qitmeer-lib/common"
	"github.com/HalalChain/qitmeer-lib/core/protocol"
	"github.com/HalalChain/qitmeer-lib/core/types/pow"
	"math/big"
	"time"
)

// testPowNetPowLimit is the highest proof of work value a block can
// have for the test network. It is the value 2^232 - 1.
var	testPowNetPowLimit = new(big.Int).Sub(new(big.Int).Lsh(common.Big1, 255), common.Big1)

// testPowNetParams defines the network parameters for the test network.
var TestPowNetParams = Params{
	Name:        "testPowNet",
	Net:         protocol.TestPowNet,
	DefaultPort: "18132",
	DNSSeeds: []DNSSeed{
		//{"testPowNet-seed.hlcwallet.info", true},
		//		//{"testPowNet-seed.qitmeer.xyz", true},
		//		//{"testPowNet-seed.qitmeer.top", true},
	},

	// Chain parameters
	GenesisBlock:             &testPowNetGenesisBlock,
	GenesisHash:              &testPowNetGenesisHash,
	PowLimit:                 testPowNetPowLimit,
	ReduceMinDifficulty:      false,
	MinDiffReductionTime:     0, // Does not apply since ReduceMinDifficulty false
	GenerateSupported:        true,
	PowConfig :&pow.PowConfig{
		PowLimitBits:             0x207fffff,
		Blake2bDPercent:          34,
		CuckarooPercent:          33,
		CuckatooPercent:          33,
		CuckarooScale:            1856,
		CuckatooScale:            1856,
		CuckarooPowLimitBits:     2000,
		CuckatooPowLimitBits:     1000,
	},

	WorkDiffAlpha:            1,
	WorkDiffWindowSize:       144,
	WorkDiffWindows:          20,
	MaximumBlockSizes:        []int{1310720},
	MaxTxSize:                1000000,
	TargetTimePerBlock:       time.Minute * 2,
	TargetTimespan:           time.Minute * 2 * 144, // TimePerBlock * WindowSize
	RetargetAdjustmentFactor: 4,

	// Subsidy parameters.
	BaseSubsidy:              2500000000, // 25 Coin
	MulSubsidy:               100,
	DivSubsidy:               101,
	SubsidyReductionInterval: 2048,
	WorkRewardProportion:     9,
	StakeRewardProportion:    0,
	BlockTaxProportion:       1,

	// Checkpoints ordered from oldest to newest.
	Checkpoints: []Checkpoint{
	},

	// Consensus rule change deployments.
	//
	Deployments: map[uint32][]ConsensusDeployment{
	},

	// Address encoding magics
	NetworkAddressPrefix: "T",
	PubKeyAddrID:         [2]byte{0x0f, 0x0f}, // starts with Tk
	PubKeyHashAddrID:     [2]byte{0x0f, 0x12}, // starts with Tm
	PKHEdwardsAddrID:     [2]byte{0x0f, 0x01}, // starts with Te
	PKHSchnorrAddrID:     [2]byte{0x0f, 0x1e}, // starts with Tr
	ScriptHashAddrID:     [2]byte{0x0e, 0xe2}, // starts with TS
	PrivateKeyID:         [2]byte{0x0c, 0xe2}, // starts with Pt

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x04, 0x35, 0x83, 0x97}, // starts with tprv
	HDPublicKeyID:  [4]byte{0x04, 0x35, 0x87, 0xd1}, // starts with tpub

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType: 11,

	CoinbaseMaturity:        16,
}
