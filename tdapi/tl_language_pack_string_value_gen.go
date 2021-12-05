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

// LanguagePackStringValueOrdinary represents TL type `languagePackStringValueOrdinary#f124a660`.
type LanguagePackStringValueOrdinary struct {
	// String value
	Value string
}

// LanguagePackStringValueOrdinaryTypeID is TL type id of LanguagePackStringValueOrdinary.
const LanguagePackStringValueOrdinaryTypeID = 0xf124a660

// construct implements constructor of LanguagePackStringValueClass.
func (l LanguagePackStringValueOrdinary) construct() LanguagePackStringValueClass { return &l }

// Ensuring interfaces in compile-time for LanguagePackStringValueOrdinary.
var (
	_ bin.Encoder     = &LanguagePackStringValueOrdinary{}
	_ bin.Decoder     = &LanguagePackStringValueOrdinary{}
	_ bin.BareEncoder = &LanguagePackStringValueOrdinary{}
	_ bin.BareDecoder = &LanguagePackStringValueOrdinary{}

	_ LanguagePackStringValueClass = &LanguagePackStringValueOrdinary{}
)

func (l *LanguagePackStringValueOrdinary) Zero() bool {
	if l == nil {
		return true
	}
	if !(l.Value == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (l *LanguagePackStringValueOrdinary) String() string {
	if l == nil {
		return "LanguagePackStringValueOrdinary(nil)"
	}
	type Alias LanguagePackStringValueOrdinary
	return fmt.Sprintf("LanguagePackStringValueOrdinary%+v", Alias(*l))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*LanguagePackStringValueOrdinary) TypeID() uint32 {
	return LanguagePackStringValueOrdinaryTypeID
}

// TypeName returns name of type in TL schema.
func (*LanguagePackStringValueOrdinary) TypeName() string {
	return "languagePackStringValueOrdinary"
}

// TypeInfo returns info about TL type.
func (l *LanguagePackStringValueOrdinary) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "languagePackStringValueOrdinary",
		ID:   LanguagePackStringValueOrdinaryTypeID,
	}
	if l == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Value",
			SchemaName: "value",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (l *LanguagePackStringValueOrdinary) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode languagePackStringValueOrdinary#f124a660 as nil")
	}
	b.PutID(LanguagePackStringValueOrdinaryTypeID)
	return l.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (l *LanguagePackStringValueOrdinary) EncodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode languagePackStringValueOrdinary#f124a660 as nil")
	}
	b.PutString(l.Value)
	return nil
}

