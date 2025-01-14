package types

import (
	"fmt"
	"strings"
	"time"
)

type Date time.Time

const layout = "2006-01-02"

func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	t, err := time.Parse(layout, s)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}
func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, d.Format(layout))), nil
}
func (d Date) Format(s string) string {
	t := time.Time(d)
	return t.Format(s)
}
