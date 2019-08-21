package rxlib

// This data type is just a face of data type RxKey. See RxKey for details.
type MasterKey interface {

	StartupResult () (byte, string)
	
	ShutdownMain  ()

	ShutdownState () (byte)
}
