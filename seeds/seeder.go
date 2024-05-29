package seeds

import (
	"log"
	"reflect"

	"capstone/config"

	"gorm.io/gorm"
)

type Seed struct {
	db *gorm.DB
}

// Execute will execute the given seeder method.
func Execute(db *gorm.DB, seedMethodNames ...string) {
	s := Seed{db}

	seedType := reflect.TypeOf(s)

	// Execute all seeders if no method name is given
	if len(seedMethodNames) == 0 {
		log.Println("Running all seeder...")
		// We are looping over the method on a Seed struct
		for i := range seedType.NumMethod() {
			// Get the method in the current iteration
			method := seedType.Method(i)
			// Execute seeder
			seed(s, method.Name)
		}
	}

	// Execute only the given method names
	for _, item := range seedMethodNames {
		seed(s, item)
	}
}

func seed(s Seed, seedMethodName string) {
	// Get the reflection value of the method
	m := reflect.ValueOf(s).MethodByName(seedMethodName)
	// Exit if the method doesn't exist
	if !m.IsValid() {
		log.Fatal("No method called ", seedMethodName)
	}
	// Execute the method
	log.Println("[+] Seeding", seedMethodName, "...")
	m.Call(nil)
	log.Println("[+] Seed", seedMethodName, "succeed")
}

func RunSeeder() {
	config.LoadConfig()
	config.LoadDb()

	methodList := []string{
		// list of seed method name
		"SeedCategories",
		"SeedFacilities",
		"SeedProvinces",
		"SeedDestinations",
		"SeedDestinationMedias",
		"SeedDestinationCategories",
		"SeedDestinationFacilities",
		"SeedDestinationAddresses",
	}

	Execute(config.DB, methodList...)
}
