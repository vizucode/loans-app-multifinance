package emailattributes

const (
	EticketTemplateIndonesian string = `
		<!DOCTYPE html>
			<html lang="id">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>E-Tiket</title>
			</head>
			<body style="margin: 0; padding: 0; background-color: #f0f0f0; font-family: Arial, sans-serif;">
				<table cellpadding="0" cellspacing="0" border="0" width="100%" style="max-width: 600px; margin: 0 auto; background-color: #ffffff;">
					<tr>
						<td style="padding: 20px; text-align: center; background-color: #007C99;">
							<img src="{{.LogoUrl}}" alt="Logo Perusahaan" style="max-width: 200px; height: auto; padding: 4px; border-radius: 4px;">
						</td>
					</tr>
					<tr>
						<td style="padding: 20px; text-align: center; background-color: #ffffff;">
							<h1 style="margin: 0; color: #007C99; font-size: 24px;">{{.JourneyType}}</h1>
						</td>
					</tr>
					<tr>
						<td style="padding: 20px;">
							<h2 style="margin: 0 0 10px 0; color: #007C99; border-bottom: 2px solid #007C99; padding-bottom: 5px;">Detail Pemesanan</h2>
							<table cellpadding="0" cellspacing="0" border="0" width="100%">
								<tr>
									<td style="padding: 5px 0;"><strong>ID Transaksi:</strong></td>
									<td style="padding: 5px 0;">{{.BookingCode}}</td>
								</tr>
								<tr>
									<td style="padding: 5px 0;"><strong>Kode Booking:</strong></td>
									<td style="padding: 5px 0;">{{.BookingId}}</td>
								</tr>
								<tr>
									<td style="padding: 5px 0;"><strong>Maskapai Penerbit:</strong></td>
									<td style="padding: 5px 0;">{{.AirlineName}}</td>
								</tr>
							</table>
						</td>
					</tr>
					<tr>
						<td style="padding: 20px;">
							<h2 style="margin: 0 0 10px 0; color: #007C99; border-bottom: 2px solid #007C99; padding-bottom: 5px;">Data Penerbangan</h2>
							<table cellpadding="0" cellspacing="0" border="0" width="100%" style="background-color: #f8f9fa; border-radius: 5px;">
								<tr>
									<td style="padding: 15px;">
										<table cellpadding="0" cellspacing="0" border="0" width="100%">
											<tr>
												<td style="padding: 5px 0;"><strong>Maskapai:</strong></td>
												<td style="padding: 5px 0;">{{.AirlineCode}}</td>
											</tr>
											<tr>
												<td style="padding: 5px 0;"><strong>Kode Penerbangan:</strong></td>
												<td style="padding: 5px 0;">{{.FlightNo}}</td>
											</tr>
											<tr>
												<td style="padding: 5px 0;"><strong>Kelas:</strong></td>
												<td style="padding: 5px 0;">{{.ClassName}} ({{.ClassCode}})</td>
											</tr>
											<tr>
												<td style="padding: 5px 0;"><strong>Keberangkatan:</strong></td>
												<td style="padding: 5px 0;">{{.DepartureCityName}} ({{.DepartureAirportCode}}) - {{.DepartureDateTime}}</td>
											</tr>
											<tr>
												<td style="padding: 5px 0;"><strong>Tiba:</strong></td>
												<td style="padding: 5px 0;">{{.ArrivalCityName}} ({{.ArrivalAirportCode}}) - {{.ArrivalDateTime}}</td>
											</tr>
										</table>
									</td>
								</tr>
							</table>
						</td>
					</tr>
					<tr>
						<td style="padding: 20px;">
							<h2 style="margin: 0 0 10px 0; color: #007C99; border-bottom: 2px solid #007C99; padding-bottom: 5px;">Detail Penerbangan</h2>
							<table cellpadding="0" cellspacing="0" border="0" width="100%" style="border-collapse: collapse;">
								<tr style="background-color: #e8f0fe;">
									<th style="padding: 10px; border: 1px solid #dddddd; text-align: left;">Kode Penerbangan</th>
									<th style="padding: 10px; border: 1px solid #dddddd; text-align: left;">Berangkat</th>
									<th style="padding: 10px; border: 1px solid #dddddd; text-align: left;">Tiba</th>
								</tr>
								{{range $index, $flight := .FlightJourney}}
									<tr>
										<td style="padding: 10px; border: 1px solid #dddddd;">{{.FlightNo}}</td>
										<td style="padding: 10px; border: 1px solid #dddddd;">{{.FlightDeparture}} ({{.IataDeparture}})</td>
										<td style="padding: 10px; border: 1px solid #dddddd;">{{.FlightDestination}} ({{.IataDestination}})</td>
									</tr>
								{{end}}
							</table>
						</td>
					</tr>
					<tr>
						<td style="padding: 20px;">
							<h2 style="margin: 0 0 10px 0; color: #007C99; border-bottom: 2px solid #007C99; padding-bottom: 5px;">Detail Penumpang</h2>
							<table cellpadding="0" cellspacing="0" border="0" width="100%" style="border-collapse: collapse;">
								<tr style="background-color: #e8f0fe;">
									<th style="padding: 10px; border: 1px solid #dddddd; text-align: left;">No</th>
									<th style="padding: 10px; border: 1px solid #dddddd; text-align: left;">Penumpang</th>
									<th style="padding: 10px; border: 1px solid #dddddd; text-align: left;">Tiket</th>
								</tr>
								{{range $index, $passenger := .Passengers}}
									<tr>
										<td style="padding: 10px; border: 1px solid #dddddd;">{{.PassengerNo}}</td>
										<td style="padding: 10px; border: 1px solid #dddddd;">{{.PassengerTitle}} {{.PassengerName}} ({{.PassengerType}})</td>
										<td style="padding: 10px; border: 1px solid #dddddd;">{{.PassengerTicket}} </td>
									</tr>
								{{end}}
							</table>
						</td>
					</tr>
					<tr>
						<td style="padding: 20px;">
							<h2 style="margin: 0 0 10px 0; color: #007C99; border-bottom: 2px solid #007C99; padding-bottom: 5px;">Detail Harga</h2>
							<table cellpadding="0" cellspacing="0" border="0" width="100%">
								<tr>
									<td style="padding: 5px 0;"><strong>Harga Tiket:</strong></td>
									<td style="padding: 5px 0; text-align: right;">IDR {{.TotalOrder}} </td>
								</tr>
								<tr style="border-top: 2px solid #007C99;">
									<td style="padding: 10px 0;"><strong>Total Keseluruhan:</strong></td>
									<td style="padding: 10px 0; text-align: right;"><strong>IDR {{.GrandTotal}} </strong></td>
								</tr>
							</table>
						</td>
					</tr>
					<tr>
						<td style="padding: 20px; text-align: center; background-color: #007C99; color: #ffffff; font-size: 14px;">
							<p style="margin: 0;">Terima kasih telah memilih layanan kami!</p>
							<p style="margin: 10px 0 0 0;">Email ini dikirim secara otomatis. Mohon untuk tidak membalas email ini.</p>
						</td>
					</tr>
				</table>
			</body>
		</html>
	`

	EticketTemplateEnglish string = `
		<!DOCTYPE html>
			<html lang="id">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>E-Tiket</title>
			</head>
			<body style="margin: 0; padding: 0; background-color: #f0f0f0; font-family: Arial, sans-serif;">
				<table cellpadding="0" cellspacing="0" border="0" width="100%" style="max-width: 600px; margin: 0 auto; background-color: #ffffff;">
					<tr>
						<td style="padding: 20px; text-align: center; background-color: #007C99;">
							<img src="{{.LogoUrl}}" alt="Logo Perusahaan" style="max-width: 200px; height: auto; padding: 4px; border-radius: 4px;">
						</td>
					</tr>
					<tr>
						<td style="padding: 20px; text-align: center; background-color: #ffffff;">
							<h1 style="margin: 0; color: #007C99; font-size: 24px;">{{.JourneyType}}</h1>
						</td>
					</tr>
					<tr>
						<td style="padding: 20px;">
							<h2 style="margin: 0 0 10px 0; color: #007C99; border-bottom: 2px solid #007C99; padding-bottom: 5px;">Order Detail</h2>
							<table cellpadding="0" cellspacing="0" border="0" width="100%">
								<tr>
									<td style="padding: 5px 0;"><strong>Transaction ID:</strong></td>
									<td style="padding: 5px 0;">{{.BookingCode}}</td>
								</tr>
								<tr>
									<td style="padding: 5px 0;"><strong>Booking Code:</strong></td>
									<td style="padding: 5px 0;">{{.BookingId}}</td>
								</tr>
								<tr>
									<td style="padding: 5px 0;"><strong>Airline Name:</strong></td>
									<td style="padding: 5px 0;">{{.AirlineName}}</td>
								</tr>
							</table>
						</td>
					</tr>
					<tr>
						<td style="padding: 20px;">
							<h2 style="margin: 0 0 10px 0; color: #007C99; border-bottom: 2px solid #007C99; padding-bottom: 5px;">Flight Data</h2>
							<table cellpadding="0" cellspacing="0" border="0" width="100%" style="background-color: #f8f9fa; border-radius: 5px;">
								<tr>
									<td style="padding: 15px;">
										<table cellpadding="0" cellspacing="0" border="0" width="100%">
											<tr>
												<td style="padding: 5px 0;"><strong>Airline:</strong></td>
												<td style="padding: 5px 0;">{{.AirlineCode}}</td>
											</tr>
											<tr>
												<td style="padding: 5px 0;"><strong>Flight Number:</strong></td>
												<td style="padding: 5px 0;">{{.FlightNo}}</td>
											</tr>
											<tr>
												<td style="padding: 5px 0;"><strong>Classes:</strong></td>
												<td style="padding: 5px 0;">{{.ClassName}} ({{.ClassCode}})</td>
											</tr>
											<tr>
												<td style="padding: 5px 0;"><strong>Departure:</strong></td>
												<td style="padding: 5px 0;">{{.DepartureCityName}} ({{.DepartureAirportCode}}) - {{.DepartureDateTime}}</td>
											</tr>
											<tr>
												<td style="padding: 5px 0;"><strong>Arrival:</strong></td>
												<td style="padding: 5px 0;">{{.ArrivalCityName}} ({{.ArrivalAirportCode}}) - {{.ArrivalDateTime}}</td>
											</tr>
										</table>
									</td>
								</tr>
							</table>
						</td>
					</tr>
					<tr>
						<td style="padding: 20px;">
							<h2 style="margin: 0 0 10px 0; color: #007C99; border-bottom: 2px solid #007C99; padding-bottom: 5px;">Flight Details</h2>
							<table cellpadding="0" cellspacing="0" border="0" width="100%" style="border-collapse: collapse;">
								<tr style="background-color: #e8f0fe;">
									<th style="padding: 10px; border: 1px solid #dddddd; text-align: left;">Flight Number</th>
									<th style="padding: 10px; border: 1px solid #dddddd; text-align: left;">Departure</th>
									<th style="padding: 10px; border: 1px solid #dddddd; text-align: left;">Arrival</th>
								</tr>
								{{range $index, $flight := .FlightJourney}}
									<tr>
										<td style="padding: 10px; border: 1px solid #dddddd;">{{.FlightNo}}</td>
										<td style="padding: 10px; border: 1px solid #dddddd;">{{.FlightDeparture}} ({{.IataDeparture}})</td>
										<td style="padding: 10px; border: 1px solid #dddddd;">{{.FlightDestination}} ({{.IataDestination}})</td>
									</tr>
								{{end}}
							</table>
						</td>
					</tr>
					<tr>
						<td style="padding: 20px;">
							<h2 style="margin: 0 0 10px 0; color: #007C99; border-bottom: 2px solid #007C99; padding-bottom: 5px;">Passenger Data</h2>
							<table cellpadding="0" cellspacing="0" border="0" width="100%" style="border-collapse: collapse;">
								<tr style="background-color: #e8f0fe;">
									<th style="padding: 10px; border: 1px solid #dddddd; text-align: left;">No</th>
									<th style="padding: 10px; border: 1px solid #dddddd; text-align: left;">Passenger</th>
									<th style="padding: 10px; border: 1px solid #dddddd; text-align: left;">Ticket Number</th>
								</tr>
								{{range $index, $passenger := .Passengers}}
									<tr>
									  
										<td style="padding: 10px; border: 1px solid #dddddd;">{{.PassengerNo}}</td>
										<td style="padding: 10px; border: 1px solid #dddddd;">{{.PassengerTitle}} {{.PassengerName}} ({{.PassengerType}})</td>
										<td style="padding: 10px; border: 1px solid #dddddd;">{{.PassengerTicket}} </td>
									</tr>
								{{end}}
							</table>
						</td>
					</tr>
					<tr>
						<td style="padding: 20px;">
							<h2 style="margin: 0 0 10px 0; color: #007C99; border-bottom: 2px solid #007C99; padding-bottom: 5px;">Price Details</h2>
							<table cellpadding="0" cellspacing="0" border="0" width="100%">
								<tr>
									<td style="padding: 5px 0;"><strong>Ticket Price:</strong></td>
									<td style="padding: 5px 0; text-align: right;">IDR {{.TotalOrder}} </td>
								</tr>
								<tr style="border-top: 2px solid #007C99;">
									<td style="padding: 10px 0;"><strong>Grand Total:</strong></td>
									<td style="padding: 10px 0; text-align: right;"><strong>IDR {{.GrandTotal}} </strong></td>
								</tr>
							</table>
						</td>
					</tr>
					<tr>
						<td style="padding: 20px; text-align: center; background-color: #007C99; color: #ffffff; font-size: 14px;">
							<p style="margin: 0;">Thank you for choosing our service!</p>
							<p style="margin: 10px 0 0 0;">This email was sent automatically. Please do not reply to this email.</p>
						</td>
					</tr>
				</table>
			</body>
			</html>
	`

	EticketTemplateRoundTripIndonesian string = `
		<!DOCTYPE html>
			<html lang="id">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>E-Tiket</title>
			</head>
			<body style="margin: 0; padding: 0; background-color: #f0f0f0; font-family: Arial, sans-serif;">
				<table cellpadding="0" cellspacing="0" border="0" width="100%" style="max-width: 600px; margin: 0 auto; background-color: #ffffff;">
					<tr>
						<td style="padding: 20px; text-align: center; background-color: #007C99;">
							<img src="{{.LogoUrl}}" alt="Logo Perusahaan" style="max-width: 200px; height: auto; padding: 4px; border-radius: 4px;">
						</td>
					</tr>
					<tr>
						<td style="padding: 20px; text-align: center; background-color: #ffffff;">
							<h1 style="margin: 0; color: #007C99; font-size: 24px;">{{.DepartureJourneyType}}</h1>
						</td>
					</tr>
					<tr>
						<td style="padding: 20px;">
							<h2 style="margin: 0 0 10px 0; color: #007C99; border-bottom: 2px solid #007C99; padding-bottom: 5px;">Detail Pemesanan</h2>
							<table cellpadding="0" cellspacing="0" border="0" width="100%">
								<tr>
									<td style="padding: 5px 0;"><strong>ID Transaksi:</strong></td>
									<td style="padding: 5px 0;">{{.BookingCode}}</td>
								</tr>
								<tr>
									<td style="padding: 5px 0;"><strong>Kode Booking:</strong></td>
									<td style="padding: 5px 0;">{{.BookingId}}</td>
								</tr>
								<tr>
									<td style="padding: 5px 0;"><strong>Maskapai Penerbit:</strong></td>
									<td style="padding: 5px 0;">{{.AirlineName}}</td>
								</tr>
							</table>
						</td>
					</tr>
					<tr>
						<td style="padding: 20px;">
							<h2 style="margin: 0 0 10px 0; color: #007C99; border-bottom: 2px solid #007C99; padding-bottom: 5px;">Data Penerbangan</h2>
							<table cellpadding="0" cellspacing="0" border="0" width="100%" style="background-color: #f8f9fa; border-radius: 5px;">
								<tr>
									<td style="padding: 15px;">
										<table cellpadding="0" cellspacing="0" border="0" width="100%">
											<tr>
												<td style="padding: 5px 0;"><strong>Maskapai:</strong></td>
												<td style="padding: 5px 0;">{{.DepartureAirlineCode}}</td>
											</tr>
											<tr>
												<td style="padding: 5px 0;"><strong>Kode Penerbangan:</strong></td>
												<td style="padding: 5px 0;">{{.DepartureFlightNo}}</td>
											</tr>
											<tr>
												<td style="padding: 5px 0;"><strong>Kelas:</strong></td>
												<td style="padding: 5px 0;">{{.DepartureClassName}} ({{.DepartureClassCode}})</td>
											</tr>
											<tr>
												<td style="padding: 5px 0;"><strong>Keberangkatan:</strong></td>
												<td style="padding: 5px 0;">{{.DepartureCityName}} ({{.DepartureAirportCode}}) - {{.DepartureDateTime}}</td>
											</tr>
											<tr>
												<td style="padding: 5px 0;"><strong>Tiba:</strong></td>
												<td style="padding: 5px 0;">{{.DepartureArrivalCityName}} ({{.DepartureArrivalAirportCode}}) - {{.DepartureArrivalDateTime}}</td>
											</tr>
										</table>
									</td>
								</tr>
							</table>
						</td>
					</tr>
					<tr>
						<td style="padding: 20px;">
							<h2 style="margin: 0 0 10px 0; color: #007C99; border-bottom: 2px solid #007C99; padding-bottom: 5px;">Detail Penerbangan</h2>
							<table cellpadding="0" cellspacing="0" border="0" width="100%" style="border-collapse: collapse;">
								<tr style="background-color: #e8f0fe;">
									<th style="padding: 10px; border: 1px solid #dddddd; text-align: left;">Kode Penerbangan</th>
									<th style="padding: 10px; border: 1px solid #dddddd; text-align: left;">Berangkat</th>
									<th style="padding: 10px; border: 1px solid #dddddd; text-align: left;">Tiba</th>
								</tr>
								{{range $index, $flight := .DepartureFlightJourney}}
									<tr>
										<td style="padding: 10px; border: 1px solid #dddddd;">{{.FlightNo}}</td>
										<td style="padding: 10px; border: 1px solid #dddddd;">{{.FlightDeparture}} ({{.IataDeparture}})</td>
										<td style="padding: 10px; border: 1px solid #dddddd;">{{.FlightDestination}} ({{.IataDestination}})</td>
									</tr>
								{{end}}
							</table>
						</td>
					</tr>
					<tr>
						<td style="padding: 20px;">
							<h2 style="margin: 0 0 10px 0; color: #007C99; border-bottom: 2px solid #007C99; padding-bottom: 5px;">Detail Penumpang</h2>
							<table cellpadding="0" cellspacing="0" border="0" width="100%" style="border-collapse: collapse;">
								<tr style="background-color: #e8f0fe;">
									<th style="padding: 10px; border: 1px solid #dddddd; text-align: left;">No</th>
									<th style="padding: 10px; border: 1px solid #dddddd; text-align: left;">Penumpang</th>
									<th style="padding: 10px; border: 1px solid #dddddd; text-align: left;">Tiket</th>
								</tr>
								{{range $index, $passenger := .Passengers}}
									<tr>
										<td style="padding: 10px; border: 1px solid #dddddd;">{{.PassengerNo}}</td>
										<td style="padding: 10px; border: 1px solid #dddddd;">{{.PassengerTitle}} {{.PassengerName}} ({{.PassengerType}})</td>
										<td style="padding: 10px; border: 1px solid #dddddd;">{{.PassengerTicket}} </td>
									</tr>
								{{end}}
							</table>
						</td>
					</tr>
					<tr>
						<td style="padding: 20px;">
							<h2 style="margin: 0 0 10px 0; color: #007C99; border-bottom: 2px solid #007C99; padding-bottom: 5px;">Detail Harga</h2>
							<table cellpadding="0" cellspacing="0" border="0" width="100%">
								<tr>
									<td style="padding: 5px 0;"><strong>Harga Tiket:</strong></td>
									<td style="padding: 5px 0; text-align: right;">IDR {{.TotalOrder}} </td>
								</tr>
								<tr style="border-top: 2px solid #007C99;">
									<td style="padding: 10px 0;"><strong>Total Keseluruhan:</strong></td>
									<td style="padding: 10px 0; text-align: right;"><strong>IDR {{.GrandTotal}} </strong></td>
								</tr>
							</table>
						</td>
					</tr>
					<tr>
						<td style="padding: 20px; text-align: center; background-color: #007C99; color: #ffffff; font-size: 14px;">
							<p style="margin: 10px 0 0 0;">Terima kasih telah menggunakan layanan kami!</p>
							<p style="margin: 10px 0 0 0;">Email ini dikirim secara otomatis. Mohon untuk tidak membalas email ini.</p>
						</td>
					</tr>
				</table>
			</body>
			</html>
<br>
<br>
<br>
<br>
<br>
<br>
<br>
<br>
			<!DOCTYPE html>
			<html lang="id">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>E-Tiket</title>
			</head>
			<body style="margin: 0; padding: 0; background-color: #f0f0f0; font-family: Arial, sans-serif;">
				<table cellpadding="0" cellspacing="0" border="0" width="100%" style="max-width: 600px; margin: 0 auto; background-color: #ffffff;">
					<tr>
						<td style="padding: 20px; text-align: center; background-color: #007C99;">
							<img src="{{.LogoUrl}}" alt="Logo Perusahaan" style="max-width: 200px; height: auto; padding: 4px; border-radius: 4px;">
						</td>
					</tr>
					<tr>
						<td style="padding: 20px; text-align: center; background-color: #ffffff;">
							<h1 style="margin: 0; color: #007C99; font-size: 24px;">{{.ReturnsJourneyType}}</h1>
						</td>
					</tr>
					<tr>
						<td style="padding: 20px;">
							<h2 style="margin: 0 0 10px 0; color: #007C99; border-bottom: 2px solid #007C99; padding-bottom: 5px;">Detail Pemesanan</h2>
							<table cellpadding="0" cellspacing="0" border="0" width="100%">
								<tr>
									<td style="padding: 5px 0;"><strong>ID Transaksi:</strong></td>
									<td style="padding: 5px 0;">{{.BookingCode}}</td>
								</tr>
								<tr>
									<td style="padding: 5px 0;"><strong>Kode Booking:</strong></td>
									<td style="padding: 5px 0;">{{.BookingId}}</td>
								</tr>
								<tr>
									<td style="padding: 5px 0;"><strong>Maskapai Penerbit:</strong></td>
									<td style="padding: 5px 0;">{{.AirlineName}}</td>
								</tr>
							</table>
						</td>
					</tr>
					<tr>
						<td style="padding: 20px;">
							<h2 style="margin: 0 0 10px 0; color: #007C99; border-bottom: 2px solid #007C99; padding-bottom: 5px;">Data Penerbangan</h2>
							<table cellpadding="0" cellspacing="0" border="0" width="100%" style="background-color: #f8f9fa; border-radius: 5px;">
								<tr>
									<td style="padding: 15px;">
										<table cellpadding="0" cellspacing="0" border="0" width="100%">
											<tr>
												<td style="padding: 5px 0;"><strong>Maskapai:</strong></td>
												<td style="padding: 5px 0;">{{.ReturnAirlineCode}}</td>
											</tr>
											<tr>
												<td style="padding: 5px 0;"><strong>Kode Penerbangan:</strong></td>
												<td style="padding: 5px 0;">{{.ReturnFlightNo}}</td>
											</tr>
											<tr>
												<td style="padding: 5px 0;"><strong>Kelas:</strong></td>
												<td style="padding: 5px 0;">{{.ReturnClassName}} ({{.ReturnClassCode}})</td>
											</tr>
											<tr>
												<td style="padding: 5px 0;"><strong>Keberangkatan:</strong></td>
												<td style="padding: 5px 0;">{{.ReturnCityName}} ({{.ReturnAirportCode}}) - {{.ReturnDepartureDateTime}}</td>
											</tr>
											<tr>
												<td style="padding: 5px 0;"><strong>Tiba:</strong></td>
												<td style="padding: 5px 0;">{{.ReturnArrivalCityName}} ({{.ReturnArrivalAirportCode}}) - {{.ReturnArrivalDateTime}}</td>
											</tr>
										</table>
									</td>
								</tr>
							</table>
						</td>
					</tr>
					<tr>
						<td style="padding: 20px;">
							<h2 style="margin: 0 0 10px 0; color: #007C99; border-bottom: 2px solid #007C99; padding-bottom: 5px;">Detail Penerbangan</h2>
							<table cellpadding="0" cellspacing="0" border="0" width="100%" style="border-collapse: collapse;">
								<tr style="background-color: #e8f0fe;">
									<th style="padding: 10px; border: 1px solid #dddddd; text-align: left;">Kode Penerbangan</th>
									<th style="padding: 10px; border: 1px solid #dddddd; text-align: left;">Berangkat</th>
									<th style="padding: 10px; border: 1px solid #dddddd; text-align: left;">Tiba</th>
								</tr>
								{{range $index, $flight := .ReturnFlightJourney}}
									<tr>
										<td style="padding: 10px; border: 1px solid #dddddd;">{{.FlightNo}}</td>
										<td style="padding: 10px; border: 1px solid #dddddd;">{{.FlightDeparture}} ({{.IataDeparture}})</td>
										<td style="padding: 10px; border: 1px solid #dddddd;">{{.FlightDestination}} ({{.IataDestination}})</td>
									</tr>
								{{end}}
							</table>
						</td>
					</tr>
					<tr>
						<td style="padding: 20px;">
							<h2 style="margin: 0 0 10px 0; color: #007C99; border-bottom: 2px solid #007C99; padding-bottom: 5px;">Detail Penumpang</h2>
							<table cellpadding="0" cellspacing="0" border="0" width="100%" style="border-collapse: collapse;">
								<tr style="background-color: #e8f0fe;">
									<th style="padding: 10px; border: 1px solid #dddddd; text-align: left;">No</th>
									<th style="padding: 10px; border: 1px solid #dddddd; text-align: left;">Penumpang</th>
									<th style="padding: 10px; border: 1px solid #dddddd; text-align: left;">Tiket</th>
								</tr>
								{{range $index, $passenger := .Passengers}}
									<tr>
										<td style="padding: 10px; border: 1px solid #dddddd;">{{.PassengerNo}}</td>
										<td style="padding: 10px; border: 1px solid #dddddd;">{{.PassengerTitle}} {{.PassengerName}} ({{.PassengerType}})</td>
										<td style="padding: 10px; border: 1px solid #dddddd;">{{.PassengerTicket}} </td>
									</tr>
								{{end}}
							</table>
						</td>
					</tr>
					<tr>
						<td style="padding: 20px;">
							<h2 style="margin: 0 0 10px 0; color: #007C99; border-bottom: 2px solid #007C99; padding-bottom: 5px;">Detail Harga</h2>
							<table cellpadding="0" cellspacing="0" border="0" width="100%">
								<tr>
									<td style="padding: 5px 0;"><strong>Harga Tiket:</strong></td>
									<td style="padding: 5px 0; text-align: right;">IDR {{.TotalOrder}} </td>
								</tr>
								<tr style="border-top: 2px solid #007C99;">
									<td style="padding: 10px 0;"><strong>Total Keseluruhan:</strong></td>
									<td style="padding: 10px 0; text-align: right;"><strong>IDR {{.GrandTotal}} </strong></td>
								</tr>
							</table>
						</td>
					</tr>
					<tr>
						<td style="padding: 20px; text-align: center; background-color: #007C99; color: #ffffff; font-size: 14px;">
						 	<p style="margin: 10px 0 0 0;">Terima kasih telah menggunakan layanan kami!</p>
							<p style="margin: 10px 0 0 0;">Email ini dikirim secara otomatis. Mohon untuk tidak membalas email ini.</p>
						</td>
					</tr>
				</table>
			</body>
			</html>
	`

	EticketTemplateRoundTripEnglish string = `
		<!DOCTYPE html>
			<html lang="id">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>E-Tiket</title>
			</head>
			<body style="margin: 0; padding: 0; background-color: #f0f0f0; font-family: Arial, sans-serif;">
				<table cellpadding="0" cellspacing="0" border="0" width="100%" style="max-width: 600px; margin: 0 auto; background-color: #ffffff;">
					<tr>
						<td style="padding: 20px; text-align: center; background-color: #007C99;">
							<img src="{{.LogoUrl}}" alt="Logo Perusahaan" style="max-width: 200px; height: auto; padding: 4px; border-radius: 4px;">
						</td>
					</tr>
					<tr>
						<td style="padding: 20px; text-align: center; background-color: #ffffff;">
							<h1 style="margin: 0; color: #007C99; font-size: 24px;">{{.DepartureJourneyType}}</h1>
						</td>
					</tr>
					<tr>
						<td style="padding: 20px;">
							<h2 style="margin: 0 0 10px 0; color: #007C99; border-bottom: 2px solid #007C99; padding-bottom: 5px;">Order Details</h2>
							<table cellpadding="0" cellspacing="0" border="0" width="100%">
								<tr>
									<td style="padding: 5px 0;"><strong>Transaction ID:</strong></td>
									<td style="padding: 5px 0;">{{.BookingCode}}</td>
								</tr>
								<tr>
									<td style="padding: 5px 0;"><strong>Booking Code:</strong></td>
									<td style="padding: 5px 0;">{{.BookingId}}</td>
								</tr>
								<tr>
									<td style="padding: 5px 0;"><strong>Airline Name:</strong></td>
									<td style="padding: 5px 0;">{{.AirlineName}}</td>
								</tr>
							</table>
						</td>
					</tr>
					<tr>
						<td style="padding: 20px;">
							<h2 style="margin: 0 0 10px 0; color: #007C99; border-bottom: 2px solid #007C99; padding-bottom: 5px;">Flight Data</h2>
							<table cellpadding="0" cellspacing="0" border="0" width="100%" style="background-color: #f8f9fa; border-radius: 5px;">
								<tr>
									<td style="padding: 15px;">
										<table cellpadding="0" cellspacing="0" border="0" width="100%">
											<tr>
												<td style="padding: 5px 0;"><strong>Airline:</strong></td>
												<td style="padding: 5px 0;">{{.DepartureAirlineCode}}</td>
											</tr>
											<tr>
												<td style="padding: 5px 0;"><strong>Flight Number:</strong></td>
												<td style="padding: 5px 0;">{{.DepartureFlightNo}}</td>
											</tr>
											<tr>
												<td style="padding: 5px 0;"><strong>Classes:</strong></td>
												<td style="padding: 5px 0;">{{.DepartureClassName}} ({{.DepartureClassCode}})</td>
											</tr>
											<tr>
												<td style="padding: 5px 0;"><strong>Departure:</strong></td>
												<td style="padding: 5px 0;">{{.DepartureCityName}} ({{.DepartureAirportCode}}) - {{.DepartureDateTime}}</td>
											</tr>
											<tr>
												<td style="padding: 5px 0;"><strong>Arrival:</strong></td>
												<td style="padding: 5px 0;">{{.DepartureArrivalCityName}} ({{.DepartureArrivalAirportCode}}) - {{.DepartureArrivalDateTime}}</td>
											</tr>
										</table>
									</td>
								</tr>
							</table>
						</td>
					</tr>
					<tr>
						<td style="padding: 20px;">
							<h2 style="margin: 0 0 10px 0; color: #007C99; border-bottom: 2px solid #007C99; padding-bottom: 5px;">Flight Details</h2>
							<table cellpadding="0" cellspacing="0" border="0" width="100%" style="border-collapse: collapse;">
								<tr style="background-color: #e8f0fe;">
									<th style="padding: 10px; border: 1px solid #dddddd; text-align: left;">Kode Penerbangan</th>
									<th style="padding: 10px; border: 1px solid #dddddd; text-align: left;">Berangkat</th>
									<th style="padding: 10px; border: 1px solid #dddddd; text-align: left;">Tiba</th>
								</tr>
								{{range $index, $flight := .DepartureFlightJourney}}
									<tr>
										<td style="padding: 10px; border: 1px solid #dddddd;">{{.FlightNo}}</td>
										<td style="padding: 10px; border: 1px solid #dddddd;">{{.FlightDeparture}} ({{.IataDeparture}})</td>
										<td style="padding: 10px; border: 1px solid #dddddd;">{{.FlightDestination}} ({{.IataDestination}})</td>
									</tr>
								{{end}}
							</table>
						</td>
					</tr>
					<tr>
						<td style="padding: 20px;">
							<h2 style="margin: 0 0 10px 0; color: #007C99; border-bottom: 2px solid #007C99; padding-bottom: 5px;">Passenger Details</h2>
							<table cellpadding="0" cellspacing="0" border="0" width="100%" style="border-collapse: collapse;">
								<tr style="background-color: #e8f0fe;">
									<th style="padding: 10px; border: 1px solid #dddddd; text-align: left;">No</th>
									<th style="padding: 10px; border: 1px solid #dddddd; text-align: left;">Penumpang</th>
									<th style="padding: 10px; border: 1px solid #dddddd; text-align: left;">Tiket</th>
								</tr>
								{{range $index, $passenger := .Passengers}}
									<tr>
										<td style="padding: 10px; border: 1px solid #dddddd;">{{.PassengerNo}}</td>
										<td style="padding: 10px; border: 1px solid #dddddd;">{{.PassengerTitle}} {{.PassengerName}} ({{.PassengerType}})</td>
										<td style="padding: 10px; border: 1px solid #dddddd;">{{.PassengerTicket}} </td>
									</tr>
								{{end}}
							</table>
						</td>
					</tr>
					<tr>
						<td style="padding: 20px;">
							<h2 style="margin: 0 0 10px 0; color: #007C99; border-bottom: 2px solid #007C99; padding-bottom: 5px;">Price Details</h2>
							<table cellpadding="0" cellspacing="0" border="0" width="100%">
								<tr>
									<td style="padding: 5px 0;"><strong>Ticket Price:</strong></td>
									<td style="padding: 5px 0; text-align: right;">IDR {{.TotalOrder}} </td>
								</tr>
								<tr style="border-top: 2px solid #007C99;">
									<td style="padding: 10px 0;"><strong>Grand Total:</strong></td>
									<td style="padding: 10px 0; text-align: right;"><strong>IDR {{.GrandTotal}} </strong></td>
								</tr>
							</table>
						</td>
					</tr>
					<tr>
						<td style="padding: 20px; text-align: center; background-color: #007C99; color: #ffffff; font-size: 14px;">
							<p style="margin: 10px 0 0 0;">Thank you for choosing our service!</p>
							<p style="margin: 10px 0 0 0;">This email was sent automatically. Please do not reply to this email.</p>
						</td>
					</tr>
				</table>
			</body>
			</html>
<br>
<br>
<br>
<br>
<br>
<br>
<br>
<br>
			<!DOCTYPE html>
			<html lang="id">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>E-Tiket</title>
			</head>
			<body style="margin: 0; padding: 0; background-color: #f0f0f0; font-family: Arial, sans-serif;">
				<table cellpadding="0" cellspacing="0" border="0" width="100%" style="max-width: 600px; margin: 0 auto; background-color: #ffffff;">
					<tr>
						<td style="padding: 20px; text-align: center; background-color: #007C99;">
							<img src="{{.LogoUrl}}" alt="Logo Perusahaan" style="max-width: 200px; height: auto; padding: 4px; border-radius: 4px;">
						</td>
					</tr>
					<tr>
						<td style="padding: 20px; text-align: center; background-color: #ffffff;">
							<h1 style="margin: 0; color: #007C99; font-size: 24px;">{{.ReturnsJourneyType}}</h1>
						</td>
					</tr>
					<tr>
						<td style="padding: 20px;">
							<h2 style="margin: 0 0 10px 0; color: #007C99; border-bottom: 2px solid #007C99; padding-bottom: 5px;">Order Details</h2>
							<table cellpadding="0" cellspacing="0" border="0" width="100%">
								<tr>
									<td style="padding: 5px 0;"><strong>Transaction ID:</strong></td>
									<td style="padding: 5px 0;">{{.BookingCode}}</td>
								</tr>
								<tr>
									<td style="padding: 5px 0;"><strong>Booking Code:</strong></td>
									<td style="padding: 5px 0;">{{.BookingId}}</td>
								</tr>
								<tr>
									<td style="padding: 5px 0;"><strong>Airline Name:</strong></td>
									<td style="padding: 5px 0;">{{.AirlineName}}</td>
								</tr>
							</table>
						</td>
					</tr>
					<tr>
						<td style="padding: 20px;">
							<h2 style="margin: 0 0 10px 0; color: #007C99; border-bottom: 2px solid #007C99; padding-bottom: 5px;">Flight Data</h2>
							<table cellpadding="0" cellspacing="0" border="0" width="100%" style="background-color: #f8f9fa; border-radius: 5px;">
								<tr>
									<td style="padding: 15px;">
										<table cellpadding="0" cellspacing="0" border="0" width="100%">
											<tr>
												<td style="padding: 5px 0;"><strong>Airline:</strong></td>
												<td style="padding: 5px 0;">{{.ReturnAirlineCode}}</td>
											</tr>
											<tr>
												<td style="padding: 5px 0;"><strong>Flight Number:</strong></td>
												<td style="padding: 5px 0;">{{.ReturnFlightNo}}</td>
											</tr>
											<tr>
												<td style="padding: 5px 0;"><strong>Classes:</strong></td>
												<td style="padding: 5px 0;">{{.ReturnClassName}} ({{.ReturnClassCode}})</td>
											</tr>
											<tr>
												<td style="padding: 5px 0;"><strong>Departure:</strong></td>
												<td style="padding: 5px 0;">{{.ReturnCityName}} ({{.ReturnAirportCode}}) - {{.ReturnDepartureDateTime}}</td>
											</tr>
											<tr>
												<td style="padding: 5px 0;"><strong>Arrival:</strong></td>
												<td style="padding: 5px 0;">{{.ReturnArrivalCityName}} ({{.ReturnArrivalAirportCode}}) - {{.ReturnArrivalDateTime}}</td>
											</tr>
										</table>
									</td>
								</tr>
							</table>
						</td>
					</tr>
					<tr>
						<td style="padding: 20px;">
							<h2 style="margin: 0 0 10px 0; color: #007C99; border-bottom: 2px solid #007C99; padding-bottom: 5px;">Flight Details</h2>
							<table cellpadding="0" cellspacing="0" border="0" width="100%" style="border-collapse: collapse;">
								<tr style="background-color: #e8f0fe;">
									<th style="padding: 10px; border: 1px solid #dddddd; text-align: left;">Kode Penerbangan</th>
									<th style="padding: 10px; border: 1px solid #dddddd; text-align: left;">Berangkat</th>
									<th style="padding: 10px; border: 1px solid #dddddd; text-align: left;">Tiba</th>
								</tr>
								{{range $index, $flight := .ReturnFlightJourney}}
									<tr>
										<td style="padding: 10px; border: 1px solid #dddddd;">{{.FlightNo}}</td>
										<td style="padding: 10px; border: 1px solid #dddddd;">{{.FlightDeparture}} ({{.IataDeparture}})</td>
										<td style="padding: 10px; border: 1px solid #dddddd;">{{.FlightDestination}} ({{.IataDestination}})</td>
									</tr>
								{{end}}
							</table>
						</td>
					</tr>
					<tr>
						<td style="padding: 20px;">
							<h2 style="margin: 0 0 10px 0; color: #007C99; border-bottom: 2px solid #007C99; padding-bottom: 5px;">Passenger Data</h2>
							<table cellpadding="0" cellspacing="0" border="0" width="100%" style="border-collapse: collapse;">
								<tr style="background-color: #e8f0fe;">
									<th style="padding: 10px; border: 1px solid #dddddd; text-align: left;">No</th>
									<th style="padding: 10px; border: 1px solid #dddddd; text-align: left;">Penumpang</th>
									<th style="padding: 10px; border: 1px solid #dddddd; text-align: left;">Tiket</th>
								</tr>
								{{range $index, $passenger := .Passengers}}
									<tr>
										<td style="padding: 10px; border: 1px solid #dddddd;">{{.PassengerNo}}</td>
										<td style="padding: 10px; border: 1px solid #dddddd;">{{.PassengerTitle}} {{.PassengerName}} ({{.PassengerType}})</td>
										<td style="padding: 10px; border: 1px solid #dddddd;">{{.PassengerTicket}} </td>
									</tr>
								{{end}}
							</table>
						</td>
					</tr>
					<tr>
						<td style="padding: 20px;">
							<h2 style="margin: 0 0 10px 0; color: #007C99; border-bottom: 2px solid #007C99; padding-bottom: 5px;">Price Details</h2>
							<table cellpadding="0" cellspacing="0" border="0" width="100%">
								<tr>
									<td style="padding: 5px 0;"><strong>Ticket Price:</strong></td>
									<td style="padding: 5px 0; text-align: right;">IDR {{.TotalOrder}} </td>
								</tr>
								<tr style="border-top: 2px solid #007C99;">
									<td style="padding: 10px 0;"><strong>Grand Total:</strong></td>
									<td style="padding: 10px 0; text-align: right;"><strong>IDR {{.GrandTotal}} </strong></td>
								</tr>
							</table>
						</td>
					</tr>
					<tr>
						<td style="padding: 20px; text-align: center; background-color: #007C99; color: #ffffff; font-size: 14px;">
							<p style="margin: 10px 0 0 0;">Thank you for choosing our service!</p>
							<p style="margin: 10px 0 0 0;">This email was sent automatically. Please do not reply to this email.</p>
						</td>
					</tr>
				</table>
			</body>
			</html>
	`
)
