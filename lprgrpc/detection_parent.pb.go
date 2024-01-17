// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: detection_parent.proto

package lpr_grpc

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

type Detection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid        string    `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	SensorName string    `protobuf:"bytes,2,opt,name=sensor_name,json=sensorName,proto3" json:"sensor_name,omitempty"`
	XType      string    `protobuf:"bytes,3,opt,name=_type,json=Type,proto3" json:"_type,omitempty"`
	Ts         float64   `protobuf:"fixed64,4,opt,name=ts,proto3" json:"ts,omitempty"`
	Label      []string  `protobuf:"bytes,5,rep,name=label,proto3" json:"label,omitempty"`
	Confidence []float64 `protobuf:"fixed64,6,rep,packed,name=confidence,proto3" json:"confidence,omitempty"`
}

func (x *Detection) Reset() {
	*x = Detection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_detection_parent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Detection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Detection) ProtoMessage() {}

func (x *Detection) ProtoReflect() protoreflect.Message {
	mi := &file_detection_parent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Detection.ProtoReflect.Descriptor instead.
func (*Detection) Descriptor() ([]byte, []int) {
	return file_detection_parent_proto_rawDescGZIP(), []int{0}
}

func (x *Detection) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Detection) GetSensorName() string {
	if x != nil {
		return x.SensorName
	}
	return ""
}

func (x *Detection) GetXType() string {
	if x != nil {
		return x.XType
	}
	return ""
}

func (x *Detection) GetTs() float64 {
	if x != nil {
		return x.Ts
	}
	return 0
}

func (x *Detection) GetLabel() []string {
	if x != nil {
		return x.Label
	}
	return nil
}

func (x *Detection) GetConfidence() []float64 {
	if x != nil {
		return x.Confidence
	}
	return nil
}

type Bounding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DetectionUid string    `protobuf:"bytes,1,opt,name=detection_uid,json=detectionUid,proto3" json:"detection_uid,omitempty"`
	BoundingType string    `protobuf:"bytes,2,opt,name=bounding_type,json=boundingType,proto3" json:"bounding_type,omitempty"`
	XValues      []float64 `protobuf:"fixed64,3,rep,packed,name=_values,json=Values,proto3" json:"_values,omitempty"` // Assuming this is a list of doubles
}

func (x *Bounding) Reset() {
	*x = Bounding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_detection_parent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bounding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bounding) ProtoMessage() {}

func (x *Bounding) ProtoReflect() protoreflect.Message {
	mi := &file_detection_parent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bounding.ProtoReflect.Descriptor instead.
func (*Bounding) Descriptor() ([]byte, []int) {
	return file_detection_parent_proto_rawDescGZIP(), []int{1}
}

func (x *Bounding) GetDetectionUid() string {
	if x != nil {
		return x.DetectionUid
	}
	return ""
}

func (x *Bounding) GetBoundingType() string {
	if x != nil {
		return x.BoundingType
	}
	return ""
}

func (x *Bounding) GetXValues() []float64 {
	if x != nil {
		return x.XValues
	}
	return nil
}

type ImageAttribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XName  string `protobuf:"bytes,1,opt,name=_name,json=Name,proto3" json:"_name,omitempty"`
	XValue string `protobuf:"bytes,2,opt,name=_value,json=Value,proto3" json:"_value,omitempty"` // Assuming _value is a string, change type if necessary
}

func (x *ImageAttribute) Reset() {
	*x = ImageAttribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_detection_parent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageAttribute) ProtoMessage() {}

func (x *ImageAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_detection_parent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageAttribute.ProtoReflect.Descriptor instead.
func (*ImageAttribute) Descriptor() ([]byte, []int) {
	return file_detection_parent_proto_rawDescGZIP(), []int{2}
}

func (x *ImageAttribute) GetXName() string {
	if x != nil {
		return x.XName
	}
	return ""
}

func (x *ImageAttribute) GetXValue() string {
	if x != nil {
		return x.XValue
	}
	return ""
}

type DetectionImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DetectionImages []string          `protobuf:"bytes,1,rep,name=detection_images,json=detectionImages,proto3" json:"detection_images,omitempty"`
	ImageAttributes []*ImageAttribute `protobuf:"bytes,2,rep,name=image_attributes,json=imageAttributes,proto3" json:"image_attributes,omitempty"`
}

func (x *DetectionImage) Reset() {
	*x = DetectionImage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_detection_parent_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetectionImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetectionImage) ProtoMessage() {}

func (x *DetectionImage) ProtoReflect() protoreflect.Message {
	mi := &file_detection_parent_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetectionImage.ProtoReflect.Descriptor instead.
func (*DetectionImage) Descriptor() ([]byte, []int) {
	return file_detection_parent_proto_rawDescGZIP(), []int{3}
}

func (x *DetectionImage) GetDetectionImages() []string {
	if x != nil {
		return x.DetectionImages
	}
	return nil
}

func (x *DetectionImage) GetImageAttributes() []*ImageAttribute {
	if x != nil {
		return x.ImageAttributes
	}
	return nil
}

type PayloadData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XName  string `protobuf:"bytes,1,opt,name=_name,json=Name,proto3" json:"_name,omitempty"`
	XValue string `protobuf:"bytes,2,opt,name=_value,json=Value,proto3" json:"_value,omitempty"` // Assuming _value is a string, change type if necessary
}

func (x *PayloadData) Reset() {
	*x = PayloadData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_detection_parent_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadData) ProtoMessage() {}

func (x *PayloadData) ProtoReflect() protoreflect.Message {
	mi := &file_detection_parent_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadData.ProtoReflect.Descriptor instead.
func (*PayloadData) Descriptor() ([]byte, []int) {
	return file_detection_parent_proto_rawDescGZIP(), []int{4}
}

func (x *PayloadData) GetXName() string {
	if x != nil {
		return x.XName
	}
	return ""
}

func (x *PayloadData) GetXValue() string {
	if x != nil {
		return x.XValue
	}
	return ""
}

type DetectionEncounter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchemaVersion          string            `protobuf:"bytes,1,opt,name=schema_version,json=schemaVersion,proto3" json:"schema_version,omitempty"`
	MessageUid             string            `protobuf:"bytes,2,opt,name=message_uid,json=messageUid,proto3" json:"message_uid,omitempty"`
	MessageParentUid       string            `protobuf:"bytes,3,opt,name=message_parent_uid,json=messageParentUid,proto3" json:"message_parent_uid,omitempty"`
	RobotUid               string            `protobuf:"bytes,4,opt,name=robot_uid,json=robotUid,proto3" json:"robot_uid,omitempty"`
	EventTs                float64           `protobuf:"fixed64,5,opt,name=event_ts,json=eventTs,proto3" json:"event_ts,omitempty"`
	UtcOffset              string            `protobuf:"bytes,6,opt,name=utc_offset,json=utcOffset,proto3" json:"utc_offset,omitempty"`
	Detections             []*Detection      `protobuf:"bytes,7,rep,name=detections,proto3" json:"detections,omitempty"`
	Boundings              []*Bounding       `protobuf:"bytes,8,rep,name=boundings,proto3" json:"boundings,omitempty"`
	ImageData              []*DetectionImage `protobuf:"bytes,9,rep,name=image_data,json=imageData,proto3" json:"image_data,omitempty"`
	PayloadData            []*PayloadData    `protobuf:"bytes,10,rep,name=payload_data,json=payloadData,proto3" json:"payload_data,omitempty"`
	DetectionsCount        int32             `protobuf:"varint,11,opt,name=detections_count,json=detectionsCount,proto3" json:"detections_count,omitempty"`
	ImageCount             int32             `protobuf:"varint,12,opt,name=image_count,json=imageCount,proto3" json:"image_count,omitempty"`
	BoundingsCount         int32             `protobuf:"varint,13,opt,name=boundings_count,json=boundingsCount,proto3" json:"boundings_count,omitempty"`
	ImageAttributesCount   int32             `protobuf:"varint,14,opt,name=image_attributes_count,json=imageAttributesCount,proto3" json:"image_attributes_count,omitempty"`
	PayloadAttributesCount int32             `protobuf:"varint,15,opt,name=payload_attributes_count,json=payloadAttributesCount,proto3" json:"payload_attributes_count,omitempty"`
}

func (x *DetectionEncounter) Reset() {
	*x = DetectionEncounter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_detection_parent_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetectionEncounter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetectionEncounter) ProtoMessage() {}

func (x *DetectionEncounter) ProtoReflect() protoreflect.Message {
	mi := &file_detection_parent_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetectionEncounter.ProtoReflect.Descriptor instead.
func (*DetectionEncounter) Descriptor() ([]byte, []int) {
	return file_detection_parent_proto_rawDescGZIP(), []int{5}
}

func (x *DetectionEncounter) GetSchemaVersion() string {
	if x != nil {
		return x.SchemaVersion
	}
	return ""
}

func (x *DetectionEncounter) GetMessageUid() string {
	if x != nil {
		return x.MessageUid
	}
	return ""
}

func (x *DetectionEncounter) GetMessageParentUid() string {
	if x != nil {
		return x.MessageParentUid
	}
	return ""
}

func (x *DetectionEncounter) GetRobotUid() string {
	if x != nil {
		return x.RobotUid
	}
	return ""
}

func (x *DetectionEncounter) GetEventTs() float64 {
	if x != nil {
		return x.EventTs
	}
	return 0
}

func (x *DetectionEncounter) GetUtcOffset() string {
	if x != nil {
		return x.UtcOffset
	}
	return ""
}

func (x *DetectionEncounter) GetDetections() []*Detection {
	if x != nil {
		return x.Detections
	}
	return nil
}

func (x *DetectionEncounter) GetBoundings() []*Bounding {
	if x != nil {
		return x.Boundings
	}
	return nil
}

func (x *DetectionEncounter) GetImageData() []*DetectionImage {
	if x != nil {
		return x.ImageData
	}
	return nil
}

func (x *DetectionEncounter) GetPayloadData() []*PayloadData {
	if x != nil {
		return x.PayloadData
	}
	return nil
}

func (x *DetectionEncounter) GetDetectionsCount() int32 {
	if x != nil {
		return x.DetectionsCount
	}
	return 0
}

func (x *DetectionEncounter) GetImageCount() int32 {
	if x != nil {
		return x.ImageCount
	}
	return 0
}

func (x *DetectionEncounter) GetBoundingsCount() int32 {
	if x != nil {
		return x.BoundingsCount
	}
	return 0
}

func (x *DetectionEncounter) GetImageAttributesCount() int32 {
	if x != nil {
		return x.ImageAttributesCount
	}
	return 0
}

func (x *DetectionEncounter) GetPayloadAttributesCount() int32 {
	if x != nil {
		return x.PayloadAttributesCount
	}
	return 0
}

type SubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Add any fields necessary for the subscription request
	ClientId  string   `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	RobotUids []string `protobuf:"bytes,2,rep,name=robot_uids,json=robotUids,proto3" json:"robot_uids,omitempty"` //filters messages to the specific robot_uids
}

