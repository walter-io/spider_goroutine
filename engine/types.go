package engine

// 请求结构
type Request struct {
	Url        string
	ParserFunc ParserFunc
}

// 解析器方法
type ParserFunc func(content []byte) ParserResult

// 解析结果
type ParserResult struct {
	Items    []Details
	Requests []Request
}

// 汽车详情字段
type Details struct {
	Name          string      // 名称
	Rank          string      // 类型
	Engine        string      // 发动机
	PowerType     string      // 动力类型
	Gearbox       string      // 变速箱
	Size          string      // 长宽高
	BodyStructure string      // 车身结构
	ListedTime    string      // 上市时间
	OilWear       interface{} // 油耗， 有可能是float64或者string，所以用interface{}
}
