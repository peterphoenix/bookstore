-- Insert authors
INSERT INTO "author" (id, name) VALUES
    (uuid_generate_v4(), 'J.K. Rowling'),
    (uuid_generate_v4(), 'George R.R. Martin'),
    (uuid_generate_v4(), 'J.R.R. Tolkien'),
    (uuid_generate_v4(), 'Neil Gaiman'),
    (uuid_generate_v4(), 'Terry Pratchett'),
    (uuid_generate_v4(), 'Isaac Asimov'),
    (uuid_generate_v4(), 'Arthur C. Clarke'),
    (uuid_generate_v4(), 'Douglas Adams'),
    (uuid_generate_v4(), 'Agatha Christie'),
    (uuid_generate_v4(), 'Stephen King'),
    (uuid_generate_v4(), 'Margaret Atwood'),
    (uuid_generate_v4(), 'Philip K. Dick'),
    (uuid_generate_v4(), 'Dan Brown'),
    (uuid_generate_v4(), 'Suzanne Collins'),
    (uuid_generate_v4(), 'Harper Lee'),
    (uuid_generate_v4(), 'F. Scott Fitzgerald'),
    (uuid_generate_v4(), 'Ernest Hemingway'),
    (uuid_generate_v4(), 'Mark Twain'),
    (uuid_generate_v4(), 'Leo Tolstoy'),
    (uuid_generate_v4(), 'Fyodor Dostoevsky'),
    (uuid_generate_v4(), 'Peter Straub'),
    (uuid_generate_v4(), 'Stephen Baxter'),
    (uuid_generate_v4(), 'Rob Wilkins'),
    (uuid_generate_v4(), 'Owen King'),
    (uuid_generate_v4(), 'Frank Herbert'),
    (uuid_generate_v4(), 'Aldous Huxley'),
    (uuid_generate_v4(), 'J.D. Salinger'),
    (uuid_generate_v4(), 'Cormac McCarthy'),
    (uuid_generate_v4(), 'Kurt Vonnegut'),
    (uuid_generate_v4(), 'Brent Weeks'),
    (uuid_generate_v4(), 'Haruki Murakami'),
    (uuid_generate_v4(), 'George Orwell'),
    (uuid_generate_v4(), 'Mary Shelley'),
    (uuid_generate_v4(), 'Oscar Wilde'),
    (uuid_generate_v4(), 'Bram Stoker'),
    (uuid_generate_v4(), 'Khaled Hosseini'),
    (uuid_generate_v4(), 'Yann Martel'),
    (uuid_generate_v4(), 'Markus Zusak'),
    (uuid_generate_v4(), 'Paulo Coelho'),
    (uuid_generate_v4(), 'Erin Morgenstern'),
    (uuid_generate_v4(), 'Gillian Flynn'),
    (uuid_generate_v4(), 'Brandon Sanderson'),
    (uuid_generate_v4(), 'Jane Austen'),
    (uuid_generate_v4(), 'John Green'),
    (uuid_generate_v4(), 'James Dashner');

