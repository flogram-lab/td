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

// UploadCDNFileReuploadNeeded represents TL type `upload.cdnFileReuploadNeeded#eea8e46e`.
// The file was cleared from the temporary RAM cache of the CDN¹ and has to be
// reuploaded.
//
// Links:
//  1) https://core.telegram.org/cdn
//
// See https://core.telegram.org/constructor/upload.cdnFileReuploadNeeded for reference.
type UploadCDNFileReuploadNeeded struct {
	// Request token (see CDN¹)
	//
	// Links:
	//  1) https://core.telegram.org/cdn
	RequestToken []byte
}

// UploadCDNFileReuploadNeededTypeID is TL type id of UploadCDNFileReuploadNeeded.
const UploadCDNFileReuploadNeededTypeID = 0xeea8e46e

// construct implements constructor of UploadCDNFileClass.
func (c UploadCDNFileReuploadNeeded) construct() UploadCDNFileClass { return &c }

// Ensuring interfaces in compile-time for UploadCDNFileReuploadNeeded.
var (
	_ bin.Encoder     = &UploadCDNFileReuploadNeeded{}
	_ bin.Decoder     = &UploadCDNFileReuploadNeeded{}
	_ bin.BareEncoder = &UploadCDNFileReuploadNeeded{}
	_ bin.BareDecoder = &UploadCDNFileReuploadNeeded{}

	_ UploadCDNFileClass = &UploadCDNFileReuploadNeeded{}
)

func (c *UploadCDNFileReuploadNeeded) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.RequestToken == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *UploadCDNFileReuploadNeeded) String() string {
	if c == nil {
		return "UploadCDNFileReuploadNeeded(nil)"
	}
	type Alias UploadCDNFileReuploadNeeded
	return fmt.Sprintf("UploadCDNFileReuploadNeeded%+v", Alias(*c))
}

// FillFrom fills UploadCDNFileReuploadNeeded from given interface.
func (c *UploadCDNFileReuploadNeeded) FillFrom(from interface {
	GetRequestToken() (value []byte)
}) {
	c.RequestToken = from.GetRequestToken()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UploadCDNFileReuploadNeeded) TypeID() uint32 {
	return UploadCDNFileReuploadNeededTypeID
}

// TypeName returns name of type in TL schema.
func (*UploadCDNFileReuploadNeeded) TypeName() string {
	return "upload.cdnFileReuploadNeeded"
}

// TypeInfo returns info about TL type.
func (c *UploadCDNFileReuploadNeeded) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "upload.cdnFileReuploadNeeded",
		ID:   UploadCDNFileReuploadNeededTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "RequestToken",
			SchemaName: "request_token",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *UploadCDNFileReuploadNeeded) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode upload.cdnFileReuploadNeeded#eea8e46e as nil")
	}
	b.PutID(UploadCDNFileReuploadNeededTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *UploadCDNFileReuploadNeeded) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode upload.cdnFileReuploadNeeded#eea8e46e as nil")
	}
	b.PutBytes(c.RequestToken)
	return nil
}

// Decode implements bin.Decoder.
func (c *UploadCDNFileReuploadNeeded) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode upload.cdnFileReuploadNeeded#eea8e46e to nil")
	}
	if err := b.ConsumeID(UploadCDNFileReuploadNeededTypeID); err != nil {
		return fmt.Errorf("unable to decode upload.cdnFileReuploadNeeded#eea8e46e: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *UploadCDNFileReuploadNeeded) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode upload.cdnFileReuploadNeeded#eea8e46e to nil")
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode upload.cdnFileReuploadNeeded#eea8e46e: field request_token: %w", err)
		}
		c.RequestToken = value
	}
	return nil
}

// GetRequestToken returns value of RequestToken field.
func (c *UploadCDNFileReuploadNeeded) GetRequestToken() (value []byte) {
	return c.RequestToken
}

// UploadCDNFile represents TL type `upload.cdnFile#a99fca4f`.
// Represent a chunk of a CDN¹ file.
//
// Links:
//  1) https://core.telegram.org/cdn
//
// See https://core.telegram.org/constructor/upload.cdnFile for reference.
type UploadCDNFile struct {
	// The data
	Bytes []byte
}

// UploadCDNFileTypeID is TL type id of UploadCDNFile.
const UploadCDNFileTypeID = 0xa99fca4f

