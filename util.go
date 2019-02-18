// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-02-18 19:46 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

import (
	"strconv"
	"strings"
)

// usually, there is no need in creating triplets by hand,
// but for test reasons
func newTestTriplet(score float64, left, right string) Triplet {
	return Triplet{Left: strings.Split(left, " "), Right: strings.Split(right, " "), Score: score}
}

// [Tokenize] Tokenize numbers via splitting, i.e. 1924 => (10 91 22 34)
func TokenizeNumbers(number string) string {
	result := []string{}
	for index, number := range number {
		result = append(result, string(number)+strconv.Itoa(index))
	}
	return strings.Join(result, " ")
}