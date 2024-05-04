-----------------------------------------------------
--                   DROP TABLES                   --
-----------------------------------------------------

DROP VIEW IF EXISTS RowCounter;

DROP TABLE IF EXISTS Anime_Studioses;
DROP TABLE IF EXISTS Anime_Characters;
DROP TABLE IF EXISTS Anime_Genres;
DROP TABLE IF EXISTS Anime_Themes;
DROP TABLE IF EXISTS Anime_Images;
DROP TABLE IF EXISTS Anime_Descriptions;
DROP TABLE IF EXISTS Episode_Descriptions;
DROP TABLE IF EXISTS Character_Images;
DROP TABLE IF EXISTS Character_Descriptions;

DROP TABLE IF EXISTS Statuses;
DROP TABLE IF EXISTS Types;
DROP TABLE IF EXISTS Studios;
DROP TABLE IF EXISTS Themes;
DROP TABLE IF EXISTS Genres;
DROP TABLE IF EXISTS Seasons;
DROP TABLE IF EXISTS Images;
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
	Url VARCHAR(2048) NOT NULL UNIQUE
);

CREATE TABLE Genres (
	Id CHAR(110) PRIMARY KEY,
	Name VARCHAR(2048) NOT NULL UNIQUE
);

CREATE TABLE Themes (
	Id CHAR(110) PRIMARY KEY,
	Name VARCHAR(2048) NOT NULL UNIQUE
);

-----------------------------------------------------
--                   MAIN TABLES                   --
-----------------------------------------------------

