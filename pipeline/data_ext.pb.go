// Code generated by protoc-gen-go. DO NOT EDIT.
// source: data_ext.proto

package pipeline

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type FeatureData struct {
	FromFeature []*Feature `protobuf:"bytes,1,rep,name=from_feature,json=fromFeature" json:"from_feature,omitempty"`
	DatasetUri  string     `protobuf:"bytes,2,opt,name=dataset_uri,json=datasetUri" json:"dataset_uri,omitempty"`
}

func (m *FeatureData) Reset()                    { *m = FeatureData{} }
func (m *FeatureData) String() string            { return proto.CompactTextString(m) }
func (*FeatureData) ProtoMessage()               {}
func (*FeatureData) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *FeatureData) GetFromFeature() []*Feature {
	if m != nil {
		return m.FromFeature
	}
	return nil
}

func (m *FeatureData) GetDatasetUri() string {
	if m != nil {
		return m.DatasetUri
	}
	return ""
}

type AddFeaturesRequest struct {
	Context  *SessionContext `protobuf:"bytes,1,opt,name=context" json:"context,omitempty"`
	Features []*FeatureData  `protobuf:"bytes,2,rep,name=features" json:"features,omitempty"`
}

func (m *AddFeaturesRequest) Reset()                    { *m = AddFeaturesRequest{} }
func (m *AddFeaturesRequest) String() string            { return proto.CompactTextString(m) }
func (*AddFeaturesRequest) ProtoMessage()               {}
func (*AddFeaturesRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *AddFeaturesRequest) GetContext() *SessionContext {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *AddFeaturesRequest) GetFeatures() []*FeatureData {
	if m != nil {
		return m.Features
	}
	return nil
}

type RemoveFeaturesRequest struct {
	Context  *SessionContext `protobuf:"bytes,1,opt,name=context" json:"context,omitempty"`
	Features []*Feature      `protobuf:"bytes,2,rep,name=features" json:"features,omitempty"`
}

func (m *RemoveFeaturesRequest) Reset()                    { *m = RemoveFeaturesRequest{} }
func (m *RemoveFeaturesRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveFeaturesRequest) ProtoMessage()               {}
func (*RemoveFeaturesRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *RemoveFeaturesRequest) GetContext() *SessionContext {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *RemoveFeaturesRequest) GetFeatures() []*Feature {
	if m != nil {
		return m.Features
	}
	return nil
}

type Sample struct {
	SampleId string `protobuf:"bytes,1,opt,name=sample_id,json=sampleId" json:"sample_id,omitempty"`
	DataUri  string `protobuf:"bytes,2,opt,name=data_uri,json=dataUri" json:"data_uri,omitempty"`
}

func (m *Sample) Reset()                    { *m = Sample{} }
func (m *Sample) String() string            { return proto.CompactTextString(m) }
func (*Sample) ProtoMessage()               {}
func (*Sample) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *Sample) GetSampleId() string {
	if m != nil {
		return m.SampleId
	}
	return ""
}

func (m *Sample) GetDataUri() string {
	if m != nil {
		return m.DataUri
	}
	return ""
}

type SampleData struct {
	FromSample []*Sample `protobuf:"bytes,1,rep,name=from_sample,json=fromSample" json:"from_sample,omitempty"`
	DatasetUri string    `protobuf:"bytes,2,opt,name=dataset_uri,json=datasetUri" json:"dataset_uri,omitempty"`
}

func (m *SampleData) Reset()                    { *m = SampleData{} }
func (m *SampleData) String() string            { return proto.CompactTextString(m) }
func (*SampleData) ProtoMessage()               {}
func (*SampleData) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *SampleData) GetFromSample() []*Sample {
	if m != nil {
		return m.FromSample
	}
	return nil
}

func (m *SampleData) GetDatasetUri() string {
	if m != nil {
		return m.DatasetUri
	}
	return ""
}

type AddSamplesRequest struct {
	Context *SessionContext `protobuf:"bytes,1,opt,name=context" json:"context,omitempty"`
	Samples []*SampleData   `protobuf:"bytes,2,rep,name=samples" json:"samples,omitempty"`
}

func (m *AddSamplesRequest) Reset()                    { *m = AddSamplesRequest{} }
func (m *AddSamplesRequest) String() string            { return proto.CompactTextString(m) }
func (*AddSamplesRequest) ProtoMessage()               {}
func (*AddSamplesRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *AddSamplesRequest) GetContext() *SessionContext {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *AddSamplesRequest) GetSamples() []*SampleData {
	if m != nil {
		return m.Samples
	}
	return nil
}

