package rxlib

type Message struct { // This data type is the type of messages being sent by threads of a mopware

	Id            string // ID of the message
	Type          string /* Message type. Possible values are: new (stands for new message), rpl
		(stands for reply), cnw (stands for continuation of a new message), crp (stands for
		continuation of a reply), err (stands for an error message) */
	ObjectId      string /* The id of another message being refered to. For example, if a
		message's type is "rpl", value of this would be the id of the message this message
		is replying. */
	SenderType    string /* The id of the thread type of the sender. Used alongside
		"SenderInsId", to uniquely the sender of the message */
	SenderInsId   string // The instance id of sender (relative to other threads of its type)
	ReceiverType  string /* The id of the thread type of the receiver. Used alongside
		"ReceiverInsId", to uniquely the sender of the message */
	ReceiverInsId string // The instance id of receiver (relative to other threads of its type)
	Message       interface {} // The actual message being sent
}
