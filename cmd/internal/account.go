package internal

type TestnetResponse struct {
	Msg     string `json:"msg"`
	Keys    Keys   `json:"keys"`
	Account string `json:"account"`
}

type TestnetAccount struct {
	Keys    Keys   `json:"keys"`
	Account string `json:"account"`
}

type Keys struct {
	ActiveKey Key `json:"active_key"`
	OwnerKey  Key `json:"owner_key"`
}

type Key struct {
	Public  string `json:"public"`
	Private string `json:"private"`
}
