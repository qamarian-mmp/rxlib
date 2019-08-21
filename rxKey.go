package rxlib

import (
	"gopkg.in/qamarian-dtp/rnet.v1"
	"sync"
)

func NewRxKey (commChan *rnet.PPO, shutChan *sync.Cond, commNet *rnet.NetCentre) (*RxKey) {
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
	commChan           *rnet.PPO
	startupResult      byte
	startupNote        string
	systemShutdownChan *sync.Cond
	shutdownSignal     bool
	shutdownState      bool
	commNetCentre      *rnet.NetCentre
}

// Master operations

func (rxk *RxKey) StartupResult () (byte, string) {
	return rxk.startupResult, rxk.startupNote
}

func (rxk *RxKey) ShutdownMain () {
	rxk.shutdownSignal = true
}

func (rxk *RxKey) ShutdownState () (bool) {
	return rxk.shutdownState
}

// Main operations

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
