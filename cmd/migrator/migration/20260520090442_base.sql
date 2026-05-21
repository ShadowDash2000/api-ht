-- +goose Up
CREATE TABLE department
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(200) NOT NULL
        CONSTRAINT department_name_min_length CHECK (char_length(name) >= 1),
    parent_id  INTEGER REFERENCES department (id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX department_parent_name_unique_idx ON department (parent_id, name) WHERE parent_id IS NOT NULL;

CREATE TABLE employee
(
    id            SERIAL PRIMARY KEY,
    department_id INTEGER      NOT NULL REFERENCES department (id),
    full_name     VARCHAR(200) NOT NULL
        CONSTRAINT employee_full_name_min_length CHECK (char_length(full_name) >= 1),
    position      VARCHAR(200) NOT NULL
        CONSTRAINT employee_position_min_length CHECK (char_length(position) >= 1),
    hired_at      DATE,
    created_at    TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE employee;
DROP TABLE department;
