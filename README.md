# Exoplanet Service

A Go microservice to manage and study exoplanets.

## Prerequisites

- Go 1.19+
- Docker (optional, for containerization)

## Project Structure

API Endpoints

1. Add an Exoplanet

Gas Giant

curl -X POST http://localhost:8080/exoplanets \
-H "Content-Type: application/json" \
-d '{
    "name": "Jupiter-2",
    "description": "A gas giant similar to Jupiter",
    "distance": 50,
    "radius": 1.2,
    "type": "GasGiant"
}'

Terrestrial

curl -X POST http://localhost:8080/exoplanets \
-H "Content-Type: application/json" \
-d '{
    "name": "Kepler-442b",
    "description": "A terrestrial planet larger than Earth",
    "distance": 120,
    "radius": 1.5,
    "mass": 5.3,
    "type": "Terrestrial"
}'

2. List All Exoplanets

curl -X GET http://localhost:8080/exoplanets

3. Get Exoplanet by ID

curl -X GET http://localhost:8080/exoplanets/{id}

4. Update an Exoplanet

Gas Giant

curl -X PUT http://localhost:8080/exoplanets/{id} \
-H "Content-Type: application/json" \
-d '{
    "name": "Jupiter-2 Updated",
    "description": "An updated description of Jupiter-2",
    "distance": 55,
    "radius": 1.3,
    "type": "GasGiant"
}'

Terrestrial

curl -X PUT http://localhost:8080/exoplanets/{id} \
-H "Content-Type: application/json" \
-d '{
    "name": "Kepler-442b Updated",
    "description": "An updated description of Kepler-442b",
    "distance": 130,
    "radius": 1.6,
    "mass": 5.5,
    "type": "Terrestrial"
}'


5. Delete an Exoplanet


curl -X DELETE http://localhost:8080/exoplanets/{id}


6. Fuel Estimation

curl -X GET "http://localhost:8080/exoplanets/{id}/fuel?crewCapacity=10"
