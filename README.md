# Ether Furnace - Advanced
An advanced version of the wismt texture tool from 3096

## License

    Ether Furnace - Advanced
    Copyright (C) 2022  3096

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License along
    with this program. If not, see <https://www.gnu.org/licenses/>.

## Usage

You must format your texture file names with <u><id.name.dds></u> (e.g. `00.PC079404_WAIST.dds`). The id will be used to identify the texture it replaces.\

Along side the `wismt` file, you also need the `wimdo` file placed in the same directory. Both files need to be modified for the replaced textures to function correctly in game.

    go run main.go <repalcement type>

### Single Item

The desired `.wismt` and `.wimdo` files for the item you want to edit should be in the `Input` folder.

The desired changes should be placed in the `Replace` folder.

    go run main.go

### Multiple Items

The desired `.wismt` and `.wimdo` files for the item you want to edit should be in a folder sharing their name (without a file extension) inside the `Input` folder.

E.g. `Input/bl000101`

The desired changes should be placed in a folder sharing the name of the files they will be inputted into (without a file extension) inside the `Replace` folder.

E.g. `Replace/bl000101`

    go run main.go mi

### Multiple Versions

The desired `.wismt` and `.wimdo` files for the item you want to edit should be in the `Input` folder.

The desired changes should be put in a folder for each respective version, starting at `1` and going in assending order, inside the `Replace` folder. 

E.g. `Replace/1`

    go run main.go mv <amount of versions>

### Multiple Versions of Multiple Items

The desired `.wismt` and `.wimdo` files for the item you want to edit should be in a folder sharing their name (without a file extension) inside the `Input` folder.

E.g. `Input/bl000101`

The desired changes should be placed inside the `Replace` folder.

E.g. `Replace/bl000101`

The desired changes should be put  in a folder sharing the name of the files they will be inputted into (without a file extension), inside a folder for each respective version, starting at `1` and going in assending order, inside the `Replace` folder. 

E.g. `Replace/1/bl000101`

    go run main.go mvmi <amount of versions>
    
or

    go run main.go mimv <amount of versions>
