package rxlib

type MasterKey interface {

	StartupResult () (byte, string)

	ShutdownMain ()

	ShutdownState () (bool)
}
