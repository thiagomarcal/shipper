// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/consignment/consignment.proto

/*
Package shipper_consignment is a generated protocol buffer package.

It is generated from these files:
	proto/consignment/consignment.proto

It has these top-level messages:
	Consignment
	Container
	GetRequest
	Response
*/
package shipper_consignment

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Consignment struct {
	Id          string       `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Description string       `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Weight      int32        `protobuf:"varint,3,opt,name=weight" json:"weight,omitempty"`
	Containers  []*Container `protobuf:"bytes,4,rep,name=containers" json:"containers,omitempty"`
	VesselId    string       `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId" json:"vessel_id,omitempty"`
}

func (m *Consignment) Reset()                    { *m = Consignment{} }
func (m *Consignment) String() string            { return proto.CompactTextString(m) }
func (*Consignment) ProtoMessage()               {}
func (*Consignment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Consignment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Consignment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Consignment) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Consignment) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Consignment) GetVesselId() string {
	if m != nil {
		return m.VesselId
	}
	return ""
}

type Container struct {
	Id         string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	CustomerId string `protobuf:"bytes,2,opt,name=customer_id,json=customerId" json:"customer_id,omitempty"`
	Origin     string `protobuf:"bytes,3,opt,name=origin" json:"origin,omitempty"`
	UserId     string `protobuf:"bytes,4,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *Container) Reset()                    { *m = Container{} }
func (m *Container) String() string            { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()               {}
func (*Container) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Container) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Container) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type GetRequest struct {
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Response struct {
	Created      bool           `protobuf:"varint,1,opt,name=created" json:"created,omitempty"`
	Consignment  *Consignment   `protobuf:"bytes,2,opt,name=consignment" json:"consignment,omitempty"`
	Consignments []*Consignment `protobuf:"bytes,3,rep,name=consignments" json:"consignments,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetConsignment() *Consignment {
	if m != nil {
		return m.Consignment
	}
	return nil
}

func (m *Response) GetConsignments() []*Consignment {
	if m != nil {
		return m.Consignments
	}
	return nil
}

func init() {
	proto.RegisterType((*Consignment)(nil), "shipper.consignment.Consignment")
	proto.RegisterType((*Container)(nil), "shipper.consignment.Container")
	proto.RegisterType((*GetRequest)(nil), "shipper.consignment.GetRequest")
	proto.RegisterType((*Response)(nil), "shipper.consignment.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ShippingService service

type ShippingServiceClient interface {
	CreateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error)
	GetConsignments(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Response, error)
}

type shippingServiceClient struct {
	c           client.Client
	serviceName string
}

func NewShippingServiceClient(serviceName string, c client.Client) ShippingServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "shipper.consignment"
	}
	return &shippingServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *shippingServiceClient) CreateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ShippingService.CreateConsignment", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingServiceClient) GetConsignments(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ShippingService.GetConsignments", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ShippingService service

type ShippingServiceHandler interface {
	CreateConsignment(context.Context, *Consignment, *Response) error
	GetConsignments(context.Context, *GetRequest, *Response) error
}

func RegisterShippingServiceHandler(s server.Server, hdlr ShippingServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&ShippingService{hdlr}, opts...))
}

type ShippingService struct {
	ShippingServiceHandler
}

func (h *ShippingService) CreateConsignment(ctx context.Context, in *Consignment, out *Response) error {
	return h.ShippingServiceHandler.CreateConsignment(ctx, in, out)
}

func (h *ShippingService) GetConsignments(ctx context.Context, in *GetRequest, out *Response) error {
	return h.ShippingServiceHandler.GetConsignments(ctx, in, out)
}

