# Pokédex
Read-Execute-Print-Loop (REPL) [Pokédex](https://bulbapedia.bulbagarden.net/wiki/Pok%C3%A9dex) Application with caching.

Features:
- Ability to catch and inspect Pokémon in the terminal
- Usage of PokéAPI to deliver up to date and full Pokémon information
- Caching of PokéAPI pages to optimize user experience
- CLI Input sanitation
- Unit Testing to ensure smooth development and maintenance

## Usage
### Prerequisites
Ensure you have Go v1.22+ installed on your system.

### Building and Running the Application
From the project's root directory, use the Go command-line tool to build the executable:<br>
```bash
go build -o pokedex
```

This command generates an executable named `pokedex`.

Execute the binary in your terminal to enter the Pokédex command line:

```bash
./pokedex
```

To then view the list of commands and their descriptions, use the `help` command:

```
pokedexCLI > help
```

Exit the CLI tool with the `exit` command.

## Unit Tests
![image](https://github.com/adamhu714/pokedex/assets/105497355/6727dd77-acf1-4e39-a1c3-773a21aef7f5)

*[Back To Top](#pokédex)* <br>
