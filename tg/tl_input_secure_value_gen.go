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

// InputSecureValue represents TL type `inputSecureValue#db21d0a7`.
// Secure value, for more info see the passport docs »¹
//
// Links:
//  1) https://core.telegram.org/passport/encryption#encryption
//
// See https://core.telegram.org/constructor/inputSecureValue for reference.
type InputSecureValue struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Secure passport¹ value type
	//
	// Links:
	//  1) https://core.telegram.org/passport
	Type SecureValueTypeClass
	// Encrypted Telegram Passport¹ element data
	//
	// Links:
	//  1) https://core.telegram.org/passport
	//
	// Use SetData and GetData helpers.
	Data SecureData
	// Encrypted passport¹ file with the front side of the document
	//
	// Links:
	//  1) https://core.telegram.org/passport
	//
	// Use SetFrontSide and GetFrontSide helpers.
	FrontSide InputSecureFileClass
	// Encrypted passport¹ file with the reverse side of the document
	//
	// Links:
	//  1) https://core.telegram.org/passport
	//
	// Use SetReverseSide and GetReverseSide helpers.
	ReverseSide InputSecureFileClass
	// Encrypted passport¹ file with a selfie of the user holding the document
	//
	// Links:
	//  1) https://core.telegram.org/passport
	//
	// Use SetSelfie and GetSelfie helpers.
	Selfie InputSecureFileClass
	// Array of encrypted passport¹ files with translated versions of the provided documents
	//
	// Links:
	//  1) https://core.telegram.org/passport
	//
	// Use SetTranslation and GetTranslation helpers.
	Translation []InputSecureFileClass
	// Array of encrypted passport¹ files with photos the of the documents
	//
	// Links:
	//  1) https://core.telegram.org/passport
	//
	// Use SetFiles and GetFiles helpers.
	Files []InputSecureFileClass
	// Plaintext verified passport¹ data
	//
	// Links:
	//  1) https://core.telegram.org/passport
	//
	// Use SetPlainData and GetPlainData helpers.
	PlainData SecurePlainDataClass
}

// InputSecureValueTypeID is TL type id of InputSecureValue.
const InputSecureValueTypeID = 0xdb21d0a7

func (i *InputSecureValue) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Flags.Zero()) {
		return false
	}
	if !(i.Type == nil) {
		return false
	}
	if !(i.Data.Zero()) {
		return false
	}
	if !(i.FrontSide == nil) {
		return false
	}
	if !(i.ReverseSide == nil) {
		return false
	}
	if !(i.Selfie == nil) {
		return false
	}
	if !(i.Translation == nil) {
		return false
	}
	if !(i.Files == nil) {
		return false
	}
	if !(i.PlainData == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputSecureValue) String() string {
	if i == nil {
		return "InputSecureValue(nil)"
	}
	type Alias InputSecureValue
	return fmt.Sprintf("InputSecureValue%+v", Alias(*i))
}

