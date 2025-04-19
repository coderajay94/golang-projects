package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"encoding/json"
	"log"

	"github.com/gocolly/colly"
	"github.com/gorilla/mux"
)

var yahooURL = "https://finance.yahoo.com/quote/"

var CompaniesData = make(map[string]CompanyDetails)

var Companies = []string{
	"MEDANTA",
	"UNIONBANK",
	"EASEMYTRIP",
	"ATULAUTO",
	// "SATIN",
	// "QUESS",
	// "JIOFIN",
	// "ACC",
	// "BAJFINANCE",
	// "CAMPUS",
	// "ANGELONE",
	// "REPCOHOME",
	// "BOROLTD",
	// "SFL",
	// "BERGEPAINT",
	// "JSWHL",
	// "AAVAS",
	// "ROSSARI",
	// "TITAGARH",
	// "SAIL",
	// "KANSAINER",
	// "BANDHANBNK",
	// "NESTLEIND",
	// "HDFCBANK",
	// "JPPOWER",
	// "KOTAKBANK",
	// "TITAN",
	// "ASIANPAINT",
	// "ADANIPORTS",
	// "DABUR",
	// "TRIDENT",
	// "GLAND",
	// "NHPC",
	// "TATAMOTORS",
	// "HINDUNILVR",
	// "RELIANCE",
	// "NMDC",
	// "SBIN",
	// "HINDUNILVR",
	// "BRITANNIA",
	// "LICI",
	// "BIOCON",
	// "COLPAL",
	// "LT",
	// "SAMMAANCAP",
}

// Struct to store company data
type CompanyDetails struct {
	CompanyCode               string    `json:"companyCode"`
	Name                      string    `json:"name"`
	Price                     string    `json:"price"`
	Week52High                string    `json:"week52High"`
	Week52Low                 string    `json:"week52Low"`
	DownBy                    float64   `json:"downBy"`
	DownByPercentage          float64   `json:"downByPercentage"`
	PotentialProfitPercentage float64   `json:"potentialProfitPercentage"`
	LastUpdated               time.Time `json:"lastUpdated"`
	DailyChange               string    `json:"dailyChange"`
	DailyChangePercentage     string    `json:"dailyChangePercentage"`
	DayRange                  string    `json:"dayRange"`
	MarketCap                 string    `json:"marketCap"`
	PERatio                   string    `json:"peRatio"`
	EPS                       string    `json:"eps"`
	YTDReturn                 string    `json:"ytdReturn"`
	OneYearReturn             string    `json:"oneYearReturn"`
	ThreeYearReturn           string    `json:"threeYearReturn"`
	FiveYearReturn            string    `json:"fiveYearReturn"`
	TotalCash                 string    `json:"totalCash"`
	DebtToEquity              string    `json:"debtToEquity"`
	CompanyProfile            string    `json:"companyProfile"`
	Sector                    string    `json:"sector"`
	Industry                  string    `json:"industry"`
}

// ScrapeCompany scrapes data for a single company
func ScrapeCompany(companyCode string) CompanyDetails {
	var companyData CompanyDetails

	c := colly.NewCollector(
		colly.AllowedDomains("finance.yahoo.com"),
	)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Delay:       2 * time.Second,
		RandomDelay: 1 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0")
		log.Println("Visiting", r.URL.String())
	})

	c.OnHTML("body", func(e *colly.HTMLElement) {
		title := e.ChildText("h1.yf-xxbei9")
		price := e.ChildText("span[data-testid='qsp-price']")
		range52 := e.ChildText("fin-streamer[data-field='fiftyTwoWeekRange']")

		// New fields
		dailyChange := e.ChildText("span[data-testid='qsp-price-change']")
		dailyChangePercent := e.ChildText("span[data-testid='qsp-price-change-percent']")
		dayRange := e.ChildText("fin-streamer[data-field='regularMarketDayRange']")
		marketCap := e.ChildText("fin-streamer[data-field='marketCap']")
		peRatio := e.ChildText("fin-streamer[data-field='trailingPE']")
		//eps := e.ChildText("td[data-test='EPS_RATIO-value']")
		eps := e.ChildText("fin-streamer[data-field='trailingEPS']")
		ytdReturn := e.ChildText("td[data-test='YTD_RETURN-value']")
		oneYearReturn := e.ChildText("td[data-test='ONE_YEAR_RETURN-value']")
		threeYearReturn := e.ChildText("td[data-test='THREE_YEAR_RETURN-value']")
		fiveYearReturn := e.ChildText("td[data-test='FIVE_YEAR_RETURN-value']")
		totalCash := e.ChildText("td[data-test='TOTAL_CASH-value']")
		debtEquity := e.ChildText("td[data-test='DEBT_EQUITY-value']")
		profile := e.ChildText("p.yf-1ja4ll8")
		sector := e.ChildText("a[data-ylk*='slk:Healthcare']")
		industry := e.ChildText("a[data-ylk*='slk:Medical%20Care%20Facilities']")

		low, high := "", ""
		if range52 != "" {
			parts := strings.Split(range52, " - ")
			if len(parts) == 2 {
				low = parts[0]
				high = parts[1]
			}
		}

		downBy := 0.0
		downByPercentage := 0.0
		potentialProfitPercentage := 0.0

		if price != "" && high != "" {
			currPrice := parsePrice(price)
			week52High := parsePrice(high)

			if currPrice < week52High {
				downBy = week52High - currPrice
				downByPercentage = (downBy / week52High) * 100
				potentialProfitPercentage = (downBy / currPrice) * 100
			}
		}

		companyData = CompanyDetails{
			CompanyCode:               companyCode,
			Name:                      title,
			Price:                     price,
			Week52High:                high,
			Week52Low:                 low,
			DownBy:                    downBy,
			DownByPercentage:          downByPercentage,
			PotentialProfitPercentage: potentialProfitPercentage,
			LastUpdated:               time.Now(),
			DailyChange:               dailyChange,
			DailyChangePercentage:     dailyChangePercent,
			DayRange:                  dayRange,
			MarketCap:                 marketCap,
			PERatio:                   peRatio,
			EPS:                       eps,
			YTDReturn:                 ytdReturn,
			OneYearReturn:             oneYearReturn,
			ThreeYearReturn:           threeYearReturn,
			FiveYearReturn:            fiveYearReturn,
			TotalCash:                 totalCash,
			DebtToEquity:              debtEquity,
			CompanyProfile:            profile,
			Sector:                    sector,
			Industry:                  industry,
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Printf("❌ Failed to scrape %s: %v\n", companyCode, err)
	})

	err := c.Visit(yahooURL + companyCode + "/")
	if err != nil {
		log.Println("Visit error:", err)
	}

	return companyData
}