-- Insert books
INSERT INTO "book" (id, title, price, image_url, isbn, description) VALUES
    (uuid_generate_v4(), 'Harry Potter and the Philosophers Stone', 299850.00, 'https://example.com/harry_potter.jpg', '9780747532743', 'A young wizard embarks on an adventure.'),
    (uuid_generate_v4(), 'A Game of Thrones', 374850.00, 'https://example.com/game_of_thrones.jpg', '9780553103540', 'A tale of power and betrayal in the Seven Kingdoms.'),
    (uuid_generate_v4(), 'The Hobbit', 224850.00, 'https://example.com/hobbit.jpg', '9780261102217', 'Bilbo Baggins goes on an unexpected adventure.'),
    (uuid_generate_v4(), 'Good Omens', 284850.00, 'https://example.com/good_omens.jpg', '9780060853983', 'A comedic take on the apocalypse by Neil Gaiman and Terry Pratchett.'),
    (uuid_generate_v4(), 'American Gods', 329850.00, 'https://example.com/american_gods.jpg', '9780060558123', 'A novel about the power of myth and gods in modern America by Neil Gaiman.'),
    (uuid_generate_v4(), 'Foundation', 269850.00, 'https://example.com/foundation.jpg', '9780553293357', 'A science fiction masterpiece by Isaac Asimov.'),
    (uuid_generate_v4(), '2001: A Space Odyssey', 239850.00, 'https://example.com/2001_odyssey.jpg', '9780451457998', 'A classic space exploration novel by Arthur C. Clarke.'),
    (uuid_generate_v4(), 'The Hitchhikers Guide to the Galaxy', 254850.00, 'https://example.com/hitchhiker.jpg', '9780345391803', 'A humorous science fiction novel by Douglas Adams.'),
    (uuid_generate_v4(), 'The Talisman', 314850.00, 'https://example.com/talisman.jpg', '9780345450661', 'A dark fantasy novel co-authored by Stephen King and Peter Straub.'),
    (uuid_generate_v4(), 'The Long Earth', 239850.00, 'https://example.com/long_earth.jpg', '9780062067777', 'A science fiction novel co-authored by Terry Pratchett and Stephen Baxter.'),
    (uuid_generate_v4(), 'Murder on the Orient Express', 194850.00, 'https://example.com/orient_express.jpg', '9780062693662', 'A classic mystery novel by Agatha Christie.'),
    (uuid_generate_v4(), 'The Shining', 314850.00, 'https://example.com/shining.jpg', '9780307743657', 'A horror novel by Stephen King.'),
    (uuid_generate_v4(), 'The Handmaids Tale', 284850.00, 'https://example.com/handmaids_tale.jpg', '9780385490818', 'A dystopian novel by Margaret Atwood.'),
    (uuid_generate_v4(), 'Do Androids Dream of Electric Sheep?', 232350.00, 'https://example.com/androids_sheep.jpg', '9780345404473', 'A science fiction novel by Philip K. Dick.'),
    (uuid_generate_v4(), 'The Da Vinci Code', 337350.00, 'https://example.com/da_vinci_code.jpg', '9780307474278', 'A mystery-thriller novel by Dan Brown.'),
    (uuid_generate_v4(), 'The Hunger Games', 224850.00, 'https://example.com/hunger_games.jpg', '9780439023481', 'A dystopian novel by Suzanne Collins.'),
    (uuid_generate_v4(), 'To Kill a Mockingbird', 164850.00, 'https://example.com/to_kill_a_mockingbird.jpg', '9780061120084', 'A classic novel about racial injustice by Harper Lee.'),
    (uuid_generate_v4(), 'Sleeping Beauties', 299850.00, 'https://example.com/sleeping_beauties.jpg', '9781501163401', 'A fantasy/horror novel co-authored by Stephen King and Owen King.'),
    (uuid_generate_v4(), 'The Mongoliad', 254850.00, 'https://example.com/mongoliad.jpg', '9781612182360', 'A historical fantasy novel co-authored by Neal Stephenson, Greg Bear, and others.'),
    (uuid_generate_v4(), 'The Great Gatsby', 179850.00, 'https://example.com/great_gatsby.jpg', '9780743273565', 'A novel about wealth and tragedy by F. Scott Fitzgerald.'),
    (uuid_generate_v4(), 'The Old Man and the Sea', 149850.00, 'https://example.com/old_man_sea.jpg', '9780684801223', 'A short novel about an epic battle between a fisherman and a giant marlin by Ernest Hemingway.'),
    (uuid_generate_v4(), 'Adventures of Huckleberry Finn', 187350.00, 'https://example.com/huck_finn.jpg', '9780486280615', 'A classic adventure novel by Mark Twain.'),
    (uuid_generate_v4(), 'War and Peace', 374850.00, 'https://example.com/war_and_peace.jpg', '9780199232765', 'A historical novel by Leo Tolstoy.'),
    (uuid_generate_v4(), 'Crime and Punishment', 292350.00, 'https://example.com/crime_and_punishment.jpg', '9780143058144', 'A psychological drama by Fyodor Dostoevsky.'),
    (uuid_generate_v4(), 'Dune', 449850.00, 'https://example.com/dune.jpg', '9780441013593', 'A science fiction novel by Frank Herbert.'),
    (uuid_generate_v4(), 'Brave New World', 209850.00, 'https://example.com/brave_new_world.jpg', '9780060850524', 'A dystopian novel by Aldous Huxley.'),
    (uuid_generate_v4(), 'The Catcher in the Rye', 187350.00, 'https://example.com/catcher_in_the_rye.jpg', '9780316769488', 'A novel about teenage alienation by J.D. Salinger.'),
    (uuid_generate_v4(), 'The Road', 224850.00, 'https://example.com/the_road.jpg', '9780307387899', 'A post-apocalyptic novel by Cormac McCarthy.'),
    (uuid_generate_v4(), 'Slaughterhouse-Five', 209850.00, 'https://example.com/slaughterhouse_five.jpg', '9780385333849', 'A novel blending history and sci-fi by Kurt Vonnegut.'),
    (uuid_generate_v4(), 'The Brothers Karamazov', 329850.00, 'https://example.com/brothers_karamazov.jpg', '9780374528379', 'A philosophical novel by Fyodor Dostoevsky.'),
    (uuid_generate_v4(), 'Raising Steam', 224850.00, 'https://example.com/raising_steam.jpg', '9780857522276', 'A Discworld novel co-authored by Terry Pratchett and Rob Wilkins.'),
    (uuid_generate_v4(), 'The Black Prism', 299850.00, 'https://example.com/black_prism.jpg', '9780316075551', 'An epic fantasy co-authored by Brent Weeks and Brandon Sanderson.'),
    (uuid_generate_v4(), 'The Wind-Up Bird Chronicle', 284850.00, 'https://example.com/wind_up_bird.jpg', '9780679775430', 'A surreal mystery novel by Haruki Murakami.'),
    (uuid_generate_v4(), '1984', 224850.00, 'https://example.com/1984.jpg', '9780451524935', 'A dystopian novel by George Orwell.'),
    (uuid_generate_v4(), 'Animal Farm', 149850.00, 'https://example.com/animal_farm.jpg', '9780451526342', 'A political allegory by George Orwell.'),
    (uuid_generate_v4(), 'Frankenstein', 194850.00, 'https://example.com/frankenstein.jpg', '9780486282114', 'A gothic novel by Mary Shelley.'),
    (uuid_generate_v4(), 'The Picture of Dorian Gray', 164850.00, 'https://example.com/dorian_gray.jpg', '9780141439570', 'A novel about moral corruption by Oscar Wilde.'),
    (uuid_generate_v4(), 'Dracula', 209850.00, 'https://example.com/dracula.jpg', '9780141439846', 'A gothic novel by Bram Stoker.'),
    (uuid_generate_v4(), 'The Kite Runner', 239850.00, 'https://example.com/kite_runner.jpg', '9781594631931', 'A novel about friendship and redemption by Khaled Hosseini.'),
    (uuid_generate_v4(), 'Life of Pi', 224850.00, 'https://example.com/life_of_pi.jpg', '9780156027328', 'A novel about survival at sea by Yann Martel.'),
    (uuid_generate_v4(), 'The Book Thief', 269850.00, 'https://example.com/book_thief.jpg', '9780375842207', 'A novel set during WWII by Markus Zusak.'),
    (uuid_generate_v4(), 'The Alchemist', 194850.00, 'https://example.com/alchemist.jpg', '9780061122415', 'A philosophical novel by Paulo Coelho.'),
    (uuid_generate_v4(), 'The Night Circus', 299850.00, 'https://example.com/night_circus.jpg', '9780307744432', 'A fantasy novel by Erin Morgenstern.'),
    (uuid_generate_v4(), 'Gone Girl', 254850.00, 'https://example.com/gone_girl.jpg', '9780307588364', 'A psychological thriller by Gillian Flynn.'),
    (uuid_generate_v4(), 'Pride and Prejudice', 149850.00, 'https://example.com/pride_and_prejudice.jpg', '9780141439518', 'A classic novel by Jane Austen.'),
    (uuid_generate_v4(), 'The Fault in Our Stars', 209850.00, 'https://example.com/fault_in_our_stars.jpg', '9780525478812', 'A young adult novel by John Green.'),
    (uuid_generate_v4(), 'The Maze Runner', 224850.00, 'https://example.com/maze_runner.jpg', '9780385737951', 'A dystopian novel by James Dashner.');


