function changeModal(name){
    // close any modal 
    $('modal-container').addClass('out');
    $('body').removeClass('modal-active');

    // open fresh modal :: OPTION : unfolding, revealing, meep-meep, sketch, bond
    var animation = "meep-meep";
    $('#'+name+' modal-container').removeAttr('class').addClass(animation);
    $('body').addClass('modal-active');
}
    
  
$('.close-animated-modal').click(function(){
    $('modal-container').addClass('out');
    $('body').removeClass('modal-active');
});

$("#signup-validation, #login-validation").parsley({
    errorClass: 'is-invalid text-danger',
    successClass: 'is-valid', // Comment this option if you don't want the field to become green when valid. Recommended in Google material design to prevent too many hints for user experience. Only report when a field is wrong.
    errorsWrapper: '<span class="form-text text-danger"></span>',
    errorTemplate: '<span></span>',
    trigger: 'change'
})
$('#signup-validation').parsley().on('form:submit', function(formInstance) {
    return false;
});

$('#login-validation').parsley().on('form:submit', function(formInstance) {
    return false;
});

$("#signup_call").click(function () {
    if ( ! $("#signup-validation").parsley().isValid() ) {
        return false;
    }
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
                alertify.success("Signup Succesfull. Please login");
            }else{
                alertify.error(response.message);
            }
        },
    });

    
});



$("#login_call").click(function () {
    if ( ! $("#login-validation").parsley().isValid() ) {
        return false;
    }
    var username = $("#login-username").val();
    var password = $("#login-password").val();

    $.ajax({
        url: "/api/go_login",
        method: "POST",
        data: {
            "username":   username,
            "password":   password
        },
        contentType: "application/x-www-form-urlencoded",
        success: function(response) {
            console.log(response);
            if(response.STATUS == 'SUCCESS'){
                alertify.success("login Succesfull. Please login");
            }else{
                alertify.error(response.message);
            }
        },
    });

    
});