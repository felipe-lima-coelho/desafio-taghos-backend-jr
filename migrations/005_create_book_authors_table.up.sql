CREATE TABLE IF NOT EXISTS book_authors (
    book_id CHAR(36) NOT NULL,
    author_id CHAR(36) NOT NULL,
    PRIMARY KEY (book_id, author_id),
    CONSTRAINT fk_book
        FOREIGN KEY (book_id)
        REFERENCES books (id)
        ON DELETE CASCADE,
    CONSTRAINT fk_author
        FOREIGN KEY (author_id)
        REFERENCES authors (id)
        ON DELETE CASCADE
);