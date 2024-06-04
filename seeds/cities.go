package seeds

import (
	"log"

	"capstone/entities"
)

func (s Seed) SeedCities() {
	cities := []entities.City{
		{
			Id:         "1101",
			ProvinceId: "11",
			Name:       "KAB. ACEH SELATAN",
		},
		{
			Id:         "1102",
			ProvinceId: "11",
			Name:       "KAB. ACEH TENGGARA",
		},
		{
			Id:         "1103",
			ProvinceId: "11",
			Name:       "KAB. ACEH TIMUR",
		},
		{
			Id:         "1104",
			ProvinceId: "11",
			Name:       "KAB. ACEH TENGAH",
		},
		{
			Id:         "1105",
			ProvinceId: "11",
			Name:       "KAB. ACEH BARAT",
		},
		{
			Id:         "1106",
			ProvinceId: "11",
			Name:       "KAB. ACEH BESAR",
		},
		{
			Id:         "1107",
			ProvinceId: "11",
			Name:       "KAB. PIDIE",
		},
		{
			Id:         "1108",
			ProvinceId: "11",
			Name:       "KAB. ACEH UTARA",
		},
		{
			Id:         "1109",
			ProvinceId: "11",
			Name:       "KAB. SIMEULUE",
		},
		{
			Id:         "1110",
			ProvinceId: "11",
			Name:       "KAB. ACEH SINGKIL",
		},
		{
			Id:         "1111",
			ProvinceId: "11",
			Name:       "KAB. BIREUEN",
		},
		{
			Id:         "1112",
			ProvinceId: "11",
			Name:       "KAB. ACEH BARAT DAYA",
		},
		{
			Id:         "1113",
			ProvinceId: "11",
			Name:       "KAB. GAYO LUES",
		},
		{
			Id:         "1114",
			ProvinceId: "11",
			Name:       "KAB. ACEH JAYA",
		},
		{
			Id:         "1115",
			ProvinceId: "11",
			Name:       "KAB. NAGAN RAYA",
		},
		{
			Id:         "1116",
			ProvinceId: "11",
			Name:       "KAB. ACEH TAMIANG",
		},
		{
			Id:         "1117",
			ProvinceId: "11",
			Name:       "KAB. BENER MERIAH",
		},
		{
			Id:         "1118",
			ProvinceId: "11",
			Name:       "KAB. PIDIE JAYA",
		},
		{
			Id:         "1171",
			ProvinceId: "11",
			Name:       "KOTA BANDA ACEH",
		},
		{
			Id:         "1172",
			ProvinceId: "11",
			Name:       "KOTA SABANG",
		},
		{
			Id:         "1173",
			ProvinceId: "11",
			Name:       "KOTA LHOKSEUMAWE",
		},
		{
			Id:         "1174",
			ProvinceId: "11",
			Name:       "KOTA LANGSA",
		},
		{
			Id:         "1175",
			ProvinceId: "11",
			Name:       "KOTA SUBULUSSALAM",
		},
		{
			Id:         "1201",
			ProvinceId: "12",
			Name:       "KAB. TAPANULI TENGAH",
		},
		{
			Id:         "1202",
			ProvinceId: "12",
			Name:       "KAB. TAPANULI UTARA",
		},
		{
			Id:         "1203",
			ProvinceId: "12",
			Name:       "KAB. TAPANULI SELATAN",
		},
		{
			Id:         "1204",
			ProvinceId: "12",
			Name:       "KAB. NIAS",
		},
		{
			Id:         "1205",
			ProvinceId: "12",
			Name:       "KAB. LANGKAT",
		},
		{
			Id:         "1206",
			ProvinceId: "12",
			Name:       "KAB. KARO",
		},
		{
			Id:         "1207",
			ProvinceId: "12",
			Name:       "KAB. DELI SERDANG",
		},
		{
			Id:         "1208",
			ProvinceId: "12",
			Name:       "KAB. SIMALUNGUN",
		},
		{
			Id:         "1209",
			ProvinceId: "12",
			Name:       "KAB. ASAHAN",
		},
		{
			Id:         "1210",
			ProvinceId: "12",
			Name:       "KAB. LABUHANBATU",
		},
		{
			Id:         "1211",
			ProvinceId: "12",
			Name:       "KAB. DAIRI",
		},
		{
			Id:         "1212",
			ProvinceId: "12",
			Name:       "KAB. TOBA",
		},
		{
			Id:         "1213",
			ProvinceId: "12",
			Name:       "KAB. MANDAILING NATAL",
		},
		{
			Id:         "1214",
			ProvinceId: "12",
			Name:       "KAB. NIAS SELATAN",
		},
		{
			Id:         "1215",
			ProvinceId: "12",
			Name:       "KAB. PAKPAK BHARAT",
		},
		{
			Id:         "1216",
			ProvinceId: "12",
			Name:       "KAB. HUMBANG HASUNDUTAN",
		},
		{
			Id:         "1217",
			ProvinceId: "12",
			Name:       "KAB. SAMOSIR",
		},
		{
			Id:         "1218",
			ProvinceId: "12",
			Name:       "KAB. SERDANG BEDAGAI",
		},
		{
			Id:         "1219",
			ProvinceId: "12",
			Name:       "KAB. BATU BARA",
		},
		{
			Id:         "1220",
			ProvinceId: "12",
			Name:       "KAB. PADANG LAWAS UTARA",
		},
		{
			Id:         "1221",
			ProvinceId: "12",
			Name:       "KAB. PADANG LAWAS",
		},
		{
			Id:         "1222",
			ProvinceId: "12",
			Name:       "KAB. LABUHANBATU SELATAN",
		},
		{
			Id:         "1223",
			ProvinceId: "12",
			Name:       "KAB. LABUHANBATU UTARA",
		},
		{
			Id:         "1224",
			ProvinceId: "12",
			Name:       "KAB. NIAS UTARA",
		},
		{
			Id:         "1225",
			ProvinceId: "12",
			Name:       "KAB. NIAS BARAT",
		},
		{
			Id:         "1271",
			ProvinceId: "12",
			Name:       "KOTA MEDAN",
		},
		{
			Id:         "1272",
			ProvinceId: "12",
			Name:       "KOTA PEMATANGSIANTAR",
		},
		{
			Id:         "1273",
			ProvinceId: "12",
			Name:       "KOTA SIBOLGA",
		},
		{
			Id:         "1274",
			ProvinceId: "12",
			Name:       "KOTA TANJUNG BALAI",
		},
		{
			Id:         "1275",
			ProvinceId: "12",
			Name:       "KOTA BINJAI",
		},
		{
			Id:         "1276",
			ProvinceId: "12",
			Name:       "KOTA TEBING TINGGI",
		},
		{
			Id:         "1277",
			ProvinceId: "12",
			Name:       "KOTA PADANGSIDIMPUAN",
		},
		{
			Id:         "1278",
			ProvinceId: "12",
			Name:       "KOTA GUNUNGSITOLI",
		},
		{
			Id:         "1301",
			ProvinceId: "13",
			Name:       "KAB. PESISIR SELATAN",
		},
		{
			Id:         "1302",
			ProvinceId: "13",
			Name:       "KAB. SOLOK",
		},
		{
			Id:         "1303",
			ProvinceId: "13",
			Name:       "KAB. SIJUNJUNG",
		},
		{
			Id:         "1304",
			ProvinceId: "13",
			Name:       "KAB. TANAH DATAR",
		},
		{
			Id:         "1305",
			ProvinceId: "13",
			Name:       "KAB. PADANG PARIAMAN",
		},
		{
			Id:         "1306",
			ProvinceId: "13",
			Name:       "KAB. AGAM",
		},
		{
			Id:         "1307",
			ProvinceId: "13",
			Name:       "KAB. LIMA PULUH KOTA",
		},
		{
			Id:         "1308",
			ProvinceId: "13",
			Name:       "KAB. PASAMAN",
		},
		{
			Id:         "1309",
			ProvinceId: "13",
			Name:       "KAB. KEPULAUAN MENTAWAI",
		},
		{
			Id:         "1310",
			ProvinceId: "13",
			Name:       "KAB. DHARMASRAYA",
		},
		{
			Id:         "1311",
			ProvinceId: "13",
			Name:       "KAB. SOLOK SELATAN",
		},
		{
			Id:         "1312",
			ProvinceId: "13",
			Name:       "KAB. PASAMAN BARAT",
		},
		{
			Id:         "1371",
			ProvinceId: "13",
			Name:       "KOTA PADANG",
		},
		{
			Id:         "1372",
			ProvinceId: "13",
			Name:       "KOTA SOLOK",
		},
		{
			Id:         "1373",
			ProvinceId: "13",
			Name:       "KOTA SAWAHLUNTO",
		},
		{
			Id:         "1374",
			ProvinceId: "13",
			Name:       "KOTA PADANG PANJANG",
		},
		{
			Id:         "1375",
			ProvinceId: "13",
			Name:       "KOTA BUKITTINGGI",
		},
		{
			Id:         "1376",
			ProvinceId: "13",
			Name:       "KOTA PAYAKUMBUH",
		},
		{
			Id:         "1377",
			ProvinceId: "13",
			Name:       "KOTA PARIAMAN",
		},
		{
			Id:         "1401",
			ProvinceId: "14",
			Name:       "KAB. KAMPAR",
		},
		{
			Id:         "1402",
			ProvinceId: "14",
			Name:       "KAB. INDRAGIRI HULU",
		},
		{
			Id:         "1403",
			ProvinceId: "14",
			Name:       "KAB. BENGKALIS",
		},
		{
			Id:         "1404",
			ProvinceId: "14",
			Name:       "KAB. INDRAGIRI HILIR",
		},
		{
			Id:         "1405",
			ProvinceId: "14",
			Name:       "KAB. PELALAWAN",
		},
		{
			Id:         "1406",
			ProvinceId: "14",
			Name:       "KAB. ROKAN HULU",
		},
		{
			Id:         "1407",
			ProvinceId: "14",
			Name:       "KAB. ROKAN HILIR",
		},
		{
			Id:         "1408",
			ProvinceId: "14",
			Name:       "KAB. SIAK",
		},
		{
			Id:         "1409",
			ProvinceId: "14",
			Name:       "KAB. KUANTAN SINGINGI",
		},
		{
			Id:         "1410",
			ProvinceId: "14",
			Name:       "KAB. KEPULAUAN MERANTI",
		},
		{
			Id:         "1471",
			ProvinceId: "14",
			Name:       "KOTA PEKANBARU",
		},
		{
			Id:         "1472",
			ProvinceId: "14",
			Name:       "KOTA DUMAI",
		},
		{
			Id:         "1501",
			ProvinceId: "15",
			Name:       "KAB. KERINCI",
		},
		{
			Id:         "1502",
			ProvinceId: "15",
			Name:       "KAB. MERANGIN",
		},
		{
			Id:         "1503",
			ProvinceId: "15",
			Name:       "KAB. SAROLANGUN",
		},
		{
			Id:         "1504",
			ProvinceId: "15",
			Name:       "KAB. BATANGHARI",
		},
		{
			Id:         "1505",
			ProvinceId: "15",
			Name:       "KAB. MUARO JAMBI",
		},
		{
			Id:         "1506",
			ProvinceId: "15",
			Name:       "KAB. TANJUNG JABUNG BARAT",
		},
		{
			Id:         "1507",
			ProvinceId: "15",
			Name:       "KAB. TANJUNG JABUNG TIMUR",
		},
		{
			Id:         "1508",
			ProvinceId: "15",
			Name:       "KAB. BUNGO",
		},
		{
			Id:         "1509",
			ProvinceId: "15",
			Name:       "KAB. TEBO",
		},
		{
			Id:         "1571",
			ProvinceId: "15",
			Name:       "KOTA JAMBI",
		},
		{
			Id:         "1572",
			ProvinceId: "15",
			Name:       "KOTA SUNGAI PENUH",
		},
		{
			Id:         "1601",
			ProvinceId: "16",
			Name:       "KAB. OGAN KOMERING ULU",
		},
		{
			Id:         "1602",
			ProvinceId: "16",
			Name:       "KAB. OGAN KOMERING ILIR",
		},
		{
			Id:         "1603",
			ProvinceId: "16",
			Name:       "KAB. MUARA ENIM",
		},
		{
			Id:         "1604",
			ProvinceId: "16",
			Name:       "KAB. LAHAT",
		},
		{
			Id:         "1605",
			ProvinceId: "16",
			Name:       "KAB. MUSI RAWAS",
		},
		{
			Id:         "1606",
			ProvinceId: "16",
			Name:       "KAB. MUSI BANYUASIN",
		},
		{
			Id:         "1607",
			ProvinceId: "16",
			Name:       "KAB. BANYUASIN",
		},
		{
			Id:         "1608",
			ProvinceId: "16",
			Name:       "KAB. OGAN KOMERING ULU TIMUR",
		},
		{
			Id:         "1609",
			ProvinceId: "16",
			Name:       "KAB. OGAN KOMERING ULU SELATAN",
		},
		{
			Id:         "1610",
			ProvinceId: "16",
			Name:       "KAB. OGAN ILIR",
		},
		{
			Id:         "1611",
			ProvinceId: "16",
			Name:       "KAB. EMPAT LAWANG",
		},
		{
			Id:         "1612",
			ProvinceId: "16",
			Name:       "KAB. PENUKAL ABAB LEMATANG ILIR",
		},
		{
			Id:         "1613",
			ProvinceId: "16",
			Name:       "KAB. MUSI RAWAS UTARA",
		},
		{
			Id:         "1671",
			ProvinceId: "16",
			Name:       "KOTA PALEMBANG",
		},
		{
			Id:         "1672",
			ProvinceId: "16",
			Name:       "KOTA PAGAR ALAM",
		},
		{
			Id:         "1673",
			ProvinceId: "16",
			Name:       "KOTA LUBUK LINGGAU",
		},
		{
			Id:         "1674",
			ProvinceId: "16",
			Name:       "KOTA PRABUMULIH",
		},
		{
			Id:         "1701",
			ProvinceId: "17",
			Name:       "KAB. BENGKULU SELATAN",
		},
		{
			Id:         "1702",
			ProvinceId: "17",
			Name:       "KAB. REJANG LEBONG",
		},
		{
			Id:         "1703",
			ProvinceId: "17",
			Name:       "KAB. BENGKULU UTARA",
		},
		{
			Id:         "1704",
			ProvinceId: "17",
			Name:       "KAB. KAUR",
		},
		{
			Id:         "1705",
			ProvinceId: "17",
			Name:       "KAB. SELUMA",
		},
		{
			Id:         "1706",
			ProvinceId: "17",
			Name:       "KAB. MUKO MUKO",
		},
		{
			Id:         "1707",
			ProvinceId: "17",
			Name:       "KAB. LEBONG",
		},
		{
			Id:         "1708",
			ProvinceId: "17",
			Name:       "KAB. KEPAHIANG",
		},
		{
			Id:         "1709",
			ProvinceId: "17",
			Name:       "KAB. BENGKULU TENGAH",
		},
		{
			Id:         "1771",
			ProvinceId: "17",
			Name:       "KOTA BENGKULU",
		},
		{
			Id:         "1801",
			ProvinceId: "18",
			Name:       "KAB. LAMPUNG SELATAN",
		},
		{
			Id:         "1802",
			ProvinceId: "18",
			Name:       "KAB. LAMPUNG TENGAH",
		},
		{
			Id:         "1803",
			ProvinceId: "18",
			Name:       "KAB. LAMPUNG UTARA",
		},
		{
			Id:         "1804",
			ProvinceId: "18",
			Name:       "KAB. LAMPUNG BARAT",
		},
		{
			Id:         "1805",
			ProvinceId: "18",
			Name:       "KAB. TULANG BAWANG",
		},
		{
			Id:         "1806",
			ProvinceId: "18",
			Name:       "KAB. TANGGAMUS",
		},
		{
			Id:         "1807",
			ProvinceId: "18",
			Name:       "KAB. LAMPUNG TIMUR",
		},
		{
			Id:         "1808",
			ProvinceId: "18",
			Name:       "KAB. WAY KANAN",
		},
		{
			Id:         "1809",
			ProvinceId: "18",
			Name:       "KAB. PESAWARAN",
		},
		{
			Id:         "1810",
			ProvinceId: "18",
			Name:       "KAB. PRINGSEWU",
		},
		{
			Id:         "1811",
			ProvinceId: "18",
			Name:       "KAB. MESUJI",
		},
		{
			Id:         "1812",
			ProvinceId: "18",
			Name:       "KAB. TULANG BAWANG BARAT",
		},
		{
			Id:         "1813",
			ProvinceId: "18",
			Name:       "KAB. PESISIR BARAT",
		},
		{
			Id:         "1871",
			ProvinceId: "18",
			Name:       "KOTA BANDAR LAMPUNG",
		},
		{
			Id:         "1872",
			ProvinceId: "18",
			Name:       "KOTA METRO",
		},
		{
			Id:         "1901",
			ProvinceId: "19",
			Name:       "KAB. BANGKA",
		},
		{
			Id:         "1902",
			ProvinceId: "19",
			Name:       "KAB. BELITUNG",
		},
		{
			Id:         "1903",
			ProvinceId: "19",
			Name:       "KAB. BANGKA SELATAN",
		},
		{
			Id:         "1904",
			ProvinceId: "19",
			Name:       "KAB. BANGKA TENGAH",
		},
		{
			Id:         "1905",
			ProvinceId: "19",
			Name:       "KAB. BANGKA BARAT",
		},
		{
			Id:         "1906",
			ProvinceId: "19",
			Name:       "KAB. BELITUNG TIMUR",
		},
		{
			Id:         "1971",
			ProvinceId: "19",
			Name:       "KOTA PANGKAL PINANG",
		},
		{
			Id:         "2101",
			ProvinceId: "21",
			Name:       "KAB. BINTAN",
		},
		{
			Id:         "2102",
			ProvinceId: "21",
			Name:       "KAB. KARIMUN",
		},
		{
			Id:         "2103",
			ProvinceId: "21",
			Name:       "KAB. NATUNA",
		},
		{
			Id:         "2104",
			ProvinceId: "21",
			Name:       "KAB. LINGGA",
		},
		{
			Id:         "2105",
			ProvinceId: "21",
			Name:       "KAB. KEPULAUAN ANAMBAS",
		},
		{
			Id:         "2171",
			ProvinceId: "21",
			Name:       "KOTA BATAM",
		},
		{
			Id:         "2172",
			ProvinceId: "21",
			Name:       "KOTA TANJUNG PINANG",
		},
		{
			Id:         "3101",
			ProvinceId: "31",
			Name:       "KAB. ADM. KEP. SERIBU",
		},
		{
			Id:         "3171",
			ProvinceId: "31",
			Name:       "KOTA ADM. JAKARTA PUSAT",
		},
		{
			Id:         "3172",
			ProvinceId: "31",
			Name:       "KOTA ADM. JAKARTA UTARA",
		},
		{
			Id:         "3173",
			ProvinceId: "31",
			Name:       "KOTA ADM. JAKARTA BARAT",
		},
		{
			Id:         "3174",
			ProvinceId: "31",
			Name:       "KOTA ADM. JAKARTA SELATAN",
		},
		{
			Id:         "3175",
			ProvinceId: "31",
			Name:       "KOTA ADM. JAKARTA TIMUR",
		},
		{
			Id:         "3201",
			ProvinceId: "32",
			Name:       "KAB. BOGOR",
		},
		{
			Id:         "3202",
			ProvinceId: "32",
			Name:       "KAB. SUKABUMI",
		},
		{
			Id:         "3203",
			ProvinceId: "32",
			Name:       "KAB. CIANJUR",
		},
		{
			Id:         "3204",
			ProvinceId: "32",
			Name:       "KAB. BANDUNG",
		},
		{
			Id:         "3205",
			ProvinceId: "32",
			Name:       "KAB. GARUT",
		},
		{
			Id:         "3206",
			ProvinceId: "32",
			Name:       "KAB. TASIKMALAYA",
		},
		{
			Id:         "3207",
			ProvinceId: "32",
			Name:       "KAB. CIAMIS",
		},
		{
			Id:         "3208",
			ProvinceId: "32",
			Name:       "KAB. KUNINGAN",
		},
		{
			Id:         "3209",
			ProvinceId: "32",
			Name:       "KAB. CIREBON",
		},
		{
			Id:         "3210",
			ProvinceId: "32",
			Name:       "KAB. MAJALENGKA",
		},
		{
			Id:         "3211",
			ProvinceId: "32",
			Name:       "KAB. SUMEDANG",
		},
		{
			Id:         "3212",
			ProvinceId: "32",
			Name:       "KAB. INDRAMAYU",
		},
		{
			Id:         "3213",
			ProvinceId: "32",
			Name:       "KAB. SUBANG",
		},
		{
			Id:         "3214",
			ProvinceId: "32",
			Name:       "KAB. PURWAKARTA",
		},
		{
			Id:         "3215",
			ProvinceId: "32",
			Name:       "KAB. KARAWANG",
		},
		{
			Id:         "3216",
			ProvinceId: "32",
			Name:       "KAB. BEKASI",
		},
		{
			Id:         "3217",
			ProvinceId: "32",
			Name:       "KAB. BANDUNG BARAT",
		},
		{
			Id:         "3218",
			ProvinceId: "32",
			Name:       "KAB. PANGANDARAN",
		},
		{
			Id:         "3271",
			ProvinceId: "32",
			Name:       "KOTA BOGOR",
		},
		{
			Id:         "3272",
			ProvinceId: "32",
			Name:       "KOTA SUKABUMI",
		},
		{
			Id:         "3273",
			ProvinceId: "32",
			Name:       "KOTA BANDUNG",
		},
		{
			Id:         "3274",
			ProvinceId: "32",
			Name:       "KOTA CIREBON",
		},
		{
			Id:         "3275",
			ProvinceId: "32",
			Name:       "KOTA BEKASI",
		},
		{
			Id:         "3276",
			ProvinceId: "32",
			Name:       "KOTA DEPOK",
		},
		{
			Id:         "3277",
			ProvinceId: "32",
			Name:       "KOTA CIMAHI",
		},
		{
			Id:         "3278",
			ProvinceId: "32",
			Name:       "KOTA TASIKMALAYA",
		},
		{
			Id:         "3279",
			ProvinceId: "32",
			Name:       "KOTA BANJAR",
		},
		{
			Id:         "3301",
			ProvinceId: "33",
			Name:       "KAB. CILACAP",
		},
		{
			Id:         "3302",
			ProvinceId: "33",
			Name:       "KAB. BANYUMAS",
		},
		{
			Id:         "3303",
			ProvinceId: "33",
			Name:       "KAB. PURBALINGGA",
		},
		{
			Id:         "3304",
			ProvinceId: "33",
			Name:       "KAB. BANJARNEGARA",
		},
		{
			Id:         "3305",
			ProvinceId: "33",
			Name:       "KAB. KEBUMEN",
		},
		{
			Id:         "3306",
			ProvinceId: "33",
			Name:       "KAB. PURWOREJO",
		},
		{
			Id:         "3307",
			ProvinceId: "33",
			Name:       "KAB. WONOSOBO",
		},
		{
			Id:         "3308",
			ProvinceId: "33",
			Name:       "KAB. MAGELANG",
		},
		{
			Id:         "3309",
			ProvinceId: "33",
			Name:       "KAB. BOYOLALI",
		},
		{
			Id:         "3310",
			ProvinceId: "33",
			Name:       "KAB. KLATEN",
		},
		{
			Id:         "3311",
			ProvinceId: "33",
			Name:       "KAB. SUKOHARJO",
		},
		{
			Id:         "3312",
			ProvinceId: "33",
			Name:       "KAB. WONOGIRI",
		},
		{
			Id:         "3313",
			ProvinceId: "33",
			Name:       "KAB. KARANGANYAR",
		},
		{
			Id:         "3314",
			ProvinceId: "33",
			Name:       "KAB. SRAGEN",
		},
		{
			Id:         "3315",
			ProvinceId: "33",
			Name:       "KAB. GROBOGAN",
		},
		{
			Id:         "3316",
			ProvinceId: "33",
			Name:       "KAB. BLORA",
		},
		{
			Id:         "3317",
			ProvinceId: "33",
			Name:       "KAB. REMBANG",
		},
		{
			Id:         "3318",
			ProvinceId: "33",
			Name:       "KAB. PATI",
		},
		{
			Id:         "3319",
			ProvinceId: "33",
			Name:       "KAB. KUDUS",
		},
		{
			Id:         "3320",
			ProvinceId: "33",
			Name:       "KAB. JEPARA",
		},
		{
			Id:         "3321",
			ProvinceId: "33",
			Name:       "KAB. DEMAK",
		},
		{
			Id:         "3322",
			ProvinceId: "33",
			Name:       "KAB. SEMARANG",
		},
		{
			Id:         "3323",
			ProvinceId: "33",
			Name:       "KAB. TEMANGGUNG",
		},
		{
			Id:         "3324",
			ProvinceId: "33",
			Name:       "KAB. KENDAL",
		},
		{
			Id:         "3325",
			ProvinceId: "33",
			Name:       "KAB. BATANG",
		},
		{
			Id:         "3326",
			ProvinceId: "33",
			Name:       "KAB. PEKALONGAN",
		},
		{
			Id:         "3327",
			ProvinceId: "33",
			Name:       "KAB. PEMALANG",
		},
		{
			Id:         "3328",
			ProvinceId: "33",
			Name:       "KAB. TEGAL",
		},
		{
			Id:         "3329",
			ProvinceId: "33",
			Name:       "KAB. BREBES",
		},
		{
			Id:         "3371",
			ProvinceId: "33",
			Name:       "KOTA MAGELANG",
		},
		{
			Id:         "3372",
			ProvinceId: "33",
			Name:       "KOTA SURAKARTA",
		},
		{
			Id:         "3373",
			ProvinceId: "33",
			Name:       "KOTA SALATIGA",
		},
		{
			Id:         "3374",
			ProvinceId: "33",
			Name:       "KOTA SEMARANG",
		},
		{
			Id:         "3375",
			ProvinceId: "33",
			Name:       "KOTA PEKALONGAN",
		},
		{
			Id:         "3376",
			ProvinceId: "33",
			Name:       "KOTA TEGAL",
		},
		{
			Id:         "3401",
			ProvinceId: "34",
			Name:       "KAB. KULON PROGO",
		},
		{
			Id:         "3402",
			ProvinceId: "34",
			Name:       "KAB. BANTUL",
		},
		{
			Id:         "3403",
			ProvinceId: "34",
			Name:       "KAB. GUNUNGKIDUL",
		},
		{
			Id:         "3404",
			ProvinceId: "34",
			Name:       "KAB. SLEMAN",
		},
		{
			Id:         "3471",
			ProvinceId: "34",
			Name:       "KOTA YOGYAKARTA",
		},
		{
			Id:         "3501",
			ProvinceId: "35",
			Name:       "KAB. PACITAN",
		},
		{
			Id:         "3502",
			ProvinceId: "35",
			Name:       "KAB. PONOROGO",
		},
		{
			Id:         "3503",
			ProvinceId: "35",
			Name:       "KAB. TRENGGALEK",
		},
		{
			Id:         "3504",
			ProvinceId: "35",
			Name:       "KAB. TULUNGAGUNG",
		},
		{
			Id:         "3505",
			ProvinceId: "35",
			Name:       "KAB. BLITAR",
		},
		{
			Id:         "3506",
			ProvinceId: "35",
			Name:       "KAB. KEDIRI",
		},
		{
			Id:         "3507",
			ProvinceId: "35",
			Name:       "KAB. MALANG",
		},
		{
			Id:         "3508",
			ProvinceId: "35",
			Name:       "KAB. LUMAJANG",
		},
		{
			Id:         "3509",
			ProvinceId: "35",
			Name:       "KAB. JEMBER",
		},
		{
			Id:         "3510",
			ProvinceId: "35",
			Name:       "KAB. BANYUWANGI",
		},
		{
			Id:         "3511",
			ProvinceId: "35",
			Name:       "KAB. BONDOWOSO",
		},
		{
			Id:         "3512",
			ProvinceId: "35",
			Name:       "KAB. SITUBONDO",
		},
		{
			Id:         "3513",
			ProvinceId: "35",
			Name:       "KAB. PROBOLINGGO",
		},
		{
			Id:         "3514",
			ProvinceId: "35",
			Name:       "KAB. PASURUAN",
		},
		{
			Id:         "3515",
			ProvinceId: "35",
			Name:       "KAB. SIDOARJO",
		},
		{
			Id:         "3516",
			ProvinceId: "35",
			Name:       "KAB. MOJOKERTO",
		},
		{
			Id:         "3517",
			ProvinceId: "35",
			Name:       "KAB. JOMBANG",
		},
		{
			Id:         "3518",
			ProvinceId: "35",
			Name:       "KAB. NGANJUK",
		},
		{
			Id:         "3519",
			ProvinceId: "35",
			Name:       "KAB. MADIUN",
		},
		{
			Id:         "3520",
			ProvinceId: "35",
			Name:       "KAB. MAGETAN",
		},
		{
			Id:         "3521",
			ProvinceId: "35",
			Name:       "KAB. NGAWI",
		},
		{
			Id:         "3522",
			ProvinceId: "35",
			Name:       "KAB. BOJONEGORO",
		},
		{
			Id:         "3523",
			ProvinceId: "35",
			Name:       "KAB. TUBAN",
		},
		{
			Id:         "3524",
			ProvinceId: "35",
			Name:       "KAB. LAMONGAN",
		},
		{
			Id:         "3525",
			ProvinceId: "35",
			Name:       "KAB. GRESIK",
		},
		{
			Id:         "3526",
			ProvinceId: "35",
			Name:       "KAB. BANGKALAN",
		},
		{
			Id:         "3527",
			ProvinceId: "35",
			Name:       "KAB. SAMPANG",
		},
		{
			Id:         "3528",
			ProvinceId: "35",
			Name:       "KAB. PAMEKASAN",
		},
		{
			Id:         "3529",
			ProvinceId: "35",
			Name:       "KAB. SUMENEP",
		},
		{
			Id:         "3571",
			ProvinceId: "35",
			Name:       "KOTA KEDIRI",
		},
		{
			Id:         "3572",
			ProvinceId: "35",
			Name:       "KOTA BLITAR",
		},
		{
			Id:         "3573",
			ProvinceId: "35",
			Name:       "KOTA MALANG",
		},
		{
			Id:         "3574",
			ProvinceId: "35",
			Name:       "KOTA PROBOLINGGO",
		},
		{
			Id:         "3575",
			ProvinceId: "35",
			Name:       "KOTA PASURUAN",
		},
		{
			Id:         "3576",
			ProvinceId: "35",
			Name:       "KOTA MOJOKERTO",
		},
		{
			Id:         "3577",
			ProvinceId: "35",
			Name:       "KOTA MADIUN",
		},
		{
			Id:         "3578",
			ProvinceId: "35",
			Name:       "KOTA SURABAYA",
		},
		{
			Id:         "3579",
			ProvinceId: "35",
			Name:       "KOTA BATU",
		},
		{
			Id:         "3601",
			ProvinceId: "36",
			Name:       "KAB. PANDEGLANG",
		},
		{
			Id:         "3602",
			ProvinceId: "36",
			Name:       "KAB. LEBAK",
		},
		{
			Id:         "3603",
			ProvinceId: "36",
			Name:       "KAB. TANGERANG",
		},
		{
			Id:         "3604",
			ProvinceId: "36",
			Name:       "KAB. SERANG",
		},
		{
			Id:         "3671",
			ProvinceId: "36",
			Name:       "KOTA TANGERANG",
		},
		{
			Id:         "3672",
			ProvinceId: "36",
			Name:       "KOTA CILEGON",
		},
		{
			Id:         "3673",
			ProvinceId: "36",
			Name:       "KOTA SERANG",
		},
		{
			Id:         "3674",
			ProvinceId: "36",
			Name:       "KOTA TANGERANG SELATAN",
		},
		{
			Id:         "5101",
			ProvinceId: "51",
			Name:       "KAB. JEMBRANA",
		},
		{
			Id:         "5102",
			ProvinceId: "51",
			Name:       "KAB. TABANAN",
		},
		{
			Id:         "5103",
			ProvinceId: "51",
			Name:       "KAB. BADUNG",
		},
		{
			Id:         "5104",
			ProvinceId: "51",
			Name:       "KAB. GIANYAR",
		},
		{
			Id:         "5105",
			ProvinceId: "51",
			Name:       "KAB. KLUNGKUNG",
		},
		{
			Id:         "5106",
			ProvinceId: "51",
			Name:       "KAB. BANGLI",
		},
		{
			Id:         "5107",
			ProvinceId: "51",
			Name:       "KAB. KARANGASEM",
		},
		{
			Id:         "5108",
			ProvinceId: "51",
			Name:       "KAB. BULELENG",
		},
		{
			Id:         "5171",
			ProvinceId: "51",
			Name:       "KOTA DENPASAR",
		},
		{
			Id:         "5201",
			ProvinceId: "52",
			Name:       "KAB. LOMBOK BARAT",
		},
		{
			Id:         "5202",
			ProvinceId: "52",
			Name:       "KAB. LOMBOK TENGAH",
		},
		{
			Id:         "5203",
			ProvinceId: "52",
			Name:       "KAB. LOMBOK TIMUR",
		},
		{
			Id:         "5204",
			ProvinceId: "52",
			Name:       "KAB. SUMBAWA",
		},
		{
			Id:         "5205",
			ProvinceId: "52",
			Name:       "KAB. DOMPU",
		},
		{
			Id:         "5206",
			ProvinceId: "52",
			Name:       "KAB. BIMA",
		},
		{
			Id:         "5207",
			ProvinceId: "52",
			Name:       "KAB. SUMBAWA BARAT",
		},
		{
			Id:         "5208",
			ProvinceId: "52",
			Name:       "KAB. LOMBOK UTARA",
		},
		{
			Id:         "5271",
			ProvinceId: "52",
			Name:       "KOTA MATARAM",
		},
		{
			Id:         "5272",
			ProvinceId: "52",
			Name:       "KOTA BIMA",
		},
		{
			Id:         "5301",
			ProvinceId: "53",
			Name:       "KAB. KUPANG",
		},
		{
			Id:         "5302",
			ProvinceId: "53",
			Name:       "KAB TIMOR TENGAH SELATAN",
		},
		{
			Id:         "5303",
			ProvinceId: "53",
			Name:       "KAB. TIMOR TENGAH UTARA",
		},
		{
			Id:         "5304",
			ProvinceId: "53",
			Name:       "KAB. BELU",
		},
		{
			Id:         "5305",
			ProvinceId: "53",
			Name:       "KAB. ALOR",
		},
		{
			Id:         "5306",
			ProvinceId: "53",
			Name:       "KAB. FLORES TIMUR",
		},
		{
			Id:         "5307",
			ProvinceId: "53",
			Name:       "KAB. SIKKA",
		},
		{
			Id:         "5308",
			ProvinceId: "53",
			Name:       "KAB. ENDE",
		},
		{
			Id:         "5309",
			ProvinceId: "53",
			Name:       "KAB. NGADA",
		},
		{
			Id:         "5310",
			ProvinceId: "53",
			Name:       "KAB. MANGGARAI",
		},
		{
			Id:         "5311",
			ProvinceId: "53",
			Name:       "KAB. SUMBA TIMUR",
		},
		{
			Id:         "5312",
			ProvinceId: "53",
			Name:       "KAB. SUMBA BARAT",
		},
		{
			Id:         "5313",
			ProvinceId: "53",
			Name:       "KAB. LEMBATA",
		},
		{
			Id:         "5314",
			ProvinceId: "53",
			Name:       "KAB. ROTE NDAO",
		},
		{
			Id:         "5315",
			ProvinceId: "53",
			Name:       "KAB. MANGGARAI BARAT",
		},
		{
			Id:         "5316",
			ProvinceId: "53",
			Name:       "KAB. NAGEKEO",
		},
		{
			Id:         "5317",
			ProvinceId: "53",
			Name:       "KAB. SUMBA TENGAH",
		},
		{
			Id:         "5318",
			ProvinceId: "53",
			Name:       "KAB. SUMBA BARAT DAYA",
		},
		{
			Id:         "5319",
			ProvinceId: "53",
			Name:       "KAB. MANGGARAI TIMUR",
		},
		{
			Id:         "5320",
			ProvinceId: "53",
			Name:       "KAB. SABU RAIJUA",
		},
		{
			Id:         "5321",
			ProvinceId: "53",
			Name:       "KAB. MALAKA",
		},
		{
			Id:         "5371",
			ProvinceId: "53",
			Name:       "KOTA KUPANG",
		},
		{
			Id:         "6101",
			ProvinceId: "61",
			Name:       "KAB. SAMBAS",
		},
		{
			Id:         "6102",
			ProvinceId: "61",
			Name:       "KAB. MEMPAWAH",
		},
		{
			Id:         "6103",
			ProvinceId: "61",
			Name:       "KAB. SANGGAU",
		},
		{
			Id:         "6104",
			ProvinceId: "61",
			Name:       "KAB. KETAPANG",
		},
		{
			Id:         "6105",
			ProvinceId: "61",
			Name:       "KAB. SINTANG",
		},
		{
			Id:         "6106",
			ProvinceId: "61",
			Name:       "KAB. KAPUAS HULU",
		},
		{
			Id:         "6107",
			ProvinceId: "61",
			Name:       "KAB. BENGKAYANG",
		},
		{
			Id:         "6108",
			ProvinceId: "61",
			Name:       "KAB. LANDAK",
		},
		{
			Id:         "6109",
			ProvinceId: "61",
			Name:       "KAB. SEKADAU",
		},
		{
			Id:         "6110",
			ProvinceId: "61",
			Name:       "KAB. MELAWI",
		},
		{
			Id:         "6111",
			ProvinceId: "61",
			Name:       "KAB. KAYONG UTARA",
		},
		{
			Id:         "6112",
			ProvinceId: "61",
			Name:       "KAB. KUBU RAYA",
		},
		{
			Id:         "6171",
			ProvinceId: "61",
			Name:       "KOTA PONTIANAK",
		},
		{
			Id:         "6172",
			ProvinceId: "61",
			Name:       "KOTA SINGKAWANG",
		},
		{
			Id:         "6201",
			ProvinceId: "62",
			Name:       "KAB. KOTAWARINGIN BARAT",
		},
		{
			Id:         "6202",
			ProvinceId: "62",
			Name:       "KAB. KOTAWARINGIN TIMUR",
		},
		{
			Id:         "6203",
			ProvinceId: "62",
			Name:       "KAB. KAPUAS",
		},
		{
			Id:         "6204",
			ProvinceId: "62",
			Name:       "KAB. BARITO SELATAN",
		},
		{
			Id:         "6205",
			ProvinceId: "62",
			Name:       "KAB. BARITO UTARA",
		},
		{
			Id:         "6206",
			ProvinceId: "62",
			Name:       "KAB. KATINGAN",
		},
		{
			Id:         "6207",
			ProvinceId: "62",
			Name:       "KAB. SERUYAN",
		},
		{
			Id:         "6208",
			ProvinceId: "62",
			Name:       "KAB. SUKAMARA",
		},
		{
			Id:         "6209",
			ProvinceId: "62",
			Name:       "KAB. LAMANDAU",
		},
		{
			Id:         "6210",
			ProvinceId: "62",
			Name:       "KAB. GUNUNG MAS",
		},
		{
			Id:         "6211",
			ProvinceId: "62",
			Name:       "KAB. PULANG PISAU",
		},
		{
			Id:         "6212",
			ProvinceId: "62",
			Name:       "KAB. MURUNG RAYA",
		},
		{
			Id:         "6213",
			ProvinceId: "62",
			Name:       "KAB. BARITO TIMUR",
		},
		{
			Id:         "6271",
			ProvinceId: "62",
			Name:       "KOTA PALANGKARAYA",
		},
		{
			Id:         "6301",
			ProvinceId: "63",
			Name:       "KAB. TANAH LAUT",
		},
		{
			Id:         "6302",
			ProvinceId: "63",
			Name:       "KAB. KOTABARU",
		},
		{
			Id:         "6303",
			ProvinceId: "63",
			Name:       "KAB. BANJAR",
		},
		{
			Id:         "6304",
			ProvinceId: "63",
			Name:       "KAB. BARITO KUALA",
		},
		{
			Id:         "6305",
			ProvinceId: "63",
			Name:       "KAB. TAPIN",
		},
		{
			Id:         "6306",
			ProvinceId: "63",
			Name:       "KAB. HULU SUNGAI SELATAN",
		},
		{
			Id:         "6307",
			ProvinceId: "63",
			Name:       "KAB. HULU SUNGAI TENGAH",
		},
		{
			Id:         "6308",
			ProvinceId: "63",
			Name:       "KAB. HULU SUNGAI UTARA",
		},
		{
			Id:         "6309",
			ProvinceId: "63",
			Name:       "KAB. TABALONG",
		},
		{
			Id:         "6310",
			ProvinceId: "63",
			Name:       "KAB. TANAH BUMBU",
		},
		{
			Id:         "6311",
			ProvinceId: "63",
			Name:       "KAB. BALANGAN",
		},
		{
			Id:         "6371",
			ProvinceId: "63",
			Name:       "KOTA BANJARMASIN",
		},
		{
			Id:         "6372",
			ProvinceId: "63",
			Name:       "KOTA BANJARBARU",
		},
		{
			Id:         "6401",
			ProvinceId: "64",
			Name:       "KAB. PASER",
		},
		{
			Id:         "6402",
			ProvinceId: "64",
			Name:       "KAB. KUTAI KARTANEGARA",
		},
		{
			Id:         "6403",
			ProvinceId: "64",
			Name:       "KAB. BERAU",
		},
		{
			Id:         "6407",
			ProvinceId: "64",
			Name:       "KAB. KUTAI BARAT",
		},
		{
			Id:         "6408",
			ProvinceId: "64",
			Name:       "KAB. KUTAI TIMUR",
		},
		{
			Id:         "6409",
			ProvinceId: "64",
			Name:       "KAB. PENAJAM PASER UTARA",
		},
		{
			Id:         "6411",
			ProvinceId: "64",
			Name:       "KAB. MAHAKAM ULU",
		},
		{
			Id:         "6471",
			ProvinceId: "64",
			Name:       "KOTA BALIKPAPAN",
		},
		{
			Id:         "6472",
			ProvinceId: "64",
			Name:       "KOTA SAMARINDA",
		},
		{
			Id:         "6474",
			ProvinceId: "64",
			Name:       "KOTA BONTANG",
		},
		{
			Id:         "6501",
			ProvinceId: "65",
			Name:       "KAB. BULUNGAN",
		},
		{
			Id:         "6502",
			ProvinceId: "65",
			Name:       "KAB. MALINAU",
		},
		{
			Id:         "6503",
			ProvinceId: "65",
			Name:       "KAB. NUNUKAN",
		},
		{
			Id:         "6504",
			ProvinceId: "65",
			Name:       "KAB. TANA TIDUNG",
		},
		{
			Id:         "6571",
			ProvinceId: "65",
			Name:       "KOTA TARAKAN",
		},
		{
			Id:         "7101",
			ProvinceId: "71",
			Name:       "KAB. BOLAANG MONGONDOW",
		},
		{
			Id:         "7102",
			ProvinceId: "71",
			Name:       "KAB. MINAHASA",
		},
		{
			Id:         "7103",
			ProvinceId: "71",
			Name:       "KAB. KEPULAUAN SANGIHE",
		},
		{
			Id:         "7104",
			ProvinceId: "71",
			Name:       "KAB. KEPULAUAN TALAUD",
		},
		{
			Id:         "7105",
			ProvinceId: "71",
			Name:       "KAB. MINAHASA SELATAN",
		},
		{
			Id:         "7106",
			ProvinceId: "71",
			Name:       "KAB. MINAHASA UTARA",
		},
		{
			Id:         "7107",
			ProvinceId: "71",
			Name:       "KAB. MINAHASA TENGGARA",
		},
		{
			Id:         "7108",
			ProvinceId: "71",
			Name:       "KAB. BOLAANG MONGONDOW UTARA",
		},
		{
			Id:         "7109",
			ProvinceId: "71",
			Name:       "KAB. KEP. SIAU TAGULANDANG BIARO",
		},
		{
			Id:         "7110",
			ProvinceId: "71",
			Name:       "KAB. BOLAANG MONGONDOW TIMUR",
		},
		{
			Id:         "7111",
			ProvinceId: "71",
			Name:       "KAB. BOLAANG MONGONDOW SELATAN",
		},
		{
			Id:         "7171",
			ProvinceId: "71",
			Name:       "KOTA MANADO",
		},
		{
			Id:         "7172",
			ProvinceId: "71",
			Name:       "KOTA BITUNG",
		},
		{
			Id:         "7173",
			ProvinceId: "71",
			Name:       "KOTA TOMOHON",
		},
		{
			Id:         "7174",
			ProvinceId: "71",
			Name:       "KOTA KOTAMOBAGU",
		},
		{
			Id:         "7201",
			ProvinceId: "72",
			Name:       "KAB. BANGGAI",
		},
		{
			Id:         "7202",
			ProvinceId: "72",
			Name:       "KAB. POSO",
		},
		{
			Id:         "7203",
			ProvinceId: "72",
			Name:       "KAB. DONGGALA",
		},
		{
			Id:         "7204",
			ProvinceId: "72",
			Name:       "KAB. TOLI TOLI",
		},
		{
			Id:         "7205",
			ProvinceId: "72",
			Name:       "KAB. BUOL",
		},
		{
			Id:         "7206",
			ProvinceId: "72",
			Name:       "KAB. MOROWALI",
		},
		{
			Id:         "7207",
			ProvinceId: "72",
			Name:       "KAB. BANGGAI KEPULAUAN",
		},
		{
			Id:         "7208",
			ProvinceId: "72",
			Name:       "KAB. PARIGI MOUTONG",
		},
		{
			Id:         "7209",
			ProvinceId: "72",
			Name:       "KAB. TOJO UNA UNA",
		},
		{
			Id:         "7210",
			ProvinceId: "72",
			Name:       "KAB. SIGI",
		},
		{
			Id:         "7211",
			ProvinceId: "72",
			Name:       "KAB. BANGGAI LAUT",
		},
		{
			Id:         "7212",
			ProvinceId: "72",
			Name:       "KAB. MOROWALI UTARA",
		},
		{
			Id:         "7271",
			ProvinceId: "72",
			Name:       "KOTA PALU",
		},
		{
			Id:         "7301",
			ProvinceId: "73",
			Name:       "KAB. KEPULAUAN SELAYAR",
		},
		{
			Id:         "7302",
			ProvinceId: "73",
			Name:       "KAB. BULUKUMBA",
		},
		{
			Id:         "7303",
			ProvinceId: "73",
			Name:       "KAB. BANTAENG",
		},
		{
			Id:         "7304",
			ProvinceId: "73",
			Name:       "KAB. JENEPONTO",
		},
		{
			Id:         "7305",
			ProvinceId: "73",
			Name:       "KAB. TAKALAR",
		},
		{
			Id:         "7306",
			ProvinceId: "73",
			Name:       "KAB. GOWA",
		},
		{
			Id:         "7307",
			ProvinceId: "73",
			Name:       "KAB. SINJAI",
		},
		{
			Id:         "7308",
			ProvinceId: "73",
			Name:       "KAB. BONE",
		},
		{
			Id:         "7309",
			ProvinceId: "73",
			Name:       "KAB. MAROS",
		},
		{
			Id:         "7310",
			ProvinceId: "73",
			Name:       "KAB. PANGKAJENE KEPULAUAN",
		},
		{
			Id:         "7311",
			ProvinceId: "73",
			Name:       "KAB. BARRU",
		},
		{
			Id:         "7312",
			ProvinceId: "73",
			Name:       "KAB. SOPPENG",
		},
		{
			Id:         "7313",
			ProvinceId: "73",
			Name:       "KAB. WAJO",
		},
		{
			Id:         "7314",
			ProvinceId: "73",
			Name:       "KAB. SIDENRENG RAPPANG",
		},
		{
			Id:         "7315",
			ProvinceId: "73",
			Name:       "KAB. PINRANG",
		},
		{
			Id:         "7316",
			ProvinceId: "73",
			Name:       "KAB. ENREKANG",
		},
		{
			Id:         "7317",
			ProvinceId: "73",
			Name:       "KAB. LUWU",
		},
		{
			Id:         "7318",
			ProvinceId: "73",
			Name:       "KAB. TANA TORAJA",
		},
		{
			Id:         "7322",
			ProvinceId: "73",
			Name:       "KAB. LUWU UTARA",
		},
		{
			Id:         "7324",
			ProvinceId: "73",
			Name:       "KAB. LUWU TIMUR",
		},
		{
			Id:         "7326",
			ProvinceId: "73",
			Name:       "KAB. TORAJA UTARA",
		},
		{
			Id:         "7371",
			ProvinceId: "73",
			Name:       "KOTA MAKASSAR",
		},
		{
			Id:         "7372",
			ProvinceId: "73",
			Name:       "KOTA PARE PARE",
		},
		{
			Id:         "7373",
			ProvinceId: "73",
			Name:       "KOTA PALOPO",
		},
		{
			Id:         "7401",
			ProvinceId: "74",
			Name:       "KAB. KOLAKA",
		},
		{
			Id:         "7402",
			ProvinceId: "74",
			Name:       "KAB. KONAWE",
		},
		{
			Id:         "7403",
			ProvinceId: "74",
			Name:       "KAB. MUNA",
		},
		{
			Id:         "7404",
			ProvinceId: "74",
			Name:       "KAB. BUTON",
		},
		{
			Id:         "7405",
			ProvinceId: "74",
			Name:       "KAB. KONAWE SELATAN",
		},
		{
			Id:         "7406",
			ProvinceId: "74",
			Name:       "KAB. BOMBANA",
		},
		{
			Id:         "7407",
			ProvinceId: "74",
			Name:       "KAB. WAKATOBI",
		},
		{
			Id:         "7408",
			ProvinceId: "74",
			Name:       "KAB. KOLAKA UTARA",
		},
		{
			Id:         "7409",
			ProvinceId: "74",
			Name:       "KAB. KONAWE UTARA",
		},
		{
			Id:         "7410",
			ProvinceId: "74",
			Name:       "KAB. BUTON UTARA",
		},
		{
			Id:         "7411",
			ProvinceId: "74",
			Name:       "KAB. KOLAKA TIMUR",
		},
		{
			Id:         "7412",
			ProvinceId: "74",
			Name:       "KAB. KONAWE KEPULAUAN",
		},
		{
			Id:         "7413",
			ProvinceId: "74",
			Name:       "KAB. MUNA BARAT",
		},
		{
			Id:         "7414",
			ProvinceId: "74",
			Name:       "KAB. BUTON TENGAH",
		},
		{
			Id:         "7415",
			ProvinceId: "74",
			Name:       "KAB. BUTON SELATAN",
		},
		{
			Id:         "7471",
			ProvinceId: "74",
			Name:       "KOTA KENDARI",
		},
		{
			Id:         "7472",
			ProvinceId: "74",
			Name:       "KOTA BAU BAU",
		},
		{
			Id:         "7501",
			ProvinceId: "75",
			Name:       "KAB. GORONTALO",
		},
		{
			Id:         "7502",
			ProvinceId: "75",
			Name:       "KAB. BOALEMO",
		},
		{
			Id:         "7503",
			ProvinceId: "75",
			Name:       "KAB. BONE BOLANGO",
		},
		{
			Id:         "7504",
			ProvinceId: "75",
			Name:       "KAB. POHUWATO",
		},
		{
			Id:         "7505",
			ProvinceId: "75",
			Name:       "KAB. GORONTALO UTARA",
		},
		{
			Id:         "7571",
			ProvinceId: "75",
			Name:       "KOTA GORONTALO",
		},
		{
			Id:         "7601",
			ProvinceId: "76",
			Name:       "KAB. PASANGKAYU",
		},
		{
			Id:         "7602",
			ProvinceId: "76",
			Name:       "KAB. MAMUJU",
		},
		{
			Id:         "7603",
			ProvinceId: "76",
			Name:       "KAB. MAMASA",
		},
		{
			Id:         "7604",
			ProvinceId: "76",
			Name:       "KAB. POLEWALI MANDAR",
		},
		{
			Id:         "7605",
			ProvinceId: "76",
			Name:       "KAB. MAJENE",
		},
		{
			Id:         "7606",
			ProvinceId: "76",
			Name:       "KAB. MAMUJU TENGAH",
		},
		{
			Id:         "8101",
			ProvinceId: "81",
			Name:       "KAB. MALUKU TENGAH",
		},
		{
			Id:         "8102",
			ProvinceId: "81",
			Name:       "KAB. MALUKU TENGGARA",
		},
		{
			Id:         "8103",
			ProvinceId: "81",
			Name:       "KAB. KEPULAUAN TANIMBAR",
		},
		{
			Id:         "8104",
			ProvinceId: "81",
			Name:       "KAB. BURU",
		},
		{
			Id:         "8105",
			ProvinceId: "81",
			Name:       "KAB. SERAM BAGIAN TIMUR",
		},
		{
			Id:         "8106",
			ProvinceId: "81",
			Name:       "KAB. SERAM BAGIAN BARAT",
		},
		{
			Id:         "8107",
			ProvinceId: "81",
			Name:       "KAB. KEPULAUAN ARU",
		},
		{
			Id:         "8108",
			ProvinceId: "81",
			Name:       "KAB. MALUKU BARAT DAYA",
		},
		{
			Id:         "8109",
			ProvinceId: "81",
			Name:       "KAB. BURU SELATAN",
		},
		{
			Id:         "8171",
			ProvinceId: "81",
			Name:       "KOTA AMBON",
		},
		{
			Id:         "8172",
			ProvinceId: "81",
			Name:       "KOTA TUAL",
		},
		{
			Id:         "8201",
			ProvinceId: "82",
			Name:       "KAB. HALMAHERA BARAT",
		},
		{
			Id:         "8202",
			ProvinceId: "82",
			Name:       "KAB. HALMAHERA TENGAH",
		},
		{
			Id:         "8203",
			ProvinceId: "82",
			Name:       "KAB. HALMAHERA UTARA",
		},
		{
			Id:         "8204",
			ProvinceId: "82",
			Name:       "KAB. HALMAHERA SELATAN",
		},
		{
			Id:         "8205",
			ProvinceId: "82",
			Name:       "KAB. KEPULAUAN SULA",
		},
		{
			Id:         "8206",
			ProvinceId: "82",
			Name:       "KAB. HALMAHERA TIMUR",
		},
		{
			Id:         "8207",
			ProvinceId: "82",
			Name:       "KAB. PULAU MOROTAI",
		},
		{
			Id:         "8208",
			ProvinceId: "82",
			Name:       "KAB. PULAU TALIABU",
		},
		{
			Id:         "8271",
			ProvinceId: "82",
			Name:       "KOTA TERNATE",
		},
		{
			Id:         "8272",
			ProvinceId: "82",
			Name:       "KOTA TIDORE KEPULAUAN",
		},
		{
			Id:         "9103",
			ProvinceId: "91",
			Name:       "KAB. JAYAPURA",
		},
		{
			Id:         "9105",
			ProvinceId: "91",
			Name:       "KAB. KEPULAUAN YAPEN",
		},
		{
			Id:         "9106",
			ProvinceId: "91",
			Name:       "KAB. BIAK NUMFOR",
		},
		{
			Id:         "9110",
			ProvinceId: "91",
			Name:       "KAB. SARMI",
		},
		{
			Id:         "9111",
			ProvinceId: "91",
			Name:       "KAB. KEEROM",
		},
		{
			Id:         "9115",
			ProvinceId: "91",
			Name:       "KAB. WAROPEN",
		},
		{
			Id:         "9119",
			ProvinceId: "91",
			Name:       "KAB. SUPIORI",
		},
		{
			Id:         "9120",
			ProvinceId: "91",
			Name:       "KAB. MAMBERAMO RAYA",
		},
		{
			Id:         "9171",
			ProvinceId: "91",
			Name:       "KOTA JAYAPURA",
		},
		{
			Id:         "9201",
			ProvinceId: "92",
			Name:       "KAB. SORONG",
		},
		{
			Id:         "9202",
			ProvinceId: "92",
			Name:       "KAB. MANOKWARI",
		},
		{
			Id:         "9203",
			ProvinceId: "92",
			Name:       "KAB. FAK FAK",
		},
		{
			Id:         "9204",
			ProvinceId: "92",
			Name:       "KAB. SORONG SELATAN",
		},
		{
			Id:         "9205",
			ProvinceId: "92",
			Name:       "KAB. RAJA AMPAT",
		},
		{
			Id:         "9206",
			ProvinceId: "92",
			Name:       "KAB. TELUK BINTUNI",
		},
		{
			Id:         "9207",
			ProvinceId: "92",
			Name:       "KAB. TELUK WONDAMA",
		},
		{
			Id:         "9208",
			ProvinceId: "92",
			Name:       "KAB. KAIMANA",
		},
		{
			Id:         "9209",
			ProvinceId: "92",
			Name:       "KAB. TAMBRAUW",
		},
		{
			Id:         "9210",
			ProvinceId: "92",
			Name:       "KAB. MAYBRAT",
		},
		{
			Id:         "9211",
			ProvinceId: "92",
			Name:       "KAB. MANOKWARI SELATAN",
		},
		{
			Id:         "9212",
			ProvinceId: "92",
			Name:       "KAB. PEGUNUNGAN ARFAK",
		},
		{
			Id:         "9271",
			ProvinceId: "92",
			Name:       "KOTA SORONG",
		},
		{
			Id:         "9301",
			ProvinceId: "93",
			Name:       "KAB. MERAUKE",
		},
		{
			Id:         "9302",
			ProvinceId: "93",
			Name:       "KAB. BOVEN DIGOEL",
		},
		{
			Id:         "9303",
			ProvinceId: "93",
			Name:       "KAB. MAPPI",
		},
		{
			Id:         "9304",
			ProvinceId: "93",
			Name:       "KAB. ASMAT",
		},
		{
			Id:         "9401",
			ProvinceId: "94",
			Name:       "KAB. NABIRE",
		},
		{
			Id:         "9402",
			ProvinceId: "94",
			Name:       "KAB. PUNCAK JAYA",
		},
		{
			Id:         "9403",
			ProvinceId: "94",
			Name:       "KAB. PANIAI",
		},
		{
			Id:         "9404",
			ProvinceId: "94",
			Name:       "KAB. MIMIKA",
		},
		{
			Id:         "9405",
			ProvinceId: "94",
			Name:       "KAB. PUNCAK",
		},
		{
			Id:         "9406",
			ProvinceId: "94",
			Name:       "KAB. DOGIYAI",
		},
		{
			Id:         "9407",
			ProvinceId: "94",
			Name:       "KAB. INTAN JAYA",
		},
		{
			Id:         "9408",
			ProvinceId: "94",
			Name:       "KAB. DEIYAI",
		},
		{
			Id:         "9501",
			ProvinceId: "95",
			Name:       "KAB. JAYAWIJAYA",
		},
		{
			Id:         "9502",
			ProvinceId: "95",
			Name:       "KAB. PEGUNUNGAN BINTANG",
		},
		{
			Id:         "9503",
			ProvinceId: "95",
			Name:       "KAB. YAHUKIMO",
		},
		{
			Id:         "9504",
			ProvinceId: "95",
			Name:       "KAB. TOLIKARA",
		},
		{
			Id:         "9505",
			ProvinceId: "95",
			Name:       "KAB. MAMBERAMO TENGAH",
		},
		{
			Id:         "9506",
			ProvinceId: "95",
			Name:       "KAB. YALIMO",
		},
		{
			Id:         "9507",
			ProvinceId: "95",
			Name:       "KAB. LANNY JAYA",
		},
		{
			Id:         "9508",
			ProvinceId: "95",
			Name:       "KAB. NDUGA",
		},
	}

	for _, city := range cities {
		if err := s.db.Where(entities.City{Name: city.Name}).
			FirstOrCreate(&city).Error; err != nil {
			log.Fatalf("failed to create city: %v", err)
		}
	}
}

// func getRandomCity(s Seed) (*entities.City, error) {
//	var city entities.City
//
//	if err := s.db.Order("RAND()").First(&city).Error; err != nil {
//		return nil, err
//	}
//
//	return &city, nil
//}
