package main

import (
	"io/ioutil"
	"strings"
	"strconv"
	"fmt"
)
type Ram struct {
    total float32
	libre  float32
	porcent float32
}

func main() {
	file_data,err := ioutil.ReadFile("/proc/meminfo")
	if err != nil {
		fmt.Println("Error al abrir el archivo")
	}else{
		cadena := string(file_data) 
		vect := strings.Split(cadena,"\n")
		redu := strings.Replace(vect[0],"MemTotal:","",-1)
		var act = strings.Replace(redu," ","",-1)
		act = strings.Replace(act,"kB","",-1)
		 libre1 , err := strconv.Atoi(act)
		 if err != nil {
			fmt.Println("Error al abrir el archivo")
		}
		 libre := libre1/1000.0
		 fmt.Println(libre);
		 nuevo1 := strings.Replace(vect[1],"MemFree:","",-1)
		 act = strings.Replace(nuevo1," ","",-1)
		 act = strings.Replace(act,"kB","",-1)
		 libre2 , err := strconv.Atoi(act)
		 if err != nil {
			fmt.Println("Error al abrir el archivo")
		}
		 usada := (libre1-libre2)/1000
		fmt.Println(usada);
		porcentaje := ((usada)*100)/libre+0.0
		fmt.Println(porcentaje)
	}
}