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

// DcOption represents TL type `dcOption#18b7a10d`.
// Data centre
//
// See https://core.telegram.org/constructor/dcOption for reference.
type DcOption struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether the specified IP is an IPv6 address
	Ipv6 bool
	// Whether this DC should only be used to download or upload files¹
	//
	// Links:
	//  1) https://core.telegram.org/api/files
	MediaOnly bool
	// Whether this DC only supports connection with transport obfuscation¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/mtproto-transports#transport-obfuscation
	TcpoOnly bool
	// Whether this is a CDN DC¹.
	//
	// Links:
	//  1) https://core.telegram.org/cdn
	CDN bool
	// If set, this IP should be used when connecting through a proxy
	Static bool
	// DC ID
	ID int
	// IP address of DC
	IPAddress string
	// Port
	Port int
	// If the tcpo_only flag is set, specifies the secret to use when connecting using transport obfuscation¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/mtproto-transports#transport-obfuscation
	//
	// Use SetSecret and GetSecret helpers.
	Secret []byte
}

// DcOptionTypeID is TL type id of DcOption.
const DcOptionTypeID = 0x18b7a10d

func (d *DcOption) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Flags.Zero()) {
		return false
	}
	if !(d.Ipv6 == false) {
		return false
	}
	if !(d.MediaOnly == false) {
		return false
	}
	if !(d.TcpoOnly == false) {
		return false
	}
	if !(d.CDN == false) {
		return false
	}
	if !(d.Static == false) {
		return false
	}
	if !(d.ID == 0) {
		return false
	}
	if !(d.IPAddress == "") {
		return false
	}
	if !(d.Port == 0) {
		return false
	}
	if !(d.Secret == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DcOption) String() string {
	if d == nil {
		return "DcOption(nil)"
	}
	type Alias DcOption
	return fmt.Sprintf("DcOption%+v", Alias(*d))
}

// FillFrom fills DcOption from given interface.
func (d *DcOption) FillFrom(from interface {
	GetIpv6() (value bool)
	GetMediaOnly() (value bool)
	GetTcpoOnly() (value bool)
	GetCDN() (value bool)
	GetStatic() (value bool)
	GetID() (value int)
	GetIPAddress() (value string)
	GetPort() (value int)
	GetSecret() (value []byte, ok bool)
}) {
	d.Ipv6 = from.GetIpv6()
	d.MediaOnly = from.GetMediaOnly()
	d.TcpoOnly = from.GetTcpoOnly()
	d.CDN = from.GetCDN()
	d.Static = from.GetStatic()
	d.ID = from.GetID()
	d.IPAddress = from.GetIPAddress()
	d.Port = from.GetPort()
	if val, ok := from.GetSecret(); ok {
		d.Secret = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DcOption) TypeID() uint32 {
	return DcOptionTypeID
}

// TypeName returns name of type in TL schema.
func (*DcOption) TypeName() string {
	return "dcOption"
}

// TypeInfo returns info about TL type.
func (d *DcOption) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "dcOption",
		ID:   DcOptionTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Flags",
			SchemaName: "flags",
		},
		{
			Name:       "Ipv6",
			SchemaName: "ipv6",
			Null:       !d.Flags.Has(0),
		},
		{
			Name:       "MediaOnly",
			SchemaName: "media_only",
			Null:       !d.Flags.Has(1),
		},
		{
			Name:       "TcpoOnly",
			SchemaName: "tcpo_only",
			Null:       !d.Flags.Has(2),
		},
		{
			Name:       "CDN",
			SchemaName: "cdn",
			Null:       !d.Flags.Has(3),
		},
		{
			Name:       "Static",
			SchemaName: "static",
			Null:       !d.Flags.Has(4),
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "IPAddress",
			SchemaName: "ip_address",
		},
		{
			Name:       "Port",
			SchemaName: "port",
		},
		{
			Name:       "Secret",
			SchemaName: "secret",
			Null:       !d.Flags.Has(10),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DcOption) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode dcOption#18b7a10d as nil")
	}
	b.PutID(DcOptionTypeID)
	if !(d.Ipv6 == false) {
		d.Flags.Set(0)
	}
	if !(d.MediaOnly == false) {
		d.Flags.Set(1)
	}
	if !(d.TcpoOnly == false) {
		d.Flags.Set(2)
	}
	if !(d.CDN == false) {
		d.Flags.Set(3)
	}
	if !(d.Static == false) {
		d.Flags.Set(4)
	}
	if !(d.Secret == nil) {
		d.Flags.Set(10)
	}
	if err := d.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode dcOption#18b7a10d: field flags: %w", err)
	}
	b.PutInt(d.ID)
	b.PutString(d.IPAddress)
	b.PutInt(d.Port)
	if d.Flags.Has(10) {
		b.PutBytes(d.Secret)
	}
	return nil
}

