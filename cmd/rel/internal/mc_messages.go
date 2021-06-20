package internal

import "fmt"

func MsgBadRootEnvFile() error {
	return fmt.Errorf(
		"problem with env root file. Set correct value to environment with key %s",
		RootEnvFileKey)
}

func MsgWithUserText(message string) error {
	return fmt.Errorf("user error. %s", message)
}
