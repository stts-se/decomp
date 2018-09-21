"use strict";

//let baseURL = window.location.protocol + '//' + window.location.host + window.location.pathname.replace(/\/$/g,"");

let baseURL = window.location.host;

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
    let decompURL = baseURL + "/decomp/" + decomper + "/" + word;
    
    console.log(decompURL);
}
