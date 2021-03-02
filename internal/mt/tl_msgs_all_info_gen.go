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

// MsgsAllInfo represents TL type `msgs_all_info#8cc0d131`.
type MsgsAllInfo struct {
	// MsgIds field of MsgsAllInfo.
	MsgIds []int64
	// Info field of MsgsAllInfo.
	Info []byte
}

// MsgsAllInfoTypeID is TL type id of MsgsAllInfo.
const MsgsAllInfoTypeID = 0x8cc0d131

func (m *MsgsAllInfo) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.MsgIds == nil) {
		return false
	}
	if !(m.Info == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MsgsAllInfo) String() string {
	if m == nil {
		return "MsgsAllInfo(nil)"
	}
	type Alias MsgsAllInfo
	return fmt.Sprintf("MsgsAllInfo%+v", Alias(*m))
}

// FillFrom fills MsgsAllInfo from given interface.
func (m *MsgsAllInfo) FillFrom(from interface {
	GetMsgIds() (value []int64)
	GetInfo() (value []byte)
}) {
	m.MsgIds = from.GetMsgIds()
	m.Info = from.GetInfo()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MsgsAllInfo) TypeID() uint32 {
	return MsgsAllInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*MsgsAllInfo) TypeName() string {
	return "msgs_all_info"
}

// TypeInfo returns info about TL type.
func (m *MsgsAllInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "msgs_all_info",
		ID:   MsgsAllInfoTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "MsgIds",
			SchemaName: "msg_ids",
		},
		{
			Name:       "Info",
			SchemaName: "info",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MsgsAllInfo) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode msgs_all_info#8cc0d131 as nil")
	}
	b.PutID(MsgsAllInfoTypeID)
	b.PutVectorHeader(len(m.MsgIds))
	for _, v := range m.MsgIds {
		b.PutLong(v)
	}
	b.PutBytes(m.Info)
	return nil
}

// GetMsgIds returns value of MsgIds field.
func (m *MsgsAllInfo) GetMsgIds() (value []int64) {
	return m.MsgIds
}

// GetInfo returns value of Info field.
func (m *MsgsAllInfo) GetInfo() (value []byte) {
	return m.Info
}

// Decode implements bin.Decoder.
func (m *MsgsAllInfo) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode msgs_all_info#8cc0d131 to nil")
	}
	if err := b.ConsumeID(MsgsAllInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode msgs_all_info#8cc0d131: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode msgs_all_info#8cc0d131: field msg_ids: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode msgs_all_info#8cc0d131: field msg_ids: %w", err)
			}
			m.MsgIds = append(m.MsgIds, value)
		}
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode msgs_all_info#8cc0d131: field info: %w", err)
		}
		m.Info = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MsgsAllInfo.
var (
	_ bin.Encoder = &MsgsAllInfo{}
	_ bin.Decoder = &MsgsAllInfo{}
)
