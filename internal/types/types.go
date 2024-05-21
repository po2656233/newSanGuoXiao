package types

import (
	mapStructure "github.com/po2656233/superplace/extend/mapstructure"
	"reflect"
)

type (
	HookType interface {
		Type() reflect.Type
		Hook() mapStructure.DecodeHookFuncType
	}
)

var (
	funcTypes []mapStructure.DecodeHookFuncType
)

func init() {
	// 需要通过json解析数据的类型，注册到此
	register(&I32I32{})
	register(&I32I64Map{})
}

func register(t HookType) {
	funcTypes = append(funcTypes, t.Hook())
}

func GetDecodeHooks() []mapStructure.DecodeHookFuncType {
	return funcTypes
}
