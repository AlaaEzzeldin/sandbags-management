# team2

### Commits to the repo

| Type         | Description                   |
|--------------|-------------------------------|
| **feat**     | new feature being added       |
| **fix**      | fixing the bug                |
| **style**    | stilization updates           |
| **refactor** | refactoring the code          |
| **test**     | everything with testing       |
| **doc**      | everything with documentation |
| **infra**    | everything with deployment    |

Specify the commit type and the message, use present tense inside, do not use semicolon at the end!


### Git Strategy 

The repo follows feature flow


### Run

The application is divided into frontend and backend apps that are put into Docker containers.

Each app has its own `Dockerfile` that are orchestrated through `docker-compose.yml`. Furthermore, postgreSQL database runs own container.

To build and run the project, do `docker-compose up --build`. It will build and start corresponding apps in separate containers.

Frontend app runs on port `:80` and backend is accessible on port `:8001`