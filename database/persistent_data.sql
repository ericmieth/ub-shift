INSERT INTO `weekday` (id, name) VALUES 
(0, 'Sonntag'),
(1, 'Montag'),
(2, 'Dienstag'),
(3, 'Mittwoch'),
(4, 'Donnerstag'),
(5, 'Freitag'),
(6, 'Samstag');


INSERT INTO `branch` (name,location) VALUES
('Albertina', 'Beethovenstraße 6, 04107 Leipzig'),
('Biowissenschaften', 'Talstraße 35, 04103 Leipzig'),
('Campus-Bibliothek', 'Universitätsstraße 3, 04109 Leipzig'),
('Chemie/ Physik', 'Johannisallee 29, 04103 Leipzig'),
('Deutsches Literaturinstitut', 'Wächterstraße 34, 04107 Leipzig'),
('Geographie', 'Johannisallee 19, 04103 Leipzig'),
('Geowissenschaften', 'Talstraße 35, 04103 Leipzig'),
('Klassische Archäologie und Ur- und Frühgeschichte', 'Ritterstraße 14, 04109 Leipzig'),
('Kunst', 'Dittrichring 18–20, 04109 Leipzig'),
('Zentralbibliothek Medizin', 'Johannisalle 34, Haus L, 04103 Leipzig'),
('Musik', 'Neumarkt 9–19 – Aufgang D, Städtisches Kaufhaus, 04109 Leipzig'),
('Orientwissenschaften', 'Schillerstraße 6, 04109 Leipzig'),
('Rechtswissenschaften', 'Burgstraße 27, 04109 Leipzig'),
('Sportwissenschaften', 'Jahnallee 59, 04109 Leipzig'),
('Verterinärmedizin', 'An den Tierkliniken 5, 04103 Leipzig');

INSERT INTO `counter` (branch_id,name) VALUES
('1', 'BIN-Raum (EG)'),
('1', 'Servicetheke (2.OG)');
