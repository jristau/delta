// Code generated by protoc-gen-go.
// source: delta.proto
// DO NOT EDIT!

/*
Package delta is a generated protocol buffer package.

It is generated from these files:
	delta.proto

It has these top-level messages:
	Point
	Span
	Network
	Mark
	Marks
	Station
	Stations
*/
package delta

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Point struct {
	// Latitude - Geographical latitude of the point for the given datum.
	Latitude float64 `protobuf:"fixed64,1,opt,name=latitude" json:"latitude,omitempty"`
	// Longitude - Geographical longitude of the point for the given datum.
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude" json:"longitude,omitempty"`
	// Elevation - Height in meters of the point for the given datum.
	Elevation float64 `protobuf:"fixed64,3,opt,name=elevation" json:"elevation,omitempty"`
	// Datum - Geographical reference system used for the latitude, longitude & elevation.
	Datum string `protobuf:"bytes,4,opt,name=datum" json:"datum,omitempty"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Span struct {
	// Start - time in Unix seconds.
	Start int64 `protobuf:"varint,1,opt,name=start" json:"start,omitempty"`
	// End - time in Unix seconds.  TODO comments about 9999 or method.
	End int64 `protobuf:"varint,2,opt,name=end" json:"end,omitempty"`
}

func (m *Span) Reset()                    { *m = Span{} }
func (m *Span) String() string            { return proto.CompactTextString(m) }
func (*Span) ProtoMessage()               {}
func (*Span) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Network struct {
	Code        string `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	External    string `protobuf:"bytes,2,opt,name=external" json:"external,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	Restricted  bool   `protobuf:"varint,4,opt,name=restricted" json:"restricted,omitempty"`
}

func (m *Network) Reset()                    { *m = Network{} }
func (m *Network) String() string            { return proto.CompactTextString(m) }
func (*Network) ProtoMessage()               {}
func (*Network) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Mark struct {
	// Code used to uniquely identify GNSS Mark.
	Code    string   `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	Name    string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Network *Network `protobuf:"bytes,3,opt,name=network" json:"network,omitempty"`
	Point   *Point   `protobuf:"bytes,4,opt,name=point" json:"point,omitempty"`
	Span    *Span    `protobuf:"bytes,5,opt,name=span" json:"span,omitempty"`
}

func (m *Mark) Reset()                    { *m = Mark{} }
func (m *Mark) String() string            { return proto.CompactTextString(m) }
func (*Mark) ProtoMessage()               {}
func (*Mark) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Mark) GetNetwork() *Network {
	if m != nil {
		return m.Network
	}
	return nil
}

func (m *Mark) GetPoint() *Point {
	if m != nil {
		return m.Point
	}
	return nil
}

func (m *Mark) GetSpan() *Span {
	if m != nil {
		return m.Span
	}
	return nil
}

type Marks struct {
	Marks map[string]*Mark `protobuf:"bytes,1,rep,name=marks" json:"marks,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Marks) Reset()                    { *m = Marks{} }
func (m *Marks) String() string            { return proto.CompactTextString(m) }
func (*Marks) ProtoMessage()               {}
func (*Marks) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Marks) GetMarks() map[string]*Mark {
	if m != nil {
		return m.Marks
	}
	return nil
}

type Station struct {
	// Code used to uniquely identify the Station.
	Code string `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	// Name used to describe the general geographical location of the Station.
	Name    string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Network *Network `protobuf:"bytes,3,opt,name=network" json:"network,omitempty"`
	Point   *Point   `protobuf:"bytes,4,opt,name=point" json:"point,omitempty"`
	Span    *Span    `protobuf:"bytes,5,opt,name=span" json:"span,omitempty"`
}

func (m *Station) Reset()                    { *m = Station{} }
func (m *Station) String() string            { return proto.CompactTextString(m) }
func (*Station) ProtoMessage()               {}
func (*Station) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Station) GetNetwork() *Network {
	if m != nil {
		return m.Network
	}
	return nil
}

func (m *Station) GetPoint() *Point {
	if m != nil {
		return m.Point
	}
	return nil
}

func (m *Station) GetSpan() *Span {
	if m != nil {
		return m.Span
	}
	return nil
}

