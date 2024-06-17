// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.2
// source: statistics.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EventType int32

const (
	EventType_LIKE EventType = 0
	EventType_VIEW EventType = 1
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0: "LIKE",
		1: "VIEW",
	}
	EventType_value = map[string]int32{
		"LIKE": 0,
		"VIEW": 1,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_statistics_proto_enumTypes[0].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_statistics_proto_enumTypes[0]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_statistics_proto_rawDescGZIP(), []int{0}
}

type UserId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserId) Reset() {
	*x = UserId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserId) ProtoMessage() {}

func (x *UserId) ProtoReflect() protoreflect.Message {
	mi := &file_statistics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserId.ProtoReflect.Descriptor instead.
func (*UserId) Descriptor() ([]byte, []int) {
	return file_statistics_proto_rawDescGZIP(), []int{0}
}

func (x *UserId) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PostId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PostId) Reset() {
	*x = PostId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostId) ProtoMessage() {}

func (x *PostId) ProtoReflect() protoreflect.Message {
	mi := &file_statistics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostId.ProtoReflect.Descriptor instead.
func (*PostId) Descriptor() ([]byte, []int) {
	return file_statistics_proto_rawDescGZIP(), []int{1}
}

func (x *PostId) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PostStatistics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId   int64  `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	AuthorId int64  `protobuf:"varint,4,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	NumLikes uint64 `protobuf:"varint,2,opt,name=num_likes,json=numLikes,proto3" json:"num_likes,omitempty"`
	NumViews uint64 `protobuf:"varint,3,opt,name=num_views,json=numViews,proto3" json:"num_views,omitempty"`
}

func (x *PostStatistics) Reset() {
	*x = PostStatistics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostStatistics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostStatistics) ProtoMessage() {}

func (x *PostStatistics) ProtoReflect() protoreflect.Message {
	mi := &file_statistics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostStatistics.ProtoReflect.Descriptor instead.
func (*PostStatistics) Descriptor() ([]byte, []int) {
	return file_statistics_proto_rawDescGZIP(), []int{2}
}

func (x *PostStatistics) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *PostStatistics) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *PostStatistics) GetNumLikes() uint64 {
	if x != nil {
		return x.NumLikes
	}
	return 0
}

func (x *PostStatistics) GetNumViews() uint64 {
	if x != nil {
		return x.NumViews
	}
	return 0
}

type UserStatistics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NumLikes uint64 `protobuf:"varint,2,opt,name=num_likes,json=numLikes,proto3" json:"num_likes,omitempty"`
	NumViews uint64 `protobuf:"varint,3,opt,name=num_views,json=numViews,proto3" json:"num_views,omitempty"`
}

func (x *UserStatistics) Reset() {
	*x = UserStatistics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserStatistics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserStatistics) ProtoMessage() {}

func (x *UserStatistics) ProtoReflect() protoreflect.Message {
	mi := &file_statistics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserStatistics.ProtoReflect.Descriptor instead.
func (*UserStatistics) Descriptor() ([]byte, []int) {
	return file_statistics_proto_rawDescGZIP(), []int{3}
}

func (x *UserStatistics) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserStatistics) GetNumLikes() uint64 {
	if x != nil {
		return x.NumLikes
	}
	return 0
}

func (x *UserStatistics) GetNumViews() uint64 {
	if x != nil {
		return x.NumViews
	}
	return 0
}

type TopKRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event EventType `protobuf:"varint,1,opt,name=event,proto3,enum=statistics_proto.EventType" json:"event,omitempty"`
	K     uint64    `protobuf:"varint,2,opt,name=k,proto3" json:"k,omitempty"`
}

func (x *TopKRequest) Reset() {
	*x = TopKRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistics_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopKRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopKRequest) ProtoMessage() {}

