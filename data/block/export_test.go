package block

func (m *MiniBlockHeader) SetMiniBlockHeaderReserved(mbhr *MiniBlockHeaderReserved) error {
	return m.setMiniBlockHeaderReserved(mbhr)
}

func (m *MiniBlockHeader) GetMiniBlockHeaderReserved() (*MiniBlockHeaderReserved, error) {
	return m.getMiniBlockHeaderReserved()
}
