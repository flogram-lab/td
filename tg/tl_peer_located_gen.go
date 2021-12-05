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

// PeerLocated represents TL type `peerLocated#ca461b5d`.
// Peer geolocated nearby
//
// See https://core.telegram.org/constructor/peerLocated for reference.
type PeerLocated struct {
	// Peer
	Peer PeerClass
	// Validity period of current data
	Expires int
	// Distance from the peer in meters
	Distance int
}

// PeerLocatedTypeID is TL type id of PeerLocated.
const PeerLocatedTypeID = 0xca461b5d

// construct implements constructor of PeerLocatedClass.
func (p PeerLocated) construct() PeerLocatedClass { return &p }

// Ensuring interfaces in compile-time for PeerLocated.
var (
	_ bin.Encoder     = &PeerLocated{}
	_ bin.Decoder     = &PeerLocated{}
	_ bin.BareEncoder = &PeerLocated{}
	_ bin.BareDecoder = &PeerLocated{}

	_ PeerLocatedClass = &PeerLocated{}
)

func (p *PeerLocated) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Peer == nil) {
		return false
	}
	if !(p.Expires == 0) {
		return false
	}
	if !(p.Distance == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PeerLocated) String() string {
	if p == nil {
		return "PeerLocated(nil)"
	}
	type Alias PeerLocated
	return fmt.Sprintf("PeerLocated%+v", Alias(*p))
}

// FillFrom fills PeerLocated from given interface.
func (p *PeerLocated) FillFrom(from interface {
	GetPeer() (value PeerClass)
	GetExpires() (value int)
	GetDistance() (value int)
}) {
	p.Peer = from.GetPeer()
	p.Expires = from.GetExpires()
	p.Distance = from.GetDistance()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PeerLocated) TypeID() uint32 {
	return PeerLocatedTypeID
}

// TypeName returns name of type in TL schema.
func (*PeerLocated) TypeName() string {
	return "peerLocated"
}

// TypeInfo returns info about TL type.
func (p *PeerLocated) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "peerLocated",
		ID:   PeerLocatedTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Expires",
			SchemaName: "expires",
		},
		{
			Name:       "Distance",
			SchemaName: "distance",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PeerLocated) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode peerLocated#ca461b5d as nil")
	}
	b.PutID(PeerLocatedTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PeerLocated) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode peerLocated#ca461b5d as nil")
	}
	if p.Peer == nil {
		return fmt.Errorf("unable to encode peerLocated#ca461b5d: field peer is nil")
	}
	if err := p.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode peerLocated#ca461b5d: field peer: %w", err)
	}
	b.PutInt(p.Expires)
	b.PutInt(p.Distance)
	return nil
}

// Decode implements bin.Decoder.
func (p *PeerLocated) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode peerLocated#ca461b5d to nil")
	}
	if err := b.ConsumeID(PeerLocatedTypeID); err != nil {
		return fmt.Errorf("unable to decode peerLocated#ca461b5d: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PeerLocated) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode peerLocated#ca461b5d to nil")
	}
	{
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode peerLocated#ca461b5d: field peer: %w", err)
		}
		p.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode peerLocated#ca461b5d: field expires: %w", err)
		}
		p.Expires = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode peerLocated#ca461b5d: field distance: %w", err)
		}
		p.Distance = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (p *PeerLocated) GetPeer() (value PeerClass) {
	return p.Peer
}

// GetExpires returns value of Expires field.
func (p *PeerLocated) GetExpires() (value int) {
	return p.Expires
}

// GetDistance returns value of Distance field.
func (p *PeerLocated) GetDistance() (value int) {
	return p.Distance
}

// PeerSelfLocated represents TL type `peerSelfLocated#f8ec284b`.
// Current peer
//
// See https://core.telegram.org/constructor/peerSelfLocated for reference.
type PeerSelfLocated struct {
	// Expiry of geolocation info for current peer
	Expires int
}

// PeerSelfLocatedTypeID is TL type id of PeerSelfLocated.
const PeerSelfLocatedTypeID = 0xf8ec284b

// construct implements constructor of PeerLocatedClass.
func (p PeerSelfLocated) construct() PeerLocatedClass { return &p }

