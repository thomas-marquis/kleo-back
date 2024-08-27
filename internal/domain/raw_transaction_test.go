package domain_test

import (
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/thomas-marquis/kleo-back/internal/domain"
)

func Test_ExtractDateFromLabel_ShouldReturnDate1(t *testing.T) {
	// Given
	regexp := regexp.MustCompile(`(\d{2})/(\d{2})`)
	tr := domain.RawTransaction{
		Label: "CARTE 28/08 REL.D'AGEN LAPLUME",
		Date:  time.Date(2020, 8, 31, 0, 0, 0, 0, time.UTC),
	}

	// When
	date, err := tr.ExtractDateFromLabel(regexp)

	// Then
	assert.NoError(t, err)

	assert.Equal(t, date.Year(), 2020)
	assert.Equal(t, date.Month(), time.Month(8))
	assert.Equal(t, date.Day(), 28)
}

func Test_ExtractDateFromLabel_ShouldReturnDate2(t *testing.T) {
	// Given
	regexp := regexp.MustCompile(`(\d{2})/(\d{2})`)
	tr := domain.RawTransaction{
		Label: "CARTE 11/12 INTERMARCHE EYSINES",
		Date:  time.Date(2020, 12, 12, 0, 0, 0, 0, time.UTC),
	}

	// When
	date, err := tr.ExtractDateFromLabel(regexp)

	// Then
	assert.NoError(t, err)

	assert.Equal(t, date.Year(), 2020)
	assert.Equal(t, date.Month(), time.December)
	assert.Equal(t, date.Day(), 11)
}

func Test_ExtractDateFromLabel_ShouldReturnDate3(t *testing.T) {
	// Given
	regexp := regexp.MustCompile(`(\d{2})/(\d{2})`)
	tr := domain.RawTransaction{
		Label: "CARTE 12/06 DHONDT EVI ST JEAN D ILLA",
		Date:  time.Date(2023, 6, 13, 0, 0, 0, 0, time.UTC),
	}

	// When
	date, err := tr.ExtractDateFromLabel(regexp)

	// Then
	assert.NoError(t, err)

	assert.Equal(t, date.Year(), 2023)
	assert.Equal(t, date.Month(), time.Month(6))
	assert.Equal(t, date.Day(), 12)
}

func Test_ExtractDateFromLabel_ShouldApplyPreviousYearWhenMonthIsDecember1(t *testing.T) {
	// Given
	regexp := regexp.MustCompile(`(\d{2})/(\d{2})`)
	tr := domain.RawTransaction{
		Label: "CARTE 31/12 INTERMARCHE SEURRE",
		Date:  time.Date(2022, 1, 3, 0, 0, 0, 0, time.UTC),
	}

	// When
	date, err := tr.ExtractDateFromLabel(regexp)

	// Then
	assert.NoError(t, err)

	assert.Equal(t, date.Year(), 2021)
	assert.Equal(t, date.Month(), time.Month(12))
	assert.Equal(t, date.Day(), 31)
}

func Test_ExtractDateFromLabel_ShouldApplyPreviousYearWhenMonthIsDecember2(t *testing.T) {
	// Given
	regexp := regexp.MustCompile(`(\d{2})/(\d{2})`)
	tr := domain.RawTransaction{
		Label: "CARTE 31/12 Smile&P*BMCA EYSINES",
		Date:  time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
	}

	// When
	date, err := tr.ExtractDateFromLabel(regexp)

	// Then
	assert.NoError(t, err)

	assert.Equal(t, date.Year(), 2022)
	assert.Equal(t, date.Month(), time.Month(12))
	assert.Equal(t, date.Day(), 31)
}

func Test_ExtractDateFromLabel_ShouldApplyBankDateIfNoDateInLabel1(t *testing.T) {
	// Given
	regexp := regexp.MustCompile(`([0-3]\d)[\.\/]([01]\d)[\.\/](\d{2})`)
	tr := domain.RawTransaction{
		Label: "PRLV Bouygues Telecom",
		Date:  time.Date(2023, 7, 17, 0, 0, 0, 0, time.UTC),
	}

	// When
	date, err := tr.ExtractDateFromLabel(regexp)

	// Then
	assert.NoError(t, err)

	assert.Equal(t, date.Year(), 2023)
	assert.Equal(t, date.Month(), time.Month(7))
	assert.Equal(t, date.Day(), 17)
}

