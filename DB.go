/*
 * Copyright (c) 2018. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package main

import "fmt"

const (
	DBAddress = "localhost:3306"
)

var (
	DBForGoInfo = DBInfo{
		Login: "root",
		Pass: "root",
		DBName: "DBForGO",
	}
)

type DBInfo struct {
	Login, Pass, DBName string

}


func (dbInfo  DBInfo) GetDataSourceName () string{

	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", dbInfo.Login, dbInfo.Pass,DBAddress,dbInfo.DBName)
}