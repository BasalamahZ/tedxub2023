package postgresql

const queryCreateTicket = `
	INSERT INTO
		ticket
	(
		nama,
		nomor_identitas,
		asal_institusi,
		domisili,
		email,
		nomor_telepon,
		line_id,
		instagram,
		create_time
	) VALUES (
		:nama,
		:nomor_identitas,
		:asal_institusi,
		:domisili,
		:email,
		:nomor_telepon,
		:line_id,
		:instagram,
		:create_time
	) RETURNING
		nama
`

const queryUpdateTicket = `
	UPDATE 
		ticket
	SET 
		status = :status, 
		nomor_tiket = :nomor_tiket,
		update_time = :update_time
	WHERE 
		id = :id
`

const queryGetAllTicket = `
	SELECT
		id,
		nama,
		nomor_identitas,
		asal_institusi,
		domisili,
		email,
		nomor_telepon,
		line_id,
		instagram,
		status,
		nomor_tiket,
		create_time
	FROM
		ticket
`
