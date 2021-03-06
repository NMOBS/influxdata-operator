// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/criterion_category_locale_availability_mode.proto

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

// Enum containing the possible CriterionCategoryLocaleAvailabilityMode.
type CriterionCategoryLocaleAvailabilityModeEnum_CriterionCategoryLocaleAvailabilityMode int32

const (
	// Not specified.
	CriterionCategoryLocaleAvailabilityModeEnum_UNSPECIFIED CriterionCategoryLocaleAvailabilityModeEnum_CriterionCategoryLocaleAvailabilityMode = 0
	// Used for return value only. Represents value unknown in this version.
	CriterionCategoryLocaleAvailabilityModeEnum_UNKNOWN CriterionCategoryLocaleAvailabilityModeEnum_CriterionCategoryLocaleAvailabilityMode = 1
	// The category is available to campaigns of all locales.
	CriterionCategoryLocaleAvailabilityModeEnum_ALL_LOCALES CriterionCategoryLocaleAvailabilityModeEnum_CriterionCategoryLocaleAvailabilityMode = 2
	// The category is available to campaigns within a list of countries,
	// regardless of language.
	CriterionCategoryLocaleAvailabilityModeEnum_COUNTRY_AND_ALL_LANGUAGES CriterionCategoryLocaleAvailabilityModeEnum_CriterionCategoryLocaleAvailabilityMode = 3
	// The category is available to campaigns within a list of languages,
	// regardless of country.
	CriterionCategoryLocaleAvailabilityModeEnum_LANGUAGE_AND_ALL_COUNTRIES CriterionCategoryLocaleAvailabilityModeEnum_CriterionCategoryLocaleAvailabilityMode = 4
	// The category is available to campaigns within a list of country, language
	// pairs.
	CriterionCategoryLocaleAvailabilityModeEnum_COUNTRY_AND_LANGUAGE CriterionCategoryLocaleAvailabilityModeEnum_CriterionCategoryLocaleAvailabilityMode = 5
)

var CriterionCategoryLocaleAvailabilityModeEnum_CriterionCategoryLocaleAvailabilityMode_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ALL_LOCALES",
	3: "COUNTRY_AND_ALL_LANGUAGES",
	4: "LANGUAGE_AND_ALL_COUNTRIES",
	5: "COUNTRY_AND_LANGUAGE",
}
var CriterionCategoryLocaleAvailabilityModeEnum_CriterionCategoryLocaleAvailabilityMode_value = map[string]int32{
	"UNSPECIFIED":                0,
	"UNKNOWN":                    1,
	"ALL_LOCALES":                2,
	"COUNTRY_AND_ALL_LANGUAGES":  3,
	"LANGUAGE_AND_ALL_COUNTRIES": 4,
	"COUNTRY_AND_LANGUAGE":       5,
}

func (x CriterionCategoryLocaleAvailabilityModeEnum_CriterionCategoryLocaleAvailabilityMode) String() string {
	return proto.EnumName(CriterionCategoryLocaleAvailabilityModeEnum_CriterionCategoryLocaleAvailabilityMode_name, int32(x))
}
func (CriterionCategoryLocaleAvailabilityModeEnum_CriterionCategoryLocaleAvailabilityMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_criterion_category_locale_availability_mode_3ad4b9809399a3b4, []int{0, 0}
}

// Describes locale availabilty mode for a criterion availability - whether
// it's available globally, or a particular country with all languages, or a
// particular language with all countries, or a country-language pair.
type CriterionCategoryLocaleAvailabilityModeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CriterionCategoryLocaleAvailabilityModeEnum) Reset() {
	*m = CriterionCategoryLocaleAvailabilityModeEnum{}
}
func (m *CriterionCategoryLocaleAvailabilityModeEnum) String() string {
	return proto.CompactTextString(m)
}
func (*CriterionCategoryLocaleAvailabilityModeEnum) ProtoMessage() {}
func (*CriterionCategoryLocaleAvailabilityModeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_criterion_category_locale_availability_mode_3ad4b9809399a3b4, []int{0}
}
func (m *CriterionCategoryLocaleAvailabilityModeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CriterionCategoryLocaleAvailabilityModeEnum.Unmarshal(m, b)
}
func (m *CriterionCategoryLocaleAvailabilityModeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CriterionCategoryLocaleAvailabilityModeEnum.Marshal(b, m, deterministic)
}
func (dst *CriterionCategoryLocaleAvailabilityModeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CriterionCategoryLocaleAvailabilityModeEnum.Merge(dst, src)
}
func (m *CriterionCategoryLocaleAvailabilityModeEnum) XXX_Size() int {
	return xxx_messageInfo_CriterionCategoryLocaleAvailabilityModeEnum.Size(m)
}
func (m *CriterionCategoryLocaleAvailabilityModeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CriterionCategoryLocaleAvailabilityModeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CriterionCategoryLocaleAvailabilityModeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CriterionCategoryLocaleAvailabilityModeEnum)(nil), "google.ads.googleads.v0.enums.CriterionCategoryLocaleAvailabilityModeEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.CriterionCategoryLocaleAvailabilityModeEnum_CriterionCategoryLocaleAvailabilityMode", CriterionCategoryLocaleAvailabilityModeEnum_CriterionCategoryLocaleAvailabilityMode_name, CriterionCategoryLocaleAvailabilityModeEnum_CriterionCategoryLocaleAvailabilityMode_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/criterion_category_locale_availability_mode.proto", fileDescriptor_criterion_category_locale_availability_mode_3ad4b9809399a3b4)
}

