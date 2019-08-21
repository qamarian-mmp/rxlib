package rxlib

// This data type is just a face of data type RxKey. See RxKey for details. This data type
// is meant to be used by a main.
type Key interface {

	StartupFailed (string)

	NowRunning ()

	Send (interface {}, string) (error)

	Read () (interface {}, error)

	Check () (bool)

	Wait ()

	NewKey (string) (Key, MasterKey, error)

	SystemShutdown ()

	CheckForShutdown () (bool)

	IndicateShutdown ()
}
