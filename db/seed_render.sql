-- ==========================
-- Schema Definitions
-- ==========================

-- USERS TABLE
CREATE SEQUENCE IF NOT EXISTS public.users_id_seq START 1;
CREATE TABLE IF NOT EXISTS public.users (
    id BIGINT PRIMARY KEY DEFAULT nextval('public.users_id_seq'),
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    password TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ
);

-- MEETUPS TABLE
CREATE SEQUENCE IF NOT EXISTS public.meetups_id_seq START 1;
CREATE TABLE IF NOT EXISTS public.meetups (
    id BIGINT PRIMARY KEY DEFAULT nextval('public.meetups_id_seq'),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    user_id BIGINT NOT NULL REFERENCES public.users(id) ON DELETE CASCADE
);

-- MEETUP_INVITATIONS TABLE
CREATE SEQUENCE IF NOT EXISTS public.meetup_invitations_id_seq START 1;
CREATE TABLE IF NOT EXISTS public.meetup_invitations (
    id BIGINT PRIMARY KEY DEFAULT nextval('public.meetup_invitations_id_seq'),
    meetup_id BIGINT NOT NULL REFERENCES public.meetups(id) ON DELETE CASCADE,
    user_id BIGINT NOT NULL REFERENCES public.users(id) ON DELETE CASCADE,
    status TEXT NOT NULL DEFAULT 'pending'
);

-- MESSAGES TABLE
CREATE SEQUENCE IF NOT EXISTS public.messages_id_seq START 1;
CREATE TABLE IF NOT EXISTS public.messages (
    id BIGINT PRIMARY KEY DEFAULT nextval('public.messages_id_seq'),
    meetup_id BIGINT NOT NULL REFERENCES public.meetups(id) ON DELETE CASCADE,
    sender_id BIGINT NOT NULL REFERENCES public.users(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    timestamp TIMESTAMPTZ DEFAULT now()
);

-- ==========================
-- Data Insertion
-- ==========================

-- USERS
INSERT INTO public.users (id, username, email, first_name, last_name, password, created_at, updated_at, deleted_at) VALUES
(2, 'ThorEdge', 'thor@gmail.com', 'Thor', 'Odisson', '$2a$10$pfWpqBPAc2vpgtagUPVyhOFyYMZHxYqHev0HRPokp3AtduiLo76tC', '2025-02-16 18:37:53.638204+02', '2025-02-16 18:37:53.638204+02', NULL),
(4, 'Odin', 'odin@gmail.com', 'Odin', 'The Maker', '$2a$10$17dusYjXHyqEAmhYEZpusu7Ho18Ds8VUus24bfDnc7JTb.YqaL1lK', '2025-02-16 21:21:41.643996+02', '2025-02-16 21:21:41.643996+02', NULL),
(5, 'Joleyne', 'joleyne@gmail.com', 'Joleyne', 'Smith', '$2a$10$Spaunj2SP4CwOSMjJ67m5OqoOf.6YseMPta22EWxTflWD29ItsZvO', '2025-02-21 06:42:47.283544+02', '2025-02-21 06:42:47.283544+02', NULL),
(11, 'Tonkata', 'tonkata@gmail.com', 'Toncho', 'Moncho', '$2a$10$V8sXdbNNPulFgz.cnRiiqeUh1C8jXqOt2a8UTHvvvRu6Q8AA.WwCC', '2025-05-20 20:13:43.297678+03', '2025-05-20 20:13:43.297678+03', NULL);

-- MEETUPS
INSERT INTO public.meetups (id, name, description, user_id) VALUES
(1, 'Valhala meeting', 'Meeting to see old friends!!!', 5),
(3, 'Corparate Meeting', 'A very important meeting.ffffffffffffffffffytuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuu', 11),
(4, 'Hunuluu Meeting', 'A very exotic meeting', 11),
(8, 'July Meetup', 'A month meetup for discussing raised issues\n', 5);

-- INVITATIONS
INSERT INTO public.meetup_invitations (id, meetup_id, user_id, status) VALUES
(36, 3, 2, 'pending'),
(37, 1, 4, 'pending'),
(39, 3, 5, 'accepted'),
(40, 1, 11, 'accepted');

-- MESSAGES
INSERT INTO public.messages (id, meetup_id, sender_id, content, timestamp) VALUES
(4, 3, 11, 'Hey, there', '2025-05-25 13:09:33.333724+03'),
(14, 3, 11, 'Hello, Mr. Smith', '2025-05-26 00:11:06.890059+03'),
(15, 3, 11, 'new content', '2025-05-26 12:50:24.749375+03'),
(16, 3, 11, 'rwerwer', '2025-05-26 13:02:49.720457+03'),
(18, 1, 5, 'Hey, there.', '2025-06-02 19:30:39.915387+03'),
(19, 1, 5, 'What''s up?', '2025-06-02 19:52:51.876387+03'),
(21, 1, 11, 'Hey, I am ok, how are you!!', '2025-06-02 20:00:30.075287+03'),
(29, 1, 5, 'Fine, thanks.', '2025-06-02 20:15:04.542479+03'),
(30, 3, 5, 'I am Mrs, hi anyway', '2025-06-02 20:15:43.849450+03');
