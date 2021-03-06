// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/rate_limit/v2/rate_limit.proto

package envoy_config_filter_http_rate_limit_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v2 "github.com/envoyproxy/go-control-plane/envoy/config/ratelimit/v2"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RateLimit struct {
	Domain                         string                     `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Stage                          uint32                     `protobuf:"varint,2,opt,name=stage,proto3" json:"stage,omitempty"`
	RequestType                    string                     `protobuf:"bytes,3,opt,name=request_type,json=requestType,proto3" json:"request_type,omitempty"`
	Timeout                        *duration.Duration         `protobuf:"bytes,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	FailureModeDeny                bool                       `protobuf:"varint,5,opt,name=failure_mode_deny,json=failureModeDeny,proto3" json:"failure_mode_deny,omitempty"`
	RateLimitedAsResourceExhausted bool                       `protobuf:"varint,6,opt,name=rate_limited_as_resource_exhausted,json=rateLimitedAsResourceExhausted,proto3" json:"rate_limited_as_resource_exhausted,omitempty"`
	RateLimitService               *v2.RateLimitServiceConfig `protobuf:"bytes,7,opt,name=rate_limit_service,json=rateLimitService,proto3" json:"rate_limit_service,omitempty"`
	XXX_NoUnkeyedLiteral           struct{}                   `json:"-"`
	XXX_unrecognized               []byte                     `json:"-"`
	XXX_sizecache                  int32                      `json:"-"`
}

func (m *RateLimit) Reset()         { *m = RateLimit{} }
func (m *RateLimit) String() string { return proto.CompactTextString(m) }
func (*RateLimit) ProtoMessage()    {}
func (*RateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_af348a51c982d3a6, []int{0}
}

func (m *RateLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimit.Unmarshal(m, b)
}
func (m *RateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimit.Marshal(b, m, deterministic)
}
func (m *RateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimit.Merge(m, src)
}
func (m *RateLimit) XXX_Size() int {
	return xxx_messageInfo_RateLimit.Size(m)
}
func (m *RateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimit proto.InternalMessageInfo

func (m *RateLimit) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimit) GetStage() uint32 {
	if m != nil {
		return m.Stage
	}
	return 0
}

func (m *RateLimit) GetRequestType() string {
	if m != nil {
		return m.RequestType
	}
	return ""
}

func (m *RateLimit) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *RateLimit) GetFailureModeDeny() bool {
	if m != nil {
		return m.FailureModeDeny
	}
	return false
}

func (m *RateLimit) GetRateLimitedAsResourceExhausted() bool {
	if m != nil {
		return m.RateLimitedAsResourceExhausted
	}
	return false
}

func (m *RateLimit) GetRateLimitService() *v2.RateLimitServiceConfig {
	if m != nil {
		return m.RateLimitService
	}
	return nil
}

func init() {
	proto.RegisterType((*RateLimit)(nil), "envoy.config.filter.http.rate_limit.v2.RateLimit")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/rate_limit/v2/rate_limit.proto", fileDescriptor_af348a51c982d3a6)
}

