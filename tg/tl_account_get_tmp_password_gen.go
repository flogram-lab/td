// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
)

// AccountGetTmpPasswordRequest represents TL type `account.getTmpPassword#449e0b51`.
// Get temporary payment password
//
// See https://core.telegram.org/method/account.getTmpPassword for reference.
type AccountGetTmpPasswordRequest struct {
	// SRP password parameters
	Password InputCheckPasswordSRPClass
	// Time during which the temporary password will be valid, in seconds; should be between
	// 60 and 86400
	Period int
}

// AccountGetTmpPasswordRequestTypeID is TL type id of AccountGetTmpPasswordRequest.
const AccountGetTmpPasswordRequestTypeID = 0x449e0b51

func (g *AccountGetTmpPasswordRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Password == nil) {
		return false
	}
	if !(g.Period == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *AccountGetTmpPasswordRequest) String() string {
	if g == nil {
		return "AccountGetTmpPasswordRequest(nil)"
	}
	type Alias AccountGetTmpPasswordRequest
	return fmt.Sprintf("AccountGetTmpPasswordRequest%+v", Alias(*g))
}

// FillFrom fills AccountGetTmpPasswordRequest from given interface.
func (g *AccountGetTmpPasswordRequest) FillFrom(from interface {
	GetPassword() (value InputCheckPasswordSRPClass)
	GetPeriod() (value int)
}) {
	g.Password = from.GetPassword()
	g.Period = from.GetPeriod()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountGetTmpPasswordRequest) TypeID() uint32 {
	return AccountGetTmpPasswordRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountGetTmpPasswordRequest) TypeName() string {
	return "account.getTmpPassword"
}

// TypeInfo returns info about TL type.
func (g *AccountGetTmpPasswordRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.getTmpPassword",
		ID:   AccountGetTmpPasswordRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Password",
			SchemaName: "password",
		},
		{
			Name:       "Period",
			SchemaName: "period",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *AccountGetTmpPasswordRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getTmpPassword#449e0b51 as nil")
	}
	b.PutID(AccountGetTmpPasswordRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *AccountGetTmpPasswordRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getTmpPassword#449e0b51 as nil")
	}
	if g.Password == nil {
		return fmt.Errorf("unable to encode account.getTmpPassword#449e0b51: field password is nil")
	}
	if err := g.Password.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.getTmpPassword#449e0b51: field password: %w", err)
	}
	b.PutInt(g.Period)
	return nil
}

// GetPassword returns value of Password field.
func (g *AccountGetTmpPasswordRequest) GetPassword() (value InputCheckPasswordSRPClass) {
	return g.Password
}

// GetPasswordAsNotEmpty returns mapped value of Password field.
func (g *AccountGetTmpPasswordRequest) GetPasswordAsNotEmpty() (*InputCheckPasswordSRP, bool) {
	return g.Password.AsNotEmpty()
}

// GetPeriod returns value of Period field.
func (g *AccountGetTmpPasswordRequest) GetPeriod() (value int) {
	return g.Period
}

// Decode implements bin.Decoder.
func (g *AccountGetTmpPasswordRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getTmpPassword#449e0b51 to nil")
	}
	if err := b.ConsumeID(AccountGetTmpPasswordRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getTmpPassword#449e0b51: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *AccountGetTmpPasswordRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getTmpPassword#449e0b51 to nil")
	}
	{
		value, err := DecodeInputCheckPasswordSRP(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.getTmpPassword#449e0b51: field password: %w", err)
		}
		g.Password = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode account.getTmpPassword#449e0b51: field period: %w", err)
		}
		g.Period = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountGetTmpPasswordRequest.
var (
	_ bin.Encoder     = &AccountGetTmpPasswordRequest{}
	_ bin.Decoder     = &AccountGetTmpPasswordRequest{}
	_ bin.BareEncoder = &AccountGetTmpPasswordRequest{}
	_ bin.BareDecoder = &AccountGetTmpPasswordRequest{}
)

// AccountGetTmpPassword invokes method account.getTmpPassword#449e0b51 returning error if any.
// Get temporary payment password
//
// Possible errors:
//  400 PASSWORD_HASH_INVALID: The provided password hash is invalid
//  400 TMP_PASSWORD_DISABLED: The temporary password is disabled
//
// See https://core.telegram.org/method/account.getTmpPassword for reference.
func (c *Client) AccountGetTmpPassword(ctx context.Context, request *AccountGetTmpPasswordRequest) (*AccountTmpPassword, error) {
	var result AccountTmpPassword

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
