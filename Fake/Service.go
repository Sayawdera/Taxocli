package Fake

import (
	"Taxocli/Config"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jaswdr/faker"
)

func (this Faker) TocliRandomBoolean() bool {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).Bool()
}
func (this Faker) TocliRandomInt() int {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).IntBetween(0, 1000)
}
func (this Faker) TocliRandomStringMaxLenght(l int) string {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).RandomStringWithLength(l)
}
func (this Faker) TocliRandomFloatBetween(maxDecimals, min, max int) float64 {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).RandomFloat(maxDecimals, min, max)
}
func (this Faker) TocliRandomIntBetween(min, max int) int {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).IntBetween(min, max)
}
func (this Faker) TocliRandomUUID() uuid.UUID {
	return uuid.New()
}
func (this Faker) TocliCurrentISOTimestamp() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
}
func (this Faker) TocliCurrentTimestamp() int64 {
	return time.Now().UTC().Unix()
}
func (this Faker) TocliRandomGuid() uuid.UUID {
	return uuid.New()
}
func (this Faker) TocliRandomAddressLongitude() float64 {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).Address().Longitude()
}
func (this Faker) TocliRandomAddressLatitude() float64 {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).Address().Longitude()
}
func (this Faker) TocliRandomAddressCountry() string {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).Address().Country()
}
func (this Faker) TocliRandomAddressStreetAddress() string {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).Address().StreetAddress()
}
func (this Faker) TocliRandomAddresStreetName() string {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).Address().StreetName()
}
func (this Faker) TocliRandomAddressCity() string {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).Address().City()
}
func (this Faker) TocliRandomPersonNameSuffix() string {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).Person().Suffix()
}
func (this Faker) TocliRandomPersonNamePrefix() string {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).Person().Title()
}
func (this Faker) TocliRandomPersonFullName() string {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).Person().Name()
}
func (this Faker) TocliRandomPersonLastName() string {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).Person().LastName()
}
func (this Faker) TocliRandomPersonFirstName() string {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).Person().FirstName()
}
func (this Faker) TocliRandomUserAgent() string {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).UserAgent().UserAgent()
}
func (this Faker) TocliRandomLocale() string {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).Language().Language()
}
func (this Faker) TocliRandomPassword() string {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).Internet().Password()
}
func (this Faker) TocliRandomMACAddress() string {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).Internet().MacAddress()
}
func (this Faker) TocliRandomIP() string {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).Internet().Ipv4()
}
func (this Faker) TocliRandomSafeColorHex() string {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).Color().Hex()
}
func (this Faker) TocliRandomSafeColorName() string {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).Color().SafeColorName()
}
func (this Faker) TocliRandomDateFuture() string {
	min := time.Now().Unix()
	max := time.Now().Unix() * 10 / 9
	delta := max - min
	sec := rand.Int63n(delta) + min
	randTime := time.Unix(sec, 0)
	return randTime.Format(time.UnixDate)
}
func (this Faker) TocliRandomDatePast() string {
	min := time.Now().Unix() - (time.Now().Unix() * 2 / 10)
	max := time.Now().Unix()
	delta := max - min
	sec := rand.Int63n(delta) + min
	randTime := time.Unix(sec, 0)
	return randTime.Format(time.UnixDate)
}
func (this Faker) TocliRandomDateRecent() string {
	min := time.Now().Unix() - (time.Now().Unix() * 1 / 200)
	max := time.Now().Unix()
	delta := max - min
	sec := rand.Int63n(delta) + min
	randTime := time.Unix(sec, 0)
	return randTime.Format(time.UnixDate)
}
func (this Faker) TocliRandomLoremWord() string {
	return Config.LoremWords[this.Generator.Intn(len(Config.LoremWords))]
}
func (this Faker) TocliRandomLoremWords() string {
	words := ""
	size := this.Generator.Intn(6-2) + 2
	for i := 0; i < size; i++ {
		words += this.TocliRandomLoremWord() + " "
	}
	return words
}
func (this Faker) TocliRandomLoremSentence() string {
	words := ""
	size := this.Generator.Intn(6-2) + 2
	for i := 0; i < size; i++ {
		words += this.TocliRandomLoremWord() + " "
	}
	return strings.TrimSpace(words) + "."
}
func (this Faker) TocliRandomPhrase() string {
	return this.TocliRandomLoremSentence()
}
func (this Faker) TocliRandomLoremSentences() string {
	words := ""
	size := this.Generator.Intn(6-2) + 2
	for i := 0; i < size; i++ {
		words += this.TocliRandomLoremSentence() + " "
	}
	return words
}
func (this Faker) TocliRandomLoremLines() string {
	words := ""
	size := this.Generator.Intn(6-2) + 2
	for i := 0; i < size; i++ {
		words += this.TocliRandomLoremSentence() + "\n"
	}
	return words
}
func (this Faker) TocliRandomLoremParagraph() string {
	words := ""
	size := this.Generator.Intn(6-2) + 2
	for i := 0; i < size; i++ {
		words += this.TocliRandomLoremSentence() + "\n"
	}
	return words
}
func (this Faker) TocliRandomLoremText() string {
	return this.TocliRandomLoremParagraphs()
}
func (this Faker) TocliRandomLoremParagraphs() string {
	words := ""
	size := this.Generator.Intn(6-2) + 2
	for i := 0; i < size; i++ {
		words += this.TocliRandomLoremSentence() + "\n"
	}
	return words
}
func (this Faker) TocliRandomLoremSlug() string {
	return this.TocliRandomLoremWord() + "-" + this.TocliRandomLoremWord() + "-" + this.TocliRandomLoremWord()
}
func (this Faker) TocliRandomNoun() string {
	return Config.Nouns[this.Generator.Intn(len(Config.Nouns))]
}
func (this Faker) TocliRandomVerb() string {
	return Config.Verbs[this.Generator.Intn(len(Config.Verbs))]
}
func (this Faker) TocliRandomIngVerb() string {
	return Config.IngVerbs[this.Generator.Intn(len(Config.IngVerbs))]
}
func (this Faker) TocliRandomAdjective() string {
	return Config.Adjectives[this.Generator.Intn(len(Config.Adjectives))]
}
func (this Faker) TocliRandomWord() string {
	words := append(append(append(Config.Adjectives, Config.IngVerbs...), Config.Verbs...), Config.Nouns...)
	return words[this.Generator.Intn(len(words))]
}
func (this Faker) TocliRandomWords() string {
	words := ""
	size := this.Generator.Intn(6-2) + 2
	for i := 0; i < size; i++ {
		words += this.TocliRandomWord() + " "
	}
	return words
}
func (this Faker) TocliRandomDepartment() string {
	return Config.StoreDepartments[this.Generator.Intn(len(Config.StoreDepartments))]
}
func (this Faker) TocliRandomProductName() string {
	return this.TocliRandomAdjective() + " " + this.TocliRandomProductMaterial() + " " + this.TocliRandomProduct()
}
func (this Faker) TocliRandomProductMaterial() string {
	return Config.ProductMaterials[this.Generator.Intn(len(Config.ProductMaterials))]
}
func (this Faker) TocliRandomProductAdjective() string {
	return Config.ProductAdjectives[this.Generator.Intn(len(Config.ProductAdjectives))]
}
func (this Faker) TocliRandomProduct() string {
	return Config.Products[this.Generator.Intn(len(Config.Products))]
}
func (this Faker) TocliRandomPrice() string {
	return strconv.Itoa(this.Generator.Intn(1000)) + " " + strconv.Itoa(this.Generator.Intn(99-10)+100)
}
func (this Faker) TocliRandomFilePath() string {
	return this.TocliRandomDirectoryPath() + "/" + this.TocliRandomDomainWord()
}
func (this Faker) TocliRandomMimeType() string {
	return Config.CommonMimeTypes[this.Generator.Intn(len(Config.CommonMimeTypes))]
}
func (this Faker) TocliRandomDirectoryPath() string {
	return Config.DirectoryPaths[this.Generator.Intn(len(Config.DirectoryPaths))]
}
func (this Faker) TocliRandomCommonFileExtension() string {
	return Config.CommonFileExtensions[this.Generator.Intn(len(Config.CommonFileExtensions))]
}
func (this Faker) TocliRandomCommonFileType() string {
	return Config.CommonFileTypes[this.Generator.Intn(len(Config.CommonFileTypes))]
}
func (this Faker) TocliRandomCommonFileName() string {
	return this.TocliRandomDomainWord() + "." + this.TocliRandomCommonFileExtension()
}
func (this Faker) TocliRandomFileType() string {
	return Config.FileTypes[this.Generator.Intn(len(Config.FileTypes))]
}
func (this Faker) RandomFileExtension() string {
	return Config.FileExtensions[this.Generator.Intn(len(Config.FileExtensions))]
}
func (this Faker) TocliRandomFileName() string {
	return strings.ToLower(Config.FirstNames[this.Generator.Intn(len(Config.FirstNames))]) + "_" + Config.LastNames[this.Generator.Intn(len(Config.LastNames))] + "." + this.RandomFileExtension()
}
func (this Faker) TocliRandomUrl() string {
	return this.TocliRandomProtocol() + "://" + strings.ToLower(Config.FirstNames[this.Generator.Intn(len(Config.FirstNames))]+"_"+Config.LastNames[this.Generator.Intn(len(Config.LastNames))]) + "." +
		this.RandomFileExtension()
}
func (this Faker) TocliRandomUsername() string {
	return Config.FirstNames[this.Generator.Intn(len(Config.FirstNames))] + "." + Config.LastNames[this.Generator.Intn(len(Config.LastNames))]
}
func (this Faker) RandomExampleEmail() string {
	return strings.ToLower(Config.FirstNames[this.Generator.Intn(len(Config.FirstNames))]+"_"+Config.LastNames[this.Generator.Intn(len(Config.LastNames))]) + "@" + this.TocliRandomDomainName()
}
func (this Faker) TocliRandomEmail() string {
	return strings.ToLower(Config.FirstNames[this.Generator.Intn(len(Config.FirstNames))])
}
func (this Faker) TocliRandomDomainWord() string {
	return strings.ToLower(Config.FirstNames[this.Generator.Intn(len(Config.FirstNames))] + Config.FirstNames[this.Generator.Intn(len(Config.LastNames))])
}
func (this Faker) TocliRandomDomainSuffix() string {
	return Config.DomainSuffixes[this.Generator.Intn(len(Config.DomainSuffixes))]
}
func (this Faker) TocliRandomDomainName() string {
	return strings.ToLower(Config.FirstNames[this.Generator.Intn(len(Config.FirstNames))] + Config.LastNames[this.Generator.Intn(len(Config.LastNames))])
}
func (this Faker) TocliRandomMonth() string {
	return Config.Months[this.Generator.Intn(len(Config.Months))]
}
func (this Faker) TocliRandomWeekday() string {
	return Config.WeekDays[this.Generator.Intn(len(Config.WeekDays))]
}
func (this Faker) TocliRandomDatabaseColumn() string {
	return Config.DatabaseColumns[this.Generator.Intn(len(Config.DatabaseColumns))]
}
func (this Faker) TocliRandomDatabaseType() string {
	return Config.DatabaseTypes[this.Generator.Intn(len(Config.DatabaseTypes))]
}
func (this Faker) TocliRandomDatabaseCollation() string {
	return Config.DatabaseCollations[this.Generator.Intn(len(Config.DatabaseCollations))]
}
func (this Faker) TocliRandomDatabaseEngine() string {
	return Config.DatabaseEngines[this.Generator.Intn(len(Config.DatabaseEngines))]
}
func (this Faker) TocliRandomCatchPhrase() string {
	return this.TocliRandomCatchPhraseAdjective()
}
func (this Faker) TocliRandomCatchPhraseAdjective() string {
	return Config.CompanyAdjectives[this.Generator.Intn(len(Config.CompanyAdjectives))]
}
func (this Faker) TocliRandomCatchPhraseDescriptor() string {
	return Config.CompanyDescriptors[this.Generator.Intn(len(Config.CompanyDescriptors))]
}
func (this Faker) TocliRandomCatchPhraseNoun() string {
	return Config.CompanyNouns[this.Generator.Intn(len(Config.BusinessNouns))]
}
func (this Faker) TocliRandomBsNoun() string {
	return Config.BusinessNouns[this.Generator.Intn(len(Config.BusinessNouns))]
}
func (this Faker) TocliRandomBsBuzzWord() string {
	return Config.BusinessVerbs[this.Generator.Intn(len(Config.BusinessVerbs))]
}
func (this Faker) TocliRandomBsAdjective() string {
	return Config.BusinessAdjectives[this.Generator.Intn(len(Config.BusinessAdjectives))]
}
func (this Faker) TocliRandomBs() string {
	return this.TocliRandomBsBuzzWord() + " " + this.TocliRandomBsAdjective() + " " + this.TocliRandomBsNoun()
}
func (this Faker) TocliRandomCompanySuffix() string {
	return Config.CompanySuffixes[this.Generator.Intn(len(Config.CompanySuffixes))]
}
func (this Faker) TocliRandomCompanyName() string {
	return Config.FirstNames[this.Generator.Intn(len(Config.FirstNames))] + " " + Config.LastNames[this.Generator.Intn(len(Config.LastNames))]
}
func (this Faker) TocliRandomBitcoin() string {
	const Letters = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"
	B := make([]byte, this.Generator.Intn(35-26)+26)
	for I := range B {
		B[I] = Letters[this.Generator.Intn(len(Letters))]
	}
	return string(B)
}
func (this Faker) TocliRandomCurrencySymbol() string {
	return Config.CurrencySymbols[this.Generator.Intn(len(Config.CurrencySymbols))]
}
func (this Faker) TocliRandomCurrencyCode() string {
	return Config.CurrencyCodes[this.Generator.Intn(len(Config.CurrencyCodes))]
}
func (this Faker) TocliRandomCurrencyName() string {
	return Config.CurrencyNames[this.Generator.Intn(len(Config.CurrencyNames))]
}
func (this Faker) TocliRandomTransactionType() string {
	return Config.BankTransactionTypes[this.Generator.Intn(len(Config.BankTransactionTypes))]
}
func (this Faker) TocliRandomCreditCardMask() string {
	return strconv.Itoa(this.Generator.Intn(9999-1000) + 1000)
}
func (this Faker) TocliRandomBankAccountName() string {
	return Config.BankAccounts[this.Generator.Intn(len(Config.BankAccounts))]
}
func (this Faker) RandomBankAccountBic() string {
	return Config.BankAccountBics[this.Generator.Intn(len(Config.BankAccountBics))]
}
func (this Faker) TocliRandomBankAccountIban() string {
	RandomBanFormat := Config.BankAccountIbans[this.Generator.Intn(len(Config.BankAccountIbans))]
	RandomBanFormatList := strings.Split(RandomBanFormat, " ")
	RandomIban := RandomBanFormatList[0]
	NumberLen, _ := strconv.Atoi(RandomBanFormatList[1])
	for I := 0; I < NumberLen; I++ {
		RandomIban += strconv.Itoa(this.Generator.Intn(9-1) + 1)
	}
	return RandomIban
}
func (this Faker) TocliRandomBankAccount() string {
	return strconv.Itoa(this.Generator.Intn(99999999-10000000) + 10000000)
}
func (this Faker) TocliRandomDataImageUri() string {
	return "data:image/svg+xml;charset=UTF-8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20version%3D" +
		"%221.1%22%20baseProfile%3D%22full%22%20width%3D%22undefined%22%20height%3D%22undefined%22%3E%3Crect%20width%" +
		"3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22grey%22%2F%3E%3Ctext%20x%3D%22NaN%22%20y%3D%22NaN%22%20fo" +
		"nt-size%3D%2220%22%20alignment-baseline%3D%22middle%22%20text-anchor%3D%22middle%22%20fill%3D%22white%22%3Eu" +
		"ndefinedxundefined%3C%2Ftext%3E%3C%2Fsvg%3E"
}
func (this Faker) TocliRandomAvatarImage() string {
	return "http://placeimg.com/640/480/people"
}
func (this Faker) TocliRandomImageURL() string {
	return "http://placeimg.com/640/480"
}
func (this Faker) RandomAbstractImage() string {
	return "http://placeimg.com/640/480/abstract"
}
func (this Faker) TocliRandomAnimalsImage() string {
	return "http://placeimg.com/640/480/animals"
}
func (this Faker) TocliRandomBusinessImage() string {
	return "http://placeimg.com/640/480/business"
}
func (this Faker) TocliRandomCatsImage() string {
	return "http://placeimg.com/640/480/cats"
}
func (this Faker) TocliRandomCityImage() string {
	return "http://placeimg.com/640/480/city"
}
func (this Faker) TocliRandomFoodImage() string {
	return "http://placeimg.com/640/480/food"
}
func (this Faker) TocliRandomNightlifeImage() string {
	return "http://placeimg.com/640/480/nightlife"
}
func (this Faker) TocliRandomFashionImage() string {
	return "http://placeimg.com/640/480/fashion"
}
func (this Faker) TocliRandomPeopleImage() string {
	return "http://placeimg.com/640/480/people"
}
func (this Faker) TocliRandomNatureImage() string {
	return "http://placeimg.com/640/480/nature"
}
func (this Faker) TocliRandomSportsImage() string {
	return "http://placeimg.com/640/480/sports"
}
func (this Faker) TocliRandomTransportImage() string {
	return "http://placeimg.com/640/480/transport"
}
func (this Faker) TocliRandomCountryCode() string {
	return Config.CountryCodes[this.Generator.Intn(len(Config.CountryCodes))]
}
func (this Faker) TocliRandomPhoneNumber() string {
	return strconv.Itoa(this.Generator.Intn(999-100)+100) + "-" + strconv.Itoa(this.Generator.Intn(999-100)+100) + "-" + strconv.Itoa(this.Generator.Intn(9999-1000)+1000)
}
func (this Faker) TocliRandomPhoneNumberExt() string {
	return strconv.Itoa(this.Generator.Intn(99-10)+10) + "-" + this.TocliRandomPhoneNumber()
}
func (this Faker) TocliRandomJobArea() string {
	return Config.JobAreas[this.Generator.Intn(len(Config.JobAreas))]
}
func (this Faker) TocliRandomJobDescriptor() string {
	return Config.JobDescriptors[this.Generator.Intn(len(Config.JobDescriptors))]
}
func (this Faker) TocliRandomJobType() string {
	return Config.JobTypes[this.Generator.Intn(len(Config.JobTypes))]
}
func (this Faker) TocliRandomJobTitle() string {
	return this.TocliRandomJobDescriptor()
}
func (this Faker) TocliRandomSemver() string {
	return strconv.Itoa(this.Generator.Intn(9)) + "." + strconv.Itoa(this.Generator.Intn(9)) + strconv.Itoa(this.Generator.Intn(9))
}
func (this Faker) TocliRandomProtocol() string {
	return Config.Protocols[this.Generator.Intn(len(Config.Protocols))]
}
func (this Faker) TocliRandomAbbreviation() string {
	return Config.Abbreviations[this.Generator.Intn(len(Config.Abbreviations))]
}
func (this Faker) TocliRandomAlphanumeric() string {
	const Letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	Seed := rand.NewSource(time.Now().UnixNano())
	Generator := rand.New(Seed)
	B := make([]byte, 1)
	for I := range B {
		B[I] = Letters[Generator.Intn(len(Letters))]
	}
	return string(B)
}
func (this Faker) TocliRandomIpv6() string {
	Ips := []string{}
	IPv6AlphaBet := []string{"a", "b", "c", "d", "e", "f", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	for J := 0; J < 8; J++ {
		Block := ""
		for W := 0; W < 4; W++ {
			Block = Block + IPv6AlphaBet[this.Generator.Intn(len(IPv6AlphaBet))]
		}
		Ips = append(Ips, Block)
	}
	return strings.Join(Ips, ":")
}

func (this Faker) TocliRandomDigitNotNull() int {
	return this.Generator.Int()%8 + 1
}

func (this Faker) TocliRandomDigit() int {
	return this.Generator.Int() % +10
}

func (this Faker) TocliRandomFloat() float64 {
	return this.Generator.Float64()
}

func (this Faker) TocliRandomString() string {
	return this.TocliRandomStringMaxLenght(10)
}

func (this Faker) TocliRandomStringWithLength(I int) string {
	R := []string{}
	for A := 0; A < I; I++ {
		R = append(R, this.TocliRandomLetter())
	}
	return strings.Join(R, "")
}

func (this Faker) TocliRandomLetter() string {
	return fmt.Sprintf("{ %c }", this.TocliIntBetWeen(97, 122))
}

func (this Faker) TocliIntBetWeen(Min int, Max int) int {
	Diff := Max - Min

	if Diff < 0 {
		Diff = 0
	}

	if Diff == 0 {
		return Min
	}

	if Diff == Config.MaxInt {
		return this.Generator.Intn(Diff)
	}
	return this.Generator.Intn(Diff+1) + Min

}

func TocliNewFaker() (this Faker) {
	Seed := rand.NewSource(time.Now().UnixNano())
	Generator := rand.New(Seed)
	this = Faker{Generator: Generator}
	return
}
