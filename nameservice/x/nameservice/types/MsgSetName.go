package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgSetName defines a SetName message
var _ sdk.Msg = &MsgSetName{}

type MsgSetName struct {
	Owner sdk.AccAddress `json:"owner" yaml:"owner"`
	Value string         `json:"value" yaml:"value"`
	Name  string         `json:"name" yaml:"name"`
}

// NewMsgSetName is a constructor function for MsgSetName
func NewMsgSetName(owner sdk.AccAddress, name string, value string) MsgSetName {
	return MsgSetName{
		Owner: owner,
		Value: value,
		Name:  name,
	}
}

func (msg MsgSetName) Route() string {
	return RouterKey
}

func (msg MsgSetName) Type() string {
	return "SetName"
}

func (msg MsgSetName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

func (msg MsgSetName) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic runs stateless checks on the message
func (msg MsgSetName) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Owner.String())
	}
	if len(msg.Name) == 0 || len(msg.Value) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Name and/or Value cannot be empty")
	}
	return nil
}
