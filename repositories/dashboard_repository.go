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
	if err := r.db.Debug().Model(&entities.Admin{}).Where("role", "admin").Count(&total).
		Error; err != nil {
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
        WITH all_months AS (
    SELECT '01' AS month UNION ALL
    SELECT '02' AS month UNION ALL
    SELECT '03' AS month UNION ALL
    SELECT '04' AS month UNION ALL
    SELECT '05' AS month UNION ALL
    SELECT '06' AS month UNION ALL
    SELECT '07' AS month UNION ALL
    SELECT '08' AS month UNION ALL
    SELECT '09' AS month UNION ALL
    SELECT '10' AS month UNION ALL
    SELECT '11' AS month UNION ALL
    SELECT '12' AS month
),
monthly_data AS (
    SELECT
        DATE_FORMAT(created_at, '%m') AS month,
        COUNT(id) AS pengguna_baru
    FROM
        users
    GROUP BY
        month
)
SELECT
    CASE all_months.month
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
    COALESCE(monthly_data.pengguna_baru, 0) AS pengguna_baru,
    (
        SELECT COUNT(u2.id)
        FROM users u2
        WHERE DATE_FORMAT(u2.created_at, '%m') <= all_months.month
    ) AS total_pengguna
FROM
    all_months
LEFT JOIN
    monthly_data ON all_months.month = monthly_data.month
ORDER BY
    all_months.month;
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
