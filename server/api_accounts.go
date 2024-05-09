package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/grokify/go-saviyntecf"
	"github.com/grokify/mogo/net/http/httputilmore"
	"github.com/grokify/mogo/pointer"
)

// PostApiV1Accounts returns the list of accounts
// (POST /api/v1/accounts)
func (a *ECFAPI) PostApiV1Accounts(w http.ResponseWriter, r *http.Request, params saviyntecf.PostApiV1AccountsParams) {
	body, resp, err := a.postApiV1AccountsProc(params)

	if err != nil {
		w.Header().Add(httputilmore.HeaderContentType, httputilmore.ContentTypeTextPlainUtf8)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if resp != nil && resp.StatusCode > 299 {
		w.WriteHeader(resp.StatusCode)
	} else {
		bytes, err := json.Marshal(body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Add(httputilmore.HeaderContentType, httputilmore.ContentTypeAppJSONUtf8)
		_, err = w.Write(bytes)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// w.WriteHeader(http.StatusOK)
	}
}

func (a *ECFAPI) postApiV1AccountsProc(params saviyntecf.PostApiV1AccountsParams) (saviyntecf.Accounts, *http.Response, error) {
	return saviyntecf.Accounts{
		Count:      pointer.Pointer(0),
		Offset:     pointer.Pointer(strconv.Itoa(pointer.Dereference(params.Offset))),
		TotalCount: pointer.Pointer(10),
	}, nil, nil
}
