package lvmsdk

type Response struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Reason string      `json:"reason"`
	Data   interface{} `json:"data"`
}
