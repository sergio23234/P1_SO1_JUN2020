function llamada_cpu() {
    fetch("http://localhost:8080/cpu")
        .then(res => res.json())
        .then(data => {
            //data ya es un objeto
           // console.log(data)
            console.log(data.cpu_info);
            actualizar_cpu(data.cpu_info);
        })
        .catch(function (error) {
            console.log("error!!");
        })
}


function actualizar_cpu(data) {
    document.getElementById("utilizado").innerHTML = "Total: " + data + " %";
    actualiza_data_cpu(data);
}

function actualiza_data_cpu(data) {
    var i = 0;
    var cambio = 0;
    
    for (i = 0; i < 10; i++) {
        var act = data_chart_cpu[i];
        if (act == 0.0) {
            data_chart_cpu[i] = parseFloat(data);
            cambio = 1;
            break;
        }
    }
    if (cambio == 0) {//no cambio nada
        for (i = 0; i < 9; i++) {
            var act = data_chart_cpu[i + 1];
            data_chart_cpu[i] = act;
        }
        data_chart_cpu[9] = parseFloat(data);
    }
    var newData = [
        { datox: '01', value: data_chart_cpu[0] },
        { datox: '02', value: data_chart_cpu[1] },
        { datox: '03', value: data_chart_cpu[2] },
        { datox: '04', value: data_chart_cpu[3] },
        { datox: '05', value: data_chart_cpu[4] },
        { datox: '06', value: data_chart_cpu[5] },
        { datox: '07', value: data_chart_cpu[6] },
        { datox: '08', value: data_chart_cpu[7] },
        { datox: '09', value: data_chart_cpu[8] },
        { datox: '10', value: data_chart_cpu[9] }
    ];
    datos_cpu.setData(newData);
}