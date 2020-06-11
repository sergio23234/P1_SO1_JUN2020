function httpGet(theUrl)
{
var xhr = new XMLHttpRequest();
xhr.withCredentials = true;
xhr.addEventListener("readystatechange", function() {
  if(this.readyState === 4) {
    console.log(this.responseText);
  }
});
xhr.open("GET", "http://localhost:8080/RAM");
xhr.send();
}

function httpGetAsync(theUrl, callback) 
{ 
 var xmlHttp = new XMLHttpRequest(); 
 xmlHttp.onreadystatechange = function() { 
  if (xmlHttp.readyState == 4 && xmlHttp.status == 200) 
   callback(xmlHttp.responseText); 
 } 
 xmlHttp.open("GET", theUrl, true); // true for asynchronous 
 xmlHttp.send(null); 
} 

function prueba(){
 var actual = httpGet("http://localhost:8080/RAM");
 console.log(actual);
}

function processData(data) {
   console.log(data);
}




