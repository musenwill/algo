package src

import (
	"fmt"

	"github.com/urfave/cli"
)

func sentenceReverse(c *cli.Context) error {
	sentenceReverseCase("how are you")
	sentenceReverseCase(" It's  a  car ")
	return nil
}

func sentenceReverseCase(sentence string) {
	fmt.Printf("origin : %s\n", sentence)
	fmt.Printf("reverse: %s\n", sentenceReverseImp(sentence))
}

func sentenceReverseImp(sentence string) string {
	runes := []rune(sentence)

	{
		firstIdx := 0
		lastIdx := len(runes) - 1
		for firstIdx < lastIdx {
			runes[firstIdx], runes[lastIdx] = runes[lastIdx], runes[firstIdx]
			firstIdx++
			lastIdx--
		}
	}

	for wordStartIdx := 0; wordStartIdx < len(runes); wordStartIdx++ {
		if runes[wordStartIdx] == ' ' {
			continue
		}
		wordEndIdx := wordStartIdx
		for ; wordEndIdx < len(runes); wordEndIdx++ {
			if runes[wordEndIdx] == ' ' {
				break
			}
		}
		wordEndIdx--
		wordReverse(runes, wordStartIdx, wordEndIdx)
		wordStartIdx = wordEndIdx
	}

	return string(runes)
}

func wordReverse(sentence []rune, firstIdx, lastIdx int) {
	for firstIdx < lastIdx {
		sentence[firstIdx], sentence[lastIdx] = sentence[lastIdx], sentence[firstIdx]
		firstIdx++
		lastIdx--
	}
}
