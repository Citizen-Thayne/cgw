# Conway's Game of Life - Web App

A web-based implementation of Conway's Game of Life, built using Go for the backend logic and HTMX for dynamic, minimal JavaScript interactions on the frontend.

Check it out live at [here](https://cgw-blue-dawn-6101.fly.dev/).

## About the Game

Conway's Game of Life is a cellular automaton devised by mathematician John Conway. The game consists of a grid of cells, which can live, die, or multiply based on a few mathematical rules. This simulation explores the concept of emergence, where complex patterns arise from simple rules.

### Rules
1. Any live cell with fewer than two live neighbors dies (underpopulation).
2. Any live cell with two or three live neighbors survives.
3. Any live cell with more than three live neighbors dies (overpopulation).
4. Any dead cell with exactly three live neighbors becomes a live cell (reproduction).

## Features

- **Interactive Grid**: Users can click to toggle cell states (alive or dead).
- **Run Simulation**: Start and pause the simulation of the game.
- **Responsive and Dynamic**: The frontend is built using HTMX, making it snappy and lightweight, without requiring a full JavaScript framework.

## Tech Stack

- **Go**: Powers the backend logic for generating new states in the Game of Life.
- **HTMX**: Handles dynamic updates on the frontend without the need for heavy JavaScript frameworks. It allows for a simpler, more lightweight interaction between the client and server.
- **HTML/CSS**: Standard markup and styles to create the grid interface and control panel.

## Installation

1. Clone this repository:
    ```bash
    git clone https://github.com/citizen-thayne/cgw.git
    ```
2. Navigate to the project directory:
    ```bash
    cd cgw
    ```
3. Install dependencies and run the Go server:
    ```bash
    go run main.go
    ```
4. Open your browser and navigate to `http://localhost:8080` to view the application.

## How It Works

1. **Grid**: The grid is rendered dynamically on the backend before being rendered by the frontend.
2. **Simulation**: The game simulates the next generation of cells based on the rules of Conway's Game of Life. The simulation runs on the backend and every 2 seconds.
3. **HTMX Requests**: Each user action triggers an HTMX request to the Go backend, which processes the game logic and returns updated HTML snippets for the grid.
4. **Polling**: The frontend polls the server every 1 second to get the latest state of the grid. (TODO: Use HTMX SSE for real-time updates.)


