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

// SendMessageAlbumRequest represents TL type `sendMessageAlbum#74bcdacf`.
type SendMessageAlbumRequest struct {
	// Target chat
	ChatID int64
	// If not 0, a message thread identifier in which the messages will be sent
	MessageThreadID int64
	// Information about the message or story to be replied; pass null if none
	ReplyTo InputMessageReplyToClass
	// Options to be used to send the messages; pass null to use default options
	Options MessageSendOptions
	// Contents of messages to be sent. At most 10 messages can be added to an album
	InputMessageContents []InputMessageContentClass
}

// SendMessageAlbumRequestTypeID is TL type id of SendMessageAlbumRequest.
const SendMessageAlbumRequestTypeID = 0x74bcdacf

// Ensuring interfaces in compile-time for SendMessageAlbumRequest.
var (
	_ bin.Encoder     = &SendMessageAlbumRequest{}
	_ bin.Decoder     = &SendMessageAlbumRequest{}
	_ bin.BareEncoder = &SendMessageAlbumRequest{}
	_ bin.BareDecoder = &SendMessageAlbumRequest{}
)

func (s *SendMessageAlbumRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ChatID == 0) {
		return false
	}
	if !(s.MessageThreadID == 0) {
		return false
	}
	if !(s.ReplyTo == nil) {
		return false
	}
	if !(s.Options.Zero()) {
		return false
	}
	if !(s.InputMessageContents == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SendMessageAlbumRequest) String() string {
	if s == nil {
		return "SendMessageAlbumRequest(nil)"
	}
	type Alias SendMessageAlbumRequest
	return fmt.Sprintf("SendMessageAlbumRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SendMessageAlbumRequest) TypeID() uint32 {
	return SendMessageAlbumRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SendMessageAlbumRequest) TypeName() string {
	return "sendMessageAlbum"
}

// TypeInfo returns info about TL type.
func (s *SendMessageAlbumRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "sendMessageAlbum",
		ID:   SendMessageAlbumRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageThreadID",
			SchemaName: "message_thread_id",
		},
		{
			Name:       "ReplyTo",
			SchemaName: "reply_to",
		},
		{
			Name:       "Options",
			SchemaName: "options",
		},
		{
			Name:       "InputMessageContents",
			SchemaName: "input_message_contents",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SendMessageAlbumRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendMessageAlbum#74bcdacf as nil")
	}
	b.PutID(SendMessageAlbumRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SendMessageAlbumRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendMessageAlbum#74bcdacf as nil")
	}
	b.PutInt53(s.ChatID)
	b.PutInt53(s.MessageThreadID)
	if s.ReplyTo == nil {
		return fmt.Errorf("unable to encode sendMessageAlbum#74bcdacf: field reply_to is nil")
	}
	if err := s.ReplyTo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode sendMessageAlbum#74bcdacf: field reply_to: %w", err)
	}
	if err := s.Options.Encode(b); err != nil {
		return fmt.Errorf("unable to encode sendMessageAlbum#74bcdacf: field options: %w", err)
	}
	b.PutInt(len(s.InputMessageContents))
	for idx, v := range s.InputMessageContents {
		if v == nil {
			return fmt.Errorf("unable to encode sendMessageAlbum#74bcdacf: field input_message_contents element with index %d is nil", idx)
		}
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare sendMessageAlbum#74bcdacf: field input_message_contents element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SendMessageAlbumRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendMessageAlbum#74bcdacf to nil")
	}
	if err := b.ConsumeID(SendMessageAlbumRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode sendMessageAlbum#74bcdacf: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SendMessageAlbumRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendMessageAlbum#74bcdacf to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode sendMessageAlbum#74bcdacf: field chat_id: %w", err)
		}
		s.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode sendMessageAlbum#74bcdacf: field message_thread_id: %w", err)
		}
		s.MessageThreadID = value
	}
	{
		value, err := DecodeInputMessageReplyTo(b)
		if err != nil {
			return fmt.Errorf("unable to decode sendMessageAlbum#74bcdacf: field reply_to: %w", err)
		}
		s.ReplyTo = value
	}
	{
		if err := s.Options.Decode(b); err != nil {
			return fmt.Errorf("unable to decode sendMessageAlbum#74bcdacf: field options: %w", err)
		}
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode sendMessageAlbum#74bcdacf: field input_message_contents: %w", err)
		}

		if headerLen > 0 {
			s.InputMessageContents = make([]InputMessageContentClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputMessageContent(b)
			if err != nil {
				return fmt.Errorf("unable to decode sendMessageAlbum#74bcdacf: field input_message_contents: %w", err)
			}
			s.InputMessageContents = append(s.InputMessageContents, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SendMessageAlbumRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode sendMessageAlbum#74bcdacf as nil")
	}
	b.ObjStart()
	b.PutID("sendMessageAlbum")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(s.ChatID)
	b.Comma()
	b.FieldStart("message_thread_id")
	b.PutInt53(s.MessageThreadID)
	b.Comma()
	b.FieldStart("reply_to")
	if s.ReplyTo == nil {
		return fmt.Errorf("unable to encode sendMessageAlbum#74bcdacf: field reply_to is nil")
	}
	if err := s.ReplyTo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode sendMessageAlbum#74bcdacf: field reply_to: %w", err)
	}
	b.Comma()
	b.FieldStart("options")
	if err := s.Options.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode sendMessageAlbum#74bcdacf: field options: %w", err)
	}
	b.Comma()
	b.FieldStart("input_message_contents")
	b.ArrStart()
	for idx, v := range s.InputMessageContents {
		if v == nil {
			return fmt.Errorf("unable to encode sendMessageAlbum#74bcdacf: field input_message_contents element with index %d is nil", idx)
		}
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode sendMessageAlbum#74bcdacf: field input_message_contents element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SendMessageAlbumRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode sendMessageAlbum#74bcdacf to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("sendMessageAlbum"); err != nil {
				return fmt.Errorf("unable to decode sendMessageAlbum#74bcdacf: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode sendMessageAlbum#74bcdacf: field chat_id: %w", err)
			}
			s.ChatID = value
		case "message_thread_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode sendMessageAlbum#74bcdacf: field message_thread_id: %w", err)
			}
			s.MessageThreadID = value
		case "reply_to":
			value, err := DecodeTDLibJSONInputMessageReplyTo(b)
			if err != nil {
				return fmt.Errorf("unable to decode sendMessageAlbum#74bcdacf: field reply_to: %w", err)
			}
			s.ReplyTo = value
		case "options":
			if err := s.Options.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode sendMessageAlbum#74bcdacf: field options: %w", err)
			}
		case "input_message_contents":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := DecodeTDLibJSONInputMessageContent(b)
				if err != nil {
					return fmt.Errorf("unable to decode sendMessageAlbum#74bcdacf: field input_message_contents: %w", err)
				}
				s.InputMessageContents = append(s.InputMessageContents, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode sendMessageAlbum#74bcdacf: field input_message_contents: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (s *SendMessageAlbumRequest) GetChatID() (value int64) {
	if s == nil {
		return
	}
	return s.ChatID
}

// GetMessageThreadID returns value of MessageThreadID field.
func (s *SendMessageAlbumRequest) GetMessageThreadID() (value int64) {
	if s == nil {
		return
	}
	return s.MessageThreadID
}

// GetReplyTo returns value of ReplyTo field.
func (s *SendMessageAlbumRequest) GetReplyTo() (value InputMessageReplyToClass) {
	if s == nil {
		return
	}
	return s.ReplyTo
}

// GetOptions returns value of Options field.
func (s *SendMessageAlbumRequest) GetOptions() (value MessageSendOptions) {
	if s == nil {
		return
	}
	return s.Options
}

// GetInputMessageContents returns value of InputMessageContents field.
func (s *SendMessageAlbumRequest) GetInputMessageContents() (value []InputMessageContentClass) {
	if s == nil {
		return
	}
	return s.InputMessageContents
}

// SendMessageAlbum invokes method sendMessageAlbum#74bcdacf returning error if any.
func (c *Client) SendMessageAlbum(ctx context.Context, request *SendMessageAlbumRequest) (*Messages, error) {
	var result Messages

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
