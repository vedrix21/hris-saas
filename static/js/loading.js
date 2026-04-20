function showLoading(message = "Processing...") {
    const el = document.getElementById("aiLoading");
    const text = document.getElementById("loadingMessage");

    if (text) text.innerText = message;
    if (el) el.style.display = "flex";
}

function hideLoading() {
    const el = document.getElementById("aiLoading");
    if (el) el.style.display = "none";
}

function handleNav(e, url) {
    e.preventDefault();

    showLoading("Opening...");

    setTimeout(() => {
        window.location.href = url;
    }, 1000); // delay biar animasi keliatan
}
