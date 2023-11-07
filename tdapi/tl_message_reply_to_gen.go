// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// MessageReplyToMessage represents TL type `messageReplyToMessage#1bb65082`.
type MessageReplyToMessage struct {
	// The identifier of the chat to which the message belongs; may be 0 if the replied
	// message is in unknown chat
	ChatID int64
	// The identifier of the message; may be 0 if the replied message is in unknown chat
	MessageID int64
	// Manually or automatically chosen quote from the replied message; may be null if none.
	// Only Bold, Italic, Underline, Strikethrough, Spoiler, and CustomEmoji entities can be
	// present in the quote
	Quote FormattedText
	// True, if the quote was manually chosen by the message sender
	IsQuoteManual bool
	// Information about origin of the message if the message was replied from another chat;
	// may be null for messages from the same chat
	Origin MessageOriginClass
	// Point in time (Unix timestamp) when the message was sent if the message was replied
	// from another chat; 0 for messages from the same chat
	OriginSendDate int32
	// Media content of the message if the message was replied from another chat; may be null
	// for messages from the same chat and messages without media.
	Content MessageContentClass
}

// MessageReplyToMessageTypeID is TL type id of MessageReplyToMessage.
const MessageReplyToMessageTypeID = 0x1bb65082

// construct implements constructor of MessageReplyToClass.
func (m MessageReplyToMessage) construct() MessageReplyToClass { return &m }

// Ensuring interfaces in compile-time for MessageReplyToMessage.
var (
	_ bin.Encoder     = &MessageReplyToMessage{}
	_ bin.Decoder     = &MessageReplyToMessage{}
	_ bin.BareEncoder = &MessageReplyToMessage{}
	_ bin.BareDecoder = &MessageReplyToMessage{}

	_ MessageReplyToClass = &MessageReplyToMessage{}
)

