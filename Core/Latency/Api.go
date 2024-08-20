package Latency

import (
	"Taxocli/Config"
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"
)

func (this *LatencyChecker) RunCommandExec() (LatencyCheckerOutputList, error) {
	AvailableTokens, err := this.DoGetTokenRequest()
	if err != nil {
		switch AvailableTokens {
		case -1:
			log.Println(" [ ERROR ]: Detected When Running The Request To The Token API")
			break
		case -2:
			log.Println(" [ ERROR ]: Detected When Trying To Decode API Response")
			break
		case -3:
			log.Println(" [ ERROR ]: Unexpected HTTP Response Code")
			break
		default:
			log.Println(" [ ERROR ]: Unexpected")
			break
		}
		return LatencyCheckerOutputList{}, err
	}
	RequiredTokens := this.GetRuns() * Config.TAXOCLI_LATENCY_TOKENS_COST
	log.Printf(" [ INFO ]: Required Tokens For This Execution { %d }, Available Tokens: { %d } ", RequiredTokens, AvailableTokens)

	if AvailableTokens < RequiredTokens {
		return LatencyCheckerOutputList{}, errors.New(" [ ERROR ]: Insufficient Tokens")
	}
	time.Sleep(Config.TAXOCLI_API_THROTTLER_TIME * time.Second)
	LatencyResults := make(map[string]float64)
	log.Printf(" [ INFO ]: Sleeping { %d }s Between Latency Requests Tests", this.GetWaitInterval())

	for I := 1; I < this.GetRuns(); I++ {
		log.Printf(" [ INFO ]: Request Number { [%d/%d] }", I, this.GetRuns())
		ResponseLatencyCheck, err := this.DoPostLatencyCheckRequest()

		if err != nil {
			log.Printf(" [ ERROR ]: Doing Latency Check Request")
			return LatencyCheckerOutputList{}, err
		}

		for KV, VK := range ResponseLatencyCheck {
			Latency := VK.(map[string]interface{})["Latency"]
			StatusCode := VK.(map[string]interface{})["StatusCode"]

			if StatusCode.(float64) != 200 {
				Latency = 1000
			}
			LatencyResults[KV] += Latency.(float64)
		}

		if this.GetRuns() > 1 {
			time.Sleep(time.Duration(this.GetWaitInterval()) * time.Second)
		}
	}

	var (
		OutputList LatencyCheckerOutputList
		Output     LatencyCheckerOutput
	)
	BestLocation, AvgLatencies := this.GetMinimumLatencies(LatencyResults)

	for I := 0; I < this.GetOutputLocationsNumber(); I++ {
		Output.AvgLatency = AvgLatencies[I]
		Output.Location = BestLocation[I]
		OutputList.Result = append(OutputList.Result, Output)
	}
	return OutputList, nil
}

func (this *LatencyCheckerOutputList) DeepCopyInto(Output *LatencyCheckerOutputList) {
	*Output = *this

	if this.Result != nil {
		this, Output := &this.Result, &Output.Result
		*Output = make([]LatencyCheckerOutput, len(*this))

		for I := range *this {
			(*Output)[I] = (*this)[I]
		}

	}

}
func (this *LatencyCheckerOutputList) DeepCopy() *LatencyCheckerOutputList {

	if this == nil {
		return nil
	}
	Output := new(LatencyCheckerOutputList)
	this.DeepCopyInto(Output)
	return Output
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
	BodyResponse := &TokenAPIResponse{}
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
	OutputKeys := make([]string, len(Latencies))
	OutputLatency := make([]float64, len(Latencies))
	Keys := make([]string, 0, len(Latencies))

	for K := range Latencies {
		Keys = append(Keys, K)
	}
	sort.SliceStable(Keys, func(i, j int) bool {
		return Latencies[Keys[i]] < Latencies[Keys[j]]
	})

	if this.GetOutputLocationsNumber() > len(Latencies) {
		this.SetOutputLocationsNumber(len(Latencies))
	}

	for I := 0; I < this.GetOutputLocationsNumber(); I++ {
		OutputKeys[I] = Keys[I]
		OutputLatency[I] = Latencies[Keys[I]]
	}
	return OutputKeys, OutputLatency
}
