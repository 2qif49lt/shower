package msg

const (
	CMD_LOGIN_REQ = 1000 + iota
	CMD_LOGIN_RSP

	CMD_ALIVE_REQ
	CMD_ALIVE_RSP

	CMD_LOGOUT_REQ
	CMD_LOGOUT_RSP

	CMD_NEW_CONN_PUSH_REQ
	CMD_NEW_CONN_PUSH_RSP
)

const (
	SUCC = 0
	FAIL = 1
)

const (
	RST_NAME_EXIST = 1000 + iota
	RST_GO_LOGIN
	RST_NO_EXIST
	RST_TRY_LATER
)
