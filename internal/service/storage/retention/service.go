package retention

import (
	"fmt"
	"sort"
)

type RetentionService interface {
	Describe(retentionID uint64) (*Retention, error)
	List(cursor uint64, limit uint64) ([]Retention, error)
	Create(Retention) (uint64, error)
	Update(retentionID uint64, retention Retention) error
	Remove(retentionID uint64) (bool, error)
}

type DummyRetentionService struct{}

func NewDummyRetentionService() *DummyRetentionService {
	return &DummyRetentionService{}
}

// вообще, по названию, ф-ция должна делать что-то другое, но по сигнатуре
// похожа на get, пока будем использовать её именно так
func (s *DummyRetentionService) Describe(retentionID uint64) (*Retention, error) {
	for _, v := range allRetention {
		if v.RetentionID == retentionID {
			return &v, nil
		}
	}

	return nil, fmt.Errorf("RetentionID=%d not found!", retentionID)
}

func (s *DummyRetentionService) List(cursor uint64, limit uint64) ([]Retention, error) {
	if limit == 0 && cursor == 0 {
		return allRetention, nil
	}

	retensionsLen := uint64(len(allRetention))
	maxUint64 := ^uint64(0)
	if cursor > retensionsLen || maxUint64-limit < cursor || cursor+limit > retensionsLen {
		return nil, fmt.Errorf("Index must be lower than %v", retensionsLen)
	}
	return allRetention[cursor : cursor+limit], nil
}

func (s *DummyRetentionService) Create(r Retention) (uint64, error) {
	for i, v := range allRetention {
		if v.RetentionID == r.RetentionID {
			//FIXME: sign to unsign convertion, there may be problems
			return uint64(i), fmt.Errorf("Entry with RetentionID=%d already exist!", r.RetentionID)
		}
	}

	// append at end and sort
	allRetention = append(allRetention, r)
	sort.SliceStable(allRetention, func(i, j int) bool {
		return allRetention[i].RetentionID < allRetention[j].RetentionID
	})

	for i, v := range allRetention {
		if v.RetentionID == r.RetentionID {
			//FIXME: sign to unsign convertion, there may be problems
			return uint64(i), nil
		}
	}
	return 0, fmt.Errorf("Empty entry with RetentionID=%d after insert!", r.RetentionID)
}

func (s *DummyRetentionService) Update(retentionID uint64, retention Retention) error {
	for i, v := range allRetention {
		if v.RetentionID == retention.RetentionID {
			allRetention[i] = retention
			return nil
		}
	}

	return fmt.Errorf("RetentionID=%d not found!", retentionID)
}

func (s *DummyRetentionService) Remove(retentionID uint64) (bool, error) {

	for i, v := range allRetention {
		if v.RetentionID == retentionID {
			allRetention = append(allRetention[:i], allRetention[i+1:]...)
			return true, nil
		}
	}

	return false, fmt.Errorf("RetentionID=%d not found!", retentionID)
}
