# Structs: Basics

In this exercise, you'll build a game store along with a set of functions to query the store.

1. Declare the following structs:
   - `item`: `id` (int), `name` (string), `price` (int)
   - `game`: embed the `item`, `genre` (string)

2. Write a `newGame` function that takes an item id, name and price and returns a game filled with
   the data supplied.

3. Stringify items and games: write a `String` function with an `item` receiver to dump an item in
   the form `<id>: <name> costs <price>` and another `String` function with a `game` receiver to
   dump a game in the form `Game <id>: <name> costs <price> of genre <genre>`.

4. Write a `newGameList` function that creates a game slice using the following data:

   | id | name | price | genre |
   |----|------|-------|-------|
   |    |      |       |       |
   | 1 | god of war | 50 | action adventure |
   | 3 | minecraft | 20 | sandbox |
   | 4 | warcraft | 40 | strategy |

5. Write a `queryById` function that returns the game in a gamelist with the given id or an "No
   such game" error.

6. Write a `listNameByPrice` function that returns the name of the game(s) with price equal or
   smaller than a given price.

Insert your code into the file `exercise.go`.




