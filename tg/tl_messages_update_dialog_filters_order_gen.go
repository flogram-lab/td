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

// MessagesUpdateDialogFiltersOrderRequest represents TL type `messages.updateDialogFiltersOrder#c563c1e4`.
// Reorder folders¹
//
// Links:
//  1) https://core.telegram.org/api/folders
//
// See https://core.telegram.org/method/messages.updateDialogFiltersOrder for reference.
type MessagesUpdateDialogFiltersOrderRequest struct {
	// New folder¹ order
	//
	// Links:
	//  1) https://core.telegram.org/api/folders
	Order []int
}

// MessagesUpdateDialogFiltersOrderRequestTypeID is TL type id of MessagesUpdateDialogFiltersOrderRequest.
const MessagesUpdateDialogFiltersOrderRequestTypeID = 0xc563c1e4

func (u *MessagesUpdateDialogFiltersOrderRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Order == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *MessagesUpdateDialogFiltersOrderRequest) String() string {
	if u == nil {
		return "MessagesUpdateDialogFiltersOrderRequest(nil)"
	}
	type Alias MessagesUpdateDialogFiltersOrderRequest
	return fmt.Sprintf("MessagesUpdateDialogFiltersOrderRequest%+v", Alias(*u))
}

// FillFrom fills MessagesUpdateDialogFiltersOrderRequest from given interface.
func (u *MessagesUpdateDialogFiltersOrderRequest) FillFrom(from interface {
	GetOrder() (value []int)
}) {
	u.Order = from.GetOrder()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesUpdateDialogFiltersOrderRequest) TypeID() uint32 {
	return MessagesUpdateDialogFiltersOrderRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesUpdateDialogFiltersOrderRequest) TypeName() string {
	return "messages.updateDialogFiltersOrder"
}

// TypeInfo returns info about TL type.
func (u *MessagesUpdateDialogFiltersOrderRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.updateDialogFiltersOrder",
		ID:   MessagesUpdateDialogFiltersOrderRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Order",
			SchemaName: "order",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *MessagesUpdateDialogFiltersOrderRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode messages.updateDialogFiltersOrder#c563c1e4 as nil")
	}
	b.PutID(MessagesUpdateDialogFiltersOrderRequestTypeID)
	b.PutVectorHeader(len(u.Order))
	for _, v := range u.Order {
		b.PutInt(v)
	}
	return nil
}

// GetOrder returns value of Order field.
func (u *MessagesUpdateDialogFiltersOrderRequest) GetOrder() (value []int) {
	return u.Order
}

// Decode implements bin.Decoder.
func (u *MessagesUpdateDialogFiltersOrderRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode messages.updateDialogFiltersOrder#c563c1e4 to nil")
	}
	if err := b.ConsumeID(MessagesUpdateDialogFiltersOrderRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.updateDialogFiltersOrder#c563c1e4: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.updateDialogFiltersOrder#c563c1e4: field order: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode messages.updateDialogFiltersOrder#c563c1e4: field order: %w", err)
			}
			u.Order = append(u.Order, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesUpdateDialogFiltersOrderRequest.
var (
	_ bin.Encoder = &MessagesUpdateDialogFiltersOrderRequest{}
	_ bin.Decoder = &MessagesUpdateDialogFiltersOrderRequest{}
)

// MessagesUpdateDialogFiltersOrder invokes method messages.updateDialogFiltersOrder#c563c1e4 returning error if any.
// Reorder folders¹
//
// Links:
//  1) https://core.telegram.org/api/folders
//
// See https://core.telegram.org/method/messages.updateDialogFiltersOrder for reference.
func (c *Client) MessagesUpdateDialogFiltersOrder(ctx context.Context, order []int) (bool, error) {
	var result BoolBox

	request := &MessagesUpdateDialogFiltersOrderRequest{
		Order: order,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
