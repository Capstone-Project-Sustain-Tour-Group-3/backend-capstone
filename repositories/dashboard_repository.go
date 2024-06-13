package repositories

import (
	"capstone/dto"
	"capstone/entities"
	"gorm.io/gorm"
)

type DashboardRepository interface {
	GetTotalAdmin() (int64, error)
	GetTotalUser() (int64, error)
	GetTotalRoute() (int64, error)
	GetTotalDestination() (int64, error)
	GetTotalVidContent() (int64, error)
	GetTotalDestinationWithCategory() ([]dto.TotalDestinationationWithCategory, error)
	GetMonthlyUsers() ([]dto.MonthlyUser, error)
	GetRouteUser() ([]dto.RouteUser, error)
}

type dashboardRepository struct {
	db *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) *dashboardRepository {
	return &dashboardRepository{db: db}
}

func (r *dashboardRepository) GetTotalAdmin() (int64, error) {
	var total int64
	if err := r.db.Model(&entities.Admin{}).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func (r *dashboardRepository) GetTotalUser() (int64, error) {
	var total int64
	if err := r.db.Model(&entities.User{}).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func (r *dashboardRepository) GetTotalRoute() (int64, error) {
	var total int64
	if err := r.db.Model(&entities.Route{}).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func (r *dashboardRepository) GetTotalDestination() (int64, error) {
	var total int64
	if err := r.db.Model(&entities.Destination{}).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func (r *dashboardRepository) GetTotalVidContent() (int64, error) {
	var total int64
	if err := r.db.Model(&entities.DestinationMedia{}).Where("type = ?", "video").Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func (r *dashboardRepository) GetTotalDestinationWithCategory() ([]dto.TotalDestinationationWithCategory, error) {
	var totalDestinationationWithCategory []dto.TotalDestinationationWithCategory
	if err := r.db.Table("destinations").
		Select("categories.name, COUNT(destinations.id) as total").
		Joins("JOIN categories ON destinations.category_id = categories.id").
		Group("categories.id").
		Scan(&totalDestinationationWithCategory).Error; err != nil {
		return nil, err
	}
	return totalDestinationationWithCategory, nil
}
func (r *dashboardRepository) GetMonthlyUsers() ([]dto.MonthlyUser, error) {
	var monthlyUsers []dto.MonthlyUser
	query := `
        SELECT
            CASE month
                WHEN '01' THEN 'Jan'
                WHEN '02' THEN 'Feb'
                WHEN '03' THEN 'Mar'
                WHEN '04' THEN 'Apr'
                WHEN '05' THEN 'Mei'
                WHEN '06' THEN 'Jun'
                WHEN '07' THEN 'Jul'
                WHEN '08' THEN 'Ags'
                WHEN '09' THEN 'Sep'
                WHEN '10' THEN 'Okt'
                WHEN '11' THEN 'Nov'
                WHEN '12' THEN 'Des'
            END AS bulan,
            pengguna_baru,
            (
                SELECT COUNT(u2.id)
                FROM users u2
                WHERE DATE_FORMAT(u2.created_at, '%m') <= month
            ) AS total_pengguna
        FROM (
            SELECT
                DATE_FORMAT(created_at, '%m') AS month,
                COUNT(id) AS pengguna_baru
            FROM
                users
            GROUP BY
                month
        ) AS bulan
        ORDER BY
            month;
    `

	if err := r.db.Raw(query).Scan(&monthlyUsers).Error; err != nil {
		return nil, err
	}

	return monthlyUsers, nil
}

func (r *dashboardRepository) GetRouteUser() ([]dto.RouteUser, error) {
	var routeUsers []dto.RouteUser
	query := `
		SELECT u.username ,r.name as nama_rute,COUNT(rd.id) AS jumlah_destinasi, r.price as biaya
		FROM routes r
		INNER JOIN route_details rd ON r.id = rd.route_id
		INNER JOIN users u ON r.user_id = u.id
		GROUP BY r.id
		ORDER BY r.created_at DESC
		LIMIT 5;
	`

	if err := r.db.Raw(query).Scan(&routeUsers).Error; err != nil {
		return nil, err
	}

	return routeUsers, nil
}
