package identityservice

import (
	"ICT/x/nameservice/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewMsgSetName = types.NewMsgSetName
	NewIdentity   = types.NewIdentity
	ModuleCdc     = types.ModuleCdc
	RegisterCodec = types.RegisterCodec
)

type (
	MsgSetName      = types.MsgSetName
	QueryResResolve = types.QueryResResolve
	QueryResNames   = types.QueryResNames
	Identity           = types.Identity
)