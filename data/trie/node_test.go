package trie

import (
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"testing"
	"time"

	"github.com/ElrondNetwork/elrond-go/config"
	"github.com/ElrondNetwork/elrond-go/data"
	"github.com/ElrondNetwork/elrond-go/data/mock"
	protobuf "github.com/ElrondNetwork/elrond-go/data/trie/proto"
	"github.com/ElrondNetwork/elrond-go/marshal"
	"github.com/ElrondNetwork/elrond-go/storage/storageUnit"
	"github.com/stretchr/testify/assert"
)

var snapshotDelay = time.Millisecond
var batchDelay = 2 * time.Second

func TestNode_hashChildrenAndNodeBranchNode(t *testing.T) {
	t.Parallel()

	marsh, hasher := getTestMarshAndHasher()
	bn, collapsedBn := getBnAndCollapsedBn()
	expectedNodeHash, _ := encodeNodeAndGetHash(collapsedBn, marsh, hasher)

	hash, err := hashChildrenAndNode(bn, marsh, hasher)
	assert.Nil(t, err)
	assert.Equal(t, expectedNodeHash, hash)
}

func TestNode_hashChildrenAndNodeExtensionNode(t *testing.T) {
	t.Parallel()

	marsh, hasher := getTestMarshAndHasher()
	en, collapsedEn := getEnAndCollapsedEn()
	expectedNodeHash, _ := encodeNodeAndGetHash(collapsedEn, marsh, hasher)

	hash, err := hashChildrenAndNode(en, marsh, hasher)
	assert.Nil(t, err)
	assert.Equal(t, expectedNodeHash, hash)
}

func TestNode_hashChildrenAndNodeLeafNode(t *testing.T) {
	t.Parallel()

	marsh, hasher := getTestMarshAndHasher()
	ln := getLn()
	expectedNodeHash, _ := encodeNodeAndGetHash(ln, marsh, hasher)

	hash, err := hashChildrenAndNode(ln, marsh, hasher)
	assert.Nil(t, err)
	assert.Equal(t, expectedNodeHash, hash)
}

func TestNode_encodeNodeAndGetHashBranchNode(t *testing.T) {
	t.Parallel()

	marsh, hasher := getTestMarshAndHasher()

	encChildren := make([][]byte, nrOfChildren)
	encChildren[1] = []byte("dog")
	encChildren[10] = []byte("doge")
	bn := newBranchNode()
	bn.EncodedChildren = encChildren

	encNode, _ := marsh.Marshal(bn)
	encNode = append(encNode, branch)
	expextedHash := hasher.Compute(string(encNode))

	hash, err := encodeNodeAndGetHash(bn, marsh, hasher)
	assert.Nil(t, err)
	assert.Equal(t, expextedHash, hash)
}

func TestNode_encodeNodeAndGetHashExtensionNode(t *testing.T) {
	t.Parallel()

	marsh, hasher := getTestMarshAndHasher()
	en := &extensionNode{CollapsedEn: protobuf.CollapsedEn{Key: []byte{2}, EncodedChild: []byte("doge")}}

	encNode, _ := marsh.Marshal(en)
	encNode = append(encNode, extension)
	expextedHash := hasher.Compute(string(encNode))

	hash, err := encodeNodeAndGetHash(en, marsh, hasher)
	assert.Nil(t, err)
	assert.Equal(t, expextedHash, hash)
}

func TestNode_encodeNodeAndGetHashLeafNode(t *testing.T) {
	t.Parallel()

	marsh, hasher := getTestMarshAndHasher()
	ln := newLeafNode([]byte("dog"), []byte("dog"))

	encNode, _ := marsh.Marshal(ln)
	encNode = append(encNode, leaf)
	expextedHash := hasher.Compute(string(encNode))

	hash, err := encodeNodeAndGetHash(ln, marsh, hasher)
	assert.Nil(t, err)
	assert.Equal(t, expextedHash, hash)
}

func TestNode_encodeNodeAndCommitToDBBranchNode(t *testing.T) {
	t.Parallel()

	db := mock.NewMemDbMock()
	marsh, hasher := getTestMarshAndHasher()
	_, collapsedBn := getBnAndCollapsedBn()
	encNode, _ := marsh.Marshal(collapsedBn)
	encNode = append(encNode, branch)
	nodeHash := hasher.Compute(string(encNode))

	err := encodeNodeAndCommitToDB(collapsedBn, db, marsh, hasher)
	assert.Nil(t, err)

	val, _ := db.Get(nodeHash)
	assert.Equal(t, encNode, val)
}

