package block

import (
	"encoding/json"
	"testing"

	"github.com/multiversx/mx-chain-core-go/data"
	"github.com/stretchr/testify/require"
)

func createJsonForHeaderV2() string {
	return `{"header":{"nonce":481,"prevHash":"nNqMnj/cTiZYVMq2WW8bh9vhiN69D/AIIm7wLn/nm+0=","prevRandSeed":"sG4G+2bvTXI/htmsnJAxfQXd9oTZe5KNZ5W766kCFBbFxl7B9yZ7YCiTyzkO5X2T","randSeed":"ACmbkLg73NmmZ8sGHFhe3MYtdK46wbKIUt+ivJsdHJld82HIAvnzjF0ezqUOdsOW","shardID":0,"timeStamp":1752768353,"round":0,"epoch":5,"blockBodyType":0,"leaderSignature":"aqchjZKetNFlMalYh+sgxcOdTxDABP7iIz5wOo5qirDI7BnhkuSQtDXhloNNZDqD","miniBlockHeaders":[{"hash":"JZ0x61Ds1sC7788L5sd3HM4CdcCsQxImNnwDY9sseUA=","senderShardID":4294967295,"receiverShardID":0,"txCount":5,"type":255,"reserved":"IAQ="},{"hash":"bvoV+3E/L3glspLncniBWc5Y913oZf9xLVdcbCdQ2R4=","senderShardID":4294967295,"receiverShardID":4294967280,"txCount":6,"type":60,"reserved":"IAU="},{"hash":"h4DQEYG9V7cCfWwC7n3HR7JRRXyxxP3wi/1aQYwaIWM=","senderShardID":4294967295,"receiverShardID":4294967280,"txCount":5,"type":60,"reserved":"IAQ="}],"peerChanges":null,"rootHash":"1vyoepMXFURwMyBaRW5bEfPVJnaD9R1VLo/w6edvSO8=","metaBlockHashes":["aZiUegDAPYMjxFoDZifKS1OtucE82KjcBcV7JOAnKrk="],"txCount":16,"epochStartMetaHash":"aZiUegDAPYMjxFoDZifKS1OtucE82KjcBcV7JOAnKrk=","receiptsHash":"DldRwCblQ7Loqy6wYJnaodHl30d3j3eH+qtFzfEv46g=","chainID":"bG9jYWxuZXQ=","softwareVersion":"Mg==","accumulatedFees":0,"developerFees":0},"scheduledRootHash":"SZRu3iHeUgmPfL99TQapZNOcqKSWYmp5rMrAuOMbmu0=","scheduledAccumulatedFees":0,"scheduledDeveloperFees":0,"scheduledGasProvided":0,"scheduledGasPenalized":0,"scheduledGasRefunded":0}`
}

func createJsonForMetaBlock() string {
	return `{"nonce":184,"epoch":2,"round":0,"timeStamp":1752766553,"shardInfo":[{"headerHash":"N4Be23RIX4Hdb/IX8N9Rn9IVrDwNv0x/aRBG3DeZ59s=","shardMiniBlockHeaders":[{"hash":"SQnnrD2Cv9UbULqY2vrdsP9pKzVrp9lUgMaKf8N/VQs=","senderShardID":0,"receiverShardID":0,"txCount":100,"type":0}],"prevRandSeed":"AtTtjVgLLCR1vcN5lhMgKAXSQ+uGgQJQAGCIRXpRur2WgyOFWVwGsvB0XNr5tT2D","round":205,"prevHash":"5DzInuk8HiY/x21RCIAaLnmEp2pNcj3GFhjV/D0ugeo=","nonce":182,"accumulatedFees":5000000000000000,"developerFees":0,"numPendingMiniBlocks":1,"lastIncludedMetaNonce":181,"shardID":0,"txCount":100,"epoch":2}],"peerInfo":null,"leaderSignature":"MaAFUyniShBNVbL01Mf5WJOAh0ypTKcjFtQ4E+wODRWpUWjb1/icT07eeEK5n7oT","prevHash":"p5RrqnclvenWpggjZazqDuNMSh/BAKXUUZOW4Ty3R80=","prevRandSeed":"n+jWtdpAJrz1G8YyxNtn6aKuMSwrpVhwzhaHVbEsTIJe0i5N3gzl73QWxaHWjUiU","randSeed":"ek0OGMLItkHOwp/AtNtM8jup4ZKUgXw2xPpMEvARqWUo+xiMai7K0Zt5n/EKl0QL","rootHash":"SA2azL3/LsUqsfofORRsha0dXjBlshUEVJELa9uBTuQ=","validatorStatsRootHash":"YWE259eQYLgZ94BeQE5Ur9/IuGlrQj3K3NMb/FnsYus=","miniBlockHeaders":null,"receiptsHash":"DldRwCblQ7Loqy6wYJnaodHl30d3j3eH+qtFzfEv46g=","epochStart":{"lastFinalizedHeaders":null,"economics":{"prevEpochStartRound":0}},"chainID":"bG9jYWxuZXQ=","softwareVersion":"Mg==","accumulatedFees":0,"accumulatedFeesInEpoch":5050000000000000,"developerFees":0,"devFeesInEpoch":0,"txCount":100}`
}

func TestPrettifyHeader(t *testing.T) {
	t.Parallel()

	t.Run("headerV2", func(t *testing.T) {
		header := &HeaderV2{}
		require.NoError(t, json.Unmarshal([]byte(createJsonForHeaderV2()), header))
		prettified, _ := PrettifyStruct(header)
		t.Log(prettified)
	})

	t.Run("metablock", func(t *testing.T) {
		header := &MetaBlock{}
		require.NoError(t, json.Unmarshal([]byte(createJsonForMetaBlock()), header))
		prettified, _ := PrettifyStruct(header)
		t.Log(prettified)
	})
}

func Test_test(t *testing.T) {
	hdr := &Header{
		Nonce:            2,
		Round:            2,
		PrevHash:         []byte("prevHash"),
		PrevRandSeed:     []byte("prevRandSeed"),
		Signature:        []byte("signature"),
		PubKeysBitmap:    []byte("00110"),
		ShardID:          0,
		RootHash:         []byte("rootHash"),
		MiniBlockHeaders: []MiniBlockHeader{},
	}

	hdrv2 := &HeaderV2{
		Header: hdr,
		ScheduledGasProvided: 0,
	}
	var h data.HeaderHandler
	h = hdrv2
	PrettifyHeader(h)
	//marshalled, _ := json.Marshal(h)
	prettified, _ := PrettifyStruct(h)
	t.Log("marshalled header", "header", prettified)

	metaHeader := &MetaBlock{
		Nonce:         2,
		Round:         2,
		PrevHash:      []byte("prevHash"),
		PrevRandSeed:  []byte("prevRandSeed"),
		Signature:     []byte("signature"),
		PubKeysBitmap: []byte("00110"),
		RootHash:      []byte("rootHash"),
		ShardInfo: []ShardData{
			{
				ShardID: 0,
				TxCount: 100,
			},
		},
	}
	h = metaHeader
	//marshalledMeta, _ := json.Marshal(h)
	prettified, _ = PrettifyStruct(h) 
	t.Log("marshalled meta header", "header", prettified)
	PrettifyHeader(h)
	//t.Log("marshalled meta header - ORIGINAL", "header", string(marshalledMeta))
}
