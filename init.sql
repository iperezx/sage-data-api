CREATE TABLE IF NOT EXISTS SageStorage.Nodes (
    nodeid          BINARY(16) NOT NULL PRIMARY KEY,
    metadata_name   VARCHAR(64),
    metadata_value  VARCHAR(64),
    geom            point srid 4326
);