func TestNode_encodeNodeAndCommitToDBExtensionNode(t *testing.T) {
	t.Parallel()

	db := mock.NewMemDbMock()
	marsh, hasher := getTestMarshAndHasher()
	_, collapsedEn := getEnAndCollapsedEn()
	encNode, _ := marsh.Marshal(collapsedEn)
	encNode = append(encNode, extension)
	nodeHash := hasher.Compute(string(encNode))

	err := encodeNodeAndCommitToDB(collapsedEn, db, marsh, hasher)
	assert.Nil(t, err)

	val, _ := db.Get(nodeHash)
	assert.Equal(t, encNode, val)
}

func TestNode_encodeNodeAndCommitToDBLeafNode(t *testing.T) {
	t.Parallel()

	db := mock.NewMemDbMock()
	marsh, hasher := getTestMarshAndHasher()
	ln := getLn()
	encNode, _ := marsh.Marshal(ln)
	encNode = append(encNode, leaf)
	nodeHash := hasher.Compute(string(encNode))

	err := encodeNodeAndCommitToDB(ln, db, marsh, hasher)
	assert.Nil(t, err)

	val, _ := db.Get(nodeHash)
	assert.Equal(t, encNode, val)
}

func TestNode_getNodeFromDBAndDecodeBranchNode(t *testing.T) {
	t.Parallel()

	db := mock.NewMemDbMock()
	marsh, hasher := getTestMarshAndHasher()
	bn, collapsedBn := getBnAndCollapsedBn()
	_ = bn.commit(false, 0, db, db, marsh, hasher)

	encNode, _ := marsh.Marshal(collapsedBn)
	encNode = append(encNode, branch)
	nodeHash := hasher.Compute(string(encNode))

	node, err := getNodeFromDBAndDecode(nodeHash, db, marsh)
	assert.Nil(t, err)

	h1, _ := encodeNodeAndGetHash(collapsedBn, marsh, hasher)
	h2, _ := encodeNodeAndGetHash(node, marsh, hasher)
	assert.Equal(t, h1, h2)
}

func TestNode_getNodeFromDBAndDecodeExtensionNode(t *testing.T) {
	t.Parallel()

	db := mock.NewMemDbMock()
	marsh, hasher := getTestMarshAndHasher()
	en, collapsedEn := getEnAndCollapsedEn()
	_ = en.commit(false, 0, db, db, marsh, hasher)

	encNode, _ := marsh.Marshal(collapsedEn)
	encNode = append(encNode, extension)
	nodeHash := hasher.Compute(string(encNode))

	node, err := getNodeFromDBAndDecode(nodeHash, db, marsh)
	assert.Nil(t, err)

	h1, _ := encodeNodeAndGetHash(collapsedEn, marsh, hasher)
	h2, _ := encodeNodeAndGetHash(node, marsh, hasher)
	assert.Equal(t, h1, h2)
}

func TestNode_getNodeFromDBAndDecodeLeafNode(t *testing.T) {
	t.Parallel()

	db := mock.NewMemDbMock()
	marsh, hasher := getTestMarshAndHasher()
	ln := getLn()
	_ = ln.commit(false, 0, db, db, marsh, hasher)

	encNode, _ := marsh.Marshal(ln)
	encNode = append(encNode, leaf)
	nodeHash := hasher.Compute(string(encNode))

	node, err := getNodeFromDBAndDecode(nodeHash, db, marsh)
	assert.Nil(t, err)
	ln = getLn()
	ln.dirty = false
	assert.Equal(t, ln, node)
}

func TestNode_resolveIfCollapsedBranchNode(t *testing.T) {
	t.Parallel()

	db := mock.NewMemDbMock()
	marsh, hasher := getTestMarshAndHasher()
	bn, collapsedBn := getBnAndCollapsedBn()
	childPos := byte(2)

	_ = bn.commit(false, 0, db, db, marsh, hasher)

	err := resolveIfCollapsed(collapsedBn, childPos, db, marsh)
	assert.Nil(t, err)
	assert.False(t, collapsedBn.isCollapsed())
}