func (m *MessageReplyToMessage) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.ChatID == 0) {
		return false
	}
	if !(m.MessageID == 0) {
		return false
	}
	if !(m.Quote.Zero()) {
		return false
	}
	if !(m.IsQuoteManual == false) {
		return false
	}
	if !(m.Origin == nil) {
		return false
	}
	if !(m.OriginSendDate == 0) {
		return false
	}
	if !(m.Content == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageReplyToMessage) String() string {
	if m == nil {
		return "MessageReplyToMessage(nil)"
	}
	type Alias MessageReplyToMessage
	return fmt.Sprintf("MessageReplyToMessage%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageReplyToMessage) TypeID() uint32 {
	return MessageReplyToMessageTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageReplyToMessage) TypeName() string {
	return "messageReplyToMessage"
}

// TypeInfo returns info about TL type.
func (m *MessageReplyToMessage) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageReplyToMessage",
		ID:   MessageReplyToMessageTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageID",
			SchemaName: "message_id",
		},
		{
			Name:       "Quote",
			SchemaName: "quote",
		},
		{
			Name:       "IsQuoteManual",
			SchemaName: "is_quote_manual",
		},
		{
			Name:       "Origin",
			SchemaName: "origin",
		},
		{
			Name:       "OriginSendDate",
			SchemaName: "origin_send_date",
		},
		{
			Name:       "Content",
			SchemaName: "content",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessageReplyToMessage) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageReplyToMessage#1bb65082 as nil")
	}
	b.PutID(MessageReplyToMessageTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageReplyToMessage) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageReplyToMessage#1bb65082 as nil")
	}
	b.PutInt53(m.ChatID)
	b.PutInt53(m.MessageID)
	if err := m.Quote.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messageReplyToMessage#1bb65082: field quote: %w", err)
	}
	b.PutBool(m.IsQuoteManual)
	if m.Origin == nil {
		return fmt.Errorf("unable to encode messageReplyToMessage#1bb65082: field origin is nil")
	}
	if err := m.Origin.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messageReplyToMessage#1bb65082: field origin: %w", err)
	}
	b.PutInt32(m.OriginSendDate)
	if m.Content == nil {
		return fmt.Errorf("unable to encode messageReplyToMessage#1bb65082: field content is nil")
	}
	if err := m.Content.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messageReplyToMessage#1bb65082: field content: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageReplyToMessage) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageReplyToMessage#1bb65082 to nil")
	}
	if err := b.ConsumeID(MessageReplyToMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode messageReplyToMessage#1bb65082: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageReplyToMessage) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageReplyToMessage#1bb65082 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode messageReplyToMessage#1bb65082: field chat_id: %w", err)
		}
		m.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode messageReplyToMessage#1bb65082: field message_id: %w", err)
		}
		m.MessageID = value
	}
	{
		if err := m.Quote.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messageReplyToMessage#1bb65082: field quote: %w", err)
		}
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode messageReplyToMessage#1bb65082: field is_quote_manual: %w", err)
		}
		m.IsQuoteManual = value
	}
	{
		value, err := DecodeMessageOrigin(b)
		if err != nil {
			return fmt.Errorf("unable to decode messageReplyToMessage#1bb65082: field origin: %w", err)
		}
		m.Origin = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode messageReplyToMessage#1bb65082: field origin_send_date: %w", err)
		}
		m.OriginSendDate = value
	}
	{
		value, err := DecodeMessageContent(b)
		if err != nil {
			return fmt.Errorf("unable to decode messageReplyToMessage#1bb65082: field content: %w", err)
		}
		m.Content = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (m *MessageReplyToMessage) EncodeTDLibJSON(b tdjson.Encoder) error {
	if m == nil {
		return fmt.Errorf("can't encode messageReplyToMessage#1bb65082 as nil")
	}
	b.ObjStart()
	b.PutID("messageReplyToMessage")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(m.ChatID)
	b.Comma()
	b.FieldStart("message_id")
	b.PutInt53(m.MessageID)
	b.Comma()
	b.FieldStart("quote")
	if err := m.Quote.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode messageReplyToMessage#1bb65082: field quote: %w", err)
	}
	b.Comma()
	b.FieldStart("is_quote_manual")
	b.PutBool(m.IsQuoteManual)
	b.Comma()
	b.FieldStart("origin")
	if m.Origin == nil {
		return fmt.Errorf("unable to encode messageReplyToMessage#1bb65082: field origin is nil")
	}
	if err := m.Origin.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode messageReplyToMessage#1bb65082: field origin: %w", err)
	}
	b.Comma()
	b.FieldStart("origin_send_date")
	b.PutInt32(m.OriginSendDate)
	b.Comma()
	b.FieldStart("content")
	if m.Content == nil {
		return fmt.Errorf("unable to encode messageReplyToMessage#1bb65082: field content is nil")
	}
	if err := m.Content.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode messageReplyToMessage#1bb65082: field content: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (m *MessageReplyToMessage) DecodeTDLibJSON(b tdjson.Decoder) error {
	if m == nil {
		return fmt.Errorf("can't decode messageReplyToMessage#1bb65082 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("messageReplyToMessage"); err != nil {
				return fmt.Errorf("unable to decode messageReplyToMessage#1bb65082: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode messageReplyToMessage#1bb65082: field chat_id: %w", err)
			}
			m.ChatID = value
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode messageReplyToMessage#1bb65082: field message_id: %w", err)
			}
			m.MessageID = value
		case "quote":
			if err := m.Quote.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode messageReplyToMessage#1bb65082: field quote: %w", err)
			}
		case "is_quote_manual":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode messageReplyToMessage#1bb65082: field is_quote_manual: %w", err)
			}
			m.IsQuoteManual = value
		case "origin":
			value, err := DecodeTDLibJSONMessageOrigin(b)
			if err != nil {
				return fmt.Errorf("unable to decode messageReplyToMessage#1bb65082: field origin: %w", err)
			}
			m.Origin = value
		case "origin_send_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode messageReplyToMessage#1bb65082: field origin_send_date: %w", err)
			}
			m.OriginSendDate = value
		case "content":
			value, err := DecodeTDLibJSONMessageContent(b)
			if err != nil {
				return fmt.Errorf("unable to decode messageReplyToMessage#1bb65082: field content: %w", err)
			}
			m.Content = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (m *MessageReplyToMessage) GetChatID() (value int64) {
	if m == nil {
		return
	}
	return m.ChatID
}

