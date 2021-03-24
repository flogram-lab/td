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

// StickersChangeStickerPositionRequest represents TL type `stickers.changeStickerPosition#ffb6d4ca`.
// Changes the absolute position of a sticker in the set to which it belongs; for bots
// only. The sticker set must have been created by the bot
//
// See https://core.telegram.org/method/stickers.changeStickerPosition for reference.
type StickersChangeStickerPositionRequest struct {
	// The sticker
	Sticker InputDocumentClass
	// The new position of the sticker, zero-based
	Position int
}

// StickersChangeStickerPositionRequestTypeID is TL type id of StickersChangeStickerPositionRequest.
const StickersChangeStickerPositionRequestTypeID = 0xffb6d4ca

func (c *StickersChangeStickerPositionRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Sticker == nil) {
		return false
	}
	if !(c.Position == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *StickersChangeStickerPositionRequest) String() string {
	if c == nil {
		return "StickersChangeStickerPositionRequest(nil)"
	}
	type Alias StickersChangeStickerPositionRequest
	return fmt.Sprintf("StickersChangeStickerPositionRequest%+v", Alias(*c))
}

// FillFrom fills StickersChangeStickerPositionRequest from given interface.
func (c *StickersChangeStickerPositionRequest) FillFrom(from interface {
	GetSticker() (value InputDocumentClass)
	GetPosition() (value int)
}) {
	c.Sticker = from.GetSticker()
	c.Position = from.GetPosition()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StickersChangeStickerPositionRequest) TypeID() uint32 {
	return StickersChangeStickerPositionRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StickersChangeStickerPositionRequest) TypeName() string {
	return "stickers.changeStickerPosition"
}

// TypeInfo returns info about TL type.
func (c *StickersChangeStickerPositionRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stickers.changeStickerPosition",
		ID:   StickersChangeStickerPositionRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Sticker",
			SchemaName: "sticker",
		},
		{
			Name:       "Position",
			SchemaName: "position",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *StickersChangeStickerPositionRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode stickers.changeStickerPosition#ffb6d4ca as nil")
	}
	b.PutID(StickersChangeStickerPositionRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *StickersChangeStickerPositionRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode stickers.changeStickerPosition#ffb6d4ca as nil")
	}
	if c.Sticker == nil {
		return fmt.Errorf("unable to encode stickers.changeStickerPosition#ffb6d4ca: field sticker is nil")
	}
	if err := c.Sticker.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stickers.changeStickerPosition#ffb6d4ca: field sticker: %w", err)
	}
	b.PutInt(c.Position)
	return nil
}

// GetSticker returns value of Sticker field.
func (c *StickersChangeStickerPositionRequest) GetSticker() (value InputDocumentClass) {
	return c.Sticker
}

// GetStickerAsNotEmpty returns mapped value of Sticker field.
func (c *StickersChangeStickerPositionRequest) GetStickerAsNotEmpty() (*InputDocument, bool) {
	return c.Sticker.AsNotEmpty()
}

// GetPosition returns value of Position field.
func (c *StickersChangeStickerPositionRequest) GetPosition() (value int) {
	return c.Position
}

// Decode implements bin.Decoder.
func (c *StickersChangeStickerPositionRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode stickers.changeStickerPosition#ffb6d4ca to nil")
	}
	if err := b.ConsumeID(StickersChangeStickerPositionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stickers.changeStickerPosition#ffb6d4ca: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *StickersChangeStickerPositionRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode stickers.changeStickerPosition#ffb6d4ca to nil")
	}
	{
		value, err := DecodeInputDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode stickers.changeStickerPosition#ffb6d4ca: field sticker: %w", err)
		}
		c.Sticker = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stickers.changeStickerPosition#ffb6d4ca: field position: %w", err)
		}
		c.Position = value
	}
	return nil
}

// Ensuring interfaces in compile-time for StickersChangeStickerPositionRequest.
var (
	_ bin.Encoder     = &StickersChangeStickerPositionRequest{}
	_ bin.Decoder     = &StickersChangeStickerPositionRequest{}
	_ bin.BareEncoder = &StickersChangeStickerPositionRequest{}
	_ bin.BareDecoder = &StickersChangeStickerPositionRequest{}
)

// StickersChangeStickerPosition invokes method stickers.changeStickerPosition#ffb6d4ca returning error if any.
// Changes the absolute position of a sticker in the set to which it belongs; for bots
// only. The sticker set must have been created by the bot
//
// Possible errors:
//  400 BOT_MISSING: This method can only be run by a bot
//  400 STICKER_INVALID: The provided sticker is invalid
//
// See https://core.telegram.org/method/stickers.changeStickerPosition for reference.
// Can be used by bots.
func (c *Client) StickersChangeStickerPosition(ctx context.Context, request *StickersChangeStickerPositionRequest) (*MessagesStickerSet, error) {
	var result MessagesStickerSet

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
