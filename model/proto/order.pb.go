// Code generated by protoc-gen-go. DO NOT EDIT.
// source: order.proto

package com_ta04_srv_order

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Order struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductId            int32    `protobuf:"varint,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	UserId               int32    `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Status               string   `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{0}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Order) GetProductId() int32 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *Order) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Order) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type User struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Username             string   `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Password             int32    `protobuf:"varint,5,opt,name=password,proto3" json:"password,omitempty"`
	PrimeNumber          int32    `protobuf:"varint,6,opt,name=prime_number,json=primeNumber,proto3" json:"prime_number,omitempty"`
	GeneratorValue       int32    `protobuf:"varint,7,opt,name=generator_value,json=generatorValue,proto3" json:"generator_value,omitempty"`
	EmailAddress         string   `protobuf:"bytes,8,opt,name=email_address,json=emailAddress,proto3" json:"email_address,omitempty"`
	PhoneNumber          string   `protobuf:"bytes,9,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	DateOfBirth          string   `protobuf:"bytes,10,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
	Address              string   `protobuf:"bytes,11,opt,name=address,proto3" json:"address,omitempty"`
	Role                 string   `protobuf:"bytes,12,opt,name=role,proto3" json:"role,omitempty"`
	Status               string   `protobuf:"bytes,13,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{1}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetPassword() int32 {
	if m != nil {
		return m.Password
	}
	return 0
}

func (m *User) GetPrimeNumber() int32 {
	if m != nil {
		return m.PrimeNumber
	}
	return 0
}

func (m *User) GetGeneratorValue() int32 {
	if m != nil {
		return m.GeneratorValue
	}
	return 0
}

func (m *User) GetEmailAddress() string {
	if m != nil {
		return m.EmailAddress
	}
	return ""
}

func (m *User) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *User) GetDateOfBirth() string {
	if m != nil {
		return m.DateOfBirth
	}
	return ""
}

func (m *User) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *User) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *User) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type GetAllOrdersRequest struct {
	UserId               int32    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllOrdersRequest) Reset()         { *m = GetAllOrdersRequest{} }
func (m *GetAllOrdersRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllOrdersRequest) ProtoMessage()    {}
func (*GetAllOrdersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{2}
}

func (m *GetAllOrdersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllOrdersRequest.Unmarshal(m, b)
}
func (m *GetAllOrdersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllOrdersRequest.Marshal(b, m, deterministic)
}
func (m *GetAllOrdersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllOrdersRequest.Merge(m, src)
}
func (m *GetAllOrdersRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllOrdersRequest.Size(m)
}
func (m *GetAllOrdersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllOrdersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllOrdersRequest proto.InternalMessageInfo

func (m *GetAllOrdersRequest) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *GetAllOrdersRequest) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type GetOneOrderRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOneOrderRequest) Reset()         { *m = GetOneOrderRequest{} }
func (m *GetOneOrderRequest) String() string { return proto.CompactTextString(m) }
func (*GetOneOrderRequest) ProtoMessage()    {}
func (*GetOneOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{3}
}

func (m *GetOneOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOneOrderRequest.Unmarshal(m, b)
}
func (m *GetOneOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOneOrderRequest.Marshal(b, m, deterministic)
}
func (m *GetOneOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOneOrderRequest.Merge(m, src)
}
func (m *GetOneOrderRequest) XXX_Size() int {
	return xxx_messageInfo_GetOneOrderRequest.Size(m)
}
func (m *GetOneOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOneOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetOneOrderRequest proto.InternalMessageInfo

func (m *GetOneOrderRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Response struct {
	Order                *Order   `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Orders               []*Order `protobuf:"bytes,3,rep,name=orders,proto3" json:"orders,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{4}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

func (m *Response) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Response) GetOrders() []*Order {
	if m != nil {
		return m.Orders
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{5}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Order)(nil), "com.ta04.srv.order.Order")
	proto.RegisterType((*User)(nil), "com.ta04.srv.order.User")
	proto.RegisterType((*GetAllOrdersRequest)(nil), "com.ta04.srv.order.GetAllOrdersRequest")
	proto.RegisterType((*GetOneOrderRequest)(nil), "com.ta04.srv.order.GetOneOrderRequest")
	proto.RegisterType((*Response)(nil), "com.ta04.srv.order.Response")
	proto.RegisterType((*Error)(nil), "com.ta04.srv.order.Error")
}

func init() {
	proto.RegisterFile("order.proto", fileDescriptor_cd01338c35d87077)
}

var fileDescriptor_cd01338c35d87077 = []byte{
	// 520 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x5f, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0xe9, 0xff, 0x66, 0xd2, 0x16, 0xc9, 0x48, 0x60, 0x16, 0x90, 0x4a, 0x40, 0xec, 0x3e,
	0x05, 0x58, 0xe0, 0x00, 0x0b, 0x82, 0xd5, 0x0a, 0x69, 0x2b, 0x05, 0x15, 0x1e, 0x23, 0xb7, 0x9e,
	0x76, 0x23, 0x25, 0x71, 0xb0, 0x9d, 0x72, 0x05, 0x6e, 0xc1, 0xb5, 0x38, 0x0e, 0xf2, 0x24, 0xa9,
	0xba, 0x4b, 0x0b, 0x3c, 0xf0, 0xe6, 0xf9, 0x7d, 0x93, 0xef, 0xb3, 0x66, 0xac, 0x80, 0xaf, 0xb4,
	0x44, 0x1d, 0x16, 0x5a, 0x59, 0xc5, 0xd8, 0x52, 0x65, 0xa1, 0x15, 0x2f, 0x5e, 0x87, 0x46, 0x6f,
	0x42, 0x52, 0x82, 0x35, 0xf4, 0x66, 0xee, 0xc0, 0x26, 0xd0, 0x4e, 0x24, 0x6f, 0x4d, 0x5b, 0x27,
	0xbd, 0xa8, 0x9d, 0x48, 0xf6, 0x08, 0xa0, 0xd0, 0x4a, 0x96, 0x4b, 0x1b, 0x27, 0x92, 0xb7, 0x89,
	0x7b, 0x35, 0xb9, 0x90, 0xec, 0x1e, 0x0c, 0x4a, 0x83, 0xda, 0x69, 0x1d, 0xd2, 0xfa, 0xae, 0xbc,
	0x90, 0xec, 0x2e, 0xf4, 0x8d, 0x15, 0xb6, 0x34, 0xbc, 0x3b, 0x6d, 0x9d, 0x78, 0x51, 0x5d, 0x05,
	0xdf, 0x3b, 0xd0, 0x9d, 0x9b, 0xfd, 0x41, 0xab, 0x44, 0x1b, 0x1b, 0xe7, 0x22, 0x43, 0x0a, 0xf2,
	0x22, 0x8f, 0xc8, 0xa5, 0xc8, 0x90, 0x3d, 0x00, 0x2f, 0x15, 0x8d, 0xda, 0x21, 0x75, 0xe8, 0x00,
	0x89, 0x47, 0x30, 0x74, 0xb1, 0xa4, 0x55, 0x71, 0xdb, 0xda, 0x69, 0x85, 0x30, 0xe6, 0x9b, 0xd2,
	0x92, 0xf7, 0x28, 0x6d, 0x5b, 0xb3, 0xc7, 0x30, 0x2a, 0x74, 0x92, 0x61, 0x9c, 0x97, 0xd9, 0x02,
	0x35, 0xef, 0x93, 0xee, 0x13, 0xbb, 0x24, 0xc4, 0x8e, 0xe1, 0xf6, 0x1a, 0x73, 0xd4, 0xc2, 0x2a,
	0x1d, 0x6f, 0x44, 0x5a, 0x22, 0x1f, 0x50, 0xd7, 0x64, 0x8b, 0x3f, 0x3b, 0xca, 0x9e, 0xc0, 0x18,
	0x33, 0x91, 0xa4, 0xb1, 0x90, 0x52, 0xa3, 0x31, 0x7c, 0x48, 0x17, 0x19, 0x11, 0x3c, 0xab, 0x18,
	0x05, 0x5e, 0xa9, 0x7c, 0x1b, 0xe8, 0x51, 0x8f, 0x4f, 0xac, 0x0e, 0x0c, 0x60, 0x2c, 0x85, 0xc5,
	0x58, 0xad, 0xe2, 0x45, 0xa2, 0xed, 0x15, 0x87, 0xaa, 0xc7, 0xc1, 0xd9, 0xea, 0xad, 0x43, 0x8c,
	0xc3, 0xa0, 0x49, 0xf1, 0x49, 0x6d, 0x4a, 0xc6, 0xa0, 0xab, 0x55, 0x8a, 0x7c, 0x44, 0x98, 0xce,
	0x3b, 0xab, 0x18, 0x5f, 0x5b, 0xc5, 0x07, 0xb8, 0x73, 0x8e, 0xf6, 0x2c, 0x4d, 0x69, 0xf3, 0x26,
	0xc2, 0xaf, 0x25, 0x1a, 0xbb, 0xbb, 0xd2, 0xd6, 0x81, 0x95, 0xb6, 0xaf, 0xf9, 0x3c, 0x05, 0x76,
	0x8e, 0x76, 0x96, 0x23, 0xf9, 0x34, 0x36, 0x37, 0xf6, 0x1b, 0xfc, 0x68, 0xc1, 0x30, 0x42, 0x53,
	0xa8, 0xdc, 0x20, 0x7b, 0x0e, 0x3d, 0x7a, 0x77, 0xa4, 0xfb, 0xa7, 0xf7, 0xc3, 0xdf, 0x9f, 0x64,
	0x58, 0xb9, 0x55, 0x7d, 0xee, 0x03, 0xd4, 0x5a, 0x69, 0x8a, 0x3e, 0xf0, 0xc1, 0x7b, 0xd7, 0x10,
	0x55, 0x7d, 0xec, 0x25, 0xf4, 0x89, 0x1a, 0xde, 0x99, 0x76, 0xfe, 0x1c, 0x51, 0x37, 0x06, 0x6f,
	0xa0, 0x47, 0x16, 0x6e, 0x88, 0x4b, 0x25, 0xb1, 0xbe, 0x3c, 0x9d, 0xdd, 0xc8, 0x33, 0x34, 0x46,
	0xac, 0x9b, 0xb7, 0xd9, 0x94, 0xa7, 0x3f, 0xdb, 0x30, 0x22, 0xa3, 0x4f, 0xa8, 0x37, 0xc9, 0x12,
	0xd9, 0x17, 0x18, 0xed, 0xce, 0x95, 0x1d, 0xef, 0x8b, 0xde, 0x33, 0xf9, 0xa3, 0x87, 0xfb, 0x1a,
	0x9b, 0x99, 0x05, 0xb7, 0xd8, 0x1c, 0xfc, 0x9d, 0x41, 0xb3, 0x67, 0x07, 0x7c, 0x6f, 0x6c, 0xe2,
	0xaf, 0xb6, 0x1f, 0x61, 0xf2, 0x4e, 0xa3, 0x7b, 0x5e, 0x8d, 0xf3, 0xe1, 0x61, 0xfd, 0x8b, 0xd9,
	0xbc, 0x90, 0xff, 0xc7, 0x6c, 0xd1, 0xa7, 0x1f, 0xd6, 0xab, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xea, 0x23, 0xb7, 0x7b, 0xbf, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for OrderService service

type OrderServiceClient interface {
	GetAllOrders(ctx context.Context, in *GetAllOrdersRequest, opts ...client.CallOption) (*Response, error)
	GetOneOrder(ctx context.Context, in *GetOneOrderRequest, opts ...client.CallOption) (*Response, error)
	CreateOneOrder(ctx context.Context, in *Order, opts ...client.CallOption) (*Response, error)
	UpdateOneOrder(ctx context.Context, in *Order, opts ...client.CallOption) (*Response, error)
}

type orderServiceClient struct {
	c           client.Client
	serviceName string
}

func NewOrderServiceClient(serviceName string, c client.Client) OrderServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "com.ta04.srv.order"
	}
	return &orderServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *orderServiceClient) GetAllOrders(ctx context.Context, in *GetAllOrdersRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "OrderService.GetAllOrders", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOneOrder(ctx context.Context, in *GetOneOrderRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "OrderService.GetOneOrder", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CreateOneOrder(ctx context.Context, in *Order, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "OrderService.CreateOneOrder", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) UpdateOneOrder(ctx context.Context, in *Order, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "OrderService.UpdateOneOrder", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OrderService service

type OrderServiceHandler interface {
	GetAllOrders(context.Context, *GetAllOrdersRequest, *Response) error
	GetOneOrder(context.Context, *GetOneOrderRequest, *Response) error
	CreateOneOrder(context.Context, *Order, *Response) error
	UpdateOneOrder(context.Context, *Order, *Response) error
}

func RegisterOrderServiceHandler(s server.Server, hdlr OrderServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&OrderService{hdlr}, opts...))
}

type OrderService struct {
	OrderServiceHandler
}

func (h *OrderService) GetAllOrders(ctx context.Context, in *GetAllOrdersRequest, out *Response) error {
	return h.OrderServiceHandler.GetAllOrders(ctx, in, out)
}

func (h *OrderService) GetOneOrder(ctx context.Context, in *GetOneOrderRequest, out *Response) error {
	return h.OrderServiceHandler.GetOneOrder(ctx, in, out)
}

func (h *OrderService) CreateOneOrder(ctx context.Context, in *Order, out *Response) error {
	return h.OrderServiceHandler.CreateOneOrder(ctx, in, out)
}

func (h *OrderService) UpdateOneOrder(ctx context.Context, in *Order, out *Response) error {
	return h.OrderServiceHandler.UpdateOneOrder(ctx, in, out)
}
