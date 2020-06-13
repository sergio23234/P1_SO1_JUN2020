function prueba(){
  fetch("http://localhost:8080/cpu")
  .then(res => res.json())
  .then(data => {
    console.log(data);
  })
  .catch(function(error){
    console.log("error!!");
  })
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









