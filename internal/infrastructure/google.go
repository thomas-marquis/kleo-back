package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/thomas-marquis/kleo-back/internal/domain"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"gopkg.in/yaml.v3"
)

type accountConfig struct {
	Name              string `yaml:"name"`
	IsActive          bool   `yaml:"isActive"`
	DateParseRegexp   string `yaml:"dateParseRegexp"`
	SheetRange        string `yaml:"sheetRange"`
	LabelColumnIndex  int    `yaml:"labelColumnIndex"`
	DateColumnIndex   int    `yaml:"dateColumnIndex"`
	AmountColumnIndex int    `yaml:"amountColumnIndex"`
	TagsColumnsIndex  []int  `yaml:"tagsColumnsIndex"`
	Users             []struct {
		Name                  string `yaml:"name"`
		AllocationColumnIndex int    `yaml:"allocationColumnIndex"`
	} `yaml:"users"`
}

type legacyInitConfig struct {
	SpreadsheetID string          `yaml:"spreadSheetId"`
	Accounts      []accountConfig `yaml:"accounts"`
	Users         []struct {
		Name  string `yaml:"name"`
		Email string `yaml:"email"`
	} `yaml:"users"`
}

type GoogleLegacyRepository struct {
	sheetsSvc      *sheets.Service
	accountsStore  map[domain.BankAccountId]*domain.BankAccount
	config         *legacyInitConfig
	accountConfigs map[domain.BankAccountId]accountConfig
}

var _ domain.LegacyRepository = &GoogleLegacyRepository{}

func NewGoogleLegacyRepository(sheetsSvc *sheets.Service) *GoogleLegacyRepository {
	logger := *log.New(os.Stdout, "GoogleLegacyRepository\t", log.LstdFlags)
	ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		logger.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		logger.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config, &logger)

	svc, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	cfg, err := loadConfig()
	if err != nil {
		logger.Fatalf("Unable to load config: %v", err)
	}

	repo := &GoogleLegacyRepository{
		sheetsSvc:      svc,
		accountsStore:  make(map[domain.BankAccountId]*domain.BankAccount, len(cfg.Accounts)),
		config:         cfg,
		accountConfigs: make(map[domain.BankAccountId]accountConfig, len(cfg.Accounts)),
	}

	err = repo.initBankAccounts()
	if err != nil {
		logger.Fatalf("Unable to init bank accounts: %v", err)
	}

	return repo
}

func (r *GoogleLegacyRepository) GetBankAccounts(ctx context.Context) ([]*domain.BankAccount, error) {
	var accounts []*domain.BankAccount = make([]*domain.BankAccount, len(r.config.Accounts))
	for _, acc := range r.accountsStore {
		accounts = append(accounts, acc)
	}

	return accounts, nil
}

