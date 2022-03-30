package logger

type Factory struct {
	contextID uint32
}

func NewFactory(contextID uint32) *Factory {
	return &Factory{contextID: contextID}
}

func (f *Factory) GetLogger(contextName string) Interface {
	return &contextualized{
		contextID:   f.contextID,
		contextName: contextName,
	}
}
