package emailattributes

const (
	PaymentConfirmIndonesian string = `
		<!DOCTYPE html>
			<html lang="id">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>{{.Title}}</title>
				<style>
				body {
					font-family: Arial, sans-serif;
					background-color: #E0E0E0;
					margin: 0;
					padding: 0;
				}
				.container {
					max-width: 600px;
					margin: 20px auto;
					background-color: #fff;
					border-radius: 8px;
					box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
					overflow: hidden;
				}
				.header {
					background-color: #007C99;
					color: white;
					text-align: center;
					padding: 15px;
				}
				.header img {
					max-width: 120px;
					margin-bottom: 10px;
					border-radius: 3px;
					padding: 4px;
				}
				.content {
					padding: 20px;
				}
				.content h1 {
					font-size: 18px;
					color: #333;
				}
				.content p {
					font-size: 16px;
					color: #555;
					margin: 5px 0;
				}
				.details {
					background-color: #f9f9f9;
					border: 1px solid #ddd;
					padding: 15px;
					margin: 20px 0;
					border-radius: 8px;
				}
				.details p {
					margin: 8px 0;
					color: #333;
					text-align: center;
				}
				.details strong {
					display: block;
					margin-bottom: 5px;
					font-size: 22px;
				}
				.payment-info {
					text-align: center;
					margin: 20px 0;
				}
				.payment-info p {
					margin: 5px 0;
				}
				.payment-info .amount {
					font-size: 22px;
					color: #007C99;
					font-weight: bold;
				}
				.important {
					background-color: #FFD700;
					padding: 10px;
					font-size: 14px;
					color: #333;
					border-radius: 5px;
					margin-top: 15px;
					text-align: center;
				}
				.footer {
					padding: 15px;
					font-size: 14px;
					text-align: center;
					color: #555;
					background-color: #f4f4f4;
				}
				.footer a {
					color: #007C99;
					text-decoration: none;
				}
				.footer p {
					margin: 5px 0;
				}
				.payment-channel {
					text-align: center;
					margin-bottom: 10px;
				}
				.payment-channel-name {
					font-size: 18px;
					font-weight: bold;
					color: #333;
					margin-bottom: 5px;
				}
				.logo-container {
					padding: 5px 0;
					text-align: center;
				}
				.logo-container img {
					max-width: 60px;
					display: inline-block;
				}
				.payment-button {
					display: block;
					width: 100%;
					max-width: 300px;
					margin: 20px auto 10px;
					padding: 12px 20px;
					background-color: #007C99;
					color: #FFFFFF !important;
					text-align: center;
					text-decoration: none;
					border-radius: 6px;
					font-weight: bold;
					transition: background-color 0.3s ease;
				}
				.payment-button:hover {
					background-color: #006080;
				}
				</style>
			</head>
			<body>
			<div class="container">
				<div class="header">
					<img src="{{.LogoUrl}}" alt="Logo">
				</div>
				<div class="content">
					<h1>Selesaikan pembayaran Anda sebelum {{.PaymentDeadline}}</h1>
					<p>No. Pesanan: <strong>{{.OrderNumber}}</strong></p>
					<p>Yth. {{.CustomerName}},</p>
					<p>Mohon selesaikan pembayaran Anda:</p>
					<div class="details">
						<p>Jumlah Total</p>
						<h1 style="text-align: center; font-size: 22px;" class="amount"><strong>Rp {{.TotalAmount}}</strong></h1>
						<a href="{{.RedirectUrl}}" class="payment-button">Lanjutkan Pembayaran</a>
					</div>
					<div class="important">
						PENTING!!! Pastikan kamu membayar sebelum tenggat waktu ditentukan!
					</div>
				</div>
				<div class="footer">
					<p><strong>Punya pertanyaan?</strong></p>
					<p>Untuk hal mendesak: <a href="tel:{{.CustomerServicePhoneNumber}}">{{.CustomerServicePhoneNumber}}</a></p>
				</div>
			</div>
			</body>
		</html>
	`

	PaymentConfirmAttachmentIndonesian string = `
		<!DOCTYPE html>
			<html lang="id">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>{{.Title}}</title>
				<style>
				body {
					font-family: Arial, sans-serif;
					background-color: #E0E0E0;
					margin: 0;
					padding: 0;
				}
				.container {
					max-width: 600px;
					margin: 20px auto;
					background-color: #fff;
					border-radius: 8px;
					box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
					overflow: hidden;
				}
				.header {
					background-color: #007C99;
					color: white;
					text-align: center;
					padding: 15px;
				}
				.header img {
					max-width: 120px;
					margin-bottom: 10px;
					border-radius: 3px;
					padding: 4px;
				}
				.content {
					padding: 20px;
				}
				.content h1 {
					font-size: 18px;
					color: #333;
				}
				.content p {
					font-size: 16px;
					color: #555;
					margin: 5px 0;
				}
				.details {
					background-color: #f9f9f9;
					border: 1px solid #ddd;
					padding: 15px;
					margin: 20px 0;
					border-radius: 8px;
				}
				.details p {
					margin: 8px 0;
					color: #333;
					text-align: center;
				}
				.details strong {
					display: block;
					margin-bottom: 5px;
					font-size: 22px;
				}
				.payment-info {
					text-align: center;
					margin: 20px 0;
				}
				.payment-info p {
					margin: 5px 0;
				}
				.payment-info .amount {
					font-size: 22px;
					color: #007C99;
					font-weight: bold;
				}
				.important {
					background-color: #FFD700;
					padding: 10px;
					font-size: 14px;
					color: #333;
					border-radius: 5px;
					margin-top: 15px;
					text-align: center;
				}
				.footer {
					padding: 15px;
					font-size: 14px;
					text-align: center;
					color: #555;
					background-color: #f4f4f4;
				}
				.footer a {
					color: #007C99;
					text-decoration: none;
				}
				.footer p {
					margin: 5px 0;
				}
				.payment-channel {
					text-align: center;
					margin-bottom: 10px;
				}
				.payment-channel-name {
					font-size: 18px;
					font-weight: bold;
					color: #333;
					margin-bottom: 5px;
				}
				.logo-container {
					padding: 5px 0;
					text-align: center;
				}
				.logo-container img {
					max-width: 60px;
					display: inline-block;
				}
				.payment-button {
					display: block;
					width: 100%;
					max-width: 300px;
					margin: 20px auto 10px;
					padding: 12px 20px;
					background-color: #007C99;
					color: #FFFFFF !important;
					text-align: center;
					text-decoration: none;
					border-radius: 6px;
					font-weight: bold;
					transition: background-color 0.3s ease;
				}
				.payment-button:hover {
					background-color: #006080;
				}
				</style>
			</head>
			<body>
			<div class="container">
				<div class="header">
					<img src="{{.LogoUrl}}" alt="Logo">
				</div>
				<div class="content">
					<h1>Selesaikan pembayaran Anda sebelum {{.PaymentDeadline}}</h1>
					<p>No. Pesanan: <strong>{{.OrderNumber}}</strong></p>
					<p>Yth. {{.CustomerName}},</p>
					<p>Mohon selesaikan pembayaran Anda:</p>
					<div class="details">
						<p>Jumlah Total</p>
						<h1 style="text-align: center; font-size: 22px;" class="amount"><strong>Rp {{.TotalAmount}}</strong></h1>
					</div>
					<div class="important">
						PENTING!!! Pastikan kamu membayar sebelum tenggat waktu ditentukan!
					</div>
				</div>
				<div class="footer">
					<p><strong>Punya pertanyaan?</strong></p>
					<p>Untuk hal mendesak: <a href="tel:{{.CustomerServicePhoneNumber}}">{{.CustomerServicePhoneNumber}}</a></p>
				</div>
			</div>
			</body>
		</html>
	`

	PaymentConfirmEnglish string = `
		<!DOCTYPE html>
			<html lang="id">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>{{.Title}}</title>
				<style>
				body {
					font-family: Arial, sans-serif;
					background-color: #E0E0E0;
					margin: 0;
					padding: 0;
				}
				.container {
					max-width: 600px;
					margin: 20px auto;
					background-color: #fff;
					border-radius: 8px;
					box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
					overflow: hidden;
				}
				.header {
					background-color: #007C99;
					color: white;
					text-align: center;
					padding: 15px;
				}
				.header img {
					max-width: 120px;
					margin-bottom: 10px;
					border-radius: 3px;
					padding: 4px;
				}
				.content {
					padding: 20px;
				}
				.content h1 {
					font-size: 18px;
					color: #333;
				}
				.content p {
					font-size: 16px;
					color: #555;
					margin: 5px 0;
				}
				.details {
					background-color: #f9f9f9;
					border: 1px solid #ddd;
					padding: 15px;
					margin: 20px 0;
					border-radius: 8px;
				}
				.details p {
					margin: 8px 0;
					color: #333;
					text-align: center;
				}
				.details strong {
					display: block;
					margin-bottom: 5px;
					font-size: 22px;
				}
				.payment-info {
					text-align: center;
					margin: 20px 0;
				}
				.payment-info p {
					margin: 5px 0;
				}
				.payment-info .amount {
					font-size: 22px;
					color: #007C99;
					font-weight: bold;
				}
				.important {
					background-color: #FFD700;
					padding: 10px;
					font-size: 14px;
					color: #333;
					border-radius: 5px;
					margin-top: 15px;
					text-align: center;
				}
				.footer {
					padding: 15px;
					font-size: 14px;
					text-align: center;
					color: #555;
					background-color: #f4f4f4;
				}
				.footer a {
					color: #007C99;
					text-decoration: none;
				}
				.footer p {
					margin: 5px 0;
				}
				.payment-channel {
					text-align: center;
					margin-bottom: 10px;
				}
				.payment-channel-name {
					font-size: 18px;
					font-weight: bold;
					color: #333;
					margin-bottom: 5px;
				}
				.logo-container {
					padding: 5px 0;
					text-align: center;
				}
				.logo-container img {
					max-width: 60px;
					display: inline-block;
				}
				.payment-button {
					display: block;
					width: 100%;
					max-width: 300px;
					margin: 20px auto 10px;
					padding: 12px 20px;
					background-color: #007C99;
					color: #FFFFFF !important;
					text-align: center;
					text-decoration: none;
					border-radius: 6px;
					font-weight: bold;
					transition: background-color 0.3s ease;
				}
				.payment-button:hover {
					background-color: #006080;
				}
				</style>
			</head>
			<body>
			<div class="container">
				<div class="header">
					<img src="{{.LogoUrl}}" alt="Logo">
				</div>
				<div class="content">
					<h1>Complete your payment before {{.PaymentDeadline}}</h1>
					<p>Order Number: <strong>{{.OrderNumber}}</strong></p>
					<p>Dear. {{.CustomerName}},</p>
					<p>Please complete your payment:</p>
					<div class="details">
						<p>Total Price</p>
						<h1 style="text-align: center; font-size: 22px;" class="amount"><strong>IDR {{.TotalAmount}}</strong></h1>
						<a href="{{.RedirectUrl}}" class="payment-button">Continue Payment</a>
					</div>
					<div class="important">
						IMPORTANT!!! Make sure you pay before the deadline!
					</div>
				</div>
				<div class="footer">
					<p><strong>ask a question?</strong></p>
					<p>For urgent matters: <a href="tel:{{.CustomerServicePhoneNumber}}">{{.CustomerServicePhoneNumber}}</a></p>
				</div>
			</div>
			</body>
		</html>
	`

	PaymentConfirmAttachmentEnglish string = `
		<!DOCTYPE html>
			<html lang="id">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>{{.Title}}</title>
				<style>
				body {
					font-family: Arial, sans-serif;
					background-color: #E0E0E0;
					margin: 0;
					padding: 0;
				}
				.container {
					max-width: 600px;
					margin: 20px auto;
					background-color: #fff;
					border-radius: 8px;
					box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
					overflow: hidden;
				}
				.header {
					background-color: #007C99;
					color: white;
					text-align: center;
					padding: 15px;
				}
				.header img {
					max-width: 120px;
					margin-bottom: 10px;
					border-radius: 3px;
					padding: 4px;
				}
				.content {
					padding: 20px;
				}
				.content h1 {
					font-size: 18px;
					color: #333;
				}
				.content p {
					font-size: 16px;
					color: #555;
					margin: 5px 0;
				}
				.details {
					background-color: #f9f9f9;
					border: 1px solid #ddd;
					padding: 15px;
					margin: 20px 0;
					border-radius: 8px;
				}
				.details p {
					margin: 8px 0;
					color: #333;
					text-align: center;
				}
				.details strong {
					display: block;
					margin-bottom: 5px;
					font-size: 22px;
				}
				.payment-info {
					text-align: center;
					margin: 20px 0;
				}
				.payment-info p {
					margin: 5px 0;
				}
				.payment-info .amount {
					font-size: 22px;
					color: #007C99;
					font-weight: bold;
				}
				.important {
					background-color: #FFD700;
					padding: 10px;
					font-size: 14px;
					color: #333;
					border-radius: 5px;
					margin-top: 15px;
					text-align: center;
				}
				.footer {
					padding: 15px;
					font-size: 14px;
					text-align: center;
					color: #555;
					background-color: #f4f4f4;
				}
				.footer a {
					color: #007C99;
					text-decoration: none;
				}
				.footer p {
					margin: 5px 0;
				}
				.payment-channel {
					text-align: center;
					margin-bottom: 10px;
				}
				.payment-channel-name {
					font-size: 18px;
					font-weight: bold;
					color: #333;
					margin-bottom: 5px;
				}
				.logo-container {
					padding: 5px 0;
					text-align: center;
				}
				.logo-container img {
					max-width: 60px;
					display: inline-block;
				}
				.payment-button {
					display: block;
					width: 100%;
					max-width: 300px;
					margin: 20px auto 10px;
					padding: 12px 20px;
					background-color: #007C99;
					color: #FFFFFF !important;
					text-align: center;
					text-decoration: none;
					border-radius: 6px;
					font-weight: bold;
					transition: background-color 0.3s ease;
				}
				.payment-button:hover {
					background-color: #006080;
				}
				</style>
			</head>
			<body>
			<div class="container">
				<div class="header">
					<img src="{{.LogoUrl}}" alt="Logo">
				</div>
				<div class="content">
					<h1>Complete your payment before {{.PaymentDeadline}}</h1>
					<p>Order Number: <strong>{{.OrderNumber}}</strong></p>
					<p>Dear. {{.CustomerName}},</p>
					<p>Please complete your payment:</p>
					<div class="details">
						<p>Total Price</p>
						<h1 style="text-align: center; font-size: 22px;" class="amount"><strong>IDR {{.TotalAmount}}</strong></h1>
						<a href="{{.RedirectUrl}}" class="payment-button">Continue Payment</a>
					</div>
					<div class="important">
						IMPORTANT!!! Make sure you pay before the deadline!
					</div>
				</div>
				<div class="footer">
					<p><strong>ask a question?</strong></p>
					<p>For urgent matters: <a href="tel:{{.CustomerServicePhoneNumber}}">{{.CustomerServicePhoneNumber}}</a></p>
				</div>
			</div>
			</body>
		</html>
	`
)
