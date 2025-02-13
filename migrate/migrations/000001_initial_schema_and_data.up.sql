-- MySQL dump 10.13  Distrib 8.0.39, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: reelplay
-- ------------------------------------------------------
-- Server version	8.0.39


--
-- Table structure for table `actors`
--

SET FOREIGN_KEY_CHECKS=0;

DROP TABLE IF EXISTS `actors`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `actors` (
                          `id` bigint NOT NULL AUTO_INCREMENT,
                          `year` bigint DEFAULT NULL,
                          `name` longtext,
                          PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=77 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `actors`
--

INSERT INTO `actors` VALUES (14,1958,'Tim Robbins'),(15,1937,'Morgan Freeman'),(16,1945,'Bob Gunton'),(17,1924,'Marlon Brando'),(18,1940,'Al Pacino'),(19,1974,'Christian Bale'),(20,1979,'Heath Ledger'),(21,1968,'Aaron Eckhart'),(22,1905,'Henry Fonda'),(23,1911,'Lee J. Cobb'),(24,1919,'Martin Balsam'),(25,1981,'Elijah Wood'),(26,1958,'Viggo Mortensen'),(27,1939,'Ian McKellen'),(28,1977,'Orlando Bloom'),(29,1966,'Robin Wright'),(30,1955,'Gary Sinise'),(31,1956,'Tom Hanks'),(32,1982,'Joshua Walters'),(33,1966,'Ryôhei Abe'),(34,1987,'Leon Shibli Ahmad'),(35,1996,'Kodi Smit-McPhee'),(36,1980,'Jóhannes Haukur Jóhannesson'),(37,1987,'Marcin Kowalczyk'),(38,1967,'Emily Watson'),(39,1964,'David Morrissey'),(40,1994,'Alex Etel'),(41,1991,'Zazie Beetz'),(42,1979,'Claire Danes'),(43,1966,'Jim Gaffigan'),(44,1974,'Leonardo DiCaprio'),(45,1981,'Joseph Gordon-Levitt'),(46,1987,'Elliot Page'),(47,1963,'Brad Pitt'),(48,1954,'Denzel Washington'),(49,1963,'Johnny Depp'),(50,1965,'Robert Downey Jr.'),(51,1970,'Matt Damon'),(52,1968,'Will Smith'),(53,1964,'Russell Crowe'),(54,1962,'Tom Cruise'),(55,1968,'Hugh Jackman'),(56,1957,'Daniel Day-Lewis'),(57,1937,'Anthony Hopkins'),(58,1948,'Samuel L. Jackson'),(59,1943,'Robert De Niro'),(60,1937,'Jack Nicholson'),(61,1974,'Joaquin Phoenix'),(62,1961,'George Clooney'),(63,1980,'Ryan Gosling'),(64,1983,'Chris Hemsworth'),(65,1976,'Benedict Cumberbatch'),(66,1982,'Eddie Redmayne'),(67,1977,'Michael Fassbender'),(68,1977,'Tom Hardy'),(69,1980,'Jake Gyllenhaal'),(70,1960,'Colin Firth'),(71,1968,'Daniel Craig'),(72,1972,'Jude Law'),(73,1969,'Edward Norton'),(74,1964,'Keanu Reeves'),(75,1969,'Matthew McConaughey'),(76,1974,'Christian Bale');

--
-- Table structure for table `bookmarks`
--

DROP TABLE IF EXISTS `bookmarks`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `bookmarks` (
                             `user_id` bigint NOT NULL,
                             `movie_id` bigint NOT NULL,
                             `created_at` datetime(3) DEFAULT NULL,
                             PRIMARY KEY (`user_id`,`movie_id`),
                             KEY `fk_bookmarks_movie` (`movie_id`),
                             CONSTRAINT `fk_bookmarks_movie` FOREIGN KEY (`movie_id`) REFERENCES `movies` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
                             CONSTRAINT `fk_bookmarks_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `bookmarks`
--

INSERT INTO `bookmarks` VALUES (1,35,'2024-11-11 10:36:52.646'),(1,37,'2025-02-11 20:38:03.977'),(7,1,'2024-10-29 15:22:51.069'),(7,3,'2024-10-29 15:24:00.296'),(7,4,'2024-11-03 10:05:19.863'),(7,37,'2024-11-14 08:30:16.479'),(12,1,'2024-11-13 13:55:05.058'),(13,42,'2024-12-09 09:38:43.407');

--
-- Table structure for table `categories`
--

DROP TABLE IF EXISTS `categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `categories` (
                              `id` bigint NOT NULL AUTO_INCREMENT,
                              `name` longtext,
                              `description` longtext,
                              PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `categories`
--

INSERT INTO `categories` VALUES (1,'Action',NULL),(2,'Adventure',NULL),(3,'Animation',NULL),(4,'Biography',NULL),(5,'Comedy',NULL),(6,'Crime',NULL),(7,'Drama',NULL),(8,'Fantasy',NULL),(9,'History',NULL),(10,'Horror',NULL),(11,'Music',NULL),(12,'Mystery',NULL),(13,'Romance',NULL),(14,'Sci-Fi','Scince Fiction'),(15,'Sport',NULL),(16,'Thriller',NULL),(17,'War',NULL);

--
-- Table structure for table `category_fits`
--

DROP TABLE IF EXISTS `category_fits`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `category_fits` (
                                 `user_id` bigint NOT NULL,
                                 `category_id` bigint NOT NULL,
                                 `fit_rate` float NOT NULL,
                                 `created_at` datetime(3) DEFAULT NULL,
                                 `updated_at` datetime(3) DEFAULT NULL,
                                 PRIMARY KEY (`user_id`,`category_id`),
                                 KEY `fk_category_fits_category` (`category_id`),
                                 CONSTRAINT `fk_category_fits_category` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
                                 CONSTRAINT `fk_category_fits_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `category_fits`
--

INSERT INTO `category_fits` VALUES (7,1,8.625,'2024-10-29 15:23:21.871','2024-11-07 15:57:09.763'),(7,2,8.5,'2024-10-29 15:23:43.780','2024-11-07 15:57:09.770'),(7,4,9,'2024-10-28 20:07:27.458','2024-11-07 15:52:28.029'),(7,5,9.5625,'2024-10-28 20:07:27.483','2024-11-07 15:57:09.773'),(7,7,7,'2024-11-03 10:24:59.199','2024-11-03 10:25:03.360'),(7,9,7,'2024-10-29 15:23:43.784','2024-10-29 15:23:49.209'),(13,1,9,'2024-12-09 11:46:08.895','2024-12-09 11:46:19.882'),(13,2,8.5,'2024-12-09 09:52:09.528','2024-12-09 11:46:27.362'),(13,5,8,'2024-12-09 09:52:53.260','2024-12-09 11:46:27.367'),(13,6,9,'2024-12-09 09:52:09.533','2024-12-09 09:52:09.533'),(13,7,8,'2024-12-09 09:39:10.600','2024-12-09 09:39:10.600'),(13,8,7,'2024-12-09 09:38:34.186','2024-12-09 09:53:16.998'),(13,9,9,'2024-12-09 11:46:08.919','2024-12-09 11:46:08.919'),(13,17,8,'2024-12-09 09:39:10.615','2024-12-09 09:39:10.615');

--
-- Table structure for table `comments`
--

DROP TABLE IF EXISTS `comments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `comments` (
                            `id` bigint NOT NULL AUTO_INCREMENT,
                            `content` longtext NOT NULL,
                            `user_id` bigint NOT NULL,
                            `movie_id` bigint NOT NULL,
                            `created_at` datetime(3) DEFAULT NULL,
                            `deleted_at` datetime(3) DEFAULT NULL,
                            PRIMARY KEY (`id`),
                            KEY `idx_comments_deleted_at` (`deleted_at`),
                            KEY `fk_comments_user` (`user_id`),
                            KEY `fk_comments_movie` (`movie_id`),
                            CONSTRAINT `fk_comments_movie` FOREIGN KEY (`movie_id`) REFERENCES `movies` (`id`),
                            CONSTRAINT `fk_comments_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comments`
--

INSERT INTO `comments` VALUES (15,'amazing movie',7,1,'2024-10-29 14:56:09.475',NULL),(16,'hay qué',7,1,'2024-10-29 14:56:42.071',NULL),(17,'love u, batman',7,3,'2024-10-29 15:23:10.614',NULL),(18,'good movie!',7,14,'2024-10-29 15:24:46.973',NULL),(19,'Legend',7,4,'2024-11-03 10:05:14.781',NULL),(20,'L',7,35,'2024-11-03 10:25:16.252',NULL),(21,'dfdf',7,38,'2024-11-07 15:53:21.776',NULL),(22,'new c',7,42,'2024-11-08 10:29:54.895',NULL),(23,'123',7,42,'2024-11-08 10:29:56.577',NULL),(24,'234234',7,42,'2024-11-08 10:29:58.765',NULL),(25,'sag312t',7,42,'2024-11-08 10:30:00.829',NULL),(26,'dsg31trv',7,42,'2024-11-08 10:30:03.709',NULL),(28,'acc gg',1,35,'2024-11-11 10:36:47.606',NULL),(29,'goat',7,2,'2024-11-14 08:30:04.833',NULL),(30,'cmt 16/11\n',7,37,'2024-11-16 21:35:37.354',NULL),(31,'good mv',13,43,'2024-12-09 09:39:15.842',NULL),(32,'good movie',1,6,'2025-02-11 20:37:08.138',NULL),(33,'phim hay ko mn\n',1,37,'2025-02-11 20:37:44.062',NULL);

--
-- Table structure for table `countries`
--

DROP TABLE IF EXISTS `countries`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `countries` (
                             `id` bigint NOT NULL AUTO_INCREMENT,
                             `name` longtext,
                             PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `countries`
--

INSERT INTO `countries` VALUES (1,'USA'),(2,'United Kingdom'),(3,'South Korea'),(4,'China'),(5,'Japan'),(6,'India'),(7,'France'),(8,'Germany'),(9,'Turkey'),(10,'Spain'),(11,'Canada'),(12,'Australia'),(13,'Russia'),(14,'Brazil'),(15,'Italy'),(16,'Philippines'),(17,'Netherlands'),(18,'Thailand'),(19,'Hong Kong'),(20,'Mexico'),(21,'Denmark'),(22,'Portugal'),(23,'Poland'),(24,'Norway'),(25,'Taiwan'),(26,'Ukraine'),(27,'Other country');

--
-- Table structure for table `directors`
--

DROP TABLE IF EXISTS `directors`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `directors` (
                             `id` bigint NOT NULL AUTO_INCREMENT,
                             `year` bigint DEFAULT NULL,
                             `name` longtext,
                             PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=859 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `directors`
--

INSERT INTO `directors` VALUES (1,1889,'Victor Fleming'),(2,1894,'John Ford'),(3,1895,'Lewis Milestone'),(4,1900,'David Hand'),(5,1901,'Clyde Geronimi'),(6,1902,'William Wyler'),(7,1906,'John Huston'),(8,1908,'Edward Dmytryk'),(9,1908,'Leo McCarey'),(10,1908,'Robert Rossen'),(11,1910,'John Sturges'),(12,1912,'Gene Kelly'),(13,1913,'Stanley Kramer'),(14,1914,'Cy Endfield'),(15,1914,'Robert Wise'),(16,1915,'Orson Welles'),(17,1918,'Robert Aldrich'),(18,1920,'Franklin J. Schaffner'),(19,1921,'George Roy Hill'),(20,1923,'Irvin Kershner'),(21,1924,'Sidney Lumet'),(22,1925,'Robert Altman'),(23,1925,'Robert Mulligan'),(24,1925,'Sam Peckinpah'),(25,1926,'Mel Brooks'),(26,1926,'Robert L. Rosen'),(27,1928,'Alan J. Pakula'),(28,1928,'Stanley Kubrick'),(29,1929,'David H. DePatie'),(30,1930,'Clint Eastwood'),(31,1931,'Hal Needham'),(32,1931,'Mike Nichols'),(33,1931,'William A. Graham'),(34,1932,'Barry Levinson'),(35,1932,'Richard Lester'),(36,1934,'Garry Marshall'),(37,1935,'Brian G. Hutton'),(38,1935,'John G. Avildsen'),(39,1935,'William Friedkin'),(40,1935,'Woody Allen'),(41,1936,'Robert Redford'),(42,1938,'Ralph Bakshi'),(43,1939,'Francis Ford Coppola'),(44,1939,'Joel Schumacher'),(45,1939,'John Badham'),(46,1939,'Michael Cimino'),(47,1939,'Wes Craven'),(48,1940,'Brian De Palma'),(49,1940,'James L. Brooks'),(50,1940,'Jay Sandrich'),(51,1940,'Ronald F. Maxwell'),(52,1940,'Terry Gilliam'),(53,1942,'Martin Scorsese'),(54,1943,'Emile Ardolino'),(55,1943,'John Langley'),(56,1943,'Michael Mann'),(57,1943,'Robert De Niro'),(58,1943,'Terrence Malick'),(59,1944,'George Lucas'),(60,1944,'Gregory Hoblit'),(61,1944,'Harold Ramis'),(62,1944,'Jim Abrahams'),(63,1944,'Jonathan Demme'),(64,1944,'Taylor Hackford'),(65,1945,'Curtis Hanson'),(66,1945,'David S. Ward'),(67,1945,'Joe Pytka'),(68,1945,'Penelope Spheeris'),(69,1945,'Ron Shelton'),(70,1945,'Stan Lathan'),(71,1945,'Steve Sohmer'),(72,1946,'David Anspaugh'),(73,1946,'David Lynch'),(74,1946,'Dennis Dugan'),(75,1946,'Joe Dante'),(76,1946,'Oliver Stone'),(77,1946,'Randal Kleiser'),(78,1946,'Richard Pepin'),(79,1946,'Steven Spielberg'),(80,1946,'Sylvester Stallone'),(81,1947,'David Zucker'),(82,1947,'Rob Reiner'),(83,1948,'John Carpenter'),(84,1948,'John Pasquin'),(85,1949,'David R. Ellis'),(86,1949,'James Widdoes'),(87,1949,'Mark DiSalle'),(88,1949,'Nancy Meyers'),(89,1949,'Randall Wallace'),(90,1949,'Rob Cohen'),(91,1949,'Roger Allers'),(92,1950,'Howard Deutch'),(93,1950,'Jerry Zucker'),(94,1950,'Joe Johnston'),(95,1950,'John Fortenberry'),(96,1950,'John Hughes'),(97,1950,'Joseph Ruben'),(98,1950,'Tom Fontana'),(99,1951,'John McTiernan'),(100,1951,'Kathryn Bigelow'),(101,1952,'Edward Zwick'),(102,1952,'Gil Junger'),(103,1952,'Gus Van Sant'),(104,1952,'Henry Selick'),(105,1952,'John Dahl'),(106,1952,'Julie Taymor'),(107,1952,'Robert Lorenz'),(108,1952,'Robert Zemeckis'),(109,1953,'Barry Sonnenfeld'),(110,1953,'Jim Jarmusch'),(111,1953,'Marc Lawrence'),(112,1953,'Martin Brest'),(113,1953,'Ron Clement'),(114,1953,'Thomas Carter'),(115,1954,'Donald Petrie'),(116,1954,'Joel and Ethan Coen'),(117,1954,'Joel Coen'),(118,1954,'Ron Howard'),(119,1955,'Andy Tennant'),(120,1955,'Bill Condon'),(121,1955,'Bill Paxton'),(122,1955,'Bill Pohlad'),(123,1955,'Catherine Hardwicke'),(124,1955,'David Frankel'),(125,1955,'David Twohy'),(126,1955,'Michael Tollin'),(127,1955,'Mike Gabriel'),(128,1955,'Peter Segal'),(129,1956,'Barry W. Blaustein'),(130,1956,'D.J. MacHale'),(131,1956,'David McNally'),(132,1956,'Gary Ross'),(133,1956,'James Gartner'),(134,1956,'John Lee Hancock'),(135,1956,'Larry Charles'),(136,1956,'Mel Gibson'),(137,1956,'Peter Farrelly'),(138,1956,'Tom Hanks'),(139,1956,'Tony Gilroy'),(140,1957,'Brad Bird'),(141,1957,'Cameron Crowe'),(142,1957,'Chris Wedge'),(143,1957,'David Silverman'),(144,1957,'Jay Roach'),(145,1957,'John Lasseter'),(146,1957,'Spike Lee'),(147,1958,'Bobby Farrelly'),(148,1958,'Chris Buck '),(149,1958,'Chris Columbus'),(150,1958,'Chuck Russell'),(151,1958,'David O. Russell'),(152,1958,'Gregg Fienberg'),(153,1958,'Keenen Ivory Wayans'),(154,1958,'R.J. Cutler'),(155,1958,'Stephen Herek'),(156,1958,'Tim Burton'),(157,1958,'Tom Shadyac'),(158,1958,'Victor Salva'),(159,1959,'Alan Taylor'),(160,1959,'Dan Gilroy'),(161,1959,'Delbert Shoopman'),(162,1959,'Frank Darabont'),(163,1959,'Nick Cassavetes'),(164,1959,'Richard LaGravenese'),(165,1959,'Sam Raimi'),(166,1960,'Doug Lefler'),(167,1960,'Edwin Baily'),(168,1960,'Eric Darnell'),(169,1960,'Gary Trousdale'),(170,1960,'Richard Linklater'),(171,1960,'Rob Marshall'),(172,1960,'Sean Penn'),(173,1960,'Steve Martino'),(174,1960,'Theodore Melfi'),(175,1960,'Tim Johnson'),(176,1961,'Aaron Sorkin'),(177,1961,'Andrew Zimmern'),(178,1961,'Brian Helgeland'),(179,1961,'Byron Allen'),(180,1961,'George Clooney'),(181,1961,'Gregory Jacobs'),(182,1961,'John Stockwell'),(183,1961,'Jonathan Mostow'),(184,1961,'Mark Dindal'),(185,1961,'Shane Black'),(186,1961,'Steve Antin'),(187,1961,'Todd Haynes'),(188,1962,'Chris Sanders'),(189,1962,'David Fincher'),(190,1962,'Frank Coraci'),(191,1962,'Kevin Lima'),(192,1962,'Paul Feig'),(193,1962,'Peter Landesman'),(194,1962,'Peter Ramsey'),(195,1962,'Peyton Reed'),(196,1962,'Stephen Sommers'),(197,1963,'Brian Robbins'),(198,1963,'Gavin O\'Connor'),(199,1963,'Grant Heslov'),(200,1963,'Greg Daniels'),(201,1963,'James Mangold'),(202,1963,'John Cameron Mitchell'),(203,1963,'John Eisendrath'),(204,1963,'Jon Turteltaub'),(205,1963,'Kent Alterman'),(206,1963,'Mark Christopher'),(207,1963,'Neil Burger'),(208,1963,'Quentin Tarantino'),(209,1963,'Rich Moore'),(210,1963,'Ronnie Baxter'),(211,1963,'Steven Soderbergh'),(212,1963,'Sue Tenney'),(213,1963,'Ted Demme'),(214,1964,'Adam Shankman'),(215,1964,'Brad Anderson'),(216,1964,'Eric Rhone'),(217,1964,'Gore Verbinski'),(218,1964,'Joss Whedon'),(219,1964,'Mark Steven Johnson'),(220,1964,'Matt Dearborn'),(221,1964,'Michael W. Watkins'),(222,1964,'Peter Berg'),(223,1964,'Tim Miller'),(224,1964,'Tom McGrath'),(225,1964,'Tommy Lynch'),(226,1965,'Andrew Stanton'),(227,1965,'Ben Stiller'),(228,1965,'Bob Kushell'),(229,1965,'Bryan Singer'),(230,1965,'D.J. Caruso'),(231,1965,'David Cherniack'),(232,1965,'Doug Liman'),(233,1965,'Jonathan Dayton'),(234,1965,'Lana Wachowski'),(235,1965,'Michael Bay'),(236,1965,'ndrew Stanton'),(237,1965,'Paul Weitz'),(238,1965,'Richard Horne'),(239,1965,'The Wachowskis'),(240,1965,'Tony Bancroft'),(241,1965,'Yaky Ortega'),(242,1966,'Adam Rifkin'),(243,1966,'Anne Fletcher'),(244,1966,'Antoine Fuqua'),(245,1966,'Bennett Miller'),(246,1966,'Boaz Yakin'),(247,1966,'Chris Renaud'),(248,1966,'David Schwimmer'),(249,1966,'Eric Radomski'),(250,1966,'Ericson Core'),(251,1966,'Greg Mottola'),(252,1966,'J.J. Abrams'),(253,1966,'James Gunn'),(254,1966,'Jon Favreau'),(255,1966,'Len Wiseman'),(256,1966,'Mark Osborne'),(257,1966,'Matt Reeves'),(258,1966,'Scott Derrickson'),(259,1966,'Tom Hertz'),(260,1966,'Tom McCarthy'),(261,1966,'Zack Snyder'),(262,1967,'Bob Koherr'),(263,1967,'Bruce Burgess'),(264,1967,'Craig Gillespie'),(265,1967,'Eric Bress'),(266,1967,'James Wong'),(267,1967,'Jason Moore'),(268,1967,'John Hamburg'),(269,1967,'Judd Apatow'),(270,1967,'Kimberly Peirce'),(271,1967,'Kirk DeMicco'),(272,1967,'Lee Unkrich'),(273,1968,'Adam McKay'),(274,1968,'Bille Woodruff'),(275,1968,'Byron Howard'),(276,1968,'Chad Stahelski'),(277,1968,'Christopher McQuarrie'),(278,1968,'David Ayer'),(279,1968,'George Nolfi'),(280,1968,'Glenn Ficarra'),(281,1968,'Jamie Brown'),(282,1968,'McG'),(283,1968,'Pete Docter'),(284,1968,'Rob Letterman'),(285,1968,'Robert Rodriguez'),(286,1969,'Andy Fickman'),(287,1969,'Brett Ratner'),(288,1969,'Chris Weitz'),(289,1969,'Curt Morgan'),(290,1969,'Daniel Myrick'),(291,1969,'Darren Aronofsky'),(292,1969,'David Dobkin'),(293,1969,'F. Gary Gray'),(294,1969,'James DeMonaco'),(295,1969,'James Gray'),(296,1969,'Joe Carnahan'),(297,1969,'Kyle Balda'),(298,1969,'Mark Andrews'),(299,1969,'Noah Baumbach'),(300,1969,'Ree Drummond'),(301,1969,'Spike Jonze'),(302,1969,'Trey Parker'),(303,1969,'Wes Anderson'),(304,1970,'Anthony and Joe Russo'),(305,1970,'Brad Meltzer'),(306,1970,'David Leitch'),(307,1970,'Greg Manwaring'),(308,1970,'Heidi Honeycutt'),(309,1970,'Liam Lynch'),(310,1970,'M. Night Shyamalan'),(311,1970,'Mara Brock Akil'),(312,1970,'Michael Charles Hill'),(313,1970,'Mike Mitchell'),(314,1970,'Paul Thomas Anderson'),(315,1970,'Paul Wheeler'),(316,1970,'Richard Lopez'),(317,1970,'Robert Ben Garant'),(318,1970,'Scott Cooper'),(319,1970,'Stephen Chbosky'),(320,1970,'Taylor Sheridan'),(321,1970,'Tim Story'),(322,1970,'Todd Phillips'),(323,1971,'Chris Miller'),(324,1971,'Craig Brewer'),(325,1971,'Francis Lawrence'),(326,1971,'Patty Jenkins'),(327,1971,'Sofia Coppola'),(328,1972,'Albert Hughes'),(329,1972,'Ava DuVernay'),(330,1972,'Ben Affleck'),(331,1972,'Eli Craig'),(332,1972,'Eli Roth'),(333,1972,'Jennifer Yuh Nelson'),(334,1972,'Justin Lin'),(335,1973,'Akiva Schaffer'),(336,1973,'Andrew Kreisberg'),(337,1973,'Don Hall'),(338,1973,'Luke Greenfield'),(339,1973,'Peter Sollett'),(340,1973,'Rian Johnson'),(341,1973,'Seth MacFarlane'),(342,1973,'Trish Sie'),(343,1974,'David Robert Mitchell'),(344,1974,'Elizabeth Banks'),(345,1974,'Jake Kasdan'),(346,1974,'Joseph Kosinski'),(347,1974,'Kenya Barris'),(348,1974,'Marc Webb'),(349,1974,'Ruben Fleischer'),(350,1974,'Seth Gordon'),(351,1974,'Steve Byrne'),(352,1974,'Will Gluck'),(353,1975,'Angelina Jolie'),(354,1975,'Bob Persichetti'),(355,1975,'Bradley Cooper'),(356,1975,'David Gordon Green'),(357,1975,'Drew Barrymore'),(358,1975,'Drew Goddard'),(359,1975,'Jamie LeClaire'),(360,1975,'Kate Shiers'),(361,1975,'Matt Tarses'),(362,1975,'Phil Lord'),(363,1975,'Rawson Marshall Thurber'),(364,1975,'Richard Kelly'),(365,1975,'Zach Braff'),(366,1976,'Colin Trevorrow'),(367,1976,'Dan Scanlon'),(368,1976,'Jeremy Saulnier'),(369,1976,'Jonathan Levine'),(370,1976,'Nicholas Stoller'),(371,1977,'Cary Joji Fukunaga'),(372,1977,'Daym Drops'),(373,1977,'Dee Rees'),(374,1977,'Jeff Wadlow'),(375,1977,'Nkechi Okoro Carroll'),(376,1977,'Tom Willis'),(377,1978,'James Franco'),(378,1978,'Josh Gordon'),(379,1978,'Nathan Greno'),(380,1979,'Aaron Cooley'),(381,1979,'Anna Boden'),(382,1979,'Darren Lynn Bousman'),(383,1979,'Jared Hess'),(384,1979,'John Krasinski'),(385,1979,'Jon M. Chu'),(386,1979,'Jordan Peele'),(387,1979,'Josh Boone'),(388,1979,'Nate Parker'),(389,1979,'Ryan Little'),(390,1980,'Anthony Uro'),(391,1980,'Matthew Lynn'),(392,1980,'Sarah Foudy'),(393,1980,'Tiller Russell'),(394,1980,'Wes Ball'),(395,1981,'Ben Hoffman'),(396,1981,'Dan Trachtenberg'),(397,1981,'Holly Sorensen'),(398,1981,'Jon Watts'),(399,1981,'Joseph Gordon-Levitt'),(400,1981,'Josh Cooley'),(401,1981,'Tarek El Moussa'),(402,1982,'Eddie Huang'),(403,1982,'Luis Manzo'),(404,1982,'Nick Zeig-Owens'),(405,1982,'Skyler Page'),(406,1983,'Robert Eggers'),(407,1984,'Jordan Vogt-Roberts'),(408,1985,'Damien Chazelle'),(409,1985,'Josh Dorsey'),(410,1986,'Ari Aster'),(411,1986,'Mike Arnold'),(412,1986,'Ryan Coogler'),(413,1987,'Brian Lazarte'),(414,1988,'Steven Caple Jr.'),(415,1928,'Patrick McGoohan'),(416,1886,'Michael Curtiz'),(417,1889,'Charlie Chaplin'),(418,1899,'Alfred Hitchcock'),(419,1906,'Harry Watt'),(420,1908,'David Lean'),(421,1914,'Ken Annakin'),(422,1923,'Richard Attenborough'),(423,1936,'Hugh Hudson'),(424,1936,'Ken Loach'),(425,1937,'Ridley Scott'),(426,1938,'Richard Marquand'),(427,1938,'Waris Hussein'),(428,1941,'Adrian Lyne'),(429,1941,'Stephen Frears'),(430,1942,'Bob Spiers'),(431,1942,'Mike Newell'),(432,1942,'Terry Jones'),(433,1943,'Mick Jackson'),(434,1943,'Roger Graef'),(435,1944,'Alan Parker'),(436,1944,'Jules Pigott'),(437,1944,'Tony Scott'),(438,1945,'Roland Joffé'),(439,1946,'Richard Loncraine'),(440,1947,'Paul Marcus'),(441,1948,'Mark Herman'),(442,1949,'Giles Foster'),(443,1949,'John Madden'),(444,1952,'Tony Kaye'),(445,1953,'Peter Lord'),(446,1954,'Anthony Minghella'),(447,1955,'Monty Don'),(448,1955,'Paul Greengrass'),(449,1956,'Danny Boyle'),(450,1956,'Peter Chelsom'),(451,1956,'Richard Curtis'),(452,1956,'Roger Michell'),(453,1957,'Charles McDougall'),(454,1957,'Paul Merton'),(455,1957,'Peter Moss'),(456,1957,'Phyllida Lloyd'),(457,1959,'Brian Percival'),(458,1959,'Simon Cowell'),(459,1960,'Julian Holmes'),(460,1960,'Kenneth Branagh'),(461,1960,'Sharon Maguire'),(462,1960,'Simon Curtis'),(463,1960,'Stephen Butchard'),(464,1960,'Stephen Daldry'),(465,1960,'Stuart Orme'),(466,1961,'Jeff Pope'),(467,1961,'Simon West'),(468,1962,'Ralph Fiennes'),(469,1963,'Armando Iannucci'),(470,1963,'David Yates'),(471,1963,'James Marsh'),(472,1963,'Richard Macer'),(473,1963,'Russell T. Davies'),(474,1964,'James Strong'),(475,1964,'Peter Cattaneo'),(476,1964,'Stephen Norrington'),(477,1965,'Daniel Percival'),(478,1965,'Paul McGuigan'),(479,1965,'Paul W.S. Anderson'),(480,1965,'Sam Mendes'),(481,1966,'Dexter Fletcher'),(482,1967,'Kevin Macdonald'),(483,1967,'Nick Love'),(484,1967,'Sam Taylor-Johnson'),(485,1968,'Guy Ritchie'),(486,1968,'Justin Chadwick'),(487,1969,'David Slade'),(488,1969,'Steve McQueen'),(489,1970,'Alan Grint'),(490,1970,'Alex Garland'),(491,1970,'Christopher Nolan'),(492,1970,'Neil Marshall'),(493,1971,'Duncan Jones'),(494,1971,'Matthew Vaughn'),(495,1972,'Greg Tiernan'),(496,1972,'Joe Wright'),(497,1972,'Julie Anne Robinson'),(498,1972,'om Hooper'),(499,1972,'Rupert Wyatt'),(500,1972,'Tom Hooper'),(501,1974,'Edgar Wright'),(502,1974,'Saul Dibb'),(503,1975,'Gareth Edwards'),(504,1975,'Tom Bidwell'),(505,1976,'Thea Sharrock'),(506,1977,'James Watkins'),(507,1977,'Mark Tonderai'),(508,1978,'Adnan Ahmad'),(509,1982,'Laura Henry-Allain'),(510,1960,'Choi Jong-il'),(511,1960,'Kang Je-gyu'),(512,1960,'Park Jin-suk'),(513,1963,'Park Chan-wook'),(514,1965,'Yoo Jong Sun'),(515,1967,'Jo Young Kwang'),(516,1967,'Kang Shin-hyo'),(517,1968,'Kim Sang-hyub'),(518,1969,'Bong Joonho'),(519,1972,'Shin Yong-hwi'),(520,1976,'Lee Dong Min'),(521,1978,'Yeon Sang-ho'),(522,1980,'Baek Ho-min'),(523,1980,'Han Joon-seo'),(524,1980,'Lee Ji-min'),(525,1985,'Choi Jung-in'),(526,1978,'Ga-Ho Lau'),(527,1951,'Zhang Yimou'),(528,1964,'Ying Da'),(529,1975,'Gao Xixi'),(530,1979,'Cai Cong'),(531,1982,'Miao Shu'),(532,1983,'Haoyuan Xu'),(533,NULL,'陈歆宇'),(534,1910,'Akira Kurosawa'),(535,1935,'Fujio Akatsuka'),(536,1935,'Isao Takahata'),(537,1941,'ayao Miyazaki'),(538,1941,'Hayao Miyazaki'),(539,1941,'Yoshiyuki Tomino'),(540,1942,'Kenji Kodama'),(541,1947,'Fumio Kurokawa'),(542,1947,'Hiroshi Sasagawa'),(543,1949,'Norifumi Suzuki'),(544,1950,'Koichi Takemoto'),(545,1954,'Hitoshi Nanba'),(546,1954,'Katsuhiro Otomo'),(547,1955,'Hideya Takahashi'),(548,1955,'Takeshi Natsuhara'),(549,1955,'Yōjirō Takita'),(550,1955,'Yukihiro Miyamoto'),(551,1956,'Kinji Yoshimoto'),(552,1959,'Kazuhiko Kato'),(553,1960,'Hiroshi Kimura'),(554,1960,'Noriyuki Abe'),(555,1962,'Keizō Kusakawa'),(556,1962,'Takashi Watanabe'),(557,1965,'Kazuki Akane'),(558,1967,'Mamoru Kanbe'),(559,1969,'Mitsuru Miura'),(560,1970,'Kenichi Imaizumi'),(561,1970,'Shugo Praico'),(562,1972,'Takashi Shimizu'),(563,1973,'Kurosugi Shinsaku'),(564,1973,'Makoto Shinkai'),(565,1974,'Ayato Matsuda'),(566,1974,'Manabu Okamoto'),(567,1974,'Toshifumi Kawase'),(568,1975,'Kazuya Nomura'),(569,1975,'Maha Harada'),(570,1976,'Naohito Takahashi'),(571,1977,'Yoshinobu Tokumoto'),(572,1979,'Rei Ishiguro'),(573,1917,'Ramanand Sagar'),(574,1945,'Shekhar Kapur'),(575,1962,'Rajkumar Hirani'),(576,1963,'Rakeysh Omprakash Mehra'),(577,1965,'Farah Khan'),(578,1969,'Gul Khan'),(579,1969,'Vikram Bhatt'),(580,1970,'Arif Khan'),(581,1971,'Aditya Chopra'),(582,1971,'Imtiaz Ali'),(583,1971,'Shimit Amin'),(584,1973,'Neeraj Pandey'),(585,1974,'Anurag Basu'),(586,1974,'Raja Krishna Menon'),(587,1976,'Nitesh Tiwari'),(588,1977,'Arvind Babbal'),(589,1982,'Ali Abbas Zafar'),(590,1982,'Neel Upadhye'),(591,1983,'Aditya Dhar'),(592,1984,'Meet Sampat'),(593,1894,'Jean Renoir'),(594,1932,'Louis Malle'),(595,1943,'Jean-Jacques Annaud'),(596,1946,'Serge Moati'),(597,1947,'Marion Sarraut'),(598,1951,'Claude-Michel Rome'),(599,1953,'ean-Pierre Jeunet'),(600,1953,'Jean-Pierre Jeunet'),(601,1953,'Yves Le Rolland'),(602,1957,'Emmanuel Chain'),(603,1959,'Luc Besson'),(604,1959,'Pierre Coffin'),(605,1959,'uc Besson'),(606,1960,'Christophe Gans'),(607,1963,'Christophe Barratier'),(608,1963,'Jean-Xavier de Lestrade'),(609,1963,'Michel Gondry'),(610,1964,'Pierre Morel'),(611,1965,'Bibo Bergeron'),(612,1965,'Olivier Megaton'),(613,1966,'Louis Leterrier'),(614,1967,'Michel Hazanavicius'),(615,1967,'Olivier Dahan'),(616,1971,'Virginie Brac'),(617,1972,'Thibaut Martin'),(618,1973,'Olivier Nakache'),(619,1975,'Fabrice Gardel'),(620,1978,'Alexandre Aja'),(621,1980,'Fabien Fournier'),(622,1980,'Jérémie Hoarau'),(623,1990,'Sophie Roze Sengelin'),(624,1929,'Heiner Carow'),(625,1929,'Heinz Oskar Wuttig'),(626,1932,'Edgar Reitz'),(627,1941,'Wolfgang Petersen'),(628,1950,'Walter Bannert'),(629,1955,'Roland Emmerich'),(630,1957,'Oliver Hirschbiegel'),(631,1963,'René Steinbach'),(632,1965,'Oliver Kalkofe'),(633,1965,'Tom Tykwer'),(634,1968,'Robert Schwentke'),(635,1973,'Florian Henckel von Donnersmarck'),(636,1973,'Marco Schnabel'),(637,1975,'Lexi Alexander'),(638,1980,'Katharina Jeschke'),(639,1982,'Ulrich del Mestre'),(640,1892,'Ernst Lubitsch'),(641,1969,'Marc Forster'),(642,1961,'Murat Saraçoğlu'),(643,1963,'Şahin Altuğ'),(644,1970,'İsa Yıldız'),(645,1972,'Faruk Aksoy'),(646,1979,'Alper Çağlar'),(647,1981,'Faruk Teber'),(648,1982,'Sema Ergenekon'),(649,1985,'Baris Yös'),(650,1985,'Sevgi Yılmaz'),(651,1974,'Jaume Collet-Serra'),(652,1935,'Narciso Ibáñez Serrador'),(653,1958,'Vicente Escrivá'),(654,1962,'José Luis Moro'),(655,1967,'Juan Carlos Fresnadillo'),(656,1971,'Joaquín Llamas'),(657,1975,'J.A. Bayona'),(658,1975,'Salvador Garcini'),(659,1978,'Roberto González'),(660,1985,'David Broncano'),(661,1989,'Carolina Iglesias'),(662,1926,'Norman Jewison'),(663,1942,'Timothy Bond'),(664,1943,'David Cronenberg'),(665,1949,'Allan Moyle'),(666,1954,'James Cameron'),(667,1957,'Claude Desrosiers'),(668,1959,'Mary Harron'),(669,1960,'Judith Beauchemin'),(670,1962,'Rick Green'),(671,1963,'François Girard'),(672,1963,'Jean-Marc Vallée'),(673,1965,'Floria Sigismondi'),(674,1967,'Denis Villeneuve'),(675,1969,'Patrick Senécal'),(676,1969,'Vincenzo Natali'),(677,1970,'Dean DeBlois'),(678,1972,'Michael Dowse'),(679,1972,'Robert Valley'),(680,1977,'Jason Reitman'),(681,1978,'Christine Doyon'),(682,1982,'Seth Rogen'),(683,1953,'Paul Haggis'),(684,1946,'Ivan Reitman'),(685,1944,'Peter Weir'),(686,1945,'George Miller'),(687,1949,'Phillip Noyce'),(688,1952,'David Stevens'),(689,1953,'Russell Mulcahy'),(690,1953,'Scott Hicks'),(691,1962,'Baz Luhrmann'),(692,1964,'Stephan Elliott'),(693,1967,'James McTeigue'),(694,1969,'Jennifer Kent'),(695,1971,'Michael Gracey'),(696,1972,'Adam Elliot'),(697,1972,'David Michôd'),(698,1973,'Annabel Crabb'),(699,1973,'Robert Luketic'),(700,1974,'Justin Kurzel'),(701,1980,'Garth Davis'),(702,1986,'Anthony Maras'),(703,1898,'Sergei Eisenstein'),(704,1933,'Elem Klimov'),(705,1948,'Valeriy Ibragimov'),(706,1967,'Fedor Bondarchuk'),(707,1970,'Genndy Tartakovsky'),(708,1977,'Diamara Nizhnikovskaya'),(709,1980,'Artyom Aksenenko'),(710,1983,'Ilya Kulikov'),(711,1961,'Timur Bekmambetov'),(712,1953,'Walcyr Carrasco'),(713,1955,'Fernando Meirelles'),(714,1965,'Carlos Saldanha'),(715,1976,'Rita Lobo'),(716,1985,'Halder Gomes'),(717,1919,'Gillo Pontecorvo'),(718,1922,'ier Paolo Pasolini'),(719,1929,'Sergio Leone'),(720,1941,'Bernardo Bertolucci'),(721,1943,'Gianfranco Albano'),(722,1952,'Roberto Benigni'),(723,1956,'Giuseppe Tornatore'),(724,1963,'Simona Ercolani'),(725,1965,'Corrado Guzzanti'),(726,1967,'Gabriele Muccino'),(727,1970,'Paolo Sorrentino'),(728,1971,'Luca Guadagnino'),(729,1975,'Ivan Cotroneo'),(730,1981,'Giuseppe Gagliardi'),(731,1897,'Frank Capra'),(732,1952,'Isabella Rossellini'),(733,1941,'George P. Cosmatos'),(734,1976,'Mac Alejandre'),(735,1938,'Paul Verhoeven'),(736,1943,'Jan de Bont'),(737,1955,'Anton Corbijn'),(738,1978,'Marc van Sambeek'),(739,1979,'Thijs Römer'),(740,1984,'Krit Sukramongkol'),(741,1946,'John Woo'),(742,1956,'Lee Tim-sing'),(743,1961,'Hung Chiu-fung'),(744,1963,'Tommy Leung'),(745,1964,'Patrick Yau'),(746,1965,'Terry Tong'),(747,1967,'Wilson Yip'),(748,1968,'Alan Mak'),(749,1968,'Poon Ka-tak'),(750,1961,'Alfonso Cuarón'),(751,1963,'Alejandro González Iñárritu'),(752,1964,'Fernanda Villeli'),(753,1964,'Guillermo del Toro'),(754,1970,'Rodolfo Antúnez'),(755,1977,'Alejandro Monteverde'),(756,1980,'Gerardo Daniel de la Cerda Barajas'),(757,1983,'Edgar Medina'),(758,1889,'Carl Theodor Dreyer'),(759,1960,'Susanne Bier'),(760,1961,'Jesper W. Nielsen'),(761,1970,'Nicolas Winding Refn'),(762,1973,'Nicolai Fuglsig'),(763,1982,'Mads Andersen'),(764,1952,'José Eduardo Moniz'),(765,1970,'Carla Albuquerque'),(766,1980,'Frederico Pombares'),(767,1938,'Jan Rybkowski'),(768,1941,'Krzysztof Kieślowski'),(769,1941,'rzysztof Kieślowski'),(770,1957,'Paweł Pawlikowski'),(771,1967,'Maciej Pieprzyca'),(772,1980,'Karol Klementewicz'),(773,1983,'Ewelina Gordziejuk'),(774,1959,'Terje Solli'),(775,1965,'Magnus Martens'),(776,1967,'Morten Tyldum'),(777,1971,'Espen Sandberg'),(778,1975,'Thora Lorentzen'),(779,1979,'Tommy Wirkola'),(780,1954,'Ang Lee'),(781,1933,'Roman Polanski'),(782,1963,'Oleksandr Tkachenko'),(783,1964,'Russell Crowe'),(784,1959,'Juan José Campanella'),(785,1963,'Gaspar Noé'),(786,1973,'Andrés Muschietti'),(787,1973,'Andy Muschietti'),(788,1980,'Bruno Bluwol'),(789,1985,'Ricardo A. Solla'),(790,1906,'Billy Wilder'),(791,1907,'Fred Zinnemann'),(792,1942,'Michael Haneke'),(793,1957,'Jaco Van Dormael'),(794,1964,'Bart De Pauw'),(795,1970,'Philippe De Schepper'),(796,1971,'Jakob Verbruggen'),(797,1969,'Danis Tanović'),(798,1960,'Alexander Witt'),(799,1972,'Alejandro Amenábar'),(800,1976,'Pablo Larraín'),(801,1978,'Mauricio Katz'),(802,1966,'Gustavo Bolívar'),(803,1932,'Miloš Forman'),(804,1963,'Alex Proyas'),(805,1977,'Marwan Hamed'),(806,1981,'Nabil Esmat'),(807,1985,'Farid Shawqy'),(808,1959,'Renny Harlin'),(809,1982,'Ilkka Vanne'),(810,1985,'Mikko Kuparinen'),(811,1973,'Yorgos Lanthimos'),(812,1975,'Papandreou A.E.'),(813,1937,'József Nepp'),(814,1977,'László Nemes'),(815,1966,'Baltasar Kormákur'),(816,1959,'Majid Majidi'),(817,1985,'Nima Bank'),(818,1949,'Jim Sheridan'),(819,1950,'Neil Jordan'),(820,1966,'Lenny Abrahamson'),(821,1969,'John Carney'),(822,1970,'Martin McDonagh'),(823,1973,'Kirsten Sheridan'),(824,1929,'Menahem Golan'),(825,1970,'Oren Peli'),(826,1977,'James Wan'),(827,1943,'Martin Campbell'),(828,1945,'Roger Donaldson'),(829,1954,'Jane Campion'),(830,1957,'Peter Burger'),(831,1961,'Peter Jackson'),(832,1964,'Andrew Niccol'),(833,1966,'Andrew Adamson'),(834,1967,'Andrew Dominik'),(835,1967,'Niki Caro'),(836,1969,'Jonathan Brough'),(837,1975,'Taika Waititi'),(838,1952,'Terry George'),(839,1975,'Saima Akram Chaudary'),(840,1983,'Mihai Bendeac'),(841,1971,'Nikola Pejaković'),(842,1981,'Dominik Dán'),(843,1963,'Gavin Hood'),(844,1979,'Neill Blomkamp'),(845,1918,'Ingmar Bergman'),(846,1946,'Lasse Hallström'),(847,1960,'Mikael Håfström'),(848,1965,'Magnus Sjöstrom'),(849,1965,'Tomas Alfredson'),(850,1974,'Måns Mårlind'),(851,1977,'Daniel Espinosa'),(852,1973,'Yasser Alazmeh'),(853,1978,'Fede Álvarez'),(854,1954,'Emir Kusturica'),(855,NULL,'Unknown'),(856,NULL,'Various directors'),(857,NULL,'Chris Buck'),(858,NULL,'Gina');

--
-- Table structure for table `episodes`
--

DROP TABLE IF EXISTS `episodes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `episodes` (
                            `id` bigint NOT NULL AUTO_INCREMENT,
                            `episode_number` bigint NOT NULL,
                            `movie_id` bigint NOT NULL,
                            `source` longtext,
                            `duration` bigint DEFAULT NULL,
                            `created_at` datetime(3) DEFAULT NULL,
                            `updated_at` datetime(3) DEFAULT NULL,
                            `deleted_at` datetime(3) DEFAULT NULL,
                            PRIMARY KEY (`id`),
                            UNIQUE KEY `idx_episode_movie` (`episode_number`,`movie_id`),
                            KEY `fk_episodes_movie` (`movie_id`),
                            KEY `idx_episodes_deleted_at` (`deleted_at`),
                            CONSTRAINT `fk_episodes_movie` FOREIGN KEY (`movie_id`) REFERENCES `movies` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=89 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `episodes`
--

INSERT INTO `episodes` VALUES (51,1,1,'https://vip.opstream12.com/20220217/224_5d301141/index.m3u8',142,'2024-10-28 20:05:37.920','2024-10-28 20:05:37.920',NULL),(52,1,2,'https://vip.opstream12.com/20220526/16828_1eb2cff5/index.m3u8',179,'2024-10-29 15:12:23.397','2024-10-29 15:12:23.397',NULL),(53,1,3,'https://vip.opstream12.com/20220508/14921_bd357a00/index.m3u8',144,'2024-10-29 15:13:09.085','2024-10-29 15:13:09.085',NULL),(54,1,4,'https://vip.opstream14.com/20220627/16018_ac1e6edf/index.m3u8',96,'2024-10-29 15:13:39.819','2024-10-29 15:13:39.819',NULL),(55,1,5,'https://vip.opstream11.com/20220311/1498_ff3b85f9/index.m3u8',201,'2024-10-29 15:14:08.374','2024-10-29 15:14:08.374',NULL),(56,1,6,'https://vip.opstream11.com/20220311/1495_da17bad1/index.m3u8',178,'2024-10-29 15:14:32.015','2024-10-29 15:14:32.015',NULL),(57,1,7,'https://vip.opstream11.com/20220311/1494_f31e11c6/index.m3u8',179,'2024-10-29 15:15:28.327','2024-10-29 15:15:28.327',NULL),(58,1,8,'https://vip.opstream15.com/20220226/222_b2f091d3/index.m3u8',143,'2024-10-29 15:16:08.175','2024-10-29 15:16:08.175',NULL),(59,1,9,'https://vip.opstream11.com/20220309/1352_798e9f08/index.m3u8',128,'2024-10-29 15:21:22.945','2024-10-29 15:21:22.945',NULL),(60,1,10,'https://vip.opstream14.com/20220317/1331_98f0f18b/index.m3u8',123,'2024-10-29 15:21:45.923','2024-10-29 15:21:45.923',NULL),(61,1,11,'https://vip.opstream12.com/20240917/23710_07db8d8b/index.m3u8',112,'2024-10-29 15:22:07.709','2024-10-29 15:22:07.709',NULL),(62,1,12,'https://vip.opstream17.com/20240122/98_d910cc1d/index.m3u8',103,'2024-10-29 15:22:27.953','2024-10-29 15:22:27.953',NULL),(63,1,14,'https://vip.opstream11.com/20220314/1901_ef246108/index.m3u8',197,'2024-10-29 15:22:44.999','2024-10-29 15:22:44.999',NULL),(64,1,35,'https://vip.opstream11.com/20241029/55859_ae8cc0d9/index.m3u8',79,'2024-11-03 10:21:30.126','2024-11-03 10:21:30.126',NULL),(67,1,36,'https://img.ophim.live/uploads/movies/nang-tho-thumb.jpg',113,'2024-11-03 10:29:45.399','2024-11-03 10:29:45.399',NULL),(68,1,37,'https://vip.opstream11.com/20231115/49131_ee15ba0c/index.m3u8',49,'2024-11-03 10:34:13.201','2024-11-03 10:34:13.201',NULL),(69,2,37,'https://vip.opstream11.com/20231118/49290_03dadf54/index.m3u8',49,'2024-11-03 10:34:22.996','2024-11-03 10:34:22.996',NULL),(70,3,37,'https://vip.opstream11.com/20231130/49634_0ef768d1/index.m3u8',49,'2024-11-03 10:34:29.911','2024-11-03 10:34:29.911',NULL),(71,4,37,'https://vip.opstream16.com/20231204/38614_e329189d/index.m3u8',49,'2024-11-03 10:34:37.186','2024-11-03 10:34:37.186',NULL),(73,5,37,'https://vip.opstream11.com/20231225/50110_78e229f6/index.m3u8',49,'2024-11-03 10:34:48.805','2024-11-03 10:34:48.805',NULL),(74,6,37,'https://vip.opstream11.com/20231225/50111_5e9235a3/index.m3u8',49,'2024-11-03 10:35:12.869','2024-11-03 10:35:12.869',NULL),(75,7,37,'ttps://vip.opstream11.com/20231225/50112_cd8b4cd1/index.m3u8',49,'2024-11-03 10:35:20.995','2024-11-03 10:35:20.995',NULL),(76,8,37,'https://vip.opstream11.com/20240101/50385_0e66d144/index.m3u8',49,'2024-11-03 10:35:28.310','2024-11-03 10:35:28.310',NULL),(78,2,36,'https://vip.opstream17.com/20240308/2275_54dbd4d6/index.m3u8',60,'2024-11-03 10:40:30.329','2024-11-03 10:40:30.329',NULL),(79,3,36,'https://vip.opstream17.com/20240308/2276_2825dc8e/index.m3u8',60,'2024-11-03 10:40:37.675','2024-11-03 10:40:37.675',NULL),(80,4,36,'https://vip.opstream17.com/20240308/2278_7ab8e765/index.m3u8',60,'2024-11-03 10:40:44.289','2024-11-03 10:40:44.289',NULL),(81,2,38,'https://vip.opstream10.com/20241003/27985_b1a5a84a/index.m3u8',60,'2024-11-03 10:43:59.226','2024-11-03 10:43:59.226',NULL),(82,1,38,'https://vip.opstream10.com/20241003/27984_b76f966a/index.m3u8',60,'2024-11-03 10:44:05.675','2024-11-03 10:44:05.675',NULL),(83,1,39,'https://vip.opstream12.com/20241006/23860_7c32ecee/index.m3u8',101,'2024-11-03 10:46:33.172','2024-11-03 10:46:33.172',NULL),(84,1,40,'https://vip.opstream11.com/20231113/49065_cc599714/index.m3u8',103,'2024-11-03 10:49:23.108','2024-11-03 10:49:23.108',NULL),(85,1,41,'https://vip.opstream14.com/20220322/2362_e5d3ee10/index.m3u8',79,'2024-11-03 10:52:09.115','2024-11-03 10:52:09.115',NULL),(86,1,42,'https://vip.opstream15.com/20220324/1930_8e4eae60/index.m3u8',120,'2024-11-08 09:48:40.556','2024-11-08 09:48:40.556',NULL),(87,1,43,'https://vip.opstream11.com/20241112/57264_5ca4c0b3/index.m3u8',121,'2024-11-14 08:38:17.628','2024-11-14 08:38:17.628',NULL),(88,1,44,'https://vip.opstream16.com/20220222/21_a5e2d3e4/index.m3u8',91,'2024-12-09 09:45:25.592','2024-12-09 09:45:25.592',NULL);

--
-- Table structure for table `movie_actor`
--

DROP TABLE IF EXISTS `movie_actor`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `movie_actor` (
                               `movie_id` bigint NOT NULL,
                               `actor_id` bigint NOT NULL,
                               PRIMARY KEY (`movie_id`,`actor_id`),
                               KEY `fk_movie_actor_actor` (`actor_id`),
                               CONSTRAINT `fk_movie_actor_actor` FOREIGN KEY (`actor_id`) REFERENCES `actors` (`id`),
                               CONSTRAINT `fk_movie_actor_movie` FOREIGN KEY (`movie_id`) REFERENCES `movies` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `movie_actor`
--

INSERT INTO `movie_actor` VALUES (1,14),(2,17),(42,17),(44,17),(2,18),(3,19),(3,20),(44,20),(3,21),(4,22),(41,22),(4,23),(4,24),(5,25),(6,25),(7,25),(5,26),(6,26),(7,26),(42,26),(5,27),(6,27),(7,27),(5,28),(6,28),(7,28),(8,29),(8,30),(8,31),(9,32),(9,33),(40,33),(9,34),(43,34),(10,35),(10,36),(10,37),(11,38),(11,39),(11,40),(40,40),(12,41),(12,42),(12,43),(14,44),(14,45),(14,46),(36,52),(43,52),(36,56),(38,58),(38,60),(39,65),(35,69),(37,69),(39,73),(35,74),(37,74);

--
-- Table structure for table `movie_category`
--

DROP TABLE IF EXISTS `movie_category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `movie_category` (
                                  `movie_id` bigint NOT NULL,
                                  `category_id` bigint NOT NULL,
                                  PRIMARY KEY (`movie_id`,`category_id`),
                                  KEY `fk_movie_category_category` (`category_id`),
                                  CONSTRAINT `fk_movie_category_category` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`),
                                  CONSTRAINT `fk_movie_category_movie` FOREIGN KEY (`movie_id`) REFERENCES `movies` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `movie_category`
--

INSERT INTO `movie_category` VALUES (3,1),(5,1),(6,1),(7,1),(9,1),(10,1),(14,1),(44,1),(1,2),(5,2),(6,2),(7,2),(10,2),(11,2),(14,2),(44,2),(37,3),(39,4),(40,4),(1,5),(2,5),(3,5),(4,5),(5,5),(6,5),(7,5),(8,5),(10,5),(12,5),(36,5),(40,5),(11,6),(35,7),(41,7),(43,7),(8,8),(37,8),(38,8),(42,8),(14,9),(36,9),(43,17);

--
-- Table structure for table `movie_director`
--

DROP TABLE IF EXISTS `movie_director`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `movie_director` (
                                  `movie_id` bigint NOT NULL,
                                  `director_id` bigint NOT NULL,
                                  PRIMARY KEY (`movie_id`,`director_id`),
                                  KEY `fk_movie_director_director` (`director_id`),
                                  CONSTRAINT `fk_movie_director_director` FOREIGN KEY (`director_id`) REFERENCES `directors` (`id`),
                                  CONSTRAINT `fk_movie_director_movie` FOREIGN KEY (`movie_id`) REFERENCES `movies` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `movie_director`
--

INSERT INTO `movie_director` VALUES (1,1),(2,1),(1,2),(2,2),(5,15),(5,25),(12,45),(11,60),(11,63),(44,65),(14,491),(43,641);

--
-- Table structure for table `movies`
--

DROP TABLE IF EXISTS `movies`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `movies` (
                          `id` bigint NOT NULL AUTO_INCREMENT,
                          `name` longtext,
                          `year` bigint DEFAULT NULL,
                          `num_episodes` bigint DEFAULT NULL,
                          `description` longtext,
                          `language` longtext,
                          `country_id` bigint DEFAULT NULL,
                          `time_for_ep` bigint DEFAULT NULL,
                          `thumbnail` longtext,
                          `trailer` longtext,
                          `rate` float DEFAULT NULL,
                          `predict_rate` float DEFAULT NULL,
                          `is_recommended` tinyint(1) DEFAULT NULL,
                          `created_at` datetime(3) DEFAULT NULL,
                          `updated_at` datetime(3) DEFAULT NULL,
                          `deleted_at` datetime(3) DEFAULT NULL,
                          `view` bigint DEFAULT '0',
                          PRIMARY KEY (`id`),
                          KEY `idx_movies_deleted_at` (`deleted_at`),
                          KEY `fk_movies_country` (`country_id`),
                          CONSTRAINT `fk_movies_country` FOREIGN KEY (`country_id`) REFERENCES `countries` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=45 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `movies`
--

INSERT INTO `movies` VALUES (1,'The Shawshank Redemption',1994,1,'Framed for murder, upright banker Andy Dufresne starts a new life at Shawshank Prison and gradually becomes close to older inmate Red.','English',1,142,'https://img.ophim.live/uploads/movies/nha-tu-shawshank-thumb.jpg','https://www.youtube.com/watch?v=P9mwtI82k6E',10,0,1,'2024-10-28 14:36:13.000','2025-02-11 20:41:25.869',NULL,12),(2,'The Godfather',1972,1,'The story begins with the wedding of the beloved daughter of “Don” Vito Corleone (Marlon Brando), the boss of the New York Mafia. His youngest son, Michael (Al Pacino), has just returned from World War II and does not hide his intention to stay out of all family business. The nature of Godfather Corleone’s business is gradually revealed, which is the underground life of those who oppose the corrupt legal system at that time, hiding under the cover of a legitimate olive oil import-export company. The Godfather lives by the old ways, protecting the weak who are bullied by the authorities and is ready to “eliminate” anyone who intends to stand in the way of the organization he is leading.','English',1,179,'https://img.ophim.live/uploads/movies/bo-gia-thumb.jpg','https://www.youtube.com/watch?v=UaVTIH8mujA',9,0,1,'2024-10-24 16:26:30.000','2024-12-09 09:52:55.434',NULL,2),(3,'The Dark Knight',2008,1,'Batman enters the most difficult phase in the fight against evil when the master criminal The Joker appears. The bosses hire The Joker to kill Batman to avenge the illegal assets he exposed. Helping The Joker are all the criminals in Gotham City, who are being narrowed down by the trio Batman - Lieutenant Gordon - Lawyer Harvey Dent. A series of attacks occur, making Batman confused. Officials are murdered, people do not dare to go out at night for fear of being killed, all entrances and exits to Gotham City are blocked. All are arranged by The Joker. Lawyer Dent and Batman\'s longtime friend Rachel Dawes are dragged to two warehouses containing timed bombs. The Joker\'s scheme makes Batman unable to save his friend in time, lawyer Dent luckily escapes death but is burned half of his face. The Joker helps Dent escape from the hospital, at the same time inciting him to take out all his hatred on Batman, the police and the bosses. Maddened by the loss of Dawes, lawyer Dent becomes the evil Two-Face, and his victims almost include Batman and the entire family of Lieutenant Gordon. Not wanting the people of Gotham to lose their spirit when they learn the truth about the leader of the city, Batman decides to take the blame for Dent...','English',NULL,144,'https://img.ophim.live/uploads/movies/ky-si-bong-dem-thumb.jpg','https://www.youtube.com/watch?v=EXeTwQWrcwY',8,0,1,'2024-10-28 14:38:47.000','2024-12-09 11:46:16.258',NULL,4),(4,'12 Angry Men',1957,1,'The defense and prosecution have retired, and the jury is entering the jury room to decide whether a young man is guilty or innocent of murdering his father. What begins as an open-and-shut murder case soon becomes a detective story that presents a succession of clues that create doubt, and a mini-drama about each juror\'s prejudices and preconceptions about the trial, the defendant, AND each other. Based on the play, all the action takes place on the jury room stage.','English',NULL,96,'https://img.ophim.live/uploads/movies/12-nguoi-dan-ong-gian-du-thumb.jpg','https://www.youtube.com/watch?v=TEN-2uTi2c0',9,0,1,'2024-10-19 09:59:23.000','2024-11-13 14:05:30.398',NULL,2),(5,'The Lord of the Rings 3: The Return of the King',2003,1,'Sauron mobilized all his forces for the final fierce battle, attacking Gondor - the land of the bloodline of mankind. Gondor under the leadership of the weak king became too fragile and dangerous. That was when the heroes showed their full potential. With the support of the Rohan warriors, the cursed ghosts in the ravine, Aragorn led the Gondor people against Sauron, he carried out the mission that fate entrusted to him, Aragorn was born to become the King of Gondor, he proved it. At the same time, Frodo and Sam continued to delve deep into the dangerous land of demons to carry out their mission, the mission to destroy the Ring of Power. It can be said that The Return of the King is a part with attractive content with extremely good and extremely epic battles. \"The Lord Of The Ring 3\" not only captivated movie-loving audiences but also conquered the hearts of film critics with 11 Oscars \'2004 including the Oscar for Best Picture and Oscar for Best Director.','English',1,201,'https://img.ophim.live/uploads/movies/chua-te-cua-nhung-chiec-nhan-3-su-tro-lai-cua-nha-vua-thumb.jpg','https://www.youtube.com/watch?v=FTdrUX323Xc',0,0,1,'2024-10-28 14:40:56.000','2024-11-14 08:33:45.879',NULL,4),(6,'The Lord of the Rings 1: The Fellowship of the Ring',2001,1,'The Lord\'s Ring, which carries the power of domination, was accidentally placed in the hands of the young hobbit Frodo. When the wizard Gandalf discovered that the ring once belonged to the tyrant Sauron, Frodo was tasked with taking the ring to the Gorge of Destruction to destroy it. Not alone on the difficult journey, Frodo also received companions from different races: Legolas of the Elf clan; Gimli of the Dwarf clan; Aragon, Boromir and three loyal Hobbit friends... The film is the first part of the classic trilogy \"The Lord of the Rings\", adapted from the novel by writer Tolkien.','English',NULL,178,'https://img.ophim.live/uploads/movies/chua-te-cua-nhung-chiec-nhan-1-hiep-hoi-nhan-than-thumb.jpg','http://www.youtube.com/watch?v=V75dMMIW2B4',9.5,10,1,'2024-10-19 09:59:25.000','2025-02-11 20:37:24.990',NULL,5),(7,'The Lord of the Rings 2: The Two Towers',2002,1,'The journey to destroy the ring continues with many dangers waiting for Frodo ahead. This time, he and Sam have a new companion, Gollum. Meanwhile, Gandalf is lucky to escape death, and his power is much greater. He, Aragon, Legolas and Gimli along with the forest gods find a way to destroy Saruman\'s plot. Following the success of part 1, The Lord Of The Ring 2 inherits and interweaves the old and new character systems, making the film\'s content even more attractive to viewers. The Lord Of The Ring 2 also won 2 Oscars for the categories of sound and effects.','English',NULL,179,'https://img.ophim.live/uploads/movies/chua-te-cua-nhung-chiec-nhan-2-hai-toa-thap-thumb.jpg','http://www.youtube.com/watch?v=V75dMMIW2B4',10,2,1,'2024-10-19 09:59:28.000','2024-12-09 09:50:40.390',NULL,2),(8,'Forrest Gump',1994,1,'The friendly goofball Forrest Gump appeared in almost every major event of the 60s and 70s.','English',NULL,143,'https://img.ophim.live/uploads/movies/forrest-gump-thumb.jpg','https://www.youtube.com/watch?v=bLvqoHBptjg',0,0,1,'2024-10-28 14:46:01.000','2024-10-29 15:16:08.187',NULL,0),(9,'High & Low The Movie',2016,1,'Five rival gangs that dominate the SWORD region unite to face off against a 500-member assault led by a legendary gang boss.','Japanese',NULL,128,'https://m.media-amazon.com/images/M/MV5BOGVmODQwMWUtMmQwZS00YjNhLWEyMzEtNTNmNDE1OTEzNWM5XkEyXkFqcGc@._V1_FMjpg_UX1000_.jpg','https://www.youtube.com/watch?v=KISlBtt4q7o',9,0,1,'2024-10-28 14:46:01.000','2024-11-16 21:41:39.854',NULL,1),(10,'Alpha',2018,1,'An epic adventure set during the last ice age, Alpha is a visually compelling story that shines a light on the origins of man\'s best friend. While on his first hunt with his tribe\'s most elite group, a young man is injured and must learn to survive alone in the wilderness. After the young man corners a lone wolf abandoned by its pack, the pair learn to rely on each other and become unlikely allies, enduring countless dangers and overwhelming odds to find their way home before winter sets in.','English',NULL,123,'https://img.ophim.live/uploads/movies/alpha-nguoi-thu-linh-thumb.jpg','https://www.youtube.com/watch?v=uIxnTi4GmCo',0,0,1,'2024-10-28 14:46:01.000','2024-12-09 12:05:09.605',NULL,1),(11,'The Water Horse',2007,1,'A lonely boy discovers a mysterious egg that hatches a sea creature of Scottish legend.\n\n','English',1,112,'https://img.ophim.live/uploads/movies/huyen-thoai-bieu-sau-thumb.jpg','https://www.youtube.com/watch?v=iuvPpCMgA9U',9,0,1,'2024-10-28 14:46:01.000','2024-12-09 09:52:06.696',NULL,1),(12,'Full Circle',2023,1,'The stories of sit-in skier Trevor Kennison and mountain climber Barry Corbet inform this documentary about the journey of adaptive athletes in extreme sports.','English',1,103,'https://img.ophim.live/uploads/movies/tro-lai-diem-xuat-phat-thumb.jpg','https://www.youtube.com/watch?v=IZh4E9ICcI4',0,0,1,'2024-10-28 14:46:01.000','2024-11-14 08:33:30.131',NULL,0),(14,'Inception',2010,1,'A thief who steals corporate secrets through the use of dream-sharing technology is given the inverse task of planting an idea in the mind of a c.e.o.','English',1,197,'https://img.ophim.live/uploads/movies/ke-danh-cap-giac-mo-thumb.jpg','https://www.youtube.com/watch?v=YoHD9XEInc0',9,9,1,'2024-10-21 09:51:55.195','2024-12-09 11:46:05.252',NULL,2),(35,'Vampyres',2015,1,'Victor Matellano directs this tale set in a stately English manor inhabited by two older female vampires and with their only cohabitant being a man imprisoned in the basement. Their lives and lifestyle are upended when a trio of campers come upon their lair and seek to uncover their dark secrets, a decision that has sexual and blood-curdling consequences.\n\n','spanish',NULL,79,'https://img.ophim.live/uploads/movies/vampyres-thumb.jpg','https://www.youtube.com/watch?v=ZinTJFP9YDY',9,7,1,'2024-11-03 10:19:08.728','2024-11-11 10:36:54.161',NULL,3),(36,'The Signal',2024,4,'The disappearance of an astronaut sends her family on a frantic search for answers. But the more they uncover, the greater the threat to them and the world.','German',NULL,73,'https://img.ophim.live/uploads/movies/tin-hieu-bi-mat-tu-khong-gian-thumb.jpg','https://www.youtube.com/watch?v=lcqXfNzoICc',7,6,1,'2024-11-03 10:29:21.181','2024-12-09 09:50:50.923',NULL,3),(37,'The Middleman\'s Love (UNCUT)',2023,8,'The story explores what happens when the perpetual middleman finally becomes someone\'s first choice.','Thai',NULL,49,'https://img.ophim.live/uploads/movies/anh-jade-nguoi-trung-gian-thumb.jpg','https://www.youtube.com/watch?v=MwUCV0ybEK4',7,3,1,'2024-11-03 10:33:53.742','2025-02-11 20:37:59.703',NULL,19),(38,'Spice Up Our Love',2024,2,'An unpredictable fantasy romance drama about a 19+ web novel author, Nam Ja Yeon, who possesses the female protagonist of her own novel, Seo Yeon Seo, and develops a romantic relationship with the male protagonist of her romance novel, Kang Ha Jun. Ja Yeon possesses Seo Yeon Seo but is horrified when she discovers that the perfect man, Kang Ha Jun in the 19+ web novel, is not her muse, Yeo Ha Jun, but Bok Gyu Hyeon of \"No Gain No Love\". Kang Ha Jun is the perfect male protagonist in the novel written by Nam Ja Yeon. Kang Ha Jun is the president of Seo Yeon Seo\'s company, GB Electric, in the novel. He appears cold but is extremely affectionate towards the woman he loves. He is the perfect chaebol man who meets Yeon Seo and learns about love. He shows interesting chemistry as he deals with her sudden change overnight and their unrequited romance.','Korean',NULL,60,'https://img.ophim.live/uploads/movies/gia-vi-tinh-yeu-thumb.jpg','https://www.youtube.com/watch?v=7_lIdB3tcpU',9,6,1,'2024-11-03 10:43:39.481','2024-12-09 09:53:15.893',NULL,6),(39,'An Invisible Victim: The Eliza Samudio Case',2024,1,'A star goalkeeper threatens a woman who is pregnant with his child. Her pleas for help go unanswered in the shadow of his fame — then tragedy strikes.\n\n','Brazilian',NULL,101,'https://img.ophim.live/uploads/movies/a-vitima-invisivel-o-caso-eliza-samudio-thumb.jpg','https://www.youtube.com/watch?v=6Ss5z1e5Zx4',0,7,1,'2024-11-03 10:46:14.631','2024-12-09 12:01:13.898',NULL,2),(40,'Forever',2023,1,'The once unbreakable bond between teenage soccer fanatics Mila and Kia begins to strain when a tough former pro becomes their new coach.','Swedish',NULL,103,'https://img.ophim.live/uploads/movies/mai-mai-thumb.jpg','https://www.youtube.com/watch?v=LitLkEX6pd0',9,8,1,'2024-11-03 10:49:08.126','2024-11-07 15:52:21.124',NULL,1),(41,'Sleeping Beauty',2011,1,'A humorous portrait of Lucy, a young college student drawn into a mysterious hidden world of unspoken desires.','English',NULL,79,'https://img.ophim.live/uploads/movies/nguoi-dep-ngu-me-thumb.jpg','https://www.youtube.com/watch?v=b4ks-1lNrxg',0,7,1,'2024-11-03 10:51:53.430','2024-11-08 10:20:35.383',NULL,1),(42,'Good Will Hunting',1997,1,'Will Hunt, a janitor at M.I.T., has a gift for math, but needs help from a psychologist to find direction in his life.','English',NULL,120,'https://img.ophim.live/uploads/movies/good-will-hunting-thumb.jpg','https://www.youtube.com/watch?v=ReIJ1lbL-Q8',9,0,1,'2024-11-08 09:48:14.441','2024-12-09 09:38:34.848',NULL,8),(43,'White Bird',2023,1,'After being expelled from Beecher Prep for his treatment of a classmate with a facial deformity, Julian has struggled to fit in at his new school. To transform his life, Julian\'s grandmother finally reveals her own story of courage of her youth in Nazi-occupied France, where a classmate shelters her from mortal danger.\n\n','English',1,121,'https://img.ophim.live/uploads/movies/phep-mau-giua-dem-dong-thumb.jpg','https://www.youtube.com/watch?v=MOi0b6pk3c8',9,6.5,1,'2024-11-14 08:37:46.048','2024-12-09 09:39:02.507',NULL,2),(44,'Blood And Oil',2019,1,'Blood and Oil When commercial oil was discovered in the small village of Oloibiri, Nigeria, the village prospered and jobs increased. Two decades after the last drops of oil were extracted and the American oil companies abandoned them, all that remained of the village were polluted water, sickly children and an economy in crisis. When a new multinational corporation once again wanted to invest in oil exploration, the villagers\' anger grew and Nigerian militants, led by Gunpowder, sought revenge on the greedy foreigners by kidnapping Powell, the company\'s CEO. Families in the village were torn apart as village elder Timpriye struggled to save the life of the man who represented all the evils that oil had brought to his village years earlier.','English',1,91,'https://img.ophim.live/uploads/movies/mau-va-dau-thumb.jpg','https://www.youtube.com/watch?v=-6qn4k4YDEs',0,6.62,1,'2024-12-09 09:44:47.550','2024-12-09 09:45:47.919',NULL,1);

--
-- Table structure for table `rates`
--

DROP TABLE IF EXISTS `rates`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `rates` (
                         `user_id` bigint NOT NULL,
                         `movie_id` bigint NOT NULL,
                         `rate` bigint NOT NULL,
                         `created_at` datetime(3) DEFAULT NULL,
                         `updated_at` datetime(3) DEFAULT NULL,
                         PRIMARY KEY (`user_id`,`movie_id`),
                         KEY `fk_rates_movie` (`movie_id`),
                         CONSTRAINT `fk_rates_movie` FOREIGN KEY (`movie_id`) REFERENCES `movies` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
                         CONSTRAINT `fk_rates_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rates`
--

INSERT INTO `rates` VALUES (1,6,9,'2025-02-11 20:37:17.768','2025-02-11 20:37:17.768'),(7,1,10,'2024-10-28 20:07:28.594','2024-10-29 14:56:27.076'),(7,3,8,'2024-12-09 09:49:22.353','2024-12-09 09:49:22.353'),(7,4,9,'2024-11-03 10:05:37.022','2024-11-03 10:05:37.022'),(7,6,10,'2024-12-09 09:50:10.318','2024-12-09 09:50:11.913'),(7,7,10,'2024-12-09 09:50:40.387','2024-12-09 09:50:40.387'),(7,9,9,'2024-11-16 21:41:39.851','2024-11-16 21:41:39.851'),(7,14,9,'2024-10-29 15:23:41.100','2024-10-29 15:23:41.100'),(7,35,9,'2024-11-03 10:24:58.626','2024-11-03 10:24:58.626'),(7,36,7,'2024-12-09 09:50:50.910','2024-12-09 09:50:50.910'),(7,37,7,'2024-12-09 09:49:42.298','2024-12-09 09:49:42.298'),(7,40,9,'2024-11-07 15:52:21.122','2024-11-07 15:52:21.122'),(7,43,9,'2024-11-14 08:39:00.040','2024-11-14 08:39:00.040'),(13,2,9,'2024-12-09 09:52:55.431','2024-12-09 09:52:55.431'),(13,11,9,'2024-12-09 09:52:06.694','2024-12-09 09:52:06.694'),(13,38,9,'2024-12-09 09:53:15.879','2024-12-09 09:53:15.879'),(13,42,9,'2024-12-09 09:38:34.835','2024-12-09 09:38:34.835');

--
-- Table structure for table `roles`
--

DROP TABLE IF EXISTS `roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `roles` (
                         `id` bigint NOT NULL AUTO_INCREMENT,
                         `name` varchar(191) DEFAULT NULL,
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `uni_roles_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `roles`
--

INSERT INTO `roles` VALUES (2,'admin'),(1,'member');

--
-- Table structure for table `user_watcheds`
--

DROP TABLE IF EXISTS `user_watcheds`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_watcheds` (
                                 `id` bigint NOT NULL AUTO_INCREMENT,
                                 `user_id` bigint NOT NULL,
                                 `episode_id` bigint NOT NULL,
                                 `last_position` bigint NOT NULL DEFAULT '0',
                                 `updated_at` datetime(3) DEFAULT NULL,
                                 PRIMARY KEY (`id`),
                                 KEY `fk_user_watcheds_user` (`user_id`),
                                 KEY `fk_user_watcheds_episode` (`episode_id`),
                                 CONSTRAINT `fk_user_watcheds_episode` FOREIGN KEY (`episode_id`) REFERENCES `episodes` (`id`),
                                 CONSTRAINT `fk_user_watcheds_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=63 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_watcheds`
--

INSERT INTO `user_watcheds` VALUES (22,7,51,6228,'2024-11-16 21:04:48.997'),(23,7,53,34,'2024-12-09 09:49:23.683'),(24,7,63,17,'2024-10-29 15:23:53.934'),(25,7,54,1055,'2024-11-03 10:05:46.002'),(26,7,64,4,'2024-11-03 10:25:18.947'),(27,7,68,1799,'2024-12-09 09:49:44.705'),(28,7,56,6,'2024-12-09 09:50:13.578'),(29,7,84,2020,'2024-11-07 15:52:48.360'),(30,7,82,6,'2024-11-08 10:32:17.785'),(31,7,81,377,'2024-11-07 15:53:31.231'),(32,7,57,4,'2024-12-09 09:50:42.480'),(33,7,83,2,'2024-11-07 15:57:33.742'),(34,7,67,0,'2024-12-09 09:50:55.057'),(35,7,85,2,'2024-11-08 10:20:39.442'),(36,7,86,6703,'2024-11-16 21:03:16.408'),(37,7,55,313,'2024-11-08 15:07:56.270'),(38,7,69,1483,'2024-11-14 08:30:32.654'),(39,7,70,0,'2024-11-08 15:09:00.466'),(40,7,71,2030,'2024-11-08 15:10:11.165'),(42,1,64,2371,'2024-11-11 10:37:00.814'),(45,7,52,6058,'2024-11-14 08:30:13.796'),(46,7,87,9,'2024-11-14 08:39:06.746'),(47,7,59,10,'2024-11-16 21:41:48.862'),(48,13,86,12,'2024-12-09 09:38:37.560'),(49,13,87,13,'2024-12-09 09:39:17.727'),(50,13,88,1568,'2024-12-09 09:45:54.561'),(51,13,61,6,'2024-12-09 09:52:11.092'),(52,13,52,3421,'2024-12-09 09:53:00.418'),(53,13,82,4,'2024-12-09 09:53:17.947'),(54,13,63,0,'2024-12-09 11:46:15.141'),(55,13,53,0,'2024-12-09 11:46:23.833'),(56,13,51,0,'2024-12-09 11:46:32.922'),(57,13,56,0,'2024-12-09 11:49:24.316'),(58,13,83,0,'2024-12-09 12:01:24.376'),(59,13,60,0,'2024-12-09 12:06:23.356'),(60,1,56,9963,'2025-02-11 20:37:34.766'),(61,1,68,689,'2025-02-11 20:37:52.716'),(62,1,69,4,'2025-02-11 20:37:59.270');

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
                         `id` bigint NOT NULL AUTO_INCREMENT,
                         `first_name` longtext,
                         `last_name` longtext,
                         `email` longtext,
                         `password` longtext,
                         `created_at` datetime(3) DEFAULT NULL,
                         `deleted_at` datetime(3) DEFAULT NULL,
                         `role_id` bigint DEFAULT NULL,
                         PRIMARY KEY (`id`),
                         KEY `idx_users_deleted_at` (`deleted_at`),
                         KEY `fk_users_role` (`role_id`),
                         CONSTRAINT `fk_users_role` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

INSERT INTO `users` VALUES (1,'Trung Quan','Vo','quanbinskt27@gmail.com','$2a$10$Y9PUJrDu76hCs5eAQL5sDO0fYBKpqBxk3gR2u0GMgCAF4vOhGlNCO','2024-10-18 08:40:16.044',NULL,2),(7,'Trung Quan','Vo','admin@gmail.com','$2a$10$yFTGazH.SJX8jCUcDA6kwOUDF4/bcnq.D5gwUn6ifcAiRIZjN11Im','2024-10-28 14:47:07.134',NULL,2),(12,'News','Bitcoin','bkfim27@gmail.com','','2024-11-11 11:00:50.974',NULL,1),(13,'Nguyễn Ngọc','Lâm 9A1','binkinacc2@gmail.com','','2024-12-09 09:37:23.421',NULL,2);


-- Dump completed on 2025-02-13  8:49:48
