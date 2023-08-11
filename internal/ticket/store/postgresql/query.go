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
