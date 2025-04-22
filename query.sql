select
	m.id as id_murid,
	m.name,
	p.status as pendidikan_terkahir,
	m.time_create,
	p.time_create as time_update
from
	murid m
left join
	(
	select
		id_murid,
		status,
		time_create
	from
		pendidikan p
	where time_create = (
		select max(time_create)
		from pendidikan p2
		where p2.id_murid = p.id_murid
	)
) as p on p.id_murid = m.id