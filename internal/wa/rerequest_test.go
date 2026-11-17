package wa

import "testing"

func TestRerequestFromPhone(t *testing.T) {
	for v, want := range map[string]bool{"": true, "1": true, "true": true, "0": false, "false": false, " Off ": false, "no": false} {
		t.Setenv("WACLI_REREQUEST_FROM_PHONE", v)
		if got := RerequestFromPhone(); got != want {
			t.Errorf("%q: got %v, want %v", v, got, want)
		}
	}
}
