--
-- PostgreSQL database dump
--

-- Dumped from database version 17.2
-- Dumped by pg_dump version 17.2

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: meetup_invitations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.meetup_invitations (
    id bigint NOT NULL,
    meetup_id bigint NOT NULL,
    user_id bigint NOT NULL,
    status text DEFAULT 'pending'::text NOT NULL
);


ALTER TABLE public.meetup_invitations OWNER TO postgres;

--
-- Name: meetup_invitations_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.meetup_invitations_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.meetup_invitations_id_seq OWNER TO postgres;

--
-- Name: meetup_invitations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.meetup_invitations_id_seq OWNED BY public.meetup_invitations.id;


--
-- Name: meetup_invitations_meetup_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.meetup_invitations_meetup_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.meetup_invitations_meetup_id_seq OWNER TO postgres;

--
-- Name: meetup_invitations_meetup_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.meetup_invitations_meetup_id_seq OWNED BY public.meetup_invitations.meetup_id;


--
-- Name: meetup_invitations_user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.meetup_invitations_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.meetup_invitations_user_id_seq OWNER TO postgres;

--
-- Name: meetup_invitations_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.meetup_invitations_user_id_seq OWNED BY public.meetup_invitations.user_id;


--
-- Name: meetups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.meetups (
    id bigint NOT NULL,
    name character varying(255) NOT NULL,
    description text,
    user_id bigint NOT NULL
);


ALTER TABLE public.meetups OWNER TO postgres;

--
-- Name: meetups_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.meetups_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.meetups_id_seq OWNER TO postgres;

--
-- Name: meetups_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.meetups_id_seq OWNED BY public.meetups.id;


--
-- Name: meetups_user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.meetups_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.meetups_user_id_seq OWNER TO postgres;

--
-- Name: meetups_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.meetups_user_id_seq OWNED BY public.meetups.user_id;


--
-- Name: messages; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.messages (
    id bigint NOT NULL,
    meetup_id bigint NOT NULL,
    sender_id bigint NOT NULL,
    content text NOT NULL,
    "timestamp" timestamp with time zone DEFAULT now()
);


ALTER TABLE public.messages OWNER TO postgres;

--
-- Name: messages_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.messages_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.messages_id_seq OWNER TO postgres;

--
-- Name: messages_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.messages_id_seq OWNED BY public.messages.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    username character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    first_name character varying(255) NOT NULL,
    last_name character varying(255) NOT NULL,
    password text,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: meetup_invitations id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.meetup_invitations ALTER COLUMN id SET DEFAULT nextval('public.meetup_invitations_id_seq'::regclass);


--
-- Name: meetup_invitations meetup_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.meetup_invitations ALTER COLUMN meetup_id SET DEFAULT nextval('public.meetup_invitations_meetup_id_seq'::regclass);


--
-- Name: meetup_invitations user_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.meetup_invitations ALTER COLUMN user_id SET DEFAULT nextval('public.meetup_invitations_user_id_seq'::regclass);


--
-- Name: meetups id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.meetups ALTER COLUMN id SET DEFAULT nextval('public.meetups_id_seq'::regclass);


--
-- Name: meetups user_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.meetups ALTER COLUMN user_id SET DEFAULT nextval('public.meetups_user_id_seq'::regclass);


