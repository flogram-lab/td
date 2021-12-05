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

// ChannelLocationEmpty represents TL type `channelLocationEmpty#bfb5ad8b`.
// No location (normal supergroup)
//
// See https://core.telegram.org/constructor/channelLocationEmpty for reference.
type ChannelLocationEmpty struct {
}

// ChannelLocationEmptyTypeID is TL type id of ChannelLocationEmpty.
const ChannelLocationEmptyTypeID = 0xbfb5ad8b

// construct implements constructor of ChannelLocationClass.
func (c ChannelLocationEmpty) construct() ChannelLocationClass { return &c }

// Ensuring interfaces in compile-time for ChannelLocationEmpty.
var (
	_ bin.Encoder     = &ChannelLocationEmpty{}
	_ bin.Decoder     = &ChannelLocationEmpty{}
	_ bin.BareEncoder = &ChannelLocationEmpty{}
	_ bin.BareDecoder = &ChannelLocationEmpty{}

	_ ChannelLocationClass = &ChannelLocationEmpty{}
)

func (c *ChannelLocationEmpty) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChannelLocationEmpty) String() string {
	if c == nil {
		return "ChannelLocationEmpty(nil)"
	}
	type Alias ChannelLocationEmpty
	return fmt.Sprintf("ChannelLocationEmpty%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelLocationEmpty) TypeID() uint32 {
	return ChannelLocationEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelLocationEmpty) TypeName() string {
	return "channelLocationEmpty"
}

// TypeInfo returns info about TL type.
func (c *ChannelLocationEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channelLocationEmpty",
		ID:   ChannelLocationEmptyTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChannelLocationEmpty) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelLocationEmpty#bfb5ad8b as nil")
	}
	b.PutID(ChannelLocationEmptyTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChannelLocationEmpty) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelLocationEmpty#bfb5ad8b as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChannelLocationEmpty) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelLocationEmpty#bfb5ad8b to nil")
	}
	if err := b.ConsumeID(ChannelLocationEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode channelLocationEmpty#bfb5ad8b: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChannelLocationEmpty) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelLocationEmpty#bfb5ad8b to nil")
	}
	return nil
}

// ChannelLocation represents TL type `channelLocation#209b82db`.
// Geographical location of supergroup (geogroups)
//
// See https://core.telegram.org/constructor/channelLocation for reference.
type ChannelLocation struct {
	// Geographical location of supergrup
	GeoPoint GeoPointClass
	// Textual description of the address
	Address string
}

// ChannelLocationTypeID is TL type id of ChannelLocation.
const ChannelLocationTypeID = 0x209b82db

// construct implements constructor of ChannelLocationClass.
func (c ChannelLocation) construct() ChannelLocationClass { return &c }

// Ensuring interfaces in compile-time for ChannelLocation.
var (
	_ bin.Encoder     = &ChannelLocation{}
	_ bin.Decoder     = &ChannelLocation{}
	_ bin.BareEncoder = &ChannelLocation{}
	_ bin.BareDecoder = &ChannelLocation{}

	_ ChannelLocationClass = &ChannelLocation{}
)

func (c *ChannelLocation) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.GeoPoint == nil) {
		return false
	}
	if !(c.Address == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChannelLocation) String() string {
	if c == nil {
		return "ChannelLocation(nil)"
	}
	type Alias ChannelLocation
	return fmt.Sprintf("ChannelLocation%+v", Alias(*c))
}

// FillFrom fills ChannelLocation from given interface.
func (c *ChannelLocation) FillFrom(from interface {
	GetGeoPoint() (value GeoPointClass)
	GetAddress() (value string)
}) {
	c.GeoPoint = from.GetGeoPoint()
	c.Address = from.GetAddress()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelLocation) TypeID() uint32 {
	return ChannelLocationTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelLocation) TypeName() string {
	return "channelLocation"
}

