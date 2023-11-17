package postgresql

const queryCreateMainEvent = `
	INSERT INTO
		mainevent
	(
		nama,
		disabilitas,
		nomor_identitas,
		asal_institusi,
		email,
		nomor_telepon,
		jumlah_tiket,
		total_harga,
		order_id,
		type,
		status,
		create_time
	) VALUES (
		:nama,
		:disabilitas,
		:nomor_identitas,
		:asal_institusi,
		:email,
		:nomor_telepon,
		:jumlah_tiket,
		:total_harga,
		:order_id,
		:type,
		:status,
		:create_time
	) RETURNING
		id
`

const queryGetMainEvent = `
	SELECT
		m.id,
		m.nama,
		m.disabilitas,
		m.nomor_identitas,
		m.asal_institusi,
		m.email,
		m.nomor_telepon,
		m.jumlah_tiket,
		m.total_harga,
		m.order_id,
		m.type,
		m.status,
		m.image_uri,
		m.nomor_tiket,
		m.checkin_status,
		m.checkin_nomor_tiket,
		m.create_time,
		m.update_time
	FROM
		mainevent m
	%s
	ORDER BY 
		m.create_time DESC
`

const queryuDeleteMainEventByEmail = `
	DELETE FROM 
		mainevent
	WHERE
		email = :email
	AND
		status = :status
`
const queryUpdateMainEvent = `
	UPDATE
		mainevent
	SET
		nama = :nama,
		disabilitas = :disabilitas,
		nomor_identitas = :nomor_identitas,
		asal_institusi = :asal_institusi,
		email = :email,
		nomor_telepon = :nomor_telepon,
		jumlah_tiket = :jumlah_tiket,
		total_harga = :total_harga,
		order_id = :order_id,
		status = :status,
		image_uri = :image_uri,
		nomor_tiket = ARRAY[:nomor_tiket],
		checkin_status = :checkin_status,
		update_time = :update_time
		%s
	WHERE
		id = :id
`
