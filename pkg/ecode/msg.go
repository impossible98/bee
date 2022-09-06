package ecode

var msg = map[int]string{
	Success:              "ok",
	InvalidParams:        "invalid params",
	ErrorCreateFavourite: "failed to create user",
	ErrorGetFavourite:    "failed to get favourite",
	ErrorGetFavouriteList:   "failed to get favourite list",
}

func GetMsg(code int) string {
	return msg[code]
}
