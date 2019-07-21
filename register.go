package rxlib

type ThreadTypeRegister struct { // This type is the data type of records found in rexa's registry
	DepIds    []string // IDs of the dependencies of the thread type
	InitFunc  InitFn   // Initialization function of threads of this type
	DnitFunc  DnitFn   // Deinitialization function of threads of this type
	Singleton bool     /* Data telling if a thread type is singleton i.e. only one instance of
		it can exist. When value is "true", thread type would be treated as singleton. If
		value is false, thread type would not be treated as singleton. */
	Daemonic  bool     /* Data telling if an instance of a thread type should be spinned up,
		on start-up. When value is "true", thread type would be treated as daemonic. If
		value is false, thread type would not be treated as daemonic. */
}
