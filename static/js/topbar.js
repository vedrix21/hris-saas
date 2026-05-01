function updateTime() {
    const now = new Date();

    const date = now.toLocaleDateString("id-ID", {
        weekday: "long",
        day: "numeric",
        month: "long",
        year: "numeric"
    });

    const time = now.toLocaleTimeString("id-ID");

    document.getElementById("currentDate").innerText = date;
    document.getElementById("currentTime").innerText = time;
}

// update tiap detik
setInterval(updateTime, 1000);
updateTime();

// dropdown
function toggleDropdown() {
    const menu = document.getElementById("dropdownMenu");
    menu.style.display = menu.style.display === "block" ? "none" : "block";
}

// klik luar nutup dropdown
window.onclick = function(e) {
    if (!e.target.closest(".user-dropdown")) {
        document.getElementById("dropdownMenu").style.display = "none";
    }
}