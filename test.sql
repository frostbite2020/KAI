--
-- PostgreSQL database dump
--

-- Dumped from database version 17.2 (Debian 17.2-1.pgdg120+1)
-- Dumped by pg_dump version 17.2

-- Started on 2024-11-28 06:13:58

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

--
-- TOC entry 4 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: pg_database_owner
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO pg_database_owner;

--
-- TOC entry 3504 (class 0 OID 0)
-- Dependencies: 4
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: pg_database_owner
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 227 (class 1259 OID 16500)
-- Name: booking_seats; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.booking_seats (
    booking_id bigint NOT NULL,
    seat_id bigint NOT NULL
);


ALTER TABLE public.booking_seats OWNER TO postgres;

--
-- TOC entry 233 (class 1259 OID 16541)
-- Name: bookings; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.bookings (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_id bigint,
    schedule_id bigint,
    booking_date timestamp with time zone,
    total_amount numeric
);


ALTER TABLE public.bookings OWNER TO postgres;

--
-- TOC entry 232 (class 1259 OID 16540)
-- Name: bookings_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.bookings_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.bookings_id_seq OWNER TO postgres;

--
-- TOC entry 3505 (class 0 OID 0)
-- Dependencies: 232
-- Name: bookings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.bookings_id_seq OWNED BY public.bookings.id;


--
-- TOC entry 235 (class 1259 OID 16561)
-- Name: carriages; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.carriages (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    train_id bigint,
    number text,
    type text,
    capacity bigint
);


ALTER TABLE public.carriages OWNER TO postgres;

--
-- TOC entry 234 (class 1259 OID 16560)
-- Name: carriages_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.carriages_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.carriages_id_seq OWNER TO postgres;

--
-- TOC entry 3506 (class 0 OID 0)
-- Dependencies: 234
-- Name: carriages_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.carriages_id_seq OWNED BY public.carriages.id;


--
-- TOC entry 218 (class 1259 OID 16430)
-- Name: cities; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cities (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text
);


ALTER TABLE public.cities OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 16429)
-- Name: cities_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.cities_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.cities_id_seq OWNER TO postgres;

--
-- TOC entry 3507 (class 0 OID 0)
-- Dependencies: 217
-- Name: cities_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.cities_id_seq OWNED BY public.cities.id;


--
-- TOC entry 222 (class 1259 OID 16455)
-- Name: routes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.routes (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    start_station_id bigint,
    end_station_id bigint,
    distance bigint,
    travel_time bigint
);


ALTER TABLE public.routes OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 16454)
-- Name: routes_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.routes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.routes_id_seq OWNER TO postgres;

--
-- TOC entry 3508 (class 0 OID 0)
-- Dependencies: 221
-- Name: routes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.routes_id_seq OWNED BY public.routes.id;


--
-- TOC entry 239 (class 1259 OID 24577)
-- Name: schedule_carriage_prices; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schedule_carriage_prices (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    schedule_id bigint,
    carriage_id bigint,
    price numeric
);


ALTER TABLE public.schedule_carriage_prices OWNER TO postgres;

--
-- TOC entry 238 (class 1259 OID 24576)
-- Name: schedule_carriage_prices_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.schedule_carriage_prices_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.schedule_carriage_prices_id_seq OWNER TO postgres;

--
-- TOC entry 3509 (class 0 OID 0)
-- Dependencies: 238
-- Name: schedule_carriage_prices_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.schedule_carriage_prices_id_seq OWNED BY public.schedule_carriage_prices.id;


--
-- TOC entry 231 (class 1259 OID 16521)
-- Name: schedules; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schedules (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    train_id bigint,
    departure text,
    arrival text,
    route_id bigint
);


ALTER TABLE public.schedules OWNER TO postgres;

--
-- TOC entry 230 (class 1259 OID 16520)
-- Name: schedules_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.schedules_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.schedules_id_seq OWNER TO postgres;

--
-- TOC entry 3510 (class 0 OID 0)
-- Dependencies: 230
-- Name: schedules_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.schedules_id_seq OWNED BY public.schedules.id;


