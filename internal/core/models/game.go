package models

import (
	"time"
	"wynn-member-api/ent"
)

type PrettyCreateAndLoginResponse struct {
	S    string `json:"s"`
	Code int    `json:"code"`
	Data struct {
		PlayerApiId       string `json:"playerApiId"`
		PlayerApiUsername string `json:"playerApiUsername"`
		PlayerUsername    string `json:"playerUsername"`
		TkUuid            string `json:"tkUuid"`
		UriDesktop        string `json:"uriDesktop"`
		UriMobile         string `json:"uriMobile"`
	} `json:"data"`
	ErrMsg string `json:"err_msg,omitempty"`
}

type PrettyDepositResponse struct {
	S    string `json:"s"`
	Code int    `json:"code"`
	Data struct {
		Ref                string  `json:"ref"`
		PlayerApiId        string  `json:"playerApiId"`
		PlayerUsername     string  `json:"playerUsername"`
		PlayerApiUsername  string  `json:"playerApiUsername"`
		AddBalance         float32 `json:"addBalance"`
		AfterMemberBalance float32 `json:"afterMemberBalance"`
		AfterAgentBalance  float32 `json:"afterAgentBalance"`
	} `json:"data"`
	ErrMsg string `json:"err_msg,omitempty"`
}

type PrettyWithdrawResponse struct {
	S    string `json:"s"`
	Code int    `json:"code"`
	Data struct {
		Ref                string  `json:"ref"`
		PlayerApiId        string  `json:"playerApiId"`
		PlayerUsername     string  `json:"playerUsername"`
		PlayerApiUsername  string  `json:"playerApiUsername"`
		AddBalance         float64 `json:"addBalance"`
		AfterMemberBalance float64 `json:"afterMemberBalance"`
		AfterAgentBalance  float64 `json:"afterAgentBalance"`
	} `json:"data"`
	ErrMsg string `json:"err_msg,omitempty"`
}

type PrettyBalanceResponse struct {
	S    string `json:"s"`
	Code int    `json:"code"`
	Data struct {
		PlayerApiId       string  `json:"playerApiId"`
		PlayerUsername    string  `json:"playerUsername"`
		PlayerApiUsername string  `json:"playerApiUsername"`
		Balance           float32 `json:"balance"`
	} `json:"data"`
	ErrMsg string `json:"err_msg,omitempty"`
}

type PGCreateMemberResponse struct {
	S    string `json:"s"`
	Code int    `json:"code"`
	Data struct {
		Id             string   `json:"_id"`
		Username       string   `json:"username"`
		Ancestor       []string `json:"ancestor"`
		Parent         string   `json:"parent"`
		ParentUsername string   `json:"parentUsername"`
		Source         string   `json:"source"`
		Role           string   `json:"role"`
		Profile        struct {
			Name     string `json:"name"`
			Mobile   string `json:"mobile"`
			LineId   string `json:"lineId"`
			Gender   string `json:"gender"`
			Facebook string `json:"facebook"`
			Nickname string `json:"nickname"`
		} `json:"profile"`
		Wallet struct {
			Balance float32 `json:"balance"`
			BuyIn   int     `json:"buyIn"`
			BuyInDs struct {
				NumberDecimal string `json:"$numberDecimal"`
			} `json:"buyInDs"`
			BuyInSpadegaming struct {
				NumberDecimal string `json:"$numberDecimal"`
			} `json:"buyInSpadegaming"`
			BuyInArcadia struct {
				NumberDecimal string `json:"$numberDecimal"`
			} `json:"buyInArcadia"`
			Points struct {
				NumberDecimal string `json:"$numberDecimal"`
			} `json:"points"`
			LastUpdate time.Time `json:"lastUpdate"`
		} `json:"wallet"`
		Pt struct {
			Pg          int `json:"pg"`
			Ds          int `json:"ds"`
			Gamatron    int `json:"gamatron"`
			Ambpoker    int `json:"ambpoker"`
			Mannaplay   int `json:"mannaplay"`
			Cq9         int `json:"cq9"`
			Evoplay     int `json:"evoplay"`
			Ambslot     int `json:"ambslot"`
			Middleware  int `json:"middleware"`
			Spinix      int `json:"spinix"`
			Microgaming struct {
				NumberDecimal string `json:"$numberDecimal"`
			} `json:"microgaming"`
			Arcadia struct {
				NumberDecimal string `json:"$numberDecimal"`
			} `json:"arcadia"`
		} `json:"pt"`
		Level           int  `json:"level"`
		IsSeamless      bool `json:"isSeamless"`
		IsFirstPassword bool `json:"isFirstPassword"`
		Social          struct {
			LineOfficial string `json:"lineOfficial"`
			Google       string `json:"google"`
			Facebook     string `json:"facebook"`
			Twitter      string `json:"twitter"`
		} `json:"social"`
		IsRegisterAngpowFeature bool `json:"isRegisterAngpowFeature"`
		ParentSocial            struct {
			Init         bool   `json:"$init"`
			LineOfficial string `json:"lineOfficial"`
			Google       string `json:"google"`
			Facebook     string `json:"facebook"`
			Twitter      string `json:"twitter"`
		} `json:"parentSocial"`
	} `json:"data"`
	ErrMsg string `json:"err_msg,omitempty"`
}