// Decode implements bin.Decoder.
func (l *LanguagePackStringValueOrdinary) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode languagePackStringValueOrdinary#f124a660 to nil")
	}
	if err := b.ConsumeID(LanguagePackStringValueOrdinaryTypeID); err != nil {
		return fmt.Errorf("unable to decode languagePackStringValueOrdinary#f124a660: %w", err)
	}
	return l.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (l *LanguagePackStringValueOrdinary) DecodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode languagePackStringValueOrdinary#f124a660 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode languagePackStringValueOrdinary#f124a660: field value: %w", err)
		}
		l.Value = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (l *LanguagePackStringValueOrdinary) EncodeTDLibJSON(b tdjson.Encoder) error {
	if l == nil {
		return fmt.Errorf("can't encode languagePackStringValueOrdinary#f124a660 as nil")
	}
	b.ObjStart()
	b.PutID("languagePackStringValueOrdinary")
	b.FieldStart("value")
	b.PutString(l.Value)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (l *LanguagePackStringValueOrdinary) DecodeTDLibJSON(b tdjson.Decoder) error {
	if l == nil {
		return fmt.Errorf("can't decode languagePackStringValueOrdinary#f124a660 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("languagePackStringValueOrdinary"); err != nil {
				return fmt.Errorf("unable to decode languagePackStringValueOrdinary#f124a660: %w", err)
			}
		case "value":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode languagePackStringValueOrdinary#f124a660: field value: %w", err)
			}
			l.Value = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetValue returns value of Value field.
func (l *LanguagePackStringValueOrdinary) GetValue() (value string) {
	return l.Value
}

// LanguagePackStringValuePluralized represents TL type `languagePackStringValuePluralized#71a812c5`.
type LanguagePackStringValuePluralized struct {
	// Value for zero objects
	ZeroValue string
	// Value for one object
	OneValue string
	// Value for two objects
	TwoValue string
	// Value for few objects
	FewValue string
	// Value for many objects
	ManyValue string
	// Default value
	OtherValue string
}

// LanguagePackStringValuePluralizedTypeID is TL type id of LanguagePackStringValuePluralized.
const LanguagePackStringValuePluralizedTypeID = 0x71a812c5

// construct implements constructor of LanguagePackStringValueClass.
func (l LanguagePackStringValuePluralized) construct() LanguagePackStringValueClass { return &l }

// Ensuring interfaces in compile-time for LanguagePackStringValuePluralized.
var (
	_ bin.Encoder     = &LanguagePackStringValuePluralized{}
	_ bin.Decoder     = &LanguagePackStringValuePluralized{}
	_ bin.BareEncoder = &LanguagePackStringValuePluralized{}
	_ bin.BareDecoder = &LanguagePackStringValuePluralized{}

	_ LanguagePackStringValueClass = &LanguagePackStringValuePluralized{}
)

func (l *LanguagePackStringValuePluralized) Zero() bool {
	if l == nil {
		return true
	}
	if !(l.ZeroValue == "") {
		return false
	}
	if !(l.OneValue == "") {
		return false
	}
	if !(l.TwoValue == "") {
		return false
	}
	if !(l.FewValue == "") {
		return false
	}
	if !(l.ManyValue == "") {
		return false
	}
	if !(l.OtherValue == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (l *LanguagePackStringValuePluralized) String() string {
	if l == nil {
		return "LanguagePackStringValuePluralized(nil)"
	}
	type Alias LanguagePackStringValuePluralized
	return fmt.Sprintf("LanguagePackStringValuePluralized%+v", Alias(*l))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*LanguagePackStringValuePluralized) TypeID() uint32 {
	return LanguagePackStringValuePluralizedTypeID
}

// TypeName returns name of type in TL schema.
func (*LanguagePackStringValuePluralized) TypeName() string {
	return "languagePackStringValuePluralized"
}

// TypeInfo returns info about TL type.
func (l *LanguagePackStringValuePluralized) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "languagePackStringValuePluralized",
		ID:   LanguagePackStringValuePluralizedTypeID,
	}
	if l == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ZeroValue",
			SchemaName: "zero_value",
		},
		{
			Name:       "OneValue",
			SchemaName: "one_value",
		},
		{
			Name:       "TwoValue",
			SchemaName: "two_value",
		},
		{
			Name:       "FewValue",
			SchemaName: "few_value",
		},
		{
			Name:       "ManyValue",
			SchemaName: "many_value",
		},
		{
			Name:       "OtherValue",
			SchemaName: "other_value",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (l *LanguagePackStringValuePluralized) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode languagePackStringValuePluralized#71a812c5 as nil")
	}
	b.PutID(LanguagePackStringValuePluralizedTypeID)
	return l.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (l *LanguagePackStringValuePluralized) EncodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode languagePackStringValuePluralized#71a812c5 as nil")
	}
	b.PutString(l.ZeroValue)
	b.PutString(l.OneValue)
	b.PutString(l.TwoValue)
	b.PutString(l.FewValue)
	b.PutString(l.ManyValue)
	b.PutString(l.OtherValue)
	return nil
}

// Decode implements bin.Decoder.
func (l *LanguagePackStringValuePluralized) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode languagePackStringValuePluralized#71a812c5 to nil")
	}
	if err := b.ConsumeID(LanguagePackStringValuePluralizedTypeID); err != nil {
		return fmt.Errorf("unable to decode languagePackStringValuePluralized#71a812c5: %w", err)
	}
	return l.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (l *LanguagePackStringValuePluralized) DecodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode languagePackStringValuePluralized#71a812c5 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode languagePackStringValuePluralized#71a812c5: field zero_value: %w", err)
		}
		l.ZeroValue = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode languagePackStringValuePluralized#71a812c5: field one_value: %w", err)
		}
		l.OneValue = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode languagePackStringValuePluralized#71a812c5: field two_value: %w", err)
		}
		l.TwoValue = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode languagePackStringValuePluralized#71a812c5: field few_value: %w", err)
		}
		l.FewValue = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode languagePackStringValuePluralized#71a812c5: field many_value: %w", err)
		}
		l.ManyValue = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode languagePackStringValuePluralized#71a812c5: field other_value: %w", err)
		}
		l.OtherValue = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (l *LanguagePackStringValuePluralized) EncodeTDLibJSON(b tdjson.Encoder) error {
	if l == nil {
		return fmt.Errorf("can't encode languagePackStringValuePluralized#71a812c5 as nil")
	}
	b.ObjStart()
	b.PutID("languagePackStringValuePluralized")
	b.FieldStart("zero_value")
	b.PutString(l.ZeroValue)
	b.FieldStart("one_value")
	b.PutString(l.OneValue)
	b.FieldStart("two_value")
	b.PutString(l.TwoValue)
	b.FieldStart("few_value")
	b.PutString(l.FewValue)
	b.FieldStart("many_value")
	b.PutString(l.ManyValue)
	b.FieldStart("other_value")
	b.PutString(l.OtherValue)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (l *LanguagePackStringValuePluralized) DecodeTDLibJSON(b tdjson.Decoder) error {
	if l == nil {
		return fmt.Errorf("can't decode languagePackStringValuePluralized#71a812c5 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("languagePackStringValuePluralized"); err != nil {
				return fmt.Errorf("unable to decode languagePackStringValuePluralized#71a812c5: %w", err)
			}
		case "zero_value":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode languagePackStringValuePluralized#71a812c5: field zero_value: %w", err)
			}
			l.ZeroValue = value
		case "one_value":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode languagePackStringValuePluralized#71a812c5: field one_value: %w", err)
			}
			l.OneValue = value
		case "two_value":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode languagePackStringValuePluralized#71a812c5: field two_value: %w", err)
			}
			l.TwoValue = value
		case "few_value":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode languagePackStringValuePluralized#71a812c5: field few_value: %w", err)
			}
			l.FewValue = value
		case "many_value":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode languagePackStringValuePluralized#71a812c5: field many_value: %w", err)
			}
			l.ManyValue = value
		case "other_value":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode languagePackStringValuePluralized#71a812c5: field other_value: %w", err)
			}
			l.OtherValue = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetZeroValue returns value of ZeroValue field.
