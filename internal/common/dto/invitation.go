package dto

type GetInvitationTodayInfoResponse struct {
	List []struct {
		Account string `json:"account"`
	} `json:"list"`
}
