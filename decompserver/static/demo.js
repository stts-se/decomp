"use strict";

//let baseURL = window.location.protocol + '//' + window.location.host + window.location.pathname.replace(/\/$/g,"");

let baseURL = window.location;

//function init() { 
    document.getElementById("decomp_button").addEventListener("click", runDecomp);
//}

function runDecomp() {
    
    let word = document.getElementById("input_word").value;
    let select = document.getElementById("decomp_select");
    var decomper = select.options[select.selectedIndex].value;
    let decompURL = baseURL + "/decomp/" + decomper + "/" + word;

    console.log(decompURL);
}
