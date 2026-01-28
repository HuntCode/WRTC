package action

import (
	"fmt"
	"net/http"
)

type wrtcClientPushAction struct{}

func NewWRTCClientPushAction() *wrtcClientPushAction {
	return &wrtcClientPushAction{}
}

func (*wrtcClientPushAction) Execute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello wrtcclient push action")
}
