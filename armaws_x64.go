package main

/*
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
*/
import "C"

import (
	"fmt"
	"strings"
	"unsafe"
	"bytes"
	"encoding/json"
	"net/http"
)

//export RVExtensionVersion
func RVExtensionVersion(output *C.char, outputsize C.size_t) {
	result := C.CString("Version 1.0")
	defer C.free(unsafe.Pointer(result))
	var size = C.strlen(result) + 1
	if size > outputsize {
		size = outputsize
	}
	C.memmove(unsafe.Pointer(output), unsafe.Pointer(result), size)
}

//export RVExtension
func RVExtension(output *C.char, outputsize C.size_t, input *C.char) {
	var temp string
	parameters := strings.Split(fmt.Sprintf(C.GoString(input)), ";")
	function := parameters[0]
	parameters = parameters[1:]

	switch function {
		case "getVersion" : {
			temp = "0.2"
		}
		case "head" : {
			temp = callHead(parameters)
		}
		case "post" : {
			temp = callPost(parameters)
		}
		default: {
			temp = callPost(parameters)
		}
	}
	
	// Return a result to Arma
	result := C.CString(temp)
	defer C.free(unsafe.Pointer(result))
	var size = C.strlen(result) + 1
	if size > outputsize {
		size = outputsize
	}
	C.memmove(unsafe.Pointer(output), unsafe.Pointer(result), size)
}

func callHead(parameters []string) string {
	url := parameters[0]
	res, err := http.Head(url)
	if err != nil {
		return fmt.Sprintf("[-1,\"%v\"]", err.Error())
	}
	return fmt.Sprintf("[0,\"%v\"]", result)
}

func callPost(parameters []string) string {
	type httpbin struct {
		Origin string `json:"origin"`
		Headers map[string]string `json:"headers"`
		Data map[string]string `json:"json"`
	}

	newhttpbin := httpbin{}
	result := "[]"

	url := parameters[0]
	parameters = parameters[1:]

	if len(parameters) % 2 != 0 {
		return fmt.Sprintf("[-1,\"%v\"]", "parameters numbers are incorrect")
	}

	params := make(map[string]string)
	for i := 0; i < len(parameters); i += 2 {
		params[parameters[i]] = parameters [i+1]
	}

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(params)
	if err != nil {
		return fmt.Sprintf("[-1,\"%v\"]", err.Error())
	}

	res, err := http.Post(url, "application/json; charset=utf-8", b)
	if err != nil {
		return fmt.Sprintf("[-1,\"%v\"]", err.Error())
	}

	err = json.NewDecoder(res.Body).Decode(&newhttpbin)
	if err != nil {
		return fmt.Sprintf("[-1,\"%v\"]", err.Error())
	}
	
	i := 0
	for key, value := range newhttpbin.Data {
		if(i > 0) {
			result = fmt.Sprintf("%v,[\"%v\",\"%v\"]",result,key,value)
		} else {
			result = fmt.Sprintf("[0,[[\"%v\",\"%v\"]",key,value)
		}
		i++
	}
	result = fmt.Sprintf("%v]]",result)
	return result
}

func main() {}