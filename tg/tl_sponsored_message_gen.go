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

// SponsoredMessage represents TL type `sponsoredMessage#fc25b828`.
// A sponsored message¹.
//
// Links:
//  1. https://core.telegram.org/api/sponsored-messages
//
// See https://core.telegram.org/constructor/sponsoredMessage for reference.
type SponsoredMessage struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether the message needs to be labeled as "recommended" instead of "sponsored"
	Recommended bool
	// Whether a profile photo bubble should be displayed for this message, like for messages
	// sent in groups. The photo shown in the bubble is obtained either from the peer
	// contained in from_id, or from chat_invite.
	ShowPeerPhoto bool
	// Message ID
	RandomID []byte
	// ID of the sender of the message
	//
	// Use SetFromID and GetFromID helpers.
	FromID PeerClass
	// Information about the chat invite hash specified in chat_invite_hash
	//
	// Use SetChatInvite and GetChatInvite helpers.
	ChatInvite ChatInviteClass
	// Chat invite
	//
	// Use SetChatInviteHash and GetChatInviteHash helpers.
	ChatInviteHash string
	// Optional link to a channel post if from_id points to a channel
	//
	// Use SetChannelPost and GetChannelPost helpers.
	ChannelPost int
	// Parameter for the bot start message if the sponsored chat is a chat with a bot.
	//
	// Use SetStartParam and GetStartParam helpers.
	StartParam string
	// Sponsored message
	Message string
	// Message entities for styled text¹
	//
	// Links:
	//  1) https://core.telegram.org/api/entities
	//
	// Use SetEntities and GetEntities helpers.
	Entities []MessageEntityClass
	// If set, contains additional information about the sponsor to be shown along with the
	// message.
	//
	// Use SetSponsorInfo and GetSponsorInfo helpers.
	SponsorInfo string
	// If set, contains additional information about the sponsored message to be shown along
	// with the message.
	//
	// Use SetAdditionalInfo and GetAdditionalInfo helpers.
	AdditionalInfo string
}

// SponsoredMessageTypeID is TL type id of SponsoredMessage.
const SponsoredMessageTypeID = 0xfc25b828

// Ensuring interfaces in compile-time for SponsoredMessage.
var (
	_ bin.Encoder     = &SponsoredMessage{}
	_ bin.Decoder     = &SponsoredMessage{}
	_ bin.BareEncoder = &SponsoredMessage{}
	_ bin.BareDecoder = &SponsoredMessage{}
)

func (s *SponsoredMessage) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Recommended == false) {
		return false
	}
	if !(s.ShowPeerPhoto == false) {
		return false
	}
	if !(s.RandomID == nil) {
		return false
	}
	if !(s.FromID == nil) {
		return false
	}
	if !(s.ChatInvite == nil) {
		return false
	}
	if !(s.ChatInviteHash == "") {
		return false
	}
	if !(s.ChannelPost == 0) {
		return false
	}
	if !(s.StartParam == "") {
		return false
	}
	if !(s.Message == "") {
		return false
	}
	if !(s.Entities == nil) {
		return false
	}
	if !(s.SponsorInfo == "") {
		return false
	}
	if !(s.AdditionalInfo == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SponsoredMessage) String() string {
	if s == nil {
		return "SponsoredMessage(nil)"
	}
	type Alias SponsoredMessage
	return fmt.Sprintf("SponsoredMessage%+v", Alias(*s))
}

