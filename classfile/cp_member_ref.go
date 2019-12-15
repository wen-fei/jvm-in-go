package classfile

type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

func (self *ConstantMemberrefInfo) ClassName() string {
	return self.cp.getClassName(self.classIndex)
}

func (self *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}

// 结构体嵌套模拟继承
// field symbol ref
type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}

// common function symbol ref
type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}

// interface function symbol ref
type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}
