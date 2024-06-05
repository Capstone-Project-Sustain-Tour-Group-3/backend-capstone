package seeds

import (
	"log"

	"capstone/entities"
)

func (s Seed) SeedProvinces() {
	provinces := []entities.Province{
		{
			Id:   "11",
			Name: "ACEH",
			Url:  "https://res.cloudinary.com/alta-minpro/image/upload/v1717500661/i1ewbj47dsetsvxwu82n.png",
		},
		{
			Id:   "12",
			Name: "SUMATERA UTARA",
			Url:  "https://res.cloudinary.com/alta-minpro/image/upload/v1717500696/byqnc7hv8jnbukayf5a0.jpg",
		},
		{
			Id:   "13",
			Name: "SUMATERA BARAT",
			Url:  "https://res.cloudinary.com/alta-minpro/image/upload/v1717500694/ybvw9ek9mb7wpbsnza6e.jpg",
		},
		{
			Id:   "14",
			Name: "RIAU",
			Url:  "https://res.cloudinary.com/alta-minpro/image/upload/v1717500688/o2pgbcqebmekozp5qh9x.jpg",
		},
		{
			Id:   "21",
			Name: "KEPULAUAN RIAU",
			Url:  "https://res.cloudinary.com/alta-minpro/image/upload/v1717500678/iwbrdbtr4hiuv3hkyvq7.jpg",
		},
		{
			Id:   "15",
			Name: "JAMBI",
			Url:  "https://res.cloudinary.com/alta-minpro/image/upload/v1717500669/sfpphdljanprgou9ozfw.jpg",
		},
		{
			Id:   "16",
			Name: "SUMATERA SELATAN",
			Url:  "https://res.cloudinary.com/alta-minpro/image/upload/v1717500695/qou9dl3mili92tkbjgxu.jpg",
		},
		{
			Id:   "19",
			Name: "BANGKA BELITUNG",
			Url:  "https://res.cloudinary.com/alta-minpro/image/upload/v1717500664/h85b2uqjhy9hdstqe9ez.jpg",
		},
		{
			Id:   "17",
			Name: "BENGKULU",
			Url:  "https://res.cloudinary.com/alta-minpro/image/upload/v1717500666/exxxdni2rsl0ffrnruub.jpg",
		},
		{
			Id:   "18",
			Name: "LAMPUNG",
			Url:  "https://res.cloudinary.com/alta-minpro/image/upload/v1717500679/xlyvm67zkknueqqwuyf4.jpg",
		},
		{
			Id:   "31",
			Name: "DKI JAKARTA",
			Url:  "https://res.cloudinary.com/alta-minpro/image/upload/v1717500669/zary2el8xl7h3e8ptilv.jpg",
		},
		{
			Id:   "32",
			Name: "JAWA BARAT",
			Url:  "https://res.cloudinary.com/alta-minpro/image/upload/v1717500668/b7dukeshrzdacmos7htd.png",
		},
		{
			Id:   "33",
			Name: "JAWA TENGAH",
			Url:  "https://res.cloudinary.com/alta-minpro/image/upload/v1717500670/ahfsrklckzqu1ifteoxw.jpg",
		},
		{
			Id:   "34",
			Name: "DI YOGYAKARTA",
			Url:  "https://res.cloudinary.com/alta-minpro/image/upload/v1717500672/qbq4nagbiwwoyq8xfiax.jpg",
		},
		{
			Id:   "35",
			Name: "JAWA TIMUR",
			Url:  "https://res.cloudinary.com/alta-minpro/image/upload/v1717500671/jfcmcj6v9gsvdjn7a4hw.jpg",
		},
		{
			Id:   "36",
			Name: "BANTEN",
			Url:  "https://res.cloudinary.com/alta-minpro/image/upload/v1717500665/cypsonvjjsiko0u78uxr.jpg",
		},
		{
			Id:   "51",
			Name: "BALI",
			Url:  "https://res.cloudinary.com/alta-minpro/image/upload/v1717500663/qbpre0kqd5r39hpeeaqm.jpg",
		},
		{
			Id:   "52",
			Name: "NUSA TENGGARA BARAT",
			Url:  "https://res.cloudinary.com/alta-minpro/image/upload/v1717500681/lil2xezqeryid7xverc0.jpg",
		},
		{
			Id:   "53",
			Name: "NUSA TENGGARA TIMUR",
			Url:  "https://res.cloudinary.com/alta-minpro/image/upload/v1717500682/iw2szz2b3vj7kv9clksq.jpg",
		},
		{
			Id:   "61",
			Name: "KALIMANTAN BARAT",
			Url:  "https://res.cloudinary.com/alta-minpro/image/upload/v1717500673/thmgquqyw56ku6qqa4ii.png",
		},
		{
			Id:   "62",
			Name: "KALIMANTAN TENGAH",
			Url:  "https://res.cloudinary.com/alta-minpro/image/upload/v1717500676/m93zhjbv1vvxlssnymkz.jpg",
		},
		{
			Id:   "63",
			Name: "KALIMANTAN SELATAN",
			Url:  "https://res.cloudinary.com/alta-minpro/image/upload/v1717500674/shhdyvaji5w1moczwzzl.jpg",
		},
		{
			Id:   "64",
			Name: "KALIMANTAN TIMUR",
			Url:  "https://res.cloudinary.com/alta-minpro/image/upload/v1717500677/c9ugqgnp26njckcdagsq.jpg",
		},
		{
			Id:   "65",
			Name: "KALIMANTAN UTARA",
			Url:  "https://res.cloudinary.com/alta-minpro/image/upload/v1717500675/xv0rz7ocrmzzerbhdmj3.jpg",
		},
		{
			Id:   "71",
			Name: "SULAWESI UTARA",
			Url:  "https://res.cloudinary.com/alta-minpro/image/upload/v1717500693/xgrv53rwy18zqllusl1t.jpg",
		},
		{
			Id:   "72",
			Name: "SULAWESI TENGAH",
			Url:  "https://res.cloudinary.com/alta-minpro/image/upload/v1717500690/hsjgaviw00glavracf03.png",
		},
		{
			Id:   "73",
			Name: "SULAWESI SELATAN",
			Url:  "https://res.cloudinary.com/alta-minpro/image/upload/v1717500692/npt2n3a9q9y4vqbmxpt7.png",
		},
		{
			Id:   "74",
			Name: "SULAWESI TENGGARA",
			Url:  "https://res.cloudinary.com/alta-minpro/image/upload/v1717500689/bojakznjjguta4gb1oxr.jpg",
		},
		{
			Id:   "75",
			Name: "GORONTALO",
			Url:  "https://res.cloudinary.com/alta-minpro/image/upload/v1717500667/gno3upr5qrqjry9bpkdn.png",
		},
		{
			Id:   "76",
			Name: "SULAWESI BARAT",
			Url:  "https://res.cloudinary.com/alta-minpro/image/upload/v1717500691/ndqlzaz96t79aztslogp.jpg",
		},
		{
			Id:   "81",
			Name: "MALUKU",
			Url:  "https://res.cloudinary.com/alta-minpro/image/upload/v1717500680/thg83qidu82ginztiwk9.jpg",
		},
		{
			Id:   "82",
			Name: "MALUKU UTARA",
			Url:  "https://res.cloudinary.com/alta-minpro/image/upload/v1717500680/kzka3kshanam7ofq1xhd.jpg",
		},
		{
			Id:   "91",
			Name: "PAPUA",
			Url:  "https://res.cloudinary.com/alta-minpro/image/upload/v1717500687/oud4plqxloysj8658z9v.jpg",
		},
		{
			Id:   "92",
			Name: "PAPUA BARAT",
			Url:  "https://res.cloudinary.com/alta-minpro/image/upload/v1717500684/ix7acvkzooy2aoh7og9l.jpg",
		},
		{
			Id:   "93",
			Name: "PAPUA SELATAN",
			Url:  "https://res.cloudinary.com/alta-minpro/image/upload/v1717500686/tekxeryuzmkoiydqghq6.jpg",
		},
		{
			Id:   "94",
			Name: "PAPUA TENGAH",
			Url:  "https://res.cloudinary.com/alta-minpro/image/upload/v1717500686/dwyh0ccnmlucgc4kjwst.jpg",
		},
		{
			Id:   "95",
			Name: "PAPUA PEGUNUNGAN",
			Url:  "https://res.cloudinary.com/alta-minpro/image/upload/v1717500685/qefq7zxevvhlsgelti8g.jpg",
		},
	}

	for _, province := range provinces {
		if err := s.db.Where(entities.Province{Name: province.Name}).
			FirstOrCreate(&province).Error; err != nil {
			log.Fatalf("failed to create province: %v", err)
		}
	}
}

// func getRandomProvince(s Seed) (*entities.Province, error) {
//	var province entities.Province
//
//	if err := s.db.Order("RAND()").First(&province).Error; err != nil {
//		return nil, err
//	}
//
//	return &province, nil
//}
