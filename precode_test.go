package gotestify

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	// Общее количество доступных кафе
	totalCount := 4

	// Создаем запрос с параметром count больше, чем доступное количество кафе
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	// Создаем фиктивный объект для записи ответа
	responseRecorder := httptest.NewRecorder()

	// Указываем обработчик для тестирования
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверяем, что статус ответа — 200 OK
	require.Equal(t, http.StatusOK, responseRecorder.Code, "Expected status code 200")

	// Получаем тело ответа
	body := responseRecorder.Body.String()

	// Проверяем, что список кафе содержит ровно `totalCount` элементов
	list := strings.Split(body, ",")
	assert.Len(t, list, totalCount, "Response should contain all available cafes")

	// Ожидаемый список кафе
	expectedCafes := []string{"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"}

	// Проверяем, что список кафе совпадает с ожидаемым
	assert.Equal(t, expectedCafes, list, "Returned cafe list does not match expected")
}

func TestMainHandlerStatusOK(t *testing.T) {
	// Создаем запрос
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	// Создаем фиктивный записывающий ответ
	responseRecorder := httptest.NewRecorder()

	// Указываем обработчик для тестирования
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверяем, что статус ответа — 200 OK
	require.Equal(t, http.StatusOK, responseRecorder.Code, "Expected status code 200")

	// Проверяем, что тело ответа не пустое
	body := responseRecorder.Body.String()
	assert.NotEmpty(t, body, "Response body should not be empty")
}

func TestMainHandlerWhenCityNotSupported(t *testing.T) {
	// Создаем запрос с неподдерживаемым городом
	req := httptest.NewRequest("GET", "/cafe?count=2&city=unknowncity", nil)

	// Создаем фиктивный объект для записи ответа
	responseRecorder := httptest.NewRecorder()

	// Вызываем обработчик
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверяем, что статус ответа — 400 Bad Request
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Expected status code 400")

	// Проверяем, что тело ответа содержит ошибку "wrong city value"
	body := responseRecorder.Body.String()
	assert.Equal(t, "wrong city value", body, "Response body should contain the error message")
}
