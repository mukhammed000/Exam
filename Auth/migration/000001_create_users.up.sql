-- Table for storing user registration information
CREATE TABLE users (
    id UUID PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    role VARCHAR(255) NOT NULL,
    native_language VARCHAR(255) NOT NULL,
);

-- Table for storing login session tokens
CREATE TABLE login_sessions (
    user_id UUID PRIMARY KEY,
    token VARCHAR(255) NOT NULL,
    expires_at TIMESTAMP NOT NULL
);

-- Table for storing user profile information
CREATE TABLE user_profiles (
    user_id UUID PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    native_language VARCHAR(255) NOT NULL,
    learning_languages VARCHAR[] NOT NULL,
    created_at TIMESTAMP NOT NULL
);

-- Table for storing refresh tokens
CREATE TABLE refresh_tokens (
    token VARCHAR(255) PRIMARY KEY,
    success BOOLEAN NOT NULL,
    message TEXT,
    user_id UUID,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Table for storing user settings
CREATE TABLE user_settings (
    user_id UUID PRIMARY KEY,
    daily_goal VARCHAR(255) NOT NULL,
    notifications_enabled BOOLEAN NOT NULL,
    preferred_learning_time VARCHAR(255) NOT NULL,
    difficulty_level VARCHAR(255) NOT NULL
);

-- Table for storing password change requests
CREATE TABLE password_changes (
    user_id UUID PRIMARY KEY,
    current_password VARCHAR(255) NOT NULL,
    new_password VARCHAR(255) NOT NULL
);

-- Table for storing password reset requests
CREATE TABLE password_resets (
    email VARCHAR(255) PRIMARY KEY,
    user_id UUID,
    reset_token VARCHAR(255) NOT NULL,
    new_password VARCHAR(255) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
