function prueba(){
  fetch("http://localhost:8080/getprin")
  .then(res => res.json())
  .then(data => {
   Actualizar_Datos(data);
  })
  .catch(function(error){
    console.log("error!!");
  })
}
function prueba_id(id_number){  
fetch("http://localhost:8080/idkill",{
	method:'post',
	headers:{
		"Content-Type":"application/json"
	},
	body: JSON.stringify({Id:id_number})})
  .then(res => res.json())
  .then(data => {
	if(data.respuesta=="1")
	{	
		alert("proceso "+id_number+" terminado");
 	}else{alert("proceso "+id_number+" no terminado");}   
  location.reload();
  })
  .catch(function(error){
    console.log("error!!");
  })
}
function Actualizar_Datos(data){
document.getElementById("total").innerHTML="Total:&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"+data.Total;
document.getElementById("zombies").innerHTML="Zombies:&nbsp;&nbsp;&nbsp;&nbsp;"+data.Zombies;
document.getElementById("corriendo").innerHTML="corriendo:&nbsp;&nbsp;"+data.Running;
document.getElementById("detenido").innerHTML="detenidos:&nbsp;&nbsp;"+data.Stoped;
document.getElementById("durmiendo").innerHTML="Durmiendo:&nbsp;"+data.Sleeped;
Graficar_Tabla(data.Tabla);
Actualizar_Arbol(data.Ordenado);
}

function Actualizar_Arbol(data){
  var tablaE = document.getElementById("Uroot");
  tablaE.parentNode.removeChild(tablaE);
  var ul=document.createElement('ul');
  ul.id = "Uroot";
  for(var i in data){
    Hijos_arbol(data[i],ul);
  }
  var Padre = document.getElementById("Arbol")
  Padre.appendChild(ul);
}
function Hijos_arbol(data,rot)
{
  var uno = document.createElement("li");
  var textnode = document.createTextNode("Nombre: "+data.Nombre+"\t\tPid: "+data.Pid);
  uno.appendChild(textnode);
  rot.appendChild(uno);
  var ul=document.createElement('ul');
  for(var i in data.Hijos){
     Hijos_arbol(data.Hijos[i],ul);
  }
  rot.appendChild(ul); 
}

function Graficar_Tabla(data){
  var tablaE = document.getElementById("Stable");
  tablaE.parentNode.removeChild(tablaE);
  var tabla = document.createElement("table");
  tabla.id = "Stable";
 var header = tabla.createTHead();
 var row = header.insertRow(0);
 var th = document.createElement('th');
 th.innerHTML = "Nombre";
 row.appendChild(th);
 var th1 = document.createElement('th');
 th1.innerHTML = "Pid";
 row.appendChild(th1);
var th2 = document.createElement('th');
 th2.innerHTML = "Estado";
 row.appendChild(th2);
 var th3 = document.createElement('th');
 th3.innerHTML = "Usuario";
 row.appendChild(th3);
 var th4 = document.createElement('th');
 th4.innerHTML = "Uso Ram";
 row.appendChild(th4);
  var tblBody = document.createElement('tbody');  
   for (var i in data) {
        var Simbolo = data[i];
        var hilera = document.createElement("tr");
        var hilera_1 = document.createElement("td");
        var subtext_t1 = document.createTextNode(Simbolo.Nombre);
        hilera_1.appendChild(subtext_t1);
        hilera.appendChild(hilera_1);
        var hilera_2 = document.createElement("td");
        var subtext_t2 = document.createTextNode(Simbolo.Pid);
        hilera_2.appendChild(subtext_t2);
        hilera.appendChild(hilera_2);
        var hilera_3 = document.createElement("td");
        var subtext_t3 = document.createTextNode(Simbolo.Estado);
        hilera_3.appendChild(subtext_t3);
        hilera.appendChild(hilera_3);
        var hilera_4 = document.createElement("td");
        var subtext_t4 = document.createTextNode(Simbolo.User);
        hilera_4.appendChild(subtext_t4);
        hilera.appendChild(hilera_4);
        var hilera_5 = document.createElement("td");
        var subtext_t5 = document.createTextNode(Simbolo.Rram);
        hilera_5.appendChild(subtext_t5);
        hilera.appendChild(hilera_5);
        tblBody.appendChild(hilera);
    }	
tabla.appendChild(tblBody);
var Padre = document.getElementById("Tabla")
Padre.appendChild(tabla);
}


// ========== por si nos sirve despues ========== 
// function httpGet(theUrl)
// {
// var xhr = new XMLHttpRequest();
// xhr.withCredentials = true;
// xhr.addEventListener("readystatechange", function() {
//   if(this.readyState === 4) {
//     console.log(this.responseText);
//   }
// });
// xhr.open("GET", "http://localhost:8080/RAM");
// xhr.send();
// }

// function httpGetAsync(theUrl, callback) 
// { 
//  var xmlHttp = new XMLHttpRequest(); 
//  xmlHttp.onreadystatechange = function() { 
//   if (xmlHttp.readyState == 4 && xmlHttp.status == 200) 
//    callback(xmlHttp.responseText); 
//  } 
//  xmlHttp.open("GET", theUrl, true); // true for asynchronous 
//  xmlHttp.send(null); 
// } 
// function processData(data) {
//   console.log(data);
// }









