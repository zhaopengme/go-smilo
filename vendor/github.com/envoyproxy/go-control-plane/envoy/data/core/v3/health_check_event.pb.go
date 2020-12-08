// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/data/core/v3/health_check_event.proto

package envoy_data_core_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type HealthCheckFailureType int32

const (
	HealthCheckFailureType_ACTIVE  HealthCheckFailureType = 0
	HealthCheckFailureType_PASSIVE HealthCheckFailureType = 1
	HealthCheckFailureType_NETWORK HealthCheckFailureType = 2
)

var HealthCheckFailureType_name = map[int32]string{
	0: "ACTIVE",
	1: "PASSIVE",
	2: "NETWORK",
}

var HealthCheckFailureType_value = map[string]int32{
	"ACTIVE":  0,
	"PASSIVE": 1,
	"NETWORK": 2,
}

func (x HealthCheckFailureType) String() string {
	return proto.EnumName(HealthCheckFailureType_name, int32(x))
}

func (HealthCheckFailureType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b418bf0decf4b3ef, []int{0}
}

type HealthCheckerType int32

const (
	HealthCheckerType_HTTP  HealthCheckerType = 0
	HealthCheckerType_TCP   HealthCheckerType = 1
	HealthCheckerType_GRPC  HealthCheckerType = 2
	HealthCheckerType_REDIS HealthCheckerType = 3
)

var HealthCheckerType_name = map[int32]string{
	0: "HTTP",
	1: "TCP",
	2: "GRPC",
	3: "REDIS",
}

var HealthCheckerType_value = map[string]int32{
	"HTTP":  0,
	"TCP":   1,
	"GRPC":  2,
	"REDIS": 3,
}

func (x HealthCheckerType) String() string {
	return proto.EnumName(HealthCheckerType_name, int32(x))
}

func (HealthCheckerType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b418bf0decf4b3ef, []int{1}
}

