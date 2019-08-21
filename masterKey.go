package rxlib

type MasterKey interface {

	// StartupResult () gives the startup result of the 'main' using the corresponding
	// key of this key.
	//
	// Outpts
	//
	// outpt 0: The startup result of the main. Possible values should be checked in
	// the variable section of this package.
	//
	// outpt 1: If value of outpt 0 is SrStartupFailed, value of this data would be a
	// text describing the reason for the failure.
	StartupResult () (byte, string)

	// ShutdownMain  () could be used to signal shutdown to the 'main' using the
	// corresponding key of this key.
	ShutdownMain  ()

	// ShutdownState () could be used to get the shutdown state of the 'main' using
	// the corresponding key of this key. True means the 'main' is down, while false
	// means the 'main' is still active or has not run at all.
	ShutdownState () (bool)
}
