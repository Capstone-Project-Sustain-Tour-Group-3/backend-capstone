package dto

import (
	"capstone/entities"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type HomepageRequest struct {
	CityName string
}

type HomepageResponse struct {
	// RecommendDestination []entities.Destination      `json:"destinasi_rekomendasi"`
	NearbyDestination       []NearbyDestination `json:"destinasi_sekitar"`
	PopularDestinationMedia []DestinationMedia  `json:"konten_destinasi_populer"`
}

type NearbyDestination struct {
	Id                 uuid.UUID `json:"id"`
	Name               string    `json:"name"`
	DestinationAddress Address   `json:"lokasi"`
	ImageUrl           *string   `json:"url_gambar"`
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
	return &HomepageRequest{
		CityName: cityName,
	}
}

func ToHomepageResponse(
	nearbyDestination, popularDestinationMedia *[]entities.Destination,
) (*HomepageResponse, error) {
	nearbyDestinationResponse := toNearbyDestinations(nearbyDestination)
	popularDestinationMediaResponse := toPopularDestinationMedias(popularDestinationMedia)

	res := &HomepageResponse{
		NearbyDestination:       nearbyDestinationResponse,
		PopularDestinationMedia: popularDestinationMediaResponse,
	}

	return res, nil
}

func toNearbyDestinations(nearbyDestination *[]entities.Destination) []NearbyDestination {
	var nearbyDestinationResponse []NearbyDestination

	lenNearbyDestination := len(*nearbyDestination)
	limitNearby := 5

	if lenNearbyDestination < limitNearby {
		limitNearby = lenNearbyDestination
	}

	for _, destination := range (*nearbyDestination)[:limitNearby] {
		nearbyDestinationResponse = append(
			nearbyDestinationResponse,
			toNearbyDestination(&destination),
		)
	}

	return nearbyDestinationResponse
}

func toPopularDestinationMedias(popularDestinations *[]entities.Destination) []DestinationMedia {
	const limit = 5
	var popularDestinationMediaResponse []DestinationMedia

	ctr := 0
	for _, destination := range *popularDestinations {
		if ctr == limit {
			break
		}

		if len(destination.DestinationMedias) == 0 {
			continue
		}

		popularDestinationMediaResponse = append(
			popularDestinationMediaResponse,
			toPopularDestinationMedia(&destination.DestinationMedias[0]),
		)
		ctr += 1
	}

	return popularDestinationMediaResponse
}

func toNearbyDestination(destination *entities.Destination) NearbyDestination {
	var province string = destination.DestinationAddress.Province.Name

	caser := cases.Title(language.Indonesian)
	city := caser.String(strings.Split(destination.DestinationAddress.City.Name, " ")[1])

	if strings.HasPrefix(province, "DKI") {
		province = "DKI " + caser.String(province[4:])
	} else if strings.HasPrefix(province, "DI") {
		province = "DI " + caser.String(province[3:])
	} else {
		province = caser.String(destination.DestinationAddress.Province.Name)
	}

	res := NearbyDestination{
		Id:   destination.Id,
		Name: destination.Name,
		DestinationAddress: Address{
			City: City{
				Name: city,
			},
			Province: Province{
				Name: province,
			},
		},
		ImageUrl: nil,
	}

	if len(destination.DestinationMedias) != 0 {
		res.ImageUrl = &destination.DestinationMedias[0].Url
	}

	return res
}

func toPopularDestinationMedia(destinationMedia *entities.DestinationMedia) DestinationMedia {
	return DestinationMedia{
		Id:    destinationMedia.Id,
		Title: destinationMedia.Title,
		Url:   destinationMedia.Url,
	}
}
