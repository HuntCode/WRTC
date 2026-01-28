package main

import (
	"signaling/src/action"
	"signaling/src/framework"
)

func init() {
	framework.GActionRouter["/wrtcclient/push"] = action.NewWRTCClientPushAction()
}
