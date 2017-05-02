CREATE TABLE users
(
  id bigserial NOT NULL,
  login character varying(50) NOT NULL,
  password_hash character varying(60),
  first_name character varying(50),
  last_name character varying(50),
  email character varying(100),
  image_url character varying(256),
  activated boolean NOT NULL,
  lang_key character varying(5),
  activation_key character varying(20),
  reset_key character varying(20),
  created_by character varying(50) NOT NULL,
  created_date timestamp without time zone NOT NULL,
  reset_date timestamp without time zone,
  last_modified_by character varying(50),
  last_modified_date timestamp without time zone,
  CONSTRAINT pk_user PRIMARY KEY (id),
  CONSTRAINT user_email_key UNIQUE (email),
  CONSTRAINT user_login_key UNIQUE (login)
);

INSERT INTO users (
    login, password_hash, first_name, last_name, email, image_url, 
            activated, lang_key, activation_key, reset_key, created_by, created_date, 
            reset_date, last_modified_by, last_modified_date) 
    VALUES (
    'admin', 
    crypt('admin123456', gen_salt('bf', 10)), --Change this PASSWORD for you own admin password
    'Administrator', 
    'Administrator', 
    'admin@localhost',
     NULL, 
     true, 
     'en_US',
     NULL, NULL, 
     'admin', 
     '2017-04-13 21:33:16.367254', 
     NULL, 'admin', NULL);

CREATE TABLE authorities
(
  name character varying(50) NOT NULL,
  CONSTRAINT pk_authority PRIMARY KEY (name)
);

INSERT INTO authorities VALUES ('ROLE_ADMIN');
INSERT INTO authorities VALUES ('ROLE_USER');

CREATE TABLE users_authorities
(
  user_id bigint NOT NULL,
  authority_name character varying(50) NOT NULL,
  CONSTRAINT user_authority_pkey PRIMARY KEY (user_id, authority_name),
  CONSTRAINT fk_authority_name FOREIGN KEY (authority_name)
      REFERENCES authorities (name) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT fk_user_id FOREIGN KEY (user_id)
      REFERENCES users (id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION
);

INSERT INTO users_authorities VALUES (
    (select id from users where login = 'admin'),
    'ROLE_ADMIN');