# media
## Introduction
Media is the service that handles image uploading and also serving images, it interacts with uacl to authorize requests.

It's split into a ngninx container used to serve images and a golang container to handle uploads

## Production Environment Variables
```
HOST - In case the service needs to run on anything other than 0.0.0.0
PORT - In case the service needs to run on anything other than 80
FILE_LOCATION - Locations of files within the container to store images
EMAIL_FROM - Email configuration.
EMAIL_PASSWORD - Email configuration.
EMAIL_LEVEL - What level of logs gets sent to the email address.
ALLOWED_ORIGINS - Cors setup.
```
## Endpoints
```
base URL for static images media.emotives.net
base URL for uploading images media-upload.emotives.net

Requests to media.emotives.net will just return 404 if the image can't be found.

GET - /healthz - Standard endpoint that just returns ok and a 200 status code. Can be used to test if the service is up
POST - /image - User authenticated response that handles image uploads.
POST - /user_profile - User authenticated response that handles user profile image uploads
```