package block

import "github.com/ElrondNetwork/elrond-go-core/data"

// GetTypeInt32 gets the miniBlock type
func (m *MiniBlockHeader) GetTypeInt32() int32 {
	if m == nil {
		return -1
	}

	return int32(m.Type)
}

// SetHash sets the miniBlock hash
func (m *MiniBlockHeader) SetHash(hash []byte) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}

	m.Hash = hash

	return nil
}

// SetSenderShardID sets the miniBlock sender shardID
func (m *MiniBlockHeader) SetSenderShardID(shardID uint32) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}

	m.SenderShardID = shardID

	return nil
}

// SetReceiverShardID sets the miniBlock receiver ShardID
func (m *MiniBlockHeader) SetReceiverShardID(shardID uint32) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}

	m.ReceiverShardID = shardID

	return nil
}

// SetTxCount sets the miniBlock txs count
func (m *MiniBlockHeader) SetTxCount(count uint32) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}

	m.TxCount = count

	return nil
}

// SetTypeInt32 sets the miniBlock type
func (m *MiniBlockHeader) SetTypeInt32(t int32) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}

	m.Type = Type(t)

	return nil
}

// SetReserved sets the miniBlock reserved field
func (m *MiniBlockHeader) SetReserved(reserved []byte) error {
	if m == nil {
		return data.ErrNilPointerReceiver
	}

	m.Reserved = reserved

	return nil
}

// GetProcessingType returns the miniBlock processing type as a int32
func (m *MiniBlockHeader) GetProcessingType() int32 {
	miniBlockHeaderReserved, err := m.getMiniBlockHeaderReserved()
	if err != nil || miniBlockHeaderReserved == nil {
		return int32(Normal)
	}

	return int32(miniBlockHeaderReserved.ExecutionType)
}

// SetProcessingType sets the miniBlock processing type
func (m *MiniBlockHeader) SetProcessingType(procType int32) error {
	var err error
	mbhr := &MiniBlockHeaderReserved{}
	if len(m.Reserved) > 0 {
		mbhr, err = m.getMiniBlockHeaderReserved()
		if err != nil {
			return err
		}
	}
	mbhr.ExecutionType = ProcessingType(procType)

	return m.setMiniBlockHeaderReserved(mbhr)
}

// GetConstructionState returns the construction state of the miniBlock
func (m *MiniBlockHeader) GetConstructionState() int32 {
	miniBlockHeaderReserved, err := m.getMiniBlockHeaderReserved()
	if err != nil || miniBlockHeaderReserved == nil {
		return int32(Final)
	}

	return int32(miniBlockHeaderReserved.State)
}

// SetConstructionState sets the construction state of the miniBlock
func (m *MiniBlockHeader) SetConstructionState(state int32) error {
	var err error
	mbhr := &MiniBlockHeaderReserved{}
	if len(m.Reserved) > 0 {
		mbhr, err = m.getMiniBlockHeaderReserved()
		if err != nil {
			return err
		}
	}
	mbhr.State = MiniBlockState(state)

	return m.setMiniBlockHeaderReserved(mbhr)
}

// IsFinal returns true if the miniBlock is final
func (m *MiniBlockHeader) IsFinal() bool {
	state := m.GetConstructionState()

	return state == int32(Final)
}

// GetMiniBlockHeaderReserved returns the MiniBlockHeader Reserved field as a MiniBlockHeaderReserved
func (m *MiniBlockHeader) getMiniBlockHeaderReserved() (*MiniBlockHeaderReserved, error) {
	if len(m.Reserved) > 0 {
		mbhr := &MiniBlockHeaderReserved{}
		err := mbhr.Unmarshal(m.Reserved)
		if err != nil {
			return nil, err
		}

		return mbhr, nil
	}
	return nil, nil
}

// SetMiniBlockHeaderReserved sets the Reserved field for the miniBlock header with the given parameter
func (m *MiniBlockHeader) setMiniBlockHeaderReserved(mbhr *MiniBlockHeaderReserved) error {
	if mbhr == nil {
		m.Reserved = nil
		return nil
	}

	reserved, err := mbhr.Marshal()
	if err != nil {
		return err
	}
	m.Reserved = reserved

	return nil
}

// GetIndexOfFirstTxProcessed returns index of the first transaction processed in the miniBlock
func (m *MiniBlockHeader) GetIndexOfFirstTxProcessed() int32 {
	miniBlockHeaderReserved, err := m.getMiniBlockHeaderReserved()
	if err != nil || miniBlockHeaderReserved == nil {
		return 0
	}

	return miniBlockHeaderReserved.IndexOfFirstTxProcessed
}

// GetIndexOfLastTxProcessed returns index of the last transaction processed in the miniBlock
func (m *MiniBlockHeader) GetIndexOfLastTxProcessed() int32 {
	miniBlockHeaderReserved, err := m.getMiniBlockHeaderReserved()
	if err != nil || miniBlockHeaderReserved == nil {
		return int32(m.TxCount) - 1
	}

	isIndexNotSetByPreviousVersion := miniBlockHeaderReserved.IndexOfLastTxProcessed == 0 &&
		miniBlockHeaderReserved.IndexOfLastTxProcessed < int32(m.TxCount-1) &&
		m.GetConstructionState() != int32(PartialExecuted)
	if isIndexNotSetByPreviousVersion {
		return int32(m.TxCount) - 1
	}

	return miniBlockHeaderReserved.IndexOfLastTxProcessed
}

// SetIndexOfLastTxProcessed sets index of the last transaction processed in the miniBlock
func (m *MiniBlockHeader) SetIndexOfLastTxProcessed(indexOfLastTxProcessed int32) error {
	var err error
	mbhr := &MiniBlockHeaderReserved{}
	if len(m.Reserved) > 0 {
		mbhr, err = m.getMiniBlockHeaderReserved()
		if err != nil {
			return err
		}
	}
	mbhr.IndexOfLastTxProcessed = indexOfLastTxProcessed

	return m.setMiniBlockHeaderReserved(mbhr)
}

// SetIndexOfFirstTxProcessed sets index of the first transaction processed in the miniBlock
func (m *MiniBlockHeader) SetIndexOfFirstTxProcessed(indexOfFirstTxProcessed int32) error {
	var err error
	mbhr := &MiniBlockHeaderReserved{}
	if len(m.Reserved) > 0 {
		mbhr, err = m.getMiniBlockHeaderReserved()
		if err != nil {
			return err
		}
	}
	mbhr.IndexOfFirstTxProcessed = indexOfFirstTxProcessed

	return m.setMiniBlockHeaderReserved(mbhr)
}

// ShallowClone returns the miniBlockHeader swallow clone
func (m *MiniBlockHeader) ShallowClone() data.MiniBlockHeaderHandler {
	if m == nil {
		return nil
	}

	mbhCopy := *m

	return &mbhCopy
}
