/*
 * @Author: Liu Yuchen
 * @Date: 2021-05-08 06:56:50
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-05-08 08:49:55
 * @Description:
 * @FilePath: /spike_system/routine/mq.go
 * @GitHub: https://github.com/liuyuchen777
 */
package routine

import (
	"fmt"
	"log"
	"spike_system/model"
)

const maxMessageNum = 20000

var BuyRecordChannel = make(chan model.RecordMsg, maxMessageNum)
var UpdateLeftChannel = make(chan model.ProductLeft, maxMessageNum)

func BuyRecordConsumer() {
	// read record msg from channel and update records
	for {
		msg := <-BuyRecordChannel
		log.Println("Got a buy record: ", msg)
		// update database
		username := msg.Username
		productname := msg.ProductName
		err := model.CreateRecord(username, productname)
		if err != nil {
			fmt.Println("Create New Record Fail!")
		} else {
			fmt.Println("Create New Record Success!")
		}
	}
}

func UpdateConsumer() {
	// read update msg from channel and update products
	for {
		msg := <-UpdateLeftChannel
		log.Println("Left nedd decreae: ", msg)
		// update database
		left := msg.Left
		productname := msg.Name
		err := model.UpdateLeft(productname, left)
		if err != nil {
			fmt.Println("Update DB Fail!")
		} else {
			fmt.Println("Update DB Success!")
		}
	}
}
