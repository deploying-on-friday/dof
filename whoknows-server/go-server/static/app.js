fetch("/api/weather")
.then(res => res.json())
.then(data => {
    console.log(data);

    const weatherDiv = document.getElementById("weather");

    const current = data.data.current_weather;

    weatherDiv.innerHTML = `
        <p>Temperature: ${current.temperature} °C</p>
        <p>Windspeed: ${current.windspeed} km/h</p> 
    `;
})
.catch(error => {
    console.error(error);
    document.getElementById("weather").innerText = "Error loading weather";
});
