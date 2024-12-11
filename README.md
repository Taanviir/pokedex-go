# Pokedex CLI Application

The Pokedex CLI Application is an interactive command-line tool for exploring the world of Pokemon. It allows users to view locations, explore areas, catch Pokemon, inspect caught Pokemon, and manage their personal Pokedex.

## Features

- **Help Menu**: Access a list of available commands and their descriptions.
- **Exit**: Exit the application gracefully.
- **Explore Locations**: Discover various locations in the Pokemon world.
- **View Pokemon in Areas**: Explore specific areas to find roaming Pokemon.
- **Catch Pokemon**: Attempt to catch a Pokemon and add it to your Pokedex.
- **Inspect Pokemon**: View detailed information about Pokemon you’ve caught.
- **View Pokedex**: Check your collection of caught Pokemon.

## Commands

| Command              | Description                                          |
|----------------------|------------------------------------------------------|
| `help`              | Displays the help menu with a list of commands.     |
| `exit`              | Exits the application.                              |
| `map`               | Lists locations in the Pokemon world.               |
| `mapb`              | Lists the previous page of location areas.          |
| `explore <location>` | Lists Pokemon roaming in the specified area.        |
| `catch <pokemon>`   | Attempts to catch the specified Pokemon.            |
| `inspect <pokemon>` | Displays details of a caught Pokemon.               |
| `pokedex`           | Shows your collection of caught Pokemon.            |

## Usage

### Starting the Application
Run the application in your terminal. Once launched, you will see a `>` prompt where you can enter commands.

### Examples

- View the help menu:
  ```
  > help
  ```
- List locations in the Pokemon world:
  ```
  > map
  ```
- Explore an area:
  ```
  > explore pallet-town
  ```
- Catch a Pokemon:
  ```
  > catch pikachu
  ```
- Inspect a caught Pokemon:
  ```
  > inspect pikachu
  ```
- View your Pokedex:
  ```
  > pokedex
  ```
- Exit the application:
  ```
  > exit
  ```

## Development

### Structure

- **CLI Commands**: Defined in the `cliCommand` struct.
- **Input Parsing**: Inputs are cleaned and tokenized using `cleanInput`.
- **REPL Loop**: Runs in the `startRepl` function to process user inputs continuously.

### Adding New Commands
To add a new command, follow these steps:
1. Create a new function with the signature `func(*config, ...string) error`.
2. Add the command to the `getCommands` function with its name, description, and callback.

Example:
```go
"newcommand": {
    name:        "newcommand",
    description: "Description of the new command.",
    callback:    newCommandFunction,
},
```

### Dependencies
- **Go Modules**: Ensure Go modules are enabled to manage dependencies.

## Contributing
Feel free to fork the repository, add features, and submit pull requests. Ensure that your code follows best practices and is well-documented.

## License
This project is open source and available under the [MIT License](LICENSE).