--
-- TOC entry 237 (class 1259 OID 16576)
-- Name: seats; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.seats (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    carriage_id bigint,
    seat_number text,
    booked boolean
);


ALTER TABLE public.seats OWNER TO postgres;

--
-- TOC entry 236 (class 1259 OID 16575)
-- Name: seats_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.seats_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.seats_id_seq OWNER TO postgres;

--
-- TOC entry 3511 (class 0 OID 0)
-- Dependencies: 236
-- Name: seats_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.seats_id_seq OWNED BY public.seats.id;


--
-- TOC entry 226 (class 1259 OID 16485)
-- Name: sessions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sessions (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    token text NOT NULL,
    expires_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.sessions OWNER TO postgres;

--
-- TOC entry 225 (class 1259 OID 16484)
-- Name: sessions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sessions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.sessions_id_seq OWNER TO postgres;

--
-- TOC entry 3512 (class 0 OID 0)
-- Dependencies: 225
-- Name: sessions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sessions_id_seq OWNED BY public.sessions.id;


--
-- TOC entry 220 (class 1259 OID 16440)
-- Name: stations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.stations (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    city_id bigint
);


ALTER TABLE public.stations OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 16439)
-- Name: stations_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.stations_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.stations_id_seq OWNER TO postgres;

--
-- TOC entry 3513 (class 0 OID 0)
-- Dependencies: 219
-- Name: stations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.stations_id_seq OWNED BY public.stations.id;


--
-- TOC entry 229 (class 1259 OID 16506)
-- Name: trains; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.trains (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    type text,
    train_number text
);


ALTER TABLE public.trains OWNER TO postgres;

--
-- TOC entry 228 (class 1259 OID 16505)
-- Name: trains_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.trains_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.trains_id_seq OWNER TO postgres;

--
-- TOC entry 3514 (class 0 OID 0)
-- Dependencies: 228
-- Name: trains_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.trains_id_seq OWNED BY public.trains.id;


--
-- TOC entry 224 (class 1259 OID 16474)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    email text NOT NULL,
    name text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 223 (class 1259 OID 16473)
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
-- TOC entry 3515 (class 0 OID 0)
-- Dependencies: 223
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 3271 (class 2604 OID 16544)
-- Name: bookings id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings ALTER COLUMN id SET DEFAULT nextval('public.bookings_id_seq'::regclass);


--
-- TOC entry 3272 (class 2604 OID 16564)
-- Name: carriages id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.carriages ALTER COLUMN id SET DEFAULT nextval('public.carriages_id_seq'::regclass);


--
-- TOC entry 3264 (class 2604 OID 16433)
-- Name: cities id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cities ALTER COLUMN id SET DEFAULT nextval('public.cities_id_seq'::regclass);


--
-- TOC entry 3266 (class 2604 OID 16458)
-- Name: routes id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.routes ALTER COLUMN id SET DEFAULT nextval('public.routes_id_seq'::regclass);


--
-- TOC entry 3274 (class 2604 OID 24580)
-- Name: schedule_carriage_prices id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schedule_carriage_prices ALTER COLUMN id SET DEFAULT nextval('public.schedule_carriage_prices_id_seq'::regclass);


--
-- TOC entry 3270 (class 2604 OID 16524)
-- Name: schedules id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schedules ALTER COLUMN id SET DEFAULT nextval('public.schedules_id_seq'::regclass);


--
-- TOC entry 3273 (class 2604 OID 16579)
-- Name: seats id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.seats ALTER COLUMN id SET DEFAULT nextval('public.seats_id_seq'::regclass);


--
-- TOC entry 3268 (class 2604 OID 16488)
-- Name: sessions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions ALTER COLUMN id SET DEFAULT nextval('public.sessions_id_seq'::regclass);


--
-- TOC entry 3265 (class 2604 OID 16443)
-- Name: stations id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.stations ALTER COLUMN id SET DEFAULT nextval('public.stations_id_seq'::regclass);


