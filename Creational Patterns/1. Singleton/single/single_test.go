package single

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	inst1 := GetInstance()
	inst2 := GetInstance()
	if inst1 != inst2 {
		t.Error("Instance should be created only once.")
	}
}
