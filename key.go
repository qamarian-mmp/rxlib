package rxlib

// This data type is just a face of data type RxKey. See RxKey for details. This data type
// is meant to be used by a main.
//
// Note that this data is not thread-safe. In other words, it should not be shared by two
// or more goroutines.
type Key interface {

	StartupFailed (string)

	NowRunning ()

	StartupResult () (byte, string)

	Send (interface {}, string) (error)

	Read () (interface {}, error)

	Check () (bool)

	Wait ()

	NewKey (string) (Key, MasterKey, error)

	SystemShutdown ()

	CheckForShutdown () (bool)

	IndicateShutdown ()

	ShutdownState () (byte)
}
