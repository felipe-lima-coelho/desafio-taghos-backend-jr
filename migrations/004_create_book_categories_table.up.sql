CREATE TABLE IF NOT EXISTS book_categories (
    book_id CHAR(36) NOT NULL,
    category_id CHAR(36) NOT NULL,
    PRIMARY KEY (book_id, category_id),
    CONSTRAINT fk_book
        FOREIGN KEY (book_id)
        REFERENCES books (id)
        ON DELETE CASCADE,
    CONSTRAINT fk_category
        FOREIGN KEY (category_id)
        REFERENCES categories (id)
        ON DELETE CASCADE
);