package Fake

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jaswdr/faker"
)

func (this Faker) RandomBoolean() bool {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).Bool()
}
func (this Faker) RandomInt() int {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).IntBetween(0, 1000)
}
func (this Faker) RandomStringMaxLenght(l int) string {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).RandomStringWithLength(l)
}
func (this Faker) RandomFloatBetween(maxDecimals, min, max int) float64 {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).RandomFloat(maxDecimals, min, max)
}
func (this Faker) RandomIntBetween(min, max int) int {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).IntBetween(min, max)
}
func (this Faker) RandomUUID() uuid.UUID {
	return uuid.New()
}
func (this Faker) CurrentISOTimestamp() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
}
func (this Faker) CurrentTimestamp() int64 {
	return time.Now().UTC().Unix()
}
func (this Faker) RandomGuid() uuid.UUID {
	return uuid.New()
}
func (this Faker) RandomAddressLongitude() float64 {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).Address().Longitude()
}
func (this Faker) RandomAddressLatitude() float64 {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).Address().Longitude()
}
func (this Faker) RandomAddressCountry() string {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).Address().Country()
}
func (this Faker) RandomAddressStreetAddress() string {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).Address().StreetAddress()
}
func (this Faker) RandomAddresStreetName() string {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).Address().StreetName()
}
func (this Faker) RandomAddressCity() string {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).Address().City()
}
func (this Faker) RandomPersonNameSuffix() string {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).Person().Suffix()
}
func (this Faker) RandomPersonNamePrefix() string {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).Person().Title()
}
func (this Faker) RandomPersonFullName() string {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).Person().Name()
}
func (this Faker) RandomPersonLastName() string {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).Person().LastName()
}
func (this Faker) RandomPersonFirstName() string {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).Person().FirstName()
}
func (this Faker) RandomUserAgent() string {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).UserAgent().UserAgent()
}
func (this Faker) RandomLocale() string {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).Language().Language()
}
func (this Faker) RandomPassword() string {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).Internet().Password()
}
func (this Faker) RandomMACAddress() string {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).Internet().MacAddress()
}
func (this Faker) RandomIP() string {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).Internet().Ipv4()
}
func (this Faker) RandomSafeColorHex() string {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).Color().Hex()
}
func (this Faker) RandomSafeColorName() string {
	return faker.NewWithSeed(rand.NewSource(time.Now().UnixNano())).Color().SafeColorName()
}
func (this Faker) RandomDateFuture() string {
	min := time.Now().Unix()
	max := time.Now().Unix() * 10 / 9
	delta := max - min
	sec := rand.Int63n(delta) + min
	randTime := time.Unix(sec, 0)
	return randTime.Format(time.UnixDate)
}
func (this Faker) RandomDatePast() string {
	min := time.Now().Unix() - (time.Now().Unix() * 2 / 10)
	max := time.Now().Unix()
	delta := max - min
	sec := rand.Int63n(delta) + min
	randTime := time.Unix(sec, 0)
	return randTime.Format(time.UnixDate)
}
func (this Faker) RandomDateRecent() string {
	min := time.Now().Unix() - (time.Now().Unix() * 1 / 200)
	max := time.Now().Unix()
	delta := max - min
	sec := rand.Int63n(delta) + min
	randTime := time.Unix(sec, 0)
	return randTime.Format(time.UnixDate)
}
func (this Faker) RandomLoremWord() string {
	return LoremWords[this.Generator.Intn(len(LoremWords))]
}
func (this Faker) RandomLoremWords() string {
	words := ""
	size := this.Generator.Intn(6-2) + 2
	for i := 0; i < size; i++ {
		words += this.RandomLoremWord() + " "
	}
	return words
}
func (this Faker) RandomLoremSentence() string {
	words := ""
	size := this.Generator.Intn(6-2) + 2
	for i := 0; i < size; i++ {
		words += this.RandomLoremWord() + " "
	}
	return strings.TrimSpace(words) + "."
}
func (this Faker) RandomPhrase() string {
	return this.RandomLoremSentence()
}
func (this Faker) RandomLoremSentences() string {
	words := ""
	size := this.Generator.Intn(6-2) + 2
	for i := 0; i < size; i++ {
		words += this.RandomLoremSentence() + " "
	}
	return words
}
func (this Faker) RandomLoremLines() string {
	words := ""
	size := this.Generator.Intn(6-2) + 2
	for i := 0; i < size; i++ {
		words += this.RandomLoremSentence() + "\n"
	}
	return words
}
func (this Faker) RandomLoremParagraph() string {
	words := ""
	size := this.Generator.Intn(6-2) + 2
	for i := 0; i < size; i++ {
		words += this.RandomLoremSentence() + "\n"
	}
	return words
}
func (this Faker) RandomLoremText() string {
	return this.RandomLoremParagraphs()
}
func (this Faker) RandomLoremParagraphs() string {
	words := ""
	size := this.Generator.Intn(6-2) + 2
	for i := 0; i < size; i++ {
		words += this.RandomLoremSentence() + "\n"
	}
	return words
}
func (this Faker) RandomLoremSlug() string {
	return this.RandomLoremWord() + "-" + this.RandomLoremWord() + "-" + this.RandomLoremWord()
}
func (this Faker) RandomNoun() string {
	return Nouns[this.Generator.Intn(len(Nouns))]
}
func (this Faker) RandomVerb() string {
	return Verbs[this.Generator.Intn(len(Verbs))]
}
func (this Faker) RandomIngVerb() string {
	return IngVerbs[this.Generator.Intn(len(IngVerbs))]
}
func (this Faker) RandomAdjective() string {
	return Adjectives[this.Generator.Intn(len(Adjectives))]
}
func (this Faker) RandomWord() string {
	words := append(append(append(Adjectives, IngVerbs...), Verbs...), Nouns...)
	return words[this.Generator.Intn(len(words))]
}
func (this Faker) RandomWords() string {
	words := ""
	size := this.Generator.Intn(6-2) + 2
	for i := 0; i < size; i++ {
		words += this.RandomWord() + " "
	}
	return words
}
func (this Faker) RandomDepartment() string {
	return StoreDepartaments[this.Generator.Intn(len(StoreDepartaments))]
}
func (this Faker) RandomProductName() string {
	return this.RandomAdjective() + " " + this.RandomProductMaterial() + " " + this.RandomProduct()
}
func (this Faker) RandomProductMaterial() string {
	return ProductMaterials[this.Generator.Intn(len(ProductMaterials))]
}
func (this Faker) RandomProductAdjective() string {
	return ProductAdjectives[this.Generator.Intn(len(ProductAdjectives))]
}
func (this Faker) RandomProduct() string {
	return Products[this.Generator.Intn(len(Products))]
}
func (this Faker) RandomPrice() string {
	return strconv.Itoa(this.Generator.Intn(1000)) + " " + strconv.Itoa(this.Generator.Intn(99-10)+100)
}
func (this Faker) RandomFilePath() string {
	return this.RandomDirectoryPath() + "/" + this.RandomDomainWord()
}
func (this Faker) RandomMimeType() string {
	return CommonMimeTypes[this.Generator.Intn(len(CommonMimeTypes))]
}
func (this Faker) RandomDirectoryPath() string {
	return DirectoryPaths[this.Generator.Intn(len(DirectoryPaths))]
}
func (this Faker) RandomCommonFileExtension() string {
	return CommonFileExtensions[this.Generator.Intn(len(CommonFileExtensions))]
}
func (this Faker) RandomCommonFileType() string {
	return CommonFIleTypes[this.Generator.Intn(len(CommonFIleTypes))]
}
func (this Faker) RandomCommonFileName() string {
	return this.RandomDomainWord() + "." + this.RandomCommonFileExtension()
}
func (this Faker) RandomFileType() string {
	return FileTypes[this.Generator.Intn(len(FileTypes))]
}
func (this Faker) RandomFileExtension() string {
	return FileExtensions[this.Generator.Intn(len(FileExtensions))]
}
func (this Faker) RandomFileName() string {
	return strings.ToLower(FirstNames[this.Generator.Intn(len(FirstNames))])+"_"+LastNames[this.Generator.Intn(len(LastNames))]+ "." this.RandomFileExtension()
}
func (this Faker) RandomUrl() string {
	return this.RandomProtocol() + "://" + strings.ToLower(FirstNames[this.Generator.Intn(len(FirstNames))]+ "_"+LastNames[this.Generator.Intn(len(LastNames))]) + "." + 
	this.RandomFileExtension()
}
func (this Faker) RandomUsername() string {
	return FirstNames[this.Generator.Intn(len(FirstNames))]+"."+LastNames[this.Generator.Intn(len(LastNames))]
}
func (this Faker) RandomExampleEmail() string {
		return stirngs.ToLower(FirstNames[this.Generator.Intn(len(FirstNames))]+"_"+LastNames[this.Generator.Intn(len(LastNames))])+"@"+this.RandomDomainName()
}
func (this Faker) RandomEmail() string {

}
func (this Faker) RandomDomainWord() string {

}
func (this Faker) RandomDomainSuffix() string {

}
func (this Faker) RandomDomainName() string {

}
func (this Faker) RandomMonth() string {

}
func (this Faker) RandomWeekday() string {

}
func (this Faker) RandomDatabaseColumn() string {

}
func (this Faker) RandomDatabaseType() string {

}
func (this Faker) RandomDatabaseCollation() string {

}
func (this Faker) RandomDatabaseEngine() string {

}
func (this Faker) RandomCatchPhrase() string {

}
func (this Faker) RandomCatchPhraseAdjective() string {

}
func (this Faker) RandomCatchPhraseDescriptor() string {

}
func (this Faker) RandomCatchPhraseNoun() string {

}
func (this Faker) RandomBsNoun() string {

}
func (this Faker) RandomBsBuzzWord() string {

}
func (this Faker) RandomBsAdjective() string {

}
func (this Faker) RandomBs() string {

}
func (this Faker) RandomCompanySuffix() string {

}
func (this Faker) RandomCompanyName() string {

}
func (this Faker) RandomBitcoin() string {

}
func (this Faker) RandomCurrencySymbol() string {

}
func (this Faker) RandomCurrencyCode() string {

}
func (this Faker) RandomCurrencyName() string {

}
func (this Faker) RandomTransactionType() string {

}
func (this Faker) RandomCreditCardMask() string {

}
func (this Faker) RandomBankAccountName() string {

}
func (this Faker) RandomBankAccountBic() string {

}
func (this Faker) RandomBankAccountIban() string {

}
func (this Faker) RandomBankAccount() string {

}
func (this Faker) RandomDataImageUri() string {

}
func (this Faker) RandomAvatarImage() string {

}
func (this Faker) RandomImageURL() string {

}
func (this Faker) RandomAbstractImage() string {

}
func (this Faker) RandomAnimalsImage() string {

}
func (this Faker) RandomBusinessImage() string {

}
func (this Faker) RandomCatsImage() string {

}
func (this Faker) RandomCityImage() string {

}
func (this Faker) RandomFoodImage() string {

}
func (this Faker) RandomNightlifeImage() string {

}
func (this Faker) RandomFashionImage() string {

}
func (this Faker) RandomPeopleImage() string {

}
func (this Faker) RandomNatureImage() string {

}
func (this Faker) RandomSportsImage() string {

}
func (this Faker) RandomTransportImage() string {

}
func (this Faker) RandomCountryCode() string {

}
func (this Faker) RandomPhoneNumber() string {

}
func (this Faker) RandomPhoneNumberExt() string {

}
func (this Faker) RandomJobArea() string {

}
func (this Faker) RandomJobDescriptor() string {

}
func (this Faker) RandomJobType() string {

}
func (this Faker) RandomJobTitle() string {

}
func (this Faker) RandomSemver() string {

}
func (this Faker) RandomProtocol() string {

}
func (this Faker) RandomAbbreviation() string {

}
func (this Faker) RandomAlphanumeric() string {

}
func (this Faker) RandomIpv6() string {

}
