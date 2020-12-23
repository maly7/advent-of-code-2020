package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	passports := strings.Split(Puzzle, "\n\n")
	valid := 0

	for _, passport := range passports {
		if checkValid(passport) {
			valid++
		}
	}

	fmt.Printf("Total valid passports: %d\n", valid)
}

func checkValid(passport string) bool {
	fields := strings.Fields(passport)

	for _, rule := range Rules {
		if !strings.Contains(passport, rule.field) {
			return false
		}

		field := getField(fields, rule.field)

		if !rule.check(field) {
			fmt.Printf("Invalid rule: %s with value %s\n", rule.field, field)
			return false
		}
	}

	return true
}

func getField(passportFields []string, field string) string {
	for _, s := range passportFields {
		if strings.Contains(s, field) {
			return strings.Split(s, field)[1]
		}
	}

	return ""
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

type rule struct {
	field string
	check func(string) bool
}

var validHcl = regexp.MustCompile("#[a-zA-Z0-9]{6}")
var validEcl = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
var validPid = regexp.MustCompile("[0-9]{9}")

var Rules = []rule{
	{"byr:", func(s string) bool {
		y, e := strconv.Atoi(s)
		return e == nil && y >= 1920 && y <= 2002
	}}, {"iyr:", func(s string) bool {
		y, e := strconv.Atoi(s)
		return e == nil && y >= 2010 && y <= 2020
	}}, {"eyr:", func(s string) bool {
		y, e := strconv.Atoi(s)
		return e == nil && y >= 2020 && y <= 2030
	},
	}, {"hgt:", func(s string) bool {
		if strings.HasSuffix(s, "in") {
			hgt, _ := strconv.Atoi(strings.Split(s, "in")[0])
			return hgt >= 59 && hgt <= 76
		} else if strings.HasSuffix(s, "cm") {
			hgt, _ := strconv.Atoi(strings.Split(s, "cm")[0])
			return hgt >= 150 && hgt <= 193
		}

		return false
	},
	}, {"hcl:", func(s string) bool {
		return len(s) == 7 && validHcl.MatchString(s)
	},
	}, {"ecl:", func(s string) bool {
		return contains(validEcl, s)
	},
	}, {"pid:", func(s string) bool {
		return len(s) == 9 && validPid.MatchString(s)
	}},
}

const Puzzle = "iyr:2015 cid:189 ecl:oth byr:1947 hcl:#6c4ab1 eyr:2026\nhgt:174cm\npid:526744288\n\npid:688706448 iyr:2017 hgt:162cm cid:174 ecl:grn byr:1943 hcl:#808e9e eyr:2025\n\necl:oth hcl:#733820 cid:124 pid:111220591\niyr:2019 eyr:2001\nbyr:1933 hgt:159in\n\npid:812929897 hgt:159cm hcl:#fffffd byr:1942 iyr:2026 cid:291\necl:oth\neyr:2024\n\ncid:83 pid:524032739 iyr:2013 ecl:amb byr:1974\nhgt:191cm hcl:#ceb3a1 eyr:2028\n\necl:gry hcl:eefed5 pid:88405792 hgt:183cm cid:221 byr:1963 eyr:2029\n\npid:777881168 ecl:grn\nhgt:181cm byr:1923 eyr:2021 iyr:2018 hcl:#18171d\n\nbyr:1941 eyr:2027 ecl:gry iyr:2016 pid:062495008 hcl:#a5e1b5 hgt:178cm\n\ncid:56\nbyr:1971\nhcl:#efcc98 pid:649868696 iyr:2011 eyr:2025 hgt:164cm\n\necl:blu\npid:117915262 eyr:2023 byr:1925 iyr:2020 hcl:#888785\nhgt:188cm\n\niyr:2012\ncid:174\neyr:2024\npid:143293382 ecl:brn byr:1946 hgt:193cm\n\neyr:2021 iyr:2011\nhgt:192cm pid:251564680\nbyr:1976\necl:blu hcl:#602927\n\nbyr:1973 ecl:blu hgt:164cm\neyr:2022 pid:695538656 iyr:2010 cid:244 hcl:#b6652a\n\niyr:2014\neyr:2027 pid:358398181 ecl:hzl hgt:74in byr:1949 cid:329\nhcl:#ceb3a1\n\ncid:211\nbyr:1954 eyr:2023 hgt:172cm ecl:blu iyr:2019 hcl:#623a2f pid:657051725\n\npid:562699115 eyr:2026 byr:2000\nhgt:162cm hcl:#602927 ecl:amb iyr:2018\n\necl:brn\niyr:2013\npid:835184859 byr:1981 hgt:157cm eyr:2027 hcl:#b6652a\n\npid:763432667 byr:1981 hcl:#cfa07d ecl:brn\niyr:2010 hgt:63in cid:107\neyr:2027\n\nbyr:2009\nhgt:177cm cid:314\nhcl:f55bf8 eyr:2025\npid:632519974\niyr:2015 ecl:amb\n\neyr:2024 pid:614239656 hgt:169cm iyr:2014 ecl:hzl byr:1992\nhcl:#602927\n\necl:blu\neyr:2026\nhcl:#efcc98\nbyr:1980 iyr:2013\nhgt:161cm\npid:065413599\n\nhgt:182cm\neyr:2025 iyr:2013 pid:939088351 hcl:#b6652a byr:1994 ecl:amb\n\nhgt:65in cid:220 ecl:amb hcl:#ceb3a1\niyr:2013 eyr:2025 pid:167894964 byr:1976\n\nhgt:185cm cid:88 ecl:blu iyr:2020\neyr:2020\nhcl:#888785 pid:582683387\nbyr:1981\n\nhcl:#866857 eyr:2020 byr:1948\npid:358943355\necl:amb hgt:164cm iyr:2019\n\npid:127467714\nhcl:#ceb3a1 byr:1991 hgt:163cm eyr:2020 iyr:2017 ecl:blu cid:229\n\ncid:156 byr:1942 eyr:2024 hcl:#cfa07d\necl:blu pid:843747591\niyr:2014 hgt:173cm\n\nhcl:#a97842 hgt:165cm\niyr:2013 ecl:#781088 byr:1952\npid:516882944\neyr:2026\n\nhgt:179cm\nbyr:1969 pid:408297435 iyr:2020 ecl:oth hcl:#cfa07d eyr:2020\n\necl:amb iyr:2013 hcl:#b6652a eyr:2023 cid:88\npid:324081998 hgt:66in byr:1945\n\niyr:2012\neyr:2024\nhcl:#18171d\npid:756726480 byr:1947 ecl:oth\nhgt:164cm\n\necl:blu\nhcl:#fffffd byr:1951 iyr:2019 pid:544645775\nhgt:153cm eyr:2027\n\npid:655906238 ecl:brn eyr:2028 byr:1959 hgt:63in cid:338\niyr:2020\n\neyr:2020\nhcl:#602927 hgt:72in iyr:2014\npid:305025767\ncid:297 byr:1957 ecl:gry\n\nhgt:155cm byr:1942 hcl:#a97842\niyr:2014 ecl:gry pid:593995708\neyr:2022\n\npid:219206471 byr:1955 eyr:2030\nhcl:#a97842 ecl:oth iyr:2015 cid:134 hgt:170cm\n\niyr:2013 cid:268\neyr:2020\nhcl:#a97842 ecl:grn pid:235279200 hgt:178cm\nbyr:1952\n\niyr:2013 pid:016384352 eyr:2027\nhcl:#866857 ecl:grn hgt:161cm byr:1943\n\necl:amb hgt:169cm pid:149540593\niyr:2012\neyr:2040 hcl:#a97842 byr:1954\n\nbyr:1938\necl:brn hcl:#b6652a eyr:2026 hgt:184cm iyr:2018 pid:832531235\n\nbyr:1945 iyr:2015 hgt:171cm eyr:2028 pid:998746896 ecl:hzl hcl:#866857\n\nhgt:73in ecl:hzl eyr:2023 cid:343 pid:458004221 iyr:2017 byr:1962 hcl:#efcc98\n\nbyr:1970 hgt:159cm pid:925022199 iyr:2013\neyr:2028 hcl:#888785\necl:hzl\n\neyr:2027 iyr:2016 ecl:gry\nhcl:#cfa07d\npid:006246552 byr:1939 cid:124 hgt:177cm\n\nbyr:1982\niyr:2016 hgt:159cm\ncid:102 hcl:#fffffd\neyr:2029\necl:grn pid:619798285\n\niyr:2018\nhgt:189cm hcl:#efcc98\nbyr:1937 eyr:2023 pid:727551553 ecl:oth\n\niyr:2014 byr:1976\neyr:2020 hcl:#7d3b0c pid:125102070 ecl:amb\nhgt:186cm\n\nhgt:187cm byr:1949\npid:027653233 eyr:2021 hcl:#341e13 ecl:hzl\niyr:2020\n\niyr:2016\nbyr:1954 pid:545631256\nhcl:#602927 eyr:2023\nhgt:191cm ecl:amb\n\npid:509762954\nhgt:190cm ecl:hzl byr:1991\neyr:2022 iyr:2019\ncid:187\n\nhcl:#c0946f eyr:2024 hgt:152cm cid:277 iyr:2015 pid:872373191 byr:1988\n\npid:544267207 cid:113\niyr:2015\nhgt:181cm\nhcl:#6b5442\necl:gry\nbyr:1971\n\necl:gry\nhgt:161cm iyr:2012 byr:1965\npid:574527322 hcl:#fffffd\n\niyr:2018 byr:1976 hcl:#b6652a\npid:024582079 hgt:169cm ecl:oth eyr:2021\n\npid:020478204\nbyr:1945 hcl:#7d3b0c\ncid:239 eyr:2025 hgt:188cm\necl:grn\niyr:2012\n\neyr:2026 pid:202653345\nbyr:1988\nhcl:#2cdc09\nhgt:185cm iyr:2010\necl:hzl\n\nhgt:183cm iyr:2017\nhcl:#18171d byr:1977 eyr:2029 pid:804559436 ecl:grn\n\nhcl:#602927 pid:812072269 hgt:170cm eyr:2026 byr:1955 iyr:2020 ecl:gry\n\neyr:2023 iyr:2010\nhcl:#cfa07d pid:592419048 byr:1943\necl:brn\nhgt:172cm\n\necl:brn iyr:2013 pid:558179058\nhcl:#fffffd eyr:2022\nbyr:1922\ncid:331 hgt:64in\n\necl:xry\nhcl:ade850 eyr:1995 pid:976028541\niyr:2030 hgt:179cm\nbyr:2030\n\necl:#2872b1 pid:158cm eyr:1927 hcl:ee8e92\niyr:2014 hgt:190cm\nbyr:2025\n\nhgt:155cm cid:283 eyr:2020 ecl:blu pid:755165290 byr:1936 hcl:#733820 iyr:2012\n\neyr:2030\nbyr:1943\ncid:323 pid:906418061 hgt:157cm ecl:amb iyr:2010\nhcl:#7d3b0c\n\nhcl:#fffffd\npid:873200829 hgt:192cm eyr:2022 ecl:blu iyr:2016 byr:1920 cid:200\n\neyr:2021\nbyr:1963\nhcl:#a97842 pid:585551405\niyr:2019 cid:91\necl:brn hgt:60cm\n\nbyr:1946\npid:520273609 hcl:#341e13 cid:66\niyr:2020 hgt:154cm eyr:2024\necl:brn\n\necl:brn hcl:#d64d7b eyr:2020\nbyr:1957 hgt:181cm iyr:2019 pid:378496967 cid:135\n\npid:002446580\neyr:2027 byr:1939 hcl:#888785\niyr:2011 cid:168\necl:oth hgt:160cm\n\niyr:2019 hgt:70in hcl:#7d3b0c byr:1983\neyr:2024 pid:369493064 cid:54 ecl:oth\n\niyr:1979 pid:170cm\nhgt:65cm eyr:1933 hcl:z\n\necl:zzz pid:193cm hcl:z eyr:2020 byr:2013 iyr:2016 hgt:177in\n\niyr:2010 hgt:187cm\nbyr:1932\nhcl:z ecl:oth pid:665967850 eyr:2030\n\neyr:2029\niyr:2013 hcl:#b6652a ecl:amb\nbyr:1936 pid:516025566\nhgt:181cm\n\nhcl:#c0946f pid:238825672 byr:2000\niyr:2013 eyr:2028 ecl:amb hgt:183cm\n\neyr:2021 hcl:#866857\ncid:77 iyr:2017 hgt:156cm pid:271118829 ecl:amb\n\niyr:2014\nhcl:#fffffd\ncid:321 hgt:159cm ecl:gry\npid:691381062 eyr:2022 byr:1991\n\npid:111506492 hcl:#c1d296 iyr:2011\nbyr:1934 hgt:176cm cid:263 eyr:2028 ecl:amb\n\niyr:2014 hgt:64in eyr:2024 cid:193 hcl:#b6652a byr:1967\necl:oth pid:138677174\n\nhgt:168cm iyr:2020 eyr:2030\nhcl:#6b5442 ecl:brn pid:975843892 byr:1927\n\nbyr:1957 ecl:amb iyr:2012 pid:177266671 eyr:2026\nhcl:#866857 hgt:162cm\n\neyr:2029\nhcl:#341e13\nhgt:175cm pid:465809700 ecl:amb byr:1974\niyr:2010\n\nhcl:#a97842 iyr:2010\nhgt:176cm eyr:2029 byr:1931 ecl:grt pid:161604244\n\neyr:2024 iyr:2018 hgt:170in byr:1959 ecl:gmt hcl:#888785\npid:94163132\n\niyr:2011\nhgt:186cm pid:998471478 byr:1956 ecl:amb\neyr:2029\nhcl:#efcc98\ncid:76\n\necl:brn\nbyr:2001 pid:378527883 iyr:2013 hcl:#83bdc5 eyr:2020 hgt:181cm\n\niyr:2017 ecl:grn hgt:172cm hcl:#888785 cid:100\neyr:2022 byr:2030\npid:311562177\n\npid:097558436\ncid:141 hgt:152cm iyr:2019\necl:brn eyr:2023\nbyr:1940\nhcl:#6b5442\n\niyr:2016 eyr:2023 byr:1992\nhgt:174cm ecl:amb\npid:691291640 cid:190 hcl:#fffffd\n\nhcl:#623a2f ecl:brn\neyr:2028 cid:227 iyr:2012 hgt:74in pid:964273950 byr:1965\n\nhcl:#ceb3a1 eyr:2028\niyr:2013 pid:175294029 hgt:150cm ecl:grn\nbyr:1936\ncid:143\n\nbyr:1935 hcl:#a97842 ecl:oth hgt:180cm iyr:2019\npid:857891916\neyr:2026\n\npid:084518249 ecl:hzl eyr:2027 hcl:#c0946f hgt:192cm cid:315 byr:1961\niyr:2010\n\nhgt:67cm pid:37925169 eyr:2022\nhcl:z iyr:2012 cid:315 byr:2028 ecl:dne\n\nhcl:#c0946f byr:1924\nhgt:176cm cid:87 pid:682212551 iyr:2011\neyr:2026\necl:gry\n\nhgt:181cm byr:1935\niyr:2018 pid:644964785\neyr:2026 ecl:amb\n\npid:789810179\necl:gry eyr:2021\ncid:159 hgt:185cm iyr:2020 hcl:#602927\nbyr:1965\n\npid:672386364\niyr:2013 eyr:2021 byr:1951 hcl:#341e13\necl:gry hgt:173cm\n\nhcl:#18171d eyr:2030 pid:957722245 iyr:2012 byr:1955\necl:grn\nhgt:154cm\n\nbyr:1955 ecl:oth\nhcl:#cfa07d\neyr:2030\niyr:2013 pid:361945273 hgt:154cm\n\niyr:2012 eyr:2027 ecl:grn hcl:#16d373\nhgt:192cm\n\npid:275525273\nbyr:1986\niyr:2017\neyr:2022\necl:grn\nhgt:75in\nhcl:#919cc0\n\neyr:2029\ncid:84 hcl:#cfa07d iyr:2013 hgt:78\necl:brn\nbyr:1925 pid:281331549\n\neyr:2027\ncid:219 iyr:2016 byr:1971 hcl:#7d3b0c hgt:179cm ecl:grn\npid:301296222\n\neyr:2030 iyr:2010 pid:995982765\nbyr:1926 ecl:amb hcl:#888785 hgt:186cm\n\nbyr:1955 iyr:2015 hgt:165cm cid:101\neyr:2027 ecl:amb hcl:#602927\npid:168654790\n\nhcl:#7d3b0c byr:1956 eyr:2029 hgt:155cm\necl:grn pid:816685992\niyr:2016\n\necl:grn hcl:#cfa07d cid:71\npid:914724136 iyr:2012 eyr:2024\nhgt:184cm byr:1938\n\necl:gry\neyr:2029 hcl:#602927 pid:255062643 iyr:2015 hgt:175cm\n\nhcl:#341e13 iyr:2017 eyr:2028\npid:459704815 byr:1922\ncid:312\necl:brn hgt:152cm\n\necl:dne eyr:1981\npid:8356519470 hgt:176 iyr:1941 byr:2006 hcl:z\n\necl:amb pid:753377589 hcl:#a97842 eyr:2022 hgt:187cm\ncid:130 iyr:2013 byr:1961\n\npid:952444443\nhcl:#bde835 byr:1963 iyr:2020 eyr:2025\necl:amb hgt:162cm\n\neyr:2027 iyr:2018 hcl:#ceb3a1 hgt:152cm pid:882429463 ecl:blu byr:1969\n\ncid:134 eyr:2021 hcl:#a97842 hgt:63in\necl:grn byr:1975 iyr:2019 pid:154078695\n\nbyr:1956 eyr:2027\npid:396230480 hcl:#b6652a\nhgt:175cm iyr:2020 ecl:oth\n\necl:grn\ncid:263 hcl:#506937 byr:1924\neyr:2030 pid:705511368 hgt:159cm\niyr:2011\n\neyr:2020 hgt:178cm ecl:grn\nbyr:1947 hcl:#888785\npid:177476829 iyr:2019\n\necl:hzl cid:211 iyr:2016 hgt:176cm pid:405182470\nbyr:1952\nhcl:#866857 eyr:2028\n\neyr:2032 cid:152 ecl:gmt hgt:150in\npid:75969209\nbyr:2019 hcl:z iyr:1940\n\nhcl:#fffffd hgt:193cm pid:607407479 cid:300 byr:1944 iyr:2017\necl:oth\neyr:2026\n\nhcl:z\ncid:125 eyr:2040 ecl:dne byr:2015 pid:733096171 hgt:63cm\niyr:1922\n\npid:575721428 hgt:152cm cid:275\nhcl:#cfa07d eyr:2028\nbyr:1935 ecl:hzl iyr:2016\n\niyr:2012\necl:grn eyr:2027 hcl:#623a2f pid:029106453 byr:1984 hgt:168cm\n\necl:blu cid:140 eyr:2028 iyr:2018 hcl:#c0946f\nhgt:163cm byr:1944\npid:709288293\n\nbyr:1936\nhgt:172cm eyr:1997 hcl:#8b8c88 cid:50\niyr:2016 pid:205477922 ecl:grn\n\nhgt:170cm pid:872750582 eyr:2027 byr:1985 iyr:2017 hcl:#d6976a ecl:blu\n\nhgt:163cm\npid:189634089 cid:116 byr:1975 eyr:2030\nhcl:#efcc98 ecl:brn iyr:2020\n\necl:amb byr:1953 hcl:#6b5442 pid:418787965\niyr:2018 hgt:193cm\neyr:2026\n\necl:#3ec898 cid:339 hcl:#866857 eyr:2025 hgt:179cm pid:591430028 iyr:1936 byr:1995\n\npid:285371937 hgt:159cm\nbyr:1922\niyr:2013 eyr:2023 hcl:#6b5442 ecl:amb\n\npid:545260883 ecl:oth\nhgt:163cm\niyr:2015 eyr:2021 byr:1975 hcl:#866857\n\necl:hzl hgt:182cm pid:053762098 eyr:2023 cid:174 hcl:#6daac4 iyr:2017 byr:1937\n\nhgt:178cm iyr:2015 byr:1956 pid:815359103\necl:blu hcl:#cfa07d eyr:2030\n\nhcl:#7d3b0c\npid:438108851 hgt:162cm byr:1930 iyr:2014 eyr:2024 ecl:amb\n\neyr:2027 iyr:2019 hcl:#90eb1c hgt:178cm\npid:314810594 cid:278 ecl:amb\nbyr:2001\n\nbyr:1949 iyr:1942 hcl:#888785 ecl:hzl hgt:184cm eyr:2027 pid:899137640\n\nhgt:153cm\neyr:2022 iyr:2011 byr:1975\nhcl:#602927\necl:amb pid:178cm\n\nhcl:#6b5442\necl:amb iyr:2018 eyr:2025 pid:418735327 byr:1922 hgt:74in\n\necl:gmt hcl:z iyr:2024\neyr:1988 hgt:75cm cid:125 pid:690872200 byr:1928\n\neyr:2024 hgt:184cm\npid:4634589837 ecl:zzz iyr:2022 byr:2000 hcl:89c187\n\niyr:2017 byr:1966 hcl:#efcc98 ecl:brn pid:473085232 eyr:2021 hgt:174cm\n\nhgt:67in eyr:2030 iyr:2014 byr:1943 hcl:#602927 cid:344\necl:oth\npid:210476779\n\nbyr:1955\necl:oth\nhgt:193cm iyr:2012 hcl:#623a2f pid:818289829 eyr:2021\n\nbyr:2018 ecl:#872a51 iyr:2024 hcl:97783d\npid:155cm hgt:174cm\neyr:1964\n\nhcl:#6b5442 hgt:157cm byr:1932 ecl:brn pid:4275535874\neyr:2024 iyr:2015\n\npid:959861097\nhgt:151cm cid:140 byr:1935\neyr:2029\niyr:2018 ecl:hzl\nhcl:#623a2f\n\nhgt:181cm pid:911791767 eyr:2027\niyr:2016 byr:1962\necl:grn hcl:#866857\n\neyr:2021\nbyr:1994\nhgt:162cm hcl:#866857 ecl:oth iyr:2014\npid:712345689\n\nhcl:#7d3b0c\nhgt:170cm pid:600132416 eyr:2025\niyr:2016 byr:1978 ecl:brn\n\nhcl:#0a9307\ncid:287 byr:1940 pid:786271493\neyr:2028 hgt:186cm\niyr:2019 ecl:oth\n\neyr:2025 hgt:190cm ecl:hzl cid:228 iyr:2019\nbyr:1932\nhcl:#623a2f pid:648307551\n\npid:304587325 iyr:2019 byr:1923 hcl:#7d3b0c\nhgt:190cm\necl:gry eyr:2030\n\nhgt:188cm eyr:2027 byr:1958 pid:572934921\nhcl:#888785 ecl:hzl iyr:2010\n\niyr:2019\nhgt:178cm ecl:grn hcl:#7d3b0c pid:007601227\nbyr:1975 eyr:2023\n\npid:808872803 byr:1929\necl:grn\neyr:2022 iyr:2019 hgt:74in hcl:#602927\n\niyr:2019\ncid:67 hcl:#602927 pid:292601338 ecl:hzl\nbyr:2001 eyr:2023 hgt:171cm\n\nbyr:1962 eyr:2022 hcl:#b6652a hgt:193cm\necl:oth\niyr:2010\n\nhgt:70in iyr:2014 hcl:#a97842\ncid:169 eyr:2020 ecl:amb\npid:329751670 byr:1959\n\nbyr:1920\necl:oth hgt:172cm cid:57 pid:515139276\neyr:2030\nhcl:#18171d\niyr:2013\n\niyr:2012\nhcl:#a97842 pid:946040810 hgt:65in\nbyr:1936 ecl:amb eyr:2020\n\nbyr:1948 hcl:#18171d\niyr:2019\necl:hzl cid:185\neyr:2023\npid:583625200 hgt:191cm\n\nhgt:154cm eyr:2022\npid:460137392 iyr:2010\necl:grn\nhcl:#ceb3a1\n\neyr:2024\niyr:2016 pid:890698391 hgt:172cm hcl:#a97842 cid:271 ecl:oth byr:1926\n\nhgt:162cm pid:340904964 hcl:#b6652a\nbyr:1966\niyr:2010\ncid:260 eyr:2028\necl:amb\n\nbyr:1933 eyr:2029 pid:642043350\niyr:2016 hcl:#b6652a ecl:grn\n\npid:602218620 eyr:2023 ecl:blu\nhcl:#623a2f\nbyr:1950 hgt:168cm iyr:2015\n\necl:gry pid:490792384\nbyr:1974\nhcl:#a97842 iyr:2016 hgt:170cm\n\niyr:2020 ecl:gry byr:2002\neyr:2029 hcl:#9f45c4\nhgt:155cm pid:604239618\n\nhgt:190cm pid:560653271 iyr:2020 cid:349\neyr:2024 ecl:blu hcl:#efcc98 byr:1936\n\neyr:2021 byr:1964 hcl:#efcc98 ecl:grn iyr:2018\nhgt:165cm pid:218376636\n\npid:186217101\niyr:2019 hgt:155cm\nbyr:2017 eyr:2022 ecl:grn cid:349 hcl:ece72e\n\niyr:2015\neyr:2026 pid:802832833\nhcl:#888785 hgt:190cm ecl:brn\nbyr:1952\ncid:202\n\ncid:151 iyr:2017 hgt:152cm hcl:#a97842 eyr:2020 ecl:hzl\npid:554959609 byr:1941\n\ncid:116\niyr:2019 hgt:159cm byr:1992 pid:662111811\nhcl:#18171d ecl:oth eyr:2024\n\necl:grn byr:1966\niyr:1950 pid:585351486\neyr:2038 hgt:178in hcl:a27d2b\n\niyr:2014 cid:238 hgt:187cm pid:523401750 ecl:amb hcl:#18171d eyr:2023 byr:1984\n\neyr:2021 byr:1957\npid:340752324\niyr:2015 hgt:157cm\nhcl:#602927 cid:70\necl:oth\n\npid:458479816 ecl:hzl\neyr:2022 hcl:z\nhgt:60cm\nbyr:2012 iyr:2005\n\ncid:57\nhgt:154cm pid:446142864\nhcl:#341e13 byr:1968 eyr:2030\niyr:2019\necl:brn\n\neyr:2028\npid:243811429 byr:1977\niyr:2011 hcl:#18171d hgt:185cm ecl:oth\n\ncid:205 byr:1976 eyr:2029 pid:649877471 hcl:#cfa07d hgt:152cm\necl:blu\niyr:2013\n\niyr:2009 pid:559014976 ecl:oth hgt:189cm byr:1936 eyr:2037\nhcl:#efcc98\n\npid:134378987 byr:1983 iyr:2013 hgt:173cm\necl:oth hcl:#ceb3a1\ncid:80\neyr:2020\n\nhgt:151cm byr:1964 ecl:grn iyr:2010 hcl:#b6652a pid:939492531\neyr:2028\n\nbyr:1961 iyr:2014 hcl:#733820 hgt:179cm\neyr:2026 ecl:gry pid:732892920\n\niyr:2018 byr:1996\npid:944007809 ecl:hzl\nhcl:#866857 eyr:2021\nhgt:155cm\n\npid:374875696 hcl:#7d3b0c\necl:oth\nhgt:193cm byr:1948 cid:238\niyr:2020\n\npid:305782299 hcl:#b6652a\necl:brn\nhgt:172cm\niyr:2018 byr:1927\n\npid:945869114 cid:95 byr:1989 hgt:173cm eyr:2025 hcl:#b6652a iyr:2012 ecl:amb\n\npid:55484149\neyr:1958\niyr:1956 ecl:grn\ncid:95 byr:2028\nhcl:c2af7e\n\nhgt:176cm ecl:amb\nhcl:#a97842 eyr:2029 pid:937928270\ncid:251\nbyr:1978\niyr:2018\n\nhgt:154cm\ncid:213 pid:767329807 ecl:hzl\niyr:2013\nhcl:#888785\neyr:2026 byr:1998\n\ncid:158 hcl:#b6652a hgt:155cm iyr:2010 eyr:2025\nbyr:1980 pid:338567803 ecl:amb\n\nhcl:#efcc98 byr:1940 hgt:62in ecl:oth pid:537307591\neyr:2030\niyr:2017\ncid:179\n\nbyr:1965 eyr:2027 pid:691913618 hgt:75in\nhcl:#6b5442 ecl:gry iyr:2012\n\nhgt:163cm byr:1964 eyr:2025\niyr:2010 hcl:#ceb3a1 ecl:oth\npid:936536544\n\npid:712946803\ncid:343\nhgt:187cm ecl:oth iyr:2020 byr:1983 eyr:2030\nhcl:#7873b3\n\necl:blu\niyr:2010\nhcl:#fffffd\neyr:2030\nhgt:175cm pid:047567505 byr:1963\n\necl:gry byr:1946 eyr:2026 hcl:#602927\nhgt:164cm\niyr:2010\n\npid:223378458\niyr:2014 cid:151 ecl:hzl hgt:171cm\neyr:2020\nhcl:#341e13 byr:1964\n\necl:brn byr:1948\nhcl:#866857\nhgt:193cm eyr:2024\niyr:2013 cid:277\n\nhcl:#623a2f byr:1943 iyr:2011 ecl:oth\nhgt:184cm\npid:371604584 eyr:2024 cid:176\n\nhcl:#efcc98\neyr:2025 pid:241834382\nhgt:178cm\nbyr:1985\niyr:2017\n\nhcl:#c0946f\nbyr:1996 pid:701366586 eyr:2026 hgt:163cm iyr:2015 ecl:oth\n\nhgt:65cm hcl:#18171d\neyr:2024 ecl:brn pid:172cm\niyr:2010\nbyr:1990\n\nhcl:#fffffd pid:68659204 hgt:161cm iyr:2025\necl:#94b8aa byr:2021 eyr:2032\n\necl:blu iyr:2018 byr:1993 cid:184\nhgt:177cm pid:289871693 hcl:#733820 eyr:2026\n\ncid:138\necl:gry hgt:174cm eyr:2024 byr:1988 iyr:2014 hcl:#341e13 pid:864852584\n\ncid:321 eyr:2028 pid:93285596 hgt:173cm\niyr:2013 ecl:gry hcl:#623a2f\nbyr:1927\n\npid:431242259 eyr:2022 ecl:hzl\nbyr:1960 hgt:151cm hcl:#efcc98 iyr:2020\n\nhcl:#866857 eyr:2029 iyr:2016 ecl:grn pid:526060780 byr:1929\ncid:310 hgt:162cm\n\necl:blu hgt:183cm cid:168\niyr:2015\neyr:2021 byr:1951 hcl:#6b5442\npid:594960553\n\nhcl:#ceb3a1\niyr:2020 byr:1951 hgt:186cm eyr:2022 ecl:amb pid:317661479\n\niyr:2016\nhgt:163in hcl:#accfa0\necl:brn\npid:307377995 byr:2000 eyr:2028\n\npid:933380459\nbyr:1938\ncid:291 hcl:#c0946f\necl:oth iyr:2018\neyr:2026 hgt:170cm\n\nbyr:1974\npid:262927116 eyr:2027 ecl:gry\nhcl:#341e13 iyr:2014 cid:232 hgt:161cm\n\nhcl:#602927\nbyr:2001 iyr:2011\nhgt:177cm eyr:2028 pid:165733929 ecl:amb\n\nbyr:1922 cid:144 pid:333716867 hgt:183cm iyr:2015\nhcl:#c25ea9 eyr:2022 ecl:blu\n\neyr:2021 cid:147 byr:1978\niyr:2020 pid:938828535\nhcl:#7d3b0c ecl:amb hgt:159cm\n\nhgt:153cm ecl:hzl\ncid:232 byr:1953 hcl:#a97842 iyr:2016 pid:356632792 eyr:2029\n\npid:745727684 ecl:gry iyr:2020\nhcl:#a97842\neyr:2025 cid:275\nhgt:65in\nbyr:1957\n\nhcl:#733820\necl:grn iyr:2019 byr:1943 eyr:2024 hgt:70in\npid:953607814\n\necl:gry eyr:2028 hcl:#cfa07d\nhgt:163cm\nbyr:1942 iyr:2019 pid:310104177\n\nhgt:190cm\neyr:2027 iyr:2010 byr:1978\necl:gry\nhcl:#964ba7\n\ncid:320\neyr:2022 hgt:169cm\necl:blu hcl:#a97842 iyr:2015 pid:669007078 byr:1986\n\niyr:2019 pid:901370677 hcl:7f2398 cid:305\necl:amb eyr:2011 hgt:190cm byr:1991\n\necl:brn\ncid:256 byr:1987 iyr:2017 eyr:2026 hcl:#623a2f pid:875646528\nhgt:160cm\n\nbyr:1955 pid:120131971 hcl:#18171d\nhgt:156cm\necl:blu\niyr:2011 eyr:2028\n\niyr:2020 ecl:brn cid:188\nhgt:157cm\neyr:2026\npid:504067323 hcl:#733820 byr:1982\n\ncid:102 hgt:177cm\nhcl:#733820 ecl:hzl byr:1984 pid:542750146 eyr:2028 iyr:2020\n\npid:419639528 iyr:2013 hgt:175cm ecl:blu\neyr:2026 byr:1999 hcl:#733820\n\nbyr:1963 eyr:2020\npid:683641152 ecl:gry cid:207 hgt:180cm\nhcl:#cfa07d\niyr:2020\n\nhgt:192cm pid:156436859 iyr:2020 hcl:#cfa07d\necl:blu byr:1963 eyr:2025 cid:147\n\neyr:2002\nhcl:z iyr:2011\npid:6830168962\nhgt:156in cid:288 byr:2029\n\neyr:2021\npid:277739802 byr:1992 ecl:hzl iyr:2020\nhcl:#7c5fe8 hgt:184cm\n\nbyr:1989 pid:066973099\niyr:2017\neyr:2022 ecl:hzl hcl:#888785 hgt:76in\n\nhcl:#866857\niyr:2016 cid:306\necl:hzl\npid:453816800 byr:1971 hgt:71in eyr:2030\n\npid:248573931 hcl:#cfa07d\niyr:2014 eyr:2024 hgt:186cm byr:1970 cid:128 ecl:blu\n\npid:172567579 ecl:brn iyr:2014 byr:1948 cid:309\nhgt:151cm hcl:#888785 eyr:2024\n\nhgt:153cm eyr:2026 byr:1929 ecl:hzl pid:684760742\nhcl:#c45e93 iyr:2018\n\npid:#d50a43\niyr:1940\necl:#7880a9 byr:2018 hcl:dc2fa7 hgt:185in eyr:1978\n\nhcl:#602927 cid:71 eyr:2020\npid:620634584 hgt:157cm byr:1991\niyr:2020 ecl:amb\n\neyr:2023\nbyr:1959 iyr:1947 hgt:152cm ecl:#503286 pid:63978523 hcl:57dd0d\n\nhgt:190cm\nbyr:1955 ecl:blu\npid:507892696\nhcl:#9bd1f0 eyr:2029\niyr:2010\n\npid:365539813\neyr:2022 hcl:#623a2f iyr:2020 hgt:184cm\necl:oth byr:1920 cid:213\n\ncid:50 ecl:oth pid:774859218 hgt:193cm\niyr:2017 byr:1925 hcl:#866857\neyr:2021\n\nhgt:189cm\niyr:2019 byr:1937\nhcl:#a97842\neyr:2025 ecl:oth\npid:787390180\n\niyr:2019 eyr:2027 hgt:183cm\necl:hzl pid:549757712\nbyr:1956\nhcl:#866857\n\npid:755580715\nhcl:#602927 hgt:187cm iyr:2017 byr:1925 eyr:2020 ecl:blu\n\niyr:2019 hgt:69in\necl:amb\nhcl:#602927 eyr:2026\npid:951019647 byr:1974\n\nbyr:1943 eyr:2034 hgt:150 pid:#36aedf ecl:oth\nhcl:z\n\neyr:2024\necl:hzl pid:824745692 iyr:2012 hcl:06ab6e\nbyr:1944\nhgt:159cm\ncid:183\n\nhgt:169cm ecl:blu\neyr:2030 iyr:2013 byr:1945 pid:791359040 hcl:#7d3b0c\n\niyr:2018\necl:hzl hgt:152cm\nhcl:#18171d eyr:2026 byr:1924 pid:534667048\n\neyr:2029 pid:933295825\niyr:2011\nhcl:#cfa07d byr:1981\nhgt:164cm ecl:grn\n\necl:amb byr:1964 iyr:2018\npid:014457573\ncid:152\neyr:2028 hgt:171cm hcl:#866857\n\nhgt:167cm\nbyr:1974 iyr:2012 ecl:amb pid:512315114\ncid:278\neyr:2028 hcl:#623a2f\n\nhgt:153cm ecl:oth iyr:2012\neyr:2027 hcl:#888785 byr:1999 pid:416990697\n\neyr:2025 ecl:blu byr:1991 hcl:#866857\nhgt:189cm pid:546461828\n\niyr:2016\nbyr:1988\nhgt:160cm eyr:2025 ecl:amb hcl:#602927\npid:562766105\n\necl:oth byr:1942\nhcl:#341e13 pid:564975864 cid:158\nhgt:159cm eyr:2028\niyr:2018\n\npid:406209763 hgt:170cm cid:331\niyr:2018 eyr:2026 byr:1981\nhcl:#733820 ecl:gry\n\npid:279164109 ecl:oth\ncid:197 hcl:#7d3b0c\neyr:2024\nhgt:185cm iyr:2020 byr:1925\n\nhcl:#efcc98 ecl:hzl\ncid:92 hgt:190cm pid:724466265 iyr:2020\neyr:2025 byr:1996\n\nbyr:1996\ncid:55 pid:906572505 ecl:grn eyr:2022 hcl:#602927 hgt:160cm iyr:2014\n\neyr:2028 hcl:#b6652a ecl:hzl hgt:186cm iyr:2016 pid:132872161 byr:1932\n\nhcl:#fffffd iyr:2019 eyr:2020 hgt:188cm\nbyr:1951 ecl:brn\npid:842126902\n\nhcl:#602927\nhgt:158cm\neyr:2023 iyr:2010\npid:681061896 byr:1977 ecl:gry\n\niyr:2018 hgt:192cm byr:1970 cid:200 ecl:grn eyr:2027\npid:164408694 hcl:#888785\n\neyr:2029\npid:447061655 iyr:2010 hcl:#341e13 ecl:oth\ncid:187 hgt:185cm byr:1943\n\nbyr:1925 iyr:2012 eyr:2025\nhgt:190cm hcl:#18171d pid:017534154 ecl:brn\n\nhgt:172cm byr:1923\neyr:2026 iyr:2015\npid:580812884 hcl:#c0946f ecl:hzl\n\nhcl:#888785 eyr:2028\nbyr:1952 ecl:brn pid:818889983\niyr:2010 hgt:180cm\n\neyr:2026 ecl:gry byr:1982 hgt:188cm hcl:#c0946f pid:610689703 iyr:2011\n\neyr:2028\niyr:2018\npid:921660781 ecl:amb\nhcl:#cfa07d hgt:178cm byr:1975\n\nbyr:1977 pid:667631009 iyr:2010\ncid:86 eyr:2022 hgt:189cm hcl:#7d3b0c ecl:oth\n\npid:214679440 hgt:190cm ecl:blu iyr:2017\neyr:2025 cid:292\n\necl:amb\niyr:2017 hcl:531ad3\nhgt:163 pid:689027667 byr:2006 eyr:2033\n\nhgt:68in byr:1928 iyr:2010 cid:227 eyr:2023\necl:hzl pid:#87bab9 hcl:#fffffd\n\necl:grn byr:1940 cid:294 hgt:152cm pid:310277488\niyr:2015 hcl:#18171d eyr:2030\n\nbyr:1965 pid:240720987\neyr:2030 ecl:oth hgt:192cm hcl:#733820\niyr:2016\n\npid:830487275\necl:blu byr:1930\nhcl:#b6652a iyr:2013 hgt:188cm eyr:2025\n\nhgt:177cm byr:1955 eyr:2030 ecl:amb pid:476675886 iyr:2016 hcl:#c0946f\n\npid:152702068 iyr:2016 hcl:#b6652a\ncid:82 ecl:blu eyr:2029 byr:1975 hgt:161cm\n\npid:136852264\neyr:2024 cid:339 ecl:oth byr:1949 iyr:2011\n\niyr:2020 pid:772739059\neyr:2025 hgt:157cm\nbyr:1945 ecl:brn\nhcl:#6b5442\n\nhcl:#18171d eyr:2022\niyr:2018 ecl:grn byr:1933 pid:053763751\n\npid:214212776 hcl:#18171d\neyr:2030\niyr:2020 byr:1988\ncid:122\nhgt:170cm ecl:oth\n\npid:883116919 iyr:2018 ecl:brn byr:1938 hgt:187cm eyr:2020\n\niyr:2020 hcl:#a97842\ncid:329 eyr:2025 byr:1946 pid:636649774\necl:grn hgt:158cm\n\neyr:2023\necl:blu hgt:161cm\nhcl:#341e13 byr:1951\niyr:2020 pid:461889565 cid:97\n\nhgt:168cm pid:492241189\neyr:2029\niyr:2013\ncid:150\nbyr:1980 hcl:#cfa07d ecl:hzl\n\nbyr:1998 ecl:gry hgt:150cm eyr:2024 pid:401735295 cid:153 hcl:#733820 iyr:2016\n\necl:hzl hgt:184cm iyr:2018\nbyr:2001\npid:453480077 eyr:2025 hcl:#a97842"
