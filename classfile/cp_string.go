package classfile

type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

// read constant pool index
func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}

// find str in constant pool
func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)
}
