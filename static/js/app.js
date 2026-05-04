// ===== SIDEBAR =====
function toggleSidebar() {
    document.querySelector(".sidebar").classList.toggle("collapsed")
    document.querySelector(".content").classList.toggle("expanded")
}

// ===== DROPDOWN =====
function toggleDropdown() {
    const menu = document.getElementById("dropdownMenu");

    if (menu.style.display === "block") {
        menu.style.display = "none";
    } else {
        menu.style.display = "block";
    }
}
// klik luar nutup dropdown
window.onclick = function(e) {
    if (!e.target.closest(".user-dropdown")) {
        const menu = document.getElementById("dropdownMenu");
        if (menu) menu.style.display = "none";
    }
}

// ===== REALTIME CLOCK =====
function updateTime() {
    const now = new Date()

    const date = now.toLocaleDateString("id-ID", {
        weekday: "long",
        year: "numeric",
        month: "long",
        day: "numeric"
    })

    // const time = now.toLocaleTimeString("id-ID")

    // 🔥 format manual biar pakai :
    const hours = String(now.getHours()).padStart(2, '0')
    const minutes = String(now.getMinutes()).padStart(2, '0')
    const seconds = String(now.getSeconds()).padStart(2, '0')

    const time = `${hours}:${minutes}:${seconds}`

    document.getElementById("currentDate").innerText = date
    document.getElementById("currentTime").innerText = time
}

setInterval(updateTime, 1000)
updateTime()

// ===== LOADING =====
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

// auto trigger saat submit form
document.addEventListener("DOMContentLoaded", function () {
    document.querySelectorAll("form").forEach(form => {
        form.addEventListener("submit", () => {
            showLoading()
        })
    })
})

// NAVIGATION HANDLER
function handleNav(e, url) {
    e.preventDefault();

    showLoading("Opening...");

    setTimeout(() => {
        window.location.href = url;
    }, 700);
}


function showPlanDetail() {
    const select = document.getElementById("planSelect");
    const selected = select.options[select.selectedIndex];

    if (!selected.value) {
        document.getElementById("planDetail").style.display = "none";
        return;
    }

    document.getElementById("planDetail").style.display = "block";

    document.getElementById("planName").innerText = selected.dataset.name;
    document.getElementById("planUsers").innerText = selected.dataset.users;
    document.getElementById("planPrice").innerText = selected.dataset.price;
    document.getElementById("planDesc").innerText = selected.dataset.desc;
}

// ===== FILE UPLOAD PREVIEW =====
document.addEventListener("DOMContentLoaded", function () {
    const fileInput = document.getElementById("fileInput");

    if (fileInput) {
        fileInput.addEventListener("change", function () {
            const name = this.files[0]?.name || "Choose file...";
            document.getElementById("fileName").innerText = name;
        });
    }
});