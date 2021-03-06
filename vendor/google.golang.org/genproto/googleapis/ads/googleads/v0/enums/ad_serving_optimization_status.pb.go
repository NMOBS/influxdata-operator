// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/ad_serving_optimization_status.proto

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

// Enum describing possible serving statuses.
type AdServingOptimizationStatusEnum_AdServingOptimizationStatus int32

const (
	// No value has been specified.
	AdServingOptimizationStatusEnum_UNSPECIFIED AdServingOptimizationStatusEnum_AdServingOptimizationStatus = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	AdServingOptimizationStatusEnum_UNKNOWN AdServingOptimizationStatusEnum_AdServingOptimizationStatus = 1
	// Ad serving is optimized based on CTR for the campaign.
	AdServingOptimizationStatusEnum_OPTIMIZE AdServingOptimizationStatusEnum_AdServingOptimizationStatus = 2
	// Ad serving is optimized based on CTR * Conversion for the campaign. If
	// the campaign is not in the conversion optimizer bidding strategy, it will
	// default to OPTIMIZED.
	AdServingOptimizationStatusEnum_CONVERSION_OPTIMIZE AdServingOptimizationStatusEnum_AdServingOptimizationStatus = 3
	// Ads are rotated evenly for 90 days, then optimized for clicks.
	AdServingOptimizationStatusEnum_ROTATE AdServingOptimizationStatusEnum_AdServingOptimizationStatus = 4
	// Show lower performing ads more evenly with higher performing ads, and do
	// not optimize.
	AdServingOptimizationStatusEnum_ROTATE_INDEFINITELY AdServingOptimizationStatusEnum_AdServingOptimizationStatus = 5
	// Ad serving optimization status is not available.
	AdServingOptimizationStatusEnum_UNAVAILABLE AdServingOptimizationStatusEnum_AdServingOptimizationStatus = 6
)

var AdServingOptimizationStatusEnum_AdServingOptimizationStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "OPTIMIZE",
	3: "CONVERSION_OPTIMIZE",
	4: "ROTATE",
	5: "ROTATE_INDEFINITELY",
	6: "UNAVAILABLE",
}
var AdServingOptimizationStatusEnum_AdServingOptimizationStatus_value = map[string]int32{
	"UNSPECIFIED":         0,
	"UNKNOWN":             1,
	"OPTIMIZE":            2,
	"CONVERSION_OPTIMIZE": 3,
	"ROTATE":              4,
	"ROTATE_INDEFINITELY": 5,
	"UNAVAILABLE":         6,
}

func (x AdServingOptimizationStatusEnum_AdServingOptimizationStatus) String() string {
	return proto.EnumName(AdServingOptimizationStatusEnum_AdServingOptimizationStatus_name, int32(x))
}
func (AdServingOptimizationStatusEnum_AdServingOptimizationStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ad_serving_optimization_status_897e6671d350f9ab, []int{0, 0}
}

// Possible ad serving statuses of a campaign.
type AdServingOptimizationStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdServingOptimizationStatusEnum) Reset()         { *m = AdServingOptimizationStatusEnum{} }
func (m *AdServingOptimizationStatusEnum) String() string { return proto.CompactTextString(m) }
func (*AdServingOptimizationStatusEnum) ProtoMessage()    {}
func (*AdServingOptimizationStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad_serving_optimization_status_897e6671d350f9ab, []int{0}
}
func (m *AdServingOptimizationStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdServingOptimizationStatusEnum.Unmarshal(m, b)
}
func (m *AdServingOptimizationStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdServingOptimizationStatusEnum.Marshal(b, m, deterministic)
}
func (dst *AdServingOptimizationStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdServingOptimizationStatusEnum.Merge(dst, src)
}
func (m *AdServingOptimizationStatusEnum) XXX_Size() int {
	return xxx_messageInfo_AdServingOptimizationStatusEnum.Size(m)
}
func (m *AdServingOptimizationStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdServingOptimizationStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdServingOptimizationStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AdServingOptimizationStatusEnum)(nil), "google.ads.googleads.v0.enums.AdServingOptimizationStatusEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.AdServingOptimizationStatusEnum_AdServingOptimizationStatus", AdServingOptimizationStatusEnum_AdServingOptimizationStatus_name, AdServingOptimizationStatusEnum_AdServingOptimizationStatus_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/ad_serving_optimization_status.proto", fileDescriptor_ad_serving_optimization_status_897e6671d350f9ab)
}

