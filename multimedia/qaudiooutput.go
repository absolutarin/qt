package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAudioOutput struct {
	core.QObject
}

type QAudioOutput_ITF interface {
	core.QObject_ITF
	QAudioOutput_PTR() *QAudioOutput
}

func PointerFromQAudioOutput(ptr QAudioOutput_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioOutput_PTR().Pointer()
	}
	return nil
}

func NewQAudioOutputFromPointer(ptr unsafe.Pointer) *QAudioOutput {
	var n = new(QAudioOutput)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAudioOutput_") {
		n.SetObjectName("QAudioOutput_" + qt.Identifier())
	}
	return n
}

func (ptr *QAudioOutput) QAudioOutput_PTR() *QAudioOutput {
	return ptr
}

func NewQAudioOutput2(audioDevice QAudioDeviceInfo_ITF, format QAudioFormat_ITF, parent core.QObject_ITF) *QAudioOutput {
	defer qt.Recovering("QAudioOutput::QAudioOutput")

	return NewQAudioOutputFromPointer(C.QAudioOutput_NewQAudioOutput2(PointerFromQAudioDeviceInfo(audioDevice), PointerFromQAudioFormat(format), core.PointerFromQObject(parent)))
}

func NewQAudioOutput(format QAudioFormat_ITF, parent core.QObject_ITF) *QAudioOutput {
	defer qt.Recovering("QAudioOutput::QAudioOutput")

	return NewQAudioOutputFromPointer(C.QAudioOutput_NewQAudioOutput(PointerFromQAudioFormat(format), core.PointerFromQObject(parent)))
}

