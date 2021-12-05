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

// LoginURLInfoOpen represents TL type `loginUrlInfoOpen#bfaf12d4`.
type LoginURLInfoOpen struct {
	// The URL to open
	URL string
	// True, if there is no need to show an ordinary open URL confirm
	SkipConfirm bool
}

// LoginURLInfoOpenTypeID is TL type id of LoginURLInfoOpen.
const LoginURLInfoOpenTypeID = 0xbfaf12d4

// construct implements constructor of LoginURLInfoClass.
func (l LoginURLInfoOpen) construct() LoginURLInfoClass { return &l }

// Ensuring interfaces in compile-time for LoginURLInfoOpen.
var (
	_ bin.Encoder     = &LoginURLInfoOpen{}
	_ bin.Decoder     = &LoginURLInfoOpen{}
	_ bin.BareEncoder = &LoginURLInfoOpen{}
	_ bin.BareDecoder = &LoginURLInfoOpen{}

	_ LoginURLInfoClass = &LoginURLInfoOpen{}
)

func (l *LoginURLInfoOpen) Zero() bool {
	if l == nil {
		return true
	}
	if !(l.URL == "") {
		return false
	}
	if !(l.SkipConfirm == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (l *LoginURLInfoOpen) String() string {
	if l == nil {
		return "LoginURLInfoOpen(nil)"
	}
	type Alias LoginURLInfoOpen
	return fmt.Sprintf("LoginURLInfoOpen%+v", Alias(*l))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*LoginURLInfoOpen) TypeID() uint32 {
	return LoginURLInfoOpenTypeID
}

// TypeName returns name of type in TL schema.
func (*LoginURLInfoOpen) TypeName() string {
	return "loginUrlInfoOpen"
}

// TypeInfo returns info about TL type.
func (l *LoginURLInfoOpen) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "loginUrlInfoOpen",
		ID:   LoginURLInfoOpenTypeID,
	}
	if l == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
		{
			Name:       "SkipConfirm",
			SchemaName: "skip_confirm",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (l *LoginURLInfoOpen) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode loginUrlInfoOpen#bfaf12d4 as nil")
	}
	b.PutID(LoginURLInfoOpenTypeID)
	return l.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (l *LoginURLInfoOpen) EncodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode loginUrlInfoOpen#bfaf12d4 as nil")
	}
	b.PutString(l.URL)
	b.PutBool(l.SkipConfirm)
	return nil
}

