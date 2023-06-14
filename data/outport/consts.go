package outport

const (
	// TopicSaveBlock is the topic that triggers a block saving
	TopicSaveBlock = "SaveBlock"
	// TopicRevertIndexedBlock is the topic that triggers a reverting of an indexed block
	TopicRevertIndexedBlock = "RevertIndexedBlock"
	// TopicSaveRoundsInfo is the topic that triggers the saving of rounds info
	TopicSaveRoundsInfo = "SaveRoundsInfo"
	// TopicSaveValidatorsPubKeys is the topic that triggers the saving of validators' public keys
	TopicSaveValidatorsPubKeys = "SaveValidatorsPubKeys"
	// TopicSaveValidatorsRating is the topic that triggers the saving of the validators' rating
	TopicSaveValidatorsRating = "SaveValidatorsRating"
	// TopicSaveAccounts is the topic that triggers the saving of accounts
	TopicSaveAccounts = "SaveAccounts"
	// TopicFinalizedBlock is the topic that triggers the handling of a finalized block
	TopicFinalizedBlock = "FinalizedBlock"
	// TopicSettings is the topic that triggers the sending of node settings
	TopicSettings = "Settings"
)
