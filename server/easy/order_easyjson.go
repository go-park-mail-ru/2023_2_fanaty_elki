// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package dto

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	entity "server/internal/domain/entity"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson120d1ca2DecodeServerEasy(in *jlexer.Lexer, out *RespOrders) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(RespOrders, 0, 8)
			} else {
				*out = RespOrders{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 *RespGetOrder
			if in.IsNull() {
				in.Skip()
				v1 = nil
			} else {
				if v1 == nil {
					v1 = new(RespGetOrder)
				}
				(*v1).UnmarshalEasyJSON(in)
			}
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson120d1ca2EncodeServerEasy(out *jwriter.Writer, in RespOrders) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
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

// MarshalJSON supports json.Marshaler interface
func (v RespOrders) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson120d1ca2EncodeServerEasy(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RespOrders) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson120d1ca2EncodeServerEasy(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RespOrders) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson120d1ca2DecodeServerEasy(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RespOrders) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson120d1ca2DecodeServerEasy(l, v)
}
func easyjson120d1ca2DecodeServerEasy1(in *jlexer.Lexer, out *RespGetOrder) {
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
		case "Status":
			out.Status = uint8(in.Uint8())
		case "Date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Date).UnmarshalJSON(data))
			}
		case "Address":
			if in.IsNull() {
				in.Skip()
				out.Address = nil
			} else {
				if out.Address == nil {
					out.Address = new(RespOrderAddress)
				}
				(*out.Address).UnmarshalEasyJSON(in)
			}
		case "Sum":
			out.Price = uint(in.Uint())
		case "DeliveryTime":
			out.DeliveryTime = uint8(in.Uint8())
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
func easyjson120d1ca2EncodeServerEasy1(out *jwriter.Writer, in RespGetOrder) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Id\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.Id))
	}
	{
		const prefix string = ",\"Status\":"
		out.RawString(prefix)
		out.Uint8(uint8(in.Status))
	}
	{
		const prefix string = ",\"Date\":"
		out.RawString(prefix)
		out.Raw((in.Date).MarshalJSON())
	}
	{
		const prefix string = ",\"Address\":"
		out.RawString(prefix)
		if in.Address == nil {
			out.RawString("null")
		} else {
			(*in.Address).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"Sum\":"
		out.RawString(prefix)
		out.Uint(uint(in.Price))
	}
	{
		const prefix string = ",\"DeliveryTime\":"
		out.RawString(prefix)
		out.Uint8(uint8(in.DeliveryTime))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RespGetOrder) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson120d1ca2EncodeServerEasy1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RespGetOrder) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson120d1ca2EncodeServerEasy1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RespGetOrder) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson120d1ca2DecodeServerEasy1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RespGetOrder) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson120d1ca2DecodeServerEasy1(l, v)
}
func easyjson120d1ca2DecodeServerEasy2(in *jlexer.Lexer, out *RespGetOneOrder) {
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
		case "Status":
			out.Status = uint8(in.Uint8())
		case "Date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Date).UnmarshalJSON(data))
			}
		case "Address":
			if in.IsNull() {
				in.Skip()
				out.Address = nil
			} else {
				if out.Address == nil {
					out.Address = new(RespOrderAddress)
				}
				(*out.Address).UnmarshalEasyJSON(in)
			}
		case "OrderItems":
			if in.IsNull() {
				in.Skip()
				out.OrderItems = nil
			} else {
				in.Delim('[')
				if out.OrderItems == nil {
					if !in.IsDelim(']') {
						out.OrderItems = make([]*OrderItems, 0, 8)
					} else {
						out.OrderItems = []*OrderItems{}
					}
				} else {
					out.OrderItems = (out.OrderItems)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *OrderItems
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(OrderItems)
						}
						(*v4).UnmarshalEasyJSON(in)
					}
					out.OrderItems = append(out.OrderItems, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Sum":
			out.Price = uint(in.Uint())
		case "DeliveryTime":
			out.DeliveryTime = uint8(in.Uint8())
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
func easyjson120d1ca2EncodeServerEasy2(out *jwriter.Writer, in RespGetOneOrder) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Id\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.Id))
	}
	{
		const prefix string = ",\"Status\":"
		out.RawString(prefix)
		out.Uint8(uint8(in.Status))
	}
	{
		const prefix string = ",\"Date\":"
		out.RawString(prefix)
		out.Raw((in.Date).MarshalJSON())
	}
	{
		const prefix string = ",\"Address\":"
		out.RawString(prefix)
		if in.Address == nil {
			out.RawString("null")
		} else {
			(*in.Address).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"OrderItems\":"
		out.RawString(prefix)
		if in.OrderItems == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.OrderItems {
				if v5 > 0 {
					out.RawByte(',')
				}
				if v6 == nil {
					out.RawString("null")
				} else {
					(*v6).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"Sum\":"
		out.RawString(prefix)
		out.Uint(uint(in.Price))
	}
	{
		const prefix string = ",\"DeliveryTime\":"
		out.RawString(prefix)
		out.Uint8(uint8(in.DeliveryTime))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RespGetOneOrder) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson120d1ca2EncodeServerEasy2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RespGetOneOrder) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson120d1ca2EncodeServerEasy2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RespGetOneOrder) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson120d1ca2DecodeServerEasy2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RespGetOneOrder) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson120d1ca2DecodeServerEasy2(l, v)
}
func easyjson120d1ca2DecodeServerEasy3(in *jlexer.Lexer, out *RespCreateOrder) {
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
		case "Status":
			out.Status = uint8(in.Uint8())
		case "Date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Date).UnmarshalJSON(data))
			}
		case "Address":
			if in.IsNull() {
				in.Skip()
				out.Address = nil
			} else {
				if out.Address == nil {
					out.Address = new(entity.Address)
				}
				easyjson120d1ca2DecodeServerInternalDomainEntity(in, out.Address)
			}
		case "Sum":
			out.Price = uint(in.Uint())
		case "DeliveryTime":
			out.DeliveryTime = uint8(in.Uint8())
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
func easyjson120d1ca2EncodeServerEasy3(out *jwriter.Writer, in RespCreateOrder) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Id\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.Id))
	}
	{
		const prefix string = ",\"Status\":"
		out.RawString(prefix)
		out.Uint8(uint8(in.Status))
	}
	{
		const prefix string = ",\"Date\":"
		out.RawString(prefix)
		out.Raw((in.Date).MarshalJSON())
	}
	{
		const prefix string = ",\"Address\":"
		out.RawString(prefix)
		if in.Address == nil {
			out.RawString("null")
		} else {
			easyjson120d1ca2EncodeServerInternalDomainEntity(out, *in.Address)
		}
	}
	{
		const prefix string = ",\"Sum\":"
		out.RawString(prefix)
		out.Uint(uint(in.Price))
	}
	{
		const prefix string = ",\"DeliveryTime\":"
		out.RawString(prefix)
		out.Uint8(uint8(in.DeliveryTime))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RespCreateOrder) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson120d1ca2EncodeServerEasy3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RespCreateOrder) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson120d1ca2EncodeServerEasy3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RespCreateOrder) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson120d1ca2DecodeServerEasy3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RespCreateOrder) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson120d1ca2DecodeServerEasy3(l, v)
}
func easyjson120d1ca2DecodeServerInternalDomainEntity(in *jlexer.Lexer, out *entity.Address) {
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
func easyjson120d1ca2EncodeServerInternalDomainEntity(out *jwriter.Writer, in entity.Address) {
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
func easyjson120d1ca2DecodeServerEasy4(in *jlexer.Lexer, out *ReqUpdateOrder) {
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
		case "Status":
			out.Status = uint8(in.Uint8())
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
func easyjson120d1ca2EncodeServerEasy4(out *jwriter.Writer, in ReqUpdateOrder) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Id\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.Id))
	}
	{
		const prefix string = ",\"Status\":"
		out.RawString(prefix)
		out.Uint8(uint8(in.Status))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ReqUpdateOrder) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson120d1ca2EncodeServerEasy4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ReqUpdateOrder) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson120d1ca2EncodeServerEasy4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ReqUpdateOrder) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson120d1ca2DecodeServerEasy4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ReqUpdateOrder) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson120d1ca2DecodeServerEasy4(l, v)
}
func easyjson120d1ca2DecodeServerEasy5(in *jlexer.Lexer, out *ReqGetOneOrder) {
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
		case "OrderId":
			out.OrderId = uint(in.Uint())
		case "UserId":
			out.UserId = uint(in.Uint())
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
func easyjson120d1ca2EncodeServerEasy5(out *jwriter.Writer, in ReqGetOneOrder) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"OrderId\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.OrderId))
	}
	{
		const prefix string = ",\"UserId\":"
		out.RawString(prefix)
		out.Uint(uint(in.UserId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ReqGetOneOrder) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson120d1ca2EncodeServerEasy5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ReqGetOneOrder) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson120d1ca2EncodeServerEasy5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ReqGetOneOrder) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson120d1ca2DecodeServerEasy5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ReqGetOneOrder) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson120d1ca2DecodeServerEasy5(l, v)
}
func easyjson120d1ca2DecodeServerEasy6(in *jlexer.Lexer, out *ReqCreateOrder) {
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
		case "UserID":
			out.UserId = uint(in.Uint())
		case "Address":
			if in.IsNull() {
				in.Skip()
				out.Address = nil
			} else {
				if out.Address == nil {
					out.Address = new(ReqCreateOrderAddress)
				}
				(*out.Address).UnmarshalEasyJSON(in)
			}
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
func easyjson120d1ca2EncodeServerEasy6(out *jwriter.Writer, in ReqCreateOrder) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"UserID\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.UserId))
	}
	{
		const prefix string = ",\"Address\":"
		out.RawString(prefix)
		if in.Address == nil {
			out.RawString("null")
		} else {
			(*in.Address).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ReqCreateOrder) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson120d1ca2EncodeServerEasy6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ReqCreateOrder) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson120d1ca2EncodeServerEasy6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ReqCreateOrder) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson120d1ca2DecodeServerEasy6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ReqCreateOrder) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson120d1ca2DecodeServerEasy6(l, v)
}
func easyjson120d1ca2DecodeServerEasy7(in *jlexer.Lexer, out *OrderItems) {
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
		case "RestaurantName":
			out.RestaurantName = string(in.String())
		case "Products":
			if in.IsNull() {
				in.Skip()
				out.Products = nil
			} else {
				in.Delim('[')
				if out.Products == nil {
					if !in.IsDelim(']') {
						out.Products = make([]*RespGetOrderProduct, 0, 8)
					} else {
						out.Products = []*RespGetOrderProduct{}
					}
				} else {
					out.Products = (out.Products)[:0]
				}
				for !in.IsDelim(']') {
					var v7 *RespGetOrderProduct
					if in.IsNull() {
						in.Skip()
						v7 = nil
					} else {
						if v7 == nil {
							v7 = new(RespGetOrderProduct)
						}
						easyjson120d1ca2DecodeServerEasy8(in, v7)
					}
					out.Products = append(out.Products, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson120d1ca2EncodeServerEasy7(out *jwriter.Writer, in OrderItems) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"RestaurantName\":"
		out.RawString(prefix[1:])
		out.String(string(in.RestaurantName))
	}
	{
		const prefix string = ",\"Products\":"
		out.RawString(prefix)
		if in.Products == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.Products {
				if v8 > 0 {
					out.RawByte(',')
				}
				if v9 == nil {
					out.RawString("null")
				} else {
					easyjson120d1ca2EncodeServerEasy8(out, *v9)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v OrderItems) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson120d1ca2EncodeServerEasy7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v OrderItems) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson120d1ca2EncodeServerEasy7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *OrderItems) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson120d1ca2DecodeServerEasy7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *OrderItems) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson120d1ca2DecodeServerEasy7(l, v)
}
func easyjson120d1ca2DecodeServerEasy8(in *jlexer.Lexer, out *RespGetOrderProduct) {
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
		case "Name":
			out.Name = string(in.String())
		case "Price":
			out.Price = float32(in.Float32())
		case "Icon":
			out.Icon = string(in.String())
		case "Count":
			out.Count = int(in.Int())
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
func easyjson120d1ca2EncodeServerEasy8(out *jwriter.Writer, in RespGetOrderProduct) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Id\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.Id))
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
		const prefix string = ",\"Icon\":"
		out.RawString(prefix)
		out.String(string(in.Icon))
	}
	{
		const prefix string = ",\"Count\":"
		out.RawString(prefix)
		out.Int(int(in.Count))
	}
	out.RawByte('}')
}
func easyjson120d1ca2DecodeServerEasy9(in *jlexer.Lexer, out *DBReqCreateOrder) {
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
		case "Products":
			if in.IsNull() {
				in.Skip()
				out.Products = nil
			} else {
				in.Delim('[')
				if out.Products == nil {
					if !in.IsDelim(']') {
						out.Products = make([]*entity.CartProduct, 0, 8)
					} else {
						out.Products = []*entity.CartProduct{}
					}
				} else {
					out.Products = (out.Products)[:0]
				}
				for !in.IsDelim(']') {
					var v10 *entity.CartProduct
					if in.IsNull() {
						in.Skip()
						v10 = nil
					} else {
						if v10 == nil {
							v10 = new(entity.CartProduct)
						}
						easyjson120d1ca2DecodeServerInternalDomainEntity1(in, v10)
					}
					out.Products = append(out.Products, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "UserId":
			out.UserId = uint(in.Uint())
		case "Status":
			out.Status = uint8(in.Uint8())
		case "Price":
			out.Price = uint(in.Uint())
		case "Date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Date).UnmarshalJSON(data))
			}
		case "Address":
			if in.IsNull() {
				in.Skip()
				out.Address = nil
			} else {
				if out.Address == nil {
					out.Address = new(DBCreateOrderAddress)
				}
				(*out.Address).UnmarshalEasyJSON(in)
			}
		case "DeliveryTime":
			out.DeliveryTime = uint8(in.Uint8())
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
func easyjson120d1ca2EncodeServerEasy9(out *jwriter.Writer, in DBReqCreateOrder) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Products\":"
		out.RawString(prefix[1:])
		if in.Products == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.Products {
				if v11 > 0 {
					out.RawByte(',')
				}
				if v12 == nil {
					out.RawString("null")
				} else {
					easyjson120d1ca2EncodeServerInternalDomainEntity1(out, *v12)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"UserId\":"
		out.RawString(prefix)
		out.Uint(uint(in.UserId))
	}
	{
		const prefix string = ",\"Status\":"
		out.RawString(prefix)
		out.Uint8(uint8(in.Status))
	}
	{
		const prefix string = ",\"Price\":"
		out.RawString(prefix)
		out.Uint(uint(in.Price))
	}
	{
		const prefix string = ",\"Date\":"
		out.RawString(prefix)
		out.Raw((in.Date).MarshalJSON())
	}
	{
		const prefix string = ",\"Address\":"
		out.RawString(prefix)
		if in.Address == nil {
			out.RawString("null")
		} else {
			(*in.Address).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"DeliveryTime\":"
		out.RawString(prefix)
		out.Uint8(uint8(in.DeliveryTime))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DBReqCreateOrder) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson120d1ca2EncodeServerEasy9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DBReqCreateOrder) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson120d1ca2EncodeServerEasy9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DBReqCreateOrder) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson120d1ca2DecodeServerEasy9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DBReqCreateOrder) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson120d1ca2DecodeServerEasy9(l, v)
}
func easyjson120d1ca2DecodeServerInternalDomainEntity1(in *jlexer.Lexer, out *entity.CartProduct) {
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
		case "ProductID":
			out.ProductID = uint(in.Uint())
		case "CartID":
			out.CartID = uint(in.Uint())
		case "ItemCount":
			out.ItemCount = int(in.Int())
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
func easyjson120d1ca2EncodeServerInternalDomainEntity1(out *jwriter.Writer, in entity.CartProduct) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ID\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.ID))
	}
	{
		const prefix string = ",\"ProductID\":"
		out.RawString(prefix)
		out.Uint(uint(in.ProductID))
	}
	{
		const prefix string = ",\"CartID\":"
		out.RawString(prefix)
		out.Uint(uint(in.CartID))
	}
	{
		const prefix string = ",\"ItemCount\":"
		out.RawString(prefix)
		out.Int(int(in.ItemCount))
	}
	out.RawByte('}')
}
