CREATE TABLE schedules (
    id SERIAL PRIMARY KEY  NOT NULL,
    habit_id int REFERENCES habits(id),
    time int
);