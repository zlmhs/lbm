package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)
func MatchNotice(c *gin.Context) {
	user := c.Param("user")
	c.HTML(http.StatusOK, "matchNotice.html", gin.H{ "USERID": user, "STATIC_FILE_HOST": m_strStaticHost, "MAIN_HOST": m_strMainHost, "EXAM_HOST": m_strExamHost})
}
func MatchAbstract(c *gin.Context) {
	user := c.Param("user")
	c.HTML(http.StatusOK, "matchAbstract.html", gin.H{ "USERID": user, "STATIC_FILE_HOST": m_strStaticHost, "MAIN_HOST": m_strMainHost, "EXAM_HOST": m_strExamHost})
}
func MatchCourseAbstract(c *gin.Context) {
	user := c.Param("user")
	c.HTML(http.StatusOK, "matchCourseAbstract.html", gin.H{ "USERID": user, "STATIC_FILE_HOST": m_strStaticHost, "MAIN_HOST": m_strMainHost, "EXAM_HOST": m_strExamHost})
}
func NewShare(c *gin.Context) {
	user := c.Param("user")
	c.HTML(http.StatusOK, "newShare.html", gin.H{ "USERID": user, "STATIC_FILE_HOST": m_strStaticHost, "MAIN_HOST": m_strMainHost, "EXAM_HOST": m_strExamHost})
}
