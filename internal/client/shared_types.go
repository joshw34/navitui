package client

type response struct {
	SubResp subResp `json:"subsonic-response"`
}

type subResp struct {
	Status   string        `json:"status"`
	SRArtist subRespArtist `json:"artists"`
	Error    subRespError  `json:"error"`
}

type subRespError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
