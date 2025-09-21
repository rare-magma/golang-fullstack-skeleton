package arranger

type EnvironmentMultiArranger struct {
	Arrangers []Arranger
}

func (a *EnvironmentMultiArranger) Arrange() error {
	for _, arranger := range a.Arrangers {
		if err := arranger.Arrange(); err != nil {
			return err
		}
	}
	return nil
}

func (a *EnvironmentMultiArranger) Close() error {
	for _, arranger := range a.Arrangers {
		if err := arranger.Close(); err != nil {
			return err
		}
	}
	return nil
}
