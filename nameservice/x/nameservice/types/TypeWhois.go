package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var MinNamePrice = sdk.Coins{sdk.NewInt64Coin("nametoken", 1)}

type Whois struct {
	Creator sdk.AccAddress `json:"creator" yaml:"owner"`
	ID      string         `json:"id" yaml:"id"`
	Value   string         `json:"value" yaml:"value"`
	Price   sdk.Coins      `json:"price" yaml:"price"`
}

func NewWhois() Whois {
	return Whois{
		Price: MinNamePrice,
	}
}

func (w Whois) string() string {
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
	Value: %s
	Price: %s`, w.Creator, w.Value, w.Price))
}
