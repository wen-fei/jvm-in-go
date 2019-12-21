package heap

import (
	"github.com/jvm-in-go/classfile"
)

type Class struct {
	accessFlags       uint16
	name              string
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        *Slots
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}

// 如果类D想访问类C，需要满足两个条件之一：C是
// public，或者C和D在同一个运行时包内。
func (self *Class) isAccessibleTo(d *Class) bool {
	return self.IsPublic() || self.getPackageName() == d.getPackageName()
}

func (self *Class) getPackageName() string {
	return self.name
}

func (self *Class) isSubClassOf(c *Class) bool {
	return self.superClass == c
}

func (self *Class) NewObject() *Object {
	return newObject(self)
}

func (self *Class) IsInterface() bool {
	return 0 != self.accessFlags&ACC_INTERFACE
}
func (self *Class) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}

func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}