// GetMessageID returns value of MessageID field.
func (m *MessageReplyToMessage) GetMessageID() (value int64) {
	if m == nil {
		return
	}
	return m.MessageID
}

// GetQuote returns value of Quote field.
func (m *MessageReplyToMessage) GetQuote() (value FormattedText) {
	if m == nil {
		return
	}
	return m.Quote
}

// GetIsQuoteManual returns value of IsQuoteManual field.
func (m *MessageReplyToMessage) GetIsQuoteManual() (value bool) {
	if m == nil {
		return
	}
	return m.IsQuoteManual
}

// GetOrigin returns value of Origin field.
func (m *MessageReplyToMessage) GetOrigin() (value MessageOriginClass) {
	if m == nil {
		return
	}
	return m.Origin
}

// GetOriginSendDate returns value of OriginSendDate field.
func (m *MessageReplyToMessage) GetOriginSendDate() (value int32) {
	if m == nil {
		return
	}
	return m.OriginSendDate
}

// GetContent returns value of Content field.
func (m *MessageReplyToMessage) GetContent() (value MessageContentClass) {
	if m == nil {
		return
	}
	return m.Content
}

// MessageReplyToStory represents TL type `messageReplyToStory#708ca939`.
type MessageReplyToStory struct {
	// The identifier of the sender of the story
	StorySenderChatID int64
	// The identifier of the story
	StoryID int32
}

// MessageReplyToStoryTypeID is TL type id of MessageReplyToStory.
const MessageReplyToStoryTypeID = 0x708ca939

// construct implements constructor of MessageReplyToClass.
func (m MessageReplyToStory) construct() MessageReplyToClass { return &m }

// Ensuring interfaces in compile-time for MessageReplyToStory.
var (
	_ bin.Encoder     = &MessageReplyToStory{}
	_ bin.Decoder     = &MessageReplyToStory{}
	_ bin.BareEncoder = &MessageReplyToStory{}
	_ bin.BareDecoder = &MessageReplyToStory{}

	_ MessageReplyToClass = &MessageReplyToStory{}
)

func (m *MessageReplyToStory) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.StorySenderChatID == 0) {
		return false
	}
	if !(m.StoryID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageReplyToStory) String() string {
	if m == nil {
		return "MessageReplyToStory(nil)"
	}
	type Alias MessageReplyToStory
	return fmt.Sprintf("MessageReplyToStory%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageReplyToStory) TypeID() uint32 {
	return MessageReplyToStoryTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageReplyToStory) TypeName() string {
	return "messageReplyToStory"
}

// TypeInfo returns info about TL type.
func (m *MessageReplyToStory) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageReplyToStory",
		ID:   MessageReplyToStoryTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "StorySenderChatID",
			SchemaName: "story_sender_chat_id",
		},
		{
			Name:       "StoryID",
			SchemaName: "story_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessageReplyToStory) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageReplyToStory#708ca939 as nil")
	}
	b.PutID(MessageReplyToStoryTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageReplyToStory) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageReplyToStory#708ca939 as nil")
	}
	b.PutInt53(m.StorySenderChatID)
	b.PutInt32(m.StoryID)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageReplyToStory) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageReplyToStory#708ca939 to nil")
	}
	if err := b.ConsumeID(MessageReplyToStoryTypeID); err != nil {
		return fmt.Errorf("unable to decode messageReplyToStory#708ca939: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageReplyToStory) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageReplyToStory#708ca939 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode messageReplyToStory#708ca939: field story_sender_chat_id: %w", err)
		}
		m.StorySenderChatID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode messageReplyToStory#708ca939: field story_id: %w", err)
		}
		m.StoryID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (m *MessageReplyToStory) EncodeTDLibJSON(b tdjson.Encoder) error {
	if m == nil {
		return fmt.Errorf("can't encode messageReplyToStory#708ca939 as nil")
	}
	b.ObjStart()
	b.PutID("messageReplyToStory")
	b.Comma()
	b.FieldStart("story_sender_chat_id")
	b.PutInt53(m.StorySenderChatID)
	b.Comma()
	b.FieldStart("story_id")
	b.PutInt32(m.StoryID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (m *MessageReplyToStory) DecodeTDLibJSON(b tdjson.Decoder) error {
	if m == nil {
		return fmt.Errorf("can't decode messageReplyToStory#708ca939 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("messageReplyToStory"); err != nil {
				return fmt.Errorf("unable to decode messageReplyToStory#708ca939: %w", err)
			}
		case "story_sender_chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode messageReplyToStory#708ca939: field story_sender_chat_id: %w", err)
			}
			m.StorySenderChatID = value
		case "story_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode messageReplyToStory#708ca939: field story_id: %w", err)
			}
			m.StoryID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetStorySenderChatID returns value of StorySenderChatID field.