-- Insert into book_author table for the corresponding author

-- "Harry Potter and the Philosopher's Stone" by J.K. Rowling
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'Harry Potter and the Philosophers Stone' AND a.name = 'J.K. Rowling';

-- "A Game of Thrones" by George R.R. Martin
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'A Game of Thrones' AND a.name = 'George R.R. Martin';

-- "The Hobbit" by J.R.R. Tolkien
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'The Hobbit' AND a.name = 'J.R.R. Tolkien';

-- "Good Omens" by Neil Gaiman and Terry Pratchett
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'Good Omens' AND a.name = 'Neil Gaiman';

INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'Good Omens' AND a.name = 'Terry Pratchett';

-- "American Gods" by Neil Gaiman
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'American Gods' AND a.name = 'Neil Gaiman';

-- "Foundation" by Isaac Asimov
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'Foundation' AND a.name = 'Isaac Asimov';

-- "2001: A Space Odyssey" by Arthur C. Clarke
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = '2001: A Space Odyssey' AND a.name = 'Arthur C. Clarke';

-- "The Hitchhiker's Guide to the Galaxy" by Douglas Adams
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'The Hitchhikers Guide to the Galaxy' AND a.name = 'Douglas Adams';

-- "The Talisman" by Stephen King and Peter Straub
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'The Talisman' AND a.name = 'Stephen King';

INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'The Talisman' AND a.name = 'Peter Straub';

-- "The Long Earth" by Terry Pratchett and Stephen Baxter
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'The Long Earth' AND a.name = 'Terry Pratchett';

INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'The Long Earth' AND a.name = 'Stephen Baxter';

-- "Murder on the Orient Express" by Agatha Christie
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'Murder on the Orient Express' AND a.name = 'Agatha Christie';

-- "The Shining" by Stephen King
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'The Shining' AND a.name = 'Stephen King';

-- "The Handmaid's Tale" by Margaret Atwood
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'The Handmaids Tale' AND a.name = 'Margaret Atwood';

-- "Do Androids Dream of Electric Sheep?" by Philip K. Dick
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'Do Androids Dream of Electric Sheep?' AND a.name = 'Philip K. Dick';

-- "The Da Vinci Code" by Dan Brown
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'The Da Vinci Code' AND a.name = 'Dan Brown';

-- "The Hunger Games" by Suzanne Collins
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'The Hunger Games' AND a.name = 'Suzanne Collins';

-- "To Kill a Mockingbird" by Harper Lee
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'To Kill a Mockingbird' AND a.name = 'Harper Lee';

-- "Sleeping Beauties" by Stephen King and Owen King
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'Sleeping Beauties' AND a.name = 'Stephen King';

INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'Sleeping Beauties' AND a.name = 'Owen King';

-- "The Mongoliad" (multi-author) by Neal Stephenson, Greg Bear
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'The Mongoliad' AND a.name = 'Neal Stephenson';

INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'The Mongoliad' AND a.name = 'Greg Bear';

-- "The Great Gatsby" by F. Scott Fitzgerald
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'The Great Gatsby' AND a.name = 'F. Scott Fitzgerald';

-- "The Old Man and the Sea" by Ernest Hemingway
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'The Old Man and the Sea' AND a.name = 'Ernest Hemingway';

-- "Adventures of Huckleberry Finn" by Mark Twain
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'Adventures of Huckleberry Finn' AND a.name = 'Mark Twain';

-- "War and Peace" by Leo Tolstoy
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'War and Peace' AND a.name = 'Leo Tolstoy';

