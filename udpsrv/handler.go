package udpsrv

import (
	"github.com/2qif49lt/logrus"
	"github.com/2qif49lt/shower/msg"
	"github.com/2qif49lt/shower/pack"
	"github.com/2qif49lt/shower/proxys"
)

type MsgHandler struct {
	Pxys *proxys.Proxys
}

func (hder *MsgHandler) Handle(inmsg *InQueueMsg, sender Sender) bool {
	cmd, body, err := pack.Unpack(inmsg.data)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"ip":       inmsg.remoteAddr.String(),
			"time":     inmsg.putTime,
			"data len": len(inmsg.data),
			"error":    err.Error(),
		}).Warnln("unpack fail")
		return false
	}

	switch cmd {
	case msg.CMD_LOGIN_REQ:
		return hder.Pxys.DoCmdLoginReq(cmd, body, sender, inmsg.putTime, inmsg.remoteAddr)
	case msg.CMD_ALIVE_REQ:
		return hder.Pxys.DoCmdAliveReq(cmd, body, sender, inmsg.putTime, inmsg.remoteAddr)
	case msg.CMD_LOGOUT_REQ:
		return hder.Pxys.DoCmdLogoutReq(cmd, body, sender, inmsg.putTime, inmsg.remoteAddr)
	}
	return false
}
