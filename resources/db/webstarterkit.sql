PGDMP          "                u            webstarterkit    10.1    10.1     ;           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                       false            <           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                       false            =           1262    16386    webstarterkit    DATABASE        CREATE DATABASE webstarterkit WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_US.UTF-8' LC_CTYPE = 'en_US.UTF-8';
    DROP DATABASE webstarterkit;
             nehme    false                        2615    2200    public    SCHEMA        CREATE SCHEMA public;
    DROP SCHEMA public;
             postgres    false            >           0    0    SCHEMA public    COMMENT     6   COMMENT ON SCHEMA public IS 'standard public schema';
                  postgres    false    3                        3079    13253    plpgsql 	   EXTENSION     ?   CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;
    DROP EXTENSION plpgsql;
                  false            ?           0    0    EXTENSION plpgsql    COMMENT     @   COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';
                       false    1            �            1259    16387    sq_admin_id    SEQUENCE     m   CREATE SEQUENCE sq_admin_id
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 "   DROP SEQUENCE public.sq_admin_id;
       public       postgres    false    3            �            1259    16389    admin    TABLE     �  CREATE TABLE admin (
    id integer DEFAULT nextval('sq_admin_id'::regclass) NOT NULL,
    name character varying(150) DEFAULT NULL::character varying,
    email character varying(240) DEFAULT NULL::character varying,
    password character varying(400) DEFAULT NULL::character varying,
    hash character varying(400) DEFAULT NULL::character varying,
    locked_on timestamp without time zone
);
    DROP TABLE public.admin;
       public         postgres    false    196    3            8          0    16389    admin 
   TABLE DATA               D   COPY admin (id, name, email, password, hash, locked_on) FROM stdin;
    public       postgres    false    197   =       @           0    0    sq_admin_id    SEQUENCE SET     2   SELECT pg_catalog.setval('sq_admin_id', 1, true);
            public       postgres    false    196            �           2606    16402    admin pk_admin 
   CONSTRAINT     E   ALTER TABLE ONLY admin
    ADD CONSTRAINT pk_admin PRIMARY KEY (id);
 8   ALTER TABLE ONLY public.admin DROP CONSTRAINT pk_admin;
       public         postgres    false    197            8   n   x�3��K��MU�/�M����� ��������\N�DC���|sG7O�wO=GW�l���$�L��O_�l���@K����\3���r� �ЬpN�?�=... ��"     