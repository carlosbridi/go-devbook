insert into usuarios(nome, nick, email, senha )
values 
("Usuário 1", "usuario_1", "usuario1@gmail.com", "$2a$10$FKkeTK8eGar99gpeMHjuAeUeH.mmoXDoU8QjU7ElsngdMSAkGD6zW"),
("Usuário 2", "usuario_2", "usuario2@gmail.com", "$2a$10$FKkeTK8eGar99gpeMHjuAeUeH.mmoXDoU8QjU7ElsngdMSAkGD6zW"),
("Usuário 3", "usuario_3", "usuario3@gmail.com", "$2a$10$FKkeTK8eGar99gpeMHjuAeUeH.mmoXDoU8QjU7ElsngdMSAkGD6zW");

insert into seguidores(usuario_id, seguidor_id )
values 
(1, 2),
(1, 3),
(2, 1),
(2, 3);

