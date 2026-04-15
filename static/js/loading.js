function showLoading(text = "Processing...") {
    const loader = document.getElementById("globalLoading");
    const loadingText = document.getElementById("loadingText");

    if (loader) loader.style.display = "flex";
    if (loadingText) loadingText.innerText = text;
}

function hideLoading() {
    const loader = document.getElementById("globalLoading");
    if (loader) loader.style.display = "none";
}