type RemoveSamplesRequest struct {
	Context   *SessionContext `protobuf:"bytes,1,opt,name=context" json:"context,omitempty"`
	SampleIds []*Sample       `protobuf:"bytes,2,rep,name=sample_ids,json=sampleIds" json:"sample_ids,omitempty"`
}

func (m *RemoveSamplesRequest) Reset()                    { *m = RemoveSamplesRequest{} }
func (m *RemoveSamplesRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveSamplesRequest) ProtoMessage()               {}
func (*RemoveSamplesRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *RemoveSamplesRequest) GetContext() *SessionContext {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *RemoveSamplesRequest) GetSampleIds() []*Sample {
	if m != nil {
		return m.SampleIds
	}
	return nil
}

type ReplacementData struct {
	Value      string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	FeatureId  string `protobuf:"bytes,2,opt,name=feature_id,json=featureId" json:"feature_id,omitempty"`
	SampleId   string `protobuf:"bytes,3,opt,name=sample_id,json=sampleId" json:"sample_id,omitempty"`
	DatasetUri string `protobuf:"bytes,4,opt,name=dataset_uri,json=datasetUri" json:"dataset_uri,omitempty"`
}

func (m *ReplacementData) Reset()                    { *m = ReplacementData{} }
func (m *ReplacementData) String() string            { return proto.CompactTextString(m) }
func (*ReplacementData) ProtoMessage()               {}
func (*ReplacementData) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *ReplacementData) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *ReplacementData) GetFeatureId() string {
	if m != nil {
		return m.FeatureId
	}
	return ""
}

func (m *ReplacementData) GetSampleId() string {
	if m != nil {
		return m.SampleId
	}
	return ""
}

func (m *ReplacementData) GetDatasetUri() string {
	if m != nil {
		return m.DatasetUri
	}
	return ""
}

type ReplaceDataRequest struct {
	Context      *SessionContext    `protobuf:"bytes,1,opt,name=context" json:"context,omitempty"`
	Replacements []*ReplacementData `protobuf:"bytes,2,rep,name=replacements" json:"replacements,omitempty"`
}

func (m *ReplaceDataRequest) Reset()                    { *m = ReplaceDataRequest{} }
func (m *ReplaceDataRequest) String() string            { return proto.CompactTextString(m) }
func (*ReplaceDataRequest) ProtoMessage()               {}
func (*ReplaceDataRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *ReplaceDataRequest) GetContext() *SessionContext {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *ReplaceDataRequest) GetReplacements() []*ReplacementData {
	if m != nil {
		return m.Replacements
	}
	return nil
}

type MaterializeRequest struct {
	Context          *SessionContext `protobuf:"bytes,1,opt,name=context" json:"context,omitempty"`
	SourceDatasetUri string          `protobuf:"bytes,2,opt,name=source_dataset_uri,json=sourceDatasetUri" json:"source_dataset_uri,omitempty"`
	DestDatasetUri   string          `protobuf:"bytes,3,opt,name=dest_dataset_uri,json=destDatasetUri" json:"dest_dataset_uri,omitempty"`
}

func (m *MaterializeRequest) Reset()                    { *m = MaterializeRequest{} }
func (m *MaterializeRequest) String() string            { return proto.CompactTextString(m) }
func (*MaterializeRequest) ProtoMessage()               {}
func (*MaterializeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *MaterializeRequest) GetContext() *SessionContext {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *MaterializeRequest) GetSourceDatasetUri() string {
	if m != nil {
		return m.SourceDatasetUri
	}
	return ""
}

func (m *MaterializeRequest) GetDestDatasetUri() string {
	if m != nil {
		return m.DestDatasetUri
	}
	return ""
}

type TrainValidationSplitRequest struct {
	Context      *SessionContext `protobuf:"bytes,1,opt,name=context" json:"context,omitempty"`
	ValSize      float32         `protobuf:"fixed32,2,opt,name=val_size,json=valSize" json:"val_size,omitempty"`
	Seed         int32           `protobuf:"varint,3,opt,name=seed" json:"seed,omitempty"`
	IsValidation bool            `protobuf:"varint,4,opt,name=is_validation,json=isValidation" json:"is_validation,omitempty"`
	DatasetUri   string          `protobuf:"bytes,5,opt,name=dataset_uri,json=datasetUri" json:"dataset_uri,omitempty"`
}

