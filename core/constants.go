package core

// HeaderType defines the type to be used for the header that is sent
type HeaderType string

const (
	// MetaHeader defines the type of *block.MetaBlock
	MetaHeader HeaderType = "MetaBlock"
	// ShardHeaderV1 defines the type of *block.Header
	ShardHeaderV1 HeaderType = "Header"
	// ShardHeaderV2 defines the type of *block.HeaderV2
	ShardHeaderV2 HeaderType = "HeaderV2"
)

// NodeType represents the node's role in the network
type NodeType string

// NodeTypeObserver signals that a node is running as observer node
const NodeTypeObserver NodeType = "observer"

// NodeTypeValidator signals that a node is running as validator node
const NodeTypeValidator NodeType = "validator"

// pkPrefixSize specifies the max numbers of chars to be displayed from one publc key
const pkPrefixSize = 12

// FileModeUserReadWrite represents the permission for a file which allows the user for reading and writing
const FileModeUserReadWrite = 0600

// FileModeReadWrite represents the permission for a file which allows reading and writing for user and group and read
// for others
const FileModeReadWrite = 0664

// MetachainShardId will be used to identify a shard ID as metachain
const MetachainShardId = uint32(0xFFFFFFFF)

// AllShardId will be used to identify that a message is for all shards
const AllShardId = uint32(0xFFFFFFF0)

// MegabyteSize represents the size in bytes of a megabyte
const MegabyteSize = 1024 * 1024

// MaxMachineIDLen is the maximum machine ID length
const MaxMachineIDLen = 10

// BuiltInFunctionClaimDeveloperRewards is the key for the claim developer rewards built-in function
const BuiltInFunctionClaimDeveloperRewards = "ClaimDeveloperRewards"

// BuiltInFunctionChangeOwnerAddress is the key for the change owner built in function built-in function
const BuiltInFunctionChangeOwnerAddress = "ChangeOwnerAddress"

// BuiltInFunctionSetUserName is the key for the set user name built-in function
const BuiltInFunctionSetUserName = "SetUserName"

// BuiltInFunctionSaveKeyValue is the key for the save key value built-in function
const BuiltInFunctionSaveKeyValue = "SaveKeyValue"

// BuiltInFunctionESDTTransfer is the key for the electronic standard digital token transfer built-in function
const BuiltInFunctionESDTTransfer = "ESDTTransfer"

// BuiltInFunctionESDTBurn is the key for the electronic standard digital token burn built-in function
const BuiltInFunctionESDTBurn = "ESDTBurn"

// BuiltInFunctionESDTFreeze is the key for the electronic standard digital token freeze built-in function
const BuiltInFunctionESDTFreeze = "ESDTFreeze"

// BuiltInFunctionESDTUnFreeze is the key for the electronic standard digital token unfreeze built-in function
const BuiltInFunctionESDTUnFreeze = "ESDTUnFreeze"

// BuiltInFunctionESDTWipe is the key for the electronic standard digital token wipe built-in function
const BuiltInFunctionESDTWipe = "ESDTWipe"

// BuiltInFunctionESDTPause is the key for the electronic standard digital token pause built-in function
const BuiltInFunctionESDTPause = "ESDTPause"

// BuiltInFunctionESDTUnPause is the key for the electronic standard digital token unpause built-in function
const BuiltInFunctionESDTUnPause = "ESDTUnPause"

// BuiltInFunctionSetESDTRole is the key for the electronic standard digital token set built-in function
const BuiltInFunctionSetESDTRole = "ESDTSetRole"

// BuiltInFunctionUnSetESDTRole is the key for the electronic standard digital token unset built-in function
const BuiltInFunctionUnSetESDTRole = "ESDTUnSetRole"

// BuiltInFunctionESDTSetLimitedTransfer is the key for the electronic standard digital token built-in function which sets the property
// for the token to be transferable only through accounts with transfer roles
const BuiltInFunctionESDTSetLimitedTransfer = "ESDTSetLimitedTransfer"

// BuiltInFunctionESDTUnSetLimitedTransfer is the key for the electronic standard digital token built-in function which unsets the property
// for the token to be transferable only through accounts with transfer roles
const BuiltInFunctionESDTUnSetLimitedTransfer = "ESDTUnSetLimitedTransfer"

// BuiltInFunctionESDTLocalMint is the key for the electronic standard digital token local mint built-in function
const BuiltInFunctionESDTLocalMint = "ESDTLocalMint"

