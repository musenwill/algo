package src

import "github.com/musenwill/algo/cmd"

func init() {
	cmd.RegisterCommand("bubbleSort", bubbleSort)
	cmd.RegisterCommand("directSort", directSort)
	cmd.RegisterCommand("insertSort", insertSort)
	cmd.RegisterCommand("quickSort", quickSort)
	cmd.RegisterCommand("binarySearch", binarySearch)
}
