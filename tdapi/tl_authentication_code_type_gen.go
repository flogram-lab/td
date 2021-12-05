// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
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
	_ = tdjson.Encoder{}
)

// AuthenticationCodeTypeTelegramMessage represents TL type `authenticationCodeTypeTelegramMessage#7bf49b2a`.
type AuthenticationCodeTypeTelegramMessage struct {
	// Length of the code
	Length int32
}

// AuthenticationCodeTypeTelegramMessageTypeID is TL type id of AuthenticationCodeTypeTelegramMessage.
const AuthenticationCodeTypeTelegramMessageTypeID = 0x7bf49b2a

// construct implements constructor of AuthenticationCodeTypeClass.
func (a AuthenticationCodeTypeTelegramMessage) construct() AuthenticationCodeTypeClass { return &a }

// Ensuring interfaces in compile-time for AuthenticationCodeTypeTelegramMessage.
var (
	_ bin.Encoder     = &AuthenticationCodeTypeTelegramMessage{}
	_ bin.Decoder     = &AuthenticationCodeTypeTelegramMessage{}
	_ bin.BareEncoder = &AuthenticationCodeTypeTelegramMessage{}
	_ bin.BareDecoder = &AuthenticationCodeTypeTelegramMessage{}

	_ AuthenticationCodeTypeClass = &AuthenticationCodeTypeTelegramMessage{}
)

func (a *AuthenticationCodeTypeTelegramMessage) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Length == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AuthenticationCodeTypeTelegramMessage) String() string {
	if a == nil {
		return "AuthenticationCodeTypeTelegramMessage(nil)"
	}
	type Alias AuthenticationCodeTypeTelegramMessage
	return fmt.Sprintf("AuthenticationCodeTypeTelegramMessage%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthenticationCodeTypeTelegramMessage) TypeID() uint32 {
	return AuthenticationCodeTypeTelegramMessageTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthenticationCodeTypeTelegramMessage) TypeName() string {
	return "authenticationCodeTypeTelegramMessage"
}

// TypeInfo returns info about TL type.
func (a *AuthenticationCodeTypeTelegramMessage) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "authenticationCodeTypeTelegramMessage",
		ID:   AuthenticationCodeTypeTelegramMessageTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Length",
			SchemaName: "length",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AuthenticationCodeTypeTelegramMessage) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode authenticationCodeTypeTelegramMessage#7bf49b2a as nil")
	}
	b.PutID(AuthenticationCodeTypeTelegramMessageTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AuthenticationCodeTypeTelegramMessage) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode authenticationCodeTypeTelegramMessage#7bf49b2a as nil")
	}
	b.PutInt32(a.Length)
	return nil
}

