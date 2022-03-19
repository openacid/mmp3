# mmp3: multi-master Paxos based kv storage

![main](https://github.com/openacid/mmp3/workflows/test/badge.svg?branch=main)
[![travis](https://travis-ci.com/openacid/mmp3.svg?branch=naive)](https://travis-ci.com/openacid/mmp3)

![](https://blog.openacid.com/post-res/mmp3/digraphqueue_demosize=1010dpi=10-b5ab13197d2fba30.jpg)

mmp3 中文介绍: https://blog.openacid.com/algo/mmp3

MMP3 is a multi-master distributed consensus protocol built on top of classic Paxos.
The application API is similar to raft except that **every node is a leader** and can handle write requests. 

- Exhibits a distributed **log-state-machine** model, similar to raft.
- Exactly three nodes in a cluster.
- Every node is a leader and is able to handle write requests without conflict.
- A message is committed in **exactly one RTT**(unlike epaxos which requires two RTT if logs interfere with others).

# Question

Any questions or suggestions, feel free to open a [issue][] :DDD.


[issue]:                          https://github.com/openacid/mmp3/issues/new/choose
