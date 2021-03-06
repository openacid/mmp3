package paxoskv

import (
	"sort"

	"golang.org/x/net/context"
)

func (s *KVServer) HandlePaxos(c context.Context, req *Request) (*Reply, error) {
	_ = c

	dd(s, "recv HandlePaxos-req: %s", req.str())

	s.Lock()
	defer s.Unlock()

	col := s.log.columns[req.Column]

	reply := &Reply{
		LastBal:   col.LastBal.Clone(),
		Instances: make(map[int64]*Ins),
	}

	if req.Bal.Less(col.LastBal) {
		return reply, nil
	}

	// granted

	col.LastBal = req.Bal.Clone()

	lsns := make([]int64, 0)
	for lsn := range req.Instances {
		lsns = append(lsns, lsn)
	}

	sort.Slice(lsns, func(i, j int) bool {
		return lsns[i] < lsns[j]
	})

	for _, lsn := range lsns {
		inst := req.Instances[lsn]
		inst = s.hdlOps(req, inst)
		reply.Instances[lsn] = inst
	}

	dd(s, "send Reply: %s", reply.str())

	return reply, nil
}

// Set impl the KV API and handles a Set request from client.
// Only the Key and Vi64 should be set in req.
func (s *KVServer) Set(c context.Context, cmd *Cmd) (*Cmd, error) {

	dd(s, "hdl Set: %s", cmd.str())

	for {
		select {
		case <-c.Done():
			return nil, c.Err()
		default:
		}

		s.Lock()
		inst := s.allocNewInst(s.Id, cmd)
		s.Unlock()

		h := NewHandler(s, inst)
		dd(h, "hdl set: %s", cmd.str())

		lsn := inst.getLSN()
		committed := h.runPaxosLoop(s.Id, lsn, s.other)

		dd(h, "committed: %s, proposed: %s", committed.str(), inst.str())
		if committed.InsId.Equal(inst.InsId) {
			return cmd, nil
		}
		dd(h, "another value committed, need to re-commit: %s", inst.str())
	}
}

// Get impl the KV-API get method.
// Only req.Key should be specified.
func (s *KVServer) Get(c context.Context, req *Cmd) (*Cmd, error) {
	_ = c
	return s.stateMachine.get(req.Key)
}
