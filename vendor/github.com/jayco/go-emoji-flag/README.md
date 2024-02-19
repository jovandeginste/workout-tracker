# go-emoji-flag

[![CircleCI](https://circleci.com/gh/jayco/go-emoji-flag.svg?style=svg)](https://circleci.com/gh/jayco/go-emoji-flag)
[![GolangCI](https://golangci.com/badges/github.com/jayco/go-emoji-flag.svg)](https://golangci.com)
[![Go Report Card](https://goreportcard.com/badge/github.com/jayco/go-emoji-flag)](https://golangci.com)

Converts a string country code to an emoji in Go.

## Install

```
go get -u github.com/jayco/go-emoji-flag
```

## Usage

Will return a flag from the supported list below or an empty string if the flag is not found.

```go
package main

import (
	"fmt"

	emoji "github.com/jayco/go-emoji-flag"
)

func main() {
	fmt.Println(emoji.GetFlag("AUS"))   // prints 🇦🇺
	fmt.Println(emoji.GetFlag("AU"))    // prints 🇦🇺
	fmt.Println(emoji.GetFlag("BOB"))   // prints
}
```

## Supported Flags:

```
 BRB/BB - 🇧🇧
 FLK/FK - 🇫🇰
 LVA/LV - 🇱🇻
 OMN/OM - 🇴🇲
 HRV/HR - 🇭🇷
 MNG/MN - 🇲🇳
 RWA/RW - 🇷🇼
 SYR/SY - 🇸🇾
 KHM/KH - 🇰🇭
 NAM/NA - 🇳🇦
 KNA/KN - 🇰🇳
 TKL/TK - 🇹🇰
 PAK/PK - 🇵🇰
 PRI/PR - 🇵🇷
 WLF/WF - 🇼🇫
 ARG/AR - 🇦🇷
 BHS/BS - 🇧🇸
 GAB/GA - 🇬🇦
 JAM/JM - 🇯🇲
 RUS/RU - 🇷🇺
 SDN/SD - 🇸🇩
 LBY/LY - 🇱🇾
 NOR/NO - 🇳🇴
 SUR/SR - 🇸🇷
 HMD/HM - 🇭🇲
 ATG/AG - 🇦🇬
 BHR/BH - 🇧🇭
 CPV/CV - 🇨🇻
 CCK/CC - 🇨🇨
 ATA/AQ - 🇦🇶
 PYF/PF - 🇵🇫
 NRU/NR - 🇳🇷
 GGY/GG - 🇬🇬
 ZAF/ZA - 🇿🇦
 ESP/ES - 🇪🇸
 BOL/BO - 🇧🇴
 DNK/DK - 🇩🇰
 REU/RE - 🇷🇪
 SLB/SB - 🇸🇧
 BGR/BG - 🇧🇬
 KGZ/KG - 🇰🇬
 MAR/MA - 🇲🇦
 TWN/TW - 🇹🇼
 ISL/IS - 🇮🇸
 MWI/MW - 🇲🇼
 THA/TH - 🇹🇭
 AZE/AZ - 🇦🇿
 HKG/HK - 🇭🇰
 COM/KM - 🇰🇲
 FRO/FO - 🇫🇴
 VGB/VG - 🇻🇬
 COD/CD - 🇨🇩
 SHN/SH - 🇸🇭
 YEM/YE - 🇾🇪
 BMU/BM - 🇧🇲
 EST/EE - 🇪🇪
 LSO/LS - 🇱🇸
 MYS/MY - 🇲🇾
 IMN/IM - 🇮🇲
 LTU/LT - 🇱🇹
 SGP/SG - 🇸🇬
 BRA/BR - 🇧🇷
 FSM/FM - 🇫🇲
 PSE/PS - 🇵🇸
 TZA/TZ - 🇹🇿
 GNQ/GQ - 🇬🇶
 KEN/KE - 🇰🇪
 ASM/AS - 🇦🇸
 AND/AD - 🇦🇩
 MTQ/MQ - 🇲🇶
 NLD/NL - 🇳🇱
 SLV/SV - 🇸🇻
 MDA/MD - 🇲🇩
 PRY/PY - 🇵🇾
 BLM/BL - 🇧🇱
 ROU/RO - 🇷🇴
 CMR/CM - 🇨🇲
 IND/IN - 🇮🇳
 MDG/MG - 🇲🇬
 MOZ/MZ - 🇲🇿
 MSR/MS - 🇲🇸
 TKM/TM - 🇹🇲
 UZB/UZ - 🇺🇿
 GUY/GY - 🇬🇾
 KAZ/KZ - 🇰🇿
 MKD/MK - 🇲🇰
 MLT/MT - 🇲🇹
 SSD/SS - 🇸🇸
 ALB/AL - 🇦🇱
 BFA/BF - 🇧🇫
 COK/CK - 🇨🇰
 GRC/GR - 🇬🇷
 GHA/GH - 🇬🇭
 MAF/MF - 🇲🇫
 COL/CO - 🇨🇴
 PRK/KP - 🇰🇵
 LIE/LI - 🇱🇮
 MCO/MC - 🇲🇨
 BTN/BT - 🇧🇹
 VAT/VA - 🇻🇦
 MYT/YT - 🇾🇹
 SVK/SK - 🇸🇰
 TTO/TT - 🇹🇹
 LAO/LA - 🇱🇦
 MLI/ML - 🇲🇱
 SAU/SA - 🇸🇦
 CHE/CH - 🇨🇭
 EGY/EG - 🇪🇬
 ETH/ET - 🇪🇹
 PER/PE - 🇵🇪
 PCN/PN - 🇵🇳
 AUT/AT - 🇦🇹
 BIH/BA - 🇧🇦
 BWA/BW - 🇧🇼
 CAN/CA - 🇨🇦
 VCT/VC - 🇻🇨
 ARE/AE - 🇦🇪
 VIR/VI - 🇻🇮
 IOT/IO - 🇮🇴
 CAF/CF - 🇨🇫
 ISR/IL - 🇮🇱
 URY/UY - 🇺🇾
 PRT/PT - 🇵🇹
 BLR/BY - 🇧🇾
 HTI/HT - 🇭🇹
 UGA/UG - 🇺🇬
 AFG/AF - 🇦🇫
 CYP/CY - 🇨🇾
 GTM/GT - 🇬🇹
 IRQ/IQ - 🇮🇶
 TUR/TR - 🇹🇷
 AGO/AO - 🇦🇴
 DJI/DJ - 🇩🇯
 IDN/ID - 🇮🇩
 SGS/GS - 🇬🇸
 SRB/RS - 🇷🇸
 SJM/SJ - 🇸🇯
 ZWE/ZW - 🇿🇼
 CUB/CU - 🇨🇺
 GNB/GW - 🇬🇼
 SWZ/SZ - 🇸🇿
 SPM/PM - 🇵🇲
 GBR/GB - 🇬🇧
 ABW/AW - 🇦🇼
 MAC/MO - 🇲🇴
 NIC/NI - 🇳🇮
 MNP/MP - 🇲🇵
 LBR/LR - 🇱🇷
 SMR/SM - 🇸🇲
 DZA/DZ - 🇩🇿
 BRN/BN - 🇧🇳
 DEU/DE - 🇩🇪
 TLS/TL - 🇹🇱
 NPL/NP - 🇳🇵
 POL/PL - 🇵🇱
 TON/TO - 🇹🇴
 BVT/BV - 🇧🇻
 CHL/CL - 🇨🇱
 ECU/EC - 🇪🇨
 MEX/MX - 🇲🇽
 CZE/CZ - 🇨🇿
 GEO/GE - 🇬🇪
 ARM/AM - 🇦🇲
 BLZ/BZ - 🇧🇿
 MHL/MH - 🇲🇭
 AUS/AU - 🇦🇺
 GRL/GL - 🇬🇱
 NGA/NG - 🇳🇬
 GUM/GU - 🇬🇺
 JPN/JP - 🇯🇵
 VEN/VE - 🇻🇪
 PAN/PA - 🇵🇦
 TJK/TJ - 🇹🇯
 BEL/BE - 🇧🇪
 DOM/DO - 🇩🇴
 ERI/ER - 🇪🇷
 MNE/ME - 🇲🇪
 LCA/LC - 🇱🇨
 SEN/SN - 🇸🇳
 FIN/FI - 🇫🇮
 GRD/GD - 🇬🇩
 IRL/IE - 🇮🇪
 JOR/JO - 🇯🇴
 PLW/PW - 🇵🇼
 SVN/SI - 🇸🇮
 BGD/BD - 🇧🇩
 TCD/TD - 🇹🇩
 JEY/JE - 🇯🇪
 MDV/MV - 🇲🇻
 MRT/MR - 🇲🇷
 NER/NE - 🇳🇪
 USA/US - 🇺🇸
 AIA/AI - 🇦🇮
 BEN/BJ - 🇧🇯
 NIU/NU - 🇳🇺
 CHN/CN - 🇨🇳
 CXR/CX - 🇨🇽
 ESH/EH - 🇪🇭
 QAT/QA - 🇶🇦
 TGO/TG - 🇹🇬
 TCA/TC - 🇹🇨
 GLP/GP - 🇬🇵
 KIR/KI - 🇰🇮
 KOR/KR - 🇰🇷
 MUS/MU - 🇲🇺
 GUF/GF - 🇬🇫
 ALA/AX - 🇦🇽
 PHL/PH - 🇵🇭
 SWE/SE - 🇸🇪
 VUT/VU - 🇻🇺
 HUN/HU - 🇭🇺
 SLE/SL - 🇸🇱
 CIV/CI - 🇨🇮
 FJI/FJ - 🇫🇯
 LKA/LK - 🇱🇰
 UKR/UA - 🇺🇦
 ZMB/ZM - 🇿🇲
 COG/CG - 🇨🇬
 DMA/DM - 🇩🇲
 HND/HN - 🇭🇳
 UMI/UM - 🇺🇲
 PNG/PG - 🇵🇬
 CRI/CR - 🇨🇷
 GMB/GM - 🇬🇲
 IRN/IR - 🇮🇷
 NZL/NZ - 🇳🇿
 BDI/BI - 🇧🇮
 FRA/FR - 🇫🇷
 ITA/IT - 🇮🇹
 LBN/LB - 🇱🇧
 KWT/KW - 🇰🇼
 WSM/WS - 🇼🇸
 STP/ST - 🇸🇹
 TUN/TN - 🇹🇳
 TUV/TV - 🇹🇻
 CYM/KY - 🇰🇾
 SYC/SC - 🇸🇨
 VNM/VN - 🇻🇳
 ATF/TF - 🇹🇫
 GIB/GI - 🇬🇮
 LUX/LU - 🇱🇺
 MMR/MM - 🇲🇲
 GIN/GN - 🇬🇳
 NCL/NC - 🇳🇨
 NFK/NF - 🇳🇫
 SOM/SO - 🇸🇴
```
