CREATE TABLE IF NOT EXISTS SageStorage.Nodes (
    nodeid          VARCHAR(64) NOT NULL PRIMARY KEY,
    metadata_name   VARCHAR(64),
    metadata_value  VARCHAR(64),
    geom            point srid 4326
);

INSERT INTO SageStorage.Nodes (nodeid, metadata_name, metadata_value, geom) VALUES 
( '4cd98fc4d2a8' ,'NodeName','Sage-NEON-01',ST_GeomFromText(' POINT(40.01631 -105.24585) ',4326 ));

