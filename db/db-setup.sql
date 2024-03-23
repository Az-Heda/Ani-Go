-----------------------------------------------------
--                   DROP TABLES                   --
-----------------------------------------------------
DROP TABLE IF EXISTS Anime_Studioses;
DROP TABLE IF EXISTS Anime_Characters;
DROP TABLE IF EXISTS Anime_Genres;
DROP TABLE IF EXISTS Anime_Images;
DROP TABLE IF EXISTS Anime_Descriptions;
DROP TABLE IF EXISTS Episode_Descriptions;
DROP TABLE IF EXISTS Character_Images;
DROP TABLE IF EXISTS Character_Descriptions;

DROP TABLE IF EXISTS Statuses;
DROP TABLE IF EXISTS Types;
DROP TABLE IF EXISTS Studios;
DROP TABLE IF EXISTS Seasons;
DROP TABLE IF EXISTS Images;
DROP TABLE IF EXISTS Genres;
DROP TABLE IF EXISTS Descriptions;

DROP TABLE IF EXISTS Character;
DROP TABLE IF EXISTS Episode;
DROP TABLE IF EXISTS Anime;

-----------------------------------------------------
--                   STAND ALONE                   --
-----------------------------------------------------

CREATE TABLE Statuses (
	Id INTEGER PRIMARY KEY AUTOINCREMENT,
	Name VARCHAR(2048) NOT NULL UNIQUE
);

CREATE TABLE Types (
	Id CHAR(110) PRIMARY KEY,
	Name VARCHAR(2048) NOT NULL UNIQUE
);

CREATE TABLE Studios (
	Id CHAR(110) PRIMARY KEY,
	Name VARCHAR(2048) NOT NULL UNIQUE
);

CREATE TABLE Seasons (
	Id CHAR(110) PRIMARY KEY,
	Season VARCHAR(2048) NOT NULL UNIQUE
);

CREATE TABLE Images (
	Id CHAR(110) PRIMARY KEY,
	Url VARCHAR(2048) NOT NULL
);

CREATE TABLE Genres (
	Id CHAR(110) PRIMARY KEY,
	Name VARCHAR(2048) NOT NULL UNIQUE
);

CREATE TABLE Descriptions (
	Id CHAR(110) PRIMARY KEY,
	Description VARCHAR(2048) NOT NULL UNIQUE
);

-----------------------------------------------------
--                   MAIN TABLES                   --
-----------------------------------------------------

CREATE TABLE Anime (
	Id CHAR(110) PRIMARY KEY,
	Title VARCHAR(2048),
	AlternativeTitle VARCHAR(2048),
	Aired INTEGER DEFAULT null,
	Duration NUMBER DEFAULT NULL,
	Url VARCHAR(2048) NOT NULL,

	Season_ID CHAR(110) DEFAULT NULL,
	Type_ID CHAR(110) DEFAULT NULL,
	CurrentStatus NUMBER DEFAULT NULL,

	FOREIGN KEY (season_ID) REFERENCES Seasons(Id),
	FOREIGN KEY (type_ID) REFERENCES Types(Id),
	FOREIGN KEY (currentStatus) REFERENCES Statuses(Id)
);

CREATE TABLE Episode (
	Id CHAR(110) PRIMARY KEY,
	Number FLOAT NOT NULL CHECK(number >= 0),
	Title VARCHAR(2048) DEFAULT NULL,
	Aired INTEGER DEFAULT NULL,
	Duration INTEGER DEFAULT NULL,

	Anime_ID CHAR(110) NOT NULL,
	FOREIGN KEY (Anime_ID) REFERENCES Anime(Id)
);

CREATE TABLE Character (
	Id CHAR(110) PRIMARY KEY,
	Name VARCHAR(2048) NOT NULL
);

----------------------------------------------------
--                 LINKING TABLES                 --
----------------------------------------------------

