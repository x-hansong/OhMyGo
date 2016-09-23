package classfile

/*
Deprecated_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/
// 标记该属性被废除，不建议使用
type DeprecatedAttribute struct {
	MarkerAttribute
}

/*
Synthetic_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/
// 用来标记源文件中不存在，由编译器生成的类成员，用于支持嵌套类和嵌套接口
type SyntheticAttribute struct {
	MarkerAttribute
}

type MarkerAttribute struct {

}

func (self *MarkerAttribute) readInfo(reader *ClassReader)  {
	// 标记属性没有值
}
