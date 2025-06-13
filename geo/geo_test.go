package geo_test

import (
	"demo/weather/geo"
	"testing"
)

func TestGetMyLocation(t *testing.T) {
	//Arrange - подготовка, expected result, данные для функции
	city := "Moscow"
	expected := geo.GeoData{
		City: "Moscow",
	}
	//Act - выполнение функции
	got, err := geo.GetMyLocation(city)
	if err != nil {
		t.Error(err)
	}
	if got.City != expected.City {
		t.Errorf("Ожидалось %v, получено %v", expected, got)
	}

	//Assert - проверка результата с expected
}

func TestGetMyLocationNoCyti(t *testing.T) {
	city := "Londonasdf"
	_, err := geo.GetMyLocation(city)
	if err != geo.ErrNoCity {
		t.Errorf("Ожидалось %v, получено %v", geo.ErrNoCity, err)
	}
}
