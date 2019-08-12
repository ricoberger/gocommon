# middleware

The package `middleware` implements various middleware handler for [httprouter](https://github.com/julienschmidt/httprouter). The package contains error, cors, log and authentication handlers.

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ricoberger/gocommon/middleware"

	"github.com/julienschmidt/httprouter"
)

func getUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) *middleware.Error {
	user, err := db.GetUsers()
	if err != nil {
		return middleware.Errorf(err, http.StatusInternalServerError, "Internal Server Error")
	}

	return middleware.WriteJSON(w, r, nil)
}

func main() {
	var router = httprouter.New()
	router.GET("/users", middleware.Log(middleware.Handle(getUsers)))

	log.Fatalln(http.ListenAndServe(":8080", router))
}
```