// FillFrom fills InputSecureValue from given interface.
func (i *InputSecureValue) FillFrom(from interface {
	GetType() (value SecureValueTypeClass)
	GetData() (value SecureData, ok bool)
	GetFrontSide() (value InputSecureFileClass, ok bool)
	GetReverseSide() (value InputSecureFileClass, ok bool)
	GetSelfie() (value InputSecureFileClass, ok bool)
	GetTranslation() (value []InputSecureFileClass, ok bool)
	GetFiles() (value []InputSecureFileClass, ok bool)
	GetPlainData() (value SecurePlainDataClass, ok bool)
}) {
	i.Type = from.GetType()
	if val, ok := from.GetData(); ok {
		i.Data = val
	}

	if val, ok := from.GetFrontSide(); ok {
		i.FrontSide = val
	}

	if val, ok := from.GetReverseSide(); ok {
		i.ReverseSide = val
	}

	if val, ok := from.GetSelfie(); ok {
		i.Selfie = val
	}

	if val, ok := from.GetTranslation(); ok {
		i.Translation = val
	}

	if val, ok := from.GetFiles(); ok {
		i.Files = val
	}

	if val, ok := from.GetPlainData(); ok {
		i.PlainData = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputSecureValue) TypeID() uint32 {
	return InputSecureValueTypeID
}

// TypeName returns name of type in TL schema.
func (*InputSecureValue) TypeName() string {
	return "inputSecureValue"
}

// TypeInfo returns info about TL type.
func (i *InputSecureValue) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputSecureValue",
		ID:   InputSecureValueTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Flags",
			SchemaName: "flags",
		},
		{
			Name:       "Type",
			SchemaName: "type",
		},
		{
			Name:       "Data",
			SchemaName: "data",
			Null:       !i.Flags.Has(0),
		},
		{
			Name:       "FrontSide",
			SchemaName: "front_side",
			Null:       !i.Flags.Has(1),
		},
		{
			Name:       "ReverseSide",
			SchemaName: "reverse_side",
			Null:       !i.Flags.Has(2),
		},
		{
			Name:       "Selfie",
			SchemaName: "selfie",
			Null:       !i.Flags.Has(3),
		},
		{
			Name:       "Translation",
			SchemaName: "translation",
			Null:       !i.Flags.Has(6),
		},
		{
			Name:       "Files",
			SchemaName: "files",
			Null:       !i.Flags.Has(4),
		},
		{
			Name:       "PlainData",
			SchemaName: "plain_data",
			Null:       !i.Flags.Has(5),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputSecureValue) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputSecureValue#db21d0a7 as nil")
	}
	b.PutID(InputSecureValueTypeID)
	if !(i.Data.Zero()) {
		i.Flags.Set(0)
	}
	if !(i.FrontSide == nil) {
		i.Flags.Set(1)
	}
	if !(i.ReverseSide == nil) {
		i.Flags.Set(2)
	}
	if !(i.Selfie == nil) {
		i.Flags.Set(3)
	}
	if !(i.Translation == nil) {
		i.Flags.Set(6)
	}
	if !(i.Files == nil) {
		i.Flags.Set(4)
	}
	if !(i.PlainData == nil) {
		i.Flags.Set(5)
	}
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputSecureValue#db21d0a7: field flags: %w", err)
	}
	if i.Type == nil {
		return fmt.Errorf("unable to encode inputSecureValue#db21d0a7: field type is nil")
	}
	if err := i.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputSecureValue#db21d0a7: field type: %w", err)
	}
	if i.Flags.Has(0) {
		if err := i.Data.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputSecureValue#db21d0a7: field data: %w", err)
		}
	}
	if i.Flags.Has(1) {
		if i.FrontSide == nil {
			return fmt.Errorf("unable to encode inputSecureValue#db21d0a7: field front_side is nil")
		}
		if err := i.FrontSide.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputSecureValue#db21d0a7: field front_side: %w", err)
		}
	}
	if i.Flags.Has(2) {
		if i.ReverseSide == nil {
			return fmt.Errorf("unable to encode inputSecureValue#db21d0a7: field reverse_side is nil")
		}
		if err := i.ReverseSide.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputSecureValue#db21d0a7: field reverse_side: %w", err)
		}
	}
	if i.Flags.Has(3) {
		if i.Selfie == nil {
			return fmt.Errorf("unable to encode inputSecureValue#db21d0a7: field selfie is nil")
		}
		if err := i.Selfie.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputSecureValue#db21d0a7: field selfie: %w", err)
		}
	}
	if i.Flags.Has(6) {
		b.PutVectorHeader(len(i.Translation))
		for idx, v := range i.Translation {
			if v == nil {
				return fmt.Errorf("unable to encode inputSecureValue#db21d0a7: field translation element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode inputSecureValue#db21d0a7: field translation element with index %d: %w", idx, err)
			}
		}
	}
	if i.Flags.Has(4) {
		b.PutVectorHeader(len(i.Files))
		for idx, v := range i.Files {
			if v == nil {
				return fmt.Errorf("unable to encode inputSecureValue#db21d0a7: field files element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode inputSecureValue#db21d0a7: field files element with index %d: %w", idx, err)
			}
		}
	}
	if i.Flags.Has(5) {
		if i.PlainData == nil {
			return fmt.Errorf("unable to encode inputSecureValue#db21d0a7: field plain_data is nil")
		}
		if err := i.PlainData.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputSecureValue#db21d0a7: field plain_data: %w", err)
		}
	}
	return nil
}

// GetType returns value of Type field.
func (i *InputSecureValue) GetType() (value SecureValueTypeClass) {
	return i.Type
}

// SetData sets value of Data conditional field.
func (i *InputSecureValue) SetData(value SecureData) {
	i.Flags.Set(0)
	i.Data = value
}

// GetData returns value of Data conditional field and
// boolean which is true if field was set.
func (i *InputSecureValue) GetData() (value SecureData, ok bool) {
	if !i.Flags.Has(0) {
		return value, false
	}
	return i.Data, true
}

// SetFrontSide sets value of FrontSide conditional field.
func (i *InputSecureValue) SetFrontSide(value InputSecureFileClass) {
	i.Flags.Set(1)
	i.FrontSide = value
}

// GetFrontSide returns value of FrontSide conditional field and
// boolean which is true if field was set.
func (i *InputSecureValue) GetFrontSide() (value InputSecureFileClass, ok bool) {
	if !i.Flags.Has(1) {
		return value, false
	}
	return i.FrontSide, true
}

// SetReverseSide sets value of ReverseSide conditional field.
func (i *InputSecureValue) SetReverseSide(value InputSecureFileClass) {
	i.Flags.Set(2)
	i.ReverseSide = value
}

