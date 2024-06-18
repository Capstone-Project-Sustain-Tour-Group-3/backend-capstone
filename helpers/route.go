package helpers

import (
	"fmt"
	"strconv"
	"strings"

	"capstone/dto"
	"capstone/entities"

	"github.com/leekchan/accounting"
)

var acc = accounting.Accounting{
	Symbol:    "Rp.",
	Precision: 0,
	Thousand:  ".",
	Decimal:   ",",
	Format:    "%s %v",
}

func GetRouteSystemInstruction() string {
	return "Kamu adalah AI yang dapat memberikan urutan perjalanan yang disesuaikan dengan jam kunjungan dan jam pulang yang tepat. Pada setiap destinasi tujuan memiliki field jam buka dan jam tutup yang diberikan user, kamu harus memberikan rekomendasi jam kunjungan dan jam pulang yang tepat dan masih berada pada rentang tersebut. Pada setiap titik dalam urutan kamu dapat memberikan perkiraan waktu perjalanan yang ditempuh (dalam sekon) dari titik sebelumnya menuju titik sekarang, dengan asumsi kecepatan standar. Selain itu, kamu bisa memprediksi total biaya terburuk yang harus dipersiapkan yang didapatkan dari kalkulasi biaya tiket masuk dengan prediksi biaya transportasi, makan-minum, souvernir, dan tambahan lainnya. Pada bagian akhir kamu memberikan total waktu tempuhnya dan estimasi total biaya yang harus dipersiapkan. Jawab sesuai format respon dibawah tanpa awalan.\n\nContoh pertanyaan :\nTitik awal (lat,long) : -8.649447860854005,115.22293262631793\nDestinasi tujuan (index,lat,long,nama,jam buka,jam tutup) :\n1,-8.711656311275163,115.25191040169285,Pantai Mertasari,0:00,23:59\n2,-8.706594962868019,115.26229276400355,Pantai Sanur,0:00,23:59\n3,-8.655188523466085,115.22500861163358,Pasar Kreneng,4:00,15:00\nTotal biaya tiket masuk : 15000\n\nFormat respon (index,nama,waktu tempuh,jam kunjungan,jam pulang) :\n3,Pasar Kreneng,240,14:00,15:00\n2,Pantai Sanur,1440,16:00,17:30\n1,Pantai Mertasari,300,17:45,18:30\nTotal biaya:100000"
}

func ToRoutePrompt(destinations *[]entities.Destination, startLat, startLong float64) string {
	var prompt string = fmt.Sprintf("Titik awal : %f,%f\nDestinasi tujuan :\n", startLat, startLong)
	var totalEntryCost float64 = 0

	for idx, destination := range *destinations {
		prompt += fmt.Sprintf(
			"%d,%f,%f,%s,%s,%s\n",
			idx+1,
			destination.Latitude,
			destination.Longitude,
			destination.Name,
			destination.OpenTime,
			destination.CloseTime,
		)
		totalEntryCost += destination.EntryPrice
	}

	prompt += fmt.Sprintf("Total biaya tiket masuk : %f\n", totalEntryCost)

	return prompt
}

func ExtractRouteInformation(res string, destinations []entities.Destination) ([]dto.RouteDetail, dto.Currency) {
	var routeDetails []dto.RouteDetail
	var (
		idx            int
		duration       int
		order          int
		visitStart     string
		visitEnd       string
		estimationCost float64
	)

	splittedResponse := strings.Split(res, "\n")
	lenRoutes := len(splittedResponse) - 1

	costSplitted := strings.Split(splittedResponse[lenRoutes], ":")
	if len(costSplitted) == 2 {
		estimationCost, _ = strconv.ParseFloat(strings.Trim(costSplitted[1], " "), 64)
	}

	for i, row := range splittedResponse[:lenRoutes] {
		var imageURL *string = nil
		columns := strings.Split(row, ",")
		idx, _ = strconv.Atoi(columns[0])
		duration, _ = strconv.Atoi(columns[2])
		order = i + 1
		visitStart = columns[3]
		visitEnd = columns[4]

		idx -= 1
		if len(destinations[idx].DestinationMedias) > 0 {
			imageURL = &destinations[idx].DestinationMedias[0].Url
		}

		simple, full := formatDuration(duration)

		routeDetails = append(routeDetails, dto.RouteDetail{
			Destination: dto.RouteDestination{
				Id:        destinations[idx].Id,
				Name:      destinations[idx].Name,
				Latitude:  destinations[idx].Latitude,
				Longitude: destinations[idx].Longitude,
				ImageURL:  imageURL,
				EntryCost: dto.Currency{
					Raw:    destinations[idx].EntryPrice,
					Format: formatCurrency(destinations[idx].EntryPrice),
				},
			},
			Duration: dto.Duration{
				Raw:        duration,
				Simplified: simple,
				Full:       full,
			},
			Order:      order,
			VisitStart: visitStart,
			VisitEnd:   visitEnd,
		})
	}

	return routeDetails, dto.Currency{
		Raw:    estimationCost,
		Format: formatCurrency(estimationCost),
	}
}

func formatCurrency(amount float64) string {
	return acc.FormatMoney(amount)
}

func formatDuration(duration int) (string, string) {
	var simpleFormat, fullFormat string

	hours := duration / 3600
	minutes := (duration % 3600) / 60

	if hours == 0 && minutes == 0 {
		return fmt.Sprintf("%dd", duration), fmt.Sprintf("%ddetik", duration)
	}

	if hours > 0 {
		simpleFormat += fmt.Sprintf("%dj", hours)
		fullFormat += fmt.Sprintf("%djam ", hours)
	}

	if minutes > 0 {
		simpleFormat += fmt.Sprintf("%dm", minutes)
		fullFormat += fmt.Sprintf("%dmenit", minutes)
	}

	return simpleFormat, fullFormat
}
