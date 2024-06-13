package seeds

import (
	"log"

	"capstone/entities"
)

func (s Seed) SeedSubdistricts() {
	subdistricts := []entities.Subdistrict{
		{
			Id:     "110101",
			CityId: "1101",
			Name:   "Bakongan",
		},
		{
			Id:     "110102",
			CityId: "1101",
			Name:   "Kluet Utara",
		},
		{
			Id:     "110103",
			CityId: "1101",
			Name:   "Kluet Selatan",
		},
		{
			Id:     "110104",
			CityId: "1101",
			Name:   "Labuhanhaji",
		},
		{
			Id:     "110105",
			CityId: "1101",
			Name:   "Meukek",
		},
		{
			Id:     "110106",
			CityId: "1101",
			Name:   "Samadua",
		},
		{
			Id:     "110107",
			CityId: "1101",
			Name:   "Sawang",
		},
		{
			Id:     "110108",
			CityId: "1101",
			Name:   "Tapaktuan",
		},
		{
			Id:     "110109",
			CityId: "1101",
			Name:   "Trumon",
		},
		{
			Id:     "110110",
			CityId: "1101",
			Name:   "Pasi Raja",
		},
		{
			Id:     "110111",
			CityId: "1101",
			Name:   "Labuhan Haji Timur",
		},
		{
			Id:     "110112",
			CityId: "1101",
			Name:   "Labuhan Haji Barat",
		},
		{
			Id:     "110113",
			CityId: "1101",
			Name:   "Kluet Tengah",
		},
		{
			Id:     "110114",
			CityId: "1101",
			Name:   "Kluet Timur",
		},
		{
			Id:     "110115",
			CityId: "1101",
			Name:   "Bakongan Timur",
		},
		{
			Id:     "110116",
			CityId: "1101",
			Name:   "Trumon Timur",
		},
		{
			Id:     "110117",
			CityId: "1101",
			Name:   "Kota Bahagia",
		},
		{
			Id:     "110118",
			CityId: "1101",
			Name:   "Trumon Tengah",
		},
		{
			Id:     "110201",
			CityId: "1102",
			Name:   "Lawe Alas",
		},
		{
			Id:     "110202",
			CityId: "1102",
			Name:   "Lawe Sigala-Gala",
		},
		{
			Id:     "110203",
			CityId: "1102",
			Name:   "Bambel",
		},
		{
			Id:     "110204",
			CityId: "1102",
			Name:   "Babussalam",
		},
		{
			Id:     "110205",
			CityId: "1102",
			Name:   "Badar",
		},
		{
			Id:     "110206",
			CityId: "1102",
			Name:   "Babul Makmur",
		},
		{
			Id:     "110207",
			CityId: "1102",
			Name:   "Darul Hasanah",
		},
		{
			Id:     "110208",
			CityId: "1102",
			Name:   "Lawe Bulan",
		},
		{
			Id:     "110209",
			CityId: "1102",
			Name:   "Bukit Tusam",
		},
		{
			Id:     "110210",
			CityId: "1102",
			Name:   "Semadam",
		},
		{
			Id:     "110211",
			CityId: "1102",
			Name:   "Babul Rahmah",
		},
		{
			Id:     "110212",
			CityId: "1102",
			Name:   "Ketambe",
		},
		{
			Id:     "110213",
			CityId: "1102",
			Name:   "Deleng Pokhkisen",
		},
		{
			Id:     "110214",
			CityId: "1102",
			Name:   "Lawe Sumur",
		},
		{
			Id:     "110215",
			CityId: "1102",
			Name:   "Tanoh Alas",
		},
		{
			Id:     "110216",
			CityId: "1102",
			Name:   "Leuser",
		},
		{
			Id:     "110301",
			CityId: "1103",
			Name:   "Darul Aman",
		},
		{
			Id:     "110302",
			CityId: "1103",
			Name:   "Julok",
		},
		{
			Id:     "110303",
			CityId: "1103",
			Name:   "Idi Rayeuk",
		},
		{
			Id:     "110304",
			CityId: "1103",
			Name:   "Birem Bayeun",
		},
		{
			Id:     "110305",
			CityId: "1103",
			Name:   "Serbajadi",
		},
		{
			Id:     "110306",
			CityId: "1103",
			Name:   "Nurussalam",
		},
		{
			Id:     "110307",
			CityId: "1103",
			Name:   "Peureulak",
		},
		{
			Id:     "110308",
			CityId: "1103",
			Name:   "Rantau Selamat",
		},
		{
			Id:     "110309",
			CityId: "1103",
			Name:   "Simpang Ulim",
		},
		{
			Id:     "110310",
			CityId: "1103",
			Name:   "Ranto Peureulak",
		},
		{
			Id:     "110311",
			CityId: "1103",
			Name:   "Pante Bidari",
		},
		{
			Id:     "110312",
			CityId: "1103",
			Name:   "Madat",
		},
		{
			Id:     "110313",
			CityId: "1103",
			Name:   "Indra Makmu",
		},
		{
			Id:     "110314",
			CityId: "1103",
			Name:   "Idi Tunong",
		},
		{
			Id:     "110315",
			CityId: "1103",
			Name:   "Banda Alam",
		},
		{
			Id:     "110316",
			CityId: "1103",
			Name:   "Peudawa",
		},
		{
			Id:     "110317",
			CityId: "1103",
			Name:   "Peureulak Timur",
		},
		{
			Id:     "110318",
			CityId: "1103",
			Name:   "Peureulak Barat",
		},
		{
			Id:     "110319",
			CityId: "1103",
			Name:   "Sungai Raya",
		},
		{
			Id:     "110320",
			CityId: "1103",
			Name:   "Simpang Jernih",
		},
		{
			Id:     "110321",
			CityId: "1103",
			Name:   "Darul Ihsan",
		},
		{
			Id:     "110322",
			CityId: "1103",
			Name:   "Darul Falah",
		},
		{
			Id:     "110323",
			CityId: "1103",
			Name:   "Idi Timur",
		},
		{
			Id:     "110324",
			CityId: "1103",
			Name:   "Peunaron",
		},
		{
			Id:     "110401",
			CityId: "1104",
			Name:   "Linge",
		},
		{
			Id:     "110402",
			CityId: "1104",
			Name:   "Silih Nara",
		},
		{
			Id:     "110403",
			CityId: "1104",
			Name:   "Bebesen",
		},
		{
			Id:     "110407",
			CityId: "1104",
			Name:   "Pegasing",
		},
		{
			Id:     "110408",
			CityId: "1104",
			Name:   "Bintang",
		},
		{
			Id:     "110410",
			CityId: "1104",
			Name:   "Ketol",
		},
		{
			Id:     "110411",
			CityId: "1104",
			Name:   "Kebayakan",
		},
		{
			Id:     "110412",
			CityId: "1104",
			Name:   "Kute Panang",
		},
		{
			Id:     "110413",
			CityId: "1104",
			Name:   "Celala",
		},
		{
			Id:     "110417",
			CityId: "1104",
			Name:   "Laut Tawar",
		},
		{
			Id:     "110418",
			CityId: "1104",
			Name:   "Atu Lintang",
		},
		{
			Id:     "110419",
			CityId: "1104",
			Name:   "Jagong Jeget",
		},
		{
			Id:     "110420",
			CityId: "1104",
			Name:   "Bies",
		},
		{
			Id:     "110421",
			CityId: "1104",
			Name:   "Rusip Antara",
		},
		{
			Id:     "110501",
			CityId: "1105",
			Name:   "Johan Pahlawan",
		},
		{
			Id:     "110502",
			CityId: "1105",
			Name:   "Kaway XVI",
		},
		{
			Id:     "110503",
			CityId: "1105",
			Name:   "Sungai Mas",
		},
		{
			Id:     "110504",
			CityId: "1105",
			Name:   "Woyla",
		},
		{
			Id:     "110505",
			CityId: "1105",
			Name:   "Samatiga",
		},
		{
			Id:     "110506",
			CityId: "1105",
			Name:   "Bubon",
		},
		{
			Id:     "110507",
			CityId: "1105",
			Name:   "Arongan Lambalek",
		},
		{
			Id:     "110508",
			CityId: "1105",
			Name:   "Pante Ceureumen",
		},
		{
			Id:     "110509",
			CityId: "1105",
			Name:   "Meureubo",
		},
		{
			Id:     "110510",
			CityId: "1105",
			Name:   "Woyla Barat",
		},
		{
			Id:     "110511",
			CityId: "1105",
			Name:   "Woyla Timur",
		},
		{
			Id:     "110512",
			CityId: "1105",
			Name:   "Panton Reu",
		},
		{
			Id:     "110601",
			CityId: "1106",
			Name:   "Lhoong",
		},
		{
			Id:     "110602",
			CityId: "1106",
			Name:   "Lhoknga",
		},
		{
			Id:     "110603",
			CityId: "1106",
			Name:   "Indrapuri",
		},
		{
			Id:     "110604",
			CityId: "1106",
			Name:   "Seulimeum",
		},
		{
			Id:     "110605",
			CityId: "1106",
			Name:   "Montasik",
		},
		{
			Id:     "110606",
			CityId: "1106",
			Name:   "Sukamakmur",
		},
		{
			Id:     "110607",
			CityId: "1106",
			Name:   "Darul Imarah",
		},
		{
			Id:     "110608",
			CityId: "1106",
			Name:   "Peukan Bada",
		},
		{
			Id:     "110609",
			CityId: "1106",
			Name:   "Mesjid Raya",
		},
		{
			Id:     "110610",
			CityId: "1106",
			Name:   "Ingin Jaya",
		},
		{
			Id:     "110611",
			CityId: "1106",
			Name:   "Kuta Baro",
		},
		{
			Id:     "110612",
			CityId: "1106",
			Name:   "Darussalam",
		},
		{
			Id:     "110613",
			CityId: "1106",
			Name:   "Pulo Aceh",
		},
		{
			Id:     "110614",
			CityId: "1106",
			Name:   "Lembah Seulawah",
		},
		{
			Id:     "110615",
			CityId: "1106",
			Name:   "Kota Jantho",
		},
		{
			Id:     "110616",
			CityId: "1106",
			Name:   "Kuta Cot Glie",
		},
		{
			Id:     "110617",
			CityId: "1106",
			Name:   "Kuta Malaka",
		},
		{
			Id:     "110618",
			CityId: "1106",
			Name:   "Simpang Tiga",
		},
		{
			Id:     "110619",
			CityId: "1106",
			Name:   "Darul Kamal",
		},
		{
			Id:     "110620",
			CityId: "1106",
			Name:   "Baitussalam",
		},
		{
			Id:     "110621",
			CityId: "1106",
			Name:   "Krueng Barona Jaya",
		},
		{
			Id:     "110622",
			CityId: "1106",
			Name:   "Leupung",
		},
		{
			Id:     "110623",
			CityId: "1106",
			Name:   "Blang Bintang",
		},
		{
			Id:     "110703",
			CityId: "1107",
			Name:   "Batee",
		},
		{
			Id:     "110704",
			CityId: "1107",
			Name:   "Delima",
		},
		{
			Id:     "110705",
			CityId: "1107",
			Name:   "Geumpang",
		},
		{
			Id:     "110706",
			CityId: "1107",
			Name:   "Glumpang Tiga",
		},
		{
			Id:     "110707",
			CityId: "1107",
			Name:   "Indrajaya",
		},
		{
			Id:     "110708",
			CityId: "1107",
			Name:   "Kembang Tanjong",
		},
		{
			Id:     "110709",
			CityId: "1107",
			Name:   "Kota Sigli",
		},
		{
			Id:     "110711",
			CityId: "1107",
			Name:   "Mila",
		},
		{
			Id:     "110712",
			CityId: "1107",
			Name:   "Muara Tiga",
		},
		{
			Id:     "110713",
			CityId: "1107",
			Name:   "Mutiara",
		},
		{
			Id:     "110714",
			CityId: "1107",
			Name:   "Padang Tiji",
		},
		{
			Id:     "110715",
			CityId: "1107",
			Name:   "Peukan Baro",
		},
		{
			Id:     "110716",
			CityId: "1107",
			Name:   "Pidie",
		},
		{
			Id:     "110717",
			CityId: "1107",
			Name:   "Sakti",
		},
		{
			Id:     "110718",
			CityId: "1107",
			Name:   "Simpang Tiga",
		},
		{
			Id:     "110719",
			CityId: "1107",
			Name:   "Tangse",
		},
		{
			Id:     "110721",
			CityId: "1107",
			Name:   "Tiro/Truseb",
		},
		{
			Id:     "110722",
			CityId: "1107",
			Name:   "Keumala",
		},
		{
			Id:     "110724",
			CityId: "1107",
			Name:   "Mutiara Timur",
		},
		{
			Id:     "110725",
			CityId: "1107",
			Name:   "Grong-grong",
		},
		{
			Id:     "110727",
			CityId: "1107",
			Name:   "Mane",
		},
		{
			Id:     "110729",
			CityId: "1107",
			Name:   "Glumpang Baro",
		},
		{
			Id:     "110731",
			CityId: "1107",
			Name:   "Titeue",
		},
		{
			Id:     "110801",
			CityId: "1108",
			Name:   "Baktiya",
		},
		{
			Id:     "110802",
			CityId: "1108",
			Name:   "Dewantara",
		},
		{
			Id:     "110803",
			CityId: "1108",
			Name:   "Kuta Makmur",
		},
		{
			Id:     "110804",
			CityId: "1108",
			Name:   "Lhoksukon",
		},
		{
			Id:     "110805",
			CityId: "1108",
			Name:   "Matangkuli",
		},
		{
			Id:     "110806",
			CityId: "1108",
			Name:   "Muara Batu",
		},
		{
			Id:     "110807",
			CityId: "1108",
			Name:   "Meurah Mulia",
		},
		{
			Id:     "110808",
			CityId: "1108",
			Name:   "Samudera",
		},
		{
			Id:     "110809",
			CityId: "1108",
			Name:   "Seunuddon",
		},
		{
			Id:     "110810",
			CityId: "1108",
			Name:   "Syamtalira Aron",
		},
		{
			Id:     "110811",
			CityId: "1108",
			Name:   "Syamtalira Bayu",
		},
		{
			Id:     "110812",
			CityId: "1108",
			Name:   "Tanah Luas",
		},
		{
			Id:     "110813",
			CityId: "1108",
			Name:   "Tanah Pasir",
		},
		{
			Id:     "110814",
			CityId: "1108",
			Name:   "T. Jambo Aye",
		},
		{
			Id:     "110815",
			CityId: "1108",
			Name:   "Sawang",
		},
		{
			Id:     "110816",
			CityId: "1108",
			Name:   "Nisam",
		},
		{
			Id:     "110817",
			CityId: "1108",
			Name:   "Cot Girek",
		},
		{
			Id:     "110818",
			CityId: "1108",
			Name:   "Langkahan",
		},
		{
			Id:     "110819",
			CityId: "1108",
			Name:   "Baktiya Barat",
		},
		{
			Id:     "110820",
			CityId: "1108",
			Name:   "Paya Bakong",
		},
		{
			Id:     "110821",
			CityId: "1108",
			Name:   "Nibong",
		},
		{
			Id:     "110822",
			CityId: "1108",
			Name:   "Simpang Kramat",
		},
		{
			Id:     "110823",
			CityId: "1108",
			Name:   "Lapang",
		},
		{
			Id:     "110824",
			CityId: "1108",
			Name:   "Pirak Timur",
		},
		{
			Id:     "110825",
			CityId: "1108",
			Name:   "Geuredong Pase",
		},
		{
			Id:     "110826",
			CityId: "1108",
			Name:   "Banda Baro",
		},
		{
			Id:     "110827",
			CityId: "1108",
			Name:   "Nisam Antara",
		},
		{
			Id:     "110901",
			CityId: "1109",
			Name:   "Simeulue Tengah",
		},
		{
			Id:     "110902",
			CityId: "1109",
			Name:   "Salang",
		},
		{
			Id:     "110903",
			CityId: "1109",
			Name:   "Teupah Barat",
		},
		{
			Id:     "110904",
			CityId: "1109",
			Name:   "Simeulue Timur",
		},
		{
			Id:     "110905",
			CityId: "1109",
			Name:   "Teluk Dalam",
		},
		{
			Id:     "110906",
			CityId: "1109",
			Name:   "Simeulue Barat",
		},
		{
			Id:     "110907",
			CityId: "1109",
			Name:   "Teupah Selatan",
		},
		{
			Id:     "110908",
			CityId: "1109",
			Name:   "Alafan",
		},
		{
			Id:     "110909",
			CityId: "1109",
			Name:   "Teupah Tengah",
		},
		{
			Id:     "110910",
			CityId: "1109",
			Name:   "Simeulue Cut",
		},
		{
			Id:     "111001",
			CityId: "1110",
			Name:   "Pulau Banyak",
		},
		{
			Id:     "111002",
			CityId: "1110",
			Name:   "Simpang Kanan",
		},
		{
			Id:     "111004",
			CityId: "1110",
			Name:   "Singkil",
		},
		{
			Id:     "111006",
			CityId: "1110",
			Name:   "Gunung Meriah",
		},
		{
			Id:     "111009",
			CityId: "1110",
			Name:   "Kota Baharu",
		},
		{
			Id:     "111010",
			CityId: "1110",
			Name:   "Singkil Utara",
		},
		{
			Id:     "111011",
			CityId: "1110",
			Name:   "Danau Paris",
		},
		{
			Id:     "111012",
			CityId: "1110",
			Name:   "Suro Makmur",
		},
		{
			Id:     "111013",
			CityId: "1110",
			Name:   "Singkohor",
		},
		{
			Id:     "111014",
			CityId: "1110",
			Name:   "Kuala Baru",
		},
		{
			Id:     "111016",
			CityId: "1110",
			Name:   "Pulau Banyak Barat",
		},
		{
			Id:     "111101",
			CityId: "1111",
			Name:   "Samalanga",
		},
		{
			Id:     "111102",
			CityId: "1111",
			Name:   "Jeunieb",
		},
		{
			Id:     "111103",
			CityId: "1111",
			Name:   "Peudada",
		},
		{
			Id:     "111104",
			CityId: "1111",
			Name:   "Jeumpa",
		},
		{
			Id:     "111105",
			CityId: "1111",
			Name:   "Peusangan",
		},
		{
			Id:     "111106",
			CityId: "1111",
			Name:   "Makmur",
		},
		{
			Id:     "111107",
			CityId: "1111",
			Name:   "Gandapura",
		},
		{
			Id:     "111108",
			CityId: "1111",
			Name:   "Pandrah",
		},
		{
			Id:     "111109",
			CityId: "1111",
			Name:   "Juli",
		},
		{
			Id:     "111110",
			CityId: "1111",
			Name:   "Jangka",
		},
		{
			Id:     "111111",
			CityId: "1111",
			Name:   "Simpang Mamplam",
		},
		{
			Id:     "111112",
			CityId: "1111",
			Name:   "Peulimbang",
		},
		{
			Id:     "111113",
			CityId: "1111",
			Name:   "Kota Juang",
		},
		{
			Id:     "111114",
			CityId: "1111",
			Name:   "Kuala",
		},
		{
			Id:     "111115",
			CityId: "1111",
			Name:   "Peusangan Siblah Krueng",
		},
		{
			Id:     "111116",
			CityId: "1111",
			Name:   "Peusangan Selatan",
		},
		{
			Id:     "111117",
			CityId: "1111",
			Name:   "Kuta Blang",
		},
		{
			Id:     "111201",
			CityId: "1112",
			Name:   "Blangpidie",
		},
		{
			Id:     "111202",
			CityId: "1112",
			Name:   "Tangan-Tangan",
		},
		{
			Id:     "111203",
			CityId: "1112",
			Name:   "Manggeng",
		},
		{
			Id:     "111204",
			CityId: "1112",
			Name:   "Susoh",
		},
		{
			Id:     "111205",
			CityId: "1112",
			Name:   "Kuala Batee",
		},
		{
			Id:     "111206",
			CityId: "1112",
			Name:   "Babah Rot",
		},
		{
			Id:     "111207",
			CityId: "1112",
			Name:   "Setia",
		},
		{
			Id:     "111208",
			CityId: "1112",
			Name:   "Jeumpa",
		},
		{
			Id:     "111209",
			CityId: "1112",
			Name:   "Lembah Sabil",
		},
		{
			Id:     "111301",
			CityId: "1113",
			Name:   "Blangkejeren",
		},
		{
			Id:     "111302",
			CityId: "1113",
			Name:   "Kutapanjang",
		},
		{
			Id:     "111303",
			CityId: "1113",
			Name:   "Rikit Gaib",
		},
		{
			Id:     "111304",
			CityId: "1113",
			Name:   "Terangun",
		},
		{
			Id:     "111305",
			CityId: "1113",
			Name:   "Pining",
		},
		{
			Id:     "111306",
			CityId: "1113",
			Name:   "Blangpegayon",
		},
		{
			Id:     "111307",
			CityId: "1113",
			Name:   "Puteri Betung",
		},
		{
			Id:     "111308",
			CityId: "1113",
			Name:   "Dabun Gelang",
		},
		{
			Id:     "111309",
			CityId: "1113",
			Name:   "Blangjerango",
		},
		{
			Id:     "111310",
			CityId: "1113",
			Name:   "Teripe Jaya",
		},
		{
			Id:     "111311",
			CityId: "1113",
			Name:   "Pantan Cuaca",
		},
		{
			Id:     "111401",
			CityId: "1114",
			Name:   "Teunom",
		},
		{
			Id:     "111402",
			CityId: "1114",
			Name:   "Krueng Sabee",
		},
		{
			Id:     "111403",
			CityId: "1114",
			Name:   "Setia Bakti",
		},
		{
			Id:     "111404",
			CityId: "1114",
			Name:   "Sampoi Niet",
		},
		{
			Id:     "111405",
			CityId: "1114",
			Name:   "Jaya",
		},
		{
			Id:     "111406",
			CityId: "1114",
			Name:   "Panga",
		},
		{
			Id:     "111407",
			CityId: "1114",
			Name:   "Indra Jaya",
		},
		{
			Id:     "111408",
			CityId: "1114",
			Name:   "Darul Hikmah",
		},
		{
			Id:     "111409",
			CityId: "1114",
			Name:   "Pasie Raya",
		},
		{
			Id:     "111501",
			CityId: "1115",
			Name:   "Kuala",
		},
		{
			Id:     "111502",
			CityId: "1115",
			Name:   "Seunagan",
		},
		{
			Id:     "111503",
			CityId: "1115",
			Name:   "Seunagan Timur",
		},
		{
			Id:     "111504",
			CityId: "1115",
			Name:   "Beutong",
		},
		{
			Id:     "111505",
			CityId: "1115",
			Name:   "Darul Makmur",
		},
		{
			Id:     "111506",
			CityId: "1115",
			Name:   "Suka Makmue",
		},
		{
			Id:     "111507",
			CityId: "1115",
			Name:   "Kuala Pesisir",
		},
		{
			Id:     "111508",
			CityId: "1115",
			Name:   "Tadu Raya",
		},
		{
			Id:     "111509",
			CityId: "1115",
			Name:   "Tripa Makmur",
		},
		{
			Id:     "111510",
			CityId: "1115",
			Name:   "Beutong Ateuh Banggalang",
		},
		{
			Id:     "111601",
			CityId: "1116",
			Name:   "Manyak Payed",
		},
		{
			Id:     "111602",
			CityId: "1116",
			Name:   "Bendahara",
		},
		{
			Id:     "111603",
			CityId: "1116",
			Name:   "Karang Baru",
		},
		{
			Id:     "111604",
			CityId: "1116",
			Name:   "Seruway",
		},
		{
			Id:     "111605",
			CityId: "1116",
			Name:   "Kota Kualasinpang",
		},
		{
			Id:     "111606",
			CityId: "1116",
			Name:   "Kejuruan Muda",
		},
		{
			Id:     "111607",
			CityId: "1116",
			Name:   "Tamiang Hulu",
		},
		{
			Id:     "111608",
			CityId: "1116",
			Name:   "Rantau",
		},
		{
			Id:     "111609",
			CityId: "1116",
			Name:   "Banda Mulia",
		},
		{
			Id:     "111610",
			CityId: "1116",
			Name:   "Bandar Pusaka",
		},
		{
			Id:     "111611",
			CityId: "1116",
			Name:   "Tenggulun",
		},
		{
			Id:     "111612",
			CityId: "1116",
			Name:   "Sekerak",
		},
		{
			Id:     "111701",
			CityId: "1117",
			Name:   "Pintu Rime Gayo",
		},
		{
			Id:     "111702",
			CityId: "1117",
			Name:   "Permata",
		},
		{
			Id:     "111703",
			CityId: "1117",
			Name:   "Syiah Utama",
		},
		{
			Id:     "111704",
			CityId: "1117",
			Name:   "Bandar",
		},
		{
			Id:     "111705",
			CityId: "1117",
			Name:   "Bukit",
		},
		{
			Id:     "111706",
			CityId: "1117",
			Name:   "Wih Pesam",
		},
		{
			Id:     "111707",
			CityId: "1117",
			Name:   "Timang gajah",
		},
		{
			Id:     "111708",
			CityId: "1117",
			Name:   "Bener Kelipah",
		},
		{
			Id:     "111709",
			CityId: "1117",
			Name:   "Mesidah",
		},
		{
			Id:     "111710",
			CityId: "1117",
			Name:   "Gajah Putih",
		},
		{
			Id:     "111801",
			CityId: "1118",
			Name:   "Meureudu",
		},
		{
			Id:     "111802",
			CityId: "1118",
			Name:   "Ulim",
		},
		{
			Id:     "111803",
			CityId: "1118",
			Name:   "Jangka Buaya",
		},
		{
			Id:     "111804",
			CityId: "1118",
			Name:   "Bandar Dua",
		},
		{
			Id:     "111805",
			CityId: "1118",
			Name:   "Meurah Dua",
		},
		{
			Id:     "111806",
			CityId: "1118",
			Name:   "Bandar Baru",
		},
		{
			Id:     "111807",
			CityId: "1118",
			Name:   "Panteraja",
		},
		{
			Id:     "111808",
			CityId: "1118",
			Name:   "Trienggadeng",
		},
		{
			Id:     "117101",
			CityId: "1171",
			Name:   "Baiturrahman",
		},
		{
			Id:     "117102",
			CityId: "1171",
			Name:   "Kuta Alam",
		},
		{
			Id:     "117103",
			CityId: "1171",
			Name:   "Meuraxa",
		},
		{
			Id:     "117104",
			CityId: "1171",
			Name:   "Syiah Kuala",
		},
		{
			Id:     "117105",
			CityId: "1171",
			Name:   "Lueng Bata",
		},
		{
			Id:     "117106",
			CityId: "1171",
			Name:   "Kuta Raja",
		},
		{
			Id:     "117107",
			CityId: "1171",
			Name:   "Banda Raya",
		},
		{
			Id:     "117108",
			CityId: "1171",
			Name:   "Jaya Baru",
		},
		{
			Id:     "117109",
			CityId: "1171",
			Name:   "Ulee Kareng",
		},
		{
			Id:     "117201",
			CityId: "1172",
			Name:   "Sukakarya",
		},
		{
			Id:     "117202",
			CityId: "1172",
			Name:   "Sukajaya",
		},
		{
			Id:     "117203",
			CityId: "1172",
			Name:   "Sukamakmue",
		},
		{
			Id:     "117301",
			CityId: "1173",
			Name:   "Muara Dua",
		},
		{
			Id:     "117302",
			CityId: "1173",
			Name:   "Banda Sakti",
		},
		{
			Id:     "117303",
			CityId: "1173",
			Name:   "Blang Mangat",
		},
		{
			Id:     "117304",
			CityId: "1173",
			Name:   "Muara Satu",
		},
		{
			Id:     "117401",
			CityId: "1174",
			Name:   "Langsa Timur",
		},
		{
			Id:     "117402",
			CityId: "1174",
			Name:   "Langsa Barat",
		},
		{
			Id:     "117403",
			CityId: "1174",
			Name:   "Langsa Kota",
		},
		{
			Id:     "117404",
			CityId: "1174",
			Name:   "Langsa Lama",
		},
		{
			Id:     "117405",
			CityId: "1174",
			Name:   "Langsa Baro",
		},
		{
			Id:     "117501",
			CityId: "1175",
			Name:   "Simpang Kiri",
		},
		{
			Id:     "117502",
			CityId: "1175",
			Name:   "Penanggalan",
		},
		{
			Id:     "117503",
			CityId: "1175",
			Name:   "Rundeng",
		},
		{
			Id:     "117504",
			CityId: "1175",
			Name:   "Sultan Daulat",
		},
		{
			Id:     "117505",
			CityId: "1175",
			Name:   "Longkib",
		},
		{
			Id:     "120101",
			CityId: "1201",
			Name:   "Barus",
		},
		{
			Id:     "120102",
			CityId: "1201",
			Name:   "Sorkam",
		},
		{
			Id:     "120103",
			CityId: "1201",
			Name:   "Pandan",
		},
		{
			Id:     "120104",
			CityId: "1201",
			Name:   "Pinangsori",
		},
		{
			Id:     "120105",
			CityId: "1201",
			Name:   "Manduamas",
		},
		{
			Id:     "120106",
			CityId: "1201",
			Name:   "Kolang",
		},
		{
			Id:     "120107",
			CityId: "1201",
			Name:   "Tapian Nauli",
		},
		{
			Id:     "120108",
			CityId: "1201",
			Name:   "Sibabangun",
		},
		{
			Id:     "120109",
			CityId: "1201",
			Name:   "Sosorgadong",
		},
		{
			Id:     "120110",
			CityId: "1201",
			Name:   "Sorkam Barat",
		},
		{
			Id:     "120111",
			CityId: "1201",
			Name:   "Sirandorung",
		},
		{
			Id:     "120112",
			CityId: "1201",
			Name:   "Andam Dewi",
		},
		{
			Id:     "120113",
			CityId: "1201",
			Name:   "Sitahuis",
		},
		{
			Id:     "120114",
			CityId: "1201",
			Name:   "Tukka",
		},
		{
			Id:     "120115",
			CityId: "1201",
			Name:   "Badiri",
		},
		{
			Id:     "120116",
			CityId: "1201",
			Name:   "Pasaribu Tobing",
		},
		{
			Id:     "120117",
			CityId: "1201",
			Name:   "Barus Utara",
		},
		{
			Id:     "120118",
			CityId: "1201",
			Name:   "Suka Bangun",
		},
		{
			Id:     "120119",
			CityId: "1201",
			Name:   "Lumut",
		},
		{
			Id:     "120120",
			CityId: "1201",
			Name:   "Sarudik",
		},
		{
			Id:     "120201",
			CityId: "1202",
			Name:   "Tarutung",
		},
		{
			Id:     "120202",
			CityId: "1202",
			Name:   "Siatas Barita",
		},
		{
			Id:     "120203",
			CityId: "1202",
			Name:   "Adian Koting",
		},
		{
			Id:     "120204",
			CityId: "1202",
			Name:   "Sipoholon",
		},
		{
			Id:     "120205",
			CityId: "1202",
			Name:   "Pahae Julu",
		},
		{
			Id:     "120206",
			CityId: "1202",
			Name:   "Pahae Jae",
		},
		{
			Id:     "120207",
			CityId: "1202",
			Name:   "Simangumban",
		},
		{
			Id:     "120208",
			CityId: "1202",
			Name:   "Purba Tua",
		},
		{
			Id:     "120209",
			CityId: "1202",
			Name:   "Siborong-Borong",
		},
		{
			Id:     "120210",
			CityId: "1202",
			Name:   "Pagaran",
		},
		{
			Id:     "120211",
			CityId: "1202",
			Name:   "Parmonangan",
		},
		{
			Id:     "120212",
			CityId: "1202",
			Name:   "Sipahutar",
		},
		{
			Id:     "120213",
			CityId: "1202",
			Name:   "Pangaribuan",
		},
		{
			Id:     "120214",
			CityId: "1202",
			Name:   "Garoga",
		},
		{
			Id:     "120215",
			CityId: "1202",
			Name:   "Muara",
		},
		{
			Id:     "120301",
			CityId: "1203",
			Name:   "Angkola Barat",
		},
		{
			Id:     "120302",
			CityId: "1203",
			Name:   "Batang Toru",
		},
		{
			Id:     "120303",
			CityId: "1203",
			Name:   "Angkola Timur",
		},
		{
			Id:     "120304",
			CityId: "1203",
			Name:   "Sipirok",
		},
		{
			Id:     "120305",
			CityId: "1203",
			Name:   "Saipar Dolok Hole",
		},
		{
			Id:     "120306",
			CityId: "1203",
			Name:   "Angkola Selatan",
		},
		{
			Id:     "120307",
			CityId: "1203",
			Name:   "Batang Angkola",
		},
		{
			Id:     "120314",
			CityId: "1203",
			Name:   "Arse",
		},
		{
			Id:     "120320",
			CityId: "1203",
			Name:   "Marancar",
		},
		{
			Id:     "120321",
			CityId: "1203",
			Name:   "Sayur Matinggi",
		},
		{
			Id:     "120322",
			CityId: "1203",
			Name:   "Aek Bilah",
		},
		{
			Id:     "120329",
			CityId: "1203",
			Name:   "Muara Batang Toru",
		},
		{
			Id:     "120330",
			CityId: "1203",
			Name:   "Tano Tombangan Angkola",
		},
		{
			Id:     "120331",
			CityId: "1203",
			Name:   "Angkola Sangkunur",
		},
		{
			Id:     "120332",
			CityId: "1203",
			Name:   "Angkola Muara Tais",
		},
		{
			Id:     "120405",
			CityId: "1204",
			Name:   "Hiliduho",
		},
		{
			Id:     "120406",
			CityId: "1204",
			Name:   "Gido",
		},
		{
			Id:     "120410",
			CityId: "1204",
			Name:   "Idanogawo",
		},
		{
			Id:     "120411",
			CityId: "1204",
			Name:   "Bawolato",
		},
		{
			Id:     "120420",
			CityId: "1204",
			Name:   "Hiliserangkai",
		},
		{
			Id:     "120421",
			CityId: "1204",
			Name:   "Botomuzoi",
		},
		{
			Id:     "120427",
			CityId: "1204",
			Name:   "Ulugawo",
		},
		{
			Id:     "120428",
			CityId: "1204",
			Name:   "Ma'u",
		},
		{
			Id:     "120429",
			CityId: "1204",
			Name:   "Somolo-molo",
		},
		{
			Id:     "120435",
			CityId: "1204",
			Name:   "Sogae'adu",
		},
		{
			Id:     "120501",
			CityId: "1205",
			Name:   "Bahorok",
		},
		{
			Id:     "120502",
			CityId: "1205",
			Name:   "Salapian",
		},
		{
			Id:     "120503",
			CityId: "1205",
			Name:   "Kuala",
		},
		{
			Id:     "120504",
			CityId: "1205",
			Name:   "Sei Bingai",
		},
		{
			Id:     "120505",
			CityId: "1205",
			Name:   "Binjai",
		},
		{
			Id:     "120506",
			CityId: "1205",
			Name:   "Selesai",
		},
		{
			Id:     "120507",
			CityId: "1205",
			Name:   "Stabat",
		},
		{
			Id:     "120508",
			CityId: "1205",
			Name:   "Wampu",
		},
		{
			Id:     "120509",
			CityId: "1205",
			Name:   "Secanggang",
		},
		{
			Id:     "120510",
			CityId: "1205",
			Name:   "Hinai",
		},
		{
			Id:     "120511",
			CityId: "1205",
			Name:   "Tanjung Pura",
		},
		{
			Id:     "120512",
			CityId: "1205",
			Name:   "Padang Tualang",
		},
		{
			Id:     "120513",
			CityId: "1205",
			Name:   "Gebang",
		},
		{
			Id:     "120514",
			CityId: "1205",
			Name:   "Babalan",
		},
		{
			Id:     "120515",
			CityId: "1205",
			Name:   "Pangkalan Susu",
		},
		{
			Id:     "120516",
			CityId: "1205",
			Name:   "Besitang",
		},
		{
			Id:     "120517",
			CityId: "1205",
			Name:   "Sei Lepan",
		},
		{
			Id:     "120518",
			CityId: "1205",
			Name:   "Berandan Barat",
		},
		{
			Id:     "120519",
			CityId: "1205",
			Name:   "Batang Serangan",
		},
		{
			Id:     "120520",
			CityId: "1205",
			Name:   "Sawit Seberang",
		},
		{
			Id:     "120521",
			CityId: "1205",
			Name:   "Sirapit",
		},
		{
			Id:     "120522",
			CityId: "1205",
			Name:   "Kutambaru",
		},
		{
			Id:     "120523",
			CityId: "1205",
			Name:   "Pematang Jaya",
		},
		{
			Id:     "120601",
			CityId: "1206",
			Name:   "Kabanjahe",
		},
		{
			Id:     "120602",
			CityId: "1206",
			Name:   "Berastagi",
		},
		{
			Id:     "120603",
			CityId: "1206",
			Name:   "Barusjahe",
		},
		{
			Id:     "120604",
			CityId: "1206",
			Name:   "Tigapanah",
		},
		{
			Id:     "120605",
			CityId: "1206",
			Name:   "Merek",
		},
		{
			Id:     "120606",
			CityId: "1206",
			Name:   "Munte",
		},
		{
			Id:     "120607",
			CityId: "1206",
			Name:   "Juhar",
		},
		{
			Id:     "120608",
			CityId: "1206",
			Name:   "Tigabinanga",
		},
		{
			Id:     "120609",
			CityId: "1206",
			Name:   "Laubaleng",
		},
		{
			Id:     "120610",
			CityId: "1206",
			Name:   "Mardingding",
		},
		{
			Id:     "120611",
			CityId: "1206",
			Name:   "Payung",
		},
		{
			Id:     "120612",
			CityId: "1206",
			Name:   "Simpang Empat",
		},
		{
			Id:     "120613",
			CityId: "1206",
			Name:   "Kutabuluh",
		},
		{
			Id:     "120614",
			CityId: "1206",
			Name:   "Dolat Rayat",
		},
		{
			Id:     "120615",
			CityId: "1206",
			Name:   "Merdeka",
		},
		{
			Id:     "120616",
			CityId: "1206",
			Name:   "Naman Teran",
		},
		{
			Id:     "120617",
			CityId: "1206",
			Name:   "Tiganderket",
		},
		{
			Id:     "120701",
			CityId: "1207",
			Name:   "Gunung Meriah",
		},
		{
			Id:     "120702",
			CityId: "1207",
			Name:   "Tanjung Morawa",
		},
		{
			Id:     "120703",
			CityId: "1207",
			Name:   "Sibolangit",
		},
		{
			Id:     "120704",
			CityId: "1207",
			Name:   "Kutalimbaru",
		},
		{
			Id:     "120705",
			CityId: "1207",
			Name:   "Pancur Batu",
		},
		{
			Id:     "120706",
			CityId: "1207",
			Name:   "Namo Rambe",
		},
		{
			Id:     "120707",
			CityId: "1207",
			Name:   "Biru-Biru",
		},
		{
			Id:     "120708",
			CityId: "1207",
			Name:   "STM Hilir",
		},
		{
			Id:     "120709",
			CityId: "1207",
			Name:   "Bangun Purba",
		},
		{
			Id:     "120719",
			CityId: "1207",
			Name:   "Galang",
		},
		{
			Id:     "120720",
			CityId: "1207",
			Name:   "STM Hulu",
		},
		{
			Id:     "120721",
			CityId: "1207",
			Name:   "Patumbak",
		},
		{
			Id:     "120722",
			CityId: "1207",
			Name:   "Deli Tua",
		},
		{
			Id:     "120723",
			CityId: "1207",
			Name:   "Sunggal",
		},
		{
			Id:     "120724",
			CityId: "1207",
			Name:   "Hamparan Perak",
		},
		{
			Id:     "120725",
			CityId: "1207",
			Name:   "Labuhan Deli",
		},
		{
			Id:     "120726",
			CityId: "1207",
			Name:   "Percut Sei Tuan",
		},
		{
			Id:     "120727",
			CityId: "1207",
			Name:   "Batang Kuis",
		},
		{
			Id:     "120728",
			CityId: "1207",
			Name:   "Lubuk Pakam",
		},
		{
			Id:     "120731",
			CityId: "1207",
			Name:   "Pagar Merbau",
		},
		{
			Id:     "120732",
			CityId: "1207",
			Name:   "Pantai Labu",
		},
		{
			Id:     "120733",
			CityId: "1207",
			Name:   "Beringin",
		},
		{
			Id:     "120801",
			CityId: "1208",
			Name:   "Siantar",
		},
		{
			Id:     "120802",
			CityId: "1208",
			Name:   "Gunung Malela",
		},
		{
			Id:     "120803",
			CityId: "1208",
			Name:   "Gunung Maligas",
		},
		{
			Id:     "120804",
			CityId: "1208",
			Name:   "Panei",
		},
		{
			Id:     "120805",
			CityId: "1208",
			Name:   "Panombeian Panei",
		},
		{
			Id:     "120806",
			CityId: "1208",
			Name:   "Jorlang Hataran",
		},
		{
			Id:     "120807",
			CityId: "1208",
			Name:   "Raya Kahean",
		},
		{
			Id:     "120808",
			CityId: "1208",
			Name:   "Bosar Maligas",
		},
		{
			Id:     "120809",
			CityId: "1208",
			Name:   "Sidamanik",
		},
		{
			Id:     "120810",
			CityId: "1208",
			Name:   "Pamatang Sidamanik",
		},
		{
			Id:     "120811",
			CityId: "1208",
			Name:   "Tanah Jawa",
		},
		{
			Id:     "120812",
			CityId: "1208",
			Name:   "Hatonduhan",
		},
		{
			Id:     "120813",
			CityId: "1208",
			Name:   "Dolok Panribuan",
		},
		{
			Id:     "120814",
			CityId: "1208",
			Name:   "Purba",
		},
		{
			Id:     "120815",
			CityId: "1208",
			Name:   "Haranggaol Horisan",
		},
		{
			Id:     "120816",
			CityId: "1208",
			Name:   "Girsang Sipangan Bolon",
		},
		{
			Id:     "120817",
			CityId: "1208",
			Name:   "Dolok Batu Nanggar",
		},
		{
			Id:     "120818",
			CityId: "1208",
			Name:   "Huta Bayu Raja",
		},
		{
			Id:     "120819",
			CityId: "1208",
			Name:   "Jawa Maraja Bah Jambi",
		},
		{
			Id:     "120820",
			CityId: "1208",
			Name:   "Dolok Pardamean",
		},
		{
			Id:     "120821",
			CityId: "1208",
			Name:   "Pematang Bandar",
		},
		{
			Id:     "120822",
			CityId: "1208",
			Name:   "Bandar Huluan",
		},
		{
			Id:     "120823",
			CityId: "1208",
			Name:   "Bandar",
		},
		{
			Id:     "120824",
			CityId: "1208",
			Name:   "Bandar Masilam",
		},
		{
			Id:     "120825",
			CityId: "1208",
			Name:   "Silimakuta",
		},
		{
			Id:     "120826",
			CityId: "1208",
			Name:   "Dolok Silou",
		},
		{
			Id:     "120827",
			CityId: "1208",
			Name:   "Silou Kahean",
		},
		{
			Id:     "120828",
			CityId: "1208",
			Name:   "Tapian Dolok",
		},
		{
			Id:     "120829",
			CityId: "1208",
			Name:   "Raya",
		},
		{
			Id:     "120830",
			CityId: "1208",
			Name:   "Ujung Padang",
		},
		{
			Id:     "120831",
			CityId: "1208",
			Name:   "Pamatang Silima Huta",
		},
		{
			Id:     "120832",
			CityId: "1208",
			Name:   "Dolog Masagal",
		},
		{
			Id:     "120908",
			CityId: "1209",
			Name:   "Meranti",
		},
		{
			Id:     "120909",
			CityId: "1209",
			Name:   "Air Joman",
		},
		{
			Id:     "120910",
			CityId: "1209",
			Name:   "Tanjung Balai",
		},
		{
			Id:     "120911",
			CityId: "1209",
			Name:   "Sei Kepayang",
		},
		{
			Id:     "120912",
			CityId: "1209",
			Name:   "Simpang Empat",
		},
		{
			Id:     "120913",
			CityId: "1209",
			Name:   "Air Batu",
		},
		{
			Id:     "120914",
			CityId: "1209",
			Name:   "Pulau Rakyat",
		},
		{
			Id:     "120915",
			CityId: "1209",
			Name:   "Bandar Pulau",
		},
		{
			Id:     "120916",
			CityId: "1209",
			Name:   "Buntu Pane",
		},
		{
			Id:     "120917",
			CityId: "1209",
			Name:   "Bandar Pasir Mandoge",
		},
		{
			Id:     "120918",
			CityId: "1209",
			Name:   "Aek Kuasan",
		},
		{
			Id:     "120919",
			CityId: "1209",
			Name:   "Kota Kisaran Barat",
		},
		{
			Id:     "120920",
			CityId: "1209",
			Name:   "Kota Kisaran Timur",
		},
		{
			Id:     "120921",
			CityId: "1209",
			Name:   "Aek Songsongan",
		},
		{
			Id:     "120922",
			CityId: "1209",
			Name:   "Rahunig",
		},
		{
			Id:     "120923",
			CityId: "1209",
			Name:   "Sei Dadap",
		},
		{
			Id:     "120924",
			CityId: "1209",
			Name:   "Sei Kepayang Barat",
		},
		{
			Id:     "120925",
			CityId: "1209",
			Name:   "Sei Kepayang Timur",
		},
		{
			Id:     "120926",
			CityId: "1209",
			Name:   "Tinggi Raja",
		},
		{
			Id:     "120927",
			CityId: "1209",
			Name:   "Setia Janji",
		},
		{
			Id:     "120928",
			CityId: "1209",
			Name:   "Silau Laut",
		},
		{
			Id:     "120929",
			CityId: "1209",
			Name:   "Rawang Panca Arga",
		},
		{
			Id:     "120930",
			CityId: "1209",
			Name:   "Pulo Bandring",
		},
		{
			Id:     "120931",
			CityId: "1209",
			Name:   "Teluk Dalam",
		},
		{
			Id:     "120932",
			CityId: "1209",
			Name:   "Aek Ledong",
		},
		{
			Id:     "121001",
			CityId: "1210",
			Name:   "Rantau Utara",
		},
		{
			Id:     "121002",
			CityId: "1210",
			Name:   "Rantau Selatan",
		},
		{
			Id:     "121007",
			CityId: "1210",
			Name:   "Bilah Barat",
		},
		{
			Id:     "121008",
			CityId: "1210",
			Name:   "Bilah Hilir",
		},
		{
			Id:     "121009",
			CityId: "1210",
			Name:   "Bilah Hulu",
		},
		{
			Id:     "121014",
			CityId: "1210",
			Name:   "Pangkatan",
		},
		{
			Id:     "121018",
			CityId: "1210",
			Name:   "Panai Tengah",
		},
		{
			Id:     "121019",
			CityId: "1210",
			Name:   "Panai Hilir",
		},
		{
			Id:     "121020",
			CityId: "1210",
			Name:   "Panai Hulu",
		},
		{
			Id:     "121101",
			CityId: "1211",
			Name:   "Sidikalang",
		},
		{
			Id:     "121102",
			CityId: "1211",
			Name:   "Sumbul",
		},
		{
			Id:     "121103",
			CityId: "1211",
			Name:   "Tigalingga",
		},
		{
			Id:     "121104",
			CityId: "1211",
			Name:   "Siempat Nempu",
		},
		{
			Id:     "121105",
			CityId: "1211",
			Name:   "Silima Pungga Pungga",
		},
		{
			Id:     "121106",
			CityId: "1211",
			Name:   "Tanah Pinem",
		},
		{
			Id:     "121107",
			CityId: "1211",
			Name:   "Siempat Nempu Hulu",
		},
		{
			Id:     "121108",
			CityId: "1211",
			Name:   "Siempat Nempu Hilir",
		},
		{
			Id:     "121109",
			CityId: "1211",
			Name:   "Pegagan Hilir",
		},
		{
			Id:     "121110",
			CityId: "1211",
			Name:   "Parbuluan",
		},
		{
			Id:     "121111",
			CityId: "1211",
			Name:   "Lae Parira",
		},
		{
			Id:     "121112",
			CityId: "1211",
			Name:   "Gunung Sitember",
		},
		{
			Id:     "121113",
			CityId: "1211",
			Name:   "Berampu",
		},
		{
			Id:     "121114",
			CityId: "1211",
			Name:   "Silahisabungan",
		},
		{
			Id:     "121115",
			CityId: "1211",
			Name:   "Sitinjo",
		},
		{
			Id:     "121201",
			CityId: "1212",
			Name:   "Balige",
		},
		{
			Id:     "121202",
			CityId: "1212",
			Name:   "Laguboti",
		},
		{
			Id:     "121203",
			CityId: "1212",
			Name:   "Silaen",
		},
		{
			Id:     "121204",
			CityId: "1212",
			Name:   "Habinsaran",
		},
		{
			Id:     "121205",
			CityId: "1212",
			Name:   "Pintu Pohan Meranti",
		},
		{
			Id:     "121206",
			CityId: "1212",
			Name:   "Borbor",
		},
		{
			Id:     "121207",
			CityId: "1212",
			Name:   "Porsea",
		},
		{
			Id:     "121208",
			CityId: "1212",
			Name:   "Ajibata",
		},
		{
			Id:     "121209",
			CityId: "1212",
			Name:   "Lumban Julu",
		},
		{
			Id:     "121210",
			CityId: "1212",
			Name:   "Uluan",
		},
		{
			Id:     "121219",
			CityId: "1212",
			Name:   "Sigumpar",
		},
		{
			Id:     "121220",
			CityId: "1212",
			Name:   "Siantar Narumonda",
		},
		{
			Id:     "121221",
			CityId: "1212",
			Name:   "Nassau",
		},
		{
			Id:     "121222",
			CityId: "1212",
			Name:   "Tampahan",
		},
		{
			Id:     "121223",
			CityId: "1212",
			Name:   "Bonatua Lunasi",
		},
		{
			Id:     "121224",
			CityId: "1212",
			Name:   "Parmaksian",
		},
		{
			Id:     "121301",
			CityId: "1213",
			Name:   "Panyabungan",
		},
		{
			Id:     "121302",
			CityId: "1213",
			Name:   "Panyabungan Utara",
		},
		{
			Id:     "121303",
			CityId: "1213",
			Name:   "Panyabungan Timur",
		},
		{
			Id:     "121304",
			CityId: "1213",
			Name:   "Panyabungan Selatan",
		},
		{
			Id:     "121305",
			CityId: "1213",
			Name:   "Panyabungan Barat",
		},
		{
			Id:     "121306",
			CityId: "1213",
			Name:   "Siabu",
		},
		{
			Id:     "121307",
			CityId: "1213",
			Name:   "Bukit Malintang",
		},
		{
			Id:     "121308",
			CityId: "1213",
			Name:   "Kotanopan",
		},
		{
			Id:     "121309",
			CityId: "1213",
			Name:   "Lembah Sorik Marapi",
		},
		{
			Id:     "121310",
			CityId: "1213",
			Name:   "Tambangan",
		},
		{
			Id:     "121311",
			CityId: "1213",
			Name:   "Ulu Pungkut",
		},
		{
			Id:     "121312",
			CityId: "1213",
			Name:   "Muara Sipongi",
		},
		{
			Id:     "121313",
			CityId: "1213",
			Name:   "Batang Natal",
		},
		{
			Id:     "121314",
			CityId: "1213",
			Name:   "Lingga Bayu",
		},
		{
			Id:     "121315",
			CityId: "1213",
			Name:   "Batahan",
		},
		{
			Id:     "121316",
			CityId: "1213",
			Name:   "Natal",
		},
		{
			Id:     "121317",
			CityId: "1213",
			Name:   "Muara Batang Gadis",
		},
		{
			Id:     "121318",
			CityId: "1213",
			Name:   "Ranto Baek",
		},
		{
			Id:     "121319",
			CityId: "1213",
			Name:   "Huta Bargot",
		},
		{
			Id:     "121320",
			CityId: "1213",
			Name:   "Puncak Sorik Marapi",
		},
		{
			Id:     "121321",
			CityId: "1213",
			Name:   "Pakantan",
		},
		{
			Id:     "121322",
			CityId: "1213",
			Name:   "Sinunukan",
		},
		{
			Id:     "121323",
			CityId: "1213",
			Name:   "Naga Juang",
		},
		{
			Id:     "121401",
			CityId: "1214",
			Name:   "Lolomatua",
		},
		{
			Id:     "121402",
			CityId: "1214",
			Name:   "Gomo",
		},
		{
			Id:     "121403",
			CityId: "1214",
			Name:   "Lahusa",
		},
		{
			Id:     "121404",
			CityId: "1214",
			Name:   "Hibala",
		},
		{
			Id:     "121405",
			CityId: "1214",
			Name:   "Pulau-Pulau Batu",
		},
		{
			Id:     "121406",
			CityId: "1214",
			Name:   "Teluk Dalam",
		},
		{
			Id:     "121407",
			CityId: "1214",
			Name:   "Amandraya",
		},
		{
			Id:     "121408",
			CityId: "1214",
			Name:   "Lolowau",
		},
		{
			Id:     "121409",
			CityId: "1214",
			Name:   "Susua",
		},
		{
			Id:     "121410",
			CityId: "1214",
			Name:   "Maniamolo",
		},
		{
			Id:     "121411",
			CityId: "1214",
			Name:   "Hilimegai",
		},
		{
			Id:     "121412",
			CityId: "1214",
			Name:   "Toma",
		},
		{
			Id:     "121413",
			CityId: "1214",
			Name:   "Mazino",
		},
		{
			Id:     "121414",
			CityId: "1214",
			Name:   "Umbunasi",
		},
		{
			Id:     "121415",
			CityId: "1214",
			Name:   "Aramo",
		},
		{
			Id:     "121416",
			CityId: "1214",
			Name:   "Pulau-Pulau Batu Timur",
		},
		{
			Id:     "121417",
			CityId: "1214",
			Name:   "Mazo",
		},
		{
			Id:     "121418",
			CityId: "1214",
			Name:   "Fanayama",
		},
		{
			Id:     "121419",
			CityId: "1214",
			Name:   "Ulunoyo",
		},
		{
			Id:     "121420",
			CityId: "1214",
			Name:   "Huruna",
		},
		{
			Id:     "121421",
			CityId: "1214",
			Name:   "O'o'u",
		},
		{
			Id:     "121422",
			CityId: "1214",
			Name:   "Onohazumba",
		},
		{
			Id:     "121423",
			CityId: "1214",
			Name:   "Hilisalawa'ahe",
		},
		{
			Id:     "121424",
			CityId: "1214",
			Name:   "Ulususua",
		},
		{
			Id:     "121425",
			CityId: "1214",
			Name:   "Sidua'ori",
		},
		{
			Id:     "121426",
			CityId: "1214",
			Name:   "Somambawa",
		},
		{
			Id:     "121427",
			CityId: "1214",
			Name:   "Boronadu",
		},
		{
			Id:     "121428",
			CityId: "1214",
			Name:   "Simuk",
		},
		{
			Id:     "121429",
			CityId: "1214",
			Name:   "Pulau-Pulau Batu Barat",
		},
		{
			Id:     "121430",
			CityId: "1214",
			Name:   "Pulau-Pulau Batu Utara",
		},
		{
			Id:     "121431",
			CityId: "1214",
			Name:   "Tanah Masa",
		},
		{
			Id:     "121432",
			CityId: "1214",
			Name:   "Luahagundre Maniamolo",
		},
		{
			Id:     "121433",
			CityId: "1214",
			Name:   "Onolalu",
		},
		{
			Id:     "121434",
			CityId: "1214",
			Name:   "Ulu Idanotae",
		},
		{
			Id:     "121435",
			CityId: "1214",
			Name:   "Idanotae",
		},
		{
			Id:     "121501",
			CityId: "1215",
			Name:   "Sitelu Tali Urang Jehe",
		},
		{
			Id:     "121502",
			CityId: "1215",
			Name:   "Kerajaan",
		},
		{
			Id:     "121503",
			CityId: "1215",
			Name:   "Salak",
		},
		{
			Id:     "121504",
			CityId: "1215",
			Name:   "Sitelu Tali Urang Julu",
		},
		{
			Id:     "121505",
			CityId: "1215",
			Name:   "Pergetteng Getteng Sengkut",
		},
		{
			Id:     "121506",
			CityId: "1215",
			Name:   "Pagindar",
		},
		{
			Id:     "121507",
			CityId: "1215",
			Name:   "Tinada",
		},
		{
			Id:     "121508",
			CityId: "1215",
			Name:   "Siempat Rube",
		},
		{
			Id:     "121601",
			CityId: "1216",
			Name:   "Parlilitan",
		},
		{
			Id:     "121602",
			CityId: "1216",
			Name:   "Pollung",
		},
		{
			Id:     "121603",
			CityId: "1216",
			Name:   "Baktiraja",
		},
		{
			Id:     "121604",
			CityId: "1216",
			Name:   "Paranginan",
		},
		{
			Id:     "121605",
			CityId: "1216",
			Name:   "Lintong Nihuta",
		},
		{
			Id:     "121606",
			CityId: "1216",
			Name:   "Dolok Sanggul",
		},
		{
			Id:     "121607",
			CityId: "1216",
			Name:   "Sijamapolang",
		},
		{
			Id:     "121608",
			CityId: "1216",
			Name:   "Onan Ganjang",
		},
		{
			Id:     "121609",
			CityId: "1216",
			Name:   "Pakkat",
		},
		{
			Id:     "121610",
			CityId: "1216",
			Name:   "Tarabintang",
		},
		{
			Id:     "121701",
			CityId: "1217",
			Name:   "Simanindo",
		},
		{
			Id:     "121702",
			CityId: "1217",
			Name:   "Onan Runggu",
		},
		{
			Id:     "121703",
			CityId: "1217",
			Name:   "Nainggolan",
		},
		{
			Id:     "121704",
			CityId: "1217",
			Name:   "Palipi",
		},
		{
			Id:     "121705",
			CityId: "1217",
			Name:   "Harian",
		},
		{
			Id:     "121706",
			CityId: "1217",
			Name:   "Sianjar Mula Mula",
		},
		{
			Id:     "121707",
			CityId: "1217",
			Name:   "Ronggur Nihuta",
		},
		{
			Id:     "121708",
			CityId: "1217",
			Name:   "Pangururan",
		},
		{
			Id:     "121709",
			CityId: "1217",
			Name:   "Sitio-tio",
		},
		{
			Id:     "121801",
			CityId: "1218",
			Name:   "Pantai Cermin",
		},
		{
			Id:     "121802",
			CityId: "1218",
			Name:   "Perbaungan",
		},
		{
			Id:     "121803",
			CityId: "1218",
			Name:   "Teluk Mengkudu",
		},
		{
			Id:     "121804",
			CityId: "1218",
			Name:   "Sei Rampah",
		},
		{
			Id:     "121805",
			CityId: "1218",
			Name:   "Tanjung Beringin",
		},
		{
			Id:     "121806",
			CityId: "1218",
			Name:   "Bandar Khalipah",
		},
		{
			Id:     "121807",
			CityId: "1218",
			Name:   "Dolok Merawan",
		},
		{
			Id:     "121808",
			CityId: "1218",
			Name:   "Sipispis",
		},
		{
			Id:     "121809",
			CityId: "1218",
			Name:   "Dolok Masihul",
		},
		{
			Id:     "121810",
			CityId: "1218",
			Name:   "Kotarih",
		},
		{
			Id:     "121811",
			CityId: "1218",
			Name:   "Silinda",
		},
		{
			Id:     "121812",
			CityId: "1218",
			Name:   "Serba Jadi",
		},
		{
			Id:     "121813",
			CityId: "1218",
			Name:   "Tebing Tinggi",
		},
		{
			Id:     "121814",
			CityId: "1218",
			Name:   "Pegajahan",
		},
		{
			Id:     "121815",
			CityId: "1218",
			Name:   "Sei Bamban",
		},
		{
			Id:     "121816",
			CityId: "1218",
			Name:   "Tebing Syahbandar",
		},
		{
			Id:     "121817",
			CityId: "1218",
			Name:   "Bintang Bayu",
		},
		{
			Id:     "121901",
			CityId: "1219",
			Name:   "Medang Deras",
		},
		{
			Id:     "121902",
			CityId: "1219",
			Name:   "Sei Suka",
		},
		{
			Id:     "121903",
			CityId: "1219",
			Name:   "Air Putih",
		},
		{
			Id:     "121904",
			CityId: "1219",
			Name:   "Lima Puluh",
		},
		{
			Id:     "121905",
			CityId: "1219",
			Name:   "Talawi",
		},
		{
			Id:     "121906",
			CityId: "1219",
			Name:   "Tanjung Tiram",
		},
		{
			Id:     "121907",
			CityId: "1219",
			Name:   "Sei Balai",
		},
		{
			Id:     "121908",
			CityId: "1219",
			Name:   "Laut Tador",
		},
		{
			Id:     "121909",
			CityId: "1219",
			Name:   "Lima Puluh Pesisir",
		},
		{
			Id:     "121910",
			CityId: "1219",
			Name:   "Datuk Lima Puluh",
		},
		{
			Id:     "121911",
			CityId: "1219",
			Name:   "Datuk Tanah Datar",
		},
		{
			Id:     "121912",
			CityId: "1219",
			Name:   "Nibung Hangus",
		},
		{
			Id:     "122001",
			CityId: "1220",
			Name:   "Dolok Sigompulon",
		},
		{
			Id:     "122002",
			CityId: "1220",
			Name:   "Dolok",
		},
		{
			Id:     "122003",
			CityId: "1220",
			Name:   "Halongonan",
		},
		{
			Id:     "122004",
			CityId: "1220",
			Name:   "Padang Bolak",
		},
		{
			Id:     "122005",
			CityId: "1220",
			Name:   "Padang Bolak Julu",
		},
		{
			Id:     "122006",
			CityId: "1220",
			Name:   "Portibi",
		},
		{
			Id:     "122007",
			CityId: "1220",
			Name:   "Batang Onang",
		},
		{
			Id:     "122008",
			CityId: "1220",
			Name:   "Simangambat",
		},
		{
			Id:     "122009",
			CityId: "1220",
			Name:   "Hulu Sihapas",
		},
		{
			Id:     "122010",
			CityId: "1220",
			Name:   "Padang Bolak Tenggara",
		},
		{
			Id:     "122011",
			CityId: "1220",
			Name:   "Halongonan Timur",
		},
		{
			Id:     "122012",
			CityId: "1220",
			Name:   "Ujung Batu",
		},
		{
			Id:     "122101",
			CityId: "1221",
			Name:   "Sosopan",
		},
		{
			Id:     "122102",
			CityId: "1221",
			Name:   "Barumun Tengah",
		},
		{
			Id:     "122103",
			CityId: "1221",
			Name:   "Huristak",
		},
		{
			Id:     "122104",
			CityId: "1221",
			Name:   "Lubuk Barumun",
		},
		{
			Id:     "122105",
			CityId: "1221",
			Name:   "Huta Raja Tinggi",
		},
		{
			Id:     "122106",
			CityId: "1221",
			Name:   "Ulu Barumun",
		},
		{
			Id:     "122107",
			CityId: "1221",
			Name:   "Barumun",
		},
		{
			Id:     "122108",
			CityId: "1221",
			Name:   "Sosa",
		},
		{
			Id:     "122109",
			CityId: "1221",
			Name:   "Batang Lubu Sutam",
		},
		{
			Id:     "122110",
			CityId: "1221",
			Name:   "Barumun Selatan",
		},
		{
			Id:     "122111",
			CityId: "1221",
			Name:   "Aek Nabara Barumun",
		},
		{
			Id:     "122112",
			CityId: "1221",
			Name:   "Sihapas Barumun",
		},
		{
			Id:     "122113",
			CityId: "1221",
			Name:   "Barumun Baru",
		},
		{
			Id:     "122114",
			CityId: "1221",
			Name:   "Ulu Sosa",
		},
		{
			Id:     "122115",
			CityId: "1221",
			Name:   "Sosa Julu",
		},
		{
			Id:     "122116",
			CityId: "1221",
			Name:   "Barumun Barat",
		},
		{
			Id:     "122117",
			CityId: "1221",
			Name:   "Sosa Timur",
		},
		{
			Id:     "122201",
			CityId: "1222",
			Name:   "Kotapinang",
		},
		{
			Id:     "122202",
			CityId: "1222",
			Name:   "Kampung Rakyat",
		},
		{
			Id:     "122203",
			CityId: "1222",
			Name:   "Torgamba",
		},
		{
			Id:     "122204",
			CityId: "1222",
			Name:   "Sungai Kanan",
		},
		{
			Id:     "122205",
			CityId: "1222",
			Name:   "Silangkitang",
		},
		{
			Id:     "122301",
			CityId: "1223",
			Name:   "Kualuh Hulu",
		},
		{
			Id:     "122302",
			CityId: "1223",
			Name:   "Kualuh Leidong",
		},
		{
			Id:     "122303",
			CityId: "1223",
			Name:   "Kualuh Hilir",
		},
		{
			Id:     "122304",
			CityId: "1223",
			Name:   "Aek Kuo",
		},
		{
			Id:     "122305",
			CityId: "1223",
			Name:   "Marbau",
		},
		{
			Id:     "122306",
			CityId: "1223",
			Name:   "Na IX - X",
		},
		{
			Id:     "122307",
			CityId: "1223",
			Name:   "Aek Natas",
		},
		{
			Id:     "122308",
			CityId: "1223",
			Name:   "Kualuh Selatan",
		},
		{
			Id:     "122401",
			CityId: "1224",
			Name:   "Lotu",
		},
		{
			Id:     "122402",
			CityId: "1224",
			Name:   "Sawo",
		},
		{
			Id:     "122403",
			CityId: "1224",
			Name:   "Tuhemberua",
		},
		{
			Id:     "122404",
			CityId: "1224",
			Name:   "Sitolu Ori",
		},
		{
			Id:     "122405",
			CityId: "1224",
			Name:   "Namohalu Esiwa",
		},
		{
			Id:     "122406",
			CityId: "1224",
			Name:   "Alasa Talumuzoi",
		},
		{
			Id:     "122407",
			CityId: "1224",
			Name:   "Alasa",
		},
		{
			Id:     "122408",
			CityId: "1224",
			Name:   "Tugala Oyo",
		},
		{
			Id:     "122409",
			CityId: "1224",
			Name:   "Afulu",
		},
		{
			Id:     "122410",
			CityId: "1224",
			Name:   "Lahewa",
		},
		{
			Id:     "122411",
			CityId: "1224",
			Name:   "Lahewa Timur",
		},
		{
			Id:     "122501",
			CityId: "1225",
			Name:   "Lahomi",
		},
		{
			Id:     "122502",
			CityId: "1225",
			Name:   "Sirombu",
		},
		{
			Id:     "122503",
			CityId: "1225",
			Name:   "Mandrehe Barat",
		},
		{
			Id:     "122504",
			CityId: "1225",
			Name:   "Moro'o",
		},
		{
			Id:     "122505",
			CityId: "1225",
			Name:   "Mandrehe",
		},
		{
			Id:     "122506",
			CityId: "1225",
			Name:   "Mandrehe Utara",
		},
		{
			Id:     "122507",
			CityId: "1225",
			Name:   "Lolofitu Moi",
		},
		{
			Id:     "122508",
			CityId: "1225",
			Name:   "Ulu Moro'o",
		},
		{
			Id:     "127101",
			CityId: "1271",
			Name:   "Medan Kota",
		},
		{
			Id:     "127102",
			CityId: "1271",
			Name:   "Medan Sunggal",
		},
		{
			Id:     "127103",
			CityId: "1271",
			Name:   "Medan Helvetia",
		},
		{
			Id:     "127104",
			CityId: "1271",
			Name:   "Medan Denai",
		},
		{
			Id:     "127105",
			CityId: "1271",
			Name:   "Medan Barat",
		},
		{
			Id:     "127106",
			CityId: "1271",
			Name:   "Medan Deli",
		},
		{
			Id:     "127107",
			CityId: "1271",
			Name:   "Medan Tuntungan",
		},
		{
			Id:     "127108",
			CityId: "1271",
			Name:   "Medan Belawan",
		},
		{
			Id:     "127109",
			CityId: "1271",
			Name:   "Medan Amplas",
		},
		{
			Id:     "127110",
			CityId: "1271",
			Name:   "Medan Area",
		},
		{
			Id:     "127111",
			CityId: "1271",
			Name:   "Medan Johor",
		},
		{
			Id:     "127112",
			CityId: "1271",
			Name:   "Medan Marelan",
		},
		{
			Id:     "127113",
			CityId: "1271",
			Name:   "Medan Labuhan",
		},
		{
			Id:     "127114",
			CityId: "1271",
			Name:   "Medan Tembung",
		},
		{
			Id:     "127115",
			CityId: "1271",
			Name:   "Medan Maimun",
		},
		{
			Id:     "127116",
			CityId: "1271",
			Name:   "Medan Polonia",
		},
		{
			Id:     "127117",
			CityId: "1271",
			Name:   "Medan Baru",
		},
		{
			Id:     "127118",
			CityId: "1271",
			Name:   "Medan Perjuangan",
		},
		{
			Id:     "127119",
			CityId: "1271",
			Name:   "Medan Petisah",
		},
		{
			Id:     "127120",
			CityId: "1271",
			Name:   "Medan Timur",
		},
		{
			Id:     "127121",
			CityId: "1271",
			Name:   "Medan Selayang",
		},
		{
			Id:     "127201",
			CityId: "1272",
			Name:   "Siantar Timur",
		},
		{
			Id:     "127202",
			CityId: "1272",
			Name:   "Siantar Barat",
		},
		{
			Id:     "127203",
			CityId: "1272",
			Name:   "Siantar Utara",
		},
		{
			Id:     "127204",
			CityId: "1272",
			Name:   "Siantar Selatan",
		},
		{
			Id:     "127205",
			CityId: "1272",
			Name:   "Siantar Marihat",
		},
		{
			Id:     "127206",
			CityId: "1272",
			Name:   "Siantar Martoba",
		},
		{
			Id:     "127207",
			CityId: "1272",
			Name:   "Siantar Sitalasari",
		},
		{
			Id:     "127208",
			CityId: "1272",
			Name:   "Siantar Marimbun",
		},
		{
			Id:     "127301",
			CityId: "1273",
			Name:   "Sibolga Utara",
		},
		{
			Id:     "127302",
			CityId: "1273",
			Name:   "Sibolga Kota",
		},
		{
			Id:     "127303",
			CityId: "1273",
			Name:   "Sibolga Selatan",
		},
		{
			Id:     "127304",
			CityId: "1273",
			Name:   "Sibolga Sambas",
		},
		{
			Id:     "127401",
			CityId: "1274",
			Name:   "Tanjungbalai Selatan",
		},
		{
			Id:     "127402",
			CityId: "1274",
			Name:   "Tanjungbalai Utara",
		},
		{
			Id:     "127403",
			CityId: "1274",
			Name:   "Sei Tualang Raso",
		},
		{
			Id:     "127404",
			CityId: "1274",
			Name:   "Teluk Nibung",
		},
		{
			Id:     "127405",
			CityId: "1274",
			Name:   "Datuk Bandar",
		},
		{
			Id:     "127406",
			CityId: "1274",
			Name:   "Datuk Bandar Timur",
		},
		{
			Id:     "127501",
			CityId: "1275",
			Name:   "Binjai Utara",
		},
		{
			Id:     "127502",
			CityId: "1275",
			Name:   "Binjai Kota",
		},
		{
			Id:     "127503",
			CityId: "1275",
			Name:   "Binjai Barat",
		},
		{
			Id:     "127504",
			CityId: "1275",
			Name:   "Binjai Timur",
		},
		{
			Id:     "127505",
			CityId: "1275",
			Name:   "Binjai Selatan",
		},
		{
			Id:     "127601",
			CityId: "1276",
			Name:   "Padang Hulu",
		},
		{
			Id:     "127602",
			CityId: "1276",
			Name:   "Rambutan",
		},
		{
			Id:     "127603",
			CityId: "1276",
			Name:   "Padang Hilir",
		},
		{
			Id:     "127604",
			CityId: "1276",
			Name:   "Bajenis",
		},
		{
			Id:     "127605",
			CityId: "1276",
			Name:   "Tebing Tinggi Kota",
		},
		{
			Id:     "127701",
			CityId: "1277",
			Name:   "Padangsidimpuan Utara",
		},
		{
			Id:     "127702",
			CityId: "1277",
			Name:   "Padangsidimpuan Selatan",
		},
		{
			Id:     "127703",
			CityId: "1277",
			Name:   "Padangsidimpuan Batunadua",
		},
		{
			Id:     "127704",
			CityId: "1277",
			Name:   "Padangsidimpuan Hutaimbaru",
		},
		{
			Id:     "127705",
			CityId: "1277",
			Name:   "Padangsidimpuan Tenggara",
		},
		{
			Id:     "127706",
			CityId: "1277",
			Name:   "Padangsidimpuan Angkola Julu",
		},
		{
			Id:     "127801",
			CityId: "1278",
			Name:   "Gunungsitoli",
		},
		{
			Id:     "127802",
			CityId: "1278",
			Name:   "Gunungsitoli Selatan",
		},
		{
			Id:     "127803",
			CityId: "1278",
			Name:   "Gunungsitoli Utara",
		},
		{
			Id:     "127804",
			CityId: "1278",
			Name:   "Gunungsitoli Idanoi",
		},
		{
			Id:     "127805",
			CityId: "1278",
			Name:   "Gunungsitoli Alo'oa",
		},
		{
			Id:     "127806",
			CityId: "1278",
			Name:   "Gunungsitoli Barat",
		},
		{
			Id:     "130101",
			CityId: "1301",
			Name:   "Pancung Soal",
		},
		{
			Id:     "130102",
			CityId: "1301",
			Name:   "Ranah Pesisir",
		},
		{
			Id:     "130103",
			CityId: "1301",
			Name:   "Lengayang",
		},
		{
			Id:     "130104",
			CityId: "1301",
			Name:   "Batang Kapas",
		},
		{
			Id:     "130105",
			CityId: "1301",
			Name:   "IV Jurai",
		},
		{
			Id:     "130106",
			CityId: "1301",
			Name:   "Bayang",
		},
		{
			Id:     "130107",
			CityId: "1301",
			Name:   "Koto XI Tarusan",
		},
		{
			Id:     "130108",
			CityId: "1301",
			Name:   "Sutera",
		},
		{
			Id:     "130109",
			CityId: "1301",
			Name:   "Linggo Sari Baganti",
		},
		{
			Id:     "130110",
			CityId: "1301",
			Name:   "Lunang",
		},
		{
			Id:     "130111",
			CityId: "1301",
			Name:   "Basa Ampek Balai Tapan",
		},
		{
			Id:     "130112",
			CityId: "1301",
			Name:   "IV Nagari Bayang Utara",
		},
		{
			Id:     "130113",
			CityId: "1301",
			Name:   "Airpura",
		},
		{
			Id:     "130114",
			CityId: "1301",
			Name:   "Ranah Ampek Hulu Tapan",
		},
		{
			Id:     "130115",
			CityId: "1301",
			Name:   "Silaut",
		},
		{
			Id:     "130203",
			CityId: "1302",
			Name:   "Pantai Cermin",
		},
		{
			Id:     "130204",
			CityId: "1302",
			Name:   "Lembah Gumanti",
		},
		{
			Id:     "130205",
			CityId: "1302",
			Name:   "Payung Sekaki",
		},
		{
			Id:     "130206",
			CityId: "1302",
			Name:   "Lembang Jaya",
		},
		{
			Id:     "130207",
			CityId: "1302",
			Name:   "Gunung Talang",
		},
		{
			Id:     "130208",
			CityId: "1302",
			Name:   "Bukit Sundi",
		},
		{
			Id:     "130209",
			CityId: "1302",
			Name:   "IX Koto Sungai Lasi",
		},
		{
			Id:     "130210",
			CityId: "1302",
			Name:   "Kubung",
		},
		{
			Id:     "130211",
			CityId: "1302",
			Name:   "X Koto Singkarak",
		},
		{
			Id:     "130212",
			CityId: "1302",
			Name:   "X Koto Diatas",
		},
		{
			Id:     "130213",
			CityId: "1302",
			Name:   "Junjung Sirih",
		},
		{
			Id:     "130217",
			CityId: "1302",
			Name:   "Hiliran Gumanti",
		},
		{
			Id:     "130218",
			CityId: "1302",
			Name:   "Tigo Lurah",
		},
		{
			Id:     "130219",
			CityId: "1302",
			Name:   "Danau Kembar",
		},
		{
			Id:     "130303",
			CityId: "1303",
			Name:   "Tanjung Gadang",
		},
		{
			Id:     "130304",
			CityId: "1303",
			Name:   "Sijunjung",
		},
		{
			Id:     "130305",
			CityId: "1303",
			Name:   "IV Nagari",
		},
		{
			Id:     "130306",
			CityId: "1303",
			Name:   "Kamang Baru",
		},
		{
			Id:     "130307",
			CityId: "1303",
			Name:   "Lubuak Tarok",
		},
		{
			Id:     "130308",
			CityId: "1303",
			Name:   "Koto VII",
		},
		{
			Id:     "130309",
			CityId: "1303",
			Name:   "Sumpur Kudus",
		},
		{
			Id:     "130310",
			CityId: "1303",
			Name:   "Kupitan",
		},
		{
			Id:     "130401",
			CityId: "1304",
			Name:   "X Koto",
		},
		{
			Id:     "130402",
			CityId: "1304",
			Name:   "Batipuh",
		},
		{
			Id:     "130403",
			CityId: "1304",
			Name:   "Rambatan",
		},
		{
			Id:     "130404",
			CityId: "1304",
			Name:   "Lima Kaum",
		},
		{
			Id:     "130405",
			CityId: "1304",
			Name:   "Tanjung Emas",
		},
		{
			Id:     "130406",
			CityId: "1304",
			Name:   "Lintau Buo",
		},
		{
			Id:     "130407",
			CityId: "1304",
			Name:   "Sungayang",
		},
		{
			Id:     "130408",
			CityId: "1304",
			Name:   "Sungai Tarab",
		},
		{
			Id:     "130409",
			CityId: "1304",
			Name:   "Pariangan",
		},
		{
			Id:     "130410",
			CityId: "1304",
			Name:   "Salimpaung",
		},
		{
			Id:     "130411",
			CityId: "1304",
			Name:   "Padang Ganting",
		},
		{
			Id:     "130412",
			CityId: "1304",
			Name:   "Tanjuang Baru",
		},
		{
			Id:     "130413",
			CityId: "1304",
			Name:   "Lintau Buo Utara",
		},
		{
			Id:     "130414",
			CityId: "1304",
			Name:   "Batipuah Selatan",
		},
		{
			Id:     "130501",
			CityId: "1305",
			Name:   "Lubuk Alung",
		},
		{
			Id:     "130502",
			CityId: "1305",
			Name:   "Batang Anai",
		},
		{
			Id:     "130503",
			CityId: "1305",
			Name:   "Nan Sabaris",
		},
		{
			Id:     "130504",
			CityId: "1305",
			Name:   "2 x 11 Enam Lingkuang",
		},
		{
			Id:     "130505",
			CityId: "1305",
			Name:   "VII Koto Sungai Sarik",
		},
		{
			Id:     "130506",
			CityId: "1305",
			Name:   "V Koto Kampung Dalam",
		},
		{
			Id:     "130507",
			CityId: "1305",
			Name:   "Sungai Garingging",
		},
		{
			Id:     "130508",
			CityId: "1305",
			Name:   "Sungai Limau",
		},
		{
			Id:     "130509",
			CityId: "1305",
			Name:   "IV Koto Aur Malintang",
		},
		{
			Id:     "130510",
			CityId: "1305",
			Name:   "Ulakan Tapakih",
		},
		{
			Id:     "130511",
			CityId: "1305",
			Name:   "Sintuak Toboh Gadang",
		},
		{
			Id:     "130512",
			CityId: "1305",
			Name:   "Padang Sago",
		},
		{
			Id:     "130513",
			CityId: "1305",
			Name:   "Batang Gasan",
		},
		{
			Id:     "130514",
			CityId: "1305",
			Name:   "V Koto Timur",
		},
		{
			Id:     "130515",
			CityId: "1305",
			Name:   "2 x 11 Kayu Tanam",
		},
		{
			Id:     "130516",
			CityId: "1305",
			Name:   "Patamuan",
		},
		{
			Id:     "130517",
			CityId: "1305",
			Name:   "Enam Lingkung",
		},
		{
			Id:     "130601",
			CityId: "1306",
			Name:   "Tanjung Mutiara",
		},
		{
			Id:     "130602",
			CityId: "1306",
			Name:   "Lubuk Basung",
		},
		{
			Id:     "130603",
			CityId: "1306",
			Name:   "Tanjung Raya",
		},
		{
			Id:     "130604",
			CityId: "1306",
			Name:   "Matur",
		},
		{
			Id:     "130605",
			CityId: "1306",
			Name:   "IV Koto",
		},
		{
			Id:     "130606",
			CityId: "1306",
			Name:   "Banuhampu",
		},
		{
			Id:     "130607",
			CityId: "1306",
			Name:   "Ampek Angkek",
		},
		{
			Id:     "130608",
			CityId: "1306",
			Name:   "Baso",
		},
		{
			Id:     "130609",
			CityId: "1306",
			Name:   "Tilatang Kamang",
		},
		{
			Id:     "130610",
			CityId: "1306",
			Name:   "Palupuh",
		},
		{
			Id:     "130611",
			CityId: "1306",
			Name:   "Palembayan",
		},
		{
			Id:     "130612",
			CityId: "1306",
			Name:   "Sungai Pua",
		},
		{
			Id:     "130613",
			CityId: "1306",
			Name:   "Ampek Nagari",
		},
		{
			Id:     "130614",
			CityId: "1306",
			Name:   "Candung",
		},
		{
			Id:     "130615",
			CityId: "1306",
			Name:   "Kamang Magek",
		},
		{
			Id:     "130616",
			CityId: "1306",
			Name:   "Malalak",
		},
		{
			Id:     "130701",
			CityId: "1307",
			Name:   "Suliki",
		},
		{
			Id:     "130702",
			CityId: "1307",
			Name:   "Guguak",
		},
		{
			Id:     "130703",
			CityId: "1307",
			Name:   "Payakumbuh",
		},
		{
			Id:     "130704",
			CityId: "1307",
			Name:   "Luak",
		},
		{
			Id:     "130705",
			CityId: "1307",
			Name:   "Harau",
		},
		{
			Id:     "130706",
			CityId: "1307",
			Name:   "Pangkalan Koto Baru",
		},
		{
			Id:     "130707",
			CityId: "1307",
			Name:   "Kapur IX",
		},
		{
			Id:     "130708",
			CityId: "1307",
			Name:   "Gunuang Omeh",
		},
		{
			Id:     "130709",
			CityId: "1307",
			Name:   "Lareh Sago Halaban",
		},
		{
			Id:     "130710",
			CityId: "1307",
			Name:   "Situjuah Limo Nagari",
		},
		{
			Id:     "130711",
			CityId: "1307",
			Name:   "Mungka",
		},
		{
			Id:     "130712",
			CityId: "1307",
			Name:   "Bukik Barisan",
		},
		{
			Id:     "130713",
			CityId: "1307",
			Name:   "Akabiluru",
		},
		{
			Id:     "130804",
			CityId: "1308",
			Name:   "Bonjol",
		},
		{
			Id:     "130805",
			CityId: "1308",
			Name:   "Lubuk Sikaping",
		},
		{
			Id:     "130807",
			CityId: "1308",
			Name:   "Panti",
		},
		{
			Id:     "130808",
			CityId: "1308",
			Name:   "Mapat Tunggul",
		},
		{
			Id:     "130812",
			CityId: "1308",
			Name:   "Duo Koto",
		},
		{
			Id:     "130813",
			CityId: "1308",
			Name:   "Tigo Nagari",
		},
		{
			Id:     "130814",
			CityId: "1308",
			Name:   "Rao",
		},
		{
			Id:     "130815",
			CityId: "1308",
			Name:   "Mapat Tunggul Selatan",
		},
		{
			Id:     "130816",
			CityId: "1308",
			Name:   "Simpang Alahan Mati",
		},
		{
			Id:     "130817",
			CityId: "1308",
			Name:   "Padang Gelugur",
		},
		{
			Id:     "130818",
			CityId: "1308",
			Name:   "Rao Utara",
		},
		{
			Id:     "130819",
			CityId: "1308",
			Name:   "Rao Selatan",
		},
		{
			Id:     "130901",
			CityId: "1309",
			Name:   "Pagai Utara",
		},
		{
			Id:     "130902",
			CityId: "1309",
			Name:   "Sipora Selatan",
		},
		{
			Id:     "130903",
			CityId: "1309",
			Name:   "Siberut Selatan",
		},
		{
			Id:     "130904",
			CityId: "1309",
			Name:   "Siberut Utara",
		},
		{
			Id:     "130905",
			CityId: "1309",
			Name:   "Siberut Barat",
		},
		{
			Id:     "130906",
			CityId: "1309",
			Name:   "Siberut Barat Daya",
		},
		{
			Id:     "130907",
			CityId: "1309",
			Name:   "Siberut Tengah",
		},
		{
			Id:     "130908",
			CityId: "1309",
			Name:   "Sipora Utara",
		},
		{
			Id:     "130909",
			CityId: "1309",
			Name:   "Sikakap",
		},
		{
			Id:     "130910",
			CityId: "1309",
			Name:   "Pagai Selatan",
		},
		{
			Id:     "131001",
			CityId: "1310",
			Name:   "Koto Baru",
		},
		{
			Id:     "131002",
			CityId: "1310",
			Name:   "Pulau Punjung",
		},
		{
			Id:     "131003",
			CityId: "1310",
			Name:   "Sungai Rumbai",
		},
		{
			Id:     "131004",
			CityId: "1310",
			Name:   "Sitiung",
		},
		{
			Id:     "131005",
			CityId: "1310",
			Name:   "Sembilan Koto",
		},
		{
			Id:     "131006",
			CityId: "1310",
			Name:   "Timpeh",
		},
		{
			Id:     "131007",
			CityId: "1310",
			Name:   "Koto Salak",
		},
		{
			Id:     "131008",
			CityId: "1310",
			Name:   "Tiumang",
		},
		{
			Id:     "131009",
			CityId: "1310",
			Name:   "Padang Laweh",
		},
		{
			Id:     "131010",
			CityId: "1310",
			Name:   "Asam Jujuhan",
		},
		{
			Id:     "131011",
			CityId: "1310",
			Name:   "Koto Besar",
		},
		{
			Id:     "131101",
			CityId: "1311",
			Name:   "Sangir",
		},
		{
			Id:     "131102",
			CityId: "1311",
			Name:   "Sungai Pagu",
		},
		{
			Id:     "131103",
			CityId: "1311",
			Name:   "Koto Parik Gadang Diateh",
		},
		{
			Id:     "131104",
			CityId: "1311",
			Name:   "Sangir Jujuan",
		},
		{
			Id:     "131105",
			CityId: "1311",
			Name:   "Sangir Batang Hari",
		},
		{
			Id:     "131106",
			CityId: "1311",
			Name:   "Pauh Duo",
		},
		{
			Id:     "131107",
			CityId: "1311",
			Name:   "Sangir Balai Janggo",
		},
		{
			Id:     "131201",
			CityId: "1312",
			Name:   "Sungaiberemas",
		},
		{
			Id:     "131202",
			CityId: "1312",
			Name:   "Lembah Melintang",
		},
		{
			Id:     "131203",
			CityId: "1312",
			Name:   "Pasaman",
		},
		{
			Id:     "131204",
			CityId: "1312",
			Name:   "Talamau",
		},
		{
			Id:     "131205",
			CityId: "1312",
			Name:   "Kinali",
		},
		{
			Id:     "131206",
			CityId: "1312",
			Name:   "Gunungtuleh",
		},
		{
			Id:     "131207",
			CityId: "1312",
			Name:   "Ranah Batahan",
		},
		{
			Id:     "131208",
			CityId: "1312",
			Name:   "Koto Balingka",
		},
		{
			Id:     "131209",
			CityId: "1312",
			Name:   "Sungaiaur",
		},
		{
			Id:     "131210",
			CityId: "1312",
			Name:   "Luhak Nan Duo",
		},
		{
			Id:     "131211",
			CityId: "1312",
			Name:   "Sasak Ranah Pesisir",
		},
		{
			Id:     "137101",
			CityId: "1371",
			Name:   "Padang Selatan",
		},
		{
			Id:     "137102",
			CityId: "1371",
			Name:   "Padang Timur",
		},
		{
			Id:     "137103",
			CityId: "1371",
			Name:   "Padang Barat",
		},
		{
			Id:     "137104",
			CityId: "1371",
			Name:   "Padang Utara",
		},
		{
			Id:     "137105",
			CityId: "1371",
			Name:   "Bungus Teluk Kabung",
		},
		{
			Id:     "137106",
			CityId: "1371",
			Name:   "Lubuk Begalung",
		},
		{
			Id:     "137107",
			CityId: "1371",
			Name:   "Lubuk Kilangan",
		},
		{
			Id:     "137108",
			CityId: "1371",
			Name:   "Pauh",
		},
		{
			Id:     "137109",
			CityId: "1371",
			Name:   "Kuranji",
		},
		{
			Id:     "137110",
			CityId: "1371",
			Name:   "Nanggalo",
		},
		{
			Id:     "137111",
			CityId: "1371",
			Name:   "Koto Tangah",
		},
		{
			Id:     "137201",
			CityId: "1372",
			Name:   "Lubuk Sikarah",
		},
		{
			Id:     "137202",
			CityId: "1372",
			Name:   "Tanjung Harapan",
		},
		{
			Id:     "137301",
			CityId: "1373",
			Name:   "Lembah Segar",
		},
		{
			Id:     "137302",
			CityId: "1373",
			Name:   "Barangin",
		},
		{
			Id:     "137303",
			CityId: "1373",
			Name:   "Silungkang",
		},
		{
			Id:     "137304",
			CityId: "1373",
			Name:   "Talawi",
		},
		{
			Id:     "137401",
			CityId: "1374",
			Name:   "Padang Panjang Timur",
		},
		{
			Id:     "137402",
			CityId: "1374",
			Name:   "Padang Panjang Barat",
		},
		{
			Id:     "137501",
			CityId: "1375",
			Name:   "Guguak Panjang",
		},
		{
			Id:     "137502",
			CityId: "1375",
			Name:   "Mandiangin Koto Selayan",
		},
		{
			Id:     "137503",
			CityId: "1375",
			Name:   "Aur Birugo Tigo Baleh",
		},
		{
			Id:     "137601",
			CityId: "1376",
			Name:   "Payakumbuh Barat",
		},
		{
			Id:     "137602",
			CityId: "1376",
			Name:   "Payakumbuh Utara",
		},
		{
			Id:     "137603",
			CityId: "1376",
			Name:   "Payakumbuh Timur",
		},
		{
			Id:     "137604",
			CityId: "1376",
			Name:   "Lamposi Tigo Nagori",
		},
		{
			Id:     "137605",
			CityId: "1376",
			Name:   "Payakumbuh Selatan",
		},
		{
			Id:     "137701",
			CityId: "1377",
			Name:   "Pariaman Tengah",
		},
		{
			Id:     "137702",
			CityId: "1377",
			Name:   "Pariaman Utara",
		},
		{
			Id:     "137703",
			CityId: "1377",
			Name:   "Pariaman Selatan",
		},
		{
			Id:     "137704",
			CityId: "1377",
			Name:   "Pariaman Timur",
		},
		{
			Id:     "140101",
			CityId: "1401",
			Name:   "Bangkinang Kota",
		},
		{
			Id:     "140102",
			CityId: "1401",
			Name:   "Kampar",
		},
		{
			Id:     "140103",
			CityId: "1401",
			Name:   "Tambang",
		},
		{
			Id:     "140104",
			CityId: "1401",
			Name:   "XIII Koto Kampar",
		},
		{
			Id:     "140105",
			CityId: "1401",
			Name:   "Kuok",
		},
		{
			Id:     "140106",
			CityId: "1401",
			Name:   "Siak Hulu",
		},
		{
			Id:     "140107",
			CityId: "1401",
			Name:   "Kampar Kiri",
		},
		{
			Id:     "140108",
			CityId: "1401",
			Name:   "Kampar Kiri Hilir",
		},
		{
			Id:     "140109",
			CityId: "1401",
			Name:   "Kampar Kiri Hulu",
		},
		{
			Id:     "140110",
			CityId: "1401",
			Name:   "Tapung",
		},
		{
			Id:     "140111",
			CityId: "1401",
			Name:   "Tapung Hilir",
		},
		{
			Id:     "140112",
			CityId: "1401",
			Name:   "Tapung Hulu",
		},
		{
			Id:     "140113",
			CityId: "1401",
			Name:   "Salo",
		},
		{
			Id:     "140114",
			CityId: "1401",
			Name:   "Rumbio Jaya",
		},
		{
			Id:     "140115",
			CityId: "1401",
			Name:   "Bangkinang",
		},
		{
			Id:     "140116",
			CityId: "1401",
			Name:   "Perhentian Raja",
		},
		{
			Id:     "140117",
			CityId: "1401",
			Name:   "Kampa",
		},
		{
			Id:     "140118",
			CityId: "1401",
			Name:   "Kampar Utara",
		},
		{
			Id:     "140119",
			CityId: "1401",
			Name:   "Kampar Kiri Tengah",
		},
		{
			Id:     "140120",
			CityId: "1401",
			Name:   "Gunung Sahilan",
		},
		{
			Id:     "140121",
			CityId: "1401",
			Name:   "Koto Kampar Hulu",
		},
		{
			Id:     "140201",
			CityId: "1402",
			Name:   "Rengat",
		},
		{
			Id:     "140202",
			CityId: "1402",
			Name:   "Rengat Barat",
		},
		{
			Id:     "140203",
			CityId: "1402",
			Name:   "Kelayang",
		},
		{
			Id:     "140204",
			CityId: "1402",
			Name:   "Pasir Penyu",
		},
		{
			Id:     "140205",
			CityId: "1402",
			Name:   "Peranap",
		},
		{
			Id:     "140206",
			CityId: "1402",
			Name:   "Siberida",
		},
		{
			Id:     "140207",
			CityId: "1402",
			Name:   "Batang Cenaku",
		},
		{
			Id:     "140208",
			CityId: "1402",
			Name:   "Batang Gangsal",
		},
		{
			Id:     "140209",
			CityId: "1402",
			Name:   "Lirik",
		},
		{
			Id:     "140210",
			CityId: "1402",
			Name:   "Kuala Cenaku",
		},
		{
			Id:     "140211",
			CityId: "1402",
			Name:   "Sungai Lala",
		},
		{
			Id:     "140212",
			CityId: "1402",
			Name:   "Lubuk Batu Jaya",
		},
		{
			Id:     "140213",
			CityId: "1402",
			Name:   "Rakit Kulim",
		},
		{
			Id:     "140214",
			CityId: "1402",
			Name:   "Batang Peranap",
		},
		{
			Id:     "140301",
			CityId: "1403",
			Name:   "Bengkalis",
		},
		{
			Id:     "140302",
			CityId: "1403",
			Name:   "Bantan",
		},
		{
			Id:     "140303",
			CityId: "1403",
			Name:   "Bukit Batu",
		},
		{
			Id:     "140309",
			CityId: "1403",
			Name:   "Mandau",
		},
		{
			Id:     "140310",
			CityId: "1403",
			Name:   "Rupat",
		},
		{
			Id:     "140311",
			CityId: "1403",
			Name:   "Rupat Utara",
		},
		{
			Id:     "140312",
			CityId: "1403",
			Name:   "Siak Kecil",
		},
		{
			Id:     "140313",
			CityId: "1403",
			Name:   "Pinggir",
		},
		{
			Id:     "140314",
			CityId: "1403",
			Name:   "Bandar Laksamana",
		},
		{
			Id:     "140315",
			CityId: "1403",
			Name:   "Talang Muandau",
		},
		{
			Id:     "140316",
			CityId: "1403",
			Name:   "Bathin Solapan",
		},
		{
			Id:     "140401",
			CityId: "1404",
			Name:   "Reteh",
		},
		{
			Id:     "140402",
			CityId: "1404",
			Name:   "Enok",
		},
		{
			Id:     "140403",
			CityId: "1404",
			Name:   "Kuala Indragiri",
		},
		{
			Id:     "140404",
			CityId: "1404",
			Name:   "Tembilahan",
		},
		{
			Id:     "140405",
			CityId: "1404",
			Name:   "Tempuling",
		},
		{
			Id:     "140406",
			CityId: "1404",
			Name:   "Gaung Anak Serka",
		},
		{
			Id:     "140407",
			CityId: "1404",
			Name:   "Mandah",
		},
		{
			Id:     "140408",
			CityId: "1404",
			Name:   "Kateman",
		},
		{
			Id:     "140409",
			CityId: "1404",
			Name:   "Keritang",
		},
		{
			Id:     "140410",
			CityId: "1404",
			Name:   "Tanah Merah",
		},
		{
			Id:     "140411",
			CityId: "1404",
			Name:   "Batang Tuaka",
		},
		{
			Id:     "140412",
			CityId: "1404",
			Name:   "Gaung",
		},
		{
			Id:     "140413",
			CityId: "1404",
			Name:   "Tembilahan Hulu",
		},
		{
			Id:     "140414",
			CityId: "1404",
			Name:   "Kemuning",
		},
		{
			Id:     "140415",
			CityId: "1404",
			Name:   "Pelangiran",
		},
		{
			Id:     "140416",
			CityId: "1404",
			Name:   "Teluk Belengkong",
		},
		{
			Id:     "140417",
			CityId: "1404",
			Name:   "Pulau Burung",
		},
		{
			Id:     "140418",
			CityId: "1404",
			Name:   "Concong",
		},
		{
			Id:     "140419",
			CityId: "1404",
			Name:   "Kempas",
		},
		{
			Id:     "140420",
			CityId: "1404",
			Name:   "Sungai Batang",
		},
		{
			Id:     "140501",
			CityId: "1405",
			Name:   "Ukui",
		},
		{
			Id:     "140502",
			CityId: "1405",
			Name:   "Pangkalan Kerinci",
		},
		{
			Id:     "140503",
			CityId: "1405",
			Name:   "Pangkalan Kuras",
		},
		{
			Id:     "140504",
			CityId: "1405",
			Name:   "Pangkalan Lesung",
		},
		{
			Id:     "140505",
			CityId: "1405",
			Name:   "Langgam",
		},
		{
			Id:     "140506",
			CityId: "1405",
			Name:   "Pelalawan",
		},
		{
			Id:     "140507",
			CityId: "1405",
			Name:   "Kerumutan",
		},
		{
			Id:     "140508",
			CityId: "1405",
			Name:   "Bunut",
		},
		{
			Id:     "140509",
			CityId: "1405",
			Name:   "Teluk Meranti",
		},
		{
			Id:     "140510",
			CityId: "1405",
			Name:   "Kuala Kampar",
		},
		{
			Id:     "140511",
			CityId: "1405",
			Name:   "Bandar Sei Kijang",
		},
		{
			Id:     "140512",
			CityId: "1405",
			Name:   "Bandar Petalangan",
		},
		{
			Id:     "140601",
			CityId: "1406",
			Name:   "Ujung Batu",
		},
		{
			Id:     "140602",
			CityId: "1406",
			Name:   "Rokan IV Koto",
		},
		{
			Id:     "140603",
			CityId: "1406",
			Name:   "Rambah",
		},
		{
			Id:     "140604",
			CityId: "1406",
			Name:   "Tambusai",
		},
		{
			Id:     "140605",
			CityId: "1406",
			Name:   "Kepenuhan",
		},
		{
			Id:     "140606",
			CityId: "1406",
			Name:   "Kunto Darussalam",
		},
		{
			Id:     "140607",
			CityId: "1406",
			Name:   "Rambah Samo",
		},
		{
			Id:     "140608",
			CityId: "1406",
			Name:   "Rambah Hilir",
		},
		{
			Id:     "140609",
			CityId: "1406",
			Name:   "Tambusai Utara",
		},
		{
			Id:     "140610",
			CityId: "1406",
			Name:   "Bangun Purba",
		},
		{
			Id:     "140611",
			CityId: "1406",
			Name:   "Tandun",
		},
		{
			Id:     "140612",
			CityId: "1406",
			Name:   "Kabun",
		},
		{
			Id:     "140613",
			CityId: "1406",
			Name:   "Bonai Darussalam",
		},
		{
			Id:     "140614",
			CityId: "1406",
			Name:   "Pagaran Tapah Darussalam",
		},
		{
			Id:     "140615",
			CityId: "1406",
			Name:   "Kepenuhan Hulu",
		},
		{
			Id:     "140616",
			CityId: "1406",
			Name:   "Pendalian IV Koto",
		},
		{
			Id:     "140701",
			CityId: "1407",
			Name:   "Kubu",
		},
		{
			Id:     "140702",
			CityId: "1407",
			Name:   "Bangko",
		},
		{
			Id:     "140703",
			CityId: "1407",
			Name:   "Tanah Putih",
		},
		{
			Id:     "140704",
			CityId: "1407",
			Name:   "Rimba Melintang",
		},
		{
			Id:     "140705",
			CityId: "1407",
			Name:   "Bagan Sinembah",
		},
		{
			Id:     "140706",
			CityId: "1407",
			Name:   "Pasir Limau Kapas",
		},
		{
			Id:     "140707",
			CityId: "1407",
			Name:   "Sinaboi",
		},
		{
			Id:     "140708",
			CityId: "1407",
			Name:   "Pujud",
		},
		{
			Id:     "140709",
			CityId: "1407",
			Name:   "Tanah Putih Tanjung Melawan",
		},
		{
			Id:     "140710",
			CityId: "1407",
			Name:   "Bangko Pusako",
		},
		{
			Id:     "140711",
			CityId: "1407",
			Name:   "Simpang Kanan",
		},
		{
			Id:     "140712",
			CityId: "1407",
			Name:   "Batu Hampar",
		},
		{
			Id:     "140713",
			CityId: "1407",
			Name:   "Rantau Kopar",
		},
		{
			Id:     "140714",
			CityId: "1407",
			Name:   "Pekaitan",
		},
		{
			Id:     "140715",
			CityId: "1407",
			Name:   "Kubu Babussalam",
		},
		{
			Id:     "140716",
			CityId: "1407",
			Name:   "Tanjung Medan",
		},
		{
			Id:     "140717",
			CityId: "1407",
			Name:   "Bagan Sinembah Raya",
		},
		{
			Id:     "140718",
			CityId: "1407",
			Name:   "Balai Jaya",
		},
		{
			Id:     "140801",
			CityId: "1408",
			Name:   "Siak",
		},
		{
			Id:     "140802",
			CityId: "1408",
			Name:   "Sungai Apit",
		},
		{
			Id:     "140803",
			CityId: "1408",
			Name:   "Minas",
		},
		{
			Id:     "140804",
			CityId: "1408",
			Name:   "Tualang",
		},
		{
			Id:     "140805",
			CityId: "1408",
			Name:   "Sungai Mandau",
		},
		{
			Id:     "140806",
			CityId: "1408",
			Name:   "Dayun",
		},
		{
			Id:     "140807",
			CityId: "1408",
			Name:   "Kerinci Kanan",
		},
		{
			Id:     "140808",
			CityId: "1408",
			Name:   "Bunga Raya",
		},
		{
			Id:     "140809",
			CityId: "1408",
			Name:   "Koto Gasib",
		},
		{
			Id:     "140810",
			CityId: "1408",
			Name:   "Kandis",
		},
		{
			Id:     "140811",
			CityId: "1408",
			Name:   "Lubuk Dalam",
		},
		{
			Id:     "140812",
			CityId: "1408",
			Name:   "Sabak Auh",
		},
		{
			Id:     "140813",
			CityId: "1408",
			Name:   "Mempura",
		},
		{
			Id:     "140814",
			CityId: "1408",
			Name:   "Pusako",
		},
		{
			Id:     "140901",
			CityId: "1409",
			Name:   "Kuantan Mudik",
		},
		{
			Id:     "140902",
			CityId: "1409",
			Name:   "Kuantan Tengah",
		},
		{
			Id:     "140903",
			CityId: "1409",
			Name:   "Singingi",
		},
		{
			Id:     "140904",
			CityId: "1409",
			Name:   "Kuantan Hilir",
		},
		{
			Id:     "140905",
			CityId: "1409",
			Name:   "Cerenti",
		},
		{
			Id:     "140906",
			CityId: "1409",
			Name:   "Benai",
		},
		{
			Id:     "140907",
			CityId: "1409",
			Name:   "Gunungtoar",
		},
		{
			Id:     "140908",
			CityId: "1409",
			Name:   "Singingi Hilir",
		},
		{
			Id:     "140909",
			CityId: "1409",
			Name:   "Pangean",
		},
		{
			Id:     "140910",
			CityId: "1409",
			Name:   "Logas Tanah Darat",
		},
		{
			Id:     "140911",
			CityId: "1409",
			Name:   "Inuman",
		},
		{
			Id:     "140912",
			CityId: "1409",
			Name:   "Hulu Kuantan",
		},
		{
			Id:     "140913",
			CityId: "1409",
			Name:   "Kuantan Hilir Seberang",
		},
		{
			Id:     "140914",
			CityId: "1409",
			Name:   "Sentajo Raya",
		},
		{
			Id:     "140915",
			CityId: "1409",
			Name:   "Pucuk Rantau",
		},
		{
			Id:     "141001",
			CityId: "1410",
			Name:   "Tebing Tinggi",
		},
		{
			Id:     "141002",
			CityId: "1410",
			Name:   "Rangsang Barat",
		},
		{
			Id:     "141003",
			CityId: "1410",
			Name:   "Rangsang",
		},
		{
			Id:     "141004",
			CityId: "1410",
			Name:   "Tebing Tinggi Barat",
		},
		{
			Id:     "141005",
			CityId: "1410",
			Name:   "Merbau",
		},
		{
			Id:     "141006",
			CityId: "1410",
			Name:   "Pulaumerbau",
		},
		{
			Id:     "141007",
			CityId: "1410",
			Name:   "Tebing Tinggi Timur",
		},
		{
			Id:     "141008",
			CityId: "1410",
			Name:   "Tasik Putri Puyu",
		},
		{
			Id:     "141009",
			CityId: "1410",
			Name:   "Rangsang Pesisir",
		},
		{
			Id:     "147101",
			CityId: "1471",
			Name:   "Sukajadi",
		},
		{
			Id:     "147102",
			CityId: "1471",
			Name:   "Pekanbaru Kota",
		},
		{
			Id:     "147103",
			CityId: "1471",
			Name:   "Sail",
		},
		{
			Id:     "147104",
			CityId: "1471",
			Name:   "Lima Puluh",
		},
		{
			Id:     "147105",
			CityId: "1471",
			Name:   "Senapelan",
		},
		{
			Id:     "147106",
			CityId: "1471",
			Name:   "Rumbai Barat",
		},
		{
			Id:     "147107",
			CityId: "1471",
			Name:   "Bukit Raya",
		},
		{
			Id:     "147108",
			CityId: "1471",
			Name:   "Binawidya",
		},
		{
			Id:     "147109",
			CityId: "1471",
			Name:   "Marpoyan Damai",
		},
		{
			Id:     "147110",
			CityId: "1471",
			Name:   "Tenayan Raya",
		},
		{
			Id:     "147111",
			CityId: "1471",
			Name:   "Payung Sekaki",
		},
		{
			Id:     "147112",
			CityId: "1471",
			Name:   "Rumbai",
		},
		{
			Id:     "147113",
			CityId: "1471",
			Name:   "Tuahmadani",
		},
		{
			Id:     "147114",
			CityId: "1471",
			Name:   "Kulim",
		},
		{
			Id:     "147115",
			CityId: "1471",
			Name:   "Rumbai Timur",
		},
		{
			Id:     "147201",
			CityId: "1472",
			Name:   "Dumai Barat",
		},
		{
			Id:     "147202",
			CityId: "1472",
			Name:   "Dumai Timur",
		},
		{
			Id:     "147203",
			CityId: "1472",
			Name:   "Bukit Kapur",
		},
		{
			Id:     "147204",
			CityId: "1472",
			Name:   "Sungai Sembilan",
		},
		{
			Id:     "147205",
			CityId: "1472",
			Name:   "Medang Kampai",
		},
		{
			Id:     "147206",
			CityId: "1472",
			Name:   "Dumai Kota",
		},
		{
			Id:     "147207",
			CityId: "1472",
			Name:   "Dumai Selatan",
		},
		{
			Id:     "150101",
			CityId: "1501",
			Name:   "Gunung Raya",
		},
		{
			Id:     "150102",
			CityId: "1501",
			Name:   "Danau Kerinci",
		},
		{
			Id:     "150104",
			CityId: "1501",
			Name:   "Sitinjau Laut",
		},
		{
			Id:     "150105",
			CityId: "1501",
			Name:   "Air Hangat",
		},
		{
			Id:     "150106",
			CityId: "1501",
			Name:   "Gunung Kerinci",
		},
		{
			Id:     "150107",
			CityId: "1501",
			Name:   "Batang Merangin",
		},
		{
			Id:     "150108",
			CityId: "1501",
			Name:   "Keliling Danau",
		},
		{
			Id:     "150109",
			CityId: "1501",
			Name:   "Kayu Aro",
		},
		{
			Id:     "150111",
			CityId: "1501",
			Name:   "Air Hangat Timur",
		},
		{
			Id:     "150115",
			CityId: "1501",
			Name:   "Gunung Tujuh",
		},
		{
			Id:     "150116",
			CityId: "1501",
			Name:   "Siulak",
		},
		{
			Id:     "150117",
			CityId: "1501",
			Name:   "Depati Tujuh",
		},
		{
			Id:     "150118",
			CityId: "1501",
			Name:   "Siulak Mukai",
		},
		{
			Id:     "150119",
			CityId: "1501",
			Name:   "Kayu Aro Barat",
		},
		{
			Id:     "150120",
			CityId: "1501",
			Name:   "Bukitkerman",
		},
		{
			Id:     "150121",
			CityId: "1501",
			Name:   "Air Hangat Barat",
		},
		{
			Id:     "150122",
			CityId: "1501",
			Name:   "Tanah Cogok",
		},
		{
			Id:     "150123",
			CityId: "1501",
			Name:   "Danau Kerinci Barat",
		},
		{
			Id:     "150201",
			CityId: "1502",
			Name:   "Jangkat",
		},
		{
			Id:     "150202",
			CityId: "1502",
			Name:   "Bangko",
		},
		{
			Id:     "150203",
			CityId: "1502",
			Name:   "Muara Siau",
		},
		{
			Id:     "150204",
			CityId: "1502",
			Name:   "Sungai Manau",
		},
		{
			Id:     "150205",
			CityId: "1502",
			Name:   "Tabir",
		},
		{
			Id:     "150206",
			CityId: "1502",
			Name:   "Pamenang",
		},
		{
			Id:     "150207",
			CityId: "1502",
			Name:   "Tabir Ulu",
		},
		{
			Id:     "150208",
			CityId: "1502",
			Name:   "Tabir Selatan",
		},
		{
			Id:     "150209",
			CityId: "1502",
			Name:   "Lembah Masurai",
		},
		{
			Id:     "150210",
			CityId: "1502",
			Name:   "Bangko Barat",
		},
		{
			Id:     "150211",
			CityId: "1502",
			Name:   "Nalo Tantan",
		},
		{
			Id:     "150212",
			CityId: "1502",
			Name:   "Batang Masumai",
		},
		{
			Id:     "150213",
			CityId: "1502",
			Name:   "Pamenang Barat",
		},
		{
			Id:     "150214",
			CityId: "1502",
			Name:   "Tabir Ilir",
		},
		{
			Id:     "150215",
			CityId: "1502",
			Name:   "Tabir Timur",
		},
		{
			Id:     "150216",
			CityId: "1502",
			Name:   "Renah Pembarap",
		},
		{
			Id:     "150217",
			CityId: "1502",
			Name:   "Pangkalan Jambu",
		},
		{
			Id:     "150218",
			CityId: "1502",
			Name:   "Jangkat Timur",
		},
		{
			Id:     "150219",
			CityId: "1502",
			Name:   "Renah Pamenang",
		},
		{
			Id:     "150220",
			CityId: "1502",
			Name:   "Pamenang Selatan",
		},
		{
			Id:     "150221",
			CityId: "1502",
			Name:   "Margo Tabir",
		},
		{
			Id:     "150222",
			CityId: "1502",
			Name:   "Tabir Lintas",
		},
		{
			Id:     "150223",
			CityId: "1502",
			Name:   "Tabir Barat",
		},
		{
			Id:     "150224",
			CityId: "1502",
			Name:   "Tiang Pumpung",
		},
		{
			Id:     "150301",
			CityId: "1503",
			Name:   "Batang Asai",
		},
		{
			Id:     "150302",
			CityId: "1503",
			Name:   "Limun",
		},
		{
			Id:     "150303",
			CityId: "1503",
			Name:   "Sarolangun",
		},
		{
			Id:     "150304",
			CityId: "1503",
			Name:   "Pauh",
		},
		{
			Id:     "150305",
			CityId: "1503",
			Name:   "Pelawan",
		},
		{
			Id:     "150306",
			CityId: "1503",
			Name:   "Mandiangin",
		},
		{
			Id:     "150307",
			CityId: "1503",
			Name:   "Air Hitam",
		},
		{
			Id:     "150308",
			CityId: "1503",
			Name:   "Bathin VIII",
		},
		{
			Id:     "150309",
			CityId: "1503",
			Name:   "Singkut",
		},
		{
			Id:     "150310",
			CityId: "1503",
			Name:   "Cermin Nan Gedang",
		},
		{
			Id:     "150311",
			CityId: "1503",
			Name:   "Mandiangin Timur",
		},
		{
			Id:     "150401",
			CityId: "1504",
			Name:   "Mersam",
		},
		{
			Id:     "150402",
			CityId: "1504",
			Name:   "Muara Tembesi",
		},
		{
			Id:     "150403",
			CityId: "1504",
			Name:   "Muara Bulian",
		},
		{
			Id:     "150404",
			CityId: "1504",
			Name:   "Batin XXIV",
		},
		{
			Id:     "150405",
			CityId: "1504",
			Name:   "Pemayung",
		},
		{
			Id:     "150406",
			CityId: "1504",
			Name:   "Maro Sebo Ulu",
		},
		{
			Id:     "150407",
			CityId: "1504",
			Name:   "Bajubang",
		},
		{
			Id:     "150408",
			CityId: "1504",
			Name:   "Maro Sebo Ilir",
		},
		{
			Id:     "150501",
			CityId: "1505",
			Name:   "Jambi Luar Kota",
		},
		{
			Id:     "150502",
			CityId: "1505",
			Name:   "Sekernan",
		},
		{
			Id:     "150503",
			CityId: "1505",
			Name:   "Kumpeh",
		},
		{
			Id:     "150504",
			CityId: "1505",
			Name:   "Maro Sebo",
		},
		{
			Id:     "150505",
			CityId: "1505",
			Name:   "Mestong",
		},
		{
			Id:     "150506",
			CityId: "1505",
			Name:   "Kumpeh Ulu",
		},
		{
			Id:     "150507",
			CityId: "1505",
			Name:   "Sungai Bahar",
		},
		{
			Id:     "150508",
			CityId: "1505",
			Name:   "Sungai Gelam",
		},
		{
			Id:     "150509",
			CityId: "1505",
			Name:   "Bahar Utara",
		},
		{
			Id:     "150510",
			CityId: "1505",
			Name:   "Bahar Selatan",
		},
		{
			Id:     "150511",
			CityId: "1505",
			Name:   "Taman Rajo",
		},
		{
			Id:     "150601",
			CityId: "1506",
			Name:   "Tungkal Ulu",
		},
		{
			Id:     "150602",
			CityId: "1506",
			Name:   "Tungkal Ilir",
		},
		{
			Id:     "150603",
			CityId: "1506",
			Name:   "Pengabuan",
		},
		{
			Id:     "150604",
			CityId: "1506",
			Name:   "Betara",
		},
		{
			Id:     "150605",
			CityId: "1506",
			Name:   "Merlung",
		},
		{
			Id:     "150606",
			CityId: "1506",
			Name:   "Tebing Tinggi",
		},
		{
			Id:     "150607",
			CityId: "1506",
			Name:   "Batang Asam",
		},
		{
			Id:     "150608",
			CityId: "1506",
			Name:   "Renah Mendaluh",
		},
		{
			Id:     "150609",
			CityId: "1506",
			Name:   "Muara Papalik",
		},
		{
			Id:     "150610",
			CityId: "1506",
			Name:   "Seberang Kota",
		},
		{
			Id:     "150611",
			CityId: "1506",
			Name:   "Bram Itam",
		},
		{
			Id:     "150612",
			CityId: "1506",
			Name:   "Kuala Betara",
		},
		{
			Id:     "150613",
			CityId: "1506",
			Name:   "Senyerang",
		},
		{
			Id:     "150701",
			CityId: "1507",
			Name:   "Muara Sabak Timur",
		},
		{
			Id:     "150702",
			CityId: "1507",
			Name:   "Nipah Panjang",
		},
		{
			Id:     "150703",
			CityId: "1507",
			Name:   "Mendahara",
		},
		{
			Id:     "150704",
			CityId: "1507",
			Name:   "Rantau Rasau",
		},
		{
			Id:     "150705",
			CityId: "1507",
			Name:   "S a d u",
		},
		{
			Id:     "150706",
			CityId: "1507",
			Name:   "Dendang",
		},
		{
			Id:     "150707",
			CityId: "1507",
			Name:   "Muara Sabak Barat",
		},
		{
			Id:     "150708",
			CityId: "1507",
			Name:   "Kuala Jambi",
		},
		{
			Id:     "150709",
			CityId: "1507",
			Name:   "Mendahara Ulu",
		},
		{
			Id:     "150710",
			CityId: "1507",
			Name:   "Geragai",
		},
		{
			Id:     "150711",
			CityId: "1507",
			Name:   "Berbak",
		},
		{
			Id:     "150801",
			CityId: "1508",
			Name:   "Tanah Tumbuh",
		},
		{
			Id:     "150802",
			CityId: "1508",
			Name:   "Rantau Pandan",
		},
		{
			Id:     "150803",
			CityId: "1508",
			Name:   "Pasar Muaro Bungo",
		},
		{
			Id:     "150804",
			CityId: "1508",
			Name:   "Jujuhan",
		},
		{
			Id:     "150805",
			CityId: "1508",
			Name:   "Tanah Sepenggal",
		},
		{
			Id:     "150806",
			CityId: "1508",
			Name:   "Pelepat",
		},
		{
			Id:     "150807",
			CityId: "1508",
			Name:   "Limbur Lubuk Mengkuang",
		},
		{
			Id:     "150808",
			CityId: "1508",
			Name:   "Muko-muko Bathin VII",
		},
		{
			Id:     "150809",
			CityId: "1508",
			Name:   "Pelepat Ilir",
		},
		{
			Id:     "150810",
			CityId: "1508",
			Name:   "Batin II Babeko",
		},
		{
			Id:     "150811",
			CityId: "1508",
			Name:   "Bathin III",
		},
		{
			Id:     "150812",
			CityId: "1508",
			Name:   "Bungo Dani",
		},
		{
			Id:     "150813",
			CityId: "1508",
			Name:   "Rimbo Tengah",
		},
		{
			Id:     "150814",
			CityId: "1508",
			Name:   "Bathin III Ulu",
		},
		{
			Id:     "150815",
			CityId: "1508",
			Name:   "Bathin II Pelayang",
		},
		{
			Id:     "150816",
			CityId: "1508",
			Name:   "Jujuhan Ilir",
		},
		{
			Id:     "150817",
			CityId: "1508",
			Name:   "Tanah Sepenggal Lintas",
		},
		{
			Id:     "150901",
			CityId: "1509",
			Name:   "Tebo Tengah",
		},
		{
			Id:     "150902",
			CityId: "1509",
			Name:   "Tebo Ilir",
		},
		{
			Id:     "150903",
			CityId: "1509",
			Name:   "Tebo Ulu",
		},
		{
			Id:     "150904",
			CityId: "1509",
			Name:   "Rimbo Bujang",
		},
		{
			Id:     "150905",
			CityId: "1509",
			Name:   "Sumay",
		},
		{
			Id:     "150906",
			CityId: "1509",
			Name:   "VII Koto",
		},
		{
			Id:     "150907",
			CityId: "1509",
			Name:   "Rimbo Ulu",
		},
		{
			Id:     "150908",
			CityId: "1509",
			Name:   "Rimbo Ilir",
		},
		{
			Id:     "150909",
			CityId: "1509",
			Name:   "Tengah Ilir",
		},
		{
			Id:     "150910",
			CityId: "1509",
			Name:   "Serai Serumpun",
		},
		{
			Id:     "150911",
			CityId: "1509",
			Name:   "VII Koto Ilir",
		},
		{
			Id:     "150912",
			CityId: "1509",
			Name:   "Muara Tabir",
		},
		{
			Id:     "157101",
			CityId: "1571",
			Name:   "Telanaipura",
		},
		{
			Id:     "157102",
			CityId: "1571",
			Name:   "Jambi Selatan",
		},
		{
			Id:     "157103",
			CityId: "1571",
			Name:   "Jambi Timur",
		},
		{
			Id:     "157104",
			CityId: "1571",
			Name:   "Pasar Jambi",
		},
		{
			Id:     "157105",
			CityId: "1571",
			Name:   "Pelayangan",
		},
		{
			Id:     "157106",
			CityId: "1571",
			Name:   "Danau Teluk",
		},
		{
			Id:     "157107",
			CityId: "1571",
			Name:   "Kota Baru",
		},
		{
			Id:     "157108",
			CityId: "1571",
			Name:   "Jelutung",
		},
		{
			Id:     "157109",
			CityId: "1571",
			Name:   "Alam Barajo",
		},
		{
			Id:     "157110",
			CityId: "1571",
			Name:   "Danau Sipin",
		},
		{
			Id:     "157111",
			CityId: "1571",
			Name:   "Paal Merah",
		},
		{
			Id:     "157201",
			CityId: "1572",
			Name:   "Sungai Penuh",
		},
		{
			Id:     "157202",
			CityId: "1572",
			Name:   "Pesisir Bukit",
		},
		{
			Id:     "157203",
			CityId: "1572",
			Name:   "Hamparan Rawang",
		},
		{
			Id:     "157204",
			CityId: "1572",
			Name:   "Tanah Kampung",
		},
		{
			Id:     "157205",
			CityId: "1572",
			Name:   "Kumun Debai",
		},
		{
			Id:     "157206",
			CityId: "1572",
			Name:   "Pondok Tinggi",
		},
		{
			Id:     "157207",
			CityId: "1572",
			Name:   "Koto Baru",
		},
		{
			Id:     "157208",
			CityId: "1572",
			Name:   "Sungai Bungkal",
		},
		{
			Id:     "160107",
			CityId: "1601",
			Name:   "Sosoh Buay Rayap",
		},
		{
			Id:     "160108",
			CityId: "1601",
			Name:   "Pengandonan",
		},
		{
			Id:     "160109",
			CityId: "1601",
			Name:   "Peninjauan",
		},
		{
			Id:     "160113",
			CityId: "1601",
			Name:   "Baturaja Barat",
		},
		{
			Id:     "160114",
			CityId: "1601",
			Name:   "Baturaja Timur",
		},
		{
			Id:     "160120",
			CityId: "1601",
			Name:   "Ulu Ogan",
		},
		{
			Id:     "160121",
			CityId: "1601",
			Name:   "Semidang Aji",
		},
		{
			Id:     "160122",
			CityId: "1601",
			Name:   "Lubuk Batang",
		},
		{
			Id:     "160128",
			CityId: "1601",
			Name:   "Lengkiti",
		},
		{
			Id:     "160129",
			CityId: "1601",
			Name:   "Sinar Peninjauan",
		},
		{
			Id:     "160130",
			CityId: "1601",
			Name:   "Lubuk Raja",
		},
		{
			Id:     "160131",
			CityId: "1601",
			Name:   "Muara Jaya",
		},
		{
			Id:     "160132",
			CityId: "1601",
			Name:   "Kedaton Peninjauan Raya",
		},
		{
			Id:     "160202",
			CityId: "1602",
			Name:   "Tanjung Lubuk",
		},
		{
			Id:     "160203",
			CityId: "1602",
			Name:   "Pedamaran",
		},
		{
			Id:     "160204",
			CityId: "1602",
			Name:   "Mesuji",
		},
		{
			Id:     "160205",
			CityId: "1602",
			Name:   "Kayu Agung",
		},
		{
			Id:     "160208",
			CityId: "1602",
			Name:   "Sirah Pulau Padang",
		},
		{
			Id:     "160211",
			CityId: "1602",
			Name:   "Tulung Selapan",
		},
		{
			Id:     "160212",
			CityId: "1602",
			Name:   "Pampangan",
		},
		{
			Id:     "160213",
			CityId: "1602",
			Name:   "Lempuing",
		},
		{
			Id:     "160214",
			CityId: "1602",
			Name:   "Air Sugihan",
		},
		{
			Id:     "160215",
			CityId: "1602",
			Name:   "Sungai Menang",
		},
		{
			Id:     "160217",
			CityId: "1602",
			Name:   "Jejawi",
		},
		{
			Id:     "160218",
			CityId: "1602",
			Name:   "Cengal",
		},
		{
			Id:     "160219",
			CityId: "1602",
			Name:   "Pangkalan Lampam",
		},
		{
			Id:     "160220",
			CityId: "1602",
			Name:   "Mesuji Makmur",
		},
		{
			Id:     "160221",
			CityId: "1602",
			Name:   "Mesuji Raya",
		},
		{
			Id:     "160222",
			CityId: "1602",
			Name:   "Lempuing Jaya",
		},
		{
			Id:     "160223",
			CityId: "1602",
			Name:   "Teluk Gelam",
		},
		{
			Id:     "160224",
			CityId: "1602",
			Name:   "Pedamaran Timur",
		},
		{
			Id:     "160301",
			CityId: "1603",
			Name:   "Tanjung Agung",
		},
		{
			Id:     "160302",
			CityId: "1603",
			Name:   "Muara Enim",
		},
		{
			Id:     "160303",
			CityId: "1603",
			Name:   "Rambang Niru",
		},
		{
			Id:     "160304",
			CityId: "1603",
			Name:   "Gunung Megang",
		},
		{
			Id:     "160306",
			CityId: "1603",
			Name:   "Gelumbang",
		},
		{
			Id:     "160307",
			CityId: "1603",
			Name:   "Lawang Kidul",
		},
		{
			Id:     "160308",
			CityId: "1603",
			Name:   "Semende Darat Laut",
		},
		{
			Id:     "160309",
			CityId: "1603",
			Name:   "Semende Darat Tengah",
		},
		{
			Id:     "160310",
			CityId: "1603",
			Name:   "Semende Darat Ulu",
		},
		{
			Id:     "160311",
			CityId: "1603",
			Name:   "Ujan Mas",
		},
		{
			Id:     "160314",
			CityId: "1603",
			Name:   "Lubai",
		},
		{
			Id:     "160315",
			CityId: "1603",
			Name:   "Rambang",
		},
		{
			Id:     "160316",
			CityId: "1603",
			Name:   "Sungai Rotan",
		},
		{
			Id:     "160317",
			CityId: "1603",
			Name:   "Lembak",
		},
		{
			Id:     "160319",
			CityId: "1603",
			Name:   "Benakat",
		},
		{
			Id:     "160321",
			CityId: "1603",
			Name:   "Kelekar",
		},
		{
			Id:     "160322",
			CityId: "1603",
			Name:   "Muara Belida",
		},
		{
			Id:     "160323",
			CityId: "1603",
			Name:   "Belimbing",
		},
		{
			Id:     "160324",
			CityId: "1603",
			Name:   "Belida Darat",
		},
		{
			Id:     "160325",
			CityId: "1603",
			Name:   "Lubai Ulu",
		},
		{
			Id:     "160326",
			CityId: "1603",
			Name:   "Empat Petulai Dangku",
		},
		{
			Id:     "160327",
			CityId: "1603",
			Name:   "Panang Enim",
		},
		{
			Id:     "160401",
			CityId: "1604",
			Name:   "Tanjungsakti Pumu",
		},
		{
			Id:     "160406",
			CityId: "1604",
			Name:   "Jarai",
		},
		{
			Id:     "160407",
			CityId: "1604",
			Name:   "Kota Agung",
		},
		{
			Id:     "160408",
			CityId: "1604",
			Name:   "Pulaupinang",
		},
		{
			Id:     "160409",
			CityId: "1604",
			Name:   "Merapi Barat",
		},
		{
			Id:     "160410",
			CityId: "1604",
			Name:   "Lahat",
		},
		{
			Id:     "160412",
			CityId: "1604",
			Name:   "Pajar Bulan",
		},
		{
			Id:     "160415",
			CityId: "1604",
			Name:   "Mulak Ulu",
		},
		{
			Id:     "160416",
			CityId: "1604",
			Name:   "Kikim Selatan",
		},
		{
			Id:     "160417",
			CityId: "1604",
			Name:   "Kikim Timur",
		},
		{
			Id:     "160418",
			CityId: "1604",
			Name:   "Kikim Tengah",
		},
		{
			Id:     "160419",
			CityId: "1604",
			Name:   "Kikim Barat",
		},
		{
			Id:     "160420",
			CityId: "1604",
			Name:   "Pseksu",
		},
		{
			Id:     "160421",
			CityId: "1604",
			Name:   "Gumay Talang",
		},
		{
			Id:     "160422",
			CityId: "1604",
			Name:   "Pagar Gunung",
		},
		{
			Id:     "160423",
			CityId: "1604",
			Name:   "Merapi Timur",
		},
		{
			Id:     "160424",
			CityId: "1604",
			Name:   "Tanjung Sakti Pumi",
		},
		{
			Id:     "160425",
			CityId: "1604",
			Name:   "Gumay Ulu",
		},
		{
			Id:     "160426",
			CityId: "1604",
			Name:   "Merapi Selatan",
		},
		{
			Id:     "160427",
			CityId: "1604",
			Name:   "Tanjungtebat",
		},
		{
			Id:     "160428",
			CityId: "1604",
			Name:   "Muarapayang",
		},
		{
			Id:     "160429",
			CityId: "1604",
			Name:   "Sukamerindu",
		},
		{
			Id:     "160430",
			CityId: "1604",
			Name:   "Mulak Sebingkai",
		},
		{
			Id:     "160431",
			CityId: "1604",
			Name:   "Lahat Selatan",
		},
		{
			Id:     "160501",
			CityId: "1605",
			Name:   "Tugumulyo",
		},
		{
			Id:     "160502",
			CityId: "1605",
			Name:   "Muara Lakitan",
		},
		{
			Id:     "160503",
			CityId: "1605",
			Name:   "Muara Kelingi",
		},
		{
			Id:     "160508",
			CityId: "1605",
			Name:   "Jayaloka",
		},
		{
			Id:     "160509",
			CityId: "1605",
			Name:   "Muara Beliti",
		},
		{
			Id:     "160510",
			CityId: "1605",
			Name:   "STL Ulu Terawas",
		},
		{
			Id:     "160511",
			CityId: "1605",
			Name:   "Selangit",
		},
		{
			Id:     "160512",
			CityId: "1605",
			Name:   "Megang Sakti",
		},
		{
			Id:     "160513",
			CityId: "1605",
			Name:   "Purwodadi",
		},
		{
			Id:     "160514",
			CityId: "1605",
			Name:   "BTS. Ulu",
		},
		{
			Id:     "160518",
			CityId: "1605",
			Name:   "Tiang Pumpung Kepungut",
		},
		{
			Id:     "160519",
			CityId: "1605",
			Name:   "Sumber Harta",
		},
		{
			Id:     "160520",
			CityId: "1605",
			Name:   "Tuah Negeri",
		},
		{
			Id:     "160521",
			CityId: "1605",
			Name:   "Suka Karya",
		},
		{
			Id:     "160601",
			CityId: "1606",
			Name:   "Sekayu",
		},
		{
			Id:     "160602",
			CityId: "1606",
			Name:   "Lais",
		},
		{
			Id:     "160603",
			CityId: "1606",
			Name:   "Sungai Keruh",
		},
		{
			Id:     "160604",
			CityId: "1606",
			Name:   "Batang Hari Leko",
		},
		{
			Id:     "160605",
			CityId: "1606",
			Name:   "Sanga Desa",
		},
		{
			Id:     "160606",
			CityId: "1606",
			Name:   "Babat Toman",
		},
		{
			Id:     "160607",
			CityId: "1606",
			Name:   "Sungai Lilin",
		},
		{
			Id:     "160608",
			CityId: "1606",
			Name:   "Keluang",
		},
		{
			Id:     "160609",
			CityId: "1606",
			Name:   "Bayung Lencir",
		},
		{
			Id:     "160610",
			CityId: "1606",
			Name:   "Plakat Tinggi",
		},
		{
			Id:     "160611",
			CityId: "1606",
			Name:   "Lalan",
		},
		{
			Id:     "160612",
			CityId: "1606",
			Name:   "Tungkal Jaya",
		},
		{
			Id:     "160613",
			CityId: "1606",
			Name:   "Lawang Wetan",
		},
		{
			Id:     "160614",
			CityId: "1606",
			Name:   "Babat Supat",
		},
		{
			Id:     "160615",
			CityId: "1606",
			Name:   "Jirak Jaya",
		},
		{
			Id:     "160701",
			CityId: "1607",
			Name:   "Banyuasin I",
		},
		{
			Id:     "160702",
			CityId: "1607",
			Name:   "Banyuasin II",
		},
		{
			Id:     "160703",
			CityId: "1607",
			Name:   "Banyuasin III",
		},
		{
			Id:     "160704",
			CityId: "1607",
			Name:   "Pulau Rimau",
		},
		{
			Id:     "160705",
			CityId: "1607",
			Name:   "Betung",
		},
		{
			Id:     "160706",
			CityId: "1607",
			Name:   "Rambutan",
		},
		{
			Id:     "160707",
			CityId: "1607",
			Name:   "Muara Padang",
		},
		{
			Id:     "160708",
			CityId: "1607",
			Name:   "Muara Telang",
		},
		{
			Id:     "160709",
			CityId: "1607",
			Name:   "Makarti Jaya",
		},
		{
			Id:     "160710",
			CityId: "1607",
			Name:   "Talang Kelapa",
		},
		{
			Id:     "160711",
			CityId: "1607",
			Name:   "Rantau Bayur",
		},
		{
			Id:     "160712",
			CityId: "1607",
			Name:   "Tanjung Lago",
		},
		{
			Id:     "160713",
			CityId: "1607",
			Name:   "Muara Sugihan",
		},
		{
			Id:     "160714",
			CityId: "1607",
			Name:   "Air Salek",
		},
		{
			Id:     "160715",
			CityId: "1607",
			Name:   "Tungkal Ilir",
		},
		{
			Id:     "160716",
			CityId: "1607",
			Name:   "Suak Tapeh",
		},
		{
			Id:     "160717",
			CityId: "1607",
			Name:   "Sembawa",
		},
		{
			Id:     "160718",
			CityId: "1607",
			Name:   "Sumber Marga Telang",
		},
		{
			Id:     "160719",
			CityId: "1607",
			Name:   "Air Kumbang",
		},
		{
			Id:     "160720",
			CityId: "1607",
			Name:   "Karangagung Ilir",
		},
		{
			Id:     "160721",
			CityId: "1607",
			Name:   "Selat Panuguan",
		},
		{
			Id:     "160801",
			CityId: "1608",
			Name:   "Martapura",
		},
		{
			Id:     "160802",
			CityId: "1608",
			Name:   "Buay Madang",
		},
		{
			Id:     "160803",
			CityId: "1608",
			Name:   "Belitang",
		},
		{
			Id:     "160804",
			CityId: "1608",
			Name:   "Cempaka",
		},
		{
			Id:     "160805",
			CityId: "1608",
			Name:   "Buay Pemuka Peliung",
		},
		{
			Id:     "160806",
			CityId: "1608",
			Name:   "Madang Suku II",
		},
		{
			Id:     "160807",
			CityId: "1608",
			Name:   "Madang Suku I",
		},
		{
			Id:     "160808",
			CityId: "1608",
			Name:   "Semendawai Suku III",
		},
		{
			Id:     "160809",
			CityId: "1608",
			Name:   "Belitang II",
		},
		{
			Id:     "160810",
			CityId: "1608",
			Name:   "Belitang III",
		},
		{
			Id:     "160811",
			CityId: "1608",
			Name:   "Bunga Mayang",
		},
		{
			Id:     "160812",
			CityId: "1608",
			Name:   "Buay Madang Timur",
		},
		{
			Id:     "160813",
			CityId: "1608",
			Name:   "Madang Suku III",
		},
		{
			Id:     "160814",
			CityId: "1608",
			Name:   "Semendawai Barat",
		},
		{
			Id:     "160815",
			CityId: "1608",
			Name:   "Semendawai Timur",
		},
		{
			Id:     "160816",
			CityId: "1608",
			Name:   "Jayapura",
		},
		{
			Id:     "160817",
			CityId: "1608",
			Name:   "Belitang Jaya",
		},
		{
			Id:     "160818",
			CityId: "1608",
			Name:   "Belitang Madang Raya",
		},
		{
			Id:     "160819",
			CityId: "1608",
			Name:   "Belitang Mulya",
		},
		{
			Id:     "160820",
			CityId: "1608",
			Name:   "Buay Pemuka Bangsa Raja",
		},
		{
			Id:     "160901",
			CityId: "1609",
			Name:   "Muara Dua",
		},
		{
			Id:     "160902",
			CityId: "1609",
			Name:   "Pulau Beringin",
		},
		{
			Id:     "160903",
			CityId: "1609",
			Name:   "Banding Agung",
		},
		{
			Id:     "160904",
			CityId: "1609",
			Name:   "Muara Dua Kisam",
		},
		{
			Id:     "160905",
			CityId: "1609",
			Name:   "Simpang",
		},
		{
			Id:     "160906",
			CityId: "1609",
			Name:   "Buay Sandang Aji",
		},
		{
			Id:     "160907",
			CityId: "1609",
			Name:   "Buay Runjung",
		},
		{
			Id:     "160908",
			CityId: "1609",
			Name:   "Mekakau Ilir",
		},
		{
			Id:     "160909",
			CityId: "1609",
			Name:   "Buay Pemaca",
		},
		{
			Id:     "160910",
			CityId: "1609",
			Name:   "Kisam Tinggi",
		},
		{
			Id:     "160911",
			CityId: "1609",
			Name:   "Kisam Ilir",
		},
		{
			Id:     "160912",
			CityId: "1609",
			Name:   "Buay Pematang Ribu Ranau Tengah",
		},
		{
			Id:     "160913",
			CityId: "1609",
			Name:   "Warkuk Ranau Selatan",
		},
		{
			Id:     "160914",
			CityId: "1609",
			Name:   "Runjung Agung",
		},
		{
			Id:     "160915",
			CityId: "1609",
			Name:   "Sungai Are",
		},
		{
			Id:     "160916",
			CityId: "1609",
			Name:   "Sindang Danau",
		},
		{
			Id:     "160917",
			CityId: "1609",
			Name:   "Buana Pemaca",
		},
		{
			Id:     "160918",
			CityId: "1609",
			Name:   "Tiga Dihaji",
		},
		{
			Id:     "160919",
			CityId: "1609",
			Name:   "Buay Rawan",
		},
		{
			Id:     "161001",
			CityId: "1610",
			Name:   "Muara Kuang",
		},
		{
			Id:     "161002",
			CityId: "1610",
			Name:   "Tanjung Batu",
		},
		{
			Id:     "161003",
			CityId: "1610",
			Name:   "Tanjung Raja",
		},
		{
			Id:     "161004",
			CityId: "1610",
			Name:   "Indralaya",
		},
		{
			Id:     "161005",
			CityId: "1610",
			Name:   "Pemulutan",
		},
		{
			Id:     "161006",
			CityId: "1610",
			Name:   "Rantau Alai",
		},
		{
			Id:     "161007",
			CityId: "1610",
			Name:   "Indralaya Utara",
		},
		{
			Id:     "161008",
			CityId: "1610",
			Name:   "Indralaya Selatan",
		},
		{
			Id:     "161009",
			CityId: "1610",
			Name:   "Pemulutan Selatan",
		},
		{
			Id:     "161010",
			CityId: "1610",
			Name:   "Pemulutan Barat",
		},
		{
			Id:     "161011",
			CityId: "1610",
			Name:   "Rantau Panjang",
		},
		{
			Id:     "161012",
			CityId: "1610",
			Name:   "Sungai Pinang",
		},
		{
			Id:     "161013",
			CityId: "1610",
			Name:   "Kandis",
		},
		{
			Id:     "161014",
			CityId: "1610",
			Name:   "Rambang Kuang",
		},
		{
			Id:     "161015",
			CityId: "1610",
			Name:   "Lubuk Keliat",
		},
		{
			Id:     "161016",
			CityId: "1610",
			Name:   "Payaraman",
		},
		{
			Id:     "161101",
			CityId: "1611",
			Name:   "Muara Pinang",
		},
		{
			Id:     "161102",
			CityId: "1611",
			Name:   "Pendopo",
		},
		{
			Id:     "161103",
			CityId: "1611",
			Name:   "Ulu Musi",
		},
		{
			Id:     "161104",
			CityId: "1611",
			Name:   "Tebing Tinggi",
		},
		{
			Id:     "161105",
			CityId: "1611",
			Name:   "Lintang Kanan",
		},
		{
			Id:     "161106",
			CityId: "1611",
			Name:   "Talang Padang",
		},
		{
			Id:     "161107",
			CityId: "1611",
			Name:   "Pasemah Air Keruh",
		},
		{
			Id:     "161108",
			CityId: "1611",
			Name:   "Sikap Dalam",
		},
		{
			Id:     "161109",
			CityId: "1611",
			Name:   "Saling",
		},
		{
			Id:     "161110",
			CityId: "1611",
			Name:   "Pendopo Barat",
		},
		{
			Id:     "161201",
			CityId: "1612",
			Name:   "Talang Ubi",
		},
		{
			Id:     "161202",
			CityId: "1612",
			Name:   "Penukal Utara",
		},
		{
			Id:     "161203",
			CityId: "1612",
			Name:   "Penukal",
		},
		{
			Id:     "161204",
			CityId: "1612",
			Name:   "Abab",
		},
		{
			Id:     "161205",
			CityId: "1612",
			Name:   "Tanah Abang",
		},
		{
			Id:     "161301",
			CityId: "1613",
			Name:   "Rupit",
		},
		{
			Id:     "161302",
			CityId: "1613",
			Name:   "Rawas Ulu",
		},
		{
			Id:     "161303",
			CityId: "1613",
			Name:   "Nibung",
		},
		{
			Id:     "161304",
			CityId: "1613",
			Name:   "Rawas Ilir",
		},
		{
			Id:     "161305",
			CityId: "1613",
			Name:   "Karang Dapo",
		},
		{
			Id:     "161306",
			CityId: "1613",
			Name:   "Karang Jaya",
		},
		{
			Id:     "161307",
			CityId: "1613",
			Name:   "Ulu Rawas",
		},
		{
			Id:     "167101",
			CityId: "1671",
			Name:   "Ilir Barat II",
		},
		{
			Id:     "167102",
			CityId: "1671",
			Name:   "Seberang Ulu I",
		},
		{
			Id:     "167103",
			CityId: "1671",
			Name:   "Seberang Ulu II",
		},
		{
			Id:     "167104",
			CityId: "1671",
			Name:   "Ilir Barat I",
		},
		{
			Id:     "167105",
			CityId: "1671",
			Name:   "Ilir Timur I",
		},
		{
			Id:     "167106",
			CityId: "1671",
			Name:   "Ilir Timur II",
		},
		{
			Id:     "167107",
			CityId: "1671",
			Name:   "Sukarami",
		},
		{
			Id:     "167108",
			CityId: "1671",
			Name:   "Sako",
		},
		{
			Id:     "167109",
			CityId: "1671",
			Name:   "Kemuning",
		},
		{
			Id:     "167110",
			CityId: "1671",
			Name:   "Kalidoni",
		},
		{
			Id:     "167111",
			CityId: "1671",
			Name:   "Bukit Kecil",
		},
		{
			Id:     "167112",
			CityId: "1671",
			Name:   "Gandus",
		},
		{
			Id:     "167113",
			CityId: "1671",
			Name:   "Kertapati",
		},
		{
			Id:     "167114",
			CityId: "1671",
			Name:   "Plaju",
		},
		{
			Id:     "167115",
			CityId: "1671",
			Name:   "Alang-alang Lebar",
		},
		{
			Id:     "167116",
			CityId: "1671",
			Name:   "Sematang Borang",
		},
		{
			Id:     "167117",
			CityId: "1671",
			Name:   "Jakabaring",
		},
		{
			Id:     "167118",
			CityId: "1671",
			Name:   "Ilir Timur Tiga",
		},
		{
			Id:     "167201",
			CityId: "1672",
			Name:   "Pagar Alam Utara",
		},
		{
			Id:     "167202",
			CityId: "1672",
			Name:   "Pagar Alam Selatan",
		},
		{
			Id:     "167203",
			CityId: "1672",
			Name:   "Dempo Utara",
		},
		{
			Id:     "167204",
			CityId: "1672",
			Name:   "Dempo Selatan",
		},
		{
			Id:     "167205",
			CityId: "1672",
			Name:   "Dempo Tengah",
		},
		{
			Id:     "167301",
			CityId: "1673",
			Name:   "Lubuk Linggau Timur I",
		},
		{
			Id:     "167302",
			CityId: "1673",
			Name:   "Lubuk Linggau Barat I",
		},
		{
			Id:     "167303",
			CityId: "1673",
			Name:   "Lubuk Linggau Selatan I",
		},
		{
			Id:     "167304",
			CityId: "1673",
			Name:   "Lubuk Linggau Utara I",
		},
		{
			Id:     "167305",
			CityId: "1673",
			Name:   "Lubuk Linggau Timur II",
		},
		{
			Id:     "167306",
			CityId: "1673",
			Name:   "Lubuk Linggau Barat II",
		},
		{
			Id:     "167307",
			CityId: "1673",
			Name:   "Lubuk Linggau Selatan II",
		},
		{
			Id:     "167308",
			CityId: "1673",
			Name:   "Lubuk Linggau Utara II",
		},
		{
			Id:     "167401",
			CityId: "1674",
			Name:   "Prabumulih Barat",
		},
		{
			Id:     "167402",
			CityId: "1674",
			Name:   "Prabumulih Timur",
		},
		{
			Id:     "167403",
			CityId: "1674",
			Name:   "Cambai",
		},
		{
			Id:     "167404",
			CityId: "1674",
			Name:   "Rambang Kapak Tengah",
		},
		{
			Id:     "167405",
			CityId: "1674",
			Name:   "Prabumulih Utara",
		},
		{
			Id:     "167406",
			CityId: "1674",
			Name:   "Prabumulih Selatan",
		},
		{
			Id:     "170101",
			CityId: "1701",
			Name:   "Kedurang",
		},
		{
			Id:     "170102",
			CityId: "1701",
			Name:   "Seginim",
		},
		{
			Id:     "170103",
			CityId: "1701",
			Name:   "Pino",
		},
		{
			Id:     "170104",
			CityId: "1701",
			Name:   "Manna",
		},
		{
			Id:     "170105",
			CityId: "1701",
			Name:   "Kota Manna",
		},
		{
			Id:     "170106",
			CityId: "1701",
			Name:   "Pino Raya",
		},
		{
			Id:     "170107",
			CityId: "1701",
			Name:   "Kedurang Ilir",
		},
		{
			Id:     "170108",
			CityId: "1701",
			Name:   "Air Nipis",
		},
		{
			Id:     "170109",
			CityId: "1701",
			Name:   "Ulu Manna",
		},
		{
			Id:     "170110",
			CityId: "1701",
			Name:   "Bunga Mas",
		},
		{
			Id:     "170111",
			CityId: "1701",
			Name:   "Pasar Manna",
		},
		{
			Id:     "170206",
			CityId: "1702",
			Name:   "Kota Padang",
		},
		{
			Id:     "170207",
			CityId: "1702",
			Name:   "Padang Ulak Tanding",
		},
		{
			Id:     "170208",
			CityId: "1702",
			Name:   "Sindang Kelingi",
		},
		{
			Id:     "170209",
			CityId: "1702",
			Name:   "Curup",
		},
		{
			Id:     "170210",
			CityId: "1702",
			Name:   "Bermani Ulu",
		},
		{
			Id:     "170211",
			CityId: "1702",
			Name:   "Selupu Rejang",
		},
		{
			Id:     "170216",
			CityId: "1702",
			Name:   "Curup Utara",
		},
		{
			Id:     "170217",
			CityId: "1702",
			Name:   "Curup Timur",
		},
		{
			Id:     "170218",
			CityId: "1702",
			Name:   "Curup Selatan",
		},
		{
			Id:     "170219",
			CityId: "1702",
			Name:   "Curup Tengah",
		},
		{
			Id:     "170220",
			CityId: "1702",
			Name:   "Binduriang",
		},
		{
			Id:     "170221",
			CityId: "1702",
			Name:   "Sindang Beliti Ulu",
		},
		{
			Id:     "170222",
			CityId: "1702",
			Name:   "Sindang Dataran",
		},
		{
			Id:     "170223",
			CityId: "1702",
			Name:   "Sindang Beliti Ilir",
		},
		{
			Id:     "170224",
			CityId: "1702",
			Name:   "Bermani Ulu Raya",
		},
		{
			Id:     "170301",
			CityId: "1703",
			Name:   "Enggano",
		},
		{
			Id:     "170306",
			CityId: "1703",
			Name:   "Kerkap",
		},
		{
			Id:     "170307",
			CityId: "1703",
			Name:   "Kota Arga Makmur",
		},
		{
			Id:     "170308",
			CityId: "1703",
			Name:   "Giri Mulya",
		},
		{
			Id:     "170309",
			CityId: "1703",
			Name:   "Padang Jaya",
		},
		{
			Id:     "170310",
			CityId: "1703",
			Name:   "Lais",
		},
		{
			Id:     "170311",
			CityId: "1703",
			Name:   "Batik Nau",
		},
		{
			Id:     "170312",
			CityId: "1703",
			Name:   "Ketahun",
		},
		{
			Id:     "170313",
			CityId: "1703",
			Name:   "Napal Putih",
		},
		{
			Id:     "170314",
			CityId: "1703",
			Name:   "Putri Hijau",
		},
		{
			Id:     "170315",
			CityId: "1703",
			Name:   "Air Besi",
		},
		{
			Id:     "170316",
			CityId: "1703",
			Name:   "Air Napal",
		},
		{
			Id:     "170319",
			CityId: "1703",
			Name:   "Hulu Palik",
		},
		{
			Id:     "170320",
			CityId: "1703",
			Name:   "Air Padang",
		},
		{
			Id:     "170321",
			CityId: "1703",
			Name:   "Arma Jaya",
		},
		{
			Id:     "170322",
			CityId: "1703",
			Name:   "Tanjung Agung Palik",
		},
		{
			Id:     "170323",
			CityId: "1703",
			Name:   "Ulok Kupai",
		},
		{
			Id:     "170324",
			CityId: "1703",
			Name:   "Pinang Raya",
		},
		{
			Id:     "170325",
			CityId: "1703",
			Name:   "Marga Sakti Sebelat",
		},
		{
			Id:     "170401",
			CityId: "1704",
			Name:   "Kinal",
		},
		{
			Id:     "170402",
			CityId: "1704",
			Name:   "Tanjung Kemuning",
		},
		{
			Id:     "170403",
			CityId: "1704",
			Name:   "Kaur Utara",
		},
		{
			Id:     "170404",
			CityId: "1704",
			Name:   "Kaur Tengah",
		},
		{
			Id:     "170405",
			CityId: "1704",
			Name:   "Kaur Selatan",
		},
		{
			Id:     "170406",
			CityId: "1704",
			Name:   "Maje",
		},
		{
			Id:     "170407",
			CityId: "1704",
			Name:   "Nasal",
		},
		{
			Id:     "170408",
			CityId: "1704",
			Name:   "Semidang Gumay",
		},
		{
			Id:     "170409",
			CityId: "1704",
			Name:   "Kelam Tengah",
		},
		{
			Id:     "170410",
			CityId: "1704",
			Name:   "Luas",
		},
		{
			Id:     "170411",
			CityId: "1704",
			Name:   "Muara Sahung",
		},
		{
			Id:     "170412",
			CityId: "1704",
			Name:   "Tetap",
		},
		{
			Id:     "170413",
			CityId: "1704",
			Name:   "Lungkang Kule",
		},
		{
			Id:     "170414",
			CityId: "1704",
			Name:   "Padang Guci Hilir",
		},
		{
			Id:     "170415",
			CityId: "1704",
			Name:   "Padang Guci Hulu",
		},
		{
			Id:     "170501",
			CityId: "1705",
			Name:   "Sukaraja",
		},
		{
			Id:     "170502",
			CityId: "1705",
			Name:   "Seluma",
		},
		{
			Id:     "170503",
			CityId: "1705",
			Name:   "Talo",
		},
		{
			Id:     "170504",
			CityId: "1705",
			Name:   "Semidang Alas",
		},
		{
			Id:     "170505",
			CityId: "1705",
			Name:   "Semidang Alas Maras",
		},
		{
			Id:     "170506",
			CityId: "1705",
			Name:   "Air Periukan",
		},
		{
			Id:     "170507",
			CityId: "1705",
			Name:   "Lubuk Sandi",
		},
		{
			Id:     "170508",
			CityId: "1705",
			Name:   "Seluma Barat",
		},
		{
			Id:     "170509",
			CityId: "1705",
			Name:   "Seluma Timur",
		},
		{
			Id:     "170510",
			CityId: "1705",
			Name:   "Seluma Utara",
		},
		{
			Id:     "170511",
			CityId: "1705",
			Name:   "Seluma Selatan",
		},
		{
			Id:     "170512",
			CityId: "1705",
			Name:   "Talo Kecil",
		},
		{
			Id:     "170513",
			CityId: "1705",
			Name:   "Ulu Talo",
		},
		{
			Id:     "170514",
			CityId: "1705",
			Name:   "Ilir Talo",
		},
		{
			Id:     "170601",
			CityId: "1706",
			Name:   "Lubuk Pinang",
		},
		{
			Id:     "170602",
			CityId: "1706",
			Name:   "Kota Mukomuko",
		},
		{
			Id:     "170603",
			CityId: "1706",
			Name:   "Teras Terunjam",
		},
		{
			Id:     "170604",
			CityId: "1706",
			Name:   "Pondok Suguh",
		},
		{
			Id:     "170605",
			CityId: "1706",
			Name:   "Ipuh",
		},
		{
			Id:     "170606",
			CityId: "1706",
			Name:   "Malin Deman",
		},
		{
			Id:     "170607",
			CityId: "1706",
			Name:   "Air Rami",
		},
		{
			Id:     "170608",
			CityId: "1706",
			Name:   "Teramang Jaya",
		},
		{
			Id:     "170609",
			CityId: "1706",
			Name:   "Selagan Raya",
		},
		{
			Id:     "170610",
			CityId: "1706",
			Name:   "Penarik",
		},
		{
			Id:     "170611",
			CityId: "1706",
			Name:   "XIV Koto",
		},
		{
			Id:     "170612",
			CityId: "1706",
			Name:   "V Koto",
		},
		{
			Id:     "170613",
			CityId: "1706",
			Name:   "Air Majunto",
		},
		{
			Id:     "170614",
			CityId: "1706",
			Name:   "Air Dikit",
		},
		{
			Id:     "170615",
			CityId: "1706",
			Name:   "Sungai Rumbai",
		},
		{
			Id:     "170701",
			CityId: "1707",
			Name:   "Lebong Utara",
		},
		{
			Id:     "170702",
			CityId: "1707",
			Name:   "Lebong Atas",
		},
		{
			Id:     "170703",
			CityId: "1707",
			Name:   "Lebong Tengah",
		},
		{
			Id:     "170704",
			CityId: "1707",
			Name:   "Lebong Selatan",
		},
		{
			Id:     "170705",
			CityId: "1707",
			Name:   "Rimbo Pengadang",
		},
		{
			Id:     "170706",
			CityId: "1707",
			Name:   "Topos",
		},
		{
			Id:     "170707",
			CityId: "1707",
			Name:   "Bingin Kuning",
		},
		{
			Id:     "170708",
			CityId: "1707",
			Name:   "Lebong Sakti",
		},
		{
			Id:     "170709",
			CityId: "1707",
			Name:   "Pelabai",
		},
		{
			Id:     "170710",
			CityId: "1707",
			Name:   "Amen",
		},
		{
			Id:     "170711",
			CityId: "1707",
			Name:   "Uram Jaya",
		},
		{
			Id:     "170712",
			CityId: "1707",
			Name:   "Pinang Belapis",
		},
		{
			Id:     "170801",
			CityId: "1708",
			Name:   "Bermani Ilir",
		},
		{
			Id:     "170802",
			CityId: "1708",
			Name:   "Ujan Mas",
		},
		{
			Id:     "170803",
			CityId: "1708",
			Name:   "Tebat Karai",
		},
		{
			Id:     "170804",
			CityId: "1708",
			Name:   "Kepahiang",
		},
		{
			Id:     "170805",
			CityId: "1708",
			Name:   "Merigi",
		},
		{
			Id:     "170806",
			CityId: "1708",
			Name:   "Kebawetan",
		},
		{
			Id:     "170807",
			CityId: "1708",
			Name:   "Seberang Musi",
		},
		{
			Id:     "170808",
			CityId: "1708",
			Name:   "Muara Kemumu",
		},
		{
			Id:     "170901",
			CityId: "1709",
			Name:   "Karang Tinggi",
		},
		{
			Id:     "170902",
			CityId: "1709",
			Name:   "Talang Empat",
		},
		{
			Id:     "170903",
			CityId: "1709",
			Name:   "Pondok Kelapa",
		},
		{
			Id:     "170904",
			CityId: "1709",
			Name:   "Pematang Tiga",
		},
		{
			Id:     "170905",
			CityId: "1709",
			Name:   "Pagar Jati",
		},
		{
			Id:     "170906",
			CityId: "1709",
			Name:   "Taba Penanjung",
		},
		{
			Id:     "170907",
			CityId: "1709",
			Name:   "Merigi Kelindang",
		},
		{
			Id:     "170908",
			CityId: "1709",
			Name:   "Merigi Sakti",
		},
		{
			Id:     "170909",
			CityId: "1709",
			Name:   "Pondok Kubang",
		},
		{
			Id:     "170910",
			CityId: "1709",
			Name:   "Bang Haji",
		},
		{
			Id:     "170911",
			CityId: "1709",
			Name:   "Semidang Lagan",
		},
		{
			Id:     "177101",
			CityId: "1771",
			Name:   "Selebar",
		},
		{
			Id:     "177102",
			CityId: "1771",
			Name:   "Gading Cempaka",
		},
		{
			Id:     "177103",
			CityId: "1771",
			Name:   "Teluk Segara",
		},
		{
			Id:     "177104",
			CityId: "1771",
			Name:   "Muara Bangka Hulu",
		},
		{
			Id:     "177105",
			CityId: "1771",
			Name:   "Kampung Melayu",
		},
		{
			Id:     "177106",
			CityId: "1771",
			Name:   "Ratu Agung",
		},
		{
			Id:     "177107",
			CityId: "1771",
			Name:   "Ratu Samban",
		},
		{
			Id:     "177108",
			CityId: "1771",
			Name:   "Sungai Serut",
		},
		{
			Id:     "177109",
			CityId: "1771",
			Name:   "Singaran Pati",
		},
		{
			Id:     "180104",
			CityId: "1801",
			Name:   "Natar",
		},
		{
			Id:     "180105",
			CityId: "1801",
			Name:   "Tanjung Bintang",
		},
		{
			Id:     "180106",
			CityId: "1801",
			Name:   "Kalianda",
		},
		{
			Id:     "180107",
			CityId: "1801",
			Name:   "Sidomulyo",
		},
		{
			Id:     "180108",
			CityId: "1801",
			Name:   "Katibung",
		},
		{
			Id:     "180109",
			CityId: "1801",
			Name:   "Penengahan",
		},
		{
			Id:     "180110",
			CityId: "1801",
			Name:   "Palas",
		},
		{
			Id:     "180113",
			CityId: "1801",
			Name:   "Jati Agung",
		},
		{
			Id:     "180114",
			CityId: "1801",
			Name:   "Ketapang",
		},
		{
			Id:     "180115",
			CityId: "1801",
			Name:   "Sragi",
		},
		{
			Id:     "180116",
			CityId: "1801",
			Name:   "Raja Basa",
		},
		{
			Id:     "180117",
			CityId: "1801",
			Name:   "Candipuro",
		},
		{
			Id:     "180118",
			CityId: "1801",
			Name:   "Merbau Mataram",
		},
		{
			Id:     "180121",
			CityId: "1801",
			Name:   "Bakauheni",
		},
		{
			Id:     "180122",
			CityId: "1801",
			Name:   "Tanjung Sari",
		},
		{
			Id:     "180123",
			CityId: "1801",
			Name:   "Way Sulan",
		},
		{
			Id:     "180124",
			CityId: "1801",
			Name:   "Way Panji",
		},
		{
			Id:     "180201",
			CityId: "1802",
			Name:   "Kalirejo",
		},
		{
			Id:     "180202",
			CityId: "1802",
			Name:   "Bangun Rejo",
		},
		{
			Id:     "180203",
			CityId: "1802",
			Name:   "Padang Ratu",
		},
		{
			Id:     "180204",
			CityId: "1802",
			Name:   "Gunung Sugih",
		},
		{
			Id:     "180205",
			CityId: "1802",
			Name:   "Trimurjo",
		},
		{
			Id:     "180206",
			CityId: "1802",
			Name:   "Punggur",
		},
		{
			Id:     "180207",
			CityId: "1802",
			Name:   "Terbanggi Besar",
		},
		{
			Id:     "180208",
			CityId: "1802",
			Name:   "Seputih Raman",
		},
		{
			Id:     "180209",
			CityId: "1802",
			Name:   "Rumbia",
		},
		{
			Id:     "180210",
			CityId: "1802",
			Name:   "Seputih Banyak",
		},
		{
			Id:     "180211",
			CityId: "1802",
			Name:   "Seputih Mataram",
		},
		{
			Id:     "180212",
			CityId: "1802",
			Name:   "Seputih Surabaya",
		},
		{
			Id:     "180213",
			CityId: "1802",
			Name:   "Terusan Nunyai",
		},
		{
			Id:     "180214",
			CityId: "1802",
			Name:   "Bumi Ratu Nuban",
		},
		{
			Id:     "180215",
			CityId: "1802",
			Name:   "Bekri",
		},
		{
			Id:     "180216",
			CityId: "1802",
			Name:   "Seputih Agung",
		},
		{
			Id:     "180217",
			CityId: "1802",
			Name:   "Way Pangubuan",
		},
		{
			Id:     "180218",
			CityId: "1802",
			Name:   "Bandar Mataram",
		},
		{
			Id:     "180219",
			CityId: "1802",
			Name:   "Pubian",
		},
		{
			Id:     "180220",
			CityId: "1802",
			Name:   "Selagai Lingga",
		},
		{
			Id:     "180221",
			CityId: "1802",
			Name:   "Anak Tuha",
		},
		{
			Id:     "180222",
			CityId: "1802",
			Name:   "Sendang Agung",
		},
		{
			Id:     "180223",
			CityId: "1802",
			Name:   "Kota Gajah",
		},
		{
			Id:     "180224",
			CityId: "1802",
			Name:   "Bumi Nabung",
		},
		{
			Id:     "180225",
			CityId: "1802",
			Name:   "Way Seputih",
		},
		{
			Id:     "180226",
			CityId: "1802",
			Name:   "Bandar Surabaya",
		},
		{
			Id:     "180227",
			CityId: "1802",
			Name:   "Anak Ratu Aji",
		},
		{
			Id:     "180228",
			CityId: "1802",
			Name:   "Putra Rumbia",
		},
		{
			Id:     "180301",
			CityId: "1803",
			Name:   "Bukit Kemuning",
		},
		{
			Id:     "180302",
			CityId: "1803",
			Name:   "Kotabumi",
		},
		{
			Id:     "180303",
			CityId: "1803",
			Name:   "Sungkai Selatan",
		},
		{
			Id:     "180304",
			CityId: "1803",
			Name:   "Tanjung Raja",
		},
		{
			Id:     "180305",
			CityId: "1803",
			Name:   "Abung Timur",
		},
		{
			Id:     "180306",
			CityId: "1803",
			Name:   "Abung Barat",
		},
		{
			Id:     "180307",
			CityId: "1803",
			Name:   "Abung Selatan",
		},
		{
			Id:     "180308",
			CityId: "1803",
			Name:   "Sungkai Utara",
		},
		{
			Id:     "180309",
			CityId: "1803",
			Name:   "Kotabumi Utara",
		},
		{
			Id:     "180310",
			CityId: "1803",
			Name:   "Kotabumi Selatan",
		},
		{
			Id:     "180311",
			CityId: "1803",
			Name:   "Abung Tengah",
		},
		{
			Id:     "180312",
			CityId: "1803",
			Name:   "Abung Tinggi",
		},
		{
			Id:     "180313",
			CityId: "1803",
			Name:   "Abung Semuli",
		},
		{
			Id:     "180314",
			CityId: "1803",
			Name:   "Abung Surakarta",
		},
		{
			Id:     "180315",
			CityId: "1803",
			Name:   "Muara Sungkai",
		},
		{
			Id:     "180316",
			CityId: "1803",
			Name:   "Bunga Mayang",
		},
		{
			Id:     "180317",
			CityId: "1803",
			Name:   "Hulu Sungkai",
		},
		{
			Id:     "180318",
			CityId: "1803",
			Name:   "Sungkai Tengah",
		},
		{
			Id:     "180319",
			CityId: "1803",
			Name:   "Abung Pekurun",
		},
		{
			Id:     "180320",
			CityId: "1803",
			Name:   "Sungkai Jaya",
		},
		{
			Id:     "180321",
			CityId: "1803",
			Name:   "Sungkai Barat",
		},
		{
			Id:     "180322",
			CityId: "1803",
			Name:   "Abung Kunang",
		},
		{
			Id:     "180323",
			CityId: "1803",
			Name:   "Blambangan Pagar",
		},
		{
			Id:     "180404",
			CityId: "1804",
			Name:   "Balik Bukit",
		},
		{
			Id:     "180405",
			CityId: "1804",
			Name:   "Sumber Jaya",
		},
		{
			Id:     "180406",
			CityId: "1804",
			Name:   "Belalau",
		},
		{
			Id:     "180407",
			CityId: "1804",
			Name:   "Way Tenong",
		},
		{
			Id:     "180408",
			CityId: "1804",
			Name:   "Sekincau",
		},
		{
			Id:     "180409",
			CityId: "1804",
			Name:   "Suoh",
		},
		{
			Id:     "180410",
			CityId: "1804",
			Name:   "Batu Brak",
		},
		{
			Id:     "180411",
			CityId: "1804",
			Name:   "Sukau",
		},
		{
			Id:     "180415",
			CityId: "1804",
			Name:   "Gedung Surian",
		},
		{
			Id:     "180418",
			CityId: "1804",
			Name:   "Kebun Tebu",
		},
		{
			Id:     "180419",
			CityId: "1804",
			Name:   "Air Hitam",
		},
		{
			Id:     "180420",
			CityId: "1804",
			Name:   "Pagar Dewa",
		},
		{
			Id:     "180421",
			CityId: "1804",
			Name:   "Batu Ketulis",
		},
		{
			Id:     "180422",
			CityId: "1804",
			Name:   "Lumbok Seminung",
		},
		{
			Id:     "180423",
			CityId: "1804",
			Name:   "Bandar Negeri Suoh",
		},
		{
			Id:     "180502",
			CityId: "1805",
			Name:   "Menggala",
		},
		{
			Id:     "180506",
			CityId: "1805",
			Name:   "Gedung Aji",
		},
		{
			Id:     "180508",
			CityId: "1805",
			Name:   "Banjar Agung",
		},
		{
			Id:     "180511",
			CityId: "1805",
			Name:   "Gedung Meneng",
		},
		{
			Id:     "180512",
			CityId: "1805",
			Name:   "Rawa Jitu Selatan",
		},
		{
			Id:     "180513",
			CityId: "1805",
			Name:   "Penawar Tama",
		},
		{
			Id:     "180518",
			CityId: "1805",
			Name:   "Rawa Jitu Timur",
		},
		{
			Id:     "180520",
			CityId: "1805",
			Name:   "Banjar Margo",
		},
		{
			Id:     "180522",
			CityId: "1805",
			Name:   "Rawa Pitu",
		},
		{
			Id:     "180523",
			CityId: "1805",
			Name:   "Penawar Aji",
		},
		{
			Id:     "180525",
			CityId: "1805",
			Name:   "Dente Teladas",
		},
		{
			Id:     "180526",
			CityId: "1805",
			Name:   "Meraksa Aji",
		},
		{
			Id:     "180527",
			CityId: "1805",
			Name:   "Gedung Aji Baru",
		},
		{
			Id:     "180529",
			CityId: "1805",
			Name:   "Banjar Baru",
		},
		{
			Id:     "180530",
			CityId: "1805",
			Name:   "Menggala Timur",
		},
		{
			Id:     "180601",
			CityId: "1806",
			Name:   "Kota Agung",
		},
		{
			Id:     "180602",
			CityId: "1806",
			Name:   "Talang Padang",
		},
		{
			Id:     "180603",
			CityId: "1806",
			Name:   "Wonosobo",
		},
		{
			Id:     "180604",
			CityId: "1806",
			Name:   "Pulau Panggung",
		},
		{
			Id:     "180609",
			CityId: "1806",
			Name:   "Cukuh Balak",
		},
		{
			Id:     "180611",
			CityId: "1806",
			Name:   "Pugung",
		},
		{
			Id:     "180612",
			CityId: "1806",
			Name:   "Semaka",
		},
		{
			Id:     "180613",
			CityId: "1806",
			Name:   "Sumberejo",
		},
		{
			Id:     "180615",
			CityId: "1806",
			Name:   "Ulu Belu",
		},
		{
			Id:     "180616",
			CityId: "1806",
			Name:   "Pematang Sawa",
		},
		{
			Id:     "180617",
			CityId: "1806",
			Name:   "Kelumbayan",
		},
		{
			Id:     "180618",
			CityId: "1806",
			Name:   "Kota Agung Barat",
		},
		{
			Id:     "180619",
			CityId: "1806",
			Name:   "Kota Agung Timur",
		},
		{
			Id:     "180620",
			CityId: "1806",
			Name:   "Gisting",
		},
		{
			Id:     "180621",
			CityId: "1806",
			Name:   "Gunung Alip",
		},
		{
			Id:     "180624",
			CityId: "1806",
			Name:   "Limau",
		},
		{
			Id:     "180625",
			CityId: "1806",
			Name:   "Bandar Negeri Semuong",
		},
		{
			Id:     "180626",
			CityId: "1806",
			Name:   "Air Naningan",
		},
		{
			Id:     "180627",
			CityId: "1806",
			Name:   "Bulok",
		},
		{
			Id:     "180628",
			CityId: "1806",
			Name:   "Klumbayan Barat",
		},
		{
			Id:     "180701",
			CityId: "1807",
			Name:   "Sukadana",
		},
		{
			Id:     "180702",
			CityId: "1807",
			Name:   "Labuhan Maringgai",
		},
		{
			Id:     "180703",
			CityId: "1807",
			Name:   "Jabung",
		},
		{
			Id:     "180704",
			CityId: "1807",
			Name:   "Pekalongan",
		},
		{
			Id:     "180705",
			CityId: "1807",
			Name:   "Sekampung",
		},
		{
			Id:     "180706",
			CityId: "1807",
			Name:   "Batanghari",
		},
		{
			Id:     "180707",
			CityId: "1807",
			Name:   "Way Jepara",
		},
		{
			Id:     "180708",
			CityId: "1807",
			Name:   "Purbolinggo",
		},
		{
			Id:     "180709",
			CityId: "1807",
			Name:   "Raman Utara",
		},
		{
			Id:     "180710",
			CityId: "1807",
			Name:   "Metro Kibang",
		},
		{
			Id:     "180711",
			CityId: "1807",
			Name:   "Marga Tiga",
		},
		{
			Id:     "180712",
			CityId: "1807",
			Name:   "Sekampung Udik",
		},
		{
			Id:     "180713",
			CityId: "1807",
			Name:   "Batanghari Nuban",
		},
		{
			Id:     "180714",
			CityId: "1807",
			Name:   "Bumi Agung",
		},
		{
			Id:     "180715",
			CityId: "1807",
			Name:   "Bandar Sribhawono",
		},
		{
			Id:     "180716",
			CityId: "1807",
			Name:   "Mataram Baru",
		},
		{
			Id:     "180717",
			CityId: "1807",
			Name:   "Melinting",
		},
		{
			Id:     "180718",
			CityId: "1807",
			Name:   "Gunung Pelindung",
		},
		{
			Id:     "180719",
			CityId: "1807",
			Name:   "Pasir Sakti",
		},
		{
			Id:     "180720",
			CityId: "1807",
			Name:   "Waway Karya",
		},
		{
			Id:     "180721",
			CityId: "1807",
			Name:   "Labuhan Ratu",
		},
		{
			Id:     "180722",
			CityId: "1807",
			Name:   "Braja Selebah",
		},
		{
			Id:     "180723",
			CityId: "1807",
			Name:   "Way Bungur",
		},
		{
			Id:     "180724",
			CityId: "1807",
			Name:   "Marga Sekampung",
		},
		{
			Id:     "180801",
			CityId: "1808",
			Name:   "Blambangan Umpu",
		},
		{
			Id:     "180802",
			CityId: "1808",
			Name:   "Kasui",
		},
		{
			Id:     "180803",
			CityId: "1808",
			Name:   "Banjit",
		},
		{
			Id:     "180804",
			CityId: "1808",
			Name:   "Baradatu",
		},
		{
			Id:     "180805",
			CityId: "1808",
			Name:   "Bahuga",
		},
		{
			Id:     "180806",
			CityId: "1808",
			Name:   "Pakuan Ratu",
		},
		{
			Id:     "180807",
			CityId: "1808",
			Name:   "Negeri Agung",
		},
		{
			Id:     "180808",
			CityId: "1808",
			Name:   "Way Tuba",
		},
		{
			Id:     "180809",
			CityId: "1808",
			Name:   "Rebang Tangkas",
		},
		{
			Id:     "180810",
			CityId: "1808",
			Name:   "Gunung Labuhan",
		},
		{
			Id:     "180811",
			CityId: "1808",
			Name:   "Negara Batin",
		},
		{
			Id:     "180812",
			CityId: "1808",
			Name:   "Negeri Besar",
		},
		{
			Id:     "180813",
			CityId: "1808",
			Name:   "Buay Bahuga",
		},
		{
			Id:     "180814",
			CityId: "1808",
			Name:   "Bumi Agung",
		},
		{
			Id:     "180815",
			CityId: "1808",
			Name:   "Umpu Semenguk",
		},
		{
			Id:     "180901",
			CityId: "1809",
			Name:   "Gedong Tataan",
		},
		{
			Id:     "180902",
			CityId: "1809",
			Name:   "Negeri Katon",
		},
		{
			Id:     "180903",
			CityId: "1809",
			Name:   "Tegineneng",
		},
		{
			Id:     "180904",
			CityId: "1809",
			Name:   "Way Lima",
		},
		{
			Id:     "180905",
			CityId: "1809",
			Name:   "Padang Cermin",
		},
		{
			Id:     "180906",
			CityId: "1809",
			Name:   "Punduh Pidada",
		},
		{
			Id:     "180907",
			CityId: "1809",
			Name:   "Kedondong",
		},
		{
			Id:     "180908",
			CityId: "1809",
			Name:   "Marga Punduh",
		},
		{
			Id:     "180909",
			CityId: "1809",
			Name:   "Way Khilau",
		},
		{
			Id:     "180910",
			CityId: "1809",
			Name:   "Teluk Pandan",
		},
		{
			Id:     "180911",
			CityId: "1809",
			Name:   "Way Ratai",
		},
		{
			Id:     "181001",
			CityId: "1810",
			Name:   "Pringsewu",
		},
		{
			Id:     "181002",
			CityId: "1810",
			Name:   "Gading Rejo",
		},
		{
			Id:     "181003",
			CityId: "1810",
			Name:   "Ambarawa",
		},
		{
			Id:     "181004",
			CityId: "1810",
			Name:   "Pardasuka",
		},
		{
			Id:     "181005",
			CityId: "1810",
			Name:   "Pagelaran",
		},
		{
			Id:     "181006",
			CityId: "1810",
			Name:   "Banyumas",
		},
		{
			Id:     "181007",
			CityId: "1810",
			Name:   "Adiluwih",
		},
		{
			Id:     "181008",
			CityId: "1810",
			Name:   "Sukoharjo",
		},
		{
			Id:     "181009",
			CityId: "1810",
			Name:   "Pagelaran Utara",
		},
		{
			Id:     "181101",
			CityId: "1811",
			Name:   "Mesuji",
		},
		{
			Id:     "181102",
			CityId: "1811",
			Name:   "Mesuji Timur",
		},
		{
			Id:     "181103",
			CityId: "1811",
			Name:   "Rawa Jitu Utara",
		},
		{
			Id:     "181104",
			CityId: "1811",
			Name:   "Way Serdang",
		},
		{
			Id:     "181105",
			CityId: "1811",
			Name:   "Simpang Pematang",
		},
		{
			Id:     "181106",
			CityId: "1811",
			Name:   "Panca Jaya",
		},
		{
			Id:     "181107",
			CityId: "1811",
			Name:   "Tanjung Raya",
		},
		{
			Id:     "181201",
			CityId: "1812",
			Name:   "Tulang Bawang Tengah",
		},
		{
			Id:     "181202",
			CityId: "1812",
			Name:   "Tumijajar",
		},
		{
			Id:     "181203",
			CityId: "1812",
			Name:   "Tulang Bawang Udik",
		},
		{
			Id:     "181204",
			CityId: "1812",
			Name:   "Gunung Terang",
		},
		{
			Id:     "181205",
			CityId: "1812",
			Name:   "Gunung Agung",
		},
		{
			Id:     "181206",
			CityId: "1812",
			Name:   "Way Kenanga",
		},
		{
			Id:     "181207",
			CityId: "1812",
			Name:   "Lambu Kibang",
		},
		{
			Id:     "181208",
			CityId: "1812",
			Name:   "Pagar Dewa",
		},
		{
			Id:     "181209",
			CityId: "1812",
			Name:   "Batu Putih",
		},
		{
			Id:     "181301",
			CityId: "1813",
			Name:   "Pesisir Tengah",
		},
		{
			Id:     "181302",
			CityId: "1813",
			Name:   "Pesisir Selatan",
		},
		{
			Id:     "181303",
			CityId: "1813",
			Name:   "Lemong",
		},
		{
			Id:     "181304",
			CityId: "1813",
			Name:   "Pesisir Utara",
		},
		{
			Id:     "181305",
			CityId: "1813",
			Name:   "Karya Penggawa",
		},
		{
			Id:     "181306",
			CityId: "1813",
			Name:   "Pulaupisang",
		},
		{
			Id:     "181307",
			CityId: "1813",
			Name:   "Way Krui",
		},
		{
			Id:     "181308",
			CityId: "1813",
			Name:   "Krui Selatan",
		},
		{
			Id:     "181309",
			CityId: "1813",
			Name:   "Ngambur",
		},
		{
			Id:     "181310",
			CityId: "1813",
			Name:   "Ngaras",
		},
		{
			Id:     "181311",
			CityId: "1813",
			Name:   "Bangkunat",
		},
		{
			Id:     "187101",
			CityId: "1871",
			Name:   "Kedaton",
		},
		{
			Id:     "187102",
			CityId: "1871",
			Name:   "Sukarame",
		},
		{
			Id:     "187103",
			CityId: "1871",
			Name:   "Tanjungkarang Barat",
		},
		{
			Id:     "187104",
			CityId: "1871",
			Name:   "Panjang",
		},
		{
			Id:     "187105",
			CityId: "1871",
			Name:   "Tanjungkarang Timur",
		},
		{
			Id:     "187106",
			CityId: "1871",
			Name:   "Tanjungkarang Pusat",
		},
		{
			Id:     "187107",
			CityId: "1871",
			Name:   "Telukbetung Selatan",
		},
		{
			Id:     "187108",
			CityId: "1871",
			Name:   "Telukbetung Barat",
		},
		{
			Id:     "187109",
			CityId: "1871",
			Name:   "Telukbetung Utara",
		},
		{
			Id:     "187110",
			CityId: "1871",
			Name:   "Rajabasa",
		},
		{
			Id:     "187111",
			CityId: "1871",
			Name:   "Tanjung Senang",
		},
		{
			Id:     "187112",
			CityId: "1871",
			Name:   "Sukabumi",
		},
		{
			Id:     "187113",
			CityId: "1871",
			Name:   "Kemiling",
		},
		{
			Id:     "187114",
			CityId: "1871",
			Name:   "Labuhan Ratu",
		},
		{
			Id:     "187115",
			CityId: "1871",
			Name:   "Way Halim",
		},
		{
			Id:     "187116",
			CityId: "1871",
			Name:   "Langkapura",
		},
		{
			Id:     "187117",
			CityId: "1871",
			Name:   "Enggal",
		},
		{
			Id:     "187118",
			CityId: "1871",
			Name:   "Kedamaian",
		},
		{
			Id:     "187119",
			CityId: "1871",
			Name:   "Telukbetung Timur",
		},
		{
			Id:     "187120",
			CityId: "1871",
			Name:   "Bumi Waras",
		},
		{
			Id:     "187201",
			CityId: "1872",
			Name:   "Metro Pusat",
		},
		{
			Id:     "187202",
			CityId: "1872",
			Name:   "Metro Utara",
		},
		{
			Id:     "187203",
			CityId: "1872",
			Name:   "Metro Barat",
		},
		{
			Id:     "187204",
			CityId: "1872",
			Name:   "Metro Timur",
		},
		{
			Id:     "187205",
			CityId: "1872",
			Name:   "Metro Selatan",
		},
		{
			Id:     "190101",
			CityId: "1901",
			Name:   "Sungailiat",
		},
		{
			Id:     "190102",
			CityId: "1901",
			Name:   "Belinyu",
		},
		{
			Id:     "190103",
			CityId: "1901",
			Name:   "Merawang",
		},
		{
			Id:     "190104",
			CityId: "1901",
			Name:   "Mendo Barat",
		},
		{
			Id:     "190105",
			CityId: "1901",
			Name:   "Pemali",
		},
		{
			Id:     "190106",
			CityId: "1901",
			Name:   "Bakam",
		},
		{
			Id:     "190107",
			CityId: "1901",
			Name:   "Riau Silip",
		},
		{
			Id:     "190108",
			CityId: "1901",
			Name:   "Puding Besar",
		},
		{
			Id:     "190201",
			CityId: "1902",
			Name:   "Tanjung Pandan",
		},
		{
			Id:     "190202",
			CityId: "1902",
			Name:   "Membalong",
		},
		{
			Id:     "190203",
			CityId: "1902",
			Name:   "Selat Nasik",
		},
		{
			Id:     "190204",
			CityId: "1902",
			Name:   "Sijuk",
		},
		{
			Id:     "190205",
			CityId: "1902",
			Name:   "Badau",
		},
		{
			Id:     "190301",
			CityId: "1903",
			Name:   "Toboali",
		},
		{
			Id:     "190302",
			CityId: "1903",
			Name:   "Lepar",
		},
		{
			Id:     "190303",
			CityId: "1903",
			Name:   "Airgegas",
		},
		{
			Id:     "190304",
			CityId: "1903",
			Name:   "Simpang Rimba",
		},
		{
			Id:     "190305",
			CityId: "1903",
			Name:   "Payung",
		},
		{
			Id:     "190306",
			CityId: "1903",
			Name:   "Tukak Sadai",
		},
		{
			Id:     "190307",
			CityId: "1903",
			Name:   "Pulaubesar",
		},
		{
			Id:     "190308",
			CityId: "1903",
			Name:   "Kepulauan Pongok",
		},
		{
			Id:     "190401",
			CityId: "1904",
			Name:   "Koba",
		},
		{
			Id:     "190402",
			CityId: "1904",
			Name:   "Pangkalan Baru",
		},
		{
			Id:     "190403",
			CityId: "1904",
			Name:   "Sungai Selan",
		},
		{
			Id:     "190404",
			CityId: "1904",
			Name:   "Simpang Katis",
		},
		{
			Id:     "190405",
			CityId: "1904",
			Name:   "Namang",
		},
		{
			Id:     "190406",
			CityId: "1904",
			Name:   "Lubuk Besar",
		},
		{
			Id:     "190501",
			CityId: "1905",
			Name:   "Mentok",
		},
		{
			Id:     "190502",
			CityId: "1905",
			Name:   "Simpang Teritip",
		},
		{
			Id:     "190503",
			CityId: "1905",
			Name:   "Jebus",
		},
		{
			Id:     "190504",
			CityId: "1905",
			Name:   "Kelapa",
		},
		{
			Id:     "190505",
			CityId: "1905",
			Name:   "Tempilang",
		},
		{
			Id:     "190506",
			CityId: "1905",
			Name:   "Parittiga",
		},
		{
			Id:     "190601",
			CityId: "1906",
			Name:   "Manggar",
		},
		{
			Id:     "190602",
			CityId: "1906",
			Name:   "Gantung",
		},
		{
			Id:     "190603",
			CityId: "1906",
			Name:   "Dendang",
		},
		{
			Id:     "190604",
			CityId: "1906",
			Name:   "Kelapa Kampit",
		},
		{
			Id:     "190605",
			CityId: "1906",
			Name:   "Damar",
		},
		{
			Id:     "190606",
			CityId: "1906",
			Name:   "Simpang Renggiang",
		},
		{
			Id:     "190607",
			CityId: "1906",
			Name:   "Simpang Pesak",
		},
		{
			Id:     "197101",
			CityId: "1971",
			Name:   "Bukit Intan",
		},
		{
			Id:     "197102",
			CityId: "1971",
			Name:   "Taman Sari",
		},
		{
			Id:     "197103",
			CityId: "1971",
			Name:   "Pangkal Balam",
		},
		{
			Id:     "197104",
			CityId: "1971",
			Name:   "Rangkui",
		},
		{
			Id:     "197105",
			CityId: "1971",
			Name:   "Gerunggang",
		},
		{
			Id:     "197106",
			CityId: "1971",
			Name:   "Gabek",
		},
		{
			Id:     "197107",
			CityId: "1971",
			Name:   "Girimaya",
		},
		{
			Id:     "210104",
			CityId: "2101",
			Name:   "Gunung Kijang",
		},
		{
			Id:     "210106",
			CityId: "2101",
			Name:   "Bintan Timur",
		},
		{
			Id:     "210107",
			CityId: "2101",
			Name:   "Bintan Utara",
		},
		{
			Id:     "210108",
			CityId: "2101",
			Name:   "Teluk Bintan",
		},
		{
			Id:     "210109",
			CityId: "2101",
			Name:   "Tambelan",
		},
		{
			Id:     "210110",
			CityId: "2101",
			Name:   "Telok Sebong",
		},
		{
			Id:     "210112",
			CityId: "2101",
			Name:   "Toapaya",
		},
		{
			Id:     "210113",
			CityId: "2101",
			Name:   "Mantang",
		},
		{
			Id:     "210114",
			CityId: "2101",
			Name:   "Bintan Pesisir",
		},
		{
			Id:     "210115",
			CityId: "2101",
			Name:   "Seri Kuala Lobam",
		},
		{
			Id:     "210201",
			CityId: "2102",
			Name:   "Moro",
		},
		{
			Id:     "210202",
			CityId: "2102",
			Name:   "Kundur",
		},
		{
			Id:     "210203",
			CityId: "2102",
			Name:   "Karimun",
		},
		{
			Id:     "210204",
			CityId: "2102",
			Name:   "Meral",
		},
		{
			Id:     "210205",
			CityId: "2102",
			Name:   "Tebing",
		},
		{
			Id:     "210206",
			CityId: "2102",
			Name:   "Buru",
		},
		{
			Id:     "210207",
			CityId: "2102",
			Name:   "Kundur Utara",
		},
		{
			Id:     "210208",
			CityId: "2102",
			Name:   "Kundur Barat",
		},
		{
			Id:     "210209",
			CityId: "2102",
			Name:   "Durai",
		},
		{
			Id:     "210210",
			CityId: "2102",
			Name:   "Meral Barat",
		},
		{
			Id:     "210211",
			CityId: "2102",
			Name:   "Ungar",
		},
		{
			Id:     "210212",
			CityId: "2102",
			Name:   "Belat",
		},
		{
			Id:     "210213",
			CityId: "2102",
			Name:   "Selat Gelam",
		},
		{
			Id:     "210214",
			CityId: "2102",
			Name:   "Sugie Besar",
		},
		{
			Id:     "210304",
			CityId: "2103",
			Name:   "Midai",
		},
		{
			Id:     "210305",
			CityId: "2103",
			Name:   "Bunguran Barat",
		},
		{
			Id:     "210306",
			CityId: "2103",
			Name:   "Serasan",
		},
		{
			Id:     "210307",
			CityId: "2103",
			Name:   "Bunguran Timur",
		},
		{
			Id:     "210308",
			CityId: "2103",
			Name:   "Bunguran Utara",
		},
		{
			Id:     "210309",
			CityId: "2103",
			Name:   "Subi",
		},
		{
			Id:     "210310",
			CityId: "2103",
			Name:   "Pulau Laut",
		},
		{
			Id:     "210311",
			CityId: "2103",
			Name:   "Pulau Tiga",
		},
		{
			Id:     "210315",
			CityId: "2103",
			Name:   "Bunguran Timur Laut",
		},
		{
			Id:     "210316",
			CityId: "2103",
			Name:   "Bunguran Tengah",
		},
		{
			Id:     "210318",
			CityId: "2103",
			Name:   "Bunguran Selatan",
		},
		{
			Id:     "210319",
			CityId: "2103",
			Name:   "Serasan Timur",
		},
		{
			Id:     "210320",
			CityId: "2103",
			Name:   "Bunguran Batubi",
		},
		{
			Id:     "210321",
			CityId: "2103",
			Name:   "Pulau Tiga Barat",
		},
		{
			Id:     "210322",
			CityId: "2103",
			Name:   "Suak Midai",
		},
		{
			Id:     "210323",
			CityId: "2103",
			Name:   "Pulau Panjang",
		},
		{
			Id:     "210324",
			CityId: "2103",
			Name:   "Pulau Seluan",
		},
		{
			Id:     "210401",
			CityId: "2104",
			Name:   "Singkep",
		},
		{
			Id:     "210402",
			CityId: "2104",
			Name:   "Lingga",
		},
		{
			Id:     "210403",
			CityId: "2104",
			Name:   "Senayang",
		},
		{
			Id:     "210404",
			CityId: "2104",
			Name:   "Singkep Barat",
		},
		{
			Id:     "210405",
			CityId: "2104",
			Name:   "Lingga Utara",
		},
		{
			Id:     "210406",
			CityId: "2104",
			Name:   "Singkep Pesisir",
		},
		{
			Id:     "210407",
			CityId: "2104",
			Name:   "Lingga Timur",
		},
		{
			Id:     "210408",
			CityId: "2104",
			Name:   "Selayar",
		},
		{
			Id:     "210409",
			CityId: "2104",
			Name:   "Singkep Selatan",
		},
		{
			Id:     "210410",
			CityId: "2104",
			Name:   "Kepulauan Posek",
		},
		{
			Id:     "210411",
			CityId: "2104",
			Name:   "Katang Bidare",
		},
		{
			Id:     "210412",
			CityId: "2104",
			Name:   "Temiang Pesisir",
		},
		{
			Id:     "210413",
			CityId: "2104",
			Name:   "Bakung Serumpun",
		},
		{
			Id:     "210501",
			CityId: "2105",
			Name:   "Siantan",
		},
		{
			Id:     "210502",
			CityId: "2105",
			Name:   "Palmatak",
		},
		{
			Id:     "210503",
			CityId: "2105",
			Name:   "Siantan Timur",
		},
		{
			Id:     "210504",
			CityId: "2105",
			Name:   "Siantan Selatan",
		},
		{
			Id:     "210505",
			CityId: "2105",
			Name:   "Jemaja Timur",
		},
		{
			Id:     "210506",
			CityId: "2105",
			Name:   "Jemaja",
		},
		{
			Id:     "210507",
			CityId: "2105",
			Name:   "Siantan Tengah",
		},
		{
			Id:     "210508",
			CityId: "2105",
			Name:   "Siantan Utara",
		},
		{
			Id:     "210509",
			CityId: "2105",
			Name:   "Jemaja Barat",
		},
		{
			Id:     "210510",
			CityId: "2105",
			Name:   "Kute Siantan",
		},
		{
			Id:     "217101",
			CityId: "2171",
			Name:   "Belakang Padang",
		},
		{
			Id:     "217102",
			CityId: "2171",
			Name:   "Batu Ampar",
		},
		{
			Id:     "217103",
			CityId: "2171",
			Name:   "Sekupang",
		},
		{
			Id:     "217104",
			CityId: "2171",
			Name:   "Nongsa",
		},
		{
			Id:     "217105",
			CityId: "2171",
			Name:   "Bulang",
		},
		{
			Id:     "217106",
			CityId: "2171",
			Name:   "Lubuk Baja",
		},
		{
			Id:     "217107",
			CityId: "2171",
			Name:   "Sei Beduk",
		},
		{
			Id:     "217108",
			CityId: "2171",
			Name:   "Galang",
		},
		{
			Id:     "217109",
			CityId: "2171",
			Name:   "Bengkong",
		},
		{
			Id:     "217110",
			CityId: "2171",
			Name:   "Batam Kota",
		},
		{
			Id:     "217111",
			CityId: "2171",
			Name:   "Sagulung",
		},
		{
			Id:     "217112",
			CityId: "2171",
			Name:   "Batu Aji",
		},
		{
			Id:     "217201",
			CityId: "2172",
			Name:   "Tanjung Pinang Barat",
		},
		{
			Id:     "217202",
			CityId: "2172",
			Name:   "Tanjung Pinang Timur",
		},
		{
			Id:     "217203",
			CityId: "2172",
			Name:   "Tanjung Pinang Kota",
		},
		{
			Id:     "217204",
			CityId: "2172",
			Name:   "Bukit Bestari",
		},
		{
			Id:     "310101",
			CityId: "3101",
			Name:   "Kepulauan Seribu Utara",
		},
		{
			Id:     "310102",
			CityId: "3101",
			Name:   "Kepulauan Seribu Selatan.",
		},
		{
			Id:     "317101",
			CityId: "3171",
			Name:   "Gambir",
		},
		{
			Id:     "317102",
			CityId: "3171",
			Name:   "Sawah Besar",
		},
		{
			Id:     "317103",
			CityId: "3171",
			Name:   "Kemayoran",
		},
		{
			Id:     "317104",
			CityId: "3171",
			Name:   "Senen",
		},
		{
			Id:     "317105",
			CityId: "3171",
			Name:   "Cempaka Putih",
		},
		{
			Id:     "317106",
			CityId: "3171",
			Name:   "Menteng",
		},
		{
			Id:     "317107",
			CityId: "3171",
			Name:   "Tanah Abang",
		},
		{
			Id:     "317108",
			CityId: "3171",
			Name:   "Johar Baru",
		},
		{
			Id:     "317201",
			CityId: "3172",
			Name:   "Penjaringan",
		},
		{
			Id:     "317202",
			CityId: "3172",
			Name:   "Tanjung Priok",
		},
		{
			Id:     "317203",
			CityId: "3172",
			Name:   "Koja",
		},
		{
			Id:     "317204",
			CityId: "3172",
			Name:   "Cilincing",
		},
		{
			Id:     "317205",
			CityId: "3172",
			Name:   "Pademangan",
		},
		{
			Id:     "317206",
			CityId: "3172",
			Name:   "Kelapa Gading",
		},
		{
			Id:     "317301",
			CityId: "3173",
			Name:   "Cengkareng",
		},
		{
			Id:     "317302",
			CityId: "3173",
			Name:   "Grogol Petamburan",
		},
		{
			Id:     "317303",
			CityId: "3173",
			Name:   "Taman Sari",
		},
		{
			Id:     "317304",
			CityId: "3173",
			Name:   "Tambora",
		},
		{
			Id:     "317305",
			CityId: "3173",
			Name:   "Kebon Jeruk",
		},
		{
			Id:     "317306",
			CityId: "3173",
			Name:   "Kalideres",
		},
		{
			Id:     "317307",
			CityId: "3173",
			Name:   "Pal Merah",
		},
		{
			Id:     "317308",
			CityId: "3173",
			Name:   "Kembangan",
		},
		{
			Id:     "317401",
			CityId: "3174",
			Name:   "Tebet",
		},
		{
			Id:     "317402",
			CityId: "3174",
			Name:   "Setiabudi",
		},
		{
			Id:     "317403",
			CityId: "3174",
			Name:   "Mampang Prapatan",
		},
		{
			Id:     "317404",
			CityId: "3174",
			Name:   "Pasar Minggu",
		},
		{
			Id:     "317405",
			CityId: "3174",
			Name:   "Kebayoran Lama",
		},
		{
			Id:     "317406",
			CityId: "3174",
			Name:   "Cilandak",
		},
		{
			Id:     "317407",
			CityId: "3174",
			Name:   "Kebayoran Baru",
		},
		{
			Id:     "317408",
			CityId: "3174",
			Name:   "Pancoran",
		},
		{
			Id:     "317409",
			CityId: "3174",
			Name:   "Jagakarsa",
		},
		{
			Id:     "317410",
			CityId: "3174",
			Name:   "Pesanggrahan",
		},
		{
			Id:     "317501",
			CityId: "3175",
			Name:   "Matraman",
		},
		{
			Id:     "317502",
			CityId: "3175",
			Name:   "Pulogadung",
		},
		{
			Id:     "317503",
			CityId: "3175",
			Name:   "Jatinegara",
		},
		{
			Id:     "317504",
			CityId: "3175",
			Name:   "Kramatjati",
		},
		{
			Id:     "317505",
			CityId: "3175",
			Name:   "Pasar Rebo",
		},
		{
			Id:     "317506",
			CityId: "3175",
			Name:   "Cakung",
		},
		{
			Id:     "317507",
			CityId: "3175",
			Name:   "Duren Sawit",
		},
		{
			Id:     "317508",
			CityId: "3175",
			Name:   "Makasar",
		},
		{
			Id:     "317509",
			CityId: "3175",
			Name:   "Ciracas",
		},
		{
			Id:     "317510",
			CityId: "3175",
			Name:   "Cipayung",
		},
		{
			Id:     "320101",
			CityId: "3201",
			Name:   "Cibinong",
		},
		{
			Id:     "320102",
			CityId: "3201",
			Name:   "Gunung Putri",
		},
		{
			Id:     "320103",
			CityId: "3201",
			Name:   "Citeureup",
		},
		{
			Id:     "320104",
			CityId: "3201",
			Name:   "Sukaraja",
		},
		{
			Id:     "320105",
			CityId: "3201",
			Name:   "Babakan Madang",
		},
		{
			Id:     "320106",
			CityId: "3201",
			Name:   "Jonggol",
		},
		{
			Id:     "320107",
			CityId: "3201",
			Name:   "Cileungsi",
		},
		{
			Id:     "320108",
			CityId: "3201",
			Name:   "Cariu",
		},
		{
			Id:     "320109",
			CityId: "3201",
			Name:   "Sukamakmur",
		},
		{
			Id:     "320110",
			CityId: "3201",
			Name:   "Parung",
		},
		{
			Id:     "320111",
			CityId: "3201",
			Name:   "Gunung Sindur",
		},
		{
			Id:     "320112",
			CityId: "3201",
			Name:   "Kemang",
		},
		{
			Id:     "320113",
			CityId: "3201",
			Name:   "Bojong Gede",
		},
		{
			Id:     "320114",
			CityId: "3201",
			Name:   "Leuwiliang",
		},
		{
			Id:     "320115",
			CityId: "3201",
			Name:   "Ciampea",
		},
		{
			Id:     "320116",
			CityId: "3201",
			Name:   "Cibungbulang",
		},
		{
			Id:     "320117",
			CityId: "3201",
			Name:   "Pamijahan",
		},
		{
			Id:     "320118",
			CityId: "3201",
			Name:   "Rumpin",
		},
		{
			Id:     "320119",
			CityId: "3201",
			Name:   "Jasinga",
		},
		{
			Id:     "320120",
			CityId: "3201",
			Name:   "Parung Panjang",
		},
		{
			Id:     "320121",
			CityId: "3201",
			Name:   "Nanggung",
		},
		{
			Id:     "320122",
			CityId: "3201",
			Name:   "Cigudeg",
		},
		{
			Id:     "320123",
			CityId: "3201",
			Name:   "Tenjo",
		},
		{
			Id:     "320124",
			CityId: "3201",
			Name:   "Ciawi",
		},
		{
			Id:     "320125",
			CityId: "3201",
			Name:   "Cisarua",
		},
		{
			Id:     "320126",
			CityId: "3201",
			Name:   "Megamendung",
		},
		{
			Id:     "320127",
			CityId: "3201",
			Name:   "Caringin",
		},
		{
			Id:     "320128",
			CityId: "3201",
			Name:   "Cijeruk",
		},
		{
			Id:     "320129",
			CityId: "3201",
			Name:   "Ciomas",
		},
		{
			Id:     "320130",
			CityId: "3201",
			Name:   "Dramaga",
		},
		{
			Id:     "320131",
			CityId: "3201",
			Name:   "Tamansari",
		},
		{
			Id:     "320132",
			CityId: "3201",
			Name:   "Klapanunggal",
		},
		{
			Id:     "320133",
			CityId: "3201",
			Name:   "Ciseeng",
		},
		{
			Id:     "320134",
			CityId: "3201",
			Name:   "Ranca Bungur",
		},
		{
			Id:     "320135",
			CityId: "3201",
			Name:   "Sukajaya",
		},
		{
			Id:     "320136",
			CityId: "3201",
			Name:   "Tanjungsari",
		},
		{
			Id:     "320137",
			CityId: "3201",
			Name:   "Tajurhalang",
		},
		{
			Id:     "320138",
			CityId: "3201",
			Name:   "Cigombong",
		},
		{
			Id:     "320139",
			CityId: "3201",
			Name:   "Leuwisadeng",
		},
		{
			Id:     "320140",
			CityId: "3201",
			Name:   "Tenjolaya",
		},
		{
			Id:     "320201",
			CityId: "3202",
			Name:   "Palabuhanratu",
		},
		{
			Id:     "320202",
			CityId: "3202",
			Name:   "Simpenan",
		},
		{
			Id:     "320203",
			CityId: "3202",
			Name:   "Cikakak",
		},
		{
			Id:     "320204",
			CityId: "3202",
			Name:   "Bantargadung",
		},
		{
			Id:     "320205",
			CityId: "3202",
			Name:   "Cisolok",
		},
		{
			Id:     "320206",
			CityId: "3202",
			Name:   "Cikidang",
		},
		{
			Id:     "320207",
			CityId: "3202",
			Name:   "Lengkong",
		},
		{
			Id:     "320208",
			CityId: "3202",
			Name:   "Jampangtengah",
		},
		{
			Id:     "320209",
			CityId: "3202",
			Name:   "Warungkiara",
		},
		{
			Id:     "320210",
			CityId: "3202",
			Name:   "Cikembar",
		},
		{
			Id:     "320211",
			CityId: "3202",
			Name:   "Cibadak",
		},
		{
			Id:     "320212",
			CityId: "3202",
			Name:   "Nagrak",
		},
		{
			Id:     "320213",
			CityId: "3202",
			Name:   "Parungkuda",
		},
		{
			Id:     "320214",
			CityId: "3202",
			Name:   "Bojonggenteng",
		},
		{
			Id:     "320215",
			CityId: "3202",
			Name:   "Parakansalak",
		},
		{
			Id:     "320216",
			CityId: "3202",
			Name:   "Cicurug",
		},
		{
			Id:     "320217",
			CityId: "3202",
			Name:   "Cidahu",
		},
		{
			Id:     "320218",
			CityId: "3202",
			Name:   "Kalapanunggal",
		},
		{
			Id:     "320219",
			CityId: "3202",
			Name:   "Kabandungan",
		},
		{
			Id:     "320220",
			CityId: "3202",
			Name:   "Waluran",
		},
		{
			Id:     "320221",
			CityId: "3202",
			Name:   "Jampangkulon",
		},
		{
			Id:     "320222",
			CityId: "3202",
			Name:   "Ciemas",
		},
		{
			Id:     "320223",
			CityId: "3202",
			Name:   "Kalibunder",
		},
		{
			Id:     "320224",
			CityId: "3202",
			Name:   "Surade",
		},
		{
			Id:     "320225",
			CityId: "3202",
			Name:   "Cibitung",
		},
		{
			Id:     "320226",
			CityId: "3202",
			Name:   "Ciracap",
		},
		{
			Id:     "320227",
			CityId: "3202",
			Name:   "Gunungguruh",
		},
		{
			Id:     "320228",
			CityId: "3202",
			Name:   "Cicantayan",
		},
		{
			Id:     "320229",
			CityId: "3202",
			Name:   "Cisaat",
		},
		{
			Id:     "320230",
			CityId: "3202",
			Name:   "Kadudampit",
		},
		{
			Id:     "320231",
			CityId: "3202",
			Name:   "Caringin",
		},
		{
			Id:     "320232",
			CityId: "3202",
			Name:   "Sukabumi",
		},
		{
			Id:     "320233",
			CityId: "3202",
			Name:   "Sukaraja",
		},
		{
			Id:     "320234",
			CityId: "3202",
			Name:   "Kebonpedes",
		},
		{
			Id:     "320235",
			CityId: "3202",
			Name:   "Cireunghas",
		},
		{
			Id:     "320236",
			CityId: "3202",
			Name:   "Sukalarang",
		},
		{
			Id:     "320237",
			CityId: "3202",
			Name:   "Pabuaran",
		},
		{
			Id:     "320238",
			CityId: "3202",
			Name:   "Purabaya",
		},
		{
			Id:     "320239",
			CityId: "3202",
			Name:   "Nyalindung",
		},
		{
			Id:     "320240",
			CityId: "3202",
			Name:   "Gegerbitung",
		},
		{
			Id:     "320241",
			CityId: "3202",
			Name:   "Sagaranten",
		},
		{
			Id:     "320242",
			CityId: "3202",
			Name:   "Curugkembar",
		},
		{
			Id:     "320243",
			CityId: "3202",
			Name:   "Cidolog",
		},
		{
			Id:     "320244",
			CityId: "3202",
			Name:   "Cidadap",
		},
		{
			Id:     "320245",
			CityId: "3202",
			Name:   "Tegalbuleud",
		},
		{
			Id:     "320246",
			CityId: "3202",
			Name:   "Cimanggu",
		},
		{
			Id:     "320247",
			CityId: "3202",
			Name:   "Ciambar",
		},
		{
			Id:     "320301",
			CityId: "3203",
			Name:   "Cianjur",
		},
		{
			Id:     "320302",
			CityId: "3203",
			Name:   "Warungkondang",
		},
		{
			Id:     "320303",
			CityId: "3203",
			Name:   "Cibeber",
		},
		{
			Id:     "320304",
			CityId: "3203",
			Name:   "Cilaku",
		},
		{
			Id:     "320305",
			CityId: "3203",
			Name:   "Ciranjang",
		},
		{
			Id:     "320306",
			CityId: "3203",
			Name:   "Bojongpicung",
		},
		{
			Id:     "320307",
			CityId: "3203",
			Name:   "Karangtengah",
		},
		{
			Id:     "320308",
			CityId: "3203",
			Name:   "Mande",
		},
		{
			Id:     "320309",
			CityId: "3203",
			Name:   "Sukaluyu",
		},
		{
			Id:     "320310",
			CityId: "3203",
			Name:   "Pacet",
		},
		{
			Id:     "320311",
			CityId: "3203",
			Name:   "Cugenang",
		},
		{
			Id:     "320312",
			CityId: "3203",
			Name:   "Cikalongkulon",
		},
		{
			Id:     "320313",
			CityId: "3203",
			Name:   "Sukaresmi",
		},
		{
			Id:     "320314",
			CityId: "3203",
			Name:   "Sukanagara",
		},
		{
			Id:     "320315",
			CityId: "3203",
			Name:   "Campaka",
		},
		{
			Id:     "320316",
			CityId: "3203",
			Name:   "Takokak",
		},
		{
			Id:     "320317",
			CityId: "3203",
			Name:   "Kadupandak",
		},
		{
			Id:     "320318",
			CityId: "3203",
			Name:   "Pagelaran",
		},
		{
			Id:     "320319",
			CityId: "3203",
			Name:   "Tanggeung",
		},
		{
			Id:     "320320",
			CityId: "3203",
			Name:   "Cibinong",
		},
		{
			Id:     "320321",
			CityId: "3203",
			Name:   "Sindangbarang",
		},
		{
			Id:     "320322",
			CityId: "3203",
			Name:   "Agrabinta",
		},
		{
			Id:     "320323",
			CityId: "3203",
			Name:   "Cidaun",
		},
		{
			Id:     "320324",
			CityId: "3203",
			Name:   "Naringgul",
		},
		{
			Id:     "320325",
			CityId: "3203",
			Name:   "Campakamulya",
		},
		{
			Id:     "320326",
			CityId: "3203",
			Name:   "Cikadu",
		},
		{
			Id:     "320327",
			CityId: "3203",
			Name:   "Gekbrong",
		},
		{
			Id:     "320328",
			CityId: "3203",
			Name:   "Cipanas",
		},
		{
			Id:     "320329",
			CityId: "3203",
			Name:   "Cijati",
		},
		{
			Id:     "320330",
			CityId: "3203",
			Name:   "Leles",
		},
		{
			Id:     "320331",
			CityId: "3203",
			Name:   "Haurwangi",
		},
		{
			Id:     "320332",
			CityId: "3203",
			Name:   "Pasirkuda",
		},
		{
			Id:     "320405",
			CityId: "3204",
			Name:   "Cileunyi",
		},
		{
			Id:     "320406",
			CityId: "3204",
			Name:   "Cimenyan",
		},
		{
			Id:     "320407",
			CityId: "3204",
			Name:   "Cilengkrang",
		},
		{
			Id:     "320408",
			CityId: "3204",
			Name:   "Bojongsoang",
		},
		{
			Id:     "320409",
			CityId: "3204",
			Name:   "Margahayu",
		},
		{
			Id:     "320410",
			CityId: "3204",
			Name:   "Margaasih",
		},
		{
			Id:     "320411",
			CityId: "3204",
			Name:   "Katapang",
		},
		{
			Id:     "320412",
			CityId: "3204",
			Name:   "Dayeuhkolot",
		},
		{
			Id:     "320413",
			CityId: "3204",
			Name:   "Banjaran",
		},
		{
			Id:     "320414",
			CityId: "3204",
			Name:   "Pameungpeuk",
		},
		{
			Id:     "320415",
			CityId: "3204",
			Name:   "Pangalengan",
		},
		{
			Id:     "320416",
			CityId: "3204",
			Name:   "Arjasari",
		},
		{
			Id:     "320417",
			CityId: "3204",
			Name:   "Cimaung",
		},
		{
			Id:     "320425",
			CityId: "3204",
			Name:   "Cicalengka",
		},
		{
			Id:     "320426",
			CityId: "3204",
			Name:   "Nagreg",
		},
		{
			Id:     "320427",
			CityId: "3204",
			Name:   "Cikancung",
		},
		{
			Id:     "320428",
			CityId: "3204",
			Name:   "Rancaekek",
		},
		{
			Id:     "320429",
			CityId: "3204",
			Name:   "Ciparay",
		},
		{
			Id:     "320430",
			CityId: "3204",
			Name:   "Pacet",
		},
		{
			Id:     "320431",
			CityId: "3204",
			Name:   "Kertasari",
		},
		{
			Id:     "320432",
			CityId: "3204",
			Name:   "Baleendah",
		},
		{
			Id:     "320433",
			CityId: "3204",
			Name:   "Majalaya",
		},
		{
			Id:     "320434",
			CityId: "3204",
			Name:   "Solokanjeruk",
		},
		{
			Id:     "320435",
			CityId: "3204",
			Name:   "Paseh",
		},
		{
			Id:     "320436",
			CityId: "3204",
			Name:   "Ibun",
		},
		{
			Id:     "320437",
			CityId: "3204",
			Name:   "Soreang",
		},
		{
			Id:     "320438",
			CityId: "3204",
			Name:   "Pasirjambu",
		},
		{
			Id:     "320439",
			CityId: "3204",
			Name:   "Ciwidey",
		},
		{
			Id:     "320440",
			CityId: "3204",
			Name:   "Rancabali",
		},
		{
			Id:     "320444",
			CityId: "3204",
			Name:   "Cangkuang",
		},
		{
			Id:     "320446",
			CityId: "3204",
			Name:   "Kutawaringin",
		},
		{
			Id:     "320501",
			CityId: "3205",
			Name:   "Garut Kota",
		},
		{
			Id:     "320502",
			CityId: "3205",
			Name:   "Karangpawitan",
		},
		{
			Id:     "320503",
			CityId: "3205",
			Name:   "Wanaraja",
		},
		{
			Id:     "320504",
			CityId: "3205",
			Name:   "Tarogong Kaler",
		},
		{
			Id:     "320505",
			CityId: "3205",
			Name:   "Tarogong Kidul",
		},
		{
			Id:     "320506",
			CityId: "3205",
			Name:   "Banyuresmi",
		},
		{
			Id:     "320507",
			CityId: "3205",
			Name:   "Samarang",
		},
		{
			Id:     "320508",
			CityId: "3205",
			Name:   "Pasirwangi",
		},
		{
			Id:     "320509",
			CityId: "3205",
			Name:   "Leles",
		},
		{
			Id:     "320510",
			CityId: "3205",
			Name:   "Kadungora",
		},
		{
			Id:     "320511",
			CityId: "3205",
			Name:   "Leuwigoong",
		},
		{
			Id:     "320512",
			CityId: "3205",
			Name:   "Cibatu",
		},
		{
			Id:     "320513",
			CityId: "3205",
			Name:   "Kersamanah",
		},
		{
			Id:     "320514",
			CityId: "3205",
			Name:   "Malangbong",
		},
		{
			Id:     "320515",
			CityId: "3205",
			Name:   "Sukawening",
		},
		{
			Id:     "320516",
			CityId: "3205",
			Name:   "Karangtengah",
		},
		{
			Id:     "320517",
			CityId: "3205",
			Name:   "Bayongbong",
		},
		{
			Id:     "320518",
			CityId: "3205",
			Name:   "Cigedug",
		},
		{
			Id:     "320519",
			CityId: "3205",
			Name:   "Cilawu",
		},
		{
			Id:     "320520",
			CityId: "3205",
			Name:   "Cisurupan",
		},
		{
			Id:     "320521",
			CityId: "3205",
			Name:   "Sukaresmi",
		},
		{
			Id:     "320522",
			CityId: "3205",
			Name:   "Cikajang",
		},
		{
			Id:     "320523",
			CityId: "3205",
			Name:   "Banjarwangi",
		},
		{
			Id:     "320524",
			CityId: "3205",
			Name:   "Singajaya",
		},
		{
			Id:     "320525",
			CityId: "3205",
			Name:   "Cihurip",
		},
		{
			Id:     "320526",
			CityId: "3205",
			Name:   "Peundeuy",
		},
		{
			Id:     "320527",
			CityId: "3205",
			Name:   "Pameungpeuk",
		},
		{
			Id:     "320528",
			CityId: "3205",
			Name:   "Cisompet",
		},
		{
			Id:     "320529",
			CityId: "3205",
			Name:   "Cibalong",
		},
		{
			Id:     "320530",
			CityId: "3205",
			Name:   "Cikelet",
		},
		{
			Id:     "320531",
			CityId: "3205",
			Name:   "Bungbulang",
		},
		{
			Id:     "320532",
			CityId: "3205",
			Name:   "Mekarmukti",
		},
		{
			Id:     "320533",
			CityId: "3205",
			Name:   "Pakenjeng",
		},
		{
			Id:     "320534",
			CityId: "3205",
			Name:   "Pamulihan",
		},
		{
			Id:     "320535",
			CityId: "3205",
			Name:   "Cisewu",
		},
		{
			Id:     "320536",
			CityId: "3205",
			Name:   "Caringin",
		},
		{
			Id:     "320537",
			CityId: "3205",
			Name:   "Talegong",
		},
		{
			Id:     "320538",
			CityId: "3205",
			Name:   "Limbangan",
		},
		{
			Id:     "320539",
			CityId: "3205",
			Name:   "Selaawi",
		},
		{
			Id:     "320540",
			CityId: "3205",
			Name:   "Cibiuk",
		},
		{
			Id:     "320541",
			CityId: "3205",
			Name:   "Pangatikan",
		},
		{
			Id:     "320542",
			CityId: "3205",
			Name:   "Sucinaraja",
		},
		{
			Id:     "320601",
			CityId: "3206",
			Name:   "Cipatujah",
		},
		{
			Id:     "320602",
			CityId: "3206",
			Name:   "Karangnunggal",
		},
		{
			Id:     "320603",
			CityId: "3206",
			Name:   "Cikalong",
		},
		{
			Id:     "320604",
			CityId: "3206",
			Name:   "Pancatengah",
		},
		{
			Id:     "320605",
			CityId: "3206",
			Name:   "Cikatomas",
		},
		{
			Id:     "320606",
			CityId: "3206",
			Name:   "Cibalong",
		},
		{
			Id:     "320607",
			CityId: "3206",
			Name:   "Parungponteng",
		},
		{
			Id:     "320608",
			CityId: "3206",
			Name:   "Bantarkalong",
		},
		{
			Id:     "320609",
			CityId: "3206",
			Name:   "Bojongasih",
		},
		{
			Id:     "320610",
			CityId: "3206",
			Name:   "Culamega",
		},
		{
			Id:     "320611",
			CityId: "3206",
			Name:   "Bojonggambir",
		},
		{
			Id:     "320612",
			CityId: "3206",
			Name:   "Sodonghilir",
		},
		{
			Id:     "320613",
			CityId: "3206",
			Name:   "Taraju",
		},
		{
			Id:     "320614",
			CityId: "3206",
			Name:   "Salawu",
		},
		{
			Id:     "320615",
			CityId: "3206",
			Name:   "Puspahiang",
		},
		{
			Id:     "320616",
			CityId: "3206",
			Name:   "Tanjungjaya",
		},
		{
			Id:     "320617",
			CityId: "3206",
			Name:   "Sukaraja",
		},
		{
			Id:     "320618",
			CityId: "3206",
			Name:   "Salopa",
		},
		{
			Id:     "320619",
			CityId: "3206",
			Name:   "Jatiwaras",
		},
		{
			Id:     "320620",
			CityId: "3206",
			Name:   "Cineam",
		},
		{
			Id:     "320621",
			CityId: "3206",
			Name:   "Karangjaya",
		},
		{
			Id:     "320622",
			CityId: "3206",
			Name:   "Manonjaya",
		},
		{
			Id:     "320623",
			CityId: "3206",
			Name:   "Gunungtanjung",
		},
		{
			Id:     "320624",
			CityId: "3206",
			Name:   "Singaparna",
		},
		{
			Id:     "320625",
			CityId: "3206",
			Name:   "Mangunreja",
		},
		{
			Id:     "320626",
			CityId: "3206",
			Name:   "Sukarame",
		},
		{
			Id:     "320627",
			CityId: "3206",
			Name:   "Cigalontang",
		},
		{
			Id:     "320628",
			CityId: "3206",
			Name:   "Leuwisari",
		},
		{
			Id:     "320629",
			CityId: "3206",
			Name:   "Padakembang",
		},
		{
			Id:     "320630",
			CityId: "3206",
			Name:   "Sariwangi",
		},
		{
			Id:     "320631",
			CityId: "3206",
			Name:   "Sukaratu",
		},
		{
			Id:     "320632",
			CityId: "3206",
			Name:   "Cisayong",
		},
		{
			Id:     "320633",
			CityId: "3206",
			Name:   "Sukahening",
		},
		{
			Id:     "320634",
			CityId: "3206",
			Name:   "Rajapolah",
		},
		{
			Id:     "320635",
			CityId: "3206",
			Name:   "Jamanis",
		},
		{
			Id:     "320636",
			CityId: "3206",
			Name:   "Ciawi",
		},
		{
			Id:     "320637",
			CityId: "3206",
			Name:   "Kadipaten",
		},
		{
			Id:     "320638",
			CityId: "3206",
			Name:   "Pagerageung",
		},
		{
			Id:     "320639",
			CityId: "3206",
			Name:   "Sukaresik",
		},
		{
			Id:     "320701",
			CityId: "3207",
			Name:   "Ciamis",
		},
		{
			Id:     "320702",
			CityId: "3207",
			Name:   "Cikoneng",
		},
		{
			Id:     "320703",
			CityId: "3207",
			Name:   "Cijeungjing",
		},
		{
			Id:     "320704",
			CityId: "3207",
			Name:   "Sadananya",
		},
		{
			Id:     "320705",
			CityId: "3207",
			Name:   "Cidolog",
		},
		{
			Id:     "320706",
			CityId: "3207",
			Name:   "Cihaurbeuti",
		},
		{
			Id:     "320707",
			CityId: "3207",
			Name:   "Panumbangan",
		},
		{
			Id:     "320708",
			CityId: "3207",
			Name:   "Panjalu",
		},
		{
			Id:     "320709",
			CityId: "3207",
			Name:   "Kawali",
		},
		{
			Id:     "320710",
			CityId: "3207",
			Name:   "Panawangan",
		},
		{
			Id:     "320711",
			CityId: "3207",
			Name:   "Cipaku",
		},
		{
			Id:     "320712",
			CityId: "3207",
			Name:   "Jatinagara",
		},
		{
			Id:     "320713",
			CityId: "3207",
			Name:   "Rajadesa",
		},
		{
			Id:     "320714",
			CityId: "3207",
			Name:   "Sukadana",
		},
		{
			Id:     "320715",
			CityId: "3207",
			Name:   "Rancah",
		},
		{
			Id:     "320716",
			CityId: "3207",
			Name:   "Tambaksari",
		},
		{
			Id:     "320717",
			CityId: "3207",
			Name:   "Lakbok",
		},
		{
			Id:     "320718",
			CityId: "3207",
			Name:   "Banjarsari",
		},
		{
			Id:     "320719",
			CityId: "3207",
			Name:   "Pamarican",
		},
		{
			Id:     "320729",
			CityId: "3207",
			Name:   "Cimaragas",
		},
		{
			Id:     "320730",
			CityId: "3207",
			Name:   "Cisaga",
		},
		{
			Id:     "320731",
			CityId: "3207",
			Name:   "Sindangkasih",
		},
		{
			Id:     "320732",
			CityId: "3207",
			Name:   "Baregbeg",
		},
		{
			Id:     "320733",
			CityId: "3207",
			Name:   "Sukamantri",
		},
		{
			Id:     "320734",
			CityId: "3207",
			Name:   "Lumbung",
		},
		{
			Id:     "320735",
			CityId: "3207",
			Name:   "Purwadadi",
		},
		{
			Id:     "320737",
			CityId: "3207",
			Name:   "Banjaranyar",
		},
		{
			Id:     "320801",
			CityId: "3208",
			Name:   "Kadugede",
		},
		{
			Id:     "320802",
			CityId: "3208",
			Name:   "Ciniru",
		},
		{
			Id:     "320803",
			CityId: "3208",
			Name:   "Subang",
		},
		{
			Id:     "320804",
			CityId: "3208",
			Name:   "Ciwaru",
		},
		{
			Id:     "320805",
			CityId: "3208",
			Name:   "Cibingbin",
		},
		{
			Id:     "320806",
			CityId: "3208",
			Name:   "Luragung",
		},
		{
			Id:     "320807",
			CityId: "3208",
			Name:   "Lebakwangi",
		},
		{
			Id:     "320808",
			CityId: "3208",
			Name:   "Garawangi",
		},
		{
			Id:     "320809",
			CityId: "3208",
			Name:   "Kuningan",
		},
		{
			Id:     "320810",
			CityId: "3208",
			Name:   "Ciawigebang",
		},
		{
			Id:     "320811",
			CityId: "3208",
			Name:   "Cidahu",
		},
		{
			Id:     "320812",
			CityId: "3208",
			Name:   "Jalaksana",
		},
		{
			Id:     "320813",
			CityId: "3208",
			Name:   "Cilimus",
		},
		{
			Id:     "320814",
			CityId: "3208",
			Name:   "Mandirancan",
		},
		{
			Id:     "320815",
			CityId: "3208",
			Name:   "Selajambe",
		},
		{
			Id:     "320816",
			CityId: "3208",
			Name:   "Kramatmulya",
		},
		{
			Id:     "320817",
			CityId: "3208",
			Name:   "Darma",
		},
		{
			Id:     "320818",
			CityId: "3208",
			Name:   "Cigugur",
		},
		{
			Id:     "320819",
			CityId: "3208",
			Name:   "Pasawahan",
		},
		{
			Id:     "320820",
			CityId: "3208",
			Name:   "Nusaherang",
		},
		{
			Id:     "320821",
			CityId: "3208",
			Name:   "Cipicung",
		},
		{
			Id:     "320822",
			CityId: "3208",
			Name:   "Pancalang",
		},
		{
			Id:     "320823",
			CityId: "3208",
			Name:   "Japara",
		},
		{
			Id:     "320824",
			CityId: "3208",
			Name:   "Cimahi",
		},
		{
			Id:     "320825",
			CityId: "3208",
			Name:   "Cilebak",
		},
		{
			Id:     "320826",
			CityId: "3208",
			Name:   "Hantara",
		},
		{
			Id:     "320827",
			CityId: "3208",
			Name:   "Kalimanggis",
		},
		{
			Id:     "320828",
			CityId: "3208",
			Name:   "Cibeureum",
		},
		{
			Id:     "320829",
			CityId: "3208",
			Name:   "Karang Kancana",
		},
		{
			Id:     "320830",
			CityId: "3208",
			Name:   "Maleber",
		},
		{
			Id:     "320831",
			CityId: "3208",
			Name:   "Sindang Agung",
		},
		{
			Id:     "320832",
			CityId: "3208",
			Name:   "Cigandamekar",
		},
		{
			Id:     "320901",
			CityId: "3209",
			Name:   "Waled",
		},
		{
			Id:     "320902",
			CityId: "3209",
			Name:   "Ciledug",
		},
		{
			Id:     "320903",
			CityId: "3209",
			Name:   "Losari",
		},
		{
			Id:     "320904",
			CityId: "3209",
			Name:   "Pabedilan",
		},
		{
			Id:     "320905",
			CityId: "3209",
			Name:   "Babakan",
		},
		{
			Id:     "320906",
			CityId: "3209",
			Name:   "Karangsembung",
		},
		{
			Id:     "320907",
			CityId: "3209",
			Name:   "Lemahabang",
		},
		{
			Id:     "320908",
			CityId: "3209",
			Name:   "Susukan Lebak",
		},
		{
			Id:     "320909",
			CityId: "3209",
			Name:   "Sedong",
		},
		{
			Id:     "320910",
			CityId: "3209",
			Name:   "Astanajapura",
		},
		{
			Id:     "320911",
			CityId: "3209",
			Name:   "Pangenan",
		},
		{
			Id:     "320912",
			CityId: "3209",
			Name:   "Mundu",
		},
		{
			Id:     "320913",
			CityId: "3209",
			Name:   "Beber",
		},
		{
			Id:     "320914",
			CityId: "3209",
			Name:   "Talun",
		},
		{
			Id:     "320915",
			CityId: "3209",
			Name:   "Sumber",
		},
		{
			Id:     "320916",
			CityId: "3209",
			Name:   "Dukupuntang",
		},
		{
			Id:     "320917",
			CityId: "3209",
			Name:   "Palimanan",
		},
		{
			Id:     "320918",
			CityId: "3209",
			Name:   "Plumbon",
		},
		{
			Id:     "320919",
			CityId: "3209",
			Name:   "Weru",
		},
		{
			Id:     "320920",
			CityId: "3209",
			Name:   "Kedawung",
		},
		{
			Id:     "320921",
			CityId: "3209",
			Name:   "Gunung Jati",
		},
		{
			Id:     "320922",
			CityId: "3209",
			Name:   "Kapetakan",
		},
		{
			Id:     "320923",
			CityId: "3209",
			Name:   "Klangenan",
		},
		{
			Id:     "320924",
			CityId: "3209",
			Name:   "Arjawinangun",
		},
		{
			Id:     "320925",
			CityId: "3209",
			Name:   "Panguragan",
		},
		{
			Id:     "320926",
			CityId: "3209",
			Name:   "Ciwaringin",
		},
		{
			Id:     "320927",
			CityId: "3209",
			Name:   "Susukan",
		},
		{
			Id:     "320928",
			CityId: "3209",
			Name:   "Gegesik",
		},
		{
			Id:     "320929",
			CityId: "3209",
			Name:   "Kaliwedi",
		},
		{
			Id:     "320930",
			CityId: "3209",
			Name:   "Gebang",
		},
		{
			Id:     "320931",
			CityId: "3209",
			Name:   "Depok",
		},
		{
			Id:     "320932",
			CityId: "3209",
			Name:   "Pasaleman",
		},
		{
			Id:     "320933",
			CityId: "3209",
			Name:   "Pabuaran",
		},
		{
			Id:     "320934",
			CityId: "3209",
			Name:   "Karangwareng",
		},
		{
			Id:     "320935",
			CityId: "3209",
			Name:   "Tengah Tani",
		},
		{
			Id:     "320936",
			CityId: "3209",
			Name:   "Plered",
		},
		{
			Id:     "320937",
			CityId: "3209",
			Name:   "Gempol",
		},
		{
			Id:     "320938",
			CityId: "3209",
			Name:   "Greged",
		},
		{
			Id:     "320939",
			CityId: "3209",
			Name:   "Suranenggala",
		},
		{
			Id:     "320940",
			CityId: "3209",
			Name:   "Jamblang",
		},
		{
			Id:     "321001",
			CityId: "3210",
			Name:   "Lemahsugih",
		},
		{
			Id:     "321002",
			CityId: "3210",
			Name:   "Bantarujeg",
		},
		{
			Id:     "321003",
			CityId: "3210",
			Name:   "Cikijing",
		},
		{
			Id:     "321004",
			CityId: "3210",
			Name:   "Talaga",
		},
		{
			Id:     "321005",
			CityId: "3210",
			Name:   "Argapura",
		},
		{
			Id:     "321006",
			CityId: "3210",
			Name:   "Maja",
		},
		{
			Id:     "321007",
			CityId: "3210",
			Name:   "Majalengka",
		},
		{
			Id:     "321008",
			CityId: "3210",
			Name:   "Sukahaji",
		},
		{
			Id:     "321009",
			CityId: "3210",
			Name:   "Rajagaluh",
		},
		{
			Id:     "321010",
			CityId: "3210",
			Name:   "Leuwimunding",
		},
		{
			Id:     "321011",
			CityId: "3210",
			Name:   "Jatiwangi",
		},
		{
			Id:     "321012",
			CityId: "3210",
			Name:   "Dawuan",
		},
		{
			Id:     "321013",
			CityId: "3210",
			Name:   "Kadipaten",
		},
		{
			Id:     "321014",
			CityId: "3210",
			Name:   "Kertajati",
		},
		{
			Id:     "321015",
			CityId: "3210",
			Name:   "Jatitujuh",
		},
		{
			Id:     "321016",
			CityId: "3210",
			Name:   "Ligung",
		},
		{
			Id:     "321017",
			CityId: "3210",
			Name:   "Sumberjaya",
		},
		{
			Id:     "321018",
			CityId: "3210",
			Name:   "Panyingkiran",
		},
		{
			Id:     "321019",
			CityId: "3210",
			Name:   "Palasah",
		},
		{
			Id:     "321020",
			CityId: "3210",
			Name:   "Cigasong",
		},
		{
			Id:     "321021",
			CityId: "3210",
			Name:   "Sindangwangi",
		},
		{
			Id:     "321022",
			CityId: "3210",
			Name:   "Banjaran",
		},
		{
			Id:     "321023",
			CityId: "3210",
			Name:   "Cingambul",
		},
		{
			Id:     "321024",
			CityId: "3210",
			Name:   "Kasokandel",
		},
		{
			Id:     "321025",
			CityId: "3210",
			Name:   "Sindang",
		},
		{
			Id:     "321026",
			CityId: "3210",
			Name:   "Malausma",
		},
		{
			Id:     "321101",
			CityId: "3211",
			Name:   "Wado",
		},
		{
			Id:     "321102",
			CityId: "3211",
			Name:   "Jatinunggal",
		},
		{
			Id:     "321103",
			CityId: "3211",
			Name:   "Darmaraja",
		},
		{
			Id:     "321104",
			CityId: "3211",
			Name:   "Cibugel",
		},
		{
			Id:     "321105",
			CityId: "3211",
			Name:   "Cisitu",
		},
		{
			Id:     "321106",
			CityId: "3211",
			Name:   "Situraja",
		},
		{
			Id:     "321107",
			CityId: "3211",
			Name:   "Conggeang",
		},
		{
			Id:     "321108",
			CityId: "3211",
			Name:   "Paseh",
		},
		{
			Id:     "321109",
			CityId: "3211",
			Name:   "Surian",
		},
		{
			Id:     "321110",
			CityId: "3211",
			Name:   "Buahdua",
		},
		{
			Id:     "321111",
			CityId: "3211",
			Name:   "Tanjungsari",
		},
		{
			Id:     "321112",
			CityId: "3211",
			Name:   "Sukasari",
		},
		{
			Id:     "321113",
			CityId: "3211",
			Name:   "Pamulihan",
		},
		{
			Id:     "321114",
			CityId: "3211",
			Name:   "Cimanggung",
		},
		{
			Id:     "321115",
			CityId: "3211",
			Name:   "Jatinangor",
		},
		{
			Id:     "321116",
			CityId: "3211",
			Name:   "Rancakalong",
		},
		{
			Id:     "321117",
			CityId: "3211",
			Name:   "Sumedang Selatan",
		},
		{
			Id:     "321118",
			CityId: "3211",
			Name:   "Sumedang Utara",
		},
		{
			Id:     "321119",
			CityId: "3211",
			Name:   "Ganeas",
		},
		{
			Id:     "321120",
			CityId: "3211",
			Name:   "Tanjungkerta",
		},
		{
			Id:     "321121",
			CityId: "3211",
			Name:   "Tanjungmedar",
		},
		{
			Id:     "321122",
			CityId: "3211",
			Name:   "Cimalaka",
		},
		{
			Id:     "321123",
			CityId: "3211",
			Name:   "Cisarua",
		},
		{
			Id:     "321124",
			CityId: "3211",
			Name:   "Tomo",
		},
		{
			Id:     "321125",
			CityId: "3211",
			Name:   "Ujungjaya",
		},
		{
			Id:     "321126",
			CityId: "3211",
			Name:   "Jatigede",
		},
		{
			Id:     "321201",
			CityId: "3212",
			Name:   "Haurgeulis",
		},
		{
			Id:     "321202",
			CityId: "3212",
			Name:   "Kroya",
		},
		{
			Id:     "321203",
			CityId: "3212",
			Name:   "Gabuswetan",
		},
		{
			Id:     "321204",
			CityId: "3212",
			Name:   "Cikedung",
		},
		{
			Id:     "321205",
			CityId: "3212",
			Name:   "Lelea",
		},
		{
			Id:     "321206",
			CityId: "3212",
			Name:   "Bangodua",
		},
		{
			Id:     "321207",
			CityId: "3212",
			Name:   "Widasari",
		},
		{
			Id:     "321208",
			CityId: "3212",
			Name:   "Kertasemaya",
		},
		{
			Id:     "321209",
			CityId: "3212",
			Name:   "Krangkeng",
		},
		{
			Id:     "321210",
			CityId: "3212",
			Name:   "Karangampel",
		},
		{
			Id:     "321211",
			CityId: "3212",
			Name:   "Juntinyuat",
		},
		{
			Id:     "321212",
			CityId: "3212",
			Name:   "Sliyeg",
		},
		{
			Id:     "321213",
			CityId: "3212",
			Name:   "Jatibarang",
		},
		{
			Id:     "321214",
			CityId: "3212",
			Name:   "Balongan",
		},
		{
			Id:     "321215",
			CityId: "3212",
			Name:   "Indramayu",
		},
		{
			Id:     "321216",
			CityId: "3212",
			Name:   "Sindang",
		},
		{
			Id:     "321217",
			CityId: "3212",
			Name:   "Cantigi",
		},
		{
			Id:     "321218",
			CityId: "3212",
			Name:   "Lohbener",
		},
		{
			Id:     "321219",
			CityId: "3212",
			Name:   "Arahan",
		},
		{
			Id:     "321220",
			CityId: "3212",
			Name:   "Losarang",
		},
		{
			Id:     "321221",
			CityId: "3212",
			Name:   "Kandanghaur",
		},
		{
			Id:     "321222",
			CityId: "3212",
			Name:   "Bongas",
		},
		{
			Id:     "321223",
			CityId: "3212",
			Name:   "Anjatan",
		},
		{
			Id:     "321224",
			CityId: "3212",
			Name:   "Sukra",
		},
		{
			Id:     "321225",
			CityId: "3212",
			Name:   "Gantar",
		},
		{
			Id:     "321226",
			CityId: "3212",
			Name:   "Trisi",
		},
		{
			Id:     "321227",
			CityId: "3212",
			Name:   "Sukagumiwang",
		},
		{
			Id:     "321228",
			CityId: "3212",
			Name:   "Kedokan Bunder",
		},
		{
			Id:     "321229",
			CityId: "3212",
			Name:   "Pasekan",
		},
		{
			Id:     "321230",
			CityId: "3212",
			Name:   "Tukdana",
		},
		{
			Id:     "321231",
			CityId: "3212",
			Name:   "Patrol",
		},
		{
			Id:     "321301",
			CityId: "3213",
			Name:   "Sagalaherang",
		},
		{
			Id:     "321302",
			CityId: "3213",
			Name:   "Cisalak",
		},
		{
			Id:     "321303",
			CityId: "3213",
			Name:   "Subang",
		},
		{
			Id:     "321304",
			CityId: "3213",
			Name:   "Kalijati",
		},
		{
			Id:     "321305",
			CityId: "3213",
			Name:   "Pabuaran",
		},
		{
			Id:     "321306",
			CityId: "3213",
			Name:   "Purwadadi",
		},
		{
			Id:     "321307",
			CityId: "3213",
			Name:   "Pagaden",
		},
		{
			Id:     "321308",
			CityId: "3213",
			Name:   "Binong",
		},
		{
			Id:     "321309",
			CityId: "3213",
			Name:   "Ciasem",
		},
		{
			Id:     "321310",
			CityId: "3213",
			Name:   "Pusakanagara",
		},
		{
			Id:     "321311",
			CityId: "3213",
			Name:   "Pamanukan",
		},
		{
			Id:     "321312",
			CityId: "3213",
			Name:   "Jalancagak",
		},
		{
			Id:     "321313",
			CityId: "3213",
			Name:   "Blanakan",
		},
		{
			Id:     "321314",
			CityId: "3213",
			Name:   "Tanjungsiang",
		},
		{
			Id:     "321315",
			CityId: "3213",
			Name:   "Compreng",
		},
		{
			Id:     "321316",
			CityId: "3213",
			Name:   "Patokbeusi",
		},
		{
			Id:     "321317",
			CityId: "3213",
			Name:   "Cibogo",
		},
		{
			Id:     "321318",
			CityId: "3213",
			Name:   "Cipunagara",
		},
		{
			Id:     "321319",
			CityId: "3213",
			Name:   "Cijambe",
		},
		{
			Id:     "321320",
			CityId: "3213",
			Name:   "Cipeundeuy",
		},
		{
			Id:     "321321",
			CityId: "3213",
			Name:   "Legonkulon",
		},
		{
			Id:     "321322",
			CityId: "3213",
			Name:   "Cikaum",
		},
		{
			Id:     "321323",
			CityId: "3213",
			Name:   "Serangpanjang",
		},
		{
			Id:     "321324",
			CityId: "3213",
			Name:   "Sukasari",
		},
		{
			Id:     "321325",
			CityId: "3213",
			Name:   "Tambakdahan",
		},
		{
			Id:     "321326",
			CityId: "3213",
			Name:   "Kasomalang",
		},
		{
			Id:     "321327",
			CityId: "3213",
			Name:   "Dawuan",
		},
		{
			Id:     "321328",
			CityId: "3213",
			Name:   "Pagaden Barat",
		},
		{
			Id:     "321329",
			CityId: "3213",
			Name:   "Ciater",
		},
		{
			Id:     "321330",
			CityId: "3213",
			Name:   "Pusakajaya",
		},
		{
			Id:     "321401",
			CityId: "3214",
			Name:   "Purwakarta",
		},
		{
			Id:     "321402",
			CityId: "3214",
			Name:   "Campaka",
		},
		{
			Id:     "321403",
			CityId: "3214",
			Name:   "Jatiluhur",
		},
		{
			Id:     "321404",
			CityId: "3214",
			Name:   "Plered",
		},
		{
			Id:     "321405",
			CityId: "3214",
			Name:   "Sukatani",
		},
		{
			Id:     "321406",
			CityId: "3214",
			Name:   "Darangdan",
		},
		{
			Id:     "321407",
			CityId: "3214",
			Name:   "Maniis",
		},
		{
			Id:     "321408",
			CityId: "3214",
			Name:   "Tegalwaru",
		},
		{
			Id:     "321409",
			CityId: "3214",
			Name:   "Wanayasa",
		},
		{
			Id:     "321410",
			CityId: "3214",
			Name:   "Pasawahan",
		},
		{
			Id:     "321411",
			CityId: "3214",
			Name:   "Bojong",
		},
		{
			Id:     "321412",
			CityId: "3214",
			Name:   "Babakancikao",
		},
		{
			Id:     "321413",
			CityId: "3214",
			Name:   "Bungursari",
		},
		{
			Id:     "321414",
			CityId: "3214",
			Name:   "Cibatu",
		},
		{
			Id:     "321415",
			CityId: "3214",
			Name:   "Sukasari",
		},
		{
			Id:     "321416",
			CityId: "3214",
			Name:   "Pondoksalam",
		},
		{
			Id:     "321417",
			CityId: "3214",
			Name:   "Kiarapedes",
		},
		{
			Id:     "321501",
			CityId: "3215",
			Name:   "Karawang Barat",
		},
		{
			Id:     "321502",
			CityId: "3215",
			Name:   "Pangkalan",
		},
		{
			Id:     "321503",
			CityId: "3215",
			Name:   "Telukjambe Timur",
		},
		{
			Id:     "321504",
			CityId: "3215",
			Name:   "Ciampel",
		},
		{
			Id:     "321505",
			CityId: "3215",
			Name:   "Klari",
		},
		{
			Id:     "321506",
			CityId: "3215",
			Name:   "Rengasdengklok",
		},
		{
			Id:     "321507",
			CityId: "3215",
			Name:   "Kutawaluya",
		},
		{
			Id:     "321508",
			CityId: "3215",
			Name:   "Batujaya",
		},
		{
			Id:     "321509",
			CityId: "3215",
			Name:   "Tirtajaya",
		},
		{
			Id:     "321510",
			CityId: "3215",
			Name:   "Pedes",
		},
		{
			Id:     "321511",
			CityId: "3215",
			Name:   "Cibuaya",
		},
		{
			Id:     "321512",
			CityId: "3215",
			Name:   "Pakisjaya",
		},
		{
			Id:     "321513",
			CityId: "3215",
			Name:   "Cikampek",
		},
		{
			Id:     "321514",
			CityId: "3215",
			Name:   "Jatisari",
		},
		{
			Id:     "321515",
			CityId: "3215",
			Name:   "Cilamaya Wetan",
		},
		{
			Id:     "321516",
			CityId: "3215",
			Name:   "Tirtamulya",
		},
		{
			Id:     "321517",
			CityId: "3215",
			Name:   "Telagasari",
		},
		{
			Id:     "321518",
			CityId: "3215",
			Name:   "Rawamerta",
		},
		{
			Id:     "321519",
			CityId: "3215",
			Name:   "Lemahabang",
		},
		{
			Id:     "321520",
			CityId: "3215",
			Name:   "Tempuran",
		},
		{
			Id:     "321521",
			CityId: "3215",
			Name:   "Majalaya",
		},
		{
			Id:     "321522",
			CityId: "3215",
			Name:   "Jayakerta",
		},
		{
			Id:     "321523",
			CityId: "3215",
			Name:   "Cilamaya Kulon",
		},
		{
			Id:     "321524",
			CityId: "3215",
			Name:   "Banyusari",
		},
		{
			Id:     "321525",
			CityId: "3215",
			Name:   "Kota Baru",
		},
		{
			Id:     "321526",
			CityId: "3215",
			Name:   "Karawang Timur",
		},
		{
			Id:     "321527",
			CityId: "3215",
			Name:   "Telukjambe Barat",
		},
		{
			Id:     "321528",
			CityId: "3215",
			Name:   "Tegalwaru",
		},
		{
			Id:     "321529",
			CityId: "3215",
			Name:   "Purwasari",
		},
		{
			Id:     "321530",
			CityId: "3215",
			Name:   "Cilebar",
		},
		{
			Id:     "321601",
			CityId: "3216",
			Name:   "Tarumajaya",
		},
		{
			Id:     "321602",
			CityId: "3216",
			Name:   "Babelan",
		},
		{
			Id:     "321603",
			CityId: "3216",
			Name:   "Sukawangi",
		},
		{
			Id:     "321604",
			CityId: "3216",
			Name:   "Tambelang",
		},
		{
			Id:     "321605",
			CityId: "3216",
			Name:   "Tambun Utara",
		},
		{
			Id:     "321606",
			CityId: "3216",
			Name:   "Tambun Selatan",
		},
		{
			Id:     "321607",
			CityId: "3216",
			Name:   "Cibitung",
		},
		{
			Id:     "321608",
			CityId: "3216",
			Name:   "Cikarang Barat",
		},
		{
			Id:     "321609",
			CityId: "3216",
			Name:   "Cikarang Utara",
		},
		{
			Id:     "321610",
			CityId: "3216",
			Name:   "Karang Bahagia",
		},
		{
			Id:     "321611",
			CityId: "3216",
			Name:   "Cikarang Timur",
		},
		{
			Id:     "321612",
			CityId: "3216",
			Name:   "Kedung Waringin",
		},
		{
			Id:     "321613",
			CityId: "3216",
			Name:   "Pebayuran",
		},
		{
			Id:     "321614",
			CityId: "3216",
			Name:   "Sukakarya",
		},
		{
			Id:     "321615",
			CityId: "3216",
			Name:   "Sukatani",
		},
		{
			Id:     "321616",
			CityId: "3216",
			Name:   "Cabangbungin",
		},
		{
			Id:     "321617",
			CityId: "3216",
			Name:   "Muaragembong",
		},
		{
			Id:     "321618",
			CityId: "3216",
			Name:   "Setu",
		},
		{
			Id:     "321619",
			CityId: "3216",
			Name:   "Cikarang Selatan",
		},
		{
			Id:     "321620",
			CityId: "3216",
			Name:   "Cikarang Pusat",
		},
		{
			Id:     "321621",
			CityId: "3216",
			Name:   "Serang Baru",
		},
		{
			Id:     "321622",
			CityId: "3216",
			Name:   "Cibarusah",
		},
		{
			Id:     "321623",
			CityId: "3216",
			Name:   "Bojongmangu",
		},
		{
			Id:     "321701",
			CityId: "3217",
			Name:   "Lembang",
		},
		{
			Id:     "321702",
			CityId: "3217",
			Name:   "Parongpong",
		},
		{
			Id:     "321703",
			CityId: "3217",
			Name:   "Cisarua",
		},
		{
			Id:     "321704",
			CityId: "3217",
			Name:   "Cikalongwetan",
		},
		{
			Id:     "321705",
			CityId: "3217",
			Name:   "Cipeundeuy",
		},
		{
			Id:     "321706",
			CityId: "3217",
			Name:   "Ngamprah",
		},
		{
			Id:     "321707",
			CityId: "3217",
			Name:   "Cipatat",
		},
		{
			Id:     "321708",
			CityId: "3217",
			Name:   "Padalarang",
		},
		{
			Id:     "321709",
			CityId: "3217",
			Name:   "Batujajar",
		},
		{
			Id:     "321710",
			CityId: "3217",
			Name:   "Cihampelas",
		},
		{
			Id:     "321711",
			CityId: "3217",
			Name:   "Cililin",
		},
		{
			Id:     "321712",
			CityId: "3217",
			Name:   "Cipongkor",
		},
		{
			Id:     "321713",
			CityId: "3217",
			Name:   "Rongga",
		},
		{
			Id:     "321714",
			CityId: "3217",
			Name:   "Sindangkerta",
		},
		{
			Id:     "321715",
			CityId: "3217",
			Name:   "Gununghalu",
		},
		{
			Id:     "321716",
			CityId: "3217",
			Name:   "Saguling",
		},
		{
			Id:     "321801",
			CityId: "3218",
			Name:   "Parigi",
		},
		{
			Id:     "321802",
			CityId: "3218",
			Name:   "Cijulang",
		},
		{
			Id:     "321803",
			CityId: "3218",
			Name:   "Cimerak",
		},
		{
			Id:     "321804",
			CityId: "3218",
			Name:   "Cigugur",
		},
		{
			Id:     "321805",
			CityId: "3218",
			Name:   "Langkaplancar",
		},
		{
			Id:     "321806",
			CityId: "3218",
			Name:   "Mangunjaya",
		},
		{
			Id:     "321807",
			CityId: "3218",
			Name:   "Padaherang",
		},
		{
			Id:     "321808",
			CityId: "3218",
			Name:   "Kalipucang",
		},
		{
			Id:     "321809",
			CityId: "3218",
			Name:   "Pangandaran",
		},
		{
			Id:     "321810",
			CityId: "3218",
			Name:   "Sidamulih",
		},
		{
			Id:     "327101",
			CityId: "3271",
			Name:   "Bogor Selatan",
		},
		{
			Id:     "327102",
			CityId: "3271",
			Name:   "Bogor Timur",
		},
		{
			Id:     "327103",
			CityId: "3271",
			Name:   "Bogor Tengah",
		},
		{
			Id:     "327104",
			CityId: "3271",
			Name:   "Bogor Barat",
		},
		{
			Id:     "327105",
			CityId: "3271",
			Name:   "Bogor Utara",
		},
		{
			Id:     "327106",
			CityId: "3271",
			Name:   "Tanah Sareal",
		},
		{
			Id:     "327201",
			CityId: "3272",
			Name:   "Gunung Puyuh",
		},
		{
			Id:     "327202",
			CityId: "3272",
			Name:   "Cikole",
		},
		{
			Id:     "327203",
			CityId: "3272",
			Name:   "Citamiang",
		},
		{
			Id:     "327204",
			CityId: "3272",
			Name:   "Warudoyong",
		},
		{
			Id:     "327205",
			CityId: "3272",
			Name:   "Baros",
		},
		{
			Id:     "327206",
			CityId: "3272",
			Name:   "Lembursitu",
		},
		{
			Id:     "327207",
			CityId: "3272",
			Name:   "Cibeureum",
		},
		{
			Id:     "327301",
			CityId: "3273",
			Name:   "Sukasari",
		},
		{
			Id:     "327302",
			CityId: "3273",
			Name:   "Coblong",
		},
		{
			Id:     "327303",
			CityId: "3273",
			Name:   "Babakan Ciparay",
		},
		{
			Id:     "327304",
			CityId: "3273",
			Name:   "Bojongloa Kaler",
		},
		{
			Id:     "327305",
			CityId: "3273",
			Name:   "Andir",
		},
		{
			Id:     "327306",
			CityId: "3273",
			Name:   "Cicendo",
		},
		{
			Id:     "327307",
			CityId: "3273",
			Name:   "Sukajadi",
		},
		{
			Id:     "327308",
			CityId: "3273",
			Name:   "Cidadap",
		},
		{
			Id:     "327309",
			CityId: "3273",
			Name:   "Bandung Wetan",
		},
		{
			Id:     "327310",
			CityId: "3273",
			Name:   "Astana Anyar",
		},
		{
			Id:     "327311",
			CityId: "3273",
			Name:   "Regol",
		},
		{
			Id:     "327312",
			CityId: "3273",
			Name:   "Batununggal",
		},
		{
			Id:     "327313",
			CityId: "3273",
			Name:   "Lengkong",
		},
		{
			Id:     "327314",
			CityId: "3273",
			Name:   "Cibeunying Kidul",
		},
		{
			Id:     "327315",
			CityId: "3273",
			Name:   "Bandung Kulon",
		},
		{
			Id:     "327316",
			CityId: "3273",
			Name:   "Kiaracondong",
		},
		{
			Id:     "327317",
			CityId: "3273",
			Name:   "Bojongloa Kidul",
		},
		{
			Id:     "327318",
			CityId: "3273",
			Name:   "Cibeunying Kaler",
		},
		{
			Id:     "327319",
			CityId: "3273",
			Name:   "Sumur Bandung",
		},
		{
			Id:     "327320",
			CityId: "3273",
			Name:   "Antapani",
		},
		{
			Id:     "327321",
			CityId: "3273",
			Name:   "Bandung Kidul",
		},
		{
			Id:     "327322",
			CityId: "3273",
			Name:   "Buahbatu",
		},
		{
			Id:     "327323",
			CityId: "3273",
			Name:   "Rancasari",
		},
		{
			Id:     "327324",
			CityId: "3273",
			Name:   "Arcamanik",
		},
		{
			Id:     "327325",
			CityId: "3273",
			Name:   "Cibiru",
		},
		{
			Id:     "327326",
			CityId: "3273",
			Name:   "Ujungberung",
		},
		{
			Id:     "327327",
			CityId: "3273",
			Name:   "Gedebage",
		},
		{
			Id:     "327328",
			CityId: "3273",
			Name:   "Panyileukan",
		},
		{
			Id:     "327329",
			CityId: "3273",
			Name:   "Cinambo",
		},
		{
			Id:     "327330",
			CityId: "3273",
			Name:   "Mandalajati",
		},
		{
			Id:     "327401",
			CityId: "3274",
			Name:   "Kejaksan",
		},
		{
			Id:     "327402",
			CityId: "3274",
			Name:   "Lemahwungkuk",
		},
		{
			Id:     "327403",
			CityId: "3274",
			Name:   "Harjamukti",
		},
		{
			Id:     "327404",
			CityId: "3274",
			Name:   "Pekalipan",
		},
		{
			Id:     "327405",
			CityId: "3274",
			Name:   "Kesambi",
		},
		{
			Id:     "327501",
			CityId: "3275",
			Name:   "Bekasi Timur",
		},
		{
			Id:     "327502",
			CityId: "3275",
			Name:   "Bekasi Barat",
		},
		{
			Id:     "327503",
			CityId: "3275",
			Name:   "Bekasi Utara",
		},
		{
			Id:     "327504",
			CityId: "3275",
			Name:   "Bekasi Selatan",
		},
		{
			Id:     "327505",
			CityId: "3275",
			Name:   "Rawalumbu",
		},
		{
			Id:     "327506",
			CityId: "3275",
			Name:   "Medansatria",
		},
		{
			Id:     "327507",
			CityId: "3275",
			Name:   "Bantargebang",
		},
		{
			Id:     "327508",
			CityId: "3275",
			Name:   "Pondokgede",
		},
		{
			Id:     "327509",
			CityId: "3275",
			Name:   "Jatiasih",
		},
		{
			Id:     "327510",
			CityId: "3275",
			Name:   "Jatisampurna",
		},
		{
			Id:     "327511",
			CityId: "3275",
			Name:   "Mustikajaya",
		},
		{
			Id:     "327512",
			CityId: "3275",
			Name:   "Pondokmelati",
		},
		{
			Id:     "327601",
			CityId: "3276",
			Name:   "Pancoran Mas",
		},
		{
			Id:     "327602",
			CityId: "3276",
			Name:   "Cimanggis",
		},
		{
			Id:     "327603",
			CityId: "3276",
			Name:   "Sawangan",
		},
		{
			Id:     "327604",
			CityId: "3276",
			Name:   "Limo",
		},
		{
			Id:     "327605",
			CityId: "3276",
			Name:   "Sukmajaya",
		},
		{
			Id:     "327606",
			CityId: "3276",
			Name:   "Beji",
		},
		{
			Id:     "327607",
			CityId: "3276",
			Name:   "Cipayung",
		},
		{
			Id:     "327608",
			CityId: "3276",
			Name:   "Cilodong",
		},
		{
			Id:     "327609",
			CityId: "3276",
			Name:   "Cinere",
		},
		{
			Id:     "327610",
			CityId: "3276",
			Name:   "Tapos",
		},
		{
			Id:     "327611",
			CityId: "3276",
			Name:   "Bojongsari",
		},
		{
			Id:     "327701",
			CityId: "3277",
			Name:   "Cimahi Selatan",
		},
		{
			Id:     "327702",
			CityId: "3277",
			Name:   "Cimahi Tengah",
		},
		{
			Id:     "327703",
			CityId: "3277",
			Name:   "Cimahi Utara",
		},
		{
			Id:     "327801",
			CityId: "3278",
			Name:   "Cihideung",
		},
		{
			Id:     "327802",
			CityId: "3278",
			Name:   "Cipedes",
		},
		{
			Id:     "327803",
			CityId: "3278",
			Name:   "Tawang",
		},
		{
			Id:     "327804",
			CityId: "3278",
			Name:   "Indihiang",
		},
		{
			Id:     "327805",
			CityId: "3278",
			Name:   "Kawalu",
		},
		{
			Id:     "327806",
			CityId: "3278",
			Name:   "Cibeureum",
		},
		{
			Id:     "327807",
			CityId: "3278",
			Name:   "Tamansari",
		},
		{
			Id:     "327808",
			CityId: "3278",
			Name:   "Mangkubumi",
		},
		{
			Id:     "327809",
			CityId: "3278",
			Name:   "Bungursari",
		},
		{
			Id:     "327810",
			CityId: "3278",
			Name:   "Purbaratu",
		},
		{
			Id:     "327901",
			CityId: "3279",
			Name:   "Banjar",
		},
		{
			Id:     "327902",
			CityId: "3279",
			Name:   "Pataruman",
		},
		{
			Id:     "327903",
			CityId: "3279",
			Name:   "Purwaharja",
		},
		{
			Id:     "327904",
			CityId: "3279",
			Name:   "Langensari",
		},
		{
			Id:     "330101",
			CityId: "3301",
			Name:   "Kedungreja",
		},
		{
			Id:     "330102",
			CityId: "3301",
			Name:   "Kesugihan",
		},
		{
			Id:     "330103",
			CityId: "3301",
			Name:   "Adipala",
		},
		{
			Id:     "330104",
			CityId: "3301",
			Name:   "Binangun",
		},
		{
			Id:     "330105",
			CityId: "3301",
			Name:   "Nusawungu",
		},
		{
			Id:     "330106",
			CityId: "3301",
			Name:   "Kroya",
		},
		{
			Id:     "330107",
			CityId: "3301",
			Name:   "Maos",
		},
		{
			Id:     "330108",
			CityId: "3301",
			Name:   "Jeruklegi",
		},
		{
			Id:     "330109",
			CityId: "3301",
			Name:   "Kawunganten",
		},
		{
			Id:     "330110",
			CityId: "3301",
			Name:   "Gandrungmangu",
		},
		{
			Id:     "330111",
			CityId: "3301",
			Name:   "Sidareja",
		},
		{
			Id:     "330112",
			CityId: "3301",
			Name:   "Karangpucung",
		},
		{
			Id:     "330113",
			CityId: "3301",
			Name:   "Cimanggu",
		},
		{
			Id:     "330114",
			CityId: "3301",
			Name:   "Majenang",
		},
		{
			Id:     "330115",
			CityId: "3301",
			Name:   "Wanareja",
		},
		{
			Id:     "330116",
			CityId: "3301",
			Name:   "Dayeuhluhur",
		},
		{
			Id:     "330117",
			CityId: "3301",
			Name:   "Sampang",
		},
		{
			Id:     "330118",
			CityId: "3301",
			Name:   "Cipari",
		},
		{
			Id:     "330119",
			CityId: "3301",
			Name:   "Patimuan",
		},
		{
			Id:     "330120",
			CityId: "3301",
			Name:   "Bantarsari",
		},
		{
			Id:     "330121",
			CityId: "3301",
			Name:   "Cilacap Selatan",
		},
		{
			Id:     "330122",
			CityId: "3301",
			Name:   "Cilacap Tengah",
		},
		{
			Id:     "330123",
			CityId: "3301",
			Name:   "Cilacap Utara",
		},
		{
			Id:     "330124",
			CityId: "3301",
			Name:   "Kampung Laut",
		},
		{
			Id:     "330201",
			CityId: "3302",
			Name:   "Lumbir",
		},
		{
			Id:     "330202",
			CityId: "3302",
			Name:   "Wangon",
		},
		{
			Id:     "330203",
			CityId: "3302",
			Name:   "Jatilawang",
		},
		{
			Id:     "330204",
			CityId: "3302",
			Name:   "Rawalo",
		},
		{
			Id:     "330205",
			CityId: "3302",
			Name:   "Kebasen",
		},
		{
			Id:     "330206",
			CityId: "3302",
			Name:   "Kemranjen",
		},
		{
			Id:     "330207",
			CityId: "3302",
			Name:   "Sumpiuh",
		},
		{
			Id:     "330208",
			CityId: "3302",
			Name:   "Tambak",
		},
		{
			Id:     "330209",
			CityId: "3302",
			Name:   "Somagede",
		},
		{
			Id:     "330210",
			CityId: "3302",
			Name:   "Kalibagor",
		},
		{
			Id:     "330211",
			CityId: "3302",
			Name:   "Banyumas",
		},
		{
			Id:     "330212",
			CityId: "3302",
			Name:   "Patikraja",
		},
		{
			Id:     "330213",
			CityId: "3302",
			Name:   "Purwojati",
		},
		{
			Id:     "330214",
			CityId: "3302",
			Name:   "Ajibarang",
		},
		{
			Id:     "330215",
			CityId: "3302",
			Name:   "Gumelar",
		},
		{
			Id:     "330216",
			CityId: "3302",
			Name:   "Pekuncen",
		},
		{
			Id:     "330217",
			CityId: "3302",
			Name:   "Cilongok",
		},
		{
			Id:     "330218",
			CityId: "3302",
			Name:   "Karanglewas",
		},
		{
			Id:     "330219",
			CityId: "3302",
			Name:   "Sokaraja",
		},
		{
			Id:     "330220",
			CityId: "3302",
			Name:   "Kembaran",
		},
		{
			Id:     "330221",
			CityId: "3302",
			Name:   "Sumbang",
		},
		{
			Id:     "330222",
			CityId: "3302",
			Name:   "Baturraden",
		},
		{
			Id:     "330223",
			CityId: "3302",
			Name:   "Kedungbanteng",
		},
		{
			Id:     "330224",
			CityId: "3302",
			Name:   "Purwokerto Selatan",
		},
		{
			Id:     "330225",
			CityId: "3302",
			Name:   "Purwokerto Barat",
		},
		{
			Id:     "330226",
			CityId: "3302",
			Name:   "Purwokerto Timur",
		},
		{
			Id:     "330227",
			CityId: "3302",
			Name:   "Purwokerto Utara",
		},
		{
			Id:     "330301",
			CityId: "3303",
			Name:   "Kemangkon",
		},
		{
			Id:     "330302",
			CityId: "3303",
			Name:   "Bukateja",
		},
		{
			Id:     "330303",
			CityId: "3303",
			Name:   "Kejobong",
		},
		{
			Id:     "330304",
			CityId: "3303",
			Name:   "Kaligondang",
		},
		{
			Id:     "330305",
			CityId: "3303",
			Name:   "Purbalingga",
		},
		{
			Id:     "330306",
			CityId: "3303",
			Name:   "Kalimanah",
		},
		{
			Id:     "330307",
			CityId: "3303",
			Name:   "Kutasari",
		},
		{
			Id:     "330308",
			CityId: "3303",
			Name:   "Mrebet",
		},
		{
			Id:     "330309",
			CityId: "3303",
			Name:   "Bobotsari",
		},
		{
			Id:     "330310",
			CityId: "3303",
			Name:   "Karangreja",
		},
		{
			Id:     "330311",
			CityId: "3303",
			Name:   "Karanganyar",
		},
		{
			Id:     "330312",
			CityId: "3303",
			Name:   "Karangmoncol",
		},
		{
			Id:     "330313",
			CityId: "3303",
			Name:   "Rembang",
		},
		{
			Id:     "330314",
			CityId: "3303",
			Name:   "Bojongsari",
		},
		{
			Id:     "330315",
			CityId: "3303",
			Name:   "Padamara",
		},
		{
			Id:     "330316",
			CityId: "3303",
			Name:   "Pengadegan",
		},
		{
			Id:     "330317",
			CityId: "3303",
			Name:   "Karangjambu",
		},
		{
			Id:     "330318",
			CityId: "3303",
			Name:   "Kertanegara",
		},
		{
			Id:     "330401",
			CityId: "3304",
			Name:   "Susukan",
		},
		{
			Id:     "330402",
			CityId: "3304",
			Name:   "Purworeja Klampok",
		},
		{
			Id:     "330403",
			CityId: "3304",
			Name:   "Mandiraja",
		},
		{
			Id:     "330404",
			CityId: "3304",
			Name:   "Purwanegara",
		},
		{
			Id:     "330405",
			CityId: "3304",
			Name:   "Bawang",
		},
		{
			Id:     "330406",
			CityId: "3304",
			Name:   "Banjarnegara",
		},
		{
			Id:     "330407",
			CityId: "3304",
			Name:   "Sigaluh",
		},
		{
			Id:     "330408",
			CityId: "3304",
			Name:   "Madukara",
		},
		{
			Id:     "330409",
			CityId: "3304",
			Name:   "Banjarmangu",
		},
		{
			Id:     "330410",
			CityId: "3304",
			Name:   "Wanadadi",
		},
		{
			Id:     "330411",
			CityId: "3304",
			Name:   "Rakit",
		},
		{
			Id:     "330412",
			CityId: "3304",
			Name:   "Punggelan",
		},
		{
			Id:     "330413",
			CityId: "3304",
			Name:   "Karangkobar",
		},
		{
			Id:     "330414",
			CityId: "3304",
			Name:   "Pagentan",
		},
		{
			Id:     "330415",
			CityId: "3304",
			Name:   "Pejawaran",
		},
		{
			Id:     "330416",
			CityId: "3304",
			Name:   "Batur",
		},
		{
			Id:     "330417",
			CityId: "3304",
			Name:   "Wanayasa",
		},
		{
			Id:     "330418",
			CityId: "3304",
			Name:   "Kalibening",
		},
		{
			Id:     "330419",
			CityId: "3304",
			Name:   "Pandanarum",
		},
		{
			Id:     "330420",
			CityId: "3304",
			Name:   "Pagedongan",
		},
		{
			Id:     "330501",
			CityId: "3305",
			Name:   "Ayah",
		},
		{
			Id:     "330502",
			CityId: "3305",
			Name:   "Buayan",
		},
		{
			Id:     "330503",
			CityId: "3305",
			Name:   "Puring",
		},
		{
			Id:     "330504",
			CityId: "3305",
			Name:   "Petanahan",
		},
		{
			Id:     "330505",
			CityId: "3305",
			Name:   "Klirong",
		},
		{
			Id:     "330506",
			CityId: "3305",
			Name:   "Buluspesantren",
		},
		{
			Id:     "330507",
			CityId: "3305",
			Name:   "Ambal",
		},
		{
			Id:     "330508",
			CityId: "3305",
			Name:   "Mirit",
		},
		{
			Id:     "330509",
			CityId: "3305",
			Name:   "Prembun",
		},
		{
			Id:     "330510",
			CityId: "3305",
			Name:   "Kutowinangun",
		},
		{
			Id:     "330511",
			CityId: "3305",
			Name:   "Alian",
		},
		{
			Id:     "330512",
			CityId: "3305",
			Name:   "Kebumen",
		},
		{
			Id:     "330513",
			CityId: "3305",
			Name:   "Pejagoan",
		},
		{
			Id:     "330514",
			CityId: "3305",
			Name:   "Sruweng",
		},
		{
			Id:     "330515",
			CityId: "3305",
			Name:   "Adimulyo",
		},
		{
			Id:     "330516",
			CityId: "3305",
			Name:   "Kuwarasan",
		},
		{
			Id:     "330517",
			CityId: "3305",
			Name:   "Rowokele",
		},
		{
			Id:     "330518",
			CityId: "3305",
			Name:   "Sempor",
		},
		{
			Id:     "330519",
			CityId: "3305",
			Name:   "Gombong",
		},
		{
			Id:     "330520",
			CityId: "3305",
			Name:   "Karanganyar",
		},
		{
			Id:     "330521",
			CityId: "3305",
			Name:   "Karanggayam",
		},
		{
			Id:     "330522",
			CityId: "3305",
			Name:   "Sadang",
		},
		{
			Id:     "330523",
			CityId: "3305",
			Name:   "Bonorowo",
		},
		{
			Id:     "330524",
			CityId: "3305",
			Name:   "Padureso",
		},
		{
			Id:     "330525",
			CityId: "3305",
			Name:   "Poncowarno",
		},
		{
			Id:     "330526",
			CityId: "3305",
			Name:   "Karangsambung",
		},
		{
			Id:     "330601",
			CityId: "3306",
			Name:   "Grabag",
		},
		{
			Id:     "330602",
			CityId: "3306",
			Name:   "Ngombol",
		},
		{
			Id:     "330603",
			CityId: "3306",
			Name:   "Purwodadi",
		},
		{
			Id:     "330604",
			CityId: "3306",
			Name:   "Bagelen",
		},
		{
			Id:     "330605",
			CityId: "3306",
			Name:   "Kaligesing",
		},
		{
			Id:     "330606",
			CityId: "3306",
			Name:   "Purworejo",
		},
		{
			Id:     "330607",
			CityId: "3306",
			Name:   "Banyuurip",
		},
		{
			Id:     "330608",
			CityId: "3306",
			Name:   "Bayan",
		},
		{
			Id:     "330609",
			CityId: "3306",
			Name:   "Kutoarjo",
		},
		{
			Id:     "330610",
			CityId: "3306",
			Name:   "Butuh",
		},
		{
			Id:     "330611",
			CityId: "3306",
			Name:   "Pituruh",
		},
		{
			Id:     "330612",
			CityId: "3306",
			Name:   "Kemiri",
		},
		{
			Id:     "330613",
			CityId: "3306",
			Name:   "Bruno",
		},
		{
			Id:     "330614",
			CityId: "3306",
			Name:   "Gebang",
		},
		{
			Id:     "330615",
			CityId: "3306",
			Name:   "Loano",
		},
		{
			Id:     "330616",
			CityId: "3306",
			Name:   "Bener",
		},
		{
			Id:     "330701",
			CityId: "3307",
			Name:   "Wadaslintang",
		},
		{
			Id:     "330702",
			CityId: "3307",
			Name:   "Kepil",
		},
		{
			Id:     "330703",
			CityId: "3307",
			Name:   "Sapuran",
		},
		{
			Id:     "330704",
			CityId: "3307",
			Name:   "Kaliwiro",
		},
		{
			Id:     "330705",
			CityId: "3307",
			Name:   "Leksono",
		},
		{
			Id:     "330706",
			CityId: "3307",
			Name:   "Selomerto",
		},
		{
			Id:     "330707",
			CityId: "3307",
			Name:   "Kalikajar",
		},
		{
			Id:     "330708",
			CityId: "3307",
			Name:   "Kertek",
		},
		{
			Id:     "330709",
			CityId: "3307",
			Name:   "Wonosobo",
		},
		{
			Id:     "330710",
			CityId: "3307",
			Name:   "Watumalang",
		},
		{
			Id:     "330711",
			CityId: "3307",
			Name:   "Mojotengah",
		},
		{
			Id:     "330712",
			CityId: "3307",
			Name:   "Garung",
		},
		{
			Id:     "330713",
			CityId: "3307",
			Name:   "Kejajar",
		},
		{
			Id:     "330714",
			CityId: "3307",
			Name:   "Sukoharjo",
		},
		{
			Id:     "330715",
			CityId: "3307",
			Name:   "Kalibawang",
		},
		{
			Id:     "330801",
			CityId: "3308",
			Name:   "Salaman",
		},
		{
			Id:     "330802",
			CityId: "3308",
			Name:   "Borobudur",
		},
		{
			Id:     "330803",
			CityId: "3308",
			Name:   "Ngluwar",
		},
		{
			Id:     "330804",
			CityId: "3308",
			Name:   "Salam",
		},
		{
			Id:     "330805",
			CityId: "3308",
			Name:   "Srumbung",
		},
		{
			Id:     "330806",
			CityId: "3308",
			Name:   "Dukun",
		},
		{
			Id:     "330807",
			CityId: "3308",
			Name:   "Sawangan",
		},
		{
			Id:     "330808",
			CityId: "3308",
			Name:   "Muntilan",
		},
		{
			Id:     "330809",
			CityId: "3308",
			Name:   "Mungkid",
		},
		{
			Id:     "330810",
			CityId: "3308",
			Name:   "Mertoyudan",
		},
		{
			Id:     "330811",
			CityId: "3308",
			Name:   "Tempuran",
		},
		{
			Id:     "330812",
			CityId: "3308",
			Name:   "Kajoran",
		},
		{
			Id:     "330813",
			CityId: "3308",
			Name:   "Kaliangkrik",
		},
		{
			Id:     "330814",
			CityId: "3308",
			Name:   "Bandongan",
		},
		{
			Id:     "330815",
			CityId: "3308",
			Name:   "Candimulyo",
		},
		{
			Id:     "330816",
			CityId: "3308",
			Name:   "Pakis",
		},
		{
			Id:     "330817",
			CityId: "3308",
			Name:   "Ngablak",
		},
		{
			Id:     "330818",
			CityId: "3308",
			Name:   "Grabag",
		},
		{
			Id:     "330819",
			CityId: "3308",
			Name:   "Tegalrejo",
		},
		{
			Id:     "330820",
			CityId: "3308",
			Name:   "Secang",
		},
		{
			Id:     "330821",
			CityId: "3308",
			Name:   "Windusari",
		},
		{
			Id:     "330901",
			CityId: "3309",
			Name:   "Selo",
		},
		{
			Id:     "330902",
			CityId: "3309",
			Name:   "Ampel",
		},
		{
			Id:     "330903",
			CityId: "3309",
			Name:   "Cepogo",
		},
		{
			Id:     "330904",
			CityId: "3309",
			Name:   "Musuk",
		},
		{
			Id:     "330905",
			CityId: "3309",
			Name:   "Boyolali",
		},
		{
			Id:     "330906",
			CityId: "3309",
			Name:   "Mojosongo",
		},
		{
			Id:     "330907",
			CityId: "3309",
			Name:   "Teras",
		},
		{
			Id:     "330908",
			CityId: "3309",
			Name:   "Sawit",
		},
		{
			Id:     "330909",
			CityId: "3309",
			Name:   "Banyudono",
		},
		{
			Id:     "330910",
			CityId: "3309",
			Name:   "Sambi",
		},
		{
			Id:     "330911",
			CityId: "3309",
			Name:   "Ngemplak",
		},
		{
			Id:     "330912",
			CityId: "3309",
			Name:   "Nogosari",
		},
		{
			Id:     "330913",
			CityId: "3309",
			Name:   "Simo",
		},
		{
			Id:     "330914",
			CityId: "3309",
			Name:   "Karanggede",
		},
		{
			Id:     "330915",
			CityId: "3309",
			Name:   "Klego",
		},
		{
			Id:     "330916",
			CityId: "3309",
			Name:   "Andong",
		},
		{
			Id:     "330917",
			CityId: "3309",
			Name:   "Kemusu",
		},
		{
			Id:     "330918",
			CityId: "3309",
			Name:   "Wonosegoro",
		},
		{
			Id:     "330919",
			CityId: "3309",
			Name:   "Juwangi",
		},
		{
			Id:     "330920",
			CityId: "3309",
			Name:   "Gladagsari",
		},
		{
			Id:     "330921",
			CityId: "3309",
			Name:   "Tamansari",
		},
		{
			Id:     "330922",
			CityId: "3309",
			Name:   "Wonosamodro",
		},
		{
			Id:     "331001",
			CityId: "3310",
			Name:   "Prambanan",
		},
		{
			Id:     "331002",
			CityId: "3310",
			Name:   "Gantiwarno",
		},
		{
			Id:     "331003",
			CityId: "3310",
			Name:   "Wedi",
		},
		{
			Id:     "331004",
			CityId: "3310",
			Name:   "Bayat",
		},
		{
			Id:     "331005",
			CityId: "3310",
			Name:   "Cawas",
		},
		{
			Id:     "331006",
			CityId: "3310",
			Name:   "Trucuk",
		},
		{
			Id:     "331007",
			CityId: "3310",
			Name:   "Kebonarum",
		},
		{
			Id:     "331008",
			CityId: "3310",
			Name:   "Jogonalan",
		},
		{
			Id:     "331009",
			CityId: "3310",
			Name:   "Manisrenggo",
		},
		{
			Id:     "331010",
			CityId: "3310",
			Name:   "Karangnongko",
		},
		{
			Id:     "331011",
			CityId: "3310",
			Name:   "Ceper",
		},
		{
			Id:     "331012",
			CityId: "3310",
			Name:   "Pedan",
		},
		{
			Id:     "331013",
			CityId: "3310",
			Name:   "Karangdowo",
		},
		{
			Id:     "331014",
			CityId: "3310",
			Name:   "Juwiring",
		},
		{
			Id:     "331015",
			CityId: "3310",
			Name:   "Wonosari",
		},
		{
			Id:     "331016",
			CityId: "3310",
			Name:   "Delanggu",
		},
		{
			Id:     "331017",
			CityId: "3310",
			Name:   "Polanharjo",
		},
		{
			Id:     "331018",
			CityId: "3310",
			Name:   "Karanganom",
		},
		{
			Id:     "331019",
			CityId: "3310",
			Name:   "Tulung",
		},
		{
			Id:     "331020",
			CityId: "3310",
			Name:   "Jatinom",
		},
		{
			Id:     "331021",
			CityId: "3310",
			Name:   "Kemalang",
		},
		{
			Id:     "331022",
			CityId: "3310",
			Name:   "Ngawen",
		},
		{
			Id:     "331023",
			CityId: "3310",
			Name:   "Kalikotes",
		},
		{
			Id:     "331024",
			CityId: "3310",
			Name:   "Klaten Utara",
		},
		{
			Id:     "331025",
			CityId: "3310",
			Name:   "Klaten Tengah",
		},
		{
			Id:     "331026",
			CityId: "3310",
			Name:   "Klaten Selatan",
		},
		{
			Id:     "331101",
			CityId: "3311",
			Name:   "Weru",
		},
		{
			Id:     "331102",
			CityId: "3311",
			Name:   "Bulu",
		},
		{
			Id:     "331103",
			CityId: "3311",
			Name:   "Tawangsari",
		},
		{
			Id:     "331104",
			CityId: "3311",
			Name:   "Sukoharjo",
		},
		{
			Id:     "331105",
			CityId: "3311",
			Name:   "Nguter",
		},
		{
			Id:     "331106",
			CityId: "3311",
			Name:   "Bendosari",
		},
		{
			Id:     "331107",
			CityId: "3311",
			Name:   "Polokarto",
		},
		{
			Id:     "331108",
			CityId: "3311",
			Name:   "Mojolaban",
		},
		{
			Id:     "331109",
			CityId: "3311",
			Name:   "Grogol",
		},
		{
			Id:     "331110",
			CityId: "3311",
			Name:   "Baki",
		},
		{
			Id:     "331111",
			CityId: "3311",
			Name:   "Gatak",
		},
		{
			Id:     "331112",
			CityId: "3311",
			Name:   "Kartasura",
		},
		{
			Id:     "331201",
			CityId: "3312",
			Name:   "Pracimantoro",
		},
		{
			Id:     "331202",
			CityId: "3312",
			Name:   "Giritontro",
		},
		{
			Id:     "331203",
			CityId: "3312",
			Name:   "Giriwoyo",
		},
		{
			Id:     "331204",
			CityId: "3312",
			Name:   "Batuwarno",
		},
		{
			Id:     "331205",
			CityId: "3312",
			Name:   "Tirtomoyo",
		},
		{
			Id:     "331206",
			CityId: "3312",
			Name:   "Nguntoronadi",
		},
		{
			Id:     "331207",
			CityId: "3312",
			Name:   "Baturetno",
		},
		{
			Id:     "331208",
			CityId: "3312",
			Name:   "Eromoko",
		},
		{
			Id:     "331209",
			CityId: "3312",
			Name:   "Wuryantoro",
		},
		{
			Id:     "331210",
			CityId: "3312",
			Name:   "Manyaran",
		},
		{
			Id:     "331211",
			CityId: "3312",
			Name:   "Selogiri",
		},
		{
			Id:     "331212",
			CityId: "3312",
			Name:   "Wonogiri",
		},
		{
			Id:     "331213",
			CityId: "3312",
			Name:   "Ngadirojo",
		},
		{
			Id:     "331214",
			CityId: "3312",
			Name:   "Sidoharjo",
		},
		{
			Id:     "331215",
			CityId: "3312",
			Name:   "Jatiroto",
		},
		{
			Id:     "331216",
			CityId: "3312",
			Name:   "Kismantoro",
		},
		{
			Id:     "331217",
			CityId: "3312",
			Name:   "Purwantoro",
		},
		{
			Id:     "331218",
			CityId: "3312",
			Name:   "Bulukerto",
		},
		{
			Id:     "331219",
			CityId: "3312",
			Name:   "Slogohimo",
		},
		{
			Id:     "331220",
			CityId: "3312",
			Name:   "Jatisrono",
		},
		{
			Id:     "331221",
			CityId: "3312",
			Name:   "Jatipurno",
		},
		{
			Id:     "331222",
			CityId: "3312",
			Name:   "Girimarto",
		},
		{
			Id:     "331223",
			CityId: "3312",
			Name:   "Karangtengah",
		},
		{
			Id:     "331224",
			CityId: "3312",
			Name:   "Paranggupito",
		},
		{
			Id:     "331225",
			CityId: "3312",
			Name:   "Puhpelem",
		},
		{
			Id:     "331301",
			CityId: "3313",
			Name:   "Jatipuro",
		},
		{
			Id:     "331302",
			CityId: "3313",
			Name:   "Jatiyoso",
		},
		{
			Id:     "331303",
			CityId: "3313",
			Name:   "Jumapolo",
		},
		{
			Id:     "331304",
			CityId: "3313",
			Name:   "Jumantono",
		},
		{
			Id:     "331305",
			CityId: "3313",
			Name:   "Matesih",
		},
		{
			Id:     "331306",
			CityId: "3313",
			Name:   "Tawangmangu",
		},
		{
			Id:     "331307",
			CityId: "3313",
			Name:   "Ngargoyoso",
		},
		{
			Id:     "331308",
			CityId: "3313",
			Name:   "Karangpandan",
		},
		{
			Id:     "331309",
			CityId: "3313",
			Name:   "Karanganyar",
		},
		{
			Id:     "331310",
			CityId: "3313",
			Name:   "Tasikmadu",
		},
		{
			Id:     "331311",
			CityId: "3313",
			Name:   "Jaten",
		},
		{
			Id:     "331312",
			CityId: "3313",
			Name:   "Colomadu",
		},
		{
			Id:     "331313",
			CityId: "3313",
			Name:   "Gondangrejo",
		},
		{
			Id:     "331314",
			CityId: "3313",
			Name:   "Kebakkramat",
		},
		{
			Id:     "331315",
			CityId: "3313",
			Name:   "Mojogedang",
		},
		{
			Id:     "331316",
			CityId: "3313",
			Name:   "Kerjo",
		},
		{
			Id:     "331317",
			CityId: "3313",
			Name:   "Jenawi",
		},
		{
			Id:     "331401",
			CityId: "3314",
			Name:   "Kalijambe",
		},
		{
			Id:     "331402",
			CityId: "3314",
			Name:   "Plupuh",
		},
		{
			Id:     "331403",
			CityId: "3314",
			Name:   "Masaran",
		},
		{
			Id:     "331404",
			CityId: "3314",
			Name:   "Kedawung",
		},
		{
			Id:     "331405",
			CityId: "3314",
			Name:   "Sambirejo",
		},
		{
			Id:     "331406",
			CityId: "3314",
			Name:   "Gondang",
		},
		{
			Id:     "331407",
			CityId: "3314",
			Name:   "Sambungmacan",
		},
		{
			Id:     "331408",
			CityId: "3314",
			Name:   "Ngrampal",
		},
		{
			Id:     "331409",
			CityId: "3314",
			Name:   "Karangmalang",
		},
		{
			Id:     "331410",
			CityId: "3314",
			Name:   "Sragen",
		},
		{
			Id:     "331411",
			CityId: "3314",
			Name:   "Sidoharjo",
		},
		{
			Id:     "331412",
			CityId: "3314",
			Name:   "Tanon",
		},
		{
			Id:     "331413",
			CityId: "3314",
			Name:   "Gemolong",
		},
		{
			Id:     "331414",
			CityId: "3314",
			Name:   "Miri",
		},
		{
			Id:     "331415",
			CityId: "3314",
			Name:   "Sumberlawang",
		},
		{
			Id:     "331416",
			CityId: "3314",
			Name:   "Mondokan",
		},
		{
			Id:     "331417",
			CityId: "3314",
			Name:   "Sukodono",
		},
		{
			Id:     "331418",
			CityId: "3314",
			Name:   "Gesi",
		},
		{
			Id:     "331419",
			CityId: "3314",
			Name:   "Tangen",
		},
		{
			Id:     "331420",
			CityId: "3314",
			Name:   "Jenar",
		},
		{
			Id:     "331501",
			CityId: "3315",
			Name:   "Kedungjati",
		},
		{
			Id:     "331502",
			CityId: "3315",
			Name:   "Karangrayung",
		},
		{
			Id:     "331503",
			CityId: "3315",
			Name:   "Penawangan",
		},
		{
			Id:     "331504",
			CityId: "3315",
			Name:   "Toroh",
		},
		{
			Id:     "331505",
			CityId: "3315",
			Name:   "Geyer",
		},
		{
			Id:     "331506",
			CityId: "3315",
			Name:   "Pulokulon",
		},
		{
			Id:     "331507",
			CityId: "3315",
			Name:   "Kradenan",
		},
		{
			Id:     "331508",
			CityId: "3315",
			Name:   "Gabus",
		},
		{
			Id:     "331509",
			CityId: "3315",
			Name:   "Ngaringan",
		},
		{
			Id:     "331510",
			CityId: "3315",
			Name:   "Wirosari",
		},
		{
			Id:     "331511",
			CityId: "3315",
			Name:   "Tawangharjo",
		},
		{
			Id:     "331512",
			CityId: "3315",
			Name:   "Grobogan",
		},
		{
			Id:     "331513",
			CityId: "3315",
			Name:   "Purwodadi",
		},
		{
			Id:     "331514",
			CityId: "3315",
			Name:   "Brati",
		},
		{
			Id:     "331515",
			CityId: "3315",
			Name:   "Klambu",
		},
		{
			Id:     "331516",
			CityId: "3315",
			Name:   "Godong",
		},
		{
			Id:     "331517",
			CityId: "3315",
			Name:   "Gubug",
		},
		{
			Id:     "331518",
			CityId: "3315",
			Name:   "Tegowanu",
		},
		{
			Id:     "331519",
			CityId: "3315",
			Name:   "Tanggungharjo",
		},
		{
			Id:     "331601",
			CityId: "3316",
			Name:   "Jati",
		},
		{
			Id:     "331602",
			CityId: "3316",
			Name:   "Randublatung",
		},
		{
			Id:     "331603",
			CityId: "3316",
			Name:   "Kradenan",
		},
		{
			Id:     "331604",
			CityId: "3316",
			Name:   "Kedungtuban",
		},
		{
			Id:     "331605",
			CityId: "3316",
			Name:   "Cepu",
		},
		{
			Id:     "331606",
			CityId: "3316",
			Name:   "Sambong",
		},
		{
			Id:     "331607",
			CityId: "3316",
			Name:   "Jiken",
		},
		{
			Id:     "331608",
			CityId: "3316",
			Name:   "Jepon",
		},
		{
			Id:     "331609",
			CityId: "3316",
			Name:   "Blora",
		},
		{
			Id:     "331610",
			CityId: "3316",
			Name:   "Tunjungan",
		},
		{
			Id:     "331611",
			CityId: "3316",
			Name:   "Banjarejo",
		},
		{
			Id:     "331612",
			CityId: "3316",
			Name:   "Ngawen",
		},
		{
			Id:     "331613",
			CityId: "3316",
			Name:   "Kunduran",
		},
		{
			Id:     "331614",
			CityId: "3316",
			Name:   "Todanan",
		},
		{
			Id:     "331615",
			CityId: "3316",
			Name:   "Bogorejo",
		},
		{
			Id:     "331616",
			CityId: "3316",
			Name:   "Japah",
		},
		{
			Id:     "331701",
			CityId: "3317",
			Name:   "Sumber",
		},
		{
			Id:     "331702",
			CityId: "3317",
			Name:   "Bulu",
		},
		{
			Id:     "331703",
			CityId: "3317",
			Name:   "Gunem",
		},
		{
			Id:     "331704",
			CityId: "3317",
			Name:   "Sale",
		},
		{
			Id:     "331705",
			CityId: "3317",
			Name:   "Sarang",
		},
		{
			Id:     "331706",
			CityId: "3317",
			Name:   "Sedan",
		},
		{
			Id:     "331707",
			CityId: "3317",
			Name:   "Pamotan",
		},
		{
			Id:     "331708",
			CityId: "3317",
			Name:   "Sulang",
		},
		{
			Id:     "331709",
			CityId: "3317",
			Name:   "Kaliori",
		},
		{
			Id:     "331710",
			CityId: "3317",
			Name:   "Rembang",
		},
		{
			Id:     "331711",
			CityId: "3317",
			Name:   "Pancur",
		},
		{
			Id:     "331712",
			CityId: "3317",
			Name:   "Kragan",
		},
		{
			Id:     "331713",
			CityId: "3317",
			Name:   "Sluke",
		},
		{
			Id:     "331714",
			CityId: "3317",
			Name:   "Lasem",
		},
		{
			Id:     "331801",
			CityId: "3318",
			Name:   "Sukolilo",
		},
		{
			Id:     "331802",
			CityId: "3318",
			Name:   "Kayen",
		},
		{
			Id:     "331803",
			CityId: "3318",
			Name:   "Tambakromo",
		},
		{
			Id:     "331804",
			CityId: "3318",
			Name:   "Winong",
		},
		{
			Id:     "331805",
			CityId: "3318",
			Name:   "Pucakwangi",
		},
		{
			Id:     "331806",
			CityId: "3318",
			Name:   "Jaken",
		},
		{
			Id:     "331807",
			CityId: "3318",
			Name:   "Batangan",
		},
		{
			Id:     "331808",
			CityId: "3318",
			Name:   "Juwana",
		},
		{
			Id:     "331809",
			CityId: "3318",
			Name:   "Jakenan",
		},
		{
			Id:     "331810",
			CityId: "3318",
			Name:   "Pati",
		},
		{
			Id:     "331811",
			CityId: "3318",
			Name:   "Gabus",
		},
		{
			Id:     "331812",
			CityId: "3318",
			Name:   "Margorejo",
		},
		{
			Id:     "331813",
			CityId: "3318",
			Name:   "Gembong",
		},
		{
			Id:     "331814",
			CityId: "3318",
			Name:   "Tlogowungu",
		},
		{
			Id:     "331815",
			CityId: "3318",
			Name:   "Wedarijaksa",
		},
		{
			Id:     "331816",
			CityId: "3318",
			Name:   "Margoyoso",
		},
		{
			Id:     "331817",
			CityId: "3318",
			Name:   "Gunungwungkal",
		},
		{
			Id:     "331818",
			CityId: "3318",
			Name:   "Cluwak",
		},
		{
			Id:     "331819",
			CityId: "3318",
			Name:   "Tayu",
		},
		{
			Id:     "331820",
			CityId: "3318",
			Name:   "Dukuhseti",
		},
		{
			Id:     "331821",
			CityId: "3318",
			Name:   "Trangkil",
		},
		{
			Id:     "331901",
			CityId: "3319",
			Name:   "Kaliwungu",
		},
		{
			Id:     "331902",
			CityId: "3319",
			Name:   "Kota Kudus",
		},
		{
			Id:     "331903",
			CityId: "3319",
			Name:   "Jati",
		},
		{
			Id:     "331904",
			CityId: "3319",
			Name:   "Undaan",
		},
		{
			Id:     "331905",
			CityId: "3319",
			Name:   "Mejobo",
		},
		{
			Id:     "331906",
			CityId: "3319",
			Name:   "Jekulo",
		},
		{
			Id:     "331907",
			CityId: "3319",
			Name:   "Bae",
		},
		{
			Id:     "331908",
			CityId: "3319",
			Name:   "Gebog",
		},
		{
			Id:     "331909",
			CityId: "3319",
			Name:   "Dawe",
		},
		{
			Id:     "332001",
			CityId: "3320",
			Name:   "Kedung",
		},
		{
			Id:     "332002",
			CityId: "3320",
			Name:   "Pecangaan",
		},
		{
			Id:     "332003",
			CityId: "3320",
			Name:   "Welahan",
		},
		{
			Id:     "332004",
			CityId: "3320",
			Name:   "Mayong",
		},
		{
			Id:     "332005",
			CityId: "3320",
			Name:   "Batealit",
		},
		{
			Id:     "332006",
			CityId: "3320",
			Name:   "Jepara",
		},
		{
			Id:     "332007",
			CityId: "3320",
			Name:   "Mlonggo",
		},
		{
			Id:     "332008",
			CityId: "3320",
			Name:   "Bangsri",
		},
		{
			Id:     "332009",
			CityId: "3320",
			Name:   "Keling",
		},
		{
			Id:     "332010",
			CityId: "3320",
			Name:   "Karimunjawa",
		},
		{
			Id:     "332011",
			CityId: "3320",
			Name:   "Tahunan",
		},
		{
			Id:     "332012",
			CityId: "3320",
			Name:   "Nalumsari",
		},
		{
			Id:     "332013",
			CityId: "3320",
			Name:   "Kalinyamatan",
		},
		{
			Id:     "332014",
			CityId: "3320",
			Name:   "Kembang",
		},
		{
			Id:     "332015",
			CityId: "3320",
			Name:   "Pakis Aji",
		},
		{
			Id:     "332016",
			CityId: "3320",
			Name:   "Donorojo",
		},
		{
			Id:     "332101",
			CityId: "3321",
			Name:   "Mranggen",
		},
		{
			Id:     "332102",
			CityId: "3321",
			Name:   "Karangawen",
		},
		{
			Id:     "332103",
			CityId: "3321",
			Name:   "Guntur",
		},
		{
			Id:     "332104",
			CityId: "3321",
			Name:   "Sayung",
		},
		{
			Id:     "332105",
			CityId: "3321",
			Name:   "Karangtengah",
		},
		{
			Id:     "332106",
			CityId: "3321",
			Name:   "Wonosalam",
		},
		{
			Id:     "332107",
			CityId: "3321",
			Name:   "Dempet",
		},
		{
			Id:     "332108",
			CityId: "3321",
			Name:   "Gajah",
		},
		{
			Id:     "332109",
			CityId: "3321",
			Name:   "Karanganyar",
		},
		{
			Id:     "332110",
			CityId: "3321",
			Name:   "Mijen",
		},
		{
			Id:     "332111",
			CityId: "3321",
			Name:   "Demak",
		},
		{
			Id:     "332112",
			CityId: "3321",
			Name:   "Bonang",
		},
		{
			Id:     "332113",
			CityId: "3321",
			Name:   "Wedung",
		},
		{
			Id:     "332114",
			CityId: "3321",
			Name:   "Kebonagung",
		},
		{
			Id:     "332201",
			CityId: "3322",
			Name:   "Getasan",
		},
		{
			Id:     "332202",
			CityId: "3322",
			Name:   "Tengaran",
		},
		{
			Id:     "332203",
			CityId: "3322",
			Name:   "Susukan",
		},
		{
			Id:     "332204",
			CityId: "3322",
			Name:   "Suruh",
		},
		{
			Id:     "332205",
			CityId: "3322",
			Name:   "Pabelan",
		},
		{
			Id:     "332206",
			CityId: "3322",
			Name:   "Tuntang",
		},
		{
			Id:     "332207",
			CityId: "3322",
			Name:   "Banyubiru",
		},
		{
			Id:     "332208",
			CityId: "3322",
			Name:   "Jambu",
		},
		{
			Id:     "332209",
			CityId: "3322",
			Name:   "Sumowono",
		},
		{
			Id:     "332210",
			CityId: "3322",
			Name:   "Ambarawa",
		},
		{
			Id:     "332211",
			CityId: "3322",
			Name:   "Bawen",
		},
		{
			Id:     "332212",
			CityId: "3322",
			Name:   "Bringin",
		},
		{
			Id:     "332213",
			CityId: "3322",
			Name:   "Bergas",
		},
		{
			Id:     "332215",
			CityId: "3322",
			Name:   "Pringapus",
		},
		{
			Id:     "332216",
			CityId: "3322",
			Name:   "Bancak",
		},
		{
			Id:     "332217",
			CityId: "3322",
			Name:   "Kaliwungu",
		},
		{
			Id:     "332218",
			CityId: "3322",
			Name:   "Ungaran Barat",
		},
		{
			Id:     "332219",
			CityId: "3322",
			Name:   "Ungaran Timur",
		},
		{
			Id:     "332220",
			CityId: "3322",
			Name:   "Bandungan",
		},
		{
			Id:     "332301",
			CityId: "3323",
			Name:   "Bulu",
		},
		{
			Id:     "332302",
			CityId: "3323",
			Name:   "Tembarak",
		},
		{
			Id:     "332303",
			CityId: "3323",
			Name:   "Temanggung",
		},
		{
			Id:     "332304",
			CityId: "3323",
			Name:   "Pringsurat",
		},
		{
			Id:     "332305",
			CityId: "3323",
			Name:   "Kaloran",
		},
		{
			Id:     "332306",
			CityId: "3323",
			Name:   "Kandangan",
		},
		{
			Id:     "332307",
			CityId: "3323",
			Name:   "Kedu",
		},
		{
			Id:     "332308",
			CityId: "3323",
			Name:   "Parakan",
		},
		{
			Id:     "332309",
			CityId: "3323",
			Name:   "Ngadirejo",
		},
		{
			Id:     "332310",
			CityId: "3323",
			Name:   "Jumo",
		},
		{
			Id:     "332311",
			CityId: "3323",
			Name:   "Tretep",
		},
		{
			Id:     "332312",
			CityId: "3323",
			Name:   "Candiroto",
		},
		{
			Id:     "332313",
			CityId: "3323",
			Name:   "Kranggan",
		},
		{
			Id:     "332314",
			CityId: "3323",
			Name:   "Tlogomulyo",
		},
		{
			Id:     "332315",
			CityId: "3323",
			Name:   "Selopampang",
		},
		{
			Id:     "332316",
			CityId: "3323",
			Name:   "Bansari",
		},
		{
			Id:     "332317",
			CityId: "3323",
			Name:   "Kledung",
		},
		{
			Id:     "332318",
			CityId: "3323",
			Name:   "Bejen",
		},
		{
			Id:     "332319",
			CityId: "3323",
			Name:   "Wonoboyo",
		},
		{
			Id:     "332320",
			CityId: "3323",
			Name:   "Gemawang",
		},
		{
			Id:     "332401",
			CityId: "3324",
			Name:   "Plantungan",
		},
		{
			Id:     "332402",
			CityId: "3324",
			Name:   "Pageruyung",
		},
		{
			Id:     "332403",
			CityId: "3324",
			Name:   "Sukorejo",
		},
		{
			Id:     "332404",
			CityId: "3324",
			Name:   "Patean",
		},
		{
			Id:     "332405",
			CityId: "3324",
			Name:   "Singorojo",
		},
		{
			Id:     "332406",
			CityId: "3324",
			Name:   "Limbangan",
		},
		{
			Id:     "332407",
			CityId: "3324",
			Name:   "Boja",
		},
		{
			Id:     "332408",
			CityId: "3324",
			Name:   "Kaliwungu",
		},
		{
			Id:     "332409",
			CityId: "3324",
			Name:   "Brangsong",
		},
		{
			Id:     "332410",
			CityId: "3324",
			Name:   "Pegandon",
		},
		{
			Id:     "332411",
			CityId: "3324",
			Name:   "Gemuh",
		},
		{
			Id:     "332412",
			CityId: "3324",
			Name:   "Weleri",
		},
		{
			Id:     "332413",
			CityId: "3324",
			Name:   "Cepiring",
		},
		{
			Id:     "332414",
			CityId: "3324",
			Name:   "Patebon",
		},
		{
			Id:     "332415",
			CityId: "3324",
			Name:   "Kendal",
		},
		{
			Id:     "332416",
			CityId: "3324",
			Name:   "Rowosari",
		},
		{
			Id:     "332417",
			CityId: "3324",
			Name:   "Kangkung",
		},
		{
			Id:     "332418",
			CityId: "3324",
			Name:   "Ringinarum",
		},
		{
			Id:     "332419",
			CityId: "3324",
			Name:   "Ngampel",
		},
		{
			Id:     "332420",
			CityId: "3324",
			Name:   "Kaliwungu Selatan",
		},
		{
			Id:     "332501",
			CityId: "3325",
			Name:   "Wonotunggal",
		},
		{
			Id:     "332502",
			CityId: "3325",
			Name:   "Bandar",
		},
		{
			Id:     "332503",
			CityId: "3325",
			Name:   "Blado",
		},
		{
			Id:     "332504",
			CityId: "3325",
			Name:   "Reban",
		},
		{
			Id:     "332505",
			CityId: "3325",
			Name:   "Bawang",
		},
		{
			Id:     "332506",
			CityId: "3325",
			Name:   "Tersono",
		},
		{
			Id:     "332507",
			CityId: "3325",
			Name:   "Gringsing",
		},
		{
			Id:     "332508",
			CityId: "3325",
			Name:   "Limpung",
		},
		{
			Id:     "332509",
			CityId: "3325",
			Name:   "Subah",
		},
		{
			Id:     "332510",
			CityId: "3325",
			Name:   "Tulis",
		},
		{
			Id:     "332511",
			CityId: "3325",
			Name:   "Batang",
		},
		{
			Id:     "332512",
			CityId: "3325",
			Name:   "Warungasem",
		},
		{
			Id:     "332513",
			CityId: "3325",
			Name:   "Kandeman",
		},
		{
			Id:     "332514",
			CityId: "3325",
			Name:   "Pecalungan",
		},
		{
			Id:     "332515",
			CityId: "3325",
			Name:   "Banyuputih",
		},
		{
			Id:     "332601",
			CityId: "3326",
			Name:   "Kandangserang",
		},
		{
			Id:     "332602",
			CityId: "3326",
			Name:   "Paninggaran",
		},
		{
			Id:     "332603",
			CityId: "3326",
			Name:   "Lebakbarang",
		},
		{
			Id:     "332604",
			CityId: "3326",
			Name:   "Petungkriyono",
		},
		{
			Id:     "332605",
			CityId: "3326",
			Name:   "Talun",
		},
		{
			Id:     "332606",
			CityId: "3326",
			Name:   "Doro",
		},
		{
			Id:     "332607",
			CityId: "3326",
			Name:   "Karanganyar",
		},
		{
			Id:     "332608",
			CityId: "3326",
			Name:   "Kajen",
		},
		{
			Id:     "332609",
			CityId: "3326",
			Name:   "Kesesi",
		},
		{
			Id:     "332610",
			CityId: "3326",
			Name:   "Sragi",
		},
		{
			Id:     "332611",
			CityId: "3326",
			Name:   "Bojong",
		},
		{
			Id:     "332612",
			CityId: "3326",
			Name:   "Wonopringgo",
		},
		{
			Id:     "332613",
			CityId: "3326",
			Name:   "Kedungwuni",
		},
		{
			Id:     "332614",
			CityId: "3326",
			Name:   "Buaran",
		},
		{
			Id:     "332615",
			CityId: "3326",
			Name:   "Tirto",
		},
		{
			Id:     "332616",
			CityId: "3326",
			Name:   "Wiradesa",
		},
		{
			Id:     "332617",
			CityId: "3326",
			Name:   "Siwalan",
		},
		{
			Id:     "332618",
			CityId: "3326",
			Name:   "Karangdadap",
		},
		{
			Id:     "332619",
			CityId: "3326",
			Name:   "Wonokerto",
		},
		{
			Id:     "332701",
			CityId: "3327",
			Name:   "Moga",
		},
		{
			Id:     "332702",
			CityId: "3327",
			Name:   "Pulosari",
		},
		{
			Id:     "332703",
			CityId: "3327",
			Name:   "Belik",
		},
		{
			Id:     "332704",
			CityId: "3327",
			Name:   "Watukumpul",
		},
		{
			Id:     "332705",
			CityId: "3327",
			Name:   "Bodeh",
		},
		{
			Id:     "332706",
			CityId: "3327",
			Name:   "Bantarbolang",
		},
		{
			Id:     "332707",
			CityId: "3327",
			Name:   "Randudongkal",
		},
		{
			Id:     "332708",
			CityId: "3327",
			Name:   "Pemalang",
		},
		{
			Id:     "332709",
			CityId: "3327",
			Name:   "Taman",
		},
		{
			Id:     "332710",
			CityId: "3327",
			Name:   "Petarukan",
		},
		{
			Id:     "332711",
			CityId: "3327",
			Name:   "Ampelgading",
		},
		{
			Id:     "332712",
			CityId: "3327",
			Name:   "Comal",
		},
		{
			Id:     "332713",
			CityId: "3327",
			Name:   "Ulujami",
		},
		{
			Id:     "332714",
			CityId: "3327",
			Name:   "Warungpring",
		},
		{
			Id:     "332801",
			CityId: "3328",
			Name:   "Margasari",
		},
		{
			Id:     "332802",
			CityId: "3328",
			Name:   "Bumijawa",
		},
		{
			Id:     "332803",
			CityId: "3328",
			Name:   "Bojong",
		},
		{
			Id:     "332804",
			CityId: "3328",
			Name:   "Balapulang",
		},
		{
			Id:     "332805",
			CityId: "3328",
			Name:   "Pagerbarang",
		},
		{
			Id:     "332806",
			CityId: "3328",
			Name:   "Lebaksiu",
		},
		{
			Id:     "332807",
			CityId: "3328",
			Name:   "Jatinegara",
		},
		{
			Id:     "332808",
			CityId: "3328",
			Name:   "Kedungbanteng",
		},
		{
			Id:     "332809",
			CityId: "3328",
			Name:   "Pangkah",
		},
		{
			Id:     "332810",
			CityId: "3328",
			Name:   "Slawi",
		},
		{
			Id:     "332811",
			CityId: "3328",
			Name:   "Adiwerna",
		},
		{
			Id:     "332812",
			CityId: "3328",
			Name:   "Talang",
		},
		{
			Id:     "332813",
			CityId: "3328",
			Name:   "Dukuhturi",
		},
		{
			Id:     "332814",
			CityId: "3328",
			Name:   "Tarub",
		},
		{
			Id:     "332815",
			CityId: "3328",
			Name:   "Kramat",
		},
		{
			Id:     "332816",
			CityId: "3328",
			Name:   "Suradadi",
		},
		{
			Id:     "332817",
			CityId: "3328",
			Name:   "Warureja",
		},
		{
			Id:     "332818",
			CityId: "3328",
			Name:   "Dukuhwaru",
		},
		{
			Id:     "332901",
			CityId: "3329",
			Name:   "Salem",
		},
		{
			Id:     "332902",
			CityId: "3329",
			Name:   "Bantarkawung",
		},
		{
			Id:     "332903",
			CityId: "3329",
			Name:   "Bumiayu",
		},
		{
			Id:     "332904",
			CityId: "3329",
			Name:   "Paguyangan",
		},
		{
			Id:     "332905",
			CityId: "3329",
			Name:   "Sirampog",
		},
		{
			Id:     "332906",
			CityId: "3329",
			Name:   "Tonjong",
		},
		{
			Id:     "332907",
			CityId: "3329",
			Name:   "Jatibarang",
		},
		{
			Id:     "332908",
			CityId: "3329",
			Name:   "Wanasari",
		},
		{
			Id:     "332909",
			CityId: "3329",
			Name:   "Brebes",
		},
		{
			Id:     "332910",
			CityId: "3329",
			Name:   "Songgom",
		},
		{
			Id:     "332911",
			CityId: "3329",
			Name:   "Kersana",
		},
		{
			Id:     "332912",
			CityId: "3329",
			Name:   "Losari",
		},
		{
			Id:     "332913",
			CityId: "3329",
			Name:   "Tanjung",
		},
		{
			Id:     "332914",
			CityId: "3329",
			Name:   "Bulakamba",
		},
		{
			Id:     "332915",
			CityId: "3329",
			Name:   "Larangan",
		},
		{
			Id:     "332916",
			CityId: "3329",
			Name:   "Ketanggungan",
		},
		{
			Id:     "332917",
			CityId: "3329",
			Name:   "Banjarharjo",
		},
		{
			Id:     "337101",
			CityId: "3371",
			Name:   "Magelang Selatan",
		},
		{
			Id:     "337102",
			CityId: "3371",
			Name:   "Magelang Utara",
		},
		{
			Id:     "337103",
			CityId: "3371",
			Name:   "Magelang Tengah",
		},
		{
			Id:     "337201",
			CityId: "3372",
			Name:   "Laweyan",
		},
		{
			Id:     "337202",
			CityId: "3372",
			Name:   "Serengan",
		},
		{
			Id:     "337203",
			CityId: "3372",
			Name:   "Pasar Kliwon",
		},
		{
			Id:     "337204",
			CityId: "3372",
			Name:   "Jebres",
		},
		{
			Id:     "337205",
			CityId: "3372",
			Name:   "Banjarsari",
		},
		{
			Id:     "337301",
			CityId: "3373",
			Name:   "Sidorejo",
		},
		{
			Id:     "337302",
			CityId: "3373",
			Name:   "Tingkir",
		},
		{
			Id:     "337303",
			CityId: "3373",
			Name:   "Argomulyo",
		},
		{
			Id:     "337304",
			CityId: "3373",
			Name:   "Sidomukti",
		},
		{
			Id:     "337401",
			CityId: "3374",
			Name:   "Semarang Tengah",
		},
		{
			Id:     "337402",
			CityId: "3374",
			Name:   "Semarang Utara",
		},
		{
			Id:     "337403",
			CityId: "3374",
			Name:   "Semarang Timur",
		},
		{
			Id:     "337404",
			CityId: "3374",
			Name:   "Gayamsari",
		},
		{
			Id:     "337405",
			CityId: "3374",
			Name:   "Genuk",
		},
		{
			Id:     "337406",
			CityId: "3374",
			Name:   "Pedurungan",
		},
		{
			Id:     "337407",
			CityId: "3374",
			Name:   "Semarang Selatan",
		},
		{
			Id:     "337408",
			CityId: "3374",
			Name:   "Candisari",
		},
		{
			Id:     "337409",
			CityId: "3374",
			Name:   "Gajahmungkur",
		},
		{
			Id:     "337410",
			CityId: "3374",
			Name:   "Tembalang",
		},
		{
			Id:     "337411",
			CityId: "3374",
			Name:   "Banyumanik",
		},
		{
			Id:     "337412",
			CityId: "3374",
			Name:   "Gunungpati",
		},
		{
			Id:     "337413",
			CityId: "3374",
			Name:   "Semarang Barat",
		},
		{
			Id:     "337414",
			CityId: "3374",
			Name:   "Mijen",
		},
		{
			Id:     "337415",
			CityId: "3374",
			Name:   "Ngaliyan",
		},
		{
			Id:     "337416",
			CityId: "3374",
			Name:   "Tugu",
		},
		{
			Id:     "337501",
			CityId: "3375",
			Name:   "Pekalongan Barat",
		},
		{
			Id:     "337502",
			CityId: "3375",
			Name:   "Pekalongan Timur",
		},
		{
			Id:     "337503",
			CityId: "3375",
			Name:   "Pekalongan Utara",
		},
		{
			Id:     "337504",
			CityId: "3375",
			Name:   "Pekalongan Selatan",
		},
		{
			Id:     "337601",
			CityId: "3376",
			Name:   "Tegal Barat",
		},
		{
			Id:     "337602",
			CityId: "3376",
			Name:   "Tegal Timur",
		},
		{
			Id:     "337603",
			CityId: "3376",
			Name:   "Tegal Selatan",
		},
		{
			Id:     "337604",
			CityId: "3376",
			Name:   "Margadana",
		},
		{
			Id:     "340101",
			CityId: "3401",
			Name:   "Temon",
		},
		{
			Id:     "340102",
			CityId: "3401",
			Name:   "Wates",
		},
		{
			Id:     "340103",
			CityId: "3401",
			Name:   "Panjatan",
		},
		{
			Id:     "340104",
			CityId: "3401",
			Name:   "Galur",
		},
		{
			Id:     "340105",
			CityId: "3401",
			Name:   "Lendah",
		},
		{
			Id:     "340106",
			CityId: "3401",
			Name:   "Sentolo",
		},
		{
			Id:     "340107",
			CityId: "3401",
			Name:   "Pengasih",
		},
		{
			Id:     "340108",
			CityId: "3401",
			Name:   "Kokap",
		},
		{
			Id:     "340109",
			CityId: "3401",
			Name:   "Girimulyo",
		},
		{
			Id:     "340110",
			CityId: "3401",
			Name:   "Nanggulan",
		},
		{
			Id:     "340111",
			CityId: "3401",
			Name:   "Samigaluh",
		},
		{
			Id:     "340112",
			CityId: "3401",
			Name:   "Kalibawang",
		},
		{
			Id:     "340201",
			CityId: "3402",
			Name:   "Srandakan",
		},
		{
			Id:     "340202",
			CityId: "3402",
			Name:   "Sanden",
		},
		{
			Id:     "340203",
			CityId: "3402",
			Name:   "Kretek",
		},
		{
			Id:     "340204",
			CityId: "3402",
			Name:   "Pundong",
		},
		{
			Id:     "340205",
			CityId: "3402",
			Name:   "Bambanglipuro",
		},
		{
			Id:     "340206",
			CityId: "3402",
			Name:   "Pandak",
		},
		{
			Id:     "340207",
			CityId: "3402",
			Name:   "Pajangan",
		},
		{
			Id:     "340208",
			CityId: "3402",
			Name:   "Bantul",
		},
		{
			Id:     "340209",
			CityId: "3402",
			Name:   "Jetis",
		},
		{
			Id:     "340210",
			CityId: "3402",
			Name:   "Imogiri",
		},
		{
			Id:     "340211",
			CityId: "3402",
			Name:   "Dlingo",
		},
		{
			Id:     "340212",
			CityId: "3402",
			Name:   "Banguntapan",
		},
		{
			Id:     "340213",
			CityId: "3402",
			Name:   "Pleret",
		},
		{
			Id:     "340214",
			CityId: "3402",
			Name:   "Piyungan",
		},
		{
			Id:     "340215",
			CityId: "3402",
			Name:   "Sewon",
		},
		{
			Id:     "340216",
			CityId: "3402",
			Name:   "Kasihan",
		},
		{
			Id:     "340217",
			CityId: "3402",
			Name:   "Sedayu",
		},
		{
			Id:     "340301",
			CityId: "3403",
			Name:   "Wonosari",
		},
		{
			Id:     "340302",
			CityId: "3403",
			Name:   "Nglipar",
		},
		{
			Id:     "340303",
			CityId: "3403",
			Name:   "Playen",
		},
		{
			Id:     "340304",
			CityId: "3403",
			Name:   "Patuk",
		},
		{
			Id:     "340305",
			CityId: "3403",
			Name:   "Paliyan",
		},
		{
			Id:     "340306",
			CityId: "3403",
			Name:   "Panggang",
		},
		{
			Id:     "340307",
			CityId: "3403",
			Name:   "Tepus",
		},
		{
			Id:     "340308",
			CityId: "3403",
			Name:   "Semanu",
		},
		{
			Id:     "340309",
			CityId: "3403",
			Name:   "Karangmojo",
		},
		{
			Id:     "340310",
			CityId: "3403",
			Name:   "Ponjong",
		},
		{
			Id:     "340311",
			CityId: "3403",
			Name:   "Rongkop",
		},
		{
			Id:     "340312",
			CityId: "3403",
			Name:   "Semin",
		},
		{
			Id:     "340313",
			CityId: "3403",
			Name:   "Ngawen",
		},
		{
			Id:     "340314",
			CityId: "3403",
			Name:   "Gedangsari",
		},
		{
			Id:     "340315",
			CityId: "3403",
			Name:   "Saptosari",
		},
		{
			Id:     "340316",
			CityId: "3403",
			Name:   "Girisubo",
		},
		{
			Id:     "340317",
			CityId: "3403",
			Name:   "Tanjungsari",
		},
		{
			Id:     "340318",
			CityId: "3403",
			Name:   "Purwosari",
		},
		{
			Id:     "340401",
			CityId: "3404",
			Name:   "Gamping",
		},
		{
			Id:     "340402",
			CityId: "3404",
			Name:   "Godean",
		},
		{
			Id:     "340403",
			CityId: "3404",
			Name:   "Moyudan",
		},
		{
			Id:     "340404",
			CityId: "3404",
			Name:   "Minggir",
		},
		{
			Id:     "340405",
			CityId: "3404",
			Name:   "Seyegan",
		},
		{
			Id:     "340406",
			CityId: "3404",
			Name:   "Mlati",
		},
		{
			Id:     "340407",
			CityId: "3404",
			Name:   "Depok",
		},
		{
			Id:     "340408",
			CityId: "3404",
			Name:   "Berbah",
		},
		{
			Id:     "340409",
			CityId: "3404",
			Name:   "Prambanan",
		},
		{
			Id:     "340410",
			CityId: "3404",
			Name:   "Kalasan",
		},
		{
			Id:     "340411",
			CityId: "3404",
			Name:   "Ngemplak",
		},
		{
			Id:     "340412",
			CityId: "3404",
			Name:   "Ngaglik",
		},
		{
			Id:     "340413",
			CityId: "3404",
			Name:   "Sleman",
		},
		{
			Id:     "340414",
			CityId: "3404",
			Name:   "Tempel",
		},
		{
			Id:     "340415",
			CityId: "3404",
			Name:   "Turi",
		},
		{
			Id:     "340416",
			CityId: "3404",
			Name:   "Pakem",
		},
		{
			Id:     "340417",
			CityId: "3404",
			Name:   "Cangkringan",
		},
		{
			Id:     "347101",
			CityId: "3471",
			Name:   "Tegalrejo",
		},
		{
			Id:     "347102",
			CityId: "3471",
			Name:   "Jetis",
		},
		{
			Id:     "347103",
			CityId: "3471",
			Name:   "Gondokusuman",
		},
		{
			Id:     "347104",
			CityId: "3471",
			Name:   "Danurejan",
		},
		{
			Id:     "347105",
			CityId: "3471",
			Name:   "Gedongtengen",
		},
		{
			Id:     "347106",
			CityId: "3471",
			Name:   "Ngampilan",
		},
		{
			Id:     "347107",
			CityId: "3471",
			Name:   "Wirobrajan",
		},
		{
			Id:     "347108",
			CityId: "3471",
			Name:   "Mantrijeron",
		},
		{
			Id:     "347109",
			CityId: "3471",
			Name:   "Kraton",
		},
		{
			Id:     "347110",
			CityId: "3471",
			Name:   "Gondomanan",
		},
		{
			Id:     "347111",
			CityId: "3471",
			Name:   "Pakualaman",
		},
		{
			Id:     "347112",
			CityId: "3471",
			Name:   "Mergangsan",
		},
		{
			Id:     "347113",
			CityId: "3471",
			Name:   "Umbulharjo",
		},
		{
			Id:     "347114",
			CityId: "3471",
			Name:   "Kotagede",
		},
		{
			Id:     "350101",
			CityId: "3501",
			Name:   "Donorojo",
		},
		{
			Id:     "350102",
			CityId: "3501",
			Name:   "Pringkuku",
		},
		{
			Id:     "350103",
			CityId: "3501",
			Name:   "Punung",
		},
		{
			Id:     "350104",
			CityId: "3501",
			Name:   "Pacitan",
		},
		{
			Id:     "350105",
			CityId: "3501",
			Name:   "Kebonagung",
		},
		{
			Id:     "350106",
			CityId: "3501",
			Name:   "Arjosari",
		},
		{
			Id:     "350107",
			CityId: "3501",
			Name:   "Nawangan",
		},
		{
			Id:     "350108",
			CityId: "3501",
			Name:   "Bandar",
		},
		{
			Id:     "350109",
			CityId: "3501",
			Name:   "Tegalombo",
		},
		{
			Id:     "350110",
			CityId: "3501",
			Name:   "Tulakan",
		},
		{
			Id:     "350111",
			CityId: "3501",
			Name:   "Ngadirojo",
		},
		{
			Id:     "350112",
			CityId: "3501",
			Name:   "Sudimoro",
		},
		{
			Id:     "350201",
			CityId: "3502",
			Name:   "Slahung",
		},
		{
			Id:     "350202",
			CityId: "3502",
			Name:   "Ngrayun",
		},
		{
			Id:     "350203",
			CityId: "3502",
			Name:   "Bungkal",
		},
		{
			Id:     "350204",
			CityId: "3502",
			Name:   "Sambit",
		},
		{
			Id:     "350205",
			CityId: "3502",
			Name:   "Sawoo",
		},
		{
			Id:     "350206",
			CityId: "3502",
			Name:   "Sooko",
		},
		{
			Id:     "350207",
			CityId: "3502",
			Name:   "Pulung",
		},
		{
			Id:     "350208",
			CityId: "3502",
			Name:   "Mlarak",
		},
		{
			Id:     "350209",
			CityId: "3502",
			Name:   "Jetis",
		},
		{
			Id:     "350210",
			CityId: "3502",
			Name:   "Siman",
		},
		{
			Id:     "350211",
			CityId: "3502",
			Name:   "Balong",
		},
		{
			Id:     "350212",
			CityId: "3502",
			Name:   "Kauman",
		},
		{
			Id:     "350213",
			CityId: "3502",
			Name:   "Badegan",
		},
		{
			Id:     "350214",
			CityId: "3502",
			Name:   "Sampung",
		},
		{
			Id:     "350215",
			CityId: "3502",
			Name:   "Sukorejo",
		},
		{
			Id:     "350216",
			CityId: "3502",
			Name:   "Babadan",
		},
		{
			Id:     "350217",
			CityId: "3502",
			Name:   "Ponorogo",
		},
		{
			Id:     "350218",
			CityId: "3502",
			Name:   "Jenangan",
		},
		{
			Id:     "350219",
			CityId: "3502",
			Name:   "Ngebel",
		},
		{
			Id:     "350220",
			CityId: "3502",
			Name:   "Jambon",
		},
		{
			Id:     "350221",
			CityId: "3502",
			Name:   "Pudak",
		},
		{
			Id:     "350301",
			CityId: "3503",
			Name:   "Panggul",
		},
		{
			Id:     "350302",
			CityId: "3503",
			Name:   "Munjungan",
		},
		{
			Id:     "350303",
			CityId: "3503",
			Name:   "Pule",
		},
		{
			Id:     "350304",
			CityId: "3503",
			Name:   "Dongko",
		},
		{
			Id:     "350305",
			CityId: "3503",
			Name:   "Tugu",
		},
		{
			Id:     "350306",
			CityId: "3503",
			Name:   "Karangan",
		},
		{
			Id:     "350307",
			CityId: "3503",
			Name:   "Kampak",
		},
		{
			Id:     "350308",
			CityId: "3503",
			Name:   "Watulimo",
		},
		{
			Id:     "350309",
			CityId: "3503",
			Name:   "Bendungan",
		},
		{
			Id:     "350310",
			CityId: "3503",
			Name:   "Gandusari",
		},
		{
			Id:     "350311",
			CityId: "3503",
			Name:   "Trenggalek",
		},
		{
			Id:     "350312",
			CityId: "3503",
			Name:   "Pogalan",
		},
		{
			Id:     "350313",
			CityId: "3503",
			Name:   "Durenan",
		},
		{
			Id:     "350314",
			CityId: "3503",
			Name:   "Suruh",
		},
		{
			Id:     "350401",
			CityId: "3504",
			Name:   "Tulungagung",
		},
		{
			Id:     "350402",
			CityId: "3504",
			Name:   "Boyolangu",
		},
		{
			Id:     "350403",
			CityId: "3504",
			Name:   "Kedungwaru",
		},
		{
			Id:     "350404",
			CityId: "3504",
			Name:   "Ngantru",
		},
		{
			Id:     "350405",
			CityId: "3504",
			Name:   "Kauman",
		},
		{
			Id:     "350406",
			CityId: "3504",
			Name:   "Pagerwojo",
		},
		{
			Id:     "350407",
			CityId: "3504",
			Name:   "Sendang",
		},
		{
			Id:     "350408",
			CityId: "3504",
			Name:   "Karangrejo",
		},
		{
			Id:     "350409",
			CityId: "3504",
			Name:   "Gondang",
		},
		{
			Id:     "350410",
			CityId: "3504",
			Name:   "Sumbergempol",
		},
		{
			Id:     "350411",
			CityId: "3504",
			Name:   "Ngunut",
		},
		{
			Id:     "350412",
			CityId: "3504",
			Name:   "Pucanglaban",
		},
		{
			Id:     "350413",
			CityId: "3504",
			Name:   "Rejotangan",
		},
		{
			Id:     "350414",
			CityId: "3504",
			Name:   "Kalidawir",
		},
		{
			Id:     "350415",
			CityId: "3504",
			Name:   "Besuki",
		},
		{
			Id:     "350416",
			CityId: "3504",
			Name:   "Campurdarat",
		},
		{
			Id:     "350417",
			CityId: "3504",
			Name:   "Bandung",
		},
		{
			Id:     "350418",
			CityId: "3504",
			Name:   "Pakel",
		},
		{
			Id:     "350419",
			CityId: "3504",
			Name:   "Tanggunggunung",
		},
		{
			Id:     "350501",
			CityId: "3505",
			Name:   "Wonodadi",
		},
		{
			Id:     "350502",
			CityId: "3505",
			Name:   "Udanawu",
		},
		{
			Id:     "350503",
			CityId: "3505",
			Name:   "Srengat",
		},
		{
			Id:     "350504",
			CityId: "3505",
			Name:   "Kademangan",
		},
		{
			Id:     "350505",
			CityId: "3505",
			Name:   "Bakung",
		},
		{
			Id:     "350506",
			CityId: "3505",
			Name:   "Ponggok",
		},
		{
			Id:     "350507",
			CityId: "3505",
			Name:   "Sanankulon",
		},
		{
			Id:     "350508",
			CityId: "3505",
			Name:   "Wonotirto",
		},
		{
			Id:     "350509",
			CityId: "3505",
			Name:   "Nglegok",
		},
		{
			Id:     "350510",
			CityId: "3505",
			Name:   "Kanigoro",
		},
		{
			Id:     "350511",
			CityId: "3505",
			Name:   "Garum",
		},
		{
			Id:     "350512",
			CityId: "3505",
			Name:   "Sutojayan",
		},
		{
			Id:     "350513",
			CityId: "3505",
			Name:   "Panggungrejo",
		},
		{
			Id:     "350514",
			CityId: "3505",
			Name:   "Talun",
		},
		{
			Id:     "350515",
			CityId: "3505",
			Name:   "Gandusari",
		},
		{
			Id:     "350516",
			CityId: "3505",
			Name:   "Binangun",
		},
		{
			Id:     "350517",
			CityId: "3505",
			Name:   "Wlingi",
		},
		{
			Id:     "350518",
			CityId: "3505",
			Name:   "Doko",
		},
		{
			Id:     "350519",
			CityId: "3505",
			Name:   "Kesamben",
		},
		{
			Id:     "350520",
			CityId: "3505",
			Name:   "Wates",
		},
		{
			Id:     "350521",
			CityId: "3505",
			Name:   "Selorejo",
		},
		{
			Id:     "350522",
			CityId: "3505",
			Name:   "Selopuro",
		},
		{
			Id:     "350601",
			CityId: "3506",
			Name:   "Semen",
		},
		{
			Id:     "350602",
			CityId: "3506",
			Name:   "Mojo",
		},
		{
			Id:     "350603",
			CityId: "3506",
			Name:   "Kras",
		},
		{
			Id:     "350604",
			CityId: "3506",
			Name:   "Ngadiluwih",
		},
		{
			Id:     "350605",
			CityId: "3506",
			Name:   "Kandat",
		},
		{
			Id:     "350606",
			CityId: "3506",
			Name:   "Wates",
		},
		{
			Id:     "350607",
			CityId: "3506",
			Name:   "Ngancar",
		},
		{
			Id:     "350608",
			CityId: "3506",
			Name:   "Puncu",
		},
		{
			Id:     "350609",
			CityId: "3506",
			Name:   "Plosoklaten",
		},
		{
			Id:     "350610",
			CityId: "3506",
			Name:   "Gurah",
		},
		{
			Id:     "350611",
			CityId: "3506",
			Name:   "Pagu",
		},
		{
			Id:     "350612",
			CityId: "3506",
			Name:   "Gampengrejo",
		},
		{
			Id:     "350613",
			CityId: "3506",
			Name:   "Grogol",
		},
		{
			Id:     "350614",
			CityId: "3506",
			Name:   "Papar",
		},
		{
			Id:     "350615",
			CityId: "3506",
			Name:   "Purwoasri",
		},
		{
			Id:     "350616",
			CityId: "3506",
			Name:   "Plemahan",
		},
		{
			Id:     "350617",
			CityId: "3506",
			Name:   "Pare",
		},
		{
			Id:     "350618",
			CityId: "3506",
			Name:   "Kepung",
		},
		{
			Id:     "350619",
			CityId: "3506",
			Name:   "Kandangan",
		},
		{
			Id:     "350620",
			CityId: "3506",
			Name:   "Tarokan",
		},
		{
			Id:     "350621",
			CityId: "3506",
			Name:   "Kunjang",
		},
		{
			Id:     "350622",
			CityId: "3506",
			Name:   "Banyakan",
		},
		{
			Id:     "350623",
			CityId: "3506",
			Name:   "Ringinrejo",
		},
		{
			Id:     "350624",
			CityId: "3506",
			Name:   "Kayen Kidul",
		},
		{
			Id:     "350625",
			CityId: "3506",
			Name:   "Ngasem",
		},
		{
			Id:     "350626",
			CityId: "3506",
			Name:   "Badas",
		},
		{
			Id:     "350701",
			CityId: "3507",
			Name:   "Donomulyo",
		},
		{
			Id:     "350702",
			CityId: "3507",
			Name:   "Pagak",
		},
		{
			Id:     "350703",
			CityId: "3507",
			Name:   "Bantur",
		},
		{
			Id:     "350704",
			CityId: "3507",
			Name:   "Sumbermanjing Wetan",
		},
		{
			Id:     "350705",
			CityId: "3507",
			Name:   "Dampit",
		},
		{
			Id:     "350706",
			CityId: "3507",
			Name:   "Ampelgading",
		},
		{
			Id:     "350707",
			CityId: "3507",
			Name:   "Poncokusumo",
		},
		{
			Id:     "350708",
			CityId: "3507",
			Name:   "Wajak",
		},
		{
			Id:     "350709",
			CityId: "3507",
			Name:   "Turen",
		},
		{
			Id:     "350710",
			CityId: "3507",
			Name:   "Gondanglegi",
		},
		{
			Id:     "350711",
			CityId: "3507",
			Name:   "Kalipare",
		},
		{
			Id:     "350712",
			CityId: "3507",
			Name:   "Sumberpucung",
		},
		{
			Id:     "350713",
			CityId: "3507",
			Name:   "Kepanjen",
		},
		{
			Id:     "350714",
			CityId: "3507",
			Name:   "Bululawang",
		},
		{
			Id:     "350715",
			CityId: "3507",
			Name:   "Tajinan",
		},
		{
			Id:     "350716",
			CityId: "3507",
			Name:   "Tumpang",
		},
		{
			Id:     "350717",
			CityId: "3507",
			Name:   "Jabung",
		},
		{
			Id:     "350718",
			CityId: "3507",
			Name:   "Pakis",
		},
		{
			Id:     "350719",
			CityId: "3507",
			Name:   "Pakisaji",
		},
		{
			Id:     "350720",
			CityId: "3507",
			Name:   "Ngajum",
		},
		{
			Id:     "350721",
			CityId: "3507",
			Name:   "Wagir",
		},
		{
			Id:     "350722",
			CityId: "3507",
			Name:   "Dau",
		},
		{
			Id:     "350723",
			CityId: "3507",
			Name:   "Karang Ploso",
		},
		{
			Id:     "350724",
			CityId: "3507",
			Name:   "Singosari",
		},
		{
			Id:     "350725",
			CityId: "3507",
			Name:   "Lawang",
		},
		{
			Id:     "350726",
			CityId: "3507",
			Name:   "Pujon",
		},
		{
			Id:     "350727",
			CityId: "3507",
			Name:   "Ngantang",
		},
		{
			Id:     "350728",
			CityId: "3507",
			Name:   "Kasembon",
		},
		{
			Id:     "350729",
			CityId: "3507",
			Name:   "Gedangan",
		},
		{
			Id:     "350730",
			CityId: "3507",
			Name:   "Tirtoyudo",
		},
		{
			Id:     "350731",
			CityId: "3507",
			Name:   "Kromengan",
		},
		{
			Id:     "350732",
			CityId: "3507",
			Name:   "Wonosari",
		},
		{
			Id:     "350733",
			CityId: "3507",
			Name:   "Pagelaran",
		},
		{
			Id:     "350801",
			CityId: "3508",
			Name:   "Tempursari",
		},
		{
			Id:     "350802",
			CityId: "3508",
			Name:   "Pronojiwo",
		},
		{
			Id:     "350803",
			CityId: "3508",
			Name:   "Candipuro",
		},
		{
			Id:     "350804",
			CityId: "3508",
			Name:   "Pasirian",
		},
		{
			Id:     "350805",
			CityId: "3508",
			Name:   "Tempeh",
		},
		{
			Id:     "350806",
			CityId: "3508",
			Name:   "Kunir",
		},
		{
			Id:     "350807",
			CityId: "3508",
			Name:   "Yosowilangun",
		},
		{
			Id:     "350808",
			CityId: "3508",
			Name:   "Rowokangkung",
		},
		{
			Id:     "350809",
			CityId: "3508",
			Name:   "Tekung",
		},
		{
			Id:     "350810",
			CityId: "3508",
			Name:   "Lumajang",
		},
		{
			Id:     "350811",
			CityId: "3508",
			Name:   "Pasrujambe",
		},
		{
			Id:     "350812",
			CityId: "3508",
			Name:   "Senduro",
		},
		{
			Id:     "350813",
			CityId: "3508",
			Name:   "Gucialit",
		},
		{
			Id:     "350814",
			CityId: "3508",
			Name:   "Padang",
		},
		{
			Id:     "350815",
			CityId: "3508",
			Name:   "Sukodono",
		},
		{
			Id:     "350816",
			CityId: "3508",
			Name:   "Kedungjajang",
		},
		{
			Id:     "350817",
			CityId: "3508",
			Name:   "Jatiroto",
		},
		{
			Id:     "350818",
			CityId: "3508",
			Name:   "Randuagung",
		},
		{
			Id:     "350819",
			CityId: "3508",
			Name:   "Klakah",
		},
		{
			Id:     "350820",
			CityId: "3508",
			Name:   "Ranuyoso",
		},
		{
			Id:     "350821",
			CityId: "3508",
			Name:   "Sumbersuko",
		},
		{
			Id:     "350901",
			CityId: "3509",
			Name:   "Jombang",
		},
		{
			Id:     "350902",
			CityId: "3509",
			Name:   "Kencong",
		},
		{
			Id:     "350903",
			CityId: "3509",
			Name:   "Sumberbaru",
		},
		{
			Id:     "350904",
			CityId: "3509",
			Name:   "Gumukmas",
		},
		{
			Id:     "350905",
			CityId: "3509",
			Name:   "Umbulsari",
		},
		{
			Id:     "350906",
			CityId: "3509",
			Name:   "Tanggul",
		},
		{
			Id:     "350907",
			CityId: "3509",
			Name:   "Semboro",
		},
		{
			Id:     "350908",
			CityId: "3509",
			Name:   "Puger",
		},
		{
			Id:     "350909",
			CityId: "3509",
			Name:   "Bangsalsari",
		},
		{
			Id:     "350910",
			CityId: "3509",
			Name:   "Balung",
		},
		{
			Id:     "350911",
			CityId: "3509",
			Name:   "Wuluhan",
		},
		{
			Id:     "350912",
			CityId: "3509",
			Name:   "Ambulu",
		},
		{
			Id:     "350913",
			CityId: "3509",
			Name:   "Rambipuji",
		},
		{
			Id:     "350914",
			CityId: "3509",
			Name:   "Panti",
		},
		{
			Id:     "350915",
			CityId: "3509",
			Name:   "Sukorambi",
		},
		{
			Id:     "350916",
			CityId: "3509",
			Name:   "Jenggawah",
		},
		{
			Id:     "350917",
			CityId: "3509",
			Name:   "Ajung",
		},
		{
			Id:     "350918",
			CityId: "3509",
			Name:   "Tempurejo",
		},
		{
			Id:     "350919",
			CityId: "3509",
			Name:   "Kaliwates",
		},
		{
			Id:     "350920",
			CityId: "3509",
			Name:   "Patrang",
		},
		{
			Id:     "350921",
			CityId: "3509",
			Name:   "Sumbersari",
		},
		{
			Id:     "350922",
			CityId: "3509",
			Name:   "Arjasa",
		},
		{
			Id:     "350923",
			CityId: "3509",
			Name:   "Mumbulsari",
		},
		{
			Id:     "350924",
			CityId: "3509",
			Name:   "Pakusari",
		},
		{
			Id:     "350925",
			CityId: "3509",
			Name:   "Jelbuk",
		},
		{
			Id:     "350926",
			CityId: "3509",
			Name:   "Mayang",
		},
		{
			Id:     "350927",
			CityId: "3509",
			Name:   "Kalisat",
		},
		{
			Id:     "350928",
			CityId: "3509",
			Name:   "Ledokombo",
		},
		{
			Id:     "350929",
			CityId: "3509",
			Name:   "Sukowono",
		},
		{
			Id:     "350930",
			CityId: "3509",
			Name:   "Silo",
		},
		{
			Id:     "350931",
			CityId: "3509",
			Name:   "Sumberjambe",
		},
		{
			Id:     "351001",
			CityId: "3510",
			Name:   "Pesanggaran",
		},
		{
			Id:     "351002",
			CityId: "3510",
			Name:   "Bangorejo",
		},
		{
			Id:     "351003",
			CityId: "3510",
			Name:   "Purwoharjo",
		},
		{
			Id:     "351004",
			CityId: "3510",
			Name:   "Tegaldlimo",
		},
		{
			Id:     "351005",
			CityId: "3510",
			Name:   "Muncar",
		},
		{
			Id:     "351006",
			CityId: "3510",
			Name:   "Cluring",
		},
		{
			Id:     "351007",
			CityId: "3510",
			Name:   "Gambiran",
		},
		{
			Id:     "351008",
			CityId: "3510",
			Name:   "Srono",
		},
		{
			Id:     "351009",
			CityId: "3510",
			Name:   "Genteng",
		},
		{
			Id:     "351010",
			CityId: "3510",
			Name:   "Glenmore",
		},
		{
			Id:     "351011",
			CityId: "3510",
			Name:   "Kalibaru",
		},
		{
			Id:     "351012",
			CityId: "3510",
			Name:   "Singojuruh",
		},
		{
			Id:     "351013",
			CityId: "3510",
			Name:   "Rogojampi",
		},
		{
			Id:     "351014",
			CityId: "3510",
			Name:   "Kabat",
		},
		{
			Id:     "351015",
			CityId: "3510",
			Name:   "Glagah",
		},
		{
			Id:     "351016",
			CityId: "3510",
			Name:   "Banyuwangi",
		},
		{
			Id:     "351017",
			CityId: "3510",
			Name:   "Giri",
		},
		{
			Id:     "351018",
			CityId: "3510",
			Name:   "Wongsorejo",
		},
		{
			Id:     "351019",
			CityId: "3510",
			Name:   "Songgon",
		},
		{
			Id:     "351020",
			CityId: "3510",
			Name:   "Sempu",
		},
		{
			Id:     "351021",
			CityId: "3510",
			Name:   "Kalipuro",
		},
		{
			Id:     "351022",
			CityId: "3510",
			Name:   "Siliragung",
		},
		{
			Id:     "351023",
			CityId: "3510",
			Name:   "Tegalsari",
		},
		{
			Id:     "351024",
			CityId: "3510",
			Name:   "Licin",
		},
		{
			Id:     "351025",
			CityId: "3510",
			Name:   "Blimbingsari",
		},
		{
			Id:     "351101",
			CityId: "3511",
			Name:   "Maesan",
		},
		{
			Id:     "351102",
			CityId: "3511",
			Name:   "Tamanan",
		},
		{
			Id:     "351103",
			CityId: "3511",
			Name:   "Tlogosari",
		},
		{
			Id:     "351104",
			CityId: "3511",
			Name:   "Sukosari",
		},
		{
			Id:     "351105",
			CityId: "3511",
			Name:   "Pujer",
		},
		{
			Id:     "351106",
			CityId: "3511",
			Name:   "Grujugan",
		},
		{
			Id:     "351107",
			CityId: "3511",
			Name:   "Curahdami",
		},
		{
			Id:     "351108",
			CityId: "3511",
			Name:   "Tenggarang",
		},
		{
			Id:     "351109",
			CityId: "3511",
			Name:   "Wonosari",
		},
		{
			Id:     "351110",
			CityId: "3511",
			Name:   "Tapen",
		},
		{
			Id:     "351111",
			CityId: "3511",
			Name:   "Bondowoso",
		},
		{
			Id:     "351112",
			CityId: "3511",
			Name:   "Wringin",
		},
		{
			Id:     "351113",
			CityId: "3511",
			Name:   "Tegalampel",
		},
		{
			Id:     "351114",
			CityId: "3511",
			Name:   "Klabang",
		},
		{
			Id:     "351115",
			CityId: "3511",
			Name:   "Cermee",
		},
		{
			Id:     "351116",
			CityId: "3511",
			Name:   "Prajekan",
		},
		{
			Id:     "351117",
			CityId: "3511",
			Name:   "Pakem",
		},
		{
			Id:     "351118",
			CityId: "3511",
			Name:   "Sumberwringin",
		},
		{
			Id:     "351119",
			CityId: "3511",
			Name:   "Sempol",
		},
		{
			Id:     "351120",
			CityId: "3511",
			Name:   "Binakal",
		},
		{
			Id:     "351121",
			CityId: "3511",
			Name:   "Taman Krocok",
		},
		{
			Id:     "351122",
			CityId: "3511",
			Name:   "Botolinggo",
		},
		{
			Id:     "351123",
			CityId: "3511",
			Name:   "Jambesari Darus Sholah",
		},
		{
			Id:     "351201",
			CityId: "3512",
			Name:   "Jatibanteng",
		},
		{
			Id:     "351202",
			CityId: "3512",
			Name:   "Besuki",
		},
		{
			Id:     "351203",
			CityId: "3512",
			Name:   "Suboh",
		},
		{
			Id:     "351204",
			CityId: "3512",
			Name:   "Mlandingan",
		},
		{
			Id:     "351205",
			CityId: "3512",
			Name:   "Kendit",
		},
		{
			Id:     "351206",
			CityId: "3512",
			Name:   "Panarukan",
		},
		{
			Id:     "351207",
			CityId: "3512",
			Name:   "Situbondo",
		},
		{
			Id:     "351208",
			CityId: "3512",
			Name:   "Panji",
		},
		{
			Id:     "351209",
			CityId: "3512",
			Name:   "Mangaran",
		},
		{
			Id:     "351210",
			CityId: "3512",
			Name:   "Kapongan",
		},
		{
			Id:     "351211",
			CityId: "3512",
			Name:   "Arjasa",
		},
		{
			Id:     "351212",
			CityId: "3512",
			Name:   "Jangkar",
		},
		{
			Id:     "351213",
			CityId: "3512",
			Name:   "Asembagus",
		},
		{
			Id:     "351214",
			CityId: "3512",
			Name:   "Banyuputih",
		},
		{
			Id:     "351215",
			CityId: "3512",
			Name:   "Sumbermalang",
		},
		{
			Id:     "351216",
			CityId: "3512",
			Name:   "Banyuglugur",
		},
		{
			Id:     "351217",
			CityId: "3512",
			Name:   "Bungatan",
		},
		{
			Id:     "351301",
			CityId: "3513",
			Name:   "Sukapura",
		},
		{
			Id:     "351302",
			CityId: "3513",
			Name:   "Sumber",
		},
		{
			Id:     "351303",
			CityId: "3513",
			Name:   "Kuripan",
		},
		{
			Id:     "351304",
			CityId: "3513",
			Name:   "Bantaran",
		},
		{
			Id:     "351305",
			CityId: "3513",
			Name:   "Leces",
		},
		{
			Id:     "351306",
			CityId: "3513",
			Name:   "Banyuanyar",
		},
		{
			Id:     "351307",
			CityId: "3513",
			Name:   "Tiris",
		},
		{
			Id:     "351308",
			CityId: "3513",
			Name:   "Krucil",
		},
		{
			Id:     "351309",
			CityId: "3513",
			Name:   "Gading",
		},
		{
			Id:     "351310",
			CityId: "3513",
			Name:   "Pakuniran",
		},
		{
			Id:     "351311",
			CityId: "3513",
			Name:   "Kotaanyar",
		},
		{
			Id:     "351312",
			CityId: "3513",
			Name:   "Paiton",
		},
		{
			Id:     "351313",
			CityId: "3513",
			Name:   "Besuk",
		},
		{
			Id:     "351314",
			CityId: "3513",
			Name:   "Kraksaan",
		},
		{
			Id:     "351315",
			CityId: "3513",
			Name:   "Krejengan",
		},
		{
			Id:     "351316",
			CityId: "3513",
			Name:   "Pejarakan",
		},
		{
			Id:     "351317",
			CityId: "3513",
			Name:   "Maron",
		},
		{
			Id:     "351318",
			CityId: "3513",
			Name:   "Gending",
		},
		{
			Id:     "351319",
			CityId: "3513",
			Name:   "Dringu",
		},
		{
			Id:     "351320",
			CityId: "3513",
			Name:   "Tegalsiwalan",
		},
		{
			Id:     "351321",
			CityId: "3513",
			Name:   "Sumberasih",
		},
		{
			Id:     "351322",
			CityId: "3513",
			Name:   "Wonomerto",
		},
		{
			Id:     "351323",
			CityId: "3513",
			Name:   "Tongas",
		},
		{
			Id:     "351324",
			CityId: "3513",
			Name:   "Lumbang",
		},
		{
			Id:     "351401",
			CityId: "3514",
			Name:   "Purwodadi",
		},
		{
			Id:     "351402",
			CityId: "3514",
			Name:   "Tutur",
		},
		{
			Id:     "351403",
			CityId: "3514",
			Name:   "Puspo",
		},
		{
			Id:     "351404",
			CityId: "3514",
			Name:   "Lumbang",
		},
		{
			Id:     "351405",
			CityId: "3514",
			Name:   "Pasrepan",
		},
		{
			Id:     "351406",
			CityId: "3514",
			Name:   "Kejayan",
		},
		{
			Id:     "351407",
			CityId: "3514",
			Name:   "Wonorejo",
		},
		{
			Id:     "351408",
			CityId: "3514",
			Name:   "Purwosari",
		},
		{
			Id:     "351409",
			CityId: "3514",
			Name:   "Sukorejo",
		},
		{
			Id:     "351410",
			CityId: "3514",
			Name:   "Prigen",
		},
		{
			Id:     "351411",
			CityId: "3514",
			Name:   "Pandaan",
		},
		{
			Id:     "351412",
			CityId: "3514",
			Name:   "Gempol",
		},
		{
			Id:     "351413",
			CityId: "3514",
			Name:   "Beji",
		},
		{
			Id:     "351414",
			CityId: "3514",
			Name:   "Bangil",
		},
		{
			Id:     "351415",
			CityId: "3514",
			Name:   "Rembang",
		},
		{
			Id:     "351416",
			CityId: "3514",
			Name:   "Kraton",
		},
		{
			Id:     "351417",
			CityId: "3514",
			Name:   "Pohjentrek",
		},
		{
			Id:     "351418",
			CityId: "3514",
			Name:   "Gondangwetan",
		},
		{
			Id:     "351419",
			CityId: "3514",
			Name:   "Winongan",
		},
		{
			Id:     "351420",
			CityId: "3514",
			Name:   "Grati",
		},
		{
			Id:     "351421",
			CityId: "3514",
			Name:   "Nguling",
		},
		{
			Id:     "351422",
			CityId: "3514",
			Name:   "Lekok",
		},
		{
			Id:     "351423",
			CityId: "3514",
			Name:   "Rejoso",
		},
		{
			Id:     "351424",
			CityId: "3514",
			Name:   "Tosari",
		},
		{
			Id:     "351501",
			CityId: "3515",
			Name:   "Tarik",
		},
		{
			Id:     "351502",
			CityId: "3515",
			Name:   "Prambon",
		},
		{
			Id:     "351503",
			CityId: "3515",
			Name:   "Krembung",
		},
		{
			Id:     "351504",
			CityId: "3515",
			Name:   "Porong",
		},
		{
			Id:     "351505",
			CityId: "3515",
			Name:   "Jabon",
		},
		{
			Id:     "351506",
			CityId: "3515",
			Name:   "Tanggulangin",
		},
		{
			Id:     "351507",
			CityId: "3515",
			Name:   "Candi",
		},
		{
			Id:     "351508",
			CityId: "3515",
			Name:   "Sidoarjo",
		},
		{
			Id:     "351509",
			CityId: "3515",
			Name:   "Tulangan",
		},
		{
			Id:     "351510",
			CityId: "3515",
			Name:   "Wonoayu",
		},
		{
			Id:     "351511",
			CityId: "3515",
			Name:   "Krian",
		},
		{
			Id:     "351512",
			CityId: "3515",
			Name:   "Balongbendo",
		},
		{
			Id:     "351513",
			CityId: "3515",
			Name:   "Taman",
		},
		{
			Id:     "351514",
			CityId: "3515",
			Name:   "Sukodono",
		},
		{
			Id:     "351515",
			CityId: "3515",
			Name:   "Buduran",
		},
		{
			Id:     "351516",
			CityId: "3515",
			Name:   "Gedangan",
		},
		{
			Id:     "351517",
			CityId: "3515",
			Name:   "Sedati",
		},
		{
			Id:     "351518",
			CityId: "3515",
			Name:   "Waru",
		},
		{
			Id:     "351601",
			CityId: "3516",
			Name:   "Jatirejo",
		},
		{
			Id:     "351602",
			CityId: "3516",
			Name:   "Gondang",
		},
		{
			Id:     "351603",
			CityId: "3516",
			Name:   "Pacet",
		},
		{
			Id:     "351604",
			CityId: "3516",
			Name:   "Trawas",
		},
		{
			Id:     "351605",
			CityId: "3516",
			Name:   "Ngoro",
		},
		{
			Id:     "351606",
			CityId: "3516",
			Name:   "Pungging",
		},
		{
			Id:     "351607",
			CityId: "3516",
			Name:   "Kutorejo",
		},
		{
			Id:     "351608",
			CityId: "3516",
			Name:   "Mojosari",
		},
		{
			Id:     "351609",
			CityId: "3516",
			Name:   "Dlanggu",
		},
		{
			Id:     "351610",
			CityId: "3516",
			Name:   "Bangsal",
		},
		{
			Id:     "351611",
			CityId: "3516",
			Name:   "Puri",
		},
		{
			Id:     "351612",
			CityId: "3516",
			Name:   "Trowulan",
		},
		{
			Id:     "351613",
			CityId: "3516",
			Name:   "Sooko",
		},
		{
			Id:     "351614",
			CityId: "3516",
			Name:   "Gedeg",
		},
		{
			Id:     "351615",
			CityId: "3516",
			Name:   "Kemlagi",
		},
		{
			Id:     "351616",
			CityId: "3516",
			Name:   "Jetis",
		},
		{
			Id:     "351617",
			CityId: "3516",
			Name:   "Dawarblandong",
		},
		{
			Id:     "351618",
			CityId: "3516",
			Name:   "Mojoanyar",
		},
		{
			Id:     "351701",
			CityId: "3517",
			Name:   "Perak",
		},
		{
			Id:     "351702",
			CityId: "3517",
			Name:   "Gudo",
		},
		{
			Id:     "351703",
			CityId: "3517",
			Name:   "Ngoro",
		},
		{
			Id:     "351704",
			CityId: "3517",
			Name:   "Bareng",
		},
		{
			Id:     "351705",
			CityId: "3517",
			Name:   "Wonosalam",
		},
		{
			Id:     "351706",
			CityId: "3517",
			Name:   "Mojoagung",
		},
		{
			Id:     "351707",
			CityId: "3517",
			Name:   "Mojowarno",
		},
		{
			Id:     "351708",
			CityId: "3517",
			Name:   "Diwek",
		},
		{
			Id:     "351709",
			CityId: "3517",
			Name:   "Jombang",
		},
		{
			Id:     "351710",
			CityId: "3517",
			Name:   "Peterongan",
		},
		{
			Id:     "351711",
			CityId: "3517",
			Name:   "Sumobito",
		},
		{
			Id:     "351712",
			CityId: "3517",
			Name:   "Kesamben",
		},
		{
			Id:     "351713",
			CityId: "3517",
			Name:   "Tembelang",
		},
		{
			Id:     "351714",
			CityId: "3517",
			Name:   "Ploso",
		},
		{
			Id:     "351715",
			CityId: "3517",
			Name:   "Plandaan",
		},
		{
			Id:     "351716",
			CityId: "3517",
			Name:   "Kabuh",
		},
		{
			Id:     "351717",
			CityId: "3517",
			Name:   "Kudu",
		},
		{
			Id:     "351718",
			CityId: "3517",
			Name:   "Bandarkedungmulyo",
		},
		{
			Id:     "351719",
			CityId: "3517",
			Name:   "Jogoroto",
		},
		{
			Id:     "351720",
			CityId: "3517",
			Name:   "Megaluh",
		},
		{
			Id:     "351721",
			CityId: "3517",
			Name:   "Ngusikan",
		},
		{
			Id:     "351801",
			CityId: "3518",
			Name:   "Sawahan",
		},
		{
			Id:     "351802",
			CityId: "3518",
			Name:   "Ngetos",
		},
		{
			Id:     "351803",
			CityId: "3518",
			Name:   "Berbek",
		},
		{
			Id:     "351804",
			CityId: "3518",
			Name:   "Loceret",
		},
		{
			Id:     "351805",
			CityId: "3518",
			Name:   "Pace",
		},
		{
			Id:     "351806",
			CityId: "3518",
			Name:   "Prambon",
		},
		{
			Id:     "351807",
			CityId: "3518",
			Name:   "Ngronggot",
		},
		{
			Id:     "351808",
			CityId: "3518",
			Name:   "Kertosono",
		},
		{
			Id:     "351809",
			CityId: "3518",
			Name:   "Patianrowo",
		},
		{
			Id:     "351810",
			CityId: "3518",
			Name:   "Baron",
		},
		{
			Id:     "351811",
			CityId: "3518",
			Name:   "Tanjunganom",
		},
		{
			Id:     "351812",
			CityId: "3518",
			Name:   "Sukomoro",
		},
		{
			Id:     "351813",
			CityId: "3518",
			Name:   "Nganjuk",
		},
		{
			Id:     "351814",
			CityId: "3518",
			Name:   "Bagor",
		},
		{
			Id:     "351815",
			CityId: "3518",
			Name:   "Wilangan",
		},
		{
			Id:     "351816",
			CityId: "3518",
			Name:   "Rejoso",
		},
		{
			Id:     "351817",
			CityId: "3518",
			Name:   "Gondang",
		},
		{
			Id:     "351818",
			CityId: "3518",
			Name:   "Ngluyu",
		},
		{
			Id:     "351819",
			CityId: "3518",
			Name:   "Lengkong",
		},
		{
			Id:     "351820",
			CityId: "3518",
			Name:   "Jatikalen",
		},
		{
			Id:     "351901",
			CityId: "3519",
			Name:   "Kebonsari",
		},
		{
			Id:     "351902",
			CityId: "3519",
			Name:   "Dolopo",
		},
		{
			Id:     "351903",
			CityId: "3519",
			Name:   "Geger",
		},
		{
			Id:     "351904",
			CityId: "3519",
			Name:   "Dagangan",
		},
		{
			Id:     "351905",
			CityId: "3519",
			Name:   "Kare",
		},
		{
			Id:     "351906",
			CityId: "3519",
			Name:   "Gemarang",
		},
		{
			Id:     "351907",
			CityId: "3519",
			Name:   "Wungu",
		},
		{
			Id:     "351908",
			CityId: "3519",
			Name:   "Madiun",
		},
		{
			Id:     "351909",
			CityId: "3519",
			Name:   "Jiwan",
		},
		{
			Id:     "351910",
			CityId: "3519",
			Name:   "Balerejo",
		},
		{
			Id:     "351911",
			CityId: "3519",
			Name:   "Mejayan",
		},
		{
			Id:     "351912",
			CityId: "3519",
			Name:   "Saradan",
		},
		{
			Id:     "351913",
			CityId: "3519",
			Name:   "Pilangkenceng",
		},
		{
			Id:     "351914",
			CityId: "3519",
			Name:   "Sawahan",
		},
		{
			Id:     "351915",
			CityId: "3519",
			Name:   "Wonoasri",
		},
		{
			Id:     "352001",
			CityId: "3520",
			Name:   "Poncol",
		},
		{
			Id:     "352002",
			CityId: "3520",
			Name:   "Parang",
		},
		{
			Id:     "352003",
			CityId: "3520",
			Name:   "Lembeyan",
		},
		{
			Id:     "352004",
			CityId: "3520",
			Name:   "Takeran",
		},
		{
			Id:     "352005",
			CityId: "3520",
			Name:   "Kawedanan",
		},
		{
			Id:     "352006",
			CityId: "3520",
			Name:   "Magetan",
		},
		{
			Id:     "352007",
			CityId: "3520",
			Name:   "Plaosan",
		},
		{
			Id:     "352008",
			CityId: "3520",
			Name:   "Panekan",
		},
		{
			Id:     "352009",
			CityId: "3520",
			Name:   "Sukomoro",
		},
		{
			Id:     "352010",
			CityId: "3520",
			Name:   "Bendo",
		},
		{
			Id:     "352011",
			CityId: "3520",
			Name:   "Maospati",
		},
		{
			Id:     "352012",
			CityId: "3520",
			Name:   "Barat",
		},
		{
			Id:     "352013",
			CityId: "3520",
			Name:   "Karangrejo",
		},
		{
			Id:     "352014",
			CityId: "3520",
			Name:   "Karas",
		},
		{
			Id:     "352015",
			CityId: "3520",
			Name:   "Kartoharjo",
		},
		{
			Id:     "352016",
			CityId: "3520",
			Name:   "Ngariboyo",
		},
		{
			Id:     "352017",
			CityId: "3520",
			Name:   "Nguntoronadi",
		},
		{
			Id:     "352018",
			CityId: "3520",
			Name:   "Sidorejo",
		},
		{
			Id:     "352101",
			CityId: "3521",
			Name:   "Sine",
		},
		{
			Id:     "352102",
			CityId: "3521",
			Name:   "Ngrambe",
		},
		{
			Id:     "352103",
			CityId: "3521",
			Name:   "Jogorogo",
		},
		{
			Id:     "352104",
			CityId: "3521",
			Name:   "Kendal",
		},
		{
			Id:     "352105",
			CityId: "3521",
			Name:   "Geneng",
		},
		{
			Id:     "352106",
			CityId: "3521",
			Name:   "Kwadungan",
		},
		{
			Id:     "352107",
			CityId: "3521",
			Name:   "Karangjati",
		},
		{
			Id:     "352108",
			CityId: "3521",
			Name:   "Padas",
		},
		{
			Id:     "352109",
			CityId: "3521",
			Name:   "Ngawi",
		},
		{
			Id:     "352110",
			CityId: "3521",
			Name:   "Paron",
		},
		{
			Id:     "352111",
			CityId: "3521",
			Name:   "Kedunggalar",
		},
		{
			Id:     "352112",
			CityId: "3521",
			Name:   "Widodaren",
		},
		{
			Id:     "352113",
			CityId: "3521",
			Name:   "Mantingan",
		},
		{
			Id:     "352114",
			CityId: "3521",
			Name:   "Pangkur",
		},
		{
			Id:     "352115",
			CityId: "3521",
			Name:   "Bringin",
		},
		{
			Id:     "352116",
			CityId: "3521",
			Name:   "Pitu",
		},
		{
			Id:     "352117",
			CityId: "3521",
			Name:   "Karanganyar",
		},
		{
			Id:     "352118",
			CityId: "3521",
			Name:   "Gerih",
		},
		{
			Id:     "352119",
			CityId: "3521",
			Name:   "Kasreman",
		},
		{
			Id:     "352201",
			CityId: "3522",
			Name:   "Ngraho",
		},
		{
			Id:     "352202",
			CityId: "3522",
			Name:   "Tambakrejo",
		},
		{
			Id:     "352203",
			CityId: "3522",
			Name:   "Ngambon",
		},
		{
			Id:     "352204",
			CityId: "3522",
			Name:   "Ngasem",
		},
		{
			Id:     "352205",
			CityId: "3522",
			Name:   "Bubulan",
		},
		{
			Id:     "352206",
			CityId: "3522",
			Name:   "Dander",
		},
		{
			Id:     "352207",
			CityId: "3522",
			Name:   "Sugihwaras",
		},
		{
			Id:     "352208",
			CityId: "3522",
			Name:   "Kedungadem",
		},
		{
			Id:     "352209",
			CityId: "3522",
			Name:   "Kepohbaru",
		},
		{
			Id:     "352210",
			CityId: "3522",
			Name:   "Baureno",
		},
		{
			Id:     "352211",
			CityId: "3522",
			Name:   "Kanor",
		},
		{
			Id:     "352212",
			CityId: "3522",
			Name:   "Sumberejo",
		},
		{
			Id:     "352213",
			CityId: "3522",
			Name:   "Balen",
		},
		{
			Id:     "352214",
			CityId: "3522",
			Name:   "Kapas",
		},
		{
			Id:     "352215",
			CityId: "3522",
			Name:   "Bojonegoro",
		},
		{
			Id:     "352216",
			CityId: "3522",
			Name:   "Kalitidu",
		},
		{
			Id:     "352217",
			CityId: "3522",
			Name:   "Malo",
		},
		{
			Id:     "352218",
			CityId: "3522",
			Name:   "Purwosari",
		},
		{
			Id:     "352219",
			CityId: "3522",
			Name:   "Padangan",
		},
		{
			Id:     "352220",
			CityId: "3522",
			Name:   "Kasiman",
		},
		{
			Id:     "352221",
			CityId: "3522",
			Name:   "Temayang",
		},
		{
			Id:     "352222",
			CityId: "3522",
			Name:   "Margomulyo",
		},
		{
			Id:     "352223",
			CityId: "3522",
			Name:   "Trucuk",
		},
		{
			Id:     "352224",
			CityId: "3522",
			Name:   "Sukosewu",
		},
		{
			Id:     "352225",
			CityId: "3522",
			Name:   "Kedewan",
		},
		{
			Id:     "352226",
			CityId: "3522",
			Name:   "Gondang",
		},
		{
			Id:     "352227",
			CityId: "3522",
			Name:   "Sekar",
		},
		{
			Id:     "352228",
			CityId: "3522",
			Name:   "Gayam",
		},
		{
			Id:     "352301",
			CityId: "3523",
			Name:   "Kenduruan",
		},
		{
			Id:     "352302",
			CityId: "3523",
			Name:   "Jatirogo",
		},
		{
			Id:     "352303",
			CityId: "3523",
			Name:   "Bangilan",
		},
		{
			Id:     "352304",
			CityId: "3523",
			Name:   "Bancar",
		},
		{
			Id:     "352305",
			CityId: "3523",
			Name:   "Senori",
		},
		{
			Id:     "352306",
			CityId: "3523",
			Name:   "Tambakboyo",
		},
		{
			Id:     "352307",
			CityId: "3523",
			Name:   "Singgahan",
		},
		{
			Id:     "352308",
			CityId: "3523",
			Name:   "Kerek",
		},
		{
			Id:     "352309",
			CityId: "3523",
			Name:   "Parengan",
		},
		{
			Id:     "352310",
			CityId: "3523",
			Name:   "Montong",
		},
		{
			Id:     "352311",
			CityId: "3523",
			Name:   "Soko",
		},
		{
			Id:     "352312",
			CityId: "3523",
			Name:   "Jenu",
		},
		{
			Id:     "352313",
			CityId: "3523",
			Name:   "Merakurak",
		},
		{
			Id:     "352314",
			CityId: "3523",
			Name:   "Rengel",
		},
		{
			Id:     "352315",
			CityId: "3523",
			Name:   "Semanding",
		},
		{
			Id:     "352316",
			CityId: "3523",
			Name:   "Tuban",
		},
		{
			Id:     "352317",
			CityId: "3523",
			Name:   "Plumpang",
		},
		{
			Id:     "352318",
			CityId: "3523",
			Name:   "Palang",
		},
		{
			Id:     "352319",
			CityId: "3523",
			Name:   "Widang",
		},
		{
			Id:     "352320",
			CityId: "3523",
			Name:   "Grabagan",
		},
		{
			Id:     "352401",
			CityId: "3524",
			Name:   "Sukorame",
		},
		{
			Id:     "352402",
			CityId: "3524",
			Name:   "Bluluk",
		},
		{
			Id:     "352403",
			CityId: "3524",
			Name:   "Modo",
		},
		{
			Id:     "352404",
			CityId: "3524",
			Name:   "Ngimbang",
		},
		{
			Id:     "352405",
			CityId: "3524",
			Name:   "Babat",
		},
		{
			Id:     "352406",
			CityId: "3524",
			Name:   "Kedungpring",
		},
		{
			Id:     "352407",
			CityId: "3524",
			Name:   "Brondong",
		},
		{
			Id:     "352408",
			CityId: "3524",
			Name:   "Laren",
		},
		{
			Id:     "352409",
			CityId: "3524",
			Name:   "Sekaran",
		},
		{
			Id:     "352410",
			CityId: "3524",
			Name:   "Maduran",
		},
		{
			Id:     "352411",
			CityId: "3524",
			Name:   "Sambeng",
		},
		{
			Id:     "352412",
			CityId: "3524",
			Name:   "Sugio",
		},
		{
			Id:     "352413",
			CityId: "3524",
			Name:   "Pucuk",
		},
		{
			Id:     "352414",
			CityId: "3524",
			Name:   "Paciran",
		},
		{
			Id:     "352415",
			CityId: "3524",
			Name:   "Solokuro",
		},
		{
			Id:     "352416",
			CityId: "3524",
			Name:   "Mantup",
		},
		{
			Id:     "352417",
			CityId: "3524",
			Name:   "Sukodadi",
		},
		{
			Id:     "352418",
			CityId: "3524",
			Name:   "Karanggeneng",
		},
		{
			Id:     "352419",
			CityId: "3524",
			Name:   "Kembangbahu",
		},
		{
			Id:     "352420",
			CityId: "3524",
			Name:   "Kalitengah",
		},
		{
			Id:     "352421",
			CityId: "3524",
			Name:   "Turi",
		},
		{
			Id:     "352422",
			CityId: "3524",
			Name:   "Lamongan",
		},
		{
			Id:     "352423",
			CityId: "3524",
			Name:   "Tikung",
		},
		{
			Id:     "352424",
			CityId: "3524",
			Name:   "Karangbinangun",
		},
		{
			Id:     "352425",
			CityId: "3524",
			Name:   "Deket",
		},
		{
			Id:     "352426",
			CityId: "3524",
			Name:   "Glagah",
		},
		{
			Id:     "352427",
			CityId: "3524",
			Name:   "Sarirejo",
		},
		{
			Id:     "352501",
			CityId: "3525",
			Name:   "Dukun",
		},
		{
			Id:     "352502",
			CityId: "3525",
			Name:   "Balongpanggang",
		},
		{
			Id:     "352503",
			CityId: "3525",
			Name:   "Panceng",
		},
		{
			Id:     "352504",
			CityId: "3525",
			Name:   "Benjeng",
		},
		{
			Id:     "352505",
			CityId: "3525",
			Name:   "Duduksampeyan",
		},
		{
			Id:     "352506",
			CityId: "3525",
			Name:   "Wringinanom",
		},
		{
			Id:     "352507",
			CityId: "3525",
			Name:   "Ujungpangkah",
		},
		{
			Id:     "352508",
			CityId: "3525",
			Name:   "Kedamean",
		},
		{
			Id:     "352509",
			CityId: "3525",
			Name:   "Sidayu",
		},
		{
			Id:     "352510",
			CityId: "3525",
			Name:   "Manyar",
		},
		{
			Id:     "352511",
			CityId: "3525",
			Name:   "Cerme",
		},
		{
			Id:     "352512",
			CityId: "3525",
			Name:   "Bungah",
		},
		{
			Id:     "352513",
			CityId: "3525",
			Name:   "Menganti",
		},
		{
			Id:     "352514",
			CityId: "3525",
			Name:   "Kebomas",
		},
		{
			Id:     "352515",
			CityId: "3525",
			Name:   "Driyorejo",
		},
		{
			Id:     "352516",
			CityId: "3525",
			Name:   "Gresik",
		},
		{
			Id:     "352517",
			CityId: "3525",
			Name:   "Sangkapura",
		},
		{
			Id:     "352518",
			CityId: "3525",
			Name:   "Tambak",
		},
		{
			Id:     "352601",
			CityId: "3526",
			Name:   "Bangkalan",
		},
		{
			Id:     "352602",
			CityId: "3526",
			Name:   "Socah",
		},
		{
			Id:     "352603",
			CityId: "3526",
			Name:   "Burneh",
		},
		{
			Id:     "352604",
			CityId: "3526",
			Name:   "Kamal",
		},
		{
			Id:     "352605",
			CityId: "3526",
			Name:   "Arosbaya",
		},
		{
			Id:     "352606",
			CityId: "3526",
			Name:   "Geger",
		},
		{
			Id:     "352607",
			CityId: "3526",
			Name:   "Klampis",
		},
		{
			Id:     "352608",
			CityId: "3526",
			Name:   "Sepulu",
		},
		{
			Id:     "352609",
			CityId: "3526",
			Name:   "Tanjung Bumi",
		},
		{
			Id:     "352610",
			CityId: "3526",
			Name:   "Kokop",
		},
		{
			Id:     "352611",
			CityId: "3526",
			Name:   "Kwanyar",
		},
		{
			Id:     "352612",
			CityId: "3526",
			Name:   "Labang",
		},
		{
			Id:     "352613",
			CityId: "3526",
			Name:   "Tanah Merah",
		},
		{
			Id:     "352614",
			CityId: "3526",
			Name:   "Tragah",
		},
		{
			Id:     "352615",
			CityId: "3526",
			Name:   "Blega",
		},
		{
			Id:     "352616",
			CityId: "3526",
			Name:   "Modung",
		},
		{
			Id:     "352617",
			CityId: "3526",
			Name:   "Konang",
		},
		{
			Id:     "352618",
			CityId: "3526",
			Name:   "Galis",
		},
		{
			Id:     "352701",
			CityId: "3527",
			Name:   "Sreseh",
		},
		{
			Id:     "352702",
			CityId: "3527",
			Name:   "Torjun",
		},
		{
			Id:     "352703",
			CityId: "3527",
			Name:   "Sampang",
		},
		{
			Id:     "352704",
			CityId: "3527",
			Name:   "Camplong",
		},
		{
			Id:     "352705",
			CityId: "3527",
			Name:   "Omben",
		},
		{
			Id:     "352706",
			CityId: "3527",
			Name:   "Kedungdung",
		},
		{
			Id:     "352707",
			CityId: "3527",
			Name:   "Jrengik",
		},
		{
			Id:     "352708",
			CityId: "3527",
			Name:   "Tambelangan",
		},
		{
			Id:     "352709",
			CityId: "3527",
			Name:   "Banyuates",
		},
		{
			Id:     "352710",
			CityId: "3527",
			Name:   "Robatal",
		},
		{
			Id:     "352711",
			CityId: "3527",
			Name:   "Sokobanah",
		},
		{
			Id:     "352712",
			CityId: "3527",
			Name:   "Ketapang",
		},
		{
			Id:     "352713",
			CityId: "3527",
			Name:   "Pangarengan",
		},
		{
			Id:     "352714",
			CityId: "3527",
			Name:   "Karangpenang",
		},
		{
			Id:     "352801",
			CityId: "3528",
			Name:   "Tlanakan",
		},
		{
			Id:     "352802",
			CityId: "3528",
			Name:   "Pademawu",
		},
		{
			Id:     "352803",
			CityId: "3528",
			Name:   "Galis",
		},
		{
			Id:     "352804",
			CityId: "3528",
			Name:   "Pamekasan",
		},
		{
			Id:     "352805",
			CityId: "3528",
			Name:   "Proppo",
		},
		{
			Id:     "352806",
			CityId: "3528",
			Name:   "Palenggaan",
		},
		{
			Id:     "352807",
			CityId: "3528",
			Name:   "Pegantenan",
		},
		{
			Id:     "352808",
			CityId: "3528",
			Name:   "Larangan",
		},
		{
			Id:     "352809",
			CityId: "3528",
			Name:   "Pakong",
		},
		{
			Id:     "352810",
			CityId: "3528",
			Name:   "Waru",
		},
		{
			Id:     "352811",
			CityId: "3528",
			Name:   "Batumarmar",
		},
		{
			Id:     "352812",
			CityId: "3528",
			Name:   "Kadur",
		},
		{
			Id:     "352813",
			CityId: "3528",
			Name:   "Pasean",
		},
		{
			Id:     "352901",
			CityId: "3529",
			Name:   "Kota Sumenep",
		},
		{
			Id:     "352902",
			CityId: "3529",
			Name:   "Kalianget",
		},
		{
			Id:     "352903",
			CityId: "3529",
			Name:   "Manding",
		},
		{
			Id:     "352904",
			CityId: "3529",
			Name:   "Talango",
		},
		{
			Id:     "352905",
			CityId: "3529",
			Name:   "Bluto",
		},
		{
			Id:     "352906",
			CityId: "3529",
			Name:   "Saronggi",
		},
		{
			Id:     "352907",
			CityId: "3529",
			Name:   "Lenteng",
		},
		{
			Id:     "352908",
			CityId: "3529",
			Name:   "Giliginting",
		},
		{
			Id:     "352909",
			CityId: "3529",
			Name:   "Guluk-Guluk",
		},
		{
			Id:     "352910",
			CityId: "3529",
			Name:   "Ganding",
		},
		{
			Id:     "352911",
			CityId: "3529",
			Name:   "Pragaan",
		},
		{
			Id:     "352912",
			CityId: "3529",
			Name:   "Ambunten",
		},
		{
			Id:     "352913",
			CityId: "3529",
			Name:   "Pasongsongan",
		},
		{
			Id:     "352914",
			CityId: "3529",
			Name:   "Dasuk",
		},
		{
			Id:     "352915",
			CityId: "3529",
			Name:   "Rubaru",
		},
		{
			Id:     "352916",
			CityId: "3529",
			Name:   "Batang Batang",
		},
		{
			Id:     "352917",
			CityId: "3529",
			Name:   "Batuputih",
		},
		{
			Id:     "352918",
			CityId: "3529",
			Name:   "Dungkek",
		},
		{
			Id:     "352919",
			CityId: "3529",
			Name:   "Gapura",
		},
		{
			Id:     "352920",
			CityId: "3529",
			Name:   "Gayam",
		},
		{
			Id:     "352921",
			CityId: "3529",
			Name:   "Nonggunong",
		},
		{
			Id:     "352922",
			CityId: "3529",
			Name:   "Raas",
		},
		{
			Id:     "352923",
			CityId: "3529",
			Name:   "Masalembu",
		},
		{
			Id:     "352924",
			CityId: "3529",
			Name:   "Arjasa",
		},
		{
			Id:     "352925",
			CityId: "3529",
			Name:   "Sapeken",
		},
		{
			Id:     "352926",
			CityId: "3529",
			Name:   "Batuan",
		},
		{
			Id:     "352927",
			CityId: "3529",
			Name:   "Kangayan",
		},
		{
			Id:     "357101",
			CityId: "3571",
			Name:   "Mojoroto",
		},
		{
			Id:     "357102",
			CityId: "3571",
			Name:   "Kota",
		},
		{
			Id:     "357103",
			CityId: "3571",
			Name:   "Pesantren",
		},
		{
			Id:     "357201",
			CityId: "3572",
			Name:   "Kepanjenkidul",
		},
		{
			Id:     "357202",
			CityId: "3572",
			Name:   "Sukorejo",
		},
		{
			Id:     "357203",
			CityId: "3572",
			Name:   "Sananwetan",
		},
		{
			Id:     "357301",
			CityId: "3573",
			Name:   "Blimbing",
		},
		{
			Id:     "357302",
			CityId: "3573",
			Name:   "Klojen",
		},
		{
			Id:     "357303",
			CityId: "3573",
			Name:   "Kedungkandang",
		},
		{
			Id:     "357304",
			CityId: "3573",
			Name:   "Sukun",
		},
		{
			Id:     "357305",
			CityId: "3573",
			Name:   "Lowokwaru",
		},
		{
			Id:     "357401",
			CityId: "3574",
			Name:   "Kademangan",
		},
		{
			Id:     "357402",
			CityId: "3574",
			Name:   "Wonoasih",
		},
		{
			Id:     "357403",
			CityId: "3574",
			Name:   "Mayangan",
		},
		{
			Id:     "357404",
			CityId: "3574",
			Name:   "Kanigaran",
		},
		{
			Id:     "357405",
			CityId: "3574",
			Name:   "Kedopok",
		},
		{
			Id:     "357501",
			CityId: "3575",
			Name:   "Gadingrejo",
		},
		{
			Id:     "357502",
			CityId: "3575",
			Name:   "Purworejo",
		},
		{
			Id:     "357503",
			CityId: "3575",
			Name:   "Bugul Kidul",
		},
		{
			Id:     "357504",
			CityId: "3575",
			Name:   "Panggungrejo",
		},
		{
			Id:     "357601",
			CityId: "3576",
			Name:   "Prajuritkulon",
		},
		{
			Id:     "357602",
			CityId: "3576",
			Name:   "Magersari",
		},
		{
			Id:     "357603",
			CityId: "3576",
			Name:   "Kranggan",
		},
		{
			Id:     "357701",
			CityId: "3577",
			Name:   "Kartoharjo",
		},
		{
			Id:     "357702",
			CityId: "3577",
			Name:   "Manguharjo",
		},
		{
			Id:     "357703",
			CityId: "3577",
			Name:   "Taman",
		},
		{
			Id:     "357801",
			CityId: "3578",
			Name:   "Karang Pilang",
		},
		{
			Id:     "357802",
			CityId: "3578",
			Name:   "Wonocolo",
		},
		{
			Id:     "357803",
			CityId: "3578",
			Name:   "Rungkut",
		},
		{
			Id:     "357804",
			CityId: "3578",
			Name:   "Wonokromo",
		},
		{
			Id:     "357805",
			CityId: "3578",
			Name:   "Tegalsari",
		},
		{
			Id:     "357806",
			CityId: "3578",
			Name:   "Sawahan",
		},
		{
			Id:     "357807",
			CityId: "3578",
			Name:   "Genteng",
		},
		{
			Id:     "357808",
			CityId: "3578",
			Name:   "Gubeng",
		},
		{
			Id:     "357809",
			CityId: "3578",
			Name:   "Sukolilo",
		},
		{
			Id:     "357810",
			CityId: "3578",
			Name:   "Tambaksari",
		},
		{
			Id:     "357811",
			CityId: "3578",
			Name:   "Simokerto",
		},
		{
			Id:     "357812",
			CityId: "3578",
			Name:   "Pabean Cantian",
		},
		{
			Id:     "357813",
			CityId: "3578",
			Name:   "Bubutan",
		},
		{
			Id:     "357814",
			CityId: "3578",
			Name:   "Tandes",
		},
		{
			Id:     "357815",
			CityId: "3578",
			Name:   "Krembangan",
		},
		{
			Id:     "357816",
			CityId: "3578",
			Name:   "Semampir",
		},
		{
			Id:     "357817",
			CityId: "3578",
			Name:   "Kenjeran",
		},
		{
			Id:     "357818",
			CityId: "3578",
			Name:   "Lakarsantri",
		},
		{
			Id:     "357819",
			CityId: "3578",
			Name:   "Benowo",
		},
		{
			Id:     "357820",
			CityId: "3578",
			Name:   "Wiyung",
		},
		{
			Id:     "357821",
			CityId: "3578",
			Name:   "Dukuh Pakis",
		},
		{
			Id:     "357822",
			CityId: "3578",
			Name:   "Gayungan",
		},
		{
			Id:     "357823",
			CityId: "3578",
			Name:   "Jambangan",
		},
		{
			Id:     "357824",
			CityId: "3578",
			Name:   "Tenggilis Mejoyo",
		},
		{
			Id:     "357825",
			CityId: "3578",
			Name:   "Gunung Anyar",
		},
		{
			Id:     "357826",
			CityId: "3578",
			Name:   "Mulyorejo",
		},
		{
			Id:     "357827",
			CityId: "3578",
			Name:   "Sukomanunggal",
		},
		{
			Id:     "357828",
			CityId: "3578",
			Name:   "Asem Rowo",
		},
		{
			Id:     "357829",
			CityId: "3578",
			Name:   "Bulak",
		},
		{
			Id:     "357830",
			CityId: "3578",
			Name:   "Pakal",
		},
		{
			Id:     "357831",
			CityId: "3578",
			Name:   "Sambikerep",
		},
		{
			Id:     "357901",
			CityId: "3579",
			Name:   "Batu",
		},
		{
			Id:     "357902",
			CityId: "3579",
			Name:   "Bumiaji",
		},
		{
			Id:     "357903",
			CityId: "3579",
			Name:   "Junrejo",
		},
		{
			Id:     "360101",
			CityId: "3601",
			Name:   "Sumur",
		},
		{
			Id:     "360102",
			CityId: "3601",
			Name:   "Cimanggu",
		},
		{
			Id:     "360103",
			CityId: "3601",
			Name:   "Cibaliung",
		},
		{
			Id:     "360104",
			CityId: "3601",
			Name:   "Cikeusik",
		},
		{
			Id:     "360105",
			CityId: "3601",
			Name:   "Cigeulis",
		},
		{
			Id:     "360106",
			CityId: "3601",
			Name:   "Panimbang",
		},
		{
			Id:     "360107",
			CityId: "3601",
			Name:   "Angsana",
		},
		{
			Id:     "360108",
			CityId: "3601",
			Name:   "Munjul",
		},
		{
			Id:     "360109",
			CityId: "3601",
			Name:   "Pagelaran",
		},
		{
			Id:     "360110",
			CityId: "3601",
			Name:   "Bojong",
		},
		{
			Id:     "360111",
			CityId: "3601",
			Name:   "Picung",
		},
		{
			Id:     "360112",
			CityId: "3601",
			Name:   "Labuan",
		},
		{
			Id:     "360113",
			CityId: "3601",
			Name:   "Menes",
		},
		{
			Id:     "360114",
			CityId: "3601",
			Name:   "Saketi",
		},
		{
			Id:     "360115",
			CityId: "3601",
			Name:   "Cipeucang",
		},
		{
			Id:     "360116",
			CityId: "3601",
			Name:   "Jiput",
		},
		{
			Id:     "360117",
			CityId: "3601",
			Name:   "Mandalawangi",
		},
		{
			Id:     "360118",
			CityId: "3601",
			Name:   "Cimanuk",
		},
		{
			Id:     "360119",
			CityId: "3601",
			Name:   "Kaduhejo",
		},
		{
			Id:     "360120",
			CityId: "3601",
			Name:   "Banjar",
		},
		{
			Id:     "360121",
			CityId: "3601",
			Name:   "Pandeglang",
		},
		{
			Id:     "360122",
			CityId: "3601",
			Name:   "Cadasari",
		},
		{
			Id:     "360123",
			CityId: "3601",
			Name:   "Cisata",
		},
		{
			Id:     "360124",
			CityId: "3601",
			Name:   "Patia",
		},
		{
			Id:     "360125",
			CityId: "3601",
			Name:   "Karang Tanjung",
		},
		{
			Id:     "360126",
			CityId: "3601",
			Name:   "Cikedal",
		},
		{
			Id:     "360127",
			CityId: "3601",
			Name:   "Cibitung",
		},
		{
			Id:     "360128",
			CityId: "3601",
			Name:   "Carita",
		},
		{
			Id:     "360129",
			CityId: "3601",
			Name:   "Sukaresmi",
		},
		{
			Id:     "360130",
			CityId: "3601",
			Name:   "Mekarjaya",
		},
		{
			Id:     "360131",
			CityId: "3601",
			Name:   "Sindangresmi",
		},
		{
			Id:     "360132",
			CityId: "3601",
			Name:   "Pulosari",
		},
		{
			Id:     "360133",
			CityId: "3601",
			Name:   "Koroncong",
		},
		{
			Id:     "360134",
			CityId: "3601",
			Name:   "Majasari",
		},
		{
			Id:     "360135",
			CityId: "3601",
			Name:   "Sobang",
		},
		{
			Id:     "360201",
			CityId: "3602",
			Name:   "Malingping",
		},
		{
			Id:     "360202",
			CityId: "3602",
			Name:   "Panggarangan",
		},
		{
			Id:     "360203",
			CityId: "3602",
			Name:   "Bayah",
		},
		{
			Id:     "360204",
			CityId: "3602",
			Name:   "Cipanas",
		},
		{
			Id:     "360205",
			CityId: "3602",
			Name:   "Muncang",
		},
		{
			Id:     "360206",
			CityId: "3602",
			Name:   "Leuwidamar",
		},
		{
			Id:     "360207",
			CityId: "3602",
			Name:   "Bojongmanik",
		},
		{
			Id:     "360208",
			CityId: "3602",
			Name:   "Gunungkencana",
		},
		{
			Id:     "360209",
			CityId: "3602",
			Name:   "Banjarsari",
		},
		{
			Id:     "360210",
			CityId: "3602",
			Name:   "Cileles",
		},
		{
			Id:     "360211",
			CityId: "3602",
			Name:   "Cimarga",
		},
		{
			Id:     "360212",
			CityId: "3602",
			Name:   "Sajira",
		},
		{
			Id:     "360213",
			CityId: "3602",
			Name:   "Maja",
		},
		{
			Id:     "360214",
			CityId: "3602",
			Name:   "Rangkasbitung",
		},
		{
			Id:     "360215",
			CityId: "3602",
			Name:   "Warunggunung",
		},
		{
			Id:     "360216",
			CityId: "3602",
			Name:   "Cijaku",
		},
		{
			Id:     "360217",
			CityId: "3602",
			Name:   "Cikulur",
		},
		{
			Id:     "360218",
			CityId: "3602",
			Name:   "Cibadak",
		},
		{
			Id:     "360219",
			CityId: "3602",
			Name:   "Cibeber",
		},
		{
			Id:     "360220",
			CityId: "3602",
			Name:   "Cilograng",
		},
		{
			Id:     "360221",
			CityId: "3602",
			Name:   "Wanasalam",
		},
		{
			Id:     "360222",
			CityId: "3602",
			Name:   "Sobang",
		},
		{
			Id:     "360223",
			CityId: "3602",
			Name:   "Curug bitung",
		},
		{
			Id:     "360224",
			CityId: "3602",
			Name:   "Kalanganyar",
		},
		{
			Id:     "360225",
			CityId: "3602",
			Name:   "Lebakgedong",
		},
		{
			Id:     "360226",
			CityId: "3602",
			Name:   "Cihara",
		},
		{
			Id:     "360227",
			CityId: "3602",
			Name:   "Cirinten",
		},
		{
			Id:     "360228",
			CityId: "3602",
			Name:   "Cigemlong",
		},
		{
			Id:     "360301",
			CityId: "3603",
			Name:   "Balaraja",
		},
		{
			Id:     "360302",
			CityId: "3603",
			Name:   "Jayanti",
		},
		{
			Id:     "360303",
			CityId: "3603",
			Name:   "Tigaraksa",
		},
		{
			Id:     "360304",
			CityId: "3603",
			Name:   "Jambe",
		},
		{
			Id:     "360305",
			CityId: "3603",
			Name:   "Cisoka",
		},
		{
			Id:     "360306",
			CityId: "3603",
			Name:   "Kresek",
		},
		{
			Id:     "360307",
			CityId: "3603",
			Name:   "Kronjo",
		},
		{
			Id:     "360308",
			CityId: "3603",
			Name:   "Mauk",
		},
		{
			Id:     "360309",
			CityId: "3603",
			Name:   "Kemiri",
		},
		{
			Id:     "360310",
			CityId: "3603",
			Name:   "Sukadiri",
		},
		{
			Id:     "360311",
			CityId: "3603",
			Name:   "Rajeg",
		},
		{
			Id:     "360312",
			CityId: "3603",
			Name:   "Pasar Kemis",
		},
		{
			Id:     "360313",
			CityId: "3603",
			Name:   "Teluknaga",
		},
		{
			Id:     "360314",
			CityId: "3603",
			Name:   "Kosambi",
		},
		{
			Id:     "360315",
			CityId: "3603",
			Name:   "Pakuhaji",
		},
		{
			Id:     "360316",
			CityId: "3603",
			Name:   "Sepatan",
		},
		{
			Id:     "360317",
			CityId: "3603",
			Name:   "Curug",
		},
		{
			Id:     "360318",
			CityId: "3603",
			Name:   "Cikupa",
		},
		{
			Id:     "360319",
			CityId: "3603",
			Name:   "Panongan",
		},
		{
			Id:     "360320",
			CityId: "3603",
			Name:   "Legok",
		},
		{
			Id:     "360322",
			CityId: "3603",
			Name:   "Pagedangan",
		},
		{
			Id:     "360323",
			CityId: "3603",
			Name:   "Cisauk",
		},
		{
			Id:     "360327",
			CityId: "3603",
			Name:   "Sukamulya",
		},
		{
			Id:     "360328",
			CityId: "3603",
			Name:   "Kelapa Dua",
		},
		{
			Id:     "360329",
			CityId: "3603",
			Name:   "Sindang Jaya",
		},
		{
			Id:     "360330",
			CityId: "3603",
			Name:   "Sepatan Timur",
		},
		{
			Id:     "360331",
			CityId: "3603",
			Name:   "Solear",
		},
		{
			Id:     "360332",
			CityId: "3603",
			Name:   "Gunung Kaler",
		},
		{
			Id:     "360333",
			CityId: "3603",
			Name:   "Mekar Baru",
		},
		{
			Id:     "360405",
			CityId: "3604",
			Name:   "Kramatwatu",
		},
		{
			Id:     "360406",
			CityId: "3604",
			Name:   "Waringinkurung",
		},
		{
			Id:     "360407",
			CityId: "3604",
			Name:   "Bojonegara",
		},
		{
			Id:     "360408",
			CityId: "3604",
			Name:   "Pulo Ampel",
		},
		{
			Id:     "360409",
			CityId: "3604",
			Name:   "Ciruas",
		},
		{
			Id:     "360411",
			CityId: "3604",
			Name:   "Kragilan",
		},
		{
			Id:     "360412",
			CityId: "3604",
			Name:   "Pontang",
		},
		{
			Id:     "360413",
			CityId: "3604",
			Name:   "Tirtayasa",
		},
		{
			Id:     "360414",
			CityId: "3604",
			Name:   "Tanara",
		},
		{
			Id:     "360415",
			CityId: "3604",
			Name:   "Cikande",
		},
		{
			Id:     "360416",
			CityId: "3604",
			Name:   "Kibin",
		},
		{
			Id:     "360417",
			CityId: "3604",
			Name:   "Carenang",
		},
		{
			Id:     "360418",
			CityId: "3604",
			Name:   "Binuang",
		},
		{
			Id:     "360419",
			CityId: "3604",
			Name:   "Petir",
		},
		{
			Id:     "360420",
			CityId: "3604",
			Name:   "Tunjung Teja",
		},
		{
			Id:     "360422",
			CityId: "3604",
			Name:   "Baros",
		},
		{
			Id:     "360423",
			CityId: "3604",
			Name:   "Cikeusal",
		},
		{
			Id:     "360424",
			CityId: "3604",
			Name:   "Pamarayan",
		},
		{
			Id:     "360425",
			CityId: "3604",
			Name:   "Kopo",
		},
		{
			Id:     "360426",
			CityId: "3604",
			Name:   "Jawilan",
		},
		{
			Id:     "360427",
			CityId: "3604",
			Name:   "Ciomas",
		},
		{
			Id:     "360428",
			CityId: "3604",
			Name:   "Pabuaran",
		},
		{
			Id:     "360429",
			CityId: "3604",
			Name:   "Padarincang",
		},
		{
			Id:     "360430",
			CityId: "3604",
			Name:   "Anyar",
		},
		{
			Id:     "360431",
			CityId: "3604",
			Name:   "Cinangka",
		},
		{
			Id:     "360432",
			CityId: "3604",
			Name:   "Mancak",
		},
		{
			Id:     "360433",
			CityId: "3604",
			Name:   "Gunungsari",
		},
		{
			Id:     "360434",
			CityId: "3604",
			Name:   "Bandung",
		},
		{
			Id:     "360435",
			CityId: "3604",
			Name:   "Lebak Wangi",
		},
		{
			Id:     "367101",
			CityId: "3671",
			Name:   "Tangerang",
		},
		{
			Id:     "367102",
			CityId: "3671",
			Name:   "Jatiuwung",
		},
		{
			Id:     "367103",
			CityId: "3671",
			Name:   "Batuceper",
		},
		{
			Id:     "367104",
			CityId: "3671",
			Name:   "Benda",
		},
		{
			Id:     "367105",
			CityId: "3671",
			Name:   "Cipondoh",
		},
		{
			Id:     "367106",
			CityId: "3671",
			Name:   "Ciledug",
		},
		{
			Id:     "367107",
			CityId: "3671",
			Name:   "Karawaci",
		},
		{
			Id:     "367108",
			CityId: "3671",
			Name:   "Periuk",
		},
		{
			Id:     "367109",
			CityId: "3671",
			Name:   "Cibodas",
		},
		{
			Id:     "367110",
			CityId: "3671",
			Name:   "Neglasari",
		},
		{
			Id:     "367111",
			CityId: "3671",
			Name:   "Pinang",
		},
		{
			Id:     "367112",
			CityId: "3671",
			Name:   "Karang Tengah",
		},
		{
			Id:     "367113",
			CityId: "3671",
			Name:   "Larangan",
		},
		{
			Id:     "367201",
			CityId: "3672",
			Name:   "Cibeber",
		},
		{
			Id:     "367202",
			CityId: "3672",
			Name:   "Cilegon",
		},
		{
			Id:     "367203",
			CityId: "3672",
			Name:   "Pulomerak",
		},
		{
			Id:     "367204",
			CityId: "3672",
			Name:   "Ciwandan",
		},
		{
			Id:     "367205",
			CityId: "3672",
			Name:   "Jombang",
		},
		{
			Id:     "367206",
			CityId: "3672",
			Name:   "Gerogol",
		},
		{
			Id:     "367207",
			CityId: "3672",
			Name:   "Purwakarta",
		},
		{
			Id:     "367208",
			CityId: "3672",
			Name:   "Citangkil",
		},
		{
			Id:     "367301",
			CityId: "3673",
			Name:   "Serang",
		},
		{
			Id:     "367302",
			CityId: "3673",
			Name:   "Kasemen",
		},
		{
			Id:     "367303",
			CityId: "3673",
			Name:   "Walantaka",
		},
		{
			Id:     "367304",
			CityId: "3673",
			Name:   "Curug",
		},
		{
			Id:     "367305",
			CityId: "3673",
			Name:   "Cipocok Jaya",
		},
		{
			Id:     "367306",
			CityId: "3673",
			Name:   "Taktakan",
		},
		{
			Id:     "367401",
			CityId: "3674",
			Name:   "Serpong",
		},
		{
			Id:     "367402",
			CityId: "3674",
			Name:   "Serpong Utara",
		},
		{
			Id:     "367403",
			CityId: "3674",
			Name:   "Pondok Aren",
		},
		{
			Id:     "367404",
			CityId: "3674",
			Name:   "Ciputat",
		},
		{
			Id:     "367405",
			CityId: "3674",
			Name:   "Ciputat Timur",
		},
		{
			Id:     "367406",
			CityId: "3674",
			Name:   "Pamulang",
		},
		{
			Id:     "367407",
			CityId: "3674",
			Name:   "Setu",
		},
		{
			Id:     "510101",
			CityId: "5101",
			Name:   "Negara",
		},
		{
			Id:     "510102",
			CityId: "5101",
			Name:   "Mendoyo",
		},
		{
			Id:     "510103",
			CityId: "5101",
			Name:   "Pekutatan",
		},
		{
			Id:     "510104",
			CityId: "5101",
			Name:   "Melaya",
		},
		{
			Id:     "510105",
			CityId: "5101",
			Name:   "Jembrana",
		},
		{
			Id:     "510201",
			CityId: "5102",
			Name:   "Selemadeg",
		},
		{
			Id:     "510202",
			CityId: "5102",
			Name:   "Selemadeg Timur",
		},
		{
			Id:     "510203",
			CityId: "5102",
			Name:   "Selemadeg Barat",
		},
		{
			Id:     "510204",
			CityId: "5102",
			Name:   "Kerambitan",
		},
		{
			Id:     "510205",
			CityId: "5102",
			Name:   "Tabanan",
		},
		{
			Id:     "510206",
			CityId: "5102",
			Name:   "Kediri",
		},
		{
			Id:     "510207",
			CityId: "5102",
			Name:   "Marga",
		},
		{
			Id:     "510208",
			CityId: "5102",
			Name:   "Penebel",
		},
		{
			Id:     "510209",
			CityId: "5102",
			Name:   "Baturiti",
		},
		{
			Id:     "510210",
			CityId: "5102",
			Name:   "Pupuan",
		},
		{
			Id:     "510301",
			CityId: "5103",
			Name:   "Kuta",
		},
		{
			Id:     "510302",
			CityId: "5103",
			Name:   "Mengwi",
		},
		{
			Id:     "510303",
			CityId: "5103",
			Name:   "Abiansemal",
		},
		{
			Id:     "510304",
			CityId: "5103",
			Name:   "Petang",
		},
		{
			Id:     "510305",
			CityId: "5103",
			Name:   "Kuta Selatan",
		},
		{
			Id:     "510306",
			CityId: "5103",
			Name:   "Kuta Utara",
		},
		{
			Id:     "510401",
			CityId: "5104",
			Name:   "Sukawati",
		},
		{
			Id:     "510402",
			CityId: "5104",
			Name:   "Blahbatuh",
		},
		{
			Id:     "510403",
			CityId: "5104",
			Name:   "Gianyar",
		},
		{
			Id:     "510404",
			CityId: "5104",
			Name:   "Tampaksiring",
		},
		{
			Id:     "510405",
			CityId: "5104",
			Name:   "Ubud",
		},
		{
			Id:     "510406",
			CityId: "5104",
			Name:   "Tegallalang",
		},
		{
			Id:     "510407",
			CityId: "5104",
			Name:   "Payangan",
		},
		{
			Id:     "510501",
			CityId: "5105",
			Name:   "Nusa Penida",
		},
		{
			Id:     "510502",
			CityId: "5105",
			Name:   "Banjarangkan",
		},
		{
			Id:     "510503",
			CityId: "5105",
			Name:   "Klungkung",
		},
		{
			Id:     "510504",
			CityId: "5105",
			Name:   "Dawan",
		},
		{
			Id:     "510601",
			CityId: "5106",
			Name:   "Susut",
		},
		{
			Id:     "510602",
			CityId: "5106",
			Name:   "Bangli",
		},
		{
			Id:     "510603",
			CityId: "5106",
			Name:   "Tembuku",
		},
		{
			Id:     "510604",
			CityId: "5106",
			Name:   "Kintamani",
		},
		{
			Id:     "510701",
			CityId: "5107",
			Name:   "Rendang",
		},
		{
			Id:     "510702",
			CityId: "5107",
			Name:   "Sidemen",
		},
		{
			Id:     "510703",
			CityId: "5107",
			Name:   "Manggis",
		},
		{
			Id:     "510704",
			CityId: "5107",
			Name:   "Karangasem",
		},
		{
			Id:     "510705",
			CityId: "5107",
			Name:   "Abang",
		},
		{
			Id:     "510706",
			CityId: "5107",
			Name:   "Bebandem",
		},
		{
			Id:     "510707",
			CityId: "5107",
			Name:   "Selat",
		},
		{
			Id:     "510708",
			CityId: "5107",
			Name:   "Kubu",
		},
		{
			Id:     "510801",
			CityId: "5108",
			Name:   "Gerokgak",
		},
		{
			Id:     "510802",
			CityId: "5108",
			Name:   "Seririt",
		},
		{
			Id:     "510803",
			CityId: "5108",
			Name:   "Busungbiu",
		},
		{
			Id:     "510804",
			CityId: "5108",
			Name:   "Banjar",
		},
		{
			Id:     "510805",
			CityId: "5108",
			Name:   "Sukasada",
		},
		{
			Id:     "510806",
			CityId: "5108",
			Name:   "Buleleng",
		},
		{
			Id:     "510807",
			CityId: "5108",
			Name:   "Sawan",
		},
		{
			Id:     "510808",
			CityId: "5108",
			Name:   "Kubutambahan",
		},
		{
			Id:     "510809",
			CityId: "5108",
			Name:   "Tejakula",
		},
		{
			Id:     "517101",
			CityId: "5171",
			Name:   "Denpasar Selatan",
		},
		{
			Id:     "517102",
			CityId: "5171",
			Name:   "Denpasar Timur",
		},
		{
			Id:     "517103",
			CityId: "5171",
			Name:   "Denpasar Barat",
		},
		{
			Id:     "517104",
			CityId: "5171",
			Name:   "Denpasar Utara",
		},
		{
			Id:     "520101",
			CityId: "5201",
			Name:   "Gerung",
		},
		{
			Id:     "520102",
			CityId: "5201",
			Name:   "Kediri",
		},
		{
			Id:     "520103",
			CityId: "5201",
			Name:   "Narmada",
		},
		{
			Id:     "520107",
			CityId: "5201",
			Name:   "Sekotong",
		},
		{
			Id:     "520108",
			CityId: "5201",
			Name:   "Labuapi",
		},
		{
			Id:     "520109",
			CityId: "5201",
			Name:   "Gunungsari",
		},
		{
			Id:     "520112",
			CityId: "5201",
			Name:   "Lingsar",
		},
		{
			Id:     "520113",
			CityId: "5201",
			Name:   "Lembar",
		},
		{
			Id:     "520114",
			CityId: "5201",
			Name:   "Batu Layar",
		},
		{
			Id:     "520115",
			CityId: "5201",
			Name:   "Kuripan",
		},
		{
			Id:     "520201",
			CityId: "5202",
			Name:   "Praya",
		},
		{
			Id:     "520202",
			CityId: "5202",
			Name:   "Jonggat",
		},
		{
			Id:     "520203",
			CityId: "5202",
			Name:   "Batukliang",
		},
		{
			Id:     "520204",
			CityId: "5202",
			Name:   "Pujut",
		},
		{
			Id:     "520205",
			CityId: "5202",
			Name:   "Praya Barat",
		},
		{
			Id:     "520206",
			CityId: "5202",
			Name:   "Praya Timur",
		},
		{
			Id:     "520207",
			CityId: "5202",
			Name:   "Janapria",
		},
		{
			Id:     "520208",
			CityId: "5202",
			Name:   "Pringgarata",
		},
		{
			Id:     "520209",
			CityId: "5202",
			Name:   "Kopang",
		},
		{
			Id:     "520210",
			CityId: "5202",
			Name:   "Praya Tengah",
		},
		{
			Id:     "520211",
			CityId: "5202",
			Name:   "Praya Barat Daya",
		},
		{
			Id:     "520212",
			CityId: "5202",
			Name:   "Batukliang Utara",
		},
		{
			Id:     "520301",
			CityId: "5203",
			Name:   "Keruak",
		},
		{
			Id:     "520302",
			CityId: "5203",
			Name:   "Sakra",
		},
		{
			Id:     "520303",
			CityId: "5203",
			Name:   "Terara",
		},
		{
			Id:     "520304",
			CityId: "5203",
			Name:   "Sikur",
		},
		{
			Id:     "520305",
			CityId: "5203",
			Name:   "Masbagik",
		},
		{
			Id:     "520306",
			CityId: "5203",
			Name:   "Sukamulia",
		},
		{
			Id:     "520307",
			CityId: "5203",
			Name:   "Selong",
		},
		{
			Id:     "520308",
			CityId: "5203",
			Name:   "Pringgabaya",
		},
		{
			Id:     "520309",
			CityId: "5203",
			Name:   "Aikmel",
		},
		{
			Id:     "520310",
			CityId: "5203",
			Name:   "Sambelia",
		},
		{
			Id:     "520311",
			CityId: "5203",
			Name:   "Montong Gading",
		},
		{
			Id:     "520312",
			CityId: "5203",
			Name:   "Pringgasela",
		},
		{
			Id:     "520313",
			CityId: "5203",
			Name:   "Suralaga",
		},
		{
			Id:     "520314",
			CityId: "5203",
			Name:   "Wanasaba",
		},
		{
			Id:     "520315",
			CityId: "5203",
			Name:   "Sembalun",
		},
		{
			Id:     "520316",
			CityId: "5203",
			Name:   "Suwela",
		},
		{
			Id:     "520317",
			CityId: "5203",
			Name:   "Labuhan Haji",
		},
		{
			Id:     "520318",
			CityId: "5203",
			Name:   "Sakra Timur",
		},
		{
			Id:     "520319",
			CityId: "5203",
			Name:   "Sakra Barat",
		},
		{
			Id:     "520320",
			CityId: "5203",
			Name:   "Jerowaru",
		},
		{
			Id:     "520321",
			CityId: "5203",
			Name:   "Lenek",
		},
		{
			Id:     "520402",
			CityId: "5204",
			Name:   "Lunyuk",
		},
		{
			Id:     "520405",
			CityId: "5204",
			Name:   "Alas",
		},
		{
			Id:     "520406",
			CityId: "5204",
			Name:   "Utan",
		},
		{
			Id:     "520407",
			CityId: "5204",
			Name:   "Batu Lanteh",
		},
		{
			Id:     "520408",
			CityId: "5204",
			Name:   "Sumbawa",
		},
		{
			Id:     "520409",
			CityId: "5204",
			Name:   "Moyo Hilir",
		},
		{
			Id:     "520410",
			CityId: "5204",
			Name:   "Moyo Hulu",
		},
		{
			Id:     "520411",
			CityId: "5204",
			Name:   "Ropang",
		},
		{
			Id:     "520412",
			CityId: "5204",
			Name:   "Lape",
		},
		{
			Id:     "520413",
			CityId: "5204",
			Name:   "Plampang",
		},
		{
			Id:     "520414",
			CityId: "5204",
			Name:   "Empang",
		},
		{
			Id:     "520417",
			CityId: "5204",
			Name:   "Alas Barat",
		},
		{
			Id:     "520418",
			CityId: "5204",
			Name:   "Labuhan Badas",
		},
		{
			Id:     "520419",
			CityId: "5204",
			Name:   "Labangka",
		},
		{
			Id:     "520420",
			CityId: "5204",
			Name:   "Buer",
		},
		{
			Id:     "520421",
			CityId: "5204",
			Name:   "Rhee",
		},
		{
			Id:     "520422",
			CityId: "5204",
			Name:   "Unter Iwes",
		},
		{
			Id:     "520423",
			CityId: "5204",
			Name:   "Moyo Utara",
		},
		{
			Id:     "520424",
			CityId: "5204",
			Name:   "Maronge",
		},
		{
			Id:     "520425",
			CityId: "5204",
			Name:   "Tarano",
		},
		{
			Id:     "520426",
			CityId: "5204",
			Name:   "Lopok",
		},
		{
			Id:     "520427",
			CityId: "5204",
			Name:   "Lenangguar",
		},
		{
			Id:     "520428",
			CityId: "5204",
			Name:   "Orong Telu",
		},
		{
			Id:     "520429",
			CityId: "5204",
			Name:   "Lantung",
		},
		{
			Id:     "520501",
			CityId: "5205",
			Name:   "Dompu",
		},
		{
			Id:     "520502",
			CityId: "5205",
			Name:   "Kempo",
		},
		{
			Id:     "520503",
			CityId: "5205",
			Name:   "Hu'u",
		},
		{
			Id:     "520504",
			CityId: "5205",
			Name:   "Kilo",
		},
		{
			Id:     "520505",
			CityId: "5205",
			Name:   "Woja",
		},
		{
			Id:     "520506",
			CityId: "5205",
			Name:   "Pekat",
		},
		{
			Id:     "520507",
			CityId: "5205",
			Name:   "Manggalewa",
		},
		{
			Id:     "520508",
			CityId: "5205",
			Name:   "Pajo",
		},
		{
			Id:     "520601",
			CityId: "5206",
			Name:   "Monta",
		},
		{
			Id:     "520602",
			CityId: "5206",
			Name:   "Bolo",
		},
		{
			Id:     "520603",
			CityId: "5206",
			Name:   "Woha",
		},
		{
			Id:     "520604",
			CityId: "5206",
			Name:   "Belo",
		},
		{
			Id:     "520605",
			CityId: "5206",
			Name:   "Wawo",
		},
		{
			Id:     "520606",
			CityId: "5206",
			Name:   "Sape",
		},
		{
			Id:     "520607",
			CityId: "5206",
			Name:   "Wera",
		},
		{
			Id:     "520608",
			CityId: "5206",
			Name:   "Donggo",
		},
		{
			Id:     "520609",
			CityId: "5206",
			Name:   "Sanggar",
		},
		{
			Id:     "520610",
			CityId: "5206",
			Name:   "Ambalawi",
		},
		{
			Id:     "520611",
			CityId: "5206",
			Name:   "Langgudu",
		},
		{
			Id:     "520612",
			CityId: "5206",
			Name:   "Lambu",
		},
		{
			Id:     "520613",
			CityId: "5206",
			Name:   "Madapangga",
		},
		{
			Id:     "520614",
			CityId: "5206",
			Name:   "Tambora",
		},
		{
			Id:     "520615",
			CityId: "5206",
			Name:   "Soromandi",
		},
		{
			Id:     "520616",
			CityId: "5206",
			Name:   "Parado",
		},
		{
			Id:     "520617",
			CityId: "5206",
			Name:   "Lambitu",
		},
		{
			Id:     "520618",
			CityId: "5206",
			Name:   "Palibelo",
		},
		{
			Id:     "520701",
			CityId: "5207",
			Name:   "Jereweh",
		},
		{
			Id:     "520702",
			CityId: "5207",
			Name:   "Taliwang",
		},
		{
			Id:     "520703",
			CityId: "5207",
			Name:   "Seteluk",
		},
		{
			Id:     "520704",
			CityId: "5207",
			Name:   "Sekongkang",
		},
		{
			Id:     "520705",
			CityId: "5207",
			Name:   "Brang Rea",
		},
		{
			Id:     "520706",
			CityId: "5207",
			Name:   "Poto Tano",
		},
		{
			Id:     "520707",
			CityId: "5207",
			Name:   "Brang Ene",
		},
		{
			Id:     "520708",
			CityId: "5207",
			Name:   "Maluk",
		},
		{
			Id:     "520801",
			CityId: "5208",
			Name:   "Tanjung",
		},
		{
			Id:     "520802",
			CityId: "5208",
			Name:   "Gangga",
		},
		{
			Id:     "520803",
			CityId: "5208",
			Name:   "Kayangan",
		},
		{
			Id:     "520804",
			CityId: "5208",
			Name:   "Bayan",
		},
		{
			Id:     "520805",
			CityId: "5208",
			Name:   "Pemenang",
		},
		{
			Id:     "527101",
			CityId: "5271",
			Name:   "Ampenan",
		},
		{
			Id:     "527102",
			CityId: "5271",
			Name:   "Mataram",
		},
		{
			Id:     "527103",
			CityId: "5271",
			Name:   "Cakranegara",
		},
		{
			Id:     "527104",
			CityId: "5271",
			Name:   "Sekarbela",
		},
		{
			Id:     "527105",
			CityId: "5271",
			Name:   "Selaprang",
		},
		{
			Id:     "527106",
			CityId: "5271",
			Name:   "Sandubaya",
		},
		{
			Id:     "527201",
			CityId: "5272",
			Name:   "RasanaE Barat",
		},
		{
			Id:     "527202",
			CityId: "5272",
			Name:   "RasanaE Timur",
		},
		{
			Id:     "527203",
			CityId: "5272",
			Name:   "Asakota",
		},
		{
			Id:     "527204",
			CityId: "5272",
			Name:   "Raba",
		},
		{
			Id:     "527205",
			CityId: "5272",
			Name:   "Mpunda",
		},
		{
			Id:     "530104",
			CityId: "5301",
			Name:   "Semau",
		},
		{
			Id:     "530105",
			CityId: "5301",
			Name:   "Kupang Barat",
		},
		{
			Id:     "530106",
			CityId: "5301",
			Name:   "Kupang Timur",
		},
		{
			Id:     "530107",
			CityId: "5301",
			Name:   "Sulamu",
		},
		{
			Id:     "530108",
			CityId: "5301",
			Name:   "Kupang Tengah",
		},
		{
			Id:     "530109",
			CityId: "5301",
			Name:   "Amarasi",
		},
		{
			Id:     "530110",
			CityId: "5301",
			Name:   "Fatuleu",
		},
		{
			Id:     "530111",
			CityId: "5301",
			Name:   "Takari",
		},
		{
			Id:     "530112",
			CityId: "5301",
			Name:   "Amfoang Selatan",
		},
		{
			Id:     "530113",
			CityId: "5301",
			Name:   "Amfoang Utara",
		},
		{
			Id:     "530116",
			CityId: "5301",
			Name:   "Nekamese",
		},
		{
			Id:     "530117",
			CityId: "5301",
			Name:   "Amarasi Barat",
		},
		{
			Id:     "530118",
			CityId: "5301",
			Name:   "Amarasi Selatan",
		},
		{
			Id:     "530119",
			CityId: "5301",
			Name:   "Amarasi Timur",
		},
		{
			Id:     "530120",
			CityId: "5301",
			Name:   "Amabi Oefeto Timur",
		},
		{
			Id:     "530121",
			CityId: "5301",
			Name:   "Amfoang Barat Daya",
		},
		{
			Id:     "530122",
			CityId: "5301",
			Name:   "Amfoang Barat Laut",
		},
		{
			Id:     "530123",
			CityId: "5301",
			Name:   "Semau Selatan",
		},
		{
			Id:     "530124",
			CityId: "5301",
			Name:   "Taebenu",
		},
		{
			Id:     "530125",
			CityId: "5301",
			Name:   "Amabi Oefeto",
		},
		{
			Id:     "530126",
			CityId: "5301",
			Name:   "Amfoang Timur",
		},
		{
			Id:     "530127",
			CityId: "5301",
			Name:   "Fatuleu Barat",
		},
		{
			Id:     "530128",
			CityId: "5301",
			Name:   "Fatuleu Tengah",
		},
		{
			Id:     "530130",
			CityId: "5301",
			Name:   "Amfoang Tengah",
		},
		{
			Id:     "530201",
			CityId: "5302",
			Name:   "Kota Soe",
		},
		{
			Id:     "530202",
			CityId: "5302",
			Name:   "Mollo Selatan",
		},
		{
			Id:     "530203",
			CityId: "5302",
			Name:   "Mollo Utara",
		},
		{
			Id:     "530204",
			CityId: "5302",
			Name:   "Amanuban Timur",
		},
		{
			Id:     "530205",
			CityId: "5302",
			Name:   "Amanuban Tengah",
		},
		{
			Id:     "530206",
			CityId: "5302",
			Name:   "Amanuban Selatan",
		},
		{
			Id:     "530207",
			CityId: "5302",
			Name:   "Amanuban Barat",
		},
		{
			Id:     "530208",
			CityId: "5302",
			Name:   "Amanatun Selatan",
		},
		{
			Id:     "530209",
			CityId: "5302",
			Name:   "Amanatun Utara",
		},
		{
			Id:     "530210",
			CityId: "5302",
			Name:   "KI'E",
		},
		{
			Id:     "530211",
			CityId: "5302",
			Name:   "Kuanfatu",
		},
		{
			Id:     "530212",
			CityId: "5302",
			Name:   "Fatumnasi",
		},
		{
			Id:     "530213",
			CityId: "5302",
			Name:   "Polen",
		},
		{
			Id:     "530214",
			CityId: "5302",
			Name:   "Batu Putih",
		},
		{
			Id:     "530215",
			CityId: "5302",
			Name:   "Boking",
		},
		{
			Id:     "530216",
			CityId: "5302",
			Name:   "Toianas",
		},
		{
			Id:     "530217",
			CityId: "5302",
			Name:   "Nunkolo",
		},
		{
			Id:     "530218",
			CityId: "5302",
			Name:   "Oenino",
		},
		{
			Id:     "530219",
			CityId: "5302",
			Name:   "Kolbano",
		},
		{
			Id:     "530220",
			CityId: "5302",
			Name:   "Kot olin",
		},
		{
			Id:     "530221",
			CityId: "5302",
			Name:   "Kualin",
		},
		{
			Id:     "530222",
			CityId: "5302",
			Name:   "Mollo Barat",
		},
		{
			Id:     "530223",
			CityId: "5302",
			Name:   "Kok Baun",
		},
		{
			Id:     "530224",
			CityId: "5302",
			Name:   "Noebana",
		},
		{
			Id:     "530225",
			CityId: "5302",
			Name:   "Santian",
		},
		{
			Id:     "530226",
			CityId: "5302",
			Name:   "Noebeba",
		},
		{
			Id:     "530227",
			CityId: "5302",
			Name:   "Kuatnana",
		},
		{
			Id:     "530228",
			CityId: "5302",
			Name:   "Fautmolo",
		},
		{
			Id:     "530229",
			CityId: "5302",
			Name:   "Fatukopa",
		},
		{
			Id:     "530230",
			CityId: "5302",
			Name:   "Mollo Tengah",
		},
		{
			Id:     "530231",
			CityId: "5302",
			Name:   "Tobu",
		},
		{
			Id:     "530232",
			CityId: "5302",
			Name:   "Nunbena",
		},
		{
			Id:     "530301",
			CityId: "5303",
			Name:   "Miomaffo Timur",
		},
		{
			Id:     "530302",
			CityId: "5303",
			Name:   "Miomaffo Barat",
		},
		{
			Id:     "530303",
			CityId: "5303",
			Name:   "Biboki Selatan",
		},
		{
			Id:     "530304",
			CityId: "5303",
			Name:   "Noemuti",
		},
		{
			Id:     "530305",
			CityId: "5303",
			Name:   "Kota Kefamenanu",
		},
		{
			Id:     "530306",
			CityId: "5303",
			Name:   "Biboki Utara",
		},
		{
			Id:     "530307",
			CityId: "5303",
			Name:   "Biboki Anleu",
		},
		{
			Id:     "530308",
			CityId: "5303",
			Name:   "Insana",
		},
		{
			Id:     "530309",
			CityId: "5303",
			Name:   "Insana Utara",
		},
		{
			Id:     "530310",
			CityId: "5303",
			Name:   "Noemuti Timur",
		},
		{
			Id:     "530311",
			CityId: "5303",
			Name:   "Miomaffo Tengah",
		},
		{
			Id:     "530312",
			CityId: "5303",
			Name:   "Musi",
		},
		{
			Id:     "530313",
			CityId: "5303",
			Name:   "Mutis",
		},
		{
			Id:     "530314",
			CityId: "5303",
			Name:   "Bikomi Selatan",
		},
		{
			Id:     "530315",
			CityId: "5303",
			Name:   "Bikomi Tengah",
		},
		{
			Id:     "530316",
			CityId: "5303",
			Name:   "Bikomi Nilulat",
		},
		{
			Id:     "530317",
			CityId: "5303",
			Name:   "Bikomi Utara",
		},
		{
			Id:     "530318",
			CityId: "5303",
			Name:   "Naibenu",
		},
		{
			Id:     "530319",
			CityId: "5303",
			Name:   "Insana Fafinesu",
		},
		{
			Id:     "530320",
			CityId: "5303",
			Name:   "Insana Barat",
		},
		{
			Id:     "530321",
			CityId: "5303",
			Name:   "Insana Tengah",
		},
		{
			Id:     "530322",
			CityId: "5303",
			Name:   "Biboki Tan Pah",
		},
		{
			Id:     "530323",
			CityId: "5303",
			Name:   "Biboki Moenleu",
		},
		{
			Id:     "530324",
			CityId: "5303",
			Name:   "Biboki Feotleu",
		},
		{
			Id:     "530401",
			CityId: "5304",
			Name:   "Lamaknen",
		},
		{
			Id:     "530402",
			CityId: "5304",
			Name:   "TasifetoTimur",
		},
		{
			Id:     "530403",
			CityId: "5304",
			Name:   "Raihat",
		},
		{
			Id:     "530404",
			CityId: "5304",
			Name:   "Tasifeto Barat",
		},
		{
			Id:     "530405",
			CityId: "5304",
			Name:   "Kakuluk Mesak",
		},
		{
			Id:     "530412",
			CityId: "5304",
			Name:   "Kota Atambua",
		},
		{
			Id:     "530413",
			CityId: "5304",
			Name:   "Raimanuk",
		},
		{
			Id:     "530417",
			CityId: "5304",
			Name:   "Lasiolat",
		},
		{
			Id:     "530418",
			CityId: "5304",
			Name:   "Lamaknen Selatan",
		},
		{
			Id:     "530421",
			CityId: "5304",
			Name:   "Atambua Barat",
		},
		{
			Id:     "530422",
			CityId: "5304",
			Name:   "Atambua Selatan",
		},
		{
			Id:     "530423",
			CityId: "5304",
			Name:   "Nanaet Duabesi",
		},
		{
			Id:     "530501",
			CityId: "5305",
			Name:   "Teluk Mutiara",
		},
		{
			Id:     "530502",
			CityId: "5305",
			Name:   "Alor Barat Laut",
		},
		{
			Id:     "530503",
			CityId: "5305",
			Name:   "Alor Barat Daya",
		},
		{
			Id:     "530504",
			CityId: "5305",
			Name:   "Alor Selatan",
		},
		{
			Id:     "530505",
			CityId: "5305",
			Name:   "Alor Timur",
		},
		{
			Id:     "530506",
			CityId: "5305",
			Name:   "Pantar",
		},
		{
			Id:     "530507",
			CityId: "5305",
			Name:   "Alor Tengah Utara",
		},
		{
			Id:     "530508",
			CityId: "5305",
			Name:   "Alor Timur Laut",
		},
		{
			Id:     "530509",
			CityId: "5305",
			Name:   "Pantar Barat",
		},
		{
			Id:     "530510",
			CityId: "5305",
			Name:   "Kabola",
		},
		{
			Id:     "530511",
			CityId: "5305",
			Name:   "Pulau Pura",
		},
		{
			Id:     "530512",
			CityId: "5305",
			Name:   "Mataru",
		},
		{
			Id:     "530513",
			CityId: "5305",
			Name:   "Pureman",
		},
		{
			Id:     "530514",
			CityId: "5305",
			Name:   "Pantar Timur",
		},
		{
			Id:     "530515",
			CityId: "5305",
			Name:   "Lembur",
		},
		{
			Id:     "530516",
			CityId: "5305",
			Name:   "Pantar Tengah",
		},
		{
			Id:     "530517",
			CityId: "5305",
			Name:   "Pantar Baru Laut",
		},
		{
			Id:     "530518",
			CityId: "5305",
			Name:   "Abad Selatan",
		},
		{
			Id:     "530601",
			CityId: "5306",
			Name:   "Wulanggitang",
		},
		{
			Id:     "530602",
			CityId: "5306",
			Name:   "Titehena",
		},
		{
			Id:     "530603",
			CityId: "5306",
			Name:   "Larantuka",
		},
		{
			Id:     "530604",
			CityId: "5306",
			Name:   "Ile Mandiri",
		},
		{
			Id:     "530605",
			CityId: "5306",
			Name:   "Tanjung Bunga",
		},
		{
			Id:     "530606",
			CityId: "5306",
			Name:   "Solor Barat",
		},
		{
			Id:     "530607",
			CityId: "5306",
			Name:   "Solor Timur",
		},
		{
			Id:     "530608",
			CityId: "5306",
			Name:   "Adonara Barat",
		},
		{
			Id:     "530609",
			CityId: "5306",
			Name:   "Wotan Ulumando",
		},
		{
			Id:     "530610",
			CityId: "5306",
			Name:   "Adonara Timur",
		},
		{
			Id:     "530611",
			CityId: "5306",
			Name:   "Kelubagolit",
		},
		{
			Id:     "530612",
			CityId: "5306",
			Name:   "Witihama",
		},
		{
			Id:     "530613",
			CityId: "5306",
			Name:   "Ile Boleng",
		},
		{
			Id:     "530614",
			CityId: "5306",
			Name:   "Demon Pagong",
		},
		{
			Id:     "530615",
			CityId: "5306",
			Name:   "Lewolema",
		},
		{
			Id:     "530616",
			CityId: "5306",
			Name:   "Ile Bura",
		},
		{
			Id:     "530617",
			CityId: "5306",
			Name:   "Adonara",
		},
		{
			Id:     "530618",
			CityId: "5306",
			Name:   "Adonara Tengah",
		},
		{
			Id:     "530619",
			CityId: "5306",
			Name:   "Solor Selatan",
		},
		{
			Id:     "530701",
			CityId: "5307",
			Name:   "Paga",
		},
		{
			Id:     "530702",
			CityId: "5307",
			Name:   "Mego",
		},
		{
			Id:     "530703",
			CityId: "5307",
			Name:   "Lela",
		},
		{
			Id:     "530704",
			CityId: "5307",
			Name:   "Nita",
		},
		{
			Id:     "530705",
			CityId: "5307",
			Name:   "Alok",
		},
		{
			Id:     "530706",
			CityId: "5307",
			Name:   "Palue",
		},
		{
			Id:     "530707",
			CityId: "5307",
			Name:   "Nelle",
		},
		{
			Id:     "530708",
			CityId: "5307",
			Name:   "Talibura",
		},
		{
			Id:     "530709",
			CityId: "5307",
			Name:   "Waigete",
		},
		{
			Id:     "530710",
			CityId: "5307",
			Name:   "Kewapante",
		},
		{
			Id:     "530711",
			CityId: "5307",
			Name:   "Bola",
		},
		{
			Id:     "530712",
			CityId: "5307",
			Name:   "Magepanda",
		},
		{
			Id:     "530713",
			CityId: "5307",
			Name:   "Waiblama",
		},
		{
			Id:     "530714",
			CityId: "5307",
			Name:   "Alok Barat",
		},
		{
			Id:     "530715",
			CityId: "5307",
			Name:   "Alok Timur",
		},
		{
			Id:     "530716",
			CityId: "5307",
			Name:   "Koting",
		},
		{
			Id:     "530717",
			CityId: "5307",
			Name:   "Tana Wawo",
		},
		{
			Id:     "530718",
			CityId: "5307",
			Name:   "Hewokloang",
		},
		{
			Id:     "530719",
			CityId: "5307",
			Name:   "Kangae",
		},
		{
			Id:     "530720",
			CityId: "5307",
			Name:   "Doreng",
		},
		{
			Id:     "530721",
			CityId: "5307",
			Name:   "Mapitara",
		},
		{
			Id:     "530801",
			CityId: "5308",
			Name:   "Nangapanda",
		},
		{
			Id:     "530802",
			CityId: "5308",
			Name:   "Pulau Ende",
		},
		{
			Id:     "530803",
			CityId: "5308",
			Name:   "Ende",
		},
		{
			Id:     "530804",
			CityId: "5308",
			Name:   "Ende Selatan",
		},
		{
			Id:     "530805",
			CityId: "5308",
			Name:   "Ndona",
		},
		{
			Id:     "530806",
			CityId: "5308",
			Name:   "Detusoko",
		},
		{
			Id:     "530807",
			CityId: "5308",
			Name:   "Wewaria",
		},
		{
			Id:     "530808",
			CityId: "5308",
			Name:   "Wolowaru",
		},
		{
			Id:     "530809",
			CityId: "5308",
			Name:   "Wolojita",
		},
		{
			Id:     "530810",
			CityId: "5308",
			Name:   "Maurole",
		},
		{
			Id:     "530811",
			CityId: "5308",
			Name:   "Maukaro",
		},
		{
			Id:     "530812",
			CityId: "5308",
			Name:   "Lio Timur",
		},
		{
			Id:     "530813",
			CityId: "5308",
			Name:   "Kota Baru",
		},
		{
			Id:     "530814",
			CityId: "5308",
			Name:   "Kelimutu",
		},
		{
			Id:     "530815",
			CityId: "5308",
			Name:   "Detukeli",
		},
		{
			Id:     "530816",
			CityId: "5308",
			Name:   "Ndona Timur",
		},
		{
			Id:     "530817",
			CityId: "5308",
			Name:   "Ndori",
		},
		{
			Id:     "530818",
			CityId: "5308",
			Name:   "Ende Utara",
		},
		{
			Id:     "530819",
			CityId: "5308",
			Name:   "Ende Tengah",
		},
		{
			Id:     "530820",
			CityId: "5308",
			Name:   "Ende Timur",
		},
		{
			Id:     "530821",
			CityId: "5308",
			Name:   "Lepembusu Kelisoke",
		},
		{
			Id:     "530901",
			CityId: "5309",
			Name:   "Aimere",
		},
		{
			Id:     "530902",
			CityId: "5309",
			Name:   "Golewa",
		},
		{
			Id:     "530906",
			CityId: "5309",
			Name:   "Bajawa",
		},
		{
			Id:     "530907",
			CityId: "5309",
			Name:   "Soa",
		},
		{
			Id:     "530909",
			CityId: "5309",
			Name:   "Riung",
		},
		{
			Id:     "530912",
			CityId: "5309",
			Name:   "Jerebuu",
		},
		{
			Id:     "530914",
			CityId: "5309",
			Name:   "Riung Barat",
		},
		{
			Id:     "530915",
			CityId: "5309",
			Name:   "Bajawa Utara",
		},
		{
			Id:     "530916",
			CityId: "5309",
			Name:   "Wolomeze",
		},
		{
			Id:     "530918",
			CityId: "5309",
			Name:   "Golewa Selatan",
		},
		{
			Id:     "530919",
			CityId: "5309",
			Name:   "Golewa Barat",
		},
		{
			Id:     "530920",
			CityId: "5309",
			Name:   "Inerie",
		},
		{
			Id:     "531001",
			CityId: "5310",
			Name:   "Wae Rii",
		},
		{
			Id:     "531003",
			CityId: "5310",
			Name:   "Ruteng",
		},
		{
			Id:     "531005",
			CityId: "5310",
			Name:   "Satar Mese",
		},
		{
			Id:     "531006",
			CityId: "5310",
			Name:   "Cibal",
		},
		{
			Id:     "531011",
			CityId: "5310",
			Name:   "Reok",
		},
		{
			Id:     "531012",
			CityId: "5310",
			Name:   "Langke Rembong",
		},
		{
			Id:     "531013",
			CityId: "5310",
			Name:   "Satar Mese Barat",
		},
		{
			Id:     "531014",
			CityId: "5310",
			Name:   "Rahong Utara",
		},
		{
			Id:     "531015",
			CityId: "5310",
			Name:   "Lelak",
		},
		{
			Id:     "531016",
			CityId: "5310",
			Name:   "Reok Barat",
		},
		{
			Id:     "531017",
			CityId: "5310",
			Name:   "Cibal barat",
		},
		{
			Id:     "531018",
			CityId: "5310",
			Name:   "Satar Mese Utara",
		},
		{
			Id:     "531101",
			CityId: "5311",
			Name:   "Kota Waingapu",
		},
		{
			Id:     "531102",
			CityId: "5311",
			Name:   "Haharu",
		},
		{
			Id:     "531103",
			CityId: "5311",
			Name:   "Lewa",
		},
		{
			Id:     "531104",
			CityId: "5311",
			Name:   "Nggaha Ori Angu",
		},
		{
			Id:     "531105",
			CityId: "5311",
			Name:   "Tabundung",
		},
		{
			Id:     "531106",
			CityId: "5311",
			Name:   "Pinu Pahar",
		},
		{
			Id:     "531107",
			CityId: "5311",
			Name:   "Pandawai",
		},
		{
			Id:     "531108",
			CityId: "5311",
			Name:   "Umalulu",
		},
		{
			Id:     "531109",
			CityId: "5311",
			Name:   "Rindi",
		},
		{
			Id:     "531110",
			CityId: "5311",
			Name:   "Pahunga Lodu",
		},
		{
			Id:     "531111",
			CityId: "5311",
			Name:   "Wulla Waijelu",
		},
		{
			Id:     "531112",
			CityId: "5311",
			Name:   "Paberiwai",
		},
		{
			Id:     "531113",
			CityId: "5311",
			Name:   "Karera",
		},
		{
			Id:     "531114",
			CityId: "5311",
			Name:   "Kahaungu Eti",
		},
		{
			Id:     "531115",
			CityId: "5311",
			Name:   "Matawai La Pawu",
		},
		{
			Id:     "531116",
			CityId: "5311",
			Name:   "Kambera",
		},
		{
			Id:     "531117",
			CityId: "5311",
			Name:   "Kambata Mapambuhang",
		},
		{
			Id:     "531118",
			CityId: "5311",
			Name:   "Lewa Tidahu",
		},
		{
			Id:     "531119",
			CityId: "5311",
			Name:   "Katala Hamu Lingu",
		},
		{
			Id:     "531120",
			CityId: "5311",
			Name:   "Kanatang",
		},
		{
			Id:     "531121",
			CityId: "5311",
			Name:   "Ngadu Ngala",
		},
		{
			Id:     "531122",
			CityId: "5311",
			Name:   "Mahu",
		},
		{
			Id:     "531204",
			CityId: "5312",
			Name:   "Tana Righu",
		},
		{
			Id:     "531210",
			CityId: "5312",
			Name:   "Loli",
		},
		{
			Id:     "531211",
			CityId: "5312",
			Name:   "Wanokaka",
		},
		{
			Id:     "531212",
			CityId: "5312",
			Name:   "Lamboya",
		},
		{
			Id:     "531215",
			CityId: "5312",
			Name:   "Kota Waikabubak",
		},
		{
			Id:     "531218",
			CityId: "5312",
			Name:   "Laboya Barat",
		},
		{
			Id:     "531301",
			CityId: "5313",
			Name:   "Naga Wutung",
		},
		{
			Id:     "531302",
			CityId: "5313",
			Name:   "Atadei",
		},
		{
			Id:     "531303",
			CityId: "5313",
			Name:   "Ile Ape",
		},
		{
			Id:     "531304",
			CityId: "5313",
			Name:   "Lebatukan",
		},
		{
			Id:     "531305",
			CityId: "5313",
			Name:   "Nubatukan",
		},
		{
			Id:     "531306",
			CityId: "5313",
			Name:   "Omesuri",
		},
		{
			Id:     "531307",
			CityId: "5313",
			Name:   "Buyasuri",
		},
		{
			Id:     "531308",
			CityId: "5313",
			Name:   "Wulandoni",
		},
		{
			Id:     "531309",
			CityId: "5313",
			Name:   "Ile Ape Timur",
		},
		{
			Id:     "531401",
			CityId: "5314",
			Name:   "Rote Barat Daya",
		},
		{
			Id:     "531402",
			CityId: "5314",
			Name:   "Rote Barat Laut",
		},
		{
			Id:     "531403",
			CityId: "5314",
			Name:   "Lobalain",
		},
		{
			Id:     "531404",
			CityId: "5314",
			Name:   "Rote Tengah",
		},
		{
			Id:     "531405",
			CityId: "5314",
			Name:   "Pantai Baru",
		},
		{
			Id:     "531406",
			CityId: "5314",
			Name:   "Rote Timur",
		},
		{
			Id:     "531407",
			CityId: "5314",
			Name:   "Rote Barat",
		},
		{
			Id:     "531408",
			CityId: "5314",
			Name:   "Rote Selatan",
		},
		{
			Id:     "531409",
			CityId: "5314",
			Name:   "Ndao Nuse",
		},
		{
			Id:     "531410",
			CityId: "5314",
			Name:   "Landu Leko",
		},
		{
			Id:     "531411",
			CityId: "5314",
			Name:   "Loaholu",
		},
		{
			Id:     "531501",
			CityId: "5315",
			Name:   "Macang Pacar",
		},
		{
			Id:     "531502",
			CityId: "5315",
			Name:   "Kuwus",
		},
		{
			Id:     "531503",
			CityId: "5315",
			Name:   "Lembor",
		},
		{
			Id:     "531504",
			CityId: "5315",
			Name:   "Sano Nggoang",
		},
		{
			Id:     "531505",
			CityId: "5315",
			Name:   "Komodo",
		},
		{
			Id:     "531506",
			CityId: "5315",
			Name:   "Boleng",
		},
		{
			Id:     "531507",
			CityId: "5315",
			Name:   "Welak",
		},
		{
			Id:     "531508",
			CityId: "5315",
			Name:   "Ndoso",
		},
		{
			Id:     "531509",
			CityId: "5315",
			Name:   "Lembor Selatan",
		},
		{
			Id:     "531510",
			CityId: "5315",
			Name:   "Mbeliling",
		},
		{
			Id:     "531511",
			CityId: "5315",
			Name:   "Pacar",
		},
		{
			Id:     "531512",
			CityId: "5315",
			Name:   "Kuwus Barat",
		},
		{
			Id:     "531601",
			CityId: "5316",
			Name:   "Aesesa",
		},
		{
			Id:     "531602",
			CityId: "5316",
			Name:   "Nangaroro",
		},
		{
			Id:     "531603",
			CityId: "5316",
			Name:   "Boawae",
		},
		{
			Id:     "531604",
			CityId: "5316",
			Name:   "Mauponggo",
		},
		{
			Id:     "531605",
			CityId: "5316",
			Name:   "Wolowae",
		},
		{
			Id:     "531606",
			CityId: "5316",
			Name:   "Keo Tengah",
		},
		{
			Id:     "531607",
			CityId: "5316",
			Name:   "Aesesa Selatan",
		},
		{
			Id:     "531701",
			CityId: "5317",
			Name:   "Katiku Tana",
		},
		{
			Id:     "531702",
			CityId: "5317",
			Name:   "Umbu Ratu Nggay Barat",
		},
		{
			Id:     "531703",
			CityId: "5317",
			Name:   "Mamboro",
		},
		{
			Id:     "531704",
			CityId: "5317",
			Name:   "Umbu Ratu Nggay",
		},
		{
			Id:     "531705",
			CityId: "5317",
			Name:   "Katiku Tana Selatan",
		},
		{
			Id:     "531706",
			CityId: "5317",
			Name:   "Umbu Ratu Nggay Tengah",
		},
		{
			Id:     "531801",
			CityId: "5318",
			Name:   "Loura",
		},
		{
			Id:     "531802",
			CityId: "5318",
			Name:   "Wewewa Utara",
		},
		{
			Id:     "531803",
			CityId: "5318",
			Name:   "Wewewa Timur",
		},
		{
			Id:     "531804",
			CityId: "5318",
			Name:   "Wewewa Barat",
		},
		{
			Id:     "531805",
			CityId: "5318",
			Name:   "Wewewa Selatan",
		},
		{
			Id:     "531806",
			CityId: "5318",
			Name:   "Kodi Bangedo",
		},
		{
			Id:     "531807",
			CityId: "5318",
			Name:   "Kodi",
		},
		{
			Id:     "531808",
			CityId: "5318",
			Name:   "Kodi Utara",
		},
		{
			Id:     "531809",
			CityId: "5318",
			Name:   "Kota Tambolaka",
		},
		{
			Id:     "531810",
			CityId: "5318",
			Name:   "Wewewa Tengah",
		},
		{
			Id:     "531811",
			CityId: "5318",
			Name:   "Kodi Balaghar",
		},
		{
			Id:     "531901",
			CityId: "5319",
			Name:   "Borong",
		},
		{
			Id:     "531902",
			CityId: "5319",
			Name:   "Lamba Leda Selatan",
		},
		{
			Id:     "531903",
			CityId: "5319",
			Name:   "Lamba Leda",
		},
		{
			Id:     "531904",
			CityId: "5319",
			Name:   "Sambi Rampas",
		},
		{
			Id:     "531905",
			CityId: "5319",
			Name:   "Elar",
		},
		{
			Id:     "531906",
			CityId: "5319",
			Name:   "Kota Komba",
		},
		{
			Id:     "531907",
			CityId: "5319",
			Name:   "Rana Mese",
		},
		{
			Id:     "531908",
			CityId: "5319",
			Name:   "Lamba Leda Timur",
		},
		{
			Id:     "531909",
			CityId: "5319",
			Name:   "Elar Selatan",
		},
		{
			Id:     "531910",
			CityId: "5319",
			Name:   "Kota Komba Utara",
		},
		{
			Id:     "531911",
			CityId: "5319",
			Name:   "Lamba Leda Utara",
		},
		{
			Id:     "531912",
			CityId: "5319",
			Name:   "Congkar",
		},
		{
			Id:     "532001",
			CityId: "5320",
			Name:   "Sabu Barat",
		},
		{
			Id:     "532002",
			CityId: "5320",
			Name:   "Sabu Tengah",
		},
		{
			Id:     "532003",
			CityId: "5320",
			Name:   "Sabu Timur",
		},
		{
			Id:     "532004",
			CityId: "5320",
			Name:   "Sabu Liae",
		},
		{
			Id:     "532005",
			CityId: "5320",
			Name:   "Hawu Mehara",
		},
		{
			Id:     "532006",
			CityId: "5320",
			Name:   "Raijua",
		},
		{
			Id:     "532101",
			CityId: "5321",
			Name:   "Malaka Tengah",
		},
		{
			Id:     "532102",
			CityId: "5321",
			Name:   "Malaka Barat",
		},
		{
			Id:     "532103",
			CityId: "5321",
			Name:   "Wewiku",
		},
		{
			Id:     "532104",
			CityId: "5321",
			Name:   "Weliman",
		},
		{
			Id:     "532105",
			CityId: "5321",
			Name:   "Rinhat",
		},
		{
			Id:     "532106",
			CityId: "5321",
			Name:   "Io Kufeu",
		},
		{
			Id:     "532107",
			CityId: "5321",
			Name:   "Sasitamean",
		},
		{
			Id:     "532108",
			CityId: "5321",
			Name:   "Laenmanen",
		},
		{
			Id:     "532109",
			CityId: "5321",
			Name:   "Malaka Timur",
		},
		{
			Id:     "532110",
			CityId: "5321",
			Name:   "Kobalima Timur",
		},
		{
			Id:     "532111",
			CityId: "5321",
			Name:   "Kobalima",
		},
		{
			Id:     "532112",
			CityId: "5321",
			Name:   "Botin Leobele",
		},
		{
			Id:     "537101",
			CityId: "5371",
			Name:   "Alak",
		},
		{
			Id:     "537102",
			CityId: "5371",
			Name:   "Maulafa",
		},
		{
			Id:     "537103",
			CityId: "5371",
			Name:   "Kelapa Lima",
		},
		{
			Id:     "537104",
			CityId: "5371",
			Name:   "Oebobo",
		},
		{
			Id:     "537105",
			CityId: "5371",
			Name:   "Kota Raja",
		},
		{
			Id:     "537106",
			CityId: "5371",
			Name:   "Kota Lama",
		},
		{
			Id:     "610101",
			CityId: "6101",
			Name:   "Sambas",
		},
		{
			Id:     "610102",
			CityId: "6101",
			Name:   "Teluk Keramat",
		},
		{
			Id:     "610103",
			CityId: "6101",
			Name:   "Jawai",
		},
		{
			Id:     "610104",
			CityId: "6101",
			Name:   "Tebas",
		},
		{
			Id:     "610105",
			CityId: "6101",
			Name:   "Pemangkat",
		},
		{
			Id:     "610106",
			CityId: "6101",
			Name:   "Sejangkung",
		},
		{
			Id:     "610107",
			CityId: "6101",
			Name:   "Selakau",
		},
		{
			Id:     "610108",
			CityId: "6101",
			Name:   "Paloh",
		},
		{
			Id:     "610109",
			CityId: "6101",
			Name:   "Sajingan Besar",
		},
		{
			Id:     "610110",
			CityId: "6101",
			Name:   "Subah",
		},
		{
			Id:     "610111",
			CityId: "6101",
			Name:   "Galing",
		},
		{
			Id:     "610112",
			CityId: "6101",
			Name:   "Tekarang",
		},
		{
			Id:     "610113",
			CityId: "6101",
			Name:   "Semparuk",
		},
		{
			Id:     "610114",
			CityId: "6101",
			Name:   "Sajad",
		},
		{
			Id:     "610115",
			CityId: "6101",
			Name:   "Sebawi",
		},
		{
			Id:     "610116",
			CityId: "6101",
			Name:   "Jawai Selatan",
		},
		{
			Id:     "610117",
			CityId: "6101",
			Name:   "Tangaran",
		},
		{
			Id:     "610118",
			CityId: "6101",
			Name:   "Salatiga",
		},
		{
			Id:     "610119",
			CityId: "6101",
			Name:   "Selakau Timur",
		},
		{
			Id:     "610201",
			CityId: "6102",
			Name:   "Mempawah Hilir",
		},
		{
			Id:     "610206",
			CityId: "6102",
			Name:   "Toho",
		},
		{
			Id:     "610207",
			CityId: "6102",
			Name:   "Sungai Pinyuh",
		},
		{
			Id:     "610208",
			CityId: "6102",
			Name:   "Jongkat",
		},
		{
			Id:     "610212",
			CityId: "6102",
			Name:   "Sungai Kunyit",
		},
		{
			Id:     "610215",
			CityId: "6102",
			Name:   "Segedong",
		},
		{
			Id:     "610216",
			CityId: "6102",
			Name:   "Anjongan",
		},
		{
			Id:     "610217",
			CityId: "6102",
			Name:   "Sadaniang",
		},
		{
			Id:     "610218",
			CityId: "6102",
			Name:   "Mempawah Timur",
		},
		{
			Id:     "610301",
			CityId: "6103",
			Name:   "Kapuas",
		},
		{
			Id:     "610302",
			CityId: "6103",
			Name:   "Mukok",
		},
		{
			Id:     "610303",
			CityId: "6103",
			Name:   "Noyan",
		},
		{
			Id:     "610304",
			CityId: "6103",
			Name:   "Jangkang",
		},
		{
			Id:     "610305",
			CityId: "6103",
			Name:   "Bonti",
		},
		{
			Id:     "610306",
			CityId: "6103",
			Name:   "Beduai",
		},
		{
			Id:     "610307",
			CityId: "6103",
			Name:   "Sekayam",
		},
		{
			Id:     "610308",
			CityId: "6103",
			Name:   "Kembayan",
		},
		{
			Id:     "610309",
			CityId: "6103",
			Name:   "Parindu",
		},
		{
			Id:     "610310",
			CityId: "6103",
			Name:   "Tayan Hulu",
		},
		{
			Id:     "610311",
			CityId: "6103",
			Name:   "Tayan Hilir",
		},
		{
			Id:     "610312",
			CityId: "6103",
			Name:   "Balai",
		},
		{
			Id:     "610313",
			CityId: "6103",
			Name:   "Toba",
		},
		{
			Id:     "610320",
			CityId: "6103",
			Name:   "Meliau",
		},
		{
			Id:     "610321",
			CityId: "6103",
			Name:   "Entikong",
		},
		{
			Id:     "610401",
			CityId: "6104",
			Name:   "Matan Hilir Utara",
		},
		{
			Id:     "610402",
			CityId: "6104",
			Name:   "Marau",
		},
		{
			Id:     "610403",
			CityId: "6104",
			Name:   "Manis Mata",
		},
		{
			Id:     "610404",
			CityId: "6104",
			Name:   "Kendawangan",
		},
		{
			Id:     "610405",
			CityId: "6104",
			Name:   "Sandai",
		},
		{
			Id:     "610407",
			CityId: "6104",
			Name:   "Sungai Laur",
		},
		{
			Id:     "610408",
			CityId: "6104",
			Name:   "Simpang Hulu",
		},
		{
			Id:     "610411",
			CityId: "6104",
			Name:   "Nanga Tayap",
		},
		{
			Id:     "610412",
			CityId: "6104",
			Name:   "Matan Hilir Selatan",
		},
		{
			Id:     "610413",
			CityId: "6104",
			Name:   "Tumbang Titi",
		},
		{
			Id:     "610414",
			CityId: "6104",
			Name:   "Jelai Hulu",
		},
		{
			Id:     "610416",
			CityId: "6104",
			Name:   "Delta Pawan",
		},
		{
			Id:     "610417",
			CityId: "6104",
			Name:   "Muara Pawan",
		},
		{
			Id:     "610418",
			CityId: "6104",
			Name:   "Benua Kayong",
		},
		{
			Id:     "610419",
			CityId: "6104",
			Name:   "Hulu Sungai",
		},
		{
			Id:     "610420",
			CityId: "6104",
			Name:   "Simpang Dua",
		},
		{
			Id:     "610421",
			CityId: "6104",
			Name:   "Air Upas",
		},
		{
			Id:     "610422",
			CityId: "6104",
			Name:   "Singkup",
		},
		{
			Id:     "610424",
			CityId: "6104",
			Name:   "Pemahan",
		},
		{
			Id:     "610425",
			CityId: "6104",
			Name:   "Sungai Melayu Rayak",
		},
		{
			Id:     "610501",
			CityId: "6105",
			Name:   "Sintang",
		},
		{
			Id:     "610502",
			CityId: "6105",
			Name:   "Tempunak",
		},
		{
			Id:     "610503",
			CityId: "6105",
			Name:   "Sepauk",
		},
		{
			Id:     "610504",
			CityId: "6105",
			Name:   "Ketungau Hilir",
		},
		{
			Id:     "610505",
			CityId: "6105",
			Name:   "Ketungau Tengah",
		},
		{
			Id:     "610506",
			CityId: "6105",
			Name:   "Ketungau Hulu",
		},
		{
			Id:     "610507",
			CityId: "6105",
			Name:   "Dedai",
		},
		{
			Id:     "610508",
			CityId: "6105",
			Name:   "Kayan Hilir",
		},
		{
			Id:     "610509",
			CityId: "6105",
			Name:   "Kayan Hulu",
		},
		{
			Id:     "610514",
			CityId: "6105",
			Name:   "Serawai",
		},
		{
			Id:     "610515",
			CityId: "6105",
			Name:   "Ambalau",
		},
		{
			Id:     "610519",
			CityId: "6105",
			Name:   "Kelam Permai",
		},
		{
			Id:     "610520",
			CityId: "6105",
			Name:   "Sungai Tebelian",
		},
		{
			Id:     "610521",
			CityId: "6105",
			Name:   "Binjai Hulu",
		},
		{
			Id:     "610601",
			CityId: "6106",
			Name:   "Putussibau Utara",
		},
		{
			Id:     "610602",
			CityId: "6106",
			Name:   "Bika",
		},
		{
			Id:     "610603",
			CityId: "6106",
			Name:   "Embaloh Hilir",
		},
		{
			Id:     "610604",
			CityId: "6106",
			Name:   "Embaloh Hulu",
		},
		{
			Id:     "610605",
			CityId: "6106",
			Name:   "Bunut Hilir",
		},
		{
			Id:     "610606",
			CityId: "6106",
			Name:   "Bunut Hulu",
		},
		{
			Id:     "610607",
			CityId: "6106",
			Name:   "Jongkong",
		},
		{
			Id:     "610608",
			CityId: "6106",
			Name:   "Hulu Gurung",
		},
		{
			Id:     "610609",
			CityId: "6106",
			Name:   "Selimbau",
		},
		{
			Id:     "610610",
			CityId: "6106",
			Name:   "Semitau",
		},
		{
			Id:     "610611",
			CityId: "6106",
			Name:   "Seberuang",
		},
		{
			Id:     "610612",
			CityId: "6106",
			Name:   "Batang Lupar",
		},
		{
			Id:     "610613",
			CityId: "6106",
			Name:   "Empanang",
		},
		{
			Id:     "610614",
			CityId: "6106",
			Name:   "Badau",
		},
		{
			Id:     "610615",
			CityId: "6106",
			Name:   "Silat Hilir",
		},
		{
			Id:     "610616",
			CityId: "6106",
			Name:   "Silat Hulu",
		},
		{
			Id:     "610617",
			CityId: "6106",
			Name:   "Putussibau Selatan",
		},
		{
			Id:     "610618",
			CityId: "6106",
			Name:   "Kalis",
		},
		{
			Id:     "610619",
			CityId: "6106",
			Name:   "Boyan Tanjung",
		},
		{
			Id:     "610620",
			CityId: "6106",
			Name:   "Mentebah",
		},
		{
			Id:     "610621",
			CityId: "6106",
			Name:   "Pengkadan",
		},
		{
			Id:     "610622",
			CityId: "6106",
			Name:   "Suhaid",
		},
		{
			Id:     "610623",
			CityId: "6106",
			Name:   "Puring Kencana",
		},
		{
			Id:     "610701",
			CityId: "6107",
			Name:   "Sungai Raya",
		},
		{
			Id:     "610702",
			CityId: "6107",
			Name:   "Samalantan",
		},
		{
			Id:     "610703",
			CityId: "6107",
			Name:   "Ledo",
		},
		{
			Id:     "610704",
			CityId: "6107",
			Name:   "Bengkayang",
		},
		{
			Id:     "610705",
			CityId: "6107",
			Name:   "Seluas",
		},
		{
			Id:     "610706",
			CityId: "6107",
			Name:   "Sanggau Ledo",
		},
		{
			Id:     "610707",
			CityId: "6107",
			Name:   "Jagoi Babang",
		},
		{
			Id:     "610708",
			CityId: "6107",
			Name:   "Monterado",
		},
		{
			Id:     "610709",
			CityId: "6107",
			Name:   "Teriak",
		},
		{
			Id:     "610710",
			CityId: "6107",
			Name:   "Suti Semarang",
		},
		{
			Id:     "610711",
			CityId: "6107",
			Name:   "Capkala",
		},
		{
			Id:     "610712",
			CityId: "6107",
			Name:   "Siding",
		},
		{
			Id:     "610713",
			CityId: "6107",
			Name:   "Lumar",
		},
		{
			Id:     "610714",
			CityId: "6107",
			Name:   "Sungai Betung",
		},
		{
			Id:     "610715",
			CityId: "6107",
			Name:   "Sungai Raya Kepulauan",
		},
		{
			Id:     "610716",
			CityId: "6107",
			Name:   "Lembah Bawang",
		},
		{
			Id:     "610717",
			CityId: "6107",
			Name:   "Tujuh Belas",
		},
		{
			Id:     "610801",
			CityId: "6108",
			Name:   "Ngabang",
		},
		{
			Id:     "610802",
			CityId: "6108",
			Name:   "Mempawah Hulu",
		},
		{
			Id:     "610803",
			CityId: "6108",
			Name:   "Menjalin",
		},
		{
			Id:     "610804",
			CityId: "6108",
			Name:   "Mandor",
		},
		{
			Id:     "610805",
			CityId: "6108",
			Name:   "Air Besar",
		},
		{
			Id:     "610806",
			CityId: "6108",
			Name:   "Menyuke",
		},
		{
			Id:     "610807",
			CityId: "6108",
			Name:   "Sengah Temila",
		},
		{
			Id:     "610808",
			CityId: "6108",
			Name:   "Meranti",
		},
		{
			Id:     "610809",
			CityId: "6108",
			Name:   "Kuala Behe",
		},
		{
			Id:     "610810",
			CityId: "6108",
			Name:   "Sebangki",
		},
		{
			Id:     "610811",
			CityId: "6108",
			Name:   "Jelimpo",
		},
		{
			Id:     "610812",
			CityId: "6108",
			Name:   "Banyuke Hulu",
		},
		{
			Id:     "610813",
			CityId: "6108",
			Name:   "Sompak",
		},
		{
			Id:     "610901",
			CityId: "6109",
			Name:   "Sekadau Hilir",
		},
		{
			Id:     "610902",
			CityId: "6109",
			Name:   "Sekadau Hulu",
		},
		{
			Id:     "610903",
			CityId: "6109",
			Name:   "Nanga Taman",
		},
		{
			Id:     "610904",
			CityId: "6109",
			Name:   "Nanga Mahap",
		},
		{
			Id:     "610905",
			CityId: "6109",
			Name:   "Belitang Hilir",
		},
		{
			Id:     "610906",
			CityId: "6109",
			Name:   "Belitang Hulu",
		},
		{
			Id:     "610907",
			CityId: "6109",
			Name:   "Belitang",
		},
		{
			Id:     "611001",
			CityId: "6110",
			Name:   "Belimbing",
		},
		{
			Id:     "611002",
			CityId: "6110",
			Name:   "Nanga Pinoh",
		},
		{
			Id:     "950835",
			CityId: "6106",
			Name:   "Nanga Kantuk",
		},
		{
			Id:     "611003",
			CityId: "6110",
			Name:   "Ella Hilir",
		},
		{
			Id:     "611004",
			CityId: "6110",
			Name:   "Menukung",
		},
		{
			Id:     "611005",
			CityId: "6110",
			Name:   "Sayan",
		},
		{
			Id:     "611006",
			CityId: "6110",
			Name:   "Tanah Pinoh",
		},
		{
			Id:     "611007",
			CityId: "6110",
			Name:   "Sokan",
		},
		{
			Id:     "611008",
			CityId: "6110",
			Name:   "Pinoh Utara",
		},
		{
			Id:     "611009",
			CityId: "6110",
			Name:   "Pinoh Selatan",
		},
		{
			Id:     "611010",
			CityId: "6110",
			Name:   "Belimbing Hulu",
		},
		{
			Id:     "611011",
			CityId: "6110",
			Name:   "Tanah Pinoh Barat",
		},
		{
			Id:     "611101",
			CityId: "6111",
			Name:   "Sukadana",
		},
		{
			Id:     "611102",
			CityId: "6111",
			Name:   "Simpang Hilir",
		},
		{
			Id:     "611103",
			CityId: "6111",
			Name:   "Teluk Batang",
		},
		{
			Id:     "611104",
			CityId: "6111",
			Name:   "Pulau Maya",
		},
		{
			Id:     "611105",
			CityId: "6111",
			Name:   "Seponti",
		},
		{
			Id:     "611106",
			CityId: "6111",
			Name:   "Kepulauan Karimata",
		},
		{
			Id:     "611201",
			CityId: "6112",
			Name:   "Sungai Raya",
		},
		{
			Id:     "611202",
			CityId: "6112",
			Name:   "Kuala Mandor B",
		},
		{
			Id:     "611203",
			CityId: "6112",
			Name:   "Sungai Ambawang",
		},
		{
			Id:     "611204",
			CityId: "6112",
			Name:   "Terentang",
		},
		{
			Id:     "611205",
			CityId: "6112",
			Name:   "Batu Ampar",
		},
		{
			Id:     "611206",
			CityId: "6112",
			Name:   "Kubu",
		},
		{
			Id:     "611207",
			CityId: "6112",
			Name:   "Rasau Jaya",
		},
		{
			Id:     "611208",
			CityId: "6112",
			Name:   "Teluk Pakedai",
		},
		{
			Id:     "611209",
			CityId: "6112",
			Name:   "Sungai Kakap",
		},
		{
			Id:     "617101",
			CityId: "6171",
			Name:   "Pontianak Selatan",
		},
		{
			Id:     "617102",
			CityId: "6171",
			Name:   "Pontianak Timur",
		},
		{
			Id:     "617103",
			CityId: "6171",
			Name:   "Pontianak Barat",
		},
		{
			Id:     "617104",
			CityId: "6171",
			Name:   "Pontianak Utara",
		},
		{
			Id:     "617105",
			CityId: "6171",
			Name:   "Pontianak Kota",
		},
		{
			Id:     "617106",
			CityId: "6171",
			Name:   "Pontianak Tenggara",
		},
		{
			Id:     "617201",
			CityId: "6172",
			Name:   "Singkawang Tengah",
		},
		{
			Id:     "617202",
			CityId: "6172",
			Name:   "Singkawang Barat",
		},
		{
			Id:     "617203",
			CityId: "6172",
			Name:   "Singkawang Timur",
		},
		{
			Id:     "617204",
			CityId: "6172",
			Name:   "Singkawang Utara",
		},
		{
			Id:     "617205",
			CityId: "6172",
			Name:   "Singkawang Selatan",
		},
		{
			Id:     "620101",
			CityId: "6201",
			Name:   "Kumai",
		},
		{
			Id:     "620102",
			CityId: "6201",
			Name:   "Arut Selatan",
		},
		{
			Id:     "620103",
			CityId: "6201",
			Name:   "Kotawaringin Lama",
		},
		{
			Id:     "620104",
			CityId: "6201",
			Name:   "Arut Utara",
		},
		{
			Id:     "620105",
			CityId: "6201",
			Name:   "Pangkalan Lada",
		},
		{
			Id:     "620106",
			CityId: "6201",
			Name:   "Pangkalan Banteng",
		},
		{
			Id:     "620201",
			CityId: "6202",
			Name:   "Kota Besi",
		},
		{
			Id:     "620202",
			CityId: "6202",
			Name:   "Cempaga",
		},
		{
			Id:     "620203",
			CityId: "6202",
			Name:   "Mentaya Hulu",
		},
		{
			Id:     "620204",
			CityId: "6202",
			Name:   "Parenggean",
		},
		{
			Id:     "620205",
			CityId: "6202",
			Name:   "Baamang",
		},
		{
			Id:     "620206",
			CityId: "6202",
			Name:   "Mentawa Baru Ketapang",
		},
		{
			Id:     "620207",
			CityId: "6202",
			Name:   "Mentaya Hilir Utara",
		},
		{
			Id:     "620208",
			CityId: "6202",
			Name:   "Mentaya Hilir Selatan",
		},
		{
			Id:     "620209",
			CityId: "6202",
			Name:   "Pulau Hanaut",
		},
		{
			Id:     "620210",
			CityId: "6202",
			Name:   "Antang Kalang",
		},
		{
			Id:     "620211",
			CityId: "6202",
			Name:   "Teluk Sampit",
		},
		{
			Id:     "620212",
			CityId: "6202",
			Name:   "Seranau",
		},
		{
			Id:     "620213",
			CityId: "6202",
			Name:   "Cempaga Hulu",
		},
		{
			Id:     "620214",
			CityId: "6202",
			Name:   "Telawang",
		},
		{
			Id:     "620215",
			CityId: "6202",
			Name:   "Bukit Santuai",
		},
		{
			Id:     "620216",
			CityId: "6202",
			Name:   "Tualan Hulu",
		},
		{
			Id:     "620217",
			CityId: "6202",
			Name:   "Telaga Antang",
		},
		{
			Id:     "620301",
			CityId: "6203",
			Name:   "Selat",
		},
		{
			Id:     "620302",
			CityId: "6203",
			Name:   "Kapuas Hilir",
		},
		{
			Id:     "620303",
			CityId: "6203",
			Name:   "Kapuas Timur",
		},
		{
			Id:     "620304",
			CityId: "6203",
			Name:   "Kapuas Kuala",
		},
		{
			Id:     "620305",
			CityId: "6203",
			Name:   "Kapuas Barat",
		},
		{
			Id:     "620306",
			CityId: "6203",
			Name:   "Pulau Petak",
		},
		{
			Id:     "620307",
			CityId: "6203",
			Name:   "Kapuas Murung",
		},
		{
			Id:     "620308",
			CityId: "6203",
			Name:   "Basarang",
		},
		{
			Id:     "620309",
			CityId: "6203",
			Name:   "Mantangai",
		},
		{
			Id:     "620310",
			CityId: "6203",
			Name:   "Timpah",
		},
		{
			Id:     "620311",
			CityId: "6203",
			Name:   "Kapuas Tengah",
		},
		{
			Id:     "620312",
			CityId: "6203",
			Name:   "Kapuas Hulu",
		},
		{
			Id:     "620313",
			CityId: "6203",
			Name:   "Tamban Catur",
		},
		{
			Id:     "620314",
			CityId: "6203",
			Name:   "Pasak Talawang",
		},
		{
			Id:     "620315",
			CityId: "6203",
			Name:   "Mandau Talawang",
		},
		{
			Id:     "620316",
			CityId: "6203",
			Name:   "Dadahup",
		},
		{
			Id:     "620317",
			CityId: "6203",
			Name:   "Bataguh",
		},
		{
			Id:     "620401",
			CityId: "6204",
			Name:   "Jenamas",
		},
		{
			Id:     "620402",
			CityId: "6204",
			Name:   "Dusun Hilir",
		},
		{
			Id:     "620403",
			CityId: "6204",
			Name:   "Karau Kuala",
		},
		{
			Id:     "620404",
			CityId: "6204",
			Name:   "Dusun Utara",
		},
		{
			Id:     "620405",
			CityId: "6204",
			Name:   "Gn. Bintang Awai",
		},
		{
			Id:     "620406",
			CityId: "6204",
			Name:   "Dusun Selatan",
		},
		{
			Id:     "620501",
			CityId: "6205",
			Name:   "Montallat",
		},
		{
			Id:     "620502",
			CityId: "6205",
			Name:   "Gunung Timang",
		},
		{
			Id:     "620503",
			CityId: "6205",
			Name:   "Gunung Purei",
		},
		{
			Id:     "620504",
			CityId: "6205",
			Name:   "Teweh Timur",
		},
		{
			Id:     "620505",
			CityId: "6205",
			Name:   "Teweh Tengah",
		},
		{
			Id:     "620506",
			CityId: "6205",
			Name:   "Lahei",
		},
		{
			Id:     "620507",
			CityId: "6205",
			Name:   "Teweh Baru",
		},
		{
			Id:     "620508",
			CityId: "6205",
			Name:   "Teweh Selatan",
		},
		{
			Id:     "620509",
			CityId: "6205",
			Name:   "Lahei Barat",
		},
		{
			Id:     "620601",
			CityId: "6206",
			Name:   "Kamipang",
		},
		{
			Id:     "620602",
			CityId: "6206",
			Name:   "Katingan Hilir",
		},
		{
			Id:     "620603",
			CityId: "6206",
			Name:   "Tewang Sangalang Garing",
		},
		{
			Id:     "620604",
			CityId: "6206",
			Name:   "Pulau Malan",
		},
		{
			Id:     "620605",
			CityId: "6206",
			Name:   "Katingan Tengah",
		},
		{
			Id:     "620606",
			CityId: "6206",
			Name:   "Sanaman Mantikei",
		},
		{
			Id:     "620607",
			CityId: "6206",
			Name:   "Marikit",
		},
		{
			Id:     "620608",
			CityId: "6206",
			Name:   "Katingan Hulu",
		},
		{
			Id:     "620609",
			CityId: "6206",
			Name:   "Mendawai",
		},
		{
			Id:     "620610",
			CityId: "6206",
			Name:   "Katingan Kuala",
		},
		{
			Id:     "620611",
			CityId: "6206",
			Name:   "Tasik Payawan",
		},
		{
			Id:     "620612",
			CityId: "6206",
			Name:   "Petak Malai",
		},
		{
			Id:     "620613",
			CityId: "6206",
			Name:   "Bukit Raya",
		},
		{
			Id:     "620701",
			CityId: "6207",
			Name:   "Seruyan Hilir",
		},
		{
			Id:     "620702",
			CityId: "6207",
			Name:   "Seruyan Tengah",
		},
		{
			Id:     "620703",
			CityId: "6207",
			Name:   "Danau Sembuluh",
		},
		{
			Id:     "620704",
			CityId: "6207",
			Name:   "Hanau",
		},
		{
			Id:     "620705",
			CityId: "6207",
			Name:   "Seruyan Hulu",
		},
		{
			Id:     "620706",
			CityId: "6207",
			Name:   "Seruyan Hilir Timur",
		},
		{
			Id:     "620707",
			CityId: "6207",
			Name:   "Seruyan Raya",
		},
		{
			Id:     "620708",
			CityId: "6207",
			Name:   "Danau Seluluk",
		},
		{
			Id:     "620709",
			CityId: "6207",
			Name:   "Batu Ampar",
		},
		{
			Id:     "620710",
			CityId: "6207",
			Name:   "Suling Tambun",
		},
		{
			Id:     "620801",
			CityId: "6208",
			Name:   "Sukamara",
		},
		{
			Id:     "620802",
			CityId: "6208",
			Name:   "Jelai",
		},
		{
			Id:     "620803",
			CityId: "6208",
			Name:   "Balai Riam",
		},
		{
			Id:     "620804",
			CityId: "6208",
			Name:   "Pantai Lunci",
		},
		{
			Id:     "620805",
			CityId: "6208",
			Name:   "Permata Kecubung",
		},
		{
			Id:     "620901",
			CityId: "6209",
			Name:   "Lamandau",
		},
		{
			Id:     "620902",
			CityId: "6209",
			Name:   "Delang",
		},
		{
			Id:     "620903",
			CityId: "6209",
			Name:   "Bulik",
		},
		{
			Id:     "620904",
			CityId: "6209",
			Name:   "Bulik Timur",
		},
		{
			Id:     "620905",
			CityId: "6209",
			Name:   "Menthobi Raya",
		},
		{
			Id:     "620906",
			CityId: "6209",
			Name:   "Sematu Jaya",
		},
		{
			Id:     "620907",
			CityId: "6209",
			Name:   "Belantikan Raya",
		},
		{
			Id:     "620908",
			CityId: "6209",
			Name:   "Batang Kawa",
		},
		{
			Id:     "621001",
			CityId: "6210",
			Name:   "Sepang",
		},
		{
			Id:     "621002",
			CityId: "6210",
			Name:   "Kurun",
		},
		{
			Id:     "621003",
			CityId: "6210",
			Name:   "Tewah",
		},
		{
			Id:     "621004",
			CityId: "6210",
			Name:   "Kahayan Hulu Utara",
		},
		{
			Id:     "621005",
			CityId: "6210",
			Name:   "Rungan",
		},
		{
			Id:     "621006",
			CityId: "6210",
			Name:   "Manuhing",
		},
		{
			Id:     "621007",
			CityId: "6210",
			Name:   "Mihing Raya",
		},
		{
			Id:     "621008",
			CityId: "6210",
			Name:   "Damang Batu",
		},
		{
			Id:     "621009",
			CityId: "6210",
			Name:   "Miri Manasa",
		},
		{
			Id:     "621010",
			CityId: "6210",
			Name:   "Rungan Hulu",
		},
		{
			Id:     "621011",
			CityId: "6210",
			Name:   "Manuhing Raya",
		},
		{
			Id:     "621012",
			CityId: "6210",
			Name:   "Rungan Barat",
		},
		{
			Id:     "621101",
			CityId: "6211",
			Name:   "Pandih Batu",
		},
		{
			Id:     "621102",
			CityId: "6211",
			Name:   "Kahayan Kuala",
		},
		{
			Id:     "621103",
			CityId: "6211",
			Name:   "Kahayan Tengah",
		},
		{
			Id:     "621104",
			CityId: "6211",
			Name:   "Banama Tingang",
		},
		{
			Id:     "621105",
			CityId: "6211",
			Name:   "Kahayan Hilir",
		},
		{
			Id:     "621106",
			CityId: "6211",
			Name:   "Maliku",
		},
		{
			Id:     "621107",
			CityId: "6211",
			Name:   "Jabiren Raya",
		},
		{
			Id:     "621108",
			CityId: "6211",
			Name:   "Sebangau Kuala",
		},
		{
			Id:     "621201",
			CityId: "6212",
			Name:   "Murung",
		},
		{
			Id:     "621202",
			CityId: "6212",
			Name:   "Tanah Siang",
		},
		{
			Id:     "621203",
			CityId: "6212",
			Name:   "Laung Tuhup",
		},
		{
			Id:     "621204",
			CityId: "6212",
			Name:   "Permata Intan",
		},
		{
			Id:     "621205",
			CityId: "6212",
			Name:   "Sumber Barito",
		},
		{
			Id:     "621206",
			CityId: "6212",
			Name:   "Barito Tuhup Raya",
		},
		{
			Id:     "621207",
			CityId: "6212",
			Name:   "Tanah Siang Selatan",
		},
		{
			Id:     "621208",
			CityId: "6212",
			Name:   "Sungai Babuat",
		},
		{
			Id:     "621209",
			CityId: "6212",
			Name:   "Seribu Riam",
		},
		{
			Id:     "621210",
			CityId: "6212",
			Name:   "Uut Murung",
		},
		{
			Id:     "621301",
			CityId: "6213",
			Name:   "Dusun Timur",
		},
		{
			Id:     "621302",
			CityId: "6213",
			Name:   "Banua Lima",
		},
		{
			Id:     "621303",
			CityId: "6213",
			Name:   "Patangkep Tutui",
		},
		{
			Id:     "621304",
			CityId: "6213",
			Name:   "Awang",
		},
		{
			Id:     "621305",
			CityId: "6213",
			Name:   "Dusun Tengah",
		},
		{
			Id:     "621306",
			CityId: "6213",
			Name:   "Pematang Karau",
		},
		{
			Id:     "621307",
			CityId: "6213",
			Name:   "Paju Epat",
		},
		{
			Id:     "621308",
			CityId: "6213",
			Name:   "Raren Batuah",
		},
		{
			Id:     "621309",
			CityId: "6213",
			Name:   "Paku",
		},
		{
			Id:     "621310",
			CityId: "6213",
			Name:   "Karusen Janang",
		},
		{
			Id:     "627101",
			CityId: "6271",
			Name:   "Pahandut",
		},
		{
			Id:     "627102",
			CityId: "6271",
			Name:   "Bukit Batu",
		},
		{
			Id:     "627103",
			CityId: "6271",
			Name:   "Jekan Raya",
		},
		{
			Id:     "627104",
			CityId: "6271",
			Name:   "Sabangau",
		},
		{
			Id:     "627105",
			CityId: "6271",
			Name:   "Rakumpit",
		},
		{
			Id:     "630101",
			CityId: "6301",
			Name:   "Takisung",
		},
		{
			Id:     "630102",
			CityId: "6301",
			Name:   "Jorong",
		},
		{
			Id:     "630103",
			CityId: "6301",
			Name:   "Pelaihari",
		},
		{
			Id:     "630104",
			CityId: "6301",
			Name:   "Kurau",
		},
		{
			Id:     "630105",
			CityId: "6301",
			Name:   "Bati Bati",
		},
		{
			Id:     "630106",
			CityId: "6301",
			Name:   "Panyipatan",
		},
		{
			Id:     "630107",
			CityId: "6301",
			Name:   "Kintap",
		},
		{
			Id:     "630108",
			CityId: "6301",
			Name:   "Tambang Ulang",
		},
		{
			Id:     "630109",
			CityId: "6301",
			Name:   "Batu Ampar",
		},
		{
			Id:     "630110",
			CityId: "6301",
			Name:   "Bajuin",
		},
		{
			Id:     "630111",
			CityId: "6301",
			Name:   "Bumi Makmur",
		},
		{
			Id:     "630201",
			CityId: "6302",
			Name:   "Pulau Sembilan",
		},
		{
			Id:     "630202",
			CityId: "6302",
			Name:   "Pulau Laut Barat",
		},
		{
			Id:     "630203",
			CityId: "6302",
			Name:   "Pulau Laut Selatan",
		},
		{
			Id:     "630204",
			CityId: "6302",
			Name:   "Pulau Laut Timur",
		},
		{
			Id:     "630205",
			CityId: "6302",
			Name:   "Pulau Sebuku",
		},
		{
			Id:     "630206",
			CityId: "6302",
			Name:   "Pulaulaut Utara",
		},
		{
			Id:     "630207",
			CityId: "6302",
			Name:   "Kelumpang Selatan",
		},
		{
			Id:     "630208",
			CityId: "6302",
			Name:   "Kelumpang Hulu",
		},
		{
			Id:     "630209",
			CityId: "6302",
			Name:   "Kelumpang Tengah",
		},
		{
			Id:     "630210",
			CityId: "6302",
			Name:   "Kelumpang Utara",
		},
		{
			Id:     "630211",
			CityId: "6302",
			Name:   "Pamukan Selatan",
		},
		{
			Id:     "630212",
			CityId: "6302",
			Name:   "Sampanahan",
		},
		{
			Id:     "630213",
			CityId: "6302",
			Name:   "Pamukan Utara",
		},
		{
			Id:     "630214",
			CityId: "6302",
			Name:   "Hampang",
		},
		{
			Id:     "630215",
			CityId: "6302",
			Name:   "Sungai Durian",
		},
		{
			Id:     "630216",
			CityId: "6302",
			Name:   "Pulau Laut Tengah",
		},
		{
			Id:     "630217",
			CityId: "6302",
			Name:   "Kelumpang Hilir",
		},
		{
			Id:     "630218",
			CityId: "6302",
			Name:   "Kelumpang Barat",
		},
		{
			Id:     "630219",
			CityId: "6302",
			Name:   "Pamukan Barat",
		},
		{
			Id:     "630220",
			CityId: "6302",
			Name:   "Pulau Laut Kepulauan",
		},
		{
			Id:     "630221",
			CityId: "6302",
			Name:   "Pulau Laut Tanjung Selayar",
		},
		{
			Id:     "630222",
			CityId: "6302",
			Name:   "Pulaulaut Sigam",
		},
		{
			Id:     "630301",
			CityId: "6303",
			Name:   "Aluh Aluh",
		},
		{
			Id:     "630302",
			CityId: "6303",
			Name:   "Kertak Hanyar",
		},
		{
			Id:     "630303",
			CityId: "6303",
			Name:   "Gambut",
		},
		{
			Id:     "630304",
			CityId: "6303",
			Name:   "Sungai Tabuk",
		},
		{
			Id:     "630305",
			CityId: "6303",
			Name:   "Martapura",
		},
		{
			Id:     "630306",
			CityId: "6303",
			Name:   "Karang Intan",
		},
		{
			Id:     "630307",
			CityId: "6303",
			Name:   "Astambul",
		},
		{
			Id:     "630308",
			CityId: "6303",
			Name:   "Simpang Empat",
		},
		{
			Id:     "630309",
			CityId: "6303",
			Name:   "Pengaron",
		},
		{
			Id:     "630310",
			CityId: "6303",
			Name:   "Sungai Pinang",
		},
		{
			Id:     "630311",
			CityId: "6303",
			Name:   "Aranio",
		},
		{
			Id:     "630312",
			CityId: "6303",
			Name:   "Mataraman",
		},
		{
			Id:     "630313",
			CityId: "6303",
			Name:   "Beruntung Baru",
		},
		{
			Id:     "630314",
			CityId: "6303",
			Name:   "Martapura Barat",
		},
		{
			Id:     "630315",
			CityId: "6303",
			Name:   "Martapura Timur",
		},
		{
			Id:     "630316",
			CityId: "6303",
			Name:   "Sambung Makmur",
		},
		{
			Id:     "630317",
			CityId: "6303",
			Name:   "Paramasan",
		},
		{
			Id:     "630318",
			CityId: "6303",
			Name:   "Telaga Bauntung",
		},
		{
			Id:     "630319",
			CityId: "6303",
			Name:   "Tatah Makmur",
		},
		{
			Id:     "630320",
			CityId: "6303",
			Name:   "Cintapuri Darussalam",
		},
		{
			Id:     "630401",
			CityId: "6304",
			Name:   "Tabunganen",
		},
		{
			Id:     "630402",
			CityId: "6304",
			Name:   "Tamban",
		},
		{
			Id:     "630403",
			CityId: "6304",
			Name:   "Anjir Pasar",
		},
		{
			Id:     "630404",
			CityId: "6304",
			Name:   "Anjir Muara",
		},
		{
			Id:     "630405",
			CityId: "6304",
			Name:   "Alalak",
		},
		{
			Id:     "630406",
			CityId: "6304",
			Name:   "Mandastana",
		},
		{
			Id:     "630407",
			CityId: "6304",
			Name:   "Rantau Badauh",
		},
		{
			Id:     "630408",
			CityId: "6304",
			Name:   "Belawang",
		},
		{
			Id:     "630409",
			CityId: "6304",
			Name:   "Cerbon",
		},
		{
			Id:     "630410",
			CityId: "6304",
			Name:   "Bakumpai",
		},
		{
			Id:     "630411",
			CityId: "6304",
			Name:   "Kuripan",
		},
		{
			Id:     "630412",
			CityId: "6304",
			Name:   "Tabukan",
		},
		{
			Id:     "630413",
			CityId: "6304",
			Name:   "Mekarsari",
		},
		{
			Id:     "630414",
			CityId: "6304",
			Name:   "Barambai",
		},
		{
			Id:     "630415",
			CityId: "6304",
			Name:   "Marabahan",
		},
		{
			Id:     "630416",
			CityId: "6304",
			Name:   "Wanaraya",
		},
		{
			Id:     "630417",
			CityId: "6304",
			Name:   "Jejangkit",
		},
		{
			Id:     "630501",
			CityId: "6305",
			Name:   "Binuang",
		},
		{
			Id:     "630502",
			CityId: "6305",
			Name:   "Tapin Selatan",
		},
		{
			Id:     "630503",
			CityId: "6305",
			Name:   "Tapin Tengah",
		},
		{
			Id:     "630504",
			CityId: "6305",
			Name:   "Tapin Utara",
		},
		{
			Id:     "630505",
			CityId: "6305",
			Name:   "Candi Laras Selatan",
		},
		{
			Id:     "630506",
			CityId: "6305",
			Name:   "Candi Laras Utara",
		},
		{
			Id:     "630507",
			CityId: "6305",
			Name:   "Bakarangan",
		},
		{
			Id:     "630508",
			CityId: "6305",
			Name:   "Piani",
		},
		{
			Id:     "630509",
			CityId: "6305",
			Name:   "Bungur",
		},
		{
			Id:     "630510",
			CityId: "6305",
			Name:   "Lokpaikat",
		},
		{
			Id:     "630511",
			CityId: "6305",
			Name:   "Salam Babaris",
		},
		{
			Id:     "630512",
			CityId: "6305",
			Name:   "Hatungun",
		},
		{
			Id:     "630601",
			CityId: "6306",
			Name:   "Sungai Raya",
		},
		{
			Id:     "630602",
			CityId: "6306",
			Name:   "Padang Batung",
		},
		{
			Id:     "630603",
			CityId: "6306",
			Name:   "Telaga Langsat",
		},
		{
			Id:     "630604",
			CityId: "6306",
			Name:   "Angkinang",
		},
		{
			Id:     "630605",
			CityId: "6306",
			Name:   "Kandangan",
		},
		{
			Id:     "630606",
			CityId: "6306",
			Name:   "Simpur",
		},
		{
			Id:     "630607",
			CityId: "6306",
			Name:   "Daha Selatan",
		},
		{
			Id:     "630608",
			CityId: "6306",
			Name:   "Daha Utara",
		},
		{
			Id:     "630609",
			CityId: "6306",
			Name:   "Kalumpang",
		},
		{
			Id:     "630610",
			CityId: "6306",
			Name:   "Loksado",
		},
		{
			Id:     "630611",
			CityId: "6306",
			Name:   "Daha Barat",
		},
		{
			Id:     "630701",
			CityId: "6307",
			Name:   "Haruyan",
		},
		{
			Id:     "630702",
			CityId: "6307",
			Name:   "Batu Benawa",
		},
		{
			Id:     "630703",
			CityId: "6307",
			Name:   "Labuan Amas Selatan",
		},
		{
			Id:     "630704",
			CityId: "6307",
			Name:   "Labuan Amas Utara",
		},
		{
			Id:     "630705",
			CityId: "6307",
			Name:   "Pandawan",
		},
		{
			Id:     "630706",
			CityId: "6307",
			Name:   "Barabai",
		},
		{
			Id:     "630707",
			CityId: "6307",
			Name:   "Batang Alai Selatan",
		},
		{
			Id:     "630708",
			CityId: "6307",
			Name:   "Batang Alai Utara",
		},
		{
			Id:     "630709",
			CityId: "6307",
			Name:   "Hantakan",
		},
		{
			Id:     "630710",
			CityId: "6307",
			Name:   "Batang Alai Timur",
		},
		{
			Id:     "630711",
			CityId: "6307",
			Name:   "Limpasu",
		},
		{
			Id:     "630801",
			CityId: "6308",
			Name:   "Danau Panggang",
		},
		{
			Id:     "630802",
			CityId: "6308",
			Name:   "Babirik",
		},
		{
			Id:     "630803",
			CityId: "6308",
			Name:   "Sungai Pandan",
		},
		{
			Id:     "630804",
			CityId: "6308",
			Name:   "Amuntai Selatan",
		},
		{
			Id:     "630805",
			CityId: "6308",
			Name:   "Amuntai Tengah",
		},
		{
			Id:     "630806",
			CityId: "6308",
			Name:   "Amuntai Utara",
		},
		{
			Id:     "630807",
			CityId: "6308",
			Name:   "Banjang",
		},
		{
			Id:     "630808",
			CityId: "6308",
			Name:   "Haur Gading",
		},
		{
			Id:     "630809",
			CityId: "6308",
			Name:   "Paminggir",
		},
		{
			Id:     "630810",
			CityId: "6308",
			Name:   "Sungai Tabukan",
		},
		{
			Id:     "630901",
			CityId: "6309",
			Name:   "Banua Lawas",
		},
		{
			Id:     "630902",
			CityId: "6309",
			Name:   "Kelua",
		},
		{
			Id:     "630903",
			CityId: "6309",
			Name:   "Tanta",
		},
		{
			Id:     "630904",
			CityId: "6309",
			Name:   "Tanjung",
		},
		{
			Id:     "630905",
			CityId: "6309",
			Name:   "Haruai",
		},
		{
			Id:     "630906",
			CityId: "6309",
			Name:   "Murung Pudak",
		},
		{
			Id:     "630907",
			CityId: "6309",
			Name:   "Muara Uya",
		},
		{
			Id:     "630908",
			CityId: "6309",
			Name:   "Muara Harus",
		},
		{
			Id:     "630909",
			CityId: "6309",
			Name:   "Pugaan",
		},
		{
			Id:     "630910",
			CityId: "6309",
			Name:   "Upau",
		},
		{
			Id:     "630911",
			CityId: "6309",
			Name:   "Jaro",
		},
		{
			Id:     "630912",
			CityId: "6309",
			Name:   "Bintang Ara",
		},
		{
			Id:     "631001",
			CityId: "6310",
			Name:   "Batu Licin",
		},
		{
			Id:     "631002",
			CityId: "6310",
			Name:   "Kusan Hilir",
		},
		{
			Id:     "631003",
			CityId: "6310",
			Name:   "Sungai Loban",
		},
		{
			Id:     "631004",
			CityId: "6310",
			Name:   "Satui",
		},
		{
			Id:     "631005",
			CityId: "6310",
			Name:   "Kusan Hulu",
		},
		{
			Id:     "631006",
			CityId: "6310",
			Name:   "Simpang Empat",
		},
		{
			Id:     "631007",
			CityId: "6310",
			Name:   "Karang Bintang",
		},
		{
			Id:     "631008",
			CityId: "6310",
			Name:   "Mantewe",
		},
		{
			Id:     "631009",
			CityId: "6310",
			Name:   "Angsana",
		},
		{
			Id:     "631010",
			CityId: "6310",
			Name:   "Kuranji",
		},
		{
			Id:     "631011",
			CityId: "6310",
			Name:   "Kusan Tengah",
		},
		{
			Id:     "631012",
			CityId: "6310",
			Name:   "Teluk Kepayang",
		},
		{
			Id:     "631101",
			CityId: "6311",
			Name:   "Juai",
		},
		{
			Id:     "631102",
			CityId: "6311",
			Name:   "Halong",
		},
		{
			Id:     "631103",
			CityId: "6311",
			Name:   "Awayan",
		},
		{
			Id:     "631104",
			CityId: "6311",
			Name:   "Batu Mandi",
		},
		{
			Id:     "631105",
			CityId: "6311",
			Name:   "Lampihong",
		},
		{
			Id:     "631106",
			CityId: "6311",
			Name:   "Paringin",
		},
		{
			Id:     "631107",
			CityId: "6311",
			Name:   "Paringin Selatan",
		},
		{
			Id:     "631108",
			CityId: "6311",
			Name:   "Tebing Tinggi",
		},
		{
			Id:     "637101",
			CityId: "6371",
			Name:   "Banjarmasin Selatan",
		},
		{
			Id:     "637102",
			CityId: "6371",
			Name:   "Banjarmasin Timur",
		},
		{
			Id:     "637103",
			CityId: "6371",
			Name:   "Banjarmasin Barat",
		},
		{
			Id:     "637104",
			CityId: "6371",
			Name:   "Banjarmasin Utara",
		},
		{
			Id:     "637105",
			CityId: "6371",
			Name:   "Banjarmasin Tengah",
		},
		{
			Id:     "637202",
			CityId: "6372",
			Name:   "Landasan Ulin",
		},
		{
			Id:     "637203",
			CityId: "6372",
			Name:   "Cempaka",
		},
		{
			Id:     "637204",
			CityId: "6372",
			Name:   "Banjarbaru Utara",
		},
		{
			Id:     "637205",
			CityId: "6372",
			Name:   "Banjarbaru Selatan",
		},
		{
			Id:     "637206",
			CityId: "6372",
			Name:   "Liang Anggang",
		},
		{
			Id:     "640101",
			CityId: "6401",
			Name:   "Batu Sopang",
		},
		{
			Id:     "640102",
			CityId: "6401",
			Name:   "Tanjung Harapan",
		},
		{
			Id:     "640103",
			CityId: "6401",
			Name:   "Paser Belengkong",
		},
		{
			Id:     "640104",
			CityId: "6401",
			Name:   "Tanah Grogot",
		},
		{
			Id:     "640105",
			CityId: "6401",
			Name:   "Kuaro",
		},
		{
			Id:     "640106",
			CityId: "6401",
			Name:   "Long Ikis",
		},
		{
			Id:     "640107",
			CityId: "6401",
			Name:   "Muara Komam",
		},
		{
			Id:     "640108",
			CityId: "6401",
			Name:   "Long Kali",
		},
		{
			Id:     "640109",
			CityId: "6401",
			Name:   "Batu Engau",
		},
		{
			Id:     "640110",
			CityId: "6401",
			Name:   "Muara Samu",
		},
		{
			Id:     "640201",
			CityId: "6402",
			Name:   "Muara Muntai",
		},
		{
			Id:     "640202",
			CityId: "6402",
			Name:   "Loa Kulu",
		},
		{
			Id:     "640203",
			CityId: "6402",
			Name:   "Loa Janan",
		},
		{
			Id:     "640204",
			CityId: "6402",
			Name:   "Anggana",
		},
		{
			Id:     "640205",
			CityId: "6402",
			Name:   "Muara Badak",
		},
		{
			Id:     "640206",
			CityId: "6402",
			Name:   "Tenggarong",
		},
		{
			Id:     "640207",
			CityId: "6402",
			Name:   "Sebulu",
		},
		{
			Id:     "640208",
			CityId: "6402",
			Name:   "Kota Bangun",
		},
		{
			Id:     "640209",
			CityId: "6402",
			Name:   "Kenohan",
		},
		{
			Id:     "640210",
			CityId: "6402",
			Name:   "Kembang Janggut",
		},
		{
			Id:     "640211",
			CityId: "6402",
			Name:   "Muara Kaman",
		},
		{
			Id:     "640212",
			CityId: "6402",
			Name:   "Tabang",
		},
		{
			Id:     "640213",
			CityId: "6402",
			Name:   "Samboja",
		},
		{
			Id:     "640214",
			CityId: "6402",
			Name:   "Muara Jawa",
		},
		{
			Id:     "640215",
			CityId: "6402",
			Name:   "Sanga Sanga",
		},
		{
			Id:     "640216",
			CityId: "6402",
			Name:   "Tenggarong Seberang",
		},
		{
			Id:     "640217",
			CityId: "6402",
			Name:   "Marang Kayu",
		},
		{
			Id:     "640218",
			CityId: "6402",
			Name:   "Muara Wis",
		},
		{
			Id:     "640219",
			CityId: "6402",
			Name:   "Kota Bangun Darat",
		},
		{
			Id:     "640220",
			CityId: "6402",
			Name:   "Semboja Barat",
		},
		{
			Id:     "640301",
			CityId: "6403",
			Name:   "Kelay",
		},
		{
			Id:     "640302",
			CityId: "6403",
			Name:   "Talisayan",
		},
		{
			Id:     "640303",
			CityId: "6403",
			Name:   "Sambaliung",
		},
		{
			Id:     "640304",
			CityId: "6403",
			Name:   "Segah",
		},
		{
			Id:     "640305",
			CityId: "6403",
			Name:   "Tanjung Redeb",
		},
		{
			Id:     "640306",
			CityId: "6403",
			Name:   "Gunung Tabur",
		},
		{
			Id:     "640307",
			CityId: "6403",
			Name:   "Pulau Derawan",
		},
		{
			Id:     "640308",
			CityId: "6403",
			Name:   "Biduk-Biduk",
		},
		{
			Id:     "640309",
			CityId: "6403",
			Name:   "Teluk Bayur",
		},
		{
			Id:     "640310",
			CityId: "6403",
			Name:   "Tabalar",
		},
		{
			Id:     "640311",
			CityId: "6403",
			Name:   "Maratua",
		},
		{
			Id:     "640312",
			CityId: "6403",
			Name:   "Batu Putih",
		},
		{
			Id:     "640313",
			CityId: "6403",
			Name:   "Biatan",
		},
		{
			Id:     "640705",
			CityId: "6407",
			Name:   "Long Iram",
		},
		{
			Id:     "640706",
			CityId: "6407",
			Name:   "Melak",
		},
		{
			Id:     "640707",
			CityId: "6407",
			Name:   "Barong Tongkok",
		},
		{
			Id:     "640708",
			CityId: "6407",
			Name:   "Damai",
		},
		{
			Id:     "640709",
			CityId: "6407",
			Name:   "Muara Lawa",
		},
		{
			Id:     "640710",
			CityId: "6407",
			Name:   "Muara Pahu",
		},
		{
			Id:     "640711",
			CityId: "6407",
			Name:   "Jempang",
		},
		{
			Id:     "640712",
			CityId: "6407",
			Name:   "Bongan",
		},
		{
			Id:     "640713",
			CityId: "6407",
			Name:   "Penyinggahan",
		},
		{
			Id:     "640714",
			CityId: "6407",
			Name:   "Bentian Besar",
		},
		{
			Id:     "640715",
			CityId: "6407",
			Name:   "Linggang Bigung",
		},
		{
			Id:     "640716",
			CityId: "6407",
			Name:   "Nyuatan",
		},
		{
			Id:     "640717",
			CityId: "6407",
			Name:   "Siluq Ngurai",
		},
		{
			Id:     "640718",
			CityId: "6407",
			Name:   "Mook Manaar Bulatn",
		},
		{
			Id:     "640719",
			CityId: "6407",
			Name:   "Tering",
		},
		{
			Id:     "640720",
			CityId: "6407",
			Name:   "Sekolaq Darat",
		},
		{
			Id:     "640801",
			CityId: "6408",
			Name:   "Muara Ancalong",
		},
		{
			Id:     "640802",
			CityId: "6408",
			Name:   "Muara Wahau",
		},
		{
			Id:     "640803",
			CityId: "6408",
			Name:   "Muara Bengkal",
		},
		{
			Id:     "640804",
			CityId: "6408",
			Name:   "Sangatta Utara",
		},
		{
			Id:     "640805",
			CityId: "6408",
			Name:   "Sangkulirang",
		},
		{
			Id:     "640806",
			CityId: "6408",
			Name:   "Busang",
		},
		{
			Id:     "640807",
			CityId: "6408",
			Name:   "Telen",
		},
		{
			Id:     "640808",
			CityId: "6408",
			Name:   "Kombeng",
		},
		{
			Id:     "640809",
			CityId: "6408",
			Name:   "Bengalon",
		},
		{
			Id:     "640810",
			CityId: "6408",
			Name:   "Kaliorang",
		},
		{
			Id:     "640811",
			CityId: "6408",
			Name:   "Sandaran",
		},
		{
			Id:     "640812",
			CityId: "6408",
			Name:   "Sangatta Selatan",
		},
		{
			Id:     "640813",
			CityId: "6408",
			Name:   "Teluk Pandan",
		},
		{
			Id:     "640814",
			CityId: "6408",
			Name:   "Rantau Pulung",
		},
		{
			Id:     "640815",
			CityId: "6408",
			Name:   "Kaubun",
		},
		{
			Id:     "640816",
			CityId: "6408",
			Name:   "Karangan",
		},
		{
			Id:     "640817",
			CityId: "6408",
			Name:   "Batu Ampar",
		},
		{
			Id:     "640818",
			CityId: "6408",
			Name:   "Long Mesangat",
		},
		{
			Id:     "640901",
			CityId: "6409",
			Name:   "Penajam",
		},
		{
			Id:     "640902",
			CityId: "6409",
			Name:   "Waru",
		},
		{
			Id:     "640903",
			CityId: "6409",
			Name:   "Babulu",
		},
		{
			Id:     "640904",
			CityId: "6409",
			Name:   "Sepaku",
		},
		{
			Id:     "641101",
			CityId: "6411",
			Name:   "Long Bagun",
		},
		{
			Id:     "641102",
			CityId: "6411",
			Name:   "Long Hubung",
		},
		{
			Id:     "641103",
			CityId: "6411",
			Name:   "Laham",
		},
		{
			Id:     "641104",
			CityId: "6411",
			Name:   "Long Apari",
		},
		{
			Id:     "641105",
			CityId: "6411",
			Name:   "Long Pahangai",
		},
		{
			Id:     "647101",
			CityId: "6471",
			Name:   "Balikpapan Timur",
		},
		{
			Id:     "647102",
			CityId: "6471",
			Name:   "Balikpapan Barat",
		},
		{
			Id:     "647103",
			CityId: "6471",
			Name:   "Balikpapan Utara",
		},
		{
			Id:     "647104",
			CityId: "6471",
			Name:   "Balikpapan Tengah",
		},
		{
			Id:     "647105",
			CityId: "6471",
			Name:   "Balikpapan Selatan",
		},
		{
			Id:     "647106",
			CityId: "6471",
			Name:   "Balikpapan Kota",
		},
		{
			Id:     "647201",
			CityId: "6472",
			Name:   "Palaran",
		},
		{
			Id:     "647202",
			CityId: "6472",
			Name:   "Samarinda Seberang",
		},
		{
			Id:     "647203",
			CityId: "6472",
			Name:   "Samarinda Ulu",
		},
		{
			Id:     "647204",
			CityId: "6472",
			Name:   "Samarinda Ilir",
		},
		{
			Id:     "647205",
			CityId: "6472",
			Name:   "Samarinda Utara",
		},
		{
			Id:     "647206",
			CityId: "6472",
			Name:   "Sungai Kunjang",
		},
		{
			Id:     "647207",
			CityId: "6472",
			Name:   "Sambutan",
		},
		{
			Id:     "647208",
			CityId: "6472",
			Name:   "Sungai Pinang",
		},
		{
			Id:     "647209",
			CityId: "6472",
			Name:   "Samarinda Kota",
		},
		{
			Id:     "647210",
			CityId: "6472",
			Name:   "Loa Janan Ilir",
		},
		{
			Id:     "647401",
			CityId: "6474",
			Name:   "Bontang Utara",
		},
		{
			Id:     "647402",
			CityId: "6474",
			Name:   "Bontang Selatan",
		},
		{
			Id:     "647403",
			CityId: "6474",
			Name:   "Bontang Barat",
		},
		{
			Id:     "650101",
			CityId: "6501",
			Name:   "Tanjung Palas",
		},
		{
			Id:     "650102",
			CityId: "6501",
			Name:   "Tanjung Palas Barat",
		},
		{
			Id:     "650103",
			CityId: "6501",
			Name:   "Tanjung Palas Utara",
		},
		{
			Id:     "650104",
			CityId: "6501",
			Name:   "Tanjung Palas Timur",
		},
		{
			Id:     "650105",
			CityId: "6501",
			Name:   "Tanjung Selor",
		},
		{
			Id:     "650106",
			CityId: "6501",
			Name:   "Tanjung Palas Tengah",
		},
		{
			Id:     "650107",
			CityId: "6501",
			Name:   "Peso",
		},
		{
			Id:     "650108",
			CityId: "6501",
			Name:   "Peso Hilir",
		},
		{
			Id:     "650109",
			CityId: "6501",
			Name:   "Sekatak",
		},
		{
			Id:     "650110",
			CityId: "6501",
			Name:   "Bunyu",
		},
		{
			Id:     "650201",
			CityId: "6502",
			Name:   "Mentarang",
		},
		{
			Id:     "650202",
			CityId: "6502",
			Name:   "Malinau Kota",
		},
		{
			Id:     "650203",
			CityId: "6502",
			Name:   "Pujungan",
		},
		{
			Id:     "650204",
			CityId: "6502",
			Name:   "Kayan Hilir",
		},
		{
			Id:     "650205",
			CityId: "6502",
			Name:   "Kayan Hulu",
		},
		{
			Id:     "650206",
			CityId: "6502",
			Name:   "Malinau Selatan",
		},
		{
			Id:     "650207",
			CityId: "6502",
			Name:   "Malinau Utara",
		},
		{
			Id:     "650208",
			CityId: "6502",
			Name:   "Malinau Barat",
		},
		{
			Id:     "650209",
			CityId: "6502",
			Name:   "Sungai Boh",
		},
		{
			Id:     "650210",
			CityId: "6502",
			Name:   "Kayan Selatan",
		},
		{
			Id:     "650211",
			CityId: "6502",
			Name:   "Bahau Hulu",
		},
		{
			Id:     "650212",
			CityId: "6502",
			Name:   "Mentarang Hulu",
		},
		{
			Id:     "650213",
			CityId: "6502",
			Name:   "Malinau Selatan Hilir",
		},
		{
			Id:     "650214",
			CityId: "6502",
			Name:   "Malinau Selatan Hulu",
		},
		{
			Id:     "650215",
			CityId: "6502",
			Name:   "Sungai Tubu",
		},
		{
			Id:     "650301",
			CityId: "6503",
			Name:   "Sebatik",
		},
		{
			Id:     "650302",
			CityId: "6503",
			Name:   "Nunukan",
		},
		{
			Id:     "650303",
			CityId: "6503",
			Name:   "Sembakung",
		},
		{
			Id:     "650304",
			CityId: "6503",
			Name:   "Lumbis",
		},
		{
			Id:     "650305",
			CityId: "6503",
			Name:   "Krayan",
		},
		{
			Id:     "650306",
			CityId: "6503",
			Name:   "Sebuku",
		},
		{
			Id:     "650307",
			CityId: "6503",
			Name:   "Krayan Selatan",
		},
		{
			Id:     "650308",
			CityId: "6503",
			Name:   "Sebatik Barat",
		},
		{
			Id:     "650309",
			CityId: "6503",
			Name:   "Nunukan Selatan",
		},
		{
			Id:     "650310",
			CityId: "6503",
			Name:   "Sebatik Timur",
		},
		{
			Id:     "650311",
			CityId: "6503",
			Name:   "Sebatik Utara",
		},
		{
			Id:     "650312",
			CityId: "6503",
			Name:   "Sebatik Tengah",
		},
		{
			Id:     "650313",
			CityId: "6503",
			Name:   "Sei Menggaris",
		},
		{
			Id:     "650314",
			CityId: "6503",
			Name:   "Tulin Onsoi",
		},
		{
			Id:     "650315",
			CityId: "6503",
			Name:   "Lumbis Ogong",
		},
		{
			Id:     "650316",
			CityId: "6503",
			Name:   "Sembakung Atulai",
		},
		{
			Id:     "650317",
			CityId: "6503",
			Name:   "Krayan Tengah",
		},
		{
			Id:     "650318",
			CityId: "6503",
			Name:   "Krayan Timur",
		},
		{
			Id:     "650319",
			CityId: "6503",
			Name:   "Krayan Barat",
		},
		{
			Id:     "650320",
			CityId: "6503",
			Name:   "Lumbis Pansiangan",
		},
		{
			Id:     "650321",
			CityId: "6503",
			Name:   "Lumbis Hulu",
		},
		{
			Id:     "650401",
			CityId: "6504",
			Name:   "Sesayap",
		},
		{
			Id:     "650402",
			CityId: "6504",
			Name:   "Sesayap Hilir",
		},
		{
			Id:     "650403",
			CityId: "6504",
			Name:   "Tana Lia",
		},
		{
			Id:     "650404",
			CityId: "6504",
			Name:   "Betayau",
		},
		{
			Id:     "650405",
			CityId: "6504",
			Name:   "Muruk Rian",
		},
		{
			Id:     "657101",
			CityId: "6571",
			Name:   "Tarakan Barat",
		},
		{
			Id:     "657102",
			CityId: "6571",
			Name:   "Tarakan Tengah",
		},
		{
			Id:     "657103",
			CityId: "6571",
			Name:   "Tarakan Timur",
		},
		{
			Id:     "657104",
			CityId: "6571",
			Name:   "Tarakan Utara",
		},
		{
			Id:     "710105",
			CityId: "7101",
			Name:   "Sang Tombolang",
		},
		{
			Id:     "710109",
			CityId: "7101",
			Name:   "Dumoga Barat",
		},
		{
			Id:     "710110",
			CityId: "7101",
			Name:   "Dumoga Timur",
		},
		{
			Id:     "710111",
			CityId: "7101",
			Name:   "Dumoga Utara",
		},
		{
			Id:     "710112",
			CityId: "7101",
			Name:   "Lolak",
		},
		{
			Id:     "710113",
			CityId: "7101",
			Name:   "Bolaang",
		},
		{
			Id:     "710114",
			CityId: "7101",
			Name:   "Lolayan",
		},
		{
			Id:     "710119",
			CityId: "7101",
			Name:   "Passi Barat",
		},
		{
			Id:     "710120",
			CityId: "7101",
			Name:   "Poigar",
		},
		{
			Id:     "710122",
			CityId: "7101",
			Name:   "Passi Timur",
		},
		{
			Id:     "710131",
			CityId: "7101",
			Name:   "Bolaang Timur",
		},
		{
			Id:     "710132",
			CityId: "7101",
			Name:   "Bilalang",
		},
		{
			Id:     "710133",
			CityId: "7101",
			Name:   "Dumoga",
		},
		{
			Id:     "710134",
			CityId: "7101",
			Name:   "Dumoga Tenggara",
		},
		{
			Id:     "710135",
			CityId: "7101",
			Name:   "Dumoga Tengah",
		},
		{
			Id:     "710201",
			CityId: "7102",
			Name:   "Tondano Barat",
		},
		{
			Id:     "710202",
			CityId: "7102",
			Name:   "Tondano Timur",
		},
		{
			Id:     "710203",
			CityId: "7102",
			Name:   "Eris",
		},
		{
			Id:     "710204",
			CityId: "7102",
			Name:   "Kombi",
		},
		{
			Id:     "710205",
			CityId: "7102",
			Name:   "Lembean Timur",
		},
		{
			Id:     "710206",
			CityId: "7102",
			Name:   "Kakas",
		},
		{
			Id:     "710207",
			CityId: "7102",
			Name:   "Tompaso",
		},
		{
			Id:     "710208",
			CityId: "7102",
			Name:   "Remboken",
		},
		{
			Id:     "710209",
			CityId: "7102",
			Name:   "Langowan Timur",
		},
		{
			Id:     "710210",
			CityId: "7102",
			Name:   "Langowan Barat",
		},
		{
			Id:     "710211",
			CityId: "7102",
			Name:   "Sonder",
		},
		{
			Id:     "710212",
			CityId: "7102",
			Name:   "Kawangkoan",
		},
		{
			Id:     "710213",
			CityId: "7102",
			Name:   "Pineleng",
		},
		{
			Id:     "710214",
			CityId: "7102",
			Name:   "Tombulu",
		},
		{
			Id:     "710215",
			CityId: "7102",
			Name:   "Tombariri",
		},
		{
			Id:     "710216",
			CityId: "7102",
			Name:   "Tondano Utara",
		},
		{
			Id:     "710217",
			CityId: "7102",
			Name:   "Langowan Selatan",
		},
		{
			Id:     "710218",
			CityId: "7102",
			Name:   "Tondano Selatan",
		},
		{
			Id:     "710219",
			CityId: "7102",
			Name:   "Langowan Utara",
		},
		{
			Id:     "710220",
			CityId: "7102",
			Name:   "Kakas Barat",
		},
		{
			Id:     "710221",
			CityId: "7102",
			Name:   "Kawangkoan Utara",
		},
		{
			Id:     "710222",
			CityId: "7102",
			Name:   "Kawangkoan Barat",
		},
		{
			Id:     "710223",
			CityId: "7102",
			Name:   "Mandolang",
		},
		{
			Id:     "710224",
			CityId: "7102",
			Name:   "Tombariri Timur",
		},
		{
			Id:     "710225",
			CityId: "7102",
			Name:   "Tompaso Barat",
		},
		{
			Id:     "710308",
			CityId: "7103",
			Name:   "Tabukan Utara",
		},
		{
			Id:     "710309",
			CityId: "7103",
			Name:   "Nusa Tabukan",
		},
		{
			Id:     "710310",
			CityId: "7103",
			Name:   "Manganitu Selatan",
		},
		{
			Id:     "710311",
			CityId: "7103",
			Name:   "Tatoareng",
		},
		{
			Id:     "710312",
			CityId: "7103",
			Name:   "Tamako",
		},
		{
			Id:     "710313",
			CityId: "7103",
			Name:   "Manganitu",
		},
		{
			Id:     "710314",
			CityId: "7103",
			Name:   "Tabukan Tengah",
		},
		{
			Id:     "710315",
			CityId: "7103",
			Name:   "Tabukan Selatan",
		},
		{
			Id:     "710316",
			CityId: "7103",
			Name:   "Kendahe",
		},
		{
			Id:     "710317",
			CityId: "7103",
			Name:   "Tahuna",
		},
		{
			Id:     "710319",
			CityId: "7103",
			Name:   "Tabukan Selatan Tengah",
		},
		{
			Id:     "710320",
			CityId: "7103",
			Name:   "Tabukan Selatan Tenggara",
		},
		{
			Id:     "710323",
			CityId: "7103",
			Name:   "Tahuna Barat",
		},
		{
			Id:     "710324",
			CityId: "7103",
			Name:   "Tahuna Timur",
		},
		{
			Id:     "710325",
			CityId: "7103",
			Name:   "Kepulauan Marore",
		},
		{
			Id:     "710401",
			CityId: "7104",
			Name:   "Lirung",
		},
		{
			Id:     "710402",
			CityId: "7104",
			Name:   "Beo",
		},
		{
			Id:     "710403",
			CityId: "7104",
			Name:   "Rainis",
		},
		{
			Id:     "710404",
			CityId: "7104",
			Name:   "Essang",
		},
		{
			Id:     "710405",
			CityId: "7104",
			Name:   "Nanusa",
		},
		{
			Id:     "710406",
			CityId: "7104",
			Name:   "Kabaruan",
		},
		{
			Id:     "710407",
			CityId: "7104",
			Name:   "Melonguane",
		},
		{
			Id:     "710408",
			CityId: "7104",
			Name:   "Gemeh",
		},
		{
			Id:     "710409",
			CityId: "7104",
			Name:   "Damau",
		},
		{
			Id:     "710410",
			CityId: "7104",
			Name:   "Tampan' Amma",
		},
		{
			Id:     "710411",
			CityId: "7104",
			Name:   "Salibabu",
		},
		{
			Id:     "710412",
			CityId: "7104",
			Name:   "Kalongan",
		},
		{
			Id:     "710413",
			CityId: "7104",
			Name:   "Miangas",
		},
		{
			Id:     "710414",
			CityId: "7104",
			Name:   "Beo Utara",
		},
		{
			Id:     "710415",
			CityId: "7104",
			Name:   "Pulutan",
		},
		{
			Id:     "710416",
			CityId: "7104",
			Name:   "Melonguane Timur",
		},
		{
			Id:     "710417",
			CityId: "7104",
			Name:   "Moronge",
		},
		{
			Id:     "710418",
			CityId: "7104",
			Name:   "Beo Selatan",
		},
		{
			Id:     "710419",
			CityId: "7104",
			Name:   "Essang Selatan",
		},
		{
			Id:     "710501",
			CityId: "7105",
			Name:   "Modoinding",
		},
		{
			Id:     "710502",
			CityId: "7105",
			Name:   "Tompaso Baru",
		},
		{
			Id:     "710503",
			CityId: "7105",
			Name:   "Ranoyapo",
		},
		{
			Id:     "710507",
			CityId: "7105",
			Name:   "Motoling",
		},
		{
			Id:     "710508",
			CityId: "7105",
			Name:   "Sinonsayang",
		},
		{
			Id:     "710509",
			CityId: "7105",
			Name:   "Tenga",
		},
		{
			Id:     "710510",
			CityId: "7105",
			Name:   "Amurang",
		},
		{
			Id:     "710512",
			CityId: "7105",
			Name:   "Tumpaan",
		},
		{
			Id:     "710513",
			CityId: "7105",
			Name:   "Tareran",
		},
		{
			Id:     "710515",
			CityId: "7105",
			Name:   "Kumelembuai",
		},
		{
			Id:     "710516",
			CityId: "7105",
			Name:   "Maesaan",
		},
		{
			Id:     "710517",
			CityId: "7105",
			Name:   "Amurang Barat",
		},
		{
			Id:     "710518",
			CityId: "7105",
			Name:   "Amurang Timur",
		},
		{
			Id:     "710519",
			CityId: "7105",
			Name:   "Tatapaan",
		},
		{
			Id:     "710521",
			CityId: "7105",
			Name:   "Motoling Barat",
		},
		{
			Id:     "710522",
			CityId: "7105",
			Name:   "Motoling Timur",
		},
		{
			Id:     "710523",
			CityId: "7105",
			Name:   "Suluun Tareran",
		},
		{
			Id:     "710601",
			CityId: "7106",
			Name:   "Kema",
		},
		{
			Id:     "710602",
			CityId: "7106",
			Name:   "Kauditan",
		},
		{
			Id:     "710603",
			CityId: "7106",
			Name:   "Airmadidi",
		},
		{
			Id:     "710604",
			CityId: "7106",
			Name:   "Wori",
		},
		{
			Id:     "710605",
			CityId: "7106",
			Name:   "Dimembe",
		},
		{
			Id:     "710606",
			CityId: "7106",
			Name:   "Likupang Barat",
		},
		{
			Id:     "710607",
			CityId: "7106",
			Name:   "Likupang Timur",
		},
		{
			Id:     "710608",
			CityId: "7106",
			Name:   "Kalawat",
		},
		{
			Id:     "710609",
			CityId: "7106",
			Name:   "Talawaan",
		},
		{
			Id:     "710610",
			CityId: "7106",
			Name:   "Likupang Selatan",
		},
		{
			Id:     "710701",
			CityId: "7107",
			Name:   "Ratahan",
		},
		{
			Id:     "710702",
			CityId: "7107",
			Name:   "Pusomaen",
		},
		{
			Id:     "710703",
			CityId: "7107",
			Name:   "Belang",
		},
		{
			Id:     "710704",
			CityId: "7107",
			Name:   "Ratatotok",
		},
		{
			Id:     "710705",
			CityId: "7107",
			Name:   "Tombatu",
		},
		{
			Id:     "710706",
			CityId: "7107",
			Name:   "Touluaan",
		},
		{
			Id:     "710707",
			CityId: "7107",
			Name:   "Touluaan Selatan",
		},
		{
			Id:     "710708",
			CityId: "7107",
			Name:   "Silian Raya",
		},
		{
			Id:     "710709",
			CityId: "7107",
			Name:   "Tombatu Timur",
		},
		{
			Id:     "710710",
			CityId: "7107",
			Name:   "Tombatu Utara",
		},
		{
			Id:     "710711",
			CityId: "7107",
			Name:   "Pasan",
		},
		{
			Id:     "710712",
			CityId: "7107",
			Name:   "Ratahan Timur",
		},
		{
			Id:     "710801",
			CityId: "7108",
			Name:   "Sangkub",
		},
		{
			Id:     "710802",
			CityId: "7108",
			Name:   "Bintauna",
		},
		{
			Id:     "710803",
			CityId: "7108",
			Name:   "Bolangitang Timur",
		},
		{
			Id:     "710804",
			CityId: "7108",
			Name:   "Bolangitang Barat",
		},
		{
			Id:     "710805",
			CityId: "7108",
			Name:   "Kaidipang",
		},
		{
			Id:     "710806",
			CityId: "7108",
			Name:   "Pinogaluman",
		},
		{
			Id:     "710901",
			CityId: "7109",
			Name:   "Siau Timur",
		},
		{
			Id:     "710902",
			CityId: "7109",
			Name:   "Siau Barat",
		},
		{
			Id:     "710903",
			CityId: "7109",
			Name:   "Tagulandang",
		},
		{
			Id:     "710904",
			CityId: "7109",
			Name:   "Siau Timur Selatan",
		},
		{
			Id:     "710905",
			CityId: "7109",
			Name:   "Siau Barat Selatan",
		},
		{
			Id:     "710906",
			CityId: "7109",
			Name:   "Tagulandang Utara",
		},
		{
			Id:     "710907",
			CityId: "7109",
			Name:   "Biaro",
		},
		{
			Id:     "710908",
			CityId: "7109",
			Name:   "Siau Barat Utara",
		},
		{
			Id:     "710909",
			CityId: "7109",
			Name:   "Siau Tengah",
		},
		{
			Id:     "710910",
			CityId: "7109",
			Name:   "Tagulandang Selatan",
		},
		{
			Id:     "711001",
			CityId: "7110",
			Name:   "Tutuyan",
		},
		{
			Id:     "711002",
			CityId: "7110",
			Name:   "Kotabunan",
		},
		{
			Id:     "711003",
			CityId: "7110",
			Name:   "Nuangan",
		},
		{
			Id:     "711004",
			CityId: "7110",
			Name:   "Modayag",
		},
		{
			Id:     "711005",
			CityId: "7110",
			Name:   "Modayag Barat",
		},
		{
			Id:     "711006",
			CityId: "7110",
			Name:   "Motongkad",
		},
		{
			Id:     "711007",
			CityId: "7110",
			Name:   "Mooat",
		},
		{
			Id:     "711101",
			CityId: "7111",
			Name:   "Bolaang Uki",
		},
		{
			Id:     "711102",
			CityId: "7111",
			Name:   "Posigadan",
		},
		{
			Id:     "711103",
			CityId: "7111",
			Name:   "Pinolosian",
		},
		{
			Id:     "711104",
			CityId: "7111",
			Name:   "Pinolosian Tengah",
		},
		{
			Id:     "711105",
			CityId: "7111",
			Name:   "Pinolosian Timur",
		},
		{
			Id:     "711106",
			CityId: "7111",
			Name:   "Helumo",
		},
		{
			Id:     "711107",
			CityId: "7111",
			Name:   "Tomini",
		},
		{
			Id:     "717101",
			CityId: "7171",
			Name:   "Bunaken",
		},
		{
			Id:     "717102",
			CityId: "7171",
			Name:   "Tuminting",
		},
		{
			Id:     "717103",
			CityId: "7171",
			Name:   "Singkil",
		},
		{
			Id:     "717104",
			CityId: "7171",
			Name:   "Wenang",
		},
		{
			Id:     "717105",
			CityId: "7171",
			Name:   "Tikala",
		},
		{
			Id:     "717106",
			CityId: "7171",
			Name:   "Sario",
		},
		{
			Id:     "717107",
			CityId: "7171",
			Name:   "Wanea",
		},
		{
			Id:     "717108",
			CityId: "7171",
			Name:   "Mapanget",
		},
		{
			Id:     "717109",
			CityId: "7171",
			Name:   "Malalayang",
		},
		{
			Id:     "717110",
			CityId: "7171",
			Name:   "Bunaken Kepulauan",
		},
		{
			Id:     "717111",
			CityId: "7171",
			Name:   "Paal Dua",
		},
		{
			Id:     "717201",
			CityId: "7172",
			Name:   "Lembeh Selatan",
		},
		{
			Id:     "717202",
			CityId: "7172",
			Name:   "Madidir",
		},
		{
			Id:     "717203",
			CityId: "7172",
			Name:   "Ranowulu",
		},
		{
			Id:     "717204",
			CityId: "7172",
			Name:   "Aertembaga",
		},
		{
			Id:     "717205",
			CityId: "7172",
			Name:   "Matuari",
		},
		{
			Id:     "717206",
			CityId: "7172",
			Name:   "Girian",
		},
		{
			Id:     "717207",
			CityId: "7172",
			Name:   "Maesa",
		},
		{
			Id:     "717208",
			CityId: "7172",
			Name:   "Lembeh Utara",
		},
		{
			Id:     "717301",
			CityId: "7173",
			Name:   "Tomohon Selatan",
		},
		{
			Id:     "717302",
			CityId: "7173",
			Name:   "Tomohon Tengah",
		},
		{
			Id:     "717303",
			CityId: "7173",
			Name:   "Tomohon Utara",
		},
		{
			Id:     "717304",
			CityId: "7173",
			Name:   "Tomohon Barat",
		},
		{
			Id:     "717305",
			CityId: "7173",
			Name:   "Tomohon Timur",
		},
		{
			Id:     "717401",
			CityId: "7174",
			Name:   "Kotamobagu Utara",
		},
		{
			Id:     "717402",
			CityId: "7174",
			Name:   "Kotamobagu Timur",
		},
		{
			Id:     "717403",
			CityId: "7174",
			Name:   "Kotamobagu Selatan",
		},
		{
			Id:     "717404",
			CityId: "7174",
			Name:   "Kotamobagu Barat",
		},
		{
			Id:     "720101",
			CityId: "7201",
			Name:   "Batui",
		},
		{
			Id:     "720102",
			CityId: "7201",
			Name:   "Bunta",
		},
		{
			Id:     "720103",
			CityId: "7201",
			Name:   "Kintom",
		},
		{
			Id:     "720104",
			CityId: "7201",
			Name:   "Luwuk",
		},
		{
			Id:     "720105",
			CityId: "7201",
			Name:   "Lamala",
		},
		{
			Id:     "720106",
			CityId: "7201",
			Name:   "Balantak",
		},
		{
			Id:     "720107",
			CityId: "7201",
			Name:   "Pagimana",
		},
		{
			Id:     "720108",
			CityId: "7201",
			Name:   "Bualemo",
		},
		{
			Id:     "720109",
			CityId: "7201",
			Name:   "Toili",
		},
		{
			Id:     "720110",
			CityId: "7201",
			Name:   "Masama",
		},
		{
			Id:     "720111",
			CityId: "7201",
			Name:   "Luwuk Timur",
		},
		{
			Id:     "720112",
			CityId: "7201",
			Name:   "Toili Barat",
		},
		{
			Id:     "720113",
			CityId: "7201",
			Name:   "Nuhon",
		},
		{
			Id:     "720114",
			CityId: "7201",
			Name:   "Moilong",
		},
		{
			Id:     "720115",
			CityId: "7201",
			Name:   "Batui Selatan",
		},
		{
			Id:     "720116",
			CityId: "7201",
			Name:   "Lobu",
		},
		{
			Id:     "720117",
			CityId: "7201",
			Name:   "Simpang Raya",
		},
		{
			Id:     "720118",
			CityId: "7201",
			Name:   "Balantak Selatan",
		},
		{
			Id:     "720119",
			CityId: "7201",
			Name:   "Balantak Utara",
		},
		{
			Id:     "720120",
			CityId: "7201",
			Name:   "Luwuk Selatan",
		},
		{
			Id:     "720121",
			CityId: "7201",
			Name:   "Luwuk Utara",
		},
		{
			Id:     "720122",
			CityId: "7201",
			Name:   "Mantoh",
		},
		{
			Id:     "720123",
			CityId: "7201",
			Name:   "Nambo",
		},
		{
			Id:     "720201",
			CityId: "7202",
			Name:   "Poso Kota",
		},
		{
			Id:     "720202",
			CityId: "7202",
			Name:   "Poso Pesisir",
		},
		{
			Id:     "720203",
			CityId: "7202",
			Name:   "Lage",
		},
		{
			Id:     "720204",
			CityId: "7202",
			Name:   "Pamona Puselemba",
		},
		{
			Id:     "720205",
			CityId: "7202",
			Name:   "Pamona Timur",
		},
		{
			Id:     "720206",
			CityId: "7202",
			Name:   "Pamona Selatan",
		},
		{
			Id:     "720207",
			CityId: "7202",
			Name:   "Lore Utara",
		},
		{
			Id:     "720208",
			CityId: "7202",
			Name:   "Lore Tengah",
		},
		{
			Id:     "720209",
			CityId: "7202",
			Name:   "Lore Selatan",
		},
		{
			Id:     "720218",
			CityId: "7202",
			Name:   "Poso Pesisir Utara",
		},
		{
			Id:     "720219",
			CityId: "7202",
			Name:   "Poso Pesisir Selatan",
		},
		{
			Id:     "720220",
			CityId: "7202",
			Name:   "Pamona Barat",
		},
		{
			Id:     "720221",
			CityId: "7202",
			Name:   "Poso Kota Selatan",
		},
		{
			Id:     "720222",
			CityId: "7202",
			Name:   "Poso Kota Utara",
		},
		{
			Id:     "720223",
			CityId: "7202",
			Name:   "Lore Barat",
		},
		{
			Id:     "720224",
			CityId: "7202",
			Name:   "Lore Timur",
		},
		{
			Id:     "720225",
			CityId: "7202",
			Name:   "Lore Piore",
		},
		{
			Id:     "720226",
			CityId: "7202",
			Name:   "Pamona Tenggara",
		},
		{
			Id:     "720227",
			CityId: "7202",
			Name:   "Pamona Utara",
		},
		{
			Id:     "720304",
			CityId: "7203",
			Name:   "Rio Pakava",
		},
		{
			Id:     "720306",
			CityId: "7203",
			Name:   "Dampelas",
		},
		{
			Id:     "720308",
			CityId: "7203",
			Name:   "Banawa",
		},
		{
			Id:     "720309",
			CityId: "7203",
			Name:   "Labuan",
		},
		{
			Id:     "720310",
			CityId: "7203",
			Name:   "Sindue",
		},
		{
			Id:     "720311",
			CityId: "7203",
			Name:   "Sirenja",
		},
		{
			Id:     "720312",
			CityId: "7203",
			Name:   "Balaesang",
		},
		{
			Id:     "720314",
			CityId: "7203",
			Name:   "Sojol",
		},
		{
			Id:     "720318",
			CityId: "7203",
			Name:   "Banawa Selatan",
		},
		{
			Id:     "720319",
			CityId: "7203",
			Name:   "Tanantovea",
		},
		{
			Id:     "720321",
			CityId: "7203",
			Name:   "Pinembani",
		},
		{
			Id:     "720324",
			CityId: "7203",
			Name:   "Sindue Tombusabora",
		},
		{
			Id:     "720325",
			CityId: "7203",
			Name:   "Sindue Tobata",
		},
		{
			Id:     "720327",
			CityId: "7203",
			Name:   "Banawa Tengah",
		},
		{
			Id:     "720330",
			CityId: "7203",
			Name:   "Sojol Utara",
		},
		{
			Id:     "720331",
			CityId: "7203",
			Name:   "Balaesang Tanjung",
		},
		{
			Id:     "720401",
			CityId: "7204",
			Name:   "Dampal Selatan",
		},
		{
			Id:     "720402",
			CityId: "7204",
			Name:   "Dampal Utara",
		},
		{
			Id:     "720403",
			CityId: "7204",
			Name:   "Dondo",
		},
		{
			Id:     "720404",
			CityId: "7204",
			Name:   "Basidondo",
		},
		{
			Id:     "720405",
			CityId: "7204",
			Name:   "Ogodeide",
		},
		{
			Id:     "720406",
			CityId: "7204",
			Name:   "Lampasio",
		},
		{
			Id:     "720407",
			CityId: "7204",
			Name:   "Baolan",
		},
		{
			Id:     "720408",
			CityId: "7204",
			Name:   "Galang",
		},
		{
			Id:     "720409",
			CityId: "7204",
			Name:   "Toli-Toli Utara",
		},
		{
			Id:     "720410",
			CityId: "7204",
			Name:   "Dako Pemean",
		},
		{
			Id:     "720501",
			CityId: "7205",
			Name:   "Momunu",
		},
		{
			Id:     "720502",
			CityId: "7205",
			Name:   "Lakea",
		},
		{
			Id:     "720503",
			CityId: "7205",
			Name:   "Bokat",
		},
		{
			Id:     "720504",
			CityId: "7205",
			Name:   "Bunobogu",
		},
		{
			Id:     "720505",
			CityId: "7205",
			Name:   "Paleleh",
		},
		{
			Id:     "720506",
			CityId: "7205",
			Name:   "Biau",
		},
		{
			Id:     "720507",
			CityId: "7205",
			Name:   "Tiloan",
		},
		{
			Id:     "720508",
			CityId: "7205",
			Name:   "Bukal",
		},
		{
			Id:     "720509",
			CityId: "7205",
			Name:   "Gadung",
		},
		{
			Id:     "720510",
			CityId: "7205",
			Name:   "Karamat",
		},
		{
			Id:     "720511",
			CityId: "7205",
			Name:   "Paleleh Barat",
		},
		{
			Id:     "720605",
			CityId: "7206",
			Name:   "Bungku Tengah",
		},
		{
			Id:     "720606",
			CityId: "7206",
			Name:   "Bungku Selatan",
		},
		{
			Id:     "720607",
			CityId: "7206",
			Name:   "Menui Kepulauan",
		},
		{
			Id:     "720608",
			CityId: "7206",
			Name:   "Bungku Barat",
		},
		{
			Id:     "720609",
			CityId: "7206",
			Name:   "Bumi Raya",
		},
		{
			Id:     "720610",
			CityId: "7206",
			Name:   "Bahodopi",
		},
		{
			Id:     "720612",
			CityId: "7206",
			Name:   "Wita Ponda",
		},
		{
			Id:     "720615",
			CityId: "7206",
			Name:   "Bungku Pesisir",
		},
		{
			Id:     "720618",
			CityId: "7206",
			Name:   "Bungku Timur",
		},
		{
			Id:     "720703",
			CityId: "7207",
			Name:   "Totikum",
		},
		{
			Id:     "720704",
			CityId: "7207",
			Name:   "Tinangkung",
		},
		{
			Id:     "720705",
			CityId: "7207",
			Name:   "Liang",
		},
		{
			Id:     "720706",
			CityId: "7207",
			Name:   "Bulagi",
		},
		{
			Id:     "720707",
			CityId: "7207",
			Name:   "Buko",
		},
		{
			Id:     "720709",
			CityId: "7207",
			Name:   "Bulagi Selatan",
		},
		{
			Id:     "720711",
			CityId: "7207",
			Name:   "Tinangkung Selatan",
		},
		{
			Id:     "720715",
			CityId: "7207",
			Name:   "Totikum Selatan",
		},
		{
			Id:     "720716",
			CityId: "7207",
			Name:   "Peling Tengah",
		},
		{
			Id:     "720717",
			CityId: "7207",
			Name:   "Bulagi Utara",
		},
		{
			Id:     "720718",
			CityId: "7207",
			Name:   "Buko Selatan",
		},
		{
			Id:     "720719",
			CityId: "7207",
			Name:   "Tinangkung Utara",
		},
		{
			Id:     "720801",
			CityId: "7208",
			Name:   "Parigi",
		},
		{
			Id:     "720802",
			CityId: "7208",
			Name:   "Ampibabo",
		},
		{
			Id:     "720803",
			CityId: "7208",
			Name:   "Tinombo",
		},
		{
			Id:     "720804",
			CityId: "7208",
			Name:   "Moutong",
		},
		{
			Id:     "720805",
			CityId: "7208",
			Name:   "Tomini",
		},
		{
			Id:     "720806",
			CityId: "7208",
			Name:   "Sausu",
		},
		{
			Id:     "720807",
			CityId: "7208",
			Name:   "Bolano Lambunu",
		},
		{
			Id:     "720808",
			CityId: "7208",
			Name:   "Kasimbar",
		},
		{
			Id:     "720809",
			CityId: "7208",
			Name:   "Torue",
		},
		{
			Id:     "720810",
			CityId: "7208",
			Name:   "Tinombo Selatan",
		},
		{
			Id:     "720811",
			CityId: "7208",
			Name:   "Parigi Selatan",
		},
		{
			Id:     "720812",
			CityId: "7208",
			Name:   "Mepanga",
		},
		{
			Id:     "720813",
			CityId: "7208",
			Name:   "Toribulu",
		},
		{
			Id:     "720814",
			CityId: "7208",
			Name:   "Taopa",
		},
		{
			Id:     "720815",
			CityId: "7208",
			Name:   "Balinggi",
		},
		{
			Id:     "720816",
			CityId: "7208",
			Name:   "Parigi Barat",
		},
		{
			Id:     "720817",
			CityId: "7208",
			Name:   "Siniu",
		},
		{
			Id:     "720818",
			CityId: "7208",
			Name:   "Palasa",
		},
		{
			Id:     "720819",
			CityId: "7208",
			Name:   "Parigi Utara",
		},
		{
			Id:     "720820",
			CityId: "7208",
			Name:   "Parigi Tengah",
		},
		{
			Id:     "720821",
			CityId: "7208",
			Name:   "Bolano",
		},
		{
			Id:     "720822",
			CityId: "7208",
			Name:   "Ongka Malino",
		},
		{
			Id:     "720823",
			CityId: "7208",
			Name:   "Sidoan",
		},
		{
			Id:     "720901",
			CityId: "7209",
			Name:   "Una Una",
		},
		{
			Id:     "720902",
			CityId: "7209",
			Name:   "Togean",
		},
		{
			Id:     "720903",
			CityId: "7209",
			Name:   "Walea Kepulauan",
		},
		{
			Id:     "720904",
			CityId: "7209",
			Name:   "Ampana Tete",
		},
		{
			Id:     "720905",
			CityId: "7209",
			Name:   "Ampana Kota",
		},
		{
			Id:     "720906",
			CityId: "7209",
			Name:   "Ulubongka",
		},
		{
			Id:     "720907",
			CityId: "7209",
			Name:   "Tojo Barat",
		},
		{
			Id:     "720908",
			CityId: "7209",
			Name:   "Tojo",
		},
		{
			Id:     "720909",
			CityId: "7209",
			Name:   "Walea Besar",
		},
		{
			Id:     "720910",
			CityId: "7209",
			Name:   "Ratolindo",
		},
		{
			Id:     "720911",
			CityId: "7209",
			Name:   "Batudaka",
		},
		{
			Id:     "720912",
			CityId: "7209",
			Name:   "Talatako",
		},
		{
			Id:     "721001",
			CityId: "7210",
			Name:   "Sigi Biromaru",
		},
		{
			Id:     "721002",
			CityId: "7210",
			Name:   "Palolo",
		},
		{
			Id:     "721003",
			CityId: "7210",
			Name:   "Nokilalaki",
		},
		{
			Id:     "721004",
			CityId: "7210",
			Name:   "Lindu",
		},
		{
			Id:     "721005",
			CityId: "7210",
			Name:   "Kulawi",
		},
		{
			Id:     "721006",
			CityId: "7210",
			Name:   "Kulawi Selatan",
		},
		{
			Id:     "721007",
			CityId: "7210",
			Name:   "Pipikoro",
		},
		{
			Id:     "721008",
			CityId: "7210",
			Name:   "Gumbasa",
		},
		{
			Id:     "721009",
			CityId: "7210",
			Name:   "Dolo Selatan",
		},
		{
			Id:     "721010",
			CityId: "7210",
			Name:   "Tanambulava",
		},
		{
			Id:     "721011",
			CityId: "7210",
			Name:   "Dolo Barat",
		},
		{
			Id:     "721012",
			CityId: "7210",
			Name:   "Dolo",
		},
		{
			Id:     "721013",
			CityId: "7210",
			Name:   "Kinovaro",
		},
		{
			Id:     "721014",
			CityId: "7210",
			Name:   "Marawola",
		},
		{
			Id:     "721015",
			CityId: "7210",
			Name:   "Marawola Barat",
		},
		{
			Id:     "721101",
			CityId: "7211",
			Name:   "Banggai",
		},
		{
			Id:     "721102",
			CityId: "7211",
			Name:   "Banggai Utara",
		},
		{
			Id:     "721103",
			CityId: "7211",
			Name:   "Bokan Kepulauan",
		},
		{
			Id:     "721104",
			CityId: "7211",
			Name:   "Bangkurung",
		},
		{
			Id:     "721105",
			CityId: "7211",
			Name:   "Labobo",
		},
		{
			Id:     "721106",
			CityId: "7211",
			Name:   "Banggai Selatan",
		},
		{
			Id:     "721107",
			CityId: "7211",
			Name:   "Banggai Tengah",
		},
		{
			Id:     "721201",
			CityId: "7212",
			Name:   "Petasia",
		},
		{
			Id:     "721202",
			CityId: "7212",
			Name:   "Petasia Timur",
		},
		{
			Id:     "721203",
			CityId: "7212",
			Name:   "Lembo Raya",
		},
		{
			Id:     "721204",
			CityId: "7212",
			Name:   "Lembo",
		},
		{
			Id:     "721205",
			CityId: "7212",
			Name:   "Mori Atas",
		},
		{
			Id:     "721206",
			CityId: "7212",
			Name:   "Mori Utara",
		},
		{
			Id:     "721207",
			CityId: "7212",
			Name:   "Soyo Jaya",
		},
		{
			Id:     "721208",
			CityId: "7212",
			Name:   "Bungku Utara",
		},
		{
			Id:     "721209",
			CityId: "7212",
			Name:   "Mamosalato",
		},
		{
			Id:     "721210",
			CityId: "7212",
			Name:   "Petasia Barat",
		},
		{
			Id:     "727101",
			CityId: "7271",
			Name:   "Palu Timur",
		},
		{
			Id:     "727102",
			CityId: "7271",
			Name:   "Palu Barat",
		},
		{
			Id:     "727103",
			CityId: "7271",
			Name:   "Palu Selatan",
		},
		{
			Id:     "727104",
			CityId: "7271",
			Name:   "Palu Utara",
		},
		{
			Id:     "727105",
			CityId: "7271",
			Name:   "Ulujadi",
		},
		{
			Id:     "727106",
			CityId: "7271",
			Name:   "Tatanga",
		},
		{
			Id:     "727107",
			CityId: "7271",
			Name:   "Tawaeli",
		},
		{
			Id:     "727108",
			CityId: "7271",
			Name:   "Mantikulore",
		},
		{
			Id:     "730101",
			CityId: "7301",
			Name:   "Benteng",
		},
		{
			Id:     "730102",
			CityId: "7301",
			Name:   "Bontoharu",
		},
		{
			Id:     "730103",
			CityId: "7301",
			Name:   "Bontomatene",
		},
		{
			Id:     "730104",
			CityId: "7301",
			Name:   "Bontomanai",
		},
		{
			Id:     "730105",
			CityId: "7301",
			Name:   "Bontosikuyu",
		},
		{
			Id:     "730106",
			CityId: "7301",
			Name:   "Pasimasunggu",
		},
		{
			Id:     "730107",
			CityId: "7301",
			Name:   "Pasimarannu",
		},
		{
			Id:     "730108",
			CityId: "7301",
			Name:   "Taka Bonerate",
		},
		{
			Id:     "730109",
			CityId: "7301",
			Name:   "Pasilambena",
		},
		{
			Id:     "730110",
			CityId: "7301",
			Name:   "Pasimasunggu Timur",
		},
		{
			Id:     "730111",
			CityId: "7301",
			Name:   "Buki",
		},
		{
			Id:     "730201",
			CityId: "7302",
			Name:   "Gantarang",
		},
		{
			Id:     "730202",
			CityId: "7302",
			Name:   "Ujung Bulu",
		},
		{
			Id:     "730203",
			CityId: "7302",
			Name:   "Bonto Bahari",
		},
		{
			Id:     "730204",
			CityId: "7302",
			Name:   "Bonto Tiro",
		},
		{
			Id:     "730205",
			CityId: "7302",
			Name:   "Herlang",
		},
		{
			Id:     "730206",
			CityId: "7302",
			Name:   "Kajang",
		},
		{
			Id:     "730207",
			CityId: "7302",
			Name:   "Bulukumpa",
		},
		{
			Id:     "730208",
			CityId: "7302",
			Name:   "Kindang",
		},
		{
			Id:     "730209",
			CityId: "7302",
			Name:   "Ujungloe",
		},
		{
			Id:     "730210",
			CityId: "7302",
			Name:   "Rilauale",
		},
		{
			Id:     "730301",
			CityId: "7303",
			Name:   "Bissappu",
		},
		{
			Id:     "730302",
			CityId: "7303",
			Name:   "Bantaeng",
		},
		{
			Id:     "730303",
			CityId: "7303",
			Name:   "Eremerasa",
		},
		{
			Id:     "730304",
			CityId: "7303",
			Name:   "Tompo Bulu",
		},
		{
			Id:     "730305",
			CityId: "7303",
			Name:   "Pajukukang",
		},
		{
			Id:     "730306",
			CityId: "7303",
			Name:   "Uluere",
		},
		{
			Id:     "730307",
			CityId: "7303",
			Name:   "Gantarang Keke",
		},
		{
			Id:     "730308",
			CityId: "7303",
			Name:   "Sinoa",
		},
		{
			Id:     "730401",
			CityId: "7304",
			Name:   "Bangkala",
		},
		{
			Id:     "730402",
			CityId: "7304",
			Name:   "Tamalatea",
		},
		{
			Id:     "730403",
			CityId: "7304",
			Name:   "Binamu",
		},
		{
			Id:     "730404",
			CityId: "7304",
			Name:   "Batang",
		},
		{
			Id:     "730405",
			CityId: "7304",
			Name:   "Kelara",
		},
		{
			Id:     "730406",
			CityId: "7304",
			Name:   "Bangkala Barat",
		},
		{
			Id:     "730407",
			CityId: "7304",
			Name:   "Bontoramba",
		},
		{
			Id:     "730408",
			CityId: "7304",
			Name:   "Turatea",
		},
		{
			Id:     "730409",
			CityId: "7304",
			Name:   "Arungkeke",
		},
		{
			Id:     "730410",
			CityId: "7304",
			Name:   "Rumbia",
		},
		{
			Id:     "730411",
			CityId: "7304",
			Name:   "Tarowang",
		},
		{
			Id:     "730501",
			CityId: "7305",
			Name:   "Mappakasunggu",
		},
		{
			Id:     "730502",
			CityId: "7305",
			Name:   "Mangarabombang",
		},
		{
			Id:     "730503",
			CityId: "7305",
			Name:   "Polongbangkeng Selatan",
		},
		{
			Id:     "730504",
			CityId: "7305",
			Name:   "Polongbangkeng Utara",
		},
		{
			Id:     "730505",
			CityId: "7305",
			Name:   "Galesong Selatan",
		},
		{
			Id:     "730506",
			CityId: "7305",
			Name:   "Galesong Utara",
		},
		{
			Id:     "730507",
			CityId: "7305",
			Name:   "Pattallassang",
		},
		{
			Id:     "730508",
			CityId: "7305",
			Name:   "Sanrobone",
		},
		{
			Id:     "730509",
			CityId: "7305",
			Name:   "Galesong",
		},
		{
			Id:     "730510",
			CityId: "7305",
			Name:   "Kepulauan Tanakeke",
		},
		{
			Id:     "730511",
			CityId: "7305",
			Name:   "Polongbangkeng Timur",
		},
		{
			Id:     "730512",
			CityId: "7305",
			Name:   "Laikang",
		},
		{
			Id:     "730601",
			CityId: "7306",
			Name:   "Bontonompo",
		},
		{
			Id:     "730602",
			CityId: "7306",
			Name:   "Bajeng",
		},
		{
			Id:     "730603",
			CityId: "7306",
			Name:   "Tompobulu",
		},
		{
			Id:     "730604",
			CityId: "7306",
			Name:   "Tinggimoncong",
		},
		{
			Id:     "730605",
			CityId: "7306",
			Name:   "Parangloe",
		},
		{
			Id:     "730606",
			CityId: "7306",
			Name:   "Bontomarannu",
		},
		{
			Id:     "730607",
			CityId: "7306",
			Name:   "Pallangga",
		},
		{
			Id:     "730608",
			CityId: "7306",
			Name:   "Somba Opu",
		},
		{
			Id:     "730609",
			CityId: "7306",
			Name:   "Bungaya",
		},
		{
			Id:     "730610",
			CityId: "7306",
			Name:   "Tombolopao",
		},
		{
			Id:     "730611",
			CityId: "7306",
			Name:   "Biringbulu",
		},
		{
			Id:     "730612",
			CityId: "7306",
			Name:   "Barombong",
		},
		{
			Id:     "730613",
			CityId: "7306",
			Name:   "Pattallasang",
		},
		{
			Id:     "730614",
			CityId: "7306",
			Name:   "Manuju",
		},
		{
			Id:     "730615",
			CityId: "7306",
			Name:   "Bontolempangang",
		},
		{
			Id:     "730616",
			CityId: "7306",
			Name:   "Bontonompo Selatan",
		},
		{
			Id:     "730617",
			CityId: "7306",
			Name:   "Parigi",
		},
		{
			Id:     "730618",
			CityId: "7306",
			Name:   "Bajeng Barat",
		},
		{
			Id:     "730701",
			CityId: "7307",
			Name:   "Sinjai Barat",
		},
		{
			Id:     "730702",
			CityId: "7307",
			Name:   "Sinjai Selatan",
		},
		{
			Id:     "730703",
			CityId: "7307",
			Name:   "Sinjai Timur",
		},
		{
			Id:     "730704",
			CityId: "7307",
			Name:   "Sinjai Tengah",
		},
		{
			Id:     "730705",
			CityId: "7307",
			Name:   "Sinjai Utara",
		},
		{
			Id:     "730706",
			CityId: "7307",
			Name:   "Bulupoddo",
		},
		{
			Id:     "730707",
			CityId: "7307",
			Name:   "Sinjai Borong",
		},
		{
			Id:     "730708",
			CityId: "7307",
			Name:   "Tellu Limpoe",
		},
		{
			Id:     "730709",
			CityId: "7307",
			Name:   "Pulau Sembilan",
		},
		{
			Id:     "730801",
			CityId: "7308",
			Name:   "Bontocani",
		},
		{
			Id:     "730802",
			CityId: "7308",
			Name:   "Kahu",
		},
		{
			Id:     "730803",
			CityId: "7308",
			Name:   "Kajuara",
		},
		{
			Id:     "730804",
			CityId: "7308",
			Name:   "Salomekko",
		},
		{
			Id:     "730805",
			CityId: "7308",
			Name:   "Tonra",
		},
		{
			Id:     "730806",
			CityId: "7308",
			Name:   "Libureng",
		},
		{
			Id:     "730807",
			CityId: "7308",
			Name:   "Mare",
		},
		{
			Id:     "730808",
			CityId: "7308",
			Name:   "Sibulue",
		},
		{
			Id:     "730809",
			CityId: "7308",
			Name:   "Barebbo",
		},
		{
			Id:     "730810",
			CityId: "7308",
			Name:   "Cina",
		},
		{
			Id:     "730811",
			CityId: "7308",
			Name:   "Ponre",
		},
		{
			Id:     "730812",
			CityId: "7308",
			Name:   "Lappariaja",
		},
		{
			Id:     "730813",
			CityId: "7308",
			Name:   "Lamuru",
		},
		{
			Id:     "730814",
			CityId: "7308",
			Name:   "Ulaweng",
		},
		{
			Id:     "730815",
			CityId: "7308",
			Name:   "Palakka",
		},
		{
			Id:     "730816",
			CityId: "7308",
			Name:   "Awangpone",
		},
		{
			Id:     "730817",
			CityId: "7308",
			Name:   "Tellu Siattinge",
		},
		{
			Id:     "730818",
			CityId: "7308",
			Name:   "Ajangale",
		},
		{
			Id:     "730819",
			CityId: "7308",
			Name:   "Dua Boccoe",
		},
		{
			Id:     "730820",
			CityId: "7308",
			Name:   "Cenrana",
		},
		{
			Id:     "730821",
			CityId: "7308",
			Name:   "Tanete Riattang",
		},
		{
			Id:     "730822",
			CityId: "7308",
			Name:   "Tanete Riattang Barat",
		},
		{
			Id:     "730823",
			CityId: "7308",
			Name:   "Tanete Riattang Timur",
		},
		{
			Id:     "730824",
			CityId: "7308",
			Name:   "Amali",
		},
		{
			Id:     "730825",
			CityId: "7308",
			Name:   "Tellulimpoe",
		},
		{
			Id:     "730826",
			CityId: "7308",
			Name:   "Bengo",
		},
		{
			Id:     "730827",
			CityId: "7308",
			Name:   "Patimpeng",
		},
		{
			Id:     "730901",
			CityId: "7309",
			Name:   "Mandai",
		},
		{
			Id:     "730902",
			CityId: "7309",
			Name:   "Camba",
		},
		{
			Id:     "730903",
			CityId: "7309",
			Name:   "Bantimurung",
		},
		{
			Id:     "730904",
			CityId: "7309",
			Name:   "Maros Baru",
		},
		{
			Id:     "730905",
			CityId: "7309",
			Name:   "Bontoa",
		},
		{
			Id:     "730906",
			CityId: "7309",
			Name:   "Malllawa",
		},
		{
			Id:     "730907",
			CityId: "7309",
			Name:   "Tanralili",
		},
		{
			Id:     "730908",
			CityId: "7309",
			Name:   "Marusu",
		},
		{
			Id:     "730909",
			CityId: "7309",
			Name:   "Simbang",
		},
		{
			Id:     "730910",
			CityId: "7309",
			Name:   "Cenrana",
		},
		{
			Id:     "730911",
			CityId: "7309",
			Name:   "Tompobulu",
		},
		{
			Id:     "730912",
			CityId: "7309",
			Name:   "Lau",
		},
		{
			Id:     "730913",
			CityId: "7309",
			Name:   "Moncongloe",
		},
		{
			Id:     "730914",
			CityId: "7309",
			Name:   "Turikale",
		},
		{
			Id:     "731001",
			CityId: "7310",
			Name:   "Liukang Tangaya",
		},
		{
			Id:     "731002",
			CityId: "7310",
			Name:   "Liukang Kalmas",
		},
		{
			Id:     "731003",
			CityId: "7310",
			Name:   "Liukang Tupabbiring",
		},
		{
			Id:     "731004",
			CityId: "7310",
			Name:   "Pangkajene",
		},
		{
			Id:     "731005",
			CityId: "7310",
			Name:   "Balocci",
		},
		{
			Id:     "731006",
			CityId: "7310",
			Name:   "Bungoro",
		},
		{
			Id:     "731007",
			CityId: "7310",
			Name:   "Labakkang",
		},
		{
			Id:     "731008",
			CityId: "7310",
			Name:   "Marang",
		},
		{
			Id:     "731009",
			CityId: "7310",
			Name:   "Segeri",
		},
		{
			Id:     "731010",
			CityId: "7310",
			Name:   "Minasa Tene",
		},
		{
			Id:     "731011",
			CityId: "7310",
			Name:   "Mandalle",
		},
		{
			Id:     "731012",
			CityId: "7310",
			Name:   "Tondong Tallasa",
		},
		{
			Id:     "731013",
			CityId: "7310",
			Name:   "Liukang Tupabbiring Utara",
		},
		{
			Id:     "731101",
			CityId: "7311",
			Name:   "Tanete Riaja",
		},
		{
			Id:     "731102",
			CityId: "7311",
			Name:   "Tanete Rilau",
		},
		{
			Id:     "731103",
			CityId: "7311",
			Name:   "Barru",
		},
		{
			Id:     "731104",
			CityId: "7311",
			Name:   "Soppeng Riaja",
		},
		{
			Id:     "731105",
			CityId: "7311",
			Name:   "Mallusetasi",
		},
		{
			Id:     "731106",
			CityId: "7311",
			Name:   "Pujananting",
		},
		{
			Id:     "731107",
			CityId: "7311",
			Name:   "Balusu",
		},
		{
			Id:     "731201",
			CityId: "7312",
			Name:   "Marioriwawo",
		},
		{
			Id:     "731202",
			CityId: "7312",
			Name:   "Liliraja",
		},
		{
			Id:     "731203",
			CityId: "7312",
			Name:   "Lilirilau",
		},
		{
			Id:     "731204",
			CityId: "7312",
			Name:   "Lalabata",
		},
		{
			Id:     "731205",
			CityId: "7312",
			Name:   "Marioriawa",
		},
		{
			Id:     "731206",
			CityId: "7312",
			Name:   "Donri Donri",
		},
		{
			Id:     "731207",
			CityId: "7312",
			Name:   "Ganra",
		},
		{
			Id:     "731208",
			CityId: "7312",
			Name:   "Citta",
		},
		{
			Id:     "731301",
			CityId: "7313",
			Name:   "Sabangparu",
		},
		{
			Id:     "731302",
			CityId: "7313",
			Name:   "Pammana",
		},
		{
			Id:     "731303",
			CityId: "7313",
			Name:   "Takkalalla",
		},
		{
			Id:     "731304",
			CityId: "7313",
			Name:   "Sajoanging",
		},
		{
			Id:     "731305",
			CityId: "7313",
			Name:   "Majauleng",
		},
		{
			Id:     "731306",
			CityId: "7313",
			Name:   "Tempe",
		},
		{
			Id:     "731307",
			CityId: "7313",
			Name:   "Belawa",
		},
		{
			Id:     "731308",
			CityId: "7313",
			Name:   "Tanasitolo",
		},
		{
			Id:     "731309",
			CityId: "7313",
			Name:   "Maniangpajo",
		},
		{
			Id:     "731310",
			CityId: "7313",
			Name:   "Pitumpanua",
		},
		{
			Id:     "731311",
			CityId: "7313",
			Name:   "Bola",
		},
		{
			Id:     "731312",
			CityId: "7313",
			Name:   "Penrang",
		},
		{
			Id:     "731313",
			CityId: "7313",
			Name:   "Gilireng",
		},
		{
			Id:     "731314",
			CityId: "7313",
			Name:   "Keera",
		},
		{
			Id:     "731401",
			CityId: "7314",
			Name:   "Panca Lautang",
		},
		{
			Id:     "731402",
			CityId: "7314",
			Name:   "Tellu Limpoe",
		},
		{
			Id:     "731403",
			CityId: "7314",
			Name:   "Watang Pulu",
		},
		{
			Id:     "731404",
			CityId: "7314",
			Name:   "Baranti",
		},
		{
			Id:     "731405",
			CityId: "7314",
			Name:   "Panca Rijang",
		},
		{
			Id:     "731406",
			CityId: "7314",
			Name:   "Kulo",
		},
		{
			Id:     "731407",
			CityId: "7314",
			Name:   "Maritengngae",
		},
		{
			Id:     "731408",
			CityId: "7314",
			Name:   "Watang Sidenreng",
		},
		{
			Id:     "731409",
			CityId: "7314",
			Name:   "Dua Pitue",
		},
		{
			Id:     "731410",
			CityId: "7314",
			Name:   "Pitu Riawa",
		},
		{
			Id:     "731411",
			CityId: "7314",
			Name:   "Pitu Riase",
		},
		{
			Id:     "731501",
			CityId: "7315",
			Name:   "Mattiro Sompe",
		},
		{
			Id:     "731502",
			CityId: "7315",
			Name:   "Suppa",
		},
		{
			Id:     "731503",
			CityId: "7315",
			Name:   "Mattiro Bulu",
		},
		{
			Id:     "731504",
			CityId: "7315",
			Name:   "Watang Sawitto",
		},
		{
			Id:     "731505",
			CityId: "7315",
			Name:   "Patampanua",
		},
		{
			Id:     "731506",
			CityId: "7315",
			Name:   "Duampanua",
		},
		{
			Id:     "731507",
			CityId: "7315",
			Name:   "Lembang",
		},
		{
			Id:     "731508",
			CityId: "7315",
			Name:   "Cempa",
		},
		{
			Id:     "731509",
			CityId: "7315",
			Name:   "Tiroang",
		},
		{
			Id:     "731510",
			CityId: "7315",
			Name:   "Lanrisang",
		},
		{
			Id:     "731511",
			CityId: "7315",
			Name:   "Paleteang",
		},
		{
			Id:     "731512",
			CityId: "7315",
			Name:   "Batulappa",
		},
		{
			Id:     "731601",
			CityId: "7316",
			Name:   "Maiwa",
		},
		{
			Id:     "731602",
			CityId: "7316",
			Name:   "Enrekang",
		},
		{
			Id:     "731603",
			CityId: "7316",
			Name:   "Baraka",
		},
		{
			Id:     "731604",
			CityId: "7316",
			Name:   "Anggeraja",
		},
		{
			Id:     "731605",
			CityId: "7316",
			Name:   "Alla",
		},
		{
			Id:     "731606",
			CityId: "7316",
			Name:   "Bungin",
		},
		{
			Id:     "731607",
			CityId: "7316",
			Name:   "Cendana",
		},
		{
			Id:     "731608",
			CityId: "7316",
			Name:   "Curio",
		},
		{
			Id:     "731609",
			CityId: "7316",
			Name:   "Malua",
		},
		{
			Id:     "731610",
			CityId: "7316",
			Name:   "Buntu Batu",
		},
		{
			Id:     "731611",
			CityId: "7316",
			Name:   "Masalle",
		},
		{
			Id:     "731612",
			CityId: "7316",
			Name:   "Baroko",
		},
		{
			Id:     "731701",
			CityId: "7317",
			Name:   "Basse Sangtempe",
		},
		{
			Id:     "731702",
			CityId: "7317",
			Name:   "Larompong",
		},
		{
			Id:     "731703",
			CityId: "7317",
			Name:   "Suli",
		},
		{
			Id:     "731704",
			CityId: "7317",
			Name:   "Bajo",
		},
		{
			Id:     "731705",
			CityId: "7317",
			Name:   "Bua Ponrang",
		},
		{
			Id:     "731706",
			CityId: "7317",
			Name:   "Walenrang",
		},
		{
			Id:     "731707",
			CityId: "7317",
			Name:   "Belopa",
		},
		{
			Id:     "731708",
			CityId: "7317",
			Name:   "Bua",
		},
		{
			Id:     "731709",
			CityId: "7317",
			Name:   "Lamasi",
		},
		{
			Id:     "731710",
			CityId: "7317",
			Name:   "Larompong Selatan",
		},
		{
			Id:     "731711",
			CityId: "7317",
			Name:   "Ponrang",
		},
		{
			Id:     "731712",
			CityId: "7317",
			Name:   "Latimojong",
		},
		{
			Id:     "731713",
			CityId: "7317",
			Name:   "Kamanre",
		},
		{
			Id:     "731714",
			CityId: "7317",
			Name:   "Belopa Utara",
		},
		{
			Id:     "731715",
			CityId: "7317",
			Name:   "Walenrang Barat",
		},
		{
			Id:     "731716",
			CityId: "7317",
			Name:   "Walenrang Utara",
		},
		{
			Id:     "731717",
			CityId: "7317",
			Name:   "Walenrang Timur",
		},
		{
			Id:     "731718",
			CityId: "7317",
			Name:   "Lamasi Timur",
		},
		{
			Id:     "731719",
			CityId: "7317",
			Name:   "Suli Barat",
		},
		{
			Id:     "731720",
			CityId: "7317",
			Name:   "Bajo Barat",
		},
		{
			Id:     "731721",
			CityId: "7317",
			Name:   "Ponrang Selatan",
		},
		{
			Id:     "731722",
			CityId: "7317",
			Name:   "Basse Sangtempe Utara",
		},
		{
			Id:     "731801",
			CityId: "7318",
			Name:   "Saluputi",
		},
		{
			Id:     "731802",
			CityId: "7318",
			Name:   "Bittuang",
		},
		{
			Id:     "731803",
			CityId: "7318",
			Name:   "Bonggakaradeng",
		},
		{
			Id:     "731805",
			CityId: "7318",
			Name:   "Makale",
		},
		{
			Id:     "731809",
			CityId: "7318",
			Name:   "Simbuang",
		},
		{
			Id:     "731811",
			CityId: "7318",
			Name:   "Rantetayo",
		},
		{
			Id:     "731812",
			CityId: "7318",
			Name:   "Mengkendek",
		},
		{
			Id:     "731813",
			CityId: "7318",
			Name:   "Sangalla",
		},
		{
			Id:     "731819",
			CityId: "7318",
			Name:   "Gandangbatu Sillanan",
		},
		{
			Id:     "731820",
			CityId: "7318",
			Name:   "Rembon",
		},
		{
			Id:     "731827",
			CityId: "7318",
			Name:   "Makale Utara",
		},
		{
			Id:     "731828",
			CityId: "7318",
			Name:   "Mappak",
		},
		{
			Id:     "731829",
			CityId: "7318",
			Name:   "Makale Selatan",
		},
		{
			Id:     "731831",
			CityId: "7318",
			Name:   "Masanda",
		},
		{
			Id:     "731833",
			CityId: "7318",
			Name:   "Sangalla Selatan",
		},
		{
			Id:     "731834",
			CityId: "7318",
			Name:   "Sangalla Utara",
		},
		{
			Id:     "731835",
			CityId: "7318",
			Name:   "Malimbong Balepe",
		},
		{
			Id:     "731837",
			CityId: "7318",
			Name:   "Rano",
		},
		{
			Id:     "731838",
			CityId: "7318",
			Name:   "Kurra",
		},
		{
			Id:     "732201",
			CityId: "7322",
			Name:   "Malangke",
		},
		{
			Id:     "732202",
			CityId: "7322",
			Name:   "Bone Bone",
		},
		{
			Id:     "732203",
			CityId: "7322",
			Name:   "Masamba",
		},
		{
			Id:     "732204",
			CityId: "7322",
			Name:   "Sabbang",
		},
		{
			Id:     "732205",
			CityId: "7322",
			Name:   "Rongkong",
		},
		{
			Id:     "732206",
			CityId: "7322",
			Name:   "Sukamaju",
		},
		{
			Id:     "732207",
			CityId: "7322",
			Name:   "Seko",
		},
		{
			Id:     "732208",
			CityId: "7322",
			Name:   "Malangke Barat",
		},
		{
			Id:     "732209",
			CityId: "7322",
			Name:   "Rampi",
		},
		{
			Id:     "732210",
			CityId: "7322",
			Name:   "Mappedeceng",
		},
		{
			Id:     "732211",
			CityId: "7322",
			Name:   "Baebunta",
		},
		{
			Id:     "732212",
			CityId: "7322",
			Name:   "Tana Lili",
		},
		{
			Id:     "732213",
			CityId: "7322",
			Name:   "Sukamaju Selatan",
		},
		{
			Id:     "732214",
			CityId: "7322",
			Name:   "Baebunta Selatan",
		},
		{
			Id:     "732215",
			CityId: "7322",
			Name:   "Sabbang Selatan",
		},
		{
			Id:     "732401",
			CityId: "7324",
			Name:   "Mangkutana",
		},
		{
			Id:     "732402",
			CityId: "7324",
			Name:   "Nuha",
		},
		{
			Id:     "732403",
			CityId: "7324",
			Name:   "Towuti",
		},
		{
			Id:     "732404",
			CityId: "7324",
			Name:   "Malili",
		},
		{
			Id:     "732405",
			CityId: "7324",
			Name:   "Angkona",
		},
		{
			Id:     "732406",
			CityId: "7324",
			Name:   "Wotu",
		},
		{
			Id:     "732407",
			CityId: "7324",
			Name:   "Burau",
		},
		{
			Id:     "732408",
			CityId: "7324",
			Name:   "Tomoni",
		},
		{
			Id:     "732409",
			CityId: "7324",
			Name:   "Tomoni Timur",
		},
		{
			Id:     "732410",
			CityId: "7324",
			Name:   "Kalaena",
		},
		{
			Id:     "732411",
			CityId: "7324",
			Name:   "Wasuponda",
		},
		{
			Id:     "732601",
			CityId: "7326",
			Name:   "Rantepao",
		},
		{
			Id:     "732602",
			CityId: "7326",
			Name:   "Sesean",
		},
		{
			Id:     "732603",
			CityId: "7326",
			Name:   "Nanggala",
		},
		{
			Id:     "732604",
			CityId: "7326",
			Name:   "Rindingallo",
		},
		{
			Id:     "732605",
			CityId: "7326",
			Name:   "Buntao",
		},
		{
			Id:     "732606",
			CityId: "7326",
			Name:   "Sa'dan",
		},
		{
			Id:     "732607",
			CityId: "7326",
			Name:   "Sanggalangi",
		},
		{
			Id:     "732608",
			CityId: "7326",
			Name:   "Sopai",
		},
		{
			Id:     "732609",
			CityId: "7326",
			Name:   "Tikala",
		},
		{
			Id:     "732610",
			CityId: "7326",
			Name:   "Balusu",
		},
		{
			Id:     "732611",
			CityId: "7326",
			Name:   "Tallunglipu",
		},
		{
			Id:     "732612",
			CityId: "7326",
			Name:   "Dende' Piongan Napo",
		},
		{
			Id:     "732613",
			CityId: "7326",
			Name:   "Buntu Pepasan",
		},
		{
			Id:     "732614",
			CityId: "7326",
			Name:   "Baruppu",
		},
		{
			Id:     "732615",
			CityId: "7326",
			Name:   "Kesu",
		},
		{
			Id:     "732616",
			CityId: "7326",
			Name:   "Tondon",
		},
		{
			Id:     "732617",
			CityId: "7326",
			Name:   "Bangkelekila",
		},
		{
			Id:     "732618",
			CityId: "7326",
			Name:   "Rantebua",
		},
		{
			Id:     "732619",
			CityId: "7326",
			Name:   "Sesean Suloara",
		},
		{
			Id:     "732620",
			CityId: "7326",
			Name:   "Kapala Pitu",
		},
		{
			Id:     "732621",
			CityId: "7326",
			Name:   "Awan Rante Karua",
		},
		{
			Id:     "737101",
			CityId: "7371",
			Name:   "Mariso",
		},
		{
			Id:     "737102",
			CityId: "7371",
			Name:   "Mamajang",
		},
		{
			Id:     "737103",
			CityId: "7371",
			Name:   "Makassar",
		},
		{
			Id:     "737104",
			CityId: "7371",
			Name:   "Ujung Pandang",
		},
		{
			Id:     "737105",
			CityId: "7371",
			Name:   "Wajo",
		},
		{
			Id:     "737106",
			CityId: "7371",
			Name:   "Bontoala",
		},
		{
			Id:     "737107",
			CityId: "7371",
			Name:   "Tallo",
		},
		{
			Id:     "737108",
			CityId: "7371",
			Name:   "Ujung Tanah",
		},
		{
			Id:     "737109",
			CityId: "7371",
			Name:   "Panakkukang",
		},
		{
			Id:     "737110",
			CityId: "7371",
			Name:   "Tamalate",
		},
		{
			Id:     "737111",
			CityId: "7371",
			Name:   "Biringkanaya",
		},
		{
			Id:     "737112",
			CityId: "7371",
			Name:   "Manggala",
		},
		{
			Id:     "737113",
			CityId: "7371",
			Name:   "Rappocini",
		},
		{
			Id:     "737114",
			CityId: "7371",
			Name:   "Tamalanrea",
		},
		{
			Id:     "737115",
			CityId: "7371",
			Name:   "Kepulauan Sangkarrang",
		},
		{
			Id:     "737201",
			CityId: "7372",
			Name:   "Bacukiki",
		},
		{
			Id:     "737202",
			CityId: "7372",
			Name:   "Ujung",
		},
		{
			Id:     "737203",
			CityId: "7372",
			Name:   "Soreang",
		},
		{
			Id:     "737204",
			CityId: "7372",
			Name:   "Bacukiki Barat",
		},
		{
			Id:     "737301",
			CityId: "7373",
			Name:   "Wara",
		},
		{
			Id:     "737302",
			CityId: "7373",
			Name:   "Wara Utara",
		},
		{
			Id:     "737303",
			CityId: "7373",
			Name:   "Wara Selatan",
		},
		{
			Id:     "737304",
			CityId: "7373",
			Name:   "Telluwanua",
		},
		{
			Id:     "737305",
			CityId: "7373",
			Name:   "Wara Timur",
		},
		{
			Id:     "737306",
			CityId: "7373",
			Name:   "Wara Barat",
		},
		{
			Id:     "737307",
			CityId: "7373",
			Name:   "Sendana",
		},
		{
			Id:     "737308",
			CityId: "7373",
			Name:   "Mungkajang",
		},
		{
			Id:     "737309",
			CityId: "7373",
			Name:   "Bara",
		},
		{
			Id:     "740101",
			CityId: "7401",
			Name:   "Wundulako",
		},
		{
			Id:     "740104",
			CityId: "7401",
			Name:   "Kolaka",
		},
		{
			Id:     "740107",
			CityId: "7401",
			Name:   "Pomalaa",
		},
		{
			Id:     "740108",
			CityId: "7401",
			Name:   "Watubangga",
		},
		{
			Id:     "740110",
			CityId: "7401",
			Name:   "Wolo",
		},
		{
			Id:     "740112",
			CityId: "7401",
			Name:   "Baula",
		},
		{
			Id:     "740114",
			CityId: "7401",
			Name:   "Latambaga",
		},
		{
			Id:     "740118",
			CityId: "7401",
			Name:   "Tanggetada",
		},
		{
			Id:     "740120",
			CityId: "7401",
			Name:   "Samaturu",
		},
		{
			Id:     "740124",
			CityId: "7401",
			Name:   "Toari",
		},
		{
			Id:     "740125",
			CityId: "7401",
			Name:   "Polinggona",
		},
		{
			Id:     "740127",
			CityId: "7401",
			Name:   "Iwoimendaa",
		},
		{
			Id:     "740201",
			CityId: "7402",
			Name:   "Lambuya",
		},
		{
			Id:     "740202",
			CityId: "7402",
			Name:   "Unaaha",
		},
		{
			Id:     "740203",
			CityId: "7402",
			Name:   "Wawotobi",
		},
		{
			Id:     "740204",
			CityId: "7402",
			Name:   "Pondidaha",
		},
		{
			Id:     "740205",
			CityId: "7402",
			Name:   "Sampara",
		},
		{
			Id:     "740210",
			CityId: "7402",
			Name:   "Abuki",
		},
		{
			Id:     "740211",
			CityId: "7402",
			Name:   "Soropia",
		},
		{
			Id:     "740215",
			CityId: "7402",
			Name:   "Tongauna",
		},
		{
			Id:     "740216",
			CityId: "7402",
			Name:   "Latoma",
		},
		{
			Id:     "740217",
			CityId: "7402",
			Name:   "Puriala",
		},
		{
			Id:     "740218",
			CityId: "7402",
			Name:   "Uepai",
		},
		{
			Id:     "740219",
			CityId: "7402",
			Name:   "Wonggeduku",
		},
		{
			Id:     "740220",
			CityId: "7402",
			Name:   "Besulutu",
		},
		{
			Id:     "740221",
			CityId: "7402",
			Name:   "Bondoala",
		},
		{
			Id:     "740223",
			CityId: "7402",
			Name:   "Routa",
		},
		{
			Id:     "740224",
			CityId: "7402",
			Name:   "Anggaberi",
		},
		{
			Id:     "740225",
			CityId: "7402",
			Name:   "Meluhu",
		},
		{
			Id:     "740228",
			CityId: "7402",
			Name:   "Amonggedo",
		},
		{
			Id:     "740231",
			CityId: "7402",
			Name:   "Asinua",
		},
		{
			Id:     "740232",
			CityId: "7402",
			Name:   "Konawe",
		},
		{
			Id:     "740233",
			CityId: "7402",
			Name:   "Kapoiala",
		},
		{
			Id:     "740236",
			CityId: "7402",
			Name:   "Lalonggasumeeto",
		},
		{
			Id:     "740237",
			CityId: "7402",
			Name:   "Onembute",
		},
		{
			Id:     "740238",
			CityId: "7402",
			Name:   "Padangguni",
		},
		{
			Id:     "740239",
			CityId: "7402",
			Name:   "Morosi",
		},
		{
			Id:     "740240",
			CityId: "7402",
			Name:   "Anggalomoare",
		},
		{
			Id:     "740241",
			CityId: "7402",
			Name:   "Wonggeduku Barat",
		},
		{
			Id:     "740242",
			CityId: "7402",
			Name:   "Tongauna Utara",
		},
		{
			Id:     "740306",
			CityId: "7403",
			Name:   "Napabalano",
		},
		{
			Id:     "740307",
			CityId: "7403",
			Name:   "Maligano",
		},
		{
			Id:     "740313",
			CityId: "7403",
			Name:   "Wakorumba Selatan",
		},
		{
			Id:     "740314",
			CityId: "7403",
			Name:   "Lasalepa",
		},
		{
			Id:     "740315",
			CityId: "7403",
			Name:   "Batalaiworu",
		},
		{
			Id:     "740316",
			CityId: "7403",
			Name:   "Katobu",
		},
		{
			Id:     "740317",
			CityId: "7403",
			Name:   "Duruka",
		},
		{
			Id:     "740318",
			CityId: "7403",
			Name:   "Lohia",
		},
		{
			Id:     "740319",
			CityId: "7403",
			Name:   "Watopute",
		},
		{
			Id:     "740320",
			CityId: "7403",
			Name:   "Kontunaga",
		},
		{
			Id:     "740323",
			CityId: "7403",
			Name:   "Kabangka",
		},
		{
			Id:     "740324",
			CityId: "7403",
			Name:   "Kabawo",
		},
		{
			Id:     "740325",
			CityId: "7403",
			Name:   "Parigi",
		},
		{
			Id:     "740326",
			CityId: "7403",
			Name:   "Bone",
		},
		{
			Id:     "740327",
			CityId: "7403",
			Name:   "Tongkuno",
		},
		{
			Id:     "740328",
			CityId: "7403",
			Name:   "Pasir Putih",
		},
		{
			Id:     "740330",
			CityId: "7403",
			Name:   "Kontu Kowuna",
		},
		{
			Id:     "740331",
			CityId: "7403",
			Name:   "Marobo",
		},
		{
			Id:     "740332",
			CityId: "7403",
			Name:   "Tongkuno Selatan",
		},
		{
			Id:     "740333",
			CityId: "7403",
			Name:   "Pasi Kolaga",
		},
		{
			Id:     "740334",
			CityId: "7403",
			Name:   "Batukara",
		},
		{
			Id:     "740337",
			CityId: "7403",
			Name:   "Towea",
		},
		{
			Id:     "740411",
			CityId: "7404",
			Name:   "Pasarwajo",
		},
		{
			Id:     "740422",
			CityId: "7404",
			Name:   "Kapontori",
		},
		{
			Id:     "740423",
			CityId: "7404",
			Name:   "Lasalimu",
		},
		{
			Id:     "740424",
			CityId: "7404",
			Name:   "Lasalimu Selatan",
		},
		{
			Id:     "740427",
			CityId: "7404",
			Name:   "Siotapina",
		},
		{
			Id:     "740428",
			CityId: "7404",
			Name:   "Wolowa",
		},
		{
			Id:     "740429",
			CityId: "7404",
			Name:   "Wabula",
		},
		{
			Id:     "740501",
			CityId: "7405",
			Name:   "Tinanggea",
		},
		{
			Id:     "740502",
			CityId: "7405",
			Name:   "Angata",
		},
		{
			Id:     "740503",
			CityId: "7405",
			Name:   "Andoolo",
		},
		{
			Id:     "740504",
			CityId: "7405",
			Name:   "Palangga",
		},
		{
			Id:     "740505",
			CityId: "7405",
			Name:   "Landono",
		},
		{
			Id:     "740506",
			CityId: "7405",
			Name:   "Lainea",
		},
		{
			Id:     "740507",
			CityId: "7405",
			Name:   "Konda",
		},
		{
			Id:     "740508",
			CityId: "7405",
			Name:   "Ranomeeto",
		},
		{
			Id:     "740509",
			CityId: "7405",
			Name:   "Kolono",
		},
		{
			Id:     "740510",
			CityId: "7405",
			Name:   "Moramo",
		},
		{
			Id:     "740511",
			CityId: "7405",
			Name:   "Laonti",
		},
		{
			Id:     "740512",
			CityId: "7405",
			Name:   "Lalembuu",
		},
		{
			Id:     "740513",
			CityId: "7405",
			Name:   "Benua",
		},
		{
			Id:     "740514",
			CityId: "7405",
			Name:   "Palangga Selatan",
		},
		{
			Id:     "740515",
			CityId: "7405",
			Name:   "Mowila",
		},
		{
			Id:     "740516",
			CityId: "7405",
			Name:   "Moramo Utara",
		},
		{
			Id:     "740517",
			CityId: "7405",
			Name:   "Buke",
		},
		{
			Id:     "740518",
			CityId: "7405",
			Name:   "Wolasi",
		},
		{
			Id:     "740519",
			CityId: "7405",
			Name:   "Laeya",
		},
		{
			Id:     "740520",
			CityId: "7405",
			Name:   "Baito",
		},
		{
			Id:     "740521",
			CityId: "7405",
			Name:   "Basala",
		},
		{
			Id:     "740522",
			CityId: "7405",
			Name:   "Ranomeeto Barat",
		},
		{
			Id:     "740523",
			CityId: "7405",
			Name:   "Kolono Timur",
		},
		{
			Id:     "740524",
			CityId: "7405",
			Name:   "Sabulakoa",
		},
		{
			Id:     "740525",
			CityId: "7405",
			Name:   "Andoolo Barat",
		},
		{
			Id:     "740601",
			CityId: "7406",
			Name:   "Poleang",
		},
		{
			Id:     "740602",
			CityId: "7406",
			Name:   "Poleang Timur",
		},
		{
			Id:     "740603",
			CityId: "7406",
			Name:   "Rarowatu",
		},
		{
			Id:     "740604",
			CityId: "7406",
			Name:   "Rumbia",
		},
		{
			Id:     "740605",
			CityId: "7406",
			Name:   "Kabaena",
		},
		{
			Id:     "740606",
			CityId: "7406",
			Name:   "Kabaena Timur",
		},
		{
			Id:     "740607",
			CityId: "7406",
			Name:   "Poleang Barat",
		},
		{
			Id:     "740608",
			CityId: "7406",
			Name:   "Mata Oleo",
		},
		{
			Id:     "740609",
			CityId: "7406",
			Name:   "Rarowatu Utara",
		},
		{
			Id:     "740610",
			CityId: "7406",
			Name:   "Poleang Utara",
		},
		{
			Id:     "740611",
			CityId: "7406",
			Name:   "Poleang Selatan",
		},
		{
			Id:     "740612",
			CityId: "7406",
			Name:   "Poleang Tenggara",
		},
		{
			Id:     "740613",
			CityId: "7406",
			Name:   "Kabaena Selatan",
		},
		{
			Id:     "740614",
			CityId: "7406",
			Name:   "Kabaena Barat",
		},
		{
			Id:     "740615",
			CityId: "7406",
			Name:   "Kabaena Utara",
		},
		{
			Id:     "740616",
			CityId: "7406",
			Name:   "Kabaena Tengah",
		},
		{
			Id:     "740617",
			CityId: "7406",
			Name:   "Kep. Masaloka Raya",
		},
		{
			Id:     "740618",
			CityId: "7406",
			Name:   "Rumbia Tengah",
		},
		{
			Id:     "740619",
			CityId: "7406",
			Name:   "Poleang Tengah",
		},
		{
			Id:     "740620",
			CityId: "7406",
			Name:   "Tontonunu",
		},
		{
			Id:     "740621",
			CityId: "7406",
			Name:   "Lantari Jaya",
		},
		{
			Id:     "740622",
			CityId: "7406",
			Name:   "Mata Usu",
		},
		{
			Id:     "740701",
			CityId: "7407",
			Name:   "Wangi-Wangi",
		},
		{
			Id:     "740702",
			CityId: "7407",
			Name:   "Kaledupa",
		},
		{
			Id:     "740703",
			CityId: "7407",
			Name:   "Tomia",
		},
		{
			Id:     "740704",
			CityId: "7407",
			Name:   "Binongko",
		},
		{
			Id:     "740705",
			CityId: "7407",
			Name:   "Wangi Wangi Selatan",
		},
		{
			Id:     "740706",
			CityId: "7407",
			Name:   "Kaledupa Selatan",
		},
		{
			Id:     "740707",
			CityId: "7407",
			Name:   "Tomia Timur",
		},
		{
			Id:     "740708",
			CityId: "7407",
			Name:   "Togo Binongko",
		},
		{
			Id:     "740801",
			CityId: "7408",
			Name:   "Lasusua",
		},
		{
			Id:     "740802",
			CityId: "7408",
			Name:   "Pakue",
		},
		{
			Id:     "740803",
			CityId: "7408",
			Name:   "Batu Putih",
		},
		{
			Id:     "740804",
			CityId: "7408",
			Name:   "Rante Angin",
		},
		{
			Id:     "740805",
			CityId: "7408",
			Name:   "Kodeoha",
		},
		{
			Id:     "740806",
			CityId: "7408",
			Name:   "Ngapa",
		},
		{
			Id:     "740807",
			CityId: "7408",
			Name:   "Wawo",
		},
		{
			Id:     "740808",
			CityId: "7408",
			Name:   "Lambai",
		},
		{
			Id:     "740809",
			CityId: "7408",
			Name:   "Watunohu",
		},
		{
			Id:     "740810",
			CityId: "7408",
			Name:   "Pakue Tengah",
		},
		{
			Id:     "740811",
			CityId: "7408",
			Name:   "Pakue Utara",
		},
		{
			Id:     "740812",
			CityId: "7408",
			Name:   "Porehu",
		},
		{
			Id:     "740813",
			CityId: "7408",
			Name:   "Katoi",
		},
		{
			Id:     "740814",
			CityId: "7408",
			Name:   "Tiwu",
		},
		{
			Id:     "740815",
			CityId: "7408",
			Name:   "Tolala",
		},
		{
			Id:     "740901",
			CityId: "7409",
			Name:   "Asera",
		},
		{
			Id:     "740902",
			CityId: "7409",
			Name:   "Wiwirano",
		},
		{
			Id:     "740903",
			CityId: "7409",
			Name:   "Langgikima",
		},
		{
			Id:     "740904",
			CityId: "7409",
			Name:   "Molawe",
		},
		{
			Id:     "740905",
			CityId: "7409",
			Name:   "Lasolo",
		},
		{
			Id:     "740906",
			CityId: "7409",
			Name:   "Lembo",
		},
		{
			Id:     "740907",
			CityId: "7409",
			Name:   "Sawa",
		},
		{
			Id:     "740908",
			CityId: "7409",
			Name:   "Oheo",
		},
		{
			Id:     "740909",
			CityId: "7409",
			Name:   "Andowia",
		},
		{
			Id:     "740910",
			CityId: "7409",
			Name:   "Motui",
		},
		{
			Id:     "740911",
			CityId: "7409",
			Name:   "Wawolesea",
		},
		{
			Id:     "740912",
			CityId: "7409",
			Name:   "Lasolo Kepulauan",
		},
		{
			Id:     "740913",
			CityId: "7409",
			Name:   "Landawe",
		},
		{
			Id:     "741001",
			CityId: "7410",
			Name:   "Kulisusu",
		},
		{
			Id:     "741002",
			CityId: "7410",
			Name:   "Kambowa",
		},
		{
			Id:     "741003",
			CityId: "7410",
			Name:   "Bonegunu",
		},
		{
			Id:     "741004",
			CityId: "7410",
			Name:   "Kulisusu Barat",
		},
		{
			Id:     "741005",
			CityId: "7410",
			Name:   "Kulisusu Utara",
		},
		{
			Id:     "741006",
			CityId: "7410",
			Name:   "Wakorumba Utara",
		},
		{
			Id:     "741101",
			CityId: "7411",
			Name:   "Tirawuta",
		},
		{
			Id:     "741102",
			CityId: "7411",
			Name:   "Loea",
		},
		{
			Id:     "741103",
			CityId: "7411",
			Name:   "Ladongi",
		},
		{
			Id:     "741104",
			CityId: "7411",
			Name:   "Poli Polia",
		},
		{
			Id:     "741105",
			CityId: "7411",
			Name:   "Lambandia",
		},
		{
			Id:     "741106",
			CityId: "7411",
			Name:   "Lalolae",
		},
		{
			Id:     "741107",
			CityId: "7411",
			Name:   "Mowewe",
		},
		{
			Id:     "741108",
			CityId: "7411",
			Name:   "Uluiwoi",
		},
		{
			Id:     "741109",
			CityId: "7411",
			Name:   "Tinondo",
		},
		{
			Id:     "741110",
			CityId: "7411",
			Name:   "Aere",
		},
		{
			Id:     "741111",
			CityId: "7411",
			Name:   "Ueesi",
		},
		{
			Id:     "741112",
			CityId: "7411",
			Name:   "Dangia",
		},
		{
			Id:     "741201",
			CityId: "7412",
			Name:   "Wawonii Barat",
		},
		{
			Id:     "741202",
			CityId: "7412",
			Name:   "Wawonii Utara",
		},
		{
			Id:     "741203",
			CityId: "7412",
			Name:   "Wawonii Timur Laut",
		},
		{
			Id:     "741204",
			CityId: "7412",
			Name:   "Wawonii Timur",
		},
		{
			Id:     "741205",
			CityId: "7412",
			Name:   "Wawonii Tenggara",
		},
		{
			Id:     "741206",
			CityId: "7412",
			Name:   "Wawonii Selatan",
		},
		{
			Id:     "741207",
			CityId: "7412",
			Name:   "Wawonii Tengah",
		},
		{
			Id:     "741301",
			CityId: "7413",
			Name:   "Sawerigadi",
		},
		{
			Id:     "741302",
			CityId: "7413",
			Name:   "Barangka",
		},
		{
			Id:     "741303",
			CityId: "7413",
			Name:   "Lawa",
		},
		{
			Id:     "741304",
			CityId: "7413",
			Name:   "Wadaga",
		},
		{
			Id:     "741305",
			CityId: "7413",
			Name:   "Tiworo Selatan",
		},
		{
			Id:     "741306",
			CityId: "7413",
			Name:   "Maginti",
		},
		{
			Id:     "741307",
			CityId: "7413",
			Name:   "Tiworo Tengah",
		},
		{
			Id:     "741308",
			CityId: "7413",
			Name:   "Tiworo Utara",
		},
		{
			Id:     "741309",
			CityId: "7413",
			Name:   "Tiworo Kepulauan",
		},
		{
			Id:     "741310",
			CityId: "7413",
			Name:   "Kusambi",
		},
		{
			Id:     "741311",
			CityId: "7413",
			Name:   "Napano Kusambi",
		},
		{
			Id:     "741401",
			CityId: "7414",
			Name:   "Lakudo",
		},
		{
			Id:     "741402",
			CityId: "7414",
			Name:   "Mawasangka Timur",
		},
		{
			Id:     "741403",
			CityId: "7414",
			Name:   "Mawasangka Tengah",
		},
		{
			Id:     "741404",
			CityId: "7414",
			Name:   "Mawasangka",
		},
		{
			Id:     "741405",
			CityId: "7414",
			Name:   "Talaga Raya",
		},
		{
			Id:     "741406",
			CityId: "7414",
			Name:   "Gu",
		},
		{
			Id:     "741407",
			CityId: "7414",
			Name:   "Sangia Wambulu",
		},
		{
			Id:     "741501",
			CityId: "7415",
			Name:   "Batauga",
		},
		{
			Id:     "741502",
			CityId: "7415",
			Name:   "Sampolawa",
		},
		{
			Id:     "741503",
			CityId: "7415",
			Name:   "Lapandewa",
		},
		{
			Id:     "741504",
			CityId: "7415",
			Name:   "Batu Atas",
		},
		{
			Id:     "741505",
			CityId: "7415",
			Name:   "Siompu Barat",
		},
		{
			Id:     "741506",
			CityId: "7415",
			Name:   "Siompu",
		},
		{
			Id:     "741507",
			CityId: "7415",
			Name:   "Kadatua",
		},
		{
			Id:     "747101",
			CityId: "7471",
			Name:   "Mandonga",
		},
		{
			Id:     "747102",
			CityId: "7471",
			Name:   "Kendari",
		},
		{
			Id:     "747103",
			CityId: "7471",
			Name:   "Baruga",
		},
		{
			Id:     "747104",
			CityId: "7471",
			Name:   "Poasia",
		},
		{
			Id:     "747105",
			CityId: "7471",
			Name:   "Kendari Barat",
		},
		{
			Id:     "747106",
			CityId: "7471",
			Name:   "Abeli",
		},
		{
			Id:     "747107",
			CityId: "7471",
			Name:   "Wua-Wua",
		},
		{
			Id:     "747108",
			CityId: "7471",
			Name:   "Kadia",
		},
		{
			Id:     "747109",
			CityId: "7471",
			Name:   "Puuwatu",
		},
		{
			Id:     "747110",
			CityId: "7471",
			Name:   "Kambu",
		},
		{
			Id:     "747111",
			CityId: "7471",
			Name:   "Nambo",
		},
		{
			Id:     "747201",
			CityId: "7472",
			Name:   "Betoambari",
		},
		{
			Id:     "747202",
			CityId: "7472",
			Name:   "Wolio",
		},
		{
			Id:     "747203",
			CityId: "7472",
			Name:   "Sorawolio",
		},
		{
			Id:     "747204",
			CityId: "7472",
			Name:   "Bungi",
		},
		{
			Id:     "747205",
			CityId: "7472",
			Name:   "Kokalukuna",
		},
		{
			Id:     "747206",
			CityId: "7472",
			Name:   "Murhum",
		},
		{
			Id:     "747207",
			CityId: "7472",
			Name:   "Lea-Lea",
		},
		{
			Id:     "747208",
			CityId: "7472",
			Name:   "Batupoaro",
		},
		{
			Id:     "750101",
			CityId: "7501",
			Name:   "Limboto",
		},
		{
			Id:     "750102",
			CityId: "7501",
			Name:   "Telaga",
		},
		{
			Id:     "750103",
			CityId: "7501",
			Name:   "Batudaa",
		},
		{
			Id:     "750104",
			CityId: "7501",
			Name:   "Tibawa",
		},
		{
			Id:     "750105",
			CityId: "7501",
			Name:   "Batudaa Pantai",
		},
		{
			Id:     "750109",
			CityId: "7501",
			Name:   "Boliyohuto",
		},
		{
			Id:     "750110",
			CityId: "7501",
			Name:   "Telaga Biru",
		},
		{
			Id:     "750111",
			CityId: "7501",
			Name:   "Bongomeme",
		},
		{
			Id:     "750113",
			CityId: "7501",
			Name:   "Tolangohula",
		},
		{
			Id:     "750114",
			CityId: "7501",
			Name:   "Mootilango",
		},
		{
			Id:     "750116",
			CityId: "7501",
			Name:   "Pulubala",
		},
		{
			Id:     "750117",
			CityId: "7501",
			Name:   "Limboto Barat",
		},
		{
			Id:     "750118",
			CityId: "7501",
			Name:   "Tilango",
		},
		{
			Id:     "750119",
			CityId: "7501",
			Name:   "Tabongo",
		},
		{
			Id:     "750120",
			CityId: "7501",
			Name:   "Biluhu",
		},
		{
			Id:     "750121",
			CityId: "7501",
			Name:   "Asparaga",
		},
		{
			Id:     "750122",
			CityId: "7501",
			Name:   "Talaga Jaya",
		},
		{
			Id:     "750123",
			CityId: "7501",
			Name:   "Bilato",
		},
		{
			Id:     "750124",
			CityId: "7501",
			Name:   "Dungaliyo",
		},
		{
			Id:     "750201",
			CityId: "7502",
			Name:   "Paguyaman",
		},
		{
			Id:     "750202",
			CityId: "7502",
			Name:   "Wonosari",
		},
		{
			Id:     "750203",
			CityId: "7502",
			Name:   "Dulupi",
		},
		{
			Id:     "750204",
			CityId: "7502",
			Name:   "Tilamuta",
		},
		{
			Id:     "750205",
			CityId: "7502",
			Name:   "Mananggu",
		},
		{
			Id:     "750206",
			CityId: "7502",
			Name:   "Botumoito",
		},
		{
			Id:     "750207",
			CityId: "7502",
			Name:   "Paguyaman Pantai",
		},
		{
			Id:     "750301",
			CityId: "7503",
			Name:   "Tapa",
		},
		{
			Id:     "750302",
			CityId: "7503",
			Name:   "Kabila",
		},
		{
			Id:     "750303",
			CityId: "7503",
			Name:   "Suwawa",
		},
		{
			Id:     "750304",
			CityId: "7503",
			Name:   "Bonepantai",
		},
		{
			Id:     "750305",
			CityId: "7503",
			Name:   "Bulango Utara",
		},
		{
			Id:     "750306",
			CityId: "7503",
			Name:   "Tilongkabila",
		},
		{
			Id:     "750307",
			CityId: "7503",
			Name:   "Botupingge",
		},
		{
			Id:     "750308",
			CityId: "7503",
			Name:   "Kabila Bone",
		},
		{
			Id:     "750309",
			CityId: "7503",
			Name:   "Bone",
		},
		{
			Id:     "750310",
			CityId: "7503",
			Name:   "Bone Raya",
		},
		{
			Id:     "750311",
			CityId: "7503",
			Name:   "Suwawa Timur",
		},
		{
			Id:     "750312",
			CityId: "7503",
			Name:   "Suwawa Selatan",
		},
		{
			Id:     "750313",
			CityId: "7503",
			Name:   "Suwawa Tengah",
		},
		{
			Id:     "750314",
			CityId: "7503",
			Name:   "Bulango Ulu",
		},
		{
			Id:     "750315",
			CityId: "7503",
			Name:   "Bulango Selatan",
		},
		{
			Id:     "750316",
			CityId: "7503",
			Name:   "Bulango Timur",
		},
		{
			Id:     "750317",
			CityId: "7503",
			Name:   "Bulawa",
		},
		{
			Id:     "750318",
			CityId: "7503",
			Name:   "Pinogu",
		},
		{
			Id:     "750401",
			CityId: "7504",
			Name:   "Popayato",
		},
		{
			Id:     "750402",
			CityId: "7504",
			Name:   "Lemito",
		},
		{
			Id:     "750403",
			CityId: "7504",
			Name:   "Randangan",
		},
		{
			Id:     "750404",
			CityId: "7504",
			Name:   "Marisa",
		},
		{
			Id:     "750405",
			CityId: "7504",
			Name:   "Paguat",
		},
		{
			Id:     "750406",
			CityId: "7504",
			Name:   "Patilanggio",
		},
		{
			Id:     "750407",
			CityId: "7504",
			Name:   "Taluditi",
		},
		{
			Id:     "750408",
			CityId: "7504",
			Name:   "Dengilo",
		},
		{
			Id:     "750409",
			CityId: "7504",
			Name:   "Buntulia",
		},
		{
			Id:     "750410",
			CityId: "7504",
			Name:   "Duhiadaa",
		},
		{
			Id:     "750411",
			CityId: "7504",
			Name:   "Wanggarasi",
		},
		{
			Id:     "750412",
			CityId: "7504",
			Name:   "Popayato Timur",
		},
		{
			Id:     "750413",
			CityId: "7504",
			Name:   "Popayato Barat",
		},
		{
			Id:     "750501",
			CityId: "7505",
			Name:   "Atinggola",
		},
		{
			Id:     "750502",
			CityId: "7505",
			Name:   "Kwandang",
		},
		{
			Id:     "750503",
			CityId: "7505",
			Name:   "Anggrek",
		},
		{
			Id:     "750504",
			CityId: "7505",
			Name:   "Sumalata",
		},
		{
			Id:     "750505",
			CityId: "7505",
			Name:   "Tolinggula",
		},
		{
			Id:     "750506",
			CityId: "7505",
			Name:   "Gentuma Raya",
		},
		{
			Id:     "750507",
			CityId: "7505",
			Name:   "Tomolito",
		},
		{
			Id:     "750508",
			CityId: "7505",
			Name:   "Ponelo Kepulauan",
		},
		{
			Id:     "750509",
			CityId: "7505",
			Name:   "Monano",
		},
		{
			Id:     "750510",
			CityId: "7505",
			Name:   "Biau",
		},
		{
			Id:     "750511",
			CityId: "7505",
			Name:   "Sumalata Timur",
		},
		{
			Id:     "757101",
			CityId: "7571",
			Name:   "Kota Barat",
		},
		{
			Id:     "757102",
			CityId: "7571",
			Name:   "Kota Selatan",
		},
		{
			Id:     "757103",
			CityId: "7571",
			Name:   "Kota Utara",
		},
		{
			Id:     "757104",
			CityId: "7571",
			Name:   "Dungingi",
		},
		{
			Id:     "757105",
			CityId: "7571",
			Name:   "Kota Timur",
		},
		{
			Id:     "757106",
			CityId: "7571",
			Name:   "Kota Tengah",
		},
		{
			Id:     "757107",
			CityId: "7571",
			Name:   "Sipatana",
		},
		{
			Id:     "757108",
			CityId: "7571",
			Name:   "Dumbo Raya",
		},
		{
			Id:     "757109",
			CityId: "7571",
			Name:   "Hulonthalangi",
		},
		{
			Id:     "760101",
			CityId: "7601",
			Name:   "Bambalamotu",
		},
		{
			Id:     "760102",
			CityId: "7601",
			Name:   "Pasangkayu",
		},
		{
			Id:     "760103",
			CityId: "7601",
			Name:   "Baras",
		},
		{
			Id:     "760104",
			CityId: "7601",
			Name:   "Sarudu",
		},
		{
			Id:     "760105",
			CityId: "7601",
			Name:   "Dapurang",
		},
		{
			Id:     "760106",
			CityId: "7601",
			Name:   "Duripoku",
		},
		{
			Id:     "760107",
			CityId: "7601",
			Name:   "Bulu Taba",
		},
		{
			Id:     "760108",
			CityId: "7601",
			Name:   "Tikke Raya",
		},
		{
			Id:     "760109",
			CityId: "7601",
			Name:   "Pedongga",
		},
		{
			Id:     "760110",
			CityId: "7601",
			Name:   "Bambaira",
		},
		{
			Id:     "760111",
			CityId: "7601",
			Name:   "Sarjo",
		},
		{
			Id:     "760112",
			CityId: "7601",
			Name:   "Lariang",
		},
		{
			Id:     "760201",
			CityId: "7602",
			Name:   "Mamuju",
		},
		{
			Id:     "760202",
			CityId: "7602",
			Name:   "Tapalang",
		},
		{
			Id:     "760203",
			CityId: "7602",
			Name:   "Kalukku",
		},
		{
			Id:     "760204",
			CityId: "7602",
			Name:   "Kalumpang",
		},
		{
			Id:     "760207",
			CityId: "7602",
			Name:   "Papalang",
		},
		{
			Id:     "760208",
			CityId: "7602",
			Name:   "Sampaga",
		},
		{
			Id:     "760211",
			CityId: "7602",
			Name:   "Tommo",
		},
		{
			Id:     "760212",
			CityId: "7602",
			Name:   "Simboro dan Kepulauan",
		},
		{
			Id:     "760213",
			CityId: "7602",
			Name:   "Tapalang Barat",
		},
		{
			Id:     "760215",
			CityId: "7602",
			Name:   "Bonehau",
		},
		{
			Id:     "760216",
			CityId: "7602",
			Name:   "Kep. Bala Balakang",
		},
		{
			Id:     "760301",
			CityId: "7603",
			Name:   "Mambi",
		},
		{
			Id:     "760302",
			CityId: "7603",
			Name:   "Aralle",
		},
		{
			Id:     "760303",
			CityId: "7603",
			Name:   "Mamasa",
		},
		{
			Id:     "760304",
			CityId: "7603",
			Name:   "Pana",
		},
		{
			Id:     "760305",
			CityId: "7603",
			Name:   "Tabulahan",
		},
		{
			Id:     "760306",
			CityId: "7603",
			Name:   "Sumarorong",
		},
		{
			Id:     "760307",
			CityId: "7603",
			Name:   "Messawa",
		},
		{
			Id:     "760308",
			CityId: "7603",
			Name:   "Sesenapadang",
		},
		{
			Id:     "760309",
			CityId: "7603",
			Name:   "Tanduk Kalua",
		},
		{
			Id:     "760310",
			CityId: "7603",
			Name:   "Tabang",
		},
		{
			Id:     "760311",
			CityId: "7603",
			Name:   "Bambang",
		},
		{
			Id:     "760312",
			CityId: "7603",
			Name:   "Balla",
		},
		{
			Id:     "760313",
			CityId: "7603",
			Name:   "Nosu",
		},
		{
			Id:     "760314",
			CityId: "7603",
			Name:   "Tawalian",
		},
		{
			Id:     "760315",
			CityId: "7603",
			Name:   "Rantebulahan Timur",
		},
		{
			Id:     "760316",
			CityId: "7603",
			Name:   "Buntumalangka",
		},
		{
			Id:     "760317",
			CityId: "7603",
			Name:   "Mehalaan",
		},
		{
			Id:     "760401",
			CityId: "7604",
			Name:   "Tinambung",
		},
		{
			Id:     "760402",
			CityId: "7604",
			Name:   "Campalagian",
		},
		{
			Id:     "760403",
			CityId: "7604",
			Name:   "Wonomulyo",
		},
		{
			Id:     "760404",
			CityId: "7604",
			Name:   "Polewali",
		},
		{
			Id:     "760405",
			CityId: "7604",
			Name:   "Tutar",
		},
		{
			Id:     "760406",
			CityId: "7604",
			Name:   "Binuang",
		},
		{
			Id:     "760407",
			CityId: "7604",
			Name:   "Tapango",
		},
		{
			Id:     "760408",
			CityId: "7604",
			Name:   "Mapilli",
		},
		{
			Id:     "760409",
			CityId: "7604",
			Name:   "Matangnga",
		},
		{
			Id:     "760410",
			CityId: "7604",
			Name:   "Luyo",
		},
		{
			Id:     "760411",
			CityId: "7604",
			Name:   "Limboro",
		},
		{
			Id:     "760412",
			CityId: "7604",
			Name:   "Balanipa",
		},
		{
			Id:     "760413",
			CityId: "7604",
			Name:   "Anreapi",
		},
		{
			Id:     "760414",
			CityId: "7604",
			Name:   "Matakali",
		},
		{
			Id:     "760415",
			CityId: "7604",
			Name:   "Allu",
		},
		{
			Id:     "760416",
			CityId: "7604",
			Name:   "Bulo",
		},
		{
			Id:     "760501",
			CityId: "7605",
			Name:   "Banggae",
		},
		{
			Id:     "760502",
			CityId: "7605",
			Name:   "Pamboang",
		},
		{
			Id:     "760503",
			CityId: "7605",
			Name:   "Sendana",
		},
		{
			Id:     "760504",
			CityId: "7605",
			Name:   "Malunda",
		},
		{
			Id:     "760505",
			CityId: "7605",
			Name:   "Ulumanda",
		},
		{
			Id:     "760506",
			CityId: "7605",
			Name:   "Tammerodo Sendana",
		},
		{
			Id:     "760507",
			CityId: "7605",
			Name:   "Tubo Sendana",
		},
		{
			Id:     "760508",
			CityId: "7605",
			Name:   "Banggae Timur",
		},
		{
			Id:     "760601",
			CityId: "7606",
			Name:   "Tobadak",
		},
		{
			Id:     "760602",
			CityId: "7606",
			Name:   "Pangale",
		},
		{
			Id:     "760603",
			CityId: "7606",
			Name:   "Budong-Budong",
		},
		{
			Id:     "760604",
			CityId: "7606",
			Name:   "Topoyo",
		},
		{
			Id:     "760605",
			CityId: "7606",
			Name:   "Karossa",
		},
		{
			Id:     "810101",
			CityId: "8101",
			Name:   "Amahai",
		},
		{
			Id:     "810102",
			CityId: "8101",
			Name:   "Teon Nila Serua",
		},
		{
			Id:     "810106",
			CityId: "8101",
			Name:   "Seram Utara",
		},
		{
			Id:     "810109",
			CityId: "8101",
			Name:   "Banda",
		},
		{
			Id:     "810111",
			CityId: "8101",
			Name:   "Tehoru",
		},
		{
			Id:     "810112",
			CityId: "8101",
			Name:   "Saparua",
		},
		{
			Id:     "810113",
			CityId: "8101",
			Name:   "Pulau Haruku",
		},
		{
			Id:     "810114",
			CityId: "8101",
			Name:   "Salahutu",
		},
		{
			Id:     "810115",
			CityId: "8101",
			Name:   "Leihitu",
		},
		{
			Id:     "810116",
			CityId: "8101",
			Name:   "Nusa Laut",
		},
		{
			Id:     "810117",
			CityId: "8101",
			Name:   "Kota Masohi",
		},
		{
			Id:     "810120",
			CityId: "8101",
			Name:   "Seram Utara Barat",
		},
		{
			Id:     "810121",
			CityId: "8101",
			Name:   "Teluk Elpaputih",
		},
		{
			Id:     "810122",
			CityId: "8101",
			Name:   "Leihitu Barat",
		},
		{
			Id:     "810123",
			CityId: "8101",
			Name:   "Telutih",
		},
		{
			Id:     "810124",
			CityId: "8101",
			Name:   "Seram Utara Timur Seti",
		},
		{
			Id:     "810125",
			CityId: "8101",
			Name:   "Seram Utara Timur Kobi",
		},
		{
			Id:     "810126",
			CityId: "8101",
			Name:   "Saparua Timur",
		},
		{
			Id:     "810201",
			CityId: "8102",
			Name:   "Kei Kecil",
		},
		{
			Id:     "810203",
			CityId: "8102",
			Name:   "Kei Besar",
		},
		{
			Id:     "810204",
			CityId: "8102",
			Name:   "Kei Besar Selatan",
		},
		{
			Id:     "810205",
			CityId: "8102",
			Name:   "Kei Besar Utara Timur",
		},
		{
			Id:     "810213",
			CityId: "8102",
			Name:   "Kei Kecil Timur",
		},
		{
			Id:     "810214",
			CityId: "8102",
			Name:   "Kei Kecil Barat",
		},
		{
			Id:     "810215",
			CityId: "8102",
			Name:   "Manyeuw",
		},
		{
			Id:     "810216",
			CityId: "8102",
			Name:   "Hoat Sorbay",
		},
		{
			Id:     "810217",
			CityId: "8102",
			Name:   "Kei Besar Utara Barat",
		},
		{
			Id:     "810218",
			CityId: "8102",
			Name:   "Kei Besar Selatan Barat",
		},
		{
			Id:     "810219",
			CityId: "8102",
			Name:   "Kei Kecil Timur Selatan",
		},
		{
			Id:     "810301",
			CityId: "8103",
			Name:   "Tanimbar Selatan",
		},
		{
			Id:     "810302",
			CityId: "8103",
			Name:   "Selaru",
		},
		{
			Id:     "810303",
			CityId: "8103",
			Name:   "Wertamrian",
		},
		{
			Id:     "810304",
			CityId: "8103",
			Name:   "Wermaktian",
		},
		{
			Id:     "810305",
			CityId: "8103",
			Name:   "Tanimbar Utara",
		},
		{
			Id:     "810306",
			CityId: "8103",
			Name:   "Fordata",
		},
		{
			Id:     "810307",
			CityId: "8103",
			Name:   "Wuar Labobar",
		},
		{
			Id:     "810308",
			CityId: "8103",
			Name:   "Kormomolin",
		},
		{
			Id:     "810309",
			CityId: "8103",
			Name:   "Nirunmas",
		},
		{
			Id:     "810318",
			CityId: "8103",
			Name:   "Molu Maru",
		},
		{
			Id:     "810401",
			CityId: "8104",
			Name:   "Namlea",
		},
		{
			Id:     "810402",
			CityId: "8104",
			Name:   "Air Buaya",
		},
		{
			Id:     "810403",
			CityId: "8104",
			Name:   "Waeapo",
		},
		{
			Id:     "810406",
			CityId: "8104",
			Name:   "Waplau",
		},
		{
			Id:     "810410",
			CityId: "8104",
			Name:   "Batabual",
		},
		{
			Id:     "810411",
			CityId: "8104",
			Name:   "Lolong Guba",
		},
		{
			Id:     "810412",
			CityId: "8104",
			Name:   "Waelata",
		},
		{
			Id:     "810413",
			CityId: "8104",
			Name:   "Fena Leisela",
		},
		{
			Id:     "810414",
			CityId: "8104",
			Name:   "Teluk Kaiely",
		},
		{
			Id:     "810415",
			CityId: "8104",
			Name:   "Lilialy",
		},
		{
			Id:     "810501",
			CityId: "8105",
			Name:   "Bula",
		},
		{
			Id:     "810502",
			CityId: "8105",
			Name:   "Seram Timur",
		},
		{
			Id:     "810503",
			CityId: "8105",
			Name:   "Werinama",
		},
		{
			Id:     "810504",
			CityId: "8105",
			Name:   "Pulau Gorom",
		},
		{
			Id:     "810505",
			CityId: "8105",
			Name:   "Wakate",
		},
		{
			Id:     "810506",
			CityId: "8105",
			Name:   "Tutuk Tolu",
		},
		{
			Id:     "810507",
			CityId: "8105",
			Name:   "Siwalalat",
		},
		{
			Id:     "810508",
			CityId: "8105",
			Name:   "Kilmury",
		},
		{
			Id:     "810509",
			CityId: "8105",
			Name:   "Pulau Panjang",
		},
		{
			Id:     "810510",
			CityId: "8105",
			Name:   "Teor",
		},
		{
			Id:     "810511",
			CityId: "8105",
			Name:   "Gorom Timur",
		},
		{
			Id:     "810512",
			CityId: "8105",
			Name:   "Bula Barat",
		},
		{
			Id:     "810513",
			CityId: "8105",
			Name:   "Kian Darat",
		},
		{
			Id:     "810514",
			CityId: "8105",
			Name:   "Siritaun Wida Timur",
		},
		{
			Id:     "810515",
			CityId: "8105",
			Name:   "Teluk Waru",
		},
		{
			Id:     "810601",
			CityId: "8106",
			Name:   "Kairatu",
		},
		{
			Id:     "810602",
			CityId: "8106",
			Name:   "Seram Barat",
		},
		{
			Id:     "810603",
			CityId: "8106",
			Name:   "Taniwel",
		},
		{
			Id:     "810604",
			CityId: "8106",
			Name:   "Huamual Belakang",
		},
		{
			Id:     "810605",
			CityId: "8106",
			Name:   "Amalatu",
		},
		{
			Id:     "810606",
			CityId: "8106",
			Name:   "Inamosol",
		},
		{
			Id:     "810607",
			CityId: "8106",
			Name:   "Kairatu Barat",
		},
		{
			Id:     "810608",
			CityId: "8106",
			Name:   "Huamual",
		},
		{
			Id:     "810609",
			CityId: "8106",
			Name:   "Kepulauan Manipa",
		},
		{
			Id:     "810610",
			CityId: "8106",
			Name:   "Taniwel Timur",
		},
		{
			Id:     "810611",
			CityId: "8106",
			Name:   "Elpaputih",
		},
		{
			Id:     "810701",
			CityId: "8107",
			Name:   "Pulau-Pulau Aru",
		},
		{
			Id:     "810702",
			CityId: "8107",
			Name:   "Aru Selatan",
		},
		{
			Id:     "810703",
			CityId: "8107",
			Name:   "Aru Tengah",
		},
		{
			Id:     "810704",
			CityId: "8107",
			Name:   "Aru Utara",
		},
		{
			Id:     "810705",
			CityId: "8107",
			Name:   "Aru Utara Timur Batuley",
		},
		{
			Id:     "810706",
			CityId: "8107",
			Name:   "Sir-Sir",
		},
		{
			Id:     "810707",
			CityId: "8107",
			Name:   "Aru Tengah Timur",
		},
		{
			Id:     "810708",
			CityId: "8107",
			Name:   "Aru Tengah Selatan",
		},
		{
			Id:     "810709",
			CityId: "8107",
			Name:   "Aru Selatan Timur",
		},
		{
			Id:     "810710",
			CityId: "8107",
			Name:   "Aru Selatan Utara",
		},
		{
			Id:     "810801",
			CityId: "8108",
			Name:   "Moa Lakor",
		},
		{
			Id:     "810802",
			CityId: "8108",
			Name:   "Damer",
		},
		{
			Id:     "810803",
			CityId: "8108",
			Name:   "Mndona Hiera",
		},
		{
			Id:     "810804",
			CityId: "8108",
			Name:   "Pulau-Pulau Babar",
		},
		{
			Id:     "810805",
			CityId: "8108",
			Name:   "Pulau-pulau Babar Timur",
		},
		{
			Id:     "810806",
			CityId: "8108",
			Name:   "Wetar",
		},
		{
			Id:     "810807",
			CityId: "8108",
			Name:   "Pulau-pulau Terselatan",
		},
		{
			Id:     "810808",
			CityId: "8108",
			Name:   "Pulau Leti",
		},
		{
			Id:     "810809",
			CityId: "8108",
			Name:   "Pulau Masela",
		},
		{
			Id:     "810810",
			CityId: "8108",
			Name:   "Dawelor Dawera",
		},
		{
			Id:     "810811",
			CityId: "8108",
			Name:   "Pulau Wetang",
		},
		{
			Id:     "810812",
			CityId: "8108",
			Name:   "Pulau Lakor",
		},
		{
			Id:     "810813",
			CityId: "8108",
			Name:   "Wetar Utara",
		},
		{
			Id:     "810814",
			CityId: "8108",
			Name:   "Wetar Barat",
		},
		{
			Id:     "810815",
			CityId: "8108",
			Name:   "Wetar Timur",
		},
		{
			Id:     "810816",
			CityId: "8108",
			Name:   "Kepulauan Romang",
		},
		{
			Id:     "810817",
			CityId: "8108",
			Name:   "Kisar Utara",
		},
		{
			Id:     "810901",
			CityId: "8109",
			Name:   "Namrole",
		},
		{
			Id:     "810902",
			CityId: "8109",
			Name:   "Waesama",
		},
		{
			Id:     "810903",
			CityId: "8109",
			Name:   "Ambalau",
		},
		{
			Id:     "810904",
			CityId: "8109",
			Name:   "Kepala Madan",
		},
		{
			Id:     "810905",
			CityId: "8109",
			Name:   "Leksula",
		},
		{
			Id:     "810906",
			CityId: "8109",
			Name:   "Fena Fafan",
		},
		{
			Id:     "817101",
			CityId: "8171",
			Name:   "Nusaniwe",
		},
		{
			Id:     "817102",
			CityId: "8171",
			Name:   "Sirimau",
		},
		{
			Id:     "817103",
			CityId: "8171",
			Name:   "Baguala",
		},
		{
			Id:     "817104",
			CityId: "8171",
			Name:   "Teluk Ambon",
		},
		{
			Id:     "817105",
			CityId: "8171",
			Name:   "Leitimur Selatan",
		},
		{
			Id:     "817201",
			CityId: "8172",
			Name:   "Pulau Dullah Utara",
		},
		{
			Id:     "817202",
			CityId: "8172",
			Name:   "Pulau Dullah Selatan",
		},
		{
			Id:     "817203",
			CityId: "8172",
			Name:   "Tayando Tam",
		},
		{
			Id:     "817204",
			CityId: "8172",
			Name:   "Pulau-Pulau Kur",
		},
		{
			Id:     "817205",
			CityId: "8172",
			Name:   "Kur Selatan",
		},
		{
			Id:     "820101",
			CityId: "8201",
			Name:   "Jailolo",
		},
		{
			Id:     "820102",
			CityId: "8201",
			Name:   "Loloda",
		},
		{
			Id:     "820103",
			CityId: "8201",
			Name:   "Ibu",
		},
		{
			Id:     "820104",
			CityId: "8201",
			Name:   "Sahu",
		},
		{
			Id:     "820105",
			CityId: "8201",
			Name:   "Jailolo Selatan",
		},
		{
			Id:     "820107",
			CityId: "8201",
			Name:   "Ibu Utara",
		},
		{
			Id:     "820108",
			CityId: "8201",
			Name:   "Ibu Selatan",
		},
		{
			Id:     "820109",
			CityId: "8201",
			Name:   "Sahu Timur",
		},
		{
			Id:     "820110",
			CityId: "8201",
			Name:   "Loloda Tengah",
		},
		{
			Id:     "820201",
			CityId: "8202",
			Name:   "Weda",
		},
		{
			Id:     "820202",
			CityId: "8202",
			Name:   "Patani",
		},
		{
			Id:     "820203",
			CityId: "8202",
			Name:   "Pulau Gebe",
		},
		{
			Id:     "820204",
			CityId: "8202",
			Name:   "Weda Utara",
		},
		{
			Id:     "820205",
			CityId: "8202",
			Name:   "Weda Selatan",
		},
		{
			Id:     "820206",
			CityId: "8202",
			Name:   "Patani Utara",
		},
		{
			Id:     "820207",
			CityId: "8202",
			Name:   "Weda Tengah",
		},
		{
			Id:     "820208",
			CityId: "8202",
			Name:   "Patani Barat",
		},
		{
			Id:     "820209",
			CityId: "8202",
			Name:   "Weda Timur",
		},
		{
			Id:     "820210",
			CityId: "8202",
			Name:   "Patani Timur",
		},
		{
			Id:     "820304",
			CityId: "8203",
			Name:   "Galela",
		},
		{
			Id:     "820305",
			CityId: "8203",
			Name:   "Tobelo",
		},
		{
			Id:     "820306",
			CityId: "8203",
			Name:   "Tobelo Selatan",
		},
		{
			Id:     "820307",
			CityId: "8203",
			Name:   "Kao",
		},
		{
			Id:     "820308",
			CityId: "8203",
			Name:   "Malifut",
		},
		{
			Id:     "820309",
			CityId: "8203",
			Name:   "Loloda Utara",
		},
		{
			Id:     "820310",
			CityId: "8203",
			Name:   "Tobelo Utara",
		},
		{
			Id:     "820311",
			CityId: "8203",
			Name:   "Tobelo Tengah",
		},
		{
			Id:     "820312",
			CityId: "8203",
			Name:   "Tobelo Timur",
		},
		{
			Id:     "820313",
			CityId: "8203",
			Name:   "Tobelo Barat",
		},
		{
			Id:     "820314",
			CityId: "8203",
			Name:   "Galela Barat",
		},
		{
			Id:     "820315",
			CityId: "8203",
			Name:   "Galela Utara",
		},
		{
			Id:     "820316",
			CityId: "8203",
			Name:   "Galela Selatan",
		},
		{
			Id:     "820319",
			CityId: "8203",
			Name:   "Loloda Kepulauan",
		},
		{
			Id:     "820320",
			CityId: "8203",
			Name:   "Kao Utara",
		},
		{
			Id:     "820321",
			CityId: "8203",
			Name:   "Kao Barat",
		},
		{
			Id:     "820322",
			CityId: "8203",
			Name:   "Kao Teluk",
		},
		{
			Id:     "820401",
			CityId: "8204",
			Name:   "Pulau Makian",
		},
		{
			Id:     "820402",
			CityId: "8204",
			Name:   "Kayoa",
		},
		{
			Id:     "820403",
			CityId: "8204",
			Name:   "Gane Timur",
		},
		{
			Id:     "820404",
			CityId: "8204",
			Name:   "Gane Barat",
		},
		{
			Id:     "820405",
			CityId: "8204",
			Name:   "Obi Selatan",
		},
		{
			Id:     "820406",
			CityId: "8204",
			Name:   "Obi",
		},
		{
			Id:     "820407",
			CityId: "8204",
			Name:   "Bacan Timur",
		},
		{
			Id:     "820408",
			CityId: "8204",
			Name:   "Bacan",
		},
		{
			Id:     "820409",
			CityId: "8204",
			Name:   "Bacan Barat",
		},
		{
			Id:     "820410",
			CityId: "8204",
			Name:   "Makian Barat",
		},
		{
			Id:     "820411",
			CityId: "8204",
			Name:   "Kayoa Barat",
		},
		{
			Id:     "820412",
			CityId: "8204",
			Name:   "Kayoa Selatan",
		},
		{
			Id:     "820413",
			CityId: "8204",
			Name:   "Kayoa Utara",
		},
		{
			Id:     "820414",
			CityId: "8204",
			Name:   "Bacan Barat Utara",
		},
		{
			Id:     "820415",
			CityId: "8204",
			Name:   "Kasiruta Barat",
		},
		{
			Id:     "820416",
			CityId: "8204",
			Name:   "Kasiruta Timur",
		},
		{
			Id:     "820417",
			CityId: "8204",
			Name:   "Bacan Selatan",
		},
		{
			Id:     "820418",
			CityId: "8204",
			Name:   "Kepulauan Botanglomang",
		},
		{
			Id:     "820419",
			CityId: "8204",
			Name:   "Mandioli Selatan",
		},
		{
			Id:     "820420",
			CityId: "8204",
			Name:   "Mandioli Utara",
		},
		{
			Id:     "820421",
			CityId: "8204",
			Name:   "Bacan Timur Selatan",
		},
		{
			Id:     "820422",
			CityId: "8204",
			Name:   "Bacan Timur Tengah",
		},
		{
			Id:     "820423",
			CityId: "8204",
			Name:   "Gane Barat Selatan",
		},
		{
			Id:     "820424",
			CityId: "8204",
			Name:   "Gane Barat Utara",
		},
		{
			Id:     "820425",
			CityId: "8204",
			Name:   "Kepulauan Joronga",
		},
		{
			Id:     "820426",
			CityId: "8204",
			Name:   "Gane Timur Selatan",
		},
		{
			Id:     "820427",
			CityId: "8204",
			Name:   "Gane Timur Tengah",
		},
		{
			Id:     "820428",
			CityId: "8204",
			Name:   "Obi Barat",
		},
		{
			Id:     "820429",
			CityId: "8204",
			Name:   "Obi Timur",
		},
		{
			Id:     "820430",
			CityId: "8204",
			Name:   "Obi Utara",
		},
		{
			Id:     "820501",
			CityId: "8205",
			Name:   "Mangoli Timur",
		},
		{
			Id:     "820502",
			CityId: "8205",
			Name:   "Sanana",
		},
		{
			Id:     "820503",
			CityId: "8205",
			Name:   "Sulabesi Barat",
		},
		{
			Id:     "820506",
			CityId: "8205",
			Name:   "Mangoli Barat",
		},
		{
			Id:     "820507",
			CityId: "8205",
			Name:   "Sulabesi Tengah",
		},
		{
			Id:     "820508",
			CityId: "8205",
			Name:   "Sulabesi Timur",
		},
		{
			Id:     "820509",
			CityId: "8205",
			Name:   "Sulabesi Selatan",
		},
		{
			Id:     "820510",
			CityId: "8205",
			Name:   "Mangoli Utara Timur",
		},
		{
			Id:     "820511",
			CityId: "8205",
			Name:   "Mangoli Tengah",
		},
		{
			Id:     "820512",
			CityId: "8205",
			Name:   "Mangoli Selatan",
		},
		{
			Id:     "820513",
			CityId: "8205",
			Name:   "Mangoli Utara",
		},
		{
			Id:     "820518",
			CityId: "8205",
			Name:   "Sanana Utara",
		},
		{
			Id:     "820601",
			CityId: "8206",
			Name:   "Wasile",
		},
		{
			Id:     "820602",
			CityId: "8206",
			Name:   "Maba",
		},
		{
			Id:     "820603",
			CityId: "8206",
			Name:   "Maba Selatan",
		},
		{
			Id:     "820604",
			CityId: "8206",
			Name:   "Wasile Selatan",
		},
		{
			Id:     "820605",
			CityId: "8206",
			Name:   "Wasile Tengah",
		},
		{
			Id:     "820606",
			CityId: "8206",
			Name:   "Wasile Utara",
		},
		{
			Id:     "820607",
			CityId: "8206",
			Name:   "Wasile Timur",
		},
		{
			Id:     "820608",
			CityId: "8206",
			Name:   "Maba Tengah",
		},
		{
			Id:     "820609",
			CityId: "8206",
			Name:   "Maba Utara",
		},
		{
			Id:     "820610",
			CityId: "8206",
			Name:   "Kota Maba",
		},
		{
			Id:     "820701",
			CityId: "8207",
			Name:   "Morotai Selatan",
		},
		{
			Id:     "820702",
			CityId: "8207",
			Name:   "Morotai Selatan Barat",
		},
		{
			Id:     "820703",
			CityId: "8207",
			Name:   "Morotai Jaya",
		},
		{
			Id:     "820704",
			CityId: "8207",
			Name:   "Morotai Utara",
		},
		{
			Id:     "820705",
			CityId: "8207",
			Name:   "Morotai Timur",
		},
		{
			Id:     "820706",
			CityId: "8207",
			Name:   "Pulau Rao",
		},
		{
			Id:     "820801",
			CityId: "8208",
			Name:   "Taliabu Barat",
		},
		{
			Id:     "820802",
			CityId: "8208",
			Name:   "Taliabu Barat Laut",
		},
		{
			Id:     "820803",
			CityId: "8208",
			Name:   "Lede",
		},
		{
			Id:     "820804",
			CityId: "8208",
			Name:   "Taliabu Utara",
		},
		{
			Id:     "820805",
			CityId: "8208",
			Name:   "Taliabu Timur",
		},
		{
			Id:     "820806",
			CityId: "8208",
			Name:   "Taliabu Timur Selatan",
		},
		{
			Id:     "820807",
			CityId: "8208",
			Name:   "Taliabu Selatan",
		},
		{
			Id:     "820808",
			CityId: "8208",
			Name:   "Tabona",
		},
		{
			Id:     "827101",
			CityId: "8271",
			Name:   "Pulau Ternate",
		},
		{
			Id:     "827102",
			CityId: "8271",
			Name:   "Kota Ternate Selatan",
		},
		{
			Id:     "827103",
			CityId: "8271",
			Name:   "Kota Ternate Utara",
		},
		{
			Id:     "827104",
			CityId: "8271",
			Name:   "Moti",
		},
		{
			Id:     "827105",
			CityId: "8271",
			Name:   "Pulau Batang Dua",
		},
		{
			Id:     "827106",
			CityId: "8271",
			Name:   "Kota Ternate Tengah",
		},
		{
			Id:     "827107",
			CityId: "8271",
			Name:   "Pulau Hiri",
		},
		{
			Id:     "827108",
			CityId: "8271",
			Name:   "Ternate Barat",
		},
		{
			Id:     "827201",
			CityId: "8272",
			Name:   "Tidore",
		},
		{
			Id:     "827202",
			CityId: "8272",
			Name:   "Oba Utara",
		},
		{
			Id:     "827203",
			CityId: "8272",
			Name:   "Oba",
		},
		{
			Id:     "827204",
			CityId: "8272",
			Name:   "Tidore Selatan",
		},
		{
			Id:     "827205",
			CityId: "8272",
			Name:   "Tidore Utara",
		},
		{
			Id:     "827206",
			CityId: "8272",
			Name:   "Oba Tengah",
		},
		{
			Id:     "827207",
			CityId: "8272",
			Name:   "Oba Selatan",
		},
		{
			Id:     "827208",
			CityId: "8272",
			Name:   "Tidore Timur",
		},
		{
			Id:     "910301",
			CityId: "9103",
			Name:   "Sentani",
		},
		{
			Id:     "910302",
			CityId: "9103",
			Name:   "Sentani Timur",
		},
		{
			Id:     "910303",
			CityId: "9103",
			Name:   "Depapre",
		},
		{
			Id:     "910304",
			CityId: "9103",
			Name:   "Sentani Barat",
		},
		{
			Id:     "910305",
			CityId: "9103",
			Name:   "Kemtuk",
		},
		{
			Id:     "910306",
			CityId: "9103",
			Name:   "Kemtuk Gresi",
		},
		{
			Id:     "910307",
			CityId: "9103",
			Name:   "Nimboran",
		},
		{
			Id:     "910308",
			CityId: "9103",
			Name:   "Nimbokrang",
		},
		{
			Id:     "910309",
			CityId: "9103",
			Name:   "Unurum Guay",
		},
		{
			Id:     "910310",
			CityId: "9103",
			Name:   "Demta",
		},
		{
			Id:     "910311",
			CityId: "9103",
			Name:   "Kaureh",
		},
		{
			Id:     "910312",
			CityId: "9103",
			Name:   "Ebungfao",
		},
		{
			Id:     "910313",
			CityId: "9103",
			Name:   "Waibu",
		},
		{
			Id:     "910314",
			CityId: "9103",
			Name:   "Nambluong",
		},
		{
			Id:     "910315",
			CityId: "9103",
			Name:   "Yapsi",
		},
		{
			Id:     "910316",
			CityId: "9103",
			Name:   "Airu",
		},
		{
			Id:     "910317",
			CityId: "9103",
			Name:   "Raveni Rara",
		},
		{
			Id:     "910318",
			CityId: "9103",
			Name:   "Gresi Selatan",
		},
		{
			Id:     "910319",
			CityId: "9103",
			Name:   "Yokari",
		},
		{
			Id:     "910501",
			CityId: "9105",
			Name:   "Yapen Selatan",
		},
		{
			Id:     "910502",
			CityId: "9105",
			Name:   "Yapen Barat",
		},
		{
			Id:     "910503",
			CityId: "9105",
			Name:   "Yapen Timur",
		},
		{
			Id:     "910504",
			CityId: "9105",
			Name:   "Angkaisera",
		},
		{
			Id:     "910505",
			CityId: "9105",
			Name:   "Poom",
		},
		{
			Id:     "910506",
			CityId: "9105",
			Name:   "Kosiwo",
		},
		{
			Id:     "910507",
			CityId: "9105",
			Name:   "Yapen Utara",
		},
		{
			Id:     "910508",
			CityId: "9105",
			Name:   "Raimbawi",
		},
		{
			Id:     "910509",
			CityId: "9105",
			Name:   "Teluk Ampimoi",
		},
		{
			Id:     "910510",
			CityId: "9105",
			Name:   "Kepulauan Ambai",
		},
		{
			Id:     "910511",
			CityId: "9105",
			Name:   "Wonawa",
		},
		{
			Id:     "910512",
			CityId: "9105",
			Name:   "Windesi",
		},
		{
			Id:     "910513",
			CityId: "9105",
			Name:   "Pulau Kurudu",
		},
		{
			Id:     "910514",
			CityId: "9105",
			Name:   "Pulau Yerui",
		},
		{
			Id:     "910515",
			CityId: "9105",
			Name:   "Anotaurei",
		},
		{
			Id:     "910516",
			CityId: "9105",
			Name:   "Yawakukat",
		},
		{
			Id:     "910517",
			CityId: "9105",
			Name:   "Nusawani",
		},
		{
			Id:     "910601",
			CityId: "9106",
			Name:   "Biak Kota",
		},
		{
			Id:     "910602",
			CityId: "9106",
			Name:   "Biak Utara",
		},
		{
			Id:     "910603",
			CityId: "9106",
			Name:   "Biak Timur",
		},
		{
			Id:     "910604",
			CityId: "9106",
			Name:   "Numfor Barat",
		},
		{
			Id:     "910605",
			CityId: "9106",
			Name:   "Numfor Timur",
		},
		{
			Id:     "910608",
			CityId: "9106",
			Name:   "Biak Barat",
		},
		{
			Id:     "910609",
			CityId: "9106",
			Name:   "Warsa",
		},
		{
			Id:     "910610",
			CityId: "9106",
			Name:   "Padaido",
		},
		{
			Id:     "910611",
			CityId: "9106",
			Name:   "Yendidori",
		},
		{
			Id:     "910612",
			CityId: "9106",
			Name:   "Samofa",
		},
		{
			Id:     "910613",
			CityId: "9106",
			Name:   "Yawosi",
		},
		{
			Id:     "910614",
			CityId: "9106",
			Name:   "Andey",
		},
		{
			Id:     "910615",
			CityId: "9106",
			Name:   "Swandiwe",
		},
		{
			Id:     "910616",
			CityId: "9106",
			Name:   "Bruyadori",
		},
		{
			Id:     "910617",
			CityId: "9106",
			Name:   "Orkeri",
		},
		{
			Id:     "910618",
			CityId: "9106",
			Name:   "Poiru",
		},
		{
			Id:     "910619",
			CityId: "9106",
			Name:   "Aimando Padaido",
		},
		{
			Id:     "910620",
			CityId: "9106",
			Name:   "Oridek",
		},
		{
			Id:     "910621",
			CityId: "9106",
			Name:   "Bondifuar",
		},
		{
			Id:     "911001",
			CityId: "9110",
			Name:   "Sarmi",
		},
		{
			Id:     "911002",
			CityId: "9110",
			Name:   "Tor Atas",
		},
		{
			Id:     "911003",
			CityId: "9110",
			Name:   "Pantai Barat",
		},
		{
			Id:     "911004",
			CityId: "9110",
			Name:   "Pantai Timur",
		},
		{
			Id:     "911005",
			CityId: "9110",
			Name:   "Bonggo",
		},
		{
			Id:     "911009",
			CityId: "9110",
			Name:   "Apawer Hulu",
		},
		{
			Id:     "911012",
			CityId: "9110",
			Name:   "Sarmi Selatan",
		},
		{
			Id:     "911013",
			CityId: "9110",
			Name:   "Sarmi Timur",
		},
		{
			Id:     "911014",
			CityId: "9110",
			Name:   "Pantai Timur Bagian Barat",
		},
		{
			Id:     "911015",
			CityId: "9110",
			Name:   "Bonggo Timur",
		},
		{
			Id:     "911101",
			CityId: "9111",
			Name:   "Waris",
		},
		{
			Id:     "911102",
			CityId: "9111",
			Name:   "Arso",
		},
		{
			Id:     "911103",
			CityId: "9111",
			Name:   "Senggi",
		},
		{
			Id:     "911104",
			CityId: "9111",
			Name:   "Web",
		},
		{
			Id:     "911105",
			CityId: "9111",
			Name:   "Skanto",
		},
		{
			Id:     "911106",
			CityId: "9111",
			Name:   "Arso Timur",
		},
		{
			Id:     "911107",
			CityId: "9111",
			Name:   "Towe",
		},
		{
			Id:     "911108",
			CityId: "9111",
			Name:   "Arso Barat",
		},
		{
			Id:     "911109",
			CityId: "9111",
			Name:   "Mannem",
		},
		{
			Id:     "911110",
			CityId: "9111",
			Name:   "Yaffi",
		},
		{
			Id:     "911111",
			CityId: "9111",
			Name:   "Kaisenar",
		},
		{
			Id:     "911501",
			CityId: "9115",
			Name:   "Waropen Bawah",
		},
		{
			Id:     "911503",
			CityId: "9115",
			Name:   "Masirei",
		},
		{
			Id:     "911507",
			CityId: "9115",
			Name:   "Risei Sayati",
		},
		{
			Id:     "911508",
			CityId: "9115",
			Name:   "Urei Faisei",
		},
		{
			Id:     "911509",
			CityId: "9115",
			Name:   "Inggerus",
		},
		{
			Id:     "911510",
			CityId: "9115",
			Name:   "Kirihi",
		},
		{
			Id:     "911511",
			CityId: "9115",
			Name:   "Oudate",
		},
		{
			Id:     "911512",
			CityId: "9115",
			Name:   "Wapoga",
		},
		{
			Id:     "911513",
			CityId: "9115",
			Name:   "Demba",
		},
		{
			Id:     "911514",
			CityId: "9115",
			Name:   "Wonti",
		},
		{
			Id:     "911515",
			CityId: "9115",
			Name:   "Soyoi Mambai",
		},
		{
			Id:     "911901",
			CityId: "9119",
			Name:   "Supiori Selatan",
		},
		{
			Id:     "911902",
			CityId: "9119",
			Name:   "Supiori Utara",
		},
		{
			Id:     "911903",
			CityId: "9119",
			Name:   "Supiori Timur",
		},
		{
			Id:     "911904",
			CityId: "9119",
			Name:   "Kepulauan Aruri",
		},
		{
			Id:     "911905",
			CityId: "9119",
			Name:   "Supiori Barat",
		},
		{
			Id:     "912001",
			CityId: "9120",
			Name:   "Mamberamo Tengah",
		},
		{
			Id:     "912002",
			CityId: "9120",
			Name:   "Mamberamo Hulu",
		},
		{
			Id:     "912003",
			CityId: "9120",
			Name:   "Rufaer",
		},
		{
			Id:     "912004",
			CityId: "9120",
			Name:   "Mamberamo Tengah Timur",
		},
		{
			Id:     "912005",
			CityId: "9120",
			Name:   "Mamberamo Hilir",
		},
		{
			Id:     "912006",
			CityId: "9120",
			Name:   "Waropen Atas",
		},
		{
			Id:     "912007",
			CityId: "9120",
			Name:   "Benuki",
		},
		{
			Id:     "912008",
			CityId: "9120",
			Name:   "Sawai",
		},
		{
			Id:     "917101",
			CityId: "9171",
			Name:   "Jayapura Utara",
		},
		{
			Id:     "917102",
			CityId: "9171",
			Name:   "Jayapura Selatan",
		},
		{
			Id:     "950840",
			CityId: "9171",
			Name:   "Distrik Ebungfa",
		},
		{
			Id:     "950841",
			CityId: "9171",
			Name:   "Distrik Waibu",
		},
		{
			Id:     "950842",
			CityId: "9501",
			Name:   "Palebaga",
		},
		{
			Id:     "950843",
			CityId: "9202",
			Name:   "Distrik Manokwari",
		},
		{
			Id:     "917103",
			CityId: "9171",
			Name:   "Abepura",
		},
		{
			Id:     "917104",
			CityId: "9171",
			Name:   "Muara Tami",
		},
		{
			Id:     "917105",
			CityId: "9171",
			Name:   "Heram",
		},
		{
			Id:     "920101",
			CityId: "9201",
			Name:   "Makbon",
		},
		{
			Id:     "920104",
			CityId: "9201",
			Name:   "Beraur",
		},
		{
			Id:     "920105",
			CityId: "9201",
			Name:   "Salawati",
		},
		{
			Id:     "920106",
			CityId: "9201",
			Name:   "Seget",
		},
		{
			Id:     "920107",
			CityId: "9206",
			Name:   "Aimas",
		},
		{
			Id:     "920108",
			CityId: "9201",
			Name:   "Klamono",
		},
		{
			Id:     "920110",
			CityId: "9201",
			Name:   "Sayosa",
		},
		{
			Id:     "920112",
			CityId: "9201",
			Name:   "Segun",
		},
		{
			Id:     "920113",
			CityId: "9201",
			Name:   "Mayamuk",
		},
		{
			Id:     "920114",
			CityId: "9201",
			Name:   "Salawati Selatan",
		},
		{
			Id:     "920117",
			CityId: "9201",
			Name:   "Klabot",
		},
		{
			Id:     "920118",
			CityId: "9201",
			Name:   "Klawak",
		},
		{
			Id:     "920120",
			CityId: "9201",
			Name:   "Maudus",
		},
		{
			Id:     "920139",
			CityId: "9201",
			Name:   "Mariat",
		},
		{
			Id:     "920140",
			CityId: "9201",
			Name:   "Klayili",
		},
		{
			Id:     "920141",
			CityId: "9201",
			Name:   "Klaso",
		},
		{
			Id:     "920142",
			CityId: "9201",
			Name:   "Moisegen",
		},
		{
			Id:     "920143",
			CityId: "9201",
			Name:   "Sorong",
		},
		{
			Id:     "920144",
			CityId: "9201",
			Name:   "Bagun",
		},
		{
			Id:     "920145",
			CityId: "9201",
			Name:   "Wemak",
		},
		{
			Id:     "920146",
			CityId: "9201",
			Name:   "Sunook",
		},
		{
			Id:     "920147",
			CityId: "9201",
			Name:   "Buk",
		},
		{
			Id:     "920148",
			CityId: "9201",
			Name:   "Saengkeduk",
		},
		{
			Id:     "920149",
			CityId: "9201",
			Name:   "Malabotom",
		},
		{
			Id:     "920150",
			CityId: "9201",
			Name:   "Konhir",
		},
		{
			Id:     "920151",
			CityId: "9201",
			Name:   "Klasafet",
		},
		{
			Id:     "920152",
			CityId: "9201",
			Name:   "Hobard",
		},
		{
			Id:     "920153",
			CityId: "9201",
			Name:   "Salawati Tengah",
		},
		{
			Id:     "920154",
			CityId: "9201",
			Name:   "Botain",
		},
		{
			Id:     "920155",
			CityId: "9201",
			Name:   "Sayosa Timur",
		},
		{
			Id:     "920203",
			CityId: "9202",
			Name:   "Warmare",
		},
		{
			Id:     "920204",
			CityId: "9202",
			Name:   "Prafi",
		},
		{
			Id:     "920205",
			CityId: "9202",
			Name:   "Masni",
		},
		{
			Id:     "920212",
			CityId: "9202",
			Name:   "Manokwari Barat",
		},
		{
			Id:     "920213",
			CityId: "9202",
			Name:   "Manokwari Timur",
		},
		{
			Id:     "920214",
			CityId: "9202",
			Name:   "Manokwari Utara",
		},
		{
			Id:     "920215",
			CityId: "9202",
			Name:   "Manokwari Selatan",
		},
		{
			Id:     "920217",
			CityId: "9202",
			Name:   "Tanah Rubuh",
		},
		{
			Id:     "920221",
			CityId: "9202",
			Name:   "Sidey",
		},
		{
			Id:     "920301",
			CityId: "9203",
			Name:   "Fak-Fak",
		},
		{
			Id:     "920302",
			CityId: "9203",
			Name:   "Fak-Fak Barat",
		},
		{
			Id:     "920303",
			CityId: "9203",
			Name:   "Fak-Fak Timur",
		},
		{
			Id:     "920304",
			CityId: "9203",
			Name:   "Kokas",
		},
		{
			Id:     "920305",
			CityId: "9203",
			Name:   "Fak-Fak Tengah",
		},
		{
			Id:     "920306",
			CityId: "9203",
			Name:   "Karas",
		},
		{
			Id:     "920307",
			CityId: "9203",
			Name:   "Bomberay",
		},
		{
			Id:     "920308",
			CityId: "9203",
			Name:   "Kramongmongga",
		},
		{
			Id:     "920309",
			CityId: "9203",
			Name:   "Teluk Patipi",
		},
		{
			Id:     "920310",
			CityId: "9203",
			Name:   "Pariwari",
		},
		{
			Id:     "920311",
			CityId: "9203",
			Name:   "Wartutin",
		},
		{
			Id:     "920312",
			CityId: "9203",
			Name:   "Fakfak Timur Tengah",
		},
		{
			Id:     "920313",
			CityId: "9203",
			Name:   "Arguni",
		},
		{
			Id:     "920314",
			CityId: "9203",
			Name:   "Mbahamdandara",
		},
		{
			Id:     "920315",
			CityId: "9203",
			Name:   "Kayauni",
		},
		{
			Id:     "920316",
			CityId: "9203",
			Name:   "Furwagi",
		},
		{
			Id:     "920317",
			CityId: "9203",
			Name:   "Tomage",
		},
		{
			Id:     "920401",
			CityId: "9204",
			Name:   "Teminabuan",
		},
		{
			Id:     "920404",
			CityId: "9204",
			Name:   "Inanwatan",
		},
		{
			Id:     "920406",
			CityId: "9204",
			Name:   "Sawiat",
		},
		{
			Id:     "920409",
			CityId: "9204",
			Name:   "Kokoda",
		},
		{
			Id:     "920410",
			CityId: "9204",
			Name:   "Moswaren",
		},
		{
			Id:     "920411",
			CityId: "9204",
			Name:   "Seremuk",
		},
		{
			Id:     "920412",
			CityId: "9204",
			Name:   "Wayer",
		},
		{
			Id:     "920414",
			CityId: "9204",
			Name:   "Kais",
		},
		{
			Id:     "920415",
			CityId: "9204",
			Name:   "Konda",
		},
		{
			Id:     "920420",
			CityId: "9204",
			Name:   "Matemani",
		},
		{
			Id:     "920421",
			CityId: "9204",
			Name:   "Kokoda Utara",
		},
		{
			Id:     "920422",
			CityId: "9204",
			Name:   "Saifi",
		},
		{
			Id:     "920424",
			CityId: "9204",
			Name:   "Fokour",
		},
		{
			Id:     "920425",
			CityId: "9204",
			Name:   "Salkma",
		},
		{
			Id:     "920426",
			CityId: "9204",
			Name:   "Kais Darat",
		},
		{
			Id:     "920501",
			CityId: "9205",
			Name:   "Misool (Misool Utara)",
		},
		{
			Id:     "920502",
			CityId: "9205",
			Name:   "Waigeo Utara",
		},
		{
			Id:     "920503",
			CityId: "9205",
			Name:   "Waigeo Selatan",
		},
		{
			Id:     "920504",
			CityId: "9205",
			Name:   "Salawati Utara",
		},
		{
			Id:     "920505",
			CityId: "9205",
			Name:   "Kepulauan Ayau",
		},
		{
			Id:     "920506",
			CityId: "9205",
			Name:   "Misool Timur",
		},
		{
			Id:     "920507",
			CityId: "9205",
			Name:   "Waigeo Barat",
		},
		{
			Id:     "920508",
			CityId: "9205",
			Name:   "Waigeo Timur",
		},
		{
			Id:     "920509",
			CityId: "9205",
			Name:   "Teluk Mayalibit",
		},
		{
			Id:     "920510",
			CityId: "9205",
			Name:   "Kofiau",
		},
		{
			Id:     "920511",
			CityId: "9206",
			Name:   "Meos Mansar",
		},
		{
			Id:     "920513",
			CityId: "9205",
			Name:   "Misool Selatan",
		},
		{
			Id:     "920514",
			CityId: "9205",
			Name:   "Warwarbomi",
		},
		{
			Id:     "920515",
			CityId: "9205",
			Name:   "Waigeo Barat Kepulauan",
		},
		{
			Id:     "920516",
			CityId: "9205",
			Name:   "Misool Barat",
		},
		{
			Id:     "920517",
			CityId: "9205",
			Name:   "Kepulauan Sembilan",
		},
		{
			Id:     "920518",
			CityId: "9205",
			Name:   "Kota Waisai",
		},
		{
			Id:     "920519",
			CityId: "9205",
			Name:   "Tiplol Mayalibit",
		},
		{
			Id:     "920520",
			CityId: "9205",
			Name:   "Batanta Utara",
		},
		{
			Id:     "920521",
			CityId: "9205",
			Name:   "Salawati Barat",
		},
		{
			Id:     "920522",
			CityId: "9205",
			Name:   "Salawati Tengah",
		},
		{
			Id:     "920523",
			CityId: "9205",
			Name:   "Supnin",
		},
		{
			Id:     "920524",
			CityId: "9205",
			Name:   "Ayau",
		},
		{
			Id:     "920525",
			CityId: "9205",
			Name:   "Batanta Selatan",
		},
		{
			Id:     "920601",
			CityId: "9206",
			Name:   "Bintuni",
		},
		{
			Id:     "920602",
			CityId: "9206",
			Name:   "Merdey",
		},
		{
			Id:     "920603",
			CityId: "9206",
			Name:   "Babo",
		},
		{
			Id:     "920604",
			CityId: "9206",
			Name:   "Aranday",
		},
		{
			Id:     "920605",
			CityId: "9206",
			Name:   "Moskona Selatan",
		},
		{
			Id:     "920606",
			CityId: "9206",
			Name:   "Moskona Utara",
		},
		{
			Id:     "920607",
			CityId: "9206",
			Name:   "Wamesa",
		},
		{
			Id:     "920608",
			CityId: "9206",
			Name:   "Fafurwar",
		},
		{
			Id:     "920609",
			CityId: "9206",
			Name:   "Tembuni",
		},
		{
			Id:     "920610",
			CityId: "9206",
			Name:   "Kuri",
		},
		{
			Id:     "920611",
			CityId: "9206",
			Name:   "Manimeri",
		},
		{
			Id:     "920612",
			CityId: "9206",
			Name:   "Tuhiba",
		},
		{
			Id:     "920613",
			CityId: "9206",
			Name:   "Dataran Beimes",
		},
		{
			Id:     "920614",
			CityId: "9206",
			Name:   "Sumuri",
		},
		{
			Id:     "920615",
			CityId: "9206",
			Name:   "Kaitaro",
		},
		{
			Id:     "920616",
			CityId: "9206",
			Name:   "Aroba",
		},
		{
			Id:     "920617",
			CityId: "9206",
			Name:   "Masyeta",
		},
		{
			Id:     "920618",
			CityId: "9206",
			Name:   "Biscoop",
		},
		{
			Id:     "920619",
			CityId: "9206",
			Name:   "Tomu",
		},
		{
			Id:     "920620",
			CityId: "9206",
			Name:   "Kamundan",
		},
		{
			Id:     "920621",
			CityId: "9206",
			Name:   "Weriagar",
		},
		{
			Id:     "920622",
			CityId: "9206",
			Name:   "Moskona Barat",
		},
		{
			Id:     "920623",
			CityId: "9206",
			Name:   "Meyado",
		},
		{
			Id:     "920624",
			CityId: "9206",
			Name:   "Moskona Timur",
		},
		{
			Id:     "920701",
			CityId: "9207",
			Name:   "Wasior",
		},
		{
			Id:     "920702",
			CityId: "9207",
			Name:   "Windesi",
		},
		{
			Id:     "920703",
			CityId: "9207",
			Name:   "Teluk Duairi",
		},
		{
			Id:     "920704",
			CityId: "9207",
			Name:   "Wondiboy",
		},
		{
			Id:     "920705",
			CityId: "9207",
			Name:   "Wamesa",
		},
		{
			Id:     "920706",
			CityId: "9207",
			Name:   "Rumberpon",
		},
		{
			Id:     "920707",
			CityId: "9207",
			Name:   "Naikere",
		},
		{
			Id:     "920708",
			CityId: "9207",
			Name:   "Rasiei",
		},
		{
			Id:     "920709",
			CityId: "9207",
			Name:   "Kuri Wamesa",
		},
		{
			Id:     "920710",
			CityId: "9207",
			Name:   "Roon",
		},
		{
			Id:     "920711",
			CityId: "9207",
			Name:   "Roswar",
		},
		{
			Id:     "920712",
			CityId: "9207",
			Name:   "Nikiwar",
		},
		{
			Id:     "920713",
			CityId: "9207",
			Name:   "Soug Jaya",
		},
		{
			Id:     "920801",
			CityId: "9208",
			Name:   "Kaimana",
		},
		{
			Id:     "920802",
			CityId: "9208",
			Name:   "Buruway",
		},
		{
			Id:     "920803",
			CityId: "9208",
			Name:   "Teluk Arguni Atas",
		},
		{
			Id:     "920804",
			CityId: "9208",
			Name:   "Teluk Etna",
		},
		{
			Id:     "920805",
			CityId: "9208",
			Name:   "Kambrau",
		},
		{
			Id:     "920806",
			CityId: "9208",
			Name:   "Teluk Arguni Bawah",
		},
		{
			Id:     "920807",
			CityId: "9208",
			Name:   "Yamor",
		},
		{
			Id:     "920901",
			CityId: "9209",
			Name:   "Fef",
		},
		{
			Id:     "920902",
			CityId: "9209",
			Name:   "Miyah",
		},
		{
			Id:     "920903",
			CityId: "9209",
			Name:   "Yembun",
		},
		{
			Id:     "920904",
			CityId: "9209",
			Name:   "Kwoor",
		},
		{
			Id:     "920905",
			CityId: "9209",
			Name:   "Sausapor",
		},
		{
			Id:     "920906",
			CityId: "9209",
			Name:   "Abun",
		},
		{
			Id:     "920907",
			CityId: "9209",
			Name:   "Syujak",
		},
		{
			Id:     "920908",
			CityId: "9209",
			Name:   "Moraid",
		},
		{
			Id:     "920909",
			CityId: "9209",
			Name:   "Kebar",
		},
		{
			Id:     "920910",
			CityId: "9209",
			Name:   "Amberbaken",
		},
		{
			Id:     "920911",
			CityId: "9209",
			Name:   "Senopi",
		},
		{
			Id:     "920912",
			CityId: "9209",
			Name:   "Mubrani",
		},
		{
			Id:     "920913",
			CityId: "9209",
			Name:   "Bikar",
		},
		{
			Id:     "920914",
			CityId: "9209",
			Name:   "Bamusbama",
		},
		{
			Id:     "920915",
			CityId: "9209",
			Name:   "Ases",
		},
		{
			Id:     "920916",
			CityId: "9209",
			Name:   "Miyah Selatan",
		},
		{
			Id:     "920917",
			CityId: "9209",
			Name:   "Ireres",
		},
		{
			Id:     "920918",
			CityId: "9209",
			Name:   "Tobouw",
		},
		{
			Id:     "920919",
			CityId: "9209",
			Name:   "Wilhem Roumbouts",
		},
		{
			Id:     "920920",
			CityId: "9209",
			Name:   "Tinggouw",
		},
		{
			Id:     "920921",
			CityId: "9209",
			Name:   "Kwesefo",
		},
		{
			Id:     "920922",
			CityId: "9209",
			Name:   "Mawabuan",
		},
		{
			Id:     "920923",
			CityId: "9209",
			Name:   "Kebar Timur",
		},
		{
			Id:     "920924",
			CityId: "9209",
			Name:   "Kebar Selatan",
		},
		{
			Id:     "920925",
			CityId: "9209",
			Name:   "Manekar",
		},
		{
			Id:     "920926",
			CityId: "9209",
			Name:   "Mpur",
		},
		{
			Id:     "920927",
			CityId: "9209",
			Name:   "Amberbaken Barat",
		},
		{
			Id:     "920928",
			CityId: "9209",
			Name:   "Kasi",
		},
		{
			Id:     "920929",
			CityId: "9209",
			Name:   "Selemkai",
		},
		{
			Id:     "921001",
			CityId: "9210",
			Name:   "Aifat",
		},
		{
			Id:     "921002",
			CityId: "9210",
			Name:   "Aifat Utara",
		},
		{
			Id:     "921003",
			CityId: "9210",
			Name:   "Aifat Timur",
		},
		{
			Id:     "921004",
			CityId: "9210",
			Name:   "Aifat Selatan",
		},
		{
			Id:     "921005",
			CityId: "9210",
			Name:   "Aitinyo Barat",
		},
		{
			Id:     "921006",
			CityId: "9210",
			Name:   "Aitinyo",
		},
		{
			Id:     "921007",
			CityId: "9210",
			Name:   "Aitinyo Utara",
		},
		{
			Id:     "921008",
			CityId: "9210",
			Name:   "Ayamaru",
		},
		{
			Id:     "921009",
			CityId: "9210",
			Name:   "Ayamaru Utara",
		},
		{
			Id:     "921010",
			CityId: "9210",
			Name:   "Ayamaru Timur",
		},
		{
			Id:     "921011",
			CityId: "9210",
			Name:   "Mare",
		},
		{
			Id:     "921012",
			CityId: "9210",
			Name:   "Aifat Timur Tengah",
		},
		{
			Id:     "921013",
			CityId: "9210",
			Name:   "Aifat Timur Jauh",
		},
		{
			Id:     "921014",
			CityId: "9210",
			Name:   "Aifat Timur Selatan",
		},
		{
			Id:     "921015",
			CityId: "9210",
			Name:   "Ayamaru Selatan",
		},
		{
			Id:     "921016",
			CityId: "9210",
			Name:   "Ayamaru Jaya",
		},
		{
			Id:     "921017",
			CityId: "9210",
			Name:   "Ayamaru Selatan Jaya",
		},
		{
			Id:     "921018",
			CityId: "9210",
			Name:   "Ayamaru Timur Selatan",
		},
		{
			Id:     "921019",
			CityId: "9210",
			Name:   "Ayamaru Utara Timur",
		},
		{
			Id:     "921020",
			CityId: "9210",
			Name:   "Ayamaru Tengah",
		},
		{
			Id:     "921021",
			CityId: "9210",
			Name:   "Ayamaru Barat",
		},
		{
			Id:     "921022",
			CityId: "9210",
			Name:   "Aitinyo Tengah",
		},
		{
			Id:     "921023",
			CityId: "9210",
			Name:   "Aitinyo Raya",
		},
		{
			Id:     "921024",
			CityId: "9210",
			Name:   "Mare Selatan",
		},
		{
			Id:     "921101",
			CityId: "9211",
			Name:   "Ransiki",
		},
		{
			Id:     "921102",
			CityId: "9211",
			Name:   "Oransbari",
		},
		{
			Id:     "921103",
			CityId: "9211",
			Name:   "Neney",
		},
		{
			Id:     "921104",
			CityId: "9211",
			Name:   "Dataran Isim",
		},
		{
			Id:     "921105",
			CityId: "9211",
			Name:   "Momi Waren",
		},
		{
			Id:     "921106",
			CityId: "9211",
			Name:   "Tahota",
		},
		{
			Id:     "921201",
			CityId: "9212",
			Name:   "Anggi",
		},
		{
			Id:     "921202",
			CityId: "9212",
			Name:   "Anggi Gida",
		},
		{
			Id:     "921203",
			CityId: "9212",
			Name:   "Membey",
		},
		{
			Id:     "921204",
			CityId: "9212",
			Name:   "Sururey",
		},
		{
			Id:     "921205",
			CityId: "9212",
			Name:   "Didohu",
		},
		{
			Id:     "921206",
			CityId: "9212",
			Name:   "Taige",
		},
		{
			Id:     "921207",
			CityId: "9212",
			Name:   "Catubouw",
		},
		{
			Id:     "921208",
			CityId: "9212",
			Name:   "Testega",
		},
		{
			Id:     "921209",
			CityId: "9212",
			Name:   "Minyambaouw",
		},
		{
			Id:     "921210",
			CityId: "9212",
			Name:   "Hingk",
		},
		{
			Id:     "927101",
			CityId: "9271",
			Name:   "Sorong",
		},
		{
			Id:     "927102",
			CityId: "9271",
			Name:   "Sorong Timur",
		},
		{
			Id:     "927103",
			CityId: "9271",
			Name:   "Sorong Barat",
		},
		{
			Id:     "927104",
			CityId: "9271",
			Name:   "Sorong Kepulauan",
		},
		{
			Id:     "927105",
			CityId: "9271",
			Name:   "Sorong Utara",
		},
		{
			Id:     "927106",
			CityId: "9271",
			Name:   "Sorong Manoi",
		},
		{
			Id:     "927107",
			CityId: "9271",
			Name:   "Sorong Kota",
		},
		{
			Id:     "927108",
			CityId: "9271",
			Name:   "Klaurung",
		},
		{
			Id:     "927109",
			CityId: "9271",
			Name:   "Malaimsimsa",
		},
		{
			Id:     "927110",
			CityId: "9271",
			Name:   "Maladum Mes",
		},
		{
			Id:     "930101",
			CityId: "9301",
			Name:   "Merauke",
		},
		{
			Id:     "930102",
			CityId: "9301",
			Name:   "Muting",
		},
		{
			Id:     "930103",
			CityId: "9301",
			Name:   "Okaba",
		},
		{
			Id:     "930104",
			CityId: "9301",
			Name:   "Kimaam",
		},
		{
			Id:     "930105",
			CityId: "9301",
			Name:   "Semangga",
		},
		{
			Id:     "930106",
			CityId: "9301",
			Name:   "Tanah Miring",
		},
		{
			Id:     "930107",
			CityId: "9301",
			Name:   "Jagebob",
		},
		{
			Id:     "930108",
			CityId: "9301",
			Name:   "Sota",
		},
		{
			Id:     "930109",
			CityId: "9301",
			Name:   "Ulilin",
		},
		{
			Id:     "930110",
			CityId: "9301",
			Name:   "Elikobal",
		},
		{
			Id:     "930111",
			CityId: "9301",
			Name:   "Kurik",
		},
		{
			Id:     "930112",
			CityId: "9301",
			Name:   "Naukenjerai",
		},
		{
			Id:     "930113",
			CityId: "9301",
			Name:   "Animha",
		},
		{
			Id:     "930114",
			CityId: "9301",
			Name:   "Malind",
		},
		{
			Id:     "930115",
			CityId: "9301",
			Name:   "Tubang",
		},
		{
			Id:     "930116",
			CityId: "9301",
			Name:   "Ngguti",
		},
		{
			Id:     "930117",
			CityId: "9301",
			Name:   "Kaptel",
		},
		{
			Id:     "930118",
			CityId: "9301",
			Name:   "Tabonji",
		},
		{
			Id:     "930119",
			CityId: "9301",
			Name:   "Waan",
		},
		{
			Id:     "930120",
			CityId: "9301",
			Name:   "Ilwayab",
		},
		{
			Id:     "930121",
			CityId: "9301",
			Name:   "Padua",
		},
		{
			Id:     "930122",
			CityId: "9301",
			Name:   "Kontuar",
		},
		{
			Id:     "930201",
			CityId: "9302",
			Name:   "Mandobo",
		},
		{
			Id:     "930202",
			CityId: "9302",
			Name:   "Mindiptana",
		},
		{
			Id:     "930203",
			CityId: "9302",
			Name:   "Waropko",
		},
		{
			Id:     "930204",
			CityId: "9302",
			Name:   "Kouh",
		},
		{
			Id:     "930205",
			CityId: "9302",
			Name:   "Jair",
		},
		{
			Id:     "930206",
			CityId: "9302",
			Name:   "Bomakia",
		},
		{
			Id:     "930207",
			CityId: "9302",
			Name:   "Kombut",
		},
		{
			Id:     "930208",
			CityId: "9302",
			Name:   "Iniyandit",
		},
		{
			Id:     "930209",
			CityId: "9302",
			Name:   "Arimop",
		},
		{
			Id:     "930210",
			CityId: "9302",
			Name:   "Fofi",
		},
		{
			Id:     "930211",
			CityId: "9302",
			Name:   "Ambatkwi",
		},
		{
			Id:     "930212",
			CityId: "9302",
			Name:   "Manggelum",
		},
		{
			Id:     "930213",
			CityId: "9302",
			Name:   "Firiwage",
		},
		{
			Id:     "930214",
			CityId: "9302",
			Name:   "Yaniruma",
		},
		{
			Id:     "930215",
			CityId: "9302",
			Name:   "Subur",
		},
		{
			Id:     "930216",
			CityId: "9302",
			Name:   "Kombay",
		},
		{
			Id:     "930217",
			CityId: "9302",
			Name:   "Ninati",
		},
		{
			Id:     "930218",
			CityId: "9302",
			Name:   "Sesnuk",
		},
		{
			Id:     "930219",
			CityId: "9302",
			Name:   "Ki",
		},
		{
			Id:     "930220",
			CityId: "9302",
			Name:   "Kawagit",
		},
		{
			Id:     "930301",
			CityId: "9303",
			Name:   "Obaa",
		},
		{
			Id:     "930302",
			CityId: "9303",
			Name:   "Mambioman Bapai",
		},
		{
			Id:     "930303",
			CityId: "9303",
			Name:   "Citak-Mitak",
		},
		{
			Id:     "930304",
			CityId: "9303",
			Name:   "Edera",
		},
		{
			Id:     "930305",
			CityId: "9303",
			Name:   "Haju",
		},
		{
			Id:     "930306",
			CityId: "9303",
			Name:   "Assue",
		},
		{
			Id:     "930307",
			CityId: "9303",
			Name:   "Kaibar",
		},
		{
			Id:     "930308",
			CityId: "9303",
			Name:   "Passue",
		},
		{
			Id:     "930309",
			CityId: "9303",
			Name:   "Minyamur",
		},
		{
			Id:     "930310",
			CityId: "9303",
			Name:   "Venaha",
		},
		{
			Id:     "930311",
			CityId: "9303",
			Name:   "Syahcame",
		},
		{
			Id:     "930312",
			CityId: "9303",
			Name:   "Yakomi",
		},
		{
			Id:     "930313",
			CityId: "9303",
			Name:   "Bamgi",
		},
		{
			Id:     "930314",
			CityId: "9303",
			Name:   "Passue Bawah",
		},
		{
			Id:     "930315",
			CityId: "9303",
			Name:   "Ti Zain",
		},
		{
			Id:     "930401",
			CityId: "9304",
			Name:   "Agats",
		},
		{
			Id:     "930402",
			CityId: "9304",
			Name:   "Atsj",
		},
		{
			Id:     "930403",
			CityId: "9304",
			Name:   "Sawa Erma",
		},
		{
			Id:     "930404",
			CityId: "9304",
			Name:   "Akat",
		},
		{
			Id:     "930405",
			CityId: "9304",
			Name:   "Fayit",
		},
		{
			Id:     "930406",
			CityId: "9304",
			Name:   "Pantai Kasuari",
		},
		{
			Id:     "930407",
			CityId: "9304",
			Name:   "Suator",
		},
		{
			Id:     "930408",
			CityId: "9304",
			Name:   "Suru-suru",
		},
		{
			Id:     "930409",
			CityId: "9304",
			Name:   "Kolf Braza",
		},
		{
			Id:     "930410",
			CityId: "9304",
			Name:   "Unir Sirau",
		},
		{
			Id:     "930411",
			CityId: "9304",
			Name:   "Joerat",
		},
		{
			Id:     "930412",
			CityId: "9304",
			Name:   "Pulau Tiga",
		},
		{
			Id:     "930413",
			CityId: "9304",
			Name:   "Jetsy",
		},
		{
			Id:     "930414",
			CityId: "9304",
			Name:   "Der Koumur",
		},
		{
			Id:     "930415",
			CityId: "9304",
			Name:   "Kopay",
		},
		{
			Id:     "930416",
			CityId: "9304",
			Name:   "Safan",
		},
		{
			Id:     "930417",
			CityId: "9304",
			Name:   "Sirets",
		},
		{
			Id:     "930418",
			CityId: "9304",
			Name:   "Ayip",
		},
		{
			Id:     "930419",
			CityId: "9304",
			Name:   "Betcbamu",
		},
		{
			Id:     "930420",
			CityId: "9304",
			Name:   "Joutu",
		},
		{
			Id:     "930421",
			CityId: "9304",
			Name:   "Aswi",
		},
		{
			Id:     "930422",
			CityId: "9304",
			Name:   "Awyu",
		},
		{
			Id:     "930423",
			CityId: "9304",
			Name:   "Koroway Buluanop",
		},
		{
			Id:     "930424",
			CityId: "9304",
			Name:   "Tomor Birip",
		},
		{
			Id:     "930425",
			CityId: "9304",
			Name:   "Sor Ep",
		},
		{
			Id:     "940101",
			CityId: "9401",
			Name:   "Nabire",
		},
		{
			Id:     "940102",
			CityId: "9401",
			Name:   "Napan",
		},
		{
			Id:     "940103",
			CityId: "9401",
			Name:   "Yaur",
		},
		{
			Id:     "940104",
			CityId: "9401",
			Name:   "Uwapa",
		},
		{
			Id:     "940105",
			CityId: "9401",
			Name:   "Wanggar",
		},
		{
			Id:     "940106",
			CityId: "9401",
			Name:   "Siriwo",
		},
		{
			Id:     "940107",
			CityId: "9401",
			Name:   "Makimi",
		},
		{
			Id:     "940108",
			CityId: "9401",
			Name:   "Teluk Umar",
		},
		{
			Id:     "940109",
			CityId: "9401",
			Name:   "Teluk Kimi",
		},
		{
			Id:     "940110",
			CityId: "9401",
			Name:   "Yaro",
		},
		{
			Id:     "940111",
			CityId: "9401",
			Name:   "Wapoga",
		},
		{
			Id:     "940112",
			CityId: "9401",
			Name:   "Nabire Barat",
		},
		{
			Id:     "940113",
			CityId: "9401",
			Name:   "Moora",
		},
		{
			Id:     "940114",
			CityId: "9401",
			Name:   "Dipa",
		},
		{
			Id:     "940115",
			CityId: "9401",
			Name:   "Menou",
		},
		{
			Id:     "940201",
			CityId: "9402",
			Name:   "Mulia",
		},
		{
			Id:     "940202",
			CityId: "9402",
			Name:   "Ilu",
		},
		{
			Id:     "940203",
			CityId: "9402",
			Name:   "Fawi",
		},
		{
			Id:     "940204",
			CityId: "9402",
			Name:   "Mewoluk",
		},
		{
			Id:     "940205",
			CityId: "9402",
			Name:   "Yamo",
		},
		{
			Id:     "940206",
			CityId: "9402",
			Name:   "Nume",
		},
		{
			Id:     "940207",
			CityId: "9402",
			Name:   "Torere",
		},
		{
			Id:     "940208",
			CityId: "9402",
			Name:   "Tingginambut",
		},
		{
			Id:     "940209",
			CityId: "9402",
			Name:   "Pagaleme",
		},
		{
			Id:     "940210",
			CityId: "9402",
			Name:   "Gurage",
		},
		{
			Id:     "940211",
			CityId: "9402",
			Name:   "Irimuli",
		},
		{
			Id:     "940212",
			CityId: "9402",
			Name:   "Muara",
		},
		{
			Id:     "940213",
			CityId: "9402",
			Name:   "Ilamburawi",
		},
		{
			Id:     "940214",
			CityId: "9402",
			Name:   "Yambi",
		},
		{
			Id:     "940215",
			CityId: "9402",
			Name:   "Lumo",
		},
		{
			Id:     "940216",
			CityId: "9402",
			Name:   "Molanikime",
		},
		{
			Id:     "940217",
			CityId: "9402",
			Name:   "Dokome",
		},
		{
			Id:     "940218",
			CityId: "9402",
			Name:   "Kalome",
		},
		{
			Id:     "940219",
			CityId: "9402",
			Name:   "Wanwi",
		},
		{
			Id:     "940220",
			CityId: "9402",
			Name:   "Yamoneri",
		},
		{
			Id:     "940221",
			CityId: "9402",
			Name:   "Waegi",
		},
		{
			Id:     "940222",
			CityId: "9402",
			Name:   "Nioga",
		},
		{
			Id:     "940223",
			CityId: "9402",
			Name:   "Gubume",
		},
		{
			Id:     "940224",
			CityId: "9402",
			Name:   "Taganombak",
		},
		{
			Id:     "940225",
			CityId: "9402",
			Name:   "Dagai",
		},
		{
			Id:     "940226",
			CityId: "9402",
			Name:   "Kiyage",
		},
		{
			Id:     "940301",
			CityId: "9403",
			Name:   "Paniai Timur",
		},
		{
			Id:     "940302",
			CityId: "9403",
			Name:   "Paniai Barat",
		},
		{
			Id:     "940303",
			CityId: "9403",
			Name:   "Aradide",
		},
		{
			Id:     "940304",
			CityId: "9403",
			Name:   "Bogabaida",
		},
		{
			Id:     "940305",
			CityId: "9403",
			Name:   "Bibida",
		},
		{
			Id:     "940306",
			CityId: "9403",
			Name:   "Dumadama",
		},
		{
			Id:     "940307",
			CityId: "9403",
			Name:   "Siriwo",
		},
		{
			Id:     "940308",
			CityId: "9403",
			Name:   "Kebo",
		},
		{
			Id:     "940309",
			CityId: "9403",
			Name:   "Yatamo",
		},
		{
			Id:     "940310",
			CityId: "9403",
			Name:   "Ekadide",
		},
		{
			Id:     "940311",
			CityId: "9403",
			Name:   "Wegee Muka",
		},
		{
			Id:     "940312",
			CityId: "9403",
			Name:   "Wegee Bino",
		},
		{
			Id:     "940313",
			CityId: "9403",
			Name:   "Pugo Dagi",
		},
		{
			Id:     "940314",
			CityId: "9403",
			Name:   "Muye",
		},
		{
			Id:     "940315",
			CityId: "9403",
			Name:   "Nakama",
		},
		{
			Id:     "940316",
			CityId: "9403",
			Name:   "Teluk Deya",
		},
		{
			Id:     "940317",
			CityId: "9403",
			Name:   "Yagai",
		},
		{
			Id:     "940318",
			CityId: "9403",
			Name:   "Youtadi",
		},
		{
			Id:     "940319",
			CityId: "9403",
			Name:   "Baya Biru",
		},
		{
			Id:     "940320",
			CityId: "9403",
			Name:   "Deiyai Miyo",
		},
		{
			Id:     "940321",
			CityId: "9403",
			Name:   "Dogomo",
		},
		{
			Id:     "940322",
			CityId: "9403",
			Name:   "Aweida",
		},
		{
			Id:     "940323",
			CityId: "9403",
			Name:   "Topiyai",
		},
		{
			Id:     "940324",
			CityId: "9403",
			Name:   "Fajar Timur",
		},
		{
			Id:     "940401",
			CityId: "9404",
			Name:   "Mimika Baru",
		},
		{
			Id:     "940402",
			CityId: "9404",
			Name:   "Agimuga",
		},
		{
			Id:     "940403",
			CityId: "9404",
			Name:   "Mimika Timur",
		},
		{
			Id:     "940404",
			CityId: "9404",
			Name:   "Mimika Barat",
		},
		{
			Id:     "940405",
			CityId: "9404",
			Name:   "Jita",
		},
		{
			Id:     "940406",
			CityId: "9404",
			Name:   "Jila",
		},
		{
			Id:     "940407",
			CityId: "9404",
			Name:   "Mimika Timur Jauh",
		},
		{
			Id:     "940408",
			CityId: "9404",
			Name:   "Mimika Tengah",
		},
		{
			Id:     "940409",
			CityId: "9404",
			Name:   "Kuala Kencana",
		},
		{
			Id:     "940410",
			CityId: "9404",
			Name:   "Tembagapura",
		},
		{
			Id:     "940411",
			CityId: "9404",
			Name:   "Mimika Barat Jauh",
		},
		{
			Id:     "940412",
			CityId: "9404",
			Name:   "Mimika Barat Tengah",
		},
		{
			Id:     "940413",
			CityId: "9404",
			Name:   "Kwamki Narama",
		},
		{
			Id:     "940414",
			CityId: "9404",
			Name:   "Hoya",
		},
		{
			Id:     "940415",
			CityId: "9404",
			Name:   "Iwaka",
		},
		{
			Id:     "940416",
			CityId: "9404",
			Name:   "Wania",
		},
		{
			Id:     "940417",
			CityId: "9404",
			Name:   "Amar",
		},
		{
			Id:     "940418",
			CityId: "9404",
			Name:   "Alama",
		},
		{
			Id:     "940501",
			CityId: "9405",
			Name:   "Ilaga",
		},
		{
			Id:     "940502",
			CityId: "9405",
			Name:   "Wangbe",
		},
		{
			Id:     "940503",
			CityId: "9405",
			Name:   "Beoga",
		},
		{
			Id:     "940504",
			CityId: "9405",
			Name:   "Doufo",
		},
		{
			Id:     "940505",
			CityId: "9405",
			Name:   "Pogoma",
		},
		{
			Id:     "940506",
			CityId: "9405",
			Name:   "Sinak",
		},
		{
			Id:     "940507",
			CityId: "9405",
			Name:   "Agandugume",
		},
		{
			Id:     "940508",
			CityId: "9405",
			Name:   "Gome",
		},
		{
			Id:     "940509",
			CityId: "9405",
			Name:   "Dervos",
		},
		{
			Id:     "940510",
			CityId: "9405",
			Name:   "Beoga Barat",
		},
		{
			Id:     "940511",
			CityId: "9405",
			Name:   "Beoga Timur",
		},
		{
			Id:     "940512",
			CityId: "9405",
			Name:   "Ogamanim",
		},
		{
			Id:     "940513",
			CityId: "9405",
			Name:   "Kembru",
		},
		{
			Id:     "940514",
			CityId: "9405",
			Name:   "Bina",
		},
		{
			Id:     "940515",
			CityId: "9405",
			Name:   "Sinak Barat",
		},
		{
			Id:     "940516",
			CityId: "9405",
			Name:   "Mage'abume",
		},
		{
			Id:     "940517",
			CityId: "9405",
			Name:   "Yugumuak",
		},
		{
			Id:     "940518",
			CityId: "9405",
			Name:   "Ilaga Utara",
		},
		{
			Id:     "940519",
			CityId: "9405",
			Name:   "Mabugi",
		},
		{
			Id:     "940520",
			CityId: "9405",
			Name:   "Omukia",
		},
		{
			Id:     "940521",
			CityId: "9405",
			Name:   "Lambewi",
		},
		{
			Id:     "940522",
			CityId: "9405",
			Name:   "Oneri",
		},
		{
			Id:     "940523",
			CityId: "9405",
			Name:   "Amungkalpia",
		},
		{
			Id:     "940524",
			CityId: "9405",
			Name:   "Gome Utara",
		},
		{
			Id:     "940525",
			CityId: "9405",
			Name:   "Erelmakawia",
		},
		{
			Id:     "940601",
			CityId: "9406",
			Name:   "Kamu",
		},
		{
			Id:     "940602",
			CityId: "9406",
			Name:   "Mapia",
		},
		{
			Id:     "940603",
			CityId: "9406",
			Name:   "Piyaiye",
		},
		{
			Id:     "940604",
			CityId: "9406",
			Name:   "Kamu Utara",
		},
		{
			Id:     "940605",
			CityId: "9406",
			Name:   "Sukikai Selatan",
		},
		{
			Id:     "940606",
			CityId: "9406",
			Name:   "Mapia Barat",
		},
		{
			Id:     "940607",
			CityId: "9406",
			Name:   "Kamu Selatan",
		},
		{
			Id:     "940608",
			CityId: "9406",
			Name:   "Kamu Timur",
		},
		{
			Id:     "940609",
			CityId: "9406",
			Name:   "Mapia Tengah",
		},
		{
			Id:     "940610",
			CityId: "9406",
			Name:   "Dogiyai",
		},
		{
			Id:     "940701",
			CityId: "9407",
			Name:   "Sugapa",
		},
		{
			Id:     "940702",
			CityId: "9407",
			Name:   "Homeyo",
		},
		{
			Id:     "940703",
			CityId: "9407",
			Name:   "Wandai",
		},
		{
			Id:     "940704",
			CityId: "9407",
			Name:   "Biandoga",
		},
		{
			Id:     "940705",
			CityId: "9407",
			Name:   "Agisiga",
		},
		{
			Id:     "940706",
			CityId: "9407",
			Name:   "Hitadipa",
		},
		{
			Id:     "940707",
			CityId: "9407",
			Name:   "Ugimba",
		},
		{
			Id:     "940708",
			CityId: "9407",
			Name:   "Tomosiga",
		},
		{
			Id:     "940801",
			CityId: "9408",
			Name:   "Tigi",
		},
		{
			Id:     "940802",
			CityId: "9408",
			Name:   "Tigi Timur",
		},
		{
			Id:     "940803",
			CityId: "9408",
			Name:   "Bowobado",
		},
		{
			Id:     "940804",
			CityId: "9408",
			Name:   "Tigi Barat",
		},
		{
			Id:     "940805",
			CityId: "9408",
			Name:   "Kapiraya",
		},
		{
			Id:     "950101",
			CityId: "9501",
			Name:   "Wamena",
		},
		{
			Id:     "950102",
			CityId: "9501",
			Name:   "Kurulu",
		},
		{
			Id:     "950103",
			CityId: "9501",
			Name:   "Asologaima",
		},
		{
			Id:     "950104",
			CityId: "9501",
			Name:   "Hubikosi",
		},
		{
			Id:     "950105",
			CityId: "9501",
			Name:   "Bolakme",
		},
		{
			Id:     "950106",
			CityId: "9501",
			Name:   "Walelagama",
		},
		{
			Id:     "950107",
			CityId: "9501",
			Name:   "Musatfak",
		},
		{
			Id:     "950108",
			CityId: "9501",
			Name:   "Wolo",
		},
		{
			Id:     "950109",
			CityId: "9501",
			Name:   "Asolokobal",
		},
		{
			Id:     "950110",
			CityId: "9501",
			Name:   "Pelebaga",
		},
		{
			Id:     "950111",
			CityId: "9501",
			Name:   "Yalengga",
		},
		{
			Id:     "950112",
			CityId: "9501",
			Name:   "Trikora",
		},
		{
			Id:     "950113",
			CityId: "9501",
			Name:   "Napua",
		},
		{
			Id:     "950114",
			CityId: "9501",
			Name:   "Walaik",
		},
		{
			Id:     "950115",
			CityId: "9501",
			Name:   "Wouma",
		},
		{
			Id:     "950116",
			CityId: "9501",
			Name:   "Hubikiak",
		},
		{
			Id:     "950117",
			CityId: "9501",
			Name:   "Ibele",
		},
		{
			Id:     "950118",
			CityId: "9501",
			Name:   "Taelarek",
		},
		{
			Id:     "950119",
			CityId: "9501",
			Name:   "Itlay Hisage",
		},
		{
			Id:     "950120",
			CityId: "9501",
			Name:   "Siepkosi",
		},
		{
			Id:     "950121",
			CityId: "9501",
			Name:   "Usilimo",
		},
		{
			Id:     "950122",
			CityId: "9501",
			Name:   "Wita Waya",
		},
		{
			Id:     "950123",
			CityId: "9501",
			Name:   "Libarek",
		},
		{
			Id:     "950124",
			CityId: "9501",
			Name:   "Wadangku",
		},
		{
			Id:     "950125",
			CityId: "9501",
			Name:   "Pisugi",
		},
		{
			Id:     "950126",
			CityId: "9501",
			Name:   "Koragi",
		},
		{
			Id:     "950127",
			CityId: "9501",
			Name:   "Tagime",
		},
		{
			Id:     "950128",
			CityId: "9501",
			Name:   "Molagalome",
		},
		{
			Id:     "950129",
			CityId: "9501",
			Name:   "Tagineri",
		},
		{
			Id:     "950130",
			CityId: "9501",
			Name:   "Silo Karno Doga",
		},
		{
			Id:     "950131",
			CityId: "9501",
			Name:   "Piramid",
		},
		{
			Id:     "950132",
			CityId: "9501",
			Name:   "Muliama",
		},
		{
			Id:     "950133",
			CityId: "9501",
			Name:   "Bugi",
		},
		{
			Id:     "950134",
			CityId: "9501",
			Name:   "Bpiri",
		},
		{
			Id:     "950135",
			CityId: "9501",
			Name:   "Welesi",
		},
		{
			Id:     "950136",
			CityId: "9501",
			Name:   "Asotipo",
		},
		{
			Id:     "950137",
			CityId: "9501",
			Name:   "Maima",
		},
		{
			Id:     "950138",
			CityId: "9501",
			Name:   "Popugoba",
		},
		{
			Id:     "950139",
			CityId: "9501",
			Name:   "Wame",
		},
		{
			Id:     "950140",
			CityId: "9501",
			Name:   "Wesaput",
		},
		{
			Id:     "950201",
			CityId: "9502",
			Name:   "Oksibil",
		},
		{
			Id:     "950202",
			CityId: "9502",
			Name:   "Kiwirok",
		},
		{
			Id:     "950203",
			CityId: "9502",
			Name:   "Okbibab",
		},
		{
			Id:     "950204",
			CityId: "9502",
			Name:   "Iwur",
		},
		{
			Id:     "950205",
			CityId: "9502",
			Name:   "Batom",
		},
		{
			Id:     "950206",
			CityId: "9502",
			Name:   "Borme",
		},
		{
			Id:     "950207",
			CityId: "9502",
			Name:   "Kiwirok Timur",
		},
		{
			Id:     "950208",
			CityId: "9502",
			Name:   "Aboy",
		},
		{
			Id:     "950209",
			CityId: "9502",
			Name:   "Pepera",
		},
		{
			Id:     "950210",
			CityId: "9502",
			Name:   "Bime",
		},
		{
			Id:     "950211",
			CityId: "9502",
			Name:   "Alemsom",
		},
		{
			Id:     "950212",
			CityId: "9502",
			Name:   "Okbape",
		},
		{
			Id:     "950213",
			CityId: "9502",
			Name:   "Kalomdol",
		},
		{
			Id:     "950214",
			CityId: "9502",
			Name:   "Oksop",
		},
		{
			Id:     "950215",
			CityId: "9502",
			Name:   "Serambakon",
		},
		{
			Id:     "950216",
			CityId: "9502",
			Name:   "Ok Aom",
		},
		{
			Id:     "950217",
			CityId: "9502",
			Name:   "Kawor",
		},
		{
			Id:     "950218",
			CityId: "9502",
			Name:   "Awinbon",
		},
		{
			Id:     "950219",
			CityId: "9502",
			Name:   "Tarup",
		},
		{
			Id:     "950220",
			CityId: "9502",
			Name:   "Okhika",
		},
		{
			Id:     "950221",
			CityId: "9502",
			Name:   "Oksamol",
		},
		{
			Id:     "950222",
			CityId: "9502",
			Name:   "Oklip",
		},
		{
			Id:     "950223",
			CityId: "9502",
			Name:   "Okbemtau",
		},
		{
			Id:     "950224",
			CityId: "9502",
			Name:   "Oksebang",
		},
		{
			Id:     "950225",
			CityId: "9502",
			Name:   "Okbab",
		},
		{
			Id:     "950226",
			CityId: "9502",
			Name:   "Batani",
		},
		{
			Id:     "950227",
			CityId: "9502",
			Name:   "Weime",
		},
		{
			Id:     "950228",
			CityId: "9502",
			Name:   "Murkim",
		},
		{
			Id:     "950229",
			CityId: "9502",
			Name:   "Mofinop",
		},
		{
			Id:     "950230",
			CityId: "9502",
			Name:   "Jetfa",
		},
		{
			Id:     "950231",
			CityId: "9502",
			Name:   "Teiraplu",
		},
		{
			Id:     "950232",
			CityId: "9502",
			Name:   "Eipumek",
		},
		{
			Id:     "950233",
			CityId: "9502",
			Name:   "Pamek",
		},
		{
			Id:     "950234",
			CityId: "9502",
			Name:   "Nongme",
		},
		{
			Id:     "950301",
			CityId: "9503",
			Name:   "Kurima",
		},
		{
			Id:     "950302",
			CityId: "9503",
			Name:   "Anggruk",
		},
		{
			Id:     "950303",
			CityId: "9503",
			Name:   "Ninia",
		},
		{
			Id:     "950304",
			CityId: "9503",
			Name:   "Silimo",
		},
		{
			Id:     "950305",
			CityId: "9503",
			Name:   "Samenage",
		},
		{
			Id:     "950306",
			CityId: "9503",
			Name:   "Nalca",
		},
		{
			Id:     "950307",
			CityId: "9503",
			Name:   "Dekai",
		},
		{
			Id:     "950308",
			CityId: "9503",
			Name:   "Obio",
		},
		{
			Id:     "950309",
			CityId: "9503",
			Name:   "Suru Suru",
		},
		{
			Id:     "950310",
			CityId: "9503",
			Name:   "Wusama",
		},
		{
			Id:     "950311",
			CityId: "9503",
			Name:   "Amuma",
		},
		{
			Id:     "950312",
			CityId: "9503",
			Name:   "Musaik",
		},
		{
			Id:     "950313",
			CityId: "9503",
			Name:   "Pasema",
		},
		{
			Id:     "950314",
			CityId: "9503",
			Name:   "Hogio",
		},
		{
			Id:     "950315",
			CityId: "9503",
			Name:   "Mugi",
		},
		{
			Id:     "950316",
			CityId: "9503",
			Name:   "Soba",
		},
		{
			Id:     "950317",
			CityId: "9503",
			Name:   "Werima",
		},
		{
			Id:     "950318",
			CityId: "9503",
			Name:   "Tangma",
		},
		{
			Id:     "950319",
			CityId: "9503",
			Name:   "Ukha",
		},
		{
			Id:     "950320",
			CityId: "9503",
			Name:   "Panggema",
		},
		{
			Id:     "950321",
			CityId: "9503",
			Name:   "Kosarek",
		},
		{
			Id:     "950322",
			CityId: "9503",
			Name:   "Nipsan",
		},
		{
			Id:     "950323",
			CityId: "9503",
			Name:   "Ubahak",
		},
		{
			Id:     "950324",
			CityId: "9503",
			Name:   "Pronggoli",
		},
		{
			Id:     "950325",
			CityId: "9503",
			Name:   "Walma",
		},
		{
			Id:     "950326",
			CityId: "9503",
			Name:   "Yahuliambut",
		},
		{
			Id:     "950327",
			CityId: "9503",
			Name:   "Hereapini",
		},
		{
			Id:     "950328",
			CityId: "9503",
			Name:   "Ubalihi",
		},
		{
			Id:     "950329",
			CityId: "9503",
			Name:   "Talambo",
		},
		{
			Id:     "950330",
			CityId: "9503",
			Name:   "Puldama",
		},
		{
			Id:     "950331",
			CityId: "9503",
			Name:   "Endomen",
		},
		{
			Id:     "950332",
			CityId: "9503",
			Name:   "Kona",
		},
		{
			Id:     "950333",
			CityId: "9503",
			Name:   "Dirwemna",
		},
		{
			Id:     "950334",
			CityId: "9503",
			Name:   "Holuon",
		},
		{
			Id:     "950335",
			CityId: "9503",
			Name:   "Lolat",
		},
		{
			Id:     "950336",
			CityId: "9503",
			Name:   "Soloikma",
		},
		{
			Id:     "950337",
			CityId: "9503",
			Name:   "Sela",
		},
		{
			Id:     "950338",
			CityId: "9503",
			Name:   "Korupun",
		},
		{
			Id:     "950339",
			CityId: "9503",
			Name:   "Langda",
		},
		{
			Id:     "950340",
			CityId: "9503",
			Name:   "Bomela",
		},
		{
			Id:     "950341",
			CityId: "9503",
			Name:   "Suntamon",
		},
		{
			Id:     "950342",
			CityId: "9503",
			Name:   "Seradala",
		},
		{
			Id:     "950343",
			CityId: "9503",
			Name:   "Sobaham",
		},
		{
			Id:     "950344",
			CityId: "9503",
			Name:   "Kabianggama",
		},
		{
			Id:     "950345",
			CityId: "9503",
			Name:   "Kwelamdua",
		},
		{
			Id:     "950346",
			CityId: "9503",
			Name:   "Kwikma",
		},
		{
			Id:     "950347",
			CityId: "9503",
			Name:   "Hilipuk",
		},
		{
			Id:     "950348",
			CityId: "9503",
			Name:   "Duram",
		},
		{
			Id:     "950349",
			CityId: "9503",
			Name:   "Yogosem",
		},
		{
			Id:     "950350",
			CityId: "9503",
			Name:   "Kayo",
		},
		{
			Id:     "950351",
			CityId: "9503",
			Name:   "Sumo",
		},
		{
			Id:     "950401",
			CityId: "9504",
			Name:   "Karubaga",
		},
		{
			Id:     "950402",
			CityId: "9504",
			Name:   "Bokondini",
		},
		{
			Id:     "950403",
			CityId: "9504",
			Name:   "Kanggime",
		},
		{
			Id:     "950404",
			CityId: "9504",
			Name:   "Kembu",
		},
		{
			Id:     "950405",
			CityId: "9504",
			Name:   "Goyage",
		},
		{
			Id:     "950406",
			CityId: "9504",
			Name:   "Wunim",
		},
		{
			Id:     "950407",
			CityId: "9504",
			Name:   "Wina",
		},
		{
			Id:     "950408",
			CityId: "9504",
			Name:   "Umagi",
		},
		{
			Id:     "950409",
			CityId: "9504",
			Name:   "Panaga",
		},
		{
			Id:     "950410",
			CityId: "9504",
			Name:   "Woniki",
		},
		{
			Id:     "950411",
			CityId: "9504",
			Name:   "Kubu",
		},
		{
			Id:     "950412",
			CityId: "9504",
			Name:   "Konda/ Kondaga",
		},
		{
			Id:     "950413",
			CityId: "9504",
			Name:   "Nelawi",
		},
		{
			Id:     "950414",
			CityId: "9504",
			Name:   "Kuari",
		},
		{
			Id:     "950415",
			CityId: "9504",
			Name:   "Bokoneri",
		},
		{
			Id:     "950416",
			CityId: "9504",
			Name:   "Bewani",
		},
		{
			Id:     "950417",
			CityId: "9504",
			Name:   "Nabunage",
		},
		{
			Id:     "950418",
			CityId: "9504",
			Name:   "Gilubandu",
		},
		{
			Id:     "950419",
			CityId: "9504",
			Name:   "Nunggawi",
		},
		{
			Id:     "950420",
			CityId: "9504",
			Name:   "Gundagi",
		},
		{
			Id:     "950421",
			CityId: "9504",
			Name:   "Numba",
		},
		{
			Id:     "950422",
			CityId: "9504",
			Name:   "Timori",
		},
		{
			Id:     "950423",
			CityId: "9504",
			Name:   "Dundu",
		},
		{
			Id:     "950424",
			CityId: "9504",
			Name:   "Geya",
		},
		{
			Id:     "950425",
			CityId: "9504",
			Name:   "Egiam",
		},
		{
			Id:     "950426",
			CityId: "9504",
			Name:   "Poganeri",
		},
		{
			Id:     "950427",
			CityId: "9504",
			Name:   "Kamboneri",
		},
		{
			Id:     "950428",
			CityId: "9504",
			Name:   "Airgaram",
		},
		{
			Id:     "950429",
			CityId: "9504",
			Name:   "Wari/Taiyeve II",
		},
		{
			Id:     "950430",
			CityId: "9504",
			Name:   "Dow",
		},
		{
			Id:     "950431",
			CityId: "9504",
			Name:   "Tagineri",
		},
		{
			Id:     "950432",
			CityId: "9504",
			Name:   "Yuneri",
		},
		{
			Id:     "950433",
			CityId: "9504",
			Name:   "Wakuwo",
		},
		{
			Id:     "950434",
			CityId: "9504",
			Name:   "Gika",
		},
		{
			Id:     "950435",
			CityId: "9504",
			Name:   "Telenggeme",
		},
		{
			Id:     "950436",
			CityId: "9504",
			Name:   "Anawi",
		},
		{
			Id:     "950437",
			CityId: "9504",
			Name:   "Wenam",
		},
		{
			Id:     "950438",
			CityId: "9504",
			Name:   "Wugi",
		},
		{
			Id:     "950439",
			CityId: "9504",
			Name:   "Danime",
		},
		{
			Id:     "950440",
			CityId: "9504",
			Name:   "Tagime",
		},
		{
			Id:     "950441",
			CityId: "9504",
			Name:   "Kai",
		},
		{
			Id:     "950442",
			CityId: "9504",
			Name:   "Aweku",
		},
		{
			Id:     "950443",
			CityId: "9504",
			Name:   "Bogonuk",
		},
		{
			Id:     "950444",
			CityId: "9504",
			Name:   "Li Anogomma",
		},
		{
			Id:     "950445",
			CityId: "9504",
			Name:   "Biuk",
		},
		{
			Id:     "950446",
			CityId: "9504",
			Name:   "Yuko",
		},
		{
			Id:     "950501",
			CityId: "9505",
			Name:   "Kobakma",
		},
		{
			Id:     "950502",
			CityId: "9505",
			Name:   "Kelila",
		},
		{
			Id:     "950503",
			CityId: "9505",
			Name:   "Eragayam",
		},
		{
			Id:     "950504",
			CityId: "9505",
			Name:   "Megambilis",
		},
		{
			Id:     "950505",
			CityId: "9505",
			Name:   "Ilugwa",
		},
		{
			Id:     "950601",
			CityId: "9506",
			Name:   "Elelim",
		},
		{
			Id:     "950602",
			CityId: "9506",
			Name:   "Apalapsili",
		},
		{
			Id:     "950603",
			CityId: "9506",
			Name:   "Abenaho",
		},
		{
			Id:     "950604",
			CityId: "9506",
			Name:   "Benawa",
		},
		{
			Id:     "950605",
			CityId: "9506",
			Name:   "Welarek",
		},
		{
			Id:     "950701",
			CityId: "9507",
			Name:   "Tiom",
		},
		{
			Id:     "950702",
			CityId: "9507",
			Name:   "Pirime",
		},
		{
			Id:     "950703",
			CityId: "9507",
			Name:   "Makki",
		},
		{
			Id:     "950704",
			CityId: "9507",
			Name:   "Gamelia",
		},
		{
			Id:     "950705",
			CityId: "9507",
			Name:   "Dimba",
		},
		{
			Id:     "950706",
			CityId: "9507",
			Name:   "Melagineri",
		},
		{
			Id:     "950707",
			CityId: "9507",
			Name:   "Balingga",
		},
		{
			Id:     "950708",
			CityId: "9507",
			Name:   "Tiomneri",
		},
		{
			Id:     "950709",
			CityId: "9507",
			Name:   "Kuyawage",
		},
		{
			Id:     "950710",
			CityId: "9507",
			Name:   "Poga",
		},
		{
			Id:     "950711",
			CityId: "9507",
			Name:   "Niname",
		},
		{
			Id:     "950712",
			CityId: "9507",
			Name:   "Nogi",
		},
		{
			Id:     "950713",
			CityId: "9507",
			Name:   "Yiginua",
		},
		{
			Id:     "950714",
			CityId: "9507",
			Name:   "Tiom Ollo",
		},
		{
			Id:     "950715",
			CityId: "9507",
			Name:   "Yugungwi",
		},
		{
			Id:     "950716",
			CityId: "9507",
			Name:   "Mokoni",
		},
		{
			Id:     "950717",
			CityId: "9507",
			Name:   "Wereka",
		},
		{
			Id:     "950718",
			CityId: "9507",
			Name:   "Milimbo",
		},
		{
			Id:     "950719",
			CityId: "9507",
			Name:   "Wiringgambut",
		},
		{
			Id:     "950720",
			CityId: "9507",
			Name:   "Gollo",
		},
		{
			Id:     "950721",
			CityId: "9507",
			Name:   "Awina",
		},
		{
			Id:     "950722",
			CityId: "9507",
			Name:   "Ayumnati",
		},
		{
			Id:     "950723",
			CityId: "9507",
			Name:   "Wano Barat",
		},
		{
			Id:     "950724",
			CityId: "9507",
			Name:   "Goa Balim",
		},
		{
			Id:     "950725",
			CityId: "9507",
			Name:   "Bruwa",
		},
		{
			Id:     "950726",
			CityId: "9507",
			Name:   "Balingga Barat",
		},
		{
			Id:     "950727",
			CityId: "9507",
			Name:   "Gupura",
		},
		{
			Id:     "950728",
			CityId: "9507",
			Name:   "Kolawa",
		},
		{
			Id:     "950729",
			CityId: "9507",
			Name:   "Gelok Beam",
		},
		{
			Id:     "950730",
			CityId: "9507",
			Name:   "Kuly Lanny",
		},
		{
			Id:     "950731",
			CityId: "9507",
			Name:   "Lannyna",
		},
		{
			Id:     "950732",
			CityId: "9507",
			Name:   "Karu",
		},
		{
			Id:     "950733",
			CityId: "9507",
			Name:   "Yiluk",
		},
		{
			Id:     "950734",
			CityId: "9507",
			Name:   "Guna",
		},
		{
			Id:     "950735",
			CityId: "9507",
			Name:   "Kelulome",
		},
		{
			Id:     "950736",
			CityId: "9507",
			Name:   "Nikogwe",
		},
		{
			Id:     "950737",
			CityId: "9507",
			Name:   "Muara",
		},
		{
			Id:     "950738",
			CityId: "9507",
			Name:   "Buguk Gona",
		},
		{
			Id:     "950739",
			CityId: "9507",
			Name:   "Melagi",
		},
		{
			Id:     "950801",
			CityId: "9508",
			Name:   "Kenyam",
		},
		{
			Id:     "950802",
			CityId: "9508",
			Name:   "Mapenduma",
		},
		{
			Id:     "950803",
			CityId: "9508",
			Name:   "Yigi",
		},
		{
			Id:     "950804",
			CityId: "9508",
			Name:   "Wosak",
		},
		{
			Id:     "950805",
			CityId: "9508",
			Name:   "Geselma",
		},
		{
			Id:     "950806",
			CityId: "9508",
			Name:   "Mugi",
		},
		{
			Id:     "950807",
			CityId: "9508",
			Name:   "Mbuwa",
		},
		{
			Id:     "950808",
			CityId: "9508",
			Name:   "Gearek",
		},
		{
			Id:     "950809",
			CityId: "9508",
			Name:   "Koroptak",
		},
		{
			Id:     "950810",
			CityId: "9508",
			Name:   "Kegayem",
		},
		{
			Id:     "950811",
			CityId: "9508",
			Name:   "Paro",
		},
		{
			Id:     "950812",
			CityId: "9508",
			Name:   "Mebarok",
		},
		{
			Id:     "950813",
			CityId: "9508",
			Name:   "Yenggelo",
		},
		{
			Id:     "950814",
			CityId: "9508",
			Name:   "Kilmid",
		},
		{
			Id:     "950815",
			CityId: "9508",
			Name:   "Alama",
		},
		{
			Id:     "950816",
			CityId: "9508",
			Name:   "Yal",
		},
		{
			Id:     "950817",
			CityId: "9508",
			Name:   "Mam",
		},
		{
			Id:     "950818",
			CityId: "9508",
			Name:   "Dal",
		},
		{
			Id:     "950819",
			CityId: "9508",
			Name:   "Nirkuri",
		},
		{
			Id:     "950820",
			CityId: "9508",
			Name:   "Inikgal",
		},
		{
			Id:     "950821",
			CityId: "9508",
			Name:   "Iniye",
		},
		{
			Id:     "950822",
			CityId: "9508",
			Name:   "Mbulmu Yalma",
		},
		{
			Id:     "950823",
			CityId: "9508",
			Name:   "Mbua Tengah",
		},
		{
			Id:     "950824",
			CityId: "9508",
			Name:   "Embetpen",
		},
		{
			Id:     "950825",
			CityId: "9508",
			Name:   "Kora",
		},
		{
			Id:     "950826",
			CityId: "9508",
			Name:   "Wusi",
		},
		{
			Id:     "950827",
			CityId: "9508",
			Name:   "Pija",
		},
		{
			Id:     "950828",
			CityId: "9508",
			Name:   "Moba",
		},
		{
			Id:     "950829",
			CityId: "9508",
			Name:   "Wutpaga",
		},
		{
			Id:     "950830",
			CityId: "9508",
			Name:   "Nenggeagin",
		},
		{
			Id:     "950831",
			CityId: "9508",
			Name:   "Krepkuri",
		},
		{
			Id:     "950832",
			CityId: "9508",
			Name:   "Pasir Putih",
		},
		{
			Id:     "950833",
			CityId: "6401",
			Name:   "Pasir Belengkong",
		},
		{
			Id:     "950834",
			CityId: "7501",
			Name:   "Telaga Jaya",
		},
		{
			Id:     "950836",
			CityId: "1375",
			Name:   "Guguk Panjang",
		},
		{
			Id:     "950837",
			CityId: "1405",
			Name:   "Segati",
		},
		{
			Id:     "950838",
			CityId: "1901",
			Name:   "Sinar Baru",
		},
		{
			Id:     "950839",
			CityId: "9304",
			Name:   "Yakapis",
		},
	}

	for _, subdistrict := range subdistricts {
		if err := s.db.Where(entities.Subdistrict{Name: subdistrict.Name}).
			FirstOrCreate(&subdistrict).Error; err != nil {
			log.Fatalf("failed to create subdistrict: %v", err)
		}
	}
}
