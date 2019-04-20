	/*
	Author: code34 nicolas_boiteux@yahoo.fr
	Copyright (C) 2013-2019 Nicolas BOITEUX

	CLASS OO_ARMAWS
	
	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.
	
	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.
	
	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <http://www.gnu.org/licenses/>. 
	*/

	#include "oop.h"

	CLASS("OO_ARMAWS")
		PRIVATE VARIABLE("string","version");
		PRIVATE VARIABLE("string","url");

		PUBLIC FUNCTION("","constructor") {
			MEMBER("version", "0.2");
			MEMBER("url", "");
		};

		PUBLIC FUNCTION("", "getVersion") {
			private["_data"];
			_data = "armaws" callExtension "getVersion";
			_data = format["Armaws: %1 Dll: %2", MEMBER("version", nil), _data];
			_data;
		};

		PUBLIC FUNCTION("string", "setUrl") {
			MEMBER("url", _this);
		};

		PUBLIC FUNCTION("array", "callWs") {
			private _line = "callWs;"+MEMBER("url", nil);
			private _result = [];

			{
					{
						if(typeName _x == "STRING") then {
							_line = _line + ";" + _x; 
						} else {
							_line = _line + ";" + str(_x);
						};
					} foreach _x;
			}foreach _this;

			private _result = parseSimpleArray ("armaws" callExtension _line);
			if((_result select 0) isEqualTo -1) then {
				diag_log format ["ARMAWS: %1", _result select 1];
				throw [-1, format ['EXCEPTION : %1', _result select 1]];
				_result = [];
			} else {
				_result = _result select 1;
			};
			_result;
		};

		PUBLIC FUNCTION("","deconstructor") {
			DELETE_VARIABLE("version");
			DELETE_VARIABLE("url");
		};
	ENDCLASS;