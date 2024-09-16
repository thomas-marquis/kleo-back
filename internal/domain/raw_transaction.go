package domain

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/google/uuid"
)

type RawTransactionId uuid.UUID

type RawTransaction struct {
	ID       RawTransactionId
	Label    string
	Date     time.Time
	Metadata map[string]interface{}
	Amount   float64
}

func NewRawTransaction(label string, date time.Time, amount float64) *RawTransaction {
	return &RawTransaction{
		ID:       RawTransactionId(uuid.New()),
		Label:    label,
		Date:     date,
		Amount:   amount,
		Metadata: make(map[string]interface{}),
	}
}

func (t *RawTransaction) ExtractDateFromLabel(regex *regexp.Regexp) (time.Time, error) {
	matches := regex.FindStringSubmatch(t.Label)

	if matches == nil {
		return t.Date, nil
	}

	if len(matches) < 3 {
		msg := fmt.Sprintf("Unable to parse label %s with the regexp %s", t.Label, regex.String())
		return t.Date, errors.New(msg) // TODO: add a test for this case
	}

	month := matches[2]
	year := t.Date.Format("2006")
	sdt := matches[1] + "/" + month + "/" + year
	dt, _ := time.Parse("02/01/2006", sdt)

	if dt.Month() == time.December && t.Date.Month() == time.January {
		dt = dt.AddDate(-1, 0, 0)
	}

	return dt, nil
}

func (t *RawTransaction) CleanupLabel(dateRegexp *regexp.Regexp) (string, error) {
	withoutLineBreak := regexp.MustCompile(`\n`).ReplaceAllString(t.Label, " ")
	withoutDate := dateRegexp.ReplaceAllString(withoutLineBreak, "")
	withoutSpaces := regexp.MustCompile(`\s+`).ReplaceAllString(withoutDate, " ")
	trimed := regexp.MustCompile(`^\s+|\s+$`).ReplaceAllString(withoutSpaces, "")

	return trimed, nil
}
