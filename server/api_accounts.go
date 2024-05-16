package server

import (
	"encoding/json"
	"net/http"

	"github.com/grokify/go-saviyntecf"
	"github.com/grokify/mogo/net/http/httputilmore"
	"github.com/grokify/mogo/pointer"
)

// PostApiV1Accounts returns the list of accounts
// (POST /api/v1/accounts)
func (a *ECFAPI) PostAccounts(w http.ResponseWriter, r *http.Request, params saviyntecf.PostAccountsParams) {
	body, resp, err := a.postAccountsProc(params)

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
	}
}

func (a *ECFAPI) postAccountsProc(params saviyntecf.PostAccountsParams) (saviyntecf.Accounts, *http.Response, error) {
	accts := saviyntecf.Accounts{
		Pagesize:   pointer.Pointer(pointer.Dereference(params.Pagesize)),
		Offset:     pointer.Pointer(pointer.Dereference(params.Offset)),
		TotalCount: pointer.Pointer(10),
	}
	resp, err := http.Get(AccountsImportCSVURL)
	if err != nil {
		return accts, resp, err
	}
	return accts, resp, nil
}
