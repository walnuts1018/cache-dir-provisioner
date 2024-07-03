package test

const (
	LabelKey   = "app"
	LabelValue = "cache-dir-provisioner"
)

type testConfig struct {
	IMAGE string
}

func (t *testConfig) envs() []string {
	return testEnvs(t)
}