// TypeInfo returns info about TL type.
func (c *ChannelLocation) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channelLocation",
		ID:   ChannelLocationTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "GeoPoint",
			SchemaName: "geo_point",
		},
		{
			Name:       "Address",
			SchemaName: "address",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChannelLocation) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelLocation#209b82db as nil")
	}
	b.PutID(ChannelLocationTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChannelLocation) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelLocation#209b82db as nil")
	}
	if c.GeoPoint == nil {
		return fmt.Errorf("unable to encode channelLocation#209b82db: field geo_point is nil")
	}
	if err := c.GeoPoint.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channelLocation#209b82db: field geo_point: %w", err)
	}
	b.PutString(c.Address)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChannelLocation) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelLocation#209b82db to nil")
	}
	if err := b.ConsumeID(ChannelLocationTypeID); err != nil {
		return fmt.Errorf("unable to decode channelLocation#209b82db: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChannelLocation) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelLocation#209b82db to nil")
	}
	{
		value, err := DecodeGeoPoint(b)
		if err != nil {
			return fmt.Errorf("unable to decode channelLocation#209b82db: field geo_point: %w", err)
		}
		c.GeoPoint = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channelLocation#209b82db: field address: %w", err)
		}
		c.Address = value
	}
	return nil
}

// GetGeoPoint returns value of GeoPoint field.
func (c *ChannelLocation) GetGeoPoint() (value GeoPointClass) {
	return c.GeoPoint
}

// GetAddress returns value of Address field.
func (c *ChannelLocation) GetAddress() (value string) {
	return c.Address
}

// ChannelLocationClassName is schema name of ChannelLocationClass.
const ChannelLocationClassName = "ChannelLocation"

// ChannelLocationClass represents ChannelLocation generic type.
//
// See https://core.telegram.org/type/ChannelLocation for reference.
//
// Example:
//  g, err := tg.DecodeChannelLocation(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.ChannelLocationEmpty: // channelLocationEmpty#bfb5ad8b
//  case *tg.ChannelLocation: // channelLocation#209b82db
//  default: panic(v)
//  }
type ChannelLocationClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() ChannelLocationClass

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

	// AsNotEmpty tries to map ChannelLocationClass to ChannelLocation.
	AsNotEmpty() (*ChannelLocation, bool)
}

// AsNotEmpty tries to map ChannelLocationEmpty to ChannelLocation.
func (c *ChannelLocationEmpty) AsNotEmpty() (*ChannelLocation, bool) {
	return nil, false
}

// AsNotEmpty tries to map ChannelLocation to ChannelLocation.
func (c *ChannelLocation) AsNotEmpty() (*ChannelLocation, bool) {
	return c, true
}

// DecodeChannelLocation implements binary de-serialization for ChannelLocationClass.
func DecodeChannelLocation(buf *bin.Buffer) (ChannelLocationClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ChannelLocationEmptyTypeID:
		// Decoding channelLocationEmpty#bfb5ad8b.
		v := ChannelLocationEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChannelLocationClass: %w", err)
		}
		return &v, nil
	case ChannelLocationTypeID:
		// Decoding channelLocation#209b82db.
		v := ChannelLocation{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChannelLocationClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ChannelLocationClass: %w", bin.NewUnexpectedID(id))
	}
}

// ChannelLocation boxes the ChannelLocationClass providing a helper.
type ChannelLocationBox struct {
	ChannelLocation ChannelLocationClass
}

// Decode implements bin.Decoder for ChannelLocationBox.
func (b *ChannelLocationBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ChannelLocationBox to nil")
	}
	v, err := DecodeChannelLocation(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ChannelLocation = v
	return nil
}

// Encode implements bin.Encode for ChannelLocationBox.
func (b *ChannelLocationBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ChannelLocation == nil {
		return fmt.Errorf("unable to encode ChannelLocationClass as nil")
	}
	return b.ChannelLocation.Encode(buf)
}