type HealthCheckEvent struct {
	HealthCheckerType HealthCheckerType `protobuf:"varint,1,opt,name=health_checker_type,json=healthCheckerType,proto3,enum=envoy.data.core.v3.HealthCheckerType" json:"health_checker_type,omitempty"`
	Host              *v3.Address       `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	ClusterName       string            `protobuf:"bytes,3,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	// Types that are valid to be assigned to Event:
	//	*HealthCheckEvent_EjectUnhealthyEvent
	//	*HealthCheckEvent_AddHealthyEvent
	//	*HealthCheckEvent_HealthCheckFailureEvent
	//	*HealthCheckEvent_DegradedHealthyHost
	//	*HealthCheckEvent_NoLongerDegradedHost
	Event                isHealthCheckEvent_Event `protobuf_oneof:"event"`
	Timestamp            *timestamp.Timestamp     `protobuf:"bytes,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *HealthCheckEvent) Reset()         { *m = HealthCheckEvent{} }
func (m *HealthCheckEvent) String() string { return proto.CompactTextString(m) }
func (*HealthCheckEvent) ProtoMessage()    {}
func (*HealthCheckEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_b418bf0decf4b3ef, []int{0}
}

func (m *HealthCheckEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckEvent.Unmarshal(m, b)
}
func (m *HealthCheckEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckEvent.Marshal(b, m, deterministic)
}
func (m *HealthCheckEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckEvent.Merge(m, src)
}
func (m *HealthCheckEvent) XXX_Size() int {
	return xxx_messageInfo_HealthCheckEvent.Size(m)
}
func (m *HealthCheckEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckEvent.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckEvent proto.InternalMessageInfo

func (m *HealthCheckEvent) GetHealthCheckerType() HealthCheckerType {
	if m != nil {
		return m.HealthCheckerType
	}
	return HealthCheckerType_HTTP
}

func (m *HealthCheckEvent) GetHost() *v3.Address {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *HealthCheckEvent) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

type isHealthCheckEvent_Event interface {
	isHealthCheckEvent_Event()
}

type HealthCheckEvent_EjectUnhealthyEvent struct {
	EjectUnhealthyEvent *HealthCheckEjectUnhealthy `protobuf:"bytes,4,opt,name=eject_unhealthy_event,json=ejectUnhealthyEvent,proto3,oneof"`
}

type HealthCheckEvent_AddHealthyEvent struct {
	AddHealthyEvent *HealthCheckAddHealthy `protobuf:"bytes,5,opt,name=add_healthy_event,json=addHealthyEvent,proto3,oneof"`
}

type HealthCheckEvent_HealthCheckFailureEvent struct {
	HealthCheckFailureEvent *HealthCheckFailure `protobuf:"bytes,7,opt,name=health_check_failure_event,json=healthCheckFailureEvent,proto3,oneof"`
}

type HealthCheckEvent_DegradedHealthyHost struct {
	DegradedHealthyHost *DegradedHealthyHost `protobuf:"bytes,8,opt,name=degraded_healthy_host,json=degradedHealthyHost,proto3,oneof"`
}

type HealthCheckEvent_NoLongerDegradedHost struct {
	NoLongerDegradedHost *NoLongerDegradedHost `protobuf:"bytes,9,opt,name=no_longer_degraded_host,json=noLongerDegradedHost,proto3,oneof"`
}

func (*HealthCheckEvent_EjectUnhealthyEvent) isHealthCheckEvent_Event() {}

func (*HealthCheckEvent_AddHealthyEvent) isHealthCheckEvent_Event() {}

func (*HealthCheckEvent_HealthCheckFailureEvent) isHealthCheckEvent_Event() {}

func (*HealthCheckEvent_DegradedHealthyHost) isHealthCheckEvent_Event() {}

func (*HealthCheckEvent_NoLongerDegradedHost) isHealthCheckEvent_Event() {}

func (m *HealthCheckEvent) GetEvent() isHealthCheckEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *HealthCheckEvent) GetEjectUnhealthyEvent() *HealthCheckEjectUnhealthy {
	if x, ok := m.GetEvent().(*HealthCheckEvent_EjectUnhealthyEvent); ok {
		return x.EjectUnhealthyEvent
	}
	return nil
}

func (m *HealthCheckEvent) GetAddHealthyEvent() *HealthCheckAddHealthy {
	if x, ok := m.GetEvent().(*HealthCheckEvent_AddHealthyEvent); ok {
		return x.AddHealthyEvent
	}
	return nil
}

func (m *HealthCheckEvent) GetHealthCheckFailureEvent() *HealthCheckFailure {
	if x, ok := m.GetEvent().(*HealthCheckEvent_HealthCheckFailureEvent); ok {
		return x.HealthCheckFailureEvent
	}
	return nil
}

func (m *HealthCheckEvent) GetDegradedHealthyHost() *DegradedHealthyHost {
	if x, ok := m.GetEvent().(*HealthCheckEvent_DegradedHealthyHost); ok {
		return x.DegradedHealthyHost
	}
	return nil
}

func (m *HealthCheckEvent) GetNoLongerDegradedHost() *NoLongerDegradedHost {
	if x, ok := m.GetEvent().(*HealthCheckEvent_NoLongerDegradedHost); ok {
		return x.NoLongerDegradedHost
	}
	return nil
}

func (m *HealthCheckEvent) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HealthCheckEvent) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HealthCheckEvent_EjectUnhealthyEvent)(nil),
		(*HealthCheckEvent_AddHealthyEvent)(nil),
		(*HealthCheckEvent_HealthCheckFailureEvent)(nil),
		(*HealthCheckEvent_DegradedHealthyHost)(nil),
		(*HealthCheckEvent_NoLongerDegradedHost)(nil),
	}
}

type HealthCheckEjectUnhealthy struct {
	FailureType          HealthCheckFailureType `protobuf:"varint,1,opt,name=failure_type,json=failureType,proto3,enum=envoy.data.core.v3.HealthCheckFailureType" json:"failure_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *HealthCheckEjectUnhealthy) Reset()         { *m = HealthCheckEjectUnhealthy{} }
func (m *HealthCheckEjectUnhealthy) String() string { return proto.CompactTextString(m) }
func (*HealthCheckEjectUnhealthy) ProtoMessage()    {}
func (*HealthCheckEjectUnhealthy) Descriptor() ([]byte, []int) {
	return fileDescriptor_b418bf0decf4b3ef, []int{1}
}

func (m *HealthCheckEjectUnhealthy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckEjectUnhealthy.Unmarshal(m, b)
}
func (m *HealthCheckEjectUnhealthy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckEjectUnhealthy.Marshal(b, m, deterministic)
}
func (m *HealthCheckEjectUnhealthy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckEjectUnhealthy.Merge(m, src)
}
func (m *HealthCheckEjectUnhealthy) XXX_Size() int {
	return xxx_messageInfo_HealthCheckEjectUnhealthy.Size(m)
}
func (m *HealthCheckEjectUnhealthy) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckEjectUnhealthy.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckEjectUnhealthy proto.InternalMessageInfo

func (m *HealthCheckEjectUnhealthy) GetFailureType() HealthCheckFailureType {
	if m != nil {
		return m.FailureType
	}
	return HealthCheckFailureType_ACTIVE
}

type HealthCheckAddHealthy struct {
	FirstCheck           bool     `protobuf:"varint,1,opt,name=first_check,json=firstCheck,proto3" json:"first_check,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheckAddHealthy) Reset()         { *m = HealthCheckAddHealthy{} }
func (m *HealthCheckAddHealthy) String() string { return proto.CompactTextString(m) }
func (*HealthCheckAddHealthy) ProtoMessage()    {}
func (*HealthCheckAddHealthy) Descriptor() ([]byte, []int) {
	return fileDescriptor_b418bf0decf4b3ef, []int{2}
}

func (m *HealthCheckAddHealthy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckAddHealthy.Unmarshal(m, b)
}
func (m *HealthCheckAddHealthy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckAddHealthy.Marshal(b, m, deterministic)
}
func (m *HealthCheckAddHealthy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckAddHealthy.Merge(m, src)
}
func (m *HealthCheckAddHealthy) XXX_Size() int {
	return xxx_messageInfo_HealthCheckAddHealthy.Size(m)
}
func (m *HealthCheckAddHealthy) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckAddHealthy.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckAddHealthy proto.InternalMessageInfo

func (m *HealthCheckAddHealthy) GetFirstCheck() bool {
	if m != nil {
		return m.FirstCheck
	}
	return false
}

type HealthCheckFailure struct {
	FailureType          HealthCheckFailureType `protobuf:"varint,1,opt,name=failure_type,json=failureType,proto3,enum=envoy.data.core.v3.HealthCheckFailureType" json:"failure_type,omitempty"`
	FirstCheck           bool                   `protobuf:"varint,2,opt,name=first_check,json=firstCheck,proto3" json:"first_check,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *HealthCheckFailure) Reset()         { *m = HealthCheckFailure{} }
func (m *HealthCheckFailure) String() string { return proto.CompactTextString(m) }
func (*HealthCheckFailure) ProtoMessage()    {}
func (*HealthCheckFailure) Descriptor() ([]byte, []int) {
	return fileDescriptor_b418bf0decf4b3ef, []int{3}
}

func (m *HealthCheckFailure) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckFailure.Unmarshal(m, b)
}
func (m *HealthCheckFailure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckFailure.Marshal(b, m, deterministic)
}
func (m *HealthCheckFailure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckFailure.Merge(m, src)
}
func (m *HealthCheckFailure) XXX_Size() int {
	return xxx_messageInfo_HealthCheckFailure.Size(m)
}
func (m *HealthCheckFailure) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckFailure.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckFailure proto.InternalMessageInfo

func (m *HealthCheckFailure) GetFailureType() HealthCheckFailureType {
	if m != nil {
		return m.FailureType
	}
	return HealthCheckFailureType_ACTIVE
}

func (m *HealthCheckFailure) GetFirstCheck() bool {
	if m != nil {
		return m.FirstCheck
	}
	return false
}

type DegradedHealthyHost struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DegradedHealthyHost) Reset()         { *m = DegradedHealthyHost{} }
func (m *DegradedHealthyHost) String() string { return proto.CompactTextString(m) }
func (*DegradedHealthyHost) ProtoMessage()    {}
func (*DegradedHealthyHost) Descriptor() ([]byte, []int) {
	return fileDescriptor_b418bf0decf4b3ef, []int{4}
}