func (x *TopKRequest) ProtoReflect() protoreflect.Message {
	mi := &file_statistics_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopKRequest.ProtoReflect.Descriptor instead.
func (*TopKRequest) Descriptor() ([]byte, []int) {
	return file_statistics_proto_rawDescGZIP(), []int{4}
}

func (x *TopKRequest) GetEvent() EventType {
	if x != nil {
		return x.Event
	}
	return EventType_LIKE
}

func (x *TopKRequest) GetK() uint64 {
	if x != nil {
		return x.K
	}
	return 0
}

type TopPosts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts         []*PostStatistics      `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
	TimeCollected *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=time_collected,json=timeCollected,proto3" json:"time_collected,omitempty"`
}

func (x *TopPosts) Reset() {
	*x = TopPosts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistics_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopPosts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopPosts) ProtoMessage() {}

func (x *TopPosts) ProtoReflect() protoreflect.Message {
	mi := &file_statistics_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopPosts.ProtoReflect.Descriptor instead.
func (*TopPosts) Descriptor() ([]byte, []int) {
	return file_statistics_proto_rawDescGZIP(), []int{5}
}

func (x *TopPosts) GetPosts() []*PostStatistics {
	if x != nil {
		return x.Posts
	}
	return nil
}

func (x *TopPosts) GetTimeCollected() *timestamppb.Timestamp {
	if x != nil {
		return x.TimeCollected
	}
	return nil
}

type TopUsers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users         []*UserStatistics      `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	TimeCollected *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=time_collected,json=timeCollected,proto3" json:"time_collected,omitempty"`
}

func (x *TopUsers) Reset() {
	*x = TopUsers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistics_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopUsers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopUsers) ProtoMessage() {}

func (x *TopUsers) ProtoReflect() protoreflect.Message {
	mi := &file_statistics_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopUsers.ProtoReflect.Descriptor instead.
func (*TopUsers) Descriptor() ([]byte, []int) {
	return file_statistics_proto_rawDescGZIP(), []int{6}
}

func (x *TopUsers) GetUsers() []*UserStatistics {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *TopUsers) GetTimeCollected() *timestamppb.Timestamp {
	if x != nil {
		return x.TimeCollected
	}
	return nil
}

var File_statistics_proto protoreflect.FileDescriptor

var file_statistics_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x18, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x18, 0x0a, 0x06,
	0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x80, 0x01, 0x0a, 0x0e, 0x50, 0x6f, 0x73, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x6e, 0x75, 0x6d, 0x5f, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x08, 0x6e, 0x75, 0x6d, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09,
	0x6e, 0x75, 0x6d, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x6e, 0x75, 0x6d, 0x56, 0x69, 0x65, 0x77, 0x73, 0x22, 0x5a, 0x0a, 0x0e, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6e,
	0x75, 0x6d, 0x5f, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08,
	0x6e, 0x75, 0x6d, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x75, 0x6d, 0x5f,
	0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6e, 0x75, 0x6d,
	0x56, 0x69, 0x65, 0x77, 0x73, 0x22, 0x4e, 0x0a, 0x0b, 0x54, 0x6f, 0x70, 0x4b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x01, 0x6b, 0x22, 0x85, 0x01, 0x0a, 0x08, 0x54, 0x6f, 0x70, 0x50, 0x6f, 0x73,
	0x74, 0x73, 0x12, 0x36, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x41, 0x0a, 0x0e, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d,
	0x74, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x22, 0x85, 0x01,
	0x0a, 0x08, 0x54, 0x6f, 0x70, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x36, 0x0a, 0x05, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x05, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x12, 0x41, 0x0a, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x2a, 0x1f, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x49, 0x4b, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04,
	0x56, 0x49, 0x45, 0x57, 0x10, 0x01, 0x32, 0x84, 0x03, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0a,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x40,
	0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x51, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x18, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x1a,
	0x20, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x4b, 0x50, 0x6f,
	0x73, 0x74, 0x73, 0x12, 0x1d, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x70, 0x4b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x70, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x00,
	0x12, 0x4b, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x4b, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x12, 0x1d, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x70, 0x4b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x70, 0x55, 0x73, 0x65, 0x72, 0x73, 0x22, 0x00, 0x42, 0x16, 0x5a,
	0x14, 0x73, 0x6f, 0x61, 0x2d, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_statistics_proto_rawDescOnce sync.Once
	file_statistics_proto_rawDescData = file_statistics_proto_rawDesc
)

