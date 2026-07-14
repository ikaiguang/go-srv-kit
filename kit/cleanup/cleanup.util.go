package cleanuputil

type CleanupManager interface {
	// Append 增加
	Append(cleanup func())
	// Cleanup 执行原则：后进先出
	Cleanup()
}

func NewCleanupManager() CleanupManager {
	return &cleanup{}
}

type cleanup struct {
	cleanupList []func()
}

func (s *cleanup) Append(cleanup func()) {
	s.cleanupList = append(s.cleanupList, cleanup)
}

func (s *cleanup) Cleanup() {
	for i := len(s.cleanupList) - 1; i >= 0; i-- {
		if s.cleanupList[i] != nil {
			s.cleanupList[i]()
		}
	}
	s.cleanupList = nil
}

func Merge(cleanupManager CleanupManager, cleanup func(), err error) (CleanupManager, error) {
	if cleanupManager == nil {
		cleanupManager = NewCleanupManager()
	}
	if err != nil {
		cleanupManager.Cleanup()
		return nil, err
	}
	cleanupManager.Append(cleanup)
	return cleanupManager, nil
}