CREATE TABLE Anime_Studioses (
	Anime_ID CHAR(110) NOT NULL,
	Studio_ID CHAR(110) NOT NULL,
	FOREIGN KEY (Anime_ID) REFERENCES Anime(Id),
	FOREIGN KEY (Studio_ID) REFERENCES Studios(Id)
);

CREATE TABLE Anime_Characters (
	Anime_ID CHAR(110) NOT NULL,
	Character_ID CHAR(110) NOT NULL,
	FOREIGN KEY (Anime_ID) REFERENCES Anime(Id),
	FOREIGN KEY (Character_ID) REFERENCES Character(Id)
);

CREATE TABLE Anime_Genres (
	Anime_ID CHAR(110) NOT NULL,
	Genre_ID CHAR(110) NOT NULL,
	FOREIGN KEY (Anime_ID) REFERENCES Anime(Id),
	FOREIGN KEY (Genre_ID) REFERENCES Genres(Id)
);

CREATE TABLE Anime_Images (
	Anime_ID CHAR(110) NOT NULL,
	Image_ID CHAR(110) NOT NULL,
	FOREIGN KEY (Anime_ID) REFERENCES Anime(Id),
	FOREIGN KEY (Image_ID) REFERENCES Images(Id)
);

CREATE TABLE Anime_Descriptions (
	Anime_ID CHAR(110) NOT NULL,
	Description_ID CHAR(110) NOT NULL,
	FOREIGN KEY (Anime_ID) REFERENCES Anime(Id),
	FOREIGN KEY (Description_ID) REFERENCES Descriptions(Id)
);

CREATE TABLE Episode_Descriptions (
	Episode_ID CHAR(110) NOT NULL,
	Description_ID CHAR(110) NOT NULL,
	FOREIGN KEY (Episode_ID) REFERENCES Episode(Id),
	FOREIGN KEY (Description_ID) REFERENCES Descriptions(Id)
);

CREATE TABLE Character_Images (
	Character_ID CHAR(110) NOT NULL,
	Image_ID CHAR(110) NOT NULL,
	FOREIGN KEY (Character_ID) REFERENCES Character(Id),
	FOREIGN KEY (Image_ID) REFERENCES Images(Id)
);

CREATE TABLE Character_Descriptions (
	Character_ID CHAR(110) NOT NULL,
	Description_ID CHAR(110) NOT NULL,
	FOREIGN KEY (Character_ID) REFERENCES Character(Id),
	FOREIGN KEY (Description_ID) REFERENCES Descriptions(Id)
);



-----------------------------------------------------
--                      VIEWS                      --
-----------------------------------------------------

CREATE VIEW RowCounter AS 
SELECT 'Anime_Studioses' AS "Table", COUNT(*) AS "Count" FROM Anime_Studioses UNION
SELECT 'Anime_Characters', COUNT(*) FROM Anime_Characters UNION
SELECT 'Anime_Genres', COUNT(*) FROM Anime_Genres UNION
SELECT 'Anime_Images', COUNT(*) FROM Anime_Images UNION
SELECT 'Anime_Descriptions', COUNT(*) FROM Anime_Descriptions UNION
SELECT 'Episode_Descriptions', COUNT(*) FROM Episode_Descriptions UNION
SELECT 'Character_Images', COUNT(*) FROM Character_Images UNION
SELECT 'Character_Descriptions', COUNT(*) FROM Character_Descriptions UNION
SELECT 'Statuses', COUNT(*) FROM Statuses UNION
SELECT 'Types', COUNT(*) FROM Types UNION
SELECT 'Studios', COUNT(*) FROM Studios UNION
SELECT 'Seasons', COUNT(*) FROM Seasons UNION
SELECT 'Images', COUNT(*) FROM Images UNION
SELECT 'Genres', COUNT(*) FROM Genres UNION
SELECT 'Descriptions', COUNT(*) FROM Descriptions UNION
SELECT 'Character', COUNT(*) FROM Character UNION
SELECT 'Episode', COUNT(*) FROM Episode UNION
SELECT 'Anime', COUNT(*) FROM Anime
ORDER BY "Table" ASC;