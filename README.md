# ARMAWS

Description: A JSON client dll for ARMA3

Author:  code34 nicolas_boiteux@yahoo.fr

Copyright (C) 2019 Nicolas BOITEUX - License MIT

# How to install:

1- Unpack the archive and copy the entire "@armaws_x64" folder into the ARMA3 root directory.

The @inibdi2 folder should look like this:

../Arma 3/@armaws/armaws_x64.dll

../Arma 3/@armaws/Addons/armaws.pbo

2- check armaws_x64.dll execution permissions, right click on it, and authorize it.

3- check in Arma3 launcher, that Battleye is turn off until BIS whitelist the dll

# Changelog

- version 0.2 : first official release

# Exemple

Send a JSON file to httpbin server and retrieve a json file convert into an arma array

	private _armaws = "new" call OO_ARMAWS;
	private _params = [["username","code34"],["message","hello world"],["id", 103],["type","soldier"]];
	["setUrl", "https://httpbin.org/post"] call _armaws;
	_result = ["callWs", _params] call _armaws;
	hintc format["%1",_result];