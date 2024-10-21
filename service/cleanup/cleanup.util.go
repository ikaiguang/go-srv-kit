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
	lastIndex := len(s.cleanupList) - 1
	for lastIndex >= 0 {
		if s.cleanupList[lastIndex] == nil {
			continue
		}
		s.cleanupList[lastIndex]()
		lastIndex--
	}
	s.cleanupList = nil
}
