"use strict";

//let baseURL = window.location.protocol + '//' + window.location.host + window.location.pathname.replace(/\/$/g,"");

let baseURL = window.location.protocol + '//' + window.location.host;

//function init() { 
document.getElementById("decomp_button").addEventListener("click", runDecomp);
document.getElementById("input_word")
    .addEventListener("keyup", function(event) {
	event.preventDefault();
	if (event.keyCode === 13) {
            document.getElementById("decomp_button").click();
	}
    });

//}

function runDecomp() {
    
    let word = document.getElementById("input_word").value;
    let select = document.getElementById("decomp_select");
    var decomper = select.options[select.selectedIndex].value;
    let decompURL = baseURL + "/decomp/" + encodeURIComponent(decomper) + "/" + encodeURIComponent(word);
    
    console.log(decompURL);
    
    let xhr = new XMLHttpRequest();
    xhr.open('GET', decompURL);
    xhr.setRequestHeader('Content-type', 'application/x-www-form-urlencoded');
    
    xhr.onload = function(evt) {
	if ( xhr.readyState === 4) {
     	    if (xhr.status === 200) {
		let res = xhr.responseText;
		
		document.getElementById("output").innerText = res;
	    } // TODO on error
	}
    };
    
    
    
    xhr.send();
}
