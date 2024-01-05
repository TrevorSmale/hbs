# Full-Stack Hotel Booking Site

# 🚧 Under Construction 🚧

## Overview

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Git](https://img.shields.io/badge/git-%23F05033.svg?style=for-the-badge&logo=git&logoColor=white)
![Linux](https://img.shields.io/badge/Linux-FCC624?style=for-the-badge&logo=linux&logoColor=black)


Utilizing Go for the Backend with Modules and Efficient Template Caching. A PostgresQL Database stores booking information. The Interactive Front End uses the Bootstrap JS framework. The whole project builds to a Podman container running on Alpine Linux for security and efficiency.

## Efficiency
The code is structured in such a way that page template rendering is checked against cached pages, reducing the number of disk calls.

## Capacity
GO's http stack is known to be fast and reliable. I know that this project can be served to 1000's of concurent users while handling errors and remaining memory safe on modestly spec'd hardware or VM's

## Portability
The project is build using go Modules in a Containerized setup, making this project highly portable and capable of running on a baremetal, VM's or in an orchestrated environment.

## Project Structure 🏗️

- **cmd/web**: This is where the `main.go` file resides, orchestrating the web application. It takes a little inspiration from the command line as it directs the show.

- **pkg**: Ah, the heart of it all! This directory hosts the handler and renderer, neatly tucked away in their respective corners. Think of it as the backstage area where the magic happens.

- **pkg/handler**: Here dwells the handler, ready to handle requests with finesse. It's like the actor on stage, reacting to the audience (or, in this case, user requests) gracefully.

- **pkg/renderer**: Right next to the handler, the renderer plays its part in crafting the perfect scenes. It's the artist, ensuring the audience (your users) gets a visually appealing experience.

- **templates**: This is where the Front-End and Client-Side shine! Imagine it as the canvas for your artistic vision, with separate modules enhancing the overall spectacle.

- **Go Version**: Go is built for backward compatability. However, new features and conventions arise. This project uses GO Version 1.21.5. In the go.mod file, you will find the version lock number.

### Go Version and Modules 🚀

This project is keeping up with the times! Embracing Go modules introduced in Go 1.11, we've bid a cheerful farewell to the older workspaces. Modules bring order to the chaos and let us manage dependencies with ease.

### Containerfile for Podman 📦

The `Containerfile` is a key player in orchestrating the magic with Podman. It's like the script for our container, dictating how the show should run. 

