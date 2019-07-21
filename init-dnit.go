package rxlib

// This type is the data type of the initialization function of mopthreads.
type InitFn func (string, string, *<-chan *Message, *chan<- *Message, byte, byte)

// This type is the data type of the deinitialization function of mopthreads.
type DnitFn func ()