func (m *MessageReplyToStory) GetStorySenderChatID() (value int64) {
	if m == nil {
		return
	}
	return m.StorySenderChatID
}

// GetStoryID returns value of StoryID field.
func (m *MessageReplyToStory) GetStoryID() (value int32) {
	if m == nil {
		return
	}
	return m.StoryID
}

// MessageReplyToClassName is schema name of MessageReplyToClass.
const MessageReplyToClassName = "MessageReplyTo"

// MessageReplyToClass represents MessageReplyTo generic type.
//
// Example:
//
//	g, err := tdapi.DecodeMessageReplyTo(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.MessageReplyToMessage: // messageReplyToMessage#1bb65082
//	case *tdapi.MessageReplyToStory: // messageReplyToStory#708ca939
//	default: panic(v)
//	}
type MessageReplyToClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() MessageReplyToClass

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

	EncodeTDLibJSON(b tdjson.Encoder) error
	DecodeTDLibJSON(b tdjson.Decoder) error
}

// DecodeMessageReplyTo implements binary de-serialization for MessageReplyToClass.
func DecodeMessageReplyTo(buf *bin.Buffer) (MessageReplyToClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessageReplyToMessageTypeID:
		// Decoding messageReplyToMessage#1bb65082.
		v := MessageReplyToMessage{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageReplyToClass: %w", err)
		}
		return &v, nil
	case MessageReplyToStoryTypeID:
		// Decoding messageReplyToStory#708ca939.
		v := MessageReplyToStory{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageReplyToClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessageReplyToClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONMessageReplyTo implements binary de-serialization for MessageReplyToClass.
func DecodeTDLibJSONMessageReplyTo(buf tdjson.Decoder) (MessageReplyToClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "messageReplyToMessage":
		// Decoding messageReplyToMessage#1bb65082.
		v := MessageReplyToMessage{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageReplyToClass: %w", err)
		}
		return &v, nil
	case "messageReplyToStory":
		// Decoding messageReplyToStory#708ca939.
		v := MessageReplyToStory{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageReplyToClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessageReplyToClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// MessageReplyTo boxes the MessageReplyToClass providing a helper.
type MessageReplyToBox struct {
	MessageReplyTo MessageReplyToClass
}

// Decode implements bin.Decoder for MessageReplyToBox.
func (b *MessageReplyToBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessageReplyToBox to nil")
	}
	v, err := DecodeMessageReplyTo(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.MessageReplyTo = v
	return nil
}

// Encode implements bin.Encode for MessageReplyToBox.
func (b *MessageReplyToBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.MessageReplyTo == nil {
		return fmt.Errorf("unable to encode MessageReplyToClass as nil")
	}
	return b.MessageReplyTo.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for MessageReplyToBox.
func (b *MessageReplyToBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessageReplyToBox to nil")
	}
	v, err := DecodeTDLibJSONMessageReplyTo(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.MessageReplyTo = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for MessageReplyToBox.
func (b *MessageReplyToBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.MessageReplyTo == nil {
		return fmt.Errorf("unable to encode MessageReplyToClass as nil")
	}
	return b.MessageReplyTo.EncodeTDLibJSON(buf)
}
