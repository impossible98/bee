package ecode

var msg = map[int]string{
	Success:               "ok",
	InvalidParams:         "invalid params",
	ErrorCreateFavourite:  "failed to create user",
	ErrorGetFavourite:     "failed to get favourite",
	ErrorGetFavouriteList: "failed to get favourite list",
	ErrorDeleteFavourite:  "failed to delele favourite",
}

func GetMsg(code int) string {
	return msg[code]
}
