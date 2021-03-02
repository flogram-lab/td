// Code generated by gotdgen, DO NOT EDIT.

package mt

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

// GetFutureSaltsRequest represents TL type `get_future_salts#b921bd04`.
type GetFutureSaltsRequest struct {
	// Num field of GetFutureSaltsRequest.
	Num int
}

// GetFutureSaltsRequestTypeID is TL type id of GetFutureSaltsRequest.
const GetFutureSaltsRequestTypeID = 0xb921bd04

func (g *GetFutureSaltsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Num == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetFutureSaltsRequest) String() string {
	if g == nil {
		return "GetFutureSaltsRequest(nil)"
	}
	type Alias GetFutureSaltsRequest
	return fmt.Sprintf("GetFutureSaltsRequest%+v", Alias(*g))
}

// FillFrom fills GetFutureSaltsRequest from given interface.
func (g *GetFutureSaltsRequest) FillFrom(from interface {
	GetNum() (value int)
}) {
	g.Num = from.GetNum()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetFutureSaltsRequest) TypeID() uint32 {
	return GetFutureSaltsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetFutureSaltsRequest) TypeName() string {
	return "get_future_salts"
}

// TypeInfo returns info about TL type.
func (g *GetFutureSaltsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "get_future_salts",
		ID:   GetFutureSaltsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Num",
			SchemaName: "num",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetFutureSaltsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode get_future_salts#b921bd04 as nil")
	}
	b.PutID(GetFutureSaltsRequestTypeID)
	b.PutInt(g.Num)
	return nil
}

// GetNum returns value of Num field.
func (g *GetFutureSaltsRequest) GetNum() (value int) {
	return g.Num
}

// Decode implements bin.Decoder.
func (g *GetFutureSaltsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode get_future_salts#b921bd04 to nil")
	}
	if err := b.ConsumeID(GetFutureSaltsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode get_future_salts#b921bd04: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode get_future_salts#b921bd04: field num: %w", err)
		}
		g.Num = value
	}
	return nil
}

// Ensuring interfaces in compile-time for GetFutureSaltsRequest.
var (
	_ bin.Encoder = &GetFutureSaltsRequest{}
	_ bin.Decoder = &GetFutureSaltsRequest{}
)

// GetFutureSalts invokes method get_future_salts#b921bd04 returning error if any.
func (c *Client) GetFutureSalts(ctx context.Context, num int) (*FutureSalts, error) {
	var result FutureSalts

	request := &GetFutureSaltsRequest{
		Num: num,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
