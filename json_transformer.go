// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-02-20 20:10 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

import (
	"encoding/json"
	"io/ioutil"
)

func WriteJSONFile(fpath string, data [][]string) error {
	head := data[0]
	body := data[1:]

	jsondata := []map[string]string{}
	for _, bodyItem := range body {
		row := map[string]string{}
		for index, cell := range bodyItem {
			nameOfCell := head[index]
			row[nameOfCell] = cell
		}
		jsondata = append(jsondata, row)
	}
	jsonAsBytes, err := json.Marshal(jsondata)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fpath, jsonAsBytes, 0644)
}

func ReadJSONFile(fpath string) ([][]string, error) {
	result := [][]string{}
	bytesOfJson, err := ioutil.ReadFile(fpath)
	if err != nil {
		return result, err
	}
	tempResult := []map[string]string{}
	err = json.Unmarshal(bytesOfJson, &tempResult)
	if err != nil {
		return result, err
	}

	// set up headers for the data down the road
	tempHeads := map[string]bool{}
	for _, row := range tempResult {
		for head, _ := range row {
			tempHeads[head] = true
		}
	}

	heads := []string{}
	for head, _ := range tempHeads {
		heads = append(heads, head)
	}

	for _, row := range tempResult {
		resultRow := []string{}
		for _, head := range heads {
			resultRow = append(resultRow, row[head])
		}
		result = append(result, resultRow)
	}
	return result, nil
}