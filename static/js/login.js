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
        `https://api.open-meteo.com/v1/forecast?latitude=${lat}&longitude=${lon}&current=temperature_2m,weather_code`,
        {
            method:"GET",
            headers:{
                "Content-Type":"application/json"
            }
        });

        if(!res.ok){
            throw new Error("Weather API failed");
        }

        const data = await res.json();

        const code = data.current?.weather_code || 0;

        // const code = 99; // TESTING
        let weather = "Cloudy";
        let icon = "☁️";


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

        const temp =
        data.current?.temperature_2m ?? "--";

        document.getElementById("weatherText")
        .innerText =
        `${weather} • ${temp}°C`;

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

    // jangan double rain
    if(document.querySelector(".rain")){
        return;
    }

    for(let i=0;i<80;i++){

        const rain = document.createElement("div");

        rain.classList.add("rain");

        rain.style.left =
        Math.random() * 100 + "vw";

        rain.style.animationDuration =
        (Math.random() * 1 + 0.5) + "s";

        rain.style.opacity =
        Math.random();

        rain.style.height =
        (Math.random() * 50 + 30) + "px";

        document.body.appendChild(rain);
    }
}



// REALTIME CLOCK
function updateLoginClock(){

    const now = new Date();

    const time =
    now.toLocaleTimeString("en-GB", {
        hour: "2-digit",
        minute: "2-digit",
        second: "2-digit"
    });

    const date =
    now.toLocaleDateString("id-ID", {
        weekday: "long",
        day: "numeric",
        month: "long",
        year: "numeric"
    });

    document.getElementById("currentTimeLogin")
    .innerText = time;

    document.getElementById("currentDateLogin")
    .innerText = date;
}

updateLoginClock();

setInterval(updateLoginClock, 1000);


// SHOW / HIDE PASSWORD
const togglePassword =
document.getElementById("togglePassword");

const passwordInput =
document.getElementById("passwordInput");

togglePassword.addEventListener("click", ()=>{

    const isPassword =
    passwordInput.type === "password";

    passwordInput.type =
    isPassword ? "text" : "password";

    togglePassword.innerHTML = isPassword
    ? '<i class="fa-solid fa-eye-slash"></i>'
    : '<i class="fa-solid fa-eye"></i>';
});