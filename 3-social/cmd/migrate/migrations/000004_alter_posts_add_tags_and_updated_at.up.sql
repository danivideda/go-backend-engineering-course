ALTER TABLE posts
ADD tags VARCHAR(100) [],
ADD updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW();