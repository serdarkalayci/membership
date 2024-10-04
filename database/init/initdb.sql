CREATE DATABASE membership;
CREATE USER membershipuser WITH PASSWORD 'membershippassword';
GRANT ALL PRIVILEGES ON DATABASE membership TO membershipuser;
ALTER DATABASE membership OWNER TO membershipuser;
\c membership membershipuser;

-- Creating tables
CREATE TABLE public.provinces (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE public.cities (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    province_id integer NOT NULL,
    CONSTRAINT cities_province FOREIGN KEY (province_id) REFERENCES public.provinces(id)
);

CREATE TABLE public.areas (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE public.membership_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE public.members (
    id UUID NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NULL,
    city_id integer NOT NULL,
    area_id integer NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    membership_type_id integer NOT NULL,
    membership_start_date DATE NULL,
    last_contact_date DATE NULL,
    occupation VARCHAR(255) NULL,
    education VARCHAR(255) NULL,
    date_of_birth DATE NULL,
    passive BOOL NOT NULL DEFAULT false,
    notes TEXT NULL,
    CONSTRAINT members_pkey PRIMARY KEY (id),
    CONSTRAINT members_city FOREIGN KEY (city_id) REFERENCES public.cities(id),
    CONSTRAINT members_area FOREIGN KEY (area_id) REFERENCES public.areas(id),
    CONSTRAINT members_membership_type FOREIGN KEY (membership_type_id) REFERENCES public.membership_types(id)
);

CREATE TABLE public.users (
    id UUID NOT NULL,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT users_username UNIQUE (username)
);

CREATE TABLE public.roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE public.user_roles (
    user_id UUID NOT NULL,
    role_id integer NOT NULL,
    CONSTRAINT user_roles_user FOREIGN KEY (user_id) REFERENCES public.users(id),
    CONSTRAINT user_roles_role FOREIGN KEY (role_id) REFERENCES public.roles(id)
);

-- Inserting Roles
INSERT INTO public.roles ("id", "name") VALUES
(1, 'Admin'),
(2, 'User');

-- Inserting Users
INSERT INTO public.users ("id", "username", "password", "email") VALUES
('00000000-0000-0000-0000-000000000001', 'admin', '$2a$10$YI0rRpOKFs0/mBI1ixP3pukU8hb5cRgcB5478kWbQRPfH9MSw5UJS', 'admin@admin.com'),
('00000000-0000-0000-0000-000000000002', 'user', '$2a$10$WZuiIeib85hrVQt5ChZ63eIxo.O.Ws9YE.hQ3tKU93ugnIGks/Fsu', 'user@user.com');

-- Inserting User Roles
INSERT INTO public.user_roles ("user_id", "role_id") VALUES
('00000000-0000-0000-0000-000000000001', 1),
('00000000-0000-0000-0000-000000000002', 2);

-- Inserting Provinces
INSERT INTO public.provinces ("id", "name") VALUES
(1, 'Drenthe'),
(2, 'Flevoland'),
(3, 'Friesland'),
(4, 'Gelderland'),
(5, 'Groningen'),
(6, 'Limburg'),
(7, 'Noord-Brabant'),
(8, 'Noord-Holland'),
(9, 'Overijssel'),
(10, 'Utrecht'),
(11, 'Zeeland'),
(12, 'ZuID-Holland');

-- Inserting Cities
INSERT INTO public.cities ("id", "province_id", "name") VALUES
(1, 1, 'Emmen'), -- Drenthe
(2, 1, 'Assen'), -- Drenthe
(3, 1, 'Hoogeveen'), -- Drenthe

(4, 2, 'Almere'), -- Flevoland
(5, 2, 'Emmeloord'), -- Flevoland
(6, 2, 'Dronten'), -- Flevoland

(7, 3, 'Sneek'), -- Friesland
(8, 3, 'Heerenveen'), -- Friesland
(9, 3, 'Harlingen'), -- Friesland

(10, 4, 'Nijmegen'), -- Gelderland
(11, 4, 'Apeldoorn'), -- Gelderland
(12, 4, 'Doetinchem'), -- Gelderland

(13, 5, 'Groningen'), -- Groningen
(14, 5, 'Winschoten'), -- Groningen
(15, 5, 'Veendam'), -- Groningen

(16, 6, 'Venlo'), -- Limburg
(17, 6, 'Sittard'), -- Limburg
(18, 6, 'Roermond'), -- Limburg

(19, 7, 'Eindhoven'), -- Noord-Brabant
(20, 7, 'Tilburg'), -- Noord-Brabant
(21, 7, 'Breda'), -- Noord-Brabant

(22, 8, 'Amsterdam'), -- Noord-Holland
(23, 8, 'Haarlem'), -- Noord-Holland
(24, 8, 'Hilversum'), -- Noord-Holland

(25, 9, 'Enschede'), -- Overijssel
(26, 9, 'Deventer'), -- Overijssel
(27, 9, 'Hengelo'), -- Overijssel

(28, 10, 'Amersfoort'), -- Utrecht
(29, 10, 'Nieuwegein'), -- Utrecht
(30, 10, 'Soest'), -- Utrecht

(31, 11, 'Vlissingen'), -- Zeeland
(32, 11, 'Goes'), -- Zeeland
(33, 11, 'Terneuzen'), -- Zeeland

(34, 12, 'Rotterdam'), -- ZuID-Holland
(35, 12, 'The Hague'), -- ZuID-Holland
(36, 12, 'Dordrecht'); -- ZuID-Holland

-- Inserting Areas
INSERT INTO public.areas ("id", "name") VALUES
(1, 'Eindhoven'),
(2, 'Amsterdam'),
(3, 'Rotterdam'),
(4, 'Utrecht');

-- Inserting Membership types
INSERT INTO public.membership_types ("id", "name") VALUES
(1, 'Kadro'),
(2, 'Aktif'),
(3, 'Sessiz'),
(4, 'Üye'),
(5, 'Potansiyel');

-- Inserting Members
INSERT INTO public.members ("id", "first_name", "last_name", "passive", "email", "phone", "city_id", "area_id", "membership_type_id", "membership_start_date", "last_contact_date", "occupation", "education", "date_of_birth") VALUES
('00000000-0000-0000-0000-000000000001', 'Oliver', 'Smith', false, 'oliver.smith@example.com', '123456789', 1, 1, 1, '2022-01-01', '2024-05-13', 'Engineer', 'Bachelor of Engineering', '1980-01-01'),
('00000000-0000-0000-0000-000000000002', 'Emma', 'Johnson', true, 'emma.johnson@example.com', '234567890', 2, 2, 2, '2022-01-02', '2024-05-12', 'Doctor', 'Doctor of Medicine', '1985-02-02'),
('00000000-0000-0000-0000-000000000003', 'William', 'Williams', false, 'william.williams@example.com', '345678901', 3, 3, 3, '2022-01-03', '2024-05-11', 'Teacher', 'Bachelor of Education', '1990-03-03'),
('00000000-0000-0000-0000-000000000004', 'Olivia', 'Brown', false, 'olivia.brown@example.com', '456789012', 4, 4, 4, '2022-01-04', '2024-05-10', 'Nurse', 'Bachelor of Nursing', '1995-04-04'),
('00000000-0000-0000-0000-000000000005', 'James', 'Jones', true, 'james.jones@example.com', '567890123', 5, 1, 5, '2022-01-05', '2024-05-09', 'Lawyer', 'Juris Doctor', '1980-05-05'),
('00000000-0000-0000-0000-000000000006', 'Sophia', 'Garcia', false, 'sophia.garcia@example.com', '678901234', 6, 2, 1, '2022-01-06', '2024-05-08', 'Accountant', 'Bachelor of Accounting', '1985-06-06'),
('00000000-0000-0000-0000-000000000007', 'Benjamin', 'Miller', false, 'benjamin.miller@example.com', '789012345', 7, 3, 2, '2022-01-07', '2024-05-07', 'Artist', 'Bachelor of Fine Arts', '1990-07-07'),
('00000000-0000-0000-0000-000000000008', 'Isabella', 'Davis', true, 'isabella.davis@example.com', '890123456', 8, 4, 3, '2022-01-08', '2024-05-06', 'Chef', 'Culinary Arts Degree', '1995-08-08'),
('00000000-0000-0000-0000-000000000009', 'Mason', 'Martinez', false, 'mason.martinez@example.com', '901234567', 9, 1, 4, '2022-01-09', '2024-05-05', 'Pilot', 'Commercial Pilot License', '1980-09-09'),
('00000000-0000-0000-0000-000000000010', 'Charlotte', 'Hernandez', false, 'charlotte.hernandez@example.com', '012345678', 10, 2, 5, '2022-01-10', '2024-05-04', 'Entrepreneur', 'Bachelor of Business Administration', '1985-10-10'),
('00000000-0000-0000-0000-000000000011', 'Liam', 'Lopez', true, 'liam.lopez@example.com', '987654221', 11, 3, 1, '2022-01-11', '2024-05-03', 'Engineer', 'Bachelor of Engineering', '1990-11-11'),
('00000000-0000-0000-0000-000000000012', 'Amelia', 'Scott', false, 'amelia.scott@example.com', '876546210', 12, 4, 2, '2022-01-12', '2024-05-02', 'Professor', 'Doctor of Philosophy', '1995-12-12'),
('00000000-0000-0000-0000-000000000013', 'Ethan', 'Green', false, 'ethan.green@example.com', '765132109', 13, 1, 3, '2022-01-13', '2024-05-01', 'Manager', 'Bachelor of Business Administration', '1981-01-13'),
('00000000-0000-0000-0000-000000000014', 'Ava', 'Adams', true, 'ava.adams@example.com', '654521098', 14, 2, 4, '2022-01-14', '2024-04-30', 'Writer', 'Bachelor of Arts', '1986-02-14'),
('00000000-0000-0000-0000-000000000015', 'Nolah', 'Campbell', false, 'nolah.campbell@example.com', '523210987', 15, 3, 5, '2022-01-15', '2024-04-29', 'Artist', 'Bachelor of Fine Arts', '1991-03-15'),
('00000000-0000-0000-0000-000000000016', 'Harper', 'Nelson', false, 'harper.nelson@example.com', '431109876', 16, 4, 1, '2022-01-16', '2024-04-28', 'Chef', 'Culinary Arts Degree', '1986-04-16'),
('00000000-0000-0000-0000-000000000017', 'Logan', 'Carter', true, 'logan.carter@example.com', '321018765', 17, 1, 2, '2022-01-17', '2024-04-27', 'Pilot', 'Commercial Pilot License', '1991-05-17'),
('00000000-0000-0000-0000-000000000018', 'Evelyn', 'Mitchell', false, 'evelyn.mitchell@example.com', '211987654', 18, 2, 3, '2022-01-18', '2024-04-26', 'Entrepreneur', 'Bachelor of Business Administration', '1986-06-18'),
('00000000-0000-0000-0000-000000000019', 'Lucas', 'Perez', false, 'lucas.perez@example.com', '119876543', 19, 3, 4, '2022-01-19', '2024-04-25', 'Engineer', 'Bachelor of Engineering', '1991-07-19'),
('00000000-0000-0000-0000-000000000020', 'Avery', 'Roberts', true, 'avery.roberts@example.com', '098765431', 20, 4, 5, '2022-01-20', '2024-04-24', 'Professor', 'Doctor of Philosophy', '1986-08-20'),
('00000000-0000-0000-0000-000000000021', 'Carter', 'Turner', false, 'carter.turner@example.com', '983654321', 21, 1, 1, '2022-01-21', '2024-04-23', 'Manager', 'Bachelor of Business Administration', '1991-09-21'),
('00000000-0000-0000-0000-000000000022', 'Scarlett', 'Phillips', false, 'scarlett.phillips@example.com', '878543210', 22, 2, 2, '2022-01-22', '2024-04-22', 'Writer', 'Bachelor of Arts', '1986-10-22'),
('00000000-0000-0000-0000-000000000023', 'Jackson', 'Campbell', true, 'jackson.campbell@example.com', '765432809', 23, 3, 3, '2022-01-23', '2024-04-21', 'Artist', 'Bachelor of Fine Arts', '1991-11-23'),
('00000000-0000-0000-0000-000000000024', 'Chloe', 'Nelson', false, 'chloe.nelson@example.com', '653321098', 24, 4, 4, '2022-01-24', '2024-04-20', 'Chef', 'Culinary Arts Degree', '1986-12-24'),
('00000000-0000-0000-0000-000000000025', 'Lincoln', 'Carter', false, 'lincoln.carter@example.com', '541213987', 25, 1, 5, '2022-01-25', '2024-04-19', 'Pilot', 'Commercial Pilot License', '1991-01-25'),
('00000000-0000-0000-0000-000000000026', 'Penelope', 'Mitchell', true, 'penelope.mitchell@example.com', '432301876', 26, 2, 1, '2022-01-26', '2024-04-18', 'Entrepreneur', 'Bachelor of Business Administration', '1986-02-26'),
('00000000-0000-0000-0000-000000000027', 'Grayson', 'Perez', false, 'grayson.perez@example.com', '321098715', 1, 3, 2, '2022-01-27', '2024-04-17', 'Engineer', 'Bachelor of Engineering', '1991-03-27'),
('00000000-0000-0000-0000-000000000028', 'Luna', 'Roberts', false, 'luna.roberts@example.com', '212987654', 2, 4, 3, '2022-01-28', '2024-04-16', 'Professor', 'Doctor of Philosophy', '1986-04-28'),
('00000000-0000-0000-0000-000000000029', 'Miles', 'Turner', true, 'miles.turner@example.com', '129876543', 3, 1, 4, '2022-01-29', '2024-04-15', 'Manager', 'Bachelor of Business Administration', '1991-05-29'),
('00000000-0000-0000-0000-000000000030', 'Stella', 'Phillips', false, 'stella.phillips@example.com', '098365432', 4, 2, 5, '2022-01-30', '2024-04-14', 'Writer', 'Bachelor of Arts', '1986-06-30'),
('00000000-0000-0000-0000-000000000031', 'Elijah', 'Campbell', false, 'elijah.campbell@example.com', '987656321', 5, 3, 1, '2022-01-31', '2024-04-13', 'Artist', 'Bachelor of Fine Arts', '1991-07-31'),
('00000000-0000-0000-0000-000000000032', 'Hazel', 'Nelson', true, 'hazel.nelson@example.com', '876533210', 6, 4, 2, '2022-02-01', '2024-04-12', 'Chef', 'Culinary Arts Degree', '1986-08-01'),
('00000000-0000-0000-0000-000000000033', 'Jacob', 'Carter', false, 'jacob.carter@example.com', '765932109', 7, 1, 3, '2022-02-02', '2024-04-11', 'Pilot', 'Commercial Pilot License', '1991-09-02'),
('00000000-0000-0000-0000-000000000034', 'Nora', 'Mitchell', false, 'nora.mitchell@example.com', '654323098', 8, 2, 4, '2022-02-03', '2024-04-10', 'Entrepreneur', 'Bachelor of Business Administration', '1986-10-03'),
('00000000-0000-0000-0000-000000000035', 'Sebastian', 'Perez', true, 'sebastian.perez@example.com', '543210187', 9, 3, 5, '2022-02-04', '2024-04-09', 'Engineer', 'Bachelor of Engineering', '1991-11-04'),
('00000000-0000-0000-0000-000000000036', 'Ella', 'Roberts', false, 'ella.roberts@example.com', '432109816', 10, 4, 1, '2022-02-05', '2024-04-08', 'Professor', 'Doctor of Philosophy', '1986-12-05'),
('00000000-0000-0000-0000-000000000037', 'Caleb', 'Turner', false, 'caleb.turner@example.com', '321098725', 11, 1, 2, '2022-02-06', '2024-04-07', 'Manager', 'Bachelor of Business Administration', '1991-01-06'),
('00000000-0000-0000-0000-000000000038', 'Mila', 'Phillips', true, 'mila.phillips@example.com', '213987654', 12, 2, 3, '2022-02-07', '2024-04-06', 'Writer', 'Bachelor of Arts', '1986-02-07'),
('00000000-0000-0000-0000-000000000039', 'Eli', 'Campbell', false, 'eli.campbell@example.com', '139876543', 13, 3, 4, '2022-02-08', '2024-04-05', 'Artist', 'Bachelor of Fine Arts', '1991-03-08'),
('00000000-0000-0000-0000-000000000040', 'Layla', 'Nelson', false, 'layla.nelson@example.com', '098765432', 14, 4, 5, '2022-02-09', '2024-04-04', 'Chef', 'Culinary Arts Degree', '1986-04-09'),
('00000000-0000-0000-0000-000000000041', 'Brayden', 'Carter', true, 'brayden.carter@example.com', '987654391', 15, 1, 1, '2022-02-10', '2024-04-03', 'Pilot', 'Commercial Pilot License', '1991-07-10'),
('00000000-0000-0000-0000-000000000042', 'Aurora', 'Mitchell', false, 'aurora.mitchell@example.com', '876143210', 16, 2, 2, '2022-02-11', '2024-04-02', 'Entrepreneur', 'Bachelor of Business Administration', '1986-08-11'),
('00000000-0000-0000-0000-000000000043', 'Levi', 'Perez', false, 'levi.perez@example.com', '765462109', 17, 3, 3, '2022-02-12', '2024-04-01', 'Engineer', 'Bachelor of Engineering', '1991-09-12'),
('00000000-0000-0000-0000-000000000044', 'Nova', 'Roeberts', true, 'nova.roeberts@example.com', '654321038', 18, 4, 4, '2022-02-13', '2024-03-31', 'Professor', 'Doctor of Philosophy', '1986-10-13'),
('00000000-0000-0000-0000-000000000045', 'Hunter', 'Turner', false, 'hunter.turner@example.com', '543217987', 19, 1, 5, '2022-02-14', '2024-03-30', 'Manager', 'Bachelor of Business Administration', '1991-11-14'),
('00000000-0000-0000-0000-000000000046', 'Emilia', 'Phillips', false, 'emilia.phillips@example.com', '432909876', 20, 2, 1, '2022-02-15', '2024-03-29', 'Writer', 'Bachelor of Arts', '1986-02-15'),
('00000000-0000-0000-0000-000000000047', 'Kai', 'Campbell', true, 'kai.campbell@example.com', '321098735', 21, 3, 2, '2022-02-16', '2024-03-28', 'Artist', 'Bachelor of Fine Arts', '1991-03-16'),
('00000000-0000-0000-0000-000000000048', 'Lila', 'Nelson', false, 'lila.nelson@example.com', '214987654', 22, 4, 3, '2022-02-17', '2024-03-27', 'Chef', 'Culinary Arts Degree', '1986-04-17'),
('00000000-0000-0000-0000-000000000049', 'Finn', 'Carter', false, 'finn.carter@example.com', '149876543', 23, 1, 4, '2022-02-18', '2024-03-26', 'Pilot', 'Commercial Pilot License', '1991-05-18'),
('00000000-0000-0000-0000-000000000050', 'Zoe', 'Mitchell', true, 'zoe.mitchell@example.com', '095765432', 24, 2, 5, '2022-02-19', '2024-03-25', 'Entrepreneur', 'Bachelor of Business Administration', '1986-06-19'),
('00000000-0000-0000-0000-000000000051', 'Axel', 'Perez', false, 'axel.perez@example.com', '987651321', 25, 3, 1, '2022-02-20', '2024-03-24', 'Engineer', 'Bachelor of Engineering', '1991-07-20'),
('00000000-0000-0000-0000-000000000052', 'Nova', 'Roberts', false, 'nova.roberts@example.com', '876943210', 26, 4, 2, '2022-02-21', '2024-03-23', 'Professor', 'Doctor of Philosophy', '1986-08-21'),
('00000000-0000-0000-0000-000000000053', 'Aria', 'Turner', true, 'aria.turner@example.com', '765431109', 1, 1, 3, '2022-02-22', '2024-03-22', 'Manager', 'Bachelor of Business Administration', '1991-09-22'),
('00000000-0000-0000-0000-000000000054', 'Zachary', 'Phillips', false, 'zachary.phillips@example.com', '694321098', 2, 2, 4, '2022-02-23', '2024-03-21', 'Writer', 'Bachelor of Arts', '1986-10-23'),
('00000000-0000-0000-0000-000000000055', 'Ayla', 'Campbell', false, 'ayla.campbell@example.com', '503210987', 3, 3, 5, '2022-02-24', '2024-03-20', 'Artist', 'Bachelor of Fine Arts', '1991-11-24'),
('00000000-0000-0000-0000-000000000056', 'Ezra', 'Nelson', true, 'ezra.nelson@example.com', '432139876', 4, 4, 1, '2022-02-25', '2024-03-19', 'Chef', 'Culinary Arts Degree', '1986-12-25'),
('00000000-0000-0000-0000-000000000057', 'Natalie', 'Carter', false, 'natalie.carter@example.com', '321098745', 5, 1, 2, '2022-02-26', '2024-03-18', 'Pilot', 'Commercial Pilot License', '1991-01-26'),
('00000000-0000-0000-0000-000000000058', 'Asher', 'Mitchell', false, 'asher.mitchell@example.com', '215987654', 6, 2, 3, '2022-02-27', '2024-03-17', 'Entrepreneur', 'Bachelor of Business Administration', '1986-02-27'),
('00000000-0000-0000-0000-000000000059', 'Scarlett', 'Perez', true, 'scarlett.perez@example.com', '159876543', 7, 3, 4, '2022-02-28', '2024-03-16', 'Engineer', 'Bachelor of Engineering', '1991-05-28'),
('00000000-0000-0000-0000-000000000060', 'Leo', 'Roberts', false, 'leo.roberts@example.com', '098665432', 8, 4, 5, '2022-03-01', '2024-03-15', 'Professor', 'Doctor of Philosophy', '1986-06-01'),
('00000000-0000-0000-0000-000000000061', 'Aurora', 'Turner', false, 'aurora.turner@example.com', '987154321', 9, 1, 1, '2022-03-02', '2024-03-14', 'Manager', 'Bachelor of Business Administration', '1991-07-02'),
('00000000-0000-0000-0000-000000000062', 'Eli', 'Phillips', true, 'eli.phillips@example.com', '876543810', 10, 2, 2, '2022-03-03', '2024-03-13', 'Writer', 'Bachelor of Arts', '1986-08-03'),
('00000000-0000-0000-0000-000000000063', 'Nora', 'Campbell', false, 'nora.campbell@example.com', '765832109', 11, 3, 3, '2022-03-04', '2024-03-12', 'Artist', 'Bachelor of Fine Arts', '1991-09-04'),
('00000000-0000-0000-0000-000000000064', 'Hudson', 'Nelson', false, 'hudson.nelson@example.com', '651321098', 12, 4, 4, '2022-03-05', '2024-03-11', 'Chef', 'Culinary Arts Degree', '1986-10-05'),
('00000000-0000-0000-0000-000000000065', 'Ella', 'Carter', true, 'ella.carter@example.com', '541210987', 13, 1, 5, '2022-03-06', '2024-03-10', 'Pilot', 'Commercial Pilot License', '1991-11-06'),
('00000000-0000-0000-0000-000000000066', 'Sebastian', 'Mitchell', false, 'sebastian.mitchell@example.com', '432149876', 14, 2, 1, '2022-03-07', '2024-03-09', 'Entrepreneur', 'Bachelor of Business Administration', '1986-12-07'),
('00000000-0000-0000-0000-000000000067', 'Ava', 'Perez', false, 'ava.perez@example.com', '321098755', 15, 3, 2, '2022-03-08', '2024-03-08', 'Engineer', 'Bachelor of Engineering', '1991-01-08'),
('00000000-0000-0000-0000-000000000068', 'Niam', 'Roberts', true, 'niam.roberts@example.com', '216987654', 16, 4, 3, '2022-03-09', '2024-03-07', 'Professor', 'Doctor of Philosophy', '1986-02-09'),
('00000000-0000-0000-0000-000000000069', 'Mila', 'Turner', false, 'mila.turner@example.com', '169876543', 17, 1, 4, '2022-03-10', '2024-03-06', 'Manager', 'Bachelor of Business Administration', '1991-05-10'),
('00000000-0000-0000-0000-000000000070', 'Henry', 'Phillips', false, 'henry.phillips@example.com', '011765432', 18, 2, 5, '2022-03-11', '2024-03-05', 'Writer', 'Bachelor of Arts', '1986-06-11'),
('00000000-0000-0000-0000-000000000071', 'Nova', 'Campbell', true, 'nova.campbell@example.com', '982654321', 19, 3, 1, '2022-03-12', '2024-03-04', 'Artist', 'Bachelor of Fine Arts', '1991-07-12'),
('00000000-0000-0000-0000-000000000072', 'Emma', 'Nelson', false, 'emma.nelson@example.com', '876513210', 20, 4, 2, '2022-03-13', '2024-03-03', 'Chef', 'Culinary Arts Degree', '1986-08-13'),
('00000000-0000-0000-0000-000000000073', 'Elajah', 'Carter', false, 'elajah.carter@example.com', '767432109', 21, 1, 3, '2022-03-14', '2024-03-02', 'Pilot', 'Commercial Pilot License', '1991-09-14'),
('00000000-0000-0000-0000-000000000074', 'Olivia', 'Mitchell', true, 'olivia.mitchell@example.com', '652321098', 22, 2, 4, '2022-03-15', '2024-03-01', 'Entrepreneur', 'Bachelor of Business Administration', '1986-10-15'),
('00000000-0000-0000-0000-000000000075', 'Lucas', 'Peresz', false, 'lucas.peresz@example.com', '543210986', 23, 3, 5, '2022-03-16', '2024-02-28', 'Engineer', 'Bachelor of Engineering', '1991-11-16'),
('00000000-0000-0000-0000-000000000076', 'Charlotte', 'Roberts', false, 'charlotte.roberts@example.com', '432101876', 24, 4, 1, '2022-03-17', '2024-02-28', 'Professor', 'Doctor of Philosophy', '1986-12-17'),
('00000000-0000-0000-0000-000000000077', 'Liam', 'Turner', true, 'liam.turner@example.com', '321098775', 25, 1, 2, '2022-03-18', '2024-02-27', 'Manager', 'Bachelor of Business Administration', '1991-01-18'),
('00000000-0000-0000-0000-000000000078', 'Avery', 'Phillips', false, 'avery.phillips@example.com', '217987654', 26, 2, 3, '2022-03-19', '2024-02-26', 'Writer', 'Bachelor of Arts', '1986-02-19'),
('00000000-0000-0000-0000-000000000079', 'Ethan', 'Campbell', false, 'ethan.campbell@example.com', '179876543', 1, 3, 4, '2022-03-20', '2024-02-25', 'Artist', 'Bachelor of Fine Arts', '1991-05-20'),
('00000000-0000-0000-0000-000000000080', 'Aria', 'Nelson', true, 'aria.nelson@example.com', '018765432', 2, 4, 5, '2022-03-21', '2024-02-24', 'Chef', 'Culinary Arts Degree', '1986-06-21'),
('00000000-0000-0000-0000-000000000081', 'Elijah', 'Carter', false, 'elijah.carter@example.com', '987184321', 3, 1, 1, '2022-03-22', '2024-02-23', 'Pilot', 'Commercial Pilot License', '1991-07-22'),
('00000000-0000-0000-0000-000000000082', 'Mila', 'Mitchell', false, 'mila.mitchell@example.com', '876543910', 4, 2, 2, '2022-03-23', '2024-02-22', 'Entrepreneur', 'Bachelor of Business Administration', '1986-08-23'),
('00000000-0000-0000-0000-000000000083', 'Landon', 'Perez', true, 'landon.perez@example.com', '761432109', 5, 3, 3, '2022-03-24', '2024-02-21', 'Engineer', 'Bachelor of Engineering', '1991-09-24'),
('00000000-0000-0000-0000-000000000084', 'Aurora', 'Roberts', false, 'aurora.roberts@example.com', '658321098', 6, 4, 4, '2022-03-25', '2024-02-20', 'Professor', 'Doctor of Philosophy', '1986-10-25'),
('00000000-0000-0000-0000-000000000085', 'Grayson', 'Turner', false, 'grayson.turner@example.com', '543290987', 7, 1, 5, '2022-03-26', '2024-02-19', 'Manager', 'Bachelor of Business Administration', '1991-11-26'),
('00000000-0000-0000-0000-000000000086', 'Natalie', 'Phillips', true, 'natalie.phillips@example.com', '432109476', 8, 2, 1, '2022-03-27', '2024-02-18', 'Writer', 'Bachelor of Arts', '1986-12-27'),
('00000000-0000-0000-0000-000000000087', 'Noah', 'Campbell', false, 'noah.campbell@example.com', '321098785', 9, 3, 2, '2022-03-28', '2024-02-17', 'Artist', 'Bachelor of Fine Arts', '1991-01-28'),
('00000000-0000-0000-0000-000000000088', 'Luna', 'Nelson', false, 'luna.nelson@example.com', '218987654', 10, 4, 3, '2022-03-29', '2024-02-16', 'Chef', 'Culinary Arts Degree', '1986-02-28'),
('00000000-0000-0000-0000-000000000089', 'Liam', 'Carter', true, 'liam.carter@example.com', '189876543', 11, 1, 4, '2022-03-30', '2024-02-15', 'Pilot', 'Commercial Pilot License', '1991-05-30'),
('00000000-0000-0000-0000-000000000090', 'Avery', 'Mitchell', false, 'avery.mitchell@example.com', '091765432', 12, 2, 5, '2022-03-31', '2024-02-14', 'Entrepreneur', 'Bachelor of Business Administration', '1986-06-01'),
('00000000-0000-0000-0000-000000000091', 'Emma', 'Perez', false, 'emma.perez@example.com', '987654321', 13, 3, 1, '2022-04-01', '2024-02-13', 'Engineer', 'Bachelor of Engineering', '1991-07-01'),
('00000000-0000-0000-0000-000000000092', 'Eli', 'Roberts', true, 'eli.roberts@example.com', '876543210', 14, 4, 2, '2022-04-02', '2024-02-12', 'Professor', 'Doctor of Philosophy', '1986-08-02'),
('00000000-0000-0000-0000-000000000093', 'Mia', 'Turner', false, 'mia.turner@example.com', '765432109', 15, 1, 3, '2022-04-03', '2024-02-11', 'Manager', 'Bachelor of Business Administration', '1991-09-03'),
('00000000-0000-0000-0000-000000000094', 'Oliver', 'Phillips', false, 'oliver.phillips@example.com', '654321098', 16, 2, 4, '2022-04-04', '2024-02-10', 'Writer', 'Bachelor of Arts', '1986-10-04'),
('00000000-0000-0000-0000-000000000095', 'Amelia', 'Campbell', true, 'amelia.campbell@example.com', '543210987', 17, 3, 5, '2022-04-05', '2024-02-09', 'Artist', 'Bachelor of Fine Arts', '1991-11-05'),
('00000000-0000-0000-0000-000000000096', 'Elijah', 'Nelson', false, 'elijah.nelson@example.com', '432109876', 18, 4, 1, '2022-04-06', '2024-02-08', 'Chef', 'Culinary Arts Degree', '1986-12-06'),
('00000000-0000-0000-0000-000000000097', 'Aurora', 'Carter', false, 'aurora.carter@example.com', '321098765', 19, 1, 2, '2022-04-07', '2024-02-07', 'Pilot', 'Commercial Pilot License', '1991-01-07'),
('00000000-0000-0000-0000-000000000098', 'Ezra', 'Mitchell', true, 'ezra.mitchell@example.com', '210987654', 20, 2, 3, '2022-04-08', '2024-02-06', 'Entrepreneur', 'Bachelor of Business Administration', '1986-02-08'),
('00000000-0000-0000-0000-000000000099', 'Isabella', 'Perez', false, 'isabella.perez@example.com', '109876543', 21, 3, 4, '2022-04-09', '2024-02-05', 'Engineer', 'Bachelor of Engineering', '1991-05-09'),
('00000000-0000-0000-0000-000000000100', 'Liam', 'Roberts', false, 'liam.roberts@example.com', '098765492', 22, 4, 5, '2022-04-10', '2024-02-04', 'Professor', 'Doctor of Philosophy', '1986-06-10');

