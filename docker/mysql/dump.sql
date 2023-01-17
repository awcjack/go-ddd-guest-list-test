CREATE TABLE `tables` (
  -- Auto increment id as table id
  `id` INT NOT NULL auto_increment,
  -- Table capacity
  `capacity` INT NOT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `guests` (
  -- Current Longest Personal name have 747 characters, Ref: https://www.guinnessworldrecords.com/world-records/67285-longest-personal-name
  -- for string more than 255 characters, it always use 2 bytes to indicate the size more than 255
  -- so set the maximum size of name to maximum size of varchar
  -- varchar will not occupy extra space with the string is shorter than the maximum size (max index size is 3072) so it won't waste the space
  -- Ref: https://dev.mysql.com/doc/refman/5.7/en/char.html
  `name` varchar(3072) NOT NULL,
  -- Expected guest number when adding guest to guest list
  `guest_number` INT NOT NULL,
  -- Corresponding table for guest
  `table_id` INT NOT NULL,
  -- Arrival status of guest
  `arrived` BOOLEAN,
  -- Guest number when arrived
  `arrived_number` INT,
  -- Arrived time
  `time_arrived` DATETIME,

  -- Unique name as id (Same name is not allowed)
  PRIMARY KEY (`name`),

  -- Foreign key that make sure table exist
  CONSTRAINT table_table_id FOREIGN KEY (`table_id`)
    REFERENCES `tables`(`id`)
);