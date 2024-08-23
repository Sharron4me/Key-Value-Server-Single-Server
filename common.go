package kvsrv

// Put or Append
type PutAppendArgs struct {
	Key   string
	Value string
	Op    string // "Put" or "Append"
	// You can add more fields if needed, like:
	// ClientID string
	// RequestID int
	RequestID int
}

type PutAppendReply struct {
	Value string
}

type GetArgs struct {
	Key string
	// Add more fields if needed, like:
	// ClientID string
}

type GetReply struct {
	Value string
}

type RequestQueue struct {
	Requests []PutAppendArgs
}
