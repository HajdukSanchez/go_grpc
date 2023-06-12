// Use Proto 3 as syntax

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.23.2
// source: proto/testpb/test.proto

// Define the package name

package testpb

import (
	studentpb "github.com/hajduksanchez/go_grpc/proto/studentpb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Create the first structure
type Test struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Test) Reset() {
	*x = Test{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_testpb_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test) ProtoMessage() {}

func (x *Test) ProtoReflect() protoreflect.Message {
	mi := &file_proto_testpb_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test.ProtoReflect.Descriptor instead.
func (*Test) Descriptor() ([]byte, []int) {
	return file_proto_testpb_test_proto_rawDescGZIP(), []int{0}
}

func (x *Test) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Test) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Question struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Answer   string `protobuf:"bytes,2,opt,name=answer,proto3" json:"answer,omitempty"`
	Question string `protobuf:"bytes,3,opt,name=question,proto3" json:"question,omitempty"`
	TestId   string `protobuf:"bytes,4,opt,name=test_id,json=testId,proto3" json:"test_id,omitempty"`
}

func (x *Question) Reset() {
	*x = Question{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_testpb_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Question) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Question) ProtoMessage() {}

func (x *Question) ProtoReflect() protoreflect.Message {
	mi := &file_proto_testpb_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Question.ProtoReflect.Descriptor instead.
func (*Question) Descriptor() ([]byte, []int) {
	return file_proto_testpb_test_proto_rawDescGZIP(), []int{1}
}

func (x *Question) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Question) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

func (x *Question) GetQuestion() string {
	if x != nil {
		return x.Question
	}
	return ""
}

func (x *Question) GetTestId() string {
	if x != nil {
		return x.TestId
	}
	return ""
}

// Response when some question is setted
type SetQuestionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *SetQuestionResponse) Reset() {
	*x = SetQuestionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_testpb_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetQuestionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetQuestionResponse) ProtoMessage() {}

