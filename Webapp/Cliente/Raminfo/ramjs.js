function llamada(){
  fetch("http://localhost:8080/ram")
  .then(res => res.json())
  .then(data => {
    //data ya es un objeto
    console.log(data);
    actualizar(data);
  })
  .catch(function(error){
    console.log("error!!");
  })
}
async function prueba(){
   await sleep(2000);
   llamada();
}
function actualizar(data){
  document.getElementById("total").innerHTML ="Total: "+data.Total+" MB";
  document.getElementById("usada").innerHTML ="Ocupada: "+data.Usada+" MB";
  document.getElementById("libre").innerHTML ="Porcentaje: "+data.Porcent+"%";
  actualiza_data(data);
}
function actualiza_data(data){
var i =0;
var cambio = 0;
for(i =0;i<10;i++){
	var act = data_chart[i];
	if(act==0.0){
		data_chart[i] = parseFloat(data.Porcent);
		cambio = 1;
		break;
	}
}
if(cambio==0){//no cambio nada
  for(i =0;i<9;i++){
	var act = data_chart[i+1];
	data_chart[i] =act ;	
 }
 data_chart[9]=parseFloat(data.Porcent);
}
var newData =[
    { datox: '01', value: data_chart[0] },
    { datox: '02', value: data_chart[1] },
    { datox: '03', value: data_chart[2] },
    { datox: '04', value: data_chart[3] },
    { datox: '05', value: data_chart[4] },    
    { datox: '06', value: data_chart[5] },
    { datox: '07', value: data_chart[6] },
    { datox: '08', value: data_chart[7] },
    { datox: '09', value: data_chart[8] },
    { datox: '10', value: data_chart[9] }
  ];
datos.setData(newData);
}

function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}
