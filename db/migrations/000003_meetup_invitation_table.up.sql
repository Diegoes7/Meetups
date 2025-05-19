CREATE TABLE meetup_invitations (
    id BIGSERIAL PRIMARY KEY,
    meetup_id BIGSERIAL REFERENCES meetups(id) ON DELETE CASCADE,
    user_id BIGSERIAL REFERENCES users(id) ON DELETE CASCADE,
    status TEXT NOT NULL DEFAULT 'pending',
    UNIQUE (meetup_id, user_id)
);
