/*
 * Copyright (c) 2018. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package main

import (
	_"fmt"
	"log"
	"net/url"
	"reflect"
)

type RegistrationInfo struct {
	UserId string `json:"userid"`
	Passwrod string `json:"password"`
}

func checkParams (params url.Values) bool{
	isOk := true
	if len(params) != 2 {
		return false

	}

	keys := reflect.ValueOf(params).MapKeys()

	k1, k2 := keys[0].String(), keys[1].String()

	if k1 != "userid" || k2 != "password"{
		isOk = false
	}

	if !isOk{
		//TODO some bug, false when params is normal
		log.Println("what???")
	}

	return isOk
}

