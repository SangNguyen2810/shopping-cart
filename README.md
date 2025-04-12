# Shopping Cart Application

## Project Overview

This full-stack e-commerce shopping cart application demonstrates modern web development practices with a React/TypeScript frontend and a Go backend. The application features product browsing, cart management, discount code application, and order processing with comprehensive error handling and state management.

## Key Features

- **Product Catalog**: Browse products with images, descriptions, and prices
- **Shopping Cart**: Add, remove, and update product quantities in real-time
- **Discount System**: Apply voucher codes with server-side validation
  - Special `HAPPYHOURS` promotion with 18% discount (10 unique codes in Redis)
  - `BUYGETONE` discount that gives the lowest priced item for free
  - Visual feedback for successful/failed code application
- **Order Processing**: Secure checkout with order confirmation
- **Responsive Design**: Optimized for mobile, tablet, and desktop
- **Error Handling**: Comprehensive error management across frontend and backend
- **State Management**: Centralized state using Zustand stores

## Technology Stack

### Frontend
- **Framework**: React 18 with TypeScript
- **UI Library**: Mantine components
- **Styling**: Tailwind CSS
- **State Management**: Zustand
- **Form Validation**: Zod
- **Build Tool**: Vite
- **HTTP Client**: Native Fetch API with custom wrapper

### Backend
- **Language**: Go
- **Framework**: Gin
- **Data Storage**:
  - In-memory for product data
  - Redis for discount code management
- **API Documentation**: Swagger
- **Validation**: Custom validation middleware

## Architecture

### Frontend Architecture

The frontend follows a component-based architecture with state management using Zustand:

```
frontend/src/
├── api/                  # API integration layer
│   ├── config.ts         # API configuration and fetch wrapper
│   ├── discount-api.ts   # Discount code validation
│   ├── order-api.ts      # Order processing
│   └── product-api.ts    # Product fetching
├── components/           # Reusable UI components
│   ├── Cart/             # Cart-related components
│   ├── ProductItem/      # Product display components
│   └── ...
├── constants/            # Application constants
├── pages/                # Page components
├── store/                # Zustand state stores
│   ├── cartStore.ts      # Shopping cart state
│   ├── voucherStore.ts   # Voucher/discount state
│   └── ...
├── types/                # TypeScript type definitions
└── utils/                # Utility functions
```

### Backend Architecture

The backend follows a clean architecture pattern with clear separation of concerns:

```
backend/
├── config/               # Application configuration
├── data/                 # Data sources and seed data
├── handlers/             # HTTP request handlers
│   ├── discount.go       # Discount code validation
│   ├── order.go          # Order processing
│   └── product.go        # Product listing
├── middleware/           # HTTP middleware
├── models/               # Data models
├── repositories/         # Data access layer
└── services/             # Business logic
    ├── discount.go       # Discount code management with Redis
    ├── order.go          # Order processing logic
    └── product.go        # Product management
```

### Data Flow

1. User browses products loaded from backend via API
2. Cart state is managed locally in the frontend
3. Discount codes are validated against backend Redis storage
4. Orders are processed through the backend with validation
5. Confirmation is returned to frontend for display

## Local Development Setup

### Prerequisites

- Node.js (v16+)
- Go (v1.18+)
- Redis (v6+)

### Frontend Setup

1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Start the development server:
   ```bash
   npm run dev
   ```

4. Access the application at http://localhost:5173

### Backend Setup

1. Navigate to the backend directory:
   ```bash
   cd backend
   ```

2. Install Go dependencies:
   ```bash
   go mod tidy
   ```

3. Ensure Redis is running:
   ```bash
   # Start Redis with Docker (optional)
   docker run -d -p 6379:6379 --name shopping-cart-redis redis:alpine
   ```

4. Start the backend server:
   ```bash
   go run main.go
   ```

5. The API will be available at http://localhost:8080/api
6. Swagger documentation is available at http://localhost:8080/swagger/index.html

## Feature Implementation Details

### Product Catalog
Products are stored in the backend and exposed through a RESTful API. The frontend fetches these products and displays them in a responsive grid layout. Each product card shows an image, name, price, and category.

### Shopping Cart
The cart is implemented using Zustand for state management, providing a seamless experience with:
- Add to cart with quantity selection
- Remove items
- Update quantities
- Real-time total calculation
- Persistent cart across page reloads

### Discount Code System
Voucher codes are managed through a Redis-backed system:
1. 10 unique `HAPPYHOURS` codes (0-9) are stored in Redis
2. When a user enters a code, it's validated through the `/discounts/validate` endpoint
3. Valid codes provide an 18% discount on the order total
4. `BUYGETONE` discount removes the lowest priced item from the total
5. Once applied, the discount is displayed with amount and percentage
6. Used codes are locked in Redis to prevent reuse
7. Error messages for invalid codes are displayed directly below the input field

Implementation highlights:
- Redis TTL (Time-to-Live) set to 24 hours for used codes
- Frontend validation with error state management
- Real-time calculation of discount amounts
- Reset functionality after order completion

### Order Processing
Orders are processed through a secure checkout flow:
1. User reviews cart contents and total
2. Applies discount code (optional)
3. Confirms order
4. Backend validates product availability and pricing
5. Confirmation is returned with order ID and details
6. Cart is cleared and ready for next purchase

## Code Quality & Best Practices

### Frontend
- **TypeScript**: Strong typing throughout for code robustness
- **Component Reusability**: Modular components with clear props interfaces
- **Responsive Design**: Mobile-first approach with Tailwind
- **Accessibility**: Semantic HTML and ARIA attributes
- **Error Handling**: Comprehensive error states and user feedback
- **State Management**: Isolation of concerns using multiple stores

### Backend
- **Clean Architecture**: Separation of concerns with handlers/services/repositories
- **Error Handling**: Structured error responses with appropriate HTTP codes
- **Validation**: Input validation at multiple levels
- **Documentation**: Swagger API documentation
- **Configuration**: Environment-based configuration system
- **Graceful Degradation**: Fallbacks when Redis is unavailable

## Testing

### Frontend Tests
Run the frontend tests with:
```bash
cd frontend
npm test
```

### Backend Tests
Run the backend tests with:
```bash
cd backend
go test ./...
```

## Deployment Considerations

### Frontend Deployment
The frontend can be built for production with:
```bash
cd frontend
npm run build
```

The resulting `dist` directory can be served from any static hosting service.

### Backend Deployment
The backend can be compiled to a binary with:
```bash
cd backend
go build -o shopping-cart-api
```

For containerization, a sample Dockerfile is included in the repository.

## Future Enhancements

1. **User Authentication**: Add login/registration for personalized experiences
2. **Payment Processing**: Integrate payment gateway for real transactions
3. **Product Search**: Add search functionality with filters
4. **Product Reviews**: Allow customers to rate and review products
5. **Order History**: Track past orders for logged-in users
6. **Wishlist**: Save products for future purchase
7. **Multi-language Support**: Internationalization for global customers

## Original Requirements

Based on the [API Documentation](https://orderfoodonline.deno.dev/public/openapi.html), this project implemented:

- Product listing with images
- Shopping cart functionality
- Order total calculation
- Quantity management in cart
- Order confirmation process
- Interactive states for elements
- Discount code system with `HAPPYHOURS` (18% discount) and `BUYGETONE` (free lowest-priced item)
- Responsive design for all device sizes
- Custom backend API implementation in Go

## License

This project is licensed under the MIT License - see the LICENSE file for details.

