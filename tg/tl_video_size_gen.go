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

// VideoSize represents TL type `videoSize#e831c556`.
// Animated profile picture¹ in MPEG4 format
//
// Links:
//  1) https://core.telegram.org/api/files#animated-profile-pictures
//
// See https://core.telegram.org/constructor/videoSize for reference.
type VideoSize struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// u for animated profile pictures, and v for trimmed and downscaled video previews
	Type string
	// File location
	Location FileLocationToBeDeprecated
	// Video width
	W int
	// Video height
	H int
	// File size
	Size int
	// Timestamp that should be shown as static preview to the user (seconds)
	//
	// Use SetVideoStartTs and GetVideoStartTs helpers.
	VideoStartTs float64
}

// VideoSizeTypeID is TL type id of VideoSize.
const VideoSizeTypeID = 0xe831c556

func (v *VideoSize) Zero() bool {
	if v == nil {
		return true
	}
	if !(v.Flags.Zero()) {
		return false
	}
	if !(v.Type == "") {
		return false
	}
	if !(v.Location.Zero()) {
		return false
	}
	if !(v.W == 0) {
		return false
	}
	if !(v.H == 0) {
		return false
	}
	if !(v.Size == 0) {
		return false
	}
	if !(v.VideoStartTs == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (v *VideoSize) String() string {
	if v == nil {
		return "VideoSize(nil)"
	}
	type Alias VideoSize
	return fmt.Sprintf("VideoSize%+v", Alias(*v))
}

// FillFrom fills VideoSize from given interface.
func (v *VideoSize) FillFrom(from interface {
	GetType() (value string)
	GetLocation() (value FileLocationToBeDeprecated)
	GetW() (value int)
	GetH() (value int)
	GetSize() (value int)
	GetVideoStartTs() (value float64, ok bool)
}) {
	v.Type = from.GetType()
	v.Location = from.GetLocation()
	v.W = from.GetW()
	v.H = from.GetH()
	v.Size = from.GetSize()
	if val, ok := from.GetVideoStartTs(); ok {
		v.VideoStartTs = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*VideoSize) TypeID() uint32 {
	return VideoSizeTypeID
}

// TypeName returns name of type in TL schema.
func (*VideoSize) TypeName() string {
	return "videoSize"
}

// TypeInfo returns info about TL type.
func (v *VideoSize) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "videoSize",
		ID:   VideoSizeTypeID,
	}
	if v == nil {
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
			Name:       "Location",
			SchemaName: "location",
		},
		{
			Name:       "W",
			SchemaName: "w",
		},
		{
			Name:       "H",
			SchemaName: "h",
		},
		{
			Name:       "Size",
			SchemaName: "size",
		},
		{
			Name:       "VideoStartTs",
			SchemaName: "video_start_ts",
			Null:       !v.Flags.Has(0),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (v *VideoSize) Encode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode videoSize#e831c556 as nil")
	}
	b.PutID(VideoSizeTypeID)
	if !(v.VideoStartTs == 0) {
		v.Flags.Set(0)
	}
	if err := v.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode videoSize#e831c556: field flags: %w", err)
	}
	b.PutString(v.Type)
	if err := v.Location.Encode(b); err != nil {
		return fmt.Errorf("unable to encode videoSize#e831c556: field location: %w", err)
	}
	b.PutInt(v.W)
	b.PutInt(v.H)
	b.PutInt(v.Size)
	if v.Flags.Has(0) {
		b.PutDouble(v.VideoStartTs)
	}
	return nil
}

// GetType returns value of Type field.
func (v *VideoSize) GetType() (value string) {
	return v.Type
}

// GetLocation returns value of Location field.
func (v *VideoSize) GetLocation() (value FileLocationToBeDeprecated) {
	return v.Location
}

// GetW returns value of W field.
func (v *VideoSize) GetW() (value int) {
	return v.W
}

// GetH returns value of H field.
func (v *VideoSize) GetH() (value int) {
	return v.H
}

// GetSize returns value of Size field.
func (v *VideoSize) GetSize() (value int) {
	return v.Size
}

// SetVideoStartTs sets value of VideoStartTs conditional field.
func (v *VideoSize) SetVideoStartTs(value float64) {
	v.Flags.Set(0)
	v.VideoStartTs = value
}

// GetVideoStartTs returns value of VideoStartTs conditional field and
// boolean which is true if field was set.
func (v *VideoSize) GetVideoStartTs() (value float64, ok bool) {
	if !v.Flags.Has(0) {
		return value, false
	}
	return v.VideoStartTs, true
}

// Decode implements bin.Decoder.
func (v *VideoSize) Decode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode videoSize#e831c556 to nil")
	}
	if err := b.ConsumeID(VideoSizeTypeID); err != nil {
		return fmt.Errorf("unable to decode videoSize#e831c556: %w", err)
	}
	{
		if err := v.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode videoSize#e831c556: field flags: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode videoSize#e831c556: field type: %w", err)
		}
		v.Type = value
	}
	{
		if err := v.Location.Decode(b); err != nil {
			return fmt.Errorf("unable to decode videoSize#e831c556: field location: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode videoSize#e831c556: field w: %w", err)
		}
		v.W = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode videoSize#e831c556: field h: %w", err)
		}
		v.H = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode videoSize#e831c556: field size: %w", err)
		}
		v.Size = value
	}
	if v.Flags.Has(0) {
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode videoSize#e831c556: field video_start_ts: %w", err)
		}
		v.VideoStartTs = value
	}
	return nil
}

// Ensuring interfaces in compile-time for VideoSize.
var (
	_ bin.Encoder = &VideoSize{}
	_ bin.Decoder = &VideoSize{}
)
