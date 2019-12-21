package base

type BytecodeReader struct {
	// store the class bytes
	code []byte
	pc   int
}

func (self *BytecodeReader) Reset(code []byte, pc int) {
	self.code = code
	self.pc = pc
}

func (self *BytecodeReader) Code() []byte {
	return self.code
}

func (self *BytecodeReader) PC() int {
	return self.pc
}

func (self *BytecodeReader) ReadInt8() int8 {
	return self.ReadInt8()
}

func (self *BytecodeReader) ReadUint8() uint8 {
	return uint8(self.ReadInt8())
}

func (self *BytecodeReader) ReadInt16() uint16 {
	byte1 := uint16(self.ReadInt8())
	byte2 := uint16(self.ReadInt8())
	return (byte1 << 8) | byte2
}

func (self *BytecodeReader) ReadInt32() int32 {
	byte1 := int32(self.ReadInt8())
	byte2 := int32(self.ReadInt8())
	byte3 := int32(self.ReadInt8())
	byte4 := int32(self.ReadInt8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}

// 0-3 字节的padding，保证defaultOffset在字节码中的地址是4的倍数
func (self *BytecodeReader) SkipPadding() {
	for self.pc%4 != 0 {
		self.ReadInt8()
	}
}

func (self *BytecodeReader) ReadInt32s(n int32) []int32 {
	ints := make([]int32, n)
	for i := range ints {
		ints[i] = self.ReadInt32()
	}
	return ints
}
