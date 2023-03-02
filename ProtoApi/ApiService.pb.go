// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.10
// source: ApiService.proto

package ProtoApi

import (
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

type Wrap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Wrap) Reset() {
	*x = Wrap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ApiService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Wrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Wrap) ProtoMessage() {}

func (x *Wrap) ProtoReflect() protoreflect.Message {
	mi := &file_ApiService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Wrap.ProtoReflect.Descriptor instead.
func (*Wrap) Descriptor() ([]byte, []int) {
	return file_ApiService_proto_rawDescGZIP(), []int{0}
}

type SheduleArray struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject    int32     `protobuf:"varint,1,opt,name=Subject,proto3" json:"Subject,omitempty"`
	Tutor      int32     `protobuf:"varint,2,opt,name=Tutor,proto3" json:"Tutor,omitempty"`
	Auditorium int32     `protobuf:"varint,3,opt,name=Auditorium,proto3" json:"Auditorium,omitempty"`
	Type       int32     `protobuf:"varint,4,opt,name=Type,proto3" json:"Type,omitempty"`
	Group      int32     `protobuf:"varint,5,opt,name=Group,proto3" json:"Group,omitempty"`
	Dates      *DateTime `protobuf:"bytes,6,opt,name=Dates,proto3" json:"Dates,omitempty"`
	Number     int32     `protobuf:"varint,7,opt,name=Number,proto3" json:"Number,omitempty"`
}

func (x *SheduleArray) Reset() {
	*x = SheduleArray{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ApiService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SheduleArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SheduleArray) ProtoMessage() {}

func (x *SheduleArray) ProtoReflect() protoreflect.Message {
	mi := &file_ApiService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SheduleArray.ProtoReflect.Descriptor instead.
func (*SheduleArray) Descriptor() ([]byte, []int) {
	return file_ApiService_proto_rawDescGZIP(), []int{1}
}

func (x *SheduleArray) GetSubject() int32 {
	if x != nil {
		return x.Subject
	}
	return 0
}

func (x *SheduleArray) GetTutor() int32 {
	if x != nil {
		return x.Tutor
	}
	return 0
}

func (x *SheduleArray) GetAuditorium() int32 {
	if x != nil {
		return x.Auditorium
	}
	return 0
}

func (x *SheduleArray) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *SheduleArray) GetGroup() int32 {
	if x != nil {
		return x.Group
	}
	return 0
}

func (x *SheduleArray) GetDates() *DateTime {
	if x != nil {
		return x.Dates
	}
	return nil
}

func (x *SheduleArray) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type DateTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year    int32 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	Month   int32 `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	Day     int32 `protobuf:"varint,3,opt,name=day,proto3" json:"day,omitempty"`
	Hours   int32 `protobuf:"varint,4,opt,name=hours,proto3" json:"hours,omitempty"`
	Minutes int32 `protobuf:"varint,5,opt,name=minutes,proto3" json:"minutes,omitempty"`
	Seconds int32 `protobuf:"varint,6,opt,name=seconds,proto3" json:"seconds,omitempty"`
}

func (x *DateTime) Reset() {
	*x = DateTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ApiService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DateTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DateTime) ProtoMessage() {}

func (x *DateTime) ProtoReflect() protoreflect.Message {
	mi := &file_ApiService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DateTime.ProtoReflect.Descriptor instead.
func (*DateTime) Descriptor() ([]byte, []int) {
	return file_ApiService_proto_rawDescGZIP(), []int{2}
}

func (x *DateTime) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *DateTime) GetMonth() int32 {
	if x != nil {
		return x.Month
	}
	return 0
}

func (x *DateTime) GetDay() int32 {
	if x != nil {
		return x.Day
	}
	return 0
}

func (x *DateTime) GetHours() int32 {
	if x != nil {
		return x.Hours
	}
	return 0
}

func (x *DateTime) GetMinutes() int32 {
	if x != nil {
		return x.Minutes
	}
	return 0
}

func (x *DateTime) GetSeconds() int32 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter string `protobuf:"bytes,1,opt,name=Filter,proto3" json:"Filter,omitempty"`
	Value  int32  `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ApiService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_ApiService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_ApiService_proto_rawDescGZIP(), []int{3}
}

func (x *Filter) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *Filter) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type SubjectsFromDb struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subjects []*StructSubject `protobuf:"bytes,1,rep,name=Subjects,proto3" json:"Subjects,omitempty"`
}

