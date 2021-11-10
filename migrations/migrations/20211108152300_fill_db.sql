-- +goose Up
INSERT INTO workplaces(name, size, removed, created, updated)
SELECT md5(RANDOM()::text),
       RANDOM() * 10 + 1,
       false,
       timestamp '2014-01-10 20:00:00' + random() * (timestamp '2014-01-20 20:00:00' - timestamp '2014-01-10 10:00:00'),
       timestamp '2014-01-10 20:00:00' + random() * (timestamp '2014-01-20 20:00:00' - timestamp '2014-01-10 10:00:00')
FROM generate_series(1, 10000);

INSERT INTO workplaces_events(workplace_id, type, status, updated)
SELECT id,
       1,
       1,
       timestamp '2014-01-10 20:00:00' + random() * (timestamp '2014-01-20 20:00:00' - timestamp '2014-01-10 10:00:00')
FROM workplaces;

INSERT INTO workplaces_events(workplace_id, type, status, updated)
SELECT id,
       random() * 6 + 2,
       1,
       timestamp '2014-01-10 20:00:00' + random() * (timestamp '2014-01-20 20:00:00' - timestamp '2014-01-10 10:00:00')
FROM workplaces;

INSERT INTO workplaces_events(workplace_id, type, status, updated)
SELECT id,
       random() * 6 + 2,
       1,
       timestamp '2014-01-10 20:00:00' + random() * (timestamp '2014-01-20 20:00:00' - timestamp '2014-01-10 10:00:00')
FROM workplaces;

INSERT INTO workplaces_events(workplace_id, type, status, updated)
SELECT id,
       random() * 6 + 2,
       1,
       timestamp '2014-01-10 20:00:00' + random() * (timestamp '2014-01-20 20:00:00' - timestamp '2014-01-10 10:00:00')
FROM workplaces;

delete
from workplaces_events
where type > 3;