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

// HelpCountry represents TL type `help.country#c3878e23`.
// Name, ISO code, localized name and phone codes/patterns of a specific country
//
// See https://core.telegram.org/constructor/help.country for reference.
type HelpCountry struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether this country should not be shown in the list
	Hidden bool
	// ISO code of country
	Iso2 string
	// Name of the country in the country's language
	DefaultName string
	// Name of the country in the user's language, if different from the original name
	//
	// Use SetName and GetName helpers.
	Name string
	// Phone codes/patterns
	CountryCodes []HelpCountryCode
}

// HelpCountryTypeID is TL type id of HelpCountry.
const HelpCountryTypeID = 0xc3878e23

func (c *HelpCountry) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.Hidden == false) {
		return false
	}
	if !(c.Iso2 == "") {
		return false
	}
	if !(c.DefaultName == "") {
		return false
	}
	if !(c.Name == "") {
		return false
	}
	if !(c.CountryCodes == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *HelpCountry) String() string {
	if c == nil {
		return "HelpCountry(nil)"
	}
	type Alias HelpCountry
	return fmt.Sprintf("HelpCountry%+v", Alias(*c))
}

// FillFrom fills HelpCountry from given interface.
func (c *HelpCountry) FillFrom(from interface {
	GetHidden() (value bool)
	GetIso2() (value string)
	GetDefaultName() (value string)
	GetName() (value string, ok bool)
	GetCountryCodes() (value []HelpCountryCode)
}) {
	c.Hidden = from.GetHidden()
	c.Iso2 = from.GetIso2()
	c.DefaultName = from.GetDefaultName()
	if val, ok := from.GetName(); ok {
		c.Name = val
	}

	c.CountryCodes = from.GetCountryCodes()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpCountry) TypeID() uint32 {
	return HelpCountryTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpCountry) TypeName() string {
	return "help.country"
}

// TypeInfo returns info about TL type.
func (c *HelpCountry) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.country",
		ID:   HelpCountryTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Flags",
			SchemaName: "flags",
		},
		{
			Name:       "Hidden",
			SchemaName: "hidden",
			Null:       !c.Flags.Has(0),
		},
		{
			Name:       "Iso2",
			SchemaName: "iso2",
		},
		{
			Name:       "DefaultName",
			SchemaName: "default_name",
		},
		{
			Name:       "Name",
			SchemaName: "name",
			Null:       !c.Flags.Has(1),
		},
		{
			Name:       "CountryCodes",
			SchemaName: "country_codes",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *HelpCountry) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode help.country#c3878e23 as nil")
	}
	b.PutID(HelpCountryTypeID)
	if !(c.Hidden == false) {
		c.Flags.Set(0)
	}
	if !(c.Name == "") {
		c.Flags.Set(1)
	}
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode help.country#c3878e23: field flags: %w", err)
	}
	b.PutString(c.Iso2)
	b.PutString(c.DefaultName)
	if c.Flags.Has(1) {
		b.PutString(c.Name)
	}
	b.PutVectorHeader(len(c.CountryCodes))
	for idx, v := range c.CountryCodes {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode help.country#c3878e23: field country_codes element with index %d: %w", idx, err)
		}
	}
	return nil
}

// SetHidden sets value of Hidden conditional field.
func (c *HelpCountry) SetHidden(value bool) {
	if value {
		c.Flags.Set(0)
		c.Hidden = true
	} else {
		c.Flags.Unset(0)
		c.Hidden = false
	}
}

// GetHidden returns value of Hidden conditional field.
func (c *HelpCountry) GetHidden() (value bool) {
	return c.Flags.Has(0)
}

// GetIso2 returns value of Iso2 field.
func (c *HelpCountry) GetIso2() (value string) {
	return c.Iso2
}

// GetDefaultName returns value of DefaultName field.
func (c *HelpCountry) GetDefaultName() (value string) {
	return c.DefaultName
}

// SetName sets value of Name conditional field.
func (c *HelpCountry) SetName(value string) {
	c.Flags.Set(1)
	c.Name = value
}

// GetName returns value of Name conditional field and
// boolean which is true if field was set.
func (c *HelpCountry) GetName() (value string, ok bool) {
	if !c.Flags.Has(1) {
		return value, false
	}
	return c.Name, true
}

// GetCountryCodes returns value of CountryCodes field.
func (c *HelpCountry) GetCountryCodes() (value []HelpCountryCode) {
	return c.CountryCodes
}

// Decode implements bin.Decoder.
func (c *HelpCountry) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode help.country#c3878e23 to nil")
	}
	if err := b.ConsumeID(HelpCountryTypeID); err != nil {
		return fmt.Errorf("unable to decode help.country#c3878e23: %w", err)
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode help.country#c3878e23: field flags: %w", err)
		}
	}
	c.Hidden = c.Flags.Has(0)
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.country#c3878e23: field iso2: %w", err)
		}
		c.Iso2 = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.country#c3878e23: field default_name: %w", err)
		}
		c.DefaultName = value
	}
	if c.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.country#c3878e23: field name: %w", err)
		}
		c.Name = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode help.country#c3878e23: field country_codes: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value HelpCountryCode
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode help.country#c3878e23: field country_codes: %w", err)
			}
			c.CountryCodes = append(c.CountryCodes, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpCountry.
var (
	_ bin.Encoder = &HelpCountry{}
	_ bin.Decoder = &HelpCountry{}
)