func (x *SubjectsFromDb) Reset() {
	*x = SubjectsFromDb{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ApiService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubjectsFromDb) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubjectsFromDb) ProtoMessage() {}

func (x *SubjectsFromDb) ProtoReflect() protoreflect.Message {
	mi := &file_ApiService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubjectsFromDb.ProtoReflect.Descriptor instead.
func (*SubjectsFromDb) Descriptor() ([]byte, []int) {
	return file_ApiService_proto_rawDescGZIP(), []int{4}
}

func (x *SubjectsFromDb) GetSubjects() []*StructSubject {
	if x != nil {
		return x.Subjects
	}
	return nil
}

type StructSubject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	SubjectItem string `protobuf:"bytes,2,opt,name=SubjectItem,proto3" json:"SubjectItem,omitempty"`
}

func (x *StructSubject) Reset() {
	*x = StructSubject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ApiService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StructSubject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StructSubject) ProtoMessage() {}

func (x *StructSubject) ProtoReflect() protoreflect.Message {
	mi := &file_ApiService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StructSubject.ProtoReflect.Descriptor instead.
func (*StructSubject) Descriptor() ([]byte, []int) {
	return file_ApiService_proto_rawDescGZIP(), []int{5}
}

func (x *StructSubject) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *StructSubject) GetSubjectItem() string {
	if x != nil {
		return x.SubjectItem
	}
	return ""
}

type TutorsFromDb struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tutors []*StructTutor `protobuf:"bytes,1,rep,name=Tutors,proto3" json:"Tutors,omitempty"`
}

func (x *TutorsFromDb) Reset() {
	*x = TutorsFromDb{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ApiService_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TutorsFromDb) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TutorsFromDb) ProtoMessage() {}

