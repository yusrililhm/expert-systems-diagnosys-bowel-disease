document.addEventListener("DOMContentLoaded", function () {
    const warn = document.getElementById("warn-text");
    const minInputLength = document.getElementById("password").minLength;

    
    document.getElementById("password").addEventListener("input", function () {
        const inputLength = document.getElementById("password").value.length;

        if (inputLength >= minInputLength) {
            warn.style.visibility = "hidden"
        }
    });

    document.getElementById("password").addEventListener("focus", function () {
        warn.style.visibility = "visible";
    });

    document.getElementById("password").addEventListener("blur", function () {
        warn.style.visibility = "hidden"
    });
});
