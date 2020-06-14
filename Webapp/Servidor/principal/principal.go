package principal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type betaInfo struct { //para la tabla de procesos se mostraria todo menos padre
	Nombre string `json:"Nombre"`
	Pid    string `json:"Pid"`
	Rram   string `json:"Rram"`
	User   string `json:"User"`
	Estado string `json:"Estado"`
	Padre  string `json:"Padre"`
}
type alphaInfo struct { //para la tabla de procesos con sus hijos
	Nombre string      `json:"Nombre"`
	Pid    string      `json:"Pid"`
	Hijos  []alphaInfo `json:"Hijos"`
}
type sendinfo struct { //informacion a enviar
	Tabla    []betaInfo  `json:"Tabla"`
	Ordenado []alphaInfo `json:"Ordenado"`
	Zomibes  int         `json:"Zombies"`
	Running  int         `json:"Running"`
	Stoped   int         `json:"Stoped"`
        Sleeped  int 	     `json:"Sleeped"`
	Total    int         `json:"Total"`
}

var zomibes int = 0
var running int = 0
var sleeping int = 0
var stoped int = 0
var total int = 0

//Principal funcion principal
func Principal() []byte {
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
	/*for i = 0; i < len(padres); i++ {
		var act alphaInfo = padres[i]
		imprimirfu(act, "")
	}*/
	var toret sendinfo
	toret.Ordenado = padres
	toret.Tabla = vectorbeta
	toret.Zomibes = zomibes
	toret.Running = running
	toret.Stoped = stoped
	toret.Total = total
	res2D := &sendinfo{
		Ordenado: padres,
		Tabla:    vectorbeta,
		Zomibes:  zomibes,
		Running:  running,
		Stoped:   stoped,
                Sleeped:   sleeping,
		Total:    total}
	res2B, err := json.Marshal(res2D)
	if err != nil {
		fmt.Println("Error al convertir")
	}
	return res2B
}
func ordenar(desorden []betaInfo) []alphaInfo {
	var i int
	var padres []alphaInfo
	for i = 0; i < len(desorden); i++ {
		var act betaInfo = desorden[i]
		if act.Padre == "0" {
			var nueva alphaInfo
			nueva.Nombre = act.Nombre
			nueva.Pid = act.Pid
			/*nueva.estado = act.estado
			nueva.ram = act.ram
			nueva.user = act.user*/
			nueva.Hijos = retornarHijos(desorden, act.Pid)
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
		if act.Padre == id {
			var nueva alphaInfo
			nueva.Nombre = act.Nombre
			nueva.Pid = act.Pid
			nueva.Hijos = retornarHijos(desorden, act.Pid)
			hijos = append(hijos, nueva)
		}
	}
	return hijos
}
func imprimirfu(act alphaInfo, temp string) {
	fmt.Println(temp + "nombre: " + act.Nombre + "\t pid: " + act.Pid)
	if len(act.Hijos) > 0 {
		var i int
		fmt.Println("")
		for i = 0; i < len(act.Hijos); i++ {
			var act alphaInfo = act.Hijos[i]
			imprimirfu(act, temp+"\t")
		}
	}
}
func leernumero(numero string, Rammem int) betaInfo {
	filedata, err := ioutil.ReadFile("/proc/" + numero + "/status")
	if err != nil {
		fmt.Println("Error al abrir el archivo")
		var toret betaInfo
		toret.Nombre = ""
		toret.Estado = ""
		toret.Pid = ""
		toret.Rram = ""
		toret.Padre = ""
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
		toret.Nombre = nombre
		toret.Estado = estado
		toret.Pid = pid
		toret.Rram = por
		toret.Padre = padre
		toret.User = userx
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
