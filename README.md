# voters-id-reader

A demo project to extract data from the Ghana Voters' ID card using Go, Gin and and Zxing golang port. This project is mainly to help me understand how to handle form data in Golang and text and image manipulation. This should not be used in production.

## Description

This project is a demo project to extract data from the Ghana Voters' ID card using Go, Gin and and the Zxing port for Golang. It works by uploading an image of the ID card and the application extracts the data from the image and returns it as a JSON response.

## Live Environment

- The Deployed Version of the API can be found at [https://voters-id-reader.onrender.com](https://voters-id-reader.onrender.com)

- The API documentation for the deployment can be found at [https://documenter.getpostman.com/view/8806007/2s8Z6x4a5Z](https://documenter.getpostman.com/view/8806007/2s8Z6x4a5Z).

## Setting up the project locally

- Clone the repository

```bash
git clone https://github.com/buabaj/voters-id-reader
```

- Navigate to the project directory.

```bash
cd voters-id-reader
```

- Run the application

```bash
go run main.go
```

- Test the endpoints using Postman or any other API testing tool

## Endpoints

| Name | Request Method | Endpoint | Parameters | Response code |
| --- | --- | --- | --- | --- |
| Index | GET | localhost:8080/ | None | 200 OK |
| Read Data from ID | POST | localhost:8080/read | None | 200 OK |
|

## Sample Data

Sample JSON data to test `read` endpoint

```form-data
{
    "file": "path/to/image/file"
}
```


Sample JSON response to `read` endpoint

```json
{
    "data": {
        "dateOfBirth": "1900-01-10",
        "dateOfRegistration": "2020-01-10",
        "firstName": "Test",
        "gender": "Male",
        "idNumber": "xxxxxxxxxx",
        "pollingStationCode": "xxxxxxx",
        "surname": "Test",
    },
    "message": "ID successfully decoded",
    "status": "success"
}
```

## Future Work

- [x] Add tests
- [ ] Learn how to test endpoints that require file uploads
- For this current version, the image needs to be clear and fit the corners of the ID card. This will be improved in the future to allow for images that are not clear and do not fit the corners of the ID card. That is, the application will be able to detect the ID card in the image and crop it out.
