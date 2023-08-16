package core

import (
	"github.com/multiversx/mx-chain-core-go/core/check"
)

// EnableEpochFlag defines a flag specific to the enableEpochs.toml
type EnableEpochFlag string

const (
	SCDeployFlag                                       EnableEpochFlag = "SCDeployFlag"
	BuiltInFunctionsFlag                               EnableEpochFlag = "BuiltInFunctionsFlag"
	RelayedTransactionsFlag                            EnableEpochFlag = "RelayedTransactionsFlag"
	PenalizedTooMuchGasFlag                            EnableEpochFlag = "PenalizedTooMuchGasFlag"
	SwitchJailWaitingFlag                              EnableEpochFlag = "SwitchJailWaitingFlag"
	BelowSignedThresholdFlag                           EnableEpochFlag = "BelowSignedThresholdFlag"
	SwitchHysteresisForMinNodesFlagInSpecificEpochOnly EnableEpochFlag = "SwitchHysteresisForMinNodesFlagInSpecificEpochOnly"
	TransactionSignedWithTxHashFlag                    EnableEpochFlag = "TransactionSignedWithTxHashFlag"
	MetaProtectionFlag                                 EnableEpochFlag = "MetaProtectionFlag"
	AheadOfTimeGasUsageFlag                            EnableEpochFlag = "AheadOfTimeGasUsageFlag"
	GasPriceModifierFlag                               EnableEpochFlag = "GasPriceModifierFlag"
	RepairCallbackFlag                                 EnableEpochFlag = "RepairCallbackFlag"
	ReturnDataToLastTransferFlagAfterEpoch             EnableEpochFlag = "ReturnDataToLastTransferFlagAfterEpoch"
	SenderInOutTransferFlag                            EnableEpochFlag = "SenderInOutTransferFlag"
	StakeFlag                                          EnableEpochFlag = "StakeFlag"
	StakingV2Flag                                      EnableEpochFlag = "StakingV2Flag"
	StakingV2OwnerFlagInSpecificEpochOnly              EnableEpochFlag = "StakingV2OwnerFlagInSpecificEpochOnly"
	StakingV2FlagAfterEpoch                            EnableEpochFlag = "StakingV2FlagAfterEpoch"
	DoubleKeyProtectionFlag                            EnableEpochFlag = "DoubleKeyProtectionFlag"
	ESDTFlag                                           EnableEpochFlag = "ESDTFlag"
	ESDTFlagInSpecificEpochOnly                        EnableEpochFlag = "ESDTFlagInSpecificEpochOnly"
	GovernanceFlag                                     EnableEpochFlag = "GovernanceFlag"
	GovernanceFlagInSpecificEpochOnly                  EnableEpochFlag = "GovernanceFlagInSpecificEpochOnly"
	DelegationManagerFlag                              EnableEpochFlag = "DelegationManagerFlag"
	DelegationSmartContractFlag                        EnableEpochFlag = "DelegationSmartContractFlag"
	DelegationSmartContractFlagInSpecificEpochOnly     EnableEpochFlag = "DelegationSmartContractFlagInSpecificEpochOnly"
	CorrectLastUnJailedFlag                            EnableEpochFlag = "CorrectLastUnJailedFlag"
	CorrectLastUnJailedFlagInSpecificEpochOnly         EnableEpochFlag = "CorrectLastUnJailedFlagInSpecificEpochOnly"
	RelayedTransactionsV2Flag                          EnableEpochFlag = "RelayedTransactionsV2Flag"
	UnBondTokensV2Flag                                 EnableEpochFlag = "UnBondTokensV2Flag"
	SaveJailedAlwaysFlag                               EnableEpochFlag = "SaveJailedAlwaysFlag"
	ReDelegateBelowMinCheckFlag                        EnableEpochFlag = "ReDelegateBelowMinCheckFlag"
	ValidatorToDelegationFlag                          EnableEpochFlag = "ValidatorToDelegationFlag"
	IncrementSCRNonceInMultiTransferFlag               EnableEpochFlag = "IncrementSCRNonceInMultiTransferFlag"
	ESDTMultiTransferFlag                              EnableEpochFlag = "ESDTMultiTransferFlag"
	GlobalMintBurnFlag                                 EnableEpochFlag = "GlobalMintBurnFlag"
	ESDTTransferRoleFlag                               EnableEpochFlag = "ESDTTransferRoleFlag"
	BuiltInFunctionOnMetaFlag                          EnableEpochFlag = "BuiltInFunctionOnMetaFlag"
	ComputeRewardCheckpointFlag                        EnableEpochFlag = "ComputeRewardCheckpointFlag"
	SCRSizeInvariantCheckFlag                          EnableEpochFlag = "SCRSizeInvariantCheckFlag"
	BackwardCompSaveKeyValueFlag                       EnableEpochFlag = "BackwardCompSaveKeyValueFlag"
	ESDTNFTCreateOnMultiShardFlag                      EnableEpochFlag = "ESDTNFTCreateOnMultiShardFlag"
	MetaESDTSetFlag                                    EnableEpochFlag = "MetaESDTSetFlag"
	AddTokensToDelegationFlag                          EnableEpochFlag = "AddTokensToDelegationFlag"
	MultiESDTTransferFixOnCallBackFlag                 EnableEpochFlag = "MultiESDTTransferFixOnCallBackFlag"
	OptimizeGasUsedInCrossMiniBlocksFlag               EnableEpochFlag = "OptimizeGasUsedInCrossMiniBlocksFlag"
	CorrectFirstQueuedFlag                             EnableEpochFlag = "CorrectFirstQueuedFlag"
	DeleteDelegatorAfterClaimRewardsFlag               EnableEpochFlag = "DeleteDelegatorAfterClaimRewardsFlag"
	RemoveNonUpdatedStorageFlag                        EnableEpochFlag = "RemoveNonUpdatedStorageFlag"
	OptimizeNFTStoreFlag                               EnableEpochFlag = "OptimizeNFTStoreFlag"
	CreateNFTThroughExecByCallerFlag                   EnableEpochFlag = "CreateNFTThroughExecByCallerFlag"
	StopDecreasingValidatorRatingWhenStuckFlag         EnableEpochFlag = "StopDecreasingValidatorRatingWhenStuckFlag"
	FrontRunningProtectionFlag                         EnableEpochFlag = "FrontRunningProtectionFlag"
	PayableBySCFlag                                    EnableEpochFlag = "PayableBySCFlag"
	CleanUpInformativeSCRsFlag                         EnableEpochFlag = "CleanUpInformativeSCRsFlag"
	StorageAPICostOptimizationFlag                     EnableEpochFlag = "StorageAPICostOptimizationFlag"
	ESDTRegisterAndSetAllRolesFlag                     EnableEpochFlag = "ESDTRegisterAndSetAllRolesFlag"
	ScheduledMiniBlocksFlag                            EnableEpochFlag = "ScheduledMiniBlocksFlag"
	CorrectJailedNotUnStakedEmptyQueueFlag             EnableEpochFlag = "CorrectJailedNotUnStakedEmptyQueueFlag"
	DoNotReturnOldBlockInBlockchainHookFlag            EnableEpochFlag = "DoNotReturnOldBlockInBlockchainHookFlag"
	AddFailedRelayedTxToInvalidMBsFlag                 EnableEpochFlag = "AddFailedRelayedTxToInvalidMBsFlag"
	SCRSizeInvariantOnBuiltInResultFlag                EnableEpochFlag = "SCRSizeInvariantOnBuiltInResultFlag"
	CheckCorrectTokenIDForTransferRoleFlag             EnableEpochFlag = "CheckCorrectTokenIDForTransferRoleFlag"
	FailExecutionOnEveryAPIErrorFlag                   EnableEpochFlag = "FailExecutionOnEveryAPIErrorFlag"
	MiniBlockPartialExecutionFlag                      EnableEpochFlag = "MiniBlockPartialExecutionFlag"
	ManagedCryptoAPIsFlag                              EnableEpochFlag = "ManagedCryptoAPIsFlag"
	ESDTMetadataContinuousCleanupFlag                  EnableEpochFlag = "ESDTMetadataContinuousCleanupFlag"
	DisableExecByCallerFlag                            EnableEpochFlag = "DisableExecByCallerFlag"
	RefactorContextFlag                                EnableEpochFlag = "RefactorContextFlag"
	CheckFunctionArgumentFlag                          EnableEpochFlag = "CheckFunctionArgumentFlag"
	CheckExecuteOnReadOnlyFlag                         EnableEpochFlag = "CheckExecuteOnReadOnlyFlag"
	SetSenderInEeiOutputTransferFlag                   EnableEpochFlag = "SetSenderInEeiOutputTransferFlag"
	FixAsyncCallbackCheckFlag                          EnableEpochFlag = "FixAsyncCallbackCheckFlag"
	SaveToSystemAccountFlag                            EnableEpochFlag = "SaveToSystemAccountFlag"
	CheckFrozenCollectionFlag                          EnableEpochFlag = "CheckFrozenCollectionFlag"
	SendAlwaysFlag                                     EnableEpochFlag = "SendAlwaysFlag"
	ValueLengthCheckFlag                               EnableEpochFlag = "ValueLengthCheckFlag"
	CheckTransferFlag                                  EnableEpochFlag = "CheckTransferFlag"
	TransferToMetaFlag                                 EnableEpochFlag = "TransferToMetaFlag"
	ESDTNFTImprovementV1Flag                           EnableEpochFlag = "ESDTNFTImprovementV1Flag"
	ChangeDelegationOwnerFlag                          EnableEpochFlag = "ChangeDelegationOwnerFlag"
	RefactorPeersMiniBlocksFlag                        EnableEpochFlag = "RefactorPeersMiniBlocksFlag"
	SCProcessorV2Flag                                  EnableEpochFlag = "SCProcessorV2Flag"
	FixAsyncCallBackArgsListFlag                       EnableEpochFlag = "FixAsyncCallBackArgsListFlag"
	FixOldTokenLiquidityFlag                           EnableEpochFlag = "FixOldTokenLiquidityFlag"
	RuntimeMemStoreLimitFlag                           EnableEpochFlag = "RuntimeMemStoreLimitFlag"
	RuntimeCodeSizeFixFlag                             EnableEpochFlag = "RuntimeCodeSizeFixFlag"
	MaxBlockchainHookCountersFlag                      EnableEpochFlag = "MaxBlockchainHookCountersFlag"
	WipeSingleNFTLiquidityDecreaseFlag                 EnableEpochFlag = "WipeSingleNFTLiquidityDecreaseFlag"
	AlwaysSaveTokenMetaDataFlag                        EnableEpochFlag = "AlwaysSaveTokenMetaDataFlag"
	SetGuardianFlag                                    EnableEpochFlag = "SetGuardianFlag"
	RelayedNonceFixFlag                                EnableEpochFlag = "RelayedNonceFixFlag"
	ConsistentTokensValuesLengthCheckFlag              EnableEpochFlag = "ConsistentTokensValuesLengthCheckFlag"
	KeepExecOrderOnCreatedSCRsFlag                     EnableEpochFlag = "KeepExecOrderOnCreatedSCRsFlag"
	MultiClaimOnDelegationFlag                         EnableEpochFlag = "MultiClaimOnDelegationFlag"
	ChangeUsernameFlag                                 EnableEpochFlag = "ChangeUsernameFlag"
	AutoBalanceDataTriesFlag                           EnableEpochFlag = "AutoBalanceDataTriesFlag"
	FixDelegationChangeOwnerOnAccountFlag              EnableEpochFlag = "FixDelegationChangeOwnerOnAccountFlag"
	FixOOGReturnCodeFlag                               EnableEpochFlag = "FixOOGReturnCodeFlag"
	DeterministicSortOnValidatorsInfoFixFlag           EnableEpochFlag = "DeterministicSortOnValidatorsInfoFixFlag"
)

