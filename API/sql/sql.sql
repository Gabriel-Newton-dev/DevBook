CREATE DATABASE IF NOT EXISTS devbook;

USe debbook;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id int auto_incremet primary key,
    nome varchar(50) not null, 
    nick varchar(20) not null unique, 
    email varchar(50) not null unique,
    senha varchar(20) not null unique, 
    criadoEm timestamp default current_timestamp()
) ENGINE=INNODB;