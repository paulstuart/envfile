package envfile

import (
	"os"
	"testing"
)

func TestEnvFile(t *testing.T) {
	Debug = testing.Verbose()
	const filename = "testdata/test.env"
	err := EnvFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	const (
		want = "1234uNme"
		name = "MY_TEST_ENV_01"
	)
	have := os.Getenv(name)
	if want != have {
		t.Errorf("want: %q -- have %q", want, have)
	}
}
