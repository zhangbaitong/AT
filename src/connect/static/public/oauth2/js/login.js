
function gologin(){
	 document.getElementById("register").style.cssText = "display: none;";
	 document.getElementById("forgetpwd").style.cssText = "display: none;";
	 document.getElementById("login").style.cssText = "display: block;";
	 
}
function goregister(){
	 document.getElementById("login").style.cssText = "display: none;";
	 document.getElementById("forgetpwd").style.cssText = "display: none;";
	 document.getElementById("register").style.cssText = "display: block;";
}
var timer;
var delay_seconds=30;
var pass_seconds=0;
var verifyCodeSend;
function CheckSendVerifyCode(){
	 pass_seconds=0;
	 //Send
	 verifyCodeSend=document.getElementById("verifyCodeSend");
	 verifyCodeSend.disabled="disabled";
	 timer = window.setInterval(SendVerifyCode,1000);
}
function SendVerifyCode(){
   ++ pass_seconds;
   verifyCodeSend.value="重新发送("+(delay_seconds-pass_seconds)+")";
   verifyCodeSend.className = "pass-button pass-text-input-disabled pass-button-verifyCodeSend ";
   
   if (pass_seconds >= delay_seconds)    
   {    
      window.clearInterval(timer); 
      if(delay_seconds<300){
        delay_seconds+=delay_seconds;
      }
      verifyCodeSend.value="发送验证码";
      verifyCodeSend.disabled="";
      verifyCodeSend.className = "pass-button pass-button-verifyCodeSend";
   }
}


function onblurAccount() {
   var is_email = /^([a-zA-Z0-9]+[_|\-|\.]?)*[a-zA-Z0-9]+@([a-zA-Z0-9]+[_|\-|\.]?)*[a-zA-Z0-9]+\.[a-zA-Z]{2,3}$/; 
   var is_num = /^1\d{10}$/

	 if(is_email.test(document.getElementById("account").value)){
	 	 document.getElementById("accountError").style.cssText = "display: none;";
	 	 document.getElementById("verifyCodeImgWrapper").style.cssText = "height: 45px;display: none;";
	 	 document.getElementById("register_submit").disabled="";
	 	 return true;
	 }else if(is_num.test(document.getElementById("account").value)){
	 	 document.getElementById("accountError").style.cssText = "display: none;";
	 	 document.getElementById("register_submit").disabled="";
	 	 document.getElementById("verifyCodeImgWrapper").style.cssText = "height: 45px;display: block;";
	 	 return true;
	 }else{
	   document.getElementById("register_submit").disabled="disabled";
	   document.getElementById("accountError").style.cssText = "display: block;";
	   document.getElementById("verifyCodeImgWrapper").style.cssText = "height: 45px;display: block;";
	   return false;
	 }
}
function onblurPassword(){
	if(document.getElementById("password").value.length>=6){
		document.getElementById("register_submit").disabled="";
		document.getElementById("passwordError").style.cssText = "display: none;";
		return true;
	}else{
		 document.getElementById("passwordError").style.cssText = "display: block;";
		 document.getElementById("register_submit").disabled="disabled";
		 return false;
	}
}
function keyupAccount(){
	if(document.getElementById("register_submit").disabled){
	 onblurAccount();
  }
}

function keyupPassword(){
	if(document.getElementById("register_submit").disabled){
	 onblurPassword();
  }
}

function validate_registerform(){
	if(onblurAccount() && onblurPassword()){
		return true;
	}else{
		return false;
	}
}

