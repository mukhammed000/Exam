CREATE TABLE languages (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE user_progress (
    user_id UUID REFERENCES users(id),
    language_id UUID REFERENCES languages(id),
    level VARCHAR(255),
    xp BIGINT,
    lessons_completed BIGINT,
    vocabulary_learned BIGINT,
    grammar_rules_mastered BIGINT,
    listening_comprehension BIGINT,
    speaking_fluency BIGINT,
    PRIMARY KEY (user_id, language_id)
);

CREATE TABLE daily_progress (
    user_id UUID REFERENCES users(id),
    date DATE,
    xp_earned BIGINT,
    lessons_completed BIGINT,
    new_words_learned BIGINT,
    minutes_practiced BIGINT,
    streak_days BIGINT,
    PRIMARY KEY (user_id, date)
);

CREATE TABLE weekly_progress (
    user_id UUID REFERENCES users(id),
    week_start DATE,
    week_end DATE,
    total_xp_earned BIGINT,
    lessons_completed BIGINT,
    new_words_learned BIGINT,
    total_minutes_practiced BIGINT,
    most_active_day DATE,
    PRIMARY KEY (user_id, week_start, week_end)
);

CREATE TABLE monthly_progress (
    user_id UUID REFERENCES users(id),
    month DATE,
    total_xp_earned BIGINT,
    lessons_completed BIGINT,
    new_words_learned BIGINT,
    total_minutes_practiced BIGINT,
    most_improved_skill VARCHAR(255),
    PRIMARY KEY (user_id, month)
);

CREATE TABLE achievements (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    title VARCHAR(255),
    description TEXT,
    earned_at TIMESTAMPTZ
);

CREATE TABLE leaderboard (
    language_id UUID REFERENCES languages(id),
    user_id UUID REFERENCES users(id),
    rank BIGINT,
    xp BIGINT,
    level VARCHAR(255),
    PRIMARY KEY (language_id, user_id)
);

CREATE TABLE skills (
    user_id UUID REFERENCES users(id),
    language_id UUID REFERENCES languages(id),
    name VARCHAR(255),
    level BIGINT,
    progress_to_next_level BIGINT,
    words_learned BIGINT,
    rules_mastered BIGINT,
    comprehension_rate BIGINT,
    fluency_rate BIGINT,
    PRIMARY KEY (user_id, language_id, name)
);