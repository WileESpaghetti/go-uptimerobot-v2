package cmd

import (
	"testing"
)

func Test_ExecuteCommand(t *testing.T) {
	//accountKey := "uxxxxxx-xxxxxxxxxxxxxxxxxxxxxxxx"
	//monitorKey := "mxxxxxxxxx-xxxxxxxxxxxxxxxxxxxxxxxx"

	t.Run("should accept api-key flag", func(t *testing.T) {
		flag := FlagApiKey

		wantKey := "test_api_key"

		rootCmd.SetArgs([]string{"--" + flag, wantKey})

		apiKeyFlag := rootCmd.PersistentFlags().Lookup(flag)
		if apiKeyFlag == nil {
			t.Errorf("flag --%s missing", flag)
		}

		// TODO check Client uses the set API key
	})
}
