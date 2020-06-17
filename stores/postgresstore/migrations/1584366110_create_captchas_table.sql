CREATE TABLE captchas (
    id VARCHAR(64) NOT NULL,
    category VARCHAR(32) NOT NULL,
    answer VARCHAR(32) NOT NULL,
    created_at BIGINT NOT NULL,
    expires_in BIGINT NOT null,
    PRIMARY KEY(id, category)
);

CREATE INDEX captchas_expires_on_idx ON captchas (expires_in);