package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type betaInfo struct { //para la tabla de procesos se mostraria todo menos padre
	nombre string
	pid    string
	ram    string
	user   string
	estado string
	padre  string
}
type alphaInfo struct { //para la tabla de procesos con sus hijos
	nombre string
	pid    string
	/*ram    string
	user   string
	estado string*/
	hijos []alphaInfo
}
type sendinfo struct { //informacion a enviar
	tabla    []betaInfo
	ordenado []alphaInfo
	zomibes  int
	running  int
	stoped   int
	total    int
}

var zomibes int = 0
var running int = 0
var sleeping int = 0
var stoped int = 0
var total int = 0

func main() {
	zomibes = 0
	running = 0
	sleeping = 0
	stoped = 0
	total = 0
	files, err := ioutil.ReadDir("/proc/")
	if err != nil {
		log.Fatal(err)
	}
	var arreglo []string
	var memoriaram = memoriaram()

	for _, f := range files {
		var nombre = f.Name()
		match, _ := regexp.MatchString("[0-9]+", nombre)
		if match {
			total++
			arreglo = append(arreglo, nombre)
		}
	}
	var i int
	var vectorbeta []betaInfo
	for i = 0; i < len(arreglo); i++ {
		var nueva betaInfo = leernumero(arreglo[i], memoriaram)
		vectorbeta = append(vectorbeta, nueva)
	}
	var padres []alphaInfo = ordenar(vectorbeta)
	for i = 0; i < len(padres); i++ {
		var act alphaInfo = padres[i]
		imprimirfu(act, "")
	}
	var toret sendinfo
	toret.ordenado = padres
	toret.tabla = vectorbeta
	toret.zomibes = zomibes
	toret.running = running
	toret.stoped = stoped
	toret.total = total
	fmt.Println("total:" + strconv.Itoa(total) + " zombies:" + strconv.Itoa(zomibes) + " corriendo:" + strconv.Itoa(running) + " dormidos:" + strconv.Itoa(sleeping) + " detenidos:" + strconv.Itoa(stoped))
}
func ordenar(desorden []betaInfo) []alphaInfo {
	var i int
	var padres []alphaInfo
	for i = 0; i < len(desorden); i++ {
		var act betaInfo = desorden[i]
		if act.padre == "0" {
			var nueva alphaInfo
			nueva.nombre = act.nombre
			nueva.pid = act.pid
			/*nueva.estado = act.estado
			nueva.ram = act.ram
			nueva.user = act.user*/
			nueva.hijos = retornarHijos(desorden, act.pid)
			padres = append(padres, nueva)
		}
	}
	return padres
}
func retornarHijos(desorden []betaInfo, id string) []alphaInfo {
	var hijos []alphaInfo
	var i int
	for i = 0; i < len(desorden); i++ {
		var act betaInfo = desorden[i]
		if act.padre == id {
			var nueva alphaInfo
			nueva.nombre = act.nombre
			nueva.pid = act.pid
			/*nueva.estado = act.estado
			nueva.ram = act.ram
			nueva.user = act.user*/
			nueva.hijos = retornarHijos(desorden, act.pid)
			hijos = append(hijos, nueva)
		}
	}
	return hijos
}
func imprimirfu(act alphaInfo, temp string) {
	fmt.Println(temp + "nombre: " + act.nombre + "\t pid: " + act.pid)
	if len(act.hijos) > 0 {
		var i int
		fmt.Println("")
		for i = 0; i < len(act.hijos); i++ {
			var act alphaInfo = act.hijos[i]
			imprimirfu(act, temp+"\t")
		}
	}
}
func leernumero(numero string, Rammem int) betaInfo {
	filedata, err := ioutil.ReadFile("/proc/" + numero + "/status")
	if err != nil {
		fmt.Println("Error al abrir el archivo")
		var toret betaInfo
		toret.nombre = ""
		toret.estado = ""
		toret.pid = ""
		toret.ram = ""
		toret.padre = ""
		return toret
	} else {
		cadena := string(filedata)
		vect := strings.Split(cadena, "\n")
		var estado string = estadovec(vect[2])
		var nombre string = retName(vect[0])
		var pid string = retPID(vect[5])
		var padre string = retPPID(vect[6])
		var por string = retMemo(vect[17], Rammem)
		var userx string = retuser(vect[8])
		var toret betaInfo
		toret.nombre = nombre
		toret.estado = estado
		toret.pid = pid
		toret.ram = por
		toret.padre = padre
		toret.user = userx
		//fmt.Println("padre:" + padre + "---->" + nombre + " --->" + estado + " --->" + pid + " ---> " + por)
		return toret
	}
}
func estadovec(estado string) string {
	redu := strings.Replace(estado, "State:", "", -1)
	redu = strings.Replace(redu, "\t", "", -1)
	var cad []string = strings.Split(redu, "(")
	var estadon string = strings.Replace(cad[1], ")", "", -1)
	verificarEstado(cad[0])
	return estadon
}
func verificarEstado(estado string) {
	if strings.HasPrefix(estado, "Z") {
		zomibes++
	} else if strings.HasPrefix(estado, "S") {
		sleeping++
	} else if strings.HasPrefix(estado, "R") {
		running++
	} else if strings.HasPrefix(estado, "T") {
		stoped++
	}
}
func retPID(estado string) string {
	redu := strings.Replace(estado, "Pid:", "", -1)
	redu = strings.Replace(redu, "\t", "", -1)
	return redu
}
func retPPID(estado string) string {
	redu := strings.Replace(estado, "PPid:", "", -1)
	redu = strings.Replace(redu, "\t", "", -1)
	return redu
}
func retName(estado string) string {
	redu := strings.Replace(estado, "Name:", "", -1)
	redu = strings.Replace(redu, "\t", "", -1)
	return redu
}
func retMemo(estado string, RAM int) string {
	if strings.HasPrefix(estado, "VmSize") {
		redu := strings.Replace(estado, "VmSize:", "", -1)
		redu = strings.Replace(redu, "\t", "", -1)
		var act = strings.Replace(redu, " ", "", -1)
		act = strings.Replace(act, "kB", "", -1)
		libre1, err := strconv.Atoi(act)
		if err != nil {
			fmt.Println("Error al convertir")
			return "--"
		}
		porcentaje := ((libre1)*100)/RAM + 0.0
		if porcentaje > 0 {
			return strconv.Itoa(porcentaje) + "%"
		} else {
			return "<1%"
		}
	}
	return "--"
}
func retuser(estado string) string {
	if strings.HasPrefix(estado, "Uid") {
		redu := strings.Replace(estado, "Uid:", "", -1)
		var redus []string = strings.Split(redu, "\t")
		return redus[1]
	}
	return "--"
}
func memoriaram() int {
	filedata, err := ioutil.ReadFile("/proc/meminfo")
	if err != nil {
		fmt.Println("Error al abrir el archivo")
	} else {
		cadena := string(filedata)
		vect := strings.Split(cadena, "\n")
		redu := strings.Replace(vect[0], "MemTotal:", "", -1)
		var act = strings.Replace(redu, " ", "", -1)
		act = strings.Replace(act, "kB", "", -1)
		libre1, err := strconv.Atoi(act)
		if err != nil {
			fmt.Println("Error al convertir")
		}
		return libre1
	}
	return 0
}
func terminarproceso(id string) {
	cmd := exec.Command("kill", "-9", id)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