func file_statistics_proto_rawDescGZIP() []byte {
	file_statistics_proto_rawDescOnce.Do(func() {
		file_statistics_proto_rawDescData = protoimpl.X.CompressGZIP(file_statistics_proto_rawDescData)
	})
	return file_statistics_proto_rawDescData
}

var file_statistics_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_statistics_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_statistics_proto_goTypes = []interface{}{
	(EventType)(0),                // 0: statistics_proto.EventType
	(*UserId)(nil),                // 1: statistics_proto.UserId
	(*PostId)(nil),                // 2: statistics_proto.PostId
	(*PostStatistics)(nil),        // 3: statistics_proto.PostStatistics
	(*UserStatistics)(nil),        // 4: statistics_proto.UserStatistics
	(*TopKRequest)(nil),           // 5: statistics_proto.TopKRequest
	(*TopPosts)(nil),              // 6: statistics_proto.TopPosts
	(*TopUsers)(nil),              // 7: statistics_proto.TopUsers
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 9: google.protobuf.Empty
}
var file_statistics_proto_depIdxs = []int32{
	0,  // 0: statistics_proto.TopKRequest.event:type_name -> statistics_proto.EventType
	3,  // 1: statistics_proto.TopPosts.posts:type_name -> statistics_proto.PostStatistics
	8,  // 2: statistics_proto.TopPosts.time_collected:type_name -> google.protobuf.Timestamp
	4,  // 3: statistics_proto.TopUsers.users:type_name -> statistics_proto.UserStatistics
	8,  // 4: statistics_proto.TopUsers.time_collected:type_name -> google.protobuf.Timestamp
	1,  // 5: statistics_proto.StatisticsService.DeleteUser:input_type -> statistics_proto.UserId
	2,  // 6: statistics_proto.StatisticsService.DeletePost:input_type -> statistics_proto.PostId
	2,  // 7: statistics_proto.StatisticsService.GetPostStatistics:input_type -> statistics_proto.PostId
	5,  // 8: statistics_proto.StatisticsService.GetTopKPosts:input_type -> statistics_proto.TopKRequest
	5,  // 9: statistics_proto.StatisticsService.GetTopKUsers:input_type -> statistics_proto.TopKRequest
	9,  // 10: statistics_proto.StatisticsService.DeleteUser:output_type -> google.protobuf.Empty
	9,  // 11: statistics_proto.StatisticsService.DeletePost:output_type -> google.protobuf.Empty
	3,  // 12: statistics_proto.StatisticsService.GetPostStatistics:output_type -> statistics_proto.PostStatistics
	6,  // 13: statistics_proto.StatisticsService.GetTopKPosts:output_type -> statistics_proto.TopPosts
	7,  // 14: statistics_proto.StatisticsService.GetTopKUsers:output_type -> statistics_proto.TopUsers
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_statistics_proto_init() }
func file_statistics_proto_init() {
	if File_statistics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_statistics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserId); i {
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
		file_statistics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostId); i {
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
		file_statistics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostStatistics); i {
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
		file_statistics_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserStatistics); i {
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
		file_statistics_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopKRequest); i {
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
		file_statistics_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopPosts); i {
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
		file_statistics_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopUsers); i {
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
			RawDescriptor: file_statistics_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_statistics_proto_goTypes,
		DependencyIndexes: file_statistics_proto_depIdxs,
		EnumInfos:         file_statistics_proto_enumTypes,
		MessageInfos:      file_statistics_proto_msgTypes,
	}.Build()
	File_statistics_proto = out.File
	file_statistics_proto_rawDesc = nil
	file_statistics_proto_goTypes = nil
	file_statistics_proto_depIdxs = nil
}
