package lvmsdk

type DoRequest struct {
	Fn      string `json:"fn"`
	Content string `json:"content"`
}

type CallRequest struct {
	Fn       string                 `json:"fn"`
	Content  string                 `json:"content"`
	Function string                 `json:"function"`
	Params   map[string]interface{} `json:"params"`
}
