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

        // const code = data.current.weather_code;

        const code = 99; // TESTING
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

// LOCATION
if(navigator.geolocation){

    navigator.geolocation.getCurrentPosition(

        async (position)=>{

            const lat = position.coords.latitude;
            const lon = position.coords.longitude;

            loadWeather(lat, lon);

            // GET CITY
            try{

                const res = await fetch(
                    `https://geocode.maps.co/reverse?lat=${lat}&lon=${lon}`
                );

                const data = await res.json();

                document.getElementById("weatherLocation")
                .innerText =
                data.address.city ||
                data.address.town ||
                data.address.village ||
                data.address.state ||
                "Indonesia";

            }catch(e){

                console.log(e);

                document.getElementById("weatherLocation")
                .innerText = "Indonesia";
            }

        },

        ()=>{

            document.getElementById("weatherText")
            .innerText =
            "Location denied";

        }
    );
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