// SetIpv6 sets value of Ipv6 conditional field.
func (d *DcOption) SetIpv6(value bool) {
	if value {
		d.Flags.Set(0)
		d.Ipv6 = true
	} else {
		d.Flags.Unset(0)
		d.Ipv6 = false
	}
}

// GetIpv6 returns value of Ipv6 conditional field.
func (d *DcOption) GetIpv6() (value bool) {
	return d.Flags.Has(0)
}

// SetMediaOnly sets value of MediaOnly conditional field.
func (d *DcOption) SetMediaOnly(value bool) {
	if value {
		d.Flags.Set(1)
		d.MediaOnly = true
	} else {
		d.Flags.Unset(1)
		d.MediaOnly = false
	}
}

// GetMediaOnly returns value of MediaOnly conditional field.
func (d *DcOption) GetMediaOnly() (value bool) {
	return d.Flags.Has(1)
}

// SetTcpoOnly sets value of TcpoOnly conditional field.
func (d *DcOption) SetTcpoOnly(value bool) {
	if value {
		d.Flags.Set(2)
		d.TcpoOnly = true
	} else {
		d.Flags.Unset(2)
		d.TcpoOnly = false
	}
}

// GetTcpoOnly returns value of TcpoOnly conditional field.
func (d *DcOption) GetTcpoOnly() (value bool) {
	return d.Flags.Has(2)
}

// SetCDN sets value of CDN conditional field.
func (d *DcOption) SetCDN(value bool) {
	if value {
		d.Flags.Set(3)
		d.CDN = true
	} else {
		d.Flags.Unset(3)
		d.CDN = false
	}
}

// GetCDN returns value of CDN conditional field.
func (d *DcOption) GetCDN() (value bool) {
	return d.Flags.Has(3)
}

// SetStatic sets value of Static conditional field.
func (d *DcOption) SetStatic(value bool) {
	if value {
		d.Flags.Set(4)
		d.Static = true
	} else {
		d.Flags.Unset(4)
		d.Static = false
	}
}

// GetStatic returns value of Static conditional field.
func (d *DcOption) GetStatic() (value bool) {
	return d.Flags.Has(4)
}

// GetID returns value of ID field.
func (d *DcOption) GetID() (value int) {
	return d.ID
}

// GetIPAddress returns value of IPAddress field.
func (d *DcOption) GetIPAddress() (value string) {
	return d.IPAddress
}

// GetPort returns value of Port field.
func (d *DcOption) GetPort() (value int) {
	return d.Port
}

// SetSecret sets value of Secret conditional field.
func (d *DcOption) SetSecret(value []byte) {
	d.Flags.Set(10)
	d.Secret = value
}

// GetSecret returns value of Secret conditional field and
// boolean which is true if field was set.
func (d *DcOption) GetSecret() (value []byte, ok bool) {
	if !d.Flags.Has(10) {
		return value, false
	}
	return d.Secret, true
}

// Decode implements bin.Decoder.
func (d *DcOption) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode dcOption#18b7a10d to nil")
	}
	if err := b.ConsumeID(DcOptionTypeID); err != nil {
		return fmt.Errorf("unable to decode dcOption#18b7a10d: %w", err)
	}
	{
		if err := d.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode dcOption#18b7a10d: field flags: %w", err)
		}
	}
	d.Ipv6 = d.Flags.Has(0)
	d.MediaOnly = d.Flags.Has(1)
	d.TcpoOnly = d.Flags.Has(2)
	d.CDN = d.Flags.Has(3)
	d.Static = d.Flags.Has(4)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode dcOption#18b7a10d: field id: %w", err)
		}
		d.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode dcOption#18b7a10d: field ip_address: %w", err)
		}
		d.IPAddress = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode dcOption#18b7a10d: field port: %w", err)
		}
		d.Port = value
	}
	if d.Flags.Has(10) {
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode dcOption#18b7a10d: field secret: %w", err)
		}
		d.Secret = value
	}
	return nil
}

// Ensuring interfaces in compile-time for DcOption.
var (
	_ bin.Encoder = &DcOption{}
	_ bin.Decoder = &DcOption{}
)
