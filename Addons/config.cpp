class CfgPatches {
	class armaws {
		units[] = {};
		weapons[] = {};
		requiredVersion = 0.1;
		requiredAddons[] = {};
		author[] = {"Code34"};
		authorUrl = "https://github.com/code34";
	};
};

class CfgFunctions
{
	class A3
	{
		class OO {
			class armaws {
				preInit = 1;
				file = "\armaws\oo_armaws.sqf";
			};
		};
	};
};