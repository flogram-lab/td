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

// UploadReuploadCdnFileRequest represents TL type `upload.reuploadCdnFile#9b2754a8`.
// Request a reupload of a certain file to a CDN DC¹.
//
// Links:
//  1) https://core.telegram.org/cdn
//
// See https://core.telegram.org/method/upload.reuploadCdnFile for reference.
type UploadReuploadCdnFileRequest struct {
	// File token
	FileToken []byte
	// Request token
	RequestToken []byte
}

// UploadReuploadCdnFileRequestTypeID is TL type id of UploadReuploadCdnFileRequest.
const UploadReuploadCdnFileRequestTypeID = 0x9b2754a8

func (r *UploadReuploadCdnFileRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.FileToken == nil) {
		return false
	}
	if !(r.RequestToken == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *UploadReuploadCdnFileRequest) String() string {
	if r == nil {
		return "UploadReuploadCdnFileRequest(nil)"
	}
	type Alias UploadReuploadCdnFileRequest
	return fmt.Sprintf("UploadReuploadCdnFileRequest%+v", Alias(*r))
}

// FillFrom fills UploadReuploadCdnFileRequest from given interface.
func (r *UploadReuploadCdnFileRequest) FillFrom(from interface {
	GetFileToken() (value []byte)
	GetRequestToken() (value []byte)
}) {
	r.FileToken = from.GetFileToken()
	r.RequestToken = from.GetRequestToken()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UploadReuploadCdnFileRequest) TypeID() uint32 {
	return UploadReuploadCdnFileRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*UploadReuploadCdnFileRequest) TypeName() string {
	return "upload.reuploadCdnFile"
}

// TypeInfo returns info about TL type.
func (r *UploadReuploadCdnFileRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "upload.reuploadCdnFile",
		ID:   UploadReuploadCdnFileRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "FileToken",
			SchemaName: "file_token",
		},
		{
			Name:       "RequestToken",
			SchemaName: "request_token",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *UploadReuploadCdnFileRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode upload.reuploadCdnFile#9b2754a8 as nil")
	}
	b.PutID(UploadReuploadCdnFileRequestTypeID)
	b.PutBytes(r.FileToken)
	b.PutBytes(r.RequestToken)
	return nil
}

// GetFileToken returns value of FileToken field.
func (r *UploadReuploadCdnFileRequest) GetFileToken() (value []byte) {
	return r.FileToken
}

// GetRequestToken returns value of RequestToken field.
func (r *UploadReuploadCdnFileRequest) GetRequestToken() (value []byte) {
	return r.RequestToken
}

// Decode implements bin.Decoder.
func (r *UploadReuploadCdnFileRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode upload.reuploadCdnFile#9b2754a8 to nil")
	}
	if err := b.ConsumeID(UploadReuploadCdnFileRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode upload.reuploadCdnFile#9b2754a8: %w", err)
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode upload.reuploadCdnFile#9b2754a8: field file_token: %w", err)
		}
		r.FileToken = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode upload.reuploadCdnFile#9b2754a8: field request_token: %w", err)
		}
		r.RequestToken = value
	}
	return nil
}

// Ensuring interfaces in compile-time for UploadReuploadCdnFileRequest.
var (
	_ bin.Encoder = &UploadReuploadCdnFileRequest{}
	_ bin.Decoder = &UploadReuploadCdnFileRequest{}
)

// UploadReuploadCdnFile invokes method upload.reuploadCdnFile#9b2754a8 returning error if any.
// Request a reupload of a certain file to a CDN DC¹.
//
// Links:
//  1) https://core.telegram.org/cdn
//
// Possible errors:
//  400 RSA_DECRYPT_FAILED: Internal RSA decryption failed
//
// See https://core.telegram.org/method/upload.reuploadCdnFile for reference.
// Can be used by bots.
func (c *Client) UploadReuploadCdnFile(ctx context.Context, request *UploadReuploadCdnFileRequest) ([]FileHash, error) {
	var result FileHashVector

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return []FileHash(result.Elems), nil
}