// Decode implements bin.Decoder.
func (l *LoginURLInfoOpen) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode loginUrlInfoOpen#bfaf12d4 to nil")
	}
	if err := b.ConsumeID(LoginURLInfoOpenTypeID); err != nil {
		return fmt.Errorf("unable to decode loginUrlInfoOpen#bfaf12d4: %w", err)
	}
	return l.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (l *LoginURLInfoOpen) DecodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode loginUrlInfoOpen#bfaf12d4 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode loginUrlInfoOpen#bfaf12d4: field url: %w", err)
		}
		l.URL = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode loginUrlInfoOpen#bfaf12d4: field skip_confirm: %w", err)
		}
		l.SkipConfirm = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (l *LoginURLInfoOpen) EncodeTDLibJSON(b tdjson.Encoder) error {
	if l == nil {
		return fmt.Errorf("can't encode loginUrlInfoOpen#bfaf12d4 as nil")
	}
	b.ObjStart()
	b.PutID("loginUrlInfoOpen")
	b.FieldStart("url")
	b.PutString(l.URL)
	b.FieldStart("skip_confirm")
	b.PutBool(l.SkipConfirm)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (l *LoginURLInfoOpen) DecodeTDLibJSON(b tdjson.Decoder) error {
	if l == nil {
		return fmt.Errorf("can't decode loginUrlInfoOpen#bfaf12d4 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("loginUrlInfoOpen"); err != nil {
				return fmt.Errorf("unable to decode loginUrlInfoOpen#bfaf12d4: %w", err)
			}
		case "url":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode loginUrlInfoOpen#bfaf12d4: field url: %w", err)
			}
			l.URL = value
		case "skip_confirm":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode loginUrlInfoOpen#bfaf12d4: field skip_confirm: %w", err)
			}
			l.SkipConfirm = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetURL returns value of URL field.
func (l *LoginURLInfoOpen) GetURL() (value string) {
	return l.URL
}

// GetSkipConfirm returns value of SkipConfirm field.
func (l *LoginURLInfoOpen) GetSkipConfirm() (value bool) {
	return l.SkipConfirm
}

// LoginURLInfoRequestConfirmation represents TL type `loginUrlInfoRequestConfirmation#7edb242f`.
type LoginURLInfoRequestConfirmation struct {
	// An HTTP URL to be opened
	URL string
	// A domain of the URL
	Domain string
	// User identifier of a bot linked with the website
	BotUserID int64
	// True, if the user needs to be requested to give the permission to the bot to send them
	// messages
	RequestWriteAccess bool
}

// LoginURLInfoRequestConfirmationTypeID is TL type id of LoginURLInfoRequestConfirmation.
const LoginURLInfoRequestConfirmationTypeID = 0x7edb242f

// construct implements constructor of LoginURLInfoClass.
func (l LoginURLInfoRequestConfirmation) construct() LoginURLInfoClass { return &l }

// Ensuring interfaces in compile-time for LoginURLInfoRequestConfirmation.
var (
	_ bin.Encoder     = &LoginURLInfoRequestConfirmation{}
	_ bin.Decoder     = &LoginURLInfoRequestConfirmation{}
	_ bin.BareEncoder = &LoginURLInfoRequestConfirmation{}
	_ bin.BareDecoder = &LoginURLInfoRequestConfirmation{}

	_ LoginURLInfoClass = &LoginURLInfoRequestConfirmation{}
)

func (l *LoginURLInfoRequestConfirmation) Zero() bool {
	if l == nil {
		return true
	}
	if !(l.URL == "") {
		return false
	}
	if !(l.Domain == "") {
		return false
	}
	if !(l.BotUserID == 0) {
		return false
	}
	if !(l.RequestWriteAccess == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (l *LoginURLInfoRequestConfirmation) String() string {
	if l == nil {
		return "LoginURLInfoRequestConfirmation(nil)"
	}
	type Alias LoginURLInfoRequestConfirmation
	return fmt.Sprintf("LoginURLInfoRequestConfirmation%+v", Alias(*l))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*LoginURLInfoRequestConfirmation) TypeID() uint32 {
	return LoginURLInfoRequestConfirmationTypeID
}

// TypeName returns name of type in TL schema.
func (*LoginURLInfoRequestConfirmation) TypeName() string {
	return "loginUrlInfoRequestConfirmation"
}

// TypeInfo returns info about TL type.
func (l *LoginURLInfoRequestConfirmation) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "loginUrlInfoRequestConfirmation",
		ID:   LoginURLInfoRequestConfirmationTypeID,
	}
	if l == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
		{
			Name:       "Domain",
			SchemaName: "domain",
		},
		{
			Name:       "BotUserID",
			SchemaName: "bot_user_id",
		},
		{
			Name:       "RequestWriteAccess",
			SchemaName: "request_write_access",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (l *LoginURLInfoRequestConfirmation) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode loginUrlInfoRequestConfirmation#7edb242f as nil")
	}
	b.PutID(LoginURLInfoRequestConfirmationTypeID)
	return l.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (l *LoginURLInfoRequestConfirmation) EncodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode loginUrlInfoRequestConfirmation#7edb242f as nil")
	}
	b.PutString(l.URL)
	b.PutString(l.Domain)
	b.PutLong(l.BotUserID)
	b.PutBool(l.RequestWriteAccess)
	return nil
}

// Decode implements bin.Decoder.
func (l *LoginURLInfoRequestConfirmation) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode loginUrlInfoRequestConfirmation#7edb242f to nil")
	}
	if err := b.ConsumeID(LoginURLInfoRequestConfirmationTypeID); err != nil {
		return fmt.Errorf("unable to decode loginUrlInfoRequestConfirmation#7edb242f: %w", err)
	}
	return l.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (l *LoginURLInfoRequestConfirmation) DecodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode loginUrlInfoRequestConfirmation#7edb242f to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode loginUrlInfoRequestConfirmation#7edb242f: field url: %w", err)
		}
		l.URL = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode loginUrlInfoRequestConfirmation#7edb242f: field domain: %w", err)
		}
		l.Domain = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode loginUrlInfoRequestConfirmation#7edb242f: field bot_user_id: %w", err)
		}
		l.BotUserID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode loginUrlInfoRequestConfirmation#7edb242f: field request_write_access: %w", err)
		}
		l.RequestWriteAccess = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (l *LoginURLInfoRequestConfirmation) EncodeTDLibJSON(b tdjson.Encoder) error {
	if l == nil {
		return fmt.Errorf("can't encode loginUrlInfoRequestConfirmation#7edb242f as nil")
	}
	b.ObjStart()
	b.PutID("loginUrlInfoRequestConfirmation")
	b.FieldStart("url")
	b.PutString(l.URL)
	b.FieldStart("domain")
	b.PutString(l.Domain)
	b.FieldStart("bot_user_id")
	b.PutLong(l.BotUserID)
	b.FieldStart("request_write_access")
	b.PutBool(l.RequestWriteAccess)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (l *LoginURLInfoRequestConfirmation) DecodeTDLibJSON(b tdjson.Decoder) error {
	if l == nil {
		return fmt.Errorf("can't decode loginUrlInfoRequestConfirmation#7edb242f to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("loginUrlInfoRequestConfirmation"); err != nil {
				return fmt.Errorf("unable to decode loginUrlInfoRequestConfirmation#7edb242f: %w", err)
			}
		case "url":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode loginUrlInfoRequestConfirmation#7edb242f: field url: %w", err)
			}
			l.URL = value
		case "domain":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode loginUrlInfoRequestConfirmation#7edb242f: field domain: %w", err)
			}
			l.Domain = value
		case "bot_user_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode loginUrlInfoRequestConfirmation#7edb242f: field bot_user_id: %w", err)
			}
			l.BotUserID = value
		case "request_write_access":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode loginUrlInfoRequestConfirmation#7edb242f: field request_write_access: %w", err)
			}
			l.RequestWriteAccess = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetURL returns value of URL field.
