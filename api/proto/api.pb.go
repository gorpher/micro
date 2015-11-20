// Code generated by protoc-gen-go.
// source: micro/micro/api/proto/api.proto
// DO NOT EDIT!

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	micro/micro/api/proto/api.proto

It has these top-level messages:
	Pair
	Request
	Response
*/
package api

import proto "github.com/golang/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type Pair struct {
	Key    string   `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Values []string `protobuf:"bytes,2,rep,name=values" json:"values,omitempty"`
}

func (m *Pair) Reset()         { *m = Pair{} }
func (m *Pair) String() string { return proto.CompactTextString(m) }
func (*Pair) ProtoMessage()    {}

type Request struct {
	Method string           `protobuf:"bytes,1,opt,name=method" json:"method,omitempty"`
	Path   string           `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Header map[string]*Pair `protobuf:"bytes,3,rep,name=header" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Get    map[string]*Pair `protobuf:"bytes,4,rep,name=get" json:"get,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Post   map[string]*Pair `protobuf:"bytes,5,rep,name=post" json:"post,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Body   string           `protobuf:"bytes,6,opt,name=body" json:"body,omitempty"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

func (m *Request) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Request) GetGet() map[string]*Pair {
	if m != nil {
		return m.Get
	}
	return nil
}

func (m *Request) GetPost() map[string]*Pair {
	if m != nil {
		return m.Post
	}
	return nil
}

type Response struct {
	StatusCode int32            `protobuf:"varint,1,opt,name=statusCode" json:"statusCode,omitempty"`
	Header     map[string]*Pair `protobuf:"bytes,2,rep,name=header" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Body       string           `protobuf:"bytes,3,opt,name=body" json:"body,omitempty"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

func (m *Response) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func init() {
}
