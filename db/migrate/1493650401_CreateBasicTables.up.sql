CREATE TABLE users
(
  id bigserial NOT NULL,
  password_hash character varying(60),
  first_name character varying(50),
  last_name character varying(50),
  email character varying(100),
  image_url character varying(256),
  activated boolean NOT NULL,
  lang_key character varying(5),
  activation_key character varying(20),
  reset_key character varying(20),
  created_date timestamp without time zone NOT NULL,
  reset_date timestamp without time zone,
  last_modified_date timestamp without time zone,
  CONSTRAINT pk_user PRIMARY KEY (id),
  CONSTRAINT user_email_key UNIQUE (email)
);

INSERT INTO users (
    password_hash, first_name, last_name, email, image_url, 
            activated, lang_key, activation_key, reset_key, created_date, 
            reset_date, last_modified_date) 
    VALUES (
    crypt('admin123456', gen_salt('bf', 10)), --Change to your password
    'Administrator', 
    'Administrator', 
    'elirenato2000@gmail.com', --Change to your email
     NULL, 
     true, 
     'en-US',
     NULL, NULL, 
     '2017-04-13 21:33:16.367254', 
     NULL, NULL);

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
    (select id from users where email = 'elirenato2000@gmail.com'),
    'ROLE_ADMIN');