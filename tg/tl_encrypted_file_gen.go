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

// EncryptedFileEmpty represents TL type `encryptedFileEmpty#c21f497e`.
// Empty constructor, unexisitng file.
//
// See https://core.telegram.org/constructor/encryptedFileEmpty for reference.
type EncryptedFileEmpty struct {
}

// EncryptedFileEmptyTypeID is TL type id of EncryptedFileEmpty.
const EncryptedFileEmptyTypeID = 0xc21f497e

func (e *EncryptedFileEmpty) Zero() bool {
	if e == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (e *EncryptedFileEmpty) String() string {
	if e == nil {
		return "EncryptedFileEmpty(nil)"
	}
	type Alias EncryptedFileEmpty
	return fmt.Sprintf("EncryptedFileEmpty%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EncryptedFileEmpty) TypeID() uint32 {
	return EncryptedFileEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*EncryptedFileEmpty) TypeName() string {
	return "encryptedFileEmpty"
}

// TypeInfo returns info about TL type.
func (e *EncryptedFileEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "encryptedFileEmpty",
		ID:   EncryptedFileEmptyTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (e *EncryptedFileEmpty) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedFileEmpty#c21f497e as nil")
	}
	b.PutID(EncryptedFileEmptyTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (e *EncryptedFileEmpty) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedFileEmpty#c21f497e to nil")
	}
	if err := b.ConsumeID(EncryptedFileEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode encryptedFileEmpty#c21f497e: %w", err)
	}
	return nil
}

// construct implements constructor of EncryptedFileClass.
func (e EncryptedFileEmpty) construct() EncryptedFileClass { return &e }

// Ensuring interfaces in compile-time for EncryptedFileEmpty.
var (
	_ bin.Encoder = &EncryptedFileEmpty{}
	_ bin.Decoder = &EncryptedFileEmpty{}

	_ EncryptedFileClass = &EncryptedFileEmpty{}
)

// EncryptedFile represents TL type `encryptedFile#4a70994c`.
// Encrypted file.
//
// See https://core.telegram.org/constructor/encryptedFile for reference.
type EncryptedFile struct {
	// File ID
	ID int64
	// Checking sum depending on user ID
	AccessHash int64
	// File size in bytes
	Size int
	// Number of data centre
	DCID int
	// 32-bit fingerprint of key used for file encryption
	KeyFingerprint int
}

// EncryptedFileTypeID is TL type id of EncryptedFile.
const EncryptedFileTypeID = 0x4a70994c

func (e *EncryptedFile) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.ID == 0) {
		return false
	}
	if !(e.AccessHash == 0) {
		return false
	}
	if !(e.Size == 0) {
		return false
	}
	if !(e.DCID == 0) {
		return false
	}
	if !(e.KeyFingerprint == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EncryptedFile) String() string {
	if e == nil {
		return "EncryptedFile(nil)"
	}
	type Alias EncryptedFile
	return fmt.Sprintf("EncryptedFile%+v", Alias(*e))
}

// FillFrom fills EncryptedFile from given interface.
func (e *EncryptedFile) FillFrom(from interface {
	GetID() (value int64)
	GetAccessHash() (value int64)
	GetSize() (value int)
	GetDCID() (value int)
	GetKeyFingerprint() (value int)
}) {
	e.ID = from.GetID()
	e.AccessHash = from.GetAccessHash()
	e.Size = from.GetSize()
	e.DCID = from.GetDCID()
	e.KeyFingerprint = from.GetKeyFingerprint()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EncryptedFile) TypeID() uint32 {
	return EncryptedFileTypeID
}

// TypeName returns name of type in TL schema.
func (*EncryptedFile) TypeName() string {
	return "encryptedFile"
}

// TypeInfo returns info about TL type.
func (e *EncryptedFile) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "encryptedFile",
		ID:   EncryptedFileTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
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
			Name:       "DCID",
			SchemaName: "dc_id",
		},
		{
			Name:       "KeyFingerprint",
			SchemaName: "key_fingerprint",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EncryptedFile) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedFile#4a70994c as nil")
	}
	b.PutID(EncryptedFileTypeID)
	b.PutLong(e.ID)
	b.PutLong(e.AccessHash)
	b.PutInt(e.Size)
	b.PutInt(e.DCID)
	b.PutInt(e.KeyFingerprint)
	return nil
}