func init() { proto.RegisterFile("proto/consignment/consignment.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xed, 0x7e, 0xf7, 0x75, 0x38, 0x8c, 0xa0, 0x41, 0xd1, 0x95, 0x7a, 0xd9, 0xa9, 0xc2,
	0xbc, 0x7b, 0x70, 0xc2, 0xd8, 0x35, 0x03, 0xaf, 0x32, 0x9b, 0x47, 0x17, 0x70, 0x49, 0x4d, 0xb2,
	0xf9, 0x5f, 0x79, 0xf2, 0xec, 0xdf, 0x26, 0x4d, 0x57, 0x17, 0x71, 0xb2, 0x5b, 0xbe, 0xef, 0xe5,
	0xf3, 0xfa, 0xe9, 0x6b, 0xe1, 0xa6, 0xd0, 0xca, 0xaa, 0xdb, 0x4c, 0x49, 0x23, 0x72, 0xb9, 0x42,
	0x69, 0xfd, 0x73, 0xea, 0xba, 0xe4, 0xd4, 0x2c, 0x45, 0x51, 0xa0, 0x4e, 0xbd, 0x56, 0xf2, 0x19,
	0x40, 0x34, 0xd9, 0x65, 0x72, 0x0c, 0x0d, 0xc1, 0x69, 0x10, 0x07, 0xa3, 0x90, 0x35, 0x04, 0x27,
	0x31, 0x44, 0x1c, 0x4d, 0xa6, 0x45, 0x61, 0x85, 0x92, 0xb4, 0xe1, 0x1a, 0x7e, 0x89, 0x9c, 0x41,
	0xe7, 0x1d, 0x45, 0xbe, 0xb4, 0xb4, 0x19, 0x07, 0xa3, 0x36, 0xdb, 0x26, 0x72, 0x0f, 0x90, 0x29,
	0x69, 0x17, 0x42, 0xa2, 0x36, 0xb4, 0x15, 0x37, 0x47, 0xd1, 0xf8, 0x3a, 0xdd, 0xe3, 0x90, 0x4e,
	0xea, 0x6b, 0xcc, 0x23, 0xc8, 0x25, 0x84, 0x1b, 0x34, 0x06, 0x5f, 0x9f, 0x05, 0xa7, 0x6d, 0xf7,
	0xdc, 0x5e, 0x55, 0x98, 0xf1, 0x64, 0x05, 0xe1, 0x0f, 0xf5, 0xc7, 0x79, 0x08, 0x51, 0xb6, 0x36,
	0x56, 0xad, 0x50, 0x97, 0x6c, 0xe5, 0x0c, 0x75, 0x69, 0xc6, 0x4b, 0x65, 0xa5, 0x45, 0x2e, 0xa4,
	0x53, 0x0e, 0xd9, 0x36, 0x91, 0x73, 0xe8, 0xae, 0x4d, 0x05, 0xb5, 0xaa, 0x46, 0x19, 0x67, 0x3c,
	0xe9, 0x03, 0x4c, 0xd1, 0x32, 0x7c, 0x5b, 0xa3, 0xb1, 0xc9, 0x47, 0x00, 0x3d, 0x86, 0xa6, 0x50,
	0xd2, 0x20, 0xa1, 0xd0, 0xcd, 0x34, 0x2e, 0x2c, 0x56, 0x06, 0x3d, 0x56, 0x47, 0xf2, 0x00, 0x91,
	0xf7, 0x96, 0x4e, 0x23, 0x1a, 0xc7, 0xff, 0x6d, 0xa0, 0x3e, 0x33, 0x1f, 0x22, 0x8f, 0xd0, 0xf7,
	0xa2, 0xa1, 0x4d, 0xb7, 0xc6, 0xc3, 0x43, 0x7e, 0x51, 0xe3, 0xaf, 0x00, 0x06, 0xf3, 0x92, 0x10,
	0x32, 0x9f, 0xa3, 0xde, 0x88, 0x0c, 0xc9, 0x13, 0x9c, 0x4c, 0x9c, 0xa8, 0xff, 0xf5, 0x0f, 0x0e,
	0xbe, 0xb8, 0xda, 0x7b, 0xa3, 0xde, 0x46, 0x72, 0x44, 0xe6, 0x30, 0x98, 0xa2, 0xf5, 0x10, 0x43,
	0x86, 0x7b, 0x99, 0xdd, 0x42, 0x0f, 0x0e, 0x7d, 0xe9, 0xb8, 0x3f, 0xf8, 0xee, 0x3b, 0x00, 0x00,
	0xff, 0xff, 0xcc, 0x8b, 0x80, 0xc0, 0xe8, 0x02, 0x00, 0x00,
}
