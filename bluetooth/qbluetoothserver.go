package bluetooth

//#include "bluetooth.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QBluetoothServer struct {
	core.QObject
}

type QBluetoothServer_ITF interface {
	core.QObject_ITF
	QBluetoothServer_PTR() *QBluetoothServer
}

func PointerFromQBluetoothServer(ptr QBluetoothServer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothServer_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothServerFromPointer(ptr unsafe.Pointer) *QBluetoothServer {
	var n = new(QBluetoothServer)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QBluetoothServer_") {
		n.SetObjectName("QBluetoothServer_" + qt.Identifier())
	}
	return n
}

func (ptr *QBluetoothServer) QBluetoothServer_PTR() *QBluetoothServer {
	return ptr
}

//QBluetoothServer::Error
type QBluetoothServer__Error int64

const (
	QBluetoothServer__NoError                       = QBluetoothServer__Error(0)
	QBluetoothServer__UnknownError                  = QBluetoothServer__Error(1)
	QBluetoothServer__PoweredOffError               = QBluetoothServer__Error(2)
	QBluetoothServer__InputOutputError              = QBluetoothServer__Error(3)
	QBluetoothServer__ServiceAlreadyRegisteredError = QBluetoothServer__Error(4)
	QBluetoothServer__UnsupportedProtocolError      = QBluetoothServer__Error(5)
)

func NewQBluetoothServer(serverType QBluetoothServiceInfo__Protocol, parent core.QObject_ITF) *QBluetoothServer {
	defer qt.Recovering("QBluetoothServer::QBluetoothServer")

	return NewQBluetoothServerFromPointer(C.QBluetoothServer_NewQBluetoothServer(C.int(serverType), core.PointerFromQObject(parent)))
}

func (ptr *QBluetoothServer) ConnectError2(f func(error QBluetoothServer__Error)) {
	defer qt.Recovering("connect QBluetoothServer::error")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QBluetoothServer) DisconnectError2() {
	defer qt.Recovering("disconnect QBluetoothServer::error")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQBluetoothServerError2
func callbackQBluetoothServerError2(ptr unsafe.Pointer, ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QBluetoothServer::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QBluetoothServer__Error))(QBluetoothServer__Error(error))
	}

}

func (ptr *QBluetoothServer) Error2(error QBluetoothServer__Error) {
	defer qt.Recovering("QBluetoothServer::error")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_Error2(ptr.Pointer(), C.int(error))
	}
}

func (ptr *QBluetoothServer) ConnectNewConnection(f func()) {
	defer qt.Recovering("connect QBluetoothServer::newConnection")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_ConnectNewConnection(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "newConnection", f)
	}
}

func (ptr *QBluetoothServer) DisconnectNewConnection() {
	defer qt.Recovering("disconnect QBluetoothServer::newConnection")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_DisconnectNewConnection(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "newConnection")
	}
}

//export callbackQBluetoothServerNewConnection
func callbackQBluetoothServerNewConnection(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QBluetoothServer::newConnection")

	if signal := qt.GetSignal(C.GoString(ptrName), "newConnection"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothServer) NewConnection() {
	defer qt.Recovering("QBluetoothServer::newConnection")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_NewConnection(ptr.Pointer())
	}
}

func (ptr *QBluetoothServer) Error() QBluetoothServer__Error {
	defer qt.Recovering("QBluetoothServer::error")

	if ptr.Pointer() != nil {
		return QBluetoothServer__Error(C.QBluetoothServer_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServer) IsListening() bool {
	defer qt.Recovering("QBluetoothServer::isListening")

	if ptr.Pointer() != nil {
		return C.QBluetoothServer_IsListening(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServer) MaxPendingConnections() int {
	defer qt.Recovering("QBluetoothServer::maxPendingConnections")

	if ptr.Pointer() != nil {
		return int(C.QBluetoothServer_MaxPendingConnections(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServer) ServerType() QBluetoothServiceInfo__Protocol {
	defer qt.Recovering("QBluetoothServer::serverType")

	if ptr.Pointer() != nil {
		return QBluetoothServiceInfo__Protocol(C.QBluetoothServer_ServerType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServer) DestroyQBluetoothServer() {
	defer qt.Recovering("QBluetoothServer::~QBluetoothServer")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_DestroyQBluetoothServer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothServer) Close() {
	defer qt.Recovering("QBluetoothServer::close")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_Close(ptr.Pointer())
	}
}

func (ptr *QBluetoothServer) HasPendingConnections() bool {
	defer qt.Recovering("QBluetoothServer::hasPendingConnections")

	if ptr.Pointer() != nil {
		return C.QBluetoothServer_HasPendingConnections(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServer) NextPendingConnection() *QBluetoothSocket {
	defer qt.Recovering("QBluetoothServer::nextPendingConnection")

	if ptr.Pointer() != nil {
		return NewQBluetoothSocketFromPointer(C.QBluetoothServer_NextPendingConnection(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBluetoothServer) SetMaxPendingConnections(numConnections int) {
	defer qt.Recovering("QBluetoothServer::setMaxPendingConnections")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_SetMaxPendingConnections(ptr.Pointer(), C.int(numConnections))
	}
}

func (ptr *QBluetoothServer) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QBluetoothServer::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QBluetoothServer) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QBluetoothServer::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQBluetoothServerTimerEvent
func callbackQBluetoothServerTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothServer::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothServerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBluetoothServer) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QBluetoothServer::timerEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QBluetoothServer) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QBluetoothServer::timerEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QBluetoothServer) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QBluetoothServer::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QBluetoothServer) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QBluetoothServer::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQBluetoothServerChildEvent
func callbackQBluetoothServerChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothServer::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQBluetoothServerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QBluetoothServer) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QBluetoothServer::childEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QBluetoothServer) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QBluetoothServer::childEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QBluetoothServer) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QBluetoothServer::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QBluetoothServer) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QBluetoothServer::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQBluetoothServerCustomEvent
func callbackQBluetoothServerCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothServer::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQBluetoothServerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QBluetoothServer) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QBluetoothServer::customEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QBluetoothServer) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QBluetoothServer::customEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}