// GetID returns value of ID field.
func (e *EncryptedFile) GetID() (value int64) {
	return e.ID
}

// GetAccessHash returns value of AccessHash field.
func (e *EncryptedFile) GetAccessHash() (value int64) {
	return e.AccessHash
}

// GetSize returns value of Size field.
func (e *EncryptedFile) GetSize() (value int) {
	return e.Size
}

// GetDCID returns value of DCID field.
func (e *EncryptedFile) GetDCID() (value int) {
	return e.DCID
}

// GetKeyFingerprint returns value of KeyFingerprint field.
func (e *EncryptedFile) GetKeyFingerprint() (value int) {
	return e.KeyFingerprint
}

// Decode implements bin.Decoder.
func (e *EncryptedFile) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedFile#4a70994c to nil")
	}
	if err := b.ConsumeID(EncryptedFileTypeID); err != nil {
		return fmt.Errorf("unable to decode encryptedFile#4a70994c: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedFile#4a70994c: field id: %w", err)
		}
		e.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedFile#4a70994c: field access_hash: %w", err)
		}
		e.AccessHash = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedFile#4a70994c: field size: %w", err)
		}
		e.Size = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedFile#4a70994c: field dc_id: %w", err)
		}
		e.DCID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedFile#4a70994c: field key_fingerprint: %w", err)
		}
		e.KeyFingerprint = value
	}
	return nil
}

// construct implements constructor of EncryptedFileClass.
func (e EncryptedFile) construct() EncryptedFileClass { return &e }

// Ensuring interfaces in compile-time for EncryptedFile.
var (
	_ bin.Encoder = &EncryptedFile{}
	_ bin.Decoder = &EncryptedFile{}

	_ EncryptedFileClass = &EncryptedFile{}
)

// EncryptedFileClass represents EncryptedFile generic type.
//
// See https://core.telegram.org/type/EncryptedFile for reference.
//
// Example:
//  g, err := tg.DecodeEncryptedFile(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.EncryptedFileEmpty: // encryptedFileEmpty#c21f497e
//  case *tg.EncryptedFile: // encryptedFile#4a70994c
//  default: panic(v)
//  }
type EncryptedFileClass interface {
	bin.Encoder
	bin.Decoder
	construct() EncryptedFileClass

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

	// AsNotEmpty tries to map EncryptedFileClass to EncryptedFile.
	AsNotEmpty() (*EncryptedFile, bool)
}

// AsInputEncryptedFileLocation tries to map EncryptedFile to InputEncryptedFileLocation.
func (e *EncryptedFile) AsInputEncryptedFileLocation() *InputEncryptedFileLocation {
	value := new(InputEncryptedFileLocation)
	value.ID = e.GetID()
	value.AccessHash = e.GetAccessHash()

	return value
}

// AsInput tries to map EncryptedFile to InputEncryptedFile.
func (e *EncryptedFile) AsInput() *InputEncryptedFile {
	value := new(InputEncryptedFile)
	value.ID = e.GetID()
	value.AccessHash = e.GetAccessHash()

	return value
}

// AsNotEmpty tries to map EncryptedFileEmpty to EncryptedFile.
func (e *EncryptedFileEmpty) AsNotEmpty() (*EncryptedFile, bool) {
	return nil, false
}

// AsNotEmpty tries to map EncryptedFile to EncryptedFile.
func (e *EncryptedFile) AsNotEmpty() (*EncryptedFile, bool) {
	return e, true
}

// DecodeEncryptedFile implements binary de-serialization for EncryptedFileClass.
func DecodeEncryptedFile(buf *bin.Buffer) (EncryptedFileClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case EncryptedFileEmptyTypeID:
		// Decoding encryptedFileEmpty#c21f497e.
		v := EncryptedFileEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EncryptedFileClass: %w", err)
		}
		return &v, nil
	case EncryptedFileTypeID:
		// Decoding encryptedFile#4a70994c.
		v := EncryptedFile{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EncryptedFileClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode EncryptedFileClass: %w", bin.NewUnexpectedID(id))
	}
}

