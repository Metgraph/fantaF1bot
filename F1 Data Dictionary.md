# F1 Data Dictionary

## User
An user of telegram registered to FF1.
Attributes:
* id: Telegram identifier for the user. It is possible to retrieve Telegram Username thanks to this id.

## Fanta
An FF1 Tournament.
Attributes:
* id: Identifier for the fanta.
* name: The name of the fanta.
* owner: The id of the user that creates and owns the fanta.

## Fanta_User
A stable that participate to a Fanta.
Attributes:
* name: The name of the Stable.
* user_id: The user owner of the Stable.
* fanta_id: The fanta to which the Stable is associated.

## Race
A one season Race,
Attributes:
* day: The day of the race.
* month: The month of the race.
* year: The year of the race.
* name: The name of the race.
* round: The Ordinal number of the race for the race's year.

## Race_Prediction_Result
The FF1 result for a one season Race. 
Attributes:
* race_year: The year of the race.
* race_round: The Ordinal number of the race.
* vsc: Flag for virtual safety car.
* sc: Flag for safety car.
* red_flag: Flag for red flag.
* dfs: Count of dfs in the race.
* double: Flag for team brace.
* poleman_win: Flag for winning poleman
* wet: Flag for wet tyres changing.
* double_point: **TODO**

## Team
A Team of 5 drivers for a one season Race of one Fanta_User.
Attributes:
* id: The identifier of the team.
* fanta_id: The id of the fanta_user that created the team.
* race_year: The year of the race.
* race_round: The Ordinal number of the race.

## Team_Driver
A driver associated to a Team.
Attributes:
* team_id: The identifier of the team.
* driver_id: The identifier of the season_driver.
* captain: Flag indicating if this team_driver is the captain of the team.

## Season_Driver
A driver with a 1 year contract.
Attributes:
* year: The year in which the driver race.
* driver_id: The id of the driver associated to the season.
* constructor_id: The id of the constructor associated to the driver in the season.
* number: The number of the driver in the season.

## Price_Driver
The price of a Season_Driver.
Attributes:
* year: The year in which the driver race.
* driver_id: The id of the driver associated to the season.
* price: The value of the driver in the season.

## Driver
A driver of F1.
Attributes:
* id: The identifier of the driver.
* name: The name of the driver.

## Constructor
A constructor of F1.
Attributes:
* id: The identifier of the constructor.
* name: The name of the constructor.




