// FillFrom fills SponsoredMessage from given interface.
func (s *SponsoredMessage) FillFrom(from interface {
	GetRecommended() (value bool)
	GetShowPeerPhoto() (value bool)
	GetRandomID() (value []byte)
	GetFromID() (value PeerClass, ok bool)
	GetChatInvite() (value ChatInviteClass, ok bool)
	GetChatInviteHash() (value string, ok bool)
	GetChannelPost() (value int, ok bool)
	GetStartParam() (value string, ok bool)
	GetMessage() (value string)
	GetEntities() (value []MessageEntityClass, ok bool)
	GetSponsorInfo() (value string, ok bool)
	GetAdditionalInfo() (value string, ok bool)
}) {
	s.Recommended = from.GetRecommended()
	s.ShowPeerPhoto = from.GetShowPeerPhoto()
	s.RandomID = from.GetRandomID()
	if val, ok := from.GetFromID(); ok {
		s.FromID = val
	}

	if val, ok := from.GetChatInvite(); ok {
		s.ChatInvite = val
	}

	if val, ok := from.GetChatInviteHash(); ok {
		s.ChatInviteHash = val
	}

	if val, ok := from.GetChannelPost(); ok {
		s.ChannelPost = val
	}

	if val, ok := from.GetStartParam(); ok {
		s.StartParam = val
	}

	s.Message = from.GetMessage()
	if val, ok := from.GetEntities(); ok {
		s.Entities = val
	}

	if val, ok := from.GetSponsorInfo(); ok {
		s.SponsorInfo = val
	}

	if val, ok := from.GetAdditionalInfo(); ok {
		s.AdditionalInfo = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SponsoredMessage) TypeID() uint32 {
	return SponsoredMessageTypeID
}

// TypeName returns name of type in TL schema.
func (*SponsoredMessage) TypeName() string {
	return "sponsoredMessage"
}

// TypeInfo returns info about TL type.
func (s *SponsoredMessage) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "sponsoredMessage",
		ID:   SponsoredMessageTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Recommended",
			SchemaName: "recommended",
			Null:       !s.Flags.Has(5),
		},
		{
			Name:       "ShowPeerPhoto",
			SchemaName: "show_peer_photo",
			Null:       !s.Flags.Has(6),
		},
		{
			Name:       "RandomID",
			SchemaName: "random_id",
		},
		{
			Name:       "FromID",
			SchemaName: "from_id",
			Null:       !s.Flags.Has(3),
		},
		{
			Name:       "ChatInvite",
			SchemaName: "chat_invite",
			Null:       !s.Flags.Has(4),
		},
		{
			Name:       "ChatInviteHash",
			SchemaName: "chat_invite_hash",
			Null:       !s.Flags.Has(4),
		},
		{
			Name:       "ChannelPost",
			SchemaName: "channel_post",
			Null:       !s.Flags.Has(2),
		},
		{
			Name:       "StartParam",
			SchemaName: "start_param",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "Message",
			SchemaName: "message",
		},
		{
			Name:       "Entities",
			SchemaName: "entities",
			Null:       !s.Flags.Has(1),
		},
		{
			Name:       "SponsorInfo",
			SchemaName: "sponsor_info",
			Null:       !s.Flags.Has(7),
		},
		{
			Name:       "AdditionalInfo",
			SchemaName: "additional_info",
			Null:       !s.Flags.Has(8),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *SponsoredMessage) SetFlags() {
	if !(s.Recommended == false) {
		s.Flags.Set(5)
	}
	if !(s.ShowPeerPhoto == false) {
		s.Flags.Set(6)
	}
	if !(s.FromID == nil) {
		s.Flags.Set(3)
	}
	if !(s.ChatInvite == nil) {
		s.Flags.Set(4)
	}
	if !(s.ChatInviteHash == "") {
		s.Flags.Set(4)
	}
	if !(s.ChannelPost == 0) {
		s.Flags.Set(2)
	}
	if !(s.StartParam == "") {
		s.Flags.Set(0)
	}
	if !(s.Entities == nil) {
		s.Flags.Set(1)
	}
	if !(s.SponsorInfo == "") {
		s.Flags.Set(7)
	}
	if !(s.AdditionalInfo == "") {
		s.Flags.Set(8)
	}
}