func (x *TutorsFromDb) ProtoReflect() protoreflect.Message {
	mi := &file_ApiService_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TutorsFromDb.ProtoReflect.Descriptor instead.
func (*TutorsFromDb) Descriptor() ([]byte, []int) {
	return file_ApiService_proto_rawDescGZIP(), []int{6}
}

func (x *TutorsFromDb) GetTutors() []*StructTutor {
	if x != nil {
		return x.Tutors
	}
	return nil
}

type StructTutor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	SecondName string `protobuf:"bytes,2,opt,name=SecondName,proto3" json:"SecondName,omitempty"`
	FirstName  string `protobuf:"bytes,3,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName   string `protobuf:"bytes,4,opt,name=LastName,proto3" json:"LastName,omitempty"`
}

func (x *StructTutor) Reset() {
	*x = StructTutor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ApiService_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StructTutor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StructTutor) ProtoMessage() {}

func (x *StructTutor) ProtoReflect() protoreflect.Message {
	mi := &file_ApiService_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StructTutor.ProtoReflect.Descriptor instead.
func (*StructTutor) Descriptor() ([]byte, []int) {
	return file_ApiService_proto_rawDescGZIP(), []int{7}
}

func (x *StructTutor) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *StructTutor) GetSecondName() string {
	if x != nil {
		return x.SecondName
	}
	return ""
}

func (x *StructTutor) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *StructTutor) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

type AuditoriumsFromDb struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auditoriums []*StructClassroom `protobuf:"bytes,1,rep,name=Auditoriums,proto3" json:"Auditoriums,omitempty"`
}

func (x *AuditoriumsFromDb) Reset() {
	*x = AuditoriumsFromDb{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ApiService_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuditoriumsFromDb) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditoriumsFromDb) ProtoMessage() {}

func (x *AuditoriumsFromDb) ProtoReflect() protoreflect.Message {
	mi := &file_ApiService_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditoriumsFromDb.ProtoReflect.Descriptor instead.
func (*AuditoriumsFromDb) Descriptor() ([]byte, []int) {
	return file_ApiService_proto_rawDescGZIP(), []int{8}
}

func (x *AuditoriumsFromDb) GetAuditoriums() []*StructClassroom {
	if x != nil {
		return x.Auditoriums
	}
	return nil
}

type StructClassroom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Classroom string `protobuf:"bytes,2,opt,name=Classroom,proto3" json:"Classroom,omitempty"`
}

func (x *StructClassroom) Reset() {
	*x = StructClassroom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ApiService_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StructClassroom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StructClassroom) ProtoMessage() {}

func (x *StructClassroom) ProtoReflect() protoreflect.Message {
	mi := &file_ApiService_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StructClassroom.ProtoReflect.Descriptor instead.
func (*StructClassroom) Descriptor() ([]byte, []int) {
	return file_ApiService_proto_rawDescGZIP(), []int{9}
}

func (x *StructClassroom) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *StructClassroom) GetClassroom() string {
	if x != nil {
		return x.Classroom
	}
	return ""
}

type TypesFromDb struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Types []*StructType `protobuf:"bytes,1,rep,name=Types,proto3" json:"Types,omitempty"`
}

func (x *TypesFromDb) Reset() {
	*x = TypesFromDb{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ApiService_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypesFromDb) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypesFromDb) ProtoMessage() {}

func (x *TypesFromDb) ProtoReflect() protoreflect.Message {
	mi := &file_ApiService_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypesFromDb.ProtoReflect.Descriptor instead.
func (*TypesFromDb) Descriptor() ([]byte, []int) {
	return file_ApiService_proto_rawDescGZIP(), []int{10}
}

func (x *TypesFromDb) GetTypes() []*StructType {
	if x != nil {
		return x.Types
	}
	return nil
}

type StructType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *StructType) Reset() {
	*x = StructType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ApiService_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StructType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StructType) ProtoMessage() {}

func (x *StructType) ProtoReflect() protoreflect.Message {
	mi := &file_ApiService_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StructType.ProtoReflect.Descriptor instead.
func (*StructType) Descriptor() ([]byte, []int) {
	return file_ApiService_proto_rawDescGZIP(), []int{11}
}

func (x *StructType) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *StructType) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type GroupsFromDb struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Groups []*StructGroup `protobuf:"bytes,1,rep,name=Groups,proto3" json:"Groups,omitempty"`
}

func (x *GroupsFromDb) Reset() {
	*x = GroupsFromDb{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ApiService_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupsFromDb) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupsFromDb) ProtoMessage() {}

func (x *GroupsFromDb) ProtoReflect() protoreflect.Message {
	mi := &file_ApiService_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupsFromDb.ProtoReflect.Descriptor instead.
func (*GroupsFromDb) Descriptor() ([]byte, []int) {
	return file_ApiService_proto_rawDescGZIP(), []int{12}
}

func (x *GroupsFromDb) GetGroups() []*StructGroup {
	if x != nil {
		return x.Groups
	}
	return nil
}

type StructGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	GroupName string `protobuf:"bytes,2,opt,name=GroupName,proto3" json:"GroupName,omitempty"`
}

func (x *StructGroup) Reset() {
	*x = StructGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ApiService_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StructGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StructGroup) ProtoMessage() {}

func (x *StructGroup) ProtoReflect() protoreflect.Message {
	mi := &file_ApiService_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StructGroup.ProtoReflect.Descriptor instead.
func (*StructGroup) Descriptor() ([]byte, []int) {
	return file_ApiService_proto_rawDescGZIP(), []int{13}
}

func (x *StructGroup) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *StructGroup) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

var File_ApiService_proto protoreflect.FileDescriptor

var file_ApiService_proto_rawDesc = []byte{
	0x0a, 0x10, 0x41, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x41, 0x70, 0x69, 0x22, 0x06, 0x0a, 0x04,
	0x57, 0x72, 0x61, 0x70, 0x22, 0xca, 0x01, 0x0a, 0x0c, 0x53, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x54, 0x75, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x54, 0x75, 0x74, 0x6f, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x41, 0x75, 0x64, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x41, 0x75, 0x64, 0x69, 0x74,
	0x6f, 0x72, 0x69, 0x75, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x28, 0x0a, 0x05, 0x44, 0x61, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x41, 0x70, 0x69, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x52, 0x05, 0x44, 0x61, 0x74, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x22, 0x90, 0x01, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65,
	0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x6f,
	0x75, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x68, 0x6f, 0x75, 0x72, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x22, 0x36, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x45, 0x0a, 0x0e,
	0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x62, 0x12, 0x33,
	0x0a, 0x08, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x41, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x08, 0x53, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x22, 0x41, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x53, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x3d, 0x0a, 0x0c, 0x54, 0x75, 0x74, 0x6f, 0x72, 0x73,
	0x46, 0x72, 0x6f, 0x6d, 0x44, 0x62, 0x12, 0x2d, 0x0a, 0x06, 0x54, 0x75, 0x74, 0x6f, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x41, 0x70,
	0x69, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x54, 0x75, 0x74, 0x6f, 0x72, 0x52, 0x06, 0x54,
	0x75, 0x74, 0x6f, 0x72, 0x73, 0x22, 0x77, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x54,
	0x75, 0x74, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x50,
	0x0a, 0x11, 0x41, 0x75, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x75, 0x6d, 0x73, 0x46, 0x72, 0x6f,
	0x6d, 0x44, 0x62, 0x12, 0x3b, 0x0a, 0x0b, 0x41, 0x75, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x75,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x41, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x72,
	0x6f, 0x6f, 0x6d, 0x52, 0x0b, 0x41, 0x75, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x75, 0x6d, 0x73,
	0x22, 0x3f, 0x0a, 0x0f, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x72,
	0x6f, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f,
	0x6d, 0x22, 0x39, 0x0a, 0x0b, 0x54, 0x79, 0x70, 0x65, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x62,
	0x12, 0x2a, 0x0a, 0x05, 0x54, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x41, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x54, 0x79, 0x70, 0x65, 0x73, 0x22, 0x30, 0x0a, 0x0a,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22, 0x3d,
	0x0a, 0x0c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x62, 0x12, 0x2d,
	0x0a, 0x06, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x41, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x06, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x3b, 0x0a,
	0x0b, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x8c, 0x03, 0x0a, 0x0e, 0x53,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x44,
	0x62, 0x52, 0x50, 0x43, 0x12, 0x0e, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x41, 0x70, 0x69, 0x2e,
	0x57, 0x72, 0x61, 0x70, 0x1a, 0x18, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x41, 0x70, 0x69, 0x2e,
	0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x62, 0x12, 0x3c,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x75, 0x74, 0x6f, 0x72, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x44,
	0x62, 0x52, 0x50, 0x43, 0x12, 0x0e, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x41, 0x70, 0x69, 0x2e,
	0x57, 0x72, 0x61, 0x70, 0x1a, 0x16, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x41, 0x70, 0x69, 0x2e,
	0x54, 0x75, 0x74, 0x6f, 0x72, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x62, 0x12, 0x45, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x75, 0x6d, 0x46, 0x72, 0x6f,
	0x6d, 0x44, 0x62, 0x52, 0x50, 0x43, 0x12, 0x0e, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x41, 0x70,
	0x69, 0x2e, 0x57, 0x72, 0x61, 0x70, 0x1a, 0x1b, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x41, 0x70,
	0x69, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x75, 0x6d, 0x73, 0x46, 0x72, 0x6f,
	0x6d, 0x44, 0x62, 0x12, 0x39, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x46, 0x72,
	0x6f, 0x6d, 0x44, 0x62, 0x52, 0x50, 0x43, 0x12, 0x0e, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x41,
	0x70, 0x69, 0x2e, 0x57, 0x72, 0x61, 0x70, 0x1a, 0x15, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x41,
	0x70, 0x69, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x62, 0x12, 0x3b,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x62,
	0x52, 0x50, 0x43, 0x12, 0x0e, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x41, 0x70, 0x69, 0x2e, 0x57,
	0x72, 0x61, 0x70, 0x1a, 0x16, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x41, 0x70, 0x69, 0x2e, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x62, 0x12, 0x3c, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x53, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x62, 0x12,
	0x10, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x41, 0x70, 0x69, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x1a, 0x16, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x41, 0x70, 0x69, 0x2e, 0x53, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x41, 0x72, 0x72, 0x61, 0x79, 0x42, 0x15, 0x5a, 0x13, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x68, 0x65, 0x65, 0x74, 0x2f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x41, 0x70, 0x69, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ApiService_proto_rawDescOnce sync.Once
	file_ApiService_proto_rawDescData = file_ApiService_proto_rawDesc
)

func file_ApiService_proto_rawDescGZIP() []byte {
	file_ApiService_proto_rawDescOnce.Do(func() {
		file_ApiService_proto_rawDescData = protoimpl.X.CompressGZIP(file_ApiService_proto_rawDescData)
	})
	return file_ApiService_proto_rawDescData
}

var file_ApiService_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_ApiService_proto_goTypes = []interface{}{
	(*Wrap)(nil),              // 0: ProtoApi.Wrap
	(*SheduleArray)(nil),      // 1: ProtoApi.SheduleArray
	(*DateTime)(nil),          // 2: ProtoApi.DateTime
	(*Filter)(nil),            // 3: ProtoApi.Filter
	(*SubjectsFromDb)(nil),    // 4: ProtoApi.SubjectsFromDb
	(*StructSubject)(nil),     // 5: ProtoApi.StructSubject
	(*TutorsFromDb)(nil),      // 6: ProtoApi.TutorsFromDb
	(*StructTutor)(nil),       // 7: ProtoApi.StructTutor
	(*AuditoriumsFromDb)(nil), // 8: ProtoApi.AuditoriumsFromDb
	(*StructClassroom)(nil),   // 9: ProtoApi.StructClassroom
	(*TypesFromDb)(nil),       // 10: ProtoApi.TypesFromDb
	(*StructType)(nil),        // 11: ProtoApi.StructType
	(*GroupsFromDb)(nil),      // 12: ProtoApi.GroupsFromDb
	(*StructGroup)(nil),       // 13: ProtoApi.StructGroup
}
var file_ApiService_proto_depIdxs = []int32{
	2,  // 0: ProtoApi.SheduleArray.Dates:type_name -> ProtoApi.DateTime
	5,  // 1: ProtoApi.SubjectsFromDb.Subjects:type_name -> ProtoApi.StructSubject
	7,  // 2: ProtoApi.TutorsFromDb.Tutors:type_name -> ProtoApi.StructTutor
	9,  // 3: ProtoApi.AuditoriumsFromDb.Auditoriums:type_name -> ProtoApi.StructClassroom
	11, // 4: ProtoApi.TypesFromDb.Types:type_name -> ProtoApi.StructType
	13, // 5: ProtoApi.GroupsFromDb.Groups:type_name -> ProtoApi.StructGroup
	0,  // 6: ProtoApi.SheduleService.GetSubjectFromDbRPC:input_type -> ProtoApi.Wrap
	0,  // 7: ProtoApi.SheduleService.GetTutorsFromDbRPC:input_type -> ProtoApi.Wrap
	0,  // 8: ProtoApi.SheduleService.GetAuditoriumFromDbRPC:input_type -> ProtoApi.Wrap
	0,  // 9: ProtoApi.SheduleService.GetTypeFromDbRPC:input_type -> ProtoApi.Wrap
	0,  // 10: ProtoApi.SheduleService.GetGroupFromDbRPC:input_type -> ProtoApi.Wrap
	3,  // 11: ProtoApi.SheduleService.GetSheduleFromDb:input_type -> ProtoApi.Filter
	4,  // 12: ProtoApi.SheduleService.GetSubjectFromDbRPC:output_type -> ProtoApi.SubjectsFromDb
	6,  // 13: ProtoApi.SheduleService.GetTutorsFromDbRPC:output_type -> ProtoApi.TutorsFromDb
	8,  // 14: ProtoApi.SheduleService.GetAuditoriumFromDbRPC:output_type -> ProtoApi.AuditoriumsFromDb
	10, // 15: ProtoApi.SheduleService.GetTypeFromDbRPC:output_type -> ProtoApi.TypesFromDb
	12, // 16: ProtoApi.SheduleService.GetGroupFromDbRPC:output_type -> ProtoApi.GroupsFromDb
	1,  // 17: ProtoApi.SheduleService.GetSheduleFromDb:output_type -> ProtoApi.SheduleArray
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_ApiService_proto_init() }
func file_ApiService_proto_init() {
	if File_ApiService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ApiService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Wrap); i {
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
		file_ApiService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SheduleArray); i {
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
		file_ApiService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DateTime); i {
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
		file_ApiService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
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
		file_ApiService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubjectsFromDb); i {
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
		file_ApiService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StructSubject); i {
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
		file_ApiService_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TutorsFromDb); i {
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
		file_ApiService_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StructTutor); i {
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
		file_ApiService_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuditoriumsFromDb); i {
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
		file_ApiService_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StructClassroom); i {
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
		file_ApiService_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TypesFromDb); i {
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
		file_ApiService_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StructType); i {
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
		file_ApiService_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupsFromDb); i {
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
		file_ApiService_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StructGroup); i {
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
			RawDescriptor: file_ApiService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ApiService_proto_goTypes,
		DependencyIndexes: file_ApiService_proto_depIdxs,
		MessageInfos:      file_ApiService_proto_msgTypes,
	}.Build()
	File_ApiService_proto = out.File
	file_ApiService_proto_rawDesc = nil
	file_ApiService_proto_goTypes = nil
	file_ApiService_proto_depIdxs = nil
}
