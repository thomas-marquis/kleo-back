package apptheme

import "time"

func FormatDate(d time.Time) string {
	return d.Format("02/01/2006")
}
