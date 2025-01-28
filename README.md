# REST API in Go 

This project is a simple REST API created with languaje GO that allows managing the SHIPPMENT domain, specifically for the UPDATE (state) microservice. The API offers the basic operation such as updating a state shippment, displaying the SWAGGER documentation technology screen as the main page.
## Project Structure

- **`main.go`**: The main class that runs the Go application and defines the API controller.

- `PATCH /api/shippment//update-state/{id}`: Allows you to update state the shippmetn, under the required ID.

## Requirements

- **GO 1.19** o superior
- **Gestor de dependencias como go mod.**

## Installation

1. **Clone the repository**

    ```bash
    git clone <https://github.com/kevinseya/microservicio-logistic-shipping-status-update.git>
    ```

2. **Install dependencies**

    ```bash
    go mod tidy
    ```

3. The application run on: `http://localhost:8080`.

## Use of endpoint

### 1. PATCH /api/shippment//update-state/{id}

Update a state of shipment. The request body must contain the user details in JSON format.
PATCH request example:
```bash
PATCH /api/shippment//update-state/{id} Content-Type: application/json
    
    {
    "state": 'In transit',
    }
```
**Response:**
```plaintext
   {
    "message": "Shipment status updated successfully",
    "shipment": {
            "id": "3287d8c6-9f50-4a35-92ec-9e5f568a6753",
            "state": "In Transit"
     }
   }
```
**Response code:**
- **`200 Update:`** Shipment status updated successfully.
- **`400 Bad Request:`** State is required in the request body.
- **`404 Not Found:`**  Shipment not found for the given ID.
- **`500 Internal Server Error:`** Server error.
