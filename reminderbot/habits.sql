CREATE TABLE habits (
    id SERIAL PRIMARY KEY NOT NULL,
    recipient_id bigint,
    content text,
    frequency text,
    created_at timestamp with time zone DEFAULT timezone('utc'::text, now()),
    updated_at timestamp with time zone DEFAULT timezone('utc'::text, now())
);