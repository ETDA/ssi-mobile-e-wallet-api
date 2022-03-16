package helpers

import (
	"github.com/stretchr/testify/assert"
	"gitlab.finema.co/finema/etda/mobile-app-api/emsgs"
	"testing"
)

func TestBirthdateTransform(t *testing.T) {
	date, ierr := BirthdateTransform("22 ต.ค. 2539")
	assert.Equal(t, ierr, nil)
	assert.Equal(t, "25391022", date)

	date, ierr = BirthdateTransform("22ต.ค.2539")
	assert.Equal(t, ierr, emsgs.InvalidDateFormat)

	date, ierr = BirthdateTransform("22 ป.ป. 2539")
	assert.Equal(t, ierr, emsgs.InvalidDateFormat)

	date, ierr = BirthdateTransform("AA ต.ค. 2539")
	assert.Equal(t, ierr, emsgs.InvalidDateFormat)
}
