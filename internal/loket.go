package internal

import (
	"encoding/json"
	"log"
	"reflect"
)

type LoketData struct {
	A int
	B int
	C int
	D int
	E int
	F int
}

func (ld *LoketData) Tambah(loket string) {
	log.Println("tambah loket:", loket)
	rValue := reflect.ValueOf(ld)
	if rValue.Kind() == reflect.Ptr {
		rValue = rValue.Elem()
	}

	field := rValue.FieldByName(loket)
	if field.IsValid() && field.CanSet() {
		field.SetInt(field.Int() + 1)
	}
}

func (ld *LoketData) Marshal() []byte {
	jsonString, err := json.Marshal(ld)
	if err != nil {
		log.Println("Gagal mengkonversi nilai: ", err)
	}
	return jsonString
}
