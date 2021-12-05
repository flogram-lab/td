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

// PollTypeRegular represents TL type `pollTypeRegular#2638f022`.
type PollTypeRegular struct {
	// True, if multiple answer options can be chosen simultaneously
	AllowMultipleAnswers bool
}

// PollTypeRegularTypeID is TL type id of PollTypeRegular.
const PollTypeRegularTypeID = 0x2638f022

// construct implements constructor of PollTypeClass.
func (p PollTypeRegular) construct() PollTypeClass { return &p }

// Ensuring interfaces in compile-time for PollTypeRegular.
var (
	_ bin.Encoder     = &PollTypeRegular{}
	_ bin.Decoder     = &PollTypeRegular{}
	_ bin.BareEncoder = &PollTypeRegular{}
	_ bin.BareDecoder = &PollTypeRegular{}

	_ PollTypeClass = &PollTypeRegular{}
)

func (p *PollTypeRegular) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.AllowMultipleAnswers == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PollTypeRegular) String() string {
	if p == nil {
		return "PollTypeRegular(nil)"
	}
	type Alias PollTypeRegular
	return fmt.Sprintf("PollTypeRegular%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PollTypeRegular) TypeID() uint32 {
	return PollTypeRegularTypeID
}

// TypeName returns name of type in TL schema.
func (*PollTypeRegular) TypeName() string {
	return "pollTypeRegular"
}

// TypeInfo returns info about TL type.
func (p *PollTypeRegular) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "pollTypeRegular",
		ID:   PollTypeRegularTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "AllowMultipleAnswers",
			SchemaName: "allow_multiple_answers",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PollTypeRegular) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pollTypeRegular#2638f022 as nil")
	}
	b.PutID(PollTypeRegularTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PollTypeRegular) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pollTypeRegular#2638f022 as nil")
	}
	b.PutBool(p.AllowMultipleAnswers)
	return nil
}

// Decode implements bin.Decoder.
func (p *PollTypeRegular) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pollTypeRegular#2638f022 to nil")
	}
	if err := b.ConsumeID(PollTypeRegularTypeID); err != nil {
		return fmt.Errorf("unable to decode pollTypeRegular#2638f022: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PollTypeRegular) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pollTypeRegular#2638f022 to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode pollTypeRegular#2638f022: field allow_multiple_answers: %w", err)
		}
		p.AllowMultipleAnswers = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *PollTypeRegular) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode pollTypeRegular#2638f022 as nil")
	}
	b.ObjStart()
	b.PutID("pollTypeRegular")
	b.FieldStart("allow_multiple_answers")
	b.PutBool(p.AllowMultipleAnswers)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *PollTypeRegular) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode pollTypeRegular#2638f022 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("pollTypeRegular"); err != nil {
				return fmt.Errorf("unable to decode pollTypeRegular#2638f022: %w", err)
			}
		case "allow_multiple_answers":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode pollTypeRegular#2638f022: field allow_multiple_answers: %w", err)
			}
			p.AllowMultipleAnswers = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetAllowMultipleAnswers returns value of AllowMultipleAnswers field.
func (p *PollTypeRegular) GetAllowMultipleAnswers() (value bool) {
	return p.AllowMultipleAnswers
}

// PollTypeQuiz represents TL type `pollTypeQuiz#27293c99`.
type PollTypeQuiz struct {
	// 0-based identifier of the correct answer option; -1 for a yet unanswered poll
	CorrectOptionID int32
	// Text that is shown when the user chooses an incorrect answer or taps on the lamp icon;
	// 0-200 characters with at most 2 line feeds; empty for a yet unanswered poll
	Explanation FormattedText
}

// PollTypeQuizTypeID is TL type id of PollTypeQuiz.
const PollTypeQuizTypeID = 0x27293c99

// construct implements constructor of PollTypeClass.
func (p PollTypeQuiz) construct() PollTypeClass { return &p }

// Ensuring interfaces in compile-time for PollTypeQuiz.
var (
	_ bin.Encoder     = &PollTypeQuiz{}
	_ bin.Decoder     = &PollTypeQuiz{}
	_ bin.BareEncoder = &PollTypeQuiz{}
	_ bin.BareDecoder = &PollTypeQuiz{}

	_ PollTypeClass = &PollTypeQuiz{}
)

