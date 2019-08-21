package rxlib

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
