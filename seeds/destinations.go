package seeds

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"capstone/entities"

	"github.com/google/uuid"
)

func (s Seed) SeedDestinations() {
	file, err := os.Open("storage/merged_data_destination.csv")
	if err != nil {
		log.Printf("failed to open file: %v", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Baca baris pertama (header) dan abaikan
	if _, err = reader.Read(); err != nil {
		log.Printf("failed to read header row: %v", err)
		return
	}

	records, err := reader.ReadAll()
	if err != nil {
		log.Printf("failed to read CSV file: %v", err)
		return
	}

	// Seed untuk angka acak
	randSource := rand.NewSource(time.Now().UnixNano())
	randGenerator := rand.New(randSource)

	// Dapatkan peta fasilitas dan alamat
	facilities, err := getAllFacilities(s)
	if err != nil {
		log.Printf("failed to get facilities: %v", err)
		return
	}

	for _, record := range records {
		destination, err2 := createDestination(s, record, randGenerator)
		if err2 != nil {
			log.Printf("failed to create destination: %v", err2)
			continue
		}

		if err = processFacilities(s, destination, record[14], facilities); err != nil {
			log.Printf("failed to process facilities: %v", err)
			continue
		}

		if err = processMedia(s, destination, record); err != nil {
			log.Printf("failed to process media: %v", err)
			continue
		}

		if err = processAddress(s, destination, record); err != nil {
			log.Printf("failed to process address: %v", err)
			continue
		}
	}
}

func createDestination(s Seed, record []string, randGenerator *rand.Rand) (entities.Destination, error) {
	entryPrice, _ := strconv.ParseFloat(record[4], 64)
	longitude, _ := strconv.ParseFloat(record[6], 64)
	latitude, _ := strconv.ParseFloat(record[5], 64)

	visitCount := randGenerator.Intn(201) + 100 // Menghasilkan angka acak antara 100 dan 300
	destination := entities.Destination{
		Id:          uuid.New(),
		CategoryId:  uuid.MustParse(record[13]),
		Name:        record[0],
		Description: record[1],
		OpenTime:    record[2],
		CloseTime:   record[3],
		EntryPrice:  entryPrice,
		Longitude:   longitude,
		Latitude:    latitude,
		VisitCount:  visitCount,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	var temp entities.Destination
	if err := s.db.Where(entities.Destination{Name: destination.Name}).
		FirstOrInit(&temp).
		Error; err != nil {
		return entities.Destination{}, err
	}

	if temp.Id != uuid.Nil {
		return temp, nil
	}

	if err := s.db.Create(&destination).Error; err != nil {
		return destination, err
	}

	return destination, nil
}

func processFacilities(s Seed, destination entities.Destination, facilitiesStr string, facilities []entities.Facility) error {
	facilitiesItems := strings.Split(facilitiesStr, ", ")
	for _, facilityItem := range facilitiesItems {
		for _, facility := range facilities {
			if facility.Name == facilityItem {
				destinationFacility := entities.DestinationFacility{
					Id:            uuid.New(),
					DestinationId: destination.Id,
					FacilityId:    facility.Id,
				}

				var temp entities.DestinationFacility
				if err := s.db.
					Where(entities.DestinationFacility{DestinationId: destination.Id, FacilityId: facility.Id}).
					FirstOrInit(&temp).
					Error; err != nil {
					return err
				}

				if temp.Id != uuid.Nil {
					continue
				}

				if err := s.db.Create(&destinationFacility).Error; err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func processMedia(s Seed, destination entities.Destination, record []string) error {
	mediaColumns := []struct {
		Url   string
		Type  string
		Title string
	}{
		{record[16], "image", record[15]},
		{record[17], "image", record[15]},
		{record[18], "image", record[15]},
		{record[19], "video", record[15]},
		{record[20], "video", record[15]},
		{record[21], "video", record[15]},
	}

	for _, media := range mediaColumns {
		if media.Url != "" {
			destinationMedia := entities.DestinationMedia{
				Id:            uuid.New(),
				DestinationId: destination.Id,
				Url:           media.Url,
				Type:          media.Type,
				Title:         media.Title,
				CreatedAt:     time.Now(),
				UpdatedAt:     time.Now(),
			}

			var temp entities.DestinationMedia
			if err := s.db.
				Where(entities.DestinationMedia{DestinationId: destination.Id, Url: media.Url}).
				FirstOrInit(&temp).
				Error; err != nil {
				return err
			}

			if temp.Id != uuid.Nil {
				continue
			}

			if err := s.db.Create(&destinationMedia).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

func processAddress(s Seed, destination entities.Destination, record []string) error {
	provinceName := record[8]
	cityName := record[9]
	subdistrictName := record[10]
	streetName := record[11]
	postalCode := record[12]

	province, err := getProvinceByName(s, provinceName)
	if err != nil {
		return err
	}

	city, err := getCityByName(s, cityName)
	if err != nil {
		return err
	}

	subdistrict, err := getSubdistrictByName(s, subdistrictName)
	if err != nil {
		return err
	}

	destinationAddress := entities.DestinationAddress{
		Id:            uuid.New(),
		DestinationId: destination.Id,
		ProvinceId:    province.Id,
		CityId:        city.Id,
		SubdistrictId: subdistrict.Id,
		StreetName:    streetName,
		PostalCode:    postalCode,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	var temp entities.DestinationAddress
	if err = s.db.
		Where(entities.DestinationAddress{DestinationId: destination.Id}).
		FirstOrInit(&temp).
		Error; err != nil {
		return err
	}

	if temp.Id != uuid.Nil {
		return nil
	}

	if err = s.db.Create(&destinationAddress).Error; err != nil {
		return err
	}

	return nil
}

func getAllFacilities(s Seed) ([]entities.Facility, error) {
	var facilities []entities.Facility

	if err := s.db.Find(&facilities).Error; err != nil {
		return nil, err
	}

	return facilities, nil
}

func getProvinceByName(s Seed, name string) (*entities.Province, error) {
	var province entities.Province
	if err := s.db.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(name)+"%").First(&province).Error; err != nil {
		return nil, err
	}
	return &province, nil
}

func getCityByName(s Seed, name string) (*entities.City, error) {
	var city entities.City
	if err := s.db.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(name)+"%").First(&city).Error; err != nil {
		return nil, err
	}
	return &city, nil
}

func getSubdistrictByName(s Seed, name string) (*entities.Subdistrict, error) {
	var subdistrict entities.Subdistrict
	if err := s.db.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(name)+"%").First(&subdistrict).Error; err != nil {
		return nil, err
	}
	return &subdistrict, nil
}