func (p *PollTypeQuiz) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.CorrectOptionID == 0) {
		return false
	}
	if !(p.Explanation.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PollTypeQuiz) String() string {
	if p == nil {
		return "PollTypeQuiz(nil)"
	}
	type Alias PollTypeQuiz
	return fmt.Sprintf("PollTypeQuiz%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PollTypeQuiz) TypeID() uint32 {
	return PollTypeQuizTypeID
}

// TypeName returns name of type in TL schema.
func (*PollTypeQuiz) TypeName() string {
	return "pollTypeQuiz"
}

// TypeInfo returns info about TL type.
func (p *PollTypeQuiz) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "pollTypeQuiz",
		ID:   PollTypeQuizTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "CorrectOptionID",
			SchemaName: "correct_option_id",
		},
		{
			Name:       "Explanation",
			SchemaName: "explanation",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PollTypeQuiz) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pollTypeQuiz#27293c99 as nil")
	}
	b.PutID(PollTypeQuizTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PollTypeQuiz) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pollTypeQuiz#27293c99 as nil")
	}
	b.PutInt32(p.CorrectOptionID)
	if err := p.Explanation.Encode(b); err != nil {
		return fmt.Errorf("unable to encode pollTypeQuiz#27293c99: field explanation: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PollTypeQuiz) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pollTypeQuiz#27293c99 to nil")
	}
	if err := b.ConsumeID(PollTypeQuizTypeID); err != nil {
		return fmt.Errorf("unable to decode pollTypeQuiz#27293c99: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PollTypeQuiz) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pollTypeQuiz#27293c99 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode pollTypeQuiz#27293c99: field correct_option_id: %w", err)
		}
		p.CorrectOptionID = value
	}
	{
		if err := p.Explanation.Decode(b); err != nil {
			return fmt.Errorf("unable to decode pollTypeQuiz#27293c99: field explanation: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *PollTypeQuiz) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode pollTypeQuiz#27293c99 as nil")
	}
	b.ObjStart()
	b.PutID("pollTypeQuiz")
	b.FieldStart("correct_option_id")
	b.PutInt32(p.CorrectOptionID)
	b.FieldStart("explanation")
	if err := p.Explanation.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode pollTypeQuiz#27293c99: field explanation: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *PollTypeQuiz) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode pollTypeQuiz#27293c99 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("pollTypeQuiz"); err != nil {
				return fmt.Errorf("unable to decode pollTypeQuiz#27293c99: %w", err)
			}
		case "correct_option_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode pollTypeQuiz#27293c99: field correct_option_id: %w", err)
			}
			p.CorrectOptionID = value
		case "explanation":
			if err := p.Explanation.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode pollTypeQuiz#27293c99: field explanation: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetCorrectOptionID returns value of CorrectOptionID field.
func (p *PollTypeQuiz) GetCorrectOptionID() (value int32) {
	return p.CorrectOptionID
}

// GetExplanation returns value of Explanation field.
func (p *PollTypeQuiz) GetExplanation() (value FormattedText) {
	return p.Explanation
}

// PollTypeClassName is schema name of PollTypeClass.
const PollTypeClassName = "PollType"

// PollTypeClass represents PollType generic type.
//
// Example:
//  g, err := tdapi.DecodePollType(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tdapi.PollTypeRegular: // pollTypeRegular#2638f022
//  case *tdapi.PollTypeQuiz: // pollTypeQuiz#27293c99
//  default: panic(v)
//  }
type PollTypeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() PollTypeClass

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

// DecodePollType implements binary de-serialization for PollTypeClass.
func DecodePollType(buf *bin.Buffer) (PollTypeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case PollTypeRegularTypeID:
		// Decoding pollTypeRegular#2638f022.
		v := PollTypeRegular{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PollTypeClass: %w", err)
		}
		return &v, nil
	case PollTypeQuizTypeID:
		// Decoding pollTypeQuiz#27293c99.
		v := PollTypeQuiz{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PollTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PollTypeClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONPollType implements binary de-serialization for PollTypeClass.
func DecodeTDLibJSONPollType(buf tdjson.Decoder) (PollTypeClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "pollTypeRegular":
		// Decoding pollTypeRegular#2638f022.
		v := PollTypeRegular{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PollTypeClass: %w", err)
		}
		return &v, nil
	case "pollTypeQuiz":
		// Decoding pollTypeQuiz#27293c99.
		v := PollTypeQuiz{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PollTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PollTypeClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// PollType boxes the PollTypeClass providing a helper.
type PollTypeBox struct {
	PollType PollTypeClass
}

// Decode implements bin.Decoder for PollTypeBox.
func (b *PollTypeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode PollTypeBox to nil")
	}
	v, err := DecodePollType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.PollType = v
	return nil
}

// Encode implements bin.Encode for PollTypeBox.
func (b *PollTypeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.PollType == nil {
		return fmt.Errorf("unable to encode PollTypeClass as nil")
	}
	return b.PollType.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for PollTypeBox.
func (b *PollTypeBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode PollTypeBox to nil")
	}
	v, err := DecodeTDLibJSONPollType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.PollType = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for PollTypeBox.
func (b *PollTypeBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.PollType == nil {
		return fmt.Errorf("unable to encode PollTypeClass as nil")
	}
	return b.PollType.EncodeTDLibJSON(buf)
}
