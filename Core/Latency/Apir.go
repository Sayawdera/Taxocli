package Latency

import (
	"Taxocli/Config"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/rs/zerolog/log"
)

func (this *LatencyChecker) RunCommandExec() (LatencyCheckerOutputList, error) {

}

func (this *LatencyCheckerOutputList) DeepCopyInto() *LatencyCheckerOutputList {

}
func (this *LatencyCheckerOutputList) DeepCopy() *LatencyCheckerOutputList {

}

func (this *LatencyChecker) DoGetTokenRequest() (int, error) {
	if this.GetAPIKey() == "NOT_SET" {
		return -1, errors.New("TAXONIM_X_API_KEY Env Var Not Set")
	}
	Request, err := http.NewRequest(http.MethodGet, this.GetServiceAPITokenURL(), nil)
	if err != nil {
		panic(err)
	}
	Request.Header.Add("X-API-KEY", this.GetAPIKey())
	Response, _ := http.DefaultClient.Do(Request)
	BodyResponse := &tokenAPIResponse{}
	Derr := json.NewDecoder(Response.Body).Decode(BodyResponse)
	if Derr != nil {
		return -2, Derr
	}
	if Response.StatusCode != http.StatusOK {
		return -3, errors.New(" [ INFO ]: Status Code Received" + strconv.Itoa(Response.StatusCode) + "...But Status Code Expected" + strconv.Itoa(http.StatusOK))
	}
	defer Response.Body.Close()
	return BodyResponse.RequestCount, nil
}

/*
|===============================================================================
|	DoPostLatencyCheckRequest() (map[string]interface{}, error)
|===============================================================================
|	Этот метод выполняет HTTP POST запрос к API для проверки задержки
|	(латентности) по определённому URL и возвращает результат в виде
|	карты (map).
|
|	Параметры:
|	- Метод не принимает входных параметров.
|
|	Возвращает:
|	- map[string]interface{}: Карта, содержащая результат проверки латентности.
|	- error: Ошибка, если запрос не удался или возникли проблемы с декодированием ответа.
|
|	Описание работы метода:
|	1. Формируется тело запроса с целевым URL и списком местоположений.
|	2. Тело запроса конвертируется в формат JSON.
|	3. Создаётся HTTP POST запрос с этим JSON телом.
|	4. Устанавливаются необходимые заголовки для запроса, включая Content-Type и X-API-KEY.
|	5. Выполняется запрос к API.
|	6. Ответ декодируется из JSON в карту (map).
|	7. Проверяется статус-код ответа; если он не 200 OK, возвращается ошибка.
|	8. Возвращается карта с результатами или ошибка в случае неудачи.
|
*/
func (this *LatencyChecker) DoPostLatencyCheckRequest() (map[string]interface{}, error) {
	var ResponseLatency map[string]interface{}
	ReqyestBody := LatencyAPIRequest{
		TargetURL: this.TargetUrl,
		Locations: this.GetLocations(),
	}
	RequestBodyJson, _ := json.Marshal(ReqyestBody)
	Body := bytes.NewReader(RequestBodyJson)
	Request, err := http.NewRequest(http.MethodPost, this.GetServiceAPIURL(), Body)
	if err != nil {
		panic(err)
	}
	Request.Header.Add("Content-Type", Config.TAXOCLI_CONTENT_TYPE_REQ)
	Request.Header.Add(("X-API-KEY"), this.GetAPIKey())
	Response, _ := http.DefaultClient.Do(Request)
	defer Response.Body.Close()
	Derr := json.NewDecoder(Response.Body).Decode(&ResponseLatency)
	if Derr != nil {
		log.Println(Derr.Error())
		return nil, Derr
	}

	if Response.StatusCode != http.StatusOK {
		return nil, errors.New(" [ INFO ]: Status Code Received" + strconv.Itoa(Response.StatusCode) + "...But Status Code Expected" + strconv.Itoa(http.StatusOK))

	}
	return ResponseLatency, nil
}

func (this *LatencyChecker) GetMinimumLatencies(Latencies map[string]float64) ([]string, []float64) {

}
