package main

import (
	"fmt"
	"jvm-in-go/classfile"
	"jvm-in-go/classpath"
	"strings"
)

// cmd test code
// run go install jvm-in-go will create jvm-in-go.exe in classpath/bin/
func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Printf("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	//fmt.Printf("classpath: %s class: %s args: %v\n",
	//	cmd.cpOption, cmd.class, cmd.args)

	//cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	//fmt.Printf("classpath:%v class: %v args:%v\n",
	//	cp, cmd.class, cmd.args)
	//
	//className := strings.Replace(cmd.class, ".", "/", -1)
	//classData, _, err := cp.ReadClass(className)
	//if err != nil {
	//	fmt.Printf("Could not find or load main class %s\n", cmd.class)
	//	return
	//}
	//fmt.Printf("class data:%v\n", classData)

	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(className, cp)
	fmt.Println(className, cp)
	printClassInfo(cf)
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constans count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: %v\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interface: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))

	for _, f := range cf.Fields() {
		fmt.Printf("    %s\n", f.Name())
	}
	fmt.Printf("Methods count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf("    %s\n", m.Name())
	}
}
