// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/common/value.proto

package common // import "google.golang.org/genproto/googleapis/ads/googleads/v1/common"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A generic data container.
type Value struct {
	// A value.
	//
	// Types that are valid to be assigned to Value:
	//	*Value_BooleanValue
	//	*Value_Int64Value
	//	*Value_FloatValue
	//	*Value_DoubleValue
	//	*Value_StringValue
	Value                isValue_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_value_1cad930ee210ecc3, []int{0}
}
func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (dst *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(dst, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

type isValue_Value interface {
	isValue_Value()
}

type Value_BooleanValue struct {
	BooleanValue bool `protobuf:"varint,1,opt,name=boolean_value,json=booleanValue,proto3,oneof"`
}

type Value_Int64Value struct {
	Int64Value int64 `protobuf:"varint,2,opt,name=int64_value,json=int64Value,proto3,oneof"`
}

type Value_FloatValue struct {
	FloatValue float32 `protobuf:"fixed32,3,opt,name=float_value,json=floatValue,proto3,oneof"`
}

type Value_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,4,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type Value_StringValue struct {
	StringValue string `protobuf:"bytes,5,opt,name=string_value,json=stringValue,proto3,oneof"`
}

func (*Value_BooleanValue) isValue_Value() {}

func (*Value_Int64Value) isValue_Value() {}

func (*Value_FloatValue) isValue_Value() {}

func (*Value_DoubleValue) isValue_Value() {}

func (*Value_StringValue) isValue_Value() {}

func (m *Value) GetValue() isValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Value) GetBooleanValue() bool {
	if x, ok := m.GetValue().(*Value_BooleanValue); ok {
		return x.BooleanValue
	}
	return false
}

func (m *Value) GetInt64Value() int64 {
	if x, ok := m.GetValue().(*Value_Int64Value); ok {
		return x.Int64Value
	}
	return 0
}

func (m *Value) GetFloatValue() float32 {
	if x, ok := m.GetValue().(*Value_FloatValue); ok {
		return x.FloatValue
	}
	return 0
}

func (m *Value) GetDoubleValue() float64 {
	if x, ok := m.GetValue().(*Value_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (m *Value) GetStringValue() string {
	if x, ok := m.GetValue().(*Value_StringValue); ok {
		return x.StringValue
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Value) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Value_OneofMarshaler, _Value_OneofUnmarshaler, _Value_OneofSizer, []interface{}{
		(*Value_BooleanValue)(nil),
		(*Value_Int64Value)(nil),
		(*Value_FloatValue)(nil),
		(*Value_DoubleValue)(nil),
		(*Value_StringValue)(nil),
	}
}

func _Value_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Value)
	// value
	switch x := m.Value.(type) {
	case *Value_BooleanValue:
		t := uint64(0)
		if x.BooleanValue {
			t = 1
		}
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *Value_Int64Value:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Int64Value))
	case *Value_FloatValue:
		b.EncodeVarint(3<<3 | proto.WireFixed32)
		b.EncodeFixed32(uint64(math.Float32bits(x.FloatValue)))
	case *Value_DoubleValue:
		b.EncodeVarint(4<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.DoubleValue))
	case *Value_StringValue:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.StringValue)
	case nil:
	default:
		return fmt.Errorf("Value.Value has unexpected type %T", x)
	}
	return nil
}

func _Value_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Value)
	switch tag {
	case 1: // value.boolean_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &Value_BooleanValue{x != 0}
		return true, err
	case 2: // value.int64_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &Value_Int64Value{int64(x)}
		return true, err
	case 3: // value.float_value
		if wire != proto.WireFixed32 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed32()
		m.Value = &Value_FloatValue{math.Float32frombits(uint32(x))}
		return true, err
	case 4: // value.double_value
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Value = &Value_DoubleValue{math.Float64frombits(x)}
		return true, err
	case 5: // value.string_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &Value_StringValue{x}
		return true, err
	default:
		return false, nil
	}
}

func _Value_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Value)
	// value
	switch x := m.Value.(type) {
	case *Value_BooleanValue:
		n += 1 // tag and wire
		n += 1
	case *Value_Int64Value:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Int64Value))
	case *Value_FloatValue:
		n += 1 // tag and wire
		n += 4
	case *Value_DoubleValue:
		n += 1 // tag and wire
		n += 8
	case *Value_StringValue:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.StringValue)))
		n += len(x.StringValue)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Value)(nil), "google.ads.googleads.v1.common.Value")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/common/value.proto", fileDescriptor_value_1cad930ee210ecc3)
}

