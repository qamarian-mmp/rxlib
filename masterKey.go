package rxlib

// This data type is just a face of data type RxKey. See RxKey for details. The data type
// is meant to be used by Rexa or the master of another main.
type MasterKey interface {

	StartupResult () (byte, string)

	ShutdownMain  ()

	ShutdownState () (byte)
}
