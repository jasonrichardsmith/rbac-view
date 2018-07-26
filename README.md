

# RBAC View



![RBAC View Screenshot](img/screen.png?raw=true)



RBAC  View is a tool to visualize your RBAC permissions. 

## Current Status

This project is considered prerelease and is under active development.

## Building

From inside this repository...

If you have npm installed and a working Go environment you can run:

```bash
make build
```

If not but you have Docker, you can run:

```bash
make builddocker
```

At the end you should have a binary called rbac-view

## Running

The binary currently only supports the following output modes, and renders based on access to a Kuberentes cluster.

- json to STDOUT
- HTML server

```bash
./rbac-view --render html (default)
```

```
./rbac-view --render json
```

## Developing

### HTML Server Frontend

The web server serves up json representations and static assets compiled into the binary.

The static assets are generated utilizing npm and then compiled into the Go binary using [fileb0x](https://github.com/UnnoTed/fileb0x).

The frontend folder has a generated json file that you can develop against, so frontend developers can start developing right away by running:

```bash
npm run dev
```

## Built with
[Go](https://golang.org/)
[npm](https://www.npmjs.com/)
[Vue.js](https://vuejs.org/)
[fileb0x](https://github.com/UnnoTed/fileb0x)


