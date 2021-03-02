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

// WebDocument represents TL type `webDocument#1c570ed1`.
// Remote document
//
// See https://core.telegram.org/constructor/webDocument for reference.
type WebDocument struct {
	// Document URL
	URL string
	// Access hash
	AccessHash int64
	// File size
	Size int
	// MIME type
	MimeType string
	// Attributes for media types
	Attributes []DocumentAttributeClass
}

// WebDocumentTypeID is TL type id of WebDocument.
const WebDocumentTypeID = 0x1c570ed1

func (w *WebDocument) Zero() bool {
	if w == nil {
		return true
	}
	if !(w.URL == "") {
		return false
	}
	if !(w.AccessHash == 0) {
		return false
	}
	if !(w.Size == 0) {
		return false
	}
	if !(w.MimeType == "") {
		return false
	}
	if !(w.Attributes == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (w *WebDocument) String() string {
	if w == nil {
		return "WebDocument(nil)"
	}
	type Alias WebDocument
	return fmt.Sprintf("WebDocument%+v", Alias(*w))
}

// FillFrom fills WebDocument from given interface.
func (w *WebDocument) FillFrom(from interface {
	GetURL() (value string)
	GetAccessHash() (value int64)
	GetSize() (value int)
	GetMimeType() (value string)
	GetAttributes() (value []DocumentAttributeClass)
}) {
	w.URL = from.GetURL()
	w.AccessHash = from.GetAccessHash()
	w.Size = from.GetSize()
	w.MimeType = from.GetMimeType()
	w.Attributes = from.GetAttributes()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*WebDocument) TypeID() uint32 {
	return WebDocumentTypeID
}

// TypeName returns name of type in TL schema.
func (*WebDocument) TypeName() string {
	return "webDocument"
}

// TypeInfo returns info about TL type.
func (w *WebDocument) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "webDocument",
		ID:   WebDocumentTypeID,
	}
	if w == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
		{
			Name:       "AccessHash",
			SchemaName: "access_hash",
		},
		{
			Name:       "Size",
			SchemaName: "size",
		},
		{
			Name:       "MimeType",
			SchemaName: "mime_type",
		},
		{
			Name:       "Attributes",
			SchemaName: "attributes",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (w *WebDocument) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webDocument#1c570ed1 as nil")
	}
	b.PutID(WebDocumentTypeID)
	b.PutString(w.URL)
	b.PutLong(w.AccessHash)
	b.PutInt(w.Size)
	b.PutString(w.MimeType)
	b.PutVectorHeader(len(w.Attributes))
	for idx, v := range w.Attributes {
		if v == nil {
			return fmt.Errorf("unable to encode webDocument#1c570ed1: field attributes element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode webDocument#1c570ed1: field attributes element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetURL returns value of URL field.
func (w *WebDocument) GetURL() (value string) {
	return w.URL
}

// GetAccessHash returns value of AccessHash field.
func (w *WebDocument) GetAccessHash() (value int64) {
	return w.AccessHash
}

// GetSize returns value of Size field.
func (w *WebDocument) GetSize() (value int) {
	return w.Size
}

// GetMimeType returns value of MimeType field.
func (w *WebDocument) GetMimeType() (value string) {
	return w.MimeType
}

// GetAttributes returns value of Attributes field.
func (w *WebDocument) GetAttributes() (value []DocumentAttributeClass) {
	return w.Attributes
}

// MapAttributes returns field Attributes wrapped in DocumentAttributeClassArray helper.
func (w *WebDocument) MapAttributes() (value DocumentAttributeClassArray) {
	return DocumentAttributeClassArray(w.Attributes)
}

// Decode implements bin.Decoder.
func (w *WebDocument) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webDocument#1c570ed1 to nil")
	}
	if err := b.ConsumeID(WebDocumentTypeID); err != nil {
		return fmt.Errorf("unable to decode webDocument#1c570ed1: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webDocument#1c570ed1: field url: %w", err)
		}
		w.URL = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode webDocument#1c570ed1: field access_hash: %w", err)
		}
		w.AccessHash = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode webDocument#1c570ed1: field size: %w", err)
		}
		w.Size = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webDocument#1c570ed1: field mime_type: %w", err)
		}
		w.MimeType = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode webDocument#1c570ed1: field attributes: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDocumentAttribute(b)
			if err != nil {
				return fmt.Errorf("unable to decode webDocument#1c570ed1: field attributes: %w", err)
			}
			w.Attributes = append(w.Attributes, value)
		}
	}
	return nil
}

// construct implements constructor of WebDocumentClass.
func (w WebDocument) construct() WebDocumentClass { return &w }

// Ensuring interfaces in compile-time for WebDocument.
var (
	_ bin.Encoder = &WebDocument{}
	_ bin.Decoder = &WebDocument{}

	_ WebDocumentClass = &WebDocument{}
)

// WebDocumentNoProxy represents TL type `webDocumentNoProxy#f9c8bcc6`.
// Remote document that can be downloaded without proxying through telegram¹
//
// Links:
//  1) https://core.telegram.org/api/files
//
// See https://core.telegram.org/constructor/webDocumentNoProxy for reference.
type WebDocumentNoProxy struct {
	// Document URL
	URL string
	// File size
	Size int
	// MIME type
	MimeType string
	// Attributes for media types
	Attributes []DocumentAttributeClass
}

