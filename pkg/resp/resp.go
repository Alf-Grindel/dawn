package resp

import (
	"github.com/alf-grindel/dawn/pkg/errno"
	"github.com/alf-grindel/dawn/pkg/utils"
	"net/http"
)

type response struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	DATA    Data   `json:"data"`
}

type Data map[string]interface{}

func WriteJson(rw http.ResponseWriter, e errno.Errno, msg Data) error {
	resp := &response{
		Code:    e.Code,
		Message: e.Message,
		DATA:    msg,
	}
	utils.ToJSON(resp, rw)
	return nil
}