// BuiltInFunctionESDTLocalBurn is the key for the electronic standard digital token local burn built-in function
const BuiltInFunctionESDTLocalBurn = "ESDTLocalBurn"

// BuiltInFunctionESDTNFTTransfer is the key for the electronic standard digital token NFT transfer built-in function
const BuiltInFunctionESDTNFTTransfer = "ESDTNFTTransfer"

// BuiltInFunctionESDTNFTCreate is the key for the electronic standard digital token NFT create built-in function
const BuiltInFunctionESDTNFTCreate = "ESDTNFTCreate"

// BuiltInFunctionESDTNFTAddQuantity is the key for the electronic standard digital token NFT add quantity built-in function
const BuiltInFunctionESDTNFTAddQuantity = "ESDTNFTAddQuantity"

// BuiltInFunctionESDTNFTCreateRoleTransfer is the key for the electronic standard digital token create role transfer function
const BuiltInFunctionESDTNFTCreateRoleTransfer = "ESDTNFTCreateRoleTransfer"

// BuiltInFunctionESDTNFTBurn is the key for the electronic standard digital token NFT burn built-in function
const BuiltInFunctionESDTNFTBurn = "ESDTNFTBurn"

// BuiltInFunctionESDTNFTAddURI is the key for the electronic standard digital token NFT add URI built-in function
const BuiltInFunctionESDTNFTAddURI = "ESDTNFTAddURI"

// BuiltInFunctionESDTNFTUpdateAttributes is the key for the electronic standard digital token NFT update attributes built-in function
const BuiltInFunctionESDTNFTUpdateAttributes = "ESDTNFTUpdateAttributes"

// BuiltInFunctionMultiESDTNFTTransfer is the key for the electronic standard digital token multi transfer built-in function
const BuiltInFunctionMultiESDTNFTTransfer = "MultiESDTNFTTransfer"

// BuiltInFunctionSetGuardian is the key for setting a guardian built-in function
const BuiltInFunctionSetGuardian = "SetGuardian"

// BuiltInFunctionGuardAccount is the built-in function key for guarding an account
const BuiltInFunctionGuardAccount = "GuardAccount"

// BuiltInFunctionUnGuardAccount is the built-in function key for un-guarding an account
const BuiltInFunctionUnGuardAccount = "UnGuardAccount"

// BuiltInFunctionMigrateDataTrie is the built-in function key for migrating the data trie
const BuiltInFunctionMigrateDataTrie = "MigrateDataTrie"

// ESDTRoleLocalMint is the constant string for the local role of mint for ESDT tokens
const ESDTRoleLocalMint = "ESDTRoleLocalMint"

// ESDTRoleLocalBurn is the constant string for the local role of burn for ESDT tokens
const ESDTRoleLocalBurn = "ESDTRoleLocalBurn"

// ESDTRoleNFTCreate is the constant string for the local role of create for ESDT NFT tokens
const ESDTRoleNFTCreate = "ESDTRoleNFTCreate"

// ESDTRoleNFTCreateMultiShard is the constant string for the local role of create for ESDT NFT tokens multishard
const ESDTRoleNFTCreateMultiShard = "ESDTRoleNFTCreateMultiShard"

// ESDTRoleNFTAddQuantity is the constant string for the local role of adding quantity for existing ESDT NFT tokens
const ESDTRoleNFTAddQuantity = "ESDTRoleNFTAddQuantity"

// ESDTRoleNFTBurn is the constant string for the local role of burn for ESDT NFT tokens
const ESDTRoleNFTBurn = "ESDTRoleNFTBurn"

// ESDTRoleNFTAddURI is the constant string for the local role of adding a URI for ESDT NFT tokens
const ESDTRoleNFTAddURI = "ESDTRoleNFTAddURI"

// ESDTRoleNFTUpdateAttributes is the constant string for the local role of updating attributes for ESDT NFT tokens
const ESDTRoleNFTUpdateAttributes = "ESDTRoleNFTUpdateAttributes"

// ESDTRoleTransfer is the constant string for the local role to transfer ESDT, only for special tokens
const ESDTRoleTransfer = "ESDTTransferRole"

// ESDTType defines the possible types in case of ESDT tokens
type ESDTType uint32

const (
	// Fungible defines the token type for ESDT fungible tokens
	Fungible ESDTType = iota
	// NonFungible defines the token type for ESDT non fungible tokens
	NonFungible
)

// FungibleESDT defines the string for the token type of fungible ESDT
const FungibleESDT = "FungibleESDT"