// WebDocumentNoProxyTypeID is TL type id of WebDocumentNoProxy.
const WebDocumentNoProxyTypeID = 0xf9c8bcc6

func (w *WebDocumentNoProxy) Zero() bool {
	if w == nil {
		return true
	}
	if !(w.URL == "") {
		return false
	}
	if !(w.Size == 0) {
		return false
	}
	if !(w.MimeType == "") {
		return false
	}
	if !(w.Attributes == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (w *WebDocumentNoProxy) String() string {
	if w == nil {
		return "WebDocumentNoProxy(nil)"
	}
	type Alias WebDocumentNoProxy
	return fmt.Sprintf("WebDocumentNoProxy%+v", Alias(*w))
}

// FillFrom fills WebDocumentNoProxy from given interface.
func (w *WebDocumentNoProxy) FillFrom(from interface {
	GetURL() (value string)
	GetSize() (value int)
	GetMimeType() (value string)
	GetAttributes() (value []DocumentAttributeClass)
}) {
	w.URL = from.GetURL()
	w.Size = from.GetSize()
	w.MimeType = from.GetMimeType()
	w.Attributes = from.GetAttributes()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*WebDocumentNoProxy) TypeID() uint32 {
	return WebDocumentNoProxyTypeID
}

// TypeName returns name of type in TL schema.
func (*WebDocumentNoProxy) TypeName() string {
	return "webDocumentNoProxy"
}

// TypeInfo returns info about TL type.
func (w *WebDocumentNoProxy) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "webDocumentNoProxy",
		ID:   WebDocumentNoProxyTypeID,
	}
	if w == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
		{
			Name:       "Size",
			SchemaName: "size",
		},
		{
			Name:       "MimeType",
			SchemaName: "mime_type",
		},
		{
			Name:       "Attributes",
			SchemaName: "attributes",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (w *WebDocumentNoProxy) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webDocumentNoProxy#f9c8bcc6 as nil")
	}
	b.PutID(WebDocumentNoProxyTypeID)
	b.PutString(w.URL)
	b.PutInt(w.Size)
	b.PutString(w.MimeType)
	b.PutVectorHeader(len(w.Attributes))
	for idx, v := range w.Attributes {
		if v == nil {
			return fmt.Errorf("unable to encode webDocumentNoProxy#f9c8bcc6: field attributes element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode webDocumentNoProxy#f9c8bcc6: field attributes element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetURL returns value of URL field.
func (w *WebDocumentNoProxy) GetURL() (value string) {
	return w.URL
}

// GetSize returns value of Size field.
func (w *WebDocumentNoProxy) GetSize() (value int) {
	return w.Size
}

// GetMimeType returns value of MimeType field.
func (w *WebDocumentNoProxy) GetMimeType() (value string) {
	return w.MimeType
}

// GetAttributes returns value of Attributes field.
func (w *WebDocumentNoProxy) GetAttributes() (value []DocumentAttributeClass) {
	return w.Attributes
}

// MapAttributes returns field Attributes wrapped in DocumentAttributeClassArray helper.
func (w *WebDocumentNoProxy) MapAttributes() (value DocumentAttributeClassArray) {
	return DocumentAttributeClassArray(w.Attributes)
}

// Decode implements bin.Decoder.
func (w *WebDocumentNoProxy) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webDocumentNoProxy#f9c8bcc6 to nil")
	}
	if err := b.ConsumeID(WebDocumentNoProxyTypeID); err != nil {
		return fmt.Errorf("unable to decode webDocumentNoProxy#f9c8bcc6: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webDocumentNoProxy#f9c8bcc6: field url: %w", err)
		}
		w.URL = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode webDocumentNoProxy#f9c8bcc6: field size: %w", err)
		}
		w.Size = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webDocumentNoProxy#f9c8bcc6: field mime_type: %w", err)
		}
		w.MimeType = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode webDocumentNoProxy#f9c8bcc6: field attributes: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDocumentAttribute(b)
			if err != nil {
				return fmt.Errorf("unable to decode webDocumentNoProxy#f9c8bcc6: field attributes: %w", err)
			}
			w.Attributes = append(w.Attributes, value)
		}
	}
	return nil
}

// construct implements constructor of WebDocumentClass.
func (w WebDocumentNoProxy) construct() WebDocumentClass { return &w }

// Ensuring interfaces in compile-time for WebDocumentNoProxy.
var (
	_ bin.Encoder = &WebDocumentNoProxy{}
	_ bin.Decoder = &WebDocumentNoProxy{}

	_ WebDocumentClass = &WebDocumentNoProxy{}
)

