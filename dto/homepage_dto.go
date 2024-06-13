package dto

import (
	"capstone/entities"

	"github.com/google/uuid"
)

type HomepageRequest struct {
	CityName string
}

type HomepageResponse struct {
	RecommendDestinations    []RecommendDestination `json:"destinasi_rekomendasi"`
	NearbyDestinations       []NearbyDestination    `json:"destinasi_sekitar"`
	PopularDestinationMedias []DestinationMedia     `json:"konten_destinasi_populer"`
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

func ToPopularDestinationMedias(popularDestinations *[]entities.Destination) []DestinationMedia {
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

func toPopularDestinationMedia(destinationMedia *entities.DestinationMedia) DestinationMedia {
	return DestinationMedia{
		Id:    destinationMedia.Id,
		Title: destinationMedia.Title,
		Url:   destinationMedia.Url,
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
