package Fake

import (
	"github.com/google/uuid"
)

type IFake interface {
	TocliRandomBoolean() bool
	TocliRandomInt() int
	TocliRandomStringMaxLenght(l int) string
	TocliRandomFloatBetween(maxDecimals, min, max int) float64
	TocliRandomIntBetween(min, max int) int
	TocliRandomUUID() uuid.UUID
	TocliCurrentISOTimestamp() string
	TocliCurrentTimestamp() int64
	TocliRandomGuid() uuid.UUID
	TocliRandomAddressLongitude() float64
	TocliRandomAddressLatitude() float64
	TocliRandomAddressCountry() string
	TocliRandomAddressStreetAddress() string
	TocliRandomAddresStreetName() string
	TocliRandomAddressCity() string
	TocliRandomPersonNameSuffix() string
	TocliRandomPersonNamePrefix() string
	TocliRandomPersonFullName() string
	TocliRandomPersonLastName() string
	TocliRandomPersonFirstName() string
	TocliRandomUserAgent() string
	TocliRandomLocale() string
	TocliRandomPassword() string
	TocliRandomMACAddress() string
	TocliRandomIP() string
	TocliRandomSafeColorHex() string
	TocliRandomSafeColorName() string
	TocliRandomDateFuture() string
	TocliRandomDatePast() string
	TocliRandomDateRecent() string
	TocliRandomLoremWord() string
	TocliRandomLoremWords() string
	TocliRandomLoremSentence() string
	TocliRandomPhrase() string
	TocliRandomLoremSentences() string
	TocliRandomLoremLines() string
	TocliRandomLoremParagraph() string
	TocliRandomLoremText() string
	TocliRandomLoremParagraphs() string
	TocliRandomLoremSlug() string
	TocliRandomNoun() string
	TocliRandomVerb() string
	TocliRandomIngVerb() string
	TocliRandomAdjective() string
	TocliRandomWord() string
	TocliRandomWords() string
	TocliRandomDepartment() string
	TocliRandomProductName() string
	TocliRandomProductMaterial() string
	TocliRandomProductAdjective() string
	TocliRandomProduct() string
	TocliRandomPrice() string
	TocliRandomFilePath() string
	TocliRandomMimeType() string
	TocliRandomDirectoryPath() string
	TocliRandomCommonFileExtension() string
	TocliRandomCommonFileType() string
	TocliRandomCommonFileName() string
	TocliRandomFileExtension() string
	TocliRandomFileName() string
	TocliRandomUrl() string
	TocliRandomUsername() string
	TocliRandomExampleEmail() string
	TocliRandomEmail() string
	TocliRandomDomainWord() string
	TocliRandomDomainSuffix() string
	TocliRandomDomainName() string
	TocliRandomMonth() string
	TocliRandomWeekday() string
	TocliRandomDatabaseColumn() string
	TocliRandomDatabaseType() string
	TocliRandomDatabaseCollation() string
	TocliRandomDatabaseEngine() string
	TocliRandomCatchPhrase() string
	TocliRandomCatchPhraseAdjective() string
	TocliRandomCatchPhraseDescriptor() string
	TocliRandomCatchPhraseNoun() string
	TocliRandomBsNoun() string
	TocliRandomBsBuzzWord() string
	TocliRandomBsAdjective() string
	TocliRandomBs() string
	TocliRandomCompanySuffix() string
	TocliRandomCompanyName() string
	TocliRandomBitcoin() string
	TocliRandomCurrencySymbol() string
	TocliRandomCurrencyCode() string
	TocliRandomCurrencyName() string
	TocliRandomTransactionType() string
	TocliRandomCreditCardMask() string
	TocliRandomBankAccountName() string
	TocliRandomBankAccountBic() string
	TocliRandomBankAccountIban() string
	TocliRandomBankAccount() string
	TocliRandomDataImageUri() string
	TocliRandomAvatarImage() string
	TocliRandomImageURL() string
	TocliRandomAbstractImage() string
	TocliRandomAnimalsImage() string
	TocliRandomBusinessImage() string
	TocliRandomCatsImage() string
	TocliRandomCityImage() string
	TocliRandomFoodImage() string
	TocliRandomNightlifeImage() string
	TocliRandomFashionImage() string
	TocliRandomPeopleImage() string
	TocliRandomNatureImage() string
	TocliRandomSportsImage() string
	TocliRandomTransportImage() string
	TocliRandomCountryCode() string
	TocliRandomPhoneNumber() string
	TocliRandomPhoneNumberExt() string
	TocliRandomJobArea() string
	TocliRandomJobDescriptor() string
	TocliRandomJobType() string
	TocliRandomJobTitle() string
	TocliRandomSemver() string
	TocliRandomProtocol() string
	TocliRandomAbbreviation() string
	TocliRandomAlphanumeric() string
	TocliRandomIpv6() string
	TocliRandomDigitNotNull() int
	TocliRandomDigit() int
	TocliRandomFloat() float64
	TocliRandomString() string
	TocliRandomStringWithLength(I int) string
	TocliRandomLetter() string
}
