package ram_inf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

//Ramtipo es el struct de tipo
type Ramtipo struct { //estructura a enviar en la ram
	Total   string `json:"Total"`
	Usada   string `json:"Usada"`
	Porcent string `json:"Porcent"`
}

//ObtenerRAM es para obtener la ram
func ObtenerRAM() []byte {
	filedata, err := ioutil.ReadFile("/proc/meminfo")
	fmt.Println("entro a ram")
	if err != nil {
		fmt.Println("Error al abrir el archivo")
		res2D := &Ramtipo{
			Total:   "0.0",
			Usada:   "0.0",
			Porcent: "0.0"}
		res2B, err := json.Marshal(res2D)
		if err != nil {
			fmt.Println("Error al convertir")
		}
		return res2B
	}
	cadena := string(filedata)
	vect := strings.Split(cadena, "\n")
	redu := strings.Replace(vect[0], "MemTotal:", "", -1)
	var act = strings.Replace(redu, " ", "", -1)
	act = strings.Replace(act, "kB", "", -1)
	libre1, err := strconv.Atoi(act)
	if err != nil {
		fmt.Println("Error al convertir")
	}
	var cien float64 = 1000.0
	var libre float64 = float64(libre1) / cien
	nuevo1 := strings.Replace(vect[1], "MemFree:", "", -1)
	act = strings.Replace(nuevo1, " ", "", -1)
	act = strings.Replace(act, "kB", "", -1)
	libre2, err := strconv.Atoi(act)
	if err != nil {
		fmt.Println("Error al convertir")
	}
	usada := float64(libre1-libre2) / cien
	porcentaje := ((usada)*100)/libre + 0.0
	res2D := &Ramtipo{
		Total:   fmt.Sprintf("%f", libre),
		Usada:   fmt.Sprintf("%f", usada),
		Porcent: fmt.Sprintf("%f", porcentaje)}
	res2B, err := json.Marshal(res2D)
	if err != nil {
		fmt.Println("Error al convertir")
	}
	return res2B
}