var fileDescriptor_ad_serving_optimization_status_897e6671d350f9ab = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4a, 0xc3, 0x30,
	0x18, 0xc7, 0xed, 0xa6, 0x53, 0x32, 0xc1, 0x52, 0x0f, 0x1e, 0x64, 0xa8, 0x7b, 0x80, 0xb4, 0xe0,
	0xd1, 0x53, 0xba, 0x65, 0x23, 0x38, 0xd3, 0xb2, 0x76, 0x15, 0x47, 0xa1, 0x54, 0x53, 0x42, 0x61,
	0x4d, 0xc6, 0xd2, 0xed, 0xe0, 0x93, 0x78, 0xf6, 0xe8, 0x33, 0xf8, 0x04, 0xde, 0x7c, 0x23, 0x69,
	0xe2, 0x86, 0x17, 0x77, 0x29, 0x7f, 0xf8, 0x7f, 0xfd, 0x7d, 0xc9, 0x2f, 0xc0, 0xe7, 0x52, 0xf2,
	0x45, 0xe1, 0xe6, 0x4c, 0xb9, 0x26, 0x36, 0x69, 0xe3, 0xb9, 0x85, 0x58, 0x57, 0xca, 0xcd, 0x59,
	0xa6, 0x8a, 0xd5, 0xa6, 0x14, 0x3c, 0x93, 0xcb, 0xba, 0xac, 0xca, 0xd7, 0xbc, 0x2e, 0xa5, 0xc8,
	0x54, 0x9d, 0xd7, 0x6b, 0x05, 0x97, 0x2b, 0x59, 0x4b, 0xa7, 0x67, 0x7e, 0x84, 0x39, 0x53, 0x70,
	0xc7, 0x80, 0x1b, 0x0f, 0x6a, 0x46, 0xff, 0xd3, 0x02, 0x57, 0x88, 0x45, 0x06, 0x13, 0xfc, 0xa1,
	0x44, 0x1a, 0x82, 0xc5, 0xba, 0xea, 0xbf, 0x59, 0xe0, 0x72, 0xcf, 0x8c, 0x73, 0x06, 0xba, 0x33,
	0x1a, 0x85, 0x78, 0x40, 0x46, 0x04, 0x0f, 0xed, 0x03, 0xa7, 0x0b, 0x8e, 0x67, 0xf4, 0x9e, 0x06,
	0x8f, 0xd4, 0xb6, 0x9c, 0x53, 0x70, 0x12, 0x84, 0x31, 0x79, 0x20, 0x73, 0x6c, 0xb7, 0x9c, 0x0b,
	0x70, 0x3e, 0x08, 0x68, 0x82, 0xa7, 0x11, 0x09, 0x68, 0xb6, 0x2b, 0xda, 0x0e, 0x00, 0x9d, 0x69,
	0x10, 0xa3, 0x18, 0xdb, 0x87, 0xcd, 0x90, 0xc9, 0x19, 0xa1, 0x43, 0x3c, 0x22, 0x94, 0xc4, 0x78,
	0xf2, 0x64, 0x1f, 0x99, 0x4d, 0x28, 0x41, 0x64, 0x82, 0xfc, 0x09, 0xb6, 0x3b, 0xfe, 0xb7, 0x05,
	0x6e, 0x5e, 0x64, 0x05, 0xf7, 0x5e, 0xd2, 0xbf, 0xde, 0x73, 0xfa, 0xb0, 0xb1, 0x14, 0x5a, 0xf3,
	0x5f, 0xd7, 0x90, 0xcb, 0x45, 0x2e, 0x38, 0x94, 0x2b, 0xee, 0xf2, 0x42, 0x68, 0x87, 0x5b, 0xf7,
	0xcb, 0x52, 0xfd, 0xf3, 0x14, 0x77, 0xfa, 0xfb, 0xde, 0x6a, 0x8f, 0x11, 0xfa, 0x68, 0xf5, 0xc6,
	0x06, 0x85, 0x98, 0x82, 0x26, 0x36, 0x29, 0xf1, 0x60, 0x63, 0x53, 0x7d, 0x6d, 0xfb, 0x14, 0x31,
	0x95, 0xee, 0xfa, 0x34, 0xf1, 0x52, 0xdd, 0x3f, 0x77, 0xf4, 0xd2, 0xdb, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x6b, 0x1e, 0x9a, 0x55, 0xfe, 0x01, 0x00, 0x00,
}
