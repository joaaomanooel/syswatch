package cmd

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	t.Run("Execute runs without error", func(t *testing.T) {
		rootCmd.SetArgs([]string{"version"})
		err := rootCmd.Execute()
		assert.NoError(t, err)
		rootCmd.SetArgs([]string{})
	})

	t.Run("Execute function calls rootCmd.Execute", func(t *testing.T) {
		rootCmd.SetArgs([]string{"version"})
		done := make(chan bool)
		go func() {
			Execute()
			done <- true
		}()
		<-done
		rootCmd.SetArgs([]string{})
	})

	t.Run("Execute with invalid command triggers os.Exit", func(t *testing.T) {
		if os.Getenv("TEST_EXIT") == "1" {
			rootCmd.SetArgs([]string{"invalid"})
			Execute()
			return
		}

		cmd := exec.Command(os.Args[0], "-test.run=TestExecute")
		cmd.Env = append(os.Environ(), "TEST_EXIT=1")
		err := cmd.Run()
		if e, ok := err.(*exec.ExitError); !ok {
			t.Fatalf("expected exit error, got %v", e)
		}
	})
}

func TestMainExit(t *testing.T) {
	t.Run("main exits with error on invalid command", func(t *testing.T) {
		rootCmd.SetArgs([]string{"invalid-command"})
		err := rootCmd.Execute()
		assert.Error(t, err)
		rootCmd.SetArgs([]string{})
	})
}

func TestRootCmd(t *testing.T) {
	t.Run("root command has correct use", func(t *testing.T) {
		assert.Equal(t, "syswatch", rootCmd.Use)
	})

	t.Run("root command has description", func(t *testing.T) {
		assert.Contains(t, rootCmd.Short, "SysWatch")
	})
}

func TestMonitorCmd(t *testing.T) {
	t.Run("monitor command exists", func(t *testing.T) {
		cmd := MonitorCmd()
		assert.NotNil(t, cmd)
		assert.Equal(t, "monitor", cmd.Use)
	})

	t.Run("monitor has interval flag", func(t *testing.T) {
		cmd := MonitorCmd()
		flag := cmd.Flag("interval")
		assert.NotNil(t, flag)
		assert.Equal(t, "2s", flag.DefValue)
	})
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
