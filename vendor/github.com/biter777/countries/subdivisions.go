package countries

import (
	"encoding/json"
	"fmt"
)

// SubdivisionCode - the code of a subdivision
type SubdivisionCode string

// Subdivision - all info about a subdivision
type Subdivision struct {
	Name            string              `json:"name"`
	Code            SubdivisionCode     `json:"code"`
	Country         CountryCode         `json:"countryCode"`
	SubdivisionType SubdivisionTypeCode `json:"type"`
}

// Type implements Typer interface
func (_ SubdivisionCode) Type() string {
	return TypeSubdivisionCode
}

// String - implements fmt.Stringer, returns an english name of the subdivision
//
//nolint:cyclop,funlen,gocyclo
func (s SubdivisionCode) String() string { //nolint:cyclop,gocyclo
	switch s {
	case SubdivisionAD02:
		return "Canillo"
	case SubdivisionAD03:
		return "Encamp"
	case SubdivisionAD04:
		return "La Massana"
	case SubdivisionAD05:
		return "Ordino"
	case SubdivisionAD06:
		return "Sant Julià de Lòria"
	case SubdivisionAD07:
		return "Andorra la Vella"
	case SubdivisionAD08:
		return "Escaldes-Engordany"
	case SubdivisionAEAJ:
		return "'Ajmān"
	case SubdivisionAEAZ:
		return "Abū Ȥaby [Abu Dhabi]"
	case SubdivisionAEDU:
		return "Dubayy"
	case SubdivisionAEFU:
		return "Al Fujayrah"
	case SubdivisionAERK:
		return "Ra’s al Khaymah"
	case SubdivisionAESH:
		return "Ash Shāriqah"
	case SubdivisionAEUQ:
		return "Umm al Qaywayn"
	case SubdivisionAFBAL:
		return "Balkh"
	case SubdivisionAFBAM:
		return "Bāmyān"
	case SubdivisionAFBDG:
		return "Bādghīs"
	case SubdivisionAFBDS:
		return "Badakhshān"
	case SubdivisionAFBGL:
		return "Baghlān"
	case SubdivisionAFDAY:
		return "Dāykundī"
	case SubdivisionAFFRA:
		return "Farāh"
	case SubdivisionAFFYB:
		return "Fāryāb"
	case SubdivisionAFGHA:
		return "Ghaznī"
	case SubdivisionAFGHO:
		return "Ghōr"
	case SubdivisionAFHEL:
		return "Helmand"
	case SubdivisionAFHER:
		return "Herāt"
	case SubdivisionAFJOW:
		return "Jowzjān"
	case SubdivisionAFKAB:
		return "Kābul"
	case SubdivisionAFKAN:
		return "Kandahār"
	case SubdivisionAFKAP:
		return "Kāpīsā"
	case SubdivisionAFKDZ:
		return "Kunduz"
	case SubdivisionAFKHO:
		return "Khōst"
	case SubdivisionAFKNR:
		return "Kunar"
	case SubdivisionAFLAG:
		return "Laghmān"
	case SubdivisionAFLOG:
		return "Lōgar"
	case SubdivisionAFNAN:
		return "Nangarhār"
	case SubdivisionAFNIM:
		return "Nīmrōz"
	case SubdivisionAFNUR:
		return "Nūristān"
	case SubdivisionAFPAN:
		return "Panjshayr"
	case SubdivisionAFPAR:
		return "Parwān"
	case SubdivisionAFPIA:
		return "Paktiyā"
	case SubdivisionAFPKA:
		return "Paktīkā"
	case SubdivisionAFSAM:
		return "Samangān"
	case SubdivisionAFSAR:
		return "Sar-e Pul"
	case SubdivisionAFTAK:
		return "Takhār"
	case SubdivisionAFURU:
		return "Uruzgān"
	case SubdivisionAFWAR:
		return "Wardak"
	case SubdivisionAFZAB:
		return "Zābul"
	case SubdivisionAG03:
		return "Saint George"
	case SubdivisionAG04:
		return "Saint John"
	case SubdivisionAG05:
		return "Saint Mary"
	case SubdivisionAG06:
		return "Saint Paul"
	case SubdivisionAG07:
		return "Saint Peter"
	case SubdivisionAG08:
		return "Saint Philip"
	case SubdivisionAG10:
		return "Barbuda"
	case SubdivisionAG11:
		return "Redonda"
	case SubdivisionAL01:
		return "Berat"
	case SubdivisionAL02:
		return "Durrës"
	case SubdivisionAL03:
		return "Elbasan"
	case SubdivisionAL04:
		return "Fier"
	case SubdivisionAL05:
		return "Gjirokastër"
	case SubdivisionAL06:
		return "Korçë"
	case SubdivisionAL07:
		return "Kukës"
	case SubdivisionAL08:
		return "Lezhë"
	case SubdivisionAL09:
		return "Dibër"
	case SubdivisionAL10:
		return "Shkodër"
	case SubdivisionAL11:
		return "Tiranë"
	case SubdivisionAL12:
		return "Vlorë"
	case SubdivisionALBR:
		return "Berat"
	case SubdivisionALBU:
		return "Bulqizë"
	case SubdivisionALDI:
		return "Dibër"
	case SubdivisionALDL:
		return "Delvinë"
	case SubdivisionALDR:
		return "Durrës"
	case SubdivisionALDV:
		return "Devoll"
	case SubdivisionALEL:
		return "Elbasan"
	case SubdivisionALER:
		return "Kolonjë"
	case SubdivisionALFR:
		return "Fier"
	case SubdivisionALGJ:
		return "Gjirokastër"
	case SubdivisionALGR:
		return "Gramsh"
	case SubdivisionALHA:
		return "Has"
	case SubdivisionALKA:
		return "Kavajë"
	case SubdivisionALKB:
		return "Kurbin"
	case SubdivisionALKC:
		return "Kuçovë"
	case SubdivisionALKO:
		return "Korçë"
	case SubdivisionALKR:
		return "Krujë"
	case SubdivisionALKU:
		return "Kukës"
	case SubdivisionALLB:
		return "Librazhd"
	case SubdivisionALLE:
		return "Lezhë"
	case SubdivisionALLU:
		return "Lushnjë"
	case SubdivisionALMK:
		return "Mallakastër"
	case SubdivisionALMM:
		return "Malësi e Madhe"
	case SubdivisionALMR:
		return "Mirditë"
	case SubdivisionALMT:
		return "Mat"
	case SubdivisionALPG:
		return "Pogradec"
	case SubdivisionALPQ:
		return "Peqin"
	case SubdivisionALPR:
		return "Përmet"
	case SubdivisionALPU:
		return "Pukë"
	case SubdivisionALSH:
		return "Shkodër"
	case SubdivisionALSK:
		return "Skrapar"
	case SubdivisionALSR:
		return "Sarandë"
	case SubdivisionALTE:
		return "Tepelenë"
	case SubdivisionALTP:
		return "Tropojë"
	case SubdivisionALTR:
		return "Tiranë"
	case SubdivisionALVL:
		return "Vlorë"
	case SubdivisionAMAG:
		return "Aragacotn"
	case SubdivisionAMAR:
		return "Ararat"
	case SubdivisionAMAV:
		return "Armavir"
	case SubdivisionAMER:
		return "Erevan"
	case SubdivisionAMGR:
		return "Gegarkunik'"
	case SubdivisionAMKT:
		return "Kotayk'"
	case SubdivisionAMLO:
		return "Lory"
	case SubdivisionAMSH:
		return "Sirak"
	case SubdivisionAMSU:
		return "Syunik'"
	case SubdivisionAMTV:
		return "Tavus"
	case SubdivisionAMVD:
		return "Vayoc Jor"
	case SubdivisionAOBGO:
		return "Bengo"
	case SubdivisionAOBGU:
		return "Benguela"
	case SubdivisionAOBIE:
		return "Bié"
	case SubdivisionAOCAB:
		return "Cabinda"
	case SubdivisionAOCCU:
		return "Cuando-Cubango"
	case SubdivisionAOCNN:
		return "Cunene"
	case SubdivisionAOCNO:
		return "Cuanza Norte"
	case SubdivisionAOCUS:
		return "Cuanza Sul"
	case SubdivisionAOHUA:
		return "Huambo"
	case SubdivisionAOHUI:
		return "Huíla"
	case SubdivisionAOLNO:
		return "Lunda Norte"
	case SubdivisionAOLSU:
		return "Lunda Sul"
	case SubdivisionAOLUA:
		return "Luanda"
	case SubdivisionAOMAL:
		return "Malange"
	case SubdivisionAOMOX:
		return "Moxico"
	case SubdivisionAONAM:
		return "Namibe"
	case SubdivisionAOUIG:
		return "Uíge"
	case SubdivisionAOZAI:
		return "Zaire"
	case SubdivisionARA:
		return "Salta"
	case SubdivisionARB:
		return "Buenos Aires"
	case SubdivisionARC:
		return "Ciudad Autónoma de Buenos Aires"
	case SubdivisionARD:
		return "San Luis"
	case SubdivisionARE:
		return "Entre Rios"
	case SubdivisionARG:
		return "Santiago del Estero"
	case SubdivisionARH:
		return "Chaco"
	case SubdivisionARJ:
		return "San Juan"
	case SubdivisionARK:
		return "Catamarca"
	case SubdivisionARL:
		return "La Pampa"
	case SubdivisionARM:
		return "Mendoza"
	case SubdivisionARN:
		return "Misiones"
	case SubdivisionARP:
		return "Formosa"
	case SubdivisionARQ:
		return "Neuquen"
	case SubdivisionARR:
		return "Rio Negro"
	case SubdivisionARS:
		return "Santa Fe"
	case SubdivisionART:
		return "Tucuman"
	case SubdivisionARU:
		return "Chubut"
	case SubdivisionARV:
		return "Tierra del Fuego"
	case SubdivisionARW:
		return "Corrientes"
	case SubdivisionARX:
		return "Cordoba"
	case SubdivisionARY:
		return "Jujuy"
	case SubdivisionARZ:
		return "Santa Cruz"
	case SubdivisionAT1:
		return "Burgenland"
	case SubdivisionAT2:
		return "Kärnten"
	case SubdivisionAT3:
		return "Niederösterreich"
	case SubdivisionAT4:
		return "Oberösterreich"
	case SubdivisionAT5:
		return "Salzburg"
	case SubdivisionAT6:
		return "Steiermark"
	case SubdivisionAT7:
		return "Tirol"
	case SubdivisionAT8:
		return "Vorarlberg"
	case SubdivisionAT9:
		return "Wien"
	case SubdivisionAUACT:
		return "Australian Capital Territory"
	case SubdivisionAUNSW:
		return "New South Wales"
	case SubdivisionAUNT:
		return "Northern Territory"
	case SubdivisionAUQLD:
		return "Queensland"
	case SubdivisionAUSA:
		return "South Australia"
	case SubdivisionAUTAS:
		return "Tasmania"
	case SubdivisionAUVIC:
		return "Victoria"
	case SubdivisionAUWA:
		return "Western Australia"
	case SubdivisionAZABS:
		return "Abşeron"
	case SubdivisionAZAGA:
		return "Ağstafa"
	case SubdivisionAZAGC:
		return "Ağcabədi"
	case SubdivisionAZAGM:
		return "Ağdam"
	case SubdivisionAZAGS:
		return "Ağdaş"
	case SubdivisionAZAGU:
		return "Ağsu"
	case SubdivisionAZAST:
		return "Astara"
	case SubdivisionAZBA:
		return "Bakı"
	case SubdivisionAZBAB:
		return "Babək"
	case SubdivisionAZBAL:
		return "Balakən"
	case SubdivisionAZBAR:
		return "Bərdə"
	case SubdivisionAZBEY:
		return "Beyləqan"
	case SubdivisionAZBIL:
		return "Biləsuvar"
	case SubdivisionAZCAB:
		return "Cəbrayıl"
	case SubdivisionAZCAL:
		return "Cəlilabab"
	case SubdivisionAZCUL:
		return "Culfa"
	case SubdivisionAZDAS:
		return "Daşkəsən"
	case SubdivisionAZFUZ:
		return "Füzuli"
	case SubdivisionAZGA:
		return "Gəncə"
	case SubdivisionAZGAD:
		return "Gədəbəy"
	case SubdivisionAZGOR:
		return "Goranboy"
	case SubdivisionAZGOY:
		return "Göyçay"
	case SubdivisionAZGYG:
		return "Göygöl"
	case SubdivisionAZHAC:
		return "Hacıqabul"
	case SubdivisionAZIMI:
		return "İmişli"
	case SubdivisionAZISM:
		return "İsmayıllı"
	case SubdivisionAZKAL:
		return "Kəlbəcər"
	case SubdivisionAZKAN:
		return "Kǝngǝrli"
	case SubdivisionAZKUR:
		return "Kürdəmir"
	case SubdivisionAZLA:
		return "Lənkəran"
	case SubdivisionAZLAC:
		return "Laçın"
	case SubdivisionAZLAN:
		return "Lənkəran"
	case SubdivisionAZLER:
		return "Lerik"
	case SubdivisionAZMAS:
		return "Masallı"
	case SubdivisionAZMI:
		return "Mingəçevir"
	case SubdivisionAZNA:
		return "Naftalan"
	case SubdivisionAZNEF:
		return "Neftçala"
	case SubdivisionAZNV:
		return "Naxçıvan"
	case SubdivisionAZNX:
		return "Naxçıvan"
	case SubdivisionAZOGU:
		return "Oğuz"
	case SubdivisionAZORD:
		return "Ordubad"
	case SubdivisionAZQAB:
		return "Qəbələ"
	case SubdivisionAZQAX:
		return "Qax"
	case SubdivisionAZQAZ:
		return "Qazax"
	case SubdivisionAZQBA:
		return "Quba"
	case SubdivisionAZQBI:
		return "Qubadlı"
	case SubdivisionAZQOB:
		return "Qobustan"
	case SubdivisionAZQUS:
		return "Qusar"
	case SubdivisionAZSA:
		return "Şəki"
	case SubdivisionAZSAB:
		return "Sabirabad"
	case SubdivisionAZSAD:
		return "Sədərək"
	case SubdivisionAZSAH:
		return "Şahbuz"
	case SubdivisionAZSAK:
		return "Şəki"
	case SubdivisionAZSAL:
		return "Salyan"
	case SubdivisionAZSAR:
		return "Şərur"
	case SubdivisionAZSAT:
		return "Saatlı"
	case SubdivisionAZSBN:
		return "Şabran"
	case SubdivisionAZSIY:
		return "Siyəzən"
	case SubdivisionAZSKR:
		return "Şəmkir"
	case SubdivisionAZSM:
		return "Sumqayıt"
	case SubdivisionAZSMI:
		return "Şamaxı"
	case SubdivisionAZSMX:
		return "Samux"
	case SubdivisionAZSR:
		return "Şirvan"
	case SubdivisionAZSUS:
		return "Şuşa"
	case SubdivisionAZTAR:
		return "Tərtər"
	case SubdivisionAZTOV:
		return "Tovuz"
	case SubdivisionAZUCA:
		return "Ucar"
	case SubdivisionAZXA:
		return "Xankəndi"
	case SubdivisionAZXAC:
		return "Xaçmaz"
	case SubdivisionAZXCI:
		return "Xocalı"
	case SubdivisionAZXIZ:
		return "Xızı"
	case SubdivisionAZXVD:
		return "Xocavənd"
	case SubdivisionAZYAR:
		return "Yardımlı"
	case SubdivisionAZYE:
		return "Yevlax"
	case SubdivisionAZYEV:
		return "Yevlax"
	case SubdivisionAZZAN:
		return "Zəngilan"
	case SubdivisionAZZAQ:
		return "Zaqatala"
	case SubdivisionAZZAR:
		return "Zərdab"
	case SubdivisionBA01:
		return "Unsko-sanski kanton"
	case SubdivisionBA02:
		return "Posavski kanton"
	case SubdivisionBA03:
		return "Tuzlanski kanton"
	case SubdivisionBA04:
		return "Zeničko-dobojski kanton"
	case SubdivisionBA05:
		return "Bosansko-podrinjski kanton"
	case SubdivisionBA06:
		return "Srednjobosanski kanton"
	case SubdivisionBA07:
		return "Hercegovačko-neretvanski kanton"
	case SubdivisionBA08:
		return "Zapadnohercegovački kanton"
	case SubdivisionBA09:
		return "Kanton Sarajevo"
	case SubdivisionBA10:
		return "Kanton br. 10 (Livanjski kanton)"
	case SubdivisionBABIH:
		return "Federacija Bosne i Hercegovine"
	case SubdivisionBABRC:
		return "Brčko distrikt"
	case SubdivisionBASRP:
		return "Republika Srpska"
	case SubdivisionBB01:
		return "Christ Church"
	case SubdivisionBB02:
		return "Saint Andrew"
	case SubdivisionBB03:
		return "Saint George"
	case SubdivisionBB04:
		return "Saint James"
	case SubdivisionBB05:
		return "Saint John"
	case SubdivisionBB06:
		return "Saint Joseph"
	case SubdivisionBB07:
		return "Saint Lucy"
	case SubdivisionBB08:
		return "Saint Michael"
	case SubdivisionBB09:
		return "Saint Peter"
	case SubdivisionBB10:
		return "Saint Philip"
	case SubdivisionBB11:
		return "Saint Thomas"
	case SubdivisionBD01:
		return "Bandarban"
	case SubdivisionBD02:
		return "Barguna"
	case SubdivisionBD03:
		return "Bogra"
	case SubdivisionBD04:
		return "Brahmanbaria"
	case SubdivisionBD05:
		return "Bagerhat"
	case SubdivisionBD06:
		return "Barisal"
	case SubdivisionBD07:
		return "Bhola"
	case SubdivisionBD08:
		return "Comilla"
	case SubdivisionBD09:
		return "Chandpur"
	case SubdivisionBD10:
		return "Chittagong"
	case SubdivisionBD11:
		return "Cox's Bazar"
	case SubdivisionBD12:
		return "Chuadanga"
	case SubdivisionBD13:
		return "Dhaka"
	case SubdivisionBD14:
		return "Dinajpur"
	case SubdivisionBD15:
		return "Faridpur"
	case SubdivisionBD16:
		return "Feni"
	case SubdivisionBD17:
		return "Gopalganj"
	case SubdivisionBD18:
		return "Gazipur"
	case SubdivisionBD19:
		return "Gaibandha"
	case SubdivisionBD20:
		return "Habiganj"
	case SubdivisionBD21:
		return "Jamalpur"
	case SubdivisionBD22:
		return "Jessore"
	case SubdivisionBD23:
		return "Jhenaidah"
	case SubdivisionBD24:
		return "Jaipurhat"
	case SubdivisionBD25:
		return "Jhalakati"
	case SubdivisionBD26:
		return "Kishorganj"
	case SubdivisionBD27:
		return "Khulna"
	case SubdivisionBD28:
		return "Kurigram"
	case SubdivisionBD29:
		return "Khagrachari"
	case SubdivisionBD30:
		return "Kushtia"
	case SubdivisionBD31:
		return "Lakshmipur"
	case SubdivisionBD32:
		return "Lalmonirhat"
	case SubdivisionBD33:
		return "Manikganj"
	case SubdivisionBD34:
		return "Mymensingh"
	case SubdivisionBD35:
		return "Munshiganj"
	case SubdivisionBD36:
		return "Madaripur"
	case SubdivisionBD37:
		return "Magura"
	case SubdivisionBD38:
		return "Moulvibazar"
	case SubdivisionBD39:
		return "Meherpur"
	case SubdivisionBD40:
		return "Narayanganj"
	case SubdivisionBD41:
		return "Netrakona"
	case SubdivisionBD42:
		return "Narsingdi"
	case SubdivisionBD43:
		return "Narail"
	case SubdivisionBD44:
		return "Natore"
	case SubdivisionBD45:
		return "Nawabganj"
	case SubdivisionBD46:
		return "Nilphamari"
	case SubdivisionBD47:
		return "Noakhali"
	case SubdivisionBD48:
		return "Naogaon"
	case SubdivisionBD49:
		return "Pabna"
	case SubdivisionBD50:
		return "Pirojpur"
	case SubdivisionBD51:
		return "Patuakhali"
	case SubdivisionBD52:
		return "Panchagarh"
	case SubdivisionBD53:
		return "Rajbari"
	case SubdivisionBD54:
		return "Rajshahi"
	case SubdivisionBD55:
		return "Rangpur"
	case SubdivisionBD56:
		return "Rangamati"
	case SubdivisionBD57:
		return "Sherpur"
	case SubdivisionBD58:
		return "Satkhira"
	case SubdivisionBD59:
		return "Sirajganj"
	case SubdivisionBD60:
		return "Sylhet"
	case SubdivisionBD61:
		return "Sunamganj"
	case SubdivisionBD62:
		return "Shariatpur"
	case SubdivisionBD63:
		return "Tangail"
	case SubdivisionBD64:
		return "Thakurgaon"
	case SubdivisionBDA:
		return "Barisal"
	case SubdivisionBDB:
		return "Chittagong"
	case SubdivisionBDC:
		return "Dhaka"
	case SubdivisionBDD:
		return "Khulna"
	case SubdivisionBDE:
		return "Rajshahi"
	case SubdivisionBDF:
		return "Rangpur"
	case SubdivisionBDG:
		return "Sylhet"
	case SubdivisionBDH:
		return "Mymensingh"
	case SubdivisionBEBRU:
		return "Bruxelles-Capitale, Région de;Brussels Hoofdstedelijk Gewest"
	case SubdivisionBEVAN:
		return "Antwerpen"
	case SubdivisionBEVBR:
		return "Vlaams-Brabant"
	case SubdivisionBEVLG:
		return "Vlaams Gewest"
	case SubdivisionBEVLI:
		return "Limburg"
	case SubdivisionBEVOV:
		return "Oost-Vlaanderen"
	case SubdivisionBEVWV:
		return "West-Vlaanderen"
	case SubdivisionBEWAL:
		return "wallonne, Région"
	case SubdivisionBEWBR:
		return "Brabant wallon"
	case SubdivisionBEWHT:
		return "Hainaut"
	case SubdivisionBEWLG:
		return "Liège"
	case SubdivisionBEWLX:
		return "Luxembourg"
	case SubdivisionBEWNA:
		return "Namur"
	case SubdivisionBF01:
		return "Boucle du Mouhoun"
	case SubdivisionBF02:
		return "Cascades"
	case SubdivisionBF03:
		return "Centre"
	case SubdivisionBF04:
		return "Centre-Est"
	case SubdivisionBF05:
		return "Centre-Nord"
	case SubdivisionBF06:
		return "Centre-Ouest"
	case SubdivisionBF07:
		return "Centre-Sud"
	case SubdivisionBF08:
		return "Est"
	case SubdivisionBF09:
		return "Hauts-Bassins"
	case SubdivisionBF10:
		return "Nord"
	case SubdivisionBF11:
		return "Plateau-Central"
	case SubdivisionBF12:
		return "Sahel"
	case SubdivisionBF13:
		return "Sud-Ouest"
	case SubdivisionBFBAL:
		return "Balé"
	case SubdivisionBFBAM:
		return "Bam"
	case SubdivisionBFBAN:
		return "Banwa"
	case SubdivisionBFBAZ:
		return "Bazèga"
	case SubdivisionBFBGR:
		return "Bougouriba"
	case SubdivisionBFBLG:
		return "Boulgou"
	case SubdivisionBFBLK:
		return "Boulkiemdé"
	case SubdivisionBFCOM:
		return "Comoé"
	case SubdivisionBFGAN:
		return "Ganzourgou"
	case SubdivisionBFGNA:
		return "Gnagna"
	case SubdivisionBFGOU:
		return "Gourma"
	case SubdivisionBFHOU:
		return "Houet"
	case SubdivisionBFIOB:
		return "Ioba"
	case SubdivisionBFKAD:
		return "Kadiogo"
	case SubdivisionBFKEN:
		return "Kénédougou"
	case SubdivisionBFKMD:
		return "Komondjari"
	case SubdivisionBFKMP:
		return "Kompienga"
	case SubdivisionBFKOP:
		return "Koulpélogo"
	case SubdivisionBFKOS:
		return "Kossi"
	case SubdivisionBFKOT:
		return "Kouritenga"
	case SubdivisionBFKOW:
		return "Kourwéogo"
	case SubdivisionBFLER:
		return "Léraba"
	case SubdivisionBFLOR:
		return "Loroum"
	case SubdivisionBFMOU:
		return "Mouhoun"
	case SubdivisionBFNAM:
		return "Namentenga"
	case SubdivisionBFNAO:
		return "Naouri"
	case SubdivisionBFNAY:
		return "Nayala"
	case SubdivisionBFNOU:
		return "Noumbiel"
	case SubdivisionBFOUB:
		return "Oubritenga"
	case SubdivisionBFOUD:
		return "Oudalan"
	case SubdivisionBFPAS:
		return "Passoré"
	case SubdivisionBFPON:
		return "Poni"
	case SubdivisionBFSEN:
		return "Séno"
	case SubdivisionBFSIS:
		return "Sissili"
	case SubdivisionBFSMT:
		return "Sanmatenga"
	case SubdivisionBFSNG:
		return "Sanguié"
	case SubdivisionBFSOM:
		return "Soum"
	case SubdivisionBFSOR:
		return "Sourou"
	case SubdivisionBFTAP:
		return "Tapoa"
	case SubdivisionBFTUI:
		return "Tui"
	case SubdivisionBFYAG:
		return "Yagha"
	case SubdivisionBFYAT:
		return "Yatenga"
	case SubdivisionBFZIR:
		return "Ziro"
	case SubdivisionBFZON:
		return "Zondoma"
	case SubdivisionBFZOU:
		return "Zoundwéogo"
	case SubdivisionBG01:
		return "Blagoevgrad"
	case SubdivisionBG02:
		return "Burgas"
	case SubdivisionBG03:
		return "Varna"
	case SubdivisionBG04:
		return "Veliko Tarnovo"
	case SubdivisionBG05:
		return "Vidin"
	case SubdivisionBG06:
		return "Vratsa"
	case SubdivisionBG07:
		return "Gabrovo"
	case SubdivisionBG08:
		return "Dobrich"
	case SubdivisionBG09:
		return "Kardzhali"
	case SubdivisionBG10:
		return "Kyustendil"
	case SubdivisionBG11:
		return "Lovech"
	case SubdivisionBG12:
		return "Montana"
	case SubdivisionBG13:
		return "Pazardzhik"
	case SubdivisionBG14:
		return "Pernik"
	case SubdivisionBG15:
		return "Pleven"
	case SubdivisionBG16:
		return "Plovdiv"
	case SubdivisionBG17:
		return "Razgrad"
	case SubdivisionBG18:
		return "Ruse"
	case SubdivisionBG19:
		return "Silistra"
	case SubdivisionBG20:
		return "Sliven"
	case SubdivisionBG21:
		return "Smolyan"
	case SubdivisionBG22:
		return "Sofia-Grad"
	case SubdivisionBG23:
		return "Sofia"
	case SubdivisionBG24:
		return "Stara Zagora"
	case SubdivisionBG25:
		return "Targovishte"
	case SubdivisionBG26:
		return "Haskovo"
	case SubdivisionBG27:
		return "Shumen"
	case SubdivisionBG28:
		return "Yambol"
	case SubdivisionBH13:
		return "Al Manāmah (Al ‘Āşimah)"
	case SubdivisionBH14:
		return "Al Janūbīyah"
	case SubdivisionBH15:
		return "Al Muḩarraq"
	case SubdivisionBH16:
		return "Al Wusţá"
	case SubdivisionBH17:
		return "Ash Shamālīyah"
	case SubdivisionBIBB:
		return "Bubanza"
	case SubdivisionBIBL:
		return "Bujumbura Rural"
	case SubdivisionBIBM:
		return "Bujumbura Mairie"
	case SubdivisionBIBR:
		return "Bururi"
	case SubdivisionBICA:
		return "Cankuzo"
	case SubdivisionBICI:
		return "Cibitoke"
	case SubdivisionBIGI:
		return "Gitega"
	case SubdivisionBIKI:
		return "Kirundo"
	case SubdivisionBIKR:
		return "Karuzi"
	case SubdivisionBIKY:
		return "Kayanza"
	case SubdivisionBIMA:
		return "Makamba"
	case SubdivisionBIMU:
		return "Muramvya"
	case SubdivisionBIMW:
		return "Mwaro"
	case SubdivisionBING:
		return "Ngozi"
	case SubdivisionBIRT:
		return "Rutana"
	case SubdivisionBIRY:
		return "Ruyigi"
	case SubdivisionBJAK:
		return "Atakora"
	case SubdivisionBJAL:
		return "Alibori"
	case SubdivisionBJAQ:
		return "Atlantique"
	case SubdivisionBJBO:
		return "Borgou"
	case SubdivisionBJCO:
		return "Collines"
	case SubdivisionBJDO:
		return "Donga"
	case SubdivisionBJKO:
		return "Kouffo"
	case SubdivisionBJLI:
		return "Littoral"
	case SubdivisionBJMO:
		return "Mono"
	case SubdivisionBJOU:
		return "Ouémé"
	case SubdivisionBJPL:
		return "Plateau"
	case SubdivisionBJZO:
		return "Zou"
	case SubdivisionBNBE:
		return "Belait"
	case SubdivisionBNBM:
		return "Brunei-Muara"
	case SubdivisionBNTE:
		return "Temburong"
	case SubdivisionBNTU:
		return "Tutong"
	case SubdivisionBOB:
		return "El Beni"
	case SubdivisionBOC:
		return "Cochabamba"
	case SubdivisionBOH:
		return "Chuquisaca"
	case SubdivisionBOL:
		return "La Paz"
	case SubdivisionBON:
		return "Pando"
	case SubdivisionBOO:
		return "Oruro"
	case SubdivisionBOP:
		return "Potosí"
	case SubdivisionBOS:
		return "Santa Cruz"
	case SubdivisionBOT:
		return "Tarija"
	case SubdivisionBQBO:
		return "Bonaire"
	case SubdivisionBQSA:
		return "Saba"
	case SubdivisionBQSE:
		return "Sint Eustatius"
	case SubdivisionBRAC:
		return "Acre"
	case SubdivisionBRAL:
		return "Alagoas"
	case SubdivisionBRAM:
		return "Amazonas"
	case SubdivisionBRAP:
		return "Amapá"
	case SubdivisionBRBA:
		return "Bahia"
	case SubdivisionBRCE:
		return "Ceará"
	case SubdivisionBRDF:
		return "Distrito Federal"
	case SubdivisionBRES:
		return "Espírito Santo"
	case SubdivisionBRFN:
		return "Fernando de Noronha"
	case SubdivisionBRGO:
		return "Goiás"
	case SubdivisionBRMA:
		return "Maranhão"
	case SubdivisionBRMG:
		return "Minas Gerais"
	case SubdivisionBRMS:
		return "Mato Grosso do Sul"
	case SubdivisionBRMT:
		return "Mato Grosso"
	case SubdivisionBRPA:
		return "Pará"
	case SubdivisionBRPB:
		return "Paraíba"
	case SubdivisionBRPE:
		return "Pernambuco"
	case SubdivisionBRPI:
		return "Piauí"
	case SubdivisionBRPR:
		return "Paraná"
	case SubdivisionBRRJ:
		return "Rio de Janeiro"
	case SubdivisionBRRN:
		return "Rio Grande do Norte"
	case SubdivisionBRRO:
		return "Rondônia"
	case SubdivisionBRRR:
		return "Roraima"
	case SubdivisionBRRS:
		return "Rio Grande do Sul"
	case SubdivisionBRSC:
		return "Santa Catarina"
	case SubdivisionBRSE:
		return "Sergipe"
	case SubdivisionBRSP:
		return "São Paulo"
	case SubdivisionBRTO:
		return "Tocantins"
	case SubdivisionBSAK:
		return "Acklins"
	case SubdivisionBSBI:
		return "Bimini"
	case SubdivisionBSBP:
		return "Black Point"
	case SubdivisionBSBY:
		return "Berry Islands"
	case SubdivisionBSCE:
		return "Central Eleuthera"
	case SubdivisionBSCI:
		return "Cat Island"
	case SubdivisionBSCK:
		return "Crooked Island and Long Cay"
	case SubdivisionBSCO:
		return "Central Abaco"
	case SubdivisionBSCS:
		return "Central Andros"
	case SubdivisionBSEG:
		return "East Grand Bahama"
	case SubdivisionBSEX:
		return "Exuma"
	case SubdivisionBSFP:
		return "City of Freeport"
	case SubdivisionBSGC:
		return "Grand Cay"
	case SubdivisionBSHI:
		return "Harbour Island"
	case SubdivisionBSHT:
		return "Hope Town"
	case SubdivisionBSIN:
		return "Inagua"
	case SubdivisionBSLI:
		return "Long Island"
	case SubdivisionBSMC:
		return "Mangrove Cay"
	case SubdivisionBSMG:
		return "Mayaguana"
	case SubdivisionBSMI:
		return "Moore's Island"
	case SubdivisionBSNE:
		return "North Eleuthera"
	case SubdivisionBSNO:
		return "North Abaco"
	case SubdivisionBSNS:
		return "North Andros"
	case SubdivisionBSRC:
		return "Rum Cay"
	case SubdivisionBSRI:
		return "Ragged Island"
	case SubdivisionBSSA:
		return "South Andros"
	case SubdivisionBSSE:
		return "South Eleuthera"
	case SubdivisionBSSO:
		return "South Abaco"
	case SubdivisionBSSS:
		return "San Salvador"
	case SubdivisionBSSW:
		return "Spanish Wells"
	case SubdivisionBSWG:
		return "West Grand Bahama"
	case SubdivisionBT11:
		return "Paro"
	case SubdivisionBT12:
		return "Chhukha"
	case SubdivisionBT13:
		return "Ha"
	case SubdivisionBT14:
		return "Samtee"
	case SubdivisionBT15:
		return "Thimphu"
	case SubdivisionBT21:
		return "Tsirang"
	case SubdivisionBT22:
		return "Dagana"
	case SubdivisionBT23:
		return "Punakha"
	case SubdivisionBT24:
		return "Wangdue Phodrang"
	case SubdivisionBT31:
		return "Sarpang"
	case SubdivisionBT32:
		return "Trongsa"
	case SubdivisionBT33:
		return "Bumthang"
	case SubdivisionBT34:
		return "Zhemgang"
	case SubdivisionBT41:
		return "Trashigang"
	case SubdivisionBT42:
		return "Monggar"
	case SubdivisionBT43:
		return "Pemagatshel"
	case SubdivisionBT44:
		return "Lhuentse"
	case SubdivisionBT45:
		return "Samdrup Jongkha"
	case SubdivisionBTGA:
		return "Gasa"
	case SubdivisionBTTY:
		return "Trashi Yangtse"
	case SubdivisionBWCE:
		return "Central"
	case SubdivisionBWGH:
		return "Ghanzi"
	case SubdivisionBWKG:
		return "Kgalagadi"
	case SubdivisionBWKL:
		return "Kgatleng"
	case SubdivisionBWKW:
		return "Kweneng"
	case SubdivisionBWNE:
		return "North-East"
	case SubdivisionBWNW:
		return "North-West"
	case SubdivisionBWSE:
		return "South-East"
	case SubdivisionBWSO:
		return "Southern"
	case SubdivisionBYBR:
		return "Bresckaja voblasć"
	case SubdivisionBYHM:
		return "Horad Minsk"
	case SubdivisionBYHO:
		return "Homieĺskaja voblasć"
	case SubdivisionBYHR:
		return "Hrodzienskaja voblasć"
	case SubdivisionBYMA:
		return "Mahilioŭskaja voblasć"
	case SubdivisionBYMI:
		return "Minskaja voblasć"
	case SubdivisionBYVI:
		return "Viciebskaja voblasć"
	case SubdivisionBZBZ:
		return "Belize"
	case SubdivisionBZCY:
		return "Cayo"
	case SubdivisionBZCZL:
		return "Corozal"
	case SubdivisionBZOW:
		return "Orange Walk"
	case SubdivisionBZSC:
		return "Stann Creek"
	case SubdivisionBZTOL:
		return "Toledo"
	case SubdivisionCAAB:
		return "Alberta"
	case SubdivisionCABC:
		return "British Columbia"
	case SubdivisionCAMB:
		return "Manitoba"
	case SubdivisionCANB:
		return "New Brunswick"
	case SubdivisionCANL:
		return "Newfoundland and Labrador"
	case SubdivisionCANS:
		return "Nova Scotia"
	case SubdivisionCANT:
		return "Northwest Territories"
	case SubdivisionCANU:
		return "Nunavut"
	case SubdivisionCAON:
		return "Ontario"
	case SubdivisionCAPE:
		return "Prince Edward Island"
	case SubdivisionCAQC:
		return "Quebec"
	case SubdivisionCASK:
		return "Saskatchewan"
	case SubdivisionCAYT:
		return "Yukon Territory"
	case SubdivisionCDBC:
		return "Bas-Congo"
	case SubdivisionCDBN:
		return "Bandundu"
	case SubdivisionCDEQ:
		return "Équateur"
	case SubdivisionCDKA:
		return "Katanga"
	case SubdivisionCDKE:
		return "Kasai-Oriental"
	case SubdivisionCDKN:
		return "Kinshasa"
	case SubdivisionCDKW:
		return "Kasai-Occidental"
	case SubdivisionCDMA:
		return "Maniema"
	case SubdivisionCDNK:
		return "Nord-Kivu"
	case SubdivisionCDOR:
		return "Orientale"
	case SubdivisionCDSK:
		return "Sud-Kivu"
	case SubdivisionCFAC:
		return "Ouham"
	case SubdivisionCFBB:
		return "Bamingui-Bangoran"
	case SubdivisionCFBGF:
		return "Bangui"
	case SubdivisionCFBK:
		return "Basse-Kotto"
	case SubdivisionCFHK:
		return "Haute-Kotto"
	case SubdivisionCFHM:
		return "Haut-Mbomou"
	case SubdivisionCFHS:
		return "Haute-Sangha / Mambéré-Kadéï"
	case SubdivisionCFKB:
		return "Gribingui"
	case SubdivisionCFKG:
		return "Kémo-Gribingui"
	case SubdivisionCFLB:
		return "Lobaye"
	case SubdivisionCFMB:
		return "Mbomou"
	case SubdivisionCFMP:
		return "Ombella-M'poko"
	case SubdivisionCFNM:
		return "Nana-Mambéré"
	case SubdivisionCFOP:
		return "Ouham-Pendé"
	case SubdivisionCFSE:
		return "Sangha"
	case SubdivisionCFUK:
		return "Ouaka"
	case SubdivisionCFVK:
		return "Vakaga"
	case SubdivisionCG11:
		return "Bouenza"
	case SubdivisionCG12:
		return "Pool"
	case SubdivisionCG13:
		return "Sangha"
	case SubdivisionCG14:
		return "Plateaux"
	case SubdivisionCG15:
		return "Cuvette-Ouest"
	case SubdivisionCG2:
		return "Lékoumou"
	case SubdivisionCG5:
		return "Kouilou"
	case SubdivisionCG7:
		return "Likouala"
	case SubdivisionCG8:
		return "Cuvette"
	case SubdivisionCG9:
		return "Niari"
	case SubdivisionCGBZV:
		return "Brazzaville"
	case SubdivisionCHAG:
		return "Aargau"
	case SubdivisionCHAI:
		return "Appenzell Innerrhoden"
	case SubdivisionCHAR:
		return "Appenzell Ausserrhoden"
	case SubdivisionCHBE:
		return "Bern"
	case SubdivisionCHBL:
		return "Basel-Landschaft"
	case SubdivisionCHBS:
		return "Basel-Stadt"
	case SubdivisionCHFR:
		return "Fribourg"
	case SubdivisionCHGE:
		return "Genève"
	case SubdivisionCHGL:
		return "Glarus"
	case SubdivisionCHGR:
		return "Graubünden"
	case SubdivisionCHJU:
		return "Jura"
	case SubdivisionCHLU:
		return "Luzern"
	case SubdivisionCHNE:
		return "Neuchâtel"
	case SubdivisionCHNW:
		return "Nidwalden"
	case SubdivisionCHOW:
		return "Obwalden"
	case SubdivisionCHSG:
		return "Sankt Gallen"
	case SubdivisionCHSH:
		return "Schaffhausen"
	case SubdivisionCHSO:
		return "Solothurn"
	case SubdivisionCHSZ:
		return "Schwyz"
	case SubdivisionCHTG:
		return "Thurgau"
	case SubdivisionCHTI:
		return "Ticino"
	case SubdivisionCHUR:
		return "Uri"
	case SubdivisionCHVD:
		return "Vaud"
	case SubdivisionCHVS:
		return "Valais"
	case SubdivisionCHZG:
		return "Zug"
	case SubdivisionCHZH:
		return "Zürich"
	case SubdivisionCI01:
		return "Lagunes (Région des)"
	case SubdivisionCI02:
		return "Haut-Sassandra (Région du)"
	case SubdivisionCI03:
		return "Savanes (Région des)"
	case SubdivisionCI04:
		return "Vallée du Bandama (Région de la)"
	case SubdivisionCI05:
		return "Moyen-Comoé (Région du)"
	case SubdivisionCI06:
		return "18 Montagnes (Région des)"
	case SubdivisionCI07:
		return "Lacs (Région des)"
	case SubdivisionCI08:
		return "Zanzan (Région du)"
	case SubdivisionCI09:
		return "Bas-Sassandra (Région du)"
	case SubdivisionCI10:
		return "Denguélé (Région du)"
	case SubdivisionCI11:
		return "Nzi-Comoé (Région)"
	case SubdivisionCI12:
		return "Marahoué (Région de la)"
	case SubdivisionCI13:
		return "Sud-Comoé (Région du)"
	case SubdivisionCI14:
		return "Worodouqou (Région du)"
	case SubdivisionCI15:
		return "Sud-Bandama (Région du)"
	case SubdivisionCI16:
		return "Agnébi (Région de l')"
	case SubdivisionCI17:
		return "Bafing (Région du)"
	case SubdivisionCI18:
		return "Fromager (Région du)"
	case SubdivisionCI19:
		return "Moyen-Cavally (Région du)"
	case SubdivisionCLAI:
		return "Aisén del General Carlos Ibáñez del Campo"
	case SubdivisionCLAN:
		return "Antofagasta"
	case SubdivisionCLAP:
		return "Arica y Parinacota"
	case SubdivisionCLAR:
		return "Araucanía"
	case SubdivisionCLAT:
		return "Atacama"
	case SubdivisionCLBI:
		return "Bío-Bío"
	case SubdivisionCLCO:
		return "Coquimbo"
	case SubdivisionCLLI:
		return "Libertador General Bernardo O'Higgins"
	case SubdivisionCLLL:
		return "Los Lagos"
	case SubdivisionCLLR:
		return "Los Ríos"
	case SubdivisionCLMA:
		return "Magallanes y Antártica Chilena"
	case SubdivisionCLML:
		return "Maule"
	case SubdivisionCLRM:
		return "Región Metropolitana de Santiago"
	case SubdivisionCLTA:
		return "Tarapacá"
	case SubdivisionCLVS:
		return "Valparaíso"
	case SubdivisionCMAD:
		return "Adamaoua"
	case SubdivisionCMCE:
		return "Centre"
	case SubdivisionCMEN:
		return "Far North"
	case SubdivisionCMES:
		return "East"
	case SubdivisionCMLT:
		return "Littoral"
	case SubdivisionCMNO:
		return "North"
	case SubdivisionCMNW:
		return "North-West (Cameroon)"
	case SubdivisionCMOU:
		return "West"
	case SubdivisionCMSU:
		return "South"
	case SubdivisionCMSW:
		return "South-West"
	case SubdivisionCNAH:
		return "Anhui Sheng"
	case SubdivisionCNBJ:
		return "Beijing Shi"
	case SubdivisionCNCQ:
		return "Chongqing Shi"
	case SubdivisionCNFJ:
		return "Fujian Sheng"
	case SubdivisionCNGD:
		return "Guangdong Sheng"
	case SubdivisionCNGS:
		return "Gansu Sheng"
	case SubdivisionCNGX:
		return "Guangxi Zhuangzu Zizhiqu"
	case SubdivisionCNGZ:
		return "Guizhou Sheng"
	case SubdivisionCNHA:
		return "Henan Sheng"
	case SubdivisionCNHB:
		return "Hubei Sheng"
	case SubdivisionCNHE:
		return "Hebei Sheng"
	case SubdivisionCNHI:
		return "Hainan Sheng"
	case SubdivisionCNHK:
		return "Hong Kong SAR (see also separate country code entry under HK)"
	case SubdivisionCNHL:
		return "Heilongjiang Sheng"
	case SubdivisionCNHN:
		return "Hunan Sheng"
	case SubdivisionCNJL:
		return "Jilin Sheng"
	case SubdivisionCNJS:
		return "Jiangsu Sheng"
	case SubdivisionCNJX:
		return "Jiangxi Sheng"
	case SubdivisionCNLN:
		return "Liaoning Sheng"
	case SubdivisionCNMO:
		return "Macao SAR (see also separate country code entry under MO)"
	case SubdivisionCNNM:
		return "Nei Mongol Zizhiqu"
	case SubdivisionCNNX:
		return "Ningxia Huizi Zizhiqu"
	case SubdivisionCNQH:
		return "Qinghai Sheng"
	case SubdivisionCNSC:
		return "Sichuan Sheng"
	case SubdivisionCNSD:
		return "Shandong Sheng"
	case SubdivisionCNSH:
		return "Shanghai Shi"
	case SubdivisionCNSN:
		return "Shaanxi Sheng"
	case SubdivisionCNSX:
		return "Shanxi Sheng"
	case SubdivisionCNTJ:
		return "Tianjin Shi"
	case SubdivisionCNTW:
		return "Taiwan Sheng (see also separate country code entry under TW)"
	case SubdivisionCNXJ:
		return "Xinjiang Uygur Zizhiqu"
	case SubdivisionCNXZ:
		return "Xizang Zizhiqu"
	case SubdivisionCNYN:
		return "Yunnan Sheng"
	case SubdivisionCNZJ:
		return "Zhejiang Sheng"
	case SubdivisionCOAMA:
		return "Amazonas"
	case SubdivisionCOANT:
		return "Antioquia"
	case SubdivisionCOARA:
		return "Arauca"
	case SubdivisionCOATL:
		return "Atlántico"
	case SubdivisionCOBOL:
		return "Bolívar"
	case SubdivisionCOBOY:
		return "Boyacá"
	case SubdivisionCOCAL:
		return "Caldas"
	case SubdivisionCOCAQ:
		return "Caquetá"
	case SubdivisionCOCAS:
		return "Casanare"
	case SubdivisionCOCAU:
		return "Cauca"
	case SubdivisionCOCES:
		return "Cesar"
	case SubdivisionCOCHO:
		return "Chocó"
	case SubdivisionCOCOR:
		return "Córdoba"
	case SubdivisionCOCUN:
		return "Cundinamarca"
	case SubdivisionCODC:
		return "Distrito Capital de Bogotá"
	case SubdivisionCOGUA:
		return "Guainía"
	case SubdivisionCOGUV:
		return "Guaviare"
	case SubdivisionCOHUI:
		return "Huila"
	case SubdivisionCOLAG:
		return "La Guajira"
	case SubdivisionCOMAG:
		return "Magdalena"
	case SubdivisionCOMET:
		return "Meta"
	case SubdivisionCONAR:
		return "Nariño"
	case SubdivisionCONSA:
		return "Norte de Santander"
	case SubdivisionCOPUT:
		return "Putumayo"
	case SubdivisionCOQUI:
		return "Quindío"
	case SubdivisionCORIS:
		return "Risaralda"
	case SubdivisionCOSAN:
		return "Santander"
	case SubdivisionCOSAP:
		return "San Andrés, Providencia y Santa Catalina"
	case SubdivisionCOSUC:
		return "Sucre"
	case SubdivisionCOTOL:
		return "Tolima"
	case SubdivisionCOVAC:
		return "Valle del Cauca"
	case SubdivisionCOVAU:
		return "Vaupés"
	case SubdivisionCOVID:
		return "Vichada"
	case SubdivisionCRA:
		return "Alajuela"
	case SubdivisionCRC:
		return "Cartago"
	case SubdivisionCRG:
		return "Guanacaste"
	case SubdivisionCRH:
		return "Heredia"
	case SubdivisionCRL:
		return "Limón"
	case SubdivisionCRP:
		return "Puntarenas"
	case SubdivisionCRSJ:
		return "San José"
	case SubdivisionCU01:
		return "Pinar del Rio"
	case SubdivisionCU02:
		return "La Habana"
	case SubdivisionCU03:
		return "Ciudad de La Habana"
	case SubdivisionCU04:
		return "Matanzas"
	case SubdivisionCU05:
		return "Villa Clara"
	case SubdivisionCU06:
		return "Cienfuegos"
	case SubdivisionCU07:
		return "Sancti Spíritus"
	case SubdivisionCU08:
		return "Ciego de Ávila"
	case SubdivisionCU09:
		return "Camagüey"
	case SubdivisionCU10:
		return "Las Tunas"
	case SubdivisionCU11:
		return "Holguín"
	case SubdivisionCU12:
		return "Granma"
	case SubdivisionCU13:
		return "Santiago de Cuba"
	case SubdivisionCU14:
		return "Guantánamo"
	case SubdivisionCU99:
		return "Isla de la Juventud"
	case SubdivisionCVB:
		return "Ilhas de Barlavento"
	case SubdivisionCVBR:
		return "Brava"
	case SubdivisionCVBV:
		return "Boa Vista"
	case SubdivisionCVCA:
		return "Santa Catarina"
	case SubdivisionCVCF:
		return "Santa Catarina de Fogo"
	case SubdivisionCVCR:
		return "Santa Cruz"
	case SubdivisionCVMA:
		return "Maio"
	case SubdivisionCVMO:
		return "Mosteiros"
	case SubdivisionCVPA:
		return "Paul"
	case SubdivisionCVPN:
		return "Porto Novo"
	case SubdivisionCVPR:
		return "Praia"
	case SubdivisionCVRB:
		return "Ribeira Brava"
	case SubdivisionCVRG:
		return "Ribeira Grande"
	case SubdivisionCVRS:
		return "Ribeira Grande de Santiago"
	case SubdivisionCVS:
		return "Ilhas de Sotavento"
	case SubdivisionCVSD:
		return "São Domingos"
	case SubdivisionCVSF:
		return "São Filipe"
	case SubdivisionCVSL:
		return "Sal"
	case SubdivisionCVSM:
		return "São Miguel"
	case SubdivisionCVSO:
		return "São Lourenço dos Órgãos"
	case SubdivisionCVSS:
		return "São Salvador do Mundo"
	case SubdivisionCVSV:
		return "São Vicente"
	case SubdivisionCVTA:
		return "Tarrafal"
	case SubdivisionCVTS:
		return "Tarrafal de São Nicolau"
	case SubdivisionCY01:
		return "Lefkosía"
	case SubdivisionCY02:
		return "Lemesós"
	case SubdivisionCY03:
		return "Lárnaka"
	case SubdivisionCY04:
		return "Ammóchostos"
	case SubdivisionCY05:
		return "Páfos"
	case SubdivisionCY06:
		return "Kerýneia"
	case SubdivisionCZ10:
		return "Praha, Hlavní mešto"
	case SubdivisionCZ101:
		return "Praha 1"
	case SubdivisionCZ102:
		return "Praha 2"
	case SubdivisionCZ103:
		return "Praha 3"
	case SubdivisionCZ104:
		return "Praha 4"
	case SubdivisionCZ105:
		return "Praha 5"
	case SubdivisionCZ106:
		return "Praha 6"
	case SubdivisionCZ107:
		return "Praha 7"
	case SubdivisionCZ108:
		return "Praha 8"
	case SubdivisionCZ109:
		return "Praha 9"
	case SubdivisionCZ110:
		return "Praha 10"
	case SubdivisionCZ111:
		return "Praha 11"
	case SubdivisionCZ112:
		return "Praha 12"
	case SubdivisionCZ113:
		return "Praha 13"
	case SubdivisionCZ114:
		return "Praha 14"
	case SubdivisionCZ115:
		return "Praha 15"
	case SubdivisionCZ116:
		return "Praha 16"
	case SubdivisionCZ117:
		return "Praha 17"
	case SubdivisionCZ118:
		return "Praha 18"
	case SubdivisionCZ119:
		return "Praha 19"
	case SubdivisionCZ120:
		return "Praha 20"
	case SubdivisionCZ121:
		return "Praha 21"
	case SubdivisionCZ122:
		return "Praha 22"
	case SubdivisionCZ20:
		return "Středočeský kraj"
	case SubdivisionCZ201:
		return "Benešov"
	case SubdivisionCZ202:
		return "Beroun"
	case SubdivisionCZ203:
		return "Kladno"
	case SubdivisionCZ204:
		return "Kolín"
	case SubdivisionCZ205:
		return "Kutná Hora"
	case SubdivisionCZ206:
		return "Mělník"
	case SubdivisionCZ207:
		return "Mladá Boleslav"
	case SubdivisionCZ208:
		return "Nymburk"
	case SubdivisionCZ209:
		return "Praha-východ"
	case SubdivisionCZ20A:
		return "Praha-západ"
	case SubdivisionCZ20B:
		return "Příbram"
	case SubdivisionCZ20C:
		return "Rakovník"
	case SubdivisionCZ31:
		return "Jihočeský kraj"
	case SubdivisionCZ311:
		return "České Budějovice"
	case SubdivisionCZ312:
		return "Český Krumlov"
	case SubdivisionCZ313:
		return "Jindřichův Hradec"
	case SubdivisionCZ314:
		return "Písek"
	case SubdivisionCZ315:
		return "Prachatice"
	case SubdivisionCZ316:
		return "Strakonice"
	case SubdivisionCZ317:
		return "Tábor"
	case SubdivisionCZ32:
		return "Plzeňský kraj"
	case SubdivisionCZ321:
		return "Domažlice"
	case SubdivisionCZ322:
		return "Klatovy"
	case SubdivisionCZ323:
		return "Plzeň-město"
	case SubdivisionCZ324:
		return "Plzeň-jih"
	case SubdivisionCZ325:
		return "Plzeň-sever"
	case SubdivisionCZ326:
		return "Rokycany"
	case SubdivisionCZ327:
		return "Tachov"
	case SubdivisionCZ41:
		return "Karlovarský kraj"
	case SubdivisionCZ411:
		return "Cheb"
	case SubdivisionCZ412:
		return "Karlovy Vary"
	case SubdivisionCZ413:
		return "Sokolov"
	case SubdivisionCZ42:
		return "Ústecký kraj"
	case SubdivisionCZ421:
		return "Děčín"
	case SubdivisionCZ422:
		return "Chomutov"
	case SubdivisionCZ423:
		return "Litoměřice"
	case SubdivisionCZ424:
		return "Louny"
	case SubdivisionCZ425:
		return "Most"
	case SubdivisionCZ426:
		return "Teplice"
	case SubdivisionCZ427:
		return "Ústí nad Labem"
	case SubdivisionCZ51:
		return "Liberecký kraj"
	case SubdivisionCZ511:
		return "Česká Lípa"
	case SubdivisionCZ512:
		return "Jablonec nad Nisou"
	case SubdivisionCZ513:
		return "Liberec"
	case SubdivisionCZ514:
		return "Semily"
	case SubdivisionCZ52:
		return "Královéhradecký kraj"
	case SubdivisionCZ521:
		return "Hradec Králové"
	case SubdivisionCZ522:
		return "Jičín"
	case SubdivisionCZ523:
		return "Náchod"
	case SubdivisionCZ524:
		return "Rychnov nad Kněžnou"
	case SubdivisionCZ525:
		return "Trutnov"
	case SubdivisionCZ53:
		return "Pardubický kraj"
	case SubdivisionCZ531:
		return "Chrudim"
	case SubdivisionCZ532:
		return "Pardubice"
	case SubdivisionCZ533:
		return "Svitavy"
	case SubdivisionCZ534:
		return "Ústí nad Orlicí"
	case SubdivisionCZ63:
		return "Kraj Vysočina"
	case SubdivisionCZ631:
		return "Havlíčkův Brod"
	case SubdivisionCZ632:
		return "Jihlava"
	case SubdivisionCZ633:
		return "Pelhřimov"
	case SubdivisionCZ634:
		return "Třebíč"
	case SubdivisionCZ635:
		return "Žďár nad Sázavou"
	case SubdivisionCZ64:
		return "Jihomoravský kraj"
	case SubdivisionCZ641:
		return "Blansko"
	case SubdivisionCZ642:
		return "Brno-město"
	case SubdivisionCZ643:
		return "Brno-venkov"
	case SubdivisionCZ644:
		return "Břeclav"
	case SubdivisionCZ645:
		return "Hodonín"
	case SubdivisionCZ646:
		return "Vyškov"
	case SubdivisionCZ647:
		return "Znojmo"
	case SubdivisionCZ71:
		return "Olomoucký kraj"
	case SubdivisionCZ711:
		return "Jeseník"
	case SubdivisionCZ712:
		return "Olomouc"
	case SubdivisionCZ713:
		return "Prostějov"
	case SubdivisionCZ714:
		return "Přerov"
	case SubdivisionCZ715:
		return "Šumperk"
	case SubdivisionCZ72:
		return "Zlínský kraj"
	case SubdivisionCZ721:
		return "Kroměříž"
	case SubdivisionCZ722:
		return "Uherské Hradiště"
	case SubdivisionCZ723:
		return "Vsetín"
	case SubdivisionCZ724:
		return "Zlín"
	case SubdivisionCZ80:
		return "Moravskoslezský kraj"
	case SubdivisionCZ801:
		return "Bruntál"
	case SubdivisionCZ802:
		return "Frýdek Místek"
	case SubdivisionCZ803:
		return "Karviná"
	case SubdivisionCZ804:
		return "Nový Jičín"
	case SubdivisionCZ805:
		return "Opava"
	case SubdivisionCZ806:
		return "Ostrava-město"
	case SubdivisionDEBB:
		return "Brandenburg"
	case SubdivisionDEBE:
		return "Berlin"
	case SubdivisionDEBW:
		return "Baden-Württemberg"
	case SubdivisionDEBY:
		return "Bayern"
	case SubdivisionDEHB:
		return "Bremen"
	case SubdivisionDEHE:
		return "Hessen"
	case SubdivisionDEHH:
		return "Hamburg"
	case SubdivisionDEMV:
		return "Mecklenburg-Vorpommern"
	case SubdivisionDENI:
		return "Niedersachsen"
	case SubdivisionDENW:
		return "Nordrhein-Westfalen"
	case SubdivisionDERP:
		return "Rheinland-Pfalz"
	case SubdivisionDESH:
		return "Schleswig-Holstein"
	case SubdivisionDESL:
		return "Saarland"
	case SubdivisionDESN:
		return "Sachsen"
	case SubdivisionDEST:
		return "Sachsen-Anhalt"
	case SubdivisionDETH:
		return "Thüringen"
	case SubdivisionDJAR:
		return "Arta"
	case SubdivisionDJAS:
		return "Ali Sabieh"
	case SubdivisionDJDI:
		return "Dikhil"
	case SubdivisionDJDJ:
		return "Djibouti"
	case SubdivisionDJOB:
		return "Obock"
	case SubdivisionDJTA:
		return "Tadjourah"
	case SubdivisionDK81:
		return "Nordjylland"
	case SubdivisionDK82:
		return "Midtjylland"
	case SubdivisionDK83:
		return "Syddanmark"
	case SubdivisionDK84:
		return "Hovedstaden"
	case SubdivisionDK85:
		return "Sjælland"
	case SubdivisionDM01:
		return "Saint Peter"
	case SubdivisionDM02:
		return "Saint Andrew"
	case SubdivisionDM03:
		return "Saint David"
	case SubdivisionDM04:
		return "Saint George"
	case SubdivisionDM05:
		return "Saint John"
	case SubdivisionDM06:
		return "Saint Joseph"
	case SubdivisionDM07:
		return "Saint Luke"
	case SubdivisionDM08:
		return "Saint Mark"
	case SubdivisionDM09:
		return "Saint Patrick"
	case SubdivisionDM10:
		return "Saint Paul"
	case SubdivisionDO01:
		return "Distrito Nacional (Santo Domingo)"
	case SubdivisionDO02:
		return "Azua"
	case SubdivisionDO03:
		return "Bahoruco"
	case SubdivisionDO04:
		return "Barahona"
	case SubdivisionDO05:
		return "Dajabón"
	case SubdivisionDO06:
		return "Duarte"
	case SubdivisionDO07:
		return "La Estrelleta [Elías Piña]"
	case SubdivisionDO08:
		return "El Seybo [El Seibo]"
	case SubdivisionDO09:
		return "Espaillat"
	case SubdivisionDO10:
		return "Independencia"
	case SubdivisionDO11:
		return "La Altagracia"
	case SubdivisionDO12:
		return "La Romana"
	case SubdivisionDO13:
		return "La Vega"
	case SubdivisionDO14:
		return "María Trinidad Sánchez"
	case SubdivisionDO15:
		return "Monte Cristi"
	case SubdivisionDO16:
		return "Pedernales"
	case SubdivisionDO17:
		return "Peravia"
	case SubdivisionDO18:
		return "Puerto Plata"
	case SubdivisionDO19:
		return "Salcedo"
	case SubdivisionDO20:
		return "Samaná"
	case SubdivisionDO21:
		return "San Cristóbal"
	case SubdivisionDO22:
		return "San Juan"
	case SubdivisionDO23:
		return "San Pedro de Macorís"
	case SubdivisionDO24:
		return "Sánchez Ramírez"
	case SubdivisionDO25:
		return "Santiago"
	case SubdivisionDO26:
		return "Santiago Rodríguez"
	case SubdivisionDO27:
		return "Valverde"
	case SubdivisionDO28:
		return "Monseñor Nouel"
	case SubdivisionDO29:
		return "Monte Plata"
	case SubdivisionDO30:
		return "Hato Mayor"
	case SubdivisionDZ01:
		return "Adrar"
	case SubdivisionDZ02:
		return "Chlef"
	case SubdivisionDZ03:
		return "Laghouat"
	case SubdivisionDZ04:
		return "Oum el Bouaghi"
	case SubdivisionDZ05:
		return "Batna"
	case SubdivisionDZ06:
		return "Béjaïa"
	case SubdivisionDZ07:
		return "Biskra"
	case SubdivisionDZ08:
		return "Béchar"
	case SubdivisionDZ09:
		return "Blida"
	case SubdivisionDZ10:
		return "Bouira"
	case SubdivisionDZ11:
		return "Tamanghasset"
	case SubdivisionDZ12:
		return "Tébessa"
	case SubdivisionDZ13:
		return "Tlemcen"
	case SubdivisionDZ14:
		return "Tiaret"
	case SubdivisionDZ15:
		return "Tizi Ouzou"
	case SubdivisionDZ16:
		return "Alger"
	case SubdivisionDZ17:
		return "Djelfa"
	case SubdivisionDZ18:
		return "Jijel"
	case SubdivisionDZ19:
		return "Sétif"
	case SubdivisionDZ20:
		return "Saïda"
	case SubdivisionDZ21:
		return "Skikda"
	case SubdivisionDZ22:
		return "Sidi Bel Abbès"
	case SubdivisionDZ23:
		return "Annaba"
	case SubdivisionDZ24:
		return "Guelma"
	case SubdivisionDZ25:
		return "Constantine"
	case SubdivisionDZ26:
		return "Médéa"
	case SubdivisionDZ27:
		return "Mostaganem"
	case SubdivisionDZ28:
		return "Msila"
	case SubdivisionDZ29:
		return "Mascara"
	case SubdivisionDZ30:
		return "Ouargla"
	case SubdivisionDZ31:
		return "Oran"
	case SubdivisionDZ32:
		return "El Bayadh"
	case SubdivisionDZ33:
		return "Illizi"
	case SubdivisionDZ34:
		return "Bordj Bou Arréridj"
	case SubdivisionDZ35:
		return "Boumerdès"
	case SubdivisionDZ36:
		return "El Tarf"
	case SubdivisionDZ37:
		return "Tindouf"
	case SubdivisionDZ38:
		return "Tissemsilt"
	case SubdivisionDZ39:
		return "El Oued"
	case SubdivisionDZ40:
		return "Khenchela"
	case SubdivisionDZ41:
		return "Souk Ahras"
	case SubdivisionDZ42:
		return "Tipaza"
	case SubdivisionDZ43:
		return "Mila"
	case SubdivisionDZ44:
		return "Aïn Defla"
	case SubdivisionDZ45:
		return "Naama"
	case SubdivisionDZ46:
		return "Aïn Témouchent"
	case SubdivisionDZ47:
		return "Ghardaïa"
	case SubdivisionDZ48:
		return "Relizane"
	case SubdivisionECA:
		return "Azuay"
	case SubdivisionECB:
		return "Bolívar"
	case SubdivisionECC:
		return "Carchi"
	case SubdivisionECD:
		return "Orellana"
	case SubdivisionECE:
		return "Esmeraldas"
	case SubdivisionECF:
		return "Cañar"
	case SubdivisionECG:
		return "Guayas"
	case SubdivisionECH:
		return "Chimborazo"
	case SubdivisionECI:
		return "Imbabura"
	case SubdivisionECL:
		return "Loja"
	case SubdivisionECM:
		return "Manabí"
	case SubdivisionECN:
		return "Napo"
	case SubdivisionECO:
		return "El Oro"
	case SubdivisionECP:
		return "Pichincha"
	case SubdivisionECR:
		return "Los Ríos"
	case SubdivisionECS:
		return "Morona-Santiago"
	case SubdivisionECSD:
		return "Santo Domingo de los Tsáchilas"
	case SubdivisionECSE:
		return "Santa Elena"
	case SubdivisionECT:
		return "Tungurahua"
	case SubdivisionECU:
		return "Sucumbíos"
	case SubdivisionECW:
		return "Galápagos"
	case SubdivisionECX:
		return "Cotopaxi"
	case SubdivisionECY:
		return "Pastaza"
	case SubdivisionECZ:
		return "Zamora-Chinchipe"
	case SubdivisionEE37:
		return "Harjumaa"
	case SubdivisionEE39:
		return "Hiiumaa"
	case SubdivisionEE44:
		return "Ida-Virumaa"
	case SubdivisionEE49:
		return "Jõgevamaa"
	case SubdivisionEE51:
		return "Järvamaa"
	case SubdivisionEE57:
		return "Läänemaa"
	case SubdivisionEE59:
		return "Lääne-Virumaa"
	case SubdivisionEE65:
		return "Põlvamaa"
	case SubdivisionEE67:
		return "Pärnumaa"
	case SubdivisionEE70:
		return "Raplamaa"
	case SubdivisionEE74:
		return "Saaremaa"
	case SubdivisionEE78:
		return "Tartumaa"
	case SubdivisionEE82:
		return "Valgamaa"
	case SubdivisionEE84:
		return "Viljandimaa"
	case SubdivisionEE86:
		return "Võrumaa"
	case SubdivisionEGALX:
		return "Al Iskandarīyah"
	case SubdivisionEGASN:
		return "Aswān"
	case SubdivisionEGAST:
		return "Asyūt"
	case SubdivisionEGBA:
		return "Al Bahr al Ahmar"
	case SubdivisionEGBH:
		return "Al Buhayrah"
	case SubdivisionEGBNS:
		return "Banī Suwayf"
	case SubdivisionEGC:
		return "Al Qāhirah"
	case SubdivisionEGDK:
		return "Ad Daqahlīyah"
	case SubdivisionEGDT:
		return "Dumyāt"
	case SubdivisionEGFYM:
		return "Al Fayyūm"
	case SubdivisionEGGH:
		return "Al Gharbīyah"
	case SubdivisionEGGZ:
		return "Al Jīzah"
	case SubdivisionEGHU:
		return "Ḩulwān"
	case SubdivisionEGIS:
		return "Al Ismā`īlīyah"
	case SubdivisionEGJS:
		return "Janūb Sīnā'"
	case SubdivisionEGKB:
		return "Al Qalyūbīyah"
	case SubdivisionEGKFS:
		return "Kafr ash Shaykh"
	case SubdivisionEGKN:
		return "Qinā"
	case SubdivisionEGMN:
		return "Al Minyā"
	case SubdivisionEGMNF:
		return "Al Minūfīyah"
	case SubdivisionEGMT:
		return "Matrūh"
	case SubdivisionEGPTS:
		return "Būr Sa`īd"
	case SubdivisionEGSHG:
		return "Sūhāj"
	case SubdivisionEGSHR:
		return "Ash Sharqīyah"
	case SubdivisionEGSIN:
		return "Shamal Sīnā'"
	case SubdivisionEGSU:
		return "As Sādis min Uktūbar"
	case SubdivisionEGSUZ:
		return "As Suways"
	case SubdivisionEGWAD:
		return "Al Wādī al Jadīd"
	case SubdivisionERAN:
		return "Ansabā"
	case SubdivisionERDK:
		return "Janūbī al Baḩrī al Aḩmar"
	case SubdivisionERDU:
		return "Al Janūbī"
	case SubdivisionERGB:
		return "Qāsh-Barkah"
	case SubdivisionERMA:
		return "Al Awsaţ"
	case SubdivisionERSK:
		return "Shimālī al Baḩrī al Aḩmar"
	case SubdivisionESA:
		return "Alicante"
	case SubdivisionESAB:
		return "Albacete"
	case SubdivisionESAL:
		return "Almería"
	case SubdivisionESAN:
		return "Andalucía"
	case SubdivisionESAR:
		return "Aragón"
	case SubdivisionESAS:
		return "Asturias, Principado de"
	case SubdivisionESAV:
		return "Ávila"
	case SubdivisionESB:
		return "Barcelona"
	case SubdivisionESBA:
		return "Badajoz"
	case SubdivisionESBI:
		return "Bizkaia"
	case SubdivisionESBU:
		return "Burgos"
	case SubdivisionESC:
		return "A Coruña"
	case SubdivisionESCA:
		return "Cádiz"
	case SubdivisionESCB:
		return "Cantabria"
	case SubdivisionESCC:
		return "Cáceres"
	case SubdivisionESCE:
		return "Ceuta"
	case SubdivisionESCL:
		return "Castilla y León"
	case SubdivisionESCM:
		return "Castilla-La Mancha"
	case SubdivisionESCN:
		return "Canarias"
	case SubdivisionESCO:
		return "Córdoba"
	case SubdivisionESCR:
		return "Ciudad Real"
	case SubdivisionESCS:
		return "Castellón"
	case SubdivisionESCT:
		return "Catalunya"
	case SubdivisionESCU:
		return "Cuenca"
	case SubdivisionESEX:
		return "Extremadura"
	case SubdivisionESGA:
		return "Galicia"
	case SubdivisionESGC:
		return "Las Palmas"
	case SubdivisionESGI:
		return "Girona"
	case SubdivisionESGR:
		return "Granada"
	case SubdivisionESGU:
		return "Guadalajara"
	case SubdivisionESH:
		return "Huelva"
	case SubdivisionESHU:
		return "Huesca"
	case SubdivisionESIB:
		return "Illes Balears"
	case SubdivisionESJ:
		return "Jaén"
	case SubdivisionESL:
		return "Lleida"
	case SubdivisionESLE:
		return "León"
	case SubdivisionESLO:
		return "La Rioja"
	case SubdivisionESLU:
		return "Lugo"
	case SubdivisionESM:
		return "Madrid"
	case SubdivisionESMA:
		return "Málaga"
	case SubdivisionESMC:
		return "Murcia, Región de"
	case SubdivisionESMD:
		return "Madrid, Comunidad de"
	case SubdivisionESML:
		return "Melilla"
	case SubdivisionESMU:
		return "Murcia"
	case SubdivisionESNA:
		return "Navarra / Nafarroa"
	case SubdivisionESNC:
		return "Navarra, Comunidad Foral de / Nafarroako Foru Komunitatea"
	case SubdivisionESO:
		return "Asturias"
	case SubdivisionESOR:
		return "Ourense"
	case SubdivisionESP:
		return "Palencia"
	case SubdivisionESPM:
		return "Balears"
	case SubdivisionESPO:
		return "Pontevedra"
	case SubdivisionESPV:
		return "País Vasco / Euskal Herria"
	case SubdivisionESRI:
		return "La Rioja"
	case SubdivisionESS:
		return "Cantabria"
	case SubdivisionESSA:
		return "Salamanca"
	case SubdivisionESSE:
		return "Sevilla"
	case SubdivisionESSG:
		return "Segovia"
	case SubdivisionESSO:
		return "Soria"
	case SubdivisionESSS:
		return "Gipuzkoa"
	case SubdivisionEST:
		return "Tarragona"
	case SubdivisionESTE:
		return "Teruel"
	case SubdivisionESTF:
		return "Santa Cruz de Tenerife"
	case SubdivisionESTO:
		return "Toledo"
	case SubdivisionESV:
		return "Valencia / València"
	case SubdivisionESVA:
		return "Valladolid"
	case SubdivisionESVC:
		return "Valenciana, Comunidad / Valenciana, Comunitat"
	case SubdivisionESVI:
		return "Álava"
	case SubdivisionESZ:
		return "Zaragoza"
	case SubdivisionESZA:
		return "Zamora"
	case SubdivisionETAA:
		return "Ādīs Ābeba"
	case SubdivisionETAF:
		return "Āfar"
	case SubdivisionETAM:
		return "Āmara"
	case SubdivisionETBE:
		return "Bīnshangul Gumuz"
	case SubdivisionETDD:
		return "Dirē Dawa"
	case SubdivisionETGA:
		return "Gambēla Hizboch"
	case SubdivisionETHA:
		return "Hārerī Hizb"
	case SubdivisionETOR:
		return "Oromīya"
	case SubdivisionETSN:
		return "YeDebub Bihēroch Bihēreseboch na Hizboch"
	case SubdivisionETSO:
		return "Sumalē"
	case SubdivisionETTI:
		return "Tigray"
	case SubdivisionFI01:
		return "Ahvenanmaan maakunta"
	case SubdivisionFI02:
		return "Etelä-Karjala"
	case SubdivisionFI03:
		return "Etelä-Pohjanmaa"
	case SubdivisionFI04:
		return "Etelä-Savo"
	case SubdivisionFI05:
		return "Kainuu"
	case SubdivisionFI06:
		return "Kanta-Häme"
	case SubdivisionFI07:
		return "Keski-Pohjanmaa"
	case SubdivisionFI08:
		return "Keski-Suomi"
	case SubdivisionFI09:
		return "Kymenlaakso"
	case SubdivisionFI10:
		return "Lappi"
	case SubdivisionFI11:
		return "Pirkanmaa"
	case SubdivisionFI12:
		return "Pohjanmaa"
	case SubdivisionFI13:
		return "Pohjois-Karjala"
	case SubdivisionFI14:
		return "Pohjois-Pohjanmaa"
	case SubdivisionFI15:
		return "Pohjois-Savo"
	case SubdivisionFI16:
		return "Päijät-Häme"
	case SubdivisionFI17:
		return "Satakunta"
	case SubdivisionFI18:
		return "Uusimaa"
	case SubdivisionFI19:
		return "Varsinais-Suomi"
	case SubdivisionFJC:
		return "Central"
	case SubdivisionFJE:
		return "Eastern"
	case SubdivisionFJN:
		return "Northern"
	case SubdivisionFJR:
		return "Rotuma"
	case SubdivisionFJW:
		return "Western"
	case SubdivisionFMKSA:
		return "Kosrae"
	case SubdivisionFMPNI:
		return "Pohnpei"
	case SubdivisionFMTRK:
		return "Chuuk"
	case SubdivisionFMYAP:
		return "Yap"
	case SubdivisionFR01:
		return "Ain"
	case SubdivisionFR02:
		return "Aisne"
	case SubdivisionFR03:
		return "Allier"
	case SubdivisionFR04:
		return "Alpes-de-Haute-Provence"
	case SubdivisionFR05:
		return "Hautes-Alpes"
	case SubdivisionFR06:
		return "Alpes-Maritimes"
	case SubdivisionFR07:
		return "Ardèche"
	case SubdivisionFR08:
		return "Ardennes"
	case SubdivisionFR09:
		return "Ariège"
	case SubdivisionFR10:
		return "Aube"
	case SubdivisionFR11:
		return "Aude"
	case SubdivisionFR12:
		return "Aveyron"
	case SubdivisionFR13:
		return "Bouches-du-Rhône"
	case SubdivisionFR14:
		return "Calvados"
	case SubdivisionFR15:
		return "Cantal"
	case SubdivisionFR16:
		return "Charente"
	case SubdivisionFR17:
		return "Charente-Maritime"
	case SubdivisionFR18:
		return "Cher"
	case SubdivisionFR19:
		return "Corrèze"
	case SubdivisionFR21:
		return "Côte-d'Or"
	case SubdivisionFR22:
		return "Côtes-d'Armor"
	case SubdivisionFR23:
		return "Creuse"
	case SubdivisionFR24:
		return "Dordogne"
	case SubdivisionFR25:
		return "Doubs"
	case SubdivisionFR26:
		return "Drôme"
	case SubdivisionFR27:
		return "Eure"
	case SubdivisionFR28:
		return "Eure-et-Loir"
	case SubdivisionFR29:
		return "Finistère"
	case SubdivisionFR2A:
		return "Corse-du-Sud"
	case SubdivisionFR2B:
		return "Haute-Corse"
	case SubdivisionFR30:
		return "Gard"
	case SubdivisionFR31:
		return "Haute-Garonne"
	case SubdivisionFR32:
		return "Gers"
	case SubdivisionFR33:
		return "Gironde"
	case SubdivisionFR34:
		return "Hérault"
	case SubdivisionFR35:
		return "Ille-et-Vilaine"
	case SubdivisionFR36:
		return "Indre"
	case SubdivisionFR37:
		return "Indre-et-Loire"
	case SubdivisionFR38:
		return "Isère"
	case SubdivisionFR39:
		return "Jura"
	case SubdivisionFR40:
		return "Landes"
	case SubdivisionFR41:
		return "Loir-et-Cher"
	case SubdivisionFR42:
		return "Loire"
	case SubdivisionFR43:
		return "Haute-Loire"
	case SubdivisionFR44:
		return "Loire-Atlantique"
	case SubdivisionFR45:
		return "Loiret"
	case SubdivisionFR46:
		return "Lot"
	case SubdivisionFR47:
		return "Lot-et-Garonne"
	case SubdivisionFR48:
		return "Lozère"
	case SubdivisionFR49:
		return "Maine-et-Loire"
	case SubdivisionFR50:
		return "Manche"
	case SubdivisionFR51:
		return "Marne"
	case SubdivisionFR52:
		return "Haute-Marne"
	case SubdivisionFR53:
		return "Mayenne"
	case SubdivisionFR54:
		return "Meurthe-et-Moselle"
	case SubdivisionFR55:
		return "Meuse"
	case SubdivisionFR56:
		return "Morbihan"
	case SubdivisionFR57:
		return "Moselle"
	case SubdivisionFR58:
		return "Nièvre"
	case SubdivisionFR59:
		return "Nord"
	case SubdivisionFR60:
		return "Oise"
	case SubdivisionFR61:
		return "Orne"
	case SubdivisionFR62:
		return "Pas-de-Calais"
	case SubdivisionFR63:
		return "Puy-de-Dôme"
	case SubdivisionFR64:
		return "Pyrénées-Atlantiques"
	case SubdivisionFR65:
		return "Hautes-Pyrénées"
	case SubdivisionFR66:
		return "Pyrénées-Orientales"
	case SubdivisionFR67:
		return "Bas-Rhin"
	case SubdivisionFR68:
		return "Haut-Rhin"
	case SubdivisionFR69:
		return "Rhône"
	case SubdivisionFR70:
		return "Haute-Saône"
	case SubdivisionFR71:
		return "Saône-et-Loire"
	case SubdivisionFR72:
		return "Sarthe"
	case SubdivisionFR73:
		return "Savoie"
	case SubdivisionFR74:
		return "Haute-Savoie"
	case SubdivisionFR75:
		return "Paris"
	case SubdivisionFR76:
		return "Seine-Maritime"
	case SubdivisionFR77:
		return "Seine-et-Marne"
	case SubdivisionFR78:
		return "Yvelines"
	case SubdivisionFR79:
		return "Deux-Sèvres"
	case SubdivisionFR80:
		return "Somme"
	case SubdivisionFR81:
		return "Tarn"
	case SubdivisionFR82:
		return "Tarn-et-Garonne"
	case SubdivisionFR83:
		return "Var"
	case SubdivisionFR84:
		return "Vaucluse"
	case SubdivisionFR85:
		return "Vendée"
	case SubdivisionFR86:
		return "Vienne"
	case SubdivisionFR87:
		return "Haute-Vienne"
	case SubdivisionFR88:
		return "Vosges"
	case SubdivisionFR89:
		return "Yonne"
	case SubdivisionFR90:
		return "Territoire de Belfort"
	case SubdivisionFR91:
		return "Essonne"
	case SubdivisionFR92:
		return "Hauts-de-Seine"
	case SubdivisionFR93:
		return "Seine-Saint-Denis"
	case SubdivisionFR94:
		return "Val-de-Marne"
	case SubdivisionFR95:
		return "Val-d'Oise"
	case SubdivisionFRARA:
		return "Auvergne-Rhône-Alpes"
	case SubdivisionFRBFC:
		return "Bourgogne-Franche-Comté"
	case SubdivisionFRBL:
		return "Saint-Barthélemy"
	case SubdivisionFRBRE:
		return "Bretagne"
	case SubdivisionFRCOR:
		return "Corse"
	case SubdivisionFRCP:
		return "Clipperton"
	case SubdivisionFRCVL:
		return "Centre-Val de Loire"
	case SubdivisionFRGES:
		return "Grand-Est"
	case SubdivisionFRGF:
		return "Guyane (française)"
	case SubdivisionFRGP:
		return "Guadeloupe"
	case SubdivisionFRGUA:
		return "Guadeloupe"
	case SubdivisionFRHDF:
		return "Hauts-de-France"
	case SubdivisionFRIDF:
		return "Île-de-France"
	case SubdivisionFRLRE:
		return "La Réunion"
	case SubdivisionFRMAY:
		return "Mayotte"
	case SubdivisionFRMF:
		return "Saint-Martin"
	case SubdivisionFRMQ:
		return "Martinique"
	case SubdivisionFRNAQ:
		return "Nouvelle-Aquitaine"
	case SubdivisionFRNC:
		return "Nouvelle-Calédonie"
	case SubdivisionFRNOR:
		return "Normandie"
	case SubdivisionFROCC:
		return "Occitanie"
	case SubdivisionFRPAC:
		return "Provence-Alpes-Côte-d’Azur"
	case SubdivisionFRPDL:
		return "Pays-de-la-Loire"
	case SubdivisionFRPF:
		return "Polynésie française"
	case SubdivisionFRPM:
		return "Saint-Pierre-et-Miquelon"
	case SubdivisionFRRE:
		return "La Réunion"
	case SubdivisionFRTF:
		return "Terres australes françaises"
	case SubdivisionFRWF:
		return "Wallis-et-Futuna"
	case SubdivisionFRYT:
		return "Mayotte"
	case SubdivisionGA1:
		return "Estuaire"
	case SubdivisionGA2:
		return "Haut-Ogooué"
	case SubdivisionGA3:
		return "Moyen-Ogooué"
	case SubdivisionGA4:
		return "Ngounié"
	case SubdivisionGA5:
		return "Nyanga"
	case SubdivisionGA6:
		return "Ogooué-Ivindo"
	case SubdivisionGA7:
		return "Ogooué-Lolo"
	case SubdivisionGA8:
		return "Ogooué-Maritime"
	case SubdivisionGA9:
		return "Woleu-Ntem"
	case SubdivisionGBABC:
		return "Armagh, Banbridge and Craigavon"
	case SubdivisionGBABD:
		return "Aberdeenshire"
	case SubdivisionGBABE:
		return "Aberdeen City"
	case SubdivisionGBAGB:
		return "Argyll and Bute"
	case SubdivisionGBAGY:
		return "Isle of Anglesey; Sir Ynys Môn"
	case SubdivisionGBAND:
		return "Ards and North Down"
	case SubdivisionGBANN:
		return "Antrim and Newtownabbey"
	case SubdivisionGBANS:
		return "Angus"
	case SubdivisionGBBAS:
		return "Bath and North East Somerset"
	case SubdivisionGBBBD:
		return "Blackburn with Darwen"
	case SubdivisionGBBDF:
		return "Bedford"
	case SubdivisionGBBDG:
		return "Barking and Dagenham"
	case SubdivisionGBBEN:
		return "Brent"
	case SubdivisionGBBEX:
		return "Bexley"
	case SubdivisionGBBFS:
		return "Belfast"
	case SubdivisionGBBGE:
		return "Bridgend; Pen-y-bont ar Ogwr"
	case SubdivisionGBBGW:
		return "Blaenau Gwent"
	case SubdivisionGBBIR:
		return "Birmingham"
	case SubdivisionGBBKM:
		return "Buckinghamshire"
	case SubdivisionGBBMH:
		return "Bournemouth"
	case SubdivisionGBBNE:
		return "Barnet"
	case SubdivisionGBBNH:
		return "Brighton and Hove"
	case SubdivisionGBBNS:
		return "Barnsley"
	case SubdivisionGBBOL:
		return "Bolton"
	case SubdivisionGBBPL:
		return "Blackpool"
	case SubdivisionGBBRC:
		return "Bracknell Forest"
	case SubdivisionGBBRD:
		return "Bradford"
	case SubdivisionGBBRY:
		return "Bromley"
	case SubdivisionGBBST:
		return "Bristol, City of"
	case SubdivisionGBBUR:
		return "Bury"
	case SubdivisionGBCAM:
		return "Cambridgeshire"
	case SubdivisionGBCAY:
		return "Caerphilly; Caerffili"
	case SubdivisionGBCBF:
		return "Central Bedfordshire"
	case SubdivisionGBCCG:
		return "Causeway Coast and Glens"
	case SubdivisionGBCGN:
		return "Ceredigion; Sir Ceredigion"
	case SubdivisionGBCHE:
		return "Cheshire East"
	case SubdivisionGBCHW:
		return "Cheshire West and Chester"
	case SubdivisionGBCLD:
		return "Calderdale"
	case SubdivisionGBCLK:
		return "Clackmannanshire"
	case SubdivisionGBCMA:
		return "Cumbria"
	case SubdivisionGBCMD:
		return "Camden"
	case SubdivisionGBCMN:
		return "Carmarthenshire; Sir Gaerfyrddin"
	case SubdivisionGBCON:
		return "Cornwall"
	case SubdivisionGBCOV:
		return "Coventry"
	case SubdivisionGBCRF:
		return "Cardiff; Caerdydd"
	case SubdivisionGBCRY:
		return "Croydon"
	case SubdivisionGBCWY:
		return "Conwy"
	case SubdivisionGBDAL:
		return "Darlington"
	case SubdivisionGBDBY:
		return "Derbyshire"
	case SubdivisionGBDEN:
		return "Denbighshire; Sir Ddinbych"
	case SubdivisionGBDER:
		return "Derby"
	case SubdivisionGBDEV:
		return "Devon"
	case SubdivisionGBDGY:
		return "Dumfries and Galloway"
	case SubdivisionGBDNC:
		return "Doncaster"
	case SubdivisionGBDND:
		return "Dundee City"
	case SubdivisionGBDOR:
		return "Dorset"
	case SubdivisionGBDRS:
		return "Derry and Strabane"
	case SubdivisionGBDUD:
		return "Dudley"
	case SubdivisionGBDUR:
		return "Durham County"
	case SubdivisionGBEAL:
		return "Ealing"
	case SubdivisionGBEAW:
		return "England and Wales"
	case SubdivisionGBEAY:
		return "East Ayrshire"
	case SubdivisionGBEDH:
		return "Edinburgh, City of"
	case SubdivisionGBEDU:
		return "East Dunbartonshire"
	case SubdivisionGBELN:
		return "East Lothian"
	case SubdivisionGBELS:
		return "Eilean Siar"
	case SubdivisionGBENF:
		return "Enfield"
	case SubdivisionGBENG:
		return "England"
	case SubdivisionGBERW:
		return "East Renfrewshire"
	case SubdivisionGBERY:
		return "East Riding of Yorkshire"
	case SubdivisionGBESS:
		return "Essex"
	case SubdivisionGBESX:
		return "East Sussex"
	case SubdivisionGBFAL:
		return "Falkirk"
	case SubdivisionGBFIF:
		return "Fife"
	case SubdivisionGBFLN:
		return "Flintshire; Sir y Fflint"
	case SubdivisionGBFMO:
		return "Fermanagh and Omagh"
	case SubdivisionGBGAT:
		return "Gateshead"
	case SubdivisionGBGBN:
		return "Great Britain"
	case SubdivisionGBGLG:
		return "Glasgow City"
	case SubdivisionGBGLS:
		return "Gloucestershire"
	case SubdivisionGBGRE:
		return "Greenwich"
	case SubdivisionGBGWN:
		return "Gwynedd"
	case SubdivisionGBHAL:
		return "Halton"
	case SubdivisionGBHAM:
		return "Hampshire"
	case SubdivisionGBHAV:
		return "Havering"
	case SubdivisionGBHCK:
		return "Hackney"
	case SubdivisionGBHEF:
		return "Herefordshire"
	case SubdivisionGBHIL:
		return "Hillingdon"
	case SubdivisionGBHLD:
		return "Highland"
	case SubdivisionGBHMF:
		return "Hammersmith and Fulham"
	case SubdivisionGBHNS:
		return "Hounslow"
	case SubdivisionGBHPL:
		return "Hartlepool"
	case SubdivisionGBHRT:
		return "Hertfordshire"
	case SubdivisionGBHRW:
		return "Harrow"
	case SubdivisionGBHRY:
		return "Haringey"
	case SubdivisionGBIOS:
		return "Isles of Scilly"
	case SubdivisionGBIOW:
		return "Isle of Wight"
	case SubdivisionGBISL:
		return "Islington"
	case SubdivisionGBIVC:
		return "Inverclyde"
	case SubdivisionGBKEC:
		return "Kensington and Chelsea"
	case SubdivisionGBKEN:
		return "Kent"
	case SubdivisionGBKHL:
		return "Kingston upon Hull"
	case SubdivisionGBKIR:
		return "Kirklees"
	case SubdivisionGBKTT:
		return "Kingston upon Thames"
	case SubdivisionGBKWL:
		return "Knowsley"
	case SubdivisionGBLAN:
		return "Lancashire"
	case SubdivisionGBLBC:
		return "Lisburn and Castlereagh"
	case SubdivisionGBLBH:
		return "Lambeth"
	case SubdivisionGBLCE:
		return "Leicester"
	case SubdivisionGBLDS:
		return "Leeds"
	case SubdivisionGBLEC:
		return "Leicestershire"
	case SubdivisionGBLEW:
		return "Lewisham"
	case SubdivisionGBLIN:
		return "Lincolnshire"
	case SubdivisionGBLIV:
		return "Liverpool"
	case SubdivisionGBLND:
		return "London, City of"
	case SubdivisionGBLUT:
		return "Luton"
	case SubdivisionGBMAN:
		return "Manchester"
	case SubdivisionGBMDB:
		return "Middlesbrough"
	case SubdivisionGBMDW:
		return "Medway"
	case SubdivisionGBMEA:
		return "Mid and East Antrim"
	case SubdivisionGBMIK:
		return "Milton Keynes"
	case SubdivisionGBMLN:
		return "Midlothian"
	case SubdivisionGBMON:
		return "Monmouthshire; Sir Fynwy"
	case SubdivisionGBMRT:
		return "Merton"
	case SubdivisionGBMRY:
		return "Moray"
	case SubdivisionGBMTY:
		return "Merthyr Tydfil; Merthyr Tudful"
	case SubdivisionGBMUL:
		return "Mid Ulster"
	case SubdivisionGBNAY:
		return "North Ayrshire"
	case SubdivisionGBNBL:
		return "Northumberland"
	case SubdivisionGBNEL:
		return "North East Lincolnshire"
	case SubdivisionGBNET:
		return "Newcastle upon Tyne"
	case SubdivisionGBNFK:
		return "Norfolk"
	case SubdivisionGBNGM:
		return "Nottingham"
	case SubdivisionGBNIR:
		return "Northern Ireland"
	case SubdivisionGBNLK:
		return "North Lanarkshire"
	case SubdivisionGBNLN:
		return "North Lincolnshire"
	case SubdivisionGBNMD:
		return "Newry, Mourne and Down"
	case SubdivisionGBNSM:
		return "North Somerset"
	case SubdivisionGBNTH:
		return "Northamptonshire"
	case SubdivisionGBNTL:
		return "Neath Port Talbot; Castell-nedd Port Talbot"
	case SubdivisionGBNTT:
		return "Nottinghamshire"
	case SubdivisionGBNTY:
		return "North Tyneside"
	case SubdivisionGBNWM:
		return "Newham"
	case SubdivisionGBNWP:
		return "Newport; Casnewydd"
	case SubdivisionGBNYK:
		return "North Yorkshire"
	case SubdivisionGBOLD:
		return "Oldham"
	case SubdivisionGBORK:
		return "Orkney Islands"
	case SubdivisionGBOXF:
		return "Oxfordshire"
	case SubdivisionGBPEM:
		return "Pembrokeshire; Sir Benfro"
	case SubdivisionGBPKN:
		return "Perth and Kinross"
	case SubdivisionGBPLY:
		return "Plymouth"
	case SubdivisionGBPOL:
		return "Poole"
	case SubdivisionGBPOR:
		return "Portsmouth"
	case SubdivisionGBPOW:
		return "Powys"
	case SubdivisionGBPTE:
		return "Peterborough"
	case SubdivisionGBRCC:
		return "Redcar and Cleveland"
	case SubdivisionGBRCH:
		return "Rochdale"
	case SubdivisionGBRCT:
		return "Rhondda, Cynon, Taff; Rhondda, Cynon, Taf"
	case SubdivisionGBRDB:
		return "Redbridge"
	case SubdivisionGBRDG:
		return "Reading"
	case SubdivisionGBRFW:
		return "Renfrewshire"
	case SubdivisionGBRIC:
		return "Richmond upon Thames"
	case SubdivisionGBROT:
		return "Rotherham"
	case SubdivisionGBRUT:
		return "Rutland"
	case SubdivisionGBSAW:
		return "Sandwell"
	case SubdivisionGBSAY:
		return "South Ayrshire"
	case SubdivisionGBSCB:
		return "Scottish Borders, The"
	case SubdivisionGBSCT:
		return "Scotland"
	case SubdivisionGBSFK:
		return "Suffolk"
	case SubdivisionGBSFT:
		return "Sefton"
	case SubdivisionGBSGC:
		return "South Gloucestershire"
	case SubdivisionGBSHF:
		return "Sheffield"
	case SubdivisionGBSHN:
		return "St. Helens"
	case SubdivisionGBSHR:
		return "Shropshire"
	case SubdivisionGBSKP:
		return "Stockport"
	case SubdivisionGBSLF:
		return "Salford"
	case SubdivisionGBSLG:
		return "Slough"
	case SubdivisionGBSLK:
		return "South Lanarkshire"
	case SubdivisionGBSND:
		return "Sunderland"
	case SubdivisionGBSOL:
		return "Solihull"
	case SubdivisionGBSOM:
		return "Somerset"
	case SubdivisionGBSOS:
		return "Southend-on-Sea"
	case SubdivisionGBSRY:
		return "Surrey"
	case SubdivisionGBSTE:
		return "Stoke-on-Trent"
	case SubdivisionGBSTG:
		return "Stirling"
	case SubdivisionGBSTH:
		return "Southampton"
	case SubdivisionGBSTN:
		return "Sutton"
	case SubdivisionGBSTS:
		return "Staffordshire"
	case SubdivisionGBSTT:
		return "Stockton-on-Tees"
	case SubdivisionGBSTY:
		return "South Tyneside"
	case SubdivisionGBSWA:
		return "Swansea; Abertawe"
	case SubdivisionGBSWD:
		return "Swindon"
	case SubdivisionGBSWK:
		return "Southwark"
	case SubdivisionGBTAM:
		return "Tameside"
	case SubdivisionGBTFW:
		return "Telford and Wrekin"
	case SubdivisionGBTHR:
		return "Thurrock"
	case SubdivisionGBTOB:
		return "Torbay"
	case SubdivisionGBTOF:
		return "Torfaen; Tor-faen"
	case SubdivisionGBTRF:
		return "Trafford"
	case SubdivisionGBTWH:
		return "Tower Hamlets"
	case SubdivisionGBUKM:
		return "United Kingdom"
	case SubdivisionGBVGL:
		return "Vale of Glamorgan, The; Bro Morgannwg"
	case SubdivisionGBWAR:
		return "Warwickshire"
	case SubdivisionGBWBK:
		return "West Berkshire"
	case SubdivisionGBWDU:
		return "West Dunbartonshire"
	case SubdivisionGBWFT:
		return "Waltham Forest"
	case SubdivisionGBWGN:
		return "Wigan"
	case SubdivisionGBWIL:
		return "Wiltshire"
	case SubdivisionGBWKF:
		return "Wakefield"
	case SubdivisionGBWLL:
		return "Walsall"
	case SubdivisionGBWLN:
		return "West Lothian"
	case SubdivisionGBWLS:
		return "Wales; Cymru"
	case SubdivisionGBWLV:
		return "Wolverhampton"
	case SubdivisionGBWND:
		return "Wandsworth"
	case SubdivisionGBWNM:
		return "Windsor and Maidenhead"
	case SubdivisionGBWOK:
		return "Wokingham"
	case SubdivisionGBWOR:
		return "Worcestershire"
	case SubdivisionGBWRL:
		return "Wirral"
	case SubdivisionGBWRT:
		return "Warrington"
	case SubdivisionGBWRX:
		return "Wrexham; Wrecsam"
	case SubdivisionGBWSM:
		return "Westminster"
	case SubdivisionGBWSX:
		return "West Sussex"
	case SubdivisionGBYOR:
		return "York"
	case SubdivisionGBZET:
		return "Shetland Islands"
	case SubdivisionGD01:
		return "Saint Andrew"
	case SubdivisionGD02:
		return "Saint David"
	case SubdivisionGD03:
		return "Saint George"
	case SubdivisionGD04:
		return "Saint John"
	case SubdivisionGD05:
		return "Saint Mark"
	case SubdivisionGD06:
		return "Saint Patrick"
	case SubdivisionGD10:
		return "Southern Grenadine Islands"
	case SubdivisionGEAB:
		return "Abkhazia"
	case SubdivisionGEAJ:
		return "Ajaria"
	case SubdivisionGEGU:
		return "Guria"
	case SubdivisionGEIM:
		return "Imeret’i"
	case SubdivisionGEKA:
		return "Kakhet’i"
	case SubdivisionGEKK:
		return "K’vemo K’art’li"
	case SubdivisionGEMM:
		return "Mts’khet’a-Mt’ianet’i"
	case SubdivisionGERL:
		return "Racha-Lech’khumi-K’vemo Svanet’i"
	case SubdivisionGESJ:
		return "Samts’khe-Javakhet’i"
	case SubdivisionGESK:
		return "Shida K’art’li"
	case SubdivisionGESZ:
		return "Samegrelo-Zemo Svanet’i"
	case SubdivisionGETB:
		return "T’bilisi"
	case SubdivisionGHAA:
		return "Greater Accra"
	case SubdivisionGHAH:
		return "Ashanti"
	case SubdivisionGHBA:
		return "Brong-Ahafo"
	case SubdivisionGHCP:
		return "Central"
	case SubdivisionGHEP:
		return "Eastern"
	case SubdivisionGHNP:
		return "Northern"
	case SubdivisionGHTV:
		return "Volta"
	case SubdivisionGHUE:
		return "Upper East"
	case SubdivisionGHUW:
		return "Upper West"
	case SubdivisionGHWP:
		return "Western"
	case SubdivisionGLKU:
		return "Kommune Kujalleq"
	case SubdivisionGLQA:
		return "Qaasuitsup Kommunia"
	case SubdivisionGLQE:
		return "Qeqqata Kommunia"
	case SubdivisionGLSM:
		return "Kommuneqarfik Sermersooq"
	case SubdivisionGMB:
		return "Banjul"
	case SubdivisionGML:
		return "Lower River"
	case SubdivisionGMM:
		return "Central River"
	case SubdivisionGMN:
		return "North Bank"
	case SubdivisionGMU:
		return "Upper River"
	case SubdivisionGMW:
		return "Western"
	case SubdivisionGNB:
		return "Boké"
	case SubdivisionGNBE:
		return "Beyla"
	case SubdivisionGNBF:
		return "Boffa"
	case SubdivisionGNBK:
		return "Boké"
	case SubdivisionGNC:
		return "Conakry"
	case SubdivisionGNCO:
		return "Coyah"
	case SubdivisionGND:
		return "Kindia"
	case SubdivisionGNDB:
		return "Dabola"
	case SubdivisionGNDI:
		return "Dinguiraye"
	case SubdivisionGNDL:
		return "Dalaba"
	case SubdivisionGNDU:
		return "Dubréka"
	case SubdivisionGNF:
		return "Faranah"
	case SubdivisionGNFA:
		return "Faranah"
	case SubdivisionGNFO:
		return "Forécariah"
	case SubdivisionGNFR:
		return "Fria"
	case SubdivisionGNGA:
		return "Gaoual"
	case SubdivisionGNGU:
		return "Guékédou"
	case SubdivisionGNK:
		return "Kankan"
	case SubdivisionGNKA:
		return "Kankan"
	case SubdivisionGNKB:
		return "Koubia"
	case SubdivisionGNKD:
		return "Kindia"
	case SubdivisionGNKE:
		return "Kérouané"
	case SubdivisionGNKN:
		return "Koundara"
	case SubdivisionGNKO:
		return "Kouroussa"
	case SubdivisionGNKS:
		return "Kissidougou"
	case SubdivisionGNL:
		return "Labé"
	case SubdivisionGNLA:
		return "Labé"
	case SubdivisionGNLE:
		return "Lélouma"
	case SubdivisionGNLO:
		return "Lola"
	case SubdivisionGNM:
		return "Mamou"
	case SubdivisionGNMC:
		return "Macenta"
	case SubdivisionGNMD:
		return "Mandiana"
	case SubdivisionGNML:
		return "Mali"
	case SubdivisionGNMM:
		return "Mamou"
	case SubdivisionGNN:
		return "Nzérékoré"
	case SubdivisionGNNZ:
		return "Nzérékoré"
	case SubdivisionGNPI:
		return "Pita"
	case SubdivisionGNSI:
		return "Siguiri"
	case SubdivisionGNTE:
		return "Télimélé"
	case SubdivisionGNTO:
		return "Tougué"
	case SubdivisionGNYO:
		return "Yomou"
	case SubdivisionGQAN:
		return "Annobón"
	case SubdivisionGQBN:
		return "Bioko Norte"
	case SubdivisionGQBS:
		return "Bioko Sur"
	case SubdivisionGQC:
		return "Región Continental"
	case SubdivisionGQCS:
		return "Centro Sur"
	case SubdivisionGQI:
		return "Región Insular"
	case SubdivisionGQKN:
		return "Kié-Ntem"
	case SubdivisionGQLI:
		return "Litoral"
	case SubdivisionGQWN:
		return "Wele-Nzas"
	case SubdivisionGR01:
		return "Aitolia kai Akarnania"
	case SubdivisionGR03:
		return "Voiotia"
	case SubdivisionGR04:
		return "Evvoias"
	case SubdivisionGR05:
		return "Evrytania"
	case SubdivisionGR06:
		return "Fthiotida"
	case SubdivisionGR07:
		return "Fokida"
	case SubdivisionGR11:
		return "Argolida"
	case SubdivisionGR12:
		return "Arkadia"
	case SubdivisionGR13:
		return "Achaïa"
	case SubdivisionGR14:
		return "Ileia"
	case SubdivisionGR15:
		return "Korinthia"
	case SubdivisionGR16:
		return "Lakonia"
	case SubdivisionGR17:
		return "Messinia"
	case SubdivisionGR21:
		return "Zakynthos"
	case SubdivisionGR22:
		return "Kerkyra"
	case SubdivisionGR23:
		return "Kefallonia"
	case SubdivisionGR24:
		return "Lefkada"
	case SubdivisionGR31:
		return "Arta"
	case SubdivisionGR32:
		return "Thesprotia"
	case SubdivisionGR33:
		return "Ioannina"
	case SubdivisionGR34:
		return "Preveza"
	case SubdivisionGR41:
		return "Karditsa"
	case SubdivisionGR42:
		return "Larisa"
	case SubdivisionGR43:
		return "Magnisia"
	case SubdivisionGR44:
		return "Trikala"
	case SubdivisionGR51:
		return "Grevena"
	case SubdivisionGR52:
		return "Drama"
	case SubdivisionGR53:
		return "Imathia"
	case SubdivisionGR54:
		return "Thessaloniki"
	case SubdivisionGR55:
		return "Kavala"
	case SubdivisionGR56:
		return "Kastoria"
	case SubdivisionGR57:
		return "Kilkis"
	case SubdivisionGR58:
		return "Kozani"
	case SubdivisionGR59:
		return "Pella"
	case SubdivisionGR61:
		return "Pieria"
	case SubdivisionGR62:
		return "Serres"
	case SubdivisionGR63:
		return "Florina"
	case SubdivisionGR64:
		return "Chalkidiki"
	case SubdivisionGR69:
		return "Agio Oros"
	case SubdivisionGR71:
		return "Evros"
	case SubdivisionGR72:
		return "Xanthi"
	case SubdivisionGR73:
		return "Rodopi"
	case SubdivisionGR81:
		return "Dodekanisos"
	case SubdivisionGR82:
		return "Kyklades"
	case SubdivisionGR83:
		return "Lesvos"
	case SubdivisionGR84:
		return "Samos"
	case SubdivisionGR85:
		return "Chios"
	case SubdivisionGR91:
		return "Irakleio"
	case SubdivisionGR92:
		return "Lasithi"
	case SubdivisionGR93:
		return "Rethymno"
	case SubdivisionGR94:
		return "Chania"
	case SubdivisionGRA:
		return "Anatoliki Makedonia kai Thraki"
	case SubdivisionGRA1:
		return "Attiki"
	case SubdivisionGRB:
		return "Kentriki Makedonia"
	case SubdivisionGRC:
		return "Dytiki Makedonia"
	case SubdivisionGRD:
		return "Ipeiros"
	case SubdivisionGRE:
		return "Thessalia"
	case SubdivisionGRF:
		return "Ionia Nisia"
	case SubdivisionGRG:
		return "Dytiki Ellada"
	case SubdivisionGRH:
		return "Sterea Ellada"
	case SubdivisionGRI:
		return "Attiki"
	case SubdivisionGRJ:
		return "Peloponnisos"
	case SubdivisionGRK:
		return "Voreio Aigaio"
	case SubdivisionGRL:
		return "Notio Aigaio"
	case SubdivisionGRM:
		return "Kriti"
	case SubdivisionGTAV:
		return "Alta Verapaz"
	case SubdivisionGTBV:
		return "Baja Verapaz"
	case SubdivisionGTCM:
		return "Chimaltenango"
	case SubdivisionGTCQ:
		return "Chiquimula"
	case SubdivisionGTES:
		return "Escuintla"
	case SubdivisionGTGU:
		return "Guatemala"
	case SubdivisionGTHU:
		return "Huehuetenango"
	case SubdivisionGTIZ:
		return "Izabal"
	case SubdivisionGTJA:
		return "Jalapa"
	case SubdivisionGTJU:
		return "Jutiapa"
	case SubdivisionGTPE:
		return "Petén"
	case SubdivisionGTPR:
		return "El Progreso"
	case SubdivisionGTQC:
		return "Quiché"
	case SubdivisionGTQZ:
		return "Quetzaltenango"
	case SubdivisionGTRE:
		return "Retalhuleu"
	case SubdivisionGTSA:
		return "Sacatepéquez"
	case SubdivisionGTSM:
		return "San Marcos"
	case SubdivisionGTSO:
		return "Sololá"
	case SubdivisionGTSR:
		return "Santa Rosa"
	case SubdivisionGTSU:
		return "Suchitepéquez"
	case SubdivisionGTTO:
		return "Totonicapán"
	case SubdivisionGTZA:
		return "Zacapa"
	case SubdivisionGWBA:
		return "Bafatá"
	case SubdivisionGWBL:
		return "Bolama"
	case SubdivisionGWBM:
		return "Biombo"
	case SubdivisionGWBS:
		return "Bissau"
	case SubdivisionGWCA:
		return "Cacheu"
	case SubdivisionGWGA:
		return "Gabú"
	case SubdivisionGWL:
		return "Leste"
	case SubdivisionGWN:
		return "Norte"
	case SubdivisionGWOI:
		return "Oio"
	case SubdivisionGWQU:
		return "Quinara"
	case SubdivisionGWS:
		return "Sul"
	case SubdivisionGWTO:
		return "Tombali"
	case SubdivisionGYBA:
		return "Barima-Waini"
	case SubdivisionGYCU:
		return "Cuyuni-Mazaruni"
	case SubdivisionGYDE:
		return "Demerara-Mahaica"
	case SubdivisionGYEB:
		return "East Berbice-Corentyne"
	case SubdivisionGYES:
		return "Essequibo Islands-West Demerara"
	case SubdivisionGYMA:
		return "Mahaica-Berbice"
	case SubdivisionGYPM:
		return "Pomeroon-Supenaam"
	case SubdivisionGYPT:
		return "Potaro-Siparuni"
	case SubdivisionGYUD:
		return "Upper Demerara-Berbice"
	case SubdivisionGYUT:
		return "Upper Takutu-Upper Essequibo"
	case SubdivisionHNAT:
		return "Atlántida"
	case SubdivisionHNCH:
		return "Choluteca"
	case SubdivisionHNCL:
		return "Colón"
	case SubdivisionHNCM:
		return "Comayagua"
	case SubdivisionHNCP:
		return "Copán"
	case SubdivisionHNCR:
		return "Cortés"
	case SubdivisionHNEP:
		return "El Paraíso"
	case SubdivisionHNFM:
		return "Francisco Morazán"
	case SubdivisionHNGD:
		return "Gracias a Dios"
	case SubdivisionHNIB:
		return "Islas de la Bahía"
	case SubdivisionHNIN:
		return "Intibucá"
	case SubdivisionHNLE:
		return "Lempira"
	case SubdivisionHNLP:
		return "La Paz"
	case SubdivisionHNOC:
		return "Ocotepeque"
	case SubdivisionHNOL:
		return "Olancho"
	case SubdivisionHNSB:
		return "Santa Bárbara"
	case SubdivisionHNVA:
		return "Valle"
	case SubdivisionHNYO:
		return "Yoro"
	case SubdivisionHR01:
		return "Zagrebačka županija"
	case SubdivisionHR02:
		return "Krapinsko-zagorska županija"
	case SubdivisionHR03:
		return "Sisačko-moslavačka županija"
	case SubdivisionHR04:
		return "Karlovačka županija"
	case SubdivisionHR05:
		return "Varaždinska županija"
	case SubdivisionHR06:
		return "Koprivničko-križevačka županija"
	case SubdivisionHR07:
		return "Bjelovarsko-bilogorska županija"
	case SubdivisionHR08:
		return "Primorsko-goranska županija"
	case SubdivisionHR09:
		return "Ličko-senjska županija"
	case SubdivisionHR10:
		return "Virovitičko-podravska županija"
	case SubdivisionHR11:
		return "Požeško-slavonska županija"
	case SubdivisionHR12:
		return "Brodsko-posavska županija"
	case SubdivisionHR13:
		return "Zadarska županija"
	case SubdivisionHR14:
		return "Osječko-baranjska županija"
	case SubdivisionHR15:
		return "Šibensko-kninska županija"
	case SubdivisionHR16:
		return "Vukovarsko-srijemska županija"
	case SubdivisionHR17:
		return "Splitsko-dalmatinska županija"
	case SubdivisionHR18:
		return "Istarska županija"
	case SubdivisionHR19:
		return "Dubrovačko-neretvanska županija"
	case SubdivisionHR20:
		return "Međimurska županija"
	case SubdivisionHR21:
		return "Grad Zagreb"
	case SubdivisionHTAR:
		return "Artibonite"
	case SubdivisionHTCE:
		return "Centre"
	case SubdivisionHTGA:
		return "Grande-Anse"
	case SubdivisionHTND:
		return "Nord"
	case SubdivisionHTNE:
		return "Nord-Est"
	case SubdivisionHTNO:
		return "Nord-Ouest"
	case SubdivisionHTOU:
		return "Ouest"
	case SubdivisionHTSD:
		return "Sud"
	case SubdivisionHTSE:
		return "Sud-Est"
	case SubdivisionHUBA:
		return "Baranya"
	case SubdivisionHUBC:
		return "Békéscsaba"
	case SubdivisionHUBE:
		return "Békés"
	case SubdivisionHUBK:
		return "Bács-Kiskun"
	case SubdivisionHUBU:
		return "Budapest"
	case SubdivisionHUBZ:
		return "Borsod-Abaúj-Zemplén"
	case SubdivisionHUCS:
		return "Csongrád"
	case SubdivisionHUDE:
		return "Debrecen"
	case SubdivisionHUDU:
		return "Dunaújváros"
	case SubdivisionHUEG:
		return "Eger"
	case SubdivisionHUER:
		return "Érd"
	case SubdivisionHUFE:
		return "Fejér"
	case SubdivisionHUGS:
		return "Győr-Moson-Sopron"
	case SubdivisionHUGY:
		return "Győr"
	case SubdivisionHUHB:
		return "Hajdú-Bihar"
	case SubdivisionHUHE:
		return "Heves"
	case SubdivisionHUHV:
		return "Hódmezővásárhely"
	case SubdivisionHUJN:
		return "Jász-Nagykun-Szolnok"
	case SubdivisionHUKE:
		return "Komárom-Esztergom"
	case SubdivisionHUKM:
		return "Kecskemét"
	case SubdivisionHUKV:
		return "Kaposvár"
	case SubdivisionHUMI:
		return "Miskolc"
	case SubdivisionHUNK:
		return "Nagykanizsa"
	case SubdivisionHUNO:
		return "Nógrád"
	case SubdivisionHUNY:
		return "Nyíregyháza"
	case SubdivisionHUPE:
		return "Pest"
	case SubdivisionHUPS:
		return "Pécs"
	case SubdivisionHUSD:
		return "Szeged"
	case SubdivisionHUSF:
		return "Székesfehérvár"
	case SubdivisionHUSH:
		return "Szombathely"
	case SubdivisionHUSK:
		return "Szolnok"
	case SubdivisionHUSN:
		return "Sopron"
	case SubdivisionHUSO:
		return "Somogy"
	case SubdivisionHUSS:
		return "Szekszárd"
	case SubdivisionHUST:
		return "Salgótarján"
	case SubdivisionHUSZ:
		return "Szabolcs-Szatmár-Bereg"
	case SubdivisionHUTB:
		return "Tatabánya"
	case SubdivisionHUTO:
		return "Tolna"
	case SubdivisionHUVA:
		return "Vas"
	case SubdivisionHUVE:
		return "Veszprém (county)"
	case SubdivisionHUVM:
		return "Veszprém"
	case SubdivisionHUZA:
		return "Zala"
	case SubdivisionHUZE:
		return "Zalaegerszeg"
	case SubdivisionIDAC:
		return "Aceh"
	case SubdivisionIDBA:
		return "Bali"
	case SubdivisionIDBB:
		return "Bangka Belitung"
	case SubdivisionIDBE:
		return "Bengkulu"
	case SubdivisionIDBT:
		return "Banten"
	case SubdivisionIDGO:
		return "Gorontalo"
	case SubdivisionIDIJ:
		return "Papua"
	case SubdivisionIDJA:
		return "Jambi"
	case SubdivisionIDJB:
		return "Jawa Barat"
	case SubdivisionIDJI:
		return "Jawa Timur"
	case SubdivisionIDJK:
		return "Jakarta Raya"
	case SubdivisionIDJT:
		return "Jawa Tengah"
	case SubdivisionIDJW:
		return "Jawa"
	case SubdivisionIDKA:
		return "Kalimantan"
	case SubdivisionIDKB:
		return "Kalimantan Barat"
	case SubdivisionIDKI:
		return "Kalimantan Timur"
	case SubdivisionIDKR:
		return "Kepulauan Riau"
	case SubdivisionIDKS:
		return "Kalimantan Selatan"
	case SubdivisionIDKT:
		return "Kalimantan Tengah"
	case SubdivisionIDLA:
		return "Lampung"
	case SubdivisionIDMA:
		return "Maluku"
	case SubdivisionIDML:
		return "Maluku"
	case SubdivisionIDMU:
		return "Maluku Utara"
	case SubdivisionIDNB:
		return "Nusa Tenggara Barat"
	case SubdivisionIDNT:
		return "Nusa Tenggara Timur"
	case SubdivisionIDNU:
		return "Nusa Tenggara"
	case SubdivisionIDPA:
		return "Papua"
	case SubdivisionIDPB:
		return "Papua Barat"
	case SubdivisionIDRI:
		return "Riau"
	case SubdivisionIDSA:
		return "Sulawesi Utara"
	case SubdivisionIDSB:
		return "Sumatra Barat"
	case SubdivisionIDSG:
		return "Sulawesi Tenggara"
	case SubdivisionIDSL:
		return "Sulawesi"
	case SubdivisionIDSM:
		return "Sumatera"
	case SubdivisionIDSN:
		return "Sulawesi Selatan"
	case SubdivisionIDSR:
		return "Sulawesi Barat"
	case SubdivisionIDSS:
		return "Sumatra Selatan"
	case SubdivisionIDST:
		return "Sulawesi Tengah"
	case SubdivisionIDSU:
		return "Sumatera Utara"
	case SubdivisionIDYO:
		return "Yogyakarta"
	case SubdivisionIEC:
		return "Connacht"
	case SubdivisionIECE:
		return "Clare"
	case SubdivisionIECN:
		return "Cavan"
	case SubdivisionIECO:
		return "Cork"
	case SubdivisionIECW:
		return "Carlow"
	case SubdivisionIED:
		return "Dublin"
	case SubdivisionIEDL:
		return "Donegal"
	case SubdivisionIEG:
		return "Galway"
	case SubdivisionIEKE:
		return "Kildare"
	case SubdivisionIEKK:
		return "Kilkenny"
	case SubdivisionIEKY:
		return "Kerry"
	case SubdivisionIEL:
		return "Leinster"
	case SubdivisionIELD:
		return "Longford"
	case SubdivisionIELH:
		return "Louth"
	case SubdivisionIELK:
		return "Limerick"
	case SubdivisionIELM:
		return "Leitrim"
	case SubdivisionIELS:
		return "Laois"
	case SubdivisionIEM:
		return "Munster"
	case SubdivisionIEMH:
		return "Meath"
	case SubdivisionIEMN:
		return "Monaghan"
	case SubdivisionIEMO:
		return "Mayo"
	case SubdivisionIEOY:
		return "Offaly"
	case SubdivisionIERN:
		return "Roscommon"
	case SubdivisionIESO:
		return "Sligo"
	case SubdivisionIETA:
		return "Tipperary"
	case SubdivisionIEU:
		return "Ulster"
	case SubdivisionIEWD:
		return "Waterford"
	case SubdivisionIEWH:
		return "Westmeath"
	case SubdivisionIEWW:
		return "Wicklow"
	case SubdivisionIEWX:
		return "Wexford"
	case SubdivisionILD:
		return "HaDarom"
	case SubdivisionILHA:
		return "Hefa"
	case SubdivisionILJM:
		return "Yerushalayim Al Quds"
	case SubdivisionILM:
		return "HaMerkaz"
	case SubdivisionILTA:
		return "Tel-Aviv"
	case SubdivisionILZ:
		return "HaZafon"
	case SubdivisionINAN:
		return "Andaman and Nicobar Islands"
	case SubdivisionINAP:
		return "Andhra Pradesh"
	case SubdivisionINAR:
		return "Arunachal Pradesh"
	case SubdivisionINAS:
		return "Assam"
	case SubdivisionINBR:
		return "Bihar"
	case SubdivisionINCH:
		return "Chandigarh"
	case SubdivisionINCT:
		return "Chhattisgarh"
	case SubdivisionINDD:
		return "Daman and Diu"
	case SubdivisionINDL:
		return "Delhi"
	case SubdivisionINDN:
		return "Dadra and Nagar Haveli"
	case SubdivisionINGA:
		return "Goa"
	case SubdivisionINGJ:
		return "Gujarat"
	case SubdivisionINHP:
		return "Himachal Pradesh"
	case SubdivisionINHR:
		return "Haryana"
	case SubdivisionINJH:
		return "Jharkhand"
	case SubdivisionINJK:
		return "Jammu and Kashmir"
	case SubdivisionINKA:
		return "Karnataka"
	case SubdivisionINKL:
		return "Kerala"
	case SubdivisionINLD:
		return "Lakshadweep"
	case SubdivisionINMH:
		return "Maharashtra"
	case SubdivisionINML:
		return "Meghalaya"
	case SubdivisionINMN:
		return "Manipur"
	case SubdivisionINMP:
		return "Madhya Pradesh"
	case SubdivisionINMZ:
		return "Mizoram"
	case SubdivisionINNL:
		return "Nagaland"
	case SubdivisionINOR:
		return "Odisha"
	case SubdivisionINPB:
		return "Punjab"
	case SubdivisionINPY:
		return "Puducherry"
	case SubdivisionINRJ:
		return "Rajasthan"
	case SubdivisionINSK:
		return "Sikkim"
	case SubdivisionINTG:
		return "Telangana"
	case SubdivisionINTN:
		return "Tamil Nadu"
	case SubdivisionINTR:
		return "Tripura"
	case SubdivisionINUP:
		return "Uttar Pradesh"
	case SubdivisionINUT:
		return "Uttarakhand"
	case SubdivisionINWB:
		return "West Bengal"
	case SubdivisionIQAN:
		return "Al Anbar"
	case SubdivisionIQAR:
		return "Arbil"
	case SubdivisionIQBA:
		return "Al Basrah"
	case SubdivisionIQBB:
		return "Babil"
	case SubdivisionIQBG:
		return "Baghdad"
	case SubdivisionIQDA:
		return "Dahuk"
	case SubdivisionIQDI:
		return "Diyala"
	case SubdivisionIQDQ:
		return "Dhi Qar"
	case SubdivisionIQKA:
		return "Karbala'"
	case SubdivisionIQMA:
		return "Maysan"
	case SubdivisionIQMU:
		return "Al Muthanna"
	case SubdivisionIQNA:
		return "An Najef"
	case SubdivisionIQNI:
		return "Ninawa"
	case SubdivisionIQQA:
		return "Al Qadisiyah"
	case SubdivisionIQSD:
		return "Salah ad Din"
	case SubdivisionIQSW:
		return "As Sulaymaniyah"
	case SubdivisionIQTS:
		return "At Ta'mim"
	case SubdivisionIQWA:
		return "Wasit"
	case SubdivisionIR01:
		return "Āzarbāyjān-e Sharqī"
	case SubdivisionIR02:
		return "Āzarbāyjān-e Gharbī"
	case SubdivisionIR03:
		return "Ardabīl"
	case SubdivisionIR04:
		return "Eşfahān"
	case SubdivisionIR05:
		return "Īlām"
	case SubdivisionIR06:
		return "Būshehr"
	case SubdivisionIR07:
		return "Tehrān"
	case SubdivisionIR08:
		return "Chahār Mahāll va Bakhtīārī"
	case SubdivisionIR10:
		return "Khūzestān"
	case SubdivisionIR11:
		return "Zanjān"
	case SubdivisionIR12:
		return "Semnān"
	case SubdivisionIR13:
		return "Sīstān va Balūchestān"
	case SubdivisionIR14:
		return "Fārs"
	case SubdivisionIR15:
		return "Kermān"
	case SubdivisionIR16:
		return "Kordestān"
	case SubdivisionIR17:
		return "Kermānshāh"
	case SubdivisionIR18:
		return "Kohgīlūyeh va Būyer Ahmad"
	case SubdivisionIR19:
		return "Gīlān"
	case SubdivisionIR20:
		return "Lorestān"
	case SubdivisionIR21:
		return "Māzandarān"
	case SubdivisionIR22:
		return "Markazī"
	case SubdivisionIR23:
		return "Hormozgān"
	case SubdivisionIR24:
		return "Hamadān"
	case SubdivisionIR25:
		return "Yazd"
	case SubdivisionIR26:
		return "Qom"
	case SubdivisionIR27:
		return "Golestān"
	case SubdivisionIR28:
		return "Qazvīn"
	case SubdivisionIR29:
		return "Khorāsān-e Janūbī"
	case SubdivisionIR30:
		return "Khorāsān-e Razavī"
	case SubdivisionIR31:
		return "Khorāsān-e Shemālī"
	case SubdivisionIS0:
		return "Reykjavík"
	case SubdivisionIS1:
		return "Höfuðborgarsvæðið"
	case SubdivisionIS2:
		return "Suðurnes"
	case SubdivisionIS3:
		return "Vesturland"
	case SubdivisionIS4:
		return "Vestfirðir"
	case SubdivisionIS5:
		return "Norðurland vestra"
	case SubdivisionIS6:
		return "Norðurland eystra"
	case SubdivisionIS7:
		return "Austurland"
	case SubdivisionIS8:
		return "Suðurland"
	case SubdivisionIT21:
		return "Piemonte"
	case SubdivisionIT23:
		return "Valle d'Aosta"
	case SubdivisionIT25:
		return "Lombardia"
	case SubdivisionIT32:
		return "Trentino-Alto Adige"
	case SubdivisionIT34:
		return "Veneto"
	case SubdivisionIT36:
		return "Friuli-Venezia Giulia"
	case SubdivisionIT42:
		return "Liguria"
	case SubdivisionIT45:
		return "Emilia-Romagna"
	case SubdivisionIT52:
		return "Toscana"
	case SubdivisionIT55:
		return "Umbria"
	case SubdivisionIT57:
		return "Marche"
	case SubdivisionIT62:
		return "Lazio"
	case SubdivisionIT65:
		return "Abruzzo"
	case SubdivisionIT67:
		return "Molise"
	case SubdivisionIT72:
		return "Campania"
	case SubdivisionIT75:
		return "Puglia"
	case SubdivisionIT77:
		return "Basilicata"
	case SubdivisionIT78:
		return "Calabria"
	case SubdivisionIT82:
		return "Sicilia"
	case SubdivisionIT88:
		return "Sardegna"
	case SubdivisionITAG:
		return "Agrigento"
	case SubdivisionITAL:
		return "Alessandria"
	case SubdivisionITAN:
		return "Ancona"
	case SubdivisionITAO:
		return "Aosta"
	case SubdivisionITAP:
		return "Ascoli Piceno"
	case SubdivisionITAQ:
		return "L'Aquila"
	case SubdivisionITAR:
		return "Arezzo"
	case SubdivisionITAT:
		return "Asti"
	case SubdivisionITAV:
		return "Avellino"
	case SubdivisionITBA:
		return "Bari"
	case SubdivisionITBG:
		return "Bergamo"
	case SubdivisionITBI:
		return "Biella"
	case SubdivisionITBL:
		return "Belluno"
	case SubdivisionITBN:
		return "Benevento"
	case SubdivisionITBO:
		return "Bologna"
	case SubdivisionITBR:
		return "Brindisi"
	case SubdivisionITBS:
		return "Brescia"
	case SubdivisionITBT:
		return "Barletta-Andria-Trani"
	case SubdivisionITBZ:
		return "Bolzano"
	case SubdivisionITCA:
		return "Cagliari"
	case SubdivisionITCB:
		return "Campobasso"
	case SubdivisionITCE:
		return "Caserta"
	case SubdivisionITCH:
		return "Chieti"
	case SubdivisionITCI:
		return "Carbonia-Iglesias"
	case SubdivisionITCL:
		return "Caltanissetta"
	case SubdivisionITCN:
		return "Cuneo"
	case SubdivisionITCO:
		return "Como"
	case SubdivisionITCR:
		return "Cremona"
	case SubdivisionITCS:
		return "Cosenza"
	case SubdivisionITCT:
		return "Catania"
	case SubdivisionITCZ:
		return "Catanzaro"
	case SubdivisionITEN:
		return "Enna"
	case SubdivisionITFC:
		return "Forlì-Cesena"
	case SubdivisionITFE:
		return "Ferrara"
	case SubdivisionITFG:
		return "Foggia"
	case SubdivisionITFI:
		return "Firenze"
	case SubdivisionITFM:
		return "Fermo"
	case SubdivisionITFR:
		return "Frosinone"
	case SubdivisionITGE:
		return "Genova"
	case SubdivisionITGO:
		return "Gorizia"
	case SubdivisionITGR:
		return "Grosseto"
	case SubdivisionITIM:
		return "Imperia"
	case SubdivisionITIS:
		return "Isernia"
	case SubdivisionITKR:
		return "Crotone"
	case SubdivisionITLC:
		return "Lecco"
	case SubdivisionITLE:
		return "Lecce"
	case SubdivisionITLI:
		return "Livorno"
	case SubdivisionITLO:
		return "Lodi"
	case SubdivisionITLT:
		return "Latina"
	case SubdivisionITLU:
		return "Lucca"
	case SubdivisionITMB:
		return "Monza e Brianza"
	case SubdivisionITMC:
		return "Macerata"
	case SubdivisionITME:
		return "Messina"
	case SubdivisionITMI:
		return "Milano"
	case SubdivisionITMN:
		return "Mantova"
	case SubdivisionITMO:
		return "Modena"
	case SubdivisionITMS:
		return "Massa-Carrara"
	case SubdivisionITMT:
		return "Matera"
	case SubdivisionITNA:
		return "Napoli"
	case SubdivisionITNO:
		return "Novara"
	case SubdivisionITNU:
		return "Nuoro"
	case SubdivisionITOG:
		return "Ogliastra"
	case SubdivisionITOR:
		return "Oristano"
	case SubdivisionITOT:
		return "Olbia-Tempio"
	case SubdivisionITPA:
		return "Palermo"
	case SubdivisionITPC:
		return "Piacenza"
	case SubdivisionITPD:
		return "Padova"
	case SubdivisionITPE:
		return "Pescara"
	case SubdivisionITPG:
		return "Perugia"
	case SubdivisionITPI:
		return "Pisa"
	case SubdivisionITPN:
		return "Pordenone"
	case SubdivisionITPO:
		return "Prato"
	case SubdivisionITPR:
		return "Parma"
	case SubdivisionITPT:
		return "Pistoia"
	case SubdivisionITPU:
		return "Pesaro e Urbino"
	case SubdivisionITPV:
		return "Pavia"
	case SubdivisionITPZ:
		return "Potenza"
	case SubdivisionITRA:
		return "Ravenna"
	case SubdivisionITRC:
		return "Reggio Calabria"
	case SubdivisionITRE:
		return "Reggio Emilia"
	case SubdivisionITRG:
		return "Ragusa"
	case SubdivisionITRI:
		return "Rieti"
	case SubdivisionITRM:
		return "Roma"
	case SubdivisionITRN:
		return "Rimini"
	case SubdivisionITRO:
		return "Rovigo"
	case SubdivisionITSA:
		return "Salerno"
	case SubdivisionITSI:
		return "Siena"
	case SubdivisionITSO:
		return "Sondrio"
	case SubdivisionITSP:
		return "La Spezia"
	case SubdivisionITSR:
		return "Siracusa"
	case SubdivisionITSS:
		return "Sassari"
	case SubdivisionITSV:
		return "Savona"
	case SubdivisionITTA:
		return "Taranto"
	case SubdivisionITTE:
		return "Teramo"
	case SubdivisionITTN:
		return "Trento"
	case SubdivisionITTO:
		return "Torino"
	case SubdivisionITTP:
		return "Trapani"
	case SubdivisionITTR:
		return "Terni"
	case SubdivisionITTS:
		return "Trieste"
	case SubdivisionITTV:
		return "Treviso"
	case SubdivisionITUD:
		return "Udine"
	case SubdivisionITVA:
		return "Varese"
	case SubdivisionITVB:
		return "Verbano-Cusio-Ossola"
	case SubdivisionITVC:
		return "Vercelli"
	case SubdivisionITVE:
		return "Venezia"
	case SubdivisionITVI:
		return "Vicenza"
	case SubdivisionITVR:
		return "Verona"
	case SubdivisionITVS:
		return "Medio Campidano"
	case SubdivisionITVT:
		return "Viterbo"
	case SubdivisionITVV:
		return "Vibo Valentia"
	case SubdivisionJM01:
		return "Kingston"
	case SubdivisionJM02:
		return "Saint Andrew"
	case SubdivisionJM03:
		return "Saint Thomas"
	case SubdivisionJM04:
		return "Portland"
	case SubdivisionJM05:
		return "Saint Mary"
	case SubdivisionJM06:
		return "Saint Ann"
	case SubdivisionJM07:
		return "Trelawny"
	case SubdivisionJM08:
		return "Saint James"
	case SubdivisionJM09:
		return "Hanover"
	case SubdivisionJM10:
		return "Westmoreland"
	case SubdivisionJM11:
		return "Saint Elizabeth"
	case SubdivisionJM12:
		return "Manchester"
	case SubdivisionJM13:
		return "Clarendon"
	case SubdivisionJM14:
		return "Saint Catherine"
	case SubdivisionJOAJ:
		return "‘Ajlūn"
	case SubdivisionJOAM:
		return "‘Ammān (Al ‘Aşimah)"
	case SubdivisionJOAQ:
		return "Al ‘Aqabah"
	case SubdivisionJOAT:
		return "Aţ Ţafīlah"
	case SubdivisionJOAZ:
		return "Az Zarqā'"
	case SubdivisionJOBA:
		return "Al Balqā'"
	case SubdivisionJOIR:
		return "Irbid"
	case SubdivisionJOJA:
		return "Jarash"
	case SubdivisionJOKA:
		return "Al Karak"
	case SubdivisionJOMA:
		return "Al Mafraq"
	case SubdivisionJOMD:
		return "Mādabā"
	case SubdivisionJOMN:
		return "Ma‘ān"
	case SubdivisionJP01:
		return "Hokkaido"
	case SubdivisionJP02:
		return "Aomori"
	case SubdivisionJP03:
		return "Iwate"
	case SubdivisionJP04:
		return "Miyagi"
	case SubdivisionJP05:
		return "Akita"
	case SubdivisionJP06:
		return "Yamagata"
	case SubdivisionJP07:
		return "Fukushima"
	case SubdivisionJP08:
		return "Ibaraki"
	case SubdivisionJP09:
		return "Tochigi"
	case SubdivisionJP10:
		return "Gunma"
	case SubdivisionJP11:
		return "Saitama"
	case SubdivisionJP12:
		return "Chiba"
	case SubdivisionJP13:
		return "Tokyo"
	case SubdivisionJP14:
		return "Kanagawa"
	case SubdivisionJP15:
		return "Niigata"
	case SubdivisionJP16:
		return "Toyama"
	case SubdivisionJP17:
		return "Ishikawa"
	case SubdivisionJP18:
		return "Fukui"
	case SubdivisionJP19:
		return "Yamanashi"
	case SubdivisionJP20:
		return "Nagano"
	case SubdivisionJP21:
		return "Gifu"
	case SubdivisionJP22:
		return "Shizuoka"
	case SubdivisionJP23:
		return "Aichi"
	case SubdivisionJP24:
		return "Mie"
	case SubdivisionJP25:
		return "Shiga"
	case SubdivisionJP26:
		return "Kyoto"
	case SubdivisionJP27:
		return "Osaka"
	case SubdivisionJP28:
		return "Hyogo"
	case SubdivisionJP29:
		return "Nara"
	case SubdivisionJP30:
		return "Wakayama"
	case SubdivisionJP31:
		return "Tottori"
	case SubdivisionJP32:
		return "Shimane"
	case SubdivisionJP33:
		return "Okayama"
	case SubdivisionJP34:
		return "Hiroshima"
	case SubdivisionJP35:
		return "Yamaguchi"
	case SubdivisionJP36:
		return "Tokushima"
	case SubdivisionJP37:
		return "Kagawa"
	case SubdivisionJP38:
		return "Ehime"
	case SubdivisionJP39:
		return "Kochi"
	case SubdivisionJP40:
		return "Fukuoka"
	case SubdivisionJP41:
		return "Saga"
	case SubdivisionJP42:
		return "Nagasaki"
	case SubdivisionJP43:
		return "Kumamoto"
	case SubdivisionJP44:
		return "Oita"
	case SubdivisionJP45:
		return "Miyazaki"
	case SubdivisionJP46:
		return "Kagoshima"
	case SubdivisionJP47:
		return "Okinawa"
	case SubdivisionKE01:
		return "Baringo"
	case SubdivisionKE02:
		return "Bomet"
	case SubdivisionKE03:
		return "Bungoma"
	case SubdivisionKE04:
		return "Busia"
	case SubdivisionKE05:
		return "Elgeyo/Marakwet"
	case SubdivisionKE06:
		return "Embu"
	case SubdivisionKE07:
		return "Garissa"
	case SubdivisionKE08:
		return "Homa Bay"
	case SubdivisionKE09:
		return "Isiolo"
	case SubdivisionKE10:
		return "Kajiado"
	case SubdivisionKE11:
		return "Kakamega"
	case SubdivisionKE12:
		return "Kericho"
	case SubdivisionKE13:
		return "Kiambu"
	case SubdivisionKE14:
		return "Kilifi"
	case SubdivisionKE15:
		return "Kirinyaga"
	case SubdivisionKE16:
		return "Kisii"
	case SubdivisionKE17:
		return "Kisumu"
	case SubdivisionKE18:
		return "Kitui"
	case SubdivisionKE19:
		return "Kwale"
	case SubdivisionKE20:
		return "Laikipia"
	case SubdivisionKE21:
		return "Lamu"
	case SubdivisionKE22:
		return "Machakos"
	case SubdivisionKE23:
		return "Makueni"
	case SubdivisionKE24:
		return "Mandera"
	case SubdivisionKE25:
		return "Marsabit"
	case SubdivisionKE26:
		return "Meru"
	case SubdivisionKE27:
		return "Migori"
	case SubdivisionKE28:
		return "Mombasa"
	case SubdivisionKE29:
		return "Murang'a"
	case SubdivisionKE30:
		return "Nairobi City"
	case SubdivisionKE31:
		return "Nakuru"
	case SubdivisionKE32:
		return "Nandi"
	case SubdivisionKE33:
		return "Narok"
	case SubdivisionKE34:
		return "Nyamira"
	case SubdivisionKE35:
		return "Nyandarua"
	case SubdivisionKE36:
		return "Nyeri"
	case SubdivisionKE37:
		return "Samburu"
	case SubdivisionKE38:
		return "Siaya"
	case SubdivisionKE39:
		return "Taita/Taveta"
	case SubdivisionKE40:
		return "Tana River"
	case SubdivisionKE41:
		return "Tharaka-Nithi"
	case SubdivisionKE42:
		return "Trans Nzoia"
	case SubdivisionKE43:
		return "Turkana"
	case SubdivisionKE44:
		return "Uasin Gishu"
	case SubdivisionKE45:
		return "Vihiga"
	case SubdivisionKE46:
		return "Wajir"
	case SubdivisionKE47:
		return "West Pokot"
	case SubdivisionKGB:
		return "Batken"
	case SubdivisionKGC:
		return "Chü"
	case SubdivisionKGGB:
		return "Bishkek"
	case SubdivisionKGJ:
		return "Jalal-Abad"
	case SubdivisionKGN:
		return "Naryn"
	case SubdivisionKGO:
		return "Osh"
	case SubdivisionKGT:
		return "Talas"
	case SubdivisionKGY:
		return "Ysyk-Köl"
	case SubdivisionKH1:
		return "Banteay Mean Chey"
	case SubdivisionKH10:
		return "Krachoh"
	case SubdivisionKH11:
		return "Mondol Kiri"
	case SubdivisionKH12:
		return "Phnom Penh"
	case SubdivisionKH13:
		return "Preah Vihear"
	case SubdivisionKH14:
		return "Prey Veaeng"
	case SubdivisionKH15:
		return "Pousaat"
	case SubdivisionKH16:
		return "Rotanak Kiri"
	case SubdivisionKH17:
		return "Siem Reab"
	case SubdivisionKH18:
		return "Krong Preah Sihanouk"
	case SubdivisionKH19:
		return "Stueng Traeng"
	case SubdivisionKH2:
		return "Battambang"
	case SubdivisionKH20:
		return "Svaay Rieng"
	case SubdivisionKH21:
		return "Taakaev"
	case SubdivisionKH22:
		return "Otdar Mean Chey"
	case SubdivisionKH23:
		return "Krong Kaeb"
	case SubdivisionKH24:
		return "Krong Pailin"
	case SubdivisionKH3:
		return "Kampong Cham"
	case SubdivisionKH4:
		return "Kampong Chhnang"
	case SubdivisionKH5:
		return "Kampong Speu"
	case SubdivisionKH6:
		return "Kampong Thom"
	case SubdivisionKH7:
		return "Kampot"
	case SubdivisionKH8:
		return "Kandal"
	case SubdivisionKH9:
		return "Kach Kong"
	case SubdivisionKIG:
		return "Gilbert Islands"
	case SubdivisionKIL:
		return "Line Islands"
	case SubdivisionKIP:
		return "Phoenix Islands"
	case SubdivisionKMA:
		return "Andjouân (Anjwān)"
	case SubdivisionKMG:
		return "Andjazîdja (Anjazījah)"
	case SubdivisionKMM:
		return "Moûhîlî (Mūhīlī)"
	case SubdivisionKN01:
		return "Christ Church Nichola Town"
	case SubdivisionKN02:
		return "Saint Anne Sandy Point"
	case SubdivisionKN03:
		return "Saint George Basseterre"
	case SubdivisionKN04:
		return "Saint George Gingerland"
	case SubdivisionKN05:
		return "Saint James Windward"
	case SubdivisionKN06:
		return "Saint John Capisterre"
	case SubdivisionKN07:
		return "Saint John Figtree"
	case SubdivisionKN08:
		return "Saint Mary Cayon"
	case SubdivisionKN09:
		return "Saint Paul Capisterre"
	case SubdivisionKN10:
		return "Saint Paul Charlestown"
	case SubdivisionKN11:
		return "Saint Peter Basseterre"
	case SubdivisionKN12:
		return "Saint Thomas Lowland"
	case SubdivisionKN13:
		return "Saint Thomas Middle Island"
	case SubdivisionKN15:
		return "Trinity Palmetto Point"
	case SubdivisionKNK:
		return "Saint Kitts"
	case SubdivisionKNN:
		return "Nevis"
	case SubdivisionKP01:
		return "P’yŏngyang"
	case SubdivisionKP02:
		return "P’yŏngan-namdo"
	case SubdivisionKP03:
		return "P’yŏngan-bukto"
	case SubdivisionKP04:
		return "Chagang-do"
	case SubdivisionKP05:
		return "Hwanghae-namdo"
	case SubdivisionKP06:
		return "Hwanghae-bukto"
	case SubdivisionKP07:
		return "Kangwŏn-do"
	case SubdivisionKP08:
		return "Hamgyŏng-namdo"
	case SubdivisionKP09:
		return "Hamgyŏng-bukto"
	case SubdivisionKP10:
		return "Yanggang-do"
	case SubdivisionKP13:
		return "Nasŏn (Najin-Sŏnbong)"
	case SubdivisionKR11:
		return "Seoul Teugbyeolsi"
	case SubdivisionKR26:
		return "Busan Gwang'yeogsi"
	case SubdivisionKR27:
		return "Daegu Gwang'yeogsi"
	case SubdivisionKR28:
		return "Incheon Gwang'yeogsi"
	case SubdivisionKR29:
		return "Gwangju Gwang'yeogsi"
	case SubdivisionKR30:
		return "Daejeon Gwang'yeogsi"
	case SubdivisionKR31:
		return "Ulsan Gwang'yeogsi"
	case SubdivisionKR41:
		return "Gyeonggido"
	case SubdivisionKR42:
		return "Gang'weondo"
	case SubdivisionKR43:
		return "Chungcheongbukdo"
	case SubdivisionKR44:
		return "Chungcheongnamdo"
	case SubdivisionKR45:
		return "Jeonrabukdo"
	case SubdivisionKR46:
		return "Jeonranamdo"
	case SubdivisionKR47:
		return "Gyeongsangbukdo"
	case SubdivisionKR48:
		return "Gyeongsangnamdo"
	case SubdivisionKR49:
		return "Jejudo"
	case SubdivisionKWAH:
		return "Al Ahmadi"
	case SubdivisionKWFA:
		return "Al Farwānīyah"
	case SubdivisionKWHA:
		return "Hawallī"
	case SubdivisionKWJA:
		return "Al Jahrrā’"
	case SubdivisionKWKU:
		return "Al Kuwayt (Al ‘Āşimah)"
	case SubdivisionKWMU:
		return "Mubārak al Kabīr"
	case SubdivisionKZAKM:
		return "Aqmola oblysy"
	case SubdivisionKZAKT:
		return "Aqtöbe oblysy"
	case SubdivisionKZALA:
		return "Almaty"
	case SubdivisionKZALM:
		return "Almaty oblysy"
	case SubdivisionKZAST:
		return "Astana"
	case SubdivisionKZATY:
		return "Atyraū oblysy"
	case SubdivisionKZKAR:
		return "Qaraghandy oblysy"
	case SubdivisionKZKUS:
		return "Qostanay oblysy"
	case SubdivisionKZKZY:
		return "Qyzylorda oblysy"
	case SubdivisionKZMAN:
		return "Mangghystaū oblysy"
	case SubdivisionKZPAV:
		return "Pavlodar oblysy"
	case SubdivisionKZSEV:
		return "Soltüstik Quzaqstan oblysy"
	case SubdivisionKZVOS:
		return "Shyghys Qazaqstan oblysy"
	case SubdivisionKZYUZ:
		return "Ongtüstik Qazaqstan oblysy"
	case SubdivisionKZZAP:
		return "Batys Quzaqstan oblysy"
	case SubdivisionKZZHA:
		return "Zhambyl oblysy"
	case SubdivisionLAAT:
		return "Attapu"
	case SubdivisionLABK:
		return "Bokèo"
	case SubdivisionLABL:
		return "Bolikhamxai"
	case SubdivisionLACH:
		return "Champasak"
	case SubdivisionLAHO:
		return "Houaphan"
	case SubdivisionLAKH:
		return "Khammouan"
	case SubdivisionLALM:
		return "Louang Namtha"
	case SubdivisionLALP:
		return "Louangphabang"
	case SubdivisionLAOU:
		return "Oudômxai"
	case SubdivisionLAPH:
		return "Phôngsali"
	case SubdivisionLASL:
		return "Salavan"
	case SubdivisionLASV:
		return "Savannakhét"
	case SubdivisionLAVI:
		return "Vientiane"
	case SubdivisionLAVT:
		return "Vientiane"
	case SubdivisionLAXA:
		return "Xaignabouli"
	case SubdivisionLAXE:
		return "Xékong"
	case SubdivisionLAXI:
		return "Xiangkhouang"
	case SubdivisionLAXS:
		return "Xaisômboun"
	case SubdivisionLBAK:
		return "Aakkâr"
	case SubdivisionLBAS:
		return "Liban-Nord"
	case SubdivisionLBBA:
		return "Beyrouth"
	case SubdivisionLBBH:
		return "Baalbek-Hermel"
	case SubdivisionLBBI:
		return "Béqaa"
	case SubdivisionLBJA:
		return "Liban-Sud"
	case SubdivisionLBJL:
		return "Mont-Liban"
	case SubdivisionLBNA:
		return "Nabatîyé"
	case SubdivisionLI01:
		return "Balzers"
	case SubdivisionLI02:
		return "Eschen"
	case SubdivisionLI03:
		return "Gamprin"
	case SubdivisionLI04:
		return "Mauren"
	case SubdivisionLI05:
		return "Planken"
	case SubdivisionLI06:
		return "Ruggell"
	case SubdivisionLI07:
		return "Schaan"
	case SubdivisionLI08:
		return "Schellenberg"
	case SubdivisionLI09:
		return "Triesen"
	case SubdivisionLI10:
		return "Triesenberg"
	case SubdivisionLI11:
		return "Vaduz"
	case SubdivisionLK1:
		return "Basnāhira paḷāta"
	case SubdivisionLK11:
		return "Kŏḷamba"
	case SubdivisionLK12:
		return "Gampaha"
	case SubdivisionLK13:
		return "Kaḷutara"
	case SubdivisionLK2:
		return "Madhyama paḷāta"
	case SubdivisionLK21:
		return "Mahanuvara"
	case SubdivisionLK22:
		return "Mātale"
	case SubdivisionLK23:
		return "Nuvara Ĕliya"
	case SubdivisionLK3:
		return "Dakuṇu paḷāta"
	case SubdivisionLK31:
		return "Gālla"
	case SubdivisionLK32:
		return "Mātara"
	case SubdivisionLK33:
		return "Hambantŏṭa"
	case SubdivisionLK4:
		return "Uturu paḷāta"
	case SubdivisionLK41:
		return "Yāpanaya"
	case SubdivisionLK42:
		return "Kilinŏchchi"
	case SubdivisionLK43:
		return "Mannārama"
	case SubdivisionLK44:
		return "Vavuniyāva"
	case SubdivisionLK45:
		return "Mulativ"
	case SubdivisionLK5:
		return "Næ̆gĕnahira paḷāta"
	case SubdivisionLK51:
		return "Maḍakalapuva"
	case SubdivisionLK52:
		return "Ampāara"
	case SubdivisionLK53:
		return "Trikuṇāmalaya"
	case SubdivisionLK6:
		return "Vayamba paḷāta"
	case SubdivisionLK61:
		return "Kuruṇægala"
	case SubdivisionLK62:
		return "Puttalama"
	case SubdivisionLK7:
		return "Uturumæ̆da paḷāta"
	case SubdivisionLK71:
		return "Anurādhapura"
	case SubdivisionLK72:
		return "Pŏḷŏnnaruva"
	case SubdivisionLK8:
		return "Ūva paḷāta"
	case SubdivisionLK81:
		return "Badulla"
	case SubdivisionLK82:
		return "Mŏṇarāgala"
	case SubdivisionLK9:
		return "Sabaragamuva paḷāta"
	case SubdivisionLK91:
		return "Ratnapura"
	case SubdivisionLK92:
		return "Kægalla"
	case SubdivisionLRBG:
		return "Bong"
	case SubdivisionLRBM:
		return "Bomi"
	case SubdivisionLRCM:
		return "Grand Cape Mount"
	case SubdivisionLRGB:
		return "Grand Bassa"
	case SubdivisionLRGG:
		return "Grand Gedeh"
	case SubdivisionLRGK:
		return "Grand Kru"
	case SubdivisionLRLO:
		return "Lofa"
	case SubdivisionLRMG:
		return "Margibi"
	case SubdivisionLRMO:
		return "Montserrado"
	case SubdivisionLRMY:
		return "Maryland"
	case SubdivisionLRNI:
		return "Nimba"
	case SubdivisionLRRI:
		return "Rivercess"
	case SubdivisionLRSI:
		return "Sinoe"
	case SubdivisionLSA:
		return "Maseru"
	case SubdivisionLSB:
		return "Butha-Buthe"
	case SubdivisionLSC:
		return "Leribe"
	case SubdivisionLSD:
		return "Berea"
	case SubdivisionLSE:
		return "Mafeteng"
	case SubdivisionLSF:
		return "Mohale's Hoek"
	case SubdivisionLSG:
		return "Quthing"
	case SubdivisionLSH:
		return "Qacha's Nek"
	case SubdivisionLSJ:
		return "Mokhotlong"
	case SubdivisionLSK:
		return "Thaba-Tseka"
	case SubdivisionLTAL:
		return "Alytaus Apskritis"
	case SubdivisionLTKL:
		return "Klaipėdos Apskritis"
	case SubdivisionLTKU:
		return "Kauno Apskritis"
	case SubdivisionLTMR:
		return "Marijampolės Apskritis"
	case SubdivisionLTPN:
		return "Panevėžio Apskritis"
	case SubdivisionLTSA:
		return "Šiaulių Apskritis"
	case SubdivisionLTTA:
		return "Tauragés Apskritis"
	case SubdivisionLTTE:
		return "Telšių Apskritis"
	case SubdivisionLTUT:
		return "Utenos Apskritis"
	case SubdivisionLTVL:
		return "Vilniaus Apskritis"
	case SubdivisionLUD:
		return "Diekirch"
	case SubdivisionLUG:
		return "Grevenmacher"
	case SubdivisionLUL:
		return "Luxembourg"
	case SubdivisionLV001:
		return "Aglonas novads"
	case SubdivisionLV002:
		return "Aizkraukles novads"
	case SubdivisionLV003:
		return "Aizputes novads"
	case SubdivisionLV004:
		return "Aknīstes novads"
	case SubdivisionLV005:
		return "Alojas novads"
	case SubdivisionLV006:
		return "Alsungas novads"
	case SubdivisionLV007:
		return "Alūksnes novads"
	case SubdivisionLV008:
		return "Amatas novads"
	case SubdivisionLV009:
		return "Apes novads"
	case SubdivisionLV010:
		return "Auces novads"
	case SubdivisionLV011:
		return "Ādažu novads"
	case SubdivisionLV012:
		return "Babītes novads"
	case SubdivisionLV013:
		return "Baldones novads"
	case SubdivisionLV014:
		return "Baltinavas novads"
	case SubdivisionLV015:
		return "Balvu novads"
	case SubdivisionLV016:
		return "Bauskas novads"
	case SubdivisionLV017:
		return "Beverīnas novads"
	case SubdivisionLV018:
		return "Brocēnu novads"
	case SubdivisionLV019:
		return "Burtnieku novads"
	case SubdivisionLV020:
		return "Carnikavas novads"
	case SubdivisionLV021:
		return "Cesvaines novads"
	case SubdivisionLV022:
		return "Cēsu novads"
	case SubdivisionLV023:
		return "Ciblas novads"
	case SubdivisionLV024:
		return "Dagdas novads"
	case SubdivisionLV025:
		return "Daugavpils novads"
	case SubdivisionLV026:
		return "Dobeles novads"
	case SubdivisionLV027:
		return "Dundagas novads"
	case SubdivisionLV028:
		return "Durbes novads"
	case SubdivisionLV029:
		return "Engures novads"
	case SubdivisionLV030:
		return "Ērgļu novads"
	case SubdivisionLV031:
		return "Garkalnes novads"
	case SubdivisionLV032:
		return "Grobiņas novads"
	case SubdivisionLV033:
		return "Gulbenes novads"
	case SubdivisionLV034:
		return "Iecavas novads"
	case SubdivisionLV035:
		return "Ikšķiles novads"
	case SubdivisionLV036:
		return "Ilūkstes novads"
	case SubdivisionLV037:
		return "Inčukalna novads"
	case SubdivisionLV038:
		return "Jaunjelgavas novads"
	case SubdivisionLV039:
		return "Jaunpiebalgas novads"
	case SubdivisionLV040:
		return "Jaunpils novads"
	case SubdivisionLV041:
		return "Jelgavas novads"
	case SubdivisionLV042:
		return "Jēkabpils novads"
	case SubdivisionLV043:
		return "Kandavas novads"
	case SubdivisionLV044:
		return "Kārsavas novads"
	case SubdivisionLV045:
		return "Kocēnu novads"
	case SubdivisionLV046:
		return "Kokneses novads"
	case SubdivisionLV047:
		return "Krāslavas novads"
	case SubdivisionLV048:
		return "Krimuldas novads"
	case SubdivisionLV049:
		return "Krustpils novads"
	case SubdivisionLV050:
		return "Kuldīgas novads"
	case SubdivisionLV051:
		return "Ķeguma novads"
	case SubdivisionLV052:
		return "Ķekavas novads"
	case SubdivisionLV053:
		return "Lielvārdes novads"
	case SubdivisionLV054:
		return "Limbažu novads"
	case SubdivisionLV055:
		return "Līgatnes novads"
	case SubdivisionLV056:
		return "Līvānu novads"
	case SubdivisionLV057:
		return "Lubānas novads"
	case SubdivisionLV058:
		return "Ludzas novads"
	case SubdivisionLV059:
		return "Madonas novads"
	case SubdivisionLV060:
		return "Mazsalacas novads"
	case SubdivisionLV061:
		return "Mālpils novads"
	case SubdivisionLV062:
		return "Mārupes novads"
	case SubdivisionLV063:
		return "Mērsraga novads"
	case SubdivisionLV064:
		return "Naukšēnu novads"
	case SubdivisionLV065:
		return "Neretas novads"
	case SubdivisionLV066:
		return "Nīcas novads"
	case SubdivisionLV067:
		return "Ogres novads"
	case SubdivisionLV068:
		return "Olaines novads"
	case SubdivisionLV069:
		return "Ozolnieku novads"
	case SubdivisionLV070:
		return "Pārgaujas novads"
	case SubdivisionLV071:
		return "Pāvilostas novads"
	case SubdivisionLV072:
		return "Pļaviņu novads"
	case SubdivisionLV073:
		return "Preiļu novads"
	case SubdivisionLV074:
		return "Priekules novads"
	case SubdivisionLV075:
		return "Priekuļu novads"
	case SubdivisionLV076:
		return "Raunas novads"
	case SubdivisionLV077:
		return "Rēzeknes novads"
	case SubdivisionLV078:
		return "Riebiņu novads"
	case SubdivisionLV079:
		return "Rojas novads"
	case SubdivisionLV080:
		return "Ropažu novads"
	case SubdivisionLV081:
		return "Rucavas novads"
	case SubdivisionLV082:
		return "Rugāju novads"
	case SubdivisionLV083:
		return "Rundāles novads"
	case SubdivisionLV084:
		return "Rūjienas novads"
	case SubdivisionLV085:
		return "Salas novads"
	case SubdivisionLV086:
		return "Salacgrīvas novads"
	case SubdivisionLV087:
		return "Salaspils novads"
	case SubdivisionLV088:
		return "Saldus novads"
	case SubdivisionLV089:
		return "Saulkrastu novads"
	case SubdivisionLV090:
		return "Sējas novads"
	case SubdivisionLV091:
		return "Siguldas novads"
	case SubdivisionLV092:
		return "Skrīveru novads"
	case SubdivisionLV093:
		return "Skrundas novads"
	case SubdivisionLV094:
		return "Smiltenes novads"
	case SubdivisionLV095:
		return "Stopiņu novads"
	case SubdivisionLV096:
		return "Strenču novads"
	case SubdivisionLV097:
		return "Talsu novads"
	case SubdivisionLV098:
		return "Tērvetes novads"
	case SubdivisionLV099:
		return "Tukuma novads"
	case SubdivisionLV100:
		return "Vaiņodes novads"
	case SubdivisionLV101:
		return "Valkas novads"
	case SubdivisionLV102:
		return "Varakļānu novads"
	case SubdivisionLV103:
		return "Vārkavas novads"
	case SubdivisionLV104:
		return "Vecpiebalgas novads"
	case SubdivisionLV105:
		return "Vecumnieku novads"
	case SubdivisionLV106:
		return "Ventspils novads"
	case SubdivisionLV107:
		return "Viesītes novads"
	case SubdivisionLV108:
		return "Viļakas novads"
	case SubdivisionLV109:
		return "Viļānu novads"
	case SubdivisionLV110:
		return "Zilupes novads"
	case SubdivisionLVDGV:
		return "Daugavpils"
	case SubdivisionLVJEL:
		return "Jelgava"
	case SubdivisionLVJKB:
		return "Jēkabpils"
	case SubdivisionLVJUR:
		return "Jūrmala"
	case SubdivisionLVLPX:
		return "Liepāja"
	case SubdivisionLVREZ:
		return "Rēzekne"
	case SubdivisionLVRIX:
		return "Rīga"
	case SubdivisionLVVEN:
		return "Ventspils"
	case SubdivisionLVVMR:
		return "Valmiera"
	case SubdivisionLYBA:
		return "Banghāzī"
	case SubdivisionLYBU:
		return "Al Buţnān"
	case SubdivisionLYDR:
		return "Darnah"
	case SubdivisionLYGT:
		return "Ghāt"
	case SubdivisionLYJA:
		return "Al Jabal al Akhḑar"
	case SubdivisionLYJB:
		return "Jaghbūb"
	case SubdivisionLYJG:
		return "Al Jabal al Gharbī"
	case SubdivisionLYJI:
		return "Al Jifārah"
	case SubdivisionLYJU:
		return "Al Jufrah"
	case SubdivisionLYKF:
		return "Al Kufrah"
	case SubdivisionLYMB:
		return "Al Marqab"
	case SubdivisionLYMI:
		return "Mişrātah"
	case SubdivisionLYMJ:
		return "Al Marj"
	case SubdivisionLYMQ:
		return "Murzuq"
	case SubdivisionLYNL:
		return "Nālūt"
	case SubdivisionLYNQ:
		return "An Nuqaţ al Khams"
	case SubdivisionLYSB:
		return "Sabhā"
	case SubdivisionLYSR:
		return "Surt"
	case SubdivisionLYTB:
		return "Ţarābulus"
	case SubdivisionLYWA:
		return "Al Wāḩāt"
	case SubdivisionLYWD:
		return "Wādī al Ḩayāt"
	case SubdivisionLYWS:
		return "Wādī ash Shāţiʾ"
	case SubdivisionLYZA:
		return "Az Zāwiyah"
	case SubdivisionMA01:
		return "Tanger-Tétouan-Al Hoceïma"
	case SubdivisionMA02:
		return "L'Oriental"
	case SubdivisionMA03:
		return "Fès-Meknès"
	case SubdivisionMA04:
		return "Rabat-Salé-Kénitra"
	case SubdivisionMA05:
		return "Béni Mellal-Khénifra"
	case SubdivisionMA06:
		return "Casablanca-Settat"
	case SubdivisionMA07:
		return "Marrakech-Safi"
	case SubdivisionMA08:
		return "Drâa-Tafilalet"
	case SubdivisionMA09:
		return "Souss-Massa"
	case SubdivisionMA10:
		return "Guelmim-Oued Noun (EH-partial)"
	case SubdivisionMA11:
		return "Laâyoune-Sakia El Hamra (EH-partial)"
	case SubdivisionMA12:
		return "Dakhla-Oued Ed-Dahab (EH)"
	case SubdivisionMAAGD:
		return "Agadir-Ida-Ou-Tanane"
	case SubdivisionMAAOU:
		return "Aousserd (EH)"
	case SubdivisionMAASZ:
		return "Assa-Zag (EH-partial)"
	case SubdivisionMAAZI:
		return "Azilal"
	case SubdivisionMABEM:
		return "Béni Mellal"
	case SubdivisionMABER:
		return "Berkane"
	case SubdivisionMABES:
		return "Benslimane"
	case SubdivisionMABOD:
		return "Boujdour (EH)"
	case SubdivisionMABOM:
		return "Boulemane"
	case SubdivisionMABRR:
		return "Berrechid"
	case SubdivisionMACAS:
		return "Casablanca"
	case SubdivisionMACHE:
		return "Chefchaouen"
	case SubdivisionMACHI:
		return "Chichaoua"
	case SubdivisionMACHT:
		return "Chtouka-Ait Baha"
	case SubdivisionMADRI:
		return "Driouch"
	case SubdivisionMAERR:
		return "Errachidia"
	case SubdivisionMAESI:
		return "Essaouira"
	case SubdivisionMAESM:
		return "Es-Semara (EH-partial)"
	case SubdivisionMAFAH:
		return "Fahs-Anjra"
	case SubdivisionMAFES:
		return "Fès"
	case SubdivisionMAFIG:
		return "Figuig"
	case SubdivisionMAFQH:
		return "Fquih Ben Salah"
	case SubdivisionMAGUE:
		return "Guelmim"
	case SubdivisionMAGUF:
		return "Guercif"
	case SubdivisionMAHAJ:
		return "El Hajeb"
	case SubdivisionMAHAO:
		return "Al Haouz"
	case SubdivisionMAHOC:
		return "Al Hoceïma"
	case SubdivisionMAIFR:
		return "Ifrane"
	case SubdivisionMAINE:
		return "Inezgane-Ait Melloul"
	case SubdivisionMAJDI:
		return "El Jadida"
	case SubdivisionMAJRA:
		return "Jerada"
	case SubdivisionMAKEN:
		return "Kénitra"
	case SubdivisionMAKES:
		return "El Kelâa des Sraghna"
	case SubdivisionMAKHE:
		return "Khemisset"
	case SubdivisionMAKHN:
		return "Khenifra"
	case SubdivisionMAKHO:
		return "Khouribga"
	case SubdivisionMALAA:
		return "Laâyoune (EH)"
	case SubdivisionMALAR:
		return "Larache"
	case SubdivisionMAMAR:
		return "Marrakech"
	case SubdivisionMAMDF:
		return "M’diq-Fnideq"
	case SubdivisionMAMED:
		return "Médiouna"
	case SubdivisionMAMEK:
		return "Meknès"
	case SubdivisionMAMID:
		return "Midelt"
	case SubdivisionMAMOH:
		return "Mohammadia"
	case SubdivisionMAMOU:
		return "Moulay Yacoub"
	case SubdivisionMANAD:
		return "Nador"
	case SubdivisionMANOU:
		return "Nouaceur"
	case SubdivisionMAOUA:
		return "Ouarzazate"
	case SubdivisionMAOUD:
		return "Oued Ed-Dahab (EH)"
	case SubdivisionMAOUJ:
		return "Oujda-Angad"
	case SubdivisionMAOUZ:
		return "Ouezzane"
	case SubdivisionMARAB:
		return "Rabat"
	case SubdivisionMAREH:
		return "Rehamna"
	case SubdivisionMASAF:
		return "Safi"
	case SubdivisionMASAL:
		return "Salé"
	case SubdivisionMASEF:
		return "Sefrou"
	case SubdivisionMASET:
		return "Settat"
	case SubdivisionMASIB:
		return "Sidi Bennour"
	case SubdivisionMASIF:
		return "Sidi Ifni"
	case SubdivisionMASIK:
		return "Sidi Kacem"
	case SubdivisionMASIL:
		return "Sidi Slimane"
	case SubdivisionMASKH:
		return "Skhirate-Témara"
	case SubdivisionMATAF:
		return "Tarfaya (EH-partial)"
	case SubdivisionMATAI:
		return "Taourirt"
	case SubdivisionMATAO:
		return "Taounate"
	case SubdivisionMATAR:
		return "Taroudant"
	case SubdivisionMATAT:
		return "Tata"
	case SubdivisionMATAZ:
		return "Taza"
	case SubdivisionMATET:
		return "Tétouan"
	case SubdivisionMATIN:
		return "Tinghir"
	case SubdivisionMATIZ:
		return "Tiznit"
	case SubdivisionMATNG:
		return "Tanger-Assilah"
	case SubdivisionMATNT:
		return "Tan-Tan (EH-partial)"
	case SubdivisionMAYUS:
		return "Youssoufia"
	case SubdivisionMAZAG:
		return "Zagora"
	case SubdivisionMCCL:
		return "La Colle"
	case SubdivisionMCCO:
		return "La Condamine"
	case SubdivisionMCFO:
		return "Fontvieille"
	case SubdivisionMCGA:
		return "La Gare"
	case SubdivisionMCJE:
		return "Jardin Exotique"
	case SubdivisionMCLA:
		return "Larvotto"
	case SubdivisionMCMA:
		return "Malbousquet"
	case SubdivisionMCMC:
		return "Monte-Carlo"
	case SubdivisionMCMG:
		return "Moneghetti"
	case SubdivisionMCMO:
		return "Monaco-Ville"
	case SubdivisionMCMU:
		return "Moulins"
	case SubdivisionMCPH:
		return "Port-Hercule"
	case SubdivisionMCSD:
		return "Sainte-Dévote"
	case SubdivisionMCSO:
		return "La Source"
	case SubdivisionMCSP:
		return "Spélugues"
	case SubdivisionMCSR:
		return "Saint-Roman"
	case SubdivisionMCVR:
		return "Vallon de la Rousse"
	case SubdivisionMDAN:
		return "Anenii Noi"
	case SubdivisionMDBA:
		return "Bălți"
	case SubdivisionMDBD:
		return "Tighina"
	case SubdivisionMDBR:
		return "Briceni"
	case SubdivisionMDBS:
		return "Basarabeasca"
	case SubdivisionMDCA:
		return "Cahul"
	case SubdivisionMDCL:
		return "Călărași"
	case SubdivisionMDCM:
		return "Cimișlia"
	case SubdivisionMDCR:
		return "Criuleni"
	case SubdivisionMDCS:
		return "Căușeni"
	case SubdivisionMDCT:
		return "Cantemir"
	case SubdivisionMDCU:
		return "Chișinău"
	case SubdivisionMDDO:
		return "Dondușeni"
	case SubdivisionMDDR:
		return "Drochia"
	case SubdivisionMDDU:
		return "Dubăsari"
	case SubdivisionMDED:
		return "Edineț"
	case SubdivisionMDFA:
		return "Fălești"
	case SubdivisionMDFL:
		return "Florești"
	case SubdivisionMDGA:
		return "Găgăuzia, Unitatea teritorială autonomă"
	case SubdivisionMDGL:
		return "Glodeni"
	case SubdivisionMDHI:
		return "Hîncești"
	case SubdivisionMDIA:
		return "Ialoveni"
	case SubdivisionMDLE:
		return "Leova"
	case SubdivisionMDNI:
		return "Nisporeni"
	case SubdivisionMDOC:
		return "Ocnița"
	case SubdivisionMDOR:
		return "Orhei"
	case SubdivisionMDRE:
		return "Rezina"
	case SubdivisionMDRI:
		return "Rîșcani"
	case SubdivisionMDSD:
		return "Șoldănești"
	case SubdivisionMDSI:
		return "Sîngerei"
	case SubdivisionMDSN:
		return "Stînga Nistrului, unitatea teritorială din"
	case SubdivisionMDSO:
		return "Soroca"
	case SubdivisionMDST:
		return "Strășeni"
	case SubdivisionMDSV:
		return "Ștefan Vodă"
	case SubdivisionMDTA:
		return "Taraclia"
	case SubdivisionMDTE:
		return "Telenești"
	case SubdivisionMDUN:
		return "Ungheni"
	case SubdivisionME01:
		return "Andrijevica"
	case SubdivisionME02:
		return "Bar"
	case SubdivisionME03:
		return "Berane"
	case SubdivisionME04:
		return "Bijelo Polje"
	case SubdivisionME05:
		return "Budva"
	case SubdivisionME06:
		return "Cetinje"
	case SubdivisionME07:
		return "Danilovgrad"
	case SubdivisionME08:
		return "Herceg-Novi"
	case SubdivisionME09:
		return "Kolašin"
	case SubdivisionME10:
		return "Kotor"
	case SubdivisionME11:
		return "Mojkovac"
	case SubdivisionME12:
		return "Nikšić"
	case SubdivisionME13:
		return "Plav"
	case SubdivisionME14:
		return "Pljevlja"
	case SubdivisionME15:
		return "Plužine"
	case SubdivisionME16:
		return "Podgorica"
	case SubdivisionME17:
		return "Rožaje"
	case SubdivisionME18:
		return "Šavnik"
	case SubdivisionME19:
		return "Tivat"
	case SubdivisionME20:
		return "Ulcinj"
	case SubdivisionME21:
		return "Žabljak"
	case SubdivisionMGA:
		return "Toamasina"
	case SubdivisionMGD:
		return "Antsiranana"
	case SubdivisionMGF:
		return "Fianarantsoa"
	case SubdivisionMGM:
		return "Mahajanga"
	case SubdivisionMGT:
		return "Antananarivo"
	case SubdivisionMGU:
		return "Toliara"
	case SubdivisionMHALK:
		return "Ailuk"
	case SubdivisionMHALL:
		return "Ailinglaplap"
	case SubdivisionMHARN:
		return "Arno"
	case SubdivisionMHAUR:
		return "Aur"
	case SubdivisionMHEBO:
		return "Ebon"
	case SubdivisionMHENI:
		return "Enewetak"
	case SubdivisionMHJAB:
		return "Jabat"
	case SubdivisionMHJAL:
		return "Jaluit"
	case SubdivisionMHKIL:
		return "Kili"
	case SubdivisionMHKWA:
		return "Kwajalein"
	case SubdivisionMHL:
		return "Ralik chain"
	case SubdivisionMHLAE:
		return "Lae"
	case SubdivisionMHLIB:
		return "Lib"
	case SubdivisionMHLIK:
		return "Likiep"
	case SubdivisionMHMAJ:
		return "Majuro"
	case SubdivisionMHMAL:
		return "Maloelap"
	case SubdivisionMHMEJ:
		return "Mejit"
	case SubdivisionMHMIL:
		return "Mili"
	case SubdivisionMHNMK:
		return "Namdrik"
	case SubdivisionMHNMU:
		return "Namu"
	case SubdivisionMHRON:
		return "Rongelap"
	case SubdivisionMHT:
		return "Ratak chain"
	case SubdivisionMHUJA:
		return "Ujae"
	case SubdivisionMHUTI:
		return "Utirik"
	case SubdivisionMHWTJ:
		return "Wotje"
	case SubdivisionMHWTN:
		return "Wotho"
	case SubdivisionMK01:
		return "Aerodrom"
	case SubdivisionMK02:
		return "Aračinovo"
	case SubdivisionMK03:
		return "Berovo"
	case SubdivisionMK04:
		return "Bitola"
	case SubdivisionMK05:
		return "Bogdanci"
	case SubdivisionMK06:
		return "Bogovinje"
	case SubdivisionMK07:
		return "Bosilovo"
	case SubdivisionMK08:
		return "Brvenica"
	case SubdivisionMK09:
		return "Butel"
	case SubdivisionMK10:
		return "Valandovo"
	case SubdivisionMK11:
		return "Vasilevo"
	case SubdivisionMK12:
		return "Vevčani"
	case SubdivisionMK13:
		return "Veles"
	case SubdivisionMK14:
		return "Vinica"
	case SubdivisionMK15:
		return "Vraneštica"
	case SubdivisionMK16:
		return "Vrapčište"
	case SubdivisionMK17:
		return "Gazi Baba"
	case SubdivisionMK18:
		return "Gevgelija"
	case SubdivisionMK19:
		return "Gostivar"
	case SubdivisionMK20:
		return "Gradsko"
	case SubdivisionMK21:
		return "Debar"
	case SubdivisionMK22:
		return "Debarca"
	case SubdivisionMK23:
		return "Delčevo"
	case SubdivisionMK24:
		return "Demir Kapija"
	case SubdivisionMK25:
		return "Demir Hisar"
	case SubdivisionMK26:
		return "Dojran"
	case SubdivisionMK27:
		return "Dolneni"
	case SubdivisionMK28:
		return "Drugovo"
	case SubdivisionMK29:
		return "Gjorče Petrov"
	case SubdivisionMK30:
		return "Želino"
	case SubdivisionMK31:
		return "Zajas"
	case SubdivisionMK32:
		return "Zelenikovo"
	case SubdivisionMK33:
		return "Zrnovci"
	case SubdivisionMK34:
		return "Ilinden"
	case SubdivisionMK35:
		return "Jegunovce"
	case SubdivisionMK36:
		return "Kavadarci"
	case SubdivisionMK37:
		return "Karbinci"
	case SubdivisionMK38:
		return "Karpoš"
	case SubdivisionMK39:
		return "Kisela Voda"
	case SubdivisionMK40:
		return "Kičevo"
	case SubdivisionMK41:
		return "Konče"
	case SubdivisionMK42:
		return "Kočani"
	case SubdivisionMK43:
		return "Kratovo"
	case SubdivisionMK44:
		return "Kriva Palanka"
	case SubdivisionMK45:
		return "Krivogaštani"
	case SubdivisionMK46:
		return "Kruševo"
	case SubdivisionMK47:
		return "Kumanovo"
	case SubdivisionMK48:
		return "Lipkovo"
	case SubdivisionMK49:
		return "Lozovo"
	case SubdivisionMK50:
		return "Mavrovo-i-Rostuša"
	case SubdivisionMK51:
		return "Makedonska Kamenica"
	case SubdivisionMK52:
		return "Makedonski Brod"
	case SubdivisionMK53:
		return "Mogila"
	case SubdivisionMK54:
		return "Negotino"
	case SubdivisionMK55:
		return "Novaci"
	case SubdivisionMK56:
		return "Novo Selo"
	case SubdivisionMK57:
		return "Oslomej"
	case SubdivisionMK58:
		return "Ohrid"
	case SubdivisionMK59:
		return "Petrovec"
	case SubdivisionMK60:
		return "Pehčevo"
	case SubdivisionMK61:
		return "Plasnica"
	case SubdivisionMK62:
		return "Prilep"
	case SubdivisionMK63:
		return "Probištip"
	case SubdivisionMK64:
		return "Radoviš"
	case SubdivisionMK65:
		return "Rankovce"
	case SubdivisionMK66:
		return "Resen"
	case SubdivisionMK67:
		return "Rosoman"
	case SubdivisionMK68:
		return "Saraj"
	case SubdivisionMK69:
		return "Sveti Nikole"
	case SubdivisionMK70:
		return "Sopište"
	case SubdivisionMK71:
		return "Staro Nagoričane"
	case SubdivisionMK72:
		return "Struga"
	case SubdivisionMK73:
		return "Strumica"
	case SubdivisionMK74:
		return "Studeničani"
	case SubdivisionMK75:
		return "Tearce"
	case SubdivisionMK76:
		return "Tetovo"
	case SubdivisionMK77:
		return "Centar"
	case SubdivisionMK78:
		return "Centar Župa"
	case SubdivisionMK79:
		return "Čair"
	case SubdivisionMK80:
		return "Čaška"
	case SubdivisionMK81:
		return "Češinovo-Obleševo"
	case SubdivisionMK82:
		return "Čučer Sandevo"
	case SubdivisionMK83:
		return "Štip"
	case SubdivisionMK84:
		return "Šuto Orizari"
	case SubdivisionML1:
		return "Kayes"
	case SubdivisionML2:
		return "Koulikoro"
	case SubdivisionML3:
		return "Sikasso"
	case SubdivisionML4:
		return "Ségou"
	case SubdivisionML5:
		return "Mopti"
	case SubdivisionML6:
		return "Tombouctou"
	case SubdivisionML7:
		return "Gao"
	case SubdivisionML8:
		return "Kidal"
	case SubdivisionMLBK0:
		return "Bamako"
	case SubdivisionMM01:
		return "Sagaing"
	case SubdivisionMM02:
		return "Bago"
	case SubdivisionMM03:
		return "Magway"
	case SubdivisionMM04:
		return "Mandalay"
	case SubdivisionMM05:
		return "Tanintharyi"
	case SubdivisionMM06:
		return "Yangon"
	case SubdivisionMM07:
		return "Ayeyarwady"
	case SubdivisionMM11:
		return "Kachin"
	case SubdivisionMM12:
		return "Kayah"
	case SubdivisionMM13:
		return "Kayin"
	case SubdivisionMM14:
		return "Chin"
	case SubdivisionMM15:
		return "Mon"
	case SubdivisionMM16:
		return "Rakhine"
	case SubdivisionMM17:
		return "Shan"
	case SubdivisionMN035:
		return "Orhon"
	case SubdivisionMN037:
		return "Darhan uul"
	case SubdivisionMN039:
		return "Hentiy"
	case SubdivisionMN041:
		return "Hövsgöl"
	case SubdivisionMN043:
		return "Hovd"
	case SubdivisionMN046:
		return "Uvs"
	case SubdivisionMN047:
		return "Töv"
	case SubdivisionMN049:
		return "Selenge"
	case SubdivisionMN051:
		return "Sühbaatar"
	case SubdivisionMN053:
		return "Ömnögovi"
	case SubdivisionMN055:
		return "Övörhangay"
	case SubdivisionMN057:
		return "Dzavhan"
	case SubdivisionMN059:
		return "Dundgovi"
	case SubdivisionMN061:
		return "Dornod"
	case SubdivisionMN063:
		return "Dornogovi"
	case SubdivisionMN064:
		return "Govi-Sumber"
	case SubdivisionMN065:
		return "Govi-Altay"
	case SubdivisionMN067:
		return "Bulgan"
	case SubdivisionMN069:
		return "Bayanhongor"
	case SubdivisionMN071:
		return "Bayan-Ölgiy"
	case SubdivisionMN073:
		return "Arhangay"
	case SubdivisionMN1:
		return "Ulanbaatar"
	case SubdivisionMR01:
		return "Hodh ech Chargui"
	case SubdivisionMR02:
		return "Hodh el Charbi"
	case SubdivisionMR03:
		return "Assaba"
	case SubdivisionMR04:
		return "Gorgol"
	case SubdivisionMR05:
		return "Brakna"
	case SubdivisionMR06:
		return "Trarza"
	case SubdivisionMR07:
		return "Adrar"
	case SubdivisionMR08:
		return "Dakhlet Nouadhibou"
	case SubdivisionMR09:
		return "Tagant"
	case SubdivisionMR10:
		return "Guidimaka"
	case SubdivisionMR11:
		return "Tiris Zemmour"
	case SubdivisionMR12:
		return "Inchiri"
	case SubdivisionMRNKC:
		return "Nouakchott"
	case SubdivisionMT01:
		return "Attard"
	case SubdivisionMT02:
		return "Balzan"
	case SubdivisionMT03:
		return "Birgu"
	case SubdivisionMT04:
		return "Birkirkara"
	case SubdivisionMT05:
		return "Birżebbuġa"
	case SubdivisionMT06:
		return "Bormla"
	case SubdivisionMT07:
		return "Dingli"
	case SubdivisionMT08:
		return "Fgura"
	case SubdivisionMT09:
		return "Floriana"
	case SubdivisionMT10:
		return "Fontana"
	case SubdivisionMT11:
		return "Gudja"
	case SubdivisionMT12:
		return "Gżira"
	case SubdivisionMT13:
		return "Għajnsielem"
	case SubdivisionMT14:
		return "Għarb"
	case SubdivisionMT15:
		return "Għargħur"
	case SubdivisionMT16:
		return "Għasri"
	case SubdivisionMT17:
		return "Għaxaq"
	case SubdivisionMT18:
		return "Ħamrun"
	case SubdivisionMT19:
		return "Iklin"
	case SubdivisionMT20:
		return "Isla"
	case SubdivisionMT21:
		return "Kalkara"
	case SubdivisionMT22:
		return "Kerċem"
	case SubdivisionMT23:
		return "Kirkop"
	case SubdivisionMT24:
		return "Lija"
	case SubdivisionMT25:
		return "Luqa"
	case SubdivisionMT26:
		return "Marsa"
	case SubdivisionMT27:
		return "Marsaskala"
	case SubdivisionMT28:
		return "Marsaxlokk"
	case SubdivisionMT29:
		return "Mdina"
	case SubdivisionMT30:
		return "Mellieħa"
	case SubdivisionMT31:
		return "Mġarr"
	case SubdivisionMT32:
		return "Mosta"
	case SubdivisionMT33:
		return "Mqabba"
	case SubdivisionMT34:
		return "Msida"
	case SubdivisionMT35:
		return "Mtarfa"
	case SubdivisionMT36:
		return "Munxar"
	case SubdivisionMT37:
		return "Nadur"
	case SubdivisionMT38:
		return "Naxxar"
	case SubdivisionMT39:
		return "Paola"
	case SubdivisionMT40:
		return "Pembroke"
	case SubdivisionMT41:
		return "Pietà"
	case SubdivisionMT42:
		return "Qala"
	case SubdivisionMT43:
		return "Qormi"
	case SubdivisionMT44:
		return "Qrendi"
	case SubdivisionMT45:
		return "Rabat Għawdex"
	case SubdivisionMT46:
		return "Rabat Malta"
	case SubdivisionMT47:
		return "Safi"
	case SubdivisionMT48:
		return "San Ġiljan"
	case SubdivisionMT49:
		return "San Ġwann"
	case SubdivisionMT50:
		return "San Lawrenz"
	case SubdivisionMT51:
		return "San Pawl il-Baħar"
	case SubdivisionMT52:
		return "Sannat"
	case SubdivisionMT53:
		return "Santa Luċija"
	case SubdivisionMT54:
		return "Santa Venera"
	case SubdivisionMT55:
		return "Siġġiewi"
	case SubdivisionMT56:
		return "Sliema"
	case SubdivisionMT57:
		return "Swieqi"
	case SubdivisionMT58:
		return "Ta’ Xbiex"
	case SubdivisionMT59:
		return "Tarxien"
	case SubdivisionMT60:
		return "Valletta"
	case SubdivisionMT61:
		return "Xagħra"
	case SubdivisionMT62:
		return "Xewkija"
	case SubdivisionMT63:
		return "Xgħajra"
	case SubdivisionMT64:
		return "Żabbar"
	case SubdivisionMT65:
		return "Żebbuġ Għawdex"
	case SubdivisionMT66:
		return "Żebbuġ Malta"
	case SubdivisionMT67:
		return "Żejtun"
	case SubdivisionMT68:
		return "Żurrieq"
	case SubdivisionMUAG:
		return "Agalega Islands"
	case SubdivisionMUBL:
		return "Black River"
	case SubdivisionMUBR:
		return "Beau Bassin-Rose Hill"
	case SubdivisionMUCC:
		return "Cargados Carajos Shoals"
	case SubdivisionMUCU:
		return "Curepipe"
	case SubdivisionMUFL:
		return "Flacq"
	case SubdivisionMUGP:
		return "Grand Port"
	case SubdivisionMUMO:
		return "Moka"
	case SubdivisionMUPA:
		return "Pamplemousses"
	case SubdivisionMUPL:
		return "Port Louis"
	case SubdivisionMUPU:
		return "Port Louis"
	case SubdivisionMUPW:
		return "Plaines Wilhems"
	case SubdivisionMUQB:
		return "Quatre Bornes"
	case SubdivisionMURO:
		return "Rodrigues Island"
	case SubdivisionMURP:
		return "Rivière du Rempart"
	case SubdivisionMUSA:
		return "Savanne"
	case SubdivisionMUVP:
		return "Vacoas-Phoenix"
	case SubdivisionMV00:
		return "Alifu Dhaalu"
	case SubdivisionMV01:
		return "Seenu"
	case SubdivisionMV02:
		return "Alifu Alifu"
	case SubdivisionMV03:
		return "Lhaviyani"
	case SubdivisionMV04:
		return "Vaavu"
	case SubdivisionMV05:
		return "Laamu"
	case SubdivisionMV07:
		return "Haa Alifu"
	case SubdivisionMV08:
		return "Thaa"
	case SubdivisionMV12:
		return "Meemu"
	case SubdivisionMV13:
		return "Raa"
	case SubdivisionMV14:
		return "Faafu"
	case SubdivisionMV17:
		return "Dhaalu"
	case SubdivisionMV20:
		return "Baa"
	case SubdivisionMV23:
		return "Haa Dhaalu"
	case SubdivisionMV24:
		return "Shaviyani"
	case SubdivisionMV25:
		return "Noonu"
	case SubdivisionMV26:
		return "Kaafu"
	case SubdivisionMV27:
		return "Gaafu Alifu"
	case SubdivisionMV28:
		return "Gaafu Dhaalu"
	case SubdivisionMV29:
		return "Gnaviyani"
	case SubdivisionMVCE:
		return "Central"
	case SubdivisionMVMLE:
		return "Male"
	case SubdivisionMVNC:
		return "North Central"
	case SubdivisionMVNO:
		return "North"
	case SubdivisionMVSC:
		return "South Central"
	case SubdivisionMVSU:
		return "South"
	case SubdivisionMVUN:
		return "Upper North"
	case SubdivisionMVUS:
		return "Upper South"
	case SubdivisionMWBA:
		return "Balaka"
	case SubdivisionMWBL:
		return "Blantyre"
	case SubdivisionMWC:
		return "Central Region"
	case SubdivisionMWCK:
		return "Chikwawa"
	case SubdivisionMWCR:
		return "Chiradzulu"
	case SubdivisionMWCT:
		return "Chitipa"
	case SubdivisionMWDE:
		return "Dedza"
	case SubdivisionMWDO:
		return "Dowa"
	case SubdivisionMWKR:
		return "Karonga"
	case SubdivisionMWKS:
		return "Kasungu"
	case SubdivisionMWLI:
		return "Lilongwe"
	case SubdivisionMWLK:
		return "Likoma"
	case SubdivisionMWMC:
		return "Mchinji"
	case SubdivisionMWMG:
		return "Mangochi"
	case SubdivisionMWMH:
		return "Machinga"
	case SubdivisionMWMU:
		return "Mulanje"
	case SubdivisionMWMW:
		return "Mwanza"
	case SubdivisionMWMZ:
		return "Mzimba"
	case SubdivisionMWN:
		return "Northern Region"
	case SubdivisionMWNB:
		return "Nkhata Bay"
	case SubdivisionMWNE:
		return "Neno"
	case SubdivisionMWNI:
		return "Ntchisi"
	case SubdivisionMWNK:
		return "Nkhotakota"
	case SubdivisionMWNS:
		return "Nsanje"
	case SubdivisionMWNU:
		return "Ntcheu"
	case SubdivisionMWPH:
		return "Phalombe"
	case SubdivisionMWRU:
		return "Rumphi"
	case SubdivisionMWS:
		return "Southern Region"
	case SubdivisionMWSA:
		return "Salima"
	case SubdivisionMWTH:
		return "Thyolo"
	case SubdivisionMWZO:
		return "Zomba"
	case SubdivisionMXAGU:
		return "Aguascalientes"
	case SubdivisionMXBCN:
		return "Baja California"
	case SubdivisionMXBCS:
		return "Baja California Sur"
	case SubdivisionMXCAM:
		return "Campeche"
	case SubdivisionMXCHH:
		return "Chihuahua"
	case SubdivisionMXCHP:
		return "Chiapas"
	case SubdivisionMXCMX:
		return "Ciudad de México"
	case SubdivisionMXCOA:
		return "Coahuila de Zaragoza"
	case SubdivisionMXCOL:
		return "Colima"
	case SubdivisionMXDUR:
		return "Durango"
	case SubdivisionMXGRO:
		return "Guerrero"
	case SubdivisionMXGUA:
		return "Guanajuato"
	case SubdivisionMXHID:
		return "Hidalgo"
	case SubdivisionMXJAL:
		return "Jalisco"
	case SubdivisionMXMEX:
		return "México"
	case SubdivisionMXMIC:
		return "Michoacán de Ocampo"
	case SubdivisionMXMOR:
		return "Morelos"
	case SubdivisionMXNAY:
		return "Nayarit"
	case SubdivisionMXNLE:
		return "Nuevo León"
	case SubdivisionMXOAX:
		return "Oaxaca"
	case SubdivisionMXPUE:
		return "Puebla"
	case SubdivisionMXQUE:
		return "Querétaro"
	case SubdivisionMXROO:
		return "Quintana Roo"
	case SubdivisionMXSIN:
		return "Sinaloa"
	case SubdivisionMXSLP:
		return "San Luis Potosí"
	case SubdivisionMXSON:
		return "Sonora"
	case SubdivisionMXTAB:
		return "Tabasco"
	case SubdivisionMXTAM:
		return "Tamaulipas"
	case SubdivisionMXTLA:
		return "Tlaxcala"
	case SubdivisionMXVER:
		return "Veracruz de Ignacio de la Llave"
	case SubdivisionMXYUC:
		return "Yucatán"
	case SubdivisionMXZAC:
		return "Zacatecas"
	case SubdivisionMY01:
		return "Johor"
	case SubdivisionMY02:
		return "Kedah"
	case SubdivisionMY03:
		return "Kelantan"
	case SubdivisionMY04:
		return "Melaka"
	case SubdivisionMY05:
		return "Negeri Sembilan"
	case SubdivisionMY06:
		return "Pahang"
	case SubdivisionMY07:
		return "Pulau Pinang"
	case SubdivisionMY08:
		return "Perak"
	case SubdivisionMY09:
		return "Perlis"
	case SubdivisionMY10:
		return "Selangor"
	case SubdivisionMY11:
		return "Terengganu"
	case SubdivisionMY12:
		return "Sabah"
	case SubdivisionMY13:
		return "Sarawak"
	case SubdivisionMY14:
		return "Wilayah Persekutuan Kuala Lumpur"
	case SubdivisionMY15:
		return "Wilayah Persekutuan Labuan"
	case SubdivisionMY16:
		return "Wilayah Persekutuan Putrajaya"
	case SubdivisionMZA:
		return "Niassa"
	case SubdivisionMZB:
		return "Manica"
	case SubdivisionMZG:
		return "Gaza"
	case SubdivisionMZI:
		return "Inhambane"
	case SubdivisionMZL:
		return "Maputo"
	case SubdivisionMZMPM:
		return "Maputo (city)"
	case SubdivisionMZN:
		return "Numpula"
	case SubdivisionMZP:
		return "Cabo Delgado"
	case SubdivisionMZQ:
		return "Zambezia"
	case SubdivisionMZS:
		return "Sofala"
	case SubdivisionMZT:
		return "Tete"
	case SubdivisionNACA:
		return "Caprivi"
	case SubdivisionNAER:
		return "Erongo"
	case SubdivisionNAHA:
		return "Hardap"
	case SubdivisionNAKA:
		return "Karas"
	case SubdivisionNAKH:
		return "Khomas"
	case SubdivisionNAKU:
		return "Kunene"
	case SubdivisionNAOD:
		return "Otjozondjupa"
	case SubdivisionNAOH:
		return "Omaheke"
	case SubdivisionNAOK:
		return "Okavango"
	case SubdivisionNAON:
		return "Oshana"
	case SubdivisionNAOS:
		return "Omusati"
	case SubdivisionNAOT:
		return "Oshikoto"
	case SubdivisionNAOW:
		return "Ohangwena"
	case SubdivisionNE1:
		return "Agadez"
	case SubdivisionNE2:
		return "Diffa"
	case SubdivisionNE3:
		return "Dosso"
	case SubdivisionNE4:
		return "Maradi"
	case SubdivisionNE5:
		return "Tahoua"
	case SubdivisionNE6:
		return "Tillabéri"
	case SubdivisionNE7:
		return "Zinder"
	case SubdivisionNE8:
		return "Niamey"
	case SubdivisionNGAB:
		return "Abia"
	case SubdivisionNGAD:
		return "Adamawa"
	case SubdivisionNGAK:
		return "Akwa Ibom"
	case SubdivisionNGAN:
		return "Anambra"
	case SubdivisionNGBA:
		return "Bauchi"
	case SubdivisionNGBE:
		return "Benue"
	case SubdivisionNGBO:
		return "Borno"
	case SubdivisionNGBY:
		return "Bayelsa"
	case SubdivisionNGCR:
		return "Cross River"
	case SubdivisionNGDE:
		return "Delta"
	case SubdivisionNGEB:
		return "Ebonyi"
	case SubdivisionNGED:
		return "Edo"
	case SubdivisionNGEK:
		return "Ekiti"
	case SubdivisionNGEN:
		return "Enugu"
	case SubdivisionNGFC:
		return "Abuja Capital Territory"
	case SubdivisionNGGO:
		return "Gombe"
	case SubdivisionNGIM:
		return "Imo"
	case SubdivisionNGJI:
		return "Jigawa"
	case SubdivisionNGKD:
		return "Kaduna"
	case SubdivisionNGKE:
		return "Kebbi"
	case SubdivisionNGKN:
		return "Kano"
	case SubdivisionNGKO:
		return "Kogi"
	case SubdivisionNGKT:
		return "Katsina"
	case SubdivisionNGKW:
		return "Kwara"
	case SubdivisionNGLA:
		return "Lagos"
	case SubdivisionNGNA:
		return "Nassarawa"
	case SubdivisionNGNI:
		return "Niger"
	case SubdivisionNGOG:
		return "Ogun"
	case SubdivisionNGON:
		return "Ondo"
	case SubdivisionNGOS:
		return "Osun"
	case SubdivisionNGOY:
		return "Oyo"
	case SubdivisionNGPL:
		return "Plateau"
	case SubdivisionNGRI:
		return "Rivers"
	case SubdivisionNGSO:
		return "Sokoto"
	case SubdivisionNGTA:
		return "Taraba"
	case SubdivisionNGYO:
		return "Yobe"
	case SubdivisionNGZA:
		return "Zamfara"
	case SubdivisionNIAN:
		return "Atlántico Norte"
	case SubdivisionNIAS:
		return "Atlántico Sur"
	case SubdivisionNIBO:
		return "Boaco"
	case SubdivisionNICA:
		return "Carazo"
	case SubdivisionNICI:
		return "Chinandega"
	case SubdivisionNICO:
		return "Chontales"
	case SubdivisionNIES:
		return "Estelí"
	case SubdivisionNIGR:
		return "Granada"
	case SubdivisionNIJI:
		return "Jinotega"
	case SubdivisionNILE:
		return "León"
	case SubdivisionNIMD:
		return "Madriz"
	case SubdivisionNIMN:
		return "Managua"
	case SubdivisionNIMS:
		return "Masaya"
	case SubdivisionNIMT:
		return "Matagalpa"
	case SubdivisionNINS:
		return "Nueva Segovia"
	case SubdivisionNIRI:
		return "Rivas"
	case SubdivisionNISJ:
		return "Río San Juan"
	case SubdivisionNLAW:
		return "Aruba"
	case SubdivisionNLBQ1:
		return "Bonaire"
	case SubdivisionNLBQ2:
		return "Saba"
	case SubdivisionNLBQ3:
		return "Sint Eustatius"
	case SubdivisionNLCW:
		return "Curaçao"
	case SubdivisionNLDR:
		return "Drenthe"
	case SubdivisionNLFL:
		return "Flevoland"
	case SubdivisionNLFR:
		return "Friesland"
	case SubdivisionNLGE:
		return "Gelderland"
	case SubdivisionNLGR:
		return "Groningen"
	case SubdivisionNLLI:
		return "Limburg"
	case SubdivisionNLNB:
		return "Noord-Brabant"
	case SubdivisionNLNH:
		return "Noord-Holland"
	case SubdivisionNLOV:
		return "Overijssel"
	case SubdivisionNLSX:
		return "Sint Maarten"
	case SubdivisionNLUT:
		return "Utrecht"
	case SubdivisionNLZE:
		return "Zeeland"
	case SubdivisionNLZH:
		return "Zuid-Holland"
	case SubdivisionNO01:
		return "Østfold"
	case SubdivisionNO02:
		return "Akershus"
	case SubdivisionNO03:
		return "Oslo"
	case SubdivisionNO04:
		return "Hedmark"
	case SubdivisionNO05:
		return "Oppland"
	case SubdivisionNO06:
		return "Buskerud"
	case SubdivisionNO07:
		return "Vestfold"
	case SubdivisionNO08:
		return "Telemark"
	case SubdivisionNO09:
		return "Aust-Agder"
	case SubdivisionNO10:
		return "Vest-Agder"
	case SubdivisionNO11:
		return "Rogaland"
	case SubdivisionNO12:
		return "Hordaland"
	case SubdivisionNO14:
		return "Sogn og Fjordane"
	case SubdivisionNO15:
		return "Møre og Romsdal"
	case SubdivisionNO18:
		return "Nordland"
	case SubdivisionNO19:
		return "Troms"
	case SubdivisionNO20:
		return "Finnmark"
	case SubdivisionNO21:
		return "Svalbard (Arctic Region)"
	case SubdivisionNO22:
		return "Jan Mayen (Arctic Region)"
	case SubdivisionNO50:
		return "Trøndelag"
	case SubdivisionNP1:
		return "Madhyamanchal"
	case SubdivisionNP2:
		return "Madhya Pashchimanchal"
	case SubdivisionNP3:
		return "Pashchimanchal"
	case SubdivisionNP4:
		return "Purwanchal"
	case SubdivisionNP5:
		return "Sudur Pashchimanchal"
	case SubdivisionNPBA:
		return "Bagmati"
	case SubdivisionNPBH:
		return "Bheri"
	case SubdivisionNPDH:
		return "Dhawalagiri"
	case SubdivisionNPGA:
		return "Gandaki"
	case SubdivisionNPJA:
		return "Janakpur"
	case SubdivisionNPKA:
		return "Karnali"
	case SubdivisionNPKO:
		return "Kosi"
	case SubdivisionNPLU:
		return "Lumbini"
	case SubdivisionNPMA:
		return "Mahakali"
	case SubdivisionNPME:
		return "Mechi"
	case SubdivisionNPNA:
		return "Narayani"
	case SubdivisionNPRA:
		return "Rapti"
	case SubdivisionNPSA:
		return "Sagarmatha"
	case SubdivisionNPSE:
		return "Seti"
	case SubdivisionNR01:
		return "Aiwo"
	case SubdivisionNR02:
		return "Anabar"
	case SubdivisionNR03:
		return "Anetan"
	case SubdivisionNR04:
		return "Anibare"
	case SubdivisionNR05:
		return "Baiti"
	case SubdivisionNR06:
		return "Boe"
	case SubdivisionNR07:
		return "Buada"
	case SubdivisionNR08:
		return "Denigomodu"
	case SubdivisionNR09:
		return "Ewa"
	case SubdivisionNR10:
		return "Ijuw"
	case SubdivisionNR11:
		return "Meneng"
	case SubdivisionNR12:
		return "Nibok"
	case SubdivisionNR13:
		return "Uaboe"
	case SubdivisionNR14:
		return "Yaren"
	case SubdivisionNZAUK:
		return "Auckland"
	case SubdivisionNZBOP:
		return "Bay of Plenty"
	case SubdivisionNZCAN:
		return "Canterbury"
	case SubdivisionNZCIT:
		return "Chatham Islands Territory"
	case SubdivisionNZGIS:
		return "Gisborne District"
	case SubdivisionNZHKB:
		return "Hawke's Bay"
	case SubdivisionNZMBH:
		return "Marlborough District"
	case SubdivisionNZMWT:
		return "Manawatu-Wanganui"
	case SubdivisionNZN:
		return "North Island"
	case SubdivisionNZNSN:
		return "Nelson City"
	case SubdivisionNZNTL:
		return "Northland"
	case SubdivisionNZOTA:
		return "Otago"
	case SubdivisionNZS:
		return "South Island"
	case SubdivisionNZSTL:
		return "Southland"
	case SubdivisionNZTAS:
		return "Tasman District"
	case SubdivisionNZTKI:
		return "Taranaki"
	case SubdivisionNZWGN:
		return "Wellington"
	case SubdivisionNZWKO:
		return "Waikato"
	case SubdivisionNZWTC:
		return "West Coast"
	case SubdivisionOMBA:
		return "Al Bāţinah"
	case SubdivisionOMBU:
		return "Al Buraymī"
	case SubdivisionOMDA:
		return "Ad Dākhilīya"
	case SubdivisionOMMA:
		return "Masqaţ"
	case SubdivisionOMMU:
		return "Musandam"
	case SubdivisionOMSH:
		return "Ash Sharqīyah"
	case SubdivisionOMWU:
		return "Al Wusţá"
	case SubdivisionOMZA:
		return "Az̧ Z̧āhirah"
	case SubdivisionOMZU:
		return "Z̧ufār"
	case SubdivisionPA1:
		return "Bocas del Toro"
	case SubdivisionPA2:
		return "Coclé"
	case SubdivisionPA3:
		return "Colón"
	case SubdivisionPA4:
		return "Chiriquí"
	case SubdivisionPA5:
		return "Darién"
	case SubdivisionPA6:
		return "Herrera"
	case SubdivisionPA7:
		return "Los Santos"
	case SubdivisionPA8:
		return "Panamá"
	case SubdivisionPA9:
		return "Veraguas"
	case SubdivisionPAEM:
		return "Emberá"
	case SubdivisionPAKY:
		return "Kuna Yala"
	case SubdivisionPANB:
		return "Ngöbe-Buglé"
	case SubdivisionPEAMA:
		return "Amazonas"
	case SubdivisionPEANC:
		return "Ancash"
	case SubdivisionPEAPU:
		return "Apurímac"
	case SubdivisionPEARE:
		return "Arequipa"
	case SubdivisionPEAYA:
		return "Ayacucho"
	case SubdivisionPECAJ:
		return "Cajamarca"
	case SubdivisionPECAL:
		return "El Callao"
	case SubdivisionPECUS:
		return "Cusco [Cuzco]"
	case SubdivisionPEHUC:
		return "Huánuco"
	case SubdivisionPEHUV:
		return "Huancavelica"
	case SubdivisionPEICA:
		return "Ica"
	case SubdivisionPEJUN:
		return "Junín"
	case SubdivisionPELAL:
		return "La Libertad"
	case SubdivisionPELAM:
		return "Lambayeque"
	case SubdivisionPELIM:
		return "Lima"
	case SubdivisionPELMA:
		return "Municipalidad Metropolitana de Lima"
	case SubdivisionPELOR:
		return "Loreto"
	case SubdivisionPEMDD:
		return "Madre de Dios"
	case SubdivisionPEMOQ:
		return "Moquegua"
	case SubdivisionPEPAS:
		return "Pasco"
	case SubdivisionPEPIU:
		return "Piura"
	case SubdivisionPEPUN:
		return "Puno"
	case SubdivisionPESAM:
		return "San Martín"
	case SubdivisionPETAC:
		return "Tacna"
	case SubdivisionPETUM:
		return "Tumbes"
	case SubdivisionPEUCA:
		return "Ucayali"
	case SubdivisionPGCPK:
		return "Chimbu"
	case SubdivisionPGCPM:
		return "Central"
	case SubdivisionPGEBR:
		return "East New Britain"
	case SubdivisionPGEHG:
		return "Eastern Highlands"
	case SubdivisionPGEPW:
		return "Enga"
	case SubdivisionPGESW:
		return "East Sepik"
	case SubdivisionPGGPK:
		return "Gulf"
	case SubdivisionPGMBA:
		return "Milne Bay"
	case SubdivisionPGMPL:
		return "Morobe"
	case SubdivisionPGMPM:
		return "Madang"
	case SubdivisionPGMRL:
		return "Manus"
	case SubdivisionPGNCD:
		return "National Capital District (Port Moresby)"
	case SubdivisionPGNIK:
		return "New Ireland"
	case SubdivisionPGNPP:
		return "Northern"
	case SubdivisionPGNSB:
		return "Bougainville"
	case SubdivisionPGSAN:
		return "Sandaun"
	case SubdivisionPGSHM:
		return "Southern Highlands"
	case SubdivisionPGWBK:
		return "West New Britain"
	case SubdivisionPGWHM:
		return "Western Highlands"
	case SubdivisionPGWPD:
		return "Western"
	case SubdivisionPH00:
		return "National Capital Region"
	case SubdivisionPH01:
		return "Ilocos (Region I)"
	case SubdivisionPH02:
		return "Cagayan Valley (Region II)"
	case SubdivisionPH03:
		return "Central Luzon (Region III)"
	case SubdivisionPH05:
		return "Bicol (Region V)"
	case SubdivisionPH06:
		return "Western Visayas (Region VI)"
	case SubdivisionPH07:
		return "Central Visayas (Region VII)"
	case SubdivisionPH08:
		return "Eastern Visayas (Region VIII)"
	case SubdivisionPH09:
		return "Zamboanga Peninsula (Region IX)"
	case SubdivisionPH10:
		return "Northern Mindanao (Region X)"
	case SubdivisionPH11:
		return "Davao (Region XI)"
	case SubdivisionPH12:
		return "Soccsksargen (Region XII)"
	case SubdivisionPH13:
		return "Caraga (Region XIII)"
	case SubdivisionPH14:
		return "Autonomous Region in Muslim Mindanao (ARMM)"
	case SubdivisionPH15:
		return "Cordillera Administrative Region (CAR)"
	case SubdivisionPH40:
		return "CALABARZON (Region IV-A)"
	case SubdivisionPH41:
		return "MIMAROPA (Region IV-B)"
	case SubdivisionPHABR:
		return "Abra"
	case SubdivisionPHAGN:
		return "Agusan del Norte"
	case SubdivisionPHAGS:
		return "Agusan del Sur"
	case SubdivisionPHAKL:
		return "Aklan"
	case SubdivisionPHALB:
		return "Albay"
	case SubdivisionPHANT:
		return "Antique"
	case SubdivisionPHAPA:
		return "Apayao"
	case SubdivisionPHAUR:
		return "Aurora"
	case SubdivisionPHBAN:
		return "Batasn"
	case SubdivisionPHBAS:
		return "Basilan"
	case SubdivisionPHBEN:
		return "Benguet"
	case SubdivisionPHBIL:
		return "Biliran"
	case SubdivisionPHBOH:
		return "Bohol"
	case SubdivisionPHBTG:
		return "Batangas"
	case SubdivisionPHBTN:
		return "Batanes"
	case SubdivisionPHBUK:
		return "Bukidnon"
	case SubdivisionPHBUL:
		return "Bulacan"
	case SubdivisionPHCAG:
		return "Cagayan"
	case SubdivisionPHCAM:
		return "Camiguin"
	case SubdivisionPHCAN:
		return "Camarines Norte"
	case SubdivisionPHCAP:
		return "Capiz"
	case SubdivisionPHCAS:
		return "Camarines Sur"
	case SubdivisionPHCAT:
		return "Catanduanes"
	case SubdivisionPHCAV:
		return "Cavite"
	case SubdivisionPHCEB:
		return "Cebu"
	case SubdivisionPHCOM:
		return "Compostela Valley"
	case SubdivisionPHDAO:
		return "Davao Oriental"
	case SubdivisionPHDAS:
		return "Davao del Sur"
	case SubdivisionPHDAV:
		return "Davao del Norte"
	case SubdivisionPHDIN:
		return "Dinagat Islands"
	case SubdivisionPHEAS:
		return "Eastern Samar"
	case SubdivisionPHGUI:
		return "Guimaras"
	case SubdivisionPHIFU:
		return "Ifugao"
	case SubdivisionPHILI:
		return "Iloilo"
	case SubdivisionPHILN:
		return "Ilocos Norte"
	case SubdivisionPHILS:
		return "Ilocos Sur"
	case SubdivisionPHISA:
		return "Isabela"
	case SubdivisionPHKAL:
		return "Kalinga-Apayso"
	case SubdivisionPHLAG:
		return "Laguna"
	case SubdivisionPHLAN:
		return "Lanao del Norte"
	case SubdivisionPHLAS:
		return "Lanao del Sur"
	case SubdivisionPHLEY:
		return "Leyte"
	case SubdivisionPHLUN:
		return "La Union"
	case SubdivisionPHMAD:
		return "Marinduque"
	case SubdivisionPHMAG:
		return "Maguindanao"
	case SubdivisionPHMAS:
		return "Masbate"
	case SubdivisionPHMDC:
		return "Mindoro Occidental"
	case SubdivisionPHMDR:
		return "Mindoro Oriental"
	case SubdivisionPHMOU:
		return "Mountain Province"
	case SubdivisionPHMSC:
		return "Misamis Occidental"
	case SubdivisionPHMSR:
		return "Misamis Oriental"
	case SubdivisionPHNCO:
		return "North Cotabato"
	case SubdivisionPHNEC:
		return "Negros Occidental"
	case SubdivisionPHNER:
		return "Negros Oriental"
	case SubdivisionPHNSA:
		return "Northern Samar"
	case SubdivisionPHNUE:
		return "Nueva Ecija"
	case SubdivisionPHNUV:
		return "Nueva Vizcaya"
	case SubdivisionPHPAM:
		return "Pampanga"
	case SubdivisionPHPAN:
		return "Pangasinan"
	case SubdivisionPHPLW:
		return "Palawan"
	case SubdivisionPHQUE:
		return "Quezon"
	case SubdivisionPHQUI:
		return "Quirino"
	case SubdivisionPHRIZ:
		return "Rizal"
	case SubdivisionPHROM:
		return "Romblon"
	case SubdivisionPHSAR:
		return "Sarangani"
	case SubdivisionPHSCO:
		return "South Cotabato"
	case SubdivisionPHSIG:
		return "Siquijor"
	case SubdivisionPHSLE:
		return "Southern Leyte"
	case SubdivisionPHSLU:
		return "Sulu"
	case SubdivisionPHSOR:
		return "Sorsogon"
	case SubdivisionPHSUK:
		return "Sultan Kudarat"
	case SubdivisionPHSUN:
		return "Surigao del Norte"
	case SubdivisionPHSUR:
		return "Surigao del Sur"
	case SubdivisionPHTAR:
		return "Tarlac"
	case SubdivisionPHTAW:
		return "Tawi-Tawi"
	case SubdivisionPHWSA:
		return "Western Samar"
	case SubdivisionPHZAN:
		return "Zamboanga del Norte"
	case SubdivisionPHZAS:
		return "Zamboanga del Sur"
	case SubdivisionPHZMB:
		return "Zambales"
	case SubdivisionPHZSI:
		return "Zamboanga Sibugay"
	case SubdivisionPKBA:
		return "Balochistan"
	case SubdivisionPKGB:
		return "Gilgit-Baltistan"
	case SubdivisionPKIS:
		return "Islamabad"
	case SubdivisionPKJK:
		return "Azad Kashmir"
	case SubdivisionPKKP:
		return "Khyber Pakhtunkhwa"
	case SubdivisionPKPB:
		return "Punjab"
	case SubdivisionPKSD:
		return "Sindh"
	case SubdivisionPKTA:
		return "Federally Administered Tribal Areas"
	case SubdivisionPLDS:
		return "Dolnośląskie"
	case SubdivisionPLKP:
		return "Kujawsko-pomorskie"
	case SubdivisionPLLB:
		return "Lubuskie"
	case SubdivisionPLLD:
		return "Łódzkie"
	case SubdivisionPLLU:
		return "Lubelskie"
	case SubdivisionPLMA:
		return "Małopolskie"
	case SubdivisionPLMZ:
		return "Mazowieckie"
	case SubdivisionPLOP:
		return "Opolskie"
	case SubdivisionPLPD:
		return "Podlaskie"
	case SubdivisionPLPK:
		return "Podkarpackie"
	case SubdivisionPLPM:
		return "Pomorskie"
	case SubdivisionPLSK:
		return "Świętokrzyskie"
	case SubdivisionPLSL:
		return "Śląskie"
	case SubdivisionPLWN:
		return "Warmińsko-mazurskie"
	case SubdivisionPLWP:
		return "Wielkopolskie"
	case SubdivisionPLZP:
		return "Zachodniopomorskie"
	case SubdivisionPSBTH:
		return "Bethlehem"
	case SubdivisionPSDEB:
		return "Deir El Balah"
	case SubdivisionPSGZA:
		return "Gaza"
	case SubdivisionPSHBN:
		return "Hebron"
	case SubdivisionPSJEM:
		return "Jerusalem"
	case SubdivisionPSJEN:
		return "Jenin"
	case SubdivisionPSJRH:
		return "Jericho - Al Aghwar"
	case SubdivisionPSKYS:
		return "Khan Yunis"
	case SubdivisionPSNBS:
		return "Nablus"
	case SubdivisionPSNGZ:
		return "North Gaza"
	case SubdivisionPSQQA:
		return "Qalqilya"
	case SubdivisionPSRBH:
		return "Ramallah"
	case SubdivisionPSRFH:
		return "Rafah"
	case SubdivisionPSSLT:
		return "Salfit"
	case SubdivisionPSTBS:
		return "Tubas"
	case SubdivisionPSTKM:
		return "Tulkarm"
	case SubdivisionPT01:
		return "Aveiro"
	case SubdivisionPT02:
		return "Beja"
	case SubdivisionPT03:
		return "Braga"
	case SubdivisionPT04:
		return "Bragança"
	case SubdivisionPT05:
		return "Castelo Branco"
	case SubdivisionPT06:
		return "Coimbra"
	case SubdivisionPT07:
		return "Évora"
	case SubdivisionPT08:
		return "Faro"
	case SubdivisionPT09:
		return "Guarda"
	case SubdivisionPT10:
		return "Leiria"
	case SubdivisionPT11:
		return "Lisboa"
	case SubdivisionPT12:
		return "Portalegre"
	case SubdivisionPT13:
		return "Porto"
	case SubdivisionPT14:
		return "Santarém"
	case SubdivisionPT15:
		return "Setúbal"
	case SubdivisionPT16:
		return "Viana do Castelo"
	case SubdivisionPT17:
		return "Vila Real"
	case SubdivisionPT18:
		return "Viseu"
	case SubdivisionPT20:
		return "Região Autónoma dos Açores"
	case SubdivisionPT30:
		return "Região Autónoma da Madeira"
	case SubdivisionPW002:
		return "Aimeliik"
	case SubdivisionPW004:
		return "Airai"
	case SubdivisionPW010:
		return "Angaur"
	case SubdivisionPW050:
		return "Hatobohei"
	case SubdivisionPW100:
		return "Kayangel"
	case SubdivisionPW150:
		return "Koror"
	case SubdivisionPW212:
		return "Melekeok"
	case SubdivisionPW214:
		return "Ngaraard"
	case SubdivisionPW218:
		return "Ngarchelong"
	case SubdivisionPW222:
		return "Ngardmau"
	case SubdivisionPW224:
		return "Ngatpang"
	case SubdivisionPW226:
		return "Ngchesar"
	case SubdivisionPW227:
		return "Ngeremlengui"
	case SubdivisionPW228:
		return "Ngiwal"
	case SubdivisionPW350:
		return "Peleliu"
	case SubdivisionPW370:
		return "Sonsorol"
	case SubdivisionPY1:
		return "Concepción"
	case SubdivisionPY10:
		return "Alto Paraná"
	case SubdivisionPY11:
		return "Central"
	case SubdivisionPY12:
		return "Ñeembucú"
	case SubdivisionPY13:
		return "Amambay"
	case SubdivisionPY14:
		return "Canindeyú"
	case SubdivisionPY15:
		return "Presidente Hayes"
	case SubdivisionPY16:
		return "Alto Paraguay"
	case SubdivisionPY19:
		return "Boquerón"
	case SubdivisionPY2:
		return "San Pedro"
	case SubdivisionPY3:
		return "Cordillera"
	case SubdivisionPY4:
		return "Guairá"
	case SubdivisionPY5:
		return "Caaguazú"
	case SubdivisionPY6:
		return "Caazapá"
	case SubdivisionPY7:
		return "Itapúa"
	case SubdivisionPY8:
		return "Misiones"
	case SubdivisionPY9:
		return "Paraguarí"
	case SubdivisionPYASU:
		return "Asunción"
	case SubdivisionQADA:
		return "Ad Dawhah"
	case SubdivisionQAKH:
		return "Al Khawr wa adh Dhakhīrah"
	case SubdivisionQAMS:
		return "Ash Shamal"
	case SubdivisionQARA:
		return "Ar Rayyan"
	case SubdivisionQAUS:
		return "Umm Salal"
	case SubdivisionQAWA:
		return "Al Wakrah"
	case SubdivisionQAZA:
		return "Az̧ Z̧a‘āyin"
	case SubdivisionROAB:
		return "Alba"
	case SubdivisionROAG:
		return "Argeș"
	case SubdivisionROAR:
		return "Arad"
	case SubdivisionROB:
		return "București"
	case SubdivisionROBC:
		return "Bacău"
	case SubdivisionROBH:
		return "Bihor"
	case SubdivisionROBN:
		return "Bistrița-Năsăud"
	case SubdivisionROBR:
		return "Brăila"
	case SubdivisionROBT:
		return "Botoșani"
	case SubdivisionROBV:
		return "Brașov"
	case SubdivisionROBZ:
		return "Buzău"
	case SubdivisionROCJ:
		return "Cluj"
	case SubdivisionROCL:
		return "Călărași"
	case SubdivisionROCS:
		return "Caraș-Severin"
	case SubdivisionROCT:
		return "Constanța"
	case SubdivisionROCV:
		return "Covasna"
	case SubdivisionRODB:
		return "Dâmbovița"
	case SubdivisionRODJ:
		return "Dolj"
	case SubdivisionROGJ:
		return "Gorj"
	case SubdivisionROGL:
		return "Galați"
	case SubdivisionROGR:
		return "Giurgiu"
	case SubdivisionROHD:
		return "Hunedoara"
	case SubdivisionROHR:
		return "Harghita"
	case SubdivisionROIF:
		return "Ilfov"
	case SubdivisionROIL:
		return "Ialomița"
	case SubdivisionROIS:
		return "Iași"
	case SubdivisionROMH:
		return "Mehedinți"
	case SubdivisionROMM:
		return "Maramureș"
	case SubdivisionROMS:
		return "Mureș"
	case SubdivisionRONT:
		return "Neamț"
	case SubdivisionROOT:
		return "Olt"
	case SubdivisionROPH:
		return "Prahova"
	case SubdivisionROSB:
		return "Sibiu"
	case SubdivisionROSJ:
		return "Sălaj"
	case SubdivisionROSM:
		return "Satu Mare"
	case SubdivisionROSV:
		return "Suceava"
	case SubdivisionROTL:
		return "Tulcea"
	case SubdivisionROTM:
		return "Timiș"
	case SubdivisionROTR:
		return "Teleorman"
	case SubdivisionROVL:
		return "Vâlcea"
	case SubdivisionROVN:
		return "Vrancea"
	case SubdivisionROVS:
		return "Vaslui"
	case SubdivisionRS00:
		return "Beograd"
	case SubdivisionRS01:
		return "Severnobački okrug"
	case SubdivisionRS02:
		return "Srednjebanatski okrug"
	case SubdivisionRS03:
		return "Severnobanatski okrug"
	case SubdivisionRS04:
		return "Južnobanatski okrug"
	case SubdivisionRS05:
		return "Zapadnobački okrug"
	case SubdivisionRS06:
		return "Južnobački okrug"
	case SubdivisionRS07:
		return "Sremski okrug"
	case SubdivisionRS08:
		return "Mačvanski okrug"
	case SubdivisionRS09:
		return "Kolubarski okrug"
	case SubdivisionRS10:
		return "Podunavski okrug"
	case SubdivisionRS11:
		return "Braničevski okrug"
	case SubdivisionRS12:
		return "Šumadijski okrug"
	case SubdivisionRS13:
		return "Pomoravski okrug"
	case SubdivisionRS14:
		return "Borski okrug"
	case SubdivisionRS15:
		return "Zaječarski okrug"
	case SubdivisionRS16:
		return "Zlatiborski okrug"
	case SubdivisionRS17:
		return "Moravički okrug"
	case SubdivisionRS18:
		return "Raški okrug"
	case SubdivisionRS19:
		return "Rasinski okrug"
	case SubdivisionRS20:
		return "Nišavski okrug"
	case SubdivisionRS21:
		return "Toplički okrug"
	case SubdivisionRS22:
		return "Pirotski okrug"
	case SubdivisionRS23:
		return "Jablanički okrug"
	case SubdivisionRS24:
		return "Pčinjski okrug"
	case SubdivisionRS25:
		return "Kosovski okrug"
	case SubdivisionRS26:
		return "Pećki okrug"
	case SubdivisionRS27:
		return "Prizrenski okrug"
	case SubdivisionRS28:
		return "Kosovsko-Mitrovački okrug"
	case SubdivisionRS29:
		return "Kosovsko-Pomoravski okrug"
	case SubdivisionRSKM:
		return "Kosovo-Metohija"
	case SubdivisionRSVO:
		return "Vojvodina"
	case SubdivisionRUAD:
		return "Adygeya, Respublika"
	case SubdivisionRUAL:
		return "Altay, Respublika"
	case SubdivisionRUALT:
		return "Altayskiy kray"
	case SubdivisionRUAMU:
		return "Amurskaya oblast'"
	case SubdivisionRUARK:
		return "Arkhangel'skaya oblast'"
	case SubdivisionRUAST:
		return "Astrakhanskaya oblast'"
	case SubdivisionRUBA:
		return "Bashkortostan, Respublika"
	case SubdivisionRUBEL:
		return "Belgorodskaya oblast'"
	case SubdivisionRUBRY:
		return "Bryanskaya oblast'"
	case SubdivisionRUBU:
		return "Buryatiya, Respublika"
	case SubdivisionRUCE:
		return "Chechenskaya Respublika"
	case SubdivisionRUCHE:
		return "Chelyabinskaya oblast'"
	case SubdivisionRUCHU:
		return "Chukotskiy avtonomnyy okrug"
	case SubdivisionRUCU:
		return "Chuvashskaya Respublika"
	case SubdivisionRUDA:
		return "Dagestan, Respublika"
	case SubdivisionRUIN:
		return "Respublika Ingushetiya"
	case SubdivisionRUIRK:
		return "Irkutiskaya oblast'"
	case SubdivisionRUIVA:
		return "Ivanovskaya oblast'"
	case SubdivisionRUKAM:
		return "Kamchatskiy kray"
	case SubdivisionRUKB:
		return "Kabardino-Balkarskaya Respublika"
	case SubdivisionRUKC:
		return "Karachayevo-Cherkesskaya Respublika"
	case SubdivisionRUKDA:
		return "Krasnodarskiy kray"
	case SubdivisionRUKEM:
		return "Kemerovskaya oblast'"
	case SubdivisionRUKGD:
		return "Kaliningradskaya oblast'"
	case SubdivisionRUKGN:
		return "Kurganskaya oblast'"
	case SubdivisionRUKHA:
		return "Khabarovskiy kray"
	case SubdivisionRUKHM:
		return "Khanty-Mansiysky avtonomnyy okrug-Yugra"
	case SubdivisionRUKIR:
		return "Kirovskaya oblast'"
	case SubdivisionRUKK:
		return "Khakasiya, Respublika"
	case SubdivisionRUKL:
		return "Kalmykiya, Respublika"
	case SubdivisionRUKLU:
		return "Kaluzhskaya oblast'"
	case SubdivisionRUKO:
		return "Komi, Respublika"
	case SubdivisionRUKOS:
		return "Kostromskaya oblast'"
	case SubdivisionRUKR:
		return "Kareliya, Respublika"
	case SubdivisionRUKRS:
		return "Kurskaya oblast'"
	case SubdivisionRUKYA:
		return "Krasnoyarskiy kray"
	case SubdivisionRULEN:
		return "Leningradskaya oblast'"
	case SubdivisionRULIP:
		return "Lipetskaya oblast'"
	case SubdivisionRUMAG:
		return "Magadanskaya oblast'"
	case SubdivisionRUME:
		return "Mariy El, Respublika"
	case SubdivisionRUMO:
		return "Mordoviya, Respublika"
	case SubdivisionRUMOS:
		return "Moskovskaya oblast'"
	case SubdivisionRUMOW:
		return "Moskva"
	case SubdivisionRUMUR:
		return "Murmanskaya oblast'"
	case SubdivisionRUNEN:
		return "Nenetskiy avtonomnyy okrug"
	case SubdivisionRUNGR:
		return "Novgorodskaya oblast'"
	case SubdivisionRUNIZ:
		return "Nizhegorodskaya oblast'"
	case SubdivisionRUNVS:
		return "Novosibirskaya oblast'"
	case SubdivisionRUOMS:
		return "Omskaya oblast'"
	case SubdivisionRUORE:
		return "Orenburgskaya oblast'"
	case SubdivisionRUORL:
		return "Orlovskaya oblast'"
	case SubdivisionRUPER:
		return "Permskiy kray"
	case SubdivisionRUPNZ:
		return "Penzenskaya oblast'"
	case SubdivisionRUPRI:
		return "Primorskiy kray"
	case SubdivisionRUPSK:
		return "Pskovskaya oblast'"
	case SubdivisionRUROS:
		return "Rostovskaya oblast'"
	case SubdivisionRURYA:
		return "Ryazanskaya oblast'"
	case SubdivisionRUSA:
		return "Sakha, Respublika [Yakutiya]"
	case SubdivisionRUSAK:
		return "Sakhalinskaya oblast'"
	case SubdivisionRUSAM:
		return "Samaraskaya oblast'"
	case SubdivisionRUSAR:
		return "Saratovskaya oblast'"
	case SubdivisionRUSE:
		return "Severnaya Osetiya-Alaniya, Respublika"
	case SubdivisionRUSMO:
		return "Smolenskaya oblast'"
	case SubdivisionRUSPE:
		return "Sankt-Peterburg"
	case SubdivisionRUSTA:
		return "Stavropol'skiy kray"
	case SubdivisionRUSVE:
		return "Sverdlovskaya oblast'"
	case SubdivisionRUTA:
		return "Tatarstan, Respublika"
	case SubdivisionRUTAM:
		return "Tambovskaya oblast'"
	case SubdivisionRUTOM:
		return "Tomskaya oblast'"
	case SubdivisionRUTUL:
		return "Tul'skaya oblast'"
	case SubdivisionRUTVE:
		return "Tverskaya oblast'"
	case SubdivisionRUTY:
		return "Tyva, Respublika [Tuva]"
	case SubdivisionRUTYU:
		return "Tyumenskaya oblast'"
	case SubdivisionRUUD:
		return "Udmurtskaya Respublika"
	case SubdivisionRUULY:
		return "Ul'yanovskaya oblast'"
	case SubdivisionRUVGG:
		return "Volgogradskaya oblast'"
	case SubdivisionRUVLA:
		return "Vladimirskaya oblast'"
	case SubdivisionRUVLG:
		return "Vologodskaya oblast'"
	case SubdivisionRUVOR:
		return "Voronezhskaya oblast'"
	case SubdivisionRUYAN:
		return "Yamalo-Nenetskiy avtonomnyy okrug"
	case SubdivisionRUYAR:
		return "Yaroslavskaya oblast'"
	case SubdivisionRUYEV:
		return "Yevreyskaya avtonomnaya oblast'"
	case SubdivisionRUZAB:
		return "Zabajkal'skij kraj"
	case SubdivisionRW01:
		return "Ville de Kigali"
	case SubdivisionRW02:
		return "Est"
	case SubdivisionRW03:
		return "Nord"
	case SubdivisionRW04:
		return "Ouest"
	case SubdivisionRW05:
		return "Sud"
	case SubdivisionSA01:
		return "Ar Riyāḍ"
	case SubdivisionSA02:
		return "Makkah"
	case SubdivisionSA03:
		return "Al Madīnah"
	case SubdivisionSA04:
		return "Ash Sharqīyah"
	case SubdivisionSA05:
		return "Al Qaşīm"
	case SubdivisionSA06:
		return "Ḥā'il"
	case SubdivisionSA07:
		return "Tabūk"
	case SubdivisionSA08:
		return "Al Ḥudūd ash Shamāliyah"
	case SubdivisionSA09:
		return "Jīzan"
	case SubdivisionSA10:
		return "Najrān"
	case SubdivisionSA11:
		return "Al Bāhah"
	case SubdivisionSA12:
		return "Al Jawf"
	case SubdivisionSA14:
		return "`Asīr"
	case SubdivisionSBCE:
		return "Central"
	case SubdivisionSBCH:
		return "Choiseul"
	case SubdivisionSBCT:
		return "Capital Territory (Honiara)"
	case SubdivisionSBGU:
		return "Guadalcanal"
	case SubdivisionSBIS:
		return "Isabel"
	case SubdivisionSBMK:
		return "Makira"
	case SubdivisionSBML:
		return "Malaita"
	case SubdivisionSBRB:
		return "Rennell and Bellona"
	case SubdivisionSBTE:
		return "Temotu"
	case SubdivisionSBWE:
		return "Western"
	case SubdivisionSC01:
		return "Anse aux Pins"
	case SubdivisionSC02:
		return "Anse Boileau"
	case SubdivisionSC03:
		return "Anse Etoile"
	case SubdivisionSC04:
		return "Anse Louis"
	case SubdivisionSC05:
		return "Anse Royale"
	case SubdivisionSC06:
		return "Baie Lazare"
	case SubdivisionSC07:
		return "Baie Sainte Anne"
	case SubdivisionSC08:
		return "Beau Vallon"
	case SubdivisionSC09:
		return "Bel Air"
	case SubdivisionSC10:
		return "Bel Ombre"
	case SubdivisionSC11:
		return "Cascade"
	case SubdivisionSC12:
		return "Glacis"
	case SubdivisionSC13:
		return "Grand Anse Mahe"
	case SubdivisionSC14:
		return "Grand Anse Praslin"
	case SubdivisionSC15:
		return "La Digue"
	case SubdivisionSC16:
		return "English River"
	case SubdivisionSC17:
		return "Mont Buxton"
	case SubdivisionSC18:
		return "Mont Fleuri"
	case SubdivisionSC19:
		return "Plaisance"
	case SubdivisionSC20:
		return "Pointe Larue"
	case SubdivisionSC21:
		return "Port Glaud"
	case SubdivisionSC22:
		return "Saint Louis"
	case SubdivisionSC23:
		return "Takamaka"
	case SubdivisionSC24:
		return "Les Mamelles"
	case SubdivisionSC25:
		return "Roche Caiman"
	case SubdivisionSDDC:
		return "Zalingei"
	case SubdivisionSDDE:
		return "Sharq Dārfūr"
	case SubdivisionSDDN:
		return "Shamāl Dārfūr"
	case SubdivisionSDDS:
		return "Janūb Dārfūr"
	case SubdivisionSDDW:
		return "Gharb Dārfūr"
	case SubdivisionSDGD:
		return "Al Qaḑārif"
	case SubdivisionSDGZ:
		return "Al Jazīrah"
	case SubdivisionSDKA:
		return "Kassalā"
	case SubdivisionSDKH:
		return "Al Kharţūm"
	case SubdivisionSDKN:
		return "Shamāl Kurdufān"
	case SubdivisionSDKS:
		return "Janūb Kurdufān"
	case SubdivisionSDNB:
		return "An Nīl al Azraq"
	case SubdivisionSDNO:
		return "Ash Shamālīyah"
	case SubdivisionSDNR:
		return "An Nīl"
	case SubdivisionSDNW:
		return "An Nīl al Abyaḑ"
	case SubdivisionSDRS:
		return "Al Baḩr al Aḩmar"
	case SubdivisionSDSI:
		return "Sinnār"
	case SubdivisionSEAB:
		return "Stockholms län"
	case SubdivisionSEAC:
		return "Västerbottens län"
	case SubdivisionSEBD:
		return "Norrbottens län"
	case SubdivisionSEC:
		return "Uppsala län"
	case SubdivisionSED:
		return "Södermanlands län"
	case SubdivisionSEE:
		return "Östergötlands län"
	case SubdivisionSEF:
		return "Jönköpings län"
	case SubdivisionSEG:
		return "Kronobergs län"
	case SubdivisionSEH:
		return "Kalmar län"
	case SubdivisionSEI:
		return "Gotlands län"
	case SubdivisionSEK:
		return "Blekinge län"
	case SubdivisionSEM:
		return "Skåne län"
	case SubdivisionSEN:
		return "Hallands län"
	case SubdivisionSEO:
		return "Västra Götalands län"
	case SubdivisionSES:
		return "Värmlands län"
	case SubdivisionSET:
		return "Örebro län"
	case SubdivisionSEU:
		return "Västmanlands län"
	case SubdivisionSEW:
		return "Dalarnas län"
	case SubdivisionSEX:
		return "Gävleborgs län"
	case SubdivisionSEY:
		return "Västernorrlands län"
	case SubdivisionSEZ:
		return "Jämtlands län"
	case SubdivisionSG01:
		return "Central Singapore"
	case SubdivisionSG02:
		return "North East"
	case SubdivisionSG03:
		return "North West"
	case SubdivisionSG04:
		return "South East"
	case SubdivisionSG05:
		return "South West"
	case SubdivisionSHAC:
		return "Ascension"
	case SubdivisionSHHL:
		return "Saint Helena"
	case SubdivisionSHTA:
		return "Tristan da Cunha"
	case SubdivisionSI001:
		return "Ajdovščina"
	case SubdivisionSI002:
		return "Beltinci"
	case SubdivisionSI003:
		return "Bled"
	case SubdivisionSI004:
		return "Bohinj"
	case SubdivisionSI005:
		return "Borovnica"
	case SubdivisionSI006:
		return "Bovec"
	case SubdivisionSI007:
		return "Brda"
	case SubdivisionSI008:
		return "Brezovica"
	case SubdivisionSI009:
		return "Brežice"
	case SubdivisionSI010:
		return "Tišina"
	case SubdivisionSI011:
		return "Celje"
	case SubdivisionSI012:
		return "Cerklje na Gorenjskem"
	case SubdivisionSI013:
		return "Cerknica"
	case SubdivisionSI014:
		return "Cerkno"
	case SubdivisionSI015:
		return "Črenšovci"
	case SubdivisionSI016:
		return "Črna na Koroškem"
	case SubdivisionSI017:
		return "Črnomelj"
	case SubdivisionSI018:
		return "Destrnik"
	case SubdivisionSI019:
		return "Divača"
	case SubdivisionSI020:
		return "Dobrepolje"
	case SubdivisionSI021:
		return "Dobrova-Polhov Gradec"
	case SubdivisionSI022:
		return "Dol pri Ljubljani"
	case SubdivisionSI023:
		return "Domžale"
	case SubdivisionSI024:
		return "Dornava"
	case SubdivisionSI025:
		return "Dravograd"
	case SubdivisionSI026:
		return "Duplek"
	case SubdivisionSI027:
		return "Gorenja vas-Poljane"
	case SubdivisionSI028:
		return "Gorišnica"
	case SubdivisionSI029:
		return "Gornja Radgona"
	case SubdivisionSI030:
		return "Gornji Grad"
	case SubdivisionSI031:
		return "Gornji Petrovci"
	case SubdivisionSI032:
		return "Grosuplje"
	case SubdivisionSI033:
		return "Šalovci"
	case SubdivisionSI034:
		return "Hrastnik"
	case SubdivisionSI035:
		return "Hrpelje-Kozina"
	case SubdivisionSI036:
		return "Idrija"
	case SubdivisionSI037:
		return "Ig"
	case SubdivisionSI038:
		return "Ilirska Bistrica"
	case SubdivisionSI039:
		return "Ivančna Gorica"
	case SubdivisionSI040:
		return "Izola/Isola"
	case SubdivisionSI041:
		return "Jesenice"
	case SubdivisionSI042:
		return "Juršinci"
	case SubdivisionSI043:
		return "Kamnik"
	case SubdivisionSI044:
		return "Kanal"
	case SubdivisionSI045:
		return "Kidričevo"
	case SubdivisionSI046:
		return "Kobarid"
	case SubdivisionSI047:
		return "Kobilje"
	case SubdivisionSI048:
		return "Kočevje"
	case SubdivisionSI049:
		return "Komen"
	case SubdivisionSI050:
		return "Koper/Capodistria"
	case SubdivisionSI051:
		return "Kozje"
	case SubdivisionSI052:
		return "Kranj"
	case SubdivisionSI053:
		return "Kranjska Gora"
	case SubdivisionSI054:
		return "Krško"
	case SubdivisionSI055:
		return "Kungota"
	case SubdivisionSI056:
		return "Kuzma"
	case SubdivisionSI057:
		return "Laško"
	case SubdivisionSI058:
		return "Lenart"
	case SubdivisionSI059:
		return "Lendava/Lendva"
	case SubdivisionSI060:
		return "Litija"
	case SubdivisionSI061:
		return "Ljubljana"
	case SubdivisionSI062:
		return "Ljubno"
	case SubdivisionSI063:
		return "Ljutomer"
	case SubdivisionSI064:
		return "Logatec"
	case SubdivisionSI065:
		return "Loška dolina"
	case SubdivisionSI066:
		return "Loški Potok"
	case SubdivisionSI067:
		return "Luče"
	case SubdivisionSI068:
		return "Lukovica"
	case SubdivisionSI069:
		return "Majšperk"
	case SubdivisionSI070:
		return "Maribor"
	case SubdivisionSI071:
		return "Medvode"
	case SubdivisionSI072:
		return "Mengeš"
	case SubdivisionSI073:
		return "Metlika"
	case SubdivisionSI074:
		return "Mežica"
	case SubdivisionSI075:
		return "Miren-Kostanjevica"
	case SubdivisionSI076:
		return "Mislinja"
	case SubdivisionSI077:
		return "Moravče"
	case SubdivisionSI078:
		return "Moravske Toplice"
	case SubdivisionSI079:
		return "Mozirje"
	case SubdivisionSI080:
		return "Murska Sobota"
	case SubdivisionSI081:
		return "Muta"
	case SubdivisionSI082:
		return "Naklo"
	case SubdivisionSI083:
		return "Nazarje"
	case SubdivisionSI084:
		return "Nova Gorica"
	case SubdivisionSI085:
		return "Novo mesto"
	case SubdivisionSI086:
		return "Odranci"
	case SubdivisionSI087:
		return "Ormož"
	case SubdivisionSI088:
		return "Osilnica"
	case SubdivisionSI089:
		return "Pesnica"
	case SubdivisionSI090:
		return "Piran/Pirano"
	case SubdivisionSI091:
		return "Pivka"
	case SubdivisionSI092:
		return "Podčetrtek"
	case SubdivisionSI093:
		return "Podvelka"
	case SubdivisionSI094:
		return "Postojna"
	case SubdivisionSI095:
		return "Preddvor"
	case SubdivisionSI096:
		return "Ptuj"
	case SubdivisionSI097:
		return "Puconci"
	case SubdivisionSI098:
		return "Rače-Fram"
	case SubdivisionSI099:
		return "Radeče"
	case SubdivisionSI100:
		return "Radenci"
	case SubdivisionSI101:
		return "Radlje ob Dravi"
	case SubdivisionSI102:
		return "Radovljica"
	case SubdivisionSI103:
		return "Ravne na Koroškem"
	case SubdivisionSI104:
		return "Ribnica"
	case SubdivisionSI105:
		return "Rogašovci"
	case SubdivisionSI106:
		return "Rogaška Slatina"
	case SubdivisionSI107:
		return "Rogatec"
	case SubdivisionSI108:
		return "Ruše"
	case SubdivisionSI109:
		return "Semič"
	case SubdivisionSI110:
		return "Sevnica"
	case SubdivisionSI111:
		return "Sežana"
	case SubdivisionSI112:
		return "Slovenj Gradec"
	case SubdivisionSI113:
		return "Slovenska Bistrica"
	case SubdivisionSI114:
		return "Slovenske Konjice"
	case SubdivisionSI115:
		return "Starče"
	case SubdivisionSI116:
		return "Sveti Jurij"
	case SubdivisionSI117:
		return "Šenčur"
	case SubdivisionSI118:
		return "Šentilj"
	case SubdivisionSI119:
		return "Šentjernej"
	case SubdivisionSI120:
		return "Šentjur"
	case SubdivisionSI121:
		return "Škocjan"
	case SubdivisionSI122:
		return "Škofja Loka"
	case SubdivisionSI123:
		return "Škofljica"
	case SubdivisionSI124:
		return "Šmarje pri Jelšah"
	case SubdivisionSI125:
		return "Šmartno ob Paki"
	case SubdivisionSI126:
		return "Šoštanj"
	case SubdivisionSI127:
		return "Štore"
	case SubdivisionSI128:
		return "Tolmin"
	case SubdivisionSI129:
		return "Trbovlje"
	case SubdivisionSI130:
		return "Trebnje"
	case SubdivisionSI131:
		return "Tržič"
	case SubdivisionSI132:
		return "Turnišče"
	case SubdivisionSI133:
		return "Velenje"
	case SubdivisionSI134:
		return "Velike Lašče"
	case SubdivisionSI135:
		return "Videm"
	case SubdivisionSI136:
		return "Vipava"
	case SubdivisionSI137:
		return "Vitanje"
	case SubdivisionSI138:
		return "Vodice"
	case SubdivisionSI139:
		return "Vojnik"
	case SubdivisionSI140:
		return "Vrhnika"
	case SubdivisionSI141:
		return "Vuzenica"
	case SubdivisionSI142:
		return "Zagorje ob Savi"
	case SubdivisionSI143:
		return "Zavrč"
	case SubdivisionSI144:
		return "Zreče"
	case SubdivisionSI146:
		return "Železniki"
	case SubdivisionSI147:
		return "Žiri"
	case SubdivisionSI148:
		return "Benedikt"
	case SubdivisionSI149:
		return "Bistrica ob Sotli"
	case SubdivisionSI150:
		return "Bloke"
	case SubdivisionSI151:
		return "Braslovče"
	case SubdivisionSI152:
		return "Cankova"
	case SubdivisionSI153:
		return "Cerkvenjak"
	case SubdivisionSI154:
		return "Dobje"
	case SubdivisionSI155:
		return "Dobrna"
	case SubdivisionSI156:
		return "Dobrovnik/Dobronak"
	case SubdivisionSI157:
		return "Dolenjske Toplice"
	case SubdivisionSI158:
		return "Grad"
	case SubdivisionSI159:
		return "Hajdina"
	case SubdivisionSI160:
		return "Hoče-Slivnica"
	case SubdivisionSI161:
		return "Hodoš/Hodos"
	case SubdivisionSI162:
		return "Horjul"
	case SubdivisionSI163:
		return "Jezersko"
	case SubdivisionSI164:
		return "Komenda"
	case SubdivisionSI165:
		return "Kostel"
	case SubdivisionSI166:
		return "Križevci"
	case SubdivisionSI167:
		return "Lovrenc na Pohorju"
	case SubdivisionSI168:
		return "Markovci"
	case SubdivisionSI169:
		return "Miklavž na Dravskem polju"
	case SubdivisionSI170:
		return "Mirna Peč"
	case SubdivisionSI171:
		return "Oplotnica"
	case SubdivisionSI172:
		return "Podlehnik"
	case SubdivisionSI173:
		return "Polzela"
	case SubdivisionSI174:
		return "Prebold"
	case SubdivisionSI175:
		return "Prevalje"
	case SubdivisionSI176:
		return "Razkrižje"
	case SubdivisionSI177:
		return "Ribnica na Pohorju"
	case SubdivisionSI178:
		return "Selnica ob Dravi"
	case SubdivisionSI179:
		return "Sodražica"
	case SubdivisionSI180:
		return "Solčava"
	case SubdivisionSI181:
		return "Sveta Ana"
	case SubdivisionSI182:
		return "Sveta Andraž v Slovenskih Goricah"
	case SubdivisionSI183:
		return "Šempeter-Vrtojba"
	case SubdivisionSI184:
		return "Tabor"
	case SubdivisionSI185:
		return "Trnovska vas"
	case SubdivisionSI186:
		return "Trzin"
	case SubdivisionSI187:
		return "Velika Polana"
	case SubdivisionSI188:
		return "Veržej"
	case SubdivisionSI189:
		return "Vransko"
	case SubdivisionSI190:
		return "Žalec"
	case SubdivisionSI191:
		return "Žetale"
	case SubdivisionSI192:
		return "Žirovnica"
	case SubdivisionSI193:
		return "Žužemberk"
	case SubdivisionSI194:
		return "Šmartno pri Litiji"
	case SubdivisionSI195:
		return "Apače"
	case SubdivisionSI196:
		return "Cirkulane"
	case SubdivisionSI197:
		return "Kosanjevica na Krki"
	case SubdivisionSI198:
		return "Makole"
	case SubdivisionSI199:
		return "Mokronog-Trebelno"
	case SubdivisionSI200:
		return "Poljčane"
	case SubdivisionSI201:
		return "Renče-Vogrsko"
	case SubdivisionSI202:
		return "Središče ob Dravi"
	case SubdivisionSI203:
		return "Straža"
	case SubdivisionSI204:
		return "Sveta Trojica v Slovenskih Goricah"
	case SubdivisionSI205:
		return "Sveti Tomaž"
	case SubdivisionSI206:
		return "Šmarjeske Topliče"
	case SubdivisionSI207:
		return "Gorje"
	case SubdivisionSI208:
		return "Log-Dragomer"
	case SubdivisionSI209:
		return "Rečica ob Savinji"
	case SubdivisionSI210:
		return "Sveti Jurij v Slovenskih Goricah"
	case SubdivisionSI211:
		return "Šentrupert"
	case SubdivisionSKBC:
		return "Banskobystrický kraj"
	case SubdivisionSKBL:
		return "Bratislavský kraj"
	case SubdivisionSKKI:
		return "Košický kraj"
	case SubdivisionSKNI:
		return "Nitriansky kraj"
	case SubdivisionSKPV:
		return "Prešovský kraj"
	case SubdivisionSKTA:
		return "Trnavský kraj"
	case SubdivisionSKTC:
		return "Trenčiansky kraj"
	case SubdivisionSKZI:
		return "Žilinský kraj"
	case SubdivisionSLE:
		return "Eastern"
	case SubdivisionSLN:
		return "Northern"
	case SubdivisionSLS:
		return "Southern (Sierra Leone)"
	case SubdivisionSLW:
		return "Western Area (Freetown)"
	case SubdivisionSM01:
		return "Acquaviva"
	case SubdivisionSM02:
		return "Chiesanuova"
	case SubdivisionSM03:
		return "Domagnano"
	case SubdivisionSM04:
		return "Faetano"
	case SubdivisionSM05:
		return "Fiorentino"
	case SubdivisionSM06:
		return "Borgo Maggiore"
	case SubdivisionSM07:
		return "San Marino"
	case SubdivisionSM08:
		return "Montegiardino"
	case SubdivisionSM09:
		return "Serravalle"
	case SubdivisionSNDB:
		return "Diourbel"
	case SubdivisionSNDK:
		return "Dakar"
	case SubdivisionSNFK:
		return "Fatick"
	case SubdivisionSNKA:
		return "Kaffrine"
	case SubdivisionSNKD:
		return "Kolda"
	case SubdivisionSNKE:
		return "Kédougou"
	case SubdivisionSNKL:
		return "Kaolack"
	case SubdivisionSNLG:
		return "Louga"
	case SubdivisionSNMT:
		return "Matam"
	case SubdivisionSNSE:
		return "Sédhiou"
	case SubdivisionSNSL:
		return "Saint-Louis"
	case SubdivisionSNTC:
		return "Tambacounda"
	case SubdivisionSNTH:
		return "Thiès"
	case SubdivisionSNZG:
		return "Ziguinchor"
	case SubdivisionSOAW:
		return "Awdal"
	case SubdivisionSOBK:
		return "Bakool"
	case SubdivisionSOBN:
		return "Banaadir"
	case SubdivisionSOBR:
		return "Bari"
	case SubdivisionSOBY:
		return "Bay"
	case SubdivisionSOGA:
		return "Galguduud"
	case SubdivisionSOGE:
		return "Gedo"
	case SubdivisionSOHI:
		return "Hiirsan"
	case SubdivisionSOJD:
		return "Jubbada Dhexe"
	case SubdivisionSOJH:
		return "Jubbada Hoose"
	case SubdivisionSOMU:
		return "Mudug"
	case SubdivisionSONU:
		return "Nugaal"
	case SubdivisionSOSA:
		return "Saneag"
	case SubdivisionSOSD:
		return "Shabeellaha Dhexe"
	case SubdivisionSOSH:
		return "Shabeellaha Hoose"
	case SubdivisionSOSO:
		return "Sool"
	case SubdivisionSOTO:
		return "Togdheer"
	case SubdivisionSOWO:
		return "Woqooyi Galbeed"
	case SubdivisionSRBR:
		return "Brokopondo"
	case SubdivisionSRCM:
		return "Commewijne"
	case SubdivisionSRCR:
		return "Coronie"
	case SubdivisionSRMA:
		return "Marowijne"
	case SubdivisionSRNI:
		return "Nickerie"
	case SubdivisionSRPM:
		return "Paramaribo"
	case SubdivisionSRPR:
		return "Para"
	case SubdivisionSRSA:
		return "Saramacca"
	case SubdivisionSRSI:
		return "Sipaliwini"
	case SubdivisionSRWA:
		return "Wanica"
	case SubdivisionSSBN:
		return "Northern Bahr el Ghazal"
	case SubdivisionSSBW:
		return "Western Bahr el Ghazal"
	case SubdivisionSSEC:
		return "Central Equatoria"
	case SubdivisionSSEE:
		return "Eastern Equatoria"
	case SubdivisionSSEW:
		return "Western Equatoria"
	case SubdivisionSSJG:
		return "Jonglei"
	case SubdivisionSSLK:
		return "Lakes"
	case SubdivisionSSNU:
		return "Upper Nile"
	case SubdivisionSSUY:
		return "Unity"
	case SubdivisionSSWR:
		return "Warrap"
	case SubdivisionSTP:
		return "Príncipe"
	case SubdivisionSTS:
		return "São Tomé"
	case SubdivisionSVAH:
		return "Ahuachapán"
	case SubdivisionSVCA:
		return "Cabañas"
	case SubdivisionSVCH:
		return "Chalatenango"
	case SubdivisionSVCU:
		return "Cuscatlán"
	case SubdivisionSVLI:
		return "La Libertad"
	case SubdivisionSVMO:
		return "Morazán"
	case SubdivisionSVPA:
		return "La Paz"
	case SubdivisionSVSA:
		return "Santa Ana"
	case SubdivisionSVSM:
		return "San Miguel"
	case SubdivisionSVSO:
		return "Sonsonate"
	case SubdivisionSVSS:
		return "San Salvador"
	case SubdivisionSVSV:
		return "San Vicente"
	case SubdivisionSVUN:
		return "La Unión"
	case SubdivisionSVUS:
		return "Usulután"
	case SubdivisionSYDI:
		return "Dimashq"
	case SubdivisionSYDR:
		return "Dar'a"
	case SubdivisionSYDY:
		return "Dayr az Zawr"
	case SubdivisionSYHA:
		return "Al Hasakah"
	case SubdivisionSYHI:
		return "Homs"
	case SubdivisionSYHL:
		return "Halab"
	case SubdivisionSYHM:
		return "Hamah"
	case SubdivisionSYID:
		return "Idlib"
	case SubdivisionSYLA:
		return "Al Ladhiqiyah"
	case SubdivisionSYQU:
		return "Al Qunaytirah"
	case SubdivisionSYRA:
		return "Ar Raqqah"
	case SubdivisionSYRD:
		return "Rif Dimashq"
	case SubdivisionSYSU:
		return "As Suwayda'"
	case SubdivisionSYTA:
		return "Tartus"
	case SubdivisionSZHH:
		return "Hhohho"
	case SubdivisionSZLU:
		return "Lubombo"
	case SubdivisionSZMA:
		return "Manzini"
	case SubdivisionSZSH:
		return "Shiselweni"
	case SubdivisionTDBA:
		return "Al Baṭḩah"
	case SubdivisionTDBG:
		return "Baḩr al Ghazāl"
	case SubdivisionTDBO:
		return "Būrkū"
	case SubdivisionTDCB:
		return "Shārī Bāqirmī"
	case SubdivisionTDEN:
		return "Innīdī"
	case SubdivisionTDGR:
		return "Qīrā"
	case SubdivisionTDHL:
		return "Ḥajjar Lamīs"
	case SubdivisionTDKA:
		return "Kānim"
	case SubdivisionTDLC:
		return "Al Buḩayrah"
	case SubdivisionTDLO:
		return "Lūqūn al Gharbī"
	case SubdivisionTDLR:
		return "Lūqūn ash Sharqī"
	case SubdivisionTDMA:
		return "Māndūl"
	case SubdivisionTDMC:
		return "Shārī al Awsaṭ"
	case SubdivisionTDME:
		return "Māyū Kībbī ash Sharqī"
	case SubdivisionTDMO:
		return "Māyū Kībbī al Gharbī"
	case SubdivisionTDND:
		return "Madīnat Injamīnā"
	case SubdivisionTDOD:
		return "Waddāy"
	case SubdivisionTDSA:
		return "Salāmāt"
	case SubdivisionTDSI:
		return "Sīlā"
	case SubdivisionTDTA:
		return "Tānjilī"
	case SubdivisionTDTI:
		return "Tibastī"
	case SubdivisionTDWF:
		return "Wādī Fīrā"
	case SubdivisionTGC:
		return "Région du Centre"
	case SubdivisionTGK:
		return "Région de la Kara"
	case SubdivisionTGM:
		return "Région Maritime"
	case SubdivisionTGP:
		return "Région des Plateaux"
	case SubdivisionTGS:
		return "Région des Savannes"
	case SubdivisionTH10:
		return "Krung Thep Maha Nakhon Bangkok"
	case SubdivisionTH11:
		return "Samut Prakan"
	case SubdivisionTH12:
		return "Nonthaburi"
	case SubdivisionTH13:
		return "Pathum Thani"
	case SubdivisionTH14:
		return "Phra Nakhon Si Ayutthaya"
	case SubdivisionTH15:
		return "Ang Thong"
	case SubdivisionTH16:
		return "Lop Buri"
	case SubdivisionTH17:
		return "Sing Buri"
	case SubdivisionTH18:
		return "Chai Nat"
	case SubdivisionTH19:
		return "Saraburi"
	case SubdivisionTH20:
		return "Chon Buri"
	case SubdivisionTH21:
		return "Rayong"
	case SubdivisionTH22:
		return "Chanthaburi"
	case SubdivisionTH23:
		return "Trat"
	case SubdivisionTH24:
		return "Chachoengsao"
	case SubdivisionTH25:
		return "Prachin Buri"
	case SubdivisionTH26:
		return "Nakhon Nayok"
	case SubdivisionTH27:
		return "Sa Kaeo"
	case SubdivisionTH30:
		return "Nakhon Ratchasima"
	case SubdivisionTH31:
		return "Buri Ram"
	case SubdivisionTH32:
		return "Surin"
	case SubdivisionTH33:
		return "Si Sa Ket"
	case SubdivisionTH34:
		return "Ubon Ratchathani"
	case SubdivisionTH35:
		return "Yasothon"
	case SubdivisionTH36:
		return "Chaiyaphum"
	case SubdivisionTH37:
		return "Amnat Charoen"
	case SubdivisionTH39:
		return "Nong Bua Lam Phu"
	case SubdivisionTH40:
		return "Khon Kaen"
	case SubdivisionTH41:
		return "Udon Thani"
	case SubdivisionTH42:
		return "Loei"
	case SubdivisionTH43:
		return "Nong Khai"
	case SubdivisionTH44:
		return "Maha Sarakham"
	case SubdivisionTH45:
		return "Roi Et"
	case SubdivisionTH46:
		return "Kalasin"
	case SubdivisionTH47:
		return "Sakon Nakhon"
	case SubdivisionTH48:
		return "Nakhon Phanom"
	case SubdivisionTH49:
		return "Mukdahan"
	case SubdivisionTH50:
		return "Chiang Mai"
	case SubdivisionTH51:
		return "Lamphun"
	case SubdivisionTH52:
		return "Lampang"
	case SubdivisionTH53:
		return "Uttaradit"
	case SubdivisionTH54:
		return "Phrae"
	case SubdivisionTH55:
		return "Nan"
	case SubdivisionTH56:
		return "Phayao"
	case SubdivisionTH57:
		return "Chiang Rai"
	case SubdivisionTH58:
		return "Mae Hong Son"
	case SubdivisionTH60:
		return "Nakhon Sawan"
	case SubdivisionTH61:
		return "Uthai Thani"
	case SubdivisionTH62:
		return "Kamphaeng Phet"
	case SubdivisionTH63:
		return "Tak"
	case SubdivisionTH64:
		return "Sukhothai"
	case SubdivisionTH65:
		return "Phitsanulok"
	case SubdivisionTH66:
		return "Phichit"
	case SubdivisionTH67:
		return "Phetchabun"
	case SubdivisionTH70:
		return "Ratchaburi"
	case SubdivisionTH71:
		return "Kanchanaburi"
	case SubdivisionTH72:
		return "Suphan Buri"
	case SubdivisionTH73:
		return "Nakhon Pathom"
	case SubdivisionTH74:
		return "Samut Sakhon"
	case SubdivisionTH75:
		return "Samut Songkhram"
	case SubdivisionTH76:
		return "Phetchaburi"
	case SubdivisionTH77:
		return "Prachuap Khiri Khan"
	case SubdivisionTH80:
		return "Nakhon Si Thammarat"
	case SubdivisionTH81:
		return "Krabi"
	case SubdivisionTH82:
		return "Phangnga"
	case SubdivisionTH83:
		return "Phuket"
	case SubdivisionTH84:
		return "Surat Thani"
	case SubdivisionTH85:
		return "Ranong"
	case SubdivisionTH86:
		return "Chumphon"
	case SubdivisionTH90:
		return "Songkhla"
	case SubdivisionTH91:
		return "Satun"
	case SubdivisionTH92:
		return "Trang"
	case SubdivisionTH93:
		return "Phatthalung"
	case SubdivisionTH94:
		return "Pattani"
	case SubdivisionTH95:
		return "Yala"
	case SubdivisionTH96:
		return "Narathiwat"
	case SubdivisionTHS:
		return "Phatthaya"
	case SubdivisionTJGB:
		return "Gorno-Badakhshan"
	case SubdivisionTJKT:
		return "Khatlon"
	case SubdivisionTJSU:
		return "Sughd"
	case SubdivisionTLAL:
		return "Aileu"
	case SubdivisionTLAN:
		return "Ainaro"
	case SubdivisionTLBA:
		return "Baucau"
	case SubdivisionTLBO:
		return "Bobonaro"
	case SubdivisionTLCO:
		return "Cova Lima"
	case SubdivisionTLDI:
		return "Díli"
	case SubdivisionTLER:
		return "Ermera"
	case SubdivisionTLLA:
		return "Lautem"
	case SubdivisionTLLI:
		return "Liquiça"
	case SubdivisionTLMF:
		return "Manufahi"
	case SubdivisionTLMT:
		return "Manatuto"
	case SubdivisionTLOE:
		return "Oecussi"
	case SubdivisionTLVI:
		return "Viqueque"
	case SubdivisionTMA:
		return "Ahal"
	case SubdivisionTMB:
		return "Balkan"
	case SubdivisionTMD:
		return "Daşoguz"
	case SubdivisionTML:
		return "Lebap"
	case SubdivisionTMM:
		return "Mary"
	case SubdivisionTMS:
		return "Aşgabat"
	case SubdivisionTN11:
		return "Tunis"
	case SubdivisionTN12:
		return "Ariana"
	case SubdivisionTN13:
		return "Ben Arous"
	case SubdivisionTN14:
		return "La Manouba"
	case SubdivisionTN21:
		return "Nabeul"
	case SubdivisionTN22:
		return "Zaghouan"
	case SubdivisionTN23:
		return "Bizerte"
	case SubdivisionTN31:
		return "Béja"
	case SubdivisionTN32:
		return "Jendouba"
	case SubdivisionTN33:
		return "Le Kef"
	case SubdivisionTN34:
		return "Siliana"
	case SubdivisionTN41:
		return "Kairouan"
	case SubdivisionTN42:
		return "Kasserine"
	case SubdivisionTN43:
		return "Sidi Bouzid"
	case SubdivisionTN51:
		return "Sousse"
	case SubdivisionTN52:
		return "Monastir"
	case SubdivisionTN53:
		return "Mahdia"
	case SubdivisionTN61:
		return "Sfax"
	case SubdivisionTN71:
		return "Gafsa"
	case SubdivisionTN72:
		return "Tozeur"
	case SubdivisionTN73:
		return "Kebili"
	case SubdivisionTN81:
		return "Gabès"
	case SubdivisionTN82:
		return "Medenine"
	case SubdivisionTN83:
		return "Tataouine"
	case SubdivisionTO01:
		return "'Eua"
	case SubdivisionTO02:
		return "Ha'apai"
	case SubdivisionTO03:
		return "Niuas"
	case SubdivisionTO04:
		return "Tongatapu"
	case SubdivisionTO05:
		return "Vava'u"
	case SubdivisionTR01:
		return "Adana"
	case SubdivisionTR02:
		return "Adıyaman"
	case SubdivisionTR03:
		return "Afyonkarahisar"
	case SubdivisionTR04:
		return "Ağrı"
	case SubdivisionTR05:
		return "Amasya"
	case SubdivisionTR06:
		return "Ankara"
	case SubdivisionTR07:
		return "Antalya"
	case SubdivisionTR08:
		return "Artvin"
	case SubdivisionTR09:
		return "Aydın"
	case SubdivisionTR10:
		return "Balıkesir"
	case SubdivisionTR11:
		return "Bilecik"
	case SubdivisionTR12:
		return "Bingöl"
	case SubdivisionTR13:
		return "Bitlis"
	case SubdivisionTR14:
		return "Bolu"
	case SubdivisionTR15:
		return "Burdur"
	case SubdivisionTR16:
		return "Bursa"
	case SubdivisionTR17:
		return "Çanakkale"
	case SubdivisionTR18:
		return "Çankırı"
	case SubdivisionTR19:
		return "Çorum"
	case SubdivisionTR20:
		return "Denizli"
	case SubdivisionTR21:
		return "Diyarbakır"
	case SubdivisionTR22:
		return "Edirne"
	case SubdivisionTR23:
		return "Elazığ"
	case SubdivisionTR24:
		return "Erzincan"
	case SubdivisionTR25:
		return "Erzurum"
	case SubdivisionTR26:
		return "Eskişehir"
	case SubdivisionTR27:
		return "Gaziantep"
	case SubdivisionTR28:
		return "Giresun"
	case SubdivisionTR29:
		return "Gümüşhane"
	case SubdivisionTR30:
		return "Hakkâri"
	case SubdivisionTR31:
		return "Hatay"
	case SubdivisionTR32:
		return "Isparta"
	case SubdivisionTR33:
		return "Mersin"
	case SubdivisionTR34:
		return "İstanbul"
	case SubdivisionTR35:
		return "İzmir"
	case SubdivisionTR36:
		return "Kars"
	case SubdivisionTR37:
		return "Kastamonu"
	case SubdivisionTR38:
		return "Kayseri"
	case SubdivisionTR39:
		return "Kırklareli"
	case SubdivisionTR40:
		return "Kırşehir"
	case SubdivisionTR41:
		return "Kocaeli"
	case SubdivisionTR42:
		return "Konya"
	case SubdivisionTR43:
		return "Kütahya"
	case SubdivisionTR44:
		return "Malatya"
	case SubdivisionTR45:
		return "Manisa"
	case SubdivisionTR46:
		return "Kahramanmaraş"
	case SubdivisionTR47:
		return "Mardin"
	case SubdivisionTR48:
		return "Muğla"
	case SubdivisionTR49:
		return "Muş"
	case SubdivisionTR50:
		return "Nevşehir"
	case SubdivisionTR51:
		return "Niğde"
	case SubdivisionTR52:
		return "Ordu"
	case SubdivisionTR53:
		return "Rize"
	case SubdivisionTR54:
		return "Sakarya"
	case SubdivisionTR55:
		return "Samsun"
	case SubdivisionTR56:
		return "Siirt"
	case SubdivisionTR57:
		return "Sinop"
	case SubdivisionTR58:
		return "Sivas"
	case SubdivisionTR59:
		return "Tekirdağ"
	case SubdivisionTR60:
		return "Tokat"
	case SubdivisionTR61:
		return "Trabzon"
	case SubdivisionTR62:
		return "Tunceli"
	case SubdivisionTR63:
		return "Şanlıurfa"
	case SubdivisionTR64:
		return "Uşak"
	case SubdivisionTR65:
		return "Van"
	case SubdivisionTR66:
		return "Yozgat"
	case SubdivisionTR67:
		return "Zonguldak"
	case SubdivisionTR68:
		return "Aksaray"
	case SubdivisionTR69:
		return "Bayburt"
	case SubdivisionTR70:
		return "Karaman"
	case SubdivisionTR71:
		return "Kırıkkale"
	case SubdivisionTR72:
		return "Batman"
	case SubdivisionTR73:
		return "Şırnak"
	case SubdivisionTR74:
		return "Bartın"
	case SubdivisionTR75:
		return "Ardahan"
	case SubdivisionTR76:
		return "Iğdır"
	case SubdivisionTR77:
		return "Yalova"
	case SubdivisionTR78:
		return "Karabük"
	case SubdivisionTR79:
		return "Kilis"
	case SubdivisionTR80:
		return "Osmaniye"
	case SubdivisionTR81:
		return "Düzce"
	case SubdivisionTTARI:
		return "Arima"
	case SubdivisionTTCHA:
		return "Chaguanas"
	case SubdivisionTTCTT:
		return "Couva-Tabaquite-Talparo"
	case SubdivisionTTDMN:
		return "Diego Martin"
	case SubdivisionTTETO:
		return "Eastern Tobago"
	case SubdivisionTTPED:
		return "Penal-Debe"
	case SubdivisionTTPOS:
		return "Port of Spain"
	case SubdivisionTTPRT:
		return "Princes Town"
	case SubdivisionTTPTF:
		return "Point Fortin"
	case SubdivisionTTRCM:
		return "Rio Claro-Mayaro"
	case SubdivisionTTSFO:
		return "San Fernando"
	case SubdivisionTTSGE:
		return "Sangre Grande"
	case SubdivisionTTSIP:
		return "Siparia"
	case SubdivisionTTSJL:
		return "San Juan-Laventille"
	case SubdivisionTTTUP:
		return "Tunapuna-Piarco"
	case SubdivisionTTWTO:
		return "Western Tobago"
	case SubdivisionTVFUN:
		return "Funafuti"
	case SubdivisionTVNIT:
		return "Niutao"
	case SubdivisionTVNKF:
		return "Nukufetau"
	case SubdivisionTVNKL:
		return "Nukulaelae"
	case SubdivisionTVNMA:
		return "Nanumea"
	case SubdivisionTVNMG:
		return "Nanumanga"
	case SubdivisionTVNUI:
		return "Nui"
	case SubdivisionTVVAI:
		return "Vaitupu"
	case SubdivisionTWCHA:
		return "Changhua"
	case SubdivisionTWCYI:
		return "Chiay City"
	case SubdivisionTWCYQ:
		return "Chiayi"
	case SubdivisionTWHSQ:
		return "Hsinchu"
	case SubdivisionTWHSZ:
		return "Hsinchui City"
	case SubdivisionTWHUA:
		return "Hualien"
	case SubdivisionTWILA:
		return "Ilan"
	case SubdivisionTWKEE:
		return "Keelung City"
	case SubdivisionTWKHH:
		return "Kaohsiung City"
	case SubdivisionTWKHQ:
		return "Kaohsiung"
	case SubdivisionTWMIA:
		return "Miaoli"
	case SubdivisionTWNAN:
		return "Nantou"
	case SubdivisionTWPEN:
		return "Penghu"
	case SubdivisionTWPIF:
		return "Pingtung"
	case SubdivisionTWTAO:
		return "Taoyuan"
	case SubdivisionTWTNN:
		return "Tainan City"
	case SubdivisionTWTNQ:
		return "Tainan"
	case SubdivisionTWTPE:
		return "Taipei City"
	case SubdivisionTWTPQ:
		return "Taipei"
	case SubdivisionTWTTT:
		return "Taitung"
	case SubdivisionTWTXG:
		return "Taichung City"
	case SubdivisionTWTXQ:
		return "Taichung"
	case SubdivisionTWYUN:
		return "Yunlin"
	case SubdivisionTZ01:
		return "Arusha"
	case SubdivisionTZ02:
		return "Dar-es-Salaam"
	case SubdivisionTZ03:
		return "Dodoma"
	case SubdivisionTZ04:
		return "Iringa"
	case SubdivisionTZ05:
		return "Kagera"
	case SubdivisionTZ06:
		return "Kaskazini Pemba"
	case SubdivisionTZ07:
		return "Kaskazini Unguja"
	case SubdivisionTZ08:
		return "Kigoma"
	case SubdivisionTZ09:
		return "Kilimanjaro"
	case SubdivisionTZ10:
		return "Kusini Pemba"
	case SubdivisionTZ11:
		return "Kusini Unguja"
	case SubdivisionTZ12:
		return "Lindi"
	case SubdivisionTZ13:
		return "Mara"
	case SubdivisionTZ14:
		return "Mbeya"
	case SubdivisionTZ15:
		return "Mjini Magharibi"
	case SubdivisionTZ16:
		return "Morogoro"
	case SubdivisionTZ17:
		return "Mtwara"
	case SubdivisionTZ18:
		return "Mwanza"
	case SubdivisionTZ19:
		return "Pwani"
	case SubdivisionTZ20:
		return "Rukwa"
	case SubdivisionTZ21:
		return "Ruvuma"
	case SubdivisionTZ22:
		return "Shinyanga"
	case SubdivisionTZ23:
		return "Singida"
	case SubdivisionTZ24:
		return "Tabora"
	case SubdivisionTZ25:
		return "Tanga"
	case SubdivisionTZ26:
		return "Manyara"
	case SubdivisionUA05:
		return "Vinnyts'ka Oblast'"
	case SubdivisionUA07:
		return "Volyns'ka Oblast'"
	case SubdivisionUA09:
		return "Luhans'ka Oblast'"
	case SubdivisionUA12:
		return "Dnipropetrovs'ka Oblast'"
	case SubdivisionUA14:
		return "Donets'ka Oblast'"
	case SubdivisionUA18:
		return "Zhytomyrs'ka Oblast'"
	case SubdivisionUA21:
		return "Zakarpats'ka Oblast'"
	case SubdivisionUA23:
		return "Zaporiz'ka Oblast'"
	case SubdivisionUA26:
		return "Ivano-Frankivs'ka Oblast'"
	case SubdivisionUA30:
		return "Kyïvs'ka mis'ka rada"
	case SubdivisionUA32:
		return "Kyïvs'ka Oblast'"
	case SubdivisionUA35:
		return "Kirovohrads'ka Oblast'"
	case SubdivisionUA40:
		return "Sevastopol"
	case SubdivisionUA43:
		return "Respublika Krym"
	case SubdivisionUA46:
		return "L'vivs'ka Oblast'"
	case SubdivisionUA48:
		return "Mykolaïvs'ka Oblast'"
	case SubdivisionUA51:
		return "Odes'ka Oblast'"
	case SubdivisionUA53:
		return "Poltavs'ka Oblast'"
	case SubdivisionUA56:
		return "Rivnens'ka Oblast'"
	case SubdivisionUA59:
		return "Sums 'ka Oblast'"
	case SubdivisionUA61:
		return "Ternopil's'ka Oblast'"
	case SubdivisionUA63:
		return "Kharkivs'ka Oblast'"
	case SubdivisionUA65:
		return "Khersons'ka Oblast'"
	case SubdivisionUA68:
		return "Khmel'nyts'ka Oblast'"
	case SubdivisionUA71:
		return "Cherkas'ka Oblast'"
	case SubdivisionUA74:
		return "Chernihivs'ka Oblast'"
	case SubdivisionUA77:
		return "Chernivets'ka Oblast'"
	case SubdivisionUG101:
		return "Kalangala"
	case SubdivisionUG102:
		return "Kampala"
	case SubdivisionUG103:
		return "Kiboga"
	case SubdivisionUG104:
		return "Luwero"
	case SubdivisionUG105:
		return "Masaka"
	case SubdivisionUG106:
		return "Mpigi"
	case SubdivisionUG107:
		return "Mubende"
	case SubdivisionUG108:
		return "Mukono"
	case SubdivisionUG109:
		return "Nakasongola"
	case SubdivisionUG110:
		return "Rakai"
	case SubdivisionUG111:
		return "Sembabule"
	case SubdivisionUG112:
		return "Kayunga"
	case SubdivisionUG113:
		return "Wakiso"
	case SubdivisionUG114:
		return "Mityana"
	case SubdivisionUG115:
		return "Nakaseke"
	case SubdivisionUG116:
		return "Lyantonde"
	case SubdivisionUG201:
		return "Bugiri"
	case SubdivisionUG202:
		return "Busia"
	case SubdivisionUG203:
		return "Iganga"
	case SubdivisionUG204:
		return "Jinja"
	case SubdivisionUG205:
		return "Kamuli"
	case SubdivisionUG206:
		return "Kapchorwa"
	case SubdivisionUG207:
		return "Katakwi"
	case SubdivisionUG208:
		return "Kumi"
	case SubdivisionUG209:
		return "Mbale"
	case SubdivisionUG210:
		return "Pallisa"
	case SubdivisionUG211:
		return "Soroti"
	case SubdivisionUG212:
		return "Tororo"
	case SubdivisionUG213:
		return "Kaberamaido"
	case SubdivisionUG214:
		return "Mayuge"
	case SubdivisionUG215:
		return "Sironko"
	case SubdivisionUG216:
		return "Amuria"
	case SubdivisionUG217:
		return "Budaka"
	case SubdivisionUG218:
		return "Bukwa"
	case SubdivisionUG219:
		return "Butaleja"
	case SubdivisionUG220:
		return "Kaliro"
	case SubdivisionUG221:
		return "Manafwa"
	case SubdivisionUG222:
		return "Namutumba"
	case SubdivisionUG223:
		return "Bududa"
	case SubdivisionUG224:
		return "Bukedea"
	case SubdivisionUG301:
		return "Adjumani"
	case SubdivisionUG302:
		return "Apac"
	case SubdivisionUG303:
		return "Arua"
	case SubdivisionUG304:
		return "Gulu"
	case SubdivisionUG305:
		return "Kitgum"
	case SubdivisionUG306:
		return "Kotido"
	case SubdivisionUG307:
		return "Lira"
	case SubdivisionUG308:
		return "Moroto"
	case SubdivisionUG309:
		return "Moyo"
	case SubdivisionUG310:
		return "Nebbi"
	case SubdivisionUG311:
		return "Nakapiripirit"
	case SubdivisionUG312:
		return "Pader"
	case SubdivisionUG313:
		return "Yumbe"
	case SubdivisionUG314:
		return "Amolatar"
	case SubdivisionUG315:
		return "Kaabong"
	case SubdivisionUG316:
		return "Koboko"
	case SubdivisionUG317:
		return "Abim"
	case SubdivisionUG318:
		return "Dokolo"
	case SubdivisionUG319:
		return "Amuru"
	case SubdivisionUG320:
		return "Maracha"
	case SubdivisionUG321:
		return "Oyam"
	case SubdivisionUG401:
		return "Bundibugyo"
	case SubdivisionUG402:
		return "Bushenyi"
	case SubdivisionUG403:
		return "Hoima"
	case SubdivisionUG404:
		return "Kabale"
	case SubdivisionUG405:
		return "Kabarole"
	case SubdivisionUG406:
		return "Kasese"
	case SubdivisionUG407:
		return "Kibaale"
	case SubdivisionUG408:
		return "Kisoro"
	case SubdivisionUG409:
		return "Masindi"
	case SubdivisionUG410:
		return "Mbarara"
	case SubdivisionUG411:
		return "Ntungamo"
	case SubdivisionUG412:
		return "Rukungiri"
	case SubdivisionUG413:
		return "Kamwenge"
	case SubdivisionUG414:
		return "Kanungu"
	case SubdivisionUG415:
		return "Kyenjojo"
	case SubdivisionUG416:
		return "Ibanda"
	case SubdivisionUG417:
		return "Isingiro"
	case SubdivisionUG418:
		return "Kiruhura"
	case SubdivisionUG419:
		return "Buliisa"
	case SubdivisionUGC:
		return "Central"
	case SubdivisionUGE:
		return "Eastern"
	case SubdivisionUGN:
		return "Northern"
	case SubdivisionUGW:
		return "Western"
	case SubdivisionUM67:
		return "Johnston Atoll"
	case SubdivisionUM71:
		return "Midway Islands"
	case SubdivisionUM76:
		return "Navassa Island"
	case SubdivisionUM79:
		return "Wake Island"
	case SubdivisionUM81:
		return "Baker Island"
	case SubdivisionUM84:
		return "Howland Island"
	case SubdivisionUM86:
		return "Jarvis Island"
	case SubdivisionUM89:
		return "Kingman Reef"
	case SubdivisionUM95:
		return "Palmyra Atoll"
	case SubdivisionUSAK:
		return "Alaska"
	case SubdivisionUSAL:
		return "Alabama"
	case SubdivisionUSAR:
		return "Arkansas"
	case SubdivisionUSAS:
		return "American Samoa"
	case SubdivisionUSAZ:
		return "Arizona"
	case SubdivisionUSCA:
		return "California"
	case SubdivisionUSCO:
		return "Colorado"
	case SubdivisionUSCT:
		return "Connecticut"
	case SubdivisionUSDC:
		return "District of Columbia"
	case SubdivisionUSDE:
		return "Delaware"
	case SubdivisionUSFL:
		return "Florida"
	case SubdivisionUSGA:
		return "Georgia"
	case SubdivisionUSGU:
		return "Guam"
	case SubdivisionUSHI:
		return "Hawaii"
	case SubdivisionUSIA:
		return "Iowa"
	case SubdivisionUSID:
		return "Idaho"
	case SubdivisionUSIL:
		return "Illinois"
	case SubdivisionUSIN:
		return "Indiana"
	case SubdivisionUSKS:
		return "Kansas"
	case SubdivisionUSKY:
		return "Kentucky"
	case SubdivisionUSLA:
		return "Louisiana"
	case SubdivisionUSMA:
		return "Massachusetts"
	case SubdivisionUSMD:
		return "Maryland"
	case SubdivisionUSME:
		return "Maine"
	case SubdivisionUSMI:
		return "Michigan"
	case SubdivisionUSMN:
		return "Minnesota"
	case SubdivisionUSMO:
		return "Missouri"
	case SubdivisionUSMP:
		return "Northern Mariana Islands"
	case SubdivisionUSMS:
		return "Mississippi"
	case SubdivisionUSMT:
		return "Montana"
	case SubdivisionUSNC:
		return "North Carolina"
	case SubdivisionUSND:
		return "North Dakota"
	case SubdivisionUSNE:
		return "Nebraska"
	case SubdivisionUSNH:
		return "New Hampshire"
	case SubdivisionUSNJ:
		return "New Jersey"
	case SubdivisionUSNM:
		return "New Mexico"
	case SubdivisionUSNV:
		return "Nevada"
	case SubdivisionUSNY:
		return "New York"
	case SubdivisionUSOH:
		return "Ohio"
	case SubdivisionUSOK:
		return "Oklahoma"
	case SubdivisionUSOR:
		return "Oregon"
	case SubdivisionUSPA:
		return "Pennsylvania"
	case SubdivisionUSPR:
		return "Puerto Rico"
	case SubdivisionUSRI:
		return "Rhode Island"
	case SubdivisionUSSC:
		return "South Carolina"
	case SubdivisionUSSD:
		return "South Dakota"
	case SubdivisionUSTN:
		return "Tennessee"
	case SubdivisionUSTX:
		return "Texas"
	case SubdivisionUSUM:
		return "United States Minor Outlying Islands"
	case SubdivisionUSUT:
		return "Utah"
	case SubdivisionUSVA:
		return "Virginia"
	case SubdivisionUSVI:
		return "Virgin Islands"
	case SubdivisionUSVT:
		return "Vermont"
	case SubdivisionUSWA:
		return "Washington"
	case SubdivisionUSWI:
		return "Wisconsin"
	case SubdivisionUSWV:
		return "West Virginia"
	case SubdivisionUSWY:
		return "Wyoming"
	case SubdivisionUYAR:
		return "Artigas"
	case SubdivisionUYCA:
		return "Canelones"
	case SubdivisionUYCL:
		return "Cerro Largo"
	case SubdivisionUYCO:
		return "Colonia"
	case SubdivisionUYDU:
		return "Durazno"
	case SubdivisionUYFD:
		return "Florida"
	case SubdivisionUYFS:
		return "Flores"
	case SubdivisionUYLA:
		return "Lavalleja"
	case SubdivisionUYMA:
		return "Maldonado"
	case SubdivisionUYMO:
		return "Montevideo"
	case SubdivisionUYPA:
		return "Paysandú"
	case SubdivisionUYRN:
		return "Río Negro"
	case SubdivisionUYRO:
		return "Rocha"
	case SubdivisionUYRV:
		return "Rivera"
	case SubdivisionUYSA:
		return "Salto"
	case SubdivisionUYSJ:
		return "San José"
	case SubdivisionUYSO:
		return "Soriano"
	case SubdivisionUYTA:
		return "Tacuarembó"
	case SubdivisionUYTT:
		return "Treinta y Tres"
	case SubdivisionUZAN:
		return "Andijon"
	case SubdivisionUZBU:
		return "Buxoro"
	case SubdivisionUZFA:
		return "Farg'ona"
	case SubdivisionUZJI:
		return "Jizzax"
	case SubdivisionUZNG:
		return "Namangan"
	case SubdivisionUZNW:
		return "Navoiy"
	case SubdivisionUZQA:
		return "Qashqadaryo"
	case SubdivisionUZQR:
		return "Qoraqalpog'iston Respublikasi"
	case SubdivisionUZSA:
		return "Samarqand"
	case SubdivisionUZSI:
		return "Sirdaryo"
	case SubdivisionUZSU:
		return "Surxondaryo"
	case SubdivisionUZTK:
		return "Toshkent"
	case SubdivisionUZTO:
		return "Toshkent"
	case SubdivisionUZXO:
		return "Xorazm"
	case SubdivisionVC01:
		return "Charlotte"
	case SubdivisionVC02:
		return "Saint Andrew"
	case SubdivisionVC03:
		return "Saint David"
	case SubdivisionVC04:
		return "Saint George"
	case SubdivisionVC05:
		return "Saint Patrick"
	case SubdivisionVC06:
		return "Grenadines"
	case SubdivisionVEA:
		return "Distrito Federal"
	case SubdivisionVEB:
		return "Anzoátegui"
	case SubdivisionVEC:
		return "Apure"
	case SubdivisionVED:
		return "Aragua"
	case SubdivisionVEE:
		return "Barinas"
	case SubdivisionVEF:
		return "Bolívar"
	case SubdivisionVEG:
		return "Carabobo"
	case SubdivisionVEH:
		return "Cojedes"
	case SubdivisionVEI:
		return "Falcón"
	case SubdivisionVEJ:
		return "Guárico"
	case SubdivisionVEK:
		return "Lara"
	case SubdivisionVEL:
		return "Mérida"
	case SubdivisionVEM:
		return "Miranda"
	case SubdivisionVEN:
		return "Monagas"
	case SubdivisionVEO:
		return "Nueva Esparta"
	case SubdivisionVEP:
		return "Portuguesa"
	case SubdivisionVER:
		return "Sucre"
	case SubdivisionVES:
		return "Táchira"
	case SubdivisionVET:
		return "Trujillo"
	case SubdivisionVEU:
		return "Yaracuy"
	case SubdivisionVEV:
		return "Zulia"
	case SubdivisionVEW:
		return "Dependencias Federales"
	case SubdivisionVEX:
		return "Vargas"
	case SubdivisionVEY:
		return "Delta Amacuro"
	case SubdivisionVEZ:
		return "Amazonas"
	case SubdivisionVN01:
		return "Lai Châu"
	case SubdivisionVN02:
		return "Lào Cai"
	case SubdivisionVN03:
		return "Hà Giang"
	case SubdivisionVN04:
		return "Cao Bằng"
	case SubdivisionVN05:
		return "Sơn La"
	case SubdivisionVN06:
		return "Yên Bái"
	case SubdivisionVN07:
		return "Tuyên Quang"
	case SubdivisionVN09:
		return "Lạng Sơn"
	case SubdivisionVN13:
		return "Quảng Ninh"
	case SubdivisionVN14:
		return "Hoà Bình"
	case SubdivisionVN15:
		return "Hà Tây"
	case SubdivisionVN18:
		return "Ninh Bình"
	case SubdivisionVN20:
		return "Thái Bình"
	case SubdivisionVN21:
		return "Thanh Hóa"
	case SubdivisionVN22:
		return "Nghệ An"
	case SubdivisionVN23:
		return "Hà Tỉnh"
	case SubdivisionVN24:
		return "Quảng Bình"
	case SubdivisionVN25:
		return "Quảng Trị"
	case SubdivisionVN26:
		return "Thừa Thiên-Huế"
	case SubdivisionVN27:
		return "Quảng Nam"
	case SubdivisionVN28:
		return "Kon Tum"
	case SubdivisionVN29:
		return "Quảng Ngãi"
	case SubdivisionVN30:
		return "Gia Lai"
	case SubdivisionVN31:
		return "Bình Định"
	case SubdivisionVN32:
		return "Phú Yên"
	case SubdivisionVN33:
		return "Đắc Lắk"
	case SubdivisionVN34:
		return "Khánh Hòa"
	case SubdivisionVN35:
		return "Lâm Đồng"
	case SubdivisionVN36:
		return "Ninh Thuận"
	case SubdivisionVN37:
		return "Tây Ninh"
	case SubdivisionVN39:
		return "Đồng Nai"
	case SubdivisionVN40:
		return "Bình Thuận"
	case SubdivisionVN41:
		return "Long An"
	case SubdivisionVN43:
		return "Bà Rịa-Vũng Tàu"
	case SubdivisionVN44:
		return "An Giang"
	case SubdivisionVN45:
		return "Đồng Tháp"
	case SubdivisionVN46:
		return "Tiền Giang"
	case SubdivisionVN47:
		return "Kiên Giang"
	case SubdivisionVN49:
		return "Vĩnh Long"
	case SubdivisionVN50:
		return "Bến Tre"
	case SubdivisionVN51:
		return "Trà Vinh"
	case SubdivisionVN52:
		return "Sóc Trăng"
	case SubdivisionVN53:
		return "Bắc Kạn"
	case SubdivisionVN54:
		return "Bắc Giang"
	case SubdivisionVN55:
		return "Bạc Liêu"
	case SubdivisionVN56:
		return "Bắc Ninh"
	case SubdivisionVN57:
		return "Bình Dương"
	case SubdivisionVN58:
		return "Bình Phước"
	case SubdivisionVN59:
		return "Cà Mau"
	case SubdivisionVN61:
		return "Hải Duong"
	case SubdivisionVN63:
		return "Hà Nam"
	case SubdivisionVN66:
		return "Hưng Yên"
	case SubdivisionVN67:
		return "Nam Định"
	case SubdivisionVN68:
		return "Phú Thọ"
	case SubdivisionVN69:
		return "Thái Nguyên"
	case SubdivisionVN70:
		return "Vĩnh Phúc"
	case SubdivisionVN71:
		return "Điện Biên"
	case SubdivisionVN72:
		return "Đắk Nông"
	case SubdivisionVN73:
		return "Hậu Giang"
	case SubdivisionVNCT:
		return "Cần Thơ"
	case SubdivisionVNDN:
		return "Đà Nẵng"
	case SubdivisionVNHN:
		return "Hà Nội"
	case SubdivisionVNHP:
		return "Hải Phòng"
	case SubdivisionVNSG:
		return "Hồ Chí Minh [Sài Gòn]"
	case SubdivisionVUMAP:
		return "Malampa"
	case SubdivisionVUPAM:
		return "Pénama"
	case SubdivisionVUSAM:
		return "Sanma"
	case SubdivisionVUSEE:
		return "Shéfa"
	case SubdivisionVUTAE:
		return "Taféa"
	case SubdivisionVUTOB:
		return "Torba"
	case SubdivisionWSAA:
		return "A'ana"
	case SubdivisionWSAL:
		return "Aiga-i-le-Tai"
	case SubdivisionWSAT:
		return "Atua"
	case SubdivisionWSFA:
		return "Fa'asaleleaga"
	case SubdivisionWSGE:
		return "Gaga'emauga"
	case SubdivisionWSGI:
		return "Gagaifomauga"
	case SubdivisionWSPA:
		return "Palauli"
	case SubdivisionWSSA:
		return "Satupa'itea"
	case SubdivisionWSTU:
		return "Tuamasaga"
	case SubdivisionWSVF:
		return "Va'a-o-Fonoti"
	case SubdivisionWSVS:
		return "Vaisigano"
	case SubdivisionYEAB:
		return "Abyān"
	case SubdivisionYEAD:
		return "'Adan"
	case SubdivisionYEAM:
		return "'Amrān"
	case SubdivisionYEBA:
		return "Al Bayḑā'"
	case SubdivisionYEDA:
		return "Aḑ Ḑāli‘"
	case SubdivisionYEDH:
		return "Dhamār"
	case SubdivisionYEHD:
		return "Ḩaḑramawt"
	case SubdivisionYEHJ:
		return "Ḩajjah"
	case SubdivisionYEIB:
		return "Ibb"
	case SubdivisionYEJA:
		return "Al Jawf"
	case SubdivisionYELA:
		return "Laḩij"
	case SubdivisionYEMA:
		return "Ma'rib"
	case SubdivisionYEMR:
		return "Al Mahrah"
	case SubdivisionYEMU:
		return "Al Ḩudaydah"
	case SubdivisionYEMW:
		return "Al Maḩwīt"
	case SubdivisionYERA:
		return "Raymah"
	case SubdivisionYESD:
		return "Şa'dah"
	case SubdivisionYESH:
		return "Shabwah"
	case SubdivisionYESN:
		return "Şan'ā'"
	case SubdivisionYETA:
		return "Tā'izz"
	case SubdivisionZAEC:
		return "Eastern Cape"
	case SubdivisionZAFS:
		return "Free State"
	case SubdivisionZAGT:
		return "Gauteng"
	case SubdivisionZALP:
		return "Limpopo"
	case SubdivisionZAMP:
		return "Mpumalanga"
	case SubdivisionZANC:
		return "Northern Cape"
	case SubdivisionZANL:
		return "Kwazulu-Natal"
	case SubdivisionZANW:
		return "North-West (South Africa)"
	case SubdivisionZAWC:
		return "Western Cape"
	case SubdivisionZM01:
		return "Western"
	case SubdivisionZM02:
		return "Central"
	case SubdivisionZM03:
		return "Eastern"
	case SubdivisionZM04:
		return "Luapula"
	case SubdivisionZM05:
		return "Northern"
	case SubdivisionZM06:
		return "North-Western"
	case SubdivisionZM07:
		return "Southern (Zambia)"
	case SubdivisionZM08:
		return "Copperbelt"
	case SubdivisionZM09:
		return "Lusaka"
	case SubdivisionZWBU:
		return "Bulawayo"
	case SubdivisionZWHA:
		return "Harare"
	case SubdivisionZWMA:
		return "Manicaland"
	case SubdivisionZWMC:
		return "Mashonaland Central"
	case SubdivisionZWME:
		return "Mashonaland East"
	case SubdivisionZWMI:
		return "Midlands"
	case SubdivisionZWMN:
		return "Matabeleland North"
	case SubdivisionZWMS:
		return "Matabeleland South"
	case SubdivisionZWMV:
		return "Masvingo"
	case SubdivisionZWMW:
		return "Mashonaland West"
	}
	return UnknownMsg
}

// Country - returns a country of the subdivision
//
//nolint:cyclop,funlen,gocyclo
func (s SubdivisionCode) Country() CountryCode {
	switch s {
	case SubdivisionAD02:
		return AD
	case SubdivisionAD03:
		return AD
	case SubdivisionAD04:
		return AD
	case SubdivisionAD05:
		return AD
	case SubdivisionAD06:
		return AD
	case SubdivisionAD07:
		return AD
	case SubdivisionAD08:
		return AD
	case SubdivisionAEAJ:
		return AE
	case SubdivisionAEAZ:
		return AE
	case SubdivisionAEDU:
		return AE
	case SubdivisionAEFU:
		return AE
	case SubdivisionAERK:
		return AE
	case SubdivisionAESH:
		return AE
	case SubdivisionAEUQ:
		return AE
	case SubdivisionAFBAL:
		return AF
	case SubdivisionAFBAM:
		return AF
	case SubdivisionAFBDG:
		return AF
	case SubdivisionAFBDS:
		return AF
	case SubdivisionAFBGL:
		return AF
	case SubdivisionAFDAY:
		return AF
	case SubdivisionAFFRA:
		return AF
	case SubdivisionAFFYB:
		return AF
	case SubdivisionAFGHA:
		return AF
	case SubdivisionAFGHO:
		return AF
	case SubdivisionAFHEL:
		return AF
	case SubdivisionAFHER:
		return AF
	case SubdivisionAFJOW:
		return AF
	case SubdivisionAFKAB:
		return AF
	case SubdivisionAFKAN:
		return AF
	case SubdivisionAFKAP:
		return AF
	case SubdivisionAFKDZ:
		return AF
	case SubdivisionAFKHO:
		return AF
	case SubdivisionAFKNR:
		return AF
	case SubdivisionAFLAG:
		return AF
	case SubdivisionAFLOG:
		return AF
	case SubdivisionAFNAN:
		return AF
	case SubdivisionAFNIM:
		return AF
	case SubdivisionAFNUR:
		return AF
	case SubdivisionAFPAN:
		return AF
	case SubdivisionAFPAR:
		return AF
	case SubdivisionAFPIA:
		return AF
	case SubdivisionAFPKA:
		return AF
	case SubdivisionAFSAM:
		return AF
	case SubdivisionAFSAR:
		return AF
	case SubdivisionAFTAK:
		return AF
	case SubdivisionAFURU:
		return AF
	case SubdivisionAFWAR:
		return AF
	case SubdivisionAFZAB:
		return AF
	case SubdivisionAG03:
		return AG
	case SubdivisionAG04:
		return AG
	case SubdivisionAG05:
		return AG
	case SubdivisionAG06:
		return AG
	case SubdivisionAG07:
		return AG
	case SubdivisionAG08:
		return AG
	case SubdivisionAG10:
		return AG
	case SubdivisionAG11:
		return AG
	case SubdivisionAL01:
		return AL
	case SubdivisionAL02:
		return AL
	case SubdivisionAL03:
		return AL
	case SubdivisionAL04:
		return AL
	case SubdivisionAL05:
		return AL
	case SubdivisionAL06:
		return AL
	case SubdivisionAL07:
		return AL
	case SubdivisionAL08:
		return AL
	case SubdivisionAL09:
		return AL
	case SubdivisionAL10:
		return AL
	case SubdivisionAL11:
		return AL
	case SubdivisionAL12:
		return AL
	case SubdivisionALBR:
		return AL
	case SubdivisionALBU:
		return AL
	case SubdivisionALDI:
		return AL
	case SubdivisionALDL:
		return AL
	case SubdivisionALDR:
		return AL
	case SubdivisionALDV:
		return AL
	case SubdivisionALEL:
		return AL
	case SubdivisionALER:
		return AL
	case SubdivisionALFR:
		return AL
	case SubdivisionALGJ:
		return AL
	case SubdivisionALGR:
		return AL
	case SubdivisionALHA:
		return AL
	case SubdivisionALKA:
		return AL
	case SubdivisionALKB:
		return AL
	case SubdivisionALKC:
		return AL
	case SubdivisionALKO:
		return AL
	case SubdivisionALKR:
		return AL
	case SubdivisionALKU:
		return AL
	case SubdivisionALLB:
		return AL
	case SubdivisionALLE:
		return AL
	case SubdivisionALLU:
		return AL
	case SubdivisionALMK:
		return AL
	case SubdivisionALMM:
		return AL
	case SubdivisionALMR:
		return AL
	case SubdivisionALMT:
		return AL
	case SubdivisionALPG:
		return AL
	case SubdivisionALPQ:
		return AL
	case SubdivisionALPR:
		return AL
	case SubdivisionALPU:
		return AL
	case SubdivisionALSH:
		return AL
	case SubdivisionALSK:
		return AL
	case SubdivisionALSR:
		return AL
	case SubdivisionALTE:
		return AL
	case SubdivisionALTP:
		return AL
	case SubdivisionALTR:
		return AL
	case SubdivisionALVL:
		return AL
	case SubdivisionAMAG:
		return AM
	case SubdivisionAMAR:
		return AM
	case SubdivisionAMAV:
		return AM
	case SubdivisionAMER:
		return AM
	case SubdivisionAMGR:
		return AM
	case SubdivisionAMKT:
		return AM
	case SubdivisionAMLO:
		return AM
	case SubdivisionAMSH:
		return AM
	case SubdivisionAMSU:
		return AM
	case SubdivisionAMTV:
		return AM
	case SubdivisionAMVD:
		return AM
	case SubdivisionAOBGO:
		return AO
	case SubdivisionAOBGU:
		return AO
	case SubdivisionAOBIE:
		return AO
	case SubdivisionAOCAB:
		return AO
	case SubdivisionAOCCU:
		return AO
	case SubdivisionAOCNN:
		return AO
	case SubdivisionAOCNO:
		return AO
	case SubdivisionAOCUS:
		return AO
	case SubdivisionAOHUA:
		return AO
	case SubdivisionAOHUI:
		return AO
	case SubdivisionAOLNO:
		return AO
	case SubdivisionAOLSU:
		return AO
	case SubdivisionAOLUA:
		return AO
	case SubdivisionAOMAL:
		return AO
	case SubdivisionAOMOX:
		return AO
	case SubdivisionAONAM:
		return AO
	case SubdivisionAOUIG:
		return AO
	case SubdivisionAOZAI:
		return AO
	case SubdivisionARA:
		return AR
	case SubdivisionARB:
		return AR
	case SubdivisionARC:
		return AR
	case SubdivisionARD:
		return AR
	case SubdivisionARE:
		return AR
	case SubdivisionARG:
		return AR
	case SubdivisionARH:
		return AR
	case SubdivisionARJ:
		return AR
	case SubdivisionARK:
		return AR
	case SubdivisionARL:
		return AR
	case SubdivisionARM:
		return AR
	case SubdivisionARN:
		return AR
	case SubdivisionARP:
		return AR
	case SubdivisionARQ:
		return AR
	case SubdivisionARR:
		return AR
	case SubdivisionARS:
		return AR
	case SubdivisionART:
		return AR
	case SubdivisionARU:
		return AR
	case SubdivisionARV:
		return AR
	case SubdivisionARW:
		return AR
	case SubdivisionARX:
		return AR
	case SubdivisionARY:
		return AR
	case SubdivisionARZ:
		return AR
	case SubdivisionAT1:
		return AT
	case SubdivisionAT2:
		return AT
	case SubdivisionAT3:
		return AT
	case SubdivisionAT4:
		return AT
	case SubdivisionAT5:
		return AT
	case SubdivisionAT6:
		return AT
	case SubdivisionAT7:
		return AT
	case SubdivisionAT8:
		return AT
	case SubdivisionAT9:
		return AT
	case SubdivisionAUACT:
		return AU
	case SubdivisionAUNSW:
		return AU
	case SubdivisionAUNT:
		return AU
	case SubdivisionAUQLD:
		return AU
	case SubdivisionAUSA:
		return AU
	case SubdivisionAUTAS:
		return AU
	case SubdivisionAUVIC:
		return AU
	case SubdivisionAUWA:
		return AU
	case SubdivisionAZABS:
		return AZ
	case SubdivisionAZAGA:
		return AZ
	case SubdivisionAZAGC:
		return AZ
	case SubdivisionAZAGM:
		return AZ
	case SubdivisionAZAGS:
		return AZ
	case SubdivisionAZAGU:
		return AZ
	case SubdivisionAZAST:
		return AZ
	case SubdivisionAZBA:
		return AZ
	case SubdivisionAZBAB:
		return AZ
	case SubdivisionAZBAL:
		return AZ
	case SubdivisionAZBAR:
		return AZ
	case SubdivisionAZBEY:
		return AZ
	case SubdivisionAZBIL:
		return AZ
	case SubdivisionAZCAB:
		return AZ
	case SubdivisionAZCAL:
		return AZ
	case SubdivisionAZCUL:
		return AZ
	case SubdivisionAZDAS:
		return AZ
	case SubdivisionAZFUZ:
		return AZ
	case SubdivisionAZGA:
		return AZ
	case SubdivisionAZGAD:
		return AZ
	case SubdivisionAZGOR:
		return AZ
	case SubdivisionAZGOY:
		return AZ
	case SubdivisionAZGYG:
		return AZ
	case SubdivisionAZHAC:
		return AZ
	case SubdivisionAZIMI:
		return AZ
	case SubdivisionAZISM:
		return AZ
	case SubdivisionAZKAL:
		return AZ
	case SubdivisionAZKAN:
		return AZ
	case SubdivisionAZKUR:
		return AZ
	case SubdivisionAZLA:
		return AZ
	case SubdivisionAZLAC:
		return AZ
	case SubdivisionAZLAN:
		return AZ
	case SubdivisionAZLER:
		return AZ
	case SubdivisionAZMAS:
		return AZ
	case SubdivisionAZMI:
		return AZ
	case SubdivisionAZNA:
		return AZ
	case SubdivisionAZNEF:
		return AZ
	case SubdivisionAZNV:
		return AZ
	case SubdivisionAZNX:
		return AZ
	case SubdivisionAZOGU:
		return AZ
	case SubdivisionAZORD:
		return AZ
	case SubdivisionAZQAB:
		return AZ
	case SubdivisionAZQAX:
		return AZ
	case SubdivisionAZQAZ:
		return AZ
	case SubdivisionAZQBA:
		return AZ
	case SubdivisionAZQBI:
		return AZ
	case SubdivisionAZQOB:
		return AZ
	case SubdivisionAZQUS:
		return AZ
	case SubdivisionAZSA:
		return AZ
	case SubdivisionAZSAB:
		return AZ
	case SubdivisionAZSAD:
		return AZ
	case SubdivisionAZSAH:
		return AZ
	case SubdivisionAZSAK:
		return AZ
	case SubdivisionAZSAL:
		return AZ
	case SubdivisionAZSAR:
		return AZ
	case SubdivisionAZSAT:
		return AZ
	case SubdivisionAZSBN:
		return AZ
	case SubdivisionAZSIY:
		return AZ
	case SubdivisionAZSKR:
		return AZ
	case SubdivisionAZSM:
		return AZ
	case SubdivisionAZSMI:
		return AZ
	case SubdivisionAZSMX:
		return AZ
	case SubdivisionAZSR:
		return AZ
	case SubdivisionAZSUS:
		return AZ
	case SubdivisionAZTAR:
		return AZ
	case SubdivisionAZTOV:
		return AZ
	case SubdivisionAZUCA:
		return AZ
	case SubdivisionAZXA:
		return AZ
	case SubdivisionAZXAC:
		return AZ
	case SubdivisionAZXCI:
		return AZ
	case SubdivisionAZXIZ:
		return AZ
	case SubdivisionAZXVD:
		return AZ
	case SubdivisionAZYAR:
		return AZ
	case SubdivisionAZYE:
		return AZ
	case SubdivisionAZYEV:
		return AZ
	case SubdivisionAZZAN:
		return AZ
	case SubdivisionAZZAQ:
		return AZ
	case SubdivisionAZZAR:
		return AZ
	case SubdivisionBA01:
		return BA
	case SubdivisionBA02:
		return BA
	case SubdivisionBA03:
		return BA
	case SubdivisionBA04:
		return BA
	case SubdivisionBA05:
		return BA
	case SubdivisionBA06:
		return BA
	case SubdivisionBA07:
		return BA
	case SubdivisionBA08:
		return BA
	case SubdivisionBA09:
		return BA
	case SubdivisionBA10:
		return BA
	case SubdivisionBABIH:
		return BA
	case SubdivisionBABRC:
		return BA
	case SubdivisionBASRP:
		return BA
	case SubdivisionBB01:
		return BB
	case SubdivisionBB02:
		return BB
	case SubdivisionBB03:
		return BB
	case SubdivisionBB04:
		return BB
	case SubdivisionBB05:
		return BB
	case SubdivisionBB06:
		return BB
	case SubdivisionBB07:
		return BB
	case SubdivisionBB08:
		return BB
	case SubdivisionBB09:
		return BB
	case SubdivisionBB10:
		return BB
	case SubdivisionBB11:
		return BB
	case SubdivisionBD01:
		return BD
	case SubdivisionBD02:
		return BD
	case SubdivisionBD03:
		return BD
	case SubdivisionBD04:
		return BD
	case SubdivisionBD05:
		return BD
	case SubdivisionBD06:
		return BD
	case SubdivisionBD07:
		return BD
	case SubdivisionBD08:
		return BD
	case SubdivisionBD09:
		return BD
	case SubdivisionBD10:
		return BD
	case SubdivisionBD11:
		return BD
	case SubdivisionBD12:
		return BD
	case SubdivisionBD13:
		return BD
	case SubdivisionBD14:
		return BD
	case SubdivisionBD15:
		return BD
	case SubdivisionBD16:
		return BD
	case SubdivisionBD17:
		return BD
	case SubdivisionBD18:
		return BD
	case SubdivisionBD19:
		return BD
	case SubdivisionBD20:
		return BD
	case SubdivisionBD21:
		return BD
	case SubdivisionBD22:
		return BD
	case SubdivisionBD23:
		return BD
	case SubdivisionBD24:
		return BD
	case SubdivisionBD25:
		return BD
	case SubdivisionBD26:
		return BD
	case SubdivisionBD27:
		return BD
	case SubdivisionBD28:
		return BD
	case SubdivisionBD29:
		return BD
	case SubdivisionBD30:
		return BD
	case SubdivisionBD31:
		return BD
	case SubdivisionBD32:
		return BD
	case SubdivisionBD33:
		return BD
	case SubdivisionBD34:
		return BD
	case SubdivisionBD35:
		return BD
	case SubdivisionBD36:
		return BD
	case SubdivisionBD37:
		return BD
	case SubdivisionBD38:
		return BD
	case SubdivisionBD39:
		return BD
	case SubdivisionBD40:
		return BD
	case SubdivisionBD41:
		return BD
	case SubdivisionBD42:
		return BD
	case SubdivisionBD43:
		return BD
	case SubdivisionBD44:
		return BD
	case SubdivisionBD45:
		return BD
	case SubdivisionBD46:
		return BD
	case SubdivisionBD47:
		return BD
	case SubdivisionBD48:
		return BD
	case SubdivisionBD49:
		return BD
	case SubdivisionBD50:
		return BD
	case SubdivisionBD51:
		return BD
	case SubdivisionBD52:
		return BD
	case SubdivisionBD53:
		return BD
	case SubdivisionBD54:
		return BD
	case SubdivisionBD55:
		return BD
	case SubdivisionBD56:
		return BD
	case SubdivisionBD57:
		return BD
	case SubdivisionBD58:
		return BD
	case SubdivisionBD59:
		return BD
	case SubdivisionBD60:
		return BD
	case SubdivisionBD61:
		return BD
	case SubdivisionBD62:
		return BD
	case SubdivisionBD63:
		return BD
	case SubdivisionBD64:
		return BD
	case SubdivisionBDA:
		return BD
	case SubdivisionBDB:
		return BD
	case SubdivisionBDC:
		return BD
	case SubdivisionBDD:
		return BD
	case SubdivisionBDE:
		return BD
	case SubdivisionBDF:
		return BD
	case SubdivisionBDG:
		return BD
	case SubdivisionBDH:
		return BD
	case SubdivisionBEBRU:
		return BE
	case SubdivisionBEVAN:
		return BE
	case SubdivisionBEVBR:
		return BE
	case SubdivisionBEVLG:
		return BE
	case SubdivisionBEVLI:
		return BE
	case SubdivisionBEVOV:
		return BE
	case SubdivisionBEVWV:
		return BE
	case SubdivisionBEWAL:
		return BE
	case SubdivisionBEWBR:
		return BE
	case SubdivisionBEWHT:
		return BE
	case SubdivisionBEWLG:
		return BE
	case SubdivisionBEWLX:
		return BE
	case SubdivisionBEWNA:
		return BE
	case SubdivisionBF01:
		return BF
	case SubdivisionBF02:
		return BF
	case SubdivisionBF03:
		return BF
	case SubdivisionBF04:
		return BF
	case SubdivisionBF05:
		return BF
	case SubdivisionBF06:
		return BF
	case SubdivisionBF07:
		return BF
	case SubdivisionBF08:
		return BF
	case SubdivisionBF09:
		return BF
	case SubdivisionBF10:
		return BF
	case SubdivisionBF11:
		return BF
	case SubdivisionBF12:
		return BF
	case SubdivisionBF13:
		return BF
	case SubdivisionBFBAL:
		return BF
	case SubdivisionBFBAM:
		return BF
	case SubdivisionBFBAN:
		return BF
	case SubdivisionBFBAZ:
		return BF
	case SubdivisionBFBGR:
		return BF
	case SubdivisionBFBLG:
		return BF
	case SubdivisionBFBLK:
		return BF
	case SubdivisionBFCOM:
		return BF
	case SubdivisionBFGAN:
		return BF
	case SubdivisionBFGNA:
		return BF
	case SubdivisionBFGOU:
		return BF
	case SubdivisionBFHOU:
		return BF
	case SubdivisionBFIOB:
		return BF
	case SubdivisionBFKAD:
		return BF
	case SubdivisionBFKEN:
		return BF
	case SubdivisionBFKMD:
		return BF
	case SubdivisionBFKMP:
		return BF
	case SubdivisionBFKOP:
		return BF
	case SubdivisionBFKOS:
		return BF
	case SubdivisionBFKOT:
		return BF
	case SubdivisionBFKOW:
		return BF
	case SubdivisionBFLER:
		return BF
	case SubdivisionBFLOR:
		return BF
	case SubdivisionBFMOU:
		return BF
	case SubdivisionBFNAM:
		return BF
	case SubdivisionBFNAO:
		return BF
	case SubdivisionBFNAY:
		return BF
	case SubdivisionBFNOU:
		return BF
	case SubdivisionBFOUB:
		return BF
	case SubdivisionBFOUD:
		return BF
	case SubdivisionBFPAS:
		return BF
	case SubdivisionBFPON:
		return BF
	case SubdivisionBFSEN:
		return BF
	case SubdivisionBFSIS:
		return BF
	case SubdivisionBFSMT:
		return BF
	case SubdivisionBFSNG:
		return BF
	case SubdivisionBFSOM:
		return BF
	case SubdivisionBFSOR:
		return BF
	case SubdivisionBFTAP:
		return BF
	case SubdivisionBFTUI:
		return BF
	case SubdivisionBFYAG:
		return BF
	case SubdivisionBFYAT:
		return BF
	case SubdivisionBFZIR:
		return BF
	case SubdivisionBFZON:
		return BF
	case SubdivisionBFZOU:
		return BF
	case SubdivisionBG01:
		return BG
	case SubdivisionBG02:
		return BG
	case SubdivisionBG03:
		return BG
	case SubdivisionBG04:
		return BG
	case SubdivisionBG05:
		return BG
	case SubdivisionBG06:
		return BG
	case SubdivisionBG07:
		return BG
	case SubdivisionBG08:
		return BG
	case SubdivisionBG09:
		return BG
	case SubdivisionBG10:
		return BG
	case SubdivisionBG11:
		return BG
	case SubdivisionBG12:
		return BG
	case SubdivisionBG13:
		return BG
	case SubdivisionBG14:
		return BG
	case SubdivisionBG15:
		return BG
	case SubdivisionBG16:
		return BG
	case SubdivisionBG17:
		return BG
	case SubdivisionBG18:
		return BG
	case SubdivisionBG19:
		return BG
	case SubdivisionBG20:
		return BG
	case SubdivisionBG21:
		return BG
	case SubdivisionBG22:
		return BG
	case SubdivisionBG23:
		return BG
	case SubdivisionBG24:
		return BG
	case SubdivisionBG25:
		return BG
	case SubdivisionBG26:
		return BG
	case SubdivisionBG27:
		return BG
	case SubdivisionBG28:
		return BG
	case SubdivisionBH13:
		return BH
	case SubdivisionBH14:
		return BH
	case SubdivisionBH15:
		return BH
	case SubdivisionBH16:
		return BH
	case SubdivisionBH17:
		return BH
	case SubdivisionBIBB:
		return BI
	case SubdivisionBIBL:
		return BI
	case SubdivisionBIBM:
		return BI
	case SubdivisionBIBR:
		return BI
	case SubdivisionBICA:
		return BI
	case SubdivisionBICI:
		return BI
	case SubdivisionBIGI:
		return BI
	case SubdivisionBIKI:
		return BI
	case SubdivisionBIKR:
		return BI
	case SubdivisionBIKY:
		return BI
	case SubdivisionBIMA:
		return BI
	case SubdivisionBIMU:
		return BI
	case SubdivisionBIMW:
		return BI
	case SubdivisionBING:
		return BI
	case SubdivisionBIRT:
		return BI
	case SubdivisionBIRY:
		return BI
	case SubdivisionBJAK:
		return BJ
	case SubdivisionBJAL:
		return BJ
	case SubdivisionBJAQ:
		return BJ
	case SubdivisionBJBO:
		return BJ
	case SubdivisionBJCO:
		return BJ
	case SubdivisionBJDO:
		return BJ
	case SubdivisionBJKO:
		return BJ
	case SubdivisionBJLI:
		return BJ
	case SubdivisionBJMO:
		return BJ
	case SubdivisionBJOU:
		return BJ
	case SubdivisionBJPL:
		return BJ
	case SubdivisionBJZO:
		return BJ
	case SubdivisionBNBE:
		return BN
	case SubdivisionBNBM:
		return BN
	case SubdivisionBNTE:
		return BN
	case SubdivisionBNTU:
		return BN
	case SubdivisionBOB:
		return BO
	case SubdivisionBOC:
		return BO
	case SubdivisionBOH:
		return BO
	case SubdivisionBOL:
		return BO
	case SubdivisionBON:
		return BO
	case SubdivisionBOO:
		return BO
	case SubdivisionBOP:
		return BO
	case SubdivisionBOS:
		return BO
	case SubdivisionBOT:
		return BO
	case SubdivisionBQBO:
		return BQ
	case SubdivisionBQSA:
		return BQ
	case SubdivisionBQSE:
		return BQ
	case SubdivisionBRAC:
		return BR
	case SubdivisionBRAL:
		return BR
	case SubdivisionBRAM:
		return BR
	case SubdivisionBRAP:
		return BR
	case SubdivisionBRBA:
		return BR
	case SubdivisionBRCE:
		return BR
	case SubdivisionBRDF:
		return BR
	case SubdivisionBRES:
		return BR
	case SubdivisionBRFN:
		return BR
	case SubdivisionBRGO:
		return BR
	case SubdivisionBRMA:
		return BR
	case SubdivisionBRMG:
		return BR
	case SubdivisionBRMS:
		return BR
	case SubdivisionBRMT:
		return BR
	case SubdivisionBRPA:
		return BR
	case SubdivisionBRPB:
		return BR
	case SubdivisionBRPE:
		return BR
	case SubdivisionBRPI:
		return BR
	case SubdivisionBRPR:
		return BR
	case SubdivisionBRRJ:
		return BR
	case SubdivisionBRRN:
		return BR
	case SubdivisionBRRO:
		return BR
	case SubdivisionBRRR:
		return BR
	case SubdivisionBRRS:
		return BR
	case SubdivisionBRSC:
		return BR
	case SubdivisionBRSE:
		return BR
	case SubdivisionBRSP:
		return BR
	case SubdivisionBRTO:
		return BR
	case SubdivisionBSAK:
		return BS
	case SubdivisionBSBI:
		return BS
	case SubdivisionBSBP:
		return BS
	case SubdivisionBSBY:
		return BS
	case SubdivisionBSCE:
		return BS
	case SubdivisionBSCI:
		return BS
	case SubdivisionBSCK:
		return BS
	case SubdivisionBSCO:
		return BS
	case SubdivisionBSCS:
		return BS
	case SubdivisionBSEG:
		return BS
	case SubdivisionBSEX:
		return BS
	case SubdivisionBSFP:
		return BS
	case SubdivisionBSGC:
		return BS
	case SubdivisionBSHI:
		return BS
	case SubdivisionBSHT:
		return BS
	case SubdivisionBSIN:
		return BS
	case SubdivisionBSLI:
		return BS
	case SubdivisionBSMC:
		return BS
	case SubdivisionBSMG:
		return BS
	case SubdivisionBSMI:
		return BS
	case SubdivisionBSNE:
		return BS
	case SubdivisionBSNO:
		return BS
	case SubdivisionBSNS:
		return BS
	case SubdivisionBSRC:
		return BS
	case SubdivisionBSRI:
		return BS
	case SubdivisionBSSA:
		return BS
	case SubdivisionBSSE:
		return BS
	case SubdivisionBSSO:
		return BS
	case SubdivisionBSSS:
		return BS
	case SubdivisionBSSW:
		return BS
	case SubdivisionBSWG:
		return BS
	case SubdivisionBT11:
		return BT
	case SubdivisionBT12:
		return BT
	case SubdivisionBT13:
		return BT
	case SubdivisionBT14:
		return BT
	case SubdivisionBT15:
		return BT
	case SubdivisionBT21:
		return BT
	case SubdivisionBT22:
		return BT
	case SubdivisionBT23:
		return BT
	case SubdivisionBT24:
		return BT
	case SubdivisionBT31:
		return BT
	case SubdivisionBT32:
		return BT
	case SubdivisionBT33:
		return BT
	case SubdivisionBT34:
		return BT
	case SubdivisionBT41:
		return BT
	case SubdivisionBT42:
		return BT
	case SubdivisionBT43:
		return BT
	case SubdivisionBT44:
		return BT
	case SubdivisionBT45:
		return BT
	case SubdivisionBTGA:
		return BT
	case SubdivisionBTTY:
		return BT
	case SubdivisionBWCE:
		return BW
	case SubdivisionBWGH:
		return BW
	case SubdivisionBWKG:
		return BW
	case SubdivisionBWKL:
		return BW
	case SubdivisionBWKW:
		return BW
	case SubdivisionBWNE:
		return BW
	case SubdivisionBWNW:
		return BW
	case SubdivisionBWSE:
		return BW
	case SubdivisionBWSO:
		return BW
	case SubdivisionBYBR:
		return BY
	case SubdivisionBYHM:
		return BY
	case SubdivisionBYHO:
		return BY
	case SubdivisionBYHR:
		return BY
	case SubdivisionBYMA:
		return BY
	case SubdivisionBYMI:
		return BY
	case SubdivisionBYVI:
		return BY
	case SubdivisionBZBZ:
		return BZ
	case SubdivisionBZCY:
		return BZ
	case SubdivisionBZCZL:
		return BZ
	case SubdivisionBZOW:
		return BZ
	case SubdivisionBZSC:
		return BZ
	case SubdivisionBZTOL:
		return BZ
	case SubdivisionCAAB:
		return CA
	case SubdivisionCABC:
		return CA
	case SubdivisionCAMB:
		return CA
	case SubdivisionCANB:
		return CA
	case SubdivisionCANL:
		return CA
	case SubdivisionCANS:
		return CA
	case SubdivisionCANT:
		return CA
	case SubdivisionCANU:
		return CA
	case SubdivisionCAON:
		return CA
	case SubdivisionCAPE:
		return CA
	case SubdivisionCAQC:
		return CA
	case SubdivisionCASK:
		return CA
	case SubdivisionCAYT:
		return CA
	case SubdivisionCDBC:
		return CD
	case SubdivisionCDBN:
		return CD
	case SubdivisionCDEQ:
		return CD
	case SubdivisionCDKA:
		return CD
	case SubdivisionCDKE:
		return CD
	case SubdivisionCDKN:
		return CD
	case SubdivisionCDKW:
		return CD
	case SubdivisionCDMA:
		return CD
	case SubdivisionCDNK:
		return CD
	case SubdivisionCDOR:
		return CD
	case SubdivisionCDSK:
		return CD
	case SubdivisionCFAC:
		return CF
	case SubdivisionCFBB:
		return CF
	case SubdivisionCFBGF:
		return CF
	case SubdivisionCFBK:
		return CF
	case SubdivisionCFHK:
		return CF
	case SubdivisionCFHM:
		return CF
	case SubdivisionCFHS:
		return CF
	case SubdivisionCFKB:
		return CF
	case SubdivisionCFKG:
		return CF
	case SubdivisionCFLB:
		return CF
	case SubdivisionCFMB:
		return CF
	case SubdivisionCFMP:
		return CF
	case SubdivisionCFNM:
		return CF
	case SubdivisionCFOP:
		return CF
	case SubdivisionCFSE:
		return CF
	case SubdivisionCFUK:
		return CF
	case SubdivisionCFVK:
		return CF
	case SubdivisionCG11:
		return CG
	case SubdivisionCG12:
		return CG
	case SubdivisionCG13:
		return CG
	case SubdivisionCG14:
		return CG
	case SubdivisionCG15:
		return CG
	case SubdivisionCG2:
		return CG
	case SubdivisionCG5:
		return CG
	case SubdivisionCG7:
		return CG
	case SubdivisionCG8:
		return CG
	case SubdivisionCG9:
		return CG
	case SubdivisionCGBZV:
		return CG
	case SubdivisionCHAG:
		return CH
	case SubdivisionCHAI:
		return CH
	case SubdivisionCHAR:
		return CH
	case SubdivisionCHBE:
		return CH
	case SubdivisionCHBL:
		return CH
	case SubdivisionCHBS:
		return CH
	case SubdivisionCHFR:
		return CH
	case SubdivisionCHGE:
		return CH
	case SubdivisionCHGL:
		return CH
	case SubdivisionCHGR:
		return CH
	case SubdivisionCHJU:
		return CH
	case SubdivisionCHLU:
		return CH
	case SubdivisionCHNE:
		return CH
	case SubdivisionCHNW:
		return CH
	case SubdivisionCHOW:
		return CH
	case SubdivisionCHSG:
		return CH
	case SubdivisionCHSH:
		return CH
	case SubdivisionCHSO:
		return CH
	case SubdivisionCHSZ:
		return CH
	case SubdivisionCHTG:
		return CH
	case SubdivisionCHTI:
		return CH
	case SubdivisionCHUR:
		return CH
	case SubdivisionCHVD:
		return CH
	case SubdivisionCHVS:
		return CH
	case SubdivisionCHZG:
		return CH
	case SubdivisionCHZH:
		return CH
	case SubdivisionCI01:
		return CI
	case SubdivisionCI02:
		return CI
	case SubdivisionCI03:
		return CI
	case SubdivisionCI04:
		return CI
	case SubdivisionCI05:
		return CI
	case SubdivisionCI06:
		return CI
	case SubdivisionCI07:
		return CI
	case SubdivisionCI08:
		return CI
	case SubdivisionCI09:
		return CI
	case SubdivisionCI10:
		return CI
	case SubdivisionCI11:
		return CI
	case SubdivisionCI12:
		return CI
	case SubdivisionCI13:
		return CI
	case SubdivisionCI14:
		return CI
	case SubdivisionCI15:
		return CI
	case SubdivisionCI16:
		return CI
	case SubdivisionCI17:
		return CI
	case SubdivisionCI18:
		return CI
	case SubdivisionCI19:
		return CI
	case SubdivisionCLAI:
		return CL
	case SubdivisionCLAN:
		return CL
	case SubdivisionCLAP:
		return CL
	case SubdivisionCLAR:
		return CL
	case SubdivisionCLAT:
		return CL
	case SubdivisionCLBI:
		return CL
	case SubdivisionCLCO:
		return CL
	case SubdivisionCLLI:
		return CL
	case SubdivisionCLLL:
		return CL
	case SubdivisionCLLR:
		return CL
	case SubdivisionCLMA:
		return CL
	case SubdivisionCLML:
		return CL
	case SubdivisionCLRM:
		return CL
	case SubdivisionCLTA:
		return CL
	case SubdivisionCLVS:
		return CL
	case SubdivisionCMAD:
		return CM
	case SubdivisionCMCE:
		return CM
	case SubdivisionCMEN:
		return CM
	case SubdivisionCMES:
		return CM
	case SubdivisionCMLT:
		return CM
	case SubdivisionCMNO:
		return CM
	case SubdivisionCMNW:
		return CM
	case SubdivisionCMOU:
		return CM
	case SubdivisionCMSU:
		return CM
	case SubdivisionCMSW:
		return CM
	case SubdivisionCNAH:
		return CN
	case SubdivisionCNBJ:
		return CN
	case SubdivisionCNCQ:
		return CN
	case SubdivisionCNFJ:
		return CN
	case SubdivisionCNGD:
		return CN
	case SubdivisionCNGS:
		return CN
	case SubdivisionCNGX:
		return CN
	case SubdivisionCNGZ:
		return CN
	case SubdivisionCNHA:
		return CN
	case SubdivisionCNHB:
		return CN
	case SubdivisionCNHE:
		return CN
	case SubdivisionCNHI:
		return CN
	case SubdivisionCNHK:
		return CN
	case SubdivisionCNHL:
		return CN
	case SubdivisionCNHN:
		return CN
	case SubdivisionCNJL:
		return CN
	case SubdivisionCNJS:
		return CN
	case SubdivisionCNJX:
		return CN
	case SubdivisionCNLN:
		return CN
	case SubdivisionCNMO:
		return CN
	case SubdivisionCNNM:
		return CN
	case SubdivisionCNNX:
		return CN
	case SubdivisionCNQH:
		return CN
	case SubdivisionCNSC:
		return CN
	case SubdivisionCNSD:
		return CN
	case SubdivisionCNSH:
		return CN
	case SubdivisionCNSN:
		return CN
	case SubdivisionCNSX:
		return CN
	case SubdivisionCNTJ:
		return CN
	case SubdivisionCNTW:
		return CN
	case SubdivisionCNXJ:
		return CN
	case SubdivisionCNXZ:
		return CN
	case SubdivisionCNYN:
		return CN
	case SubdivisionCNZJ:
		return CN
	case SubdivisionCOAMA:
		return CO
	case SubdivisionCOANT:
		return CO
	case SubdivisionCOARA:
		return CO
	case SubdivisionCOATL:
		return CO
	case SubdivisionCOBOL:
		return CO
	case SubdivisionCOBOY:
		return CO
	case SubdivisionCOCAL:
		return CO
	case SubdivisionCOCAQ:
		return CO
	case SubdivisionCOCAS:
		return CO
	case SubdivisionCOCAU:
		return CO
	case SubdivisionCOCES:
		return CO
	case SubdivisionCOCHO:
		return CO
	case SubdivisionCOCOR:
		return CO
	case SubdivisionCOCUN:
		return CO
	case SubdivisionCODC:
		return CO
	case SubdivisionCOGUA:
		return CO
	case SubdivisionCOGUV:
		return CO
	case SubdivisionCOHUI:
		return CO
	case SubdivisionCOLAG:
		return CO
	case SubdivisionCOMAG:
		return CO
	case SubdivisionCOMET:
		return CO
	case SubdivisionCONAR:
		return CO
	case SubdivisionCONSA:
		return CO
	case SubdivisionCOPUT:
		return CO
	case SubdivisionCOQUI:
		return CO
	case SubdivisionCORIS:
		return CO
	case SubdivisionCOSAN:
		return CO
	case SubdivisionCOSAP:
		return CO
	case SubdivisionCOSUC:
		return CO
	case SubdivisionCOTOL:
		return CO
	case SubdivisionCOVAC:
		return CO
	case SubdivisionCOVAU:
		return CO
	case SubdivisionCOVID:
		return CO
	case SubdivisionCRA:
		return CR
	case SubdivisionCRC:
		return CR
	case SubdivisionCRG:
		return CR
	case SubdivisionCRH:
		return CR
	case SubdivisionCRL:
		return CR
	case SubdivisionCRP:
		return CR
	case SubdivisionCRSJ:
		return CR
	case SubdivisionCU01:
		return CU
	case SubdivisionCU02:
		return CU
	case SubdivisionCU03:
		return CU
	case SubdivisionCU04:
		return CU
	case SubdivisionCU05:
		return CU
	case SubdivisionCU06:
		return CU
	case SubdivisionCU07:
		return CU
	case SubdivisionCU08:
		return CU
	case SubdivisionCU09:
		return CU
	case SubdivisionCU10:
		return CU
	case SubdivisionCU11:
		return CU
	case SubdivisionCU12:
		return CU
	case SubdivisionCU13:
		return CU
	case SubdivisionCU14:
		return CU
	case SubdivisionCU99:
		return CU
	case SubdivisionCVB:
		return CV
	case SubdivisionCVBR:
		return CV
	case SubdivisionCVBV:
		return CV
	case SubdivisionCVCA:
		return CV
	case SubdivisionCVCF:
		return CV
	case SubdivisionCVCR:
		return CV
	case SubdivisionCVMA:
		return CV
	case SubdivisionCVMO:
		return CV
	case SubdivisionCVPA:
		return CV
	case SubdivisionCVPN:
		return CV
	case SubdivisionCVPR:
		return CV
	case SubdivisionCVRB:
		return CV
	case SubdivisionCVRG:
		return CV
	case SubdivisionCVRS:
		return CV
	case SubdivisionCVS:
		return CV
	case SubdivisionCVSD:
		return CV
	case SubdivisionCVSF:
		return CV
	case SubdivisionCVSL:
		return CV
	case SubdivisionCVSM:
		return CV
	case SubdivisionCVSO:
		return CV
	case SubdivisionCVSS:
		return CV
	case SubdivisionCVSV:
		return CV
	case SubdivisionCVTA:
		return CV
	case SubdivisionCVTS:
		return CV
	case SubdivisionCY01:
		return CY
	case SubdivisionCY02:
		return CY
	case SubdivisionCY03:
		return CY
	case SubdivisionCY04:
		return CY
	case SubdivisionCY05:
		return CY
	case SubdivisionCY06:
		return CY
	case SubdivisionCZ10:
		return CZ
	case SubdivisionCZ101:
		return CZ
	case SubdivisionCZ102:
		return CZ
	case SubdivisionCZ103:
		return CZ
	case SubdivisionCZ104:
		return CZ
	case SubdivisionCZ105:
		return CZ
	case SubdivisionCZ106:
		return CZ
	case SubdivisionCZ107:
		return CZ
	case SubdivisionCZ108:
		return CZ
	case SubdivisionCZ109:
		return CZ
	case SubdivisionCZ110:
		return CZ
	case SubdivisionCZ111:
		return CZ
	case SubdivisionCZ112:
		return CZ
	case SubdivisionCZ113:
		return CZ
	case SubdivisionCZ114:
		return CZ
	case SubdivisionCZ115:
		return CZ
	case SubdivisionCZ116:
		return CZ
	case SubdivisionCZ117:
		return CZ
	case SubdivisionCZ118:
		return CZ
	case SubdivisionCZ119:
		return CZ
	case SubdivisionCZ120:
		return CZ
	case SubdivisionCZ121:
		return CZ
	case SubdivisionCZ122:
		return CZ
	case SubdivisionCZ20:
		return CZ
	case SubdivisionCZ201:
		return CZ
	case SubdivisionCZ202:
		return CZ
	case SubdivisionCZ203:
		return CZ
	case SubdivisionCZ204:
		return CZ
	case SubdivisionCZ205:
		return CZ
	case SubdivisionCZ206:
		return CZ
	case SubdivisionCZ207:
		return CZ
	case SubdivisionCZ208:
		return CZ
	case SubdivisionCZ209:
		return CZ
	case SubdivisionCZ20A:
		return CZ
	case SubdivisionCZ20B:
		return CZ
	case SubdivisionCZ20C:
		return CZ
	case SubdivisionCZ31:
		return CZ
	case SubdivisionCZ311:
		return CZ
	case SubdivisionCZ312:
		return CZ
	case SubdivisionCZ313:
		return CZ
	case SubdivisionCZ314:
		return CZ
	case SubdivisionCZ315:
		return CZ
	case SubdivisionCZ316:
		return CZ
	case SubdivisionCZ317:
		return CZ
	case SubdivisionCZ32:
		return CZ
	case SubdivisionCZ321:
		return CZ
	case SubdivisionCZ322:
		return CZ
	case SubdivisionCZ323:
		return CZ
	case SubdivisionCZ324:
		return CZ
	case SubdivisionCZ325:
		return CZ
	case SubdivisionCZ326:
		return CZ
	case SubdivisionCZ327:
		return CZ
	case SubdivisionCZ41:
		return CZ
	case SubdivisionCZ411:
		return CZ
	case SubdivisionCZ412:
		return CZ
	case SubdivisionCZ413:
		return CZ
	case SubdivisionCZ42:
		return CZ
	case SubdivisionCZ421:
		return CZ
	case SubdivisionCZ422:
		return CZ
	case SubdivisionCZ423:
		return CZ
	case SubdivisionCZ424:
		return CZ
	case SubdivisionCZ425:
		return CZ
	case SubdivisionCZ426:
		return CZ
	case SubdivisionCZ427:
		return CZ
	case SubdivisionCZ51:
		return CZ
	case SubdivisionCZ511:
		return CZ
	case SubdivisionCZ512:
		return CZ
	case SubdivisionCZ513:
		return CZ
	case SubdivisionCZ514:
		return CZ
	case SubdivisionCZ52:
		return CZ
	case SubdivisionCZ521:
		return CZ
	case SubdivisionCZ522:
		return CZ
	case SubdivisionCZ523:
		return CZ
	case SubdivisionCZ524:
		return CZ
	case SubdivisionCZ525:
		return CZ
	case SubdivisionCZ53:
		return CZ
	case SubdivisionCZ531:
		return CZ
	case SubdivisionCZ532:
		return CZ
	case SubdivisionCZ533:
		return CZ
	case SubdivisionCZ534:
		return CZ
	case SubdivisionCZ63:
		return CZ
	case SubdivisionCZ631:
		return CZ
	case SubdivisionCZ632:
		return CZ
	case SubdivisionCZ633:
		return CZ
	case SubdivisionCZ634:
		return CZ
	case SubdivisionCZ635:
		return CZ
	case SubdivisionCZ64:
		return CZ
	case SubdivisionCZ641:
		return CZ
	case SubdivisionCZ642:
		return CZ
	case SubdivisionCZ643:
		return CZ
	case SubdivisionCZ644:
		return CZ
	case SubdivisionCZ645:
		return CZ
	case SubdivisionCZ646:
		return CZ
	case SubdivisionCZ647:
		return CZ
	case SubdivisionCZ71:
		return CZ
	case SubdivisionCZ711:
		return CZ
	case SubdivisionCZ712:
		return CZ
	case SubdivisionCZ713:
		return CZ
	case SubdivisionCZ714:
		return CZ
	case SubdivisionCZ715:
		return CZ
	case SubdivisionCZ72:
		return CZ
	case SubdivisionCZ721:
		return CZ
	case SubdivisionCZ722:
		return CZ
	case SubdivisionCZ723:
		return CZ
	case SubdivisionCZ724:
		return CZ
	case SubdivisionCZ80:
		return CZ
	case SubdivisionCZ801:
		return CZ
	case SubdivisionCZ802:
		return CZ
	case SubdivisionCZ803:
		return CZ
	case SubdivisionCZ804:
		return CZ
	case SubdivisionCZ805:
		return CZ
	case SubdivisionCZ806:
		return CZ
	case SubdivisionDEBB:
		return DE
	case SubdivisionDEBE:
		return DE
	case SubdivisionDEBW:
		return DE
	case SubdivisionDEBY:
		return DE
	case SubdivisionDEHB:
		return DE
	case SubdivisionDEHE:
		return DE
	case SubdivisionDEHH:
		return DE
	case SubdivisionDEMV:
		return DE
	case SubdivisionDENI:
		return DE
	case SubdivisionDENW:
		return DE
	case SubdivisionDERP:
		return DE
	case SubdivisionDESH:
		return DE
	case SubdivisionDESL:
		return DE
	case SubdivisionDESN:
		return DE
	case SubdivisionDEST:
		return DE
	case SubdivisionDETH:
		return DE
	case SubdivisionDJAR:
		return DJ
	case SubdivisionDJAS:
		return DJ
	case SubdivisionDJDI:
		return DJ
	case SubdivisionDJDJ:
		return DJ
	case SubdivisionDJOB:
		return DJ
	case SubdivisionDJTA:
		return DJ
	case SubdivisionDK81:
		return DK
	case SubdivisionDK82:
		return DK
	case SubdivisionDK83:
		return DK
	case SubdivisionDK84:
		return DK
	case SubdivisionDK85:
		return DK
	case SubdivisionDM01:
		return DM
	case SubdivisionDM02:
		return DM
	case SubdivisionDM03:
		return DM
	case SubdivisionDM04:
		return DM
	case SubdivisionDM05:
		return DM
	case SubdivisionDM06:
		return DM
	case SubdivisionDM07:
		return DM
	case SubdivisionDM08:
		return DM
	case SubdivisionDM09:
		return DM
	case SubdivisionDM10:
		return DM
	case SubdivisionDO01:
		return DO
	case SubdivisionDO02:
		return DO
	case SubdivisionDO03:
		return DO
	case SubdivisionDO04:
		return DO
	case SubdivisionDO05:
		return DO
	case SubdivisionDO06:
		return DO
	case SubdivisionDO07:
		return DO
	case SubdivisionDO08:
		return DO
	case SubdivisionDO09:
		return DO
	case SubdivisionDO10:
		return DO
	case SubdivisionDO11:
		return DO
	case SubdivisionDO12:
		return DO
	case SubdivisionDO13:
		return DO
	case SubdivisionDO14:
		return DO
	case SubdivisionDO15:
		return DO
	case SubdivisionDO16:
		return DO
	case SubdivisionDO17:
		return DO
	case SubdivisionDO18:
		return DO
	case SubdivisionDO19:
		return DO
	case SubdivisionDO20:
		return DO
	case SubdivisionDO21:
		return DO
	case SubdivisionDO22:
		return DO
	case SubdivisionDO23:
		return DO
	case SubdivisionDO24:
		return DO
	case SubdivisionDO25:
		return DO
	case SubdivisionDO26:
		return DO
	case SubdivisionDO27:
		return DO
	case SubdivisionDO28:
		return DO
	case SubdivisionDO29:
		return DO
	case SubdivisionDO30:
		return DO
	case SubdivisionDZ01:
		return DZ
	case SubdivisionDZ02:
		return DZ
	case SubdivisionDZ03:
		return DZ
	case SubdivisionDZ04:
		return DZ
	case SubdivisionDZ05:
		return DZ
	case SubdivisionDZ06:
		return DZ
	case SubdivisionDZ07:
		return DZ
	case SubdivisionDZ08:
		return DZ
	case SubdivisionDZ09:
		return DZ
	case SubdivisionDZ10:
		return DZ
	case SubdivisionDZ11:
		return DZ
	case SubdivisionDZ12:
		return DZ
	case SubdivisionDZ13:
		return DZ
	case SubdivisionDZ14:
		return DZ
	case SubdivisionDZ15:
		return DZ
	case SubdivisionDZ16:
		return DZ
	case SubdivisionDZ17:
		return DZ
	case SubdivisionDZ18:
		return DZ
	case SubdivisionDZ19:
		return DZ
	case SubdivisionDZ20:
		return DZ
	case SubdivisionDZ21:
		return DZ
	case SubdivisionDZ22:
		return DZ
	case SubdivisionDZ23:
		return DZ
	case SubdivisionDZ24:
		return DZ
	case SubdivisionDZ25:
		return DZ
	case SubdivisionDZ26:
		return DZ
	case SubdivisionDZ27:
		return DZ
	case SubdivisionDZ28:
		return DZ
	case SubdivisionDZ29:
		return DZ
	case SubdivisionDZ30:
		return DZ
	case SubdivisionDZ31:
		return DZ
	case SubdivisionDZ32:
		return DZ
	case SubdivisionDZ33:
		return DZ
	case SubdivisionDZ34:
		return DZ
	case SubdivisionDZ35:
		return DZ
	case SubdivisionDZ36:
		return DZ
	case SubdivisionDZ37:
		return DZ
	case SubdivisionDZ38:
		return DZ
	case SubdivisionDZ39:
		return DZ
	case SubdivisionDZ40:
		return DZ
	case SubdivisionDZ41:
		return DZ
	case SubdivisionDZ42:
		return DZ
	case SubdivisionDZ43:
		return DZ
	case SubdivisionDZ44:
		return DZ
	case SubdivisionDZ45:
		return DZ
	case SubdivisionDZ46:
		return DZ
	case SubdivisionDZ47:
		return DZ
	case SubdivisionDZ48:
		return DZ
	case SubdivisionECA:
		return EC
	case SubdivisionECB:
		return EC
	case SubdivisionECC:
		return EC
	case SubdivisionECD:
		return EC
	case SubdivisionECE:
		return EC
	case SubdivisionECF:
		return EC
	case SubdivisionECG:
		return EC
	case SubdivisionECH:
		return EC
	case SubdivisionECI:
		return EC
	case SubdivisionECL:
		return EC
	case SubdivisionECM:
		return EC
	case SubdivisionECN:
		return EC
	case SubdivisionECO:
		return EC
	case SubdivisionECP:
		return EC
	case SubdivisionECR:
		return EC
	case SubdivisionECS:
		return EC
	case SubdivisionECSD:
		return EC
	case SubdivisionECSE:
		return EC
	case SubdivisionECT:
		return EC
	case SubdivisionECU:
		return EC
	case SubdivisionECW:
		return EC
	case SubdivisionECX:
		return EC
	case SubdivisionECY:
		return EC
	case SubdivisionECZ:
		return EC
	case SubdivisionEE37:
		return EE
	case SubdivisionEE39:
		return EE
	case SubdivisionEE44:
		return EE
	case SubdivisionEE49:
		return EE
	case SubdivisionEE51:
		return EE
	case SubdivisionEE57:
		return EE
	case SubdivisionEE59:
		return EE
	case SubdivisionEE65:
		return EE
	case SubdivisionEE67:
		return EE
	case SubdivisionEE70:
		return EE
	case SubdivisionEE74:
		return EE
	case SubdivisionEE78:
		return EE
	case SubdivisionEE82:
		return EE
	case SubdivisionEE84:
		return EE
	case SubdivisionEE86:
		return EE
	case SubdivisionEGALX:
		return EG
	case SubdivisionEGASN:
		return EG
	case SubdivisionEGAST:
		return EG
	case SubdivisionEGBA:
		return EG
	case SubdivisionEGBH:
		return EG
	case SubdivisionEGBNS:
		return EG
	case SubdivisionEGC:
		return EG
	case SubdivisionEGDK:
		return EG
	case SubdivisionEGDT:
		return EG
	case SubdivisionEGFYM:
		return EG
	case SubdivisionEGGH:
		return EG
	case SubdivisionEGGZ:
		return EG
	case SubdivisionEGHU:
		return EG
	case SubdivisionEGIS:
		return EG
	case SubdivisionEGJS:
		return EG
	case SubdivisionEGKB:
		return EG
	case SubdivisionEGKFS:
		return EG
	case SubdivisionEGKN:
		return EG
	case SubdivisionEGMN:
		return EG
	case SubdivisionEGMNF:
		return EG
	case SubdivisionEGMT:
		return EG
	case SubdivisionEGPTS:
		return EG
	case SubdivisionEGSHG:
		return EG
	case SubdivisionEGSHR:
		return EG
	case SubdivisionEGSIN:
		return EG
	case SubdivisionEGSU:
		return EG
	case SubdivisionEGSUZ:
		return EG
	case SubdivisionEGWAD:
		return EG
	case SubdivisionERAN:
		return ER
	case SubdivisionERDK:
		return ER
	case SubdivisionERDU:
		return ER
	case SubdivisionERGB:
		return ER
	case SubdivisionERMA:
		return ER
	case SubdivisionERSK:
		return ER
	case SubdivisionESA:
		return ES
	case SubdivisionESAB:
		return ES
	case SubdivisionESAL:
		return ES
	case SubdivisionESAN:
		return ES
	case SubdivisionESAR:
		return ES
	case SubdivisionESAS:
		return ES
	case SubdivisionESAV:
		return ES
	case SubdivisionESB:
		return ES
	case SubdivisionESBA:
		return ES
	case SubdivisionESBI:
		return ES
	case SubdivisionESBU:
		return ES
	case SubdivisionESC:
		return ES
	case SubdivisionESCA:
		return ES
	case SubdivisionESCB:
		return ES
	case SubdivisionESCC:
		return ES
	case SubdivisionESCE:
		return ES
	case SubdivisionESCL:
		return ES
	case SubdivisionESCM:
		return ES
	case SubdivisionESCN:
		return ES
	case SubdivisionESCO:
		return ES
	case SubdivisionESCR:
		return ES
	case SubdivisionESCS:
		return ES
	case SubdivisionESCT:
		return ES
	case SubdivisionESCU:
		return ES
	case SubdivisionESEX:
		return ES
	case SubdivisionESGA:
		return ES
	case SubdivisionESGC:
		return ES
	case SubdivisionESGI:
		return ES
	case SubdivisionESGR:
		return ES
	case SubdivisionESGU:
		return ES
	case SubdivisionESH:
		return ES
	case SubdivisionESHU:
		return ES
	case SubdivisionESIB:
		return ES
	case SubdivisionESJ:
		return ES
	case SubdivisionESL:
		return ES
	case SubdivisionESLE:
		return ES
	case SubdivisionESLO:
		return ES
	case SubdivisionESLU:
		return ES
	case SubdivisionESM:
		return ES
	case SubdivisionESMA:
		return ES
	case SubdivisionESMC:
		return ES
	case SubdivisionESMD:
		return ES
	case SubdivisionESML:
		return ES
	case SubdivisionESMU:
		return ES
	case SubdivisionESNA:
		return ES
	case SubdivisionESNC:
		return ES
	case SubdivisionESO:
		return ES
	case SubdivisionESOR:
		return ES
	case SubdivisionESP:
		return ES
	case SubdivisionESPM:
		return ES
	case SubdivisionESPO:
		return ES
	case SubdivisionESPV:
		return ES
	case SubdivisionESRI:
		return ES
	case SubdivisionESS:
		return ES
	case SubdivisionESSA:
		return ES
	case SubdivisionESSE:
		return ES
	case SubdivisionESSG:
		return ES
	case SubdivisionESSO:
		return ES
	case SubdivisionESSS:
		return ES
	case SubdivisionEST:
		return ES
	case SubdivisionESTE:
		return ES
	case SubdivisionESTF:
		return ES
	case SubdivisionESTO:
		return ES
	case SubdivisionESV:
		return ES
	case SubdivisionESVA:
		return ES
	case SubdivisionESVC:
		return ES
	case SubdivisionESVI:
		return ES
	case SubdivisionESZ:
		return ES
	case SubdivisionESZA:
		return ES
	case SubdivisionETAA:
		return ET
	case SubdivisionETAF:
		return ET
	case SubdivisionETAM:
		return ET
	case SubdivisionETBE:
		return ET
	case SubdivisionETDD:
		return ET
	case SubdivisionETGA:
		return ET
	case SubdivisionETHA:
		return ET
	case SubdivisionETOR:
		return ET
	case SubdivisionETSN:
		return ET
	case SubdivisionETSO:
		return ET
	case SubdivisionETTI:
		return ET
	case SubdivisionFI01:
		return FI
	case SubdivisionFI02:
		return FI
	case SubdivisionFI03:
		return FI
	case SubdivisionFI04:
		return FI
	case SubdivisionFI05:
		return FI
	case SubdivisionFI06:
		return FI
	case SubdivisionFI07:
		return FI
	case SubdivisionFI08:
		return FI
	case SubdivisionFI09:
		return FI
	case SubdivisionFI10:
		return FI
	case SubdivisionFI11:
		return FI
	case SubdivisionFI12:
		return FI
	case SubdivisionFI13:
		return FI
	case SubdivisionFI14:
		return FI
	case SubdivisionFI15:
		return FI
	case SubdivisionFI16:
		return FI
	case SubdivisionFI17:
		return FI
	case SubdivisionFI18:
		return FI
	case SubdivisionFI19:
		return FI
	case SubdivisionFJC:
		return FJ
	case SubdivisionFJE:
		return FJ
	case SubdivisionFJN:
		return FJ
	case SubdivisionFJR:
		return FJ
	case SubdivisionFJW:
		return FJ
	case SubdivisionFMKSA:
		return FM
	case SubdivisionFMPNI:
		return FM
	case SubdivisionFMTRK:
		return FM
	case SubdivisionFMYAP:
		return FM
	case SubdivisionFR01:
		return FR
	case SubdivisionFR02:
		return FR
	case SubdivisionFR03:
		return FR
	case SubdivisionFR04:
		return FR
	case SubdivisionFR05:
		return FR
	case SubdivisionFR06:
		return FR
	case SubdivisionFR07:
		return FR
	case SubdivisionFR08:
		return FR
	case SubdivisionFR09:
		return FR
	case SubdivisionFR10:
		return FR
	case SubdivisionFR11:
		return FR
	case SubdivisionFR12:
		return FR
	case SubdivisionFR13:
		return FR
	case SubdivisionFR14:
		return FR
	case SubdivisionFR15:
		return FR
	case SubdivisionFR16:
		return FR
	case SubdivisionFR17:
		return FR
	case SubdivisionFR18:
		return FR
	case SubdivisionFR19:
		return FR
	case SubdivisionFR21:
		return FR
	case SubdivisionFR22:
		return FR
	case SubdivisionFR23:
		return FR
	case SubdivisionFR24:
		return FR
	case SubdivisionFR25:
		return FR
	case SubdivisionFR26:
		return FR
	case SubdivisionFR27:
		return FR
	case SubdivisionFR28:
		return FR
	case SubdivisionFR29:
		return FR
	case SubdivisionFR2A:
		return FR
	case SubdivisionFR2B:
		return FR
	case SubdivisionFR30:
		return FR
	case SubdivisionFR31:
		return FR
	case SubdivisionFR32:
		return FR
	case SubdivisionFR33:
		return FR
	case SubdivisionFR34:
		return FR
	case SubdivisionFR35:
		return FR
	case SubdivisionFR36:
		return FR
	case SubdivisionFR37:
		return FR
	case SubdivisionFR38:
		return FR
	case SubdivisionFR39:
		return FR
	case SubdivisionFR40:
		return FR
	case SubdivisionFR41:
		return FR
	case SubdivisionFR42:
		return FR
	case SubdivisionFR43:
		return FR
	case SubdivisionFR44:
		return FR
	case SubdivisionFR45:
		return FR
	case SubdivisionFR46:
		return FR
	case SubdivisionFR47:
		return FR
	case SubdivisionFR48:
		return FR
	case SubdivisionFR49:
		return FR
	case SubdivisionFR50:
		return FR
	case SubdivisionFR51:
		return FR
	case SubdivisionFR52:
		return FR
	case SubdivisionFR53:
		return FR
	case SubdivisionFR54:
		return FR
	case SubdivisionFR55:
		return FR
	case SubdivisionFR56:
		return FR
	case SubdivisionFR57:
		return FR
	case SubdivisionFR58:
		return FR
	case SubdivisionFR59:
		return FR
	case SubdivisionFR60:
		return FR
	case SubdivisionFR61:
		return FR
	case SubdivisionFR62:
		return FR
	case SubdivisionFR63:
		return FR
	case SubdivisionFR64:
		return FR
	case SubdivisionFR65:
		return FR
	case SubdivisionFR66:
		return FR
	case SubdivisionFR67:
		return FR
	case SubdivisionFR68:
		return FR
	case SubdivisionFR69:
		return FR
	case SubdivisionFR70:
		return FR
	case SubdivisionFR71:
		return FR
	case SubdivisionFR72:
		return FR
	case SubdivisionFR73:
		return FR
	case SubdivisionFR74:
		return FR
	case SubdivisionFR75:
		return FR
	case SubdivisionFR76:
		return FR
	case SubdivisionFR77:
		return FR
	case SubdivisionFR78:
		return FR
	case SubdivisionFR79:
		return FR
	case SubdivisionFR80:
		return FR
	case SubdivisionFR81:
		return FR
	case SubdivisionFR82:
		return FR
	case SubdivisionFR83:
		return FR
	case SubdivisionFR84:
		return FR
	case SubdivisionFR85:
		return FR
	case SubdivisionFR86:
		return FR
	case SubdivisionFR87:
		return FR
	case SubdivisionFR88:
		return FR
	case SubdivisionFR89:
		return FR
	case SubdivisionFR90:
		return FR
	case SubdivisionFR91:
		return FR
	case SubdivisionFR92:
		return FR
	case SubdivisionFR93:
		return FR
	case SubdivisionFR94:
		return FR
	case SubdivisionFR95:
		return FR
	case SubdivisionFRARA:
		return FR
	case SubdivisionFRBFC:
		return FR
	case SubdivisionFRBL:
		return FR
	case SubdivisionFRBRE:
		return FR
	case SubdivisionFRCOR:
		return FR
	case SubdivisionFRCP:
		return FR
	case SubdivisionFRCVL:
		return FR
	case SubdivisionFRGES:
		return FR
	case SubdivisionFRGF:
		return FR
	case SubdivisionFRGP:
		return FR
	case SubdivisionFRGUA:
		return FR
	case SubdivisionFRHDF:
		return FR
	case SubdivisionFRIDF:
		return FR
	case SubdivisionFRLRE:
		return FR
	case SubdivisionFRMAY:
		return FR
	case SubdivisionFRMF:
		return FR
	case SubdivisionFRMQ:
		return FR
	case SubdivisionFRNAQ:
		return FR
	case SubdivisionFRNC:
		return FR
	case SubdivisionFRNOR:
		return FR
	case SubdivisionFROCC:
		return FR
	case SubdivisionFRPAC:
		return FR
	case SubdivisionFRPDL:
		return FR
	case SubdivisionFRPF:
		return FR
	case SubdivisionFRPM:
		return FR
	case SubdivisionFRRE:
		return FR
	case SubdivisionFRTF:
		return FR
	case SubdivisionFRWF:
		return FR
	case SubdivisionFRYT:
		return FR
	case SubdivisionGA1:
		return GA
	case SubdivisionGA2:
		return GA
	case SubdivisionGA3:
		return GA
	case SubdivisionGA4:
		return GA
	case SubdivisionGA5:
		return GA
	case SubdivisionGA6:
		return GA
	case SubdivisionGA7:
		return GA
	case SubdivisionGA8:
		return GA
	case SubdivisionGA9:
		return GA
	case SubdivisionGBABC:
		return GB
	case SubdivisionGBABD:
		return GB
	case SubdivisionGBABE:
		return GB
	case SubdivisionGBAGB:
		return GB
	case SubdivisionGBAGY:
		return GB
	case SubdivisionGBAND:
		return GB
	case SubdivisionGBANN:
		return GB
	case SubdivisionGBANS:
		return GB
	case SubdivisionGBBAS:
		return GB
	case SubdivisionGBBBD:
		return GB
	case SubdivisionGBBDF:
		return GB
	case SubdivisionGBBDG:
		return GB
	case SubdivisionGBBEN:
		return GB
	case SubdivisionGBBEX:
		return GB
	case SubdivisionGBBFS:
		return GB
	case SubdivisionGBBGE:
		return GB
	case SubdivisionGBBGW:
		return GB
	case SubdivisionGBBIR:
		return GB
	case SubdivisionGBBKM:
		return GB
	case SubdivisionGBBMH:
		return GB
	case SubdivisionGBBNE:
		return GB
	case SubdivisionGBBNH:
		return GB
	case SubdivisionGBBNS:
		return GB
	case SubdivisionGBBOL:
		return GB
	case SubdivisionGBBPL:
		return GB
	case SubdivisionGBBRC:
		return GB
	case SubdivisionGBBRD:
		return GB
	case SubdivisionGBBRY:
		return GB
	case SubdivisionGBBST:
		return GB
	case SubdivisionGBBUR:
		return GB
	case SubdivisionGBCAM:
		return GB
	case SubdivisionGBCAY:
		return GB
	case SubdivisionGBCBF:
		return GB
	case SubdivisionGBCCG:
		return GB
	case SubdivisionGBCGN:
		return GB
	case SubdivisionGBCHE:
		return GB
	case SubdivisionGBCHW:
		return GB
	case SubdivisionGBCLD:
		return GB
	case SubdivisionGBCLK:
		return GB
	case SubdivisionGBCMA:
		return GB
	case SubdivisionGBCMD:
		return GB
	case SubdivisionGBCMN:
		return GB
	case SubdivisionGBCON:
		return GB
	case SubdivisionGBCOV:
		return GB
	case SubdivisionGBCRF:
		return GB
	case SubdivisionGBCRY:
		return GB
	case SubdivisionGBCWY:
		return GB
	case SubdivisionGBDAL:
		return GB
	case SubdivisionGBDBY:
		return GB
	case SubdivisionGBDEN:
		return GB
	case SubdivisionGBDER:
		return GB
	case SubdivisionGBDEV:
		return GB
	case SubdivisionGBDGY:
		return GB
	case SubdivisionGBDNC:
		return GB
	case SubdivisionGBDND:
		return GB
	case SubdivisionGBDOR:
		return GB
	case SubdivisionGBDRS:
		return GB
	case SubdivisionGBDUD:
		return GB
	case SubdivisionGBDUR:
		return GB
	case SubdivisionGBEAL:
		return GB
	case SubdivisionGBEAW:
		return GB
	case SubdivisionGBEAY:
		return GB
	case SubdivisionGBEDH:
		return GB
	case SubdivisionGBEDU:
		return GB
	case SubdivisionGBELN:
		return GB
	case SubdivisionGBELS:
		return GB
	case SubdivisionGBENF:
		return GB
	case SubdivisionGBENG:
		return GB
	case SubdivisionGBERW:
		return GB
	case SubdivisionGBERY:
		return GB
	case SubdivisionGBESS:
		return GB
	case SubdivisionGBESX:
		return GB
	case SubdivisionGBFAL:
		return GB
	case SubdivisionGBFIF:
		return GB
	case SubdivisionGBFLN:
		return GB
	case SubdivisionGBFMO:
		return GB
	case SubdivisionGBGAT:
		return GB
	case SubdivisionGBGBN:
		return GB
	case SubdivisionGBGLG:
		return GB
	case SubdivisionGBGLS:
		return GB
	case SubdivisionGBGRE:
		return GB
	case SubdivisionGBGWN:
		return GB
	case SubdivisionGBHAL:
		return GB
	case SubdivisionGBHAM:
		return GB
	case SubdivisionGBHAV:
		return GB
	case SubdivisionGBHCK:
		return GB
	case SubdivisionGBHEF:
		return GB
	case SubdivisionGBHIL:
		return GB
	case SubdivisionGBHLD:
		return GB
	case SubdivisionGBHMF:
		return GB
	case SubdivisionGBHNS:
		return GB
	case SubdivisionGBHPL:
		return GB
	case SubdivisionGBHRT:
		return GB
	case SubdivisionGBHRW:
		return GB
	case SubdivisionGBHRY:
		return GB
	case SubdivisionGBIOS:
		return GB
	case SubdivisionGBIOW:
		return GB
	case SubdivisionGBISL:
		return GB
	case SubdivisionGBIVC:
		return GB
	case SubdivisionGBKEC:
		return GB
	case SubdivisionGBKEN:
		return GB
	case SubdivisionGBKHL:
		return GB
	case SubdivisionGBKIR:
		return GB
	case SubdivisionGBKTT:
		return GB
	case SubdivisionGBKWL:
		return GB
	case SubdivisionGBLAN:
		return GB
	case SubdivisionGBLBC:
		return GB
	case SubdivisionGBLBH:
		return GB
	case SubdivisionGBLCE:
		return GB
	case SubdivisionGBLDS:
		return GB
	case SubdivisionGBLEC:
		return GB
	case SubdivisionGBLEW:
		return GB
	case SubdivisionGBLIN:
		return GB
	case SubdivisionGBLIV:
		return GB
	case SubdivisionGBLND:
		return GB
	case SubdivisionGBLUT:
		return GB
	case SubdivisionGBMAN:
		return GB
	case SubdivisionGBMDB:
		return GB
	case SubdivisionGBMDW:
		return GB
	case SubdivisionGBMEA:
		return GB
	case SubdivisionGBMIK:
		return GB
	case SubdivisionGBMLN:
		return GB
	case SubdivisionGBMON:
		return GB
	case SubdivisionGBMRT:
		return GB
	case SubdivisionGBMRY:
		return GB
	case SubdivisionGBMTY:
		return GB
	case SubdivisionGBMUL:
		return GB
	case SubdivisionGBNAY:
		return GB
	case SubdivisionGBNBL:
		return GB
	case SubdivisionGBNEL:
		return GB
	case SubdivisionGBNET:
		return GB
	case SubdivisionGBNFK:
		return GB
	case SubdivisionGBNGM:
		return GB
	case SubdivisionGBNIR:
		return GB
	case SubdivisionGBNLK:
		return GB
	case SubdivisionGBNLN:
		return GB
	case SubdivisionGBNMD:
		return GB
	case SubdivisionGBNSM:
		return GB
	case SubdivisionGBNTH:
		return GB
	case SubdivisionGBNTL:
		return GB
	case SubdivisionGBNTT:
		return GB
	case SubdivisionGBNTY:
		return GB
	case SubdivisionGBNWM:
		return GB
	case SubdivisionGBNWP:
		return GB
	case SubdivisionGBNYK:
		return GB
	case SubdivisionGBOLD:
		return GB
	case SubdivisionGBORK:
		return GB
	case SubdivisionGBOXF:
		return GB
	case SubdivisionGBPEM:
		return GB
	case SubdivisionGBPKN:
		return GB
	case SubdivisionGBPLY:
		return GB
	case SubdivisionGBPOL:
		return GB
	case SubdivisionGBPOR:
		return GB
	case SubdivisionGBPOW:
		return GB
	case SubdivisionGBPTE:
		return GB
	case SubdivisionGBRCC:
		return GB
	case SubdivisionGBRCH:
		return GB
	case SubdivisionGBRCT:
		return GB
	case SubdivisionGBRDB:
		return GB
	case SubdivisionGBRDG:
		return GB
	case SubdivisionGBRFW:
		return GB
	case SubdivisionGBRIC:
		return GB
	case SubdivisionGBROT:
		return GB
	case SubdivisionGBRUT:
		return GB
	case SubdivisionGBSAW:
		return GB
	case SubdivisionGBSAY:
		return GB
	case SubdivisionGBSCB:
		return GB
	case SubdivisionGBSCT:
		return GB
	case SubdivisionGBSFK:
		return GB
	case SubdivisionGBSFT:
		return GB
	case SubdivisionGBSGC:
		return GB
	case SubdivisionGBSHF:
		return GB
	case SubdivisionGBSHN:
		return GB
	case SubdivisionGBSHR:
		return GB
	case SubdivisionGBSKP:
		return GB
	case SubdivisionGBSLF:
		return GB
	case SubdivisionGBSLG:
		return GB
	case SubdivisionGBSLK:
		return GB
	case SubdivisionGBSND:
		return GB
	case SubdivisionGBSOL:
		return GB
	case SubdivisionGBSOM:
		return GB
	case SubdivisionGBSOS:
		return GB
	case SubdivisionGBSRY:
		return GB
	case SubdivisionGBSTE:
		return GB
	case SubdivisionGBSTG:
		return GB
	case SubdivisionGBSTH:
		return GB
	case SubdivisionGBSTN:
		return GB
	case SubdivisionGBSTS:
		return GB
	case SubdivisionGBSTT:
		return GB
	case SubdivisionGBSTY:
		return GB
	case SubdivisionGBSWA:
		return GB
	case SubdivisionGBSWD:
		return GB
	case SubdivisionGBSWK:
		return GB
	case SubdivisionGBTAM:
		return GB
	case SubdivisionGBTFW:
		return GB
	case SubdivisionGBTHR:
		return GB
	case SubdivisionGBTOB:
		return GB
	case SubdivisionGBTOF:
		return GB
	case SubdivisionGBTRF:
		return GB
	case SubdivisionGBTWH:
		return GB
	case SubdivisionGBUKM:
		return GB
	case SubdivisionGBVGL:
		return GB
	case SubdivisionGBWAR:
		return GB
	case SubdivisionGBWBK:
		return GB
	case SubdivisionGBWDU:
		return GB
	case SubdivisionGBWFT:
		return GB
	case SubdivisionGBWGN:
		return GB
	case SubdivisionGBWIL:
		return GB
	case SubdivisionGBWKF:
		return GB
	case SubdivisionGBWLL:
		return GB
	case SubdivisionGBWLN:
		return GB
	case SubdivisionGBWLS:
		return GB
	case SubdivisionGBWLV:
		return GB
	case SubdivisionGBWND:
		return GB
	case SubdivisionGBWNM:
		return GB
	case SubdivisionGBWOK:
		return GB
	case SubdivisionGBWOR:
		return GB
	case SubdivisionGBWRL:
		return GB
	case SubdivisionGBWRT:
		return GB
	case SubdivisionGBWRX:
		return GB
	case SubdivisionGBWSM:
		return GB
	case SubdivisionGBWSX:
		return GB
	case SubdivisionGBYOR:
		return GB
	case SubdivisionGBZET:
		return GB
	case SubdivisionGD01:
		return GD
	case SubdivisionGD02:
		return GD
	case SubdivisionGD03:
		return GD
	case SubdivisionGD04:
		return GD
	case SubdivisionGD05:
		return GD
	case SubdivisionGD06:
		return GD
	case SubdivisionGD10:
		return GD
	case SubdivisionGEAB:
		return GE
	case SubdivisionGEAJ:
		return GE
	case SubdivisionGEGU:
		return GE
	case SubdivisionGEIM:
		return GE
	case SubdivisionGEKA:
		return GE
	case SubdivisionGEKK:
		return GE
	case SubdivisionGEMM:
		return GE
	case SubdivisionGERL:
		return GE
	case SubdivisionGESJ:
		return GE
	case SubdivisionGESK:
		return GE
	case SubdivisionGESZ:
		return GE
	case SubdivisionGETB:
		return GE
	case SubdivisionGHAA:
		return GH
	case SubdivisionGHAH:
		return GH
	case SubdivisionGHBA:
		return GH
	case SubdivisionGHCP:
		return GH
	case SubdivisionGHEP:
		return GH
	case SubdivisionGHNP:
		return GH
	case SubdivisionGHTV:
		return GH
	case SubdivisionGHUE:
		return GH
	case SubdivisionGHUW:
		return GH
	case SubdivisionGHWP:
		return GH
	case SubdivisionGLKU:
		return GL
	case SubdivisionGLQA:
		return GL
	case SubdivisionGLQE:
		return GL
	case SubdivisionGLSM:
		return GL
	case SubdivisionGMB:
		return GM
	case SubdivisionGML:
		return GM
	case SubdivisionGMM:
		return GM
	case SubdivisionGMN:
		return GM
	case SubdivisionGMU:
		return GM
	case SubdivisionGMW:
		return GM
	case SubdivisionGNB:
		return GN
	case SubdivisionGNBE:
		return GN
	case SubdivisionGNBF:
		return GN
	case SubdivisionGNBK:
		return GN
	case SubdivisionGNC:
		return GN
	case SubdivisionGNCO:
		return GN
	case SubdivisionGND:
		return GN
	case SubdivisionGNDB:
		return GN
	case SubdivisionGNDI:
		return GN
	case SubdivisionGNDL:
		return GN
	case SubdivisionGNDU:
		return GN
	case SubdivisionGNF:
		return GN
	case SubdivisionGNFA:
		return GN
	case SubdivisionGNFO:
		return GN
	case SubdivisionGNFR:
		return GN
	case SubdivisionGNGA:
		return GN
	case SubdivisionGNGU:
		return GN
	case SubdivisionGNK:
		return GN
	case SubdivisionGNKA:
		return GN
	case SubdivisionGNKB:
		return GN
	case SubdivisionGNKD:
		return GN
	case SubdivisionGNKE:
		return GN
	case SubdivisionGNKN:
		return GN
	case SubdivisionGNKO:
		return GN
	case SubdivisionGNKS:
		return GN
	case SubdivisionGNL:
		return GN
	case SubdivisionGNLA:
		return GN
	case SubdivisionGNLE:
		return GN
	case SubdivisionGNLO:
		return GN
	case SubdivisionGNM:
		return GN
	case SubdivisionGNMC:
		return GN
	case SubdivisionGNMD:
		return GN
	case SubdivisionGNML:
		return GN
	case SubdivisionGNMM:
		return GN
	case SubdivisionGNN:
		return GN
	case SubdivisionGNNZ:
		return GN
	case SubdivisionGNPI:
		return GN
	case SubdivisionGNSI:
		return GN
	case SubdivisionGNTE:
		return GN
	case SubdivisionGNTO:
		return GN
	case SubdivisionGNYO:
		return GN
	case SubdivisionGQAN:
		return GQ
	case SubdivisionGQBN:
		return GQ
	case SubdivisionGQBS:
		return GQ
	case SubdivisionGQC:
		return GQ
	case SubdivisionGQCS:
		return GQ
	case SubdivisionGQI:
		return GQ
	case SubdivisionGQKN:
		return GQ
	case SubdivisionGQLI:
		return GQ
	case SubdivisionGQWN:
		return GQ
	case SubdivisionGR01:
		return GR
	case SubdivisionGR03:
		return GR
	case SubdivisionGR04:
		return GR
	case SubdivisionGR05:
		return GR
	case SubdivisionGR06:
		return GR
	case SubdivisionGR07:
		return GR
	case SubdivisionGR11:
		return GR
	case SubdivisionGR12:
		return GR
	case SubdivisionGR13:
		return GR
	case SubdivisionGR14:
		return GR
	case SubdivisionGR15:
		return GR
	case SubdivisionGR16:
		return GR
	case SubdivisionGR17:
		return GR
	case SubdivisionGR21:
		return GR
	case SubdivisionGR22:
		return GR
	case SubdivisionGR23:
		return GR
	case SubdivisionGR24:
		return GR
	case SubdivisionGR31:
		return GR
	case SubdivisionGR32:
		return GR
	case SubdivisionGR33:
		return GR
	case SubdivisionGR34:
		return GR
	case SubdivisionGR41:
		return GR
	case SubdivisionGR42:
		return GR
	case SubdivisionGR43:
		return GR
	case SubdivisionGR44:
		return GR
	case SubdivisionGR51:
		return GR
	case SubdivisionGR52:
		return GR
	case SubdivisionGR53:
		return GR
	case SubdivisionGR54:
		return GR
	case SubdivisionGR55:
		return GR
	case SubdivisionGR56:
		return GR
	case SubdivisionGR57:
		return GR
	case SubdivisionGR58:
		return GR
	case SubdivisionGR59:
		return GR
	case SubdivisionGR61:
		return GR
	case SubdivisionGR62:
		return GR
	case SubdivisionGR63:
		return GR
	case SubdivisionGR64:
		return GR
	case SubdivisionGR69:
		return GR
	case SubdivisionGR71:
		return GR
	case SubdivisionGR72:
		return GR
	case SubdivisionGR73:
		return GR
	case SubdivisionGR81:
		return GR
	case SubdivisionGR82:
		return GR
	case SubdivisionGR83:
		return GR
	case SubdivisionGR84:
		return GR
	case SubdivisionGR85:
		return GR
	case SubdivisionGR91:
		return GR
	case SubdivisionGR92:
		return GR
	case SubdivisionGR93:
		return GR
	case SubdivisionGR94:
		return GR
	case SubdivisionGRA:
		return GR
	case SubdivisionGRA1:
		return GR
	case SubdivisionGRB:
		return GR
	case SubdivisionGRC:
		return GR
	case SubdivisionGRD:
		return GR
	case SubdivisionGRE:
		return GR
	case SubdivisionGRF:
		return GR
	case SubdivisionGRG:
		return GR
	case SubdivisionGRH:
		return GR
	case SubdivisionGRI:
		return GR
	case SubdivisionGRJ:
		return GR
	case SubdivisionGRK:
		return GR
	case SubdivisionGRL:
		return GR
	case SubdivisionGRM:
		return GR
	case SubdivisionGTAV:
		return GT
	case SubdivisionGTBV:
		return GT
	case SubdivisionGTCM:
		return GT
	case SubdivisionGTCQ:
		return GT
	case SubdivisionGTES:
		return GT
	case SubdivisionGTGU:
		return GT
	case SubdivisionGTHU:
		return GT
	case SubdivisionGTIZ:
		return GT
	case SubdivisionGTJA:
		return GT
	case SubdivisionGTJU:
		return GT
	case SubdivisionGTPE:
		return GT
	case SubdivisionGTPR:
		return GT
	case SubdivisionGTQC:
		return GT
	case SubdivisionGTQZ:
		return GT
	case SubdivisionGTRE:
		return GT
	case SubdivisionGTSA:
		return GT
	case SubdivisionGTSM:
		return GT
	case SubdivisionGTSO:
		return GT
	case SubdivisionGTSR:
		return GT
	case SubdivisionGTSU:
		return GT
	case SubdivisionGTTO:
		return GT
	case SubdivisionGTZA:
		return GT
	case SubdivisionGWBA:
		return GW
	case SubdivisionGWBL:
		return GW
	case SubdivisionGWBM:
		return GW
	case SubdivisionGWBS:
		return GW
	case SubdivisionGWCA:
		return GW
	case SubdivisionGWGA:
		return GW
	case SubdivisionGWL:
		return GW
	case SubdivisionGWN:
		return GW
	case SubdivisionGWOI:
		return GW
	case SubdivisionGWQU:
		return GW
	case SubdivisionGWS:
		return GW
	case SubdivisionGWTO:
		return GW
	case SubdivisionGYBA:
		return GY
	case SubdivisionGYCU:
		return GY
	case SubdivisionGYDE:
		return GY
	case SubdivisionGYEB:
		return GY
	case SubdivisionGYES:
		return GY
	case SubdivisionGYMA:
		return GY
	case SubdivisionGYPM:
		return GY
	case SubdivisionGYPT:
		return GY
	case SubdivisionGYUD:
		return GY
	case SubdivisionGYUT:
		return GY
	case SubdivisionHNAT:
		return HN
	case SubdivisionHNCH:
		return HN
	case SubdivisionHNCL:
		return HN
	case SubdivisionHNCM:
		return HN
	case SubdivisionHNCP:
		return HN
	case SubdivisionHNCR:
		return HN
	case SubdivisionHNEP:
		return HN
	case SubdivisionHNFM:
		return HN
	case SubdivisionHNGD:
		return HN
	case SubdivisionHNIB:
		return HN
	case SubdivisionHNIN:
		return HN
	case SubdivisionHNLE:
		return HN
	case SubdivisionHNLP:
		return HN
	case SubdivisionHNOC:
		return HN
	case SubdivisionHNOL:
		return HN
	case SubdivisionHNSB:
		return HN
	case SubdivisionHNVA:
		return HN
	case SubdivisionHNYO:
		return HN
	case SubdivisionHR01:
		return HR
	case SubdivisionHR02:
		return HR
	case SubdivisionHR03:
		return HR
	case SubdivisionHR04:
		return HR
	case SubdivisionHR05:
		return HR
	case SubdivisionHR06:
		return HR
	case SubdivisionHR07:
		return HR
	case SubdivisionHR08:
		return HR
	case SubdivisionHR09:
		return HR
	case SubdivisionHR10:
		return HR
	case SubdivisionHR11:
		return HR
	case SubdivisionHR12:
		return HR
	case SubdivisionHR13:
		return HR
	case SubdivisionHR14:
		return HR
	case SubdivisionHR15:
		return HR
	case SubdivisionHR16:
		return HR
	case SubdivisionHR17:
		return HR
	case SubdivisionHR18:
		return HR
	case SubdivisionHR19:
		return HR
	case SubdivisionHR20:
		return HR
	case SubdivisionHR21:
		return HR
	case SubdivisionHTAR:
		return HT
	case SubdivisionHTCE:
		return HT
	case SubdivisionHTGA:
		return HT
	case SubdivisionHTND:
		return HT
	case SubdivisionHTNE:
		return HT
	case SubdivisionHTNO:
		return HT
	case SubdivisionHTOU:
		return HT
	case SubdivisionHTSD:
		return HT
	case SubdivisionHTSE:
		return HT
	case SubdivisionHUBA:
		return HU
	case SubdivisionHUBC:
		return HU
	case SubdivisionHUBE:
		return HU
	case SubdivisionHUBK:
		return HU
	case SubdivisionHUBU:
		return HU
	case SubdivisionHUBZ:
		return HU
	case SubdivisionHUCS:
		return HU
	case SubdivisionHUDE:
		return HU
	case SubdivisionHUDU:
		return HU
	case SubdivisionHUEG:
		return HU
	case SubdivisionHUER:
		return HU
	case SubdivisionHUFE:
		return HU
	case SubdivisionHUGS:
		return HU
	case SubdivisionHUGY:
		return HU
	case SubdivisionHUHB:
		return HU
	case SubdivisionHUHE:
		return HU
	case SubdivisionHUHV:
		return HU
	case SubdivisionHUJN:
		return HU
	case SubdivisionHUKE:
		return HU
	case SubdivisionHUKM:
		return HU
	case SubdivisionHUKV:
		return HU
	case SubdivisionHUMI:
		return HU
	case SubdivisionHUNK:
		return HU
	case SubdivisionHUNO:
		return HU
	case SubdivisionHUNY:
		return HU
	case SubdivisionHUPE:
		return HU
	case SubdivisionHUPS:
		return HU
	case SubdivisionHUSD:
		return HU
	case SubdivisionHUSF:
		return HU
	case SubdivisionHUSH:
		return HU
	case SubdivisionHUSK:
		return HU
	case SubdivisionHUSN:
		return HU
	case SubdivisionHUSO:
		return HU
	case SubdivisionHUSS:
		return HU
	case SubdivisionHUST:
		return HU
	case SubdivisionHUSZ:
		return HU
	case SubdivisionHUTB:
		return HU
	case SubdivisionHUTO:
		return HU
	case SubdivisionHUVA:
		return HU
	case SubdivisionHUVE:
		return HU
	case SubdivisionHUVM:
		return HU
	case SubdivisionHUZA:
		return HU
	case SubdivisionHUZE:
		return HU
	case SubdivisionIDAC:
		return ID
	case SubdivisionIDBA:
		return ID
	case SubdivisionIDBB:
		return ID
	case SubdivisionIDBE:
		return ID
	case SubdivisionIDBT:
		return ID
	case SubdivisionIDGO:
		return ID
	case SubdivisionIDIJ:
		return ID
	case SubdivisionIDJA:
		return ID
	case SubdivisionIDJB:
		return ID
	case SubdivisionIDJI:
		return ID
	case SubdivisionIDJK:
		return ID
	case SubdivisionIDJT:
		return ID
	case SubdivisionIDJW:
		return ID
	case SubdivisionIDKA:
		return ID
	case SubdivisionIDKB:
		return ID
	case SubdivisionIDKI:
		return ID
	case SubdivisionIDKR:
		return ID
	case SubdivisionIDKS:
		return ID
	case SubdivisionIDKT:
		return ID
	case SubdivisionIDLA:
		return ID
	case SubdivisionIDMA:
		return ID
	case SubdivisionIDML:
		return ID
	case SubdivisionIDMU:
		return ID
	case SubdivisionIDNB:
		return ID
	case SubdivisionIDNT:
		return ID
	case SubdivisionIDNU:
		return ID
	case SubdivisionIDPA:
		return ID
	case SubdivisionIDPB:
		return ID
	case SubdivisionIDRI:
		return ID
	case SubdivisionIDSA:
		return ID
	case SubdivisionIDSB:
		return ID
	case SubdivisionIDSG:
		return ID
	case SubdivisionIDSL:
		return ID
	case SubdivisionIDSM:
		return ID
	case SubdivisionIDSN:
		return ID
	case SubdivisionIDSR:
		return ID
	case SubdivisionIDSS:
		return ID
	case SubdivisionIDST:
		return ID
	case SubdivisionIDSU:
		return ID
	case SubdivisionIDYO:
		return ID
	case SubdivisionIEC:
		return IE
	case SubdivisionIECE:
		return IE
	case SubdivisionIECN:
		return IE
	case SubdivisionIECO:
		return IE
	case SubdivisionIECW:
		return IE
	case SubdivisionIED:
		return IE
	case SubdivisionIEDL:
		return IE
	case SubdivisionIEG:
		return IE
	case SubdivisionIEKE:
		return IE
	case SubdivisionIEKK:
		return IE
	case SubdivisionIEKY:
		return IE
	case SubdivisionIEL:
		return IE
	case SubdivisionIELD:
		return IE
	case SubdivisionIELH:
		return IE
	case SubdivisionIELK:
		return IE
	case SubdivisionIELM:
		return IE
	case SubdivisionIELS:
		return IE
	case SubdivisionIEM:
		return IE
	case SubdivisionIEMH:
		return IE
	case SubdivisionIEMN:
		return IE
	case SubdivisionIEMO:
		return IE
	case SubdivisionIEOY:
		return IE
	case SubdivisionIERN:
		return IE
	case SubdivisionIESO:
		return IE
	case SubdivisionIETA:
		return IE
	case SubdivisionIEU:
		return IE
	case SubdivisionIEWD:
		return IE
	case SubdivisionIEWH:
		return IE
	case SubdivisionIEWW:
		return IE
	case SubdivisionIEWX:
		return IE
	case SubdivisionILD:
		return IL
	case SubdivisionILHA:
		return IL
	case SubdivisionILJM:
		return IL
	case SubdivisionILM:
		return IL
	case SubdivisionILTA:
		return IL
	case SubdivisionILZ:
		return IL
	case SubdivisionINAN:
		return IN
	case SubdivisionINAP:
		return IN
	case SubdivisionINAR:
		return IN
	case SubdivisionINAS:
		return IN
	case SubdivisionINBR:
		return IN
	case SubdivisionINCH:
		return IN
	case SubdivisionINCT:
		return IN
	case SubdivisionINDD:
		return IN
	case SubdivisionINDL:
		return IN
	case SubdivisionINDN:
		return IN
	case SubdivisionINGA:
		return IN
	case SubdivisionINGJ:
		return IN
	case SubdivisionINHP:
		return IN
	case SubdivisionINHR:
		return IN
	case SubdivisionINJH:
		return IN
	case SubdivisionINJK:
		return IN
	case SubdivisionINKA:
		return IN
	case SubdivisionINKL:
		return IN
	case SubdivisionINLD:
		return IN
	case SubdivisionINMH:
		return IN
	case SubdivisionINML:
		return IN
	case SubdivisionINMN:
		return IN
	case SubdivisionINMP:
		return IN
	case SubdivisionINMZ:
		return IN
	case SubdivisionINNL:
		return IN
	case SubdivisionINOR:
		return IN
	case SubdivisionINPB:
		return IN
	case SubdivisionINPY:
		return IN
	case SubdivisionINRJ:
		return IN
	case SubdivisionINSK:
		return IN
	case SubdivisionINTG:
		return IN
	case SubdivisionINTN:
		return IN
	case SubdivisionINTR:
		return IN
	case SubdivisionINUP:
		return IN
	case SubdivisionINUT:
		return IN
	case SubdivisionINWB:
		return IN
	case SubdivisionIQAN:
		return IQ
	case SubdivisionIQAR:
		return IQ
	case SubdivisionIQBA:
		return IQ
	case SubdivisionIQBB:
		return IQ
	case SubdivisionIQBG:
		return IQ
	case SubdivisionIQDA:
		return IQ
	case SubdivisionIQDI:
		return IQ
	case SubdivisionIQDQ:
		return IQ
	case SubdivisionIQKA:
		return IQ
	case SubdivisionIQMA:
		return IQ
	case SubdivisionIQMU:
		return IQ
	case SubdivisionIQNA:
		return IQ
	case SubdivisionIQNI:
		return IQ
	case SubdivisionIQQA:
		return IQ
	case SubdivisionIQSD:
		return IQ
	case SubdivisionIQSW:
		return IQ
	case SubdivisionIQTS:
		return IQ
	case SubdivisionIQWA:
		return IQ
	case SubdivisionIR01:
		return IR
	case SubdivisionIR02:
		return IR
	case SubdivisionIR03:
		return IR
	case SubdivisionIR04:
		return IR
	case SubdivisionIR05:
		return IR
	case SubdivisionIR06:
		return IR
	case SubdivisionIR07:
		return IR
	case SubdivisionIR08:
		return IR
	case SubdivisionIR10:
		return IR
	case SubdivisionIR11:
		return IR
	case SubdivisionIR12:
		return IR
	case SubdivisionIR13:
		return IR
	case SubdivisionIR14:
		return IR
	case SubdivisionIR15:
		return IR
	case SubdivisionIR16:
		return IR
	case SubdivisionIR17:
		return IR
	case SubdivisionIR18:
		return IR
	case SubdivisionIR19:
		return IR
	case SubdivisionIR20:
		return IR
	case SubdivisionIR21:
		return IR
	case SubdivisionIR22:
		return IR
	case SubdivisionIR23:
		return IR
	case SubdivisionIR24:
		return IR
	case SubdivisionIR25:
		return IR
	case SubdivisionIR26:
		return IR
	case SubdivisionIR27:
		return IR
	case SubdivisionIR28:
		return IR
	case SubdivisionIR29:
		return IR
	case SubdivisionIR30:
		return IR
	case SubdivisionIR31:
		return IR
	case SubdivisionIS0:
		return IS
	case SubdivisionIS1:
		return IS
	case SubdivisionIS2:
		return IS
	case SubdivisionIS3:
		return IS
	case SubdivisionIS4:
		return IS
	case SubdivisionIS5:
		return IS
	case SubdivisionIS6:
		return IS
	case SubdivisionIS7:
		return IS
	case SubdivisionIS8:
		return IS
	case SubdivisionIT21:
		return IT
	case SubdivisionIT23:
		return IT
	case SubdivisionIT25:
		return IT
	case SubdivisionIT32:
		return IT
	case SubdivisionIT34:
		return IT
	case SubdivisionIT36:
		return IT
	case SubdivisionIT42:
		return IT
	case SubdivisionIT45:
		return IT
	case SubdivisionIT52:
		return IT
	case SubdivisionIT55:
		return IT
	case SubdivisionIT57:
		return IT
	case SubdivisionIT62:
		return IT
	case SubdivisionIT65:
		return IT
	case SubdivisionIT67:
		return IT
	case SubdivisionIT72:
		return IT
	case SubdivisionIT75:
		return IT
	case SubdivisionIT77:
		return IT
	case SubdivisionIT78:
		return IT
	case SubdivisionIT82:
		return IT
	case SubdivisionIT88:
		return IT
	case SubdivisionITAG:
		return IT
	case SubdivisionITAL:
		return IT
	case SubdivisionITAN:
		return IT
	case SubdivisionITAO:
		return IT
	case SubdivisionITAP:
		return IT
	case SubdivisionITAQ:
		return IT
	case SubdivisionITAR:
		return IT
	case SubdivisionITAT:
		return IT
	case SubdivisionITAV:
		return IT
	case SubdivisionITBA:
		return IT
	case SubdivisionITBG:
		return IT
	case SubdivisionITBI:
		return IT
	case SubdivisionITBL:
		return IT
	case SubdivisionITBN:
		return IT
	case SubdivisionITBO:
		return IT
	case SubdivisionITBR:
		return IT
	case SubdivisionITBS:
		return IT
	case SubdivisionITBT:
		return IT
	case SubdivisionITBZ:
		return IT
	case SubdivisionITCA:
		return IT
	case SubdivisionITCB:
		return IT
	case SubdivisionITCE:
		return IT
	case SubdivisionITCH:
		return IT
	case SubdivisionITCI:
		return IT
	case SubdivisionITCL:
		return IT
	case SubdivisionITCN:
		return IT
	case SubdivisionITCO:
		return IT
	case SubdivisionITCR:
		return IT
	case SubdivisionITCS:
		return IT
	case SubdivisionITCT:
		return IT
	case SubdivisionITCZ:
		return IT
	case SubdivisionITEN:
		return IT
	case SubdivisionITFC:
		return IT
	case SubdivisionITFE:
		return IT
	case SubdivisionITFG:
		return IT
	case SubdivisionITFI:
		return IT
	case SubdivisionITFM:
		return IT
	case SubdivisionITFR:
		return IT
	case SubdivisionITGE:
		return IT
	case SubdivisionITGO:
		return IT
	case SubdivisionITGR:
		return IT
	case SubdivisionITIM:
		return IT
	case SubdivisionITIS:
		return IT
	case SubdivisionITKR:
		return IT
	case SubdivisionITLC:
		return IT
	case SubdivisionITLE:
		return IT
	case SubdivisionITLI:
		return IT
	case SubdivisionITLO:
		return IT
	case SubdivisionITLT:
		return IT
	case SubdivisionITLU:
		return IT
	case SubdivisionITMB:
		return IT
	case SubdivisionITMC:
		return IT
	case SubdivisionITME:
		return IT
	case SubdivisionITMI:
		return IT
	case SubdivisionITMN:
		return IT
	case SubdivisionITMO:
		return IT
	case SubdivisionITMS:
		return IT
	case SubdivisionITMT:
		return IT
	case SubdivisionITNA:
		return IT
	case SubdivisionITNO:
		return IT
	case SubdivisionITNU:
		return IT
	case SubdivisionITOG:
		return IT
	case SubdivisionITOR:
		return IT
	case SubdivisionITOT:
		return IT
	case SubdivisionITPA:
		return IT
	case SubdivisionITPC:
		return IT
	case SubdivisionITPD:
		return IT
	case SubdivisionITPE:
		return IT
	case SubdivisionITPG:
		return IT
	case SubdivisionITPI:
		return IT
	case SubdivisionITPN:
		return IT
	case SubdivisionITPO:
		return IT
	case SubdivisionITPR:
		return IT
	case SubdivisionITPT:
		return IT
	case SubdivisionITPU:
		return IT
	case SubdivisionITPV:
		return IT
	case SubdivisionITPZ:
		return IT
	case SubdivisionITRA:
		return IT
	case SubdivisionITRC:
		return IT
	case SubdivisionITRE:
		return IT
	case SubdivisionITRG:
		return IT
	case SubdivisionITRI:
		return IT
	case SubdivisionITRM:
		return IT
	case SubdivisionITRN:
		return IT
	case SubdivisionITRO:
		return IT
	case SubdivisionITSA:
		return IT
	case SubdivisionITSI:
		return IT
	case SubdivisionITSO:
		return IT
	case SubdivisionITSP:
		return IT
	case SubdivisionITSR:
		return IT
	case SubdivisionITSS:
		return IT
	case SubdivisionITSV:
		return IT
	case SubdivisionITTA:
		return IT
	case SubdivisionITTE:
		return IT
	case SubdivisionITTN:
		return IT
	case SubdivisionITTO:
		return IT
	case SubdivisionITTP:
		return IT
	case SubdivisionITTR:
		return IT
	case SubdivisionITTS:
		return IT
	case SubdivisionITTV:
		return IT
	case SubdivisionITUD:
		return IT
	case SubdivisionITVA:
		return IT
	case SubdivisionITVB:
		return IT
	case SubdivisionITVC:
		return IT
	case SubdivisionITVE:
		return IT
	case SubdivisionITVI:
		return IT
	case SubdivisionITVR:
		return IT
	case SubdivisionITVS:
		return IT
	case SubdivisionITVT:
		return IT
	case SubdivisionITVV:
		return IT
	case SubdivisionJM01:
		return JM
	case SubdivisionJM02:
		return JM
	case SubdivisionJM03:
		return JM
	case SubdivisionJM04:
		return JM
	case SubdivisionJM05:
		return JM
	case SubdivisionJM06:
		return JM
	case SubdivisionJM07:
		return JM
	case SubdivisionJM08:
		return JM
	case SubdivisionJM09:
		return JM
	case SubdivisionJM10:
		return JM
	case SubdivisionJM11:
		return JM
	case SubdivisionJM12:
		return JM
	case SubdivisionJM13:
		return JM
	case SubdivisionJM14:
		return JM
	case SubdivisionJOAJ:
		return JO
	case SubdivisionJOAM:
		return JO
	case SubdivisionJOAQ:
		return JO
	case SubdivisionJOAT:
		return JO
	case SubdivisionJOAZ:
		return JO
	case SubdivisionJOBA:
		return JO
	case SubdivisionJOIR:
		return JO
	case SubdivisionJOJA:
		return JO
	case SubdivisionJOKA:
		return JO
	case SubdivisionJOMA:
		return JO
	case SubdivisionJOMD:
		return JO
	case SubdivisionJOMN:
		return JO
	case SubdivisionJP01:
		return JP
	case SubdivisionJP02:
		return JP
	case SubdivisionJP03:
		return JP
	case SubdivisionJP04:
		return JP
	case SubdivisionJP05:
		return JP
	case SubdivisionJP06:
		return JP
	case SubdivisionJP07:
		return JP
	case SubdivisionJP08:
		return JP
	case SubdivisionJP09:
		return JP
	case SubdivisionJP10:
		return JP
	case SubdivisionJP11:
		return JP
	case SubdivisionJP12:
		return JP
	case SubdivisionJP13:
		return JP
	case SubdivisionJP14:
		return JP
	case SubdivisionJP15:
		return JP
	case SubdivisionJP16:
		return JP
	case SubdivisionJP17:
		return JP
	case SubdivisionJP18:
		return JP
	case SubdivisionJP19:
		return JP
	case SubdivisionJP20:
		return JP
	case SubdivisionJP21:
		return JP
	case SubdivisionJP22:
		return JP
	case SubdivisionJP23:
		return JP
	case SubdivisionJP24:
		return JP
	case SubdivisionJP25:
		return JP
	case SubdivisionJP26:
		return JP
	case SubdivisionJP27:
		return JP
	case SubdivisionJP28:
		return JP
	case SubdivisionJP29:
		return JP
	case SubdivisionJP30:
		return JP
	case SubdivisionJP31:
		return JP
	case SubdivisionJP32:
		return JP
	case SubdivisionJP33:
		return JP
	case SubdivisionJP34:
		return JP
	case SubdivisionJP35:
		return JP
	case SubdivisionJP36:
		return JP
	case SubdivisionJP37:
		return JP
	case SubdivisionJP38:
		return JP
	case SubdivisionJP39:
		return JP
	case SubdivisionJP40:
		return JP
	case SubdivisionJP41:
		return JP
	case SubdivisionJP42:
		return JP
	case SubdivisionJP43:
		return JP
	case SubdivisionJP44:
		return JP
	case SubdivisionJP45:
		return JP
	case SubdivisionJP46:
		return JP
	case SubdivisionJP47:
		return JP
	case SubdivisionKE01:
		return KE
	case SubdivisionKE02:
		return KE
	case SubdivisionKE03:
		return KE
	case SubdivisionKE04:
		return KE
	case SubdivisionKE05:
		return KE
	case SubdivisionKE06:
		return KE
	case SubdivisionKE07:
		return KE
	case SubdivisionKE08:
		return KE
	case SubdivisionKE09:
		return KE
	case SubdivisionKE10:
		return KE
	case SubdivisionKE11:
		return KE
	case SubdivisionKE12:
		return KE
	case SubdivisionKE13:
		return KE
	case SubdivisionKE14:
		return KE
	case SubdivisionKE15:
		return KE
	case SubdivisionKE16:
		return KE
	case SubdivisionKE17:
		return KE
	case SubdivisionKE18:
		return KE
	case SubdivisionKE19:
		return KE
	case SubdivisionKE20:
		return KE
	case SubdivisionKE21:
		return KE
	case SubdivisionKE22:
		return KE
	case SubdivisionKE23:
		return KE
	case SubdivisionKE24:
		return KE
	case SubdivisionKE25:
		return KE
	case SubdivisionKE26:
		return KE
	case SubdivisionKE27:
		return KE
	case SubdivisionKE28:
		return KE
	case SubdivisionKE29:
		return KE
	case SubdivisionKE30:
		return KE
	case SubdivisionKE31:
		return KE
	case SubdivisionKE32:
		return KE
	case SubdivisionKE33:
		return KE
	case SubdivisionKE34:
		return KE
	case SubdivisionKE35:
		return KE
	case SubdivisionKE36:
		return KE
	case SubdivisionKE37:
		return KE
	case SubdivisionKE38:
		return KE
	case SubdivisionKE39:
		return KE
	case SubdivisionKE40:
		return KE
	case SubdivisionKE41:
		return KE
	case SubdivisionKE42:
		return KE
	case SubdivisionKE43:
		return KE
	case SubdivisionKE44:
		return KE
	case SubdivisionKE45:
		return KE
	case SubdivisionKE46:
		return KE
	case SubdivisionKE47:
		return KE
	case SubdivisionKGB:
		return KG
	case SubdivisionKGC:
		return KG
	case SubdivisionKGGB:
		return KG
	case SubdivisionKGJ:
		return KG
	case SubdivisionKGN:
		return KG
	case SubdivisionKGO:
		return KG
	case SubdivisionKGT:
		return KG
	case SubdivisionKGY:
		return KG
	case SubdivisionKH1:
		return KH
	case SubdivisionKH10:
		return KH
	case SubdivisionKH11:
		return KH
	case SubdivisionKH12:
		return KH
	case SubdivisionKH13:
		return KH
	case SubdivisionKH14:
		return KH
	case SubdivisionKH15:
		return KH
	case SubdivisionKH16:
		return KH
	case SubdivisionKH17:
		return KH
	case SubdivisionKH18:
		return KH
	case SubdivisionKH19:
		return KH
	case SubdivisionKH2:
		return KH
	case SubdivisionKH20:
		return KH
	case SubdivisionKH21:
		return KH
	case SubdivisionKH22:
		return KH
	case SubdivisionKH23:
		return KH
	case SubdivisionKH24:
		return KH
	case SubdivisionKH3:
		return KH
	case SubdivisionKH4:
		return KH
	case SubdivisionKH5:
		return KH
	case SubdivisionKH6:
		return KH
	case SubdivisionKH7:
		return KH
	case SubdivisionKH8:
		return KH
	case SubdivisionKH9:
		return KH
	case SubdivisionKIG:
		return KI
	case SubdivisionKIL:
		return KI
	case SubdivisionKIP:
		return KI
	case SubdivisionKMA:
		return KM
	case SubdivisionKMG:
		return KM
	case SubdivisionKMM:
		return KM
	case SubdivisionKN01:
		return KN
	case SubdivisionKN02:
		return KN
	case SubdivisionKN03:
		return KN
	case SubdivisionKN04:
		return KN
	case SubdivisionKN05:
		return KN
	case SubdivisionKN06:
		return KN
	case SubdivisionKN07:
		return KN
	case SubdivisionKN08:
		return KN
	case SubdivisionKN09:
		return KN
	case SubdivisionKN10:
		return KN
	case SubdivisionKN11:
		return KN
	case SubdivisionKN12:
		return KN
	case SubdivisionKN13:
		return KN
	case SubdivisionKN15:
		return KN
	case SubdivisionKNK:
		return KN
	case SubdivisionKNN:
		return KN
	case SubdivisionKP01:
		return KP
	case SubdivisionKP02:
		return KP
	case SubdivisionKP03:
		return KP
	case SubdivisionKP04:
		return KP
	case SubdivisionKP05:
		return KP
	case SubdivisionKP06:
		return KP
	case SubdivisionKP07:
		return KP
	case SubdivisionKP08:
		return KP
	case SubdivisionKP09:
		return KP
	case SubdivisionKP10:
		return KP
	case SubdivisionKP13:
		return KP
	case SubdivisionKR11:
		return KR
	case SubdivisionKR26:
		return KR
	case SubdivisionKR27:
		return KR
	case SubdivisionKR28:
		return KR
	case SubdivisionKR29:
		return KR
	case SubdivisionKR30:
		return KR
	case SubdivisionKR31:
		return KR
	case SubdivisionKR41:
		return KR
	case SubdivisionKR42:
		return KR
	case SubdivisionKR43:
		return KR
	case SubdivisionKR44:
		return KR
	case SubdivisionKR45:
		return KR
	case SubdivisionKR46:
		return KR
	case SubdivisionKR47:
		return KR
	case SubdivisionKR48:
		return KR
	case SubdivisionKR49:
		return KR
	case SubdivisionKWAH:
		return KW
	case SubdivisionKWFA:
		return KW
	case SubdivisionKWHA:
		return KW
	case SubdivisionKWJA:
		return KW
	case SubdivisionKWKU:
		return KW
	case SubdivisionKWMU:
		return KW
	case SubdivisionKZAKM:
		return KZ
	case SubdivisionKZAKT:
		return KZ
	case SubdivisionKZALA:
		return KZ
	case SubdivisionKZALM:
		return KZ
	case SubdivisionKZAST:
		return KZ
	case SubdivisionKZATY:
		return KZ
	case SubdivisionKZKAR:
		return KZ
	case SubdivisionKZKUS:
		return KZ
	case SubdivisionKZKZY:
		return KZ
	case SubdivisionKZMAN:
		return KZ
	case SubdivisionKZPAV:
		return KZ
	case SubdivisionKZSEV:
		return KZ
	case SubdivisionKZVOS:
		return KZ
	case SubdivisionKZYUZ:
		return KZ
	case SubdivisionKZZAP:
		return KZ
	case SubdivisionKZZHA:
		return KZ
	case SubdivisionLAAT:
		return LA
	case SubdivisionLABK:
		return LA
	case SubdivisionLABL:
		return LA
	case SubdivisionLACH:
		return LA
	case SubdivisionLAHO:
		return LA
	case SubdivisionLAKH:
		return LA
	case SubdivisionLALM:
		return LA
	case SubdivisionLALP:
		return LA
	case SubdivisionLAOU:
		return LA
	case SubdivisionLAPH:
		return LA
	case SubdivisionLASL:
		return LA
	case SubdivisionLASV:
		return LA
	case SubdivisionLAVI:
		return LA
	case SubdivisionLAVT:
		return LA
	case SubdivisionLAXA:
		return LA
	case SubdivisionLAXE:
		return LA
	case SubdivisionLAXI:
		return LA
	case SubdivisionLAXS:
		return LA
	case SubdivisionLBAK:
		return LB
	case SubdivisionLBAS:
		return LB
	case SubdivisionLBBA:
		return LB
	case SubdivisionLBBH:
		return LB
	case SubdivisionLBBI:
		return LB
	case SubdivisionLBJA:
		return LB
	case SubdivisionLBJL:
		return LB
	case SubdivisionLBNA:
		return LB
	case SubdivisionLI01:
		return LI
	case SubdivisionLI02:
		return LI
	case SubdivisionLI03:
		return LI
	case SubdivisionLI04:
		return LI
	case SubdivisionLI05:
		return LI
	case SubdivisionLI06:
		return LI
	case SubdivisionLI07:
		return LI
	case SubdivisionLI08:
		return LI
	case SubdivisionLI09:
		return LI
	case SubdivisionLI10:
		return LI
	case SubdivisionLI11:
		return LI
	case SubdivisionLK1:
		return LK
	case SubdivisionLK11:
		return LK
	case SubdivisionLK12:
		return LK
	case SubdivisionLK13:
		return LK
	case SubdivisionLK2:
		return LK
	case SubdivisionLK21:
		return LK
	case SubdivisionLK22:
		return LK
	case SubdivisionLK23:
		return LK
	case SubdivisionLK3:
		return LK
	case SubdivisionLK31:
		return LK
	case SubdivisionLK32:
		return LK
	case SubdivisionLK33:
		return LK
	case SubdivisionLK4:
		return LK
	case SubdivisionLK41:
		return LK
	case SubdivisionLK42:
		return LK
	case SubdivisionLK43:
		return LK
	case SubdivisionLK44:
		return LK
	case SubdivisionLK45:
		return LK
	case SubdivisionLK5:
		return LK
	case SubdivisionLK51:
		return LK
	case SubdivisionLK52:
		return LK
	case SubdivisionLK53:
		return LK
	case SubdivisionLK6:
		return LK
	case SubdivisionLK61:
		return LK
	case SubdivisionLK62:
		return LK
	case SubdivisionLK7:
		return LK
	case SubdivisionLK71:
		return LK
	case SubdivisionLK72:
		return LK
	case SubdivisionLK8:
		return LK
	case SubdivisionLK81:
		return LK
	case SubdivisionLK82:
		return LK
	case SubdivisionLK9:
		return LK
	case SubdivisionLK91:
		return LK
	case SubdivisionLK92:
		return LK
	case SubdivisionLRBG:
		return LR
	case SubdivisionLRBM:
		return LR
	case SubdivisionLRCM:
		return LR
	case SubdivisionLRGB:
		return LR
	case SubdivisionLRGG:
		return LR
	case SubdivisionLRGK:
		return LR
	case SubdivisionLRLO:
		return LR
	case SubdivisionLRMG:
		return LR
	case SubdivisionLRMO:
		return LR
	case SubdivisionLRMY:
		return LR
	case SubdivisionLRNI:
		return LR
	case SubdivisionLRRI:
		return LR
	case SubdivisionLRSI:
		return LR
	case SubdivisionLSA:
		return LS
	case SubdivisionLSB:
		return LS
	case SubdivisionLSC:
		return LS
	case SubdivisionLSD:
		return LS
	case SubdivisionLSE:
		return LS
	case SubdivisionLSF:
		return LS
	case SubdivisionLSG:
		return LS
	case SubdivisionLSH:
		return LS
	case SubdivisionLSJ:
		return LS
	case SubdivisionLSK:
		return LS
	case SubdivisionLTAL:
		return LT
	case SubdivisionLTKL:
		return LT
	case SubdivisionLTKU:
		return LT
	case SubdivisionLTMR:
		return LT
	case SubdivisionLTPN:
		return LT
	case SubdivisionLTSA:
		return LT
	case SubdivisionLTTA:
		return LT
	case SubdivisionLTTE:
		return LT
	case SubdivisionLTUT:
		return LT
	case SubdivisionLTVL:
		return LT
	case SubdivisionLUD:
		return LU
	case SubdivisionLUG:
		return LU
	case SubdivisionLUL:
		return LU
	case SubdivisionLV001:
		return LV
	case SubdivisionLV002:
		return LV
	case SubdivisionLV003:
		return LV
	case SubdivisionLV004:
		return LV
	case SubdivisionLV005:
		return LV
	case SubdivisionLV006:
		return LV
	case SubdivisionLV007:
		return LV
	case SubdivisionLV008:
		return LV
	case SubdivisionLV009:
		return LV
	case SubdivisionLV010:
		return LV
	case SubdivisionLV011:
		return LV
	case SubdivisionLV012:
		return LV
	case SubdivisionLV013:
		return LV
	case SubdivisionLV014:
		return LV
	case SubdivisionLV015:
		return LV
	case SubdivisionLV016:
		return LV
	case SubdivisionLV017:
		return LV
	case SubdivisionLV018:
		return LV
	case SubdivisionLV019:
		return LV
	case SubdivisionLV020:
		return LV
	case SubdivisionLV021:
		return LV
	case SubdivisionLV022:
		return LV
	case SubdivisionLV023:
		return LV
	case SubdivisionLV024:
		return LV
	case SubdivisionLV025:
		return LV
	case SubdivisionLV026:
		return LV
	case SubdivisionLV027:
		return LV
	case SubdivisionLV028:
		return LV
	case SubdivisionLV029:
		return LV
	case SubdivisionLV030:
		return LV
	case SubdivisionLV031:
		return LV
	case SubdivisionLV032:
		return LV
	case SubdivisionLV033:
		return LV
	case SubdivisionLV034:
		return LV
	case SubdivisionLV035:
		return LV
	case SubdivisionLV036:
		return LV
	case SubdivisionLV037:
		return LV
	case SubdivisionLV038:
		return LV
	case SubdivisionLV039:
		return LV
	case SubdivisionLV040:
		return LV
	case SubdivisionLV041:
		return LV
	case SubdivisionLV042:
		return LV
	case SubdivisionLV043:
		return LV
	case SubdivisionLV044:
		return LV
	case SubdivisionLV045:
		return LV
	case SubdivisionLV046:
		return LV
	case SubdivisionLV047:
		return LV
	case SubdivisionLV048:
		return LV
	case SubdivisionLV049:
		return LV
	case SubdivisionLV050:
		return LV
	case SubdivisionLV051:
		return LV
	case SubdivisionLV052:
		return LV
	case SubdivisionLV053:
		return LV
	case SubdivisionLV054:
		return LV
	case SubdivisionLV055:
		return LV
	case SubdivisionLV056:
		return LV
	case SubdivisionLV057:
		return LV
	case SubdivisionLV058:
		return LV
	case SubdivisionLV059:
		return LV
	case SubdivisionLV060:
		return LV
	case SubdivisionLV061:
		return LV
	case SubdivisionLV062:
		return LV
	case SubdivisionLV063:
		return LV
	case SubdivisionLV064:
		return LV
	case SubdivisionLV065:
		return LV
	case SubdivisionLV066:
		return LV
	case SubdivisionLV067:
		return LV
	case SubdivisionLV068:
		return LV
	case SubdivisionLV069:
		return LV
	case SubdivisionLV070:
		return LV
	case SubdivisionLV071:
		return LV
	case SubdivisionLV072:
		return LV
	case SubdivisionLV073:
		return LV
	case SubdivisionLV074:
		return LV
	case SubdivisionLV075:
		return LV
	case SubdivisionLV076:
		return LV
	case SubdivisionLV077:
		return LV
	case SubdivisionLV078:
		return LV
	case SubdivisionLV079:
		return LV
	case SubdivisionLV080:
		return LV
	case SubdivisionLV081:
		return LV
	case SubdivisionLV082:
		return LV
	case SubdivisionLV083:
		return LV
	case SubdivisionLV084:
		return LV
	case SubdivisionLV085:
		return LV
	case SubdivisionLV086:
		return LV
	case SubdivisionLV087:
		return LV
	case SubdivisionLV088:
		return LV
	case SubdivisionLV089:
		return LV
	case SubdivisionLV090:
		return LV
	case SubdivisionLV091:
		return LV
	case SubdivisionLV092:
		return LV
	case SubdivisionLV093:
		return LV
	case SubdivisionLV094:
		return LV
	case SubdivisionLV095:
		return LV
	case SubdivisionLV096:
		return LV
	case SubdivisionLV097:
		return LV
	case SubdivisionLV098:
		return LV
	case SubdivisionLV099:
		return LV
	case SubdivisionLV100:
		return LV
	case SubdivisionLV101:
		return LV
	case SubdivisionLV102:
		return LV
	case SubdivisionLV103:
		return LV
	case SubdivisionLV104:
		return LV
	case SubdivisionLV105:
		return LV
	case SubdivisionLV106:
		return LV
	case SubdivisionLV107:
		return LV
	case SubdivisionLV108:
		return LV
	case SubdivisionLV109:
		return LV
	case SubdivisionLV110:
		return LV
	case SubdivisionLVDGV:
		return LV
	case SubdivisionLVJEL:
		return LV
	case SubdivisionLVJKB:
		return LV
	case SubdivisionLVJUR:
		return LV
	case SubdivisionLVLPX:
		return LV
	case SubdivisionLVREZ:
		return LV
	case SubdivisionLVRIX:
		return LV
	case SubdivisionLVVEN:
		return LV
	case SubdivisionLVVMR:
		return LV
	case SubdivisionLYBA:
		return LY
	case SubdivisionLYBU:
		return LY
	case SubdivisionLYDR:
		return LY
	case SubdivisionLYGT:
		return LY
	case SubdivisionLYJA:
		return LY
	case SubdivisionLYJB:
		return LY
	case SubdivisionLYJG:
		return LY
	case SubdivisionLYJI:
		return LY
	case SubdivisionLYJU:
		return LY
	case SubdivisionLYKF:
		return LY
	case SubdivisionLYMB:
		return LY
	case SubdivisionLYMI:
		return LY
	case SubdivisionLYMJ:
		return LY
	case SubdivisionLYMQ:
		return LY
	case SubdivisionLYNL:
		return LY
	case SubdivisionLYNQ:
		return LY
	case SubdivisionLYSB:
		return LY
	case SubdivisionLYSR:
		return LY
	case SubdivisionLYTB:
		return LY
	case SubdivisionLYWA:
		return LY
	case SubdivisionLYWD:
		return LY
	case SubdivisionLYWS:
		return LY
	case SubdivisionLYZA:
		return LY
	case SubdivisionMA01:
		return MA
	case SubdivisionMA02:
		return MA
	case SubdivisionMA03:
		return MA
	case SubdivisionMA04:
		return MA
	case SubdivisionMA05:
		return MA
	case SubdivisionMA06:
		return MA
	case SubdivisionMA07:
		return MA
	case SubdivisionMA08:
		return MA
	case SubdivisionMA09:
		return MA
	case SubdivisionMA10:
		return MA
	case SubdivisionMA11:
		return MA
	case SubdivisionMA12:
		return MA
	case SubdivisionMAAGD:
		return MA
	case SubdivisionMAAOU:
		return MA
	case SubdivisionMAASZ:
		return MA
	case SubdivisionMAAZI:
		return MA
	case SubdivisionMABEM:
		return MA
	case SubdivisionMABER:
		return MA
	case SubdivisionMABES:
		return MA
	case SubdivisionMABOD:
		return MA
	case SubdivisionMABOM:
		return MA
	case SubdivisionMABRR:
		return MA
	case SubdivisionMACAS:
		return MA
	case SubdivisionMACHE:
		return MA
	case SubdivisionMACHI:
		return MA
	case SubdivisionMACHT:
		return MA
	case SubdivisionMADRI:
		return MA
	case SubdivisionMAERR:
		return MA
	case SubdivisionMAESI:
		return MA
	case SubdivisionMAESM:
		return MA
	case SubdivisionMAFAH:
		return MA
	case SubdivisionMAFES:
		return MA
	case SubdivisionMAFIG:
		return MA
	case SubdivisionMAFQH:
		return MA
	case SubdivisionMAGUE:
		return MA
	case SubdivisionMAGUF:
		return MA
	case SubdivisionMAHAJ:
		return MA
	case SubdivisionMAHAO:
		return MA
	case SubdivisionMAHOC:
		return MA
	case SubdivisionMAIFR:
		return MA
	case SubdivisionMAINE:
		return MA
	case SubdivisionMAJDI:
		return MA
	case SubdivisionMAJRA:
		return MA
	case SubdivisionMAKEN:
		return MA
	case SubdivisionMAKES:
		return MA
	case SubdivisionMAKHE:
		return MA
	case SubdivisionMAKHN:
		return MA
	case SubdivisionMAKHO:
		return MA
	case SubdivisionMALAA:
		return MA
	case SubdivisionMALAR:
		return MA
	case SubdivisionMAMAR:
		return MA
	case SubdivisionMAMDF:
		return MA
	case SubdivisionMAMED:
		return MA
	case SubdivisionMAMEK:
		return MA
	case SubdivisionMAMID:
		return MA
	case SubdivisionMAMOH:
		return MA
	case SubdivisionMAMOU:
		return MA
	case SubdivisionMANAD:
		return MA
	case SubdivisionMANOU:
		return MA
	case SubdivisionMAOUA:
		return MA
	case SubdivisionMAOUD:
		return MA
	case SubdivisionMAOUJ:
		return MA
	case SubdivisionMAOUZ:
		return MA
	case SubdivisionMARAB:
		return MA
	case SubdivisionMAREH:
		return MA
	case SubdivisionMASAF:
		return MA
	case SubdivisionMASAL:
		return MA
	case SubdivisionMASEF:
		return MA
	case SubdivisionMASET:
		return MA
	case SubdivisionMASIB:
		return MA
	case SubdivisionMASIF:
		return MA
	case SubdivisionMASIK:
		return MA
	case SubdivisionMASIL:
		return MA
	case SubdivisionMASKH:
		return MA
	case SubdivisionMATAF:
		return MA
	case SubdivisionMATAI:
		return MA
	case SubdivisionMATAO:
		return MA
	case SubdivisionMATAR:
		return MA
	case SubdivisionMATAT:
		return MA
	case SubdivisionMATAZ:
		return MA
	case SubdivisionMATET:
		return MA
	case SubdivisionMATIN:
		return MA
	case SubdivisionMATIZ:
		return MA
	case SubdivisionMATNG:
		return MA
	case SubdivisionMATNT:
		return MA
	case SubdivisionMAYUS:
		return MA
	case SubdivisionMAZAG:
		return MA
	case SubdivisionMCCL:
		return MC
	case SubdivisionMCCO:
		return MC
	case SubdivisionMCFO:
		return MC
	case SubdivisionMCGA:
		return MC
	case SubdivisionMCJE:
		return MC
	case SubdivisionMCLA:
		return MC
	case SubdivisionMCMA:
		return MC
	case SubdivisionMCMC:
		return MC
	case SubdivisionMCMG:
		return MC
	case SubdivisionMCMO:
		return MC
	case SubdivisionMCMU:
		return MC
	case SubdivisionMCPH:
		return MC
	case SubdivisionMCSD:
		return MC
	case SubdivisionMCSO:
		return MC
	case SubdivisionMCSP:
		return MC
	case SubdivisionMCSR:
		return MC
	case SubdivisionMCVR:
		return MC
	case SubdivisionMDAN:
		return MD
	case SubdivisionMDBA:
		return MD
	case SubdivisionMDBD:
		return MD
	case SubdivisionMDBR:
		return MD
	case SubdivisionMDBS:
		return MD
	case SubdivisionMDCA:
		return MD
	case SubdivisionMDCL:
		return MD
	case SubdivisionMDCM:
		return MD
	case SubdivisionMDCR:
		return MD
	case SubdivisionMDCS:
		return MD
	case SubdivisionMDCT:
		return MD
	case SubdivisionMDCU:
		return MD
	case SubdivisionMDDO:
		return MD
	case SubdivisionMDDR:
		return MD
	case SubdivisionMDDU:
		return MD
	case SubdivisionMDED:
		return MD
	case SubdivisionMDFA:
		return MD
	case SubdivisionMDFL:
		return MD
	case SubdivisionMDGA:
		return MD
	case SubdivisionMDGL:
		return MD
	case SubdivisionMDHI:
		return MD
	case SubdivisionMDIA:
		return MD
	case SubdivisionMDLE:
		return MD
	case SubdivisionMDNI:
		return MD
	case SubdivisionMDOC:
		return MD
	case SubdivisionMDOR:
		return MD
	case SubdivisionMDRE:
		return MD
	case SubdivisionMDRI:
		return MD
	case SubdivisionMDSD:
		return MD
	case SubdivisionMDSI:
		return MD
	case SubdivisionMDSN:
		return MD
	case SubdivisionMDSO:
		return MD
	case SubdivisionMDST:
		return MD
	case SubdivisionMDSV:
		return MD
	case SubdivisionMDTA:
		return MD
	case SubdivisionMDTE:
		return MD
	case SubdivisionMDUN:
		return MD
	case SubdivisionME01:
		return ME
	case SubdivisionME02:
		return ME
	case SubdivisionME03:
		return ME
	case SubdivisionME04:
		return ME
	case SubdivisionME05:
		return ME
	case SubdivisionME06:
		return ME
	case SubdivisionME07:
		return ME
	case SubdivisionME08:
		return ME
	case SubdivisionME09:
		return ME
	case SubdivisionME10:
		return ME
	case SubdivisionME11:
		return ME
	case SubdivisionME12:
		return ME
	case SubdivisionME13:
		return ME
	case SubdivisionME14:
		return ME
	case SubdivisionME15:
		return ME
	case SubdivisionME16:
		return ME
	case SubdivisionME17:
		return ME
	case SubdivisionME18:
		return ME
	case SubdivisionME19:
		return ME
	case SubdivisionME20:
		return ME
	case SubdivisionME21:
		return ME
	case SubdivisionMGA:
		return MG
	case SubdivisionMGD:
		return MG
	case SubdivisionMGF:
		return MG
	case SubdivisionMGM:
		return MG
	case SubdivisionMGT:
		return MG
	case SubdivisionMGU:
		return MG
	case SubdivisionMHALK:
		return MH
	case SubdivisionMHALL:
		return MH
	case SubdivisionMHARN:
		return MH
	case SubdivisionMHAUR:
		return MH
	case SubdivisionMHEBO:
		return MH
	case SubdivisionMHENI:
		return MH
	case SubdivisionMHJAB:
		return MH
	case SubdivisionMHJAL:
		return MH
	case SubdivisionMHKIL:
		return MH
	case SubdivisionMHKWA:
		return MH
	case SubdivisionMHL:
		return MH
	case SubdivisionMHLAE:
		return MH
	case SubdivisionMHLIB:
		return MH
	case SubdivisionMHLIK:
		return MH
	case SubdivisionMHMAJ:
		return MH
	case SubdivisionMHMAL:
		return MH
	case SubdivisionMHMEJ:
		return MH
	case SubdivisionMHMIL:
		return MH
	case SubdivisionMHNMK:
		return MH
	case SubdivisionMHNMU:
		return MH
	case SubdivisionMHRON:
		return MH
	case SubdivisionMHT:
		return MH
	case SubdivisionMHUJA:
		return MH
	case SubdivisionMHUTI:
		return MH
	case SubdivisionMHWTJ:
		return MH
	case SubdivisionMHWTN:
		return MH
	case SubdivisionMK01:
		return MK
	case SubdivisionMK02:
		return MK
	case SubdivisionMK03:
		return MK
	case SubdivisionMK04:
		return MK
	case SubdivisionMK05:
		return MK
	case SubdivisionMK06:
		return MK
	case SubdivisionMK07:
		return MK
	case SubdivisionMK08:
		return MK
	case SubdivisionMK09:
		return MK
	case SubdivisionMK10:
		return MK
	case SubdivisionMK11:
		return MK
	case SubdivisionMK12:
		return MK
	case SubdivisionMK13:
		return MK
	case SubdivisionMK14:
		return MK
	case SubdivisionMK15:
		return MK
	case SubdivisionMK16:
		return MK
	case SubdivisionMK17:
		return MK
	case SubdivisionMK18:
		return MK
	case SubdivisionMK19:
		return MK
	case SubdivisionMK20:
		return MK
	case SubdivisionMK21:
		return MK
	case SubdivisionMK22:
		return MK
	case SubdivisionMK23:
		return MK
	case SubdivisionMK24:
		return MK
	case SubdivisionMK25:
		return MK
	case SubdivisionMK26:
		return MK
	case SubdivisionMK27:
		return MK
	case SubdivisionMK28:
		return MK
	case SubdivisionMK29:
		return MK
	case SubdivisionMK30:
		return MK
	case SubdivisionMK31:
		return MK
	case SubdivisionMK32:
		return MK
	case SubdivisionMK33:
		return MK
	case SubdivisionMK34:
		return MK
	case SubdivisionMK35:
		return MK
	case SubdivisionMK36:
		return MK
	case SubdivisionMK37:
		return MK
	case SubdivisionMK38:
		return MK
	case SubdivisionMK39:
		return MK
	case SubdivisionMK40:
		return MK
	case SubdivisionMK41:
		return MK
	case SubdivisionMK42:
		return MK
	case SubdivisionMK43:
		return MK
	case SubdivisionMK44:
		return MK
	case SubdivisionMK45:
		return MK
	case SubdivisionMK46:
		return MK
	case SubdivisionMK47:
		return MK
	case SubdivisionMK48:
		return MK
	case SubdivisionMK49:
		return MK
	case SubdivisionMK50:
		return MK
	case SubdivisionMK51:
		return MK
	case SubdivisionMK52:
		return MK
	case SubdivisionMK53:
		return MK
	case SubdivisionMK54:
		return MK
	case SubdivisionMK55:
		return MK
	case SubdivisionMK56:
		return MK
	case SubdivisionMK57:
		return MK
	case SubdivisionMK58:
		return MK
	case SubdivisionMK59:
		return MK
	case SubdivisionMK60:
		return MK
	case SubdivisionMK61:
		return MK
	case SubdivisionMK62:
		return MK
	case SubdivisionMK63:
		return MK
	case SubdivisionMK64:
		return MK
	case SubdivisionMK65:
		return MK
	case SubdivisionMK66:
		return MK
	case SubdivisionMK67:
		return MK
	case SubdivisionMK68:
		return MK
	case SubdivisionMK69:
		return MK
	case SubdivisionMK70:
		return MK
	case SubdivisionMK71:
		return MK
	case SubdivisionMK72:
		return MK
	case SubdivisionMK73:
		return MK
	case SubdivisionMK74:
		return MK
	case SubdivisionMK75:
		return MK
	case SubdivisionMK76:
		return MK
	case SubdivisionMK77:
		return MK
	case SubdivisionMK78:
		return MK
	case SubdivisionMK79:
		return MK
	case SubdivisionMK80:
		return MK
	case SubdivisionMK81:
		return MK
	case SubdivisionMK82:
		return MK
	case SubdivisionMK83:
		return MK
	case SubdivisionMK84:
		return MK
	case SubdivisionML1:
		return ML
	case SubdivisionML2:
		return ML
	case SubdivisionML3:
		return ML
	case SubdivisionML4:
		return ML
	case SubdivisionML5:
		return ML
	case SubdivisionML6:
		return ML
	case SubdivisionML7:
		return ML
	case SubdivisionML8:
		return ML
	case SubdivisionMLBK0:
		return ML
	case SubdivisionMM01:
		return MM
	case SubdivisionMM02:
		return MM
	case SubdivisionMM03:
		return MM
	case SubdivisionMM04:
		return MM
	case SubdivisionMM05:
		return MM
	case SubdivisionMM06:
		return MM
	case SubdivisionMM07:
		return MM
	case SubdivisionMM11:
		return MM
	case SubdivisionMM12:
		return MM
	case SubdivisionMM13:
		return MM
	case SubdivisionMM14:
		return MM
	case SubdivisionMM15:
		return MM
	case SubdivisionMM16:
		return MM
	case SubdivisionMM17:
		return MM
	case SubdivisionMN035:
		return MN
	case SubdivisionMN037:
		return MN
	case SubdivisionMN039:
		return MN
	case SubdivisionMN041:
		return MN
	case SubdivisionMN043:
		return MN
	case SubdivisionMN046:
		return MN
	case SubdivisionMN047:
		return MN
	case SubdivisionMN049:
		return MN
	case SubdivisionMN051:
		return MN
	case SubdivisionMN053:
		return MN
	case SubdivisionMN055:
		return MN
	case SubdivisionMN057:
		return MN
	case SubdivisionMN059:
		return MN
	case SubdivisionMN061:
		return MN
	case SubdivisionMN063:
		return MN
	case SubdivisionMN064:
		return MN
	case SubdivisionMN065:
		return MN
	case SubdivisionMN067:
		return MN
	case SubdivisionMN069:
		return MN
	case SubdivisionMN071:
		return MN
	case SubdivisionMN073:
		return MN
	case SubdivisionMN1:
		return MN
	case SubdivisionMR01:
		return MR
	case SubdivisionMR02:
		return MR
	case SubdivisionMR03:
		return MR
	case SubdivisionMR04:
		return MR
	case SubdivisionMR05:
		return MR
	case SubdivisionMR06:
		return MR
	case SubdivisionMR07:
		return MR
	case SubdivisionMR08:
		return MR
	case SubdivisionMR09:
		return MR
	case SubdivisionMR10:
		return MR
	case SubdivisionMR11:
		return MR
	case SubdivisionMR12:
		return MR
	case SubdivisionMRNKC:
		return MR
	case SubdivisionMT01:
		return MT
	case SubdivisionMT02:
		return MT
	case SubdivisionMT03:
		return MT
	case SubdivisionMT04:
		return MT
	case SubdivisionMT05:
		return MT
	case SubdivisionMT06:
		return MT
	case SubdivisionMT07:
		return MT
	case SubdivisionMT08:
		return MT
	case SubdivisionMT09:
		return MT
	case SubdivisionMT10:
		return MT
	case SubdivisionMT11:
		return MT
	case SubdivisionMT12:
		return MT
	case SubdivisionMT13:
		return MT
	case SubdivisionMT14:
		return MT
	case SubdivisionMT15:
		return MT
	case SubdivisionMT16:
		return MT
	case SubdivisionMT17:
		return MT
	case SubdivisionMT18:
		return MT
	case SubdivisionMT19:
		return MT
	case SubdivisionMT20:
		return MT
	case SubdivisionMT21:
		return MT
	case SubdivisionMT22:
		return MT
	case SubdivisionMT23:
		return MT
	case SubdivisionMT24:
		return MT
	case SubdivisionMT25:
		return MT
	case SubdivisionMT26:
		return MT
	case SubdivisionMT27:
		return MT
	case SubdivisionMT28:
		return MT
	case SubdivisionMT29:
		return MT
	case SubdivisionMT30:
		return MT
	case SubdivisionMT31:
		return MT
	case SubdivisionMT32:
		return MT
	case SubdivisionMT33:
		return MT
	case SubdivisionMT34:
		return MT
	case SubdivisionMT35:
		return MT
	case SubdivisionMT36:
		return MT
	case SubdivisionMT37:
		return MT
	case SubdivisionMT38:
		return MT
	case SubdivisionMT39:
		return MT
	case SubdivisionMT40:
		return MT
	case SubdivisionMT41:
		return MT
	case SubdivisionMT42:
		return MT
	case SubdivisionMT43:
		return MT
	case SubdivisionMT44:
		return MT
	case SubdivisionMT45:
		return MT
	case SubdivisionMT46:
		return MT
	case SubdivisionMT47:
		return MT
	case SubdivisionMT48:
		return MT
	case SubdivisionMT49:
		return MT
	case SubdivisionMT50:
		return MT
	case SubdivisionMT51:
		return MT
	case SubdivisionMT52:
		return MT
	case SubdivisionMT53:
		return MT
	case SubdivisionMT54:
		return MT
	case SubdivisionMT55:
		return MT
	case SubdivisionMT56:
		return MT
	case SubdivisionMT57:
		return MT
	case SubdivisionMT58:
		return MT
	case SubdivisionMT59:
		return MT
	case SubdivisionMT60:
		return MT
	case SubdivisionMT61:
		return MT
	case SubdivisionMT62:
		return MT
	case SubdivisionMT63:
		return MT
	case SubdivisionMT64:
		return MT
	case SubdivisionMT65:
		return MT
	case SubdivisionMT66:
		return MT
	case SubdivisionMT67:
		return MT
	case SubdivisionMT68:
		return MT
	case SubdivisionMUAG:
		return MU
	case SubdivisionMUBL:
		return MU
	case SubdivisionMUBR:
		return MU
	case SubdivisionMUCC:
		return MU
	case SubdivisionMUCU:
		return MU
	case SubdivisionMUFL:
		return MU
	case SubdivisionMUGP:
		return MU
	case SubdivisionMUMO:
		return MU
	case SubdivisionMUPA:
		return MU
	case SubdivisionMUPL:
		return MU
	case SubdivisionMUPU:
		return MU
	case SubdivisionMUPW:
		return MU
	case SubdivisionMUQB:
		return MU
	case SubdivisionMURO:
		return MU
	case SubdivisionMURP:
		return MU
	case SubdivisionMUSA:
		return MU
	case SubdivisionMUVP:
		return MU
	case SubdivisionMV00:
		return MV
	case SubdivisionMV01:
		return MV
	case SubdivisionMV02:
		return MV
	case SubdivisionMV03:
		return MV
	case SubdivisionMV04:
		return MV
	case SubdivisionMV05:
		return MV
	case SubdivisionMV07:
		return MV
	case SubdivisionMV08:
		return MV
	case SubdivisionMV12:
		return MV
	case SubdivisionMV13:
		return MV
	case SubdivisionMV14:
		return MV
	case SubdivisionMV17:
		return MV
	case SubdivisionMV20:
		return MV
	case SubdivisionMV23:
		return MV
	case SubdivisionMV24:
		return MV
	case SubdivisionMV25:
		return MV
	case SubdivisionMV26:
		return MV
	case SubdivisionMV27:
		return MV
	case SubdivisionMV28:
		return MV
	case SubdivisionMV29:
		return MV
	case SubdivisionMVCE:
		return MV
	case SubdivisionMVMLE:
		return MV
	case SubdivisionMVNC:
		return MV
	case SubdivisionMVNO:
		return MV
	case SubdivisionMVSC:
		return MV
	case SubdivisionMVSU:
		return MV
	case SubdivisionMVUN:
		return MV
	case SubdivisionMVUS:
		return MV
	case SubdivisionMWBA:
		return MW
	case SubdivisionMWBL:
		return MW
	case SubdivisionMWC:
		return MW
	case SubdivisionMWCK:
		return MW
	case SubdivisionMWCR:
		return MW
	case SubdivisionMWCT:
		return MW
	case SubdivisionMWDE:
		return MW
	case SubdivisionMWDO:
		return MW
	case SubdivisionMWKR:
		return MW
	case SubdivisionMWKS:
		return MW
	case SubdivisionMWLI:
		return MW
	case SubdivisionMWLK:
		return MW
	case SubdivisionMWMC:
		return MW
	case SubdivisionMWMG:
		return MW
	case SubdivisionMWMH:
		return MW
	case SubdivisionMWMU:
		return MW
	case SubdivisionMWMW:
		return MW
	case SubdivisionMWMZ:
		return MW
	case SubdivisionMWN:
		return MW
	case SubdivisionMWNB:
		return MW
	case SubdivisionMWNE:
		return MW
	case SubdivisionMWNI:
		return MW
	case SubdivisionMWNK:
		return MW
	case SubdivisionMWNS:
		return MW
	case SubdivisionMWNU:
		return MW
	case SubdivisionMWPH:
		return MW
	case SubdivisionMWRU:
		return MW
	case SubdivisionMWS:
		return MW
	case SubdivisionMWSA:
		return MW
	case SubdivisionMWTH:
		return MW
	case SubdivisionMWZO:
		return MW
	case SubdivisionMXAGU:
		return MX
	case SubdivisionMXBCN:
		return MX
	case SubdivisionMXBCS:
		return MX
	case SubdivisionMXCAM:
		return MX
	case SubdivisionMXCHH:
		return MX
	case SubdivisionMXCHP:
		return MX
	case SubdivisionMXCMX:
		return MX
	case SubdivisionMXCOA:
		return MX
	case SubdivisionMXCOL:
		return MX
	case SubdivisionMXDUR:
		return MX
	case SubdivisionMXGRO:
		return MX
	case SubdivisionMXGUA:
		return MX
	case SubdivisionMXHID:
		return MX
	case SubdivisionMXJAL:
		return MX
	case SubdivisionMXMEX:
		return MX
	case SubdivisionMXMIC:
		return MX
	case SubdivisionMXMOR:
		return MX
	case SubdivisionMXNAY:
		return MX
	case SubdivisionMXNLE:
		return MX
	case SubdivisionMXOAX:
		return MX
	case SubdivisionMXPUE:
		return MX
	case SubdivisionMXQUE:
		return MX
	case SubdivisionMXROO:
		return MX
	case SubdivisionMXSIN:
		return MX
	case SubdivisionMXSLP:
		return MX
	case SubdivisionMXSON:
		return MX
	case SubdivisionMXTAB:
		return MX
	case SubdivisionMXTAM:
		return MX
	case SubdivisionMXTLA:
		return MX
	case SubdivisionMXVER:
		return MX
	case SubdivisionMXYUC:
		return MX
	case SubdivisionMXZAC:
		return MX
	case SubdivisionMY01:
		return MY
	case SubdivisionMY02:
		return MY
	case SubdivisionMY03:
		return MY
	case SubdivisionMY04:
		return MY
	case SubdivisionMY05:
		return MY
	case SubdivisionMY06:
		return MY
	case SubdivisionMY07:
		return MY
	case SubdivisionMY08:
		return MY
	case SubdivisionMY09:
		return MY
	case SubdivisionMY10:
		return MY
	case SubdivisionMY11:
		return MY
	case SubdivisionMY12:
		return MY
	case SubdivisionMY13:
		return MY
	case SubdivisionMY14:
		return MY
	case SubdivisionMY15:
		return MY
	case SubdivisionMY16:
		return MY
	case SubdivisionMZA:
		return MZ
	case SubdivisionMZB:
		return MZ
	case SubdivisionMZG:
		return MZ
	case SubdivisionMZI:
		return MZ
	case SubdivisionMZL:
		return MZ
	case SubdivisionMZMPM:
		return MZ
	case SubdivisionMZN:
		return MZ
	case SubdivisionMZP:
		return MZ
	case SubdivisionMZQ:
		return MZ
	case SubdivisionMZS:
		return MZ
	case SubdivisionMZT:
		return MZ
	case SubdivisionNACA:
		return NA
	case SubdivisionNAER:
		return NA
	case SubdivisionNAHA:
		return NA
	case SubdivisionNAKA:
		return NA
	case SubdivisionNAKH:
		return NA
	case SubdivisionNAKU:
		return NA
	case SubdivisionNAOD:
		return NA
	case SubdivisionNAOH:
		return NA
	case SubdivisionNAOK:
		return NA
	case SubdivisionNAON:
		return NA
	case SubdivisionNAOS:
		return NA
	case SubdivisionNAOT:
		return NA
	case SubdivisionNAOW:
		return NA
	case SubdivisionNE1:
		return NE
	case SubdivisionNE2:
		return NE
	case SubdivisionNE3:
		return NE
	case SubdivisionNE4:
		return NE
	case SubdivisionNE5:
		return NE
	case SubdivisionNE6:
		return NE
	case SubdivisionNE7:
		return NE
	case SubdivisionNE8:
		return NE
	case SubdivisionNGAB:
		return NG
	case SubdivisionNGAD:
		return NG
	case SubdivisionNGAK:
		return NG
	case SubdivisionNGAN:
		return NG
	case SubdivisionNGBA:
		return NG
	case SubdivisionNGBE:
		return NG
	case SubdivisionNGBO:
		return NG
	case SubdivisionNGBY:
		return NG
	case SubdivisionNGCR:
		return NG
	case SubdivisionNGDE:
		return NG
	case SubdivisionNGEB:
		return NG
	case SubdivisionNGED:
		return NG
	case SubdivisionNGEK:
		return NG
	case SubdivisionNGEN:
		return NG
	case SubdivisionNGFC:
		return NG
	case SubdivisionNGGO:
		return NG
	case SubdivisionNGIM:
		return NG
	case SubdivisionNGJI:
		return NG
	case SubdivisionNGKD:
		return NG
	case SubdivisionNGKE:
		return NG
	case SubdivisionNGKN:
		return NG
	case SubdivisionNGKO:
		return NG
	case SubdivisionNGKT:
		return NG
	case SubdivisionNGKW:
		return NG
	case SubdivisionNGLA:
		return NG
	case SubdivisionNGNA:
		return NG
	case SubdivisionNGNI:
		return NG
	case SubdivisionNGOG:
		return NG
	case SubdivisionNGON:
		return NG
	case SubdivisionNGOS:
		return NG
	case SubdivisionNGOY:
		return NG
	case SubdivisionNGPL:
		return NG
	case SubdivisionNGRI:
		return NG
	case SubdivisionNGSO:
		return NG
	case SubdivisionNGTA:
		return NG
	case SubdivisionNGYO:
		return NG
	case SubdivisionNGZA:
		return NG
	case SubdivisionNIAN:
		return NI
	case SubdivisionNIAS:
		return NI
	case SubdivisionNIBO:
		return NI
	case SubdivisionNICA:
		return NI
	case SubdivisionNICI:
		return NI
	case SubdivisionNICO:
		return NI
	case SubdivisionNIES:
		return NI
	case SubdivisionNIGR:
		return NI
	case SubdivisionNIJI:
		return NI
	case SubdivisionNILE:
		return NI
	case SubdivisionNIMD:
		return NI
	case SubdivisionNIMN:
		return NI
	case SubdivisionNIMS:
		return NI
	case SubdivisionNIMT:
		return NI
	case SubdivisionNINS:
		return NI
	case SubdivisionNIRI:
		return NI
	case SubdivisionNISJ:
		return NI
	case SubdivisionNLAW:
		return NL
	case SubdivisionNLBQ1:
		return NL
	case SubdivisionNLBQ2:
		return NL
	case SubdivisionNLBQ3:
		return NL
	case SubdivisionNLCW:
		return NL
	case SubdivisionNLDR:
		return NL
	case SubdivisionNLFL:
		return NL
	case SubdivisionNLFR:
		return NL
	case SubdivisionNLGE:
		return NL
	case SubdivisionNLGR:
		return NL
	case SubdivisionNLLI:
		return NL
	case SubdivisionNLNB:
		return NL
	case SubdivisionNLNH:
		return NL
	case SubdivisionNLOV:
		return NL
	case SubdivisionNLSX:
		return NL
	case SubdivisionNLUT:
		return NL
	case SubdivisionNLZE:
		return NL
	case SubdivisionNLZH:
		return NL
	case SubdivisionNO01:
		return NO
	case SubdivisionNO02:
		return NO
	case SubdivisionNO03:
		return NO
	case SubdivisionNO04:
		return NO
	case SubdivisionNO05:
		return NO
	case SubdivisionNO06:
		return NO
	case SubdivisionNO07:
		return NO
	case SubdivisionNO08:
		return NO
	case SubdivisionNO09:
		return NO
	case SubdivisionNO10:
		return NO
	case SubdivisionNO11:
		return NO
	case SubdivisionNO12:
		return NO
	case SubdivisionNO14:
		return NO
	case SubdivisionNO15:
		return NO
	case SubdivisionNO18:
		return NO
	case SubdivisionNO19:
		return NO
	case SubdivisionNO20:
		return NO
	case SubdivisionNO21:
		return NO
	case SubdivisionNO22:
		return NO
	case SubdivisionNO50:
		return NO
	case SubdivisionNP1:
		return NP
	case SubdivisionNP2:
		return NP
	case SubdivisionNP3:
		return NP
	case SubdivisionNP4:
		return NP
	case SubdivisionNP5:
		return NP
	case SubdivisionNPBA:
		return NP
	case SubdivisionNPBH:
		return NP
	case SubdivisionNPDH:
		return NP
	case SubdivisionNPGA:
		return NP
	case SubdivisionNPJA:
		return NP
	case SubdivisionNPKA:
		return NP
	case SubdivisionNPKO:
		return NP
	case SubdivisionNPLU:
		return NP
	case SubdivisionNPMA:
		return NP
	case SubdivisionNPME:
		return NP
	case SubdivisionNPNA:
		return NP
	case SubdivisionNPRA:
		return NP
	case SubdivisionNPSA:
		return NP
	case SubdivisionNPSE:
		return NP
	case SubdivisionNR01:
		return NR
	case SubdivisionNR02:
		return NR
	case SubdivisionNR03:
		return NR
	case SubdivisionNR04:
		return NR
	case SubdivisionNR05:
		return NR
	case SubdivisionNR06:
		return NR
	case SubdivisionNR07:
		return NR
	case SubdivisionNR08:
		return NR
	case SubdivisionNR09:
		return NR
	case SubdivisionNR10:
		return NR
	case SubdivisionNR11:
		return NR
	case SubdivisionNR12:
		return NR
	case SubdivisionNR13:
		return NR
	case SubdivisionNR14:
		return NR
	case SubdivisionNZAUK:
		return NZ
	case SubdivisionNZBOP:
		return NZ
	case SubdivisionNZCAN:
		return NZ
	case SubdivisionNZCIT:
		return NZ
	case SubdivisionNZGIS:
		return NZ
	case SubdivisionNZHKB:
		return NZ
	case SubdivisionNZMBH:
		return NZ
	case SubdivisionNZMWT:
		return NZ
	case SubdivisionNZN:
		return NZ
	case SubdivisionNZNSN:
		return NZ
	case SubdivisionNZNTL:
		return NZ
	case SubdivisionNZOTA:
		return NZ
	case SubdivisionNZS:
		return NZ
	case SubdivisionNZSTL:
		return NZ
	case SubdivisionNZTAS:
		return NZ
	case SubdivisionNZTKI:
		return NZ
	case SubdivisionNZWGN:
		return NZ
	case SubdivisionNZWKO:
		return NZ
	case SubdivisionNZWTC:
		return NZ
	case SubdivisionOMBA:
		return OM
	case SubdivisionOMBU:
		return OM
	case SubdivisionOMDA:
		return OM
	case SubdivisionOMMA:
		return OM
	case SubdivisionOMMU:
		return OM
	case SubdivisionOMSH:
		return OM
	case SubdivisionOMWU:
		return OM
	case SubdivisionOMZA:
		return OM
	case SubdivisionOMZU:
		return OM
	case SubdivisionPA1:
		return PA
	case SubdivisionPA2:
		return PA
	case SubdivisionPA3:
		return PA
	case SubdivisionPA4:
		return PA
	case SubdivisionPA5:
		return PA
	case SubdivisionPA6:
		return PA
	case SubdivisionPA7:
		return PA
	case SubdivisionPA8:
		return PA
	case SubdivisionPA9:
		return PA
	case SubdivisionPAEM:
		return PA
	case SubdivisionPAKY:
		return PA
	case SubdivisionPANB:
		return PA
	case SubdivisionPEAMA:
		return PE
	case SubdivisionPEANC:
		return PE
	case SubdivisionPEAPU:
		return PE
	case SubdivisionPEARE:
		return PE
	case SubdivisionPEAYA:
		return PE
	case SubdivisionPECAJ:
		return PE
	case SubdivisionPECAL:
		return PE
	case SubdivisionPECUS:
		return PE
	case SubdivisionPEHUC:
		return PE
	case SubdivisionPEHUV:
		return PE
	case SubdivisionPEICA:
		return PE
	case SubdivisionPEJUN:
		return PE
	case SubdivisionPELAL:
		return PE
	case SubdivisionPELAM:
		return PE
	case SubdivisionPELIM:
		return PE
	case SubdivisionPELMA:
		return PE
	case SubdivisionPELOR:
		return PE
	case SubdivisionPEMDD:
		return PE
	case SubdivisionPEMOQ:
		return PE
	case SubdivisionPEPAS:
		return PE
	case SubdivisionPEPIU:
		return PE
	case SubdivisionPEPUN:
		return PE
	case SubdivisionPESAM:
		return PE
	case SubdivisionPETAC:
		return PE
	case SubdivisionPETUM:
		return PE
	case SubdivisionPEUCA:
		return PE
	case SubdivisionPGCPK:
		return PG
	case SubdivisionPGCPM:
		return PG
	case SubdivisionPGEBR:
		return PG
	case SubdivisionPGEHG:
		return PG
	case SubdivisionPGEPW:
		return PG
	case SubdivisionPGESW:
		return PG
	case SubdivisionPGGPK:
		return PG
	case SubdivisionPGMBA:
		return PG
	case SubdivisionPGMPL:
		return PG
	case SubdivisionPGMPM:
		return PG
	case SubdivisionPGMRL:
		return PG
	case SubdivisionPGNCD:
		return PG
	case SubdivisionPGNIK:
		return PG
	case SubdivisionPGNPP:
		return PG
	case SubdivisionPGNSB:
		return PG
	case SubdivisionPGSAN:
		return PG
	case SubdivisionPGSHM:
		return PG
	case SubdivisionPGWBK:
		return PG
	case SubdivisionPGWHM:
		return PG
	case SubdivisionPGWPD:
		return PG
	case SubdivisionPH00:
		return PH
	case SubdivisionPH01:
		return PH
	case SubdivisionPH02:
		return PH
	case SubdivisionPH03:
		return PH
	case SubdivisionPH05:
		return PH
	case SubdivisionPH06:
		return PH
	case SubdivisionPH07:
		return PH
	case SubdivisionPH08:
		return PH
	case SubdivisionPH09:
		return PH
	case SubdivisionPH10:
		return PH
	case SubdivisionPH11:
		return PH
	case SubdivisionPH12:
		return PH
	case SubdivisionPH13:
		return PH
	case SubdivisionPH14:
		return PH
	case SubdivisionPH15:
		return PH
	case SubdivisionPH40:
		return PH
	case SubdivisionPH41:
		return PH
	case SubdivisionPHABR:
		return PH
	case SubdivisionPHAGN:
		return PH
	case SubdivisionPHAGS:
		return PH
	case SubdivisionPHAKL:
		return PH
	case SubdivisionPHALB:
		return PH
	case SubdivisionPHANT:
		return PH
	case SubdivisionPHAPA:
		return PH
	case SubdivisionPHAUR:
		return PH
	case SubdivisionPHBAN:
		return PH
	case SubdivisionPHBAS:
		return PH
	case SubdivisionPHBEN:
		return PH
	case SubdivisionPHBIL:
		return PH
	case SubdivisionPHBOH:
		return PH
	case SubdivisionPHBTG:
		return PH
	case SubdivisionPHBTN:
		return PH
	case SubdivisionPHBUK:
		return PH
	case SubdivisionPHBUL:
		return PH
	case SubdivisionPHCAG:
		return PH
	case SubdivisionPHCAM:
		return PH
	case SubdivisionPHCAN:
		return PH
	case SubdivisionPHCAP:
		return PH
	case SubdivisionPHCAS:
		return PH
	case SubdivisionPHCAT:
		return PH
	case SubdivisionPHCAV:
		return PH
	case SubdivisionPHCEB:
		return PH
	case SubdivisionPHCOM:
		return PH
	case SubdivisionPHDAO:
		return PH
	case SubdivisionPHDAS:
		return PH
	case SubdivisionPHDAV:
		return PH
	case SubdivisionPHDIN:
		return PH
	case SubdivisionPHEAS:
		return PH
	case SubdivisionPHGUI:
		return PH
	case SubdivisionPHIFU:
		return PH
	case SubdivisionPHILI:
		return PH
	case SubdivisionPHILN:
		return PH
	case SubdivisionPHILS:
		return PH
	case SubdivisionPHISA:
		return PH
	case SubdivisionPHKAL:
		return PH
	case SubdivisionPHLAG:
		return PH
	case SubdivisionPHLAN:
		return PH
	case SubdivisionPHLAS:
		return PH
	case SubdivisionPHLEY:
		return PH
	case SubdivisionPHLUN:
		return PH
	case SubdivisionPHMAD:
		return PH
	case SubdivisionPHMAG:
		return PH
	case SubdivisionPHMAS:
		return PH
	case SubdivisionPHMDC:
		return PH
	case SubdivisionPHMDR:
		return PH
	case SubdivisionPHMOU:
		return PH
	case SubdivisionPHMSC:
		return PH
	case SubdivisionPHMSR:
		return PH
	case SubdivisionPHNCO:
		return PH
	case SubdivisionPHNEC:
		return PH
	case SubdivisionPHNER:
		return PH
	case SubdivisionPHNSA:
		return PH
	case SubdivisionPHNUE:
		return PH
	case SubdivisionPHNUV:
		return PH
	case SubdivisionPHPAM:
		return PH
	case SubdivisionPHPAN:
		return PH
	case SubdivisionPHPLW:
		return PH
	case SubdivisionPHQUE:
		return PH
	case SubdivisionPHQUI:
		return PH
	case SubdivisionPHRIZ:
		return PH
	case SubdivisionPHROM:
		return PH
	case SubdivisionPHSAR:
		return PH
	case SubdivisionPHSCO:
		return PH
	case SubdivisionPHSIG:
		return PH
	case SubdivisionPHSLE:
		return PH
	case SubdivisionPHSLU:
		return PH
	case SubdivisionPHSOR:
		return PH
	case SubdivisionPHSUK:
		return PH
	case SubdivisionPHSUN:
		return PH
	case SubdivisionPHSUR:
		return PH
	case SubdivisionPHTAR:
		return PH
	case SubdivisionPHTAW:
		return PH
	case SubdivisionPHWSA:
		return PH
	case SubdivisionPHZAN:
		return PH
	case SubdivisionPHZAS:
		return PH
	case SubdivisionPHZMB:
		return PH
	case SubdivisionPHZSI:
		return PH
	case SubdivisionPKBA:
		return PK
	case SubdivisionPKGB:
		return PK
	case SubdivisionPKIS:
		return PK
	case SubdivisionPKJK:
		return PK
	case SubdivisionPKKP:
		return PK
	case SubdivisionPKPB:
		return PK
	case SubdivisionPKSD:
		return PK
	case SubdivisionPKTA:
		return PK
	case SubdivisionPLDS:
		return PL
	case SubdivisionPLKP:
		return PL
	case SubdivisionPLLB:
		return PL
	case SubdivisionPLLD:
		return PL
	case SubdivisionPLLU:
		return PL
	case SubdivisionPLMA:
		return PL
	case SubdivisionPLMZ:
		return PL
	case SubdivisionPLOP:
		return PL
	case SubdivisionPLPD:
		return PL
	case SubdivisionPLPK:
		return PL
	case SubdivisionPLPM:
		return PL
	case SubdivisionPLSK:
		return PL
	case SubdivisionPLSL:
		return PL
	case SubdivisionPLWN:
		return PL
	case SubdivisionPLWP:
		return PL
	case SubdivisionPLZP:
		return PL
	case SubdivisionPSBTH:
		return PS
	case SubdivisionPSDEB:
		return PS
	case SubdivisionPSGZA:
		return PS
	case SubdivisionPSHBN:
		return PS
	case SubdivisionPSJEM:
		return PS
	case SubdivisionPSJEN:
		return PS
	case SubdivisionPSJRH:
		return PS
	case SubdivisionPSKYS:
		return PS
	case SubdivisionPSNBS:
		return PS
	case SubdivisionPSNGZ:
		return PS
	case SubdivisionPSQQA:
		return PS
	case SubdivisionPSRBH:
		return PS
	case SubdivisionPSRFH:
		return PS
	case SubdivisionPSSLT:
		return PS
	case SubdivisionPSTBS:
		return PS
	case SubdivisionPSTKM:
		return PS
	case SubdivisionPT01:
		return PT
	case SubdivisionPT02:
		return PT
	case SubdivisionPT03:
		return PT
	case SubdivisionPT04:
		return PT
	case SubdivisionPT05:
		return PT
	case SubdivisionPT06:
		return PT
	case SubdivisionPT07:
		return PT
	case SubdivisionPT08:
		return PT
	case SubdivisionPT09:
		return PT
	case SubdivisionPT10:
		return PT
	case SubdivisionPT11:
		return PT
	case SubdivisionPT12:
		return PT
	case SubdivisionPT13:
		return PT
	case SubdivisionPT14:
		return PT
	case SubdivisionPT15:
		return PT
	case SubdivisionPT16:
		return PT
	case SubdivisionPT17:
		return PT
	case SubdivisionPT18:
		return PT
	case SubdivisionPT20:
		return PT
	case SubdivisionPT30:
		return PT
	case SubdivisionPW002:
		return PW
	case SubdivisionPW004:
		return PW
	case SubdivisionPW010:
		return PW
	case SubdivisionPW050:
		return PW
	case SubdivisionPW100:
		return PW
	case SubdivisionPW150:
		return PW
	case SubdivisionPW212:
		return PW
	case SubdivisionPW214:
		return PW
	case SubdivisionPW218:
		return PW
	case SubdivisionPW222:
		return PW
	case SubdivisionPW224:
		return PW
	case SubdivisionPW226:
		return PW
	case SubdivisionPW227:
		return PW
	case SubdivisionPW228:
		return PW
	case SubdivisionPW350:
		return PW
	case SubdivisionPW370:
		return PW
	case SubdivisionPY1:
		return PY
	case SubdivisionPY10:
		return PY
	case SubdivisionPY11:
		return PY
	case SubdivisionPY12:
		return PY
	case SubdivisionPY13:
		return PY
	case SubdivisionPY14:
		return PY
	case SubdivisionPY15:
		return PY
	case SubdivisionPY16:
		return PY
	case SubdivisionPY19:
		return PY
	case SubdivisionPY2:
		return PY
	case SubdivisionPY3:
		return PY
	case SubdivisionPY4:
		return PY
	case SubdivisionPY5:
		return PY
	case SubdivisionPY6:
		return PY
	case SubdivisionPY7:
		return PY
	case SubdivisionPY8:
		return PY
	case SubdivisionPY9:
		return PY
	case SubdivisionPYASU:
		return PY
	case SubdivisionQADA:
		return QA
	case SubdivisionQAKH:
		return QA
	case SubdivisionQAMS:
		return QA
	case SubdivisionQARA:
		return QA
	case SubdivisionQAUS:
		return QA
	case SubdivisionQAWA:
		return QA
	case SubdivisionQAZA:
		return QA
	case SubdivisionROAB:
		return RO
	case SubdivisionROAG:
		return RO
	case SubdivisionROAR:
		return RO
	case SubdivisionROB:
		return RO
	case SubdivisionROBC:
		return RO
	case SubdivisionROBH:
		return RO
	case SubdivisionROBN:
		return RO
	case SubdivisionROBR:
		return RO
	case SubdivisionROBT:
		return RO
	case SubdivisionROBV:
		return RO
	case SubdivisionROBZ:
		return RO
	case SubdivisionROCJ:
		return RO
	case SubdivisionROCL:
		return RO
	case SubdivisionROCS:
		return RO
	case SubdivisionROCT:
		return RO
	case SubdivisionROCV:
		return RO
	case SubdivisionRODB:
		return RO
	case SubdivisionRODJ:
		return RO
	case SubdivisionROGJ:
		return RO
	case SubdivisionROGL:
		return RO
	case SubdivisionROGR:
		return RO
	case SubdivisionROHD:
		return RO
	case SubdivisionROHR:
		return RO
	case SubdivisionROIF:
		return RO
	case SubdivisionROIL:
		return RO
	case SubdivisionROIS:
		return RO
	case SubdivisionROMH:
		return RO
	case SubdivisionROMM:
		return RO
	case SubdivisionROMS:
		return RO
	case SubdivisionRONT:
		return RO
	case SubdivisionROOT:
		return RO
	case SubdivisionROPH:
		return RO
	case SubdivisionROSB:
		return RO
	case SubdivisionROSJ:
		return RO
	case SubdivisionROSM:
		return RO
	case SubdivisionROSV:
		return RO
	case SubdivisionROTL:
		return RO
	case SubdivisionROTM:
		return RO
	case SubdivisionROTR:
		return RO
	case SubdivisionROVL:
		return RO
	case SubdivisionROVN:
		return RO
	case SubdivisionROVS:
		return RO
	case SubdivisionRS00:
		return RS
	case SubdivisionRS01:
		return RS
	case SubdivisionRS02:
		return RS
	case SubdivisionRS03:
		return RS
	case SubdivisionRS04:
		return RS
	case SubdivisionRS05:
		return RS
	case SubdivisionRS06:
		return RS
	case SubdivisionRS07:
		return RS
	case SubdivisionRS08:
		return RS
	case SubdivisionRS09:
		return RS
	case SubdivisionRS10:
		return RS
	case SubdivisionRS11:
		return RS
	case SubdivisionRS12:
		return RS
	case SubdivisionRS13:
		return RS
	case SubdivisionRS14:
		return RS
	case SubdivisionRS15:
		return RS
	case SubdivisionRS16:
		return RS
	case SubdivisionRS17:
		return RS
	case SubdivisionRS18:
		return RS
	case SubdivisionRS19:
		return RS
	case SubdivisionRS20:
		return RS
	case SubdivisionRS21:
		return RS
	case SubdivisionRS22:
		return RS
	case SubdivisionRS23:
		return RS
	case SubdivisionRS24:
		return RS
	case SubdivisionRS25:
		return RS
	case SubdivisionRS26:
		return RS
	case SubdivisionRS27:
		return RS
	case SubdivisionRS28:
		return RS
	case SubdivisionRS29:
		return RS
	case SubdivisionRSKM:
		return RS
	case SubdivisionRSVO:
		return RS
	case SubdivisionRUAD:
		return RU
	case SubdivisionRUAL:
		return RU
	case SubdivisionRUALT:
		return RU
	case SubdivisionRUAMU:
		return RU
	case SubdivisionRUARK:
		return RU
	case SubdivisionRUAST:
		return RU
	case SubdivisionRUBA:
		return RU
	case SubdivisionRUBEL:
		return RU
	case SubdivisionRUBRY:
		return RU
	case SubdivisionRUBU:
		return RU
	case SubdivisionRUCE:
		return RU
	case SubdivisionRUCHE:
		return RU
	case SubdivisionRUCHU:
		return RU
	case SubdivisionRUCU:
		return RU
	case SubdivisionRUDA:
		return RU
	case SubdivisionRUIN:
		return RU
	case SubdivisionRUIRK:
		return RU
	case SubdivisionRUIVA:
		return RU
	case SubdivisionRUKAM:
		return RU
	case SubdivisionRUKB:
		return RU
	case SubdivisionRUKC:
		return RU
	case SubdivisionRUKDA:
		return RU
	case SubdivisionRUKEM:
		return RU
	case SubdivisionRUKGD:
		return RU
	case SubdivisionRUKGN:
		return RU
	case SubdivisionRUKHA:
		return RU
	case SubdivisionRUKHM:
		return RU
	case SubdivisionRUKIR:
		return RU
	case SubdivisionRUKK:
		return RU
	case SubdivisionRUKL:
		return RU
	case SubdivisionRUKLU:
		return RU
	case SubdivisionRUKO:
		return RU
	case SubdivisionRUKOS:
		return RU
	case SubdivisionRUKR:
		return RU
	case SubdivisionRUKRS:
		return RU
	case SubdivisionRUKYA:
		return RU
	case SubdivisionRULEN:
		return RU
	case SubdivisionRULIP:
		return RU
	case SubdivisionRUMAG:
		return RU
	case SubdivisionRUME:
		return RU
	case SubdivisionRUMO:
		return RU
	case SubdivisionRUMOS:
		return RU
	case SubdivisionRUMOW:
		return RU
	case SubdivisionRUMUR:
		return RU
	case SubdivisionRUNEN:
		return RU
	case SubdivisionRUNGR:
		return RU
	case SubdivisionRUNIZ:
		return RU
	case SubdivisionRUNVS:
		return RU
	case SubdivisionRUOMS:
		return RU
	case SubdivisionRUORE:
		return RU
	case SubdivisionRUORL:
		return RU
	case SubdivisionRUPER:
		return RU
	case SubdivisionRUPNZ:
		return RU
	case SubdivisionRUPRI:
		return RU
	case SubdivisionRUPSK:
		return RU
	case SubdivisionRUROS:
		return RU
	case SubdivisionRURYA:
		return RU
	case SubdivisionRUSA:
		return RU
	case SubdivisionRUSAK:
		return RU
	case SubdivisionRUSAM:
		return RU
	case SubdivisionRUSAR:
		return RU
	case SubdivisionRUSE:
		return RU
	case SubdivisionRUSMO:
		return RU
	case SubdivisionRUSPE:
		return RU
	case SubdivisionRUSTA:
		return RU
	case SubdivisionRUSVE:
		return RU
	case SubdivisionRUTA:
		return RU
	case SubdivisionRUTAM:
		return RU
	case SubdivisionRUTOM:
		return RU
	case SubdivisionRUTUL:
		return RU
	case SubdivisionRUTVE:
		return RU
	case SubdivisionRUTY:
		return RU
	case SubdivisionRUTYU:
		return RU
	case SubdivisionRUUD:
		return RU
	case SubdivisionRUULY:
		return RU
	case SubdivisionRUVGG:
		return RU
	case SubdivisionRUVLA:
		return RU
	case SubdivisionRUVLG:
		return RU
	case SubdivisionRUVOR:
		return RU
	case SubdivisionRUYAN:
		return RU
	case SubdivisionRUYAR:
		return RU
	case SubdivisionRUYEV:
		return RU
	case SubdivisionRUZAB:
		return RU
	case SubdivisionRW01:
		return RW
	case SubdivisionRW02:
		return RW
	case SubdivisionRW03:
		return RW
	case SubdivisionRW04:
		return RW
	case SubdivisionRW05:
		return RW
	case SubdivisionSA01:
		return SA
	case SubdivisionSA02:
		return SA
	case SubdivisionSA03:
		return SA
	case SubdivisionSA04:
		return SA
	case SubdivisionSA05:
		return SA
	case SubdivisionSA06:
		return SA
	case SubdivisionSA07:
		return SA
	case SubdivisionSA08:
		return SA
	case SubdivisionSA09:
		return SA
	case SubdivisionSA10:
		return SA
	case SubdivisionSA11:
		return SA
	case SubdivisionSA12:
		return SA
	case SubdivisionSA14:
		return SA
	case SubdivisionSBCE:
		return SB
	case SubdivisionSBCH:
		return SB
	case SubdivisionSBCT:
		return SB
	case SubdivisionSBGU:
		return SB
	case SubdivisionSBIS:
		return SB
	case SubdivisionSBMK:
		return SB
	case SubdivisionSBML:
		return SB
	case SubdivisionSBRB:
		return SB
	case SubdivisionSBTE:
		return SB
	case SubdivisionSBWE:
		return SB
	case SubdivisionSC01:
		return SC
	case SubdivisionSC02:
		return SC
	case SubdivisionSC03:
		return SC
	case SubdivisionSC04:
		return SC
	case SubdivisionSC05:
		return SC
	case SubdivisionSC06:
		return SC
	case SubdivisionSC07:
		return SC
	case SubdivisionSC08:
		return SC
	case SubdivisionSC09:
		return SC
	case SubdivisionSC10:
		return SC
	case SubdivisionSC11:
		return SC
	case SubdivisionSC12:
		return SC
	case SubdivisionSC13:
		return SC
	case SubdivisionSC14:
		return SC
	case SubdivisionSC15:
		return SC
	case SubdivisionSC16:
		return SC
	case SubdivisionSC17:
		return SC
	case SubdivisionSC18:
		return SC
	case SubdivisionSC19:
		return SC
	case SubdivisionSC20:
		return SC
	case SubdivisionSC21:
		return SC
	case SubdivisionSC22:
		return SC
	case SubdivisionSC23:
		return SC
	case SubdivisionSC24:
		return SC
	case SubdivisionSC25:
		return SC
	case SubdivisionSDDC:
		return SD
	case SubdivisionSDDE:
		return SD
	case SubdivisionSDDN:
		return SD
	case SubdivisionSDDS:
		return SD
	case SubdivisionSDDW:
		return SD
	case SubdivisionSDGD:
		return SD
	case SubdivisionSDGZ:
		return SD
	case SubdivisionSDKA:
		return SD
	case SubdivisionSDKH:
		return SD
	case SubdivisionSDKN:
		return SD
	case SubdivisionSDKS:
		return SD
	case SubdivisionSDNB:
		return SD
	case SubdivisionSDNO:
		return SD
	case SubdivisionSDNR:
		return SD
	case SubdivisionSDNW:
		return SD
	case SubdivisionSDRS:
		return SD
	case SubdivisionSDSI:
		return SD
	case SubdivisionSEAB:
		return SE
	case SubdivisionSEAC:
		return SE
	case SubdivisionSEBD:
		return SE
	case SubdivisionSEC:
		return SE
	case SubdivisionSED:
		return SE
	case SubdivisionSEE:
		return SE
	case SubdivisionSEF:
		return SE
	case SubdivisionSEG:
		return SE
	case SubdivisionSEH:
		return SE
	case SubdivisionSEI:
		return SE
	case SubdivisionSEK:
		return SE
	case SubdivisionSEM:
		return SE
	case SubdivisionSEN:
		return SE
	case SubdivisionSEO:
		return SE
	case SubdivisionSES:
		return SE
	case SubdivisionSET:
		return SE
	case SubdivisionSEU:
		return SE
	case SubdivisionSEW:
		return SE
	case SubdivisionSEX:
		return SE
	case SubdivisionSEY:
		return SE
	case SubdivisionSEZ:
		return SE
	case SubdivisionSG01:
		return SG
	case SubdivisionSG02:
		return SG
	case SubdivisionSG03:
		return SG
	case SubdivisionSG04:
		return SG
	case SubdivisionSG05:
		return SG
	case SubdivisionSHAC:
		return SH
	case SubdivisionSHHL:
		return SH
	case SubdivisionSHTA:
		return SH
	case SubdivisionSI001:
		return SI
	case SubdivisionSI002:
		return SI
	case SubdivisionSI003:
		return SI
	case SubdivisionSI004:
		return SI
	case SubdivisionSI005:
		return SI
	case SubdivisionSI006:
		return SI
	case SubdivisionSI007:
		return SI
	case SubdivisionSI008:
		return SI
	case SubdivisionSI009:
		return SI
	case SubdivisionSI010:
		return SI
	case SubdivisionSI011:
		return SI
	case SubdivisionSI012:
		return SI
	case SubdivisionSI013:
		return SI
	case SubdivisionSI014:
		return SI
	case SubdivisionSI015:
		return SI
	case SubdivisionSI016:
		return SI
	case SubdivisionSI017:
		return SI
	case SubdivisionSI018:
		return SI
	case SubdivisionSI019:
		return SI
	case SubdivisionSI020:
		return SI
	case SubdivisionSI021:
		return SI
	case SubdivisionSI022:
		return SI
	case SubdivisionSI023:
		return SI
	case SubdivisionSI024:
		return SI
	case SubdivisionSI025:
		return SI
	case SubdivisionSI026:
		return SI
	case SubdivisionSI027:
		return SI
	case SubdivisionSI028:
		return SI
	case SubdivisionSI029:
		return SI
	case SubdivisionSI030:
		return SI
	case SubdivisionSI031:
		return SI
	case SubdivisionSI032:
		return SI
	case SubdivisionSI033:
		return SI
	case SubdivisionSI034:
		return SI
	case SubdivisionSI035:
		return SI
	case SubdivisionSI036:
		return SI
	case SubdivisionSI037:
		return SI
	case SubdivisionSI038:
		return SI
	case SubdivisionSI039:
		return SI
	case SubdivisionSI040:
		return SI
	case SubdivisionSI041:
		return SI
	case SubdivisionSI042:
		return SI
	case SubdivisionSI043:
		return SI
	case SubdivisionSI044:
		return SI
	case SubdivisionSI045:
		return SI
	case SubdivisionSI046:
		return SI
	case SubdivisionSI047:
		return SI
	case SubdivisionSI048:
		return SI
	case SubdivisionSI049:
		return SI
	case SubdivisionSI050:
		return SI
	case SubdivisionSI051:
		return SI
	case SubdivisionSI052:
		return SI
	case SubdivisionSI053:
		return SI
	case SubdivisionSI054:
		return SI
	case SubdivisionSI055:
		return SI
	case SubdivisionSI056:
		return SI
	case SubdivisionSI057:
		return SI
	case SubdivisionSI058:
		return SI
	case SubdivisionSI059:
		return SI
	case SubdivisionSI060:
		return SI
	case SubdivisionSI061:
		return SI
	case SubdivisionSI062:
		return SI
	case SubdivisionSI063:
		return SI
	case SubdivisionSI064:
		return SI
	case SubdivisionSI065:
		return SI
	case SubdivisionSI066:
		return SI
	case SubdivisionSI067:
		return SI
	case SubdivisionSI068:
		return SI
	case SubdivisionSI069:
		return SI
	case SubdivisionSI070:
		return SI
	case SubdivisionSI071:
		return SI
	case SubdivisionSI072:
		return SI
	case SubdivisionSI073:
		return SI
	case SubdivisionSI074:
		return SI
	case SubdivisionSI075:
		return SI
	case SubdivisionSI076:
		return SI
	case SubdivisionSI077:
		return SI
	case SubdivisionSI078:
		return SI
	case SubdivisionSI079:
		return SI
	case SubdivisionSI080:
		return SI
	case SubdivisionSI081:
		return SI
	case SubdivisionSI082:
		return SI
	case SubdivisionSI083:
		return SI
	case SubdivisionSI084:
		return SI
	case SubdivisionSI085:
		return SI
	case SubdivisionSI086:
		return SI
	case SubdivisionSI087:
		return SI
	case SubdivisionSI088:
		return SI
	case SubdivisionSI089:
		return SI
	case SubdivisionSI090:
		return SI
	case SubdivisionSI091:
		return SI
	case SubdivisionSI092:
		return SI
	case SubdivisionSI093:
		return SI
	case SubdivisionSI094:
		return SI
	case SubdivisionSI095:
		return SI
	case SubdivisionSI096:
		return SI
	case SubdivisionSI097:
		return SI
	case SubdivisionSI098:
		return SI
	case SubdivisionSI099:
		return SI
	case SubdivisionSI100:
		return SI
	case SubdivisionSI101:
		return SI
	case SubdivisionSI102:
		return SI
	case SubdivisionSI103:
		return SI
	case SubdivisionSI104:
		return SI
	case SubdivisionSI105:
		return SI
	case SubdivisionSI106:
		return SI
	case SubdivisionSI107:
		return SI
	case SubdivisionSI108:
		return SI
	case SubdivisionSI109:
		return SI
	case SubdivisionSI110:
		return SI
	case SubdivisionSI111:
		return SI
	case SubdivisionSI112:
		return SI
	case SubdivisionSI113:
		return SI
	case SubdivisionSI114:
		return SI
	case SubdivisionSI115:
		return SI
	case SubdivisionSI116:
		return SI
	case SubdivisionSI117:
		return SI
	case SubdivisionSI118:
		return SI
	case SubdivisionSI119:
		return SI
	case SubdivisionSI120:
		return SI
	case SubdivisionSI121:
		return SI
	case SubdivisionSI122:
		return SI
	case SubdivisionSI123:
		return SI
	case SubdivisionSI124:
		return SI
	case SubdivisionSI125:
		return SI
	case SubdivisionSI126:
		return SI
	case SubdivisionSI127:
		return SI
	case SubdivisionSI128:
		return SI
	case SubdivisionSI129:
		return SI
	case SubdivisionSI130:
		return SI
	case SubdivisionSI131:
		return SI
	case SubdivisionSI132:
		return SI
	case SubdivisionSI133:
		return SI
	case SubdivisionSI134:
		return SI
	case SubdivisionSI135:
		return SI
	case SubdivisionSI136:
		return SI
	case SubdivisionSI137:
		return SI
	case SubdivisionSI138:
		return SI
	case SubdivisionSI139:
		return SI
	case SubdivisionSI140:
		return SI
	case SubdivisionSI141:
		return SI
	case SubdivisionSI142:
		return SI
	case SubdivisionSI143:
		return SI
	case SubdivisionSI144:
		return SI
	case SubdivisionSI146:
		return SI
	case SubdivisionSI147:
		return SI
	case SubdivisionSI148:
		return SI
	case SubdivisionSI149:
		return SI
	case SubdivisionSI150:
		return SI
	case SubdivisionSI151:
		return SI
	case SubdivisionSI152:
		return SI
	case SubdivisionSI153:
		return SI
	case SubdivisionSI154:
		return SI
	case SubdivisionSI155:
		return SI
	case SubdivisionSI156:
		return SI
	case SubdivisionSI157:
		return SI
	case SubdivisionSI158:
		return SI
	case SubdivisionSI159:
		return SI
	case SubdivisionSI160:
		return SI
	case SubdivisionSI161:
		return SI
	case SubdivisionSI162:
		return SI
	case SubdivisionSI163:
		return SI
	case SubdivisionSI164:
		return SI
	case SubdivisionSI165:
		return SI
	case SubdivisionSI166:
		return SI
	case SubdivisionSI167:
		return SI
	case SubdivisionSI168:
		return SI
	case SubdivisionSI169:
		return SI
	case SubdivisionSI170:
		return SI
	case SubdivisionSI171:
		return SI
	case SubdivisionSI172:
		return SI
	case SubdivisionSI173:
		return SI
	case SubdivisionSI174:
		return SI
	case SubdivisionSI175:
		return SI
	case SubdivisionSI176:
		return SI
	case SubdivisionSI177:
		return SI
	case SubdivisionSI178:
		return SI
	case SubdivisionSI179:
		return SI
	case SubdivisionSI180:
		return SI
	case SubdivisionSI181:
		return SI
	case SubdivisionSI182:
		return SI
	case SubdivisionSI183:
		return SI
	case SubdivisionSI184:
		return SI
	case SubdivisionSI185:
		return SI
	case SubdivisionSI186:
		return SI
	case SubdivisionSI187:
		return SI
	case SubdivisionSI188:
		return SI
	case SubdivisionSI189:
		return SI
	case SubdivisionSI190:
		return SI
	case SubdivisionSI191:
		return SI
	case SubdivisionSI192:
		return SI
	case SubdivisionSI193:
		return SI
	case SubdivisionSI194:
		return SI
	case SubdivisionSI195:
		return SI
	case SubdivisionSI196:
		return SI
	case SubdivisionSI197:
		return SI
	case SubdivisionSI198:
		return SI
	case SubdivisionSI199:
		return SI
	case SubdivisionSI200:
		return SI
	case SubdivisionSI201:
		return SI
	case SubdivisionSI202:
		return SI
	case SubdivisionSI203:
		return SI
	case SubdivisionSI204:
		return SI
	case SubdivisionSI205:
		return SI
	case SubdivisionSI206:
		return SI
	case SubdivisionSI207:
		return SI
	case SubdivisionSI208:
		return SI
	case SubdivisionSI209:
		return SI
	case SubdivisionSI210:
		return SI
	case SubdivisionSI211:
		return SI
	case SubdivisionSKBC:
		return SK
	case SubdivisionSKBL:
		return SK
	case SubdivisionSKKI:
		return SK
	case SubdivisionSKNI:
		return SK
	case SubdivisionSKPV:
		return SK
	case SubdivisionSKTA:
		return SK
	case SubdivisionSKTC:
		return SK
	case SubdivisionSKZI:
		return SK
	case SubdivisionSLE:
		return SL
	case SubdivisionSLN:
		return SL
	case SubdivisionSLS:
		return SL
	case SubdivisionSLW:
		return SL
	case SubdivisionSM01:
		return SM
	case SubdivisionSM02:
		return SM
	case SubdivisionSM03:
		return SM
	case SubdivisionSM04:
		return SM
	case SubdivisionSM05:
		return SM
	case SubdivisionSM06:
		return SM
	case SubdivisionSM07:
		return SM
	case SubdivisionSM08:
		return SM
	case SubdivisionSM09:
		return SM
	case SubdivisionSNDB:
		return SN
	case SubdivisionSNDK:
		return SN
	case SubdivisionSNFK:
		return SN
	case SubdivisionSNKA:
		return SN
	case SubdivisionSNKD:
		return SN
	case SubdivisionSNKE:
		return SN
	case SubdivisionSNKL:
		return SN
	case SubdivisionSNLG:
		return SN
	case SubdivisionSNMT:
		return SN
	case SubdivisionSNSE:
		return SN
	case SubdivisionSNSL:
		return SN
	case SubdivisionSNTC:
		return SN
	case SubdivisionSNTH:
		return SN
	case SubdivisionSNZG:
		return SN
	case SubdivisionSOAW:
		return SO
	case SubdivisionSOBK:
		return SO
	case SubdivisionSOBN:
		return SO
	case SubdivisionSOBR:
		return SO
	case SubdivisionSOBY:
		return SO
	case SubdivisionSOGA:
		return SO
	case SubdivisionSOGE:
		return SO
	case SubdivisionSOHI:
		return SO
	case SubdivisionSOJD:
		return SO
	case SubdivisionSOJH:
		return SO
	case SubdivisionSOMU:
		return SO
	case SubdivisionSONU:
		return SO
	case SubdivisionSOSA:
		return SO
	case SubdivisionSOSD:
		return SO
	case SubdivisionSOSH:
		return SO
	case SubdivisionSOSO:
		return SO
	case SubdivisionSOTO:
		return SO
	case SubdivisionSOWO:
		return SO
	case SubdivisionSRBR:
		return SR
	case SubdivisionSRCM:
		return SR
	case SubdivisionSRCR:
		return SR
	case SubdivisionSRMA:
		return SR
	case SubdivisionSRNI:
		return SR
	case SubdivisionSRPM:
		return SR
	case SubdivisionSRPR:
		return SR
	case SubdivisionSRSA:
		return SR
	case SubdivisionSRSI:
		return SR
	case SubdivisionSRWA:
		return SR
	case SubdivisionSSBN:
		return SS
	case SubdivisionSSBW:
		return SS
	case SubdivisionSSEC:
		return SS
	case SubdivisionSSEE:
		return SS
	case SubdivisionSSEW:
		return SS
	case SubdivisionSSJG:
		return SS
	case SubdivisionSSLK:
		return SS
	case SubdivisionSSNU:
		return SS
	case SubdivisionSSUY:
		return SS
	case SubdivisionSSWR:
		return SS
	case SubdivisionSTP:
		return ST
	case SubdivisionSTS:
		return ST
	case SubdivisionSVAH:
		return SV
	case SubdivisionSVCA:
		return SV
	case SubdivisionSVCH:
		return SV
	case SubdivisionSVCU:
		return SV
	case SubdivisionSVLI:
		return SV
	case SubdivisionSVMO:
		return SV
	case SubdivisionSVPA:
		return SV
	case SubdivisionSVSA:
		return SV
	case SubdivisionSVSM:
		return SV
	case SubdivisionSVSO:
		return SV
	case SubdivisionSVSS:
		return SV
	case SubdivisionSVSV:
		return SV
	case SubdivisionSVUN:
		return SV
	case SubdivisionSVUS:
		return SV
	case SubdivisionSYDI:
		return SY
	case SubdivisionSYDR:
		return SY
	case SubdivisionSYDY:
		return SY
	case SubdivisionSYHA:
		return SY
	case SubdivisionSYHI:
		return SY
	case SubdivisionSYHL:
		return SY
	case SubdivisionSYHM:
		return SY
	case SubdivisionSYID:
		return SY
	case SubdivisionSYLA:
		return SY
	case SubdivisionSYQU:
		return SY
	case SubdivisionSYRA:
		return SY
	case SubdivisionSYRD:
		return SY
	case SubdivisionSYSU:
		return SY
	case SubdivisionSYTA:
		return SY
	case SubdivisionSZHH:
		return SZ
	case SubdivisionSZLU:
		return SZ
	case SubdivisionSZMA:
		return SZ
	case SubdivisionSZSH:
		return SZ
	case SubdivisionTDBA:
		return TD
	case SubdivisionTDBG:
		return TD
	case SubdivisionTDBO:
		return TD
	case SubdivisionTDCB:
		return TD
	case SubdivisionTDEN:
		return TD
	case SubdivisionTDGR:
		return TD
	case SubdivisionTDHL:
		return TD
	case SubdivisionTDKA:
		return TD
	case SubdivisionTDLC:
		return TD
	case SubdivisionTDLO:
		return TD
	case SubdivisionTDLR:
		return TD
	case SubdivisionTDMA:
		return TD
	case SubdivisionTDMC:
		return TD
	case SubdivisionTDME:
		return TD
	case SubdivisionTDMO:
		return TD
	case SubdivisionTDND:
		return TD
	case SubdivisionTDOD:
		return TD
	case SubdivisionTDSA:
		return TD
	case SubdivisionTDSI:
		return TD
	case SubdivisionTDTA:
		return TD
	case SubdivisionTDTI:
		return TD
	case SubdivisionTDWF:
		return TD
	case SubdivisionTGC:
		return TG
	case SubdivisionTGK:
		return TG
	case SubdivisionTGM:
		return TG
	case SubdivisionTGP:
		return TG
	case SubdivisionTGS:
		return TG
	case SubdivisionTH10:
		return TH
	case SubdivisionTH11:
		return TH
	case SubdivisionTH12:
		return TH
	case SubdivisionTH13:
		return TH
	case SubdivisionTH14:
		return TH
	case SubdivisionTH15:
		return TH
	case SubdivisionTH16:
		return TH
	case SubdivisionTH17:
		return TH
	case SubdivisionTH18:
		return TH
	case SubdivisionTH19:
		return TH
	case SubdivisionTH20:
		return TH
	case SubdivisionTH21:
		return TH
	case SubdivisionTH22:
		return TH
	case SubdivisionTH23:
		return TH
	case SubdivisionTH24:
		return TH
	case SubdivisionTH25:
		return TH
	case SubdivisionTH26:
		return TH
	case SubdivisionTH27:
		return TH
	case SubdivisionTH30:
		return TH
	case SubdivisionTH31:
		return TH
	case SubdivisionTH32:
		return TH
	case SubdivisionTH33:
		return TH
	case SubdivisionTH34:
		return TH
	case SubdivisionTH35:
		return TH
	case SubdivisionTH36:
		return TH
	case SubdivisionTH37:
		return TH
	case SubdivisionTH39:
		return TH
	case SubdivisionTH40:
		return TH
	case SubdivisionTH41:
		return TH
	case SubdivisionTH42:
		return TH
	case SubdivisionTH43:
		return TH
	case SubdivisionTH44:
		return TH
	case SubdivisionTH45:
		return TH
	case SubdivisionTH46:
		return TH
	case SubdivisionTH47:
		return TH
	case SubdivisionTH48:
		return TH
	case SubdivisionTH49:
		return TH
	case SubdivisionTH50:
		return TH
	case SubdivisionTH51:
		return TH
	case SubdivisionTH52:
		return TH
	case SubdivisionTH53:
		return TH
	case SubdivisionTH54:
		return TH
	case SubdivisionTH55:
		return TH
	case SubdivisionTH56:
		return TH
	case SubdivisionTH57:
		return TH
	case SubdivisionTH58:
		return TH
	case SubdivisionTH60:
		return TH
	case SubdivisionTH61:
		return TH
	case SubdivisionTH62:
		return TH
	case SubdivisionTH63:
		return TH
	case SubdivisionTH64:
		return TH
	case SubdivisionTH65:
		return TH
	case SubdivisionTH66:
		return TH
	case SubdivisionTH67:
		return TH
	case SubdivisionTH70:
		return TH
	case SubdivisionTH71:
		return TH
	case SubdivisionTH72:
		return TH
	case SubdivisionTH73:
		return TH
	case SubdivisionTH74:
		return TH
	case SubdivisionTH75:
		return TH
	case SubdivisionTH76:
		return TH
	case SubdivisionTH77:
		return TH
	case SubdivisionTH80:
		return TH
	case SubdivisionTH81:
		return TH
	case SubdivisionTH82:
		return TH
	case SubdivisionTH83:
		return TH
	case SubdivisionTH84:
		return TH
	case SubdivisionTH85:
		return TH
	case SubdivisionTH86:
		return TH
	case SubdivisionTH90:
		return TH
	case SubdivisionTH91:
		return TH
	case SubdivisionTH92:
		return TH
	case SubdivisionTH93:
		return TH
	case SubdivisionTH94:
		return TH
	case SubdivisionTH95:
		return TH
	case SubdivisionTH96:
		return TH
	case SubdivisionTHS:
		return TH
	case SubdivisionTJGB:
		return TJ
	case SubdivisionTJKT:
		return TJ
	case SubdivisionTJSU:
		return TJ
	case SubdivisionTLAL:
		return TL
	case SubdivisionTLAN:
		return TL
	case SubdivisionTLBA:
		return TL
	case SubdivisionTLBO:
		return TL
	case SubdivisionTLCO:
		return TL
	case SubdivisionTLDI:
		return TL
	case SubdivisionTLER:
		return TL
	case SubdivisionTLLA:
		return TL
	case SubdivisionTLLI:
		return TL
	case SubdivisionTLMF:
		return TL
	case SubdivisionTLMT:
		return TL
	case SubdivisionTLOE:
		return TL
	case SubdivisionTLVI:
		return TL
	case SubdivisionTMA:
		return TM
	case SubdivisionTMB:
		return TM
	case SubdivisionTMD:
		return TM
	case SubdivisionTML:
		return TM
	case SubdivisionTMM:
		return TM
	case SubdivisionTMS:
		return TM
	case SubdivisionTN11:
		return TN
	case SubdivisionTN12:
		return TN
	case SubdivisionTN13:
		return TN
	case SubdivisionTN14:
		return TN
	case SubdivisionTN21:
		return TN
	case SubdivisionTN22:
		return TN
	case SubdivisionTN23:
		return TN
	case SubdivisionTN31:
		return TN
	case SubdivisionTN32:
		return TN
	case SubdivisionTN33:
		return TN
	case SubdivisionTN34:
		return TN
	case SubdivisionTN41:
		return TN
	case SubdivisionTN42:
		return TN
	case SubdivisionTN43:
		return TN
	case SubdivisionTN51:
		return TN
	case SubdivisionTN52:
		return TN
	case SubdivisionTN53:
		return TN
	case SubdivisionTN61:
		return TN
	case SubdivisionTN71:
		return TN
	case SubdivisionTN72:
		return TN
	case SubdivisionTN73:
		return TN
	case SubdivisionTN81:
		return TN
	case SubdivisionTN82:
		return TN
	case SubdivisionTN83:
		return TN
	case SubdivisionTO01:
		return TO
	case SubdivisionTO02:
		return TO
	case SubdivisionTO03:
		return TO
	case SubdivisionTO04:
		return TO
	case SubdivisionTO05:
		return TO
	case SubdivisionTR01:
		return TR
	case SubdivisionTR02:
		return TR
	case SubdivisionTR03:
		return TR
	case SubdivisionTR04:
		return TR
	case SubdivisionTR05:
		return TR
	case SubdivisionTR06:
		return TR
	case SubdivisionTR07:
		return TR
	case SubdivisionTR08:
		return TR
	case SubdivisionTR09:
		return TR
	case SubdivisionTR10:
		return TR
	case SubdivisionTR11:
		return TR
	case SubdivisionTR12:
		return TR
	case SubdivisionTR13:
		return TR
	case SubdivisionTR14:
		return TR
	case SubdivisionTR15:
		return TR
	case SubdivisionTR16:
		return TR
	case SubdivisionTR17:
		return TR
	case SubdivisionTR18:
		return TR
	case SubdivisionTR19:
		return TR
	case SubdivisionTR20:
		return TR
	case SubdivisionTR21:
		return TR
	case SubdivisionTR22:
		return TR
	case SubdivisionTR23:
		return TR
	case SubdivisionTR24:
		return TR
	case SubdivisionTR25:
		return TR
	case SubdivisionTR26:
		return TR
	case SubdivisionTR27:
		return TR
	case SubdivisionTR28:
		return TR
	case SubdivisionTR29:
		return TR
	case SubdivisionTR30:
		return TR
	case SubdivisionTR31:
		return TR
	case SubdivisionTR32:
		return TR
	case SubdivisionTR33:
		return TR
	case SubdivisionTR34:
		return TR
	case SubdivisionTR35:
		return TR
	case SubdivisionTR36:
		return TR
	case SubdivisionTR37:
		return TR
	case SubdivisionTR38:
		return TR
	case SubdivisionTR39:
		return TR
	case SubdivisionTR40:
		return TR
	case SubdivisionTR41:
		return TR
	case SubdivisionTR42:
		return TR
	case SubdivisionTR43:
		return TR
	case SubdivisionTR44:
		return TR
	case SubdivisionTR45:
		return TR
	case SubdivisionTR46:
		return TR
	case SubdivisionTR47:
		return TR
	case SubdivisionTR48:
		return TR
	case SubdivisionTR49:
		return TR
	case SubdivisionTR50:
		return TR
	case SubdivisionTR51:
		return TR
	case SubdivisionTR52:
		return TR
	case SubdivisionTR53:
		return TR
	case SubdivisionTR54:
		return TR
	case SubdivisionTR55:
		return TR
	case SubdivisionTR56:
		return TR
	case SubdivisionTR57:
		return TR
	case SubdivisionTR58:
		return TR
	case SubdivisionTR59:
		return TR
	case SubdivisionTR60:
		return TR
	case SubdivisionTR61:
		return TR
	case SubdivisionTR62:
		return TR
	case SubdivisionTR63:
		return TR
	case SubdivisionTR64:
		return TR
	case SubdivisionTR65:
		return TR
	case SubdivisionTR66:
		return TR
	case SubdivisionTR67:
		return TR
	case SubdivisionTR68:
		return TR
	case SubdivisionTR69:
		return TR
	case SubdivisionTR70:
		return TR
	case SubdivisionTR71:
		return TR
	case SubdivisionTR72:
		return TR
	case SubdivisionTR73:
		return TR
	case SubdivisionTR74:
		return TR
	case SubdivisionTR75:
		return TR
	case SubdivisionTR76:
		return TR
	case SubdivisionTR77:
		return TR
	case SubdivisionTR78:
		return TR
	case SubdivisionTR79:
		return TR
	case SubdivisionTR80:
		return TR
	case SubdivisionTR81:
		return TR
	case SubdivisionTTARI:
		return TT
	case SubdivisionTTCHA:
		return TT
	case SubdivisionTTCTT:
		return TT
	case SubdivisionTTDMN:
		return TT
	case SubdivisionTTETO:
		return TT
	case SubdivisionTTPED:
		return TT
	case SubdivisionTTPOS:
		return TT
	case SubdivisionTTPRT:
		return TT
	case SubdivisionTTPTF:
		return TT
	case SubdivisionTTRCM:
		return TT
	case SubdivisionTTSFO:
		return TT
	case SubdivisionTTSGE:
		return TT
	case SubdivisionTTSIP:
		return TT
	case SubdivisionTTSJL:
		return TT
	case SubdivisionTTTUP:
		return TT
	case SubdivisionTTWTO:
		return TT
	case SubdivisionTVFUN:
		return TV
	case SubdivisionTVNIT:
		return TV
	case SubdivisionTVNKF:
		return TV
	case SubdivisionTVNKL:
		return TV
	case SubdivisionTVNMA:
		return TV
	case SubdivisionTVNMG:
		return TV
	case SubdivisionTVNUI:
		return TV
	case SubdivisionTVVAI:
		return TV
	case SubdivisionTWCHA:
		return TW
	case SubdivisionTWCYI:
		return TW
	case SubdivisionTWCYQ:
		return TW
	case SubdivisionTWHSQ:
		return TW
	case SubdivisionTWHSZ:
		return TW
	case SubdivisionTWHUA:
		return TW
	case SubdivisionTWILA:
		return TW
	case SubdivisionTWKEE:
		return TW
	case SubdivisionTWKHH:
		return TW
	case SubdivisionTWKHQ:
		return TW
	case SubdivisionTWMIA:
		return TW
	case SubdivisionTWNAN:
		return TW
	case SubdivisionTWPEN:
		return TW
	case SubdivisionTWPIF:
		return TW
	case SubdivisionTWTAO:
		return TW
	case SubdivisionTWTNN:
		return TW
	case SubdivisionTWTNQ:
		return TW
	case SubdivisionTWTPE:
		return TW
	case SubdivisionTWTPQ:
		return TW
	case SubdivisionTWTTT:
		return TW
	case SubdivisionTWTXG:
		return TW
	case SubdivisionTWTXQ:
		return TW
	case SubdivisionTWYUN:
		return TW
	case SubdivisionTZ01:
		return TZ
	case SubdivisionTZ02:
		return TZ
	case SubdivisionTZ03:
		return TZ
	case SubdivisionTZ04:
		return TZ
	case SubdivisionTZ05:
		return TZ
	case SubdivisionTZ06:
		return TZ
	case SubdivisionTZ07:
		return TZ
	case SubdivisionTZ08:
		return TZ
	case SubdivisionTZ09:
		return TZ
	case SubdivisionTZ10:
		return TZ
	case SubdivisionTZ11:
		return TZ
	case SubdivisionTZ12:
		return TZ
	case SubdivisionTZ13:
		return TZ
	case SubdivisionTZ14:
		return TZ
	case SubdivisionTZ15:
		return TZ
	case SubdivisionTZ16:
		return TZ
	case SubdivisionTZ17:
		return TZ
	case SubdivisionTZ18:
		return TZ
	case SubdivisionTZ19:
		return TZ
	case SubdivisionTZ20:
		return TZ
	case SubdivisionTZ21:
		return TZ
	case SubdivisionTZ22:
		return TZ
	case SubdivisionTZ23:
		return TZ
	case SubdivisionTZ24:
		return TZ
	case SubdivisionTZ25:
		return TZ
	case SubdivisionTZ26:
		return TZ
	case SubdivisionUA05:
		return UA
	case SubdivisionUA07:
		return UA
	case SubdivisionUA09:
		return UA
	case SubdivisionUA12:
		return UA
	case SubdivisionUA14:
		return UA
	case SubdivisionUA18:
		return UA
	case SubdivisionUA21:
		return UA
	case SubdivisionUA23:
		return UA
	case SubdivisionUA26:
		return UA
	case SubdivisionUA30:
		return UA
	case SubdivisionUA32:
		return UA
	case SubdivisionUA35:
		return UA
	case SubdivisionUA40:
		return UA
	case SubdivisionUA43:
		return UA
	case SubdivisionUA46:
		return UA
	case SubdivisionUA48:
		return UA
	case SubdivisionUA51:
		return UA
	case SubdivisionUA53:
		return UA
	case SubdivisionUA56:
		return UA
	case SubdivisionUA59:
		return UA
	case SubdivisionUA61:
		return UA
	case SubdivisionUA63:
		return UA
	case SubdivisionUA65:
		return UA
	case SubdivisionUA68:
		return UA
	case SubdivisionUA71:
		return UA
	case SubdivisionUA74:
		return UA
	case SubdivisionUA77:
		return UA
	case SubdivisionUG101:
		return UG
	case SubdivisionUG102:
		return UG
	case SubdivisionUG103:
		return UG
	case SubdivisionUG104:
		return UG
	case SubdivisionUG105:
		return UG
	case SubdivisionUG106:
		return UG
	case SubdivisionUG107:
		return UG
	case SubdivisionUG108:
		return UG
	case SubdivisionUG109:
		return UG
	case SubdivisionUG110:
		return UG
	case SubdivisionUG111:
		return UG
	case SubdivisionUG112:
		return UG
	case SubdivisionUG113:
		return UG
	case SubdivisionUG114:
		return UG
	case SubdivisionUG115:
		return UG
	case SubdivisionUG116:
		return UG
	case SubdivisionUG201:
		return UG
	case SubdivisionUG202:
		return UG
	case SubdivisionUG203:
		return UG
	case SubdivisionUG204:
		return UG
	case SubdivisionUG205:
		return UG
	case SubdivisionUG206:
		return UG
	case SubdivisionUG207:
		return UG
	case SubdivisionUG208:
		return UG
	case SubdivisionUG209:
		return UG
	case SubdivisionUG210:
		return UG
	case SubdivisionUG211:
		return UG
	case SubdivisionUG212:
		return UG
	case SubdivisionUG213:
		return UG
	case SubdivisionUG214:
		return UG
	case SubdivisionUG215:
		return UG
	case SubdivisionUG216:
		return UG
	case SubdivisionUG217:
		return UG
	case SubdivisionUG218:
		return UG
	case SubdivisionUG219:
		return UG
	case SubdivisionUG220:
		return UG
	case SubdivisionUG221:
		return UG
	case SubdivisionUG222:
		return UG
	case SubdivisionUG223:
		return UG
	case SubdivisionUG224:
		return UG
	case SubdivisionUG301:
		return UG
	case SubdivisionUG302:
		return UG
	case SubdivisionUG303:
		return UG
	case SubdivisionUG304:
		return UG
	case SubdivisionUG305:
		return UG
	case SubdivisionUG306:
		return UG
	case SubdivisionUG307:
		return UG
	case SubdivisionUG308:
		return UG
	case SubdivisionUG309:
		return UG
	case SubdivisionUG310:
		return UG
	case SubdivisionUG311:
		return UG
	case SubdivisionUG312:
		return UG
	case SubdivisionUG313:
		return UG
	case SubdivisionUG314:
		return UG
	case SubdivisionUG315:
		return UG
	case SubdivisionUG316:
		return UG
	case SubdivisionUG317:
		return UG
	case SubdivisionUG318:
		return UG
	case SubdivisionUG319:
		return UG
	case SubdivisionUG320:
		return UG
	case SubdivisionUG321:
		return UG
	case SubdivisionUG401:
		return UG
	case SubdivisionUG402:
		return UG
	case SubdivisionUG403:
		return UG
	case SubdivisionUG404:
		return UG
	case SubdivisionUG405:
		return UG
	case SubdivisionUG406:
		return UG
	case SubdivisionUG407:
		return UG
	case SubdivisionUG408:
		return UG
	case SubdivisionUG409:
		return UG
	case SubdivisionUG410:
		return UG
	case SubdivisionUG411:
		return UG
	case SubdivisionUG412:
		return UG
	case SubdivisionUG413:
		return UG
	case SubdivisionUG414:
		return UG
	case SubdivisionUG415:
		return UG
	case SubdivisionUG416:
		return UG
	case SubdivisionUG417:
		return UG
	case SubdivisionUG418:
		return UG
	case SubdivisionUG419:
		return UG
	case SubdivisionUGC:
		return UG
	case SubdivisionUGE:
		return UG
	case SubdivisionUGN:
		return UG
	case SubdivisionUGW:
		return UG
	case SubdivisionUM67:
		return UM
	case SubdivisionUM71:
		return UM
	case SubdivisionUM76:
		return UM
	case SubdivisionUM79:
		return UM
	case SubdivisionUM81:
		return UM
	case SubdivisionUM84:
		return UM
	case SubdivisionUM86:
		return UM
	case SubdivisionUM89:
		return UM
	case SubdivisionUM95:
		return UM
	case SubdivisionUSAK:
		return US
	case SubdivisionUSAL:
		return US
	case SubdivisionUSAR:
		return US
	case SubdivisionUSAS:
		return US
	case SubdivisionUSAZ:
		return US
	case SubdivisionUSCA:
		return US
	case SubdivisionUSCO:
		return US
	case SubdivisionUSCT:
		return US
	case SubdivisionUSDC:
		return US
	case SubdivisionUSDE:
		return US
	case SubdivisionUSFL:
		return US
	case SubdivisionUSGA:
		return US
	case SubdivisionUSGU:
		return US
	case SubdivisionUSHI:
		return US
	case SubdivisionUSIA:
		return US
	case SubdivisionUSID:
		return US
	case SubdivisionUSIL:
		return US
	case SubdivisionUSIN:
		return US
	case SubdivisionUSKS:
		return US
	case SubdivisionUSKY:
		return US
	case SubdivisionUSLA:
		return US
	case SubdivisionUSMA:
		return US
	case SubdivisionUSMD:
		return US
	case SubdivisionUSME:
		return US
	case SubdivisionUSMI:
		return US
	case SubdivisionUSMN:
		return US
	case SubdivisionUSMO:
		return US
	case SubdivisionUSMP:
		return US
	case SubdivisionUSMS:
		return US
	case SubdivisionUSMT:
		return US
	case SubdivisionUSNC:
		return US
	case SubdivisionUSND:
		return US
	case SubdivisionUSNE:
		return US
	case SubdivisionUSNH:
		return US
	case SubdivisionUSNJ:
		return US
	case SubdivisionUSNM:
		return US
	case SubdivisionUSNV:
		return US
	case SubdivisionUSNY:
		return US
	case SubdivisionUSOH:
		return US
	case SubdivisionUSOK:
		return US
	case SubdivisionUSOR:
		return US
	case SubdivisionUSPA:
		return US
	case SubdivisionUSPR:
		return US
	case SubdivisionUSRI:
		return US
	case SubdivisionUSSC:
		return US
	case SubdivisionUSSD:
		return US
	case SubdivisionUSTN:
		return US
	case SubdivisionUSTX:
		return US
	case SubdivisionUSUM:
		return US
	case SubdivisionUSUT:
		return US
	case SubdivisionUSVA:
		return US
	case SubdivisionUSVI:
		return US
	case SubdivisionUSVT:
		return US
	case SubdivisionUSWA:
		return US
	case SubdivisionUSWI:
		return US
	case SubdivisionUSWV:
		return US
	case SubdivisionUSWY:
		return US
	case SubdivisionUYAR:
		return UY
	case SubdivisionUYCA:
		return UY
	case SubdivisionUYCL:
		return UY
	case SubdivisionUYCO:
		return UY
	case SubdivisionUYDU:
		return UY
	case SubdivisionUYFD:
		return UY
	case SubdivisionUYFS:
		return UY
	case SubdivisionUYLA:
		return UY
	case SubdivisionUYMA:
		return UY
	case SubdivisionUYMO:
		return UY
	case SubdivisionUYPA:
		return UY
	case SubdivisionUYRN:
		return UY
	case SubdivisionUYRO:
		return UY
	case SubdivisionUYRV:
		return UY
	case SubdivisionUYSA:
		return UY
	case SubdivisionUYSJ:
		return UY
	case SubdivisionUYSO:
		return UY
	case SubdivisionUYTA:
		return UY
	case SubdivisionUYTT:
		return UY
	case SubdivisionUZAN:
		return UZ
	case SubdivisionUZBU:
		return UZ
	case SubdivisionUZFA:
		return UZ
	case SubdivisionUZJI:
		return UZ
	case SubdivisionUZNG:
		return UZ
	case SubdivisionUZNW:
		return UZ
	case SubdivisionUZQA:
		return UZ
	case SubdivisionUZQR:
		return UZ
	case SubdivisionUZSA:
		return UZ
	case SubdivisionUZSI:
		return UZ
	case SubdivisionUZSU:
		return UZ
	case SubdivisionUZTK:
		return UZ
	case SubdivisionUZTO:
		return UZ
	case SubdivisionUZXO:
		return UZ
	case SubdivisionVC01:
		return VC
	case SubdivisionVC02:
		return VC
	case SubdivisionVC03:
		return VC
	case SubdivisionVC04:
		return VC
	case SubdivisionVC05:
		return VC
	case SubdivisionVC06:
		return VC
	case SubdivisionVEA:
		return VE
	case SubdivisionVEB:
		return VE
	case SubdivisionVEC:
		return VE
	case SubdivisionVED:
		return VE
	case SubdivisionVEE:
		return VE
	case SubdivisionVEF:
		return VE
	case SubdivisionVEG:
		return VE
	case SubdivisionVEH:
		return VE
	case SubdivisionVEI:
		return VE
	case SubdivisionVEJ:
		return VE
	case SubdivisionVEK:
		return VE
	case SubdivisionVEL:
		return VE
	case SubdivisionVEM:
		return VE
	case SubdivisionVEN:
		return VE
	case SubdivisionVEO:
		return VE
	case SubdivisionVEP:
		return VE
	case SubdivisionVER:
		return VE
	case SubdivisionVES:
		return VE
	case SubdivisionVET:
		return VE
	case SubdivisionVEU:
		return VE
	case SubdivisionVEV:
		return VE
	case SubdivisionVEW:
		return VE
	case SubdivisionVEX:
		return VE
	case SubdivisionVEY:
		return VE
	case SubdivisionVEZ:
		return VE
	case SubdivisionVN01:
		return VN
	case SubdivisionVN02:
		return VN
	case SubdivisionVN03:
		return VN
	case SubdivisionVN04:
		return VN
	case SubdivisionVN05:
		return VN
	case SubdivisionVN06:
		return VN
	case SubdivisionVN07:
		return VN
	case SubdivisionVN09:
		return VN
	case SubdivisionVN13:
		return VN
	case SubdivisionVN14:
		return VN
	case SubdivisionVN15:
		return VN
	case SubdivisionVN18:
		return VN
	case SubdivisionVN20:
		return VN
	case SubdivisionVN21:
		return VN
	case SubdivisionVN22:
		return VN
	case SubdivisionVN23:
		return VN
	case SubdivisionVN24:
		return VN
	case SubdivisionVN25:
		return VN
	case SubdivisionVN26:
		return VN
	case SubdivisionVN27:
		return VN
	case SubdivisionVN28:
		return VN
	case SubdivisionVN29:
		return VN
	case SubdivisionVN30:
		return VN
	case SubdivisionVN31:
		return VN
	case SubdivisionVN32:
		return VN
	case SubdivisionVN33:
		return VN
	case SubdivisionVN34:
		return VN
	case SubdivisionVN35:
		return VN
	case SubdivisionVN36:
		return VN
	case SubdivisionVN37:
		return VN
	case SubdivisionVN39:
		return VN
	case SubdivisionVN40:
		return VN
	case SubdivisionVN41:
		return VN
	case SubdivisionVN43:
		return VN
	case SubdivisionVN44:
		return VN
	case SubdivisionVN45:
		return VN
	case SubdivisionVN46:
		return VN
	case SubdivisionVN47:
		return VN
	case SubdivisionVN49:
		return VN
	case SubdivisionVN50:
		return VN
	case SubdivisionVN51:
		return VN
	case SubdivisionVN52:
		return VN
	case SubdivisionVN53:
		return VN
	case SubdivisionVN54:
		return VN
	case SubdivisionVN55:
		return VN
	case SubdivisionVN56:
		return VN
	case SubdivisionVN57:
		return VN
	case SubdivisionVN58:
		return VN
	case SubdivisionVN59:
		return VN
	case SubdivisionVN61:
		return VN
	case SubdivisionVN63:
		return VN
	case SubdivisionVN66:
		return VN
	case SubdivisionVN67:
		return VN
	case SubdivisionVN68:
		return VN
	case SubdivisionVN69:
		return VN
	case SubdivisionVN70:
		return VN
	case SubdivisionVN71:
		return VN
	case SubdivisionVN72:
		return VN
	case SubdivisionVN73:
		return VN
	case SubdivisionVNCT:
		return VN
	case SubdivisionVNDN:
		return VN
	case SubdivisionVNHN:
		return VN
	case SubdivisionVNHP:
		return VN
	case SubdivisionVNSG:
		return VN
	case SubdivisionVUMAP:
		return VU
	case SubdivisionVUPAM:
		return VU
	case SubdivisionVUSAM:
		return VU
	case SubdivisionVUSEE:
		return VU
	case SubdivisionVUTAE:
		return VU
	case SubdivisionVUTOB:
		return VU
	case SubdivisionWSAA:
		return WS
	case SubdivisionWSAL:
		return WS
	case SubdivisionWSAT:
		return WS
	case SubdivisionWSFA:
		return WS
	case SubdivisionWSGE:
		return WS
	case SubdivisionWSGI:
		return WS
	case SubdivisionWSPA:
		return WS
	case SubdivisionWSSA:
		return WS
	case SubdivisionWSTU:
		return WS
	case SubdivisionWSVF:
		return WS
	case SubdivisionWSVS:
		return WS
	case SubdivisionYEAB:
		return YE
	case SubdivisionYEAD:
		return YE
	case SubdivisionYEAM:
		return YE
	case SubdivisionYEBA:
		return YE
	case SubdivisionYEDA:
		return YE
	case SubdivisionYEDH:
		return YE
	case SubdivisionYEHD:
		return YE
	case SubdivisionYEHJ:
		return YE
	case SubdivisionYEIB:
		return YE
	case SubdivisionYEJA:
		return YE
	case SubdivisionYELA:
		return YE
	case SubdivisionYEMA:
		return YE
	case SubdivisionYEMR:
		return YE
	case SubdivisionYEMU:
		return YE
	case SubdivisionYEMW:
		return YE
	case SubdivisionYERA:
		return YE
	case SubdivisionYESD:
		return YE
	case SubdivisionYESH:
		return YE
	case SubdivisionYESN:
		return YE
	case SubdivisionYETA:
		return YE
	case SubdivisionZAEC:
		return ZA
	case SubdivisionZAFS:
		return ZA
	case SubdivisionZAGT:
		return ZA
	case SubdivisionZALP:
		return ZA
	case SubdivisionZAMP:
		return ZA
	case SubdivisionZANC:
		return ZA
	case SubdivisionZANL:
		return ZA
	case SubdivisionZANW:
		return ZA
	case SubdivisionZAWC:
		return ZA
	case SubdivisionZM01:
		return ZM
	case SubdivisionZM02:
		return ZM
	case SubdivisionZM03:
		return ZM
	case SubdivisionZM04:
		return ZM
	case SubdivisionZM05:
		return ZM
	case SubdivisionZM06:
		return ZM
	case SubdivisionZM07:
		return ZM
	case SubdivisionZM08:
		return ZM
	case SubdivisionZM09:
		return ZM
	case SubdivisionZWBU:
		return ZW
	case SubdivisionZWHA:
		return ZW
	case SubdivisionZWMA:
		return ZW
	case SubdivisionZWMC:
		return ZW
	case SubdivisionZWME:
		return ZW
	case SubdivisionZWMI:
		return ZW
	case SubdivisionZWMN:
		return ZW
	case SubdivisionZWMS:
		return ZW
	case SubdivisionZWMV:
		return ZW
	case SubdivisionZWMW:
		return ZW
	}

	return Unknown
}

// Info - return CapitalCode as Capital info
func (s SubdivisionCode) Info() *Subdivision {
	return &Subdivision{
		Name:    s.String(),
		Code:    s,
		Country: s.Country(),
	}
}

// IsValid - returns true, if code is correct
func (s SubdivisionCode) IsValid() bool {
	return s.String() != UnknownMsg
}

// SubdivisionType - returns the subdivision type code
//
//nolint:cyclop,funlen,gocyclo
func (s SubdivisionCode) SubdivisionType() SubdivisionTypeCode {
	switch s {
	case SubdivisionAD02:
		return SubdivisionTypeParish
	case SubdivisionAD03:
		return SubdivisionTypeParish
	case SubdivisionAD04:
		return SubdivisionTypeParish
	case SubdivisionAD05:
		return SubdivisionTypeParish
	case SubdivisionAD06:
		return SubdivisionTypeParish
	case SubdivisionAD07:
		return SubdivisionTypeParish
	case SubdivisionAD08:
		return SubdivisionTypeParish
	case SubdivisionAEAJ:
		return SubdivisionTypeEmirate
	case SubdivisionAEAZ:
		return SubdivisionTypeEmirate
	case SubdivisionAEDU:
		return SubdivisionTypeEmirate
	case SubdivisionAEFU:
		return SubdivisionTypeEmirate
	case SubdivisionAERK:
		return SubdivisionTypeEmirate
	case SubdivisionAESH:
		return SubdivisionTypeEmirate
	case SubdivisionAEUQ:
		return SubdivisionTypeEmirate
	case SubdivisionAFBAL:
		return SubdivisionTypeProvince
	case SubdivisionAFBAM:
		return SubdivisionTypeProvince
	case SubdivisionAFBDG:
		return SubdivisionTypeProvince
	case SubdivisionAFBDS:
		return SubdivisionTypeProvince
	case SubdivisionAFBGL:
		return SubdivisionTypeProvince
	case SubdivisionAFDAY:
		return SubdivisionTypeProvince
	case SubdivisionAFFRA:
		return SubdivisionTypeProvince
	case SubdivisionAFFYB:
		return SubdivisionTypeProvince
	case SubdivisionAFGHA:
		return SubdivisionTypeProvince
	case SubdivisionAFGHO:
		return SubdivisionTypeProvince
	case SubdivisionAFHEL:
		return SubdivisionTypeProvince
	case SubdivisionAFHER:
		return SubdivisionTypeProvince
	case SubdivisionAFJOW:
		return SubdivisionTypeProvince
	case SubdivisionAFKAB:
		return SubdivisionTypeProvince
	case SubdivisionAFKAN:
		return SubdivisionTypeProvince
	case SubdivisionAFKAP:
		return SubdivisionTypeProvince
	case SubdivisionAFKDZ:
		return SubdivisionTypeProvince
	case SubdivisionAFKHO:
		return SubdivisionTypeProvince
	case SubdivisionAFKNR:
		return SubdivisionTypeProvince
	case SubdivisionAFLAG:
		return SubdivisionTypeProvince
	case SubdivisionAFLOG:
		return SubdivisionTypeProvince
	case SubdivisionAFNAN:
		return SubdivisionTypeProvince
	case SubdivisionAFNIM:
		return SubdivisionTypeProvince
	case SubdivisionAFNUR:
		return SubdivisionTypeProvince
	case SubdivisionAFPAN:
		return SubdivisionTypeProvince
	case SubdivisionAFPAR:
		return SubdivisionTypeProvince
	case SubdivisionAFPIA:
		return SubdivisionTypeProvince
	case SubdivisionAFPKA:
		return SubdivisionTypeProvince
	case SubdivisionAFSAM:
		return SubdivisionTypeProvince
	case SubdivisionAFSAR:
		return SubdivisionTypeProvince
	case SubdivisionAFTAK:
		return SubdivisionTypeProvince
	case SubdivisionAFURU:
		return SubdivisionTypeProvince
	case SubdivisionAFWAR:
		return SubdivisionTypeProvince
	case SubdivisionAFZAB:
		return SubdivisionTypeProvince
	case SubdivisionAG03:
		return SubdivisionTypeParish
	case SubdivisionAG04:
		return SubdivisionTypeParish
	case SubdivisionAG05:
		return SubdivisionTypeParish
	case SubdivisionAG06:
		return SubdivisionTypeParish
	case SubdivisionAG07:
		return SubdivisionTypeParish
	case SubdivisionAG08:
		return SubdivisionTypeParish
	case SubdivisionAG10:
		return SubdivisionTypeDependency
	case SubdivisionAG11:
		return SubdivisionTypeDependency
	case SubdivisionAL01:
		return SubdivisionTypeCounty
	case SubdivisionAL02:
		return SubdivisionTypeCounty
	case SubdivisionAL03:
		return SubdivisionTypeCounty
	case SubdivisionAL04:
		return SubdivisionTypeCounty
	case SubdivisionAL05:
		return SubdivisionTypeCounty
	case SubdivisionAL06:
		return SubdivisionTypeCounty
	case SubdivisionAL07:
		return SubdivisionTypeCounty
	case SubdivisionAL08:
		return SubdivisionTypeCounty
	case SubdivisionAL09:
		return SubdivisionTypeCounty
	case SubdivisionAL10:
		return SubdivisionTypeCounty
	case SubdivisionAL11:
		return SubdivisionTypeCounty
	case SubdivisionAL12:
		return SubdivisionTypeCounty
	case SubdivisionALBR:
		return SubdivisionTypeDistrict
	case SubdivisionALBU:
		return SubdivisionTypeDistrict
	case SubdivisionALDI:
		return SubdivisionTypeDistrict
	case SubdivisionALDL:
		return SubdivisionTypeDistrict
	case SubdivisionALDR:
		return SubdivisionTypeDistrict
	case SubdivisionALDV:
		return SubdivisionTypeDistrict
	case SubdivisionALEL:
		return SubdivisionTypeDistrict
	case SubdivisionALER:
		return SubdivisionTypeDistrict
	case SubdivisionALFR:
		return SubdivisionTypeDistrict
	case SubdivisionALGJ:
		return SubdivisionTypeDistrict
	case SubdivisionALGR:
		return SubdivisionTypeDistrict
	case SubdivisionALHA:
		return SubdivisionTypeDistrict
	case SubdivisionALKA:
		return SubdivisionTypeDistrict
	case SubdivisionALKB:
		return SubdivisionTypeDistrict
	case SubdivisionALKC:
		return SubdivisionTypeDistrict
	case SubdivisionALKO:
		return SubdivisionTypeDistrict
	case SubdivisionALKR:
		return SubdivisionTypeDistrict
	case SubdivisionALKU:
		return SubdivisionTypeDistrict
	case SubdivisionALLB:
		return SubdivisionTypeDistrict
	case SubdivisionALLE:
		return SubdivisionTypeDistrict
	case SubdivisionALLU:
		return SubdivisionTypeDistrict
	case SubdivisionALMK:
		return SubdivisionTypeDistrict
	case SubdivisionALMM:
		return SubdivisionTypeDistrict
	case SubdivisionALMR:
		return SubdivisionTypeDistrict
	case SubdivisionALMT:
		return SubdivisionTypeDistrict
	case SubdivisionALPG:
		return SubdivisionTypeDistrict
	case SubdivisionALPQ:
		return SubdivisionTypeDistrict
	case SubdivisionALPR:
		return SubdivisionTypeDistrict
	case SubdivisionALPU:
		return SubdivisionTypeDistrict
	case SubdivisionALSH:
		return SubdivisionTypeDistrict
	case SubdivisionALSK:
		return SubdivisionTypeDistrict
	case SubdivisionALSR:
		return SubdivisionTypeDistrict
	case SubdivisionALTE:
		return SubdivisionTypeDistrict
	case SubdivisionALTP:
		return SubdivisionTypeDistrict
	case SubdivisionALTR:
		return SubdivisionTypeDistrict
	case SubdivisionALVL:
		return SubdivisionTypeDistrict
	case SubdivisionAMAG:
		return SubdivisionTypeProvince
	case SubdivisionAMAR:
		return SubdivisionTypeProvince
	case SubdivisionAMAV:
		return SubdivisionTypeProvince
	case SubdivisionAMER:
		return SubdivisionTypeProvince
	case SubdivisionAMGR:
		return SubdivisionTypeProvince
	case SubdivisionAMKT:
		return SubdivisionTypeProvince
	case SubdivisionAMLO:
		return SubdivisionTypeProvince
	case SubdivisionAMSH:
		return SubdivisionTypeProvince
	case SubdivisionAMSU:
		return SubdivisionTypeProvince
	case SubdivisionAMTV:
		return SubdivisionTypeProvince
	case SubdivisionAMVD:
		return SubdivisionTypeProvince
	case SubdivisionAOBGO:
		return SubdivisionTypeProvince
	case SubdivisionAOBGU:
		return SubdivisionTypeProvince
	case SubdivisionAOBIE:
		return SubdivisionTypeProvince
	case SubdivisionAOCAB:
		return SubdivisionTypeProvince
	case SubdivisionAOCCU:
		return SubdivisionTypeProvince
	case SubdivisionAOCNN:
		return SubdivisionTypeProvince
	case SubdivisionAOCNO:
		return SubdivisionTypeProvince
	case SubdivisionAOCUS:
		return SubdivisionTypeProvince
	case SubdivisionAOHUA:
		return SubdivisionTypeProvince
	case SubdivisionAOHUI:
		return SubdivisionTypeProvince
	case SubdivisionAOLNO:
		return SubdivisionTypeProvince
	case SubdivisionAOLSU:
		return SubdivisionTypeProvince
	case SubdivisionAOLUA:
		return SubdivisionTypeProvince
	case SubdivisionAOMAL:
		return SubdivisionTypeProvince
	case SubdivisionAOMOX:
		return SubdivisionTypeProvince
	case SubdivisionAONAM:
		return SubdivisionTypeProvince
	case SubdivisionAOUIG:
		return SubdivisionTypeProvince
	case SubdivisionAOZAI:
		return SubdivisionTypeProvince
	case SubdivisionARA:
		return SubdivisionTypeProvince
	case SubdivisionARB:
		return SubdivisionTypeProvince
	case SubdivisionARC:
		return SubdivisionTypeCity
	case SubdivisionARD:
		return SubdivisionTypeProvince
	case SubdivisionARE:
		return SubdivisionTypeProvince
	case SubdivisionARG:
		return SubdivisionTypeProvince
	case SubdivisionARH:
		return SubdivisionTypeProvince
	case SubdivisionARJ:
		return SubdivisionTypeProvince
	case SubdivisionARK:
		return SubdivisionTypeProvince
	case SubdivisionARL:
		return SubdivisionTypeProvince
	case SubdivisionARM:
		return SubdivisionTypeProvince
	case SubdivisionARN:
		return SubdivisionTypeProvince
	case SubdivisionARP:
		return SubdivisionTypeProvince
	case SubdivisionARQ:
		return SubdivisionTypeProvince
	case SubdivisionARR:
		return SubdivisionTypeProvince
	case SubdivisionARS:
		return SubdivisionTypeProvince
	case SubdivisionART:
		return SubdivisionTypeProvince
	case SubdivisionARU:
		return SubdivisionTypeProvince
	case SubdivisionARV:
		return SubdivisionTypeProvince
	case SubdivisionARW:
		return SubdivisionTypeProvince
	case SubdivisionARX:
		return SubdivisionTypeProvince
	case SubdivisionARY:
		return SubdivisionTypeProvince
	case SubdivisionARZ:
		return SubdivisionTypeProvince
	case SubdivisionAT1:
		return SubdivisionTypeState
	case SubdivisionAT2:
		return SubdivisionTypeState
	case SubdivisionAT3:
		return SubdivisionTypeState
	case SubdivisionAT4:
		return SubdivisionTypeState
	case SubdivisionAT5:
		return SubdivisionTypeState
	case SubdivisionAT6:
		return SubdivisionTypeState
	case SubdivisionAT7:
		return SubdivisionTypeState
	case SubdivisionAT8:
		return SubdivisionTypeState
	case SubdivisionAT9:
		return SubdivisionTypeState
	case SubdivisionAUACT:
		return SubdivisionTypeTerritory
	case SubdivisionAUNSW:
		return SubdivisionTypeState
	case SubdivisionAUNT:
		return SubdivisionTypeTerritory
	case SubdivisionAUQLD:
		return SubdivisionTypeState
	case SubdivisionAUSA:
		return SubdivisionTypeState
	case SubdivisionAUTAS:
		return SubdivisionTypeState
	case SubdivisionAUVIC:
		return SubdivisionTypeState
	case SubdivisionAUWA:
		return SubdivisionTypeState
	case SubdivisionAZABS:
		return SubdivisionTypeRayon
	case SubdivisionAZAGA:
		return SubdivisionTypeRayon
	case SubdivisionAZAGC:
		return SubdivisionTypeRayon
	case SubdivisionAZAGM:
		return SubdivisionTypeRayon
	case SubdivisionAZAGS:
		return SubdivisionTypeRayon
	case SubdivisionAZAGU:
		return SubdivisionTypeRayon
	case SubdivisionAZAST:
		return SubdivisionTypeRayon
	case SubdivisionAZBA:
		return SubdivisionTypeMunicipality
	case SubdivisionAZBAB:
		return SubdivisionTypeRayon
	case SubdivisionAZBAL:
		return SubdivisionTypeRayon
	case SubdivisionAZBAR:
		return SubdivisionTypeRayon
	case SubdivisionAZBEY:
		return SubdivisionTypeRayon
	case SubdivisionAZBIL:
		return SubdivisionTypeRayon
	case SubdivisionAZCAB:
		return SubdivisionTypeRayon
	case SubdivisionAZCAL:
		return SubdivisionTypeRayon
	case SubdivisionAZCUL:
		return SubdivisionTypeRayon
	case SubdivisionAZDAS:
		return SubdivisionTypeRayon
	case SubdivisionAZFUZ:
		return SubdivisionTypeRayon
	case SubdivisionAZGA:
		return SubdivisionTypeMunicipality
	case SubdivisionAZGAD:
		return SubdivisionTypeRayon
	case SubdivisionAZGOR:
		return SubdivisionTypeRayon
	case SubdivisionAZGOY:
		return SubdivisionTypeRayon
	case SubdivisionAZGYG:
		return SubdivisionTypeRayon
	case SubdivisionAZHAC:
		return SubdivisionTypeRayon
	case SubdivisionAZIMI:
		return SubdivisionTypeRayon
	case SubdivisionAZISM:
		return SubdivisionTypeRayon
	case SubdivisionAZKAL:
		return SubdivisionTypeRayon
	case SubdivisionAZKAN:
		return SubdivisionTypeRayon
	case SubdivisionAZKUR:
		return SubdivisionTypeRayon
	case SubdivisionAZLA:
		return SubdivisionTypeMunicipality
	case SubdivisionAZLAC:
		return SubdivisionTypeRayon
	case SubdivisionAZLAN:
		return SubdivisionTypeRayon
	case SubdivisionAZLER:
		return SubdivisionTypeRayon
	case SubdivisionAZMAS:
		return SubdivisionTypeRayon
	case SubdivisionAZMI:
		return SubdivisionTypeMunicipality
	case SubdivisionAZNA:
		return SubdivisionTypeMunicipality
	case SubdivisionAZNEF:
		return SubdivisionTypeRayon
	case SubdivisionAZNV:
		return SubdivisionTypeMunicipality
	case SubdivisionAZNX:
		return SubdivisionTypeAutonomousRepublic
	case SubdivisionAZOGU:
		return SubdivisionTypeRayon
	case SubdivisionAZORD:
		return SubdivisionTypeRayon
	case SubdivisionAZQAB:
		return SubdivisionTypeRayon
	case SubdivisionAZQAX:
		return SubdivisionTypeRayon
	case SubdivisionAZQAZ:
		return SubdivisionTypeRayon
	case SubdivisionAZQBA:
		return SubdivisionTypeRayon
	case SubdivisionAZQBI:
		return SubdivisionTypeRayon
	case SubdivisionAZQOB:
		return SubdivisionTypeRayon
	case SubdivisionAZQUS:
		return SubdivisionTypeRayon
	case SubdivisionAZSA:
		return SubdivisionTypeMunicipality
	case SubdivisionAZSAB:
		return SubdivisionTypeRayon
	case SubdivisionAZSAD:
		return SubdivisionTypeRayon
	case SubdivisionAZSAH:
		return SubdivisionTypeRayon
	case SubdivisionAZSAK:
		return SubdivisionTypeRayon
	case SubdivisionAZSAL:
		return SubdivisionTypeRayon
	case SubdivisionAZSAR:
		return SubdivisionTypeRayon
	case SubdivisionAZSAT:
		return SubdivisionTypeRayon
	case SubdivisionAZSBN:
		return SubdivisionTypeRayon
	case SubdivisionAZSIY:
		return SubdivisionTypeRayon
	case SubdivisionAZSKR:
		return SubdivisionTypeRayon
	case SubdivisionAZSM:
		return SubdivisionTypeMunicipality
	case SubdivisionAZSMI:
		return SubdivisionTypeRayon
	case SubdivisionAZSMX:
		return SubdivisionTypeRayon
	case SubdivisionAZSR:
		return SubdivisionTypeMunicipality
	case SubdivisionAZSUS:
		return SubdivisionTypeRayon
	case SubdivisionAZTAR:
		return SubdivisionTypeRayon
	case SubdivisionAZTOV:
		return SubdivisionTypeRayon
	case SubdivisionAZUCA:
		return SubdivisionTypeRayon
	case SubdivisionAZXA:
		return SubdivisionTypeMunicipality
	case SubdivisionAZXAC:
		return SubdivisionTypeRayon
	case SubdivisionAZXCI:
		return SubdivisionTypeRayon
	case SubdivisionAZXIZ:
		return SubdivisionTypeRayon
	case SubdivisionAZXVD:
		return SubdivisionTypeRayon
	case SubdivisionAZYAR:
		return SubdivisionTypeRayon
	case SubdivisionAZYE:
		return SubdivisionTypeMunicipality
	case SubdivisionAZYEV:
		return SubdivisionTypeRayon
	case SubdivisionAZZAN:
		return SubdivisionTypeRayon
	case SubdivisionAZZAQ:
		return SubdivisionTypeRayon
	case SubdivisionAZZAR:
		return SubdivisionTypeRayon
	case SubdivisionBA01:
		return SubdivisionTypeCanton
	case SubdivisionBA02:
		return SubdivisionTypeCanton
	case SubdivisionBA03:
		return SubdivisionTypeCanton
	case SubdivisionBA04:
		return SubdivisionTypeCanton
	case SubdivisionBA05:
		return SubdivisionTypeCanton
	case SubdivisionBA06:
		return SubdivisionTypeCanton
	case SubdivisionBA07:
		return SubdivisionTypeCanton
	case SubdivisionBA08:
		return SubdivisionTypeCanton
	case SubdivisionBA09:
		return SubdivisionTypeCanton
	case SubdivisionBA10:
		return SubdivisionTypeCanton
	case SubdivisionBABIH:
		return SubdivisionTypeEntity
	case SubdivisionBABRC:
		return SubdivisionTypeDistrict
	case SubdivisionBASRP:
		return SubdivisionTypeEntity
	case SubdivisionBB01:
		return SubdivisionTypeParish
	case SubdivisionBB02:
		return SubdivisionTypeParish
	case SubdivisionBB03:
		return SubdivisionTypeParish
	case SubdivisionBB04:
		return SubdivisionTypeParish
	case SubdivisionBB05:
		return SubdivisionTypeParish
	case SubdivisionBB06:
		return SubdivisionTypeParish
	case SubdivisionBB07:
		return SubdivisionTypeParish
	case SubdivisionBB08:
		return SubdivisionTypeParish
	case SubdivisionBB09:
		return SubdivisionTypeParish
	case SubdivisionBB10:
		return SubdivisionTypeParish
	case SubdivisionBB11:
		return SubdivisionTypeParish
	case SubdivisionBD01:
		return SubdivisionTypeDistrict
	case SubdivisionBD02:
		return SubdivisionTypeDistrict
	case SubdivisionBD03:
		return SubdivisionTypeDistrict
	case SubdivisionBD04:
		return SubdivisionTypeDistrict
	case SubdivisionBD05:
		return SubdivisionTypeDistrict
	case SubdivisionBD06:
		return SubdivisionTypeDistrict
	case SubdivisionBD07:
		return SubdivisionTypeDistrict
	case SubdivisionBD08:
		return SubdivisionTypeDistrict
	case SubdivisionBD09:
		return SubdivisionTypeDistrict
	case SubdivisionBD10:
		return SubdivisionTypeDistrict
	case SubdivisionBD11:
		return SubdivisionTypeDistrict
	case SubdivisionBD12:
		return SubdivisionTypeDistrict
	case SubdivisionBD13:
		return SubdivisionTypeDistrict
	case SubdivisionBD14:
		return SubdivisionTypeDistrict
	case SubdivisionBD15:
		return SubdivisionTypeDistrict
	case SubdivisionBD16:
		return SubdivisionTypeDistrict
	case SubdivisionBD17:
		return SubdivisionTypeDistrict
	case SubdivisionBD18:
		return SubdivisionTypeDistrict
	case SubdivisionBD19:
		return SubdivisionTypeDistrict
	case SubdivisionBD20:
		return SubdivisionTypeDistrict
	case SubdivisionBD21:
		return SubdivisionTypeDistrict
	case SubdivisionBD22:
		return SubdivisionTypeDistrict
	case SubdivisionBD23:
		return SubdivisionTypeDistrict
	case SubdivisionBD24:
		return SubdivisionTypeDistrict
	case SubdivisionBD25:
		return SubdivisionTypeDistrict
	case SubdivisionBD26:
		return SubdivisionTypeDistrict
	case SubdivisionBD27:
		return SubdivisionTypeDistrict
	case SubdivisionBD28:
		return SubdivisionTypeDistrict
	case SubdivisionBD29:
		return SubdivisionTypeDistrict
	case SubdivisionBD30:
		return SubdivisionTypeDistrict
	case SubdivisionBD31:
		return SubdivisionTypeDistrict
	case SubdivisionBD32:
		return SubdivisionTypeDistrict
	case SubdivisionBD33:
		return SubdivisionTypeDistrict
	case SubdivisionBD34:
		return SubdivisionTypeDistrict
	case SubdivisionBD35:
		return SubdivisionTypeDistrict
	case SubdivisionBD36:
		return SubdivisionTypeDistrict
	case SubdivisionBD37:
		return SubdivisionTypeDistrict
	case SubdivisionBD38:
		return SubdivisionTypeDistrict
	case SubdivisionBD39:
		return SubdivisionTypeDistrict
	case SubdivisionBD40:
		return SubdivisionTypeDistrict
	case SubdivisionBD41:
		return SubdivisionTypeDistrict
	case SubdivisionBD42:
		return SubdivisionTypeDistrict
	case SubdivisionBD43:
		return SubdivisionTypeDistrict
	case SubdivisionBD44:
		return SubdivisionTypeDistrict
	case SubdivisionBD45:
		return SubdivisionTypeDistrict
	case SubdivisionBD46:
		return SubdivisionTypeDistrict
	case SubdivisionBD47:
		return SubdivisionTypeDistrict
	case SubdivisionBD48:
		return SubdivisionTypeDistrict
	case SubdivisionBD49:
		return SubdivisionTypeDistrict
	case SubdivisionBD50:
		return SubdivisionTypeDistrict
	case SubdivisionBD51:
		return SubdivisionTypeDistrict
	case SubdivisionBD52:
		return SubdivisionTypeDistrict
	case SubdivisionBD53:
		return SubdivisionTypeDistrict
	case SubdivisionBD54:
		return SubdivisionTypeDistrict
	case SubdivisionBD55:
		return SubdivisionTypeDistrict
	case SubdivisionBD56:
		return SubdivisionTypeDistrict
	case SubdivisionBD57:
		return SubdivisionTypeDistrict
	case SubdivisionBD58:
		return SubdivisionTypeDistrict
	case SubdivisionBD59:
		return SubdivisionTypeDistrict
	case SubdivisionBD60:
		return SubdivisionTypeDistrict
	case SubdivisionBD61:
		return SubdivisionTypeDistrict
	case SubdivisionBD62:
		return SubdivisionTypeDistrict
	case SubdivisionBD63:
		return SubdivisionTypeDistrict
	case SubdivisionBD64:
		return SubdivisionTypeDistrict
	case SubdivisionBDA:
		return SubdivisionTypeDivision
	case SubdivisionBDB:
		return SubdivisionTypeDivision
	case SubdivisionBDC:
		return SubdivisionTypeDivision
	case SubdivisionBDD:
		return SubdivisionTypeDivision
	case SubdivisionBDE:
		return SubdivisionTypeDivision
	case SubdivisionBDF:
		return SubdivisionTypeDivision
	case SubdivisionBDG:
		return SubdivisionTypeDivision
	case SubdivisionBDH:
		return SubdivisionTypeDivision
	case SubdivisionBEBRU:
		return SubdivisionTypeRegion
	case SubdivisionBEVAN:
		return SubdivisionTypeProvince
	case SubdivisionBEVBR:
		return SubdivisionTypeProvince
	case SubdivisionBEVLG:
		return SubdivisionTypeRegion
	case SubdivisionBEVLI:
		return SubdivisionTypeProvince
	case SubdivisionBEVOV:
		return SubdivisionTypeProvince
	case SubdivisionBEVWV:
		return SubdivisionTypeProvince
	case SubdivisionBEWAL:
		return SubdivisionTypeRegion
	case SubdivisionBEWBR:
		return SubdivisionTypeProvince
	case SubdivisionBEWHT:
		return SubdivisionTypeProvince
	case SubdivisionBEWLG:
		return SubdivisionTypeProvince
	case SubdivisionBEWLX:
		return SubdivisionTypeProvince
	case SubdivisionBEWNA:
		return SubdivisionTypeProvince
	case SubdivisionBF01:
		return SubdivisionTypeRegion
	case SubdivisionBF02:
		return SubdivisionTypeRegion
	case SubdivisionBF03:
		return SubdivisionTypeRegion
	case SubdivisionBF04:
		return SubdivisionTypeRegion
	case SubdivisionBF05:
		return SubdivisionTypeRegion
	case SubdivisionBF06:
		return SubdivisionTypeRegion
	case SubdivisionBF07:
		return SubdivisionTypeRegion
	case SubdivisionBF08:
		return SubdivisionTypeRegion
	case SubdivisionBF09:
		return SubdivisionTypeRegion
	case SubdivisionBF10:
		return SubdivisionTypeRegion
	case SubdivisionBF11:
		return SubdivisionTypeRegion
	case SubdivisionBF12:
		return SubdivisionTypeRegion
	case SubdivisionBF13:
		return SubdivisionTypeRegion
	case SubdivisionBFBAL:
		return SubdivisionTypeProvince
	case SubdivisionBFBAM:
		return SubdivisionTypeProvince
	case SubdivisionBFBAN:
		return SubdivisionTypeProvince
	case SubdivisionBFBAZ:
		return SubdivisionTypeProvince
	case SubdivisionBFBGR:
		return SubdivisionTypeProvince
	case SubdivisionBFBLG:
		return SubdivisionTypeProvince
	case SubdivisionBFBLK:
		return SubdivisionTypeProvince
	case SubdivisionBFCOM:
		return SubdivisionTypeProvince
	case SubdivisionBFGAN:
		return SubdivisionTypeProvince
	case SubdivisionBFGNA:
		return SubdivisionTypeProvince
	case SubdivisionBFGOU:
		return SubdivisionTypeProvince
	case SubdivisionBFHOU:
		return SubdivisionTypeProvince
	case SubdivisionBFIOB:
		return SubdivisionTypeProvince
	case SubdivisionBFKAD:
		return SubdivisionTypeProvince
	case SubdivisionBFKEN:
		return SubdivisionTypeProvince
	case SubdivisionBFKMD:
		return SubdivisionTypeProvince
	case SubdivisionBFKMP:
		return SubdivisionTypeProvince
	case SubdivisionBFKOP:
		return SubdivisionTypeProvince
	case SubdivisionBFKOS:
		return SubdivisionTypeProvince
	case SubdivisionBFKOT:
		return SubdivisionTypeProvince
	case SubdivisionBFKOW:
		return SubdivisionTypeProvince
	case SubdivisionBFLER:
		return SubdivisionTypeProvince
	case SubdivisionBFLOR:
		return SubdivisionTypeProvince
	case SubdivisionBFMOU:
		return SubdivisionTypeProvince
	case SubdivisionBFNAM:
		return SubdivisionTypeProvince
	case SubdivisionBFNAO:
		return SubdivisionTypeProvince
	case SubdivisionBFNAY:
		return SubdivisionTypeProvince
	case SubdivisionBFNOU:
		return SubdivisionTypeProvince
	case SubdivisionBFOUB:
		return SubdivisionTypeProvince
	case SubdivisionBFOUD:
		return SubdivisionTypeProvince
	case SubdivisionBFPAS:
		return SubdivisionTypeProvince
	case SubdivisionBFPON:
		return SubdivisionTypeProvince
	case SubdivisionBFSEN:
		return SubdivisionTypeProvince
	case SubdivisionBFSIS:
		return SubdivisionTypeProvince
	case SubdivisionBFSMT:
		return SubdivisionTypeProvince
	case SubdivisionBFSNG:
		return SubdivisionTypeProvince
	case SubdivisionBFSOM:
		return SubdivisionTypeProvince
	case SubdivisionBFSOR:
		return SubdivisionTypeProvince
	case SubdivisionBFTAP:
		return SubdivisionTypeProvince
	case SubdivisionBFTUI:
		return SubdivisionTypeProvince
	case SubdivisionBFYAG:
		return SubdivisionTypeProvince
	case SubdivisionBFYAT:
		return SubdivisionTypeProvince
	case SubdivisionBFZIR:
		return SubdivisionTypeProvince
	case SubdivisionBFZON:
		return SubdivisionTypeProvince
	case SubdivisionBFZOU:
		return SubdivisionTypeProvince
	case SubdivisionBG01:
		return SubdivisionTypeRegion
	case SubdivisionBG02:
		return SubdivisionTypeRegion
	case SubdivisionBG03:
		return SubdivisionTypeRegion
	case SubdivisionBG04:
		return SubdivisionTypeRegion
	case SubdivisionBG05:
		return SubdivisionTypeRegion
	case SubdivisionBG06:
		return SubdivisionTypeRegion
	case SubdivisionBG07:
		return SubdivisionTypeRegion
	case SubdivisionBG08:
		return SubdivisionTypeRegion
	case SubdivisionBG09:
		return SubdivisionTypeRegion
	case SubdivisionBG10:
		return SubdivisionTypeRegion
	case SubdivisionBG11:
		return SubdivisionTypeRegion
	case SubdivisionBG12:
		return SubdivisionTypeRegion
	case SubdivisionBG13:
		return SubdivisionTypeRegion
	case SubdivisionBG14:
		return SubdivisionTypeRegion
	case SubdivisionBG15:
		return SubdivisionTypeRegion
	case SubdivisionBG16:
		return SubdivisionTypeRegion
	case SubdivisionBG17:
		return SubdivisionTypeRegion
	case SubdivisionBG18:
		return SubdivisionTypeRegion
	case SubdivisionBG19:
		return SubdivisionTypeRegion
	case SubdivisionBG20:
		return SubdivisionTypeRegion
	case SubdivisionBG21:
		return SubdivisionTypeRegion
	case SubdivisionBG22:
		return SubdivisionTypeRegion
	case SubdivisionBG23:
		return SubdivisionTypeRegion
	case SubdivisionBG24:
		return SubdivisionTypeRegion
	case SubdivisionBG25:
		return SubdivisionTypeRegion
	case SubdivisionBG26:
		return SubdivisionTypeRegion
	case SubdivisionBG27:
		return SubdivisionTypeRegion
	case SubdivisionBG28:
		return SubdivisionTypeRegion
	case SubdivisionBH13:
		return SubdivisionTypeGovernorate
	case SubdivisionBH14:
		return SubdivisionTypeGovernorate
	case SubdivisionBH15:
		return SubdivisionTypeGovernorate
	case SubdivisionBH16:
		return SubdivisionTypeGovernorate
	case SubdivisionBH17:
		return SubdivisionTypeGovernorate
	case SubdivisionBIBB:
		return SubdivisionTypeProvince
	case SubdivisionBIBL:
		return SubdivisionTypeProvince
	case SubdivisionBIBM:
		return SubdivisionTypeProvince
	case SubdivisionBIBR:
		return SubdivisionTypeProvince
	case SubdivisionBICA:
		return SubdivisionTypeProvince
	case SubdivisionBICI:
		return SubdivisionTypeProvince
	case SubdivisionBIGI:
		return SubdivisionTypeProvince
	case SubdivisionBIKI:
		return SubdivisionTypeProvince
	case SubdivisionBIKR:
		return SubdivisionTypeProvince
	case SubdivisionBIKY:
		return SubdivisionTypeProvince
	case SubdivisionBIMA:
		return SubdivisionTypeProvince
	case SubdivisionBIMU:
		return SubdivisionTypeProvince
	case SubdivisionBIMW:
		return SubdivisionTypeProvince
	case SubdivisionBING:
		return SubdivisionTypeProvince
	case SubdivisionBIRT:
		return SubdivisionTypeProvince
	case SubdivisionBIRY:
		return SubdivisionTypeProvince
	case SubdivisionBJAK:
		return SubdivisionTypeDepartment
	case SubdivisionBJAL:
		return SubdivisionTypeDepartment
	case SubdivisionBJAQ:
		return SubdivisionTypeDepartment
	case SubdivisionBJBO:
		return SubdivisionTypeDepartment
	case SubdivisionBJCO:
		return SubdivisionTypeDepartment
	case SubdivisionBJDO:
		return SubdivisionTypeDepartment
	case SubdivisionBJKO:
		return SubdivisionTypeDepartment
	case SubdivisionBJLI:
		return SubdivisionTypeDepartment
	case SubdivisionBJMO:
		return SubdivisionTypeDepartment
	case SubdivisionBJOU:
		return SubdivisionTypeDepartment
	case SubdivisionBJPL:
		return SubdivisionTypeDepartment
	case SubdivisionBJZO:
		return SubdivisionTypeDepartment
	case SubdivisionBNBE:
		return SubdivisionTypeDistrict
	case SubdivisionBNBM:
		return SubdivisionTypeDistrict
	case SubdivisionBNTE:
		return SubdivisionTypeDistrict
	case SubdivisionBNTU:
		return SubdivisionTypeDistrict
	case SubdivisionBOB:
		return SubdivisionTypeDepartment
	case SubdivisionBOC:
		return SubdivisionTypeDepartment
	case SubdivisionBOH:
		return SubdivisionTypeDepartment
	case SubdivisionBOL:
		return SubdivisionTypeDepartment
	case SubdivisionBON:
		return SubdivisionTypeDepartment
	case SubdivisionBOO:
		return SubdivisionTypeDepartment
	case SubdivisionBOP:
		return SubdivisionTypeDepartment
	case SubdivisionBOS:
		return SubdivisionTypeDepartment
	case SubdivisionBOT:
		return SubdivisionTypeDepartment
	case SubdivisionBQBO:
		return SubdivisionTypeSpecialMunicipality
	case SubdivisionBQSA:
		return SubdivisionTypeSpecialMunicipality
	case SubdivisionBQSE:
		return SubdivisionTypeSpecialMunicipality
	case SubdivisionBRAC:
		return SubdivisionTypeState
	case SubdivisionBRAL:
		return SubdivisionTypeState
	case SubdivisionBRAM:
		return SubdivisionTypeState
	case SubdivisionBRAP:
		return SubdivisionTypeState
	case SubdivisionBRBA:
		return SubdivisionTypeState
	case SubdivisionBRCE:
		return SubdivisionTypeState
	case SubdivisionBRDF:
		return SubdivisionTypeFederalDistrict
	case SubdivisionBRES:
		return SubdivisionTypeState
	case SubdivisionBRFN:
		return SubdivisionTypeState
	case SubdivisionBRGO:
		return SubdivisionTypeState
	case SubdivisionBRMA:
		return SubdivisionTypeState
	case SubdivisionBRMG:
		return SubdivisionTypeState
	case SubdivisionBRMS:
		return SubdivisionTypeState
	case SubdivisionBRMT:
		return SubdivisionTypeState
	case SubdivisionBRPA:
		return SubdivisionTypeState
	case SubdivisionBRPB:
		return SubdivisionTypeState
	case SubdivisionBRPE:
		return SubdivisionTypeState
	case SubdivisionBRPI:
		return SubdivisionTypeState
	case SubdivisionBRPR:
		return SubdivisionTypeState
	case SubdivisionBRRJ:
		return SubdivisionTypeState
	case SubdivisionBRRN:
		return SubdivisionTypeState
	case SubdivisionBRRO:
		return SubdivisionTypeState
	case SubdivisionBRRR:
		return SubdivisionTypeState
	case SubdivisionBRRS:
		return SubdivisionTypeState
	case SubdivisionBRSC:
		return SubdivisionTypeState
	case SubdivisionBRSE:
		return SubdivisionTypeState
	case SubdivisionBRSP:
		return SubdivisionTypeState
	case SubdivisionBRTO:
		return SubdivisionTypeState
	case SubdivisionBSAK:
		return SubdivisionTypeDistrict
	case SubdivisionBSBI:
		return SubdivisionTypeDistrict
	case SubdivisionBSBP:
		return SubdivisionTypeDistrict
	case SubdivisionBSBY:
		return SubdivisionTypeDistrict
	case SubdivisionBSCE:
		return SubdivisionTypeDistrict
	case SubdivisionBSCI:
		return SubdivisionTypeDistrict
	case SubdivisionBSCK:
		return SubdivisionTypeDistrict
	case SubdivisionBSCO:
		return SubdivisionTypeDistrict
	case SubdivisionBSCS:
		return SubdivisionTypeDistrict
	case SubdivisionBSEG:
		return SubdivisionTypeDistrict
	case SubdivisionBSEX:
		return SubdivisionTypeDistrict
	case SubdivisionBSFP:
		return SubdivisionTypeDistrict
	case SubdivisionBSGC:
		return SubdivisionTypeDistrict
	case SubdivisionBSHI:
		return SubdivisionTypeDistrict
	case SubdivisionBSHT:
		return SubdivisionTypeDistrict
	case SubdivisionBSIN:
		return SubdivisionTypeDistrict
	case SubdivisionBSLI:
		return SubdivisionTypeDistrict
	case SubdivisionBSMC:
		return SubdivisionTypeDistrict
	case SubdivisionBSMG:
		return SubdivisionTypeDistrict
	case SubdivisionBSMI:
		return SubdivisionTypeDistrict
	case SubdivisionBSNE:
		return SubdivisionTypeDistrict
	case SubdivisionBSNO:
		return SubdivisionTypeDistrict
	case SubdivisionBSNS:
		return SubdivisionTypeDistrict
	case SubdivisionBSRC:
		return SubdivisionTypeDistrict
	case SubdivisionBSRI:
		return SubdivisionTypeDistrict
	case SubdivisionBSSA:
		return SubdivisionTypeDistrict
	case SubdivisionBSSE:
		return SubdivisionTypeDistrict
	case SubdivisionBSSO:
		return SubdivisionTypeDistrict
	case SubdivisionBSSS:
		return SubdivisionTypeDistrict
	case SubdivisionBSSW:
		return SubdivisionTypeDistrict
	case SubdivisionBSWG:
		return SubdivisionTypeDistrict
	case SubdivisionBT11:
		return SubdivisionTypeDistrict
	case SubdivisionBT12:
		return SubdivisionTypeDistrict
	case SubdivisionBT13:
		return SubdivisionTypeDistrict
	case SubdivisionBT14:
		return SubdivisionTypeDistrict
	case SubdivisionBT15:
		return SubdivisionTypeDistrict
	case SubdivisionBT21:
		return SubdivisionTypeDistrict
	case SubdivisionBT22:
		return SubdivisionTypeDistrict
	case SubdivisionBT23:
		return SubdivisionTypeDistrict
	case SubdivisionBT24:
		return SubdivisionTypeDistrict
	case SubdivisionBT31:
		return SubdivisionTypeDistrict
	case SubdivisionBT32:
		return SubdivisionTypeDistrict
	case SubdivisionBT33:
		return SubdivisionTypeDistrict
	case SubdivisionBT34:
		return SubdivisionTypeDistrict
	case SubdivisionBT41:
		return SubdivisionTypeDistrict
	case SubdivisionBT42:
		return SubdivisionTypeDistrict
	case SubdivisionBT43:
		return SubdivisionTypeDistrict
	case SubdivisionBT44:
		return SubdivisionTypeDistrict
	case SubdivisionBT45:
		return SubdivisionTypeDistrict
	case SubdivisionBTGA:
		return SubdivisionTypeDistrict
	case SubdivisionBTTY:
		return SubdivisionTypeDistrict
	case SubdivisionBWCE:
		return SubdivisionTypeDistrict
	case SubdivisionBWGH:
		return SubdivisionTypeDistrict
	case SubdivisionBWKG:
		return SubdivisionTypeDistrict
	case SubdivisionBWKL:
		return SubdivisionTypeDistrict
	case SubdivisionBWKW:
		return SubdivisionTypeDistrict
	case SubdivisionBWNE:
		return SubdivisionTypeDistrict
	case SubdivisionBWNW:
		return SubdivisionTypeDistrict
	case SubdivisionBWSE:
		return SubdivisionTypeDistrict
	case SubdivisionBWSO:
		return SubdivisionTypeDistrict
	case SubdivisionBYBR:
		return SubdivisionTypeOblast
	case SubdivisionBYHM:
		return SubdivisionTypeCity
	case SubdivisionBYHO:
		return SubdivisionTypeOblast
	case SubdivisionBYHR:
		return SubdivisionTypeOblast
	case SubdivisionBYMA:
		return SubdivisionTypeOblast
	case SubdivisionBYMI:
		return SubdivisionTypeOblast
	case SubdivisionBYVI:
		return SubdivisionTypeOblast
	case SubdivisionBZBZ:
		return SubdivisionTypeDistrict
	case SubdivisionBZCY:
		return SubdivisionTypeDistrict
	case SubdivisionBZCZL:
		return SubdivisionTypeDistrict
	case SubdivisionBZOW:
		return SubdivisionTypeDistrict
	case SubdivisionBZSC:
		return SubdivisionTypeDistrict
	case SubdivisionBZTOL:
		return SubdivisionTypeDistrict
	case SubdivisionCAAB:
		return SubdivisionTypeProvince
	case SubdivisionCABC:
		return SubdivisionTypeProvince
	case SubdivisionCAMB:
		return SubdivisionTypeProvince
	case SubdivisionCANB:
		return SubdivisionTypeProvince
	case SubdivisionCANL:
		return SubdivisionTypeProvince
	case SubdivisionCANS:
		return SubdivisionTypeProvince
	case SubdivisionCANT:
		return SubdivisionTypeTerritory
	case SubdivisionCANU:
		return SubdivisionTypeTerritory
	case SubdivisionCAON:
		return SubdivisionTypeProvince
	case SubdivisionCAPE:
		return SubdivisionTypeProvince
	case SubdivisionCAQC:
		return SubdivisionTypeProvince
	case SubdivisionCASK:
		return SubdivisionTypeProvince
	case SubdivisionCAYT:
		return SubdivisionTypeTerritory
	case SubdivisionCDBC:
		return SubdivisionTypeProvince
	case SubdivisionCDBN:
		return SubdivisionTypeProvince
	case SubdivisionCDEQ:
		return SubdivisionTypeProvince
	case SubdivisionCDKA:
		return SubdivisionTypeProvince
	case SubdivisionCDKE:
		return SubdivisionTypeProvince
	case SubdivisionCDKN:
		return SubdivisionTypeCity
	case SubdivisionCDKW:
		return SubdivisionTypeProvince
	case SubdivisionCDMA:
		return SubdivisionTypeProvince
	case SubdivisionCDNK:
		return SubdivisionTypeProvince
	case SubdivisionCDOR:
		return SubdivisionTypeProvince
	case SubdivisionCDSK:
		return SubdivisionTypeProvince
	case SubdivisionCFAC:
		return SubdivisionTypePrefecture
	case SubdivisionCFBB:
		return SubdivisionTypePrefecture
	case SubdivisionCFBGF:
		return SubdivisionTypeCommune
	case SubdivisionCFBK:
		return SubdivisionTypePrefecture
	case SubdivisionCFHK:
		return SubdivisionTypePrefecture
	case SubdivisionCFHM:
		return SubdivisionTypePrefecture
	case SubdivisionCFHS:
		return SubdivisionTypePrefecture
	case SubdivisionCFKB:
		return SubdivisionTypeEconomicPrefecture
	case SubdivisionCFKG:
		return SubdivisionTypePrefecture
	case SubdivisionCFLB:
		return SubdivisionTypePrefecture
	case SubdivisionCFMB:
		return SubdivisionTypePrefecture
	case SubdivisionCFMP:
		return SubdivisionTypePrefecture
	case SubdivisionCFNM:
		return SubdivisionTypePrefecture
	case SubdivisionCFOP:
		return SubdivisionTypePrefecture
	case SubdivisionCFSE:
		return SubdivisionTypeEconomicPrefecture
	case SubdivisionCFUK:
		return SubdivisionTypePrefecture
	case SubdivisionCFVK:
		return SubdivisionTypePrefecture
	case SubdivisionCG11:
		return SubdivisionTypeRegion
	case SubdivisionCG12:
		return SubdivisionTypeRegion
	case SubdivisionCG13:
		return SubdivisionTypeRegion
	case SubdivisionCG14:
		return SubdivisionTypeRegion
	case SubdivisionCG15:
		return SubdivisionTypeRegion
	case SubdivisionCG2:
		return SubdivisionTypeRegion
	case SubdivisionCG5:
		return SubdivisionTypeRegion
	case SubdivisionCG7:
		return SubdivisionTypeRegion
	case SubdivisionCG8:
		return SubdivisionTypeRegion
	case SubdivisionCG9:
		return SubdivisionTypeRegion
	case SubdivisionCGBZV:
		return SubdivisionTypeCapitalDistrict
	case SubdivisionCHAG:
		return SubdivisionTypeCanton
	case SubdivisionCHAI:
		return SubdivisionTypeCanton
	case SubdivisionCHAR:
		return SubdivisionTypeCanton
	case SubdivisionCHBE:
		return SubdivisionTypeCanton
	case SubdivisionCHBL:
		return SubdivisionTypeCanton
	case SubdivisionCHBS:
		return SubdivisionTypeCanton
	case SubdivisionCHFR:
		return SubdivisionTypeCanton
	case SubdivisionCHGE:
		return SubdivisionTypeCanton
	case SubdivisionCHGL:
		return SubdivisionTypeCanton
	case SubdivisionCHGR:
		return SubdivisionTypeCanton
	case SubdivisionCHJU:
		return SubdivisionTypeCanton
	case SubdivisionCHLU:
		return SubdivisionTypeCanton
	case SubdivisionCHNE:
		return SubdivisionTypeCanton
	case SubdivisionCHNW:
		return SubdivisionTypeCanton
	case SubdivisionCHOW:
		return SubdivisionTypeCanton
	case SubdivisionCHSG:
		return SubdivisionTypeCanton
	case SubdivisionCHSH:
		return SubdivisionTypeCanton
	case SubdivisionCHSO:
		return SubdivisionTypeCanton
	case SubdivisionCHSZ:
		return SubdivisionTypeCanton
	case SubdivisionCHTG:
		return SubdivisionTypeCanton
	case SubdivisionCHTI:
		return SubdivisionTypeCanton
	case SubdivisionCHUR:
		return SubdivisionTypeCanton
	case SubdivisionCHVD:
		return SubdivisionTypeCanton
	case SubdivisionCHVS:
		return SubdivisionTypeCanton
	case SubdivisionCHZG:
		return SubdivisionTypeCanton
	case SubdivisionCHZH:
		return SubdivisionTypeCanton
	case SubdivisionCI01:
		return SubdivisionTypeRegion
	case SubdivisionCI02:
		return SubdivisionTypeRegion
	case SubdivisionCI03:
		return SubdivisionTypeRegion
	case SubdivisionCI04:
		return SubdivisionTypeRegion
	case SubdivisionCI05:
		return SubdivisionTypeRegion
	case SubdivisionCI06:
		return SubdivisionTypeRegion
	case SubdivisionCI07:
		return SubdivisionTypeRegion
	case SubdivisionCI08:
		return SubdivisionTypeRegion
	case SubdivisionCI09:
		return SubdivisionTypeRegion
	case SubdivisionCI10:
		return SubdivisionTypeRegion
	case SubdivisionCI11:
		return SubdivisionTypeRegion
	case SubdivisionCI12:
		return SubdivisionTypeRegion
	case SubdivisionCI13:
		return SubdivisionTypeRegion
	case SubdivisionCI14:
		return SubdivisionTypeRegion
	case SubdivisionCI15:
		return SubdivisionTypeRegion
	case SubdivisionCI16:
		return SubdivisionTypeRegion
	case SubdivisionCI17:
		return SubdivisionTypeRegion
	case SubdivisionCI18:
		return SubdivisionTypeRegion
	case SubdivisionCI19:
		return SubdivisionTypeRegion
	case SubdivisionCLAI:
		return SubdivisionTypeRegion
	case SubdivisionCLAN:
		return SubdivisionTypeRegion
	case SubdivisionCLAP:
		return SubdivisionTypeRegion
	case SubdivisionCLAR:
		return SubdivisionTypeRegion
	case SubdivisionCLAT:
		return SubdivisionTypeRegion
	case SubdivisionCLBI:
		return SubdivisionTypeRegion
	case SubdivisionCLCO:
		return SubdivisionTypeRegion
	case SubdivisionCLLI:
		return SubdivisionTypeRegion
	case SubdivisionCLLL:
		return SubdivisionTypeRegion
	case SubdivisionCLLR:
		return SubdivisionTypeRegion
	case SubdivisionCLMA:
		return SubdivisionTypeRegion
	case SubdivisionCLML:
		return SubdivisionTypeRegion
	case SubdivisionCLRM:
		return SubdivisionTypeRegion
	case SubdivisionCLTA:
		return SubdivisionTypeRegion
	case SubdivisionCLVS:
		return SubdivisionTypeRegion
	case SubdivisionCMAD:
		return SubdivisionTypeProvince
	case SubdivisionCMCE:
		return SubdivisionTypeProvince
	case SubdivisionCMEN:
		return SubdivisionTypeProvince
	case SubdivisionCMES:
		return SubdivisionTypeProvince
	case SubdivisionCMLT:
		return SubdivisionTypeProvince
	case SubdivisionCMNO:
		return SubdivisionTypeProvince
	case SubdivisionCMNW:
		return SubdivisionTypeProvince
	case SubdivisionCMOU:
		return SubdivisionTypeProvince
	case SubdivisionCMSU:
		return SubdivisionTypeProvince
	case SubdivisionCMSW:
		return SubdivisionTypeProvince
	case SubdivisionCNAH:
		return SubdivisionTypeProvince
	case SubdivisionCNBJ:
		return SubdivisionTypeMunicipality
	case SubdivisionCNCQ:
		return SubdivisionTypeMunicipality
	case SubdivisionCNFJ:
		return SubdivisionTypeProvince
	case SubdivisionCNGD:
		return SubdivisionTypeProvince
	case SubdivisionCNGS:
		return SubdivisionTypeProvince
	case SubdivisionCNGX:
		return SubdivisionTypeAutonomousRegion
	case SubdivisionCNGZ:
		return SubdivisionTypeProvince
	case SubdivisionCNHA:
		return SubdivisionTypeProvince
	case SubdivisionCNHB:
		return SubdivisionTypeProvince
	case SubdivisionCNHE:
		return SubdivisionTypeProvince
	case SubdivisionCNHI:
		return SubdivisionTypeProvince
	case SubdivisionCNHK:
		return SubdivisionTypeSpecialAdministrativeRegion
	case SubdivisionCNHL:
		return SubdivisionTypeProvince
	case SubdivisionCNHN:
		return SubdivisionTypeProvince
	case SubdivisionCNJL:
		return SubdivisionTypeProvince
	case SubdivisionCNJS:
		return SubdivisionTypeProvince
	case SubdivisionCNJX:
		return SubdivisionTypeProvince
	case SubdivisionCNLN:
		return SubdivisionTypeProvince
	case SubdivisionCNMO:
		return SubdivisionTypeSpecialAdministrativeRegion
	case SubdivisionCNNM:
		return SubdivisionTypeAutonomousRegion
	case SubdivisionCNNX:
		return SubdivisionTypeAutonomousRegion
	case SubdivisionCNQH:
		return SubdivisionTypeProvince
	case SubdivisionCNSC:
		return SubdivisionTypeProvince
	case SubdivisionCNSD:
		return SubdivisionTypeProvince
	case SubdivisionCNSH:
		return SubdivisionTypeMunicipality
	case SubdivisionCNSN:
		return SubdivisionTypeProvince
	case SubdivisionCNSX:
		return SubdivisionTypeProvince
	case SubdivisionCNTJ:
		return SubdivisionTypeMunicipality
	case SubdivisionCNTW:
		return SubdivisionTypeProvince
	case SubdivisionCNXJ:
		return SubdivisionTypeAutonomousRegion
	case SubdivisionCNXZ:
		return SubdivisionTypeAutonomousRegion
	case SubdivisionCNYN:
		return SubdivisionTypeProvince
	case SubdivisionCNZJ:
		return SubdivisionTypeProvince
	case SubdivisionCOAMA:
		return SubdivisionTypeDepartment
	case SubdivisionCOANT:
		return SubdivisionTypeDepartment
	case SubdivisionCOARA:
		return SubdivisionTypeDepartment
	case SubdivisionCOATL:
		return SubdivisionTypeDepartment
	case SubdivisionCOBOL:
		return SubdivisionTypeDepartment
	case SubdivisionCOBOY:
		return SubdivisionTypeDepartment
	case SubdivisionCOCAL:
		return SubdivisionTypeDepartment
	case SubdivisionCOCAQ:
		return SubdivisionTypeDepartment
	case SubdivisionCOCAS:
		return SubdivisionTypeDepartment
	case SubdivisionCOCAU:
		return SubdivisionTypeDepartment
	case SubdivisionCOCES:
		return SubdivisionTypeDepartment
	case SubdivisionCOCHO:
		return SubdivisionTypeDepartment
	case SubdivisionCOCOR:
		return SubdivisionTypeDepartment
	case SubdivisionCOCUN:
		return SubdivisionTypeDepartment
	case SubdivisionCODC:
		return SubdivisionTypeCapitalDistrict
	case SubdivisionCOGUA:
		return SubdivisionTypeDepartment
	case SubdivisionCOGUV:
		return SubdivisionTypeDepartment
	case SubdivisionCOHUI:
		return SubdivisionTypeDepartment
	case SubdivisionCOLAG:
		return SubdivisionTypeDepartment
	case SubdivisionCOMAG:
		return SubdivisionTypeDepartment
	case SubdivisionCOMET:
		return SubdivisionTypeDepartment
	case SubdivisionCONAR:
		return SubdivisionTypeDepartment
	case SubdivisionCONSA:
		return SubdivisionTypeDepartment
	case SubdivisionCOPUT:
		return SubdivisionTypeDepartment
	case SubdivisionCOQUI:
		return SubdivisionTypeDepartment
	case SubdivisionCORIS:
		return SubdivisionTypeDepartment
	case SubdivisionCOSAN:
		return SubdivisionTypeDepartment
	case SubdivisionCOSAP:
		return SubdivisionTypeDepartment
	case SubdivisionCOSUC:
		return SubdivisionTypeDepartment
	case SubdivisionCOTOL:
		return SubdivisionTypeDepartment
	case SubdivisionCOVAC:
		return SubdivisionTypeDepartment
	case SubdivisionCOVAU:
		return SubdivisionTypeDepartment
	case SubdivisionCOVID:
		return SubdivisionTypeDepartment
	case SubdivisionCRA:
		return SubdivisionTypeProvince
	case SubdivisionCRC:
		return SubdivisionTypeProvince
	case SubdivisionCRG:
		return SubdivisionTypeProvince
	case SubdivisionCRH:
		return SubdivisionTypeProvince
	case SubdivisionCRL:
		return SubdivisionTypeProvince
	case SubdivisionCRP:
		return SubdivisionTypeProvince
	case SubdivisionCRSJ:
		return SubdivisionTypeProvince
	case SubdivisionCU01:
		return SubdivisionTypeProvince
	case SubdivisionCU02:
		return SubdivisionTypeProvince
	case SubdivisionCU03:
		return SubdivisionTypeProvince
	case SubdivisionCU04:
		return SubdivisionTypeProvince
	case SubdivisionCU05:
		return SubdivisionTypeProvince
	case SubdivisionCU06:
		return SubdivisionTypeProvince
	case SubdivisionCU07:
		return SubdivisionTypeProvince
	case SubdivisionCU08:
		return SubdivisionTypeProvince
	case SubdivisionCU09:
		return SubdivisionTypeProvince
	case SubdivisionCU10:
		return SubdivisionTypeProvince
	case SubdivisionCU11:
		return SubdivisionTypeProvince
	case SubdivisionCU12:
		return SubdivisionTypeProvince
	case SubdivisionCU13:
		return SubdivisionTypeProvince
	case SubdivisionCU14:
		return SubdivisionTypeProvince
	case SubdivisionCU99:
		return SubdivisionTypeSpecialMunicipality
	case SubdivisionCVB:
		return SubdivisionTypeGeographicalRegion
	case SubdivisionCVBR:
		return SubdivisionTypeMunicipality
	case SubdivisionCVBV:
		return SubdivisionTypeMunicipality
	case SubdivisionCVCA:
		return SubdivisionTypeMunicipality
	case SubdivisionCVCF:
		return SubdivisionTypeMunicipality
	case SubdivisionCVCR:
		return SubdivisionTypeMunicipality
	case SubdivisionCVMA:
		return SubdivisionTypeMunicipality
	case SubdivisionCVMO:
		return SubdivisionTypeMunicipality
	case SubdivisionCVPA:
		return SubdivisionTypeMunicipality
	case SubdivisionCVPN:
		return SubdivisionTypeMunicipality
	case SubdivisionCVPR:
		return SubdivisionTypeMunicipality
	case SubdivisionCVRB:
		return SubdivisionTypeMunicipality
	case SubdivisionCVRG:
		return SubdivisionTypeMunicipality
	case SubdivisionCVRS:
		return SubdivisionTypeMunicipality
	case SubdivisionCVS:
		return SubdivisionTypeGeographicalRegion
	case SubdivisionCVSD:
		return SubdivisionTypeMunicipality
	case SubdivisionCVSF:
		return SubdivisionTypeMunicipality
	case SubdivisionCVSL:
		return SubdivisionTypeMunicipality
	case SubdivisionCVSM:
		return SubdivisionTypeMunicipality
	case SubdivisionCVSO:
		return SubdivisionTypeMunicipality
	case SubdivisionCVSS:
		return SubdivisionTypeMunicipality
	case SubdivisionCVSV:
		return SubdivisionTypeMunicipality
	case SubdivisionCVTA:
		return SubdivisionTypeMunicipality
	case SubdivisionCVTS:
		return SubdivisionTypeMunicipality
	case SubdivisionCY01:
		return SubdivisionTypeDistrict
	case SubdivisionCY02:
		return SubdivisionTypeDistrict
	case SubdivisionCY03:
		return SubdivisionTypeDistrict
	case SubdivisionCY04:
		return SubdivisionTypeDistrict
	case SubdivisionCY05:
		return SubdivisionTypeDistrict
	case SubdivisionCY06:
		return SubdivisionTypeDistrict
	case SubdivisionCZ10:
		return SubdivisionTypeCapitalCity
	case SubdivisionCZ101:
		return SubdivisionTypeDistrict
	case SubdivisionCZ102:
		return SubdivisionTypeDistrict
	case SubdivisionCZ103:
		return SubdivisionTypeDistrict
	case SubdivisionCZ104:
		return SubdivisionTypeDistrict
	case SubdivisionCZ105:
		return SubdivisionTypeDistrict
	case SubdivisionCZ106:
		return SubdivisionTypeDistrict
	case SubdivisionCZ107:
		return SubdivisionTypeDistrict
	case SubdivisionCZ108:
		return SubdivisionTypeDistrict
	case SubdivisionCZ109:
		return SubdivisionTypeDistrict
	case SubdivisionCZ110:
		return SubdivisionTypeDistrict
	case SubdivisionCZ111:
		return SubdivisionTypeDistrict
	case SubdivisionCZ112:
		return SubdivisionTypeDistrict
	case SubdivisionCZ113:
		return SubdivisionTypeDistrict
	case SubdivisionCZ114:
		return SubdivisionTypeDistrict
	case SubdivisionCZ115:
		return SubdivisionTypeDistrict
	case SubdivisionCZ116:
		return SubdivisionTypeDistrict
	case SubdivisionCZ117:
		return SubdivisionTypeDistrict
	case SubdivisionCZ118:
		return SubdivisionTypeDistrict
	case SubdivisionCZ119:
		return SubdivisionTypeDistrict
	case SubdivisionCZ120:
		return SubdivisionTypeDistrict
	case SubdivisionCZ121:
		return SubdivisionTypeDistrict
	case SubdivisionCZ122:
		return SubdivisionTypeDistrict
	case SubdivisionCZ20:
		return SubdivisionTypeRegion
	case SubdivisionCZ201:
		return SubdivisionTypeDistrict
	case SubdivisionCZ202:
		return SubdivisionTypeDistrict
	case SubdivisionCZ203:
		return SubdivisionTypeDistrict
	case SubdivisionCZ204:
		return SubdivisionTypeDistrict
	case SubdivisionCZ205:
		return SubdivisionTypeDistrict
	case SubdivisionCZ206:
		return SubdivisionTypeDistrict
	case SubdivisionCZ207:
		return SubdivisionTypeDistrict
	case SubdivisionCZ208:
		return SubdivisionTypeDistrict
	case SubdivisionCZ209:
		return SubdivisionTypeDistrict
	case SubdivisionCZ20A:
		return SubdivisionTypeDistrict
	case SubdivisionCZ20B:
		return SubdivisionTypeDistrict
	case SubdivisionCZ20C:
		return SubdivisionTypeDistrict
	case SubdivisionCZ31:
		return SubdivisionTypeRegion
	case SubdivisionCZ311:
		return SubdivisionTypeDistrict
	case SubdivisionCZ312:
		return SubdivisionTypeDistrict
	case SubdivisionCZ313:
		return SubdivisionTypeDistrict
	case SubdivisionCZ314:
		return SubdivisionTypeDistrict
	case SubdivisionCZ315:
		return SubdivisionTypeDistrict
	case SubdivisionCZ316:
		return SubdivisionTypeDistrict
	case SubdivisionCZ317:
		return SubdivisionTypeDistrict
	case SubdivisionCZ32:
		return SubdivisionTypeRegion
	case SubdivisionCZ321:
		return SubdivisionTypeDistrict
	case SubdivisionCZ322:
		return SubdivisionTypeDistrict
	case SubdivisionCZ323:
		return SubdivisionTypeDistrict
	case SubdivisionCZ324:
		return SubdivisionTypeDistrict
	case SubdivisionCZ325:
		return SubdivisionTypeDistrict
	case SubdivisionCZ326:
		return SubdivisionTypeDistrict
	case SubdivisionCZ327:
		return SubdivisionTypeDistrict
	case SubdivisionCZ41:
		return SubdivisionTypeRegion
	case SubdivisionCZ411:
		return SubdivisionTypeDistrict
	case SubdivisionCZ412:
		return SubdivisionTypeDistrict
	case SubdivisionCZ413:
		return SubdivisionTypeDistrict
	case SubdivisionCZ42:
		return SubdivisionTypeRegion
	case SubdivisionCZ421:
		return SubdivisionTypeDistrict
	case SubdivisionCZ422:
		return SubdivisionTypeDistrict
	case SubdivisionCZ423:
		return SubdivisionTypeDistrict
	case SubdivisionCZ424:
		return SubdivisionTypeDistrict
	case SubdivisionCZ425:
		return SubdivisionTypeDistrict
	case SubdivisionCZ426:
		return SubdivisionTypeDistrict
	case SubdivisionCZ427:
		return SubdivisionTypeDistrict
	case SubdivisionCZ51:
		return SubdivisionTypeRegion
	case SubdivisionCZ511:
		return SubdivisionTypeDistrict
	case SubdivisionCZ512:
		return SubdivisionTypeDistrict
	case SubdivisionCZ513:
		return SubdivisionTypeDistrict
	case SubdivisionCZ514:
		return SubdivisionTypeDistrict
	case SubdivisionCZ52:
		return SubdivisionTypeRegion
	case SubdivisionCZ521:
		return SubdivisionTypeDistrict
	case SubdivisionCZ522:
		return SubdivisionTypeDistrict
	case SubdivisionCZ523:
		return SubdivisionTypeDistrict
	case SubdivisionCZ524:
		return SubdivisionTypeDistrict
	case SubdivisionCZ525:
		return SubdivisionTypeDistrict
	case SubdivisionCZ53:
		return SubdivisionTypeRegion
	case SubdivisionCZ531:
		return SubdivisionTypeDistrict
	case SubdivisionCZ532:
		return SubdivisionTypeDistrict
	case SubdivisionCZ533:
		return SubdivisionTypeDistrict
	case SubdivisionCZ534:
		return SubdivisionTypeDistrict
	case SubdivisionCZ63:
		return SubdivisionTypeRegion
	case SubdivisionCZ631:
		return SubdivisionTypeDistrict
	case SubdivisionCZ632:
		return SubdivisionTypeDistrict
	case SubdivisionCZ633:
		return SubdivisionTypeDistrict
	case SubdivisionCZ634:
		return SubdivisionTypeDistrict
	case SubdivisionCZ635:
		return SubdivisionTypeDistrict
	case SubdivisionCZ64:
		return SubdivisionTypeRegion
	case SubdivisionCZ641:
		return SubdivisionTypeDistrict
	case SubdivisionCZ642:
		return SubdivisionTypeDistrict
	case SubdivisionCZ643:
		return SubdivisionTypeDistrict
	case SubdivisionCZ644:
		return SubdivisionTypeDistrict
	case SubdivisionCZ645:
		return SubdivisionTypeDistrict
	case SubdivisionCZ646:
		return SubdivisionTypeDistrict
	case SubdivisionCZ647:
		return SubdivisionTypeDistrict
	case SubdivisionCZ71:
		return SubdivisionTypeRegion
	case SubdivisionCZ711:
		return SubdivisionTypeDistrict
	case SubdivisionCZ712:
		return SubdivisionTypeDistrict
	case SubdivisionCZ713:
		return SubdivisionTypeDistrict
	case SubdivisionCZ714:
		return SubdivisionTypeDistrict
	case SubdivisionCZ715:
		return SubdivisionTypeDistrict
	case SubdivisionCZ72:
		return SubdivisionTypeRegion
	case SubdivisionCZ721:
		return SubdivisionTypeDistrict
	case SubdivisionCZ722:
		return SubdivisionTypeDistrict
	case SubdivisionCZ723:
		return SubdivisionTypeDistrict
	case SubdivisionCZ724:
		return SubdivisionTypeDistrict
	case SubdivisionCZ80:
		return SubdivisionTypeRegion
	case SubdivisionCZ801:
		return SubdivisionTypeDistrict
	case SubdivisionCZ802:
		return SubdivisionTypeDistrict
	case SubdivisionCZ803:
		return SubdivisionTypeDistrict
	case SubdivisionCZ804:
		return SubdivisionTypeDistrict
	case SubdivisionCZ805:
		return SubdivisionTypeDistrict
	case SubdivisionCZ806:
		return SubdivisionTypeDistrict
	case SubdivisionDEBB:
		return SubdivisionTypeState
	case SubdivisionDEBE:
		return SubdivisionTypeState
	case SubdivisionDEBW:
		return SubdivisionTypeState
	case SubdivisionDEBY:
		return SubdivisionTypeState
	case SubdivisionDEHB:
		return SubdivisionTypeState
	case SubdivisionDEHE:
		return SubdivisionTypeState
	case SubdivisionDEHH:
		return SubdivisionTypeState
	case SubdivisionDEMV:
		return SubdivisionTypeState
	case SubdivisionDENI:
		return SubdivisionTypeState
	case SubdivisionDENW:
		return SubdivisionTypeState
	case SubdivisionDERP:
		return SubdivisionTypeState
	case SubdivisionDESH:
		return SubdivisionTypeState
	case SubdivisionDESL:
		return SubdivisionTypeState
	case SubdivisionDESN:
		return SubdivisionTypeState
	case SubdivisionDEST:
		return SubdivisionTypeState
	case SubdivisionDETH:
		return SubdivisionTypeState
	case SubdivisionDJAR:
		return SubdivisionTypeRegion
	case SubdivisionDJAS:
		return SubdivisionTypeRegion
	case SubdivisionDJDI:
		return SubdivisionTypeRegion
	case SubdivisionDJDJ:
		return SubdivisionTypeCity
	case SubdivisionDJOB:
		return SubdivisionTypeRegion
	case SubdivisionDJTA:
		return SubdivisionTypeRegion
	case SubdivisionDK81:
		return SubdivisionTypeRegion
	case SubdivisionDK82:
		return SubdivisionTypeRegion
	case SubdivisionDK83:
		return SubdivisionTypeRegion
	case SubdivisionDK84:
		return SubdivisionTypeRegion
	case SubdivisionDK85:
		return SubdivisionTypeRegion
	case SubdivisionDM01:
		return SubdivisionTypeParish
	case SubdivisionDM02:
		return SubdivisionTypeParish
	case SubdivisionDM03:
		return SubdivisionTypeParish
	case SubdivisionDM04:
		return SubdivisionTypeParish
	case SubdivisionDM05:
		return SubdivisionTypeParish
	case SubdivisionDM06:
		return SubdivisionTypeParish
	case SubdivisionDM07:
		return SubdivisionTypeParish
	case SubdivisionDM08:
		return SubdivisionTypeParish
	case SubdivisionDM09:
		return SubdivisionTypeParish
	case SubdivisionDM10:
		return SubdivisionTypeParish
	case SubdivisionDO01:
		return SubdivisionTypeDistrict
	case SubdivisionDO02:
		return SubdivisionTypeProvince
	case SubdivisionDO03:
		return SubdivisionTypeProvince
	case SubdivisionDO04:
		return SubdivisionTypeProvince
	case SubdivisionDO05:
		return SubdivisionTypeProvince
	case SubdivisionDO06:
		return SubdivisionTypeProvince
	case SubdivisionDO07:
		return SubdivisionTypeProvince
	case SubdivisionDO08:
		return SubdivisionTypeProvince
	case SubdivisionDO09:
		return SubdivisionTypeProvince
	case SubdivisionDO10:
		return SubdivisionTypeProvince
	case SubdivisionDO11:
		return SubdivisionTypeProvince
	case SubdivisionDO12:
		return SubdivisionTypeProvince
	case SubdivisionDO13:
		return SubdivisionTypeProvince
	case SubdivisionDO14:
		return SubdivisionTypeProvince
	case SubdivisionDO15:
		return SubdivisionTypeProvince
	case SubdivisionDO16:
		return SubdivisionTypeProvince
	case SubdivisionDO17:
		return SubdivisionTypeProvince
	case SubdivisionDO18:
		return SubdivisionTypeProvince
	case SubdivisionDO19:
		return SubdivisionTypeProvince
	case SubdivisionDO20:
		return SubdivisionTypeProvince
	case SubdivisionDO21:
		return SubdivisionTypeProvince
	case SubdivisionDO22:
		return SubdivisionTypeProvince
	case SubdivisionDO23:
		return SubdivisionTypeProvince
	case SubdivisionDO24:
		return SubdivisionTypeProvince
	case SubdivisionDO25:
		return SubdivisionTypeProvince
	case SubdivisionDO26:
		return SubdivisionTypeProvince
	case SubdivisionDO27:
		return SubdivisionTypeProvince
	case SubdivisionDO28:
		return SubdivisionTypeProvince
	case SubdivisionDO29:
		return SubdivisionTypeProvince
	case SubdivisionDO30:
		return SubdivisionTypeProvince
	case SubdivisionDZ01:
		return SubdivisionTypeProvince
	case SubdivisionDZ02:
		return SubdivisionTypeProvince
	case SubdivisionDZ03:
		return SubdivisionTypeProvince
	case SubdivisionDZ04:
		return SubdivisionTypeProvince
	case SubdivisionDZ05:
		return SubdivisionTypeProvince
	case SubdivisionDZ06:
		return SubdivisionTypeProvince
	case SubdivisionDZ07:
		return SubdivisionTypeProvince
	case SubdivisionDZ08:
		return SubdivisionTypeProvince
	case SubdivisionDZ09:
		return SubdivisionTypeProvince
	case SubdivisionDZ10:
		return SubdivisionTypeProvince
	case SubdivisionDZ11:
		return SubdivisionTypeProvince
	case SubdivisionDZ12:
		return SubdivisionTypeProvince
	case SubdivisionDZ13:
		return SubdivisionTypeProvince
	case SubdivisionDZ14:
		return SubdivisionTypeProvince
	case SubdivisionDZ15:
		return SubdivisionTypeProvince
	case SubdivisionDZ16:
		return SubdivisionTypeProvince
	case SubdivisionDZ17:
		return SubdivisionTypeProvince
	case SubdivisionDZ18:
		return SubdivisionTypeProvince
	case SubdivisionDZ19:
		return SubdivisionTypeProvince
	case SubdivisionDZ20:
		return SubdivisionTypeProvince
	case SubdivisionDZ21:
		return SubdivisionTypeProvince
	case SubdivisionDZ22:
		return SubdivisionTypeProvince
	case SubdivisionDZ23:
		return SubdivisionTypeProvince
	case SubdivisionDZ24:
		return SubdivisionTypeProvince
	case SubdivisionDZ25:
		return SubdivisionTypeProvince
	case SubdivisionDZ26:
		return SubdivisionTypeProvince
	case SubdivisionDZ27:
		return SubdivisionTypeProvince
	case SubdivisionDZ28:
		return SubdivisionTypeProvince
	case SubdivisionDZ29:
		return SubdivisionTypeProvince
	case SubdivisionDZ30:
		return SubdivisionTypeProvince
	case SubdivisionDZ31:
		return SubdivisionTypeProvince
	case SubdivisionDZ32:
		return SubdivisionTypeProvince
	case SubdivisionDZ33:
		return SubdivisionTypeProvince
	case SubdivisionDZ34:
		return SubdivisionTypeProvince
	case SubdivisionDZ35:
		return SubdivisionTypeProvince
	case SubdivisionDZ36:
		return SubdivisionTypeProvince
	case SubdivisionDZ37:
		return SubdivisionTypeProvince
	case SubdivisionDZ38:
		return SubdivisionTypeProvince
	case SubdivisionDZ39:
		return SubdivisionTypeProvince
	case SubdivisionDZ40:
		return SubdivisionTypeProvince
	case SubdivisionDZ41:
		return SubdivisionTypeProvince
	case SubdivisionDZ42:
		return SubdivisionTypeProvince
	case SubdivisionDZ43:
		return SubdivisionTypeProvince
	case SubdivisionDZ44:
		return SubdivisionTypeProvince
	case SubdivisionDZ45:
		return SubdivisionTypeProvince
	case SubdivisionDZ46:
		return SubdivisionTypeProvince
	case SubdivisionDZ47:
		return SubdivisionTypeProvince
	case SubdivisionDZ48:
		return SubdivisionTypeProvince
	case SubdivisionECA:
		return SubdivisionTypeProvince
	case SubdivisionECB:
		return SubdivisionTypeProvince
	case SubdivisionECC:
		return SubdivisionTypeProvince
	case SubdivisionECD:
		return SubdivisionTypeProvince
	case SubdivisionECE:
		return SubdivisionTypeProvince
	case SubdivisionECF:
		return SubdivisionTypeProvince
	case SubdivisionECG:
		return SubdivisionTypeProvince
	case SubdivisionECH:
		return SubdivisionTypeProvince
	case SubdivisionECI:
		return SubdivisionTypeProvince
	case SubdivisionECL:
		return SubdivisionTypeProvince
	case SubdivisionECM:
		return SubdivisionTypeProvince
	case SubdivisionECN:
		return SubdivisionTypeProvince
	case SubdivisionECO:
		return SubdivisionTypeProvince
	case SubdivisionECP:
		return SubdivisionTypeProvince
	case SubdivisionECR:
		return SubdivisionTypeProvince
	case SubdivisionECS:
		return SubdivisionTypeProvince
	case SubdivisionECSD:
		return SubdivisionTypeProvince
	case SubdivisionECSE:
		return SubdivisionTypeProvince
	case SubdivisionECT:
		return SubdivisionTypeProvince
	case SubdivisionECU:
		return SubdivisionTypeProvince
	case SubdivisionECW:
		return SubdivisionTypeProvince
	case SubdivisionECX:
		return SubdivisionTypeProvince
	case SubdivisionECY:
		return SubdivisionTypeProvince
	case SubdivisionECZ:
		return SubdivisionTypeProvince
	case SubdivisionEE37:
		return SubdivisionTypeCounty
	case SubdivisionEE39:
		return SubdivisionTypeCounty
	case SubdivisionEE44:
		return SubdivisionTypeCounty
	case SubdivisionEE49:
		return SubdivisionTypeCounty
	case SubdivisionEE51:
		return SubdivisionTypeCounty
	case SubdivisionEE57:
		return SubdivisionTypeCounty
	case SubdivisionEE59:
		return SubdivisionTypeCounty
	case SubdivisionEE65:
		return SubdivisionTypeCounty
	case SubdivisionEE67:
		return SubdivisionTypeCounty
	case SubdivisionEE70:
		return SubdivisionTypeCounty
	case SubdivisionEE74:
		return SubdivisionTypeCounty
	case SubdivisionEE78:
		return SubdivisionTypeCounty
	case SubdivisionEE82:
		return SubdivisionTypeCounty
	case SubdivisionEE84:
		return SubdivisionTypeCounty
	case SubdivisionEE86:
		return SubdivisionTypeCounty
	case SubdivisionEGALX:
		return SubdivisionTypeGovernorate
	case SubdivisionEGASN:
		return SubdivisionTypeGovernorate
	case SubdivisionEGAST:
		return SubdivisionTypeGovernorate
	case SubdivisionEGBA:
		return SubdivisionTypeGovernorate
	case SubdivisionEGBH:
		return SubdivisionTypeGovernorate
	case SubdivisionEGBNS:
		return SubdivisionTypeGovernorate
	case SubdivisionEGC:
		return SubdivisionTypeGovernorate
	case SubdivisionEGDK:
		return SubdivisionTypeGovernorate
	case SubdivisionEGDT:
		return SubdivisionTypeGovernorate
	case SubdivisionEGFYM:
		return SubdivisionTypeGovernorate
	case SubdivisionEGGH:
		return SubdivisionTypeGovernorate
	case SubdivisionEGGZ:
		return SubdivisionTypeGovernorate
	case SubdivisionEGHU:
		return SubdivisionTypeGovernorate
	case SubdivisionEGIS:
		return SubdivisionTypeGovernorate
	case SubdivisionEGJS:
		return SubdivisionTypeGovernorate
	case SubdivisionEGKB:
		return SubdivisionTypeGovernorate
	case SubdivisionEGKFS:
		return SubdivisionTypeGovernorate
	case SubdivisionEGKN:
		return SubdivisionTypeGovernorate
	case SubdivisionEGMN:
		return SubdivisionTypeGovernorate
	case SubdivisionEGMNF:
		return SubdivisionTypeGovernorate
	case SubdivisionEGMT:
		return SubdivisionTypeGovernorate
	case SubdivisionEGPTS:
		return SubdivisionTypeGovernorate
	case SubdivisionEGSHG:
		return SubdivisionTypeGovernorate
	case SubdivisionEGSHR:
		return SubdivisionTypeGovernorate
	case SubdivisionEGSIN:
		return SubdivisionTypeGovernorate
	case SubdivisionEGSU:
		return SubdivisionTypeGovernorate
	case SubdivisionEGSUZ:
		return SubdivisionTypeGovernorate
	case SubdivisionEGWAD:
		return SubdivisionTypeGovernorate
	case SubdivisionERAN:
		return SubdivisionTypeProvince
	case SubdivisionERDK:
		return SubdivisionTypeProvince
	case SubdivisionERDU:
		return SubdivisionTypeProvince
	case SubdivisionERGB:
		return SubdivisionTypeProvince
	case SubdivisionERMA:
		return SubdivisionTypeProvince
	case SubdivisionERSK:
		return SubdivisionTypeProvince
	case SubdivisionESA:
		return SubdivisionTypeProvince
	case SubdivisionESAB:
		return SubdivisionTypeProvince
	case SubdivisionESAL:
		return SubdivisionTypeProvince
	case SubdivisionESAN:
		return SubdivisionTypeAutonomousCommunity
	case SubdivisionESAR:
		return SubdivisionTypeAutonomousCommunity
	case SubdivisionESAS:
		return SubdivisionTypeAutonomousCommunity
	case SubdivisionESAV:
		return SubdivisionTypeProvince
	case SubdivisionESB:
		return SubdivisionTypeProvince
	case SubdivisionESBA:
		return SubdivisionTypeProvince
	case SubdivisionESBI:
		return SubdivisionTypeProvince
	case SubdivisionESBU:
		return SubdivisionTypeProvince
	case SubdivisionESC:
		return SubdivisionTypeProvince
	case SubdivisionESCA:
		return SubdivisionTypeProvince
	case SubdivisionESCB:
		return SubdivisionTypeAutonomousCommunity
	case SubdivisionESCC:
		return SubdivisionTypeProvince
	case SubdivisionESCE:
		return SubdivisionTypeAutonomousCity
	case SubdivisionESCL:
		return SubdivisionTypeAutonomousCommunity
	case SubdivisionESCM:
		return SubdivisionTypeAutonomousCommunity
	case SubdivisionESCN:
		return SubdivisionTypeAutonomousCommunity
	case SubdivisionESCO:
		return SubdivisionTypeProvince
	case SubdivisionESCR:
		return SubdivisionTypeProvince
	case SubdivisionESCS:
		return SubdivisionTypeProvince
	case SubdivisionESCT:
		return SubdivisionTypeAutonomousCommunity
	case SubdivisionESCU:
		return SubdivisionTypeProvince
	case SubdivisionESEX:
		return SubdivisionTypeAutonomousCommunity
	case SubdivisionESGA:
		return SubdivisionTypeAutonomousCommunity
	case SubdivisionESGC:
		return SubdivisionTypeProvince
	case SubdivisionESGI:
		return SubdivisionTypeProvince
	case SubdivisionESGR:
		return SubdivisionTypeProvince
	case SubdivisionESGU:
		return SubdivisionTypeProvince
	case SubdivisionESH:
		return SubdivisionTypeProvince
	case SubdivisionESHU:
		return SubdivisionTypeProvince
	case SubdivisionESIB:
		return SubdivisionTypeAutonomousCommunity
	case SubdivisionESJ:
		return SubdivisionTypeProvince
	case SubdivisionESL:
		return SubdivisionTypeProvince
	case SubdivisionESLE:
		return SubdivisionTypeProvince
	case SubdivisionESLO:
		return SubdivisionTypeProvince
	case SubdivisionESLU:
		return SubdivisionTypeProvince
	case SubdivisionESM:
		return SubdivisionTypeProvince
	case SubdivisionESMA:
		return SubdivisionTypeProvince
	case SubdivisionESMC:
		return SubdivisionTypeAutonomousCommunity
	case SubdivisionESMD:
		return SubdivisionTypeAutonomousCommunity
	case SubdivisionESML:
		return SubdivisionTypeAutonomousCity
	case SubdivisionESMU:
		return SubdivisionTypeProvince
	case SubdivisionESNA:
		return SubdivisionTypeProvince
	case SubdivisionESNC:
		return SubdivisionTypeAutonomousCommunity
	case SubdivisionESO:
		return SubdivisionTypeProvince
	case SubdivisionESOR:
		return SubdivisionTypeProvince
	case SubdivisionESP:
		return SubdivisionTypeProvince
	case SubdivisionESPM:
		return SubdivisionTypeProvince
	case SubdivisionESPO:
		return SubdivisionTypeProvince
	case SubdivisionESPV:
		return SubdivisionTypeAutonomousCommunity
	case SubdivisionESRI:
		return SubdivisionTypeAutonomousCommunity
	case SubdivisionESS:
		return SubdivisionTypeProvince
	case SubdivisionESSA:
		return SubdivisionTypeProvince
	case SubdivisionESSE:
		return SubdivisionTypeProvince
	case SubdivisionESSG:
		return SubdivisionTypeProvince
	case SubdivisionESSO:
		return SubdivisionTypeProvince
	case SubdivisionESSS:
		return SubdivisionTypeProvince
	case SubdivisionEST:
		return SubdivisionTypeProvince
	case SubdivisionESTE:
		return SubdivisionTypeProvince
	case SubdivisionESTF:
		return SubdivisionTypeProvince
	case SubdivisionESTO:
		return SubdivisionTypeProvince
	case SubdivisionESV:
		return SubdivisionTypeProvince
	case SubdivisionESVA:
		return SubdivisionTypeProvince
	case SubdivisionESVC:
		return SubdivisionTypeAutonomousCommunity
	case SubdivisionESVI:
		return SubdivisionTypeProvince
	case SubdivisionESZ:
		return SubdivisionTypeProvince
	case SubdivisionESZA:
		return SubdivisionTypeProvince
	case SubdivisionETAA:
		return SubdivisionTypeAdministration
	case SubdivisionETAF:
		return SubdivisionTypeState
	case SubdivisionETAM:
		return SubdivisionTypeState
	case SubdivisionETBE:
		return SubdivisionTypeState
	case SubdivisionETDD:
		return SubdivisionTypeAdministration
	case SubdivisionETGA:
		return SubdivisionTypeState
	case SubdivisionETHA:
		return SubdivisionTypeState
	case SubdivisionETOR:
		return SubdivisionTypeState
	case SubdivisionETSN:
		return SubdivisionTypeState
	case SubdivisionETSO:
		return SubdivisionTypeState
	case SubdivisionETTI:
		return SubdivisionTypeState
	case SubdivisionFI01:
		return SubdivisionTypeRegion
	case SubdivisionFI02:
		return SubdivisionTypeRegion
	case SubdivisionFI03:
		return SubdivisionTypeRegion
	case SubdivisionFI04:
		return SubdivisionTypeRegion
	case SubdivisionFI05:
		return SubdivisionTypeRegion
	case SubdivisionFI06:
		return SubdivisionTypeRegion
	case SubdivisionFI07:
		return SubdivisionTypeRegion
	case SubdivisionFI08:
		return SubdivisionTypeRegion
	case SubdivisionFI09:
		return SubdivisionTypeRegion
	case SubdivisionFI10:
		return SubdivisionTypeRegion
	case SubdivisionFI11:
		return SubdivisionTypeRegion
	case SubdivisionFI12:
		return SubdivisionTypeRegion
	case SubdivisionFI13:
		return SubdivisionTypeRegion
	case SubdivisionFI14:
		return SubdivisionTypeRegion
	case SubdivisionFI15:
		return SubdivisionTypeRegion
	case SubdivisionFI16:
		return SubdivisionTypeRegion
	case SubdivisionFI17:
		return SubdivisionTypeRegion
	case SubdivisionFI18:
		return SubdivisionTypeRegion
	case SubdivisionFI19:
		return SubdivisionTypeRegion
	case SubdivisionFJC:
		return SubdivisionTypeDivision
	case SubdivisionFJE:
		return SubdivisionTypeDivision
	case SubdivisionFJN:
		return SubdivisionTypeDivision
	case SubdivisionFJR:
		return SubdivisionTypeDependency
	case SubdivisionFJW:
		return SubdivisionTypeDivision
	case SubdivisionFMKSA:
		return SubdivisionTypeState
	case SubdivisionFMPNI:
		return SubdivisionTypeState
	case SubdivisionFMTRK:
		return SubdivisionTypeState
	case SubdivisionFMYAP:
		return SubdivisionTypeState
	case SubdivisionFR01:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR02:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR03:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR04:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR05:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR06:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR07:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR08:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR09:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR10:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR11:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR12:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR13:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR14:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR15:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR16:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR17:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR18:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR19:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR21:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR22:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR23:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR24:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR25:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR26:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR27:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR28:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR29:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR2A:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR2B:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR30:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR31:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR32:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR33:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR34:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR35:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR36:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR37:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR38:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR39:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR40:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR41:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR42:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR43:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR44:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR45:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR46:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR47:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR48:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR49:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR50:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR51:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR52:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR53:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR54:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR55:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR56:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR57:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR58:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR59:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR60:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR61:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR62:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR63:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR64:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR65:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR66:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR67:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR68:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR69:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR70:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR71:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR72:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR73:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR74:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR75:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR76:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR77:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR78:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR79:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR80:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR81:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR82:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR83:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR84:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR85:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR86:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR87:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR88:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR89:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR90:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR91:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR92:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR93:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR94:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFR95:
		return SubdivisionTypeMetropolitanDepartment
	case SubdivisionFRARA:
		return SubdivisionTypeMetropolitanRegion
	case SubdivisionFRBFC:
		return SubdivisionTypeMetropolitanRegion
	case SubdivisionFRBL:
		return SubdivisionTypeOverseasTerritorialCollectivity
	case SubdivisionFRBRE:
		return SubdivisionTypeMetropolitanRegion
	case SubdivisionFRCOR:
		return SubdivisionTypeMetropolitanRegion
	case SubdivisionFRCP:
		return SubdivisionTypeDependency
	case SubdivisionFRCVL:
		return SubdivisionTypeMetropolitanRegion
	case SubdivisionFRGES:
		return SubdivisionTypeMetropolitanRegion
	case SubdivisionFRGF:
		return SubdivisionTypeOverseasTerritorialCollectivity
	case SubdivisionFRGP:
		return SubdivisionTypeOverseasDepartment
	case SubdivisionFRGUA:
		return SubdivisionTypeOverseasRegion
	case SubdivisionFRHDF:
		return SubdivisionTypeMetropolitanRegion
	case SubdivisionFRIDF:
		return SubdivisionTypeMetropolitanRegion
	case SubdivisionFRLRE:
		return SubdivisionTypeOverseasRegion
	case SubdivisionFRMAY:
		return SubdivisionTypeOverseasRegion
	case SubdivisionFRMF:
		return SubdivisionTypeOverseasTerritorialCollectivity
	case SubdivisionFRMQ:
		return SubdivisionTypeOverseasTerritorialCollectivity
	case SubdivisionFRNAQ:
		return SubdivisionTypeMetropolitanRegion
	case SubdivisionFRNC:
		return SubdivisionTypeOverseasTerritorialCollectivity
	case SubdivisionFRNOR:
		return SubdivisionTypeMetropolitanRegion
	case SubdivisionFROCC:
		return SubdivisionTypeMetropolitanRegion
	case SubdivisionFRPAC:
		return SubdivisionTypeMetropolitanRegion
	case SubdivisionFRPDL:
		return SubdivisionTypeMetropolitanRegion
	case SubdivisionFRPF:
		return SubdivisionTypeOverseasTerritorialCollectivity
	case SubdivisionFRPM:
		return SubdivisionTypeOverseasTerritorialCollectivity
	case SubdivisionFRRE:
		return SubdivisionTypeOverseasDepartment
	case SubdivisionFRTF:
		return SubdivisionTypeOverseasTerritorialCollectivity
	case SubdivisionFRWF:
		return SubdivisionTypeOverseasTerritorialCollectivity
	case SubdivisionFRYT:
		return SubdivisionTypeOverseasDepartment
	case SubdivisionGA1:
		return SubdivisionTypeProvince
	case SubdivisionGA2:
		return SubdivisionTypeProvince
	case SubdivisionGA3:
		return SubdivisionTypeProvince
	case SubdivisionGA4:
		return SubdivisionTypeProvince
	case SubdivisionGA5:
		return SubdivisionTypeProvince
	case SubdivisionGA6:
		return SubdivisionTypeProvince
	case SubdivisionGA7:
		return SubdivisionTypeProvince
	case SubdivisionGA8:
		return SubdivisionTypeProvince
	case SubdivisionGA9:
		return SubdivisionTypeProvince
	case SubdivisionGBABC:
		return SubdivisionTypeDistrict
	case SubdivisionGBABD:
		return SubdivisionTypeCouncilArea
	case SubdivisionGBABE:
		return SubdivisionTypeCouncilArea
	case SubdivisionGBAGB:
		return SubdivisionTypeCouncilArea
	case SubdivisionGBAGY:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBAND:
		return SubdivisionTypeDistrict
	case SubdivisionGBANN:
		return SubdivisionTypeDistrict
	case SubdivisionGBANS:
		return SubdivisionTypeCouncilArea
	case SubdivisionGBBAS:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBBBD:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBBDF:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBBDG:
		return SubdivisionTypeLondonBorough
	case SubdivisionGBBEN:
		return SubdivisionTypeLondonBorough
	case SubdivisionGBBEX:
		return SubdivisionTypeLondonBorough
	case SubdivisionGBBFS:
		return SubdivisionTypeDistrict
	case SubdivisionGBBGE:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBBGW:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBBIR:
		return SubdivisionTypeMetropolitanDistrict
	case SubdivisionGBBKM:
		return SubdivisionTypeTwoTierCounty
	case SubdivisionGBBMH:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBBNE:
		return SubdivisionTypeLondonBorough
	case SubdivisionGBBNH:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBBNS:
		return SubdivisionTypeMetropolitanDistrict
	case SubdivisionGBBOL:
		return SubdivisionTypeMetropolitanDistrict
	case SubdivisionGBBPL:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBBRC:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBBRD:
		return SubdivisionTypeMetropolitanDistrict
	case SubdivisionGBBRY:
		return SubdivisionTypeLondonBorough
	case SubdivisionGBBST:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBBUR:
		return SubdivisionTypeMetropolitanDistrict
	case SubdivisionGBCAM:
		return SubdivisionTypeTwoTierCounty
	case SubdivisionGBCAY:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBCBF:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBCCG:
		return SubdivisionTypeDistrict
	case SubdivisionGBCGN:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBCHE:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBCHW:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBCLD:
		return SubdivisionTypeMetropolitanDistrict
	case SubdivisionGBCLK:
		return SubdivisionTypeCouncilArea
	case SubdivisionGBCMA:
		return SubdivisionTypeTwoTierCounty
	case SubdivisionGBCMD:
		return SubdivisionTypeLondonBorough
	case SubdivisionGBCMN:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBCON:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBCOV:
		return SubdivisionTypeMetropolitanDistrict
	case SubdivisionGBCRF:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBCRY:
		return SubdivisionTypeLondonBorough
	case SubdivisionGBCWY:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBDAL:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBDBY:
		return SubdivisionTypeTwoTierCounty
	case SubdivisionGBDEN:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBDER:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBDEV:
		return SubdivisionTypeTwoTierCounty
	case SubdivisionGBDGY:
		return SubdivisionTypeCouncilArea
	case SubdivisionGBDNC:
		return SubdivisionTypeMetropolitanDistrict
	case SubdivisionGBDND:
		return SubdivisionTypeCouncilArea
	case SubdivisionGBDOR:
		return SubdivisionTypeTwoTierCounty
	case SubdivisionGBDRS:
		return SubdivisionTypeDistrict
	case SubdivisionGBDUD:
		return SubdivisionTypeMetropolitanDistrict
	case SubdivisionGBDUR:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBEAL:
		return SubdivisionTypeLondonBorough
	case SubdivisionGBEAW:
		return SubdivisionTypeNation
	case SubdivisionGBEAY:
		return SubdivisionTypeCouncilArea
	case SubdivisionGBEDH:
		return SubdivisionTypeCouncilArea
	case SubdivisionGBEDU:
		return SubdivisionTypeCouncilArea
	case SubdivisionGBELN:
		return SubdivisionTypeCouncilArea
	case SubdivisionGBELS:
		return SubdivisionTypeCouncilArea
	case SubdivisionGBENF:
		return SubdivisionTypeLondonBorough
	case SubdivisionGBENG:
		return SubdivisionTypeCountry
	case SubdivisionGBERW:
		return SubdivisionTypeCouncilArea
	case SubdivisionGBERY:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBESS:
		return SubdivisionTypeTwoTierCounty
	case SubdivisionGBESX:
		return SubdivisionTypeTwoTierCounty
	case SubdivisionGBFAL:
		return SubdivisionTypeCouncilArea
	case SubdivisionGBFIF:
		return SubdivisionTypeCouncilArea
	case SubdivisionGBFLN:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBFMO:
		return SubdivisionTypeDistrict
	case SubdivisionGBGAT:
		return SubdivisionTypeMetropolitanDistrict
	case SubdivisionGBGBN:
		return SubdivisionTypeNation
	case SubdivisionGBGLG:
		return SubdivisionTypeCouncilArea
	case SubdivisionGBGLS:
		return SubdivisionTypeTwoTierCounty
	case SubdivisionGBGRE:
		return SubdivisionTypeLondonBorough
	case SubdivisionGBGWN:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBHAL:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBHAM:
		return SubdivisionTypeTwoTierCounty
	case SubdivisionGBHAV:
		return SubdivisionTypeLondonBorough
	case SubdivisionGBHCK:
		return SubdivisionTypeLondonBorough
	case SubdivisionGBHEF:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBHIL:
		return SubdivisionTypeLondonBorough
	case SubdivisionGBHLD:
		return SubdivisionTypeCouncilArea
	case SubdivisionGBHMF:
		return SubdivisionTypeLondonBorough
	case SubdivisionGBHNS:
		return SubdivisionTypeLondonBorough
	case SubdivisionGBHPL:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBHRT:
		return SubdivisionTypeTwoTierCounty
	case SubdivisionGBHRW:
		return SubdivisionTypeLondonBorough
	case SubdivisionGBHRY:
		return SubdivisionTypeLondonBorough
	case SubdivisionGBIOS:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBIOW:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBISL:
		return SubdivisionTypeLondonBorough
	case SubdivisionGBIVC:
		return SubdivisionTypeCouncilArea
	case SubdivisionGBKEC:
		return SubdivisionTypeLondonBorough
	case SubdivisionGBKEN:
		return SubdivisionTypeTwoTierCounty
	case SubdivisionGBKHL:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBKIR:
		return SubdivisionTypeMetropolitanDistrict
	case SubdivisionGBKTT:
		return SubdivisionTypeLondonBorough
	case SubdivisionGBKWL:
		return SubdivisionTypeMetropolitanDistrict
	case SubdivisionGBLAN:
		return SubdivisionTypeTwoTierCounty
	case SubdivisionGBLBC:
		return SubdivisionTypeDistrict
	case SubdivisionGBLBH:
		return SubdivisionTypeLondonBorough
	case SubdivisionGBLCE:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBLDS:
		return SubdivisionTypeMetropolitanDistrict
	case SubdivisionGBLEC:
		return SubdivisionTypeTwoTierCounty
	case SubdivisionGBLEW:
		return SubdivisionTypeLondonBorough
	case SubdivisionGBLIN:
		return SubdivisionTypeTwoTierCounty
	case SubdivisionGBLIV:
		return SubdivisionTypeMetropolitanDistrict
	case SubdivisionGBLND:
		return SubdivisionTypeCityCorporation
	case SubdivisionGBLUT:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBMAN:
		return SubdivisionTypeMetropolitanDistrict
	case SubdivisionGBMDB:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBMDW:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBMEA:
		return SubdivisionTypeDistrict
	case SubdivisionGBMIK:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBMLN:
		return SubdivisionTypeCouncilArea
	case SubdivisionGBMON:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBMRT:
		return SubdivisionTypeLondonBorough
	case SubdivisionGBMRY:
		return SubdivisionTypeCouncilArea
	case SubdivisionGBMTY:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBMUL:
		return SubdivisionTypeDistrict
	case SubdivisionGBNAY:
		return SubdivisionTypeCouncilArea
	case SubdivisionGBNBL:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBNEL:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBNET:
		return SubdivisionTypeMetropolitanDistrict
	case SubdivisionGBNFK:
		return SubdivisionTypeTwoTierCounty
	case SubdivisionGBNGM:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBNIR:
		return SubdivisionTypeProvince
	case SubdivisionGBNLK:
		return SubdivisionTypeCouncilArea
	case SubdivisionGBNLN:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBNMD:
		return SubdivisionTypeDistrict
	case SubdivisionGBNSM:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBNTH:
		return SubdivisionTypeTwoTierCounty
	case SubdivisionGBNTL:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBNTT:
		return SubdivisionTypeTwoTierCounty
	case SubdivisionGBNTY:
		return SubdivisionTypeMetropolitanDistrict
	case SubdivisionGBNWM:
		return SubdivisionTypeLondonBorough
	case SubdivisionGBNWP:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBNYK:
		return SubdivisionTypeTwoTierCounty
	case SubdivisionGBOLD:
		return SubdivisionTypeMetropolitanDistrict
	case SubdivisionGBORK:
		return SubdivisionTypeCouncilArea
	case SubdivisionGBOXF:
		return SubdivisionTypeTwoTierCounty
	case SubdivisionGBPEM:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBPKN:
		return SubdivisionTypeCouncilArea
	case SubdivisionGBPLY:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBPOL:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBPOR:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBPOW:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBPTE:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBRCC:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBRCH:
		return SubdivisionTypeMetropolitanDistrict
	case SubdivisionGBRCT:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBRDB:
		return SubdivisionTypeLondonBorough
	case SubdivisionGBRDG:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBRFW:
		return SubdivisionTypeCouncilArea
	case SubdivisionGBRIC:
		return SubdivisionTypeLondonBorough
	case SubdivisionGBROT:
		return SubdivisionTypeMetropolitanDistrict
	case SubdivisionGBRUT:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBSAW:
		return SubdivisionTypeMetropolitanDistrict
	case SubdivisionGBSAY:
		return SubdivisionTypeCouncilArea
	case SubdivisionGBSCB:
		return SubdivisionTypeCouncilArea
	case SubdivisionGBSCT:
		return SubdivisionTypeCountry
	case SubdivisionGBSFK:
		return SubdivisionTypeTwoTierCounty
	case SubdivisionGBSFT:
		return SubdivisionTypeMetropolitanDistrict
	case SubdivisionGBSGC:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBSHF:
		return SubdivisionTypeMetropolitanDistrict
	case SubdivisionGBSHN:
		return SubdivisionTypeMetropolitanDistrict
	case SubdivisionGBSHR:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBSKP:
		return SubdivisionTypeMetropolitanDistrict
	case SubdivisionGBSLF:
		return SubdivisionTypeMetropolitanDistrict
	case SubdivisionGBSLG:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBSLK:
		return SubdivisionTypeCouncilArea
	case SubdivisionGBSND:
		return SubdivisionTypeMetropolitanDistrict
	case SubdivisionGBSOL:
		return SubdivisionTypeMetropolitanDistrict
	case SubdivisionGBSOM:
		return SubdivisionTypeTwoTierCounty
	case SubdivisionGBSOS:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBSRY:
		return SubdivisionTypeTwoTierCounty
	case SubdivisionGBSTE:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBSTG:
		return SubdivisionTypeCouncilArea
	case SubdivisionGBSTH:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBSTN:
		return SubdivisionTypeLondonBorough
	case SubdivisionGBSTS:
		return SubdivisionTypeTwoTierCounty
	case SubdivisionGBSTT:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBSTY:
		return SubdivisionTypeMetropolitanDistrict
	case SubdivisionGBSWA:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBSWD:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBSWK:
		return SubdivisionTypeLondonBorough
	case SubdivisionGBTAM:
		return SubdivisionTypeMetropolitanDistrict
	case SubdivisionGBTFW:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBTHR:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBTOB:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBTOF:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBTRF:
		return SubdivisionTypeMetropolitanDistrict
	case SubdivisionGBTWH:
		return SubdivisionTypeLondonBorough
	case SubdivisionGBUKM:
		return SubdivisionTypeNation
	case SubdivisionGBVGL:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBWAR:
		return SubdivisionTypeTwoTierCounty
	case SubdivisionGBWBK:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBWDU:
		return SubdivisionTypeCouncilArea
	case SubdivisionGBWFT:
		return SubdivisionTypeLondonBorough
	case SubdivisionGBWGN:
		return SubdivisionTypeMetropolitanDistrict
	case SubdivisionGBWIL:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBWKF:
		return SubdivisionTypeMetropolitanDistrict
	case SubdivisionGBWLL:
		return SubdivisionTypeMetropolitanDistrict
	case SubdivisionGBWLN:
		return SubdivisionTypeCouncilArea
	case SubdivisionGBWLS:
		return SubdivisionTypeCountry
	case SubdivisionGBWLV:
		return SubdivisionTypeMetropolitanDistrict
	case SubdivisionGBWND:
		return SubdivisionTypeLondonBorough
	case SubdivisionGBWNM:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBWOK:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBWOR:
		return SubdivisionTypeTwoTierCounty
	case SubdivisionGBWRL:
		return SubdivisionTypeMetropolitanDistrict
	case SubdivisionGBWRT:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBWRX:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBWSM:
		return SubdivisionTypeLondonBorough
	case SubdivisionGBWSX:
		return SubdivisionTypeTwoTierCounty
	case SubdivisionGBYOR:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionGBZET:
		return SubdivisionTypeCouncilArea
	case SubdivisionGD01:
		return SubdivisionTypeParish
	case SubdivisionGD02:
		return SubdivisionTypeParish
	case SubdivisionGD03:
		return SubdivisionTypeParish
	case SubdivisionGD04:
		return SubdivisionTypeParish
	case SubdivisionGD05:
		return SubdivisionTypeParish
	case SubdivisionGD06:
		return SubdivisionTypeParish
	case SubdivisionGD10:
		return SubdivisionTypeDependency
	case SubdivisionGEAB:
		return SubdivisionTypeAutonomousRepublic
	case SubdivisionGEAJ:
		return SubdivisionTypeAutonomousRepublic
	case SubdivisionGEGU:
		return SubdivisionTypeRegion
	case SubdivisionGEIM:
		return SubdivisionTypeRegion
	case SubdivisionGEKA:
		return SubdivisionTypeRegion
	case SubdivisionGEKK:
		return SubdivisionTypeRegion
	case SubdivisionGEMM:
		return SubdivisionTypeRegion
	case SubdivisionGERL:
		return SubdivisionTypeRegion
	case SubdivisionGESJ:
		return SubdivisionTypeRegion
	case SubdivisionGESK:
		return SubdivisionTypeRegion
	case SubdivisionGESZ:
		return SubdivisionTypeRegion
	case SubdivisionGETB:
		return SubdivisionTypeCity
	case SubdivisionGHAA:
		return SubdivisionTypeRegion
	case SubdivisionGHAH:
		return SubdivisionTypeRegion
	case SubdivisionGHBA:
		return SubdivisionTypeRegion
	case SubdivisionGHCP:
		return SubdivisionTypeRegion
	case SubdivisionGHEP:
		return SubdivisionTypeRegion
	case SubdivisionGHNP:
		return SubdivisionTypeRegion
	case SubdivisionGHTV:
		return SubdivisionTypeRegion
	case SubdivisionGHUE:
		return SubdivisionTypeRegion
	case SubdivisionGHUW:
		return SubdivisionTypeRegion
	case SubdivisionGHWP:
		return SubdivisionTypeRegion
	case SubdivisionGLKU:
		return SubdivisionTypeMunicipality
	case SubdivisionGLQA:
		return SubdivisionTypeMunicipality
	case SubdivisionGLQE:
		return SubdivisionTypeMunicipality
	case SubdivisionGLSM:
		return SubdivisionTypeMunicipality
	case SubdivisionGMB:
		return SubdivisionTypeCity
	case SubdivisionGML:
		return SubdivisionTypeDivision
	case SubdivisionGMM:
		return SubdivisionTypeDivision
	case SubdivisionGMN:
		return SubdivisionTypeDivision
	case SubdivisionGMU:
		return SubdivisionTypeDivision
	case SubdivisionGMW:
		return SubdivisionTypeDivision
	case SubdivisionGNB:
		return SubdivisionTypeGovernorate
	case SubdivisionGNBE:
		return SubdivisionTypePrefecture
	case SubdivisionGNBF:
		return SubdivisionTypePrefecture
	case SubdivisionGNBK:
		return SubdivisionTypePrefecture
	case SubdivisionGNC:
		return SubdivisionTypeSpecialZone
	case SubdivisionGNCO:
		return SubdivisionTypePrefecture
	case SubdivisionGND:
		return SubdivisionTypeGovernorate
	case SubdivisionGNDB:
		return SubdivisionTypePrefecture
	case SubdivisionGNDI:
		return SubdivisionTypePrefecture
	case SubdivisionGNDL:
		return SubdivisionTypePrefecture
	case SubdivisionGNDU:
		return SubdivisionTypePrefecture
	case SubdivisionGNF:
		return SubdivisionTypeGovernorate
	case SubdivisionGNFA:
		return SubdivisionTypePrefecture
	case SubdivisionGNFO:
		return SubdivisionTypePrefecture
	case SubdivisionGNFR:
		return SubdivisionTypePrefecture
	case SubdivisionGNGA:
		return SubdivisionTypePrefecture
	case SubdivisionGNGU:
		return SubdivisionTypePrefecture
	case SubdivisionGNK:
		return SubdivisionTypeGovernorate
	case SubdivisionGNKA:
		return SubdivisionTypePrefecture
	case SubdivisionGNKB:
		return SubdivisionTypePrefecture
	case SubdivisionGNKD:
		return SubdivisionTypePrefecture
	case SubdivisionGNKE:
		return SubdivisionTypePrefecture
	case SubdivisionGNKN:
		return SubdivisionTypePrefecture
	case SubdivisionGNKO:
		return SubdivisionTypePrefecture
	case SubdivisionGNKS:
		return SubdivisionTypePrefecture
	case SubdivisionGNL:
		return SubdivisionTypeGovernorate
	case SubdivisionGNLA:
		return SubdivisionTypePrefecture
	case SubdivisionGNLE:
		return SubdivisionTypePrefecture
	case SubdivisionGNLO:
		return SubdivisionTypePrefecture
	case SubdivisionGNM:
		return SubdivisionTypeGovernorate
	case SubdivisionGNMC:
		return SubdivisionTypePrefecture
	case SubdivisionGNMD:
		return SubdivisionTypePrefecture
	case SubdivisionGNML:
		return SubdivisionTypePrefecture
	case SubdivisionGNMM:
		return SubdivisionTypePrefecture
	case SubdivisionGNN:
		return SubdivisionTypeGovernorate
	case SubdivisionGNNZ:
		return SubdivisionTypePrefecture
	case SubdivisionGNPI:
		return SubdivisionTypePrefecture
	case SubdivisionGNSI:
		return SubdivisionTypePrefecture
	case SubdivisionGNTE:
		return SubdivisionTypePrefecture
	case SubdivisionGNTO:
		return SubdivisionTypePrefecture
	case SubdivisionGNYO:
		return SubdivisionTypePrefecture
	case SubdivisionGQAN:
		return SubdivisionTypeProvince
	case SubdivisionGQBN:
		return SubdivisionTypeProvince
	case SubdivisionGQBS:
		return SubdivisionTypeProvince
	case SubdivisionGQC:
		return SubdivisionTypeRegion
	case SubdivisionGQCS:
		return SubdivisionTypeProvince
	case SubdivisionGQI:
		return SubdivisionTypeRegion
	case SubdivisionGQKN:
		return SubdivisionTypeProvince
	case SubdivisionGQLI:
		return SubdivisionTypeProvince
	case SubdivisionGQWN:
		return SubdivisionTypeProvince
	case SubdivisionGR01:
		return SubdivisionTypeDepartment
	case SubdivisionGR03:
		return SubdivisionTypeDepartment
	case SubdivisionGR04:
		return SubdivisionTypeDepartment
	case SubdivisionGR05:
		return SubdivisionTypeDepartment
	case SubdivisionGR06:
		return SubdivisionTypeDepartment
	case SubdivisionGR07:
		return SubdivisionTypeDepartment
	case SubdivisionGR11:
		return SubdivisionTypeDepartment
	case SubdivisionGR12:
		return SubdivisionTypeDepartment
	case SubdivisionGR13:
		return SubdivisionTypeDepartment
	case SubdivisionGR14:
		return SubdivisionTypeDepartment
	case SubdivisionGR15:
		return SubdivisionTypeDepartment
	case SubdivisionGR16:
		return SubdivisionTypeDepartment
	case SubdivisionGR17:
		return SubdivisionTypeDepartment
	case SubdivisionGR21:
		return SubdivisionTypeDepartment
	case SubdivisionGR22:
		return SubdivisionTypeDepartment
	case SubdivisionGR23:
		return SubdivisionTypeDepartment
	case SubdivisionGR24:
		return SubdivisionTypeDepartment
	case SubdivisionGR31:
		return SubdivisionTypeDepartment
	case SubdivisionGR32:
		return SubdivisionTypeDepartment
	case SubdivisionGR33:
		return SubdivisionTypeDepartment
	case SubdivisionGR34:
		return SubdivisionTypeDepartment
	case SubdivisionGR41:
		return SubdivisionTypeDepartment
	case SubdivisionGR42:
		return SubdivisionTypeDepartment
	case SubdivisionGR43:
		return SubdivisionTypeDepartment
	case SubdivisionGR44:
		return SubdivisionTypeDepartment
	case SubdivisionGR51:
		return SubdivisionTypeDepartment
	case SubdivisionGR52:
		return SubdivisionTypeDepartment
	case SubdivisionGR53:
		return SubdivisionTypeDepartment
	case SubdivisionGR54:
		return SubdivisionTypeDepartment
	case SubdivisionGR55:
		return SubdivisionTypeDepartment
	case SubdivisionGR56:
		return SubdivisionTypeDepartment
	case SubdivisionGR57:
		return SubdivisionTypeDepartment
	case SubdivisionGR58:
		return SubdivisionTypeDepartment
	case SubdivisionGR59:
		return SubdivisionTypeDepartment
	case SubdivisionGR61:
		return SubdivisionTypeDepartment
	case SubdivisionGR62:
		return SubdivisionTypeDepartment
	case SubdivisionGR63:
		return SubdivisionTypeDepartment
	case SubdivisionGR64:
		return SubdivisionTypeDepartment
	case SubdivisionGR69:
		return SubdivisionTypeSelfGovernedPart
	case SubdivisionGR71:
		return SubdivisionTypeDepartment
	case SubdivisionGR72:
		return SubdivisionTypeDepartment
	case SubdivisionGR73:
		return SubdivisionTypeDepartment
	case SubdivisionGR81:
		return SubdivisionTypeDepartment
	case SubdivisionGR82:
		return SubdivisionTypeDepartment
	case SubdivisionGR83:
		return SubdivisionTypeDepartment
	case SubdivisionGR84:
		return SubdivisionTypeDepartment
	case SubdivisionGR85:
		return SubdivisionTypeDepartment
	case SubdivisionGR91:
		return SubdivisionTypeDepartment
	case SubdivisionGR92:
		return SubdivisionTypeDepartment
	case SubdivisionGR93:
		return SubdivisionTypeDepartment
	case SubdivisionGR94:
		return SubdivisionTypeDepartment
	case SubdivisionGRA:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionGRA1:
		return SubdivisionTypeDepartment
	case SubdivisionGRB:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionGRC:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionGRD:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionGRE:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionGRF:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionGRG:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionGRH:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionGRI:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionGRJ:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionGRK:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionGRL:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionGRM:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionGTAV:
		return SubdivisionTypeDepartment
	case SubdivisionGTBV:
		return SubdivisionTypeDepartment
	case SubdivisionGTCM:
		return SubdivisionTypeDepartment
	case SubdivisionGTCQ:
		return SubdivisionTypeDepartment
	case SubdivisionGTES:
		return SubdivisionTypeDepartment
	case SubdivisionGTGU:
		return SubdivisionTypeDepartment
	case SubdivisionGTHU:
		return SubdivisionTypeDepartment
	case SubdivisionGTIZ:
		return SubdivisionTypeDepartment
	case SubdivisionGTJA:
		return SubdivisionTypeDepartment
	case SubdivisionGTJU:
		return SubdivisionTypeDepartment
	case SubdivisionGTPE:
		return SubdivisionTypeDepartment
	case SubdivisionGTPR:
		return SubdivisionTypeDepartment
	case SubdivisionGTQC:
		return SubdivisionTypeDepartment
	case SubdivisionGTQZ:
		return SubdivisionTypeDepartment
	case SubdivisionGTRE:
		return SubdivisionTypeDepartment
	case SubdivisionGTSA:
		return SubdivisionTypeDepartment
	case SubdivisionGTSM:
		return SubdivisionTypeDepartment
	case SubdivisionGTSO:
		return SubdivisionTypeDepartment
	case SubdivisionGTSR:
		return SubdivisionTypeDepartment
	case SubdivisionGTSU:
		return SubdivisionTypeDepartment
	case SubdivisionGTTO:
		return SubdivisionTypeDepartment
	case SubdivisionGTZA:
		return SubdivisionTypeDepartment
	case SubdivisionGWBA:
		return SubdivisionTypeRegion
	case SubdivisionGWBL:
		return SubdivisionTypeRegion
	case SubdivisionGWBM:
		return SubdivisionTypeRegion
	case SubdivisionGWBS:
		return SubdivisionTypeAutonomousSector
	case SubdivisionGWCA:
		return SubdivisionTypeRegion
	case SubdivisionGWGA:
		return SubdivisionTypeRegion
	case SubdivisionGWL:
		return SubdivisionTypeProvince
	case SubdivisionGWN:
		return SubdivisionTypeProvince
	case SubdivisionGWOI:
		return SubdivisionTypeRegion
	case SubdivisionGWQU:
		return SubdivisionTypeRegion
	case SubdivisionGWS:
		return SubdivisionTypeProvince
	case SubdivisionGWTO:
		return SubdivisionTypeRegion
	case SubdivisionGYBA:
		return SubdivisionTypeRegion
	case SubdivisionGYCU:
		return SubdivisionTypeRegion
	case SubdivisionGYDE:
		return SubdivisionTypeRegion
	case SubdivisionGYEB:
		return SubdivisionTypeRegion
	case SubdivisionGYES:
		return SubdivisionTypeRegion
	case SubdivisionGYMA:
		return SubdivisionTypeRegion
	case SubdivisionGYPM:
		return SubdivisionTypeRegion
	case SubdivisionGYPT:
		return SubdivisionTypeRegion
	case SubdivisionGYUD:
		return SubdivisionTypeRegion
	case SubdivisionGYUT:
		return SubdivisionTypeRegion
	case SubdivisionHNAT:
		return SubdivisionTypeDepartment
	case SubdivisionHNCH:
		return SubdivisionTypeDepartment
	case SubdivisionHNCL:
		return SubdivisionTypeDepartment
	case SubdivisionHNCM:
		return SubdivisionTypeDepartment
	case SubdivisionHNCP:
		return SubdivisionTypeDepartment
	case SubdivisionHNCR:
		return SubdivisionTypeDepartment
	case SubdivisionHNEP:
		return SubdivisionTypeDepartment
	case SubdivisionHNFM:
		return SubdivisionTypeDepartment
	case SubdivisionHNGD:
		return SubdivisionTypeDepartment
	case SubdivisionHNIB:
		return SubdivisionTypeDepartment
	case SubdivisionHNIN:
		return SubdivisionTypeDepartment
	case SubdivisionHNLE:
		return SubdivisionTypeDepartment
	case SubdivisionHNLP:
		return SubdivisionTypeDepartment
	case SubdivisionHNOC:
		return SubdivisionTypeDepartment
	case SubdivisionHNOL:
		return SubdivisionTypeDepartment
	case SubdivisionHNSB:
		return SubdivisionTypeDepartment
	case SubdivisionHNVA:
		return SubdivisionTypeDepartment
	case SubdivisionHNYO:
		return SubdivisionTypeDepartment
	case SubdivisionHR01:
		return SubdivisionTypeCounty
	case SubdivisionHR02:
		return SubdivisionTypeCounty
	case SubdivisionHR03:
		return SubdivisionTypeCounty
	case SubdivisionHR04:
		return SubdivisionTypeCounty
	case SubdivisionHR05:
		return SubdivisionTypeCounty
	case SubdivisionHR06:
		return SubdivisionTypeCounty
	case SubdivisionHR07:
		return SubdivisionTypeCounty
	case SubdivisionHR08:
		return SubdivisionTypeCounty
	case SubdivisionHR09:
		return SubdivisionTypeCounty
	case SubdivisionHR10:
		return SubdivisionTypeCounty
	case SubdivisionHR11:
		return SubdivisionTypeCounty
	case SubdivisionHR12:
		return SubdivisionTypeCounty
	case SubdivisionHR13:
		return SubdivisionTypeCounty
	case SubdivisionHR14:
		return SubdivisionTypeCounty
	case SubdivisionHR15:
		return SubdivisionTypeCounty
	case SubdivisionHR16:
		return SubdivisionTypeCounty
	case SubdivisionHR17:
		return SubdivisionTypeCounty
	case SubdivisionHR18:
		return SubdivisionTypeCounty
	case SubdivisionHR19:
		return SubdivisionTypeCounty
	case SubdivisionHR20:
		return SubdivisionTypeCounty
	case SubdivisionHR21:
		return SubdivisionTypeCity
	case SubdivisionHTAR:
		return SubdivisionTypeDepartment
	case SubdivisionHTCE:
		return SubdivisionTypeDepartment
	case SubdivisionHTGA:
		return SubdivisionTypeDepartment
	case SubdivisionHTND:
		return SubdivisionTypeDepartment
	case SubdivisionHTNE:
		return SubdivisionTypeDepartment
	case SubdivisionHTNO:
		return SubdivisionTypeDepartment
	case SubdivisionHTOU:
		return SubdivisionTypeDepartment
	case SubdivisionHTSD:
		return SubdivisionTypeDepartment
	case SubdivisionHTSE:
		return SubdivisionTypeDepartment
	case SubdivisionHUBA:
		return SubdivisionTypeCounty
	case SubdivisionHUBC:
		return SubdivisionTypeCityWithCountyRights
	case SubdivisionHUBE:
		return SubdivisionTypeCounty
	case SubdivisionHUBK:
		return SubdivisionTypeCounty
	case SubdivisionHUBU:
		return SubdivisionTypeCapitalCity
	case SubdivisionHUBZ:
		return SubdivisionTypeCounty
	case SubdivisionHUCS:
		return SubdivisionTypeCounty
	case SubdivisionHUDE:
		return SubdivisionTypeCityWithCountyRights
	case SubdivisionHUDU:
		return SubdivisionTypeCityWithCountyRights
	case SubdivisionHUEG:
		return SubdivisionTypeCityWithCountyRights
	case SubdivisionHUER:
		return SubdivisionTypeCityWithCountyRights
	case SubdivisionHUFE:
		return SubdivisionTypeCounty
	case SubdivisionHUGS:
		return SubdivisionTypeCounty
	case SubdivisionHUGY:
		return SubdivisionTypeCityWithCountyRights
	case SubdivisionHUHB:
		return SubdivisionTypeCounty
	case SubdivisionHUHE:
		return SubdivisionTypeCounty
	case SubdivisionHUHV:
		return SubdivisionTypeCityWithCountyRights
	case SubdivisionHUJN:
		return SubdivisionTypeCounty
	case SubdivisionHUKE:
		return SubdivisionTypeCounty
	case SubdivisionHUKM:
		return SubdivisionTypeCityWithCountyRights
	case SubdivisionHUKV:
		return SubdivisionTypeCityWithCountyRights
	case SubdivisionHUMI:
		return SubdivisionTypeCityWithCountyRights
	case SubdivisionHUNK:
		return SubdivisionTypeCityWithCountyRights
	case SubdivisionHUNO:
		return SubdivisionTypeCounty
	case SubdivisionHUNY:
		return SubdivisionTypeCityWithCountyRights
	case SubdivisionHUPE:
		return SubdivisionTypeCounty
	case SubdivisionHUPS:
		return SubdivisionTypeCityWithCountyRights
	case SubdivisionHUSD:
		return SubdivisionTypeCityWithCountyRights
	case SubdivisionHUSF:
		return SubdivisionTypeCityWithCountyRights
	case SubdivisionHUSH:
		return SubdivisionTypeCityWithCountyRights
	case SubdivisionHUSK:
		return SubdivisionTypeCityWithCountyRights
	case SubdivisionHUSN:
		return SubdivisionTypeCityWithCountyRights
	case SubdivisionHUSO:
		return SubdivisionTypeCounty
	case SubdivisionHUSS:
		return SubdivisionTypeCityWithCountyRights
	case SubdivisionHUST:
		return SubdivisionTypeCityWithCountyRights
	case SubdivisionHUSZ:
		return SubdivisionTypeCounty
	case SubdivisionHUTB:
		return SubdivisionTypeCityWithCountyRights
	case SubdivisionHUTO:
		return SubdivisionTypeCounty
	case SubdivisionHUVA:
		return SubdivisionTypeCounty
	case SubdivisionHUVE:
		return SubdivisionTypeCounty
	case SubdivisionHUVM:
		return SubdivisionTypeCityWithCountyRights
	case SubdivisionHUZA:
		return SubdivisionTypeCounty
	case SubdivisionHUZE:
		return SubdivisionTypeCityWithCountyRights
	case SubdivisionIDAC:
		return SubdivisionTypeAutonomousProvince
	case SubdivisionIDBA:
		return SubdivisionTypeProvince
	case SubdivisionIDBB:
		return SubdivisionTypeProvince
	case SubdivisionIDBE:
		return SubdivisionTypeProvince
	case SubdivisionIDBT:
		return SubdivisionTypeProvince
	case SubdivisionIDGO:
		return SubdivisionTypeProvince
	case SubdivisionIDIJ:
		return SubdivisionTypeGeographicalUnit
	case SubdivisionIDJA:
		return SubdivisionTypeProvince
	case SubdivisionIDJB:
		return SubdivisionTypeProvince
	case SubdivisionIDJI:
		return SubdivisionTypeProvince
	case SubdivisionIDJK:
		return SubdivisionTypeSpecialDistrict
	case SubdivisionIDJT:
		return SubdivisionTypeProvince
	case SubdivisionIDJW:
		return SubdivisionTypeGeographicalUnit
	case SubdivisionIDKA:
		return SubdivisionTypeGeographicalUnit
	case SubdivisionIDKB:
		return SubdivisionTypeProvince
	case SubdivisionIDKI:
		return SubdivisionTypeProvince
	case SubdivisionIDKR:
		return SubdivisionTypeProvince
	case SubdivisionIDKS:
		return SubdivisionTypeProvince
	case SubdivisionIDKT:
		return SubdivisionTypeProvince
	case SubdivisionIDLA:
		return SubdivisionTypeProvince
	case SubdivisionIDMA:
		return SubdivisionTypeProvince
	case SubdivisionIDML:
		return SubdivisionTypeGeographicalUnit
	case SubdivisionIDMU:
		return SubdivisionTypeProvince
	case SubdivisionIDNB:
		return SubdivisionTypeProvince
	case SubdivisionIDNT:
		return SubdivisionTypeProvince
	case SubdivisionIDNU:
		return SubdivisionTypeGeographicalUnit
	case SubdivisionIDPA:
		return SubdivisionTypeProvince
	case SubdivisionIDPB:
		return SubdivisionTypeProvince
	case SubdivisionIDRI:
		return SubdivisionTypeProvince
	case SubdivisionIDSA:
		return SubdivisionTypeProvince
	case SubdivisionIDSB:
		return SubdivisionTypeProvince
	case SubdivisionIDSG:
		return SubdivisionTypeProvince
	case SubdivisionIDSL:
		return SubdivisionTypeGeographicalUnit
	case SubdivisionIDSM:
		return SubdivisionTypeGeographicalUnit
	case SubdivisionIDSN:
		return SubdivisionTypeProvince
	case SubdivisionIDSR:
		return SubdivisionTypeProvince
	case SubdivisionIDSS:
		return SubdivisionTypeProvince
	case SubdivisionIDST:
		return SubdivisionTypeProvince
	case SubdivisionIDSU:
		return SubdivisionTypeProvince
	case SubdivisionIDYO:
		return SubdivisionTypeSpecialRegion
	case SubdivisionIEC:
		return SubdivisionTypeProvince
	case SubdivisionIECE:
		return SubdivisionTypeCounty
	case SubdivisionIECN:
		return SubdivisionTypeCounty
	case SubdivisionIECO:
		return SubdivisionTypeCounty
	case SubdivisionIECW:
		return SubdivisionTypeCounty
	case SubdivisionIED:
		return SubdivisionTypeCounty
	case SubdivisionIEDL:
		return SubdivisionTypeCounty
	case SubdivisionIEG:
		return SubdivisionTypeCounty
	case SubdivisionIEKE:
		return SubdivisionTypeCounty
	case SubdivisionIEKK:
		return SubdivisionTypeCounty
	case SubdivisionIEKY:
		return SubdivisionTypeCounty
	case SubdivisionIEL:
		return SubdivisionTypeProvince
	case SubdivisionIELD:
		return SubdivisionTypeCounty
	case SubdivisionIELH:
		return SubdivisionTypeCounty
	case SubdivisionIELK:
		return SubdivisionTypeCounty
	case SubdivisionIELM:
		return SubdivisionTypeCounty
	case SubdivisionIELS:
		return SubdivisionTypeCounty
	case SubdivisionIEM:
		return SubdivisionTypeProvince
	case SubdivisionIEMH:
		return SubdivisionTypeCounty
	case SubdivisionIEMN:
		return SubdivisionTypeCounty
	case SubdivisionIEMO:
		return SubdivisionTypeCounty
	case SubdivisionIEOY:
		return SubdivisionTypeCounty
	case SubdivisionIERN:
		return SubdivisionTypeCounty
	case SubdivisionIESO:
		return SubdivisionTypeCounty
	case SubdivisionIETA:
		return SubdivisionTypeCounty
	case SubdivisionIEU:
		return SubdivisionTypeProvince
	case SubdivisionIEWD:
		return SubdivisionTypeCounty
	case SubdivisionIEWH:
		return SubdivisionTypeCounty
	case SubdivisionIEWW:
		return SubdivisionTypeCounty
	case SubdivisionIEWX:
		return SubdivisionTypeCounty
	case SubdivisionILD:
		return SubdivisionTypeDistrict
	case SubdivisionILHA:
		return SubdivisionTypeDistrict
	case SubdivisionILJM:
		return SubdivisionTypeDistrict
	case SubdivisionILM:
		return SubdivisionTypeDistrict
	case SubdivisionILTA:
		return SubdivisionTypeDistrict
	case SubdivisionILZ:
		return SubdivisionTypeDistrict
	case SubdivisionINAN:
		return SubdivisionTypeUnionTerritory
	case SubdivisionINAP:
		return SubdivisionTypeState
	case SubdivisionINAR:
		return SubdivisionTypeState
	case SubdivisionINAS:
		return SubdivisionTypeState
	case SubdivisionINBR:
		return SubdivisionTypeState
	case SubdivisionINCH:
		return SubdivisionTypeUnionTerritory
	case SubdivisionINCT:
		return SubdivisionTypeState
	case SubdivisionINDD:
		return SubdivisionTypeUnionTerritory
	case SubdivisionINDL:
		return SubdivisionTypeUnionTerritory
	case SubdivisionINDN:
		return SubdivisionTypeUnionTerritory
	case SubdivisionINGA:
		return SubdivisionTypeState
	case SubdivisionINGJ:
		return SubdivisionTypeState
	case SubdivisionINHP:
		return SubdivisionTypeState
	case SubdivisionINHR:
		return SubdivisionTypeState
	case SubdivisionINJH:
		return SubdivisionTypeState
	case SubdivisionINJK:
		return SubdivisionTypeState
	case SubdivisionINKA:
		return SubdivisionTypeState
	case SubdivisionINKL:
		return SubdivisionTypeState
	case SubdivisionINLD:
		return SubdivisionTypeUnionTerritory
	case SubdivisionINMH:
		return SubdivisionTypeState
	case SubdivisionINML:
		return SubdivisionTypeState
	case SubdivisionINMN:
		return SubdivisionTypeState
	case SubdivisionINMP:
		return SubdivisionTypeState
	case SubdivisionINMZ:
		return SubdivisionTypeState
	case SubdivisionINNL:
		return SubdivisionTypeState
	case SubdivisionINOR:
		return SubdivisionTypeState
	case SubdivisionINPB:
		return SubdivisionTypeState
	case SubdivisionINPY:
		return SubdivisionTypeUnionTerritory
	case SubdivisionINRJ:
		return SubdivisionTypeState
	case SubdivisionINSK:
		return SubdivisionTypeState
	case SubdivisionINTG:
		return SubdivisionTypeState
	case SubdivisionINTN:
		return SubdivisionTypeState
	case SubdivisionINTR:
		return SubdivisionTypeState
	case SubdivisionINUP:
		return SubdivisionTypeState
	case SubdivisionINUT:
		return SubdivisionTypeState
	case SubdivisionINWB:
		return SubdivisionTypeState
	case SubdivisionIQAN:
		return SubdivisionTypeGovernorate
	case SubdivisionIQAR:
		return SubdivisionTypeGovernorate
	case SubdivisionIQBA:
		return SubdivisionTypeGovernorate
	case SubdivisionIQBB:
		return SubdivisionTypeGovernorate
	case SubdivisionIQBG:
		return SubdivisionTypeGovernorate
	case SubdivisionIQDA:
		return SubdivisionTypeGovernorate
	case SubdivisionIQDI:
		return SubdivisionTypeGovernorate
	case SubdivisionIQDQ:
		return SubdivisionTypeGovernorate
	case SubdivisionIQKA:
		return SubdivisionTypeGovernorate
	case SubdivisionIQMA:
		return SubdivisionTypeGovernorate
	case SubdivisionIQMU:
		return SubdivisionTypeGovernorate
	case SubdivisionIQNA:
		return SubdivisionTypeGovernorate
	case SubdivisionIQNI:
		return SubdivisionTypeGovernorate
	case SubdivisionIQQA:
		return SubdivisionTypeGovernorate
	case SubdivisionIQSD:
		return SubdivisionTypeGovernorate
	case SubdivisionIQSW:
		return SubdivisionTypeGovernorate
	case SubdivisionIQTS:
		return SubdivisionTypeGovernorate
	case SubdivisionIQWA:
		return SubdivisionTypeGovernorate
	case SubdivisionIR01:
		return SubdivisionTypeProvince
	case SubdivisionIR02:
		return SubdivisionTypeProvince
	case SubdivisionIR03:
		return SubdivisionTypeProvince
	case SubdivisionIR04:
		return SubdivisionTypeProvince
	case SubdivisionIR05:
		return SubdivisionTypeProvince
	case SubdivisionIR06:
		return SubdivisionTypeProvince
	case SubdivisionIR07:
		return SubdivisionTypeProvince
	case SubdivisionIR08:
		return SubdivisionTypeProvince
	case SubdivisionIR10:
		return SubdivisionTypeProvince
	case SubdivisionIR11:
		return SubdivisionTypeProvince
	case SubdivisionIR12:
		return SubdivisionTypeProvince
	case SubdivisionIR13:
		return SubdivisionTypeProvince
	case SubdivisionIR14:
		return SubdivisionTypeProvince
	case SubdivisionIR15:
		return SubdivisionTypeProvince
	case SubdivisionIR16:
		return SubdivisionTypeProvince
	case SubdivisionIR17:
		return SubdivisionTypeProvince
	case SubdivisionIR18:
		return SubdivisionTypeProvince
	case SubdivisionIR19:
		return SubdivisionTypeProvince
	case SubdivisionIR20:
		return SubdivisionTypeProvince
	case SubdivisionIR21:
		return SubdivisionTypeProvince
	case SubdivisionIR22:
		return SubdivisionTypeProvince
	case SubdivisionIR23:
		return SubdivisionTypeProvince
	case SubdivisionIR24:
		return SubdivisionTypeProvince
	case SubdivisionIR25:
		return SubdivisionTypeProvince
	case SubdivisionIR26:
		return SubdivisionTypeProvince
	case SubdivisionIR27:
		return SubdivisionTypeProvince
	case SubdivisionIR28:
		return SubdivisionTypeProvince
	case SubdivisionIR29:
		return SubdivisionTypeProvince
	case SubdivisionIR30:
		return SubdivisionTypeProvince
	case SubdivisionIR31:
		return SubdivisionTypeProvince
	case SubdivisionIS0:
		return SubdivisionTypeCity
	case SubdivisionIS1:
		return SubdivisionTypeRegion
	case SubdivisionIS2:
		return SubdivisionTypeRegion
	case SubdivisionIS3:
		return SubdivisionTypeRegion
	case SubdivisionIS4:
		return SubdivisionTypeRegion
	case SubdivisionIS5:
		return SubdivisionTypeRegion
	case SubdivisionIS6:
		return SubdivisionTypeRegion
	case SubdivisionIS7:
		return SubdivisionTypeRegion
	case SubdivisionIS8:
		return SubdivisionTypeRegion
	case SubdivisionIT21:
		return SubdivisionTypeRegion
	case SubdivisionIT23:
		return SubdivisionTypeRegion
	case SubdivisionIT25:
		return SubdivisionTypeRegion
	case SubdivisionIT32:
		return SubdivisionTypeRegion
	case SubdivisionIT34:
		return SubdivisionTypeRegion
	case SubdivisionIT36:
		return SubdivisionTypeRegion
	case SubdivisionIT42:
		return SubdivisionTypeRegion
	case SubdivisionIT45:
		return SubdivisionTypeRegion
	case SubdivisionIT52:
		return SubdivisionTypeRegion
	case SubdivisionIT55:
		return SubdivisionTypeRegion
	case SubdivisionIT57:
		return SubdivisionTypeRegion
	case SubdivisionIT62:
		return SubdivisionTypeRegion
	case SubdivisionIT65:
		return SubdivisionTypeRegion
	case SubdivisionIT67:
		return SubdivisionTypeRegion
	case SubdivisionIT72:
		return SubdivisionTypeRegion
	case SubdivisionIT75:
		return SubdivisionTypeRegion
	case SubdivisionIT77:
		return SubdivisionTypeRegion
	case SubdivisionIT78:
		return SubdivisionTypeRegion
	case SubdivisionIT82:
		return SubdivisionTypeRegion
	case SubdivisionIT88:
		return SubdivisionTypeRegion
	case SubdivisionITAG:
		return SubdivisionTypeProvince
	case SubdivisionITAL:
		return SubdivisionTypeProvince
	case SubdivisionITAN:
		return SubdivisionTypeProvince
	case SubdivisionITAO:
		return SubdivisionTypeProvince
	case SubdivisionITAP:
		return SubdivisionTypeProvince
	case SubdivisionITAQ:
		return SubdivisionTypeProvince
	case SubdivisionITAR:
		return SubdivisionTypeProvince
	case SubdivisionITAT:
		return SubdivisionTypeProvince
	case SubdivisionITAV:
		return SubdivisionTypeProvince
	case SubdivisionITBA:
		return SubdivisionTypeProvince
	case SubdivisionITBG:
		return SubdivisionTypeProvince
	case SubdivisionITBI:
		return SubdivisionTypeProvince
	case SubdivisionITBL:
		return SubdivisionTypeProvince
	case SubdivisionITBN:
		return SubdivisionTypeProvince
	case SubdivisionITBO:
		return SubdivisionTypeProvince
	case SubdivisionITBR:
		return SubdivisionTypeProvince
	case SubdivisionITBS:
		return SubdivisionTypeProvince
	case SubdivisionITBT:
		return SubdivisionTypeProvince
	case SubdivisionITBZ:
		return SubdivisionTypeProvince
	case SubdivisionITCA:
		return SubdivisionTypeProvince
	case SubdivisionITCB:
		return SubdivisionTypeProvince
	case SubdivisionITCE:
		return SubdivisionTypeProvince
	case SubdivisionITCH:
		return SubdivisionTypeProvince
	case SubdivisionITCI:
		return SubdivisionTypeProvince
	case SubdivisionITCL:
		return SubdivisionTypeProvince
	case SubdivisionITCN:
		return SubdivisionTypeProvince
	case SubdivisionITCO:
		return SubdivisionTypeProvince
	case SubdivisionITCR:
		return SubdivisionTypeProvince
	case SubdivisionITCS:
		return SubdivisionTypeProvince
	case SubdivisionITCT:
		return SubdivisionTypeProvince
	case SubdivisionITCZ:
		return SubdivisionTypeProvince
	case SubdivisionITEN:
		return SubdivisionTypeProvince
	case SubdivisionITFC:
		return SubdivisionTypeProvince
	case SubdivisionITFE:
		return SubdivisionTypeProvince
	case SubdivisionITFG:
		return SubdivisionTypeProvince
	case SubdivisionITFI:
		return SubdivisionTypeProvince
	case SubdivisionITFM:
		return SubdivisionTypeProvince
	case SubdivisionITFR:
		return SubdivisionTypeProvince
	case SubdivisionITGE:
		return SubdivisionTypeProvince
	case SubdivisionITGO:
		return SubdivisionTypeProvince
	case SubdivisionITGR:
		return SubdivisionTypeProvince
	case SubdivisionITIM:
		return SubdivisionTypeProvince
	case SubdivisionITIS:
		return SubdivisionTypeProvince
	case SubdivisionITKR:
		return SubdivisionTypeProvince
	case SubdivisionITLC:
		return SubdivisionTypeProvince
	case SubdivisionITLE:
		return SubdivisionTypeProvince
	case SubdivisionITLI:
		return SubdivisionTypeProvince
	case SubdivisionITLO:
		return SubdivisionTypeProvince
	case SubdivisionITLT:
		return SubdivisionTypeProvince
	case SubdivisionITLU:
		return SubdivisionTypeProvince
	case SubdivisionITMB:
		return SubdivisionTypeProvince
	case SubdivisionITMC:
		return SubdivisionTypeProvince
	case SubdivisionITME:
		return SubdivisionTypeProvince
	case SubdivisionITMI:
		return SubdivisionTypeProvince
	case SubdivisionITMN:
		return SubdivisionTypeProvince
	case SubdivisionITMO:
		return SubdivisionTypeProvince
	case SubdivisionITMS:
		return SubdivisionTypeProvince
	case SubdivisionITMT:
		return SubdivisionTypeProvince
	case SubdivisionITNA:
		return SubdivisionTypeProvince
	case SubdivisionITNO:
		return SubdivisionTypeProvince
	case SubdivisionITNU:
		return SubdivisionTypeProvince
	case SubdivisionITOG:
		return SubdivisionTypeProvince
	case SubdivisionITOR:
		return SubdivisionTypeProvince
	case SubdivisionITOT:
		return SubdivisionTypeProvince
	case SubdivisionITPA:
		return SubdivisionTypeProvince
	case SubdivisionITPC:
		return SubdivisionTypeProvince
	case SubdivisionITPD:
		return SubdivisionTypeProvince
	case SubdivisionITPE:
		return SubdivisionTypeProvince
	case SubdivisionITPG:
		return SubdivisionTypeProvince
	case SubdivisionITPI:
		return SubdivisionTypeProvince
	case SubdivisionITPN:
		return SubdivisionTypeProvince
	case SubdivisionITPO:
		return SubdivisionTypeProvince
	case SubdivisionITPR:
		return SubdivisionTypeProvince
	case SubdivisionITPT:
		return SubdivisionTypeProvince
	case SubdivisionITPU:
		return SubdivisionTypeProvince
	case SubdivisionITPV:
		return SubdivisionTypeProvince
	case SubdivisionITPZ:
		return SubdivisionTypeProvince
	case SubdivisionITRA:
		return SubdivisionTypeProvince
	case SubdivisionITRC:
		return SubdivisionTypeProvince
	case SubdivisionITRE:
		return SubdivisionTypeProvince
	case SubdivisionITRG:
		return SubdivisionTypeProvince
	case SubdivisionITRI:
		return SubdivisionTypeProvince
	case SubdivisionITRM:
		return SubdivisionTypeProvince
	case SubdivisionITRN:
		return SubdivisionTypeProvince
	case SubdivisionITRO:
		return SubdivisionTypeProvince
	case SubdivisionITSA:
		return SubdivisionTypeProvince
	case SubdivisionITSI:
		return SubdivisionTypeProvince
	case SubdivisionITSO:
		return SubdivisionTypeProvince
	case SubdivisionITSP:
		return SubdivisionTypeProvince
	case SubdivisionITSR:
		return SubdivisionTypeProvince
	case SubdivisionITSS:
		return SubdivisionTypeProvince
	case SubdivisionITSV:
		return SubdivisionTypeProvince
	case SubdivisionITTA:
		return SubdivisionTypeProvince
	case SubdivisionITTE:
		return SubdivisionTypeProvince
	case SubdivisionITTN:
		return SubdivisionTypeProvince
	case SubdivisionITTO:
		return SubdivisionTypeProvince
	case SubdivisionITTP:
		return SubdivisionTypeProvince
	case SubdivisionITTR:
		return SubdivisionTypeProvince
	case SubdivisionITTS:
		return SubdivisionTypeProvince
	case SubdivisionITTV:
		return SubdivisionTypeProvince
	case SubdivisionITUD:
		return SubdivisionTypeProvince
	case SubdivisionITVA:
		return SubdivisionTypeProvince
	case SubdivisionITVB:
		return SubdivisionTypeProvince
	case SubdivisionITVC:
		return SubdivisionTypeProvince
	case SubdivisionITVE:
		return SubdivisionTypeProvince
	case SubdivisionITVI:
		return SubdivisionTypeProvince
	case SubdivisionITVR:
		return SubdivisionTypeProvince
	case SubdivisionITVS:
		return SubdivisionTypeProvince
	case SubdivisionITVT:
		return SubdivisionTypeProvince
	case SubdivisionITVV:
		return SubdivisionTypeProvince
	case SubdivisionJM01:
		return SubdivisionTypeParish
	case SubdivisionJM02:
		return SubdivisionTypeParish
	case SubdivisionJM03:
		return SubdivisionTypeParish
	case SubdivisionJM04:
		return SubdivisionTypeParish
	case SubdivisionJM05:
		return SubdivisionTypeParish
	case SubdivisionJM06:
		return SubdivisionTypeParish
	case SubdivisionJM07:
		return SubdivisionTypeParish
	case SubdivisionJM08:
		return SubdivisionTypeParish
	case SubdivisionJM09:
		return SubdivisionTypeParish
	case SubdivisionJM10:
		return SubdivisionTypeParish
	case SubdivisionJM11:
		return SubdivisionTypeParish
	case SubdivisionJM12:
		return SubdivisionTypeParish
	case SubdivisionJM13:
		return SubdivisionTypeParish
	case SubdivisionJM14:
		return SubdivisionTypeParish
	case SubdivisionJOAJ:
		return SubdivisionTypeGovernorate
	case SubdivisionJOAM:
		return SubdivisionTypeGovernorate
	case SubdivisionJOAQ:
		return SubdivisionTypeGovernorate
	case SubdivisionJOAT:
		return SubdivisionTypeGovernorate
	case SubdivisionJOAZ:
		return SubdivisionTypeGovernorate
	case SubdivisionJOBA:
		return SubdivisionTypeGovernorate
	case SubdivisionJOIR:
		return SubdivisionTypeGovernorate
	case SubdivisionJOJA:
		return SubdivisionTypeGovernorate
	case SubdivisionJOKA:
		return SubdivisionTypeGovernorate
	case SubdivisionJOMA:
		return SubdivisionTypeGovernorate
	case SubdivisionJOMD:
		return SubdivisionTypeGovernorate
	case SubdivisionJOMN:
		return SubdivisionTypeGovernorate
	case SubdivisionJP01:
		return SubdivisionTypePrefecture
	case SubdivisionJP02:
		return SubdivisionTypePrefecture
	case SubdivisionJP03:
		return SubdivisionTypePrefecture
	case SubdivisionJP04:
		return SubdivisionTypePrefecture
	case SubdivisionJP05:
		return SubdivisionTypePrefecture
	case SubdivisionJP06:
		return SubdivisionTypePrefecture
	case SubdivisionJP07:
		return SubdivisionTypePrefecture
	case SubdivisionJP08:
		return SubdivisionTypePrefecture
	case SubdivisionJP09:
		return SubdivisionTypePrefecture
	case SubdivisionJP10:
		return SubdivisionTypePrefecture
	case SubdivisionJP11:
		return SubdivisionTypePrefecture
	case SubdivisionJP12:
		return SubdivisionTypePrefecture
	case SubdivisionJP13:
		return SubdivisionTypePrefecture
	case SubdivisionJP14:
		return SubdivisionTypePrefecture
	case SubdivisionJP15:
		return SubdivisionTypePrefecture
	case SubdivisionJP16:
		return SubdivisionTypePrefecture
	case SubdivisionJP17:
		return SubdivisionTypePrefecture
	case SubdivisionJP18:
		return SubdivisionTypePrefecture
	case SubdivisionJP19:
		return SubdivisionTypePrefecture
	case SubdivisionJP20:
		return SubdivisionTypePrefecture
	case SubdivisionJP21:
		return SubdivisionTypePrefecture
	case SubdivisionJP22:
		return SubdivisionTypePrefecture
	case SubdivisionJP23:
		return SubdivisionTypePrefecture
	case SubdivisionJP24:
		return SubdivisionTypePrefecture
	case SubdivisionJP25:
		return SubdivisionTypePrefecture
	case SubdivisionJP26:
		return SubdivisionTypePrefecture
	case SubdivisionJP27:
		return SubdivisionTypePrefecture
	case SubdivisionJP28:
		return SubdivisionTypePrefecture
	case SubdivisionJP29:
		return SubdivisionTypePrefecture
	case SubdivisionJP30:
		return SubdivisionTypePrefecture
	case SubdivisionJP31:
		return SubdivisionTypePrefecture
	case SubdivisionJP32:
		return SubdivisionTypePrefecture
	case SubdivisionJP33:
		return SubdivisionTypePrefecture
	case SubdivisionJP34:
		return SubdivisionTypePrefecture
	case SubdivisionJP35:
		return SubdivisionTypePrefecture
	case SubdivisionJP36:
		return SubdivisionTypePrefecture
	case SubdivisionJP37:
		return SubdivisionTypePrefecture
	case SubdivisionJP38:
		return SubdivisionTypePrefecture
	case SubdivisionJP39:
		return SubdivisionTypePrefecture
	case SubdivisionJP40:
		return SubdivisionTypePrefecture
	case SubdivisionJP41:
		return SubdivisionTypePrefecture
	case SubdivisionJP42:
		return SubdivisionTypePrefecture
	case SubdivisionJP43:
		return SubdivisionTypePrefecture
	case SubdivisionJP44:
		return SubdivisionTypePrefecture
	case SubdivisionJP45:
		return SubdivisionTypePrefecture
	case SubdivisionJP46:
		return SubdivisionTypePrefecture
	case SubdivisionJP47:
		return SubdivisionTypePrefecture
	case SubdivisionKE01:
		return SubdivisionTypeCounty
	case SubdivisionKE02:
		return SubdivisionTypeCounty
	case SubdivisionKE03:
		return SubdivisionTypeCounty
	case SubdivisionKE04:
		return SubdivisionTypeCounty
	case SubdivisionKE05:
		return SubdivisionTypeCounty
	case SubdivisionKE06:
		return SubdivisionTypeCounty
	case SubdivisionKE07:
		return SubdivisionTypeCounty
	case SubdivisionKE08:
		return SubdivisionTypeCounty
	case SubdivisionKE09:
		return SubdivisionTypeCounty
	case SubdivisionKE10:
		return SubdivisionTypeCounty
	case SubdivisionKE11:
		return SubdivisionTypeCounty
	case SubdivisionKE12:
		return SubdivisionTypeCounty
	case SubdivisionKE13:
		return SubdivisionTypeCounty
	case SubdivisionKE14:
		return SubdivisionTypeCounty
	case SubdivisionKE15:
		return SubdivisionTypeCounty
	case SubdivisionKE16:
		return SubdivisionTypeCounty
	case SubdivisionKE17:
		return SubdivisionTypeCounty
	case SubdivisionKE18:
		return SubdivisionTypeCounty
	case SubdivisionKE19:
		return SubdivisionTypeCounty
	case SubdivisionKE20:
		return SubdivisionTypeCounty
	case SubdivisionKE21:
		return SubdivisionTypeCounty
	case SubdivisionKE22:
		return SubdivisionTypeCounty
	case SubdivisionKE23:
		return SubdivisionTypeCounty
	case SubdivisionKE24:
		return SubdivisionTypeCounty
	case SubdivisionKE25:
		return SubdivisionTypeCounty
	case SubdivisionKE26:
		return SubdivisionTypeCounty
	case SubdivisionKE27:
		return SubdivisionTypeCounty
	case SubdivisionKE28:
		return SubdivisionTypeCounty
	case SubdivisionKE29:
		return SubdivisionTypeCounty
	case SubdivisionKE30:
		return SubdivisionTypeCounty
	case SubdivisionKE31:
		return SubdivisionTypeCounty
	case SubdivisionKE32:
		return SubdivisionTypeCounty
	case SubdivisionKE33:
		return SubdivisionTypeCounty
	case SubdivisionKE34:
		return SubdivisionTypeCounty
	case SubdivisionKE35:
		return SubdivisionTypeCounty
	case SubdivisionKE36:
		return SubdivisionTypeCounty
	case SubdivisionKE37:
		return SubdivisionTypeCounty
	case SubdivisionKE38:
		return SubdivisionTypeCounty
	case SubdivisionKE39:
		return SubdivisionTypeCounty
	case SubdivisionKE40:
		return SubdivisionTypeCounty
	case SubdivisionKE41:
		return SubdivisionTypeCounty
	case SubdivisionKE42:
		return SubdivisionTypeCounty
	case SubdivisionKE43:
		return SubdivisionTypeCounty
	case SubdivisionKE44:
		return SubdivisionTypeCounty
	case SubdivisionKE45:
		return SubdivisionTypeCounty
	case SubdivisionKE46:
		return SubdivisionTypeCounty
	case SubdivisionKE47:
		return SubdivisionTypeCounty
	case SubdivisionKGB:
		return SubdivisionTypeRegion
	case SubdivisionKGC:
		return SubdivisionTypeRegion
	case SubdivisionKGGB:
		return SubdivisionTypeCity
	case SubdivisionKGJ:
		return SubdivisionTypeRegion
	case SubdivisionKGN:
		return SubdivisionTypeRegion
	case SubdivisionKGO:
		return SubdivisionTypeRegion
	case SubdivisionKGT:
		return SubdivisionTypeRegion
	case SubdivisionKGY:
		return SubdivisionTypeRegion
	case SubdivisionKH1:
		return SubdivisionTypeProvince
	case SubdivisionKH10:
		return SubdivisionTypeProvince
	case SubdivisionKH11:
		return SubdivisionTypeProvince
	case SubdivisionKH12:
		return SubdivisionTypeAutonomousMunicipality
	case SubdivisionKH13:
		return SubdivisionTypeProvince
	case SubdivisionKH14:
		return SubdivisionTypeProvince
	case SubdivisionKH15:
		return SubdivisionTypeProvince
	case SubdivisionKH16:
		return SubdivisionTypeProvince
	case SubdivisionKH17:
		return SubdivisionTypeProvince
	case SubdivisionKH18:
		return SubdivisionTypeAutonomousMunicipality
	case SubdivisionKH19:
		return SubdivisionTypeProvince
	case SubdivisionKH2:
		return SubdivisionTypeProvince
	case SubdivisionKH20:
		return SubdivisionTypeProvince
	case SubdivisionKH21:
		return SubdivisionTypeProvince
	case SubdivisionKH22:
		return SubdivisionTypeProvince
	case SubdivisionKH23:
		return SubdivisionTypeAutonomousMunicipality
	case SubdivisionKH24:
		return SubdivisionTypeAutonomousMunicipality
	case SubdivisionKH3:
		return SubdivisionTypeProvince
	case SubdivisionKH4:
		return SubdivisionTypeProvince
	case SubdivisionKH5:
		return SubdivisionTypeProvince
	case SubdivisionKH6:
		return SubdivisionTypeProvince
	case SubdivisionKH7:
		return SubdivisionTypeProvince
	case SubdivisionKH8:
		return SubdivisionTypeProvince
	case SubdivisionKH9:
		return SubdivisionTypeProvince
	case SubdivisionKIG:
		return SubdivisionTypeIslandGroup
	case SubdivisionKIL:
		return SubdivisionTypeIslandGroup
	case SubdivisionKIP:
		return SubdivisionTypeIslandGroup
	case SubdivisionKMA:
		return SubdivisionTypeIsland
	case SubdivisionKMG:
		return SubdivisionTypeIsland
	case SubdivisionKMM:
		return SubdivisionTypeIsland
	case SubdivisionKN01:
		return SubdivisionTypeParish
	case SubdivisionKN02:
		return SubdivisionTypeParish
	case SubdivisionKN03:
		return SubdivisionTypeParish
	case SubdivisionKN04:
		return SubdivisionTypeParish
	case SubdivisionKN05:
		return SubdivisionTypeParish
	case SubdivisionKN06:
		return SubdivisionTypeParish
	case SubdivisionKN07:
		return SubdivisionTypeParish
	case SubdivisionKN08:
		return SubdivisionTypeParish
	case SubdivisionKN09:
		return SubdivisionTypeParish
	case SubdivisionKN10:
		return SubdivisionTypeParish
	case SubdivisionKN11:
		return SubdivisionTypeParish
	case SubdivisionKN12:
		return SubdivisionTypeParish
	case SubdivisionKN13:
		return SubdivisionTypeParish
	case SubdivisionKN15:
		return SubdivisionTypeParish
	case SubdivisionKNK:
		return SubdivisionTypeState
	case SubdivisionKNN:
		return SubdivisionTypeState
	case SubdivisionKP01:
		return SubdivisionTypeCapitalCity
	case SubdivisionKP02:
		return SubdivisionTypeProvince
	case SubdivisionKP03:
		return SubdivisionTypeProvince
	case SubdivisionKP04:
		return SubdivisionTypeProvince
	case SubdivisionKP05:
		return SubdivisionTypeProvince
	case SubdivisionKP06:
		return SubdivisionTypeProvince
	case SubdivisionKP07:
		return SubdivisionTypeProvince
	case SubdivisionKP08:
		return SubdivisionTypeProvince
	case SubdivisionKP09:
		return SubdivisionTypeProvince
	case SubdivisionKP10:
		return SubdivisionTypeProvince
	case SubdivisionKP13:
		return SubdivisionTypeSpecialCity
	case SubdivisionKR11:
		return SubdivisionTypeCapitalMetropolitanCity
	case SubdivisionKR26:
		return SubdivisionTypeMetropolitanCities
	case SubdivisionKR27:
		return SubdivisionTypeMetropolitanCities
	case SubdivisionKR28:
		return SubdivisionTypeMetropolitanCities
	case SubdivisionKR29:
		return SubdivisionTypeMetropolitanCities
	case SubdivisionKR30:
		return SubdivisionTypeMetropolitanCities
	case SubdivisionKR31:
		return SubdivisionTypeMetropolitanCities
	case SubdivisionKR41:
		return SubdivisionTypeProvince
	case SubdivisionKR42:
		return SubdivisionTypeProvince
	case SubdivisionKR43:
		return SubdivisionTypeProvince
	case SubdivisionKR44:
		return SubdivisionTypeProvince
	case SubdivisionKR45:
		return SubdivisionTypeProvince
	case SubdivisionKR46:
		return SubdivisionTypeProvince
	case SubdivisionKR47:
		return SubdivisionTypeProvince
	case SubdivisionKR48:
		return SubdivisionTypeProvince
	case SubdivisionKR49:
		return SubdivisionTypeProvince
	case SubdivisionKWAH:
		return SubdivisionTypeGovernorate
	case SubdivisionKWFA:
		return SubdivisionTypeGovernorate
	case SubdivisionKWHA:
		return SubdivisionTypeGovernorate
	case SubdivisionKWJA:
		return SubdivisionTypeGovernorate
	case SubdivisionKWKU:
		return SubdivisionTypeGovernorate
	case SubdivisionKWMU:
		return SubdivisionTypeGovernorate
	case SubdivisionKZAKM:
		return SubdivisionTypeRegion
	case SubdivisionKZAKT:
		return SubdivisionTypeRegion
	case SubdivisionKZALA:
		return SubdivisionTypeCity
	case SubdivisionKZALM:
		return SubdivisionTypeRegion
	case SubdivisionKZAST:
		return SubdivisionTypeCity
	case SubdivisionKZATY:
		return SubdivisionTypeRegion
	case SubdivisionKZKAR:
		return SubdivisionTypeRegion
	case SubdivisionKZKUS:
		return SubdivisionTypeRegion
	case SubdivisionKZKZY:
		return SubdivisionTypeRegion
	case SubdivisionKZMAN:
		return SubdivisionTypeRegion
	case SubdivisionKZPAV:
		return SubdivisionTypeRegion
	case SubdivisionKZSEV:
		return SubdivisionTypeRegion
	case SubdivisionKZVOS:
		return SubdivisionTypeRegion
	case SubdivisionKZYUZ:
		return SubdivisionTypeRegion
	case SubdivisionKZZAP:
		return SubdivisionTypeRegion
	case SubdivisionKZZHA:
		return SubdivisionTypeRegion
	case SubdivisionLAAT:
		return SubdivisionTypeProvince
	case SubdivisionLABK:
		return SubdivisionTypeProvince
	case SubdivisionLABL:
		return SubdivisionTypeProvince
	case SubdivisionLACH:
		return SubdivisionTypeProvince
	case SubdivisionLAHO:
		return SubdivisionTypeProvince
	case SubdivisionLAKH:
		return SubdivisionTypeProvince
	case SubdivisionLALM:
		return SubdivisionTypeProvince
	case SubdivisionLALP:
		return SubdivisionTypeProvince
	case SubdivisionLAOU:
		return SubdivisionTypeProvince
	case SubdivisionLAPH:
		return SubdivisionTypeProvince
	case SubdivisionLASL:
		return SubdivisionTypeProvince
	case SubdivisionLASV:
		return SubdivisionTypeProvince
	case SubdivisionLAVI:
		return SubdivisionTypeProvince
	case SubdivisionLAVT:
		return SubdivisionTypePrefecture
	case SubdivisionLAXA:
		return SubdivisionTypeProvince
	case SubdivisionLAXE:
		return SubdivisionTypeProvince
	case SubdivisionLAXI:
		return SubdivisionTypeProvince
	case SubdivisionLAXS:
		return SubdivisionTypeProvince
	case SubdivisionLBAK:
		return SubdivisionTypeGovernorate
	case SubdivisionLBAS:
		return SubdivisionTypeGovernorate
	case SubdivisionLBBA:
		return SubdivisionTypeGovernorate
	case SubdivisionLBBH:
		return SubdivisionTypeGovernorate
	case SubdivisionLBBI:
		return SubdivisionTypeGovernorate
	case SubdivisionLBJA:
		return SubdivisionTypeGovernorate
	case SubdivisionLBJL:
		return SubdivisionTypeGovernorate
	case SubdivisionLBNA:
		return SubdivisionTypeGovernorate
	case SubdivisionLI01:
		return SubdivisionTypeCommune
	case SubdivisionLI02:
		return SubdivisionTypeCommune
	case SubdivisionLI03:
		return SubdivisionTypeCommune
	case SubdivisionLI04:
		return SubdivisionTypeCommune
	case SubdivisionLI05:
		return SubdivisionTypeCommune
	case SubdivisionLI06:
		return SubdivisionTypeCommune
	case SubdivisionLI07:
		return SubdivisionTypeCommune
	case SubdivisionLI08:
		return SubdivisionTypeCommune
	case SubdivisionLI09:
		return SubdivisionTypeCommune
	case SubdivisionLI10:
		return SubdivisionTypeCommune
	case SubdivisionLI11:
		return SubdivisionTypeCommune
	case SubdivisionLK1:
		return SubdivisionTypeProvince
	case SubdivisionLK11:
		return SubdivisionTypeDistrict
	case SubdivisionLK12:
		return SubdivisionTypeDistrict
	case SubdivisionLK13:
		return SubdivisionTypeDistrict
	case SubdivisionLK2:
		return SubdivisionTypeProvince
	case SubdivisionLK21:
		return SubdivisionTypeDistrict
	case SubdivisionLK22:
		return SubdivisionTypeDistrict
	case SubdivisionLK23:
		return SubdivisionTypeDistrict
	case SubdivisionLK3:
		return SubdivisionTypeProvince
	case SubdivisionLK31:
		return SubdivisionTypeDistrict
	case SubdivisionLK32:
		return SubdivisionTypeDistrict
	case SubdivisionLK33:
		return SubdivisionTypeDistrict
	case SubdivisionLK4:
		return SubdivisionTypeProvince
	case SubdivisionLK41:
		return SubdivisionTypeDistrict
	case SubdivisionLK42:
		return SubdivisionTypeDistrict
	case SubdivisionLK43:
		return SubdivisionTypeDistrict
	case SubdivisionLK44:
		return SubdivisionTypeDistrict
	case SubdivisionLK45:
		return SubdivisionTypeDistrict
	case SubdivisionLK5:
		return SubdivisionTypeProvince
	case SubdivisionLK51:
		return SubdivisionTypeDistrict
	case SubdivisionLK52:
		return SubdivisionTypeDistrict
	case SubdivisionLK53:
		return SubdivisionTypeDistrict
	case SubdivisionLK6:
		return SubdivisionTypeProvince
	case SubdivisionLK61:
		return SubdivisionTypeDistrict
	case SubdivisionLK62:
		return SubdivisionTypeDistrict
	case SubdivisionLK7:
		return SubdivisionTypeProvince
	case SubdivisionLK71:
		return SubdivisionTypeDistrict
	case SubdivisionLK72:
		return SubdivisionTypeDistrict
	case SubdivisionLK8:
		return SubdivisionTypeProvince
	case SubdivisionLK81:
		return SubdivisionTypeDistrict
	case SubdivisionLK82:
		return SubdivisionTypeDistrict
	case SubdivisionLK9:
		return SubdivisionTypeProvince
	case SubdivisionLK91:
		return SubdivisionTypeDistrict
	case SubdivisionLK92:
		return SubdivisionTypeDistrict
	case SubdivisionLRBG:
		return SubdivisionTypeCounty
	case SubdivisionLRBM:
		return SubdivisionTypeCounty
	case SubdivisionLRCM:
		return SubdivisionTypeCounty
	case SubdivisionLRGB:
		return SubdivisionTypeCounty
	case SubdivisionLRGG:
		return SubdivisionTypeCounty
	case SubdivisionLRGK:
		return SubdivisionTypeCounty
	case SubdivisionLRLO:
		return SubdivisionTypeCounty
	case SubdivisionLRMG:
		return SubdivisionTypeCounty
	case SubdivisionLRMO:
		return SubdivisionTypeCounty
	case SubdivisionLRMY:
		return SubdivisionTypeCounty
	case SubdivisionLRNI:
		return SubdivisionTypeCounty
	case SubdivisionLRRI:
		return SubdivisionTypeCounty
	case SubdivisionLRSI:
		return SubdivisionTypeCounty
	case SubdivisionLSA:
		return SubdivisionTypeDistrict
	case SubdivisionLSB:
		return SubdivisionTypeDistrict
	case SubdivisionLSC:
		return SubdivisionTypeDistrict
	case SubdivisionLSD:
		return SubdivisionTypeDistrict
	case SubdivisionLSE:
		return SubdivisionTypeDistrict
	case SubdivisionLSF:
		return SubdivisionTypeDistrict
	case SubdivisionLSG:
		return SubdivisionTypeDistrict
	case SubdivisionLSH:
		return SubdivisionTypeDistrict
	case SubdivisionLSJ:
		return SubdivisionTypeDistrict
	case SubdivisionLSK:
		return SubdivisionTypeDistrict
	case SubdivisionLTAL:
		return SubdivisionTypeCounty
	case SubdivisionLTKL:
		return SubdivisionTypeCounty
	case SubdivisionLTKU:
		return SubdivisionTypeCounty
	case SubdivisionLTMR:
		return SubdivisionTypeCounty
	case SubdivisionLTPN:
		return SubdivisionTypeCounty
	case SubdivisionLTSA:
		return SubdivisionTypeCounty
	case SubdivisionLTTA:
		return SubdivisionTypeCounty
	case SubdivisionLTTE:
		return SubdivisionTypeCounty
	case SubdivisionLTUT:
		return SubdivisionTypeCounty
	case SubdivisionLTVL:
		return SubdivisionTypeCounty
	case SubdivisionLUD:
		return SubdivisionTypeDistrict
	case SubdivisionLUG:
		return SubdivisionTypeDistrict
	case SubdivisionLUL:
		return SubdivisionTypeDistrict
	case SubdivisionLV001:
		return SubdivisionTypeMunicipality
	case SubdivisionLV002:
		return SubdivisionTypeMunicipality
	case SubdivisionLV003:
		return SubdivisionTypeMunicipality
	case SubdivisionLV004:
		return SubdivisionTypeMunicipality
	case SubdivisionLV005:
		return SubdivisionTypeMunicipality
	case SubdivisionLV006:
		return SubdivisionTypeMunicipality
	case SubdivisionLV007:
		return SubdivisionTypeMunicipality
	case SubdivisionLV008:
		return SubdivisionTypeMunicipality
	case SubdivisionLV009:
		return SubdivisionTypeMunicipality
	case SubdivisionLV010:
		return SubdivisionTypeMunicipality
	case SubdivisionLV011:
		return SubdivisionTypeMunicipality
	case SubdivisionLV012:
		return SubdivisionTypeMunicipality
	case SubdivisionLV013:
		return SubdivisionTypeMunicipality
	case SubdivisionLV014:
		return SubdivisionTypeMunicipality
	case SubdivisionLV015:
		return SubdivisionTypeMunicipality
	case SubdivisionLV016:
		return SubdivisionTypeMunicipality
	case SubdivisionLV017:
		return SubdivisionTypeMunicipality
	case SubdivisionLV018:
		return SubdivisionTypeMunicipality
	case SubdivisionLV019:
		return SubdivisionTypeMunicipality
	case SubdivisionLV020:
		return SubdivisionTypeMunicipality
	case SubdivisionLV021:
		return SubdivisionTypeMunicipality
	case SubdivisionLV022:
		return SubdivisionTypeMunicipality
	case SubdivisionLV023:
		return SubdivisionTypeMunicipality
	case SubdivisionLV024:
		return SubdivisionTypeMunicipality
	case SubdivisionLV025:
		return SubdivisionTypeMunicipality
	case SubdivisionLV026:
		return SubdivisionTypeMunicipality
	case SubdivisionLV027:
		return SubdivisionTypeMunicipality
	case SubdivisionLV028:
		return SubdivisionTypeMunicipality
	case SubdivisionLV029:
		return SubdivisionTypeMunicipality
	case SubdivisionLV030:
		return SubdivisionTypeMunicipality
	case SubdivisionLV031:
		return SubdivisionTypeMunicipality
	case SubdivisionLV032:
		return SubdivisionTypeMunicipality
	case SubdivisionLV033:
		return SubdivisionTypeMunicipality
	case SubdivisionLV034:
		return SubdivisionTypeMunicipality
	case SubdivisionLV035:
		return SubdivisionTypeMunicipality
	case SubdivisionLV036:
		return SubdivisionTypeMunicipality
	case SubdivisionLV037:
		return SubdivisionTypeMunicipality
	case SubdivisionLV038:
		return SubdivisionTypeMunicipality
	case SubdivisionLV039:
		return SubdivisionTypeMunicipality
	case SubdivisionLV040:
		return SubdivisionTypeMunicipality
	case SubdivisionLV041:
		return SubdivisionTypeMunicipality
	case SubdivisionLV042:
		return SubdivisionTypeMunicipality
	case SubdivisionLV043:
		return SubdivisionTypeMunicipality
	case SubdivisionLV044:
		return SubdivisionTypeMunicipality
	case SubdivisionLV045:
		return SubdivisionTypeMunicipality
	case SubdivisionLV046:
		return SubdivisionTypeMunicipality
	case SubdivisionLV047:
		return SubdivisionTypeMunicipality
	case SubdivisionLV048:
		return SubdivisionTypeMunicipality
	case SubdivisionLV049:
		return SubdivisionTypeMunicipality
	case SubdivisionLV050:
		return SubdivisionTypeMunicipality
	case SubdivisionLV051:
		return SubdivisionTypeMunicipality
	case SubdivisionLV052:
		return SubdivisionTypeMunicipality
	case SubdivisionLV053:
		return SubdivisionTypeMunicipality
	case SubdivisionLV054:
		return SubdivisionTypeMunicipality
	case SubdivisionLV055:
		return SubdivisionTypeMunicipality
	case SubdivisionLV056:
		return SubdivisionTypeMunicipality
	case SubdivisionLV057:
		return SubdivisionTypeMunicipality
	case SubdivisionLV058:
		return SubdivisionTypeMunicipality
	case SubdivisionLV059:
		return SubdivisionTypeMunicipality
	case SubdivisionLV060:
		return SubdivisionTypeMunicipality
	case SubdivisionLV061:
		return SubdivisionTypeMunicipality
	case SubdivisionLV062:
		return SubdivisionTypeMunicipality
	case SubdivisionLV063:
		return SubdivisionTypeMunicipality
	case SubdivisionLV064:
		return SubdivisionTypeMunicipality
	case SubdivisionLV065:
		return SubdivisionTypeMunicipality
	case SubdivisionLV066:
		return SubdivisionTypeMunicipality
	case SubdivisionLV067:
		return SubdivisionTypeMunicipality
	case SubdivisionLV068:
		return SubdivisionTypeMunicipality
	case SubdivisionLV069:
		return SubdivisionTypeMunicipality
	case SubdivisionLV070:
		return SubdivisionTypeMunicipality
	case SubdivisionLV071:
		return SubdivisionTypeMunicipality
	case SubdivisionLV072:
		return SubdivisionTypeMunicipality
	case SubdivisionLV073:
		return SubdivisionTypeMunicipality
	case SubdivisionLV074:
		return SubdivisionTypeMunicipality
	case SubdivisionLV075:
		return SubdivisionTypeMunicipality
	case SubdivisionLV076:
		return SubdivisionTypeMunicipality
	case SubdivisionLV077:
		return SubdivisionTypeMunicipality
	case SubdivisionLV078:
		return SubdivisionTypeMunicipality
	case SubdivisionLV079:
		return SubdivisionTypeMunicipality
	case SubdivisionLV080:
		return SubdivisionTypeMunicipality
	case SubdivisionLV081:
		return SubdivisionTypeMunicipality
	case SubdivisionLV082:
		return SubdivisionTypeMunicipality
	case SubdivisionLV083:
		return SubdivisionTypeMunicipality
	case SubdivisionLV084:
		return SubdivisionTypeMunicipality
	case SubdivisionLV085:
		return SubdivisionTypeMunicipality
	case SubdivisionLV086:
		return SubdivisionTypeMunicipality
	case SubdivisionLV087:
		return SubdivisionTypeMunicipality
	case SubdivisionLV088:
		return SubdivisionTypeMunicipality
	case SubdivisionLV089:
		return SubdivisionTypeMunicipality
	case SubdivisionLV090:
		return SubdivisionTypeMunicipality
	case SubdivisionLV091:
		return SubdivisionTypeMunicipality
	case SubdivisionLV092:
		return SubdivisionTypeMunicipality
	case SubdivisionLV093:
		return SubdivisionTypeMunicipality
	case SubdivisionLV094:
		return SubdivisionTypeMunicipality
	case SubdivisionLV095:
		return SubdivisionTypeMunicipality
	case SubdivisionLV096:
		return SubdivisionTypeMunicipality
	case SubdivisionLV097:
		return SubdivisionTypeMunicipality
	case SubdivisionLV098:
		return SubdivisionTypeMunicipality
	case SubdivisionLV099:
		return SubdivisionTypeMunicipality
	case SubdivisionLV100:
		return SubdivisionTypeMunicipality
	case SubdivisionLV101:
		return SubdivisionTypeMunicipality
	case SubdivisionLV102:
		return SubdivisionTypeMunicipality
	case SubdivisionLV103:
		return SubdivisionTypeMunicipality
	case SubdivisionLV104:
		return SubdivisionTypeMunicipality
	case SubdivisionLV105:
		return SubdivisionTypeMunicipality
	case SubdivisionLV106:
		return SubdivisionTypeMunicipality
	case SubdivisionLV107:
		return SubdivisionTypeMunicipality
	case SubdivisionLV108:
		return SubdivisionTypeMunicipality
	case SubdivisionLV109:
		return SubdivisionTypeMunicipality
	case SubdivisionLV110:
		return SubdivisionTypeMunicipality
	case SubdivisionLVDGV:
		return SubdivisionTypeRepublicanCity
	case SubdivisionLVJEL:
		return SubdivisionTypeRepublicanCity
	case SubdivisionLVJKB:
		return SubdivisionTypeRepublicanCity
	case SubdivisionLVJUR:
		return SubdivisionTypeRepublicanCity
	case SubdivisionLVLPX:
		return SubdivisionTypeRepublicanCity
	case SubdivisionLVREZ:
		return SubdivisionTypeRepublicanCity
	case SubdivisionLVRIX:
		return SubdivisionTypeRepublicanCity
	case SubdivisionLVVEN:
		return SubdivisionTypeRepublicanCity
	case SubdivisionLVVMR:
		return SubdivisionTypeRepublicanCity
	case SubdivisionLYBA:
		return SubdivisionTypePopularates
	case SubdivisionLYBU:
		return SubdivisionTypePopularates
	case SubdivisionLYDR:
		return SubdivisionTypePopularates
	case SubdivisionLYGT:
		return SubdivisionTypePopularates
	case SubdivisionLYJA:
		return SubdivisionTypePopularates
	case SubdivisionLYJB:
		return SubdivisionTypePopularates
	case SubdivisionLYJG:
		return SubdivisionTypePopularates
	case SubdivisionLYJI:
		return SubdivisionTypePopularates
	case SubdivisionLYJU:
		return SubdivisionTypePopularates
	case SubdivisionLYKF:
		return SubdivisionTypePopularates
	case SubdivisionLYMB:
		return SubdivisionTypePopularates
	case SubdivisionLYMI:
		return SubdivisionTypePopularates
	case SubdivisionLYMJ:
		return SubdivisionTypePopularates
	case SubdivisionLYMQ:
		return SubdivisionTypePopularates
	case SubdivisionLYNL:
		return SubdivisionTypePopularates
	case SubdivisionLYNQ:
		return SubdivisionTypePopularates
	case SubdivisionLYSB:
		return SubdivisionTypePopularates
	case SubdivisionLYSR:
		return SubdivisionTypePopularates
	case SubdivisionLYTB:
		return SubdivisionTypePopularates
	case SubdivisionLYWA:
		return SubdivisionTypePopularates
	case SubdivisionLYWD:
		return SubdivisionTypePopularates
	case SubdivisionLYWS:
		return SubdivisionTypePopularates
	case SubdivisionLYZA:
		return SubdivisionTypePopularates
	case SubdivisionMA01:
		return SubdivisionTypeRegion
	case SubdivisionMA02:
		return SubdivisionTypeRegion
	case SubdivisionMA03:
		return SubdivisionTypeRegion
	case SubdivisionMA04:
		return SubdivisionTypeRegion
	case SubdivisionMA05:
		return SubdivisionTypeRegion
	case SubdivisionMA06:
		return SubdivisionTypeRegion
	case SubdivisionMA07:
		return SubdivisionTypeRegion
	case SubdivisionMA08:
		return SubdivisionTypeRegion
	case SubdivisionMA09:
		return SubdivisionTypeRegion
	case SubdivisionMA10:
		return SubdivisionTypeRegion
	case SubdivisionMA11:
		return SubdivisionTypeRegion
	case SubdivisionMA12:
		return SubdivisionTypeRegion
	case SubdivisionMAAGD:
		return SubdivisionTypePrefecture
	case SubdivisionMAAOU:
		return SubdivisionTypeProvince
	case SubdivisionMAASZ:
		return SubdivisionTypeProvince
	case SubdivisionMAAZI:
		return SubdivisionTypeProvince
	case SubdivisionMABEM:
		return SubdivisionTypeProvince
	case SubdivisionMABER:
		return SubdivisionTypeProvince
	case SubdivisionMABES:
		return SubdivisionTypeProvince
	case SubdivisionMABOD:
		return SubdivisionTypeProvince
	case SubdivisionMABOM:
		return SubdivisionTypeProvince
	case SubdivisionMABRR:
		return SubdivisionTypeProvince
	case SubdivisionMACAS:
		return SubdivisionTypePrefecture
	case SubdivisionMACHE:
		return SubdivisionTypeProvince
	case SubdivisionMACHI:
		return SubdivisionTypeProvince
	case SubdivisionMACHT:
		return SubdivisionTypeProvince
	case SubdivisionMADRI:
		return SubdivisionTypeProvince
	case SubdivisionMAERR:
		return SubdivisionTypeProvince
	case SubdivisionMAESI:
		return SubdivisionTypeProvince
	case SubdivisionMAESM:
		return SubdivisionTypeProvince
	case SubdivisionMAFAH:
		return SubdivisionTypeProvince
	case SubdivisionMAFES:
		return SubdivisionTypePrefecture
	case SubdivisionMAFIG:
		return SubdivisionTypeProvince
	case SubdivisionMAFQH:
		return SubdivisionTypeProvince
	case SubdivisionMAGUE:
		return SubdivisionTypeProvince
	case SubdivisionMAGUF:
		return SubdivisionTypeProvince
	case SubdivisionMAHAJ:
		return SubdivisionTypeProvince
	case SubdivisionMAHAO:
		return SubdivisionTypeProvince
	case SubdivisionMAHOC:
		return SubdivisionTypeProvince
	case SubdivisionMAIFR:
		return SubdivisionTypeProvince
	case SubdivisionMAINE:
		return SubdivisionTypePrefecture
	case SubdivisionMAJDI:
		return SubdivisionTypeProvince
	case SubdivisionMAJRA:
		return SubdivisionTypeProvince
	case SubdivisionMAKEN:
		return SubdivisionTypeProvince
	case SubdivisionMAKES:
		return SubdivisionTypeProvince
	case SubdivisionMAKHE:
		return SubdivisionTypeProvince
	case SubdivisionMAKHN:
		return SubdivisionTypeProvince
	case SubdivisionMAKHO:
		return SubdivisionTypeProvince
	case SubdivisionMALAA:
		return SubdivisionTypeProvince
	case SubdivisionMALAR:
		return SubdivisionTypeProvince
	case SubdivisionMAMAR:
		return SubdivisionTypePrefecture
	case SubdivisionMAMDF:
		return SubdivisionTypePrefecture
	case SubdivisionMAMED:
		return SubdivisionTypeProvince
	case SubdivisionMAMEK:
		return SubdivisionTypePrefecture
	case SubdivisionMAMID:
		return SubdivisionTypeProvince
	case SubdivisionMAMOH:
		return SubdivisionTypePrefecture
	case SubdivisionMAMOU:
		return SubdivisionTypeProvince
	case SubdivisionMANAD:
		return SubdivisionTypeProvince
	case SubdivisionMANOU:
		return SubdivisionTypeProvince
	case SubdivisionMAOUA:
		return SubdivisionTypeProvince
	case SubdivisionMAOUD:
		return SubdivisionTypeProvince
	case SubdivisionMAOUJ:
		return SubdivisionTypePrefecture
	case SubdivisionMAOUZ:
		return SubdivisionTypeProvince
	case SubdivisionMARAB:
		return SubdivisionTypePrefecture
	case SubdivisionMAREH:
		return SubdivisionTypeProvince
	case SubdivisionMASAF:
		return SubdivisionTypeProvince
	case SubdivisionMASAL:
		return SubdivisionTypePrefecture
	case SubdivisionMASEF:
		return SubdivisionTypeProvince
	case SubdivisionMASET:
		return SubdivisionTypeProvince
	case SubdivisionMASIB:
		return SubdivisionTypeProvince
	case SubdivisionMASIF:
		return SubdivisionTypeProvince
	case SubdivisionMASIK:
		return SubdivisionTypeProvince
	case SubdivisionMASIL:
		return SubdivisionTypeProvince
	case SubdivisionMASKH:
		return SubdivisionTypePrefecture
	case SubdivisionMATAF:
		return SubdivisionTypeProvince
	case SubdivisionMATAI:
		return SubdivisionTypeProvince
	case SubdivisionMATAO:
		return SubdivisionTypeProvince
	case SubdivisionMATAR:
		return SubdivisionTypeProvince
	case SubdivisionMATAT:
		return SubdivisionTypeProvince
	case SubdivisionMATAZ:
		return SubdivisionTypeProvince
	case SubdivisionMATET:
		return SubdivisionTypeProvince
	case SubdivisionMATIN:
		return SubdivisionTypeProvince
	case SubdivisionMATIZ:
		return SubdivisionTypeProvince
	case SubdivisionMATNG:
		return SubdivisionTypePrefecture
	case SubdivisionMATNT:
		return SubdivisionTypeProvince
	case SubdivisionMAYUS:
		return SubdivisionTypeProvince
	case SubdivisionMAZAG:
		return SubdivisionTypeProvince
	case SubdivisionMCCL:
		return SubdivisionTypeQuarter
	case SubdivisionMCCO:
		return SubdivisionTypeQuarter
	case SubdivisionMCFO:
		return SubdivisionTypeQuarter
	case SubdivisionMCGA:
		return SubdivisionTypeQuarter
	case SubdivisionMCJE:
		return SubdivisionTypeQuarter
	case SubdivisionMCLA:
		return SubdivisionTypeQuarter
	case SubdivisionMCMA:
		return SubdivisionTypeQuarter
	case SubdivisionMCMC:
		return SubdivisionTypeQuarter
	case SubdivisionMCMG:
		return SubdivisionTypeQuarter
	case SubdivisionMCMO:
		return SubdivisionTypeQuarter
	case SubdivisionMCMU:
		return SubdivisionTypeQuarter
	case SubdivisionMCPH:
		return SubdivisionTypeQuarter
	case SubdivisionMCSD:
		return SubdivisionTypeQuarter
	case SubdivisionMCSO:
		return SubdivisionTypeQuarter
	case SubdivisionMCSP:
		return SubdivisionTypeQuarter
	case SubdivisionMCSR:
		return SubdivisionTypeQuarter
	case SubdivisionMCVR:
		return SubdivisionTypeQuarter
	case SubdivisionMDAN:
		return SubdivisionTypeDistrict
	case SubdivisionMDBA:
		return SubdivisionTypeCity
	case SubdivisionMDBD:
		return SubdivisionTypeCity
	case SubdivisionMDBR:
		return SubdivisionTypeDistrict
	case SubdivisionMDBS:
		return SubdivisionTypeDistrict
	case SubdivisionMDCA:
		return SubdivisionTypeDistrict
	case SubdivisionMDCL:
		return SubdivisionTypeDistrict
	case SubdivisionMDCM:
		return SubdivisionTypeDistrict
	case SubdivisionMDCR:
		return SubdivisionTypeDistrict
	case SubdivisionMDCS:
		return SubdivisionTypeDistrict
	case SubdivisionMDCT:
		return SubdivisionTypeDistrict
	case SubdivisionMDCU:
		return SubdivisionTypeCity
	case SubdivisionMDDO:
		return SubdivisionTypeDistrict
	case SubdivisionMDDR:
		return SubdivisionTypeDistrict
	case SubdivisionMDDU:
		return SubdivisionTypeDistrict
	case SubdivisionMDED:
		return SubdivisionTypeDistrict
	case SubdivisionMDFA:
		return SubdivisionTypeDistrict
	case SubdivisionMDFL:
		return SubdivisionTypeDistrict
	case SubdivisionMDGA:
		return SubdivisionTypeAutonomousTerritorialUnit
	case SubdivisionMDGL:
		return SubdivisionTypeDistrict
	case SubdivisionMDHI:
		return SubdivisionTypeDistrict
	case SubdivisionMDIA:
		return SubdivisionTypeDistrict
	case SubdivisionMDLE:
		return SubdivisionTypeDistrict
	case SubdivisionMDNI:
		return SubdivisionTypeDistrict
	case SubdivisionMDOC:
		return SubdivisionTypeDistrict
	case SubdivisionMDOR:
		return SubdivisionTypeDistrict
	case SubdivisionMDRE:
		return SubdivisionTypeDistrict
	case SubdivisionMDRI:
		return SubdivisionTypeDistrict
	case SubdivisionMDSD:
		return SubdivisionTypeDistrict
	case SubdivisionMDSI:
		return SubdivisionTypeDistrict
	case SubdivisionMDSN:
		return SubdivisionTypeTerritorialUnit
	case SubdivisionMDSO:
		return SubdivisionTypeDistrict
	case SubdivisionMDST:
		return SubdivisionTypeDistrict
	case SubdivisionMDSV:
		return SubdivisionTypeDistrict
	case SubdivisionMDTA:
		return SubdivisionTypeDistrict
	case SubdivisionMDTE:
		return SubdivisionTypeDistrict
	case SubdivisionMDUN:
		return SubdivisionTypeDistrict
	case SubdivisionME01:
		return SubdivisionTypeMunicipality
	case SubdivisionME02:
		return SubdivisionTypeMunicipality
	case SubdivisionME03:
		return SubdivisionTypeMunicipality
	case SubdivisionME04:
		return SubdivisionTypeMunicipality
	case SubdivisionME05:
		return SubdivisionTypeMunicipality
	case SubdivisionME06:
		return SubdivisionTypeMunicipality
	case SubdivisionME07:
		return SubdivisionTypeMunicipality
	case SubdivisionME08:
		return SubdivisionTypeMunicipality
	case SubdivisionME09:
		return SubdivisionTypeMunicipality
	case SubdivisionME10:
		return SubdivisionTypeMunicipality
	case SubdivisionME11:
		return SubdivisionTypeMunicipality
	case SubdivisionME12:
		return SubdivisionTypeMunicipality
	case SubdivisionME13:
		return SubdivisionTypeMunicipality
	case SubdivisionME14:
		return SubdivisionTypeMunicipality
	case SubdivisionME15:
		return SubdivisionTypeMunicipality
	case SubdivisionME16:
		return SubdivisionTypeMunicipality
	case SubdivisionME17:
		return SubdivisionTypeMunicipality
	case SubdivisionME18:
		return SubdivisionTypeMunicipality
	case SubdivisionME19:
		return SubdivisionTypeMunicipality
	case SubdivisionME20:
		return SubdivisionTypeMunicipality
	case SubdivisionME21:
		return SubdivisionTypeMunicipality
	case SubdivisionMGA:
		return SubdivisionTypeAutonomousProvince
	case SubdivisionMGD:
		return SubdivisionTypeAutonomousProvince
	case SubdivisionMGF:
		return SubdivisionTypeAutonomousProvince
	case SubdivisionMGM:
		return SubdivisionTypeAutonomousProvince
	case SubdivisionMGT:
		return SubdivisionTypeAutonomousProvince
	case SubdivisionMGU:
		return SubdivisionTypeAutonomousProvince
	case SubdivisionMHALK:
		return SubdivisionTypeMunicipality
	case SubdivisionMHALL:
		return SubdivisionTypeMunicipality
	case SubdivisionMHARN:
		return SubdivisionTypeMunicipality
	case SubdivisionMHAUR:
		return SubdivisionTypeMunicipality
	case SubdivisionMHEBO:
		return SubdivisionTypeMunicipality
	case SubdivisionMHENI:
		return SubdivisionTypeMunicipality
	case SubdivisionMHJAB:
		return SubdivisionTypeMunicipality
	case SubdivisionMHJAL:
		return SubdivisionTypeMunicipality
	case SubdivisionMHKIL:
		return SubdivisionTypeMunicipality
	case SubdivisionMHKWA:
		return SubdivisionTypeMunicipality
	case SubdivisionMHL:
		return SubdivisionTypeChainsOfIslands
	case SubdivisionMHLAE:
		return SubdivisionTypeMunicipality
	case SubdivisionMHLIB:
		return SubdivisionTypeMunicipality
	case SubdivisionMHLIK:
		return SubdivisionTypeMunicipality
	case SubdivisionMHMAJ:
		return SubdivisionTypeMunicipality
	case SubdivisionMHMAL:
		return SubdivisionTypeMunicipality
	case SubdivisionMHMEJ:
		return SubdivisionTypeMunicipality
	case SubdivisionMHMIL:
		return SubdivisionTypeMunicipality
	case SubdivisionMHNMK:
		return SubdivisionTypeMunicipality
	case SubdivisionMHNMU:
		return SubdivisionTypeMunicipality
	case SubdivisionMHRON:
		return SubdivisionTypeMunicipality
	case SubdivisionMHT:
		return SubdivisionTypeChainsOfIslands
	case SubdivisionMHUJA:
		return SubdivisionTypeMunicipality
	case SubdivisionMHUTI:
		return SubdivisionTypeMunicipality
	case SubdivisionMHWTJ:
		return SubdivisionTypeMunicipality
	case SubdivisionMHWTN:
		return SubdivisionTypeMunicipality
	case SubdivisionMK01:
		return SubdivisionTypeMunicipality
	case SubdivisionMK02:
		return SubdivisionTypeMunicipality
	case SubdivisionMK03:
		return SubdivisionTypeMunicipality
	case SubdivisionMK04:
		return SubdivisionTypeMunicipality
	case SubdivisionMK05:
		return SubdivisionTypeMunicipality
	case SubdivisionMK06:
		return SubdivisionTypeMunicipality
	case SubdivisionMK07:
		return SubdivisionTypeMunicipality
	case SubdivisionMK08:
		return SubdivisionTypeMunicipality
	case SubdivisionMK09:
		return SubdivisionTypeMunicipality
	case SubdivisionMK10:
		return SubdivisionTypeMunicipality
	case SubdivisionMK11:
		return SubdivisionTypeMunicipality
	case SubdivisionMK12:
		return SubdivisionTypeMunicipality
	case SubdivisionMK13:
		return SubdivisionTypeMunicipality
	case SubdivisionMK14:
		return SubdivisionTypeMunicipality
	case SubdivisionMK15:
		return SubdivisionTypeMunicipality
	case SubdivisionMK16:
		return SubdivisionTypeMunicipality
	case SubdivisionMK17:
		return SubdivisionTypeMunicipality
	case SubdivisionMK18:
		return SubdivisionTypeMunicipality
	case SubdivisionMK19:
		return SubdivisionTypeMunicipality
	case SubdivisionMK20:
		return SubdivisionTypeMunicipality
	case SubdivisionMK21:
		return SubdivisionTypeMunicipality
	case SubdivisionMK22:
		return SubdivisionTypeMunicipality
	case SubdivisionMK23:
		return SubdivisionTypeMunicipality
	case SubdivisionMK24:
		return SubdivisionTypeMunicipality
	case SubdivisionMK25:
		return SubdivisionTypeMunicipality
	case SubdivisionMK26:
		return SubdivisionTypeMunicipality
	case SubdivisionMK27:
		return SubdivisionTypeMunicipality
	case SubdivisionMK28:
		return SubdivisionTypeMunicipality
	case SubdivisionMK29:
		return SubdivisionTypeMunicipality
	case SubdivisionMK30:
		return SubdivisionTypeMunicipality
	case SubdivisionMK31:
		return SubdivisionTypeMunicipality
	case SubdivisionMK32:
		return SubdivisionTypeMunicipality
	case SubdivisionMK33:
		return SubdivisionTypeMunicipality
	case SubdivisionMK34:
		return SubdivisionTypeMunicipality
	case SubdivisionMK35:
		return SubdivisionTypeMunicipality
	case SubdivisionMK36:
		return SubdivisionTypeMunicipality
	case SubdivisionMK37:
		return SubdivisionTypeMunicipality
	case SubdivisionMK38:
		return SubdivisionTypeMunicipality
	case SubdivisionMK39:
		return SubdivisionTypeMunicipality
	case SubdivisionMK40:
		return SubdivisionTypeMunicipality
	case SubdivisionMK41:
		return SubdivisionTypeMunicipality
	case SubdivisionMK42:
		return SubdivisionTypeMunicipality
	case SubdivisionMK43:
		return SubdivisionTypeMunicipality
	case SubdivisionMK44:
		return SubdivisionTypeMunicipality
	case SubdivisionMK45:
		return SubdivisionTypeMunicipality
	case SubdivisionMK46:
		return SubdivisionTypeMunicipality
	case SubdivisionMK47:
		return SubdivisionTypeMunicipality
	case SubdivisionMK48:
		return SubdivisionTypeMunicipality
	case SubdivisionMK49:
		return SubdivisionTypeMunicipality
	case SubdivisionMK50:
		return SubdivisionTypeMunicipality
	case SubdivisionMK51:
		return SubdivisionTypeMunicipality
	case SubdivisionMK52:
		return SubdivisionTypeMunicipality
	case SubdivisionMK53:
		return SubdivisionTypeMunicipality
	case SubdivisionMK54:
		return SubdivisionTypeMunicipality
	case SubdivisionMK55:
		return SubdivisionTypeMunicipality
	case SubdivisionMK56:
		return SubdivisionTypeMunicipality
	case SubdivisionMK57:
		return SubdivisionTypeMunicipality
	case SubdivisionMK58:
		return SubdivisionTypeMunicipality
	case SubdivisionMK59:
		return SubdivisionTypeMunicipality
	case SubdivisionMK60:
		return SubdivisionTypeMunicipality
	case SubdivisionMK61:
		return SubdivisionTypeMunicipality
	case SubdivisionMK62:
		return SubdivisionTypeMunicipality
	case SubdivisionMK63:
		return SubdivisionTypeMunicipality
	case SubdivisionMK64:
		return SubdivisionTypeMunicipality
	case SubdivisionMK65:
		return SubdivisionTypeMunicipality
	case SubdivisionMK66:
		return SubdivisionTypeMunicipality
	case SubdivisionMK67:
		return SubdivisionTypeMunicipality
	case SubdivisionMK68:
		return SubdivisionTypeMunicipality
	case SubdivisionMK69:
		return SubdivisionTypeMunicipality
	case SubdivisionMK70:
		return SubdivisionTypeMunicipality
	case SubdivisionMK71:
		return SubdivisionTypeMunicipality
	case SubdivisionMK72:
		return SubdivisionTypeMunicipality
	case SubdivisionMK73:
		return SubdivisionTypeMunicipality
	case SubdivisionMK74:
		return SubdivisionTypeMunicipality
	case SubdivisionMK75:
		return SubdivisionTypeMunicipality
	case SubdivisionMK76:
		return SubdivisionTypeMunicipality
	case SubdivisionMK77:
		return SubdivisionTypeMunicipality
	case SubdivisionMK78:
		return SubdivisionTypeMunicipality
	case SubdivisionMK79:
		return SubdivisionTypeMunicipality
	case SubdivisionMK80:
		return SubdivisionTypeMunicipality
	case SubdivisionMK81:
		return SubdivisionTypeMunicipality
	case SubdivisionMK82:
		return SubdivisionTypeMunicipality
	case SubdivisionMK83:
		return SubdivisionTypeMunicipality
	case SubdivisionMK84:
		return SubdivisionTypeMunicipality
	case SubdivisionML1:
		return SubdivisionTypeRegion
	case SubdivisionML2:
		return SubdivisionTypeRegion
	case SubdivisionML3:
		return SubdivisionTypeRegion
	case SubdivisionML4:
		return SubdivisionTypeRegion
	case SubdivisionML5:
		return SubdivisionTypeRegion
	case SubdivisionML6:
		return SubdivisionTypeRegion
	case SubdivisionML7:
		return SubdivisionTypeRegion
	case SubdivisionML8:
		return SubdivisionTypeRegion
	case SubdivisionMLBK0:
		return SubdivisionTypeDistrict
	case SubdivisionMM01:
		return SubdivisionTypeDivision
	case SubdivisionMM02:
		return SubdivisionTypeDivision
	case SubdivisionMM03:
		return SubdivisionTypeDivision
	case SubdivisionMM04:
		return SubdivisionTypeDivision
	case SubdivisionMM05:
		return SubdivisionTypeDivision
	case SubdivisionMM06:
		return SubdivisionTypeDivision
	case SubdivisionMM07:
		return SubdivisionTypeDivision
	case SubdivisionMM11:
		return SubdivisionTypeState
	case SubdivisionMM12:
		return SubdivisionTypeState
	case SubdivisionMM13:
		return SubdivisionTypeState
	case SubdivisionMM14:
		return SubdivisionTypeState
	case SubdivisionMM15:
		return SubdivisionTypeState
	case SubdivisionMM16:
		return SubdivisionTypeState
	case SubdivisionMM17:
		return SubdivisionTypeState
	case SubdivisionMN035:
		return SubdivisionTypeMunicipality
	case SubdivisionMN037:
		return SubdivisionTypeMunicipality
	case SubdivisionMN039:
		return SubdivisionTypeProvince
	case SubdivisionMN041:
		return SubdivisionTypeProvince
	case SubdivisionMN043:
		return SubdivisionTypeProvince
	case SubdivisionMN046:
		return SubdivisionTypeProvince
	case SubdivisionMN047:
		return SubdivisionTypeProvince
	case SubdivisionMN049:
		return SubdivisionTypeProvince
	case SubdivisionMN051:
		return SubdivisionTypeProvince
	case SubdivisionMN053:
		return SubdivisionTypeProvince
	case SubdivisionMN055:
		return SubdivisionTypeProvince
	case SubdivisionMN057:
		return SubdivisionTypeProvince
	case SubdivisionMN059:
		return SubdivisionTypeProvince
	case SubdivisionMN061:
		return SubdivisionTypeProvince
	case SubdivisionMN063:
		return SubdivisionTypeProvince
	case SubdivisionMN064:
		return SubdivisionTypeMunicipality
	case SubdivisionMN065:
		return SubdivisionTypeProvince
	case SubdivisionMN067:
		return SubdivisionTypeProvince
	case SubdivisionMN069:
		return SubdivisionTypeProvince
	case SubdivisionMN071:
		return SubdivisionTypeProvince
	case SubdivisionMN073:
		return SubdivisionTypeProvince
	case SubdivisionMN1:
		return SubdivisionTypeMunicipality
	case SubdivisionMR01:
		return SubdivisionTypeRegion
	case SubdivisionMR02:
		return SubdivisionTypeRegion
	case SubdivisionMR03:
		return SubdivisionTypeRegion
	case SubdivisionMR04:
		return SubdivisionTypeRegion
	case SubdivisionMR05:
		return SubdivisionTypeRegion
	case SubdivisionMR06:
		return SubdivisionTypeRegion
	case SubdivisionMR07:
		return SubdivisionTypeRegion
	case SubdivisionMR08:
		return SubdivisionTypeRegion
	case SubdivisionMR09:
		return SubdivisionTypeRegion
	case SubdivisionMR10:
		return SubdivisionTypeRegion
	case SubdivisionMR11:
		return SubdivisionTypeRegion
	case SubdivisionMR12:
		return SubdivisionTypeRegion
	case SubdivisionMRNKC:
		return SubdivisionTypeDistrict
	case SubdivisionMT01:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT02:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT03:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT04:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT05:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT06:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT07:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT08:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT09:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT10:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT11:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT12:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT13:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT14:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT15:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT16:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT17:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT18:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT19:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT20:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT21:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT22:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT23:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT24:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT25:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT26:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT27:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT28:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT29:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT30:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT31:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT32:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT33:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT34:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT35:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT36:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT37:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT38:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT39:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT40:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT41:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT42:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT43:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT44:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT45:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT46:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT47:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT48:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT49:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT50:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT51:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT52:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT53:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT54:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT55:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT56:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT57:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT58:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT59:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT60:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT61:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT62:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT63:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT64:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT65:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT66:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT67:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMT68:
		return SubdivisionTypeLocalCouncil
	case SubdivisionMUAG:
		return SubdivisionTypeDependency
	case SubdivisionMUBL:
		return SubdivisionTypeDistrict
	case SubdivisionMUBR:
		return SubdivisionTypeCity
	case SubdivisionMUCC:
		return SubdivisionTypeDependency
	case SubdivisionMUCU:
		return SubdivisionTypeCity
	case SubdivisionMUFL:
		return SubdivisionTypeDistrict
	case SubdivisionMUGP:
		return SubdivisionTypeDistrict
	case SubdivisionMUMO:
		return SubdivisionTypeDistrict
	case SubdivisionMUPA:
		return SubdivisionTypeDistrict
	case SubdivisionMUPL:
		return SubdivisionTypeDistrict
	case SubdivisionMUPU:
		return SubdivisionTypeCity
	case SubdivisionMUPW:
		return SubdivisionTypeDistrict
	case SubdivisionMUQB:
		return SubdivisionTypeCity
	case SubdivisionMURO:
		return SubdivisionTypeDependency
	case SubdivisionMURP:
		return SubdivisionTypeDistrict
	case SubdivisionMUSA:
		return SubdivisionTypeDistrict
	case SubdivisionMUVP:
		return SubdivisionTypeCity
	case SubdivisionMV00:
		return SubdivisionTypeAdministrativeAtoll
	case SubdivisionMV01:
		return SubdivisionTypeAdministrativeAtoll
	case SubdivisionMV02:
		return SubdivisionTypeAdministrativeAtoll
	case SubdivisionMV03:
		return SubdivisionTypeAdministrativeAtoll
	case SubdivisionMV04:
		return SubdivisionTypeAdministrativeAtoll
	case SubdivisionMV05:
		return SubdivisionTypeAdministrativeAtoll
	case SubdivisionMV07:
		return SubdivisionTypeAdministrativeAtoll
	case SubdivisionMV08:
		return SubdivisionTypeAdministrativeAtoll
	case SubdivisionMV12:
		return SubdivisionTypeAdministrativeAtoll
	case SubdivisionMV13:
		return SubdivisionTypeAdministrativeAtoll
	case SubdivisionMV14:
		return SubdivisionTypeAdministrativeAtoll
	case SubdivisionMV17:
		return SubdivisionTypeAdministrativeAtoll
	case SubdivisionMV20:
		return SubdivisionTypeAdministrativeAtoll
	case SubdivisionMV23:
		return SubdivisionTypeAdministrativeAtoll
	case SubdivisionMV24:
		return SubdivisionTypeAdministrativeAtoll
	case SubdivisionMV25:
		return SubdivisionTypeAdministrativeAtoll
	case SubdivisionMV26:
		return SubdivisionTypeAdministrativeAtoll
	case SubdivisionMV27:
		return SubdivisionTypeAdministrativeAtoll
	case SubdivisionMV28:
		return SubdivisionTypeAdministrativeAtoll
	case SubdivisionMV29:
		return SubdivisionTypeAdministrativeAtoll
	case SubdivisionMVCE:
		return SubdivisionTypeProvince
	case SubdivisionMVMLE:
		return SubdivisionTypeCity
	case SubdivisionMVNC:
		return SubdivisionTypeProvince
	case SubdivisionMVNO:
		return SubdivisionTypeProvince
	case SubdivisionMVSC:
		return SubdivisionTypeProvince
	case SubdivisionMVSU:
		return SubdivisionTypeProvince
	case SubdivisionMVUN:
		return SubdivisionTypeProvince
	case SubdivisionMVUS:
		return SubdivisionTypeProvince
	case SubdivisionMWBA:
		return SubdivisionTypeDistrict
	case SubdivisionMWBL:
		return SubdivisionTypeDistrict
	case SubdivisionMWC:
		return SubdivisionTypeRegion
	case SubdivisionMWCK:
		return SubdivisionTypeDistrict
	case SubdivisionMWCR:
		return SubdivisionTypeDistrict
	case SubdivisionMWCT:
		return SubdivisionTypeDistrict
	case SubdivisionMWDE:
		return SubdivisionTypeDistrict
	case SubdivisionMWDO:
		return SubdivisionTypeDistrict
	case SubdivisionMWKR:
		return SubdivisionTypeDistrict
	case SubdivisionMWKS:
		return SubdivisionTypeDistrict
	case SubdivisionMWLI:
		return SubdivisionTypeDistrict
	case SubdivisionMWLK:
		return SubdivisionTypeDistrict
	case SubdivisionMWMC:
		return SubdivisionTypeDistrict
	case SubdivisionMWMG:
		return SubdivisionTypeDistrict
	case SubdivisionMWMH:
		return SubdivisionTypeDistrict
	case SubdivisionMWMU:
		return SubdivisionTypeDistrict
	case SubdivisionMWMW:
		return SubdivisionTypeDistrict
	case SubdivisionMWMZ:
		return SubdivisionTypeDistrict
	case SubdivisionMWN:
		return SubdivisionTypeRegion
	case SubdivisionMWNB:
		return SubdivisionTypeDistrict
	case SubdivisionMWNE:
		return SubdivisionTypeDistrict
	case SubdivisionMWNI:
		return SubdivisionTypeDistrict
	case SubdivisionMWNK:
		return SubdivisionTypeDistrict
	case SubdivisionMWNS:
		return SubdivisionTypeDistrict
	case SubdivisionMWNU:
		return SubdivisionTypeDistrict
	case SubdivisionMWPH:
		return SubdivisionTypeDistrict
	case SubdivisionMWRU:
		return SubdivisionTypeDistrict
	case SubdivisionMWS:
		return SubdivisionTypeRegion
	case SubdivisionMWSA:
		return SubdivisionTypeDistrict
	case SubdivisionMWTH:
		return SubdivisionTypeDistrict
	case SubdivisionMWZO:
		return SubdivisionTypeDistrict
	case SubdivisionMXAGU:
		return SubdivisionTypeState
	case SubdivisionMXBCN:
		return SubdivisionTypeState
	case SubdivisionMXBCS:
		return SubdivisionTypeState
	case SubdivisionMXCAM:
		return SubdivisionTypeState
	case SubdivisionMXCHH:
		return SubdivisionTypeState
	case SubdivisionMXCHP:
		return SubdivisionTypeState
	case SubdivisionMXCMX:
		return SubdivisionTypeFederalDistrict
	case SubdivisionMXCOA:
		return SubdivisionTypeState
	case SubdivisionMXCOL:
		return SubdivisionTypeState
	case SubdivisionMXDUR:
		return SubdivisionTypeState
	case SubdivisionMXGRO:
		return SubdivisionTypeState
	case SubdivisionMXGUA:
		return SubdivisionTypeState
	case SubdivisionMXHID:
		return SubdivisionTypeState
	case SubdivisionMXJAL:
		return SubdivisionTypeState
	case SubdivisionMXMEX:
		return SubdivisionTypeState
	case SubdivisionMXMIC:
		return SubdivisionTypeState
	case SubdivisionMXMOR:
		return SubdivisionTypeState
	case SubdivisionMXNAY:
		return SubdivisionTypeState
	case SubdivisionMXNLE:
		return SubdivisionTypeState
	case SubdivisionMXOAX:
		return SubdivisionTypeState
	case SubdivisionMXPUE:
		return SubdivisionTypeState
	case SubdivisionMXQUE:
		return SubdivisionTypeState
	case SubdivisionMXROO:
		return SubdivisionTypeState
	case SubdivisionMXSIN:
		return SubdivisionTypeState
	case SubdivisionMXSLP:
		return SubdivisionTypeState
	case SubdivisionMXSON:
		return SubdivisionTypeState
	case SubdivisionMXTAB:
		return SubdivisionTypeState
	case SubdivisionMXTAM:
		return SubdivisionTypeState
	case SubdivisionMXTLA:
		return SubdivisionTypeState
	case SubdivisionMXVER:
		return SubdivisionTypeState
	case SubdivisionMXYUC:
		return SubdivisionTypeState
	case SubdivisionMXZAC:
		return SubdivisionTypeState
	case SubdivisionMY01:
		return SubdivisionTypeState
	case SubdivisionMY02:
		return SubdivisionTypeState
	case SubdivisionMY03:
		return SubdivisionTypeState
	case SubdivisionMY04:
		return SubdivisionTypeState
	case SubdivisionMY05:
		return SubdivisionTypeState
	case SubdivisionMY06:
		return SubdivisionTypeState
	case SubdivisionMY07:
		return SubdivisionTypeState
	case SubdivisionMY08:
		return SubdivisionTypeState
	case SubdivisionMY09:
		return SubdivisionTypeState
	case SubdivisionMY10:
		return SubdivisionTypeState
	case SubdivisionMY11:
		return SubdivisionTypeState
	case SubdivisionMY12:
		return SubdivisionTypeState
	case SubdivisionMY13:
		return SubdivisionTypeState
	case SubdivisionMY14:
		return SubdivisionTypeFederalTerritories
	case SubdivisionMY15:
		return SubdivisionTypeFederalTerritories
	case SubdivisionMY16:
		return SubdivisionTypeFederalTerritories
	case SubdivisionMZA:
		return SubdivisionTypeProvince
	case SubdivisionMZB:
		return SubdivisionTypeProvince
	case SubdivisionMZG:
		return SubdivisionTypeProvince
	case SubdivisionMZI:
		return SubdivisionTypeProvince
	case SubdivisionMZL:
		return SubdivisionTypeProvince
	case SubdivisionMZMPM:
		return SubdivisionTypeCity
	case SubdivisionMZN:
		return SubdivisionTypeProvince
	case SubdivisionMZP:
		return SubdivisionTypeProvince
	case SubdivisionMZQ:
		return SubdivisionTypeProvince
	case SubdivisionMZS:
		return SubdivisionTypeProvince
	case SubdivisionMZT:
		return SubdivisionTypeProvince
	case SubdivisionNACA:
		return SubdivisionTypeRegion
	case SubdivisionNAER:
		return SubdivisionTypeRegion
	case SubdivisionNAHA:
		return SubdivisionTypeRegion
	case SubdivisionNAKA:
		return SubdivisionTypeRegion
	case SubdivisionNAKH:
		return SubdivisionTypeRegion
	case SubdivisionNAKU:
		return SubdivisionTypeRegion
	case SubdivisionNAOD:
		return SubdivisionTypeRegion
	case SubdivisionNAOH:
		return SubdivisionTypeRegion
	case SubdivisionNAOK:
		return SubdivisionTypeRegion
	case SubdivisionNAON:
		return SubdivisionTypeRegion
	case SubdivisionNAOS:
		return SubdivisionTypeRegion
	case SubdivisionNAOT:
		return SubdivisionTypeRegion
	case SubdivisionNAOW:
		return SubdivisionTypeRegion
	case SubdivisionNE1:
		return SubdivisionTypeDepartment
	case SubdivisionNE2:
		return SubdivisionTypeDepartment
	case SubdivisionNE3:
		return SubdivisionTypeDepartment
	case SubdivisionNE4:
		return SubdivisionTypeDepartment
	case SubdivisionNE5:
		return SubdivisionTypeDepartment
	case SubdivisionNE6:
		return SubdivisionTypeDepartment
	case SubdivisionNE7:
		return SubdivisionTypeDepartment
	case SubdivisionNE8:
		return SubdivisionTypeCapitalDistrict
	case SubdivisionNGAB:
		return SubdivisionTypeState
	case SubdivisionNGAD:
		return SubdivisionTypeState
	case SubdivisionNGAK:
		return SubdivisionTypeState
	case SubdivisionNGAN:
		return SubdivisionTypeState
	case SubdivisionNGBA:
		return SubdivisionTypeState
	case SubdivisionNGBE:
		return SubdivisionTypeState
	case SubdivisionNGBO:
		return SubdivisionTypeState
	case SubdivisionNGBY:
		return SubdivisionTypeState
	case SubdivisionNGCR:
		return SubdivisionTypeState
	case SubdivisionNGDE:
		return SubdivisionTypeState
	case SubdivisionNGEB:
		return SubdivisionTypeState
	case SubdivisionNGED:
		return SubdivisionTypeState
	case SubdivisionNGEK:
		return SubdivisionTypeState
	case SubdivisionNGEN:
		return SubdivisionTypeState
	case SubdivisionNGFC:
		return SubdivisionTypeCapitalTerritory
	case SubdivisionNGGO:
		return SubdivisionTypeState
	case SubdivisionNGIM:
		return SubdivisionTypeState
	case SubdivisionNGJI:
		return SubdivisionTypeState
	case SubdivisionNGKD:
		return SubdivisionTypeState
	case SubdivisionNGKE:
		return SubdivisionTypeState
	case SubdivisionNGKN:
		return SubdivisionTypeState
	case SubdivisionNGKO:
		return SubdivisionTypeState
	case SubdivisionNGKT:
		return SubdivisionTypeState
	case SubdivisionNGKW:
		return SubdivisionTypeState
	case SubdivisionNGLA:
		return SubdivisionTypeState
	case SubdivisionNGNA:
		return SubdivisionTypeState
	case SubdivisionNGNI:
		return SubdivisionTypeState
	case SubdivisionNGOG:
		return SubdivisionTypeState
	case SubdivisionNGON:
		return SubdivisionTypeState
	case SubdivisionNGOS:
		return SubdivisionTypeState
	case SubdivisionNGOY:
		return SubdivisionTypeState
	case SubdivisionNGPL:
		return SubdivisionTypeState
	case SubdivisionNGRI:
		return SubdivisionTypeState
	case SubdivisionNGSO:
		return SubdivisionTypeState
	case SubdivisionNGTA:
		return SubdivisionTypeState
	case SubdivisionNGYO:
		return SubdivisionTypeState
	case SubdivisionNGZA:
		return SubdivisionTypeState
	case SubdivisionNIAN:
		return SubdivisionTypeAutonomousRegion
	case SubdivisionNIAS:
		return SubdivisionTypeAutonomousRegion
	case SubdivisionNIBO:
		return SubdivisionTypeDepartment
	case SubdivisionNICA:
		return SubdivisionTypeDepartment
	case SubdivisionNICI:
		return SubdivisionTypeDepartment
	case SubdivisionNICO:
		return SubdivisionTypeDepartment
	case SubdivisionNIES:
		return SubdivisionTypeDepartment
	case SubdivisionNIGR:
		return SubdivisionTypeDepartment
	case SubdivisionNIJI:
		return SubdivisionTypeDepartment
	case SubdivisionNILE:
		return SubdivisionTypeDepartment
	case SubdivisionNIMD:
		return SubdivisionTypeDepartment
	case SubdivisionNIMN:
		return SubdivisionTypeDepartment
	case SubdivisionNIMS:
		return SubdivisionTypeDepartment
	case SubdivisionNIMT:
		return SubdivisionTypeDepartment
	case SubdivisionNINS:
		return SubdivisionTypeDepartment
	case SubdivisionNIRI:
		return SubdivisionTypeDepartment
	case SubdivisionNISJ:
		return SubdivisionTypeDepartment
	case SubdivisionNLAW:
		return SubdivisionTypeCountry
	case SubdivisionNLBQ1:
		return SubdivisionTypeSpecialMunicipality
	case SubdivisionNLBQ2:
		return SubdivisionTypeSpecialMunicipality
	case SubdivisionNLBQ3:
		return SubdivisionTypeSpecialMunicipality
	case SubdivisionNLCW:
		return SubdivisionTypeCountry
	case SubdivisionNLDR:
		return SubdivisionTypeProvince
	case SubdivisionNLFL:
		return SubdivisionTypeProvince
	case SubdivisionNLFR:
		return SubdivisionTypeProvince
	case SubdivisionNLGE:
		return SubdivisionTypeProvince
	case SubdivisionNLGR:
		return SubdivisionTypeProvince
	case SubdivisionNLLI:
		return SubdivisionTypeProvince
	case SubdivisionNLNB:
		return SubdivisionTypeProvince
	case SubdivisionNLNH:
		return SubdivisionTypeProvince
	case SubdivisionNLOV:
		return SubdivisionTypeProvince
	case SubdivisionNLSX:
		return SubdivisionTypeCountry
	case SubdivisionNLUT:
		return SubdivisionTypeProvince
	case SubdivisionNLZE:
		return SubdivisionTypeProvince
	case SubdivisionNLZH:
		return SubdivisionTypeProvince
	case SubdivisionNO01:
		return SubdivisionTypeCounty
	case SubdivisionNO02:
		return SubdivisionTypeCounty
	case SubdivisionNO03:
		return SubdivisionTypeCounty
	case SubdivisionNO04:
		return SubdivisionTypeCounty
	case SubdivisionNO05:
		return SubdivisionTypeCounty
	case SubdivisionNO06:
		return SubdivisionTypeCounty
	case SubdivisionNO07:
		return SubdivisionTypeCounty
	case SubdivisionNO08:
		return SubdivisionTypeCounty
	case SubdivisionNO09:
		return SubdivisionTypeCounty
	case SubdivisionNO10:
		return SubdivisionTypeCounty
	case SubdivisionNO11:
		return SubdivisionTypeCounty
	case SubdivisionNO12:
		return SubdivisionTypeCounty
	case SubdivisionNO14:
		return SubdivisionTypeCounty
	case SubdivisionNO15:
		return SubdivisionTypeCounty
	case SubdivisionNO18:
		return SubdivisionTypeCounty
	case SubdivisionNO19:
		return SubdivisionTypeCounty
	case SubdivisionNO20:
		return SubdivisionTypeCounty
	case SubdivisionNO21:
		return SubdivisionTypeArcticRegion
	case SubdivisionNO22:
		return SubdivisionTypeArcticRegion
	case SubdivisionNO50:
		return SubdivisionTypeCounty
	case SubdivisionNP1:
		return SubdivisionTypeDevelopmentRegion
	case SubdivisionNP2:
		return SubdivisionTypeDevelopmentRegion
	case SubdivisionNP3:
		return SubdivisionTypeDevelopmentRegion
	case SubdivisionNP4:
		return SubdivisionTypeDevelopmentRegion
	case SubdivisionNP5:
		return SubdivisionTypeDevelopmentRegion
	case SubdivisionNPBA:
		return SubdivisionTypeZone
	case SubdivisionNPBH:
		return SubdivisionTypeZone
	case SubdivisionNPDH:
		return SubdivisionTypeZone
	case SubdivisionNPGA:
		return SubdivisionTypeZone
	case SubdivisionNPJA:
		return SubdivisionTypeZone
	case SubdivisionNPKA:
		return SubdivisionTypeZone
	case SubdivisionNPKO:
		return SubdivisionTypeZone
	case SubdivisionNPLU:
		return SubdivisionTypeZone
	case SubdivisionNPMA:
		return SubdivisionTypeZone
	case SubdivisionNPME:
		return SubdivisionTypeZone
	case SubdivisionNPNA:
		return SubdivisionTypeZone
	case SubdivisionNPRA:
		return SubdivisionTypeZone
	case SubdivisionNPSA:
		return SubdivisionTypeZone
	case SubdivisionNPSE:
		return SubdivisionTypeZone
	case SubdivisionNR01:
		return SubdivisionTypeDistrict
	case SubdivisionNR02:
		return SubdivisionTypeDistrict
	case SubdivisionNR03:
		return SubdivisionTypeDistrict
	case SubdivisionNR04:
		return SubdivisionTypeDistrict
	case SubdivisionNR05:
		return SubdivisionTypeDistrict
	case SubdivisionNR06:
		return SubdivisionTypeDistrict
	case SubdivisionNR07:
		return SubdivisionTypeDistrict
	case SubdivisionNR08:
		return SubdivisionTypeDistrict
	case SubdivisionNR09:
		return SubdivisionTypeDistrict
	case SubdivisionNR10:
		return SubdivisionTypeDistrict
	case SubdivisionNR11:
		return SubdivisionTypeDistrict
	case SubdivisionNR12:
		return SubdivisionTypeDistrict
	case SubdivisionNR13:
		return SubdivisionTypeDistrict
	case SubdivisionNR14:
		return SubdivisionTypeDistrict
	case SubdivisionNZAUK:
		return SubdivisionTypeRegionalCouncil
	case SubdivisionNZBOP:
		return SubdivisionTypeRegionalCouncil
	case SubdivisionNZCAN:
		return SubdivisionTypeRegionalCouncil
	case SubdivisionNZCIT:
		return SubdivisionTypeSpecialIslandAuthority
	case SubdivisionNZGIS:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionNZHKB:
		return SubdivisionTypeRegionalCouncil
	case SubdivisionNZMBH:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionNZMWT:
		return SubdivisionTypeRegionalCouncil
	case SubdivisionNZN:
		return SubdivisionTypeIsland
	case SubdivisionNZNSN:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionNZNTL:
		return SubdivisionTypeRegionalCouncil
	case SubdivisionNZOTA:
		return SubdivisionTypeRegionalCouncil
	case SubdivisionNZS:
		return SubdivisionTypeIsland
	case SubdivisionNZSTL:
		return SubdivisionTypeRegionalCouncil
	case SubdivisionNZTAS:
		return SubdivisionTypeUnitaryAuthority
	case SubdivisionNZTKI:
		return SubdivisionTypeRegionalCouncil
	case SubdivisionNZWGN:
		return SubdivisionTypeRegionalCouncil
	case SubdivisionNZWKO:
		return SubdivisionTypeRegionalCouncil
	case SubdivisionNZWTC:
		return SubdivisionTypeRegionalCouncil
	case SubdivisionOMBA:
		return SubdivisionTypeRegion
	case SubdivisionOMBU:
		return SubdivisionTypeGovernorate
	case SubdivisionOMDA:
		return SubdivisionTypeRegion
	case SubdivisionOMMA:
		return SubdivisionTypeGovernorate
	case SubdivisionOMMU:
		return SubdivisionTypeGovernorate
	case SubdivisionOMSH:
		return SubdivisionTypeRegion
	case SubdivisionOMWU:
		return SubdivisionTypeRegion
	case SubdivisionOMZA:
		return SubdivisionTypeRegion
	case SubdivisionOMZU:
		return SubdivisionTypeGovernorate
	case SubdivisionPA1:
		return SubdivisionTypeProvince
	case SubdivisionPA2:
		return SubdivisionTypeProvince
	case SubdivisionPA3:
		return SubdivisionTypeProvince
	case SubdivisionPA4:
		return SubdivisionTypeProvince
	case SubdivisionPA5:
		return SubdivisionTypeProvince
	case SubdivisionPA6:
		return SubdivisionTypeProvince
	case SubdivisionPA7:
		return SubdivisionTypeProvince
	case SubdivisionPA8:
		return SubdivisionTypeProvince
	case SubdivisionPA9:
		return SubdivisionTypeProvince
	case SubdivisionPAEM:
		return SubdivisionTypeIndigenousRegion
	case SubdivisionPAKY:
		return SubdivisionTypeIndigenousRegion
	case SubdivisionPANB:
		return SubdivisionTypeIndigenousRegion
	case SubdivisionPEAMA:
		return SubdivisionTypeRegion
	case SubdivisionPEANC:
		return SubdivisionTypeRegion
	case SubdivisionPEAPU:
		return SubdivisionTypeRegion
	case SubdivisionPEARE:
		return SubdivisionTypeRegion
	case SubdivisionPEAYA:
		return SubdivisionTypeRegion
	case SubdivisionPECAJ:
		return SubdivisionTypeRegion
	case SubdivisionPECAL:
		return SubdivisionTypeConstitutionalProvince
	case SubdivisionPECUS:
		return SubdivisionTypeRegion
	case SubdivisionPEHUC:
		return SubdivisionTypeRegion
	case SubdivisionPEHUV:
		return SubdivisionTypeRegion
	case SubdivisionPEICA:
		return SubdivisionTypeRegion
	case SubdivisionPEJUN:
		return SubdivisionTypeRegion
	case SubdivisionPELAL:
		return SubdivisionTypeRegion
	case SubdivisionPELAM:
		return SubdivisionTypeRegion
	case SubdivisionPELIM:
		return SubdivisionTypeRegion
	case SubdivisionPELMA:
		return SubdivisionTypeMunicipality
	case SubdivisionPELOR:
		return SubdivisionTypeRegion
	case SubdivisionPEMDD:
		return SubdivisionTypeRegion
	case SubdivisionPEMOQ:
		return SubdivisionTypeRegion
	case SubdivisionPEPAS:
		return SubdivisionTypeRegion
	case SubdivisionPEPIU:
		return SubdivisionTypeRegion
	case SubdivisionPEPUN:
		return SubdivisionTypeRegion
	case SubdivisionPESAM:
		return SubdivisionTypeRegion
	case SubdivisionPETAC:
		return SubdivisionTypeRegion
	case SubdivisionPETUM:
		return SubdivisionTypeRegion
	case SubdivisionPEUCA:
		return SubdivisionTypeRegion
	case SubdivisionPGCPK:
		return SubdivisionTypeProvince
	case SubdivisionPGCPM:
		return SubdivisionTypeProvince
	case SubdivisionPGEBR:
		return SubdivisionTypeProvince
	case SubdivisionPGEHG:
		return SubdivisionTypeProvince
	case SubdivisionPGEPW:
		return SubdivisionTypeProvince
	case SubdivisionPGESW:
		return SubdivisionTypeProvince
	case SubdivisionPGGPK:
		return SubdivisionTypeProvince
	case SubdivisionPGMBA:
		return SubdivisionTypeProvince
	case SubdivisionPGMPL:
		return SubdivisionTypeProvince
	case SubdivisionPGMPM:
		return SubdivisionTypeProvince
	case SubdivisionPGMRL:
		return SubdivisionTypeProvince
	case SubdivisionPGNCD:
		return SubdivisionTypeDistrict
	case SubdivisionPGNIK:
		return SubdivisionTypeProvince
	case SubdivisionPGNPP:
		return SubdivisionTypeProvince
	case SubdivisionPGNSB:
		return SubdivisionTypeAutonomousRegion
	case SubdivisionPGSAN:
		return SubdivisionTypeProvince
	case SubdivisionPGSHM:
		return SubdivisionTypeProvince
	case SubdivisionPGWBK:
		return SubdivisionTypeProvince
	case SubdivisionPGWHM:
		return SubdivisionTypeProvince
	case SubdivisionPGWPD:
		return SubdivisionTypeProvince
	case SubdivisionPH00:
		return SubdivisionTypeRegion
	case SubdivisionPH01:
		return SubdivisionTypeRegion
	case SubdivisionPH02:
		return SubdivisionTypeRegion
	case SubdivisionPH03:
		return SubdivisionTypeRegion
	case SubdivisionPH05:
		return SubdivisionTypeRegion
	case SubdivisionPH06:
		return SubdivisionTypeRegion
	case SubdivisionPH07:
		return SubdivisionTypeRegion
	case SubdivisionPH08:
		return SubdivisionTypeRegion
	case SubdivisionPH09:
		return SubdivisionTypeRegion
	case SubdivisionPH10:
		return SubdivisionTypeRegion
	case SubdivisionPH11:
		return SubdivisionTypeRegion
	case SubdivisionPH12:
		return SubdivisionTypeRegion
	case SubdivisionPH13:
		return SubdivisionTypeRegion
	case SubdivisionPH14:
		return SubdivisionTypeRegion
	case SubdivisionPH15:
		return SubdivisionTypeRegion
	case SubdivisionPH40:
		return SubdivisionTypeRegion
	case SubdivisionPH41:
		return SubdivisionTypeRegion
	case SubdivisionPHABR:
		return SubdivisionTypeProvince
	case SubdivisionPHAGN:
		return SubdivisionTypeProvince
	case SubdivisionPHAGS:
		return SubdivisionTypeProvince
	case SubdivisionPHAKL:
		return SubdivisionTypeProvince
	case SubdivisionPHALB:
		return SubdivisionTypeProvince
	case SubdivisionPHANT:
		return SubdivisionTypeProvince
	case SubdivisionPHAPA:
		return SubdivisionTypeProvince
	case SubdivisionPHAUR:
		return SubdivisionTypeProvince
	case SubdivisionPHBAN:
		return SubdivisionTypeProvince
	case SubdivisionPHBAS:
		return SubdivisionTypeProvince
	case SubdivisionPHBEN:
		return SubdivisionTypeProvince
	case SubdivisionPHBIL:
		return SubdivisionTypeProvince
	case SubdivisionPHBOH:
		return SubdivisionTypeProvince
	case SubdivisionPHBTG:
		return SubdivisionTypeProvince
	case SubdivisionPHBTN:
		return SubdivisionTypeProvince
	case SubdivisionPHBUK:
		return SubdivisionTypeProvince
	case SubdivisionPHBUL:
		return SubdivisionTypeProvince
	case SubdivisionPHCAG:
		return SubdivisionTypeProvince
	case SubdivisionPHCAM:
		return SubdivisionTypeProvince
	case SubdivisionPHCAN:
		return SubdivisionTypeProvince
	case SubdivisionPHCAP:
		return SubdivisionTypeProvince
	case SubdivisionPHCAS:
		return SubdivisionTypeProvince
	case SubdivisionPHCAT:
		return SubdivisionTypeProvince
	case SubdivisionPHCAV:
		return SubdivisionTypeProvince
	case SubdivisionPHCEB:
		return SubdivisionTypeProvince
	case SubdivisionPHCOM:
		return SubdivisionTypeProvince
	case SubdivisionPHDAO:
		return SubdivisionTypeProvince
	case SubdivisionPHDAS:
		return SubdivisionTypeProvince
	case SubdivisionPHDAV:
		return SubdivisionTypeProvince
	case SubdivisionPHDIN:
		return SubdivisionTypeProvince
	case SubdivisionPHEAS:
		return SubdivisionTypeProvince
	case SubdivisionPHGUI:
		return SubdivisionTypeProvince
	case SubdivisionPHIFU:
		return SubdivisionTypeProvince
	case SubdivisionPHILI:
		return SubdivisionTypeProvince
	case SubdivisionPHILN:
		return SubdivisionTypeProvince
	case SubdivisionPHILS:
		return SubdivisionTypeProvince
	case SubdivisionPHISA:
		return SubdivisionTypeProvince
	case SubdivisionPHKAL:
		return SubdivisionTypeProvince
	case SubdivisionPHLAG:
		return SubdivisionTypeProvince
	case SubdivisionPHLAN:
		return SubdivisionTypeProvince
	case SubdivisionPHLAS:
		return SubdivisionTypeProvince
	case SubdivisionPHLEY:
		return SubdivisionTypeProvince
	case SubdivisionPHLUN:
		return SubdivisionTypeProvince
	case SubdivisionPHMAD:
		return SubdivisionTypeProvince
	case SubdivisionPHMAG:
		return SubdivisionTypeProvince
	case SubdivisionPHMAS:
		return SubdivisionTypeProvince
	case SubdivisionPHMDC:
		return SubdivisionTypeProvince
	case SubdivisionPHMDR:
		return SubdivisionTypeProvince
	case SubdivisionPHMOU:
		return SubdivisionTypeProvince
	case SubdivisionPHMSC:
		return SubdivisionTypeProvince
	case SubdivisionPHMSR:
		return SubdivisionTypeProvince
	case SubdivisionPHNCO:
		return SubdivisionTypeProvince
	case SubdivisionPHNEC:
		return SubdivisionTypeProvince
	case SubdivisionPHNER:
		return SubdivisionTypeProvince
	case SubdivisionPHNSA:
		return SubdivisionTypeProvince
	case SubdivisionPHNUE:
		return SubdivisionTypeProvince
	case SubdivisionPHNUV:
		return SubdivisionTypeProvince
	case SubdivisionPHPAM:
		return SubdivisionTypeProvince
	case SubdivisionPHPAN:
		return SubdivisionTypeProvince
	case SubdivisionPHPLW:
		return SubdivisionTypeProvince
	case SubdivisionPHQUE:
		return SubdivisionTypeProvince
	case SubdivisionPHQUI:
		return SubdivisionTypeProvince
	case SubdivisionPHRIZ:
		return SubdivisionTypeProvince
	case SubdivisionPHROM:
		return SubdivisionTypeProvince
	case SubdivisionPHSAR:
		return SubdivisionTypeProvince
	case SubdivisionPHSCO:
		return SubdivisionTypeProvince
	case SubdivisionPHSIG:
		return SubdivisionTypeProvince
	case SubdivisionPHSLE:
		return SubdivisionTypeProvince
	case SubdivisionPHSLU:
		return SubdivisionTypeProvince
	case SubdivisionPHSOR:
		return SubdivisionTypeProvince
	case SubdivisionPHSUK:
		return SubdivisionTypeProvince
	case SubdivisionPHSUN:
		return SubdivisionTypeProvince
	case SubdivisionPHSUR:
		return SubdivisionTypeProvince
	case SubdivisionPHTAR:
		return SubdivisionTypeProvince
	case SubdivisionPHTAW:
		return SubdivisionTypeProvince
	case SubdivisionPHWSA:
		return SubdivisionTypeProvince
	case SubdivisionPHZAN:
		return SubdivisionTypeProvince
	case SubdivisionPHZAS:
		return SubdivisionTypeProvince
	case SubdivisionPHZMB:
		return SubdivisionTypeProvince
	case SubdivisionPHZSI:
		return SubdivisionTypeProvince
	case SubdivisionPKBA:
		return SubdivisionTypeProvince
	case SubdivisionPKGB:
		return SubdivisionTypeArea
	case SubdivisionPKIS:
		return SubdivisionTypeCapitalTerritory
	case SubdivisionPKJK:
		return SubdivisionTypeArea
	case SubdivisionPKKP:
		return SubdivisionTypeProvince
	case SubdivisionPKPB:
		return SubdivisionTypeProvince
	case SubdivisionPKSD:
		return SubdivisionTypeProvince
	case SubdivisionPKTA:
		return SubdivisionTypeTerritory
	case SubdivisionPLDS:
		return SubdivisionTypeProvince
	case SubdivisionPLKP:
		return SubdivisionTypeProvince
	case SubdivisionPLLB:
		return SubdivisionTypeProvince
	case SubdivisionPLLD:
		return SubdivisionTypeProvince
	case SubdivisionPLLU:
		return SubdivisionTypeProvince
	case SubdivisionPLMA:
		return SubdivisionTypeProvince
	case SubdivisionPLMZ:
		return SubdivisionTypeProvince
	case SubdivisionPLOP:
		return SubdivisionTypeProvince
	case SubdivisionPLPD:
		return SubdivisionTypeProvince
	case SubdivisionPLPK:
		return SubdivisionTypeProvince
	case SubdivisionPLPM:
		return SubdivisionTypeProvince
	case SubdivisionPLSK:
		return SubdivisionTypeProvince
	case SubdivisionPLSL:
		return SubdivisionTypeProvince
	case SubdivisionPLWN:
		return SubdivisionTypeProvince
	case SubdivisionPLWP:
		return SubdivisionTypeProvince
	case SubdivisionPLZP:
		return SubdivisionTypeProvince
	case SubdivisionPSBTH:
		return SubdivisionTypeGovernorate
	case SubdivisionPSDEB:
		return SubdivisionTypeGovernorate
	case SubdivisionPSGZA:
		return SubdivisionTypeGovernorate
	case SubdivisionPSHBN:
		return SubdivisionTypeGovernorate
	case SubdivisionPSJEM:
		return SubdivisionTypeGovernorate
	case SubdivisionPSJEN:
		return SubdivisionTypeGovernorate
	case SubdivisionPSJRH:
		return SubdivisionTypeGovernorate
	case SubdivisionPSKYS:
		return SubdivisionTypeGovernorate
	case SubdivisionPSNBS:
		return SubdivisionTypeGovernorate
	case SubdivisionPSNGZ:
		return SubdivisionTypeGovernorate
	case SubdivisionPSQQA:
		return SubdivisionTypeGovernorate
	case SubdivisionPSRBH:
		return SubdivisionTypeGovernorate
	case SubdivisionPSRFH:
		return SubdivisionTypeGovernorate
	case SubdivisionPSSLT:
		return SubdivisionTypeGovernorate
	case SubdivisionPSTBS:
		return SubdivisionTypeGovernorate
	case SubdivisionPSTKM:
		return SubdivisionTypeGovernorate
	case SubdivisionPT01:
		return SubdivisionTypeDistrict
	case SubdivisionPT02:
		return SubdivisionTypeDistrict
	case SubdivisionPT03:
		return SubdivisionTypeDistrict
	case SubdivisionPT04:
		return SubdivisionTypeDistrict
	case SubdivisionPT05:
		return SubdivisionTypeDistrict
	case SubdivisionPT06:
		return SubdivisionTypeDistrict
	case SubdivisionPT07:
		return SubdivisionTypeDistrict
	case SubdivisionPT08:
		return SubdivisionTypeDistrict
	case SubdivisionPT09:
		return SubdivisionTypeDistrict
	case SubdivisionPT10:
		return SubdivisionTypeDistrict
	case SubdivisionPT11:
		return SubdivisionTypeDistrict
	case SubdivisionPT12:
		return SubdivisionTypeDistrict
	case SubdivisionPT13:
		return SubdivisionTypeDistrict
	case SubdivisionPT14:
		return SubdivisionTypeDistrict
	case SubdivisionPT15:
		return SubdivisionTypeDistrict
	case SubdivisionPT16:
		return SubdivisionTypeDistrict
	case SubdivisionPT17:
		return SubdivisionTypeDistrict
	case SubdivisionPT18:
		return SubdivisionTypeDistrict
	case SubdivisionPT20:
		return SubdivisionTypeAutonomousRegion
	case SubdivisionPT30:
		return SubdivisionTypeAutonomousRegion
	case SubdivisionPW002:
		return SubdivisionTypeState
	case SubdivisionPW004:
		return SubdivisionTypeState
	case SubdivisionPW010:
		return SubdivisionTypeState
	case SubdivisionPW050:
		return SubdivisionTypeState
	case SubdivisionPW100:
		return SubdivisionTypeState
	case SubdivisionPW150:
		return SubdivisionTypeState
	case SubdivisionPW212:
		return SubdivisionTypeState
	case SubdivisionPW214:
		return SubdivisionTypeState
	case SubdivisionPW218:
		return SubdivisionTypeState
	case SubdivisionPW222:
		return SubdivisionTypeState
	case SubdivisionPW224:
		return SubdivisionTypeState
	case SubdivisionPW226:
		return SubdivisionTypeState
	case SubdivisionPW227:
		return SubdivisionTypeState
	case SubdivisionPW228:
		return SubdivisionTypeState
	case SubdivisionPW350:
		return SubdivisionTypeState
	case SubdivisionPW370:
		return SubdivisionTypeState
	case SubdivisionPY1:
		return SubdivisionTypeDepartment
	case SubdivisionPY10:
		return SubdivisionTypeDepartment
	case SubdivisionPY11:
		return SubdivisionTypeDepartment
	case SubdivisionPY12:
		return SubdivisionTypeDepartment
	case SubdivisionPY13:
		return SubdivisionTypeDepartment
	case SubdivisionPY14:
		return SubdivisionTypeDepartment
	case SubdivisionPY15:
		return SubdivisionTypeDepartment
	case SubdivisionPY16:
		return SubdivisionTypeDepartment
	case SubdivisionPY19:
		return SubdivisionTypeDepartment
	case SubdivisionPY2:
		return SubdivisionTypeDepartment
	case SubdivisionPY3:
		return SubdivisionTypeDepartment
	case SubdivisionPY4:
		return SubdivisionTypeDepartment
	case SubdivisionPY5:
		return SubdivisionTypeDepartment
	case SubdivisionPY6:
		return SubdivisionTypeDepartment
	case SubdivisionPY7:
		return SubdivisionTypeDepartment
	case SubdivisionPY8:
		return SubdivisionTypeDepartment
	case SubdivisionPY9:
		return SubdivisionTypeDepartment
	case SubdivisionPYASU:
		return SubdivisionTypeCapitalDistrict
	case SubdivisionQADA:
		return SubdivisionTypeMunicipality
	case SubdivisionQAKH:
		return SubdivisionTypeMunicipality
	case SubdivisionQAMS:
		return SubdivisionTypeMunicipality
	case SubdivisionQARA:
		return SubdivisionTypeMunicipality
	case SubdivisionQAUS:
		return SubdivisionTypeMunicipality
	case SubdivisionQAWA:
		return SubdivisionTypeMunicipality
	case SubdivisionQAZA:
		return SubdivisionTypeMunicipality
	case SubdivisionROAB:
		return SubdivisionTypeDepartment
	case SubdivisionROAG:
		return SubdivisionTypeDepartment
	case SubdivisionROAR:
		return SubdivisionTypeDepartment
	case SubdivisionROB:
		return SubdivisionTypeMunicipality
	case SubdivisionROBC:
		return SubdivisionTypeDepartment
	case SubdivisionROBH:
		return SubdivisionTypeDepartment
	case SubdivisionROBN:
		return SubdivisionTypeDepartment
	case SubdivisionROBR:
		return SubdivisionTypeDepartment
	case SubdivisionROBT:
		return SubdivisionTypeDepartment
	case SubdivisionROBV:
		return SubdivisionTypeDepartment
	case SubdivisionROBZ:
		return SubdivisionTypeDepartment
	case SubdivisionROCJ:
		return SubdivisionTypeDepartment
	case SubdivisionROCL:
		return SubdivisionTypeDepartment
	case SubdivisionROCS:
		return SubdivisionTypeDepartment
	case SubdivisionROCT:
		return SubdivisionTypeDepartment
	case SubdivisionROCV:
		return SubdivisionTypeDepartment
	case SubdivisionRODB:
		return SubdivisionTypeDepartment
	case SubdivisionRODJ:
		return SubdivisionTypeDepartment
	case SubdivisionROGJ:
		return SubdivisionTypeDepartment
	case SubdivisionROGL:
		return SubdivisionTypeDepartment
	case SubdivisionROGR:
		return SubdivisionTypeDepartment
	case SubdivisionROHD:
		return SubdivisionTypeDepartment
	case SubdivisionROHR:
		return SubdivisionTypeDepartment
	case SubdivisionROIF:
		return SubdivisionTypeDepartment
	case SubdivisionROIL:
		return SubdivisionTypeDepartment
	case SubdivisionROIS:
		return SubdivisionTypeDepartment
	case SubdivisionROMH:
		return SubdivisionTypeDepartment
	case SubdivisionROMM:
		return SubdivisionTypeDepartment
	case SubdivisionROMS:
		return SubdivisionTypeDepartment
	case SubdivisionRONT:
		return SubdivisionTypeDepartment
	case SubdivisionROOT:
		return SubdivisionTypeDepartment
	case SubdivisionROPH:
		return SubdivisionTypeDepartment
	case SubdivisionROSB:
		return SubdivisionTypeDepartment
	case SubdivisionROSJ:
		return SubdivisionTypeDepartment
	case SubdivisionROSM:
		return SubdivisionTypeDepartment
	case SubdivisionROSV:
		return SubdivisionTypeDepartment
	case SubdivisionROTL:
		return SubdivisionTypeDepartment
	case SubdivisionROTM:
		return SubdivisionTypeDepartment
	case SubdivisionROTR:
		return SubdivisionTypeDepartment
	case SubdivisionROVL:
		return SubdivisionTypeDepartment
	case SubdivisionROVN:
		return SubdivisionTypeDepartment
	case SubdivisionROVS:
		return SubdivisionTypeDepartment
	case SubdivisionRS00:
		return SubdivisionTypeCity
	case SubdivisionRS01:
		return SubdivisionTypeDistrict
	case SubdivisionRS02:
		return SubdivisionTypeDistrict
	case SubdivisionRS03:
		return SubdivisionTypeDistrict
	case SubdivisionRS04:
		return SubdivisionTypeDistrict
	case SubdivisionRS05:
		return SubdivisionTypeDistrict
	case SubdivisionRS06:
		return SubdivisionTypeDistrict
	case SubdivisionRS07:
		return SubdivisionTypeDistrict
	case SubdivisionRS08:
		return SubdivisionTypeDistrict
	case SubdivisionRS09:
		return SubdivisionTypeDistrict
	case SubdivisionRS10:
		return SubdivisionTypeDistrict
	case SubdivisionRS11:
		return SubdivisionTypeDistrict
	case SubdivisionRS12:
		return SubdivisionTypeDistrict
	case SubdivisionRS13:
		return SubdivisionTypeDistrict
	case SubdivisionRS14:
		return SubdivisionTypeDistrict
	case SubdivisionRS15:
		return SubdivisionTypeDistrict
	case SubdivisionRS16:
		return SubdivisionTypeDistrict
	case SubdivisionRS17:
		return SubdivisionTypeDistrict
	case SubdivisionRS18:
		return SubdivisionTypeDistrict
	case SubdivisionRS19:
		return SubdivisionTypeDistrict
	case SubdivisionRS20:
		return SubdivisionTypeDistrict
	case SubdivisionRS21:
		return SubdivisionTypeDistrict
	case SubdivisionRS22:
		return SubdivisionTypeDistrict
	case SubdivisionRS23:
		return SubdivisionTypeDistrict
	case SubdivisionRS24:
		return SubdivisionTypeDistrict
	case SubdivisionRS25:
		return SubdivisionTypeDistrict
	case SubdivisionRS26:
		return SubdivisionTypeDistrict
	case SubdivisionRS27:
		return SubdivisionTypeDistrict
	case SubdivisionRS28:
		return SubdivisionTypeDistrict
	case SubdivisionRS29:
		return SubdivisionTypeDistrict
	case SubdivisionRSKM:
		return SubdivisionTypeAutonomousProvince
	case SubdivisionRSVO:
		return SubdivisionTypeAutonomousProvince
	case SubdivisionRUAD:
		return SubdivisionTypeRepublic
	case SubdivisionRUAL:
		return SubdivisionTypeRepublic
	case SubdivisionRUALT:
		return SubdivisionTypeAdministrativeTerritory
	case SubdivisionRUAMU:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUARK:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUAST:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUBA:
		return SubdivisionTypeRepublic
	case SubdivisionRUBEL:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUBRY:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUBU:
		return SubdivisionTypeRepublic
	case SubdivisionRUCE:
		return SubdivisionTypeRepublic
	case SubdivisionRUCHE:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUCHU:
		return SubdivisionTypeAutonomousDistrict
	case SubdivisionRUCU:
		return SubdivisionTypeRepublic
	case SubdivisionRUDA:
		return SubdivisionTypeRepublic
	case SubdivisionRUIN:
		return SubdivisionTypeRepublic
	case SubdivisionRUIRK:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUIVA:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUKAM:
		return SubdivisionTypeAdministrativeTerritory
	case SubdivisionRUKB:
		return SubdivisionTypeRepublic
	case SubdivisionRUKC:
		return SubdivisionTypeRepublic
	case SubdivisionRUKDA:
		return SubdivisionTypeAdministrativeTerritory
	case SubdivisionRUKEM:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUKGD:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUKGN:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUKHA:
		return SubdivisionTypeAdministrativeTerritory
	case SubdivisionRUKHM:
		return SubdivisionTypeAutonomousDistrict
	case SubdivisionRUKIR:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUKK:
		return SubdivisionTypeRepublic
	case SubdivisionRUKL:
		return SubdivisionTypeRepublic
	case SubdivisionRUKLU:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUKO:
		return SubdivisionTypeRepublic
	case SubdivisionRUKOS:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUKR:
		return SubdivisionTypeRepublic
	case SubdivisionRUKRS:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUKYA:
		return SubdivisionTypeAdministrativeTerritory
	case SubdivisionRULEN:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRULIP:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUMAG:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUME:
		return SubdivisionTypeRepublic
	case SubdivisionRUMO:
		return SubdivisionTypeRepublic
	case SubdivisionRUMOS:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUMOW:
		return SubdivisionTypeAutonomousCity
	case SubdivisionRUMUR:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUNEN:
		return SubdivisionTypeAutonomousDistrict
	case SubdivisionRUNGR:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUNIZ:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUNVS:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUOMS:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUORE:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUORL:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUPER:
		return SubdivisionTypeAdministrativeTerritory
	case SubdivisionRUPNZ:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUPRI:
		return SubdivisionTypeAdministrativeTerritory
	case SubdivisionRUPSK:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUROS:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRURYA:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUSA:
		return SubdivisionTypeRepublic
	case SubdivisionRUSAK:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUSAM:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUSAR:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUSE:
		return SubdivisionTypeRepublic
	case SubdivisionRUSMO:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUSPE:
		return SubdivisionTypeAutonomousCity
	case SubdivisionRUSTA:
		return SubdivisionTypeAdministrativeTerritory
	case SubdivisionRUSVE:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUTA:
		return SubdivisionTypeRepublic
	case SubdivisionRUTAM:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUTOM:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUTUL:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUTVE:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUTY:
		return SubdivisionTypeRepublic
	case SubdivisionRUTYU:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUUD:
		return SubdivisionTypeRepublic
	case SubdivisionRUULY:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUVGG:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUVLA:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUVLG:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUVOR:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUYAN:
		return SubdivisionTypeAutonomousDistrict
	case SubdivisionRUYAR:
		return SubdivisionTypeAdministrativeRegion
	case SubdivisionRUYEV:
		return SubdivisionTypeAutonomousRegion
	case SubdivisionRUZAB:
		return SubdivisionTypeAdministrativeTerritory
	case SubdivisionRW01:
		return SubdivisionTypeTownCouncil
	case SubdivisionRW02:
		return SubdivisionTypeProvince
	case SubdivisionRW03:
		return SubdivisionTypeProvince
	case SubdivisionRW04:
		return SubdivisionTypeProvince
	case SubdivisionRW05:
		return SubdivisionTypeProvince
	case SubdivisionSA01:
		return SubdivisionTypeProvince
	case SubdivisionSA02:
		return SubdivisionTypeProvince
	case SubdivisionSA03:
		return SubdivisionTypeProvince
	case SubdivisionSA04:
		return SubdivisionTypeProvince
	case SubdivisionSA05:
		return SubdivisionTypeProvince
	case SubdivisionSA06:
		return SubdivisionTypeProvince
	case SubdivisionSA07:
		return SubdivisionTypeProvince
	case SubdivisionSA08:
		return SubdivisionTypeProvince
	case SubdivisionSA09:
		return SubdivisionTypeProvince
	case SubdivisionSA10:
		return SubdivisionTypeProvince
	case SubdivisionSA11:
		return SubdivisionTypeProvince
	case SubdivisionSA12:
		return SubdivisionTypeProvince
	case SubdivisionSA14:
		return SubdivisionTypeProvince
	case SubdivisionSBCE:
		return SubdivisionTypeProvince
	case SubdivisionSBCH:
		return SubdivisionTypeProvince
	case SubdivisionSBCT:
		return SubdivisionTypeCapitalTerritory
	case SubdivisionSBGU:
		return SubdivisionTypeProvince
	case SubdivisionSBIS:
		return SubdivisionTypeProvince
	case SubdivisionSBMK:
		return SubdivisionTypeProvince
	case SubdivisionSBML:
		return SubdivisionTypeProvince
	case SubdivisionSBRB:
		return SubdivisionTypeProvince
	case SubdivisionSBTE:
		return SubdivisionTypeProvince
	case SubdivisionSBWE:
		return SubdivisionTypeProvince
	case SubdivisionSC01:
		return SubdivisionTypeDistrict
	case SubdivisionSC02:
		return SubdivisionTypeDistrict
	case SubdivisionSC03:
		return SubdivisionTypeDistrict
	case SubdivisionSC04:
		return SubdivisionTypeDistrict
	case SubdivisionSC05:
		return SubdivisionTypeDistrict
	case SubdivisionSC06:
		return SubdivisionTypeDistrict
	case SubdivisionSC07:
		return SubdivisionTypeDistrict
	case SubdivisionSC08:
		return SubdivisionTypeDistrict
	case SubdivisionSC09:
		return SubdivisionTypeDistrict
	case SubdivisionSC10:
		return SubdivisionTypeDistrict
	case SubdivisionSC11:
		return SubdivisionTypeDistrict
	case SubdivisionSC12:
		return SubdivisionTypeDistrict
	case SubdivisionSC13:
		return SubdivisionTypeDistrict
	case SubdivisionSC14:
		return SubdivisionTypeDistrict
	case SubdivisionSC15:
		return SubdivisionTypeDistrict
	case SubdivisionSC16:
		return SubdivisionTypeDistrict
	case SubdivisionSC17:
		return SubdivisionTypeDistrict
	case SubdivisionSC18:
		return SubdivisionTypeDistrict
	case SubdivisionSC19:
		return SubdivisionTypeDistrict
	case SubdivisionSC20:
		return SubdivisionTypeDistrict
	case SubdivisionSC21:
		return SubdivisionTypeDistrict
	case SubdivisionSC22:
		return SubdivisionTypeDistrict
	case SubdivisionSC23:
		return SubdivisionTypeDistrict
	case SubdivisionSC24:
		return SubdivisionTypeDistrict
	case SubdivisionSC25:
		return SubdivisionTypeDistrict
	case SubdivisionSDDC:
		return SubdivisionTypeState
	case SubdivisionSDDE:
		return SubdivisionTypeState
	case SubdivisionSDDN:
		return SubdivisionTypeState
	case SubdivisionSDDS:
		return SubdivisionTypeState
	case SubdivisionSDDW:
		return SubdivisionTypeState
	case SubdivisionSDGD:
		return SubdivisionTypeState
	case SubdivisionSDGZ:
		return SubdivisionTypeState
	case SubdivisionSDKA:
		return SubdivisionTypeState
	case SubdivisionSDKH:
		return SubdivisionTypeState
	case SubdivisionSDKN:
		return SubdivisionTypeState
	case SubdivisionSDKS:
		return SubdivisionTypeState
	case SubdivisionSDNB:
		return SubdivisionTypeState
	case SubdivisionSDNO:
		return SubdivisionTypeState
	case SubdivisionSDNR:
		return SubdivisionTypeState
	case SubdivisionSDNW:
		return SubdivisionTypeState
	case SubdivisionSDRS:
		return SubdivisionTypeState
	case SubdivisionSDSI:
		return SubdivisionTypeState
	case SubdivisionSEAB:
		return SubdivisionTypeCounty
	case SubdivisionSEAC:
		return SubdivisionTypeCounty
	case SubdivisionSEBD:
		return SubdivisionTypeCounty
	case SubdivisionSEC:
		return SubdivisionTypeCounty
	case SubdivisionSED:
		return SubdivisionTypeCounty
	case SubdivisionSEE:
		return SubdivisionTypeCounty
	case SubdivisionSEF:
		return SubdivisionTypeCounty
	case SubdivisionSEG:
		return SubdivisionTypeCounty
	case SubdivisionSEH:
		return SubdivisionTypeCounty
	case SubdivisionSEI:
		return SubdivisionTypeCounty
	case SubdivisionSEK:
		return SubdivisionTypeCounty
	case SubdivisionSEM:
		return SubdivisionTypeCounty
	case SubdivisionSEN:
		return SubdivisionTypeCounty
	case SubdivisionSEO:
		return SubdivisionTypeCounty
	case SubdivisionSES:
		return SubdivisionTypeCounty
	case SubdivisionSET:
		return SubdivisionTypeCounty
	case SubdivisionSEU:
		return SubdivisionTypeCounty
	case SubdivisionSEW:
		return SubdivisionTypeCounty
	case SubdivisionSEX:
		return SubdivisionTypeCounty
	case SubdivisionSEY:
		return SubdivisionTypeCounty
	case SubdivisionSEZ:
		return SubdivisionTypeCounty
	case SubdivisionSG01:
		return SubdivisionTypeDistrict
	case SubdivisionSG02:
		return SubdivisionTypeDistrict
	case SubdivisionSG03:
		return SubdivisionTypeDistrict
	case SubdivisionSG04:
		return SubdivisionTypeDistrict
	case SubdivisionSG05:
		return SubdivisionTypeDistrict
	case SubdivisionSHAC:
		return SubdivisionTypeGeographicalEntity
	case SubdivisionSHHL:
		return SubdivisionTypeGeographicalEntity
	case SubdivisionSHTA:
		return SubdivisionTypeGeographicalEntity
	case SubdivisionSI001:
		return SubdivisionTypeMunicipality
	case SubdivisionSI002:
		return SubdivisionTypeMunicipality
	case SubdivisionSI003:
		return SubdivisionTypeMunicipality
	case SubdivisionSI004:
		return SubdivisionTypeMunicipality
	case SubdivisionSI005:
		return SubdivisionTypeMunicipality
	case SubdivisionSI006:
		return SubdivisionTypeMunicipality
	case SubdivisionSI007:
		return SubdivisionTypeMunicipality
	case SubdivisionSI008:
		return SubdivisionTypeMunicipality
	case SubdivisionSI009:
		return SubdivisionTypeMunicipality
	case SubdivisionSI010:
		return SubdivisionTypeMunicipality
	case SubdivisionSI011:
		return SubdivisionTypeMunicipality
	case SubdivisionSI012:
		return SubdivisionTypeMunicipality
	case SubdivisionSI013:
		return SubdivisionTypeMunicipality
	case SubdivisionSI014:
		return SubdivisionTypeMunicipality
	case SubdivisionSI015:
		return SubdivisionTypeMunicipality
	case SubdivisionSI016:
		return SubdivisionTypeMunicipality
	case SubdivisionSI017:
		return SubdivisionTypeMunicipality
	case SubdivisionSI018:
		return SubdivisionTypeMunicipality
	case SubdivisionSI019:
		return SubdivisionTypeMunicipality
	case SubdivisionSI020:
		return SubdivisionTypeMunicipality
	case SubdivisionSI021:
		return SubdivisionTypeMunicipality
	case SubdivisionSI022:
		return SubdivisionTypeMunicipality
	case SubdivisionSI023:
		return SubdivisionTypeMunicipality
	case SubdivisionSI024:
		return SubdivisionTypeMunicipality
	case SubdivisionSI025:
		return SubdivisionTypeMunicipality
	case SubdivisionSI026:
		return SubdivisionTypeMunicipality
	case SubdivisionSI027:
		return SubdivisionTypeMunicipality
	case SubdivisionSI028:
		return SubdivisionTypeMunicipality
	case SubdivisionSI029:
		return SubdivisionTypeMunicipality
	case SubdivisionSI030:
		return SubdivisionTypeMunicipality
	case SubdivisionSI031:
		return SubdivisionTypeMunicipality
	case SubdivisionSI032:
		return SubdivisionTypeMunicipality
	case SubdivisionSI033:
		return SubdivisionTypeMunicipality
	case SubdivisionSI034:
		return SubdivisionTypeMunicipality
	case SubdivisionSI035:
		return SubdivisionTypeMunicipality
	case SubdivisionSI036:
		return SubdivisionTypeMunicipality
	case SubdivisionSI037:
		return SubdivisionTypeMunicipality
	case SubdivisionSI038:
		return SubdivisionTypeMunicipality
	case SubdivisionSI039:
		return SubdivisionTypeMunicipality
	case SubdivisionSI040:
		return SubdivisionTypeMunicipality
	case SubdivisionSI041:
		return SubdivisionTypeMunicipality
	case SubdivisionSI042:
		return SubdivisionTypeMunicipality
	case SubdivisionSI043:
		return SubdivisionTypeMunicipality
	case SubdivisionSI044:
		return SubdivisionTypeMunicipality
	case SubdivisionSI045:
		return SubdivisionTypeMunicipality
	case SubdivisionSI046:
		return SubdivisionTypeMunicipality
	case SubdivisionSI047:
		return SubdivisionTypeMunicipality
	case SubdivisionSI048:
		return SubdivisionTypeMunicipality
	case SubdivisionSI049:
		return SubdivisionTypeMunicipality
	case SubdivisionSI050:
		return SubdivisionTypeMunicipality
	case SubdivisionSI051:
		return SubdivisionTypeMunicipality
	case SubdivisionSI052:
		return SubdivisionTypeMunicipality
	case SubdivisionSI053:
		return SubdivisionTypeMunicipality
	case SubdivisionSI054:
		return SubdivisionTypeMunicipality
	case SubdivisionSI055:
		return SubdivisionTypeMunicipality
	case SubdivisionSI056:
		return SubdivisionTypeMunicipality
	case SubdivisionSI057:
		return SubdivisionTypeMunicipality
	case SubdivisionSI058:
		return SubdivisionTypeMunicipality
	case SubdivisionSI059:
		return SubdivisionTypeMunicipality
	case SubdivisionSI060:
		return SubdivisionTypeMunicipality
	case SubdivisionSI061:
		return SubdivisionTypeMunicipality
	case SubdivisionSI062:
		return SubdivisionTypeMunicipality
	case SubdivisionSI063:
		return SubdivisionTypeMunicipality
	case SubdivisionSI064:
		return SubdivisionTypeMunicipality
	case SubdivisionSI065:
		return SubdivisionTypeMunicipality
	case SubdivisionSI066:
		return SubdivisionTypeMunicipality
	case SubdivisionSI067:
		return SubdivisionTypeMunicipality
	case SubdivisionSI068:
		return SubdivisionTypeMunicipality
	case SubdivisionSI069:
		return SubdivisionTypeMunicipality
	case SubdivisionSI070:
		return SubdivisionTypeMunicipality
	case SubdivisionSI071:
		return SubdivisionTypeMunicipality
	case SubdivisionSI072:
		return SubdivisionTypeMunicipality
	case SubdivisionSI073:
		return SubdivisionTypeMunicipality
	case SubdivisionSI074:
		return SubdivisionTypeMunicipality
	case SubdivisionSI075:
		return SubdivisionTypeMunicipality
	case SubdivisionSI076:
		return SubdivisionTypeMunicipality
	case SubdivisionSI077:
		return SubdivisionTypeMunicipality
	case SubdivisionSI078:
		return SubdivisionTypeMunicipality
	case SubdivisionSI079:
		return SubdivisionTypeMunicipality
	case SubdivisionSI080:
		return SubdivisionTypeMunicipality
	case SubdivisionSI081:
		return SubdivisionTypeMunicipality
	case SubdivisionSI082:
		return SubdivisionTypeMunicipality
	case SubdivisionSI083:
		return SubdivisionTypeMunicipality
	case SubdivisionSI084:
		return SubdivisionTypeMunicipality
	case SubdivisionSI085:
		return SubdivisionTypeMunicipality
	case SubdivisionSI086:
		return SubdivisionTypeMunicipality
	case SubdivisionSI087:
		return SubdivisionTypeMunicipality
	case SubdivisionSI088:
		return SubdivisionTypeMunicipality
	case SubdivisionSI089:
		return SubdivisionTypeMunicipality
	case SubdivisionSI090:
		return SubdivisionTypeMunicipality
	case SubdivisionSI091:
		return SubdivisionTypeMunicipality
	case SubdivisionSI092:
		return SubdivisionTypeMunicipality
	case SubdivisionSI093:
		return SubdivisionTypeMunicipality
	case SubdivisionSI094:
		return SubdivisionTypeMunicipality
	case SubdivisionSI095:
		return SubdivisionTypeMunicipality
	case SubdivisionSI096:
		return SubdivisionTypeMunicipality
	case SubdivisionSI097:
		return SubdivisionTypeMunicipality
	case SubdivisionSI098:
		return SubdivisionTypeMunicipality
	case SubdivisionSI099:
		return SubdivisionTypeMunicipality
	case SubdivisionSI100:
		return SubdivisionTypeMunicipality
	case SubdivisionSI101:
		return SubdivisionTypeMunicipality
	case SubdivisionSI102:
		return SubdivisionTypeMunicipality
	case SubdivisionSI103:
		return SubdivisionTypeMunicipality
	case SubdivisionSI104:
		return SubdivisionTypeMunicipality
	case SubdivisionSI105:
		return SubdivisionTypeMunicipality
	case SubdivisionSI106:
		return SubdivisionTypeMunicipality
	case SubdivisionSI107:
		return SubdivisionTypeMunicipality
	case SubdivisionSI108:
		return SubdivisionTypeMunicipality
	case SubdivisionSI109:
		return SubdivisionTypeMunicipality
	case SubdivisionSI110:
		return SubdivisionTypeMunicipality
	case SubdivisionSI111:
		return SubdivisionTypeMunicipality
	case SubdivisionSI112:
		return SubdivisionTypeMunicipality
	case SubdivisionSI113:
		return SubdivisionTypeMunicipality
	case SubdivisionSI114:
		return SubdivisionTypeMunicipality
	case SubdivisionSI115:
		return SubdivisionTypeMunicipality
	case SubdivisionSI116:
		return SubdivisionTypeMunicipality
	case SubdivisionSI117:
		return SubdivisionTypeMunicipality
	case SubdivisionSI118:
		return SubdivisionTypeMunicipality
	case SubdivisionSI119:
		return SubdivisionTypeMunicipality
	case SubdivisionSI120:
		return SubdivisionTypeMunicipality
	case SubdivisionSI121:
		return SubdivisionTypeMunicipality
	case SubdivisionSI122:
		return SubdivisionTypeMunicipality
	case SubdivisionSI123:
		return SubdivisionTypeMunicipality
	case SubdivisionSI124:
		return SubdivisionTypeMunicipality
	case SubdivisionSI125:
		return SubdivisionTypeMunicipality
	case SubdivisionSI126:
		return SubdivisionTypeMunicipality
	case SubdivisionSI127:
		return SubdivisionTypeMunicipality
	case SubdivisionSI128:
		return SubdivisionTypeMunicipality
	case SubdivisionSI129:
		return SubdivisionTypeMunicipality
	case SubdivisionSI130:
		return SubdivisionTypeMunicipality
	case SubdivisionSI131:
		return SubdivisionTypeMunicipality
	case SubdivisionSI132:
		return SubdivisionTypeMunicipality
	case SubdivisionSI133:
		return SubdivisionTypeMunicipality
	case SubdivisionSI134:
		return SubdivisionTypeMunicipality
	case SubdivisionSI135:
		return SubdivisionTypeMunicipality
	case SubdivisionSI136:
		return SubdivisionTypeMunicipality
	case SubdivisionSI137:
		return SubdivisionTypeMunicipality
	case SubdivisionSI138:
		return SubdivisionTypeMunicipality
	case SubdivisionSI139:
		return SubdivisionTypeMunicipality
	case SubdivisionSI140:
		return SubdivisionTypeMunicipality
	case SubdivisionSI141:
		return SubdivisionTypeMunicipality
	case SubdivisionSI142:
		return SubdivisionTypeMunicipality
	case SubdivisionSI143:
		return SubdivisionTypeMunicipality
	case SubdivisionSI144:
		return SubdivisionTypeMunicipality
	case SubdivisionSI146:
		return SubdivisionTypeMunicipality
	case SubdivisionSI147:
		return SubdivisionTypeMunicipality
	case SubdivisionSI148:
		return SubdivisionTypeMunicipality
	case SubdivisionSI149:
		return SubdivisionTypeMunicipality
	case SubdivisionSI150:
		return SubdivisionTypeMunicipality
	case SubdivisionSI151:
		return SubdivisionTypeMunicipality
	case SubdivisionSI152:
		return SubdivisionTypeMunicipality
	case SubdivisionSI153:
		return SubdivisionTypeMunicipality
	case SubdivisionSI154:
		return SubdivisionTypeMunicipality
	case SubdivisionSI155:
		return SubdivisionTypeMunicipality
	case SubdivisionSI156:
		return SubdivisionTypeMunicipality
	case SubdivisionSI157:
		return SubdivisionTypeMunicipality
	case SubdivisionSI158:
		return SubdivisionTypeMunicipality
	case SubdivisionSI159:
		return SubdivisionTypeMunicipality
	case SubdivisionSI160:
		return SubdivisionTypeMunicipality
	case SubdivisionSI161:
		return SubdivisionTypeMunicipality
	case SubdivisionSI162:
		return SubdivisionTypeMunicipality
	case SubdivisionSI163:
		return SubdivisionTypeMunicipality
	case SubdivisionSI164:
		return SubdivisionTypeMunicipality
	case SubdivisionSI165:
		return SubdivisionTypeMunicipality
	case SubdivisionSI166:
		return SubdivisionTypeMunicipality
	case SubdivisionSI167:
		return SubdivisionTypeMunicipality
	case SubdivisionSI168:
		return SubdivisionTypeMunicipality
	case SubdivisionSI169:
		return SubdivisionTypeMunicipality
	case SubdivisionSI170:
		return SubdivisionTypeMunicipality
	case SubdivisionSI171:
		return SubdivisionTypeMunicipality
	case SubdivisionSI172:
		return SubdivisionTypeMunicipality
	case SubdivisionSI173:
		return SubdivisionTypeMunicipality
	case SubdivisionSI174:
		return SubdivisionTypeMunicipality
	case SubdivisionSI175:
		return SubdivisionTypeMunicipality
	case SubdivisionSI176:
		return SubdivisionTypeMunicipality
	case SubdivisionSI177:
		return SubdivisionTypeMunicipality
	case SubdivisionSI178:
		return SubdivisionTypeMunicipality
	case SubdivisionSI179:
		return SubdivisionTypeMunicipality
	case SubdivisionSI180:
		return SubdivisionTypeMunicipality
	case SubdivisionSI181:
		return SubdivisionTypeMunicipality
	case SubdivisionSI182:
		return SubdivisionTypeMunicipality
	case SubdivisionSI183:
		return SubdivisionTypeMunicipality
	case SubdivisionSI184:
		return SubdivisionTypeMunicipality
	case SubdivisionSI185:
		return SubdivisionTypeMunicipality
	case SubdivisionSI186:
		return SubdivisionTypeMunicipality
	case SubdivisionSI187:
		return SubdivisionTypeMunicipality
	case SubdivisionSI188:
		return SubdivisionTypeMunicipality
	case SubdivisionSI189:
		return SubdivisionTypeMunicipality
	case SubdivisionSI190:
		return SubdivisionTypeMunicipality
	case SubdivisionSI191:
		return SubdivisionTypeMunicipality
	case SubdivisionSI192:
		return SubdivisionTypeMunicipality
	case SubdivisionSI193:
		return SubdivisionTypeMunicipality
	case SubdivisionSI194:
		return SubdivisionTypeMunicipality
	case SubdivisionSI195:
		return SubdivisionTypeMunicipality
	case SubdivisionSI196:
		return SubdivisionTypeMunicipality
	case SubdivisionSI197:
		return SubdivisionTypeMunicipality
	case SubdivisionSI198:
		return SubdivisionTypeMunicipality
	case SubdivisionSI199:
		return SubdivisionTypeMunicipality
	case SubdivisionSI200:
		return SubdivisionTypeMunicipality
	case SubdivisionSI201:
		return SubdivisionTypeMunicipality
	case SubdivisionSI202:
		return SubdivisionTypeMunicipality
	case SubdivisionSI203:
		return SubdivisionTypeMunicipality
	case SubdivisionSI204:
		return SubdivisionTypeMunicipality
	case SubdivisionSI205:
		return SubdivisionTypeMunicipality
	case SubdivisionSI206:
		return SubdivisionTypeMunicipality
	case SubdivisionSI207:
		return SubdivisionTypeMunicipality
	case SubdivisionSI208:
		return SubdivisionTypeMunicipality
	case SubdivisionSI209:
		return SubdivisionTypeMunicipality
	case SubdivisionSI210:
		return SubdivisionTypeMunicipality
	case SubdivisionSI211:
		return SubdivisionTypeMunicipality
	case SubdivisionSKBC:
		return SubdivisionTypeRegion
	case SubdivisionSKBL:
		return SubdivisionTypeRegion
	case SubdivisionSKKI:
		return SubdivisionTypeRegion
	case SubdivisionSKNI:
		return SubdivisionTypeRegion
	case SubdivisionSKPV:
		return SubdivisionTypeRegion
	case SubdivisionSKTA:
		return SubdivisionTypeRegion
	case SubdivisionSKTC:
		return SubdivisionTypeRegion
	case SubdivisionSKZI:
		return SubdivisionTypeRegion
	case SubdivisionSLE:
		return SubdivisionTypeProvince
	case SubdivisionSLN:
		return SubdivisionTypeProvince
	case SubdivisionSLS:
		return SubdivisionTypeProvince
	case SubdivisionSLW:
		return SubdivisionTypeArea
	case SubdivisionSM01:
		return SubdivisionTypeMunicipalities
	case SubdivisionSM02:
		return SubdivisionTypeMunicipalities
	case SubdivisionSM03:
		return SubdivisionTypeMunicipalities
	case SubdivisionSM04:
		return SubdivisionTypeMunicipalities
	case SubdivisionSM05:
		return SubdivisionTypeMunicipalities
	case SubdivisionSM06:
		return SubdivisionTypeMunicipalities
	case SubdivisionSM07:
		return SubdivisionTypeMunicipalities
	case SubdivisionSM08:
		return SubdivisionTypeMunicipalities
	case SubdivisionSM09:
		return SubdivisionTypeMunicipalities
	case SubdivisionSNDB:
		return SubdivisionTypeRegion
	case SubdivisionSNDK:
		return SubdivisionTypeRegion
	case SubdivisionSNFK:
		return SubdivisionTypeRegion
	case SubdivisionSNKA:
		return SubdivisionTypeRegion
	case SubdivisionSNKD:
		return SubdivisionTypeRegion
	case SubdivisionSNKE:
		return SubdivisionTypeRegion
	case SubdivisionSNKL:
		return SubdivisionTypeRegion
	case SubdivisionSNLG:
		return SubdivisionTypeRegion
	case SubdivisionSNMT:
		return SubdivisionTypeRegion
	case SubdivisionSNSE:
		return SubdivisionTypeRegion
	case SubdivisionSNSL:
		return SubdivisionTypeRegion
	case SubdivisionSNTC:
		return SubdivisionTypeRegion
	case SubdivisionSNTH:
		return SubdivisionTypeRegion
	case SubdivisionSNZG:
		return SubdivisionTypeRegion
	case SubdivisionSOAW:
		return SubdivisionTypeRegion
	case SubdivisionSOBK:
		return SubdivisionTypeRegion
	case SubdivisionSOBN:
		return SubdivisionTypeRegion
	case SubdivisionSOBR:
		return SubdivisionTypeRegion
	case SubdivisionSOBY:
		return SubdivisionTypeRegion
	case SubdivisionSOGA:
		return SubdivisionTypeRegion
	case SubdivisionSOGE:
		return SubdivisionTypeRegion
	case SubdivisionSOHI:
		return SubdivisionTypeRegion
	case SubdivisionSOJD:
		return SubdivisionTypeRegion
	case SubdivisionSOJH:
		return SubdivisionTypeRegion
	case SubdivisionSOMU:
		return SubdivisionTypeRegion
	case SubdivisionSONU:
		return SubdivisionTypeRegion
	case SubdivisionSOSA:
		return SubdivisionTypeRegion
	case SubdivisionSOSD:
		return SubdivisionTypeRegion
	case SubdivisionSOSH:
		return SubdivisionTypeRegion
	case SubdivisionSOSO:
		return SubdivisionTypeRegion
	case SubdivisionSOTO:
		return SubdivisionTypeRegion
	case SubdivisionSOWO:
		return SubdivisionTypeRegion
	case SubdivisionSRBR:
		return SubdivisionTypeDistrict
	case SubdivisionSRCM:
		return SubdivisionTypeDistrict
	case SubdivisionSRCR:
		return SubdivisionTypeDistrict
	case SubdivisionSRMA:
		return SubdivisionTypeDistrict
	case SubdivisionSRNI:
		return SubdivisionTypeDistrict
	case SubdivisionSRPM:
		return SubdivisionTypeDistrict
	case SubdivisionSRPR:
		return SubdivisionTypeDistrict
	case SubdivisionSRSA:
		return SubdivisionTypeDistrict
	case SubdivisionSRSI:
		return SubdivisionTypeDistrict
	case SubdivisionSRWA:
		return SubdivisionTypeDistrict
	case SubdivisionSSBN:
		return SubdivisionTypeState
	case SubdivisionSSBW:
		return SubdivisionTypeState
	case SubdivisionSSEC:
		return SubdivisionTypeState
	case SubdivisionSSEE:
		return SubdivisionTypeState
	case SubdivisionSSEW:
		return SubdivisionTypeState
	case SubdivisionSSJG:
		return SubdivisionTypeState
	case SubdivisionSSLK:
		return SubdivisionTypeState
	case SubdivisionSSNU:
		return SubdivisionTypeState
	case SubdivisionSSUY:
		return SubdivisionTypeState
	case SubdivisionSSWR:
		return SubdivisionTypeState
	case SubdivisionSTP:
		return SubdivisionTypeMunicipality
	case SubdivisionSTS:
		return SubdivisionTypeMunicipality
	case SubdivisionSVAH:
		return SubdivisionTypeDepartment
	case SubdivisionSVCA:
		return SubdivisionTypeDepartment
	case SubdivisionSVCH:
		return SubdivisionTypeDepartment
	case SubdivisionSVCU:
		return SubdivisionTypeDepartment
	case SubdivisionSVLI:
		return SubdivisionTypeDepartment
	case SubdivisionSVMO:
		return SubdivisionTypeDepartment
	case SubdivisionSVPA:
		return SubdivisionTypeDepartment
	case SubdivisionSVSA:
		return SubdivisionTypeDepartment
	case SubdivisionSVSM:
		return SubdivisionTypeDepartment
	case SubdivisionSVSO:
		return SubdivisionTypeDepartment
	case SubdivisionSVSS:
		return SubdivisionTypeDepartment
	case SubdivisionSVSV:
		return SubdivisionTypeDepartment
	case SubdivisionSVUN:
		return SubdivisionTypeDepartment
	case SubdivisionSVUS:
		return SubdivisionTypeDepartment
	case SubdivisionSYDI:
		return SubdivisionTypeGovernorate
	case SubdivisionSYDR:
		return SubdivisionTypeGovernorate
	case SubdivisionSYDY:
		return SubdivisionTypeGovernorate
	case SubdivisionSYHA:
		return SubdivisionTypeGovernorate
	case SubdivisionSYHI:
		return SubdivisionTypeGovernorate
	case SubdivisionSYHL:
		return SubdivisionTypeGovernorate
	case SubdivisionSYHM:
		return SubdivisionTypeGovernorate
	case SubdivisionSYID:
		return SubdivisionTypeGovernorate
	case SubdivisionSYLA:
		return SubdivisionTypeGovernorate
	case SubdivisionSYQU:
		return SubdivisionTypeGovernorate
	case SubdivisionSYRA:
		return SubdivisionTypeGovernorate
	case SubdivisionSYRD:
		return SubdivisionTypeGovernorate
	case SubdivisionSYSU:
		return SubdivisionTypeGovernorate
	case SubdivisionSYTA:
		return SubdivisionTypeGovernorate
	case SubdivisionSZHH:
		return SubdivisionTypeDistrict
	case SubdivisionSZLU:
		return SubdivisionTypeDistrict
	case SubdivisionSZMA:
		return SubdivisionTypeDistrict
	case SubdivisionSZSH:
		return SubdivisionTypeDistrict
	case SubdivisionTDBA:
		return SubdivisionTypeRegion
	case SubdivisionTDBG:
		return SubdivisionTypeRegion
	case SubdivisionTDBO:
		return SubdivisionTypeRegion
	case SubdivisionTDCB:
		return SubdivisionTypeRegion
	case SubdivisionTDEN:
		return SubdivisionTypeRegion
	case SubdivisionTDGR:
		return SubdivisionTypeRegion
	case SubdivisionTDHL:
		return SubdivisionTypeRegion
	case SubdivisionTDKA:
		return SubdivisionTypeRegion
	case SubdivisionTDLC:
		return SubdivisionTypeRegion
	case SubdivisionTDLO:
		return SubdivisionTypeRegion
	case SubdivisionTDLR:
		return SubdivisionTypeRegion
	case SubdivisionTDMA:
		return SubdivisionTypeRegion
	case SubdivisionTDMC:
		return SubdivisionTypeRegion
	case SubdivisionTDME:
		return SubdivisionTypeRegion
	case SubdivisionTDMO:
		return SubdivisionTypeRegion
	case SubdivisionTDND:
		return SubdivisionTypeRegion
	case SubdivisionTDOD:
		return SubdivisionTypeRegion
	case SubdivisionTDSA:
		return SubdivisionTypeRegion
	case SubdivisionTDSI:
		return SubdivisionTypeRegion
	case SubdivisionTDTA:
		return SubdivisionTypeRegion
	case SubdivisionTDTI:
		return SubdivisionTypeRegion
	case SubdivisionTDWF:
		return SubdivisionTypeRegion
	case SubdivisionTGC:
		return SubdivisionTypeRegion
	case SubdivisionTGK:
		return SubdivisionTypeRegion
	case SubdivisionTGM:
		return SubdivisionTypeRegion
	case SubdivisionTGP:
		return SubdivisionTypeRegion
	case SubdivisionTGS:
		return SubdivisionTypeRegion
	case SubdivisionTH10:
		return SubdivisionTypeMunicipality
	case SubdivisionTH11:
		return SubdivisionTypeProvince
	case SubdivisionTH12:
		return SubdivisionTypeProvince
	case SubdivisionTH13:
		return SubdivisionTypeProvince
	case SubdivisionTH14:
		return SubdivisionTypeProvince
	case SubdivisionTH15:
		return SubdivisionTypeProvince
	case SubdivisionTH16:
		return SubdivisionTypeProvince
	case SubdivisionTH17:
		return SubdivisionTypeProvince
	case SubdivisionTH18:
		return SubdivisionTypeProvince
	case SubdivisionTH19:
		return SubdivisionTypeProvince
	case SubdivisionTH20:
		return SubdivisionTypeProvince
	case SubdivisionTH21:
		return SubdivisionTypeProvince
	case SubdivisionTH22:
		return SubdivisionTypeProvince
	case SubdivisionTH23:
		return SubdivisionTypeProvince
	case SubdivisionTH24:
		return SubdivisionTypeProvince
	case SubdivisionTH25:
		return SubdivisionTypeProvince
	case SubdivisionTH26:
		return SubdivisionTypeProvince
	case SubdivisionTH27:
		return SubdivisionTypeProvince
	case SubdivisionTH30:
		return SubdivisionTypeProvince
	case SubdivisionTH31:
		return SubdivisionTypeProvince
	case SubdivisionTH32:
		return SubdivisionTypeProvince
	case SubdivisionTH33:
		return SubdivisionTypeProvince
	case SubdivisionTH34:
		return SubdivisionTypeProvince
	case SubdivisionTH35:
		return SubdivisionTypeProvince
	case SubdivisionTH36:
		return SubdivisionTypeProvince
	case SubdivisionTH37:
		return SubdivisionTypeProvince
	case SubdivisionTH39:
		return SubdivisionTypeProvince
	case SubdivisionTH40:
		return SubdivisionTypeProvince
	case SubdivisionTH41:
		return SubdivisionTypeProvince
	case SubdivisionTH42:
		return SubdivisionTypeProvince
	case SubdivisionTH43:
		return SubdivisionTypeProvince
	case SubdivisionTH44:
		return SubdivisionTypeProvince
	case SubdivisionTH45:
		return SubdivisionTypeProvince
	case SubdivisionTH46:
		return SubdivisionTypeProvince
	case SubdivisionTH47:
		return SubdivisionTypeProvince
	case SubdivisionTH48:
		return SubdivisionTypeProvince
	case SubdivisionTH49:
		return SubdivisionTypeProvince
	case SubdivisionTH50:
		return SubdivisionTypeProvince
	case SubdivisionTH51:
		return SubdivisionTypeProvince
	case SubdivisionTH52:
		return SubdivisionTypeProvince
	case SubdivisionTH53:
		return SubdivisionTypeProvince
	case SubdivisionTH54:
		return SubdivisionTypeProvince
	case SubdivisionTH55:
		return SubdivisionTypeProvince
	case SubdivisionTH56:
		return SubdivisionTypeProvince
	case SubdivisionTH57:
		return SubdivisionTypeProvince
	case SubdivisionTH58:
		return SubdivisionTypeProvince
	case SubdivisionTH60:
		return SubdivisionTypeProvince
	case SubdivisionTH61:
		return SubdivisionTypeProvince
	case SubdivisionTH62:
		return SubdivisionTypeProvince
	case SubdivisionTH63:
		return SubdivisionTypeProvince
	case SubdivisionTH64:
		return SubdivisionTypeProvince
	case SubdivisionTH65:
		return SubdivisionTypeProvince
	case SubdivisionTH66:
		return SubdivisionTypeProvince
	case SubdivisionTH67:
		return SubdivisionTypeProvince
	case SubdivisionTH70:
		return SubdivisionTypeProvince
	case SubdivisionTH71:
		return SubdivisionTypeProvince
	case SubdivisionTH72:
		return SubdivisionTypeProvince
	case SubdivisionTH73:
		return SubdivisionTypeProvince
	case SubdivisionTH74:
		return SubdivisionTypeProvince
	case SubdivisionTH75:
		return SubdivisionTypeProvince
	case SubdivisionTH76:
		return SubdivisionTypeProvince
	case SubdivisionTH77:
		return SubdivisionTypeProvince
	case SubdivisionTH80:
		return SubdivisionTypeProvince
	case SubdivisionTH81:
		return SubdivisionTypeProvince
	case SubdivisionTH82:
		return SubdivisionTypeProvince
	case SubdivisionTH83:
		return SubdivisionTypeProvince
	case SubdivisionTH84:
		return SubdivisionTypeProvince
	case SubdivisionTH85:
		return SubdivisionTypeProvince
	case SubdivisionTH86:
		return SubdivisionTypeProvince
	case SubdivisionTH90:
		return SubdivisionTypeProvince
	case SubdivisionTH91:
		return SubdivisionTypeProvince
	case SubdivisionTH92:
		return SubdivisionTypeProvince
	case SubdivisionTH93:
		return SubdivisionTypeProvince
	case SubdivisionTH94:
		return SubdivisionTypeProvince
	case SubdivisionTH95:
		return SubdivisionTypeProvince
	case SubdivisionTH96:
		return SubdivisionTypeProvince
	case SubdivisionTHS:
		return SubdivisionTypeProvince
	case SubdivisionTJGB:
		return SubdivisionTypeAutonomousRegion
	case SubdivisionTJKT:
		return SubdivisionTypeRegion
	case SubdivisionTJSU:
		return SubdivisionTypeRegion
	case SubdivisionTLAL:
		return SubdivisionTypeDistrict
	case SubdivisionTLAN:
		return SubdivisionTypeDistrict
	case SubdivisionTLBA:
		return SubdivisionTypeDistrict
	case SubdivisionTLBO:
		return SubdivisionTypeDistrict
	case SubdivisionTLCO:
		return SubdivisionTypeDistrict
	case SubdivisionTLDI:
		return SubdivisionTypeDistrict
	case SubdivisionTLER:
		return SubdivisionTypeDistrict
	case SubdivisionTLLA:
		return SubdivisionTypeDistrict
	case SubdivisionTLLI:
		return SubdivisionTypeDistrict
	case SubdivisionTLMF:
		return SubdivisionTypeDistrict
	case SubdivisionTLMT:
		return SubdivisionTypeDistrict
	case SubdivisionTLOE:
		return SubdivisionTypeDistrict
	case SubdivisionTLVI:
		return SubdivisionTypeDistrict
	case SubdivisionTMA:
		return SubdivisionTypeRegion
	case SubdivisionTMB:
		return SubdivisionTypeRegion
	case SubdivisionTMD:
		return SubdivisionTypeRegion
	case SubdivisionTML:
		return SubdivisionTypeRegion
	case SubdivisionTMM:
		return SubdivisionTypeRegion
	case SubdivisionTMS:
		return SubdivisionTypeCity
	case SubdivisionTN11:
		return SubdivisionTypeGovernorate
	case SubdivisionTN12:
		return SubdivisionTypeGovernorate
	case SubdivisionTN13:
		return SubdivisionTypeGovernorate
	case SubdivisionTN14:
		return SubdivisionTypeGovernorate
	case SubdivisionTN21:
		return SubdivisionTypeGovernorate
	case SubdivisionTN22:
		return SubdivisionTypeGovernorate
	case SubdivisionTN23:
		return SubdivisionTypeGovernorate
	case SubdivisionTN31:
		return SubdivisionTypeGovernorate
	case SubdivisionTN32:
		return SubdivisionTypeGovernorate
	case SubdivisionTN33:
		return SubdivisionTypeGovernorate
	case SubdivisionTN34:
		return SubdivisionTypeGovernorate
	case SubdivisionTN41:
		return SubdivisionTypeGovernorate
	case SubdivisionTN42:
		return SubdivisionTypeGovernorate
	case SubdivisionTN43:
		return SubdivisionTypeGovernorate
	case SubdivisionTN51:
		return SubdivisionTypeGovernorate
	case SubdivisionTN52:
		return SubdivisionTypeGovernorate
	case SubdivisionTN53:
		return SubdivisionTypeGovernorate
	case SubdivisionTN61:
		return SubdivisionTypeGovernorate
	case SubdivisionTN71:
		return SubdivisionTypeGovernorate
	case SubdivisionTN72:
		return SubdivisionTypeGovernorate
	case SubdivisionTN73:
		return SubdivisionTypeGovernorate
	case SubdivisionTN81:
		return SubdivisionTypeGovernorate
	case SubdivisionTN82:
		return SubdivisionTypeGovernorate
	case SubdivisionTN83:
		return SubdivisionTypeGovernorate
	case SubdivisionTO01:
		return SubdivisionTypeDivision
	case SubdivisionTO02:
		return SubdivisionTypeDivision
	case SubdivisionTO03:
		return SubdivisionTypeDivision
	case SubdivisionTO04:
		return SubdivisionTypeDivision
	case SubdivisionTO05:
		return SubdivisionTypeDivision
	case SubdivisionTR01:
		return SubdivisionTypeProvince
	case SubdivisionTR02:
		return SubdivisionTypeProvince
	case SubdivisionTR03:
		return SubdivisionTypeProvince
	case SubdivisionTR04:
		return SubdivisionTypeProvince
	case SubdivisionTR05:
		return SubdivisionTypeProvince
	case SubdivisionTR06:
		return SubdivisionTypeProvince
	case SubdivisionTR07:
		return SubdivisionTypeProvince
	case SubdivisionTR08:
		return SubdivisionTypeProvince
	case SubdivisionTR09:
		return SubdivisionTypeProvince
	case SubdivisionTR10:
		return SubdivisionTypeProvince
	case SubdivisionTR11:
		return SubdivisionTypeProvince
	case SubdivisionTR12:
		return SubdivisionTypeProvince
	case SubdivisionTR13:
		return SubdivisionTypeProvince
	case SubdivisionTR14:
		return SubdivisionTypeProvince
	case SubdivisionTR15:
		return SubdivisionTypeProvince
	case SubdivisionTR16:
		return SubdivisionTypeProvince
	case SubdivisionTR17:
		return SubdivisionTypeProvince
	case SubdivisionTR18:
		return SubdivisionTypeProvince
	case SubdivisionTR19:
		return SubdivisionTypeProvince
	case SubdivisionTR20:
		return SubdivisionTypeProvince
	case SubdivisionTR21:
		return SubdivisionTypeProvince
	case SubdivisionTR22:
		return SubdivisionTypeProvince
	case SubdivisionTR23:
		return SubdivisionTypeProvince
	case SubdivisionTR24:
		return SubdivisionTypeProvince
	case SubdivisionTR25:
		return SubdivisionTypeProvince
	case SubdivisionTR26:
		return SubdivisionTypeProvince
	case SubdivisionTR27:
		return SubdivisionTypeProvince
	case SubdivisionTR28:
		return SubdivisionTypeProvince
	case SubdivisionTR29:
		return SubdivisionTypeProvince
	case SubdivisionTR30:
		return SubdivisionTypeProvince
	case SubdivisionTR31:
		return SubdivisionTypeProvince
	case SubdivisionTR32:
		return SubdivisionTypeProvince
	case SubdivisionTR33:
		return SubdivisionTypeProvince
	case SubdivisionTR34:
		return SubdivisionTypeProvince
	case SubdivisionTR35:
		return SubdivisionTypeProvince
	case SubdivisionTR36:
		return SubdivisionTypeProvince
	case SubdivisionTR37:
		return SubdivisionTypeProvince
	case SubdivisionTR38:
		return SubdivisionTypeProvince
	case SubdivisionTR39:
		return SubdivisionTypeProvince
	case SubdivisionTR40:
		return SubdivisionTypeProvince
	case SubdivisionTR41:
		return SubdivisionTypeProvince
	case SubdivisionTR42:
		return SubdivisionTypeProvince
	case SubdivisionTR43:
		return SubdivisionTypeProvince
	case SubdivisionTR44:
		return SubdivisionTypeProvince
	case SubdivisionTR45:
		return SubdivisionTypeProvince
	case SubdivisionTR46:
		return SubdivisionTypeProvince
	case SubdivisionTR47:
		return SubdivisionTypeProvince
	case SubdivisionTR48:
		return SubdivisionTypeProvince
	case SubdivisionTR49:
		return SubdivisionTypeProvince
	case SubdivisionTR50:
		return SubdivisionTypeProvince
	case SubdivisionTR51:
		return SubdivisionTypeProvince
	case SubdivisionTR52:
		return SubdivisionTypeProvince
	case SubdivisionTR53:
		return SubdivisionTypeProvince
	case SubdivisionTR54:
		return SubdivisionTypeProvince
	case SubdivisionTR55:
		return SubdivisionTypeProvince
	case SubdivisionTR56:
		return SubdivisionTypeProvince
	case SubdivisionTR57:
		return SubdivisionTypeProvince
	case SubdivisionTR58:
		return SubdivisionTypeProvince
	case SubdivisionTR59:
		return SubdivisionTypeProvince
	case SubdivisionTR60:
		return SubdivisionTypeProvince
	case SubdivisionTR61:
		return SubdivisionTypeProvince
	case SubdivisionTR62:
		return SubdivisionTypeProvince
	case SubdivisionTR63:
		return SubdivisionTypeProvince
	case SubdivisionTR64:
		return SubdivisionTypeProvince
	case SubdivisionTR65:
		return SubdivisionTypeProvince
	case SubdivisionTR66:
		return SubdivisionTypeProvince
	case SubdivisionTR67:
		return SubdivisionTypeProvince
	case SubdivisionTR68:
		return SubdivisionTypeProvince
	case SubdivisionTR69:
		return SubdivisionTypeProvince
	case SubdivisionTR70:
		return SubdivisionTypeProvince
	case SubdivisionTR71:
		return SubdivisionTypeProvince
	case SubdivisionTR72:
		return SubdivisionTypeProvince
	case SubdivisionTR73:
		return SubdivisionTypeProvince
	case SubdivisionTR74:
		return SubdivisionTypeProvince
	case SubdivisionTR75:
		return SubdivisionTypeProvince
	case SubdivisionTR76:
		return SubdivisionTypeProvince
	case SubdivisionTR77:
		return SubdivisionTypeProvince
	case SubdivisionTR78:
		return SubdivisionTypeProvince
	case SubdivisionTR79:
		return SubdivisionTypeProvince
	case SubdivisionTR80:
		return SubdivisionTypeProvince
	case SubdivisionTR81:
		return SubdivisionTypeProvince
	case SubdivisionTTARI:
		return SubdivisionTypeBorough
	case SubdivisionTTCHA:
		return SubdivisionTypeBorough
	case SubdivisionTTCTT:
		return SubdivisionTypeRegion
	case SubdivisionTTDMN:
		return SubdivisionTypeRegion
	case SubdivisionTTETO:
		return SubdivisionTypeRegion
	case SubdivisionTTPED:
		return SubdivisionTypeRegion
	case SubdivisionTTPOS:
		return SubdivisionTypeCity
	case SubdivisionTTPRT:
		return SubdivisionTypeRegion
	case SubdivisionTTPTF:
		return SubdivisionTypeBorough
	case SubdivisionTTRCM:
		return SubdivisionTypeRegion
	case SubdivisionTTSFO:
		return SubdivisionTypeCity
	case SubdivisionTTSGE:
		return SubdivisionTypeRegion
	case SubdivisionTTSIP:
		return SubdivisionTypeRegion
	case SubdivisionTTSJL:
		return SubdivisionTypeRegion
	case SubdivisionTTTUP:
		return SubdivisionTypeRegion
	case SubdivisionTTWTO:
		return SubdivisionTypeRegion
	case SubdivisionTVFUN:
		return SubdivisionTypeTownCouncil
	case SubdivisionTVNIT:
		return SubdivisionTypeIslandCouncil
	case SubdivisionTVNKF:
		return SubdivisionTypeIslandCouncil
	case SubdivisionTVNKL:
		return SubdivisionTypeIslandCouncil
	case SubdivisionTVNMA:
		return SubdivisionTypeIslandCouncil
	case SubdivisionTVNMG:
		return SubdivisionTypeIslandCouncil
	case SubdivisionTVNUI:
		return SubdivisionTypeIslandCouncil
	case SubdivisionTVVAI:
		return SubdivisionTypeIslandCouncil
	case SubdivisionTWCHA:
		return SubdivisionTypeDistrict
	case SubdivisionTWCYI:
		return SubdivisionTypeMunicipality
	case SubdivisionTWCYQ:
		return SubdivisionTypeDistrict
	case SubdivisionTWHSQ:
		return SubdivisionTypeDistrict
	case SubdivisionTWHSZ:
		return SubdivisionTypeMunicipality
	case SubdivisionTWHUA:
		return SubdivisionTypeDistrict
	case SubdivisionTWILA:
		return SubdivisionTypeDistrict
	case SubdivisionTWKEE:
		return SubdivisionTypeMunicipality
	case SubdivisionTWKHH:
		return SubdivisionTypeSpecialMunicipality
	case SubdivisionTWKHQ:
		return SubdivisionTypeDistrict
	case SubdivisionTWMIA:
		return SubdivisionTypeDistrict
	case SubdivisionTWNAN:
		return SubdivisionTypeDistrict
	case SubdivisionTWPEN:
		return SubdivisionTypeDistrict
	case SubdivisionTWPIF:
		return SubdivisionTypeDistrict
	case SubdivisionTWTAO:
		return SubdivisionTypeDistrict
	case SubdivisionTWTNN:
		return SubdivisionTypeMunicipality
	case SubdivisionTWTNQ:
		return SubdivisionTypeDistrict
	case SubdivisionTWTPE:
		return SubdivisionTypeSpecialMunicipality
	case SubdivisionTWTPQ:
		return SubdivisionTypeDistrict
	case SubdivisionTWTTT:
		return SubdivisionTypeDistrict
	case SubdivisionTWTXG:
		return SubdivisionTypeMunicipality
	case SubdivisionTWTXQ:
		return SubdivisionTypeDistrict
	case SubdivisionTWYUN:
		return SubdivisionTypeDistrict
	case SubdivisionTZ01:
		return SubdivisionTypeRegion
	case SubdivisionTZ02:
		return SubdivisionTypeRegion
	case SubdivisionTZ03:
		return SubdivisionTypeRegion
	case SubdivisionTZ04:
		return SubdivisionTypeRegion
	case SubdivisionTZ05:
		return SubdivisionTypeRegion
	case SubdivisionTZ06:
		return SubdivisionTypeRegion
	case SubdivisionTZ07:
		return SubdivisionTypeRegion
	case SubdivisionTZ08:
		return SubdivisionTypeRegion
	case SubdivisionTZ09:
		return SubdivisionTypeRegion
	case SubdivisionTZ10:
		return SubdivisionTypeRegion
	case SubdivisionTZ11:
		return SubdivisionTypeRegion
	case SubdivisionTZ12:
		return SubdivisionTypeRegion
	case SubdivisionTZ13:
		return SubdivisionTypeRegion
	case SubdivisionTZ14:
		return SubdivisionTypeRegion
	case SubdivisionTZ15:
		return SubdivisionTypeRegion
	case SubdivisionTZ16:
		return SubdivisionTypeRegion
	case SubdivisionTZ17:
		return SubdivisionTypeRegion
	case SubdivisionTZ18:
		return SubdivisionTypeRegion
	case SubdivisionTZ19:
		return SubdivisionTypeRegion
	case SubdivisionTZ20:
		return SubdivisionTypeRegion
	case SubdivisionTZ21:
		return SubdivisionTypeRegion
	case SubdivisionTZ22:
		return SubdivisionTypeRegion
	case SubdivisionTZ23:
		return SubdivisionTypeRegion
	case SubdivisionTZ24:
		return SubdivisionTypeRegion
	case SubdivisionTZ25:
		return SubdivisionTypeRegion
	case SubdivisionTZ26:
		return SubdivisionTypeRegion
	case SubdivisionUA05:
		return SubdivisionTypeProvince
	case SubdivisionUA07:
		return SubdivisionTypeProvince
	case SubdivisionUA09:
		return SubdivisionTypeProvince
	case SubdivisionUA12:
		return SubdivisionTypeProvince
	case SubdivisionUA14:
		return SubdivisionTypeProvince
	case SubdivisionUA18:
		return SubdivisionTypeProvince
	case SubdivisionUA21:
		return SubdivisionTypeProvince
	case SubdivisionUA23:
		return SubdivisionTypeProvince
	case SubdivisionUA26:
		return SubdivisionTypeProvince
	case SubdivisionUA30:
		return SubdivisionTypeCity
	case SubdivisionUA32:
		return SubdivisionTypeProvince
	case SubdivisionUA35:
		return SubdivisionTypeProvince
	case SubdivisionUA40:
		return SubdivisionTypeCity
	case SubdivisionUA43:
		return SubdivisionTypeAutonomousRepublic
	case SubdivisionUA46:
		return SubdivisionTypeProvince
	case SubdivisionUA48:
		return SubdivisionTypeProvince
	case SubdivisionUA51:
		return SubdivisionTypeProvince
	case SubdivisionUA53:
		return SubdivisionTypeProvince
	case SubdivisionUA56:
		return SubdivisionTypeProvince
	case SubdivisionUA59:
		return SubdivisionTypeProvince
	case SubdivisionUA61:
		return SubdivisionTypeProvince
	case SubdivisionUA63:
		return SubdivisionTypeProvince
	case SubdivisionUA65:
		return SubdivisionTypeProvince
	case SubdivisionUA68:
		return SubdivisionTypeProvince
	case SubdivisionUA71:
		return SubdivisionTypeProvince
	case SubdivisionUA74:
		return SubdivisionTypeProvince
	case SubdivisionUA77:
		return SubdivisionTypeProvince
	case SubdivisionUG101:
		return SubdivisionTypeDistrict
	case SubdivisionUG102:
		return SubdivisionTypeDistrict
	case SubdivisionUG103:
		return SubdivisionTypeDistrict
	case SubdivisionUG104:
		return SubdivisionTypeDistrict
	case SubdivisionUG105:
		return SubdivisionTypeDistrict
	case SubdivisionUG106:
		return SubdivisionTypeDistrict
	case SubdivisionUG107:
		return SubdivisionTypeDistrict
	case SubdivisionUG108:
		return SubdivisionTypeDistrict
	case SubdivisionUG109:
		return SubdivisionTypeDistrict
	case SubdivisionUG110:
		return SubdivisionTypeDistrict
	case SubdivisionUG111:
		return SubdivisionTypeDistrict
	case SubdivisionUG112:
		return SubdivisionTypeDistrict
	case SubdivisionUG113:
		return SubdivisionTypeDistrict
	case SubdivisionUG114:
		return SubdivisionTypeDistrict
	case SubdivisionUG115:
		return SubdivisionTypeDistrict
	case SubdivisionUG116:
		return SubdivisionTypeDistrict
	case SubdivisionUG201:
		return SubdivisionTypeDistrict
	case SubdivisionUG202:
		return SubdivisionTypeDistrict
	case SubdivisionUG203:
		return SubdivisionTypeDistrict
	case SubdivisionUG204:
		return SubdivisionTypeDistrict
	case SubdivisionUG205:
		return SubdivisionTypeDistrict
	case SubdivisionUG206:
		return SubdivisionTypeDistrict
	case SubdivisionUG207:
		return SubdivisionTypeDistrict
	case SubdivisionUG208:
		return SubdivisionTypeDistrict
	case SubdivisionUG209:
		return SubdivisionTypeDistrict
	case SubdivisionUG210:
		return SubdivisionTypeDistrict
	case SubdivisionUG211:
		return SubdivisionTypeDistrict
	case SubdivisionUG212:
		return SubdivisionTypeDistrict
	case SubdivisionUG213:
		return SubdivisionTypeDistrict
	case SubdivisionUG214:
		return SubdivisionTypeDistrict
	case SubdivisionUG215:
		return SubdivisionTypeDistrict
	case SubdivisionUG216:
		return SubdivisionTypeDistrict
	case SubdivisionUG217:
		return SubdivisionTypeDistrict
	case SubdivisionUG218:
		return SubdivisionTypeDistrict
	case SubdivisionUG219:
		return SubdivisionTypeDistrict
	case SubdivisionUG220:
		return SubdivisionTypeDistrict
	case SubdivisionUG221:
		return SubdivisionTypeDistrict
	case SubdivisionUG222:
		return SubdivisionTypeDistrict
	case SubdivisionUG223:
		return SubdivisionTypeDistrict
	case SubdivisionUG224:
		return SubdivisionTypeDistrict
	case SubdivisionUG301:
		return SubdivisionTypeDistrict
	case SubdivisionUG302:
		return SubdivisionTypeDistrict
	case SubdivisionUG303:
		return SubdivisionTypeDistrict
	case SubdivisionUG304:
		return SubdivisionTypeDistrict
	case SubdivisionUG305:
		return SubdivisionTypeDistrict
	case SubdivisionUG306:
		return SubdivisionTypeDistrict
	case SubdivisionUG307:
		return SubdivisionTypeDistrict
	case SubdivisionUG308:
		return SubdivisionTypeDistrict
	case SubdivisionUG309:
		return SubdivisionTypeDistrict
	case SubdivisionUG310:
		return SubdivisionTypeDistrict
	case SubdivisionUG311:
		return SubdivisionTypeDistrict
	case SubdivisionUG312:
		return SubdivisionTypeDistrict
	case SubdivisionUG313:
		return SubdivisionTypeDistrict
	case SubdivisionUG314:
		return SubdivisionTypeDistrict
	case SubdivisionUG315:
		return SubdivisionTypeDistrict
	case SubdivisionUG316:
		return SubdivisionTypeDistrict
	case SubdivisionUG317:
		return SubdivisionTypeDistrict
	case SubdivisionUG318:
		return SubdivisionTypeDistrict
	case SubdivisionUG319:
		return SubdivisionTypeDistrict
	case SubdivisionUG320:
		return SubdivisionTypeDistrict
	case SubdivisionUG321:
		return SubdivisionTypeDistrict
	case SubdivisionUG401:
		return SubdivisionTypeDistrict
	case SubdivisionUG402:
		return SubdivisionTypeDistrict
	case SubdivisionUG403:
		return SubdivisionTypeDistrict
	case SubdivisionUG404:
		return SubdivisionTypeDistrict
	case SubdivisionUG405:
		return SubdivisionTypeDistrict
	case SubdivisionUG406:
		return SubdivisionTypeDistrict
	case SubdivisionUG407:
		return SubdivisionTypeDistrict
	case SubdivisionUG408:
		return SubdivisionTypeDistrict
	case SubdivisionUG409:
		return SubdivisionTypeDistrict
	case SubdivisionUG410:
		return SubdivisionTypeDistrict
	case SubdivisionUG411:
		return SubdivisionTypeDistrict
	case SubdivisionUG412:
		return SubdivisionTypeDistrict
	case SubdivisionUG413:
		return SubdivisionTypeDistrict
	case SubdivisionUG414:
		return SubdivisionTypeDistrict
	case SubdivisionUG415:
		return SubdivisionTypeDistrict
	case SubdivisionUG416:
		return SubdivisionTypeDistrict
	case SubdivisionUG417:
		return SubdivisionTypeDistrict
	case SubdivisionUG418:
		return SubdivisionTypeDistrict
	case SubdivisionUG419:
		return SubdivisionTypeDistrict
	case SubdivisionUGC:
		return SubdivisionTypeGeographicalRegion
	case SubdivisionUGE:
		return SubdivisionTypeGeographicalRegion
	case SubdivisionUGN:
		return SubdivisionTypeGeographicalRegion
	case SubdivisionUGW:
		return SubdivisionTypeGeographicalRegion
	case SubdivisionUM67:
		return SubdivisionTypeTerritory
	case SubdivisionUM71:
		return SubdivisionTypeTerritory
	case SubdivisionUM76:
		return SubdivisionTypeTerritory
	case SubdivisionUM79:
		return SubdivisionTypeTerritory
	case SubdivisionUM81:
		return SubdivisionTypeTerritory
	case SubdivisionUM84:
		return SubdivisionTypeTerritory
	case SubdivisionUM86:
		return SubdivisionTypeTerritory
	case SubdivisionUM89:
		return SubdivisionTypeTerritory
	case SubdivisionUM95:
		return SubdivisionTypeTerritory
	case SubdivisionUSAK:
		return SubdivisionTypeState
	case SubdivisionUSAL:
		return SubdivisionTypeState
	case SubdivisionUSAR:
		return SubdivisionTypeState
	case SubdivisionUSAS:
		return SubdivisionTypeOutlyingArea
	case SubdivisionUSAZ:
		return SubdivisionTypeState
	case SubdivisionUSCA:
		return SubdivisionTypeState
	case SubdivisionUSCO:
		return SubdivisionTypeState
	case SubdivisionUSCT:
		return SubdivisionTypeState
	case SubdivisionUSDC:
		return SubdivisionTypeDistrict
	case SubdivisionUSDE:
		return SubdivisionTypeState
	case SubdivisionUSFL:
		return SubdivisionTypeState
	case SubdivisionUSGA:
		return SubdivisionTypeState
	case SubdivisionUSGU:
		return SubdivisionTypeOutlyingArea
	case SubdivisionUSHI:
		return SubdivisionTypeState
	case SubdivisionUSIA:
		return SubdivisionTypeState
	case SubdivisionUSID:
		return SubdivisionTypeState
	case SubdivisionUSIL:
		return SubdivisionTypeState
	case SubdivisionUSIN:
		return SubdivisionTypeState
	case SubdivisionUSKS:
		return SubdivisionTypeState
	case SubdivisionUSKY:
		return SubdivisionTypeState
	case SubdivisionUSLA:
		return SubdivisionTypeState
	case SubdivisionUSMA:
		return SubdivisionTypeState
	case SubdivisionUSMD:
		return SubdivisionTypeState
	case SubdivisionUSME:
		return SubdivisionTypeState
	case SubdivisionUSMI:
		return SubdivisionTypeState
	case SubdivisionUSMN:
		return SubdivisionTypeState
	case SubdivisionUSMO:
		return SubdivisionTypeState
	case SubdivisionUSMP:
		return SubdivisionTypeOutlyingArea
	case SubdivisionUSMS:
		return SubdivisionTypeState
	case SubdivisionUSMT:
		return SubdivisionTypeState
	case SubdivisionUSNC:
		return SubdivisionTypeState
	case SubdivisionUSND:
		return SubdivisionTypeState
	case SubdivisionUSNE:
		return SubdivisionTypeState
	case SubdivisionUSNH:
		return SubdivisionTypeState
	case SubdivisionUSNJ:
		return SubdivisionTypeState
	case SubdivisionUSNM:
		return SubdivisionTypeState
	case SubdivisionUSNV:
		return SubdivisionTypeState
	case SubdivisionUSNY:
		return SubdivisionTypeState
	case SubdivisionUSOH:
		return SubdivisionTypeState
	case SubdivisionUSOK:
		return SubdivisionTypeState
	case SubdivisionUSOR:
		return SubdivisionTypeState
	case SubdivisionUSPA:
		return SubdivisionTypeState
	case SubdivisionUSPR:
		return SubdivisionTypeOutlyingArea
	case SubdivisionUSRI:
		return SubdivisionTypeState
	case SubdivisionUSSC:
		return SubdivisionTypeState
	case SubdivisionUSSD:
		return SubdivisionTypeState
	case SubdivisionUSTN:
		return SubdivisionTypeState
	case SubdivisionUSTX:
		return SubdivisionTypeState
	case SubdivisionUSUM:
		return SubdivisionTypeOutlyingArea
	case SubdivisionUSUT:
		return SubdivisionTypeState
	case SubdivisionUSVA:
		return SubdivisionTypeState
	case SubdivisionUSVI:
		return SubdivisionTypeOutlyingArea
	case SubdivisionUSVT:
		return SubdivisionTypeState
	case SubdivisionUSWA:
		return SubdivisionTypeState
	case SubdivisionUSWI:
		return SubdivisionTypeState
	case SubdivisionUSWV:
		return SubdivisionTypeState
	case SubdivisionUSWY:
		return SubdivisionTypeState
	case SubdivisionUYAR:
		return SubdivisionTypeDepartment
	case SubdivisionUYCA:
		return SubdivisionTypeDepartment
	case SubdivisionUYCL:
		return SubdivisionTypeDepartment
	case SubdivisionUYCO:
		return SubdivisionTypeDepartment
	case SubdivisionUYDU:
		return SubdivisionTypeDepartment
	case SubdivisionUYFD:
		return SubdivisionTypeDepartment
	case SubdivisionUYFS:
		return SubdivisionTypeDepartment
	case SubdivisionUYLA:
		return SubdivisionTypeDepartment
	case SubdivisionUYMA:
		return SubdivisionTypeDepartment
	case SubdivisionUYMO:
		return SubdivisionTypeDepartment
	case SubdivisionUYPA:
		return SubdivisionTypeDepartment
	case SubdivisionUYRN:
		return SubdivisionTypeDepartment
	case SubdivisionUYRO:
		return SubdivisionTypeDepartment
	case SubdivisionUYRV:
		return SubdivisionTypeDepartment
	case SubdivisionUYSA:
		return SubdivisionTypeDepartment
	case SubdivisionUYSJ:
		return SubdivisionTypeDepartment
	case SubdivisionUYSO:
		return SubdivisionTypeDepartment
	case SubdivisionUYTA:
		return SubdivisionTypeDepartment
	case SubdivisionUYTT:
		return SubdivisionTypeDepartment
	case SubdivisionUZAN:
		return SubdivisionTypeRegion
	case SubdivisionUZBU:
		return SubdivisionTypeRegion
	case SubdivisionUZFA:
		return SubdivisionTypeRegion
	case SubdivisionUZJI:
		return SubdivisionTypeRegion
	case SubdivisionUZNG:
		return SubdivisionTypeRegion
	case SubdivisionUZNW:
		return SubdivisionTypeRegion
	case SubdivisionUZQA:
		return SubdivisionTypeRegion
	case SubdivisionUZQR:
		return SubdivisionTypeRepublic
	case SubdivisionUZSA:
		return SubdivisionTypeRegion
	case SubdivisionUZSI:
		return SubdivisionTypeRegion
	case SubdivisionUZSU:
		return SubdivisionTypeRegion
	case SubdivisionUZTK:
		return SubdivisionTypeCity
	case SubdivisionUZTO:
		return SubdivisionTypeRegion
	case SubdivisionUZXO:
		return SubdivisionTypeRegion
	case SubdivisionVC01:
		return SubdivisionTypeParish
	case SubdivisionVC02:
		return SubdivisionTypeParish
	case SubdivisionVC03:
		return SubdivisionTypeParish
	case SubdivisionVC04:
		return SubdivisionTypeParish
	case SubdivisionVC05:
		return SubdivisionTypeParish
	case SubdivisionVC06:
		return SubdivisionTypeParish
	case SubdivisionVEA:
		return SubdivisionTypeFederalDistrict
	case SubdivisionVEB:
		return SubdivisionTypeState
	case SubdivisionVEC:
		return SubdivisionTypeState
	case SubdivisionVED:
		return SubdivisionTypeState
	case SubdivisionVEE:
		return SubdivisionTypeState
	case SubdivisionVEF:
		return SubdivisionTypeState
	case SubdivisionVEG:
		return SubdivisionTypeState
	case SubdivisionVEH:
		return SubdivisionTypeState
	case SubdivisionVEI:
		return SubdivisionTypeState
	case SubdivisionVEJ:
		return SubdivisionTypeState
	case SubdivisionVEK:
		return SubdivisionTypeState
	case SubdivisionVEL:
		return SubdivisionTypeState
	case SubdivisionVEM:
		return SubdivisionTypeState
	case SubdivisionVEN:
		return SubdivisionTypeState
	case SubdivisionVEO:
		return SubdivisionTypeState
	case SubdivisionVEP:
		return SubdivisionTypeState
	case SubdivisionVER:
		return SubdivisionTypeState
	case SubdivisionVES:
		return SubdivisionTypeState
	case SubdivisionVET:
		return SubdivisionTypeState
	case SubdivisionVEU:
		return SubdivisionTypeState
	case SubdivisionVEV:
		return SubdivisionTypeState
	case SubdivisionVEW:
		return SubdivisionTypeFederalDependency
	case SubdivisionVEX:
		return SubdivisionTypeState
	case SubdivisionVEY:
		return SubdivisionTypeState
	case SubdivisionVEZ:
		return SubdivisionTypeState
	case SubdivisionVN01:
		return SubdivisionTypeProvince
	case SubdivisionVN02:
		return SubdivisionTypeProvince
	case SubdivisionVN03:
		return SubdivisionTypeProvince
	case SubdivisionVN04:
		return SubdivisionTypeProvince
	case SubdivisionVN05:
		return SubdivisionTypeProvince
	case SubdivisionVN06:
		return SubdivisionTypeProvince
	case SubdivisionVN07:
		return SubdivisionTypeProvince
	case SubdivisionVN09:
		return SubdivisionTypeProvince
	case SubdivisionVN13:
		return SubdivisionTypeProvince
	case SubdivisionVN14:
		return SubdivisionTypeProvince
	case SubdivisionVN15:
		return SubdivisionTypeProvince
	case SubdivisionVN18:
		return SubdivisionTypeProvince
	case SubdivisionVN20:
		return SubdivisionTypeProvince
	case SubdivisionVN21:
		return SubdivisionTypeProvince
	case SubdivisionVN22:
		return SubdivisionTypeProvince
	case SubdivisionVN23:
		return SubdivisionTypeProvince
	case SubdivisionVN24:
		return SubdivisionTypeProvince
	case SubdivisionVN25:
		return SubdivisionTypeProvince
	case SubdivisionVN26:
		return SubdivisionTypeProvince
	case SubdivisionVN27:
		return SubdivisionTypeProvince
	case SubdivisionVN28:
		return SubdivisionTypeProvince
	case SubdivisionVN29:
		return SubdivisionTypeProvince
	case SubdivisionVN30:
		return SubdivisionTypeProvince
	case SubdivisionVN31:
		return SubdivisionTypeProvince
	case SubdivisionVN32:
		return SubdivisionTypeProvince
	case SubdivisionVN33:
		return SubdivisionTypeProvince
	case SubdivisionVN34:
		return SubdivisionTypeProvince
	case SubdivisionVN35:
		return SubdivisionTypeProvince
	case SubdivisionVN36:
		return SubdivisionTypeProvince
	case SubdivisionVN37:
		return SubdivisionTypeProvince
	case SubdivisionVN39:
		return SubdivisionTypeProvince
	case SubdivisionVN40:
		return SubdivisionTypeProvince
	case SubdivisionVN41:
		return SubdivisionTypeProvince
	case SubdivisionVN43:
		return SubdivisionTypeProvince
	case SubdivisionVN44:
		return SubdivisionTypeProvince
	case SubdivisionVN45:
		return SubdivisionTypeProvince
	case SubdivisionVN46:
		return SubdivisionTypeProvince
	case SubdivisionVN47:
		return SubdivisionTypeProvince
	case SubdivisionVN49:
		return SubdivisionTypeProvince
	case SubdivisionVN50:
		return SubdivisionTypeProvince
	case SubdivisionVN51:
		return SubdivisionTypeProvince
	case SubdivisionVN52:
		return SubdivisionTypeProvince
	case SubdivisionVN53:
		return SubdivisionTypeProvince
	case SubdivisionVN54:
		return SubdivisionTypeProvince
	case SubdivisionVN55:
		return SubdivisionTypeProvince
	case SubdivisionVN56:
		return SubdivisionTypeProvince
	case SubdivisionVN57:
		return SubdivisionTypeProvince
	case SubdivisionVN58:
		return SubdivisionTypeProvince
	case SubdivisionVN59:
		return SubdivisionTypeProvince
	case SubdivisionVN61:
		return SubdivisionTypeProvince
	case SubdivisionVN63:
		return SubdivisionTypeProvince
	case SubdivisionVN66:
		return SubdivisionTypeProvince
	case SubdivisionVN67:
		return SubdivisionTypeProvince
	case SubdivisionVN68:
		return SubdivisionTypeProvince
	case SubdivisionVN69:
		return SubdivisionTypeProvince
	case SubdivisionVN70:
		return SubdivisionTypeProvince
	case SubdivisionVN71:
		return SubdivisionTypeProvince
	case SubdivisionVN72:
		return SubdivisionTypeProvince
	case SubdivisionVN73:
		return SubdivisionTypeProvince
	case SubdivisionVNCT:
		return SubdivisionTypeMunicipality
	case SubdivisionVNDN:
		return SubdivisionTypeMunicipality
	case SubdivisionVNHN:
		return SubdivisionTypeMunicipality
	case SubdivisionVNHP:
		return SubdivisionTypeMunicipality
	case SubdivisionVNSG:
		return SubdivisionTypeMunicipality
	case SubdivisionVUMAP:
		return SubdivisionTypeProvince
	case SubdivisionVUPAM:
		return SubdivisionTypeProvince
	case SubdivisionVUSAM:
		return SubdivisionTypeProvince
	case SubdivisionVUSEE:
		return SubdivisionTypeProvince
	case SubdivisionVUTAE:
		return SubdivisionTypeProvince
	case SubdivisionVUTOB:
		return SubdivisionTypeProvince
	case SubdivisionWSAA:
		return SubdivisionTypeDistrict
	case SubdivisionWSAL:
		return SubdivisionTypeDistrict
	case SubdivisionWSAT:
		return SubdivisionTypeDistrict
	case SubdivisionWSFA:
		return SubdivisionTypeDistrict
	case SubdivisionWSGE:
		return SubdivisionTypeDistrict
	case SubdivisionWSGI:
		return SubdivisionTypeDistrict
	case SubdivisionWSPA:
		return SubdivisionTypeDistrict
	case SubdivisionWSSA:
		return SubdivisionTypeDistrict
	case SubdivisionWSTU:
		return SubdivisionTypeDistrict
	case SubdivisionWSVF:
		return SubdivisionTypeDistrict
	case SubdivisionWSVS:
		return SubdivisionTypeDistrict
	case SubdivisionYEAB:
		return SubdivisionTypeGovernorate
	case SubdivisionYEAD:
		return SubdivisionTypeGovernorate
	case SubdivisionYEAM:
		return SubdivisionTypeGovernorate
	case SubdivisionYEBA:
		return SubdivisionTypeGovernorate
	case SubdivisionYEDA:
		return SubdivisionTypeGovernorate
	case SubdivisionYEDH:
		return SubdivisionTypeGovernorate
	case SubdivisionYEHD:
		return SubdivisionTypeGovernorate
	case SubdivisionYEHJ:
		return SubdivisionTypeGovernorate
	case SubdivisionYEIB:
		return SubdivisionTypeGovernorate
	case SubdivisionYEJA:
		return SubdivisionTypeGovernorate
	case SubdivisionYELA:
		return SubdivisionTypeGovernorate
	case SubdivisionYEMA:
		return SubdivisionTypeGovernorate
	case SubdivisionYEMR:
		return SubdivisionTypeGovernorate
	case SubdivisionYEMU:
		return SubdivisionTypeGovernorate
	case SubdivisionYEMW:
		return SubdivisionTypeGovernorate
	case SubdivisionYERA:
		return SubdivisionTypeGovernorate
	case SubdivisionYESD:
		return SubdivisionTypeGovernorate
	case SubdivisionYESH:
		return SubdivisionTypeGovernorate
	case SubdivisionYESN:
		return SubdivisionTypeGovernorate
	case SubdivisionYETA:
		return SubdivisionTypeGovernorate
	case SubdivisionZAEC:
		return SubdivisionTypeProvince
	case SubdivisionZAFS:
		return SubdivisionTypeProvince
	case SubdivisionZAGT:
		return SubdivisionTypeProvince
	case SubdivisionZALP:
		return SubdivisionTypeProvince
	case SubdivisionZAMP:
		return SubdivisionTypeProvince
	case SubdivisionZANC:
		return SubdivisionTypeProvince
	case SubdivisionZANL:
		return SubdivisionTypeProvince
	case SubdivisionZANW:
		return SubdivisionTypeProvince
	case SubdivisionZAWC:
		return SubdivisionTypeProvince
	case SubdivisionZM01:
		return SubdivisionTypeProvince
	case SubdivisionZM02:
		return SubdivisionTypeProvince
	case SubdivisionZM03:
		return SubdivisionTypeProvince
	case SubdivisionZM04:
		return SubdivisionTypeProvince
	case SubdivisionZM05:
		return SubdivisionTypeProvince
	case SubdivisionZM06:
		return SubdivisionTypeProvince
	case SubdivisionZM07:
		return SubdivisionTypeProvince
	case SubdivisionZM08:
		return SubdivisionTypeProvince
	case SubdivisionZM09:
		return SubdivisionTypeProvince
	case SubdivisionZWBU:
		return SubdivisionTypeCity
	case SubdivisionZWHA:
		return SubdivisionTypeCity
	case SubdivisionZWMA:
		return SubdivisionTypeProvince
	case SubdivisionZWMC:
		return SubdivisionTypeProvince
	case SubdivisionZWME:
		return SubdivisionTypeProvince
	case SubdivisionZWMI:
		return SubdivisionTypeProvince
	case SubdivisionZWMN:
		return SubdivisionTypeProvince
	case SubdivisionZWMS:
		return SubdivisionTypeProvince
	case SubdivisionZWMV:
		return SubdivisionTypeProvince
	case SubdivisionZWMW:
		return SubdivisionTypeProvince
	}
	return SubdivisionTypeUnknown
}

// Type implements Typer interface
func (_ Subdivision) Type() string {
	return TypeSubdivision
}

// Value implements database/sql/driver.Valuer
func (s Subdivision) Value() (Value, error) {
	return json.Marshal(s)
}

// Scan implements database/sql.Scanner
func (s *Subdivision) Scan(src interface{}) error {
	if s == nil {
		return fmt.Errorf("countries::Scan: Subdivision scan err: subdivision == nil")
	}
	switch src := src.(type) {
	case *Subdivision:
		*s = *src
	case Subdivision:
		*s = src
	default:
		return fmt.Errorf("countries::Scan: Subdivision scan err: unexpected value of type %T for %T", src, *s)
	}
	return nil
}

// AllSubdivisions - return all subdivision codes
//
//nolint:funlen
func AllSubdivisions() []SubdivisionCode {
	return []SubdivisionCode{
		SubdivisionUnknown,
		SubdivisionAD02,
		SubdivisionAD03,
		SubdivisionAD04,
		SubdivisionAD05,
		SubdivisionAD06,
		SubdivisionAD07,
		SubdivisionAD08,
		SubdivisionAEAJ,
		SubdivisionAEAZ,
		SubdivisionAEDU,
		SubdivisionAEFU,
		SubdivisionAERK,
		SubdivisionAESH,
		SubdivisionAEUQ,
		SubdivisionAFBAL,
		SubdivisionAFBAM,
		SubdivisionAFBDG,
		SubdivisionAFBDS,
		SubdivisionAFBGL,
		SubdivisionAFDAY,
		SubdivisionAFFRA,
		SubdivisionAFFYB,
		SubdivisionAFGHA,
		SubdivisionAFGHO,
		SubdivisionAFHEL,
		SubdivisionAFHER,
		SubdivisionAFJOW,
		SubdivisionAFKAB,
		SubdivisionAFKAN,
		SubdivisionAFKAP,
		SubdivisionAFKDZ,
		SubdivisionAFKHO,
		SubdivisionAFKNR,
		SubdivisionAFLAG,
		SubdivisionAFLOG,
		SubdivisionAFNAN,
		SubdivisionAFNIM,
		SubdivisionAFNUR,
		SubdivisionAFPAN,
		SubdivisionAFPAR,
		SubdivisionAFPIA,
		SubdivisionAFPKA,
		SubdivisionAFSAM,
		SubdivisionAFSAR,
		SubdivisionAFTAK,
		SubdivisionAFURU,
		SubdivisionAFWAR,
		SubdivisionAFZAB,
		SubdivisionAG03,
		SubdivisionAG04,
		SubdivisionAG05,
		SubdivisionAG06,
		SubdivisionAG07,
		SubdivisionAG08,
		SubdivisionAG10,
		SubdivisionAG11,
		SubdivisionAL01,
		SubdivisionAL02,
		SubdivisionAL03,
		SubdivisionAL04,
		SubdivisionAL05,
		SubdivisionAL06,
		SubdivisionAL07,
		SubdivisionAL08,
		SubdivisionAL09,
		SubdivisionAL10,
		SubdivisionAL11,
		SubdivisionAL12,
		SubdivisionALBR,
		SubdivisionALBU,
		SubdivisionALDI,
		SubdivisionALDL,
		SubdivisionALDR,
		SubdivisionALDV,
		SubdivisionALEL,
		SubdivisionALER,
		SubdivisionALFR,
		SubdivisionALGJ,
		SubdivisionALGR,
		SubdivisionALHA,
		SubdivisionALKA,
		SubdivisionALKB,
		SubdivisionALKC,
		SubdivisionALKO,
		SubdivisionALKR,
		SubdivisionALKU,
		SubdivisionALLB,
		SubdivisionALLE,
		SubdivisionALLU,
		SubdivisionALMK,
		SubdivisionALMM,
		SubdivisionALMR,
		SubdivisionALMT,
		SubdivisionALPG,
		SubdivisionALPQ,
		SubdivisionALPR,
		SubdivisionALPU,
		SubdivisionALSH,
		SubdivisionALSK,
		SubdivisionALSR,
		SubdivisionALTE,
		SubdivisionALTP,
		SubdivisionALTR,
		SubdivisionALVL,
		SubdivisionAMAG,
		SubdivisionAMAR,
		SubdivisionAMAV,
		SubdivisionAMER,
		SubdivisionAMGR,
		SubdivisionAMKT,
		SubdivisionAMLO,
		SubdivisionAMSH,
		SubdivisionAMSU,
		SubdivisionAMTV,
		SubdivisionAMVD,
		SubdivisionAOBGO,
		SubdivisionAOBGU,
		SubdivisionAOBIE,
		SubdivisionAOCAB,
		SubdivisionAOCCU,
		SubdivisionAOCNN,
		SubdivisionAOCNO,
		SubdivisionAOCUS,
		SubdivisionAOHUA,
		SubdivisionAOHUI,
		SubdivisionAOLNO,
		SubdivisionAOLSU,
		SubdivisionAOLUA,
		SubdivisionAOMAL,
		SubdivisionAOMOX,
		SubdivisionAONAM,
		SubdivisionAOUIG,
		SubdivisionAOZAI,
		SubdivisionARA,
		SubdivisionARB,
		SubdivisionARC,
		SubdivisionARD,
		SubdivisionARE,
		SubdivisionARG,
		SubdivisionARH,
		SubdivisionARJ,
		SubdivisionARK,
		SubdivisionARL,
		SubdivisionARM,
		SubdivisionARN,
		SubdivisionARP,
		SubdivisionARQ,
		SubdivisionARR,
		SubdivisionARS,
		SubdivisionART,
		SubdivisionARU,
		SubdivisionARV,
		SubdivisionARW,
		SubdivisionARX,
		SubdivisionARY,
		SubdivisionARZ,
		SubdivisionAT1,
		SubdivisionAT2,
		SubdivisionAT3,
		SubdivisionAT4,
		SubdivisionAT5,
		SubdivisionAT6,
		SubdivisionAT7,
		SubdivisionAT8,
		SubdivisionAT9,
		SubdivisionAUACT,
		SubdivisionAUNSW,
		SubdivisionAUNT,
		SubdivisionAUQLD,
		SubdivisionAUSA,
		SubdivisionAUTAS,
		SubdivisionAUVIC,
		SubdivisionAUWA,
		SubdivisionAZABS,
		SubdivisionAZAGA,
		SubdivisionAZAGC,
		SubdivisionAZAGM,
		SubdivisionAZAGS,
		SubdivisionAZAGU,
		SubdivisionAZAST,
		SubdivisionAZBA,
		SubdivisionAZBAB,
		SubdivisionAZBAL,
		SubdivisionAZBAR,
		SubdivisionAZBEY,
		SubdivisionAZBIL,
		SubdivisionAZCAB,
		SubdivisionAZCAL,
		SubdivisionAZCUL,
		SubdivisionAZDAS,
		SubdivisionAZFUZ,
		SubdivisionAZGA,
		SubdivisionAZGAD,
		SubdivisionAZGOR,
		SubdivisionAZGOY,
		SubdivisionAZGYG,
		SubdivisionAZHAC,
		SubdivisionAZIMI,
		SubdivisionAZISM,
		SubdivisionAZKAL,
		SubdivisionAZKAN,
		SubdivisionAZKUR,
		SubdivisionAZLA,
		SubdivisionAZLAC,
		SubdivisionAZLAN,
		SubdivisionAZLER,
		SubdivisionAZMAS,
		SubdivisionAZMI,
		SubdivisionAZNA,
		SubdivisionAZNEF,
		SubdivisionAZNV,
		SubdivisionAZNX,
		SubdivisionAZOGU,
		SubdivisionAZORD,
		SubdivisionAZQAB,
		SubdivisionAZQAX,
		SubdivisionAZQAZ,
		SubdivisionAZQBA,
		SubdivisionAZQBI,
		SubdivisionAZQOB,
		SubdivisionAZQUS,
		SubdivisionAZSA,
		SubdivisionAZSAB,
		SubdivisionAZSAD,
		SubdivisionAZSAH,
		SubdivisionAZSAK,
		SubdivisionAZSAL,
		SubdivisionAZSAR,
		SubdivisionAZSAT,
		SubdivisionAZSBN,
		SubdivisionAZSIY,
		SubdivisionAZSKR,
		SubdivisionAZSM,
		SubdivisionAZSMI,
		SubdivisionAZSMX,
		SubdivisionAZSR,
		SubdivisionAZSUS,
		SubdivisionAZTAR,
		SubdivisionAZTOV,
		SubdivisionAZUCA,
		SubdivisionAZXA,
		SubdivisionAZXAC,
		SubdivisionAZXCI,
		SubdivisionAZXIZ,
		SubdivisionAZXVD,
		SubdivisionAZYAR,
		SubdivisionAZYE,
		SubdivisionAZYEV,
		SubdivisionAZZAN,
		SubdivisionAZZAQ,
		SubdivisionAZZAR,
		SubdivisionBA01,
		SubdivisionBA02,
		SubdivisionBA03,
		SubdivisionBA04,
		SubdivisionBA05,
		SubdivisionBA06,
		SubdivisionBA07,
		SubdivisionBA08,
		SubdivisionBA09,
		SubdivisionBA10,
		SubdivisionBABIH,
		SubdivisionBABRC,
		SubdivisionBASRP,
		SubdivisionBB01,
		SubdivisionBB02,
		SubdivisionBB03,
		SubdivisionBB04,
		SubdivisionBB05,
		SubdivisionBB06,
		SubdivisionBB07,
		SubdivisionBB08,
		SubdivisionBB09,
		SubdivisionBB10,
		SubdivisionBB11,
		SubdivisionBD01,
		SubdivisionBD02,
		SubdivisionBD03,
		SubdivisionBD04,
		SubdivisionBD05,
		SubdivisionBD06,
		SubdivisionBD07,
		SubdivisionBD08,
		SubdivisionBD09,
		SubdivisionBD10,
		SubdivisionBD11,
		SubdivisionBD12,
		SubdivisionBD13,
		SubdivisionBD14,
		SubdivisionBD15,
		SubdivisionBD16,
		SubdivisionBD17,
		SubdivisionBD18,
		SubdivisionBD19,
		SubdivisionBD20,
		SubdivisionBD21,
		SubdivisionBD22,
		SubdivisionBD23,
		SubdivisionBD24,
		SubdivisionBD25,
		SubdivisionBD26,
		SubdivisionBD27,
		SubdivisionBD28,
		SubdivisionBD29,
		SubdivisionBD30,
		SubdivisionBD31,
		SubdivisionBD32,
		SubdivisionBD33,
		SubdivisionBD34,
		SubdivisionBD35,
		SubdivisionBD36,
		SubdivisionBD37,
		SubdivisionBD38,
		SubdivisionBD39,
		SubdivisionBD40,
		SubdivisionBD41,
		SubdivisionBD42,
		SubdivisionBD43,
		SubdivisionBD44,
		SubdivisionBD45,
		SubdivisionBD46,
		SubdivisionBD47,
		SubdivisionBD48,
		SubdivisionBD49,
		SubdivisionBD50,
		SubdivisionBD51,
		SubdivisionBD52,
		SubdivisionBD53,
		SubdivisionBD54,
		SubdivisionBD55,
		SubdivisionBD56,
		SubdivisionBD57,
		SubdivisionBD58,
		SubdivisionBD59,
		SubdivisionBD60,
		SubdivisionBD61,
		SubdivisionBD62,
		SubdivisionBD63,
		SubdivisionBD64,
		SubdivisionBDA,
		SubdivisionBDB,
		SubdivisionBDC,
		SubdivisionBDD,
		SubdivisionBDE,
		SubdivisionBDF,
		SubdivisionBDG,
		SubdivisionBDH,
		SubdivisionBEBRU,
		SubdivisionBEVAN,
		SubdivisionBEVBR,
		SubdivisionBEVLG,
		SubdivisionBEVLI,
		SubdivisionBEVOV,
		SubdivisionBEVWV,
		SubdivisionBEWAL,
		SubdivisionBEWBR,
		SubdivisionBEWHT,
		SubdivisionBEWLG,
		SubdivisionBEWLX,
		SubdivisionBEWNA,
		SubdivisionBF01,
		SubdivisionBF02,
		SubdivisionBF03,
		SubdivisionBF04,
		SubdivisionBF05,
		SubdivisionBF06,
		SubdivisionBF07,
		SubdivisionBF08,
		SubdivisionBF09,
		SubdivisionBF10,
		SubdivisionBF11,
		SubdivisionBF12,
		SubdivisionBF13,
		SubdivisionBFBAL,
		SubdivisionBFBAM,
		SubdivisionBFBAN,
		SubdivisionBFBAZ,
		SubdivisionBFBGR,
		SubdivisionBFBLG,
		SubdivisionBFBLK,
		SubdivisionBFCOM,
		SubdivisionBFGAN,
		SubdivisionBFGNA,
		SubdivisionBFGOU,
		SubdivisionBFHOU,
		SubdivisionBFIOB,
		SubdivisionBFKAD,
		SubdivisionBFKEN,
		SubdivisionBFKMD,
		SubdivisionBFKMP,
		SubdivisionBFKOP,
		SubdivisionBFKOS,
		SubdivisionBFKOT,
		SubdivisionBFKOW,
		SubdivisionBFLER,
		SubdivisionBFLOR,
		SubdivisionBFMOU,
		SubdivisionBFNAM,
		SubdivisionBFNAO,
		SubdivisionBFNAY,
		SubdivisionBFNOU,
		SubdivisionBFOUB,
		SubdivisionBFOUD,
		SubdivisionBFPAS,
		SubdivisionBFPON,
		SubdivisionBFSEN,
		SubdivisionBFSIS,
		SubdivisionBFSMT,
		SubdivisionBFSNG,
		SubdivisionBFSOM,
		SubdivisionBFSOR,
		SubdivisionBFTAP,
		SubdivisionBFTUI,
		SubdivisionBFYAG,
		SubdivisionBFYAT,
		SubdivisionBFZIR,
		SubdivisionBFZON,
		SubdivisionBFZOU,
		SubdivisionBG01,
		SubdivisionBG02,
		SubdivisionBG03,
		SubdivisionBG04,
		SubdivisionBG05,
		SubdivisionBG06,
		SubdivisionBG07,
		SubdivisionBG08,
		SubdivisionBG09,
		SubdivisionBG10,
		SubdivisionBG11,
		SubdivisionBG12,
		SubdivisionBG13,
		SubdivisionBG14,
		SubdivisionBG15,
		SubdivisionBG16,
		SubdivisionBG17,
		SubdivisionBG18,
		SubdivisionBG19,
		SubdivisionBG20,
		SubdivisionBG21,
		SubdivisionBG22,
		SubdivisionBG23,
		SubdivisionBG24,
		SubdivisionBG25,
		SubdivisionBG26,
		SubdivisionBG27,
		SubdivisionBG28,
		SubdivisionBH13,
		SubdivisionBH14,
		SubdivisionBH15,
		SubdivisionBH16,
		SubdivisionBH17,
		SubdivisionBIBB,
		SubdivisionBIBL,
		SubdivisionBIBM,
		SubdivisionBIBR,
		SubdivisionBICA,
		SubdivisionBICI,
		SubdivisionBIGI,
		SubdivisionBIKI,
		SubdivisionBIKR,
		SubdivisionBIKY,
		SubdivisionBIMA,
		SubdivisionBIMU,
		SubdivisionBIMW,
		SubdivisionBING,
		SubdivisionBIRT,
		SubdivisionBIRY,
		SubdivisionBJAK,
		SubdivisionBJAL,
		SubdivisionBJAQ,
		SubdivisionBJBO,
		SubdivisionBJCO,
		SubdivisionBJDO,
		SubdivisionBJKO,
		SubdivisionBJLI,
		SubdivisionBJMO,
		SubdivisionBJOU,
		SubdivisionBJPL,
		SubdivisionBJZO,
		SubdivisionBNBE,
		SubdivisionBNBM,
		SubdivisionBNTE,
		SubdivisionBNTU,
		SubdivisionBOB,
		SubdivisionBOC,
		SubdivisionBOH,
		SubdivisionBOL,
		SubdivisionBON,
		SubdivisionBOO,
		SubdivisionBOP,
		SubdivisionBOS,
		SubdivisionBOT,
		SubdivisionBQBO,
		SubdivisionBQSA,
		SubdivisionBQSE,
		SubdivisionBRAC,
		SubdivisionBRAL,
		SubdivisionBRAM,
		SubdivisionBRAP,
		SubdivisionBRBA,
		SubdivisionBRCE,
		SubdivisionBRDF,
		SubdivisionBRES,
		SubdivisionBRFN,
		SubdivisionBRGO,
		SubdivisionBRMA,
		SubdivisionBRMG,
		SubdivisionBRMS,
		SubdivisionBRMT,
		SubdivisionBRPA,
		SubdivisionBRPB,
		SubdivisionBRPE,
		SubdivisionBRPI,
		SubdivisionBRPR,
		SubdivisionBRRJ,
		SubdivisionBRRN,
		SubdivisionBRRO,
		SubdivisionBRRR,
		SubdivisionBRRS,
		SubdivisionBRSC,
		SubdivisionBRSE,
		SubdivisionBRSP,
		SubdivisionBRTO,
		SubdivisionBSAK,
		SubdivisionBSBI,
		SubdivisionBSBP,
		SubdivisionBSBY,
		SubdivisionBSCE,
		SubdivisionBSCI,
		SubdivisionBSCK,
		SubdivisionBSCO,
		SubdivisionBSCS,
		SubdivisionBSEG,
		SubdivisionBSEX,
		SubdivisionBSFP,
		SubdivisionBSGC,
		SubdivisionBSHI,
		SubdivisionBSHT,
		SubdivisionBSIN,
		SubdivisionBSLI,
		SubdivisionBSMC,
		SubdivisionBSMG,
		SubdivisionBSMI,
		SubdivisionBSNE,
		SubdivisionBSNO,
		SubdivisionBSNS,
		SubdivisionBSRC,
		SubdivisionBSRI,
		SubdivisionBSSA,
		SubdivisionBSSE,
		SubdivisionBSSO,
		SubdivisionBSSS,
		SubdivisionBSSW,
		SubdivisionBSWG,
		SubdivisionBT11,
		SubdivisionBT12,
		SubdivisionBT13,
		SubdivisionBT14,
		SubdivisionBT15,
		SubdivisionBT21,
		SubdivisionBT22,
		SubdivisionBT23,
		SubdivisionBT24,
		SubdivisionBT31,
		SubdivisionBT32,
		SubdivisionBT33,
		SubdivisionBT34,
		SubdivisionBT41,
		SubdivisionBT42,
		SubdivisionBT43,
		SubdivisionBT44,
		SubdivisionBT45,
		SubdivisionBTGA,
		SubdivisionBTTY,
		SubdivisionBWCE,
		SubdivisionBWGH,
		SubdivisionBWKG,
		SubdivisionBWKL,
		SubdivisionBWKW,
		SubdivisionBWNE,
		SubdivisionBWNW,
		SubdivisionBWSE,
		SubdivisionBWSO,
		SubdivisionBYBR,
		SubdivisionBYHM,
		SubdivisionBYHO,
		SubdivisionBYHR,
		SubdivisionBYMA,
		SubdivisionBYMI,
		SubdivisionBYVI,
		SubdivisionBZBZ,
		SubdivisionBZCY,
		SubdivisionBZCZL,
		SubdivisionBZOW,
		SubdivisionBZSC,
		SubdivisionBZTOL,
		SubdivisionCAAB,
		SubdivisionCABC,
		SubdivisionCAMB,
		SubdivisionCANB,
		SubdivisionCANL,
		SubdivisionCANS,
		SubdivisionCANT,
		SubdivisionCANU,
		SubdivisionCAON,
		SubdivisionCAPE,
		SubdivisionCAQC,
		SubdivisionCASK,
		SubdivisionCAYT,
		SubdivisionCDBC,
		SubdivisionCDBN,
		SubdivisionCDEQ,
		SubdivisionCDKA,
		SubdivisionCDKE,
		SubdivisionCDKN,
		SubdivisionCDKW,
		SubdivisionCDMA,
		SubdivisionCDNK,
		SubdivisionCDOR,
		SubdivisionCDSK,
		SubdivisionCFAC,
		SubdivisionCFBB,
		SubdivisionCFBGF,
		SubdivisionCFBK,
		SubdivisionCFHK,
		SubdivisionCFHM,
		SubdivisionCFHS,
		SubdivisionCFKB,
		SubdivisionCFKG,
		SubdivisionCFLB,
		SubdivisionCFMB,
		SubdivisionCFMP,
		SubdivisionCFNM,
		SubdivisionCFOP,
		SubdivisionCFSE,
		SubdivisionCFUK,
		SubdivisionCFVK,
		SubdivisionCG11,
		SubdivisionCG12,
		SubdivisionCG13,
		SubdivisionCG14,
		SubdivisionCG15,
		SubdivisionCG2,
		SubdivisionCG5,
		SubdivisionCG7,
		SubdivisionCG8,
		SubdivisionCG9,
		SubdivisionCGBZV,
		SubdivisionCHAG,
		SubdivisionCHAI,
		SubdivisionCHAR,
		SubdivisionCHBE,
		SubdivisionCHBL,
		SubdivisionCHBS,
		SubdivisionCHFR,
		SubdivisionCHGE,
		SubdivisionCHGL,
		SubdivisionCHGR,
		SubdivisionCHJU,
		SubdivisionCHLU,
		SubdivisionCHNE,
		SubdivisionCHNW,
		SubdivisionCHOW,
		SubdivisionCHSG,
		SubdivisionCHSH,
		SubdivisionCHSO,
		SubdivisionCHSZ,
		SubdivisionCHTG,
		SubdivisionCHTI,
		SubdivisionCHUR,
		SubdivisionCHVD,
		SubdivisionCHVS,
		SubdivisionCHZG,
		SubdivisionCHZH,
		SubdivisionCI01,
		SubdivisionCI02,
		SubdivisionCI03,
		SubdivisionCI04,
		SubdivisionCI05,
		SubdivisionCI06,
		SubdivisionCI07,
		SubdivisionCI08,
		SubdivisionCI09,
		SubdivisionCI10,
		SubdivisionCI11,
		SubdivisionCI12,
		SubdivisionCI13,
		SubdivisionCI14,
		SubdivisionCI15,
		SubdivisionCI16,
		SubdivisionCI17,
		SubdivisionCI18,
		SubdivisionCI19,
		SubdivisionCLAI,
		SubdivisionCLAN,
		SubdivisionCLAP,
		SubdivisionCLAR,
		SubdivisionCLAT,
		SubdivisionCLBI,
		SubdivisionCLCO,
		SubdivisionCLLI,
		SubdivisionCLLL,
		SubdivisionCLLR,
		SubdivisionCLMA,
		SubdivisionCLML,
		SubdivisionCLRM,
		SubdivisionCLTA,
		SubdivisionCLVS,
		SubdivisionCMAD,
		SubdivisionCMCE,
		SubdivisionCMEN,
		SubdivisionCMES,
		SubdivisionCMLT,
		SubdivisionCMNO,
		SubdivisionCMNW,
		SubdivisionCMOU,
		SubdivisionCMSU,
		SubdivisionCMSW,
		SubdivisionCNAH,
		SubdivisionCNBJ,
		SubdivisionCNCQ,
		SubdivisionCNFJ,
		SubdivisionCNGD,
		SubdivisionCNGS,
		SubdivisionCNGX,
		SubdivisionCNGZ,
		SubdivisionCNHA,
		SubdivisionCNHB,
		SubdivisionCNHE,
		SubdivisionCNHI,
		SubdivisionCNHK,
		SubdivisionCNHL,
		SubdivisionCNHN,
		SubdivisionCNJL,
		SubdivisionCNJS,
		SubdivisionCNJX,
		SubdivisionCNLN,
		SubdivisionCNMO,
		SubdivisionCNNM,
		SubdivisionCNNX,
		SubdivisionCNQH,
		SubdivisionCNSC,
		SubdivisionCNSD,
		SubdivisionCNSH,
		SubdivisionCNSN,
		SubdivisionCNSX,
		SubdivisionCNTJ,
		SubdivisionCNTW,
		SubdivisionCNXJ,
		SubdivisionCNXZ,
		SubdivisionCNYN,
		SubdivisionCNZJ,
		SubdivisionCOAMA,
		SubdivisionCOANT,
		SubdivisionCOARA,
		SubdivisionCOATL,
		SubdivisionCOBOL,
		SubdivisionCOBOY,
		SubdivisionCOCAL,
		SubdivisionCOCAQ,
		SubdivisionCOCAS,
		SubdivisionCOCAU,
		SubdivisionCOCES,
		SubdivisionCOCHO,
		SubdivisionCOCOR,
		SubdivisionCOCUN,
		SubdivisionCODC,
		SubdivisionCOGUA,
		SubdivisionCOGUV,
		SubdivisionCOHUI,
		SubdivisionCOLAG,
		SubdivisionCOMAG,
		SubdivisionCOMET,
		SubdivisionCONAR,
		SubdivisionCONSA,
		SubdivisionCOPUT,
		SubdivisionCOQUI,
		SubdivisionCORIS,
		SubdivisionCOSAN,
		SubdivisionCOSAP,
		SubdivisionCOSUC,
		SubdivisionCOTOL,
		SubdivisionCOVAC,
		SubdivisionCOVAU,
		SubdivisionCOVID,
		SubdivisionCRA,
		SubdivisionCRC,
		SubdivisionCRG,
		SubdivisionCRH,
		SubdivisionCRL,
		SubdivisionCRP,
		SubdivisionCRSJ,
		SubdivisionCU01,
		SubdivisionCU02,
		SubdivisionCU03,
		SubdivisionCU04,
		SubdivisionCU05,
		SubdivisionCU06,
		SubdivisionCU07,
		SubdivisionCU08,
		SubdivisionCU09,
		SubdivisionCU10,
		SubdivisionCU11,
		SubdivisionCU12,
		SubdivisionCU13,
		SubdivisionCU14,
		SubdivisionCU99,
		SubdivisionCVB,
		SubdivisionCVBR,
		SubdivisionCVBV,
		SubdivisionCVCA,
		SubdivisionCVCF,
		SubdivisionCVCR,
		SubdivisionCVMA,
		SubdivisionCVMO,
		SubdivisionCVPA,
		SubdivisionCVPN,
		SubdivisionCVPR,
		SubdivisionCVRB,
		SubdivisionCVRG,
		SubdivisionCVRS,
		SubdivisionCVS,
		SubdivisionCVSD,
		SubdivisionCVSF,
		SubdivisionCVSL,
		SubdivisionCVSM,
		SubdivisionCVSO,
		SubdivisionCVSS,
		SubdivisionCVSV,
		SubdivisionCVTA,
		SubdivisionCVTS,
		SubdivisionCY01,
		SubdivisionCY02,
		SubdivisionCY03,
		SubdivisionCY04,
		SubdivisionCY05,
		SubdivisionCY06,
		SubdivisionCZ10,
		SubdivisionCZ101,
		SubdivisionCZ102,
		SubdivisionCZ103,
		SubdivisionCZ104,
		SubdivisionCZ105,
		SubdivisionCZ106,
		SubdivisionCZ107,
		SubdivisionCZ108,
		SubdivisionCZ109,
		SubdivisionCZ110,
		SubdivisionCZ111,
		SubdivisionCZ112,
		SubdivisionCZ113,
		SubdivisionCZ114,
		SubdivisionCZ115,
		SubdivisionCZ116,
		SubdivisionCZ117,
		SubdivisionCZ118,
		SubdivisionCZ119,
		SubdivisionCZ120,
		SubdivisionCZ121,
		SubdivisionCZ122,
		SubdivisionCZ20,
		SubdivisionCZ201,
		SubdivisionCZ202,
		SubdivisionCZ203,
		SubdivisionCZ204,
		SubdivisionCZ205,
		SubdivisionCZ206,
		SubdivisionCZ207,
		SubdivisionCZ208,
		SubdivisionCZ209,
		SubdivisionCZ20A,
		SubdivisionCZ20B,
		SubdivisionCZ20C,
		SubdivisionCZ31,
		SubdivisionCZ311,
		SubdivisionCZ312,
		SubdivisionCZ313,
		SubdivisionCZ314,
		SubdivisionCZ315,
		SubdivisionCZ316,
		SubdivisionCZ317,
		SubdivisionCZ32,
		SubdivisionCZ321,
		SubdivisionCZ322,
		SubdivisionCZ323,
		SubdivisionCZ324,
		SubdivisionCZ325,
		SubdivisionCZ326,
		SubdivisionCZ327,
		SubdivisionCZ41,
		SubdivisionCZ411,
		SubdivisionCZ412,
		SubdivisionCZ413,
		SubdivisionCZ42,
		SubdivisionCZ421,
		SubdivisionCZ422,
		SubdivisionCZ423,
		SubdivisionCZ424,
		SubdivisionCZ425,
		SubdivisionCZ426,
		SubdivisionCZ427,
		SubdivisionCZ51,
		SubdivisionCZ511,
		SubdivisionCZ512,
		SubdivisionCZ513,
		SubdivisionCZ514,
		SubdivisionCZ52,
		SubdivisionCZ521,
		SubdivisionCZ522,
		SubdivisionCZ523,
		SubdivisionCZ524,
		SubdivisionCZ525,
		SubdivisionCZ53,
		SubdivisionCZ531,
		SubdivisionCZ532,
		SubdivisionCZ533,
		SubdivisionCZ534,
		SubdivisionCZ63,
		SubdivisionCZ631,
		SubdivisionCZ632,
		SubdivisionCZ633,
		SubdivisionCZ634,
		SubdivisionCZ635,
		SubdivisionCZ64,
		SubdivisionCZ641,
		SubdivisionCZ642,
		SubdivisionCZ643,
		SubdivisionCZ644,
		SubdivisionCZ645,
		SubdivisionCZ646,
		SubdivisionCZ647,
		SubdivisionCZ71,
		SubdivisionCZ711,
		SubdivisionCZ712,
		SubdivisionCZ713,
		SubdivisionCZ714,
		SubdivisionCZ715,
		SubdivisionCZ72,
		SubdivisionCZ721,
		SubdivisionCZ722,
		SubdivisionCZ723,
		SubdivisionCZ724,
		SubdivisionCZ80,
		SubdivisionCZ801,
		SubdivisionCZ802,
		SubdivisionCZ803,
		SubdivisionCZ804,
		SubdivisionCZ805,
		SubdivisionCZ806,
		SubdivisionDEBB,
		SubdivisionDEBE,
		SubdivisionDEBW,
		SubdivisionDEBY,
		SubdivisionDEHB,
		SubdivisionDEHE,
		SubdivisionDEHH,
		SubdivisionDEMV,
		SubdivisionDENI,
		SubdivisionDENW,
		SubdivisionDERP,
		SubdivisionDESH,
		SubdivisionDESL,
		SubdivisionDESN,
		SubdivisionDEST,
		SubdivisionDETH,
		SubdivisionDJAR,
		SubdivisionDJAS,
		SubdivisionDJDI,
		SubdivisionDJDJ,
		SubdivisionDJOB,
		SubdivisionDJTA,
		SubdivisionDK81,
		SubdivisionDK82,
		SubdivisionDK83,
		SubdivisionDK84,
		SubdivisionDK85,
		SubdivisionDM01,
		SubdivisionDM02,
		SubdivisionDM03,
		SubdivisionDM04,
		SubdivisionDM05,
		SubdivisionDM06,
		SubdivisionDM07,
		SubdivisionDM08,
		SubdivisionDM09,
		SubdivisionDM10,
		SubdivisionDO01,
		SubdivisionDO02,
		SubdivisionDO03,
		SubdivisionDO04,
		SubdivisionDO05,
		SubdivisionDO06,
		SubdivisionDO07,
		SubdivisionDO08,
		SubdivisionDO09,
		SubdivisionDO10,
		SubdivisionDO11,
		SubdivisionDO12,
		SubdivisionDO13,
		SubdivisionDO14,
		SubdivisionDO15,
		SubdivisionDO16,
		SubdivisionDO17,
		SubdivisionDO18,
		SubdivisionDO19,
		SubdivisionDO20,
		SubdivisionDO21,
		SubdivisionDO22,
		SubdivisionDO23,
		SubdivisionDO24,
		SubdivisionDO25,
		SubdivisionDO26,
		SubdivisionDO27,
		SubdivisionDO28,
		SubdivisionDO29,
		SubdivisionDO30,
		SubdivisionDZ01,
		SubdivisionDZ02,
		SubdivisionDZ03,
		SubdivisionDZ04,
		SubdivisionDZ05,
		SubdivisionDZ06,
		SubdivisionDZ07,
		SubdivisionDZ08,
		SubdivisionDZ09,
		SubdivisionDZ10,
		SubdivisionDZ11,
		SubdivisionDZ12,
		SubdivisionDZ13,
		SubdivisionDZ14,
		SubdivisionDZ15,
		SubdivisionDZ16,
		SubdivisionDZ17,
		SubdivisionDZ18,
		SubdivisionDZ19,
		SubdivisionDZ20,
		SubdivisionDZ21,
		SubdivisionDZ22,
		SubdivisionDZ23,
		SubdivisionDZ24,
		SubdivisionDZ25,
		SubdivisionDZ26,
		SubdivisionDZ27,
		SubdivisionDZ28,
		SubdivisionDZ29,
		SubdivisionDZ30,
		SubdivisionDZ31,
		SubdivisionDZ32,
		SubdivisionDZ33,
		SubdivisionDZ34,
		SubdivisionDZ35,
		SubdivisionDZ36,
		SubdivisionDZ37,
		SubdivisionDZ38,
		SubdivisionDZ39,
		SubdivisionDZ40,
		SubdivisionDZ41,
		SubdivisionDZ42,
		SubdivisionDZ43,
		SubdivisionDZ44,
		SubdivisionDZ45,
		SubdivisionDZ46,
		SubdivisionDZ47,
		SubdivisionDZ48,
		SubdivisionECA,
		SubdivisionECB,
		SubdivisionECC,
		SubdivisionECD,
		SubdivisionECE,
		SubdivisionECF,
		SubdivisionECG,
		SubdivisionECH,
		SubdivisionECI,
		SubdivisionECL,
		SubdivisionECM,
		SubdivisionECN,
		SubdivisionECO,
		SubdivisionECP,
		SubdivisionECR,
		SubdivisionECS,
		SubdivisionECSD,
		SubdivisionECSE,
		SubdivisionECT,
		SubdivisionECU,
		SubdivisionECW,
		SubdivisionECX,
		SubdivisionECY,
		SubdivisionECZ,
		SubdivisionEE37,
		SubdivisionEE39,
		SubdivisionEE44,
		SubdivisionEE49,
		SubdivisionEE51,
		SubdivisionEE57,
		SubdivisionEE59,
		SubdivisionEE65,
		SubdivisionEE67,
		SubdivisionEE70,
		SubdivisionEE74,
		SubdivisionEE78,
		SubdivisionEE82,
		SubdivisionEE84,
		SubdivisionEE86,
		SubdivisionEGALX,
		SubdivisionEGASN,
		SubdivisionEGAST,
		SubdivisionEGBA,
		SubdivisionEGBH,
		SubdivisionEGBNS,
		SubdivisionEGC,
		SubdivisionEGDK,
		SubdivisionEGDT,
		SubdivisionEGFYM,
		SubdivisionEGGH,
		SubdivisionEGGZ,
		SubdivisionEGHU,
		SubdivisionEGIS,
		SubdivisionEGJS,
		SubdivisionEGKB,
		SubdivisionEGKFS,
		SubdivisionEGKN,
		SubdivisionEGMN,
		SubdivisionEGMNF,
		SubdivisionEGMT,
		SubdivisionEGPTS,
		SubdivisionEGSHG,
		SubdivisionEGSHR,
		SubdivisionEGSIN,
		SubdivisionEGSU,
		SubdivisionEGSUZ,
		SubdivisionEGWAD,
		SubdivisionERAN,
		SubdivisionERDK,
		SubdivisionERDU,
		SubdivisionERGB,
		SubdivisionERMA,
		SubdivisionERSK,
		SubdivisionESA,
		SubdivisionESAB,
		SubdivisionESAL,
		SubdivisionESAN,
		SubdivisionESAR,
		SubdivisionESAS,
		SubdivisionESAV,
		SubdivisionESB,
		SubdivisionESBA,
		SubdivisionESBI,
		SubdivisionESBU,
		SubdivisionESC,
		SubdivisionESCA,
		SubdivisionESCB,
		SubdivisionESCC,
		SubdivisionESCE,
		SubdivisionESCL,
		SubdivisionESCM,
		SubdivisionESCN,
		SubdivisionESCO,
		SubdivisionESCR,
		SubdivisionESCS,
		SubdivisionESCT,
		SubdivisionESCU,
		SubdivisionESEX,
		SubdivisionESGA,
		SubdivisionESGC,
		SubdivisionESGI,
		SubdivisionESGR,
		SubdivisionESGU,
		SubdivisionESH,
		SubdivisionESHU,
		SubdivisionESIB,
		SubdivisionESJ,
		SubdivisionESL,
		SubdivisionESLE,
		SubdivisionESLO,
		SubdivisionESLU,
		SubdivisionESM,
		SubdivisionESMA,
		SubdivisionESMC,
		SubdivisionESMD,
		SubdivisionESML,
		SubdivisionESMU,
		SubdivisionESNA,
		SubdivisionESNC,
		SubdivisionESO,
		SubdivisionESOR,
		SubdivisionESP,
		SubdivisionESPM,
		SubdivisionESPO,
		SubdivisionESPV,
		SubdivisionESRI,
		SubdivisionESS,
		SubdivisionESSA,
		SubdivisionESSE,
		SubdivisionESSG,
		SubdivisionESSO,
		SubdivisionESSS,
		SubdivisionEST,
		SubdivisionESTE,
		SubdivisionESTF,
		SubdivisionESTO,
		SubdivisionESV,
		SubdivisionESVA,
		SubdivisionESVC,
		SubdivisionESVI,
		SubdivisionESZ,
		SubdivisionESZA,
		SubdivisionETAA,
		SubdivisionETAF,
		SubdivisionETAM,
		SubdivisionETBE,
		SubdivisionETDD,
		SubdivisionETGA,
		SubdivisionETHA,
		SubdivisionETOR,
		SubdivisionETSN,
		SubdivisionETSO,
		SubdivisionETTI,
		SubdivisionFI01,
		SubdivisionFI02,
		SubdivisionFI03,
		SubdivisionFI04,
		SubdivisionFI05,
		SubdivisionFI06,
		SubdivisionFI07,
		SubdivisionFI08,
		SubdivisionFI09,
		SubdivisionFI10,
		SubdivisionFI11,
		SubdivisionFI12,
		SubdivisionFI13,
		SubdivisionFI14,
		SubdivisionFI15,
		SubdivisionFI16,
		SubdivisionFI17,
		SubdivisionFI18,
		SubdivisionFI19,
		SubdivisionFJC,
		SubdivisionFJE,
		SubdivisionFJN,
		SubdivisionFJR,
		SubdivisionFJW,
		SubdivisionFMKSA,
		SubdivisionFMPNI,
		SubdivisionFMTRK,
		SubdivisionFMYAP,
		SubdivisionFR01,
		SubdivisionFR02,
		SubdivisionFR03,
		SubdivisionFR04,
		SubdivisionFR05,
		SubdivisionFR06,
		SubdivisionFR07,
		SubdivisionFR08,
		SubdivisionFR09,
		SubdivisionFR10,
		SubdivisionFR11,
		SubdivisionFR12,
		SubdivisionFR13,
		SubdivisionFR14,
		SubdivisionFR15,
		SubdivisionFR16,
		SubdivisionFR17,
		SubdivisionFR18,
		SubdivisionFR19,
		SubdivisionFR21,
		SubdivisionFR22,
		SubdivisionFR23,
		SubdivisionFR24,
		SubdivisionFR25,
		SubdivisionFR26,
		SubdivisionFR27,
		SubdivisionFR28,
		SubdivisionFR29,
		SubdivisionFR2A,
		SubdivisionFR2B,
		SubdivisionFR30,
		SubdivisionFR31,
		SubdivisionFR32,
		SubdivisionFR33,
		SubdivisionFR34,
		SubdivisionFR35,
		SubdivisionFR36,
		SubdivisionFR37,
		SubdivisionFR38,
		SubdivisionFR39,
		SubdivisionFR40,
		SubdivisionFR41,
		SubdivisionFR42,
		SubdivisionFR43,
		SubdivisionFR44,
		SubdivisionFR45,
		SubdivisionFR46,
		SubdivisionFR47,
		SubdivisionFR48,
		SubdivisionFR49,
		SubdivisionFR50,
		SubdivisionFR51,
		SubdivisionFR52,
		SubdivisionFR53,
		SubdivisionFR54,
		SubdivisionFR55,
		SubdivisionFR56,
		SubdivisionFR57,
		SubdivisionFR58,
		SubdivisionFR59,
		SubdivisionFR60,
		SubdivisionFR61,
		SubdivisionFR62,
		SubdivisionFR63,
		SubdivisionFR64,
		SubdivisionFR65,
		SubdivisionFR66,
		SubdivisionFR67,
		SubdivisionFR68,
		SubdivisionFR69,
		SubdivisionFR70,
		SubdivisionFR71,
		SubdivisionFR72,
		SubdivisionFR73,
		SubdivisionFR74,
		SubdivisionFR75,
		SubdivisionFR76,
		SubdivisionFR77,
		SubdivisionFR78,
		SubdivisionFR79,
		SubdivisionFR80,
		SubdivisionFR81,
		SubdivisionFR82,
		SubdivisionFR83,
		SubdivisionFR84,
		SubdivisionFR85,
		SubdivisionFR86,
		SubdivisionFR87,
		SubdivisionFR88,
		SubdivisionFR89,
		SubdivisionFR90,
		SubdivisionFR91,
		SubdivisionFR92,
		SubdivisionFR93,
		SubdivisionFR94,
		SubdivisionFR95,
		SubdivisionFRARA,
		SubdivisionFRBFC,
		SubdivisionFRBL,
		SubdivisionFRBRE,
		SubdivisionFRCOR,
		SubdivisionFRCP,
		SubdivisionFRCVL,
		SubdivisionFRGES,
		SubdivisionFRGF,
		SubdivisionFRGP,
		SubdivisionFRGUA,
		SubdivisionFRHDF,
		SubdivisionFRIDF,
		SubdivisionFRLRE,
		SubdivisionFRMAY,
		SubdivisionFRMF,
		SubdivisionFRMQ,
		SubdivisionFRNAQ,
		SubdivisionFRNC,
		SubdivisionFRNOR,
		SubdivisionFROCC,
		SubdivisionFRPAC,
		SubdivisionFRPDL,
		SubdivisionFRPF,
		SubdivisionFRPM,
		SubdivisionFRRE,
		SubdivisionFRTF,
		SubdivisionFRWF,
		SubdivisionFRYT,
		SubdivisionGA1,
		SubdivisionGA2,
		SubdivisionGA3,
		SubdivisionGA4,
		SubdivisionGA5,
		SubdivisionGA6,
		SubdivisionGA7,
		SubdivisionGA8,
		SubdivisionGA9,
		SubdivisionGBABC,
		SubdivisionGBABD,
		SubdivisionGBABE,
		SubdivisionGBAGB,
		SubdivisionGBAGY,
		SubdivisionGBAND,
		SubdivisionGBANN,
		SubdivisionGBANS,
		SubdivisionGBBAS,
		SubdivisionGBBBD,
		SubdivisionGBBDF,
		SubdivisionGBBDG,
		SubdivisionGBBEN,
		SubdivisionGBBEX,
		SubdivisionGBBFS,
		SubdivisionGBBGE,
		SubdivisionGBBGW,
		SubdivisionGBBIR,
		SubdivisionGBBKM,
		SubdivisionGBBMH,
		SubdivisionGBBNE,
		SubdivisionGBBNH,
		SubdivisionGBBNS,
		SubdivisionGBBOL,
		SubdivisionGBBPL,
		SubdivisionGBBRC,
		SubdivisionGBBRD,
		SubdivisionGBBRY,
		SubdivisionGBBST,
		SubdivisionGBBUR,
		SubdivisionGBCAM,
		SubdivisionGBCAY,
		SubdivisionGBCBF,
		SubdivisionGBCCG,
		SubdivisionGBCGN,
		SubdivisionGBCHE,
		SubdivisionGBCHW,
		SubdivisionGBCLD,
		SubdivisionGBCLK,
		SubdivisionGBCMA,
		SubdivisionGBCMD,
		SubdivisionGBCMN,
		SubdivisionGBCON,
		SubdivisionGBCOV,
		SubdivisionGBCRF,
		SubdivisionGBCRY,
		SubdivisionGBCWY,
		SubdivisionGBDAL,
		SubdivisionGBDBY,
		SubdivisionGBDEN,
		SubdivisionGBDER,
		SubdivisionGBDEV,
		SubdivisionGBDGY,
		SubdivisionGBDNC,
		SubdivisionGBDND,
		SubdivisionGBDOR,
		SubdivisionGBDRS,
		SubdivisionGBDUD,
		SubdivisionGBDUR,
		SubdivisionGBEAL,
		SubdivisionGBEAW,
		SubdivisionGBEAY,
		SubdivisionGBEDH,
		SubdivisionGBEDU,
		SubdivisionGBELN,
		SubdivisionGBELS,
		SubdivisionGBENF,
		SubdivisionGBENG,
		SubdivisionGBERW,
		SubdivisionGBERY,
		SubdivisionGBESS,
		SubdivisionGBESX,
		SubdivisionGBFAL,
		SubdivisionGBFIF,
		SubdivisionGBFLN,
		SubdivisionGBFMO,
		SubdivisionGBGAT,
		SubdivisionGBGBN,
		SubdivisionGBGLG,
		SubdivisionGBGLS,
		SubdivisionGBGRE,
		SubdivisionGBGWN,
		SubdivisionGBHAL,
		SubdivisionGBHAM,
		SubdivisionGBHAV,
		SubdivisionGBHCK,
		SubdivisionGBHEF,
		SubdivisionGBHIL,
		SubdivisionGBHLD,
		SubdivisionGBHMF,
		SubdivisionGBHNS,
		SubdivisionGBHPL,
		SubdivisionGBHRT,
		SubdivisionGBHRW,
		SubdivisionGBHRY,
		SubdivisionGBIOS,
		SubdivisionGBIOW,
		SubdivisionGBISL,
		SubdivisionGBIVC,
		SubdivisionGBKEC,
		SubdivisionGBKEN,
		SubdivisionGBKHL,
		SubdivisionGBKIR,
		SubdivisionGBKTT,
		SubdivisionGBKWL,
		SubdivisionGBLAN,
		SubdivisionGBLBC,
		SubdivisionGBLBH,
		SubdivisionGBLCE,
		SubdivisionGBLDS,
		SubdivisionGBLEC,
		SubdivisionGBLEW,
		SubdivisionGBLIN,
		SubdivisionGBLIV,
		SubdivisionGBLND,
		SubdivisionGBLUT,
		SubdivisionGBMAN,
		SubdivisionGBMDB,
		SubdivisionGBMDW,
		SubdivisionGBMEA,
		SubdivisionGBMIK,
		SubdivisionGBMLN,
		SubdivisionGBMON,
		SubdivisionGBMRT,
		SubdivisionGBMRY,
		SubdivisionGBMTY,
		SubdivisionGBMUL,
		SubdivisionGBNAY,
		SubdivisionGBNBL,
		SubdivisionGBNEL,
		SubdivisionGBNET,
		SubdivisionGBNFK,
		SubdivisionGBNGM,
		SubdivisionGBNIR,
		SubdivisionGBNLK,
		SubdivisionGBNLN,
		SubdivisionGBNMD,
		SubdivisionGBNSM,
		SubdivisionGBNTH,
		SubdivisionGBNTL,
		SubdivisionGBNTT,
		SubdivisionGBNTY,
		SubdivisionGBNWM,
		SubdivisionGBNWP,
		SubdivisionGBNYK,
		SubdivisionGBOLD,
		SubdivisionGBORK,
		SubdivisionGBOXF,
		SubdivisionGBPEM,
		SubdivisionGBPKN,
		SubdivisionGBPLY,
		SubdivisionGBPOL,
		SubdivisionGBPOR,
		SubdivisionGBPOW,
		SubdivisionGBPTE,
		SubdivisionGBRCC,
		SubdivisionGBRCH,
		SubdivisionGBRCT,
		SubdivisionGBRDB,
		SubdivisionGBRDG,
		SubdivisionGBRFW,
		SubdivisionGBRIC,
		SubdivisionGBROT,
		SubdivisionGBRUT,
		SubdivisionGBSAW,
		SubdivisionGBSAY,
		SubdivisionGBSCB,
		SubdivisionGBSCT,
		SubdivisionGBSFK,
		SubdivisionGBSFT,
		SubdivisionGBSGC,
		SubdivisionGBSHF,
		SubdivisionGBSHN,
		SubdivisionGBSHR,
		SubdivisionGBSKP,
		SubdivisionGBSLF,
		SubdivisionGBSLG,
		SubdivisionGBSLK,
		SubdivisionGBSND,
		SubdivisionGBSOL,
		SubdivisionGBSOM,
		SubdivisionGBSOS,
		SubdivisionGBSRY,
		SubdivisionGBSTE,
		SubdivisionGBSTG,
		SubdivisionGBSTH,
		SubdivisionGBSTN,
		SubdivisionGBSTS,
		SubdivisionGBSTT,
		SubdivisionGBSTY,
		SubdivisionGBSWA,
		SubdivisionGBSWD,
		SubdivisionGBSWK,
		SubdivisionGBTAM,
		SubdivisionGBTFW,
		SubdivisionGBTHR,
		SubdivisionGBTOB,
		SubdivisionGBTOF,
		SubdivisionGBTRF,
		SubdivisionGBTWH,
		SubdivisionGBUKM,
		SubdivisionGBVGL,
		SubdivisionGBWAR,
		SubdivisionGBWBK,
		SubdivisionGBWDU,
		SubdivisionGBWFT,
		SubdivisionGBWGN,
		SubdivisionGBWIL,
		SubdivisionGBWKF,
		SubdivisionGBWLL,
		SubdivisionGBWLN,
		SubdivisionGBWLS,
		SubdivisionGBWLV,
		SubdivisionGBWND,
		SubdivisionGBWNM,
		SubdivisionGBWOK,
		SubdivisionGBWOR,
		SubdivisionGBWRL,
		SubdivisionGBWRT,
		SubdivisionGBWRX,
		SubdivisionGBWSM,
		SubdivisionGBWSX,
		SubdivisionGBYOR,
		SubdivisionGBZET,
		SubdivisionGD01,
		SubdivisionGD02,
		SubdivisionGD03,
		SubdivisionGD04,
		SubdivisionGD05,
		SubdivisionGD06,
		SubdivisionGD10,
		SubdivisionGEAB,
		SubdivisionGEAJ,
		SubdivisionGEGU,
		SubdivisionGEIM,
		SubdivisionGEKA,
		SubdivisionGEKK,
		SubdivisionGEMM,
		SubdivisionGERL,
		SubdivisionGESJ,
		SubdivisionGESK,
		SubdivisionGESZ,
		SubdivisionGETB,
		SubdivisionGHAA,
		SubdivisionGHAH,
		SubdivisionGHBA,
		SubdivisionGHCP,
		SubdivisionGHEP,
		SubdivisionGHNP,
		SubdivisionGHTV,
		SubdivisionGHUE,
		SubdivisionGHUW,
		SubdivisionGHWP,
		SubdivisionGLKU,
		SubdivisionGLQA,
		SubdivisionGLQE,
		SubdivisionGLSM,
		SubdivisionGMB,
		SubdivisionGML,
		SubdivisionGMM,
		SubdivisionGMN,
		SubdivisionGMU,
		SubdivisionGMW,
		SubdivisionGNB,
		SubdivisionGNBE,
		SubdivisionGNBF,
		SubdivisionGNBK,
		SubdivisionGNC,
		SubdivisionGNCO,
		SubdivisionGND,
		SubdivisionGNDB,
		SubdivisionGNDI,
		SubdivisionGNDL,
		SubdivisionGNDU,
		SubdivisionGNF,
		SubdivisionGNFA,
		SubdivisionGNFO,
		SubdivisionGNFR,
		SubdivisionGNGA,
		SubdivisionGNGU,
		SubdivisionGNK,
		SubdivisionGNKA,
		SubdivisionGNKB,
		SubdivisionGNKD,
		SubdivisionGNKE,
		SubdivisionGNKN,
		SubdivisionGNKO,
		SubdivisionGNKS,
		SubdivisionGNL,
		SubdivisionGNLA,
		SubdivisionGNLE,
		SubdivisionGNLO,
		SubdivisionGNM,
		SubdivisionGNMC,
		SubdivisionGNMD,
		SubdivisionGNML,
		SubdivisionGNMM,
		SubdivisionGNN,
		SubdivisionGNNZ,
		SubdivisionGNPI,
		SubdivisionGNSI,
		SubdivisionGNTE,
		SubdivisionGNTO,
		SubdivisionGNYO,
		SubdivisionGQAN,
		SubdivisionGQBN,
		SubdivisionGQBS,
		SubdivisionGQC,
		SubdivisionGQCS,
		SubdivisionGQI,
		SubdivisionGQKN,
		SubdivisionGQLI,
		SubdivisionGQWN,
		SubdivisionGR01,
		SubdivisionGR03,
		SubdivisionGR04,
		SubdivisionGR05,
		SubdivisionGR06,
		SubdivisionGR07,
		SubdivisionGR11,
		SubdivisionGR12,
		SubdivisionGR13,
		SubdivisionGR14,
		SubdivisionGR15,
		SubdivisionGR16,
		SubdivisionGR17,
		SubdivisionGR21,
		SubdivisionGR22,
		SubdivisionGR23,
		SubdivisionGR24,
		SubdivisionGR31,
		SubdivisionGR32,
		SubdivisionGR33,
		SubdivisionGR34,
		SubdivisionGR41,
		SubdivisionGR42,
		SubdivisionGR43,
		SubdivisionGR44,
		SubdivisionGR51,
		SubdivisionGR52,
		SubdivisionGR53,
		SubdivisionGR54,
		SubdivisionGR55,
		SubdivisionGR56,
		SubdivisionGR57,
		SubdivisionGR58,
		SubdivisionGR59,
		SubdivisionGR61,
		SubdivisionGR62,
		SubdivisionGR63,
		SubdivisionGR64,
		SubdivisionGR69,
		SubdivisionGR71,
		SubdivisionGR72,
		SubdivisionGR73,
		SubdivisionGR81,
		SubdivisionGR82,
		SubdivisionGR83,
		SubdivisionGR84,
		SubdivisionGR85,
		SubdivisionGR91,
		SubdivisionGR92,
		SubdivisionGR93,
		SubdivisionGR94,
		SubdivisionGRA,
		SubdivisionGRA1,
		SubdivisionGRB,
		SubdivisionGRC,
		SubdivisionGRD,
		SubdivisionGRE,
		SubdivisionGRF,
		SubdivisionGRG,
		SubdivisionGRH,
		SubdivisionGRI,
		SubdivisionGRJ,
		SubdivisionGRK,
		SubdivisionGRL,
		SubdivisionGRM,
		SubdivisionGTAV,
		SubdivisionGTBV,
		SubdivisionGTCM,
		SubdivisionGTCQ,
		SubdivisionGTES,
		SubdivisionGTGU,
		SubdivisionGTHU,
		SubdivisionGTIZ,
		SubdivisionGTJA,
		SubdivisionGTJU,
		SubdivisionGTPE,
		SubdivisionGTPR,
		SubdivisionGTQC,
		SubdivisionGTQZ,
		SubdivisionGTRE,
		SubdivisionGTSA,
		SubdivisionGTSM,
		SubdivisionGTSO,
		SubdivisionGTSR,
		SubdivisionGTSU,
		SubdivisionGTTO,
		SubdivisionGTZA,
		SubdivisionGWBA,
		SubdivisionGWBL,
		SubdivisionGWBM,
		SubdivisionGWBS,
		SubdivisionGWCA,
		SubdivisionGWGA,
		SubdivisionGWL,
		SubdivisionGWN,
		SubdivisionGWOI,
		SubdivisionGWQU,
		SubdivisionGWS,
		SubdivisionGWTO,
		SubdivisionGYBA,
		SubdivisionGYCU,
		SubdivisionGYDE,
		SubdivisionGYEB,
		SubdivisionGYES,
		SubdivisionGYMA,
		SubdivisionGYPM,
		SubdivisionGYPT,
		SubdivisionGYUD,
		SubdivisionGYUT,
		SubdivisionHNAT,
		SubdivisionHNCH,
		SubdivisionHNCL,
		SubdivisionHNCM,
		SubdivisionHNCP,
		SubdivisionHNCR,
		SubdivisionHNEP,
		SubdivisionHNFM,
		SubdivisionHNGD,
		SubdivisionHNIB,
		SubdivisionHNIN,
		SubdivisionHNLE,
		SubdivisionHNLP,
		SubdivisionHNOC,
		SubdivisionHNOL,
		SubdivisionHNSB,
		SubdivisionHNVA,
		SubdivisionHNYO,
		SubdivisionHR01,
		SubdivisionHR02,
		SubdivisionHR03,
		SubdivisionHR04,
		SubdivisionHR05,
		SubdivisionHR06,
		SubdivisionHR07,
		SubdivisionHR08,
		SubdivisionHR09,
		SubdivisionHR10,
		SubdivisionHR11,
		SubdivisionHR12,
		SubdivisionHR13,
		SubdivisionHR14,
		SubdivisionHR15,
		SubdivisionHR16,
		SubdivisionHR17,
		SubdivisionHR18,
		SubdivisionHR19,
		SubdivisionHR20,
		SubdivisionHR21,
		SubdivisionHTAR,
		SubdivisionHTCE,
		SubdivisionHTGA,
		SubdivisionHTND,
		SubdivisionHTNE,
		SubdivisionHTNO,
		SubdivisionHTOU,
		SubdivisionHTSD,
		SubdivisionHTSE,
		SubdivisionHUBA,
		SubdivisionHUBC,
		SubdivisionHUBE,
		SubdivisionHUBK,
		SubdivisionHUBU,
		SubdivisionHUBZ,
		SubdivisionHUCS,
		SubdivisionHUDE,
		SubdivisionHUDU,
		SubdivisionHUEG,
		SubdivisionHUER,
		SubdivisionHUFE,
		SubdivisionHUGS,
		SubdivisionHUGY,
		SubdivisionHUHB,
		SubdivisionHUHE,
		SubdivisionHUHV,
		SubdivisionHUJN,
		SubdivisionHUKE,
		SubdivisionHUKM,
		SubdivisionHUKV,
		SubdivisionHUMI,
		SubdivisionHUNK,
		SubdivisionHUNO,
		SubdivisionHUNY,
		SubdivisionHUPE,
		SubdivisionHUPS,
		SubdivisionHUSD,
		SubdivisionHUSF,
		SubdivisionHUSH,
		SubdivisionHUSK,
		SubdivisionHUSN,
		SubdivisionHUSO,
		SubdivisionHUSS,
		SubdivisionHUST,
		SubdivisionHUSZ,
		SubdivisionHUTB,
		SubdivisionHUTO,
		SubdivisionHUVA,
		SubdivisionHUVE,
		SubdivisionHUVM,
		SubdivisionHUZA,
		SubdivisionHUZE,
		SubdivisionIDAC,
		SubdivisionIDBA,
		SubdivisionIDBB,
		SubdivisionIDBE,
		SubdivisionIDBT,
		SubdivisionIDGO,
		SubdivisionIDIJ,
		SubdivisionIDJA,
		SubdivisionIDJB,
		SubdivisionIDJI,
		SubdivisionIDJK,
		SubdivisionIDJT,
		SubdivisionIDJW,
		SubdivisionIDKA,
		SubdivisionIDKB,
		SubdivisionIDKI,
		SubdivisionIDKR,
		SubdivisionIDKS,
		SubdivisionIDKT,
		SubdivisionIDLA,
		SubdivisionIDMA,
		SubdivisionIDML,
		SubdivisionIDMU,
		SubdivisionIDNB,
		SubdivisionIDNT,
		SubdivisionIDNU,
		SubdivisionIDPA,
		SubdivisionIDPB,
		SubdivisionIDRI,
		SubdivisionIDSA,
		SubdivisionIDSB,
		SubdivisionIDSG,
		SubdivisionIDSL,
		SubdivisionIDSM,
		SubdivisionIDSN,
		SubdivisionIDSR,
		SubdivisionIDSS,
		SubdivisionIDST,
		SubdivisionIDSU,
		SubdivisionIDYO,
		SubdivisionIEC,
		SubdivisionIECE,
		SubdivisionIECN,
		SubdivisionIECO,
		SubdivisionIECW,
		SubdivisionIED,
		SubdivisionIEDL,
		SubdivisionIEG,
		SubdivisionIEKE,
		SubdivisionIEKK,
		SubdivisionIEKY,
		SubdivisionIEL,
		SubdivisionIELD,
		SubdivisionIELH,
		SubdivisionIELK,
		SubdivisionIELM,
		SubdivisionIELS,
		SubdivisionIEM,
		SubdivisionIEMH,
		SubdivisionIEMN,
		SubdivisionIEMO,
		SubdivisionIEOY,
		SubdivisionIERN,
		SubdivisionIESO,
		SubdivisionIETA,
		SubdivisionIEU,
		SubdivisionIEWD,
		SubdivisionIEWH,
		SubdivisionIEWW,
		SubdivisionIEWX,
		SubdivisionILD,
		SubdivisionILHA,
		SubdivisionILJM,
		SubdivisionILM,
		SubdivisionILTA,
		SubdivisionILZ,
		SubdivisionINAN,
		SubdivisionINAP,
		SubdivisionINAR,
		SubdivisionINAS,
		SubdivisionINBR,
		SubdivisionINCH,
		SubdivisionINCT,
		SubdivisionINDD,
		SubdivisionINDL,
		SubdivisionINDN,
		SubdivisionINGA,
		SubdivisionINGJ,
		SubdivisionINHP,
		SubdivisionINHR,
		SubdivisionINJH,
		SubdivisionINJK,
		SubdivisionINKA,
		SubdivisionINKL,
		SubdivisionINLD,
		SubdivisionINMH,
		SubdivisionINML,
		SubdivisionINMN,
		SubdivisionINMP,
		SubdivisionINMZ,
		SubdivisionINNL,
		SubdivisionINOR,
		SubdivisionINPB,
		SubdivisionINPY,
		SubdivisionINRJ,
		SubdivisionINSK,
		SubdivisionINTG,
		SubdivisionINTN,
		SubdivisionINTR,
		SubdivisionINUP,
		SubdivisionINUT,
		SubdivisionINWB,
		SubdivisionIQAN,
		SubdivisionIQAR,
		SubdivisionIQBA,
		SubdivisionIQBB,
		SubdivisionIQBG,
		SubdivisionIQDA,
		SubdivisionIQDI,
		SubdivisionIQDQ,
		SubdivisionIQKA,
		SubdivisionIQMA,
		SubdivisionIQMU,
		SubdivisionIQNA,
		SubdivisionIQNI,
		SubdivisionIQQA,
		SubdivisionIQSD,
		SubdivisionIQSW,
		SubdivisionIQTS,
		SubdivisionIQWA,
		SubdivisionIR01,
		SubdivisionIR02,
		SubdivisionIR03,
		SubdivisionIR04,
		SubdivisionIR05,
		SubdivisionIR06,
		SubdivisionIR07,
		SubdivisionIR08,
		SubdivisionIR10,
		SubdivisionIR11,
		SubdivisionIR12,
		SubdivisionIR13,
		SubdivisionIR14,
		SubdivisionIR15,
		SubdivisionIR16,
		SubdivisionIR17,
		SubdivisionIR18,
		SubdivisionIR19,
		SubdivisionIR20,
		SubdivisionIR21,
		SubdivisionIR22,
		SubdivisionIR23,
		SubdivisionIR24,
		SubdivisionIR25,
		SubdivisionIR26,
		SubdivisionIR27,
		SubdivisionIR28,
		SubdivisionIR29,
		SubdivisionIR30,
		SubdivisionIR31,
		SubdivisionIS0,
		SubdivisionIS1,
		SubdivisionIS2,
		SubdivisionIS3,
		SubdivisionIS4,
		SubdivisionIS5,
		SubdivisionIS6,
		SubdivisionIS7,
		SubdivisionIS8,
		SubdivisionIT21,
		SubdivisionIT23,
		SubdivisionIT25,
		SubdivisionIT32,
		SubdivisionIT34,
		SubdivisionIT36,
		SubdivisionIT42,
		SubdivisionIT45,
		SubdivisionIT52,
		SubdivisionIT55,
		SubdivisionIT57,
		SubdivisionIT62,
		SubdivisionIT65,
		SubdivisionIT67,
		SubdivisionIT72,
		SubdivisionIT75,
		SubdivisionIT77,
		SubdivisionIT78,
		SubdivisionIT82,
		SubdivisionIT88,
		SubdivisionITAG,
		SubdivisionITAL,
		SubdivisionITAN,
		SubdivisionITAO,
		SubdivisionITAP,
		SubdivisionITAQ,
		SubdivisionITAR,
		SubdivisionITAT,
		SubdivisionITAV,
		SubdivisionITBA,
		SubdivisionITBG,
		SubdivisionITBI,
		SubdivisionITBL,
		SubdivisionITBN,
		SubdivisionITBO,
		SubdivisionITBR,
		SubdivisionITBS,
		SubdivisionITBT,
		SubdivisionITBZ,
		SubdivisionITCA,
		SubdivisionITCB,
		SubdivisionITCE,
		SubdivisionITCH,
		SubdivisionITCI,
		SubdivisionITCL,
		SubdivisionITCN,
		SubdivisionITCO,
		SubdivisionITCR,
		SubdivisionITCS,
		SubdivisionITCT,
		SubdivisionITCZ,
		SubdivisionITEN,
		SubdivisionITFC,
		SubdivisionITFE,
		SubdivisionITFG,
		SubdivisionITFI,
		SubdivisionITFM,
		SubdivisionITFR,
		SubdivisionITGE,
		SubdivisionITGO,
		SubdivisionITGR,
		SubdivisionITIM,
		SubdivisionITIS,
		SubdivisionITKR,
		SubdivisionITLC,
		SubdivisionITLE,
		SubdivisionITLI,
		SubdivisionITLO,
		SubdivisionITLT,
		SubdivisionITLU,
		SubdivisionITMB,
		SubdivisionITMC,
		SubdivisionITME,
		SubdivisionITMI,
		SubdivisionITMN,
		SubdivisionITMO,
		SubdivisionITMS,
		SubdivisionITMT,
		SubdivisionITNA,
		SubdivisionITNO,
		SubdivisionITNU,
		SubdivisionITOG,
		SubdivisionITOR,
		SubdivisionITOT,
		SubdivisionITPA,
		SubdivisionITPC,
		SubdivisionITPD,
		SubdivisionITPE,
		SubdivisionITPG,
		SubdivisionITPI,
		SubdivisionITPN,
		SubdivisionITPO,
		SubdivisionITPR,
		SubdivisionITPT,
		SubdivisionITPU,
		SubdivisionITPV,
		SubdivisionITPZ,
		SubdivisionITRA,
		SubdivisionITRC,
		SubdivisionITRE,
		SubdivisionITRG,
		SubdivisionITRI,
		SubdivisionITRM,
		SubdivisionITRN,
		SubdivisionITRO,
		SubdivisionITSA,
		SubdivisionITSI,
		SubdivisionITSO,
		SubdivisionITSP,
		SubdivisionITSR,
		SubdivisionITSS,
		SubdivisionITSV,
		SubdivisionITTA,
		SubdivisionITTE,
		SubdivisionITTN,
		SubdivisionITTO,
		SubdivisionITTP,
		SubdivisionITTR,
		SubdivisionITTS,
		SubdivisionITTV,
		SubdivisionITUD,
		SubdivisionITVA,
		SubdivisionITVB,
		SubdivisionITVC,
		SubdivisionITVE,
		SubdivisionITVI,
		SubdivisionITVR,
		SubdivisionITVS,
		SubdivisionITVT,
		SubdivisionITVV,
		SubdivisionJM01,
		SubdivisionJM02,
		SubdivisionJM03,
		SubdivisionJM04,
		SubdivisionJM05,
		SubdivisionJM06,
		SubdivisionJM07,
		SubdivisionJM08,
		SubdivisionJM09,
		SubdivisionJM10,
		SubdivisionJM11,
		SubdivisionJM12,
		SubdivisionJM13,
		SubdivisionJM14,
		SubdivisionJOAJ,
		SubdivisionJOAM,
		SubdivisionJOAQ,
		SubdivisionJOAT,
		SubdivisionJOAZ,
		SubdivisionJOBA,
		SubdivisionJOIR,
		SubdivisionJOJA,
		SubdivisionJOKA,
		SubdivisionJOMA,
		SubdivisionJOMD,
		SubdivisionJOMN,
		SubdivisionJP01,
		SubdivisionJP02,
		SubdivisionJP03,
		SubdivisionJP04,
		SubdivisionJP05,
		SubdivisionJP06,
		SubdivisionJP07,
		SubdivisionJP08,
		SubdivisionJP09,
		SubdivisionJP10,
		SubdivisionJP11,
		SubdivisionJP12,
		SubdivisionJP13,
		SubdivisionJP14,
		SubdivisionJP15,
		SubdivisionJP16,
		SubdivisionJP17,
		SubdivisionJP18,
		SubdivisionJP19,
		SubdivisionJP20,
		SubdivisionJP21,
		SubdivisionJP22,
		SubdivisionJP23,
		SubdivisionJP24,
		SubdivisionJP25,
		SubdivisionJP26,
		SubdivisionJP27,
		SubdivisionJP28,
		SubdivisionJP29,
		SubdivisionJP30,
		SubdivisionJP31,
		SubdivisionJP32,
		SubdivisionJP33,
		SubdivisionJP34,
		SubdivisionJP35,
		SubdivisionJP36,
		SubdivisionJP37,
		SubdivisionJP38,
		SubdivisionJP39,
		SubdivisionJP40,
		SubdivisionJP41,
		SubdivisionJP42,
		SubdivisionJP43,
		SubdivisionJP44,
		SubdivisionJP45,
		SubdivisionJP46,
		SubdivisionJP47,
		SubdivisionKE01,
		SubdivisionKE02,
		SubdivisionKE03,
		SubdivisionKE04,
		SubdivisionKE05,
		SubdivisionKE06,
		SubdivisionKE07,
		SubdivisionKE08,
		SubdivisionKE09,
		SubdivisionKE10,
		SubdivisionKE11,
		SubdivisionKE12,
		SubdivisionKE13,
		SubdivisionKE14,
		SubdivisionKE15,
		SubdivisionKE16,
		SubdivisionKE17,
		SubdivisionKE18,
		SubdivisionKE19,
		SubdivisionKE20,
		SubdivisionKE21,
		SubdivisionKE22,
		SubdivisionKE23,
		SubdivisionKE24,
		SubdivisionKE25,
		SubdivisionKE26,
		SubdivisionKE27,
		SubdivisionKE28,
		SubdivisionKE29,
		SubdivisionKE30,
		SubdivisionKE31,
		SubdivisionKE32,
		SubdivisionKE33,
		SubdivisionKE34,
		SubdivisionKE35,
		SubdivisionKE36,
		SubdivisionKE37,
		SubdivisionKE38,
		SubdivisionKE39,
		SubdivisionKE40,
		SubdivisionKE41,
		SubdivisionKE42,
		SubdivisionKE43,
		SubdivisionKE44,
		SubdivisionKE45,
		SubdivisionKE46,
		SubdivisionKE47,
		SubdivisionKGB,
		SubdivisionKGC,
		SubdivisionKGGB,
		SubdivisionKGJ,
		SubdivisionKGN,
		SubdivisionKGO,
		SubdivisionKGT,
		SubdivisionKGY,
		SubdivisionKH1,
		SubdivisionKH10,
		SubdivisionKH11,
		SubdivisionKH12,
		SubdivisionKH13,
		SubdivisionKH14,
		SubdivisionKH15,
		SubdivisionKH16,
		SubdivisionKH17,
		SubdivisionKH18,
		SubdivisionKH19,
		SubdivisionKH2,
		SubdivisionKH20,
		SubdivisionKH21,
		SubdivisionKH22,
		SubdivisionKH23,
		SubdivisionKH24,
		SubdivisionKH3,
		SubdivisionKH4,
		SubdivisionKH5,
		SubdivisionKH6,
		SubdivisionKH7,
		SubdivisionKH8,
		SubdivisionKH9,
		SubdivisionKIG,
		SubdivisionKIL,
		SubdivisionKIP,
		SubdivisionKMA,
		SubdivisionKMG,
		SubdivisionKMM,
		SubdivisionKN01,
		SubdivisionKN02,
		SubdivisionKN03,
		SubdivisionKN04,
		SubdivisionKN05,
		SubdivisionKN06,
		SubdivisionKN07,
		SubdivisionKN08,
		SubdivisionKN09,
		SubdivisionKN10,
		SubdivisionKN11,
		SubdivisionKN12,
		SubdivisionKN13,
		SubdivisionKN15,
		SubdivisionKNK,
		SubdivisionKNN,
		SubdivisionKP01,
		SubdivisionKP02,
		SubdivisionKP03,
		SubdivisionKP04,
		SubdivisionKP05,
		SubdivisionKP06,
		SubdivisionKP07,
		SubdivisionKP08,
		SubdivisionKP09,
		SubdivisionKP10,
		SubdivisionKP13,
		SubdivisionKR11,
		SubdivisionKR26,
		SubdivisionKR27,
		SubdivisionKR28,
		SubdivisionKR29,
		SubdivisionKR30,
		SubdivisionKR31,
		SubdivisionKR41,
		SubdivisionKR42,
		SubdivisionKR43,
		SubdivisionKR44,
		SubdivisionKR45,
		SubdivisionKR46,
		SubdivisionKR47,
		SubdivisionKR48,
		SubdivisionKR49,
		SubdivisionKWAH,
		SubdivisionKWFA,
		SubdivisionKWHA,
		SubdivisionKWJA,
		SubdivisionKWKU,
		SubdivisionKWMU,
		SubdivisionKZAKM,
		SubdivisionKZAKT,
		SubdivisionKZALA,
		SubdivisionKZALM,
		SubdivisionKZAST,
		SubdivisionKZATY,
		SubdivisionKZKAR,
		SubdivisionKZKUS,
		SubdivisionKZKZY,
		SubdivisionKZMAN,
		SubdivisionKZPAV,
		SubdivisionKZSEV,
		SubdivisionKZVOS,
		SubdivisionKZYUZ,
		SubdivisionKZZAP,
		SubdivisionKZZHA,
		SubdivisionLAAT,
		SubdivisionLABK,
		SubdivisionLABL,
		SubdivisionLACH,
		SubdivisionLAHO,
		SubdivisionLAKH,
		SubdivisionLALM,
		SubdivisionLALP,
		SubdivisionLAOU,
		SubdivisionLAPH,
		SubdivisionLASL,
		SubdivisionLASV,
		SubdivisionLAVI,
		SubdivisionLAVT,
		SubdivisionLAXA,
		SubdivisionLAXE,
		SubdivisionLAXI,
		SubdivisionLAXS,
		SubdivisionLBAK,
		SubdivisionLBAS,
		SubdivisionLBBA,
		SubdivisionLBBH,
		SubdivisionLBBI,
		SubdivisionLBJA,
		SubdivisionLBJL,
		SubdivisionLBNA,
		SubdivisionLI01,
		SubdivisionLI02,
		SubdivisionLI03,
		SubdivisionLI04,
		SubdivisionLI05,
		SubdivisionLI06,
		SubdivisionLI07,
		SubdivisionLI08,
		SubdivisionLI09,
		SubdivisionLI10,
		SubdivisionLI11,
		SubdivisionLK1,
		SubdivisionLK11,
		SubdivisionLK12,
		SubdivisionLK13,
		SubdivisionLK2,
		SubdivisionLK21,
		SubdivisionLK22,
		SubdivisionLK23,
		SubdivisionLK3,
		SubdivisionLK31,
		SubdivisionLK32,
		SubdivisionLK33,
		SubdivisionLK4,
		SubdivisionLK41,
		SubdivisionLK42,
		SubdivisionLK43,
		SubdivisionLK44,
		SubdivisionLK45,
		SubdivisionLK5,
		SubdivisionLK51,
		SubdivisionLK52,
		SubdivisionLK53,
		SubdivisionLK6,
		SubdivisionLK61,
		SubdivisionLK62,
		SubdivisionLK7,
		SubdivisionLK71,
		SubdivisionLK72,
		SubdivisionLK8,
		SubdivisionLK81,
		SubdivisionLK82,
		SubdivisionLK9,
		SubdivisionLK91,
		SubdivisionLK92,
		SubdivisionLRBG,
		SubdivisionLRBM,
		SubdivisionLRCM,
		SubdivisionLRGB,
		SubdivisionLRGG,
		SubdivisionLRGK,
		SubdivisionLRLO,
		SubdivisionLRMG,
		SubdivisionLRMO,
		SubdivisionLRMY,
		SubdivisionLRNI,
		SubdivisionLRRI,
		SubdivisionLRSI,
		SubdivisionLSA,
		SubdivisionLSB,
		SubdivisionLSC,
		SubdivisionLSD,
		SubdivisionLSE,
		SubdivisionLSF,
		SubdivisionLSG,
		SubdivisionLSH,
		SubdivisionLSJ,
		SubdivisionLSK,
		SubdivisionLTAL,
		SubdivisionLTKL,
		SubdivisionLTKU,
		SubdivisionLTMR,
		SubdivisionLTPN,
		SubdivisionLTSA,
		SubdivisionLTTA,
		SubdivisionLTTE,
		SubdivisionLTUT,
		SubdivisionLTVL,
		SubdivisionLUD,
		SubdivisionLUG,
		SubdivisionLUL,
		SubdivisionLV001,
		SubdivisionLV002,
		SubdivisionLV003,
		SubdivisionLV004,
		SubdivisionLV005,
		SubdivisionLV006,
		SubdivisionLV007,
		SubdivisionLV008,
		SubdivisionLV009,
		SubdivisionLV010,
		SubdivisionLV011,
		SubdivisionLV012,
		SubdivisionLV013,
		SubdivisionLV014,
		SubdivisionLV015,
		SubdivisionLV016,
		SubdivisionLV017,
		SubdivisionLV018,
		SubdivisionLV019,
		SubdivisionLV020,
		SubdivisionLV021,
		SubdivisionLV022,
		SubdivisionLV023,
		SubdivisionLV024,
		SubdivisionLV025,
		SubdivisionLV026,
		SubdivisionLV027,
		SubdivisionLV028,
		SubdivisionLV029,
		SubdivisionLV030,
		SubdivisionLV031,
		SubdivisionLV032,
		SubdivisionLV033,
		SubdivisionLV034,
		SubdivisionLV035,
		SubdivisionLV036,
		SubdivisionLV037,
		SubdivisionLV038,
		SubdivisionLV039,
		SubdivisionLV040,
		SubdivisionLV041,
		SubdivisionLV042,
		SubdivisionLV043,
		SubdivisionLV044,
		SubdivisionLV045,
		SubdivisionLV046,
		SubdivisionLV047,
		SubdivisionLV048,
		SubdivisionLV049,
		SubdivisionLV050,
		SubdivisionLV051,
		SubdivisionLV052,
		SubdivisionLV053,
		SubdivisionLV054,
		SubdivisionLV055,
		SubdivisionLV056,
		SubdivisionLV057,
		SubdivisionLV058,
		SubdivisionLV059,
		SubdivisionLV060,
		SubdivisionLV061,
		SubdivisionLV062,
		SubdivisionLV063,
		SubdivisionLV064,
		SubdivisionLV065,
		SubdivisionLV066,
		SubdivisionLV067,
		SubdivisionLV068,
		SubdivisionLV069,
		SubdivisionLV070,
		SubdivisionLV071,
		SubdivisionLV072,
		SubdivisionLV073,
		SubdivisionLV074,
		SubdivisionLV075,
		SubdivisionLV076,
		SubdivisionLV077,
		SubdivisionLV078,
		SubdivisionLV079,
		SubdivisionLV080,
		SubdivisionLV081,
		SubdivisionLV082,
		SubdivisionLV083,
		SubdivisionLV084,
		SubdivisionLV085,
		SubdivisionLV086,
		SubdivisionLV087,
		SubdivisionLV088,
		SubdivisionLV089,
		SubdivisionLV090,
		SubdivisionLV091,
		SubdivisionLV092,
		SubdivisionLV093,
		SubdivisionLV094,
		SubdivisionLV095,
		SubdivisionLV096,
		SubdivisionLV097,
		SubdivisionLV098,
		SubdivisionLV099,
		SubdivisionLV100,
		SubdivisionLV101,
		SubdivisionLV102,
		SubdivisionLV103,
		SubdivisionLV104,
		SubdivisionLV105,
		SubdivisionLV106,
		SubdivisionLV107,
		SubdivisionLV108,
		SubdivisionLV109,
		SubdivisionLV110,
		SubdivisionLVDGV,
		SubdivisionLVJEL,
		SubdivisionLVJKB,
		SubdivisionLVJUR,
		SubdivisionLVLPX,
		SubdivisionLVREZ,
		SubdivisionLVRIX,
		SubdivisionLVVEN,
		SubdivisionLVVMR,
		SubdivisionLYBA,
		SubdivisionLYBU,
		SubdivisionLYDR,
		SubdivisionLYGT,
		SubdivisionLYJA,
		SubdivisionLYJB,
		SubdivisionLYJG,
		SubdivisionLYJI,
		SubdivisionLYJU,
		SubdivisionLYKF,
		SubdivisionLYMB,
		SubdivisionLYMI,
		SubdivisionLYMJ,
		SubdivisionLYMQ,
		SubdivisionLYNL,
		SubdivisionLYNQ,
		SubdivisionLYSB,
		SubdivisionLYSR,
		SubdivisionLYTB,
		SubdivisionLYWA,
		SubdivisionLYWD,
		SubdivisionLYWS,
		SubdivisionLYZA,
		SubdivisionMA01,
		SubdivisionMA02,
		SubdivisionMA03,
		SubdivisionMA04,
		SubdivisionMA05,
		SubdivisionMA06,
		SubdivisionMA07,
		SubdivisionMA08,
		SubdivisionMA09,
		SubdivisionMA10,
		SubdivisionMA11,
		SubdivisionMA12,
		SubdivisionMAAGD,
		SubdivisionMAAOU,
		SubdivisionMAASZ,
		SubdivisionMAAZI,
		SubdivisionMABEM,
		SubdivisionMABER,
		SubdivisionMABES,
		SubdivisionMABOD,
		SubdivisionMABOM,
		SubdivisionMABRR,
		SubdivisionMACAS,
		SubdivisionMACHE,
		SubdivisionMACHI,
		SubdivisionMACHT,
		SubdivisionMADRI,
		SubdivisionMAERR,
		SubdivisionMAESI,
		SubdivisionMAESM,
		SubdivisionMAFAH,
		SubdivisionMAFES,
		SubdivisionMAFIG,
		SubdivisionMAFQH,
		SubdivisionMAGUE,
		SubdivisionMAGUF,
		SubdivisionMAHAJ,
		SubdivisionMAHAO,
		SubdivisionMAHOC,
		SubdivisionMAIFR,
		SubdivisionMAINE,
		SubdivisionMAJDI,
		SubdivisionMAJRA,
		SubdivisionMAKEN,
		SubdivisionMAKES,
		SubdivisionMAKHE,
		SubdivisionMAKHN,
		SubdivisionMAKHO,
		SubdivisionMALAA,
		SubdivisionMALAR,
		SubdivisionMAMAR,
		SubdivisionMAMDF,
		SubdivisionMAMED,
		SubdivisionMAMEK,
		SubdivisionMAMID,
		SubdivisionMAMOH,
		SubdivisionMAMOU,
		SubdivisionMANAD,
		SubdivisionMANOU,
		SubdivisionMAOUA,
		SubdivisionMAOUD,
		SubdivisionMAOUJ,
		SubdivisionMAOUZ,
		SubdivisionMARAB,
		SubdivisionMAREH,
		SubdivisionMASAF,
		SubdivisionMASAL,
		SubdivisionMASEF,
		SubdivisionMASET,
		SubdivisionMASIB,
		SubdivisionMASIF,
		SubdivisionMASIK,
		SubdivisionMASIL,
		SubdivisionMASKH,
		SubdivisionMATAF,
		SubdivisionMATAI,
		SubdivisionMATAO,
		SubdivisionMATAR,
		SubdivisionMATAT,
		SubdivisionMATAZ,
		SubdivisionMATET,
		SubdivisionMATIN,
		SubdivisionMATIZ,
		SubdivisionMATNG,
		SubdivisionMATNT,
		SubdivisionMAYUS,
		SubdivisionMAZAG,
		SubdivisionMCCL,
		SubdivisionMCCO,
		SubdivisionMCFO,
		SubdivisionMCGA,
		SubdivisionMCJE,
		SubdivisionMCLA,
		SubdivisionMCMA,
		SubdivisionMCMC,
		SubdivisionMCMG,
		SubdivisionMCMO,
		SubdivisionMCMU,
		SubdivisionMCPH,
		SubdivisionMCSD,
		SubdivisionMCSO,
		SubdivisionMCSP,
		SubdivisionMCSR,
		SubdivisionMCVR,
		SubdivisionMDAN,
		SubdivisionMDBA,
		SubdivisionMDBD,
		SubdivisionMDBR,
		SubdivisionMDBS,
		SubdivisionMDCA,
		SubdivisionMDCL,
		SubdivisionMDCM,
		SubdivisionMDCR,
		SubdivisionMDCS,
		SubdivisionMDCT,
		SubdivisionMDCU,
		SubdivisionMDDO,
		SubdivisionMDDR,
		SubdivisionMDDU,
		SubdivisionMDED,
		SubdivisionMDFA,
		SubdivisionMDFL,
		SubdivisionMDGA,
		SubdivisionMDGL,
		SubdivisionMDHI,
		SubdivisionMDIA,
		SubdivisionMDLE,
		SubdivisionMDNI,
		SubdivisionMDOC,
		SubdivisionMDOR,
		SubdivisionMDRE,
		SubdivisionMDRI,
		SubdivisionMDSD,
		SubdivisionMDSI,
		SubdivisionMDSN,
		SubdivisionMDSO,
		SubdivisionMDST,
		SubdivisionMDSV,
		SubdivisionMDTA,
		SubdivisionMDTE,
		SubdivisionMDUN,
		SubdivisionME01,
		SubdivisionME02,
		SubdivisionME03,
		SubdivisionME04,
		SubdivisionME05,
		SubdivisionME06,
		SubdivisionME07,
		SubdivisionME08,
		SubdivisionME09,
		SubdivisionME10,
		SubdivisionME11,
		SubdivisionME12,
		SubdivisionME13,
		SubdivisionME14,
		SubdivisionME15,
		SubdivisionME16,
		SubdivisionME17,
		SubdivisionME18,
		SubdivisionME19,
		SubdivisionME20,
		SubdivisionME21,
		SubdivisionMGA,
		SubdivisionMGD,
		SubdivisionMGF,
		SubdivisionMGM,
		SubdivisionMGT,
		SubdivisionMGU,
		SubdivisionMHALK,
		SubdivisionMHALL,
		SubdivisionMHARN,
		SubdivisionMHAUR,
		SubdivisionMHEBO,
		SubdivisionMHENI,
		SubdivisionMHJAB,
		SubdivisionMHJAL,
		SubdivisionMHKIL,
		SubdivisionMHKWA,
		SubdivisionMHL,
		SubdivisionMHLAE,
		SubdivisionMHLIB,
		SubdivisionMHLIK,
		SubdivisionMHMAJ,
		SubdivisionMHMAL,
		SubdivisionMHMEJ,
		SubdivisionMHMIL,
		SubdivisionMHNMK,
		SubdivisionMHNMU,
		SubdivisionMHRON,
		SubdivisionMHT,
		SubdivisionMHUJA,
		SubdivisionMHUTI,
		SubdivisionMHWTJ,
		SubdivisionMHWTN,
		SubdivisionMK01,
		SubdivisionMK02,
		SubdivisionMK03,
		SubdivisionMK04,
		SubdivisionMK05,
		SubdivisionMK06,
		SubdivisionMK07,
		SubdivisionMK08,
		SubdivisionMK09,
		SubdivisionMK10,
		SubdivisionMK11,
		SubdivisionMK12,
		SubdivisionMK13,
		SubdivisionMK14,
		SubdivisionMK15,
		SubdivisionMK16,
		SubdivisionMK17,
		SubdivisionMK18,
		SubdivisionMK19,
		SubdivisionMK20,
		SubdivisionMK21,
		SubdivisionMK22,
		SubdivisionMK23,
		SubdivisionMK24,
		SubdivisionMK25,
		SubdivisionMK26,
		SubdivisionMK27,
		SubdivisionMK28,
		SubdivisionMK29,
		SubdivisionMK30,
		SubdivisionMK31,
		SubdivisionMK32,
		SubdivisionMK33,
		SubdivisionMK34,
		SubdivisionMK35,
		SubdivisionMK36,
		SubdivisionMK37,
		SubdivisionMK38,
		SubdivisionMK39,
		SubdivisionMK40,
		SubdivisionMK41,
		SubdivisionMK42,
		SubdivisionMK43,
		SubdivisionMK44,
		SubdivisionMK45,
		SubdivisionMK46,
		SubdivisionMK47,
		SubdivisionMK48,
		SubdivisionMK49,
		SubdivisionMK50,
		SubdivisionMK51,
		SubdivisionMK52,
		SubdivisionMK53,
		SubdivisionMK54,
		SubdivisionMK55,
		SubdivisionMK56,
		SubdivisionMK57,
		SubdivisionMK58,
		SubdivisionMK59,
		SubdivisionMK60,
		SubdivisionMK61,
		SubdivisionMK62,
		SubdivisionMK63,
		SubdivisionMK64,
		SubdivisionMK65,
		SubdivisionMK66,
		SubdivisionMK67,
		SubdivisionMK68,
		SubdivisionMK69,
		SubdivisionMK70,
		SubdivisionMK71,
		SubdivisionMK72,
		SubdivisionMK73,
		SubdivisionMK74,
		SubdivisionMK75,
		SubdivisionMK76,
		SubdivisionMK77,
		SubdivisionMK78,
		SubdivisionMK79,
		SubdivisionMK80,
		SubdivisionMK81,
		SubdivisionMK82,
		SubdivisionMK83,
		SubdivisionMK84,
		SubdivisionML1,
		SubdivisionML2,
		SubdivisionML3,
		SubdivisionML4,
		SubdivisionML5,
		SubdivisionML6,
		SubdivisionML7,
		SubdivisionML8,
		SubdivisionMLBK0,
		SubdivisionMM01,
		SubdivisionMM02,
		SubdivisionMM03,
		SubdivisionMM04,
		SubdivisionMM05,
		SubdivisionMM06,
		SubdivisionMM07,
		SubdivisionMM11,
		SubdivisionMM12,
		SubdivisionMM13,
		SubdivisionMM14,
		SubdivisionMM15,
		SubdivisionMM16,
		SubdivisionMM17,
		SubdivisionMN035,
		SubdivisionMN037,
		SubdivisionMN039,
		SubdivisionMN041,
		SubdivisionMN043,
		SubdivisionMN046,
		SubdivisionMN047,
		SubdivisionMN049,
		SubdivisionMN051,
		SubdivisionMN053,
		SubdivisionMN055,
		SubdivisionMN057,
		SubdivisionMN059,
		SubdivisionMN061,
		SubdivisionMN063,
		SubdivisionMN064,
		SubdivisionMN065,
		SubdivisionMN067,
		SubdivisionMN069,
		SubdivisionMN071,
		SubdivisionMN073,
		SubdivisionMN1,
		SubdivisionMR01,
		SubdivisionMR02,
		SubdivisionMR03,
		SubdivisionMR04,
		SubdivisionMR05,
		SubdivisionMR06,
		SubdivisionMR07,
		SubdivisionMR08,
		SubdivisionMR09,
		SubdivisionMR10,
		SubdivisionMR11,
		SubdivisionMR12,
		SubdivisionMRNKC,
		SubdivisionMT01,
		SubdivisionMT02,
		SubdivisionMT03,
		SubdivisionMT04,
		SubdivisionMT05,
		SubdivisionMT06,
		SubdivisionMT07,
		SubdivisionMT08,
		SubdivisionMT09,
		SubdivisionMT10,
		SubdivisionMT11,
		SubdivisionMT12,
		SubdivisionMT13,
		SubdivisionMT14,
		SubdivisionMT15,
		SubdivisionMT16,
		SubdivisionMT17,
		SubdivisionMT18,
		SubdivisionMT19,
		SubdivisionMT20,
		SubdivisionMT21,
		SubdivisionMT22,
		SubdivisionMT23,
		SubdivisionMT24,
		SubdivisionMT25,
		SubdivisionMT26,
		SubdivisionMT27,
		SubdivisionMT28,
		SubdivisionMT29,
		SubdivisionMT30,
		SubdivisionMT31,
		SubdivisionMT32,
		SubdivisionMT33,
		SubdivisionMT34,
		SubdivisionMT35,
		SubdivisionMT36,
		SubdivisionMT37,
		SubdivisionMT38,
		SubdivisionMT39,
		SubdivisionMT40,
		SubdivisionMT41,
		SubdivisionMT42,
		SubdivisionMT43,
		SubdivisionMT44,
		SubdivisionMT45,
		SubdivisionMT46,
		SubdivisionMT47,
		SubdivisionMT48,
		SubdivisionMT49,
		SubdivisionMT50,
		SubdivisionMT51,
		SubdivisionMT52,
		SubdivisionMT53,
		SubdivisionMT54,
		SubdivisionMT55,
		SubdivisionMT56,
		SubdivisionMT57,
		SubdivisionMT58,
		SubdivisionMT59,
		SubdivisionMT60,
		SubdivisionMT61,
		SubdivisionMT62,
		SubdivisionMT63,
		SubdivisionMT64,
		SubdivisionMT65,
		SubdivisionMT66,
		SubdivisionMT67,
		SubdivisionMT68,
		SubdivisionMUAG,
		SubdivisionMUBL,
		SubdivisionMUBR,
		SubdivisionMUCC,
		SubdivisionMUCU,
		SubdivisionMUFL,
		SubdivisionMUGP,
		SubdivisionMUMO,
		SubdivisionMUPA,
		SubdivisionMUPL,
		SubdivisionMUPU,
		SubdivisionMUPW,
		SubdivisionMUQB,
		SubdivisionMURO,
		SubdivisionMURP,
		SubdivisionMUSA,
		SubdivisionMUVP,
		SubdivisionMV00,
		SubdivisionMV01,
		SubdivisionMV02,
		SubdivisionMV03,
		SubdivisionMV04,
		SubdivisionMV05,
		SubdivisionMV07,
		SubdivisionMV08,
		SubdivisionMV12,
		SubdivisionMV13,
		SubdivisionMV14,
		SubdivisionMV17,
		SubdivisionMV20,
		SubdivisionMV23,
		SubdivisionMV24,
		SubdivisionMV25,
		SubdivisionMV26,
		SubdivisionMV27,
		SubdivisionMV28,
		SubdivisionMV29,
		SubdivisionMVCE,
		SubdivisionMVMLE,
		SubdivisionMVNC,
		SubdivisionMVNO,
		SubdivisionMVSC,
		SubdivisionMVSU,
		SubdivisionMVUN,
		SubdivisionMVUS,
		SubdivisionMWBA,
		SubdivisionMWBL,
		SubdivisionMWC,
		SubdivisionMWCK,
		SubdivisionMWCR,
		SubdivisionMWCT,
		SubdivisionMWDE,
		SubdivisionMWDO,
		SubdivisionMWKR,
		SubdivisionMWKS,
		SubdivisionMWLI,
		SubdivisionMWLK,
		SubdivisionMWMC,
		SubdivisionMWMG,
		SubdivisionMWMH,
		SubdivisionMWMU,
		SubdivisionMWMW,
		SubdivisionMWMZ,
		SubdivisionMWN,
		SubdivisionMWNB,
		SubdivisionMWNE,
		SubdivisionMWNI,
		SubdivisionMWNK,
		SubdivisionMWNS,
		SubdivisionMWNU,
		SubdivisionMWPH,
		SubdivisionMWRU,
		SubdivisionMWS,
		SubdivisionMWSA,
		SubdivisionMWTH,
		SubdivisionMWZO,
		SubdivisionMXAGU,
		SubdivisionMXBCN,
		SubdivisionMXBCS,
		SubdivisionMXCAM,
		SubdivisionMXCHH,
		SubdivisionMXCHP,
		SubdivisionMXCMX,
		SubdivisionMXCOA,
		SubdivisionMXCOL,
		SubdivisionMXDUR,
		SubdivisionMXGRO,
		SubdivisionMXGUA,
		SubdivisionMXHID,
		SubdivisionMXJAL,
		SubdivisionMXMEX,
		SubdivisionMXMIC,
		SubdivisionMXMOR,
		SubdivisionMXNAY,
		SubdivisionMXNLE,
		SubdivisionMXOAX,
		SubdivisionMXPUE,
		SubdivisionMXQUE,
		SubdivisionMXROO,
		SubdivisionMXSIN,
		SubdivisionMXSLP,
		SubdivisionMXSON,
		SubdivisionMXTAB,
		SubdivisionMXTAM,
		SubdivisionMXTLA,
		SubdivisionMXVER,
		SubdivisionMXYUC,
		SubdivisionMXZAC,
		SubdivisionMY01,
		SubdivisionMY02,
		SubdivisionMY03,
		SubdivisionMY04,
		SubdivisionMY05,
		SubdivisionMY06,
		SubdivisionMY07,
		SubdivisionMY08,
		SubdivisionMY09,
		SubdivisionMY10,
		SubdivisionMY11,
		SubdivisionMY12,
		SubdivisionMY13,
		SubdivisionMY14,
		SubdivisionMY15,
		SubdivisionMY16,
		SubdivisionMZA,
		SubdivisionMZB,
		SubdivisionMZG,
		SubdivisionMZI,
		SubdivisionMZL,
		SubdivisionMZMPM,
		SubdivisionMZN,
		SubdivisionMZP,
		SubdivisionMZQ,
		SubdivisionMZS,
		SubdivisionMZT,
		SubdivisionNACA,
		SubdivisionNAER,
		SubdivisionNAHA,
		SubdivisionNAKA,
		SubdivisionNAKH,
		SubdivisionNAKU,
		SubdivisionNAOD,
		SubdivisionNAOH,
		SubdivisionNAOK,
		SubdivisionNAON,
		SubdivisionNAOS,
		SubdivisionNAOT,
		SubdivisionNAOW,
		SubdivisionNE1,
		SubdivisionNE2,
		SubdivisionNE3,
		SubdivisionNE4,
		SubdivisionNE5,
		SubdivisionNE6,
		SubdivisionNE7,
		SubdivisionNE8,
		SubdivisionNGAB,
		SubdivisionNGAD,
		SubdivisionNGAK,
		SubdivisionNGAN,
		SubdivisionNGBA,
		SubdivisionNGBE,
		SubdivisionNGBO,
		SubdivisionNGBY,
		SubdivisionNGCR,
		SubdivisionNGDE,
		SubdivisionNGEB,
		SubdivisionNGED,
		SubdivisionNGEK,
		SubdivisionNGEN,
		SubdivisionNGFC,
		SubdivisionNGGO,
		SubdivisionNGIM,
		SubdivisionNGJI,
		SubdivisionNGKD,
		SubdivisionNGKE,
		SubdivisionNGKN,
		SubdivisionNGKO,
		SubdivisionNGKT,
		SubdivisionNGKW,
		SubdivisionNGLA,
		SubdivisionNGNA,
		SubdivisionNGNI,
		SubdivisionNGOG,
		SubdivisionNGON,
		SubdivisionNGOS,
		SubdivisionNGOY,
		SubdivisionNGPL,
		SubdivisionNGRI,
		SubdivisionNGSO,
		SubdivisionNGTA,
		SubdivisionNGYO,
		SubdivisionNGZA,
		SubdivisionNIAN,
		SubdivisionNIAS,
		SubdivisionNIBO,
		SubdivisionNICA,
		SubdivisionNICI,
		SubdivisionNICO,
		SubdivisionNIES,
		SubdivisionNIGR,
		SubdivisionNIJI,
		SubdivisionNILE,
		SubdivisionNIMD,
		SubdivisionNIMN,
		SubdivisionNIMS,
		SubdivisionNIMT,
		SubdivisionNINS,
		SubdivisionNIRI,
		SubdivisionNISJ,
		SubdivisionNLAW,
		SubdivisionNLBQ1,
		SubdivisionNLBQ2,
		SubdivisionNLBQ3,
		SubdivisionNLCW,
		SubdivisionNLDR,
		SubdivisionNLFL,
		SubdivisionNLFR,
		SubdivisionNLGE,
		SubdivisionNLGR,
		SubdivisionNLLI,
		SubdivisionNLNB,
		SubdivisionNLNH,
		SubdivisionNLOV,
		SubdivisionNLSX,
		SubdivisionNLUT,
		SubdivisionNLZE,
		SubdivisionNLZH,
		SubdivisionNO01,
		SubdivisionNO02,
		SubdivisionNO03,
		SubdivisionNO04,
		SubdivisionNO05,
		SubdivisionNO06,
		SubdivisionNO07,
		SubdivisionNO08,
		SubdivisionNO09,
		SubdivisionNO10,
		SubdivisionNO11,
		SubdivisionNO12,
		SubdivisionNO14,
		SubdivisionNO15,
		SubdivisionNO18,
		SubdivisionNO19,
		SubdivisionNO20,
		SubdivisionNO21,
		SubdivisionNO22,
		SubdivisionNO50,
		SubdivisionNP1,
		SubdivisionNP2,
		SubdivisionNP3,
		SubdivisionNP4,
		SubdivisionNP5,
		SubdivisionNPBA,
		SubdivisionNPBH,
		SubdivisionNPDH,
		SubdivisionNPGA,
		SubdivisionNPJA,
		SubdivisionNPKA,
		SubdivisionNPKO,
		SubdivisionNPLU,
		SubdivisionNPMA,
		SubdivisionNPME,
		SubdivisionNPNA,
		SubdivisionNPRA,
		SubdivisionNPSA,
		SubdivisionNPSE,
		SubdivisionNR01,
		SubdivisionNR02,
		SubdivisionNR03,
		SubdivisionNR04,
		SubdivisionNR05,
		SubdivisionNR06,
		SubdivisionNR07,
		SubdivisionNR08,
		SubdivisionNR09,
		SubdivisionNR10,
		SubdivisionNR11,
		SubdivisionNR12,
		SubdivisionNR13,
		SubdivisionNR14,
		SubdivisionNZAUK,
		SubdivisionNZBOP,
		SubdivisionNZCAN,
		SubdivisionNZCIT,
		SubdivisionNZGIS,
		SubdivisionNZHKB,
		SubdivisionNZMBH,
		SubdivisionNZMWT,
		SubdivisionNZN,
		SubdivisionNZNSN,
		SubdivisionNZNTL,
		SubdivisionNZOTA,
		SubdivisionNZS,
		SubdivisionNZSTL,
		SubdivisionNZTAS,
		SubdivisionNZTKI,
		SubdivisionNZWGN,
		SubdivisionNZWKO,
		SubdivisionNZWTC,
		SubdivisionOMBA,
		SubdivisionOMBU,
		SubdivisionOMDA,
		SubdivisionOMMA,
		SubdivisionOMMU,
		SubdivisionOMSH,
		SubdivisionOMWU,
		SubdivisionOMZA,
		SubdivisionOMZU,
		SubdivisionPA1,
		SubdivisionPA2,
		SubdivisionPA3,
		SubdivisionPA4,
		SubdivisionPA5,
		SubdivisionPA6,
		SubdivisionPA7,
		SubdivisionPA8,
		SubdivisionPA9,
		SubdivisionPAEM,
		SubdivisionPAKY,
		SubdivisionPANB,
		SubdivisionPEAMA,
		SubdivisionPEANC,
		SubdivisionPEAPU,
		SubdivisionPEARE,
		SubdivisionPEAYA,
		SubdivisionPECAJ,
		SubdivisionPECAL,
		SubdivisionPECUS,
		SubdivisionPEHUC,
		SubdivisionPEHUV,
		SubdivisionPEICA,
		SubdivisionPEJUN,
		SubdivisionPELAL,
		SubdivisionPELAM,
		SubdivisionPELIM,
		SubdivisionPELMA,
		SubdivisionPELOR,
		SubdivisionPEMDD,
		SubdivisionPEMOQ,
		SubdivisionPEPAS,
		SubdivisionPEPIU,
		SubdivisionPEPUN,
		SubdivisionPESAM,
		SubdivisionPETAC,
		SubdivisionPETUM,
		SubdivisionPEUCA,
		SubdivisionPGCPK,
		SubdivisionPGCPM,
		SubdivisionPGEBR,
		SubdivisionPGEHG,
		SubdivisionPGEPW,
		SubdivisionPGESW,
		SubdivisionPGGPK,
		SubdivisionPGMBA,
		SubdivisionPGMPL,
		SubdivisionPGMPM,
		SubdivisionPGMRL,
		SubdivisionPGNCD,
		SubdivisionPGNIK,
		SubdivisionPGNPP,
		SubdivisionPGNSB,
		SubdivisionPGSAN,
		SubdivisionPGSHM,
		SubdivisionPGWBK,
		SubdivisionPGWHM,
		SubdivisionPGWPD,
		SubdivisionPH00,
		SubdivisionPH01,
		SubdivisionPH02,
		SubdivisionPH03,
		SubdivisionPH05,
		SubdivisionPH06,
		SubdivisionPH07,
		SubdivisionPH08,
		SubdivisionPH09,
		SubdivisionPH10,
		SubdivisionPH11,
		SubdivisionPH12,
		SubdivisionPH13,
		SubdivisionPH14,
		SubdivisionPH15,
		SubdivisionPH40,
		SubdivisionPH41,
		SubdivisionPHABR,
		SubdivisionPHAGN,
		SubdivisionPHAGS,
		SubdivisionPHAKL,
		SubdivisionPHALB,
		SubdivisionPHANT,
		SubdivisionPHAPA,
		SubdivisionPHAUR,
		SubdivisionPHBAN,
		SubdivisionPHBAS,
		SubdivisionPHBEN,
		SubdivisionPHBIL,
		SubdivisionPHBOH,
		SubdivisionPHBTG,
		SubdivisionPHBTN,
		SubdivisionPHBUK,
		SubdivisionPHBUL,
		SubdivisionPHCAG,
		SubdivisionPHCAM,
		SubdivisionPHCAN,
		SubdivisionPHCAP,
		SubdivisionPHCAS,
		SubdivisionPHCAT,
		SubdivisionPHCAV,
		SubdivisionPHCEB,
		SubdivisionPHCOM,
		SubdivisionPHDAO,
		SubdivisionPHDAS,
		SubdivisionPHDAV,
		SubdivisionPHDIN,
		SubdivisionPHEAS,
		SubdivisionPHGUI,
		SubdivisionPHIFU,
		SubdivisionPHILI,
		SubdivisionPHILN,
		SubdivisionPHILS,
		SubdivisionPHISA,
		SubdivisionPHKAL,
		SubdivisionPHLAG,
		SubdivisionPHLAN,
		SubdivisionPHLAS,
		SubdivisionPHLEY,
		SubdivisionPHLUN,
		SubdivisionPHMAD,
		SubdivisionPHMAG,
		SubdivisionPHMAS,
		SubdivisionPHMDC,
		SubdivisionPHMDR,
		SubdivisionPHMOU,
		SubdivisionPHMSC,
		SubdivisionPHMSR,
		SubdivisionPHNCO,
		SubdivisionPHNEC,
		SubdivisionPHNER,
		SubdivisionPHNSA,
		SubdivisionPHNUE,
		SubdivisionPHNUV,
		SubdivisionPHPAM,
		SubdivisionPHPAN,
		SubdivisionPHPLW,
		SubdivisionPHQUE,
		SubdivisionPHQUI,
		SubdivisionPHRIZ,
		SubdivisionPHROM,
		SubdivisionPHSAR,
		SubdivisionPHSCO,
		SubdivisionPHSIG,
		SubdivisionPHSLE,
		SubdivisionPHSLU,
		SubdivisionPHSOR,
		SubdivisionPHSUK,
		SubdivisionPHSUN,
		SubdivisionPHSUR,
		SubdivisionPHTAR,
		SubdivisionPHTAW,
		SubdivisionPHWSA,
		SubdivisionPHZAN,
		SubdivisionPHZAS,
		SubdivisionPHZMB,
		SubdivisionPHZSI,
		SubdivisionPKBA,
		SubdivisionPKGB,
		SubdivisionPKIS,
		SubdivisionPKJK,
		SubdivisionPKKP,
		SubdivisionPKPB,
		SubdivisionPKSD,
		SubdivisionPKTA,
		SubdivisionPLDS,
		SubdivisionPLKP,
		SubdivisionPLLB,
		SubdivisionPLLD,
		SubdivisionPLLU,
		SubdivisionPLMA,
		SubdivisionPLMZ,
		SubdivisionPLOP,
		SubdivisionPLPD,
		SubdivisionPLPK,
		SubdivisionPLPM,
		SubdivisionPLSK,
		SubdivisionPLSL,
		SubdivisionPLWN,
		SubdivisionPLWP,
		SubdivisionPLZP,
		SubdivisionPSBTH,
		SubdivisionPSDEB,
		SubdivisionPSGZA,
		SubdivisionPSHBN,
		SubdivisionPSJEM,
		SubdivisionPSJEN,
		SubdivisionPSJRH,
		SubdivisionPSKYS,
		SubdivisionPSNBS,
		SubdivisionPSNGZ,
		SubdivisionPSQQA,
		SubdivisionPSRBH,
		SubdivisionPSRFH,
		SubdivisionPSSLT,
		SubdivisionPSTBS,
		SubdivisionPSTKM,
		SubdivisionPT01,
		SubdivisionPT02,
		SubdivisionPT03,
		SubdivisionPT04,
		SubdivisionPT05,
		SubdivisionPT06,
		SubdivisionPT07,
		SubdivisionPT08,
		SubdivisionPT09,
		SubdivisionPT10,
		SubdivisionPT11,
		SubdivisionPT12,
		SubdivisionPT13,
		SubdivisionPT14,
		SubdivisionPT15,
		SubdivisionPT16,
		SubdivisionPT17,
		SubdivisionPT18,
		SubdivisionPT20,
		SubdivisionPT30,
		SubdivisionPW002,
		SubdivisionPW004,
		SubdivisionPW010,
		SubdivisionPW050,
		SubdivisionPW100,
		SubdivisionPW150,
		SubdivisionPW212,
		SubdivisionPW214,
		SubdivisionPW218,
		SubdivisionPW222,
		SubdivisionPW224,
		SubdivisionPW226,
		SubdivisionPW227,
		SubdivisionPW228,
		SubdivisionPW350,
		SubdivisionPW370,
		SubdivisionPY1,
		SubdivisionPY10,
		SubdivisionPY11,
		SubdivisionPY12,
		SubdivisionPY13,
		SubdivisionPY14,
		SubdivisionPY15,
		SubdivisionPY16,
		SubdivisionPY19,
		SubdivisionPY2,
		SubdivisionPY3,
		SubdivisionPY4,
		SubdivisionPY5,
		SubdivisionPY6,
		SubdivisionPY7,
		SubdivisionPY8,
		SubdivisionPY9,
		SubdivisionPYASU,
		SubdivisionQADA,
		SubdivisionQAKH,
		SubdivisionQAMS,
		SubdivisionQARA,
		SubdivisionQAUS,
		SubdivisionQAWA,
		SubdivisionQAZA,
		SubdivisionROAB,
		SubdivisionROAG,
		SubdivisionROAR,
		SubdivisionROB,
		SubdivisionROBC,
		SubdivisionROBH,
		SubdivisionROBN,
		SubdivisionROBR,
		SubdivisionROBT,
		SubdivisionROBV,
		SubdivisionROBZ,
		SubdivisionROCJ,
		SubdivisionROCL,
		SubdivisionROCS,
		SubdivisionROCT,
		SubdivisionROCV,
		SubdivisionRODB,
		SubdivisionRODJ,
		SubdivisionROGJ,
		SubdivisionROGL,
		SubdivisionROGR,
		SubdivisionROHD,
		SubdivisionROHR,
		SubdivisionROIF,
		SubdivisionROIL,
		SubdivisionROIS,
		SubdivisionROMH,
		SubdivisionROMM,
		SubdivisionROMS,
		SubdivisionRONT,
		SubdivisionROOT,
		SubdivisionROPH,
		SubdivisionROSB,
		SubdivisionROSJ,
		SubdivisionROSM,
		SubdivisionROSV,
		SubdivisionROTL,
		SubdivisionROTM,
		SubdivisionROTR,
		SubdivisionROVL,
		SubdivisionROVN,
		SubdivisionROVS,
		SubdivisionRS00,
		SubdivisionRS01,
		SubdivisionRS02,
		SubdivisionRS03,
		SubdivisionRS04,
		SubdivisionRS05,
		SubdivisionRS06,
		SubdivisionRS07,
		SubdivisionRS08,
		SubdivisionRS09,
		SubdivisionRS10,
		SubdivisionRS11,
		SubdivisionRS12,
		SubdivisionRS13,
		SubdivisionRS14,
		SubdivisionRS15,
		SubdivisionRS16,
		SubdivisionRS17,
		SubdivisionRS18,
		SubdivisionRS19,
		SubdivisionRS20,
		SubdivisionRS21,
		SubdivisionRS22,
		SubdivisionRS23,
		SubdivisionRS24,
		SubdivisionRS25,
		SubdivisionRS26,
		SubdivisionRS27,
		SubdivisionRS28,
		SubdivisionRS29,
		SubdivisionRSKM,
		SubdivisionRSVO,
		SubdivisionRUAD,
		SubdivisionRUAL,
		SubdivisionRUALT,
		SubdivisionRUAMU,
		SubdivisionRUARK,
		SubdivisionRUAST,
		SubdivisionRUBA,
		SubdivisionRUBEL,
		SubdivisionRUBRY,
		SubdivisionRUBU,
		SubdivisionRUCE,
		SubdivisionRUCHE,
		SubdivisionRUCHU,
		SubdivisionRUCU,
		SubdivisionRUDA,
		SubdivisionRUIN,
		SubdivisionRUIRK,
		SubdivisionRUIVA,
		SubdivisionRUKAM,
		SubdivisionRUKB,
		SubdivisionRUKC,
		SubdivisionRUKDA,
		SubdivisionRUKEM,
		SubdivisionRUKGD,
		SubdivisionRUKGN,
		SubdivisionRUKHA,
		SubdivisionRUKHM,
		SubdivisionRUKIR,
		SubdivisionRUKK,
		SubdivisionRUKL,
		SubdivisionRUKLU,
		SubdivisionRUKO,
		SubdivisionRUKOS,
		SubdivisionRUKR,
		SubdivisionRUKRS,
		SubdivisionRUKYA,
		SubdivisionRULEN,
		SubdivisionRULIP,
		SubdivisionRUMAG,
		SubdivisionRUME,
		SubdivisionRUMO,
		SubdivisionRUMOS,
		SubdivisionRUMOW,
		SubdivisionRUMUR,
		SubdivisionRUNEN,
		SubdivisionRUNGR,
		SubdivisionRUNIZ,
		SubdivisionRUNVS,
		SubdivisionRUOMS,
		SubdivisionRUORE,
		SubdivisionRUORL,
		SubdivisionRUPER,
		SubdivisionRUPNZ,
		SubdivisionRUPRI,
		SubdivisionRUPSK,
		SubdivisionRUROS,
		SubdivisionRURYA,
		SubdivisionRUSA,
		SubdivisionRUSAK,
		SubdivisionRUSAM,
		SubdivisionRUSAR,
		SubdivisionRUSE,
		SubdivisionRUSMO,
		SubdivisionRUSPE,
		SubdivisionRUSTA,
		SubdivisionRUSVE,
		SubdivisionRUTA,
		SubdivisionRUTAM,
		SubdivisionRUTOM,
		SubdivisionRUTUL,
		SubdivisionRUTVE,
		SubdivisionRUTY,
		SubdivisionRUTYU,
		SubdivisionRUUD,
		SubdivisionRUULY,
		SubdivisionRUVGG,
		SubdivisionRUVLA,
		SubdivisionRUVLG,
		SubdivisionRUVOR,
		SubdivisionRUYAN,
		SubdivisionRUYAR,
		SubdivisionRUYEV,
		SubdivisionRUZAB,
		SubdivisionRW01,
		SubdivisionRW02,
		SubdivisionRW03,
		SubdivisionRW04,
		SubdivisionRW05,
		SubdivisionSA01,
		SubdivisionSA02,
		SubdivisionSA03,
		SubdivisionSA04,
		SubdivisionSA05,
		SubdivisionSA06,
		SubdivisionSA07,
		SubdivisionSA08,
		SubdivisionSA09,
		SubdivisionSA10,
		SubdivisionSA11,
		SubdivisionSA12,
		SubdivisionSA14,
		SubdivisionSBCE,
		SubdivisionSBCH,
		SubdivisionSBCT,
		SubdivisionSBGU,
		SubdivisionSBIS,
		SubdivisionSBMK,
		SubdivisionSBML,
		SubdivisionSBRB,
		SubdivisionSBTE,
		SubdivisionSBWE,
		SubdivisionSC01,
		SubdivisionSC02,
		SubdivisionSC03,
		SubdivisionSC04,
		SubdivisionSC05,
		SubdivisionSC06,
		SubdivisionSC07,
		SubdivisionSC08,
		SubdivisionSC09,
		SubdivisionSC10,
		SubdivisionSC11,
		SubdivisionSC12,
		SubdivisionSC13,
		SubdivisionSC14,
		SubdivisionSC15,
		SubdivisionSC16,
		SubdivisionSC17,
		SubdivisionSC18,
		SubdivisionSC19,
		SubdivisionSC20,
		SubdivisionSC21,
		SubdivisionSC22,
		SubdivisionSC23,
		SubdivisionSC24,
		SubdivisionSC25,
		SubdivisionSDDC,
		SubdivisionSDDE,
		SubdivisionSDDN,
		SubdivisionSDDS,
		SubdivisionSDDW,
		SubdivisionSDGD,
		SubdivisionSDGZ,
		SubdivisionSDKA,
		SubdivisionSDKH,
		SubdivisionSDKN,
		SubdivisionSDKS,
		SubdivisionSDNB,
		SubdivisionSDNO,
		SubdivisionSDNR,
		SubdivisionSDNW,
		SubdivisionSDRS,
		SubdivisionSDSI,
		SubdivisionSEAB,
		SubdivisionSEAC,
		SubdivisionSEBD,
		SubdivisionSEC,
		SubdivisionSED,
		SubdivisionSEE,
		SubdivisionSEF,
		SubdivisionSEG,
		SubdivisionSEH,
		SubdivisionSEI,
		SubdivisionSEK,
		SubdivisionSEM,
		SubdivisionSEN,
		SubdivisionSEO,
		SubdivisionSES,
		SubdivisionSET,
		SubdivisionSEU,
		SubdivisionSEW,
		SubdivisionSEX,
		SubdivisionSEY,
		SubdivisionSEZ,
		SubdivisionSG01,
		SubdivisionSG02,
		SubdivisionSG03,
		SubdivisionSG04,
		SubdivisionSG05,
		SubdivisionSHAC,
		SubdivisionSHHL,
		SubdivisionSHTA,
		SubdivisionSI001,
		SubdivisionSI002,
		SubdivisionSI003,
		SubdivisionSI004,
		SubdivisionSI005,
		SubdivisionSI006,
		SubdivisionSI007,
		SubdivisionSI008,
		SubdivisionSI009,
		SubdivisionSI010,
		SubdivisionSI011,
		SubdivisionSI012,
		SubdivisionSI013,
		SubdivisionSI014,
		SubdivisionSI015,
		SubdivisionSI016,
		SubdivisionSI017,
		SubdivisionSI018,
		SubdivisionSI019,
		SubdivisionSI020,
		SubdivisionSI021,
		SubdivisionSI022,
		SubdivisionSI023,
		SubdivisionSI024,
		SubdivisionSI025,
		SubdivisionSI026,
		SubdivisionSI027,
		SubdivisionSI028,
		SubdivisionSI029,
		SubdivisionSI030,
		SubdivisionSI031,
		SubdivisionSI032,
		SubdivisionSI033,
		SubdivisionSI034,
		SubdivisionSI035,
		SubdivisionSI036,
		SubdivisionSI037,
		SubdivisionSI038,
		SubdivisionSI039,
		SubdivisionSI040,
		SubdivisionSI041,
		SubdivisionSI042,
		SubdivisionSI043,
		SubdivisionSI044,
		SubdivisionSI045,
		SubdivisionSI046,
		SubdivisionSI047,
		SubdivisionSI048,
		SubdivisionSI049,
		SubdivisionSI050,
		SubdivisionSI051,
		SubdivisionSI052,
		SubdivisionSI053,
		SubdivisionSI054,
		SubdivisionSI055,
		SubdivisionSI056,
		SubdivisionSI057,
		SubdivisionSI058,
		SubdivisionSI059,
		SubdivisionSI060,
		SubdivisionSI061,
		SubdivisionSI062,
		SubdivisionSI063,
		SubdivisionSI064,
		SubdivisionSI065,
		SubdivisionSI066,
		SubdivisionSI067,
		SubdivisionSI068,
		SubdivisionSI069,
		SubdivisionSI070,
		SubdivisionSI071,
		SubdivisionSI072,
		SubdivisionSI073,
		SubdivisionSI074,
		SubdivisionSI075,
		SubdivisionSI076,
		SubdivisionSI077,
		SubdivisionSI078,
		SubdivisionSI079,
		SubdivisionSI080,
		SubdivisionSI081,
		SubdivisionSI082,
		SubdivisionSI083,
		SubdivisionSI084,
		SubdivisionSI085,
		SubdivisionSI086,
		SubdivisionSI087,
		SubdivisionSI088,
		SubdivisionSI089,
		SubdivisionSI090,
		SubdivisionSI091,
		SubdivisionSI092,
		SubdivisionSI093,
		SubdivisionSI094,
		SubdivisionSI095,
		SubdivisionSI096,
		SubdivisionSI097,
		SubdivisionSI098,
		SubdivisionSI099,
		SubdivisionSI100,
		SubdivisionSI101,
		SubdivisionSI102,
		SubdivisionSI103,
		SubdivisionSI104,
		SubdivisionSI105,
		SubdivisionSI106,
		SubdivisionSI107,
		SubdivisionSI108,
		SubdivisionSI109,
		SubdivisionSI110,
		SubdivisionSI111,
		SubdivisionSI112,
		SubdivisionSI113,
		SubdivisionSI114,
		SubdivisionSI115,
		SubdivisionSI116,
		SubdivisionSI117,
		SubdivisionSI118,
		SubdivisionSI119,
		SubdivisionSI120,
		SubdivisionSI121,
		SubdivisionSI122,
		SubdivisionSI123,
		SubdivisionSI124,
		SubdivisionSI125,
		SubdivisionSI126,
		SubdivisionSI127,
		SubdivisionSI128,
		SubdivisionSI129,
		SubdivisionSI130,
		SubdivisionSI131,
		SubdivisionSI132,
		SubdivisionSI133,
		SubdivisionSI134,
		SubdivisionSI135,
		SubdivisionSI136,
		SubdivisionSI137,
		SubdivisionSI138,
		SubdivisionSI139,
		SubdivisionSI140,
		SubdivisionSI141,
		SubdivisionSI142,
		SubdivisionSI143,
		SubdivisionSI144,
		SubdivisionSI146,
		SubdivisionSI147,
		SubdivisionSI148,
		SubdivisionSI149,
		SubdivisionSI150,
		SubdivisionSI151,
		SubdivisionSI152,
		SubdivisionSI153,
		SubdivisionSI154,
		SubdivisionSI155,
		SubdivisionSI156,
		SubdivisionSI157,
		SubdivisionSI158,
		SubdivisionSI159,
		SubdivisionSI160,
		SubdivisionSI161,
		SubdivisionSI162,
		SubdivisionSI163,
		SubdivisionSI164,
		SubdivisionSI165,
		SubdivisionSI166,
		SubdivisionSI167,
		SubdivisionSI168,
		SubdivisionSI169,
		SubdivisionSI170,
		SubdivisionSI171,
		SubdivisionSI172,
		SubdivisionSI173,
		SubdivisionSI174,
		SubdivisionSI175,
		SubdivisionSI176,
		SubdivisionSI177,
		SubdivisionSI178,
		SubdivisionSI179,
		SubdivisionSI180,
		SubdivisionSI181,
		SubdivisionSI182,
		SubdivisionSI183,
		SubdivisionSI184,
		SubdivisionSI185,
		SubdivisionSI186,
		SubdivisionSI187,
		SubdivisionSI188,
		SubdivisionSI189,
		SubdivisionSI190,
		SubdivisionSI191,
		SubdivisionSI192,
		SubdivisionSI193,
		SubdivisionSI194,
		SubdivisionSI195,
		SubdivisionSI196,
		SubdivisionSI197,
		SubdivisionSI198,
		SubdivisionSI199,
		SubdivisionSI200,
		SubdivisionSI201,
		SubdivisionSI202,
		SubdivisionSI203,
		SubdivisionSI204,
		SubdivisionSI205,
		SubdivisionSI206,
		SubdivisionSI207,
		SubdivisionSI208,
		SubdivisionSI209,
		SubdivisionSI210,
		SubdivisionSI211,
		SubdivisionSKBC,
		SubdivisionSKBL,
		SubdivisionSKKI,
		SubdivisionSKNI,
		SubdivisionSKPV,
		SubdivisionSKTA,
		SubdivisionSKTC,
		SubdivisionSKZI,
		SubdivisionSLE,
		SubdivisionSLN,
		SubdivisionSLS,
		SubdivisionSLW,
		SubdivisionSM01,
		SubdivisionSM02,
		SubdivisionSM03,
		SubdivisionSM04,
		SubdivisionSM05,
		SubdivisionSM06,
		SubdivisionSM07,
		SubdivisionSM08,
		SubdivisionSM09,
		SubdivisionSNDB,
		SubdivisionSNDK,
		SubdivisionSNFK,
		SubdivisionSNKA,
		SubdivisionSNKD,
		SubdivisionSNKE,
		SubdivisionSNKL,
		SubdivisionSNLG,
		SubdivisionSNMT,
		SubdivisionSNSE,
		SubdivisionSNSL,
		SubdivisionSNTC,
		SubdivisionSNTH,
		SubdivisionSNZG,
		SubdivisionSOAW,
		SubdivisionSOBK,
		SubdivisionSOBN,
		SubdivisionSOBR,
		SubdivisionSOBY,
		SubdivisionSOGA,
		SubdivisionSOGE,
		SubdivisionSOHI,
		SubdivisionSOJD,
		SubdivisionSOJH,
		SubdivisionSOMU,
		SubdivisionSONU,
		SubdivisionSOSA,
		SubdivisionSOSD,
		SubdivisionSOSH,
		SubdivisionSOSO,
		SubdivisionSOTO,
		SubdivisionSOWO,
		SubdivisionSRBR,
		SubdivisionSRCM,
		SubdivisionSRCR,
		SubdivisionSRMA,
		SubdivisionSRNI,
		SubdivisionSRPM,
		SubdivisionSRPR,
		SubdivisionSRSA,
		SubdivisionSRSI,
		SubdivisionSRWA,
		SubdivisionSSBN,
		SubdivisionSSBW,
		SubdivisionSSEC,
		SubdivisionSSEE,
		SubdivisionSSEW,
		SubdivisionSSJG,
		SubdivisionSSLK,
		SubdivisionSSNU,
		SubdivisionSSUY,
		SubdivisionSSWR,
		SubdivisionSTP,
		SubdivisionSTS,
		SubdivisionSVAH,
		SubdivisionSVCA,
		SubdivisionSVCH,
		SubdivisionSVCU,
		SubdivisionSVLI,
		SubdivisionSVMO,
		SubdivisionSVPA,
		SubdivisionSVSA,
		SubdivisionSVSM,
		SubdivisionSVSO,
		SubdivisionSVSS,
		SubdivisionSVSV,
		SubdivisionSVUN,
		SubdivisionSVUS,
		SubdivisionSYDI,
		SubdivisionSYDR,
		SubdivisionSYDY,
		SubdivisionSYHA,
		SubdivisionSYHI,
		SubdivisionSYHL,
		SubdivisionSYHM,
		SubdivisionSYID,
		SubdivisionSYLA,
		SubdivisionSYQU,
		SubdivisionSYRA,
		SubdivisionSYRD,
		SubdivisionSYSU,
		SubdivisionSYTA,
		SubdivisionSZHH,
		SubdivisionSZLU,
		SubdivisionSZMA,
		SubdivisionSZSH,
		SubdivisionTDBA,
		SubdivisionTDBG,
		SubdivisionTDBO,
		SubdivisionTDCB,
		SubdivisionTDEN,
		SubdivisionTDGR,
		SubdivisionTDHL,
		SubdivisionTDKA,
		SubdivisionTDLC,
		SubdivisionTDLO,
		SubdivisionTDLR,
		SubdivisionTDMA,
		SubdivisionTDMC,
		SubdivisionTDME,
		SubdivisionTDMO,
		SubdivisionTDND,
		SubdivisionTDOD,
		SubdivisionTDSA,
		SubdivisionTDSI,
		SubdivisionTDTA,
		SubdivisionTDTI,
		SubdivisionTDWF,
		SubdivisionTGC,
		SubdivisionTGK,
		SubdivisionTGM,
		SubdivisionTGP,
		SubdivisionTGS,
		SubdivisionTH10,
		SubdivisionTH11,
		SubdivisionTH12,
		SubdivisionTH13,
		SubdivisionTH14,
		SubdivisionTH15,
		SubdivisionTH16,
		SubdivisionTH17,
		SubdivisionTH18,
		SubdivisionTH19,
		SubdivisionTH20,
		SubdivisionTH21,
		SubdivisionTH22,
		SubdivisionTH23,
		SubdivisionTH24,
		SubdivisionTH25,
		SubdivisionTH26,
		SubdivisionTH27,
		SubdivisionTH30,
		SubdivisionTH31,
		SubdivisionTH32,
		SubdivisionTH33,
		SubdivisionTH34,
		SubdivisionTH35,
		SubdivisionTH36,
		SubdivisionTH37,
		SubdivisionTH39,
		SubdivisionTH40,
		SubdivisionTH41,
		SubdivisionTH42,
		SubdivisionTH43,
		SubdivisionTH44,
		SubdivisionTH45,
		SubdivisionTH46,
		SubdivisionTH47,
		SubdivisionTH48,
		SubdivisionTH49,
		SubdivisionTH50,
		SubdivisionTH51,
		SubdivisionTH52,
		SubdivisionTH53,
		SubdivisionTH54,
		SubdivisionTH55,
		SubdivisionTH56,
		SubdivisionTH57,
		SubdivisionTH58,
		SubdivisionTH60,
		SubdivisionTH61,
		SubdivisionTH62,
		SubdivisionTH63,
		SubdivisionTH64,
		SubdivisionTH65,
		SubdivisionTH66,
		SubdivisionTH67,
		SubdivisionTH70,
		SubdivisionTH71,
		SubdivisionTH72,
		SubdivisionTH73,
		SubdivisionTH74,
		SubdivisionTH75,
		SubdivisionTH76,
		SubdivisionTH77,
		SubdivisionTH80,
		SubdivisionTH81,
		SubdivisionTH82,
		SubdivisionTH83,
		SubdivisionTH84,
		SubdivisionTH85,
		SubdivisionTH86,
		SubdivisionTH90,
		SubdivisionTH91,
		SubdivisionTH92,
		SubdivisionTH93,
		SubdivisionTH94,
		SubdivisionTH95,
		SubdivisionTH96,
		SubdivisionTHS,
		SubdivisionTJGB,
		SubdivisionTJKT,
		SubdivisionTJSU,
		SubdivisionTLAL,
		SubdivisionTLAN,
		SubdivisionTLBA,
		SubdivisionTLBO,
		SubdivisionTLCO,
		SubdivisionTLDI,
		SubdivisionTLER,
		SubdivisionTLLA,
		SubdivisionTLLI,
		SubdivisionTLMF,
		SubdivisionTLMT,
		SubdivisionTLOE,
		SubdivisionTLVI,
		SubdivisionTMA,
		SubdivisionTMB,
		SubdivisionTMD,
		SubdivisionTML,
		SubdivisionTMM,
		SubdivisionTMS,
		SubdivisionTN11,
		SubdivisionTN12,
		SubdivisionTN13,
		SubdivisionTN14,
		SubdivisionTN21,
		SubdivisionTN22,
		SubdivisionTN23,
		SubdivisionTN31,
		SubdivisionTN32,
		SubdivisionTN33,
		SubdivisionTN34,
		SubdivisionTN41,
		SubdivisionTN42,
		SubdivisionTN43,
		SubdivisionTN51,
		SubdivisionTN52,
		SubdivisionTN53,
		SubdivisionTN61,
		SubdivisionTN71,
		SubdivisionTN72,
		SubdivisionTN73,
		SubdivisionTN81,
		SubdivisionTN82,
		SubdivisionTN83,
		SubdivisionTO01,
		SubdivisionTO02,
		SubdivisionTO03,
		SubdivisionTO04,
		SubdivisionTO05,
		SubdivisionTR01,
		SubdivisionTR02,
		SubdivisionTR03,
		SubdivisionTR04,
		SubdivisionTR05,
		SubdivisionTR06,
		SubdivisionTR07,
		SubdivisionTR08,
		SubdivisionTR09,
		SubdivisionTR10,
		SubdivisionTR11,
		SubdivisionTR12,
		SubdivisionTR13,
		SubdivisionTR14,
		SubdivisionTR15,
		SubdivisionTR16,
		SubdivisionTR17,
		SubdivisionTR18,
		SubdivisionTR19,
		SubdivisionTR20,
		SubdivisionTR21,
		SubdivisionTR22,
		SubdivisionTR23,
		SubdivisionTR24,
		SubdivisionTR25,
		SubdivisionTR26,
		SubdivisionTR27,
		SubdivisionTR28,
		SubdivisionTR29,
		SubdivisionTR30,
		SubdivisionTR31,
		SubdivisionTR32,
		SubdivisionTR33,
		SubdivisionTR34,
		SubdivisionTR35,
		SubdivisionTR36,
		SubdivisionTR37,
		SubdivisionTR38,
		SubdivisionTR39,
		SubdivisionTR40,
		SubdivisionTR41,
		SubdivisionTR42,
		SubdivisionTR43,
		SubdivisionTR44,
		SubdivisionTR45,
		SubdivisionTR46,
		SubdivisionTR47,
		SubdivisionTR48,
		SubdivisionTR49,
		SubdivisionTR50,
		SubdivisionTR51,
		SubdivisionTR52,
		SubdivisionTR53,
		SubdivisionTR54,
		SubdivisionTR55,
		SubdivisionTR56,
		SubdivisionTR57,
		SubdivisionTR58,
		SubdivisionTR59,
		SubdivisionTR60,
		SubdivisionTR61,
		SubdivisionTR62,
		SubdivisionTR63,
		SubdivisionTR64,
		SubdivisionTR65,
		SubdivisionTR66,
		SubdivisionTR67,
		SubdivisionTR68,
		SubdivisionTR69,
		SubdivisionTR70,
		SubdivisionTR71,
		SubdivisionTR72,
		SubdivisionTR73,
		SubdivisionTR74,
		SubdivisionTR75,
		SubdivisionTR76,
		SubdivisionTR77,
		SubdivisionTR78,
		SubdivisionTR79,
		SubdivisionTR80,
		SubdivisionTR81,
		SubdivisionTTARI,
		SubdivisionTTCHA,
		SubdivisionTTCTT,
		SubdivisionTTDMN,
		SubdivisionTTETO,
		SubdivisionTTPED,
		SubdivisionTTPOS,
		SubdivisionTTPRT,
		SubdivisionTTPTF,
		SubdivisionTTRCM,
		SubdivisionTTSFO,
		SubdivisionTTSGE,
		SubdivisionTTSIP,
		SubdivisionTTSJL,
		SubdivisionTTTUP,
		SubdivisionTTWTO,
		SubdivisionTVFUN,
		SubdivisionTVNIT,
		SubdivisionTVNKF,
		SubdivisionTVNKL,
		SubdivisionTVNMA,
		SubdivisionTVNMG,
		SubdivisionTVNUI,
		SubdivisionTVVAI,
		SubdivisionTWCHA,
		SubdivisionTWCYI,
		SubdivisionTWCYQ,
		SubdivisionTWHSQ,
		SubdivisionTWHSZ,
		SubdivisionTWHUA,
		SubdivisionTWILA,
		SubdivisionTWKEE,
		SubdivisionTWKHH,
		SubdivisionTWKHQ,
		SubdivisionTWMIA,
		SubdivisionTWNAN,
		SubdivisionTWPEN,
		SubdivisionTWPIF,
		SubdivisionTWTAO,
		SubdivisionTWTNN,
		SubdivisionTWTNQ,
		SubdivisionTWTPE,
		SubdivisionTWTPQ,
		SubdivisionTWTTT,
		SubdivisionTWTXG,
		SubdivisionTWTXQ,
		SubdivisionTWYUN,
		SubdivisionTZ01,
		SubdivisionTZ02,
		SubdivisionTZ03,
		SubdivisionTZ04,
		SubdivisionTZ05,
		SubdivisionTZ06,
		SubdivisionTZ07,
		SubdivisionTZ08,
		SubdivisionTZ09,
		SubdivisionTZ10,
		SubdivisionTZ11,
		SubdivisionTZ12,
		SubdivisionTZ13,
		SubdivisionTZ14,
		SubdivisionTZ15,
		SubdivisionTZ16,
		SubdivisionTZ17,
		SubdivisionTZ18,
		SubdivisionTZ19,
		SubdivisionTZ20,
		SubdivisionTZ21,
		SubdivisionTZ22,
		SubdivisionTZ23,
		SubdivisionTZ24,
		SubdivisionTZ25,
		SubdivisionTZ26,
		SubdivisionUA05,
		SubdivisionUA07,
		SubdivisionUA09,
		SubdivisionUA12,
		SubdivisionUA14,
		SubdivisionUA18,
		SubdivisionUA21,
		SubdivisionUA23,
		SubdivisionUA26,
		SubdivisionUA30,
		SubdivisionUA32,
		SubdivisionUA35,
		SubdivisionUA40,
		SubdivisionUA43,
		SubdivisionUA46,
		SubdivisionUA48,
		SubdivisionUA51,
		SubdivisionUA53,
		SubdivisionUA56,
		SubdivisionUA59,
		SubdivisionUA61,
		SubdivisionUA63,
		SubdivisionUA65,
		SubdivisionUA68,
		SubdivisionUA71,
		SubdivisionUA74,
		SubdivisionUA77,
		SubdivisionUG101,
		SubdivisionUG102,
		SubdivisionUG103,
		SubdivisionUG104,
		SubdivisionUG105,
		SubdivisionUG106,
		SubdivisionUG107,
		SubdivisionUG108,
		SubdivisionUG109,
		SubdivisionUG110,
		SubdivisionUG111,
		SubdivisionUG112,
		SubdivisionUG113,
		SubdivisionUG114,
		SubdivisionUG115,
		SubdivisionUG116,
		SubdivisionUG201,
		SubdivisionUG202,
		SubdivisionUG203,
		SubdivisionUG204,
		SubdivisionUG205,
		SubdivisionUG206,
		SubdivisionUG207,
		SubdivisionUG208,
		SubdivisionUG209,
		SubdivisionUG210,
		SubdivisionUG211,
		SubdivisionUG212,
		SubdivisionUG213,
		SubdivisionUG214,
		SubdivisionUG215,
		SubdivisionUG216,
		SubdivisionUG217,
		SubdivisionUG218,
		SubdivisionUG219,
		SubdivisionUG220,
		SubdivisionUG221,
		SubdivisionUG222,
		SubdivisionUG223,
		SubdivisionUG224,
		SubdivisionUG301,
		SubdivisionUG302,
		SubdivisionUG303,
		SubdivisionUG304,
		SubdivisionUG305,
		SubdivisionUG306,
		SubdivisionUG307,
		SubdivisionUG308,
		SubdivisionUG309,
		SubdivisionUG310,
		SubdivisionUG311,
		SubdivisionUG312,
		SubdivisionUG313,
		SubdivisionUG314,
		SubdivisionUG315,
		SubdivisionUG316,
		SubdivisionUG317,
		SubdivisionUG318,
		SubdivisionUG319,
		SubdivisionUG320,
		SubdivisionUG321,
		SubdivisionUG401,
		SubdivisionUG402,
		SubdivisionUG403,
		SubdivisionUG404,
		SubdivisionUG405,
		SubdivisionUG406,
		SubdivisionUG407,
		SubdivisionUG408,
		SubdivisionUG409,
		SubdivisionUG410,
		SubdivisionUG411,
		SubdivisionUG412,
		SubdivisionUG413,
		SubdivisionUG414,
		SubdivisionUG415,
		SubdivisionUG416,
		SubdivisionUG417,
		SubdivisionUG418,
		SubdivisionUG419,
		SubdivisionUGC,
		SubdivisionUGE,
		SubdivisionUGN,
		SubdivisionUGW,
		SubdivisionUM67,
		SubdivisionUM71,
		SubdivisionUM76,
		SubdivisionUM79,
		SubdivisionUM81,
		SubdivisionUM84,
		SubdivisionUM86,
		SubdivisionUM89,
		SubdivisionUM95,
		SubdivisionUSAK,
		SubdivisionUSAL,
		SubdivisionUSAR,
		SubdivisionUSAS,
		SubdivisionUSAZ,
		SubdivisionUSCA,
		SubdivisionUSCO,
		SubdivisionUSCT,
		SubdivisionUSDC,
		SubdivisionUSDE,
		SubdivisionUSFL,
		SubdivisionUSGA,
		SubdivisionUSGU,
		SubdivisionUSHI,
		SubdivisionUSIA,
		SubdivisionUSID,
		SubdivisionUSIL,
		SubdivisionUSIN,
		SubdivisionUSKS,
		SubdivisionUSKY,
		SubdivisionUSLA,
		SubdivisionUSMA,
		SubdivisionUSMD,
		SubdivisionUSME,
		SubdivisionUSMI,
		SubdivisionUSMN,
		SubdivisionUSMO,
		SubdivisionUSMP,
		SubdivisionUSMS,
		SubdivisionUSMT,
		SubdivisionUSNC,
		SubdivisionUSND,
		SubdivisionUSNE,
		SubdivisionUSNH,
		SubdivisionUSNJ,
		SubdivisionUSNM,
		SubdivisionUSNV,
		SubdivisionUSNY,
		SubdivisionUSOH,
		SubdivisionUSOK,
		SubdivisionUSOR,
		SubdivisionUSPA,
		SubdivisionUSPR,
		SubdivisionUSRI,
		SubdivisionUSSC,
		SubdivisionUSSD,
		SubdivisionUSTN,
		SubdivisionUSTX,
		SubdivisionUSUM,
		SubdivisionUSUT,
		SubdivisionUSVA,
		SubdivisionUSVI,
		SubdivisionUSVT,
		SubdivisionUSWA,
		SubdivisionUSWI,
		SubdivisionUSWV,
		SubdivisionUSWY,
		SubdivisionUYAR,
		SubdivisionUYCA,
		SubdivisionUYCL,
		SubdivisionUYCO,
		SubdivisionUYDU,
		SubdivisionUYFD,
		SubdivisionUYFS,
		SubdivisionUYLA,
		SubdivisionUYMA,
		SubdivisionUYMO,
		SubdivisionUYPA,
		SubdivisionUYRN,
		SubdivisionUYRO,
		SubdivisionUYRV,
		SubdivisionUYSA,
		SubdivisionUYSJ,
		SubdivisionUYSO,
		SubdivisionUYTA,
		SubdivisionUYTT,
		SubdivisionUZAN,
		SubdivisionUZBU,
		SubdivisionUZFA,
		SubdivisionUZJI,
		SubdivisionUZNG,
		SubdivisionUZNW,
		SubdivisionUZQA,
		SubdivisionUZQR,
		SubdivisionUZSA,
		SubdivisionUZSI,
		SubdivisionUZSU,
		SubdivisionUZTK,
		SubdivisionUZTO,
		SubdivisionUZXO,
		SubdivisionVC01,
		SubdivisionVC02,
		SubdivisionVC03,
		SubdivisionVC04,
		SubdivisionVC05,
		SubdivisionVC06,
		SubdivisionVEA,
		SubdivisionVEB,
		SubdivisionVEC,
		SubdivisionVED,
		SubdivisionVEE,
		SubdivisionVEF,
		SubdivisionVEG,
		SubdivisionVEH,
		SubdivisionVEI,
		SubdivisionVEJ,
		SubdivisionVEK,
		SubdivisionVEL,
		SubdivisionVEM,
		SubdivisionVEN,
		SubdivisionVEO,
		SubdivisionVEP,
		SubdivisionVER,
		SubdivisionVES,
		SubdivisionVET,
		SubdivisionVEU,
		SubdivisionVEV,
		SubdivisionVEW,
		SubdivisionVEX,
		SubdivisionVEY,
		SubdivisionVEZ,
		SubdivisionVN01,
		SubdivisionVN02,
		SubdivisionVN03,
		SubdivisionVN04,
		SubdivisionVN05,
		SubdivisionVN06,
		SubdivisionVN07,
		SubdivisionVN09,
		SubdivisionVN13,
		SubdivisionVN14,
		SubdivisionVN15,
		SubdivisionVN18,
		SubdivisionVN20,
		SubdivisionVN21,
		SubdivisionVN22,
		SubdivisionVN23,
		SubdivisionVN24,
		SubdivisionVN25,
		SubdivisionVN26,
		SubdivisionVN27,
		SubdivisionVN28,
		SubdivisionVN29,
		SubdivisionVN30,
		SubdivisionVN31,
		SubdivisionVN32,
		SubdivisionVN33,
		SubdivisionVN34,
		SubdivisionVN35,
		SubdivisionVN36,
		SubdivisionVN37,
		SubdivisionVN39,
		SubdivisionVN40,
		SubdivisionVN41,
		SubdivisionVN43,
		SubdivisionVN44,
		SubdivisionVN45,
		SubdivisionVN46,
		SubdivisionVN47,
		SubdivisionVN49,
		SubdivisionVN50,
		SubdivisionVN51,
		SubdivisionVN52,
		SubdivisionVN53,
		SubdivisionVN54,
		SubdivisionVN55,
		SubdivisionVN56,
		SubdivisionVN57,
		SubdivisionVN58,
		SubdivisionVN59,
		SubdivisionVN61,
		SubdivisionVN63,
		SubdivisionVN66,
		SubdivisionVN67,
		SubdivisionVN68,
		SubdivisionVN69,
		SubdivisionVN70,
		SubdivisionVN71,
		SubdivisionVN72,
		SubdivisionVN73,
		SubdivisionVNCT,
		SubdivisionVNDN,
		SubdivisionVNHN,
		SubdivisionVNHP,
		SubdivisionVNSG,
		SubdivisionVUMAP,
		SubdivisionVUPAM,
		SubdivisionVUSAM,
		SubdivisionVUSEE,
		SubdivisionVUTAE,
		SubdivisionVUTOB,
		SubdivisionWSAA,
		SubdivisionWSAL,
		SubdivisionWSAT,
		SubdivisionWSFA,
		SubdivisionWSGE,
		SubdivisionWSGI,
		SubdivisionWSPA,
		SubdivisionWSSA,
		SubdivisionWSTU,
		SubdivisionWSVF,
		SubdivisionWSVS,
		SubdivisionYEAB,
		SubdivisionYEAD,
		SubdivisionYEAM,
		SubdivisionYEBA,
		SubdivisionYEDA,
		SubdivisionYEDH,
		SubdivisionYEHD,
		SubdivisionYEHJ,
		SubdivisionYEIB,
		SubdivisionYEJA,
		SubdivisionYELA,
		SubdivisionYEMA,
		SubdivisionYEMR,
		SubdivisionYEMU,
		SubdivisionYEMW,
		SubdivisionYERA,
		SubdivisionYESD,
		SubdivisionYESH,
		SubdivisionYESN,
		SubdivisionYETA,
		SubdivisionZAEC,
		SubdivisionZAFS,
		SubdivisionZAGT,
		SubdivisionZALP,
		SubdivisionZAMP,
		SubdivisionZANC,
		SubdivisionZANL,
		SubdivisionZANW,
		SubdivisionZAWC,
		SubdivisionZM01,
		SubdivisionZM02,
		SubdivisionZM03,
		SubdivisionZM04,
		SubdivisionZM05,
		SubdivisionZM06,
		SubdivisionZM07,
		SubdivisionZM08,
		SubdivisionZM09,
		SubdivisionZWBU,
		SubdivisionZWHA,
		SubdivisionZWMA,
		SubdivisionZWMC,
		SubdivisionZWME,
		SubdivisionZWMI,
		SubdivisionZWMN,
		SubdivisionZWMS,
		SubdivisionZWMV,
		SubdivisionZWMW,
	}
}

// AllSubdivisionsInfo - return all subdivision codes as []Subdivision
func AllSubdivisionsInfo() []*Subdivision {
	all := AllSubdivisions()
	subdivisions := make([]*Subdivision, 0, len(all))
	for _, v := range all {
		subdivisions = append(subdivisions, v.Info())
	}
	return subdivisions
}

// AllSubdivisionsByCountryCode - returns all the subdivisions, mapped to their country code
func AllSubdivisionsByCountryCode() map[CountryCode][]SubdivisionCode {
	resp := map[CountryCode][]SubdivisionCode{}

	for _, s := range AllSubdivisions() {
		c := s.Country()
		if _, ok := resp[c]; !ok {
			resp[c] = []SubdivisionCode{}
		}
		resp[c] = append(resp[c], s)
	}

	return resp
}

// SubdivisionsByCountryCode - returns all subdivisions for a particular country code
func SubdivisionsByCountryCode(c CountryCode) []SubdivisionCode {
	return AllSubdivisionsByCountryCode()[c]
}

// TotalSubdivisions - returns number of subdivisions in the package
func TotalSubdivisions() int {
	return 4885
}
