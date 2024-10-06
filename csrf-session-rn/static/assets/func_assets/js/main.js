"use strict";
document.addEventListener('DOMContentLoaded', function() {
	content_ready_scripts();
});


function content_ready_scripts() {
	redirectIndexPage('admin-form');
}


function redirectIndexPage(selector) {
	const form = document.getElementById(selector);
	if (form) {
		form.addEventListener('submit', function(e) {
			alert(selector);
			e.preventDefault();
			if (validateForm()) {
				// window.location.href = 'index.html';
        let url = "https://127.0.0.1:8443/public";
        //B get data
        const inputs = form.querySelectorAll('input');
        const formData = {};
        inputs.forEach(input => {
          if(input.name){
            formData[input.name] = input.value;
          }
        });
        console.log('get data...');
        console.log(formData);
        //E get data
				fetch(url, {
						method: 'POST',
						headers: {
							'Content-Type': 'application/json',
              'X-CSRF-Token': formData["csrf"]
						},
						body: JSON.stringify(formData)
					})
					.then(response => response.json())
					.then(data => {
            // B get data return
            alert("data return");
            alert(data);
            console.log('return data...');
						console.log(data);
						// E get data return
					})
					.catch(error => {
						console.error('Error:', error);
					});
			}
		});

		function validateForm() {
			const passwordField = document.getElementById('password');
			if (
				(passwordField.value.trim() === '')
			) {
				alert('Password is required.');
				return false;
			}
			return true;
		}
	}
}