package heap

import "github.com/jvm-in-go/classfile"

type Field struct {
	ClassMember
	constValueIndex uint
	slotId          uint
}

func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
	}
	return fields
}

func (self *Field) isLongOrDouble() bool {
	return self.descriptor == "J" || self.descriptor == "D"
}

func (self *Field) SlotId() uint {
	return self.slotId
}

func (self *Field) ConstValueIndex() uint {
	return 0
}

func (self *Field) Descriptor() string {
	return self.descriptor
}
