CREATE TABLE IF NOT EXISTS cakes(
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    title varchar (255),
    description varchar (255),
    rating float,
    image varchar (255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

