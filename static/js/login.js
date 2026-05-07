// LOGIN LOADING
const form = document.getElementById("loginForm");
const btn = document.getElementById("loginBtn");

form.addEventListener("submit", function () {

    btn.classList.add("loading");

    btn.innerHTML = "Signing In...";

    btn.disabled = true;
});

// WEATHER
async function loadWeather(lat, lon) {

    try {

        const res = await fetch(
            `https://api.open-meteo.com/v1/forecast?latitude=${lat}&longitude=${lon}&current=temperature_2m,weather_code`
        );

        const data = await res.json();

        const code = data.current.weather_code;

        // const code = 99; // TESTING
        let weather = "Cloudy";
        let icon = "☁️";

        // HIDE LIGHTNING TOGGLE
        document.getElementById("effectToggle")
        .style.display = "none";

        if(code === 0){
            weather = "Clear Sky";
            icon = "☀️";
        }
        else if(code <= 3){
            weather = "Cloudy";
            icon = "☁️";
        }
        else if(code <= 67){
            weather = "Rainy";
            icon = "🌧️";

            document.body.classList.add("rainy");
            createRain();
        }
        else if(code <= 99){
            weather = "Storm";
            icon = "⛈️";

            document.body.classList.add("rainy");

            createRain();

        }

        document.getElementById("weatherText")
        .innerText =
        `${weather} • ${data.current.temperature_2m}°C`;

        document.getElementById("weatherIcon")
        .innerText = icon;

    } catch(err){

        console.log(err);

        document.getElementById("weatherText")
        .innerText = "Weather unavailable";
    }
}

// LOCATION + CITY
if (navigator.geolocation) {

    navigator.geolocation.getCurrentPosition(

        async (position) => {

            const lat = position.coords.latitude;
            const lon = position.coords.longitude;

            // LOAD WEATHER
            loadWeather(lat, lon);

            // LOAD CITY
            try {

                const res = await fetch(
                    `https://api.bigdatacloud.net/data/reverse-geocode-client?latitude=${lat}&longitude=${lon}&localityLanguage=id`
                );

                const data = await res.json();

                const city =
                    data.city ||
                    data.locality ||
                    data.principalSubdivision ||
                    "Indonesia";

                document.getElementById("weatherLocation")
                .innerText = city;

            } catch (e) {

                console.log(e);

                document.getElementById("weatherLocation")
                .innerText = "Indonesia";
            }

        },

        () => {

            document.getElementById("weatherText")
            .innerText = "Location denied";

            document.getElementById("weatherLocation")
            .innerText = "Indonesia";
        }

    );

} else {

    document.getElementById("weatherLocation")
    .innerText = "Indonesia";
}


// AUTO DAY MODE
function setDayMode(){

    const hour = new Date().getHours();

    document.body.classList.remove(
        "morning",
        "afternoon",
        "evening",
        "night"
    );

    if(hour >= 5 && hour < 11){

        document.body.classList.add("morning");

    }else if(hour >= 11 && hour < 16){

        document.body.classList.add("afternoon");

    }else if(hour >= 16 && hour < 19){

        document.body.classList.add("evening");

    }else{

        document.body.classList.add("night");
    }
}

setDayMode();

// RAIN EFFECT
function createRain(){

    for(let i=0;i<80;i++){

        const rain = document.createElement("div");

        rain.classList.add("rain");

        rain.style.left = Math.random() * 100 + "vw";

        rain.style.animationDuration =
        (Math.random() * 1 + 0.5) + "s";

        rain.style.opacity = Math.random();

        rain.style.height =
        (Math.random() * 50 + 30) + "px";

        document.body.appendChild(rain);
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

    document.getElementById("currentDateLogin").innerText = date
    document.getElementById("currentTimeLogin").innerText = time
}

setInterval(updateTime, 1000)
updateTime()