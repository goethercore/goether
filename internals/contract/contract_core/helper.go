package contract_core

import ("github.com/ayoseun/geth-lite/types")


func ToCallArg(msg types.ParamObject) interface{} {
arg := types.ParamObject{
	To:  msg.To,
	From: msg.From,
	Data: msg.Data,
}
return  []interface{}{arg, "latest"}
}
