// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/feed_mapping_criterion_type.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

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

// Possible placeholder types for a feed mapping.
type FeedMappingCriterionTypeEnum_FeedMappingCriterionType int32

const (
	// Not specified.
	FeedMappingCriterionTypeEnum_UNSPECIFIED FeedMappingCriterionTypeEnum_FeedMappingCriterionType = 0
	// Used for return value only. Represents value unknown in this version.
	FeedMappingCriterionTypeEnum_UNKNOWN FeedMappingCriterionTypeEnum_FeedMappingCriterionType = 1
	// Allows campaign targeting at locations within a location feed.
	FeedMappingCriterionTypeEnum_CAMPAIGN_LOCATION_TARGETS FeedMappingCriterionTypeEnum_FeedMappingCriterionType = 2
	// Allows url targeting for your dynamic search ads within a page feed.
	FeedMappingCriterionTypeEnum_DSA_PAGE_FEED FeedMappingCriterionTypeEnum_FeedMappingCriterionType = 3
)

var FeedMappingCriterionTypeEnum_FeedMappingCriterionType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "CAMPAIGN_LOCATION_TARGETS",
	3: "DSA_PAGE_FEED",
}
var FeedMappingCriterionTypeEnum_FeedMappingCriterionType_value = map[string]int32{
	"UNSPECIFIED":               0,
	"UNKNOWN":                   1,
	"CAMPAIGN_LOCATION_TARGETS": 2,
	"DSA_PAGE_FEED":             3,
}

func (x FeedMappingCriterionTypeEnum_FeedMappingCriterionType) String() string {
	return proto.EnumName(FeedMappingCriterionTypeEnum_FeedMappingCriterionType_name, int32(x))
}
func (FeedMappingCriterionTypeEnum_FeedMappingCriterionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_feed_mapping_criterion_type_d98ffdc43b560cb8, []int{0, 0}
}

// Container for enum describing possible criterion types for a feed mapping.
type FeedMappingCriterionTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedMappingCriterionTypeEnum) Reset()         { *m = FeedMappingCriterionTypeEnum{} }
func (m *FeedMappingCriterionTypeEnum) String() string { return proto.CompactTextString(m) }
func (*FeedMappingCriterionTypeEnum) ProtoMessage()    {}
func (*FeedMappingCriterionTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_feed_mapping_criterion_type_d98ffdc43b560cb8, []int{0}
}
func (m *FeedMappingCriterionTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedMappingCriterionTypeEnum.Unmarshal(m, b)
}
func (m *FeedMappingCriterionTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedMappingCriterionTypeEnum.Marshal(b, m, deterministic)
}
func (dst *FeedMappingCriterionTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedMappingCriterionTypeEnum.Merge(dst, src)
}
func (m *FeedMappingCriterionTypeEnum) XXX_Size() int {
	return xxx_messageInfo_FeedMappingCriterionTypeEnum.Size(m)
}
func (m *FeedMappingCriterionTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedMappingCriterionTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FeedMappingCriterionTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*FeedMappingCriterionTypeEnum)(nil), "google.ads.googleads.v0.enums.FeedMappingCriterionTypeEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.FeedMappingCriterionTypeEnum_FeedMappingCriterionType", FeedMappingCriterionTypeEnum_FeedMappingCriterionType_name, FeedMappingCriterionTypeEnum_FeedMappingCriterionType_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/feed_mapping_criterion_type.proto", fileDescriptor_feed_mapping_criterion_type_d98ffdc43b560cb8)
}

