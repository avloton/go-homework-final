package templates

const OrderThanks = `
<!DOCTYPE html>
				<html lang="ru">
				<head>
					<meta charset="UTF-8">
					<meta name="viewport" content="width=device-width, initial-scale=1.0">
					<title>Спасибо за заказ! - Пекарня "Ароматная булочная"</title>
					<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.4.0/css/all.min.css">
					<link href="https://fonts.googleapis.com/css2?family=Playfair+Display:wght@400;700&family=Roboto:wght@300;400;500&display=swap" rel="stylesheet">
					<style>
						* {
							margin: 0;
							padding: 0;
							box-sizing: border-box;
						}

						:root {
							--primary-color: #f8b595;
							--secondary-color: #d89a6d;
							--dark-color: #5c3d2e;
							--light-color: #fff8f0;
							--text-color: #333;
							--shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
							--border-radius: 10px;
						}

						body {
							font-family: 'Roboto', sans-serif;
							line-height: 1.6;
							color: var(--text-color);
							background-color: var(--light-color);
							min-height: 100vh;
							display: flex;
							flex-direction: column;
							justify-content: center;
							align-items: center;
							padding: 20px;
							background-image: 
								radial-gradient(circle at 10% 20%, rgba(248, 181, 149, 0.05) 0%, transparent 20%),
								radial-gradient(circle at 90% 80%, rgba(216, 154, 109, 0.05) 0%, transparent 20%);
						}

						h1, h2, h3, h4 {
							font-family: 'Playfair Display', serif;
							font-weight: 700;
							margin-bottom: 1rem;
						}

						h2 {
							font-size: 2rem;
							color: var(--dark-color);
						}

						.thank-you-container {
							max-width: 600px;
							width: 100%;
							text-align: center;
						}

						.thank-you-card {
							background-color: white;
							padding: 3rem 2.5rem;
							border-radius: var(--border-radius);
							box-shadow: var(--shadow);
							margin-bottom: 2rem;
							position: relative;
							overflow: hidden;
						}

						.thank-you-card::before {
							content: '';
							position: absolute;
							top: 0;
							left: 0;
							right: 0;
							height: 5px;
							background: linear-gradient(90deg, var(--primary-color), var(--secondary-color));
						}

						.thank-you-icon {
							font-size: 4rem;
							color: var(--secondary-color);
							margin-bottom: 1.5rem;
							animation: bounce 1s ease-in-out;
						}

						@keyframes bounce {
							0%, 20%, 50%, 80%, 100% {transform: translateY(0);}
							40% {transform: translateY(-15px);}
							60% {transform: translateY(-7px);}
						}

						.thank-you-message h2 {
							margin-bottom: 1.5rem;
							font-size: 2.2rem;
						}

						.thank-you-message p {
							font-size: 1.1rem;
							margin-bottom: 2rem;
							color: #555;
							line-height: 1.7;
						}

						.order-details {
							background-color: rgba(248, 181, 149, 0.1);
							padding: 1.5rem;
							border-radius: var(--border-radius);
							margin: 2rem 0;
							border-left: 4px solid var(--secondary-color);
							text-align: left;
						}

						.order-details h3 {
							font-size: 1.3rem;
							margin-bottom: 1rem;
							color: var(--dark-color);
							display: flex;
							align-items: center;
							gap: 0.5rem;
						}

						.order-details h3 i {
							color: var(--secondary-color);
						}

						.order-details ul {
							list-style: none;
							padding-left: 0;
						}

						.order-details li {
							margin-bottom: 0.5rem;
							padding-left: 1.5rem;
							position: relative;
						}

						.order-details li:before {
							content: '✓';
							position: absolute;
							left: 0;
							color: var(--secondary-color);
							font-weight: bold;
						}

						.button-container {
							display: flex;
							flex-wrap: wrap;
							gap: 1rem;
							justify-content: center;
							margin-top: 2rem;
						}

						.btn {
							display: inline-flex;
							align-items: center;
							justify-content: center;
							gap: 0.7rem;
							background-color: var(--primary-color);
							color: white;
							padding: 1rem 2rem;
							border-radius: 50px;
							font-weight: 500;
							font-size: 1rem;
							transition: all 0.3s;
							border: none;
							cursor: pointer;
							text-decoration: none;
							min-width: 200px;
							box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
						}

						.btn:hover {
							background-color: var(--secondary-color);
							transform: translateY(-3px);
							box-shadow: 0 6px 12px rgba(0, 0, 0, 0.15);
						}

						.btn-secondary {
							background-color: white;
							color: var(--dark-color);
							border: 2px solid var(--primary-color);
						}

						.btn-secondary:hover {
							background-color: #f9f9f9;
							border-color: var(--secondary-color);
						}

						.btn i {
							font-size: 1.1rem;
						}

						.contact-info {
							margin-top: 2.5rem;
							padding-top: 1.5rem;
							border-top: 1px solid #eee;
							color: #666;
							font-size: 0.95rem;
						}

						.contact-info p {
							margin-bottom: 0.5rem;
							display: flex;
							align-items: center;
							justify-content: center;
							gap: 0.5rem;
						}

						.contact-info i {
							color: var(--secondary-color);
						}

						/* Анимация появления */
						.thank-you-card {
							animation: fadeIn 0.8s ease-out;
						}

						@keyframes fadeIn {
							from { opacity: 0; transform: translateY(20px); }
							to { opacity: 1; transform: translateY(0); }
						}

						/* Декоративный элемент в виде булочки */
						.bread-decoration {
							position: absolute;
							width: 60px;
							height: 60px;
							background-color: var(--light-color);
							border-radius: 50%;
							display: flex;
							align-items: center;
							justify-content: center;
							color: var(--secondary-color);
							font-size: 1.8rem;
							box-shadow: var(--shadow);
							z-index: -1;
						}

						.bread-1 {
							top: 10%;
							left: 5%;
							animation: float 6s ease-in-out infinite;
						}

						.bread-2 {
							bottom: 15%;
							right: 7%;
							animation: float 8s ease-in-out infinite 1s;
						}

						@keyframes float {
							0%, 100% { transform: translateY(0) rotate(0deg); }
							50% { transform: translateY(-15px) rotate(5deg); }
						}

						/* Адаптивность */
						@media (max-width: 768px) {
							.thank-you-card {
								padding: 2rem 1.5rem;
							}
							
							h2 {
								font-size: 1.8rem;
							}
							
							.button-container {
								flex-direction: column;
								align-items: center;
							}
							
							.btn {
								width: 100%;
								max-width: 300px;
							}
							
							.bread-decoration {
								display: none;
							}
						}

						@media (max-width: 480px) {
							body {
								padding: 15px;
							}
							
							.thank-you-card {
								padding: 1.5rem 1rem;
							}
							
							.thank-you-icon {
								font-size: 3rem;
							}
						}
					</style>
				</head>
				<body>
					<!-- Декоративные элементы -->
					<div class="bread-decoration bread-1">
						<i class="fas fa-bread-slice"></i>
					</div>
					
					<div class="bread-decoration bread-2">
						<i class="fas fa-cookie-bite"></i>
					</div>

					<div class="thank-you-container">
						<div class="thank-you-card">
							<div class="thank-you-icon">
								<i class="fas fa-check-circle"></i>
							</div>
							
							<div class="thank-you-message">
								<h2>Спасибо за ваш заказ!</h2>
								<p>Мы получили вашу заявку и уже начали готовить для вас свежую выпечку. Наш оператор свяжется с вами в течение 30 минут для подтверждения деталей заказа.</p>
								
								<div class="order-details">
									<h3><i class="fas fa-info-circle"></i> Что дальше?</h3>
									<ul>
										<li>Подтверждение заказа по телефону</li>
										<li>Приготовление из свежих ингредиентов</li>
										<li>Доставка в выбранное вами время</li>
										<li>СМС-уведомление о статусе заказа</li>
									</ul>
								</div>
								
								<div class="button-container">
									<form action="/show_orders" method="get">
										<button type="submit" class="btn">
											<i class="fas fa-clipboard-list"></i> Отследить заказ
										</button>
									</form>
									
									<form action="/" method="get">
										<button type="submit" class="btn btn-secondary">
											<i class="fas fa-home"></i> Вернуться на главную
										</button>
									</form>
								</div>
								
								<div class="contact-info">
									<p><i class="fas fa-phone"></i> Есть вопросы? Звоните: +7 (999) 123-45-67</p>
									<p><i class="fas fa-clock"></i> Время работы: ежедневно с 7:00 до 21:00</p>
								</div>
							</div>
						</div>
						
						<div style="color: #777; font-size: 0.9rem; margin-top: 1rem;">
							<p>Пекарня "Ароматная булочная" — свежая выпечка с любовью!</p>
						</div>
					</div>

					<script>
						// Простая анимация для иконки
						document.addEventListener('DOMContentLoaded', function() {
							const icon = document.querySelector('.thank-you-icon');
							
							// Добавляем класс с анимацией через небольшой таймаут
							setTimeout(() => {
								icon.style.animation = 'bounce 1s ease-in-out';
							}, 300);
							
							// Добавляем небольшие эффекты при наведении на кнопки
							const buttons = document.querySelectorAll('.btn');
							buttons.forEach(button => {
								button.addEventListener('mouseenter', function() {
									this.style.transform = 'translateY(-3px)';
								});
								
								button.addEventListener('mouseleave', function() {
									this.style.transform = 'translateY(0)';
								});
							});
						});
					</script>
				</body>
				</html>
`

