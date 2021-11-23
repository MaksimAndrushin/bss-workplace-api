-- +goose Up
CREATE TABLE workplaces
(
    id      BIGSERIAL,
    name    TEXT      NOT NULL,
    size    INT       NOT NULL,
    removed BOOL      NOT NULL,
    created TIMESTAMP NOT NULL,
    updated TIMESTAMP
) PARTITION BY HASH (id);

create table workplaces_p0 partition of workplaces(primary key (id)) for values with (modulus 4, remainder 0);
create table workplaces_p1 partition of workplaces(primary key (id)) for values with (modulus 4, remainder 1);
create table workplaces_p2 partition of workplaces(primary key (id)) for values with (modulus 4, remainder 2);
create table workplaces_p3 partition of workplaces(primary key (id)) for values with (modulus 4, remainder 3);

CREATE INDEX wrkpl_p0_removed_idx ON workplaces_p0 USING btree (removed);
CREATE INDEX wrkpl_p1_removed_idx ON workplaces_p1 USING btree (removed);
CREATE INDEX wrkpl_p2_removed_idx ON workplaces_p2 USING btree (removed);
CREATE INDEX wrkpl_p3_removed_idx ON workplaces_p3 USING btree (removed);

CREATE TABLE workplaces_events
(
    id           BIGSERIAL,
    workplace_id BIGSERIAL NOT NULL,
    type         INT       NOT NULL,
    status       INT       NOT NULL,
    updated      TIMESTAMP,
    payload      JSONB
) partition by hash(workplace_id);

create table workplaces_event_p0 partition of workplaces_events(workplace_id) for values with (modulus 4, remainder 0);
create unique index workplaces_events_p0_uidx_id on workplaces_event_p0(id);

create table workplaces_event_p1 partition of workplaces_events(workplace_id) for values with (modulus 4, remainder 1);
create unique index workplaces_events_p1_uidx_id on workplaces_event_p1(id);

create table workplaces_event_p2 partition of workplaces_events(workplace_id) for values with (modulus 4, remainder 2);
create unique index workplaces_events_p2_uidx_id on workplaces_event_p2(id);

create table workplaces_event_p3 partition of workplaces_events(workplace_id) for values with (modulus 4, remainder 3);
create unique index workplaces_events_p3_uidx_id on workplaces_event_p3(id);

CREATE INDEX wrkpl_event_p0_status_idx ON workplaces_event_p0 USING btree (status);
CREATE INDEX wrkpl_event_p1_status_idx ON workplaces_event_p1 USING btree (status);
CREATE INDEX wrkpl_event_p2_status_idx ON workplaces_event_p2 USING btree (status);
CREATE INDEX wrkpl_event_p3_status_idx ON workplaces_event_p3 USING btree (status);

-- +goose Down
DROP TABLE workplaces;
DROP TABLE workplaces_events;
