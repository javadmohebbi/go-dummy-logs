package godummylogs

type DummyLogs struct {
	Config Config
}

func New(c string) *DummyLogs {
	return &DummyLogs{
		Config: ParseConfig(c),
	}
}

// generate
func (d DummyLogs) Generate() {
	// auth.log
	a := NewAuthLog(d.Config)
	a.Generate()
}