// EncryptedFile boxes the EncryptedFileClass providing a helper.
type EncryptedFileBox struct {
	EncryptedFile EncryptedFileClass
}

// Decode implements bin.Decoder for EncryptedFileBox.
func (b *EncryptedFileBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode EncryptedFileBox to nil")
	}
	v, err := DecodeEncryptedFile(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.EncryptedFile = v
	return nil
}

// Encode implements bin.Encode for EncryptedFileBox.
func (b *EncryptedFileBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.EncryptedFile == nil {
		return fmt.Errorf("unable to encode EncryptedFileClass as nil")
	}
	return b.EncryptedFile.Encode(buf)
}

// EncryptedFileClassArray is adapter for slice of EncryptedFileClass.
type EncryptedFileClassArray []EncryptedFileClass

// Sort sorts slice of EncryptedFileClass.
func (s EncryptedFileClassArray) Sort(less func(a, b EncryptedFileClass) bool) EncryptedFileClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of EncryptedFileClass.
func (s EncryptedFileClassArray) SortStable(less func(a, b EncryptedFileClass) bool) EncryptedFileClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of EncryptedFileClass.
func (s EncryptedFileClassArray) Retain(keep func(x EncryptedFileClass) bool) EncryptedFileClassArray {
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
func (s EncryptedFileClassArray) First() (v EncryptedFileClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s EncryptedFileClassArray) Last() (v EncryptedFileClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *EncryptedFileClassArray) PopFirst() (v EncryptedFileClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero EncryptedFileClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *EncryptedFileClassArray) Pop() (v EncryptedFileClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsEncryptedFile returns copy with only EncryptedFile constructors.
func (s EncryptedFileClassArray) AsEncryptedFile() (to EncryptedFileArray) {
	for _, elem := range s {
		value, ok := elem.(*EncryptedFile)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyNotEmpty appends only NotEmpty constructors to
// given slice.
func (s EncryptedFileClassArray) AppendOnlyNotEmpty(to []*EncryptedFile) []*EncryptedFile {
	for _, elem := range s {
		value, ok := elem.AsNotEmpty()
		if !ok {
			continue
		}
		to = append(to, value)
	}

	return to
}

// AsNotEmpty returns copy with only NotEmpty constructors.
func (s EncryptedFileClassArray) AsNotEmpty() (to []*EncryptedFile) {
	return s.AppendOnlyNotEmpty(to)
}

// FirstAsNotEmpty returns first element of slice (if exists).
func (s EncryptedFileClassArray) FirstAsNotEmpty() (v *EncryptedFile, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// LastAsNotEmpty returns last element of slice (if exists).
func (s EncryptedFileClassArray) LastAsNotEmpty() (v *EncryptedFile, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopFirstAsNotEmpty returns element of slice (if exists).
func (s *EncryptedFileClassArray) PopFirstAsNotEmpty() (v *EncryptedFile, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopAsNotEmpty returns element of slice (if exists).
func (s *EncryptedFileClassArray) PopAsNotEmpty() (v *EncryptedFile, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// EncryptedFileArray is adapter for slice of EncryptedFile.
type EncryptedFileArray []EncryptedFile

// Sort sorts slice of EncryptedFile.
func (s EncryptedFileArray) Sort(less func(a, b EncryptedFile) bool) EncryptedFileArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of EncryptedFile.
func (s EncryptedFileArray) SortStable(less func(a, b EncryptedFile) bool) EncryptedFileArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of EncryptedFile.
func (s EncryptedFileArray) Retain(keep func(x EncryptedFile) bool) EncryptedFileArray {
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
func (s EncryptedFileArray) First() (v EncryptedFile, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s EncryptedFileArray) Last() (v EncryptedFile, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *EncryptedFileArray) PopFirst() (v EncryptedFile, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero EncryptedFile
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *EncryptedFileArray) Pop() (v EncryptedFile, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
