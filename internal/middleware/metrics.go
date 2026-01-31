package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// HTTP500Errors - счетчик HTTP 500 ошибок с метками (labels).
	HTTP500Errors = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_500_errors_total",                                    // Имя метрики в Prometheus
			Help: "Total number of HTTP 500 Internal Server Error responses", // Описание метрики
		},
		[]string{"method", "endpoint"}, // Метки для группировки метрик
	)

	// HTTPRequestsTotal - общее количество HTTP запросов.
	// Позволяет отслеживать нагрузку на приложение по методам, эндпоинтам и статусам.
	HTTPRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "status"},
	)

	// HTTPRequestDuration - длительность обработки HTTP запросов в секундах.
	// Histogram позволяет анализировать производительность и находить медленные запросы.
	HTTPRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: []float64{0.001, 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10}, // Бакеты для гистограммы
		},
		[]string{"method", "endpoint", "status"},
	)

	// HTTP404Errors - счетчик HTTP 404 ошибок (страница не найдена).
	// Помогает отслеживать неверные URL и битые ссылки.
	HTTP404Errors = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_404_errors_total",
			Help: "Total number of HTTP 404 Not Found responses",
		},
		[]string{"method", "endpoint"},
	)

	// HTTP405Errors - счетчик HTTP 405 ошибок (метод не разрешен).
	// Помогает отслеживать неправильное использование HTTP методов.
	HTTP405Errors = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_405_errors_total",
			Help: "Total number of HTTP 405 Method Not Allowed responses",
		},
		[]string{"method", "endpoint"},
	)
)

// MetricsMiddleware создает middleware для сбора метрик 500 ошибок.
// Middleware - это функция, которая оборачивает HTTP обработчик и выполняет
// дополнительную логику до/после обработки запроса.
//
// Параметры:
//   - next: следующий обработчик в цепочке middleware
//
// Возвращает:
//   - http.Handler: новый обработчик с логикой сбора метрик
func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Засекаем время начала обработки запроса для метрики длительности.
		start := time.Now()

		// Создаем обертку над оригинальным ResponseWriter.
		// Это необходимо, чтобы перехватить HTTP статус код ответа,
		// так как стандартный ResponseWriter не позволяет узнать статус после записи.
		rw := &responseWriter{
			ResponseWriter: w,             // Встраиваем оригинальный ResponseWriter
			statusCode:     http.StatusOK, // Статус по умолчанию (200), если не установлен явно
		}

		// Вызываем следующий обработчик в цепочке.
		// Передаем обертку rw вместо оригинального w, чтобы перехватить статус код.
		next.ServeHTTP(rw, r)

		// Вычисляем длительность обработки запроса.
		duration := time.Since(start).Seconds()

		// Получаем статус код и конвертируем в строку для меток.
		statusCode := rw.statusCode
		status := strconv.Itoa(statusCode)

		// Получаем endpoint (путь запроса) для меток.
		endpoint := r.URL.Path

		// Записываем общую метрику запросов (все запросы с метками method, endpoint, status).
		HTTPRequestsTotal.WithLabelValues(r.Method, endpoint, status).Inc()

		// Записываем метрику длительности запроса.
		HTTPRequestDuration.WithLabelValues(r.Method, endpoint, status).Observe(duration)

		// Записываем специфичные метрики для разных типов ошибок.
		switch statusCode {
		case http.StatusInternalServerError: // 500
			HTTP500Errors.WithLabelValues(r.Method, endpoint).Inc()
		case http.StatusNotFound: // 404
			HTTP404Errors.WithLabelValues(r.Method, endpoint).Inc()
		case http.StatusMethodNotAllowed: // 405
			HTTP405Errors.WithLabelValues(r.Method, endpoint).Inc()
		}
	})
}

// responseWriter - обертка над http.ResponseWriter для перехвата HTTP статус кода.
type responseWriter struct {
	http.ResponseWriter      // Встраиваем оригинальный ResponseWriter для делегирования методов
	statusCode          int  // Сохраняем HTTP статус код ответа
	written             bool // Флаг, указывающий, был ли уже записан заголовок ответа
}

// WriteHeader перехватывает вызов WriteHeader и сохраняет статус код.
// Это позволяет middleware узнать, какой статус код был установлен обработчиком.
func (rw *responseWriter) WriteHeader(code int) {
	// Проверяем, не был ли уже записан заголовок.
	// В HTTP можно вызвать WriteHeader только один раз.
	if !rw.written {
		rw.statusCode = code                // Сохраняем статус код
		rw.written = true                   // Помечаем, что заголовок записан
		rw.ResponseWriter.WriteHeader(code) // Вызываем оригинальный метод
	}
}

// Write перехватывает запись тела ответа.
// Если Write вызывается до WriteHeader, статус код по умолчанию - 200 OK.
func (rw *responseWriter) Write(b []byte) (int, error) {
	// Если заголовок еще не был записан, устанавливаем статус 200 OK.
	// Это стандартное поведение HTTP: если Write вызывается первым, статус = 200.
	if !rw.written {
		rw.statusCode = http.StatusOK // Статус по умолчанию
		rw.written = true             // Помечаем, что началась запись
	}
	// Делегируем запись оригинальному ResponseWriter
	return rw.ResponseWriter.Write(b)
}
