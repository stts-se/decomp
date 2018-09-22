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

document.getElementById("add_prefix_button").addEventListener("click", addPrefix);
document.getElementById("remove_prefix_button").addEventListener("click", removePrefix);
document.getElementById("add_suffix_button").addEventListener("click", addSuffix);
document.getElementById("remove_suffix_button").addEventListener("click", removeSuffix);

//}

function runDecomp() {
    
    let word = document.getElementById("input_word").value.toLowerCase();
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
		let resp = xhr.responseText;
		let respObj = JSON.parse(resp);
		let res = "";
		// If no compound parts, just present the original
		// word
		if ( respObj.length === 0) {
		    res = word; 
		}
		for (var i=0; i < respObj.length; i++) {
		    //console.log("PARTS", resObj[i].parts);
		    res += respObj[i].parts.join(" ") + "\n";
		}
		
		
		document.getElementById("output").innerText = res;
	    } 
	} else {
	    document.getElementById("message_div").innerText = xhr.responseText;
	}
    };

    xhr.onerror = function(evt) {
	document.getElementById("message_div").innerText = "ERROR: Failed to call server";
    };
    
    xhr.send();
}


// TODO Woohoo, copy and paste below !!!

function addPrefix() {
    
    let word = document.getElementById("prefix_input").value.toLowerCase();

    if (word.trim() === "") {
	return;
    }
    
    let select = document.getElementById("decomp_select");
    var decomper = select.options[select.selectedIndex].value;
    let decompURL = baseURL + "/decomp/" + encodeURIComponent(decomper) + "/add_prefix/" + encodeURIComponent(word);
    
    console.log(decompURL);
    
    let xhr = new XMLHttpRequest();
    xhr.open('GET', decompURL);
    xhr.setRequestHeader('Content-type', 'application/x-www-form-urlencoded');
    
    xhr.onload = function(evt) {
	if ( xhr.readyState === 4) {
     	    if (xhr.status === 200) {
		let res = xhr.responseText;
		
		document.getElementById("message_div").innerText = res;
	    } // TODO on error
	}
    };

    xhr.onerror = function(evt) {
	document.getElementById("message_div").innerText = "ERROR: Failed to call server";
    };
    
    xhr.send();
}

function removePrefix() {
    
    let word = document.getElementById("prefix_input").value.toLowerCase();

    if (word.trim() === "") {
	return;
    }
    
    let select = document.getElementById("decomp_select");
    var decomper = select.options[select.selectedIndex].value;
    let decompURL = baseURL + "/decomp/" + encodeURIComponent(decomper) + "/remove_prefix/" + encodeURIComponent(word);
    
    console.log(decompURL);
    
    let xhr = new XMLHttpRequest();
    xhr.open('GET', decompURL);
    xhr.setRequestHeader('Content-type', 'application/x-www-form-urlencoded');
    
    xhr.onload = function(evt) {
	if ( xhr.readyState === 4) {
     	    if (xhr.status === 200) {
		let res = xhr.responseText;
		
		document.getElementById("message_div").innerText = res;
	    } // TODO on error
	}
    };

    xhr.onerror = function(evt) {
	document.getElementById("message_div").innerText = "ERROR: Failed to call server";
    };
    
    xhr.send();
}


function addSuffix() {
    
    let word = document.getElementById("suffix_input").value.toLowerCase();

    if (word.trim() === "") {
	return;
    }
    
    let select = document.getElementById("decomp_select");
    var decomper = select.options[select.selectedIndex].value;
    let decompURL = baseURL + "/decomp/" + encodeURIComponent(decomper) + "/add_suffix/" + encodeURIComponent(word);
    
    console.log(decompURL);
    
    let xhr = new XMLHttpRequest();
    xhr.open('GET', decompURL);
    xhr.setRequestHeader('Content-type', 'application/x-www-form-urlencoded');
    
    xhr.onload = function(evt) {
	if ( xhr.readyState === 4) {
     	    if (xhr.status === 200) {
		let res = xhr.responseText;
		
		document.getElementById("message_div").innerText = res;
	    } // TODO on error
	}
    };

    xhr.onerror = function(evt) {
	document.getElementById("message_div").innerText = "ERROR: Failed to call server";
    };
    
    xhr.send();
}

function removeSuffix() {
    
    let word = document.getElementById("suffix_input").value.toLowerCase();
    
    if (word.trim() === "") {
	return;
    }
    
    let select = document.getElementById("decomp_select");
    var decomper = select.options[select.selectedIndex].value;
    let decompURL = baseURL + "/decomp/" + encodeURIComponent(decomper) + "/remove_suffix/" + encodeURIComponent(word);
    
    console.log(decompURL);
    
    let xhr = new XMLHttpRequest();
    xhr.open('GET', decompURL);
    xhr.setRequestHeader('Content-type', 'application/x-www-form-urlencoded');
    
    xhr.onload = function(evt) {
	if ( xhr.readyState === 4) {
     	    if (xhr.status === 200) {
		let res = xhr.responseText;
		
		document.getElementById("message_div").innerText = res;
	    } // TODO on error
	}
    };
    
    xhr.onerror = function(evt) {
	document.getElementById("message_div").innerText = "ERROR: Failed to call server";
    };
    
    xhr.send();
}
