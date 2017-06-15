package controllers

import (
	"xsbPro/common"
)

var (
	// GlobalConf common.IConfigInfo
	m_strMainHost   string
	m_strExamHost   string
	m_strStaticHost string
)

// func InitControllers(examHost string) {
func InitControllers(conf common.IConfigInfo) {
	m_strMainHost   = conf.GetMainHost()
	m_strExamHost   = conf.GetExamHost()
	m_strStaticHost = conf.GetStaticFileServer()
	// GlobalConf = conf
}