// Decode implements bin.Decoder.
func (a *AuthenticationCodeTypeTelegramMessage) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode authenticationCodeTypeTelegramMessage#7bf49b2a to nil")
	}
	if err := b.ConsumeID(AuthenticationCodeTypeTelegramMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode authenticationCodeTypeTelegramMessage#7bf49b2a: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AuthenticationCodeTypeTelegramMessage) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode authenticationCodeTypeTelegramMessage#7bf49b2a to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode authenticationCodeTypeTelegramMessage#7bf49b2a: field length: %w", err)
		}
		a.Length = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (a *AuthenticationCodeTypeTelegramMessage) EncodeTDLibJSON(b tdjson.Encoder) error {
	if a == nil {
		return fmt.Errorf("can't encode authenticationCodeTypeTelegramMessage#7bf49b2a as nil")
	}
	b.ObjStart()
	b.PutID("authenticationCodeTypeTelegramMessage")
	b.FieldStart("length")
	b.PutInt32(a.Length)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (a *AuthenticationCodeTypeTelegramMessage) DecodeTDLibJSON(b tdjson.Decoder) error {
	if a == nil {
		return fmt.Errorf("can't decode authenticationCodeTypeTelegramMessage#7bf49b2a to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("authenticationCodeTypeTelegramMessage"); err != nil {
				return fmt.Errorf("unable to decode authenticationCodeTypeTelegramMessage#7bf49b2a: %w", err)
			}
		case "length":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode authenticationCodeTypeTelegramMessage#7bf49b2a: field length: %w", err)
			}
			a.Length = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetLength returns value of Length field.
func (a *AuthenticationCodeTypeTelegramMessage) GetLength() (value int32) {
	return a.Length
}

// AuthenticationCodeTypeSMS represents TL type `authenticationCodeTypeSms#3960e288`.
type AuthenticationCodeTypeSMS struct {
	// Length of the code
	Length int32
}

// AuthenticationCodeTypeSMSTypeID is TL type id of AuthenticationCodeTypeSMS.
const AuthenticationCodeTypeSMSTypeID = 0x3960e288

// construct implements constructor of AuthenticationCodeTypeClass.
func (a AuthenticationCodeTypeSMS) construct() AuthenticationCodeTypeClass { return &a }

// Ensuring interfaces in compile-time for AuthenticationCodeTypeSMS.
var (
	_ bin.Encoder     = &AuthenticationCodeTypeSMS{}
	_ bin.Decoder     = &AuthenticationCodeTypeSMS{}
	_ bin.BareEncoder = &AuthenticationCodeTypeSMS{}
	_ bin.BareDecoder = &AuthenticationCodeTypeSMS{}

	_ AuthenticationCodeTypeClass = &AuthenticationCodeTypeSMS{}
)

func (a *AuthenticationCodeTypeSMS) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Length == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AuthenticationCodeTypeSMS) String() string {
	if a == nil {
		return "AuthenticationCodeTypeSMS(nil)"
	}
	type Alias AuthenticationCodeTypeSMS
	return fmt.Sprintf("AuthenticationCodeTypeSMS%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthenticationCodeTypeSMS) TypeID() uint32 {
	return AuthenticationCodeTypeSMSTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthenticationCodeTypeSMS) TypeName() string {
	return "authenticationCodeTypeSms"
}

// TypeInfo returns info about TL type.
func (a *AuthenticationCodeTypeSMS) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "authenticationCodeTypeSms",
		ID:   AuthenticationCodeTypeSMSTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Length",
			SchemaName: "length",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AuthenticationCodeTypeSMS) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode authenticationCodeTypeSms#3960e288 as nil")
	}
	b.PutID(AuthenticationCodeTypeSMSTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AuthenticationCodeTypeSMS) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode authenticationCodeTypeSms#3960e288 as nil")
	}
	b.PutInt32(a.Length)
	return nil
}

