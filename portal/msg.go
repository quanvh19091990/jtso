package portal

type (
	TabAsso struct {
		Shortname string
		Profiles  string
	}

	NewRouter struct {
		Hostname  string `json:"hostname"`
		Shortname string `json:"shortname"`
		Family    string `json:"family"`
	}

	DeletedRouter struct {
		Shortname string `json:"shortname"`
	}

	AddProfile struct {
		Shortname string   `json:"shortname"`
		Profiles  []string `json:"profiles"`
	}

	UpdateDoc struct {
		Profile string `json:"profile"`
	}

	DelProfile struct {
		Shortname string `json:"shortname"`
	}

	Reply struct {
		Status string `json:"status"`
		Msg    string `json:"msg"`
	}

	ReplyDoc struct {
		Status string `json:"status"`
		Img    string `json:"img"`
		Desc   string `json:"desc"`
		Tele   string `json:"tele"`
		Graf   string `json:"graf"`
		Kapa   string `json:"kapa"`
	}

	Credential struct {
		NetconfUser string `json:"netuser"`
		NetconfPwd  string `json:"netpwd"`
		GnmiUser    string `json:"gnmiuser"`
		GnmiPwd     string `json:"gnmipwd"`
		UseTls      string `json:"usetls"`
	}
)
