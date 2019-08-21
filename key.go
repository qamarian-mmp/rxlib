package rxlib

type Key interface {

	// StartupFailed() should be called if the main is unable to startup successfully.
	// The reason for startup failure should be provided as the input of this method.
	StartupFailed (string)

	// NowRunning () should be called if the main is able to startup successfully.
	NowRunning ()

	// Send () could be used to send messages to other mains in the system.
	Send (interface {}, string) (error)

	// Read () could be used to read received messages, on at a time.
	Read () (interface {}, error)

	// Check () could be used to check if there is any new message that could be read.
	Check () (bool)

	// Wait () could be used to keep a main waiting until a new message has been
	// received. When a new message has been received, Read () could then be used to
	// read the new message. Unlike a combination of a for loop and method Check (),
	// Wait does not waste CPU cycles while waiting. In other words, whenever a main
	// needs to wait till a new message has been received, it is recommended this
	// method is called.
	Wait ()

	// NewKey () could be used to get a new key. Let's assume, a new main would be
	// created at runtime, the main would also require a key, to join the system. In
	// such a situation, this is the method that helps out. This method creates a new
	// key that could be used by the new main, to communicate with other mains in the
	// system.
	//
	// Outpts
	//
	// outpt 0: The new key.
	//
	// outpt 1: A master key that could be used to communicate with the new key.
	//
	// outpt 2: On success, value would be nil. On failure, value would be an error.
	NewKey (string) (Key, MasterKey, error)

	// SystemShutdown () could be used by a main, to gracefully shutdown the whole
	// system.
	SystemShutdown ()

	// CheckForShutdown () could be used by a main, to check if it has been asked to
	// shutdown. True would mean it has been asked to shutdown, while false would
	// mean it is yet to be asked to shutdown.
	CheckForShutdown () (bool)

	// IndicateShutdown () could be used by a main, to show that is has been shutdown.
	// Before any main returns, they should call this method on their key, otherwise
	// the main would be assumed to still be running, and the system may become unable
	// to shutdown gracefully.
	IndicateShutdown ()
}