func (x *SubscriptionRequest) Reset() {
	*x = SubscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_detection_parent_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionRequest) ProtoMessage() {}

func (x *SubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_detection_parent_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionRequest.ProtoReflect.Descriptor instead.
func (*SubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_detection_parent_proto_rawDescGZIP(), []int{6}
}

func (x *SubscriptionRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *SubscriptionRequest) GetRobotUids() []string {
	if x != nil {
		return x.RobotUids
	}
	return nil
}

var File_detection_parent_proto protoreflect.FileDescriptor

var file_detection_parent_proto_rawDesc = []byte{
	0x0a, 0x16, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6c, 0x70, 0x72, 0x67, 0x72, 0x70,
	0x63, 0x22, 0x99, 0x01, 0x0a, 0x09, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x02, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x01, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x6d, 0x0a,
	0x08, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x74,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x69, 0x64, 0x12, 0x23,
	0x0a, 0x0d, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x01, 0x52, 0x06, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x3c, 0x0a, 0x0e,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x13,
	0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x7f, 0x0a, 0x0e, 0x44, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x29, 0x0a, 0x10,
	0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x42, 0x0a, 0x10, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x6c, 0x70, 0x72, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x0f, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x22, 0x39, 0x0a, 0x0b, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x13, 0x0a, 0x05, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x15, 0x0a, 0x06, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x9c, 0x05, 0x0a, 0x12, 0x44, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x25, 0x0a,
	0x0e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x55, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x55, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x5f, 0x75, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x55, 0x69, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x74, 0x63, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x74, 0x63, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x32, 0x0a, 0x0a, 0x64, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x6c, 0x70, 0x72, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0a, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2f,
	0x0a, 0x09, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x6c, 0x70, 0x72, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x6f, 0x75, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x52, 0x09, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x36, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x09, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6c, 0x70, 0x72, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x09, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x37, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x6c, 0x70, 0x72, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x29, 0x0a, 0x10, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x64, 0x65, 0x74, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f,
	0x62, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x16, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x18, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x16, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x51, 0x0a, 0x13, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x6f, 0x62,
	0x6f, 0x74, 0x5f, 0x75, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x72,
	0x6f, 0x62, 0x6f, 0x74, 0x55, 0x69, 0x64, 0x73, 0x32, 0x68, 0x0a, 0x10, 0x45, 0x6e, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x15,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x44, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x2e, 0x6c, 0x70, 0x72, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6c, 0x70, 0x72, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x30, 0x01, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x66, 0x6c, 0x61, 0x73, 0x68, 0x2d, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x70,
	0x72, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_detection_parent_proto_rawDescOnce sync.Once
	file_detection_parent_proto_rawDescData = file_detection_parent_proto_rawDesc
)

func file_detection_parent_proto_rawDescGZIP() []byte {
	file_detection_parent_proto_rawDescOnce.Do(func() {
		file_detection_parent_proto_rawDescData = protoimpl.X.CompressGZIP(file_detection_parent_proto_rawDescData)
	})
	return file_detection_parent_proto_rawDescData
}

var file_detection_parent_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_detection_parent_proto_goTypes = []interface{}{
	(*Detection)(nil),           // 0: lprgrpc.Detection
	(*Bounding)(nil),            // 1: lprgrpc.Bounding
	(*ImageAttribute)(nil),      // 2: lprgrpc.ImageAttribute
	(*DetectionImage)(nil),      // 3: lprgrpc.DetectionImage
	(*PayloadData)(nil),         // 4: lprgrpc.PayloadData
	(*DetectionEncounter)(nil),  // 5: lprgrpc.DetectionEncounter
	(*SubscriptionRequest)(nil), // 6: lprgrpc.SubscriptionRequest
}
var file_detection_parent_proto_depIdxs = []int32{
	2, // 0: lprgrpc.DetectionImage.image_attributes:type_name -> lprgrpc.ImageAttribute
	0, // 1: lprgrpc.DetectionEncounter.detections:type_name -> lprgrpc.Detection
	1, // 2: lprgrpc.DetectionEncounter.boundings:type_name -> lprgrpc.Bounding
	3, // 3: lprgrpc.DetectionEncounter.image_data:type_name -> lprgrpc.DetectionImage
	4, // 4: lprgrpc.DetectionEncounter.payload_data:type_name -> lprgrpc.PayloadData
	6, // 5: lprgrpc.EncounterService.SubscribeToDetections:input_type -> lprgrpc.SubscriptionRequest
	5, // 6: lprgrpc.EncounterService.SubscribeToDetections:output_type -> lprgrpc.DetectionEncounter
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_detection_parent_proto_init() }
func file_detection_parent_proto_init() {
	if File_detection_parent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_detection_parent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Detection); i {
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
		file_detection_parent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bounding); i {
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
		file_detection_parent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageAttribute); i {
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
		file_detection_parent_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetectionImage); i {
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
		file_detection_parent_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayloadData); i {
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
		file_detection_parent_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetectionEncounter); i {
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
		file_detection_parent_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionRequest); i {
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
			RawDescriptor: file_detection_parent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_detection_parent_proto_goTypes,
		DependencyIndexes: file_detection_parent_proto_depIdxs,
		MessageInfos:      file_detection_parent_proto_msgTypes,
	}.Build()
	File_detection_parent_proto = out.File
	file_detection_parent_proto_rawDesc = nil
	file_detection_parent_proto_goTypes = nil
	file_detection_parent_proto_depIdxs = nil
}