func TestNode_resolveIfCollapsedExtensionNode(t *testing.T) {
	t.Parallel()

	db := mock.NewMemDbMock()
	marsh, hasher := getTestMarshAndHasher()
	en, collapsedEn := getEnAndCollapsedEn()

	_ = en.commit(false, 0, db, db, marsh, hasher)

	err := resolveIfCollapsed(collapsedEn, 0, db, marsh)
	assert.Nil(t, err)
	assert.False(t, collapsedEn.isCollapsed())
}

func TestNode_resolveIfCollapsedLeafNode(t *testing.T) {
	t.Parallel()

	db := mock.NewMemDbMock()
	marsh, hasher := getTestMarshAndHasher()
	ln := getLn()

	_ = ln.commit(false, 0, db, db, marsh, hasher)

	err := resolveIfCollapsed(ln, 0, db, marsh)
	assert.Nil(t, err)
	assert.False(t, ln.isCollapsed())
}

func TestNode_resolveIfCollapsedNilNode(t *testing.T) {
	t.Parallel()

	db := mock.NewMemDbMock()
	marsh, _ := getTestMarshAndHasher()
	var node *extensionNode

	err := resolveIfCollapsed(node, 0, db, marsh)
	assert.Equal(t, ErrNilNode, err)
}

func TestNode_concat(t *testing.T) {
	t.Parallel()

	a := []byte{1, 2, 3}
	var b byte
	b = 4
	ab := []byte{1, 2, 3, 4}
	assert.Equal(t, ab, concat(a, b))
}

