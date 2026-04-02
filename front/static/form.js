var formSection = document.getElementsByClassName("form-section")[0];
var editForm = document.getElementsByClassName("edit-form")[0];

function edit(id, text, texts){
	if (id.length < 1) {
		
		editForm.action = "/book";
	} else{
		document.getElementsByClassName("edit-method")[0].value = "PUT";
		editForm.action = "/update/book/"+id;
	}

	document.getElementsByClassName("form-values")[0].value = text;
	document.getElementsByClassName("form-values")[1].value = texts;
	formSection.style.display = "grid";
}

function formCancel(){
	formSection.style.display = "none";
}