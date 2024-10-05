# Merchant-Bank API

This project is a simple API that simulates interactions between merchants and banks.

## Features

- Login
- Payment
- Logout

## How to Run

1. Clone the repository:

   ```sh
   git clone https://github.com/randytjioe/merchant-bank-api
   cd merchant-bank-api
   ```

2. Run the application:

   ```sh
   go run main.go
   ```

3. The API will be available at `http://localhost:5050`.

## API Endpoints

- `POST /login`
- `POST /logout`
- `POST /payment`

## Testing

To run tests on Postman:

#### First, log in:

- URL: `POST http://localhost:5050/login`

- Body (JSON):

```json
{
  "username": "JohnDoe",
  "password": "password123"
}
```

#### Successful Login Response

![alt text](<Screenshot from 2024-10-05 13-37-25.png>)

#### Failed Login Response

![alt text](<Screenshot from 2024-10-05 13-37-56-4.png>)

#### Make a Payment:

- URL: `POST http://localhost:5050/payment`

- Body (JSON):

```json
{
  "merchant_id": "12345",
  "amount": 100
}
```

#### Expected Responses:

![alt text](<Screenshot from 2024-10-05 13-37-43.png>)

##### If not logged in

![alt text](<Screenshot from 2024-10-05 13-37-56-1.png>)

##### If balance is insufficient

![alt text](<Screenshot from 2024-10-05 13-37-56-2.png>)

##### If paying an amount of 0 or less

![alt text](<Screenshot from 2024-10-05 13-37-56-3.png>)

#### Logout Response

![alt text](<Screenshot from 2024-10-05 13-37-56.png>)
