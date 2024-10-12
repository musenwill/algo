package src

import "github.com/musenwill/algo/cmd"

func init() {
	cmd.RegisterCommand("bubbleSort", bubbleSort)
	cmd.RegisterCommand("directSort", directSort)
	cmd.RegisterCommand("insertSort", insertSort)
	cmd.RegisterCommand("quickSort", quickSort)
	cmd.RegisterCommand("mergeSort", mergeSort)
	cmd.RegisterCommand("shellSort", shellSort)
	cmd.RegisterCommand("heapSort", heapSort)
	cmd.RegisterCommand("binarySearch", binarySearch)

	cmd.RegisterCommand("hanoi", hanoi)
	cmd.RegisterCommand("gcd", gcd)
	cmd.RegisterCommand("pow", pow)
	cmd.RegisterCommand("topK", topK)
	cmd.RegisterCommand("reverseLink", reverseLink)
	cmd.RegisterCommand("sortLink", sortLink)
	cmd.RegisterCommand("mergeLink", mergeLink)
	cmd.RegisterCommand("swapLink", swapLink)
	cmd.RegisterCommand("circleLink", circleLink)
	cmd.RegisterCommand("lru", lru)
	cmd.RegisterCommand("treeFM", treeFM)
	cmd.RegisterCommand("treeLM", treeLM)
	cmd.RegisterCommand("channel", channel)
}
