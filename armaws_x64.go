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
	parameters := strings.Split(fmt.Sprintf(C.GoString(input)), ";")
	temp := callWS(parameters)

	// Return a result to Arma
	result := C.CString(temp)
	defer C.free(unsafe.Pointer(result))
	var size = C.strlen(result) + 1
	if size > outputsize {
		size = outputsize
	}
	C.memmove(unsafe.Pointer(output), unsafe.Pointer(result), size)
}

func callWS(parameters []string) string {
	type httpbin struct {
		Origin string `json:"origin"`
		Headers map[string]string `json:"headers"`
		Data map[string]string `json:"json"`
	}

	url := parameters[0]
	parameters = parameters[1:]

	params := make(map[string]string)
	for i := 0; i < len(parameters); i += 2 {
		params[parameters[i]] = parameters [i+1]
	}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(params)

	res, _ := http.Post(url, "application/json; charset=utf-8", b)
	newhttpbin := httpbin{}
	json.NewDecoder(res.Body).Decode(&newhttpbin)

	var result string
	i := 0
	for key, value := range newhttpbin.Data {
		if(i > 0) {
			result = fmt.Sprintf("%v,[%v,%v]",result,key,value)
		} else {
			result = fmt.Sprintf("[[%v,%v]",key,value)
		}
		i++
	}
	result = fmt.Sprintf("%v]",result)
	return result
}

func main() {}