// construct implements constructor of UploadCDNFileClass.
func (c UploadCDNFile) construct() UploadCDNFileClass { return &c }

// Ensuring interfaces in compile-time for UploadCDNFile.
var (
	_ bin.Encoder     = &UploadCDNFile{}
	_ bin.Decoder     = &UploadCDNFile{}
	_ bin.BareEncoder = &UploadCDNFile{}
	_ bin.BareDecoder = &UploadCDNFile{}

	_ UploadCDNFileClass = &UploadCDNFile{}
)

func (c *UploadCDNFile) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Bytes == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *UploadCDNFile) String() string {
	if c == nil {
		return "UploadCDNFile(nil)"
	}
	type Alias UploadCDNFile
	return fmt.Sprintf("UploadCDNFile%+v", Alias(*c))
}

// FillFrom fills UploadCDNFile from given interface.
func (c *UploadCDNFile) FillFrom(from interface {
	GetBytes() (value []byte)
}) {
	c.Bytes = from.GetBytes()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UploadCDNFile) TypeID() uint32 {
	return UploadCDNFileTypeID
}

// TypeName returns name of type in TL schema.
func (*UploadCDNFile) TypeName() string {
	return "upload.cdnFile"
}

// TypeInfo returns info about TL type.
func (c *UploadCDNFile) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "upload.cdnFile",
		ID:   UploadCDNFileTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Bytes",
			SchemaName: "bytes",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *UploadCDNFile) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode upload.cdnFile#a99fca4f as nil")
	}
	b.PutID(UploadCDNFileTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *UploadCDNFile) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode upload.cdnFile#a99fca4f as nil")
	}
	b.PutBytes(c.Bytes)
	return nil
}

// Decode implements bin.Decoder.
func (c *UploadCDNFile) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode upload.cdnFile#a99fca4f to nil")
	}
	if err := b.ConsumeID(UploadCDNFileTypeID); err != nil {
		return fmt.Errorf("unable to decode upload.cdnFile#a99fca4f: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *UploadCDNFile) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode upload.cdnFile#a99fca4f to nil")
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode upload.cdnFile#a99fca4f: field bytes: %w", err)
		}
		c.Bytes = value
	}
	return nil
}

// GetBytes returns value of Bytes field.
func (c *UploadCDNFile) GetBytes() (value []byte) {
	return c.Bytes
}

// UploadCDNFileClassName is schema name of UploadCDNFileClass.
const UploadCDNFileClassName = "upload.CdnFile"

// UploadCDNFileClass represents upload.CdnFile generic type.
//
// See https://core.telegram.org/type/upload.CdnFile for reference.
//
// Example:
//  g, err := tg.DecodeUploadCDNFile(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.UploadCDNFileReuploadNeeded: // upload.cdnFileReuploadNeeded#eea8e46e
//  case *tg.UploadCDNFile: // upload.cdnFile#a99fca4f
//  default: panic(v)
//  }
type UploadCDNFileClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() UploadCDNFileClass

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
}

// DecodeUploadCDNFile implements binary de-serialization for UploadCDNFileClass.
func DecodeUploadCDNFile(buf *bin.Buffer) (UploadCDNFileClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case UploadCDNFileReuploadNeededTypeID:
		// Decoding upload.cdnFileReuploadNeeded#eea8e46e.
		v := UploadCDNFileReuploadNeeded{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UploadCDNFileClass: %w", err)
		}
		return &v, nil
	case UploadCDNFileTypeID:
		// Decoding upload.cdnFile#a99fca4f.
		v := UploadCDNFile{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UploadCDNFileClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode UploadCDNFileClass: %w", bin.NewUnexpectedID(id))
	}
}

// UploadCDNFile boxes the UploadCDNFileClass providing a helper.
type UploadCDNFileBox struct {
	CdnFile UploadCDNFileClass
}

// Decode implements bin.Decoder for UploadCDNFileBox.
func (b *UploadCDNFileBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode UploadCDNFileBox to nil")
	}
	v, err := DecodeUploadCDNFile(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.CdnFile = v
	return nil
}

// Encode implements bin.Encode for UploadCDNFileBox.
func (b *UploadCDNFileBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.CdnFile == nil {
		return fmt.Errorf("unable to encode UploadCDNFileClass as nil")
	}
	return b.CdnFile.Encode(buf)
}
