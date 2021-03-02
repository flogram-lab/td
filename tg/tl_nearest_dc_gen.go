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

// NearestDc represents TL type `nearestDc#8e1a1775`.
// Nearest data centre, according to geo-ip.
//
// See https://core.telegram.org/constructor/nearestDc for reference.
type NearestDc struct {
	// Country code determined by geo-ip
	Country string
	// Number of current data centre
	ThisDC int
	// Number of nearest data centre
	NearestDC int
}

// NearestDcTypeID is TL type id of NearestDc.
const NearestDcTypeID = 0x8e1a1775

func (n *NearestDc) Zero() bool {
	if n == nil {
		return true
	}
	if !(n.Country == "") {
		return false
	}
	if !(n.ThisDC == 0) {
		return false
	}
	if !(n.NearestDC == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (n *NearestDc) String() string {
	if n == nil {
		return "NearestDc(nil)"
	}
	type Alias NearestDc
	return fmt.Sprintf("NearestDc%+v", Alias(*n))
}

// FillFrom fills NearestDc from given interface.
func (n *NearestDc) FillFrom(from interface {
	GetCountry() (value string)
	GetThisDC() (value int)
	GetNearestDC() (value int)
}) {
	n.Country = from.GetCountry()
	n.ThisDC = from.GetThisDC()
	n.NearestDC = from.GetNearestDC()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*NearestDc) TypeID() uint32 {
	return NearestDcTypeID
}

// TypeName returns name of type in TL schema.
func (*NearestDc) TypeName() string {
	return "nearestDc"
}

// TypeInfo returns info about TL type.
func (n *NearestDc) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "nearestDc",
		ID:   NearestDcTypeID,
	}
	if n == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Country",
			SchemaName: "country",
		},
		{
			Name:       "ThisDC",
			SchemaName: "this_dc",
		},
		{
			Name:       "NearestDC",
			SchemaName: "nearest_dc",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (n *NearestDc) Encode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode nearestDc#8e1a1775 as nil")
	}
	b.PutID(NearestDcTypeID)
	b.PutString(n.Country)
	b.PutInt(n.ThisDC)
	b.PutInt(n.NearestDC)
	return nil
}

// GetCountry returns value of Country field.
func (n *NearestDc) GetCountry() (value string) {
	return n.Country
}

// GetThisDC returns value of ThisDC field.
func (n *NearestDc) GetThisDC() (value int) {
	return n.ThisDC
}

// GetNearestDC returns value of NearestDC field.
func (n *NearestDc) GetNearestDC() (value int) {
	return n.NearestDC
}

// Decode implements bin.Decoder.
func (n *NearestDc) Decode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode nearestDc#8e1a1775 to nil")
	}
	if err := b.ConsumeID(NearestDcTypeID); err != nil {
		return fmt.Errorf("unable to decode nearestDc#8e1a1775: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode nearestDc#8e1a1775: field country: %w", err)
		}
		n.Country = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode nearestDc#8e1a1775: field this_dc: %w", err)
		}
		n.ThisDC = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode nearestDc#8e1a1775: field nearest_dc: %w", err)
		}
		n.NearestDC = value
	}
	return nil
}

// Ensuring interfaces in compile-time for NearestDc.
var (
	_ bin.Encoder = &NearestDc{}
	_ bin.Decoder = &NearestDc{}
)
