// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package dto

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

func easyjsonF4fdf71eDecodeServerEasy(in *jlexer.Lexer, out *RespOrderAddress) {
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
		case "City":
			out.City = string(in.String())
		case "Street":
			out.Street = string(in.String())
		case "House":
			out.House = string(in.String())
		case "Flat":
			out.Flat = uint(in.Uint())
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
func easyjsonF4fdf71eEncodeServerEasy(out *jwriter.Writer, in RespOrderAddress) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"City\":"
		out.RawString(prefix[1:])
		out.String(string(in.City))
	}
	{
		const prefix string = ",\"Street\":"
		out.RawString(prefix)
		out.String(string(in.Street))
	}
	{
		const prefix string = ",\"House\":"
		out.RawString(prefix)
		out.String(string(in.House))
	}
	{
		const prefix string = ",\"Flat\":"
		out.RawString(prefix)
		out.Uint(uint(in.Flat))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RespOrderAddress) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF4fdf71eEncodeServerEasy(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RespOrderAddress) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF4fdf71eEncodeServerEasy(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RespOrderAddress) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF4fdf71eDecodeServerEasy(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RespOrderAddress) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF4fdf71eDecodeServerEasy(l, v)
}
func easyjsonF4fdf71eDecodeServerEasy1(in *jlexer.Lexer, out *RespGetAddresses) {
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
		case "Addresses":
			if in.IsNull() {
				in.Skip()
				out.Addresses = nil
			} else {
				in.Delim('[')
				if out.Addresses == nil {
					if !in.IsDelim(']') {
						out.Addresses = make([]*RespGetAddress, 0, 8)
					} else {
						out.Addresses = []*RespGetAddress{}
					}
				} else {
					out.Addresses = (out.Addresses)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *RespGetAddress
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(RespGetAddress)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Addresses = append(out.Addresses, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "CurrentAddressesId":
			out.CurrentAddressesId = uint(in.Uint())
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
func easyjsonF4fdf71eEncodeServerEasy1(out *jwriter.Writer, in RespGetAddresses) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Addresses\":"
		out.RawString(prefix[1:])
		if in.Addresses == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Addresses {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"CurrentAddressesId\":"
		out.RawString(prefix)
		out.Uint(uint(in.CurrentAddressesId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RespGetAddresses) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF4fdf71eEncodeServerEasy1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RespGetAddresses) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF4fdf71eEncodeServerEasy1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RespGetAddresses) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF4fdf71eDecodeServerEasy1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RespGetAddresses) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF4fdf71eDecodeServerEasy1(l, v)
}
func easyjsonF4fdf71eDecodeServerEasy2(in *jlexer.Lexer, out *RespGetAddress) {
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
		case "Id":
			out.Id = uint(in.Uint())
		case "City":
			out.City = string(in.String())
		case "Street":
			out.Street = string(in.String())
		case "House":
			out.House = string(in.String())
		case "Flat":
			out.Flat = uint(in.Uint())
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
func easyjsonF4fdf71eEncodeServerEasy2(out *jwriter.Writer, in RespGetAddress) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Id\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.Id))
	}
	{
		const prefix string = ",\"City\":"
		out.RawString(prefix)
		out.String(string(in.City))
	}
	{
		const prefix string = ",\"Street\":"
		out.RawString(prefix)
		out.String(string(in.Street))
	}
	{
		const prefix string = ",\"House\":"
		out.RawString(prefix)
		out.String(string(in.House))
	}
	{
		const prefix string = ",\"Flat\":"
		out.RawString(prefix)
		out.Uint(uint(in.Flat))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RespGetAddress) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF4fdf71eEncodeServerEasy2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RespGetAddress) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF4fdf71eEncodeServerEasy2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RespGetAddress) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF4fdf71eDecodeServerEasy2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RespGetAddress) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF4fdf71eDecodeServerEasy2(l, v)
}
func easyjsonF4fdf71eDecodeServerEasy3(in *jlexer.Lexer, out *ReqCreateOrderAddress) {
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
		case "City":
			out.City = string(in.String())
		case "Street":
			out.Street = string(in.String())
		case "House":
			out.House = string(in.String())
		case "Flat":
			out.Flat = uint(in.Uint())
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
func easyjsonF4fdf71eEncodeServerEasy3(out *jwriter.Writer, in ReqCreateOrderAddress) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"City\":"
		out.RawString(prefix[1:])
		out.String(string(in.City))
	}
	{
		const prefix string = ",\"Street\":"
		out.RawString(prefix)
		out.String(string(in.Street))
	}
	{
		const prefix string = ",\"House\":"
		out.RawString(prefix)
		out.String(string(in.House))
	}
	{
		const prefix string = ",\"Flat\":"
		out.RawString(prefix)
		out.Uint(uint(in.Flat))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ReqCreateOrderAddress) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF4fdf71eEncodeServerEasy3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ReqCreateOrderAddress) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF4fdf71eEncodeServerEasy3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ReqCreateOrderAddress) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF4fdf71eDecodeServerEasy3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ReqCreateOrderAddress) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF4fdf71eDecodeServerEasy3(l, v)
}
func easyjsonF4fdf71eDecodeServerEasy4(in *jlexer.Lexer, out *ReqCreateAddress) {
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
		case "City":
			out.City = string(in.String())
		case "Street":
			out.Street = string(in.String())
		case "House":
			out.House = string(in.String())
		case "Flat":
			out.Flat = uint(in.Uint())
		case "Cookie":
			out.Cookie = string(in.String())
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
func easyjsonF4fdf71eEncodeServerEasy4(out *jwriter.Writer, in ReqCreateAddress) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"City\":"
		out.RawString(prefix[1:])
		out.String(string(in.City))
	}
	{
		const prefix string = ",\"Street\":"
		out.RawString(prefix)
		out.String(string(in.Street))
	}
	{
		const prefix string = ",\"House\":"
		out.RawString(prefix)
		out.String(string(in.House))
	}
	{
		const prefix string = ",\"Flat\":"
		out.RawString(prefix)
		out.Uint(uint(in.Flat))
	}
	{
		const prefix string = ",\"Cookie\":"
		out.RawString(prefix)
		out.String(string(in.Cookie))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ReqCreateAddress) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF4fdf71eEncodeServerEasy4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ReqCreateAddress) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF4fdf71eEncodeServerEasy4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ReqCreateAddress) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF4fdf71eDecodeServerEasy4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ReqCreateAddress) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF4fdf71eDecodeServerEasy4(l, v)
}
func easyjsonF4fdf71eDecodeServerEasy5(in *jlexer.Lexer, out *DBReqUpdateUserAddress) {
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
		case "UserId":
			out.UserId = uint(in.Uint())
		case "AddressId":
			out.AddressId = uint(in.Uint())
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
func easyjsonF4fdf71eEncodeServerEasy5(out *jwriter.Writer, in DBReqUpdateUserAddress) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"UserId\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.UserId))
	}
	{
		const prefix string = ",\"AddressId\":"
		out.RawString(prefix)
		out.Uint(uint(in.AddressId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DBReqUpdateUserAddress) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF4fdf71eEncodeServerEasy5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DBReqUpdateUserAddress) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF4fdf71eEncodeServerEasy5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DBReqUpdateUserAddress) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF4fdf71eDecodeServerEasy5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DBReqUpdateUserAddress) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF4fdf71eDecodeServerEasy5(l, v)
}
func easyjsonF4fdf71eDecodeServerEasy6(in *jlexer.Lexer, out *DBReqDeleteUserAddress) {
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
		case "UserId":
			out.UserId = uint(in.Uint())
		case "AddressId":
			out.AddressId = uint(in.Uint())
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
func easyjsonF4fdf71eEncodeServerEasy6(out *jwriter.Writer, in DBReqDeleteUserAddress) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"UserId\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.UserId))
	}
	{
		const prefix string = ",\"AddressId\":"
		out.RawString(prefix)
		out.Uint(uint(in.AddressId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DBReqDeleteUserAddress) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF4fdf71eEncodeServerEasy6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DBReqDeleteUserAddress) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF4fdf71eEncodeServerEasy6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DBReqDeleteUserAddress) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF4fdf71eDecodeServerEasy6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DBReqDeleteUserAddress) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF4fdf71eDecodeServerEasy6(l, v)
}
func easyjsonF4fdf71eDecodeServerEasy7(in *jlexer.Lexer, out *DBReqCreateUserAddress) {
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
		case "UserId":
			out.UserId = uint(in.Uint())
		case "City":
			out.City = string(in.String())
		case "Street":
			out.Street = string(in.String())
		case "House":
			out.House = string(in.String())
		case "Flat":
			out.Flat = uint(in.Uint())
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
func easyjsonF4fdf71eEncodeServerEasy7(out *jwriter.Writer, in DBReqCreateUserAddress) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"UserId\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.UserId))
	}
	{
		const prefix string = ",\"City\":"
		out.RawString(prefix)
		out.String(string(in.City))
	}
	{
		const prefix string = ",\"Street\":"
		out.RawString(prefix)
		out.String(string(in.Street))
	}
	{
		const prefix string = ",\"House\":"
		out.RawString(prefix)
		out.String(string(in.House))
	}
	{
		const prefix string = ",\"Flat\":"
		out.RawString(prefix)
		out.Uint(uint(in.Flat))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DBReqCreateUserAddress) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF4fdf71eEncodeServerEasy7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DBReqCreateUserAddress) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF4fdf71eEncodeServerEasy7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DBReqCreateUserAddress) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF4fdf71eDecodeServerEasy7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DBReqCreateUserAddress) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF4fdf71eDecodeServerEasy7(l, v)
}
func easyjsonF4fdf71eDecodeServerEasy8(in *jlexer.Lexer, out *DBCreateOrderAddress) {
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
		case "City":
			out.City = string(in.String())
		case "Street":
			out.Street = string(in.String())
		case "House":
			out.House = string(in.String())
		case "Flat":
			out.Flat = uint(in.Uint())
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
func easyjsonF4fdf71eEncodeServerEasy8(out *jwriter.Writer, in DBCreateOrderAddress) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"City\":"
		out.RawString(prefix[1:])
		out.String(string(in.City))
	}
	{
		const prefix string = ",\"Street\":"
		out.RawString(prefix)
		out.String(string(in.Street))
	}
	{
		const prefix string = ",\"House\":"
		out.RawString(prefix)
		out.String(string(in.House))
	}
	{
		const prefix string = ",\"Flat\":"
		out.RawString(prefix)
		out.Uint(uint(in.Flat))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DBCreateOrderAddress) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF4fdf71eEncodeServerEasy8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DBCreateOrderAddress) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF4fdf71eEncodeServerEasy8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DBCreateOrderAddress) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF4fdf71eDecodeServerEasy8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DBCreateOrderAddress) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF4fdf71eDecodeServerEasy8(l, v)
}