type PGDepositResponse struct {
	S    string `json:"s"`
	Code int    `json:"code"`
	Data struct {
		Username string `json:"username"`
		Wallet   struct {
			Balance float32 `json:"balance"`
			BuyIn   int     `json:"buyIn"`
			BuyInDs struct {
				NumberDecimal string `json:"$numberDecimal"`
			} `json:"buyInDs"`
			BuyInSpadegaming struct {
				NumberDecimal string `json:"$numberDecimal"`
			} `json:"buyInSpadegaming"`
			BuyInArcadia struct {
				NumberDecimal string `json:"$numberDecimal"`
			} `json:"buyInArcadia"`
			Points struct {
				NumberDecimal string `json:"$numberDecimal"`
			} `json:"points"`
			LastUpdate time.Time `json:"lastUpdate"`
		} `json:"wallet"`
		Balance struct {
			Before float32 `json:"before"`
			After  float32 `json:"after"`
		} `json:"balance"`
		RefId string `json:"refId"`
	} `json:"data"`
	ErrMsg string `json:"err_msg,omitempty"`
}

type PGWithdrawResponse struct {
	S    string `json:"s"`
	Code int    `json:"code"`
	Data struct {
		Username string `json:"username"`
		Wallet   struct {
			Balance float32 `json:"balance"`
			BuyIn   float32 `json:"buyIn"`
			BuyInDs struct {
				NumberDecimal string `json:"$numberDecimal"`
			} `json:"buyInDs"`
			BuyInSpadegaming struct {
				NumberDecimal string `json:"$numberDecimal"`
			} `json:"buyInSpadegaming"`
			BuyInArcadia struct {
				NumberDecimal string `json:"$numberDecimal"`
			} `json:"buyInArcadia"`
			Points struct {
				NumberDecimal string `json:"$numberDecimal"`
			} `json:"points"`
			LastUpdate time.Time `json:"lastUpdate"`
		} `json:"wallet"`
		Balance struct {
			Before int `json:"before"`
			After  int `json:"after"`
		} `json:"balance"`
		RefId string `json:"refId"`
	} `json:"data"`
	ErrMsg string `json:"err_msg,omitempty"`
}

type PGBalanceResponse struct {
	S    string `json:"s"`
	Code int    `json:"code"`
	Data struct {
		Balance     float32 `json:"balance"`
		Outstanding float32 `json:"outstanding"`
	} `json:"data"`
	ErrMsg string `json:"err_msg,omitempty"`
}

type PGLoginResponse struct {
	S    string `json:"s"`
	Code int    `json:"code"`
	Data struct {
		Url string `json:"url"`
	} `json:"data"`
	ErrMsg string `json:"err_msg,omitempty"`
}

type SAGameLoginResponse struct {
	S    string `json:"s"`
	Code int    `json:"code"`
	Data struct {
		Data struct {
			ErrorMsgId  string `json:"ErrorMsgId"`
			ErrorMsg    string `json:"ErrorMsg"`
			Token       string `json:"Token"`
			DisplayName string `json:"DisplayName"`
			LobbyCode   string `json:"lobby_code"`
		} `json:"data"`
		Url string `json:"url"`
	} `json:"data"`
	ErrMsg string `json:"err_msg,omitempty"`
}

type SAGameRegisterResponse struct {
	S      string `json:"s"`
	Code   int    `json:"code"`
	Data   string `json:"data"`
	ErrMsg string `json:"err_msg,omitempty"`
}

type SAGameProfileResponse struct {
	S    string `json:"s"`
	Code int    `json:"code"`
	Data struct {
		ErrorMsgId     string `json:"ErrorMsgId"`
		ErrorMsg       string `json:"ErrorMsg"`
		IsSuccess      string `json:"IsSuccess"`
		Username       string `json:"Username"`
		Balance        string `json:"Balance"`
		Online         string `json:"Online"`
		Betted         string `json:"Betted"`
		BettedAmount   string `json:"BettedAmount"`
		MaxBalance     string `json:"MaxBalance"`
		MaxWinning     string `json:"MaxWinning"`
		WithholdAmount string `json:"WithholdAmount"`
	} `json:"data"`
	ErrMsg string `json:"err_msg,omitempty"`
}

type SAGameDepositResponse struct {
	S    string `json:"s"`
	Code int    `json:"code"`
	Data struct {
		ErrorMsgId   string `json:"ErrorMsgId"`
		ErrorMsg     string `json:"ErrorMsg"`
		Username     string `json:"Username"`
		Balance      string `json:"Balance"`
		CreditAmount string `json:"CreditAmount"`
		OrderId      string `json:"OrderId"`
	} `json:"data"`
	ErrMsg string `json:"err_msg,omitempty"`
}

type SAGameWithdrawResponse struct {
	S    string `json:"s"`
	Code int    `json:"code"`
	Data struct {
		ErrorMsgId  string `json:"ErrorMsgId"`
		ErrorMsg    string `json:"ErrorMsg"`
		Username    string `json:"Username"`
		Balance     string `json:"Balance"`
		DebitAmount string `json:"DebitAmount"`
		OrderId     string `json:"OrderId"`
	} `json:"data"`
	ErrMsg string `json:"err_msg,omitempty"`
}

type GameAccountListResponse struct {
	Id        int       `json:"id"`
	Uuid      string    `json:"uuid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Edges     struct {
		Game   *ent.Game     `json:"game"`
		Detail []interface{} `json:"detail,omitempty"`
	} `json:"edges"`
}
