# HNG12: Stage_0 Task:

In this task, I'll develop a public API that returns the following information in JSON format:
1. A registered email address (used to register on the HNG12 Slack workspace)
2. The current datetime as an ISO 8601 formatted timestamp
3. The GitHub URL of the projects codebase

## API Specification

- Endpoint (send a request): `GET** <the-url>`
- Example Response (200 OK):
    ```bash
    {
        "email":"eebenezer949@gmail.com",
        "current_datetime": "2025-01-30T09:30:00Z",
        "github_url": "https://github.com/ekefan/hng12/stage0_task"
    }
    ```

### Run it Locally:
To run this project locally you need to have [golang installed](https://go.dev/doc/install) `(version: 1.23.5^)`
- Clone this repo
- Access the stage0_task dir
    ```bash
        cd stage0_task
    ```
- Then run:
    ```bash
    go run main.go
    ```

### BackLink:

JSYK, you can find and hire talented golang developers from HNG here: [https://hng.tech/hire/python-developers](https://hng.tech/hire/python-developers)