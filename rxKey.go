package rxlib

import (
	"gopkg.in/qamarian-dtp/rnet.v1"
	"sync"
)

// NewRxKey () helps create a new 'Rx Key'. The function requires three (3) inputs namely:
//	- a communication channel (PPO) for the key
//	- a data that could be used to signal shutdown to the system
//	- the communication network that powers the system
func NewRxKey (commChan *rnet.PPO, shutChan *sync.Cond, commNet *rnet.NetCentre) (*RxKey){
	return &RxKey {
		commChan:           commChan,
		startupResult:      SrUnavailable,
		startupNote:          "",
		systemShutdownChan: shutChan,
		shutdownSignal:     false,
		shutdownState:      SsNotApplicable,
		commNetCentre:      commNet,
	}
}

// This data type should never be manipulated directly. To manipulate it, use any of its
// faces that suit you.
type RxKey struct {
	commChan           *rnet.PPO       // The channel the key uses for communication.
	startupResult      byte            // The startup result of the key's main,
	startupNote        string          // The startup note of the key's main.
	systemShutdownChan *sync.Cond      /* The data that could be used to signal
		shutdown to the system. */
	shutdownSignal     bool            /* The data indicating if the key's main has
		been asked to shutdown or not. */
	shutdownState      byte            /* The data indicating if the key's main has
		been shutdown or not. */
	commNetCentre      *rnet.NetCentre /* The network making it possible for the main
		to communicate with other mains in the system. */
}





// ----- Master key methods -----

// StartupResult () gives the startup result of the main using the key.
//
// Outpts
//
// outpt 0: The startup result of the main. Possible values should be checked in the
// variable section of this package.
//
// outpt 1: If value of outpt 0 is SrStartupFailed, value of this data would be a text
// describing the reason for the failure, otherwise, value would be an empty string.
func (rxk *RxKey) StartupResult () (byte, string) {
	return rxk.startupResult, rxk.startupNote
}

// ShutdownMain  () could be used to signal shutdown to the main using this key.
func (rxk *RxKey) ShutdownMain () {
	rxk.shutdownSignal = true
}





// ----- Normal key methods -----

// StartupFailed () should be called if the main is unable to startup successfully. The
// reason for startup failure should be provided as the input of this method.
func (rxk *RxKey) StartupFailed (note string) {
	rxk.startupResult = SrStartupFailed
	rxk.startupNote = note
}

// NowRunning () should be called if the main is able to startup successfully.
func (rxk *RxKey) NowRunning () {
	rxk.startupResult = SrStartedUp
	rxk.shutdownState = SsStillRunning
}

// Send () could be used to send messages to the other mains in the system.
func (rxk *RxKey) Send (mssg interface {}, recipient string) (error) {
	return rxk.commChan.Send (mssg, recipient)
}

// Read () could be used to read received messages, on at a time.
func (rxk *RxKey) Read () (interface {}, error) {
	return rxk.commChan.Read ()
}

// Check () could be used to check if there is any new message that could be read. True
// would mean there is a message that could be read, while false would means there is no
// message that could be read.
func (rxk *RxKey) Check () (bool) {
	return rxk.commChan.Check ()
}

// Wait () could be used to keep a main waiting until a new message has been received.
// When a new message has been received, Read () could then be used to read the new
// message. Unlike a combination of a for loop and method Check (), Wait does not waste
// CPU cycles while waiting. In other words, whenever a main needs to wait till a new
// message has been received, it is recommended this method is called.
func (rxk *RxKey) Wait () {
	rxk.commChan.Wait ()
}

// NewKey () could be used to get a new key. Let's assume, a new main would be created at
// runtime, the main would also require a key, to join the system. In such a situation,
// this is the method that helps out. This method creates a new key that could be used by
// the new main, to communicate with the other mains in the system.
//
// Outpts
//
// outpt 0: The new key.
//
// outpt 1: A master key that could be used to communicate with the new key.
//
// outpt 2: On success, value would be nil. On failure, value would be an error.
func (rxk *RxKey) NewKey (id string) (Key, MasterKey, error) {
	commChan, errX := rxk.commNetCentre.NewPPO (id)
	if errX != nil {
		return nil, nil, errX
	}
	key := NewRxKey (commChan, rxk.systemShutdownChan, rxk.commNetCentre)
	return key, key, nil
}

// SystemShutdown () could be used by a main, to gracefully shutdown the whole system.
func (rxk *RxKey) SystemShutdown () {
	rxk.systemShutdownChan.Signal ()
}

// CheckForShutdown () could be used by a main, to check if it has been asked to shutdown.
// True would mean it has been asked to shutdown, while false would mean it is yet to be
// asked to shutdown.
func (rxk *RxKey) CheckForShutdown () (bool) {
	return rxk.shutdownSignal
}

// IndicateShutdown () could be used by a main, to show that is has been shutdown. Before
// any main returns, they should call  this method on their key, otherwise the main would
// be assumed to still be running, and the system may become unable to shutdown
// gracefully.
func (rxk *RxKey) IndicateShutdown () {
	rxk.shutdownState = SsHasShutdown
}





// ----- Common methods -----

// ShutdownState () could be used to get the shutdown state of the main using this key.
// Check posssible values in the variable section.
func (rxk *RxKey) ShutdownState () (byte) {
	return rxk.shutdownState
}

var (
	// Startup results
	SrUnavailable   byte = 0 // This means the main has neither started up nor failed.
	SrStartupFailed byte = 1 // This means the main could not start up successfully.
	SrStartedUp     byte = 2 // This means the main started up successfully.

	// Shutdown states
	SsNotApplicable byte = 0 // This means the main is yet to start up successfully.
	SsStillRunning  byte = 1 // This means the main is still running.
	SsHasShutdown   byte = 2 // This means the main has shutdown.
)
