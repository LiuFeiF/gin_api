package res

type Response struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

type Responses struct { //这种是返回的内层数据列表，以及有多少个数据。curd常用
	ItemList interface{} `json:"itemList"`
	Total    int         `json:"total"`
}
