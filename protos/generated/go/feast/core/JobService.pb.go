// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feast/core/JobService.proto

package core // import "github.com/gojek/feast/protos/generated/go/feast/core"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import specs "github.com/gojek/feast/protos/generated/go/feast/specs"
import empty "github.com/golang/protobuf/ptypes/empty"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type JobServiceTypes struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobServiceTypes) Reset()         { *m = JobServiceTypes{} }
func (m *JobServiceTypes) String() string { return proto.CompactTextString(m) }
func (*JobServiceTypes) ProtoMessage()    {}
func (*JobServiceTypes) Descriptor() ([]byte, []int) {
	return fileDescriptor_JobService_edcd183b773c9f62, []int{0}
}
func (m *JobServiceTypes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobServiceTypes.Unmarshal(m, b)
}
func (m *JobServiceTypes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobServiceTypes.Marshal(b, m, deterministic)
}
func (dst *JobServiceTypes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobServiceTypes.Merge(dst, src)
}
func (m *JobServiceTypes) XXX_Size() int {
	return xxx_messageInfo_JobServiceTypes.Size(m)
}
func (m *JobServiceTypes) XXX_DiscardUnknown() {
	xxx_messageInfo_JobServiceTypes.DiscardUnknown(m)
}

var xxx_messageInfo_JobServiceTypes proto.InternalMessageInfo