var fileDescriptor_criterion_category_locale_availability_mode_3ad4b9809399a3b4 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x4a, 0xc3, 0x30,
	0x1c, 0xc7, 0xed, 0xe6, 0x1f, 0xc8, 0x0e, 0x96, 0xe2, 0x41, 0x85, 0x09, 0xee, 0xe2, 0x41, 0x49,
	0x0b, 0x1e, 0x3d, 0x65, 0x5d, 0x2d, 0xc3, 0x9a, 0x0e, 0x6b, 0x27, 0x4a, 0x21, 0x64, 0x6d, 0x08,
	0x85, 0xb6, 0x19, 0x4d, 0x37, 0xd8, 0xeb, 0x78, 0xf4, 0xe6, 0x6b, 0xf8, 0x0e, 0x1e, 0x7c, 0x13,
	0x69, 0xba, 0x8e, 0x5d, 0x94, 0x5d, 0xca, 0xb7, 0x7c, 0x7e, 0xf9, 0x84, 0x7c, 0x7f, 0xc0, 0xe7,
	0x42, 0xf0, 0x8c, 0x99, 0x34, 0x91, 0x66, 0x13, 0xeb, 0xb4, 0xb4, 0x4c, 0x56, 0x2c, 0x72, 0x69,
	0xc6, 0x65, 0x5a, 0xb1, 0x32, 0x15, 0x05, 0x89, 0x69, 0xc5, 0xb8, 0x28, 0x57, 0x24, 0x13, 0x31,
	0xcd, 0x18, 0xa1, 0x4b, 0x9a, 0x66, 0x74, 0x96, 0x66, 0x69, 0xb5, 0x22, 0xb9, 0x48, 0x18, 0x9c,
	0x97, 0xa2, 0x12, 0x46, 0xbf, 0xb1, 0x40, 0x9a, 0x48, 0xb8, 0x11, 0xc2, 0xa5, 0x05, 0x95, 0x70,
	0xf0, 0xa3, 0x81, 0x6b, 0xbb, 0x95, 0xda, 0x6b, 0xa7, 0xa7, 0x94, 0x68, 0xcb, 0xf8, 0x28, 0x12,
	0xe6, 0x14, 0x8b, 0x7c, 0xf0, 0xa9, 0x81, 0xab, 0x1d, 0xe7, 0x8d, 0x63, 0xd0, 0x0b, 0x71, 0x30,
	0x71, 0xec, 0xf1, 0xfd, 0xd8, 0x19, 0xe9, 0x7b, 0x46, 0x0f, 0x1c, 0x85, 0xf8, 0x01, 0xfb, 0x2f,
	0x58, 0xd7, 0x6a, 0x8a, 0x3c, 0x8f, 0x78, 0xbe, 0x8d, 0x3c, 0x27, 0xd0, 0x3b, 0x46, 0x1f, 0x9c,
	0xd9, 0x7e, 0x88, 0x9f, 0x9f, 0x5e, 0x09, 0xc2, 0x23, 0xa2, 0x20, 0xc2, 0x6e, 0x88, 0x5c, 0x27,
	0xd0, 0xbb, 0xc6, 0x05, 0x38, 0x6f, 0x7f, 0x37, 0xbc, 0x99, 0x1f, 0x3b, 0x81, 0xbe, 0x6f, 0x9c,
	0x82, 0x93, 0xed, 0xe3, 0xed, 0xac, 0x7e, 0x30, 0xfc, 0xd6, 0xc0, 0x65, 0x2c, 0x72, 0xf8, 0x6f,
	0x13, 0xc3, 0x9b, 0x1d, 0x9f, 0x35, 0xa9, 0x6b, 0x9d, 0x68, 0x6f, 0xc3, 0xb5, 0x8e, 0x8b, 0x8c,
	0x16, 0x1c, 0x8a, 0x92, 0x9b, 0x9c, 0x15, 0xaa, 0xf4, 0x76, 0x73, 0xf3, 0x54, 0xfe, 0xb1, 0xc8,
	0x3b, 0xf5, 0x7d, 0xef, 0x74, 0x5d, 0x84, 0x3e, 0x3a, 0x7d, 0xb7, 0x51, 0xa1, 0x44, 0xc2, 0x26,
	0xd6, 0x69, 0x6a, 0xc1, 0xba, 0x72, 0xf9, 0xd5, 0xf2, 0x08, 0x25, 0x32, 0xda, 0xf0, 0x68, 0x6a,
	0x45, 0x8a, 0xcf, 0x0e, 0xd5, 0xa5, 0xb7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc4, 0xfc, 0x96,
	0x5b, 0x3c, 0x02, 0x00, 0x00,
}
