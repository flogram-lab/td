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

// MessagesGetSavedGifsRequest represents TL type `messages.getSavedGifs#83bf3d52`.
// Get saved GIFs
//
// See https://core.telegram.org/method/messages.getSavedGifs for reference.
type MessagesGetSavedGifsRequest struct {
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int
}

// MessagesGetSavedGifsRequestTypeID is TL type id of MessagesGetSavedGifsRequest.
const MessagesGetSavedGifsRequestTypeID = 0x83bf3d52

func (g *MessagesGetSavedGifsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetSavedGifsRequest) String() string {
	if g == nil {
		return "MessagesGetSavedGifsRequest(nil)"
	}
	type Alias MessagesGetSavedGifsRequest
	return fmt.Sprintf("MessagesGetSavedGifsRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetSavedGifsRequest from given interface.
func (g *MessagesGetSavedGifsRequest) FillFrom(from interface {
	GetHash() (value int)
}) {
	g.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetSavedGifsRequest) TypeID() uint32 {
	return MessagesGetSavedGifsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetSavedGifsRequest) TypeName() string {
	return "messages.getSavedGifs"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetSavedGifsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getSavedGifs",
		ID:   MessagesGetSavedGifsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetSavedGifsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getSavedGifs#83bf3d52 as nil")
	}
	b.PutID(MessagesGetSavedGifsRequestTypeID)
	b.PutInt(g.Hash)
	return nil
}

// GetHash returns value of Hash field.
func (g *MessagesGetSavedGifsRequest) GetHash() (value int) {
	return g.Hash
}

// Decode implements bin.Decoder.
func (g *MessagesGetSavedGifsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getSavedGifs#83bf3d52 to nil")
	}
	if err := b.ConsumeID(MessagesGetSavedGifsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getSavedGifs#83bf3d52: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getSavedGifs#83bf3d52: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetSavedGifsRequest.
var (
	_ bin.Encoder = &MessagesGetSavedGifsRequest{}
	_ bin.Decoder = &MessagesGetSavedGifsRequest{}
)

// MessagesGetSavedGifs invokes method messages.getSavedGifs#83bf3d52 returning error if any.
// Get saved GIFs
//
// See https://core.telegram.org/method/messages.getSavedGifs for reference.
func (c *Client) MessagesGetSavedGifs(ctx context.Context, hash int) (MessagesSavedGifsClass, error) {
	var result MessagesSavedGifsBox

	request := &MessagesGetSavedGifsRequest{
		Hash: hash,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.SavedGifs, nil
}
