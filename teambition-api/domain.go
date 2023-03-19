package teambitionapi

// 自定义字段值
type Customfield struct {
	CFID  string  `json:"cfId,omitempty"`  // 自定义字段ID
	Type  string  `json:"type,omitempty"`  // 自定义字段类型
	Value []Value `json:"value,omitempty"` // 字段值集合
}

// 字段值
type Value struct {
	ID         string `json:"id,omitempty"`         // 字段值ID
	MetaString string `json:"metaString,omitempty"` // 字段值元属性
	Title      string `json:"title,omitempty"`      // 字段值内容
}