--
-- TOC entry 3269 (class 2604 OID 16509)
-- Name: trains id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trains ALTER COLUMN id SET DEFAULT nextval('public.trains_id_seq'::regclass);


--
-- TOC entry 3267 (class 2604 OID 16477)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 3486 (class 0 OID 16500)
-- Dependencies: 227
-- Data for Name: booking_seats; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.booking_seats (booking_id, seat_id) FROM stdin;
\.


--
-- TOC entry 3492 (class 0 OID 16541)
-- Dependencies: 233
-- Data for Name: bookings; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.bookings (id, created_at, updated_at, deleted_at, user_id, schedule_id, booking_date, total_amount) FROM stdin;
\.


--
-- TOC entry 3494 (class 0 OID 16561)
-- Dependencies: 235
-- Data for Name: carriages; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.carriages (id, created_at, updated_at, deleted_at, train_id, number, type, capacity) FROM stdin;
1	2024-11-27 22:12:15.213947+00	2024-11-27 22:12:15.213947+00	\N	3	1A	Executive	50
2	2024-11-27 22:13:26.996409+00	2024-11-27 22:13:26.996409+00	\N	3	2B	Business	100
4	2024-11-27 22:14:09.309481+00	2024-11-27 22:14:09.309481+00	\N	3	3C	Economy	200
\.


--
-- TOC entry 3477 (class 0 OID 16430)
-- Dependencies: 218
-- Data for Name: cities; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.cities (id, created_at, updated_at, deleted_at, name) FROM stdin;
1	2024-11-27 22:03:34.430299+00	2024-11-27 22:03:34.430299+00	\N	Jakarta
2	2024-11-27 22:06:01.74748+00	2024-11-27 22:06:01.74748+00	\N	Malang
\.


--
-- TOC entry 3481 (class 0 OID 16455)
-- Dependencies: 222
-- Data for Name: routes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.routes (id, created_at, updated_at, deleted_at, start_station_id, end_station_id, distance, travel_time) FROM stdin;
1	2024-11-27 22:10:03.968115+00	2024-11-27 22:10:03.968115+00	\N	1	2	855000	55800
2	2024-11-27 22:10:50.50875+00	2024-11-27 22:10:50.50875+00	\N	1	3	865000	57600
\.


--
-- TOC entry 3498 (class 0 OID 24577)
-- Dependencies: 239
-- Data for Name: schedule_carriage_prices; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.schedule_carriage_prices (id, created_at, updated_at, deleted_at, schedule_id, carriage_id, price) FROM stdin;
1	2024-11-27 22:22:57.656817+00	2024-11-27 22:22:57.656817+00	\N	1	2	750000
2	2024-11-27 22:23:35.456084+00	2024-11-27 22:23:35.456084+00	\N	1	1	1350000
3	2024-11-27 22:23:54.993135+00	2024-11-27 22:23:54.993135+00	\N	1	4	300000
\.


--
-- TOC entry 3490 (class 0 OID 16521)
-- Dependencies: 231
-- Data for Name: schedules; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.schedules (id, created_at, updated_at, deleted_at, train_id, departure, arrival, route_id) FROM stdin;
1	2024-11-27 22:19:14.434417+00	2024-11-27 22:19:14.434417+00	\N	3	2024-12-01T08:00:00Z	2024-12-01T23:30:00Z	1
\.


--
-- TOC entry 3496 (class 0 OID 16576)
-- Dependencies: 237
-- Data for Name: seats; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.seats (id, created_at, updated_at, deleted_at, carriage_id, seat_number, booked) FROM stdin;
\.


