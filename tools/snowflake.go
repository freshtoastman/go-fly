package tools

import (
	"errors"
	"sync"
	"time"
)

/*
* Snowflake
*
* 1                                               42           52             64
* +-----------------------------------------------+------------+---------------+
* | timestamp(ms)                                 | workerid   | sequence      |
* +-----------------------------------------------+------------+---------------+
* | 0000000000 0000000000 0000000000 0000000000 0 | 0000000000 | 0000000000 00 |
* +-----------------------------------------------+------------+---------------+
*
* 1. 41位時间截(毫秒级)，注意这是時间截的差值（當前時间截 - 开始時间截)。可以使用约70年: (1L << 41) / (1000L * 60 * 60 * 24 * 365) = 69
* 2. 10位數據機器位，可以部署在1024个节点
* 3. 12位序列，毫秒内的計數，同一機器，同一時间截並發4096个序號
 */

const (
	twepoch        = int64(1483228800000)             //开始時间截 (2017-01-01)
	workeridBits   = uint(10)                         //機器id所占的位數
	sequenceBits   = uint(12)                         //序列所占的位數
	workeridMax    = int64(-1 ^ (-1 << workeridBits)) //支持的最大機器id數量
	sequenceMask   = int64(-1 ^ (-1 << sequenceBits)) //
	workeridShift  = sequenceBits                     //機器id左移位數
	timestampShift = sequenceBits + workeridBits      //時间戳左移位數
)

// A Snowflake struct holds the basic information needed for a snowflake generator worker
type Snowflake struct {
	sync.Mutex
	timestamp int64
	workerid  int64
	sequence  int64
}

// NewNode returns a new snowflake worker that can be used to generate snowflake IDs
func NewSnowflake(workerid int64) (*Snowflake, error) {

	if workerid < 0 || workerid > workeridMax {
		return nil, errors.New("workerid must be between 0 and 1023")
	}

	return &Snowflake{
		timestamp: 0,
		workerid:  workerid,
		sequence:  0,
	}, nil
}

// Generate creates and returns a unique snowflake ID
func (s *Snowflake) Generate() int64 {

	s.Lock()

	now := time.Now().UnixNano() / 1000000

	if s.timestamp == now {
		s.sequence = (s.sequence + 1) & sequenceMask

		if s.sequence == 0 {
			for now <= s.timestamp {
				now = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		s.sequence = 0
	}

	s.timestamp = now

	r := int64((now-twepoch)<<timestampShift | (s.workerid << workeridShift) | (s.sequence))

	s.Unlock()
	return r
}
