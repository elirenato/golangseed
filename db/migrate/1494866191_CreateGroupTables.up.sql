-- Migration goes here.

CREATE TABLE groups
(
  id bigserial NOT NULL,
  name character varying(50),
  image_url character varying(256),
  created_date timestamp without time zone NOT NULL,
  last_modified_date timestamp without time zone,
  owner_id bigint NOT NULL,
  CONSTRAINT pk_groups PRIMARY KEY (id),
  CONSTRAINT name_key UNIQUE (name),
  CONSTRAINT fk_user_id FOREIGN KEY (owner_id)
      REFERENCES users (id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION
);