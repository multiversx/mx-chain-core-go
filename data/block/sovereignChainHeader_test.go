package block

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/data"
)

func TestSovereignChainHeader_GetEpochStartHandler(t *testing.T) {
	t.Parallel()

	economics := Economics{
		NodePrice: big.NewInt(100),
	}
	epochStartData := EpochStartCrossChainData{
		Round: 4,
	}
	epochStartSov := EpochStartSovereign{
		Economics:                     economics,
		LastFinalizedCrossChainHeader: epochStartData,
	}
	sovHdr := &SovereignChainHeader{
		EpochStart: epochStartSov,
	}

	require.Equal(t, &epochStartSov, sovHdr.GetEpochStartHandler())
	require.Equal(t, &economics, sovHdr.GetEpochStartHandler().GetEconomicsHandler())
	require.Empty(t, sovHdr.GetEpochStartHandler().GetLastFinalizedHeaderHandlers())

	epochStartData.ShardID = 8
	err := sovHdr.GetEpochStartHandler().SetLastFinalizedHeaders([]data.EpochStartShardDataHandler{&epochStartData})
	require.Nil(t, err)
	require.Empty(t, sovHdr.GetEpochStartHandler().GetLastFinalizedHeaderHandlers())

	epochStartData.ShardID = core.MainChainShardId
	epochStartSov.LastFinalizedCrossChainHeader = epochStartData
	err = sovHdr.GetEpochStartHandler().SetLastFinalizedHeaders([]data.EpochStartShardDataHandler{&epochStartData})
	require.Nil(t, err)

	require.Equal(t, []data.EpochStartShardDataHandler{&epochStartData}, sovHdr.GetEpochStartHandler().GetLastFinalizedHeaderHandlers())
	require.Equal(t, &epochStartSov, sovHdr.GetEpochStartHandler())
}

func TestSovereignChainHeader_GetEconomicsHandler(t *testing.T) {
	t.Parallel()

	economics := Economics{
		NodePrice: big.NewInt(100),
	}
	sovHdr := &SovereignChainHeader{
		EpochStart: EpochStartSovereign{
			Economics: economics,
		},
	}
	require.Equal(t, &economics, sovHdr.GetEpochStartHandler().GetEconomicsHandler())

	economics2 := Economics{
		NodePrice: big.NewInt(200),
	}
	err := sovHdr.GetEpochStartHandler().SetEconomics(&economics2)
	require.Nil(t, err)
	require.Equal(t, &economics2, sovHdr.GetEpochStartHandler().GetEconomicsHandler())
}
