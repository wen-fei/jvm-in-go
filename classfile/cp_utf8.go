package classfile

type ConstantUtf8Info struct {
	str string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readeBytes(length)
	// bytes -> str in Go
	self.str = decodeMUTF8(bytes)
}

// simply function of MTF8 -> UTF8
func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}
