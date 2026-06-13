-- name: InsertAntrian :execresult
insert into antrian (urut)
select
    ifnull(max(urut),0)+1 as urut
from antrian where date(tanggal) = current_date() ;
-- select * from antrian where id = (select last_insert_id()) limit 1;

-- name: GetAllLoket :many
select max(urut) as urut, loket from antrian where tanggal = current_date() and panggil = 0 group by loket;

-- name: GetAntrian :one
select id, urut, loket, DATE_FORMAT(tanggal, '%Y-%m-%d') AS tanggal from antrian
where tanggal = current_date() and id = ?;

-- name: ListAntrian :many
select id,loket, urut from antrian
where tanggal = current_date() and loket = ? and panggil = 0
order by urut;

-- name: SelesaiAntrian :exec
update antrian set panggil=1 where id = ?;
