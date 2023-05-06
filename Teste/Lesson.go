// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Teste

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Lesson struct {
	_tab flatbuffers.Table
}

func GetRootAsLesson(buf []byte, offset flatbuffers.UOffsetT) *Lesson {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Lesson{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Lesson) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Lesson) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Lesson) Id() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Lesson) MutateId(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func (rcv *Lesson) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Lesson) Description() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Lesson) CourseId() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Lesson) MutateCourseId(n int32) bool {
	return rcv._tab.MutateInt32Slot(10, n)
}

func LessonStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func LessonAddId(builder *flatbuffers.Builder, id int32) {
	builder.PrependInt32Slot(0, id, 0)
}
func LessonAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(name), 0)
}
func LessonAddDescription(builder *flatbuffers.Builder, description flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(description), 0)
}
func LessonAddCourseId(builder *flatbuffers.Builder, courseId int32) {
	builder.PrependInt32Slot(3, courseId, 0)
}
func LessonEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
