package cmd

func HandleCmd() error {
	return newRootCmd().Execute()
}
