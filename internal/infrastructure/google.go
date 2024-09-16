package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/thomas-marquis/kleo-back/internal/domain"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"gopkg.in/yaml.v3"
)

var (
	categoryNameMetaKey = "categoryName"
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
	CategoryMap []categoryMapConfig `yaml:"categoryMap"`
}

type categoryMapConfig struct {
	OldLabel         string `yaml:"oldLabel"`
	NewLabel         string `yaml:"newLabel"`
	Value            string `yaml:"value"`
	Description      string `yaml:"description"`
	SubCategoryValue string `yaml:"subCategory"`
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
	var accounts = make([]*domain.BankAccount, len(r.config.Accounts))
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
	var rawTransactions []*domain.RawTransaction = make([]*domain.RawTransaction, 0, len(*values))

	fmt.Println("#First lines:")
	for i, row := range *values {
		if i <= 3 {
			fmt.Println(row)
		}

		// bank date
		rawDate := row[accCfg.DateColumnIndex].(string)
		if strings.TrimSpace(rawDate) == "" {
			continue
		}

		dateFormats := [3]string{"02/01/2006", "02/1/2006", "02/01/06"}
		var (
			t   time.Time
			err error
		)
		for _, format := range dateFormats {
			t, err = time.Parse(format, rawDate)
			if err == nil {
				break
			}
		}
		if err != nil {
			return nil, fmt.Errorf("error while parsing date: %s", err.Error())
		}

		a, err := getFloatValueFromRow(row, accCfg.AmountColumnIndex)
		if err != nil {
			return nil, err
		}

		label := row[accCfg.LabelColumnIndex].(string)

		rawTransaction := domain.NewRawTransaction(label, t, a)
		rawTransactions = append(rawTransactions, rawTransaction)
	}

	return rawTransactions, nil
}

func (r *GoogleLegacyRepository) GetCategoryByOldLabel(ctx context.Context, oldLabel string) (*domain.Category, error) {
	return nil, nil
}

func (r *GoogleLegacyRepository) GetDateParseRegexpByAccountId(ctx context.Context, accountID domain.BankAccountId) (regexp.Regexp, error) {
	accCfg, ok := r.accountConfigs[accountID]
	if !ok {
		return regexp.Regexp{}, fmt.Errorf("account not found")
	}
	reg := regexp.MustCompile(accCfg.DateParseRegexp)
	return *reg, nil
}

func (r *GoogleLegacyRepository) GetCategoryFromMetadata(ctx context.Context, metadata map[string]interface{}) (*domain.Category, error) {
	// catName, ok := metadata[categoryNameMetaKey]
	// if !ok {
	// 	return nil, fmt.Errorf("category not found in metadata")
	// }
	cat := &domain.Category{}
	return cat, nil
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

// Returns a float value from a row, handling the case where the value is empty or badly formatted
func getFloatValueFromRow(row []interface{}, col int) (float64, error) {
	value := row[col].(string)
	value = strings.ReplaceAll(value, ",", ".")
	value = strings.TrimSpace(value)
	value = strings.ReplaceAll(value, " ", "")
	value = strings.ReplaceAll(value, "\u202f", "")
	if value == "" {
		value = "0.0"
	}
	f, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return 0, fmt.Errorf("error while parsing amount %s: %s", value, err.Error())
	}
	return f, nil
}
