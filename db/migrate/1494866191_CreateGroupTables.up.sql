-- Migration goes here.

CREATE TABLE groups
(
  id bigserial NOT NULL,
  name character varying(50),
  image_url character varying(256),
  created_date timestamp without time zone NOT NULL,
  last_modified_date timestamp without time zone,
  CONSTRAINT pk_groups PRIMARY KEY (id),
  CONSTRAINT name_key UNIQUE (name)
);