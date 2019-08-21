package rxlib

// RxLog is an abstract data type. It serves as Rexa's log. Although, rexa has a default
// implementation of this ADT, feel free to create another implementation of this ADT,
// and replace the default with whatever you create.
type RxLog interface {

	// Record () could be used to add a new record to the log. The first input
	// should be the string you want to add, while the second input should be the
	// type of record you are adding. See the variable section for possible values
	// of the second input.
	Record (string, byte) (error)
}

var (
	// Log record types
	LrtStandard byte = 0 // The could be used to represent a normal log.
	LrtWarning  byte = 1 // The could be used to represent a warning log.
	LrtError    byte = 2 // The could be used to represent an error log.
)
