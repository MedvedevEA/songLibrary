-- +goose Up
-- +goose StatementBegin
--trgm
CREATE SCHEMA trgm;
CREATE EXTENSION pg_trgm WITH SCHEMA trgm;
--groups
CREATE TABLE IF NOT EXISTS public.groups
(
    group_id uuid NOT NULL DEFAULT gen_random_uuid(),
    name character varying COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT groups_pk PRIMARY KEY (group_id)
)
TABLESPACE pg_default;
ALTER TABLE IF EXISTS public.groups
    OWNER to postgres;
CREATE INDEX IF NOT EXISTS groups_name_lower_trgm
    ON public.groups USING gist
    (lower(name::text) COLLATE pg_catalog."default" trgm.gist_trgm_ops)
    TABLESPACE pg_default;
CREATE UNIQUE INDEX IF NOT EXISTS groups_name_lower_unique_btree
    ON public.groups USING btree
    (lower(name::text) COLLATE pg_catalog."default" ASC NULLS LAST)
    TABLESPACE pg_default;
--songs
CREATE TABLE IF NOT EXISTS public.songs
(
    song_id uuid NOT NULL DEFAULT gen_random_uuid(),
    group_id uuid NOT NULL,
    name character varying COLLATE pg_catalog."default" NOT NULL,
    release_date date,
    text character varying COLLATE pg_catalog."default",
    link character varying COLLATE pg_catalog."default",
    CONSTRAINT songs_pk PRIMARY KEY (song_id),
    CONSTRAINT group_fk FOREIGN KEY (group_id)
        REFERENCES public.groups (group_id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE RESTRICT
        NOT VALID
)
TABLESPACE pg_default;
ALTER TABLE IF EXISTS public.songs
    OWNER to postgres;
CREATE INDEX IF NOT EXISTS songs_group_id_btree
    ON public.songs USING btree
    (group_id ASC NULLS LAST)
    TABLESPACE pg_default;
CREATE INDEX IF NOT EXISTS songs_link_lower_trgm
    ON public.songs USING gist
    (lower(link::text) COLLATE pg_catalog."default" trgm.gist_trgm_ops)
    TABLESPACE pg_default;
CREATE INDEX IF NOT EXISTS songs_name_lower_trgm
    ON public.songs USING gist
    (lower(name::text) COLLATE pg_catalog."default" trgm.gist_trgm_ops)
    TABLESPACE pg_default;
CREATE INDEX IF NOT EXISTS songs_release_date_btree
    ON public.songs USING btree
    (release_date ASC NULLS LAST)
    TABLESPACE pg_default;
CREATE INDEX IF NOT EXISTS songs_text_lower_trgm
    ON public.songs USING gist
    (lower(text::text) COLLATE pg_catalog."default" trgm.gist_trgm_ops)
    TABLESPACE pg_default;
--add_group
CREATE OR REPLACE FUNCTION public.add_group(
	request jsonb DEFAULT '{}'::jsonb)
    RETURNS json
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE PARALLEL UNSAFE
AS $BODY$
DECLARE response json;
BEGIN
	WITH t AS(
		INSERT INTO public.groups
		VALUES (
			gen_random_uuid(),
			(add_group.request ->> 'name')::character varying
		)
		RETURNING group_id,"name"
	)
	SELECT row_to_json(t) INTO response FROM t;
	RETURN response;
END;
$BODY$;
ALTER FUNCTION public.add_group(jsonb)
    OWNER TO postgres;
--add_song
CREATE OR REPLACE FUNCTION public.add_song(
	request jsonb DEFAULT '{}'::jsonb)
    RETURNS json
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE PARALLEL UNSAFE
AS $BODY$
DECLARE response json;
BEGIN
	WITH t AS (
		SELECT group_id, "name"
		FROM public.groups
		WHERE LOWER(groups.name)=LOWER((add_song.request ->>'group')::character varying)
	),
	t2 AS (
		INSERT INTO public.songs
		VALUES (
			gen_random_uuid(),
			(SELECT group_id FROM t),
			(add_song.request ->> 'name')::character varying,
			(add_song.request ->> 'release_date')::date,
			(add_song.request ->> 'text')::character varying,
			(add_song.request ->> 'link')::character varying
		)
		RETURNING song_id,"name",release_date,"text","link"
	)
	SELECT row_to_json(t3) INTO response FROM (SELECT t2.*, row_to_json(t) AS group  FROM t,t2) t3;
	RETURN response;
END;
$BODY$;
ALTER FUNCTION public.add_song(jsonb)
    OWNER TO postgres;
--get_group
CREATE OR REPLACE FUNCTION public.get_group(
	request jsonb DEFAULT '{}'::jsonb)
    RETURNS SETOF json 
    LANGUAGE 'plpgsql'
    COST 100
    STABLE PARALLEL UNSAFE
    ROWS 1000
AS $BODY$
BEGIN
	RETURN QUERY
	WITH t AS (
		SELECT group_id,"name" 
		FROM public.groups
		WHERE groups.group_id=(get_group.request ->> 'group_id')::uuid
	)
	SELECT row_to_json(t) 
	FROM t;
END;
$BODY$;
ALTER FUNCTION public.get_group(jsonb)
    OWNER TO postgres;
--get_groups
CREATE OR REPLACE FUNCTION public.get_groups(
	request jsonb DEFAULT '{}'::jsonb)
    RETURNS json
    LANGUAGE 'plpgsql'
    COST 100
    STABLE PARALLEL UNSAFE
AS $BODY$
DECLARE response json;
BEGIN
	WITH t AS (
		SELECT group_id, "name"
		FROM public.groups
		WHERE 
			(get_groups.request ->>'name') IS null OR 
			LOWER(groups.name) LIKE '%' || LOWER((get_groups.request ->>'name')::character varying) || '%'
	)
	SELECT json_build_object( 
		'data',
		COALESCE(
			(SELECT json_agg(row_to_json(t2)) FROM (SELECT * FROM t OFFSET (get_groups.request ->>'offset')::integer LIMIT (get_groups.request ->>'limit')::integer) t2), 
			'[]'::json),
		'all_record_count', 
		(SELECT count(*) FROM t)
	) INTO response;
RETURN response;
END;
$BODY$;
ALTER FUNCTION public.get_groups(jsonb)
    OWNER TO postgres;
--get_song
CREATE OR REPLACE FUNCTION public.get_song(
	request jsonb DEFAULT '{}'::jsonb)
    RETURNS SETOF json 
    LANGUAGE 'plpgsql'
    COST 100
    STABLE PARALLEL UNSAFE
    ROWS 1000
AS $BODY$
BEGIN
	RETURN QUERY
	WITH t AS (
		SELECT 
			_s.song_id,
			_s.name,
			_s.release_date,
			_s.text,
			_s.link,
			row_to_json(_g) AS "group"
		FROM public.songs _s
		JOIN public.groups _g
		ON _g.group_id=_s.group_id
		WHERE _s.song_id=(get_song.request ->>'song_id')::uuid
	)
	SELECT row_to_json(t) 
	FROM t;
END;
$BODY$;
ALTER FUNCTION public.get_song(jsonb)
    OWNER TO postgres;
--get_song_text
CREATE OR REPLACE FUNCTION public.get_song_text(
	request jsonb DEFAULT '{}'::json)
    RETURNS SETOF json 
    LANGUAGE 'plpgsql'
    COST 100
    STABLE PARALLEL UNSAFE
    ROWS 1000
AS $BODY$
BEGIN
	RETURN QUERY
	WITH t AS (
		SELECT *
		FROM public.songs
		WHERE songs.song_id=(get_song_text.request ->> 'song_id')::uuid
	),
	t2 AS (
		SELECT ROW_NUMBER() OVER() AS "index",	"text"
		FROM (SELECT string_to_table(t.text,E'\n\n') AS "text" FROM t)
	)
	SELECT json_build_object(
		'data',
		COALESCE(
			(SELECT json_agg(row_to_json(t3)) FROM (SELECT * FROM t2 OFFSET (get_song_text.request ->>'offset')::integer LIMIT (get_song_text.request ->>'limit')::integer) t3),
			'[]'::json),
		'all_record_count',
		(SELECT COUNT(*) FROM t2)) 
	FROM t;
END;
$BODY$;
ALTER FUNCTION public.get_song_text(jsonb)
    OWNER TO postgres;
--get_songs
CREATE OR REPLACE FUNCTION public.get_songs(
	request jsonb DEFAULT '{}'::json)
    RETURNS json
    LANGUAGE 'plpgsql'
    COST 100
    STABLE PARALLEL UNSAFE
AS $BODY$
DECLARE response json;
BEGIN
	WITH t AS (
		SELECT 
			_s.song_id,
			_s.name,
			_s.release_date,
			_s.text,
			_s.link,
			row_to_json(_g) AS "group"
		FROM public.songs _s
		JOIN public.groups _g
		ON _g.group_id=_s.group_id
		WHERE 
			((get_songs.request ->>'group') IS null OR LOWER(_g.name) LIKE '%' || LOWER((get_songs.request ->>'group')::character varying) || '%') AND
			((get_songs.request ->>'name') IS null OR LOWER(_s.name) LIKE '%' || LOWER((get_songs.request ->>'name')::character varying) || '%') AND
			((get_songs.request ->>'release_date') IS null OR _s.release_date = (get_songs.request ->>'release_date')::date) AND
			((get_songs.request ->>'text') IS null OR LOWER(_s.text) LIKE '%' || LOWER((get_songs.request ->>'text')::character varying) || '%') AND
			((get_songs.request ->>'link') IS null OR LOWER(_s.link) LIKE '%' || LOWER((get_songs.request ->>'link')::character varying) || '%') 
	)
	SELECT json_build_object( 
		'data',
		COALESCE(
			(SELECT json_agg(row_to_json(t2)) FROM (SELECT * FROM t OFFSET (get_songs.request ->>'offset')::integer LIMIT (get_songs.request ->>'limit')::integer) t2), 
			'[]'::json),
		'all_record_count', 
		(SELECT count(*) from t)
	) INTO response;
RETURN response;
END;
$BODY$;
ALTER FUNCTION public.get_songs(jsonb)
    OWNER TO postgres;
--remove_group
CREATE OR REPLACE FUNCTION public.remove_group(
	request jsonb DEFAULT '{}'::jsonb)
    RETURNS SETOF json 
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE PARALLEL UNSAFE
    ROWS 1000
AS $BODY$
BEGIN
	RETURN QUERY
	WITH t AS (
		DELETE FROM public.groups
		WHERE groups.group_id=(remove_group.request ->>'group_id')::uuid
		RETURNING group_id
	)
	SELECT row_to_json(t) 
	FROM t;
END;
$BODY$;
ALTER FUNCTION public.remove_group(jsonb)
    OWNER TO postgres;
--remove_song
CREATE OR REPLACE FUNCTION public.remove_song(
	request jsonb DEFAULT '{}'::jsonb)
    RETURNS SETOF json 
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE PARALLEL UNSAFE
    ROWS 1000
AS $BODY$
BEGIN
	RETURN QUERY
	WITH t AS (
		DELETE FROM public.songs
		WHERE songs.song_id=(remove_song.request ->>'song_id')::uuid
		RETURNING song_id
	)
	SELECT row_to_json(t)
	FROM t;
END;
$BODY$;
ALTER FUNCTION public.remove_song(jsonb)
    OWNER TO postgres;
--update_group
CREATE OR REPLACE FUNCTION public.update_group(
	request jsonb DEFAULT '{}'::jsonb)
    RETURNS SETOF json 
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE PARALLEL UNSAFE
    ROWS 1000
AS $BODY$
BEGIN
	RETURN QUERY
	WITH t AS (
		UPDATE public.groups
		SET	"name"=COALESCE((update_group.request ->>'name')::character varying,"name")
		WHERE groups.group_id=(update_group.request ->>'group_id')::uuid
		RETURNING group_id
	)
	SELECT row_to_json(t) 
	FROM t;
END;
$BODY$;
ALTER FUNCTION public.update_group(jsonb)
    OWNER TO postgres;
--update_song
CREATE OR REPLACE FUNCTION public.update_song(
	request jsonb DEFAULT '{}'::jsonb)
    RETURNS SETOF json 
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE PARALLEL UNSAFE
    ROWS 1000
AS $BODY$
BEGIN
	RETURN QUERY
	WITH t AS (
		UPDATE public.songs
		SET	
			"group_id"=COALESCE((update_song.request ->>'group_id')::uuid,group_id),
			"name"=COALESCE((update_song.request ->>'name')::character varying,"name"),
			release_date= CASE WHEN (update_song.request ->>'set_release_date')::boolean=true OR  (update_song.request ->>'release_date') IS NOT null THEN (update_song.request ->>'release_date')::date ELSE release_date END,
			"text"= CASE WHEN (update_song.request ->>'set_text')::boolean=true OR (update_song.request ->>'text') IS NOT null THEN (update_song.request ->>'text')::character varying ELSE songs.text END,
			"link"= CASE WHEN (update_song.request ->>'set_link')::boolean=true OR (update_song.request ->>'link') IS NOT null THEN (update_song.request ->>'link')::character varying ELSE songs.link END
		WHERE songs.song_id=(update_song.request ->>'song_id')::uuid
		RETURNING song_id
	)
	SELECT row_to_json(t)
	FROM t;
END;
$BODY$;
ALTER FUNCTION public.update_song(jsonb)
    OWNER TO postgres;
-- +goose StatementEnd
