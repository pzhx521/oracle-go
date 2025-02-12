// desc:
// @author renshiwei
// Date: 2023/4/7 15:33

package withdraw

import (
	"github.com/NodeDAO/oracle-go/app/consensusModule"
	"github.com/NodeDAO/oracle-go/consensus/beacon"
	"github.com/NodeDAO/oracle-go/contracts/withdrawOracle"
	consensusApi "github.com/attestantio/go-eth2-client/api/v1"
	"math/big"
)

type WithdrawHelper struct {
	// param
	refSlot          *big.Int
	consensusVersion *big.Int

	executionBlock           *beacon.ExecutionBlock
	delayedExitSlashStandard *big.Int
	clVaultMinSettleLimit    *big.Int

	validatorExaMap        map[string]*ValidatorExa
	requireReportValidator map[string]*ValidatorExa

	clBalance      *big.Int
	clVaultBalance *big.Int

	delayedExitTokenIds []*big.Int

	largeExitDelayedRequestIds []*big.Int

	totalOperatorClCapital *big.Int
	totalNftCount          *big.Int

	withdrawInfos      []withdrawOracle.WithdrawInfo
	exitValidatorInfos []withdrawOracle.ExitValidatorInfo

	// res
	reportData *withdrawOracle.WithdrawOracleReportData

	// report
	oracle              *Oracle
	hashConsensusHelper *consensusModule.HashConsensusHelper
}

type ValidatorExa struct {
	Validator *consensusApi.Validator

	IsExited          bool
	ExitedSlot        *big.Int
	ExitedBlockHeight *big.Int

	TokenId    *big.Int
	OperatorId *big.Int
	// Whether oracle needs to make a report
	IsNeedOracleReportExit bool
	// IsNeedOracleReportExit = true And then to calculate
	ExitedAmount *big.Int
	// Whether the tokenId is owned by the pledge pool
	IsOwnerLiqPool bool

	// 1.slashed 2.exited 3.Not OracleReportExit
	SlashAmount *big.Int

	IsDelayedExit bool
}

type EffectiveOperator struct {
	VnftCount      uint64
	OperatorReward withdrawOracle.WithdrawInfo
}

type Oracle struct {
}

// version
const (
	CONSENSUS_VERSION = 1
	CONTRACT_VERSION  = 1
)