// GetReverseSide returns value of ReverseSide conditional field and
// boolean which is true if field was set.
func (i *InputSecureValue) GetReverseSide() (value InputSecureFileClass, ok bool) {
	if !i.Flags.Has(2) {
		return value, false
	}
	return i.ReverseSide, true
}

// SetSelfie sets value of Selfie conditional field.
func (i *InputSecureValue) SetSelfie(value InputSecureFileClass) {
	i.Flags.Set(3)
	i.Selfie = value
}

// GetSelfie returns value of Selfie conditional field and
// boolean which is true if field was set.
func (i *InputSecureValue) GetSelfie() (value InputSecureFileClass, ok bool) {
	if !i.Flags.Has(3) {
		return value, false
	}
	return i.Selfie, true
}

// SetTranslation sets value of Translation conditional field.
func (i *InputSecureValue) SetTranslation(value []InputSecureFileClass) {
	i.Flags.Set(6)
	i.Translation = value
}

// GetTranslation returns value of Translation conditional field and
// boolean which is true if field was set.
func (i *InputSecureValue) GetTranslation() (value []InputSecureFileClass, ok bool) {
	if !i.Flags.Has(6) {
		return value, false
	}
	return i.Translation, true
}

// MapTranslation returns field Translation wrapped in InputSecureFileClassArray helper.
func (i *InputSecureValue) MapTranslation() (value InputSecureFileClassArray, ok bool) {
	if !i.Flags.Has(6) {
		return value, false
	}
	return InputSecureFileClassArray(i.Translation), true
}

// SetFiles sets value of Files conditional field.
func (i *InputSecureValue) SetFiles(value []InputSecureFileClass) {
	i.Flags.Set(4)
	i.Files = value
}

// GetFiles returns value of Files conditional field and
// boolean which is true if field was set.
func (i *InputSecureValue) GetFiles() (value []InputSecureFileClass, ok bool) {
	if !i.Flags.Has(4) {
		return value, false
	}
	return i.Files, true
}

// MapFiles returns field Files wrapped in InputSecureFileClassArray helper.
func (i *InputSecureValue) MapFiles() (value InputSecureFileClassArray, ok bool) {
	if !i.Flags.Has(4) {
		return value, false
	}
	return InputSecureFileClassArray(i.Files), true
}

// SetPlainData sets value of PlainData conditional field.
func (i *InputSecureValue) SetPlainData(value SecurePlainDataClass) {
	i.Flags.Set(5)
	i.PlainData = value
}

// GetPlainData returns value of PlainData conditional field and
// boolean which is true if field was set.
func (i *InputSecureValue) GetPlainData() (value SecurePlainDataClass, ok bool) {
	if !i.Flags.Has(5) {
		return value, false
	}
	return i.PlainData, true
}

// Decode implements bin.Decoder.
func (i *InputSecureValue) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputSecureValue#db21d0a7 to nil")
	}
	if err := b.ConsumeID(InputSecureValueTypeID); err != nil {
		return fmt.Errorf("unable to decode inputSecureValue#db21d0a7: %w", err)
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputSecureValue#db21d0a7: field flags: %w", err)
		}
	}
	{
		value, err := DecodeSecureValueType(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputSecureValue#db21d0a7: field type: %w", err)
		}
		i.Type = value
	}
	if i.Flags.Has(0) {
		if err := i.Data.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputSecureValue#db21d0a7: field data: %w", err)
		}
	}
	if i.Flags.Has(1) {
		value, err := DecodeInputSecureFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputSecureValue#db21d0a7: field front_side: %w", err)
		}
		i.FrontSide = value
	}
	if i.Flags.Has(2) {
		value, err := DecodeInputSecureFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputSecureValue#db21d0a7: field reverse_side: %w", err)
		}
		i.ReverseSide = value
	}
	if i.Flags.Has(3) {
		value, err := DecodeInputSecureFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputSecureValue#db21d0a7: field selfie: %w", err)
		}
		i.Selfie = value
	}
	if i.Flags.Has(6) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode inputSecureValue#db21d0a7: field translation: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputSecureFile(b)
			if err != nil {
				return fmt.Errorf("unable to decode inputSecureValue#db21d0a7: field translation: %w", err)
			}
			i.Translation = append(i.Translation, value)
		}
	}
	if i.Flags.Has(4) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode inputSecureValue#db21d0a7: field files: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputSecureFile(b)
			if err != nil {
				return fmt.Errorf("unable to decode inputSecureValue#db21d0a7: field files: %w", err)
			}
			i.Files = append(i.Files, value)
		}
	}
	if i.Flags.Has(5) {
		value, err := DecodeSecurePlainData(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputSecureValue#db21d0a7: field plain_data: %w", err)
		}
		i.PlainData = value
	}
	return nil
}

// Ensuring interfaces in compile-time for InputSecureValue.
var (
	_ bin.Encoder = &InputSecureValue{}
	_ bin.Decoder = &InputSecureValue{}
)