func TestNode_hasValidHash(t *testing.T) {
	t.Parallel()

	marsh, hasher := getTestMarshAndHasher()
	bn, _ := getBnAndCollapsedBn()
	ok, err := hasValidHash(bn)
	assert.Nil(t, err)
	assert.False(t, ok)

	_ = bn.setHash(marsh, hasher)
	bn.dirty = false

	ok, err = hasValidHash(bn)
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestNode_hasValidHashNilNode(t *testing.T) {
	t.Parallel()

	var node *branchNode
	ok, err := hasValidHash(node)
	assert.Equal(t, ErrNilNode, err)
	assert.False(t, ok)
}

func TestNode_decodeNodeBranchNode(t *testing.T) {
	t.Parallel()

	marsh, hasher := getTestMarshAndHasher()
	_, collapsedBn := getBnAndCollapsedBn()

	encNode, _ := marsh.Marshal(collapsedBn)
	encNode = append(encNode, branch)

	node, err := decodeNode(encNode, marsh)
	assert.Nil(t, err)

	h1, _ := encodeNodeAndGetHash(collapsedBn, marsh, hasher)
	h2, _ := encodeNodeAndGetHash(node, marsh, hasher)
	assert.Equal(t, h1, h2)
}

func TestNode_decodeNodeExtensionNode(t *testing.T) {
	t.Parallel()

	marsh, hasher := getTestMarshAndHasher()
	_, collapsedEn := getEnAndCollapsedEn()

	encNode, _ := marsh.Marshal(collapsedEn)
	encNode = append(encNode, extension)

	node, err := decodeNode(encNode, marsh)
	assert.Nil(t, err)

	h1, _ := encodeNodeAndGetHash(collapsedEn, marsh, hasher)
	h2, _ := encodeNodeAndGetHash(node, marsh, hasher)
	assert.Equal(t, h1, h2)
}

func TestNode_decodeNodeLeafNode(t *testing.T) {
	t.Parallel()

	marsh, hasher := getTestMarshAndHasher()
	ln := getLn()

	encNode, _ := marsh.Marshal(ln)
	encNode = append(encNode, leaf)

	node, err := decodeNode(encNode, marsh)
	assert.Nil(t, err)
	ln.dirty = false

	h1, _ := encodeNodeAndGetHash(ln, marsh, hasher)
	h2, _ := encodeNodeAndGetHash(node, marsh, hasher)
	assert.Equal(t, h1, h2)
}

func TestNode_decodeNodeInvalidNode(t *testing.T) {
	t.Parallel()

	marsh, _ := getTestMarshAndHasher()
	ln := getLn()
	invalidNode := byte(6)

	encNode, _ := marsh.Marshal(ln)
	encNode = append(encNode, invalidNode)

	node, err := decodeNode(encNode, marsh)
	assert.Nil(t, node)
	assert.Equal(t, ErrInvalidNode, err)
}

func TestNode_decodeNodeInvalidEncoding(t *testing.T) {
	t.Parallel()

	marsh, _ := getTestMarshAndHasher()

	var encNode []byte

	node, err := decodeNode(encNode, marsh)
	assert.Nil(t, node)
	assert.Equal(t, ErrInvalidEncoding, err)
}

func TestNode_getEmptyNodeOfTypeBranchNode(t *testing.T) {
	t.Parallel()

	bn, err := getEmptyNodeOfType(branch)
	assert.Nil(t, err)
	assert.IsType(t, &branchNode{}, bn)
}

func TestNode_getEmptyNodeOfTypeExtensionNode(t *testing.T) {
	t.Parallel()

	en, err := getEmptyNodeOfType(extension)
	assert.Nil(t, err)
	assert.IsType(t, &extensionNode{}, en)
}

func TestNode_getEmptyNodeOfTypeLeafNode(t *testing.T) {
	t.Parallel()

	ln, err := getEmptyNodeOfType(leaf)
	assert.Nil(t, err)
	assert.IsType(t, &leafNode{}, ln)
}

func TestNode_getEmptyNodeOfTypeWrongNode(t *testing.T) {
	t.Parallel()

	n, err := getEmptyNodeOfType(6)
	assert.Equal(t, ErrInvalidNode, err)
	assert.Nil(t, n)
}

func TestNode_childPosOutOfRange(t *testing.T) {
	t.Parallel()

	assert.True(t, childPosOutOfRange(17))
	assert.False(t, childPosOutOfRange(5))
}

func TestMarshalingAndUnmarshalingWithCapnp(t *testing.T) {
	_, collapsedBn := getBnAndCollapsedBn()
	collapsedBn.dirty = false
	marsh := marshal.CapnpMarshalizer{}
	bn := newBranchNode()

	encBn, err := marsh.Marshal(collapsedBn)
	assert.Nil(t, err)
	assert.NotNil(t, encBn)

	err = marsh.Unmarshal(bn, encBn)
	assert.Nil(t, err)
	assert.Equal(t, collapsedBn, bn)
}

func TestKeyBytesToHex(t *testing.T) {
	t.Parallel()

	var test = []struct {
		key, hex []byte
	}{
		{[]byte("doe"), []byte{6, 4, 6, 15, 6, 5, 16}},
		{[]byte("dog"), []byte{6, 4, 6, 15, 6, 7, 16}},
	}

	for i := range test {
		assert.Equal(t, test[i].hex, keyBytesToHex(test[i].key))
	}
}

func TestPrefixLen(t *testing.T) {
	t.Parallel()

	var test = []struct {
		a, b   []byte
		length int
	}{
		{[]byte("doe"), []byte("dog"), 2},
		{[]byte("dog"), []byte("dogglesworth"), 3},
		{[]byte("mouse"), []byte("mouse"), 5},
		{[]byte("caterpillar"), []byte("cats"), 3},
		{[]byte("caterpillar"), []byte(""), 0},
		{[]byte(""), []byte("caterpillar"), 0},
		{[]byte("a"), []byte("caterpillar"), 0},
	}

	for i := range test {
		assert.Equal(t, test[i].length, prefixLen(test[i].a, test[i].b))
	}
}

func TestGetOldHashesIfNodeIsCollapsed(t *testing.T) {
	t.Parallel()

	msh, hsh := getTestMarshAndHasher()
	evictionWaitListSize := 100
	evictionWaitList, _ := mock.NewEvictionWaitingList(evictionWaitListSize, mock.NewMemDbMock(), msh)

	tr := &patriciaMerkleTrie{
		db:                    mock.NewMemDbMock(),
		dbEvictionWaitingList: evictionWaitList,
		oldHashes:             make([][]byte, 0),
		oldRoot:               make([]byte, 0),
		marshalizer:           msh,
		hasher:                hsh,
	}

	_ = tr.Update([]byte("doe"), []byte("reindeer"))
	_ = tr.Update([]byte("dog"), []byte("puppy"))
	_ = tr.Update([]byte("dogglesworth"), []byte("cat"))

	rootHash, _ := tr.Root()
	rootKey := []byte{6, 4, 6, 15, 6}
	nextNode, _, _ := tr.root.getNext(rootKey, tr.db, tr.marshalizer)

	_ = tr.Commit()

	tr.root = &extensionNode{
		CollapsedEn: protobuf.CollapsedEn{
			Key:          rootKey,
			EncodedChild: nextNode.getHash(),
		},
		child: nil,
		hash:  rootHash,
		dirty: false,
	}
	_ = tr.Update([]byte("doeee"), []byte("value of doeee"))

	assert.Equal(t, 3, len(tr.oldHashes))
}

func TestClearOldHashesAndOldRootOnCommit(t *testing.T) {
	t.Parallel()

	msh, hsh := getTestMarshAndHasher()
	evictionWaitListSize := 100
	evictionWaitList, _ := mock.NewEvictionWaitingList(evictionWaitListSize, mock.NewMemDbMock(), msh)

	tr := &patriciaMerkleTrie{
		db:                    mock.NewMemDbMock(),
		dbEvictionWaitingList: evictionWaitList,
		oldHashes:             make([][]byte, 0),
		oldRoot:               make([]byte, 0),
		marshalizer:           msh,
		hasher:                hsh,
	}

	_ = tr.Update([]byte("doe"), []byte("reindeer"))
	_ = tr.Update([]byte("dog"), []byte("puppy"))
	_ = tr.Update([]byte("dogglesworth"), []byte("cat"))
	_ = tr.Commit()
	root, _ := tr.Root()

	_ = tr.Update([]byte("doeee"), []byte("value of doeee"))

	assert.Equal(t, 3, len(tr.oldHashes))
	assert.Equal(t, root, tr.oldRoot)

	_ = tr.Commit()

	assert.Equal(t, 0, len(tr.oldHashes))
	assert.Equal(t, 0, len(tr.oldRoot))
}

func TestTrieDatabasePruning(t *testing.T) {
	t.Parallel()

	msh, hsh := getTestMarshAndHasher()
	size := 5
	evictionWaitList, _ := mock.NewEvictionWaitingList(size, mock.NewMemDbMock(), msh)

	tr := &patriciaMerkleTrie{
		db:                    mock.NewMemDbMock(),
		dbEvictionWaitingList: evictionWaitList,
		oldHashes:             make([][]byte, 0),
		oldRoot:               make([]byte, 0),
		marshalizer:           msh,
		hasher:                hsh,
	}

	_ = tr.Update([]byte("doe"), []byte("reindeer"))
	_ = tr.Update([]byte("dog"), []byte("puppy"))
	_ = tr.Update([]byte("dogglesworth"), []byte("cat"))
	_ = tr.Commit()

	key := []byte{6, 4, 6, 15, 6, 7, 16}
	oldHashes := make([][]byte, 0)
	n := tr.root
	rootHash, _ := tr.Root()
	oldHashes = append(oldHashes, rootHash)

	for i := 0; i < 3; i++ {
		n, key, _ = n.getNext(key, tr.db, tr.marshalizer)
		oldHashes = append(oldHashes, n.getHash())
	}

	_ = tr.Update([]byte("dog"), []byte("doee"))
	_ = tr.Commit()

	err := tr.Prune(rootHash, data.OldRoot)
	assert.Nil(t, err)

	for i := range oldHashes {
		encNode, err := tr.db.Get(oldHashes[i])
		assert.Nil(t, encNode)
		assert.NotNil(t, err)
	}
}

func TestTrieResetOldHashes(t *testing.T) {
	t.Parallel()

	msh, hsh := getTestMarshAndHasher()
	evictionWaitListSize := 100
	evictionWaitList, _ := mock.NewEvictionWaitingList(evictionWaitListSize, mock.NewMemDbMock(), msh)

	tr := &patriciaMerkleTrie{
		db:                    mock.NewMemDbMock(),
		dbEvictionWaitingList: evictionWaitList,
		oldHashes:             make([][]byte, 0),
		oldRoot:               make([]byte, 0),
		marshalizer:           msh,
		hasher:                hsh,
	}

	_ = tr.Update([]byte("doe"), []byte("reindeer"))
	_ = tr.Update([]byte("dog"), []byte("puppy"))
	_ = tr.Update([]byte("dogglesworth"), []byte("cat"))
	_ = tr.Commit()

	_ = tr.Update([]byte("doeee"), []byte("value of doeee"))

	assert.NotEqual(t, 0, len(tr.oldHashes))
	assert.NotEqual(t, 0, len(tr.oldRoot))

	expectedHashes := tr.oldHashes
	hashes := tr.ResetOldHashes()
	assert.Equal(t, expectedHashes, hashes)
	assert.Equal(t, 0, len(tr.oldHashes))
	assert.Equal(t, 0, len(tr.oldRoot))
}

func TestTrieAddHashesToOldHashes(t *testing.T) {
	t.Parallel()

	msh, hsh := getTestMarshAndHasher()
	evictionWaitListSize := 100
	evictionWaitList, _ := mock.NewEvictionWaitingList(evictionWaitListSize, mock.NewMemDbMock(), msh)
	hashes := [][]byte{[]byte("one"), []byte("two"), []byte("three")}

	tr := &patriciaMerkleTrie{
		db:                    mock.NewMemDbMock(),
		dbEvictionWaitingList: evictionWaitList,
		oldHashes:             make([][]byte, 0),
		oldRoot:               make([]byte, 0),
		marshalizer:           msh,
		hasher:                hsh,
	}

	_ = tr.Update([]byte("doe"), []byte("reindeer"))
	_ = tr.Update([]byte("dog"), []byte("puppy"))
	_ = tr.Update([]byte("dogglesworth"), []byte("cat"))
	_ = tr.Commit()

	_ = tr.Update([]byte("doeee"), []byte("value of doeee"))

	expectedHLength := len(tr.oldHashes) + len(hashes)
	tr.AppendToOldHashes(hashes)
	assert.Equal(t, expectedHLength, len(tr.oldHashes))
}

func TestRecreateTrieFromSnapshotDb(t *testing.T) {
	t.Parallel()

	testVals := []struct {
		key   []byte
		value []byte
	}{
		{[]byte("doe"), []byte("reindeer")},
		{[]byte("dog"), []byte("puppy")},
		{[]byte("dogglesworth"), []byte("cat")},
	}

	msh, hsh := getTestMarshAndHasher()
	evictionWaitListSize := 100
	evictionWaitList, _ := mock.NewEvictionWaitingList(evictionWaitListSize, mock.NewMemDbMock(), msh)

	tempDir, _ := ioutil.TempDir("", "leveldb_temp")
	cfg := config.DBConfig{
		FilePath:          tempDir,
		Type:              string(storageUnit.LvlDbSerial),
		BatchDelaySeconds: 1,
		MaxBatchSize:      1,
		MaxOpenFiles:      10,
	}

	tr := &patriciaMerkleTrie{
		db:                    mock.NewMemDbMock(),
		snapshots:             make([]data.DBWriteCacher, 0),
		snapshotDbCfg:         cfg,
		dbEvictionWaitingList: evictionWaitList,
		marshalizer:           msh,
		hasher:                hsh,
	}

	for _, testVal := range testVals {
		_ = tr.Update(testVal.key, testVal.value)
	}

	_ = tr.Snapshot()
	collapsedRoot, _ := tr.root.getCollapsed(tr.marshalizer, tr.hasher)

	snapshotTrie := &patriciaMerkleTrie{
		root:        collapsedRoot,
		db:          tr.snapshots[0],
		marshalizer: msh,
		hasher:      hsh,
	}

	for tr.snapshotInProgress {
		time.Sleep(snapshotDelay)
	}

	for _, testVal := range testVals {
		val, err := snapshotTrie.Get(testVal.key)
		assert.Nil(t, err)
		assert.Equal(t, testVal.value, val)
	}
}

func TestEachSnapshotCreatesOwnDatabase(t *testing.T) {
	t.Parallel()

	testVals := []struct {
		key   []byte
		value []byte
	}{
		{[]byte("doe"), []byte("reindeer")},
		{[]byte("dog"), []byte("puppy")},
		{[]byte("dogglesworth"), []byte("cat")},
	}

	msh, hsh := getTestMarshAndHasher()
	evictionWaitListSize := 100
	evictionWaitList, _ := mock.NewEvictionWaitingList(evictionWaitListSize, mock.NewMemDbMock(), msh)

	tempDir, _ := ioutil.TempDir("", "leveldb_temp")
	cfg := config.DBConfig{
		FilePath:          tempDir,
		Type:              string(storageUnit.LvlDbSerial),
		BatchDelaySeconds: 1,
		MaxBatchSize:      1,
		MaxOpenFiles:      10,
	}

	tr := &patriciaMerkleTrie{
		db:                    mock.NewMemDbMock(),
		snapshots:             make([]data.DBWriteCacher, 0),
		snapshotId:            0,
		snapshotDbCfg:         cfg,
		dbEvictionWaitingList: evictionWaitList,
		marshalizer:           msh,
		hasher:                hsh,
	}

	for _, testVal := range testVals {
		_ = tr.Update(testVal.key, testVal.value)
		_ = tr.Snapshot()
		for tr.snapshotInProgress {
			time.Sleep(snapshotDelay)
		}

		snapshotId := strconv.Itoa(tr.snapshotId - 1)
		snapshotPath := path.Join(tr.snapshotDbCfg.FilePath, snapshotId)
		f, _ := os.Stat(snapshotPath)
		assert.True(t, f.IsDir())
	}

	assert.Equal(t, len(testVals), tr.snapshotId)
}

func TestDeleteOldSnapshots(t *testing.T) {
	t.Parallel()

	testVals := []struct {
		key   []byte
		value []byte
	}{
		{[]byte("doe"), []byte("reindeer")},
		{[]byte("dog"), []byte("puppy")},
		{[]byte("dogglesworth"), []byte("cat")},
		{[]byte("horse"), []byte("mustang")},
	}

	msh, hsh := getTestMarshAndHasher()
	evictionWaitListSize := 100
	evictionWaitList, _ := mock.NewEvictionWaitingList(evictionWaitListSize, mock.NewMemDbMock(), msh)

	tempDir, _ := ioutil.TempDir("", "leveldb_temp")
	cfg := config.DBConfig{
		FilePath:          tempDir,
		Type:              string(storageUnit.LvlDbSerial),
		BatchDelaySeconds: 1,
		MaxBatchSize:      1,
		MaxOpenFiles:      10,
	}

	tr := &patriciaMerkleTrie{
		db:                    mock.NewMemDbMock(),
		snapshots:             make([]data.DBWriteCacher, 0),
		snapshotId:            0,
		snapshotDbCfg:         cfg,
		dbEvictionWaitingList: evictionWaitList,
		marshalizer:           msh,
		hasher:                hsh,
	}

	for _, testVal := range testVals {
		_ = tr.Update(testVal.key, testVal.value)
		_ = tr.Snapshot()
		for tr.snapshotInProgress {
			time.Sleep(snapshotDelay)
		}
	}

	snapshots, _ := ioutil.ReadDir(tr.snapshotDbCfg.FilePath)
	assert.Equal(t, 2, len(snapshots))
	assert.Equal(t, "2", snapshots[0].Name())
	assert.Equal(t, "3", snapshots[1].Name())
}

func TestNode_getDirtyHashes(t *testing.T) {
	t.Parallel()

	testVals := []struct {
		key   []byte
		value []byte
	}{
		{[]byte("doe"), []byte("reindeer")},
		{[]byte("dog"), []byte("puppy")},
		{[]byte("dogglesworth"), []byte("cat")},
	}

	msh, hsh := getTestMarshAndHasher()
	evictionWaitList := &mock.EvictionWaitingList{
		Cache:       make(map[string][][]byte),
		CacheSize:   100,
		Db:          mock.NewMemDbMock(),
		Marshalizer: msh,
	}

	tr := &patriciaMerkleTrie{
		db:                    mock.NewMemDbMock(),
		dbEvictionWaitingList: evictionWaitList,
		oldHashes:             make([][]byte, 0),
		oldRoot:               make([]byte, 0),
		marshalizer:           msh,
		hasher:                hsh,
	}

	for _, testVal := range testVals {
		_ = tr.Update(testVal.key, testVal.value)
	}

	hashes, err := tr.root.getDirtyHashes()
	assert.Nil(t, err)
	assert.NotNil(t, hashes)
	assert.Equal(t, 6, len(hashes))
}

func TestPruningAndPruningCancellingOnTrieRollback(t *testing.T) {
	t.Parallel()

	testVals := []struct {
		key   []byte
		value []byte
	}{
		{[]byte("doe"), []byte("reindeer")},
		{[]byte("dog"), []byte("puppy")},
		{[]byte("dogglesworth"), []byte("cat")},
		{[]byte("horse"), []byte("stallion")},
	}

	msh, hsh := getTestMarshAndHasher()
	evictionWaitList := &mock.EvictionWaitingList{
		Cache:       make(map[string][][]byte),
		CacheSize:   100,
		Db:          mock.NewMemDbMock(),
		Marshalizer: msh,
	}

	tr := &patriciaMerkleTrie{
		db:                    mock.NewMemDbMock(),
		dbEvictionWaitingList: evictionWaitList,
		oldHashes:             make([][]byte, 0),
		oldRoot:               make([]byte, 0),
		marshalizer:           msh,
		hasher:                hsh,
	}

	rootHashes := make([][]byte, 0)
	rootHashes = append(rootHashes)
	for _, testVal := range testVals {
		_ = tr.Update(testVal.key, testVal.value)
		_ = tr.Commit()
		rootHashes = append(rootHashes, tr.root.getHash())
	}

	for i := 0; i < len(rootHashes); i++ {
		_, err := tr.Recreate(rootHashes[i])
		assert.Nil(t, err)
	}

	tr.CancelPrune(rootHashes[0], data.NewRoot)
	finalizeTrieState(t, 1, tr, rootHashes)
	finalizeTrieState(t, 2, tr, rootHashes)
	rollbackTrieState(t, 3, tr, rootHashes)

	_, err := tr.Recreate(rootHashes[2])
	assert.Nil(t, err)
}

func finalizeTrieState(t *testing.T, index int, tr data.Trie, rootHashes [][]byte) {
	err := tr.Prune(rootHashes[index-1], data.OldRoot)
	assert.Nil(t, err)
	tr.CancelPrune(rootHashes[index], data.NewRoot)

	_, err = tr.Recreate(rootHashes[index-1])
	assert.NotNil(t, err)
}

func rollbackTrieState(t *testing.T, index int, tr data.Trie, rootHashes [][]byte) {
	err := tr.Prune(rootHashes[index], data.NewRoot)
	assert.Nil(t, err)
	tr.CancelPrune(rootHashes[index-1], data.OldRoot)

	_, err = tr.Recreate(rootHashes[index])
	assert.NotNil(t, err)
}

func TestPruningIsBufferedWhileSnapshoting(t *testing.T) {
	t.Parallel()

	nrVals := 100000
	index := 0
	var rootHashes [][]byte

	msh, hsh := getTestMarshAndHasher()
	evictionWaitListSize := 100
	evictionWaitList, _ := mock.NewEvictionWaitingList(evictionWaitListSize, mock.NewMemDbMock(), msh)

	tempDir, _ := ioutil.TempDir("", "leveldb_temp")
	cfg := config.DBConfig{
		FilePath:          tempDir,
		Type:              string(storageUnit.LvlDbSerial),
		BatchDelaySeconds: 1,
		MaxBatchSize:      40000,
		MaxOpenFiles:      10,
	}

	tr := &patriciaMerkleTrie{
		db:                    mock.NewMemDbMock(),
		snapshots:             make([]data.DBWriteCacher, 0),
		snapshotDbCfg:         cfg,
		dbEvictionWaitingList: evictionWaitList,
		marshalizer:           msh,
		hasher:                hsh,
	}

	for i := 0; i < nrVals; i++ {
		_ = tr.Update(tr.hasher.Compute(strconv.Itoa(index)), tr.hasher.Compute(strconv.Itoa(index)))
		index++
	}

	_ = tr.Commit()
	rootHash := tr.root.getHash()
	rootHashes = append(rootHashes, rootHash)
	_ = tr.Snapshot()

	nrRounds := 10
	nrUpdates := 1000
	for i := 0; i < nrRounds; i++ {
		for j := 0; j < nrUpdates; j++ {
			_ = tr.Update(tr.hasher.Compute(strconv.Itoa(index)), tr.hasher.Compute(strconv.Itoa(index)))
			index++
		}
		_ = tr.Commit()

		previousRootHashIndex := len(rootHashes) - 1
		currentRootHash := tr.root.getHash()

		_ = tr.Prune(rootHashes[previousRootHashIndex], data.OldRoot)
		_ = tr.Prune(currentRootHash, data.NewRoot)
		rootHashes = append(rootHashes, currentRootHash)
	}
	assert.Equal(t, 21, len(evictionWaitList.Cache))
	assert.NotEqual(t, 0, len(tr.pruningBuffer))

	for len(tr.pruningBuffer) != 0 {
		time.Sleep(snapshotDelay)
	}

	for i := range rootHashes {
		trie, err := tr.Recreate(rootHashes[i])
		assert.Nil(t, trie)
		assert.NotNil(t, err)
	}

	time.Sleep(batchDelay)
	val, err := tr.snapshots[0].Get(rootHash)
	assert.NotNil(t, val)
	assert.Nil(t, err)
}