var fileDescriptor_feed_mapping_criterion_type_d98ffdc43b560cb8 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xd1, 0x4a, 0xfb, 0x30,
	0x18, 0xc5, 0xff, 0xed, 0xe0, 0x2f, 0x64, 0x88, 0xb5, 0x57, 0x0a, 0xf6, 0x62, 0x7b, 0x80, 0xb4,
	0xe0, 0x5d, 0xbc, 0x90, 0xac, 0xcb, 0x4a, 0xd1, 0x75, 0xc5, 0x75, 0x13, 0xa4, 0x10, 0xea, 0x12,
	0x43, 0x65, 0x6d, 0x4a, 0xb3, 0x0d, 0xf6, 0x0a, 0x3e, 0x86, 0x97, 0x3e, 0x8a, 0x8f, 0xe2, 0xa5,
	0x4f, 0x20, 0x6d, 0xdc, 0xee, 0xea, 0x4d, 0x38, 0xe4, 0x7c, 0xdf, 0x8f, 0xf3, 0x1d, 0x70, 0x2b,
	0xa4, 0x14, 0x6b, 0xee, 0x66, 0x4c, 0xb9, 0x5a, 0x36, 0x6a, 0xe7, 0xb9, 0xbc, 0xdc, 0x16, 0xca,
	0x7d, 0xe1, 0x9c, 0xd1, 0x22, 0xab, 0xaa, 0xbc, 0x14, 0x74, 0x55, 0xe7, 0x1b, 0x5e, 0xe7, 0xb2,
	0xa4, 0x9b, 0x7d, 0xc5, 0x61, 0x55, 0xcb, 0x8d, 0xb4, 0x1d, 0xbd, 0x05, 0x33, 0xa6, 0xe0, 0x11,
	0x00, 0x77, 0x1e, 0x6c, 0x01, 0xc3, 0x37, 0x03, 0x5c, 0x4d, 0x38, 0x67, 0x53, 0xcd, 0xf0, 0x0f,
	0x88, 0x64, 0x5f, 0x71, 0x52, 0x6e, 0x8b, 0xe1, 0x2b, 0xb8, 0xe8, 0xf2, 0xed, 0x33, 0xd0, 0x5f,
	0x44, 0xf3, 0x98, 0xf8, 0xe1, 0x24, 0x24, 0x63, 0xeb, 0x9f, 0xdd, 0x07, 0x27, 0x8b, 0xe8, 0x2e,
	0x9a, 0x3d, 0x46, 0x96, 0x61, 0x3b, 0xe0, 0xd2, 0xc7, 0xd3, 0x18, 0x87, 0x41, 0x44, 0xef, 0x67,
	0x3e, 0x4e, 0xc2, 0x59, 0x44, 0x13, 0xfc, 0x10, 0x90, 0x64, 0x6e, 0x99, 0xf6, 0x39, 0x38, 0x1d,
	0xcf, 0x31, 0x8d, 0x71, 0x40, 0xe8, 0x84, 0x90, 0xb1, 0xd5, 0x1b, 0x7d, 0x1b, 0x60, 0xb0, 0x92,
	0x05, 0xfc, 0x33, 0xf2, 0xc8, 0xe9, 0xca, 0x13, 0x37, 0x07, 0xc7, 0xc6, 0xd3, 0xe8, 0x77, 0x5f,
	0xc8, 0x75, 0x56, 0x0a, 0x28, 0x6b, 0xe1, 0x0a, 0x5e, 0xb6, 0x75, 0x1c, 0x3a, 0xac, 0x72, 0xd5,
	0x51, 0xe9, 0x4d, 0xfb, 0xbe, 0x9b, 0xbd, 0x00, 0xe3, 0x0f, 0xd3, 0x09, 0x34, 0x0a, 0x33, 0x05,
	0xb5, 0x6c, 0xd4, 0xd2, 0x83, 0x4d, 0x37, 0xea, 0xf3, 0xe0, 0xa7, 0x98, 0xa9, 0xf4, 0xe8, 0xa7,
	0x4b, 0x2f, 0x6d, 0xfd, 0x2f, 0x73, 0xa0, 0x3f, 0x11, 0xc2, 0x4c, 0x21, 0x74, 0x9c, 0x40, 0x68,
	0xe9, 0x21, 0xd4, 0xce, 0x3c, 0xff, 0x6f, 0x83, 0x5d, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x53,
	0x05, 0xac, 0xf5, 0xea, 0x01, 0x00, 0x00,
}
