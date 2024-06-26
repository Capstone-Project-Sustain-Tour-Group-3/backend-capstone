package dto

import (
	"strings"

	"capstone/entities"

	"github.com/google/uuid"
)

type HomepageRequest struct {
	CityName string
}

type HomepageResponse struct {
	RecommendDestinations    []RecommendDestination `json:"destinasi_rekomendasi"`
	NearbyDestinations       []NearbyDestination    `json:"destinasi_sekitar"`
	PopularDestinationMedias []PopularDestination   `json:"destinasi_populer"`
}

type RecommendDestination struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"nama"`
	ImageUrl *string   `json:"url_gambar"`
	Province Province  `json:"provinsi"`
}

type NearbyDestination struct {
	Id                 uuid.UUID `json:"id"`
	Name               string    `json:"nama"`
	DestinationAddress Address   `json:"lokasi"`
	ImageUrl           *string   `json:"url_gambar"`
}

type PopularDestination struct {
	Id      uuid.UUID        `json:"id"`
	Name    string           `json:"nama"`
	Content DestinationMedia `json:"konten"`
}

type DestinationMedia struct {
	Id    uuid.UUID `json:"id"`
	Title string    `json:"judul"`
	Url   string    `json:"url"`
}

type Address struct {
	City     City     `json:"kota"`
	Province Province `json:"provinsi"`
}

type City struct {
	Name string `json:"nama"`
}

type Province struct {
	Name string `json:"nama"`
}

func ToHomepageRequest(cityName string) *HomepageRequest {
	var translatedCityName string
	switch strings.ToLower(cityName) {
	case "north sumatra":
		translatedCityName = "Sumatera Utara"
	case "west sumatra":
		translatedCityName = "Sumatera Barat"
	case "south sumatra":
		translatedCityName = "Sumatera Selatan"
	case "riau islands":
		translatedCityName = "Kepulauan Riau"
	case "bangka belitung islands":
		translatedCityName = "Bangka Belitung"
	case "west java":
		translatedCityName = "Jawa Barat"
	case "central java":
		translatedCityName = "Jawa Tengah"
	case "east java":
		translatedCityName = "Jawa Timur"
	case "west nusa tenggara":
		translatedCityName = "Nusa Tenggara Barat"
	case "east nusa tenggara":
		translatedCityName = "Nusa Tenggara Timur"
	case "west kalimantan":
		translatedCityName = "Kalimantan Barat"
	case "central kalimantan":
		translatedCityName = "Kalimantan Tengah"
	case "north kalimantan":
		translatedCityName = "Kalimantan Utara"
	case "east kalimantan":
		translatedCityName = "Kalimantan Timur"
	case "south kalimantan":
		translatedCityName = "Kalimantan Selatan"
	case "south sulawesi":
		translatedCityName = "Sulawesi Selatan"
	case "north sulawesi":
		translatedCityName = "Sulawesi Utara"
	case "central sulawesi":
		translatedCityName = "Sulawesi Tengah"
	case "west sulawesi":
		translatedCityName = "Sulawesi Barat"
	case "southeast sulawesi":
		translatedCityName = "Sulawesi Tenggara"
	case "south east sulawesi":
		translatedCityName = "Sulawesi Tenggara"
	case "north maluku":
		translatedCityName = "Maluku Utara"
	case "west papua":
		translatedCityName = "Papua Barat"
	case "south papua":
		translatedCityName = "Papua Selatan"
	case "central papua":
		translatedCityName = "Papua Tengah"
	case "mountain papua":
		translatedCityName = "Papua Pegunungan"
	case "southwest papua":
		translatedCityName = "Papua Barat Daya"
	case "south west papua":
		translatedCityName = "Papua Barat Daya"
	default:
		translatedCityName = cityName
	}

	return &HomepageRequest{
		CityName: translatedCityName,
	}
}

func ToRecommendDestinations(recommendDestinations *[]entities.Destination) []RecommendDestination {
	var recommendDestinationResponse []RecommendDestination

	for _, destination := range *recommendDestinations {
		recommendDestinationResponse = append(
			recommendDestinationResponse,
			toRecommendDestination(&destination),
		)
	}

	return recommendDestinationResponse
}

func ToNearbyDestinations(nearbyDestination *[]entities.Destination) []NearbyDestination {
	var nearbyDestinationResponse []NearbyDestination

	for _, destination := range *nearbyDestination {
		nearbyDestinationResponse = append(
			nearbyDestinationResponse,
			toNearbyDestination(&destination),
		)
	}

	return nearbyDestinationResponse
}

func ToPopularDestinations(popularDestinations *[]entities.Destination) []PopularDestination {
	const limit = 5
	var popularDestinationResponse []PopularDestination

	ctr := 0
	for _, destination := range *popularDestinations {
		if ctr == limit {
			break
		}

		if len(destination.DestinationMedias) == 0 {
			continue
		}

		popularDestinationResponse = append(
			popularDestinationResponse,
			toPopularDestination(&destination),
		)
		ctr += 1
	}

	return popularDestinationResponse
}

func toNearbyDestination(destination *entities.Destination) NearbyDestination {
	res := NearbyDestination{
		Id:   destination.Id,
		Name: destination.Name,
		DestinationAddress: Address{
			City: City{
				Name: destination.DestinationAddress.City.Name,
			},
			Province: Province{
				Name: destination.DestinationAddress.Province.Name,
			},
		},
		ImageUrl: nil,
	}

	if len(destination.DestinationMedias) != 0 {
		res.ImageUrl = &destination.DestinationMedias[0].Url
	}

	return res
}

func toPopularDestination(destination *entities.Destination) PopularDestination {
	return PopularDestination{
		Id:   destination.Id,
		Name: destination.Name,
		Content: DestinationMedia{
			Id:    destination.DestinationMedias[0].Id,
			Title: destination.DestinationMedias[0].Title,
			Url:   destination.DestinationMedias[0].Url,
		},
	}
}

func toRecommendDestination(destination *entities.Destination) RecommendDestination {
	var imageUrl *string = nil

	if len(destination.DestinationMedias) != 0 {
		imageUrl = &destination.DestinationMedias[0].Url
	}

	return RecommendDestination{
		Id:       destination.Id,
		Name:     destination.Name,
		ImageUrl: imageUrl,
		Province: Province{
			Name: destination.DestinationAddress.Province.Name,
		},
	}
}
