// desc:
// @author renshiwei
// Date: 2023/4/11 11:02

package consensusModule

import (
	"context"
	"github.com/NodeDAO/oracle-go/contracts/hashConsensus"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type ReportContract interface {
	GetConsensusContractAddress(ctx context.Context) (common.Address, error)

	GetConsensusContract(ctx context.Context) (*hashConsensus.HashConsensus, error)

	CheckContractVersions(ctx context.Context) error

	IsContractReportable(ctx context.Context) (bool, error)

	IsMainDataSubmitted(ctx context.Context) (bool, error)
}

type HashConsensusHelper struct {
	ReportContract ReportContract
}

type MemberInfo struct {
	IsReportMember              bool
	IsFastLane                  bool
	LastReportRefSlot           *big.Int
	FastLaneLengthSlot          *big.Int
	CurrentFrameRefSlot         *big.Int
	DeadlineSlot                *big.Int
	CurrentFrameMemberReport    [32]byte
	CurrentFrameConsensusReport [32]byte
}