--
-- TOC entry 3485 (class 0 OID 16485)
-- Dependencies: 226
-- Data for Name: sessions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sessions (id, user_id, token, expires_at, created_at, updated_at) FROM stdin;
1	1	DJhbAsfXCAdcDi0LyVqjT6ZcylJrIWX5DCyRydIY9WE=	2024-11-28 14:08:24.632772+00	2024-11-27 14:08:24.632772+00	2024-11-27 14:08:24.632772+00
2	1	nDnRfhpJfctkOEXHHWJzFhL00IEDwc8RaRI_chdDsUk=	2024-11-28 14:28:35.979498+00	2024-11-27 14:28:35.981229+00	2024-11-27 14:28:35.981229+00
3	1	zTJrKqHMHStHoBm2ksJaJmlNoU5VIbrMGUSVRXGSUyg=	2024-11-28 14:35:52.31664+00	2024-11-27 14:35:52.317157+00	2024-11-27 14:35:52.317157+00
\.


--
-- TOC entry 3479 (class 0 OID 16440)
-- Dependencies: 220
-- Data for Name: stations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.stations (id, created_at, updated_at, deleted_at, name, city_id) FROM stdin;
1	2024-11-27 22:05:17.972293+00	2024-11-27 22:05:17.972293+00	\N	PASAR SENEN (PSE)	1
2	2024-11-27 22:06:19.731874+00	2024-11-27 22:06:19.731874+00	\N	MALANG (ML)	2
3	2024-11-27 22:06:32.483436+00	2024-11-27 22:06:32.483436+00	\N	KEPANJEN (KPN)	2
\.


--
-- TOC entry 3488 (class 0 OID 16506)
-- Dependencies: 229
-- Data for Name: trains; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.trains (id, created_at, updated_at, deleted_at, name, type, train_number) FROM stdin;
3	2024-11-27 16:16:57.554245+00	2024-11-27 16:16:57.554245+00	\N	Matarmaja	Intercity Train	1001
\.


--
-- TOC entry 3483 (class 0 OID 16474)
-- Dependencies: 224
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, email, name, created_at, updated_at) FROM stdin;
1	ardiansyahrafi4@gmail.com		2024-11-27 14:08:24.626796+00	2024-11-27 14:08:24.626796+00
\.


--
-- TOC entry 3516 (class 0 OID 0)
-- Dependencies: 232
-- Name: bookings_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.bookings_id_seq', 1, false);


--
-- TOC entry 3517 (class 0 OID 0)
-- Dependencies: 234
-- Name: carriages_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.carriages_id_seq', 4, true);


--
-- TOC entry 3518 (class 0 OID 0)
-- Dependencies: 217
-- Name: cities_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.cities_id_seq', 2, true);


--
-- TOC entry 3519 (class 0 OID 0)
-- Dependencies: 221
-- Name: routes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.routes_id_seq', 2, true);


--
-- TOC entry 3520 (class 0 OID 0)
-- Dependencies: 238
-- Name: schedule_carriage_prices_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.schedule_carriage_prices_id_seq', 3, true);


--
-- TOC entry 3521 (class 0 OID 0)
-- Dependencies: 230
-- Name: schedules_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.schedules_id_seq', 1, true);


--
-- TOC entry 3522 (class 0 OID 0)
-- Dependencies: 236
-- Name: seats_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.seats_id_seq', 1, false);


--
-- TOC entry 3523 (class 0 OID 0)
-- Dependencies: 225
-- Name: sessions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sessions_id_seq', 3, true);


--
-- TOC entry 3524 (class 0 OID 0)
-- Dependencies: 219
-- Name: stations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.stations_id_seq', 3, true);


--
-- TOC entry 3525 (class 0 OID 0)
-- Dependencies: 228
-- Name: trains_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.trains_id_seq', 3, true);


--
-- TOC entry 3526 (class 0 OID 0)
-- Dependencies: 223
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 1, true);


--
-- TOC entry 3293 (class 2606 OID 16504)
-- Name: booking_seats booking_seats_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.booking_seats
    ADD CONSTRAINT booking_seats_pkey PRIMARY KEY (booking_id, seat_id);


--
-- TOC entry 3301 (class 2606 OID 16548)
-- Name: bookings bookings_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings
    ADD CONSTRAINT bookings_pkey PRIMARY KEY (id);


--
-- TOC entry 3304 (class 2606 OID 16568)
-- Name: carriages carriages_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.carriages
    ADD CONSTRAINT carriages_pkey PRIMARY KEY (id);


