package kvsrv

import (
	"log"
	"strings"
	"sync"
)

const Debug = false

func DPrintf(format string, a ...interface{}) (n int, err error) {
	if Debug {
		log.Printf(format, a...)
	}
	return
}

type KVServer struct {
	mu            sync.Mutex
	kv_store      map[string]string
	seen_requests map[int]bool
}

func (kv *KVServer) Get(args *GetArgs, reply *GetReply) {
	// Your code here.
	kv.mu.Lock()
	defer kv.mu.Unlock()
	value, exist := kv.kv_store[args.Key]
	if exist {
		reply.Value = value
	} else {
		reply.Value = ""
	}
}

func (kv *KVServer) Put(args *PutAppendArgs, reply *PutAppendReply) {
	// Your code here.
	kv.mu.Lock()
	defer kv.mu.Unlock()
	_, exists := kv.seen_requests[args.RequestID]
	if exists {
		return
	}
	kv.kv_store[args.Key] = args.Value
	kv.seen_requests[args.RequestID] = true
}

func (kv *KVServer) Append(args *PutAppendArgs, reply *PutAppendReply) {
	// Your code here.
	kv.mu.Lock()
	defer kv.mu.Unlock()
	_, exists := kv.seen_requests[args.RequestID]
	if exists {
		reply.Value = strings.Split(kv.kv_store[args.Key], args.Value)[0]
		return
	}
	value, exist := kv.kv_store[args.Key]
	old_value := ""
	if exist {
		old_value = value
		kv.kv_store[args.Key] = value + args.Value
	} else {
		kv.kv_store[args.Key] = args.Value
	}
	kv.seen_requests[args.RequestID] = true
	reply.Value = old_value
}

func StartKVServer() *KVServer {
	kv := new(KVServer)
	kv.kv_store = make(map[string]string)
	kv.seen_requests = make(map[int]bool)
	return kv
}