// Decode implements bin.Decoder.
func (a *AuthenticationCodeTypeSMS) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode authenticationCodeTypeSms#3960e288 to nil")
	}
	if err := b.ConsumeID(AuthenticationCodeTypeSMSTypeID); err != nil {
		return fmt.Errorf("unable to decode authenticationCodeTypeSms#3960e288: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AuthenticationCodeTypeSMS) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode authenticationCodeTypeSms#3960e288 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode authenticationCodeTypeSms#3960e288: field length: %w", err)
		}
		a.Length = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (a *AuthenticationCodeTypeSMS) EncodeTDLibJSON(b tdjson.Encoder) error {
	if a == nil {
		return fmt.Errorf("can't encode authenticationCodeTypeSms#3960e288 as nil")
	}
	b.ObjStart()
	b.PutID("authenticationCodeTypeSms")
	b.FieldStart("length")
	b.PutInt32(a.Length)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (a *AuthenticationCodeTypeSMS) DecodeTDLibJSON(b tdjson.Decoder) error {
	if a == nil {
		return fmt.Errorf("can't decode authenticationCodeTypeSms#3960e288 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("authenticationCodeTypeSms"); err != nil {
				return fmt.Errorf("unable to decode authenticationCodeTypeSms#3960e288: %w", err)
			}
		case "length":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode authenticationCodeTypeSms#3960e288: field length: %w", err)
			}
			a.Length = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetLength returns value of Length field.
func (a *AuthenticationCodeTypeSMS) GetLength() (value int32) {
	return a.Length
}

// AuthenticationCodeTypeCall represents TL type `authenticationCodeTypeCall#61876c67`.
type AuthenticationCodeTypeCall struct {
	// Length of the code
	Length int32
}

// AuthenticationCodeTypeCallTypeID is TL type id of AuthenticationCodeTypeCall.
const AuthenticationCodeTypeCallTypeID = 0x61876c67

// construct implements constructor of AuthenticationCodeTypeClass.
func (a AuthenticationCodeTypeCall) construct() AuthenticationCodeTypeClass { return &a }

// Ensuring interfaces in compile-time for AuthenticationCodeTypeCall.
var (
	_ bin.Encoder     = &AuthenticationCodeTypeCall{}
	_ bin.Decoder     = &AuthenticationCodeTypeCall{}
	_ bin.BareEncoder = &AuthenticationCodeTypeCall{}
	_ bin.BareDecoder = &AuthenticationCodeTypeCall{}

	_ AuthenticationCodeTypeClass = &AuthenticationCodeTypeCall{}
)

func (a *AuthenticationCodeTypeCall) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Length == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AuthenticationCodeTypeCall) String() string {
	if a == nil {
		return "AuthenticationCodeTypeCall(nil)"
	}
	type Alias AuthenticationCodeTypeCall
	return fmt.Sprintf("AuthenticationCodeTypeCall%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthenticationCodeTypeCall) TypeID() uint32 {
	return AuthenticationCodeTypeCallTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthenticationCodeTypeCall) TypeName() string {
	return "authenticationCodeTypeCall"
}

// TypeInfo returns info about TL type.
func (a *AuthenticationCodeTypeCall) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "authenticationCodeTypeCall",
		ID:   AuthenticationCodeTypeCallTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Length",
			SchemaName: "length",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AuthenticationCodeTypeCall) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode authenticationCodeTypeCall#61876c67 as nil")
	}
	b.PutID(AuthenticationCodeTypeCallTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AuthenticationCodeTypeCall) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode authenticationCodeTypeCall#61876c67 as nil")
	}
	b.PutInt32(a.Length)
	return nil
}

// Decode implements bin.Decoder.
func (a *AuthenticationCodeTypeCall) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode authenticationCodeTypeCall#61876c67 to nil")
	}
	if err := b.ConsumeID(AuthenticationCodeTypeCallTypeID); err != nil {
		return fmt.Errorf("unable to decode authenticationCodeTypeCall#61876c67: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AuthenticationCodeTypeCall) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode authenticationCodeTypeCall#61876c67 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode authenticationCodeTypeCall#61876c67: field length: %w", err)
		}
		a.Length = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (a *AuthenticationCodeTypeCall) EncodeTDLibJSON(b tdjson.Encoder) error {
	if a == nil {
		return fmt.Errorf("can't encode authenticationCodeTypeCall#61876c67 as nil")
	}
	b.ObjStart()
	b.PutID("authenticationCodeTypeCall")
	b.FieldStart("length")
	b.PutInt32(a.Length)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (a *AuthenticationCodeTypeCall) DecodeTDLibJSON(b tdjson.Decoder) error {
	if a == nil {
		return fmt.Errorf("can't decode authenticationCodeTypeCall#61876c67 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("authenticationCodeTypeCall"); err != nil {
				return fmt.Errorf("unable to decode authenticationCodeTypeCall#61876c67: %w", err)
			}
		case "length":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode authenticationCodeTypeCall#61876c67: field length: %w", err)
			}
			a.Length = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetLength returns value of Length field.
func (a *AuthenticationCodeTypeCall) GetLength() (value int32) {
	return a.Length
}

// AuthenticationCodeTypeFlashCall represents TL type `authenticationCodeTypeFlashCall#533379a2`.
type AuthenticationCodeTypeFlashCall struct {
	// Pattern of the phone number from which the call will be made
	Pattern string
}

// AuthenticationCodeTypeFlashCallTypeID is TL type id of AuthenticationCodeTypeFlashCall.
const AuthenticationCodeTypeFlashCallTypeID = 0x533379a2

// construct implements constructor of AuthenticationCodeTypeClass.
func (a AuthenticationCodeTypeFlashCall) construct() AuthenticationCodeTypeClass { return &a }

// Ensuring interfaces in compile-time for AuthenticationCodeTypeFlashCall.
var (
	_ bin.Encoder     = &AuthenticationCodeTypeFlashCall{}
	_ bin.Decoder     = &AuthenticationCodeTypeFlashCall{}
	_ bin.BareEncoder = &AuthenticationCodeTypeFlashCall{}
	_ bin.BareDecoder = &AuthenticationCodeTypeFlashCall{}

	_ AuthenticationCodeTypeClass = &AuthenticationCodeTypeFlashCall{}
)

func (a *AuthenticationCodeTypeFlashCall) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Pattern == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AuthenticationCodeTypeFlashCall) String() string {
	if a == nil {
		return "AuthenticationCodeTypeFlashCall(nil)"
	}
	type Alias AuthenticationCodeTypeFlashCall
	return fmt.Sprintf("AuthenticationCodeTypeFlashCall%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthenticationCodeTypeFlashCall) TypeID() uint32 {
	return AuthenticationCodeTypeFlashCallTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthenticationCodeTypeFlashCall) TypeName() string {
	return "authenticationCodeTypeFlashCall"
}

// TypeInfo returns info about TL type.
func (a *AuthenticationCodeTypeFlashCall) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "authenticationCodeTypeFlashCall",
		ID:   AuthenticationCodeTypeFlashCallTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Pattern",
			SchemaName: "pattern",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AuthenticationCodeTypeFlashCall) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode authenticationCodeTypeFlashCall#533379a2 as nil")
	}
	b.PutID(AuthenticationCodeTypeFlashCallTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AuthenticationCodeTypeFlashCall) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode authenticationCodeTypeFlashCall#533379a2 as nil")
	}
	b.PutString(a.Pattern)
	return nil
}