CREATE TABLE Anime (
	Id CHAR(110) PRIMARY KEY,
	Title VARCHAR(2048),
	AlternativeTitle VARCHAR(2048),
	Aired INTEGER DEFAULT null,
	Duration INTEGER DEFAULT NULL,
	Broadcast INTEGER DEFAULT NULL,
	Url VARCHAR(2048) NOT NULL,

	Season_ID CHAR(110) DEFAULT NULL,
	Type_ID CHAR(110) DEFAULT NULL,
	CurrentStatus INTEGER DEFAULT NULL,

	FOREIGN KEY (season_ID) REFERENCES Seasons(Id) ON UPDATE CASCADE ON DELETE SET NULL,
	FOREIGN KEY (type_ID) REFERENCES Types(Id) ON UPDATE CASCADE ON DELETE SET NULL,
	FOREIGN KEY (currentStatus) REFERENCES Statuses(Id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE TABLE Episode (
	Id CHAR(110) PRIMARY KEY,
	Number FLOAT NOT NULL CHECK(number >= 0),
	Title VARCHAR(2048) DEFAULT NULL,
	Aired INTEGER DEFAULT NULL,
	Duration INTEGER DEFAULT NULL,

	Anime_ID CHAR(110) NOT NULL,
	FOREIGN KEY (Anime_ID) REFERENCES Anime(Id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE Character (
	Id CHAR(110) PRIMARY KEY,
	Name VARCHAR(2048) NOT NULL,
	Url VARCHAR(2048) NOT NULL
);

CREATE TABLE Descriptions (
	Id CHAR(110) PRIMARY KEY,
	Description VARCHAR(2048) NOT NULL,
	Anime_ID VARCHAR(110) DEFAULT NULL REFERENCES Anime(Id) ON UPDATE CASCADE ON DELETE SET NULL,
	Episode_ID VARCHAR(110) DEFAULT NULL REFERENCES Episodes(Id) ON UPDATE CASCADE ON DELETE SET NULL,
	Character_ID VARCHAR(110) DEFAULT NULL REFERENCES Character(Id) ON UPDATE CASCADE ON DELETE SET NULL
);

----------------------------------------------------
--                 LINKING TABLES                 --
----------------------------------------------------

CREATE TABLE Anime_Studioses (
	Anime_ID CHAR(110) NOT NULL,
	Studio_ID CHAR(110) NOT NULL,
	FOREIGN KEY (Anime_ID) REFERENCES Anime(Id) ON UPDATE CASCADE ON DELETE CASCADE,
	FOREIGN KEY (Studio_ID) REFERENCES Studios(Id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE Anime_Characters (
	Anime_ID CHAR(110) NOT NULL,
	Character_ID CHAR(110) NOT NULL,
	FOREIGN KEY (Anime_ID) REFERENCES Anime(Id) ON UPDATE CASCADE ON DELETE CASCADE,
	FOREIGN KEY (Character_ID) REFERENCES Character(Id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE Anime_Genres (
	Anime_ID CHAR(110) NOT NULL,
	Genre_ID CHAR(110) NOT NULL,
	FOREIGN KEY (Anime_ID) REFERENCES Anime(Id) ON UPDATE CASCADE ON DELETE CASCADE,
	FOREIGN KEY (Genre_ID) REFERENCES Genres(Id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE Anime_Themes (
	Anime_ID CHAR(110) NOT NULL,
	Theme_ID CHAR(110) NOT NULL,
	FOREIGN KEY (Anime_ID) REFERENCES Anime(Id) ON UPDATE CASCADE ON DELETE CASCADE,
	FOREIGN KEY (Theme_ID) REFERENCES Themes(Id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE Anime_Images (
	Anime_ID CHAR(110) NOT NULL,
	Image_ID CHAR(110) NOT NULL,
	IsDefault INTEGER DEFAULT 0,
	FOREIGN KEY (Anime_ID) REFERENCES Anime(Id) ON UPDATE CASCADE ON DELETE CASCADE,
	FOREIGN KEY (Image_ID) REFERENCES Images(Id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE Character_Images (
	Character_ID CHAR(110) NOT NULL,
	Image_ID CHAR(110) NOT NULL,
	IsDefault INTEGER DEFAULT 0,
	FOREIGN KEY (Character_ID) REFERENCES Character(Id) ON UPDATE CASCADE ON DELETE CASCADE,
	FOREIGN KEY (Image_ID) REFERENCES Images(Id) ON UPDATE CASCADE ON DELETE CASCADE
);


-----------------------------------------------------
--                      VIEWS                      --
-----------------------------------------------------

CREATE VIEW RowCounter AS 
SELECT 'Anime_Characters' AS "Table", COUNT(*) AS "Count" FROM Anime_Characters UNION
SELECT 'Anime_Genres', COUNT(*) FROM Anime_Genres UNION
SELECT 'Anime_Images', COUNT(*) FROM Anime_Images UNION
SELECT 'Anime_Studioses', COUNT(*) FROM Anime_Studioses UNION
SELECT 'Anime_Themes', COUNT(*) FROM Anime_Themes UNION
SELECT 'Anime', COUNT(*) FROM Anime UNION
SELECT 'Character_Images', COUNT(*) FROM Character_Images UNION
SELECT 'Character', COUNT(*) FROM Character UNION
SELECT 'Descriptions', COUNT(*) FROM Descriptions UNION
SELECT 'Episode', COUNT(*) FROM Episode UNION
SELECT 'Genres', COUNT(*) FROM Genres UNION
SELECT 'Images', COUNT(*) FROM Images UNION
SELECT 'Seasons', COUNT(*) FROM Seasons UNION
SELECT 'Statuses', COUNT(*) FROM Statuses UNION
SELECT 'Studios', COUNT(*) FROM Studios UNION
SELECT 'Themes', COUNT(*) FROM Themes UNION
SELECT 'Types', COUNT(*) FROM Types
ORDER BY "Table" ASC;


-----------------------------------------------------
--                     INSERT                      --
-----------------------------------------------------

INSERT INTO Statuses (Id, Name) VALUES
(0, 'Watched'),
(1, 'Airing'),
(2, 'Started'),
(3, 'Plan to watch'),
(4, 'New');