func (l *LanguagePackStringValuePluralized) GetZeroValue() (value string) {
	return l.ZeroValue
}

// GetOneValue returns value of OneValue field.
func (l *LanguagePackStringValuePluralized) GetOneValue() (value string) {
	return l.OneValue
}

// GetTwoValue returns value of TwoValue field.
func (l *LanguagePackStringValuePluralized) GetTwoValue() (value string) {
	return l.TwoValue
}

// GetFewValue returns value of FewValue field.
func (l *LanguagePackStringValuePluralized) GetFewValue() (value string) {
	return l.FewValue
}

// GetManyValue returns value of ManyValue field.
func (l *LanguagePackStringValuePluralized) GetManyValue() (value string) {
	return l.ManyValue
}

// GetOtherValue returns value of OtherValue field.
func (l *LanguagePackStringValuePluralized) GetOtherValue() (value string) {
	return l.OtherValue
}

// LanguagePackStringValueDeleted represents TL type `languagePackStringValueDeleted#6d5cb6fa`.
type LanguagePackStringValueDeleted struct {
}

// LanguagePackStringValueDeletedTypeID is TL type id of LanguagePackStringValueDeleted.
const LanguagePackStringValueDeletedTypeID = 0x6d5cb6fa

// construct implements constructor of LanguagePackStringValueClass.
func (l LanguagePackStringValueDeleted) construct() LanguagePackStringValueClass { return &l }

// Ensuring interfaces in compile-time for LanguagePackStringValueDeleted.
var (
	_ bin.Encoder     = &LanguagePackStringValueDeleted{}
	_ bin.Decoder     = &LanguagePackStringValueDeleted{}
	_ bin.BareEncoder = &LanguagePackStringValueDeleted{}
	_ bin.BareDecoder = &LanguagePackStringValueDeleted{}

	_ LanguagePackStringValueClass = &LanguagePackStringValueDeleted{}
)

