// db project main.go
package main

import (
	"github.com/gin-gonic/gin"
)

type myForm struct {
	Colors []string `form:"colors[]"`
}

func (formHandler (c *gin.Context) {
	var fakeForm myForm
	c.ShouldBind(&fakeForm)
	c.JSON(200, gin.H{"color": fakeForm.Colors})
})
func main() {
	formHandler()
}