func (x *SetQuestionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_testpb_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetQuestionResponse.ProtoReflect.Descriptor instead.
func (*SetQuestionResponse) Descriptor() ([]byte, []int) {
	return file_proto_testpb_test_proto_rawDescGZIP(), []int{2}
}

func (x *SetQuestionResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

// Request to get a Test
type GetTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTestRequest) Reset() {
	*x = GetTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_testpb_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTestRequest) ProtoMessage() {}

func (x *GetTestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_testpb_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTestRequest.ProtoReflect.Descriptor instead.
func (*GetTestRequest) Descriptor() ([]byte, []int) {
	return file_proto_testpb_test_proto_rawDescGZIP(), []int{3}
}

func (x *GetTestRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Response when some test is setted
type SetTestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SetTestResponse) Reset() {
	*x = SetTestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_testpb_test_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetTestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTestResponse) ProtoMessage() {}

func (x *SetTestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_testpb_test_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTestResponse.ProtoReflect.Descriptor instead.
func (*SetTestResponse) Descriptor() ([]byte, []int) {
	return file_proto_testpb_test_proto_rawDescGZIP(), []int{4}
}

func (x *SetTestResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SetTestResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Enroll a student in a test
type EnrollStudentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId string `protobuf:"bytes,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	TestId    string `protobuf:"bytes,2,opt,name=test_id,json=testId,proto3" json:"test_id,omitempty"`
}

func (x *EnrollStudentRequest) Reset() {
	*x = EnrollStudentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_testpb_test_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnrollStudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnrollStudentRequest) ProtoMessage() {}

func (x *EnrollStudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_testpb_test_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnrollStudentRequest.ProtoReflect.Descriptor instead.
func (*EnrollStudentRequest) Descriptor() ([]byte, []int) {
	return file_proto_testpb_test_proto_rawDescGZIP(), []int{5}
}

func (x *EnrollStudentRequest) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *EnrollStudentRequest) GetTestId() string {
	if x != nil {
		return x.TestId
	}
	return ""
}

// Response when some student is enrolled
type SetEnrollStudentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *SetEnrollStudentResponse) Reset() {
	*x = SetEnrollStudentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_testpb_test_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetEnrollStudentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetEnrollStudentResponse) ProtoMessage() {}

func (x *SetEnrollStudentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_testpb_test_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetEnrollStudentResponse.ProtoReflect.Descriptor instead.
func (*SetEnrollStudentResponse) Descriptor() ([]byte, []int) {
	return file_proto_testpb_test_proto_rawDescGZIP(), []int{6}
}

func (x *SetEnrollStudentResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

// Get all student for a specific test
type GetStudentsPerTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestId string `protobuf:"bytes,1,opt,name=test_id,json=testId,proto3" json:"test_id,omitempty"`
}

func (x *GetStudentsPerTestRequest) Reset() {
	*x = GetStudentsPerTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_testpb_test_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStudentsPerTestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStudentsPerTestRequest) ProtoMessage() {}

func (x *GetStudentsPerTestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_testpb_test_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStudentsPerTestRequest.ProtoReflect.Descriptor instead.
func (*GetStudentsPerTestRequest) Descriptor() ([]byte, []int) {
	return file_proto_testpb_test_proto_rawDescGZIP(), []int{7}
}

func (x *GetStudentsPerTestRequest) GetTestId() string {
	if x != nil {
		return x.TestId
	}
	return ""
}

var File_proto_testpb_test_proto protoreflect.FileDescriptor

var file_proto_testpb_test_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x70, 0x62,
	0x2f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a,
	0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x67, 0x0a, 0x08, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65,
	0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x22, 0x25, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x35, 0x0a, 0x0f,
	0x53, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x4e, 0x0a, 0x14, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x53, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65,
	0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x22, 0x2a, 0x0a, 0x18, 0x53, 0x65, 0x74, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22,
	0x34, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x65,
	0x72, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x65, 0x73, 0x74, 0x49, 0x64, 0x32, 0xbf, 0x02, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65,
	0x73, 0x74, 0x12, 0x2c, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x12, 0x0a, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x53, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3b, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x0e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x19, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x4d, 0x0a,
	0x0d, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x1a,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x53, 0x65, 0x74, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x53, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x49, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x65, 0x72, 0x54, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x65, 0x72, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x30, 0x01, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x6a, 0x64, 0x75, 0x6b, 0x73, 0x61, 0x6e, 0x63,
	0x68, 0x65, 0x7a, 0x2f, 0x67, 0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_testpb_test_proto_rawDescOnce sync.Once
	file_proto_testpb_test_proto_rawDescData = file_proto_testpb_test_proto_rawDesc
)

func file_proto_testpb_test_proto_rawDescGZIP() []byte {
	file_proto_testpb_test_proto_rawDescOnce.Do(func() {
		file_proto_testpb_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_testpb_test_proto_rawDescData)
	})
	return file_proto_testpb_test_proto_rawDescData
}

var file_proto_testpb_test_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_testpb_test_proto_goTypes = []interface{}{
	(*Test)(nil),                      // 0: test.Test
	(*Question)(nil),                  // 1: test.Question
	(*SetQuestionResponse)(nil),       // 2: test.SetQuestionResponse
	(*GetTestRequest)(nil),            // 3: test.GetTestRequest
	(*SetTestResponse)(nil),           // 4: test.SetTestResponse
	(*EnrollStudentRequest)(nil),      // 5: test.EnrollStudentRequest
	(*SetEnrollStudentResponse)(nil),  // 6: test.SetEnrollStudentResponse
	(*GetStudentsPerTestRequest)(nil), // 7: test.GetStudentsPerTestRequest
	(*studentpb.Student)(nil),         // 8: student.Student
}
var file_proto_testpb_test_proto_depIdxs = []int32{
	3, // 0: test.TestService.GetTest:input_type -> test.GetTestRequest
	0, // 1: test.TestService.SetTest:input_type -> test.Test
	1, // 2: test.TestService.SetQuestions:input_type -> test.Question
	5, // 3: test.TestService.EnrollStudent:input_type -> test.EnrollStudentRequest
	7, // 4: test.TestService.GetStudentsPerTest:input_type -> test.GetStudentsPerTestRequest
	0, // 5: test.TestService.GetTest:output_type -> test.Test
	4, // 6: test.TestService.SetTest:output_type -> test.SetTestResponse
	2, // 7: test.TestService.SetQuestions:output_type -> test.SetQuestionResponse
	6, // 8: test.TestService.EnrollStudent:output_type -> test.SetEnrollStudentResponse
	8, // 9: test.TestService.GetStudentsPerTest:output_type -> student.Student
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_testpb_test_proto_init() }
func file_proto_testpb_test_proto_init() {
	if File_proto_testpb_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_testpb_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_testpb_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Question); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_testpb_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetQuestionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_testpb_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTestRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_testpb_test_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetTestResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_testpb_test_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnrollStudentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_testpb_test_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetEnrollStudentResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_testpb_test_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStudentsPerTestRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_testpb_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_testpb_test_proto_goTypes,
		DependencyIndexes: file_proto_testpb_test_proto_depIdxs,
		MessageInfos:      file_proto_testpb_test_proto_msgTypes,
	}.Build()
	File_proto_testpb_test_proto = out.File
	file_proto_testpb_test_proto_rawDesc = nil
	file_proto_testpb_test_proto_goTypes = nil
	file_proto_testpb_test_proto_depIdxs = nil
}
