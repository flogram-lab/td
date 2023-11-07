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

// GetChatBoostsRequest represents TL type `getChatBoosts#ab5eaa38`.
type GetChatBoostsRequest struct {
	// Identifier of the chat
	ChatID int64
	// Pass true to receive only boosts received from gift codes and giveaways created by the
	// chat
	OnlyGiftCodes bool
	// Offset of the first entry to return as received from the previous request; use empty
	// string to get the first chunk of results
	Offset string
	// The maximum number of boosts to be returned; up to 100. For optimal performance, the
	// number of returned boosts can be smaller than the specified limit
	Limit int32
}

// GetChatBoostsRequestTypeID is TL type id of GetChatBoostsRequest.
const GetChatBoostsRequestTypeID = 0xab5eaa38

// Ensuring interfaces in compile-time for GetChatBoostsRequest.
var (
	_ bin.Encoder     = &GetChatBoostsRequest{}
	_ bin.Decoder     = &GetChatBoostsRequest{}
	_ bin.BareEncoder = &GetChatBoostsRequest{}
	_ bin.BareDecoder = &GetChatBoostsRequest{}
)

func (g *GetChatBoostsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.OnlyGiftCodes == false) {
		return false
	}
	if !(g.Offset == "") {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetChatBoostsRequest) String() string {
	if g == nil {
		return "GetChatBoostsRequest(nil)"
	}
	type Alias GetChatBoostsRequest
	return fmt.Sprintf("GetChatBoostsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetChatBoostsRequest) TypeID() uint32 {
	return GetChatBoostsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetChatBoostsRequest) TypeName() string {
	return "getChatBoosts"
}

// TypeInfo returns info about TL type.
func (g *GetChatBoostsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getChatBoosts",
		ID:   GetChatBoostsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "OnlyGiftCodes",
			SchemaName: "only_gift_codes",
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetChatBoostsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatBoosts#ab5eaa38 as nil")
	}
	b.PutID(GetChatBoostsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetChatBoostsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatBoosts#ab5eaa38 as nil")
	}
	b.PutInt53(g.ChatID)
	b.PutBool(g.OnlyGiftCodes)
	b.PutString(g.Offset)
	b.PutInt32(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetChatBoostsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatBoosts#ab5eaa38 to nil")
	}
	if err := b.ConsumeID(GetChatBoostsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getChatBoosts#ab5eaa38: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetChatBoostsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatBoosts#ab5eaa38 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getChatBoosts#ab5eaa38: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode getChatBoosts#ab5eaa38: field only_gift_codes: %w", err)
		}
		g.OnlyGiftCodes = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getChatBoosts#ab5eaa38: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getChatBoosts#ab5eaa38: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetChatBoostsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatBoosts#ab5eaa38 as nil")
	}
	b.ObjStart()
	b.PutID("getChatBoosts")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.Comma()
	b.FieldStart("only_gift_codes")
	b.PutBool(g.OnlyGiftCodes)
	b.Comma()
	b.FieldStart("offset")
	b.PutString(g.Offset)
	b.Comma()
	b.FieldStart("limit")
	b.PutInt32(g.Limit)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetChatBoostsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatBoosts#ab5eaa38 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getChatBoosts"); err != nil {
				return fmt.Errorf("unable to decode getChatBoosts#ab5eaa38: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getChatBoosts#ab5eaa38: field chat_id: %w", err)
			}
			g.ChatID = value
		case "only_gift_codes":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode getChatBoosts#ab5eaa38: field only_gift_codes: %w", err)
			}
			g.OnlyGiftCodes = value
		case "offset":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getChatBoosts#ab5eaa38: field offset: %w", err)
			}
			g.Offset = value
		case "limit":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getChatBoosts#ab5eaa38: field limit: %w", err)
			}
			g.Limit = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetChatBoostsRequest) GetChatID() (value int64) {
	if g == nil {
		return
	}
	return g.ChatID
}

// GetOnlyGiftCodes returns value of OnlyGiftCodes field.
func (g *GetChatBoostsRequest) GetOnlyGiftCodes() (value bool) {
	if g == nil {
		return
	}
	return g.OnlyGiftCodes
}

// GetOffset returns value of Offset field.
func (g *GetChatBoostsRequest) GetOffset() (value string) {
	if g == nil {
		return
	}
	return g.Offset
}

// GetLimit returns value of Limit field.
func (g *GetChatBoostsRequest) GetLimit() (value int32) {
	if g == nil {
		return
	}
	return g.Limit
}

// GetChatBoosts invokes method getChatBoosts#ab5eaa38 returning error if any.
func (c *Client) GetChatBoosts(ctx context.Context, request *GetChatBoostsRequest) (*FoundChatBoosts, error) {
	var result FoundChatBoosts

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
