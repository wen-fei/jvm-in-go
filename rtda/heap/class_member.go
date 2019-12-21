package heap

import "github.com/jvm-in-go/classfile"

type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
}

func (self *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	self.accessFlags = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descriptor()
}

func (self *ClassMember) IsPublic() bool {
	return 0 != self.class.accessFlags&ACC_PUBLIC
}

func (self *ClassMember) IsProtected() bool {
	return 0 != self.class.accessFlags&ACC_PROTECTED
}

func (self *ClassMember) IsPrivate() bool {
	return 0 != self.class.accessFlags&ACC_PRIVATE
}

func (self *ClassMember) IsStatic() bool {
	return 0 != self.accessFlags&ACC_STATIC
}

func (self *ClassMember) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}

func (self *ClassMember) Name() string {
	return self.name
}
func (self *ClassMember) Descriptor() string {
	return self.descriptor
}
func (self *ClassMember) Class() *Class {
	return self.class
}
