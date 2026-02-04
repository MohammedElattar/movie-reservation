CREATE TABLE IF NOT EXISTS access_tokens(
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    token text NOT NULL,
    expires_at TIMESTAMPTZ,

    CONSTRAINT fk_access_tokens_user_id
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);
