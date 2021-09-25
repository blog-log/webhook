package config

import "os"

func SetEnv(vars map[string]string) error {
	for k, v := range vars {
		if err := os.Setenv(k, v); err != nil {
			return err
		}
	}
	return nil
}

func UnsetEnv(vars map[string]string) error {
	for k := range vars {
		if err := os.Unsetenv(k); err != nil {
			return err
		}
	}
	return nil
}