func (m *TrainValidationSplitRequest) Reset()                    { *m = TrainValidationSplitRequest{} }
func (m *TrainValidationSplitRequest) String() string            { return proto.CompactTextString(m) }
func (*TrainValidationSplitRequest) ProtoMessage()               {}
func (*TrainValidationSplitRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func (m *TrainValidationSplitRequest) GetContext() *SessionContext {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *TrainValidationSplitRequest) GetValSize() float32 {
	if m != nil {
		return m.ValSize
	}
	return 0
}

func (m *TrainValidationSplitRequest) GetSeed() int32 {
	if m != nil {
		return m.Seed
	}
	return 0
}

func (m *TrainValidationSplitRequest) GetIsValidation() bool {
	if m != nil {
		return m.IsValidation
	}
	return false
}

func (m *TrainValidationSplitRequest) GetDatasetUri() string {
	if m != nil {
		return m.DatasetUri
	}
	return ""
}

type RevertRequest struct {
	DatasetUri string `protobuf:"bytes,1,opt,name=dataset_uri,json=datasetUri" json:"dataset_uri,omitempty"`
}

func (m *RevertRequest) Reset()                    { *m = RevertRequest{} }
func (m *RevertRequest) String() string            { return proto.CompactTextString(m) }
func (*RevertRequest) ProtoMessage()               {}
func (*RevertRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

func (m *RevertRequest) GetDatasetUri() string {
	if m != nil {
		return m.DatasetUri
	}
	return ""
}

func init() {
	proto.RegisterType((*FeatureData)(nil), "FeatureData")
	proto.RegisterType((*AddFeaturesRequest)(nil), "AddFeaturesRequest")
	proto.RegisterType((*RemoveFeaturesRequest)(nil), "RemoveFeaturesRequest")
	proto.RegisterType((*Sample)(nil), "Sample")
	proto.RegisterType((*SampleData)(nil), "SampleData")
	proto.RegisterType((*AddSamplesRequest)(nil), "AddSamplesRequest")
	proto.RegisterType((*RemoveSamplesRequest)(nil), "RemoveSamplesRequest")
	proto.RegisterType((*ReplacementData)(nil), "ReplacementData")
	proto.RegisterType((*ReplaceDataRequest)(nil), "ReplaceDataRequest")
	proto.RegisterType((*MaterializeRequest)(nil), "MaterializeRequest")
	proto.RegisterType((*TrainValidationSplitRequest)(nil), "TrainValidationSplitRequest")
	proto.RegisterType((*RevertRequest)(nil), "RevertRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DataExt service

type DataExtClient interface {
	// Add and remove features to/from datasets
	AddFeatures(ctx context.Context, in *AddFeaturesRequest, opts ...grpc.CallOption) (*Response, error)
	RemoveFeatures(ctx context.Context, in *RemoveFeaturesRequest, opts ...grpc.CallOption) (*Response, error)
	// Add and remove records (rows) to/from datasets
	AddSamples(ctx context.Context, in *AddSamplesRequest, opts ...grpc.CallOption) (*Response, error)
	RemoveSamples(ctx context.Context, in *RemoveSamplesRequest, opts ...grpc.CallOption) (*Response, error)
	// Replace individual data points in a set
	ReplaceData(ctx context.Context, in *ReplaceDataRequest, opts ...grpc.CallOption) (*Response, error)
	// Persist the dataset with modifications applied for future use
	Materialize(ctx context.Context, in *MaterializeRequest, opts ...grpc.CallOption) (*Response, error)
	// Deterministic split of a dataset into training a validation
	// Filters out all but validation records or training records depending on is_validation
	TrainValidationSplit(ctx context.Context, in *TrainValidationSplitRequest, opts ...grpc.CallOption) (*Response, error)
	// Revert the dataset to the original state
	Revert(ctx context.Context, in *RevertRequest, opts ...grpc.CallOption) (*Response, error)
}

type dataExtClient struct {
	cc *grpc.ClientConn
}

func NewDataExtClient(cc *grpc.ClientConn) DataExtClient {
	return &dataExtClient{cc}
}

func (c *dataExtClient) AddFeatures(ctx context.Context, in *AddFeaturesRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/DataExt/AddFeatures", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataExtClient) RemoveFeatures(ctx context.Context, in *RemoveFeaturesRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/DataExt/RemoveFeatures", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataExtClient) AddSamples(ctx context.Context, in *AddSamplesRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/DataExt/AddSamples", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataExtClient) RemoveSamples(ctx context.Context, in *RemoveSamplesRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/DataExt/RemoveSamples", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataExtClient) ReplaceData(ctx context.Context, in *ReplaceDataRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/DataExt/ReplaceData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataExtClient) Materialize(ctx context.Context, in *MaterializeRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/DataExt/Materialize", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataExtClient) TrainValidationSplit(ctx context.Context, in *TrainValidationSplitRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/DataExt/TrainValidationSplit", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataExtClient) Revert(ctx context.Context, in *RevertRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/DataExt/Revert", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DataExt service

type DataExtServer interface {
	// Add and remove features to/from datasets
	AddFeatures(context.Context, *AddFeaturesRequest) (*Response, error)
	RemoveFeatures(context.Context, *RemoveFeaturesRequest) (*Response, error)
	// Add and remove records (rows) to/from datasets
	AddSamples(context.Context, *AddSamplesRequest) (*Response, error)
	RemoveSamples(context.Context, *RemoveSamplesRequest) (*Response, error)
	// Replace individual data points in a set
	ReplaceData(context.Context, *ReplaceDataRequest) (*Response, error)
	// Persist the dataset with modifications applied for future use
	Materialize(context.Context, *MaterializeRequest) (*Response, error)
	// Deterministic split of a dataset into training a validation
	// Filters out all but validation records or training records depending on is_validation
	TrainValidationSplit(context.Context, *TrainValidationSplitRequest) (*Response, error)
	// Revert the dataset to the original state
	Revert(context.Context, *RevertRequest) (*Response, error)
}

func RegisterDataExtServer(s *grpc.Server, srv DataExtServer) {
	s.RegisterService(&_DataExt_serviceDesc, srv)
}

func _DataExt_AddFeatures_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFeaturesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataExtServer).AddFeatures(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataExt/AddFeatures",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataExtServer).AddFeatures(ctx, req.(*AddFeaturesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataExt_RemoveFeatures_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFeaturesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataExtServer).RemoveFeatures(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataExt/RemoveFeatures",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataExtServer).RemoveFeatures(ctx, req.(*RemoveFeaturesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataExt_AddSamples_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSamplesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataExtServer).AddSamples(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataExt/AddSamples",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataExtServer).AddSamples(ctx, req.(*AddSamplesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataExt_RemoveSamples_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveSamplesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataExtServer).RemoveSamples(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataExt/RemoveSamples",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataExtServer).RemoveSamples(ctx, req.(*RemoveSamplesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataExt_ReplaceData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplaceDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataExtServer).ReplaceData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataExt/ReplaceData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataExtServer).ReplaceData(ctx, req.(*ReplaceDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataExt_Materialize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MaterializeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataExtServer).Materialize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataExt/Materialize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataExtServer).Materialize(ctx, req.(*MaterializeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataExt_TrainValidationSplit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrainValidationSplitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataExtServer).TrainValidationSplit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataExt/TrainValidationSplit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataExtServer).TrainValidationSplit(ctx, req.(*TrainValidationSplitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataExt_Revert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataExtServer).Revert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataExt/Revert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataExtServer).Revert(ctx, req.(*RevertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataExt_serviceDesc = grpc.ServiceDesc{
	ServiceName: "DataExt",
	HandlerType: (*DataExtServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddFeatures",
			Handler:    _DataExt_AddFeatures_Handler,
		},
		{
			MethodName: "RemoveFeatures",
			Handler:    _DataExt_RemoveFeatures_Handler,
		},
		{
			MethodName: "AddSamples",
			Handler:    _DataExt_AddSamples_Handler,
		},
		{
			MethodName: "RemoveSamples",
			Handler:    _DataExt_RemoveSamples_Handler,
		},
		{
			MethodName: "ReplaceData",
			Handler:    _DataExt_ReplaceData_Handler,
		},
		{
			MethodName: "Materialize",
			Handler:    _DataExt_Materialize_Handler,
		},
		{
			MethodName: "TrainValidationSplit",
			Handler:    _DataExt_TrainValidationSplit_Handler,
		},
		{
			MethodName: "Revert",
			Handler:    _DataExt_Revert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "data_ext.proto",
}

func init() { proto.RegisterFile("data_ext.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 645 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xfd, 0x8b, 0x33, 0x4e, 0xd3, 0x32, 0x6d, 0x51, 0xda, 0x82, 0xa8, 0x0c, 0x45, 0x41,
	0xc0, 0x16, 0xb5, 0x70, 0xa7, 0x50, 0x90, 0x7a, 0xe0, 0xb2, 0xe1, 0x47, 0x82, 0x83, 0xb5, 0xc4,
	0x53, 0xb1, 0x92, 0x63, 0x1b, 0xef, 0x26, 0xaa, 0x7a, 0xe6, 0x1d, 0x38, 0xf3, 0x20, 0xbc, 0x1b,
	0xf2, 0xae, 0x13, 0xc7, 0x76, 0x05, 0x22, 0xdc, 0xbc, 0xf3, 0xb3, 0xf3, 0xcd, 0xf7, 0xcd, 0x8e,
	0xa1, 0x1b, 0x0a, 0x2d, 0x02, 0xba, 0xd4, 0x2c, 0xcd, 0x12, 0x9d, 0xec, 0xc1, 0x30, 0xc9, 0xc8,
	0x7e, 0xfb, 0x9f, 0xc1, 0x7b, 0x43, 0x42, 0x8f, 0x33, 0x3a, 0x13, 0x5a, 0xe0, 0x23, 0xe8, 0x5c,
	0x64, 0xc9, 0x28, 0xb8, 0xb0, 0xb6, 0x9e, 0x73, 0xb0, 0xdc, 0xf7, 0x8e, 0x5d, 0x56, 0xc4, 0x70,
	0x2f, 0xf7, 0x16, 0x07, 0xbc, 0x0b, 0x5e, 0x7e, 0xb3, 0x22, 0x1d, 0x8c, 0x33, 0xd9, 0x5b, 0x3a,
	0x70, 0xfa, 0x6d, 0x0e, 0x85, 0xe9, 0x7d, 0x26, 0x7d, 0x09, 0x78, 0x1a, 0x86, 0x45, 0xb8, 0xe2,
	0xf4, 0x6d, 0x4c, 0x4a, 0xe3, 0x43, 0x68, 0x0d, 0x93, 0x58, 0xd3, 0xa5, 0xee, 0x39, 0x07, 0x4e,
	0xdf, 0x3b, 0xde, 0x60, 0x03, 0x52, 0x4a, 0x26, 0xf1, 0x2b, 0x6b, 0xe6, 0x53, 0x3f, 0xf6, 0xc1,
	0x2d, 0x90, 0xa8, 0xde, 0x92, 0x81, 0xd2, 0x61, 0x73, 0x70, 0xf9, 0xcc, 0xeb, 0x7f, 0x85, 0x1d,
	0x4e, 0xa3, 0x64, 0x42, 0xff, 0x51, 0xed, 0x7e, 0xa3, 0x5a, 0xd9, 0x78, 0x59, 0xe9, 0x05, 0xac,
	0x0d, 0xc4, 0x28, 0x8d, 0x08, 0xf7, 0xa1, 0xad, 0xcc, 0x57, 0x20, 0x43, 0x73, 0x79, 0x9b, 0xbb,
	0xd6, 0x70, 0x1e, 0xe2, 0x2e, 0xb8, 0x86, 0xf6, 0x92, 0x99, 0x56, 0x7e, 0xce, 0x69, 0xf9, 0x08,
	0x60, 0x6f, 0x30, 0x94, 0xf7, 0xc1, 0x90, 0x1a, 0xd8, 0xcc, 0x82, 0xf1, 0x16, 0xb3, 0x11, 0x1c,
	0x72, 0x5f, 0x51, 0xef, 0xaf, 0x7c, 0x13, 0xdc, 0x3c, 0x0d, 0x43, 0x1b, 0xbd, 0x08, 0x01, 0x87,
	0xd0, 0xb2, 0x28, 0xa6, 0xfd, 0x7b, 0xac, 0x04, 0xca, 0xa7, 0x3e, 0x5f, 0xc2, 0xb6, 0xe5, 0x7a,
	0xf1, 0x4a, 0x0f, 0x00, 0x66, 0xd4, 0x4d, 0x8b, 0xcd, 0x7a, 0x6e, 0x4f, 0x49, 0x54, 0xfe, 0x77,
	0x07, 0x36, 0x38, 0xa5, 0x91, 0x18, 0xd2, 0x88, 0x62, 0x6d, 0x08, 0xdb, 0x86, 0xd5, 0x89, 0x88,
	0xc6, 0x54, 0x50, 0x6e, 0x0f, 0x78, 0x07, 0xa0, 0x90, 0x28, 0x57, 0xc3, 0x72, 0xd3, 0x2e, 0x2c,
	0xe7, 0x61, 0x55, 0xab, 0xe5, 0x9a, 0x56, 0x35, 0x62, 0x57, 0x1a, 0xc4, 0x8e, 0x01, 0x0b, 0x14,
	0x86, 0x89, 0x7f, 0xef, 0xf7, 0x19, 0x74, 0xb2, 0xb2, 0x8d, 0x69, 0xc7, 0x9b, 0xac, 0xd6, 0x1b,
	0xaf, 0x44, 0xf9, 0x3f, 0x1c, 0xc0, 0xb7, 0x42, 0x53, 0x26, 0x45, 0x24, 0xaf, 0x68, 0x81, 0xba,
	0x8f, 0x01, 0x55, 0x32, 0xce, 0x86, 0x14, 0x34, 0x27, 0x67, 0xd3, 0x7a, 0xce, 0x66, 0x6d, 0x62,
	0x1f, 0x36, 0x43, 0x52, 0xba, 0x12, 0x6b, 0xb9, 0xea, 0xe6, 0xf6, 0x32, 0xd2, 0xff, 0xe5, 0xc0,
	0xfe, 0xbb, 0x4c, 0xc8, 0xf8, 0x83, 0x88, 0x64, 0x28, 0xb4, 0x4c, 0xe2, 0x41, 0x1a, 0x49, 0xbd,
	0x00, 0xc4, 0x5d, 0x70, 0x27, 0x22, 0x0a, 0x94, 0xbc, 0x22, 0x03, 0x6c, 0x89, 0xb7, 0x26, 0x22,
	0x1a, 0xc8, 0x2b, 0x42, 0x84, 0x15, 0x45, 0x64, 0xf5, 0x5a, 0xe5, 0xe6, 0x1b, 0xef, 0xc1, 0xba,
	0x54, 0xc1, 0x64, 0x56, 0xd6, 0xa8, 0xe5, 0xf2, 0x8e, 0x54, 0x25, 0x94, 0xba, 0xa0, 0xab, 0x0d,
	0x41, 0x9f, 0xc2, 0x3a, 0xa7, 0x09, 0x65, 0x33, 0xc0, 0xb5, 0x0c, 0xa7, 0x9e, 0x71, 0xfc, 0x73,
	0x19, 0x5a, 0x39, 0x01, 0xaf, 0x2f, 0x35, 0x1e, 0x81, 0x37, 0xb7, 0xd7, 0x70, 0x8b, 0x35, 0xb7,
	0xdc, 0x5e, 0x9b, 0x71, 0x52, 0x69, 0x12, 0x2b, 0xf2, 0x6f, 0xe0, 0x73, 0xe8, 0x56, 0xb7, 0x13,
	0xde, 0x62, 0xd7, 0xae, 0xab, 0x6a, 0xda, 0x13, 0x80, 0xf2, 0x3d, 0x23, 0xb2, 0xc6, 0xe3, 0xae,
	0x86, 0x9f, 0xe4, 0x4d, 0xcd, 0xbd, 0x4b, 0xdc, 0x61, 0xd7, 0xbd, 0xd3, 0x6a, 0xd2, 0x11, 0x78,
	0x73, 0xa3, 0x8d, 0x5b, 0xac, 0x39, 0xe8, 0x8d, 0x84, 0xb9, 0x99, 0xc4, 0x2d, 0xd6, 0x9c, 0xd0,
	0x6a, 0xc2, 0x29, 0x6c, 0x5f, 0x37, 0x2a, 0x78, 0x9b, 0xfd, 0x61, 0x82, 0xaa, 0x57, 0x1c, 0xc2,
	0x9a, 0x95, 0x0b, 0xbb, 0xac, 0xa2, 0x5b, 0x25, 0xec, 0x25, 0x7c, 0x72, 0x53, 0x99, 0x52, 0x24,
	0x63, 0xfa, 0xb2, 0x66, 0xfe, 0x6f, 0x27, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x57, 0x0b,
	0x09, 0xfd, 0x06, 0x00, 0x00,
}