--
-- TOC entry 3276 (class 2606 OID 16437)
-- Name: cities cities_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cities
    ADD CONSTRAINT cities_pkey PRIMARY KEY (id);


--
-- TOC entry 3283 (class 2606 OID 16460)
-- Name: routes routes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.routes
    ADD CONSTRAINT routes_pkey PRIMARY KEY (id);


--
-- TOC entry 3311 (class 2606 OID 24584)
-- Name: schedule_carriage_prices schedule_carriage_prices_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schedule_carriage_prices
    ADD CONSTRAINT schedule_carriage_prices_pkey PRIMARY KEY (id);


--
-- TOC entry 3299 (class 2606 OID 16528)
-- Name: schedules schedules_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schedules
    ADD CONSTRAINT schedules_pkey PRIMARY KEY (id);


--
-- TOC entry 3308 (class 2606 OID 16583)
-- Name: seats seats_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.seats
    ADD CONSTRAINT seats_pkey PRIMARY KEY (id);


--
-- TOC entry 3289 (class 2606 OID 16492)
-- Name: sessions sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (id);


--
-- TOC entry 3280 (class 2606 OID 16447)
-- Name: stations stations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.stations
    ADD CONSTRAINT stations_pkey PRIMARY KEY (id);


--
-- TOC entry 3296 (class 2606 OID 16513)
-- Name: trains trains_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trains
    ADD CONSTRAINT trains_pkey PRIMARY KEY (id);


--
-- TOC entry 3291 (class 2606 OID 16494)
-- Name: sessions uni_sessions_token; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT uni_sessions_token UNIQUE (token);


--
-- TOC entry 3285 (class 2606 OID 16483)
-- Name: users uni_users_email; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT uni_users_email UNIQUE (email);


--
-- TOC entry 3287 (class 2606 OID 16481)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 3302 (class 1259 OID 16559)
-- Name: idx_bookings_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bookings_deleted_at ON public.bookings USING btree (deleted_at);


--
-- TOC entry 3305 (class 1259 OID 16574)
-- Name: idx_carriages_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_carriages_deleted_at ON public.carriages USING btree (deleted_at);


--
-- TOC entry 3277 (class 1259 OID 16438)
-- Name: idx_cities_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_cities_deleted_at ON public.cities USING btree (deleted_at);


--
-- TOC entry 3281 (class 1259 OID 16471)
-- Name: idx_routes_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_routes_deleted_at ON public.routes USING btree (deleted_at);


--
-- TOC entry 3309 (class 1259 OID 24595)
-- Name: idx_schedule_carriage_prices_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_schedule_carriage_prices_deleted_at ON public.schedule_carriage_prices USING btree (deleted_at);


--
-- TOC entry 3297 (class 1259 OID 16539)
-- Name: idx_schedules_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_schedules_deleted_at ON public.schedules USING btree (deleted_at);


--
-- TOC entry 3306 (class 1259 OID 16589)
-- Name: idx_seats_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_seats_deleted_at ON public.seats USING btree (deleted_at);


--
-- TOC entry 3278 (class 1259 OID 16453)
-- Name: idx_stations_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_stations_deleted_at ON public.stations USING btree (deleted_at);


--
-- TOC entry 3294 (class 1259 OID 16514)
-- Name: idx_trains_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_trains_deleted_at ON public.trains USING btree (deleted_at);


--
-- TOC entry 3319 (class 2606 OID 16590)
-- Name: booking_seats fk_booking_seats_booking; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.booking_seats
    ADD CONSTRAINT fk_booking_seats_booking FOREIGN KEY (booking_id) REFERENCES public.bookings(id);


--
-- TOC entry 3320 (class 2606 OID 16595)
-- Name: booking_seats fk_booking_seats_seat; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.booking_seats
    ADD CONSTRAINT fk_booking_seats_seat FOREIGN KEY (seat_id) REFERENCES public.seats(id);


