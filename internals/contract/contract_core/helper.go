package contract_core

import ("github.com/goethercore/goether/types")


func ToCallArg(msg types.ParamObject) interface{} {
arg := types.ParamObject{
	To:  msg.To,
	From: msg.From,
	Data: msg.Data,
}
return  []interface{}{arg, "latest"}
}
