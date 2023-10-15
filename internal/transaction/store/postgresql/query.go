package postgresql

const queryCreateTransaction = `
	INSERT INTO
		transaction
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
		jumlah_tiket,
		harga,
		nomor_tiket,
		response_midtrans,
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
		:jumlah_tiket,
		:harga,
		:nomor_tiket,
		:response_midtrans,
		:create_time
	) RETURNING
		id
`

const queryuDeleteTransactionByEmail = `
	DELETE FROM 
		transaction
	WHERE
		email = :email
`
