package sessions

import (
	"fmt"
    "github.com/gin-contrib/sessions"
    "github.com/gin-gonic/gin"
)

func SetSession(c *gin.Context, key string, value interface{}) error {
    session := sessions.Default(c)
    session.Set(key, value)
    return session.Save()
}

func GetSession(c *gin.Context, key string) (interface{}, error) {
    session := sessions.Default(c)
    value := session.Get(key)
    if value == nil {
        return nil, fmt.Errorf("no session value for key: %s", key)
    }
    return value, nil
}
