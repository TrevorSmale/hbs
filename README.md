# hbs
Full-Stack Hotel Booking Site

## Overview

![Linux](https://img.shields.io/badge/Linux-FCC624?style=for-the-badge&logo=linux&logoColor=black)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Git](https://img.shields.io/badge/git-%23F05033.svg?style=for-the-badge&logo=git&logoColor=white)
![Visual Studio Code](https://img.shields.io/badge/Visual%20Studio%20Code-0078d7.svg?style=for-the-badge&logo=visual-studio-code&logoColor=white)
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

Utilizing Go for the Backend with Modules and Efficient Template Caching. A PostgresQL Database stores booking information. The Interactive Front End uses the Bootstrap JS framework. The whole project builds to a Podman container running on Alpine Linux for security and efficiency.

## Efficiency

The code is structured in such a way that page template rendering is checked against cached pages, reducing the number of disk calls.

## Capacity

GO's http stack is known to be fast and reliable. I know that this project can be served to 1000's of concurent users while handling errors and remaining memory safe on modestly spec'd hardware or VM's

## Portability

The project is build using go Modules in a Containerized setup, making this project highly portable and capable of running on a baremetal, VM's or in an orchestrated environment.



