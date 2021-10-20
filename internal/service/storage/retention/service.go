package retention

//"errors"

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

// это get
func (s *DummyRetentionService) Describe(retentionID uint64) (*Retention, error) {
	/*for _, v := range allRetention {
	    if v.retentionID == retentionID {
	      return &v, nil
	    }
	  }
	  return allRetention[retentionID], errors.Printf("retentionID %d not found!", retentionID)*/
	return nil, nil
}

// это list
func (s *DummyRetentionService) List(cursor uint64, limit uint64) ([]Retention, error) {
	//return allRetention, nil
	return nil, nil
}

// это new
func (s *DummyRetentionService) Create(r Retention) (uint64, error) {
	/*for _, v := range allRetention {
	    if v.retentionID == r.retentionID {
	      return nil, errors.Printf("retentionID %d already exist!", retentionID)
	    }
	  }

	  append(allRetention, r)
	  sort.SliceStable(allRetention, func(i, j int) bool {
	    return allRetention[i].retentionID < allRetention[j].retentionID
	  })
	*/
	return 0, nil
}

// это edit
func (s *DummyRetentionService) Update(retentionID uint64, retention Retention) error {
	/*
	  for i, v := range allRetention {
	    if v.retentionID == retention.retentionID {
	      allRetention[i] = retention
	      return nil
	    }
	  }
	  // если элемента нет, считаем, что надо создать, заного отработает поиск, но, пока поф, это заглушка
	  _, err := s.New(retention)
	  return err*/
	return nil
}

// это delete
func (s *DummyRetentionService) Remove(retentionID uint64) (bool, error) {
	/*
	  for i, v := range allRetention {
	    if v.retentionID == retentionID {
	        allRetention = append(slice[:i], slice[i+1:]...)
	      return nil
	    }
	  }
	  // Элемента нет, удалять нечего, всё ок?
	  return nil*/
	return true, nil
}
