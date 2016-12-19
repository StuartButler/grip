package grip

import (
	"os"
	"os/exec"
	"testing"

	"github.com/tychoish/grip/level"
	"github.com/tychoish/grip/message"
)

func TestConditionalSendFatalExits(t *testing.T) {
	std.UseNativeLogger()
	if os.Getenv("SHOULD_CRASH") == "1" {
		std.EmergencyFatal(message.NewLinesMessage(level.Emergency, "foo"))
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestConditionalSendFatalExits")
	cmd.Env = append(os.Environ(), "SHOULD_CRASH=1")
	err := cmd.Run()
	if err == nil {
		t.Errorf("sendFatal should have exited 1, instead: %s", err.Error())
	}
}
