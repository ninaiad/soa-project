// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.2
// source: posts.proto

package posts_proto

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

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorId int64  `protobuf:"varint,1,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Text     string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_posts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_posts_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *CreateRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
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
		mi := &file_posts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostId) ProtoMessage() {}

func (x *PostId) ProtoReflect() protoreflect.Message {
	mi := &file_posts_proto_msgTypes[1]
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
	return file_posts_proto_rawDescGZIP(), []int{1}
}

func (x *PostId) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
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
		mi := &file_posts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserId) ProtoMessage() {}

func (x *UserId) ProtoReflect() protoreflect.Message {
	mi := &file_posts_proto_msgTypes[2]
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
	return file_posts_proto_rawDescGZIP(), []int{2}
}

func (x *UserId) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorId int64  `protobuf:"varint,1,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	PostId   int64  `protobuf:"varint,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	Text     string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_posts_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_posts_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateRequest) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *UpdateRequest) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *UpdateRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type AuthoredPostId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorId int64 `protobuf:"varint,1,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	PostId   int64 `protobuf:"varint,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
}

func (x *AuthoredPostId) Reset() {
	*x = AuthoredPostId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthoredPostId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthoredPostId) ProtoMessage() {}

func (x *AuthoredPostId) ProtoReflect() protoreflect.Message {
	mi := &file_posts_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthoredPostId.ProtoReflect.Descriptor instead.
func (*AuthoredPostId) Descriptor() ([]byte, []int) {
	return file_posts_proto_rawDescGZIP(), []int{4}
}

func (x *AuthoredPostId) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *AuthoredPostId) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Text        string                 `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	TimeUpdated *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=time_updated,json=timeUpdated,proto3" json:"time_updated,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_posts_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_posts_proto_rawDescGZIP(), []int{5}
}

func (x *Post) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Post) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Post) GetTimeUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.TimeUpdated
	}
	return nil
}

type PageOfPostsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorId int64 `protobuf:"varint,1,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	PageNum  int32 `protobuf:"varint,2,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *PageOfPostsRequest) Reset() {
	*x = PageOfPostsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageOfPostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageOfPostsRequest) ProtoMessage() {}

func (x *PageOfPostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_posts_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageOfPostsRequest.ProtoReflect.Descriptor instead.
func (*PageOfPostsRequest) Descriptor() ([]byte, []int) {
	return file_posts_proto_rawDescGZIP(), []int{6}
}

func (x *PageOfPostsRequest) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *PageOfPostsRequest) GetPageNum() int32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *PageOfPostsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type PageOfPosts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNum  int32   `protobuf:"varint,1,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	PageSize int32   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	AuthorId int64   `protobuf:"varint,3,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Posts    []*Post `protobuf:"bytes,4,rep,name=posts,proto3" json:"posts,omitempty"`
}

func (x *PageOfPosts) Reset() {
	*x = PageOfPosts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageOfPosts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageOfPosts) ProtoMessage() {}

func (x *PageOfPosts) ProtoReflect() protoreflect.Message {
	mi := &file_posts_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageOfPosts.ProtoReflect.Descriptor instead.
func (*PageOfPosts) Descriptor() ([]byte, []int) {
	return file_posts_proto_rawDescGZIP(), []int{7}
}

func (x *PageOfPosts) GetPageNum() int32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *PageOfPosts) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PageOfPosts) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *PageOfPosts) GetPosts() []*Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

var File_posts_proto protoreflect.FileDescriptor

