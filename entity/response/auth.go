package response

type Auth struct {
	Id      uint   `json:"id,omitempty"`
	Name 	string `json:"name,omitempty"`
	Email   string `json:"email,omitempty"`
	Role    string `json:"role,omitempty"`
}

type UserGoogleInfo struct {
	Id				string `json:"id"`
	Email			string `json:"email"`
	VerifiedEmail	string `json:"verifiedEmail"`
	Name 			string `json:"name"`
}

type AuthToken struct {
	Name 	string `json:"name"`
	Email   string `json:"email"`
	Token	string `json:"token"`
}

// // UserGoogleInfo adalah struktur data untuk menyimpan informasi pengguna dari Google.
// type UserGoogleInfo struct {
// 	ID            string `json:"id"`
// 	Email         string `json:"email"`
// 	Name          string `json:"name"`
// 	Picture       string `json:"picture"`
// 	// Tambahan informasi pengguna lainnya sesuai kebutuhan.
// }

// NewUserGoogleInfo membuat instance baru dari UserGoogleInfo.
// func NewUserGoogleInfo(id, email, name, picture string) *UserGoogleInfo {
// 	return &UserGoogleInfo{
// 		ID:      id,
// 		Email:   email,
// 		Name:    name,
// 		Picture: picture,
// 	}
// }