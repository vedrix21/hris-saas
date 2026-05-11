const dateElement = document.getElementById("currentDate");

const now = new Date();

const options = {
    weekday: "long",
    year: "numeric",
    month: "long",
    day: "numeric"
};

dateElement.innerText = now.toLocaleDateString("en-US", options);