// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: servicefunctionchain.proto

package model

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ServiceFunctionChain_Type int32

const (
	ServiceFunctionChain_ServiceFunctionPod ServiceFunctionChain_Type = 0
	ServiceFunctionChain_InterfaceInput     ServiceFunctionChain_Type = 1
	ServiceFunctionChain_InterfaceOutput    ServiceFunctionChain_Type = 2
	ServiceFunctionChain_PodInput           ServiceFunctionChain_Type = 3
	ServiceFunctionChain_PodOutput          ServiceFunctionChain_Type = 4
)

var ServiceFunctionChain_Type_name = map[int32]string{
	0: "ServiceFunctionPod",
	1: "InterfaceInput",
	2: "InterfaceOutput",
	3: "PodInput",
	4: "PodOutput",
}
var ServiceFunctionChain_Type_value = map[string]int32{
	"ServiceFunctionPod": 0,
	"InterfaceInput":     1,
	"InterfaceOutput":    2,
	"PodInput":           3,
	"PodOutput":          4,
}

func (x ServiceFunctionChain_Type) String() string {
	return proto.EnumName(ServiceFunctionChain_Type_name, int32(x))
}
func (ServiceFunctionChain_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_servicefunctionchain_a09a261c5a0e6415, []int{0, 0}
}