type JobServiceTypes_SubmitImportJobRequest struct {
	ImportSpec           *specs.ImportSpec `protobuf:"bytes,1,opt,name=importSpec,proto3" json:"importSpec,omitempty"`
	Name                 string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *JobServiceTypes_SubmitImportJobRequest) Reset() {
	*m = JobServiceTypes_SubmitImportJobRequest{}
}
func (m *JobServiceTypes_SubmitImportJobRequest) String() string { return proto.CompactTextString(m) }
func (*JobServiceTypes_SubmitImportJobRequest) ProtoMessage()    {}
func (*JobServiceTypes_SubmitImportJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_JobService_edcd183b773c9f62, []int{0, 0}
}
func (m *JobServiceTypes_SubmitImportJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobServiceTypes_SubmitImportJobRequest.Unmarshal(m, b)
}
func (m *JobServiceTypes_SubmitImportJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobServiceTypes_SubmitImportJobRequest.Marshal(b, m, deterministic)
}
func (dst *JobServiceTypes_SubmitImportJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobServiceTypes_SubmitImportJobRequest.Merge(dst, src)
}
func (m *JobServiceTypes_SubmitImportJobRequest) XXX_Size() int {
	return xxx_messageInfo_JobServiceTypes_SubmitImportJobRequest.Size(m)
}
func (m *JobServiceTypes_SubmitImportJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JobServiceTypes_SubmitImportJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JobServiceTypes_SubmitImportJobRequest proto.InternalMessageInfo

func (m *JobServiceTypes_SubmitImportJobRequest) GetImportSpec() *specs.ImportSpec {
	if m != nil {
		return m.ImportSpec
	}
	return nil
}

func (m *JobServiceTypes_SubmitImportJobRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type JobServiceTypes_SubmitImportJobResponse struct {
	JobId                string   `protobuf:"bytes,1,opt,name=jobId,proto3" json:"jobId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobServiceTypes_SubmitImportJobResponse) Reset() {
	*m = JobServiceTypes_SubmitImportJobResponse{}
}
func (m *JobServiceTypes_SubmitImportJobResponse) String() string { return proto.CompactTextString(m) }
func (*JobServiceTypes_SubmitImportJobResponse) ProtoMessage()    {}
func (*JobServiceTypes_SubmitImportJobResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_JobService_edcd183b773c9f62, []int{0, 1}
}
func (m *JobServiceTypes_SubmitImportJobResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobServiceTypes_SubmitImportJobResponse.Unmarshal(m, b)
}
func (m *JobServiceTypes_SubmitImportJobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobServiceTypes_SubmitImportJobResponse.Marshal(b, m, deterministic)
}
func (dst *JobServiceTypes_SubmitImportJobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobServiceTypes_SubmitImportJobResponse.Merge(dst, src)
}
func (m *JobServiceTypes_SubmitImportJobResponse) XXX_Size() int {
	return xxx_messageInfo_JobServiceTypes_SubmitImportJobResponse.Size(m)
}
func (m *JobServiceTypes_SubmitImportJobResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JobServiceTypes_SubmitImportJobResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JobServiceTypes_SubmitImportJobResponse proto.InternalMessageInfo

func (m *JobServiceTypes_SubmitImportJobResponse) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

type JobServiceTypes_ListJobsResponse struct {
	Jobs                 []*JobServiceTypes_JobDetail `protobuf:"bytes,1,rep,name=jobs,proto3" json:"jobs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *JobServiceTypes_ListJobsResponse) Reset()         { *m = JobServiceTypes_ListJobsResponse{} }
func (m *JobServiceTypes_ListJobsResponse) String() string { return proto.CompactTextString(m) }
func (*JobServiceTypes_ListJobsResponse) ProtoMessage()    {}
func (*JobServiceTypes_ListJobsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_JobService_edcd183b773c9f62, []int{0, 2}
}
func (m *JobServiceTypes_ListJobsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobServiceTypes_ListJobsResponse.Unmarshal(m, b)
}
func (m *JobServiceTypes_ListJobsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobServiceTypes_ListJobsResponse.Marshal(b, m, deterministic)
}
func (dst *JobServiceTypes_ListJobsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobServiceTypes_ListJobsResponse.Merge(dst, src)
}
func (m *JobServiceTypes_ListJobsResponse) XXX_Size() int {
	return xxx_messageInfo_JobServiceTypes_ListJobsResponse.Size(m)
}
func (m *JobServiceTypes_ListJobsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JobServiceTypes_ListJobsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JobServiceTypes_ListJobsResponse proto.InternalMessageInfo

func (m *JobServiceTypes_ListJobsResponse) GetJobs() []*JobServiceTypes_JobDetail {
	if m != nil {
		return m.Jobs
	}
	return nil
}

type JobServiceTypes_GetJobRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobServiceTypes_GetJobRequest) Reset()         { *m = JobServiceTypes_GetJobRequest{} }
func (m *JobServiceTypes_GetJobRequest) String() string { return proto.CompactTextString(m) }
func (*JobServiceTypes_GetJobRequest) ProtoMessage()    {}
func (*JobServiceTypes_GetJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_JobService_edcd183b773c9f62, []int{0, 3}
}
func (m *JobServiceTypes_GetJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobServiceTypes_GetJobRequest.Unmarshal(m, b)
}
func (m *JobServiceTypes_GetJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobServiceTypes_GetJobRequest.Marshal(b, m, deterministic)
}
func (dst *JobServiceTypes_GetJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobServiceTypes_GetJobRequest.Merge(dst, src)
}
func (m *JobServiceTypes_GetJobRequest) XXX_Size() int {
	return xxx_messageInfo_JobServiceTypes_GetJobRequest.Size(m)
}
func (m *JobServiceTypes_GetJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JobServiceTypes_GetJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JobServiceTypes_GetJobRequest proto.InternalMessageInfo

func (m *JobServiceTypes_GetJobRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type JobServiceTypes_GetJobResponse struct {
	Job                  *JobServiceTypes_JobDetail `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *JobServiceTypes_GetJobResponse) Reset()         { *m = JobServiceTypes_GetJobResponse{} }
func (m *JobServiceTypes_GetJobResponse) String() string { return proto.CompactTextString(m) }
func (*JobServiceTypes_GetJobResponse) ProtoMessage()    {}
func (*JobServiceTypes_GetJobResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_JobService_edcd183b773c9f62, []int{0, 4}
}
func (m *JobServiceTypes_GetJobResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobServiceTypes_GetJobResponse.Unmarshal(m, b)
}
func (m *JobServiceTypes_GetJobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobServiceTypes_GetJobResponse.Marshal(b, m, deterministic)
}
func (dst *JobServiceTypes_GetJobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobServiceTypes_GetJobResponse.Merge(dst, src)
}
func (m *JobServiceTypes_GetJobResponse) XXX_Size() int {
	return xxx_messageInfo_JobServiceTypes_GetJobResponse.Size(m)
}
func (m *JobServiceTypes_GetJobResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JobServiceTypes_GetJobResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JobServiceTypes_GetJobResponse proto.InternalMessageInfo

func (m *JobServiceTypes_GetJobResponse) GetJob() *JobServiceTypes_JobDetail {
	if m != nil {
		return m.Job
	}
	return nil
}

type JobServiceTypes_AbortJobRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobServiceTypes_AbortJobRequest) Reset()         { *m = JobServiceTypes_AbortJobRequest{} }
func (m *JobServiceTypes_AbortJobRequest) String() string { return proto.CompactTextString(m) }
func (*JobServiceTypes_AbortJobRequest) ProtoMessage()    {}
func (*JobServiceTypes_AbortJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_JobService_edcd183b773c9f62, []int{0, 5}
}
func (m *JobServiceTypes_AbortJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobServiceTypes_AbortJobRequest.Unmarshal(m, b)
}
func (m *JobServiceTypes_AbortJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobServiceTypes_AbortJobRequest.Marshal(b, m, deterministic)
}
func (dst *JobServiceTypes_AbortJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobServiceTypes_AbortJobRequest.Merge(dst, src)
}
func (m *JobServiceTypes_AbortJobRequest) XXX_Size() int {
	return xxx_messageInfo_JobServiceTypes_AbortJobRequest.Size(m)
}
func (m *JobServiceTypes_AbortJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JobServiceTypes_AbortJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JobServiceTypes_AbortJobRequest proto.InternalMessageInfo

func (m *JobServiceTypes_AbortJobRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type JobServiceTypes_AbortJobResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobServiceTypes_AbortJobResponse) Reset()         { *m = JobServiceTypes_AbortJobResponse{} }
func (m *JobServiceTypes_AbortJobResponse) String() string { return proto.CompactTextString(m) }
func (*JobServiceTypes_AbortJobResponse) ProtoMessage()    {}
func (*JobServiceTypes_AbortJobResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_JobService_edcd183b773c9f62, []int{0, 6}
}
func (m *JobServiceTypes_AbortJobResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobServiceTypes_AbortJobResponse.Unmarshal(m, b)
}
func (m *JobServiceTypes_AbortJobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobServiceTypes_AbortJobResponse.Marshal(b, m, deterministic)
}
func (dst *JobServiceTypes_AbortJobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobServiceTypes_AbortJobResponse.Merge(dst, src)
}
func (m *JobServiceTypes_AbortJobResponse) XXX_Size() int {
	return xxx_messageInfo_JobServiceTypes_AbortJobResponse.Size(m)
}
func (m *JobServiceTypes_AbortJobResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JobServiceTypes_AbortJobResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JobServiceTypes_AbortJobResponse proto.InternalMessageInfo

func (m *JobServiceTypes_AbortJobResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// Expanded view of a given job. Returns job information, as well
// as latest metrics.
type JobServiceTypes_JobDetail struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ExtId                string               `protobuf:"bytes,2,opt,name=extId,proto3" json:"extId,omitempty"`
	Type                 string               `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Runner               string               `protobuf:"bytes,4,opt,name=runner,proto3" json:"runner,omitempty"`
	Status               string               `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Entities             []string             `protobuf:"bytes,6,rep,name=entities,proto3" json:"entities,omitempty"`
	Features             []string             `protobuf:"bytes,7,rep,name=features,proto3" json:"features,omitempty"`
	Metrics              map[string]float64   `protobuf:"bytes,8,rep,name=metrics,proto3" json:"metrics,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	LastUpdated          *timestamp.Timestamp `protobuf:"bytes,9,opt,name=lastUpdated,proto3" json:"lastUpdated,omitempty"`
	Created              *timestamp.Timestamp `protobuf:"bytes,10,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *JobServiceTypes_JobDetail) Reset()         { *m = JobServiceTypes_JobDetail{} }
func (m *JobServiceTypes_JobDetail) String() string { return proto.CompactTextString(m) }
func (*JobServiceTypes_JobDetail) ProtoMessage()    {}
func (*JobServiceTypes_JobDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_JobService_edcd183b773c9f62, []int{0, 7}
}
func (m *JobServiceTypes_JobDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobServiceTypes_JobDetail.Unmarshal(m, b)
}
func (m *JobServiceTypes_JobDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobServiceTypes_JobDetail.Marshal(b, m, deterministic)
}
func (dst *JobServiceTypes_JobDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobServiceTypes_JobDetail.Merge(dst, src)
}
func (m *JobServiceTypes_JobDetail) XXX_Size() int {
	return xxx_messageInfo_JobServiceTypes_JobDetail.Size(m)
}
func (m *JobServiceTypes_JobDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_JobServiceTypes_JobDetail.DiscardUnknown(m)
}

var xxx_messageInfo_JobServiceTypes_JobDetail proto.InternalMessageInfo

func (m *JobServiceTypes_JobDetail) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *JobServiceTypes_JobDetail) GetExtId() string {
	if m != nil {
		return m.ExtId
	}
	return ""
}

func (m *JobServiceTypes_JobDetail) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *JobServiceTypes_JobDetail) GetRunner() string {
	if m != nil {
		return m.Runner
	}
	return ""
}

func (m *JobServiceTypes_JobDetail) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *JobServiceTypes_JobDetail) GetEntities() []string {
	if m != nil {
		return m.Entities
	}
	return nil
}

func (m *JobServiceTypes_JobDetail) GetFeatures() []string {
	if m != nil {
		return m.Features
	}
	return nil
}

func (m *JobServiceTypes_JobDetail) GetMetrics() map[string]float64 {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *JobServiceTypes_JobDetail) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

func (m *JobServiceTypes_JobDetail) GetCreated() *timestamp.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func init() {
	proto.RegisterType((*JobServiceTypes)(nil), "feast.core.JobServiceTypes")
	proto.RegisterType((*JobServiceTypes_SubmitImportJobRequest)(nil), "feast.core.JobServiceTypes.SubmitImportJobRequest")
	proto.RegisterType((*JobServiceTypes_SubmitImportJobResponse)(nil), "feast.core.JobServiceTypes.SubmitImportJobResponse")
	proto.RegisterType((*JobServiceTypes_ListJobsResponse)(nil), "feast.core.JobServiceTypes.ListJobsResponse")
	proto.RegisterType((*JobServiceTypes_GetJobRequest)(nil), "feast.core.JobServiceTypes.GetJobRequest")
	proto.RegisterType((*JobServiceTypes_GetJobResponse)(nil), "feast.core.JobServiceTypes.GetJobResponse")
	proto.RegisterType((*JobServiceTypes_AbortJobRequest)(nil), "feast.core.JobServiceTypes.AbortJobRequest")
	proto.RegisterType((*JobServiceTypes_AbortJobResponse)(nil), "feast.core.JobServiceTypes.AbortJobResponse")
	proto.RegisterType((*JobServiceTypes_JobDetail)(nil), "feast.core.JobServiceTypes.JobDetail")
	proto.RegisterMapType((map[string]float64)(nil), "feast.core.JobServiceTypes.JobDetail.MetricsEntry")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// JobServiceClient is the client API for JobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JobServiceClient interface {
	// Submit a job to feast to run. Returns the job id.
	SubmitJob(ctx context.Context, in *JobServiceTypes_SubmitImportJobRequest, opts ...grpc.CallOption) (*JobServiceTypes_SubmitImportJobResponse, error)
	// List all jobs submitted to feast.
	ListJobs(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*JobServiceTypes_ListJobsResponse, error)
	// Get Job with ID
	GetJob(ctx context.Context, in *JobServiceTypes_GetJobRequest, opts ...grpc.CallOption) (*JobServiceTypes_GetJobResponse, error)
	// Abort job with given ID
	AbortJob(ctx context.Context, in *JobServiceTypes_AbortJobRequest, opts ...grpc.CallOption) (*JobServiceTypes_AbortJobResponse, error)
}

type jobServiceClient struct {
	cc *grpc.ClientConn
}

func NewJobServiceClient(cc *grpc.ClientConn) JobServiceClient {
	return &jobServiceClient{cc}
}

func (c *jobServiceClient) SubmitJob(ctx context.Context, in *JobServiceTypes_SubmitImportJobRequest, opts ...grpc.CallOption) (*JobServiceTypes_SubmitImportJobResponse, error) {
	out := new(JobServiceTypes_SubmitImportJobResponse)
	err := c.cc.Invoke(ctx, "/feast.core.JobService/SubmitJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) ListJobs(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*JobServiceTypes_ListJobsResponse, error) {
	out := new(JobServiceTypes_ListJobsResponse)
	err := c.cc.Invoke(ctx, "/feast.core.JobService/ListJobs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) GetJob(ctx context.Context, in *JobServiceTypes_GetJobRequest, opts ...grpc.CallOption) (*JobServiceTypes_GetJobResponse, error) {
	out := new(JobServiceTypes_GetJobResponse)
	err := c.cc.Invoke(ctx, "/feast.core.JobService/GetJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) AbortJob(ctx context.Context, in *JobServiceTypes_AbortJobRequest, opts ...grpc.CallOption) (*JobServiceTypes_AbortJobResponse, error) {
	out := new(JobServiceTypes_AbortJobResponse)
	err := c.cc.Invoke(ctx, "/feast.core.JobService/AbortJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobServiceServer is the server API for JobService service.
type JobServiceServer interface {
	// Submit a job to feast to run. Returns the job id.
	SubmitJob(context.Context, *JobServiceTypes_SubmitImportJobRequest) (*JobServiceTypes_SubmitImportJobResponse, error)
	// List all jobs submitted to feast.
	ListJobs(context.Context, *empty.Empty) (*JobServiceTypes_ListJobsResponse, error)
	// Get Job with ID
	GetJob(context.Context, *JobServiceTypes_GetJobRequest) (*JobServiceTypes_GetJobResponse, error)
	// Abort job with given ID
	AbortJob(context.Context, *JobServiceTypes_AbortJobRequest) (*JobServiceTypes_AbortJobResponse, error)
}

func RegisterJobServiceServer(s *grpc.Server, srv JobServiceServer) {
	s.RegisterService(&_JobService_serviceDesc, srv)
}

func _JobService_SubmitJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobServiceTypes_SubmitImportJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).SubmitJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feast.core.JobService/SubmitJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).SubmitJob(ctx, req.(*JobServiceTypes_SubmitImportJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_ListJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).ListJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feast.core.JobService/ListJobs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).ListJobs(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_GetJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobServiceTypes_GetJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).GetJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feast.core.JobService/GetJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).GetJob(ctx, req.(*JobServiceTypes_GetJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_AbortJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobServiceTypes_AbortJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).AbortJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feast.core.JobService/AbortJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).AbortJob(ctx, req.(*JobServiceTypes_AbortJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _JobService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "feast.core.JobService",
	HandlerType: (*JobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitJob",
			Handler:    _JobService_SubmitJob_Handler,
		},
		{
			MethodName: "ListJobs",
			Handler:    _JobService_ListJobs_Handler,
		},
		{
			MethodName: "GetJob",
			Handler:    _JobService_GetJob_Handler,
		},
		{
			MethodName: "AbortJob",
			Handler:    _JobService_AbortJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feast/core/JobService.proto",
}

func init() {
	proto.RegisterFile("feast/core/JobService.proto", fileDescriptor_JobService_edcd183b773c9f62)
}

var fileDescriptor_JobService_edcd183b773c9f62 = []byte{
	// 621 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x4e, 0xdb, 0x4c,
	0x10, 0x55, 0x62, 0x08, 0xf1, 0xf0, 0x7d, 0x80, 0x56, 0x15, 0x58, 0x4b, 0x25, 0x52, 0xa4, 0x4a,
	0xe9, 0x8f, 0x6c, 0x29, 0xb4, 0xa2, 0x45, 0xbd, 0x29, 0x2a, 0xaa, 0x82, 0x40, 0x42, 0x86, 0xde,
	0xf4, 0xa6, 0xb2, 0x9d, 0xc1, 0xdd, 0x10, 0x7b, 0x5d, 0xef, 0x1a, 0x35, 0xef, 0xd2, 0x37, 0xe8,
	0x13, 0xf5, 0x6d, 0xaa, 0xdd, 0xb5, 0x63, 0x63, 0xaa, 0x94, 0xde, 0xf9, 0xec, 0x9c, 0xd9, 0xb3,
	0x73, 0x3c, 0x33, 0xb0, 0x7b, 0x8d, 0x81, 0x90, 0x5e, 0xc4, 0x73, 0xf4, 0x4e, 0x79, 0x78, 0x89,
	0xf9, 0x2d, 0x8b, 0xd0, 0xcd, 0x72, 0x2e, 0x39, 0x01, 0x1d, 0x74, 0x55, 0x90, 0x3e, 0x36, 0x44,
	0x91, 0x61, 0x24, 0xbc, 0x71, 0x92, 0xf1, 0x5c, 0x5e, 0x66, 0x18, 0x19, 0x26, 0xdd, 0x8d, 0x39,
	0x8f, 0x67, 0xe8, 0x69, 0x14, 0x16, 0xd7, 0x1e, 0x26, 0x99, 0x9c, 0x97, 0xc1, 0xbd, 0x76, 0x50,
	0xb2, 0x04, 0x85, 0x0c, 0x92, 0xcc, 0x10, 0xf6, 0x7f, 0xf5, 0x60, 0xb3, 0x16, 0xbf, 0x9a, 0x67,
	0x28, 0x28, 0xc2, 0xf6, 0x65, 0x11, 0x26, 0x4c, 0x1a, 0xad, 0x53, 0x1e, 0xfa, 0xf8, 0xad, 0x40,
	0x21, 0xc9, 0x21, 0x00, 0x5b, 0xe8, 0x3b, 0x9d, 0x41, 0x67, 0xb8, 0x3e, 0xda, 0x71, 0xcd, 0x53,
	0xf5, 0xf3, 0xdc, 0xfa, 0x79, 0x7e, 0x83, 0x4a, 0x08, 0xac, 0xa4, 0x41, 0x82, 0x4e, 0x77, 0xd0,
	0x19, 0xda, 0xbe, 0xfe, 0xa6, 0x1e, 0xec, 0xdc, 0x93, 0x11, 0x19, 0x4f, 0x05, 0x92, 0x47, 0xb0,
	0x3a, 0xe5, 0xe1, 0x78, 0xa2, 0x25, 0x6c, 0xdf, 0x00, 0x7a, 0x0e, 0x5b, 0x67, 0x4c, 0x28, 0xa2,
	0x58, 0x30, 0xdf, 0xc2, 0xca, 0x94, 0x87, 0xc2, 0xe9, 0x0c, 0xac, 0xe1, 0xfa, 0xe8, 0xa9, 0x5b,
	0xdb, 0xe6, 0xb6, 0xca, 0x52, 0xf8, 0x03, 0xca, 0x80, 0xcd, 0x7c, 0x9d, 0x42, 0xf7, 0xe0, 0xff,
	0x8f, 0xd8, 0xac, 0x6e, 0x03, 0xba, 0xac, 0x92, 0xec, 0xb2, 0x09, 0x1d, 0xc3, 0x46, 0x45, 0x28,
	0xd5, 0x0e, 0xc1, 0x9a, 0xf2, 0xb0, 0x2c, 0xfc, 0x81, 0x62, 0x2a, 0x83, 0x3e, 0x81, 0xcd, 0xf7,
	0xe1, 0x5d, 0x2f, 0xdb, 0x6a, 0xfb, 0xb0, 0x55, 0x53, 0x4a, 0xbd, 0x36, 0xe7, 0xa7, 0x05, 0xf6,
	0xe2, 0xe6, 0x76, 0x54, 0xb9, 0x86, 0xdf, 0xe5, 0x78, 0x52, 0xba, 0x6c, 0x80, 0xb2, 0x5e, 0xce,
	0x33, 0x74, 0x2c, 0x63, 0xbd, 0xfa, 0x26, 0xdb, 0xd0, 0xcb, 0x8b, 0x34, 0xc5, 0xdc, 0x59, 0xd1,
	0xa7, 0x25, 0x52, 0xe7, 0x42, 0x06, 0xb2, 0x10, 0xce, 0xaa, 0x39, 0x37, 0x88, 0x50, 0xe8, 0x63,
	0x2a, 0x99, 0x64, 0x28, 0x9c, 0xde, 0xc0, 0x1a, 0xda, 0xfe, 0x02, 0xab, 0xd8, 0x35, 0x06, 0xb2,
	0xc8, 0x51, 0x38, 0x6b, 0x26, 0x56, 0x61, 0x72, 0x06, 0x6b, 0x09, 0xca, 0x9c, 0x45, 0xc2, 0xe9,
	0xeb, 0x1f, 0x34, 0x7a, 0x90, 0x67, 0xee, 0xb9, 0x49, 0x3a, 0x49, 0x65, 0x3e, 0xf7, 0xab, 0x2b,
	0xc8, 0x3b, 0x58, 0x9f, 0x05, 0x42, 0x7e, 0xca, 0x26, 0x81, 0xc4, 0x89, 0x63, 0xeb, 0xbf, 0x40,
	0x5d, 0xd3, 0xe2, 0x6e, 0xd5, 0xe2, 0xee, 0x55, 0xd5, 0xe2, 0x7e, 0x93, 0x4e, 0x5e, 0xc1, 0x5a,
	0x94, 0xa3, 0xce, 0x84, 0xbf, 0x66, 0x56, 0x54, 0x7a, 0x04, 0xff, 0x35, 0x1f, 0x43, 0xb6, 0xc0,
	0xba, 0xc1, 0x79, 0x69, 0xba, 0xfa, 0x54, 0xae, 0xdf, 0x06, 0xb3, 0xc2, 0xf4, 0x76, 0xc7, 0x37,
	0xe0, 0xa8, 0xfb, 0xa6, 0x33, 0xfa, 0x61, 0x01, 0xd4, 0x35, 0x12, 0x09, 0xb6, 0xe9, 0xf7, 0x53,
	0x1e, 0x92, 0xa5, 0x46, 0xfc, 0x79, 0xfa, 0xe8, 0xc1, 0x3f, 0xe5, 0x94, 0x2d, 0x74, 0x01, 0xfd,
	0x6a, 0x68, 0xc8, 0xf6, 0xbd, 0x8a, 0x4f, 0xd4, 0xae, 0xa0, 0x2f, 0x97, 0x5d, 0x7c, 0x6f, 0xe4,
	0xbe, 0x40, 0xcf, 0x8c, 0x05, 0x79, 0xb6, 0x2c, 0xef, 0xce, 0x6c, 0xd1, 0xe7, 0x0f, 0xa1, 0x96,
	0x02, 0x08, 0xfd, 0x6a, 0x12, 0xc8, 0x8b, 0x65, 0x79, 0xad, 0x91, 0x5a, 0x5e, 0x47, 0x7b, 0xb8,
	0x8e, 0xaf, 0xa0, 0xb1, 0x64, 0x8f, 0x1b, 0x5b, 0xf0, 0x42, 0xb9, 0xf3, 0xf9, 0x75, 0xcc, 0xe4,
	0xd7, 0x22, 0x74, 0x23, 0x9e, 0x78, 0x31, 0x9f, 0xe2, 0x8d, 0x67, 0xd6, 0xb0, 0xf6, 0x4e, 0x78,
	0x31, 0xa6, 0x98, 0xab, 0x16, 0xf1, 0x62, 0xee, 0xd5, 0x9b, 0x3c, 0xec, 0xe9, 0xf8, 0xc1, 0xef,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xc7, 0xf4, 0xb6, 0x7c, 0xde, 0x05, 0x00, 0x00,
}
