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
		total_harga,
		tanggal,
		order_id,
		status_payment,
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
		:total_harga,
		:tanggal,
		:order_id,
		:status_payment,
		:response_midtrans,
		:create_time
	) RETURNING
		id
`

const queryGetTransaction = `
	SELECT
		t.id,
		t.nama,
		t.jenis_kelamin,
		t.nomor_identitas,
		t.asal_institusi,
		t.domisili,
		t.email,
		t.nomor_telepon,
		t.line_id,
		t.instagram,
		t.jumlah_tiket,
		t.total_harga,
		t.tanggal,
		t.order_id,
		t.status_payment,
		t.response_midtrans,
		t.nomor_tiket,
		t.checkin_status,
		t.checkin_nomor_tiket,
		t.create_time,
		t.update_time
	FROM
		transaction t
	%s
`

const queryuDeleteTransactionByEmail = `
	DELETE FROM 
		transaction
	WHERE
		email = :email
	AND
		tanggal = :tanggal
	AND
		status_payment = :status_payment
`
