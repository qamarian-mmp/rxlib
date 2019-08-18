package rxlib

type Key interface {

	State (byte, string)

	StartupFailed (string)

	NowRunning ()

	Send (interface {}, string) error

	Read (interface {}, error)

	Shutdown ()

	ShutdownIssued () bool

	MarkShutdown (string)
}

var (
	StateStartup        byte = 0
	StateFailedStartup  byte = 1
	StateRunning        byte = 2
	StateError          byte = 3
	StateDone           byte = 4
)
