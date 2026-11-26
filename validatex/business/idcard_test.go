package business_test

import (
	"testing"
	"time"

	"github.com/gtkit/stringx/validatex/business"
)

func TestParseChinaIDCard(t *testing.T) {
	now := time.Date(2026, 3, 28, 0, 0, 0, 0, time.UTC)
	info, err := business.ParseChinaIDCard("11010519491231002x", now)
	if err != nil {
		t.Fatalf("ParseChinaIDCard returned error: %v", err)
	}
	want := business.ChinaIDCardInfo{
		Number:       "11010519491231002X",
		ProvinceCode: "11",
		BirthDate:    time.Date(1949, 12, 31, 0, 0, 0, 0, time.UTC),
		Age:          76,
		Gender:       business.GenderFemale,
	}
	if info != want {
		t.Fatalf("ParseChinaIDCard = %+v, want %+v", info, want)
	}

	if info, err := business.ParseChinaIDCard("110105194912310021", now); err == nil || info != (business.ChinaIDCardInfo{}) {
		t.Fatalf("ParseChinaIDCard with invalid checksum = %+v, %v, want zero value and error", info, err)
	}
}
