package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName // this was defined in your key.go file

// MsgSetName defines a SetName message
type MsgSetName struct {
	ReadableName string         `json:"name"`
	Owner        sdk.AccAddress `json:"owner"`
	Certificate  bool		    `json:"certified"`
}

func NewMsgSetName(name string, owner sdk.AccAddress, certified bool) MsgSetName {
	return MsgSetName{
		ReadableName: name,
		Owner: owner,
		Certificate: certified,
	}
}

// Route should return the name of the module
func (msg MsgSetName) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSetName) Type() string { return "set_name" }

func (msg MsgSetName) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.ReadableName) == 0 || !msg.Certificate {
		return sdk.ErrUnknownRequest("Account must be certified or contain a name")
	}
	return nil
}

func (msg MsgSetName) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgSetName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

type MsgRegister struct {
	Name     string         `json:"name"`
	Register sdk.AccAddress `json:"account"`
}

func (msg MsgRegister) Route() string { return RouterKey }

func (msg MsgRegister) Type() string { return "register_name" }

func (msg MsgRegister) ValidateBasic() sdk.Error {
	if msg.Register.Empty() {
		return sdk.ErrInvalidAddress(msg.Register.String())
	}
	if len(msg.Name) == 0 {
		return sdk.ErrUnknownRequest("Name cannot be empty")
	}

	return nil
}

func (msg MsgRegister) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgRegister) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Buyer}
}

