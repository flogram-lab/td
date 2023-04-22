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

// MessagesBotApp represents TL type `messages.botApp#eb50adf5`.
// Contains information about a named bot web app¹
//
// Links:
//  1. https://core.telegram.org/api/bots/webapps#named-bot-web-apps
//
// See https://core.telegram.org/constructor/messages.botApp for reference.
type MessagesBotApp struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether the web app was never used by the user, and confirmation must be asked from
	// the user before opening it.
	Inactive bool
	// The bot is asking permission to send messages to the user: if the user agrees, set the
	// write_allowed flag when invoking messages.requestAppWebView¹.
	//
	// Links:
	//  1) https://core.telegram.org/method/messages.requestAppWebView
	RequestWriteAccess bool
	// Bot app information
	App BotAppClass
}

// MessagesBotAppTypeID is TL type id of MessagesBotApp.
const MessagesBotAppTypeID = 0xeb50adf5

// Ensuring interfaces in compile-time for MessagesBotApp.
var (
	_ bin.Encoder     = &MessagesBotApp{}
	_ bin.Decoder     = &MessagesBotApp{}
	_ bin.BareEncoder = &MessagesBotApp{}
	_ bin.BareDecoder = &MessagesBotApp{}
)

func (b *MessagesBotApp) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.Flags.Zero()) {
		return false
	}
	if !(b.Inactive == false) {
		return false
	}
	if !(b.RequestWriteAccess == false) {
		return false
	}
	if !(b.App == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *MessagesBotApp) String() string {
	if b == nil {
		return "MessagesBotApp(nil)"
	}
	type Alias MessagesBotApp
	return fmt.Sprintf("MessagesBotApp%+v", Alias(*b))
}

// FillFrom fills MessagesBotApp from given interface.
func (b *MessagesBotApp) FillFrom(from interface {
	GetInactive() (value bool)
	GetRequestWriteAccess() (value bool)
	GetApp() (value BotAppClass)
}) {
	b.Inactive = from.GetInactive()
	b.RequestWriteAccess = from.GetRequestWriteAccess()
	b.App = from.GetApp()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesBotApp) TypeID() uint32 {
	return MessagesBotAppTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesBotApp) TypeName() string {
	return "messages.botApp"
}

// TypeInfo returns info about TL type.
func (b *MessagesBotApp) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.botApp",
		ID:   MessagesBotAppTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Inactive",
			SchemaName: "inactive",
			Null:       !b.Flags.Has(0),
		},
		{
			Name:       "RequestWriteAccess",
			SchemaName: "request_write_access",
			Null:       !b.Flags.Has(1),
		},
		{
			Name:       "App",
			SchemaName: "app",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (b *MessagesBotApp) SetFlags() {
	if !(b.Inactive == false) {
		b.Flags.Set(0)
	}
	if !(b.RequestWriteAccess == false) {
		b.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (b *MessagesBotApp) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode messages.botApp#eb50adf5 as nil")
	}
	buf.PutID(MessagesBotAppTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *MessagesBotApp) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode messages.botApp#eb50adf5 as nil")
	}
	b.SetFlags()
	if err := b.Flags.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode messages.botApp#eb50adf5: field flags: %w", err)
	}
	if b.App == nil {
		return fmt.Errorf("unable to encode messages.botApp#eb50adf5: field app is nil")
	}
	if err := b.App.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode messages.botApp#eb50adf5: field app: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *MessagesBotApp) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode messages.botApp#eb50adf5 to nil")
	}
	if err := buf.ConsumeID(MessagesBotAppTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.botApp#eb50adf5: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *MessagesBotApp) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode messages.botApp#eb50adf5 to nil")
	}
	{
		if err := b.Flags.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode messages.botApp#eb50adf5: field flags: %w", err)
		}
	}
	b.Inactive = b.Flags.Has(0)
	b.RequestWriteAccess = b.Flags.Has(1)
	{
		value, err := DecodeBotApp(buf)
		if err != nil {
			return fmt.Errorf("unable to decode messages.botApp#eb50adf5: field app: %w", err)
		}
		b.App = value
	}
	return nil
}

// SetInactive sets value of Inactive conditional field.
func (b *MessagesBotApp) SetInactive(value bool) {
	if value {
		b.Flags.Set(0)
		b.Inactive = true
	} else {
		b.Flags.Unset(0)
		b.Inactive = false
	}
}

// GetInactive returns value of Inactive conditional field.
func (b *MessagesBotApp) GetInactive() (value bool) {
	if b == nil {
		return
	}
	return b.Flags.Has(0)
}

// SetRequestWriteAccess sets value of RequestWriteAccess conditional field.
func (b *MessagesBotApp) SetRequestWriteAccess(value bool) {
	if value {
		b.Flags.Set(1)
		b.RequestWriteAccess = true
	} else {
		b.Flags.Unset(1)
		b.RequestWriteAccess = false
	}
}

// GetRequestWriteAccess returns value of RequestWriteAccess conditional field.
func (b *MessagesBotApp) GetRequestWriteAccess() (value bool) {
	if b == nil {
		return
	}
	return b.Flags.Has(1)
}

// GetApp returns value of App field.
func (b *MessagesBotApp) GetApp() (value BotAppClass) {
	if b == nil {
		return
	}
	return b.App
}

// GetAppAsModified returns mapped value of App field.
func (b *MessagesBotApp) GetAppAsModified() (*BotApp, bool) {
	return b.App.AsModified()
}