// Ensuring interfaces in compile-time for PeerSelfLocated.
var (
	_ bin.Encoder     = &PeerSelfLocated{}
	_ bin.Decoder     = &PeerSelfLocated{}
	_ bin.BareEncoder = &PeerSelfLocated{}
	_ bin.BareDecoder = &PeerSelfLocated{}

	_ PeerLocatedClass = &PeerSelfLocated{}
)

func (p *PeerSelfLocated) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Expires == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PeerSelfLocated) String() string {
	if p == nil {
		return "PeerSelfLocated(nil)"
	}
	type Alias PeerSelfLocated
	return fmt.Sprintf("PeerSelfLocated%+v", Alias(*p))
}

// FillFrom fills PeerSelfLocated from given interface.
func (p *PeerSelfLocated) FillFrom(from interface {
	GetExpires() (value int)
}) {
	p.Expires = from.GetExpires()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PeerSelfLocated) TypeID() uint32 {
	return PeerSelfLocatedTypeID
}

// TypeName returns name of type in TL schema.
func (*PeerSelfLocated) TypeName() string {
	return "peerSelfLocated"
}

// TypeInfo returns info about TL type.
func (p *PeerSelfLocated) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "peerSelfLocated",
		ID:   PeerSelfLocatedTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Expires",
			SchemaName: "expires",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PeerSelfLocated) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode peerSelfLocated#f8ec284b as nil")
	}
	b.PutID(PeerSelfLocatedTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PeerSelfLocated) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode peerSelfLocated#f8ec284b as nil")
	}
	b.PutInt(p.Expires)
	return nil
}

// Decode implements bin.Decoder.
func (p *PeerSelfLocated) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode peerSelfLocated#f8ec284b to nil")
	}
	if err := b.ConsumeID(PeerSelfLocatedTypeID); err != nil {
		return fmt.Errorf("unable to decode peerSelfLocated#f8ec284b: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PeerSelfLocated) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode peerSelfLocated#f8ec284b to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode peerSelfLocated#f8ec284b: field expires: %w", err)
		}
		p.Expires = value
	}
	return nil
}

// GetExpires returns value of Expires field.
func (p *PeerSelfLocated) GetExpires() (value int) {
	return p.Expires
}

// PeerLocatedClassName is schema name of PeerLocatedClass.
const PeerLocatedClassName = "PeerLocated"

// PeerLocatedClass represents PeerLocated generic type.
//
// See https://core.telegram.org/type/PeerLocated for reference.
//
// Example:
//  g, err := tg.DecodePeerLocated(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.PeerLocated: // peerLocated#ca461b5d
//  case *tg.PeerSelfLocated: // peerSelfLocated#f8ec284b
//  default: panic(v)
//  }
type PeerLocatedClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() PeerLocatedClass

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

	// Validity period of current data
	GetExpires() (value int)
}

// DecodePeerLocated implements binary de-serialization for PeerLocatedClass.
func DecodePeerLocated(buf *bin.Buffer) (PeerLocatedClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case PeerLocatedTypeID:
		// Decoding peerLocated#ca461b5d.
		v := PeerLocated{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PeerLocatedClass: %w", err)
		}
		return &v, nil
	case PeerSelfLocatedTypeID:
		// Decoding peerSelfLocated#f8ec284b.
		v := PeerSelfLocated{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PeerLocatedClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PeerLocatedClass: %w", bin.NewUnexpectedID(id))
	}
}

// PeerLocated boxes the PeerLocatedClass providing a helper.
type PeerLocatedBox struct {
	PeerLocated PeerLocatedClass
}

// Decode implements bin.Decoder for PeerLocatedBox.
func (b *PeerLocatedBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode PeerLocatedBox to nil")
	}
	v, err := DecodePeerLocated(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.PeerLocated = v
	return nil
}

// Encode implements bin.Encode for PeerLocatedBox.
func (b *PeerLocatedBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.PeerLocated == nil {
		return fmt.Errorf("unable to encode PeerLocatedClass as nil")
	}
	return b.PeerLocated.Encode(buf)
}
