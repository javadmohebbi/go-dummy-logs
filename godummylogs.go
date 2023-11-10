package godummylogs

type DummyLogs struct {
	Config Config
}

func New(c string) *DummyLogs {
	return &DummyLogs{
		Config: ParseConfig(c),
	}
}