type Stations struct {
	Stations map[string]*Station `protobuf:"bytes,1,rep,name=stations" json:"stations,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Stations) Reset()                    { *m = Stations{} }
func (m *Stations) String() string            { return proto.CompactTextString(m) }
func (*Stations) ProtoMessage()               {}
func (*Stations) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Stations) GetStations() map[string]*Station {
	if m != nil {
		return m.Stations
	}
	return nil
}

func init() {
	proto.RegisterType((*Point)(nil), "delta.Point")
	proto.RegisterType((*Span)(nil), "delta.Span")
	proto.RegisterType((*Network)(nil), "delta.Network")
	proto.RegisterType((*Mark)(nil), "delta.Mark")
	proto.RegisterType((*Marks)(nil), "delta.Marks")
	proto.RegisterType((*Station)(nil), "delta.Station")
	proto.RegisterType((*Stations)(nil), "delta.Stations")
}

func init() { proto.RegisterFile("delta.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x53, 0x4d, 0x6b, 0xe3, 0x30,
	0x10, 0xc5, 0xb1, 0xb5, 0xb6, 0xc7, 0xbb, 0x61, 0x11, 0x0b, 0x6b, 0xc2, 0x7e, 0x64, 0xcd, 0x1e,
	0x72, 0xd9, 0x1c, 0xb2, 0x97, 0xdd, 0x1e, 0x0b, 0x39, 0x95, 0x96, 0xa2, 0xdc, 0x7a, 0x53, 0x63,
	0x51, 0x4c, 0x1c, 0xd9, 0xc8, 0x72, 0xda, 0x50, 0xfa, 0x27, 0x7a, 0x2f, 0xf4, 0xa7, 0xd6, 0x1a,
	0x29, 0x4e, 0x02, 0xed, 0xbd, 0x17, 0x33, 0x33, 0xef, 0xc9, 0xef, 0xcd, 0x13, 0x82, 0x24, 0x17,
	0xa5, 0xe6, 0xd3, 0x5a, 0x55, 0xba, 0xa2, 0x04, 0x9b, 0xac, 0x05, 0x72, 0x59, 0x15, 0x52, 0xd3,
	0x11, 0x44, 0x25, 0xd7, 0x85, 0x6e, 0x73, 0x91, 0x7a, 0x63, 0x6f, 0xe2, 0xb1, 0xbe, 0xa7, 0xdf,
	0x20, 0x2e, 0x2b, 0x79, 0x63, 0xc1, 0x01, 0x82, 0xfb, 0x81, 0x41, 0x45, 0x29, 0x36, 0x1d, 0xb9,
	0x92, 0xa9, 0x6f, 0xd1, 0x7e, 0x40, 0xbf, 0x00, 0xc9, 0xb9, 0x6e, 0xd7, 0x69, 0xd0, 0x21, 0x31,
	0xb3, 0x4d, 0x36, 0x85, 0x60, 0x51, 0x73, 0x44, 0x1b, 0xcd, 0x95, 0x46, 0x49, 0x9f, 0xd9, 0x86,
	0x7e, 0x06, 0x5f, 0xc8, 0x1c, 0x95, 0x7c, 0x66, 0xca, 0xec, 0x1e, 0xc2, 0x0b, 0xa1, 0x6f, 0x2b,
	0xb5, 0xa2, 0x14, 0x82, 0x65, 0xe5, 0x4c, 0xc6, 0x0c, 0x6b, 0x63, 0x5e, 0xdc, 0x69, 0xa1, 0x24,
	0x2f, 0xf1, 0x54, 0xcc, 0xfa, 0x9e, 0x8e, 0xcd, 0xde, 0xcd, 0x52, 0x15, 0x75, 0x6f, 0x30, 0x66,
	0x87, 0x23, 0xfa, 0x03, 0x40, 0x89, 0x46, 0xab, 0x62, 0xa9, 0x45, 0x8e, 0x3e, 0x23, 0x76, 0x30,
	0xc9, 0x9e, 0x3c, 0x08, 0xce, 0xf9, 0x1b, 0xd2, 0xdd, 0x4c, 0xf2, 0xb5, 0x70, 0xb2, 0x58, 0xd3,
	0x09, 0x84, 0xd2, 0xba, 0x45, 0xb9, 0x64, 0x36, 0x9c, 0xda, 0xe8, 0xdd, 0x0e, 0x6c, 0x07, 0xd3,
	0x0c, 0x48, 0x6d, 0xe2, 0x47, 0xd5, 0x64, 0xf6, 0xd1, 0xf1, 0xf0, 0x4a, 0x98, 0x85, 0xe8, 0x4f,
	0x08, 0x9a, 0x2e, 0xab, 0x94, 0x20, 0x25, 0x71, 0x14, 0x13, 0x1f, 0x43, 0x20, 0x7b, 0x00, 0x62,
	0xec, 0x35, 0xf4, 0x0f, 0x90, 0xb5, 0x29, 0x3a, 0x83, 0x7e, 0x47, 0xfd, 0xea, 0xa8, 0x08, 0xda,
	0xef, 0x5c, 0x6a, 0xb5, 0x65, 0x96, 0x35, 0x9a, 0x03, 0xec, 0x87, 0x26, 0xf4, 0x95, 0xd8, 0xba,
	0xdd, 0x4c, 0x49, 0x7f, 0x01, 0xd9, 0xf0, 0xb2, 0xb5, 0xbb, 0xed, 0x95, 0xcd, 0x19, 0x66, 0x91,
	0x93, 0xc1, 0x3f, 0x2f, 0x7b, 0xf6, 0x20, 0x5c, 0x68, 0x7b, 0xdb, 0xef, 0x34, 0xa1, 0x47, 0x0f,
	0x22, 0x67, 0xb1, 0xa1, 0xff, 0x21, 0x6a, 0x5c, 0xed, 0x82, 0xfa, 0xbe, 0x3b, 0xe1, 0xc6, 0x7d,
	0x61, 0xe3, 0xea, 0xe9, 0xa3, 0x33, 0xf8, 0x74, 0x04, 0xbd, 0x12, 0xda, 0xef, 0xe3, 0xd0, 0x86,
	0xc7, 0xbf, 0x3e, 0xc8, 0xed, 0x34, 0xbc, 0xb2, 0x6f, 0xf0, 0xfa, 0x03, 0xbe, 0xc8, 0xbf, 0x2f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x84, 0xc5, 0x9c, 0x04, 0xa0, 0x03, 0x00, 0x00,
}