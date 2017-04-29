function closeToast(){
    $(".toastBg").remove()
}
function yifanToast(message,size) {
    var toastContent= '<div class="toastBg">'+'<div class="toastBox '+size+'">'+'<div class="toastHeight30"></div>'+' <div class="toastMessage">'+' <p id="message">'+message+'   </p>'+' <div class="button">'+'<div class="toastHeight"></div>'+'<button class="toastButton" onclick="closeToast()">确定</button>'+'</div>'+' </div>'+' </div>';
    $("html").prepend(toastContent);
}