const FeedbackTanks = `
<!DOCTYPE html>
				<html lang="ru">
				<head>
					<meta charset="UTF-8">
					<meta name="viewport" content="width=device-width, initial-scale=1.0">
					<title>Спасибо! - Пекарня "Ароматная булочная"</title>
					<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.4.0/css/all.min.css">
					<link href="https://fonts.googleapis.com/css2?family=Playfair+Display:wght@400;700&family=Roboto:wght@300;400;500&display=swap" rel="stylesheet">
					<style>
						:root {
							--primary-color: #f8b595;
							--secondary-color: #d89a6d;
							--dark-color: #5c3d2e;
							--light-color: #fff8f0;
						}

						body {
							font-family: 'Roboto', sans-serif;
							background-color: var(--light-color);
							margin: 0;
							padding: 0;
							display: flex;
							justify-content: center;
							align-items: center;
							min-height: 100vh;
						}

						.thank-you-message {
							background: white;
							padding: 3rem;
							border-radius: 12px;
							box-shadow: 0 8px 25px rgba(92, 61, 46, 0.1);
							text-align: center;
							max-width: 500px;
							width: 90%;
							border-top: 5px solid var(--secondary-color);
						}

						h2 {
							font-family: 'Playfair Display', serif;
							color: var(--dark-color);
							margin-bottom: 2rem;
							font-size: 1.8rem;
							line-height: 1.4;
						}

						.thank-icon {
							font-size: 3.5rem;
							color: var(--secondary-color);
							margin-bottom: 1.5rem;
							animation: fadeIn 0.8s ease-out;
						}

						button {
							background-color: var(--primary-color);
							color: white;
							border: none;
							padding: 1rem 2rem;
							border-radius: 50px;
							font-family: 'Roboto', sans-serif;
							font-size: 1rem;
							font-weight: 500;
							cursor: pointer;
							transition: all 0.3s;
							display: inline-flex;
							align-items: center;
							gap: 0.7rem;
							box-shadow: 0 4px 10px rgba(248, 181, 149, 0.3);
						}

						button:hover {
							background-color: var(--secondary-color);
							transform: translateY(-2px);
							box-shadow: 0 6px 15px rgba(216, 154, 109, 0.4);
						}

						button i {
							font-size: 1.1rem;
						}

						@keyframes fadeIn {
							from { opacity: 0; transform: translateY(10px); }
							to { opacity: 1; transform: translateY(0); }
						}

						@media (max-width: 600px) {
							.thank-you-message {
								padding: 2rem;
							}
							
							h2 {
								font-size: 1.5rem;
							}
							
							.thank-icon {
								font-size: 2.8rem;
							}
						}
					</style>
				</head>
				<body>
					<div class="thank-you-message">
						<div class="thank-icon">
							<i class="fas fa-comment-dots"></i>
						</div>
						
						<h2>Спасибо за ваше сообщение, мы обязательно ответим!</h2>
						
						<form action="/" method="get">
							<button type="submit">
								<i class="fas fa-home"></i>
								Вернуться на главную
							</button>
						</form>
					</div>
				</body>
				</html>
`