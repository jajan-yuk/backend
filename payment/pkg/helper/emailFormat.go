package helper

import "fmt"

func RegisterEmail(name, role, url string) string {
	body := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
			<head>
				<title>Jajan Yuk - Email Verification</title>
			</head>
			<body>
				<h2>Jajan Yuk - Email Verification</h2>
				<p>Hi, %s</p>
				<p>Terima kasih telah mendaftar di Jajan Yuk Apps sebagai %s</p>
				<p>Untuk melanjutkan proses registrasi, silahkan klik link berikut</p>
				<p><a href="%s">Klik disini</a></p>
				</br>
				<p>Terima kasih,</p>
				<p>Jajan Yuk Team</p>
			</body>
		</html>
	`, name, role, url)

	return body
}

func SuccesActivateEmail(name string) string {
	body := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
			<head>
				<title>Jajan Yuk - Success Activate Account</title>
			</head>
			<body>
				<h2>Welcome - Jajan Yuk Apps</h2>
				<p>Hi, %s</p>
				<p>Akun anda sudah terdaftar pada platform kami, silahkan login untuk melanjutkan</p>
				</br>
				<p>Terima kasih,</p>
				<p>Jajan Yuk Team</p>
			</body>
		</html>
	`, name)

	return body
}

func ResetPassword(name, url, tanggal, expire string) string {
	body := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
			<head>
				<title>Jajan Yuk - Reset Password</title>
			</head>
			<body>
				<h2>Request Reset Password - Jajan Yuk Apps</h2>
				<p>Hi, %s</p>
				<p>Permintaan Reset Password Telah anda lakukan pada <b>%s</b></p>
				<p>Klik link dibawah ini sebelum expired pada <b>%s</b></p>
				</br>
				<p><a href="%s">Klik disini</a></p>
				</br>
				<p>Terima kasih,</p>
				<p>Jajan Yuk Team</p>
			</body>
		</html>
	`, name, tanggal, expire, url)

	return body
}
