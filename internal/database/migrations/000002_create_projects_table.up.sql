CREATE TABLE IF NOT EXISTS projects (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    creation_date TIMESTAMP NOT NULL, 
    modification_date TIMESTAMP
);

CREATE TABLE IF NOT EXISTS user_projects (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    project_id INTEGER NOT NULL
);

ALTER TABLE user_projects ADD CONSTRAINT user_id_fk FOREIGN KEY(user_id) REFERENCES users(id);
ALTER TABLE user_projects ADD CONSTRAINT project_id_fk FOREIGN KEY(project_id) REFERENCES projects(id);

