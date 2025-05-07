package utils

import (
	"log"
	"sync"
	"time"
)

const (
	epoch          = 903024000000
	timestampBits  = 41
	machineIDBits  = 10
	sequenceDits   = 12
	maxTimestamp   = -1 ^ (-1 << timestampBits)
	maxMachineID   = -1 ^ (-1 << machineIDBits)
	maxSequenceNum = -1 ^ (-1 << sequenceDits)
	machineShift   = sequenceDits
	timestampShift = sequenceDits + machineIDBits
)

type Snowflake struct {
	sync.Mutex
	timestamp   int64
	machineID   int64
	sequenceNum int64
}

func NewSnowflake(machineID int64) *Snowflake {
	if machineID < 0 || machineID > maxMachineID {
		log.Fatalf("utils.snowflake: machineId must be between 0 and %d \n", maxMachineID-1)
	}
	return &Snowflake{
		timestamp:   0,
		machineID:   machineID,
		sequenceNum: 0,
	}
}

func (s *Snowflake) GenerateID() int64 {
	s.Lock()
	now := time.Now().UnixNano() / 1e6
	if s.timestamp == now {
		s.sequenceNum = (s.sequenceNum + 1) & maxSequenceNum
		if s.sequenceNum == 0 {
			for now <= s.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		s.sequenceNum = 0
	}
	t := now - epoch
	if t > maxTimestamp {
		s.Unlock()
		log.Fatalf("utils.snowflake: epoch must be between 0 and %d\n", maxTimestamp-1)
	}
	s.timestamp = now
	r := (t)<<timestampShift | (s.machineID << machineShift) | (s.sequenceNum)
	s.Unlock()
	return r
}