func (m *DegradedHealthyHost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DegradedHealthyHost.Unmarshal(m, b)
}
func (m *DegradedHealthyHost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DegradedHealthyHost.Marshal(b, m, deterministic)
}
func (m *DegradedHealthyHost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DegradedHealthyHost.Merge(m, src)
}
func (m *DegradedHealthyHost) XXX_Size() int {
	return xxx_messageInfo_DegradedHealthyHost.Size(m)
}
func (m *DegradedHealthyHost) XXX_DiscardUnknown() {
	xxx_messageInfo_DegradedHealthyHost.DiscardUnknown(m)
}

var xxx_messageInfo_DegradedHealthyHost proto.InternalMessageInfo

type NoLongerDegradedHost struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NoLongerDegradedHost) Reset()         { *m = NoLongerDegradedHost{} }
func (m *NoLongerDegradedHost) String() string { return proto.CompactTextString(m) }
func (*NoLongerDegradedHost) ProtoMessage()    {}
func (*NoLongerDegradedHost) Descriptor() ([]byte, []int) {
	return fileDescriptor_b418bf0decf4b3ef, []int{5}
}

func (m *NoLongerDegradedHost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NoLongerDegradedHost.Unmarshal(m, b)
}
func (m *NoLongerDegradedHost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NoLongerDegradedHost.Marshal(b, m, deterministic)
}
func (m *NoLongerDegradedHost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoLongerDegradedHost.Merge(m, src)
}
func (m *NoLongerDegradedHost) XXX_Size() int {
	return xxx_messageInfo_NoLongerDegradedHost.Size(m)
}
func (m *NoLongerDegradedHost) XXX_DiscardUnknown() {
	xxx_messageInfo_NoLongerDegradedHost.DiscardUnknown(m)
}

