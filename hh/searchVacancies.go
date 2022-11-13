package hh

import (
	"encoding/json"
	"fmt"
	"headHunterBot/myLog"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

func getVacanciesByText(text string, pageIndex int) Response {
	query := fmt.Sprintf("https://api.hh.ru/vacancies?text=%s&area=40&page=%d&per_page=7", url.QueryEscape(text), pageIndex)

	req, err := http.NewRequest(http.MethodGet, query, nil)

	if err != nil {
		log.Fatalf("client: error assigning http request: %s\n", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Telegram/1.0 (rahman.abdirashov@gmail.com)")

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("client: error making http request: %s\n", err)
	}
	resBody, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("client: error reading body request: %s\n", err)
	}

	var response Response
	json.Unmarshal(resBody, &response)

	myLog.LogHHResponse(text, string(resBody))
	return response
}

func GetVacancyInformationById(id string) string {
	query := fmt.Sprintf("https://api.hh.ru/vacancies/%s", id)

	req, err := http.NewRequest(http.MethodGet, query, nil)

	if err != nil {
		log.Fatalf("client: error assigning http request: %s\n", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Telegram/1.0 (rahman.abdirashov@gmail.com)")

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("client: error making http request: %s\n", err)
	}
	resBody, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("client: error reading body request: %s\n", err)
	}

	var vacancy Vacancy
	json.Unmarshal(resBody, &vacancy)
	myLog.LogHHResponse("", string(resBody))
	fmt.Printf("%+v\n", vacancy)

	returnMessage := "" //vacancy.Employment.Name + "\n\n"
	returnMessage += fmt.Sprintf("\n\n<b>Вакансия:</b> %s\n", vacancy.Name)

	if vacancy.Salary.From != 0 && vacancy.Salary.To != 0 {
		returnMessage += fmt.Sprintf("<b>Зарплата</b>: %d - %d %s\n", vacancy.Salary.From, vacancy.Salary.To, vacancy.Salary.Currency)
	} else if vacancy.Salary.From != 0 {
		returnMessage += fmt.Sprintf("<b>Зарплата</b>: %d %s\n", vacancy.Salary.From, vacancy.Salary.Currency)
	} else if vacancy.Salary.To != 0 {
		returnMessage += fmt.Sprintf("<b>Зарплата</b>: %d %s\n", vacancy.Salary.To, vacancy.Salary.Currency)
	}
	if vacancy.Salary.Gross {
		returnMessage += "<b>Доход</b>: Валовой доход\n"
	}
	if len(vacancy.Address.Raw) > 0 {
		returnMessage += fmt.Sprintf("<b>Адрес</b>: %s\n", vacancy.Address.Raw)
	}
	returnMessage += fmt.Sprintf("<b>Работадатель</b>: %s\n", vacancy.Employer.Name)
	if len(vacancy.Schedule.Name) > 0 {
		returnMessage += fmt.Sprintf("<b>График</b>: %s", vacancy.Schedule.Name)
	}
	if len(vacancy.Employment.Name) > 0 {
		if len(vacancy.Schedule.Name) > 0 {
			returnMessage += ", " + vacancy.Employment.Name
		} else {
			returnMessage += fmt.Sprintf("<b>График</b>: %s\n", vacancy.Employment.Name)
		}
	} else {
		returnMessage += "\n"
	}
	returnMessage += "\n\n"
	if len(vacancy.Experience.Name) > 0 {
		returnMessage += fmt.Sprintf("<b>Опыт</b>: %s\n", vacancy.Experience.Name)
	}
	if len(vacancy.KeySkills) > 0 {
		returnMessage += "<b>Ключевые навыки</b>: "
		for i := 0; i < len(vacancy.KeySkills); i++ {
			returnMessage += vacancy.KeySkills[i].Name
			if i+1 != len(vacancy.KeySkills) {
				returnMessage += ", "
			}
		}
		returnMessage += "\n"
	}
	if len(vacancy.Snippet.Requirement) > 0 {
		returnMessage += fmt.Sprintf("<b>Требования</b>: %s\n", vacancy.Snippet.Requirement)
	}
	if len(vacancy.Snippet.Responsibility) > 0 {
		returnMessage += fmt.Sprintf("<b>Ответственности</b>: %s\n", vacancy.Snippet.Responsibility)
	}
	if len(vacancy.AlternateURL) > 0 {
		returnMessage += fmt.Sprintf("<b>Ссылка</b>: %s\n", vacancy.AlternateURL)
	}

	return returnMessage
}

func GetVacanciesTextByRange(text string, index int) (string, int) {
	var vacancies Response = getVacanciesByText(text, index)

	returnMessage := fmt.Sprintf("Найдено %d вакансий", vacancies.Found)

	for i := 0; i < len(vacancies.Items); i++ {
		vacancy := vacancies.Items[i]
		var vacancyText string

		vacancyText += fmt.Sprintf("\n\n<b>Вакансия:</b> %s\n", vacancy.Name)

		if vacancy.Salary.From != 0 && vacancy.Salary.To != 0 {
			vacancyText += fmt.Sprintf("<b>Зарплата</b>: %d - %d %s\n", vacancy.Salary.From, vacancy.Salary.To, vacancy.Salary.Currency)
		} else if vacancy.Salary.From != 0 {
			vacancyText += fmt.Sprintf("<b>Зарплата</b>: %d %s\n", vacancy.Salary.From, vacancy.Salary.Currency)
		} else if vacancy.Salary.To != 0 {
			vacancyText += fmt.Sprintf("<b>Зарплата</b>: %d %s\n", vacancy.Salary.To, vacancy.Salary.Currency)
		}
		if len(vacancy.Address.Raw) > 0 {
			vacancyText += fmt.Sprintf("<b>Адрес</b>: %s\n", vacancy.Address.Raw)
		}
		vacancyText += fmt.Sprintf("<b>Работадатель</b>: %s\n", vacancy.Employer.Name)
		if len(vacancy.Schedule.Name) > 0 {
			vacancyText += fmt.Sprintf("<b>График</b>: %s\n", vacancy.Schedule.Name)
		}
		vacancyText += fmt.Sprintf("<b>Детальнее</b>: /%s\n", vacancy.Id)
		returnMessage += vacancyText
	}
	return returnMessage, vacancies.Found
}
