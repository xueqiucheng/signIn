package handler

import (
	"net/http"
	"net/http/httputil"
	"regexp"

	"github.com/Azure/go-ntlmssp"
	"github.com/gin-gonic/gin"
)

func Index(context *gin.Context) {

	user := context.Query("a")
	password := context.Query("b")
	url := "http://erp.careerintlinc.com/sitepages/aportal/default.aspx"

	client := &http.Client{
		Transport: ntlmssp.Negotiator{
			RoundTripper: &http.Transport{},
		},
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)

	request.SetBasicAuth(user, password)

	resp, err := client.Do(request)

	if err != nil {
		defer func() {
			if r := recover(); r != nil {
				context.String(http.StatusOK, "login failed[code 1]")
			}
		}()
	}
	defer resp.Body.Close()

	res, err := httputil.DumpResponse(resp, true)

	if err != nil {
		defer func() {
			if r := recover(); r != nil {
				context.String(http.StatusOK, "login failed[code 2]")
			}
		}()
	}

	if len(res) > 0 {
		pattern := `class="status">(.*?)</div>`
		reg := regexp.MustCompile(pattern)
		params := reg.FindStringSubmatch(string(res))
		if params != nil {
			context.String(http.StatusOK, params[1])
		} else {
			context.String(http.StatusOK, "login failed[code 3]")
		}
	} else {
		context.String(http.StatusOK, "login failed[code 4]")
	}

}