// NonFungibleESDT defines the string for the token type of non fungible ESDT
const NonFungibleESDT = "NonFungibleESDT"

// SemiFungibleESDT defines the string for the token type of semi fungible ESDT
const SemiFungibleESDT = "SemiFungibleESDT"

// MaxRoyalty defines 100% as uint32
const MaxRoyalty = uint32(10000)

// RelayedTransaction is the key for the electronic meta/gassless/relayed transaction standard
const RelayedTransaction = "relayedTx"

// RelayedTransactionV2 is the key for the optimized electronic meta/gassless/relayed transaction standard
const RelayedTransactionV2 = "relayedTxV2"

// SCDeployInitFunctionName is the key for the function which is called at smart contract deploy time
const SCDeployInitFunctionName = "_init"

// ProtectedKeyPrefix is the key prefix which is protected from writing in the trie - only for special builtin functions
const ProtectedKeyPrefix = "E" + "L" + "R" + "O" + "N" + "D"

// DelegationSystemSCKey is the key under which there is data in case of system delegation smart contracts
const DelegationSystemSCKey = "delegation"

// ESDTKeyIdentifier is the key prefix for esdt tokens
const ESDTKeyIdentifier = "esdt"

// ESDTRoleIdentifier is the key prefix for esdt role identifier
const ESDTRoleIdentifier = "role"

// ESDTNFTLatestNonceIdentifier is the key prefix for esdt latest nonce identifier
const ESDTNFTLatestNonceIdentifier = "nonce"

// GuardiansKeyIdentifier is the key prefix for guardians
const GuardiansKeyIdentifier = "guardians"

// MaxNumShards represents the maximum number of shards possible in the system
const MaxNumShards = 256

// MinMetaTxExtraGasCost is the constant defined for minimum gas value to be sent in meta transaction
const MinMetaTxExtraGasCost = uint64(1_000_000)

// MaxLeafSize represents maximum amount of data which can be saved under one leaf
const MaxLeafSize = uint64(1 << 26) //64MB

// MaxBufferSizeToSendTrieNodes represents max buffer size to send in bytes used when resolving trie nodes
// Every trie node that has a greater size than this constant is considered a large trie node and should be split in
// smaller chunks
const MaxBufferSizeToSendTrieNodes = 1 << 18 //256KB

// MaxUserNameLength represents the maximum number of bytes a UserName can have
const MaxUserNameLength = 32

// MinLenArgumentsESDTTransfer defines the min length of arguments for the ESDT transfer
const MinLenArgumentsESDTTransfer = 2

// MinLenArgumentsESDTNFTTransfer defines the minimum length for esdt nft transfer
const MinLenArgumentsESDTNFTTransfer = 4

// MaxLenForESDTIssueMint defines the maximum length in bytes for the issued/minted balance
const MaxLenForESDTIssueMint = 100

// BaseOperationCostString represents the field name for base operation costs
const BaseOperationCostString = "BaseOperationCost"

// BuiltInCostString represents the field name for built in operation costs
const BuiltInCostString = "BuiltInCost"

// ESDTSCAddress is the hard-coded address for esdt issuing smart contract
var ESDTSCAddress = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 255, 255}

// SCDeployIdentifier is the identifier for a smart contract deploy
const SCDeployIdentifier = "SCDeploy"

// SCUpgradeIdentifier is the identifier for a smart contract upgrade
const SCUpgradeIdentifier = "SCUpgrade"

// WriteLogIdentifier is the identifier for the information log that is generated by a smart contract call/esdt transfer
const WriteLogIdentifier = "writeLog"

// SignalErrorOperation is the identifier for the log that is generated when a smart contract is executed but return an error
const SignalErrorOperation = "signalError"

// CompletedTxEventIdentifier is the identifier for the log that is generated when the execution of a smart contract call is done
const CompletedTxEventIdentifier = "completedTxEvent"

// InternalVMErrorsOperation is the identifier for the log that is generated when the execution of a smart contract generates an interval vm error
const InternalVMErrorsOperation = "internalVMErrors"

// GasRefundForRelayerMessage is the return message for to the smart contract result with refund for the relayer
const GasRefundForRelayerMessage = "gas refund for relayer"

// InitialVersionOfTransaction defines the initial version for a transaction
const InitialVersionOfTransaction = uint32(1)

// DefaultAddressPrefix is the default hrp of MultiversX/Elrond
const DefaultAddressPrefix = "erd"

// TopicRequestSuffix represents the topic name suffix for requests
const TopicRequestSuffix = "_REQUEST"
