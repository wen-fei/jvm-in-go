package classfile

// label class、interface、field or method which is deprecated
type DeprecatedAttribute struct {
	MarkerAttribute
}

// label file which is not exists int source file and produce by translater
type SyntheticAttribute struct {
	MarkerAttribute
}

type MarkerAttribute struct{}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}
