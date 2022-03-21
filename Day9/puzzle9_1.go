package day9

import (
	"math"
)

func computeRiskLevel(c, l, r, u, d int) int {
	if c < l && c < r && c < u && c < d {
		return 1 + c
	}
	return 0
}

func Puzzle9_1(heightmap [][]int) int {
	riskLevel := 0
	for i := 1; i < len(heightmap)-1; i++ {
		for j := 1; j < len(heightmap[i])-1; j++ {
			riskLevel += computeRiskLevel(heightmap[i][j], heightmap[i][j-1], heightmap[i][j+1], heightmap[i-1][j], heightmap[i+1][j])
		}
	}

	for j := 0; j < len(heightmap[0]); j++ {
		i := len(heightmap) - 1
		if j == 0 {
			riskLevel += computeRiskLevel(heightmap[0][j], math.MaxInt, heightmap[0][j+1], math.MaxInt, heightmap[1][j])
			riskLevel += computeRiskLevel(heightmap[i][j], math.MaxInt, heightmap[i][j+1], heightmap[i-1][j], math.MaxInt)
		} else if j == len(heightmap[0])-1 {
			riskLevel += computeRiskLevel(heightmap[0][j], heightmap[0][j-1], math.MaxInt, math.MaxInt, heightmap[1][j])
			riskLevel += computeRiskLevel(heightmap[i][j], heightmap[i][j-1], math.MaxInt, heightmap[i-1][j], math.MaxInt)
		} else {
			riskLevel += computeRiskLevel(heightmap[0][j], heightmap[0][j-1], heightmap[0][j+1], math.MaxInt, heightmap[1][j])
			riskLevel += computeRiskLevel(heightmap[i][j], heightmap[i][j-1], heightmap[i][j+1], heightmap[i-1][j], math.MaxInt)
		}
	}
	return riskLevel
}