var file_posts_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70,
	0x6f, 0x73, 0x74, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x18, 0x0a, 0x06, 0x50, 0x6f,
	0x73, 0x74, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x18, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x59,
	0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70,
	0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x46, 0x0a, 0x0e, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x65, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49,
	0x64, 0x22, 0x69, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x3d, 0x0a,
	0x0c, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0b, 0x74, 0x69, 0x6d, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x69, 0x0a, 0x12,
	0x50, 0x61, 0x67, 0x65, 0x4f, 0x66, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x8b, 0x01, 0x0a, 0x0b, 0x50, 0x61, 0x67, 0x65,
	0x4f, 0x66, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e,
	0x75, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x05,
	0x70, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x05,
	0x70, 0x6f, 0x73, 0x74, 0x73, 0x32, 0xa0, 0x03, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f,
	0x73, 0x74, 0x49, 0x64, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0a, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x65, 0x64, 0x50,
	0x6f, 0x73, 0x74, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x3b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x70, 0x6f, 0x73,
	0x74, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x65,
	0x64, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x1a, 0x11, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x4f, 0x66, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x1f,
	0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x67,
	0x65, 0x4f, 0x66, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x4f, 0x66, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x70, 0x6f, 0x73, 0x74,
	0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x73, 0x6f, 0x61, 0x2d,
	0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_posts_proto_rawDescOnce sync.Once
	file_posts_proto_rawDescData = file_posts_proto_rawDesc
)

func file_posts_proto_rawDescGZIP() []byte {
	file_posts_proto_rawDescOnce.Do(func() {
		file_posts_proto_rawDescData = protoimpl.X.CompressGZIP(file_posts_proto_rawDescData)
	})
	return file_posts_proto_rawDescData
}

var file_posts_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_posts_proto_goTypes = []interface{}{
	(*CreateRequest)(nil),         // 0: posts_proto.CreateRequest
	(*PostId)(nil),                // 1: posts_proto.PostId
	(*UserId)(nil),                // 2: posts_proto.UserId
	(*UpdateRequest)(nil),         // 3: posts_proto.UpdateRequest
	(*AuthoredPostId)(nil),        // 4: posts_proto.AuthoredPostId
	(*Post)(nil),                  // 5: posts_proto.Post
	(*PageOfPostsRequest)(nil),    // 6: posts_proto.PageOfPostsRequest
	(*PageOfPosts)(nil),           // 7: posts_proto.PageOfPosts
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 9: google.protobuf.Empty
}
var file_posts_proto_depIdxs = []int32{
	8, // 0: posts_proto.Post.time_updated:type_name -> google.protobuf.Timestamp
	5, // 1: posts_proto.PageOfPosts.posts:type_name -> posts_proto.Post
	0, // 2: posts_proto.PostsServer.CreatePost:input_type -> posts_proto.CreateRequest
	3, // 3: posts_proto.PostsServer.UpdatePost:input_type -> posts_proto.UpdateRequest
	4, // 4: posts_proto.PostsServer.DeletePost:input_type -> posts_proto.AuthoredPostId
	4, // 5: posts_proto.PostsServer.GetPost:input_type -> posts_proto.AuthoredPostId
	6, // 6: posts_proto.PostsServer.GetPageOfPosts:input_type -> posts_proto.PageOfPostsRequest
	2, // 7: posts_proto.PostsServer.DeleteUser:input_type -> posts_proto.UserId
	1, // 8: posts_proto.PostsServer.CreatePost:output_type -> posts_proto.PostId
	9, // 9: posts_proto.PostsServer.UpdatePost:output_type -> google.protobuf.Empty
	9, // 10: posts_proto.PostsServer.DeletePost:output_type -> google.protobuf.Empty
	5, // 11: posts_proto.PostsServer.GetPost:output_type -> posts_proto.Post
	7, // 12: posts_proto.PostsServer.GetPageOfPosts:output_type -> posts_proto.PageOfPosts
	9, // 13: posts_proto.PostsServer.DeleteUser:output_type -> google.protobuf.Empty
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_posts_proto_init() }
func file_posts_proto_init() {
	if File_posts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_posts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_posts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_posts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_posts_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_posts_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthoredPostId); i {
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
		file_posts_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post); i {
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
		file_posts_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageOfPostsRequest); i {
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
		file_posts_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageOfPosts); i {
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
			RawDescriptor: file_posts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_posts_proto_goTypes,
		DependencyIndexes: file_posts_proto_depIdxs,
		MessageInfos:      file_posts_proto_msgTypes,
	}.Build()
	File_posts_proto = out.File
	file_posts_proto_rawDesc = nil
	file_posts_proto_goTypes = nil
	file_posts_proto_depIdxs = nil
}