// WebDocumentClass represents WebDocument generic type.
//
// See https://core.telegram.org/type/WebDocument for reference.
//
// Example:
//  g, err := tg.DecodeWebDocument(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.WebDocument: // webDocument#1c570ed1
//  case *tg.WebDocumentNoProxy: // webDocumentNoProxy#f9c8bcc6
//  default: panic(v)
//  }
type WebDocumentClass interface {
	bin.Encoder
	bin.Decoder
	construct() WebDocumentClass

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

	// Document URL
	GetURL() (value string)
	// File size
	GetSize() (value int)
	// MIME type
	GetMimeType() (value string)
	// Attributes for media types
	GetAttributes() (value []DocumentAttributeClass)
	// Attributes for media types
	MapAttributes() (value DocumentAttributeClassArray)
}

// AsInput tries to map WebDocument to InputWebDocument.
func (w *WebDocument) AsInput() *InputWebDocument {
	value := new(InputWebDocument)
	value.URL = w.GetURL()
	value.Size = w.GetSize()
	value.MimeType = w.GetMimeType()
	value.Attributes = w.GetAttributes()

	return value
}

// DecodeWebDocument implements binary de-serialization for WebDocumentClass.
func DecodeWebDocument(buf *bin.Buffer) (WebDocumentClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case WebDocumentTypeID:
		// Decoding webDocument#1c570ed1.
		v := WebDocument{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode WebDocumentClass: %w", err)
		}
		return &v, nil
	case WebDocumentNoProxyTypeID:
		// Decoding webDocumentNoProxy#f9c8bcc6.
		v := WebDocumentNoProxy{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode WebDocumentClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode WebDocumentClass: %w", bin.NewUnexpectedID(id))
	}
}

// WebDocument boxes the WebDocumentClass providing a helper.
type WebDocumentBox struct {
	WebDocument WebDocumentClass
}

// Decode implements bin.Decoder for WebDocumentBox.
func (b *WebDocumentBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode WebDocumentBox to nil")
	}
	v, err := DecodeWebDocument(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.WebDocument = v
	return nil
}

// Encode implements bin.Encode for WebDocumentBox.
func (b *WebDocumentBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.WebDocument == nil {
		return fmt.Errorf("unable to encode WebDocumentClass as nil")
	}
	return b.WebDocument.Encode(buf)
}

// WebDocumentClassArray is adapter for slice of WebDocumentClass.
type WebDocumentClassArray []WebDocumentClass

// Sort sorts slice of WebDocumentClass.
func (s WebDocumentClassArray) Sort(less func(a, b WebDocumentClass) bool) WebDocumentClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of WebDocumentClass.
func (s WebDocumentClassArray) SortStable(less func(a, b WebDocumentClass) bool) WebDocumentClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of WebDocumentClass.
func (s WebDocumentClassArray) Retain(keep func(x WebDocumentClass) bool) WebDocumentClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s WebDocumentClassArray) First() (v WebDocumentClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s WebDocumentClassArray) Last() (v WebDocumentClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *WebDocumentClassArray) PopFirst() (v WebDocumentClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero WebDocumentClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *WebDocumentClassArray) Pop() (v WebDocumentClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsWebDocument returns copy with only WebDocument constructors.
func (s WebDocumentClassArray) AsWebDocument() (to WebDocumentArray) {
	for _, elem := range s {
		value, ok := elem.(*WebDocument)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsWebDocumentNoProxy returns copy with only WebDocumentNoProxy constructors.
func (s WebDocumentClassArray) AsWebDocumentNoProxy() (to WebDocumentNoProxyArray) {
	for _, elem := range s {
		value, ok := elem.(*WebDocumentNoProxy)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// WebDocumentArray is adapter for slice of WebDocument.
type WebDocumentArray []WebDocument

// Sort sorts slice of WebDocument.
func (s WebDocumentArray) Sort(less func(a, b WebDocument) bool) WebDocumentArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of WebDocument.
func (s WebDocumentArray) SortStable(less func(a, b WebDocument) bool) WebDocumentArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of WebDocument.
func (s WebDocumentArray) Retain(keep func(x WebDocument) bool) WebDocumentArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s WebDocumentArray) First() (v WebDocument, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s WebDocumentArray) Last() (v WebDocument, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *WebDocumentArray) PopFirst() (v WebDocument, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero WebDocument
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *WebDocumentArray) Pop() (v WebDocument, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// WebDocumentNoProxyArray is adapter for slice of WebDocumentNoProxy.
type WebDocumentNoProxyArray []WebDocumentNoProxy

// Sort sorts slice of WebDocumentNoProxy.
func (s WebDocumentNoProxyArray) Sort(less func(a, b WebDocumentNoProxy) bool) WebDocumentNoProxyArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of WebDocumentNoProxy.
func (s WebDocumentNoProxyArray) SortStable(less func(a, b WebDocumentNoProxy) bool) WebDocumentNoProxyArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of WebDocumentNoProxy.
func (s WebDocumentNoProxyArray) Retain(keep func(x WebDocumentNoProxy) bool) WebDocumentNoProxyArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s WebDocumentNoProxyArray) First() (v WebDocumentNoProxy, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s WebDocumentNoProxyArray) Last() (v WebDocumentNoProxy, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *WebDocumentNoProxyArray) PopFirst() (v WebDocumentNoProxy, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero WebDocumentNoProxy
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *WebDocumentNoProxyArray) Pop() (v WebDocumentNoProxy, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