var fileDescriptor_value_1cad930ee210ecc3 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0xd1, 0x41, 0x4b, 0xf3, 0x30,
	0x18, 0x07, 0xf0, 0xa5, 0x7b, 0xf7, 0xaa, 0xd9, 0xbc, 0xec, 0x24, 0x22, 0xa3, 0x4e, 0x84, 0xe2,
	0x21, 0xa5, 0x28, 0x1e, 0xe2, 0xa9, 0x53, 0xd8, 0x8e, 0x63, 0x87, 0x1e, 0xa4, 0x20, 0xd9, 0x5a,
	0x43, 0x21, 0xcb, 0x33, 0x96, 0x6c, 0x1f, 0xc8, 0xa3, 0x5f, 0x44, 0xf0, 0x7b, 0x78, 0xf1, 0x53,
	0x48, 0xf2, 0x64, 0xbd, 0xe9, 0xa9, 0x0f, 0xff, 0xfe, 0xf2, 0xa7, 0x7d, 0x42, 0x6f, 0x24, 0x80,
	0x54, 0x75, 0x2a, 0x2a, 0x93, 0xe2, 0xe8, 0xa6, 0x7d, 0x96, 0xae, 0x60, 0xbd, 0x06, 0x9d, 0xee,
	0x85, 0xda, 0xd5, 0x6c, 0xb3, 0x05, 0x0b, 0xc3, 0x11, 0x02, 0x26, 0x2a, 0xc3, 0x5a, 0xcb, 0xf6,
	0x19, 0x43, 0x7b, 0x7e, 0x71, 0xe8, 0xda, 0x34, 0xa9, 0xd0, 0x1a, 0xac, 0xb0, 0x0d, 0x68, 0x83,
	0xa7, 0xc7, 0x1f, 0x84, 0xf6, 0x0a, 0xd7, 0x36, 0xbc, 0xa6, 0xa7, 0x4b, 0x00, 0x55, 0x0b, 0xfd,
	0xe2, 0xeb, 0xcf, 0x48, 0x4c, 0x92, 0xe3, 0x59, 0x67, 0x31, 0x08, 0x31, 0xb2, 0x4b, 0xda, 0x6f,
	0xb4, 0xbd, 0xbf, 0x0b, 0x28, 0x8a, 0x49, 0xd2, 0x9d, 0x75, 0x16, 0xd4, 0x87, 0x2d, 0x79, 0x55,
	0x20, 0x6c, 0x20, 0xdd, 0x98, 0x24, 0x91, 0x23, 0x3e, 0x44, 0x72, 0x45, 0x07, 0x15, 0xec, 0x96,
	0xaa, 0x0e, 0xe6, 0x5f, 0x4c, 0x12, 0x32, 0xeb, 0x2c, 0xfa, 0x98, 0xb6, 0xc8, 0xd8, 0x6d, 0xa3,
	0x65, 0x40, 0xbd, 0x98, 0x24, 0x27, 0x0e, 0x61, 0xea, 0xd1, 0xe4, 0x88, 0xf6, 0xfc, 0xdb, 0xc9,
	0x17, 0xa1, 0xe3, 0x15, 0xac, 0xd9, 0xdf, 0xeb, 0x98, 0x50, 0x7f, 0x6c, 0xee, 0x7e, 0x7e, 0x4e,
	0x9e, 0x9f, 0x82, 0x96, 0xa0, 0x84, 0x96, 0x0c, 0xb6, 0x32, 0x95, 0xb5, 0xf6, 0xab, 0x39, 0x2c,
	0x7e, 0xd3, 0x98, 0xdf, 0xee, 0xe1, 0x01, 0x1f, 0x6f, 0x51, 0x77, 0x9a, 0xe7, 0xef, 0xd1, 0x68,
	0x8a, 0x65, 0x79, 0x65, 0x18, 0x8e, 0x6e, 0x2a, 0x32, 0xf6, 0xe8, 0xd9, 0xe7, 0x01, 0x94, 0x79,
	0x65, 0xca, 0x16, 0x94, 0x45, 0x56, 0x22, 0xf8, 0x8e, 0xc6, 0x98, 0x72, 0x9e, 0x57, 0x86, 0xf3,
	0x96, 0x70, 0x5e, 0x64, 0x9c, 0x23, 0x5a, 0xfe, 0xf7, 0x5f, 0x77, 0xfb, 0x13, 0x00, 0x00, 0xff,
	0xff, 0xce, 0x35, 0xdb, 0xd7, 0x24, 0x02, 0x00, 0x00,
}