// CheckHandlerCompatibility checks if the provided handler is compatible with this mx-chain-core-go version
func CheckHandlerCompatibility(handler EnableEpochsHandler) error {
	if check.IfNil(handler) {
		return ErrNilEnableEpochsHandler
	}

	// allFlags slice must contain all flags defined above
	allFlags := []EnableEpochFlag{
		SCDeployFlag,
		BuiltInFunctionsFlag,
		RelayedTransactionsFlag,
		PenalizedTooMuchGasFlag,
		SwitchJailWaitingFlag,
		BelowSignedThresholdFlag,
		SwitchHysteresisForMinNodesFlagInSpecificEpochOnly,
		TransactionSignedWithTxHashFlag,
		MetaProtectionFlag,
		AheadOfTimeGasUsageFlag,
		GasPriceModifierFlag,
		RepairCallbackFlag,
		ReturnDataToLastTransferFlagAfterEpoch,
		SenderInOutTransferFlag,
		StakeFlag,
		StakingV2Flag,
		StakingV2OwnerFlagInSpecificEpochOnly,
		StakingV2FlagAfterEpoch,
		DoubleKeyProtectionFlag,
		ESDTFlag,
		ESDTFlagInSpecificEpochOnly,
		GovernanceFlag,
		GovernanceFlagInSpecificEpochOnly,
		DelegationManagerFlag,
		DelegationSmartContractFlag,
		DelegationSmartContractFlagInSpecificEpochOnly,
		CorrectLastUnJailedFlagInSpecificEpochOnly,
		CorrectLastUnJailedFlag,
		RelayedTransactionsV2Flag,
		UnBondTokensV2Flag,
		SaveJailedAlwaysFlag,
		ReDelegateBelowMinCheckFlag,
		ValidatorToDelegationFlag,
		IncrementSCRNonceInMultiTransferFlag,
		ESDTMultiTransferFlag,
		GlobalMintBurnFlag,
		ESDTTransferRoleFlag,
		BuiltInFunctionOnMetaFlag,
		ComputeRewardCheckpointFlag,
		SCRSizeInvariantCheckFlag,
		BackwardCompSaveKeyValueFlag,
		ESDTNFTCreateOnMultiShardFlag,
		MetaESDTSetFlag,
		AddTokensToDelegationFlag,
		MultiESDTTransferFixOnCallBackFlag,
		OptimizeGasUsedInCrossMiniBlocksFlag,
		CorrectFirstQueuedFlag,
		DeleteDelegatorAfterClaimRewardsFlag,
		RemoveNonUpdatedStorageFlag,
		OptimizeNFTStoreFlag,
		CreateNFTThroughExecByCallerFlag,
		StopDecreasingValidatorRatingWhenStuckFlag,
		FrontRunningProtectionFlag,
		PayableBySCFlag,
		CleanUpInformativeSCRsFlag,
		StorageAPICostOptimizationFlag,
		ESDTRegisterAndSetAllRolesFlag,
		ScheduledMiniBlocksFlag,
		CorrectJailedNotUnStakedEmptyQueueFlag,
		DoNotReturnOldBlockInBlockchainHookFlag,
		AddFailedRelayedTxToInvalidMBsFlag,
		SCRSizeInvariantOnBuiltInResultFlag,
		CheckCorrectTokenIDForTransferRoleFlag,
		FailExecutionOnEveryAPIErrorFlag,
		MiniBlockPartialExecutionFlag,
		ManagedCryptoAPIsFlag,
		ESDTMetadataContinuousCleanupFlag,
		DisableExecByCallerFlag,
		RefactorContextFlag,
		CheckFunctionArgumentFlag,
		CheckExecuteOnReadOnlyFlag,
		SetSenderInEeiOutputTransferFlag,
		FixAsyncCallbackCheckFlag,
		SaveToSystemAccountFlag,
		CheckFrozenCollectionFlag,
		SendAlwaysFlag,
		ValueLengthCheckFlag,
		CheckTransferFlag,
		TransferToMetaFlag,
		ESDTNFTImprovementV1Flag,
		ChangeDelegationOwnerFlag,
		RefactorPeersMiniBlocksFlag,
		SCProcessorV2Flag,
		FixAsyncCallBackArgsListFlag,
		FixOldTokenLiquidityFlag,
		RuntimeMemStoreLimitFlag,
		RuntimeCodeSizeFixFlag,
		MaxBlockchainHookCountersFlag,
		WipeSingleNFTLiquidityDecreaseFlag,
		AlwaysSaveTokenMetaDataFlag,
		SetGuardianFlag,
		RelayedNonceFixFlag,
		ConsistentTokensValuesLengthCheckFlag,
		KeepExecOrderOnCreatedSCRsFlag,
		MultiClaimOnDelegationFlag,
		ChangeUsernameFlag,
		AutoBalanceDataTriesFlag,
		FixDelegationChangeOwnerOnAccountFlag,
		FixOOGReturnCodeFlag,
		DeterministicSortOnValidatorsInfoFixFlag,
	}

	for _, flag := range allFlags {
		if !handler.IsFlagDefined(flag) {
			return ErrInvalidEnableEpochsHandler
		}
	}

	return nil
}
