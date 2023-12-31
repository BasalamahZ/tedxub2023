package postgresql

const queryCreateTicket = `
	INSERT INTO
		ticket
	(
		nama,
		jenis_kelamin,
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
		:jenis_kelamin,
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
		jenis_kelamin,
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

const queryCountEmail = `
		SELECT
			COUNT(email)
		FROM
			ticket
		WHERE
			email = :email
`

const queryCountNumberIdentity = `
		SELECT
			COUNT(nomor_identitas)
		FROM
			ticket
		WHERE
			nomor_identitas = :nomor_identitas
`
