(function() {
    'use strict';
    window.addEventListener('load', function() {
        // Fetch all the forms we want to apply custom Bootstrap validation styles to
        var forms = document.getElementsByClassName('signup-validation');
        // Loop over them and prevent submission
        var validation = Array.prototype.filter.call(forms, function(form) {
            form.addEventListener('submit', function(event) {
            if (form.checkValidity() === false) {
                    event.preventDefault();
                    event.stopPropagation();
                }
                callSubmit();
                form.classList.add('was-validated');
            }, false);
        });
    }, false);
})();



$( "#signin-username" ).keydown(function() {
    // alert( "Handler for .keydown() called." );
    $( ".signup-validation" ).addClass( "was-validated" );
});



function callSubmit() {
    
   
    var username = $("#signin-username").val();
    var email = $("#signin-email").val();
    var password = $("#signin-password").val();
    var mobile = $("#signin-mobile").val();
    var country = $("#signin-country option:selected").attr('data-val');
    var country_code = $("#signin-country option:selected").attr('value');

    $.ajax({
        url: "/api/go_signup",
        method: "POST",
        data: {
            "username":   username,
            "email":      email,
            "password":   password,
            "mobile":     mobile,
            "country" :   country, 
            "country_code" : country_code 
        },
        contentType: "application/x-www-form-urlencoded",
        success: function(response) {
            console.log(response);
            if(response.STATUS == 'SUCCESS'){
                console.log("SUCC");
            }else{
                console.log("ERR");
                alert(response.message);
            }
        },
    });

    
};