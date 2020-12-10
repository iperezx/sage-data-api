CREATE TABLE IF NOT EXISTS SageStorage.Nodes (
    nodeid          VARCHAR(64) NOT NULL,
    metadata_name   VARCHAR(64),
    metadata_value  VARCHAR(64)
);

INSERT INTO SageStorage.Nodes (nodeid, metadata_name, metadata_value) VALUES 
( '4cd98fc4d2a8' ,'NodeName','Sage-NEON-01');

