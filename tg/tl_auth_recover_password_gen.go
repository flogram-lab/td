// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdp"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}
var _ = errors.Is
var _ = sort.Ints
var _ = tdp.Format

// AuthRecoverPasswordRequest represents TL type `auth.recoverPassword#4ea56e92`.
// Reset the 2FA password¹ using the recovery code sent using auth.requestPasswordRecovery².
//
// Links:
//  1) https://core.telegram.org/api/srp
//  2) https://core.telegram.org/method/auth.requestPasswordRecovery
//
// See https://core.telegram.org/method/auth.recoverPassword for reference.
type AuthRecoverPasswordRequest struct {
	// Code received via email
	Code string
}

// AuthRecoverPasswordRequestTypeID is TL type id of AuthRecoverPasswordRequest.
const AuthRecoverPasswordRequestTypeID = 0x4ea56e92

func (r *AuthRecoverPasswordRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Code == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *AuthRecoverPasswordRequest) String() string {
	if r == nil {
		return "AuthRecoverPasswordRequest(nil)"
	}
	type Alias AuthRecoverPasswordRequest
	return fmt.Sprintf("AuthRecoverPasswordRequest%+v", Alias(*r))
}

// FillFrom fills AuthRecoverPasswordRequest from given interface.
func (r *AuthRecoverPasswordRequest) FillFrom(from interface {
	GetCode() (value string)
}) {
	r.Code = from.GetCode()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthRecoverPasswordRequest) TypeID() uint32 {
	return AuthRecoverPasswordRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthRecoverPasswordRequest) TypeName() string {
	return "auth.recoverPassword"
}

// TypeInfo returns info about TL type.
func (r *AuthRecoverPasswordRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.recoverPassword",
		ID:   AuthRecoverPasswordRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Code",
			SchemaName: "code",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *AuthRecoverPasswordRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode auth.recoverPassword#4ea56e92 as nil")
	}
	b.PutID(AuthRecoverPasswordRequestTypeID)
	b.PutString(r.Code)
	return nil
}

// GetCode returns value of Code field.
func (r *AuthRecoverPasswordRequest) GetCode() (value string) {
	return r.Code
}

// Decode implements bin.Decoder.
func (r *AuthRecoverPasswordRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode auth.recoverPassword#4ea56e92 to nil")
	}
	if err := b.ConsumeID(AuthRecoverPasswordRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.recoverPassword#4ea56e92: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.recoverPassword#4ea56e92: field code: %w", err)
		}
		r.Code = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AuthRecoverPasswordRequest.
var (
	_ bin.Encoder = &AuthRecoverPasswordRequest{}
	_ bin.Decoder = &AuthRecoverPasswordRequest{}
)

// AuthRecoverPassword invokes method auth.recoverPassword#4ea56e92 returning error if any.
// Reset the 2FA password¹ using the recovery code sent using auth.requestPasswordRecovery².
//
// Links:
//  1) https://core.telegram.org/api/srp
//  2) https://core.telegram.org/method/auth.requestPasswordRecovery
//
// Possible errors:
//  400 CODE_EMPTY: The provided code is empty
//
// See https://core.telegram.org/method/auth.recoverPassword for reference.
func (c *Client) AuthRecoverPassword(ctx context.Context, code string) (AuthAuthorizationClass, error) {
	var result AuthAuthorizationBox

	request := &AuthRecoverPasswordRequest{
		Code: code,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Authorization, nil
}