func Test_ExtractDateFromLabel_ShouldApplyBankDateIfNoDateInLabel2(t *testing.T) {
	// Given
	regexp := regexp.MustCompile(`([0-3]\d)[\.\/]([01]\d)[\.\/](\d{2})`)
	tr := domain.RawTransaction{
		Label: "vers MLLE DIBLING MELANIE",
		Date:  time.Date(2023, 5, 2, 0, 0, 0, 0, time.UTC),
	}

	// When
	date, err := tr.ExtractDateFromLabel(regexp)

	// Then
	assert.NoError(t, err)

	assert.Equal(t, date.Year(), 2023)
	assert.Equal(t, date.Month(), time.Month(5))
	assert.Equal(t, date.Day(), 2)
}

func Test_ExtractDateFromLabel_ShouldExtractDateWithYear(t *testing.T) {
	// Given
	regexp := regexp.MustCompile(`([0-3]\d)[\.\/]([01]\d)[\.\/](\d{2})`)
	tr := domain.RawTransaction{
		Label: "ACHAT CB BINH DUONG 13.03.15",
		Date:  time.Date(2015, 3, 15, 0, 0, 0, 0, time.UTC),
	}

	// When
	date, err := tr.ExtractDateFromLabel(regexp)

	// Then
	assert.NoError(t, err)

	assert.Equal(t, date.Year(), 2015)
	assert.Equal(t, date.Month(), time.Month(3))
	assert.Equal(t, date.Day(), 13)
}

func Test_ExtractDateFromLabel_ShouldExtractDateWithYearInPriority(t *testing.T) {
	// Given
	regexp := regexp.MustCompile(`([0-3]\d)[\.\/]([01]\d)[\.\/](\d{2})`)
	tr := domain.RawTransaction{
		Label: "ACHAT CB VEMMAEUROPE.CO 30.12.13\nEUR 128,00 CARTE NO 421",
		Date:  time.Date(2014, 1, 2, 0, 0, 0, 0, time.UTC),
	}

	// When
	date, err := tr.ExtractDateFromLabel(regexp)

	// Then
	assert.NoError(t, err)

	assert.Equal(t, date.Year(), 2013)
	assert.Equal(t, date.Month(), time.Month(12))
	assert.Equal(t, date.Day(), 30)
}

func Test_CleanupLabel_ShouldRemoveLineBreaks(t *testing.T) {
	// Given
	regexp := regexp.MustCompile(`([0-3]\d)[\.\/]([01]\d)[\.\/](\d{2})`)
	tr := domain.RawTransaction{
		Label: "ACHAT CB VEMMAEUROPE.CO\nEUR 128,00 CARTE NO 421",
	}

	// When
	label, err := tr.CleanupLabel(regexp)

	// Then
	assert.NoError(t, err)

	assert.Equal(t, label, "ACHAT CB VEMMAEUROPE.CO EUR 128,00 CARTE NO 421")
}

func Test_CleanupLabel_ShouldRemoveDate(t *testing.T) {
	// Given
	regexp := regexp.MustCompile(`([0-3]\d)[\.\/]([01]\d)[\.\/](\d{2})`)
	tr := domain.RawTransaction{
		Label: "ACHAT CB VEMMAEUROPE.CO 30.12.13EUR 128,00 CARTE NO 421",
	}

	// When
	label, err := tr.CleanupLabel(regexp)

	// Then
	assert.NoError(t, err)

	assert.Equal(t, label, "ACHAT CB VEMMAEUROPE.CO EUR 128,00 CARTE NO 421")
}

func Test_CleanupLabel_ShouldRemoveConsecutiveSpaces(t *testing.T) {
	// Given
	regexp := regexp.MustCompile(`([0-3]\d)[\.\/]([01]\d)[\.\/](\d{2})`)
	tr := domain.RawTransaction{
		Label: "ACHAT CB VEMMAEUROPE.CO 30.12.13   EUR 128,00 CARTE NO 421",
	}

	// When
	label, err := tr.CleanupLabel(regexp)

	// Then
	assert.NoError(t, err)

	assert.Equal(t, label, "ACHAT CB VEMMAEUROPE.CO EUR 128,00 CARTE NO 421")
}

func Test_CleanupLabel_ShouldRemoveStartAndEndSpaces(t *testing.T) {
	// Given
	regexp := regexp.MustCompile(`([0-3]\d)[\.\/]([01]\d)[\.\/](\d{2})`)
	tr := domain.RawTransaction{
		Label: " ACHAT CB VEMMAEUROPE.CO 30.12.13EUR 128,00 CARTE NO 421  \n ",
	}

	// When
	label, err := tr.CleanupLabel(regexp)

	// Then
	assert.NoError(t, err)

	assert.Equal(t, label, "ACHAT CB VEMMAEUROPE.CO EUR 128,00 CARTE NO 421")
}
