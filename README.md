# Usererror

Simple package for returning errors when the user makes a mistake.

## Example usage for setting http response
 
```go
import (
	"net/http"

	"github.com/knoebber/usererror"
)

func createUser(username string) error {
	if userExists(username) {
		return usererror.Format("username %q already exists", username)
	}

	return createUser(username)
}

func handleError(w http.ResponseWriter, err error) {
	if uErr := usererror.Convert(err); uErr != nil {
		// Set status 200 and a message in response body.
		setMessage(w, uErr.Message)
	} else {
		// Set status 500.
		setInternalError(w, err)
	}
}
```
