# Pokédex
Read-Execute-Print-Loop (REPL) [Pokédex](https://bulbapedia.bulbagarden.net/wiki/Pok%C3%A9dex) Application with caching.

Features:
- Ability to catch and inspect Pokémon in the terminal
- Usage of [PokéAPI](https://pokeapi.co/) to deliver up to date and full Pokémon information
- Caching of PokéAPI pages to optimize user experience
- CLI input sanitation
- Unit testing to ensure smooth and robust development and maintenance

## Contents

* [Getting Started](#getting-started)
  * [Prerequisites](#prerequisites)
  * [Building and Running The Application](#building-and-running-the-application)
* [Commands](#commands)
  * [catch <pokemon_name/id>](#catch-pokemon_nameid)
  * [inspect <pokemon_name/id>](#inspect-pokemon_nameid)
  * [pokedex](#pokedex)
  * [map](#map)
  * [mapb](#mapb)
  * [explore <location_area_name/id>](#explore-location_area_nameid)
  * [help](#help)
  * [exit](#exit)
* [Unit Tests](#unit-tests)

## Getting Started
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

The command line should show the following prompt:

```
pokedexCLI > 
```

Exit the CLI tool with the `exit` command or by pressing `CTRL + C`.

*[Back To Top](#pokédex)* <br>

## Commands

* [catch <pokemon_name/id>](#catch-pokemon_nameid)
* [inspect <pokemon_name/id>](#inspect-pokemon_nameid)
* [pokedex](#pokedex)
* [map](#map)
* [mapb](#mapb)
* [explore <location_area_name/id>](#explore-location_area_nameid)
* [help](#help)
* [exit](#exit)

### catch <pokemon_name/id>

Accepts 1 argument: <pokemon_name/id>

Attempt to catch a pokemon. Rolls a random integer between 1 and the pokémon's base experience. Stronger pokémon have higher base experience. If the roll is 50 or lower, the pokémon is caught and added to the pokédex.

Example usage:

``` bash
pokedexCLI > catch mew
```
``` bash
pokedexCLI > catch 151
```

*[Back To Top](#pokédex)* &nbsp; *[Back To Commands](#commands)*<br>

### inspect <pokemon_name/id>

Accepts 1 argument: <pokemon_name/id>

Prints information about a specified pokémon, if they have been added to the pokédex.

Example usage:

``` bash
pokedexCLI > inspect mew
```
``` bash
pokedexCLI > inspect 151
```

*[Back To Top](#pokédex)* &nbsp; *[Back To Commands](#commands)*<br>

### pokedex

Lists all the pokémon that have been added to the pokédex so far.

*[Back To Top](#pokédex)* &nbsp; *[Back To Commands](#commands)*<br>

### map

Lists all the location areas on the following page of location areas from the PokéAPI.

*[Back To Top](#pokédex)* &nbsp; *[Back To Commands](#commands)*<br>

### mapb

Lists all the location areas on the previous page of location areas from the PokéAPI.

*[Back To Top](#pokédex)* &nbsp; *[Back To Commands](#commands)*<br>

### explore <location_area_name/id>

Accepts 1 argument: <location_area_name/id>

Lists the pokemon in a specified location area.

Example usage:

``` bash
pokedexCLI > explore eterna-city-area
```
``` bash
pokedexCLI > explore 2
```

*[Back To Top](#pokédex)* &nbsp; *[Back To Commands](#commands)*<br>

### help

Lists all commands and their description. 

*[Back To Top](#pokédex)* &nbsp; *[Back To Commands](#commands)*<br>

### exit

Exits the CLI tool.

*[Back To Top](#pokédex)* &nbsp; *[Back To Commands](#commands)*<br>

## Unit Tests
![image](https://github.com/adamhu714/pokedex/assets/105497355/6727dd77-acf1-4e39-a1c3-773a21aef7f5)

*[Back To Top](#pokédex)* <br>
