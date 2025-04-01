package todolist

import "testing"

func TestRun(t *testing.T) {

	t.Run("run with no args should give error", func(t *testing.T) {
		err := NewCli().Run()
		assertError(t, err, NoArgsError)
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Error("wanted an error but didn't get one")
		return
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

