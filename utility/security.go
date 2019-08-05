package utility

//import "fmt"
import "net/http"

func AuthenticateRequest(header http.Header) bool {
	//fmt.Println("authentciation:", header.Get("Authentication"))
	if header.Get("Authentication") == "" || header.Get("Authentication") != "hasToken" {
		return false
	}
	return true
}
