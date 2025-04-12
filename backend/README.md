# Shopping Cart Backend

A Go-based backend server implementing the [Order Food Online API](https://orderfoodonline.deno.dev/public/openapi.html) specification.

## Features

- Product listing endpoint (`GET /api/products`)
- Order placement with discount codes (`POST /api/orders`)
- CORS support for frontend integration
- Discount code support:
  - `HAPPYHOURS`: 18% off total order
  - `BUYGETONE`: Get the lowest priced item for free

## Setup

1. Install Go 1.16 or later
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Run the server:
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080`.

## API Endpoints

### GET /api/products
Returns a list of available products.

### POST /api/orders
Places a new order. Request body should be in JSON format:
```json
{
  "items": [
    {
      "productId": "string",
      "quantity": 0
    }
  ],
  "discountCode": "string"
}
```

## Frontend Integration

The backend is configured to work with the Vite-based frontend running on port 5173. CORS is enabled for this origin. 