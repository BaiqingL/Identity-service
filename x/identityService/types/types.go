package types

import (
	"fmt"
	"strings"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Identity struct {
	ReadableName string         `json:"name"`
	Owner        sdk.AccAddress `json:"owner"`
	Certificate  bool		    `json:"certified"`
}

func NewIdentity() Identity {
	return Whois{
		Certificate: false,
	}
}


func (c Identity) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Identity: %s
Owner: %s
Certified: %s`, c.ReadableName, c.Owner, c.Certificate))
}