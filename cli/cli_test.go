package cli

import "testing"

func TestRun(t *testing.T) {

	t.Run("run with no args should give error", func(t *testing.T) {
		err := NewCli().Run()
		if err == nil {
			t.Error("expected error but found none")
		}
	})
}
