package memory

func LittleEndian(address [2]byte) uint16 {
	return uint16(address[1])<<8 | uint16(address[0])
}
