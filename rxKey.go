package rxlib

import (
	"gopkg.in/qamarian-dtp/rnet.v1"
	"sync"
)

// NewRxKey () helps create a new Rx Key. This function is meant to be used by rexa.
func NewRxKey (commChan *rnet.PPO, shutChan *sync.Cond, commNet *rnet.NetCentre) (*RxKey){
	return &RxKey {
		commChan:           commChan,
		startupResult:      SrUnavailable,
		startupNote:          "",
		systemShutdownChan: shutChan,
		shutdownSignal:     false,
		shutdownState:      false,
		commNetCentre:      commNet,
	}
}

type RxKey struct {
	commChan           *rnet.PPO       // The channel the key uses for communication.
	startupResult      byte            // The startup result of the key's main,
	startupNote        string          // The startup note of the key's main.
	systemShutdownChan *sync.Cond      /* The data that could be used to signal
		shutdown to the system. */
	shutdownSignal     bool            /* The data indicating if the key's main has
		been asked to shutdown or not. */
	shutdownState      bool            /* The data indicating if the key's main has
		been shutdown or not. */
	commNetCentre      *rnet.NetCentre /* The network making it possible for the main
		of the system to communicate. */
}

// Master key methods

func (rxk *RxKey) StartupResult () (byte, string) {
	return rxk.startupResult, rxk.startupNote
}

func (rxk *RxKey) ShutdownMain () {
	rxk.shutdownSignal = true
}

func (rxk *RxKey) ShutdownState () (bool) {
	return rxk.shutdownState
}

// Normal key methods

func (rxk *RxKey) StartupFailed (note string) {
	rxk.startupResult = SrStartupFailed
	rxk.startupNote = note
}

func (rxk *RxKey) NowRunning () {
	rxk.startupResult = SrNowRunning
}

func (rxk *RxKey) Send (mssg interface {}, recipient string) (error) {
	return rxk.commChan.Send (mssg, recipient)
}

func (rxk *RxKey) Read () (interface {}, error) {
	return rxk.commChan.Read ()
}

func (rxk *RxKey) Check () (bool) {
	return rxk.commChan.Check ()
}

func (rxk *RxKey) Wait () {
	rxk.commChan.Wait ()
}

func (rxk *RxKey) NewKey (id string) (Key, MasterKey, error) {
	commChan, errX := rxk.commNetCentre.NewPPO (id)
	if errX != nil {
		return nil, nil, errX
	}
	key := NewRxKey (commChan, rxk.systemShutdownChan, rxk.commNetCentre)
	return key, key, nil
}

func (rxk *RxKey) SystemShutdown () {
	rxk.systemShutdownChan.Signal ()
}

func (rxk *RxKey) CheckForShutdown () (bool) {
	return rxk.shutdownSignal
}

func (rxk *RxKey) IndicateShutdown () {
	rxk.shutdownState = true
}

var (
	SrUnavailable   byte = 0
	SrStartupFailed byte = 1
	SrNowRunning    byte = 2
)
