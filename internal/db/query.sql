-- name: InsertAntrian :exec
insert into antrian (loket, urut)
select
    sqlc.arg(loket) AS loket,
    ifnull(max(urut),0)+1 as urut
from antrian
where date(tanggal) = current_date()
  and antrian.loket = sqlc.arg(loket);

-- name: GetAllLoket :many
select max(urut) as urut, loket from antrian where tanggal = current_date() and panggil = 0 group by loket;

-- name: GetAntrian :one
select id, loket, urut from antrian
where tanggal = current_date() and id = ?;

-- name: ListAntrian :many
select id,loket, urut from antrian
where tanggal = current_date() and loket = ? and panggil = 0
order by urut;

-- name: SelesaiAntrian :exec
update antrian set panggil=1 where id = ?;