var xxx_messageInfo_NoLongerDegradedHost proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("envoy.data.core.v3.HealthCheckFailureType", HealthCheckFailureType_name, HealthCheckFailureType_value)
	proto.RegisterEnum("envoy.data.core.v3.HealthCheckerType", HealthCheckerType_name, HealthCheckerType_value)
	proto.RegisterType((*HealthCheckEvent)(nil), "envoy.data.core.v3.HealthCheckEvent")
	proto.RegisterType((*HealthCheckEjectUnhealthy)(nil), "envoy.data.core.v3.HealthCheckEjectUnhealthy")
	proto.RegisterType((*HealthCheckAddHealthy)(nil), "envoy.data.core.v3.HealthCheckAddHealthy")
	proto.RegisterType((*HealthCheckFailure)(nil), "envoy.data.core.v3.HealthCheckFailure")
	proto.RegisterType((*DegradedHealthyHost)(nil), "envoy.data.core.v3.DegradedHealthyHost")
	proto.RegisterType((*NoLongerDegradedHost)(nil), "envoy.data.core.v3.NoLongerDegradedHost")
}

func init() {
	proto.RegisterFile("envoy/data/core/v3/health_check_event.proto", fileDescriptor_b418bf0decf4b3ef)
}

var fileDescriptor_b418bf0decf4b3ef = []byte{
	// 743 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xcd, 0x6e, 0xda, 0x4a,
	0x14, 0xc7, 0x31, 0xdf, 0x1c, 0xa2, 0x7b, 0x9d, 0x49, 0xb8, 0xe1, 0x22, 0xb5, 0xa1, 0x48, 0x6d,
	0x29, 0x4d, 0x6c, 0x05, 0xba, 0x88, 0x88, 0x54, 0x09, 0x13, 0x5a, 0xa2, 0x54, 0x29, 0x72, 0x68,
	0xb3, 0xaa, 0xac, 0x09, 0x1e, 0xc0, 0x2d, 0x78, 0x90, 0x6d, 0x50, 0xd9, 0x76, 0xd5, 0x67, 0xe8,
	0x43, 0x74, 0xd1, 0x37, 0xe8, 0xa2, 0xef, 0x54, 0x65, 0x55, 0xcd, 0x8c, 0x81, 0x00, 0x8e, 0xc2,
	0xa6, 0x3b, 0xe6, 0xf8, 0x9c, 0xff, 0xef, 0x7c, 0x02, 0xcf, 0x89, 0x3d, 0xa1, 0x53, 0xd5, 0xc4,
	0x1e, 0x56, 0x3b, 0xd4, 0x21, 0xea, 0xa4, 0xa2, 0xf6, 0x09, 0x1e, 0x78, 0x7d, 0xa3, 0xd3, 0x27,
	0x9d, 0x4f, 0x06, 0x99, 0x10, 0xdb, 0x53, 0x46, 0x0e, 0xf5, 0x28, 0x42, 0xdc, 0x59, 0x61, 0xce,
	0x0a, 0x73, 0x56, 0x26, 0x95, 0x5c, 0x41, 0x08, 0x74, 0xa8, 0xdd, 0xb5, 0x7a, 0x73, 0x09, 0x6c,
	0x9a, 0x0e, 0x71, 0x5d, 0x11, 0x97, 0xdb, 0xef, 0x51, 0xda, 0x1b, 0x10, 0x95, 0xbf, 0xae, 0xc7,
	0x5d, 0xd5, 0xb3, 0x86, 0xc4, 0xf5, 0xf0, 0x70, 0xe4, 0x3b, 0x3c, 0x1a, 0x9b, 0x23, 0xac, 0x62,
	0xdb, 0xa6, 0x1e, 0xf6, 0x2c, 0x6a, 0xbb, 0xea, 0x84, 0x38, 0xae, 0x45, 0x6d, 0xcb, 0xee, 0xf9,
	0x2e, 0x7b, 0x13, 0x3c, 0xb0, 0x4c, 0xec, 0x11, 0x75, 0xf6, 0x43, 0x7c, 0x28, 0xfc, 0x88, 0x83,
	0xdc, 0xe4, 0x19, 0xd7, 0x59, 0xc2, 0x0d, 0x96, 0x2f, 0x32, 0x60, 0xe7, 0x76, 0x15, 0xc4, 0x31,
	0xbc, 0xe9, 0x88, 0x64, 0xa5, 0xbc, 0x54, 0xfc, 0xa7, 0xfc, 0x58, 0x59, 0xaf, 0x43, 0xb9, 0x25,
	0x41, 0x9c, 0xf6, 0x74, 0x44, 0xb4, 0xe4, 0x8d, 0x16, 0xfb, 0x22, 0x85, 0x65, 0x49, 0xdf, 0xee,
	0xaf, 0x7e, 0x44, 0x47, 0x10, 0xed, 0x53, 0xd7, 0xcb, 0x86, 0xf3, 0x52, 0x31, 0x5d, 0x7e, 0xe0,
	0x2b, 0x8a, 0x2e, 0xcc, 0x35, 0x6b, 0xa2, 0x0b, 0x3a, 0x77, 0x45, 0x25, 0xd8, 0xea, 0x0c, 0xc6,
	0xae, 0x47, 0x1c, 0xc3, 0xc6, 0x43, 0x92, 0x8d, 0xe4, 0xa5, 0x62, 0x4a, 0x4b, 0xdc, 0x68, 0x51,
	0x27, 0x9c, 0x97, 0xf4, 0xb4, 0xff, 0xf1, 0x02, 0x0f, 0x09, 0xea, 0x40, 0x86, 0x7c, 0x24, 0x1d,
	0xcf, 0x18, 0xdb, 0x82, 0x3d, 0x15, 0x83, 0xc8, 0x46, 0x39, 0xef, 0xf0, 0x9e, 0x0a, 0x1a, 0x2c,
	0xf6, 0xdd, 0x2c, 0xb4, 0x19, 0xd2, 0x77, 0xc8, 0x92, 0x45, 0x34, 0xe9, 0x0a, 0xb6, 0xb1, 0x69,
	0x1a, 0xcb, 0x80, 0x18, 0x07, 0x3c, 0xbb, 0x07, 0x50, 0x33, 0xcd, 0xe6, 0x5c, 0xfc, 0x5f, 0x3c,
	0x7f, 0x09, 0x61, 0x02, 0xb9, 0xa5, 0x1d, 0xea, 0x62, 0x6b, 0x30, 0x76, 0x88, 0x4f, 0x48, 0x70,
	0xc2, 0x93, 0x7b, 0x08, 0xaf, 0x44, 0x4c, 0x33, 0xa4, 0xef, 0xf5, 0xd7, 0xac, 0x02, 0xf3, 0x01,
	0x32, 0x26, 0xe9, 0x39, 0xd8, 0x24, 0x8b, 0x22, 0xf8, 0x50, 0x92, 0x9c, 0xf0, 0x34, 0x88, 0x70,
	0xea, 0x07, 0xcc, 0xb2, 0xa7, 0xae, 0xc7, 0xda, 0x63, 0xae, 0x9b, 0x11, 0x86, 0x3d, 0x9b, 0x1a,
	0x03, 0x6a, 0xf7, 0x88, 0x63, 0x2c, 0x40, 0x0c, 0x90, 0xe2, 0x80, 0x62, 0x10, 0xe0, 0x82, 0xbe,
	0xe1, 0x11, 0x73, 0x90, 0x20, 0xec, 0xda, 0x01, 0x76, 0x74, 0x0c, 0xa9, 0xf9, 0x29, 0x64, 0xe3,
	0x5c, 0x34, 0xa7, 0x88, 0x63, 0x51, 0x66, 0xc7, 0xa2, 0xb4, 0x67, 0x1e, 0xfa, 0xc2, 0xb9, 0xaa,
	0x7e, 0xfb, 0xf5, 0xf5, 0x61, 0x09, 0x8a, 0x6b, 0x19, 0x94, 0xf1, 0x60, 0xd4, 0xc7, 0xca, 0xea,
	0x45, 0x68, 0x5b, 0x10, 0xe3, 0xed, 0x47, 0x91, 0xdf, 0x9a, 0x54, 0xf8, 0x2e, 0xc1, 0xff, 0x77,
	0xee, 0x0b, 0xba, 0x82, 0xad, 0xd9, 0xc8, 0x6e, 0x9d, 0x4d, 0x69, 0xb3, 0x89, 0xad, 0xdc, 0x4e,
	0xba, 0xbb, 0x30, 0x57, 0x8f, 0x59, 0xd6, 0x15, 0x38, 0xda, 0x24, 0xeb, 0xa5, 0x94, 0x0a, 0x36,
	0x64, 0x02, 0xd7, 0x0f, 0xed, 0x43, 0xba, 0x6b, 0x39, 0xae, 0x27, 0x56, 0x8d, 0xa7, 0x9a, 0xd4,
	0x81, 0x9b, 0xb8, 0x6b, 0xf5, 0x05, 0x63, 0xaa, 0x70, 0xb8, 0x01, 0x73, 0x21, 0x5b, 0xf8, 0x29,
	0x01, 0x5a, 0xaf, 0xed, 0xaf, 0x75, 0x66, 0xb5, 0x8c, 0xf0, 0x5a, 0x19, 0x47, 0xac, 0x8c, 0x03,
	0x28, 0x6d, 0x50, 0x86, 0x8f, 0x2b, 0x9c, 0xc1, 0x4e, 0xc0, 0xba, 0x57, 0xcb, 0x4c, 0xe9, 0xd0,
	0xff, 0xe7, 0x0f, 0x50, 0x0a, 0x88, 0x29, 0x9c, 0xc3, 0x6e, 0xd0, 0x62, 0x57, 0x2b, 0x4c, 0x4b,
	0x81, 0x83, 0xbb, 0xb4, 0x82, 0x82, 0x4a, 0x2f, 0xe1, 0xbf, 0xe0, 0xe6, 0x20, 0x80, 0x78, 0xad,
	0xde, 0x3e, 0x7b, 0xdf, 0x90, 0x43, 0x28, 0x0d, 0x89, 0x56, 0xed, 0xf2, 0x92, 0x3d, 0x24, 0xf6,
	0xb8, 0x68, 0xb4, 0xaf, 0xde, 0xea, 0xe7, 0x72, 0xb8, 0x74, 0x02, 0xdb, 0x6b, 0xff, 0xd6, 0x28,
	0x09, 0xd1, 0x66, 0xbb, 0xdd, 0x92, 0x43, 0x28, 0x01, 0x91, 0x76, 0xbd, 0x25, 0x4b, 0xcc, 0xf4,
	0x5a, 0x6f, 0xd5, 0xe5, 0x30, 0x4a, 0x41, 0x4c, 0x6f, 0x9c, 0x9e, 0x5d, 0xca, 0x11, 0xed, 0x04,
	0xf2, 0x16, 0x15, 0xf3, 0x1a, 0x39, 0xf4, 0xf3, 0x34, 0x60, 0x74, 0x5a, 0x66, 0xf5, 0x7a, 0x5a,
	0xec, 0x16, 0x5b, 0xd2, 0x75, 0x9c, 0x1f, 0x65, 0xe5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x90,
	0xb3, 0x90, 0x03, 0x36, 0x07, 0x00, 0x00,
}