// ServiceFunctionChain is used to store definition of service function chain defined via CRD.
type ServiceFunctionChain struct {
	// name of the custom network
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// type of the custom network
	Network              string                                  `protobuf:"bytes,2,opt,name=network,proto3" json:"network,omitempty"`
	Chain                []*ServiceFunctionChain_ServiceFunction `protobuf:"bytes,3,rep,name=chain,proto3" json:"chain,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *ServiceFunctionChain) Reset()         { *m = ServiceFunctionChain{} }
func (m *ServiceFunctionChain) String() string { return proto.CompactTextString(m) }
func (*ServiceFunctionChain) ProtoMessage()    {}
func (*ServiceFunctionChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_servicefunctionchain_a09a261c5a0e6415, []int{0}
}
func (m *ServiceFunctionChain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceFunctionChain.Unmarshal(m, b)
}
func (m *ServiceFunctionChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceFunctionChain.Marshal(b, m, deterministic)
}
func (dst *ServiceFunctionChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceFunctionChain.Merge(dst, src)
}
func (m *ServiceFunctionChain) XXX_Size() int {
	return xxx_messageInfo_ServiceFunctionChain.Size(m)
}
func (m *ServiceFunctionChain) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceFunctionChain.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceFunctionChain proto.InternalMessageInfo

func (m *ServiceFunctionChain) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServiceFunctionChain) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *ServiceFunctionChain) GetChain() []*ServiceFunctionChain_ServiceFunction {
	if m != nil {
		return m.Chain
	}
	return nil
}

type ServiceFunctionChain_ServiceFunction struct {
	Name                 string                    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 ServiceFunctionChain_Type `protobuf:"varint,2,opt,name=type,proto3,enum=model.ServiceFunctionChain_Type" json:"type,omitempty"`
	PodSelector          map[string]string         `protobuf:"bytes,3,rep,name=pod_selector,json=podSelector,proto3" json:"pod_selector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	InputInterface       string                    `protobuf:"bytes,4,opt,name=input_interface,json=inputInterface,proto3" json:"input_interface,omitempty"`
	OutputInterface      string                    `protobuf:"bytes,5,opt,name=output_interface,json=outputInterface,proto3" json:"output_interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ServiceFunctionChain_ServiceFunction) Reset()         { *m = ServiceFunctionChain_ServiceFunction{} }
func (m *ServiceFunctionChain_ServiceFunction) String() string { return proto.CompactTextString(m) }
func (*ServiceFunctionChain_ServiceFunction) ProtoMessage()    {}
func (*ServiceFunctionChain_ServiceFunction) Descriptor() ([]byte, []int) {
	return fileDescriptor_servicefunctionchain_a09a261c5a0e6415, []int{0, 0}
}
func (m *ServiceFunctionChain_ServiceFunction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceFunctionChain_ServiceFunction.Unmarshal(m, b)
}
func (m *ServiceFunctionChain_ServiceFunction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceFunctionChain_ServiceFunction.Marshal(b, m, deterministic)
}
func (dst *ServiceFunctionChain_ServiceFunction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceFunctionChain_ServiceFunction.Merge(dst, src)
}
func (m *ServiceFunctionChain_ServiceFunction) XXX_Size() int {
	return xxx_messageInfo_ServiceFunctionChain_ServiceFunction.Size(m)
}
func (m *ServiceFunctionChain_ServiceFunction) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceFunctionChain_ServiceFunction.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceFunctionChain_ServiceFunction proto.InternalMessageInfo

func (m *ServiceFunctionChain_ServiceFunction) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServiceFunctionChain_ServiceFunction) GetType() ServiceFunctionChain_Type {
	if m != nil {
		return m.Type
	}
	return ServiceFunctionChain_ServiceFunctionPod
}

func (m *ServiceFunctionChain_ServiceFunction) GetPodSelector() map[string]string {
	if m != nil {
		return m.PodSelector
	}
	return nil
}

func (m *ServiceFunctionChain_ServiceFunction) GetInputInterface() string {
	if m != nil {
		return m.InputInterface
	}
	return ""
}

func (m *ServiceFunctionChain_ServiceFunction) GetOutputInterface() string {
	if m != nil {
		return m.OutputInterface
	}
	return ""
}

func init() {
	proto.RegisterType((*ServiceFunctionChain)(nil), "model.ServiceFunctionChain")
	proto.RegisterType((*ServiceFunctionChain_ServiceFunction)(nil), "model.ServiceFunctionChain.ServiceFunction")
	proto.RegisterMapType((map[string]string)(nil), "model.ServiceFunctionChain.ServiceFunction.PodSelectorEntry")
	proto.RegisterEnum("model.ServiceFunctionChain_Type", ServiceFunctionChain_Type_name, ServiceFunctionChain_Type_value)
}

func init() {
	proto.RegisterFile("servicefunctionchain.proto", fileDescriptor_servicefunctionchain_a09a261c5a0e6415)
}

var fileDescriptor_servicefunctionchain_a09a261c5a0e6415 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0xcd, 0x3f, 0xb5, 0xd3, 0xda, 0x84, 0xb1, 0x48, 0xe8, 0xa9, 0xf4, 0x62, 0x45, 0xc8,
	0xa1, 0x7a, 0x10, 0x11, 0x41, 0x44, 0xa1, 0x27, 0x43, 0xea, 0xbd, 0xc4, 0xec, 0x14, 0x43, 0xdb,
	0xdd, 0x90, 0x6e, 0x2a, 0x79, 0x3b, 0x1f, 0xc1, 0x47, 0x92, 0x6c, 0xb6, 0x15, 0x43, 0x11, 0xbc,
	0xcd, 0x7c, 0xf3, 0x9b, 0x99, 0x2f, 0x93, 0x85, 0xfe, 0x9a, 0xf2, 0x4d, 0x9a, 0xd0, 0xbc, 0xe0,
	0x89, 0x4c, 0x05, 0x4f, 0xde, 0xe3, 0x94, 0x07, 0x59, 0x2e, 0xa4, 0x40, 0x67, 0x25, 0x18, 0x2d,
	0x87, 0x9f, 0x36, 0xf4, 0xa6, 0x35, 0xf5, 0xac, 0xa9, 0xc7, 0x8a, 0x42, 0x04, 0x9b, 0xc7, 0x2b,
	0xf2, 0x8d, 0x81, 0x31, 0x6a, 0x45, 0x2a, 0x46, 0x1f, 0x8e, 0x38, 0xc9, 0x0f, 0x91, 0x2f, 0x7c,
	0x53, 0xc9, 0xdb, 0x14, 0x1f, 0xc0, 0x51, 0xc3, 0x7d, 0x6b, 0x60, 0x8d, 0xda, 0xe3, 0xcb, 0x40,
	0x4d, 0x0f, 0xf6, 0x4d, 0x6e, 0x8a, 0x51, 0xdd, 0xd9, 0xff, 0x32, 0xc1, 0x6d, 0x94, 0xf6, 0x9a,
	0xb8, 0x06, 0x5b, 0x96, 0x19, 0x29, 0x07, 0xdd, 0xf1, 0xe0, 0xaf, 0x4d, 0xaf, 0x65, 0x46, 0x91,
	0xa2, 0x71, 0x06, 0x9d, 0x4c, 0xb0, 0xd9, 0x9a, 0x96, 0x94, 0x48, 0x91, 0x6b, 0x9f, 0x77, 0xff,
	0xf0, 0x19, 0x84, 0x82, 0x4d, 0x75, 0xfb, 0x13, 0x97, 0x79, 0x19, 0xb5, 0xb3, 0x1f, 0x05, 0xcf,
	0xc1, 0x4d, 0x79, 0x56, 0xc8, 0x59, 0xca, 0x25, 0xe5, 0xf3, 0x38, 0x21, 0xdf, 0x56, 0xae, 0xbb,
	0x4a, 0x9e, 0x6c, 0x55, 0xbc, 0x00, 0x4f, 0x14, 0xf2, 0x37, 0xe9, 0x28, 0xd2, 0xad, 0xf5, 0x1d,
	0xda, 0xbf, 0x07, 0xaf, 0xb9, 0x14, 0x3d, 0xb0, 0x16, 0x54, 0xea, 0x8b, 0x54, 0x21, 0xf6, 0xc0,
	0xd9, 0xc4, 0xcb, 0x82, 0xf4, 0x3f, 0xa9, 0x93, 0x5b, 0xf3, 0xc6, 0x18, 0x32, 0xb0, 0xab, 0x13,
	0xe0, 0x19, 0x60, 0xe3, 0x63, 0x42, 0xc1, 0xbc, 0x03, 0x44, 0xe8, 0xee, 0x96, 0x4d, 0x2a, 0x97,
	0x9e, 0x81, 0xa7, 0xe0, 0xee, 0xb4, 0x17, 0xe5, 0xc7, 0x33, 0xb1, 0x03, 0xc7, 0xa1, 0x60, 0x35,
	0x62, 0xe1, 0x09, 0xb4, 0x42, 0xc1, 0x74, 0xd1, 0x7e, 0x3b, 0x54, 0x0f, 0xea, 0xea, 0x3b, 0x00,
	0x00, 0xff, 0xff, 0x45, 0x05, 0x0a, 0x0d, 0x6e, 0x02, 0x00, 0x00,
}