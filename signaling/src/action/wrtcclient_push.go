package action

import (
	"fmt"
	"net/http"
	"signaling/src/framework"
)

type wrtcClientPushAction struct{}

func NewWRTCClientPushAction() *wrtcClientPushAction {
	return &wrtcClientPushAction{}
}

func writeHtmlErrorResponse(w http.ResponseWriter, status int, err string) {
	w.WriteHeader(status)
	w.Write([]byte(err))
}

func (*wrtcClientPushAction) Execute(w http.ResponseWriter, cr *framework.ComRequest) {
	fmt.Println("hello wrtcclient push action")
}
