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

// mainPowLimit is the highest proof of work value a block can
// have for the main network. It is the value 2^224 - 1.
var mainPowLimit    = new(big.Int).Sub(new(big.Int).Lsh(common.Big1, 224), common.Big1)

// MainNetParams defines the network parameters for the main network.
var MainNetParams = Params{
	Name:        "mainnet",
	Net:         protocol.MainNet,
	DefaultPort: "8130",
	DNSSeeds: []DNSSeed{
		{"seed.qitmeer.io", true},
		{"seed2.qitmeer.io", true},
		{"seed3.qitmeer.io", true},
	},

	// Chain parameters
	GenesisBlock:             &genesisBlock,
	GenesisHash:              &genesisHash,

	PowConfig :&pow.PowConfig{
		Blake2bdPowLimit:                 mainPowLimit,
		Blake2bdPowLimitBits:             0x1d00ffff,
		Blake2bDPercent:          100,
		CuckarooPercent:          0,
		CuckatooPercent:          0,
		CuckarooDiffScale:            1856,
		CuckatooDiffScale:            1856,
		CuckarooMinDifficulty:     1000,
		CuckatooMinDifficulty:     1000,
	},
	ReduceMinDifficulty:      false,
	MinDiffReductionTime:     0, // Does not apply since ReduceMinDifficulty false
	GenerateSupported:        false,
	WorkDiffAlpha:            1,
	WorkDiffWindowSize:       144,
	WorkDiffWindows:          20,
	MaximumBlockSizes:        []int{393216},
	MaxTxSize:                393216,
	TargetTimePerBlock:       time.Minute * 5,
	TargetTimespan:           time.Minute * 5 * 144, // TimePerBlock * WindowSize
	RetargetAdjustmentFactor: 4,


	// Subsidy parameters.
	BaseSubsidy:              3119582664, // 21m
	MulSubsidy:               100,
	DivSubsidy:               101,
	SubsidyReductionInterval: 6144,
	WorkRewardProportion:     9,
	StakeRewardProportion:    0,
	BlockTaxProportion:       1,

	// Checkpoints ordered from oldest to newest.
	Checkpoints: []Checkpoint{
	},

	Deployments: map[uint32][]ConsensusDeployment{
	},

	// Address encoding magics
	NetworkAddressPrefix: "N",
	PubKeyAddrID:         [2]byte{0x0c, 0x3e}, // starts with Nk
	PubKeyHashAddrID:     [2]byte{0x0c, 0x41}, // starts with Nm
	PKHEdwardsAddrID:     [2]byte{0x0c, 0x30}, // starts with Ne
	PKHSchnorrAddrID:     [2]byte{0x0c, 0x4e}, // starts with Nr
	ScriptHashAddrID:     [2]byte{0x0c, 0x12}, // starts with NS
	PrivateKeyID:         [2]byte{0x0c, 0xd1}, // starts with Pm

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x03, 0xb8, 0xc4, 0x22}, // starts with nprv
	HDPublicKeyID:  [4]byte{0x03, 0xb8, 0xc8, 0x58}, // starts with npub

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	// TODO : register coin type
	// https://github.com/satoshilabs/slips/blob/master/slip-0044.md
	HDCoinType: 223,

	CoinbaseMaturity:        256,
}
