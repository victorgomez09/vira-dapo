CREATE TABLE IF NOT EXISTS collections (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    index varchar(100),
    project_id INTEGER NOT NULL,
    creation_date TIMESTAMP NOT NULL, 
    modification_date TIMESTAMP
);

ALTER TABLE collections ADD CONSTRAINT project_fk FOREIGN KEY(project_id) REFERENCES projects(id);