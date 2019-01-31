// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

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

func easyjson5b12d0fdDecodeGithubComMhc18GoModels(in *jlexer.Lexer, out *Like) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = int16(in.Int16())
		case "ts":
			out.Ts = int32(in.Int32())
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
func easyjson5b12d0fdEncodeGithubComMhc18GoModels(out *jwriter.Writer, in Like) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Id != 0 {
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int16(int16(in.Id))
	}
	if in.Ts != 0 {
		const prefix string = ",\"ts\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Ts))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Like) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5b12d0fdEncodeGithubComMhc18GoModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Like) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5b12d0fdEncodeGithubComMhc18GoModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Like) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5b12d0fdDecodeGithubComMhc18GoModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Like) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5b12d0fdDecodeGithubComMhc18GoModels(l, v)
}
func easyjson5b12d0fdDecodeGithubComMhc18GoModels1(in *jlexer.Lexer, out *AccountFormEntry) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			if in.IsNull() {
				in.Skip()
				out.Id = nil
			} else {
				if out.Id == nil {
					out.Id = new(uint16)
				}
				*out.Id = uint16(in.Uint16())
			}
		case "email":
			if in.IsNull() {
				in.Skip()
				out.Email = nil
			} else {
				if out.Email == nil {
					out.Email = new(string)
				}
				*out.Email = string(in.String())
			}
		case "fname":
			out.Fname = string(in.String())
		case "sname":
			out.Sname = string(in.String())
		case "phone":
			out.Phone = string(in.String())
		case "sex":
			if in.IsNull() {
				in.Skip()
				out.Sex = nil
			} else {
				if out.Sex == nil {
					out.Sex = new(string)
				}
				*out.Sex = string(in.String())
			}
		case "birth":
			if in.IsNull() {
				in.Skip()
				out.Birth = nil
			} else {
				if out.Birth == nil {
					out.Birth = new(uint32)
				}
				*out.Birth = uint32(in.Uint32())
			}
		case "country":
			out.Country = string(in.String())
		case "city":
			out.City = string(in.String())
		case "joined":
			if in.IsNull() {
				in.Skip()
				out.Joined = nil
			} else {
				if out.Joined == nil {
					out.Joined = new(uint32)
				}
				*out.Joined = uint32(in.Uint32())
			}
		case "status":
			if in.IsNull() {
				in.Skip()
				out.Status = nil
			} else {
				if out.Status == nil {
					out.Status = new(string)
				}
				*out.Status = string(in.String())
			}
		case "interests":
			if in.IsNull() {
				in.Skip()
				out.Interests = nil
			} else {
				if out.Interests == nil {
					out.Interests = new([]string)
				}
				if in.IsNull() {
					in.Skip()
					*out.Interests = nil
				} else {
					in.Delim('[')
					if *out.Interests == nil {
						if !in.IsDelim(']') {
							*out.Interests = make([]string, 0, 4)
						} else {
							*out.Interests = []string{}
						}
					} else {
						*out.Interests = (*out.Interests)[:0]
					}
					for !in.IsDelim(']') {
						var v1 string
						v1 = string(in.String())
						*out.Interests = append(*out.Interests, v1)
						in.WantComma()
					}
					in.Delim(']')
				}
			}
		case "premium":
			if in.IsNull() {
				in.Skip()
				out.Premium = nil
			} else {
				if out.Premium == nil {
					out.Premium = new(struct {
						Start  *uint32
						Finish *uint32
					})
				}
				easyjson5b12d0fdDecode(in, &*out.Premium)
			}
		case "likes":
			if in.IsNull() {
				in.Skip()
				out.Likes = nil
			} else {
				if out.Likes == nil {
					out.Likes = new([]*Like)
				}
				if in.IsNull() {
					in.Skip()
					*out.Likes = nil
				} else {
					in.Delim('[')
					if *out.Likes == nil {
						if !in.IsDelim(']') {
							*out.Likes = make([]*Like, 0, 8)
						} else {
							*out.Likes = []*Like{}
						}
					} else {
						*out.Likes = (*out.Likes)[:0]
					}
					for !in.IsDelim(']') {
						var v2 *Like
						if in.IsNull() {
							in.Skip()
							v2 = nil
						} else {
							if v2 == nil {
								v2 = new(Like)
							}
							(*v2).UnmarshalEasyJSON(in)
						}
						*out.Likes = append(*out.Likes, v2)
						in.WantComma()
					}
					in.Delim(']')
				}
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
func easyjson5b12d0fdEncodeGithubComMhc18GoModels1(out *jwriter.Writer, in AccountFormEntry) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Id != nil {
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint16(uint16(*in.Id))
	}
	if in.Email != nil {
		const prefix string = ",\"email\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Email))
	}
	if in.Fname != "" {
		const prefix string = ",\"fname\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Fname))
	}
	if in.Sname != "" {
		const prefix string = ",\"sname\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Sname))
	}
	if in.Phone != "" {
		const prefix string = ",\"phone\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Phone))
	}
	if in.Sex != nil {
		const prefix string = ",\"sex\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Sex))
	}
	if in.Birth != nil {
		const prefix string = ",\"birth\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint32(uint32(*in.Birth))
	}
	if in.Country != "" {
		const prefix string = ",\"country\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Country))
	}
	if in.City != "" {
		const prefix string = ",\"city\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.City))
	}
	if in.Joined != nil {
		const prefix string = ",\"joined\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint32(uint32(*in.Joined))
	}
	if in.Status != nil {
		const prefix string = ",\"status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Status))
	}
	if in.Interests != nil {
		const prefix string = ",\"interests\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if *in.Interests == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v3, v4 := range *in.Interests {
				if v3 > 0 {
					out.RawByte(',')
				}
				out.String(string(v4))
			}
			out.RawByte(']')
		}
	}
	if in.Premium != nil {
		const prefix string = ",\"premium\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson5b12d0fdEncode(out, *in.Premium)
	}
	if in.Likes != nil {
		const prefix string = ",\"likes\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if *in.Likes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range *in.Likes {
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
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AccountFormEntry) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5b12d0fdEncodeGithubComMhc18GoModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AccountFormEntry) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5b12d0fdEncodeGithubComMhc18GoModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AccountFormEntry) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5b12d0fdDecodeGithubComMhc18GoModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AccountFormEntry) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5b12d0fdDecodeGithubComMhc18GoModels1(l, v)
}
func easyjson5b12d0fdDecode(in *jlexer.Lexer, out *struct {
	Start  *uint32
	Finish *uint32
}) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "start":
			if in.IsNull() {
				in.Skip()
				out.Start = nil
			} else {
				if out.Start == nil {
					out.Start = new(uint32)
				}
				*out.Start = uint32(in.Uint32())
			}
		case "finish":
			if in.IsNull() {
				in.Skip()
				out.Finish = nil
			} else {
				if out.Finish == nil {
					out.Finish = new(uint32)
				}
				*out.Finish = uint32(in.Uint32())
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
func easyjson5b12d0fdEncode(out *jwriter.Writer, in struct {
	Start  *uint32
	Finish *uint32
}) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Start != nil {
		const prefix string = ",\"start\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint32(uint32(*in.Start))
	}
	if in.Finish != nil {
		const prefix string = ",\"finish\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint32(uint32(*in.Finish))
	}
	out.RawByte('}')
}