// Encode implements bin.Encoder.
func (s *SponsoredMessage) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sponsoredMessage#fc25b828 as nil")
	}
	b.PutID(SponsoredMessageTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SponsoredMessage) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sponsoredMessage#fc25b828 as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode sponsoredMessage#fc25b828: field flags: %w", err)
	}
	b.PutBytes(s.RandomID)
	if s.Flags.Has(3) {
		if s.FromID == nil {
			return fmt.Errorf("unable to encode sponsoredMessage#fc25b828: field from_id is nil")
		}
		if err := s.FromID.Encode(b); err != nil {
			return fmt.Errorf("unable to encode sponsoredMessage#fc25b828: field from_id: %w", err)
		}
	}
	if s.Flags.Has(4) {
		if s.ChatInvite == nil {
			return fmt.Errorf("unable to encode sponsoredMessage#fc25b828: field chat_invite is nil")
		}
		if err := s.ChatInvite.Encode(b); err != nil {
			return fmt.Errorf("unable to encode sponsoredMessage#fc25b828: field chat_invite: %w", err)
		}
	}
	if s.Flags.Has(4) {
		b.PutString(s.ChatInviteHash)
	}
	if s.Flags.Has(2) {
		b.PutInt(s.ChannelPost)
	}
	if s.Flags.Has(0) {
		b.PutString(s.StartParam)
	}
	b.PutString(s.Message)
	if s.Flags.Has(1) {
		b.PutVectorHeader(len(s.Entities))
		for idx, v := range s.Entities {
			if v == nil {
				return fmt.Errorf("unable to encode sponsoredMessage#fc25b828: field entities element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode sponsoredMessage#fc25b828: field entities element with index %d: %w", idx, err)
			}
		}
	}
	if s.Flags.Has(7) {
		b.PutString(s.SponsorInfo)
	}
	if s.Flags.Has(8) {
		b.PutString(s.AdditionalInfo)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SponsoredMessage) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sponsoredMessage#fc25b828 to nil")
	}
	if err := b.ConsumeID(SponsoredMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode sponsoredMessage#fc25b828: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SponsoredMessage) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sponsoredMessage#fc25b828 to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#fc25b828: field flags: %w", err)
		}
	}
	s.Recommended = s.Flags.Has(5)
	s.ShowPeerPhoto = s.Flags.Has(6)
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#fc25b828: field random_id: %w", err)
		}
		s.RandomID = value
	}
	if s.Flags.Has(3) {
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#fc25b828: field from_id: %w", err)
		}
		s.FromID = value
	}
	if s.Flags.Has(4) {
		value, err := DecodeChatInvite(b)
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#fc25b828: field chat_invite: %w", err)
		}
		s.ChatInvite = value
	}
	if s.Flags.Has(4) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#fc25b828: field chat_invite_hash: %w", err)
		}
		s.ChatInviteHash = value
	}
	if s.Flags.Has(2) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#fc25b828: field channel_post: %w", err)
		}
		s.ChannelPost = value
	}
	if s.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#fc25b828: field start_param: %w", err)
		}
		s.StartParam = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#fc25b828: field message: %w", err)
		}
		s.Message = value
	}
	if s.Flags.Has(1) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#fc25b828: field entities: %w", err)
		}

		if headerLen > 0 {
			s.Entities = make([]MessageEntityClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode sponsoredMessage#fc25b828: field entities: %w", err)
			}
			s.Entities = append(s.Entities, value)
		}
	}
	if s.Flags.Has(7) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#fc25b828: field sponsor_info: %w", err)
		}
		s.SponsorInfo = value
	}
	if s.Flags.Has(8) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode sponsoredMessage#fc25b828: field additional_info: %w", err)
		}
		s.AdditionalInfo = value
	}
	return nil
}

// SetRecommended sets value of Recommended conditional field.
func (s *SponsoredMessage) SetRecommended(value bool) {
	if value {
		s.Flags.Set(5)
		s.Recommended = true
	} else {
		s.Flags.Unset(5)
		s.Recommended = false
	}
}

// GetRecommended returns value of Recommended conditional field.
func (s *SponsoredMessage) GetRecommended() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(5)
}

// SetShowPeerPhoto sets value of ShowPeerPhoto conditional field.
func (s *SponsoredMessage) SetShowPeerPhoto(value bool) {
	if value {
		s.Flags.Set(6)
		s.ShowPeerPhoto = true
	} else {
		s.Flags.Unset(6)
		s.ShowPeerPhoto = false
	}
}

// GetShowPeerPhoto returns value of ShowPeerPhoto conditional field.
func (s *SponsoredMessage) GetShowPeerPhoto() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(6)
}

