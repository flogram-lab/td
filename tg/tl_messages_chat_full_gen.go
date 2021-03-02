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

// MessagesChatFull represents TL type `messages.chatFull#e5d7d19c`.
// Extended info on chat and auxiliary data.
//
// See https://core.telegram.org/constructor/messages.chatFull for reference.
type MessagesChatFull struct {
	// Extended info on a chat
	FullChat ChatFullClass
	// List containing basic info on chat
	Chats []ChatClass
	// List of users mentioned above
	Users []UserClass
}

// MessagesChatFullTypeID is TL type id of MessagesChatFull.
const MessagesChatFullTypeID = 0xe5d7d19c

func (c *MessagesChatFull) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.FullChat == nil) {
		return false
	}
	if !(c.Chats == nil) {
		return false
	}
	if !(c.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *MessagesChatFull) String() string {
	if c == nil {
		return "MessagesChatFull(nil)"
	}
	type Alias MessagesChatFull
	return fmt.Sprintf("MessagesChatFull%+v", Alias(*c))
}

// FillFrom fills MessagesChatFull from given interface.
func (c *MessagesChatFull) FillFrom(from interface {
	GetFullChat() (value ChatFullClass)
	GetChats() (value []ChatClass)
	GetUsers() (value []UserClass)
}) {
	c.FullChat = from.GetFullChat()
	c.Chats = from.GetChats()
	c.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesChatFull) TypeID() uint32 {
	return MessagesChatFullTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesChatFull) TypeName() string {
	return "messages.chatFull"
}

// TypeInfo returns info about TL type.
func (c *MessagesChatFull) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.chatFull",
		ID:   MessagesChatFullTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "FullChat",
			SchemaName: "full_chat",
		},
		{
			Name:       "Chats",
			SchemaName: "chats",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *MessagesChatFull) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode messages.chatFull#e5d7d19c as nil")
	}
	b.PutID(MessagesChatFullTypeID)
	if c.FullChat == nil {
		return fmt.Errorf("unable to encode messages.chatFull#e5d7d19c: field full_chat is nil")
	}
	if err := c.FullChat.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.chatFull#e5d7d19c: field full_chat: %w", err)
	}
	b.PutVectorHeader(len(c.Chats))
	for idx, v := range c.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode messages.chatFull#e5d7d19c: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.chatFull#e5d7d19c: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(c.Users))
	for idx, v := range c.Users {
		if v == nil {
			return fmt.Errorf("unable to encode messages.chatFull#e5d7d19c: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.chatFull#e5d7d19c: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetFullChat returns value of FullChat field.
func (c *MessagesChatFull) GetFullChat() (value ChatFullClass) {
	return c.FullChat
}

// GetChats returns value of Chats field.
func (c *MessagesChatFull) GetChats() (value []ChatClass) {
	return c.Chats
}

// MapChats returns field Chats wrapped in ChatClassArray helper.
func (c *MessagesChatFull) MapChats() (value ChatClassArray) {
	return ChatClassArray(c.Chats)
}

// GetUsers returns value of Users field.
func (c *MessagesChatFull) GetUsers() (value []UserClass) {
	return c.Users
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (c *MessagesChatFull) MapUsers() (value UserClassArray) {
	return UserClassArray(c.Users)
}

// Decode implements bin.Decoder.
func (c *MessagesChatFull) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode messages.chatFull#e5d7d19c to nil")
	}
	if err := b.ConsumeID(MessagesChatFullTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.chatFull#e5d7d19c: %w", err)
	}
	{
		value, err := DecodeChatFull(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.chatFull#e5d7d19c: field full_chat: %w", err)
		}
		c.FullChat = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.chatFull#e5d7d19c: field chats: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.chatFull#e5d7d19c: field chats: %w", err)
			}
			c.Chats = append(c.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.chatFull#e5d7d19c: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.chatFull#e5d7d19c: field users: %w", err)
			}
			c.Users = append(c.Users, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesChatFull.
var (
	_ bin.Encoder = &MessagesChatFull{}
	_ bin.Decoder = &MessagesChatFull{}
)
