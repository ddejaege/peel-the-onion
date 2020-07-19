package main

import (
    "fmt"
    "encoding/ascii85"
    "io/ioutil"
)

func main() {
    
    layer := "layer0"
    data, err := ioutil.ReadFile(layer + ".in")
    if err != nil {
        fmt.Println("Error reading file", err)
        return
    }

    layerZeroEncodedBytes := []byte(string(data))
    layerOneDecodedBytes := make([]byte, len(layerZeroEncodedBytes))
    
    nLayerOneDecodedBytes, _, err := ascii85.Decode(layerOneDecodedBytes, layerZeroEncodedBytes, true)
    if err != nil {
        fmt.Println("Error decoding input", err)
    }
    
    layerOneDecodedBytes = layerOneDecodedBytes[:nLayerOneDecodedBytes]
    ioutil.WriteFile(layer + ".out", layerOneDecodedBytes, 0664)
}
