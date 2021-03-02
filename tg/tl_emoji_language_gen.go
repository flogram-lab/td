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

// EmojiLanguage represents TL type `emojiLanguage#b3fb5361`.
// Emoji language
//
// See https://core.telegram.org/constructor/emojiLanguage for reference.
type EmojiLanguage struct {
	// Language code
	LangCode string
}

// EmojiLanguageTypeID is TL type id of EmojiLanguage.
const EmojiLanguageTypeID = 0xb3fb5361

func (e *EmojiLanguage) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.LangCode == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EmojiLanguage) String() string {
	if e == nil {
		return "EmojiLanguage(nil)"
	}
	type Alias EmojiLanguage
	return fmt.Sprintf("EmojiLanguage%+v", Alias(*e))
}

// FillFrom fills EmojiLanguage from given interface.
func (e *EmojiLanguage) FillFrom(from interface {
	GetLangCode() (value string)
}) {
	e.LangCode = from.GetLangCode()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EmojiLanguage) TypeID() uint32 {
	return EmojiLanguageTypeID
}

// TypeName returns name of type in TL schema.
func (*EmojiLanguage) TypeName() string {
	return "emojiLanguage"
}

// TypeInfo returns info about TL type.
func (e *EmojiLanguage) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "emojiLanguage",
		ID:   EmojiLanguageTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "LangCode",
			SchemaName: "lang_code",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EmojiLanguage) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode emojiLanguage#b3fb5361 as nil")
	}
	b.PutID(EmojiLanguageTypeID)
	b.PutString(e.LangCode)
	return nil
}

// GetLangCode returns value of LangCode field.
func (e *EmojiLanguage) GetLangCode() (value string) {
	return e.LangCode
}

// Decode implements bin.Decoder.
func (e *EmojiLanguage) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode emojiLanguage#b3fb5361 to nil")
	}
	if err := b.ConsumeID(EmojiLanguageTypeID); err != nil {
		return fmt.Errorf("unable to decode emojiLanguage#b3fb5361: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode emojiLanguage#b3fb5361: field lang_code: %w", err)
		}
		e.LangCode = value
	}
	return nil
}

// Ensuring interfaces in compile-time for EmojiLanguage.
var (
	_ bin.Encoder = &EmojiLanguage{}
	_ bin.Decoder = &EmojiLanguage{}
)