var fileDescriptor_af348a51c982d3a6 = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x6f, 0xd3, 0x30,
	0x14, 0x57, 0xba, 0xb5, 0xdd, 0x3c, 0x3e, 0x86, 0x2f, 0x84, 0x49, 0x8c, 0x30, 0x24, 0x54, 0x55,
	0xc8, 0x16, 0x2d, 0x12, 0x67, 0xc2, 0xb8, 0x20, 0x90, 0xa6, 0xc0, 0x3d, 0xf2, 0xea, 0xd7, 0xcc,
	0x52, 0x62, 0x07, 0xfb, 0x25, 0x6a, 0xf8, 0x03, 0x38, 0x70, 0xe1, 0xca, 0xdf, 0xc9, 0x71, 0x07,
	0x84, 0x12, 0x27, 0xed, 0xc6, 0x69, 0x37, 0xfb, 0xfd, 0x3e, 0xde, 0x27, 0x79, 0x0b, 0xba, 0x36,
	0x0d, 0x5f, 0x19, 0xbd, 0x56, 0x19, 0x5f, 0xab, 0x1c, 0xc1, 0xf2, 0x2b, 0xc4, 0x92, 0x5b, 0x81,
	0x90, 0xe6, 0xaa, 0x50, 0xc8, 0xeb, 0xc5, 0x8d, 0x1f, 0x2b, 0xad, 0x41, 0x43, 0x5f, 0x76, 0x42,
	0xe6, 0x85, 0xcc, 0x0b, 0x59, 0x2b, 0x64, 0x37, 0xa8, 0xf5, 0xe2, 0xe4, 0xc5, 0xad, 0x04, 0x2d,
	0xb6, 0xf3, 0xcc, 0x9d, 0x37, 0x3b, 0x39, 0xcd, 0x8c, 0xc9, 0x72, 0xe0, 0xdd, 0xef, 0xb2, 0x5a,
	0x73, 0x59, 0x59, 0x81, 0xca, 0xe8, 0x01, 0xaf, 0x64, 0x29, 0xb8, 0xd0, 0xda, 0x60, 0x17, 0x76,
	0xbc, 0x50, 0x59, 0xeb, 0xd5, 0xe3, 0x8f, 0x6b, 0x91, 0x2b, 0x29, 0x10, 0xf8, 0xf0, 0xf0, 0xc0,
	0xd9, 0x8f, 0x3d, 0x72, 0x98, 0x08, 0x84, 0x4f, 0x6d, 0x4e, 0xfa, 0x8c, 0x4c, 0xa4, 0x29, 0x84,
	0xd2, 0x61, 0x10, 0x05, 0xb3, 0xc3, 0x78, 0x7a, 0x1d, 0xef, 0xdb, 0x51, 0x14, 0x24, 0x7d, 0x98,
	0x3e, 0x25, 0x63, 0x87, 0x22, 0x83, 0x70, 0x14, 0x05, 0xb3, 0xfb, 0x1d, 0x3e, 0x1f, 0x85, 0x24,
	0xf1, 0x51, 0xfa, 0x9c, 0xdc, 0xb3, 0xf0, 0xad, 0x02, 0x87, 0x29, 0x36, 0x25, 0x84, 0x7b, 0xad,
	0x4b, 0x72, 0xd4, 0xc7, 0xbe, 0x36, 0x25, 0xd0, 0x25, 0x99, 0xa2, 0x2a, 0xc0, 0x54, 0x18, 0xee,
	0x47, 0xc1, 0xec, 0x68, 0xf1, 0x84, 0xf9, 0xde, 0xd8, 0xd0, 0x1b, 0x3b, 0xef, 0x7b, 0x4b, 0x06,
	0x26, 0x9d, 0x93, 0x47, 0x6b, 0xa1, 0xf2, 0xca, 0x42, 0x5a, 0x18, 0x09, 0xa9, 0x04, 0xdd, 0x84,
	0xe3, 0x28, 0x98, 0x1d, 0x24, 0x0f, 0x7b, 0xe0, 0xb3, 0x91, 0x70, 0x0e, 0xba, 0xa1, 0x1f, 0xc9,
	0xd9, 0x6e, 0xc0, 0x20, 0x53, 0xe1, 0x52, 0x0b, 0xce, 0x54, 0x76, 0x05, 0x29, 0x6c, 0xae, 0x44,
	0xe5, 0x10, 0x64, 0x38, 0xe9, 0xc4, 0xa7, 0x76, 0x68, 0x1d, 0xe4, 0x3b, 0x97, 0xf4, 0xb4, 0x0f,
	0x03, 0x8b, 0x2a, 0x42, 0x77, 0x5e, 0xa9, 0x03, 0x5b, 0xab, 0x15, 0x84, 0xd3, 0xae, 0xee, 0xd7,
	0xec, 0xd6, 0x82, 0xb7, 0x8b, 0x63, 0xf5, 0x82, 0x6d, 0x27, 0xfa, 0xc5, 0x4b, 0xde, 0x77, 0x9c,
	0xf8, 0xe0, 0x3a, 0x1e, 0xff, 0x0c, 0x46, 0xc7, 0x41, 0x72, 0x6c, 0xff, 0x63, 0xc4, 0xdf, 0xff,
	0xfc, 0xfe, 0xfb, 0x6b, 0xfc, 0x8a, 0xce, 0xbd, 0x2b, 0x6c, 0x10, 0xb4, 0x6b, 0x37, 0xd9, 0x9f,
	0x8e, 0xdb, 0xdd, 0x4e, 0x9f, 0x66, 0x49, 0xde, 0x28, 0xe3, 0x8b, 0x28, 0xad, 0xd9, 0x34, 0xec,
	0x6e, 0x07, 0x17, 0x3f, 0xd8, 0x56, 0x77, 0xd1, 0x4e, 0xfc, 0x22, 0xb8, 0x9c, 0x74, 0xa3, 0x5f,
	0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xbb, 0xa6, 0x24, 0xa9, 0xec, 0x02, 0x00, 0x00,
}
