CREATE TABLE channels (
    id bigserial not null,
    channel_id varchar,
    channel_name varchar,
    channel_info text
);

CREATE TABLE videos (
    id bigserial not null,
    video_id varchar,
    video_title varchar,
    publish_date varchar,
    video_info text
);

CREATE TABLE playlist (
    id bigserial not null,
    playlist_id varchar,
    playlist_title text,
    embeded_html text,
    video_count bigint
);