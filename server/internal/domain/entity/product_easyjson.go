// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package entity

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonCf3f67efDecodeServerInternalDomainEntity(in *jlexer.Lexer, out *Product) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ID":
			out.ID = uint(in.Uint())
		case "Name":
			out.Name = string(in.String())
		case "Price":
			out.Price = float32(in.Float32())
		case "CookingTime":
			out.CookingTime = int(in.Int())
		case "Portion":
			out.Portion = string(in.String())
		case "Description":
			out.Description = string(in.String())
		case "Icon":
			out.Icon = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCf3f67efEncodeServerInternalDomainEntity(out *jwriter.Writer, in Product) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ID\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.ID))
	}
	{
		const prefix string = ",\"Name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"Price\":"
		out.RawString(prefix)
		out.Float32(float32(in.Price))
	}
	{
		const prefix string = ",\"CookingTime\":"
		out.RawString(prefix)
		out.Int(int(in.CookingTime))
	}
	{
		const prefix string = ",\"Portion\":"
		out.RawString(prefix)
		out.String(string(in.Portion))
	}
	{
		const prefix string = ",\"Description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"Icon\":"
		out.RawString(prefix)
		out.String(string(in.Icon))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Product) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCf3f67efEncodeServerInternalDomainEntity(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Product) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCf3f67efEncodeServerInternalDomainEntity(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Product) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCf3f67efDecodeServerInternalDomainEntity(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Product) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCf3f67efDecodeServerInternalDomainEntity(l, v)
}
