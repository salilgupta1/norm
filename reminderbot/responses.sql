CREATE TABLE responses (
    id SERIAL PRIMARY KEY NOT NULL,
    recipient_id bigint,
    habit_id int REFERENCES habits(id),
    response text,
    sent_at timestamp with time zone DEFAULT timezone('utc'::text, now()),
    responded_at timestamp with time zone
);