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
		t.image_uri,
		t.nomor_tiket,
		t.checkin_status,
		t.checkin_nomor_tiket,
		t.create_time,
		t.update_time
	FROM
		transaction t
	%s
	ORDER BY 
		t.create_time DESC
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
const queryUpdateTransaction = `
	UPDATE
		transaction
	SET
		nama = :nama,
		jenis_kelamin = :jenis_kelamin,
		nomor_identitas = :nomor_identitas,
		asal_institusi = :asal_institusi,
		domisili = :domisili,
		email = :email,
		nomor_telepon = :nomor_telepon,
		line_id = :line_id,
		instagram = :instagram,
		jumlah_tiket = :jumlah_tiket,
		total_harga = :total_harga,
		tanggal = :tanggal,
		order_id = :order_id,
		status_payment = :status_payment,
		image_uri = :image_uri,
		nomor_tiket = ARRAY[:nomor_tiket],
		checkin_status = :checkin_status,
		update_time = :update_time
		%s
	WHERE
		id = :id
`