func (l *LanguagePackStringValueDeleted) Zero() bool {
	if l == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (l *LanguagePackStringValueDeleted) String() string {
	if l == nil {
		return "LanguagePackStringValueDeleted(nil)"
	}
	type Alias LanguagePackStringValueDeleted
	return fmt.Sprintf("LanguagePackStringValueDeleted%+v", Alias(*l))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*LanguagePackStringValueDeleted) TypeID() uint32 {
	return LanguagePackStringValueDeletedTypeID
}

// TypeName returns name of type in TL schema.
func (*LanguagePackStringValueDeleted) TypeName() string {
	return "languagePackStringValueDeleted"
}

// TypeInfo returns info about TL type.
func (l *LanguagePackStringValueDeleted) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "languagePackStringValueDeleted",
		ID:   LanguagePackStringValueDeletedTypeID,
	}
	if l == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (l *LanguagePackStringValueDeleted) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode languagePackStringValueDeleted#6d5cb6fa as nil")
	}
	b.PutID(LanguagePackStringValueDeletedTypeID)
	return l.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (l *LanguagePackStringValueDeleted) EncodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode languagePackStringValueDeleted#6d5cb6fa as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (l *LanguagePackStringValueDeleted) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode languagePackStringValueDeleted#6d5cb6fa to nil")
	}
	if err := b.ConsumeID(LanguagePackStringValueDeletedTypeID); err != nil {
		return fmt.Errorf("unable to decode languagePackStringValueDeleted#6d5cb6fa: %w", err)
	}
	return l.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (l *LanguagePackStringValueDeleted) DecodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode languagePackStringValueDeleted#6d5cb6fa to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (l *LanguagePackStringValueDeleted) EncodeTDLibJSON(b tdjson.Encoder) error {
	if l == nil {
		return fmt.Errorf("can't encode languagePackStringValueDeleted#6d5cb6fa as nil")
	}
	b.ObjStart()
	b.PutID("languagePackStringValueDeleted")
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (l *LanguagePackStringValueDeleted) DecodeTDLibJSON(b tdjson.Decoder) error {
	if l == nil {
		return fmt.Errorf("can't decode languagePackStringValueDeleted#6d5cb6fa to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("languagePackStringValueDeleted"); err != nil {
				return fmt.Errorf("unable to decode languagePackStringValueDeleted#6d5cb6fa: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// LanguagePackStringValueClassName is schema name of LanguagePackStringValueClass.
const LanguagePackStringValueClassName = "LanguagePackStringValue"

// LanguagePackStringValueClass represents LanguagePackStringValue generic type.
//
// Example:
//  g, err := tdapi.DecodeLanguagePackStringValue(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tdapi.LanguagePackStringValueOrdinary: // languagePackStringValueOrdinary#f124a660
//  case *tdapi.LanguagePackStringValuePluralized: // languagePackStringValuePluralized#71a812c5
//  case *tdapi.LanguagePackStringValueDeleted: // languagePackStringValueDeleted#6d5cb6fa
//  default: panic(v)
//  }
type LanguagePackStringValueClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() LanguagePackStringValueClass

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

// DecodeLanguagePackStringValue implements binary de-serialization for LanguagePackStringValueClass.
func DecodeLanguagePackStringValue(buf *bin.Buffer) (LanguagePackStringValueClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case LanguagePackStringValueOrdinaryTypeID:
		// Decoding languagePackStringValueOrdinary#f124a660.
		v := LanguagePackStringValueOrdinary{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode LanguagePackStringValueClass: %w", err)
		}
		return &v, nil
	case LanguagePackStringValuePluralizedTypeID:
		// Decoding languagePackStringValuePluralized#71a812c5.
		v := LanguagePackStringValuePluralized{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode LanguagePackStringValueClass: %w", err)
		}
		return &v, nil
	case LanguagePackStringValueDeletedTypeID:
		// Decoding languagePackStringValueDeleted#6d5cb6fa.
		v := LanguagePackStringValueDeleted{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode LanguagePackStringValueClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode LanguagePackStringValueClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONLanguagePackStringValue implements binary de-serialization for LanguagePackStringValueClass.
func DecodeTDLibJSONLanguagePackStringValue(buf tdjson.Decoder) (LanguagePackStringValueClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "languagePackStringValueOrdinary":
		// Decoding languagePackStringValueOrdinary#f124a660.
		v := LanguagePackStringValueOrdinary{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode LanguagePackStringValueClass: %w", err)
		}
		return &v, nil
	case "languagePackStringValuePluralized":
		// Decoding languagePackStringValuePluralized#71a812c5.
		v := LanguagePackStringValuePluralized{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode LanguagePackStringValueClass: %w", err)
		}
		return &v, nil
	case "languagePackStringValueDeleted":
		// Decoding languagePackStringValueDeleted#6d5cb6fa.
		v := LanguagePackStringValueDeleted{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode LanguagePackStringValueClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode LanguagePackStringValueClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// LanguagePackStringValue boxes the LanguagePackStringValueClass providing a helper.
type LanguagePackStringValueBox struct {
	LanguagePackStringValue LanguagePackStringValueClass
}

// Decode implements bin.Decoder for LanguagePackStringValueBox.
func (b *LanguagePackStringValueBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode LanguagePackStringValueBox to nil")
	}
	v, err := DecodeLanguagePackStringValue(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.LanguagePackStringValue = v
	return nil
}

// Encode implements bin.Encode for LanguagePackStringValueBox.
func (b *LanguagePackStringValueBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.LanguagePackStringValue == nil {
		return fmt.Errorf("unable to encode LanguagePackStringValueClass as nil")
	}
	return b.LanguagePackStringValue.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for LanguagePackStringValueBox.
func (b *LanguagePackStringValueBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode LanguagePackStringValueBox to nil")
	}
	v, err := DecodeTDLibJSONLanguagePackStringValue(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.LanguagePackStringValue = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for LanguagePackStringValueBox.
func (b *LanguagePackStringValueBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.LanguagePackStringValue == nil {
		return fmt.Errorf("unable to encode LanguagePackStringValueClass as nil")
	}
	return b.LanguagePackStringValue.EncodeTDLibJSON(buf)
}