--
-- Name: messages id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.messages ALTER COLUMN id SET DEFAULT nextval('public.messages_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: meetup_invitations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.meetup_invitations (id, meetup_id, user_id, status) FROM stdin;
36	3	2	pending
37	1	4	pending
40	1	11	accepted
39	3	5	accepted
\.


--
-- Data for Name: meetups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.meetups (id, name, description, user_id) FROM stdin;
1	Valhala meeting	Meeting to see old friends!!!	5
4	Hunuluu Meeting	A very exotic meeting	11
3	Corparate Meeting	A very important meeting.ffffffffffffffffffytuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuu	11
8	July Meetup	A month meetup for discussing raised issues\n 	5
\.


--
-- Data for Name: messages; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.messages (id, meetup_id, sender_id, content, "timestamp") FROM stdin;
4	3	11	Hey, there	2025-05-25 13:09:33.333724+03
14	3	11	Hello, Mr. Smith	2025-05-26 00:11:06.890059+03
16	3	11	rwerwer	2025-05-26 13:02:49.720457+03
15	3	11	new content	2025-05-26 12:50:24.749375+03
18	1	5	Hey, there.	2025-06-02 19:30:39.915387+03
19	1	5	What's up?	2025-06-02 19:52:51.876387+03
21	1	11	Hey, I am ok, how are you!!	2025-06-02 20:00:30.075287+03
29	1	5	Fine, thanks.	2025-06-02 20:15:04.542479+03
30	3	5	I am Mrs, hi anyway	2025-06-02 20:15:43.84945+03
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, username, email, first_name, last_name, password, created_at, updated_at, deleted_at) FROM stdin;
2	ThorEdge	thor@gmail.com	Thor	Odisson	$2a$10$pfWpqBPAc2vpgtagUPVyhOFyYMZHxYqHev0HRPokp3AtduiLo76tC	2025-02-16 18:37:53.638204+02	2025-02-16 18:37:53.638204+02	\N
4	Odin	odin@gmail.com	Odin	The Maker	$2a$10$17dusYjXHyqEAmhYEZpusu7Ho18Ds8VUus24bfDnc7JTb.YqaL1lK	2025-02-16 21:21:41.643996+02	2025-02-16 21:21:41.643996+02	\N
5	Joleyne	joleyne@gmail.com	Joleyne	Smith	$2a$10$Spaunj2SP4CwOSMjJ67m5OqoOf.6YseMPta22EWxTflWD29ItsZvO	2025-02-21 06:42:47.283544+02	2025-02-21 06:42:47.283544+02	\N
11	Tonkata	tonkata@gmail.com	Toncho	Moncho	$2a$10$V8sXdbNNPulFgz.cnRiiqeUh1C8jXqOt2a8UTHvvvRu6Q8AA.WwCC	2025-05-20 20:13:43.297678+03	2025-05-20 20:13:43.297678+03	\N
\.


--
-- Name: meetup_invitations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.meetup_invitations_id_seq', 40, true);


--
-- Name: meetup_invitations_meetup_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.meetup_invitations_meetup_id_seq', 1, false);


--
-- Name: meetup_invitations_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.meetup_invitations_user_id_seq', 1, false);


--
-- Name: meetups_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.meetups_id_seq', 8, true);


--
-- Name: meetups_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.meetups_user_id_seq', 1, false);


--
-- Name: messages_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.messages_id_seq', 30, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 11, true);


--
-- Name: meetup_invitations meetup_invitations_meetup_id_user_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.meetup_invitations
    ADD CONSTRAINT meetup_invitations_meetup_id_user_id_key UNIQUE (meetup_id, user_id);


--
-- Name: meetup_invitations meetup_invitations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.meetup_invitations
    ADD CONSTRAINT meetup_invitations_pkey PRIMARY KEY (id);


--
-- Name: meetups meetups_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.meetups
    ADD CONSTRAINT meetups_pkey PRIMARY KEY (id);


--
-- Name: meetups meetups_username_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.meetups
    ADD CONSTRAINT meetups_username_key UNIQUE (name);


--
-- Name: messages messages_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.messages
    ADD CONSTRAINT messages_pkey PRIMARY KEY (id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: users users_username_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- Name: meetup_invitations meetup_invitations_meetup_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.meetup_invitations
    ADD CONSTRAINT meetup_invitations_meetup_id_fkey FOREIGN KEY (meetup_id) REFERENCES public.meetups(id) ON DELETE CASCADE;


--
-- Name: meetup_invitations meetup_invitations_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.meetup_invitations
    ADD CONSTRAINT meetup_invitations_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: meetups meetups_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.meetups
    ADD CONSTRAINT meetups_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: messages messages_meetup_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.messages
    ADD CONSTRAINT messages_meetup_id_fkey FOREIGN KEY (meetup_id) REFERENCES public.meetups(id) ON DELETE CASCADE;


--
-- Name: messages messages_sender_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.messages
    ADD CONSTRAINT messages_sender_id_fkey FOREIGN KEY (sender_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

