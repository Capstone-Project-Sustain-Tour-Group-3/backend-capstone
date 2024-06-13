package dto

type TotalDestinationationWithCategory struct {
	Name  string `json:"nama_kategori"`
	Total int    `json:"total"`
}

type MonthlyUser struct {
	Bulan         string `json:"bulan"`
	PenggunaBaru  int    `json:"pengguna_baru"`
	TotalPengguna int    `json:"total_pengguna"`
}
type TotalData struct {
	TotalAdmin int64 `json:"total_admin"`
	TotalUser  int64 `json:"total_user"`
	TotalRute  int64 `json:"total_rute"`
}

type TotalContentVid struct {
	TotalContent   int64 `json:"total_content"`
	TotalDestinasi int64 `json:"total_destinasi"`
}

type DashboardResponse struct {
	TotalData           TotalData                           `json:"total_data"`
	PertumbuhanPengguna []MonthlyUser                       `json:"pertumbuhan_pengguna"`
	TotalKontenVideo    TotalContentVid                     `json:"total_content_vid"`
	KategoriDestinasi   []TotalDestinationationWithCategory `json:"kategori_destinasi"`
	RuteUser            []RouteUser                         `json:"route_user"`
}

type RouteUser struct {
	Username        string  `json:"username"`
	NamaRute        string  `json:"nama_rute"`
	JumlahDestinasi int     `json:"jumlah_destinasi"`
	Biaya           float64 `json:"biaya"`
}
