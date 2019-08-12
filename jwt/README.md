# jwt

The package `jwt` is used to create and parse JWT tokens.

```go
import (
	"fmt"
	"time"

	"github.com/ricoberger/gocommon/jwt"

	jwtgo "github.com/dgrijalva/jwt-go"
)

func main() {
	token, err := jwt.Create(jwtgo.MapClaims{
		"exp":  time.Now().Add(time.Hour * 24 * 7).Unix(),
		"data": "hello world",
	}, "mysecret")
	if err != nil {
		return
	}

	claims, err := jwt.Parse(token, "mysecret")
	if err != nil {
		return
	}

	fmt.Println(claims)
}
```