function goforgetpwd(){
	 document.getElementById("login").style.cssText = "display: none;";
	 document.getElementById("register").style.cssText = "display: none;";
	 document.getElementById("forgetpwd").style.cssText = "display: block;";
}
function switchpage(a) {
    switch (a) {
    case 1:
        document.getElementById("switcher_qr_login").className = "switch_btn_focus";
        document.getElementById("switcher_web_login").className = "switch_btn";
        document.getElementById("switch_bottom").style.cssText = "left: 0px; width: 64px; position: absolute;";
        document.getElementById("qr_login").style.cssText = "display: block;";
        document.getElementById("web_login").style.cssText = "display: none;";
        document.getElementById("bottom_qr_login").style.cssText = "display: block;";
        document.getElementById("web_register").style.cssText = "display: none;";
        break;
    case 2:
        document.getElementById("switcher_qr_login").className = "switch_btn";
        document.getElementById("switcher_web_login").className = "switch_btn_focus";
        document.getElementById("switch_bottom").style.cssText = "left: 154px; width: 96px; position: absolute;";
        document.getElementById("qr_login").style.cssText = "display: none;";
        document.getElementById("web_login").style.cssText = "display: block;";
        document.getElementById("bottom_qr_login").style.cssText = "display: none;";
        document.getElementById("web_register").style.cssText = "display: none;";
        break;
    }
}
function register(){
	document.getElementById("switcher_qr_login").className = "switch_btn";
  document.getElementById("switcher_web_login").className = "switch_btn_focus";
  document.getElementById("switch_bottom").style.cssText = "left: 154px; width: 96px; position: absolute;";
  document.getElementById("web_login").style.cssText = "display: none;";
	document.getElementById("qr_login").style.cssText = "display: none;";
  document.getElementById("web_register").style.cssText = "display: block;";
}
function onfocusU(a) {
    document.getElementById("uin_tips").className = "input_tips_focus";
    document.getElementById("u").parentNode.className = "inputOuter_focus";
}
function onblurU(a) {
    document.getElementById("uin_tips").className = "input_tips";
    document.getElementById("u").parentNode.className = "inputOuter";
}
function onfocusP(a) {
    document.getElementById("pwd_tips").className = "input_tips_focus";
    document.getElementById("p").parentNode.className = "inputOuter_focus";
}
function onblurP(a) {
    document.getElementById("pwd_tips").className = "input_tips";
    document.getElementById("p").parentNode.className = "inputOuter";
}


function onfocusU1(a) {
    document.getElementById("uin_tips1").className = "input_tips_focus";
    document.getElementById("u1").parentNode.className = "inputOuter_focus";
}
function onblurU1(a) {
    document.getElementById("uin_tips1").className = "input_tips";
    document.getElementById("u1").parentNode.className = "inputOuter";
}
function onfocusP1(a) {
    document.getElementById("pwd_tips1").className = "input_tips_focus";
    document.getElementById("p1").parentNode.className = "inputOuter_focus";
}
function onblurP1(a) {
    document.getElementById("pwd_tips1").className = "input_tips";
    document.getElementById("p1").parentNode.className = "inputOuter";
}

function keyupU() {
    if (document.getElementById("u").value == "") {
        document.getElementById("uin_tips").className = "input_tips_focus";
        document.getElementById("uin_del").style.cssText = "display: none;";
        document.getElementById("uin_tips").style.cssText = "display: block;";
    } else {
        document.getElementById("uin_del").style.cssText = "display: block;";
        document.getElementById("uin_tips").style.cssText = "display: none;";
    }
}
function uindel() {
    document.getElementById("uin_del").style.cssText = "display: none;";
    document.getElementById("uin_tips").style.cssText = "display: block;";
    document.getElementById("uin_tips").className = "input_tips";
    document.getElementById("u").parentNode.className = "inputOuter";
    document.getElementById("u").value = "";
}
function detectCapsLockP(c) {
    var b = c.keyCode || c.which;
    var a = c.shiftKey || (b == 16) || false;
    if (((b >= 65 && b <= 90) && !a) || ((b >= 97 && b <= 122) && a)) {
        document.getElementById("caps_lock_tips").style.cssText = "display: block;";
    } else {
        document.getElementById("caps_lock_tips").style.cssText = "display: none;";
    }
    document.getElementById("pwd_tips").style.cssText = "display: none;";
}
function keyupP() {
    if (document.getElementById("p").value == "") {
        document.getElementById("pwd_tips").className = "input_tips_focus";
        document.getElementById("pwd_tips").style.cssText = "display: block;";
    } else {
        document.getElementById("pwd_tips").style.cssText = "display: none;";
    }
}