func (l *LoginURLInfoRequestConfirmation) GetURL() (value string) {
	return l.URL
}

// GetDomain returns value of Domain field.
func (l *LoginURLInfoRequestConfirmation) GetDomain() (value string) {
	return l.Domain
}

// GetBotUserID returns value of BotUserID field.
func (l *LoginURLInfoRequestConfirmation) GetBotUserID() (value int64) {
	return l.BotUserID
}

// GetRequestWriteAccess returns value of RequestWriteAccess field.
func (l *LoginURLInfoRequestConfirmation) GetRequestWriteAccess() (value bool) {
	return l.RequestWriteAccess
}

// LoginURLInfoClassName is schema name of LoginURLInfoClass.
const LoginURLInfoClassName = "LoginUrlInfo"

// LoginURLInfoClass represents LoginUrlInfo generic type.
//
// Example:
//  g, err := tdapi.DecodeLoginURLInfo(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tdapi.LoginURLInfoOpen: // loginUrlInfoOpen#bfaf12d4
//  case *tdapi.LoginURLInfoRequestConfirmation: // loginUrlInfoRequestConfirmation#7edb242f
//  default: panic(v)
//  }
type LoginURLInfoClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() LoginURLInfoClass

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

	// The URL to open
	GetURL() (value string)
}

// DecodeLoginURLInfo implements binary de-serialization for LoginURLInfoClass.
func DecodeLoginURLInfo(buf *bin.Buffer) (LoginURLInfoClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case LoginURLInfoOpenTypeID:
		// Decoding loginUrlInfoOpen#bfaf12d4.
		v := LoginURLInfoOpen{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode LoginURLInfoClass: %w", err)
		}
		return &v, nil
	case LoginURLInfoRequestConfirmationTypeID:
		// Decoding loginUrlInfoRequestConfirmation#7edb242f.
		v := LoginURLInfoRequestConfirmation{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode LoginURLInfoClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode LoginURLInfoClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONLoginURLInfo implements binary de-serialization for LoginURLInfoClass.
func DecodeTDLibJSONLoginURLInfo(buf tdjson.Decoder) (LoginURLInfoClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "loginUrlInfoOpen":
		// Decoding loginUrlInfoOpen#bfaf12d4.
		v := LoginURLInfoOpen{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode LoginURLInfoClass: %w", err)
		}
		return &v, nil
	case "loginUrlInfoRequestConfirmation":
		// Decoding loginUrlInfoRequestConfirmation#7edb242f.
		v := LoginURLInfoRequestConfirmation{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode LoginURLInfoClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode LoginURLInfoClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// LoginURLInfo boxes the LoginURLInfoClass providing a helper.
type LoginURLInfoBox struct {
	LoginUrlInfo LoginURLInfoClass
}

// Decode implements bin.Decoder for LoginURLInfoBox.
func (b *LoginURLInfoBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode LoginURLInfoBox to nil")
	}
	v, err := DecodeLoginURLInfo(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.LoginUrlInfo = v
	return nil
}

// Encode implements bin.Encode for LoginURLInfoBox.
func (b *LoginURLInfoBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.LoginUrlInfo == nil {
		return fmt.Errorf("unable to encode LoginURLInfoClass as nil")
	}
	return b.LoginUrlInfo.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for LoginURLInfoBox.
func (b *LoginURLInfoBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode LoginURLInfoBox to nil")
	}
	v, err := DecodeTDLibJSONLoginURLInfo(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.LoginUrlInfo = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for LoginURLInfoBox.
func (b *LoginURLInfoBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.LoginUrlInfo == nil {
		return fmt.Errorf("unable to encode LoginURLInfoClass as nil")
	}
	return b.LoginUrlInfo.EncodeTDLibJSON(buf)
}