-- "Crime and Punishment" by Fyodor Dostoevsky
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'Crime and Punishment' AND a.name = 'Fyodor Dostoevsky';

-- "Dune" by Frank Herbert (Assuming Frank Herbert is added in the author table)
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'Dune' AND a.name = 'Frank Herbert';

-- "Brave New World" by Aldous Huxley (Assuming Aldous Huxley is added in the author table)
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'Brave New World' AND a.name = 'Aldous Huxley';

-- "The Catcher in the Rye" by J.D. Salinger (Assuming J.D. Salinger is added in the author table)
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'The Catcher in the Rye' AND a.name = 'J.D. Salinger';

-- "The Road" by Cormac McCarthy (Assuming Cormac McCarthy is added in the author table)
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'The Road' AND a.name = 'Cormac McCarthy';

-- "Slaughterhouse-Five" by Kurt Vonnegut (Assuming Kurt Vonnegut is added in the author table)
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'Slaughterhouse-Five' AND a.name = 'Kurt Vonnegut';

-- "The Brothers Karamazov" by Fyodor Dostoevsky
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'The Brothers Karamazov' AND a.name = 'Fyodor Dostoevsky';

-- "Raising Steam" by Terry Pratchett and Rob Wilkins
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'Raising Steam' AND a.name = 'Terry Pratchett';

INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'Raising Steam' AND a.name = 'Rob Wilkins';

-- "The Black Prism" by Brent Weeks (Assuming Brent Weeks is added in the author table)
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'The Black Prism' AND a.name = 'Brent Weeks';

-- "The Wind-Up Bird Chronicle" by Haruki Murakami (Assuming Haruki Murakami is added in the author table)
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'The Wind-Up Bird Chronicle' AND a.name = 'Haruki Murakami';

-- "1984" by George Orwell (Assuming George Orwell is added in the author table)
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = '1984' AND a.name = 'George Orwell';

-- "Animal Farm" by George Orwell
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'Animal Farm' AND a.name = 'George Orwell';

-- "Frankenstein" by Mary Shelley (Assuming Mary Shelley is added in the author table)
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'Frankenstein' AND a.name = 'Mary Shelley';

-- "The Picture of Dorian Gray" by Oscar Wilde (Assuming Oscar Wilde is added in the author table)
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'The Picture of Dorian Gray' AND a.name = 'Oscar Wilde';

-- "Dracula" by Bram Stoker (Assuming Bram Stoker is added in the author table)
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'Dracula' AND a.name = 'Bram Stoker';

-- "The Kite Runner" by Khaled Hosseini (Assuming Khaled Hosseini is added in the author table)
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'The Kite Runner' AND a.name = 'Khaled Hosseini';

-- "Life of Pi" by Yann Martel (Assuming Yann Martel is added in the author table)
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'Life of Pi' AND a.name = 'Yann Martel';

-- "The Book Thief" by Markus Zusak (Assuming Markus Zusak is added in the author table)
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'The Book Thief' AND a.name = 'Markus Zusak';

-- "The Alchemist" by Paulo Coelho (Assuming Paulo Coelho is added in the author table)
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'The Alchemist' AND a.name = 'Paulo Coelho';

-- "The Night Circus" by Erin Morgenstern (Assuming Erin Morgenstern is added in the author table)
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'The Night Circus' AND a.name = 'Erin Morgenstern';

-- "Gone Girl" by Gillian Flynn (Assuming Gillian Flynn is added in the author table)
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'Gone Girl' AND a.name = 'Gillian Flynn';

-- "Pride and Prejudice" by Jane Austen (Assuming Jane Austen is added in the author table)
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'Pride and Prejudice' AND a.name = 'Jane Austen';

-- "The Fault in Our Stars" by John Green (Assuming John Green is added in the author table)
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'The Fault in Our Stars' AND a.name = 'John Green';

-- "The Maze Runner" by James Dashner (Assuming James Dashner is added in the author table)
INSERT INTO "book_author" (book_id, author_id)
SELECT b.id, a.id FROM "book" b, "author" a
WHERE b.title = 'The Maze Runner' AND a.name = 'James Dashner';
