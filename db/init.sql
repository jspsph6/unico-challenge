USE fair;

CREATE TABLE IF NOT EXISTS fair (
    id int not null auto_increment,
    external_id int not null,
    setcens int not null,
    fair_name varchar(255) not null,
    register varchar(255) not null,
    code_district int not null,
    code_subprefecture int not null,
    subprefecture varchar(255) not null,
    areap int not null,
    district_name varchar(255) not null,
    region5 varchar(255) not null,
    region8 varchar(255) not null,
    street varchar(255) not null,
    number varchar(255) null,
    neighborhood varchar(255) not null,
    reference varchar(255) null,
    coordinates point not null,
    PRIMARY KEY (id)
);

SET character_set_client = utf8;
SET character_set_connection = utf8;
SET character_set_results = utf8;
SET collation_connection = utf8_general_ci;