create table antrian (
    id int primary key auto_increment,
    tanggal datetime not null default current_date(),
    loket int not null default 0,
    urut int not null,
    panggil int not null default 0
);

create table user (
    id int primary key auto_increment,
    nama varchar(10) not null,
    password varchar(255) not null
);