--
-- TOC entry 3324 (class 2606 OID 16554)
-- Name: bookings fk_bookings_schedule; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings
    ADD CONSTRAINT fk_bookings_schedule FOREIGN KEY (schedule_id) REFERENCES public.schedules(id);


--
-- TOC entry 3325 (class 2606 OID 16549)
-- Name: bookings fk_bookings_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings
    ADD CONSTRAINT fk_bookings_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 3328 (class 2606 OID 16584)
-- Name: seats fk_carriages_seats; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.seats
    ADD CONSTRAINT fk_carriages_seats FOREIGN KEY (carriage_id) REFERENCES public.carriages(id);


--
-- TOC entry 3326 (class 2606 OID 24606)
-- Name: carriages fk_carriages_train; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.carriages
    ADD CONSTRAINT fk_carriages_train FOREIGN KEY (train_id) REFERENCES public.trains(id);


--
-- TOC entry 3312 (class 2606 OID 16448)
-- Name: stations fk_cities_stations; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.stations
    ADD CONSTRAINT fk_cities_stations FOREIGN KEY (city_id) REFERENCES public.cities(id);


--
-- TOC entry 3314 (class 2606 OID 16461)
-- Name: routes fk_routes_end_station; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.routes
    ADD CONSTRAINT fk_routes_end_station FOREIGN KEY (end_station_id) REFERENCES public.stations(id);


--
-- TOC entry 3315 (class 2606 OID 24596)
-- Name: routes fk_routes_start_station; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.routes
    ADD CONSTRAINT fk_routes_start_station FOREIGN KEY (start_station_id) REFERENCES public.stations(id);


--
-- TOC entry 3329 (class 2606 OID 24590)
-- Name: schedule_carriage_prices fk_schedule_carriage_prices_carriage; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schedule_carriage_prices
    ADD CONSTRAINT fk_schedule_carriage_prices_carriage FOREIGN KEY (carriage_id) REFERENCES public.carriages(id);


--
-- TOC entry 3330 (class 2606 OID 24585)
-- Name: schedule_carriage_prices fk_schedule_carriage_prices_schedule; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schedule_carriage_prices
    ADD CONSTRAINT fk_schedule_carriage_prices_schedule FOREIGN KEY (schedule_id) REFERENCES public.schedules(id);


--
-- TOC entry 3321 (class 2606 OID 16534)
-- Name: schedules fk_schedules_route; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schedules
    ADD CONSTRAINT fk_schedules_route FOREIGN KEY (route_id) REFERENCES public.routes(id);


--
-- TOC entry 3322 (class 2606 OID 16529)
-- Name: schedules fk_schedules_train; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schedules
    ADD CONSTRAINT fk_schedules_train FOREIGN KEY (train_id) REFERENCES public.trains(id);


--
-- TOC entry 3317 (class 2606 OID 16495)
-- Name: sessions fk_sessions_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT fk_sessions_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 3313 (class 2606 OID 24601)
-- Name: stations fk_stations_city; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.stations
    ADD CONSTRAINT fk_stations_city FOREIGN KEY (city_id) REFERENCES public.cities(id);


--
-- TOC entry 3316 (class 2606 OID 16466)
-- Name: routes fk_stations_routes; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.routes
    ADD CONSTRAINT fk_stations_routes FOREIGN KEY (start_station_id) REFERENCES public.stations(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- TOC entry 3327 (class 2606 OID 16569)
-- Name: carriages fk_trains_carriages; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.carriages
    ADD CONSTRAINT fk_trains_carriages FOREIGN KEY (train_id) REFERENCES public.trains(id);


--
-- TOC entry 3323 (class 2606 OID 16600)
-- Name: schedules fk_trains_schedules; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schedules
    ADD CONSTRAINT fk_trains_schedules FOREIGN KEY (train_id) REFERENCES public.trains(id);


--
-- TOC entry 3318 (class 2606 OID 16515)
-- Name: sessions fk_users_sessions; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT fk_users_sessions FOREIGN KEY (user_id) REFERENCES public.users(id);


-- Completed on 2024-11-28 06:13:59

--
-- PostgreSQL database dump complete
--

