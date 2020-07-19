package main

import (
	"fmt"
	"encoding/ascii85"
	"io/ioutil"
	"math/bits"
)

func main() {

	layer := "layer1"
	data, err := ioutil.ReadFile(layer + ".in")
	if err != nil {
		fmt.Println("Error reading file, err")
		return
	}

	layerOneEncodedBytes := []byte(string(data))
	layerTwoDecodedBytes := make([]byte, len(layerOneEncodedBytes))

	nLayerTwoDecodedBytes, _, err := ascii85.Decode(layerTwoDecodedBytes, layerOneEncodedBytes, true)
	if err != nil {
		fmt.Println("Error decoding input", err)
	}

	layerTwoDecodedBytes = layerTwoDecodedBytes[:nLayerTwoDecodedBytes]
	for i, v := range layerTwoDecodedBytes {
		var value uint8 = uint8(v)
		// "01010101" will flip every other bit, which = 85
		var flipper uint8 = 85
		var flipped uint8 = value ^ flipper
		// rotate them "left" -1 (right 1)
		var rotatedAndFlipped uint8 = bits.RotateLeft8(flipped, -1)
		layerTwoDecodedBytes[i] = rotatedAndFlipped
	}

	ioutil.WriteFile(layer + ".out", layerTwoDecodedBytes, 0664)

}
