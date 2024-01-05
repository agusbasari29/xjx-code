package response

type ResponseCredential struct {
	Token     string `json:"token"`
	UserID    uint   `json:"user_id"`
	Issuer    string `json:"issuer"`
	IssuedAt  int64  `json:"issued_at"`
	ExpiresAt int64  `json:"expired_at"`
}
