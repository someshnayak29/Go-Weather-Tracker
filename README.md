Go-Weather Tracker

Go-Weather Tracker is a simple Go application that fetches weather data from OpenWeatherMap API based on user queries and serves the data over HTTP.

Features

Fetches current weather data for a specified city using OpenWeatherMap API.
Provides a /weather/ endpoint to query weather data for a specific city.
Includes a /hello endpoint to test server connectivity.

Setup:

Clone the repository:

git clone <repository-url>
cd Go-Weather-Tracker

Create .apiConfig file:

Create a file named .apiConfig in the project root and add your OpenWeatherMap API key in JSON format:

{
    "OpenWeatherMapApiKey": "your-api-key-here"
}

/weather/{city} Endpoint
Replace {city} with the name of the city for which you want to fetch weather data.

Response:

{
    "name": "New York",
    "main": {
        "temp": 293.15
    }
}