// GetRandomID returns value of RandomID field.
func (s *SponsoredMessage) GetRandomID() (value []byte) {
	if s == nil {
		return
	}
	return s.RandomID
}

// SetFromID sets value of FromID conditional field.
func (s *SponsoredMessage) SetFromID(value PeerClass) {
	s.Flags.Set(3)
	s.FromID = value
}

// GetFromID returns value of FromID conditional field and
// boolean which is true if field was set.
func (s *SponsoredMessage) GetFromID() (value PeerClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(3) {
		return value, false
	}
	return s.FromID, true
}

// SetChatInvite sets value of ChatInvite conditional field.
func (s *SponsoredMessage) SetChatInvite(value ChatInviteClass) {
	s.Flags.Set(4)
	s.ChatInvite = value
}

// GetChatInvite returns value of ChatInvite conditional field and
// boolean which is true if field was set.
func (s *SponsoredMessage) GetChatInvite() (value ChatInviteClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(4) {
		return value, false
	}
	return s.ChatInvite, true
}

// SetChatInviteHash sets value of ChatInviteHash conditional field.
func (s *SponsoredMessage) SetChatInviteHash(value string) {
	s.Flags.Set(4)
	s.ChatInviteHash = value
}

// GetChatInviteHash returns value of ChatInviteHash conditional field and
// boolean which is true if field was set.
func (s *SponsoredMessage) GetChatInviteHash() (value string, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(4) {
		return value, false
	}
	return s.ChatInviteHash, true
}

// SetChannelPost sets value of ChannelPost conditional field.
func (s *SponsoredMessage) SetChannelPost(value int) {
	s.Flags.Set(2)
	s.ChannelPost = value
}

// GetChannelPost returns value of ChannelPost conditional field and
// boolean which is true if field was set.
func (s *SponsoredMessage) GetChannelPost() (value int, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(2) {
		return value, false
	}
	return s.ChannelPost, true
}

// SetStartParam sets value of StartParam conditional field.
func (s *SponsoredMessage) SetStartParam(value string) {
	s.Flags.Set(0)
	s.StartParam = value
}

// GetStartParam returns value of StartParam conditional field and
// boolean which is true if field was set.
func (s *SponsoredMessage) GetStartParam() (value string, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.StartParam, true
}

// GetMessage returns value of Message field.
func (s *SponsoredMessage) GetMessage() (value string) {
	if s == nil {
		return
	}
	return s.Message
}

// SetEntities sets value of Entities conditional field.
func (s *SponsoredMessage) SetEntities(value []MessageEntityClass) {
	s.Flags.Set(1)
	s.Entities = value
}

// GetEntities returns value of Entities conditional field and
// boolean which is true if field was set.
func (s *SponsoredMessage) GetEntities() (value []MessageEntityClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(1) {
		return value, false
	}
	return s.Entities, true
}

// SetSponsorInfo sets value of SponsorInfo conditional field.
func (s *SponsoredMessage) SetSponsorInfo(value string) {
	s.Flags.Set(7)
	s.SponsorInfo = value
}

// GetSponsorInfo returns value of SponsorInfo conditional field and
// boolean which is true if field was set.
func (s *SponsoredMessage) GetSponsorInfo() (value string, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(7) {
		return value, false
	}
	return s.SponsorInfo, true
}

// SetAdditionalInfo sets value of AdditionalInfo conditional field.
func (s *SponsoredMessage) SetAdditionalInfo(value string) {
	s.Flags.Set(8)
	s.AdditionalInfo = value
}

// GetAdditionalInfo returns value of AdditionalInfo conditional field and
// boolean which is true if field was set.
func (s *SponsoredMessage) GetAdditionalInfo() (value string, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(8) {
		return value, false
	}
	return s.AdditionalInfo, true
}

// MapEntities returns field Entities wrapped in MessageEntityClassArray helper.
func (s *SponsoredMessage) MapEntities() (value MessageEntityClassArray, ok bool) {
	if !s.Flags.Has(1) {
		return value, false
	}
	return MessageEntityClassArray(s.Entities), true
}