func (r *GoogleLegacyRepository) GetRawTransactionsByAccountId(ctx context.Context, accountID domain.BankAccountId) ([]*domain.RawTransaction, error) {
	accCfg, ok := r.accountConfigs[accountID]
	if !ok {
		return nil, fmt.Errorf("account not found")
	}

	values, err := r.getValuesFromSheet(r.config.SpreadsheetID, accCfg.SheetRange)
	if err != nil {
		return nil, fmt.Errorf("error while retrieving values from sheet: %s", err.Error())
	}
	var rawTr []*domain.RawTransaction

	fmt.Println("#First lines:")
	for i, row := range *values {
		// 	trTags := make([]entities.Tag, 0)
		//
		if i <= 3 {
			fmt.Println(row)
		}
		//
		// 	// bank date
		// 	rawDate := row[param.DateColumnIndex].(string)
		// 	if strings.TrimSpace(rawDate) == "" {
		// 		continue
		// 	}
		//
		// 	dateFormats := [3]string{"02/01/2006", "02/1/2006", "02/01/06"}
		// 	var t time.Time
		// 	var err error
		// 	for _, format := range dateFormats {
		// 		t, err = time.Parse(format, rawDate)
		// 		if err == nil {
		// 			break
		// 		}
		// 	}
		// 	if err != nil {
		// 		return nil, fmt.Errorf("error while parsing date: %s", err.Error())
		// 	}
		//
		// 	a, err := getFloatValueFromRow(row, param.AmountColumnIndex)
		// 	if err != nil {
		// 		return nil, err
		// 	}
		//
		// 	// allocation
		// 	var allocs []entities.Allocation
		// 	if param.DefaultAllocation != (entities.Allocation{}) {
		// 		allocs = []entities.Allocation{param.DefaultAllocation}
		// 	} else {
		// 		allocs = make([]entities.Allocation, 0, len(param.AllocationColIdx))
		// 		for user, colIdx := range param.AllocationColIdx {
		// 			rate, err := getFloatValueFromRow(row, colIdx)
		// 			if err != nil {
		// 				return nil, err
		// 			}
		//
		// 			allocs = append(allocs, entities.Allocation{
		// 				User:      user,
		// 				Rate:      float32(rate),
		// 				Household: OurHousehold,
		// 			})
		// 		}
		// 	}
		//
		// 	// tags
		// 	for _, colIdx := range param.TagsColsIdx {
		// 		var tagLabel string
		// 		if len(row) <= colIdx {
		// 			tagLabel = ""
		// 		} else {
		// 			tagLabel = row[colIdx].(string)
		// 			tagLabel = strings.TrimSpace(tagLabel)
		// 		}
		// 		if tagLabel == "" {
		// 			continue
		// 		}
		//
		// 		var (
		// 			tag *entities.Tag
		// 			ok  bool
		// 		)
		// 		_, ok = s.tagMap[tagLabel]
		// 		if !ok {
		// 			tag = &entities.Tag{
		// 				Label: tagLabel,
		// 			}
		// 			logger.Printf("\tSaving new tag '%s'", tagLabel)
		// 			createdTag, err := s.qualifyingService.CreateTag(tag.Label, tag.Description, *OurHousehold)
		// 			if err != nil {
		// 				return nil, err
		// 			}
		// 			s.tagMap[tagLabel] = &createdTag
		// 		}
		// 		trTags = append(trTags, *s.tagMap[tagLabel])
		// 	}
		//
		// 	// build raw transaction
		// 	newRawTr := entities.RawTransaction{
		// 		Account: *param.Account,
		// 		Date:    t,
		// 		Amount:  a,
		// 		Label:   row[param.LabelColumnIndex].(string),
		// 	}
		// 	rawTr = append(rawTr, &newRawTr)
		//
		// 	// category
		// 	var categoryLabel string
		// 	if len(row) <= param.CategoryColIdx {
		// 		categoryLabel = ""
		// 	} else {
		// 		categoryLabel = row[param.CategoryColIdx].(string)
		// 		categoryLabel = strings.TrimSpace(categoryLabel)
		// 	}
		// 	if categoryLabel != "" {
		// 		cat, ok := s.categMap[categoryLabel]
		// 		if !ok {
		// 			return nil, fmt.Errorf("unknown category label: %s", categoryLabel)
		// 		}
		// 		s.initRepository.AddCategory(newRawTr, *cat)
		// 	}
		//
		// 	// handle taggging
		// 	s.initRepository.AddTags(newRawTr, trTags)
		//
		// 	// handle allocations
		// 	s.initRepository.AddAllocations(newRawTr, allocs)
		// }
		//
		// logger.Println("#First parsed raw transactions:")
		// for i, tr := range rawTr {
		// 	if i <= 3 {
		// 		logger.Println(tr)
		// 	}
	}

	return rawTr, nil
}

func (r *GoogleLegacyRepository) GetCategoryByOldLabel(ctx context.Context, oldLabel string) (*domain.Category, error) {
	return nil, nil
}

func (r *GoogleLegacyRepository) GetDateParseRegexpByAccountId(ctx context.Context, accountID domain.BankAccountId) (regexp.Regexp, error) {

	return nil, nil
}

func (r *GoogleLegacyRepository) GetCategoryFromMetadata(ctx context.Context, metadata map[string]interface{}) (*domain.Category, error) {
	return nil, nil
}

func (r *GoogleLegacyRepository) getUsers() ([]*domain.User, error) {
	var users []*domain.User = make([]*domain.User, len(r.config.Users))
	for _, u := range r.config.Users {
		users = append(users, domain.NewUser(u.Name, u.Email))
	}
	return users, nil
}

func (r *GoogleLegacyRepository) getUserByName(name string) (*domain.User, error) {
	for _, u := range r.config.Users {
		if u.Name == name {
			return domain.NewUser(u.Name, u.Email), nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

func (r *GoogleLegacyRepository) initBankAccounts() error {
	for _, accCfg := range r.config.Accounts {
		a := domain.NewBankAccount(accCfg.Name)

		for _, userCfg := range accCfg.Users {
			u, err := r.getUserByName(userCfg.Name)
			if err != nil {
				return fmt.Errorf("user not found: %w", err)
			}
			a.AssociateUser(u)
		}

		r.accountsStore[a.ID] = a
		r.accountConfigs[a.ID] = accCfg
	}

	return nil
}

func (g *GoogleLegacyRepository) getValuesFromSheet(spreadsheetId string, readRange string) (*[][]interface{}, error) {
	fmt.Printf("Fetch data from spreadsheet %s with range %s\n", spreadsheetId, readRange)
	resp, err := g.sheetsSvc.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	}

	fmt.Printf("%d rows fetched from spreadsheet %s with range %s.\n", len(resp.Values), spreadsheetId, readRange)

	return &resp.Values, nil
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config, logger *log.Logger) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config, logger)
		saveToken(tokFile, tok, logger)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config, logger *log.Logger) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	logger.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		logger.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		logger.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token, logger *log.Logger) {
	logger.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		logger.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func loadConfig() (*legacyInitConfig, error) {
	f, err := os.Open("legacy_init.yaml")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	cfg := &legacyInitConfig{}
	err = yaml.NewDecoder(f).Decode(cfg)
	return cfg, err
}
