# Echo Gonertia React Example

This is a example for start using the [Golang Inertia Adapter](https://github.com/romsar/gonertia) using Echo & React

This is inspired by [React Example w/ Gonertia](https://github.com/sdil/gonertia_react_example/tree/master)

## Demo

[demo.webm](https://github.com/user-attachments/assets/139e8fc9-340f-41ec-a8b2-c1f3864fe129)

## Technologies

- [Go](https://go.dev/)
- [Echo](https://echo.labstack.com/)
- [Gonertia](https://github.com/romsar/gonertia)
- [go-playground/validator](https://github.com/go-playground/validator)
- [Gorilla Sessions](https://github.com/gorilla/sessions)
- [React](https://react.dev/)
- [TypeScript](https://www.typescriptlang.org/)
- [Inertia.js](https://inertiajs.com/)
- [Tailwind CSS](https://tailwindcss.com/)
- [Vite](https://vite.dev/)
- [Vitest](https://vitest.dev/)
- [Biome](https://biomejs.dev/)

## Technical Details 

- Local and production development with Docker
- Hot reload (server-side and frontend)
- CSRF
- Validation
- Flash messages on validation errors
- Custom error handling
- Simple page display
- Post requests uging Form helper
- Using layouts
- Testing (frontend)

## Instruction

1. Launch app using docker compose

```bash
make compose/up
```

2. Access app

http://localhost:8080/

## Development

Rebuild docker image

```bash
make compose/build
```

Enter docker container

````bash
make shell/app
make shell/front
````

Run lint and format

```bash
make check
```

Run tests

```bash
make test
```

Simulate production environment

```bash
make compose/build/release
make compose/up/release
```

## How to contribute

- Create a local branch or fork
- Upload the changes
- Make PR (understandable)
