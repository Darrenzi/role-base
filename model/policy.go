package model

type Policy struct {
	Sub    string `form:"sub"`
	Obj    string `form:"obj"`
	Method string `form:"method"`
}
