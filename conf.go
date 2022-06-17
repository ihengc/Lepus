package Lepus

/********************************************************
* @author: Ihc
* @date: 2022/6/17 0017 11:49
* @version: 1.0
* @description:
*********************************************************/

type Option func(conf *ApplicationConf)

type ApplicationConf struct {
	Name    string
	Host    string
	Port    int
	AppMode AppMode
}