func (ptr *QAudioOutput) BufferSize() int {
	defer qt.Recovering("QAudioOutput::bufferSize")

	if ptr.Pointer() != nil {
		return int(C.QAudioOutput_BufferSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioOutput) BytesFree() int {
	defer qt.Recovering("QAudioOutput::bytesFree")

	if ptr.Pointer() != nil {
		return int(C.QAudioOutput_BytesFree(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioOutput) Category() string {
	defer qt.Recovering("QAudioOutput::category")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioOutput_Category(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioOutput) ElapsedUSecs() int64 {
	defer qt.Recovering("QAudioOutput::elapsedUSecs")

	if ptr.Pointer() != nil {
		return int64(C.QAudioOutput_ElapsedUSecs(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioOutput) Error() QAudio__Error {
	defer qt.Recovering("QAudioOutput::error")

	if ptr.Pointer() != nil {
		return QAudio__Error(C.QAudioOutput_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioOutput) ConnectNotify(f func()) {
	defer qt.Recovering("connect QAudioOutput::notify")

	if ptr.Pointer() != nil {
		C.QAudioOutput_ConnectNotify(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "notify", f)
	}
}

func (ptr *QAudioOutput) DisconnectNotify() {
	defer qt.Recovering("disconnect QAudioOutput::notify")

	if ptr.Pointer() != nil {
		C.QAudioOutput_DisconnectNotify(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "notify")
	}
}

//export callbackQAudioOutputNotify
func callbackQAudioOutputNotify(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAudioOutput::notify")

	if signal := qt.GetSignal(C.GoString(ptrName), "notify"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAudioOutput) Notify() {
	defer qt.Recovering("QAudioOutput::notify")

	if ptr.Pointer() != nil {
		C.QAudioOutput_Notify(ptr.Pointer())
	}
}

func (ptr *QAudioOutput) NotifyInterval() int {
	defer qt.Recovering("QAudioOutput::notifyInterval")

	if ptr.Pointer() != nil {
		return int(C.QAudioOutput_NotifyInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioOutput) PeriodSize() int {
	defer qt.Recovering("QAudioOutput::periodSize")

	if ptr.Pointer() != nil {
		return int(C.QAudioOutput_PeriodSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioOutput) ProcessedUSecs() int64 {
	defer qt.Recovering("QAudioOutput::processedUSecs")

	if ptr.Pointer() != nil {
		return int64(C.QAudioOutput_ProcessedUSecs(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioOutput) Reset() {
	defer qt.Recovering("QAudioOutput::reset")

	if ptr.Pointer() != nil {
		C.QAudioOutput_Reset(ptr.Pointer())
	}
}

func (ptr *QAudioOutput) Resume() {
	defer qt.Recovering("QAudioOutput::resume")

	if ptr.Pointer() != nil {
		C.QAudioOutput_Resume(ptr.Pointer())
	}
}

func (ptr *QAudioOutput) SetBufferSize(value int) {
	defer qt.Recovering("QAudioOutput::setBufferSize")

	if ptr.Pointer() != nil {
		C.QAudioOutput_SetBufferSize(ptr.Pointer(), C.int(value))
	}
}

func (ptr *QAudioOutput) SetCategory(category string) {
	defer qt.Recovering("QAudioOutput::setCategory")

	if ptr.Pointer() != nil {
		C.QAudioOutput_SetCategory(ptr.Pointer(), C.CString(category))
	}
}

func (ptr *QAudioOutput) SetNotifyInterval(ms int) {
	defer qt.Recovering("QAudioOutput::setNotifyInterval")

	if ptr.Pointer() != nil {
		C.QAudioOutput_SetNotifyInterval(ptr.Pointer(), C.int(ms))
	}
}

func (ptr *QAudioOutput) SetVolume(volume float64) {
	defer qt.Recovering("QAudioOutput::setVolume")

	if ptr.Pointer() != nil {
		C.QAudioOutput_SetVolume(ptr.Pointer(), C.double(volume))
	}
}

func (ptr *QAudioOutput) Start2() *core.QIODevice {
	defer qt.Recovering("QAudioOutput::start")

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QAudioOutput_Start2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAudioOutput) Start(device core.QIODevice_ITF) {
	defer qt.Recovering("QAudioOutput::start")

	if ptr.Pointer() != nil {
		C.QAudioOutput_Start(ptr.Pointer(), core.PointerFromQIODevice(device))
	}
}

func (ptr *QAudioOutput) State() QAudio__State {
	defer qt.Recovering("QAudioOutput::state")

	if ptr.Pointer() != nil {
		return QAudio__State(C.QAudioOutput_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioOutput) ConnectStateChanged(f func(state QAudio__State)) {
	defer qt.Recovering("connect QAudioOutput::stateChanged")

	if ptr.Pointer() != nil {
		C.QAudioOutput_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QAudioOutput) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QAudioOutput::stateChanged")

	if ptr.Pointer() != nil {
		C.QAudioOutput_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQAudioOutputStateChanged
func callbackQAudioOutputStateChanged(ptr unsafe.Pointer, ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QAudioOutput::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QAudio__State))(QAudio__State(state))
	}

}

func (ptr *QAudioOutput) StateChanged(state QAudio__State) {
	defer qt.Recovering("QAudioOutput::stateChanged")

	if ptr.Pointer() != nil {
		C.QAudioOutput_StateChanged(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QAudioOutput) Stop() {
	defer qt.Recovering("QAudioOutput::stop")

	if ptr.Pointer() != nil {
		C.QAudioOutput_Stop(ptr.Pointer())
	}
}

func (ptr *QAudioOutput) Suspend() {
	defer qt.Recovering("QAudioOutput::suspend")

	if ptr.Pointer() != nil {
		C.QAudioOutput_Suspend(ptr.Pointer())
	}
}

func (ptr *QAudioOutput) Volume() float64 {
	defer qt.Recovering("QAudioOutput::volume")

	if ptr.Pointer() != nil {
		return float64(C.QAudioOutput_Volume(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioOutput) DestroyQAudioOutput() {
	defer qt.Recovering("QAudioOutput::~QAudioOutput")

	if ptr.Pointer() != nil {
		C.QAudioOutput_DestroyQAudioOutput(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAudioOutput) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAudioOutput::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAudioOutput) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAudioOutput::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAudioOutputTimerEvent
func callbackQAudioOutputTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioOutput::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAudioOutputFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAudioOutput) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAudioOutput::timerEvent")

	if ptr.Pointer() != nil {
		C.QAudioOutput_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAudioOutput) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAudioOutput::timerEvent")

	if ptr.Pointer() != nil {
		C.QAudioOutput_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAudioOutput) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAudioOutput::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAudioOutput) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAudioOutput::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAudioOutputChildEvent
func callbackQAudioOutputChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioOutput::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAudioOutputFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAudioOutput) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAudioOutput::childEvent")

	if ptr.Pointer() != nil {
		C.QAudioOutput_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAudioOutput) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAudioOutput::childEvent")

	if ptr.Pointer() != nil {
		C.QAudioOutput_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAudioOutput) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAudioOutput::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAudioOutput) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAudioOutput::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAudioOutputCustomEvent
func callbackQAudioOutputCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioOutput::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAudioOutputFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAudioOutput) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAudioOutput::customEvent")

	if ptr.Pointer() != nil {
		C.QAudioOutput_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAudioOutput) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAudioOutput::customEvent")

	if ptr.Pointer() != nil {
		C.QAudioOutput_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}