func parsePrice(priceStr string) float64 {
	priceStr = strings.ReplaceAll(priceStr, ",", "")
	var price float64
	_, err := fmt.Sscanf(priceStr, "%f", &price)
	if err != nil {
		log.Printf("Error parsing price: %s", err)
		return 0.0
	}
	return price
}

func startPolling(Companies []string) []CompanyDetails {
	var companyData []CompanyDetails
	for _, company := range Companies {
		data := ScrapeCompany(company + ".NS")
		data.CompanyCode = company
		companyData = append(companyData, data)
		CompaniesData[company] = data
		fmt.Printf("Company Code: %s\n", data.CompanyCode)
		fmt.Printf("Company: %s\nPrice: %s\n52 Week High: %s\n52 Week Low: %s\nDown By: %.2f\nDown By Percentage: %.2f%%\nPotential Profit Percentage: %.2f%%\n\n",
			data.Name, data.Price, data.Week52High, data.Week52Low, data.DownBy, data.DownByPercentage, data.PotentialProfitPercentage)
	}
	return companyData
}

func GetCompanyData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	companyCode := params["companyCode"]
	w.Header().Set("Content-Type", "application/json")

	if data, exists := CompaniesData[companyCode]; exists {
		json.NewEncoder(w).Encode(data)
		return
	}
	// } else {
	// 	//http.Error(w, "Company not found", http.StatusNotFound)
	// }

	fmt.Println("Scrap new data for Company Code:", companyCode)
	data := startPolling([]string{companyCode})
	if len(data) > 0 {
		json.NewEncoder(w).Encode(data[0])
		CompaniesData[companyCode] = data[0]
		// Save the updated data to JSON file
		file := "data.json"
		err := SaveMapToJSON(file, CompaniesData)
		if err != nil {
			log.Fatal("Save failed:", err)
		}
		fmt.Println("Company data updated successfully.")
		return
	}
	http.Error(w, "Company not found", http.StatusNotFound)
}

func GetAllCompanies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Initialize a slice with enough capacity based on the number of companies
	companies := make([]CompanyDetails, 0, len(CompaniesData))

	// Iterate over the Companies map and append matching data to the companies slice
	for _, company := range Companies {
		if data, exists := CompaniesData[company]; exists {
			companies = append(companies, data)
		}
	}
	// Print the total number of companies found
	fmt.Println("total companies found", len(companies))

	// Return the companies data in JSON format
	json.NewEncoder(w).Encode(companies)
}

func LoadMapFromJSON(filePath string) (map[string]CompanyDetails, error) {
	var data map[string]CompanyDetails

	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	log.Println("✅ Data loaded from", filePath)
	return data, nil
}

func SaveMapToJSON(filePath string, data map[string]CompanyDetails) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return err
	}

	log.Println("✅ Data saved to", filePath)
	return nil
}

func main() {

	//if len(Companies) == len(CompaniesData) {
	startPolling(Companies)

	file := "data.json"

	// Save
	err := SaveMapToJSON(file, CompaniesData)
	if err != nil {
		log.Fatal("Save failed:", err)
	}

	// Load
	loadedData, err := LoadMapFromJSON(file)
	if err != nil {
		fmt.Println("Load failed:", err)
	}

	for code, details := range loadedData {
		fmt.Printf("📈 %s => %+v\n", code, details)
	}

	r := mux.NewRouter()
	r.HandleFunc("/companies", GetAllCompanies).Methods("GET")
	r.HandleFunc("/companies/{companyCode}", GetCompanyData).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