// Decode implements bin.Decoder.
func (a *AuthenticationCodeTypeFlashCall) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode authenticationCodeTypeFlashCall#533379a2 to nil")
	}
	if err := b.ConsumeID(AuthenticationCodeTypeFlashCallTypeID); err != nil {
		return fmt.Errorf("unable to decode authenticationCodeTypeFlashCall#533379a2: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AuthenticationCodeTypeFlashCall) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode authenticationCodeTypeFlashCall#533379a2 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode authenticationCodeTypeFlashCall#533379a2: field pattern: %w", err)
		}
		a.Pattern = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (a *AuthenticationCodeTypeFlashCall) EncodeTDLibJSON(b tdjson.Encoder) error {
	if a == nil {
		return fmt.Errorf("can't encode authenticationCodeTypeFlashCall#533379a2 as nil")
	}
	b.ObjStart()
	b.PutID("authenticationCodeTypeFlashCall")
	b.FieldStart("pattern")
	b.PutString(a.Pattern)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (a *AuthenticationCodeTypeFlashCall) DecodeTDLibJSON(b tdjson.Decoder) error {
	if a == nil {
		return fmt.Errorf("can't decode authenticationCodeTypeFlashCall#533379a2 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("authenticationCodeTypeFlashCall"); err != nil {
				return fmt.Errorf("unable to decode authenticationCodeTypeFlashCall#533379a2: %w", err)
			}
		case "pattern":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode authenticationCodeTypeFlashCall#533379a2: field pattern: %w", err)
			}
			a.Pattern = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetPattern returns value of Pattern field.
func (a *AuthenticationCodeTypeFlashCall) GetPattern() (value string) {
	return a.Pattern
}

// AuthenticationCodeTypeClassName is schema name of AuthenticationCodeTypeClass.
const AuthenticationCodeTypeClassName = "AuthenticationCodeType"

// AuthenticationCodeTypeClass represents AuthenticationCodeType generic type.
//
// Example:
//  g, err := tdapi.DecodeAuthenticationCodeType(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tdapi.AuthenticationCodeTypeTelegramMessage: // authenticationCodeTypeTelegramMessage#7bf49b2a
//  case *tdapi.AuthenticationCodeTypeSMS: // authenticationCodeTypeSms#3960e288
//  case *tdapi.AuthenticationCodeTypeCall: // authenticationCodeTypeCall#61876c67
//  case *tdapi.AuthenticationCodeTypeFlashCall: // authenticationCodeTypeFlashCall#533379a2
//  default: panic(v)
//  }
type AuthenticationCodeTypeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() AuthenticationCodeTypeClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	EncodeTDLibJSON(b tdjson.Encoder) error
	DecodeTDLibJSON(b tdjson.Decoder) error
}

// DecodeAuthenticationCodeType implements binary de-serialization for AuthenticationCodeTypeClass.
func DecodeAuthenticationCodeType(buf *bin.Buffer) (AuthenticationCodeTypeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case AuthenticationCodeTypeTelegramMessageTypeID:
		// Decoding authenticationCodeTypeTelegramMessage#7bf49b2a.
		v := AuthenticationCodeTypeTelegramMessage{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthenticationCodeTypeClass: %w", err)
		}
		return &v, nil
	case AuthenticationCodeTypeSMSTypeID:
		// Decoding authenticationCodeTypeSms#3960e288.
		v := AuthenticationCodeTypeSMS{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthenticationCodeTypeClass: %w", err)
		}
		return &v, nil
	case AuthenticationCodeTypeCallTypeID:
		// Decoding authenticationCodeTypeCall#61876c67.
		v := AuthenticationCodeTypeCall{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthenticationCodeTypeClass: %w", err)
		}
		return &v, nil
	case AuthenticationCodeTypeFlashCallTypeID:
		// Decoding authenticationCodeTypeFlashCall#533379a2.
		v := AuthenticationCodeTypeFlashCall{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthenticationCodeTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode AuthenticationCodeTypeClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONAuthenticationCodeType implements binary de-serialization for AuthenticationCodeTypeClass.
func DecodeTDLibJSONAuthenticationCodeType(buf tdjson.Decoder) (AuthenticationCodeTypeClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "authenticationCodeTypeTelegramMessage":
		// Decoding authenticationCodeTypeTelegramMessage#7bf49b2a.
		v := AuthenticationCodeTypeTelegramMessage{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthenticationCodeTypeClass: %w", err)
		}
		return &v, nil
	case "authenticationCodeTypeSms":
		// Decoding authenticationCodeTypeSms#3960e288.
		v := AuthenticationCodeTypeSMS{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthenticationCodeTypeClass: %w", err)
		}
		return &v, nil
	case "authenticationCodeTypeCall":
		// Decoding authenticationCodeTypeCall#61876c67.
		v := AuthenticationCodeTypeCall{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthenticationCodeTypeClass: %w", err)
		}
		return &v, nil
	case "authenticationCodeTypeFlashCall":
		// Decoding authenticationCodeTypeFlashCall#533379a2.
		v := AuthenticationCodeTypeFlashCall{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthenticationCodeTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode AuthenticationCodeTypeClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// AuthenticationCodeType boxes the AuthenticationCodeTypeClass providing a helper.
type AuthenticationCodeTypeBox struct {
	AuthenticationCodeType AuthenticationCodeTypeClass
}

// Decode implements bin.Decoder for AuthenticationCodeTypeBox.
func (b *AuthenticationCodeTypeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode AuthenticationCodeTypeBox to nil")
	}
	v, err := DecodeAuthenticationCodeType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.AuthenticationCodeType = v
	return nil
}

// Encode implements bin.Encode for AuthenticationCodeTypeBox.
func (b *AuthenticationCodeTypeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.AuthenticationCodeType == nil {
		return fmt.Errorf("unable to encode AuthenticationCodeTypeClass as nil")
	}
	return b.AuthenticationCodeType.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for AuthenticationCodeTypeBox.
func (b *AuthenticationCodeTypeBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode AuthenticationCodeTypeBox to nil")
	}
	v, err := DecodeTDLibJSONAuthenticationCodeType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.AuthenticationCodeType = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for AuthenticationCodeTypeBox.
func (b *AuthenticationCodeTypeBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.AuthenticationCodeType == nil {
		return fmt.Errorf("unable to encode AuthenticationCodeTypeClass as nil")
	}
	return b.AuthenticationCodeType.EncodeTDLibJSON(buf)
}
