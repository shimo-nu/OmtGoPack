package src

import (
	"github.com/gin-contrib/sessions"
	"github.com/google/uuid"
)

// func SessionStart(c *gin.Context) {
func SessionStart() {
	session := session.Default(c)
	uuidObj, _ := uuid.NewUUID()

	fmt.Println(uuidObj.String())
}

func SessionCheck() bool {

	return
}
